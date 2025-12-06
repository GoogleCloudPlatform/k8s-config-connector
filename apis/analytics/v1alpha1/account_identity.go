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
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"

	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ identity.Identity = &AccountIdentity{}

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

func (i *AccountIdentity) SetID(id string) {
	i.id = id
	return
}

func (i *AccountIdentity) FromExternal(ref string) error {
	tokens := strings.Split(ref, "/")
	if len(tokens) != 2 || tokens[0] != "accounts" {
		return fmt.Errorf("format of AnalyticsAccount external=%q was not known (use accounts/{{accountID}})", ref)
	}
	i.id = tokens[1]
	if i.id == "" {
		return fmt.Errorf("accountID was empty in external=%q", ref)
	}
	return nil
}

var _ identity.Resource = &AnalyticsAccount{}

func (obj *AnalyticsAccount) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	newIdentity := &AccountIdentity{}

	// Attempt to get the service-generated resource ID.
	newIdentity.id = common.ValueOf(obj.Spec.ResourceID)
	if newIdentity.id == "" && obj.Status.ExternalRef != nil { // Reconciliation after creation is completed.
		err := newIdentity.FromExternal(common.ValueOf(obj.Status.ExternalRef))
		if err != nil {
			return nil, err
		}
	}

	// Attempt to ensure ID is immutable, by verifying against previously-set `status.externalRef`.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		if newIdentity.String() != externalRef {
			return nil, fmt.Errorf("cannot update AnalyticsAccount identity (old=%q, new=%q): identity is immutable", externalRef, newIdentity.String())
		}
	}

	return newIdentity, nil
}
