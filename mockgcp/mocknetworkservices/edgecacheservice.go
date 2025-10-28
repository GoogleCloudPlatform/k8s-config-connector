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

// +tool:mockgcp-support
// proto.service: google.cloud.networkservices.v1.NetworkServices
// proto.message: google.cloud.networkservices.v1.EdgeCacheService

package mocknetworkservices

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/google/cloud/networkservices/v1"
	commonpb "cloud.google.com/go/networkservices/apiv1/networkservicespb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

func (s *NetworkServicesServer) GetEdgeCacheService(ctx context.Context, req *pb.GetEdgeCacheServiceRequest) (*pb.EdgeCacheService, error) {
	name, err := s.parseEdgeCacheServiceName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.EdgeCacheService{}
	obj.Name = fqn
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *NetworkServicesServer) ListEdgeCacheServices(ctx context.Context, req *pb.ListEdgeCacheServicesRequest) (*pb.ListEdgeCacheServicesResponse, error) {
	response := &pb.ListEdgeCacheServicesResponse{}

	findKind := (&pb.EdgeCacheService{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, findKind, storage.ListOptions{
		Prefix: req.Parent + "/edgeCacheServices/",
	}, func(obj proto.Message) error {
		edgeCacheService := obj.(*pb.EdgeCacheService)
		response.EdgeCacheServices = append(response.EdgeCacheServices, edgeCacheService)
		return nil
	}); err != nil {
		return nil, err
	}
	return response, nil
}

func (s *NetworkServicesServer) CreateEdgeCacheService(ctx context.Context, req *pb.CreateEdgeCacheServiceRequest) (*longrunningpb.Operation, error) {
	reqName := req.Parent + "/edgeCacheServices/" + req.EdgeCacheServiceId
	name, err := s.parseEdgeCacheServiceName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	obj := proto.Clone(req.EdgeCacheService).(*pb.EdgeCacheService)
	obj.Name = fqn

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &commonpb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		EndTime:    timestamppb.New(now),
		Target:     name.String(),
		Verb:       "create",
		ApiVersion: "v1",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		result := proto.Clone(obj).(*pb.EdgeCacheService)
		result.Name = fqn
		return result, nil
	})
}

func (s *NetworkServicesServer) UpdateEdgeCacheService(ctx context.Context, req *pb.UpdateEdgeCacheServiceRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseEdgeCacheServiceName(req.EdgeCacheService.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	existing := &pb.EdgeCacheService{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		return nil, err
	}

	now := time.Now()

	updated := proto.Clone(req.EdgeCacheService).(*pb.EdgeCacheService)
	updated.Name = fqn

	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &commonpb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		EndTime:    timestamppb.New(now),
		Target:     name.String(),
		Verb:       "update",
		ApiVersion: "v1",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		result := proto.Clone(updated).(*pb.EdgeCacheService)
		result.Name = fqn
		return result, nil
	})
}

func (s *NetworkServicesServer) DeleteEdgeCacheService(ctx context.Context, req *pb.DeleteEdgeCacheServiceRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseEdgeCacheServiceName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.EdgeCacheService{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	now := time.Now()
	lroMetadata := &commonpb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		EndTime:    timestamppb.New(now),
		Target:     name.String(),
		Verb:       "delete",
		ApiVersion: "v1",
	}
	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.DoneLRO(ctx, lroPrefix, lroMetadata, &emptypb.Empty{})
}

type edgeCacheServiceName struct {
	Project              *projects.ProjectData
	Location             string
	EdgeCacheServiceName string
}

func (n *edgeCacheServiceName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/edgeCacheServices/" + n.EdgeCacheServiceName
}

// parseEdgeCacheServiceName parses a string into an edgeCacheServiceName.
// The expected form is `projects/*/locations/*/edgeCacheServices/*`.
func (s *NetworkServicesServer) parseEdgeCacheServiceName(name string) (*edgeCacheServiceName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "edgeCacheServices" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &edgeCacheServiceName{
			Project:              project,
			Location:             tokens[3],
			EdgeCacheServiceName: tokens[5],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
