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

package dataflow

import (
	"context"
	"testing"

	pb "cloud.google.com/go/dataflow/apiv1beta3/dataflowpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataflow/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func TestDataflowJob_Export(t *testing.T) {
	ctx := context.Background()

	actualJob := &pb.Job{
		Id:           "job-12345",
		Name:         "word-count-job",
		Type:         pb.JobType_JOB_TYPE_BATCH,
		CurrentState: pb.JobState_JOB_STATE_RUNNING,
		Location:     "us-central1",
		ProjectId:    "test-project",
		Labels: map[string]string{
			"env": "production",
		},
		Environment: &pb.Environment{
			TempStoragePrefix: "gs://test-bucket/tmp",
		},
	}

	adapter := &dataflowJobAdapter{
		id: &krm.DataflowJobIdentity{
			Project:  "test-project",
			Location: "us-central1",
			Job:      "job-12345",
		},
		actual: actualJob,
	}

	u, err := adapter.Export(ctx)
	if err != nil {
		t.Fatalf("Export failed: %v", err)
	}

	if u.GetName() != "job-12345" {
		t.Errorf("expected name to be %q, got %q", "job-12345", u.GetName())
	}

	gvk := u.GroupVersionKind()
	if gvk != krm.DataflowJobGVK {
		t.Errorf("expected GVK to be %s, got %s", krm.DataflowJobGVK, gvk)
	}

	labels := u.GetLabels()
	if labels["env"] != "production" {
		t.Errorf("expected label env to be %q, got %q", "production", labels["env"])
	}

	annotations := u.GetAnnotations()
	if annotations["cnrm.cloud.google.com/project-id"] != "test-project" {
		t.Errorf("expected project ID annotation to be %q, got %q", "test-project", annotations["cnrm.cloud.google.com/project-id"])
	}

	region, found, err := unstructured.NestedString(u.Object, "spec", "region")
	if err != nil || !found || region != "us-central1" {
		t.Errorf("spec.region not set correctly: got %q (found=%v, err=%v)", region, found, err)
	}

	resourceID, found, err := unstructured.NestedString(u.Object, "spec", "resourceID")
	if err != nil || !found || resourceID != "job-12345" {
		t.Errorf("spec.resourceID not set correctly: got %q (found=%v, err=%v)", resourceID, found, err)
	}

	tempGcs, found, err := unstructured.NestedString(u.Object, "spec", "tempGcsLocation")
	if err != nil || !found || tempGcs != "gs://test-bucket/tmp" {
		t.Errorf("spec.tempGcsLocation not set correctly: got %q (found=%v, err=%v)", tempGcs, found, err)
	}
}
