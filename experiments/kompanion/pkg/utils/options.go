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

package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/pflag"
)

const (
	// flag names.
	kubeconfigFlag = "kubeconfig"

	targetNamespacesFlag = "target-namespaces"
	ignoreNamespacesFlag = "exclude-namespaces"

	targetObjectsFlag = "target-objects"

	workerRoutinesFlag = "worker-routines"
)

type ClusterCrawlOptions struct {
	Kubeconfig string

	TargetNamespaces []string
	IgnoreNamespaces []string

	TargetObjects []string
	IgnoreObjects []string

	WorkerRoutines int
}

func (o *ClusterCrawlOptions) ClusterCrawlAddFlags(flags *pflag.FlagSet) {
	flags.StringVarP(&o.Kubeconfig, kubeconfigFlag, "", o.Kubeconfig, "path to the kubeconfig file.")

	flags.StringArrayVarP(&o.TargetNamespaces, targetNamespacesFlag, "", o.TargetNamespaces, "namespace prefix to target the export tool. Targets all if empty. Can be specified multiple times.")
	flags.StringArrayVarP(&o.IgnoreNamespaces, ignoreNamespacesFlag, "", o.IgnoreNamespaces, "namespace prefix to ignore. Excludes nothing if empty. Can be specified multiple times. Defaults to \"kube\".")

	flags.StringArrayVarP(&o.TargetObjects, targetObjectsFlag, "", o.TargetObjects, "object name prefix to target. Targets all if empty. Can be specified multiple times.")

	flags.IntVarP(&o.WorkerRoutines, workerRoutinesFlag, "", o.WorkerRoutines, "Configure the number of worker routines to export namespaces with. Defaults to 10. ")
}

func (opts *ClusterCrawlOptions) ValidateClusterCrawlFlags() error {
	if opts.WorkerRoutines <= 0 || opts.WorkerRoutines > 100 {
		return fmt.Errorf("invalid value %d for flag %s. Supported values are [1,100]", opts.WorkerRoutines, workerRoutinesFlag)
	}
	if opts.Kubeconfig != "" {
		if _, err := os.Stat(opts.Kubeconfig); err != nil {
			return fmt.Errorf("error reading kube config file %s, got %v", opts.Kubeconfig, err)
		}
	}

	return nil
}

func (o *ClusterCrawlOptions) ClusterCrawlPrint() {
	log.Printf("kubeconfig set to %q.\n", o.Kubeconfig)
	log.Printf("targetNamespaces set to %v.\n", o.TargetNamespaces)
	log.Printf("ignoreNamespaces set to %v.\n", o.IgnoreNamespaces)
	log.Printf("targetObjects set to %v.\n", o.TargetObjects)
	log.Printf("workerRountines set to %d.\n", o.WorkerRoutines)
}

func NewClusterCrawlOptions() ClusterCrawlOptions {
	o := ClusterCrawlOptions{
		Kubeconfig:       "",
		TargetNamespaces: []string{},
		IgnoreNamespaces: []string{"kube"},
		TargetObjects:    []string{},
		WorkerRoutines:   10,
	}
	return o
}
