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
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &BigQueryAnalyticsHubListingIdentity{}
	_ identity.Resource   = &BigQueryAnalyticsHubListing{}
)

var BigQueryAnalyticsHubListingIdentityFormat = gcpurls.Template[BigQueryAnalyticsHubListingIdentity]("analyticshub.googleapis.com", "projects/{project}/locations/{location}/dataExchanges/{dataExchange}/listings/{listing}")

// +k8s:deepcopy-gen=false
type BigQueryAnalyticsHubListingIdentity struct {
	Project      string
	Location     string
	DataExchange string
	Listing      string
}

func (i *BigQueryAnalyticsHubListingIdentity) String() string {
	return BigQueryAnalyticsHubListingIdentityFormat.ToString(*i)
}

func (i *BigQueryAnalyticsHubListingIdentity) FromExternal(ref string) error {
	parsed, match, err := BigQueryAnalyticsHubListingIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of BigQueryAnalyticsHubListing external=%q was not known (use %s): %w", ref, BigQueryAnalyticsHubListingIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of BigQueryAnalyticsHubListing external=%q was not known (use %s)", ref, BigQueryAnalyticsHubListingIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *BigQueryAnalyticsHubListingIdentity) Host() string {
	return BigQueryAnalyticsHubListingIdentityFormat.Host()
}

func getIdentityFromBigQueryAnalyticsHubListingSpec(ctx context.Context, reader client.Reader, obj client.Object) (*BigQueryAnalyticsHubListingIdentity, error) {
	listing := &BigQueryAnalyticsHubListing{}
	switch t := obj.(type) {
	case *BigQueryAnalyticsHubListing:
		listing = t
	case *unstructured.Unstructured:
		if err := runtime.DefaultUnstructuredConverter.FromUnstructured(t.Object, listing); err != nil {
			return nil, fmt.Errorf("failed to convert unstructured to BigQueryAnalyticsHubListing: %w", err)
		}
	default:
		return nil, fmt.Errorf("expected *BigQueryAnalyticsHubListing or *unstructured.Unstructured, got %T", obj)
	}

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

	dataExchangeRef := listing.Spec.DataExchangeRef.DeepCopy()
	if err := dataExchangeRef.Normalize(ctx, reader, listing.Namespace); err != nil {
		return nil, fmt.Errorf("cannot resolve dataExchangeRef: %w", err)
	}
	dataExchangeExternal := dataExchangeRef.External
	dataExchangeID := &BigQueryAnalyticsHubDataExchangeIdentity{}
	if err := dataExchangeID.FromExternal(dataExchangeExternal); err != nil {
		return nil, fmt.Errorf("cannot parse dataExchangeRef: %w", err)
	}

	if dataExchangeID.Project != projectID {
		return nil, fmt.Errorf("dataExchangeRef.project must match spec.projectRef")
	}
	if dataExchangeID.Location != location {
		return nil, fmt.Errorf("dataExchangeRef.location must match spec.location")
	}

	identity := &BigQueryAnalyticsHubListingIdentity{
		Project:      projectID,
		Location:     location,
		DataExchange: dataExchangeID.DataExchange,
		Listing:      resourceID,
	}
	return identity, nil
}

func (obj *BigQueryAnalyticsHubListing) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromBigQueryAnalyticsHubListingSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &BigQueryAnalyticsHubListingIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change BigQueryAnalyticsHubListing identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
