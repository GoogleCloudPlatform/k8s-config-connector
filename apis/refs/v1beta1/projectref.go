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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// The Project that this resource belongs to.
type ProjectRef struct {
	/* The `projectID` field of a project, when not managed by Config Connector. */
	External string `json:"external,omitempty"`
	/* The `name` field of a `Project` resource. */
	Name string `json:"name,omitempty"`
	/* The `namespace` field of a `Project` resource. */
	Namespace string `json:"namespace,omitempty"`
	// The kind of the Project resource; optional but must be `Project` if provided.
	// +optional
	Kind string `json:"kind,omitempty"`
}

var _ Ref = &ProjectRef{}

func (r *ProjectRef) GetGVK() schema.GroupVersionKind {
	return schema.GroupVersionKind{
		Group:   "resourcemanager.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "Project",
	}
}

func (r *ProjectRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *ProjectRef) GetExternal() string {
	return r.External
}

func (r *ProjectRef) SetExternal(ref string) {
	r.External = ref
}

func (r *ProjectRef) ValidateExternal(ref string) error {
	id := &Project{}
	if err := id.FromExternal(r.GetExternal()); err != nil {
		return err
	}
	return nil
}

func (r *ProjectRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	// No status.externalRef, so can't use default method
	// return Normalize(ctx, reader, r, defaultNamespace)

	project, err := ResolveProject(ctx, reader, defaultNamespace, r)
	if err != nil {
		return err
	}

	r.External = "projects/" + project.ProjectID
	return nil
}

// AsProjectRef converts a generic ResourceRef into a ProjectRef
func AsProjectRef(in *v1alpha1.ResourceRef) *ProjectRef {
	if in == nil {
		return nil
	}
	return &ProjectRef{
		Namespace: in.Namespace,
		Name:      in.Name,
		External:  in.External,
		Kind:      in.Kind,
	}
}

type Project struct {
	ProjectID string
}

var _ identity.Identity = &Project{}

func (p *Project) String() string {
	return "projects/" + p.ProjectID
}

func (p *Project) FromExternal(ref string) error {
	tokens := strings.Split(ref, "/")
	if len(tokens) == 1 {
		p.ProjectID = tokens[0]
		return nil
	}
	if len(tokens) == 2 && tokens[0] == "projects" {
		p.ProjectID = tokens[1]
		return nil
	}
	return fmt.Errorf("unknown format for project %q (use projects/{projectId})", ref)
}

// ResolveProjectFromAnnotation resolves the projectID to use for a resource,
// it should be used for resources which do not have a projectRef
func ResolveProjectFromAnnotation(ctx context.Context, reader client.Reader, src client.Object) (*Project, error) {
	if projectID := src.GetAnnotations()["cnrm.cloud.google.com/project-id"]; projectID != "" {
		return &Project{ProjectID: projectID}, nil
	}

	return nil, fmt.Errorf("project-id annotation not set on resource")
}

// ResolveProject will resolve a ProjectRef to a Project, with the ProjectID.
func ResolveProject(ctx context.Context, reader client.Reader, otherNamespace string, ref *ProjectRef) (*Project, error) {
	if ref == nil {
		return nil, nil
	}

	if ref.Kind != "" {
		if ref.Kind != "Project" {
			return nil, fmt.Errorf("kind is optional on project reference, but must be \"Project\" if provided")
		}
	}

	if ref.External != "" {
		if ref.Name != "" {
			return nil, fmt.Errorf("cannot specify both name and external on project reference")
		}

		tokens := strings.Split(ref.External, "/")
		if len(tokens) == 1 {
			return &Project{ProjectID: tokens[0]}, nil
		}
		if len(tokens) == 2 && tokens[0] == "projects" {
			return &Project{ProjectID: tokens[1]}, nil
		}
		return nil, fmt.Errorf("format of project external=%q was not known (use projects/<projectId> or <projectId>)", ref.External)
	}

	if ref.Name == "" {
		return nil, fmt.Errorf("must specify either name or external on project reference")
	}

	key := types.NamespacedName{
		Namespace: ref.Namespace,
		Name:      ref.Name,
	}
	if key.Namespace == "" {
		key.Namespace = otherNamespace
	}

	project := &unstructured.Unstructured{}
	project.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "resourcemanager.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "Project",
	})
	if err := reader.Get(ctx, key, project); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, fmt.Errorf("referenced Project %v not found", key)
		}
		return nil, fmt.Errorf("error reading referenced Project %v: %w", key, err)
	}

	projectID, err := GetResourceID(project)
	if err != nil {
		return nil, err
	}

	return &Project{
		ProjectID: projectID,
	}, nil
}

func ResolveProjectID(ctx context.Context, reader client.Reader, obj *unstructured.Unstructured) (string, error) {
	projectRefExternal, _, _ := unstructured.NestedString(obj.Object, "spec", "projectRef", "external")
	if projectRefExternal != "" {
		projectRef := ProjectRef{
			External: projectRefExternal,
		}

		project, err := ResolveProject(ctx, reader, obj.GetNamespace(), &projectRef)
		if err != nil {
			return "", fmt.Errorf("cannot parse projectRef.external %q in %v %v/%v: %w", projectRefExternal, obj.GetKind(), obj.GetNamespace(), obj.GetName(), err)
		}
		return project.ProjectID, nil
	}

	projectRefName, _, _ := unstructured.NestedString(obj.Object, "spec", "projectRef", "name")
	if projectRefName != "" {
		projectRefNamespace, _, _ := unstructured.NestedString(obj.Object, "spec", "projectRef", "namespace")

		projectRef := ProjectRef{
			Name:      projectRefName,
			Namespace: projectRefNamespace,
		}
		if projectRef.Namespace == "" {
			projectRef.Namespace = obj.GetNamespace()
		}

		project, err := ResolveProject(ctx, reader, obj.GetNamespace(), &projectRef)
		if err != nil {
			return "", fmt.Errorf("cannot parse projectRef in %v %v/%v: %w", obj.GetKind(), obj.GetNamespace(), obj.GetName(), err)
		}
		return project.ProjectID, nil
	}

	if projectID := obj.GetAnnotations()["cnrm.cloud.google.com/project-id"]; projectID != "" {
		return projectID, nil
	}

	return "", fmt.Errorf("cannot find project id for %v %v/%v", obj.GetKind(), obj.GetNamespace(), obj.GetName())
}
