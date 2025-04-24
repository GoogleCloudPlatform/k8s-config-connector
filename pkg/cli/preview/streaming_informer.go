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

package preview

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	toolscache "k8s.io/client-go/tools/cache"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// streamingInformer is an informer that streams events from the Kubernetes API server.
// It implements cache.Informer, but allows us to intercept / mutate operations.
type streamingInformer struct {
	streamingClient *StreamingClient
	typeInfo        *typeInfo
	mutex           sync.Mutex

	eventHandlerRegistrations []*eventHandlerRegistration

	resyncPeriod time.Duration

	hasSynced atomic.Bool
	objects   objects
}

// objects is a map of objects by their namespaced name.
type objects struct {
	mutex sync.Mutex
	store map[types.NamespacedName]Object
}

// OnListObject is called from list, the lock should be held
func (o *objects) OnListObject(obj Object, isInInitialList bool, eventHandlerRegistrations []*eventHandlerRegistration) {
	id := types.NamespacedName{
		Namespace: obj.GetNamespace(),
		Name:      obj.GetName(),
	}
	o.store[id] = obj
	for _, handler := range eventHandlerRegistrations {
		handler.handler.OnAdd(obj, false)
	}
}

// OnWatchAdd is called from watch, the lock is not held
func (o *objects) OnWatchAdd(obj Object, eventHandlerRegistrations []*eventHandlerRegistration) {
	id := types.NamespacedName{
		Namespace: obj.GetNamespace(),
		Name:      obj.GetName(),
	}

	o.mutex.Lock()
	defer o.mutex.Unlock()

	o.store[id] = obj
	for _, handler := range eventHandlerRegistrations {
		handler.handler.OnAdd(obj, false)
	}
}

var _ cache.Informer = &streamingInformer{}

// newStreamingInformer creates a new streaming informer.
func newStreamingInformer(streamingClient *StreamingClient, typeInfo *typeInfo) (*streamingInformer, error) {
	s := &streamingInformer{
		streamingClient: streamingClient,
		typeInfo:        typeInfo,
	}
	s.objects = objects{
		store: make(map[types.NamespacedName]Object),
	}
	return s, nil
}

// Start starts the informer.
func (i *streamingInformer) Start(ctx context.Context) error {
	go i.run(ctx)
	return nil
}

// run runs the informer.
func (i *streamingInformer) run(ctx context.Context) {
	log := klog.FromContext(ctx)
	for {
		err := i.runOnce(ctx)
		if err != nil {
			// If the context is closed, don't log as much
			if err := ctx.Err(); err != nil {
				log.V(2).Info("context closed; stopping informer")
				return
			}
			log.Error(err, "running list/watch for informer (will retry)")
		}

		// TODO: Backoff
		time.Sleep(2 * time.Second)
	}
}

// runOnce runs the informer once.
func (i *streamingInformer) runOnce(ctx context.Context) error {
	listMetadata, err := i.doList(ctx)
	if err != nil {
		return err
	}

	i.hasSynced.Store(true)
	watchListener := &watchListener{objects: &i.objects, eventHandlerRegistrations: i.eventHandlerRegistrations}
	watchOptions := WatchOptions{
		ResourceVersion:     listMetadata.ResourceVersion,
		AllowWatchBookmarks: true,
	}
	if err := i.streamingClient.Watch(ctx, i.typeInfo, watchOptions, watchListener); err != nil {
		return err
	}

	return fmt.Errorf("watch finished unexpectedly")
}

// doList lists the objects for the given type.
func (i *streamingInformer) doList(ctx context.Context) (*ListMetadata, error) {
	i.objects.mutex.Lock()
	defer i.objects.mutex.Unlock()

	isInInitialList := !i.hasSynced.Load()
	listListener := &listListener{objects: &i.objects, isInInitialList: isInInitialList, eventHandlerRegistrations: i.eventHandlerRegistrations}
	if err := i.streamingClient.List(ctx, i.typeInfo, listListener); err != nil {
		return nil, err
	}
	return &listListener.metadata, nil
}

// listListener is a listener for list operations.
type listListener struct {
	isInInitialList           bool
	eventHandlerRegistrations []*eventHandlerRegistration
	objects                   *objects
	metadata                  ListMetadata
}

// OnListBegin is called when the list operation begins.
func (i *listListener) OnListBegin(metadata ListMetadata) {
	i.metadata = metadata
}

// OnListObject is called when an object is listed.
func (i *listListener) OnListObject(obj Object) error {
	i.objects.OnListObject(obj, i.isInInitialList, i.eventHandlerRegistrations)
	return nil
}

// OnListEnd is called when the list operation ends.
func (i *listListener) OnListEnd() {
}

type watchListener struct {
	objects                   *objects
	eventHandlerRegistrations []*eventHandlerRegistration
}

// OnWatchEvent is called when a watch event occurs.
func (i *watchListener) OnWatchEvent(eventType string, obj Object) error {
	switch eventType {
	case "ADDED":
		i.objects.OnWatchAdd(obj, i.eventHandlerRegistrations)
		return nil
	case "BOOKMARK":
		klog.Infof("BOOKMARK %+v", obj)
		return nil
	default:
		return fmt.Errorf("unknown event type: %q", eventType)
	}
}

// AddEventHandler adds an event handler to the shared informer using the shared informer's resync
// period.  Events to a single handler are delivered sequentially, but there is no coordination
// between different handlers.
// It returns a registration handle for the handler that can be used to remove
// the handler again.
func (i *streamingInformer) AddEventHandler(handler toolscache.ResourceEventHandler) (toolscache.ResourceEventHandlerRegistration, error) {
	return i.AddEventHandlerWithResyncPeriod(handler, i.resyncPeriod)
}

// AddEventHandlerWithResyncPeriod adds an event handler to the shared informer using the
// specified resync period.  Events to a single handler are delivered sequentially, but there is
// no coordination between different handlers.
// It returns a registration handle for the handler that can be used to remove
// the handler again and an error if the handler cannot be added.
func (i *streamingInformer) AddEventHandlerWithResyncPeriod(handler toolscache.ResourceEventHandler, resyncPeriod time.Duration) (toolscache.ResourceEventHandlerRegistration, error) {
	i.mutex.Lock()
	defer i.mutex.Unlock()

	registration := &eventHandlerRegistration{handler: handler, resyncPeriod: resyncPeriod}

	// TODO: Propagate correctly to ListWatch
	i.eventHandlerRegistrations = append(i.eventHandlerRegistrations, registration)

	if i.hasSynced.Load() {
		// TODO: Snapshot
		for _, obj := range i.objects.store {
			registration.handler.OnAdd(obj, true)
		}
	}

	return registration, nil
}

// RemoveEventHandler removes a formerly added event handler given by
// its registration handle.
// This function is guaranteed to be idempotent, and thread-safe.
func (i *streamingInformer) RemoveEventHandler(handle toolscache.ResourceEventHandlerRegistration) error {
	panic("not implemented")
}

// AddIndexers adds more indexers to this store.  If you call this after you already have data
// in the store, the results are undefined.
func (i *streamingInformer) AddIndexers(indexers toolscache.Indexers) error {
	panic("not implemented")
}

// HasSynced return true if the informers underlying store has synced.
func (i *streamingInformer) HasSynced() bool {
	return i.hasSynced.Load()
}

// WaitForCacheSync waits for all the caches to sync.  Returns false if it could not sync a cache.
func (i *streamingInformer) WaitForCacheSync(ctx context.Context) bool {
	timer := time.NewTicker(50 * time.Millisecond)
	defer timer.Stop()

	for {
		if i.HasSynced() {
			return true
		}

		// Check for timer or ctx.Done()
		select {
		case <-timer.C:
			// Check again
		case <-ctx.Done():
			return false
		}
	}
}

type eventHandlerRegistration struct {
	informer     *streamingInformer
	handler      toolscache.ResourceEventHandler
	resyncPeriod time.Duration
}

// HasSynced reports if both the parent has synced and all pre-sync
// events have been delivered.
func (i *eventHandlerRegistration) HasSynced() bool {
	return i.informer.HasSynced()
}

// Get retrieves an obj for the given object key from the Kubernetes Cluster.
// obj must be a struct pointer so that obj can be updated with the response
// returned by the Server.
func (i *streamingInformer) Get(ctx context.Context, key client.ObjectKey, obj client.Object, opts ...client.GetOption) error {
	if len(opts) != 0 {
		return fmt.Errorf("options not implemented: %v", opts)
	}

	i.objects.mutex.Lock()
	defer i.objects.mutex.Unlock()

	existing, ok := i.objects.store[key]
	if !ok {
		return apierrors.NewNotFound(i.typeInfo.GroupResource(), key.String())
	}

	// TODO: How do we want to copy objects?
	b, err := json.Marshal(existing)
	if err != nil {
		return fmt.Errorf("error copying %T: %w", obj, err)
	}
	if err := json.Unmarshal(b, obj); err != nil {
		return fmt.Errorf("error copying %T: %w", obj, err)
	}
	return nil
}
