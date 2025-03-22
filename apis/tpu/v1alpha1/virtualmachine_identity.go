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

// +tool:krm-identity
// proto.service: google.cloud.tpu.v2.Tpu
// proto.message: google.cloud.tpu.v2.Node
// crd.type: TPUVirtualMachine
// crd.version: v1alpha1

import (
	"context"
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ identity.Identity = &TPUVirtualMachineIdentity{}

// +k8s:deepcopy-gen=false
type TPUVirtualMachineIdentity struct {
	ParentID *parent.ProjectAndLocationParent
	Node     string
}

func (i *TPUVirtualMachineIdentity) String() string {
	return i.ParentID.String() + "/nodes/" + i.Node
}

func (i *TPUVirtualMachineIdentity) FromExternal(external string) error {
	tokens := strings.Split(external, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "nodes" {
		i.ParentID = &parent.ProjectAndLocationParent{
			ProjectID: tokens[1],
			Location:  tokens[3],
		}
		i.Node = tokens[5]
		return nil
	}
	return fmt.Errorf("format of TPUVirtualMachine external=%q was not known (use projects/{{projectID}}/locations/{{location}}/nodes/{{nodeID}})", external)
}

var _ identity.Resource = &TPUVirtualMachine{}

func (obj *TPUVirtualMachine) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	// Get parent ID
	parentID, err := obj.GetParentIdentity(ctx, reader)
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

	id := &TPUVirtualMachineIdentity{
		ParentID: parentID.(*parent.ProjectAndLocationParent),
		Node:     resourceID,
	}

	// Attempt to ensure ID is immutable, by verifying against previously-set `status.externalRef`.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		previousID := &TPUVirtualMachineIdentity{}
		if err := previousID.FromExternal(externalRef); err != nil {
			return nil, err
		}
		if id.String() != previousID.String() {
			return nil, fmt.Errorf("cannot update TPUVirtualMachine identity (old=%q, new=%q): identity is immutable", previousID.String(), id.String())
		}
	}

	return id, nil
}

func (obj *TPUVirtualMachine) GetParentIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	// TODO: Can we extract helper?

	// Normalize projectRef
	if err := obj.Spec.ProjectRef.Normalize(ctx, reader, obj.GetNamespace()); err != nil {
		return nil, err
	}

	location := obj.Spec.Location

	external := obj.Spec.ProjectRef.External + "/locations/" + location

	// Get parent identity
	parentID := &parent.ProjectAndLocationParent{}
	if err := parentID.FromExternal(external); err != nil {
		return nil, err
	}
	return parentID, nil
}
