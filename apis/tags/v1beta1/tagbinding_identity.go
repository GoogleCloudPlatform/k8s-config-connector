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

	"sigs.k8s.io/controller-runtime/pkg/client"
)

// TagBindingIdentity defines the resource reference to TagsTagBinding, which "External" field
// holds the GCP identifier for the KRM object.
type TagBindingIdentity struct {
	parent   string
	tagValue string
}

func (i *TagBindingIdentity) String() string {
	return fmt.Sprintf("tagBindings/%s/%s", i.parent, i.tagValue)
}

func (i *TagBindingIdentity) Parent() string {
	return i.parent
}

func (i *TagBindingIdentity) TagValue() string {
	return i.tagValue
}

// New builds a TagBindingIdentity from the Config Connector TagBinding object.
func NewTagBindingIdentity(ctx context.Context, reader client.Reader, obj *TagsTagBinding) (*TagBindingIdentity, error) {
	parent, err := resolveParent(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	tagValue, err := resolveTagValue(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	return &TagBindingIdentity{
		parent:   parent,
		tagValue: tagValue,
	}, nil
}

func resolveParent(ctx context.Context, reader client.Reader, obj *TagsTagBinding) (string, error) {
	if obj.Spec.ParentRef.External != "" {
		return obj.Spec.ParentRef.External, nil
	}
	if obj.Spec.ParentRef.Name != "" {
		// TODO: This is not quite right, we need to resolve the project number.
		// For now, we will just use the name.
		return fmt.Sprintf("//cloudresourcemanager.googleapis.com/projects/%s", obj.Spec.ParentRef.Name), nil
	}
	return "", fmt.Errorf("parentRef is required")
}

func resolveTagValue(ctx context.Context, reader client.Reader, obj *TagsTagBinding) (string, error) {
	if obj.Spec.TagValueRef.External != "" {
		return obj.Spec.TagValueRef.External, nil
	}
	if obj.Spec.TagValueRef.Name != "" {
		// TODO: This is not quite right, we need to resolve the tag value name.
		// For now, we will just use the name.
		return fmt.Sprintf("tagValues/%s", obj.Spec.TagValueRef.Name), nil
	}
	return "", fmt.Errorf("tagValueRef is required")
}
