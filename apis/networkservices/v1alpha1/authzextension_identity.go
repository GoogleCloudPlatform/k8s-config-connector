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
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ identity.Identity = &AuthzExtensionIdentity{}

// AuthzExtensionIdentity is the identity of a NetworkServicesAuthzExtension.
type AuthzExtensionIdentity struct {
	parent *parent.ProjectAndLocationParent
	id     string
}

func (i *AuthzExtensionIdentity) String() string {
	return i.parent.String() + "/authzExtensions/" + i.id
}

func (i *AuthzExtensionIdentity) ID() string {
	return i.id
}

func (i *AuthzExtensionIdentity) Parent() *parent.ProjectAndLocationParent {
	return i.parent
}

func (i *AuthzExtensionIdentity) FromExternal(ref string) error {
	tokens := strings.Split(ref, "/authzExtensions/")
	if len(tokens) != 2 {
		return fmt.Errorf("format of NetworkServicesAuthzExtension external=%q was not known (use projects/{{projectID}}/locations/{{location}}/authzExtensions/{{authzExtensionID}})", ref)
	}
	i.parent = &parent.ProjectAndLocationParent{}
	if err := i.parent.FromExternal(tokens[0]); err != nil {
		return err
	}
	i.id = tokens[1]
	if i.id == "" {
		return fmt.Errorf("authzExtensionID was empty in external=%q", ref)
	}
	return nil
}

var _ identity.Resource = &NetworkServicesAuthzExtension{}

func (obj *NetworkServicesAuthzExtension) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	id := &AuthzExtensionIdentity{
		parent: &parent.ProjectAndLocationParent{},
	}

	// Resolve user-configured Parent
	if err := obj.Spec.ProjectAndLocationRef.Build(ctx, reader, obj.GetNamespace(), id.parent); err != nil {
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
	id.id = resourceID

	// Use approved External
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		statusIdentity := &AuthzExtensionIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, fmt.Errorf("cannot parse existing externalRef=%q: %w", externalRef, err)
		}
		if statusIdentity.String() != id.String() {
			return nil, fmt.Errorf("existing externalRef=%q does not match the identity resolved from spec: %q", externalRef, id.String())
		}
	}
	return id, nil
}

// NewAuthzExtensionIdentity builds an AuthzExtensionIdentity from the Config Connector AuthzExtension object.
func NewAuthzExtensionIdentity(ctx context.Context, reader client.Reader, obj *NetworkServicesAuthzExtension) (*AuthzExtensionIdentity, error) {
	id := &AuthzExtensionIdentity{
		parent: &parent.ProjectAndLocationParent{},
	}

	// Resolve user-configured Parent
	if err := obj.Spec.ProjectAndLocationRef.Build(ctx, reader, obj.GetNamespace(), id.parent); err != nil {
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
	id.id = resourceID

	return id, nil
}
