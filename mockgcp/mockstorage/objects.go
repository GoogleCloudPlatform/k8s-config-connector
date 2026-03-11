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
	"fmt"
	"strings"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/httpmux"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/storage/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

type objects struct {
	*MockService
	pb.UnimplementedObjectsServerServer
}

func (s *objects) ListObjects(ctx context.Context, req *pb.ListObjectsRequest) (*pb.Objects, error) {
	httpmux.SetExpiresHeader(ctx, time.Now())

	ret := &pb.Objects{}
	ret.Kind = PtrTo("storage#objects")

	kind := (&pb.Object{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, kind, storage.ListOptions{}, func(obj proto.Message) error {
		o := obj.(*pb.Object)
		if o.GetBucket() == req.GetBucket() {
			// Basic prefix matching for GCS
			if req.GetPrefix() != "" && !strings.HasPrefix(o.GetName(), req.GetPrefix()) {
				return nil
			}

			// GCS ListObjects also returns prefixes (delimited by /)
			if req.GetDelimiter() == "/" {
				name := o.GetName()
				if req.GetPrefix() != "" {
					name = strings.TrimPrefix(name, req.GetPrefix())
				}
				if i := strings.Index(name, "/"); i != -1 {
					prefix := o.GetName()[:len(o.GetName())-len(name)+i+1]
					found := false
					for _, p := range ret.Prefixes {
						if p == prefix {
							found = true
							break
						}
					}
					if !found {
						ret.Prefixes = append(ret.Prefixes, prefix)
					}
					return nil
				}
			}

			ret.Items = append(ret.Items, o)
		}
		return nil
	}); err != nil {
		return nil, status.Errorf(codes.Internal, "error listing objects: %v", err)
	}

	return ret, nil
}

func (s *objects) InsertObject(ctx context.Context, req *pb.InsertObjectRequest) (*pb.Object, error) {
	obj := req.GetObject()
	if obj == nil {
		obj = &pb.Object{}
	}
	if obj.Bucket == nil {
		obj.Bucket = PtrTo(req.GetBucket())
	}
	if obj.Name == nil {
		obj.Name = PtrTo(req.GetName())
	}

	if obj.GetBucket() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "bucket is required")
	}
	if obj.GetName() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "name is required")
	}

	obj.Kind = PtrTo("storage#object")
	obj.Generation = PtrTo(int64(1))
	obj.Metageneration = PtrTo(int64(1))
	obj.Etag = PtrTo("mock-etag")

	fqn := fmt.Sprintf("b/%s/o/%s", obj.GetBucket(), obj.GetName())

	// Check if already exists to allow overwriting
	existing := &pb.Object{}
	if err := s.storage.Get(ctx, fqn, existing); err == nil {
		obj.Generation = PtrTo(existing.GetGeneration() + 1)
		if err := s.storage.Update(ctx, fqn, obj); err != nil {
			return nil, status.Errorf(codes.Internal, "error updating object: %v", err)
		}
	} else {
		if err := s.storage.Create(ctx, fqn, obj); err != nil {
			return nil, status.Errorf(codes.Internal, "error creating object: %v", err)
		}
	}

	return obj, nil
}

func (s *objects) GetObject(ctx context.Context, req *pb.GetObjectRequest) (*pb.Object, error) {
	fqn := fmt.Sprintf("b/%s/o/%s", req.GetBucket(), req.GetName())
	obj := &pb.Object{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *objects) DeleteObject(ctx context.Context, req *pb.DeleteObjectRequest) (*empty.Empty, error) {
	fqn := fmt.Sprintf("b/%s/o/%s", req.GetBucket(), req.GetName())
	obj := &pb.Object{}
	if err := s.storage.Delete(ctx, fqn, obj); err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}
