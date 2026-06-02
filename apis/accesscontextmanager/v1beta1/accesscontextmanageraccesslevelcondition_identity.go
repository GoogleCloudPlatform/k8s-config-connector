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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &AccessContextManagerAccessLevelConditionIdentity{}
	_ identity.Resource   = &AccessContextManagerAccessLevelCondition{}

	// AccessContextManagerAccessLevelConditionIdentityURL is the format for the externalRef of an AccessContextManagerAccessLevelCondition.
	// Since there is no direct asset name for a condition in GCP Asset Inventory, we use a synthetic format
	// based on the parent AccessLevel.
	AccessContextManagerAccessLevelConditionIdentityURL = "accessPolicies/{accessPolicy}/accessLevels/{accessLevel}/condition"

	accessContextManagerAccessLevelConditionIdentityFormat = gcpurls.Template[AccessContextManagerAccessLevelConditionIdentity](
		"accesscontextmanager.googleapis.com",
		AccessContextManagerAccessLevelConditionIdentityURL,
	)
)

// AccessContextManagerAccessLevelConditionIdentity represents the identity of an AccessContextManagerAccessLevelCondition.
// +k8s:deepcopy-gen=false
type AccessContextManagerAccessLevelConditionIdentity struct {
	AccessPolicy string
	AccessLevel  string
}

func (i *AccessContextManagerAccessLevelConditionIdentity) String() string {
	return accessContextManagerAccessLevelConditionIdentityFormat.ToString(*i)
}

func (i *AccessContextManagerAccessLevelConditionIdentity) FromExternal(ref string) error {
	parsed, match, err := accessContextManagerAccessLevelConditionIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of AccessContextManagerAccessLevelCondition external=%q was not known (use %s): %w", ref, accessContextManagerAccessLevelConditionIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of AccessContextManagerAccessLevelCondition external=%q was not known (use %s)", ref, accessContextManagerAccessLevelConditionIdentityFormat.CanonicalForm())
	}
	*i = *parsed
	return nil
}

func (i *AccessContextManagerAccessLevelConditionIdentity) Host() string {
	return accessContextManagerAccessLevelConditionIdentityFormat.Host()
}

func (obj *AccessContextManagerAccessLevelCondition) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	// Get parent AccessLevel identity
	if obj.Spec.AccessLevelRef == nil {
		return nil, fmt.Errorf("spec.accessLevelRef is required")
	}

	accessLevelId := &AccessLevelIdentity{}
	external := obj.Spec.AccessLevelRef.GetExternal()
	if external != "" {
		if err := accessLevelId.FromExternal(external); err != nil {
			return nil, fmt.Errorf("cannot parse spec.accessLevelRef.external=%q: %w", external, err)
		}
	} else {
		// If external is not provided, we would need to resolve the reference.
		// For identity, we usually expect the externalRef to be eventually populated or the spec to have enough info.
		// However, the direct controller will handle resolution during reconcile.
		// Here we try to get it from status if available.
		externalRef := common.ValueOf(obj.Status.ExternalRef)
		if externalRef != "" {
			id := &AccessContextManagerAccessLevelConditionIdentity{}
			if err := id.FromExternal(externalRef); err != nil {
				return nil, fmt.Errorf("cannot parse status.externalRef=%q: %w", externalRef, err)
			}
			return id, nil
		}
		return nil, fmt.Errorf("cannot determine identity: spec.accessLevelRef.external and status.externalRef are both unset")
	}

	return &AccessContextManagerAccessLevelConditionIdentity{
		AccessPolicy: accessLevelId.Parent,
		AccessLevel:  accessLevelId.AccessLevel,
	}, nil
}
