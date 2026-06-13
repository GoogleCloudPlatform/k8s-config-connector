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

package v1beta1

import (
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
)

var _ identity.Identity = &ComputeDiskIdentity{}

type ComputeDiskIdentity struct {
	ProjectID string
	Zone      string
	Region    string
	Name      string
}

func (i *ComputeDiskIdentity) String() string {
	if i.Region != "" {
		return fmt.Sprintf("projects/%s/regions/%s/disks/%s", i.ProjectID, i.Region, i.Name)
	}
	return fmt.Sprintf("projects/%s/zones/%s/disks/%s", i.ProjectID, i.Zone, i.Name)
}

func (i *ComputeDiskIdentity) FromExternal(s string) error {
	if s == "" {
		return fmt.Errorf("value cannot be empty")
	}
	tokens := strings.Split(s, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "zones" && tokens[4] == "disks" {
		i.ProjectID = tokens[1]
		i.Zone = tokens[3]
		i.Region = ""
		i.Name = tokens[5]
		return nil
	}

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "regions" && tokens[4] == "disks" {
		i.ProjectID = tokens[1]
		i.Zone = ""
		i.Region = tokens[3]
		i.Name = tokens[5]
		return nil
	}
	return fmt.Errorf("invalid format: %s, expected projects/{project}/zones/{zone}/disks/{name} or projects/{project}/regions/{region}/disks/{name}", s)
}
