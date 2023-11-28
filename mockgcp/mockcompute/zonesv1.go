// Copyright 2022 Google LLC
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
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
)

type ZonesV1 struct {
	*MockService
	pb.UnimplementedZonesServer
}

func (s *ZonesV1) Get(ctx context.Context, req *pb.GetZoneRequest) (*pb.Zone, error) {
	reqName := "projects/" + req.GetProject() + "/zones/" + req.GetZone()
	name, err := s.parseZoneName(reqName)
	if err != nil {
		return nil, err
	}

	region := strings.Join(strings.Split(name.Zone, "-")[:2], "-")
	region = fmt.Sprintf("https://www.googleapis.com/compute/v1/projects/%s/regions/%s", name.Project.ID, region)

	obj := &pb.Zone{}

	obj.Kind = PtrTo("compute#zone")
	obj.Name = PtrTo(name.Zone)
	obj.Status = PtrTo("UP")
	obj.Region = &region
	obj.SelfLink = PtrTo("https://www.googleapis.com/compute/v1" + name.String())

	return obj, nil
}

type zoneName struct {
	Project *projects.ProjectData
	Zone    string
}

func (n *zoneName) String() string {
	return "projects/" + n.Project.ID + "/zones/" + n.Zone
}

// parseZoneName parses a string into a zoneName.
// The expected form is `projects/*/zones/`.
func (s *MockService) parseZoneName(name string) (*zoneName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "zones" {
		project, err := s.projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &zoneName{
			Project: project,
			Zone:    tokens[3],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
