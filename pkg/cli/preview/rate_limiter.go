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
	"net/url"

	"golang.org/x/time/rate"
	"k8s.io/klog/v2"
)

func (c *interceptingGCPClient) getOrCreateRateLimiter(u *url.URL) (*rate.Limiter, error) {
	if u == nil {
		klog.Fatal("nil URL for GCP call")
	}
	klog.V(2).Info("getOrCreateRateLimiter", "API", u.Host)
	if _, ok := c.rateLimiters[u.Host]; !ok {
		c.rateLimiters[u.Host] = rate.NewLimiter(rate.Limit(c.qps), c.burst)
	}
	return c.rateLimiters[u.Host], nil
}
