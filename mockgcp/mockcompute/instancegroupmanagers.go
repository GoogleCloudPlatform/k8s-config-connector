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

package mockcompute

import (
	"context"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
)

type InstanceGroupManagersV1 struct {
	*MockService
	pb.UnimplementedInstanceGroupManagersServer
}

func (s *InstanceGroupManagersV1) Get(ctx context.Context, req *pb.GetInstanceGroupManagerRequest) (*pb.InstanceGroupManager, error) {
	reqName := "projects/" + req.GetProject() + "/zones/" + req.GetZone() + "/instanceGroupManagers/" + req.GetInstanceGroupManager()
	name, err := s.parseZonalInstanceGroupManagerName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.InstanceGroupManager{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			// Synthetic support for GKE IGMs
			if strings.HasPrefix(name.Name, "gke-") {
				obj = &pb.InstanceGroupManager{
					Name:          PtrTo(name.Name),
					SelfLink:      PtrTo(buildComputeSelfLink(ctx, fqn)),
					Zone:          PtrTo(buildComputeSelfLink(ctx, "projects/"+name.Project.ID+"/zones/"+name.Zone)),
					Id:            PtrTo(s.generateID()),
					InstanceGroup: PtrTo(buildComputeSelfLink(ctx, fqn[:len(fqn)-4])), // approximate
					Status: &pb.InstanceGroupManagerStatus{
						IsStable: PtrTo(true),
					},
				}
				return obj, nil
			}
		}
		return nil, err
	}

	return obj, nil
}

type zonalInstanceGroupManagerName struct {
	Project *projects.ProjectData
	Zone    string
	Name    string
}

func (n *zonalInstanceGroupManagerName) String() string {
	return "projects/" + n.Project.ID + "/zones/" + n.Zone + "/instanceGroupManagers/" + n.Name
}

// parseZonalInstanceGroupManagerName parses a string into a zonalInstanceGroupManagerName.
// The expected form is `projects/*/zones/*/instanceGroupManagers/*`.
func (s *MockService) parseZonalInstanceGroupManagerName(name string) (*zonalInstanceGroupManagerName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "zones" && tokens[4] == "instanceGroupManagers" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &zonalInstanceGroupManagerName{
			Project: project,
			Zone:    tokens[3],
			Name:    tokens[5],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
