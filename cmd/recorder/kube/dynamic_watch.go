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

package kube

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/dynamic"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

// DynamicWatch is a handle to a watch; closing will free resources.
type DynamicWatch interface {
	// Close stops the watch and frees resources.
	Close()
}

// watchRestartDelay is the time between a Watch being dropped and attempting to resume it
const watchRestartDelay = 30 * time.Second

// watchListener is the interface that is implemented by object observers.
type watchListener interface {
	// OnRestartWatch is called when a watch is being restarted.
	// Observers should record whether any objects are dropped from the subsequent OnAdded calls,
	// and treat this as removed.
	OnRestartWatch()

	// OnBookmark is called for each bookmark event.
	// It is called after the initial list, following OnRestartWatch.
	OnBookmark()

	// OnAdded is called whenever an object is added, in response to the ADDED watch event.
	// It is also called for every object after watches are started / restarted.
	OnAdded(obj *unstructured.Unstructured)

	// OnDeleted is called whenever an object is deleted, in response to the DELETED watch event.
	OnDeleted(obj *unstructured.Unstructured)

	// OnModified is called whenever an object is modified, in response to the MODIFIED watch event.
	OnModified(obj *unstructured.Unstructured)
}

// dynamicWatchers manages a set of watches.
type dynamicWatchers struct {
	client dynamic.Interface
}

// NewDynamicWatchers constructs a dynamicWatchers object.
func newDynamicWatchers(client dynamic.Interface) *dynamicWatchers {
	return &dynamicWatchers{
		client: client,
	}
}

// dynamicWatch manages a single watch.
type dynamicWatch struct {
	gvr schema.GroupVersionResource

	client dynamic.Interface

	listener atomic.Pointer[watchListener]

	mutex  sync.Mutex
	closer func()
}

// Close implements Closeable
func (d *dynamicWatch) Close() {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	d.listener.Store(nil)

	if d.closer != nil {
		d.closer()
		d.closer = nil
	}
}

// AddWatch registers a watch for changes to 'trigger' filtered by 'options' to raise an event on 'target'
func (dw *dynamicWatchers) AddWatch(ctx context.Context, gvr schema.GroupVersionResource, listener watchListener) DynamicWatch {
	ctx, closer := context.WithCancel(ctx)

	dkw := &dynamicWatch{
		closer: closer,
		gvr:    gvr,
		client: dw.client,
	}

	dkw.listener.Store(&listener)

	go func() {
		dkw.watchForever(ctx)
	}()

	return dkw
}

func (w *dynamicWatch) watchForever(ctx context.Context) {
	for {
		listener := w.listener.Load()
		if listener == nil {
			return
		}

		w.watchUntilClosed(ctx)

		time.Sleep(watchRestartDelay)
	}
}

// A Watch will be closed when the pod loses connection to the API server.
// If a Watch is opened with no ResourceVersion then we will receive an 'ADDED'
// event for all Watch objects[1]. This will result in "over-notification"
// from this Watch but it will ensure we always Reconcile when needed.
//
// [1] https://github.com/kubernetes/kubernetes/issues/54878#issuecomment-357575276
func (w *dynamicWatch) watchUntilClosed(ctx context.Context) {
	log := log.FromContext(ctx)

	listener := w.listener.Load()
	if listener == nil {
		return
	}
	(*listener).OnRestartWatch()

	resource := w.client.Resource(w.gvr)

	options := metav1.ListOptions{}
	// Though we don't use the resource version, bookmarks help keep TCP connections healthy and for SendInitialEvents.
	options.AllowWatchBookmarks = true

	timeoutSeconds := int64(30 * 60)
	options.TimeoutSeconds = &timeoutSeconds

	// TODO: Once the WatchList feature gate is on by default, we can use SendInitialEvents to get a precise notification of the bookmark.
	// Until then, we rely on a quirk that the bookmark is not sent until after the initial rows are sent (I believe)
	// Even if that's not the case, we "only" end up with some slightly incorrect statistics.
	// sendInitialEvents := true
	// options.SendInitialEvents = &sendInitialEvents
	// options.ResourceVersionMatch = metav1.ResourceVersionMatchNotOlderThan
	// options.ResourceVersion = // unset for consistent read

	// Note that we must send ResourceVersion "0", not empty, to receive bookmark events
	options.ResourceVersion = "0"

	events, err := resource.Watch(ctx, options)
	if err != nil {
		log.WithValues("gvr", w.gvr.String()).Error(err, "failed to start watch")
		return
	}

	log.WithValues("gvr", w.gvr.String()).Info("watch started")

	// Always clean up watchers
	defer events.Stop()

	for clientEvent := range events.ResultChan() {
		listener := w.listener.Load()
		if listener == nil {
			break
		}

		switch clientEvent.Type {
		case watch.Bookmark:
			log.V(2).WithValues("gvr", w.gvr.String()).Info("got watch bookmark")
			(*listener).OnBookmark()
			continue
		case watch.Error:
			log.Error(fmt.Errorf("unexpected error from watch: %v", clientEvent.Object), "error during watch")
			return
		}

		u := clientEvent.Object.(*unstructured.Unstructured)

		switch clientEvent.Type {
		case watch.Deleted:
			(*listener).OnDeleted(u)

		case watch.Added:
			(*listener).OnAdded(u)

		case watch.Modified:
			(*listener).OnModified(u)

		default:
			log.Error(fmt.Errorf("unexpected event type from watch: %v", clientEvent.Type), "error during watch")
			return
		}
	}

	log.WithValues("gvr", w.gvr.String()).Info("watch closed")
}
