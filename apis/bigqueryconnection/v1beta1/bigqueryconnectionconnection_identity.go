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

package v1beta1

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
	_ identity.ServerGeneratedIdentity = &BigQueryConnectionConnectionIdentity{}
	_ identity.Resource                = &BigQueryConnectionConnection{}
)

var BigQueryConnectionConnectionIdentityFormat = gcpurls.Template[BigQueryConnectionConnectionIdentity]("bigqueryconnection.googleapis.com", "projects/{project}/locations/{location}/connections/{connection}")

// BigQueryConnectionConnectionIdentity is the identity of a GCP BigQueryConnectionConnection resource.
// +k8s:deepcopy-gen=false
type BigQueryConnectionConnectionIdentity struct {
	Project    string
	Location   string
	Connection string
}

func (i *BigQueryConnectionConnectionIdentity) HasIdentitySpecified() bool {
	return i.Connection != ""
}

func (i *BigQueryConnectionConnectionIdentity) String() string {
	return BigQueryConnectionConnectionIdentityFormat.ToString(*i)
}

func (i *BigQueryConnectionConnectionIdentity) FromExternal(ref string) error {
	parsed, match, err := BigQueryConnectionConnectionIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of BigQueryConnectionConnection external=%q was not known (use %s): %w", ref, BigQueryConnectionConnectionIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of BigQueryConnectionConnection external=%q was not known (use %s)", ref, BigQueryConnectionConnectionIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *BigQueryConnectionConnectionIdentity) Host() string {
	return BigQueryConnectionConnectionIdentityFormat.Host()
}

func getIdentityFromBigQueryConnectionConnectionSpec(ctx context.Context, reader client.Reader, obj *BigQueryConnectionConnection) (*BigQueryConnectionConnectionIdentity, error) {
	// For BigQueryConnectionConnection, resourceID is optional and can be empty.
	// We retrieve it directly from Spec.ResourceID to avoid falling back to GetName().
	resourceID := common.ValueOf(obj.Spec.ResourceID)

	location := obj.Spec.Location

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project: %w", err)
	}

	identity := &BigQueryConnectionConnectionIdentity{
		Project:    projectID,
		Location:   location,
		Connection: resourceID,
	}
	return identity, nil
}

func (obj *BigQueryConnectionConnection) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromBigQueryConnectionConnectionSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &BigQueryConnectionConnectionIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if specIdentity.Connection == "" {
			specIdentity.Connection = statusIdentity.Connection
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change BigQueryConnectionConnection identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
