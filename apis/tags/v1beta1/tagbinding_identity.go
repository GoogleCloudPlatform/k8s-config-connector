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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ identity.Identity = &TagBindingIdentity{}

const TagBindingIDURL = "tagBindings/{{parentWithFullURL}}/tagValues/{{tagValueID}"

type TagBindingIdentity struct {
	// parent is the full resource name of the resource to which the tag is bound.
	// e.g. `//cloudresourcemanager.googleapis.com/projects/12345`
	parent string
	// tagValue is the resource name of the TagValue.
	tagValue string
}

func (i *TagBindingIdentity) String() string {
	// The name of the TagBinding. This is a String of the form:
	// `tagBindings/{full-resource-name}/{tag-value-name}`
	// (e.g. `tagBindings/%2F%2Fcloudresourcemanager.googleapis.com%2Fprojects%2F123/tagValues/456`).
	// The parent needs to be URL-encoded.
	return "tagBindings/" + url.PathEscape(i.parent) + "/tagValues/" + i.tagValue
}

func (i *TagBindingIdentity) ParentWithFullURL() string {
	return i.parent
}

func (i *TagBindingIdentity) TagValue() string {
	return "tagValues/" + i.tagValue
}

func (i *TagBindingIdentity) FromExternal(ref string) error {
	name, found := strings.CutPrefix(ref, "tagBindings/")
	if !found {
		return fmt.Errorf("format of TagBinding missing prefix `tagBindings`. got=%q (use %s)", ref, TagBindingIDURL)
	}
	tokens := strings.Split(name, "/tagValues/")
	if len(tokens) != 2 {
		return fmt.Errorf("format of TagBinding external=%q was not known (use %s)", ref, TagBindingIDURL)
	}

	// TODO: parent can be other resources in addition to Project.
	parent := &TagBindingProject{}
	if err := parent.FromExternal(tokens[0]); err != nil {
		return err
	}
	i.parent = parent.String()
	if i.parent == "" {
		return fmt.Errorf("parent was empty in external=%q", ref)
	}
	i.tagValue = tokens[1]
	if i.tagValue == "" {
		return fmt.Errorf("tagValue was empty in external=%q", ref)
	}
	return nil
}

var _ identity.Resource = &TagsTagBinding{}

func (obj *TagsTagBinding) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {

	if obj.Spec.ResourceID != nil {
		// This is a legacy special use. The prefix tagBinding is to support backward compatibility with existing resources.
		externalRef := "tagBindings/" + *obj.Spec.ResourceID
		identity := &TagBindingIdentity{}
		err := identity.FromExternal(externalRef)
		return identity, err

	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, err
	}

	return &TagBindingIdentity{
		parent:   obj.Spec.ParentRef.External,
		tagValue: obj.Spec.TagValueRef.GetExternal(),
	}, nil
}
