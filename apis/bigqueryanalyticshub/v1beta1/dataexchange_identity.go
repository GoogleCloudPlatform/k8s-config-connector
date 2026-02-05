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
	_ identity.IdentityV2 = &BigQueryAnalyticsHubDataExchangeIdentity{}
	_ identity.Resource   = &BigQueryAnalyticsHubDataExchange{}
)

var BigQueryAnalyticsHubDataExchangeIdentityFormat = gcpurls.Template[BigQueryAnalyticsHubDataExchangeIdentity]("analyticshub.googleapis.com", "projects/{project}/locations/{location}/dataExchanges/{dataExchange}")

// +k8s:deepcopy-gen=false
type BigQueryAnalyticsHubDataExchangeIdentity struct {
	Project      string
	Location     string
	DataExchange string
}

func (i *BigQueryAnalyticsHubDataExchangeIdentity) String() string {
	return BigQueryAnalyticsHubDataExchangeIdentityFormat.ToString(*i)
}

func (i *BigQueryAnalyticsHubDataExchangeIdentity) FromExternal(ref string) error {
	parsed, match, err := BigQueryAnalyticsHubDataExchangeIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of BigQueryAnalyticsHubDataExchange external=%q was not known (use %s): %w", ref, BigQueryAnalyticsHubDataExchangeIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of BigQueryAnalyticsHubDataExchange external=%q was not known (use %s)", ref, BigQueryAnalyticsHubDataExchangeIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *BigQueryAnalyticsHubDataExchangeIdentity) Host() string {
	return BigQueryAnalyticsHubDataExchangeIdentityFormat.Host()
}

func ParseDataExchangeIdentity(external string) (*BigQueryAnalyticsHubDataExchangeIdentity, error) {
	id := &BigQueryAnalyticsHubDataExchangeIdentity{}
	if err := id.FromExternal(external); err != nil {
		return nil, err
	}
	return id, nil
}

func (i *BigQueryAnalyticsHubDataExchangeIdentity) ID() string {
	return i.DataExchange
}

func getIdentityFromBigQueryAnalyticsHubDataExchangeSpec(ctx context.Context, reader client.Reader, obj client.Object) (*BigQueryAnalyticsHubDataExchangeIdentity, error) {
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

	identity := &BigQueryAnalyticsHubDataExchangeIdentity{
		Project:      projectID,
		Location:     location,
		DataExchange: resourceID,
	}
	return identity, nil
}

func (obj *BigQueryAnalyticsHubDataExchange) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromBigQueryAnalyticsHubDataExchangeSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &BigQueryAnalyticsHubDataExchangeIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change BigQueryAnalyticsHubDataExchange identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
