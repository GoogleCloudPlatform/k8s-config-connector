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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	// TagsLocationTagBindingIdentityURL is the format for the externalRef of a TagsLocationTagBinding.
	TagsLocationTagBindingIdentityURL = "tagBindings/{{tagBinding}}/tagValues/{{tagValue}}"
)

var _ identity.Identity = &TagsLocationTagBindingIdentity{}

// TagsLocationTagBindingIdentity represents the identity of a TagsLocationTagBinding.
// +k8s:deepcopy-gen=false
type TagsLocationTagBindingIdentity struct {
	Parent   string
	TagValue string
}

func (i *TagsLocationTagBindingIdentity) String() string {
	return "tagBindings/" + i.Parent + "/tagValues/" + i.TagValue
}

func (i *TagsLocationTagBindingIdentity) FromExternal(ref string) error {
	// TODO: Should be able to parse https://docs.cloud.google.com/asset-inventory/docs/asset-names
	// But that format is //cloudresourcemanager.googleapis.com/tagBindings/TAG_BINDING
	// which is not the format used by the service.

	ref = strings.TrimPrefix(ref, "//cloudresourcemanager.googleapis.com/")

	tokens := strings.Split(ref, "/")
	if len(tokens) == 4 && tokens[0] == "tagBindings" && tokens[2] == "tagValues" {
		i.Parent = tokens[1]
		i.TagValue = tokens[3]
		if i.TagValue == "" {
			return fmt.Errorf("tagValue was empty in external=%q", ref)
		}
		return nil
	}

	return fmt.Errorf("format of TagBinding external=%q was not known (use %s)", ref, TagsLocationTagBindingIdentityURL)
}

var _ identity.Resource = &TagsLocationTagBinding{}

func (obj *TagsLocationTagBinding) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	// Get desired resource ID
	resourceID := common.ValueOf(obj.Spec.ResourceID)

	// Server-generated ID; do not fallback to name
	// if resourceID == "" {
	// 	resourceID = obj.GetName()
	// }

	var specIdentity *TagsLocationTagBindingIdentity
	if resourceID != "" {
		specIdentity = &TagsLocationTagBindingIdentity{}
		if !strings.HasPrefix(resourceID, "tagBindings/") {
			resourceID = "tagBindings/" + resourceID
		}
		if err := specIdentity.FromExternal(resourceID); err != nil {
			return nil, fmt.Errorf("cannot parse spec.resourceID=%q: %w", resourceID, err)
		}
	}

	// Validate against the ID stored in status.externalRef
	var statusIdentity *TagsLocationTagBindingIdentity
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		statusIdentity = &TagsLocationTagBindingIdentity{}
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
