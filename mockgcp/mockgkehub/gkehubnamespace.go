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

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/gkehub/v1"
)

func (s *GkeHubV1) GetNamespace(ctx context.Context, req *pb.GetNamespaceRequest) (*pb.Namespace, error) {
	name, err := s.parseNamespaceName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Namespace{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *GkeHubV1) CreateNamespace(ctx context.Context, req *pb.CreateNamespaceRequest) (*longrunning.Operation, error) {
	reqName := req.Parent + "/namespaces/" + req.ScopeNamespaceId
	name, err := s.parseNamespaceName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := timestamppb.Now()

	obj := proto.Clone(req.Resource).(*pb.Namespace)
	obj.Name = fqn
	obj.CreateTime = now
	obj.UpdateTime = now
	obj.State = &pb.NamespaceLifecycleState{Code: pb.NamespaceLifecycleState_READY}
	obj.Uid = "111111111111111111111" // Stable UID for testing

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}
	metadata := &pb.OperationMetadata{
		Target:     fqn,
		CreateTime: now,
		EndTime:    now,
	}
	return s.operations.StartLRO(ctx, name.String(), metadata, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *GkeHubV1) UpdateNamespace(ctx context.Context, req *pb.UpdateNamespaceRequest) (*longrunning.Operation, error) {
	reqName := req.GetName()

	name, err := s.parseNamespaceName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Namespace{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}
	now := timestamppb.Now()
	obj.UpdateTime = now
	// Required. A list of fields to be updated in this request.
	paths := req.GetUpdateMask().GetPaths()

	for _, path := range paths {
		switch path {
		case "labels":
			obj.Labels = req.Resource.GetLabels()
		case "namespaceLabels":
			obj.NamespaceLabels = req.Resource.GetNamespaceLabels()
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
		return obj, nil
	})
}

func (s *GkeHubV1) DeleteNamespace(ctx context.Context, req *pb.DeleteNamespaceRequest) (*longrunning.Operation, error) {
	name, err := s.parseNamespaceName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := timestamppb.Now()

	oldObj := &pb.Namespace{}
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
	return s.operations.DoneLRO(ctx, name.String(), metadata, &emptypb.Empty{})
}
