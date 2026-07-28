// Copyright 2025 Google LLC
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
// proto.message: google.storage.control.v2.AnywhereCache

package mockstorage

import (
	"context"
	"fmt"
	"strings"
	"time"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	// Note we use "real" protos (not mockgcp) ones as it's GRPC API.
	pb "cloud.google.com/go/storage/control/apiv2/controlpb"
	storagepb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/storage/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// AnywhereCache States (lowercase representation from GCP API).
	anywhereCacheStateCreating = "creating"
	anywhereCacheStateRunning  = "running"
	anywhereCacheStatePaused   = "paused"
	anywhereCacheStateDisabled = "disabled"
)

type StorageControlService struct {
	*MockService
	pb.UnimplementedStorageControlServer
}

func (s *StorageControlService) GetAnywhereCache(ctx context.Context, req *pb.GetAnywhereCacheRequest) (*pb.AnywhereCache, error) {
	fqn := req.GetName()
	ret := &pb.AnywhereCache{}
	if err := s.storage.Get(ctx, fqn, ret); err != nil {
		return nil, err
	}
	return ret, nil
}

func (s *StorageControlService) CreateAnywhereCache(ctx context.Context, req *pb.CreateAnywhereCacheRequest) (*longrunningpb.Operation, error) {
	zone := req.GetAnywhereCache().GetZone()
	fqn := fmt.Sprintf("%s/anywhereCaches/%s", req.GetParent(), zone)

	now := time.Now()

	obj := proto.CloneOf(req.GetAnywhereCache())
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.State = anywhereCacheStateCreating

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}
	op, err := s.operations.StartLRO(ctx, fqn, &pb.CreateAnywhereCacheMetadata{AnywhereCacheId: &zone}, func() (proto.Message, error) {
		result := proto.CloneOf(obj)
		result.State = anywhereCacheStateRunning
		if err := s.storage.Update(ctx, fqn, result); err != nil {
			return nil, err
		}
		return result, nil
	})
	if err != nil {
		return op, err
	}
	return op, err

}

func (s *StorageControlService) UpdateAnywhereCache(ctx context.Context, req *pb.UpdateAnywhereCacheRequest) (*longrunningpb.Operation, error) {
	fqn := req.GetAnywhereCache().GetName()

	obj := &pb.AnywhereCache{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	now := time.Now()
	obj.UpdateTime = timestamppb.New(now)

	if patch := req.GetAnywhereCache(); patch != nil {
		if patch.AdmissionPolicy != "" {
			obj.AdmissionPolicy = patch.AdmissionPolicy
		}
		if patch.Ttl != nil {
			obj.Ttl = patch.Ttl
		}
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op, err := s.operations.StartLRO(ctx, fqn, &pb.AnywhereCache{}, func() (proto.Message, error) {
		result := proto.CloneOf(obj)
		if err := s.storage.Update(ctx, fqn, result); err != nil {
			return nil, err
		}
		return result, nil
	})
	if err != nil {
		return op, err
	}
	return op, err
}

func (s *StorageControlService) ListAnywhereCaches(ctx context.Context, req *pb.ListAnywhereCachesRequest) (*pb.ListAnywhereCachesResponse, error) {
	var caches []*pb.AnywhereCache
	fqn_parent := req.GetParent()

	cacheKind := (&pb.AnywhereCache{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, cacheKind, storage.ListOptions{Prefix: fqn_parent}, func(obj proto.Message) error {
		cache := obj.(*pb.AnywhereCache)
		caches = append(caches, cache)
		return nil
	}); err != nil {
		return nil, err
	}

	return &pb.ListAnywhereCachesResponse{
		AnywhereCaches: caches,
	}, nil
}

func (s *StorageControlService) PauseAnywhereCache(ctx context.Context, req *pb.PauseAnywhereCacheRequest) (*pb.AnywhereCache, error) {
	fqn := req.GetName()

	obj := &pb.AnywhereCache{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	now := time.Now()
	obj.UpdateTime = timestamppb.New(now)
	obj.State = anywhereCacheStatePaused

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *StorageControlService) ResumeAnywhereCache(ctx context.Context, req *pb.ResumeAnywhereCacheRequest) (*pb.AnywhereCache, error) {
	fqn := req.GetName()

	obj := &pb.AnywhereCache{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	now := time.Now()
	obj.UpdateTime = timestamppb.New(now)
	obj.State = anywhereCacheStateRunning

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *StorageControlService) DisableAnywhereCache(ctx context.Context, req *pb.DisableAnywhereCacheRequest) (*pb.AnywhereCache, error) {
	fqn := req.GetName()

	obj := &pb.AnywhereCache{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	now := time.Now()
	obj.UpdateTime = timestamppb.New(now)
	obj.State = anywhereCacheStateDisabled

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func parseBucketIDFromControlName(name string) (string, error) {
	// e.g. "projects/_/buckets/bucket-123/managedFolders/folder-456"
	// e.g. "projects/_/buckets/bucket-123"
	tokens := strings.Split(name, "/")
	if len(tokens) >= 4 && tokens[0] == "projects" && tokens[2] == "buckets" {
		return tokens[3], nil
	}
	return "", fmt.Errorf("invalid name: %s", name)
}

func (s *StorageControlService) checkBucketExists(ctx context.Context, bucketID string) error {
	bucketObj := &storagepb.Bucket{}
	if err := s.storage.Get(ctx, "buckets/"+bucketID, bucketObj); err != nil {
		// Return standard NotFound error if bucket doesn't exist.
		return status.Errorf(codes.NotFound, "The specified bucket does not exist.")
	}
	return nil
}

func (s *StorageControlService) GetFolder(ctx context.Context, req *pb.GetFolderRequest) (*pb.Folder, error) {
	fqn := req.GetName()
	bucketID, err := parseBucketIDFromControlName(fqn)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid name: %v", err)
	}
	if err := s.checkBucketExists(ctx, bucketID); err != nil {
		return nil, err
	}

	ret := &pb.Folder{}
	if err := s.storage.Get(ctx, fqn, ret); err != nil {
		return nil, err
	}
	if !strings.HasSuffix(ret.Name, "/") {
		ret.Name += "/"
	}
	return ret, nil
}

func (s *StorageControlService) CreateFolder(ctx context.Context, req *pb.CreateFolderRequest) (*pb.Folder, error) {
	parent := req.GetParent()
	bucketID, err := parseBucketIDFromControlName(parent)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid parent: %v", err)
	}
	if err := s.checkBucketExists(ctx, bucketID); err != nil {
		return nil, err
	}

	folderID := req.GetFolderId()
	fqn := fmt.Sprintf("%s/folders/%s", parent, folderID)

	now := time.Now()
	obj := proto.Clone(req.GetFolder()).(*pb.Folder)
	obj.Name = fqn
	if !strings.HasSuffix(obj.Name, "/") {
		obj.Name += "/"
	}
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.Metageneration = 1

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *StorageControlService) DeleteFolder(ctx context.Context, req *pb.DeleteFolderRequest) (*emptypb.Empty, error) {
	fqn := req.GetName()
	bucketID, err := parseBucketIDFromControlName(fqn)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid name: %v", err)
	}
	if err := s.checkBucketExists(ctx, bucketID); err != nil {
		return nil, err
	}

	deleted := &pb.Folder{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *StorageControlService) GetManagedFolder(ctx context.Context, req *pb.GetManagedFolderRequest) (*pb.ManagedFolder, error) {
	fqn := req.GetName()
	bucketID, err := parseBucketIDFromControlName(fqn)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid name: %v", err)
	}
	if err := s.checkBucketExists(ctx, bucketID); err != nil {
		return nil, err
	}

	ret := &pb.ManagedFolder{}
	if err := s.storage.Get(ctx, fqn, ret); err != nil {
		return nil, err
	}
	if !strings.HasSuffix(ret.Name, "/") {
		ret.Name += "/"
	}
	return ret, nil
}

func (s *StorageControlService) CreateManagedFolder(ctx context.Context, req *pb.CreateManagedFolderRequest) (*pb.ManagedFolder, error) {
	parent := req.GetParent()
	bucketID, err := parseBucketIDFromControlName(parent)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid parent: %v", err)
	}
	if err := s.checkBucketExists(ctx, bucketID); err != nil {
		return nil, err
	}

	managedFolderID := req.GetManagedFolderId()
	fqn := fmt.Sprintf("%s/managedFolders/%s", parent, managedFolderID)

	now := time.Now()
	obj := proto.Clone(req.GetManagedFolder()).(*pb.ManagedFolder)
	obj.Name = fqn
	if !strings.HasSuffix(obj.Name, "/") {
		obj.Name += "/"
	}
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.Metageneration = 1

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *StorageControlService) DeleteManagedFolder(ctx context.Context, req *pb.DeleteManagedFolderRequest) (*emptypb.Empty, error) {
	fqn := req.GetName()
	bucketID, err := parseBucketIDFromControlName(fqn)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid name: %v", err)
	}
	if err := s.checkBucketExists(ctx, bucketID); err != nil {
		return nil, err
	}

	deleted := &pb.ManagedFolder{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
