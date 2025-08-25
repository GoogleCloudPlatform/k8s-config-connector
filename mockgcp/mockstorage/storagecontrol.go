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
	"log"
	"time"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	// Note we use "real" protos (not mockgcp) ones as it's GRPC API.
	pb "cloud.google.com/go/storage/control/apiv2/controlpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"google.golang.org/protobuf/proto"
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

	obj := proto.Clone(req.GetAnywhereCache()).(*pb.AnywhereCache)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.State = anywhereCacheStateCreating

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}
	op, err := s.operations.StartLRO(ctx, fqn, &pb.CreateAnywhereCacheMetadata{AnywhereCacheId: &zone}, func() (proto.Message, error) {
		result := proto.Clone(obj).(*pb.AnywhereCache)
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
		result := proto.Clone(obj).(*pb.AnywhereCache)
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

func (s *StorageControlService) GetManagedFolder(ctx context.Context, req *pb.GetManagedFolderRequest) (*pb.ManagedFolder, error) {
	obj := &pb.ManagedFolder{}
	// Since a ManagedFolder's Name is stored with a trailing "/", we need to add it to the end of the request Name.
	storedName := req.GetName() + "/"
	if err := s.storage.Get(ctx, storedName, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The specified managed folder does not exist.")
		}
		log.Printf("An unexpected error occurred: %v", err)
		return nil, err
	}
	return obj, nil
}

func (s *StorageControlService) ListManagedFolders(ctx context.Context, req *pb.ListManagedFoldersRequest) (*pb.ListManagedFoldersResponse, error) {
	findPrefix := req.GetPrefix()

	response := &pb.ListManagedFoldersResponse{}

	findKind := (&pb.ManagedFolder{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, findKind, storage.ListOptions{Prefix: findPrefix}, func(obj proto.Message) error {
		managedFolder := obj.(*pb.ManagedFolder)
		response.ManagedFolders = append(response.ManagedFolders, managedFolder)
		return nil
	}); err != nil {
		return nil, err
	}
	return response, nil
}
func (s *StorageControlService) CreateManagedFolder(ctx context.Context, req *pb.CreateManagedFolderRequest) (*pb.ManagedFolder, error) {
	fqn := fmt.Sprintf("%s/managedFolders/%s/", req.GetParent(), req.GetManagedFolderId())

	now := time.Now()
	generation := int64(1)

	obj := proto.Clone(req.GetManagedFolder()).(*pb.ManagedFolder)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.Metageneration = generation

	log.Printf("Checking for an empty name: %s", obj.Name)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *StorageControlService) DeleteManagedFolder(ctx context.Context, req *pb.DeleteManagedFolderRequest) (*emptypb.Empty, error) {
	deleted := &pb.ManagedFolder{}
	// Since a ManagedFolder's Name is stored with a trailing "/", we need to add it to the end of the request Name.
	storedName := req.GetName() + "/"
	if err := s.storage.Delete(ctx, storedName, deleted); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
