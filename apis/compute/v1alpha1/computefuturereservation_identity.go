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
	_ identity.IdentityV2 = &ComputeFutureReservationIdentity{}
	_ identity.Resource   = &ComputeFutureReservation{}
)

var ComputeFutureReservationIdentityFormat = gcpurls.Template[ComputeFutureReservationIdentity](
	"compute.googleapis.com",
	"projects/{project}/zones/{zone}/futureReservations/{futureReservation}",
)

// ComputeFutureReservationIdentity is the identity of a GCP ComputeFutureReservation resource.
// +k8s:deepcopy-gen=false
type ComputeFutureReservationIdentity struct {
	Project           string
	Zone              string
	FutureReservation string
}

func (i *ComputeFutureReservationIdentity) String() string {
	return ComputeFutureReservationIdentityFormat.ToString(*i)
}

func (i *ComputeFutureReservationIdentity) FromExternal(ref string) error {
	parsed, match, err := ComputeFutureReservationIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of ComputeFutureReservation external=%q was not known (use %s): %w", ref, ComputeFutureReservationIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of ComputeFutureReservation external=%q was not known (use %s)", ref, ComputeFutureReservationIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *ComputeFutureReservationIdentity) Host() string {
	return ComputeFutureReservationIdentityFormat.Host()
}

func getIdentityFromComputeFutureReservationSpec(ctx context.Context, reader client.Reader, obj *ComputeFutureReservation) (*ComputeFutureReservationIdentity, error) {
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

	identity := &ComputeFutureReservationIdentity{
		Project:           projectID,
		Zone:              location,
		FutureReservation: resourceID,
	}
	return identity, nil
}

func (obj *ComputeFutureReservation) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromComputeFutureReservationSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &ComputeFutureReservationIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change ComputeFutureReservation identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
