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

package deletiondefender

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/util/workqueue"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

func TestPredicate(t *testing.T) {
	for _, tc := range []struct {
		name       string
		deleting   bool
		finalizers []string
		want       bool
	}{
		{name: "live resource with deletion defense", finalizers: []string{k8s.DeletionDefenderFinalizerName}},
		{name: "deleting resource needing finalization", deleting: true, finalizers: []string{k8s.ControllerFinalizerName, k8s.DeletionDefenderFinalizerName}, want: true},
		{name: "deletion defense already finalized", deleting: true, finalizers: []string{k8s.ControllerFinalizerName}},
		{name: "unmanaged resource"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			obj := &metav1.PartialObjectMetadata{
				ObjectMeta: metav1.ObjectMeta{Finalizers: tc.finalizers},
			}
			if tc.deleting {
				deletedAt := metav1.NewTime(time.Unix(1, 0))
				obj.SetDeletionTimestamp(&deletedAt)
			}
			p := Predicate{}
			if got := p.Create(event.CreateEvent{Object: obj}); got != tc.want {
				t.Errorf("Create() = %v, want %v; initial lists must resume pending deletions", got, tc.want)
			}
			if got := p.Update(event.UpdateEvent{ObjectOld: &metav1.PartialObjectMetadata{}, ObjectNew: obj}); got != tc.want {
				t.Errorf("Update() = %v, want %v", got, tc.want)
			}
			if got := p.Generic(event.GenericEvent{Object: obj}); got != tc.want {
				t.Errorf("Generic() = %v, want %v", got, tc.want)
			}
			if p.Delete(event.DeleteEvent{Object: obj}) {
				t.Error("Delete() = true; an object already removed from the API server needs no finalization")
			}
		})
	}
}

func TestPredicateDropsEmptyEvents(t *testing.T) {
	p := Predicate{}
	if p.Create(event.CreateEvent{}) || p.Update(event.UpdateEvent{}) || p.Generic(event.GenericEvent{}) || p.Delete(event.DeleteEvent{}) {
		t.Error("events without object metadata must not be queued")
	}
}

func TestResyncDoesNotQueueLiveResourcesAheadOfDeletion(t *testing.T) {
	ctx := context.Background()
	queue := workqueue.NewTypedRateLimitingQueue(workqueue.DefaultTypedControllerRateLimiter[reconcile.Request]())
	defer queue.ShutDown()
	p := Predicate{}
	h := &handler.EnqueueRequestForObject{}
	for i := 0; i < 33000; i++ {
		obj := &metav1.PartialObjectMetadata{
			ObjectMeta: metav1.ObjectMeta{
				Namespace:  "test-namespace",
				Name:       fmt.Sprintf("live-policy-%d", i),
				Finalizers: []string{k8s.DeletionDefenderFinalizerName, k8s.ControllerFinalizerName},
			},
		}
		create := event.CreateEvent{Object: obj}
		if p.Create(create) {
			h.Create(ctx, create, queue)
		}
		// Informer resyncs deliver updates even when the object is unchanged.
		update := event.UpdateEvent{ObjectOld: obj, ObjectNew: obj.DeepCopy()}
		if p.Update(update) {
			h.Update(ctx, update, queue)
		}
	}
	if got := queue.Len(); got != 0 {
		t.Fatalf("live resources queued %d requests before deletion, want 0", got)
	}

	old := &metav1.PartialObjectMetadata{
		ObjectMeta: metav1.ObjectMeta{
			Namespace:  "test-namespace",
			Name:       "deleting-policy",
			Finalizers: []string{k8s.DeletionDefenderFinalizerName, k8s.ControllerFinalizerName},
		},
	}
	deleting := old.DeepCopy()
	deletedAt := metav1.NewTime(time.Unix(1, 0))
	deleting.SetDeletionTimestamp(&deletedAt)
	update := event.UpdateEvent{ObjectOld: old, ObjectNew: deleting}
	if p.Update(update) {
		h.Update(ctx, update, queue)
	}
	if got := queue.Len(); got != 1 {
		t.Fatalf("pending deletion queued %d requests, want 1", got)
	}
	item, shutdown := queue.Get()
	defer queue.Done(item)
	want := reconcile.Request{NamespacedName: types.NamespacedName{Namespace: old.Namespace, Name: old.Name}}
	if shutdown || item != want {
		t.Fatalf("first queued request = %v (shutdown %v), want %v", item, shutdown, want)
	}
}
