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

func TestSetupValidateNormalizesFields(t *testing.T) {
	opts := newSetupOptions()
	opts.contexts = []string{" cluster-a ", "cluster-b "}
	opts.bucket = " gs://my-ha-bucket/ "

	err := opts.validate()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	wantContexts := []string{"cluster-a", "cluster-b"}
	if !slices.Equal(opts.contexts, wantContexts) {
		t.Fatalf("got contexts %v, want %v", opts.contexts, wantContexts)
	}
	if opts.bucket != "my-ha-bucket" {
		t.Fatalf("got bucket %q, want %q", opts.bucket, "my-ha-bucket")
	}
}

func TestSetupValidateErrors(t *testing.T) {
	tests := []struct {
		name              string
		contexts          []string
		bucket            string
		wantErrSubstrings []string
	}{
		{
			name:              "fewer than 2 contexts",
			contexts:          []string{"cluster-a"},
			bucket:            "my-bucket",
			wantErrSubstrings: []string{"at least 2 cluster contexts must be specified"},
		},
		{
			name:              "duplicate contexts padded with whitespace",
			contexts:          []string{"cluster-a", " cluster-a "},
			bucket:            "my-bucket",
			wantErrSubstrings: []string{"duplicate context \"cluster-a\""},
		},
		{
			name:              "empty bucket",
			contexts:          []string{"cluster-a", "cluster-b"},
			bucket:            "",
			wantErrSubstrings: []string{"--bucket is required"},
		},
		{
			name:              "whitespace-only bucket",
			contexts:          []string{"cluster-a", "cluster-b"},
			bucket:            "   ",
			wantErrSubstrings: []string{"--bucket is required"},
		},
		{
			name:              "bare gs:// prefix without bucket name",
			contexts:          []string{"cluster-a", "cluster-b"},
			bucket:            "gs://   ",
			wantErrSubstrings: []string{"--bucket is required"},
		},
		{
			name:              "accumulates multiple flag errors",
			contexts:          []string{"cluster-a"},
			bucket:            "",
			wantErrSubstrings: []string{"at least 2 cluster contexts", "--bucket is required"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			opts := newSetupOptions()
			opts.contexts = slices.Clone(tc.contexts)
			opts.bucket = tc.bucket

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

func TestRunSetup(t *testing.T) {
	useTestKubeconfig(t, "cluster-a", "cluster-b")
	opts := newSetupOptions()
	opts.contexts = []string{"cluster-a", "cluster-b"}
	opts.bucket = "my-ha-bucket"

	err := runSetup(context.Background(), opts)

	if err == nil || !strings.Contains(err.Error(), "ha setup is not yet implemented") {
		t.Fatalf("expected not yet implemented error from runSetup, got %v", err)
	}
}
