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

package lint

import (
	"context"
	"fmt"
	"log"
	"sort"
	"strings"
	"sync"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/cmd/lint/options"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog/v2"
)

func Execute(ctx context.Context, opts *options.Options) error {
	kindsToLint, err := opts.Validate()
	if err != nil {
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
		return fmt.Errorf("error creating Kubernetes clientset: %w", err)
	}

	discoveryClient := clientset.Discovery()
	var resources []schema.GroupVersionResource
	resources, err = getResourcesWithKindFilter(discoveryClient, resources, kindsToLint)
	if err != nil {
		return fmt.Errorf("error fetching resources: %w", err)
	}

	namespaces, err := clientset.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
	if err != nil {
		return fmt.Errorf("error fetching namespaces: %w", err)
	}

	q := &taskQueue{}
	r := &Result{}

	dynamicClient, err := dynamic.NewForConfig(config)
	if err != nil {
		return fmt.Errorf("error creating dynamic client: %w", err)
	}

	perNamespace := len(opts.TargetNamespaces) > 0
	if perNamespace {
		for _, ns := range namespaces.Items {
			if shouldExclude(ns.Name, opts.IgnoreNamespaces, opts.TargetNamespaces) {
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

	if opts.Verbose > 0 {
		log.Printf("Starting worker threads to process Config Connector objects")
	}

	var wg sync.WaitGroup
	var errs []error
	var errsMutex sync.Mutex

	for i := 0; i < opts.WorkerRoutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for {
				task := q.GetWork()
				if task == nil {
					return
				}

				if err := task.Run(ctx); err != nil {
					errsMutex.Lock()
					errs = append(errs, err)
					errsMutex.Unlock()
				}
			}
		}()
	}
	wg.Wait()

	if len(errs) > 0 {
		for _, err := range errs {
			log.Printf("Error: %v", err)
		}
	}

	r.Print()
	return nil
}

func getRESTConfig(ctx context.Context, opts *options.Options) (*rest.Config, error) {
	var loadingRules clientcmd.ClientConfigLoader
	if opts.Kubeconfig != "" {
		loadingRules = &clientcmd.ClientConfigLoadingRules{ExplicitPath: opts.Kubeconfig}
	} else {
		loadingRules = clientcmd.NewDefaultClientConfigLoadingRules()
	}
	return clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		loadingRules,
		&clientcmd.ConfigOverrides{}).ClientConfig()
}

func shouldExclude(name string, excludes []string, includes []string) bool {
	for _, exclude := range excludes {
		if strings.Contains(name, exclude) {
			log.Printf("Excluding %s as it contains %s", name, exclude)
			return true
		}
	}

	if len(includes) == 0 {
		return false
	}

	for _, include := range includes {
		if strings.Contains(name, include) {
			log.Printf("Including %s as it contains %s", name, include)
			return false
		}
	}

	log.Printf("Excluding %s as nothing targets it %v", name, includes)
	return true
}

func getResourcesWithKindFilter(discoveryClient discovery.DiscoveryInterface, resources []schema.GroupVersionResource, kinds map[string]bool) ([]schema.GroupVersionResource, error) {
	apiResourceLists, err := discoveryClient.ServerPreferredResources()
	if err != nil {
		return nil, fmt.Errorf("failed to get preferred resources: %w", err)
	}

	for _, apiResourceList := range apiResourceLists {
		if !strings.Contains(apiResourceList.GroupVersion, ".cnrm.cloud.google.com/") {
			continue
		}

		apiResourceListGroupVersion, err := schema.ParseGroupVersion(apiResourceList.GroupVersion)
		if err != nil {
			klog.Warningf("skipping unparseable groupVersion %q", apiResourceList.GroupVersion)
			continue
		}

		for _, apiResource := range apiResourceList.APIResources {
			if !apiResource.Namespaced {
				continue
			}
			if !contains(apiResource.Verbs, "list") {
				continue
			}
			if !kinds[apiResource.Kind] {
				continue
			}
			gvr := schema.GroupVersionResource{
				Group:    apiResource.Group,
				Version:  apiResource.Version,
				Resource: apiResource.Name,
			}

			if gvr.Group == "" {
				gvr.Group = apiResourceListGroupVersion.Group
			}

			if gvr.Version == "" {
				gvr.Version = apiResourceListGroupVersion.Version
			}

			resources = append(resources, gvr)
		}
	}
	sort.Slice(resources, func(i, j int) bool {
		return resources[i].String() < resources[j].String()
	})
	return resources, nil
}

func contains(slice []string, str string) bool {
	for _, s := range slice {
		if strings.ToLower(s) == strings.ToLower(str) {
			return true
		}
	}
	return false
}

type Result struct {
	lock sync.Mutex
	// Map[namespace, Map[gk, []resourceID]]
	resources map[string]map[string][]string
}

func (r *Result) Print() {
	log.Println("Following Resources should include `cnrm.cloud.google.com/deletion-policy: abandon` annotation")
	r.lock.Lock()
	defer r.lock.Unlock()
	for ns, nsMap := range r.resources {
		log.Printf("Namespace: %s\n", ns)
		for gk, ids := range nsMap {
			log.Printf("Group Kind: %s\n", gk)
			for _, id := range ids {
				log.Printf("- %s\n", id)
			}
		}
	}
}

func (r *Result) addNewResource(namespace string, gk string, id string) {
	r.lock.Lock()
	defer r.lock.Unlock()
	if r.resources == nil {
		r.resources = make(map[string]map[string][]string)
	}

	nsMap, present := r.resources[namespace]
	if !present {
		nsMap = make(map[string][]string)
		r.resources[namespace] = nsMap
	}
	gkMap, present := nsMap[gk]
	if !present {
		gkMap = []string{id}
		nsMap[gk] = gkMap
	} else {
		gkMap = append(gkMap, id)
		nsMap[gk] = gkMap
	}
}

type lintResourcesTask struct {
	Result *Result

	Namespace string

	DynamicClient dynamic.Interface

	Resources []schema.GroupVersionResource
}

func (t *lintResourcesTask) Run(ctx context.Context) error {
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
			if !(r.GetAnnotations()["cnrm.cloud.google.com/deletion-policy"] == "abandon") {
				gk := gvr.Resource + "." + gvr.Group
				t.Result.addNewResource(t.Namespace, gk, r.GetName())
			}
		}
	}

	return nil
}

type Task interface {
	Run(ctx context.Context) error
}

type taskQueue struct {
	mu    sync.Mutex
	tasks []Task
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
