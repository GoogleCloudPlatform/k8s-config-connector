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
	_ identity.IdentityV2 = &DatabaseMigrationConversionWorkspaceIdentity{}
	_ identity.Resource   = &DatabaseMigrationConversionWorkspace{}
)

var DatabaseMigrationConversionWorkspaceIdentityFormat = gcpurls.Template[DatabaseMigrationConversionWorkspaceIdentity]("datamigration.googleapis.com", "projects/{project}/locations/{location}/conversionWorkspaces/{conversionWorkspace}")

// +k8s:deepcopy-gen=false
type DatabaseMigrationConversionWorkspaceIdentity struct {
	Project             string
	Location            string
	ConversionWorkspace string
}

func (i *DatabaseMigrationConversionWorkspaceIdentity) String() string {
	return DatabaseMigrationConversionWorkspaceIdentityFormat.ToString(*i)
}

func (i *DatabaseMigrationConversionWorkspaceIdentity) FromExternal(ref string) error {
	parsed, match, err := DatabaseMigrationConversionWorkspaceIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of DatabaseMigrationConversionWorkspace external=%q was not known (use %s): %w", ref, DatabaseMigrationConversionWorkspaceIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of DatabaseMigrationConversionWorkspace external=%q was not known (use %s)", ref, DatabaseMigrationConversionWorkspaceIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *DatabaseMigrationConversionWorkspaceIdentity) Host() string {
	return DatabaseMigrationConversionWorkspaceIdentityFormat.Host()
}

func getIdentityFromDatabaseMigrationConversionWorkspaceSpec(ctx context.Context, reader client.Reader, obj client.Object) (*DatabaseMigrationConversionWorkspaceIdentity, error) {
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

	identity := &DatabaseMigrationConversionWorkspaceIdentity{
		Project:             projectID,
		Location:            location,
		ConversionWorkspace: resourceID,
	}
	return identity, nil
}

func (obj *DatabaseMigrationConversionWorkspace) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromDatabaseMigrationConversionWorkspaceSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &DatabaseMigrationConversionWorkspaceIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change DatabaseMigrationConversionWorkspace identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
