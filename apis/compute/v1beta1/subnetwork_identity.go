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

package v1beta1

import (
	"fmt"
	"strings"

	common "github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/reference"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
)

type SubnetworkIdentity struct {
	ResourceID string
	Parent     *parent.ComputeParent
}

func (i *SubnetworkIdentity) String() string {
	return i.Parent.String() + "/subnetworks/" + i.ResourceID
}

func (i *SubnetworkIdentity) FromExternal(external string) error {
	id, err := ParseComputeSubnetworkExternal(external)
	if err != nil {
		return fmt.Errorf("error parsing ComputeSubnetworkID from %q: %w", external, err)
	}
	i.Parent.ProjectID = id.Parent.ProjectID
	i.ResourceID = id.ResourceID
	return nil
}

func ParseComputeSubnetworkExternal(external string) (*SubnetworkIdentity, error) {
	if external == "" {
		return nil, fmt.Errorf("empty ComputeSubnetwork external value")
	}
	trimmedExternal := common.FixStaleComputeExternalFormat(external)
	tokens := strings.Split(trimmedExternal, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "regions" && tokens[4] == "subnetworks" {
		return &SubnetworkIdentity{
			Parent:     &parent.ComputeParent{ProjectID: tokens[1], Location: tokens[3]},
			ResourceID: tokens[5],
		}, nil
	}
	return nil, fmt.Errorf("format of ComputeSubnetwork external=%q was not known (use https://www.googleapis.com/compute/{{version}}/projects/{{projectId}}/regions/{{region}}/subnetworks/{{subnetworkId}} or projects/{{projectId}}/regions/{{region}}/subnetworks/{{subnetworkId}})", external)
}
