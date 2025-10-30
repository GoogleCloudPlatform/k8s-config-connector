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
	"net/url"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"

	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ identity.Identity = &TagBindingIdentity{}

const TagBindingIDURL = "tagBindings/{{parentWithFullURL}}/tagValues/{{tagValueID}"

// +k8s:deepcopy-gen=false
type TagBindingIdentity struct {
	// The Parent uniform the diverse parent kind, with the full resource name.
	parent *TagBindingParent
	// tagValue is the resource name of the TagValue.
	tagValue string
}

func (i *TagBindingIdentity) String() string {
	return fmt.Sprintf("tagBindings/%s/tagValues/%s", url.PathEscape(i.parent.String()), i.tagValue)
}

func (i *TagBindingIdentity) TagValue() string {
	return "tagValues/" + i.tagValue
}

func (i *TagBindingIdentity) Parent() *TagBindingParent {
	return i.parent
}

func (i *TagBindingIdentity) FromExternal(ref string) error {
	// Legacy Terraform reconciler trims the `tagBindings` prefix, and the direct controller keeps the real value of tagBinding name from the GCP
	ref = strings.TrimPrefix(ref, "tagBindings/")

	tokens := strings.Split(ref, "/")
	if len(tokens) != 3 || tokens[1] != "tagValues" {
		return fmt.Errorf("format of TagBinding external=%q was not known (use %s)", ref, TagBindingIDURL)
	}

	i.parent = &TagBindingParent{}
	if err := i.parent.FromExternal(tokens[0]); err != nil {
		return err
	}
	if i.tagValue == "" {
		return fmt.Errorf("tagValue was empty in external=%q", ref)
	}
	return nil
}

var _ identity.Resource = &TagsTagBinding{}

func (obj *TagsTagBinding) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	newIdentity := &TagBindingIdentity{
		parent:   &TagBindingParent{},
		tagValue: "",
	}

	parentRef := obj.Spec.ParentRef
	if err := parentRef.Normalize(ctx, reader, obj.GetNamespace()); err != nil {
		return nil, err
	}
	if err := newIdentity.parent.FromExternal(parentRef.External); err != nil {
		return nil, fmt.Errorf("parsing parentRef.external=%q: %w", parentRef.External, err)
	}

	if err := obj.Spec.TagValueRef.Normalize(ctx, reader, obj.GetNamespace()); err != nil {
		return nil, err
	}
	newIdentity.tagValue = obj.Spec.TagValueRef.GetExternal()

	return newIdentity, nil
}
