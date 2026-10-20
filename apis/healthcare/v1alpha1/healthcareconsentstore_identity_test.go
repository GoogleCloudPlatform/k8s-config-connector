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

	"github.com/google/go-cmp/cmp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
}
