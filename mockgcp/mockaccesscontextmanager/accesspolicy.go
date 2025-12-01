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
// proto.message: google.identity.accesscontextmanager.v1.AccessPolicy

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
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/accesscontextmanager/apiv1/accesscontextmanagerpb"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

func (s *AccessContextManagerV1) GetAccessPolicy(ctx context.Context, req *pb.GetAccessPolicyRequest) (*pb.AccessPolicy, error) {
	name, err := s.parseAccessPolicyName(req.GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.AccessPolicy{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "AccessPolicy %q not found.", fqn)
		}
		return nil, status.Errorf(codes.Internal, "Error getting AccessPolicy: %v", err)
	}

	return obj, nil
}

func (s *AccessContextManagerV1) CreateAccessPolicy(ctx context.Context, req *pb.AccessPolicy) (*longrunningpb.Operation, error) {
	// Check if a policy already exists for this parent.
	var policyExists bool
	policyKind := (&pb.AccessPolicy{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, policyKind, storage.ListOptions{}, func(p proto.Message) error {
		policy := p.(*pb.AccessPolicy)
		if policy.Parent == req.Parent {
			policyExists = true
		}
		return nil
	}); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to check for existing policies: %v", err)
	}

	if policyExists {
		return nil, status.Errorf(codes.Aborted, "organization %q can only have one access policy", req.Parent)
	}

	// Name is output-only and should be ignored, but the resource ID is server-assigned.
	policyID := strconv.FormatInt(time.Now().UnixNano(), 10)
	name := &accessPolicyName{AccessPolicy: policyID}
	fqn := name.String()

	obj := ProtoClone(req)
	obj.Name = fqn
	obj.Etag = computeEtag(obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, status.Errorf(codes.Internal, "Error creating AccessPolicy: %v", err)
	}

	metadata := &pb.AccessContextManagerOperationMetadata{}
	lroPrefix := "operations/" + fqn + "/create/{{operationID}}"
	return s.operations.DoneLRO(ctx, lroPrefix, metadata, obj)
}

func (s *AccessContextManagerV1) UpdateAccessPolicy(ctx context.Context, req *pb.UpdateAccessPolicyRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseAccessPolicyName(req.GetPolicy().GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	existing := &pb.AccessPolicy{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "AccessPolicy %q not found.", fqn)
		}
		return nil, status.Errorf(codes.Internal, "Error getting AccessPolicy for update: %v", err)
	}

	// An etag is not required from the client, but if it is provided, it must match the current etag of the resource.
	if etag := req.GetPolicy().GetEtag(); etag != "" && etag != existing.GetEtag() {
		return nil, status.Errorf(codes.Aborted, "etag mismatch")
	}

	updated := ProtoClone(existing)
	paths := req.GetUpdateMask().GetPaths()
	for _, path := range paths {
		switch path {
		case "title":
			updated.Title = req.GetPolicy().GetTitle()
		case "scopes":
			updated.Scopes = req.GetPolicy().GetScopes()
		default:
			return nil, fmt.Errorf("UpdateAccessLevel: unsupported update_mask path (in mock) %q", path)
		}
	}

	updated.UpdateTime = timestamppb.Now()
	updated.Etag = computeEtag(updated)

	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, status.Errorf(codes.Internal, "Error updating AccessPolicy: %v", err)
	}

	metadata := &pb.AccessContextManagerOperationMetadata{}
	lroPrefix := "operations/" + fqn + "/update/{{operationID}}"
	lro, err := s.operations.DoneLRO(ctx, lroPrefix, metadata, updated)
	// Technically the operation is empty when it is later retrieved (response and metadata are not set), but that's kinda hard for us to do here.
	return lro, err
}

func (s *AccessContextManagerV1) DeleteAccessPolicy(ctx context.Context, req *pb.DeleteAccessPolicyRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseAccessPolicyName(req.GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	if err := s.storage.Delete(ctx, fqn, &pb.AccessPolicy{}); err != nil {
		if status.Code(err) == codes.NotFound {
			// Deleting a non-existent policy should succeed idempotently
			metadata := &pb.AccessContextManagerOperationMetadata{}
			return s.operations.DoneLRO(ctx, fqn, metadata, &emptypb.Empty{})
		}
		return nil, status.Errorf(codes.Internal, "Error deleting AccessPolicy: %v", err)
	}

	metadata := &pb.AccessContextManagerOperationMetadata{}
	lroPrefix := "operations/" + fqn + "/delete/{{operationID}}"
	return s.operations.DoneLRO(ctx, lroPrefix, metadata, &emptypb.Empty{})
}

func (s *AccessContextManagerV1) ListAccessPolicies(ctx context.Context, req *pb.ListAccessPoliciesRequest) (*pb.ListAccessPoliciesResponse, error) {
	if req.GetParent() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "parent is required")
	}

	response := &pb.ListAccessPoliciesResponse{}
	policyKind := (&pb.AccessPolicy{}).ProtoReflect().Descriptor()

	if err := s.storage.List(ctx, policyKind, storage.ListOptions{}, func(p proto.Message) error {
		policy := p.(*pb.AccessPolicy)
		if policy.Parent == req.GetParent() {
			response.AccessPolicies = append(response.AccessPolicies, policy)
		}
		return nil
	}); err != nil {
		return nil, status.Errorf(codes.Internal, "Error listing AccessPolicies: %v", err)
	}

	return response, nil
}

type accessPolicyName struct {
	AccessPolicy string
}

func (n *accessPolicyName) String() string {
	return "accessPolicies/" + n.AccessPolicy
}

func (s *MockService) parseAccessPolicyName(name string) (*accessPolicyName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 2 && tokens[0] == "accessPolicies" {
		return &accessPolicyName{AccessPolicy: tokens[1]}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
