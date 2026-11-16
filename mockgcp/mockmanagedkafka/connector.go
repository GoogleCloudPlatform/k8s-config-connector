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
// proto.service: google.cloud.managedkafka.v1.ManagedKafkaConnect
// proto.message: google.cloud.managedkafka.v1.Connector

package mockmanagedkafka

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "cloud.google.com/go/managedkafka/apiv1/managedkafkapb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

func (s *managedKafkaConnect) CreateConnector(ctx context.Context, req *pb.CreateConnectorRequest) (*pb.Connector, error) {
	reqName := fmt.Sprintf("%s/connectors/%s", req.GetParent(), req.GetConnectorId())
	name, err := s.parseConnectorName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.GetConnector()).(*pb.Connector)
	obj.Name = fqn
	if obj.Configs == nil {
		obj.Configs = make(map[string]string)
	}
	if obj.Configs["name"] == "" {
		obj.Configs["name"] = name.Connector
	}
	if obj.GetTaskRestartPolicy() == nil {
		obj.RestartPolicy = &pb.Connector_TaskRestartPolicy{
			TaskRestartPolicy: &pb.TaskRetryPolicy{
				MinimumBackoff: durationpb.New(60 * time.Second),
				MaximumBackoff: durationpb.New(900 * time.Second),
			},
		}
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *managedKafkaConnect) GetConnector(ctx context.Context, req *pb.GetConnectorRequest) (*pb.Connector, error) {
	name, err := s.parseConnectorName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Connector{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Connector %s not found", name.Connector)
		}
		return nil, err
	}

	return obj, nil
}

func (s *managedKafkaConnect) ListConnectors(ctx context.Context, req *pb.ListConnectorsRequest) (*pb.ListConnectorsResponse, error) {
	name, err := s.parseConnectorName(req.GetParent() + "/connectors/dummy")
	if err != nil {
		return nil, err
	}

	response := &pb.ListConnectorsResponse{}

	findPrefix := strings.TrimSuffix(name.String(), "dummy")

	listKind := (&pb.Connector{}).ProtoReflect().Descriptor()

	if err := s.storage.List(ctx, listKind, storage.ListOptions{}, func(obj proto.Message) error {
		connector := obj.(*pb.Connector)
		if strings.HasPrefix(connector.GetName(), findPrefix) {
			response.Connectors = append(response.Connectors, connector)
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return response, nil
}

func (s *managedKafkaConnect) UpdateConnector(ctx context.Context, req *pb.UpdateConnectorRequest) (*pb.Connector, error) {
	name, err := s.parseConnectorName(req.GetConnector().GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	obj := &pb.Connector{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}
	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask must be provided")
	}

	for _, path := range paths {
		switch path {
		case "configs":
			obj.Configs = req.GetConnector().GetConfigs()
			if obj.Configs == nil {
				obj.Configs = make(map[string]string)
			}
			if obj.Configs["name"] == "" {
				obj.Configs["name"] = name.Connector
			}
		case "taskRestartPolicy", "task_restart_policy":
			if req.GetConnector().GetTaskRestartPolicy() != nil {
				obj.RestartPolicy = &pb.Connector_TaskRestartPolicy{
					TaskRestartPolicy: req.GetConnector().GetTaskRestartPolicy(),
				}
			}
		case "name":
			obj.Name = req.GetConnector().GetName()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "field %q is not yet handled in mock", path)
		}
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *managedKafkaConnect) DeleteConnector(ctx context.Context, req *pb.DeleteConnectorRequest) (*emptypb.Empty, error) {
	name, err := s.parseConnectorName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.Connector{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

type connectorName struct {
	Project        *projects.ProjectData
	Location       string
	ConnectCluster string
	Connector      string
}

func (n *connectorName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/connectClusters/%s/connectors/%s", n.Project.ID, n.Location, n.ConnectCluster, n.Connector)
}

func (s *MockService) parseConnectorName(name string) (*connectorName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "connectClusters" && tokens[6] == "connectors" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		return &connectorName{
			Project:        project,
			Location:       tokens[3],
			ConnectCluster: tokens[5],
			Connector:      tokens[7],
		}, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "invalid connector name %q", name)
}
