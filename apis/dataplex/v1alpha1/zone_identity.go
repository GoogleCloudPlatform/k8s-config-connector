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
	// refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1" // Removed as ProjectRef is not used directly for parent resolution here.
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// ZoneIdentity defines the resource reference to DataplexZone, which "External" field
// holds the GCP identifier for the KRM object.
type ZoneIdentity struct {
	parent *ZoneParent
	id     string
}

func (i *ZoneIdentity) String() string {
	return i.parent.String() + "/zones/" + i.id
}

func (i *ZoneIdentity) ID() string {
	return i.id
}

func (i *ZoneIdentity) Parent() *ZoneParent {
	return i.parent
}

// ZoneParent represents the parent hierarchy for a Dataplex Zone.
type ZoneParent struct {
	ProjectID string
	Location  string
	LakeID    string
}

func (p *ZoneParent) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/lakes/%s", p.ProjectID, p.Location, p.LakeID)
}

// NewZoneIdentity builds a ZoneIdentity from the Config Connector DataplexZone object.
func NewZoneIdentity(ctx context.Context, reader client.Reader, obj *DataplexZone) (*ZoneIdentity, error) {
	// Get Parent details from LakeRef
	if obj.Spec.LakeRef == nil || obj.Spec.LakeRef.Name == "" {
		return nil, fmt.Errorf("spec.lakeRef.name is required")
	}
	lakeRefName := obj.Spec.LakeRef.Name

	// Parse lakeRefName: "projects/{project_id}/locations/{location_id}/lakes/{lake_id}"
	lakeTokens := strings.Split(lakeRefName, "/")
	if len(lakeTokens) != 6 || lakeTokens[0] != "projects" || lakeTokens[2] != "locations" || lakeTokens[4] != "lakes" {
		return nil, fmt.Errorf("format of spec.lakeRef.name=%q was not known (use projects/{{project}}/locations/{{location}}/lakes/{{lake}})", lakeRefName)
	}
	projectID := lakeTokens[1]
	location := lakeTokens[3]
	lakeID := lakeTokens[5]

	// Get desired Zone ID
	resourceID := common.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID for zone")
	}

	parentFromSpec := &ZoneParent{
		ProjectID: projectID,
		Location:  location,
		LakeID:    lakeID,
	}

	// Use approved External if available and validate
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		actualParent, actualResourceID, err := ParseZoneExternal(externalRef)
		if err != nil {
			return nil, fmt.Errorf("failed to parse existing externalRef %q: %w", externalRef, err)
		}
		if actualParent.ProjectID != parentFromSpec.ProjectID {
			return nil, fmt.Errorf("parent project changed, expect %q, got %q", actualParent.ProjectID, parentFromSpec.ProjectID)
		}
		if actualParent.Location != parentFromSpec.Location {
			return nil, fmt.Errorf("parent location changed, expect %q, got %q", actualParent.Location, parentFromSpec.Location)
		}
		if actualParent.LakeID != parentFromSpec.LakeID {
			return nil, fmt.Errorf("parent lake changed, expect %q, got %q", actualParent.LakeID, parentFromSpec.LakeID)
		}
		// We allow metadata.name or spec.resourceID to be specified even if externalRef exists,
		// as long as the specified ID matches the actual resource ID from externalRef.
		if actualResourceID != resourceID {
			// If the user tries to change the resource ID after creation, we should error out.
			// This check handles the case where the user changes metadata.name or spec.resourceID
			// after the resource has been created and its externalRef is populated.
			return nil, fmt.Errorf("cannot change resource ID from %q to %q for an existing resource", actualResourceID, resourceID)
		}
		// If validation passes, use the validated parent and ID derived from externalRef
		// This ensures we use the canonical identifiers from the existing resource.
		return &ZoneIdentity{
			parent: actualParent,
			id:     actualResourceID,
		}, nil
	}

	// If no externalRef, construct identity from spec
	return &ZoneIdentity{
		parent: parentFromSpec,
		id:     resourceID,
	}, nil
}

// ParseZoneExternal parses the external representation of a Zone name.
// Format: projects/{{project}}/locations/{{location}}/lakes/{{lake}}/zones/{{zone}}
func ParseZoneExternal(external string) (parent *ZoneParent, resourceID string, err error) {
	tokens := strings.Split(external, "/")
	if len(tokens) != 8 ||
		tokens[0] != "projects" ||
		tokens[2] != "locations" ||
		tokens[4] != "lakes" ||
		tokens[6] != "zones" {
		return nil, "", fmt.Errorf("format of DataplexZone external=%q was not known (use projects/{{project}}/locations/{{location}}/lakes/{{lake}}/zones/{{zone}})", external)
	}
	parent = &ZoneParent{
		ProjectID: tokens[1],
		Location:  tokens[3],
		LakeID:    tokens[5],
	}
	resourceID = tokens[7]
	return parent, resourceID, nil
}
