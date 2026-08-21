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

package mockconfig

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/config/v1"
)

type ConfigV1 struct {
	*MockService
	pb.UnimplementedConfigServer
}

type deploymentGroupName struct {
	Project             *projects.ProjectData
	Location            string
	DeploymentGroupName string
}

func (n *deploymentGroupName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/deploymentGroups/" + n.DeploymentGroupName
}

func (s *MockService) parseDeploymentGroupName(name string) (*deploymentGroupName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "deploymentGroups" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		return &deploymentGroupName{
			Project:             project,
			Location:            tokens[3],
			DeploymentGroupName: tokens[5],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

func (s *ConfigV1) GetDeploymentGroup(ctx context.Context, req *pb.GetDeploymentGroupRequest) (*pb.DeploymentGroup, error) {
	name, err := s.parseDeploymentGroupName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.DeploymentGroup{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *ConfigV1) CreateDeploymentGroup(ctx context.Context, req *pb.CreateDeploymentGroupRequest) (*longrunning.Operation, error) {
	reqName := req.Parent + "/deploymentGroups/" + req.DeploymentGroupId
	name, err := s.parseDeploymentGroupName(reqName)
	if err != nil {
		return nil, err
	}

	now := time.Now()

	fqn := name.String()

	obj := proto.Clone(req.DeploymentGroup).(*pb.DeploymentGroup)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.State = pb.DeploymentGroup_ACTIVE

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	opMetadata := &pb.OperationMetadata{
		ApiVersion:            "v1",
		CreateTime:            timestamppb.New(now),
		Verb:                  "create",
		RequestedCancellation: false,
		Target:                fqn,
	}
	opPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, opPrefix, opMetadata, func() (proto.Message, error) {
		opMetadata.EndTime = timestamppb.Now()
		lroResponse := proto.Clone(obj).(*pb.DeploymentGroup)
		lroResponse.Labels = nil
		lroResponse.Annotations = nil
		return lroResponse, nil
	})
}

func (s *ConfigV1) UpdateDeploymentGroup(ctx context.Context, req *pb.UpdateDeploymentGroupRequest) (*longrunning.Operation, error) {
	reqName := req.GetDeploymentGroup().GetName()

	name, err := s.parseDeploymentGroupName(reqName)
	if err != nil {
		return nil, err
	}

	now := time.Now()

	fqn := name.String()
	obj := &pb.DeploymentGroup{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	paths := req.GetUpdateMask().GetPaths()
	for _, path := range paths {
		switch path {
		case "labels":
			obj.Labels = req.GetDeploymentGroup().GetLabels()
		case "annotations":
			obj.Annotations = req.GetDeploymentGroup().GetAnnotations()
		case "deploymentUnits":
			obj.DeploymentUnits = req.GetDeploymentGroup().GetDeploymentUnits()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
		}
	}

	obj.UpdateTime = timestamppb.New(now)

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	opMetadata := &pb.OperationMetadata{
		ApiVersion:            "v1",
		CreateTime:            timestamppb.New(now),
		Verb:                  "update",
		RequestedCancellation: false,
		Target:                fqn,
	}
	opPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, opPrefix, opMetadata, func() (proto.Message, error) {
		opMetadata.EndTime = timestamppb.Now()
		lroResponse := proto.Clone(obj).(*pb.DeploymentGroup)
		lroResponse.Labels = nil
		lroResponse.Annotations = nil
		return lroResponse, nil
	})
}

func (s *ConfigV1) DeleteDeploymentGroup(ctx context.Context, req *pb.DeleteDeploymentGroupRequest) (*longrunning.Operation, error) {
	name, err := s.parseDeploymentGroupName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	oldObj := &pb.DeploymentGroup{}
	if err := s.storage.Delete(ctx, fqn, oldObj); err != nil {
		return nil, err
	}

	responseObj := proto.Clone(oldObj).(*pb.DeploymentGroup)
	responseObj.State = pb.DeploymentGroup_DELETED
	responseObj.CreateTime = nil
	responseObj.UpdateTime = nil
	responseObj.Labels = nil
	responseObj.Annotations = nil

	opMetadata := &pb.OperationMetadata{
		ApiVersion:            "v1",
		CreateTime:            timestamppb.New(now),
		Verb:                  "delete",
		RequestedCancellation: false,
		Target:                fqn,
	}
	opPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, opPrefix, opMetadata, func() (proto.Message, error) {
		opMetadata.EndTime = timestamppb.Now()
		return responseObj, nil
	})
}
