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

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// ParentRef is a reference to a parent resource.
// +kcc:ref=Project
type TagsTagBindingParentRef struct {
	// Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Name string `json:"name,omitempty"`

	// Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/
	Namespace string `json:"namespace,omitempty"`

	// Allowed value: string of the format `//cloudresourcemanager.googleapis.com/projects/{{value}}`,
	// where {{value}} is the `number` field of a `Project` resource.
	External string `json:"external,omitempty"`
}

// TagsTagBindingParentRef implements the standard refs.Ref interface
var _ refs.Ref = &TagsTagBindingParentRef{}

func (r *TagsTagBindingParentRef) GetGVK() schema.GroupVersionKind {
	// Add other kinds here as they are supported
	return refs.ProjectGVK
}

func (r *TagsTagBindingParentRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *TagsTagBindingParentRef) GetExternal() string {
	return r.External
}

func (r *TagsTagBindingParentRef) SetExternal(ref string) {
	r.External = ref
}

func (r *TagsTagBindingParentRef) ValidateExternal(ref string) error {
	gvk := r.GetGVK()
	switch gvk {
	case refs.ProjectGVK:
		id := &refs.ProjectRef{
			Name:      r.Name,
			Namespace: r.Namespace,
			External:  r.External,
		}
		if err := id.ValidateExternal(ref); err != nil {
			return err
		}
		return nil
	default:
		return fmt.Errorf("format of %s external=%q was not known", gvk.Kind, ref)
	}
}

func (r *TagsTagBindingParentRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	gvk := r.GetGVK()
	switch gvk {
	case refs.ProjectGVK:
		id := &refs.ProjectRef{
			Name:      r.Name,
			Namespace: r.Namespace,
			External:  r.External,
		}
		if err := id.Normalize(ctx, reader, defaultNamespace); err != nil {
			return err
		}
		// Include the service qualification in case of ambiguities
		r.External = "//cloudresourcemanager.googleapis.com/" + id.External
		return nil
	default:
		return fmt.Errorf("unsupported gvk for reference %v", gvk)
	}
}
