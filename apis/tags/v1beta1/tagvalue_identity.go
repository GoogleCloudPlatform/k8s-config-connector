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
	// TagsTagValueIdentityURL is the format for the externalRef of a TagsTagValue.
	TagsTagValueIdentityURL = "tagValues/{{tagValue}}"
)

var _ identity.Identity = &TagsTagValueIdentity{}

// TagsTagValueIdentity represents the identity of a TagsTagValue.
// +k8s:deepcopy-gen=false
type TagsTagValueIdentity struct {
	TagValue string
}

func (i *TagsTagValueIdentity) String() string {
	return "tagValues/" + i.TagValue
}

func (i *TagsTagValueIdentity) FromExternal(ref string) error {
	// Should be able to parse https://docs.cloud.google.com/asset-inventory/docs/asset-names
	ref = strings.TrimPrefix(ref, "//cloudresourcemanager.googleapis.com/")

	tokens := strings.Split(ref, "/")
	if len(tokens) == 2 && tokens[0] == "tagValues" {
		i.TagValue = tokens[1]
		if i.TagValue == "" {
			return fmt.Errorf("tagValue was empty in external=%q", ref)
		}
		return nil
	}

	return fmt.Errorf("format of TagValue external=%q was not known (use %s)", ref, TagsTagValueIdentityURL)
}

var _ identity.Resource = &TagsTagValue{}

func (obj *TagsTagValue) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	// Get desired resource ID
	resourceID := common.ValueOf(obj.Spec.ResourceID)

	// Server-generated ID; do not fallback to name
	// if resourceID == "" {
	// 	resourceID = obj.GetName()
	// }

	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	newIdentity := &TagsTagValueIdentity{
		TagValue: resourceID,
	}

	// Validate against the ID stored in status.externalRef
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		statusIdentity := &TagsTagValueIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, fmt.Errorf("cannot parse existing externalRef=%q: %w", externalRef, err)
		}
		if statusIdentity.String() != newIdentity.String() {
			return nil, fmt.Errorf("existing externalRef=%q does not match the identity resolved from spec: %q", externalRef, newIdentity.String())
		}
	}
	return newIdentity, nil
}
