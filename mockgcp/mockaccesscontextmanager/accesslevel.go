// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package mockaccesscontextmanager

import (
	"context"

	"google.golang.org/protobuf/proto"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/identity/accesscontextmanager/v1"
	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AccessContextManagerV1 struct {
	*MockService
	pb.UnimplementedAccessContextManagerServer
}

func (s *AccessContextManagerV1) GetAccessLevel(ctx context.Context, req *pb.GetAccessLevelRequest) (*pb.AccessLevel, error) {
	name, err := s.parseAccessLevelName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.AccessLevel{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *AccessContextManagerV1) CreateAccessLevel(ctx context.Context, req *pb.CreateAccessLevelRequest) (*longrunning.Operation, error) {
	reqName := req.GetAccessLevel().GetName()
	name, err := s.parseAccessLevelName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.AccessLevel).(*pb.AccessLevel)
	obj.Name = fqn

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return s.operations.NewLRO(ctx)
}

func (s *AccessContextManagerV1) UpdateAccessLevel(ctx context.Context, req *pb.UpdateAccessLevelRequest) (*longrunning.Operation, error) {
	reqName := req.GetAccessLevel().GetName()

	name, err := s.parseAccessLevelName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.AccessLevel{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// Required. A list of fields to be updated in this request.
	paths := req.GetUpdateMask().GetPaths()

	// TODO: Some sort of helper for fieldmask?
	for _, path := range paths {
		switch path {
		case "title":
			obj.Title = req.GetAccessLevel().GetTitle()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
		}
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return s.operations.NewLRO(ctx)
}

func (s *AccessContextManagerV1) DeleteAccessLevel(ctx context.Context, req *pb.DeleteAccessLevelRequest) (*longrunning.Operation, error) {
	name, err := s.parseAccessLevelName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	oldObj := &pb.AccessLevel{}
	if err := s.storage.Delete(ctx, fqn, oldObj); err != nil {
		return nil, err
	}

	return s.operations.NewLRO(ctx)
}
