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

// +tool:mockgcp-support
// proto.service: google.identity.accesscontextmanager.v1.AccessContextManager
// proto.message: google.identity.accesscontextmanager.v1.GcpUserAccessBinding

package mockaccesscontextmanager

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "cloud.google.com/go/accesscontextmanager/apiv1/accesscontextmanagerpb"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

func (s *AccessContextManagerV1) GetGcpUserAccessBinding(ctx context.Context, req *pb.GetGcpUserAccessBindingRequest) (*pb.GcpUserAccessBinding, error) {
	name, err := s.parseGcpUserAccessBindingName(req.GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.GcpUserAccessBinding{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "GcpUserAccessBinding %q not found.", fqn)
		}
		return nil, err
	}

	return obj, nil
}

type gcpUserAccessBindingName struct {
	Organization string
	BindingID    string
}

func (n *gcpUserAccessBindingName) String() string {
	return "organizations/" + n.Organization + "/gcpUserAccessBindings/" + n.BindingID
}

func (s *AccessContextManagerV1) parseGcpUserAccessBindingName(name string) (*gcpUserAccessBindingName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 4 && tokens[0] == "organizations" && tokens[2] == "gcpUserAccessBindings" {
		return &gcpUserAccessBindingName{Organization: tokens[1], BindingID: tokens[3]}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

func (s *AccessContextManagerV1) ListGcpUserAccessBindings(ctx context.Context, req *pb.ListGcpUserAccessBindingsRequest) (*pb.ListGcpUserAccessBindingsResponse, error) {
	parentOrg, err := s.parseOrganizationParent(req.GetParent())
	if err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("organizations/%s/gcpUserAccessBindings/", parentOrg)

	response := &pb.ListGcpUserAccessBindingsResponse{}
	kind := (&pb.GcpUserAccessBinding{}).ProtoReflect().Descriptor()

	if err := s.storage.List(ctx, kind, storage.ListOptions{Prefix: prefix}, func(p proto.Message) error {
		binding := p.(*pb.GcpUserAccessBinding)
		response.GcpUserAccessBindings = append(response.GcpUserAccessBindings, binding)
		return nil
	}); err != nil {
		return nil, status.Errorf(codes.Internal, "Error listing GcpUserAccessBindings: %v", err)
	}

	return response, nil
}

func (s *AccessContextManagerV1) parseOrganizationParent(parent string) (string, error) {
	tokens := strings.Split(parent, "/")
	if len(tokens) == 2 && tokens[0] == "organizations" {
		return tokens[1], nil
	}
	return "", status.Errorf(codes.InvalidArgument, "parent %q is not valid", parent)
}

func (s *AccessContextManagerV1) CreateGcpUserAccessBinding(ctx context.Context, req *pb.CreateGcpUserAccessBindingRequest) (*longrunningpb.Operation, error) {
	if req.GetGcpUserAccessBinding() == nil {
		return nil, status.Errorf(codes.InvalidArgument, "gcp_user_access_binding is required")
	}

	parentOrg, err := s.parseOrganizationParent(req.GetParent())
	if err != nil {
		return nil, err
	}

	groupKey := req.GetGcpUserAccessBinding().GetGroupKey()
	if groupKey == "" {
		return nil, status.Errorf(codes.InvalidArgument, "group_key is required")
	}

	prefix := fmt.Sprintf("organizations/%s/gcpUserAccessBindings/", parentOrg)
	kind := (&pb.GcpUserAccessBinding{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, kind, storage.ListOptions{Prefix: prefix}, func(p proto.Message) error {
		existing := p.(*pb.GcpUserAccessBinding)
		if existing.GetGroupKey() == groupKey {
			return status.Errorf(codes.AlreadyExists, "GcpUserAccessBinding with group_key %q already exists in organization %q", groupKey, parentOrg)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	bindingID := strconv.FormatInt(time.Now().UnixNano(), 10)
	if req.GetGcpUserAccessBinding().GetName() != "" {
		parsed, err := s.parseGcpUserAccessBindingName(req.GetGcpUserAccessBinding().GetName())
		if err == nil && parsed.BindingID != "" {
			bindingID = parsed.BindingID
		}
	}

	name := &gcpUserAccessBindingName{
		Organization: parentOrg,
		BindingID:    bindingID,
	}
	fqn := name.String()

	obj := proto.CloneOf(req.GetGcpUserAccessBinding())
	obj.Name = fqn

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	metadata := &pb.GcpUserAccessBindingOperationMetadata{}
	lroPrefix := "operations/" + fqn + "/create/{{operationID}}"
	return s.operations.DoneLRO(ctx, lroPrefix, metadata, obj)
}

func (s *AccessContextManagerV1) UpdateGcpUserAccessBinding(ctx context.Context, req *pb.UpdateGcpUserAccessBindingRequest) (*longrunningpb.Operation, error) {
	if req.GetGcpUserAccessBinding() == nil {
		return nil, status.Errorf(codes.InvalidArgument, "gcp_user_access_binding is required for update")
	}
	name, err := s.parseGcpUserAccessBindingName(req.GetGcpUserAccessBinding().GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	existing := &pb.GcpUserAccessBinding{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "GcpUserAccessBinding %q not found.", fqn)
		}
		return nil, err
	}

	updated := proto.CloneOf(existing)
	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask is required")
	}

	for _, path := range paths {
		switch path {
		case "access_levels":
			updated.AccessLevels = req.GetGcpUserAccessBinding().GetAccessLevels()
		default:
			return nil, fmt.Errorf("UpdateGcpUserAccessBinding: unsupported update_mask path (in mock) %q", path)
		}
	}

	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}

	metadata := &pb.GcpUserAccessBindingOperationMetadata{}
	lroPrefix := "operations/" + fqn + "/update/{{operationID}}"
	return s.operations.DoneLRO(ctx, lroPrefix, metadata, updated)
}

func (s *AccessContextManagerV1) DeleteGcpUserAccessBinding(ctx context.Context, req *pb.DeleteGcpUserAccessBindingRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseGcpUserAccessBindingName(req.GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	metadata := &pb.GcpUserAccessBindingOperationMetadata{}
	lroPrefix := "operations/" + fqn + "/delete/{{operationID}}"

	if err := s.storage.Delete(ctx, fqn, &pb.GcpUserAccessBinding{}); err != nil {
		if status.Code(err) == codes.NotFound {
			return s.operations.DoneLRO(ctx, lroPrefix, metadata, &emptypb.Empty{})
		}
		return nil, err
	}

	return s.operations.DoneLRO(ctx, lroPrefix, metadata, &emptypb.Empty{})
}
