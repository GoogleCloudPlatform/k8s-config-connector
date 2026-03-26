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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	common "github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/reference"
)

type ComputeVPNTunnelIdentity struct {
	id     string
	parent parent.ComputeParent
}

func (i *ComputeVPNTunnelIdentity) Parent() parent.ComputeParent {
	return i.parent
}

func (i *ComputeVPNTunnelIdentity) String() string {
	return i.parent.String() + "/vpnTunnels/" + i.id
}

func (i *ComputeVPNTunnelIdentity) ID() string {
	return i.id
}

func (i *ComputeVPNTunnelIdentity) FromExternal(external string) error {
	id, err := ParseComputeVPNTunnelExternal(external)
	if err != nil {
		return fmt.Errorf("error parsing ComputeVPNTunnelID from %q: %w", external, err)
	}
	i.parent.ProjectID = id.Parent().ProjectID
	i.parent.Location = id.Parent().Location
	i.id = id.id
	return nil
}

func ParseComputeVPNTunnelExternal(external string) (*ComputeVPNTunnelIdentity, error) {
	if external == "" {
		return nil, fmt.Errorf("empty ComputeVPNTunnel external value")
	}
	trimmedExternal := common.FixStaleComputeExternalFormat(external)
	tokens := strings.Split(trimmedExternal, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "regions" && tokens[4] == "vpnTunnels" {
		return &ComputeVPNTunnelIdentity{
			parent: parent.ComputeParent{ProjectID: tokens[1], Location: tokens[3]},
			id:     tokens[5],
		}, nil
	}
	return nil, fmt.Errorf("format of computevpntunnel external=%q was not known (use https://www.googleapis.com/compute/{{version}}/projects/{{projectId}}/regions/{{region}}/vpnTunnels/{{name}} or projects/{{projectId}}/regions/{{region}}/vpnTunnels/{{name}})", external)
}
