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

package mocksql

import (
	"context"
	"fmt"
	"time"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/sql/v1beta4"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
	"k8s.io/apimachinery/pkg/util/uuid"
	"k8s.io/klog/v2"
)

type operations struct {
	storage storage.Storage
}

func (s *operations) startLRO(ctx context.Context, op *pb.Operation, obj proto.Message, callback func() (proto.Message, error)) (*pb.Operation, error) {
	log := klog.FromContext(ctx)

	now := time.Now()

	switch obj := obj.(type) {
	case *pb.DatabaseInstance:
		if op.OperationType == pb.Operation_CREATE_REPLICA {
			op.TargetId = obj.MasterInstanceName
			op.TargetLink = fmt.Sprintf("https://sqladmin.googleapis.com/sql/v1beta4/projects/%s/instances/%s", op.TargetProject, op.TargetId)
		} else {
			op.TargetId = obj.Name
			op.TargetLink = fmt.Sprintf("https://sqladmin.googleapis.com/sql/v1beta4/projects/%s/instances/%s", op.TargetProject, op.TargetId)
		}
	case *pb.User:
		op.TargetId = obj.Instance
		op.TargetLink = fmt.Sprintf("https://sqladmin.googleapis.com/sql/v1beta4/projects/%s/instances/%s", obj.Project, obj.Instance)
	default:
		klog.Fatalf("unhandled type %T", obj)
	}

	op.Kind = "sql#operation"
	op.User = "user@example.com" // TODO
	if op.Status == 0 {
		op.Status = pb.Operation_PENDING
	}

	if op.Status == pb.Operation_RUNNING {
		// StartTime is set when transitions to RUNNING
		op.StartTime = timestamppb.New(now)
	}
	// if op.Status == pb.Operation_DONE {
	// 	// StartTime is set when transitions to RUNNING
	// 	op.StartTime = timestamppb.New(now)
	// 	op.EndTime = timestamppb.New(now)
	// }

	op.InsertTime = timestamppb.New(now)
	op.Name = string(uuid.NewUUID())

	fqn := "projects/" + op.TargetProject + "/operations/" + op.Name
	op.SelfLink = "https://sqladmin.googleapis.com/sql/v1beta4/" + fqn

	log.Info("storing operation", "fqn", fqn)
	if err := s.storage.Create(ctx, fqn, op); err != nil {
		return nil, status.Errorf(codes.Internal, "error creating LRO: %v", err)
	}

	go func() {
		result, err := callback()
		finished := &pb.Operation{}
		if err2 := s.storage.Get(ctx, fqn, finished); err2 != nil {
			klog.Warningf("error getting LRO: %v", err2)
			return
		}

		finished.Status = pb.Operation_DONE
		finished.StartTime = timestamppb.New(time.Now())
		finished.EndTime = timestamppb.New(time.Now())

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
