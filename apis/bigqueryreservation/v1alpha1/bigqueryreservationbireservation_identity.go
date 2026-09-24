// Copyright 2025 Google LLC
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
	_ identity.IdentityV2 = &BigQueryReservationBiReservationIdentity{}
	_ identity.Resource   = &BigQueryReservationBiReservation{}
)

var BigQueryReservationBiReservationIdentityFormat = gcpurls.Template[BigQueryReservationBiReservationIdentity]("bigqueryreservation.googleapis.com", "projects/{project}/locations/{location}/biReservation")

// BigQueryReservationBiReservationIdentity is the identity of a GCP BigQueryReservationBiReservation resource.
// +k8s:deepcopy-gen=false
type BigQueryReservationBiReservationIdentity struct {
	Project  string
	Location string
}

func (i *BigQueryReservationBiReservationIdentity) String() string {
	return BigQueryReservationBiReservationIdentityFormat.ToString(*i)
}

func (i *BigQueryReservationBiReservationIdentity) FromExternal(ref string) error {
	parsed, match, err := BigQueryReservationBiReservationIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of BigQueryReservationBiReservation external=%q was not known (use %s): %w", ref, BigQueryReservationBiReservationIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of BigQueryReservationBiReservation external=%q was not known (use %s)", ref, BigQueryReservationBiReservationIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *BigQueryReservationBiReservationIdentity) Host() string {
	return BigQueryReservationBiReservationIdentityFormat.Host()
}

func getIdentityFromBigQueryReservationBiReservationSpec(ctx context.Context, reader client.Reader, obj *BigQueryReservationBiReservation) (*BigQueryReservationBiReservationIdentity, error) {
	location, err := refs.GetLocation(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve location: %w", err)
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project: %w", err)
	}

	identity := &BigQueryReservationBiReservationIdentity{
		Project:  projectID,
		Location: location,
	}
	return identity, nil
}

func (obj *BigQueryReservationBiReservation) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromBigQueryReservationBiReservationSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &BigQueryReservationBiReservationIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change BigQueryReservationBiReservation identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
