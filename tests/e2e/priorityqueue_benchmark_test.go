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

package e2e

import (
	"fmt"
	"sort"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/util/workqueue"
	"sigs.k8s.io/controller-runtime/pkg/controller/priorityqueue"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	kccwq "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/workqueue"
)

// TestPriorityQueueBacklogBenchmark simulates a production workload with 1,000 resources
// generating background drift resyncs, while user spec mutations arrive intermittently.
// It measures p50 and p99 actuation latency from mutation enqueue to completion, comparing
// the standard FIFO queue against the starvation-bounded PriorityQueue.
func TestPriorityQueueBacklogBenchmark(t *testing.T) {
	const (
		numSimulatedResources  = 1000
		numMutations           = 50
		simulatedReconcileWork = 50 * time.Microsecond
		numWorkers             = 10
	)

	runBenchmark := func(t *testing.T, usePQ bool) (p50 time.Duration, p99 time.Duration) {
		var queue workqueue.TypedRateLimitingInterface[reconcile.Request]
		var pq priorityqueue.PriorityQueue[reconcile.Request]
		if usePQ {
			p := kccwq.New[reconcile.Request]("bench-pq", nil)
			queue = p
			pq = p
		} else {
			queue = workqueue.NewTypedRateLimitingQueueWithConfig[reconcile.Request](
				workqueue.DefaultTypedControllerRateLimiter[reconcile.Request](),
				workqueue.TypedRateLimitingQueueConfig[reconcile.Request]{Name: "bench-fifo"},
			)
		}
		defer queue.ShutDown()

		// Enqueue 1,000 background drift resyncs
		for i := 0; i < numSimulatedResources; i++ {
			req := reconcile.Request{
				NamespacedName: types.NamespacedName{
					Namespace: "bench-ns",
					Name:      fmt.Sprintf("resync-%d", i),
				},
			}
			if usePQ && pq != nil {
				// RequeueAfter drift resyncs enter at PriorityLow
				pq.AddWithOpts(priorityqueue.AddOpts{Priority: kccwq.PriorityLow}, req)
			} else {
				queue.Add(req)
			}
		}

		mutationStartTimes := make(map[reconcile.Request]time.Time)
		mutationLatencies := make([]time.Duration, 0, numMutations)
		var latMu sync.Mutex

		var activeWorkers atomic.Int32
		stopWorkers := make(chan struct{})
		var wg sync.WaitGroup

		// Start worker pool simulating controllers executing reconciles
		for w := 0; w < numWorkers; w++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for {
					select {
					case <-stopWorkers:
						return
					default:
					}

					item, shutdown := queue.Get()
					if shutdown {
						return
					}
					activeWorkers.Add(1)

					// Simulate lightweight reconciliation work
					time.Sleep(simulatedReconcileWork)

					latMu.Lock()
					if start, ok := mutationStartTimes[item]; ok {
						mutationLatencies = append(mutationLatencies, time.Since(start))
					}
					latMu.Unlock()

					queue.Done(item)
					activeWorkers.Add(-1)
				}
			}()
		}

		// Inject mutations with spacing to simulate user actions
		for m := 0; m < numMutations; m++ {
			req := reconcile.Request{
				NamespacedName: types.NamespacedName{
					Namespace: "bench-ns",
					Name:      fmt.Sprintf("user-mutation-%d", m),
				},
			}
			latMu.Lock()
			mutationStartTimes[req] = time.Now()
			latMu.Unlock()

			if usePQ && pq != nil {
				// Watch event arrives at PriorityDefault
				pq.Add(req)
			} else {
				queue.Add(req)
			}
			time.Sleep(100 * time.Microsecond)
		}

		// Wait until all user mutations have been processed
		deadline := time.After(10 * time.Second)
		for {
			latMu.Lock()
			done := len(mutationLatencies) >= numMutations
			latMu.Unlock()
			if done {
				break
			}
			select {
			case <-deadline:
				t.Fatalf("timed out waiting for mutations to complete (completed: %d/%d)", len(mutationLatencies), numMutations)
			default:
				time.Sleep(1 * time.Millisecond)
			}
		}

		close(stopWorkers)
		queue.ShutDown()
		wg.Wait()

		sort.Slice(mutationLatencies, func(i, j int) bool {
			return mutationLatencies[i] < mutationLatencies[j]
		})

		p50 = mutationLatencies[len(mutationLatencies)*50/100]
		p99 = mutationLatencies[len(mutationLatencies)*99/100]
		return p50, p99
	}

	fifoP50, fifoP99 := runBenchmark(t, false)
	pqP50, pqP99 := runBenchmark(t, true)

	t.Logf("FIFO Queue (1000 resync backlog): p50 = %v, p99 = %v", fifoP50, fifoP99)
	t.Logf("Priority Queue (1000 resync backlog): p50 = %v, p99 = %v", pqP50, pqP99)

	// In the PriorityQueue, user mutations must not sit behind the 1,000 backlogged drift resyncs.
	// Latency should be significantly faster than FIFO.
	if pqP99 >= fifoP99 {
		t.Errorf("expected PriorityQueue p99 (%v) to be lower than FIFO p99 (%v)", pqP99, fifoP99)
	}
}
