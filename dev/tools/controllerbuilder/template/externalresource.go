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

package template

const ExternalResourceTemplate = `
package {{.Service}}

import (
	"fmt"
	"strings"
)

// TODO(user): Define resource identity
type {{.Kind}}Identity struct {
	project    string
	location   string
	{{.KindToLower}} string
}

// Parent builds a {{.Kind}} parent
func (c *{{.Kind}}Identity) Parent() string {
	// TODO(user): Define resource parent
}

// FullyQualifiedName builds a {{.Kind}} resource fully qualified name
func (c *{{.Kind}}Identity) FullyQualifiedName() string {
	// TODO(user): Define resource fully qualified name
}

// AsExternalRef builds a externalRef from a {{.Kind}}
func (c *{{.Kind}}Identity) AsExternalRef() *string {
	e := serviceDomain + "/" + c.FullyQualifiedName()
	return &e
}

// asID builds a {{.Kind}}Identity from a external reference
func asID(externalRef string) (*{{.Kind}}Identity, error) {
	// TODO(user): Build resource identity from external reference
	if !strings.HasPrefix(externalRef, serviceDomain) {
		return nil, fmt.Errorf("externalRef should have prefix %s, got %s", serviceDomain, externalRef)
	}
	path := strings.TrimPrefix(externalRef, serviceDomain+"/")
	tokens := strings.Split(path, "/")

	// TODO(user): Confirm the format of your resources, and verify it like the example below
	if len(tokens) != 6 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "{{.KindToLower}}s" {
		return nil, fmt.Errorf("externalRef should be %s/projects/<project>/locations/<location>/{{.KindToLower}}s/<{{.KindToLower}}>, got %s",
			serviceDomain, externalRef)
	}
	return &{{.Kind}}Identity{
		project:    tokens[1],
		location:   tokens[3],
		{{.KindToLower}}: tokens[5],
	}, nil
}

// BuildID builds a unique identifier {{.Kind}}Identity from resource components
func BuildID(project, location string) *{{.Kind}}Identity {
	// TODO(user): Build resource identity from resource components, i.e. project, location, resource id
	return &{{.Kind}}Identity{
		project:    project,
		location:   location,
		{{.KindToLower}}: {{.KindToLower}},
	}
}
`
