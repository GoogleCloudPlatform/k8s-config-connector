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

func (s *MapManagementServer) GetStyleConfig(ctx context.Context, req *pb.GetStyleConfigRequest) (*pb.StyleConfig, error) {
	name, err := s.parseStyleConfigName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.StyleConfig{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *MapManagementServer) CreateStyleConfig(ctx context.Context, req *pb.CreateStyleConfigRequest) (*pb.StyleConfig, error) {
	tokens := strings.Split(req.GetParent(), "/")
	if len(tokens) != 2 || tokens[0] != "projects" {
		return nil, status.Errorf(codes.InvalidArgument, "parent %q is malformed", req.GetParent())
	}

	project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
	if err != nil {
		return nil, err
	}

	displayName := req.GetStyleConfig().GetDisplayName()
	hash := md5.Sum([]byte(displayName))
	styleID := hex.EncodeToString(hash[:12])

	name := &styleConfigName{
		Project:     project,
		StyleConfig: styleID,
	}
	fqn := name.String()

	now := time.Now()
	obj := proto.Clone(req.GetStyleConfig()).(*pb.StyleConfig)
	obj.Name = fqn
	obj.StyleId = styleID
	obj.JsonStyleSheet = "null"
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *MapManagementServer) UpdateStyleConfig(ctx context.Context, req *pb.UpdateStyleConfigRequest) (*pb.StyleConfig, error) {
	if req.GetStyleConfig() == nil {
		return nil, status.Errorf(codes.InvalidArgument, "style_config is required")
	}

	name, err := s.parseStyleConfigName(req.GetStyleConfig().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.StyleConfig{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	if req.GetUpdateMask() != nil && len(req.GetUpdateMask().GetPaths()) > 0 {
		if err := fields.UpdateByFieldMask(obj, req.GetStyleConfig(), req.GetUpdateMask().GetPaths()); err != nil {
			return nil, err
		}
	} else {
		proto.Merge(obj, req.GetStyleConfig())
	}

	obj.Name = fqn
	obj.StyleId = name.StyleConfig
	obj.JsonStyleSheet = "null"
	obj.UpdateTime = timestamppb.New(time.Now())

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *MapManagementServer) DeleteStyleConfig(ctx context.Context, req *pb.DeleteStyleConfigRequest) (*emptypb.Empty, error) {
	name, err := s.parseStyleConfigName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	oldObj := &pb.StyleConfig{}
	if err := s.storage.Delete(ctx, fqn, oldObj); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *MapManagementServer) ListStyleConfigs(ctx context.Context, req *pb.ListStyleConfigsRequest) (*pb.ListStyleConfigsResponse, error) {
	tokens := strings.Split(req.GetParent(), "/")
	if len(tokens) != 2 || tokens[0] != "projects" {
		return nil, status.Errorf(codes.InvalidArgument, "parent %q is malformed", req.GetParent())
	}

	project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
	if err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%d/styleConfigs/", project.Number)
	var styleConfigs []*pb.StyleConfig
	if err := s.storage.List(ctx, (&pb.StyleConfig{}).ProtoReflect().Descriptor(), storage.ListOptions{Prefix: prefix}, func(obj proto.Message) error {
		styleConfigs = append(styleConfigs, obj.(*pb.StyleConfig))
		return nil
	}); err != nil {
		return nil, err
	}

	return &pb.ListStyleConfigsResponse{
		StyleConfigs: styleConfigs,
	}, nil
}
