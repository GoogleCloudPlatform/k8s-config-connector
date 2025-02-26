// Copyright 2024 Google LLC
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
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	ApigeeEnvgroupIDToken  = "envgroups"
	ApigeeEnvgroupIDFormat = ApigeeOrganizationIDFormat + "/" + ApigeeEnvgroupIDToken + "/{{envgroupID}}"
)

var _ identity.Identity = &ApigeeEnvgroupIdentity{}

type ApigeeEnvgroupIdentity struct {
	ParentID   *ApigeeOrganizationIdentity
	ResourceID string
}

func (i *ApigeeEnvgroupIdentity) String() string {
	return i.ParentID.String() + "/" + ApigeeEnvgroupIDToken + "/" + i.ResourceID
}

func (i *ApigeeEnvgroupIdentity) FromExternal(ref string) error {
	requiredTokens := len(strings.Split(ApigeeEnvgroupIDFormat, "/"))

	tokens := strings.Split(ref, "/")
	if len(tokens) != requiredTokens || tokens[len(tokens)-2] != ApigeeEnvgroupIDToken {
		return fmt.Errorf("format of ApigeeEnvgroup ref=%q was not known (use %q)", ref, ApigeeEnvgroupIDFormat)
	}

	parentID := &ApigeeOrganizationIdentity{}
	if err := parentID.FromExternal(strings.Join(tokens[:len(tokens)-2], "/")); err != nil {
		return fmt.Errorf("format of ApigeeEnvgroup ref=%q was not known (use %q)", ref, ApigeeEnvgroupIDFormat)
	}

	resourceID := tokens[len(tokens)-1]

	i.ParentID = parentID
	i.ResourceID = resourceID

	return nil
}

var _ identity.Resource = &ApigeeEnvgroup{}

func (obj *ApigeeEnvgroup) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	// Get parent ID
	parentID, err := obj.GetParentIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}

	// Get desired ID
	resourceID := common.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	id := &ApigeeEnvgroupIdentity{
		ParentID:   parentID.(*ApigeeOrganizationIdentity),
		ResourceID: resourceID,
	}

	// Attempt to ensure ID is immutable, by verifying against previously-set `status.externalRef`.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		previousID := &ApigeeEnvgroupIdentity{}
		if err := previousID.FromExternal(externalRef); err != nil {
			return nil, err
		}
		if id.String() != previousID.String() {
			return nil, fmt.Errorf("cannot update ApigeeEnvgroup identity (old=%q, new=%q): identity is immutable", previousID.String(), id.String())
		}
	}

	return id, nil
}

func (obj *ApigeeEnvgroup) GetParentIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	// Normalize parent reference
	if err := obj.Spec.OrganizationRef.Normalize(ctx, reader, obj.GetNamespace()); err != nil {
		return nil, err
	}
	// Get parent identity
	parentID := &ApigeeOrganizationIdentity{}
	if err := parentID.FromExternal(obj.Spec.OrganizationRef.External); err != nil {
		return nil, err
	}
	return parentID, nil
}
