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
)

func TestClassifier_Classify(t *testing.T) {
	defaultFixture := &fixture.Data{
		FixtureName: "widget-test",
		PrimaryCreate: &fixture.KRMResource{
			Kind:  "ExampleWidget",
			Group: "example.cnrm.cloud.google.com",
			Name:  "widget-1",
		},
		Dependencies: []*fixture.KRMResource{
			{
				Kind:  "ExampleNetwork",
				Group: "example.cnrm.cloud.google.com",
				Name:  "network-1",
			},
		},
	}

	tests := []struct {
		name            string
		fix             *fixture.Data
		entries         []*logparser.Entry
		wantPhaseSource string
		wantRoles       []string
		wantKinds       []string
		wantPhases      []string
		wantParentIdxs  []int
	}{
		{
			name: "explicit Phase comments with DEPENDENCY and primary phases",
			fix:  defaultFixture,
			entries: []*logparser.Entry{
				{
					Phase: logparser.PhaseDependency,
					Request: logparser.HTTPRequest{
						Method: "POST",
						URL:    "https://example.googleapis.com/v1/projects/p1/locations/us-central1/networks",
						Body:   `{"name":"network-1"}`,
					},
					Response: logparser.HTTPResponse{Status: "200 OK", StatusCode: 200},
				},
				{
					Phase: logparser.PhaseCreate,
					Request: logparser.HTTPRequest{
						Method: "POST",
						URL:    "https://example.googleapis.com/v1/projects/p1/locations/us-central1/widgets?widgetId=widget-1",
						Body:   `{"name":"widget-1"}`,
					},
					Response: logparser.HTTPResponse{Status: "200 OK", StatusCode: 200},
				},
				{
					Phase: logparser.PhaseUpdate,
					Request: logparser.HTTPRequest{
						Method: "PATCH",
						URL:    "https://example.googleapis.com/v1/projects/p1/locations/us-central1/widgets/widget-1?updateMask=description",
						Body:   `{"description":"v2"}`,
					},
					Response: logparser.HTTPResponse{Status: "200 OK", StatusCode: 200},
				},
				{
					Phase: logparser.PhaseDelete,
					Request: logparser.HTTPRequest{
						Method: "DELETE",
						URL:    "https://example.googleapis.com/v1/projects/p1/locations/us-central1/widgets/widget-1",
					},
					Response: logparser.HTTPResponse{Status: "200 OK", StatusCode: 200},
				},
			},
			wantPhaseSource: PhaseSourceExplicit,
			wantRoles:       []string{RoleDependency, RolePrimary, RolePrimary, RolePrimary},
			wantKinds:       []string{"ExampleNetwork", "ExampleWidget", "ExampleWidget", "ExampleWidget"},
			wantPhases:      []string{logparser.PhaseCreate, logparser.PhaseCreate, logparser.PhaseUpdate, logparser.PhaseDelete},
			wantParentIdxs:  []int{0, 0, 0, 0},
		},
		{
			name: "interleaved dependency and primary LRO polls op-net vs op-widget",
			fix:  defaultFixture,
			entries: []*logparser.Entry{
				{
					Phase: logparser.PhaseCreate,
					Request: logparser.HTTPRequest{
						Method: "POST",
						URL:    "https://example.googleapis.com/v1/projects/p1/locations/us-central1/networks?networkId=network-1",
						Body:   `{"name":"network-1"}`,
					},
					Response: logparser.HTTPResponse{
						Status:     "200 OK",
						StatusCode: 200,
						Body:       `{"name":"projects/p1/locations/us-central1/operations/op-net","done":false}`,
					},
				},
				{
					Phase: logparser.PhaseCreate,
					Request: logparser.HTTPRequest{
						Method: "POST",
						URL:    "https://example.googleapis.com/v1/projects/p1/locations/us-central1/widgets?widgetId=widget-1",
						Body:   `{"name":"widget-1"}`,
					},
					Response: logparser.HTTPResponse{
						Status:     "200 OK",
						StatusCode: 200,
						Body:       `{"name":"projects/p1/locations/us-central1/operations/op-widget","done":false}`,
					},
				},
				{
					Phase: logparser.PhaseCreate,
					Request: logparser.HTTPRequest{
						Method: "GET",
						URL:    "https://example.googleapis.com/v1/projects/p1/locations/us-central1/operations/op-net",
					},
					Response: logparser.HTTPResponse{
						Status:     "200 OK",
						StatusCode: 200,
						Body:       `{"name":"projects/p1/locations/us-central1/operations/op-net","done":true}`,
					},
				},
				{
					Phase: logparser.PhaseCreate,
					Request: logparser.HTTPRequest{
						Method: "GET",
						URL:    "https://example.googleapis.com/v1/projects/p1/locations/us-central1/operations/op-widget",
					},
					Response: logparser.HTTPResponse{
						Status:     "200 OK",
						StatusCode: 200,
						Body:       `{"name":"projects/p1/locations/us-central1/operations/op-widget","done":true,"response":{"name":"widget-1"}}`,
					},
				},
			},
			wantPhaseSource: PhaseSourceExplicit,
			wantRoles:       []string{RoleDependency, RolePrimary, RoleDependency, RolePrimary},
			wantKinds:       []string{"ExampleNetwork", "ExampleWidget", "ExampleNetwork", "ExampleWidget"},
			wantPhases:      []string{logparser.PhaseCreate, logparser.PhaseCreate, logparser.PhaseCreate, logparser.PhaseCreate},
			wantParentIdxs:  []int{0, 0, 1, 2},
		},
		{
			name: "unmarked entries inferred as PhaseCreate when Tier 2 is enabled",
			fix:  defaultFixture,
			entries: []*logparser.Entry{
				{
					Request: logparser.HTTPRequest{
						Method: "POST",
						URL:    "https://example.googleapis.com/v1/projects/p1/locations/us-central1/widgets?widgetId=widget-1",
					},
					Response: logparser.HTTPResponse{Status: "200 OK", StatusCode: 200},
				},
			},
			wantPhaseSource: PhaseSourceInferred,
			wantRoles:       []string{RolePrimary},
			wantKinds:       []string{"ExampleWidget"},
			wantPhases:      []string{logparser.PhaseCreate},
			wantParentIdxs:  []int{0},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			clf := New(tc.fix)

			// Act
			classified, phaseSource := clf.Classify(tc.entries)

			// Assert
			if phaseSource != tc.wantPhaseSource {
				t.Fatalf("phaseSource = %q, want %q", phaseSource, tc.wantPhaseSource)
			}
			if len(classified) != len(tc.wantRoles) {
				t.Fatalf("len(classified) = %d, want %d", len(classified), len(tc.wantRoles))
			}
			var gotRoles, gotKinds, gotPhases []string
			var gotParents []int
			for _, ce := range classified {
				gotRoles = append(gotRoles, ce.Role)
				gotKinds = append(gotKinds, ce.ResourceKind)
				gotPhases = append(gotPhases, ce.Phase)
				gotParents = append(gotParents, ce.ParentIndex)
			}
			if !reflect.DeepEqual(gotRoles, tc.wantRoles) {
				t.Errorf("Roles = %v, want %v", gotRoles, tc.wantRoles)
			}
			if !reflect.DeepEqual(gotKinds, tc.wantKinds) {
				t.Errorf("Kinds = %v, want %v", gotKinds, tc.wantKinds)
			}
			if !reflect.DeepEqual(gotPhases, tc.wantPhases) {
				t.Errorf("Phases = %v, want %v", gotPhases, tc.wantPhases)
			}
			if !reflect.DeepEqual(gotParents, tc.wantParentIdxs) {
				t.Errorf("ParentIndices = %v, want %v", gotParents, tc.wantParentIdxs)
			}
		})
	}
}
