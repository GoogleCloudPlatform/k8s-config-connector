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
	"net/url"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/cmd/commonparams"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/cmd/export"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/cmd/export/parameters"

	"github.com/spf13/cobra"
)

var (
	exportParams = parameters.Parameters{}
	exportCmd    = &cobra.Command{
		Use:   "export [URL]",
		Short: "Export a GCP resource to Config Connector YAML",
		Long:  `Export a GCP resource to Config Connector YAML`,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return fmt.Errorf("expected a URL argument")
			}
			if len(args) > 1 {
				return fmt.Errorf("expected only a single URL argument")
			}
			uri := args[0]
			parsedURL, err := url.Parse(uri)
			if err != nil {
				return fmt.Errorf("error parsing URL argument '%v': %w", uri, err)
			}
			if parsedURL.Host == "" {
				return fmt.Errorf("invalid URL argument '%v': host portion is empty", uri)
			}
			if parsedURL.Path == "" || parsedURL.Path == "/" {
				return fmt.Errorf("invalid URL argument '%v': path portion is empty", uri)
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()

			fillRootFlagsOnExportParams(&exportParams)
			exportParams.URI = args[0]
			if err := parameters.Validate(&exportParams); err != nil {
				return err
			}
			rootCmd.SilenceUsage = true
			return export.Execute(ctx, &exportParams)
		},
	}
)

func init() {
	commonparams.AddOAuth2TokenParam(exportCmd, &exportParams.GCPAccessToken)
	commonparams.AddIAMFormatParam(exportCmd, &exportParams.IAMFormat)
	commonparams.AddFilterDeletedIAMMembersParam(exportCmd, &exportParams.FilterDeletedIAMMembers)
	commonparams.AddOutputParam(exportCmd, &exportParams.Output)
	commonparams.AddResourceFormatParam(exportCmd, &exportParams.ResourceFormat)
}

func fillRootFlagsOnExportParams(params *parameters.Parameters) {
	params.Verbose = verbose
}
