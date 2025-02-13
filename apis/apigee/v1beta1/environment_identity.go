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
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	EnvironmentIDToken  = "environments"
	EnvironmentIDFormat = ApigeeOrganizationIDFormat + "/" + EnvironmentIDToken + "/{{environmentID}}"
)

var _ identity.Identity = &ApigeeEnvironmentIdentity{}

type ApigeeEnvironmentIdentity struct {
	ParentID   *ApigeeOrganizationIdentity
	ResourceID string
}

func (i *ApigeeEnvironmentIdentity) String() string {
	return i.ParentID.String() + "/" + EnvironmentIDToken + "/" + i.ResourceID
}

func (i *ApigeeEnvironmentIdentity) FromExternal(ref string) error {
	requiredTokens := len(strings.Split(EnvironmentIDFormat, "/"))

	tokens := strings.Split(ref, "/")
	if len(tokens) != requiredTokens || tokens[len(tokens)-2] != EnvironmentIDToken {
		return fmt.Errorf("format of ApigeeEnvironment ref=%q was not known (use %q)", ref, EnvironmentIDFormat)
	}

	parentID := &ApigeeOrganizationIdentity{}
	if err := parentID.FromExternal(strings.Join(tokens[:len(tokens)-2], "/")); err != nil {
		return fmt.Errorf("format of ApigeeEnvironment ref=%q was not known (use %q)", ref, EnvironmentIDFormat)
	}

	resourceID := tokens[len(tokens)-1]

	i.ParentID = parentID
	i.ResourceID = resourceID

	return nil
}

var _ identity.Resource = &ApigeeEnvironment{}

func (obj *ApigeeEnvironment) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	// Get parent ID
	parentID, err := obj.GetParentIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}

	// Get resource ID
	resourceID := common.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	id := &ApigeeEnvironmentIdentity{
		ParentID:   parentID.(*ApigeeOrganizationIdentity),
		ResourceID: resourceID,
	}

	// Attempt to ensure ID is immutable, by verifying against previously-set `status.externalRef`.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		previousID := &ApigeeEnvironmentIdentity{}
		if err := previousID.FromExternal(externalRef); err != nil {
			return nil, err
		}
		if id.String() != previousID.String() {
			return nil, fmt.Errorf("cannot update ApigeeEnvironment identity (old=%q, new=%q): identity is immutable", previousID.String(), id.String())
		}
	}

	return id, nil
}

func (obj *ApigeeEnvironment) GetParentIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	// Normalize parent reference
	if err := obj.Spec.ApigeeOrganizationRef.Normalize(ctx, reader, obj.GetNamespace()); err != nil {
		return nil, err
	}
	// Get parent identity
	parentID := &ApigeeOrganizationIdentity{}
	if err := parentID.FromExternal(obj.Spec.ApigeeOrganizationRef.External); err != nil {
		return nil, err
	}
	return parentID, nil
}
