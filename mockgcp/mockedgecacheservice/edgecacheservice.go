// Copyright 2025 Google LLC
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

package mockedgecacheservice

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/fields"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/edgecacheservice/v1"
)

func (s *EdgeCacheServicesServer) GetEdgeCacheService(ctx context.Context, req *pb.GetEdgeCacheServiceRequest) (*pb.EdgeCacheService, error) {
	name, err := s.parseEdgeCacheServiceName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.EdgeCacheService{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *EdgeCacheServicesServer) CreateEdgeCacheService(ctx context.Context, req *pb.CreateEdgeCacheServiceRequest) (*longrunningpb.Operation, error) {
	reqName := req.EdgeCacheServiceId
	if reqName == "" {
		return nil, status.Error(codes.InvalidArgument, "edge_cache_service_id is required")
	}

	name, err := s.parseEdgeCacheServiceName(fmt.Sprintf("%s/edgeCacheServices/%s", req.Parent, reqName))
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.EdgeCacheService).(*pb.EdgeCacheService)
	obj.Name = fqn
	obj.CreateTime = timestamppb.Now()
	obj.UpdateTime = timestamppb.Now()

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op, err := s.operations.StartLRO(ctx, "", nil, func() (proto.Message, error) {
		return obj, nil
	})
	if err != nil {
		return nil, err
	}
	return s.toLongRunning(op)
}

func (s *EdgeCacheServicesServer) UpdateEdgeCacheService(ctx context.Context, req *pb.UpdateEdgeCacheServiceRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseEdgeCacheServiceName(req.EdgeCacheService.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.EdgeCacheService{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// Clone the existing object to update it
	updated := proto.Clone(obj).(*pb.EdgeCacheService)
	updated.UpdateTime = timestamppb.Now()

	if req.UpdateMask != nil {
		if err := fields.UpdateByFieldMask(updated, req.EdgeCacheService, req.UpdateMask.Paths); err != nil {
			return nil, err
		}
	} else {
		// Full replacement if no mask?
		// Note: We should probably preserve system fields like create_time if replacement comes from user.
		// But here we rely on the input to be complete or merge logic?
		// Standard Update: if mask is empty, replace resource?
		// We'll assume full replace but keep immutable name/create_time.
		updated = proto.Clone(req.EdgeCacheService).(*pb.EdgeCacheService)
		updated.Name = obj.Name
		updated.CreateTime = obj.CreateTime
		updated.UpdateTime = timestamppb.Now()
	}

	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}

	op, err := s.operations.StartLRO(ctx, "", nil, func() (proto.Message, error) {
		return updated, nil
	})
	if err != nil {
		return nil, err
	}
	return s.toLongRunning(op)
}

func (s *EdgeCacheServicesServer) DeleteEdgeCacheService(ctx context.Context, req *pb.DeleteEdgeCacheServiceRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseEdgeCacheServiceName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	oldObj := &pb.EdgeCacheService{}
	if err := s.storage.Delete(ctx, fqn, oldObj); err != nil {
		return nil, err
	}

	op, err := s.operations.StartLRO(ctx, "", nil, func() (proto.Message, error) {
		return &emptypb.Empty{}, nil
	})
	if err != nil {
		return nil, err
	}
	return s.toLongRunning(op)
}

// toLongRunning converts the genproto LRO to the type expected by the interface.
func (s *EdgeCacheServicesServer) toLongRunning(op proto.Message) (*longrunningpb.Operation, error) {
	b, err := proto.Marshal(op)
	if err != nil {
		return nil, err
	}
	out := &longrunningpb.Operation{}
	if err := proto.Unmarshal(b, out); err != nil {
		return nil, err
	}
	return out, nil

}

type edgeCacheServiceName struct {
	Project   *projects.ProjectData
	Location  string
	ServiceID string
}

func (n *edgeCacheServiceName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/edgeCacheServices/%s", n.Project.ID, n.Location, n.ServiceID)
}

func (s *EdgeCacheServicesServer) parseEdgeCacheServiceName(name string) (*edgeCacheServiceName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "edgeCacheServices" {
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}
		return &edgeCacheServiceName{
			Project:   project,
			Location:  tokens[3],
			ServiceID: tokens[5],
		}, nil
	}
	return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("invalid EdgeCacheService name %q", name))
}
