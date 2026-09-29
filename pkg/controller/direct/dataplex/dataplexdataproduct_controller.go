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
// proto.service: google.cloud.dataplex.v1.DataProductService
// proto.message: google.cloud.dataplex.v1.DataProduct
// crd.type: DataplexDataProduct
// crd.version: v1alpha1

package dataplex

import (
	"context"
	"fmt"

	gcp "cloud.google.com/go/dataplex/apiv1"
	pb "cloud.google.com/go/dataplex/apiv1/dataplexpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataplex/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
	"google.golang.org/protobuf/proto"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.DataplexDataProductGVK, NewDataProductModel)
}

func NewDataProductModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &dataProductModel{config: config}, nil
}

var _ directbase.Model = &dataProductModel{}

type dataProductModel struct {
	config *config.ControllerConfig
}

func (m *dataProductModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.DataplexDataProduct{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	copied := obj.DeepCopy()
	mapCtx := &direct.MapContext{}
	desired := DataplexDataProductSpec_ToProto(mapCtx, &copied.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	idI, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id, ok := idI.(*krm.DataProductIdentity)
	if !ok {
		return nil, fmt.Errorf("unexpected identity type %T", idI)
	}

	adapter := &dataProductAdapter{
		id:      id,
		desired: desired,
		reader:  reader,
	}

	// Get GCP client
	gcpClient, err := newGCPClient(ctx, m.config)
	if err != nil {
		return nil, fmt.Errorf("building gcp client: %w", err)
	}
	client, err := gcpClient.dataProductClient(ctx)
	if err != nil {
		return nil, err
	}
	adapter.gcpClient = client

	return adapter, nil
}

func (m *dataProductModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type dataProductAdapter struct {
	gcpClient *gcp.DataProductClient
	id        *krm.DataProductIdentity
	desired   *pb.DataProduct
	actual    *pb.DataProduct
	reader    client.Reader
}

var _ directbase.Adapter = &dataProductAdapter{}

func (a *dataProductAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting dataplex dataproduct", "name", a.id)

	req := &pb.GetDataProductRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetDataProduct(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting dataplex dataproduct %q from gcp: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *dataProductAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating dataplex dataproduct", "name", a.id)

	req := &pb.CreateDataProductRequest{
		Parent:        a.id.ParentString(),
		DataProductId: a.id.DataProduct,
		DataProduct:   a.desired,
	}

	op, err := a.gcpClient.CreateDataProduct(ctx, req)
	if err != nil {
		return fmt.Errorf("creating dataplex dataproduct %s: %w", a.id.String(), err)
	}

	_, err = op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting create dataplex dataproduct %s failed: %w", a.id, err)
	}

	log.V(2).Info("successfully created dataplex dataproduct in gcp", "name", a.id)

	// Fetch fully-populated resource after creation
	actual, err := a.gcpClient.GetDataProduct(ctx, &pb.GetDataProductRequest{Name: a.id.String()})
	if err != nil {
		return fmt.Errorf("fetching dataplex dataproduct %s after create: %w", a.id, err)
	}

	return a.updateStatus(ctx, createOp, actual)
}

func (a *dataProductAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating dataplex dataproduct", "name", a.id)

	mapCtx := &direct.MapContext{}

	// Mask actual to only contain spec fields for correct diffing
	maskedActualSpec := DataplexDataProductSpec_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	maskedActual := DataplexDataProductSpec_ToProto(mapCtx, maskedActualSpec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	clonedDesired := proto.Clone(a.desired).(*pb.DataProduct)

	diffs, updateMask, err := common.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return err
	}

	if diffs == nil || !diffs.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id.String())
		return a.updateStatus(ctx, updateOp, a.actual)
	}

	diffs.Object = updateOp.GetUnstructured()
	structuredreporting.ReportDiff(ctx, diffs)

	req := &pb.UpdateDataProductRequest{
		DataProduct: clonedDesired,
		UpdateMask:  updateMask,
	}

	req.DataProduct.Name = a.id.String()

	op, err := a.gcpClient.UpdateDataProduct(ctx, req)
	if err != nil {
		return fmt.Errorf("updating dataplex dataproduct %s: %w", a.id.String(), err)
	}

	_, err = op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for update of dataplex dataproduct %s: %w", a.id.String(), err)
	}

	log.V(2).Info("successfully updated dataplex dataproduct", "name", a.id)

	// Fetch fully-populated resource after update
	actual, err := a.gcpClient.GetDataProduct(ctx, &pb.GetDataProductRequest{Name: a.id.String()})
	if err != nil {
		return fmt.Errorf("fetching dataplex dataproduct %s after update: %w", a.id, err)
	}

	return a.updateStatus(ctx, updateOp, actual)
}

func (a *dataProductAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.DataProduct) error {
	mapCtx := &direct.MapContext{}
	status := &krm.DataplexDataProductStatus{}
	status.ObservedState = DataplexDataProductObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.PtrTo(a.id.String())
	return op.UpdateStatus(ctx, status, nil)
}

func (a *dataProductAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	log := klog.FromContext(ctx)

	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}

	obj := &krm.DataplexDataProduct{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(DataplexDataProductSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, fmt.Errorf("converting to unstructured: %w", err)
	}

	u := &unstructured.Unstructured{Object: uObj}
	u.SetName(a.id.DataProduct)
	u.SetGroupVersionKind(krm.DataplexDataProductGVK)

	log.V(2).Info("exporting dataplex dataproduct", "name", a.id)
	return u, nil
}

func (a *dataProductAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting dataplex dataproduct", "name", a.id)

	req := &pb.DeleteDataProductRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteDataProduct(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting dataplex dataproduct %s: %w", a.id.String(), err)
	}

	log.V(2).Info("waiting for deletion of dataplex dataproduct", "name", a.id)
	err = op.Wait(ctx)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("waiting for delete of dataplex dataproduct %s: %w", a.id.String(), err)
	}

	log.V(2).Info("successfully deleted dataplex dataproduct", "name", a.id)
	return true, nil
}
