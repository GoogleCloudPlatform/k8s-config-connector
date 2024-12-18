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

package firestore

import (
	"fmt"
	"strings"
)

type FirestoreDatabaseIdentity struct {
	project           string
	location          string
	firestoredatabase string
}

// Parent builds a FirestoreDatabase parent
func (c *FirestoreDatabaseIdentity) Parent() string {
	return "projects/" + c.project
}

// FullyQualifiedName builds a FirestoreDatabase resource fully qualified name
func (c *FirestoreDatabaseIdentity) FullyQualifiedName() string {
	return c.Parent() + "/databases/" + c.firestoredatabase
}

// AsExternalRef builds a externalRef from a FirestoreDatabase
func (c *FirestoreDatabaseIdentity) AsExternalRef() *string {
	e := serviceDomain + "/" + c.FullyQualifiedName()
	return &e
}

// asID builds a FirestoreDatabaseIdentity from a external reference
func asID(externalRef string) (*FirestoreDatabaseIdentity, error) {
	if !strings.HasPrefix(externalRef, serviceDomain) {
		return nil, fmt.Errorf("externalRef should have prefix %s, got %s", serviceDomain, externalRef)
	}
	path := strings.TrimPrefix(externalRef, serviceDomain+"/")
	tokens := strings.Split(path, "/")

	if len(tokens) != 4 || tokens[0] != "projects" || tokens[2] != "databases" {
		return nil, fmt.Errorf("externalRef should be %s/projects/{{projectID}}/databases/{{firestoredatabase}}, got %s",
			serviceDomain, externalRef)
	}
	return &FirestoreDatabaseIdentity{
		project:           tokens[1],
		firestoredatabase: tokens[3],
	}, nil
}

// BuildID builds a unique identifier FirestoreDatabaseIdentity from resource components
func BuildID(project, firestoredatabase string) *FirestoreDatabaseIdentity {
	return &FirestoreDatabaseIdentity{
		project:           project,
		firestoredatabase: firestoredatabase,
	}
}
