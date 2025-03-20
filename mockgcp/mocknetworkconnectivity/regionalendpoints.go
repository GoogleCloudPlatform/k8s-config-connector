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
// proto.service: mockgcp.cloud.networkconnectivity.v1.ProjectsLocationsRegionalEndpoints
// proto.message: RegionalEndpoint

package mocknetworkconnectivity

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/networkconnectivity/v1"
)

func (s *networkConnectivityV1) GetRegionalEndpoint(ctx context.Context, req *pb.GetRegionalEndpointRequest) (*pb.RegionalEndpoint, error) {
	name, err := s.parseRegionalEndpointName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.RegionalEndpoint{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *networkConnectivityV1) CreateRegionalEndpoint(ctx context.Context, req *pb.CreateRegionalEndpointRequest) (*pb.RegionalEndpoint, error) {
	reqName := fmt.Sprintf("%s/regionalEndpoints/%s", req.GetParent(), req.GetRegionalEndpointId())
	name, err := s.parseRegionalEndpointName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := req.GetRegionalEndpoint()
	obj.Name = fqn

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *networkConnectivityV1) DeleteRegionalEndpoint(ctx context.Context, req *pb.DeleteRegionalEndpointRequest) (*emptypb.Empty, error) {
	name, err := s.parseRegionalEndpointName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.RegionalEndpoint{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

type regionalEndpointName struct {
	Project             *projects.ProjectData
	Location            string
	RegionalEndpointId string
}

func (n *regionalEndpointName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/regionalEndpoints/" + n.RegionalEndpointId
}

// parseRegionalEndpointName parses a string into an regionalEndpointName.
// The expected form is `projects/*/locations/*/regionalEndpoints/*`.
func (s *MockService) parseRegionalEndpointName(name string) (*regionalEndpointName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "regionalEndpoints" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &regionalEndpointName{
			Project:             project,
			Location:            tokens[3],
			RegionalEndpointId: tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

```


