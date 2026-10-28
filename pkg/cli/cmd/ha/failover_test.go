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

package ha

import (
	"context"
	"slices"
	"strings"
	"testing"
)

func TestFailoverValidateNormalizesFields(t *testing.T) {
	opts := newFailoverOptions()
	opts.contexts = []string{" cluster-a ", " cluster-b"}
	opts.targetCluster = " cluster-b "

	err := opts.validate()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	wantContexts := []string{"cluster-a", "cluster-b"}
	if !slices.Equal(opts.contexts, wantContexts) {
		t.Fatalf("got contexts %v, want %v", opts.contexts, wantContexts)
	}
	if opts.targetCluster != "cluster-b" {
		t.Fatalf("got targetCluster %q, want %q", opts.targetCluster, "cluster-b")
	}
}

func TestFailoverValidateErrors(t *testing.T) {
	tests := []struct {
		name              string
		contexts          []string
		targetCluster     string
		wantErrSubstrings []string
	}{
		{
			name:              "missing target-cluster",
			contexts:          []string{"cluster-a", "cluster-b"},
			wantErrSubstrings: []string{"--target-cluster is required"},
		},
		{
			name:              "target-cluster not in contexts",
			contexts:          []string{"cluster-a", "cluster-b"},
			targetCluster:     "cluster-c",
			wantErrSubstrings: []string{"must be one of the contexts specified in --contexts"},
		},
		{
			name:              "accumulates contexts and target-cluster errors",
			contexts:          []string{"cluster-a"},
			targetCluster:     "",
			wantErrSubstrings: []string{"at least 2 cluster contexts", "--target-cluster is required"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			opts := newFailoverOptions()
			opts.contexts = slices.Clone(tc.contexts)
			opts.targetCluster = tc.targetCluster

			err := opts.validate()

			if err == nil {
				t.Fatalf("expected error containing %v, got nil", tc.wantErrSubstrings)
			}
			for _, sub := range tc.wantErrSubstrings {
				if !strings.Contains(err.Error(), sub) {
					t.Fatalf("expected error %q to contain %q", err.Error(), sub)
				}
			}
		})
	}
}

func TestRunFailover(t *testing.T) {
	useTestKubeconfig(t, "cluster-a", "cluster-b")
	opts := newFailoverOptions()
	opts.contexts = []string{"cluster-a", "cluster-b"}
	opts.targetCluster = "cluster-b"

	err := runFailover(context.Background(), opts)

	if err == nil || !strings.Contains(err.Error(), "ha failover is not yet implemented") {
		t.Fatalf("expected not yet implemented error from runFailover, got %v", err)
	}
}
