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

package ratelimiter

import (
	"fmt"
	"math"
	"strconv"
	"sync"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"k8s.io/apimachinery/pkg/types"
)

// ParseBackoffMaxDelay extracts and validates the annotation value.
// Returns:
//   - (nil, nil) if annotation is unset or empty.
//   - (ptr(0), nil) if set to "0" (halt retries).
//   - (ptr(duration), nil) if set to a positive integer.
//   - (nil, error) if value is negative or not a valid integer.
func ParseBackoffMaxDelay(annotations map[string]string) (*time.Duration, error) {
	val, ok := annotations[k8s.BackoffMaxDelayInSecondsAnnotation]
	if !ok || val == "" {
		return nil, nil
	}
	seconds, err := strconv.ParseInt(val, 10, 32)
	if err != nil {
		return nil, fmt.Errorf("invalid value %q for annotation %s: must be an integer", val, k8s.BackoffMaxDelayInSecondsAnnotation)
	}
	if seconds < 0 {
		return nil, fmt.Errorf("invalid value %q for annotation %s: must be non-negative", val, k8s.BackoffMaxDelayInSecondsAnnotation)
	}
	d := time.Duration(seconds) * time.Second
	return &d, nil
}

// DynamicRateLimiter tracks consecutive reconciliation failures per resource
// and calculates exponential retry intervals clamped to the resource's max delay.
type DynamicRateLimiter struct {
	mu       sync.Mutex
	failures map[types.NamespacedName]int
}

func NewDynamicRateLimiter() *DynamicRateLimiter {
	return &DynamicRateLimiter{
		failures: make(map[types.NamespacedName]int),
	}
}

// NextDelay increments the failure count and calculates the exponential backoff:
// baseDelay * 2^(failures-1), capped at maxDelay.
func (r *DynamicRateLimiter) NextDelay(nn types.NamespacedName, baseDelay, maxDelay time.Duration) time.Duration {
	r.mu.Lock()
	defer r.mu.Unlock()

	exp := r.failures[nn]
	r.failures[nn]++

	backoff := float64(baseDelay) * math.Pow(2, float64(exp))
	if backoff > float64(maxDelay) || backoff <= 0 {
		return maxDelay
	}
	return time.Duration(backoff)
}

// Forget clears the failure count when reconciliation succeeds (matching controller-runtime semantics).
func (r *DynamicRateLimiter) Forget(nn types.NamespacedName) {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.failures, nn)
}
