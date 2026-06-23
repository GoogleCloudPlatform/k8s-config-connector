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

package mockservicedirectory

import (
	"context"
	"fmt"
	"net"
	"sort"
	"strings"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/servicedirectory/v1beta1"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *RegistrationServiceV1) GetEndpoint(ctx context.Context, req *pb.GetEndpointRequest) (*pb.Endpoint, error) {
	name, err := s.parseEndpointName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Endpoint{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *RegistrationServiceV1) CreateEndpoint(ctx context.Context, req *pb.CreateEndpointRequest) (*pb.Endpoint, error) {
	if err := s.validateIP(req.Parent, req.GetEndpointId(), req.GetEndpoint()); err != nil {
		return nil, err
	}

	reqName := req.Parent + "/endpoints/" + req.GetEndpointId()
	name, err := s.parseEndpointName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.CloneOf(req.Endpoint)
	obj.Name = fqn

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *RegistrationServiceV1) validateIP(parent, endpointID string, endpoint *pb.Endpoint) error {
	if endpoint == nil {
		return nil
	}
	address := endpoint.GetAddress()
	if address == "" {
		return nil
	}
	if net.ParseIP(address) != nil {
		return nil
	}

	// Address is not empty and not a valid IP
	var metadataStrings []string
	keys := make([]string, 0, len(endpoint.Metadata))
	for k := range endpoint.Metadata {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		metadataStrings = append(metadataStrings, fmt.Sprintf("metadata { key: %q value: %q }", k, endpoint.Metadata[k]))
	}

	endpointParts := []string{}
	if endpoint.Address != "" {
		endpointParts = append(endpointParts, fmt.Sprintf("address: %q", endpoint.Address))
	}
	if endpoint.Port != 0 {
		endpointParts = append(endpointParts, fmt.Sprintf("port: %d", endpoint.Port))
	}
	if len(metadataStrings) > 0 {
		endpointParts = append(endpointParts, strings.Join(metadataStrings, " "))
	}
	if endpoint.Network != "" {
		endpointParts = append(endpointParts, fmt.Sprintf("network: %q", endpoint.Network))
	}
	endpointStr := strings.Join(endpointParts, " ")

	errMsg := fmt.Sprintf("IP address invalid. IP must be empty, IPv4 or IPv6; request 'go/debugonly   parent: %q endpoint_id: %q endpoint { %s }'", parent, endpointID, endpointStr)
	return status.Errorf(codes.InvalidArgument, "%s", errMsg)
}

func (s *RegistrationServiceV1) UpdateEndpoint(ctx context.Context, req *pb.UpdateEndpointRequest) (*pb.Endpoint, error) {
	reqName := req.GetEndpoint().GetName()

	name, err := s.parseEndpointName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Endpoint{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// Apply field mask updates
	paths := req.GetUpdateMask().GetPaths()
	for _, path := range paths {
		switch path {
		case "address":
			obj.Address = req.GetEndpoint().GetAddress()
		case "port":
			obj.Port = req.GetEndpoint().GetPort()
		case "network":
			obj.Network = req.GetEndpoint().GetNetwork()
		case "metadata":
			obj.Metadata = req.GetEndpoint().GetMetadata()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not supported", path)
		}
	}

	if err := s.validateIP(name.serviceName.String(), name.EndpointName, obj); err != nil {
		return nil, err
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *RegistrationServiceV1) DeleteEndpoint(ctx context.Context, req *pb.DeleteEndpointRequest) (*empty.Empty, error) {
	name, err := s.parseEndpointName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	oldObj := &pb.Endpoint{}
	if err := s.storage.Delete(ctx, fqn, oldObj); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

type endpointName struct {
	serviceName
	EndpointName string
}

func (n *endpointName) String() string {
	return n.serviceName.String() + "/endpoints/" + n.EndpointName
}

// parseEndpointName parses a string into an endpointName.
// The expected form is projects/<projectID>/locations/<location>/namespaces/<namespace>/services/<service>/endpoints/<endpoint>
func (s *MockService) parseEndpointName(name string) (*endpointName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 10 && tokens[8] == "endpoints" {
		servicename, err := s.parseServiceName(strings.Join(tokens[0:8], "/"))
		if err != nil {
			return nil, err
		}

		name := &endpointName{
			serviceName:  *servicename,
			EndpointName: tokens[9],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
