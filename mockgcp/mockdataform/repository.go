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

package mockdataform

import (
	"context"
	"net/http"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/httpmux"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/dataform/v1beta1"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

type RepositoryV1Beta1 struct {
	*MockService
	pb.UnimplementedDataformServer
}

func (r *RepositoryV1Beta1) ListRepositories(context.Context, *pb.ListRepositoriesRequest) (*pb.ListRepositoriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRepositories not implemented")
}

func (r *RepositoryV1Beta1) GetRepository(ctx context.Context, request *pb.GetRepositoryRequest) (*pb.Repository, error) {
	name, err := r.parseDataformRepository(request.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Repository{}
	if err := r.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (r *RepositoryV1Beta1) CreateRepository(ctx context.Context, request *pb.CreateRepositoryRequest) (*pb.Repository, error) {
	reqName := request.Parent + "/repositories/" + request.RepositoryId
	name, err := r.parseDataformRepository(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(request.Repository).(*pb.Repository)
	obj.Name = fqn

	if err := r.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}
func (r *RepositoryV1Beta1) UpdateRepository(ctx context.Context, request *pb.UpdateRepositoryRequest) (*pb.Repository, error) {
	name, err := r.parseDataformRepository(request.Repository.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Repository{}
	if err := r.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	updateMask := request.GetUpdateMask()
	for _, path := range updateMask.Paths {
		switch path {
		case "npmrc_environment_variables_secret_version", "npmrcEnvironmentVariablesSecretVersion":
			obj.NpmrcEnvironmentVariablesSecretVersion = request.GetRepository().GetNpmrcEnvironmentVariablesSecretVersion()
		case "git_remote_settings", "gitRemoteSettings":
			obj.GitRemoteSettings = request.GetRepository().GetGitRemoteSettings()
		case "workspace_compilation_overrides", "workspaceCompilationOverrides":
			obj.WorkspaceCompilationOverrides = request.GetRepository().GetWorkspaceCompilationOverrides()
		case "service_account", "serviceAccount":
			obj.ServiceAccount = request.GetRepository().GetServiceAccount()
		case "display_name", "displayName":
			obj.DisplayName = request.GetRepository().GetDisplayName()
		case "set_authenticated_user_admin", "setAuthenticatedUserAdmin":
			obj.SetAuthenticatedUserAdmin = request.GetRepository().GetSetAuthenticatedUserAdmin()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "field %q is not yet handled in mock", path)
		}
	}

	if err := r.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}
func (r *RepositoryV1Beta1) DeleteRepository(ctx context.Context, request *pb.DeleteRepositoryRequest) (*empty.Empty, error) {
	name, err := r.parseDataformRepository(request.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.Repository{}
	if err := r.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	httpmux.SetStatusCode(ctx, http.StatusNoContent)

	return &empty.Empty{}, nil
}

// 'projects/${projectId}/locations/us-west2/repositories/dataformrepository-${uniqueId}'
type repositoryName struct {
	Project      *projects.ProjectData
	Location     string
	RepositoryID string
}

func (n *repositoryName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/repositories/" + n.RepositoryID
}

// parseDataformRepository parses a string into a repositoryName.
// The expected form is projects/<projectID>/locations/<region>/repositories/<repositoryID>
func (s *MockService) parseDataformRepository(name string) (*repositoryName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "repositories" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &repositoryName{
			Project:      project,
			Location:     tokens[3],
			RepositoryID: tokens[5],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
