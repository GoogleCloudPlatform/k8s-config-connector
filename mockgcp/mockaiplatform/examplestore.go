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

// +tool:mockgcp-support
// proto.service: google.cloud.aiplatform.v1beta1.ExampleStoreService
// proto.message: google.cloud.aiplatform.v1beta1.ExampleStore

package mockaiplatform

import (
	"context"
	"fmt"
	"strings"
	"time"

	longrunning "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

type exampleStoreService struct {
	*MockService
	pb.UnimplementedExampleStoreServiceServer
}

func (s *exampleStoreService) GetExampleStore(ctx context.Context, req *pb.GetExampleStoreRequest) (*pb.ExampleStore, error) {
	name, err := s.parseExampleStoreName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.ExampleStore{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *exampleStoreService) CreateExampleStore(ctx context.Context, req *pb.CreateExampleStoreRequest) (*longrunning.Operation, error) {
	reqName := req.ExampleStore.GetName()
	name, err := s.parseExampleStoreName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	obj := proto.CloneOf(req.ExampleStore)
	obj.Name = fqn

	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.CreateExampleStoreOperationMetadata{}
	op.GenericMetadata = &pb.GenericOperationMetadata{
		CreateTime: timestamppb.New(now),
		UpdateTime: timestamppb.New(now),
	}
	opPrefix := name.String()
	return s.operations.StartLRO(ctx, opPrefix, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *exampleStoreService) UpdateExampleStore(ctx context.Context, req *pb.UpdateExampleStoreRequest) (*longrunning.Operation, error) {
	reqName := req.ExampleStore.GetName()
	name, err := s.parseExampleStoreName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	obj := &pb.ExampleStore{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		obj.DisplayName = req.ExampleStore.DisplayName
		obj.Description = req.ExampleStore.Description
	} else {
		for _, path := range paths {
			switch path {
			case "display_name", "displayName":
				obj.DisplayName = req.ExampleStore.DisplayName
			case "description":
				obj.Description = req.ExampleStore.Description
			}
		}
	}

	obj.UpdateTime = timestamppb.New(now)

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.UpdateExampleStoreOperationMetadata{}
	op.GenericMetadata = &pb.GenericOperationMetadata{
		CreateTime: timestamppb.New(now),
		UpdateTime: timestamppb.New(now),
	}
	opPrefix := fqn
	return s.operations.StartLRO(ctx, opPrefix, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *exampleStoreService) ListExampleStores(ctx context.Context, req *pb.ListExampleStoresRequest) (*pb.ListExampleStoresResponse, error) {
	project, err := s.Projects.GetProjectByID(req.Parent)
	if err != nil {
		return nil, err
	}

	findPrefix := fmt.Sprintf("projects/%v/", project.ID)

	var exampleStores []*pb.ExampleStore

	exampleStoreKind := (&pb.ExampleStore{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, exampleStoreKind, storage.ListOptions{}, func(obj proto.Message) error {
		exampleStore := obj.(*pb.ExampleStore)
		if strings.HasPrefix(exampleStore.GetName(), findPrefix) {
			exampleStores = append(exampleStores, exampleStore)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return &pb.ListExampleStoresResponse{
		ExampleStores: exampleStores,
	}, nil
}

func (s *exampleStoreService) DeleteExampleStore(ctx context.Context, req *pb.DeleteExampleStoreRequest) (*longrunning.Operation, error) {
	name, err := s.parseExampleStoreName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	deleted := &pb.ExampleStore{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	op := &pb.DeleteExampleStoreOperationMetadata{}
	op.GenericMetadata = &pb.GenericOperationMetadata{
		CreateTime: timestamppb.New(now),
		UpdateTime: timestamppb.New(now),
	}
	opPrefix := fqn
	return s.operations.DoneLRO(ctx, opPrefix, op, &emptypb.Empty{})
}

type ExampleStoreName struct {
	Project        *projects.ProjectData
	Location       string
	ExampleStoreId string
}

func (s *exampleStoreService) parseExampleStoreName(name string) (*ExampleStoreName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "exampleStores" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}
		return &ExampleStoreName{
			Project:        project,
			Location:       tokens[3],
			ExampleStoreId: tokens[5],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "invalid name %q", name)
}

func (n *ExampleStoreName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/exampleStores/%s", n.Project.ID, n.Location, n.ExampleStoreId)
}
