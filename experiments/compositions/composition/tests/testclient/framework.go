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

package testclient

import (
	"testing"
	"time"

	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/util/retry"
)

const opDuration = 2 * time.Second

// Poll - polls for op to return without an error. If op has not executed
// without error by timeout, Poll fails the test.
func Poll(t *testing.T, op func() error, timeout time.Duration) {
	t.Helper()

	retryFrequency := wait.Backoff{
		Duration: opDuration,
		Steps:    int(timeout / opDuration),
	}
	err := retry.OnError(retryFrequency, func(_ error) bool { return true }, op)
	if err != nil {
		t.Errorf("timeout")
		t.FailNow()
	}
}

// getFrequency - calculates the cadence at which test checks, time boxed
// by duration, are retried until timeout
func getFrequency(t *testing.T, timeout time.Duration) wait.Backoff {
	t.Helper()
	return wait.Backoff{
		Duration: opDuration,
		Steps:    int(timeout / opDuration),
	}
}
