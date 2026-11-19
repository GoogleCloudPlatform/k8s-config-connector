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

package mockces

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/ces/apiv1/cespb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/fields"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

func (s *AgentServiceV1) GetDeployment(ctx context.Context, req *pb.GetDeploymentRequest) (*pb.Deployment, error) {
	name, err := s.parseDeploymentName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Deployment{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "com.google.apps.framework.request.NotFoundException: Deployment %s not found.", fqn)
		}
		return nil, err
	}
	// Real GCP GetDeployment does not return etag
	result := proto.Clone(obj).(*pb.Deployment)
	result.Etag = ""
	return result, nil
}

func (s *AgentServiceV1) CreateDeployment(ctx context.Context, req *pb.CreateDeploymentRequest) (*pb.Deployment, error) {
	parent, err := s.parseAppName(req.GetParent())
	if err != nil {
		return nil, err
	}

	parentApp := &pb.App{}
	if err := s.storage.Get(ctx, parent.String(), parentApp); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Parent resource '%s' was not found", parent.String())
		}
		return nil, err
	}

	deploymentID := req.GetDeploymentId()
	if deploymentID == "" {
		return nil, status.Errorf(codes.InvalidArgument, "deploymentId is required")
	}

	name := &deploymentName{
		Project:      parent.Project,
		Location:     parent.Location,
		AppID:        parent.AppID,
		DeploymentID: deploymentID,
	}

	fqn := name.String()

	existing := &pb.Deployment{}
	if err := s.storage.Get(ctx, fqn, existing); err == nil {
		return nil, status.Errorf(codes.AlreadyExists, "Resource '%s' already exists", fqn)
	}

	obj := proto.Clone(req.GetDeployment()).(*pb.Deployment)
	obj.Name = fqn
	now := time.Now()
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// Real GCP CreateDeployment returns the deployment without etag
	result := proto.Clone(obj).(*pb.Deployment)
	result.Etag = ""
	return result, nil
}

func (s *AgentServiceV1) UpdateDeployment(ctx context.Context, req *pb.UpdateDeploymentRequest) (*pb.Deployment, error) {
	reqDeployment := req.GetDeployment()
	if reqDeployment == nil {
		return nil, status.Errorf(codes.InvalidArgument, "deployment is required")
	}

	name, err := s.parseDeploymentName(reqDeployment.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Deployment{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "com.google.apps.framework.request.NotFoundException: Deployment %s not found.", fqn)
		}
		return nil, err
	}

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		proto.Merge(obj, reqDeployment)
	} else {
		if err := fields.UpdateByFieldMask(obj, reqDeployment, paths); err != nil {
			return nil, err
		}
	}

	obj.UpdateTime = timestamppb.New(time.Now())

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// Real GCP UpdateDeployment returns etag
	result := proto.Clone(obj).(*pb.Deployment)
	result.Etag = "abcdef0123A="
	return result, nil
}

func (s *AgentServiceV1) DeleteDeployment(ctx context.Context, req *pb.DeleteDeploymentRequest) (*emptypb.Empty, error) {
	name, err := s.parseDeploymentName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.Deployment{}
	if err := s.storage.Get(ctx, fqn, deleted); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "com.google.apps.framework.request.NotFoundException: Deployment %s not found.", fqn)
		}
		return nil, err
	}

	if err := s.storage.Delete(ctx, fqn, &pb.Deployment{}); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *AgentServiceV1) ListDeployments(ctx context.Context, req *pb.ListDeploymentsRequest) (*pb.ListDeploymentsResponse, error) {
	parent, err := s.parseAppName(req.GetParent())
	if err != nil {
		return nil, err
	}

	response := &pb.ListDeploymentsResponse{}
	findPrefix := fmt.Sprintf("projects/%s/locations/%s/apps/%s/deployments/", parent.Project.ID, parent.Location, parent.AppID)

	deploymentKind := (&pb.Deployment{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, deploymentKind, storage.ListOptions{}, func(obj proto.Message) error {
		item := obj.(*pb.Deployment)
		if strings.HasPrefix(item.GetName(), findPrefix) {
			itemCopy := proto.Clone(item).(*pb.Deployment)
			itemCopy.Etag = ""
			response.Deployments = append(response.Deployments, itemCopy)
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return response, nil
}
