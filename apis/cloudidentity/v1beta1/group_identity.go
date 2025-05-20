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
	id string
}

func (i *GroupIdentity) String() string {
	return "groups/" + i.id
}

func (i *GroupIdentity) ID() string {
	return i.id
}

// NewGroupIdentity New builds a GroupIdentity from the Config Connector Group object.
func NewGroupIdentity(ctx context.Context, reader client.Reader, obj *CloudIdentityGroup) (*GroupIdentity, error) {
	// Get desired ID
	resourceID := common.ValueOf(obj.Spec.ResourceID)

	/*
	   Special handling for backward compatibility.
	   Ideally resourceID should be {groupID}, but due to "wrong" configuration in servicemappings:
	   https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/575f419d12d6967374e7d2fea9871757991060e4/config/servicemappings/cloudidentity.yaml#L31
	   the TF-based resource's resourceID is groups/{groupID}. Make sure both formats are accepted by direct controller.
	*/
	if id, err := ParseGroupExternal(resourceID); err == nil {
		resourceID = id
	}

	// Use approved External
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		actualResourceID, err := ParseGroupExternal(externalRef)
		if err != nil {
			return nil, err
		}
		if resourceID != "" && actualResourceID != resourceID {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
				resourceID, actualResourceID)
		}
		resourceID = actualResourceID
	}

	return &GroupIdentity{
		id: resourceID,
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
