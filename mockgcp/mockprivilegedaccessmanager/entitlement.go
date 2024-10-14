// Copyright 2022 Google LLC
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

package mockprivilegedaccessmanager

import (
	"context"
	"time"

	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/fields"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/privilegedaccessmanager/v1"
)

type PrivilegedAccessManager struct {
	*MockService
	pb.UnimplementedPrivilegedAccessManagerServer
}

func (s *PrivilegedAccessManager) GetEntitlement(ctx context.Context, req *pb.GetEntitlementRequest) (*pb.Entitlement, error) {
	name, err := s.parseEntitlementName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Entitlement{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *PrivilegedAccessManager) CreateEntitlement(ctx context.Context, req *pb.CreateEntitlementRequest) (*longrunning.Operation, error) {
	reqName := req.Parent + "/entitlements/" + req.EntitlementId
	name, err := s.parseEntitlementName(reqName)
	if err != nil {
		return nil, err
	}

	now := timestamppb.New(time.Now())
	fqn := name.String()

	obj := proto.Clone(req.Entitlement).(*pb.Entitlement)
	obj.Name = fqn
	obj.CreateTime = now
	obj.UpdateTime = now
	obj.Etag = fields.ComputeWeakEtag(obj)
	obj.State = pb.Entitlement_AVAILABLE
	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	metadata := constructOperationMetadata(fqn, "create")
	return s.operations.StartLRO(ctx, name.parent(), metadata, func() (proto.Message, error) {
		result := proto.Clone(obj).(*pb.Entitlement)
		metadata.EndTime = now
		return result, nil
	})
}

func (s *PrivilegedAccessManager) UpdateEntitlement(ctx context.Context, req *pb.UpdateEntitlementRequest) (*longrunning.Operation, error) {
	reqName := req.GetEntitlement().GetName()

	name, err := s.parseEntitlementName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Entitlement{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// Required. A list of fields to be updated in this request.
	paths := req.GetUpdateMask().GetPaths()

	for _, path := range paths {
		switch path {
		case "eligibleUsers":
			obj.EligibleUsers = req.GetEntitlement().GetEligibleUsers()
		case "approvalWorkflow":
			obj.ApprovalWorkflow = req.GetEntitlement().GetApprovalWorkflow()
		case "privilegedAccess":
			obj.PrivilegedAccess = req.GetEntitlement().GetPrivilegedAccess()
		case "maxRequestDuration":
			obj.MaxRequestDuration = req.GetEntitlement().GetMaxRequestDuration()
		case "requesterJustificationConfig":
			obj.RequesterJustificationConfig = req.GetEntitlement().GetRequesterJustificationConfig()
		case "additionalNotificationTargets":
			obj.AdditionalNotificationTargets = req.GetEntitlement().GetAdditionalNotificationTargets()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
		}
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	metadata := constructOperationMetadata(fqn, "update")
	return s.operations.StartLRO(ctx, name.parent(), metadata, func() (proto.Message, error) {
		result := proto.Clone(obj).(*pb.Entitlement)
		now := timestamppb.New(time.Now())
		metadata.EndTime = now
		return result, nil
	})
}

func (s *PrivilegedAccessManager) DeleteEntitlement(ctx context.Context, req *pb.DeleteEntitlementRequest) (*longrunning.Operation, error) {
	name, err := s.parseEntitlementName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	oldObj := &pb.Entitlement{}
	if err := s.storage.Delete(ctx, fqn, oldObj); err != nil {
		return nil, err
	}
	metadata := constructOperationMetadata(fqn, "delete")
	return s.operations.StartLRO(ctx, name.parent(), metadata, func() (proto.Message, error) {
		result := proto.Clone(oldObj).(*pb.Entitlement)
		result.State = pb.Entitlement_DELETED
		now := timestamppb.New(time.Now())
		metadata.EndTime = now
		return result, nil
	})
}

func constructOperationMetadata(target, verb string) *pb.OperationMetadata {
	now := timestamppb.New(time.Now())
	return &pb.OperationMetadata{
		Target:                target,
		CreateTime:            now,
		ApiVersion:            "v1",
		RequestedCancellation: false,
		Verb:                  verb,
	}
}
