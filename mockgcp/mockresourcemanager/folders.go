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

package mockresourcemanager

import (
	"context"
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
	"k8s.io/klog/v2"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/resourcemanager/v3"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

type Folders struct {
	*MockService
	pb.UnimplementedFoldersServer
}

func (s *Folders) GetFolder(ctx context.Context, req *pb.GetFolderRequest) (*pb.Folder, error) {
	name, err := s.parseFolderName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Folder{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// We should verify that this is part of one of our projects, but ... it's a mock

	return obj, nil
}

func (s *Folders) SearchFolders(ctx context.Context, req *pb.SearchFoldersRequest) (*pb.SearchFoldersResponse, error) {
	log := klog.FromContext(ctx)

	// TODO: Implement search properly
	log.Info("SearchFolders is stub implemented", "request", req)

	response := &pb.SearchFoldersResponse{}
	return response, nil
}

func (s *Folders) CreateFolder(ctx context.Context, req *pb.CreateFolderRequest) (*longrunningpb.Operation, error) {
	parent := req.GetFolder().GetParent()
	if strings.HasPrefix(parent, "folders/") {
		_, err := s.parseFolderName(parent)
		if err != nil {
			return nil, err
		}
	} else if strings.HasPrefix(parent, "organizations/") {
		// We should check that the org exists, permissions etc, but ... it's a mock
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "parent %q is not valid", parent)
	}

	name := &folderName{
		ID: time.Now().UnixNano(),
	}

	fqn := name.String()
	now := timestamppb.Now()

	obj := proto.Clone(req.Folder).(*pb.Folder)

	obj.CreateTime = now
	obj.UpdateTime = now
	obj.Etag = base64.StdEncoding.EncodeToString(computeEtag(obj))
	obj.Name = fqn
	obj.State = pb.Folder_ACTIVE

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return s.operations.StartLRO(ctx, "", nil, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *Folders) UpdateFolder(ctx context.Context, req *pb.UpdateFolderRequest) (*longrunningpb.Operation, error) {
	reqName := req.GetFolder().GetName()
	name, err := s.parseFolderName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Folder{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// Only the `display_name` field can be changed.

	paths := req.GetUpdateMask().GetPaths()
	for _, path := range paths {
		switch path {
		case "display_name":
			obj.DisplayName = req.GetFolder().GetDisplayName()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not implemented by mockgcp", path)
		}
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	metadata := &pb.UpdateFolderMetadata{}
	return s.operations.StartLRO(ctx, "", metadata, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *Folders) MoveFolder(ctx context.Context, req *pb.MoveFolderRequest) (*longrunningpb.Operation, error) {
	reqName := req.GetName()
	name, err := s.parseFolderName(reqName)
	if err != nil {
		return nil, err
	}

	destinationParent := req.GetDestinationParent()

	fqn := name.String()
	obj := &pb.Folder{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	obj.Parent = destinationParent

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	metadata := &pb.MoveFolderMetadata{}
	return s.operations.StartLRO(ctx, "", metadata, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *Folders) DeleteFolder(ctx context.Context, req *pb.DeleteFolderRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseFolderName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Folder{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	obj.State = pb.Folder_DELETE_REQUESTED
	obj.UpdateTime = timestamppb.Now()

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// We should verify that this is part of one of our projects, but ... it's a mock
	lro, err := s.operations.DoneLRO(ctx, "", nil, obj)
	if err != nil {
		return nil, err
	}
	// Does not return name for finished LROs
	lro.Name = ""
	return lro, nil
}

type folderName struct {
	ID int64
}

func (n *folderName) String() string {
	return fmt.Sprintf("folders/%d", n.ID)
}

// parseFolderName parses a string into a folderName.
// The expected form is folders/<folderName>
func (s *MockService) parseFolderName(name string) (*folderName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 2 && tokens[0] == "folders" {

		n, err := strconv.ParseInt(tokens[1], 10, 64)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid (bad id)", name)
		}
		name := &folderName{
			ID: n,
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
