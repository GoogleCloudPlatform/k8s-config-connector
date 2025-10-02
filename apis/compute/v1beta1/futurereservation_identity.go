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

package v1beta1

import (
	"context"
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// FutureReservationIdentity defines the resource reference to ComputeFutureReservation, which "External" field
// holds the GCP identifier for the KRM object.
type FutureReservationIdentity struct {
	parent *FutureReservationParent
	id     string
}

func (i *FutureReservationIdentity) String() string {
	return i.parent.String() + "/futureReservations/" + i.id
}

func (i *FutureReservationIdentity) ID() string {
	return i.id
}

func (i *FutureReservationIdentity) Parent() *FutureReservationParent {
	return i.parent
}

type FutureReservationParent struct {
	ProjectID string
	Zone      string
}

func (p *FutureReservationParent) String() string {
	return "projects/" + p.ProjectID + "/zones/" + p.Zone
}

// New builds a FutureReservationIdentity from the Config Connector FutureReservation object.
func NewFutureReservationIdentity(ctx context.Context, reader client.Reader, obj *ComputeFutureReservation) (*FutureReservationIdentity, error) {

	// Get Parent
	projectRef, err := refsv1beta1.ResolveProject(ctx, reader, obj.GetNamespace(), obj.Spec.ProjectRef)
	if err != nil {
		return nil, err
	}
	projectID := projectRef.ProjectID
	if projectID == "" {
		return nil, fmt.Errorf("cannot resolve project")
	}
	zone := obj.Spec.Zone

	// Get desired ID
	resourceID := common.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	// Use approved External
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		actualParent, actualResourceID, err := ParseFutureReservationExternal(externalRef)
		if err != nil {
			return nil, err
		}
		if actualParent.ProjectID != projectID {
			return nil, fmt.Errorf("spec.projectRef changed, expect %s, got %s", actualParent.ProjectID, projectID)
		}
		if actualParent.Zone != zone {
			return nil, fmt.Errorf("spec.zone changed, expect %s, got %s", actualParent.Zone, zone)
		}
		if actualResourceID != resourceID {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
				resourceID, actualResourceID)
		}
	}
	return &FutureReservationIdentity{
		parent: &FutureReservationParent{
			ProjectID: projectID,
			Zone:      zone,
		},
		id: resourceID,
	}, nil
}

func ParseFutureReservationExternal(external string) (parent *FutureReservationParent, resourceID string, err error) {
	tokens := strings.Split(external, "/")
	if len(tokens) != 6 || tokens[0] != "projects" || tokens[2] != "zones" || tokens[4] != "futureReservations" {
		return nil, "", fmt.Errorf("format of ComputeFutureReservation external=%q was not known (use projects/{{projectID}}/zones/{{zone}}/futureReservations/{{futurereservationID}})", external)
	}
	parent = &FutureReservationParent{
		ProjectID: tokens[1],
		Zone:      tokens[3],
	}
	resourceID = tokens[5]
	return parent, resourceID, nil
}
