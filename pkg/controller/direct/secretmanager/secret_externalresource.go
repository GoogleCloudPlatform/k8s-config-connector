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

package secretmanager

import (
	"fmt"
	"strings"
)

// The Identifier for ConfigConnector to track the SecretManagerSecret resource from the GCP service.
type SecretManagerSecretIdentity struct {
	Parent *parent
	Secret string
}

type parent struct {
	Project string
}

func (p *parent) String() string {
	return fmt.Sprintf("projects/%s", p.Project)
}

// FullyQualifiedName returns both parent and resource ID in the full url format.
func (c *SecretManagerSecretIdentity) FullyQualifiedName() string {
	return fmt.Sprintf("%s/secrets/%s", c.Parent, c.Secret)
}

// AsExternalRef builds a externalRef from a SecretManagerSecret
func (c *SecretManagerSecretIdentity) AsExternalRef() *string {
	e := serviceDomain + "/" + c.FullyQualifiedName()
	return &e
}

// asID builds a SecretManagerSecretIdentity from a `status.externalRef`
func asID(externalRef string) (*SecretManagerSecretIdentity, error) {
	if !strings.HasPrefix(externalRef, serviceDomain) {
		return nil, fmt.Errorf("externalRef should have prefix %s, got %s", serviceDomain, externalRef)
	}
	path := strings.TrimPrefix(externalRef, serviceDomain+"/")
	tokens := strings.Split(path, "/")

	if len(tokens) != 4 || tokens[0] != "projects" || tokens[2] != "secrets" {
		return nil, fmt.Errorf("externalRef should be %s/projects/<project>/secrets/<Secret>, got %s",
			serviceDomain, externalRef)
	}
	return &SecretManagerSecretIdentity{
		Parent: &parent{Project: tokens[1]},
		Secret: tokens[3],
	}, nil
}

// BuildID builds the ID for ConfigConnector to track the SecretManagerSecret resource from the GCP service.
func BuildID(project, resourceID string) *SecretManagerSecretIdentity {
	return &SecretManagerSecretIdentity{
		Parent: &parent{Project: project},
		Secret: resourceID,
	}
}
