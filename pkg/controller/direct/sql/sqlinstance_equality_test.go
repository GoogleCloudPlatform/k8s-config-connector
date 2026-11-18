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
	"reflect"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/changecookies"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	api "google.golang.org/api/sqladmin/v1beta4"
	"google.golang.org/protobuf/proto"

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

type mockOperation struct {
	cookieResult bool
}

func (m *mockOperation) UpdateStatus(ctx context.Context, typedStatus any, readyCondition *v1alpha1.Condition) error {
	return nil
}
func (m *mockOperation) RequestRequeue() {}
func (m *mockOperation) SetLastModifiedCookie(ctx context.Context, desired, actual proto.Message) error {
	return nil
}
func (m *mockOperation) CompareLastModifiedCookie(desired, actual proto.Message) (bool, error) {
	return m.cookieResult, nil
}
func (m *mockOperation) SetMBUR(ctx context.Context, computeMBUR func() (string, error)) error {
	return nil
}
func (m *mockOperation) CompareMBUR(computeMBUR func() (string, error)) (bool, error) {
	return false, nil
}

func TestCompareLastModifiedCookie_DatabaseVersionUnmodeledEnum(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	// 1. Verify that unmodeled database versions (e.g. POSTGRES_UNKNOWN_1 and POSTGRES_UNKNOWN_2)
	// both unmarshal to SQL_DATABASE_VERSION_UNSPECIFIED (0) in toProto,
	// demonstrating that changecookies proto hashing alone produces colliding cookies.
	adapter := &sqlInstanceAdapter{}
	settings := &api.Settings{
		Tier: "db-custom-2-7680",
		BackupConfiguration: &api.BackupConfiguration{
			Enabled: true,
		},
		IpConfiguration:    &api.IpConfiguration{},
		LocationPreference: &api.LocationPreference{},
	}
	unknown1GCP := &api.DatabaseInstance{
		Name:            "test-instance",
		DatabaseVersion: "POSTGRES_UNKNOWN_1",
		Settings:        settings,
	}
	unknown2GCP := &api.DatabaseInstance{
		Name:            "test-instance",
		DatabaseVersion: "POSTGRES_UNKNOWN_2",
		Settings:        settings,
	}

	unknown1Proto, err := adapter.toProto(ctx, unknown1GCP)
	if err != nil {
		t.Fatalf("toProto(unknown1) failed: %v", err)
	}
	unknown2Proto, err := adapter.toProto(ctx, unknown2GCP)
	if err != nil {
		t.Fatalf("toProto(unknown2) failed: %v", err)
	}

	cookie1, err := changecookies.ComputeChangeCookie(unknown1Proto, unknown1Proto)
	if err != nil {
		t.Fatalf("ComputeChangeCookie(unknown1) failed: %v", err)
	}
	cookie2, err := changecookies.ComputeChangeCookie(unknown2Proto, unknown1Proto)
	if err != nil {
		t.Fatalf("ComputeChangeCookie(unknown2, unknown1) failed: %v", err)
	}
	if cookie1 != cookie2 {
		t.Fatalf("expected colliding cookies due to unmodeled enum, but cookies differed: %q vs %q", cookie1, cookie2)
	}

	// 2. Verify that adapter.CompareLastModifiedCookie detects the version change
	// and returns false even when op.CompareLastModifiedCookie returns true.
	mockOp := &mockOperation{cookieResult: true}

	unknown2Desired := &krm.SQLInstance{
		Spec: krm.SQLInstanceSpec{
			DatabaseVersion: direct.LazyPtr("POSTGRES_UNKNOWN_2"),
			Settings: krm.InstanceSettings{
				Tier: "db-custom-2-7680",
			},
		},
	}
	adapter.desired = unknown2Desired

	// When actual is POSTGRES_UNKNOWN_1 and desired is POSTGRES_UNKNOWN_2, CompareLastModifiedCookie must return false.
	upToDate, err := adapter.CompareLastModifiedCookie(ctx, mockOp, unknown1GCP)
	if err != nil {
		t.Fatalf("CompareLastModifiedCookie failed: %v", err)
	}
	if upToDate {
		t.Errorf("expected CompareLastModifiedCookie to return false when databaseVersion differs (POSTGRES_UNKNOWN_1 vs POSTGRES_UNKNOWN_2), got true")
	}

	// When actual matches desired (POSTGRES_UNKNOWN_2 vs POSTGRES_UNKNOWN_2), CompareLastModifiedCookie must return true.
	upToDate, err = adapter.CompareLastModifiedCookie(ctx, mockOp, unknown2GCP)
	if err != nil {
		t.Fatalf("CompareLastModifiedCookie failed: %v", err)
	}
	if !upToDate {
		t.Errorf("expected CompareLastModifiedCookie to return true when databaseVersion matches, got false")
	}

	// When actual is POSTGRES_17 and desired is POSTGRES_18, CompareLastModifiedCookie must return false.
	pg17GCP := &api.DatabaseInstance{
		Name:            "test-instance",
		DatabaseVersion: "POSTGRES_17",
		Settings:        settings,
	}
	pg18GCP := &api.DatabaseInstance{
		Name:            "test-instance",
		DatabaseVersion: "POSTGRES_18",
		Settings:        settings,
	}
	pg18Desired := &krm.SQLInstance{
		Spec: krm.SQLInstanceSpec{
			DatabaseVersion: direct.LazyPtr("POSTGRES_18"),
			Settings: krm.InstanceSettings{
				Tier: "db-custom-2-7680",
			},
		},
	}
	adapter.desired = pg18Desired
	upToDate, err = adapter.CompareLastModifiedCookie(ctx, mockOp, pg17GCP)
	if err != nil {
		t.Fatalf("CompareLastModifiedCookie failed: %v", err)
	}
	if upToDate {
		t.Errorf("expected CompareLastModifiedCookie to return false when databaseVersion differs (POSTGRES_17 vs POSTGRES_18), got true")
	}
	upToDate, err = adapter.CompareLastModifiedCookie(ctx, mockOp, pg18GCP)
	if err != nil {
		t.Fatalf("CompareLastModifiedCookie failed: %v", err)
	}
	if !upToDate {
		t.Errorf("expected CompareLastModifiedCookie to return true when databaseVersion matches, got false")
	}

	// When actual is POSTGRES_16 (modeled) and desired is POSTGRES_17 (unmodeled), CompareLastModifiedCookie must return false.
	pg16GCP := &api.DatabaseInstance{
		Name:            "test-instance",
		DatabaseVersion: "POSTGRES_16",
		Settings:        settings,
	}
	pg17Desired := &krm.SQLInstance{
		Spec: krm.SQLInstanceSpec{
			DatabaseVersion: direct.LazyPtr("POSTGRES_17"),
			Settings: krm.InstanceSettings{
				Tier: "db-custom-2-7680",
			},
		},
	}
	adapter.desired = pg17Desired
	upToDate, err = adapter.CompareLastModifiedCookie(ctx, mockOp, pg16GCP)
	if err != nil {
		t.Fatalf("CompareLastModifiedCookie failed: %v", err)
	}
	if upToDate {
		t.Errorf("expected CompareLastModifiedCookie to return false when databaseVersion differs (POSTGRES_16 vs POSTGRES_17), got true")
	}

	// When desired.Spec.DatabaseVersion is nil, CompareLastModifiedCookie should defer to cookie comparison.
	nilVersionDesired := &krm.SQLInstance{
		Spec: krm.SQLInstanceSpec{
			Settings: krm.InstanceSettings{
				Tier: "db-custom-2-7680",
			},
		},
	}
	adapter.desired = nilVersionDesired
	upToDate, err = adapter.CompareLastModifiedCookie(ctx, mockOp, pg18GCP)
	if err != nil {
		t.Fatalf("CompareLastModifiedCookie failed: %v", err)
	}
	if !upToDate {
		t.Errorf("expected CompareLastModifiedCookie to return true when databaseVersion is nil and cookie matches, got false")
	}
}
