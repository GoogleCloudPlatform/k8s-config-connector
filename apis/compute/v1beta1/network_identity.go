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
	"context"
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	parent "github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	commonref "github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/reference"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type NetworkIdentity struct {
	id     string
	parent parent.ProjectParent
}

func (i *NetworkIdentity) Parent() parent.ProjectParent {
	return i.parent
}

func (i *NetworkIdentity) String() string {
	return i.parent.String() + "/global/networks/" + i.id
}

func (i *NetworkIdentity) ID() string {
	return i.id
}

func (i *NetworkIdentity) FromExternal(external string) error {
	id, err := ParseComputeNetworkExternal(external)
	if err != nil {
		return fmt.Errorf("error parsing ComputeNetworkID from %q: %w", external, err)
	}
	i.parent.ProjectID = id.Parent().ProjectID
	i.id = id.id
	return nil
}

func ParseComputeNetworkExternal(external string) (*NetworkIdentity, error) {
	if external == "" {
		return nil, fmt.Errorf("empty ComputeNetwork external value")
	}
	trimmedExternal := commonref.FixStaleComputeExternalFormat(external)
	tokens := strings.Split(trimmedExternal, "/")
	if len(tokens) == 5 && tokens[0] == "projects" && tokens[2] == "global" && tokens[3] == "networks" {
		return &NetworkIdentity{
			parent: parent.ProjectParent{ProjectID: tokens[1]},
			id:     tokens[4],
		}, nil
	}
	return nil, fmt.Errorf("format of computenetwork external=%q was not known (use https://www.googleapis.com/compute/{{version}}/projects/{{projectId}}/global/networks/{{networkId}} or projects/{{projectId}}/global/networks/{{networkId}})", external)
}

// NewNetworkIdentity builds a NetworkIdentity from the Config Connector ComputeNetwork object.
func NewNetworkIdentity(ctx context.Context, reader client.Reader, obj *ComputeNetwork) (*NetworkIdentity, error) {
	// Get Parent
	projectID, err := refsv1beta1.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Get desired ID
	resourceID := common.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	// Use approved External
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		actual, err := ParseComputeNetworkExternal(externalRef)
		if err != nil {
			return nil, err
		}
		if actual.parent.ProjectID != projectID {
			return nil, fmt.Errorf("spec.projectRef changed, expect %s, got %s", actual.parent.ProjectID, projectID)
		}
		if actual.id != resourceID {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
				resourceID, actual.id)
		}
	}
	return &NetworkIdentity{
		parent: parent.ProjectParent{ProjectID: projectID},
		id:     resourceID,
	}, nil
}
