package apis

const IdentityTemplate = `// Copyright 2025 Google LLC
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

package {{ .Version }}

{{ if .ResourcePattern -}}
{{- if eq .ParentStyle "multi" }}
// Resource name pattern(s), from google.api.resource on {{ .ProtoMessageFullName }}:
//
{{- range .ResourcePatterns }}
//	{{ . }}
{{- end }}
//
// TODO: this resource supports multiple parent hierarchies (e.g. project/folder/organization).
// The scaffolder does not model multi-parent shapes yet; implement hierarchical parent references
// and update the Parent struct and Parse{{.ProtoMessageName}}External to match.
{{- else }}
// Resource name pattern, from google.api.resource on {{ .ProtoMessageFullName }}:
//
//	{{ .ResourcePattern }}
{{- if or (eq .ParentStyle "other") (eq .ParentStyle "organization") (eq .ParentStyle "folder") }}
//
// TODO: the scaffolder does not model this parent shape. The Parent struct and
// Parse{{.ProtoMessageName}}External below assume projects/locations, so they are
// wrong for this resource: rewrite them to match the pattern above. The scaffolded
// Spec needs the matching change, since it carries projectRef and location.
{{- end }}
{{- end }}
{{- else }}
// TODO: {{ .ProtoMessageFullName }} declares no google.api.resource pattern, so the
// parent shape and collection segment below are guesses. Verify them against the
// service's API documentation before relying on this file.
{{- end }}

import (
	"context"
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// {{.ProtoMessageName}}Identity defines the resource reference to {{.Kind}}, which "External" field
// holds the GCP identifier for the KRM object.
type {{.ProtoMessageName}}Identity struct {
	parent *{{.ProtoMessageName}}Parent
	id string
}

func (i *{{.ProtoMessageName}}Identity) String() string {
	return  i.parent.String() + "/{{.Collection}}/" + i.id
}

func (i *{{.ProtoMessageName}}Identity) ID() string {
	return i.id
}

func (i *{{.ProtoMessageName}}Identity) Parent() *{{.ProtoMessageName}}Parent {
	return  i.parent
}

{{ if eq .ParentStyle "project" -}}
type {{.ProtoMessageName}}Parent struct {
	ProjectID string
}

func (p *{{.ProtoMessageName}}Parent) String() string {
	return "projects/" + p.ProjectID
}
{{- else -}}
type {{.ProtoMessageName}}Parent struct {
	ProjectID string
	Location  string
}

func (p *{{.ProtoMessageName}}Parent) String() string {
	return "projects/" + p.ProjectID + "/locations/" + p.Location
}
{{- end }}


// New builds a {{.ProtoMessageName}}Identity from the Config Connector {{.ProtoMessageName}} object.
func New{{.ProtoMessageName}}Identity(ctx context.Context, reader client.Reader, obj *{{.Kind}}) (*{{.ProtoMessageName}}Identity, error) {

	// Get Parent
	projectRef, err := refsv1beta1.ResolveProject(ctx, reader, obj.GetNamespace(), obj.Spec.ProjectRef)
	if err != nil {
		return nil, err
	}
	projectID := projectRef.ProjectID
	if projectID == "" {
		return nil, fmt.Errorf("cannot resolve project")
	}
{{- if ne .ParentStyle "project" }}
	location := obj.Spec.Location
{{- end }}

	// Get desired ID
	resourceID := common.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	// Use approved External
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		actualParent, actualResourceID, err := Parse{{.ProtoMessageName}}External(externalRef)
		if err != nil {
			return nil, err
		}
		if actualParent.ProjectID != projectID {
			return nil, fmt.Errorf("spec.projectRef changed, expect %s, got %s", actualParent.ProjectID, projectID)
		}
{{- if ne .ParentStyle "project" }}
		if actualParent.Location != location {
			return nil, fmt.Errorf("spec.location changed, expect %s, got %s", actualParent.Location, location)
		}
{{- end }}
		if actualResourceID != resourceID {
			return nil, fmt.Errorf("cannot reset ` + "`" + `metadata.name` + "`" + ` or ` + "`" + `spec.resourceID` + "`" + ` to %s, since it has already assigned to %s",
				resourceID, actualResourceID)
		}
	}
	return &{{.ProtoMessageName}}Identity{
		parent: &{{.ProtoMessageName}}Parent{
			ProjectID: projectID,
{{- if ne .ParentStyle "project" }}
			Location:  location,
{{- end }}
		},
		id: resourceID,
	}, nil
}

func Parse{{.ProtoMessageName}}External(external string) (parent *{{.ProtoMessageName}}Parent, resourceID string, err error) {
	tokens := strings.Split(external, "/")
{{- if eq .ParentStyle "project" }}
	if len(tokens) != 4 || tokens[0] != "projects" || tokens[2] != "{{.Collection}}" {
		return nil, "", fmt.Errorf("format of {{.Kind}} external=%q was not known (use projects/{{"{{"}}projectID{{"}}"}}/{{.Collection}}/{{"{{"}}{{.ProtoMessageName | ToLower }}ID{{"}}"}})", external)
	}
	parent = &{{.ProtoMessageName}}Parent{
		ProjectID: tokens[1],
	}
	resourceID = tokens[3]
{{- else }}
	if len(tokens) != 6 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "{{.Collection}}" {
		return nil, "", fmt.Errorf("format of {{.Kind}} external=%q was not known (use projects/{{"{{"}}projectID{{"}}"}}/locations/{{"{{"}}location{{"}}"}}/{{.Collection}}/{{"{{"}}{{.ProtoMessageName | ToLower }}ID{{"}}"}})", external)
	}
	parent = &{{.ProtoMessageName}}Parent{
		ProjectID: tokens[1],
		Location:  tokens[3],
	}
	resourceID = tokens[5]
{{- end }}
	return parent, resourceID, nil
}
`
