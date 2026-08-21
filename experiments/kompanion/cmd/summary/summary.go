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
	"context"
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/GoogleCloudPlatform/k8s-config-connector/experiments/kompanion/pkg/utils"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

const (
	examples = `
	# Summarize Config Connector resources across all namespaces, excludes \"kube\" namespaces by default
	kompanion summary
	`
)

func BuildSummaryCmd() *cobra.Command {
	opts := NewSummaryOptions()
	cmd := &cobra.Command{
		Use:     "summary",
		Short:   "summarize Config Connector resources",
		Example: examples,
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunSummarize(cmd.Context(), opts)
		},
		Args: cobra.ExactArgs(0),
	}

	flags := cmd.Flags()
	opts.AddFlags(flags)

	return cmd
}

type Result struct {
	// Map[namespace, Map[gk, Map[status, count]]]
	counts map[string]map[string]map[string]int
	// Map[gk, Map[status, count]]
	totals map[string]map[string]int
	// secures access to counts & totals
	countLock sync.Mutex
}

func (r *Result) increment(namespace string, gk string, status string) {
	r.countLock.Lock()
	defer r.countLock.Unlock()
	if r.counts == nil {
		r.counts = make(map[string]map[string]map[string]int)
	}
	nsMap, present := r.counts[namespace]
	if !present {
		nsMap = make(map[string]map[string]int)
		r.counts[namespace] = nsMap
	}
	gkMap, present := nsMap[gk]
	if !present {
		gkMap = make(map[string]int)
		nsMap[gk] = gkMap
	}
	count := gkMap[status]
	// Can ignore not present because count will default to 0.
	gkMap[status] = count + 1

	if r.totals == nil {
		r.totals = make(map[string]map[string]int)
	}
	gkMap, present = r.totals[gk]
	if !present {
		gkMap = make(map[string]int)
		r.totals[gk] = gkMap
	}
	count = gkMap[status]
	// Can ignore not present because count will default to 0.
	gkMap[status] = count + 1
}

func (r *Result) Print() {
	r.countLock.Lock()
	defer r.countLock.Unlock()
	for ns, nsMap := range r.counts {
		log.Printf("Namespace: %s\n", ns)
		for gk, gkMap := range nsMap {
			first := true
			info := "("
			for status, count := range gkMap {
				if first {
					info += fmt.Sprintf("%s: %d", status, count)
					first = false
				} else {
					info += fmt.Sprintf(", %s: %d", status, count)
				}
			}
			info += ")"
			log.Printf("- GroupKind: %s %s\n", gk, info)
		}
	}
	log.Printf("Totals: \n")
	for gk, gkMap := range r.totals {
		first := true
		info := "("
		for status, count := range gkMap {
			if first {
				info += fmt.Sprintf("%s: %d", status, count)
				first = false
			} else {
				info += fmt.Sprintf(", %s: %d", status, count)
			}
		}
		info += ")"
		log.Printf("- GroupKind: %s %s\n", gk, info)
	}
}

// Task is implemented by our namespace-collection routine, or anything else we want to run in parallel.
type Task interface {
	Run(ctx context.Context) error
}

// tracks the namespaces to be exported.
// thread safe.
type taskQueue struct {
	mu    sync.Mutex
	tasks []Task // will be treated as a FIFO queue
}

func (n *taskQueue) GetWork() Task {
	n.mu.Lock()
	defer n.mu.Unlock()

	if len(n.tasks) == 0 {
		return nil
	}

	workItem := n.tasks[0]
	n.tasks = n.tasks[1:]

	return workItem
}

func (n *taskQueue) AddTask(t Task) {
	n.mu.Lock()
	defer n.mu.Unlock()

	n.tasks = append(n.tasks, t)
}

type dumpResourcesTask struct {
	// Result accumulates the overall summary
	Summary *Result

	// Namespace is the namespace to filter down
	Namespace string

	DynamicClient *dynamic.DynamicClient

	// Resources is the list of resources to query
	Resources []schema.GroupVersionResource
}

func (t *dumpResourcesTask) Run(ctx context.Context) error {
	// klog := klog.FromContext(ctx)

	for _, gvr := range t.Resources {
		var resources *unstructured.UnstructuredList
		if t.Namespace != "" {
			r, err := t.DynamicClient.Resource(gvr).Namespace(t.Namespace).List(ctx, metav1.ListOptions{})
			if err != nil {
				return fmt.Errorf("fetching gvr %s resources in namespace %s: %w", gvr, t.Namespace, err)
			}
			resources = r

		} else {
			r, err := t.DynamicClient.Resource(gvr).List(ctx, metav1.ListOptions{})
			if err != nil {
				return fmt.Errorf("fetching gvr %s resources: %w", gvr, err)
			}
			resources = r
		}

		for _, r := range resources.Items {
			status, err := getStatus(r, gvr)
			if err != nil {
				log.Printf("Got error retrieving status of type gvr %s, %v", gvr, err)
			}
			gk := gvr.Resource + "." + gvr.Group
			t.Summary.increment(r.GetNamespace(), gk, status)
		}
	}

	return nil
}

func getRESTConfig(ctx context.Context, opts *SummaryOptions) (*rest.Config, error) {
	var loadingRules clientcmd.ClientConfigLoader
	if opts.kubeconfig != "" {
		loadingRules = &clientcmd.ClientConfigLoadingRules{ExplicitPath: opts.kubeconfig}
	} else {
		loadingRules = clientcmd.NewDefaultClientConfigLoadingRules()
	}

	return clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		loadingRules,
		&clientcmd.ConfigOverrides{
			// ClusterInfo: clientcmdapi.Cluster{Server: masterUrl},
		}).ClientConfig()
}

func RunSummarize(ctx context.Context, opts *SummaryOptions) error {
	log.Printf("Running kompanion summary with kubeconfig: %s", opts.kubeconfig)

	if err := opts.validateFlags(); err != nil {
		opts.Print()
		return err
	}

	config, err := getRESTConfig(ctx, opts)
	if err != nil {
		return fmt.Errorf("error building kubeconfig: %w", err)
	}

	// We rely more on server-side rate limiting now, so give it a high client-side QPS
	if config.QPS == 0 {
		config.QPS = 100
		config.Burst = 20
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return fmt.Errorf("error creating Kubernetes clientset: %sw", err)
	}

	dynamicClient, err := dynamic.NewForConfig(config)
	if err != nil {
		return fmt.Errorf("error creating dynamic client: %w", err)
	}

	// use the discovery client to iterate over all api resources
	discoveryClient := clientset.Discovery()
	var resources []schema.GroupVersionResource
	resources, err = utils.GetResources(discoveryClient, resources)
	if err != nil {
		return fmt.Errorf("error fetching resources: %w", err)
	}

	namespaces, err := clientset.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
	if err != nil {
		return fmt.Errorf("error fetching namespaces: %w", err)
	}

	// create the work log for go routine workers to use
	q := &taskQueue{}

	summary := &Result{}

	// Parallize across resources, unless we are scoped to a few namespaces
	// The thought is that if users target a particular namespace (or a few), they may not have cluster-wide permission.
	perNamespace := len(opts.targetNamespaces) > 0
	if perNamespace {
		for _, ns := range namespaces.Items {
			if shouldExclude(ns.Name, opts.ignoreNamespaces, opts.targetNamespaces) {
				continue
			}

			q.AddTask(&dumpResourcesTask{
				Summary:       summary,
				Namespace:     ns.Name,
				Resources:     resources,
				DynamicClient: dynamicClient,
			})
		}
	} else {
		for _, resource := range resources {
			q.AddTask(&dumpResourcesTask{
				Summary:       summary,
				Resources:     []schema.GroupVersionResource{resource},
				DynamicClient: dynamicClient,
			})
		}
	}

	log.Printf("Starting worker threads to process Config Connector objects")
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

	log.Printf("Dumping Config Connector summaries")
	wg.Wait()

	summary.Print()

	return nil
}

func getStatus(r unstructured.Unstructured, gvr schema.GroupVersionResource) (string, error) {
	id := types.NamespacedName{
		Namespace: r.GetNamespace(),
		Name:      r.GetName(),
	}
	slices, present, err := unstructured.NestedSlice(r.Object, "status", "conditions")
	if !present {
		return "no status", nil
	}
	if err != nil {
		return "error", fmt.Errorf("error retrieving .status.healthy from resource %s: %w", id, err)
	}
	if len(slices) < 1 {
		return "no status", nil
	} else if len(slices) > 1 {
		return "unknown", nil
	}
	element := slices[0]
	condition, ok := element.(map[string]interface{})
	if !ok {
		return "unknown", fmt.Errorf("resource %s of type gvr %s failed to cast %T to condition", id, gvr, element)
	}
	val, present, err := unstructured.NestedString(condition, "status")
	if !present {
		return "unknown", nil
	}
	if err != nil {
		return "unknown", fmt.Errorf("error retrieving .status.healthy from resource %s: %w", id, err)
	}
	return val, nil
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
