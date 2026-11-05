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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &HealthcareConsentStoreIdentity{}
	_ identity.Resource   = &HealthcareConsentStore{}
)

var HealthcareConsentStoreIdentityFormat = gcpurls.Template[HealthcareConsentStoreIdentity](
	"healthcare.googleapis.com",
	"projects/{project}/locations/{location}/datasets/{dataset}/consentStores/{consentStore}",
)

// HealthcareConsentStoreIdentity is the identity of a GCP HealthcareConsentStore resource.
// +k8s:deepcopy-gen=false
type HealthcareConsentStoreIdentity struct {
	Project      string
	Location     string
	Dataset      string
	ConsentStore string
}

func (i *HealthcareConsentStoreIdentity) String() string {
	return HealthcareConsentStoreIdentityFormat.ToString(*i)
}

func (i *HealthcareConsentStoreIdentity) FromExternal(ref string) error {
	parsed, match, err := HealthcareConsentStoreIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of HealthcareConsentStore external=%q was not known (use %s): %w", ref, HealthcareConsentStoreIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of HealthcareConsentStore external=%q was not known (use %s)", ref, HealthcareConsentStoreIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *HealthcareConsentStoreIdentity) Host() string {
	return HealthcareConsentStoreIdentityFormat.Host()
}

func (i *HealthcareConsentStoreIdentity) ParentString() string {
	return fmt.Sprintf("projects/%s/locations/%s/datasets/%s", i.Project, i.Location, i.Dataset)
}

var DatasetIdentityFormat = gcpurls.Template[HealthcareDatasetIdentity](
	"healthcare.googleapis.com",
	"projects/{project}/locations/{location}/datasets/{dataset}",
)

// HealthcareDatasetIdentity is the identity of a GCP HealthcareDataset resource.
// +k8s:deepcopy-gen=false
type HealthcareDatasetIdentity struct {
	Project  string
	Location string
	Dataset  string
}

func (i *HealthcareDatasetIdentity) String() string {
	return DatasetIdentityFormat.ToString(*i)
}

func (i *HealthcareDatasetIdentity) FromExternal(ref string) error {
	parsed, match, err := DatasetIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of HealthcareDataset external=%q was not known (use %s): %w", ref, DatasetIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of HealthcareDataset external=%q was not known (use %s)", ref, DatasetIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func getIdentityFromHealthcareConsentStoreSpec(ctx context.Context, reader client.Reader, obj *HealthcareConsentStore) (*HealthcareConsentStoreIdentity, error) {
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

	return &HealthcareConsentStoreIdentity{
		Project:      datasetID.Project,
		Location:     datasetID.Location,
		Dataset:      datasetID.Dataset,
		ConsentStore: resourceID,
	}, nil
}

func (obj *HealthcareConsentStore) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	return getIdentityFromHealthcareConsentStoreSpec(ctx, reader, obj)
}
