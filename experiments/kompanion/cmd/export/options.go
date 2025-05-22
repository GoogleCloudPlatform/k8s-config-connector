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

package export

import (
	"log"

	"github.com/GoogleCloudPlatform/k8s-config-connector/experiments/kompanion/pkg/utils"
	"github.com/spf13/pflag"
)

const (
	reportNamePrefixFlag = "report-prefix"
)

type ExportOptions struct {
	utils.ClusterCrawlOptions

	ReportNamePrefix string
}

func (opts *ExportOptions) AddFlags(flags *pflag.FlagSet) {
	opts.ClusterCrawlAddFlags(flags)

	flags.StringVarP(&opts.ReportNamePrefix, reportNamePrefixFlag, "", opts.ReportNamePrefix, "Prefix for the report name. The tool appends a timestamp to this in the format \"YYYYMMDD-HHMMSS.milliseconds\".")
}

func (opts *ExportOptions) validateFlags() error {
	if err := opts.ValidateClusterCrawlFlags(); err != nil {
		return err
	}

	return nil
}

func (opts *ExportOptions) Print() {
	opts.ClusterCrawlPrint()
	log.Printf("reportNamePrefix set to %q.\n", opts.ReportNamePrefix)
}

func NewExportOptions() *ExportOptions {
	opts := ExportOptions{
		ReportNamePrefix: "report",
	}
	opts.ClusterCrawlOptions = utils.NewClusterCrawlOptions()
	return &opts
}
