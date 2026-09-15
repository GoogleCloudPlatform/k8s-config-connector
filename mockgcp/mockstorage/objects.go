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
	"strings"
	"time"

	grpcpb "cloud.google.com/go/storage/control/apiv2/controlpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/httpmux"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/storage/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
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

	bucketName := req.GetBucket()
	prefix := req.GetPrefix()
	delimiter := req.GetDelimiter()

	var prefixes []string

	// List HTTP folders from s.storage
	folderKind := (&pb.Folder{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, folderKind, storage.ListOptions{}, func(obj proto.Message) error {
		folder := obj.(*pb.Folder)
		if folder.GetBucket() != bucketName {
			return nil
		}
		name := folder.GetName() // e.g. "testfolder/"
		if strings.HasPrefix(name, prefix) {
			rel := name[len(prefix):]
			if delimiter != "" {
				idx := strings.Index(rel, delimiter)
				if idx >= 0 {
					p := prefix + rel[:idx+len(delimiter)]
					prefixes = append(prefixes, p)
				}
			} else {
				prefixes = append(prefixes, name)
			}
		}
		return nil
	}); err != nil {
		return nil, err
	}

	// List HTTP managed folders from s.storage
	httpManagedFolderKind := (&pb.ManagedFolder{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, httpManagedFolderKind, storage.ListOptions{}, func(obj proto.Message) error {
		folder := obj.(*pb.ManagedFolder)
		if folder.GetBucket() != bucketName {
			return nil
		}
		name := folder.GetName() // e.g. "testmanagedfolder/"
		if strings.HasPrefix(name, prefix) {
			rel := name[len(prefix):]
			if delimiter != "" {
				idx := strings.Index(rel, delimiter)
				if idx >= 0 {
					p := prefix + rel[:idx+len(delimiter)]
					prefixes = append(prefixes, p)
				}
			} else {
				prefixes = append(prefixes, name)
			}
		}
		return nil
	}); err != nil {
		return nil, err
	}

	// List control folders from s.storage
	controlFolderKind := (&grpcpb.Folder{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, controlFolderKind, storage.ListOptions{}, func(obj proto.Message) error {
		folder := obj.(*grpcpb.Folder)
		bucketID, err := parseBucketIDFromControlName(folder.GetName())
		if err != nil || bucketID != bucketName {
			return nil
		}
		tokens := strings.Split(folder.GetName(), "/folders/")
		if len(tokens) < 2 {
			return nil
		}
		name := tokens[1]
		if !strings.HasSuffix(name, "/") {
			name += "/"
		}
		if strings.HasPrefix(name, prefix) {
			rel := name[len(prefix):]
			if delimiter != "" {
				idx := strings.Index(rel, delimiter)
				if idx >= 0 {
					p := prefix + rel[:idx+len(delimiter)]
					prefixes = append(prefixes, p)
				}
			} else {
				prefixes = append(prefixes, name)
			}
		}
		return nil
	}); err != nil {
		return nil, err
	}

	// List control managed folders from s.storage
	controlManagedFolderKind := (&grpcpb.ManagedFolder{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, controlManagedFolderKind, storage.ListOptions{}, func(obj proto.Message) error {
		folder := obj.(*grpcpb.ManagedFolder)
		bucketID, err := parseBucketIDFromControlName(folder.GetName())
		if err != nil || bucketID != bucketName {
			return nil
		}
		tokens := strings.Split(folder.GetName(), "/managedFolders/")
		if len(tokens) < 2 {
			return nil
		}
		name := tokens[1]
		if !strings.HasSuffix(name, "/") {
			name += "/"
		}
		if strings.HasPrefix(name, prefix) {
			rel := name[len(prefix):]
			if delimiter != "" {
				idx := strings.Index(rel, delimiter)
				if idx >= 0 {
					p := prefix + rel[:idx+len(delimiter)]
					prefixes = append(prefixes, p)
				}
			} else {
				prefixes = append(prefixes, name)
			}
		}
		return nil
	}); err != nil {
		return nil, err
	}

	uniquePrefixes := make(map[string]bool)
	var finalPrefixes []string
	for _, p := range prefixes {
		if !uniquePrefixes[p] {
			uniquePrefixes[p] = true
			finalPrefixes = append(finalPrefixes, p)
		}
	}

	ret := &pb.Objects{}
	ret.Kind = PtrTo("storage#objects")
	ret.Prefixes = finalPrefixes
	return ret, nil
}

func (s *objects) GetObject(ctx context.Context, req *pb.GetObjectRequest) (*pb.Object, error) {
	// A stub implementation, just to support deletion (for now)

	return nil, status.Errorf(codes.NotFound, "No such object: %s/%s", req.GetBucket(), req.GetName())
}
