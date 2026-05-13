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
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/fields"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/servicedirectory/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
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
	reqName := req.Parent + "/endpoints/" + req.GetEndpointId()
	name, err := s.parseEndpointName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.Endpoint).(*pb.Endpoint)
	obj.Name = fqn

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}
	return obj, nil
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

	if req.UpdateMask != nil {
		if err := fields.UpdateByFieldMask(obj, req.Endpoint, req.UpdateMask.Paths); err != nil {
			return nil, err
		}
	} else {
		proto.Merge(obj, req.Endpoint)
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

func (s *RegistrationServiceV1) ListEndpoints(ctx context.Context, req *pb.ListEndpointsRequest) (*pb.ListEndpointsResponse, error) {
	name, err := s.parseServiceName(req.Parent)
	if err != nil {
		return nil, err
	}

	prefix := name.String() + "/endpoints/"

	response := &pb.ListEndpointsResponse{}
	if err := s.storage.List(ctx, (&pb.Endpoint{}).ProtoReflect().Descriptor(), storage.ListOptions{Prefix: prefix}, func(obj proto.Message) error {
		response.Endpoints = append(response.Endpoints, obj.(*pb.Endpoint))
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}

type endpointName struct {
	serviceName
	EndpointId string
}

func (n *endpointName) String() string {
	return n.serviceName.String() + "/endpoints/" + n.EndpointId
}

// parseEndpointName parses a string into a endpointName.
// The expected form is projects/<projectID>/locations/<location>/namespaces/<namespace>/services/<service>/endpoints/<endpoint>
func (s *MockService) parseEndpointName(name string) (*endpointName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 10 && tokens[8] == "endpoints" {
		servicename, err := s.parseServiceName(strings.Join(tokens[0:8], "/"))
		if err != nil {
			return nil, err
		}

		name := &endpointName{
			serviceName: *servicename,
			EndpointId:  tokens[9],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
