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
	_ identity.IdentityV2 = &BigQueryMigrationWorkflowIdentity{}
	_ identity.Resource   = &BigQueryMigrationWorkflow{}
)

var BigQueryMigrationWorkflowIdentityFormat = gcpurls.Template[BigQueryMigrationWorkflowIdentity]("bigquerymigration.googleapis.com", "projects/{project}/locations/{location}/workflows/{workflow}")

// +k8s:deepcopy-gen=false
type BigQueryMigrationWorkflowIdentity struct {
	Project  string
	Location string
	Workflow string
}

func (i *BigQueryMigrationWorkflowIdentity) String() string {
	return BigQueryMigrationWorkflowIdentityFormat.ToString(*i)
}

func (i *BigQueryMigrationWorkflowIdentity) FromExternal(ref string) error {
	parsed, match, err := BigQueryMigrationWorkflowIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of BigQueryMigrationWorkflow external=%q was not known (use %s): %w", ref, BigQueryMigrationWorkflowIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of BigQueryMigrationWorkflow external=%q was not known (use %s)", ref, BigQueryMigrationWorkflowIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *BigQueryMigrationWorkflowIdentity) Host() string {
	return BigQueryMigrationWorkflowIdentityFormat.Host()
}

func getIdentityFromBigQueryMigrationWorkflowSpec(ctx context.Context, reader client.Reader, obj client.Object) (*BigQueryMigrationWorkflowIdentity, error) {
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

	identity := &BigQueryMigrationWorkflowIdentity{
		Project:  projectID,
		Location: location,
		Workflow: resourceID,
	}
	return identity, nil
}

func (obj *BigQueryMigrationWorkflow) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromBigQueryMigrationWorkflowSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		statusIdentity := &BigQueryMigrationWorkflowIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change BigQueryMigrationWorkflow identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
