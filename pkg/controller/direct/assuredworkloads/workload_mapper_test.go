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

package assuredworkloads

import (
	"reflect"
	"testing"

	pb "cloud.google.com/go/assuredworkloads/apiv1/assuredworkloadspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/assuredworkloads/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func TestWorkload_SaaEnrollmentResponse_ToProto(t *testing.T) {
	mapCtx := &direct.MapContext{}
	in := &krm.Workload_SaaEnrollmentResponse{
		SetupErrors: []string{"SETUP_ERROR_UNSPECIFIED", "ERROR_INVALID_BASE_SETUP"},
		SetupStatus: direct.PtrTo("SETUP_STATE_UNSPECIFIED"),
	}

	out := Workload_SaaEnrollmentResponse_ToProto(mapCtx, in)

	if mapCtx.Err() != nil {
		t.Fatalf("unexpected error: %v", mapCtx.Err())
	}

	if out == nil {
		t.Fatal("expected non-nil output")
	}

	expectedErrors := []pb.Workload_SaaEnrollmentResponse_SetupError{
		pb.Workload_SaaEnrollmentResponse_SETUP_ERROR_UNSPECIFIED,
		pb.Workload_SaaEnrollmentResponse_ERROR_INVALID_BASE_SETUP,
	}

	if !reflect.DeepEqual(out.SetupErrors, expectedErrors) {
		t.Errorf("expected SetupErrors %v, got %v", expectedErrors, out.SetupErrors)
	}

	if out.SetupStatus == nil || *out.SetupStatus != pb.Workload_SaaEnrollmentResponse_SETUP_STATE_UNSPECIFIED {
		t.Errorf("expected SetupStatus %v, got %v", pb.Workload_SaaEnrollmentResponse_SETUP_STATE_UNSPECIFIED, out.SetupStatus)
	}
}

func TestWorkload_SaaEnrollmentResponse_FromProto(t *testing.T) {
	mapCtx := &direct.MapContext{}
	in := &pb.Workload_SaaEnrollmentResponse{
		SetupErrors: []pb.Workload_SaaEnrollmentResponse_SetupError{
			pb.Workload_SaaEnrollmentResponse_SETUP_ERROR_UNSPECIFIED,
			pb.Workload_SaaEnrollmentResponse_ERROR_INVALID_BASE_SETUP,
		},
		SetupStatus: pb.Workload_SaaEnrollmentResponse_STATUS_COMPLETE.Enum(),
	}

	out := Workload_SaaEnrollmentResponse_FromProto(mapCtx, in)

	if mapCtx.Err() != nil {
		t.Fatalf("unexpected error: %v", mapCtx.Err())
	}

	if out == nil {
		t.Fatal("expected non-nil output")
	}

	expectedErrors := []string{"SETUP_ERROR_UNSPECIFIED", "ERROR_INVALID_BASE_SETUP"}

	if !reflect.DeepEqual(out.SetupErrors, expectedErrors) {
		t.Errorf("expected SetupErrors %v, got %v", expectedErrors, out.SetupErrors)
	}

	if out.SetupStatus == nil || *out.SetupStatus != "STATUS_COMPLETE" {
		t.Errorf("expected SetupStatus %v, got %v", "STATUS_COMPLETE", out.SetupStatus)
	}
}
