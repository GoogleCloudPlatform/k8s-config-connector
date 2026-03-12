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

package mockrun

import (
	"context"
	"fmt"
	"strings"

	iampb "cloud.google.com/go/iam/apiv1/iampb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/fields"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/run/v2"
	"github.com/google/uuid"
	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type workerPools struct {
	*MockService
	pb.UnimplementedWorkerPoolsServer
}

func (s *workerPools) GetWorkerPool(ctx context.Context, req *pb.GetWorkerPoolRequest) (*pb.WorkerPool, error) {
	name, err := s.parseWorkerPoolName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.WorkerPool{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *workerPools) CreateWorkerPool(ctx context.Context, req *pb.CreateWorkerPoolRequest) (*longrunning.Operation, error) {
	reqName := req.Parent + "/workerPools/" + req.WorkerPoolId
	name, err := s.parseWorkerPoolName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.WorkerPool).(*pb.WorkerPool)
	obj.Name = fqn
	obj.CreateTime = timestamppb.Now()
	obj.UpdateTime = timestamppb.Now()
	obj.Etag = fields.ComputeWeakEtag(obj)
	obj.Creator = "test@google.com"
	obj.LastModifier = "test@google.com"

	obj.Uid = uuid.NewString()
	obj.Generation = 1

	if obj.Template == nil {
		obj.Template = &pb.WorkerPoolRevisionTemplate{}
	}

	// Set TerminalCondition to Ready if not set (Mocking immediate success)
	if obj.TerminalCondition == nil {
		obj.TerminalCondition = &pb.Condition{
			LastTransitionTime: timestamppb.Now(),
			State:              pb.Condition_CONDITION_SUCCEEDED,
			Type:               "Ready",
		}
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}
	return s.operations.StartLRO(ctx, req.Parent, obj, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *workerPools) DeleteWorkerPool(ctx context.Context, req *pb.DeleteWorkerPoolRequest) (*longrunning.Operation, error) {
	name, err := s.parseWorkerPoolName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.WorkerPool{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}
	if err := s.storage.Delete(ctx, fqn, &pb.WorkerPool{}); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, lroPrefix, obj, func() (protoreflect.ProtoMessage, error) {
		return obj, nil
	})
}

func (s *workerPools) UpdateWorkerPool(ctx context.Context, req *pb.UpdateWorkerPoolRequest) (*longrunning.Operation, error) {
	name, err := s.parseWorkerPoolName(req.GetWorkerPool().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.WorkerPool{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	updated := req.GetWorkerPool()

	// Basic update logic
	if updated.Labels != nil {
		obj.Labels = updated.Labels
	}
	if updated.Annotations != nil {
		obj.Annotations = updated.Annotations
	}
	if updated.Template != nil {
		obj.Template = updated.Template
	}

	obj.UpdateTime = timestamppb.Now()
	obj.Generation++

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, lroPrefix, obj, func() (protoreflect.ProtoMessage, error) {
		return obj, nil
	})
}

func (s *workerPools) GetIamPolicy(ctx context.Context, req *iampb.GetIamPolicyRequest) (*iampb.Policy, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIamPolicy not implemented")
}

func (s *workerPools) SetIamPolicy(ctx context.Context, req *iampb.SetIamPolicyRequest) (*iampb.Policy, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetIamPolicy not implemented")
}

func (s *workerPools) TestIamPermissions(ctx context.Context, req *iampb.TestIamPermissionsRequest) (*iampb.TestIamPermissionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestIamPermissions not implemented")
}

type workerPoolName struct {
	Project    *projects.ProjectData
	Location   string
	WorkerPool string
}

func (n *workerPoolName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/workerPools/%s", n.Project.ID, n.Location, n.WorkerPool)
}

func (s *MockService) parseWorkerPoolName(name string) (*workerPoolName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "workerPools" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		n := &workerPoolName{
			Project:    project,
			Location:   tokens[3],
			WorkerPool: tokens[5],
		}
		return n, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
