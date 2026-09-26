// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

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
	desiredPb.Name = id.String()

	// Get bigqueryreservation GCP client
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
	// TODO: implement export support
	return nil, nil
}

type ReservationGroupAdapter struct {
	id        *krm.BigQueryReservationReservationGroupIdentity
	gcpClient *gcp.Client
	desired   *pb.ReservationGroup
}

var _ directbase.Adapter = &ReservationGroupAdapter{}

func (a *ReservationGroupAdapter) Find(ctx context.Context) (bool, error) {
	if a.id.ReservationGroup == "" {
		return false, nil
	}

	req := &pb.GetReservationGroupRequest{
		Name: a.id.String(),
	}
	actual, err := a.gcpClient.GetReservationGroup(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ReservationGroup %q: %w", a.id.String(), err)
	}

	if err := a.id.FromExternal(actual.Name); err != nil {
		return false, err
	}
	return true, nil
}

func (a *ReservationGroupAdapter) Create(ctx context.Context, op *directbase.CreateOperation) error {
	u := op.GetUnstructured()
	klog.V(2).Infof("creating ReservationGroup %s/%s", u.GetNamespace(), u.GetName())

	req := &pb.CreateReservationGroupRequest{
		Parent:             a.id.ParentString(),
		ReservationGroupId: a.id.ReservationGroup,
		ReservationGroup:   a.desired,
	}
	created, err := a.gcpClient.CreateReservationGroup(ctx, req)
	if err != nil {
		return fmt.Errorf("creating ReservationGroup %s: %w", a.id.String(), err)
	}

	// GET after create to ensure we have all fields (though for ReservationGroup there aren't many)
	actual, err := a.gcpClient.GetReservationGroup(ctx, &pb.GetReservationGroupRequest{Name: created.Name})
	if err != nil {
		return fmt.Errorf("getting ReservationGroup %s after create: %w", created.Name, err)
	}

	return a.updateStatus(ctx, op, actual)
}

func (a *ReservationGroupAdapter) Update(ctx context.Context, op *directbase.UpdateOperation) error {
	u := op.GetUnstructured()
	klog.V(2).Infof("updating ReservationGroup %s/%s", u.GetNamespace(), u.GetName())

	req := &pb.GetReservationGroupRequest{
		Name: a.id.String(),
	}
	actual, err := a.gcpClient.GetReservationGroup(ctx, req)
	if err != nil {
		return fmt.Errorf("getting ReservationGroup %s: %w", a.id.String(), err)
	}

	diffs, updateMask, err := a.compare(ctx, actual, a.desired)
	if err != nil {
		return err
	}

	if !diffs.HasDiff() {
		klog.V(2).Infof("no diff for ReservationGroup %s/%s", u.GetNamespace(), u.GetName())
		return a.updateStatus(ctx, op, actual)
	}

	if updateMask == nil || len(updateMask.Paths) == 0 {
		klog.V(2).Infof("no mutable diff for ReservationGroup %s/%s", u.GetNamespace(), u.GetName())
		return a.updateStatus(ctx, op, actual)
	}

	structuredreporting.ReportDiff(ctx, diffs)

	updateReq := &pb.UpdateReservationGroupRequest{
		ReservationGroup: a.desired,
		UpdateMask:       updateMask,
	}
	updated, err := a.gcpClient.UpdateReservationGroup(ctx, updateReq)
	if err != nil {
		return fmt.Errorf("updating ReservationGroup %s: %w", a.id.String(), err)
	}

	// GET after update
	actual, err = a.gcpClient.GetReservationGroup(ctx, &pb.GetReservationGroupRequest{Name: updated.Name})
	if err != nil {
		return fmt.Errorf("getting ReservationGroup %s after update: %w", updated.Name, err)
	}

	return a.updateStatus(ctx, op, actual)
}

func (a *ReservationGroupAdapter) compare(ctx context.Context, actual, desired *pb.ReservationGroup) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, BigQueryReservationReservationGroupSpec_v1alpha1_FromProto, BigQueryReservationReservationGroupSpec_v1alpha1_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = desired.Name

	return common.DiffForTopLevelFields(ctx, desired.ProtoReflect(), maskedActual.ProtoReflect())
}

func (a *ReservationGroupAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.ReservationGroup) error {
	mapCtx := &direct.MapContext{}
	status := BigQueryReservationReservationGroupStatus_v1alpha1_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(latest.Name)
	return op.UpdateStatus(ctx, status, nil)
}

func (a *ReservationGroupAdapter) Delete(ctx context.Context, op *directbase.DeleteOperation) (bool, error) {
	u := op.GetUnstructured()
	klog.V(2).Infof("deleting ReservationGroup %s/%s", u.GetNamespace(), u.GetName())

	req := &pb.DeleteReservationGroupRequest{
		Name: a.id.String(),
	}
	err := a.gcpClient.DeleteReservationGroup(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting ReservationGroup %s: %w", a.id.String(), err)
	}

	return true, nil
}

func (a *ReservationGroupAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	return nil, nil
}
