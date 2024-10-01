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
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type entitlementName struct {
	Container     string
	Location      string
	EntitlementID string
}

func (n *entitlementName) String() string {
	return n.parent() + "/entitlements/" + n.EntitlementID
}

func (n *entitlementName) parent() string {
	return n.Container + "/locations/" + n.Location
}

// parseEntitlementName parses a string into a entitlementName.
// The expected form is projects/<projectID>/locations/<region>/entitlements/<entitlementID>,
// or folders/<folderID>/locations/<region>/entitlements/<entitlementID>, or
// organizations/<organizationID>/locations/<region>/entitlements/<entitlementID>.
func (s *MockService) parseEntitlementName(name string) (*entitlementName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 &&
		(tokens[0] == "projects" || tokens[0] == "folders" || tokens[0] == "organizations") &&
		tokens[2] == "locations" && tokens[4] == "entitlements" {

		return &entitlementName{
			Container:     fmt.Sprintf("%s/%s", tokens[0], tokens[1]),
			Location:      tokens[3],
			EntitlementID: tokens[5],
		}, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
