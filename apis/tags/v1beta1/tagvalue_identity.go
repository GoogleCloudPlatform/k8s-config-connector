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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	// TagsTagValueIdentityURL is the format for the externalRef of a TagsTagValue.
	TagsTagValueIdentityURL = "tagValues/{{tagValue}}"
)

var _ identity.Identity = &TagsTagValueIdentity{}

var tagValueURL = gcpurls.Template[TagsTagValueIdentity](
	"cloudresourcemanager.googleapis.com",
	"tagValues/{TagValue}",
)

// TagsTagValueIdentity represents the identity of a TagsTagValue.
// +k8s:deepcopy-gen=false
type TagsTagValueIdentity struct {
	TagValue string
}

func (i *TagsTagValueIdentity) String() string {
	return "tagValues/" + i.TagValue
}

func (i *TagsTagValueIdentity) FromExternal(ref string) error {
	out, match, err := tagValueURL.Parse(ref)
	if err != nil {
		return err
	}
	if !match {
		return fmt.Errorf("format of TagValue external=%q was not known (use %s)", ref, TagsTagValueIdentityURL)
	}
	if out.TagValue == "" {
		return fmt.Errorf("tagValue was empty in external=%q", ref)
	}
	*i = *out
	return nil
}

var _ identity.Resource = &TagsTagValue{}

func (obj *TagsTagValue) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	// Get desired resource ID
	resourceID := common.ValueOf(obj.Spec.ResourceID)

	// Server-generated ID; do not fallback to name
	// if resourceID == "" {
	// 	resourceID = obj.GetName()
	// }

	var specIdentity *TagsTagValueIdentity
	if resourceID != "" {
		specIdentity = &TagsTagValueIdentity{
			TagValue: resourceID,
		}
	}

	// Validate against the ID stored in status.externalRef
	var statusIdentity *TagsTagValueIdentity
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		statusIdentity = &TagsTagValueIdentity{}
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
