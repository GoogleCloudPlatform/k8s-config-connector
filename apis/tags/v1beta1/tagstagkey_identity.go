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
	// TagsTagKeyIdentityURL is the format for the externalRef of a TagsTagKey.
	TagsTagKeyIdentityURL = "tagKeys/{{tagKey}}"
	ServiceDomain         = "cloudresourcemanager.googleapis.com"
)

var _ identity.Identity = &TagsTagKeyIdentity{}

// TagsTagKeyIdentity represents the identity of a TagsTagKey.
// +k8s:deepcopy-gen=false
type TagsTagKeyIdentity struct {
	TagKey string
}

func (i *TagsTagKeyIdentity) String() string {
	return "tagKeys/" + i.TagKey
}

func (i *TagsTagKeyIdentity) FromExternal(ref string) error {
	// Should be able to parse https://docs.cloud.google.com/asset-inventory/docs/asset-names
	ref = strings.TrimPrefix(ref, "//cloudresourcemanager.googleapis.com/")

	tokens := strings.Split(ref, "/")
	if len(tokens) == 2 && tokens[0] == "tagKeys" {
		i.TagKey = tokens[1]
		if i.TagKey == "" {
			return fmt.Errorf("tagKey was empty in external=%q", ref)
		}
		return nil
	}

	return fmt.Errorf("format of TagKey external=%q was not known (use %s)", ref, TagsTagKeyIdentityURL)
}

var _ identity.Resource = &TagsTagKey{}

func (obj *TagsTagKey) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	// Get desired resource ID
	resourceID := common.ValueOf(obj.Spec.ResourceID)

	// Server-generated ID; do not fallback to name
	// if resourceID == "" {
	// 	resourceID = obj.GetName()
	// }

	var specIdentity *TagsTagKeyIdentity
	if resourceID != "" {
		specIdentity = &TagsTagKeyIdentity{
			TagKey: resourceID,
		}
	}

	// Validate against the ID stored in status.externalRef
	var statusIdentity *TagsTagKeyIdentity
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		statusIdentity = &TagsTagKeyIdentity{}
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
