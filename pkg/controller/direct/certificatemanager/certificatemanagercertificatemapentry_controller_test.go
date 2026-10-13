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

package certificatemanager

import (
	"context"
	"reflect"
	"testing"

	pb "cloud.google.com/go/certificatemanager/apiv1/certificatemanagerpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/certificatemanager/v1beta1"
)

func TestNormalizeActualCertificates(t *testing.T) {
	tests := []struct {
		name         string
		entryProject string
		actual       []string
		desired      []string
		expected     []string
	}{
		{
			name:         "normalizes project number to entry project ID",
			entryProject: "test-project",
			actual: []string{
				"projects/123456789012/locations/global/certificates/my-cert",
			},
			desired: []string{
				"projects/test-project/locations/global/certificates/my-cert",
			},
			expected: []string{
				"projects/test-project/locations/global/certificates/my-cert",
			},
		},
		{
			name:         "leaves existing project ID unchanged",
			entryProject: "test-project",
			actual: []string{
				"projects/test-project/locations/global/certificates/my-cert",
			},
			desired: []string{
				"projects/test-project/locations/global/certificates/my-cert",
			},
			expected: []string{
				"projects/test-project/locations/global/certificates/my-cert",
			},
		},
		{
			name:         "does not normalize when desired project does not match entry project",
			entryProject: "test-project",
			actual: []string{
				"projects/123456789012/locations/global/certificates/my-cert",
			},
			desired: []string{
				"projects/other-project/locations/global/certificates/my-cert",
			},
			expected: []string{
				"projects/123456789012/locations/global/certificates/my-cert",
			},
		},
		{
			name:         "multiple certificates mapped correctly",
			entryProject: "test-project",
			actual: []string{
				"projects/123456789012/locations/global/certificates/cert-1",
				"projects/123456789012/locations/global/certificates/cert-2",
			},
			desired: []string{
				"projects/test-project/locations/global/certificates/cert-1",
				"projects/test-project/locations/global/certificates/cert-2",
			},
			expected: []string{
				"projects/test-project/locations/global/certificates/cert-1",
				"projects/test-project/locations/global/certificates/cert-2",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := normalizeActualCertificates(tc.actual, tc.desired, tc.entryProject)
			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("normalizeActualCertificates() = %v, want %v", got, tc.expected)
			}
		})
	}
}

func TestCompareCertificateMapEntry_CrossProjectDetection(t *testing.T) {
	ctx := context.Background()

	adapter := &CertificateMapEntryAdapter{
		id: &krm.CertificateManagerCertificateMapEntryIdentity{
			Project:             "project-a",
			CertificateMap:      "map-1",
			CertificateMapEntry: "entry-1",
		},
	}

	actual := &pb.CertificateMapEntry{
		Name: "projects/111111111111/locations/global/certificateMaps/map-1/certificateMapEntries/entry-1",
		Match: &pb.CertificateMapEntry_Hostname{
			Hostname: "example.com",
		},
		Certificates: []string{
			"projects/111111111111/locations/global/certificates/my-cert",
		},
	}

	// 1. Desired points to same cert using Project ID "project-a": should be 0 diff
	desiredSameProject := &pb.CertificateMapEntry{
		Name: "projects/project-a/locations/global/certificateMaps/map-1/certificateMapEntries/entry-1",
		Match: &pb.CertificateMapEntry_Hostname{
			Hostname: "example.com",
		},
		Certificates: []string{
			"projects/project-a/locations/global/certificates/my-cert",
		},
	}

	diffs, updateMask, err := adapter.compareCertificateMapEntry(ctx, actual, desiredSameProject)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if diffs.HasDiff() {
		t.Errorf("expected no diff when project number resolves to same project ID, got diffs: %v", diffs.Fields)
	}
	if len(updateMask.GetPaths()) != 0 {
		t.Errorf("expected empty update mask, got: %v", updateMask.GetPaths())
	}

	// 2. Desired points to cert in Project B ("project-b"): must detect diff (no false cross-project equivalence!)
	desiredDiffProject := &pb.CertificateMapEntry{
		Name: "projects/project-a/locations/global/certificateMaps/map-1/certificateMapEntries/entry-1",
		Match: &pb.CertificateMapEntry_Hostname{
			Hostname: "example.com",
		},
		Certificates: []string{
			"projects/project-b/locations/global/certificates/my-cert",
		},
	}

	diffs, updateMask, err = adapter.compareCertificateMapEntry(ctx, actual, desiredDiffProject)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !diffs.HasDiff() {
		t.Errorf("expected diff when certificates are in different projects, but diff was empty")
	}
	foundCertPath := false
	for _, p := range updateMask.GetPaths() {
		if p == "certificates" {
			foundCertPath = true
			break
		}
	}
	if !foundCertPath {
		t.Errorf("expected update mask to include 'certificates', got: %v", updateMask.GetPaths())
	}
}
