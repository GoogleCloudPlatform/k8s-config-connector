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

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/commands/exportcsv"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/commands/generatemapper"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/commands/generatetypes"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/commands/updatetypes"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/options"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/scaffold"
	cctemplate "github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/template/controller"
	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type ControllerOptions struct {
	Kind      string
	ProtoName string
}

func buildAddCommand(baseOptions *options.GenerateOptions) *cobra.Command {
	opts := ControllerOptions{}
	addCmd := &cobra.Command{
		Use:   "add",
		Short: "add direct controller",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if opts.Kind == "" {
				return fmt.Errorf("--kind is required")
			}
			if opts.ProtoName == "" {
				return fmt.Errorf("--proto-resource is required")
			}

			if baseOptions.APIVersion == "" {
				return fmt.Errorf("--api-version is required")
			}
			_, err := schema.ParseGroupVersion(baseOptions.APIVersion)
			if err != nil {
				return fmt.Errorf("unable to parse --api-version: %w", err)
			}

			if baseOptions.ServiceName == "" {
				return fmt.Errorf("--service is required")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			gv, _ := schema.ParseGroupVersion(baseOptions.APIVersion)
			gcpTokens := strings.Split(baseOptions.ServiceName, ".")
			version := gcpTokens[len(gcpTokens)-1]
			if version[0] != 'v' {
				return fmt.Errorf("--service does not contain GCP version")
			}
			serviceName := strings.TrimSuffix(gv.Group, ".cnrm.cloud.google.com")
			cArgs := &cctemplate.ControllerArgs{
				KCCService:    serviceName,
				KCCVersion:    gv.Version,
				Kind:          opts.Kind,
				ProtoResource: opts.ProtoName,
				ProtoVersion:  version,
			}
			return scaffold.Scaffold(serviceName, opts.ProtoName, cArgs)
		},
	}
	addCmd.Flags().StringVarP(&opts.ProtoName, "proto-resource", "p", "", "the GCP resource proto name. It should match the name in the proto apis. i.e. For resource google.storage.v1.bucket, the `--proto-resource` should be `bucket`. If `--kind` is not given, the `--proto-resource` value will also be used as the kind name with a capital letter `Storage`.")
	addCmd.Flags().StringVarP(&opts.Kind, "kind", "k", "", "the KCC resource Kind. requires `--proto-resource`.")
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
	rootCmd.AddCommand(exportcsv.BuildCommand(&generateOptions))
	rootCmd.AddCommand(exportcsv.BuildPromptCommand(&generateOptions))

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
