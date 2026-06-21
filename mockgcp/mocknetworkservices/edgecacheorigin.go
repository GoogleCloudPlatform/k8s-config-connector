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

package mocknetworkservices

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/networkservices/v1"
	"google.golang.org/genproto/googleapis/longrunning"
)

type edgeCacheOriginName struct {
	Project             *projects.ProjectData
	Location            string
	EdgeCacheOriginName string
}

func (n *edgeCacheOriginName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/edgeCacheOrigins/%s", n.Project.ID, n.Location, n.EdgeCacheOriginName)
}

func (s *NetworkServicesServer) parseEdgeCacheOriginName(name string) (*edgeCacheOriginName, error) {
	if name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "name cannot be empty")
	}
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "edgeCacheOrigins" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}
		return &edgeCacheOriginName{
			Project:             project,
			Location:            tokens[3],
			EdgeCacheOriginName: tokens[5],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

func (s *NetworkServicesServer) GetEdgeCacheOrigin(ctx context.Context, req *pb.GetEdgeCacheOriginRequest) (*pb.EdgeCacheOrigin, error) {
	name, err := s.parseEdgeCacheOriginName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.EdgeCacheOrigin{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "EdgeCacheOrigin %q not found", name)
		}
		return nil, err
	}

	return obj, nil
}

func (s *NetworkServicesServer) CreateEdgeCacheOrigin(ctx context.Context, req *pb.CreateEdgeCacheOriginRequest) (*longrunning.Operation, error) {
	reqName := fmt.Sprintf("%s/edgeCacheOrigins/%s", req.GetParent(), req.GetEdgeCacheOriginId())
	name, err := s.parseEdgeCacheOriginName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	obj := proto.CloneOf(req.EdgeCacheOrigin)
	obj.Name = fqn

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		Target:     name.String(),
		Verb:       "create",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *NetworkServicesServer) UpdateEdgeCacheOrigin(ctx context.Context, req *pb.UpdateEdgeCacheOriginRequest) (*longrunning.Operation, error) {
	reqName := req.GetEdgeCacheOrigin().GetName()

	name, err := s.parseEdgeCacheOriginName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	existing := &pb.EdgeCacheOrigin{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		return nil, err
	}

	// Update existing fields (simple complete update for mock)
	obj := proto.CloneOf(req.EdgeCacheOrigin)
	obj.Name = fqn

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		Target:     name.String(),
		Verb:       "update",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *NetworkServicesServer) DeleteEdgeCacheOrigin(ctx context.Context, req *pb.DeleteEdgeCacheOriginRequest) (*longrunning.Operation, error) {
	name, err := s.parseEdgeCacheOriginName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	existing := &pb.EdgeCacheOrigin{}
	if err := s.storage.Delete(ctx, fqn, existing); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		Target:     name.String(),
		Verb:       "delete",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		return existing, nil
	})
}

func (s *NetworkServicesServer) ListEdgeCacheOrigins(ctx context.Context, req *pb.ListEdgeCacheOriginsRequest) (*pb.ListEdgeCacheOriginsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "ListEdgeCacheOrigins is not implemented in mock")
}
