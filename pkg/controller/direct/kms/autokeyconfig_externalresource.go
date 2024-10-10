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

package kms

import (
	"fmt"
	"strings"
)

// The Identifier for ConfigConnector to track the KMSAutokeyConfig resource from the GCP service.
type KMSAutokeyConfigIdentity struct {
	Parent        *parent
}

type parent struct {
	Folder  string
}

func (p *parent) String() string {
	return fmt.Sprintf("folders/%s", p.Folder)
}

// FullyQualifiedName returns both parent and resource ID in the full url format.
func (c *KMSAutokeyConfigIdentity) FullyQualifiedName() string {
	return fmt.Sprintf("%s/autokeyconfig", c.Parent)
}

// AsExternalRef builds a externalRef from a KMSAutokeyConfig
func (c *KMSAutokeyConfigIdentity) AsExternalRef() *string {
	e := serviceDomain + "/" + c.FullyQualifiedName()
	return &e
}

// asID builds a KMSAutokeyConfigIdentity from a `status.externalRef`
func asID(externalRef string) (*KMSAutokeyConfigIdentity, error) {
	// TODO(user): Build resource identity from external reference
	if !strings.HasPrefix(externalRef, serviceDomain) {
		return nil, fmt.Errorf("externalRef should have prefix %s, got %s", serviceDomain, externalRef)
	}
	path := strings.TrimPrefix(externalRef, serviceDomain+"/")
	tokens := strings.Split(path, "/")

	// TODO(user): Confirm the format of your resources, and verify it like the example below
	if len(tokens) != 3 || tokens[0] != "folders" || tokens[2] != "autokeyConfig" {
		return nil, fmt.Errorf("externalRef should be %s/projects/<project>/locations/<location>/autokeyconfigs/<AutokeyConfig>, got %s",
			serviceDomain, externalRef)
	}
	return &KMSAutokeyConfigIdentity{
		Parent:        &parent{Folder: tokens[1]},
	}, nil
}

// BuildID builds the ID for ConfigConnector to track the KMSAutokeyConfig resource from the GCP service.
func BuildID(folder string) *KMSAutokeyConfigIdentity {
	// TODO(user): Build resource identity from resource components, i.e. project, location, resource id
	return &KMSAutokeyConfigIdentity{
		Parent:        &parent{Folder: folder},
	}
}
