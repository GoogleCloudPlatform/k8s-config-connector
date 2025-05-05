// Copyright 2024 Google LLC
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

package mockcloudidentity

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/google/apps/cloudidentity/v1beta1"
	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

type groupsMembershipsServer struct {
	*MockService
	pb.UnimplementedGroupsMembershipsServerServer
}

func (s *groupsMembershipsServer) GetGroupsMembership(ctx context.Context, req *pb.GetGroupsMembershipRequest) (*pb.Membership, error) {
	//GET https://cloudidentity.googleapis.com/v1/{name=groups/*/memberships/*}
	name, err := s.parseMembershipName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Membership{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Error(4006): Membership does not exist.")
		}
		return nil, err
	}

	return obj, nil
}

func (s *groupsMembershipsServer) CreateGroupsMembership(ctx context.Context, req *pb.CreateGroupsMembershipRequest) (*longrunning.Operation, error) {
	// POST https://cloudidentity.googleapis.com/v1/{parent=groups/*}/memberships
	groupName, err := s.parseGroupName(*req.Parent)
	if err != nil {
		return nil, err
	}
	reqName := fmt.Sprintf("%s/memberships/%x", groupName.String(), time.Now().UnixMilli())
	name, err := s.parseMembershipName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := timestamppb.Now()

	obj := proto.Clone(req.GroupsMembership).(*pb.Membership)
	obj.Name = PtrTo(fmt.Sprintf("%s", name.String()))
	obj.CreateTime = now
	obj.UpdateTime = now
	obj.Type = PtrTo("USER")
	// Legacy field, not available in API v1 but in v1beta1
	obj.MemberKey = obj.PreferredMemberKey

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	retObj := proto.Clone(obj).(*pb.Membership)
	// output-only fields are not populated in LRO
	retObj.CreateTime = nil
	retObj.UpdateTime = nil
	retObj.Type = nil
	return buildLRO(retObj)
}

// No patch in mockgcp/generated/google/apps/cloudidentity/v1beta1/service_grpc.pb.go
// func (s *groupsMembershipsServer) PatchGroupsMembership(ctx context.Context, req *pb.PatchGroupsMembershipRequest) (*longrunning.Operation, error) {}

// TODO: implement these from mockgcp/generated/google/apps/cloudidentity/v1beta1/service_grpc.pb.go ?
// Modifies the `MembershipRole`s of a `Membership`.
// ModifyMembershipRolesGroupsMembership(ctx context.Context, in *ModifyMembershipRolesGroupsMembershipRequest, opts ...grpc.CallOption) (*ModifyMembershipRolesResponse, error)
// Searches direct groups of a member.
// SearchDirectGroupsGroupsMembership(ctx context.Context, in *SearchDirectGroupsGroupsMembershipRequest, opts ...grpc.CallOption) (*SearchDirectGroupsResponse, error)
// Search transitive groups of a member. **Note:** This feature is only available to Google Workspace Enterprise Standard, Enterprise Plus, and Enterprise for Education; and Cloud Identity Premium accounts. A transitive group is any group that has a direct or indirect membership to the member. Actor must have view permissions all transitive groups.
// SearchTransitiveGroupsGroupsMembership(ctx context.Context, in *SearchTransitiveGroupsGroupsMembershipRequest, opts ...grpc.CallOption) (*SearchTransitiveGroupsResponse, error)
// Search transitive memberships of a group. **Note:** This feature is only available to Google Workspace Enterprise Standard, Enterprise Plus, and Enterprise for Education; and Cloud Identity Premium accounts. A transitive membership is any direct or indirect membership of a group. Actor must have view permissions to all transitive memberships.
// SearchTransitiveMembershipsGroupsMembership(ctx context.Context, in *SearchTransitiveMembershipsGroupsMembershipRequest, opts ...grpc.CallOption) (*SearchTransitiveMembershipsResponse, error)

func (s *groupsMembershipsServer) ModifyMembershipRolesGroupsMembership(ctx context.Context, req *pb.ModifyMembershipRolesGroupsMembershipRequest) (*pb.ModifyMembershipRolesResponse, error) {
	name, err := s.parseMembershipName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Membership{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	gm := req.GetGroupsMembership()
	if gm == nil {
		return &pb.ModifyMembershipRolesResponse{
			Membership: &pb.Membership{Name: obj.Name, Roles: obj.Roles},
		}, nil
	}

	// Only contains the added/removed/updated roles in response
	var modifiedRoles []*pb.MembershipRole
	for i := range gm.RemoveRoles {
		for j, role := range obj.Roles {
			if gm.RemoveRoles[i] == *obj.Roles[j].Name {
				obj.Roles = append(obj.Roles[:j], obj.Roles[j+1:]...)
				modifiedRoles = append(modifiedRoles, role)
				break
			}
		}
	}

	for _, role := range gm.AddRoles {
		obj.Roles = append(obj.Roles, role)
		modifiedRoles = append(modifiedRoles, role)
	}

	// From proto defn:
	// The fully-qualified names of fields to update. May only contain the field `expiry_detail.expire_time`.
	// FieldMask *string `protobuf:"bytes,1,opt,name=field_mask,json=fieldMask" json:"field_mask,omitempty"`
	// The `MembershipRole`s to be updated. Only `MEMBER` `MembershipRoles` can currently be updated. May only contain a `MembershipRole` with `name` `MEMBER`.
	// MembershipRole *MembershipRole `protobuf:"bytes,2,opt,name=membership_role,json=membershipRole" json:"membership_role,omitempty"`
	for i := range gm.UpdateRolesParams {
		umr := gm.UpdateRolesParams[i].GetMembershipRole()
		if umr == nil {
			continue
		}
		if *umr.Name != "MEMBER" {
			continue
		}
		if umr.ExpiryDetail == nil {
			continue
		}
		if umr.ExpiryDetail.ExpireTime == nil {
			continue
		}
		if *gm.UpdateRolesParams[i].FieldMask != "expiry_detail.expire_time" {
			continue
		}

		for j, role := range obj.Roles {
			if *umr.Name == *obj.Roles[j].Name {
				obj.Roles[j].ExpiryDetail = proto.Clone(umr.ExpiryDetail).(*pb.ExpiryDetail)
				modifiedRoles = append(modifiedRoles, role)
				break
			}
		}
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}
	return &pb.ModifyMembershipRolesResponse{Membership: &pb.Membership{Name: obj.Name, Roles: modifiedRoles}}, nil
}

func (s *groupsMembershipsServer) DeleteGroupsMembership(ctx context.Context, req *pb.DeleteGroupsMembershipRequest) (*longrunning.Operation, error) {
	name, err := s.parseMembershipName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.Membership{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	// Returns a non-standard LRO
	lro := &longrunning.Operation{}
	lro.Done = true
	return lro, nil
}

type membershipName struct {
	Group      string
	Membership string
}

func (n *membershipName) String() string {
	return fmt.Sprintf("groups/%s/memberships/%s", n.Group, n.Membership)
}

func (s *MockService) parseMembershipName(name string) (*membershipName, error) {
	// From GET https://cloudidentity.googleapis.com/v1/{name=groups/*/memberships/*}
	// name: groups/{group}/memberships/{membership}
	tokens := strings.Split(name, "/")

	if len(tokens) == 4 && tokens[0] == "groups" && tokens[2] == "memberships" {
		name := &membershipName{
			Group:      tokens[1],
			Membership: tokens[3],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
