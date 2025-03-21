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

package bigqueryreservation

import (
	"context"
	"fmt"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigqueryreservation/v1alpha1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	gcp "cloud.google.com/go/bigquery/reservation/apiv1"

	pb "cloud.google.com/go/bigquery/reservation/apiv1/reservationpb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.BigQueryReservationAssignmentGVK, NewAssignmentModel)
}

func NewAssignmentModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelAssignment{config: *config}, nil
}

var _ directbase.Model = &modelAssignment{}

type modelAssignment struct {
	config config.ControllerConfig
}

func (m *modelAssignment) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building Assignment client: %w", err)
	}
	return gcpClient, err
}

func (m *modelAssignment) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.BigQueryReservationAssignment{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewAssignmentIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Get bigqueryreservation GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	// Normalize Assignee
	if obj.Spec.Assignee.ProjectRef != nil {
		project, err := refsv1beta1.ResolveProject(ctx, reader, u.GetNamespace(), obj.Spec.Assignee.ProjectRef)
		if err != nil {
			return nil, err
		}
		obj.Spec.Assignee.ProjectRef.External = project.ProjectID
	}
	if obj.Spec.Assignee.FolderRef != nil {
		folder, err := refsv1beta1.ResolveFolder(ctx, reader, u, obj.Spec.Assignee.FolderRef)
		if err != nil {
			return nil, err
		}
		obj.Spec.Assignee.FolderRef.External = folder.FolderID
	}
	if obj.Spec.Assignee.OrganizationRef != nil {
		org, err := refsv1beta1.ResolveOrganization(ctx, reader, u, obj.Spec.Assignee.OrganizationRef)
		if err != nil {
			return nil, err
		}
		obj.Spec.Assignee.OrganizationRef.External = org.OrganizationID
	}

	return &AssignmentAdapter{
		id:            id,
		gcpClient:     gcpClient,
		desired:       obj,
		destinationId: obj.Spec.ReservationRef.External,
	}, nil
}

func (m *modelAssignment) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type AssignmentAdapter struct {
	id        *krm.AssignmentIdentity
	gcpClient *gcp.Client
	desired   *krm.BigQueryReservationAssignment
	actual    *pb.Assignment
	// The reservation to move the assignment to
	destinationId string
}

var _ directbase.Adapter = &AssignmentAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *AssignmentAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting BigqueryReservationAssignment", "name", a.id.String())

	req := &pb.ListAssignmentsRequest{Parent: a.id.Parent().String()}
	assignmentIterator := a.gcpClient.ListAssignments(ctx, req)
	if assignmentIterator == nil {
		fmt.Printf("not found Assignment %q\n", a.id.String())
		return false, nil
	}

	// There is no more items when error is iterator.Done.
	for assignmentpb, err := assignmentIterator.Next(); err == nil; {
		if assignmentpb.Name == a.id.String() {
			a.actual = assignmentpb
			return true, nil
		}
		assignmentpb, err = assignmentIterator.Next()
	}

	return false, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *AssignmentAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating Assignment", "name", a.id.String())
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	assignment := BigqueryReservationAssignmentSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &pb.CreateAssignmentRequest{
		Parent:       a.id.Parent().String(),
		AssignmentId: a.id.AssignmentID(),
		Assignment:   assignment,
	}
	created, err := a.gcpClient.CreateAssignment(ctx, req)
	if err != nil {
		return fmt.Errorf("creating Assignment %s: %w", a.id.String(), err)
	}

	log.V(2).Info("successfully created Assignment", "name", a.id.String())

	status := &krm.BigQueryReservationAssignmentStatus{}
	status.ObservedState = BigqueryReservationAssignmentObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	// update the externalRef in the KRM resoruce
	status.ExternalRef = direct.LazyPtr(created.GetName())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *AssignmentAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating or moving the Assignment", "name", a.id.String())
	mapCtx := &direct.MapContext{}
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	desiredSpec := &a.desired.DeepCopy().Spec

	// Get the reservation to move the assignment from
	currentReservation, _, err := krm.ParseAssignmentExternal(a.actual.GetName())
	if err != nil {
		return err
	}

	var updated *pb.Assignment
	status := &krm.BigQueryReservationAssignmentStatus{}
	// Case1: Move the assignment to another reservation
	if currentReservation.String() != a.destinationId {
		log.V(2).Info("moving assignment to another reservation", "current", a.id.String())
		req := &pb.MoveAssignmentRequest{
			Name:          a.actual.GetName(),
			DestinationId: a.destinationId,
			AssignmentId:  a.id.AssignmentID(),
		}
		// if user wants to retain the assignmentID
		if desiredSpec.ResourceID != nil {
			req.AssignmentId = direct.ValueOf(desiredSpec.ResourceID)
		}
		updated, err = a.gcpClient.MoveAssignment(ctx, req)
		if err != nil {
			return fmt.Errorf("moving Assignment %s: %w", a.id.String(), err)
		}

		// Rebuild the externalRef
		status.ExternalRef = direct.LazyPtr(updated.GetName())
	}

	/* 	if len(paths) == 0 {
	   		log.V(2).Info("no field needs update", "name", a.id.String())
	   		status := &krm.BigQueryReservationAssignmentStatus{}
	   		status.ObservedState = BigqueryReservationAssignmentObservedState_FromProto(mapCtx, a.actual)
	   		if mapCtx.Err() != nil {
	   			return mapCtx.Err()
	   		}
	   		return updateOp.UpdateStatus(ctx, status, nil)
	   	}
	   	updateMask := &fieldmaskpb.FieldMask{
	   		Paths: paths}

	   	// TODO(contributor): Complete the gcp "UPDATE" or "PATCH" request.
	   	req := &pb.UpdateAssignmentRequest{
	   		UpdateMask: updateMask,
	   		Assignment: desiredPb,
	   	}
	   	updated, err := a.gcpClient.UpdateAssignment(ctx, req)
	   	if err != nil {
	   		return fmt.Errorf("updating Assignment %s: %w", a.id.String(), err)
	   	}
	   	log.V(2).Info("successfully updated Assignment", "name", a.id.String()) */

	status.ObservedState = BigqueryReservationAssignmentObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *AssignmentAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.BigQueryReservationAssignment{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(BigqueryReservationAssignmentSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ReservationRef = &krm.ReservationRef{External: a.destinationId}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.actual.Name)
	u.SetGroupVersionKind(krm.BigQueryReservationAssignmentGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *AssignmentAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting Assignment", "name", a.id.String())

	req := &pb.DeleteAssignmentRequest{Name: a.id.String()}
	err := a.gcpClient.DeleteAssignment(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent Assignment, assuming it was already deleted", "name", a.id.String())
			return true, nil
		}
		return false, fmt.Errorf("deleting Assignment %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully deleted Assignment", "name", a.id.String())

	return true, nil
}
