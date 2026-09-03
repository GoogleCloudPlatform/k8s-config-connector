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
	"strings"

	pb "cloud.google.com/go/apihub/apiv1/apihubpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *ApiHubServer) GetDeployment(ctx context.Context, req *pb.GetDeploymentRequest) (*pb.Deployment, error) {
	name, err := s.parseDeploymentName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Deployment{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource %q was not found", fqn)
		}
		return nil, err
	}

	s.enrichDeployment(obj)
	return obj, nil
}

func (s *ApiHubServer) CreateDeployment(ctx context.Context, req *pb.CreateDeploymentRequest) (*pb.Deployment, error) {
	reqName := req.Parent + "/deployments/" + req.DeploymentId
	name, err := s.parseDeploymentName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.Deployment).(*pb.Deployment)
	obj.Name = fqn

	now := timestamppb.Now()
	obj.CreateTime = now
	obj.UpdateTime = now

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	s.enrichDeployment(obj)
	return obj, nil
}

func (s *ApiHubServer) UpdateDeployment(ctx context.Context, req *pb.UpdateDeploymentRequest) (*pb.Deployment, error) {
	deploymentName := req.GetDeployment().GetName()

	name, err := s.parseDeploymentName(deploymentName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Deployment{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource %q was not found", fqn)
		}
		return nil, err
	}

	updateMask := req.GetUpdateMask()
	paths := updateMask.GetPaths()
	if len(paths) == 0 {
		// Treat as all paths if empty
		paths = []string{
			"display_name", "description", "documentation", "deployment_type", "resource_uri", "endpoints", "slo", "environment",
		}
	}

	for _, path := range paths {
		switch path {
		case "display_name", "displayName":
			obj.DisplayName = req.GetDeployment().GetDisplayName()
		case "description":
			obj.Description = req.GetDeployment().GetDescription()
		case "documentation":
			obj.Documentation = req.GetDeployment().GetDocumentation()
		case "deployment_type", "deploymentType":
			obj.DeploymentType = req.GetDeployment().GetDeploymentType()
		case "resource_uri", "resourceUri":
			obj.ResourceUri = req.GetDeployment().GetResourceUri()
		case "endpoints":
			obj.Endpoints = req.GetDeployment().GetEndpoints()
		case "slo":
			obj.Slo = req.GetDeployment().GetSlo()
		case "environment":
			obj.Environment = req.GetDeployment().GetEnvironment()
		}
	}

	obj.UpdateTime = timestamppb.Now()

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	s.enrichDeployment(obj)
	return obj, nil
}

func (s *ApiHubServer) DeleteDeployment(ctx context.Context, req *pb.DeleteDeploymentRequest) (*emptypb.Empty, error) {
	name, err := s.parseDeploymentName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Deployment{}
	if err := s.storage.Delete(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource %q was not found", fqn)
		}
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *ApiHubServer) ListDeployments(ctx context.Context, req *pb.ListDeploymentsRequest) (*pb.ListDeploymentsResponse, error) {
	name, err := s.parseDeploymentName(req.Parent + "/deployments/dummy")
	if err != nil {
		return nil, err
	}

	response := &pb.ListDeploymentsResponse{}

	findPrefix := "projects/" + name.Project.ID + "/locations/" + name.Location + "/deployments/"

	deploymentKind := (&pb.Deployment{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, deploymentKind, storage.ListOptions{}, func(obj proto.Message) error {
		deployment := obj.(*pb.Deployment)
		if strings.HasPrefix(deployment.GetName(), findPrefix) {
			s.enrichDeployment(deployment)
			response.Deployments = append(response.Deployments, deployment)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}

func (s *ApiHubServer) enrichDeployment(obj *pb.Deployment) {
	if obj == nil || obj.Name == "" {
		return
	}
	name, err := s.parseDeploymentName(obj.Name)
	if err != nil {
		return
	}
	projectID := name.Project.ID
	location := name.Location

	if obj.DeploymentType != nil {
		obj.DeploymentType.Attribute = "projects/" + projectID + "/locations/" + location + "/attributes/system-deployment-type"
		if ev := obj.DeploymentType.GetEnumValues(); ev != nil {
			for _, val := range ev.Values {
				if val.Id == "apigee" {
					val.DisplayName = "Apigee"
					val.Description = "Apigee"
					val.Immutable = true
				}
			}
		}
	}

	if obj.Environment != nil {
		obj.Environment.Attribute = "projects/" + projectID + "/locations/" + location + "/attributes/system-environment"
		if ev := obj.Environment.GetEnumValues(); ev != nil {
			for _, val := range ev.Values {
				if val.Id == "staging" {
					val.DisplayName = "Staging"
					val.Description = "Staging"
				} else if val.Id == "prod" {
					val.DisplayName = "Production"
					val.Description = "Production"
				}
			}
		}
	}

	if obj.Slo != nil {
		obj.Slo.Attribute = "projects/" + projectID + "/locations/" + location + "/attributes/system-slo"
		if ev := obj.Slo.GetEnumValues(); ev != nil {
			for _, val := range ev.Values {
				if val.Id == "99-9" {
					val.DisplayName = "99.9%"
					val.Description = "99.9% SLO"
				} else if val.Id == "99-99" {
					val.DisplayName = "99.99%"
					val.Description = "99.99% SLO"
				}
			}
		}
	}
}
