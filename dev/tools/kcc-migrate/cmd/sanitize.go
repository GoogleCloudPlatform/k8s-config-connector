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
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/kcc-migrate/pkg/sanitizer"
	"github.com/spf13/cobra"
)

var (
	sanitizeInputDir        string
	sanitizeOutputDir       string
	sanitizeInjectAbandon   bool
	sanitizeInjectConflict  bool
	sanitizeInjectStateSpec bool
	sanitizeTargetNS        string
)

var sanitizeCmd = &cobra.Command{
	Use:   "sanitize",
	Short: "Sanitize exported KCC manifests and inject safe adoption annotations",
	Long: `The sanitize command strips dynamic status blocks, server-generated metadata
(uid, resourceVersion, generation, managedFields), and ACM/Config Sync annotations.
It injects 'cnrm.cloud.google.com/deletion-policy: abandon',
'cnrm.cloud.google.com/management-conflict-prevention-policy: none', and
'cnrm.cloud.google.com/state-into-spec: absent'
to guarantee zero GCP resource deletion, adoption lockouts, or immutability webhook rejections during migration.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if sanitizeInputDir == "" || sanitizeOutputDir == "" {
			return fmt.Errorf("both --input-dir and --output-dir are required")
		}

		opts := sanitizer.SanitizeOptions{
			InjectAbandon:       sanitizeInjectAbandon,
			InjectConflictNone:  sanitizeInjectConflict,
			InjectStateIntoSpec: sanitizeInjectStateSpec,
			TargetNamespace:     sanitizeTargetNS,
		}

		fmt.Printf("==> Starting sanitization pipeline...\n")
		fmt.Printf("    Input Directory:       %s\n", sanitizeInputDir)
		fmt.Printf("    Output Directory:      %s\n", sanitizeOutputDir)
		fmt.Printf("    Inject Deletion Policy: abandon    = %v\n", opts.InjectAbandon)
		fmt.Printf("    Inject Conflict Policy: none       = %v\n", opts.InjectConflictNone)
		fmt.Printf("    Inject State-Into-Spec: absent     = %v\n", opts.InjectStateIntoSpec)
		if opts.TargetNamespace != "" {
			fmt.Printf("    Target Namespace:      %s\n", opts.TargetNamespace)
		}

		totalFiles := 0
		totalResources := 0

		err := filepath.Walk(sanitizeInputDir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				return nil
			}
			if !strings.HasSuffix(path, ".yaml") && !strings.HasSuffix(path, ".yml") {
				return nil
			}

			relPath, err := filepath.Rel(sanitizeInputDir, path)
			if err != nil {
				return err
			}
			destPath := filepath.Join(sanitizeOutputDir, relPath)

			count, err := sanitizer.SanitizeFile(path, destPath, opts)
			if err != nil {
				return fmt.Errorf("error processing %s: %w", path, err)
			}

			if count > 0 {
				totalFiles++
				totalResources += count
				fmt.Printf("    [OK] %s -> %s (%d resources)\n", relPath, filepath.Base(destPath), count)
			}
			return nil
		})

		if err != nil {
			return err
		}

		fmt.Printf("\n==> Sanitization Complete!\n")
		fmt.Printf("    Files Processed:   %d\n", totalFiles)
		fmt.Printf("    Resources Cleaned: %d\n", totalResources)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(sanitizeCmd)
	sanitizeCmd.Flags().StringVarP(&sanitizeInputDir, "input-dir", "i", "", "Directory containing raw exported manifests")
	sanitizeCmd.Flags().StringVarP(&sanitizeOutputDir, "output-dir", "o", "", "Directory to write sanitized manifests")
	sanitizeCmd.Flags().BoolVar(&sanitizeInjectAbandon, "inject-abandon", true, "Inject 'cnrm.cloud.google.com/deletion-policy: abandon'")
	sanitizeCmd.Flags().BoolVar(&sanitizeInjectConflict, "inject-conflict-none", true, "Inject 'cnrm.cloud.google.com/management-conflict-prevention-policy: none'")
	sanitizeCmd.Flags().BoolVar(&sanitizeInjectStateSpec, "inject-state-into-spec", true, "Inject 'cnrm.cloud.google.com/state-into-spec: absent'")
	sanitizeCmd.Flags().StringVarP(&sanitizeTargetNS, "target-namespace", "n", "", "Optionally override namespace in all manifests")
}
