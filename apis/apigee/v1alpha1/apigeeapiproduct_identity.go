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
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &ApigeeAPIProductIdentity{}
	_ identity.Resource   = &ApigeeAPIProduct{}
)

var ApigeeAPIProductIdentityFormat = gcpurls.Template[ApigeeAPIProductIdentity]("apigee.googleapis.com", "organizations/{organization}/apiproducts/{apiproduct}")

// +k8s:deepcopy-gen=false
type ApigeeAPIProductIdentity struct {
	Organization string
	Apiproduct   string
}

func (i *ApigeeAPIProductIdentity) String() string {
	return ApigeeAPIProductIdentityFormat.ToString(*i)
}

func (i *ApigeeAPIProductIdentity) FromExternal(ref string) error {
	parsed, match, err := ApigeeAPIProductIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of ApigeeAPIProduct external=%q was not known (use %s): %w", ref, ApigeeAPIProductIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of ApigeeAPIProduct external=%q was not known (use %s)", ref, ApigeeAPIProductIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *ApigeeAPIProductIdentity) Host() string {
	return ApigeeAPIProductIdentityFormat.Host()
}

func getIdentityFromApigeeAPIProductSpec(ctx context.Context, reader client.Reader, obj *ApigeeAPIProduct) (*ApigeeAPIProductIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	if obj.Spec.OrganizationRef == nil {
		return nil, fmt.Errorf("organization is required")
	}

	if err := obj.Spec.OrganizationRef.Normalize(ctx, reader, obj.GetNamespace()); err != nil {
		return nil, err
	}

	// External is in format "organizations/{organization}"
	orgExternal := obj.Spec.OrganizationRef.External
	tokens := strings.Split(orgExternal, "/")
	if len(tokens) != 2 || tokens[0] != "organizations" {
		return nil, fmt.Errorf("organization external reference %q is invalid", orgExternal)
	}

	identity := &ApigeeAPIProductIdentity{
		Organization: tokens[1],
		Apiproduct:   resourceID,
	}
	return identity, nil
}

func (obj *ApigeeAPIProduct) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromApigeeAPIProductSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &ApigeeAPIProductIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change ApigeeAPIProduct identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
