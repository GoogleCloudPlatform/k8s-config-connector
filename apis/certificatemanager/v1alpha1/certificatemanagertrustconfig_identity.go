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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &CertificateManagerTrustConfigIdentity{}
	_ identity.Resource   = &CertificateManagerTrustConfig{}
)

var CertificateManagerTrustConfigIdentityFormat = gcpurls.Template[CertificateManagerTrustConfigIdentity]("certificatemanager.googleapis.com", "projects/{project}/locations/{location}/trustConfigs/{trustConfig}")

// CertificateManagerTrustConfigIdentity is the identity of a GCP CertificateManagerTrustConfig resource.
// +k8s:deepcopy-gen=false
type CertificateManagerTrustConfigIdentity struct {
	Project     string
	Location    string
	TrustConfig string
}

func (i *CertificateManagerTrustConfigIdentity) String() string {
	return CertificateManagerTrustConfigIdentityFormat.ToString(*i)
}

func (i *CertificateManagerTrustConfigIdentity) FromExternal(ref string) error {
	parsed, match, err := CertificateManagerTrustConfigIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of CertificateManagerTrustConfig external=%q was not known (use %s): %w", ref, CertificateManagerTrustConfigIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of CertificateManagerTrustConfig external=%q was not known (use %s)", ref, CertificateManagerTrustConfigIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *CertificateManagerTrustConfigIdentity) Host() string {
	return CertificateManagerTrustConfigIdentityFormat.Host()
}

func getIdentityFromCertificateManagerTrustConfigSpec(ctx context.Context, reader client.Reader, obj *CertificateManagerTrustConfig) (*CertificateManagerTrustConfigIdentity, error) {
	if obj.Spec.ProjectRef == nil {
		return nil, fmt.Errorf("spec.projectRef must be specified")
	}
	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("resolving projectRef: %w", err)
	}

	location := common.ValueOf(obj.Spec.Location)
	if location == "" {
		return nil, fmt.Errorf("spec.location must be specified")
	}

	// Resolve user-configured ID
	resourceID := common.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	identity := &CertificateManagerTrustConfigIdentity{
		Project:     projectID,
		Location:    location,
		TrustConfig: resourceID,
	}
	return identity, nil
}

func (obj *CertificateManagerTrustConfig) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromCertificateManagerTrustConfigSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &CertificateManagerTrustConfigIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change CertificateManagerTrustConfig identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
