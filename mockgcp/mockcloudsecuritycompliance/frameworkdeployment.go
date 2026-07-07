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

package mockcloudsecuritycompliance

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	pb "cloud.google.com/go/cloudsecuritycompliance/apiv1/cloudsecuritycompliancepb"
	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
)

type deploymentServer struct {
	*MockService
	pb.UnimplementedDeploymentServer
}

func (s *deploymentServer) GetFrameworkDeployment(ctx context.Context, req *pb.GetFrameworkDeploymentRequest) (*pb.FrameworkDeployment, error) {
	name := req.GetName()
	if name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "name is required")
	}

	obj := &pb.FrameworkDeployment{}
	if err := s.storage.Get(ctx, name, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Requested entity was not found.")
		}
		return nil, err
	}

	return obj, nil
}

func (s *deploymentServer) CreateFrameworkDeployment(ctx context.Context, req *pb.CreateFrameworkDeploymentRequest) (*longrunningpb.Operation, error) {
	parent := req.GetParent()
	if parent == "" {
		return nil, status.Errorf(codes.InvalidArgument, "parent is required")
	}
	deploymentID := req.GetFrameworkDeploymentId()
	if deploymentID == "" {
		return nil, status.Errorf(codes.InvalidArgument, "framework_deployment_id is required")
	}

	fqn := parent + "/frameworkDeployments/" + deploymentID

	obj := proto.Clone(req.GetFrameworkDeployment()).(*pb.FrameworkDeployment)
	obj.Name = fqn

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	prefix := parent
	return s.operations.DoneLRO(ctx, prefix, nil, obj)
}

func (s *deploymentServer) DeleteFrameworkDeployment(ctx context.Context, req *pb.DeleteFrameworkDeploymentRequest) (*longrunningpb.Operation, error) {
	name := req.GetName()
	if name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "name is required")
	}

	obj := &pb.FrameworkDeployment{}
	if err := s.storage.Delete(ctx, name, obj); err != nil {
		return nil, err
	}

	return s.operations.DoneLRO(ctx, name, nil, obj)
}
