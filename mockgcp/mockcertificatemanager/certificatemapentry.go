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

package mockcertificatemanager

import (
	"context"
	"fmt"

	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"k8s.io/klog/v2"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/certificatemanager/v1"
)

func (s *CertificateManagerV1) GetCertificateMapEntry(ctx context.Context, req *pb.GetCertificateMapEntryRequest) (*pb.CertificateMapEntry, error) {
	name, err := s.parseCertificateMapEntryName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.CertificateMapEntry{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *CertificateManagerV1) CreateCertificateMapEntry(ctx context.Context, req *pb.CreateCertificateMapEntryRequest) (*longrunning.Operation, error) {
	reqName := req.Parent + "/certificateMapEntries/" + req.CertificateMapEntryId
	name, err := s.parseCertificateMapEntryName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.CertificateMapEntry).(*pb.CertificateMapEntry)
	obj.Name = fqn
	now := timestamppb.Now()
	obj.CreateTime = now

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}
	lroMetadata := &pb.OperationMetadata{
		ApiVersion:            "v1",
		CreateTime:            now,
		RequestedCancellation: false,
		Target:                fqn,
		Verb:                  "create",
	}
	lroPrefix := fmt.Sprintf("projects/%s/locations/global", name.Project.ID)
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		result := proto.Clone(obj).(*pb.CertificateMapEntry)
		result.Labels = nil
		lroMetadata.RequestedCancellation = false
		return result, nil
	})
}

func (s *CertificateManagerV1) UpdateCertificateMapEntry(ctx context.Context, req *pb.UpdateCertificateMapEntryRequest) (*longrunning.Operation, error) {
	reqName := req.GetCertificateMapEntry().GetName()

	name, err := s.parseCertificateMapEntryName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.CertificateMapEntry{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// Required. The update mask applies to the resource.
	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		klog.Warningf("update_mask was not provided in request, should be required")
	}
	// TODO: Some sort of helper for fieldmask?
	for _, path := range paths {
		switch path {
		case "description":
			obj.Description = req.GetCertificateMapEntry().GetDescription()
		case "labels":
			obj.Labels = req.GetCertificateMapEntry().GetLabels()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
		}
	}
	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}
	now := timestamppb.Now()
	lroMetadata := &pb.OperationMetadata{
		ApiVersion:            "v1",
		CreateTime:            now,
		RequestedCancellation: false,
		Target:                fqn,
		Verb:                  "update",
	}
	lroPrefix := fmt.Sprintf("projects/%s/locations/global", name.Project.ID)

	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		result := proto.Clone(obj).(*pb.CertificateMapEntry)
		result.Labels = nil
		return result, nil
	})
}

func (s *CertificateManagerV1) DeleteCertificateMapEntry(ctx context.Context, req *pb.DeleteCertificateMapEntryRequest) (*longrunning.Operation, error) {
	name, err := s.parseCertificateMapEntryName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deletedObj := &pb.CertificateMapEntry{}
	if err := s.storage.Delete(ctx, fqn, deletedObj); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/global", name.Project.ID)
	lroMetadata := &pb.OperationMetadata{
		ApiVersion:            "v1",
		CreateTime:            timestamppb.Now(),
		RequestedCancellation: false,
		Target:                fqn,
		Verb:                  "delete",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		return &emptypb.Empty{}, nil
	})
}
