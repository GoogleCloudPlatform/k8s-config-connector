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

package mockgkehub

import (
	"context"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/gkehub/v1beta1"
)

type GKEHubMembership struct {
	*MockService
	pb.UnimplementedGkeHubMembershipServiceServer
}

func (s *GKEHubMembership) GetMembership(ctx context.Context, req *pb.GetMembershipRequest) (*pb.Membership, error) {
	name, err := s.parseMembershipName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Membership{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *GKEHubMembership) CreateMembership(ctx context.Context, req *pb.CreateMembershipRequest) (*longrunning.Operation, error) {
	reqName := req.Parent + "/memberships/" + req.MembershipId
	name, err := s.parseMembershipName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := timestamppb.Now()
	obj := proto.Clone(req.Resource).(*pb.Membership)
	obj.Name = fqn
	// The real gcp generated the values below
	obj.InfrastructureType = 2
	obj.ExternalId = "c772f869-1d6c-4d50-a92e-816c48322246"
	if obj.Authority != nil {
		obj.Authority.IdentityProvider = obj.Authority.Issuer
		obj.Authority.WorkloadIdentityPool = "${projectId}.svc.id.goog"
	}
	obj.UniqueId = "12345678"
	obj.CreateTime = now
	obj.State = &pb.MembershipState{Code: pb.MembershipState_READY}
	obj.UpdateTime = now

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	metadata := &pb.OperationMetadata{
		Target:     fqn,
		CreateTime: now,
		EndTime:    now,
	}
	return s.operations.StartLRO(ctx, name.String(), metadata, func() (proto.Message, error) {
		result := proto.Clone(obj).(*pb.Membership)
		result.CreateTime = now
		result.UpdateTime = now
		result.State = &pb.MembershipState{Code: pb.MembershipState_READY}
		return result, nil
	})
}

func (s *GKEHubMembership) UpdateMembership(ctx context.Context, req *pb.UpdateMembershipRequest) (*longrunning.Operation, error) {
	reqName := req.GetName()

	name, err := s.parseMembershipName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Membership{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	now := timestamppb.Now()
	// Required. A list of fields to be updated in this request.
	paths := req.GetUpdateMask().GetPaths()

	// TODO: Some sort of helper for fieldmask?
	for _, path := range paths {
		switch path {
		case "description":
			obj.Description = req.Resource.GetDescription()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
		}
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	metadata := &pb.OperationMetadata{
		Target:     fqn,
		CreateTime: now,
		EndTime:    now,
	}
	return s.operations.StartLRO(ctx, name.String(), metadata, func() (proto.Message, error) {
		result := proto.Clone(obj).(*pb.Membership)
		result.UpdateTime = now
		result.State = &pb.MembershipState{Code: pb.MembershipState_READY}
		return result, nil
	})
}

func (s *GKEHubMembership) DeleteMembership(ctx context.Context, req *pb.DeleteMembershipRequest) (*longrunning.Operation, error) {
	name, err := s.parseMembershipName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := timestamppb.Now()
	oldObj := &pb.Membership{}
	if err := s.storage.Delete(ctx, fqn, oldObj); err != nil {
		if status.Code(err) == codes.NotFound {
			return s.operations.NewLRO(ctx)
		}
		return &longrunningpb.Operation{}, err
	}

	metadata := &pb.OperationMetadata{
		Target:     fqn,
		CreateTime: now,
		EndTime:    now,
	}
	return s.operations.DoneLRO(ctx, name.String(), metadata, &pb.Membership{})
}
