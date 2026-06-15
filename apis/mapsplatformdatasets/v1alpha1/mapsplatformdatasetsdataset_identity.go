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
	_ identity.IdentityV2 = &MapsPlatformDatasetsDatasetIdentity{}
	_ identity.Resource   = &MapsPlatformDatasetsDataset{}
)

var MapsPlatformDatasetsDatasetIdentityFormat = gcpurls.Template[MapsPlatformDatasetsDatasetIdentity]("mapsplatformdatasets.googleapis.com", "projects/{project}/datasets/{dataset}")

// +k8s:deepcopy-gen=false
type MapsPlatformDatasetsDatasetIdentity struct {
	Project string
	Dataset string
}

func (i *MapsPlatformDatasetsDatasetIdentity) String() string {
	return MapsPlatformDatasetsDatasetIdentityFormat.ToString(*i)
}

func (i *MapsPlatformDatasetsDatasetIdentity) FromExternal(ref string) error {
	parsed, match, err := MapsPlatformDatasetsDatasetIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of MapsPlatformDatasetsDataset external=%q was not known (use %s): %w", ref, MapsPlatformDatasetsDatasetIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of MapsPlatformDatasetsDataset external=%q was not known (use %s)", ref, MapsPlatformDatasetsDatasetIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *MapsPlatformDatasetsDatasetIdentity) Host() string {
	return MapsPlatformDatasetsDatasetIdentityFormat.Host()
}

func getIdentityFromMapsPlatformDatasetsDatasetSpec(ctx context.Context, reader client.Reader, obj client.Object) (*MapsPlatformDatasetsDatasetIdentity, error) {
	_, ok := obj.(*MapsPlatformDatasetsDataset)
	if !ok {
		return nil, fmt.Errorf("object is not a MapsPlatformDatasetsDataset")
	}
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project")
	}

	identity := &MapsPlatformDatasetsDatasetIdentity{
		Project: projectID,
		Dataset: resourceID,
	}
	return identity, nil
}

func (obj *MapsPlatformDatasetsDataset) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromMapsPlatformDatasetsDatasetSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &MapsPlatformDatasetsDatasetIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change MapsPlatformDatasetsDataset identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
