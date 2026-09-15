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

package main

import (
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func getSortedKeys(m map[string]KindTarget) []string {
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func TestTargetResolver(t *testing.T) {
	repoRoot, err := findRepoRoot()
	if err != nil {
		t.Fatalf("findRepoRoot() failed: %v", err)
	}

	resolver, err := NewTargetResolver(repoRoot)
	if err != nil {
		t.Fatalf("NewTargetResolver() failed: %v", err)
	}

	t.Run("Rule 1 - Manifest changes with stable v1beta1 subtest folders (Exact match)", func(t *testing.T) {
		changedFiles := []string{
			"pkg/test/resourcefixture/testdata/basic/storage/v1beta1/storagebucket/storagebucketbasic/create.yaml",
			"pkg/test/resourcefixture/testdata/basic/pubsub/v1beta1/pubsubsubscription/basicpubsubsubscription/dependencies.yaml",
		}

		got, err := resolver.ResolveTargets(changedFiles)
		if err != nil {
			t.Fatalf("ResolveTargets() failed: %v", err)
		}

		want := map[string]KindTarget{
			"PubSubSubscription": {
				All:   false,
				Tests: []string{"basicpubsubsubscription"},
			},
			"StorageBucket": {
				All:   false,
				Tests: []string{"storagebucketbasic"},
			},
		}

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("ResolveTargets() mismatch (-want +got):\n%s", diff)
		}
	})

	t.Run("Rule 2 - Direct controller, types, identity (Key-only exact match and All=true)", func(t *testing.T) {
		changedFiles := []string{
			"pkg/controller/direct/compute/computerouternat_controller.go",
			"apis/cloudidentity/v1beta1/cloudidentitygroup_identity.go",
		}

		got, err := resolver.ResolveTargets(changedFiles)
		if err != nil {
			t.Fatalf("ResolveTargets() failed: %v", err)
		}

		// Exact match on keys only
		wantKinds := []string{"CloudIdentityGroup", "ComputeRouterNAT"}
		gotKinds := getSortedKeys(got)

		if diff := cmp.Diff(wantKinds, gotKinds); diff != "" {
			t.Errorf("Resolved Kind keys mismatch (-want +got):\n%s", diff)
		}

		// Verify that tests are present and All is true for whole-kind changes
		for _, kind := range wantKinds {
			target := got[kind]
			if !target.All {
				t.Errorf("expected All=true for Kind %q", kind)
			}
			if len(target.Tests) == 0 {
				t.Errorf("expected non-empty test folders for Kind %q", kind)
			}
		}
	})

	t.Run("Rule 3 - Reference changes (Selective match and All=false)", func(t *testing.T) {
		changedFiles := []string{
			"apis/compute/v1beta1/computenetwork_reference.go",
		}

		got, err := resolver.ResolveTargets(changedFiles)
		if err != nil {
			t.Fatalf("ResolveTargets() failed: %v", err)
		}

		// The expected kinds must be either the same set or a subset of the actual output (because more resources
		// might reference ComputeNetworkRef in the future).
		expectedKinds := []string{
			"ComputeRoute",
			"ComputeRouter",
		}

		for _, kind := range expectedKinds {
			target, exists := got[kind]
			if !exists || len(target.Tests) == 0 {
				t.Errorf("expected output to contain non-empty test list for Kind %q, got empty/missing", kind)
			}
			if target.All {
				t.Errorf("expected All=false for reference change on Kind %q", kind)
			}
		}
	})
}
