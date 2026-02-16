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
	"io"
	"sort"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/powertools/kubecli"
	"github.com/spf13/cobra"
	"google.golang.org/api/iterator"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/yaml"

	"cloud.google.com/go/storage"
)

type restoreOptions struct {
	targetCluster         string
	targetClusterLocation string
	sourceBucket          string
	backupTimestamp       string
	project               string
	dryRun                bool
}

func NewRestoreCmd() *cobra.Command {
	options := &restoreOptions{}

	cmd := &cobra.Command{
		Use:   "restore",
		Short: "Restore resources from GCS to a target cluster",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runRestore(cmd.Context(), options)
		},
	}

	cmd.Flags().StringVar(&options.targetCluster, "target-cluster", "", "Name of the target cluster")
	cmd.Flags().StringVar(&options.targetClusterLocation, "target-cluster-location", "", "Location of the target cluster")
	cmd.Flags().StringVar(&options.sourceBucket, "source-bucket", "", "Source GCS bucket name")
	cmd.Flags().StringVar(&options.backupTimestamp, "backup-timestamp", "", "Backup timestamp (YYYY-MM-DD-HH-MM-SS)")
	cmd.Flags().StringVar(&options.project, "project", "", "GCP project ID")
	cmd.Flags().BoolVar(&options.dryRun, "dry-run", false, "Perform a dry-run validation")

	return cmd
}

func runRestore(ctx context.Context, options *restoreOptions) error {
	if options.sourceBucket == "" {
		return fmt.Errorf("--source-bucket is required")
	}
	if options.backupTimestamp == "" {
		return fmt.Errorf("--backup-timestamp is required")
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

	fmt.Printf("Loading backup from gs://%s/%s/...\n", options.sourceBucket, options.backupTimestamp)

	prefix := options.backupTimestamp + "/"
	it := gcsClient.Bucket(options.sourceBucket).Objects(ctx, &storage.Query{Prefix: prefix})

	var objects []*unstructured.Unstructured
	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return fmt.Errorf("iterating GCS objects: %w", err)
		}

		if !strings.HasSuffix(attrs.Name, ".yaml") {
			continue
		}

		obj, err := loadObject(ctx, gcsClient, options.sourceBucket, attrs.Name)
		if err != nil {
			fmt.Printf("Warning: failed to load %s: %v\n", attrs.Name, err)
			continue
		}
		objects = append(objects, obj)
	}

	if len(objects) == 0 {
		return fmt.Errorf("no resources found in backup")
	}

	fmt.Printf("Found %d resources in backup. Validating...\n", len(objects))
	if err := validateResources(ctx, kubeClient, objects); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}

	fmt.Println("Sorting resources for restore...")
	sortResources(objects)

	fmt.Println("Initiating restore...")
	for _, obj := range objects {
		if err := applyObject(ctx, kubeClient, obj, options.dryRun); err != nil {
			fmt.Printf("Warning: failed to restore %s/%s (%s): %v\n", obj.GetNamespace(), obj.GetName(), obj.GetKind(), err)
		}
	}

	fmt.Println("Restore complete.")
	return nil
}

func loadObject(ctx context.Context, gcsClient *storage.Client, bucket, objectName string) (*unstructured.Unstructured, error) {
	rc, err := gcsClient.Bucket(bucket).Object(objectName).NewReader(ctx)
	if err != nil {
		return nil, err
	}
	defer rc.Close()

	data, err := io.ReadAll(rc)
	if err != nil {
		return nil, err
	}

	obj := &unstructured.Unstructured{}
	if err := yaml.Unmarshal(data, &obj.Object); err != nil {
		return nil, fmt.Errorf("unmarshaling YAML: %w", err)
	}

	return obj, nil
}

func validateResources(ctx context.Context, kubeClient *kubecli.Client, objects []*unstructured.Unstructured) error {
	// Check if GVKs exist in the target cluster
	resourceLists, err := kubeClient.DiscoveryClient.ServerPreferredResources()
	if err != nil {
		return fmt.Errorf("discovering server resources: %w", err)
	}

	supportedGVKs := make(map[string]bool)
	for _, list := range resourceLists {
		for _, r := range list.APIResources {
			gv, _ := schema.ParseGroupVersion(list.GroupVersion)
			gvk := gv.WithKind(r.Kind)
			supportedGVKs[gvk.String()] = true
		}
	}

	for _, obj := range objects {
		gvk := obj.GroupVersionKind()
		if !supportedGVKs[gvk.String()] {
			return fmt.Errorf("resource type %v not supported in target cluster (is the CRD installed?)", gvk)
		}
	}
	return nil
}

func sortResources(objects []*unstructured.Unstructured) {
	// Simple priority-based sort
	// Priority 1: Projects, Folders, Organizations
	// Priority 2: Networks
	// Priority 3: Everything else
	getPriority := func(kind string) int {
		switch kind {
		case "Project", "Folder", "Organization":
			return 1
		case "ComputeNetwork", "ComputeSubnetwork":
			return 2
		default:
			return 3
		}
	}

	sort.SliceStable(objects, func(i, j int) bool {
		pi := getPriority(objects[i].GetKind())
		pj := getPriority(objects[j].GetKind())
		if pi != pj {
			return pi < pj
		}
		return objects[i].GetKind() < objects[j].GetKind()
	})
}

func applyObject(ctx context.Context, kubeClient *kubecli.Client, obj *unstructured.Unstructured, dryRun bool) error {
	// Sanitization
	unstructured.RemoveNestedField(obj.Object, "metadata", "uid")
	unstructured.RemoveNestedField(obj.Object, "metadata", "resourceVersion")
	unstructured.RemoveNestedField(obj.Object, "metadata", "managedFields")

	// Acquisition
	annotations := obj.GetAnnotations()
	if annotations == nil {
		annotations = make(map[string]string)
	}
	annotations["cnrm.cloud.google.com/deletion-policy"] = "abandon"
	obj.SetAnnotations(annotations)

	if dryRun {
		fmt.Printf("Dry-run: would restore %s/%s (%s)\n", obj.GetNamespace(), obj.GetName(), obj.GetKind())
		return nil
	}

	// Apply to cluster
	if err := kubeClient.Patch(ctx, obj, client.Apply, client.FieldOwner("config-connector-backup"), client.ForceOwnership); err != nil {
		return fmt.Errorf("applying object: %w", err)
	}

	fmt.Printf("Restored %s/%s (%s)\n", obj.GetNamespace(), obj.GetName(), obj.GetKind())
	return nil
}
