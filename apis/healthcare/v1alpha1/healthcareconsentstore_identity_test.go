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

package v1alpha1

import (
	"context"
	"testing"

	k8sv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	"github.com/google/go-cmp/cmp"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestHealthcareConsentStoreIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *HealthcareConsentStoreIdentity
	}{
		{
			name: "valid relative reference",
			ref:  "projects/my-project/locations/us-central1/datasets/my-dataset/consentStores/my-consent-store",
			want: &HealthcareConsentStoreIdentity{
				Project:      "my-project",
				Location:     "us-central1",
				Dataset:      "my-dataset",
				ConsentStore: "my-consent-store",
			},
		},
		{
			name: "valid full url with https",
			ref:  "https://healthcare.googleapis.com/projects/my-project/locations/us-central1/datasets/my-dataset/consentStores/my-consent-store",
			want: &HealthcareConsentStoreIdentity{
				Project:      "my-project",
				Location:     "us-central1",
				Dataset:      "my-dataset",
				ConsentStore: "my-consent-store",
			},
		},
		{
			name: "valid full url with //",
			ref:  "//healthcare.googleapis.com/projects/my-project/locations/us-central1/datasets/my-dataset/consentStores/my-consent-store",
			want: &HealthcareConsentStoreIdentity{
				Project:      "my-project",
				Location:     "us-central1",
				Dataset:      "my-dataset",
				ConsentStore: "my-consent-store",
			},
		},
		{
			name:    "invalid format missing consentStores",
			ref:     "projects/my-project/locations/us-central1/datasets/my-dataset",
			wantErr: true,
		},
		{
			name:    "empty string",
			ref:     "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &HealthcareConsentStoreIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Fatalf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				if diff := cmp.Diff(tt.want, i); diff != "" {
					t.Errorf("FromExternal() mismatch (-want +got):\n%s", diff)
				}
				expectedString := "projects/my-project/locations/us-central1/datasets/my-dataset/consentStores/my-consent-store"
				if got := i.String(); got != expectedString {
					t.Errorf("String() = %v, want %v", got, expectedString)
				}
				expectedParent := "projects/my-project/locations/us-central1/datasets/my-dataset"
				if got := i.ParentString(); got != expectedParent {
					t.Errorf("ParentString() = %v, want %v", got, expectedParent)
				}
				if gotHost := i.Host(); gotHost != "healthcare.googleapis.com" {
					t.Errorf("Host() = %v, want %v", gotHost, "healthcare.googleapis.com")
				}
			}
		})
	}
}

func TestHealthcareDatasetIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *HealthcareDatasetIdentity
	}{
		{
			name: "valid relative reference",
			ref:  "projects/my-project/locations/us-central1/datasets/my-dataset",
			want: &HealthcareDatasetIdentity{
				Project:  "my-project",
				Location: "us-central1",
				Dataset:  "my-dataset",
			},
		},
		{
			name: "valid full url with https",
			ref:  "https://healthcare.googleapis.com/projects/my-project/locations/us-central1/datasets/my-dataset",
			want: &HealthcareDatasetIdentity{
				Project:  "my-project",
				Location: "us-central1",
				Dataset:  "my-dataset",
			},
		},
		{
			name: "valid full url with //",
			ref:  "//healthcare.googleapis.com/projects/my-project/locations/us-central1/datasets/my-dataset",
			want: &HealthcareDatasetIdentity{
				Project:  "my-project",
				Location: "us-central1",
				Dataset:  "my-dataset",
			},
		},
		{
			name:    "invalid format missing dataset",
			ref:     "projects/my-project/locations/us-central1",
			wantErr: true,
		},
		{
			name:    "empty string",
			ref:     "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &HealthcareDatasetIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Fatalf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				if diff := cmp.Diff(tt.want, i); diff != "" {
					t.Errorf("FromExternal() mismatch (-want +got):\n%s", diff)
				}
				expectedString := "projects/my-project/locations/us-central1/datasets/my-dataset"
				if got := i.String(); got != expectedString {
					t.Errorf("String() = %v, want %v", got, expectedString)
				}
			}
		})
	}
}

func TestGetIdentityFromHealthcareConsentStoreSpec(t *testing.T) {
	ctx := context.Background()

	resourceID := "custom-store-id"
	obj := &HealthcareConsentStore{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-store",
			Namespace: "test-ns",
		},
		Spec: HealthcareConsentStoreSpec{
			Dataset:    "projects/test-project/locations/us-central1/datasets/test-dataset",
			ResourceID: &resourceID,
		},
	}

	got, err := getIdentityFromHealthcareConsentStoreSpec(ctx, nil, obj)
	if err != nil {
		t.Fatalf("getIdentityFromHealthcareConsentStoreSpec() unexpected error: %v", err)
	}

	want := &HealthcareConsentStoreIdentity{
		Project:      "test-project",
		Location:     "us-central1",
		Dataset:      "test-dataset",
		ConsentStore: "custom-store-id",
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("getIdentityFromHealthcareConsentStoreSpec() mismatch (-want +got):\n%s", diff)
	}

	// Test default to metadata.name when resourceID is not set
	objNoResourceID := &HealthcareConsentStore{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-store",
			Namespace: "test-ns",
		},
		Spec: HealthcareConsentStoreSpec{
			Dataset: "projects/test-project/locations/us-central1/datasets/test-dataset",
		},
	}

	gotDefault, err := getIdentityFromHealthcareConsentStoreSpec(ctx, nil, objNoResourceID)
	if err != nil {
		t.Fatalf("getIdentityFromHealthcareConsentStoreSpec() unexpected error: %v", err)
	}

	wantDefault := &HealthcareConsentStoreIdentity{
		Project:      "test-project",
		Location:     "us-central1",
		Dataset:      "test-dataset",
		ConsentStore: "test-store",
	}

	if diff := cmp.Diff(wantDefault, gotDefault); diff != "" {
		t.Errorf("getIdentityFromHealthcareConsentStoreSpec() mismatch (-want +got):\n%s", diff)
	}

	// Test missing dataset
	objMissingDataset := &HealthcareConsentStore{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-store",
			Namespace: "test-ns",
		},
		Spec: HealthcareConsentStoreSpec{},
	}
	if _, err := getIdentityFromHealthcareConsentStoreSpec(ctx, nil, objMissingDataset); err == nil {
		t.Errorf("getIdentityFromHealthcareConsentStoreSpec() expected error for missing dataset, got nil")
	}

	// Test invalid dataset format
	objInvalidDataset := &HealthcareConsentStore{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-store",
			Namespace: "test-ns",
		},
		Spec: HealthcareConsentStoreSpec{
			Dataset: "invalid/dataset",
		},
	}
	if _, err := getIdentityFromHealthcareConsentStoreSpec(ctx, nil, objInvalidDataset); err == nil {
		t.Errorf("getIdentityFromHealthcareConsentStoreSpec() expected error for invalid dataset, got nil")
	}
}

func TestHealthcareConsentStore_GetIdentity(t *testing.T) {
	ctx := context.Background()

	obj := &HealthcareConsentStore{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-consent-store",
			Namespace: "default",
		},
		Spec: HealthcareConsentStoreSpec{
			Dataset: "projects/test-project/locations/us-central1/datasets/test-dataset",
		},
	}

	gotIdentity, err := obj.GetIdentity(ctx, nil)
	if err != nil {
		t.Fatalf("GetIdentity() error = %v", err)
	}

	got, ok := gotIdentity.(*HealthcareConsentStoreIdentity)
	if !ok {
		t.Fatalf("GetIdentity() returned %T, want *HealthcareConsentStoreIdentity", gotIdentity)
	}

	want := &HealthcareConsentStoreIdentity{
		Project:      "test-project",
		Location:     "us-central1",
		Dataset:      "test-dataset",
		ConsentStore: "my-consent-store",
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("GetIdentity() mismatch (-want +got):\n%s", diff)
	}
}

func TestHealthcareConsentStoreRef_Normalize(t *testing.T) {
	ctx := context.Background()

	scheme := runtime.NewScheme()
	_ = AddToScheme(scheme)
	_ = corev1.AddToScheme(scheme)

	readyStore := &HealthcareConsentStore{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "ready-store",
			Namespace: "test-ns",
		},
		Spec: HealthcareConsentStoreSpec{
			Dataset: "projects/test-project/locations/us-central1/datasets/test-dataset",
		},
		Status: HealthcareConsentStoreStatus{
			Conditions: []k8sv1alpha1.Condition{
				{
					Type:   "Ready",
					Status: "True",
				},
			},
		},
	}

	notReadyStore := &HealthcareConsentStore{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "not-ready-store",
			Namespace: "test-ns",
		},
		Spec: HealthcareConsentStoreSpec{
			Dataset: "projects/test-project/locations/us-central1/datasets/test-dataset",
		},
		Status: HealthcareConsentStoreStatus{
			Conditions: []k8sv1alpha1.Condition{
				{
					Type:   "Ready",
					Status: "False",
				},
			},
		},
	}

	fakeClient := fake.NewClientBuilder().WithScheme(scheme).WithObjects(readyStore, notReadyStore).Build()

	tests := []struct {
		name             string
		ref              *HealthcareConsentStoreRef
		defaultNamespace string
		wantErr          bool
		wantExternal     string
	}{
		{
			name: "external reference valid",
			ref: &HealthcareConsentStoreRef{
				External: "projects/p1/locations/l1/datasets/d1/consentStores/cs1",
			},
			defaultNamespace: "test-ns",
			wantErr:          false,
			wantExternal:     "projects/p1/locations/l1/datasets/d1/consentStores/cs1",
		},
		{
			name: "external reference invalid",
			ref: &HealthcareConsentStoreRef{
				External: "invalid/format",
			},
			defaultNamespace: "test-ns",
			wantErr:          true,
		},
		{
			name: "by name and namespace when ready",
			ref: &HealthcareConsentStoreRef{
				Name:      "ready-store",
				Namespace: "test-ns",
			},
			defaultNamespace: "test-ns",
			wantErr:          false,
			wantExternal:     "projects/test-project/locations/us-central1/datasets/test-dataset/consentStores/ready-store",
		},
		{
			name: "by name with default namespace when ready",
			ref: &HealthcareConsentStoreRef{
				Name: "ready-store",
			},
			defaultNamespace: "test-ns",
			wantErr:          false,
			wantExternal:     "projects/test-project/locations/us-central1/datasets/test-dataset/consentStores/ready-store",
		},
		{
			name: "by name and namespace when not ready",
			ref: &HealthcareConsentStoreRef{
				Name:      "not-ready-store",
				Namespace: "test-ns",
			},
			defaultNamespace: "test-ns",
			wantErr:          true,
		},
		{
			name: "by name when not found",
			ref: &HealthcareConsentStoreRef{
				Name:      "non-existent",
				Namespace: "test-ns",
			},
			defaultNamespace: "test-ns",
			wantErr:          true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.ref.Normalize(ctx, fakeClient, tt.defaultNamespace)
			if (err != nil) != tt.wantErr {
				t.Fatalf("Normalize() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				if tt.ref.External != tt.wantExternal {
					t.Errorf("Normalize() ref.External = %v, want %v", tt.ref.External, tt.wantExternal)
				}
				if tt.ref.Name != "" || tt.ref.Namespace != "" {
					t.Errorf("Normalize() Name/Namespace not cleared: name=%q, ns=%q", tt.ref.Name, tt.ref.Namespace)
				}
			}
		})
	}
}

func TestHealthcareConsentStoreRef_Interface(t *testing.T) {
	ref := &HealthcareConsentStoreRef{
		Name:      "test-store",
		Namespace: "test-ns",
	}

	if got := ref.GetGVK(); got != HealthcareConsentStoreGVK {
		t.Errorf("GetGVK() = %v, want %v", got, HealthcareConsentStoreGVK)
	}

	wantNN := types.NamespacedName{
		Name:      "test-store",
		Namespace: "test-ns",
	}
	if gotNN := ref.GetNamespacedName(); gotNN != wantNN {
		t.Errorf("GetNamespacedName() = %v, want %v", gotNN, wantNN)
	}

	externalVal := "projects/p1/locations/l1/datasets/d1/consentStores/cs1"
	ref.SetExternal(externalVal)
	if got := ref.GetExternal(); got != externalVal {
		t.Errorf("GetExternal() = %v, want %v", got, externalVal)
	}
	if ref.Name != "" || ref.Namespace != "" {
		t.Errorf("SetExternal() did not clear Name/Namespace")
	}

	if err := ref.ValidateExternal(externalVal); err != nil {
		t.Errorf("ValidateExternal(%q) unexpected error: %v", externalVal, err)
	}
	if err := ref.ValidateExternal("invalid"); err == nil {
		t.Errorf("ValidateExternal(invalid) expected error, got nil")
	}

	id, err := ref.ParseExternalToIdentity()
	if err != nil {
		t.Fatalf("ParseExternalToIdentity() error = %v", err)
	}
	storeID, ok := id.(*HealthcareConsentStoreIdentity)
	if !ok {
		t.Fatalf("ParseExternalToIdentity() returned %T, want *HealthcareConsentStoreIdentity", id)
	}
	wantID := &HealthcareConsentStoreIdentity{
		Project:      "p1",
		Location:     "l1",
		Dataset:      "d1",
		ConsentStore: "cs1",
	}
	if diff := cmp.Diff(wantID, storeID); diff != "" {
		t.Errorf("ParseExternalToIdentity() mismatch (-want +got):\n%s", diff)
	}
}
