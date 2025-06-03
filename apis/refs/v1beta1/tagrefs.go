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
	"k8s.io/klog/v2"
)

// TODO Support organization level tags
// TODO Support identifying by Parent and ShortName in refs, if it is possible

// Tags can be assigned as `tagKeys/[tag_key_id]` mapped to `tagValues/[tag_value_id]`,
// or `[org id, project id, or project number]/[tag_key_shortname]` mapped to `[value_shortname]`
type TagValueRef struct {
	// Should be in the format `[CONTAINER_ID]/[tag_key_shortname]/[tag_value_shortname]` for the tag key holding the value,
	// where `CONTAINER_ID` is `[org_id]`, `[project_id]`, or `[project_number]`.
	External string `json:"external,omitempty"`
	// The `metadata.name` field of a `TagValue` resource.
	Name string `json:"name,omitempty"`
	// The k8s `namespace` of a `TagValue` resource.
	Namespace string `json:"namespace,omitempty"`
}

type TagValue struct {
	Ref        *TagValueRef
	ResourceID string
}

// ResolveTagValueRef will resolve a TagValueRef to a TagValue.
func ResolveTagValueRef(ctx context.Context, reader client.Reader, src client.Object, ref *TagValueRef) (*TagValueRef, error) {
	if ref == nil {
		return nil, nil
	}

	klog.Infof("ResolveTagValueRef name: %s, External: %s", ref.Name, ref.External)

	if ref.Name == "" && ref.External == "" {
		return nil, fmt.Errorf("must specify either name or external on TagValueRef")
	}
	if ref.Name != "" && ref.External != "" {
		return nil, fmt.Errorf("cannot specify both name and external on TagValueRef")
	}

	// External should be in the `tagKeys/[tag_key_id]/[tag_value_id]` format or `[CONTAINER_ID]/[tag_key_shortname]/[tag_value_shortname]`
	if ref.External != "" {
		tokens := strings.Split(ref.External, "/")
		if len(tokens) == 3 {
			ref = &TagValueRef{
				External: fmt.Sprintf("%s/%s/%s", tokens[0], tokens[1], tokens[2]),
			}
			return ref, nil
		}
		return nil, fmt.Errorf("format of TagValueRef external=%q was not known (use [CONTAINER_ID]/[tag_key_shortname]/[tag_value_shortname])", ref.External)
	}

	// Namespace is referring to k8s namespace, not tags namespace
	key := types.NamespacedName{
		Namespace: ref.Namespace,
		Name:      ref.Name,
	}
	if key.Namespace == "" {
		key.Namespace = src.GetNamespace()
	}

	// Fetch object from k8s cluster to construct the external form
	tagValue := &unstructured.Unstructured{}
	tagValue.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "tags.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "TagsTagValue",
	})
	if err := reader.Get(ctx, key, tagValue); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, fmt.Errorf("referenced TagsTagValue %v not found.", key)
		}
		return nil, fmt.Errorf("error reading referenced TagsTagValue %v: %w.", key, err)
	}

	tagKey, err := ResolveTagKeyForObject(ctx, reader, tagValue)
	if err != nil {
		return nil, err
	}

	ref = &TagValueRef{
		// `[CONTAINER_ID]/[key_shortname]/[value_shortname]` format
		External: fmt.Sprintf("%s/%s", tagKey.Ref.External, key.Name),
	}

	klog.Infof("ResolveTagValueRef 2 name: %s, External: %s", ref.Name, ref.External)

	return ref, nil
}

type TagKeyRef struct {
	// If provided must be in the format `[CONTAINER_ID]/[tag_key_shortname]` holding the key, where `CONTAINER_ID` is `[org_id]`, `[project_id]`, or `[project_number]`.
	External string `json:"external,omitempty"`
	// The `metadata.name` field of a `TagKey` resource.
	Name string `json:"name,omitempty"`
	// The `metadata.namespace` field of a `TagKey` resource.
	Namespace string `json:"namespace,omitempty"`
}

type TagKey struct {
	Ref        *TagKeyRef
	ResourceID string
}

// ResolveTagKeyRef will resolve a TagKeyRef to a TagKey.
func ResolveTagKeyRef(ctx context.Context, reader client.Reader, src client.Object, ref *TagKeyRef) (*TagKey, error) {
	if ref == nil {
		return nil, nil
	}

	klog.Infof("ResolveTagKeyRef name: %s, External: %s", ref.Name, ref.External)

	if ref.Name == "" && ref.External == "" {
		return nil, fmt.Errorf("must specify either name or external on TagKeyRef")
	}
	if ref.Name != "" && ref.External != "" {
		return nil, fmt.Errorf("cannot specify both name and external on TagKeyRef")
	}

	// External should be in the `[CONTAINER_ID]/[tag_key_shortname]` format
	if ref.External != "" {
		tokens := strings.Split(ref.External, "/")
		if len(tokens) == 2 {
			ref = &TagKeyRef{
				External: fmt.Sprintf("%s/%s", tokens[0], tokens[1]),
			}
			return &TagKey{Ref: ref, ResourceID: tokens[1]}, nil
		}
		return nil, fmt.Errorf("format of TagKeyRef external=%q was not known (use [CONTAINER_ID]/[tag_key_shortname])", ref.External)
	}

	key := types.NamespacedName{
		Namespace: ref.Namespace,
		Name:      ref.Name,
	}
	if key.Namespace == "" {
		key.Namespace = src.GetNamespace()
	}

	// Fetch object from k8s cluster to construct the external form
	tagKey := &unstructured.Unstructured{}
	tagKey.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "tags.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "TagsTagKey",
	})
	if err := reader.Get(ctx, key, tagKey); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, fmt.Errorf("referenced TagsTagKey %v not found", key)
		}
		return nil, fmt.Errorf("error reading referenced TagsTagKey %v: %w", key, err)
	}

	// TODO: Gives the Name back when running e2e samples tests. Should give the GCP resource ID. Or we completely ignore it in favor of using key shortname
	tagKeyResourceID, err := GetResourceID(tagKey)
	if err != nil {
		return nil, err
	}

	projectID, err := ResolveProjectID(ctx, reader, tagKey)
	if err != nil {
		return nil, err
	}

	ref = &TagKeyRef{
		// `[CONTAINER]/[key_shortname]` format`
		External: fmt.Sprintf("%s/%s", projectID, key.Name),
	}

	klog.Infof("ResolveTagKeyRef 2 name: %s, External: %s", ref.Name, ref.External)

	return &TagKey{Ref: ref, ResourceID: tagKeyResourceID}, nil
}

func ResolveTagKeyForObject(ctx context.Context, reader client.Reader, obj *unstructured.Unstructured) (*TagKey, error) {
	tagKeyRefExternal, _, err := unstructured.NestedString(obj.Object, "spec", "parentRef", "external")
	if err != nil {
		return nil, fmt.Errorf("error fetching parentRef.external %w", err)
	}
	klog.Infof("ResolveTagKeyForObject External: %s", tagKeyRefExternal)
	if tagKeyRefExternal != "" {
		return ResolveTagKeyRef(ctx, reader, obj, &TagKeyRef{External: tagKeyRefExternal})
	}

	tagKeyRefName, _, err := unstructured.NestedString(obj.Object, "spec", "parentRef", "name")
	if err != nil {
		return nil, fmt.Errorf("error fetching parentRef.name %w", err)
	}
	klog.Infof("ResolveTagKeyForObject Name: %s", tagKeyRefName)
	if tagKeyRefName != "" {
		tagKeyRefNamespace, _, err := unstructured.NestedString(obj.Object, "spec", "parentRef", "namespace")
		if err != nil {
			return nil, fmt.Errorf("error fetching parentRef.namespace %w", err)
		}

		tagKeyRef := TagKeyRef{
			Name:      tagKeyRefName,
			Namespace: tagKeyRefNamespace,
		}
		if tagKeyRef.Namespace == "" {
			tagKeyRef.Namespace = obj.GetNamespace()
		}

		return ResolveTagKeyRef(ctx, reader, obj, &tagKeyRef)
	}

	return nil, fmt.Errorf("cannot find tagKeyRef for %v %v/%v", obj.GetKind(), obj.GetNamespace(), obj.GetName())
}
