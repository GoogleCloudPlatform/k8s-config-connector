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

package mocksecuritycentermanagement

import (
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type eventThreatDetectionCustomModuleName struct {
	ParentType                         string // "projects", "folders", "organizations"
	ParentID                           string
	Location                           string
	EventThreatDetectionCustomModuleID string
}

func (n *eventThreatDetectionCustomModuleName) String() string {
	return n.ParentType + "/" + n.ParentID + "/locations/" + n.Location + "/eventThreatDetectionCustomModules/" + n.EventThreatDetectionCustomModuleID
}

// parseEventThreatDetectionCustomModuleName parses a string into an eventThreatDetectionCustomModuleName.
// Expected forms:
// - organizations/<orgID>/locations/<location>/eventThreatDetectionCustomModules/<moduleID>
// - folders/<folderID>/locations/<location>/eventThreatDetectionCustomModules/<moduleID>
// - projects/<projectID>/locations/<location>/eventThreatDetectionCustomModules/<moduleID>
func (s *MockService) parseEventThreatDetectionCustomModuleName(name string) (*eventThreatDetectionCustomModuleName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && (tokens[0] == "projects" || tokens[0] == "folders" || tokens[0] == "organizations") && tokens[2] == "locations" && tokens[4] == "eventThreatDetectionCustomModules" {
		return &eventThreatDetectionCustomModuleName{
			ParentType:                         tokens[0],
			ParentID:                           tokens[1],
			Location:                           tokens[3],
			EventThreatDetectionCustomModuleID: tokens[5],
		}, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
