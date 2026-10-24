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
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/kcc-migrate/pkg/cluster"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/kcc-migrate/pkg/sorter"
	"github.com/spf13/cobra"
)

var (
	adoptKubeconfig string
	adoptContext    string
	adoptInputDir   string
	adoptDryRun     bool
	adoptSort       bool
)

var tierDescriptions = map[int]string{
	0:  "Tier 0: Root Foundations & Identities (KMSKeyRing, ComputeNetwork, IAMServiceAccount, StorageBucket)",
	1:  "Tier 1: Sub-foundations & Subnets (KMSCryptoKey, ComputeSubnetwork, SecretManagerSecret, PubSubTopic)",
	2:  "Tier 2: Core Stateful Engines & Firewalls (SQLInstance, SpannerInstance, RedisInstance, ComputeFirewall)",
	3:  "Tier 3: Stateful Children & Databases (SQLDatabase, SQLUser, DNSRecordSet, PubSubSubscription)",
	4:  "Tier 4: Routing & Endpoints (ComputeBackendService, ComputeForwardingRule)",
	5:  "Tier 5: IAM Policy Members & Bindings (IAMPolicyMember, IAMPolicy, IAMPartialPolicy)",
	10: "Tier 10: Standard Custom Resources",
}

var adoptCmd = &cobra.Command{
	Use:   "adopt",
	Short: "Adopt sanitized manifests into the target Standalone KCC cluster",
	Long: `Applies sanitized KCC manifests to the target cluster.
With 'deletion-policy: abandon' and 'management-conflict-prevention-policy: none'
embedded, KCC reconcilers will adopt existing live GCP resources without downtime or recreation.
Topological DAG sorting is enabled by default to eliminate reference resolution errors.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if adoptInputDir == "" {
			return fmt.Errorf("--input-dir is required")
		}

		client := cluster.NewKubectlClient(adoptKubeconfig, adoptContext)

		if adoptSort {
			tmpDir, err := os.MkdirTemp("", "kcc-adopt-sorted-*")
			if err != nil {
				return fmt.Errorf("creating temp dir for sorted manifests: %w", err)
			}
			defer os.RemoveAll(tmpDir)

			count, err := sorter.SortDirectory(adoptInputDir, tmpDir)
			if err != nil {
				return fmt.Errorf("topological sorting failed: %w", err)
			}
			if count == 0 {
				fmt.Printf("==> No KCC resources found in %s to adopt.\n", adoptInputDir)
				return nil
			}
			fmt.Printf("==> Topologically sorted %d resources across dependency tiers\n", count)

			tierFiles, err := os.ReadDir(tmpDir)
			if err != nil {
				return fmt.Errorf("reading sorted tier directory: %w", err)
			}

			// Sort filenames to ensure tier0 -> tier1 -> ... -> tier5
			var fileNames []string
			for _, f := range tierFiles {
				if strings.HasSuffix(f.Name(), ".yaml") {
					fileNames = append(fileNames, f.Name())
				}
			}
			sort.Strings(fileNames)

			for _, fn := range fileNames {
				var tierNum int
				fmt.Sscanf(fn, "tier%d_ordered.yaml", &tierNum)
				desc, ok := tierDescriptions[tierNum]
				if !ok {
					desc = fmt.Sprintf("Tier %d: Dependent Resources", tierNum)
				}

				tierPath := filepath.Join(tmpDir, fn)
				if adoptDryRun {
					fmt.Printf("==> [DRY-RUN] Applying %s...\n", desc)
					out, err := client.ApplyFile(tierPath, true)
					if err != nil {
						return fmt.Errorf("dry-run failed on %s: %w", fn, err)
					}
					fmt.Println(out)
				} else {
					fmt.Printf("==> [LIVE] Applying %s...\n", desc)
					out, err := client.ApplyFile(tierPath, false)
					if err != nil {
						return fmt.Errorf("live adoption failed on %s: %w", fn, err)
					}
					fmt.Println(out)
				}
			}

			if adoptDryRun {
				fmt.Println("==> Dry-run succeeded across all tiers! Ready for live adoption.")
			} else {
				fmt.Println("==> All tiers applied in topological order. Standalone KCC is reconciling existing cloud infrastructure.")
			}
			return nil
		}

		if adoptDryRun {
			fmt.Printf("==> Executing SERVER DRY-RUN apply (unordered) from: %s\n", adoptInputDir)
			out, err := client.ApplyDirectory(adoptInputDir, true)
			if err != nil {
				return fmt.Errorf("dry-run apply failed: %w", err)
			}
			fmt.Println(out)
			fmt.Println("==> Dry-run succeeded! Safe to execute live adoption.")
			return nil
		}

		fmt.Printf("==> Executing LIVE adoption apply (unordered) from: %s\n", adoptInputDir)
		out, err := client.ApplyDirectory(adoptInputDir, false)
		if err != nil {
			return fmt.Errorf("live adoption failed: %w", err)
		}
		fmt.Println(out)
		fmt.Println("==> Adoption manifests applied! Standalone KCC is reconciling existing cloud infrastructure.")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(adoptCmd)
	adoptCmd.Flags().StringVar(&adoptKubeconfig, "kubeconfig", "", "Path to kubeconfig file")
	adoptCmd.Flags().StringVar(&adoptContext, "context", "", "Kubernetes context to use")
	adoptCmd.Flags().StringVarP(&adoptInputDir, "input-dir", "i", "", "Directory containing sanitized manifests")
	adoptCmd.Flags().BoolVar(&adoptDryRun, "dry-run", false, "Perform server-side dry run only")
	adoptCmd.Flags().BoolVar(&adoptSort, "sort", true, "Topologically sort manifests across dependency tiers before applying")
}
