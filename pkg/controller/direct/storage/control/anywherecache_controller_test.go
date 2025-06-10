// Copyright 2025 Google LLC
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

package storagecontrol

import (
	"context"
	"fmt"
	"testing"

	pb "cloud.google.com/go/storage/control/apiv2/controlpb"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/storage/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	v1 "k8s.io/api/core/v1"
)

type ProtoMatcher struct {
	msg proto.Message
}

func ProtoEq(msg proto.Message) gomock.Matcher {
	return &ProtoMatcher{msg: msg}
}

func (m *ProtoMatcher) Matches(x interface{}) bool {
	if x, ok := x.(proto.Message); ok {
		return proto.Equal(m.msg, x)
	}
	return false
}

func (m *ProtoMatcher) String() string {
	return fmt.Sprintf("is equal to %v", m.msg)
}

func TestIsStateChangeRequested(t *testing.T) {
	tests := []struct {
		name         string
		actualState  string
		desiredState *string // Use pointer to simulate nil for default
		expected     bool
	}{
		{
			name:         "states are different (running to paused)",
			actualState:  anywhereCacheStateRunning,
			desiredState: direct.LazyPtr(anywhereCacheStatePaused),
			expected:     true,
		},
		{
			name:         "states are different (paused to running)",
			actualState:  anywhereCacheStatePaused,
			desiredState: direct.LazyPtr(anywhereCacheStateRunning),
			expected:     true,
		},
		{
			name:         "states are different (disabled to running)",
			actualState:  anywhereCacheStateDisabled,
			desiredState: direct.LazyPtr(anywhereCacheStateRunning),
			expected:     true,
		},
		{
			name:         "states are the same (running to running)",
			actualState:  anywhereCacheStateRunning,
			desiredState: direct.LazyPtr(anywhereCacheStateRunning),
			expected:     false,
		},
		{
			name:         "states are the same (paused to paused)",
			actualState:  anywhereCacheStatePaused,
			desiredState: direct.LazyPtr(anywhereCacheStatePaused),
			expected:     false,
		},
		{
			name:         "desired state is nil (defaults to running) and actual is running",
			actualState:  anywhereCacheStateRunning,
			desiredState: nil,
			expected:     false,
		},
		{
			name:         "desired state is nil (defaults to running) and actual is paused",
			actualState:  anywhereCacheStatePaused,
			desiredState: nil,
			expected:     true,
		},
		{
			name:         "desired state is nil (defaults to running) and actual is creating",
			actualState:  anywhereCacheStateCreating,
			desiredState: nil,
			expected:     true,
		},
		{
			name:         "actual state is creating, desired is running",
			actualState:  anywhereCacheStateCreating,
			desiredState: direct.LazyPtr(anywhereCacheStateRunning),
			expected:     true,
		},
		{
			name:         "actual state is running, desired is creating (invalid but tests comparison)",
			actualState:  anywhereCacheStateRunning,
			desiredState: direct.LazyPtr(anywhereCacheStateCreating),
			expected:     true,
		},
		{
			name:         "actual state is invalid, desired is running",
			actualState:  anywhereCacheStateInvalid,
			desiredState: direct.LazyPtr(anywhereCacheStateRunning),
			expected:     true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			adapter := &AnywhereCacheAdapter{
				actual: &pb.AnywhereCache{
					State: tc.actualState,
				},
				desired: &krm.StorageAnywhereCache{
					Spec: krm.StorageAnywhereCacheSpec{
						DesiredState: tc.desiredState,
					},
				},
			}

			if got := adapter.IsStateChangeRequested(); got != tc.expected {
				t.Errorf("IsStateChangeRequested() got = %v, want %v for actualState '%s' and desiredState '%v'", got, tc.expected, tc.actualState, tc.desiredState)
			}
		})
	}
}

func TestAnywhereCacheAdapter_Update(t *testing.T) {
	ctx := context.Background()
	projectID := "test-project"
	bucketID := "test-bucket"
	cacheID := "sample-id"
	zone := "us-central1-a"
	id := krm.GetAnywhereCacheIdentity(&krm.AnywhereCacheParent{BucketName: bucketID}, cacheID)
	anywhereCacheName := fmt.Sprintf("projects/_/buckets/%s/anywhereCaches/%s", bucketID, cacheID)

	// Helper function to get a desired KRM and actual object
	getDesiredKRM := func(desiredState string, admissionPolicy string, ttl string) *krm.StorageAnywhereCache {
		return &krm.StorageAnywhereCache{
			Spec: krm.StorageAnywhereCacheSpec{
				BucketRef:       &refs.StorageBucketRef{External: fmt.Sprintf("projects/%s/buckets/%s", projectID, bucketID)},
				Zone:            &zone,
				AdmissionPolicy: &admissionPolicy, // Desired value
				Ttl:             &ttl,
				DesiredState:    &desiredState,
			},
			Status: krm.StorageAnywhereCacheStatus{},
		}
	}
	getActualPB := func(mapCtx *direct.MapContext, state string, admissionPolicy string, ttl string, pendingUpdate bool) *pb.AnywhereCache {
		return &pb.AnywhereCache{
			Name:            anywhereCacheName,
			Zone:            zone,
			AdmissionPolicy: admissionPolicy,
			Ttl:             direct.StringDuration_ToProto(mapCtx, &ttl),
			State:           state,
			PendingUpdate:   pendingUpdate,
		}
	}
	getAdapter := func(mockClient *MockAnywhereCacheAPI, desiredKrm *krm.StorageAnywhereCache, actualPb *pb.AnywhereCache) *AnywhereCacheAdapter {
		return &AnywhereCacheAdapter{
			id:        id,
			gcpClient: mockClient,
			desired:   desiredKrm,
			actual:    actualPb,
		}
	}

	t.Run("Metadata change requested, no state change, no previous update is in progress", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		// Arrange
		mapCtx := &direct.MapContext{}
		mockClient := NewMockAnywhereCacheAPI(ctrl)
		mockUpdateOp := NewMockDirectBaseUpdateOperation(ctrl)

		desiredKrm := getDesiredKRM(anywhereCacheStateRunning, "admit-on-first-miss", "86400s")

		desiredPb := StorageAnywhereCacheSpec_ToProto(mapCtx, &desiredKrm.Spec)
		desiredPb.Name = anywhereCacheName
		if mapCtx.Err() != nil {
			t.Fatalf("unexpected error: %v", mapCtx.Err())
		}

		actualPb := getActualPB(mapCtx, anywhereCacheStateRunning, "admit-on-second-miss", "10002s", false)
		if mapCtx.Err() != nil {
			t.Fatalf("unexpected error creating actualPb: %v", mapCtx.Err())
		}

		// Setup Mocks
		expectedUpdateReq := &pb.UpdateAnywhereCacheRequest{
			UpdateMask:    &fieldmaskpb.FieldMask{Paths: []string{"admission_policy", "ttl"}},
			AnywhereCache: desiredPb,
		}
		mockClient.EXPECT().UpdateAnywhereCache(ctx, ProtoEq(expectedUpdateReq), gomock.Any()).Return(nil, nil).Times(1)
		mockClient.EXPECT().GetAnywhereCache(ctx, ProtoEq(&pb.GetAnywhereCacheRequest{Name: anywhereCacheName}), gomock.Any()).Return(getActualPB(mapCtx, anywhereCacheStateRunning, "admit-on-first-miss", "86400s", true), nil).Times(1)
		mockUpdateOp.EXPECT().UpdateStatus(ctx, gomock.Any(), getReadyCondition(v1.ConditionFalse, k8s.Updating, k8s.UpdatingMessage)).Return(nil).Times(1)
		if mapCtx.Err() != nil {
			t.Fatalf("unexpected error setting mocks: %v", mapCtx.Err())
		}

		// Execute & Assert
		err := getAdapter(mockClient, desiredKrm, actualPb).UpdateCache(ctx, mockUpdateOp)
		assert.NoError(t, err)
	})

	t.Run("Cache is in creating state, so any updates should be delayed", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		// Arrange
		mapCtx := &direct.MapContext{}
		mockClient := NewMockAnywhereCacheAPI(ctrl)
		mockUpdateOp := NewMockDirectBaseUpdateOperation(ctrl)

		desiredKrm := getDesiredKRM(anywhereCacheStateRunning, "admit-on-second-miss", "56400s")

		actualPb := getActualPB(mapCtx, anywhereCacheStateCreating, "admit-on-first-miss", "86400s", false)
		if mapCtx.Err() != nil {
			t.Fatalf("unexpected error creating actualPb: %v", mapCtx.Err())
		}

		// Setup Mocks
		mockUpdateOp.EXPECT().UpdateStatus(ctx, gomock.Any(), getReadyCondition(v1.ConditionFalse, k8s.Creating, k8s.CreatingMessage)).Return(nil).Times(1)

		// Execute and Assert
		err := getAdapter(mockClient, desiredKrm, actualPb).UpdateCache(ctx, mockUpdateOp)
		assert.NoError(t, err)
	})

	t.Run("Metadata change requested on RUNNING cache, a previous update is pending, should delay update and set Updating status", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		// Arrange
		mapCtx := &direct.MapContext{}
		mockClient := NewMockAnywhereCacheAPI(ctrl)
		mockUpdateOp := NewMockDirectBaseUpdateOperation(ctrl)

		desiredKrm := getDesiredKRM(anywhereCacheStateRunning, "admit-on-second-first", "56000s")
		desiredPb := StorageAnywhereCacheSpec_ToProto(mapCtx, &desiredKrm.Spec)
		desiredPb.Name = anywhereCacheName
		if mapCtx.Err() != nil {
			t.Fatalf("unexpected error: %v", mapCtx.Err())
		}

		actualPb := getActualPB(mapCtx, anywhereCacheStateRunning, "admit-on-first-miss", "45000s", true)
		if mapCtx.Err() != nil {
			t.Fatalf("unexpected error creating actualPb: %v", mapCtx.Err())
		}

		// Setup Mocks
		mockUpdateOp.EXPECT().UpdateStatus(ctx, gomock.Any(), getReadyCondition(v1.ConditionFalse, k8s.Updating, k8s.UpdatingMessage)).Return(nil).Times(1)

		// Execute and Assert
		err := getAdapter(mockClient, desiredKrm, actualPb).UpdateCache(ctx, mockUpdateOp)
		assert.NoError(t, err)
	})

	t.Run("No changes: should do nothing and set UpToDate status", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		// Arrange
		mapCtx := &direct.MapContext{}
		mockClient := NewMockAnywhereCacheAPI(ctrl)
		mockUpdateOp := NewMockDirectBaseUpdateOperation(ctrl)

		desiredKrm := getDesiredKRM(anywhereCacheStateRunning, "admit-on-first-miss", "86400s")

		desiredPb := StorageAnywhereCacheSpec_ToProto(mapCtx, &desiredKrm.Spec)
		desiredPb.Name = anywhereCacheName
		if mapCtx.Err() != nil {
			t.Fatalf("unexpected error: %v", mapCtx.Err())
		}

		actualPb := getActualPB(mapCtx, anywhereCacheStateRunning, "admit-on-first-miss", "86400s", false)
		if mapCtx.Err() != nil {
			t.Fatalf("unexpected error creating actualPb: %v", mapCtx.Err())
		}

		// Setup Mocks
		mockUpdateOp.EXPECT().UpdateStatus(ctx, gomock.Any(), getReadyCondition(v1.ConditionTrue, k8s.UpToDate, k8s.UpToDateMessage)).Return(nil).Times(1)

		// Assert and Execute
		err := getAdapter(mockClient, desiredKrm, actualPb).UpdateCache(ctx, mockUpdateOp)
		assert.NoError(t, err)
	})

	t.Run("State change requested: RUNNING to PAUSED", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		// Arrange
		mapCtx := &direct.MapContext{}
		mockClient := NewMockAnywhereCacheAPI(ctrl)
		mockUpdateOp := NewMockDirectBaseUpdateOperation(ctrl)

		desiredKrm := getDesiredKRM(anywhereCacheStatePaused, "admit-on-first-miss", "86400s")

		actualPb := getActualPB(mapCtx, anywhereCacheStateRunning, "admit-on-first-miss", "86400s", false)
		if mapCtx.Err() != nil {
			t.Fatalf("unexpected error creating actualPb: %v", mapCtx.Err())
		}

		// Setup Mocks
		mockClient.EXPECT().PauseAnywhereCache(ctx, ProtoEq(&pb.PauseAnywhereCacheRequest{Name: anywhereCacheName})).Return(nil, nil).Times(1)
		mockClient.EXPECT().GetAnywhereCache(ctx, ProtoEq(&pb.GetAnywhereCacheRequest{Name: anywhereCacheName}), gomock.Any()).Return(getActualPB(mapCtx, anywhereCacheStatePaused, "admit-on-first-miss", "86400s", false), nil).Times(1)
		mockUpdateOp.EXPECT().UpdateStatus(ctx, gomock.Any(), getReadyCondition(v1.ConditionTrue, k8s.UpToDate, k8s.UpToDateMessage)).Return(nil).Times(1)
		if mapCtx.Err() != nil {
			t.Fatalf("unexpected error while setting mocks: %v", mapCtx.Err())
		}

		// Assert and Execute
		err := getAdapter(mockClient, desiredKrm, actualPb).UpdateCache(ctx, mockUpdateOp)
		assert.NoError(t, err)
	})

	t.Run("State change requested: RUNNING to DISABLED - a previous update is in progress", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		// Arrange
		mapCtx := &direct.MapContext{}
		mockClient := NewMockAnywhereCacheAPI(ctrl)
		mockUpdateOp := NewMockDirectBaseUpdateOperation(ctrl)

		desiredKrm := getDesiredKRM(anywhereCacheStateDisabled, "admit-on-first-miss", "86400s")

		actualPb := getActualPB(mapCtx, anywhereCacheStateRunning, "admit-on-first-miss", "86400s", true)
		if mapCtx.Err() != nil {
			t.Fatalf("unexpected error creating actualPb: %v", mapCtx.Err())
		}

		// Setup Mocks
		mockClient.EXPECT().DisableAnywhereCache(ctx, ProtoEq(&pb.DisableAnywhereCacheRequest{Name: anywhereCacheName})).Return(nil, nil).Times(1)
		mockClient.EXPECT().GetAnywhereCache(ctx, ProtoEq(&pb.GetAnywhereCacheRequest{Name: anywhereCacheName}), gomock.Any()).Return(getActualPB(mapCtx, anywhereCacheStateDisabled, "admit-on-first-miss", "86400s", true), nil).Times(1)
		mockUpdateOp.EXPECT().RequestRequeue().Return().Times(1)
		mockUpdateOp.EXPECT().UpdateStatus(ctx, gomock.Any(), getReadyCondition(v1.ConditionFalse, k8s.Updating, k8s.UpdatingMessage)).Return(nil).Times(1)
		if mapCtx.Err() != nil {
			t.Fatalf("unexpected error while setting mocks: %v", mapCtx.Err())
		}

		// Assert and Execute
		err := getAdapter(mockClient, desiredKrm, actualPb).UpdateCache(ctx, mockUpdateOp)
		assert.NoError(t, err)
	})

	t.Run("State change requested: DISABLED to PAUSED + Metadata Changes Requested", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		// Arrange
		mapCtx := &direct.MapContext{}
		mockClient := NewMockAnywhereCacheAPI(ctrl)
		mockUpdateOp := NewMockDirectBaseUpdateOperation(ctrl)

		desiredKrm := getDesiredKRM(anywhereCacheStatePaused, "admit-on-second-miss", "86400s")

		actualPb := getActualPB(mapCtx, anywhereCacheStateDisabled, "admit-on-first-miss", "86400s", true)
		if mapCtx.Err() != nil {
			t.Fatalf("unexpected error creating actualPb: %v", mapCtx.Err())
		}

		// Setup Mocks
		mockClient.EXPECT().ResumeAnywhereCache(ctx, ProtoEq(&pb.ResumeAnywhereCacheRequest{Name: anywhereCacheName})).Return(nil, nil).Times(1)
		mockClient.EXPECT().PauseAnywhereCache(ctx, ProtoEq(&pb.PauseAnywhereCacheRequest{Name: anywhereCacheName})).Return(nil, nil).Times(1)
		mockClient.EXPECT().GetAnywhereCache(ctx, ProtoEq(&pb.GetAnywhereCacheRequest{Name: anywhereCacheName}), gomock.Any()).Return(getActualPB(mapCtx, anywhereCacheStatePaused, "admit-on-first-miss", "86400s", true), nil).Times(1)
		mockUpdateOp.EXPECT().RequestRequeue().Return().Times(1)
		mockUpdateOp.EXPECT().UpdateStatus(ctx, gomock.Any(), getReadyCondition(v1.ConditionFalse, k8s.Updating, k8s.UpdatingMessage)).Return(nil).Times(1)
		if mapCtx.Err() != nil {
			t.Fatalf("unexpected error while setting mocks: %v", mapCtx.Err())
		}

		// Assert and Execute
		err := getAdapter(mockClient, desiredKrm, actualPb).UpdateCache(ctx, mockUpdateOp)
		assert.NoError(t, err)
	})

	t.Run("State change requested: DISABLED to RUNNING", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		// Arrange
		mapCtx := &direct.MapContext{}
		mockClient := NewMockAnywhereCacheAPI(ctrl)
		mockUpdateOp := NewMockDirectBaseUpdateOperation(ctrl)

		desiredKrm := getDesiredKRM(anywhereCacheStateRunning, "admit-on-first-miss", "86400s")

		actualPb := getActualPB(mapCtx, anywhereCacheStateDisabled, "admit-on-first-miss", "86400s", false)
		if mapCtx.Err() != nil {
			t.Fatalf("unexpected error creating actualPb: %v", mapCtx.Err())
		}

		// Setup Mocks
		mockClient.EXPECT().ResumeAnywhereCache(ctx, ProtoEq(&pb.ResumeAnywhereCacheRequest{Name: anywhereCacheName})).Return(nil, nil).Times(1)
		mockClient.EXPECT().GetAnywhereCache(ctx, ProtoEq(&pb.GetAnywhereCacheRequest{Name: anywhereCacheName}), gomock.Any()).Return(getActualPB(mapCtx, anywhereCacheStateRunning, "admit-on-first-miss", "86400s", false), nil).Times(1)
		mockUpdateOp.EXPECT().UpdateStatus(ctx, gomock.Any(), getReadyCondition(v1.ConditionTrue, k8s.UpToDate, k8s.UpToDateMessage)).Return(nil).Times(1)
		if mapCtx.Err() != nil {
			t.Fatalf("unexpected error while setting mocks: %v", mapCtx.Err())
		}

		// Assert and Execute
		err := getAdapter(mockClient, desiredKrm, actualPb).UpdateCache(ctx, mockUpdateOp)
		assert.NoError(t, err)
	})

	t.Run("Metadata change requested in DISABLED state: should return error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mapCtx := &direct.MapContext{}
		mockClient := NewMockAnywhereCacheAPI(ctrl)
		mockUpdateOp := NewMockDirectBaseUpdateOperation(ctrl)

		desiredKrm := getDesiredKRM(anywhereCacheStateDisabled, "admit-on-second-miss", "86400s")

		desiredPb := StorageAnywhereCacheSpec_ToProto(mapCtx, &desiredKrm.Spec)
		desiredPb.Name = anywhereCacheName
		if mapCtx.Err() != nil {
			t.Fatalf("unexpected error: %v", mapCtx.Err())
		}

		actualPb := getActualPB(mapCtx, anywhereCacheStateDisabled, "admit-on-first-miss", "86400s", false)
		if mapCtx.Err() != nil {
			t.Fatalf("unexpected error creating actualPb: %v", mapCtx.Err())
		}

		// Assert and Execute
		err := getAdapter(mockClient, desiredKrm, actualPb).UpdateCache(ctx, mockUpdateOp)
		assert.ErrorContains(t, err, "AnywhereCache cannot be updated in 'paused' or 'disabled' state. Please change the cache state to 'running' to apply metadata updates")
	})
}
