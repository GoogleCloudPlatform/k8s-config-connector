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
	"os"

	"github.com/spf13/cobra"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/cmd/apply"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/cmd/apply/parameters"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/cmd/commonparams"
)

const (
	applyCommandName = "apply"
)

var (
	applyParams = parameters.Parameters{}
	applyCmd    = &cobra.Command{
		Use:    applyCommandName,
		Hidden: true,
		Short:  "Apply a KRM resource configuration file to Google Cloud Platform backend",
		Long:   `Apply a KRM resource configuration file to Google Cloud Platform backend`,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			if err := parameters.Validate(&applyParams, os.Stdin); err != nil {
				return err
			}
			rootCmd.SilenceUsage = true
			return apply.Execute(ctx, &applyParams, os.Stdout)
		},
		Args: cobra.NoArgs,
	}
)

func init() {
	commonparams.AddOAuth2TokenParam(applyCmd, &applyParams.OAuth2Token)
	inputUsage := "the input file path containing the KRM file to be applied."
	applyCmd.Flags().StringVarP(&applyParams.Input, parameters.InputParam, "i", "", inputUsage)
}
