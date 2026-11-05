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
	"sigs.k8s.io/controller-runtime/pkg/client/fake"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"
	"github.com/google/go-cmp/cmp"
)

func TestBeyondCorpAppConnector_GetIdentity(t *testing.T) {
	ctx := context.Background()

	resourceID := "custom-connector"
	tests := []struct {
		name         string
		obj          *BeyondCorpAppConnector
		wantIdentity *BeyondCorpAppConnectorIdentity
		wantErr      bool
	}{
		{
			name: "valid spec with explicit resourceID and projectRef.external",
			obj: &BeyondCorpAppConnector{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-connector",
					Namespace: "test-ns",
				},
				Spec: BeyondCorpAppConnectorSpec{
					ResourceID: &resourceID,
					Region:     "us-central1",
					ProjectRef: refs.ProjectRef{
						External: "projects/my-project",
					},
				},
			},
			wantIdentity: &BeyondCorpAppConnectorIdentity{
				Project:      "my-project",
				Location:     "us-central1",
				AppConnector: "custom-connector",
			},
			wantErr: false,
		},
		{
			name: "valid spec defaulting to metadata.name and project-id annotation",
			obj: &BeyondCorpAppConnector{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-connector",
					Namespace: "test-ns",
					Annotations: map[string]string{
						"cnrm.cloud.google.com/project-id": "annotated-project",
					},
				},
				Spec: BeyondCorpAppConnectorSpec{
					Region: "europe-west1",
				},
			},
			wantIdentity: &BeyondCorpAppConnectorIdentity{
				Project:      "annotated-project",
				Location:     "europe-west1",
				AppConnector: "test-connector",
			},
			wantErr: false,
		},
		{
			name: "valid spec with matching status.externalRef",
			obj: &BeyondCorpAppConnector{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-connector",
					Namespace: "test-ns",
				},
				Spec: BeyondCorpAppConnectorSpec{
					ResourceID: &resourceID,
					Region:     "us-central1",
					ProjectRef: refs.ProjectRef{
						External: "projects/my-project",
					},
				},
				Status: BeyondCorpAppConnectorStatus{
					ExternalRef: common.LazyPtr("projects/my-project/locations/us-central1/appConnectors/custom-connector"),
				},
			},
			wantIdentity: &BeyondCorpAppConnectorIdentity{
				Project:      "my-project",
				Location:     "us-central1",
				AppConnector: "custom-connector",
			},
			wantErr: false,
		},
		{
			name: "status.externalRef does not match spec identity",
			obj: &BeyondCorpAppConnector{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-connector",
					Namespace: "test-ns",
				},
				Spec: BeyondCorpAppConnectorSpec{
					ResourceID: &resourceID,
					Region:     "us-central1",
					ProjectRef: refs.ProjectRef{
						External: "projects/my-project",
					},
				},
				Status: BeyondCorpAppConnectorStatus{
					ExternalRef: common.LazyPtr("projects/other-project/locations/us-central1/appConnectors/custom-connector"),
				},
			},
			wantErr: true,
		},
		{
			name: "malformed status.externalRef",
			obj: &BeyondCorpAppConnector{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-connector",
					Namespace: "test-ns",
				},
				Spec: BeyondCorpAppConnectorSpec{
					ResourceID: &resourceID,
					Region:     "us-central1",
					ProjectRef: refs.ProjectRef{
						External: "projects/my-project",
					},
				},
				Status: BeyondCorpAppConnectorStatus{
					ExternalRef: common.LazyPtr("invalid-external-ref"),
				},
			},
			wantErr: true,
		},
		{
			name: "missing region",
			obj: &BeyondCorpAppConnector{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-connector",
					Namespace: "test-ns",
					Annotations: map[string]string{
						"cnrm.cloud.google.com/project-id": "my-project",
					},
				},
				Spec: BeyondCorpAppConnectorSpec{},
			},
			wantErr: true,
		},
		{
			name: "missing project",
			obj: &BeyondCorpAppConnector{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-connector",
					Namespace: "test-ns",
				},
				Spec: BeyondCorpAppConnectorSpec{
					Region: "us-central1",
				},
			},
			wantErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			gotIdentity, err := tc.obj.GetIdentity(ctx, nil)
			if (err != nil) != tc.wantErr {
				t.Fatalf("GetIdentity() error = %v, wantErr %v", err, tc.wantErr)
			}
			if tc.wantErr {
				return
			}
			if diff := cmp.Diff(tc.wantIdentity, gotIdentity); diff != "" {
				t.Errorf("GetIdentity() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestBeyondCorpAppConnectorRef_Normalize(t *testing.T) {
	ctx := context.Background()

	scheme := runtime.NewScheme()
	_ = AddToScheme(scheme)
	_ = corev1.AddToScheme(scheme)

	externalRef := "projects/my-project/locations/us-central1/appConnectors/custom-connector"
	readyObj := &BeyondCorpAppConnector{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "ready-connector",
			Namespace: "test-ns",
		},
		Status: BeyondCorpAppConnectorStatus{
			ExternalRef: &externalRef,
		},
	}

	notReadyObj := &BeyondCorpAppConnector{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "not-ready-connector",
			Namespace: "test-ns",
		},
		Status: BeyondCorpAppConnectorStatus{},
	}

	fakeClient := fake.NewClientBuilder().WithScheme(scheme).WithObjects(readyObj, notReadyObj).Build()

	tests := []struct {
		name             string
		ref              *BeyondCorpAppConnectorRef
		defaultNamespace string
		wantErr          bool
		wantExternal     string
	}{
		{
			name: "external reference valid",
			ref: &BeyondCorpAppConnectorRef{
				External: "projects/p1/locations/l1/appConnectors/c1",
			},
			defaultNamespace: "test-ns",
			wantErr:          false,
			wantExternal:     "projects/p1/locations/l1/appConnectors/c1",
		},
		{
			name: "external reference invalid",
			ref: &BeyondCorpAppConnectorRef{
				External: "invalid/format",
			},
			defaultNamespace: "test-ns",
			wantErr:          true,
		},
		{
			name: "by name and namespace with status.externalRef set",
			ref: &BeyondCorpAppConnectorRef{
				Name:      "ready-connector",
				Namespace: "test-ns",
			},
			defaultNamespace: "test-ns",
			wantErr:          false,
			wantExternal:     "projects/my-project/locations/us-central1/appConnectors/custom-connector",
		},
		{
			name: "by name with default namespace with status.externalRef set",
			ref: &BeyondCorpAppConnectorRef{
				Name: "ready-connector",
			},
			defaultNamespace: "test-ns",
			wantErr:          false,
			wantExternal:     "projects/my-project/locations/us-central1/appConnectors/custom-connector",
		},
		{
			name: "by name and namespace when status.externalRef not set (not ready)",
			ref: &BeyondCorpAppConnectorRef{
				Name:      "not-ready-connector",
				Namespace: "test-ns",
			},
			defaultNamespace: "test-ns",
			wantErr:          true,
		},
		{
			name: "by name when not found",
			ref: &BeyondCorpAppConnectorRef{
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
