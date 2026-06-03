// Copyright 2026 Google LLC
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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &AccountIdentity{}
	_ identity.Resource   = &AnalyticsAccount{}
)

var AccountIdentityFormat = gcpurls.Template[AccountIdentity]("analyticsadmin.googleapis.com", "accounts/{account}")

// +k8s:deepcopy-gen=false
type AccountIdentity struct {
	Account string
}

func (i *AccountIdentity) String() string {
	return AccountIdentityFormat.ToString(*i)
}

func (i *AccountIdentity) FromExternal(ref string) error {
	parsed, match, err := AccountIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of AnalyticsAccount external=%q was not known (use %s): %w", ref, AccountIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of AnalyticsAccount external=%q was not known (use %s)", ref, AccountIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *AccountIdentity) Host() string {
	return AccountIdentityFormat.Host()
}

func (i *AccountIdentity) ExternalIdentifier() *string {
	return &i.Account
}

func (i *AccountIdentity) ID() string {
	return i.Account
}

func (i *AccountIdentity) SetID(id string) {
	i.Account = id
}

func getIdentityFromAnalyticsAccountSpec(ctx context.Context, reader client.Reader, obj *AnalyticsAccount) (*AccountIdentity, error) {
	resourceID, err := refsv1beta1.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	identity := &AccountIdentity{
		Account: resourceID,
	}
	return identity, nil
}

func (obj *AnalyticsAccount) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromAnalyticsAccountSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &AccountIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		// Account IDs are service-generated, so we only validate if the spec has an explicit resourceID.
		if obj.Spec.ResourceID != nil && statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change AnalyticsAccount identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
		return statusIdentity, nil
	}

	return specIdentity, nil
}
