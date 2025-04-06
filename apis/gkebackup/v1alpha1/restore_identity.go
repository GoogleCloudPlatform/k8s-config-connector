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
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// RestoreIdentity defines the resource reference to GKEBackupRestore, which "External" field
// holds the GCP identifier for the KRM object.
type RestoreIdentity struct {
	parent *RestoreParent
	id     string
}

func (i *RestoreIdentity) String() string {
	return i.parent.String() + "/restores/" + i.id
}

func (i *RestoreIdentity) ID() string {
	return i.id
}

func (i *RestoreIdentity) Parent() *RestoreParent {
	return i.parent
}

type RestoreParent struct {
	RestorePlan string
}

func (p *RestoreParent) String() string {
	return p.RestorePlan
}

// New builds a RestoreIdentity from the Config Connector Restore object.
func NewRestoreIdentity(ctx context.Context, reader client.Reader, obj *GKEBackupRestore) (*RestoreIdentity, error) {
	// Get Parent
	restorePlanRef := obj.Spec.RestorePlanRef
	restorePlan, err := restorePlanRef.NormalizedExternal(ctx, reader, obj.GetNamespace())
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
		actualParent, actualResourceID, err := ParseRestoreExternal(externalRef)
		if err != nil {
			return nil, err
		}
		if actualParent.RestorePlan != restorePlan {
			return nil, fmt.Errorf("spec.restorePlanRef changed, expect %s, got %s", actualParent.RestorePlan, restorePlan)
		}
		if actualResourceID != resourceID {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
				resourceID, actualResourceID)
		}
	}
	return &RestoreIdentity{
		parent: &RestoreParent{
			RestorePlan: restorePlan,
		},
		id: resourceID,
	}, nil
}

func ParseRestoreExternal(external string) (parent *RestoreParent, resourceID string, err error) {
	tokens := strings.Split(external, "/")
	if len(tokens) != 8 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "restorePlans" || tokens[6] != "restores" {
		return nil, "", fmt.Errorf("format of GKEBackupRestore external=%q was not known (use projects/{{projectID}}/locations/{{location}}/restorePlans/{{restoreplanID}}/restores/{{restoreID}})", external)
	}
	restorePlan := strings.Join(tokens[:len(tokens)-2], "/")
	parent = &RestoreParent{
		RestorePlan: restorePlan,
	}
	resourceID = tokens[7]
	return parent, resourceID, nil
}
