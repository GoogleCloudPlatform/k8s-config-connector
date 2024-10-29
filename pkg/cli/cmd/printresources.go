// Copyright 2022 Google LLC
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

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/cmd/printresources"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/cmd/printresources/parameters"

	"github.com/spf13/cobra"
)

const (
	printResourcesCommandName = "print-resources"
)

var (
	printResourcesParams         = parameters.Parameters{}
	printResourcesCmdDescription = fmt.Sprintf("Print the list of resources for %v", commandName)
	printResourcesCmd            = &cobra.Command{
		Use:   printResourcesCommandName,
		Short: printResourcesCmdDescription,
		Long:  printResourcesCmdDescription,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := parameters.Validate(&printResourcesParams); err != nil {
				return err
			}
			rootCmd.SilenceUsage = true
			return printresources.Execute(&printResourcesParams, os.Stdout)
		},
		Args: cobra.NoArgs,
	}
	outputFormat string
)

func init() {
	inputUsage := fmt.Sprintf("specify the format of the output, options are '%v', '%v', or '%v' (default: '%v')",
		parameters.TableOutputFormat, parameters.YAMLOutputFormat, parameters.JSONOutputFormat, parameters.DefaultOutputFormat)
	printResourcesCmd.Flags().StringVarP(&printResourcesParams.OutputFormat, parameters.OutputFormatParam, "o",
		parameters.DefaultOutputFormat, inputUsage)
}
