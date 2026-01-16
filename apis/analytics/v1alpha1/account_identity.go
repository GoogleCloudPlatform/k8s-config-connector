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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
)

// AccountIdentity defines the resource reference to AnalyticsAccount, which "External" field
// holds the GCP identifier for the KRM object.
// +k8s:deepcopy-gen=false
type AccountIdentity struct {
	id string
}

func (i *AccountIdentity) String() string {
	return "accounts/" + i.id
}

func (i *AccountIdentity) ID() string {
	return i.id
}

func (i *AccountIdentity) SetID(id string) {
	i.id = id
	return
}

// New builds a AccountIdentity from the Config Connector Account object.
func NewAccountIdentity(ctx context.Context, reader client.Reader, obj *AnalyticsAccount) (*AccountIdentity, error) {
	// Attempt to get the service-generated resource ID.
	resourceID := common.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" && obj.Status.ExternalRef != nil { // Reconciliation after creation is completed.
		savedResourceID, err := ParseAccountExternal(common.ValueOf(obj.Status.ExternalRef))
		if err != nil {
			return nil, err
		}
		resourceID = savedResourceID
	}

	id := &AccountIdentity{
		id: resourceID,
	}

	// Attempt to ensure ID is immutable, by verifying against previously-set `status.externalRef`.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		if id.String() != externalRef {
			return nil, fmt.Errorf("cannot update AnalyticsAccount identity (old=%q, new=%q): identity is immutable", externalRef, id.String())
		}
	}

	return id, nil
}

func ParseAccountExternal(external string) (resourceID string, err error) {
	tokens := strings.Split(external, "/")
	if len(tokens) != 2 || tokens[0] != "accounts" {
		return "", fmt.Errorf("format of AnalyticsAccount external=%q was not known (use accounts/{{accountID}})", external)
	}
	resourceID = tokens[1]
	return resourceID, nil
}

var _ identity.Identity = &AccountIdentity{}
