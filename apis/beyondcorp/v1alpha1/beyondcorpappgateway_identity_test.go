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

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"
	k8sv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	"github.com/google/go-cmp/cmp"
)

func TestBeyondCorpAppGatewayIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name         string
		ref          string
		wantProject  string
		wantLocation string
		wantID       string
		wantErr      bool
	}{
		{
			name:         "valid external ref",
			ref:          "projects/my-project/locations/us-central1/appGateways/my-gateway",
			wantProject:  "my-project",
			wantLocation: "us-central1",
			wantID:       "my-gateway",
			wantErr:      false,
		},
		{
			name:         "valid full url with //",
			ref:          "//beyondcorp.googleapis.com/projects/my-project/locations/us-central1/appGateways/my-gateway",
			wantProject:  "my-project",
			wantLocation: "us-central1",
			wantID:       "my-gateway",
			wantErr:      false,
		},
		{
			name:    "invalid external ref missing appGateways",
			ref:     "projects/my-project/locations/us-central1/my-gateway",
			wantErr: true,
		},
		{
			name:    "invalid empty ref",
			ref:     "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BeyondCorpAppGatewayIdentity{}
			if err := i.FromExternal(tt.ref); (err != nil) != tt.wantErr {
				t.Fatalf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				want := &BeyondCorpAppGatewayIdentity{
					Project:    tt.wantProject,
					Location:   tt.wantLocation,
					AppGateway: tt.wantID,
				}
				if diff := cmp.Diff(want, i); diff != "" {
					t.Errorf("FromExternal() mismatch (-want +got):\n%s", diff)
				}
				canonicalRef := "projects/" + tt.wantProject + "/locations/" + tt.wantLocation + "/appGateways/" + tt.wantID
				if got := i.String(); got != canonicalRef {
					t.Errorf("String() = %v, want %v", got, canonicalRef)
				}
				wantParent := "projects/" + tt.wantProject + "/locations/" + tt.wantLocation
				if got := i.ParentString(); got != wantParent {
					t.Errorf("ParentString() = %v, want %v", got, wantParent)
				}
				if gotHost := i.Host(); gotHost != "beyondcorp.googleapis.com" {
					t.Errorf("Host() = %v, want %v", gotHost, "beyondcorp.googleapis.com")
				}
			}
		})
	}
}

func TestGetIdentityFromBeyondCorpAppGatewaySpec(t *testing.T) {
	ctx := context.Background()

	resourceID := "custom-gateway-id"
	obj := &BeyondCorpAppGateway{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-gateway",
			Namespace: "test-ns",
		},
		Spec: BeyondCorpAppGatewaySpec{
			ProjectRef: refs.ProjectRef{
				External: "test-project",
			},
			Region:     "us-central1",
			ResourceID: &resourceID,
		},
	}

	got, err := getIdentityFromBeyondCorpAppGatewaySpec(ctx, nil, obj)
	if err != nil {
		t.Fatalf("getIdentityFromBeyondCorpAppGatewaySpec() unexpected error: %v", err)
	}

	want := &BeyondCorpAppGatewayIdentity{
		Project:    "test-project",
		Location:   "us-central1",
		AppGateway: "custom-gateway-id",
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("getIdentityFromBeyondCorpAppGatewaySpec() mismatch (-want +got):\n%s", diff)
	}

	// Test default to metadata.name when resourceID is not set
	objNoResourceID := &BeyondCorpAppGateway{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-gateway",
			Namespace: "test-ns",
		},
		Spec: BeyondCorpAppGatewaySpec{
			ProjectRef: refs.ProjectRef{
				External: "test-project",
			},
			Region: "us-central1",
		},
	}

	gotDefault, err := getIdentityFromBeyondCorpAppGatewaySpec(ctx, nil, objNoResourceID)
	if err != nil {
		t.Fatalf("getIdentityFromBeyondCorpAppGatewaySpec() unexpected error: %v", err)
	}

	wantDefault := &BeyondCorpAppGatewayIdentity{
		Project:    "test-project",
		Location:   "us-central1",
		AppGateway: "test-gateway",
	}

	if diff := cmp.Diff(wantDefault, gotDefault); diff != "" {
		t.Errorf("getIdentityFromBeyondCorpAppGatewaySpec() mismatch (-want +got):\n%s", diff)
	}

	// Test missing region
	objMissingRegion := &BeyondCorpAppGateway{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-gateway",
			Namespace: "test-ns",
		},
		Spec: BeyondCorpAppGatewaySpec{
			ProjectRef: refs.ProjectRef{
				External: "test-project",
			},
		},
	}
	if _, err := getIdentityFromBeyondCorpAppGatewaySpec(ctx, nil, objMissingRegion); err == nil {
		t.Errorf("getIdentityFromBeyondCorpAppGatewaySpec() expected error for missing region, got nil")
	}
}

func TestBeyondCorpAppGateway_GetIdentity(t *testing.T) {
	ctx := context.Background()

	obj := &BeyondCorpAppGateway{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-gateway",
			Namespace: "default",
		},
		Spec: BeyondCorpAppGatewaySpec{
			ProjectRef: refs.ProjectRef{
				External: "test-project",
			},
			Region: "us-central1",
		},
	}

	gotIdentity, err := obj.GetIdentity(ctx, nil)
	if err != nil {
		t.Fatalf("GetIdentity() error = %v", err)
	}

	got, ok := gotIdentity.(*BeyondCorpAppGatewayIdentity)
	if !ok {
		t.Fatalf("GetIdentity() returned %T, want *BeyondCorpAppGatewayIdentity", gotIdentity)
	}

	want := &BeyondCorpAppGatewayIdentity{
		Project:    "test-project",
		Location:   "us-central1",
		AppGateway: "my-gateway",
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("GetIdentity() mismatch (-want +got):\n%s", diff)
	}
}

func TestBeyondCorpAppGatewayRef_Normalize(t *testing.T) {
	ctx := context.Background()

	scheme := runtime.NewScheme()
	_ = AddToScheme(scheme)
	_ = corev1.AddToScheme(scheme)

	readyExternalRef := "projects/test-project/locations/us-central1/appGateways/ready-gateway"
	readyGateway := &BeyondCorpAppGateway{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "ready-gateway",
			Namespace: "test-ns",
		},
		Spec: BeyondCorpAppGatewaySpec{
			ProjectRef: refs.ProjectRef{
				External: "test-project",
			},
			Region: "us-central1",
		},
		Status: BeyondCorpAppGatewayStatus{
			ExternalRef: &readyExternalRef,
			Conditions: []k8sv1alpha1.Condition{
				{
					Type:   "Ready",
					Status: "True",
				},
			},
		},
	}

	notReadyGateway := &BeyondCorpAppGateway{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "not-ready-gateway",
			Namespace: "test-ns",
		},
		Spec: BeyondCorpAppGatewaySpec{
			ProjectRef: refs.ProjectRef{
				External: "test-project",
			},
			Region: "us-central1",
		},
		Status: BeyondCorpAppGatewayStatus{
			Conditions: []k8sv1alpha1.Condition{
				{
					Type:   "Ready",
					Status: "False",
				},
			},
		},
	}

	fakeClient := fake.NewClientBuilder().WithScheme(scheme).WithObjects(readyGateway, notReadyGateway).Build()

	tests := []struct {
		name             string
		ref              *BeyondCorpAppGatewayRef
		defaultNamespace string
		wantErr          bool
		wantExternal     string
	}{
		{
			name: "external reference valid",
			ref: &BeyondCorpAppGatewayRef{
				External: "projects/p1/locations/us-central1/appGateways/gw1",
			},
			defaultNamespace: "test-ns",
			wantErr:          false,
			wantExternal:     "projects/p1/locations/us-central1/appGateways/gw1",
		},
		{
			name: "external reference invalid",
			ref: &BeyondCorpAppGatewayRef{
				External: "invalid/format",
			},
			defaultNamespace: "test-ns",
			wantErr:          true,
		},
		{
			name: "by name and namespace when ready",
			ref: &BeyondCorpAppGatewayRef{
				Name:      "ready-gateway",
				Namespace: "test-ns",
			},
			defaultNamespace: "test-ns",
			wantErr:          false,
			wantExternal:     "projects/test-project/locations/us-central1/appGateways/ready-gateway",
		},
		{
			name: "by name with default namespace when ready",
			ref: &BeyondCorpAppGatewayRef{
				Name: "ready-gateway",
			},
			defaultNamespace: "test-ns",
			wantErr:          false,
			wantExternal:     "projects/test-project/locations/us-central1/appGateways/ready-gateway",
		},
		{
			name: "by name and namespace when not ready",
			ref: &BeyondCorpAppGatewayRef{
				Name:      "not-ready-gateway",
				Namespace: "test-ns",
			},
			defaultNamespace: "test-ns",
			wantErr:          true,
		},
		{
			name: "by name when not found",
			ref: &BeyondCorpAppGatewayRef{
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

func TestBeyondCorpAppGatewayRef_Interface(t *testing.T) {
	ref := &BeyondCorpAppGatewayRef{
		Name:      "test-gateway",
		Namespace: "test-ns",
	}

	if got := ref.GetGVK(); got != BeyondCorpAppGatewayGVK {
		t.Errorf("GetGVK() = %v, want %v", got, BeyondCorpAppGatewayGVK)
	}

	wantNN := types.NamespacedName{
		Name:      "test-gateway",
		Namespace: "test-ns",
	}
	if gotNN := ref.GetNamespacedName(); gotNN != wantNN {
		t.Errorf("GetNamespacedName() = %v, want %v", gotNN, wantNN)
	}

	externalVal := "projects/p1/locations/us-central1/appGateways/gw1"
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
	gwID, ok := id.(*BeyondCorpAppGatewayIdentity)
	if !ok {
		t.Fatalf("ParseExternalToIdentity() returned %T, want *BeyondCorpAppGatewayIdentity", id)
	}
	wantID := &BeyondCorpAppGatewayIdentity{
		Project:    "p1",
		Location:   "us-central1",
		AppGateway: "gw1",
	}
	if diff := cmp.Diff(wantID, gwID); diff != "" {
		t.Errorf("ParseExternalToIdentity() mismatch (-want +got):\n%s", diff)
	}
}
