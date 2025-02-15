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

// InstanceAttachmentIdentity defines the resource reference to ApigeeInstanceAttachment, which "External" field
// holds the GCP identifier for the KRM object.
type InstanceAttachmentIdentity struct {
	parent *InstanceIdentity
	id     string
}

func (i *InstanceAttachmentIdentity) String() string {
	return i.parent.String() + "/attachments/" + i.id
}

func (i *InstanceAttachmentIdentity) ID() string {
	return i.id
}

func (i *InstanceAttachmentIdentity) Parent() *InstanceIdentity {
	return i.parent
}

// New builds a GoogleCloudApigeeV1InstanceAttachmentIdentity from the Config Connector GoogleCloudApigeeV1InstanceAttachment object.
func NewInstanceAttachmentIdentity(ctx context.Context, reader client.Reader, obj *ApigeeInstanceAttachment) (*InstanceAttachmentIdentity, error) {
	// Get Parent
	instanceExternal, err := obj.Spec.InstanceRef.NormalizedExternal(ctx, reader, obj.GetNamespace())
	if err != nil {
		return nil, err
	}
	parentInstanceID, err := NewInstanceIdentityFromNormalizedExternal(instanceExternal)
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
		actualParent, actualResourceID, err := ParseInstanceAttachmentExternal(externalRef)
		if err != nil {
			return nil, err
		}
		if actualParent.String() != parentInstanceID.String() {
			return nil, fmt.Errorf("parent instance (spec.instanceRef) changed, expected %s, got %s", actualParent.String(), parentInstanceID.String())
		}
		if actualResourceID != resourceID {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
				resourceID, actualResourceID)
		}
	}
	return &InstanceAttachmentIdentity{
		parent: parentInstanceID,
		id:     resourceID,
	}, nil
}

func ParseInstanceAttachmentExternal(external string) (parent *InstanceIdentity, resourceID string, err error) {
	tokens := strings.Split(external, "/")
	if len(tokens) != 6 || tokens[0] != "organizations" || tokens[2] != "instances" || tokens[4] != "attachments" {
		return nil, "", fmt.Errorf("format of ApigeeInstanceAttachment external=%q was not known (use organizations/{{organizationID}}/instances/{{instanceID}}/attachments/{{instanceattachmentID}})", external)
	}
	parent, err = NewInstanceIdentityFromNormalizedExternal(strings.Join(tokens[:4], "/"))
	if err != nil {
		return nil, "", err
	}
	resourceID = tokens[5]
	return parent, resourceID, nil
}
