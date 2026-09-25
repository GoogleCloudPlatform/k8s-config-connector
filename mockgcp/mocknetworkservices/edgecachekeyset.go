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

type EdgeCacheKeysetsServer struct {
	*MockService
	edgecachepb.UnimplementedEdgeCacheKeysetsServerServer
}

type edgeCacheKeysetName struct {
	Project  *projects.ProjectData
	Location string
	ID       string
}

func (n *edgeCacheKeysetName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/edgeCacheKeysets/" + n.ID
}

func (s *MockService) parseEdgeCacheKeysetName(name string) (*edgeCacheKeysetName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "edgeCacheKeysets" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}
		return &edgeCacheKeysetName{
			Project:  project,
			Location: tokens[3],
			ID:       tokens[5],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "invalid edgeCacheKeyset name %q", name)
}

func (s *EdgeCacheKeysetsServer) GetEdgeCacheKeyset(ctx context.Context, req *edgecachepb.GetEdgeCacheKeysetRequest) (*edgecachepb.EdgeCacheKeyset, error) {
	name, err := s.parseEdgeCacheKeysetName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &edgecachepb.EdgeCacheKeyset{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}
	return obj, nil
}

func (s *EdgeCacheKeysetsServer) ListEdgeCacheKeysets(ctx context.Context, req *edgecachepb.ListEdgeCacheKeysetsRequest) (*edgecachepb.ListEdgeCacheKeysetsResponse, error) {
	response := &edgecachepb.ListEdgeCacheKeysetsResponse{}

	prefixName, err := s.parseEdgeCacheKeysetName(req.Parent + "/edgeCacheKeysets/placeholder")
	if err != nil {
		return nil, err
	}
	prefix := strings.TrimSuffix(prefixName.String(), "placeholder")

	findKind := (&edgecachepb.EdgeCacheKeyset{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, findKind, storage.ListOptions{
		Prefix: prefix,
	}, func(obj proto.Message) error {
		keyset := obj.(*edgecachepb.EdgeCacheKeyset)
		response.EdgeCacheKeysets = append(response.EdgeCacheKeysets, keyset)
		return nil
	}); err != nil {
		return nil, err
	}
	return response, nil
}

func (s *EdgeCacheKeysetsServer) CreateEdgeCacheKeyset(ctx context.Context, req *edgecachepb.CreateEdgeCacheKeysetRequest) (*longrunningpb.Operation, error) {
	reqName := req.Parent + "/edgeCacheKeysets/" + req.EdgeCacheKeysetId
	name, err := s.parseEdgeCacheKeysetName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	obj := proto.Clone(req.EdgeCacheKeyset).(*edgecachepb.EdgeCacheKeyset)
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

func (s *EdgeCacheKeysetsServer) PatchEdgeCacheKeyset(ctx context.Context, req *edgecachepb.PatchEdgeCacheKeysetRequest) (*longrunningpb.Operation, error) {
	reqName := req.GetName()
	if reqName == "" && req.GetEdgeCacheKeyset() != nil {
		reqName = req.GetEdgeCacheKeyset().GetName()
	}
	name, err := s.parseEdgeCacheKeysetName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	obj := &edgecachepb.EdgeCacheKeyset{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		reqObj := req.GetEdgeCacheKeyset()
		reqObj.CreateTime = obj.CreateTime
		reqObj.UpdateTime = timestamppb.New(now)
		reqObj.Name = obj.Name
		obj = reqObj
	} else {
		for _, path := range paths {
			switch path {
			case "description":
				obj.Description = req.GetEdgeCacheKeyset().GetDescription()
			case "labels":
				obj.Labels = req.GetEdgeCacheKeyset().GetLabels()
			case "publicKeys":
				obj.PublicKeys = req.GetEdgeCacheKeyset().GetPublicKeys()
			case "validationSharedKeys":
				obj.ValidationSharedKeys = req.GetEdgeCacheKeyset().GetValidationSharedKeys()
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

func (s *EdgeCacheKeysetsServer) DeleteEdgeCacheKeyset(ctx context.Context, req *edgecachepb.DeleteEdgeCacheKeysetRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseEdgeCacheKeysetName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	deleted := &edgecachepb.EdgeCacheKeyset{}
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
