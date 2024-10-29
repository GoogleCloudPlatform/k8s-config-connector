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

package mocksecuresourcemanager

import (
	"context"
	"fmt"
	"strings"
	"time"

	longrunning "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/securesourcemanager/v1"
)

func (s *secureSourceManagerServer) GetRepository(ctx context.Context, req *pb.GetRepositoryRequest) (*pb.Repository, error) {
	name, err := s.parseRepositoryName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Repository{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *secureSourceManagerServer) CreateRepository(ctx context.Context, req *pb.CreateRepositoryRequest) (*longrunning.Operation, error) {
	reqName := req.Parent + "/repositories/" + req.RepositoryId
	name, err := s.parseRepositoryName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	obj := proto.Clone(req.Repository).(*pb.Repository)
	obj.Name = fqn

	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)

	instanceName, err := s.parseInstanceName(req.GetRepository().GetInstance())
	if err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("%s-%d", instanceName.InstanceID, name.Project.Number)
	domain := "." + name.Location + ".sourcemanager.dev"
	obj.Uris = &pb.Repository_URIs{
		Html:     prefix + domain + fmt.Sprintf("%s/%s", name.Project.ID, req.GetRepositoryId()),
		Api:      prefix + "-api" + domain + fmt.Sprintf("/v1/projects/%s/locations/%s/repositories/%s", name.Project.ID, name.Location, req.GetRepositoryId()),
		GitHttps: prefix + "-git" + domain + fmt.Sprintf("%s/%s.git", name.Project.ID, req.GetRepositoryId()),
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		Target:     name.String(),
		Verb:       "create",
		ApiVersion: "v1",
	}
	opPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, opPrefix, op, func() (proto.Message, error) {
		op.EndTime = timestamppb.Now()
		return obj, nil
	})
}

func (s *secureSourceManagerServer) DeleteRepository(ctx context.Context, req *pb.DeleteRepositoryRequest) (*longrunning.Operation, error) {
	name, err := s.parseRepositoryName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	deleted := &pb.Repository{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	op := &pb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		Target:     name.String(),
		Verb:       "delete",
		ApiVersion: "v1",
	}
	opPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, opPrefix, op, func() (proto.Message, error) {
		op.EndTime = timestamppb.Now()
		return &emptypb.Empty{}, nil
	})
}

type RepositoryName struct {
	Project      *projects.ProjectData
	Location     string
	RepositoryID string
}

func (n *RepositoryName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/repositories/%s", n.Project.ID, n.Location, n.RepositoryID)
}

// func (n *RepositoryName) Target() string {
// 	return fmt.Sprintf("projects/%s/locations/%s/repositories/%s", n.Project.ID, n.Location, n.RepositoryID)
// }

// parseRepositoryName parses a string into a RepositoryName.
// The expected form is projects/*/locations/*/repositories/*
func (s *MockService) parseRepositoryName(name string) (*RepositoryName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "repositories" {
		projectName, err := projects.ParseProjectName(tokens[0] + "/" + tokens[1])
		if err != nil {
			return nil, err
		}
		project, err := s.Projects.GetProject(projectName)
		if err != nil {
			return nil, err
		}

		name := &RepositoryName{
			Project:      project,
			Location:     tokens[3],
			RepositoryID: tokens[5],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
