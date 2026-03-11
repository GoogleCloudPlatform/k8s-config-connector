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
	"strconv"
	"strings"

	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/certificatemanager/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

func (s *CertificateManagerV1) GetCertificateIssuanceConfig(ctx context.Context, req *pb.GetCertificateIssuanceConfigRequest) (*pb.CertificateIssuanceConfig, error) {
	name, err := s.parseCertificateIssuanceConfigName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.CertificateIssuanceConfig{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *CertificateManagerV1) CreateCertificateIssuanceConfig(ctx context.Context, req *pb.CreateCertificateIssuanceConfigRequest) (*longrunning.Operation, error) {
	reqName := req.Parent + "/certificateIssuanceConfigs/" + req.CertificateIssuanceConfigId
	name, err := s.parseCertificateIssuanceConfigName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.CertificateIssuanceConfig).(*pb.CertificateIssuanceConfig)
	obj.Name = fqn

	now := timestamppb.Now()
	obj.CreateTime = now
	obj.UpdateTime = now

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroMetadata := &pb.OperationMetadata{
		ApiVersion: "v1",
		CreateTime: now,
		Target:     fqn,
		Verb:       "create",
	}

	return s.operations.StartLRO(ctx, req.Parent, lroMetadata, func() (proto.Message, error) {
		result := proto.Clone(obj).(*pb.CertificateIssuanceConfig)
		return result, nil
	})
}

func (s *CertificateManagerV1) ListCertificateIssuanceConfigs(ctx context.Context, req *pb.ListCertificateIssuanceConfigsRequest) (*pb.ListCertificateIssuanceConfigsResponse, error) {
	objs := []*pb.CertificateIssuanceConfig{}
	kind := (&pb.CertificateIssuanceConfig{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, kind, storage.ListOptions{}, func(obj proto.Message) error {
		objs = append(objs, obj.(*pb.CertificateIssuanceConfig))
		return nil
	}); err != nil {
		return nil, err
	}

	// Filter by location
	tokens := strings.Split(req.Parent, "/")
	if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "locations" {
		location := tokens[3]
		filtered := []*pb.CertificateIssuanceConfig{}
		for _, obj := range objs {
			name, err := s.parseCertificateIssuanceConfigName(obj.Name)
			if err == nil && name.Location == location {
				filtered = append(filtered, obj)
			}
		}
		objs = filtered
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "parent %q is not valid", req.Parent)
	}

	// Simple pagination implementation
	pageSize := int(req.GetPageSize())
	if pageSize == 0 {
		pageSize = 100 // Default page size
	}
	if pageSize > 1000 {
		pageSize = 1000 // Max page size
	}

	startIndex := 0
	if req.GetPageToken() != "" {
		var err error
		startIndex, err = strconv.Atoi(req.GetPageToken())
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "page_token %q is not valid", req.GetPageToken())
		}
	}

	if startIndex < 0 || startIndex > len(objs) {
		startIndex = len(objs)
	}

	endIndex := startIndex + pageSize
	if endIndex > len(objs) {
		endIndex = len(objs)
	}

	nextPageToken := ""
	if endIndex < len(objs) {
		nextPageToken = strconv.Itoa(endIndex)
	}

	return &pb.ListCertificateIssuanceConfigsResponse{
		CertificateIssuanceConfigs: objs[startIndex:endIndex],
		NextPageToken:              nextPageToken,
	}, nil
}

func (s *CertificateManagerV1) DeleteCertificateIssuanceConfig(ctx context.Context, req *pb.DeleteCertificateIssuanceConfigRequest) (*longrunning.Operation, error) {
	name, err := s.parseCertificateIssuanceConfigName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deletedObj := &pb.CertificateIssuanceConfig{}
	if err := s.storage.Delete(ctx, fqn, deletedObj); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		ApiVersion: "v1",
		CreateTime: timestamppb.Now(),
		Target:     fqn,
		Verb:       "delete",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		return &emptypb.Empty{}, nil
	})
}
