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
	"net/http"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/httpmux"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	gw_v1 "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/dataform/v1"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

type TeamFolderV1 struct {
	*MockService
	gw_v1.UnimplementedDataformServer
}

func (r *TeamFolderV1) GetTeamFolder(ctx context.Context, request *gw_v1.GetTeamFolderRequest) (*gw_v1.TeamFolder, error) {
	name, err := r.parseDataformTeamFolder(request.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &gw_v1.TeamFolder{}
	if err := r.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (r *TeamFolderV1) CreateTeamFolder(ctx context.Context, request *gw_v1.CreateTeamFolderRequest) (*gw_v1.TeamFolder, error) {
	name, err := r.parseDataformTeamFolder(request.TeamFolder.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(request.TeamFolder).(*gw_v1.TeamFolder)
	obj.Name = fqn

	if err := r.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (r *TeamFolderV1) UpdateTeamFolder(ctx context.Context, request *gw_v1.UpdateTeamFolderRequest) (*gw_v1.TeamFolder, error) {
	name, err := r.parseDataformTeamFolder(request.TeamFolder.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &gw_v1.TeamFolder{}
	if err := r.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	updateMask := request.GetUpdateMask()
	for _, path := range updateMask.Paths {
		switch path {
		case "display_name", "displayName":
			obj.DisplayName = request.GetTeamFolder().GetDisplayName()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "field %q is not yet handled in mock", path)
		}
	}

	if err := r.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (r *TeamFolderV1) DeleteTeamFolder(ctx context.Context, request *gw_v1.DeleteTeamFolderRequest) (*empty.Empty, error) {
	name, err := r.parseDataformTeamFolder(request.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &gw_v1.TeamFolder{}
	if err := r.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	httpmux.SetStatusCode(ctx, http.StatusNoContent)

	return &empty.Empty{}, nil
}

type teamFolderName struct {
	Project      *projects.ProjectData
	Location     string
	TeamFolderID string
}

func (n *teamFolderName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/teamFolders/" + n.TeamFolderID
}

func (s *MockService) parseDataformTeamFolder(name string) (*teamFolderName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "teamFolders" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &teamFolderName{
			Project:      project,
			Location:     tokens[3],
			TeamFolderID: tokens[5],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
