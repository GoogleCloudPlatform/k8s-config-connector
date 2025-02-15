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

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// GroupRef is a reference to a Group resource.
type GroupRef struct {
	/* The group for the resource
	   Allowed value: The Google Cloud resource name of a `CloudIdentityGroup` resource (format: `groups/{{name}}`). */
	External string `json:"external,omitempty"`
	/* The `name` field of a `CloudIdentityGroup` resource. */
	Name string `json:"name,omitempty"`
	/* The `namespace` field of a `CloudIdentityGroup` resource. */
	Namespace string `json:"namespace,omitempty"`
	// The kind of the Group resource; optional but must be `CloudIdentityGroup` if provided.
	// +optional
	Kind string `json:"kind,omitempty"`
}

// AsGroupRef converts a generic ResourceRef into a GroupRef
func AsGroupRef(in *v1alpha1.ResourceRef) *GroupRef {
	if in == nil {
		return nil
	}
	return &GroupRef{
		Namespace: in.Namespace,
		Name:      in.Name,
		External:  in.External,
		Kind:      in.Kind,
	}
}

type Group struct {
	GroupID string
}

// ResolveGroup will resolve a GroupRef to a Group, with the GroupID.
func ResolveGroup(ctx context.Context, reader client.Reader, otherNamespace string, ref *GroupRef) (*Group, error) {
	if ref == nil {
		return nil, nil
	}

	if ref.Kind != "" {
		if ref.Kind != "CloudIdentityGroup" {
			return nil, fmt.Errorf("kind is optional on group reference, but must be \"CloudIdentityGroup\" if provided")
		}
	}

	if ref.External != "" {
		if ref.Name != "" {
			return nil, fmt.Errorf("cannot specify both name and external on group reference")
		}

		tokens := strings.Split(ref.External, "/")
		if len(tokens) == 1 {
			return &Group{GroupID: tokens[0]}, nil
		}
		if len(tokens) == 2 && tokens[0] == "groups" {
			return &Group{GroupID: tokens[1]}, nil
		}
		return nil, fmt.Errorf("format of group external=%q was not known (use groups/<groupId> or <groupId>)", ref.External)
	}

	if ref.Name == "" {
		return nil, fmt.Errorf("must specify either name or external on group reference")
	}

	key := types.NamespacedName{
		Namespace: ref.Namespace,
		Name:      ref.Name,
	}
	if key.Namespace == "" {
		key.Namespace = otherNamespace
	}

	group := &unstructured.Unstructured{}
	group.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "cloudidentity.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "CloudIdentityGroup",
	})
	if err := reader.Get(ctx, key, group); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, fmt.Errorf("referenced CloudIdentityGroup %v not found", key)
		}
		return nil, fmt.Errorf("error reading referenced CloudIdentityGroup %v: %w", key, err)
	}

	groupID, err := GetResourceID(group)
	if err != nil {
		return nil, err
	}

	return &Group{
		GroupID: groupID,
	}, nil
}

func ResolveGroupID(ctx context.Context, reader client.Reader, obj *unstructured.Unstructured) (string, error) {
	groupRefExternal, _, _ := unstructured.NestedString(obj.Object, "spec", "groupRef", "external")
	if groupRefExternal != "" {
		groupRef := GroupRef{
			External: groupRefExternal,
		}

		group, err := ResolveGroup(ctx, reader, obj.GetNamespace(), &groupRef)
		if err != nil {
			return "", fmt.Errorf("cannot parse groupRef.external %q in %v %v/%v: %w", groupRefExternal, obj.GetKind(), obj.GetNamespace(), obj.GetName(), err)
		}
		return group.GroupID, nil
	}

	groupRefName, _, _ := unstructured.NestedString(obj.Object, "spec", "groupRef", "name")
	if groupRefName != "" {
		groupRefNamespace, _, _ := unstructured.NestedString(obj.Object, "spec", "groupRef", "namespace")

		groupRef := GroupRef{
			Name:      groupRefName,
			Namespace: groupRefNamespace,
		}
		if groupRef.Namespace == "" {
			groupRef.Namespace = obj.GetNamespace()
		}

		group, err := ResolveGroup(ctx, reader, obj.GetNamespace(), &groupRef)
		if err != nil {
			return "", fmt.Errorf("cannot parse groupRef in %v %v/%v: %w", obj.GetKind(), obj.GetNamespace(), obj.GetName(), err)
		}
		return group.GroupID, nil
	}

	return "", fmt.Errorf("cannot find group id for %v %v/%v", obj.GetKind(), obj.GetNamespace(), obj.GetName())
}
