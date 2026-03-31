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
	"testing"

	api "google.golang.org/api/sqladmin/v1beta4"
)

func TestDiffInstances_StrictPointersMatch(t *testing.T) {
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
