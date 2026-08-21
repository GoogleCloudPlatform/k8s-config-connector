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
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/networkservices/v1"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

type edgeCacheServiceName struct {
	Project              *projects.ProjectData
	EdgeCacheServiceName string
}

func (n *edgeCacheServiceName) String() string {
	return "projects/" + n.Project.ID + "/locations/global/edgeCacheServices/" + n.EdgeCacheServiceName
}

func (s *NetworkServicesServer) parseEdgeCacheServiceName(name string) (*edgeCacheServiceName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[3] == "global" && tokens[4] == "edgeCacheServices" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		return &edgeCacheServiceName{
			Project:              project,
			EdgeCacheServiceName: tokens[5],
		}, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}

func (s *NetworkServicesServer) GetEdgeCacheService(ctx context.Context, req *pb.GetEdgeCacheServiceRequest) (*pb.EdgeCacheService, error) {
	name, err := s.parseEdgeCacheServiceName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.EdgeCacheService{}
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

	prefixName, err := s.parseEdgeCacheServiceName(req.Parent + "/edgeCacheServices/placeholder-name")
	if err != nil {
		return nil, err
	}
	prefix := strings.TrimSuffix(prefixName.String(), "placeholder-name")

	findKind := (&pb.EdgeCacheService{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, findKind, storage.ListOptions{
		Prefix: prefix,
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

	obj := proto.CloneOf(req.EdgeCacheService)
	obj.Name = fqn

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/global", name.Project.ID)
	lroMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		Target:     name.String(),
		Verb:       "create",
		ApiVersion: "v1",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())
		result := proto.CloneOf(obj)
		return result, nil
	})
}

func (s *NetworkServicesServer) UpdateEdgeCacheService(ctx context.Context, req *pb.UpdateEdgeCacheServiceRequest) (*longrunningpb.Operation, error) {
	reqName := req.GetEdgeCacheService().GetName()

	name, err := s.parseEdgeCacheServiceName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.EdgeCacheService{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		req.EdgeCacheService.Name = obj.Name
		obj = req.EdgeCacheService
	} else {
		// Field mask updates
		for _, path := range paths {
			switch path {
			case "description":
				obj.Description = req.GetEdgeCacheService().GetDescription()
			case "labels":
				obj.Labels = req.GetEdgeCacheService().GetLabels()
			case "disableHttp2":
				obj.DisableHttp2 = req.GetEdgeCacheService().GetDisableHttp2()
			case "disableQuic":
				obj.DisableQuic = req.GetEdgeCacheService().GetDisableQuic()
			case "edgeSecurityPolicy":
				obj.EdgeSecurityPolicy = req.GetEdgeCacheService().GetEdgeSecurityPolicy()
			case "edgeSslCertificates":
				obj.EdgeSslCertificates = req.GetEdgeCacheService().GetEdgeSslCertificates()
			case "logConfig":
				obj.LogConfig = req.GetEdgeCacheService().GetLogConfig()
			case "requireTls":
				obj.RequireTls = req.GetEdgeCacheService().GetRequireTls()
			case "routing":
				obj.Routing = req.GetEdgeCacheService().GetRouting()
			case "sslPolicy":
				obj.SslPolicy = req.GetEdgeCacheService().GetSslPolicy()
			default:
				return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
			}
		}
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	now := time.Now()
	lroPrefix := fmt.Sprintf("projects/%s/locations/global", name.Project.ID)
	lroMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		Target:     name.String(),
		Verb:       "update",
		ApiVersion: "v1",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())

		result := proto.CloneOf(obj)
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
	lroMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		Target:     name.String(),
		Verb:       "delete",
		ApiVersion: "v1",
	}
	lroPrefix := fmt.Sprintf("projects/%s/locations/global", name.Project.ID)
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())

		result := &emptypb.Empty{}
		return result, nil
	})
}
