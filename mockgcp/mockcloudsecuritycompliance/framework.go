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
	"google.golang.org/protobuf/types/known/emptypb"

	pb "cloud.google.com/go/cloudsecuritycompliance/apiv1/cloudsecuritycompliancepb"
)

func (s *configServer) GetFramework(ctx context.Context, req *pb.GetFrameworkRequest) (*pb.Framework, error) {
	name := req.GetName()
	if name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "name is required")
	}

	obj := &pb.Framework{}
	if err := s.storage.Get(ctx, name, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Requested entity was not found.")
		}
		return nil, err
	}

	return obj, nil
}

func (s *configServer) CreateFramework(ctx context.Context, req *pb.CreateFrameworkRequest) (*pb.Framework, error) {
	parent := req.GetParent()
	if parent == "" {
		return nil, status.Errorf(codes.InvalidArgument, "parent is required")
	}
	frameworkID := req.GetFrameworkId()
	if frameworkID == "" {
		return nil, status.Errorf(codes.InvalidArgument, "framework_id is required")
	}

	fqn := parent + "/frameworks/" + frameworkID

	obj := proto.Clone(req.GetFramework()).(*pb.Framework)
	obj.Name = fqn

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *configServer) UpdateFramework(ctx context.Context, req *pb.UpdateFrameworkRequest) (*pb.Framework, error) {
	reqObj := req.GetFramework()
	if reqObj == nil {
		return nil, status.Errorf(codes.InvalidArgument, "framework is required")
	}
	fqn := reqObj.GetName()
	if fqn == "" {
		return nil, status.Errorf(codes.InvalidArgument, "name is required")
	}

	existing := &pb.Framework{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		return nil, err
	}

	// Just update the storage with the new object
	obj := proto.Clone(reqObj).(*pb.Framework)

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *configServer) DeleteFramework(ctx context.Context, req *pb.DeleteFrameworkRequest) (*emptypb.Empty, error) {
	name := req.GetName()
	if name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "name is required")
	}

	obj := &pb.Framework{}
	if err := s.storage.Delete(ctx, name, obj); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
