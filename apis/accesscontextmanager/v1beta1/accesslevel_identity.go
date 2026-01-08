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

package v1beta1

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	util "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util/identity"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	// AccessLevelIdentityURL is the format for the externalRef of an AccessPolicy.
	// This format is used in the AccessContextManager API (and is different.
	// See: https://docs.cloud.google.com/asset-inventory/docs/asset-names
	AccessLevelIdentityURL = "accessPolicies/{{accessPolicyID}}/accessLevels/{{accessLevel}}"
)

var _ identity.Identity = &AccessLevelIdentity{}
var _ identity.Resource = &AccessContextManagerAccessLevel{}

var parser = regexp.MustCompile(`((//)?accesscontextmanager.googleapis.com)?/?accessPolicies/(?P<accessPolicies>[[:alpha:]]+)/accessLevels/(?P<accessLevels>[[:alpha:]]+)`)

// AccessLevelIdentity represents the identity of an accessLevel.
// +k8s:deepcopy-gen=false
type AccessLevelIdentity struct {
	Parent      string
	AccessLevel string
}

func (i *AccessLevelIdentity) String() string {
	return "accessPolicies/" + i.Parent + "/accessLevels/" + i.AccessLevel
}

func (i *AccessLevelIdentity) FromExternal(ref string) error {
	// TODO: Should be able to parse https://docs.cloud.google.com/asset-inventory/docs/asset-names
	// But that format is //cloudresourcemanager.googleapis.com/accessPolicieis/ACCESS_POLICY_ID/accessLevels/ACCESS_LEVEL
	// which is not the format used by the service.

	err, identityMap := util.ParseIdentityMap(ref, parser, 2)
	if err != nil {
		return fmt.Errorf("format of AccessLevel external=%q was not known (use %s): %w", ref, AccessLevelIdentityURL, err)
	}

	i.Parent = identityMap["accessPolicies"]
	i.AccessLevel = identityMap["accessLevels"]

	return nil
}

func (obj *AccessContextManagerAccessLevel) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	// Get desired resource ID
	resourceID := common.ValueOf(obj.Spec.ResourceID)

	// Server-generated ID; do not fallback to name
	// if resourceID == "" {
	// 	resourceID = obj.GetName()
	// }

	var specIdentity *AccessLevelIdentity
	if resourceID != "" {
		specIdentity = &AccessLevelIdentity{}
		if !strings.HasPrefix(resourceID, "accessPolicies/") {
			resourceID = "accessPolicies/" + resourceID
		}
		if err := specIdentity.FromExternal(resourceID); err != nil {
			return nil, fmt.Errorf("cannot parse spec.resourceID=%q: %w", resourceID, err)
		}
	}

	// Validate against the ID stored in status.externalRef
	var statusIdentity *AccessLevelIdentity
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		statusIdentity = &AccessLevelIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, fmt.Errorf("cannot parse existing externalRef=%q: %w", externalRef, err)
		}
	}

	if specIdentity != nil {
		if statusIdentity != nil && statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("existing externalRef=%q does not match the identity resolved from spec: %q", externalRef, specIdentity.String())
		}
		return specIdentity, nil
	}

	if statusIdentity != nil {
		return statusIdentity, nil
	}

	return nil, fmt.Errorf("cannot determine identity: spec.resourceID and status.externalRef are both unset")
}
