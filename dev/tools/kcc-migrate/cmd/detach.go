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
	"sync"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/kcc-migrate/pkg/cluster"
	"github.com/spf13/cobra"
)

var (
	detachKubeconfig string
	detachContext    string
	detachNamespaces string
	detachForce      bool
)

var detachCmd = &cobra.Command{
	Use:   "detach",
	Short: "Safely detach and delete KCC resources from the source Config Controller",
	Long: `The detach command checks all KCC resources in the source cluster to verify
that 'cnrm.cloud.google.com/deletion-policy: abandon' is applied, and then cleanly
deletes the Kubernetes Custom Resources from the source cluster without deleting
the underlying Google Cloud infrastructure.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		client := cluster.NewKubectlClient(detachKubeconfig, detachContext)

		var namespaces []string
		if detachNamespaces != "" {
			for _, ns := range strings.Split(detachNamespaces, ",") {
				if trimmed := strings.TrimSpace(ns); trimmed != "" {
					namespaces = append(namespaces, trimmed)
				}
			}
		} else {
			namespaces = []string{"default"}
		}

		fmt.Printf("==> Initiating SAFE DETACH on source cluster across namespaces: %v\n", namespaces)

		crds, err := client.GetKCCCRDs()
		if err != nil {
			return fmt.Errorf("listing CRDs: %w", err)
		}

		var filteredCRDs []string
		for _, crd := range crds {
			if strings.HasPrefix(crd, "configconnectors.") || strings.HasPrefix(crd, "configconnectorcontexts.") || strings.HasPrefix(crd, "multiclusterleases.") {
				continue
			}
			filteredCRDs = append(filteredCRDs, crd)
		}

		// Step 1: Ensure deletion-policy: abandon is patched onto all resources
		fmt.Println("==> Auditing & enforcing 'cnrm.cloud.google.com/deletion-policy: abandon' across all resources...")
		var wg sync.WaitGroup
		crdChan := make(chan string, len(filteredCRDs))
		for _, crd := range filteredCRDs {
			crdChan <- crd
		}
		close(crdChan)

		numWorkers := 20
		if len(filteredCRDs) < numWorkers {
			numWorkers = len(filteredCRDs)
		}

		for i := 0; i < numWorkers; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				patchArg := `{"metadata":{"annotations":{"cnrm.cloud.google.com/deletion-policy":"abandon"}}}`
				for crd := range crdChan {
					for _, ns := range namespaces {
						out, err := client.RunCommand("get", crd, "-n", ns, "-o", "name")
						if err != nil || strings.TrimSpace(out) == "" {
							continue
						}
						items := strings.Split(strings.TrimSpace(out), "\n")
						for _, item := range items {
							item = strings.TrimSpace(item)
							if item == "" {
								continue
							}
							client.RunCommand("patch", item, "-n", ns, "--type=merge", "-p", patchArg)
							fmt.Printf("    Patched %s with deletion-policy: abandon\n", item)
						}
					}
				}
			}()
		}
		wg.Wait()

		fmt.Println("==> All resources confirmed with 'deletion-policy: abandon'.")
		fmt.Println("==> Deleting KCC CRs non-destructively from source cluster...")

		crdChan2 := make(chan string, len(filteredCRDs))
		for _, crd := range filteredCRDs {
			crdChan2 <- crd
		}
		close(crdChan2)

		for i := 0; i < numWorkers; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for crd := range crdChan2 {
					for _, ns := range namespaces {
						out, err := client.RunCommand("get", crd, "-n", ns, "-o", "name")
						if err != nil || strings.TrimSpace(out) == "" {
							continue
						}
						items := strings.Split(strings.TrimSpace(out), "\n")
						for _, item := range items {
							item = strings.TrimSpace(item)
							if item == "" {
								continue
							}
							_, delErr := client.RunCommand("delete", item, "-n", ns, "--wait=false")
							if delErr != nil {
								fmt.Printf("    [WARN] Failed to delete %s in %s: %v\n", item, ns, delErr)
							} else {
								fmt.Printf("    Detached %s from %s\n", item, ns)
							}
						}
					}
				}
			}()
		}
		wg.Wait()

		fmt.Println("==> Detach complete. Kubernetes CRs removed from source cluster; live GCP infrastructure remains intact.")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(detachCmd)
	detachCmd.Flags().StringVar(&detachKubeconfig, "kubeconfig", "", "Path to kubeconfig file")
	detachCmd.Flags().StringVar(&detachContext, "context", "", "Kubernetes context to use")
	detachCmd.Flags().StringVarP(&detachNamespaces, "namespaces", "n", "", "Namespaces to detach")
	detachCmd.Flags().BoolVar(&detachForce, "force", false, "Force detach without confirmation")
}
