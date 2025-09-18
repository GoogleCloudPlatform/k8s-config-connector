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

package v1beta1

import (
	"context"
	"fmt"
	"strings"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type TagValueRef struct {
	/* The SQLInstance selfLink, when not managed by Config Connector. */
	External string `json:"external,omitempty"`
	/* The `name` field of a `SQLInstance` resource. */
	Name string `json:"name,omitempty"`
	/* The `namespace` field of a `SQLInstance` resource. */
	Namespace string `json:"namespace,omitempty"`
}

type TagValue struct {
	value string
}

func (s *TagValue) String() string {
	return "`tagKeys/" + s.value
}

func ResolveSTagValueRef(ctx context.Context, reader client.Reader, obj client.Object, ref *TagValueRef) (*TagValue, error) {
	if ref == nil {
		return nil, nil
	}

	if ref.Name == "" && ref.External == "" {
		return nil, fmt.Errorf("must specify either name or external on tagValueRef")
	}
	if ref.External != "" && ref.Name != "" {
		return nil, fmt.Errorf("cannot specify both spec.tagValueRef.name and spec.tagValueRef.external")
	}

	if ref.External != "" {
		// External should be in the format `tagKeys/{{value}}`.
		tokens := strings.Split(ref.External, "/")
		if len(tokens) == 2 && tokens[0] == "tagKeys" {
			return &TagValue{
				value: tokens[1],
			}, nil
		}
		return nil, fmt.Errorf("format of tagvalueRef external=%q was not known (use tagKeys/{{value}})", ref.External)
	}

	key := types.NamespacedName{
		Namespace: ref.Namespace,
		Name:      ref.Name,
	}
	if key.Namespace == "" {
		key.Namespace = obj.GetNamespace()
	}

	tagvalue := &unstructured.Unstructured{}
	tagvalue.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "tags.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "TagsTagValue",
	})
	if err := reader.Get(ctx, key, tagvalue); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, fmt.Errorf("referenced tagvalue %v not found", key)
		}
		return nil, fmt.Errorf("error reading referenced tagvalue %v: %w", key, err)
	}

	resourceID, _, err := unstructured.NestedString(tagvalue.Object, "spec", "resourceID")
	if err != nil {
		return nil, fmt.Errorf("reading spec.resourceID from SQLInstance %s/%s: %w", tagvalue.GetNamespace(), tagvalue.GetName(), err)
	}
	if resourceID == "" {
		resourceID = tagvalue.GetName()
	}

	return &TagValue{
		value: resourceID,
	}, nil
}
