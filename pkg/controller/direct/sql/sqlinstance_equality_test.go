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

package sql

import (
	"reflect"
	"testing"

	api "google.golang.org/api/sqladmin/v1beta4"
)

func TestDiffInstances_StrictPointersMatch(t *testing.T) {
	t.Parallel()
	// desired has all optional struct pointers as nil (typical for a minimal KRM spec)
	desired := &api.DatabaseInstance{
		Settings: &api.Settings{},
	}

	// actual has these fields populated with "empty" defaults by the GCP API
	actual := &api.DatabaseInstance{
		DiskEncryptionConfiguration: &api.DiskEncryptionConfiguration{
			Kind: "sql#diskEncryptionConfiguration",
		},
		ReplicaConfiguration: &api.ReplicaConfiguration{
			Kind: "sql#replicaConfiguration",
		},
		ReplicationCluster: &api.ReplicationCluster{},
		Settings: &api.Settings{
			BackupConfiguration: &api.BackupConfiguration{
				Kind: "sql#backupConfiguration",
			},
			DataCacheConfig: &api.DataCacheConfig{},
			IpConfiguration: &api.IpConfiguration{
				Ipv4Enabled: true,
				SslMode:     "ALLOW_UNENCRYPTED_AND_ENCRYPTED",
			},
			LocationPreference: &api.LocationPreference{
				Kind: "sql#locationPreference",
			},
		},
	}

	diff := DiffInstances(desired, actual)

	// We want HasDiff() to be false because these are semantically equivalent.
	// Nil values in desired (KRM) should match default/empty values in actual (GCP).
	if diff.HasDiff() {
		t.Errorf("DiffInstances() identified unexpected diffs: %v", diff.Fields)
	}
}

func TestDiffInstances_NilSettings(t *testing.T) {
	t.Parallel()
	// desired has Settings as nil
	desired := &api.DatabaseInstance{
		Settings: nil,
	}

	// actual has Settings as an empty object
	actual := &api.DatabaseInstance{
		Settings: &api.Settings{},
	}

	diff := DiffInstances(desired, actual)

	if diff.HasDiff() {
		t.Errorf("DiffInstances() identified unexpected diffs when Settings is nil in desired: %v", diff.Fields)
	}
}

func TestDiffInstances_NilUserLabels(t *testing.T) {
	t.Parallel()
	// desired has nil UserLabels
	desired := &api.DatabaseInstance{
		Settings: &api.Settings{
			UserLabels: nil,
		},
	}

	// actual has an empty UserLabels map
	actual := &api.DatabaseInstance{
		Settings: &api.Settings{
			UserLabels: map[string]string{},
		},
	}

	diff := DiffInstances(desired, actual)

	if diff.HasDiff() {
		t.Errorf("DiffInstances() identified unexpected diffs when UserLabels is nil in desired: %v", diff.Fields)
	}
}

func TestDiffInstances_AuthorizedNetworksSorting(t *testing.T) {
	t.Parallel()
	desired := &api.DatabaseInstance{
		Settings: &api.Settings{
			IpConfiguration: &api.IpConfiguration{
				AuthorizedNetworks: []*api.AclEntry{
					{Name: "b", Value: "2.2.2.2"},
					{Name: "a", Value: "1.1.1.1"},
					{Name: "", Value: "4.4.4.4"},
					{Name: "", Value: "3.3.3.3"},
				},
			},
		},
	}

	actual := &api.DatabaseInstance{
		Settings: &api.Settings{
			IpConfiguration: &api.IpConfiguration{
				AuthorizedNetworks: []*api.AclEntry{
					{Name: "", Value: "3.3.3.3"},
					{Name: "a", Value: "1.1.1.1"},
					{Name: "b", Value: "2.2.2.2"},
					{Name: "", Value: "4.4.4.4"},
				},
			},
		},
	}

	diff := DiffInstances(desired, actual)

	if diff.HasDiff() {
		t.Errorf("DiffInstances() identified unexpected diffs due to AuthorizedNetworks sorting: %v", diff.Fields)
	}
}

func TestDiffInstances_DatabaseFlagsSorting(t *testing.T) {
	t.Parallel()
	desired := &api.DatabaseInstance{
		Settings: &api.Settings{
			DatabaseFlags: []*api.DatabaseFlags{
				{Name: "b", Value: "v2"},
				{Name: "a", Value: "v1"},
			},
		},
	}

	actual := &api.DatabaseInstance{
		Settings: &api.Settings{
			DatabaseFlags: []*api.DatabaseFlags{
				{Name: "a", Value: "v1"},
				{Name: "b", Value: "v2"},
			},
		},
	}

	diff := DiffInstances(desired, actual)

	if diff.HasDiff() {
		t.Errorf("DiffInstances() identified unexpected diffs due to DatabaseFlags sorting: %v", diff.Fields)
	}
}

func TestDiffInstances_UserLabelsDiff(t *testing.T) {
	t.Parallel()
	desired := &api.DatabaseInstance{
		Settings: &api.Settings{
			UserLabels: map[string]string{
				"key1": "val1",
				"key2": "val2-changed",
				"key3": "val3-new",
			},
		},
	}

	actual := &api.DatabaseInstance{
		Settings: &api.Settings{
			UserLabels: map[string]string{
				"key1": "val1",
				"key2": "val2",
				"key4": "val4-removed",
			},
		},
	}

	diff := DiffInstances(desired, actual)

	if !diff.HasDiff() {
		t.Errorf("DiffInstances() expected diffs, but got none")
	}

	expectedDiffs := map[string]struct {
		Old any
		New any
	}{
		".settings.userLabels[*\"key2\"]": {Old: "val2", New: "val2-changed"},
		".settings.userLabels[+\"key3\"]": {Old: nil, New: "val3-new"},
		".settings.userLabels[-\"key4\"]": {Old: "val4-removed", New: nil},
	}

	if len(diff.Fields) != len(expectedDiffs) {
		t.Fatalf("DiffInstances() expected %d diffs, got %d. Fields: %v", len(expectedDiffs), len(diff.Fields), diff.Fields)
	}

	for _, field := range diff.Fields {
		expected, ok := expectedDiffs[field.ID]
		if !ok {
			t.Errorf("Unexpected diff field ID %q", field.ID)
			continue
		}
		if field.Old != expected.Old {
			t.Errorf("For field %q, expected Old=%v, got %v", field.ID, expected.Old, field.Old)
		}
		if field.New != expected.New {
			t.Errorf("For field %q, expected New=%v, got %v", field.ID, expected.New, field.New)
		}
	}
}

func TestDiffInstances_AvailabilityTypeCasing(t *testing.T) {
	desired := &api.DatabaseInstance{
		Settings: &api.Settings{
			AvailabilityType: "Zonal",
		},
	}

	actual := &api.DatabaseInstance{
		Settings: &api.Settings{
			AvailabilityType: "ZONAL",
		},
	}

	diff := DiffInstances(desired, actual)

	if diff.HasDiff() {
		t.Errorf("DiffInstances() expected no diffs, but got: %v", diff.Fields)
	}
}

func TestDiffInstances_ServerCa_NoDiff(t *testing.T) {
	t.Parallel()
	desired := &api.DatabaseInstance{
		Settings: &api.Settings{
			IpConfiguration: &api.IpConfiguration{
				ServerCaMode:                  "CUSTOMER_MANAGED_CAS_CA",
				ServerCaPool:                  "projects/test-project/locations/us-central1/caPools/my-ca-pool",
				CustomSubjectAlternativeNames: []string{"a.example.com", "b.example.com"},
			},
		},
	}

	actual := &api.DatabaseInstance{
		Settings: &api.Settings{
			IpConfiguration: &api.IpConfiguration{
				ServerCaMode:                  "CUSTOMER_MANAGED_CAS_CA",
				ServerCaPool:                  "projects/test-project/locations/us-central1/caPools/my-ca-pool",
				CustomSubjectAlternativeNames: []string{"a.example.com", "b.example.com"},
			},
		},
	}

	diff := DiffInstances(desired, actual)

	if diff.HasDiff() {
		t.Errorf("DiffInstances() expected no diffs for matching Server CA and SANs, but got: %v", diff.Fields)
	}
}

func TestDiffInstances_ServerCa_DefaultNoDiff(t *testing.T) {
	t.Parallel()
	// Unset serverCaMode in desired should match default GOOGLE_MANAGED_INTERNAL_CA in actual
	desired := &api.DatabaseInstance{
		Settings: &api.Settings{
			IpConfiguration: &api.IpConfiguration{
				ServerCaMode: "",
			},
		},
	}

	actual := &api.DatabaseInstance{
		Settings: &api.Settings{
			IpConfiguration: &api.IpConfiguration{
				ServerCaMode: "GOOGLE_MANAGED_INTERNAL_CA",
			},
		},
	}

	diff := DiffInstances(desired, actual)

	if diff.HasDiff() {
		t.Errorf("DiffInstances() expected no diffs for default GOOGLE_MANAGED_INTERNAL_CA, but got: %v", diff.Fields)
	}
}

func TestDiffInstances_ServerCa_WithDiff(t *testing.T) {
	t.Parallel()
	desired := &api.DatabaseInstance{
		Settings: &api.Settings{
			IpConfiguration: &api.IpConfiguration{
				ServerCaMode:                  "CUSTOMER_MANAGED_CAS_CA",
				ServerCaPool:                  "projects/test-project/locations/us-central1/caPools/pool-2",
				CustomSubjectAlternativeNames: []string{"c.example.com", "a.example.com"},
			},
		},
	}

	actual := &api.DatabaseInstance{
		Settings: &api.Settings{
			IpConfiguration: &api.IpConfiguration{
				ServerCaMode:                  "GOOGLE_MANAGED_CAS_CA",
				ServerCaPool:                  "projects/test-project/locations/us-central1/caPools/pool-1",
				CustomSubjectAlternativeNames: []string{"b.example.com", "a.example.com"},
			},
		},
	}

	diff := DiffInstances(desired, actual)

	if !diff.HasDiff() {
		t.Fatalf("DiffInstances() expected diffs, but got none")
	}

	expectedDiffs := map[string]struct {
		Old any
		New any
	}{
		".settings.ipConfiguration.serverCaMode":                  {Old: "GOOGLE_MANAGED_CAS_CA", New: "CUSTOMER_MANAGED_CAS_CA"},
		".settings.ipConfiguration.serverCaPool":                  {Old: "projects/test-project/locations/us-central1/caPools/pool-1", New: "projects/test-project/locations/us-central1/caPools/pool-2"},
		".settings.ipConfiguration.customSubjectAlternativeNames": {Old: []string{"b.example.com", "a.example.com"}, New: []string{"c.example.com", "a.example.com"}},
	}

	if len(diff.Fields) != len(expectedDiffs) {
		t.Fatalf("DiffInstances() expected %d diffs, got %d. Fields: %v", len(expectedDiffs), len(diff.Fields), diff.Fields)
	}

	for _, field := range diff.Fields {
		expected, ok := expectedDiffs[field.ID]
		if !ok {
			t.Errorf("Unexpected diff field ID %q", field.ID)
			continue
		}
		if !reflect.DeepEqual(field.Old, expected.Old) {
			t.Errorf("For field %q, expected Old=%v, got %v", field.ID, expected.Old, field.Old)
		}
		if !reflect.DeepEqual(field.New, expected.New) {
			t.Errorf("For field %q, expected New=%v, got %v", field.ID, expected.New, field.New)
		}
	}
}
