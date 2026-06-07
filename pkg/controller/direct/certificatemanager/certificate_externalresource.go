// Copyright 2026 Google LLC
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

package certificatemanager

import (
	"fmt"
	"strings"
)

// The Identifier for ConfigConnector to track the CertificateManagerCertificate resource from the GCP service.
type CertificateManagerCertificateIdentity struct {
	Parent      *parent
	Certificate string
}

// FullyQualifiedName returns both parent and resource ID in the full url format.
func (c *CertificateManagerCertificateIdentity) FullyQualifiedName() string {
	return fmt.Sprintf("%s/certificates/%s", c.Parent.String(), c.Certificate)
}

// AsExternalRef builds an externalRef from a CertificateManagerCertificate
func (c *CertificateManagerCertificateIdentity) AsExternalRef() *string {
	e := serviceDomain + "/" + c.FullyQualifiedName()
	return &e
}

// asCertificateID builds a CertificateManagerCertificateIdentity from a `status.externalRef`
func asCertificateID(externalRef string) (*CertificateManagerCertificateIdentity, error) {
	if !strings.HasPrefix(externalRef, serviceDomain) {
		return nil, fmt.Errorf("externalRef should have prefix %s, got %s", serviceDomain, externalRef)
	}
	path := strings.TrimPrefix(externalRef, serviceDomain+"/")
	tokens := strings.Split(path, "/")

	if len(tokens) != 6 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "certificates" {
		return nil, fmt.Errorf("externalRef should be %s/projects/<project>/locations/<location>/certificates/<certificate>, got %s",
			serviceDomain, externalRef)
	}
	return &CertificateManagerCertificateIdentity{
		Parent:      &parent{Project: tokens[1], Location: tokens[3]},
		Certificate: tokens[5],
	}, nil
}

// BuildCertificateID builds the ID for ConfigConnector to track the CertificateManagerCertificate resource from the GCP service.
func BuildCertificateID(project, location, resourceID string) *CertificateManagerCertificateIdentity {
	return &CertificateManagerCertificateIdentity{
		Parent:      &parent{Project: project, Location: location},
		Certificate: resourceID,
	}
}
