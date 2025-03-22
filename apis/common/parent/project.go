// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package parent

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

var _ Parent = &ProjectParent{}

type ProjectParent struct {
	ProjectID string
}

func (i *ProjectParent) FromExternal(external string) error {
	tokens := strings.Split(external, "/")
	if len(tokens) == 2 && tokens[0] == "projects" {
		i.ProjectID = tokens[1]
		return nil
	}
	return fmt.Errorf("format of project reference %q was not known (use projects/{{projectID}})", external)
}

var _ ParentBuilder = &ProjectRef{}

// Project specifies the resource's GCP hierarchy (Project/Folder/Organization).
// +kubebuilder:object:generate:=true
type ProjectRef struct {
	/* The `projectID` field of a project, when not managed by Config Connector. */
	External string `json:"external,omitempty"`
	/* The `name` field of a `Project` resource. */
	Name string `json:"name,omitempty"`
	/* The `namespace` field of a `Project` resource. */
	Namespace string `json:"namespace,omitempty"`
}

// Builds a the ProjectParent from ProjectRef.
// If `projectRef.external` is given, parse projectID from External, otherwise find the ConfigConnector project
// according to `projectRef.name` and `projectRef.namespace`.
func (p *ProjectRef) Build(ctx context.Context, reader client.Reader, othernamespace string, parent Parent) error {
	projectParent, ok := parent.(*ProjectParent)
	if !ok {
		return fmt.Errorf("build invalid parent, except %T", &ProjectParent{})
	}
	if p.External != "" {
		if p.Name != "" {
			return fmt.Errorf("cannot specify both name and external on project reference")
		}

		tokens := strings.Split(p.External, "/")
		if len(tokens) == 1 {
			projectParent.ProjectID = tokens[0]
			return nil
		}
		if len(tokens) == 2 && tokens[0] == "projects" {
			projectParent.ProjectID = tokens[1]
			return nil
		}
		return fmt.Errorf("format of project external=%q was not known (use projects/<projectId> or <projectId>)", p.External)
	}

	if p.Name == "" {
		return fmt.Errorf("must specify either name or external on project reference")
	}

	key := types.NamespacedName{
		Namespace: p.Namespace,
		Name:      p.Name,
	}
	if key.Namespace == "" {
		key.Namespace = othernamespace
	}

	project := &unstructured.Unstructured{}
	project.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "resourcemanager.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "Project",
	})
	if err := reader.Get(ctx, key, project); err != nil {
		if apierrors.IsNotFound(err) {
			return fmt.Errorf("referenced Project %v not found", key)
		}
		return fmt.Errorf("error reading referenced Project %v: %w", key, err)
	}

	projectID, _, err := unstructured.NestedString(project.Object, "spec", "resourceID")
	if err != nil {
		return fmt.Errorf("reading spec.resourceID from %v %v/%v: %w", project.GroupVersionKind().Kind, p.Namespace, p.Name, err)
	}
	if projectID == "" {
		projectID = project.GetName()
	}
	projectParent.ProjectID = projectID
	return nil
}

func ParseProjectParent(external string) (*ProjectParent, error) {
	tokens := strings.Split(external, "/")
	if len(tokens) != 2 || tokens[0] != "projects" {
		return nil, fmt.Errorf("format of Project external=%q was not known (use projects/<projectId>)", external)
	}

	return &ProjectParent{
		ProjectID: tokens[1],
	}, nil
}

func (p *ProjectParent) String() string {
	return "projects/" + p.ProjectID
}

func (p *ProjectParent) MatchActual(actualI Parent) error {
	actual, ok := actualI.(*ProjectParent)
	if !ok {
		return fmt.Errorf("parent format changed, desired %T", p)
	}
	if p.ProjectID != actual.ProjectID {
		return fmt.Errorf("spec.projectRef changed, desired %s, actual %s", p.ProjectID, actual.ProjectID)
	}
	return nil
}
