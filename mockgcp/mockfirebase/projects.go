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

package mockfirebase

import (
	"context"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/firebase/v1beta1"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type projectsServer struct {
	*MockService
	pb.UnimplementedProjectsServerServer
}

type projectName struct {
	Project *projects.ProjectData
}

func (n *projectName) String() string {
	return "projects/" + n.Project.ID
}

func (s *MockService) parseProjectName(name string) (*projectName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 2 && tokens[0] == "projects" {
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}
		return &projectName{
			Project: project,
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "invalid project name %q", name)
}

func (s *projectsServer) defaultFirebaseProject(project *projects.ProjectData) *pb.FirebaseProject {
	fqn := "projects/" + project.ID
	return &pb.FirebaseProject{
		Name:          PtrTo(fqn),
		ProjectId:     PtrTo(project.ID),
		ProjectNumber: PtrTo(project.Number),
		DisplayName:   PtrTo(project.ID),
		State:         PtrTo("ACTIVE"),
		Etag:          PtrTo("abcdef0123A="),
		Resources: &pb.DefaultResources{
			HostingSite: PtrTo(project.ID),
		},
	}
}

func (s *projectsServer) GetProject(ctx context.Context, req *pb.GetProjectRequest) (*pb.FirebaseProject, error) {
	name, err := s.parseProjectName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.FirebaseProject{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			obj = s.defaultFirebaseProject(name.Project)
			if err := s.storage.Create(ctx, fqn, obj); err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	return obj, nil
}

func (s *projectsServer) AddFirebaseProject(ctx context.Context, req *pb.AddFirebaseProjectRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseProjectName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := s.defaultFirebaseProject(name.Project)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return s.operations.DoneLRO(ctx, fqn, nil, obj)
}

func (s *projectsServer) PatchProject(ctx context.Context, req *pb.PatchProjectRequest) (*pb.FirebaseProject, error) {
	name, err := s.parseProjectName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.FirebaseProject{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			obj = s.defaultFirebaseProject(name.Project)
		} else {
			return nil, err
		}
	}

	if req.Project != nil {
		if req.Project.DisplayName != nil {
			obj.DisplayName = req.Project.DisplayName
		}
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func PtrTo[T any](t T) *T {
	return &t
}
