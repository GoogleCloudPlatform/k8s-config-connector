// Copyright 2024 Google LLC
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
	"reflect"
	"strings"

	gcp "cloud.google.com/go/bigquery/reservation/apiv1"
	pb "cloud.google.com/go/bigquery/reservation/apiv1/reservationpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigqueryreservation/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.BigQueryReservationReservationGVK, NewReservationModel)
}

func NewReservationModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelReservation{config: *config}, nil
}

var _ directbase.Model = &modelReservation{}

type modelReservation struct {
	config config.ControllerConfig
}

func (m *modelReservation) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building Reservation client: %w", err)
	}
	return gcpClient, err
}

func (m *modelReservation) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.BigQueryReservationReservation{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewReservationIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Get bigqueryreservation GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &ReservationAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelReservation) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type ReservationAdapter struct {
	id        *krm.ReservationIdentity
	gcpClient *gcp.Client
	desired   *krm.BigQueryReservationReservation
	actual    *pb.Reservation
}

var _ directbase.Adapter = &ReservationAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *ReservationAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting Reservation", "name", a.id.String())

	req := &pb.GetReservationRequest{Name: a.id.String()}
	reservationpb, err := a.gcpClient.GetReservation(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting Reservation %q: %w", a.id, err)
	}

	a.actual = reservationpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *ReservationAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating Reservation", "name", a.id.String())
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	desiredPb := BigQueryReservationReservationSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &pb.CreateReservationRequest{
		Parent:        a.id.Parent().String(),
		ReservationId: a.id.ID(),
		Reservation:   desiredPb,
	}
	created, err := a.gcpClient.CreateReservation(ctx, req)
	if err != nil {
		return fmt.Errorf("creating Reservation %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully created Reservation", "name", created.Name)

	status := &krm.BigQueryReservationReservationStatus{}
	status.ObservedState = BigQueryReservationReservationObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *ReservationAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating Reservation", "name", a.id.String())
	mapCtx := &direct.MapContext{}

	desiredSpec := &a.desired.DeepCopy().Spec
	desiredPb := BigQueryReservationReservationSpec_ToProto(mapCtx, desiredSpec)
	desiredPb.Name = a.id.String()
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	paths := []string{}

	if !reflect.DeepEqual(desiredPb.SlotCapacity, a.actual.SlotCapacity) {
		paths = append(paths, "slot_capacity")
	}
	if !reflect.DeepEqual(desiredPb.IgnoreIdleSlots, a.actual.IgnoreIdleSlots) {
		paths = append(paths, "ignore_idle_slots")
	}
	if !reflect.DeepEqual(desiredPb.Concurrency, a.actual.Concurrency) {
		paths = append(paths, "concurrency")
	}

	// Handle secondaryLocation field which can be modified by API during failover
	shouldUpdateSecondaryLocation := false

	if a.actual.SecondaryLocation == "" && desiredPb.SecondaryLocation != "" {
		// Case 1: Switching from non-failover to failover mode
		shouldUpdateSecondaryLocation = true
	} else if a.actual.SecondaryLocation != "" && desiredPb.SecondaryLocation == "" {
		// Case 2: Switching from failover to non-failover mode
		if a.actual.PrimaryLocation != desiredSpec.Location {
			return fmt.Errorf("updating Reservation %s: %s", a.id.String(), "Reservation is not available for change in secondary region")
		}
		shouldUpdateSecondaryLocation = true
	} // In all other cases, ignore the diff in this field

	if shouldUpdateSecondaryLocation {
		if desiredPb.Edition != pb.Edition_ENTERPRISE_PLUS {
			return fmt.Errorf("updating Reservation %s: %s", a.id.String(), "secondaryLocation is only available for ENTERPRISE_PLUS")
		}
		paths = append(paths, "secondary_location")
	}

	if desiredPb.Autoscale != nil && a.actual.Autoscale != nil && desiredPb.Autoscale.MaxSlots != a.actual.Autoscale.MaxSlots {
		paths = append(paths, "autoscale")
	} else if desiredPb.Autoscale != nil && a.actual.Autoscale == nil {
		paths = append(paths, "autoscale")
	}

	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id.String())
		status := &krm.BigQueryReservationReservationStatus{}
		status.ObservedState = BigQueryReservationReservationObservedState_FromProto(mapCtx, a.actual)
		if mapCtx.Err() != nil {
			return mapCtx.Err()
		}
		return updateOp.UpdateStatus(ctx, status, nil)
	}
	updateMask := &fieldmaskpb.FieldMask{
		Paths: paths,
	}

	req := &pb.UpdateReservationRequest{
		UpdateMask:  updateMask,
		Reservation: desiredPb,
	}
	updated, err := a.gcpClient.UpdateReservation(ctx, req)
	if err != nil {
		return fmt.Errorf("updating Reservation %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully updated Reservation", "name", updated.Name)

	status := &krm.BigQueryReservationReservationStatus{}
	status.ObservedState = BigQueryReservationReservationObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *ReservationAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.BigQueryReservationReservation{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(BigQueryReservationReservationSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Parent().ProjectID}
	obj.Spec.Location = a.id.Parent().Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	tokens := strings.Split(a.actual.Name, "/")
	u.SetName(tokens[5])
	u.SetGroupVersionKind(krm.BigQueryReservationReservationGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *ReservationAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting Reservation", "name", a.id)

	req := &pb.DeleteReservationRequest{Name: a.id.String()}
	err := a.gcpClient.DeleteReservation(ctx, req)
	if err != nil {
		return false, fmt.Errorf("deleting Reservation %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully deleted Reservation", "name", a.id.String())

	return true, nil
}
