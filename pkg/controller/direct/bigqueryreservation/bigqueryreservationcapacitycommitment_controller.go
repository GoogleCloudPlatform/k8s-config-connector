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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/tags"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.BigQueryReservationCapacityCommitmentGVK, NewCapacityCommitmentModel)
}

func NewCapacityCommitmentModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelCapacityCommitment{config: *config}, nil
}

var _ directbase.Model = &modelCapacityCommitment{}

type modelCapacityCommitment struct {
	config config.ControllerConfig
}

func (m *modelCapacityCommitment) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building CapacityCommitment client: %w", err)
	}
	return gcpClient, err
}

func (m *modelCapacityCommitment) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.BigQueryReservationCapacityCommitment{}
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
	id := idInterface.(*krm.BigQueryReservationCapacityCommitmentIdentity)

	mapCtx := &direct.MapContext{}
	desiredPb := BigQueryReservationCapacityCommitmentSpec_v1alpha1_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	// Get bigqueryreservation GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	var enforceSingleAdminProject bool
	if obj.Spec.EnforceSingleAdminProjectPerOrg != nil {
		enforceSingleAdminProject = (*obj.Spec.EnforceSingleAdminProjectPerOrg == "true")
	}

	return &CapacityCommitmentAdapter{
		id:                        id,
		gcpClient:                 gcpClient,
		desired:                   desiredPb,
		enforceSingleAdminProject: enforceSingleAdminProject,
		actual:                    nil,
	}, nil
}

func (m *modelCapacityCommitment) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	id := &krm.BigQueryReservationCapacityCommitmentIdentity{}
	if err := id.FromExternal(url); err != nil {
		// Not recognized
		return nil, nil
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &CapacityCommitmentAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}

type CapacityCommitmentAdapter struct {
	id                        *krm.BigQueryReservationCapacityCommitmentIdentity
	gcpClient                 *gcp.Client
	desired                   *pb.CapacityCommitment
	enforceSingleAdminProject bool
	actual                    *pb.CapacityCommitment
}

var _ directbase.Adapter = &CapacityCommitmentAdapter{}

func (a *CapacityCommitmentAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting BigQueryReservationCapacityCommitment", "name", a.id.String())

	req := &pb.GetCapacityCommitmentRequest{Name: a.id.String()}
	cc, err := a.gcpClient.GetCapacityCommitment(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting CapacityCommitment %q: %w", a.id.String(), err)
	}

	a.actual = cc
	return true, nil
}

func (a *CapacityCommitmentAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating CapacityCommitment", "name", a.id.String())

	req := &pb.CreateCapacityCommitmentRequest{
		Parent:                          "projects/" + a.id.Project + "/locations/" + a.id.Location,
		CapacityCommitmentId:            a.id.CapacityCommitment,
		CapacityCommitment:              a.desired,
		EnforceSingleAdminProjectPerOrg: a.enforceSingleAdminProject,
	}

	created, err := a.gcpClient.CreateCapacityCommitment(ctx, req)
	if err != nil {
		return fmt.Errorf("creating CapacityCommitment %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully created CapacityCommitment", "name", created.Name)

	return a.updateStatus(ctx, createOp, created)
}

func (a *CapacityCommitmentAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating CapacityCommitment", "name", a.id.String())

	desiredPb := a.desired
	desiredPb.Name = a.id.String()

	diffs, updateMask, err := compareCapacityCommitment(ctx, a.actual, desiredPb)
	if err != nil {
		return err
	}

	latest := a.actual
	if diffs.HasDiff() {
		diffs.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, diffs)

		req := &pb.UpdateCapacityCommitmentRequest{
			CapacityCommitment: desiredPb,
			UpdateMask:         updateMask,
		}

		updated, err := a.gcpClient.UpdateCapacityCommitment(ctx, req)
		if err != nil {
			return fmt.Errorf("updating CapacityCommitment %s: %w", a.id.String(), err)
		}
		latest = updated
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func (a *CapacityCommitmentAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.CapacityCommitment) error {
	mapCtx := &direct.MapContext{}
	status := BigQueryReservationCapacityCommitmentStatus_v1alpha1_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return op.UpdateStatus(ctx, status, nil)
}

func (a *CapacityCommitmentAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.BigQueryReservationCapacityCommitment{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(BigQueryReservationCapacityCommitmentSpec_v1alpha1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.Location = a.id.Location
	obj.Spec.ResourceID = &a.id.CapacityCommitment
	obj.Spec.ProjectRef.External = a.id.Project

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.CapacityCommitment)
	u.SetGroupVersionKind(krm.BigQueryReservationCapacityCommitmentGVK)
	return u, nil
}

func (a *CapacityCommitmentAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting CapacityCommitment", "name", a.id.String())

	req := &pb.DeleteCapacityCommitmentRequest{Name: a.id.String()}
	err := a.gcpClient.DeleteCapacityCommitment(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent CapacityCommitment, assuming it was already deleted", "name", a.id.String())
			return true, nil
		}
		return false, fmt.Errorf("deleting CapacityCommitment %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully deleted CapacityCommitment", "name", a.id.String())
	return true, nil
}

func compareCapacityCommitment(ctx context.Context, actual, desired *pb.CapacityCommitment) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, BigQueryReservationCapacityCommitmentSpec_v1alpha1_FromProto, BigQueryReservationCapacityCommitmentSpec_v1alpha1_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = desired.Name
	diffs, updateMask, err := tags.DiffForTopLevelFields(ctx, desired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}
