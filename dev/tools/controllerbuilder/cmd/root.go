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

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/scaffold"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/template"
	"github.com/spf13/cobra"
)

var (
	serviceName string
	// TODO: Resource and kind name should be the same. Validation the uppercase/lowercase.
	kind       string
	apiVersion string

	addCmd = &cobra.Command{
		Use:   "add",
		Short: "add direct controller",
		RunE: func(cmd *cobra.Command, args []string) error {
			// TODO(check kcc root)
			cArgs := &template.ControllerArgs{
				Service:     serviceName,
				Version:     apiVersion,
				Kind:        kind,
				KindToLower: strings.ToLower(kind),
			}
			path, err := scaffold.BuildControllerPath(serviceName, kind)
			if err != nil {
				return err
			}
			return scaffold.Scaffold(path, cArgs)
		},
	}
)

func init() {
	addCmd.PersistentFlags().StringVarP(&apiVersion, "version", "v", "v1alpha1", "the KRM API version. used to import the KRM API")
	addCmd.PersistentFlags().StringVarP(&serviceName, "service", "s", "", "the GCP service name")
	addCmd.PersistentFlags().StringVarP(&kind, "resourceInKind", "r", "", "the GCP resource name under the GCP service. should be in camel case ")
}

func Execute() {
	if err := addCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
