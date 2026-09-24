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
// proto.service: google.cloud.dataplex.v1.CatalogService
// proto.message: google.cloud.dataplex.v1.AspectType
// crd.type: DataplexAspectType
// crd.version: v1alpha1

package dataplex

import (
	"context"
	"fmt"

	gcp "cloud.google.com/go/dataplex/apiv1"
	pb "cloud.google.com/go/dataplex/apiv1/dataplexpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataplex/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
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
	registry.RegisterModel(krm.DataplexAspectTypeGVK, NewAspectTypeModel)
}

func NewAspectTypeModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &aspectTypeModel{config: config}, nil
}

var _ directbase.Model = &aspectTypeModel{}

type aspectTypeModel struct {
	config *config.ControllerConfig
}

func (m *aspectTypeModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.DataplexAspectType{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	copied := obj.DeepCopy()
	mapCtx := &direct.MapContext{}
	desired := DataplexAspectTypeSpec_ToProto(mapCtx, &copied.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	idI, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id, ok := idI.(*krm.AspectTypeIdentity)
	if !ok {
		return nil, fmt.Errorf("unexpected identity type %T", idI)
	}

	adapter := &aspectTypeAdapter{
		id:      id,
		desired: desired,
		reader:  reader,
	}

	// Get GCP client
	gcpClient, err := newGCPClient(ctx, m.config)
	if err != nil {
		return nil, fmt.Errorf("building gcp client: %w", err)
	}
	catalogClient, err := gcpClient.catalogClient(ctx)
	if err != nil {
		return nil, err
	}
	adapter.gcpClient = catalogClient

	return adapter, nil
}

func (m *aspectTypeModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type aspectTypeAdapter struct {
	gcpClient *gcp.CatalogClient
	id        *krm.AspectTypeIdentity
	desired   *pb.AspectType
	actual    *pb.AspectType
	reader    client.Reader
}

var _ directbase.Adapter = &aspectTypeAdapter{}

func (a *aspectTypeAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting dataplex aspecttype", "name", a.id)

	req := &pb.GetAspectTypeRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetAspectType(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting dataplex aspecttype %q from gcp: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *aspectTypeAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating dataplex aspecttype", "name", a.id)

	req := &pb.CreateAspectTypeRequest{
		Parent:       a.id.ParentString(),
		AspectTypeId: a.id.AspectType,
		AspectType:   a.desired,
	}

	op, err := a.gcpClient.CreateAspectType(ctx, req)
	if err != nil {
		return fmt.Errorf("creating dataplex aspecttype %s: %w", a.id.String(), err)
	}

	_, err = op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting create dataplex aspecttype %s failed: %w", a.id, err)
	}

	log.V(2).Info("successfully created dataplex aspecttype in gcp", "name", a.id)

	// Fetch fully-populated resource after creation
	actual, err := a.gcpClient.GetAspectType(ctx, &pb.GetAspectTypeRequest{Name: a.id.String()})
	if err != nil {
		return fmt.Errorf("fetching dataplex aspecttype %s after create: %w", a.id, err)
	}

	return a.updateStatus(ctx, createOp, actual)
}

func (a *aspectTypeAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating dataplex aspecttype", "name", a.id)

	mapCtx := &direct.MapContext{}

	// Mask actual to only contain spec fields for correct diffing
	maskedActualSpec := DataplexAspectTypeSpec_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	maskedActual := DataplexAspectTypeSpec_ToProto(mapCtx, maskedActualSpec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	clonedDesired := proto.Clone(a.desired).(*pb.AspectType)

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

	req := &pb.UpdateAspectTypeRequest{
		AspectType: clonedDesired,
		UpdateMask: updateMask,
	}

	req.AspectType.Name = a.id.String()

	op, err := a.gcpClient.UpdateAspectType(ctx, req)
	if err != nil {
		return fmt.Errorf("updating dataplex aspecttype %s: %w", a.id.String(), err)
	}

	_, err = op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for update of dataplex aspecttype %s: %w", a.id.String(), err)
	}

	log.V(2).Info("successfully updated dataplex aspecttype", "name", a.id)

	// Fetch fully-populated resource after update
	actual, err := a.gcpClient.GetAspectType(ctx, &pb.GetAspectTypeRequest{Name: a.id.String()})
	if err != nil {
		return fmt.Errorf("fetching dataplex aspecttype %s after update: %w", a.id, err)
	}

	return a.updateStatus(ctx, updateOp, actual)
}

func (a *aspectTypeAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.AspectType) error {
	mapCtx := &direct.MapContext{}
	status := &krm.DataplexAspectTypeStatus{}
	status.ObservedState = DataplexAspectTypeObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.PtrTo(a.id.String())
	return op.UpdateStatus(ctx, status, nil)
}

func (a *aspectTypeAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	log := klog.FromContext(ctx)

	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}

	obj := &krm.DataplexAspectType{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(DataplexAspectTypeSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ParentRef = &parent.ProjectAndLocationRef{
		ProjectRef: &refs.ProjectRef{External: a.id.Project},
		Location:   a.id.Location,
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u := &unstructured.Unstructured{Object: uObj}
	u.SetName(a.id.AspectType)
	u.SetGroupVersionKind(krm.DataplexAspectTypeGVK)

	log.Info("exported object", "obj", u, "gvk", u.GroupVersionKind())
	return u, nil
}

// Delete implements the Adapter interface.
func (a *aspectTypeAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting dataplex aspecttype", "name", a.id)

	req := &pb.DeleteAspectTypeRequest{
		Name: a.id.String(),
	}
	op, err := a.gcpClient.DeleteAspectType(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent dataplex aspecttype, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting dataplex aspecttype %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully deleted dataplex aspecttype", "name", a.id)

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting for deletion of dataplex aspecttype %s: %w", a.id.String(), err)
	}
	return true, nil
}
