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

package resourcewatcher

import (
	"context"
	"fmt"

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
)

type ResourceWatcher struct {
	dynamicClient dynamic.Interface
	logger        logr.Logger
}

// New creates a new ResourceWatcher that uses a dynamic client
// to monitor the status of requested resources
func New(clientConfig *rest.Config, logger logr.Logger) (*ResourceWatcher, error) {
	dynamicClient, err := dynamic.NewForConfig(clientConfig)
	if err != nil {
		return nil, err
	}
	return NewWithClient(dynamicClient, logger), nil
}

func NewWithClient(dynamicClient dynamic.Interface, logger logr.Logger) *ResourceWatcher {
	return &ResourceWatcher{
		dynamicClient: dynamicClient,
		logger:        logger.WithName("resourcewatcher"),
	}
}

// isResourceReady returns whether a resource identified by the given GVK
// and NamespacedName is ready. Note that a 'reason' for failure is returned only
// when the resource is not ready and no fatal error has occurred.
func isResourceReady(event watch.Event) (ok bool, reason string, err error) {
	if event.Type != watch.Modified && event.Type != watch.Added {
		return false, fmt.Sprintf("got watch event of type '%v', want event type '%v' or '%v'", event.Type, watch.Modified, watch.Added), nil
	}
	u, ok := event.Object.(*unstructured.Unstructured)
	if !ok {
		return false, "", fmt.Errorf("error casting event object '%v' of kind '%v' to unstructured", event.Object, event.Object.GetObjectKind())
	}
	resource, err := k8s.NewResource(u)
	if err != nil {
		return false, "", fmt.Errorf("error converting unstructured to resource: %w", err)
	}
	if !k8s.IsResourceReady(resource) {
		return false, "resource not ready", nil
	}
	return true, "", nil
}

// WaitForResourceToBeReady waits for the resource identified by the given GVK
// and NamespacedName. It blocks until the resource is ready, an error occurs, or a context
// cancellation occurs. Note that a nil return value signifies that the resource is ready and
// no errors have occurred.
func (r *ResourceWatcher) WaitForResourceToBeReady(ctx context.Context, nn types.NamespacedName, gvk schema.GroupVersionKind) error {
	logger := r.logger.WithValues("resource", nn, "resourceGVK", gvk)
	client := r.dynamicClient.Resource(k8s.ToGVR(gvk)).Namespace(nn.Namespace)
	nameSelector := fields.OneTermEqualSelector("metadata.name", nn.Name).String()
	watch, err := client.Watch(ctx, metav1.ListOptions{FieldSelector: nameSelector})
	if err != nil {
		return fmt.Errorf("error creating watch on resource: %w", err)
	}
	logger.Info("successfully created watch on resource")
	defer watch.Stop()
	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf("context was cancelled: %v", ctx.Err())
		case event, ok := <-watch.ResultChan():
			if !ok {
				return fmt.Errorf("watch channel was closed")
			}
			ok, reason, err := isResourceReady(event)
			if err != nil {
				return fmt.Errorf("error checking if resource is ready: %w", err)
			}
			if !ok {
				logger.Info("resource not ready", "reason", reason)
				continue
			}
			logger.Info("resource is ready")
			return nil
		}
	}
}
