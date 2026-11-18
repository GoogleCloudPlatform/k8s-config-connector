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
	_ identity.IdentityV2 = &HealthcareHL7V2StoreIdentity{}
	_ identity.Resource   = &HealthcareHL7V2Store{}
)

var HealthcareHL7V2StoreIdentityFormat = gcpurls.Template[HealthcareHL7V2StoreIdentity](
	"healthcare.googleapis.com",
	"projects/{project}/locations/{location}/datasets/{dataset}/hl7V2Stores/{hl7V2Store}",
)

// HealthcareHL7V2StoreIdentity is the identity of a GCP HealthcareHL7V2Store resource.
// +k8s:deepcopy-gen=false
type HealthcareHL7V2StoreIdentity struct {
	Project    string
	Location   string
	Dataset    string
	Hl7V2Store string
}

func (i *HealthcareHL7V2StoreIdentity) String() string {
	return HealthcareHL7V2StoreIdentityFormat.ToString(*i)
}

func (i *HealthcareHL7V2StoreIdentity) FromExternal(ref string) error {
	parsed, match, err := HealthcareHL7V2StoreIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of HealthcareHL7V2Store external=%q was not known (use %s): %w", ref, HealthcareHL7V2StoreIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of HealthcareHL7V2Store external=%q was not known (use %s)", ref, HealthcareHL7V2StoreIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *HealthcareHL7V2StoreIdentity) Host() string {
	return HealthcareHL7V2StoreIdentityFormat.Host()
}

func (i *HealthcareHL7V2StoreIdentity) ParentString() string {
	return fmt.Sprintf("projects/%s/locations/%s/datasets/%s", i.Project, i.Location, i.Dataset)
}

func getIdentityFromHealthcareHL7V2StoreSpec(ctx context.Context, reader client.Reader, obj *HealthcareHL7V2Store) (*HealthcareHL7V2StoreIdentity, error) {
	if obj.Spec.Dataset == "" {
		return nil, fmt.Errorf("spec.dataset is required")
	}

	datasetID := &HealthcareDatasetIdentity{}
	if err := datasetID.FromExternal(obj.Spec.Dataset); err != nil {
		return nil, fmt.Errorf("parsing spec.dataset %q: %w", obj.Spec.Dataset, err)
	}

	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID: %w", err)
	}

	return &HealthcareHL7V2StoreIdentity{
		Project:    datasetID.Project,
		Location:   datasetID.Location,
		Dataset:    datasetID.Dataset,
		Hl7V2Store: resourceID,
	}, nil
}

func (obj *HealthcareHL7V2Store) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromHealthcareHL7V2StoreSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		statusIdentity := &HealthcareHL7V2StoreIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if specIdentity.Hl7V2Store == "" {
			specIdentity.Hl7V2Store = statusIdentity.Hl7V2Store
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change HealthcareHL7V2Store identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
