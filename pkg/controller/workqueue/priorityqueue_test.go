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

package workqueue

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/util/workqueue"
	"sigs.k8s.io/controller-runtime/pkg/controller/priorityqueue"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

func newTestRequest(name string) reconcile.Request {
	return reconcile.Request{
		NamespacedName: types.NamespacedName{
			Namespace: "default",
			Name:      name,
		},
	}
}

// TestFIFOOrderingWithinPriorityTiers verifies that items within the same priority tier
// are dequeued strictly in FIFO order.
func TestFIFOOrderingWithinPriorityTiers(t *testing.T) {
	q := New[reconcile.Request]("test-fifo", nil)
	defer q.ShutDown()

	// High priority tier FIFO
	req1 := newTestRequest("high-1")
	req2 := newTestRequest("high-2")
	req3 := newTestRequest("high-3")

	q.Add(req1)
	q.Add(req2)
	q.Add(req3)

	for _, expected := range []reconcile.Request{req1, req2, req3} {
		item, priority, shutdown := q.GetWithPriority()
		if shutdown {
			t.Fatalf("expected item %v, got shutdown", expected)
		}
		if item != expected {
			t.Errorf("expected item %v, got %v", expected, item)
		}
		if priority != PriorityDefault {
			t.Errorf("expected priority %v, got %v", PriorityDefault, priority)
		}
		q.Done(item)
	}

	// Low priority tier FIFO (using AddWithOpts with PriorityLow)
	low1 := newTestRequest("low-1")
	low2 := newTestRequest("low-2")
	low3 := newTestRequest("low-3")

	q.AddWithOpts(priorityqueue.AddOpts{Priority: PriorityLow}, low1, low2, low3)

	for _, expected := range []reconcile.Request{low1, low2, low3} {
		item, priority, shutdown := q.GetWithPriority()
		if shutdown {
			t.Fatalf("expected item %v, got shutdown", expected)
		}
		if item != expected {
			t.Errorf("expected item %v, got %v", expected, item)
		}
		if priority != PriorityLow {
			t.Errorf("expected priority %v, got %v", PriorityLow, priority)
		}
		q.Done(item)
	}
}

// TestPriorityOrdering verifies that high priority items are dequeued before low priority items.
func TestPriorityOrdering(t *testing.T) {
	q := New[reconcile.Request]("test-priority", nil)
	defer q.ShutDown()

	low := newTestRequest("low")
	high := newTestRequest("high")

	// Add low first, then high
	q.AddWithOpts(priorityqueue.AddOpts{Priority: PriorityLow}, low)
	q.Add(high)

	item1, pri1, _ := q.GetWithPriority()
	q.Done(item1)
	if item1 != high || pri1 != PriorityDefault {
		t.Fatalf("expected high item first (%v, pri %d), got %v, pri %d", high, PriorityDefault, item1, pri1)
	}

	item2, pri2, _ := q.GetWithPriority()
	q.Done(item2)
	if item2 != low || pri2 != PriorityLow {
		t.Fatalf("expected low item second (%v, pri %d), got %v, pri %d", low, PriorityLow, item2, pri2)
	}
}

// TestPromotionOnImmediateUpdate verifies that when an item pending in RequeueAfter (delayed)
// receives an immediate Add (watch event), it is immediately promoted to high priority and made ready.
func TestPromotionOnImmediateUpdate(t *testing.T) {
	q := New[reconcile.Request]("test-promotion", nil)
	defer q.ShutDown()

	req := newTestRequest("pending-resync")

	// 1. Resource scheduled for delayed drift resync (1 hour in the future)
	q.AddAfter(req, time.Hour)

	// In the delayed state, queue length is 0 (not ready yet)
	if q.Len() != 0 {
		t.Fatalf("expected ready queue length to be 0 for delayed item, got %d", q.Len())
	}

	// 2. User performs an update or watch event fires -> immediate Add
	q.Add(req)

	if q.Len() != 1 {
		t.Fatalf("expected queue length to be 1 immediately after promotion, got %d", q.Len())
	}

	// 3. Dequeue must immediately return the promoted item with PriorityDefault without waiting 1 hour
	doneCh := make(chan struct{})
	go func() {
		defer close(doneCh)
		item, priority, shutdown := q.GetWithPriority()
		if shutdown {
			t.Errorf("unexpected shutdown")
			return
		}
		if item != req {
			t.Errorf("expected promoted item %v, got %v", req, item)
		}
		if priority != PriorityDefault {
			t.Errorf("expected promoted priority %v, got %v", PriorityDefault, priority)
		}
		q.Done(item)
	}()

	select {
	case <-doneCh:
		// Succeeded immediately
	case <-time.After(2 * time.Second):
		t.Fatalf("timed out waiting for promoted item to dequeue")
	}
}

// TestPromotionWhenAlreadyReadyLowPriority verifies that an item sitting in the low-priority
// ready queue gets promoted to high priority when an immediate Add arrives.
func TestPromotionWhenAlreadyReadyLowPriority(t *testing.T) {
	q := New[reconcile.Request]("test-promotion-ready", nil)
	defer q.ShutDown()

	reqLow := newTestRequest("item-low")
	reqHighOther := newTestRequest("item-high-other")

	// Add item at low priority
	q.AddWithOpts(priorityqueue.AddOpts{Priority: PriorityLow}, reqLow)
	// Add another high priority item
	q.Add(reqHighOther)

	// Now promote reqLow via immediate Add
	q.Add(reqLow)

	// Now both are high priority. Since reqHighOther was added to high first, it pops first,
	// then reqLow pops with PriorityDefault.
	item1, pri1, _ := q.GetWithPriority()
	q.Done(item1)
	if item1 != reqHighOther || pri1 != PriorityDefault {
		t.Fatalf("expected item1 to be %v (pri %d), got %v (pri %d)", reqHighOther, PriorityDefault, item1, pri1)
	}

	item2, pri2, _ := q.GetWithPriority()
	q.Done(item2)
	if item2 != reqLow || pri2 != PriorityDefault {
		t.Fatalf("expected promoted item2 to be %v with pri %d, got %v with pri %d", reqLow, PriorityDefault, item2, pri2)
	}
}

// TestBoundedStarvationCeiling verifies that under a continuous stream of high-priority events,
// a ready low-priority item is guaranteed to drain within StarvationCeiling (8) high-priority dequeues.
func TestBoundedStarvationCeiling(t *testing.T) {
	q := New[reconcile.Request]("test-starvation", nil)
	defer q.ShutDown()

	lowReq := newTestRequest("background-drift-resync")
	q.AddWithOpts(priorityqueue.AddOpts{Priority: PriorityLow}, lowReq)

	// Continuously enqueue 20 high-priority requests
	for i := 0; i < 20; i++ {
		highReq := newTestRequest(fmt.Sprintf("user-mutation-%d", i))
		q.Add(highReq)
	}

	lowServedAt := -1
	for i := 0; i < 21; i++ {
		item, priority, shutdown := q.GetWithPriority()
		if shutdown {
			t.Fatalf("unexpected shutdown at iteration %d", i)
		}
		q.Done(item)

		if item == lowReq {
			lowServedAt = i
			if priority != PriorityLow {
				t.Errorf("expected lowReq to have priority %d, got %d", PriorityLow, priority)
			}
			break
		}
	}

	if lowServedAt == -1 {
		t.Fatalf("low priority item was starved and never served")
	}

	// With StarvationCeiling = 8, 8 high priority items are served first (indices 0..7),
	// and the 9th item (index 8) MUST be the low-priority item.
	if lowServedAt != StarvationCeiling {
		t.Fatalf("expected low priority item to be served at index %d, but was served at %d", StarvationCeiling, lowServedAt)
	}
}

// TestRateLimitedAndErrorRetries verifies rate-limited retry integration and that retries
// are scheduled at PriorityDefault.
func TestRateLimitedAndErrorRetries(t *testing.T) {
	rateLimiter := workqueue.NewTypedItemExponentialFailureRateLimiter[reconcile.Request](10*time.Millisecond, 100*time.Millisecond)
	q := New[reconcile.Request]("test-retry", rateLimiter)
	defer q.ShutDown()

	req := newTestRequest("failing-resource")

	// Initially num requeues is 0
	if q.NumRequeues(req) != 0 {
		t.Fatalf("expected 0 requeues initially, got %d", q.NumRequeues(req))
	}

	// Add rate limited
	q.AddRateLimited(req)

	if q.NumRequeues(req) != 1 {
		t.Fatalf("expected 1 requeue after AddRateLimited, got %d", q.NumRequeues(req))
	}

	// Should be dequeued with PriorityDefault once the rate limit timer elapses
	doneCh := make(chan struct{})
	go func() {
		defer close(doneCh)
		item, priority, shutdown := q.GetWithPriority()
		if shutdown {
			t.Errorf("unexpected shutdown")
			return
		}
		if item != req {
			t.Errorf("expected %v, got %v", req, item)
		}
		if priority != PriorityDefault {
			t.Errorf("expected PriorityDefault %d, got %d", PriorityDefault, priority)
		}
		q.Done(item)
	}()

	select {
	case <-doneCh:
		// OK
	case <-time.After(2 * time.Second):
		t.Fatalf("timed out waiting for rate limited item to dequeue")
	}

	// Forget clears the rate limiter
	q.Forget(req)
	if q.NumRequeues(req) != 0 {
		t.Fatalf("expected 0 requeues after Forget, got %d", q.NumRequeues(req))
	}
}

// TestShutdownDrain verifies queue shutdown behavior.
func TestShutdownDrain(t *testing.T) {
	q := New[reconcile.Request]("test-shutdown", nil)

	req1 := newTestRequest("req-1")
	req2 := newTestRequest("req-2")

	q.Add(req1)
	q.Add(req2)

	q.ShutDown()

	if !q.ShuttingDown() {
		t.Fatalf("expected ShuttingDown() to be true")
	}

	// Ready items should drain
	item1, shutdown1 := q.Get()
	if shutdown1 || item1 != req1 {
		t.Fatalf("expected req1 before shutdown drain, got %v, shutdown %v", item1, shutdown1)
	}
	q.Done(item1)

	item2, shutdown2 := q.Get()
	if shutdown2 || item2 != req2 {
		t.Fatalf("expected req2 before shutdown drain, got %v, shutdown %v", item2, shutdown2)
	}
	q.Done(item2)

	// Once empty, Get returns shutdown = true
	_, shutdown3 := q.Get()
	if !shutdown3 {
		t.Fatalf("expected shutdown to be true when draining an empty shut down queue")
	}
}

// TestConcurrentAccess verifies thread-safety under concurrent adds and gets.
func TestConcurrentAccess(t *testing.T) {
	q := New[reconcile.Request]("test-concurrency", nil)
	defer q.ShutDown()

	const numWorkers = 5
	const numItems = 200

	var wg sync.WaitGroup
	wg.Add(numWorkers)

	consumed := make(map[reconcile.Request]int)
	var mu sync.Mutex

	for w := 0; w < numWorkers; w++ {
		go func() {
			defer wg.Done()
			for {
				item, shutdown := q.Get()
				if shutdown {
					return
				}
				mu.Lock()
				consumed[item]++
				mu.Unlock()
				q.Done(item)
			}
		}()
	}

	// Produce items concurrently
	var prodWg sync.WaitGroup
	prodWg.Add(2)

	go func() {
		defer prodWg.Done()
		for i := 0; i < numItems; i++ {
			q.Add(newTestRequest(fmt.Sprintf("high-%d", i)))
		}
	}()

	go func() {
		defer prodWg.Done()
		for i := 0; i < numItems; i++ {
			q.AddWithOpts(priorityqueue.AddOpts{Priority: PriorityLow}, newTestRequest(fmt.Sprintf("low-%d", i)))
		}
	}()

	prodWg.Wait()

	// Wait a bit for processing to complete, then shut down
	time.Sleep(100 * time.Millisecond)
	q.ShutDown()
	wg.Wait()

	mu.Lock()
	defer mu.Unlock()
	if len(consumed) != numItems*2 {
		t.Fatalf("expected %d items consumed, got %d", numItems*2, len(consumed))
	}
}
