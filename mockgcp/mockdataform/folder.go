// Copyright 2026 Google LLC
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

package mockdataform

import (
	"context"
	"strings"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/dataform/v1beta1"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (r *RepositoryV1Beta1) GetFolder(ctx context.Context, request *pb.GetFolderRequest) (*pb.Folder, error) {
	name, err := r.parseDataformFolder(request.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Folder{}
	if err := r.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (r *RepositoryV1Beta1) CreateFolder(ctx context.Context, request *pb.CreateFolderRequest) (*pb.Folder, error) {
	reqName := request.Parent + "/folders/" + request.FolderId
	name, err := r.parseDataformFolder(reqName)
	if err != nil {
		return nil, err
	}

	// Real GCP assigns a server-generated UUID for the folder name/ID.
	// To match this behavior and ensure stable E2E tests, we use a deterministic static UUID.
	name.FolderID = "2ba614be-5147-410a-ad89-bab4ff528497"
	fqn := name.String()

	obj := proto.Clone(request.Folder).(*pb.Folder)
	obj.Name = fqn
	obj.CreatorIamPrincipal = proto.String("user:overseer-kcc-tester@" + name.Project.ID + ".iam.gserviceaccount.com")
	now := time.Now()
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)

	if err := r.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (r *RepositoryV1Beta1) UpdateFolder(ctx context.Context, request *pb.UpdateFolderRequest) (*pb.Folder, error) {
	name, err := r.parseDataformFolder(request.Folder.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Folder{}
	if err := r.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	now := time.Now()
	obj.UpdateTime = timestamppb.New(now)

	updateMask := request.GetUpdateMask()
	if updateMask == nil || len(updateMask.Paths) == 0 {
		obj.DisplayName = request.GetFolder().GetDisplayName()
		if request.GetFolder().ContainingFolder != nil {
			obj.ContainingFolder = request.GetFolder().ContainingFolder
		}
	} else {
		for _, path := range updateMask.Paths {
			switch path {
			case "display_name", "displayName":
				obj.DisplayName = request.GetFolder().GetDisplayName()
			case "containing_folder", "containingFolder":
				obj.ContainingFolder = request.GetFolder().ContainingFolder
			default:
				return nil, status.Errorf(codes.InvalidArgument, "field %q is not yet handled in mock", path)
			}
		}
	}

	if err := r.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (r *RepositoryV1Beta1) DeleteFolder(ctx context.Context, request *pb.DeleteFolderRequest) (*empty.Empty, error) {
	name, err := r.parseDataformFolder(request.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.Folder{}
	if err := r.storage.Delete(ctx, fqn, deleted); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return &empty.Empty{}, nil
}

// 'projects/${projectId}/locations/us-central1/folders/${folderID}'
type folderName struct {
	Project  *projects.ProjectData
	Location string
	FolderID string
}

func (n *folderName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/folders/" + n.FolderID
}

// parseDataformFolder parses a string into a folderName.
// The expected form is projects/<projectID>/locations/<region>/folders/<folderID>
func (s *MockService) parseDataformFolder(name string) (*folderName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "folders" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &folderName{
			Project:  project,
			Location: tokens[3],
			FolderID: tokens[5],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
