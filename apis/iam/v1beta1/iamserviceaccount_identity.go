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

package v1beta1

import (
	"context"
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &IAMServiceAccountIdentity{}
	_ identity.Resource   = &IAMServiceAccount{}
)

var IAMServiceAccountIdentityFormat = gcpurls.Template[IAMServiceAccountIdentity](
	"iam.googleapis.com",
	"projects/{project}/serviceAccounts/{account}",
)

// IAMServiceAccountIdentity is the identity of a GCP IAMServiceAccount resource.
// +k8s:deepcopy-gen=false
type IAMServiceAccountIdentity struct {
	Project string
	Account string
}

func (i *IAMServiceAccountIdentity) String() string {
	return IAMServiceAccountIdentityFormat.ToString(*i)
}

func (i *IAMServiceAccountIdentity) FromExternal(ref string) error {
	parsed, match, err := IAMServiceAccountIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of IAMServiceAccount external=%q was not known (use %s): %w", ref, IAMServiceAccountIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of IAMServiceAccount external=%q was not known (use %s)", ref, IAMServiceAccountIdentityFormat.CanonicalForm())
	}

	if strings.Contains(parsed.Account, "@") {
		return fmt.Errorf("format of IAMServiceAccount external=%q was not known (use %s): email format is not allowed in identity", ref, IAMServiceAccountIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *IAMServiceAccountIdentity) Host() string {
	return IAMServiceAccountIdentityFormat.Host()
}

func (i *IAMServiceAccountIdentity) ParentString() string {
	return "projects/" + i.Project
}

func getIdentityFromIAMServiceAccountSpec(ctx context.Context, reader client.Reader, obj *IAMServiceAccount) (*IAMServiceAccountIdentity, error) {
	resourceID := common.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project: %w", err)
	}

	identity := &IAMServiceAccountIdentity{
		Project: projectID,
		Account: resourceID,
	}
	return identity, nil
}

func (obj *IAMServiceAccount) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromIAMServiceAccountSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status externalRef, if present.
	statusExternalRef := common.ValueOf(obj.Status.ExternalRef)
	if statusExternalRef != "" {
		parsed, match, err := IAMServiceAccountIdentityFormat.Parse(statusExternalRef)
		if err != nil || !match {
			return nil, fmt.Errorf("format of IAMServiceAccount status.externalRef=%q was not known (use %s)", statusExternalRef, IAMServiceAccountIdentityFormat.CanonicalForm())
		}
		// Do NOT allow email suffix in status.externalRef. Compare directly without stripping.
		if strings.Contains(parsed.Account, "@") {
			return nil, fmt.Errorf("format of IAMServiceAccount status.externalRef=%q was not known (use %s): email format is not allowed in identity", statusExternalRef, IAMServiceAccountIdentityFormat.CanonicalForm())
		}
		statusIdentity := &IAMServiceAccountIdentity{
			Project: parsed.Project,
			Account: parsed.Account,
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change IAMServiceAccount identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
