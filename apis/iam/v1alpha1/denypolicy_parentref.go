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

package v1alpha1

import (
	"context"
	"fmt"

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// IAMDenyPolicyParentRef is a reference to a parent resource.
// +kcc:ref=Project;Folder;Organization
type IAMDenyPolicyParentRef struct {
	// Kind to which we are attaching the policy. Defaults to Project if not specified.
	// +optional
	// +kubebuilder:validation:Optional
	Kind string `json:"kind,omitempty"`

	// Name of the referent.
	Name string `json:"name,omitempty"`

	// Namespace of the referent.
	// +optional
	// +kubebuilder:validation:Optional
	Namespace string `json:"namespace,omitempty"`

	// External string for the referent.
	// +optional
	// +kubebuilder:validation:Optional
	External string `json:"external,omitempty"`
}

var _ refs.Ref = &IAMDenyPolicyParentRef{}

func (r *IAMDenyPolicyParentRef) GetGVK() schema.GroupVersionKind {
	return schema.GroupVersionKind{Kind: r.Kind}
}

func (r *IAMDenyPolicyParentRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Namespace: r.Namespace,
		Name:      r.Name,
	}
}

func (r *IAMDenyPolicyParentRef) GetExternal() string {
	return r.External
}

func (r *IAMDenyPolicyParentRef) SetExternal(ref string) {
	r.External = ref
}

func (r *IAMDenyPolicyParentRef) ValidateExternal(ref string) error {
	if ref == "" {
		return fmt.Errorf("external cannot be empty")
	}
	return nil
}

func (r *IAMDenyPolicyParentRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	if r.External != "" {
		return nil
	}

	kind := r.Kind
	if kind == "" {
		kind = "Project"
	}

	nn := r.GetNamespacedName()
	if nn.Namespace == "" {
		nn.Namespace = defaultNamespace
	}

	switch kind {
	case "Project":
		projectRef := &refs.ProjectRef{
			Name:      nn.Name,
			Namespace: nn.Namespace,
			External:  r.External,
		}
		normalized, err := refs.ResolveProject(ctx, reader, nn.Namespace, projectRef)
		if err != nil {
			return err
		}
		// url-encoded format for project
		// For project, the format is cloudresourcemanager.googleapis.com%2Fprojects%2F{ProjectID}
		// but we store the project ID in External, and we'll format it in the adapter.
		r.External = normalized.ProjectID

	case "Folder":
		folderRef := &refs.FolderRef{
			Name:      nn.Name,
			Namespace: nn.Namespace,
			External:  r.External,
		}
		src := &unstructured.Unstructured{}
		src.SetNamespace(nn.Namespace)
		normalized, err := refs.ResolveFolder(ctx, reader, src, folderRef)
		if err != nil {
			return err
		}
		r.External = normalized.FolderID

	case "Organization":
		orgRef := &refs.OrganizationRef{
			External: r.External,
		}
		// OrganizationRef does not use Name/Namespace because there's no CRD
		src := &unstructured.Unstructured{}
		src.SetNamespace(nn.Namespace)
		normalized, err := refs.ResolveOrganization(ctx, reader, src, orgRef)
		if err != nil {
			return err
		}
		r.External = normalized.OrganizationID
	default:
		return fmt.Errorf("IAMDenyPolicy does not support parent kind %q", r.Kind)
	}

	return nil
}
