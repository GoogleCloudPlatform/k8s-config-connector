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
	"sync"
	"sync/atomic"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
)

// KubeView listens to a watch, and stores the value of a function on each object.
type KubeView[V any] struct {
	// mapFn holds the function we evaluate for each object
	mapFn func(*unstructured.Unstructured) V

	// mutex protects against concurrent access for the subsequent fields.
	mutex sync.Mutex

	// values holds the value of the function for each object.
	values map[types.NamespacedName]V

	// seenSinceLastRestart is populated on a new watch until we receive the first bookmark.
	// This allows us to track object deletions (in between watches)
	seenSinceLastRestart map[types.NamespacedName]bool

	// syncedOnce is set once we have observed the full set of objects at least once,
	// because we have seen a bookmark event that indicates all the objects have been sent.
	syncedOnce atomic.Bool

	// watch is used to cleanup the watch when we are done.
	watch DynamicWatch
}

func (m *KubeView[V]) Close() {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if m.watch != nil {
		m.watch.Close()
		m.watch = nil
	}
}

var _ watchListener = &kubeViewListener[string]{}

// WatchKube constructs a KubeView.
func WatchKube[V any](ctx context.Context, kube *Target, gvr schema.GroupVersionResource, fn func(*unstructured.Unstructured) V) *KubeView[V] {
	k := &KubeView[V]{
		mapFn:  fn,
		values: make(map[types.NamespacedName]V),
	}

	listener := &kubeViewListener[V]{
		KubeView: k,
	}

	watchHandle := kube.dynamicWatchers.AddWatch(ctx, gvr, listener)
	k.watch = watchHandle

	return k
}

// Snapshot returns a deep-copy of the map
func (m *KubeView[V]) Snapshot() map[types.NamespacedName]V {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	out := make(map[types.NamespacedName]V, len(m.values))
	for k, v := range m.values {
		out[k] = v
	}
	return out
}

// HasSyncedOnce is true if we have seen all the objects at least once.
// This allows us to wait for the initial "list" to complete, though
// that list may be implemented via a watch.
func (m *KubeView[V]) HasSyncedOnce() bool {
	return m.syncedOnce.Load()
}

// kubeViewListener implements the watchListener, populating KubeView.
// This avoids polluting the public interface of KubeView.
type kubeViewListener[V any] struct {
	*KubeView[V]
}

// OnRestartWatch implements WatchListener.
func (l *kubeViewListener[V]) OnRestartWatch() {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	l.seenSinceLastRestart = make(map[types.NamespacedName]bool)
}

// OnBookmark implements WatchListener.
func (l *kubeViewListener[V]) OnBookmark() {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	// kube-apiserver sends a bookmark once the initial stream is done.
	if l.seenSinceLastRestart != nil {
		for id := range l.values {
			if !l.seenSinceLastRestart[id] {
				delete(l.values, id)
			}
		}

		l.seenSinceLastRestart = nil
	}

	l.syncedOnce.Store(true)
}

// OnAdded implements WatchListener.
func (l *kubeViewListener[V]) OnAdded(obj *unstructured.Unstructured) {
	l.setValue(obj)
}

// OnModified implements WatchListener.
func (l *kubeViewListener[V]) OnModified(obj *unstructured.Unstructured) {
	l.setValue(obj)
}

func (l *kubeViewListener[V]) setValue(obj *unstructured.Unstructured) {
	id := types.NamespacedName{
		Namespace: obj.GetNamespace(),
		Name:      obj.GetName(),
	}
	newV := l.mapFn(obj)

	l.mutex.Lock()
	defer l.mutex.Unlock()

	l.values[id] = newV
	if l.seenSinceLastRestart != nil {
		l.seenSinceLastRestart[id] = true
	}
}

// OnDeleted implements WatchListener.
func (l *kubeViewListener[V]) OnDeleted(obj *unstructured.Unstructured) {
	id := types.NamespacedName{
		Namespace: obj.GetNamespace(),
		Name:      obj.GetName(),
	}

	l.mutex.Lock()
	defer l.mutex.Unlock()

	delete(l.values, id)
}
