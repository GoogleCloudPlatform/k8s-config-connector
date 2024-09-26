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

package controller

const ExternalResourceTemplate = `
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

package {{.KCCService}}

import (
	"fmt"
	"strings"
)

// The Identifier for ConfigConnector to track the {{.Kind}} resource from the GCP service.
type {{.Kind}}Identity struct {
	Parent *parent
	{{.ProtoResource}} string
}

type parent struct {
	Project string
	Location string
}

func (p *parent) String() string{
	return fmt.Sprintf("projects/%s/locations/%s", p.Project, p.Location)
}

// FullyQualifiedName returns both parent and resource ID in the full url format. 
func (c *{{.Kind}}Identity) FullyQualifiedName() string {
	// TODO(user): Edit the URL path
	return fmt.Sprintf("%s/{{.ProtoResource | ToLower}}s/%s", c.Parent, c.{{.ProtoResource}})
}

// AsExternalRef builds a externalRef from a {{.Kind}}
func (c *{{.Kind}}Identity) AsExternalRef() *string {
	e := serviceDomain + "/" + c.FullyQualifiedName()
	return &e
}

// asID builds a {{.Kind}}Identity from a ` + "`" + `status.externalRef` + "`" + `
func asID(externalRef string) (*{{.Kind}}Identity, error) {
	// TODO(user): Build resource identity from external reference
	if !strings.HasPrefix(externalRef, serviceDomain) {
		return nil, fmt.Errorf("externalRef should have prefix %s, got %s", serviceDomain, externalRef)
	}
	path := strings.TrimPrefix(externalRef, serviceDomain+"/")
	tokens := strings.Split(path, "/")

	// TODO(user): Confirm the format of your resources, and verify it like the example below
	if len(tokens) != 6 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "{{.ProtoResource | ToLower}}s" {
		return nil, fmt.Errorf("externalRef should be %s/projects/<project>/locations/<location>/{{.ProtoResource | ToLower}}s/<{{.ProtoResource}}>, got %s",
			serviceDomain, externalRef)
	}
	return &{{.Kind}}Identity{
		Parent: &parent{Project: tokens[1], Location: tokens[3]},
		{{.ProtoResource}}: tokens[5],
	}, nil
}

// BuildID builds the ID for ConfigConnector to track the {{.Kind}} resource from the GCP service.
func BuildID(project, location, resourceID string) *{{.Kind}}Identity {
	// TODO(user): Build resource identity from resource components, i.e. project, location, resource id
	return &{{.Kind}}Identity{
		Parent: &parent{Project: project, Location: location},
		{{.ProtoResource}}: resourceID,
	}
}
`
