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

package mockdesigncenter

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/designcenter/apiv1/designcenterpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/fields"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

type DesignCenterServer struct {
	*MockService
	pb.UnimplementedDesignCenterServer
}

func (s *DesignCenterServer) GetSpace(ctx context.Context, req *pb.GetSpaceRequest) (*pb.Space, error) {
	name, err := s.parseSpaceName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Space{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *DesignCenterServer) CreateSpace(ctx context.Context, req *pb.CreateSpaceRequest) (*pb.Space, error) {
	parent, err := s.parseParent(req.Parent)
	if err != nil {
		return nil, err
	}

	if req.SpaceId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "spaceId is required")
	}

	if req.Space == nil {
		return nil, status.Errorf(codes.InvalidArgument, "space is required")
	}

	spaceID := req.SpaceId
	fqn := fmt.Sprintf("%s/spaces/%s", parent.String(), spaceID)

	name, err := s.parseSpaceName(fqn)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	obj := proto.Clone(req.Space).(*pb.Space)
	obj.Name = name.String()
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *DesignCenterServer) UpdateSpace(ctx context.Context, req *pb.UpdateSpaceRequest) (*pb.Space, error) {
	if req.Space == nil {
		return nil, status.Errorf(codes.InvalidArgument, "space is required")
	}

	name, err := s.parseSpaceName(req.Space.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	existing := &pb.Space{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	updated := proto.Clone(existing).(*pb.Space)
	now := time.Now()
	updated.UpdateTime = timestamppb.New(now)

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask is required")
	}

	if err := fields.UpdateByFieldMask(updated, req.Space, paths); err != nil {
		return nil, err
	}

	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}

	return updated, nil
}

func (s *DesignCenterServer) DeleteSpace(ctx context.Context, req *pb.DeleteSpaceRequest) (*emptypb.Empty, error) {
	name, err := s.parseSpaceName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.Space{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *DesignCenterServer) ListSpaces(ctx context.Context, req *pb.ListSpacesRequest) (*pb.ListSpacesResponse, error) {
	parent, err := s.parseParent(req.Parent)
	if err != nil {
		return nil, err
	}

	prefix := parent.String() + "/spaces/"
	response := &pb.ListSpacesResponse{}

	findOpts := storage.ListOptions{
		Prefix: prefix,
	}
	spaceKind := (&pb.Space{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, spaceKind, findOpts, func(item proto.Message) error {
		obj := item.(*pb.Space)
		response.Spaces = append(response.Spaces, obj)
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}
