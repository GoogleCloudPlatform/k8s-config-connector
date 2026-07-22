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
	"k8s.io/klog/v2"

	pb "cloud.google.com/go/certificatemanager/apiv1/certificatemanagerpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
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

	now := timestamppb.Now()
	obj.CreateTime = now
	obj.UpdateTime = now

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		ApiVersion: "v1",
		CreateTime: now,
		Target:     fqn,
		Verb:       "create",
	}

	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		result := proto.CloneOf(obj)
		return result, nil
	})
}

func (s *CertificateManagerV1) ListTrustConfigs(ctx context.Context, req *pb.ListTrustConfigsRequest) (*pb.ListTrustConfigsResponse, error) {
	objs := []*pb.TrustConfig{}
	kind := (&pb.TrustConfig{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, kind, storage.ListOptions{}, func(obj proto.Message) error {
		objs = append(objs, obj.(*pb.TrustConfig))
		return nil
	}); err != nil {
		return nil, err
	}

	// Filter by location
	tokens := strings.Split(req.Parent, "/")
	if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "locations" {
		location := tokens[3]
		filtered := []*pb.TrustConfig{}
		for _, obj := range objs {
			name, err := s.parseTrustConfigName(obj.Name)
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

	return &pb.ListTrustConfigsResponse{
		TrustConfigs:  objs[startIndex:endIndex],
		NextPageToken: nextPageToken,
	}, nil
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
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		klog.Warningf("update_mask was not provided in request, should be required")
	}

	for _, path := range paths {
		switch path {
		case "description":
			obj.Description = req.GetTrustConfig().GetDescription()
		case "labels":
			obj.Labels = req.GetTrustConfig().GetLabels()
		case "trust_stores", "trustStores":
			obj.TrustStores = req.GetTrustConfig().GetTrustStores()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
		}
	}

	now := timestamppb.Now()
	obj.UpdateTime = now

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		ApiVersion:            "v1",
		CreateTime:            now,
		RequestedCancellation: false,
		Target:                fqn,
		Verb:                  "update",
	}

	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		result := proto.CloneOf(obj)
		return result, nil
	})
}

func (s *CertificateManagerV1) DeleteTrustConfig(ctx context.Context, req *pb.DeleteTrustConfigRequest) (*longrunning.Operation, error) {
	name, err := s.parseTrustConfigName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	oldObj := &pb.TrustConfig{}
	if err := s.storage.Delete(ctx, fqn, oldObj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
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
		lroMetadata.EndTime = timestamppb.Now()
		return &emptypb.Empty{}, nil
	})
}
