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

package gcpwatch

import (
	"context"
	"sync"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/iam/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/event"
)

type DependencyTracker struct {
	fetcher Fetcher

	gcpResources map[gcpResouceKey]*dependenciesByResource

	controllersMutex sync.Mutex
	controllers      map[string]*ControllerRegistration
}

type ControllerRegistration struct {
	tracker       *DependencyTracker
	controllerKey string
	queue         chan event.GenericEvent
}

func (t *DependencyTracker) RegisterController(key string, queue chan event.GenericEvent) *ControllerRegistration {
	t.controllersMutex.Lock()
	defer t.controllersMutex.Unlock()

	if _, exists := t.controllers[key]; exists {
		klog.Fatalf("multiple controllers with key %q", key)
	}

	registration := &ControllerRegistration{
		controllerKey: key,
		queue:         queue,
		tracker:       t,
	}

	t.controllers[key] = registration

	return registration
}

type gcpResouceKey struct {
	Kind     string
	External string
}

type dependenciesByResource struct {
	etag string

	dependenciesMutex sync.Mutex
	dependencies      []dependency
}

type dependency struct {
	controller *ControllerRegistration
	namespace  string
	name       string
}

type Fetcher interface {
	IsSupported(kind string, external string) bool
	Fetch(ctx context.Context, kind string, external string) (*ResourceInfo, error)
}

type ResourceInfo struct {
	Etag string
}

func NewDependencyTracker(fetcher Fetcher) *DependencyTracker {
	tracker := &DependencyTracker{
		fetcher:      fetcher,
		controllers:  make(map[string]*ControllerRegistration),
		gcpResources: make(map[gcpResouceKey]*dependenciesByResource),
	}

	return tracker
}

type PollConfig struct {
	InitialDelay time.Duration
	MinInterval  time.Duration
	PollInterval time.Duration
}

func (t *DependencyTracker) PollForever(ctx context.Context, pc *PollConfig) {
	nextPoll := time.Now().Add(pc.InitialDelay)
	for {
		if ctx.Err() != nil {
			return
		}

		// todo acpana jitter
		time.Sleep(pc.PollInterval)

		if time.Now().Before(nextPoll) {
			continue
		}

		nextPoll = time.Now().Add(pc.MinInterval)

		if err := t.pollOnce(ctx); err != nil {
			klog.Warningf("error during drift-correction polling: %v", err)
		}
	}
}

func (t *DependencyTracker) copyKeysUnderLock() []gcpResouceKey {
	t.controllersMutex.Lock()
	defer t.controllersMutex.Unlock()

	keys := make([]gcpResouceKey, 0, len(t.gcpResources))
	for key := range t.gcpResources {
		keys = append(keys, key)
	}
	return keys
}

func (t *DependencyTracker) getGCPResourcesUnderLock(key gcpResouceKey) *dependenciesByResource {
	t.controllersMutex.Lock()
	defer t.controllersMutex.Unlock()

	return t.gcpResources[key]
}

func (t *DependencyTracker) pollOnce(ctx context.Context) error {
	log := klog.FromContext(ctx)

	keys := t.copyKeysUnderLock()

	for _, key := range keys {
		gcpResource := t.getGCPResourcesUnderLock(key)
		if gcpResource == nil {
			continue
		}

		latest, err := t.fetcher.Fetch(ctx, key.Kind, key.External)
		if err != nil {
			// TODO: Remove if not found?
			log.Error(err, "error fetching iam", "kind", key.Kind, "external", key.External)
			continue
		}

		maybeNotifyDependenciesUnderLock(ctx, gcpResource, key, latest)
	}

	return nil
}

func maybeNotifyDependenciesUnderLock(ctx context.Context, gcpResource *dependenciesByResource, key gcpResouceKey, latest *ResourceInfo) {
	log := klog.FromContext(ctx)

	gcpResource.dependenciesMutex.Lock()
	defer gcpResource.dependenciesMutex.Unlock()

	log.Info("got iam etag", "newEtag", latest.Etag, "kind", key.Kind, "external", key.External, "oldEtag", gcpResource.etag)
	if latest.Etag != gcpResource.etag {
		gcpResource.etag = latest.Etag

		for _, dep := range gcpResource.dependencies {
			log.Info("triggering notification", "controller", dep.controller.controllerKey, "namespace", dep.namespace, "name", dep.name)
			genEvent := event.GenericEvent{}
			genEvent.Object = &unstructured.Unstructured{}
			genEvent.Object.SetNamespace(dep.namespace)
			genEvent.Object.SetName(dep.name)
			dep.controller.queue <- genEvent
		}
	}
}

// todo acpana expose reason why not added?
func (r *ControllerRegistration) Add(ctx context.Context, policy *v1beta1.IAMPartialPolicy, etag string) bool {
	log := klog.FromContext(ctx)
	if r == nil {
		return false
	}

	if etag == "" {
		return false
	}

	// TODO: Resolve Name -> External

	t := r.tracker
	kind := policy.Spec.ResourceReference.Kind
	if kind == "" {
		return false
	}

	external := policy.Spec.ResourceReference.External
	if external == "" {
		return false
	}

	if !t.fetcher.IsSupported(kind, external) {
		return false
	}

	log.Info("adding watch on policy", "policy", policy, "kind", kind, "external", external)

	t.addUnderLock(kind, external, r, policy.GetNamespace(), policy.GetName(), etag)
	return true
}

func (t *DependencyTracker) addUnderLock(kind string, external string, controller *ControllerRegistration, namespace string, name string, etag string) {
	target := t.getTarget(kind, external)

	target.dependenciesMutex.Lock()
	defer target.dependenciesMutex.Unlock()

	if target.etag == "" {
		target.etag = etag
		// TODO: Trigger now?
	}

	for i := range target.dependencies {
		dep := &target.dependencies[i]
		if dep.controller == controller && dep.namespace == namespace && dep.name == name {
			return
		}
	}
	target.dependencies = append(target.dependencies, dependency{
		controller: controller,
		namespace:  namespace,
		name:       name,
	})
}

func (t *DependencyTracker) getTarget(kind string, external string) *dependenciesByResource {
	t.controllersMutex.Lock()
	defer t.controllersMutex.Unlock()

	key := gcpResouceKey{
		Kind:     kind,
		External: external,
	}

	target := t.gcpResources[key]
	if target == nil {
		target = &dependenciesByResource{}
		t.gcpResources[key] = target
	}
	return target
}
