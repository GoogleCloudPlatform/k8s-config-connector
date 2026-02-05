// Copyright 2026 Google LLC
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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/cmd/preview"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/cmd/preview/parameters"
	"github.com/spf13/cobra"
)

const (
	previewKubeconfigFlag = "kubeconfig"
	previewTimeoutFlag    = "timeout"

	previewReportNamePrefixFlag = "report-prefix"
)

const (
	previewExamples = `
	# preview Config Connector resources
	config-connector preview
	`
)

var (
	previewParams = parameters.Parameters{}
	previewCmd    = &cobra.Command{
		Use:     "preview",
		Short:   "preview Config Connector resources",
		Example: previewExamples,
		Hidden:  true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return preview.Execute(cmd.Context(), &previewParams)
		},
		Args: cobra.ExactArgs(0),
	}
)

func init() {
	// TODO: Add scope
	previewCmd.Flags().StringVarP(&previewParams.Kubeconfig, previewKubeconfigFlag, "", previewParams.Kubeconfig, "path to the kubeconfig file.")
	previewCmd.Flags().IntVarP(&previewParams.Timeout, previewTimeoutFlag, "", 15, "timeout in minutes. Default to 15 minutes.")
	previewCmd.Flags().StringVarP(&previewParams.ReportNamePrefix, previewReportNamePrefixFlag, "", "preview-report", "Prefix for the report name. The tool appends a timestamp to this in the format \"YYYYMMDD-HHMMSS.milliseconds\".")
	previewCmd.Flags().BoolVarP(&previewParams.FullReport, "full-report", "f", false, "Write a full report with all the captured objects. Disabled by default.")
	previewCmd.Flags().Float64VarP(&previewParams.GCPQPS, "gcpQPS", "q", 5.0, "Maximum qps for GCP API requests, per service. Default to 5.0. Set gcpQPS to 0 to disable rate limiting.")
	previewCmd.Flags().IntVarP(&previewParams.GCPBurst, "gcpBurst", "b", 5, "Maximum burst for GCP API requests, per service. Default to 5. Set gcpQPS to 0 to disable rate limiting.")
	previewCmd.Flags().StringVarP(&previewParams.Namespace, "namespace", "n", "", "Namespace to preview. If not specified, all namespaces will be previewed.")
	previewCmd.Flags().IntVarP(&previewParams.Verbose, "verbose", "v", 0, "Log verbosity level. Default to 0.")
	previewCmd.Flags().BoolVarP(&previewParams.InCluster, "in-cluster", "", false, "Run in GKE cluster. Default to false.")
}
