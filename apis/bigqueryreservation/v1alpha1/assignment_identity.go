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
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// AssignmentIdentity defines the resource identity to BigQueryReservationAssignment,
// which is to be used in the controller.
type AssignmentIdentity struct {
	// The reservation to which the assignment is currently attached
	parent *BQReservation
	// The assignment resourceID
	id string
}

func (i *AssignmentIdentity) String() string {
	return i.parent.String() + "/assignments/" + i.id
}

func (i *AssignmentIdentity) ID() string {
	return i.id
}

func (i *AssignmentIdentity) Parent() *BQReservation {
	return i.parent
}

func (i *AssignmentIdentity) GetProjetID() string {
	return i.parent.ProjectID
}

type BQReservation struct {
	ProjectID       string
	Location        string
	ReservationName string
}

func (p *BQReservation) String() string {
	return "projects/" + p.ProjectID + "/locations/" + p.Location + "/reservations/" + p.ReservationName
}

// New builds a AssignmentIdentity from the Config Connector BigQueryReservationAssignment object.
func NewAssignmentIdentity(ctx context.Context, reader client.Reader, obj *BigQueryReservationAssignment) (*AssignmentIdentity, error) {

	// Get the reservation to move the assignment to.
	// Uses the "spec.ReservationRef"
	var name string
	var err error
	if obj.Spec.ReservationRef.External != "" {
		name = obj.Spec.ReservationRef.External
	} else {
		name, err = obj.Spec.ReservationRef.NormalizedExternal(ctx, reader, obj.GetNamespace())
		if err != nil {
			return nil, err
		}
	}

	reservationParent, reservationID, err := ParseReservationExternal(name)
	if err != nil {
		return nil, err
	}

	parent := &BQReservation{
		ProjectID:       reservationParent.ProjectID,
		Location:        reservationParent.Location,
		ReservationName: reservationID,
	}

	// Get desired assignment resourceID
	// Only works for resource acquisition
	resourceID := common.ValueOf(obj.Spec.ResourceID)

	// Use approved External
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// The assignment is already existing
		actualParent, actualResourceID, err := ParseAssignmentExternal(externalRef)
		if err != nil {
			return nil, err
		}
		// Need to reset the parent to the actual parent
		//  when moving the assignment from one reservation to another.
		if actualParent.ProjectID != parent.ProjectID {
			parent.ProjectID = actualParent.ProjectID
		}
		if actualParent.Location != parent.Location {
			parent.Location = actualParent.Location
		}
		if actualParent.ReservationName != parent.ReservationName {
			parent.ReservationName = actualParent.ReservationName
		}
		// For BigQueryReservationAssignment, the GCP resourceID is output only.
		if resourceID == "" {
			resourceID = actualResourceID
		} else if resourceID != actualResourceID {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
				resourceID, actualResourceID)
		}
	}

	return &AssignmentIdentity{
		parent: parent,
		id:     resourceID,
	}, nil
}

func ParseAssignmentExternal(external string) (parent *BQReservation, resourceID string, err error) {
	tokens := strings.Split(external, "/")
	if len(tokens) != 8 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "reservations" || tokens[6] != "assignments" {
		return nil, "", fmt.Errorf("format of BigqueryReservation external=%q was not known (use projects/{{projectID}}/locations/{{location}}/reservations/{{reservationName}}/assignments/{{assignmentID}})", external)
	}
	parent = &BQReservation{
		ProjectID:       tokens[1],
		Location:        tokens[3],
		ReservationName: tokens[5],
	}
	resourceID = tokens[7]
	return parent, resourceID, nil
}
