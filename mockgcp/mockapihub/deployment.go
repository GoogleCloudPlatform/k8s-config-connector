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

package mockapihub

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "cloud.google.com/go/apihub/apiv1/apihubpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
)

type ApiHubV1Service struct {
	*MockService
	pb.UnimplementedApiHubServer
}

func (s *ApiHubV1Service) GetDeployment(ctx context.Context, req *pb.GetDeploymentRequest) (*pb.Deployment, error) {
	name, err := s.parseDeploymentName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Deployment{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *ApiHubV1Service) CreateDeployment(ctx context.Context, req *pb.CreateDeploymentRequest) (*pb.Deployment, error) {
	reqName := req.Parent + "/deployments/" + req.DeploymentId
	name, err := s.parseDeploymentName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.Deployment).(*pb.Deployment)
	obj.Name = fqn

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *ApiHubV1Service) UpdateDeployment(ctx context.Context, req *pb.UpdateDeploymentRequest) (*pb.Deployment, error) {
	name, err := s.parseDeploymentName(req.GetDeployment().GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.Deployment{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	paths := req.GetUpdateMask().GetPaths()
	for _, path := range paths {
		switch path {
		case "displayName":
			obj.DisplayName = req.GetDeployment().GetDisplayName()
		case "description":
			obj.Description = req.GetDeployment().GetDescription()
		case "documentation":
			obj.Documentation = req.GetDeployment().GetDocumentation()
		case "deploymentType":
			obj.DeploymentType = req.GetDeployment().GetDeploymentType()
		case "resourceUri":
			obj.ResourceUri = req.GetDeployment().GetResourceUri()
		case "endpoints":
			obj.Endpoints = req.GetDeployment().GetEndpoints()
		case "slo":
			obj.Slo = req.GetDeployment().GetSlo()
		case "environment":
			obj.Environment = req.GetDeployment().GetEnvironment()
		case "attributes":
			obj.Attributes = req.GetDeployment().GetAttributes()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
		}
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *ApiHubV1Service) DeleteDeployment(ctx context.Context, req *pb.DeleteDeploymentRequest) (*emptypb.Empty, error) {
	name, err := s.parseDeploymentName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.Deployment{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

type deploymentName struct {
	Project        *projects.ProjectData
	Location       string
	DeploymentName string
}

func (n *deploymentName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/deployments/%s", n.Project.ID, n.Location, n.DeploymentName)
}

// parseDeploymentName parses a string into a deploymentName.
// The expected form is `projects/*/locations/*/deployments/*`.
func (s *ApiHubV1Service) parseDeploymentName(name string) (*deploymentName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "deployments" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &deploymentName{
			Project:        project,
			Location:       tokens[3],
			DeploymentName: tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
