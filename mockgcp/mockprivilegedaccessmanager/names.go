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

package mockprivilegedaccessmanager

import (
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
)

type entitlementName struct {
	Project       *projects.ProjectData
	Location      string
	EntitlementID string
}

func (n *entitlementName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/entitlements/" + n.EntitlementID
}

// parseEntitlementName parses a string into a entitlementName.
// The expected form is projects/<projectID>/locations/<region>/entitlements/<entitlementID>
func (s *MockService) parseEntitlementName(name string) (*entitlementName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "entitlements" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &entitlementName{
			Project:       project,
			Location:      tokens[3],
			EntitlementID: tokens[5],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
