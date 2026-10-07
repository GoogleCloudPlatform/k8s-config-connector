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
	_ identity.IdentityV2 = &StorageInsightsReportConfigIdentity{}
	_ identity.Resource   = &StorageInsightsReportConfig{}
)

var StorageInsightsReportConfigIdentityFormat = gcpurls.Template[StorageInsightsReportConfigIdentity]("storageinsights.googleapis.com", "projects/{project}/locations/{location}/reportConfigs/{reportConfig}")

// StorageInsightsReportConfigIdentity is the identity of a GCP StorageInsightsReportConfig resource.
// +k8s:deepcopy-gen=false
type StorageInsightsReportConfigIdentity struct {
	Project      string
	Location     string
	ReportConfig string
}

func (i *StorageInsightsReportConfigIdentity) String() string {
	return StorageInsightsReportConfigIdentityFormat.ToString(*i)
}

func (i *StorageInsightsReportConfigIdentity) FromExternal(ref string) error {
	parsed, match, err := StorageInsightsReportConfigIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of StorageInsightsReportConfig external=%q was not known (use %s): %w", ref, StorageInsightsReportConfigIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of StorageInsightsReportConfig external=%q was not known (use %s)", ref, StorageInsightsReportConfigIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *StorageInsightsReportConfigIdentity) Host() string {
	return StorageInsightsReportConfigIdentityFormat.Host()
}

func (i *StorageInsightsReportConfigIdentity) ParentString() string {
	return fmt.Sprintf("projects/%s/locations/%s", i.Project, i.Location)
}

func getIdentityFromStorageInsightsReportConfigSpec(ctx context.Context, reader client.Reader, obj *StorageInsightsReportConfig) (*StorageInsightsReportConfigIdentity, error) {
	// StorageInsightsReportConfig only supports service-generated IDs (GCP assigns a UUID upon creation).
	resourceID := common.ValueOf(obj.Spec.ResourceID)

	location, err := refs.GetLocation(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve location")
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project")
	}

	identity := &StorageInsightsReportConfigIdentity{
		Project:      projectID,
		Location:     location,
		ReportConfig: resourceID,
	}
	return identity, nil
}

func (obj *StorageInsightsReportConfig) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromStorageInsightsReportConfigSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		statusIdentity := &StorageInsightsReportConfigIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if specIdentity.ReportConfig == "" {
			if statusIdentity.Project != specIdentity.Project || statusIdentity.Location != specIdentity.Location {
				return nil, fmt.Errorf("cannot change StorageInsightsReportConfig parent (old parent=%s/%s, new parent=%s/%s)", statusIdentity.Project, statusIdentity.Location, specIdentity.Project, specIdentity.Location)
			}
			return statusIdentity, nil
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change StorageInsightsReportConfig identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
