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

package mockservicenetworking

import (
	"context"
	"fmt"
	"strings"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/servicenetworking/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"k8s.io/klog/v2"
)

type connectionsV1 struct {
	*MockService
	pb.UnimplementedServicesConnectionsServerServer
}

func (s *connectionsV1) ListServicesConnections(ctx context.Context, req *pb.ListServicesConnectionsRequest) (*pb.ListConnectionsResponse, error) {
	log := klog.FromContext(ctx)
	log.Info("ListServicesConnections", "request", req)

	serviceName, err := s.parseServiceName(req.Parent)
	if err != nil {
		return nil, err
	}

	network, err := s.parseNetworkName(req.GetNetwork())
	if err != nil {
		return nil, err
	}

	// fqn := name.String()

	response := &pb.ListConnectionsResponse{}

	// Network must not have any subnets depending on it
	subnetKind := (&pb.Connection{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, subnetKind, storage.ListOptions{}, func(obj proto.Message) error {
		connection := obj.(*pb.Connection)
		if connection.Service == serviceName.String() {
			if connection.Network == network.String() {
				response.Connections = append(response.Connections, connection)
			}
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}

func (s *connectionsV1) CreateServicesConnection(ctx context.Context, req *pb.CreateServicesConnectionRequest) (*longrunningpb.Operation, error) {
	log := klog.FromContext(ctx)
	log.Info("CreateServicesConnection", "request", req)

	// Parse the parent (service name) from the request
	serviceName, err := s.parseServiceName(req.GetParent())
	if err != nil {
		return nil, err
	}

	// Parse the network from the connection
	network, err := s.parseNetworkName(req.GetServicesConnection().GetNetwork())
	if err != nil {
		return nil, err
	}

	// Create a connection name based on the service and network
	connectionName := &serviceConnectionName{
		Service: serviceName.Name,
		Name:    network.Project.ID,
	}

	// Use the same fqn format as in PatchServicesConnection
	fqn := connectionName.String() + network.String()
	obj := &pb.Connection{
		Network:               network.String(),
		Peering:               "servicenetworking-googleapis-com",
		Service:               serviceName.String(),
		ReservedPeeringRanges: req.GetServicesConnection().GetReservedPeeringRanges(),
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return s.operations.StartLRO(ctx, "", nil, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *connectionsV1) PatchServicesConnection(ctx context.Context, req *pb.PatchServicesConnectionRequest) (*longrunningpb.Operation, error) {
	log := klog.FromContext(ctx)
	log.Info("PatchServicesConnection", "request", req)

	connectionName, err := s.parseServiceConnectionName(req.GetName())
	if err != nil {
		return nil, err
	}

	network, err := s.parseNetworkName(req.GetServicesConnection().GetNetwork())
	if err != nil {
		return nil, err
	}

	if connectionName.Name == "-" {
		connectionName.Name = network.Project.ID
	}

	create := false

	fqn := connectionName.String() + network.String()
	obj := &pb.Connection{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			create = true
		} else {
			return nil, err
		}
	}

	obj.Network = network.String()
	obj.Peering = "servicenetworking-googleapis-com"
	obj.Service = "services/servicenetworking.googleapis.com"

	// TODO: Check fieldPath
	obj.ReservedPeeringRanges = req.GetServicesConnection().GetReservedPeeringRanges()

	if create {
		if err := s.storage.Create(ctx, fqn, obj); err != nil {
			return nil, err
		}
	} else {
		if err := s.storage.Update(ctx, fqn, obj); err != nil {
			return nil, err
		}
	}

	return s.operations.StartLRO(ctx, "", nil, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *connectionsV1) DeleteConnectionServicesConnection(ctx context.Context, req *pb.DeleteConnectionServicesConnectionRequest) (*longrunningpb.Operation, error) {
	log := klog.FromContext(ctx)
	log.Info("DeleteServicesConnection", "request", req)

	connectionName, err := s.parseServiceConnectionName(req.GetName())
	if err != nil {
		return nil, err
	}

	network, err := s.parseNetworkName(req.GetServicesConnection().GetConsumerNetwork())
	if err != nil {
		return nil, err
	}

	if connectionName.Name == "-" {
		connectionName.Name = network.Project.ID
	}

	fqn := connectionName.String() + network.String()
	obj := &pb.Connection{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Connection %s does not exist", fqn)
		}
		return nil, err
	}

	if err := s.storage.Delete(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return s.operations.StartLRO(ctx, "", nil, func() (proto.Message, error) {
		return &emptypb.Empty{}, nil
	})
}

type serviceName struct {
	Name string
}

func (n *serviceName) String() string {
	return "services/" + n.Name
}

// parseServiceName parses a string into a serviceName.
// The expected form is `services/*`.
func (s *MockService) parseServiceName(name string) (*serviceName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 2 && tokens[0] == "services" {
		name := &serviceName{
			Name: tokens[1],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}

type serviceConnectionName struct {
	Service string
	Name    string
}

func (n *serviceConnectionName) String() string {
	return "services/" + n.Service + "/connections/" + n.Name
}

// parseServiceConnectionName parses a string into a serviceConnectionName.
// The expected form is `services/*/connections/-`.
func (s *MockService) parseServiceConnectionName(name string) (*serviceConnectionName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 4 && tokens[0] == "services" && tokens[2] == "connections" {
		name := &serviceConnectionName{
			Service: tokens[1],
			Name:    tokens[3],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}

type networkName struct {
	Project *projects.ProjectData
	Name    string
}

func (n *networkName) String() string {
	return fmt.Sprintf("projects/%d/global/networks/%s", n.Project.Number, n.Name)
}

// parseNetworkName parses a string into a networkName.
// The expected form is `projects/*/global/networks/{name}`.
func (s *MockService) parseNetworkName(name string) (*networkName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 5 && tokens[0] == "projects" && tokens[2] == "global" && tokens[3] == "networks" {
		project, err := s.Projects.GetProjectByNumber(tokens[1])
		if err != nil {
			return nil, err
		}
		name := &networkName{
			Project: project,
			Name:    tokens[4],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
