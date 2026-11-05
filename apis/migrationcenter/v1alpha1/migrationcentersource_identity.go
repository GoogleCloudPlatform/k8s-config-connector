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
	_ identity.IdentityV2 = &MigrationCenterSourceIdentity{}
	_ identity.Resource   = &MigrationCenterSource{}
)

var MigrationCenterSourceIdentityFormat = gcpurls.Template[MigrationCenterSourceIdentity]("migrationcenter.googleapis.com", "projects/{project}/locations/{location}/sources/{source}")

// MigrationCenterSourceIdentity is the identity of a GCP MigrationCenterSource resource.
// +k8s:deepcopy-gen=false
type MigrationCenterSourceIdentity struct {
	Project  string
	Location string
	Source   string
}

func (i *MigrationCenterSourceIdentity) String() string {
	return MigrationCenterSourceIdentityFormat.ToString(*i)
}

func (i *MigrationCenterSourceIdentity) FromExternal(ref string) error {
	parsed, match, err := MigrationCenterSourceIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of MigrationCenterSource external=%q was not known (use %s): %w", ref, MigrationCenterSourceIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of MigrationCenterSource external=%q was not known (use %s)", ref, MigrationCenterSourceIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *MigrationCenterSourceIdentity) Host() string {
	return MigrationCenterSourceIdentityFormat.Host()
}

func (i *MigrationCenterSourceIdentity) ParentString() string {
	return fmt.Sprintf("projects/%s/locations/%s", i.Project, i.Location)
}

func getIdentityFromMigrationCenterSourceSpec(ctx context.Context, reader client.Reader, obj *MigrationCenterSource) (*MigrationCenterSourceIdentity, error) {
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

	sourceIdentity := &MigrationCenterSourceIdentity{
		Project:  projectID,
		Location: location,
		Source:   resourceID,
	}
	return sourceIdentity, nil
}

func (obj *MigrationCenterSource) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromMigrationCenterSourceSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		if !strings.Contains(externalRef, "/") {
			if externalRef != specIdentity.Source {
				return nil, fmt.Errorf("cannot change MigrationCenterSource identity (old=%q, new=%q)", externalRef, specIdentity.Source)
			}
		} else {
			statusIdentity := &MigrationCenterSourceIdentity{}
			if err := statusIdentity.FromExternal(externalRef); err != nil {
				return nil, err
			}

			if statusIdentity.String() != specIdentity.String() {
				return nil, fmt.Errorf("cannot change MigrationCenterSource identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
			}
		}
	}

	return specIdentity, nil
}
