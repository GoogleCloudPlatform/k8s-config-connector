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
	"context"
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// NetworkPolicyIdentity defines the resource reference to VMwareEngineNetworkPolicy, which "External" field
// holds the GCP identifier for the KRM object.
type NetworkPolicyIdentity struct {
	parent *NetworkPolicyParent
	id     string
}

func (i *NetworkPolicyIdentity) String() string {
	return i.parent.String() + "/networkpolicys/" + i.id
}

func (i *NetworkPolicyIdentity) ID() string {
	return i.id
}

func (i *NetworkPolicyIdentity) Parent() *NetworkPolicyParent {
	return i.parent
}

type NetworkPolicyParent struct {
	ProjectID string
	Location  string
}

func (p *NetworkPolicyParent) String() string {
	return "projects/" + p.ProjectID + "/locations/" + p.Location
}

// New builds a NetworkPolicyIdentity from the Config Connector NetworkPolicy object.
func NewNetworkPolicyIdentity(ctx context.Context, reader client.Reader, obj *VMwareEngineNetworkPolicy) (*NetworkPolicyIdentity, error) {

	// Get Parent
	projectRef, err := refsv1beta1.ResolveProject(ctx, reader, obj.GetNamespace(), obj.Spec.ProjectRef)
	if err != nil {
		return nil, err
	}
	projectID := projectRef.ProjectID
	if projectID == "" {
		return nil, fmt.Errorf("cannot resolve project")
	}
	location := obj.Spec.Location

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
		actualParent, actualResourceID, err := ParseNetworkPolicyExternal(externalRef)
		if err != nil {
			return nil, err
		}
		if actualParent.ProjectID != projectID {
			return nil, fmt.Errorf("spec.projectRef changed, expect %s, got %s", actualParent.ProjectID, projectID)
		}
		if actualParent.Location != location {
			return nil, fmt.Errorf("spec.location changed, expect %s, got %s", actualParent.Location, location)
		}
		if actualResourceID != resourceID {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
				resourceID, actualResourceID)
		}
	}
	return &NetworkPolicyIdentity{
		parent: &NetworkPolicyParent{
			ProjectID: projectID,
			Location:  location,
		},
		id: resourceID,
	}, nil
}

func ParseNetworkPolicyExternal(external string) (parent *NetworkPolicyParent, resourceID string, err error) {
	tokens := strings.Split(external, "/")
	if len(tokens) != 6 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "networkpolicys" {
		return nil, "", fmt.Errorf("format of VMwareEngineNetworkPolicy external=%q was not known (use projects/{{projectID}}/locations/{{location}}/networkpolicys/{{networkpolicyID}})", external)
	}
	parent = &NetworkPolicyParent{
		ProjectID: tokens[1],
		Location:  tokens[3],
	}
	resourceID = tokens[5]
	return parent, resourceID, nil
}
