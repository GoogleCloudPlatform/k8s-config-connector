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

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/cmd/bulkexport"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/cmd/bulkexport/parameters"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/cmd/commonparams"

	"github.com/spf13/cobra"
)

const (
	onErrorDefault        = parameters.HaltOnErrorOption
	bulkExportCommandName = "bulk-export"
)

var (
	bulkExportParams = parameters.Parameters{}
	bulkExportCmd    = &cobra.Command{
		Use:   bulkExportCommandName,
		Short: "Perform a bulk export of GCP resources to Config Connector YAML using asset inventory",
		Long:  `Perform a bulk export of GCP resources to Config Connector YAML using asset inventory`,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()

			fillRootFlagsOnBulkExportParams(&bulkExportParams)
			if err := parameters.Validate(&bulkExportParams, os.Stdin); err != nil {
				return err
			}
			rootCmd.SilenceUsage = true
			return bulkexport.Execute(ctx, &bulkExportParams)
		},
		Args: cobra.NoArgs,
	}
)

func init() {
	commonparams.AddOAuth2TokenParam(bulkExportCmd, &bulkExportParams.OAuth2Token)
	commonparams.AddIAMFormatParam(bulkExportCmd, &bulkExportParams.IAMFormat)
	commonparams.AddFilterDeletedIAMMembersParam(bulkExportCmd, &bulkExportParams.FilterDeletedIAMMembers)
	commonparams.AddOutputParam(bulkExportCmd, &bulkExportParams.Output)
	commonparams.AddResourceFormatParam(bulkExportCmd, &bulkExportParams.ResourceFormat)
	inputUsage := fmt.Sprintf("an optional input file path containing an asset inventory export, cannot be used with piped input or '%v'", parameters.StorageKeyParam)
	bulkExportCmd.Flags().StringVarP(&bulkExportParams.Input, parameters.InputParam, "i", "", inputUsage)
	onErrorUsage := fmt.Sprintf("control the behavior when a recoverable error occurs, options are '%v', '%v', or '%v'", parameters.ContinueOnErrorOption, parameters.HaltOnErrorOption, parameters.IgnoreOnErrorOption)
	bulkExportCmd.Flags().StringVar(&bulkExportParams.OnError, parameters.OnErrorParam, onErrorDefault, onErrorUsage)
	storageKeyUsage := "an optional cloud storage key where an asset inventory export will be stored, example: 'gs://your-bucket-name/your/prefix/path'"
	bulkExportCmd.Flags().StringVarP(&bulkExportParams.StorageKey, parameters.StorageKeyParam, "s", "", storageKeyUsage)
	projectUsage := fmt.Sprintf("an optional project id for which a cloud asset inventory will be exported to a temporary bucket; use the '%v' parameter to avoid the creation of a temporary bucket", parameters.StorageKeyParam)
	bulkExportCmd.Flags().StringVar(&bulkExportParams.ProjectID, parameters.ProjectIDParam, "", projectUsage)
	folderUsage := fmt.Sprintf("an optional folder id for which a cloud asset inventory will be exported to a temporary bucket; use the '%v' parameter to avoid the creation of a temporary bucket", parameters.StorageKeyParam)
	bulkExportCmd.Flags().IntVar(&bulkExportParams.FolderID, parameters.FolderIDParam, 0, folderUsage)
	organizationUsage := fmt.Sprintf("an optional organization id for which a cloud asset inventory will be exported to a temporary bucket; use the '%v' parameter to avoid the creation of a temporary bucket", parameters.StorageKeyParam)
	bulkExportCmd.Flags().IntVar(&bulkExportParams.OrganizationID, parameters.OrganizationIDParam, 0, organizationUsage)
}

func fillRootFlagsOnBulkExportParams(params *parameters.Parameters) {
	params.Verbose = verbose
}
