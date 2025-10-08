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

// AccountIdentity defines the resource reference to AnalyticsAccount, which "External" field
// holds the GCP identifier for the KRM object.
type AccountIdentity struct {
	id string
}

func (i *AccountIdentity) String() string {
	return "accounts/" + i.id
}

func (i *AccountIdentity) ID() string {
	return i.id
}

// New builds a AccountIdentity from the Config Connector Account object.
func NewAccountIdentity(ctx context.Context, reader client.Reader, obj *AnalyticsAccount) (*AccountIdentity, error) {
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
		actualResourceID, err := ParseAccountExternal(externalRef)
		if err != nil {
			return nil, err
		}
		if actualResourceID != resourceID {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
				resourceID, actualResourceID)
		}
	}
	return &AccountIdentity{
		id: resourceID,
	}, nil
}

func ParseAccountExternal(external string) (resourceID string, err error) {
	tokens := strings.Split(external, "/")
	if len(tokens) != 2 || tokens[0] != "accounts" {
		return "", fmt.Errorf("format of AnalyticsAccount external=%q was not known (use accounts/{{accountID}})", external)
	}
	resourceID = tokens[1]
	return resourceID, nil
}
