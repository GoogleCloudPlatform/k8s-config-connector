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
// proto.service: google.cloud.dataplex.v1.DataScanService
// proto.message: google.cloud.dataplex.v1.DataScan
// crd.type: DataplexDataScan
// crd.version: v1alpha1

package dataplex

import (
	"context"
	"fmt"

	gcp "cloud.google.com/go/dataplex/apiv1"
	pb "cloud.google.com/go/dataplex/apiv1/dataplexpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataplex/v1alpha1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/label"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.DataplexDataScanGVK, NewDataScanModel)
}

func NewDataScanModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &dataScanModel{config: config}, nil
}

var _ directbase.Model = &dataScanModel{}

type dataScanModel struct {
	config *config.ControllerConfig
}

func (m *dataScanModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.DataplexDataScan{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	copied := obj.DeepCopy()
	mapCtx := &direct.MapContext{}
	desired := DataplexDataScanSpec_ToProto(mapCtx, &copied.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	desired.Labels = label.NewGCPLabelsFromK8sLabels(u.GetLabels())

	idI, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id, ok := idI.(*krm.DataScanIdentity)
	if !ok {
		return nil, fmt.Errorf("unexpected identity type %T", idI)
	}

	adapter := &dataScanAdapter{
		id:      id,
		desired: desired,
		reader:  reader,
	}

	// Get GCP Client
	gcpClient, err := newGCPClient(ctx, m.config)
	if err != nil {
		return nil, fmt.Errorf("building gcp client: %w", err)
	}
	client, err := gcpClient.dataScanRESTClient(ctx)
	if err != nil {
		return nil, err
	}
	adapter.gcpClient = client

	return adapter, nil
}

func (m *dataScanModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type dataScanAdapter struct {
	gcpClient *gcp.DataScanClient
	id        *krm.DataScanIdentity
	desired   *pb.DataScan
	actual    *pb.DataScan
	reader    client.Reader
}

var _ directbase.Adapter = &dataScanAdapter{}

func (a *dataScanAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting dataplex datascan", "name", a.id)

	req := &pb.GetDataScanRequest{
		Name: a.id.String(),
		View: pb.GetDataScanRequest_FULL,
	}
	actual, err := a.gcpClient.GetDataScan(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting dataplex datascan %q: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *dataScanAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating dataplex datascan", "name", a.id)

	req := &pb.CreateDataScanRequest{
		Parent:     a.id.ParentString(),
		DataScan:   a.desired,
		DataScanId: a.id.DataScan,
	}
	op, err := a.gcpClient.CreateDataScan(ctx, req)
	if err != nil {
		return fmt.Errorf("creating dataplex datascan %s: %w", a.id.String(), err)
	}

	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting create dataplex datascan %s failed: %w", a.id, err)
	}

	log.V(2).Info("successfully created dataplex datascan in gcp", "name", a.id)

	reqGet := &pb.GetDataScanRequest{
		Name: a.id.String(),
		View: pb.GetDataScanRequest_FULL,
	}
	created, err = a.gcpClient.GetDataScan(ctx, reqGet)
	if err != nil {
		return fmt.Errorf("getting dataplex datascan %s after create: %w", a.id.String(), err)
	}

	return a.updateStatus(ctx, createOp, created)
}

func (a *dataScanAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating dataplex datascan", "name", a.id)

	a.desired.Name = a.id.String()

	diffs, updateMask, err := compareDataScan(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	var updated *pb.DataScan
	if !diffs.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id)
		updated = a.actual
	} else {
		structuredreporting.ReportDiff(ctx, diffs)
		req := &pb.UpdateDataScanRequest{
			UpdateMask: updateMask,
			DataScan:   a.desired,
		}
		op, err := a.gcpClient.UpdateDataScan(ctx, req)
		if err != nil {
			return fmt.Errorf("updating dataplex datascan %s: %w", a.id.String(), err)
		}
		updated, err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("waiting for update of dataplex datascan %s: %w", a.id.String(), err)
		}
		log.V(2).Info("successfully updated dataplex datascan", "name", a.id)

		reqGet := &pb.GetDataScanRequest{
			Name: a.id.String(),
			View: pb.GetDataScanRequest_FULL,
		}
		updated, err = a.gcpClient.GetDataScan(ctx, reqGet)
		if err != nil {
			return fmt.Errorf("getting dataplex datascan %s after update: %w", a.id.String(), err)
		}
	}

	return a.updateStatus(ctx, updateOp, updated)
}

func (a *dataScanAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	log := klog.FromContext(ctx)

	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}

	obj := &krm.DataplexDataScan{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(DataplexDataScanSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refsv1beta1.ProjectRef{External: a.id.Project}
	obj.Spec.Location = direct.PtrTo(a.id.Location)

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u := &unstructured.Unstructured{Object: uObj}
	u.SetName(a.id.DataScan)
	u.SetGroupVersionKind(krm.DataplexDataScanGVK)

	log.Info("exported object", "obj", u, "gvk", u.GroupVersionKind())
	return u, nil
}

func (a *dataScanAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting dataplex datascan", "name", a.id)

	req := &pb.DeleteDataScanRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteDataScan(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent DataplexDataScan, assuming it was already deleted", "name", a.id.String())
			return true, nil
		}
		return false, fmt.Errorf("deleting dataplex datascan %s: %w", a.id.String(), err)
	}
	err = op.Wait(ctx)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("waiting for deletion of dataplex datascan %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully deleted dataplex datascan", "name", a.id)
	return true, nil
}

func (a *dataScanAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.DataScan) error {
	mapCtx := &direct.MapContext{}
	status := &krm.DataplexDataScanStatus{}
	status.ObservedState = DataplexDataScanObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.PtrTo(a.id.String())
	return op.UpdateStatus(ctx, status, nil)
}

func compareDataScan(ctx context.Context, actual, desired *pb.DataScan) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, DataplexDataScanSpec_FromProto, DataplexDataScanSpec_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = desired.Name
	maskedActual.Labels = actual.Labels

	clonedDesired := proto.Clone(desired).(*pb.DataScan)

	populateDefaults := func(obj *pb.DataScan) {
		if obj.ExecutionSpec == nil {
			obj.ExecutionSpec = &pb.DataScan_ExecutionSpec{}
		}
		if obj.ExecutionSpec.Trigger == nil {
			obj.ExecutionSpec.Trigger = &pb.Trigger{
				Mode: &pb.Trigger_OnDemand_{
					OnDemand: &pb.Trigger_OnDemand{},
				},
			}
		}
		if len(obj.Labels) == 0 {
			obj.Labels = nil
		}
	}

	populateDefaults(maskedActual)
	populateDefaults(clonedDesired)

	diffs, updateMask, err := common.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}
