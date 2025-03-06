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
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/storage/control/v2"
)

func (s *MockService) GetManagedFolder(ctx context.Context, req *pb.GetManagedFolderRequest) (*pb.ManagedFolder, error) {
	name, err := s.parseManagedFolderName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.ManagedFolder{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *MockService) CreateManagedFolder(ctx context.Context, req *pb.CreateManagedFolderRequest) (*pb.ManagedFolder, error) {
	reqName := req.Parent + "/managedFolders/" + req.ManagedFolderId
	name, err := s.parseManagedFolderName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	obj := proto.Clone(req.ManagedFolder).(*pb.ManagedFolder)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.Metageneration = 1

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *MockService) DeleteManagedFolder(ctx context.Context, req *pb.DeleteManagedFolderRequest) (*emptypb.Empty, error) {
	name, err := s.parseManagedFolderName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.ManagedFolder{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

type managedFolderName struct {
	Project       *projects.ProjectData
	Bucket        string
	ManagedFolder string
}

func (n *managedFolderName) String() string {
	return "projects/" + n.Project.ID + "/buckets/" + n.Bucket + "/managedFolders/" + n.ManagedFolder
}

// parseManagedFolderName parses a string into an managedFolderName .
// The expected form is `projects/*/buckets/*/managedFolders/*`.
func (s *MockService) parseManagedFolderName(name string) (*managedFolderName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "buckets" && tokens[4] == "managedFolders" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &managedFolderName{
			Project:       project,
			Bucket:        tokens[3],
			ManagedFolder: tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

```


