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

package bigqueryreservation

import (
	"context"
	"fmt"

	gcp "cloud.google.com/go/bigquery/reservation/apiv1"
	pb "cloud.google.com/go/bigquery/reservation/apiv1/reservationpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigqueryreservation/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"google.golang.org/api/option"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.BigQueryReservationReservationGroupGVK, NewReservationGroupModel)
}

func NewReservationGroupModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelReservationGroup{config: *config}, nil
}

var _ directbase.Model = &modelReservationGroup{}

type modelReservationGroup struct {
	config config.ControllerConfig
}

func (m *modelReservationGroup) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building ReservationGroup client: %w", err)
	}
	return gcpClient, err
}

func (m *modelReservationGroup) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.BigQueryReservationReservationGroup{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}

	typedID, ok := id.(*krm.BigQueryReservationReservationGroupIdentity)
	if !ok {
		return nil, fmt.Errorf("unexpected identity type %T", id)
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &ReservationGroupAdapter{
		id:        typedID,
		inner:     obj,
		gcpClient: gcpClient,
	}, nil
}

func (m *modelReservationGroup) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type ReservationGroupAdapter struct {
	id        *krm.BigQueryReservationReservationGroupIdentity
	inner     *krm.BigQueryReservationReservationGroup
	gcpClient *gcp.Client
}

var _ directbase.Adapter = &ReservationGroupAdapter{}

func (a *ReservationGroupAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ReservationGroup", "name", a.id.String())

	req := &pb.GetReservationGroupRequest{
		Name: a.id.String(),
	}

	reservationGroup, err := a.gcpClient.GetReservationGroup(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ReservationGroup %q: %w", a.id.String(), err)
	}

	mapCtx := &direct.MapContext{}
	observedState := BigQueryReservationReservationGroupObservedState_FromProto(mapCtx, reservationGroup)
	if mapCtx.Err() != nil {
		return true, mapCtx.Err()
	}
	if observedState != nil {
		a.inner.Status.ObservedState = observedState
	}

	return true, nil
}

func (a *ReservationGroupAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating ReservationGroup", "name", a.id.String())

	mapCtx := &direct.MapContext{}
	reservationGroup := BigQueryReservationReservationGroupSpec_ToProto(mapCtx, &a.inner.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	parent := fmt.Sprintf("projects/%s/locations/%s", a.id.Project, a.id.Location)
	req := &pb.CreateReservationGroupRequest{
		Parent:             parent,
		ReservationGroupId: a.id.Reservation_group,
		ReservationGroup:   reservationGroup,
	}

	created, err := a.gcpClient.CreateReservationGroup(ctx, req)
	if err != nil {
		return fmt.Errorf("creating ReservationGroup %q: %w", a.id.String(), err)
	}

	log.V(2).Info("successfully created ReservationGroup", "name", a.id.String())

	observedState := BigQueryReservationReservationGroupObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	if observedState != nil {
		a.inner.Status.ObservedState = observedState
	}

	return nil
}

func (a *ReservationGroupAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	// Update is not supported by the GCP API.
	return nil
}

func (a *ReservationGroupAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.inner == nil {
		return nil, fmt.Errorf("export error: inner obj is nil")
	}
	obj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(a.inner)
	if err != nil {
		return nil, err
	}
	return &unstructured.Unstructured{Object: obj}, nil
}

func (a *ReservationGroupAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ReservationGroup", "name", a.id.String())

	req := &pb.DeleteReservationGroupRequest{
		Name: a.id.String(),
	}

	err := a.gcpClient.DeleteReservationGroup(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting ReservationGroup %q: %w", a.id.String(), err)
	}

	log.V(2).Info("successfully deleted ReservationGroup", "name", a.id.String())
	return false, nil
}
