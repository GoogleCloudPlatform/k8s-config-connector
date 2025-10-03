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
	"time"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/resourcemanager/v1"
	v3 "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/resourcemanager/v3"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"k8s.io/klog/v2"
)

type ProjectsV1 struct {
	*MockService
	pb.UnimplementedProjectsServerServer
}

// Retrieves the project identified by the specified `name` (for example,
// `projects/415104041262`).
func (s *ProjectsV1) GetProject(ctx context.Context, req *pb.GetProjectRequest) (*pb.Project, error) {
	reqV3 := &v3.GetProjectRequest{
		Name: "projects/" + req.GetProjectId(),
	}

	responseV3, err := s.projectsV3.GetProject(ctx, reqV3)
	if err != nil {
		// Terraform string-matches against the error message (!!!)
		if status.Code(err) == codes.PermissionDenied {
			// This API actually returns a 403 in the project-not-found case, unlike other APIs
			return nil, status.Error(codes.PermissionDenied, "The caller does not have permission")
		}

		klog.Infof("error is %T %+v", err, err)
		return nil, err
	}

	responseV1 := &pb.Project{}
	if err := projectToV1(responseV3, responseV1); err != nil {
		return nil, err
	}
	return responseV1, nil
}

// Request that a new project be created.
func (s *ProjectsV1) CreateProject(ctx context.Context, req *pb.CreateProjectRequest) (*longrunningpb.Operation, error) {
	reqV3 := &v3.CreateProjectRequest{
		Project: &v3.Project{
			ProjectId:   req.GetProject().GetProjectId(),
			DisplayName: req.GetProject().GetName(),
			Labels:      req.GetProject().GetLabels(),
		},
	}

	if req.Project.Parent != nil {
		plural := req.Project.Parent.GetType() + "s" // A bit of a hack!
		reqV3.Project.Parent = plural + "/" + req.Project.Parent.GetId()
	}

	lro, err := s.projectsV3.CreateProject(ctx, reqV3)
	if err != nil {
		return nil, err
	}

	lrov1, err := lroV3ToV1(lro)
	if err != nil {
		return nil, err
	}

	// We actually only return the name from this operation
	return &longrunningpb.Operation{
		Name: lrov1.Name,
	}, nil
}

// Request that a new project be created.
func (s *ProjectsV1) DeleteProject(ctx context.Context, req *pb.DeleteProjectRequest) (*pb.Empty, error) {
	reqV3 := &v3.DeleteProjectRequest{
		Name: "projects/" + req.GetProjectId(),
	}

	op, err := s.projectsV3.DeleteProject(ctx, reqV3)
	if err != nil {
		return nil, err
	}

	// V1 does not return an LRO (this method is actually fast anyway, we just mark the projet for deletion)
	if _, err := s.operations.Wait(ctx, op.Name, time.Minute); err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}

// Updates a project.
// Only the V3 `display_name` and `labels` fields can be change.
func (s *ProjectsV1) UpdateProject(ctx context.Context, req *pb.UpdateProjectRequest) (*pb.Project, error) {
	reqV3 := &v3.UpdateProjectRequest{
		Project: &v3.Project{
			Name:        "projects/" + req.GetProject().GetProjectId(),
			DisplayName: req.GetProject().GetName(),
			Labels:      req.GetProject().GetLabels(),
		},
	}

	lro, err := s.projectsV3.UpdateProject(ctx, reqV3)
	if err != nil {
		return nil, err
	}
	if !lro.Done {
		return nil, fmt.Errorf("expected updateProject to be immediate")
	}

	// TODO: Get object from lro
	project, err := s.GetProject(ctx, &pb.GetProjectRequest{ProjectId: PtrTo(req.GetProject().GetProjectId())})
	if err != nil {
		return nil, fmt.Errorf("error fetching project after update: %w", err)
	}
	return project, nil
	// updatedV3, err := s.projectsInternal.updateProject(ctx, reqV3)
	// if err != nil {
	// 	return nil, err
	// }

	// updatedV1 := &pb.Project{}
	// if err := projectToV1(updatedV3, updatedV1); err != nil {
	// 	return nil, err
	// }

	// return updatedV1, nil
}

// Convert a V3 LRO (with an embedded V3 result) to a V1 LRO
func lroV3ToV1(lro *longrunningpb.Operation) (*longrunningpb.Operation, error) {
	if response, ok := lro.Result.(*longrunningpb.Operation_Response); ok {
		projectV3 := &v3.Project{}
		if err := anypb.UnmarshalTo(response.Response, projectV3, proto.UnmarshalOptions{}); err != nil {
			return nil, err
		}
		projectV1 := &pb.Project{}
		if err := projectToV1(projectV3, projectV1); err != nil {
			return nil, err
		}
		anyV1, err := anypb.New(projectV1)
		if err != nil {
			return nil, err
		}
		response.Response = anyV1
	}

	return lro, nil
}

// Convert a V3 project to a V1 project
func projectToV1(in *v3.Project, out *pb.Project) error {
	projectNumber, err := strconv.ParseInt(strings.TrimPrefix(in.Name, "projects/"), 10, 64)
	if err != nil {
		return fmt.Errorf("cannot parse project number from %q", in.Name)
	}
	out.ProjectNumber = &projectNumber
	out.ProjectId = &in.ProjectId
	if in.DisplayName != "" {
		out.Name = &in.DisplayName
	}
	out.CreateTime = in.CreateTime
	out.Labels = in.Labels

	switch in.State {
	case v3.Project_ACTIVE:
		out.LifecycleState = PtrTo("ACTIVE")
	case v3.Project_DELETE_REQUESTED:
		out.LifecycleState = PtrTo("DELETE_REQUESTED")
	default:
		out.LifecycleState = nil
	}

	parent := in.GetParent()
	if strings.HasPrefix(parent, "organizations/") {
		out.Parent = &pb.ResourceId{
			Type: PtrTo("organization"),
			Id:   PtrTo(strings.TrimPrefix(parent, "organizations/")),
		}
	} else if strings.HasPrefix(parent, "folders/") {
		out.Parent = &pb.ResourceId{
			Type: PtrTo("folder"),
			Id:   PtrTo(strings.TrimPrefix(parent, "folders/")),
		}
	}
	return nil
}
