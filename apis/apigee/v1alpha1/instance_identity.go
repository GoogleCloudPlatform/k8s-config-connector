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

package v1alpha1

import (
	"context"
	"fmt"
	"strings"

	apigeev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/apigee/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"

	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	InstanceIDToken  = "instances"
	InstanceIDFormat = apigeev1beta1.OrganizationIDFormat + "/" + InstanceIDToken + "/{{instanceID}}"
)

// InstanceIdentity uniquely defines a ApigeeInstance object.
type InstanceIdentity struct {
	ParentID   *apigeev1beta1.OrganizationIdentity
	ResourceID string
}

func (i *InstanceIdentity) String() string {
	return i.ParentID.String() + "/" + InstanceIDToken + "/" + i.ResourceID
}

// NewInstanceIdentity parses a string-format ApigeeInstance reference into a InstanceIdentity object.
func NewInstanceIdentity(ref string) (*InstanceIdentity, error) {
	requiredTokens := len(strings.Split(InstanceIDFormat, "/"))

	tokens := strings.Split(ref, "/")
	if len(tokens) != requiredTokens || tokens[len(tokens)-2] != InstanceIDToken {
		return nil, fmt.Errorf("format of ApigeeInstance ref=%q was not known (use %q)", ref, InstanceIDFormat)
	}

	parentID, err := apigeev1beta1.NewOrganizationIdentity(strings.Join(tokens[:len(tokens)-2], "/"))
	if err != nil {
		return nil, fmt.Errorf("format of ApigeeInstance ref=%q was not known (use %q)", ref, InstanceIDFormat)
	}

	resourceID := tokens[len(tokens)-1]

	id := &InstanceIdentity{
		ParentID:   parentID,
		ResourceID: resourceID,
	}

	return id, nil
}

// GetIdentity reads the identity from the ApigeeInstance resource.
func (obj *ApigeeInstance) GetIdentity(ctx context.Context, reader client.Reader) (*InstanceIdentity, error) {
	// Normalize parent reference
	if err := obj.Spec.OrganizationRef.Normalize(ctx, reader, obj.GetNamespace()); err != nil {
		return nil, err
	}
	// Get parent identity
	parentID, err := apigeev1beta1.NewOrganizationIdentity(obj.Spec.OrganizationRef.External)
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

	id := &InstanceIdentity{
		ParentID:   parentID,
		ResourceID: resourceID,
	}

	// Attempt to ensure ID is immutable, by verifying against previously-set `status.externalRef`.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		previousID, err := apigeev1beta1.NewInstanceIdentity(externalRef)
		if err != nil {
			return nil, err
		}
		if id.String() != previousID.String() {
			return nil, fmt.Errorf("cannot update ApigeeInstance identity (old=%q, new=%q): identity is immutable", previousID.String(), id.String())
		}
	}

	return id, nil
}
