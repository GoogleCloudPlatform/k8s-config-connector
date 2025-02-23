// Copyright 2022 Google LLC
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

package mockiam

import (
	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/iam/v1beta"
	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"k8s.io/klog/v2"
	"strings"
)

func (s *ServerV1Beta) GetWorkloadIdentityPool(ctx context.Context, req *pb.GetWorkloadIdentityPoolRequest) (*pb.WorkloadIdentityPool, error) {
	klog.Infof("reqName: %", req.Name)
	name, err := s.parseWorkloadIdentityPoolName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.WorkloadIdentityPool{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *ServerV1Beta) CreateWorkloadIdentityPool(ctx context.Context, req *pb.CreateWorkloadIdentityPoolRequest) (*longrunningpb.Operation, error) {
	workloadIdentityPoolName := req.GetParent() + "/workloadIdentityPools/" + req.GetWorkloadIdentityPoolId()
	name, err := s.parseWorkloadIdentityPoolName(workloadIdentityPoolName)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := proto.Clone(req.GetWorkloadIdentityPool()).(*pb.WorkloadIdentityPool)
	obj.Name = fqn
	obj.State = pb.WorkloadIdentityPool_ACTIVE

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}
	metadata := &pb.WorkloadIdentityPoolOperationMetadata{}
	return s.operations.StartLRO(ctx, fqn, metadata, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *ServerV1Beta) UpdateWorkloadIdentityPool(ctx context.Context, req *pb.UpdateWorkloadIdentityPoolRequest) (*longrunning.Operation, error) {
	reqName := req.GetWorkloadIdentityPool().GetName()

	name, err := s.parseWorkloadIdentityPoolName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.WorkloadIdentityPool{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// Required. A list of fields to be updated in this request.
	paths := req.GetUpdateMask().GetPaths()

	// TODO: Some sort of helper for fieldmask?
	for _, path := range paths {
		switch path {
		case "displayName":
			obj.DisplayName = req.GetWorkloadIdentityPool().GetDisplayName()
		case "description":
			obj.Description = req.GetWorkloadIdentityPool().GetDescription()
		case "state":
			obj.State = req.GetWorkloadIdentityPool().GetState()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
		}
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}
	return s.operations.NewLRO(ctx)

	/* opMetadata := &pb.WorkloadIdentityPoolOperationMetadata{}
	opPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, opPrefix, opMetadata, func() (proto.Message, error) {
		return obj, nil
	}) */
}

func (s *ServerV1Beta) DeleteWorkloadIdentityPool(ctx context.Context, req *pb.DeleteWorkloadIdentityPoolRequest) (*longrunning.Operation, error) {
	name, err := s.parseWorkloadIdentityPoolName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	oldObj := &pb.WorkloadIdentityPool{}
	if err := s.storage.Delete(ctx, fqn, oldObj); err != nil {
		return nil, err
	}

	opMetadata := &pb.WorkloadIdentityPoolOperationMetadata{}
	opPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, opPrefix, opMetadata, func() (proto.Message, error) {
		return &emptypb.Empty{}, nil
	})
}

type WorkloadIdentityPoolName struct {
	Project                  *projects.ProjectData
	Location                 string
	WorkloadIdentityPoolName string
}

func (n *WorkloadIdentityPoolName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/workloadIdentityPools/" + n.WorkloadIdentityPoolName
}

// parseWorkloadIdentityPoolName parses a string into a WorkloadIdentityPoolName.
// The expected form is projects/<projectID>/locations/global/workloadIdentityPools/<WorkloadIdentityPoolName>
func (s *MockService) parseWorkloadIdentityPoolName(name string) (*WorkloadIdentityPoolName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "workloadIdentityPools" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &WorkloadIdentityPoolName{
			Project:                  project,
			Location:                 tokens[3],
			WorkloadIdentityPoolName: tokens[5],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
