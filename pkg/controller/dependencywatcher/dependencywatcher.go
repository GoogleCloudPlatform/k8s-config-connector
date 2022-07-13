// Copyright 2022 Google LLC
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

package dependencywatcher

import (
	"context"
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"

	"github.com/go-logr/logr"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
	klog "sigs.k8s.io/controller-runtime/pkg/log"
)

type DependencyWatcher struct {
	dynamicClient dynamic.Interface
	logger        logr.Logger
	resource      *k8s.Resource
}

// CreateWatchForResource creates a DependencyWatcher that uses a dynamicClient to monitor the
// dependencies of a given resource
func CreateWatchForResource(resource *k8s.Resource, clientConfig *rest.Config) (*DependencyWatcher, error) {
	dynamicClient, err := dynamic.NewForConfig(clientConfig)
	if err != nil {
		return nil, err
	}
	return CreateWatchForResourceWithClient(resource, dynamicClient), nil
}

func CreateWatchForResourceWithClient(resource *k8s.Resource, dynamicClient dynamic.Interface) *DependencyWatcher {
	dependencyWatcherName := fmt.Sprintf("%v-dependencywatcher", strings.ToLower(resource.GroupVersionKind().Kind))
	return &DependencyWatcher{
		dynamicClient: dynamicClient,
		resource:      resource,
		logger:        klog.Log.WithName(dependencyWatcherName).WithValues("resource", resource.GetNamespacedName(), "resourceGVK", resource.GroupVersionKind()),
	}
}

// isReferenceReady returns whether or not a resource identified by the given GVK
// and NamespacedName is ready. Note that a 'reason' for failure is returned only
// when the resource is not ready and no fatal error has occurred.
func isReferenceReady(event watch.Event) (ok bool, reason string, err error) {
	if event.Type != watch.Modified && event.Type != watch.Added {
		return false, fmt.Sprintf("got watch event of type '%v', want event type '%v' or '%v'", event.Type, watch.Modified, watch.Added), nil
	}
	u, ok := event.Object.(*unstructured.Unstructured)
	if !ok {
		return false, "", fmt.Errorf("error casting event object '%v' of kind '%v' to unstructured", event.Object, event.Object.GetObjectKind())
	}
	refResource, err := k8s.NewResource(u)
	if err != nil {
		return false, "", fmt.Errorf("error converting unstructured to resource: %v", err)
	}
	if !k8s.IsResourceReady(refResource) {
		return false, "referenced resource not ready", nil
	}
	return true, "", nil
}

// WaitForReferenceToBeReady waits for the resource identified by the given GVK
// and NamespacedName. It blocks until the resource is ready, an error occurs, or a context
// cancellation occurs. Note that a nil return value signifies that the resource is ready and
// no errors have occurred.
func (d *DependencyWatcher) WaitForReferenceToBeReady(ctx context.Context, refNN types.NamespacedName, refGVK schema.GroupVersionKind) error {
	logger := d.logger.WithValues("reference", refNN).WithValues("referenceGVK", refGVK)

	client := d.dynamicClient.Resource(k8s.ToGVR(refGVK)).Namespace(refNN.Namespace)
	nameSelector := fields.OneTermEqualSelector("metadata.name", refNN.Name).String()
	watch, err := client.Watch(ctx, metav1.ListOptions{FieldSelector: nameSelector})
	if err != nil {
		return fmt.Errorf("error creating watch on reference: %v", err)
	}
	defer watch.Stop()
	logger.Info("successfully created watch on reference")
	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf("context was cancelled: %v", ctx.Err())
		case event, ok := <-watch.ResultChan():
			if !ok {
				return fmt.Errorf("watch channel was closed")
			}
			ok, reason, err := isReferenceReady(event)
			if err != nil {
				return fmt.Errorf("error checking if reference is ready: %v", err)
			}
			if !ok {
				logger.Info("reference not ready", "reason", reason)
				continue
			}
			logger.Info("reference is ready")
			return nil
		}
	}
}
