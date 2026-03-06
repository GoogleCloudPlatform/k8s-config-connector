// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package backup

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/powertools/kubecli"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
	"google.golang.org/api/option"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/yaml"

	"cloud.google.com/go/storage"
)

type createOptions struct {
	kubecli.ClusterOptions
	cluster   string
	location  string
	project   string
	bucket    string
	namespace string
}

func NewCreateCmd() *cobra.Command {
	options := &createOptions{}

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Trigger an immediate backup",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runCreate(cmd, options)
		},
	}

	options.ClusterOptions.AddFlags(cmd)
	cmd.Flags().StringVar(&options.cluster, "cluster", "", "Name of the cluster")
	cmd.Flags().StringVar(&options.location, "location", "", "Region of the cluster")
	cmd.Flags().StringVar(&options.project, "project", "", "GCP project ID")
	cmd.Flags().StringVar(&options.bucket, "bucket", "", "GCS bucket name for backups")
	cmd.Flags().StringVar(&options.namespace, "namespace", "cnrm-system", "Namespace where Config Connector is installed")

	return cmd
}

func runCreate(cmd *cobra.Command, options *createOptions) error {
	ctx := cmd.Context()
	if options.bucket == "" {
		return fmt.Errorf("--bucket is required")
	}

	kubeClient, err := kubecli.NewClient(ctx, options.ClusterOptions)
	if err != nil {
		return fmt.Errorf("creating kubernetes client: %w", err)
	}

	var gcsOptions []option.ClientOption
	if httpClient := ctx.Value(oauth2.HTTPClient); httpClient != nil {
		gcsOptions = append(gcsOptions, option.WithHTTPClient(httpClient.(*http.Client)))
	}
	gcsClient, err := storage.NewClient(ctx, gcsOptions...)
	if err != nil {
		return fmt.Errorf("creating GCS client: %w", err)
	}
	defer gcsClient.Close()

	timestamp := time.Now().UTC().Format("2006-01-02-15-04-05")
	clusterName := options.cluster
	if clusterName == "" {
		clusterName = "default-cluster"
	}

	fmt.Fprintf(cmd.OutOrStdout(), "Starting backup for cluster %q to bucket %q (timestamp: %s)...\n", clusterName, options.bucket, timestamp)

	// Discover KCC resources
	_, resourceLists, err := kubeClient.DiscoveryClient.ServerGroupsAndResources()
	if err != nil {
		return fmt.Errorf("discovering server resources: %w", err)
	}

	stats := make(map[string]int)
	backedUpGKs := make(map[string]bool)

	for _, resourceList := range resourceLists {
		gv, err := schema.ParseGroupVersion(resourceList.GroupVersion)
		if err != nil {
			continue
		}

		if !strings.HasSuffix(gv.Group, ".cnrm.cloud.google.com") {
			continue
		}

		for _, resource := range resourceList.APIResources {
			if strings.Contains(resource.Name, "/status") || strings.Contains(resource.Name, "/finalizers") {
				continue
			}

			gk := fmt.Sprintf("%s/%s", gv.Group, resource.Kind)
			if backedUpGKs[gk] {
				continue
			}

			gvk := gv.WithKind(resource.Kind)
			count, err := backupResource(ctx, kubeClient, gcsClient, options.bucket, clusterName, timestamp, gvk)
			if err != nil {
				fmt.Fprintf(cmd.ErrOrStderr(), "Warning: failed to backup %v: %v\n", gvk, err)
			}
			if count > 0 {
				stats[gk] = count
			}
			backedUpGKs[gk] = true
		}
	}

	if err := writeSummary(ctx, gcsClient, options.bucket, clusterName, timestamp, stats); err != nil {
		fmt.Fprintf(cmd.ErrOrStderr(), "Warning: failed to write summary.json: %v\n", err)
	}

	total := 0
	for _, count := range stats {
		total += count
	}
	fmt.Fprintf(cmd.OutOrStdout(), "\nBackup completed successfully. Total resources backed up: %d\n", total)

	return nil
}

func writeSummary(ctx context.Context, gcsClient *storage.Client, bucket, cluster, timestamp string, stats map[string]int) error {
	data, err := json.MarshalIndent(stats, "", "  ")
	if err != nil {
		return err
	}

	objectName := fmt.Sprintf("%s/%s/summary.json", cluster, timestamp)
	wc := gcsClient.Bucket(bucket).Object(objectName).NewWriter(ctx)
	if _, err := wc.Write(data); err != nil {
		_ = wc.Close()
		return err
	}
	return wc.Close()
}

func backupResource(ctx context.Context, kubeClient *kubecli.Client, gcsClient *storage.Client, bucket, cluster, timestamp string, gvk schema.GroupVersionKind) (int, error) {
	limit := int64(500)
	continueToken := ""
	count := 0

	for {
		list := &unstructured.UnstructuredList{}
		list.SetGroupVersionKind(gvk)

		listOptions := []client.ListOption{
			client.Limit(limit),
		}
		if continueToken != "" {
			listOptions = append(listOptions, client.Continue(continueToken))
		}

		if err := kubeClient.List(ctx, list, listOptions...); err != nil {
			return count, err
		}

		for _, item := range list.Items {
			if err := backupObject(ctx, gcsClient, bucket, cluster, timestamp, item); err != nil {
				// We don't have the cmd here, so we'll just use fmt.Printf for now or pass cmd around.
				// For now, let's keep it simple.
				fmt.Printf("Warning: failed to backup object %s/%s (%s): %v\n", item.GetNamespace(), item.GetName(), item.GetKind(), err)
				continue
			}
			count++
		}

		continueToken = list.GetContinue()
		if continueToken == "" {
			break
		}
	}

	return count, nil
}

func backupObject(ctx context.Context, gcsClient *storage.Client, bucket, cluster, timestamp string, obj unstructured.Unstructured) error {
	sanitizeObject(&obj)

	data, err := yaml.Marshal(obj.Object)
	if err != nil {
		return err
	}

	name := obj.GetName()
	namespace := obj.GetNamespace()
	kind := strings.ToLower(obj.GetKind())

	objectName := fmt.Sprintf("%s/%s/%s/%s/%s.yaml", cluster, timestamp, namespace, kind, name)

	wc := gcsClient.Bucket(bucket).Object(objectName).NewWriter(ctx)
	if _, err := wc.Write(data); err != nil {
		_ = wc.Close()
		return err
	}
	return wc.Close()
}

func sanitizeObject(obj *unstructured.Unstructured) {
	unstructured.RemoveNestedField(obj.Object, "metadata", "uid")
	unstructured.RemoveNestedField(obj.Object, "metadata", "resourceVersion")
	unstructured.RemoveNestedField(obj.Object, "metadata", "generation")
	unstructured.RemoveNestedField(obj.Object, "metadata", "managedFields")
	unstructured.RemoveNestedField(obj.Object, "metadata", "creationTimestamp")
	unstructured.RemoveNestedField(obj.Object, "metadata", "ownerReferences")
}
