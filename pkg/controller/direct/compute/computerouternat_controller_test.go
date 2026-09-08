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

package compute

import (
	"context"
	"testing"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestRouterNATAdapter_RequestBuilding(t *testing.T) {
	s := runtime.NewScheme()
	s.AddKnownTypes(krm.GroupVersion, &krm.ComputeRouterNAT{}, &unstructured.Unstructured{})

	tests := []struct {
		name        string
		obj         *krm.ComputeRouterNAT
		objs        []runtime.Object
		wantRouter  string
		wantProject string
		wantRegion  string
	}{
		{
			name: "short name routerRef",
			obj: &krm.ComputeRouterNAT{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "nat-1",
					Namespace: "ns-1",
					Annotations: map[string]string{
						"cnrm.cloud.google.com/project-id": "proj-1",
					},
				},
				Spec: krm.ComputeRouterNATSpec{
					Region: "region-1",
					RouterRef: krm.ComputeRouterRef{
						External: "router-1",
					},
				},
			},
			wantRouter:  "router-1",
			wantProject: "proj-1",
			wantRegion:  "region-1",
		},
		{
			name: "canonical path routerRef",
			obj: &krm.ComputeRouterNAT{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "nat-1",
					Namespace: "ns-1",
				},
				Spec: krm.ComputeRouterNATSpec{
					RouterRef: krm.ComputeRouterRef{
						External: "projects/proj-1/regions/region-1/routers/router-1",
					},
				},
			},
			wantRouter:  "router-1",
			wantProject: "proj-1",
			wantRegion:  "region-1",
		},
		{
			name: "full URI routerRef",
			obj: &krm.ComputeRouterNAT{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "nat-1",
					Namespace: "ns-1",
				},
				Spec: krm.ComputeRouterNATSpec{
					RouterRef: krm.ComputeRouterRef{
						External: "https://www.googleapis.com/compute/v1/projects/proj-1/regions/region-1/routers/router-1",
					},
				},
			},
			wantRouter:  "router-1",
			wantProject: "proj-1",
			wantRegion:  "region-1",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.TODO()
			fakeClient := fake.NewClientBuilder().WithScheme(s).WithRuntimeObjects(tc.objs...).Build()

			ident, err := tc.obj.GetIdentity(ctx, fakeClient)
			if err != nil {
				t.Fatalf("unexpected error resolving identity: %v", err)
			}

			id := ident.(*krm.ComputeRouterNATIdentity)

			if id.Router != tc.wantRouter {
				t.Errorf("expected router %q, got %q", tc.wantRouter, id.Router)
			}
			if id.Project != tc.wantProject {
				t.Errorf("expected project %q, got %q", tc.wantProject, id.Project)
			}
			if id.Region != tc.wantRegion {
				t.Errorf("expected region %q, got %q", tc.wantRegion, id.Region)
			}
		})
	}
}
