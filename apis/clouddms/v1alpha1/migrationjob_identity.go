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
	_ identity.IdentityV2 = &CloudDMSMigrationJobIdentity{}
	_ identity.Resource   = &CloudDMSMigrationJob{}
)

var CloudDMSMigrationJobIdentityFormat = gcpurls.Template[CloudDMSMigrationJobIdentity]("datamigration.googleapis.com", "projects/{project}/locations/{location}/migrationJobs/{migrationJob}")

// +k8s:deepcopy-gen=false
type CloudDMSMigrationJobIdentity struct {
	Project      string
	Location     string
	MigrationJob string
}

func (i *CloudDMSMigrationJobIdentity) String() string {
	return CloudDMSMigrationJobIdentityFormat.ToString(*i)
}

func (i *CloudDMSMigrationJobIdentity) FromExternal(ref string) error {
	parsed, match, err := CloudDMSMigrationJobIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of CloudDMSMigrationJob external=%q was not known (use %s): %w", ref, CloudDMSMigrationJobIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of CloudDMSMigrationJob external=%q was not known (use %s)", ref, CloudDMSMigrationJobIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *CloudDMSMigrationJobIdentity) Host() string {
	return CloudDMSMigrationJobIdentityFormat.Host()
}

func getIdentityFromCloudDMSMigrationJobSpec(ctx context.Context, reader client.Reader, obj client.Object) (*CloudDMSMigrationJobIdentity, error) {
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

	identity := &CloudDMSMigrationJobIdentity{
		Project:      projectID,
		Location:     location,
		MigrationJob: resourceID,
	}
	return identity, nil
}

func (obj *CloudDMSMigrationJob) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromCloudDMSMigrationJobSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &CloudDMSMigrationJobIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change CloudDMSMigrationJob identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
