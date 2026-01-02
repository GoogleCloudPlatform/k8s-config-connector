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

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refsv1beta1.Ref = &TagsTagBindingRef{}

// TagsTagBindingRef is a reference to a TagsTagBinding resource.
type TagsTagBindingRef struct {
	// A reference to an externally managed TagsTagBinding resource.
	// Should be in the format "tagBindings/{{tagBindingID}}".
	External string `json:"external,omitempty"`

	// The name of a TagsTagBinding resource.
	Name string `json:"name,omitempty"`

	// The namespace of a TagsTagBinding resource.
	Namespace string `json:"namespace,omitempty"`
}

func (r *TagsTagBindingRef) GetGVK() schema.GroupVersionKind {
	return TagsTagBindingGVK
}

func (r *TagsTagBindingRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *TagsTagBindingRef) GetExternal() string {
	return r.External
}

func (r *TagsTagBindingRef) SetExternal(ref string) {
	r.External = ref
}

func (r *TagsTagBindingRef) ValidateExternal(ref string) error {
	id := &TagsTagBindingIdentity{}
	if err := id.FromExternal(r.GetExternal()); err != nil {
		return err
	}
	return nil
}

func (r *TagsTagBindingRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	fallback := func(u *unstructured.Unstructured) string {
		name, _, _ := unstructured.NestedString(u.Object, "status", "name")
		if name != "" {
			return "tagBindings/" + name
		}
		return ""
	}
	return refsv1beta1.NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback)
}
