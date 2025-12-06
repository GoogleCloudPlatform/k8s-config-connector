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
// proto.service: google.cloud.apphub.v1
// proto.message: google.cloud.apphub.v1.DiscoveredService

package mockapphub

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/apphub/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

func (s *AppHubV1Service) ListDiscoveredServices(ctx context.Context, req *pb.ListDiscoveredServicesRequest) (*pb.ListDiscoveredServicesResponse, error) {

	response := &pb.ListDiscoveredServicesResponse{}

	findPrefix := req.GetParent()
	findKind := (&pb.DiscoveredService{}).ProtoReflect().Descriptor()

	if err := s.storage.List(ctx, findKind, storage.ListOptions{
		Prefix: findPrefix,
	}, func(obj proto.Message) error {
		discoveredService := obj.(*pb.DiscoveredService)
		response.DiscoveredServices = append(response.DiscoveredServices, discoveredService)
		return nil
	}); err != nil {
		return nil, err
	}
	return response, nil
}

func (s *AppHubV1Service) GetDiscoveredService(ctx context.Context, req *pb.GetDiscoveredServiceRequest) (*pb.DiscoveredService, error) {
	name, err := s.parseDiscoveredServiceName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.DiscoveredService{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

type discoveredServiceName struct {
	Project            *projects.ProjectData
	Location           string
	DiscoveredServices string
}

func (n *discoveredServiceName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/discoveredServices/%s", n.Project.ID, n.Location, n.DiscoveredServices)
}

// parseApplicationName parses a string into an applicationName.
// The expected form is `projects/*/locations/*/discoveredServices/*`.
func (s *AppHubV1Service) parseDiscoveredServiceName(name string) (*discoveredServiceName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "discoveredServices" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &discoveredServiceName{
			Project:            project,
			Location:           tokens[3],
			DiscoveredServices: tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
