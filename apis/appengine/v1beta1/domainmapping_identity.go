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
	"context"
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// DomainMappingIdentity defines the resource reference to AppEngineDomainMapping, which "External" field
// holds the GCP identifier for the KRM object.
type DomainMappingIdentity struct {
	parent *DomainMappingParent
	id     string
}

func (i *DomainMappingIdentity) String() string {
	return i.parent.String() + "/domainMappings/" + i.id
}

func (i *DomainMappingIdentity) ID() string {
	return i.id
}

func (i *DomainMappingIdentity) Parent() *DomainMappingParent {
	return i.parent
}

type DomainMappingParent struct {
	ProjectID string
	Location  string
}

func (p *DomainMappingParent) String() string {
	return "apps/" + p.ProjectID
}

// New builds a DomainMappingIdentity from the Config Connector AppEngineDomainMapping object.
func NewDomainMappingIdentity(ctx context.Context, reader client.Reader, obj *AppEngineDomainMapping) (*DomainMappingIdentity, error) {

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

	// '.status.externalRef' could be unset before the upgrade so only parse it
	// when the value is non-nil.
	var externalRef string
	if obj.Status.ExternalRef != nil {
		// Use approved ExternalRef
		externalRef = common.ValueOf(obj.Status.ExternalRef)
	}

	if externalRef != "" {
		// Validate desired with actual
		actualParent, actualResourceID, err := ParseDomainMappingExternal(externalRef)
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
	return &DomainMappingIdentity{
		parent: &DomainMappingParent{
			ProjectID: projectID,
			Location:  location,
		},
		id: resourceID,
	}, nil
}

func ParseDomainMappingExternal(external string) (parent *DomainMappingParent, resourceID string, err error) {
	external = strings.TrimPrefix(external, "//appengine.googleapis.com/")
	tokens := strings.Split(external, "/")
	if len(tokens) != 4 || tokens[0] != "apps" || tokens[2] != "domainMappings" {
		return nil, "", fmt.Errorf("format of AppEngineDomainMapping external=%q was not known (use apps/{{projectID}}/domainMappings/{{domainMappingID}})", external)
	}
	parent = &DomainMappingParent{
		ProjectID: tokens[1],
		Location:  "", // AppEngine doesn't use location in the path, but we store it for consistency
	}
	resourceID = tokens[3]
	return parent, resourceID, nil
}
