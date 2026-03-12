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
	"fmt"
	"strconv"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/resourcemanager/v3"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

type ProjectsInternal struct {
	storage storage.Storage
}

var _ projects.ProjectStore = &ProjectsInternal{}

func projectNotFoundError(project string) error {
	// This error follows a very specific format
	// For privacy reasons we don't want to reveal if the project exists.
	// Terraform also string-matches against the error(!!!)

	msg := fmt.Sprintf("Project '%s' not found or permission denied.", project)

	return status.Error(codes.PermissionDenied, msg)
}

// GetProjectByNumber returns the project with the specified project number, or an error if not found.
// Note that the project number must still be passed as a string, to keep terraform happy.
func (s *ProjectsInternal) GetProjectByNumber(projectNumberAsString string) (*projects.ProjectData, error) {
	ctx := context.TODO()

	projectNumber, err := strconv.ParseInt(projectNumberAsString, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid project number %q", projectNumberAsString)
	}

	project, err := s.tryGetProjectByNumber(ctx, projectNumber)
	if err != nil {
		return nil, err
	}

	if project == nil {
		// Terraform passes the project ID as 0000000 and expects that back in the error, not 0 (!!!)
		return nil, projectNotFoundError(projectNumberAsString)
	}
	return toProjectData(project)
}

// GetProjectByID returns the project with the specified project id, or an error if not found.
func (s *ProjectsInternal) GetProjectByID(projectID string) (*projects.ProjectData, error) {
	ctx := context.TODO()

	project, err := s.tryGetProjectByID(ctx, projectID)
	if err != nil {
		return nil, err
	}

	if project == nil {
		return nil, projectNotFoundError(projectID)
	}
	return toProjectData(project)
}

// GetProject returns the project for the parsed ProjectName.
func (s *ProjectsInternal) GetProject(project *projects.ProjectName) (*projects.ProjectData, error) {
	if project.ProjectID != "" {
		return s.GetProjectByID(project.ProjectID)
	} else {
		return s.GetProjectByNumber(project.OriginalValue)
	}
}

// GetProjectByIDOrNumber returns the project for the provided id or number.
func (s *ProjectsInternal) GetProjectByIDOrNumber(projectIDOrNumber string) (*projects.ProjectData, error) {
	projectName, err := projects.ParseProjectIDOrNumber(projectIDOrNumber)
	if err != nil {
		return nil, err
	}
	return s.GetProject(projectName)
}

func toProjectData(project *pb.Project) (*projects.ProjectData, error) {
	data := &projects.ProjectData{
		ID: project.ProjectId,
	}
	projectNumber, err := strconv.ParseInt(strings.TrimPrefix(project.Name, "projects/"), 10, 64)
	if err != nil {
		return nil, fmt.Errorf("cannot parse project number from %q", project.Name)
	}
	data.Number = projectNumber
	return data, nil
}

// tryGetProjectByID returns the project with the specified project id, or nil if not found.
func (s *ProjectsInternal) tryGetProjectByID(ctx context.Context, projectID string) (*pb.Project, error) {
	fqn := "projects/" + projectID

	obj := &pb.Project{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return obj, nil
}

// tryGetProjectByID returns the project with the specified project number, or nil if not found.
func (s *ProjectsInternal) tryGetProjectByNumber(ctx context.Context, projectNumber int64) (*pb.Project, error) {
	projectKind := (&pb.Project{}).ProtoReflect().Descriptor()

	matchProjectName := "projects/" + strconv.FormatInt(projectNumber, 10)

	var found *pb.Project
	// TODO: This is terribly inefficient!
	if err := s.storage.List(ctx, projectKind, storage.ListOptions{}, func(obj proto.Message) error {
		project := obj.(*pb.Project)
		if project.Name == matchProjectName {
			found = project
		}
		return nil
	}); err != nil {
		return nil, status.Errorf(codes.Internal, "error reading projects: %v", err)
	}

	return found, nil

}

// GetProject returns the project for the parsed ProjectName.
func (s *ProjectsInternal) tryGetProject(ctx context.Context, project *projects.ProjectName) (*pb.Project, error) {
	if project.ProjectID != "" {
		return s.tryGetProjectByID(ctx, project.ProjectID)
	} else {
		return s.tryGetProjectByNumber(ctx, project.ProjectNumber)
	}
}

// // updateProject updates the project.  If the project is not found, it returns permission denied.
// func (s *ProjectsInternal) updateProject(ctx context.Context, req *pb.UpdateProjectRequest) (*pb.Project, error) {
// 	projectName, err := projects.ParseProjectName(req.GetProject().GetName())
// 	if err != nil {
// 		return nil, err
// 	}

// 	obj, err := s.projectsInternal.tryGetProject(ctx, projectName)
// 	if err != nil {
// 		return nil, status.Errorf(codes.Internal, "error reading project: %v", err)
// 	}
// 	if obj == nil {
// 		return nil, status.Error(codes.PermissionDenied, "permission denied")
// 	}

// 	fqn := "projects/" + obj.ProjectID

// 	// Only the `display_name` and `labels` fields can be change.
// 	paths := req.GetUpdateMask().GetPaths()
// 	if len(paths) == 0 {
// 		if len(req.GetProject().GetLabels()) != 0 {
// 			paths = append(paths, "labels")
// 		}
// 		if len(req.GetProject().GetDisplayName()) != 0 {
// 			paths = append(paths, "display_name")
// 		}
// 	}

// 	// TODO: Some sort of helper for fieldmask?
// 	for _, path := range paths {
// 		switch path {
// 		case "display_name":
// 			obj.DisplayName = req.GetProject().DisplayName
// 		case "labels":
// 			obj.Labels = req.GetProject().GetLabels()
// 		default:
// 			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
// 		}
// 	}

// 	if err := s.storage.Update(ctx, fqn, obj); err != nil {
// 		return nil, status.Errorf(codes.Internal, "error updating project: %v", err)
// 	}

// 	return obj, nil
// }

// updateProject updates the project.  If the project is not found, it returns permission denied.
func (s *ProjectsInternal) mutateProject(ctx context.Context, name string, mutator func(project *pb.Project) error) (*pb.Project, error) {
	projectName, err := projects.ParseProjectName(name)
	if err != nil {
		return nil, err
	}

	obj, err := s.tryGetProject(ctx, projectName)
	if err != nil {
		return nil, err
	}
	if obj == nil {
		return nil, status.Error(codes.PermissionDenied, "permission denied")
	}

	fqn := "projects/" + obj.ProjectId

	if err := mutator(obj); err != nil {
		return nil, err
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}
