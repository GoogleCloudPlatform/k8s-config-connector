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

// +tool:controller
// proto.service: google.cloud.compute.v1.Reservations
// proto.message: google.cloud.compute.v1.Reservation
// crd.type: ComputeReservation
// crd.version: v1beta1

package compute

import (
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"

	compute "cloud.google.com/go/compute/apiv1"
	computepb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.ComputeReservationGVK, NewReservationModel)
}

func NewReservationModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &reservationModel{config: config}, nil
}

var _ directbase.Model = &reservationModel{}

type reservationModel struct {
	config *config.ControllerConfig
}

func (m *reservationModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ComputeReservation{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}

	gcpClient, err := newGCPClient(m.config)
	if err != nil {
		return nil, fmt.Errorf("building gcp client: %w", err)
	}
	reservationsClient, err := gcpClient.newReservationsClient(ctx)
	if err != nil {
		return nil, err
	}

	return &ReservationAdapter{
		gcpClient: reservationsClient,
		id:        id.(*krm.ComputeReservationIdentity),
		desired:   obj,
		reader:    reader,
	}, nil
}

func (m *reservationModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type ReservationAdapter struct {
	gcpClient *compute.ReservationsClient
	id        *krm.ComputeReservationIdentity
	desired   *krm.ComputeReservation
	actual    *computepb.Reservation
	reader    client.Reader
}

var _ directbase.Adapter = &ReservationAdapter{}

// Find retrieves the GCP resource.
func (a *ReservationAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting Reservation", "name", a.id)

	req := &computepb.GetReservationRequest{
		Project:     a.id.Project,
		Zone:        a.id.Zone,
		Reservation: a.id.Reservation,
	}
	actual, err := a.gcpClient.Get(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting Reservation %q: %w", a.id, err)
	}

	a.actual = actual
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *ReservationAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating Reservation", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := ComputeReservationSpec_v1beta1_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	resource.Name = direct.LazyPtr(a.id.Reservation)
	// match realGCP log: "zone": "projects/${projectId}/global/zones/us-central1-a"
	resource.Zone = direct.LazyPtr(fmt.Sprintf("projects/%s/global/zones/%s", a.id.Project, a.id.Zone))

	req := &computepb.InsertReservationRequest{
		Project:             a.id.Project,
		Zone:                a.id.Zone,
		ReservationResource: resource,
	}
	op, err := a.gcpClient.Insert(ctx, req)
	if err != nil {
		return fmt.Errorf("creating Reservation %s: %w", a.id, err)
	}

	err = op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("compute Reservation %s waiting creation: %w", a.id.String(), err)
	}
	log.Info("successfully created compute Reservation in gcp", "name", a.id)

	// Get the created resource
	created, err := a.get(ctx)
	if err != nil {
		return fmt.Errorf("getting Reservation %s: %w", a.id, err)
	}

	status := ComputeReservationStatus_v1beta1_FromProto(mapCtx, created)
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *ReservationAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating Reservation", "name", a.id)

	mapCtx := &direct.MapContext{}
	desired := a.desired.DeepCopy()
	desiredPb := ComputeReservationSpec_v1beta1_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	desiredPb.Name = direct.LazyPtr(a.id.Reservation)

	// Handle output-only fields from GCP
	a.assignGCPDefaults(desiredPb, a.actual)

	paths, report, err := common.CompareProtoMessageStructuredDiff(desiredPb, a.actual, common.BasicDiff)
	if err != nil {
		return err
	}

	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		if a.desired.Status.ExternalRef == nil {
			status := ComputeReservationStatus_v1beta1_FromProto(mapCtx, a.actual)
			status.ExternalRef = direct.LazyPtr(a.id.String())
			return updateOp.UpdateStatus(ctx, status, nil)
		}
		return nil
	}

	for path := range paths {
		if path != "specific_reservation.count" {
			return fmt.Errorf("field %q is immutable", path)
		}
	}

	report.Object = updateOp.GetUnstructured()
	structuredreporting.ReportDiff(ctx, report)

	updateOp.RecordUpdatingEvent()

	req := &computepb.ResizeReservationRequest{
		Project:     a.id.Project,
		Zone:        a.id.Zone,
		Reservation: a.id.Reservation,
		ReservationsResizeRequestResource: &computepb.ReservationsResizeRequest{
			SpecificSkuCount: direct.PtrTo(int64(desiredPb.GetSpecificReservation().GetCount())),
		},
	}
	op, err := a.gcpClient.Resize(ctx, req)
	if err != nil {
		return fmt.Errorf("resizing compute Reservation %s: %w", a.id.String(), err)
	}

	err = op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for resize of compute Reservation %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully resized compute Reservation", "name", a.id)

	// Get the updated resource
	updated, err := a.get(ctx)
	if err != nil {
		return fmt.Errorf("getting ComputeReservation %s: %w", a.id, err)
	}

	status := ComputeReservationStatus_v1beta1_FromProto(mapCtx, updated)
	if a.desired.Status.ExternalRef == nil {
		status.ExternalRef = direct.LazyPtr(a.id.String())
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *ReservationAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ComputeReservation{}
	mapCtx := &direct.MapContext{}
	obj.Spec = *ComputeReservationSpec_v1beta1_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	// Note: project is not explicitly in spec but inferred from namespace, we set it if needed
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.Reservation)
	u.SetGroupVersionKind(krm.ComputeReservationGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *ReservationAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting Reservation", "name", a.id)

	req := &computepb.DeleteReservationRequest{
		Project:     a.id.Project,
		Zone:        a.id.Zone,
		Reservation: a.id.Reservation,
	}
	op, err := a.gcpClient.Delete(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent compute Reservation, assuming it was already deleted", "name", a.id.String())
			return true, nil
		}
		return false, fmt.Errorf("deleting compute Reservation %s: %w", a.id.String(), err)
	}

	if !op.Done() {
		err = op.Wait(ctx)
		if err != nil {
			return false, fmt.Errorf("waiting for deletion of compute Reservation %s: %w", a.id.String(), err)
		}
	}
	log.Info("successfully deleted compute Reservation", "name", a.id)

	return true, nil
}

func (a *ReservationAdapter) get(ctx context.Context) (*computepb.Reservation, error) {
	getReq := &computepb.GetReservationRequest{
		Project:     a.id.Project,
		Zone:        a.id.Zone,
		Reservation: a.id.Reservation,
	}
	resource, err := a.gcpClient.Get(ctx, getReq)
	if err != nil {
		return nil, fmt.Errorf("getting ComputeReservation %s: %w", a.id, err)
	}
	return resource, nil
}

func (a *ReservationAdapter) assignGCPDefaults(desired *computepb.Reservation, actual *computepb.Reservation) {
	// Output-only fields
	desired.SelfLink = actual.SelfLink
	desired.CreationTimestamp = actual.CreationTimestamp
	desired.Status = actual.Status
	desired.Kind = actual.Kind
	desired.Id = actual.Id
	desired.Commitment = actual.Commitment
	desired.ResourceStatus = actual.ResourceStatus

	// Convert Zone to match GCP API, the immutability is handled in reservation_identity.go
	desired.Zone = actual.Zone

	// GCP default fields if not specified by user
	if desired.ReservationSharingPolicy == nil {
		desired.ReservationSharingPolicy = actual.ReservationSharingPolicy
	}
	if desired.ShareSettings == nil {
		desired.ShareSettings = actual.ShareSettings
	}
	if desired.SpecificReservation != nil && desired.SpecificReservation.AssuredCount == nil {
		desired.SpecificReservation.AssuredCount = actual.SpecificReservation.AssuredCount
	}
}
