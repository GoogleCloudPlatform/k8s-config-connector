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
	common "github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/reference"
)

// ComputeNodeGroupIdentity is the identity of a ComputeNodeGroup.
// TODO: Move this to nodegroup_identity.go once it is generated/implemented
type ComputeNodeGroupIdentity struct {
	Project   string
	Zone      string
	NodeGroup string
}

var _ identity.Identity = &ComputeNodeGroupIdentity{}

func (i *ComputeNodeGroupIdentity) String() string {
	return fmt.Sprintf("projects/%s/zones/%s/nodeGroups/%s", i.Project, i.Zone, i.NodeGroup)
}

func (i *ComputeNodeGroupIdentity) FromExternal(ref string) error {
	if ref == "" {
		return fmt.Errorf("empty external reference")
	}

	trimmed := common.FixStaleComputeExternalFormat(ref)
	tokens := strings.Split(trimmed, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "zones" && tokens[4] == "nodeGroups" {
		i.Project = tokens[1]
		i.Zone = tokens[3]
		i.NodeGroup = tokens[5]
		return nil
	}

	return fmt.Errorf("invalid format for ComputeNodeGroup external reference: %q. "+
		"Expected format: projects/{projectID}/zones/{zone}/nodeGroups/{nodeGroup}", ref)
}
