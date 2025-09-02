// Copyright 2025 Google LLC
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

package lint

import (
	"context"
	"fmt"
	"log"
	"sort"
	"strings"
	"sync"

	"github.com/GoogleCloudPlatform/k8s-config-connector/experiments/kompanion/pkg/utils"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

const (
	examples = `
	# lint Config Connector resources across all namespaces, excludes \"kube\" namespaces by default
	kompanion lint

	# exclude certain namespace prefixes
	kompanion lint --exclude-namespaces=kube --exclude-namespaces=my-team

	# target only specific namespace prefixes
	kompanion lint --target-namespaces=my-team

	# include more resource kind in linting
	kompanion lint --include-resources=ComputeNetwork

	# ignore more resource kind in linting
	kompanion lint --ignore-resources=StorageBucket
	`
)

var abandonOnDeleteCheckKind = map[string]bool{
	"AlloyDBCluster":    true,
	"AlloyDBInstance":   true,
	"BigQueryDataset":   true,
	"BigQueryTable":     true,
	"BigtableInstance":  true,
	"ContainerCluster":  true,
	"ContainerNodePool": true,
	"KMKeyRing":         true,
	"KMSCryptoKey":      true,
	"RedisCluster":      true,
	"SpannerDatabase":   true,
	"SpannerInstance":   true,
	"SQLDatabase":       true,
	"SQLInstance":       true,
	"StorageBucket":     true,
}

const (
	kubeconfigFlag = "kubeconfig"

	targetNamespacesFlag = "target-namespaces"
	ignoreNamespacesFlag = "exclude-namespaces"

	includeResourcesFlag = "include-resources"
	ignoreResourcesFlag  = "ignore-resources"

	workerRoutinesFlag = "worker-routines"
)

type LintOptions struct {
	kubeconfig string

	targetNamespaces []string
	ignoreNamespaces []string

	includeResources []string
	ignoreResources  []string

	workerRountines int

	debug bool
}

func (opts *LintOptions) validate() error {
	for _, kind := range opts.includeResources {
		abandonOnDeleteCheckKind[kind] = true
	}

	for _, kind := range opts.ignoreResources {
		// Using delete is more idiomatic for removing an item from a set.
		delete(abandonOnDeleteCheckKind, kind)
	}

	// Collect the keys to be linted so we can sort them for stable output.
	var kindsToLint []string
	for kind := range abandonOnDeleteCheckKind {
		kindsToLint = append(kindsToLint, kind)
	}
	sort.Strings(kindsToLint)
	if opts.debug {
		log.Println("Running lint on the following resource kinds:")
		for _, kind := range kindsToLint {
			log.Printf("* %s", kind)
		}
	}
	if opts.workerRountines <= 0 || opts.workerRountines > 100 {
		return fmt.Errorf("invalid value %d for flag %s. Supported values are [1,100]", opts.workerRountines, workerRoutinesFlag)
	}
	return nil
}

func BuildLintCmd() *cobra.Command {
	var opts LintOptions
	cmd := &cobra.Command{
		Use:     "lint",
		Short:   "lint Config Connector resources",
		Example: examples,
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunLint(cmd.Context(), &opts)
		},
		Args: cobra.ExactArgs(0),
	}

	cmd.Flags().StringVarP(&opts.kubeconfig, kubeconfigFlag, "", opts.kubeconfig, "path to the kubeconfig file.")

	cmd.Flags().StringArrayVarP(&opts.targetNamespaces, targetNamespacesFlag, "", []string{}, "namespace prefix to target the lint tool. Targets all if empty. Can be specified multiple times.")
	cmd.Flags().StringArrayVarP(&opts.ignoreNamespaces, ignoreNamespacesFlag, "", []string{"kube"}, "namespace prefix to ignore. Excludes nothing if empty. Can be specified multiple times. Defaults to \"kube\".")

	cmd.Flags().StringSliceVarP(&opts.includeResources, includeResourcesFlag, "", []string{}, "resource kind to include in the lint tool. Can be specified multiple times.")
	cmd.Flags().StringSliceVarP(&opts.ignoreResources, ignoreResourcesFlag, "", []string{}, "resource kind to ignore in the lint tool. Can be specified multiple times.")
	cmd.Flags().IntVarP(&opts.workerRountines, workerRoutinesFlag, "", 10, "Configure the number of worker routines to lint namespaces with. Defaults to 10. ")
	cmd.Flags().BoolVarP(&opts.debug, "debug", "d", false, "Enable debug output")

	return cmd
}

func getRESTConfig(ctx context.Context, opts *LintOptions) (*rest.Config, error) {
	var loadingRules clientcmd.ClientConfigLoader
	if opts.kubeconfig != "" {
		loadingRules = &clientcmd.ClientConfigLoadingRules{ExplicitPath: opts.kubeconfig}
	} else {
		loadingRules = clientcmd.NewDefaultClientConfigLoadingRules()
	}
	return clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		loadingRules,
		&clientcmd.ConfigOverrides{}).ClientConfig()
}

func RunLint(ctx context.Context, opts *LintOptions) error {
	if err := opts.validate(); err != nil {
		return err
	}

	config, err := getRESTConfig(ctx, opts)
	if err != nil {
		return fmt.Errorf("error building kubeconfig: %w", err)
	}

	if config.QPS == 0 {
		config.QPS = 100
		config.Burst = 20
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return fmt.Errorf("error creating Kubernetes clientset: %sw", err)
	}

	// use the discovery client to iterate over all api resources
	discoveryClient := clientset.Discovery()
	var resources []schema.GroupVersionResource
	resources, err = utils.GetResourcesWithKindFilter(discoveryClient, resources, abandonOnDeleteCheckKind)
	if err != nil {
		return fmt.Errorf("error fetching resources: %w", err)
	}

	namespaces, err := clientset.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
	if err != nil {
		return fmt.Errorf("error fetching namespaces: %w", err)
	}

	// create the work log for go routine workers to use
	q := &taskQueue{}

	r := &Result{}

	dynamicClient, err := dynamic.NewForConfig(config)
	if err != nil {
		return fmt.Errorf("error creating dynamic client: %w", err)
	}

	// Parallize across resources, unless we are scoped to a few namespaces
	// The thought is that if users target a particular namespace (or a few), they may not have cluster-wide permission.
	perNamespace := len(opts.targetNamespaces) > 0
	if perNamespace {
		for _, ns := range namespaces.Items {
			if shouldExclude(ns.Name, opts.ignoreNamespaces, opts.targetNamespaces) {
				continue
			}

			q.AddTask(&lintResourcesTask{
				Result:        r,
				Namespace:     ns.Name,
				Resources:     resources,
				DynamicClient: dynamicClient,
			})
		}
	} else {
		for _, resource := range resources {
			q.AddTask(&lintResourcesTask{
				Result:        r,
				Resources:     []schema.GroupVersionResource{resource},
				DynamicClient: dynamicClient,
			})
		}
	}
	if opts.debug {
		log.Printf("Starting worker threads to process Config Connector objects")
	}
	var wg sync.WaitGroup

	var errs []error
	var errsMutex sync.Mutex

	for i := 0; i < opts.workerRountines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for {
				task := q.GetWork()
				if task == nil {
					// no work
					return
				}

				//log.Printf("Starting work on %v", task)
				if err := task.Run(ctx); err != nil {
					errsMutex.Lock()
					errs = append(errs, err)
					errsMutex.Unlock()
				}
			}
		}()
	}
	wg.Wait()
	r.Print()
	return nil
}

func shouldExclude(name string, excludes []string, includes []string) bool {
	// todo acpana: maps
	for _, exclude := range excludes {
		if strings.Contains(name, exclude) {
			log.Printf("Excluding %s as it contains %s", name, exclude)
			return true
		}
	}

	if len(includes) == 0 {
		return false // no includes means includes all in this case
	}

	for _, include := range includes {
		if strings.Contains(name, include) {
			log.Printf("Including %s as it contains %s", name, include)
			return false
		}
	}

	// by default exclude if nothing that has been defined included this namespace.
	log.Printf("Excluding %s as nothing targets it %s", name, includes)
	return true
}
