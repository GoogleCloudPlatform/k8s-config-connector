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

package contentwarehouseschema

import (
	"context"
	"fmt"

	api "google.golang.org/api/contentwarehouse/v1"
	"google.golang.org/api/option"
	pb "google.golang.org/genproto/googleapis/cloud/contentwarehouse/v1"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/contentwarehouse/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/contentwarehouse"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/tags"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.ContentWarehouseSchemaGVK, NewModel)
}

func NewModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &model{config: config}, nil
}

var _ directbase.Model = &model{}

type model struct {
	config *config.ControllerConfig
}

func (m *model) client(ctx context.Context) (*api.Service, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := api.NewService(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building ContentWarehouse client: %w", err)
	}
	return gcpClient, nil
}

func (m *model) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ContentWarehouseSchema{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	// Always call common.NormalizeReferences to resolve references
	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	idVal, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id, ok := idVal.(*krm.ContentWarehouseSchemaIdentity)
	if !ok {
		return nil, fmt.Errorf("unexpected identity type: %T", idVal)
	}

	// Convert the KRM spec to proto format
	mapCtx := &direct.MapContext{}
	desired := contentwarehouse.ContentWarehouseSchemaSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &adapter{
		id:      id,
		service: gcpClient,
		desired: desired,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type adapter struct {
	id      *krm.ContentWarehouseSchemaIdentity
	service *api.Service
	desired *pb.DocumentSchema
	actual  *api.GoogleCloudContentwarehouseV1DocumentSchema
}

var _ directbase.Adapter = &adapter{}

func (a *adapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	fqn := a.id.String()
	log.V(2).Info("getting ContentWarehouseSchema", "name", fqn)

	call := a.service.Projects.Locations.DocumentSchemas.Get(fqn)
	actual, err := call.Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ContentWarehouseSchema %q: %w", fqn, err)
	}

	a.actual = actual
	return true, nil
}

func (a *adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	fqn := a.id.String()
	log.V(2).Info("creating ContentWarehouseSchema", "name", fqn)

	parent := fmt.Sprintf("projects/%s/locations/%s", a.id.Project, a.id.Location)

	apiSchema := &api.GoogleCloudContentwarehouseV1DocumentSchema{}
	if err := common.ProtoToAPI(a.desired, apiSchema); err != nil {
		return err
	}

	call := a.service.Projects.Locations.DocumentSchemas.Create(parent, apiSchema)
	created, err := call.Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("creating ContentWarehouseSchema %s: %w", fqn, err)
	}

	createdPb := &pb.DocumentSchema{}
	if err := common.APIToProto(created, createdPb); err != nil {
		return err
	}

	return a.updateStatus(ctx, createOp, createdPb)
}

func (a *adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	fqn := a.id.String()
	log.V(2).Info("updating ContentWarehouseSchema", "name", fqn)

	diffs, _, err := a.compare(ctx)
	if err != nil {
		return err
	}

	var latestPb *pb.DocumentSchema = &pb.DocumentSchema{}
	if err := common.APIToProto(a.actual, latestPb); err != nil {
		return err
	}

	if diffs.HasDiff() {
		diffs.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, diffs)

		apiSchema := &api.GoogleCloudContentwarehouseV1DocumentSchema{}
		if err := common.ProtoToAPI(a.desired, apiSchema); err != nil {
			return err
		}

		req := &api.GoogleCloudContentwarehouseV1UpdateDocumentSchemaRequest{
			DocumentSchema: apiSchema,
		}

		call := a.service.Projects.Locations.DocumentSchemas.Patch(fqn, req)
		updated, err := call.Context(ctx).Do()
		if err != nil {
			return fmt.Errorf("updating ContentWarehouseSchema %s: %w", fqn, err)
		}

		latestPb = &pb.DocumentSchema{}
		if err := common.APIToProto(updated, latestPb); err != nil {
			return err
		}
	}

	return a.updateStatus(ctx, updateOp, latestPb)
}

func (a *adapter) compare(ctx context.Context) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	actualPb := &pb.DocumentSchema{}
	if err := common.APIToProto(a.actual, actualPb); err != nil {
		return nil, nil, err
	}

	maskedActual, err := mappers.OnlySpecFields(actualPb, contentwarehouse.ContentWarehouseSchemaSpec_FromProto, contentwarehouse.ContentWarehouseSchemaSpec_ToProto)
	if err != nil {
		return nil, nil, err
	}

	diffs, updateMask, err := tags.DiffForTopLevelFields(ctx, a.desired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}

	return diffs, updateMask, nil
}

func (a *adapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.DocumentSchema) error {
	mapCtx := &direct.MapContext{}
	status := &krm.ContentWarehouseSchemaStatus{}
	status.ObservedState = contentwarehouse.ContentWarehouseSchemaObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return op.UpdateStatus(ctx, status, nil)
}

func (a *adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ContentWarehouseSchema{}
	mapCtx := &direct.MapContext{}
	actualPb := &pb.DocumentSchema{}
	if err := common.APIToProto(a.actual, actualPb); err != nil {
		return nil, err
	}
	obj.Spec = direct.ValueOf(contentwarehouse.ContentWarehouseSchemaSpec_FromProto(mapCtx, actualPb))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.Location = &a.id.Location
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Project}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.DocumentSchema)
	u.SetGroupVersionKind(krm.ContentWarehouseSchemaGVK)
	u.Object = uObj
	return u, nil
}

func (a *adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	fqn := a.id.String()
	log.V(2).Info("deleting ContentWarehouseSchema", "name", fqn)

	call := a.service.Projects.Locations.DocumentSchemas.Delete(fqn)
	_, err := call.Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting ContentWarehouseSchema %s: %w", fqn, err)
	}
	return true, nil
}
