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

package updatetypes

import (
	"fmt"
	"os"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/options"

	"github.com/spf13/cobra"
)

type baseUpdateTypeOptions struct {
	*options.GenerateOptions

	ServiceName string

	ignoredFields    string // TODO: could be part of GenerateOptions
	apiDirectory     string
	apiGoPackagePath string
}

func (o *baseUpdateTypeOptions) InitDefaults() error {
	root, err := options.RepoRoot()
	if err != nil {
		return err
	}
	o.apiDirectory = root + "/apis/"
	o.apiGoPackagePath = "github.com/GoogleCloudPlatform/k8s-config-connector/apis/"
	return nil
}

func (o *baseUpdateTypeOptions) BindFlags(cmd *cobra.Command) {
	// TODO: Update this flag to accept a file path pointing to the ignored fields YAML file.
	cmd.Flags().StringVar(&o.ignoredFields, "ignored-fields", o.ignoredFields, "Comma-separated list of fields to ignore")
	cmd.Flags().StringVar(&o.apiDirectory, "api-dir", o.apiDirectory, "Base directory for APIs")
	cmd.Flags().StringVar(&o.apiGoPackagePath, "api-go-package-path", o.apiGoPackagePath, "API Go package path")
	cmd.Flags().StringVarP(&o.ServiceName, "service", "s", o.ServiceName, "the GCP service name")
}

func BuildCommand(baseOptions *options.GenerateOptions) *cobra.Command {
	opt := &baseUpdateTypeOptions{
		GenerateOptions: baseOptions,
	}
	if err := opt.InitDefaults(); err != nil {
		fmt.Fprintf(os.Stderr, "Error initializing defaults: %v\n", err)
		os.Exit(1)
	}

	cmd := &cobra.Command{
		Use:   "update-types",
		Short: "update KRM types for a proto service",
	}

	opt.BindFlags(cmd)

	// subcommands
	cmd.AddCommand(buildInsertCommand(opt))
	cmd.AddCommand(buildSyncCommand(opt))

	return cmd
}
