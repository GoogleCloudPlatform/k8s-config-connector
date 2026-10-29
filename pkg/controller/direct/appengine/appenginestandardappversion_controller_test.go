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

package appengine

import (
	"context"
	"reflect"
	"sort"
	"testing"

	pb "cloud.google.com/go/appengine/apiv1/appenginepb"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/appengine/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
)

func TestStandardAppVersion_CompareSpecifiedSpec(t *testing.T) {
	ctx := context.Background()

	actual := &pb.Version{
		Scaling: &pb.Version_AutomaticScaling{
			AutomaticScaling: &pb.AutomaticScaling{
				MinIdleInstances:      1,
				MaxIdleInstances:      3,
				MaxConcurrentRequests: 10, // unmanaged in desired KRM
				StandardSchedulerSettings: &pb.StandardSchedulerSettings{
					TargetCpuUtilization: 0.5,
					MinInstances:         2,
					MaxInstances:         10,
				},
			},
		},
		InstanceClass: "F1",
		Entrypoint: &pb.Entrypoint{
			Command: &pb.Entrypoint_Shell{
				Shell: "gunicorn -b :$PORT main:app",
			},
		},
		Deployment: &pb.Deployment{
			Files: map[string]*pb.FileInfo{
				"main.py": {
					SourceUrl: "https://storage.googleapis.com/staging.mock-project.appspot.com/main.py",
					Sha1Sum:   "abc123sha1",
				},
			},
		},
	}

	desired := &krm.AppEngineStandardAppVersionSpec{
		AutomaticScaling: &krm.AutomaticScaling{
			MinIdleInstances: direct.LazyPtr(2), // changed from 1 to 2
			MaxIdleInstances: direct.LazyPtr(4), // changed from 3 to 4
			// MaxConcurrentRequests is omitted/unmanaged
			StandardSchedulerSettings: &krm.StandardSchedulerSettings{
				TargetCPUUtilization: direct.LazyPtr(0.5), // unchanged
				MinInstances:         direct.LazyPtr(2),   // unchanged
				MaxInstances:         direct.LazyPtr(10),  // unchanged
			},
		},
		InstanceClass: direct.LazyPtr("F2"), // changed from F1 to F2
		Entrypoint: &krm.Entrypoint{
			Shell: "gunicorn -b :$PORT main:app", // server doesn't return this; normalizer ignores diff
		},
		Deployment: &krm.Deployment{
			Files: []krm.DeploymentFile{
				{
					Name:      "main.py",
					SourceUrl: "https://storage.googleapis.com/staging.mock-project.appspot.com/main.py",
					// Sha1Sum omitted; normalizer ignores diff
				},
			},
		},
	}

	diffRes, err := common.CompareSpecifiedSpec(
		ctx,
		desired,
		actual,
		AppEngineStandardAppVersionSpec_v1alpha1_FromProto,
		AppEngineStandardAppVersionSpec_v1alpha1_ToProto,
		normalizeStandardAppVersionForDiff,
		true, // fineGrained
	)
	if err != nil {
		t.Fatalf("unexpected error from CompareSpecifiedSpec: %v", err)
	}

	if diffRes.Empty() {
		t.Fatalf("expected non-empty diff")
	}

	gotPaths := diffRes.FieldMask.GetPaths()
	sort.Strings(gotPaths)
	wantPaths := []string{
		"automatic_scaling.max_idle_instances",
		"automatic_scaling.min_idle_instances",
		"instance_class",
	}
	sort.Strings(wantPaths)

	if !reflect.DeepEqual(gotPaths, wantPaths) {
		t.Fatalf("unexpected paths in fieldmask: got %v, want %v", gotPaths, wantPaths)
	}

	// Verify unmanaged field is preserved in MergedDesired
	if diffRes.MergedDesired.GetAutomaticScaling().GetMaxConcurrentRequests() != 10 {
		t.Fatalf("expected MergedDesired to preserve unmanaged max_concurrent_requests=10, got %d",
			diffRes.MergedDesired.GetAutomaticScaling().GetMaxConcurrentRequests())
	}
}
