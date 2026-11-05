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

package documentai

import (
	"context"
	"testing"

	pb "cloud.google.com/go/documentai/apiv1/documentaipb"
	krmv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/documentai/v1alpha1"
)

func TestProcessorDefaultVersionAdapter_diff(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name       string
		desiredVer string
		actualVer  string
		aliases    []*pb.ProcessorVersionAlias
		expectDiff bool
	}{
		{
			name:       "exact match",
			desiredVer: "projects/123/locations/us/processors/p1/processorVersions/v1",
			actualVer:  "projects/123/locations/us/processors/p1/processorVersions/v1",
			expectDiff: false,
		},
		{
			name:       "diff versions",
			desiredVer: "projects/123/locations/us/processors/p1/processorVersions/v2",
			actualVer:  "projects/123/locations/us/processors/p1/processorVersions/v1",
			expectDiff: true,
		},
		{
			name:       "alias matches resolved version",
			desiredVer: "projects/123/locations/us/processors/p1/processorVersions/rc",
			actualVer:  "projects/123/locations/us/processors/p1/processorVersions/v1",
			aliases: []*pb.ProcessorVersionAlias{
				{
					Alias:            "projects/123/locations/us/processors/p1/processorVersions/rc",
					ProcessorVersion: "projects/123/locations/us/processors/p1/processorVersions/v1",
				},
			},
			expectDiff: false,
		},
		{
			name:       "short alias name matches full path alias",
			desiredVer: "stable",
			actualVer:  "projects/123/locations/us/processors/p1/processorVersions/v1",
			aliases: []*pb.ProcessorVersionAlias{
				{
					Alias:            "projects/123/locations/us/processors/p1/processorVersions/stable",
					ProcessorVersion: "projects/123/locations/us/processors/p1/processorVersions/v1",
				},
			},
			expectDiff: false,
		},
		{
			name:       "short alias name does not match different alias",
			desiredVer: "rc",
			actualVer:  "projects/123/locations/us/processors/p1/processorVersions/v1",
			aliases: []*pb.ProcessorVersionAlias{
				{
					Alias:            "projects/123/locations/us/processors/p1/processorVersions/stable",
					ProcessorVersion: "projects/123/locations/us/processors/p1/processorVersions/v1",
				},
			},
			expectDiff: true,
		},
		{
			name:       "alias does not match actual version",
			desiredVer: "projects/123/locations/us/processors/p1/processorVersions/rc",
			actualVer:  "projects/123/locations/us/processors/p1/processorVersions/v2",
			aliases: []*pb.ProcessorVersionAlias{
				{
					Alias:            "projects/123/locations/us/processors/p1/processorVersions/rc",
					ProcessorVersion: "projects/123/locations/us/processors/p1/processorVersions/v1",
				},
			},
			expectDiff: true,
		},
		{
			name:       "project number vs project ID diff suppression",
			desiredVer: "projects/my-project/locations/us/processors/p1/processorVersions/v1",
			actualVer:  "projects/123456789/locations/us/processors/p1/processorVersions/v1",
			expectDiff: false,
		},
		{
			name:       "different project IDs without number should not be suppressed",
			desiredVer: "projects/project-a/locations/us/processors/p1/processorVersions/v1",
			actualVer:  "projects/project-b/locations/us/processors/p1/processorVersions/v1",
			expectDiff: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			adapter := &ProcessorDefaultVersionAdapter{
				desired: &krmv1alpha1.DocumentAIProcessorDefaultVersion{
					Spec: krmv1alpha1.DocumentAIProcessorDefaultVersionSpec{
						Version: tc.desiredVer,
					},
				},
				actual: &pb.Processor{
					DefaultProcessorVersion: tc.actualVer,
					ProcessorVersionAliases: tc.aliases,
				},
			}

			diff, err := adapter.diff(ctx)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if diff.HasDiff() != tc.expectDiff {
				t.Errorf("expected diff %v, got %v", tc.expectDiff, diff.HasDiff())
			}
		})
	}
}

func TestProcessorDefaultVersionAdapter_Export(t *testing.T) {
	ctx := context.Background()

	id := &krmv1alpha1.DocumentAIProcessorDefaultVersionIdentity{
		Project:   "test-project",
		Location:  "us",
		Processor: "test-proc",
	}

	adapter := &ProcessorDefaultVersionAdapter{
		id: id,
		actual: &pb.Processor{
			DefaultProcessorVersion: "projects/123/locations/us/processors/test-proc/processorVersions/v1",
		},
	}

	u, err := adapter.Export(ctx)
	if err != nil {
		t.Fatalf("unexpected export error: %v", err)
	}
	if u == nil {
		t.Fatal("expected non-nil export result")
	}
	if u.GetName() != "test-proc" {
		t.Errorf("expected name 'test-proc', got %q", u.GetName())
	}
	if u.GetKind() != "DocumentAIProcessorDefaultVersion" {
		t.Errorf("expected kind 'DocumentAIProcessorDefaultVersion', got %q", u.GetKind())
	}
}
