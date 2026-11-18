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
	_ identity.IdentityV2 = &HealthcareDICOMStoreIdentity{}
	_ identity.Resource   = &HealthcareDICOMStore{}
)

var HealthcareDICOMStoreIdentityFormat = gcpurls.Template[HealthcareDICOMStoreIdentity](
	"healthcare.googleapis.com",
	"projects/{project}/locations/{location}/datasets/{dataset}/dicomStores/{dicomStore}",
)

// HealthcareDICOMStoreIdentity is the identity of a GCP HealthcareDICOMStore resource.
// +k8s:deepcopy-gen=false
type HealthcareDICOMStoreIdentity struct {
	Project    string
	Location   string
	Dataset    string
	DicomStore string
}

func (i *HealthcareDICOMStoreIdentity) String() string {
	return HealthcareDICOMStoreIdentityFormat.ToString(*i)
}

func (i *HealthcareDICOMStoreIdentity) FromExternal(ref string) error {
	parsed, match, err := HealthcareDICOMStoreIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of HealthcareDICOMStore external=%q was not known (use %s): %w", ref, HealthcareDICOMStoreIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of HealthcareDICOMStore external=%q was not known (use %s)", ref, HealthcareDICOMStoreIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *HealthcareDICOMStoreIdentity) Host() string {
	return HealthcareDICOMStoreIdentityFormat.Host()
}

func (i *HealthcareDICOMStoreIdentity) ParentString() string {
	return fmt.Sprintf("projects/%s/locations/%s/datasets/%s", i.Project, i.Location, i.Dataset)
}

func getIdentityFromHealthcareDICOMStoreSpec(ctx context.Context, reader client.Reader, obj *HealthcareDICOMStore) (*HealthcareDICOMStoreIdentity, error) {
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

	return &HealthcareDICOMStoreIdentity{
		Project:    datasetID.Project,
		Location:   datasetID.Location,
		Dataset:    datasetID.Dataset,
		DicomStore: resourceID,
	}, nil
}

func (obj *HealthcareDICOMStore) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromHealthcareDICOMStoreSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		statusIdentity := &HealthcareDICOMStoreIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if specIdentity.DicomStore == "" {
			specIdentity.DicomStore = statusIdentity.DicomStore
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change HealthcareDICOMStore identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
