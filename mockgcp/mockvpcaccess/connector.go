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

// +tool:mockgcp-support
// krm.apiVersion: vpcaccess.cnrm.cloud.google.com/v1beta1
// krm.kind: VPCAccessConnector
// proto.service: google.cloud.vpcaccess.v1.VpcAccessService
// proto.resource: Connector

package mockvpcaccess

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/vpcaccess/v1"
)

func (s *vpcAccessService) CreateConnector(ctx context.Context, req *pb.CreateConnectorRequest) (*longrunning.Operation, error) {
	reqName := fmt.Sprintf("%s/connectors/%s", req.GetParent(), req.GetConnectorId())
	name, err := s.parseConnectorName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	obj := proto.Clone(req.GetConnector()).(*pb.Connector)
	obj.Name = fqn

	if obj.MinInstances == 0 {
		obj.MinInstances = 3
	}
	if obj.MaxInstances == 0 {
		obj.MaxInstances = 4
	}

	obj.State = pb.Connector_READY

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Region)
	metadata := &pb.OperationMetadata{
		Method:     "google.cloud.vpcaccess.v1beta1.VpcAccessService.CreateConnector",
		CreateTime: timestamppb.New(now),
		Target:     name.String(),
	}
	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}

func (s *vpcAccessService) GetConnector(ctx context.Context, req *pb.GetConnectorRequest) (*pb.Connector, error) {
	name, err := s.parseConnectorName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Connector{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The connector is not found.")
		}
		return nil, err
	}
	return obj, nil

}

func (s *vpcAccessService) DeleteConnector(ctx context.Context, req *pb.DeleteConnectorRequest) (*longrunning.Operation, error) {
	name, err := s.parseConnectorName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	deleted := &pb.Connector{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Region)
	metadata := &pb.OperationMetadata{
		Method:     "google.cloud.vpcaccess.v1beta1.VpcAccessService.DeleteConnector",
		CreateTime: timestamppb.New(now),
		Target:     name.String(),
	}
	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()
		return &emptypb.Empty{}, nil
	})
}

type connectorName struct {
	Project *projects.ProjectData
	Region  string
	Name    string
}

func (n *connectorName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/connectors/%s", n.Project.ID, n.Region, n.Name)
}

func (s *vpcAccessService) parseConnectorName(name string) (*connectorName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "connectors" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		n := &connectorName{
			Project: project,
			Region:  tokens[3],
			Name:    tokens[5],
		}

		return n, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "invalid connector name %q", name)
}
