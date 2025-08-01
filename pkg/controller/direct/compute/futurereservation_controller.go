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

package compute

import (
	"context"
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	computev1beta "cloud.google.com/go/compute/apiv1beta"
	computepb "cloud.google.com/go/compute/apiv1beta/computepb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.FutureReservationGVK, NewFutureReservationModel)
}

func NewFutureReservationModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &futureReservationModel{config: *config}, nil
}

var _ directbase.Model = &futureReservationModel{}

type futureReservationModel struct {
	config config.ControllerConfig
}

func (m *futureReservationModel) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.FutureReservation{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewFutureReservationIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Get compute GCP client
	gcpClient, err := newGCPClient(&m.config)
	if err != nil {
		return nil, fmt.Errorf("building FutureReservation client: %w", err)
	}
	futureReservationClient, err := gcpClient.newFutureReservationsClient(ctx)
	if err != nil {
		return nil, err
	}

	return &FutureReservationAdapter{
		id:        id,
		gcpClient: futureReservationClient,
		desired:   obj,
	}, nil
}

func (m *futureReservationModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type FutureReservationAdapter struct {
	id        *krm.FutureReservationIdentity
	gcpClient *computev1beta.FutureReservationsClient
	desired   *krm.FutureReservation
	actual    *computepb.FutureReservation
}

var _ directbase.Adapter = &FutureReservationAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *FutureReservationAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting FutureReservation", "name", a.id)

	req := &computepb.GetFutureReservationRequest{
		Project:           a.id.Parent().ProjectID,
		Zone:              a.id.Parent().Location,
		FutureReservation: a.id.ID(),
	}
	futurereservationpb, err := a.gcpClient.Get(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting FutureReservation %q: %w", a.id, err)
	}

	a.actual = futurereservationpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *FutureReservationAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating FutureReservation", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := FutureReservationSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	resource.Name = direct.LazyPtr(a.id.ID())

	req := &computepb.InsertFutureReservationRequest{
		Project:                   a.id.Parent().ProjectID,
		Zone:                      a.id.Parent().Location,
		FutureReservationResource: resource,
	}
	op, err := a.gcpClient.Insert(ctx, req)
	if err != nil {
		return fmt.Errorf("creating FutureReservation %s: %w", a.id, err)
	}
	err = op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("FutureReservation %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created FutureReservation", "name", a.id)

	// Get the created resource
	created, err := a.get(ctx)
	if err != nil {
		return fmt.Errorf("getting FutureReservation %s: %w", a.id, err)
	}

	status := &krm.FutureReservationStatus{}
	status.ObservedState = FutureReservationObservedState_FromProto(mapCtx, created)
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *FutureReservationAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating FutureReservation", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := FutureReservationSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	resource.Name = direct.LazyPtr(a.id.ID())

	paths, err := common.CompareProtoMessage(resource, a.actual, common.BasicDiff)
	if err != nil {
		return err
	}

	updated := a.actual
	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
	} else {
		log.V(2).Info("fields need update", "name", a.id, "paths", paths)
		updateMask := strings.Join(sets.List(paths), ",")

		req := &computepb.UpdateFutureReservationRequest{
			Project:                   a.id.Parent().ProjectID,
			Zone:                      a.id.Parent().Location,
			UpdateMask:                direct.LazyPtr(updateMask),
			FutureReservation:         a.id.ID(),
			FutureReservationResource: resource,
		}
		op, err := a.gcpClient.Update(ctx, req)
		if err != nil {
			return fmt.Errorf("updating FutureReservation %s: %w", a.id.String(), err)
		}
		err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("FutureReservation %s waiting for update: %w", a.id.String(), err)
		}
		log.V(2).Info("successfully updated FutureReservation", "name", a.id.String())

		// Get the updated resource
		updated, err = a.get(ctx)
		if err != nil {
			return fmt.Errorf("getting FutureReservation %s: %w", a.id, err)
		}
	}

	status := &krm.FutureReservationStatus{}
	status.ObservedState = FutureReservationObservedState_FromProto(mapCtx, updated)
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *FutureReservationAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.FutureReservation{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(FutureReservationSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Parent().ProjectID}
	obj.Spec.Location = a.id.Parent().Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.ID())
	u.SetGroupVersionKind(krm.FutureReservationGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *FutureReservationAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting FutureReservation", "name", a.id)

	req := &computepb.DeleteFutureReservationRequest{
		Project:           a.id.Parent().ProjectID,
		Zone:              a.id.Parent().Location,
		FutureReservation: a.id.ID(),
	}
	op, err := a.gcpClient.Delete(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent FutureReservation, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting FutureReservation %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted FutureReservation", "name", a.id)

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting for deletion of FutureReservation %s: %w", a.id, err)
	}
	return true, nil
}

func (a *FutureReservationAdapter) get(ctx context.Context) (*computepb.FutureReservation, error) {
	getReq := &computepb.GetFutureReservationRequest{
		Project:           a.id.Parent().ProjectID,
		Zone:              a.id.Parent().Location,
		FutureReservation: a.id.ID(),
	}
	resource, err := a.gcpClient.Get(ctx, getReq)
	if err != nil {
		return nil, fmt.Errorf("getting FutureReservation %s: %w", a.id, err)
	}
	return resource, nil
}
