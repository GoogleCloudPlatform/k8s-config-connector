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

package v1alpha1

import (
	"context"
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// FloorSettingIdentity defines the resource reference to ModelArmorFloorSetting, which "External" field
// holds the GCP identifier for the KRM object.
type FloorSettingIdentity struct {
	parent *FloorSettingParent
}

func (i *FloorSettingIdentity) String() string {
	return i.parent.String() + "/floorSetting"
}

func (i *FloorSettingIdentity) Parent() *FloorSettingParent {
	return i.parent
}

type FloorSettingParent struct {
	ProjectID string
	Location  string
}

func (p *FloorSettingParent) String() string {
	return "projects/" + p.ProjectID + "/locations/" + p.Location
}

// NewFloorSettingIdentity builds a FloorSettingIdentity from the Config Connector FloorSetting object.
func NewFloorSettingIdentity(ctx context.Context, reader client.Reader, obj *ModelArmorFloorSetting) (*FloorSettingIdentity, error) {
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

	// Use approved External
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		actualParent, err := ParseFloorSettingExternal(externalRef)
		if err != nil {
			return nil, err
		}
		if actualParent.ProjectID != projectID {
			return nil, fmt.Errorf("spec.projectRef changed, expect %s, got %s", actualParent.ProjectID, projectID)
		}
		if actualParent.Location != location {
			return nil, fmt.Errorf("spec.location changed, expect %s, got %s", actualParent.Location, location)
		}
	}
	return &FloorSettingIdentity{
		parent: &FloorSettingParent{
			ProjectID: projectID,
			Location:  location,
		},
	}, nil
}

func ParseFloorSettingExternal(external string) (parent *FloorSettingParent, err error) {
	tokens := strings.Split(external, "/")
	if len(tokens) != 5 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "floorSetting" {
		return nil, fmt.Errorf("format of ModelArmorFloorSetting external=%q was not known (use projects/{{projectID}}/locations/{{location}}/floorSetting)", external)
	}
	parent = &FloorSettingParent{
		ProjectID: tokens[1],
		Location:  tokens[3],
	}
	return parent, nil
}
