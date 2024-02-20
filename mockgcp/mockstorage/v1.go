// Copyright 2023 Google LLC
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

package mockstorage

import (
	"context"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/storage/v1"
	"github.com/golang/protobuf/ptypes/empty"
)

type StorageV1 struct {
	*MockService
	pb.UnimplementedStorageServer
}

func (s *StorageV1) GetBucket(ctx context.Context, req *pb.GetBucketRequest) (*pb.Bucket, error) {
	fqn := req.GetBucket()

	obj := &pb.Bucket{}

	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *StorageV1) InsertBucket(ctx context.Context, req *pb.InsertBucketRequest) (*pb.Bucket, error) {
	fqn := req.GetBucket().GetName()

	now := timestamppb.Now()

	obj := proto.Clone(req.GetBucket()).(*pb.Bucket)
	obj.Name = fqn

	obj.Id = fqn
	obj.Name = fqn
	obj.ProjectNumber = 0 // todo

	obj.Location = "US"
	obj.LocationType = "multi-region"
	obj.TimeCreated = now
	obj.Updated = now

	if obj.StorageClass == "" {
		obj.StorageClass = "STANDARD"
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *StorageV1) PatchBucket(ctx context.Context, req *pb.PatchBucketRequest) (*pb.Bucket, error) {
	fqn := req.GetBucket()

	obj := &pb.Bucket{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *StorageV1) DeleteBucket(ctx context.Context, req *pb.DeleteBucketRequest) (*empty.Empty, error) {
	fqn := req.GetBucket()

	deletedObj := &pb.Bucket{}
	if err := s.storage.Delete(ctx, fqn, deletedObj); err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}
