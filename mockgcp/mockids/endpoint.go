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

package mockids

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/ids/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

type CloudIDSEndpointV1 struct {
	*MockService
	pb.UnimplementedIDSServer
}

func (c *CloudIDSEndpointV1) ListEndpoints(ctx context.Context, request *pb.ListEndpointsRequest) (*pb.ListEndpointsResponse, error) {
	project, err := c.Projects.GetProjectByID(request.Parent)
	if err != nil {
		return nil, err
	}
	findPrefix := fmt.Sprintf("projects/%v/", project.ID)
	endpointKind := (&pb.Endpoint{}).ProtoReflect().Descriptor()

	var endpoints []*pb.Endpoint
	if err := c.storage.List(ctx, endpointKind, storage.ListOptions{}, func(obj proto.Message) error {
		endpoint := obj.(*pb.Endpoint)
		if strings.HasPrefix(endpoint.Name, findPrefix) {
			endpoints = append(endpoints, endpoint)
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return &pb.ListEndpointsResponse{
		Endpoints: endpoints,
	}, nil
}

func (c *CloudIDSEndpointV1) GetEndpoint(ctx context.Context, request *pb.GetEndpointRequest) (*pb.Endpoint, error) {
	endpoint, err := c.parseEndpointName(request.Name)
	if err != nil {
		return nil, err
	}

	fqn := endpoint.String()
	obj := &pb.Endpoint{}
	if err := c.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (c *CloudIDSEndpointV1) CreateEndpoint(ctx context.Context, request *pb.CreateEndpointRequest) (*longrunning.Operation, error) {
	// validate endpoint id:
	// This value must start with a lowercase letter followed by up to 62
	// lowercase letters, numbers, or hyphens, and cannot end with a hyphen.
	// Values that do not match this pattern will trigger an INVALID_ARGUMENT
	// error.
	re := regexp.MustCompile(`^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$`)
	if !re.MatchString(request.EndpointId) {
		return nil, status.Errorf(codes.InvalidArgument, "endpointID %q is malformed", request.EndpointId)
	}

	reqName := request.Parent + "/endpoints/" + request.EndpointId
	endpoint, err := c.parseEndpointName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := endpoint.String()
	obj := proto.Clone(request.Endpoint).(*pb.Endpoint)
	if err := c.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return c.operations.NewLRO(ctx)
}

func (c *CloudIDSEndpointV1) DeleteEndpoint(ctx context.Context, request *pb.DeleteEndpointRequest) (*longrunning.Operation, error) {
	endpoint, err := c.parseEndpointName(request.Name)
	if err != nil {
		return nil, err
	}

	fqn := endpoint.String()
	oldObj := &pb.Endpoint{}
	if err := c.storage.Delete(ctx, fqn, oldObj); err != nil {
		return nil, err
	}

	return c.operations.NewLRO(ctx)

}
