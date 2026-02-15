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

package options

import (
	"github.com/spf13/pflag"
)

const (
	kubeconfigFlag       = "kubeconfig"
	timeoutFlag          = "timeout"
	reportNamePrefixFlag = "report-prefix"
	fullReportFlag       = "full-report"
	gcpQPSFlag           = "gcp-qps"
	gcpBurstFlag         = "gcp-burst"
	namespaceFlag        = "namespace"
	verboseFlag          = "verbose"
	inClusterFlag        = "in-cluster"
)

type Options struct {
	Kubeconfig       string
	Timeout          int
	ReportNamePrefix string
	FullReport       bool
	GCPQPS           float64
	GCPBurst         int
	Namespace        string
	Verbose          int
	InCluster        bool
}

func NewOptions() *Options {
	return &Options{
		Timeout:          15,
		ReportNamePrefix: "preview-report",
		GCPQPS:           5.0,
		GCPBurst:         5,
	}
}

func (o *Options) AddFlags(flags *pflag.FlagSet) {
	flags.StringVarP(&o.Kubeconfig, kubeconfigFlag, "", o.Kubeconfig, "path to the kubeconfig file.")
	flags.IntVarP(&o.Timeout, timeoutFlag, "", o.Timeout, "timeout in minutes. Default to 15 minutes.")
	flags.StringVarP(&o.ReportNamePrefix, reportNamePrefixFlag, "", o.ReportNamePrefix, "Prefix for the report name. The tool appends a timestamp to this in the format \"YYYYMMDD-HHMMSS.milliseconds\".")
	flags.BoolVarP(&o.FullReport, fullReportFlag, "f", o.FullReport, "Write a full report with all the captured objects. Disabled by default.")
	flags.Float64VarP(&o.GCPQPS, gcpQPSFlag, "q", o.GCPQPS, "Maximum qps for GCP API requests, per service. Default to 5.0. Set gcp-qps to 0 to disable rate limiting.")
	flags.IntVarP(&o.GCPBurst, gcpBurstFlag, "b", o.GCPBurst, "Maximum burst for GCP API requests, per service. Default to 5. Set gcp-qps to 0 to disable rate limiting.")
	flags.StringVarP(&o.Namespace, namespaceFlag, "n", o.Namespace, "Namespace to preview. If not specified, all namespaces will be previewed.")
	flags.IntVarP(&o.Verbose, verboseFlag, "v", o.Verbose, "Log verbosity level. Default to 0.")
	flags.BoolVarP(&o.InCluster, inClusterFlag, "", o.InCluster, "Run in GKE cluster. Default to false.")
}
