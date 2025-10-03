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
	ApigeeEnvgroupAttachmentIDToken  = "attachments"
	ApigeeEnvgroupAttachmentIDFormat = ApigeeEnvgroupIDFormat + "/" + ApigeeEnvgroupAttachmentIDToken + "/{{attachmentID}}"
)

var _ identity.Identity = &ApigeeEnvgroupAttachmentIdentity{}

type ApigeeEnvgroupAttachmentIdentity struct {
	ParentID   *ApigeeEnvgroupIdentity
	ResourceID string
}

func (i *ApigeeEnvgroupAttachmentIdentity) String() string {
	if i.ResourceID == "" {
		// Initially, the identity of the ApigeeEnvgroupAttachment is unknown, until it is acquired or created.
		// This is because the resource uses a service-generated ID.
		return ""
	}
	return i.ParentID.String() + "/" + ApigeeEnvgroupAttachmentIDToken + "/" + i.ResourceID
}

func (i *ApigeeEnvgroupAttachmentIdentity) FromExternal(ref string) error {
	requiredTokens := len(strings.Split(ApigeeEnvgroupAttachmentIDFormat, "/"))

	tokens := strings.Split(ref, "/")
	if len(tokens) != requiredTokens || tokens[len(tokens)-2] != ApigeeEnvgroupAttachmentIDToken {
		return fmt.Errorf("format of ApigeeEnvgroupAttachment ref=%q was not known (use %q)", ref, ApigeeEnvgroupAttachmentIDFormat)
	}

	parentID := &ApigeeEnvgroupIdentity{}
	if err := parentID.FromExternal(strings.Join(tokens[:len(tokens)-2], "/")); err != nil {
		return fmt.Errorf("format of ApigeeEnvgroupAttachment ref=%q was not known (use %q)", ref, ApigeeEnvgroupAttachmentIDFormat)
	}

	resourceID := tokens[len(tokens)-1]

	i.ParentID = parentID
	i.ResourceID = resourceID

	return nil
}

var _ identity.Resource = &ApigeeEnvgroupAttachment{}

func (obj *ApigeeEnvgroupAttachment) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	// Get parent ID
	parentID, err := obj.GetParentIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}

	// Get service-generated resource ID
	resourceID := common.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" && obj.Status.ExternalRef != nil {
		savedID := &ApigeeEnvgroupAttachmentIdentity{}
		if err := savedID.FromExternal(common.ValueOf(obj.Status.ExternalRef)); err != nil {
			return nil, err
		}
		resourceID = savedID.ResourceID
	}

	id := &ApigeeEnvgroupAttachmentIdentity{
		ParentID:   parentID.(*ApigeeEnvgroupIdentity),
		ResourceID: resourceID,
	}

	// Attempt to ensure ID is immutable, by verifying against previously-set `status.externalRef`.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		previousID := &ApigeeEnvgroupAttachmentIdentity{}
		if err := previousID.FromExternal(externalRef); err != nil {
			return nil, err
		}
		if id.String() != previousID.String() {
			return nil, fmt.Errorf("cannot update ApigeeEnvgroupAttachment identity (old=%q, new=%q): identity is immutable", previousID.String(), id.String())
		}
	}

	return id, nil
}

func (obj *ApigeeEnvgroupAttachment) GetParentIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	// Normalize parent reference
	if err := obj.Spec.EnvgroupRef.Normalize(ctx, reader, obj.GetNamespace()); err != nil {
		return nil, err
	}
	// Get parent identity
	parentID := &ApigeeEnvgroupIdentity{}
	if err := parentID.FromExternal(obj.Spec.EnvgroupRef.External); err != nil {
		return nil, err
	}
	return parentID, nil
}
