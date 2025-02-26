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

package v1alpha1

import (
	"context"
	"fmt"
	"strings"

	apigeev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/apigee/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	ApigeeEndpointAttachmentIDToken  = "endpointAttachments"
	ApigeeEndpointAttachmentIDFormat = apigeev1beta1.ApigeeOrganizationIDFormat + "/" + ApigeeEndpointAttachmentIDToken + "/{{attachmentID}}"
)

var _ identity.Identity = &ApigeeEndpointAttachmentIdentity{}

type ApigeeEndpointAttachmentIdentity struct {
	ParentID   *apigeev1beta1.ApigeeOrganizationIdentity
	ResourceID string
}

func (i *ApigeeEndpointAttachmentIdentity) String() string {
	return i.ParentID.String() + "/" + ApigeeEndpointAttachmentIDToken + "/" + i.ResourceID
}

func (i *ApigeeEndpointAttachmentIdentity) FromExternal(ref string) error {
	requiredTokens := len(strings.Split(ApigeeEndpointAttachmentIDFormat, "/"))

	tokens := strings.Split(ref, "/")
	if len(tokens) != requiredTokens || tokens[len(tokens)-2] != ApigeeEndpointAttachmentIDToken {
		return fmt.Errorf("format of ApigeeEndpointAttachment ref=%q was not known (use %q)", ref, ApigeeEndpointAttachmentIDFormat)
	}

	parentID := &apigeev1beta1.ApigeeOrganizationIdentity{}
	if err := parentID.FromExternal(strings.Join(tokens[:len(tokens)-2], "/")); err != nil {
		return fmt.Errorf("format of ApigeeEndpointAttachment ref=%q was not known (use %q)", ref, ApigeeEndpointAttachmentIDFormat)
	}

	resourceID := tokens[len(tokens)-1]

	i.ParentID = parentID
	i.ResourceID = resourceID

	return nil
}

var _ identity.Resource = &ApigeeEndpointAttachment{}

func (obj *ApigeeEndpointAttachment) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
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

	id := &ApigeeEndpointAttachmentIdentity{
		ParentID:   parentID.(*apigeev1beta1.ApigeeOrganizationIdentity),
		ResourceID: resourceID,
	}

	// Attempt to ensure ID is immutable, by verifying against previously-set `status.externalRef`.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		previousID := &ApigeeEndpointAttachmentIdentity{}
		if err := previousID.FromExternal(externalRef); err != nil {
			return nil, err
		}
		if id.String() != previousID.String() {
			return nil, fmt.Errorf("cannot update ApigeeEndpointAttachment identity (old=%q, new=%q): identity is immutable", previousID.String(), id.String())
		}
	}

	return id, nil
}

func (obj *ApigeeEndpointAttachment) GetParentIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	// Normalize parent reference
	if err := obj.Spec.OrganizationRef.Normalize(ctx, reader, obj.GetNamespace()); err != nil {
		return nil, err
	}
	// Get parent identity
	parentID := &apigeev1beta1.ApigeeOrganizationIdentity{}
	if err := parentID.FromExternal(obj.Spec.OrganizationRef.External); err != nil {
		return nil, err
	}
	return parentID, nil
}
