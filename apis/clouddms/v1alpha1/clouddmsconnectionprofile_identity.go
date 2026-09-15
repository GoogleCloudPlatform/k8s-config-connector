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
	_ identity.IdentityV2 = &CloudDMSConnectionProfileIdentity{}
	_ identity.Resource   = &CloudDMSConnectionProfile{}
)

var CloudDMSConnectionProfileIdentityFormat = gcpurls.Template[CloudDMSConnectionProfileIdentity]("datamigration.googleapis.com", "projects/{project}/locations/{location}/connectionProfiles/{connectionprofile}")

// CloudDMSConnectionProfileIdentity is the identity of a GCP CloudDMSConnectionProfile resource.
// +k8s:deepcopy-gen=false
type CloudDMSConnectionProfileIdentity struct {
	Project           string
	Location          string
	ConnectionProfile string
}

func (i *CloudDMSConnectionProfileIdentity) String() string {
	return CloudDMSConnectionProfileIdentityFormat.ToString(*i)
}

func (i *CloudDMSConnectionProfileIdentity) FromExternal(ref string) error {
	parsed, match, err := CloudDMSConnectionProfileIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of CloudDMSConnectionProfile external=%q was not known (use %s): %w", ref, CloudDMSConnectionProfileIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of CloudDMSConnectionProfile external=%q was not known (use %s)", ref, CloudDMSConnectionProfileIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *CloudDMSConnectionProfileIdentity) Host() string {
	return CloudDMSConnectionProfileIdentityFormat.Host()
}

func getIdentityFromCloudDMSConnectionProfileSpec(ctx context.Context, reader client.Reader, obj *CloudDMSConnectionProfile) (*CloudDMSConnectionProfileIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	location, err := refs.GetLocation(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve location")
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project")
	}

	identity := &CloudDMSConnectionProfileIdentity{
		Project:           projectID,
		Location:          location,
		ConnectionProfile: resourceID,
	}
	return identity, nil
}

func (obj *CloudDMSConnectionProfile) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromCloudDMSConnectionProfileSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		actualIdentity := &CloudDMSConnectionProfileIdentity{}
		if err := actualIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}
		if actualIdentity.Project != specIdentity.Project {
			return nil, fmt.Errorf("spec.projectRef changed, expect %s, got %s", actualIdentity.Project, specIdentity.Project)
		}
		if actualIdentity.Location != specIdentity.Location {
			return nil, fmt.Errorf("spec.location changed, expect %s, got %s", actualIdentity.Location, specIdentity.Location)
		}
		if actualIdentity.ConnectionProfile != specIdentity.ConnectionProfile {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
				specIdentity.ConnectionProfile, actualIdentity.ConnectionProfile)
		}
	}
	return specIdentity, nil
}
