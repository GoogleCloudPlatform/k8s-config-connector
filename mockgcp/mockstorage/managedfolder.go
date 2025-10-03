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

// +tool:mockgcp-support
// proto.service: google.storage.control.v2.StorageControl
// proto.message: google.storage.control.v2.ManagedFolder

package mockstorage

import (
	"context"
	"time"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/storage/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

type managedFolders struct {
	*MockService
	pb.UnimplementedManagedFoldersServerServer
}

func (s *managedFolders) GetManagedFolder(ctx context.Context, req *pb.GetManagedFolderRequest) (*pb.ManagedFolder, error) {
	obj := &pb.ManagedFolder{}
	if err := s.storage.Get(ctx, req.GetName(), obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *managedFolders) ListManagedFolders(ctx context.Context, req *pb.ListManagedFoldersRequest) (*pb.ManagedFolders, error) {
	findPrefix := req.GetPrefix()

	response := &pb.ManagedFolders{}

	findKind := (&pb.ManagedFolder{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, findKind, storage.ListOptions{Prefix: findPrefix}, func(obj proto.Message) error {
		managedFolder := obj.(*pb.ManagedFolder)
		response.Items = append(response.Items, managedFolder)
		return nil
	}); err != nil {
		return nil, err
	}
	return response, nil
}
func (s *managedFolders) InsertManagedFolder(ctx context.Context, req *pb.InsertManagedFolderRequest) (*pb.ManagedFolder, error) {
	now := time.Now()

	generation := int64(1)
	obj := proto.Clone(req.GetManagedFolder()).(*pb.ManagedFolder)
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.Metageneration = &generation

	if err := s.storage.Create(ctx, req.GetManagedFolder().GetName(), obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *managedFolders) DeleteManagedFolder(ctx context.Context, req *pb.DeleteManagedFolderRequest) (*emptypb.Empty, error) {
	deleted := &pb.ManagedFolder{}
	if err := s.storage.Delete(ctx, req.GetName(), deleted); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
