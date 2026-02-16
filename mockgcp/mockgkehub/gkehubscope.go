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

package mockgkehub

import (
	"context"

	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/gkehub/v1beta"
)

type GKEHubScopes struct {
	*MockService
	pb.UnimplementedGkeHubScopesServer
}

func (s *GKEHubScopes) GetScope(ctx context.Context, req *pb.GetScopeRequest) (*pb.Scope, error) {
	name, err := s.parseScopeName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Scope{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *GKEHubScopes) CreateScope(ctx context.Context, req *pb.CreateScopeRequest) (*longrunning.Operation, error) {
	reqName := req.Parent + "/scopes/" + req.ScopeId
	name, err := s.parseScopeName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := timestamppb.Now()
	obj := proto.Clone(req.Resource).(*pb.Scope)
	obj.Name = fqn
	obj.Uid = "c772f869-1d6c-4d50-a92e-816c48322246"
	obj.CreateTime = now
	obj.UpdateTime = now
	obj.State = &pb.ScopeState{Code: pb.ScopeState_READY}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	metadata := &pb.OperationMetadata{
		Target:     fqn,
		CreateTime: now,
		EndTime:    now,
	}
	return s.operations.StartLRO(ctx, fqn, metadata, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *GKEHubScopes) PatchScope(ctx context.Context, req *pb.UpdateScopeRequest) (*longrunning.Operation, error) {
	name, err := s.parseScopeName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Scope{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	now := timestamppb.Now()
	paths := req.GetUpdateMask().GetPaths()

	for _, path := range paths {
		switch path {
		case "labels":
			obj.Labels = req.Resource.GetLabels()
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
	return s.operations.StartLRO(ctx, fqn, metadata, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *GKEHubScopes) DeleteScope(ctx context.Context, req *pb.DeleteScopeRequest) (*longrunning.Operation, error) {
	name, err := s.parseScopeName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := timestamppb.Now()
	oldObj := &pb.Scope{}
	if err := s.storage.Delete(ctx, fqn, oldObj); err != nil {
		return nil, err
	}

	metadata := &pb.OperationMetadata{
		Target:     fqn,
		CreateTime: now,
		EndTime:    now,
	}
	return s.operations.StartLRO(ctx, fqn, metadata, func() (proto.Message, error) {
		return &pb.Scope{}, nil
	})
}

func (s *GKEHubScopes) GetMembershipBinding(ctx context.Context, req *pb.GetMembershipBindingRequest) (*pb.MembershipBinding, error) {
	name, err := s.parseMembershipBindingName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.MembershipBinding{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *GKEHubScopes) CreateMembershipBinding(ctx context.Context, req *pb.CreateMembershipBindingRequest) (*longrunning.Operation, error) {
	reqName := req.Parent + "/bindings/" + req.MembershipBindingId
	name, err := s.parseMembershipBindingName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := timestamppb.Now()
	obj := proto.Clone(req.Resource).(*pb.MembershipBinding)
	obj.Name = fqn
	obj.Uid = "c772f869-1d6c-4d50-a92e-816c48322247"
	obj.CreateTime = now
	obj.UpdateTime = now
	obj.State = &pb.MembershipBindingState{Code: pb.MembershipBindingState_READY}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	metadata := &pb.OperationMetadata{
		Target:     fqn,
		CreateTime: now,
		EndTime:    now,
	}
	return s.operations.StartLRO(ctx, fqn, metadata, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *GKEHubScopes) PatchMembershipBinding(ctx context.Context, req *pb.UpdateMembershipBindingRequest) (*longrunning.Operation, error) {
	name, err := s.parseMembershipBindingName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.MembershipBinding{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	now := timestamppb.Now()
	paths := req.GetUpdateMask().GetPaths()

	for _, path := range paths {
		switch path {
		case "labels":
			obj.Labels = req.Resource.GetLabels()
		case "scope":
			obj.Scope = req.Resource.GetScope()
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
	return s.operations.StartLRO(ctx, fqn, metadata, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *GKEHubScopes) DeleteMembershipBinding(ctx context.Context, req *pb.DeleteMembershipBindingRequest) (*longrunning.Operation, error) {
	name, err := s.parseMembershipBindingName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := timestamppb.Now()
	oldObj := &pb.MembershipBinding{}
	if err := s.storage.Delete(ctx, fqn, oldObj); err != nil {
		return nil, err
	}

	metadata := &pb.OperationMetadata{
		Target:     fqn,
		CreateTime: now,
		EndTime:    now,
	}
	return s.operations.StartLRO(ctx, fqn, metadata, func() (proto.Message, error) {
		return &pb.MembershipBinding{}, nil
	})
}

func (s *GKEHubScopes) GetScopeRBACRoleBinding(ctx context.Context, req *pb.GetScopeRBACRoleBindingRequest) (*pb.RBACRoleBinding, error) {
	name, err := s.parseScopeRBACRoleBindingName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.RBACRoleBinding{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *GKEHubScopes) CreateScopeRBACRoleBinding(ctx context.Context, req *pb.CreateScopeRBACRoleBindingRequest) (*longrunning.Operation, error) {
	reqName := req.Parent + "/rbacrolebindings/" + req.RbacrolebindingId
	name, err := s.parseScopeRBACRoleBindingName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := timestamppb.Now()
	obj := proto.Clone(req.Resource).(*pb.RBACRoleBinding)
	obj.Name = fqn
	obj.Uid = "c772f869-1d6c-4d50-a92e-816c48322248"
	obj.CreateTime = now
	obj.UpdateTime = now
	obj.State = &pb.RBACRoleBindingState{Code: pb.RBACRoleBindingState_READY}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	metadata := &pb.OperationMetadata{
		Target:     fqn,
		CreateTime: now,
		EndTime:    now,
	}
	return s.operations.StartLRO(ctx, fqn, metadata, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *GKEHubScopes) PatchScopeRBACRoleBinding(ctx context.Context, req *pb.UpdateScopeRBACRoleBindingRequest) (*longrunning.Operation, error) {
	name, err := s.parseScopeRBACRoleBindingName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.RBACRoleBinding{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	now := timestamppb.Now()
	paths := req.GetUpdateMask().GetPaths()

	for _, path := range paths {
		switch path {
		case "labels":
			obj.Labels = req.Resource.GetLabels()
		case "user":
			obj.User = req.Resource.GetUser()
		case "group":
			obj.Group = req.Resource.GetGroup()
		case "role":
			obj.Role = req.Resource.GetRole()
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
	return s.operations.StartLRO(ctx, fqn, metadata, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *GKEHubScopes) DeleteScopeRBACRoleBinding(ctx context.Context, req *pb.DeleteScopeRBACRoleBindingRequest) (*longrunning.Operation, error) {
	name, err := s.parseScopeRBACRoleBindingName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := timestamppb.Now()
	oldObj := &pb.RBACRoleBinding{}
	if err := s.storage.Delete(ctx, fqn, oldObj); err != nil {
		return nil, err
	}

	metadata := &pb.OperationMetadata{
		Target:     fqn,
		CreateTime: now,
		EndTime:    now,
	}
	return s.operations.StartLRO(ctx, fqn, metadata, func() (proto.Message, error) {
		return &pb.RBACRoleBinding{}, nil
	})
}
