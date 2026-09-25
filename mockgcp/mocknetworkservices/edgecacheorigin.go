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
	"strings"
	"time"

	pb "cloud.google.com/go/networkservices/apiv1/networkservicespb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	edgecachepb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/networkservices/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type EdgeCacheOriginsServer struct {
	*MockService
	edgecachepb.UnimplementedEdgeCacheOriginsServerServer
}

type edgeCacheOriginName struct {
	Project  *projects.ProjectData
	Location string
	ID       string
}

func (n *edgeCacheOriginName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/edgeCacheOrigins/" + n.ID
}

func (s *MockService) parseEdgeCacheOriginName(name string) (*edgeCacheOriginName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "edgeCacheOrigins" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}
		return &edgeCacheOriginName{
			Project:  project,
			Location: tokens[3],
			ID:       tokens[5],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "invalid edgeCacheOrigin name %q", name)
}

func (s *EdgeCacheOriginsServer) GetEdgeCacheOrigin(ctx context.Context, req *edgecachepb.GetEdgeCacheOriginRequest) (*edgecachepb.EdgeCacheOrigin, error) {
	name, err := s.parseEdgeCacheOriginName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &edgecachepb.EdgeCacheOrigin{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}
	return obj, nil
}

func (s *EdgeCacheOriginsServer) ListEdgeCacheOrigins(ctx context.Context, req *edgecachepb.ListEdgeCacheOriginsRequest) (*edgecachepb.ListEdgeCacheOriginsResponse, error) {
	response := &edgecachepb.ListEdgeCacheOriginsResponse{}

	prefixName, err := s.parseEdgeCacheOriginName(req.Parent + "/edgeCacheOrigins/placeholder")
	if err != nil {
		return nil, err
	}
	prefix := strings.TrimSuffix(prefixName.String(), "placeholder")

	findKind := (&edgecachepb.EdgeCacheOrigin{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, findKind, storage.ListOptions{
		Prefix: prefix,
	}, func(obj proto.Message) error {
		origin := obj.(*edgecachepb.EdgeCacheOrigin)
		response.EdgeCacheOrigins = append(response.EdgeCacheOrigins, origin)
		return nil
	}); err != nil {
		return nil, err
	}
	return response, nil
}

func (s *EdgeCacheOriginsServer) CreateEdgeCacheOrigin(ctx context.Context, req *edgecachepb.CreateEdgeCacheOriginRequest) (*longrunningpb.Operation, error) {
	reqName := req.Parent + "/edgeCacheOrigins/" + req.EdgeCacheOriginId
	name, err := s.parseEdgeCacheOriginName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	obj := proto.Clone(req.EdgeCacheOrigin).(*edgecachepb.EdgeCacheOrigin)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix := "projects/" + name.Project.ID + "/locations/" + name.Location
	lroMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		Target:     fqn,
		Verb:       "create",
		ApiVersion: "v1",
	}

	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())
		return proto.Clone(obj), nil
	})
}

func (s *EdgeCacheOriginsServer) PatchEdgeCacheOrigin(ctx context.Context, req *edgecachepb.PatchEdgeCacheOriginRequest) (*longrunningpb.Operation, error) {
	reqName := req.GetName()
	if reqName == "" && req.GetEdgeCacheOrigin() != nil {
		reqName = req.GetEdgeCacheOrigin().GetName()
	}
	name, err := s.parseEdgeCacheOriginName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	obj := &edgecachepb.EdgeCacheOrigin{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		reqObj := req.GetEdgeCacheOrigin()
		reqObj.CreateTime = obj.CreateTime
		reqObj.UpdateTime = timestamppb.New(now)
		reqObj.Name = obj.Name
		obj = reqObj
	} else {
		for _, path := range paths {
			switch path {
			case "description":
				obj.Description = req.GetEdgeCacheOrigin().GetDescription()
			case "labels":
				obj.Labels = req.GetEdgeCacheOrigin().GetLabels()
			case "originAddress":
				obj.OriginAddress = req.GetEdgeCacheOrigin().GetOriginAddress()
			case "protocol":
				obj.Protocol = req.GetEdgeCacheOrigin().GetProtocol()
			case "port":
				obj.Port = req.GetEdgeCacheOrigin().GetPort()
			case "maxAttempts":
				obj.MaxAttempts = req.GetEdgeCacheOrigin().GetMaxAttempts()
			case "timeout":
				obj.Timeout = req.GetEdgeCacheOrigin().GetTimeout()
			case "retryConditions":
				obj.RetryConditions = req.GetEdgeCacheOrigin().GetRetryConditions()
			case "awsV4Authentication":
				obj.AwsV4Authentication = req.GetEdgeCacheOrigin().GetAwsV4Authentication()
			case "originOverrideAction":
				obj.OriginOverrideAction = req.GetEdgeCacheOrigin().GetOriginOverrideAction()
			default:
				return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q is not supported", path)
			}
		}
		obj.UpdateTime = timestamppb.New(now)
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix := "projects/" + name.Project.ID + "/locations/" + name.Location
	lroMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		Target:     fqn,
		Verb:       "update",
		ApiVersion: "v1",
	}

	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())
		return proto.Clone(obj), nil
	})
}

func (s *EdgeCacheOriginsServer) DeleteEdgeCacheOrigin(ctx context.Context, req *edgecachepb.DeleteEdgeCacheOriginRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseEdgeCacheOriginName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	deleted := &edgecachepb.EdgeCacheOrigin{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	lroPrefix := "projects/" + name.Project.ID + "/locations/" + name.Location
	lroMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		Target:     fqn,
		Verb:       "delete",
		ApiVersion: "v1",
	}

	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())
		return &emptypb.Empty{}, nil
	})
}
