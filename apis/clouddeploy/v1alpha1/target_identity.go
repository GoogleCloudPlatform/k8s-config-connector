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
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ identity.Identity = &TargetIdentity{}
var _ identity.Resource = &CloudDeployTarget{}

// TargetIdentity defines the resource reference to CloudDeployTarget, which "External" field
// holds the GCP identifier for the KRM object.
type TargetIdentity struct {
	parent *TargetParent
	id     string
}

func (i *TargetIdentity) String() string {
	return i.parent.String() + "/targets/" + i.id
}

func (i *TargetIdentity) ID() string {
	return i.id
}

func (i *TargetIdentity) Parent() *TargetParent {
	return i.parent
}

type TargetParent struct {
	ProjectID string
	Location  string
}

func (p *TargetParent) String() string {
	return "projects/" + p.ProjectID + "/locations/" + p.Location
}

// FromExternal parses a GCP resource string (e.g., "projects/{project}/locations/{location}/targets/{target}")
// and populates the TargetIdentity fields.
func (i *TargetIdentity) FromExternal(ref string) error {
	parts := strings.Split(ref, "/")
	if len(parts) != 6 || parts[0] != "projects" || parts[2] != "locations" || parts[4] != "targets" {
		return fmt.Errorf("invalid format for CloudDeployTarget external=%q", ref)
	}
	i.parent = &TargetParent{
		ProjectID: parts[1],
		Location:  parts[3],
	}
	i.id = parts[5]
	return nil
}

// GetIdentity constructs a TargetIdentity from the object's Spec (Project, Location, Name/ResourceID).
// It also validates against Status.ExternalRef if present to ensure consistency.
func (obj *CloudDeployTarget) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	// Build Identity from Spec
	project, err := refsv1beta1.ResolveProject(ctx, reader, obj.GetNamespace(), obj.Spec.ProjectRef)
	if err != nil {
		return nil, err
	}
	projectID := project.ProjectID
	location := common.ValueOf(obj.Spec.Location)
	if location == "" {
		return nil, fmt.Errorf("spec.location is required")
	}
	id := common.ValueOf(obj.Spec.ResourceID)
	if id == "" {
		id = obj.GetName()
	}
	newIdentity := &TargetIdentity{
		parent: &TargetParent{
			ProjectID: projectID,
			Location:  location,
		},
		id: id,
	}

	// Validate against the ID stored in status.externalRef, if any
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		statusIdentity := &TargetIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, fmt.Errorf("cannot parse existing externalRef=%q: %w", externalRef, err)
		}
		if statusIdentity.String() != newIdentity.String() {
			return nil, fmt.Errorf("existing externalRef=%q does not match the identity resolved from spec: %q", externalRef, newIdentity.String())
		}
	}
	return newIdentity, nil
}

// New builds a TargetIdentity from the Config Connector Target object.
func NewTargetIdentity(ctx context.Context, reader client.Reader, obj *CloudDeployTarget) (*TargetIdentity, error) {
	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	return id.(*TargetIdentity), nil
}
