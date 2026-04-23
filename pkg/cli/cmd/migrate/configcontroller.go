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

package migrate

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
	"sigs.k8s.io/yaml"
)

type ConfigControllerOptions struct {
	SourceKubeconfig string
	Output           string
}

func NewConfigControllerCmd() *cobra.Command {
	opts := &ConfigControllerOptions{}
	cmd := &cobra.Command{
		Use:   "config-controller",
		Short: "Migrate from a Config Controller cluster to a Manual Installation",
		Long: `Migrates Config Connector resources from a Config Controller instance to a Manual Installation on a separate management cluster.
This tool will:
1. Export Config Connector resources from the source cluster.
2. Strip cluster-specific operational statuses, generated names, and old annotations.
3. Add management-conflict-prevention-policy: none to take acquisition safely.
4. Output the transformed YAMLs to be imported into the target management cluster.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runConfigControllerMigrate(cmd.Context(), opts)
		},
	}
	cmd.Flags().StringVar(&opts.SourceKubeconfig, "source-kubeconfig", "", "Path to the kubeconfig for the source Config Controller cluster")
	cmd.Flags().StringVarP(&opts.Output, "output", "o", "-", "Output directory to save the YAML files, or '-' for stdout")

	// Mark source-kubeconfig as required
	cmd.MarkFlagRequired("source-kubeconfig")
	return cmd
}

func runConfigControllerMigrate(ctx context.Context, opts *ConfigControllerOptions) error {
	sourceConfig, err := clientcmd.BuildConfigFromFlags("", opts.SourceKubeconfig)
	if err != nil {
		return fmt.Errorf("error building source kubeconfig: %w", err)
	}

	discoveryClient, err := discovery.NewDiscoveryClientForConfig(sourceConfig)
	if err != nil {
		return fmt.Errorf("error creating discovery client: %w", err)
	}

	sourceDynamic, err := dynamic.NewForConfig(sourceConfig)
	if err != nil {
		return fmt.Errorf("error creating source dynamic client: %w", err)
	}

	// 1. Discover all KCC CRDs
	serverResources, err := discoveryClient.ServerPreferredResources()
	if err != nil {
		// Can happen if some APIs are failing, but we still get the successful ones
		fmt.Fprintf(os.Stderr, "Warning: error during discovery: %v\n", err)
	}

	var gvrs []schema.GroupVersionResource
	for _, group := range serverResources {
		gv, err := schema.ParseGroupVersion(group.GroupVersion)
		if err != nil {
			continue
		}
		// Filter for Config Connector groups
		if !strings.HasSuffix(gv.Group, ".cnrm.cloud.google.com") {
			continue
		}
		for _, resource := range group.APIResources {
			if strings.Contains(resource.Name, "/") {
				// Skip subresources
				continue
			}
			gvrs = append(gvrs, schema.GroupVersionResource{
				Group:    gv.Group,
				Version:  gv.Version,
				Resource: resource.Name,
			})
		}
	}

	var outWriter io.Writer = os.Stdout
	if opts.Output != "" && opts.Output != "-" {
		err := os.MkdirAll(opts.Output, 0755)
		if err != nil {
			return fmt.Errorf("failed to create output directory: %w", err)
		}
	}

	for _, gvr := range gvrs {
		list, err := sourceDynamic.Resource(gvr).Namespace(metav1.NamespaceAll).List(ctx, metav1.ListOptions{})
		if err != nil {
			continue // Might be permission issue or resource not existing
		}

		for _, item := range list.Items {
			// Strip cluster-specific fields
			stripUnwantedFields(&item)

			// Output
			yamlBytes, err := yaml.Marshal(item.Object)
			if err != nil {
				return fmt.Errorf("failed to marshal object %s/%s: %w", item.GetNamespace(), item.GetName(), err)
			}

			if opts.Output == "-" {
				fmt.Fprintln(outWriter, "---")
				fmt.Fprintln(outWriter, string(yamlBytes))
			} else {
				// Save to file
				filename := fmt.Sprintf("%s_%s_%s.yaml", item.GetKind(), item.GetNamespace(), item.GetName())
				if item.GetNamespace() == "" {
					filename = fmt.Sprintf("%s_%s.yaml", item.GetKind(), item.GetName())
				}
				filePath := filepath.Join(opts.Output, filename)
				err := os.WriteFile(filePath, yamlBytes, 0644)
				if err != nil {
					return fmt.Errorf("failed to write file %s: %w", filePath, err)
				}
			}
		}
	}

	return nil
}

func stripUnwantedFields(u *unstructured.Unstructured) {
	unstructured.RemoveNestedField(u.Object, "status")
	unstructured.RemoveNestedField(u.Object, "metadata", "creationTimestamp")
	unstructured.RemoveNestedField(u.Object, "metadata", "generation")
	unstructured.RemoveNestedField(u.Object, "metadata", "resourceVersion")
	unstructured.RemoveNestedField(u.Object, "metadata", "uid")
	unstructured.RemoveNestedField(u.Object, "metadata", "managedFields")
	unstructured.RemoveNestedField(u.Object, "metadata", "generateName")
	unstructured.RemoveNestedField(u.Object, "metadata", "ownerReferences")
	unstructured.RemoveNestedField(u.Object, "metadata", "finalizers")

	annotations := u.GetAnnotations()
	if annotations == nil {
		annotations = make(map[string]string)
	}
	delete(annotations, "kubectl.kubernetes.io/last-applied-configuration")

	// Set take acquisition mode
	annotations["cnrm.cloud.google.com/management-conflict-prevention-policy"] = "none"
	u.SetAnnotations(annotations)
}
