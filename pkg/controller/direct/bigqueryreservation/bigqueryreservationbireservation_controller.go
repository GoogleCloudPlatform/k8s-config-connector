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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.BigQueryReservationBiReservationGVK, NewBiReservationModel)
}

func NewBiReservationModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelBiReservation{config: *config}, nil
}

var _ directbase.Model = &modelBiReservation{}

type modelBiReservation struct {
	config config.ControllerConfig
}

func (m *modelBiReservation) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building BiReservation client: %w", err)
	}
	return gcpClient, err
}

func (m *modelBiReservation) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.BigQueryReservationBiReservation{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	idInterface, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := idInterface.(*krm.BigQueryReservationBiReservationIdentity)

	mapCtx := &direct.MapContext{}
	desiredPb := BigQueryReservationBiReservationSpec_v1alpha1_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	// Get bigqueryreservation GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &BiReservationAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   desiredPb,
		krmObj:    obj,
	}, nil
}

func (m *modelBiReservation) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type BiReservationAdapter struct {
	id        *krm.BigQueryReservationBiReservationIdentity
	gcpClient *gcp.Client
	desired   *pb.BiReservation
	actual    *pb.BiReservation
	krmObj    *krm.BigQueryReservationBiReservation
}

var _ directbase.Adapter = &BiReservationAdapter{}

// Find retrieves the GCP resource.
func (a *BiReservationAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting BiReservation", "name", a.id.String())

	req := &pb.GetBiReservationRequest{Name: a.id.String()}
	bireservationpb, err := a.gcpClient.GetBiReservation(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting BiReservation %q: %w", a.id.String(), err)
	}

	a.actual = bireservationpb
	return true, nil
}

// Create is a no-op for singleton resource
func (a *BiReservationAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating BiReservation is a no-op since it is a singleton", "name", a.id.String())
	return nil
}

// Update updates the resource in GCP
func (a *BiReservationAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating BiReservation", "name", a.id.String())

	diffs, updateMask, err := a.diff(ctx)
	if err != nil {
		return err
	}

	if !diffs.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id.String())
		return a.updateStatus(ctx, updateOp, a.actual)
	}

	structuredreporting.ReportDiff(ctx, diffs)

	req := &pb.UpdateBiReservationRequest{
		BiReservation: a.desired,
		UpdateMask:    updateMask,
	}

	// Make sure Name is set on update
	req.BiReservation.Name = a.id.String()

	updated, err := a.gcpClient.UpdateBiReservation(ctx, req)
	if err != nil {
		return fmt.Errorf("updating BiReservation %s: %w", a.id.String(), err)
	}

	return a.updateStatus(ctx, updateOp, updated)
}

func (a *BiReservationAdapter) diff(ctx context.Context) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	mapCtx := &direct.MapContext{}

	// Mask actual to only contain spec fields for correct diffing
	maskedActualSpec := BigQueryReservationBiReservationSpec_v1alpha1_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return nil, nil, mapCtx.Err()
	}
	maskedActual := BigQueryReservationBiReservationSpec_v1alpha1_ToProto(mapCtx, maskedActualSpec)
	if mapCtx.Err() != nil {
		return nil, nil, mapCtx.Err()
	}

	clonedDesired := proto.Clone(a.desired).(*pb.BiReservation)

	return common.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
}

func (a *BiReservationAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.BiReservation) error {
	mapCtx := &direct.MapContext{}
	status := &krm.BigQueryReservationBiReservationStatus{}
	status.ObservedState = BigQueryReservationBiReservationObservedState_v1alpha1_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return op.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *BiReservationAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.BigQueryReservationBiReservation{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(BigQueryReservationBiReservationSpec_v1alpha1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Project}
	obj.Spec.Location = a.id.Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	// Use the static name for the singleton resource.
	u.SetName("bireservation")
	u.SetGroupVersionKind(krm.BigQueryReservationBiReservationGVK)
	u.Object = uObj
	return u, nil
}

// Delete is a no-op for singleton resource
func (a *BiReservationAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting BiReservation is a no-op since it is a singleton", "name", a.id.String())
	return true, nil
}
