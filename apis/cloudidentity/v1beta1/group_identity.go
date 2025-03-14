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
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// GroupIdentity defines the resource reference to CloudIdentityGroup, which "External" field
// holds the GCP identifier for the KRM object.
type GroupIdentity struct {
	id                      string
	serviceGeneratedIDKnown bool
}

// HasKnownID tells whether Config Connector knows the resource identity.
// If not, Config Connector saves one GCP GET call, and starts the CREATE call directly.
// This is mostly for GCP services that do not allow user to specify ID, but assign an ID when creating the object.
func (i *GroupIdentity) HasKnownID() bool {
	return i.serviceGeneratedIDKnown
}

func (i *GroupIdentity) String() string {
	return "groups/" + i.id
}

func (i *GroupIdentity) ID() string {
	return i.id
}

// NewGroupIdentity New builds a GroupIdentity from the Config Connector Group object.
func NewGroupIdentity(ctx context.Context, reader client.Reader, obj *CloudIdentityGroup) (*GroupIdentity, error) {
	known := false
	// Get desired ID
	desiredResourceID := common.ValueOf(obj.Spec.ResourceID)
	if desiredResourceID != "" {
		known = true
	}

	// Use approved External
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		known = true
		// Validate desired with actual
		actualResourceID, err := ParseGroupExternal(externalRef)
		if err != nil {
			return nil, err
		}
		if desiredResourceID != "" && actualResourceID != desiredResourceID {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
				desiredResourceID, actualResourceID)
		}
		desiredResourceID = actualResourceID
	}

	return &GroupIdentity{
		id:                      desiredResourceID,
		serviceGeneratedIDKnown: known,
	}, nil
}

func ParseGroupExternal(external string) (resourceID string, err error) {
	tokens := strings.Split(external, "/")
	if len(tokens) != 2 || tokens[0] != "groups" {
		return "", fmt.Errorf("format of CloudIdentityGroup external=%q was not known (use groups/{{groupID}})", external)
	}
	resourceID = tokens[1]
	return resourceID, nil
}
