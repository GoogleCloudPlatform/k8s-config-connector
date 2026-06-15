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
	_ identity.IdentityV2 = &StorageInsightsDatasetConfigIdentity{}
	_ identity.Resource   = &StorageInsightsDatasetConfig{}
)

var StorageInsightsDatasetConfigIdentityFormat = gcpurls.Template[StorageInsightsDatasetConfigIdentity]("storageinsights.googleapis.com", "projects/{project}/locations/{location}/datasetConfigs/{datasetconfig}")

// StorageInsightsDatasetConfigIdentity is the identity of a GCP StorageInsightsDatasetConfig resource.
// +k8s:deepcopy-gen=false
type StorageInsightsDatasetConfigIdentity struct {
	Project       string
	Location      string
	DatasetConfig string
}

func (i *StorageInsightsDatasetConfigIdentity) String() string {
	return StorageInsightsDatasetConfigIdentityFormat.ToString(*i)
}

func (i *StorageInsightsDatasetConfigIdentity) FromExternal(ref string) error {
	parsed, match, err := StorageInsightsDatasetConfigIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of StorageInsightsDatasetConfig external=%q was not known (use %s): %w", ref, StorageInsightsDatasetConfigIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of StorageInsightsDatasetConfig external=%q was not known (use %s)", ref, StorageInsightsDatasetConfigIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *StorageInsightsDatasetConfigIdentity) Host() string {
	return StorageInsightsDatasetConfigIdentityFormat.Host()
}

func (i *StorageInsightsDatasetConfigIdentity) ParentString() string {
	return fmt.Sprintf("projects/%s/locations/%s", i.Project, i.Location)
}

func getIdentityFromStorageInsightsDatasetConfigSpec(ctx context.Context, reader client.Reader, obj *StorageInsightsDatasetConfig) (*StorageInsightsDatasetConfigIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID: %w", err)
	}

	location, err := refs.GetLocation(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve location: %w", err)
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project: %w", err)
	}

	identity := &StorageInsightsDatasetConfigIdentity{
		Project:       projectID,
		Location:      location,
		DatasetConfig: resourceID,
	}
	return identity, nil
}

func (obj *StorageInsightsDatasetConfig) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromStorageInsightsDatasetConfigSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	if obj.Status.ExternalRef != nil && *obj.Status.ExternalRef != "" {
		statusIdentity := &StorageInsightsDatasetConfigIdentity{}
		if err := statusIdentity.FromExternal(*obj.Status.ExternalRef); err != nil {
			return nil, err
		}
		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change StorageInsightsDatasetConfig identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
