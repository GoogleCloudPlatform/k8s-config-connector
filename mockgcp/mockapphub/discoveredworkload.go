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
// proto.service: google.cloud.apphub.v1.AppHub
// proto.message: google.cloud.apphub.v1.DiscoveredWorkload

package mockapphub

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	pb "cloud.google.com/go/apphub/apiv1/apphubpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

func (s *AppHubV1Service) ListDiscoveredWorkloads(ctx context.Context, req *pb.ListDiscoveredWorkloadsRequest) (*pb.ListDiscoveredWorkloadsResponse, error) {
	response := &pb.ListDiscoveredWorkloadsResponse{}

	findPrefix := req.GetParent()
	findKind := (&pb.DiscoveredWorkload{}).ProtoReflect().Descriptor()

	if err := s.storage.List(ctx, findKind, storage.ListOptions{
		Prefix: findPrefix,
	}, func(obj proto.Message) error {
		discoveredWorkload := obj.(*pb.DiscoveredWorkload)
		response.DiscoveredWorkloads = append(response.DiscoveredWorkloads, discoveredWorkload)
		return nil
	}); err != nil {
		return nil, err
	}
	return response, nil
}

func (s *AppHubV1Service) GetDiscoveredWorkload(ctx context.Context, req *pb.GetDiscoveredWorkloadRequest) (*pb.DiscoveredWorkload, error) {
	name, err := s.parseDiscoveredWorkloadName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.DiscoveredWorkload{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			// Auto-populate discovered workload for mock testing so that Get succeeds!
			obj = &pb.DiscoveredWorkload{
				Name: fqn,
				WorkloadReference: &pb.WorkloadReference{
					Uri: "//container.googleapis.com/projects/" + name.Project.ID + "/locations/" + name.Location + "/clusters/mock-cluster/k8s/namespaces/mock-ns/apps/" + name.DiscoveredWorkloads,
				},
				WorkloadProperties: &pb.WorkloadProperties{
					GcpProject: name.Project.ID,
					Location:   name.Location,
					Zone:       name.Location + "-a",
				},
			}
			if err := s.storage.Create(ctx, fqn, obj); err != nil {
				return nil, err
			}
			return obj, nil
		}
		return nil, err
	}

	return obj, nil
}

type discoveredWorkloadName struct {
	Project             *projects.ProjectData
	Location            string
	DiscoveredWorkloads string
}

func (n *discoveredWorkloadName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/discoveredworkloads/%s", n.Project.ID, n.Location, n.DiscoveredWorkloads)
}

// parseDiscoveredWorkloadName parses a string into a discoveredWorkloadName.
// The expected form is `projects/*/locations/*/discoveredWorkloads/*` or `projects/*/locations/*/discoveredworkloads/*`.
func (s *AppHubV1Service) parseDiscoveredWorkloadName(name string) (*discoveredWorkloadName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && (tokens[4] == "discoveredWorkloads" || tokens[4] == "discoveredworkloads") {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &discoveredWorkloadName{
			Project:             project,
			Location:            tokens[3],
			DiscoveredWorkloads: tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
