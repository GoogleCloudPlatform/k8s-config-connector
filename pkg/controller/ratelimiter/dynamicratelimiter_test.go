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
	"strings"
	"testing"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"k8s.io/apimachinery/pkg/types"
)

func TestParseBackoffMaxDelay(t *testing.T) {
	tests := []struct {
		name        string
		annotations map[string]string
		want        *time.Duration
		wantErr     string
	}{
		{
			name:        "nil annotations",
			annotations: nil,
			want:        nil,
			wantErr:     "",
		},
		{
			name:        "annotation not set",
			annotations: map[string]string{"foo": "bar"},
			want:        nil,
			wantErr:     "",
		},
		{
			name: "empty annotation value",
			annotations: map[string]string{
				k8s.BackoffMaxDelayInSecondsAnnotation: "",
			},
			want:    nil,
			wantErr: "",
		},
		{
			name: "zero value - halt retries",
			annotations: map[string]string{
				k8s.BackoffMaxDelayInSecondsAnnotation: "0",
			},
			want:    durationPtr(0),
			wantErr: "",
		},
		{
			name: "positive value 10s",
			annotations: map[string]string{
				k8s.BackoffMaxDelayInSecondsAnnotation: "10",
			},
			want:    durationPtr(10 * time.Second),
			wantErr: "",
		},
		{
			name: "positive value 600s",
			annotations: map[string]string{
				k8s.BackoffMaxDelayInSecondsAnnotation: "600",
			},
			want:    durationPtr(600 * time.Second),
			wantErr: "",
		},
		{
			name: "negative value",
			annotations: map[string]string{
				k8s.BackoffMaxDelayInSecondsAnnotation: "-1",
			},
			want:    nil,
			wantErr: "must be non-negative",
		},
		{
			name: "non-integer string",
			annotations: map[string]string{
				k8s.BackoffMaxDelayInSecondsAnnotation: "invalid",
			},
			want:    nil,
			wantErr: "must be an integer",
		},
		{
			name: "floating point string",
			annotations: map[string]string{
				k8s.BackoffMaxDelayInSecondsAnnotation: "1.5",
			},
			want:    nil,
			wantErr: "must be an integer",
		},
		{
			name: "integer overflow exceeding int32",
			annotations: map[string]string{
				k8s.BackoffMaxDelayInSecondsAnnotation: "9999999999999999999",
			},
			want:    nil,
			wantErr: "must be an integer",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := ParseBackoffMaxDelay(tc.annotations)
			if tc.wantErr != "" {
				if err == nil {
					t.Fatalf("expected error containing %q, got nil", tc.wantErr)
				}
				if !strings.Contains(err.Error(), tc.wantErr) {
					t.Fatalf("expected error containing %q, got %v", tc.wantErr, err)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if tc.want == nil {
				if got != nil {
					t.Fatalf("expected nil duration, got %v", *got)
				}
				return
			}
			if got == nil {
				t.Fatalf("expected %v, got nil", *tc.want)
			}
			if *got != *tc.want {
				t.Fatalf("expected %v, got %v", *tc.want, *got)
			}
		})
	}
}

func TestDynamicRateLimiter_NextDelayAndForget(t *testing.T) {
	rl := NewDynamicRateLimiter()

	nn1 := types.NamespacedName{Namespace: "default", Name: "res1"}
	nn2 := types.NamespacedName{Namespace: "default", Name: "res2"}

	baseDelay := 1 * time.Second
	maxDelay := 10 * time.Second

	// 1st failure for nn1: 1 * 2^0 = 1s
	if delay := rl.NextDelay(nn1, baseDelay, maxDelay); delay != 1*time.Second {
		t.Fatalf("expected 1s on first failure, got %v", delay)
	}

	// 2nd failure for nn1: 1 * 2^1 = 2s
	if delay := rl.NextDelay(nn1, baseDelay, maxDelay); delay != 2*time.Second {
		t.Fatalf("expected 2s on second failure, got %v", delay)
	}

	// 1st failure for nn2 (isolated from nn1): 1 * 2^0 = 1s
	if delay := rl.NextDelay(nn2, baseDelay, maxDelay); delay != 1*time.Second {
		t.Fatalf("expected 1s for nn2 first failure, got %v", delay)
	}

	// 3rd failure for nn1: 1 * 2^2 = 4s
	if delay := rl.NextDelay(nn1, baseDelay, maxDelay); delay != 4*time.Second {
		t.Fatalf("expected 4s on third failure, got %v", delay)
	}

	// 4th failure for nn1: 1 * 2^3 = 8s
	if delay := rl.NextDelay(nn1, baseDelay, maxDelay); delay != 8*time.Second {
		t.Fatalf("expected 8s on fourth failure, got %v", delay)
	}

	// 5th failure for nn1: 1 * 2^4 = 16s -> capped at 10s
	if delay := rl.NextDelay(nn1, baseDelay, maxDelay); delay != 10*time.Second {
		t.Fatalf("expected 10s (capped) on fifth failure, got %v", delay)
	}

	// 6th failure for nn1: capped at 10s
	if delay := rl.NextDelay(nn1, baseDelay, maxDelay); delay != 10*time.Second {
		t.Fatalf("expected 10s (capped) on sixth failure, got %v", delay)
	}

	// Forget nn1 on success
	rl.Forget(nn1)

	// After forget, failure count is reset to 0: next delay is 1s
	if delay := rl.NextDelay(nn1, baseDelay, maxDelay); delay != 1*time.Second {
		t.Fatalf("expected 1s after forget, got %v", delay)
	}

	// nn2 was not forgotten: 2nd failure for nn2 should be 2s
	if delay := rl.NextDelay(nn2, baseDelay, maxDelay); delay != 2*time.Second {
		t.Fatalf("expected 2s for nn2 second failure, got %v", delay)
	}
}

func TestDynamicRateLimiter_OverflowProtection(t *testing.T) {
	rl := NewDynamicRateLimiter()
	nn := types.NamespacedName{Namespace: "default", Name: "overflow-test"}

	baseDelay := 1 * time.Second
	maxDelay := 120 * time.Second

	// Simulate 100 consecutive failures
	for i := 0; i < 100; i++ {
		delay := rl.NextDelay(nn, baseDelay, maxDelay)
		if delay > maxDelay || delay <= 0 {
			t.Fatalf("iteration %d: delay %v exceeded maxDelay or was non-positive", i, delay)
		}
	}
}

func durationPtr(d time.Duration) *time.Duration {
	return &d
}
