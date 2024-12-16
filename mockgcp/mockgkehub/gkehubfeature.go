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
	"fmt"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/gkehub/v1beta"
)

type GKEHubFeature struct {
	*MockService
	pb.UnimplementedGkeHubServer
}

func (s *GKEHubFeature) GetFeature(ctx context.Context, req *pb.GetFeatureRequest) (*pb.Feature, error) {
	name, err := s.parseFeatureName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Feature{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *GKEHubFeature) CreateFeature(ctx context.Context, req *pb.CreateFeatureRequest) (*longrunning.Operation, error) {
	reqName := req.Parent + "/features/" + req.FeatureId
	name, err := s.parseFeatureName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := timestamppb.Now()

	obj := proto.Clone(req.Resource).(*pb.Feature)
	obj.Name = fqn

	// Mimic the GCP API validation logic.
	for id, spec := range obj.MembershipSpecs {
		acmSpec := spec.GetConfigmanagement()
		if acmSpec != nil {
			if acmSpec.GetConfigSync() == nil && acmSpec.GetHierarchyController() == nil && acmSpec.GetPolicyController() == nil {
				return nil, fmt.Errorf("none of configsync or hierarchycontroller or policycontroller is specified under configmanagement for memebership %s", id)
			}
		}
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}
	metadata := &pb.OperationMetadata{
		Target:     fqn,
		CreateTime: now,
		EndTime:    now,
	}
	return s.operations.StartLRO(ctx, name.String(), metadata, func() (proto.Message, error) {
		result := proto.Clone(obj).(*pb.Feature)
		result.CreateTime = now
		result.UpdateTime = now
		result.ResourceState = &pb.FeatureResourceState{State: pb.FeatureResourceState_ACTIVE}
		return result, nil
	})
}

func (s *GKEHubFeature) UpdateFeature(ctx context.Context, req *pb.UpdateFeatureRequest) (*longrunning.Operation, error) {
	reqName := req.GetName()

	name, err := s.parseFeatureName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Feature{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}
	now := timestamppb.Now()
	obj.UpdateTime = now
	// Required. A list of fields to be updated in this request.
	paths := req.GetUpdateMask().GetPaths()

	// TODO: Some sort of helper for fieldmask?
	for _, path := range paths {
		switch path {
		case "labels":
			obj.Labels = req.Resource.GetLabels()
		// Spec is in the GCP API, not a KRM Spec
		case "spec":
			obj.Spec = req.GetResource().Spec
		case "membershipSpecs":
			obj.MembershipSpecs = updateMembershipSpecsMap(obj.MembershipSpecs, req.GetResource().GetMembershipSpecs())
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
		result := proto.Clone(obj).(*pb.Feature)
		result.UpdateTime = now
		result.ResourceState = &pb.FeatureResourceState{State: pb.FeatureResourceState_ACTIVE}
		return result, nil
	})
}

func updateMembershipSpecsMap(membershipSpecs, membershipSpecsPatch map[string]*pb.MembershipFeatureSpec) map[string]*pb.MembershipFeatureSpec {
	if membershipSpecs == nil {
		membershipSpecs = make(map[string]*pb.MembershipFeatureSpec)
	}
	for k, v := range membershipSpecsPatch {
		membershipSpecs[k] = v
	}
	return membershipSpecs
}

func (s *GKEHubFeature) DeleteFeature(ctx context.Context, req *pb.DeleteFeatureRequest) (*longrunning.Operation, error) {
	name, err := s.parseFeatureName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := timestamppb.Now()

	oldObj := &pb.Feature{}
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
	return s.operations.DoneLRO(ctx, name.String(), metadata, &pb.Feature{})
}
