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

package bigqueryreservation

import (
	"context"
	"testing"

	pb "cloud.google.com/go/bigquery/reservation/apiv1/reservationpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigqueryreservation/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/lifecyclehandler"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestComputeAssignmentUpdate_SchedulingPolicy(t *testing.T) {
	int64Ptr := func(v int64) *int64 { return &v }

	tests := []struct {
		name         string
		desiredSpec  *krm.BigQueryReservationAssignmentSpec
		actual       *pb.Assignment
		wantPaths    []string
		wantMaxSlots *int64
		wantConcur   *int64
	}{
		{
			name: "desired.Spec.SchedulingPolicy == nil when actual.SchedulingPolicy is non-nil (no diff)",
			desiredSpec: &krm.BigQueryReservationAssignmentSpec{
				SchedulingPolicy: nil,
			},
			actual: &pb.Assignment{
				Name: "projects/test-project/locations/us-west2/reservations/res1/assignments/as1",
				SchedulingPolicy: &pb.SchedulingPolicy{
					Concurrency: int64Ptr(10),
					MaxSlots:    int64Ptr(2000),
				},
			},
			wantPaths: []string{},
		},
		{
			name: "partial desired.Spec.SchedulingPolicy: only concurrency changed, preserves maxSlots from actual",
			desiredSpec: &krm.BigQueryReservationAssignmentSpec{
				SchedulingPolicy: &krm.SchedulingPolicy{
					Concurrency: int64Ptr(20),
					MaxSlots:    nil,
				},
			},
			actual: &pb.Assignment{
				Name: "projects/test-project/locations/us-west2/reservations/res1/assignments/as1",
				SchedulingPolicy: &pb.SchedulingPolicy{
					Concurrency: int64Ptr(10),
					MaxSlots:    int64Ptr(2000),
				},
			},
			wantPaths:    []string{"scheduling_policy"},
			wantConcur:   int64Ptr(20),
			wantMaxSlots: int64Ptr(2000),
		},
		{
			name: "partial desired.Spec.SchedulingPolicy: only concurrency set, matches actual concurrency, preserves maxSlots (no diff)",
			desiredSpec: &krm.BigQueryReservationAssignmentSpec{
				SchedulingPolicy: &krm.SchedulingPolicy{
					Concurrency: int64Ptr(10),
					MaxSlots:    nil,
				},
			},
			actual: &pb.Assignment{
				Name: "projects/test-project/locations/us-west2/reservations/res1/assignments/as1",
				SchedulingPolicy: &pb.SchedulingPolicy{
					Concurrency: int64Ptr(10),
					MaxSlots:    int64Ptr(2000),
				},
			},
			wantPaths:    []string{},
			wantConcur:   int64Ptr(10),
			wantMaxSlots: int64Ptr(2000),
		},
		{
			name: "partial desired.Spec.SchedulingPolicy: only maxSlots changed, preserves concurrency from actual",
			desiredSpec: &krm.BigQueryReservationAssignmentSpec{
				SchedulingPolicy: &krm.SchedulingPolicy{
					Concurrency: nil,
					MaxSlots:    int64Ptr(4000),
				},
			},
			actual: &pb.Assignment{
				Name: "projects/test-project/locations/us-west2/reservations/res1/assignments/as1",
				SchedulingPolicy: &pb.SchedulingPolicy{
					Concurrency: int64Ptr(10),
					MaxSlots:    int64Ptr(2000),
				},
			},
			wantPaths:    []string{"scheduling_policy"},
			wantConcur:   int64Ptr(10),
			wantMaxSlots: int64Ptr(4000),
		},
		{
			name: "both concurrency and maxSlots specified and changed",
			desiredSpec: &krm.BigQueryReservationAssignmentSpec{
				SchedulingPolicy: &krm.SchedulingPolicy{
					Concurrency: int64Ptr(30),
					MaxSlots:    int64Ptr(5000),
				},
			},
			actual: &pb.Assignment{
				Name: "projects/test-project/locations/us-west2/reservations/res1/assignments/as1",
				SchedulingPolicy: &pb.SchedulingPolicy{
					Concurrency: int64Ptr(10),
					MaxSlots:    int64Ptr(2000),
				},
			},
			wantPaths:    []string{"scheduling_policy"},
			wantConcur:   int64Ptr(30),
			wantMaxSlots: int64Ptr(5000),
		},
		{
			name: "actual.SchedulingPolicy is nil and desired.Spec.SchedulingPolicy is set",
			desiredSpec: &krm.BigQueryReservationAssignmentSpec{
				SchedulingPolicy: &krm.SchedulingPolicy{
					Concurrency: int64Ptr(10),
					MaxSlots:    int64Ptr(2000),
				},
			},
			actual: &pb.Assignment{
				Name:             "projects/test-project/locations/us-west2/reservations/res1/assignments/as1",
				SchedulingPolicy: nil,
			},
			wantPaths:    []string{"scheduling_policy"},
			wantConcur:   int64Ptr(10),
			wantMaxSlots: int64Ptr(2000),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mapCtx := &direct.MapContext{}
			desiredPb := BigQueryReservationAssignmentSpec_ToProto(mapCtx, tc.desiredSpec)
			if mapCtx.Err() != nil {
				t.Fatalf("BigQueryReservationAssignmentSpec_ToProto failed: %v", mapCtx.Err())
			}

			paths := computeAssignmentUpdate(tc.desiredSpec, desiredPb, tc.actual, nil)

			if len(paths) != len(tc.wantPaths) {
				t.Fatalf("computeAssignmentUpdate() paths = %v, want %v", paths, tc.wantPaths)
			}
			for i, p := range paths {
				if p != tc.wantPaths[i] {
					t.Errorf("computeAssignmentUpdate() paths[%d] = %q, want %q", i, p, tc.wantPaths[i])
				}
			}

			if tc.wantConcur != nil {
				if desiredPb.SchedulingPolicy == nil || desiredPb.SchedulingPolicy.Concurrency == nil {
					t.Errorf("desiredPb.SchedulingPolicy.Concurrency = nil, want %v", *tc.wantConcur)
				} else if *desiredPb.SchedulingPolicy.Concurrency != *tc.wantConcur {
					t.Errorf("desiredPb.SchedulingPolicy.Concurrency = %v, want %v", *desiredPb.SchedulingPolicy.Concurrency, *tc.wantConcur)
				}
			}

			if tc.wantMaxSlots != nil {
				if desiredPb.SchedulingPolicy == nil || desiredPb.SchedulingPolicy.MaxSlots == nil {
					t.Errorf("desiredPb.SchedulingPolicy.MaxSlots = nil, want %v", *tc.wantMaxSlots)
				} else if *desiredPb.SchedulingPolicy.MaxSlots != *tc.wantMaxSlots {
					t.Errorf("desiredPb.SchedulingPolicy.MaxSlots = %v, want %v", *desiredPb.SchedulingPolicy.MaxSlots, *tc.wantMaxSlots)
				}
			}
		})
	}
}

func TestUpdateAssignment_NilSchedulingPolicy_NoClientCall(t *testing.T) {
	ctx := context.Background()
	int64Ptr := func(v int64) *int64 { return &v }

	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(krm.BigQueryReservationAssignmentGVK)
	u.SetName("test-assignment")
	u.SetNamespace("test-ns")

	scheme := runtime.NewScheme()
	if err := krm.AddToScheme(scheme); err != nil {
		t.Fatalf("failed to add to scheme: %v", err)
	}
	k8sClient := fake.NewClientBuilder().
		WithScheme(scheme).
		WithStatusSubresource(&krm.BigQueryReservationAssignment{}).
		Build()
	if err := k8sClient.Create(ctx, u); err != nil {
		t.Fatalf("failed to create object in fake client: %v", err)
	}

	obj := &krm.BigQueryReservationAssignment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-assignment",
			Namespace: "test-ns",
		},
		Spec: krm.BigQueryReservationAssignmentSpec{
			ReservationRef: &krm.ReservationRef{
				External: "projects/test-project/locations/us-west2/reservations/res1",
			},
		},
	}
	id, err := krm.NewAssignmentIdentity(ctx, k8sClient, obj)
	if err != nil {
		t.Fatalf("NewAssignmentIdentity failed: %v", err)
	}

	lh := lifecyclehandler.LifecycleHandler{Recorder: record.NewFakeRecorder(100)}
	updateOp := directbase.NewUpdateOperation(lh, k8sClient, u)

	adapter := &AssignmentAdapter{
		id: id,
		// gcpClient is nil. If UpdateAssignment is called, it would panic.
		gcpClient: nil,
		actual: &pb.Assignment{
			Name: "projects/test-project/locations/us-west2/reservations/res1/assignments/as1",
			SchedulingPolicy: &pb.SchedulingPolicy{
				Concurrency: int64Ptr(10),
				MaxSlots:    int64Ptr(2000),
			},
		},
	}

	desiredSpec := &krm.BigQueryReservationAssignmentSpec{
		SchedulingPolicy: nil,
	}

	// Should identify no diff and update status without calling gcpClient.UpdateAssignment
	if err := adapter.updateAssignment(ctx, updateOp, desiredSpec); err != nil {
		t.Fatalf("updateAssignment() unexpected error: %v", err)
	}
}
