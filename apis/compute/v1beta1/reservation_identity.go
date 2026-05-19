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
	reference "github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/reference"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &ComputeReservationIdentity{}
	_ identity.Resource   = &ComputeReservation{}
)

var ComputeReservationIdentityFormat = gcpurls.Template[ComputeReservationIdentity]("compute.googleapis.com", "projects/{project}/zones/{zone}/reservations/{reservation}")

// +k8s:deepcopy-gen=false
type ComputeReservationIdentity struct {
	Project     string
	Zone        string
	Reservation string
}

func (i *ComputeReservationIdentity) String() string {
	return ComputeReservationIdentityFormat.ToString(*i)
}

func (i *ComputeReservationIdentity) FromExternal(ref string) error {
	focused := reference.FixStaleComputeExternalFormat(ref)

	parsed, match, err := ComputeReservationIdentityFormat.Parse(focused)
	if err != nil {
		return fmt.Errorf("format of ComputeReservation external=%q was not known (use %s): %w", ref, ComputeReservationIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of ComputeReservation external=%q was not known (use %s)", ref, ComputeReservationIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *ComputeReservationIdentity) Host() string {
	return ComputeReservationIdentityFormat.Host()
}

func getIdentityFromComputeReservationSpec(ctx context.Context, reader client.Reader, obj client.Object) (*ComputeReservationIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	u, ok := obj.(*ComputeReservation)
	if !ok {
		return nil, fmt.Errorf("object is not a ComputeReservation")
	}

	zone := common.ValueOf(u.Spec.Zone)
	if zone == "" {
		return nil, fmt.Errorf("cannot resolve zone")
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project")
	}

	identity := &ComputeReservationIdentity{
		Project:     projectID,
		Zone:        zone,
		Reservation: resourceID,
	}
	return identity, nil
}

func (obj *ComputeReservation) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromComputeReservationSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &ComputeReservationIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change ComputeReservation identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
