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
	"fmt"

	"github.com/spf13/pflag"
)

const (
	kubeconfigFlag = "kubeconfig"

	targetNamespacesFlag = "target-namespaces"
	ignoreNamespacesFlag = "exclude-namespaces"

	includeResourcesFlag = "include-resources"
	ignoreResourcesFlag  = "ignore-resources"

	workerRoutinesFlag = "worker-routines"
	verboseFlag        = "verbose"
)

var abandonOnDeleteCheckKind = map[string]bool{
	"AlloyDBCluster":    true,
	"AlloyDBInstance":   true,
	"BigQueryDataset":   true,
	"BigQueryTable":     true,
	"BigtableInstance":  true,
	"ContainerCluster":  true,
	"ContainerNodePool": true,
	"KMSKeyRing":        true,
	"KMSCryptoKey":      true,
	"RedisCluster":      true,
	"SpannerDatabase":   true,
	"SpannerInstance":   true,
	"SQLDatabase":       true,
	"SQLInstance":       true,
	"StorageBucket":     true,
}

type Options struct {
	Kubeconfig string

	TargetNamespaces []string
	IgnoreNamespaces []string

	IncludeResources []string
	IgnoreResources  []string

	WorkerRoutines int

	Verbose int
}

func NewOptions() *Options {
	return &Options{
		IgnoreNamespaces: []string{"kube"},
		WorkerRoutines:   10,
	}
}

func (o *Options) AddFlags(flags *pflag.FlagSet) {
	flags.StringVarP(&o.Kubeconfig, kubeconfigFlag, "", o.Kubeconfig, "path to the kubeconfig file.")

	flags.StringArrayVarP(&o.TargetNamespaces, targetNamespacesFlag, "", []string{}, "namespace prefix to target the lint tool. Targets all if empty. Can be specified multiple times.")
	flags.StringArrayVarP(&o.IgnoreNamespaces, ignoreNamespacesFlag, "", o.IgnoreNamespaces, "namespace prefix to ignore. Excludes nothing if empty. Can be specified multiple times. Defaults to \"kube\".")

	flags.StringSliceVarP(&o.IncludeResources, includeResourcesFlag, "", []string{}, "resource kind to include in the lint tool. Can be specified multiple times.")
	flags.StringSliceVarP(&o.IgnoreResources, ignoreResourcesFlag, "", []string{}, "resource kind to ignore in the lint tool. Can be specified multiple times.")
	flags.IntVarP(&o.WorkerRoutines, workerRoutinesFlag, "", o.WorkerRoutines, "Configure the number of worker routines to lint namespaces with. Defaults to 10. ")
	flags.IntVarP(&o.Verbose, verboseFlag, "v", 0, "Log verbosity level. Default to 0.")
}

func (o *Options) Validate() (map[string]bool, error) {
	kindsToLintMap := make(map[string]bool)
	for k, v := range abandonOnDeleteCheckKind {
		kindsToLintMap[k] = v
	}

	for _, kind := range o.IncludeResources {
		kindsToLintMap[kind] = true
	}

	for _, kind := range o.IgnoreResources {
		delete(kindsToLintMap, kind)
	}

	if o.WorkerRoutines <= 0 || o.WorkerRoutines > 100 {
		return nil, fmt.Errorf("invalid value %d for flag %s. Supported values are [1,100]", o.WorkerRoutines, workerRoutinesFlag)
	}
	return kindsToLintMap, nil
}
