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

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/kcc-migrate/pkg/cluster"
	"github.com/spf13/cobra"
)

var (
	adoptKubeconfig string
	adoptContext    string
	adoptInputDir   string
	adoptDryRun     bool
)

var adoptCmd = &cobra.Command{
	Use:   "adopt",
	Short: "Adopt sanitized manifests into the target Standalone KCC cluster",
	Long: `Applies sanitized KCC manifests to the target cluster.
With 'deletion-policy: abandon' and 'management-conflict-prevention-policy: none'
embedded, KCC reconcilers will adopt existing live GCP resources without downtime or recreation.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if adoptInputDir == "" {
			return fmt.Errorf("--input-dir is required")
		}

		client := cluster.NewKubectlClient(adoptKubeconfig, adoptContext)

		if adoptDryRun {
			fmt.Printf("==> Executing SERVER DRY-RUN apply from: %s\n", adoptInputDir)
			out, err := client.ApplyDirectory(adoptInputDir, true)
			if err != nil {
				return fmt.Errorf("dry-run apply failed: %w", err)
			}
			fmt.Println(out)
			fmt.Println("==> Dry-run succeeded! Ready for live adoption.")
			return nil
		}

		fmt.Printf("==> Executing LIVE adoption apply from: %s\n", adoptInputDir)
		out, err := client.ApplyDirectory(adoptInputDir, false)
		if err != nil {
			return fmt.Errorf("live adoption failed: %w", err)
		}
		fmt.Println(out)

		fmt.Println("==> Adoption manifests submitted. Standalone KCC is reconciling existing cloud infrastructure.")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(adoptCmd)
	adoptCmd.Flags().StringVar(&adoptKubeconfig, "kubeconfig", "", "Path to kubeconfig file")
	adoptCmd.Flags().StringVar(&adoptContext, "context", "", "Kubernetes context to use")
	adoptCmd.Flags().StringVarP(&adoptInputDir, "input-dir", "i", "", "Directory containing sanitized manifests")
	adoptCmd.Flags().BoolVar(&adoptDryRun, "dry-run", false, "Perform server-side dry run only")
}
