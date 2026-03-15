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
	"google.golang.org/protobuf/proto"
)


type objects struct {
	*MockService
	pb.UnimplementedObjectsServerServer
}

func (s *objects) ListObjects(ctx context.Context, req *pb.ListObjectsRequest) (*pb.Objects, error) {
	httpmux.SetExpiresHeader(ctx, time.Now())

	var items []*pb.Object
	prefixes := make(map[string]bool)

	kind := (&pb.Object{}).ProtoReflect().Descriptor()
	prefix := fmt.Sprintf("buckets/%s/objects/", req.GetBucket())
	if err := s.storage.List(ctx, kind, storage.ListOptions{Prefix: prefix}, func(obj proto.Message) error {
		object := obj.(*pb.Object)

		name := object.GetName()
		if req.GetPrefix() != "" && !strings.HasPrefix(name, req.GetPrefix()) {
			return nil
		}

		if req.GetDelimiter() != "" {
			relativeName := strings.TrimPrefix(name, req.GetPrefix())
			if idx := strings.Index(relativeName, req.GetDelimiter()); idx != -1 {
				prefix := req.GetPrefix() + relativeName[:idx+1]
				prefixes[prefix] = true
				return nil
			}
		}

		items = append(items, object)
		return nil
	}); err != nil {
		return nil, err
	}

	ret := &pb.Objects{}
	ret.Kind = PtrTo("storage#objects")
	ret.Items = items
	for p := range prefixes {
		ret.Prefixes = append(ret.Prefixes, p)
	}
	return ret, nil
}

func (s *objects) GetObject(ctx context.Context, req *pb.GetObjectRequest) (*pb.Object, error) {
	fqn := fmt.Sprintf("buckets/%s/objects/%s", req.GetBucket(), req.GetName())

	obj := &pb.Object{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *objects) InsertObject(ctx context.Context, req *pb.InsertObjectRequest) (*pb.Object, error) {
	bucketName := req.GetBucket()
	objectName := req.GetObject().GetName()

	fqn := fmt.Sprintf("buckets/%s/objects/%s", bucketName, objectName)
	obj := proto.Clone(req.GetObject()).(*pb.Object)
	obj.Bucket = PtrTo(bucketName)
	obj.Name = PtrTo(objectName)
	obj.Kind = PtrTo("storage#object")

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}
