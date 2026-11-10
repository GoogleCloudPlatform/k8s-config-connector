// Copyright 2026 High Five Labs
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

// The
// OrganizationSecurityProfileGroupService's create, get, update and delete of
// security profiles and security profile groups, in the shape of the project
// service's (securityprofile.go, securityprofilegroup.go), whose names must be
// a project's.

package mocknetworksecurity

import (
	"context"
	"strings"
	"time"

	pbv1 "cloud.google.com/go/networksecurity/apiv1/networksecuritypb"
	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// OrganizationSecurityProfileGroupServer serves an organization's security
// profiles and security profile groups
// (organizations/{organization}/locations/{location}/securityProfiles/{id},
// .../securityProfileGroups/{id}).
type OrganizationSecurityProfileGroupServer struct {
	*MockService
	pbv1.UnimplementedOrganizationSecurityProfileGroupServiceServer
}

// parseOrganizationParent returns the location prefix
// organizations/{organization}/locations/{location}.
func parseOrganizationParent(parent string) (string, error) {
	tokens := strings.Split(parent, "/")
	if len(tokens) == 4 && tokens[0] == "organizations" && tokens[1] != "" && tokens[2] == "locations" && tokens[3] != "" {
		return parent, nil
	}
	return "", status.Errorf(codes.InvalidArgument, "parent %q is not valid", parent)
}

// parseOrganizationName checks name is
// organizations/{organization}/locations/{location}/{collection}/{id} and
// returns its location prefix.
func parseOrganizationName(name, collection string) (string, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[4] == collection && tokens[5] != "" {
		if prefix, err := parseOrganizationParent(strings.Join(tokens[:4], "/")); err == nil {
			return prefix, nil
		}
	}
	return "", status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

func (s *OrganizationSecurityProfileGroupServer) lro(ctx context.Context, prefix, target, verb string, result func() protoreflect.ProtoMessage) (*longrunning.Operation, error) {
	lroMetadata := &pbv1.OperationMetadata{
		CreateTime: timestamppb.New(time.Now()),
		Target:     target,
		Verb:       verb,
		ApiVersion: "v1",
	}
	return s.operations.StartLRO(ctx, prefix, lroMetadata, func() (protoreflect.ProtoMessage, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())
		return result(), nil
	})
}

func (s *OrganizationSecurityProfileGroupServer) CreateSecurityProfile(ctx context.Context, req *pbv1.CreateSecurityProfileRequest) (*longrunning.Operation, error) {
	prefix, err := parseOrganizationParent(req.Parent)
	if err != nil {
		return nil, err
	}
	fqn := req.Parent + "/securityProfiles/" + req.SecurityProfileId
	obj := proto.CloneOf(req.SecurityProfile)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(time.Now())
	obj.UpdateTime = obj.CreateTime
	obj.Etag = "dummy-etag"
	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}
	return s.lro(ctx, prefix, fqn, "create", func() protoreflect.ProtoMessage { return proto.CloneOf(obj) })
}

func (s *OrganizationSecurityProfileGroupServer) GetSecurityProfile(ctx context.Context, req *pbv1.GetSecurityProfileRequest) (*pbv1.SecurityProfile, error) {
	if _, err := parseOrganizationName(req.Name, "securityProfiles"); err != nil {
		return nil, err
	}
	obj := &pbv1.SecurityProfile{}
	if err := s.storage.Get(ctx, req.Name, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", req.Name)
		}
		return nil, err
	}
	return obj, nil
}

func (s *OrganizationSecurityProfileGroupServer) UpdateSecurityProfile(ctx context.Context, req *pbv1.UpdateSecurityProfileRequest) (*longrunning.Operation, error) {
	fqn := req.GetSecurityProfile().GetName()
	prefix, err := parseOrganizationName(fqn, "securityProfiles")
	if err != nil {
		return nil, err
	}
	existing := &pbv1.SecurityProfile{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		return nil, err
	}
	obj := proto.CloneOf(req.SecurityProfile)
	obj.CreateTime = existing.CreateTime
	obj.UpdateTime = timestamppb.New(time.Now())
	obj.Etag = "updated-etag"
	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}
	return s.lro(ctx, prefix, fqn, "update", func() protoreflect.ProtoMessage { return proto.CloneOf(obj) })
}

func (s *OrganizationSecurityProfileGroupServer) DeleteSecurityProfile(ctx context.Context, req *pbv1.DeleteSecurityProfileRequest) (*longrunning.Operation, error) {
	prefix, err := parseOrganizationName(req.Name, "securityProfiles")
	if err != nil {
		return nil, err
	}
	if err := s.storage.Delete(ctx, req.Name, &pbv1.SecurityProfile{}); err != nil {
		return nil, err
	}
	return s.lro(ctx, prefix, req.Name, "delete", func() protoreflect.ProtoMessage { return &emptypb.Empty{} })
}

func (s *OrganizationSecurityProfileGroupServer) CreateSecurityProfileGroup(ctx context.Context, req *pbv1.CreateSecurityProfileGroupRequest) (*longrunning.Operation, error) {
	prefix, err := parseOrganizationParent(req.Parent)
	if err != nil {
		return nil, err
	}
	fqn := req.Parent + "/securityProfileGroups/" + req.SecurityProfileGroupId
	obj := proto.CloneOf(req.SecurityProfileGroup)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(time.Now())
	obj.UpdateTime = obj.CreateTime
	obj.Etag = "dummy-etag"
	obj.DataPathId = 12345
	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}
	return s.lro(ctx, prefix, fqn, "create", func() protoreflect.ProtoMessage { return proto.CloneOf(obj) })
}

func (s *OrganizationSecurityProfileGroupServer) GetSecurityProfileGroup(ctx context.Context, req *pbv1.GetSecurityProfileGroupRequest) (*pbv1.SecurityProfileGroup, error) {
	if _, err := parseOrganizationName(req.Name, "securityProfileGroups"); err != nil {
		return nil, err
	}
	obj := &pbv1.SecurityProfileGroup{}
	if err := s.storage.Get(ctx, req.Name, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", req.Name)
		}
		return nil, err
	}
	return obj, nil
}

func (s *OrganizationSecurityProfileGroupServer) UpdateSecurityProfileGroup(ctx context.Context, req *pbv1.UpdateSecurityProfileGroupRequest) (*longrunning.Operation, error) {
	fqn := req.GetSecurityProfileGroup().GetName()
	prefix, err := parseOrganizationName(fqn, "securityProfileGroups")
	if err != nil {
		return nil, err
	}
	existing := &pbv1.SecurityProfileGroup{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		return nil, err
	}
	obj := proto.CloneOf(req.SecurityProfileGroup)
	obj.CreateTime = existing.CreateTime
	obj.UpdateTime = timestamppb.New(time.Now())
	obj.Etag = "updated-etag"
	obj.DataPathId = existing.DataPathId
	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}
	return s.lro(ctx, prefix, fqn, "update", func() protoreflect.ProtoMessage { return proto.CloneOf(obj) })
}

func (s *OrganizationSecurityProfileGroupServer) DeleteSecurityProfileGroup(ctx context.Context, req *pbv1.DeleteSecurityProfileGroupRequest) (*longrunning.Operation, error) {
	prefix, err := parseOrganizationName(req.Name, "securityProfileGroups")
	if err != nil {
		return nil, err
	}
	if err := s.storage.Delete(ctx, req.Name, &pbv1.SecurityProfileGroup{}); err != nil {
		return nil, err
	}
	return s.lro(ctx, prefix, req.Name, "delete", func() protoreflect.ProtoMessage { return &emptypb.Empty{} })
}
