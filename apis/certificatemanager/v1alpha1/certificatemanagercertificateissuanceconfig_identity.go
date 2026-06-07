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

package v1alpha1

import (
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &CertificateManagerCertificateIssuanceConfigIdentity{}
	_ identity.Resource   = &CertificateManagerCertificateIssuanceConfig{}
)

var CertificateManagerCertificateIssuanceConfigIdentityFormat = gcpurls.Template[CertificateManagerCertificateIssuanceConfigIdentity]("certificatemanager.googleapis.com", "projects/{project}/locations/{location}/certificateIssuanceConfigs/{certificateissuanceconfig}")

// +k8s:deepcopy-gen=false
type CertificateManagerCertificateIssuanceConfigIdentity struct {
	Project                   string
	Location                  string
	CertificateIssuanceConfig string
}

func (i *CertificateManagerCertificateIssuanceConfigIdentity) String() string {
	return CertificateManagerCertificateIssuanceConfigIdentityFormat.ToString(*i)
}

func (i *CertificateManagerCertificateIssuanceConfigIdentity) FromExternal(ref string) error {
	parsed, match, err := CertificateManagerCertificateIssuanceConfigIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of CertificateManagerCertificateIssuanceConfig external=%q was not known (use %s): %w", ref, CertificateManagerCertificateIssuanceConfigIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of CertificateManagerCertificateIssuanceConfig external=%q was not known (use %s)", ref, CertificateManagerCertificateIssuanceConfigIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *CertificateManagerCertificateIssuanceConfigIdentity) Host() string {
	return CertificateManagerCertificateIssuanceConfigIdentityFormat.Host()
}

func getIdentityFromCertificateManagerCertificateIssuanceConfigSpec(ctx context.Context, reader client.Reader, obj *CertificateManagerCertificateIssuanceConfig) (*CertificateManagerCertificateIssuanceConfigIdentity, error) {
	if obj.Spec.ProjectAndLocationRef == nil {
		return nil, fmt.Errorf("inline projectRef/location must be specified")
	}
	if obj.Spec.ProjectAndLocationRef.ProjectRef == nil {
		return nil, fmt.Errorf("spec.projectRef must be specified")
	}

	// Resolve user-configured Parent (project and location)
	p := &parent.ProjectAndLocationParent{}
	if err := obj.Spec.ProjectAndLocationRef.Build(ctx, reader, obj.GetNamespace(), p); err != nil {
		return nil, err
	}

	// Resolve user-configured ID
	resourceID := common.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	identity := &CertificateManagerCertificateIssuanceConfigIdentity{
		Project:                   p.ProjectID,
		Location:                  p.Location,
		CertificateIssuanceConfig: resourceID,
	}
	return identity, nil
}

func (obj *CertificateManagerCertificateIssuanceConfig) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromCertificateManagerCertificateIssuanceConfigSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &CertificateManagerCertificateIssuanceConfigIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change CertificateManagerCertificateIssuanceConfig identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
