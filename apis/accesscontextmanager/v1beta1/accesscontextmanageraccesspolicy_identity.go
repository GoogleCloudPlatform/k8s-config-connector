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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &AccessPolicyIdentityV2{}
	_ identity.Resource   = &AccessContextManagerAccessPolicy{}
)

// AccessContextManagerAccessPolicyIdentityFormat matches accessPolicies/{accessPolicy}
var AccessContextManagerAccessPolicyIdentityFormat = gcpurls.Template[AccessPolicyIdentityV2]("accesscontextmanager.googleapis.com", "accessPolicies/{accessPolicy}")

// +k8s:deepcopy-gen=false
type AccessPolicyIdentityV2 struct {
	AccessPolicy string
}

func (i *AccessPolicyIdentityV2) String() string {
	return AccessContextManagerAccessPolicyIdentityFormat.ToString(*i)
}

func (i *AccessPolicyIdentityV2) FromExternal(ref string) error {
	parsed, match, err := AccessContextManagerAccessPolicyIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of AccessContextManagerAccessPolicy external=%q was not known (use %s): %w", ref, AccessContextManagerAccessPolicyIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of AccessContextManagerAccessPolicy external=%q was not known (use %s)", ref, AccessContextManagerAccessPolicyIdentityFormat.CanonicalForm())
	}
	*i = *parsed
	return nil
}

func (i *AccessPolicyIdentityV2) Host() string {
	return AccessContextManagerAccessPolicyIdentityFormat.Host()
}

func getIdentityFromAccessContextManagerAccessPolicySpec(ctx context.Context, reader client.Reader, obj *AccessContextManagerAccessPolicy) (*AccessPolicyIdentityV2, error) {
	resourceID := common.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	identity := &AccessPolicyIdentityV2{
		AccessPolicy: resourceID,
	}
	return identity, nil
}

func (obj *AccessContextManagerAccessPolicy) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromAccessContextManagerAccessPolicySpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &AccessPolicyIdentityV2{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change AccessContextManagerAccessPolicy identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}

// AccessPolicyIdentity is the identity of an AccessContextManagerAccessPolicy.
type AccessPolicyIdentity struct {
	resourceID string
	title      string
}

func (i *AccessPolicyIdentity) String() string {
	// return "projects/" + i.ProjectID + "/accesspolicys/" + i.resourceID
	return "/accesspolicys/" + i.resourceID
}

func (i *AccessPolicyIdentity) ResourceID() string {
	return i.resourceID
}

func (i *AccessPolicyIdentity) Title() string {
	return i.title
}

// New builds an AccessPolicyIdentity from the Config Connector AccessPolicy object.
func NewAccessPolicyIdentity(ctx context.Context, reader client.Reader, obj *AccessContextManagerAccessPolicy) (*AccessPolicyIdentity, error) {
	return &AccessPolicyIdentity{
		resourceID: *obj.Spec.ResourceID,
		title:      *obj.Spec.Title,
	}, nil
}

func ParseAccessPolicyExternal(external string) (resourceID string, err error) {
	// pattern: "accessPolicies/{access_policy}"
	tokens := strings.Split(external, "/")
	if len(tokens) != 2 || tokens[0] != "accesspolicys" {
		return "", fmt.Errorf("format of AccessContextManagerAccessPolicy external=%q was not known (use accessPolicies/{{access_policy}})", external)
	}
	resourceID = tokens[1]
	return resourceID, nil
}
