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

package contactcenterinsights

import (
	"context"
	"testing"

	contactcenterinsightspb "cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/contactcenterinsights/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func TestCCInsightsQAScorecardRevision_Export(t *testing.T) {
	ctx := context.Background()

	id := &krm.CCInsightsQAScorecardRevisionIdentity{
		Project:     "test-project",
		Location:    "us-central1",
		QaScorecard: "scorecard-123",
		Revision:    "rev-456",
	}

	actual := &contactcenterinsightspb.QaScorecardRevision{
		Name: "projects/test-project/locations/us-central1/qaScorecards/scorecard-123/revisions/rev-456",
		Snapshot: &contactcenterinsightspb.QaScorecard{
			DisplayName: "Scorecard Snapshot",
			Description: "Test scorecard snapshot description",
		},
	}

	adapter := &QaScorecardRevisionAdapter{
		id:     id,
		actual: actual,
	}

	u, err := adapter.Export(ctx)
	if err != nil {
		t.Fatalf("Export failed: %v", err)
	}

	if u == nil {
		t.Fatal("Export returned nil unstructured object")
	}

	// Verify GVK
	if got, want := u.GroupVersionKind(), krm.CCInsightsQAScorecardRevisionGVK; got != want {
		t.Errorf("GroupVersionKind() = %v, want %v", got, want)
	}

	// Verify Spec fields (ParentRef, ResourceID, Snapshot)
	spec, found, err := unstructured.NestedMap(u.Object, "spec")
	if err != nil {
		t.Fatalf("NestedMap spec failed: %v", err)
	}
	if !found {
		t.Fatal("spec field not found in unstructured object")
	}

	parentRef, found, _ := unstructured.NestedString(spec, "parentRef", "external")
	if !found || parentRef != "projects/test-project/locations/us-central1/qaScorecards/scorecard-123" {
		t.Errorf("parentRef.external = %q, want %q", parentRef, "projects/test-project/locations/us-central1/qaScorecards/scorecard-123")
	}

	resourceID, found, _ := unstructured.NestedString(spec, "resourceID")
	if !found || resourceID != "rev-456" {
		t.Errorf("resourceID = %q, want %q", resourceID, "rev-456")
	}

	displayName, found, _ := unstructured.NestedString(spec, "snapshot", "displayName")
	if !found || displayName != "Scorecard Snapshot" {
		t.Errorf("snapshot.displayName = %q, want %q", displayName, "Scorecard Snapshot")
	}
}

func TestIsInitialScorecardRevision(t *testing.T) {
	tests := []struct {
		name           string
		rev            *contactcenterinsightspb.QaScorecardRevision
		totalRevisions int
		want           bool
	}{
		{
			name: "initial provisioned revision with numeric ID and sole revision",
			rev: &contactcenterinsightspb.QaScorecardRevision{
				Name:  "projects/123/locations/us-central1/qaScorecards/sc-1/revisions/9642586111448304558",
				State: contactcenterinsightspb.QaScorecardRevision_EDITABLE,
			},
			totalRevisions: 1,
			want:           true,
		},
		{
			name: "editable revision but multiple revisions exist",
			rev: &contactcenterinsightspb.QaScorecardRevision{
				Name:  "projects/123/locations/us-central1/qaScorecards/sc-1/revisions/9642586111448304558",
				State: contactcenterinsightspb.QaScorecardRevision_EDITABLE,
			},
			totalRevisions: 2,
			want:           false,
		},
		{
			name: "user-defined non-numeric revision ID",
			rev: &contactcenterinsightspb.QaScorecardRevision{
				Name:  "projects/123/locations/us-central1/qaScorecards/sc-1/revisions/rev-user-1",
				State: contactcenterinsightspb.QaScorecardRevision_EDITABLE,
			},
			totalRevisions: 1,
			want:           false,
		},
		{
			name: "non-editable state",
			rev: &contactcenterinsightspb.QaScorecardRevision{
				Name:  "projects/123/locations/us-central1/qaScorecards/sc-1/revisions/9642586111448304558",
				State: contactcenterinsightspb.QaScorecardRevision_READY,
			},
			totalRevisions: 1,
			want:           false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := isInitialScorecardRevision(tc.rev, tc.totalRevisions)
			if got != tc.want {
				t.Errorf("isInitialScorecardRevision() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestCCInsightsQAScorecardRevision_UpdateSnapshotNormalization(t *testing.T) {
	ctx := context.Background()

	desired := &krm.CCInsightsQAScorecardRevisionSpec{
		Snapshot: &krm.QaScorecard{
			DisplayName: direct.LazyPtr("Scorecard Snapshot Minimal 123"),
		},
	}

	actual := &contactcenterinsightspb.QaScorecardRevision{
		Name: "projects/test-project/locations/us-central1/qaScorecards/scorecard-123/revisions/rev-456",
		Snapshot: &contactcenterinsightspb.QaScorecard{
			DisplayName: "Scorecard for revision 123",
		},
	}

	normalize := func(ctx context.Context, pb *contactcenterinsightspb.QaScorecardRevision) error {
		if actual != nil {
			pb.Snapshot = actual.GetSnapshot()
		}
		return nil
	}

	_, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		desired,
		actual,
		CCInsightsQAScorecardRevisionSpec_FromProto,
		CCInsightsQAScorecardRevisionSpec_ToProto,
		normalize,
	)
	if err != nil {
		t.Fatalf("CompareGreenfieldSpec failed: %v", err)
	}

	if report.HasDiff() {
		t.Errorf("expected no diff after normalization, got diff on paths: %v", updateMask.Paths)
	}
}
