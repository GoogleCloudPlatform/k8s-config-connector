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
// proto.service: google.identity.accesscontextmanager.v1.AccessContextManager
// proto.message: google.identity.accesscontextmanager.v1.AccessLevel

package mockaccesscontextmanager

import (
	"context"
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "cloud.google.com/go/accesscontextmanager/apiv1/accesscontextmanagerpb"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

func (s *AccessContextManagerV1) GetAccessLevel(ctx context.Context, req *pb.GetAccessLevelRequest) (*pb.AccessLevel, error) {
	name, err := s.parseAccessLevelName(req.GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.AccessLevel{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "AccessLevel %q not found.", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *AccessContextManagerV1) ListAccessLevels(ctx context.Context, req *pb.ListAccessLevelsRequest) (*pb.ListAccessLevelsResponse, error) {
	name, err := s.parseAccessLevelName(req.GetParent() + "/accessLevels/dummy")
	if err != nil {
		return nil, err
	}
	prefix := strings.TrimSuffix(name.String(), "dummy")

	response := &pb.ListAccessLevelsResponse{}
	kind := (&pb.AccessLevel{}).ProtoReflect().Descriptor()

	if err := s.storage.List(ctx, kind, storage.ListOptions{Prefix: prefix}, func(p proto.Message) error {
		accessLevel := p.(*pb.AccessLevel)
		response.AccessLevels = append(response.AccessLevels, accessLevel)
		return nil
	}); err != nil {
		return nil, status.Errorf(codes.Internal, "Error listing AccessLevels: %v", err)
	}

	return response, nil
}

func (s *AccessContextManagerV1) CreateAccessLevel(ctx context.Context, req *pb.CreateAccessLevelRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseAccessLevelName(req.GetAccessLevel().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := ProtoClone(req.GetAccessLevel())
	obj.Name = fqn

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	metadata := &pb.AccessContextManagerOperationMetadata{}
	lroPrefix := "operations/" + fqn + "/create/{{operationID}}"
	return s.operations.DoneLRO(ctx, lroPrefix, metadata, obj)
}

func (s *AccessContextManagerV1) UpdateAccessLevel(ctx context.Context, req *pb.UpdateAccessLevelRequest) (*longrunningpb.Operation, error) {
	if req.GetAccessLevel() == nil {
		return nil, status.Errorf(codes.InvalidArgument, "access_level is required for update")
	}
	name, err := s.parseAccessLevelName(req.GetAccessLevel().GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	existing := &pb.AccessLevel{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		return nil, err
	}

	updated := ProtoClone(existing)
	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask is required")
	}

	for _, path := range paths {
		switch path {
		case "title":
			updated.Title = req.GetAccessLevel().GetTitle()
		case "description":
			updated.Description = req.GetAccessLevel().GetDescription()
		case "basic":
			updated.Level = &pb.AccessLevel_Basic{Basic: req.GetAccessLevel().GetBasic()}
		case "basic.conditions":
			basic := existing.GetBasic()
			if basic == nil {
				basic = &pb.BasicLevel{}
			}
			basic.Conditions = req.GetAccessLevel().GetBasic().GetConditions()
			updated.Level = &pb.AccessLevel_Basic{Basic: basic}
		case "custom":
			updated.Level = &pb.AccessLevel_Custom{Custom: req.GetAccessLevel().GetCustom()}
		default:
			return nil, fmt.Errorf("UpdateAccessLevel: unsupported update_mask path (in mock) %q", path)
		}
	}

	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}

	metadata := &pb.AccessContextManagerOperationMetadata{}
	lroPrefix := "operations/" + fqn + "/update/{{operationID}}"
	return s.operations.DoneLRO(ctx, lroPrefix, metadata, updated)
}

func (s *AccessContextManagerV1) DeleteAccessLevel(ctx context.Context, req *pb.DeleteAccessLevelRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseAccessLevelName(req.GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	if err := s.storage.Delete(ctx, fqn, &pb.AccessLevel{}); err != nil {
		if status.Code(err) == codes.NotFound {
			// Deleting a non-existent object should succeed idempotently.
			metadata := &pb.AccessContextManagerOperationMetadata{}
			return s.operations.DoneLRO(ctx, fqn, metadata, &emptypb.Empty{})
		}
		return nil, err
	}

	metadata := &pb.AccessContextManagerOperationMetadata{}
	lroPrefix := "operations/" + fqn + "/delete/{{operationID}}"
	return s.operations.DoneLRO(ctx, lroPrefix, metadata, &emptypb.Empty{})
}

type accessLevelName struct {
	AccessPolicy string
	AccessLevel  string
}

func (n *accessLevelName) String() string {
	return "accessPolicies/" + n.AccessPolicy + "/accessLevels/" + n.AccessLevel
}

// parseAccessLevelName parses a string into a accessLevelName.
// The expected form is accessPolicies/{access_policy}/accessLevels/{access_level}
func (s *AccessContextManagerV1) parseAccessLevelName(name string) (*accessLevelName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 4 && tokens[0] == "accessPolicies" && tokens[2] == "accessLevels" {
		return &accessLevelName{AccessPolicy: tokens[1], AccessLevel: tokens[3]}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
