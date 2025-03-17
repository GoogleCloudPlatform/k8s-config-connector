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
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// AssignmentIdentity defines the resource reference to BigQueryReservationAssignment,
// which "External" field holds the GCP identifier for the KRM object.
type AssignmentIdentity struct {
	parent    *AssignmentParent
	projectID string
	id        string
}

func (i *AssignmentIdentity) String() string {
	return i.parent.String() + "/assignments/" + i.id
}

func (i *AssignmentIdentity) ID() string {
	return i.id
}

func (i *AssignmentIdentity) Parent() *AssignmentParent {
	return i.parent
}

func (i *AssignmentIdentity) GetProjetID() string {
	return i.projectID
}

type AssignmentParent struct {
	ProjectID       string
	Location        string
	ReservationName string
}

func (p *AssignmentParent) String() string {
	return "projects/" + p.ProjectID + "/locations/" + p.Location + "/reservations/" + p.ReservationName
}

// New builds a AssignmentIdentity from the Config Connector BigQueryReservationAssignment object.
func NewAssignmentIdentity(ctx context.Context, reader client.Reader, obj *BigQueryReservationAssignment) (*AssignmentIdentity, error) {

	// Get Parent
	projectRef, err := refsv1beta1.ResolveProject(ctx, reader, obj.GetNamespace(), obj.Spec.ProjectRef)
	if err != nil {
		return nil, err
	}
	projectID := projectRef.ProjectID
	if projectID == "" {
		return nil, fmt.Errorf("cannot resolve project")
	}
	location := *obj.Spec.Location
	reservationName := *obj.Spec.ReservationName

	// Get desired ID
	// Only works for resource aquisition
	resourceID := common.ValueOf(obj.Spec.ResourceID)

	// Use approved External
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		actualParent, actualResourceID, err := ParseAssignmentExternal(externalRef)
		if err != nil {
			return nil, err
		}
		// Moving an Assignment from one reservation to another reservation
		//  can override these values
		if actualParent.ProjectID != projectID {
			projectID = actualParent.ProjectID
		}
		if actualParent.Location != location {
			location = actualParent.Location
		}
		if actualParent.ReservationName != reservationName {
			reservationName = actualParent.ReservationName
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
		parent: &AssignmentParent{
			ProjectID:       projectID,
			Location:        location,
			ReservationName: reservationName,
		},
		projectID: projectID,
		id:        resourceID,
	}, nil
}

func ParseAssignmentExternal(external string) (parent *AssignmentParent, resourceID string, err error) {
	tokens := strings.Split(external, "/")
	if len(tokens) != 8 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "reservations" || tokens[6] != "assignments" {
		return nil, "", fmt.Errorf("format of BigqueryReservation external=%q was not known (use projects/{{projectID}}/locations/{{location}}/reservations/{{reservationName}}/assignments/{{assignmentID}})", external)
	}
	parent = &AssignmentParent{
		ProjectID:       tokens[1],
		Location:        tokens[3],
		ReservationName: tokens[5],
	}
	resourceID = tokens[7]
	return parent, resourceID, nil
}
