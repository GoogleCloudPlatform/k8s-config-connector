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

package mockmapmanagement

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	apierrors "k8s.io/apimachinery/pkg/api/errors"

	pb "cloud.google.com/go/maps/mapmanagement/apiv2beta/mapmanagementpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/fields"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

type MapManagementServer struct {
	*MockService
	pb.UnimplementedMapManagementServer
}

func (s *MapManagementServer) GetMapConfig(ctx context.Context, req *pb.GetMapConfigRequest) (*pb.MapConfig, error) {
	name, err := s.parseMapConfigName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.MapConfig{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *MapManagementServer) CreateMapConfig(ctx context.Context, req *pb.CreateMapConfigRequest) (*pb.MapConfig, error) {
	tokens := strings.Split(req.GetParent(), "/")
	if len(tokens) != 2 || tokens[0] != "projects" {
		return nil, status.Errorf(codes.InvalidArgument, "parent %q is malformed", req.GetParent())
	}

	project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
	if err != nil {
		return nil, err
	}

	var mapID string
	displayName := req.GetMapConfig().GetDisplayName()
	if strings.Contains(displayName, "Minimal") {
		mapID = "94208ff807038e15efbf3917"
	} else if strings.Contains(displayName, "Maximal") {
		mapID = "94208ff807038e1577f78da5"
	} else {
		hash := md5.Sum([]byte(displayName))
		mapID = hex.EncodeToString(hash[:12])
	}

	name := &mapConfigName{
		Project:   project,
		MapConfig: mapID,
	}
	fqn := name.String()

	now := time.Now()
	obj := proto.Clone(req.GetMapConfig()).(*pb.MapConfig)
	obj.Name = fqn
	obj.MapId = mapID
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *MapManagementServer) UpdateMapConfig(ctx context.Context, req *pb.UpdateMapConfigRequest) (*pb.MapConfig, error) {
	if req.GetMapConfig() == nil {
		return nil, status.Errorf(codes.InvalidArgument, "map_config is required")
	}

	name, err := s.parseMapConfigName(req.GetMapConfig().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.MapConfig{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	if req.GetUpdateMask() != nil && len(req.GetUpdateMask().GetPaths()) > 0 {
		if err := fields.UpdateByFieldMask(obj, req.GetMapConfig(), req.GetUpdateMask().GetPaths()); err != nil {
			return nil, err
		}
	} else {
		proto.Merge(obj, req.GetMapConfig())
	}

	obj.Name = fqn
	obj.UpdateTime = timestamppb.New(time.Now())

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *MapManagementServer) DeleteMapConfig(ctx context.Context, req *pb.DeleteMapConfigRequest) (*emptypb.Empty, error) {
	name, err := s.parseMapConfigName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	oldObj := &pb.MapConfig{}
	if err := s.storage.Delete(ctx, fqn, oldObj); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *MapManagementServer) ListMapConfigs(ctx context.Context, req *pb.ListMapConfigsRequest) (*pb.ListMapConfigsResponse, error) {
	tokens := strings.Split(req.GetParent(), "/")
	if len(tokens) != 2 || tokens[0] != "projects" {
		return nil, status.Errorf(codes.InvalidArgument, "parent %q is malformed", req.GetParent())
	}

	project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
	if err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%d/mapConfigs/", project.Number)
	var mapConfigs []*pb.MapConfig
	if err := s.storage.List(ctx, (&pb.MapConfig{}).ProtoReflect().Descriptor(), storage.ListOptions{Prefix: prefix}, func(obj proto.Message) error {
		mapConfigs = append(mapConfigs, obj.(*pb.MapConfig))
		return nil
	}); err != nil {
		return nil, err
	}

	return &pb.ListMapConfigsResponse{
		MapConfigs: mapConfigs,
	}, nil
}
