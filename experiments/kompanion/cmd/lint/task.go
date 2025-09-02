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
	"sync"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
)

type lintResourcesTask struct {
	Result *Result

	// Namespace is the namespace to filter down
	Namespace string

	DynamicClient *dynamic.DynamicClient

	// Resources is the list of resources to query
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
