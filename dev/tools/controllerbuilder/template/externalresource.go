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

const (
	// TODO(user): Add service domain
	serviceDomain = "//{{.KindToLower}}.googleapis.com"
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

// ExternalRef builds a externalRef from a {{.Kind}}
func (c *{{.Kind}}Identity) ExternalRef() *string {
	e := serviceDomain + "/" + c.FullyQualifiedName()
	return &e
}

// BuildIDFromExternal builds a {{.Kind}}Identity from a external reference
func BuildIDFromExternal(external string) (*{{.Kind}}Identity, error) {
	// TODO(user): Build resource identity from external reference
}

// BuildID builds a {{.Kind}}Identity from resource components
func BuildID(project, location string) *{{.Kind}}Identity {
	// TODO(user): Build resource identity from resource components, i.e. project, location, resource id
}
