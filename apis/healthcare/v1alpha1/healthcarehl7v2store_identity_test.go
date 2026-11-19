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

func TestHealthcareHL7V2StoreIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *HealthcareHL7V2StoreIdentity
	}{
		{
			name: "valid relative reference",
			ref:  "projects/my-project/locations/us-central1/datasets/my-dataset/hl7V2Stores/my-hl7v2-store",
			want: &HealthcareHL7V2StoreIdentity{
				Project:    "my-project",
				Location:   "us-central1",
				Dataset:    "my-dataset",
				Hl7V2Store: "my-hl7v2-store",
			},
		},
		{
			name: "valid full url with https",
			ref:  "https://healthcare.googleapis.com/projects/my-project/locations/us-central1/datasets/my-dataset/hl7V2Stores/my-hl7v2-store",
			want: &HealthcareHL7V2StoreIdentity{
				Project:    "my-project",
				Location:   "us-central1",
				Dataset:    "my-dataset",
				Hl7V2Store: "my-hl7v2-store",
			},
		},
		{
			name: "valid full url with //",
			ref:  "//healthcare.googleapis.com/projects/my-project/locations/us-central1/datasets/my-dataset/hl7V2Stores/my-hl7v2-store",
			want: &HealthcareHL7V2StoreIdentity{
				Project:    "my-project",
				Location:   "us-central1",
				Dataset:    "my-dataset",
				Hl7V2Store: "my-hl7v2-store",
			},
		},
		{
			name:    "invalid format missing hl7V2Stores",
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
			i := &HealthcareHL7V2StoreIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Fatalf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				if diff := cmp.Diff(tt.want, i); diff != "" {
					t.Errorf("FromExternal() mismatch (-want +got):\n%s", diff)
				}
				expectedString := "projects/my-project/locations/us-central1/datasets/my-dataset/hl7V2Stores/my-hl7v2-store"
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

func TestGetIdentityFromHealthcareHL7V2StoreSpec(t *testing.T) {
	ctx := context.Background()

	resourceID := "custom-store-id"
	obj := &HealthcareHL7V2Store{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-store",
			Namespace: "test-ns",
		},
		Spec: HealthcareHL7V2StoreSpec{
			Dataset:    "projects/test-project/locations/us-central1/datasets/test-dataset",
			ResourceID: &resourceID,
		},
	}

	got, err := getIdentityFromHealthcareHL7V2StoreSpec(ctx, nil, obj)
	if err != nil {
		t.Fatalf("getIdentityFromHealthcareHL7V2StoreSpec() unexpected error: %v", err)
	}

	want := &HealthcareHL7V2StoreIdentity{
		Project:    "test-project",
		Location:   "us-central1",
		Dataset:    "test-dataset",
		Hl7V2Store: "custom-store-id",
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("getIdentityFromHealthcareHL7V2StoreSpec() mismatch (-want +got):\n%s", diff)
	}

	// Test default to metadata.name when resourceID is not set
	objNoResourceID := &HealthcareHL7V2Store{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-store",
			Namespace: "test-ns",
		},
		Spec: HealthcareHL7V2StoreSpec{
			Dataset: "projects/test-project/locations/us-central1/datasets/test-dataset",
		},
	}

	gotDefault, err := getIdentityFromHealthcareHL7V2StoreSpec(ctx, nil, objNoResourceID)
	if err != nil {
		t.Fatalf("getIdentityFromHealthcareHL7V2StoreSpec() unexpected error: %v", err)
	}

	wantDefault := &HealthcareHL7V2StoreIdentity{
		Project:    "test-project",
		Location:   "us-central1",
		Dataset:    "test-dataset",
		Hl7V2Store: "test-store",
	}

	if diff := cmp.Diff(wantDefault, gotDefault); diff != "" {
		t.Errorf("getIdentityFromHealthcareHL7V2StoreSpec() mismatch (-want +got):\n%s", diff)
	}

	// Test missing dataset
	objMissingDataset := &HealthcareHL7V2Store{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-store",
			Namespace: "test-ns",
		},
		Spec: HealthcareHL7V2StoreSpec{},
	}
	if _, err := getIdentityFromHealthcareHL7V2StoreSpec(ctx, nil, objMissingDataset); err == nil {
		t.Errorf("getIdentityFromHealthcareHL7V2StoreSpec() expected error for missing dataset, got nil")
	}

	// Test invalid dataset format
	objInvalidDataset := &HealthcareHL7V2Store{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-store",
			Namespace: "test-ns",
		},
		Spec: HealthcareHL7V2StoreSpec{
			Dataset: "invalid/dataset",
		},
	}
	if _, err := getIdentityFromHealthcareHL7V2StoreSpec(ctx, nil, objInvalidDataset); err == nil {
		t.Errorf("getIdentityFromHealthcareHL7V2StoreSpec() expected error for invalid dataset, got nil")
	}
}

func TestHealthcareHL7V2Store_GetIdentity(t *testing.T) {
	ctx := context.Background()

	obj := &HealthcareHL7V2Store{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-hl7v2-store",
			Namespace: "default",
		},
		Spec: HealthcareHL7V2StoreSpec{
			Dataset: "projects/test-project/locations/us-central1/datasets/test-dataset",
		},
	}

	gotIdentity, err := obj.GetIdentity(ctx, nil)
	if err != nil {
		t.Fatalf("GetIdentity() error = %v", err)
	}

	got, ok := gotIdentity.(*HealthcareHL7V2StoreIdentity)
	if !ok {
		t.Fatalf("GetIdentity() returned %T, want *HealthcareHL7V2StoreIdentity", gotIdentity)
	}

	want := &HealthcareHL7V2StoreIdentity{
		Project:    "test-project",
		Location:   "us-central1",
		Dataset:    "test-dataset",
		Hl7V2Store: "my-hl7v2-store",
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("GetIdentity() mismatch (-want +got):\n%s", diff)
	}
}

func TestHealthcareHL7V2StoreRef_Normalize(t *testing.T) {
	ctx := context.Background()

	scheme := runtime.NewScheme()
	_ = AddToScheme(scheme)
	_ = corev1.AddToScheme(scheme)

	readyStore := &HealthcareHL7V2Store{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "ready-store",
			Namespace: "test-ns",
		},
		Spec: HealthcareHL7V2StoreSpec{
			Dataset: "projects/test-project/locations/us-central1/datasets/test-dataset",
		},
		Status: HealthcareHL7V2StoreStatus{
			Conditions: []k8sv1alpha1.Condition{
				{
					Type:   "Ready",
					Status: "True",
				},
			},
		},
	}

	notReadyStore := &HealthcareHL7V2Store{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "not-ready-store",
			Namespace: "test-ns",
		},
		Spec: HealthcareHL7V2StoreSpec{
			Dataset: "projects/test-project/locations/us-central1/datasets/test-dataset",
		},
		Status: HealthcareHL7V2StoreStatus{
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
		ref              *HealthcareHL7V2StoreRef
		defaultNamespace string
		wantErr          bool
		wantExternal     string
	}{
		{
			name: "external reference valid",
			ref: &HealthcareHL7V2StoreRef{
				External: "projects/p1/locations/l1/datasets/d1/hl7V2Stores/hs1",
			},
			defaultNamespace: "test-ns",
			wantErr:          false,
			wantExternal:     "projects/p1/locations/l1/datasets/d1/hl7V2Stores/hs1",
		},
		{
			name: "external reference invalid",
			ref: &HealthcareHL7V2StoreRef{
				External: "invalid/format",
			},
			defaultNamespace: "test-ns",
			wantErr:          true,
		},
		{
			name: "by name and namespace when ready",
			ref: &HealthcareHL7V2StoreRef{
				Name:      "ready-store",
				Namespace: "test-ns",
			},
			defaultNamespace: "test-ns",
			wantErr:          false,
			wantExternal:     "projects/test-project/locations/us-central1/datasets/test-dataset/hl7V2Stores/ready-store",
		},
		{
			name: "by name with default namespace when ready",
			ref: &HealthcareHL7V2StoreRef{
				Name: "ready-store",
			},
			defaultNamespace: "test-ns",
			wantErr:          false,
			wantExternal:     "projects/test-project/locations/us-central1/datasets/test-dataset/hl7V2Stores/ready-store",
		},
		{
			name: "by name and namespace when not ready",
			ref: &HealthcareHL7V2StoreRef{
				Name:      "not-ready-store",
				Namespace: "test-ns",
			},
			defaultNamespace: "test-ns",
			wantErr:          true,
		},
		{
			name: "by name when not found",
			ref: &HealthcareHL7V2StoreRef{
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

func TestHealthcareHL7V2StoreRef_Interface(t *testing.T) {
	ref := &HealthcareHL7V2StoreRef{
		Name:      "test-store",
		Namespace: "test-ns",
	}

	if got := ref.GetGVK(); got != HealthcareHL7V2StoreGVK {
		t.Errorf("GetGVK() = %v, want %v", got, HealthcareHL7V2StoreGVK)
	}

	wantNN := types.NamespacedName{
		Name:      "test-store",
		Namespace: "test-ns",
	}
	if gotNN := ref.GetNamespacedName(); gotNN != wantNN {
		t.Errorf("GetNamespacedName() = %v, want %v", gotNN, wantNN)
	}

	externalVal := "projects/p1/locations/l1/datasets/d1/hl7V2Stores/hs1"
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
	storeID, ok := id.(*HealthcareHL7V2StoreIdentity)
	if !ok {
		t.Fatalf("ParseExternalToIdentity() returned %T, want *HealthcareHL7V2StoreIdentity", id)
	}
	wantID := &HealthcareHL7V2StoreIdentity{
		Project:    "p1",
		Location:   "l1",
		Dataset:    "d1",
		Hl7V2Store: "hs1",
	}
	if diff := cmp.Diff(wantID, storeID); diff != "" {
		t.Errorf("ParseExternalToIdentity() mismatch (-want +got):\n%s", diff)
	}
}
