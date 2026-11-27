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

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "kcc-migrate",
	Short: "KCC Migration & Safe Adoption CLI for enterprise platforms",
	Long: `kcc-migrate is an enterprise-grade migration automation CLI designed for enterprise platforms
to safely migrate Google Cloud infrastructure-as-code from Google Cloud Config Controller (ACP)
to Self-Managed Standalone Config Connector (KCC) on GKE Standard.

It guarantees zero-downtime, non-destructive resource adoption by stripping server metadata,
injecting abandon deletion policies, and validating schemas before cutover.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func init() {
	// Global flags can be added here if needed
}
