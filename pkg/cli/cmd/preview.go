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

func NewPreviewCmd() *cobra.Command {
	params := parameters.Parameters{}
	cmd := &cobra.Command{
		Use:     "preview",
		Short:   "preview Config Connector resources",
		Example: previewExamples,
		Hidden:  true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return preview.Execute(cmd.Context(), &params)
		},
		Args: cobra.ExactArgs(0),
	}

	cmd.Flags().StringVarP(&params.Kubeconfig, previewKubeconfigFlag, "", params.Kubeconfig, "path to the kubeconfig file.")
	cmd.Flags().IntVarP(&params.Timeout, previewTimeoutFlag, "", 15, "timeout in minutes. Default to 15 minutes.")
	cmd.Flags().StringVarP(&params.ReportNamePrefix, previewReportNamePrefixFlag, "", "preview-report", "Prefix for the report name. The tool appends a timestamp to this in the format \"YYYYMMDD-HHMMSS.milliseconds\".")
	cmd.Flags().BoolVarP(&params.FullReport, "full-report", "f", false, "Write a full report with all the captured objects. Disabled by default.")
	cmd.Flags().Float64VarP(&params.GCPQPS, "gcp-qps", "q", 5.0, "Maximum qps for GCP API requests, per service. Default to 5.0. Set gcp-qps to 0 to disable rate limiting.")
	cmd.Flags().IntVarP(&params.GCPBurst, "gcp-burst", "b", 5, "Maximum burst for GCP API requests, per service. Default to 5. Set gcp-qps to 0 to disable rate limiting.")
	cmd.Flags().StringVarP(&params.Namespace, "namespace", "n", "", "Namespace to preview. If not specified, all namespaces will be previewed.")
	cmd.Flags().IntVarP(&params.Verbose, "verbose", "v", 0, "Log verbosity level. Default to 0.")
	cmd.Flags().BoolVarP(&params.InCluster, "in-cluster", "", false, "Run in GKE cluster. Default to false.")

	return cmd
}
