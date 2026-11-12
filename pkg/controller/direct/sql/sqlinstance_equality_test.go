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
	"context"
	"encoding/json"
	"reflect"
	"slices"
	"strings"
	"testing"

	computev1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/changecookies"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	api "google.golang.org/api/sqladmin/v1beta4"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/sql/v1beta1"
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
	// Unset serverCaMode in KRM should adopt the live value from actual and produce no diff.
	actual := &api.DatabaseInstance{
		Settings: &api.Settings{
			BackupConfiguration: &api.BackupConfiguration{},
			IpConfiguration: &api.IpConfiguration{
				Ipv4Enabled:  true,
				SslMode:      "ALLOW_UNENCRYPTED_AND_ENCRYPTED",
				ServerCaMode: "GOOGLE_MANAGED_INTERNAL_CA",
			},
		},
	}

	krmObj := &krm.SQLInstance{
		Spec: krm.SQLInstanceSpec{
			Settings: krm.InstanceSettings{
				IpConfiguration: &krm.InstanceIpConfiguration{},
			},
		},
	}

	desired, err := SQLInstanceKRMToGCP(krmObj, actual, nil)
	if err != nil {
		t.Fatalf("SQLInstanceKRMToGCP() failed: %v", err)
	}

	diff := DiffIpConfiguration(desired.Settings.IpConfiguration, actual.Settings.IpConfiguration)

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

func TestDiffInstances_ServerCa_CustomSANs_OrderIndependent(t *testing.T) {
	t.Parallel()
	desired := &api.DatabaseInstance{
		Settings: &api.Settings{
			IpConfiguration: &api.IpConfiguration{
				CustomSubjectAlternativeNames: []string{"a.example.com", "b.example.com", "c.example.com"},
			},
		},
	}

	actual := &api.DatabaseInstance{
		Settings: &api.Settings{
			IpConfiguration: &api.IpConfiguration{
				CustomSubjectAlternativeNames: []string{"c.example.com", "a.example.com", "b.example.com"},
			},
		},
	}

	diff := DiffInstances(desired, actual)

	if diff.HasDiff() {
		t.Errorf("DiffInstances() expected no diffs for re-ordered SANs, but got: %v", diff.Fields)
	}
}

func TestDiffInstances_ServerCaPool_PrefixNormalization(t *testing.T) {
	t.Parallel()
	desired := &api.DatabaseInstance{
		Settings: &api.Settings{
			IpConfiguration: &api.IpConfiguration{
				ServerCaPool: "projects/test-project/locations/us-central1/caPools/my-ca-pool",
			},
		},
	}

	actual := &api.DatabaseInstance{
		Settings: &api.Settings{
			IpConfiguration: &api.IpConfiguration{
				ServerCaPool: "//privateca.googleapis.com/projects/test-project/locations/us-central1/caPools/my-ca-pool",
			},
		},
	}

	diff := DiffInstances(desired, actual)

	if diff.HasDiff() {
		t.Errorf("DiffInstances() expected no diffs for normalized CA pool prefixes, but got: %v", diff.Fields)
	}
}

func TestToProto_PreservesServerCaFields(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	adapter := &sqlInstanceAdapter{}

	base := &api.DatabaseInstance{
		Name: "my-instance",
		Settings: &api.Settings{
			IpConfiguration: &api.IpConfiguration{
				ServerCaMode:                  "CUSTOMER_MANAGED_CAS_CA",
				ServerCaPool:                  "projects/test-project/locations/us-central1/caPools/pool-1",
				CustomSubjectAlternativeNames: []string{"db.example.com"},
			},
		},
	}
	updated := &api.DatabaseInstance{
		Name: "my-instance",
		Settings: &api.Settings{
			IpConfiguration: &api.IpConfiguration{
				ServerCaMode:                  "CUSTOMER_MANAGED_CAS_CA",
				ServerCaPool:                  "projects/test-project/locations/us-central1/caPools/pool-2",
				CustomSubjectAlternativeNames: []string{"db.example.com", "db-replica.example.com"},
			},
		},
	}

	baseProto, err := adapter.toProto(ctx, base)
	if err != nil {
		t.Fatalf("toProto(base) failed: %v", err)
	}
	updatedProto, err := adapter.toProto(ctx, updated)
	if err != nil {
		t.Fatalf("toProto(updated) failed: %v", err)
	}

	if got := baseProto.GetSettings().GetIpConfiguration().GetServerCaMode(); got != "CUSTOMER_MANAGED_CAS_CA" {
		t.Errorf("expected ServerCaMode CUSTOMER_MANAGED_CAS_CA, got %q", got)
	}
	if got := baseProto.GetSettings().GetIpConfiguration().GetServerCaPool(); got != "projects/test-project/locations/us-central1/caPools/pool-1" {
		t.Errorf("expected ServerCaPool pool-1, got %q", got)
	}
	if got := baseProto.GetSettings().GetIpConfiguration().GetCustomSubjectAlternativeNames(); !reflect.DeepEqual(got, []string{"db.example.com"}) {
		t.Errorf("expected CustomSubjectAlternativeNames [db.example.com], got %v", got)
	}

	cookie1, err := changecookies.ComputeChangeCookie(baseProto, baseProto)
	if err != nil {
		t.Fatalf("ComputeChangeCookie(base) failed: %v", err)
	}
	cookie2, err := changecookies.ComputeChangeCookie(updatedProto, baseProto)
	if err != nil {
		t.Fatalf("ComputeChangeCookie(updated) failed: %v", err)
	}
	if cookie1 == cookie2 {
		t.Errorf("expected change cookie to differ when serverCaPool and customSubjectAlternativeNames change, got identical cookie %q", cookie1)
	}
}

func TestDiffInstances_PscConfig_NetworkAttachmentUri_NoDiff(t *testing.T) {
	t.Parallel()
	desired := &api.DatabaseInstance{
		Settings: &api.Settings{
			IpConfiguration: &api.IpConfiguration{
				PscConfig: &api.PscConfig{
					PscEnabled:           true,
					NetworkAttachmentUri: "projects/test-project/regions/us-central1/networkAttachments/my-na",
				},
			},
		},
	}

	actual := &api.DatabaseInstance{
		Settings: &api.Settings{
			IpConfiguration: &api.IpConfiguration{
				PscConfig: &api.PscConfig{
					PscEnabled:           true,
					NetworkAttachmentUri: "https://www.googleapis.com/compute/v1/projects/test-project/regions/us-central1/networkAttachments/my-na",
				},
			},
		},
	}

	diff := DiffInstances(desired, actual)

	if diff.HasDiff() {
		t.Errorf("DiffInstances() expected no diffs for normalized networkAttachmentUri prefixes, but got: %v", diff.Fields)
	}
}

func TestDiffInstances_PscConfig_NetworkAttachmentUri_WithDiff(t *testing.T) {
	t.Parallel()
	desired := &api.DatabaseInstance{
		Settings: &api.Settings{
			IpConfiguration: &api.IpConfiguration{
				PscConfig: &api.PscConfig{
					PscEnabled:           true,
					NetworkAttachmentUri: "projects/test-project/regions/us-central1/networkAttachments/my-na-2",
				},
			},
		},
	}

	actual := &api.DatabaseInstance{
		Settings: &api.Settings{
			IpConfiguration: &api.IpConfiguration{
				PscConfig: &api.PscConfig{
					PscEnabled:           true,
					NetworkAttachmentUri: "projects/test-project/regions/us-central1/networkAttachments/my-na-1",
				},
			},
		},
	}

	diff := DiffInstances(desired, actual)

	if !diff.HasDiff() {
		t.Fatalf("DiffInstances() expected diff when networkAttachmentUri differs, got none")
	}
	if len(diff.Fields) != 1 {
		t.Fatalf("DiffInstances() expected 1 diff, got %d. Fields: %v", len(diff.Fields), diff.Fields)
	}
	if diff.Fields[0].ID != ".settings.ipConfiguration.pscConfig.networkAttachmentUri" {
		t.Errorf("expected diff ID .settings.ipConfiguration.pscConfig.networkAttachmentUri, got %q", diff.Fields[0].ID)
	}
}

func TestToProto_PreservesPscConfigNetworkAttachmentUri(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	adapter := &sqlInstanceAdapter{}

	base := &api.DatabaseInstance{
		Name: "my-instance",
		Settings: &api.Settings{
			IpConfiguration: &api.IpConfiguration{
				PscConfig: &api.PscConfig{
					PscEnabled:           true,
					NetworkAttachmentUri: "projects/test-project/regions/us-central1/networkAttachments/na-1",
				},
			},
		},
	}
	updated := &api.DatabaseInstance{
		Name: "my-instance",
		Settings: &api.Settings{
			IpConfiguration: &api.IpConfiguration{
				PscConfig: &api.PscConfig{
					PscEnabled:           true,
					NetworkAttachmentUri: "projects/test-project/regions/us-central1/networkAttachments/na-2",
				},
			},
		},
	}

	baseProto, err := adapter.toProto(ctx, base)
	if err != nil {
		t.Fatalf("toProto(base) failed: %v", err)
	}
	updatedProto, err := adapter.toProto(ctx, updated)
	if err != nil {
		t.Fatalf("toProto(updated) failed: %v", err)
	}

	if got := baseProto.GetSettings().GetIpConfiguration().GetPscConfig().GetNetworkAttachmentUri(); got != "projects/test-project/regions/us-central1/networkAttachments/na-1" {
		t.Errorf("expected NetworkAttachmentUri na-1, got %q", got)
	}

	cookie1, err := changecookies.ComputeChangeCookie(baseProto, baseProto)
	if err != nil {
		t.Fatalf("ComputeChangeCookie(base) failed: %v", err)
	}
	cookie2, err := changecookies.ComputeChangeCookie(updatedProto, baseProto)
	if err != nil {
		t.Fatalf("ComputeChangeCookie(updated) failed: %v", err)
	}
	if cookie1 == cookie2 {
		t.Errorf("expected change cookie to differ when networkAttachmentUri changes, got identical cookie %q", cookie1)
	}
}

func TestSQLInstanceKRMToGCP_PscConfig_CreateWithoutNetworkAttachment(t *testing.T) {
	t.Parallel()
	krmObj := &krm.SQLInstance{
		Spec: krm.SQLInstanceSpec{
			DatabaseVersion: direct.PtrTo("POSTGRES_16"),
			Region:          direct.PtrTo("us-central1"),
			Settings: krm.InstanceSettings{
				Tier: "db-custom-1-3840",
				IpConfiguration: &krm.InstanceIpConfiguration{
					Ipv4Enabled: direct.PtrTo(false),
					PscConfig: []krm.InstancePscConfig{
						{
							PscEnabled:              direct.PtrTo(true),
							AllowedConsumerProjects: []string{"test-project"},
						},
					},
				},
			},
		},
	}

	desiredGCP, err := SQLInstanceKRMToGCP(krmObj, nil, nil)
	if err != nil {
		t.Fatalf("SQLInstanceKRMToGCP() failed: %v", err)
	}

	psc := desiredGCP.Settings.IpConfiguration.PscConfig
	if psc == nil {
		t.Fatalf("expected non-nil PscConfig")
	}
	if psc.NetworkAttachmentUri != "" {
		t.Errorf("expected empty NetworkAttachmentUri on create without networkAttachmentRef, got %q", psc.NetworkAttachmentUri)
	}
	if len(psc.NullFields) != 0 {
		t.Errorf("expected empty NullFields on create without networkAttachmentRef, got %v", psc.NullFields)
	}
	for _, f := range psc.ForceSendFields {
		if f == "NetworkAttachmentUri" {
			t.Errorf("did not expect NetworkAttachmentUri in ForceSendFields when networkAttachmentRef is nil")
		}
	}

	rawJSON, err := json.Marshal(psc)
	if err != nil {
		t.Fatalf("json.Marshal(psc) failed: %v", err)
	}
	if strings.Contains(string(rawJSON), "networkAttachmentUri") {
		t.Errorf("expected create JSON to omit networkAttachmentUri completely, got %s", string(rawJSON))
	}
}

func TestSQLInstanceKRMToGCP_PscConfig_SetNetworkAttachment_PrefixNormalizationAndForceSend(t *testing.T) {
	t.Parallel()
	wantRelative := "projects/test-project/regions/us-central1/networkAttachments/my-na"
	inputs := []string{
		"projects/test-project/regions/us-central1/networkAttachments/my-na",
		"//compute.googleapis.com/projects/test-project/regions/us-central1/networkAttachments/my-na",
		"https://www.googleapis.com/compute/v1/projects/test-project/regions/us-central1/networkAttachments/my-na",
		"https://compute.googleapis.com/compute/v1/projects/test-project/regions/us-central1/networkAttachments/my-na",
		"https://compute.googleapis.com/v1/projects/test-project/regions/us-central1/networkAttachments/my-na",
		"https://compute.googleapis.com/projects/test-project/regions/us-central1/networkAttachments/my-na",
	}

	for _, external := range inputs {
		krmObj := &krm.SQLInstance{
			Spec: krm.SQLInstanceSpec{
				Settings: krm.InstanceSettings{
					Tier: "db-custom-1-3840",
					IpConfiguration: &krm.InstanceIpConfiguration{
						PscConfig: []krm.InstancePscConfig{
							{
								PscEnabled: direct.PtrTo(true),
								NetworkAttachmentRef: &computev1alpha1.ComputeNetworkAttachmentRef{
									External: external,
								},
							},
						},
					},
				},
			},
		}

		desiredGCP, err := SQLInstanceKRMToGCP(krmObj, nil, nil)
		if err != nil {
			t.Fatalf("SQLInstanceKRMToGCP(%q) failed: %v", external, err)
		}

		psc := desiredGCP.Settings.IpConfiguration.PscConfig
		if psc.NetworkAttachmentUri != wantRelative {
			t.Errorf("for input %q: expected NetworkAttachmentUri %q, got %q", external, wantRelative, psc.NetworkAttachmentUri)
		}
		hasForceSend := false
		for _, f := range psc.ForceSendFields {
			if f == "NetworkAttachmentUri" {
				hasForceSend = true
			}
		}
		if !hasForceSend {
			t.Errorf("for input %q: expected NetworkAttachmentUri in ForceSendFields, got %v", external, psc.ForceSendFields)
		}
	}
}

func TestSQLInstanceKRMToGCP_PscConfig_ClearNetworkAttachment_NullFieldsAndDiff(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	adapter := &sqlInstanceAdapter{}

	// User's desired KRM has pscConfig managed, with networkAttachmentRef removed/nil.
	krmWithoutAttachment := &krm.SQLInstance{
		Spec: krm.SQLInstanceSpec{
			DatabaseVersion: direct.PtrTo("POSTGRES_16"),
			Region:          direct.PtrTo("us-central1"),
			Settings: krm.InstanceSettings{
				Tier: "db-custom-1-3840",
				IpConfiguration: &krm.InstanceIpConfiguration{
					Ipv4Enabled: direct.PtrTo(false),
					PscConfig: []krm.InstancePscConfig{
						{
							PscEnabled:              direct.PtrTo(true),
							AllowedConsumerProjects: []string{"test-project"},
							NetworkAttachmentRef:    nil,
						},
					},
				},
			},
		},
	}

	// Live GCP currently has networkAttachmentUri attached.
	actualWithAttachment := &api.DatabaseInstance{
		DatabaseVersion: "POSTGRES_16",
		InstanceType:    "CLOUD_SQL_INSTANCE",
		Region:          "us-central1",
		Settings: &api.Settings{
			ActivationPolicy:          "ALWAYS",
			AuthorizedGaeApplications: []string{},
			AvailabilityType:          "ZONAL",
			BackupConfiguration:       &api.BackupConfiguration{},
			ConnectorEnforcement:      "NOT_REQUIRED",
			DataDiskType:              "PD_SSD",
			Edition:                   "ENTERPRISE",
			PricingPlan:               "PER_USE",
			ReplicationType:           "SYNCHRONOUS",
			StorageAutoResize:         direct.PtrTo(true),
			Tier:                      "db-custom-1-3840",
			UserLabels:                map[string]string{"managed-by-cnrm": "true"},
			IpConfiguration: &api.IpConfiguration{
				Ipv4Enabled: false,
				SslMode:     "ALLOW_UNENCRYPTED_AND_ENCRYPTED",
				PscConfig: &api.PscConfig{
					PscEnabled:              true,
					AllowedConsumerProjects: []string{"test-project"},
					NetworkAttachmentUri:    "projects/test-project/regions/us-central1/networkAttachments/na-1",
				},
			},
		},
	}

	// Step 1: Before clear, SQLInstanceKRMToGCP must populate NullFields: ["NetworkAttachmentUri"],
	// serialize `"networkAttachmentUri":null` in JSON, report a diff, and produce a distinct cookie.
	desiredClearGCP, err := SQLInstanceKRMToGCP(krmWithoutAttachment, actualWithAttachment, nil)
	if err != nil {
		t.Fatalf("SQLInstanceKRMToGCP(clear) failed: %v", err)
	}

	psc := desiredClearGCP.Settings.IpConfiguration.PscConfig
	if !reflect.DeepEqual(psc.NullFields, []string{"NetworkAttachmentUri"}) {
		t.Fatalf("expected NullFields [NetworkAttachmentUri] when detaching networkAttachmentUri, got %v", psc.NullFields)
	}

	rawJSON, err := json.Marshal(psc)
	if err != nil {
		t.Fatalf("json.Marshal(psc) failed: %v", err)
	}
	if !strings.Contains(string(rawJSON), `"networkAttachmentUri":null`) {
		t.Errorf("expected wire JSON to contain \"networkAttachmentUri\":null, got %s", string(rawJSON))
	}

	diff := DiffInstances(desiredClearGCP, actualWithAttachment)
	if !diff.HasDiff() {
		t.Fatalf("expected DiffInstances() to report diff when clearing networkAttachmentUri, got none")
	}
	if len(diff.Fields) != 1 || diff.Fields[0].ID != ".settings.ipConfiguration.pscConfig.networkAttachmentUri" {
		t.Errorf("expected single diff on .settings.ipConfiguration.pscConfig.networkAttachmentUri, got %v", diff.Fields)
	}

	desiredClearProto, err := adapter.toProto(ctx, desiredClearGCP)
	if err != nil {
		t.Fatalf("toProto(desiredClearGCP) failed: %v", err)
	}
	actualWithAttachmentProto, err := adapter.toProto(ctx, actualWithAttachment)
	if err != nil {
		t.Fatalf("toProto(actualWithAttachment) failed: %v", err)
	}
	previousCookie, err := changecookies.ComputeChangeCookie(actualWithAttachmentProto, actualWithAttachmentProto)
	if err != nil {
		t.Fatalf("ComputeChangeCookie(previous) failed: %v", err)
	}
	currentCookie, err := changecookies.ComputeChangeCookie(desiredClearProto, actualWithAttachmentProto)
	if err != nil {
		t.Fatalf("ComputeChangeCookie(current) failed: %v", err)
	}
	if previousCookie == currentCookie {
		t.Errorf("expected change cookie to differ when detaching networkAttachmentUri, got identical cookie %q", currentCookie)
	}

	// Step 2: After GCP detaches the network attachment (NetworkAttachmentUri == ""),
	// re-reconciliation must have empty NullFields, zero diff, and matching change cookies.
	actualAfterClear := &api.DatabaseInstance{
		DatabaseVersion: "POSTGRES_16",
		InstanceType:    "CLOUD_SQL_INSTANCE",
		Region:          "us-central1",
		Settings: &api.Settings{
			ActivationPolicy:          "ALWAYS",
			AuthorizedGaeApplications: []string{},
			AvailabilityType:          "ZONAL",
			BackupConfiguration:       &api.BackupConfiguration{},
			ConnectorEnforcement:      "NOT_REQUIRED",
			DataDiskType:              "PD_SSD",
			Edition:                   "ENTERPRISE",
			PricingPlan:               "PER_USE",
			ReplicationType:           "SYNCHRONOUS",
			StorageAutoResize:         direct.PtrTo(true),
			Tier:                      "db-custom-1-3840",
			UserLabels:                map[string]string{"managed-by-cnrm": "true"},
			IpConfiguration: &api.IpConfiguration{
				Ipv4Enabled: false,
				SslMode:     "ALLOW_UNENCRYPTED_AND_ENCRYPTED",
				PscConfig: &api.PscConfig{
					PscEnabled:              true,
					AllowedConsumerProjects: []string{"test-project"},
					NetworkAttachmentUri:    "",
				},
			},
		},
	}

	desiredSteadyGCP, err := SQLInstanceKRMToGCP(krmWithoutAttachment, actualAfterClear, nil)
	if err != nil {
		t.Fatalf("SQLInstanceKRMToGCP(steady) failed: %v", err)
	}
	if len(desiredSteadyGCP.Settings.IpConfiguration.PscConfig.NullFields) != 0 {
		t.Errorf("expected empty NullFields once GCP attachment is cleared, got %v", desiredSteadyGCP.Settings.IpConfiguration.PscConfig.NullFields)
	}
	steadyDiff := DiffInstances(desiredSteadyGCP, actualAfterClear)
	if steadyDiff.HasDiff() {
		t.Errorf("expected zero diff after networkAttachmentUri is cleared, got %v", steadyDiff.Fields)
	}
}

func TestInstancePscConfigGCPToKRM_NetworkAttachmentUri(t *testing.T) {
	t.Parallel()
	uri := "projects/test-project/regions/us-central1/networkAttachments/my-na"
	got := InstancePscConfigGCPToKRM(&api.PscConfig{
		PscEnabled:           true,
		NetworkAttachmentUri: uri,
	})
	if len(got) != 1 || got[0].NetworkAttachmentRef == nil || got[0].NetworkAttachmentRef.External != uri {
		t.Fatalf("expected NetworkAttachmentRef.External = %q, got %+v", uri, got)
	}

	gotEmpty := InstancePscConfigGCPToKRM(&api.PscConfig{
		PscEnabled:           true,
		NetworkAttachmentUri: "",
	})
	if len(gotEmpty) != 1 || gotEmpty[0].NetworkAttachmentRef != nil {
		t.Fatalf("expected nil NetworkAttachmentRef when NetworkAttachmentUri is empty, got %+v", gotEmpty)
	}
}

func TestSQLInstanceKRMToGCP_PscConfig_EdgeCases(t *testing.T) {
	t.Parallel()

	// 1. Empty External on NetworkAttachmentRef must not force-send "".
	pscEmptyExt := InstancePscConfigKRMToGCP([]krm.InstancePscConfig{
		{
			PscEnabled:           direct.PtrTo(true),
			NetworkAttachmentRef: &computev1alpha1.ComputeNetworkAttachmentRef{External: ""},
		},
	})
	if slices.Contains(pscEmptyExt.ForceSendFields, "NetworkAttachmentUri") {
		t.Errorf("expected NetworkAttachmentUri not to be force-sent when External is empty, got ForceSendFields=%v", pscEmptyExt.ForceSendFields)
	}

	// 2. Project-number URIs are preserved verbatim by Cloud SQL and match without diff.
	projNumURI := "projects/123456789012/regions/europe-north2/networkAttachments/my-na"
	diffProjNum := DiffPscConfig(
		&api.PscConfig{PscEnabled: true, NetworkAttachmentUri: "https://www.googleapis.com/compute/v1/" + projNumURI},
		&api.PscConfig{PscEnabled: true, NetworkAttachmentUri: projNumURI},
	)
	if diffProjNum.HasDiff() {
		t.Errorf("expected no diff for project-number networkAttachmentUri with URL prefix, got %v", diffProjNum.Fields)
	}

	// 3. Removing the entire pscConfig block from KRM sets desired PscConfig=nil and reports diffs
	// on pscEnabled, allowedConsumerProjects, and networkAttachmentUri rather than silently ignoring it.
	krmNoPscBlock := &krm.SQLInstance{
		Spec: krm.SQLInstanceSpec{
			DatabaseVersion: direct.PtrTo("POSTGRES_16"),
			Region:          direct.PtrTo("europe-north2"),
			Settings: krm.InstanceSettings{
				Tier: "db-custom-1-3840",
				IpConfiguration: &krm.InstanceIpConfiguration{
					Ipv4Enabled: direct.PtrTo(false),
				},
			},
		},
	}
	actualWithPsc := &api.DatabaseInstance{
		DatabaseVersion: "POSTGRES_16",
		InstanceType:    "CLOUD_SQL_INSTANCE",
		Region:          "europe-north2",
		Settings: &api.Settings{
			ActivationPolicy:          "ALWAYS",
			AuthorizedGaeApplications: []string{},
			AvailabilityType:          "ZONAL",
			BackupConfiguration:       &api.BackupConfiguration{},
			ConnectorEnforcement:      "NOT_REQUIRED",
			DataDiskType:              "PD_SSD",
			Edition:                   "ENTERPRISE",
			PricingPlan:               "PER_USE",
			ReplicationType:           "SYNCHRONOUS",
			StorageAutoResize:         direct.PtrTo(true),
			Tier:                      "db-custom-1-3840",
			UserLabels:                map[string]string{"managed-by-cnrm": "true"},
			IpConfiguration: &api.IpConfiguration{
				Ipv4Enabled: false,
				SslMode:     "ALLOW_UNENCRYPTED_AND_ENCRYPTED",
				PscConfig: &api.PscConfig{
					PscEnabled:              true,
					AllowedConsumerProjects: []string{"test-project"},
					NetworkAttachmentUri:    "projects/test-project/regions/europe-north2/networkAttachments/my-na",
				},
			},
		},
	}
	desiredNoPsc, err := SQLInstanceKRMToGCP(krmNoPscBlock, actualWithPsc, nil)
	if err != nil {
		t.Fatalf("SQLInstanceKRMToGCP(no pscConfig) failed: %v", err)
	}
	if desiredNoPsc.Settings.IpConfiguration.PscConfig != nil {
		t.Errorf("expected nil PscConfig when pscConfig block is omitted, got %+v", desiredNoPsc.Settings.IpConfiguration.PscConfig)
	}
	diffNoPsc := DiffInstances(desiredNoPsc, actualWithPsc)
	if !diffNoPsc.HasDiff() {
		t.Fatalf("expected diff when pscConfig block is removed from an active PSC instance")
	}
}
