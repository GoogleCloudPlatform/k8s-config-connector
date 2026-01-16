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

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	storage "github.com/GoogleCloudPlatform/k8s-config-connector/apis/storage/v1beta1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// ParentRef is a reference to a parent resource.
// +kcc:ref=Project;StorageBucket
type TagsTagBindingParentRef struct {
	// Kind of the referent.
	// +optional
	// +kubebuilder:validation:Optional
	// +kubebuilder:default=Project
	Kind string `json:"kind,omitempty"`

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
	_, id, _ := r.resolveReference()

	if id == nil {
		return schema.GroupVersionKind{}
	}

	return id.GetGVK()
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
	_, id, err := r.resolveReference()
	if err != nil {
		return err
	}
	if id == nil {
		return fmt.Errorf("unknown format for external reference %q", r.External)
	}

	if err := id.ValidateExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *TagsTagBindingParentRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	service, id, err := r.resolveReference()
	if err != nil {
		return err
	}
	if id == nil {
		return fmt.Errorf("unknown format for external reference %q", r.External)
	}

	if err := id.Normalize(ctx, reader, defaultNamespace); err != nil {
		return err
	}

	// Include the service qualification in case of ambiguities
	r.External = "//" + service + "/" + id.GetExternal()

	return nil
}

func (r *TagsTagBindingParentRef) resolveReference() (string, refs.Ref, error) {
	kind := r.Kind
	if kind == "" {
		kind = "Project"
	}

	switch kind {
	case "StorageBucket":
		return "storage.googleapis.com", &storage.StorageBucketRef{
			Name:      r.Name,
			Namespace: r.Namespace,
			External:  strings.TrimPrefix(r.External, "//storage.googleapis.com/"),
		}, nil
	case "Project":
		if r.External != "" && !isProjectExternal(r.External) {
			return "", nil, fmt.Errorf("unknown format for a Project reference in %q, please set Kind for non-project references or use projects/{project} for project references", r.External)
		}
		return "cloudresourcemanager.googleapis.com", &refs.ProjectRef{
			Name:      r.Name,
			Namespace: r.Namespace,
			External:  r.External,
		}, nil
	default:
		return "", nil, nil
	}
}

func isProjectExternal(external string) bool {
	if strings.HasPrefix(external, "//cloudresourcemanager.googleapis.com/") {
		return true
	}
	if strings.HasPrefix(external, "projects/") {
		return true
	}
	// "123456789" is also a valid project external (project number)
	if !strings.Contains(external, "/") {
		return true
	}
	return false
}
