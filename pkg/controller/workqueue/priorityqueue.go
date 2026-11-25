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
	"container/heap"
	"sync"
	"sync/atomic"
	"time"

	"k8s.io/client-go/util/workqueue"
	"sigs.k8s.io/controller-runtime/pkg/controller/priorityqueue"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

const (
	// PriorityDefault is the high priority tier (0) for watch events, dependency wakeups,
	// and rate-limited error retries.
	PriorityDefault = 0

	// PriorityLow is the low priority tier (-100) for periodic RequeueAfter drift resyncs.
	PriorityLow = -100

	// StarvationCeiling defines the maximum consecutive high-priority items served before
	// forcing the dequeue of at least one ready low-priority item, if available.
	StarvationCeiling = 8
)

// PriorityWorkQueue wraps controller-runtime's PriorityQueue interface and
// client-go's TypedRateLimitingInterface to provide bounded-starvation priority scheduling.
type PriorityWorkQueue[T comparable] interface {
	priorityqueue.PriorityQueue[T]
}

type queueItem[T comparable] struct {
	key          T
	priority     int
	readyAt      time.Time
	addedCounter uint64
	index        int // index in the delayed heap
}

// delayedHeap implements heap.Interface for items waiting on a timer.
type delayedHeap[T comparable] []*queueItem[T]

func (h delayedHeap[T]) Len() int           { return len(h) }
func (h delayedHeap[T]) Less(i, j int) bool { return h[i].readyAt.Before(h[j].readyAt) }
func (h delayedHeap[T]) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].index = i
	h[j].index = j
}
func (h *delayedHeap[T]) Push(x any) {
	n := len(*h)
	it := x.(*queueItem[T])
	it.index = n
	*h = append(*h, it)
}
func (h *delayedHeap[T]) Pop() any {
	old := *h
	n := len(old)
	it := old[n-1]
	old[n-1] = nil
	it.index = -1
	*h = old[0 : n-1]
	return it
}

// fifoQueue maintains FIFO ordering for ready items within a priority tier.
type fifoQueue[T comparable] struct {
	items []T
	set   map[T]struct{}
}

func newFIFOQueue[T comparable]() *fifoQueue[T] {
	return &fifoQueue[T]{
		items: make([]T, 0),
		set:   make(map[T]struct{}),
	}
}

func (q *fifoQueue[T]) add(key T) bool {
	if _, exists := q.set[key]; exists {
		return false
	}
	q.set[key] = struct{}{}
	q.items = append(q.items, key)
	return true
}

func (q *fifoQueue[T]) remove(key T) bool {
	if _, exists := q.set[key]; !exists {
		return false
	}
	delete(q.set, key)
	for i, item := range q.items {
		if item == key {
			q.items = append(q.items[:i], q.items[i+1:]...)
			break
		}
	}
	return true
}

func (q *fifoQueue[T]) pop() (T, bool) {
	var zero T
	if len(q.items) == 0 {
		return zero, false
	}
	key := q.items[0]
	q.items = q.items[1:]
	delete(q.set, key)
	return key, true
}

func (q *fifoQueue[T]) has(key T) bool {
	_, ok := q.set[key]
	return ok
}

func (q *fifoQueue[T]) len() int {
	return len(q.items)
}

type boundedPriorityQueue[T comparable] struct {
	name        string
	rateLimiter workqueue.TypedRateLimiter[T]

	lock sync.Mutex
	cond *sync.Cond

	// Ready queues per priority tier
	highReady *fifoQueue[T]
	lowReady  *fifoQueue[T]

	// Delayed heap for items with After > 0
	delayed delayedHeap[T]

	// Tracking maps
	// items maps key to currently queued item metadata
	items map[T]*queueItem[T]
	// processing tracks keys currently being processed (between Get and Done)
	processing map[T]struct{}
	// dirty tracks keys requested while being processed, mapping to the highest pending priority
	dirty map[T]int

	starvationCeiling    int
	consecutiveHighCount int
	addedCounter         uint64

	shuttingDown    atomic.Bool
	stopCh          chan struct{}
	delayedUpdateCh chan struct{}

	metrics queueMetrics[T]

	// Clock hooks for testing
	now func() time.Time
}

// Option configures boundedPriorityQueue options.
type Option[T comparable] func(*options[T])

type options[T comparable] struct {
	metricsProvider   workqueue.MetricsProvider
	starvationCeiling int
}

// WithMetricsProvider specifies a custom MetricsProvider.
func WithMetricsProvider[T comparable](mp workqueue.MetricsProvider) Option[T] {
	return func(o *options[T]) {
		o.metricsProvider = mp
	}
}

// WithStarvationCeiling sets the starvation ceiling.
func WithStarvationCeiling[T comparable](ceiling int) Option[T] {
	return func(o *options[T]) {
		o.starvationCeiling = ceiling
	}
}

// New constructs a new starvation-bounded priority work queue.
func New[T comparable](name string, rateLimiter workqueue.TypedRateLimiter[T], opts ...Option[T]) PriorityWorkQueue[T] {
	if rateLimiter == nil {
		rateLimiter = workqueue.DefaultTypedControllerRateLimiter[T]()
	}

	optConfig := &options[T]{
		starvationCeiling: StarvationCeiling,
	}
	for _, opt := range opts {
		opt(optConfig)
	}

	q := &boundedPriorityQueue[T]{
		name:              name,
		rateLimiter:       rateLimiter,
		highReady:         newFIFOQueue[T](),
		lowReady:          newFIFOQueue[T](),
		delayed:           make(delayedHeap[T], 0),
		items:             make(map[T]*queueItem[T]),
		processing:        make(map[T]struct{}),
		dirty:             make(map[T]int),
		starvationCeiling: optConfig.starvationCeiling,
		stopCh:            make(chan struct{}),
		delayedUpdateCh:   make(chan struct{}, 1),
		now:               time.Now,
	}
	q.cond = sync.NewCond(&q.lock)
	q.metrics = newQueueMetrics[T](optConfig.metricsProvider, name, q.now)

	heap.Init(&q.delayed)
	go q.delayedLoop()
	if _, ok := q.metrics.(noMetrics[T]); !ok {
		go q.updateUnfinishedWorkLoop()
	}

	return q
}

// NewQueueFunc returns a NewQueue constructor compatible with controller.Options.NewQueue.
func NewQueueFunc(rateLimiter workqueue.TypedRateLimiter[reconcile.Request], opts ...Option[reconcile.Request]) func(controllerName string, rl workqueue.TypedRateLimiter[reconcile.Request]) workqueue.TypedRateLimitingInterface[reconcile.Request] {
	return func(controllerName string, rl workqueue.TypedRateLimiter[reconcile.Request]) workqueue.TypedRateLimitingInterface[reconcile.Request] {
		limiter := rateLimiter
		if limiter == nil {
			limiter = rl
		}
		return New[reconcile.Request](controllerName, limiter, opts...)
	}
}

func (q *boundedPriorityQueue[T]) notifyDelayedUpdate() {
	select {
	case q.delayedUpdateCh <- struct{}{}:
	default:
	}
}

func (q *boundedPriorityQueue[T]) enqueueReadyLocked(key T, priority int) {
	if _, isProcessing := q.processing[key]; isProcessing {
		// Key is currently being processed by a worker.
		// Mark it dirty so it gets enqueued when Done() is invoked.
		currPri, exists := q.dirty[key]
		if !exists || priority > currPri {
			q.dirty[key] = priority
		}
		return
	}

	it := &queueItem[T]{
		key:          key,
		priority:     priority,
		addedCounter: q.addedCounter,
		index:        -1,
	}
	q.items[key] = it

	var added bool
	if priority >= PriorityDefault {
		added = q.highReady.add(key)
	} else {
		added = q.lowReady.add(key)
	}
	if added {
		q.metrics.add(key)
	}
	q.cond.Broadcast()
}

func (q *boundedPriorityQueue[T]) Add(item T) {
	q.AddWithOpts(priorityqueue.AddOpts{Priority: PriorityDefault}, item)
}

func (q *boundedPriorityQueue[T]) AddAfter(item T, after time.Duration) {
	q.AddWithOpts(priorityqueue.AddOpts{After: after, Priority: PriorityLow}, item)
}

func (q *boundedPriorityQueue[T]) AddRateLimited(item T) {
	q.AddWithOpts(priorityqueue.AddOpts{RateLimited: true, Priority: PriorityDefault}, item)
}

func (q *boundedPriorityQueue[T]) AddWithOpts(o priorityqueue.AddOpts, items ...T) {
	if q.shuttingDown.Load() {
		return
	}

	q.lock.Lock()
	defer q.lock.Unlock()

	if q.shuttingDown.Load() {
		return
	}

	for _, key := range items {
		effectivePriority := o.Priority
		after := o.After

		if o.RateLimited {
			rlAfter := q.rateLimiter.When(key)
			if after == 0 || rlAfter < after {
				after = rlAfter
			}
			// Rate-limited error retries always route to PriorityDefault
			effectivePriority = PriorityDefault
			q.metrics.retry()
		} else if after > 0 {
			// Periodic drift resyncs (after > 0) route to PriorityLow unless explicitly set higher
			if o.Priority <= PriorityDefault {
				effectivePriority = PriorityLow
			}
			q.metrics.retry()
		}

		// Handle keys currently being reconciled by a worker (processing set)
		if _, isProcessing := q.processing[key]; isProcessing {
			if after == 0 {
				// Immediate arrival while processing:
				// If a delayed entry existed, cancel it because the immediate event supersedes it.
				if existing, exists := q.items[key]; exists && existing.index >= 0 {
					heap.Remove(&q.delayed, existing.index)
					existing.index = -1
					delete(q.items, key)
				}
				currPri, wasDirty := q.dirty[key]
				if !wasDirty || effectivePriority > currPri {
					q.dirty[key] = effectivePriority
				}
			} else {
				// Delayed arrival while processing:
				// If already dirty from an immediate event, immediate event takes precedence.
				if _, isDirty := q.dirty[key]; isDirty {
					continue
				}
				existing, exists := q.items[key]
				if exists && existing.index >= 0 {
					newReadyAt := q.now().Add(after)
					if newReadyAt.Before(existing.readyAt) {
						existing.readyAt = newReadyAt
						heap.Fix(&q.delayed, existing.index)
						q.notifyDelayedUpdate()
					}
					if effectivePriority > existing.priority {
						existing.priority = effectivePriority
					}
				} else if !exists {
					q.addedCounter++
					it := &queueItem[T]{
						key:          key,
						priority:     effectivePriority,
						readyAt:      q.now().Add(after),
						addedCounter: q.addedCounter,
						index:        -1,
					}
					q.items[key] = it
					heap.Push(&q.delayed, it)
					q.notifyDelayedUpdate()
				}
			}
			continue
		}

		existing, exists := q.items[key]
		if exists {
			// The item is already queued (either ready or delayed).
			if after == 0 {
				// Immediate arrival promotions:
				if existing.index >= 0 {
					heap.Remove(&q.delayed, existing.index)
					existing.index = -1
					if effectivePriority > existing.priority {
						existing.priority = effectivePriority
					} else if existing.priority < PriorityDefault {
						existing.priority = PriorityDefault
					}
					existing.readyAt = time.Time{}
					var added bool
					if existing.priority >= PriorityDefault {
						added = q.highReady.add(key)
					} else {
						added = q.lowReady.add(key)
					}
					if added {
						q.metrics.add(key)
					}
					q.cond.Broadcast()
				} else {
					// Already in ready queue
					if existing.priority == PriorityLow && effectivePriority >= PriorityDefault {
						q.lowReady.remove(key)
						q.highReady.add(key)
					}
					if effectivePriority > existing.priority {
						existing.priority = effectivePriority
					}
				}
			} else {
				// Incoming is delayed.
				if existing.index < 0 {
					// Already ready in ready queue; keep it ready.
					if effectivePriority > existing.priority {
						if existing.priority == PriorityLow && effectivePriority >= PriorityDefault {
							q.lowReady.remove(key)
							q.highReady.add(key)
						}
						existing.priority = effectivePriority
					}
				} else {
					// Both existing and incoming are delayed.
					newReadyAt := q.now().Add(after)
					if newReadyAt.Before(existing.readyAt) {
						existing.readyAt = newReadyAt
						heap.Fix(&q.delayed, existing.index)
						q.notifyDelayedUpdate()
					}
					if effectivePriority > existing.priority {
						existing.priority = effectivePriority
					}
				}
			}
			continue
		}

		// New item
		q.addedCounter++
		if after > 0 {
			it := &queueItem[T]{
				key:          key,
				priority:     effectivePriority,
				readyAt:      q.now().Add(after),
				addedCounter: q.addedCounter,
				index:        -1,
			}
			q.items[key] = it
			heap.Push(&q.delayed, it)
			q.notifyDelayedUpdate()
		} else {
			q.enqueueReadyLocked(key, effectivePriority)
		}
	}
}

func (q *boundedPriorityQueue[T]) Get() (item T, shutdown bool) {
	item, _, shutdown = q.GetWithPriority()
	return item, shutdown
}

func (q *boundedPriorityQueue[T]) GetWithPriority() (item T, priority int, shutdown bool) {
	q.lock.Lock()
	defer q.lock.Unlock()

	for {
		if q.shuttingDown.Load() && q.highReady.len() == 0 && q.lowReady.len() == 0 {
			var zero T
			return zero, 0, true
		}

		highLen := q.highReady.len()
		lowLen := q.lowReady.len()

		if highLen > 0 || lowLen > 0 {
			serveLow := false
			if lowLen > 0 {
				if highLen == 0 || q.consecutiveHighCount >= q.starvationCeiling {
					serveLow = true
				}
			}

			var key T
			var found bool
			if serveLow {
				key, found = q.lowReady.pop()
				q.consecutiveHighCount = 0
			} else {
				key, found = q.highReady.pop()
				q.consecutiveHighCount++
			}

			if found {
				it := q.items[key]
				pri := PriorityDefault
				if it != nil {
					pri = it.priority
				}
				delete(q.items, key)
				q.processing[key] = struct{}{}
				delete(q.dirty, key)
				q.metrics.get(key)
				return key, pri, false
			}
		}

		if q.shuttingDown.Load() {
			var zero T
			return zero, 0, true
		}

		q.cond.Wait()
	}
}

func (q *boundedPriorityQueue[T]) Done(item T) {
	q.lock.Lock()
	defer q.lock.Unlock()

	delete(q.processing, item)
	q.metrics.done(item)

	if pri, isDirty := q.dirty[item]; isDirty {
		delete(q.dirty, item)
		q.addedCounter++
		q.enqueueReadyLocked(item, pri)
	}
}

func (q *boundedPriorityQueue[T]) ShutDown() {
	if q.shuttingDown.Swap(true) {
		return
	}
	close(q.stopCh)
	q.lock.Lock()
	q.cond.Broadcast()
	q.lock.Unlock()
	q.notifyDelayedUpdate()
}

func (q *boundedPriorityQueue[T]) ShutDownWithDrain() {
	q.ShutDown()
}

func (q *boundedPriorityQueue[T]) ShuttingDown() bool {
	return q.shuttingDown.Load()
}

func (q *boundedPriorityQueue[T]) Forget(item T) {
	q.rateLimiter.Forget(item)
}

func (q *boundedPriorityQueue[T]) NumRequeues(item T) int {
	return q.rateLimiter.NumRequeues(item)
}

func (q *boundedPriorityQueue[T]) Len() int {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.highReady.len() + q.lowReady.len()
}

func (q *boundedPriorityQueue[T]) updateUnfinishedWorkLoop() {
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()
	for {
		select {
		case <-q.stopCh:
			return
		case <-ticker.C:
			q.metrics.updateUnfinishedWork()
		}
	}
}

func (q *boundedPriorityQueue[T]) delayedLoop() {
	timer := time.NewTimer(time.Hour)
	defer timer.Stop()

	for {
		q.lock.Lock()
		if q.shuttingDown.Load() && len(q.delayed) == 0 {
			q.lock.Unlock()
			return
		}

		now := q.now()
		for len(q.delayed) > 0 {
			next := q.delayed[0]
			if next.readyAt.After(now) {
				break
			}
			it := heap.Pop(&q.delayed).(*queueItem[T])
			it.index = -1
			q.enqueueReadyLocked(it.key, it.priority)
		}

		var waitDuration time.Duration
		if len(q.delayed) > 0 {
			waitDuration = q.delayed[0].readyAt.Sub(now)
			if waitDuration < 0 {
				waitDuration = 0
			}
		} else {
			waitDuration = time.Hour
		}
		q.lock.Unlock()

		if !timer.Stop() {
			select {
			case <-timer.C:
			default:
			}
		}
		timer.Reset(waitDuration)

		select {
		case <-q.stopCh:
			return
		case <-q.delayedUpdateCh:
		case <-timer.C:
		}
	}
}
