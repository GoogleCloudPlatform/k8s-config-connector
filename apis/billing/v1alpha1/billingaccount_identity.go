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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &BillingAccountIdentity{}
	_ identity.Resource   = &BillingAccount{}
)

var BillingAccountIdentityFormat = gcpurls.Template[BillingAccountIdentity]("cloudbilling.googleapis.com", "billingAccounts/{billingAccount}")

// BillingAccountIdentity is the identity of a GCP BillingAccount resource.
// +k8s:deepcopy-gen=false
type BillingAccountIdentity struct {
	BillingAccount string
}

func (i *BillingAccountIdentity) String() string {
	return BillingAccountIdentityFormat.ToString(*i)
}

func (i *BillingAccountIdentity) FromExternal(ref string) error {
	if ref == "" {
		return fmt.Errorf("BillingAccount external reference cannot be empty")
	}

	if !strings.Contains(ref, "/") {
		i.BillingAccount = ref
		return nil
	}

	parsed, match, err := BillingAccountIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of BillingAccount external=%q was not known (use %s): %w", ref, BillingAccountIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of BillingAccount external=%q was not known (use %s)", ref, BillingAccountIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *BillingAccountIdentity) Host() string {
	return BillingAccountIdentityFormat.Host()
}

func getIdentityFromBillingAccountSpec(ctx context.Context, reader client.Reader, obj *BillingAccount) (*BillingAccountIdentity, error) {
	resourceID := common.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	identity := &BillingAccountIdentity{
		BillingAccount: resourceID,
	}
	return identity, nil
}

func (obj *BillingAccount) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromBillingAccountSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		statusIdentity := &BillingAccountIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot update BillingAccount identity (old=%q, new=%q): identity is immutable", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
