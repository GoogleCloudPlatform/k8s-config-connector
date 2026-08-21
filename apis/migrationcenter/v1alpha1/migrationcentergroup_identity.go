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
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &MigrationCenterGroupIdentity{}
	_ identity.Resource   = &MigrationCenterGroup{}
)

var MigrationCenterGroupIdentityFormat = gcpurls.Template[MigrationCenterGroupIdentity]("migrationcenter.googleapis.com", "projects/{project}/locations/{location}/groups/{group}")

// MigrationCenterGroupIdentity is the identity of a GCP MigrationCenterGroup resource.
// +k8s:deepcopy-gen=false
type MigrationCenterGroupIdentity struct {
	Project  string
	Location string
	Group    string
}

func (i *MigrationCenterGroupIdentity) String() string {
	return MigrationCenterGroupIdentityFormat.ToString(*i)
}

func (i *MigrationCenterGroupIdentity) FromExternal(ref string) error {
	parsed, match, err := MigrationCenterGroupIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of MigrationCenterGroup external=%q was not known (use %s): %w", ref, MigrationCenterGroupIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of MigrationCenterGroup external=%q was not known (use %s)", ref, MigrationCenterGroupIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *MigrationCenterGroupIdentity) Host() string {
	return MigrationCenterGroupIdentityFormat.Host()
}

func (i *MigrationCenterGroupIdentity) ParentString() string {
	return fmt.Sprintf("projects/%s/locations/%s", i.Project, i.Location)
}

func getIdentityFromMigrationCenterGroupSpec(ctx context.Context, reader client.Reader, obj *MigrationCenterGroup) (*MigrationCenterGroupIdentity, error) {
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

	identity := &MigrationCenterGroupIdentity{
		Project:  projectID,
		Location: location,
		Group:    resourceID,
	}
	return identity, nil
}

func (obj *MigrationCenterGroup) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromMigrationCenterGroupSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		if !strings.Contains(externalRef, "/") {
			if externalRef != specIdentity.Group {
				return nil, fmt.Errorf("cannot change MigrationCenterGroup identity (old=%q, new=%q)", externalRef, specIdentity.Group)
			}
		} else {
			statusIdentity := &MigrationCenterGroupIdentity{}
			if err := statusIdentity.FromExternal(externalRef); err != nil {
				return nil, err
			}

			if statusIdentity.String() != specIdentity.String() {
				return nil, fmt.Errorf("cannot change MigrationCenterGroup identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
			}
		}
	}

	return specIdentity, nil
}
