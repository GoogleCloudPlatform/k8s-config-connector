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

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/kcc-migrate/pkg/linter"
	"github.com/spf13/cobra"
)

var (
	lintInputDir string
	lintStrict   bool
)

var lintCmd = &cobra.Command{
	Use:   "lint",
	Short: "Validate manifests against KCC safety guardrails before adoption",
	Long: `The lint command checks all YAML manifests in a directory to ensure:
1. Valid Kubernetes apiVersion, kind, and metadata.name.
2. No leaked runtime status or server-assigned metadata (uid, resourceVersion).
3. Presence of required safety annotations (deletion-policy: abandon, conflict-prevention: none).`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if lintInputDir == "" {
			return fmt.Errorf("--input-dir is required")
		}

		fmt.Printf("==> Linting manifests in: %s (strict-safety=%v)\n\n", lintInputDir, lintStrict)

		report, err := linter.LintDirectory(lintInputDir, lintStrict)
		if err != nil {
			return fmt.Errorf("lint error: %w", err)
		}

		warningCount := 0
		errorCount := 0

		for _, issue := range report.Issues {
			if issue.Severity == "ERROR" {
				errorCount++
				fmt.Fprintf(os.Stderr, "  [ERROR] %s (%s/%s): %s\n", issue.FilePath, issue.Kind, issue.Name, issue.Message)
			} else {
				warningCount++
				fmt.Printf("  [WARN]  %s (%s/%s): %s\n", issue.FilePath, issue.Kind, issue.Name, issue.Message)
			}
		}

		fmt.Printf("\n==> Lint Summary:\n")
		fmt.Printf("    Files Scanned:     %d\n", report.TotalFiles)
		fmt.Printf("    Resources Checked: %d\n", report.TotalResources)
		fmt.Printf("    Warnings:          %d\n", warningCount)
		fmt.Printf("    Errors:            %d\n", errorCount)

		if !report.Passed {
			return fmt.Errorf("linting failed with %d errors; aborting adoption until resolved", errorCount)
		}

		fmt.Println("    Status:            ALL CHECKS PASSED (Ready for Safe Adoption)")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(lintCmd)
	lintCmd.Flags().StringVarP(&lintInputDir, "input-dir", "i", "", "Directory containing manifests to lint")
	lintCmd.Flags().BoolVarP(&lintStrict, "strict", "s", true, "Fail on missing safety annotations")
}
