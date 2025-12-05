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

// +tool:mockgcp-support-but-no-proto
// proto.service: google.storage.v1.FoldersServer
// proto.message: google.storage.v1.Folder

package mockstorage

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/httpmux"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/storage/v1"
)

type folder struct {
	*MockService
	pb.UnimplementedFoldersServerServer
}

func (s *folder) GetFolder(ctx context.Context, req *pb.GetFolderRequest) (*pb.Folder, error) {
	name, err := s.buildFolderName(req.GetBucket(), req.GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.Folder{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	httpmux.SetExpiresHeader(ctx, time.Now())

	return obj, nil
}

func (s *folder) InsertFolder(ctx context.Context, req *pb.InsertFolderRequest) (*pb.Folder, error) {
	name, err := s.buildFolderName(req.GetBucket(), req.GetFolder().GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	now := timestamppb.Now()

	generation := int64(1)
	obj := proto.Clone(req.GetFolder()).(*pb.Folder)
	obj.Bucket = &name.Bucket
	obj.CreateTime = now
	obj.UpdateTime = now
	obj.Id = PtrTo(fmt.Sprintf("%d", time.Now().UnixNano()))
	obj.Metageneration = &generation
	obj.Kind = PtrTo("storage#folder")
	obj.SelfLink = PtrTo(fmt.Sprintf("https://www.googleapis.com/storage/v1/b/%s/folders/%s", name.Bucket, name.Folder))
	obj.Name = PtrTo(name.Folder + "/")

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *folder) DeleteFolder(ctx context.Context, req *pb.DeleteFolderRequest) (*emptypb.Empty, error) {
	name, err := s.buildFolderName(req.GetBucket(), req.GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	deleted := &pb.Folder{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	httpmux.SetStatusCode(ctx, http.StatusNoContent)

	return &emptypb.Empty{}, nil
}

type folderName struct {
	Bucket string
	Folder string
}

func (n *folderName) String() string {
	return fmt.Sprintf("buckets/%s/folders/%s", n.Bucket, n.Folder)
}

// buildFolderName builds a folderName from the constituent parts.
func (s *MockService) buildFolderName(bucket string, folder string) (*folderName, error) {
	folder = strings.TrimSuffix(folder, "/")

	name := &folderName{
		Bucket: bucket,
		Folder: folder,
	}

	return name, nil
}
