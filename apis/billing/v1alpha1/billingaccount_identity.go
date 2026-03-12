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

var _ identity.Identity = &BillingAccountIdentity{}

type BillingAccountIdentity struct {
	BillingAccountID string
}

func (i *BillingAccountIdentity) String() string {
	return "billingAccounts/" + i.BillingAccountID
}

func (i *BillingAccountIdentity) FromExternal(ref string) error {
	tokens := strings.Split(ref, "/")
	if len(tokens) == 2 && tokens[0] == "billingAccounts" {
		i.BillingAccountID = tokens[1]
		return nil
	}

	return fmt.Errorf("format of BillingAccount ref=%q was not known (use %q)", ref, "billingAccounts/{billingAccountID}")
}

var _ identity.Resource = &BillingAccount{}

func (obj *BillingAccount) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {

	// Get resource ID
	resourceID := common.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	id := &BillingAccountIdentity{
		BillingAccountID: resourceID,
	}

	// Attempt to ensure ID is immutable, by verifying against previously-set `status.externalRef`.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		previousID := &BillingAccountIdentity{}
		if err := previousID.FromExternal(externalRef); err != nil {
			return nil, err
		}
		if id.String() != previousID.String() {
			return nil, fmt.Errorf("cannot update BillingAccount identity (old=%q, new=%q): identity is immutable", previousID.String(), id.String())
		}
	}

	return id, nil
}

func (obj *BillingAccount) GetParentIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	return nil, nil
}
