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

package ratelimiter

import (
	"time"

	"golang.org/x/time/rate"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/util/flowcontrol"
	"k8s.io/client-go/util/workqueue"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

type RateLimiter = workqueue.TypedRateLimiter[reconcile.Request]

func NewRateLimiter() RateLimiter {
	// This is based on workqueue.DefaultControllerRateLimiter, but with different parameters better suited to KRM reconciliation.
	// Context is in b/188203307

	// We have both overall and per-item rate limiting.
	// The overall is a token bucket and the per-item is exponential
	// The per-item rate limiter initially retries much more slowly (2 seconds vs 2 milliseconds),
	// and a much faster ultimate limit (120 seconds instead of 1000 seconds).

	// If we implement b/190097904 we should revisit these values, in particular the max delay could
	// likely be much higher again.

	return workqueue.NewTypedMaxOfRateLimiter[reconcile.Request](
		workqueue.NewTypedItemExponentialFailureRateLimiter[reconcile.Request](2*time.Second, 120*time.Second),
		// 10 qps, 100 bucket size.  This is only for retry speed and its only the overall factor (not per item)
		&workqueue.TypedBucketRateLimiter[reconcile.Request]{Limiter: rate.NewLimiter(rate.Limit(10), 100)},
	)
}

// RequeueRateLimiter slows down the periodic object re-reconcile, so that we can remain responsive to new changes.
//
// KCC schedules its objects for periodic-requeuing by returning from a
// successful Reconcile invocation with the RequeueAfter value set,
// by default to 10 minutes (with some skew to avoid a thundering-herd).
//
// The problem there is that once you have "too many objects", every object gets
// requeued, and RequeueAfter time passes before we can clear the backlog.
// So we end up with every object queued up for re-reconciliation.
//
// The big problem with that is that then _new_ changes - in particular
// user-initiated changes - are added to the back of the queue.  Then these
// user-initiated changes experience a long delay while every other object
// gets reconciled, before they get their turn.  We want to remain responsive
// to user-changes, even when there are lots of objects being re-reconciled.
//
// The workaround is to introduce an additional delay on RequeueAfter,
// to avoid the backlog building up.  We do this using a dedicated
// rate limiter, that (currently) is configured to keep the requeue
// traffic to 5 qps.  That will hopefully leave enough capacity for
// more latency sensitive reconciliations, at the expense of a longer
// delay in re-reconciliation.
func RequeueRateLimiter() RateLimiter {
	return workqueue.NewTypedMaxOfRateLimiter[reconcile.Request](
		// 5 qps, 50 bucket size.  This is the overall factor, and must be slower than the NewRateLimiter limit, to leave "room" for new items.
		&workqueue.TypedBucketRateLimiter[reconcile.Request]{Limiter: rate.NewLimiter(rate.Limit(5), 50)},
	)
}

// SetMasterRateLimiter sets the kubernetes client level rate limiter.
// This rate limiter is shared among all requests created by the client.
// If specified, it will override the QPS and Burst fields.
//
// By default, this rate limiter uses tokenBucketRateLimiter(20.0, 30).
// In ConfigConnector, this becomes a bottleneck when re-reconciliate a large amount of ConfigConnector resources.
//
// One potential downside of bumping this rate limit is that ConfigConnector could hit GCP service quotes due to the
// more aggressive GCP requests. For your information, the IAM quota has Read request 6,000 per minute, and Write requests 600 per minute. https://cloud.google.com/iam/quotas
func SetMasterRateLimiter(restConfig *rest.Config, qps float32, burst int) {
	restConfig.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(qps, burst)
}
