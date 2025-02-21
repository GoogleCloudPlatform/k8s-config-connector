// Copyright 2025 Google LLC
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

package main

import (
	"fmt"
	"os"

	"github.com/GoogleCloudPlatform/k8s-config-connector/experiments/conductor/cmd/runner"
	// "github.com/GoogleCloudPlatform/k8s-config-connector/experiments/kompanion/cmd/summary"
	// "github.com/GoogleCloudPlatform/k8s-config-connector/experiments/kompanion/pkg/version"
	"github.com/spf13/cobra"
)

func BuildRootCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "conductor subcommand",
		Short: "conductor is an alpha cli tool for managing ML tools.",
	}

	rootCmd.AddCommand(runner.BuildRunnerCmd())
	// rootCmd.AddCommand(summary.BuildSummaryCmd())

	// rootCmd.Version = version.GetVersion()
	// rootCmd.CompletionOptions.DisableDefaultCmd = true

	return rootCmd
}

func main() {
	rootCmd := BuildRootCommand()
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
