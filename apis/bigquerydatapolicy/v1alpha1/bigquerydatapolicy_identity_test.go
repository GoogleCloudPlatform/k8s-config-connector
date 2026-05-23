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

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestParseDataPolicyExternal(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		expectedP  *DataPolicyParent
		expectedID string
		hasError   bool
	}{
		{
			name:  "valid external ref",
			input: "projects/my-project/locations/us-central1/datapolicies/my-datapolicy",
			expectedP: &DataPolicyParent{
				ProjectID: "my-project",
				Location:  "us-central1",
			},
			expectedID: "my-datapolicy",
			hasError:   false,
		},
		{
			name:     "invalid external ref - wrong segments count",
			input:    "projects/my-project/locations/us-central1/datapolicies",
			hasError: true,
		},
		{
			name:     "invalid external ref - wrong static keyword projects",
			input:    "project/my-project/locations/us-central1/datapolicies/my-datapolicy",
			hasError: true,
		},
		{
			name:     "invalid external ref - wrong static keyword datapolicies",
			input:    "projects/my-project/locations/us-central1/policies/my-datapolicy",
			hasError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			parent, id, err := ParseDataPolicyExternal(tc.input)
			if tc.hasError {
				if err == nil {
					t.Fatalf("expected error for %q, but got none", tc.input)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if parent.ProjectID != tc.expectedP.ProjectID || parent.Location != tc.expectedP.Location {
				t.Errorf("expected parent %+v, got %+v", tc.expectedP, parent)
			}
			if id != tc.expectedID {
				t.Errorf("expected ID %q, got %q", tc.expectedID, id)
			}
		})
	}
}

func TestNewDataPolicyIdentity(t *testing.T) {
	scheme := runtime.NewScheme()
	_ = SchemeBuilder.AddToScheme(scheme)

	// Create fake client with some dummy project objects
	fakeClient := fake.NewClientBuilder().WithScheme(scheme).Build()

	tests := []struct {
		name        string
		obj         *BigQueryDataPolicy
		expectedErr bool
		expectedStr string
	}{
		{
			name: "valid inline projectRef and location",
			obj: &BigQueryDataPolicy{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "my-datapolicy",
					Namespace: "default",
				},
				Spec: BigQueryDataPolicySpec{
					Parent: &Parent{
						Location: "us-central1",
						ProjectRef: &refsv1beta1.ProjectRef{
							External: "my-project",
						},
					},
				},
			},
			expectedErr: false,
			expectedStr: "projects/my-project/locations/us-central1/datapolicies/my-datapolicy",
		},
		{
			name: "valid with externalRef set matching spec",
			obj: &BigQueryDataPolicy{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "my-datapolicy",
					Namespace: "default",
				},
				Spec: BigQueryDataPolicySpec{
					Parent: &Parent{
						Location: "us-central1",
						ProjectRef: &refsv1beta1.ProjectRef{
							External: "my-project",
						},
					},
				},
				Status: BigQueryDataPolicyStatus{
					ExternalRef: ptrTo("projects/my-project/locations/us-central1/datapolicies/my-datapolicy"),
				},
			},
			expectedErr: false,
			expectedStr: "projects/my-project/locations/us-central1/datapolicies/my-datapolicy",
		},
		{
			name: "invalid with mismatched project in externalRef",
			obj: &BigQueryDataPolicy{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "my-datapolicy",
					Namespace: "default",
				},
				Spec: BigQueryDataPolicySpec{
					Parent: &Parent{
						Location: "us-central1",
						ProjectRef: &refsv1beta1.ProjectRef{
							External: "my-project",
						},
					},
				},
				Status: BigQueryDataPolicyStatus{
					ExternalRef: ptrTo("projects/another-project/locations/us-central1/datapolicies/my-datapolicy"),
				},
			},
			expectedErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			id, err := NewDataPolicyIdentity(ctx, fakeClient, tc.obj)
			if tc.expectedErr {
				if err == nil {
					t.Fatal("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if id.String() != tc.expectedStr {
				t.Errorf("expected identity string %q, got %q", tc.expectedStr, id.String())
			}
			if id.ID() != tc.obj.GetName() {
				t.Errorf("expected ID %q, got %q", tc.obj.GetName(), id.ID())
			}
		})
	}
}

func ptrTo[T any](t T) *T {
	return &t
}

func (obj *BigQueryDataPolicy) GetNamespace() string {
	return obj.Namespace
}
