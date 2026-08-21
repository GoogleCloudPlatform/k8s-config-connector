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

package summary

import (
	"fmt"
	"log"

	"github.com/spf13/pflag"
)

const (
	// flag names.
	kubeconfigFlag       = "kubeconfig"
	reportNamePrefixFlag = "report-prefix"

	targetNamespacesFlag = "target-namespaces"
	ignoreNamespacesFlag = "exclude-namespaces"

	targetObjectsFlag = "target-objects"

	workerRoutinesFlag = "worker-routines"
)

type SummaryOptions struct {
	kubeconfig       string
	reportNamePrefix string

	targetNamespaces []string
	ignoreNamespaces []string

	targetObjects []string

	workerRountines int
}

func (o *SummaryOptions) AddFlags(flags *pflag.FlagSet) {
	flags.StringVarP(&o.kubeconfig, kubeconfigFlag, "", o.kubeconfig, "path to the kubeconfig file.")
	flags.StringVarP(&o.reportNamePrefix, reportNamePrefixFlag, "", o.reportNamePrefix, "Prefix for the report name. The tool appends a timestamp to this in the format \"YYYYMMDD-HHMMSS.milliseconds\".")

	flags.StringArrayVarP(&o.targetNamespaces, targetNamespacesFlag, "", o.targetNamespaces, "namespace prefix to target the export tool. Targets all if empty. Can be specified multiple times.")
	flags.StringArrayVarP(&o.ignoreNamespaces, ignoreNamespacesFlag, "", o.ignoreNamespaces, "namespace prefix to ignore. Excludes nothing if empty. Can be specified multiple times. Defaults to \"kube\".")

	flags.StringArrayVarP(&o.targetObjects, targetObjectsFlag, "", o.targetObjects, "object name prefix to target. Targets all if empty. Can be specified multiple times.")

	flags.IntVarP(&o.workerRountines, workerRoutinesFlag, "", o.workerRountines, "Configure the number of worker routines to export namespaces with. Defaults to 10. ")
}

func (opts *SummaryOptions) validateFlags() error {
	if opts.workerRountines <= 0 || opts.workerRountines > 100 {
		return fmt.Errorf("invalid value %d for flag %s. Supported values are [1,100]", opts.workerRountines, workerRoutinesFlag)
	}

	return nil
}

func (o *SummaryOptions) Print() {
	log.Printf("kubeconfig set to %q.\n", o.kubeconfig)
	log.Printf("reportNamePrefix set to %q.\n", o.reportNamePrefix)
	log.Printf("targetNamespaces set to %v.\n", o.targetNamespaces)
	log.Printf("ignoreNamespaces set to %v.\n", o.ignoreNamespaces)
	log.Printf("targetObjects set to %v.\n", o.targetObjects)
	log.Printf("workerRountines set to %d.\n", o.workerRountines)
}

func NewSummaryOptions() *SummaryOptions {
	o := SummaryOptions{
		kubeconfig:       "",
		reportNamePrefix: "report",
		targetNamespaces: []string{},
		ignoreNamespaces: []string{"kube"},
		targetObjects:    []string{},
		workerRountines:  10,
	}
	return &o
}
