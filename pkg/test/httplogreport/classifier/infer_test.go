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

package classifier

import (
	"reflect"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/httplogreport/fixture"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/httplogreport/logparser"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func TestClassifyByInference(t *testing.T) {
	createObj := &unstructured.Unstructured{
		Object: map[string]any{
			"apiVersion": "example.cnrm.cloud.google.com/v1alpha1",
			"kind":       "ExampleCluster",
			"metadata":   map[string]any{"name": "cluster-1"},
			"spec": map[string]any{
				"description":     "initial",
				"secondaryConfig": "enabled-on-create",
			},
		},
	}
	updateObj := &unstructured.Unstructured{
		Object: map[string]any{
			"apiVersion": "example.cnrm.cloud.google.com/v1alpha1",
			"kind":       "ExampleCluster",
			"metadata":   map[string]any{"name": "cluster-1"},
			"spec": map[string]any{
				"description":     "updated-in-update-phase",
				"secondaryConfig": "enabled-on-create",
			},
		},
	}

	tests := []struct {
		name       string
		fix        *fixture.Data
		rawLog     string
		wantPhases []string
		wantRoles  []string
		wantKinds  []string
	}{
		{
			name: "composite create and update phase inference",
			fix: &fixture.Data{
				PrimaryCreate: &fixture.KRMResource{
					Kind:  "ExampleCluster",
					Group: "example.cnrm.cloud.google.com",
					Name:  "cluster-1",
					Obj:   createObj,
				},
				PrimaryUpdate: &fixture.KRMResource{
					Kind:  "ExampleCluster",
					Group: "example.cnrm.cloud.google.com",
					Name:  "cluster-1",
					Obj:   updateObj,
				},
			},
			rawLog: `POST https://example.googleapis.com/v1/projects/p/locations/l/clusters?clusterId=cluster-1
Content-Type: application/json

{"description": "initial"}

200 OK

{"description": "initial"}

---

PATCH https://example.googleapis.com/v1/projects/p/locations/l/clusters/cluster-1?updateMask=secondaryConfig
Content-Type: application/json

{"secondaryConfig": "enabled-on-create"}

200 OK

{"description": "initial", "secondaryConfig": "enabled-on-create"}

---

GET https://example.googleapis.com/v1/projects/p/locations/l/clusters/cluster-1

200 OK

{"description": "initial", "secondaryConfig": "enabled-on-create"}

---

PATCH https://example.googleapis.com/v1/projects/p/locations/l/clusters/cluster-1?updateMask=description
Content-Type: application/json

{"description": "updated-in-update-phase"}

200 OK

{"description": "updated-in-update-phase", "secondaryConfig": "enabled-on-create"}

---

DELETE https://example.googleapis.com/v1/projects/p/locations/l/clusters/cluster-1

200 OK

{}
`,
			wantPhases: []string{
				logparser.PhaseCreate,
				logparser.PhaseCreate,
				logparser.PhaseCreate,
				logparser.PhaseUpdate,
				logparser.PhaseDelete,
			},
			wantRoles: []string{
				RolePrimary,
				RolePrimary,
				RolePrimary,
				RolePrimary,
				RolePrimary,
			},
			wantKinds: []string{
				"ExampleCluster",
				"ExampleCluster",
				"ExampleCluster",
				"ExampleCluster",
				"ExampleCluster",
			},
		},
		{
			name: "legacy unmarked dependency collection and placeholder inference",
			fix: &fixture.Data{
				PrimaryCreate: &fixture.KRMResource{
					Kind:  "ExampleCluster",
					Group: "example.cnrm.cloud.google.com",
					Name:  "cluster-1",
					Obj:   createObj,
				},
				Dependencies: []*fixture.KRMResource{
					{
						Kind:  "ComputeNetwork",
						Group: "compute.cnrm.cloud.google.com",
						Name:  "net-dep",
					},
				},
			},
			rawLog: `POST https://compute.googleapis.com/compute/v1/projects/p/global/networks
Content-Type: application/json

{"name": "${networkID}"}

200 OK

{}

---

POST https://example.googleapis.com/v1/projects/p/locations/l/clusters?clusterId=cluster-1
Content-Type: application/json

{"description": "initial"}

200 OK

{}

---

DELETE https://example.googleapis.com/v1/projects/p/locations/l/clusters/cluster-1

200 OK

{}
`,
			wantPhases: []string{
				logparser.PhaseCreate,
				logparser.PhaseCreate,
				logparser.PhaseDelete,
			},
			wantRoles: []string{
				RoleDependency,
				RolePrimary,
				RolePrimary,
			},
			wantKinds: []string{
				"ComputeNetwork",
				"ExampleCluster",
				"ExampleCluster",
			},
		},
		{
			name: "hierarchical parent dependency does not swallow child primary resource",
			fix: &fixture.Data{
				PrimaryCreate: &fixture.KRMResource{
					Kind:  "DataplexGlossaryTerm",
					Group: "dataplex.cnrm.cloud.google.com",
					Name:  "dpgt-1",
				},
				Dependencies: []*fixture.KRMResource{
					{
						Kind:  "DataplexGlossary",
						Group: "dataplex.cnrm.cloud.google.com",
						Name:  "dpg-1",
					},
				},
			},
			rawLog: `POST https://dataplex.googleapis.com/v1/projects/p/locations/l/glossaries?glossaryId=dpg-1
Content-Type: application/json

{"displayName": "glossary"}

200 OK

{}

---

POST https://dataplex.googleapis.com/v1/projects/p/locations/l/glossaries/dpg-1/terms?termId=dpgt-1
Content-Type: application/json

{"displayName": "term"}

200 OK

{}

---

DELETE https://dataplex.googleapis.com/v1/projects/p/locations/l/glossaries/dpg-1/terms/dpgt-1

200 OK

{}
`,
			wantPhases: []string{
				logparser.PhaseCreate,
				logparser.PhaseCreate,
				logparser.PhaseDelete,
			},
			wantRoles: []string{
				RoleDependency,
				RolePrimary,
				RolePrimary,
			},
			wantKinds: []string{
				"DataplexGlossary",
				"DataplexGlossaryTerm",
				"DataplexGlossaryTerm",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			entries, err := logparser.Parse(tc.rawLog)
			if err != nil {
				t.Fatalf("Parse failed: %v", err)
			}
			clf := New(tc.fix)

			// Act
			classified, phaseSource := clf.Classify(entries)

			// Assert
			if phaseSource != PhaseSourceInferred {
				t.Fatalf("expected PhaseSourceInferred, got %s", phaseSource)
			}
			var gotPhases, gotRoles, gotKinds []string
			for _, ce := range classified {
				gotPhases = append(gotPhases, ce.Phase)
				gotRoles = append(gotRoles, ce.Role)
				gotKinds = append(gotKinds, ce.ResourceKind)
			}
			if !reflect.DeepEqual(gotPhases, tc.wantPhases) {
				t.Errorf("inferred phases = %v, want %v", gotPhases, tc.wantPhases)
			}
			if !reflect.DeepEqual(gotRoles, tc.wantRoles) {
				t.Errorf("inferred roles = %v, want %v", gotRoles, tc.wantRoles)
			}
			if !reflect.DeepEqual(gotKinds, tc.wantKinds) {
				t.Errorf("inferred kinds = %v, want %v", gotKinds, tc.wantKinds)
			}
		})
	}
}
