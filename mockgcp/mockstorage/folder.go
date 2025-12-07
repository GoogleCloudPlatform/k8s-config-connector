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
// proto.service: google.storage.v1
// proto.message: google.storage.v1.Folder

package mockstorage

import (
	"context"
	"time"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/storage/v1"
)

type folder struct {
	*MockService
	pb.UnimplementedFoldersServerServer
}

func (s *folder) GetFolder(ctx context.Context, req *pb.GetFolderRequest) (*pb.Folder, error) {
	obj := &pb.Folder{}
	if err := s.storage.Get(ctx, req.GetName(), obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *folder) InsertFolder(ctx context.Context, req *pb.InsertFolderRequest) (*pb.Folder, error) {
	now := time.Now()
	generation := int64(1)
	obj := proto.Clone(req.GetFolder()).(*pb.Folder)
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.Metageneration = &generation

	if err := s.storage.Create(ctx, req.GetFolder().GetName(), obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *folder) DeleteFolder(ctx context.Context, req *pb.DeleteFolderRequest) (*emptypb.Empty, error) {
	fqn := req.GetName()
	deleted := &pb.Folder{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
