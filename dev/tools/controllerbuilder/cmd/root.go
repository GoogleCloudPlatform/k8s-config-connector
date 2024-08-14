// Copyright 2024 Google LLC
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
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/commands/generatemapper"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/commands/generatetypes"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/commands/updatetypes"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/options"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/scaffold"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/template"
	"github.com/spf13/cobra"
)

func buildAddCommand(baseOptions *options.GenerateOptions) *cobra.Command {
	// TODO: Resource and kind name should be the same. Validation the uppercase/lowercase.
	kind := ""

	addCmd := &cobra.Command{
		Use:   "add",
		Short: "add direct controller",
		RunE: func(cmd *cobra.Command, args []string) error {
			// TODO(check kcc root)
			cArgs := &template.ControllerArgs{
				Service:     baseOptions.ServiceName,
				Version:     baseOptions.APIVersion,
				Kind:        kind,
				KindToLower: strings.ToLower(kind),
			}
			return scaffold.Scaffold(baseOptions.ServiceName, kind, cArgs)
		},
	}
	addCmd.PersistentFlags().StringVarP(&kind, "resourceInKind", "r", "", "the GCP resource name under the GCP service. should be in camel case ")

	return addCmd
}

func Execute() {
	var generateOptions options.GenerateOptions
	generateOptions.InitDefaults()

	rootCmd := &cobra.Command{}
	generateOptions.BindPersistentFlags(rootCmd)

	rootCmd.AddCommand(buildAddCommand(&generateOptions))
	rootCmd.AddCommand(generatetypes.BuildCommand(&generateOptions))
	rootCmd.AddCommand(generatemapper.BuildCommand(&generateOptions))
	rootCmd.AddCommand(updatetypes.BuildCommand(&generateOptions))

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
