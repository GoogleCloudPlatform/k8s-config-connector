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

package resourceoverrides_test

import (
	"context"
	"sort"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/krmtotf"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/resourceoverrides"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/resourceoverrides/operations"
	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/appengine"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// splitTrafficLiveState returns the live state the Terraform controller has after reading
// the service from App Engine, with the shardBy App Engine returns ("" when it returns none).
func splitTrafficLiveState(shardBy string, allocations map[string]interface{}) *terraform.InstanceState {
	split := map[string]interface{}{"allocations": allocations}
	if shardBy != "" {
		split["shard_by"] = shardBy
	}
	return krmtotf.MapToInstanceState(appengine.ResourceAppEngineServiceSplitTraffic(), map[string]interface{}{
		"id":      "apps/my-project/services/default",
		"project": "my-project",
		"service": "default",
		"split":   []interface{}{split},
	})
}

// splitTrafficConfig returns the desired Terraform config, with the given spec.split.shardBy
// ("" when it is omitted).
func splitTrafficConfig(shardBy string, allocations map[string]interface{}) *terraform.ResourceConfig {
	split := map[string]interface{}{"allocations": allocations}
	if shardBy != "" {
		split["shard_by"] = shardBy
	}
	return krmtotf.MapToResourceConfig(appengine.ResourceAppEngineServiceSplitTraffic(), map[string]interface{}{
		"project": "my-project",
		"service": "default",
		"split":   []interface{}{split},
	})
}

// migrateTrafficLiveState returns the live state the Terraform controller has after reading
// the service from App Engine: split comes from the API and migrate_traffic is absent.
func migrateTrafficLiveState(allocations map[string]interface{}) *terraform.InstanceState {
	return krmtotf.MapToInstanceState(appengine.ResourceAppEngineServiceSplitTraffic(), map[string]interface{}{
		"id":      "apps/my-project/services/default",
		"project": "my-project",
		"service": "default",
		"split": []interface{}{
			map[string]interface{}{
				"shard_by":    "IP",
				"allocations": allocations,
			},
		},
	})
}

// migrateTrafficConfig returns the desired Terraform config. A nil migrateTraffic means the
// field is omitted from the spec.
func migrateTrafficConfig(migrateTraffic interface{}, allocations map[string]interface{}) *terraform.ResourceConfig {
	m := map[string]interface{}{
		"project": "my-project",
		"service": "default",
		"split": []interface{}{
			map[string]interface{}{
				"shard_by":    "IP",
				"allocations": allocations,
			},
		},
	}
	if migrateTraffic != nil {
		m["migrate_traffic"] = migrateTraffic
	}
	return krmtotf.MapToResourceConfig(appengine.ResourceAppEngineServiceSplitTraffic(), m)
}

func applySplitTrafficOverride(t *testing.T, liveState *terraform.InstanceState, config *terraform.ResourceConfig) {
	t.Helper()
	gvk := schema.GroupVersionKind{Group: "appengine.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "AppEngineServiceSplitTraffic"}
	op := &operations.PreTerraformApply{LiveState: liveState, TerraformConfig: config}
	if err := resourceoverrides.Handler.PreTerraformApply(context.Background(), gvk, op); err != nil {
		t.Fatalf("PreTerraformApply failed: %v", err)
	}
}

func splitTrafficDiff(t *testing.T, liveState *terraform.InstanceState, config *terraform.ResourceConfig) *terraform.InstanceDiff {
	t.Helper()
	diff, err := appengine.ResourceAppEngineServiceSplitTraffic().Diff(context.Background(), liveState, config, nil)
	if err != nil {
		t.Fatalf("Diff failed: %v", err)
	}
	return diff
}

func diffKeys(diff *terraform.InstanceDiff) []string {
	var keys []string
	if diff != nil {
		for k := range diff.Attributes {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)
	return keys
}

// Without the override, a one-version split whose shardBy App Engine dropped is a diff on
// split.0.shard_by alone, so Terraform updates the split on every reconcile.
func TestAppEngineServiceSplitTraffic_ShardByDroppedWithoutOverride(t *testing.T) {
	allocations := map[string]interface{}{"v2": "1"}
	liveState := splitTrafficLiveState("", allocations)
	config := splitTrafficConfig("IP", allocations)

	got := diffKeys(splitTrafficDiff(t, liveState, config))

	if diff := cmp.Diff([]string{"split.0.shard_by"}, got); diff != "" {
		t.Errorf("unexpected diff keys without the override (-want +got):\n%s", diff)
	}
}

func TestAppEngineServiceSplitTraffic_ShardByDroppedByAppEngine(t *testing.T) {
	tests := []struct {
		name string
		// liveShardBy is the shardBy App Engine returns ("" for none).
		liveShardBy       string
		liveAllocations   map[string]interface{}
		configShardBy     string
		configAllocations map[string]interface{}
		// wantDiffKeys are the attributes Terraform would update.
		wantDiffKeys []string
		// wantShardBy is the shard_by value Terraform would send with an update.
		wantShardBy string
	}{
		{
			name:              "one version, App Engine dropped shardBy: nothing to update",
			liveAllocations:   map[string]interface{}{"v2": "1"},
			configShardBy:     "IP",
			configAllocations: map[string]interface{}{"v2": "1"},
			wantShardBy:       "IP",
		},
		{
			name:              "one version, shardBy omitted: nothing to update",
			liveAllocations:   map[string]interface{}{"v2": "1"},
			configAllocations: map[string]interface{}{"v2": "1"},
		},
		{
			name:              "one version, shardBy UNSPECIFIED: nothing to update",
			liveAllocations:   map[string]interface{}{"v2": "1"},
			configShardBy:     "UNSPECIFIED",
			configAllocations: map[string]interface{}{"v2": "1"},
			wantShardBy:       "UNSPECIFIED",
		},
		{
			name:              "traffic moves to another single version: update split, send the desired shardBy",
			liveAllocations:   map[string]interface{}{"v1": "1"},
			configShardBy:     "IP",
			configAllocations: map[string]interface{}{"v2": "1"},
			wantDiffKeys:      []string{"split.0.allocations.v1", "split.0.allocations.v2"},
			wantShardBy:       "IP",
		},
		{
			name:              "one version to two versions: update split, send the desired shardBy",
			liveAllocations:   map[string]interface{}{"v2": "1"},
			configShardBy:     "IP",
			configAllocations: map[string]interface{}{"v1": "0.5", "v2": "0.5"},
			wantDiffKeys:      []string{"split.0.allocations.%", "split.0.allocations.v1", "split.0.allocations.v2"},
			wantShardBy:       "IP",
		},
		{
			name:              "two versions to one version: update split, send the desired shardBy",
			liveShardBy:       "IP",
			liveAllocations:   map[string]interface{}{"v1": "0.5", "v2": "0.5"},
			configShardBy:     "IP",
			configAllocations: map[string]interface{}{"v2": "1"},
			wantDiffKeys:      []string{"split.0.allocations.%", "split.0.allocations.v1", "split.0.allocations.v2"},
			wantShardBy:       "IP",
		},
		{
			name:              "two versions without shardBy are left alone",
			liveAllocations:   map[string]interface{}{"v1": "0.5", "v2": "0.5"},
			configShardBy:     "IP",
			configAllocations: map[string]interface{}{"v1": "0.5", "v2": "0.5"},
			wantDiffKeys:      []string{"split.0.shard_by"},
			wantShardBy:       "IP",
		},
		{
			name:              "one version, App Engine kept a different shardBy: update shardBy",
			liveShardBy:       "IP",
			liveAllocations:   map[string]interface{}{"v2": "1"},
			configShardBy:     "COOKIE",
			configAllocations: map[string]interface{}{"v2": "1"},
			wantDiffKeys:      []string{"split.0.shard_by"},
			wantShardBy:       "COOKIE",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			liveState := splitTrafficLiveState(tc.liveShardBy, tc.liveAllocations)
			config := splitTrafficConfig(tc.configShardBy, tc.configAllocations)

			applySplitTrafficOverride(t, liveState, config)
			diff := splitTrafficDiff(t, liveState, config)

			if d := cmp.Diff(tc.wantDiffKeys, diffKeys(diff)); d != "" {
				t.Errorf("unexpected diff keys (-want +got):\n%s", d)
			}
			// Update builds its request from the live state merged with the diff.
			if got := liveState.MergeDiff(diff).Attributes["split.0.shard_by"]; got != tc.wantShardBy {
				t.Errorf("shard_by sent with an update = %q, want %q", got, tc.wantShardBy)
			}
		})
	}
}

func TestAppEngineServiceSplitTraffic_ShardByOverrideSkipsCreate(t *testing.T) {
	// Attributes that would be filled in if the resource existed, but no ID: Terraform
	// treats an InstanceState without an ID as empty, so this is a create.
	liveState := &terraform.InstanceState{
		Attributes: map[string]string{
			"split.#":                "1",
			"split.0.allocations.%":  "1",
			"split.0.allocations.v2": "1",
		},
	}
	if !liveState.Empty() {
		t.Fatalf("test setup: live state without an ID should be empty")
	}
	config := splitTrafficConfig("IP", map[string]interface{}{"v2": "1"})

	applySplitTrafficOverride(t, liveState, config)

	if got, ok := liveState.Attributes["split.0.shard_by"]; ok {
		t.Errorf("override set split.0.shard_by = %q on a resource that does not exist yet", got)
	}
}

// Without the override, a change to migrateTraffic alone is a diff on migrate_traffic only.
// Terraform then sends an update with an empty update mask, which App Engine rejects when
// migrateTraffic=true.
func TestAppEngineServiceSplitTraffic_MigrateTrafficOnlyChangeWithoutOverride(t *testing.T) {
	allocations := map[string]interface{}{"v1": "0.6", "v2": "0.4"}
	liveState := migrateTrafficLiveState(allocations)
	config := migrateTrafficConfig(true, allocations)

	got := diffKeys(splitTrafficDiff(t, liveState, config))

	if diff := cmp.Diff([]string{"migrate_traffic"}, got); diff != "" {
		t.Errorf("unexpected diff keys without the override (-want +got):\n%s", diff)
	}
}

func TestAppEngineServiceSplitTraffic_PreTerraformApply(t *testing.T) {
	sameAllocations := map[string]interface{}{"v1": "0.6", "v2": "0.4"}

	tests := []struct {
		name           string
		migrateTraffic interface{}
		// liveMigrateTraffic is a migrate_traffic value already in the live state ("" for
		// none).
		liveMigrateTraffic string
		liveAllocations    map[string]interface{}
		configAllocations  map[string]interface{}
		// wantDiffKeys are the attributes Terraform would update.
		wantDiffKeys []string
		// wantMigrateTraffic is the migrate_traffic value Terraform would send with an
		// update.
		wantMigrateTraffic string
	}{
		{
			name:               "only migrateTraffic changed to true: nothing to update",
			migrateTraffic:     true,
			liveAllocations:    sameAllocations,
			configAllocations:  sameAllocations,
			wantMigrateTraffic: "true",
		},
		{
			name:               "only migrateTraffic changed from false to true: stale live value is replaced",
			migrateTraffic:     true,
			liveMigrateTraffic: "false",
			liveAllocations:    sameAllocations,
			configAllocations:  sameAllocations,
			wantMigrateTraffic: "true",
		},
		{
			name:               "migrateTraffic false and split unchanged: nothing to update",
			migrateTraffic:     false,
			liveAllocations:    sameAllocations,
			configAllocations:  sameAllocations,
			wantMigrateTraffic: "false",
		},
		{
			name:               "migrateTraffic omitted and split unchanged: nothing to update",
			liveAllocations:    sameAllocations,
			configAllocations:  sameAllocations,
			wantMigrateTraffic: "false",
		},
		{
			name:               "migrateTraffic removed from the spec: stale live true is reset to false",
			liveMigrateTraffic: "true",
			liveAllocations:    sameAllocations,
			configAllocations:  sameAllocations,
			wantMigrateTraffic: "false",
		},
		{
			name:               "split changed with migrateTraffic omitted: update split and send migrateTraffic=false",
			liveAllocations:    sameAllocations,
			configAllocations:  map[string]interface{}{"v1": "0.1", "v2": "0.9"},
			wantDiffKeys:       []string{"split.0.allocations.v1", "split.0.allocations.v2"},
			wantMigrateTraffic: "false",
		},
		{
			name:               "allocations written differently are equal",
			migrateTraffic:     true,
			liveAllocations:    sameAllocations,
			configAllocations:  map[string]interface{}{"v1": "0.60", "v2": "0.40"},
			wantMigrateTraffic: "true",
		},
		{
			name:               "split changed with migrateTraffic true: update split and send migrateTraffic=true",
			migrateTraffic:     true,
			liveAllocations:    sameAllocations,
			configAllocations:  map[string]interface{}{"v1": "0.1", "v2": "0.9"},
			wantDiffKeys:       []string{"split.0.allocations.v1", "split.0.allocations.v2"},
			wantMigrateTraffic: "true",
		},
		{
			name:               "version added to the split is still a diff",
			migrateTraffic:     false,
			liveAllocations:    map[string]interface{}{"v1": "0.5", "v2": "0.5"},
			configAllocations:  map[string]interface{}{"v1": "0.5", "v2": "0.3", "v3": "0.2"},
			wantDiffKeys:       []string{"split.0.allocations.%", "split.0.allocations.v2", "split.0.allocations.v3"},
			wantMigrateTraffic: "false",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			liveState := migrateTrafficLiveState(tc.liveAllocations)
			if tc.liveMigrateTraffic != "" {
				liveState.Attributes["migrate_traffic"] = tc.liveMigrateTraffic
			}
			config := migrateTrafficConfig(tc.migrateTraffic, tc.configAllocations)

			applySplitTrafficOverride(t, liveState, config)
			diff := splitTrafficDiff(t, liveState, config)

			if d := cmp.Diff(tc.wantDiffKeys, diffKeys(diff)); d != "" {
				t.Errorf("unexpected diff keys (-want +got):\n%s", d)
			}
			// Update builds its request from the live state merged with the diff.
			got := liveState.MergeDiff(diff).Attributes["migrate_traffic"]
			if got != tc.wantMigrateTraffic {
				t.Errorf("migrate_traffic sent with an update = %q, want %q", got, tc.wantMigrateTraffic)
			}
		})
	}
}

func TestAppEngineServiceSplitTraffic_PreTerraformApply_Create(t *testing.T) {
	liveState := &terraform.InstanceState{}
	config := migrateTrafficConfig(true, map[string]interface{}{"v1": "1"})

	applySplitTrafficOverride(t, liveState, config)

	if _, found := liveState.Attributes["migrate_traffic"]; found {
		t.Errorf("override changed the live state of a resource that does not exist yet: %v", liveState.Attributes)
	}
}

// A value that is not yet known (for example from an unresolved reference) is not a
// concrete bool, so the live state is left alone.
func TestAppEngineServiceSplitTraffic_PreTerraformApply_UnknownValue(t *testing.T) {
	liveState := migrateTrafficLiveState(map[string]interface{}{"v1": "1"})
	liveState.Attributes["migrate_traffic"] = "true"
	// Terraform's placeholder for an unknown value (hcl2shim.UnknownVariableValue, which
	// is in an internal package).
	const unknownValue = "74D93920-ED26-11E3-AC10-0800200C9A66"
	config := terraform.NewResourceConfigRaw(map[string]interface{}{
		"project":         "my-project",
		"service":         "default",
		"migrate_traffic": unknownValue,
	})

	applySplitTrafficOverride(t, liveState, config)

	if got := liveState.Attributes["migrate_traffic"]; got != "true" {
		t.Errorf("migrate_traffic in live state = %q, want unchanged %q", got, "true")
	}
}
