// Copyright 2025 Google LLC
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

package v1alpha1

import (
	"fmt"
	"strings"
)

type ConnectionProfileIdentity struct {
	Project           string
	Location          string
	ConnectionProfile string
}

func (i *ConnectionProfileIdentity) FromExternal(external string) error {
	tokens := strings.Split(external, "/")
	if len(tokens) != 6 {
		return fmt.Errorf("invalid external format: %q, expected projects/{project}/locations/{location}/connectionProfiles/{connectionProfile}", external)
	}
	if tokens[0] != "projects" {
		return fmt.Errorf("invalid external format: %q, expected projects segment", external)
	}
	if tokens[2] != "locations" {
		return fmt.Errorf("invalid external format: %q, expected locations segment", external)
	}
	if tokens[4] != "connectionProfiles" {
		return fmt.Errorf("invalid external format: %q, expected connectionProfiles segment", external)
	}
	i.Project = tokens[1]
	i.Location = tokens[3]
	i.ConnectionProfile = tokens[5]
	return nil
}
