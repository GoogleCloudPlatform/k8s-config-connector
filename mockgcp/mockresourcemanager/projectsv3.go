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
	"hash/adler32"
	"strconv"
	"time"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	apierrors "k8s.io/apimachinery/pkg/api/errors"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/resourcemanager/v3"
)

type ProjectsV3 struct {
	*MockService
	pb.UnimplementedProjectsServer
}

func (s *ProjectsV3) GetProject(ctx context.Context, req *pb.GetProjectRequest) (*pb.Project, error) {
	projectName, err := projects.ParseProjectName(req.Name)
	if err != nil {
		return nil, err
	}

	project, err := s.projectsInternal.tryGetProject(ctx, projectName)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error reading project: %v", err)
	}
	if project == nil {
		// This API actually returns a 403 in the project-not-found case, unlike other APIs
		msg := fmt.Sprintf("Permission 'resourcemanager.projects.get' denied on resource '//cloudresourcemanager.googleapis.com/%s' (or it may not exist).", projectName.String())
		return nil, status.Error(codes.PermissionDenied, msg)
	}
	return project, nil
}

func (s *ProjectsV3) CreateProject(ctx context.Context, req *pb.CreateProjectRequest) (*longrunningpb.Operation, error) {
	projectID := req.GetProject().GetProjectId()
	if projectID == "" {
		return nil, status.Errorf(codes.InvalidArgument, "projectId is required")
	}

	hasher := adler32.New()
	hasher.Write([]byte(projectID))
	projectNumber := int64(hasher.Sum32()) // TODO: Check project number is unique? (and maybe require projects to be created)

	project := proto.Clone(req.GetProject()).(*pb.Project)
	project.Name = "projects/" + strconv.FormatInt(projectNumber, 10)
	project.ProjectId = projectID
	project.DisplayName = req.GetProject().GetDisplayName()
	project.State = pb.Project_ACTIVE
	project.CreateTime = timestamppb.Now()
	project.DeleteTime = nil
	project.UpdateTime = nil

	// TODO: What should the etag be?  We don't want to expose internal details (though it is only a mock),
	// but we want uniqueness
	project.Etag = fmt.Sprintf("%d", time.Now().UnixNano())

	fqn := "projects/" + project.ProjectId

	if err := s.storage.Create(ctx, fqn, project); err != nil {
		if apierrors.IsAlreadyExists(err) {
			return nil, status.Errorf(codes.AlreadyExists, "Requested entity already exists")
		}
		return nil, status.Errorf(codes.Internal, "error creating project: %v", err)
	}

	lro, err := s.operations.NewLRO(ctx)
	if err != nil {
		return nil, err
	}
	any, err := anypb.New(project)
	if err != nil {
		return nil, err
	}
	response := &longrunningpb.Operation_Response{}
	response.Response = any

	lro.Done = true
	lro.Result = response
	return lro, nil
}

func (s *ProjectsV3) DeleteProject(ctx context.Context, req *pb.DeleteProjectRequest) (*longrunningpb.Operation, error) {
	projectName, err := projects.ParseProjectName(req.Name)
	if err != nil {
		return nil, err
	}

	project, err := s.projectsInternal.tryGetProject(ctx, projectName)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error reading project: %v", err)
	}
	if project == nil {
		// This API actually returns a 403 in the project-not-found case, unlike other APIs
		msg := fmt.Sprintf("Permission 'resourcemanager.projects.get' denied on resource '//cloudresourcemanager.googleapis.com/%s' (or it may not exist).", projectName.String())
		return nil, status.Error(codes.PermissionDenied, msg)
	}

	fqn := "projects/" + project.ProjectId

	kind := (&pb.Project{}).ProtoReflect().Descriptor()
	if err := s.storage.Delete(ctx, kind, fqn); err != nil {
		return nil, status.Errorf(codes.Internal, "error deleting project: %v", err)
	}

	lro, err := s.operations.NewLRO(ctx)
	if err != nil {
		return nil, err
	}
	lro.Done = true

	return lro, nil
}
