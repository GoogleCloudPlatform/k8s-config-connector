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

package securitycenter

import (
	"context"
	"fmt"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/securitycenter/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	securitycenter "cloud.google.com/go/securitycenter/apiv1"
	pb "cloud.google.com/go/securitycenter/apiv1/securitycenterpb"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func NewBigQueryExportModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelBigQueryExport{config: *config}, nil
}

var _ directbase.Model = &modelBigQueryExport{}

type modelBigQueryExport struct {
	config config.ControllerConfig
}

func (m *modelBigQueryExport) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.SecurityCenterBigQueryExport{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}

	gcpClient, err := newGCPClient(ctx, &m.config)
	if err != nil {
		return nil, err
	}
	client, err := gcpClient.newSecurityCenterClient(ctx)
	if err != nil {
		return nil, err
	}

	return &BigQueryExportAdapter{
		id:        id.(*krm.SecurityCenterBigQueryExportIdentity),
		gcpClient: client,
		desired:   obj,
		reader:    reader,
	}, nil
}

func (m *modelBigQueryExport) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type BigQueryExportAdapter struct {
	id        *krm.SecurityCenterBigQueryExportIdentity
	gcpClient *securitycenter.Client
	desired   *krm.SecurityCenterBigQueryExport
	actual    *pb.BigQueryExport
	reader    client.Reader
}

var _ directbase.Adapter = &BigQueryExportAdapter{}

func (a *BigQueryExportAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting SecurityCenterBigQueryExport", "name", a.id)

	res, err := a.gcpClient.GetBigQueryExport(ctx, &pb.GetBigQueryExportRequest{Name: a.id.String()})
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting SecurityCenterBigQueryExport %q: %w", a.id, err)
	}

	a.actual = res
	return true, nil
}

func (a *BigQueryExportAdapter) normalizeReferences(ctx context.Context) error {
	if a.desired.Spec.DatasetRef != nil {
		if err := a.desired.Spec.DatasetRef.Normalize(ctx, a.reader, a.desired.Namespace); err != nil {
			return err
		}
	}
	return nil
}

func (a *BigQueryExportAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating SecurityCenterBigQueryExport", "name", a.id)

	if err := a.normalizeReferences(ctx); err != nil {
		return err
	}

	mapCtx := &direct.MapContext{}
	desiredPb := SecurityCenterBigQueryExportSpec_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &pb.CreateBigQueryExportRequest{
		Parent:           a.id.Parent().String(),
		BigQueryExport:   desiredPb,
		BigQueryExportId: a.id.ID(),
	}
	created, err := a.gcpClient.CreateBigQueryExport(ctx, req)
	if err != nil {
		return fmt.Errorf("creating SecurityCenterBigQueryExport %s: %w", a.id, err)
	}

	log.V(2).Info("successfully created SecurityCenterBigQueryExport", "name", a.id)

	status := &krm.SecurityCenterBigQueryExportStatus{}
	status.ObservedState = SecurityCenterBigQueryExportObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *BigQueryExportAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating SecurityCenterBigQueryExport", "name", a.id)

	if err := a.normalizeReferences(ctx); err != nil {
		return err
	}

	mapCtx := &direct.MapContext{}
	desiredPb := SecurityCenterBigQueryExportSpec_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	paths, report, err := common.CompareProtoMessageStructuredDiff(desiredPb, a.actual, common.BasicDiff)
	if err != nil {
		return err
	}

	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		if a.desired.Status.ExternalRef == nil {
			status := &krm.SecurityCenterBigQueryExportStatus{}
			status.ObservedState = SecurityCenterBigQueryExportObservedState_FromProto(mapCtx, a.actual)
			status.ExternalRef = direct.LazyPtr(a.id.String())
			return updateOp.UpdateStatus(ctx, status, nil)
		}
		return nil
	}

	report.Object = updateOp.GetUnstructured()
	structuredreporting.ReportDiff(ctx, report)

	updateOp.RecordUpdatingEvent()

	desiredPb.Name = a.id.String()
	req := &pb.UpdateBigQueryExportRequest{
		BigQueryExport: desiredPb,
		UpdateMask:     &fieldmaskpb.FieldMask{Paths: sets.List(paths)},
	}
	updated, err := a.gcpClient.UpdateBigQueryExport(ctx, req)
	if err != nil {
		return fmt.Errorf("updating SecurityCenterBigQueryExport %s: %w", a.id, err)
	}

	log.V(2).Info("successfully updated SecurityCenterBigQueryExport", "name", a.id)

	status := &krm.SecurityCenterBigQueryExportStatus{}
	status.ObservedState = SecurityCenterBigQueryExportObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	if a.desired.Status.ExternalRef == nil {
		status.ExternalRef = direct.LazyPtr(a.id.String())
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *BigQueryExportAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting SecurityCenterBigQueryExport", "name", a.id)

	err := a.gcpClient.DeleteBigQueryExport(ctx, &pb.DeleteBigQueryExportRequest{Name: a.id.String()})
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting SecurityCenterBigQueryExport %s: %w", a.id, err)
	}

	log.V(2).Info("successfully deleted SecurityCenterBigQueryExport", "name", a.id)
	return true, nil
}

func (a *BigQueryExportAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	return nil, nil
}
