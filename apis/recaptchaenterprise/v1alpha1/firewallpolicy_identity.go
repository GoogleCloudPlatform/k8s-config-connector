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

// FirewallPolicyIdentity defines the resource reference to ReCAPTCHAEnterpriseFirewallPolicy, which "External" field
// holds the GCP identifier for the KRM object.
type FirewallPolicyIdentity struct {
	parent *FirewallPolicyParent
	id     string
}

func (i *FirewallPolicyIdentity) String() string {
	return i.parent.String() + "/firewallpolicies/" + i.id
}

func (i *FirewallPolicyIdentity) ID() string {
	return i.id
}

func (i *FirewallPolicyIdentity) Parent() *FirewallPolicyParent {
	return i.parent
}

type FirewallPolicyParent struct {
	ProjectID string
}

func (p *FirewallPolicyParent) String() string {
	return "projects/" + p.ProjectID
}

// New builds a FirewallPolicyIdentity from the Config Connector FirewallPolicy object.
func NewFirewallPolicyIdentity(ctx context.Context, reader client.Reader, obj *ReCAPTCHAEnterpriseFirewallPolicy) (*FirewallPolicyIdentity, error) {
	// Get Parent
	projectRef, err := refsv1beta1.ResolveProject(ctx, reader, obj.GetNamespace(), obj.Spec.ProjectRef)
	if err != nil {
		return nil, err
	}
	projectID := projectRef.ProjectID
	if projectID == "" {
		return nil, fmt.Errorf("cannot resolve project")
	}

	// Get desired ID
	// This resource has a server-generated ID. This means user should not know
	// the ID before the resource is created, and 'metadata.name' won't be used
	// as the default resource ID. So empty value for 'spec.resourceID' should
	// also be valid:
	// 1. When 'spec.resourceID' is not set or set to an empty value, the
	//    intention is to create the resource.
	// 2. When 'spec.resourceID' is set, the intention is to acquire an existing
	//    resource.
	//    2.1. When 'spec.resourceID' is set but the corresponding GCP resource
	//         is not found, then it is a real error.
	resourceID := common.ValueOf(obj.Spec.ResourceID)

	// Use approved External
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		actualParent, actualResourceID, err := ParseFirewallPolicyExternal(externalRef)
		if err != nil {
			return nil, err
		}
		if actualParent.ProjectID != projectID {
			return nil, fmt.Errorf("spec.projectRef changed, expect %s, got %s", actualParent.ProjectID, projectID)
		}
		if resourceID != "" && actualResourceID != resourceID {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
				resourceID, actualResourceID)
		}
		resourceID = actualResourceID
	}
	return &FirewallPolicyIdentity{
		parent: &FirewallPolicyParent{
			ProjectID: projectID,
		},
		id: resourceID,
	}, nil
}

func ParseFirewallPolicyExternal(external string) (parent *FirewallPolicyParent, resourceID string, err error) {
	tokens := strings.Split(external, "/")
	if len(tokens) != 4 || tokens[0] != "projects" || tokens[2] != "firewallpolicies" {
		return nil, "", fmt.Errorf("format of ReCAPTCHAEnterpriseFirewallPolicy external=%q was not known (use projects/{{projectID}}/firewallpolicies/{{firewallpolicyID}})", external)
	}
	parent = &FirewallPolicyParent{
		ProjectID: tokens[1],
	}
	resourceID = tokens[3]
	return parent, resourceID, nil
}
