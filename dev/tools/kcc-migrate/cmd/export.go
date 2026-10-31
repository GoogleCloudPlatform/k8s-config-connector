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

package cmd

import (
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/kcc-migrate/pkg/cluster"
	"github.com/spf13/cobra"
)

var (
	exportKubeconfig string
	exportContext    string
	exportNamespaces string
	exportOutputDir  string
)

var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export live KCC resources from a source Config Controller cluster",
	Long: `Discovers all installed KCC CRDs (*.cnrm.cloud.google.com) on the source cluster
and extracts all live custom resources in the specified namespaces.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if exportOutputDir == "" {
			return fmt.Errorf("--output-dir is required")
		}

		client := cluster.NewKubectlClient(exportKubeconfig, exportContext)

		fmt.Println("==> Querying source cluster for installed KCC CRDs...")
		crds, err := client.GetKCCCRDs()
		if err != nil {
			return fmt.Errorf("failed to discover KCC CRDs: %w", err)
		}
		fmt.Printf("    Discovered %d KCC CustomResourceDefinitions.\n", len(crds))

		var namespaces []string
		if exportNamespaces != "" {
			for _, ns := range strings.Split(exportNamespaces, ",") {
				if trimmed := strings.TrimSpace(ns); trimmed != "" {
					namespaces = append(namespaces, trimmed)
				}
			}
		} else {
			namespaces = []string{"default"}
		}

		fmt.Printf("==> Exporting live resources across namespaces: %v\n", namespaces)
		count, err := client.ExportResources(crds, namespaces, exportOutputDir)
		if err != nil {
			return fmt.Errorf("export error: %w", err)
		}

		fmt.Printf("==> Successfully exported %d resource files into %s\n", count, exportOutputDir)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(exportCmd)
	exportCmd.Flags().StringVar(&exportKubeconfig, "kubeconfig", "", "Path to kubeconfig file")
	exportCmd.Flags().StringVar(&exportContext, "context", "", "Kubernetes context to use")
	exportCmd.Flags().StringVarP(&exportNamespaces, "namespaces", "n", "", "Comma-separated list of namespaces to export")
	exportCmd.Flags().StringVarP(&exportOutputDir, "output-dir", "o", "", "Directory to write raw exported manifests")
}
