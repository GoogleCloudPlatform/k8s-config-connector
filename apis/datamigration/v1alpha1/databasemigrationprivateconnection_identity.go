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
	_ identity.IdentityV2 = &DatabaseMigrationPrivateConnectionIdentity{}
	_ identity.Resource   = &DatabaseMigrationPrivateConnection{}
)

var DatabaseMigrationPrivateConnectionIdentityFormat = gcpurls.Template[DatabaseMigrationPrivateConnectionIdentity]("datamigration.googleapis.com", "projects/{project}/locations/{location}/privateConnections/{privateconnection}")

// DatabaseMigrationPrivateConnectionIdentity is the identity of a GCP DatabaseMigrationPrivateConnection resource.
// +k8s:deepcopy-gen=false
type DatabaseMigrationPrivateConnectionIdentity struct {
	Project           string
	Location          string
	PrivateConnection string
}

func (i *DatabaseMigrationPrivateConnectionIdentity) String() string {
	return DatabaseMigrationPrivateConnectionIdentityFormat.ToString(*i)
}

func (i *DatabaseMigrationPrivateConnectionIdentity) FromExternal(ref string) error {
	parsed, match, err := DatabaseMigrationPrivateConnectionIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of DatabaseMigrationPrivateConnection external=%q was not known (use %s): %w", ref, DatabaseMigrationPrivateConnectionIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of DatabaseMigrationPrivateConnection external=%q was not known (use %s)", ref, DatabaseMigrationPrivateConnectionIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *DatabaseMigrationPrivateConnectionIdentity) Host() string {
	return DatabaseMigrationPrivateConnectionIdentityFormat.Host()
}

func getIdentityFromDatabaseMigrationPrivateConnectionSpec(ctx context.Context, reader client.Reader, obj *DatabaseMigrationPrivateConnection) (*DatabaseMigrationPrivateConnectionIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, err
	}

	location, err := refs.GetLocation(obj)
	if err != nil {
		return nil, err
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	identity := &DatabaseMigrationPrivateConnectionIdentity{
		Project:           projectID,
		Location:          location,
		PrivateConnection: resourceID,
	}
	return identity, nil
}

func (obj *DatabaseMigrationPrivateConnection) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromDatabaseMigrationPrivateConnectionSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &DatabaseMigrationPrivateConnectionIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change DatabaseMigrationPrivateConnection identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
