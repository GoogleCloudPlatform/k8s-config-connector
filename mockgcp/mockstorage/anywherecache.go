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

package mockstorage

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/storage/v1"
	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

type anywhereCaches struct {
	*MockService
	pb.UnimplementedAnywhereCachesServerServer
}

func (s *anywhereCaches) GetAnywhereCache(ctx context.Context, req *pb.GetAnywhereCacheRequest) (*pb.AnywhereCache, error) {
	fqn := fmt.Sprintf("projects/_/buckets/%s/anywhereCaches/%s", req.GetBucket(), req.GetAnywhereCacheId())
	ret := &pb.AnywhereCache{}
	if err := s.storage.Get(ctx, fqn, ret); err != nil {
		return nil, err
	}
	return ret, nil
}

func (s *anywhereCaches) InsertAnywhereCache(ctx context.Context, req *pb.InsertAnywhereCacheRequest) (*longrunningpb.Operation, error) {
	fqn := fmt.Sprintf("projects/_/buckets/%s/anywhereCaches/%s", req.GetBucket(), req.GetAnywhereCache().GetZone())

	now := time.Now()

	obj := proto.Clone(req.GetAnywhereCache()).(*pb.AnywhereCache)
	obj.Id = PtrTo(fqn)
	obj.Kind = PtrTo("storage#anywhereCache")
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.State = PtrTo("creating")

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
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


func (s *anywhereCaches) UpdateAnywhereCache(ctx context.Context, req *pb.UpdateAnywhereCacheRequest) (*longrunningpb.Operation, error) {
	fqn := fmt.Sprintf("projects/_/buckets/%s/anywhereCaches/%s", req.GetBucket(), req.GetAnywhereCacheId())

	obj := &pb.AnywhereCache{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	now := time.Now()
	obj.UpdateTime = timestamppb.New(now)

	if patch := req.GetAnywhereCache(); patch != nil {
		if patch.AdmissionPolicy != nil {
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

func (s *anywhereCaches) ListAnywhereCaches(ctx context.Context, req *pb.ListAnywhereCachesRequest) (*pb.AnywhereCaches, error) {
	var caches []*pb.AnywhereCache
	fqn_parent := fmt.Sprintf("projects/_/buckets/%s", req.GetBucket())

	cacheKind := (&pb.AnywhereCache{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, cacheKind, storage.ListOptions{Prefix: fqn_parent}, func(obj proto.Message) error {
		cache := obj.(*pb.AnywhereCache)
		caches = append(caches, cache)
		return nil
	}); err != nil {
		return nil, err
	}

	return &pb.AnywhereCaches{
		Items:         caches,
		NextPageToken: nil,
		Kind:          PtrTo("storage#anywhereCaches"),
	}, nil
}

func (s *anywhereCaches) PauseAnywhereCache(ctx context.Context, req *pb.PauseAnywhereCacheRequest) (*pb.AnywhereCache, error) {
	fqn := fmt.Sprintf("projects/_/buckets/%s/anywhereCaches/%s", req.GetBucket(), req.GetAnywhereCacheId())

	obj := &pb.AnywhereCache{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	now := time.Now()
	obj.UpdateTime = timestamppb.New(now)
	obj.State = PtrTo("paused")

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *anywhereCaches) ResumeAnywhereCache(ctx context.Context, req *pb.ResumeAnywhereCacheRequest) (*pb.AnywhereCache, error) {
	fqn := fmt.Sprintf("projects/_/buckets/%s/anywhereCaches/%s", req.GetBucket(), req.GetAnywhereCacheId())

	obj := &pb.AnywhereCache{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	now := time.Now()
	obj.UpdateTime = timestamppb.New(now)
	obj.State = PtrTo("running")

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *anywhereCaches) DisableAnywhereCache(ctx context.Context, req *pb.DisableAnywhereCacheRequest) (*pb.AnywhereCache, error) {
	fqn := fmt.Sprintf("projects/_/buckets/%s/anywhereCaches/%s", req.GetBucket(), req.GetAnywhereCacheId())

	obj := &pb.AnywhereCache{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	now := time.Now()
	obj.UpdateTime = timestamppb.New(now)
	obj.State = PtrTo("disabled")

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}
