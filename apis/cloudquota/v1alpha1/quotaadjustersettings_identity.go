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

// QuotaAdjusterSettingsIdentity defines the resource reference to APIQuotaAdjusterSettings, which "External" field
// holds the GCP identifier for the KRM object.
type QuotaAdjusterSettingsIdentity struct {
	parent *QuotaAdjusterSettingsParent
}

func (i *QuotaAdjusterSettingsIdentity) String() string {
	return i.parent.String() + "/quotaAdjusterSettings"
}

func (i *QuotaAdjusterSettingsIdentity) ID() string {
	return "quotaAdjusterSettings"
}

func (i *QuotaAdjusterSettingsIdentity) Parent() *QuotaAdjusterSettingsParent {
	return i.parent
}

// QuotaAdjusterSettingsParent defines the GCP project, location for the given QuotaAdjusterSettings resource.
type QuotaAdjusterSettingsParent struct {
	ProjectID string
}

func (p *QuotaAdjusterSettingsParent) String() string {
	return "projects/" + p.ProjectID + "/locations/global"
}

// New builds a QuotaAdjusterSettingsIdentity from the Config Connector QuotaAdjusterSettings object.
func NewQuotaAdjusterSettingsIdentity(ctx context.Context, reader client.Reader, obj *APIQuotaAdjusterSettings) (*QuotaAdjusterSettingsIdentity, error) {

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
	resourceID := common.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}
	if resourceID != "quotaAdjusterSettings" {
		return nil, fmt.Errorf("resourceID or .metadata.name can only be 'quotaAdjusterSettings'")
	}

	// Use approved External
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		actualParent, actualResourceID, err := ParseQuotaAdjusterSettingsExternal(externalRef)
		if err != nil {
			return nil, err
		}
		if actualParent.ProjectID != projectID {
			return nil, fmt.Errorf("spec.projectRef changed, expect %s, got %s", actualParent.ProjectID, projectID)
		}
		if actualResourceID != resourceID {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
				resourceID, actualResourceID)
		}
	}
	return &QuotaAdjusterSettingsIdentity{
		parent: &QuotaAdjusterSettingsParent{
			ProjectID: projectID,
		},
	}, nil
}

func ParseQuotaAdjusterSettingsExternal(external string) (parent *QuotaAdjusterSettingsParent, resourceID string, err error) {
	tokens := strings.Split(external, "/")
	if len(tokens) != 5 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[3] != "global" || tokens[4] != "quotaAdjusterSettings" {
		return nil, "", fmt.Errorf("format of APIQuotaAdjusterSettings external=%q was not known (use projects/{{projectID}}/locations/global/quotaAdjusterSettings)", external)
	}
	parent = &QuotaAdjusterSettingsParent{
		ProjectID: tokens[1],
	}
	resourceID = tokens[4]
	return parent, resourceID, nil
}
