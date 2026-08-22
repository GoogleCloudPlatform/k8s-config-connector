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

package mockcertificatemanager

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"k8s.io/klog/v2"

	pb "cloud.google.com/go/certificatemanager/apiv1/certificatemanagerpb"
)

func (s *CertificateManagerV1) GetTrustConfig(ctx context.Context, req *pb.GetTrustConfigRequest) (*pb.TrustConfig, error) {
	name, err := s.parseTrustConfigName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.TrustConfig{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *CertificateManagerV1) CreateTrustConfig(ctx context.Context, req *pb.CreateTrustConfigRequest) (*longrunning.Operation, error) {
	reqName := req.Parent + "/trustConfigs/" + req.TrustConfigId
	name, err := s.parseTrustConfigName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.TrustConfig).(*pb.TrustConfig)
	obj.Name = fqn

	normalizePEMs(obj)

	now := timestamppb.Now()
	obj.CreateTime = now
	obj.UpdateTime = now
	obj.Etag = "abcdef0123A="

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

	return s.operations.StartLRO(ctx, req.Parent, lroMetadata, func() (proto.Message, error) {
		result := proto.Clone(obj).(*pb.TrustConfig)
		return result, nil
	})
}

func (s *CertificateManagerV1) UpdateTrustConfig(ctx context.Context, req *pb.UpdateTrustConfigRequest) (*longrunning.Operation, error) {
	reqName := req.GetTrustConfig().GetName()

	name, err := s.parseTrustConfigName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.TrustConfig{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		klog.Warningf("update_mask was not provided in request, should be required")
	}

	// Update fields based on update mask
	for _, path := range paths {
		switch path {
		case "description":
			obj.Description = req.GetTrustConfig().GetDescription()
		case "trust_stores":
			obj.TrustStores = req.GetTrustConfig().GetTrustStores()
		case "labels":
			obj.Labels = req.GetTrustConfig().GetLabels()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
		}
	}

	normalizePEMs(obj)

	now := timestamppb.Now()
	obj.UpdateTime = now
	obj.Etag = "abcdef0123A="

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroMetadata := &pb.OperationMetadata{
		ApiVersion:            "v1",
		CreateTime:            now,
		RequestedCancellation: false,
		Target:                fqn,
		Verb:                  "update",
	}
	lroPrefix := fmt.Sprintf("projects/%s/locations/global", name.Project.ID)

	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		result := proto.Clone(obj).(*pb.TrustConfig)
		return result, nil
	})
}

func (s *CertificateManagerV1) DeleteTrustConfig(ctx context.Context, req *pb.DeleteTrustConfigRequest) (*longrunning.Operation, error) {
	name, err := s.parseTrustConfigName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.TrustConfig{}
	if err := s.storage.Delete(ctx, fqn, obj); err != nil {
		return nil, err
	}

	now := timestamppb.Now()
	lroMetadata := &pb.OperationMetadata{
		ApiVersion:            "v1",
		CreateTime:            now,
		RequestedCancellation: false,
		Target:                fqn,
		Verb:                  "delete",
	}

	parent := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)

	return s.operations.StartLRO(ctx, parent, lroMetadata, func() (proto.Message, error) {
		return &emptypb.Empty{}, nil
	})
}

func normalizePEMs(trustConfig *pb.TrustConfig) {
	for _, store := range trustConfig.TrustStores {
		for _, anchor := range store.TrustAnchors {
			if anchor.PemCertificate != "" && !strings.HasSuffix(anchor.PemCertificate, "\n") {
				anchor.PemCertificate += "\n"
			}
		}
		for _, ca := range store.IntermediateCas {
			if ca.PemCertificate != "" && !strings.HasSuffix(ca.PemCertificate, "\n") {
				ca.PemCertificate += "\n"
			}
		}
	}
}
