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
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
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
		return nil, fmt.Errorf("building ReservationGroup REST client: %w", err)
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

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	idInterface, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := idInterface.(*krm.BigQueryReservationReservationGroupIdentity)

	mapCtx := &direct.MapContext{}
	desiredPb := BigQueryReservationReservationGroupSpec_v1alpha1_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &ReservationGroupAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   desiredPb,
	}, nil
}

func (m *modelReservationGroup) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	id := &krm.BigQueryReservationReservationGroupIdentity{}
	if err := id.FromExternal(url); err != nil {
		// Not recognized
		return nil, nil
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &ReservationGroupAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}

type ReservationGroupAdapter struct {
	id        *krm.BigQueryReservationReservationGroupIdentity
	gcpClient *gcp.Client
	desired   *pb.ReservationGroup
	actual    *pb.ReservationGroup
}

var _ directbase.Adapter = &ReservationGroupAdapter{}

func (a *ReservationGroupAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting BigQueryReservationReservationGroup", "name", a.id.String())

	req := &pb.GetReservationGroupRequest{Name: a.id.String()}
	rg, err := a.gcpClient.GetReservationGroup(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ReservationGroup %q: %w", a.id.String(), err)
	}

	a.actual = rg
	return true, nil
}

func (a *ReservationGroupAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating ReservationGroup", "name", a.id.String())

	req := &pb.CreateReservationGroupRequest{
		Parent:             a.id.ParentString(),
		ReservationGroupId: a.id.ReservationGroup,
		ReservationGroup:   a.desired,
	}

	created, err := a.gcpClient.CreateReservationGroup(ctx, req)
	if err != nil {
		return fmt.Errorf("creating ReservationGroup %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully created ReservationGroup", "name", created.Name)

	return a.updateStatus(ctx, createOp, created)
}

func (a *ReservationGroupAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating ReservationGroup", "name", a.id.String())

	desiredPb := a.desired
	desiredPb.Name = a.id.String()

	diffs, err := compareReservationGroup(ctx, a.actual, desiredPb)
	if err != nil {
		return err
	}

	if diffs.HasDiff() {
		diffs.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, diffs)
		return fmt.Errorf("BigQueryReservationReservationGroup is immutable and cannot be updated. Detected changes: %v", diffs.Fields)
	}

	return a.updateStatus(ctx, updateOp, a.actual)
}

func (a *ReservationGroupAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.ReservationGroup) error {
	status := &krm.BigQueryReservationReservationGroupStatus{}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return op.UpdateStatus(ctx, status, nil)
}

func (a *ReservationGroupAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.BigQueryReservationReservationGroup{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(BigQueryReservationReservationGroupSpec_v1alpha1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.Location = &a.id.Location
	obj.Spec.ResourceID = &a.id.ReservationGroup
	obj.Spec.ProjectRef = &refsv1beta1.ProjectRef{External: a.id.Project}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.ReservationGroup)
	u.SetGroupVersionKind(krm.BigQueryReservationReservationGroupGVK)
	return u, nil
}

func (a *ReservationGroupAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ReservationGroup", "name", a.id.String())

	req := &pb.DeleteReservationGroupRequest{Name: a.id.String()}
	err := a.gcpClient.DeleteReservationGroup(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent ReservationGroup, assuming it was already deleted", "name", a.id.String())
			return true, nil
		}
		return false, fmt.Errorf("deleting ReservationGroup %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully deleted ReservationGroup", "name", a.id.String())
	return true, nil
}

func compareReservationGroup(ctx context.Context, actual, desired *pb.ReservationGroup) (*structuredreporting.Diff, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, BigQueryReservationReservationGroupSpec_v1alpha1_FromProto, BigQueryReservationReservationGroupSpec_v1alpha1_ToProto)
	if err != nil {
		return nil, err
	}
	maskedActual.Name = desired.Name
	diffs, _, err := common.DiffForTopLevelFields(ctx, desired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, err
	}
	return diffs, nil
}
