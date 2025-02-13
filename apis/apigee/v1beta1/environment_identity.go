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
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	EnvironmentIDToken  = "environments"
	EnvironmentIDFormat = OrganizationIDFormat + "/" + EnvironmentIDToken + "/{{environmentID}}"
)

// OrganizationIdentity uniquely defines a ApigeeEnvironment object.
type EnvironmentIdentity struct {
	ParentID   *OrganizationIdentity
	ResourceID string
}

func (i *EnvironmentIdentity) String() string {
	return i.ParentID.String() + "/" + EnvironmentIDToken + "/" + i.ResourceID
}

// NewEnvironmentIdentity parses a string-format ApigeeEnvironment reference into a EnvironmentIdentity object.
func NewEnvironmentIdentity(ref string) (*EnvironmentIdentity, error) {
	requiredTokens := len(strings.Split(EnvironmentIDFormat, "/"))

	tokens := strings.Split(ref, "/")
	if len(tokens) != requiredTokens || tokens[len(tokens)-2] != EnvironmentIDToken {
		return nil, fmt.Errorf("format of ApigeeEnvironment ref=%q was not known (use %q)", ref, EnvironmentIDFormat)
	}

	parentID, err := NewOrganizationIdentity(strings.Join(tokens[:len(tokens)-2], "/"))
	if err != nil {
		return nil, fmt.Errorf("format of ApigeeEnvironment ref=%q was not known (use %q)", ref, EnvironmentIDFormat)
	}

	resourceID := tokens[len(tokens)-1]

	id := &EnvironmentIdentity{
		ParentID:   parentID,
		ResourceID: resourceID,
	}

	return id, nil
}

// GetIdentity reads the identity from the ApigeeEnvironment resource.
func (obj *ApigeeEnvironment) GetIdentity(ctx context.Context, reader client.Reader) (*EnvironmentIdentity, error) {
	// Normalize parent reference
	if err := obj.Spec.ApigeeOrganizationRef.Normalize(ctx, reader, obj.GetNamespace()); err != nil {
		return nil, err
	}
	// Get parent identity
	parentID, err := NewOrganizationIdentity(obj.Spec.ApigeeOrganizationRef.External)
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

	id := &EnvironmentIdentity{
		ParentID:   parentID,
		ResourceID: resourceID,
	}

	// Attempt to ensure ID is immutable, by verifying against previously-set `status.externalRef`.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		previousID, err := NewEnvironmentIdentity(externalRef)
		if err != nil {
			return nil, err
		}
		if id.String() != previousID.String() {
			return nil, fmt.Errorf("cannot update ApigeeEnvironment identity (old=%q, new=%q): identity is immutable", previousID.String(), id.String())
		}
	}

	return id, nil
}
