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

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/commands/apply"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/commands/detectnewfields"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/commands/exportcsv"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/commands/generatecontroller"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/commands/generatedirectreconciler"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/commands/generatefuzzer"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/commands/generatemapper"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/commands/generatetypes"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/commands/updatetypes"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/options"
	"github.com/spf13/cobra"
)

func Execute() {
	var generateOptions options.GenerateOptions
	if err := generateOptions.InitDefaults(); err != nil {
		fmt.Fprintf(os.Stderr, "Error initializing defaults: %v\n", err)
		os.Exit(1)
	}
	rootCmd := &cobra.Command{}
	generateOptions.BindPersistentFlags(rootCmd)

	rootCmd.AddCommand(generatedirectreconciler.BuildCommand(&generateOptions))
	rootCmd.AddCommand(generatecontroller.BuildCommand(&generateOptions))
	rootCmd.AddCommand(generatetypes.BuildCommand(&generateOptions))
	rootCmd.AddCommand(generatetypes.BuildV2Command(&generateOptions))
	rootCmd.AddCommand(generatemapper.BuildCommand(&generateOptions))
	rootCmd.AddCommand(updatetypes.BuildCommand(&generateOptions))
	rootCmd.AddCommand(exportcsv.BuildCommand(&generateOptions))
	rootCmd.AddCommand(exportcsv.BuildPromptCommand(&generateOptions))
	rootCmd.AddCommand(apply.BuildCommand(&generateOptions))
	rootCmd.AddCommand(detectnewfields.BuildCommand(&generateOptions))
	rootCmd.AddCommand(generatefuzzer.BuildCommand(&generateOptions))

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
