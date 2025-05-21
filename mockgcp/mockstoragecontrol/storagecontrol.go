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

package mockstoragecontrol

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"

	// Note we use "real" protos (not mockgcp) ones as it's GRPC API.
	pb "cloud.google.com/go/storage/control/apiv2/controlpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"google.golang.org/protobuf/proto"
	// empty "github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

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
	obj.State = "running"

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}
	op, err := s.operations.StartLRO(ctx, fqn, &pb.CreateAnywhereCacheMetadata{AnywhereCacheId: &zone}, func() (proto.Message, error) {
		result := proto.Clone(obj).(*pb.AnywhereCache)
		return result, nil
	})
	if err != nil {
		return op, err
	}
	response, err := anypb.New(obj)
	if err != nil {
		return op, err
	}
	op.Result = &longrunningpb.Operation_Response{
		Response: response,
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
		return result, nil
	})
	if err != nil {
		return op, err
	}
	response, err := anypb.New(obj)
	if err != nil {
		return op, err
	}
	op.Result = &longrunningpb.Operation_Response{
		Response: response,
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
		AnywhereCaches:         caches,
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
	obj.State = "paused"

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
	obj.State = "running"

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
	obj.State = "disabled"

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}


// func (s *StorageControlService) CreateFolder(ctx context.Context, req *pb.CreateFolderRequest) (*pb.Folder, error) {
// 	return nil, fmt.Errorf("method CreateFolder not implemented")
// }

// func (s *StorageControlService) DeleteFolder(ctx context.Context, req *pb.DeleteFolderRequest) (*empty.Empty, error) {
// 	return nil, fmt.Errorf("method DeleteFolder not implemented")
// }

// func (s *StorageControlService) GetFolder(ctx context.Context, req *pb.GetFolderRequest) (*pb.Folder, error) {
// 	return nil, fmt.Errorf("method GetFolder not implemented")
// }

// func (s *StorageControlService) ListFolders(ctx context.Context, req *pb.ListFoldersRequest) (*pb.ListFoldersResponse, error) {
// 	return nil, fmt.Errorf("method ListFolders not implemented")
// }

// func (s *StorageControlService) RenameFolder(ctx context.Context, req *pb.RenameFolderRequest) (*longrunningpb.Operation, error) {
// 	return nil, fmt.Errorf("method RenameFolder not implemented")
// }

// func (s *StorageControlService) GetStorageLayout(ctx context.Context, req *pb.GetStorageLayoutRequest) (*pb.StorageLayout, error) {
// 	return nil, fmt.Errorf("method GetStorageLayout not implemented")
// }

// func (s *StorageControlService) CreateManagedFolder(ctx context.Context, req *pb.CreateManagedFolderRequest) (*pb.ManagedFolder, error) {
// 	return nil, fmt.Errorf("method CreateManagedFolder not implemented")
// }

// func (s *StorageControlService) DeleteManagedFolder(ctx context.Context, req *pb.DeleteManagedFolderRequest) (*empty.Empty, error) {
// 	return nil, fmt.Errorf("method DeleteManagedFolder not implemented")
// }

// func (s *StorageControlService) GetManagedFolder(ctx context.Context, req *pb.GetManagedFolderRequest) (*pb.ManagedFolder, error) {
// 	return nil, fmt.Errorf("method GetManagedFolder not implemented")
// }

// func (s *StorageControlService) ListManagedFolders(ctx context.Context, req *pb.ListManagedFoldersRequest) (*pb.ListManagedFoldersResponse, error) {
// 	return nil, fmt.Errorf("method ListManagedFolders not implemented")
// }
