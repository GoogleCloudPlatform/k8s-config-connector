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

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/jitter"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/go-logr/logr"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
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

// IsReferenceReady returns whether or not a resource identified by the given GVK
// and NamespacedName is ready. Note that a 'reason' for failure is returned only
// when the resource is not ready and no fatal error has occurred.
func (d *DependencyWatcher) IsReferenceReady(ctx context.Context, refNN types.NamespacedName, refGVK schema.GroupVersionKind) (ok bool, reason string, err error) {
	client := d.dynamicClient.Resource(k8s.ToGVR(refGVK)).Namespace(refNN.Namespace)
	u, err := client.Get(ctx, refNN.Name, metav1.GetOptions{})
	if err != nil {
		if apierrors.IsNotFound(err) {
			return false, "referenced resource does not exist", nil
		}
		return false, "", fmt.Errorf("error getting resource: %v", err)
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

// WatchReferenceUntilReady creates a Watch on a resource identified by the given GVK
// and NamespacedName to monitor any changes on the resource until the resource is
// ready or the Watch expires, whichever come first. If the resource becomes ready,
// the readyHandler will be called and the Watch will be terminated.
func (d *DependencyWatcher) WatchReferenceUntilReady(ctx context.Context, refNN types.NamespacedName, refGVK schema.GroupVersionKind, readyHandler func()) error {
	logger := d.logger.WithValues("reference", refNN).WithValues("referenceGVK", refGVK)

	client := d.dynamicClient.Resource(k8s.ToGVR(refGVK)).Namespace(refNN.Namespace)
	nameSelector := fields.OneTermEqualSelector("metadata.name", refNN.Name).String()
	timeout := int64(jitter.GenerateJitteredReenqueuePeriod().Seconds())
	watch, err := client.Watch(ctx, metav1.ListOptions{FieldSelector: nameSelector, TimeoutSeconds: &timeout})
	if err != nil {
		return fmt.Errorf("error creating watch on resource: %v", err)
	}
	go func() {
		defer watch.Stop()
		for {
			_, ok := <-watch.ResultChan()
			if !ok {
				logger.Info("the watch created by resource on referenced resource was terminated prematurely either due to timeout or error")
				return
			}
			ok, reason, err := d.IsReferenceReady(ctx, refNN, refGVK)
			if err != nil {
				logger.Error(err, "error checking if reference is ready")
				return
			}
			if !ok {
				logger.Info("resource not ready, handler not invoked", "reason", reason)
				continue
			}
			logger.Info("resource is ready; handler invoked")
			readyHandler()
			return
		}
	}()
	return nil
}
