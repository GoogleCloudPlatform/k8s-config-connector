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

// +tool:mockgcp-support
// proto.service: google.cloud.networkmanagement.v1.VpcFlowLogsService
// proto.message: google.cloud.networkmanagement.v1.VpcFlowLogsConfig

package mocknetworkmanagement

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	pb "cloud.google.com/go/networkmanagement/apiv1/networkmanagementpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
)

type vpcFlowLogsService struct {
	*MockService
	pb.UnimplementedVpcFlowLogsServiceServer
}

func (s *vpcFlowLogsService) GetVpcFlowLogsConfig(ctx context.Context, req *pb.GetVpcFlowLogsConfigRequest) (*pb.VpcFlowLogsConfig, error) {
	name, err := s.parseVpcFlowLogsConfigName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.VpcFlowLogsConfig{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}
	return obj, nil
}

func (s *vpcFlowLogsService) CreateVpcFlowLogsConfig(ctx context.Context, req *pb.CreateVpcFlowLogsConfigRequest) (*longrunningpb.Operation, error) {
	reqName := req.Parent + "/vpcFlowLogsConfigs/" + req.VpcFlowLogsConfigId
	name, err := s.parseVpcFlowLogsConfigName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	obj := proto.CloneOf(req.VpcFlowLogsConfig)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)

	// Set target resource state
	targetResourceState := pb.VpcFlowLogsConfig_TARGET_RESOURCE_EXISTS
	obj.TargetResourceState = &targetResourceState

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return s.operations.StartLRO(ctx, fqn, nil, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *vpcFlowLogsService) UpdateVpcFlowLogsConfig(ctx context.Context, req *pb.UpdateVpcFlowLogsConfigRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseVpcFlowLogsConfigName(req.VpcFlowLogsConfig.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	actual := &pb.VpcFlowLogsConfig{}
	if err := s.storage.Get(ctx, fqn, actual); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	// Simple mock update using field masks
	paths := req.UpdateMask.GetPaths()
	if len(paths) == 0 {
		// If empty, update all writable fields
		paths = []string{"description", "state", "aggregation_interval", "flow_sampling", "metadata", "metadata_fields", "filter_expr", "labels"}
	}

	for _, path := range paths {
		switch path {
		case "description":
			actual.Description = req.VpcFlowLogsConfig.Description
		case "state":
			actual.State = req.VpcFlowLogsConfig.State
		case "aggregation_interval", "aggregationInterval":
			actual.AggregationInterval = req.VpcFlowLogsConfig.AggregationInterval
		case "flow_sampling", "flowSampling":
			actual.FlowSampling = req.VpcFlowLogsConfig.FlowSampling
		case "metadata":
			actual.Metadata = req.VpcFlowLogsConfig.Metadata
		case "metadata_fields", "metadataFields":
			actual.MetadataFields = req.VpcFlowLogsConfig.MetadataFields
		case "filter_expr", "filterExpr":
			actual.FilterExpr = req.VpcFlowLogsConfig.FilterExpr
		case "labels":
			actual.Labels = req.VpcFlowLogsConfig.Labels
		}
	}

	actual.UpdateTime = timestamppb.New(now)

	if err := s.storage.Update(ctx, fqn, actual); err != nil {
		return nil, err
	}

	return s.operations.StartLRO(ctx, fqn, nil, func() (proto.Message, error) {
		return actual, nil
	})
}

func (s *vpcFlowLogsService) DeleteVpcFlowLogsConfig(ctx context.Context, req *pb.DeleteVpcFlowLogsConfigRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseVpcFlowLogsConfigName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.VpcFlowLogsConfig{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return s.operations.StartLRO(ctx, fqn, nil, func() (proto.Message, error) {
		return &emptypb.Empty{}, nil
	})
}

type vpcFlowLogsConfigName struct {
	Project  *projects.ProjectData
	Location string
	ID       string
}

func (n *vpcFlowLogsConfigName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/vpcFlowLogsConfigs/%s", n.Project.ID, n.Location, n.ID)
}

func (s *MockService) parseVpcFlowLogsConfigName(name string) (*vpcFlowLogsConfigName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "vpcFlowLogsConfigs" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}
		return &vpcFlowLogsConfigName{
			Project:  project,
			Location: tokens[3],
			ID:       tokens[5],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid vpc flow logs config name", name)
}
