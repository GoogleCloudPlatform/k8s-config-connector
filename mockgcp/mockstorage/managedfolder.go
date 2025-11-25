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
// proto.service: google.storage.control.v2.StorageControl
// proto.message: google.storage.control.v2.ManagedFolder

package mockstorage

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"k8s.io/klog/v2"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/httpmux"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/storage/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

type managedFolders struct {
	*MockService
	pb.UnimplementedManagedFoldersServerServer
}

func (s *managedFolders) GetManagedFolder(ctx context.Context, req *pb.GetManagedFolderRequest) (*pb.ManagedFolder, error) {
	name, err := s.buildManagedFolderName(req.GetBucket(), req.GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.ManagedFolder{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	httpmux.SetExpiresHeader(ctx, time.Now())

	return obj, nil
}

func (s *managedFolders) ListManagedFolders(ctx context.Context, req *pb.ListManagedFoldersRequest) (*pb.ManagedFolders, error) {
	findPrefix := req.GetPrefix()

	response := &pb.ManagedFolders{
		Kind: PtrTo("storage#managedFolders"),
	}

	findKind := (&pb.ManagedFolder{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, findKind, storage.ListOptions{Prefix: findPrefix}, func(obj proto.Message) error {
		managedFolder := obj.(*pb.ManagedFolder)
		response.Items = append(response.Items, managedFolder)
		return nil
	}); err != nil {
		return nil, err
	}

	httpmux.SetExpiresHeader(ctx, time.Now())
	return response, nil
}

func (s *managedFolders) getBucket(ctx context.Context, bucketName string) (*pb.Bucket, error) {
	bucketObj := &pb.Bucket{}
	if err := s.storage.Get(ctx, "buckets/"+bucketName, bucketObj); err != nil {
		return nil, err
	}
	return bucketObj, nil
}

func (s *managedFolders) InsertManagedFolder(ctx context.Context, req *pb.InsertManagedFolderRequest) (*pb.ManagedFolder, error) {
	name, err := s.buildManagedFolderName(req.GetBucket(), req.GetManagedFolder().GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	now := time.Now()

	bucket, err := s.getBucket(ctx, name.Bucket)
	if err != nil {
		return nil, err
	}

	if !bucket.GetIamConfiguration().GetUniformBucketLevelAccess().GetEnabled() {
		return nil, status.Errorf(codes.FailedPrecondition, "Uniform bucket-level access is required to be enabled on the bucket in order to perform this operation. Read more at https://cloud.google.com/storage/docs/uniform-bucket-level-access")
	}

	generation := int64(1)
	obj := proto.Clone(req.GetManagedFolder()).(*pb.ManagedFolder)
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.Metageneration = &generation
	obj.Id = PtrTo(fmt.Sprintf("%s/%s/", name.Bucket, name.Folder))
	obj.Kind = PtrTo("storage#managedFolder")
	obj.SelfLink = PtrTo(fmt.Sprintf("https://www.googleapis.com/storage/v1/b/%s/managedFolders/%s", name.Bucket, url.PathEscape(name.Folder+"/")))
	obj.Name = PtrTo(name.Folder + "/")

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *managedFolders) DeleteManagedFolder(ctx context.Context, req *pb.DeleteManagedFolderRequest) (*emptypb.Empty, error) {
	name, err := s.buildManagedFolderName(req.GetBucket(), req.GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	deleted := &pb.ManagedFolder{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	httpmux.SetStatusCode(ctx, http.StatusNoContent)

	return &emptypb.Empty{}, nil
}

type managedFolderName struct {
	Bucket string
	Folder string
}

func (n *managedFolderName) String() string {
	return fmt.Sprintf("buckets/%s/managedFolders/%s", n.Bucket, n.Folder)
}

// parseFolderName parses a string into an folderName.
// The expected form is `buckets/*/folders/*`.
func (s *MockService) buildManagedFolderName(bucket string, folder string) (*managedFolderName, error) {
	klog.Infof("folder before trim: %q", folder)
	folder = strings.TrimSuffix(folder, "/")

	klog.Infof("folder after trim: %q", folder)
	name := &managedFolderName{
		Bucket: bucket,
		Folder: folder,
	}

	return name, nil
}
