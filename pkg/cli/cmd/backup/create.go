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
	"fmt"
	"strings"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/powertools/kubecli"
	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/yaml"

	"cloud.google.com/go/storage"
)

type createOptions struct {
	cluster  string
	location string
	project  string
	bucket   string
}

func NewCreateCmd() *cobra.Command {
	options := &createOptions{}

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Trigger an immediate backup",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runCreate(cmd.Context(), options)
		},
	}

	cmd.Flags().StringVar(&options.cluster, "cluster", "", "Name of the cluster")
	cmd.Flags().StringVar(&options.location, "location", "", "Region of the cluster")
	cmd.Flags().StringVar(&options.project, "project", "", "GCP project ID")
	cmd.Flags().StringVar(&options.bucket, "bucket", "", "GCS bucket name for backups")

	return cmd
}

func runCreate(ctx context.Context, options *createOptions) error {
	if options.bucket == "" {
		return fmt.Errorf("--bucket is required")
	}

	kubeClient, err := kubecli.NewClient(ctx, kubecli.ClusterOptions{})
	if err != nil {
		return fmt.Errorf("creating kubernetes client: %w", err)
	}

	gcsClient, err := storage.NewClient(ctx)
	if err != nil {
		return fmt.Errorf("creating GCS client: %w", err)
	}
	defer gcsClient.Close()

	timestamp := time.Now().Format("2006-01-02-15-04-05")

	// Discover KCC resources
	resourceLists, err := kubeClient.DiscoveryClient.ServerPreferredResources()
	if err != nil {
		return fmt.Errorf("discovering server resources: %w", err)
	}

	for _, resourceList := range resourceLists {
		gv, err := schema.ParseGroupVersion(resourceList.GroupVersion)
		if err != nil {
			continue
		}

		if !strings.HasSuffix(gv.Group, ".cnrm.cloud.google.com") {
			continue
		}

		for _, resource := range resourceList.APIResources {
			// Skip subresources
			if strings.Contains(resource.Name, "/") {
				continue
			}

			// Skip resources that don't support list
			supportsList := false
			for _, verb := range resource.Verbs {
				if verb == "list" {
					supportsList = true
					break
				}
			}
			if !supportsList {
				continue
			}

			gvk := gv.WithKind(resource.Kind)
			if err := backupResource(ctx, kubeClient, gcsClient, options.bucket, timestamp, gvk); err != nil {
				fmt.Printf("Warning: failed to backup %v: %v\n", gvk, err)
			}
		}
	}

	return nil
}

func backupResource(ctx context.Context, kubeClient *kubecli.Client, gcsClient *storage.Client, bucket, timestamp string, gvk schema.GroupVersionKind) error {
	list := &unstructured.UnstructuredList{}
	list.SetGroupVersionKind(gvk)

	if err := kubeClient.List(ctx, list); err != nil {
		return err
	}

	for _, item := range list.Items {
		if err := backupObject(ctx, gcsClient, bucket, timestamp, item); err != nil {
			return err
		}
	}

	return nil
}

func backupObject(ctx context.Context, gcsClient *storage.Client, bucket, timestamp string, obj unstructured.Unstructured) error {
	// Sanitization - Remove live fields
	unstructured.RemoveNestedField(obj.Object, "metadata", "uid")
	unstructured.RemoveNestedField(obj.Object, "metadata", "resourceVersion")
	unstructured.RemoveNestedField(obj.Object, "metadata", "generation")
	unstructured.RemoveNestedField(obj.Object, "metadata", "managedFields")
	unstructured.RemoveNestedField(obj.Object, "metadata", "creationTimestamp")

	// Remove status
	unstructured.RemoveNestedField(obj.Object, "status")

	// Remove common system annotations
	annotations := obj.GetAnnotations()
	if annotations != nil {
		delete(annotations, "kubectl.kubernetes.io/last-applied-configuration")
		delete(annotations, "deployment.kubernetes.io/revision")
		delete(annotations, "cnrm.cloud.google.com/state-into-spec") // Sometimes added by KCC
		if len(annotations) == 0 {
			unstructured.RemoveNestedField(obj.Object, "metadata", "annotations")
		} else {
			obj.SetAnnotations(annotations)
		}
	}

	data, err := yaml.Marshal(obj.Object)
	if err != nil {
		return err
	}

	namespace := obj.GetNamespace()
	if namespace == "" {
		namespace = "_cluster_scoped"
	}
	kind := strings.ToLower(obj.GetKind())
	name := obj.GetName()

	objectName := fmt.Sprintf("%s/%s/%s/%s.yaml", timestamp, namespace, kind, name)

	wc := gcsClient.Bucket(bucket).Object(objectName).NewWriter(ctx)
	if _, err := wc.Write(data); err != nil {
		return err
	}
	return wc.Close()
}
