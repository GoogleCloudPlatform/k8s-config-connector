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

package bigquerydatatransfer

import (
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigquerydatatransfer/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

type BigQueryDataTransferConfigIdentity struct {
	projectID        string
	location         string
	transferConfigID string
}

// Parent builds a BigQueryDataTransferConfig parent
func (c *BigQueryDataTransferConfigIdentity) Parent() string {
	return fmt.Sprintf("projects/%s/locations/%s", c.projectID, c.location)
}

// FullyQualifiedName builds a BigQueryDataTransferConfig resource fully qualified name
func (c *BigQueryDataTransferConfigIdentity) FullyQualifiedName() string {
	return fmt.Sprintf("projects/%s/locations/%s/transferConfigs/%s", c.projectID, c.location, c.transferConfigID)
}

// AsExternalRef builds a externalRef from a BigQueryDataTransferConfig
func (c *BigQueryDataTransferConfigIdentity) AsExternalRef() *string {
	e := serviceDomain + "/" + c.FullyQualifiedName()
	return &e
}

// asID builds a BigQueryDataTransferConfigIdentity from a external reference
func asID(externalRef string) (*BigQueryDataTransferConfigIdentity, error) {
	if !strings.HasPrefix(externalRef, serviceDomain) {
		return nil, fmt.Errorf("externalRef should have prefix %s, got %s", serviceDomain, externalRef)
	}
	path := strings.TrimPrefix(externalRef, serviceDomain+"/")
	tokens := strings.Split(path, "/")
	if len(tokens) != 6 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "transferConfigs" {
		return nil, fmt.Errorf("externalRef should be %s/projects/<project>/locations/<location>/transferConfigs/<transferConfigID>, got %s",
			serviceDomain, externalRef)
	}
	return &BigQueryDataTransferConfigIdentity{
		projectID:        tokens[1],
		location:         tokens[3],
		transferConfigID: tokens[5],
	}, nil
}

// BuildID builds a unique identifier BigQueryDataTransferConfigIdentity from resource components
func BuildID(projectID, location, transferConfigID string) *BigQueryDataTransferConfigIdentity {
	return &BigQueryDataTransferConfigIdentity{
		projectID:        projectID,
		location:         location,
		transferConfigID: transferConfigID,
	}
}

// parseServiceGeneratedIDFromName extracts the service generated UUID from the name field of the resource. e.g. "projects/{project_id}/locations/{region}/transferConfigs/{config_id}"
func parseServiceGeneratedIDFromName(s string) (string, error) {
	tokens := strings.Split(s, "/")
	if len(tokens) != 6 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "transferConfigs" {
		return "", fmt.Errorf("service generated name should have format projects/<project>/locations/<location>/transferConfigs/<transferConfigID>, got %s", s)
	}
	return tokens[5], nil
}

// parseServiceGeneratedIDFromExternalRef extracts the service generated UUID from the externalRef.
func parseServiceGeneratedIDFromExternalRef(obj *krm.BigQueryDataTransferConfig) (string, error) {
	if obj == nil || obj.Status.ExternalRef == nil {
		return "", nil // it is OK to have "" resource ID prior to resource creation.
	}
	s := direct.ValueOf(obj.Status.ExternalRef)
	s = strings.TrimPrefix(s, serviceDomain+"/")
	tokens := strings.Split(s, "/")
	if len(tokens) != 6 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "transferConfigs" {
		return "", fmt.Errorf("externalRef should be %s/projects/<project>/locations/<location>/transferConfigs/<transferConfigID>, got %s",
			serviceDomain, s)
	}
	return tokens[5], nil
}
