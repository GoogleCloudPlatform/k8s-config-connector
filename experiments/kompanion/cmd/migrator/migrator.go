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

package migrator

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/GoogleCloudPlatform/k8s-config-connector/experiments/kompanion/pkg/utils"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
)

const (
	examples = `
	# export Config Connector resources across all namespaces, excludes \"kube\" namespaces by default
	kompanion migrator

	# exclude certain namespace prefixes
	kompanion migrator --exclude-namespaces=kube --exclude-namespaces=my-team

	# target only specific namespace prefixes
	kompanion migrator --target-namespaces=my-team

	# target only specific namespace prefixes AND specific object prefixes
	kompanion migrator --target-namespaces=my-team --target-objects=logging
	`
)

func BuildMigratorCmd() *cobra.Command {
	opts := NewMigratorOptions()
	cmd := &cobra.Command{
		Use:     "migrator",
		Short:   "migrates Config Connector resources",
		Example: examples,
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunMigrator(cmd.Context(), opts)
		},
		Args: cobra.ExactArgs(0),
	}

	flags := cmd.Flags()
	opts.AddFlags(flags)

	return cmd
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
	// Namespace is the namespace to filter down
	Namespace string

	DynamicClient *dynamic.DynamicClient

	OutputChannel chan string

	// Resources is the list of resources to query
	Resources []schema.GroupVersionResource

	ShouldExcludeObject func(id types.NamespacedName) bool
}

func (t *dumpResourcesTask) Run(ctx context.Context) error {
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
			id := types.NamespacedName{
				Namespace: r.GetNamespace(),
				Name:      r.GetName(),
			}
			if t.ShouldExcludeObject(id) {
				continue
			}
			r = pruneResource(r)

			data, err := yaml.Marshal(r)
			if err != nil {
				return fmt.Errorf("error marshalling resource %s: %w", id, err)
			}
			t.OutputChannel <- string(data)
		}
	}

	return nil
}

func pruneResource(r unstructured.Unstructured) unstructured.Unstructured {
	newAnnotations := make(map[string]string)
	annotations := r.GetAnnotations()
	for key, value := range annotations {
		if !strings.Contains(key, "cnrm.cloud.google.com") {
			continue
		}
		newAnnotations[key] = value
	}
	if len(newAnnotations) == 0 {
		r.SetAnnotations(nil)
	} else {
		r.SetAnnotations(newAnnotations)
	}

	r.SetCreationTimestamp(metav1.Time{})
	r.SetFinalizers(nil)
	r.SetGeneration(0)
	r.SetManagedFields(nil)
	r.SetResourceVersion("")
	r.SetUID("")

	unstructured.RemoveNestedField(r.Object, "status", "conditions")
	unstructured.RemoveNestedField(r.Object, "status", "creationTimestamp")
	unstructured.RemoveNestedField(r.Object, "status", "healthy")
	unstructured.RemoveNestedField(r.Object, "status", "observedGeneration")
	unstructured.RemoveNestedField(r.Object, "status", "selfLink") // Are we sure?
	status, present, err := unstructured.NestedMap(r.Object, "status")
	if err != nil {
		log.Fatalf("Could not get the status field %v\r", err)
	}
	if present && len(status) == 0 {
		unstructured.RemoveNestedField(r.Object, "status")
	}

	return r
}

func RunMigrator(ctx context.Context, opts *MigratorOptions) error {
	log.Printf("Running kompanion export with kubeconfig: %s", opts.Kubeconfig)

	if err := opts.validateFlags(); err != nil {
		return err
	}

	config, err := utils.GetRESTConfig(ctx, opts.Kubeconfig)
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

	outputChannel := make(chan string, 2*opts.WorkerRoutines)

	namespaces, err := clientset.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
	if err != nil {
		return fmt.Errorf("error fetching namespaces: %w", err)
	}
	shouldExcludeObject := func(id types.NamespacedName) bool {
		if shouldExclude(id.Namespace, opts.IgnoreNamespaces, opts.TargetNamespaces) {
			return true
		}
		return shouldExclude(id.Name, opts.IgnoreObjects, opts.TargetObjects)
	}

	// write the namespaces to the migrator file
	filedir := filepath.Dir(opts.MigrationFile)
	if err := os.MkdirAll(filedir, 0o700); err != nil {
		return fmt.Errorf("error creating directory %s: %w", filedir, err)
	}
	migfile, err := os.Create(opts.MigrationFile)
	if err != nil {
		return fmt.Errorf("error creating migration file %s: %w", opts.MigrationFile, err)
	}
	defer migfile.Close()
	nsGvr := schema.GroupVersionResource{
		Group:    "",
		Version:  "v1",
		Resource: "namespaces",
	}
	nss, err := dynamicClient.Resource(nsGvr).List(ctx, metav1.ListOptions{})
	if err != nil {
		return fmt.Errorf("fetching namespace %s resources: %w", nsGvr, err)
	}
	for _, ns := range nss.Items {
		ns = pruneResource(ns)
		if len(ns.GetAnnotations()) <= 0 {
			continue
		}
		data, err := yaml.Marshal(ns)
		if err != nil {
			return fmt.Errorf("error marshalling namespace: %w", err)
		}
		migfile.WriteString(string(data[:]))
		migfile.WriteString("---\n")
	}

	// create the work log for go routine workers to use
	q := &taskQueue{}

	// Parallize across resources, unless we are scoped to a few namespaces
	// The thought is that if users target a particular namespace (or a few), they may not have cluster-wide permission.
	perNamespace := len(opts.TargetNamespaces) > 0
	if perNamespace {
		for _, ns := range namespaces.Items {
			if shouldExclude(ns.Name, opts.IgnoreNamespaces, opts.TargetNamespaces) {
				continue
			}

			q.AddTask(&dumpResourcesTask{
				Namespace:     ns.Name,
				Resources:     resources,
				DynamicClient: dynamicClient,
				OutputChannel: outputChannel,
				// ReportDir:           reportTempDir,
				ShouldExcludeObject: shouldExcludeObject,
			})
		}
	} else {
		for _, resource := range resources {
			q.AddTask(&dumpResourcesTask{
				Resources:     []schema.GroupVersionResource{resource},
				DynamicClient: dynamicClient,
				OutputChannel: outputChannel,
				// ReportDir:           reportTempDir,
				ShouldExcludeObject: shouldExcludeObject,
			})
		}
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
					// no work
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

	go func() {
		for data := range outputChannel {
			migfile.WriteString(data)
			migfile.WriteString("---\n")
		}
	}()

	wg.Wait()
	close(outputChannel)
	return nil
}

func shouldExclude(name string, excludes []string, includes []string) bool {
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
