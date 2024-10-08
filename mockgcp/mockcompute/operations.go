// Copyright 2024 Google LLC
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

package mockcompute

import (
	"context"
	"fmt"
	"time"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
	"google.golang.org/protobuf/proto"
	"k8s.io/apimachinery/pkg/util/uuid"
	"k8s.io/klog/v2"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

type computeOperations struct {
	storage storage.Storage
}

func newComputeOperationsService(storage storage.Storage) *computeOperations {
	return &computeOperations{
		storage: storage,
	}
}

func (s *computeOperations) globalOperationFQN(projectID string, name string) string {
	return "projects/" + projectID + "/global/operations/" + name
}

func (s *computeOperations) regionalOperationFQN(projectID string, region string, name string) string {
	return "projects/" + projectID + "/regions/" + region + "/operations/" + name
}

// Deprecated: use startGlobalLRO
func (s *computeOperations) newLRO(ctx context.Context, projectID string) (*pb.Operation, error) {
	log := klog.FromContext(ctx)

	now := time.Now()
	millis := now.UnixMilli()
	id := string(uuid.NewUUID())
	nanos := now.UnixNano()

	op := &pb.Operation{}

	op.StartTime = PtrTo(formatTime(now))
	op.InsertTime = PtrTo(formatTime(now))
	op.Id = PtrTo(uint64(nanos))

	op.Progress = PtrTo(int32(0))

	name := fmt.Sprintf("operation-%d-%s", millis, id)
	op.Name = PtrTo(name)
	op.Kind = PtrTo("compute#operation")
	fqn := s.globalOperationFQN(projectID, name)
	op.SelfLink = PtrTo("https://www.googleapis.com/compute/v1/" + fqn)

	op.Status = PtrTo(pb.Operation_DONE)

	log.Info("storing operation", "fqn", fqn)
	if err := s.storage.Create(ctx, fqn, op); err != nil {
		return nil, err
	}
	return op, nil
}

func (s *computeOperations) startLRO0(ctx context.Context, op *pb.Operation, fqn string, callback func() (proto.Message, error)) (*pb.Operation, error) {
	log := klog.FromContext(ctx)

	now := time.Now()
	nanos := now.UnixNano()

	if op == nil {
		op = &pb.Operation{}
	}

	op.StartTime = PtrTo(formatTime(now))
	op.InsertTime = PtrTo(formatTime(now))
	op.Id = PtrTo(uint64(nanos))

	if op.Progress == nil {
		op.Progress = PtrTo(int32(0))
	}

	if op.Status == nil {
		op.Status = PtrTo(pb.Operation_RUNNING)
	}

	op.Kind = PtrTo("compute#operation")
	op.SelfLink = PtrTo("https://www.googleapis.com/compute/v1/" + fqn)

	log.Info("storing operation", "fqn", fqn)
	if err := s.storage.Create(ctx, fqn, op); err != nil {
		return nil, err
	}

	go func() {
		result, err := callback()
		finished := &pb.Operation{}
		if err2 := s.storage.Get(ctx, fqn, finished); err2 != nil {
			klog.Warningf("error getting LRO: %v", err2)
			return
		}

		finished.Progress = PtrTo(int32(100))
		finished.Status = PtrTo(pb.Operation_DONE)
		finished.EndTime = PtrTo(formatTime(time.Now()))

		if err != nil {
			klog.Warningf("TODO: handle LRO error %v", err)
		} else {
			klog.Warningf("TODO: handle LRO result %v", result)
		}
		if err := s.storage.Update(ctx, fqn, finished); err != nil {
			klog.Warningf("error updating LRO: %v", err)
			return
		}
	}()

	return op, nil
}

func (s *computeOperations) startRegionalLRO(ctx context.Context, projectID string, region string, op *pb.Operation, callback func() (proto.Message, error)) (*pb.Operation, error) {
	now := time.Now()
	millis := now.UnixMilli()
	id := uuid.NewUUID()

	name := fmt.Sprintf("operation-%d-%s", millis, id)
	fqn := s.regionalOperationFQN(projectID, region, name)

	op.Name = PtrTo(name)
	op.Region = PtrTo(fmt.Sprintf("https://www.googleapis.com/compute/v1/projects/%s/regions/%s", projectID, region))
	return s.startLRO0(ctx, op, fqn, callback)
}

func (s *computeOperations) startGlobalLRO(ctx context.Context, projectID string, op *pb.Operation, callback func() (proto.Message, error)) (*pb.Operation, error) {
	now := time.Now()
	millis := now.UnixMilli()
	id := uuid.NewUUID()

	name := fmt.Sprintf("operation-%d-%s", millis, id)
	fqn := s.globalOperationFQN(projectID, name)

	op.Name = PtrTo(name)
	return s.startLRO0(ctx, op, fqn, callback)
}

// Gets the latest state of a long-running operation.  Clients can use this
// method to poll the operation result at intervals as recommended by the API
// service.
func (s *computeOperations) getOperation(ctx context.Context, fqn string) (*pb.Operation, error) {
	op := &pb.Operation{}
	if err := s.storage.Get(ctx, fqn, op); err != nil {
		return nil, err
	}

	return op, nil
}

func formatTime(t time.Time) string {
	return t.Format(time.RFC3339)
}
