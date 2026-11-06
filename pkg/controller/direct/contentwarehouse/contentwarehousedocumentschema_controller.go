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

// +generated:controller
// krm.group: contentwarehouse.cnrm.cloud.google.com
// krm.version: v1alpha1
// krm.kind: ContentWarehouseDocumentSchema

package contentwarehouse

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/projects"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/contentwarehouse/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	api "google.golang.org/api/contentwarehouse/v1"
	"google.golang.org/api/option"
	pb "google.golang.org/genproto/googleapis/cloud/contentwarehouse/v1"
	"google.golang.org/protobuf/encoding/protojson"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.ContentWarehouseDocumentSchemaGVK, NewDocumentSchemaModel)
}

func NewDocumentSchemaModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelDocumentSchema{config: config}, nil
}

var _ directbase.Model = &modelDocumentSchema{}

type modelDocumentSchema struct {
	config *config.ControllerConfig
}

func (m *modelDocumentSchema) client(ctx context.Context) (*api.Service, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := api.NewService(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building ContentWarehouse service: %w", err)
	}
	return gcpClient, err
}

func (m *modelDocumentSchema) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ContentWarehouseDocumentSchema{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	idVal, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := idVal.(*krm.ContentWarehouseDocumentSchemaIdentity)

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &DocumentSchemaAdapter{
		id:            id,
		gcpClient:     gcpClient,
		desired:       obj,
		projectMapper: m.config.ProjectMapper,
	}, nil
}

func (m *modelDocumentSchema) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//contentwarehouse.googleapis.com/") {
		return nil, nil
	}
	id := &krm.ContentWarehouseDocumentSchemaIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &DocumentSchemaAdapter{
		id:            id,
		gcpClient:     gcpClient,
		projectMapper: m.config.ProjectMapper,
	}, nil
}

type DocumentSchemaAdapter struct {
	id            *krm.ContentWarehouseDocumentSchemaIdentity
	gcpClient     *api.Service
	desired       *krm.ContentWarehouseDocumentSchema
	actual        *pb.DocumentSchema
	projectMapper *projects.ProjectMapper
}

var _ directbase.Adapter = &DocumentSchemaAdapter{}

func (a *DocumentSchemaAdapter) fullyQualifiedName(ctx context.Context) (string, error) {
	projectNumber, err := a.projectMapper.LookupProjectNumber(ctx, a.id.Project)
	if err != nil {
		return "", fmt.Errorf("error converting project ID %s to project number: %w", a.id.Project, err)
	}
	return fmt.Sprintf("projects/%d/locations/%s/documentSchemas/%s", projectNumber, a.id.Location, a.id.DocumentSchema), nil
}

func (a *DocumentSchemaAdapter) parentName(ctx context.Context) (string, error) {
	projectNumber, err := a.projectMapper.LookupProjectNumber(ctx, a.id.Project)
	if err != nil {
		return "", fmt.Errorf("error converting project ID %s to project number: %w", a.id.Project, err)
	}
	return fmt.Sprintf("projects/%d/locations/%s", projectNumber, a.id.Location), nil
}

func (a *DocumentSchemaAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting DocumentSchema", "name", a.id)
	if !a.id.HasIdentitySpecified() {
		return false, nil
	}

	name, err := a.fullyQualifiedName(ctx)
	if err != nil {
		return false, err
	}

	restObj, err := a.gcpClient.Projects.Locations.DocumentSchemas.Get(name).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting DocumentSchema %q: %w", a.id, err)
	}

	actual, err := DocumentSchemaRESTToProto(restObj)
	if err != nil {
		return false, err
	}

	a.actual = actual
	return true, nil
}

func (a *DocumentSchemaAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating DocumentSchema", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := ContentWarehouseDocumentSchemaSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	parent, err := a.parentName(ctx)
	if err != nil {
		return err
	}

	restObj, err := DocumentSchemaProtoToREST(resource)
	if err != nil {
		return err
	}

	createdREST, err := a.gcpClient.Projects.Locations.DocumentSchemas.Create(parent, restObj).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("creating DocumentSchema %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created DocumentSchema", "name", a.id)

	created, err := DocumentSchemaRESTToProto(createdREST)
	if err != nil {
		return err
	}

	status := &krm.ContentWarehouseDocumentSchemaStatus{}
	status.ObservedState = ContentWarehouseDocumentSchemaObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	parsedID := &krm.ContentWarehouseDocumentSchemaIdentity{}
	if err := parsedID.FromExternal(created.Name); err != nil {
		return err
	}
	a.id.DocumentSchema = parsedID.DocumentSchema
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *DocumentSchemaAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating DocumentSchema", "name", a.id)

	mapCtx := &direct.MapContext{}
	resource, report, _, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		ContentWarehouseDocumentSchemaSpec_FromProto,
		ContentWarehouseDocumentSchemaSpec_ToProto,
		nil,
	)
	if err != nil {
		return err
	}

	name, err := a.fullyQualifiedName(ctx)
	if err != nil {
		return err
	}

	updated := a.actual
	if !report.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id)
	} else {
		report.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, report)
		resource.Name = name

		restObj, err := DocumentSchemaProtoToREST(resource)
		if err != nil {
			return err
		}

		req := &api.GoogleCloudContentwarehouseV1UpdateDocumentSchemaRequest{
			DocumentSchema: restObj,
		}

		updatedREST, err := a.gcpClient.Projects.Locations.DocumentSchemas.Patch(name, req).Context(ctx).Do()
		if err != nil {
			return fmt.Errorf("updating DocumentSchema %s: %w", a.id, err)
		}
		log.V(2).Info("successfully updated DocumentSchema", "name", a.id)

		updated, err = DocumentSchemaRESTToProto(updatedREST)
		if err != nil {
			return err
		}
	}

	status := &krm.ContentWarehouseDocumentSchemaStatus{}
	status.ObservedState = ContentWarehouseDocumentSchemaObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *DocumentSchemaAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ContentWarehouseDocumentSchema{}
	mapCtx := &direct.MapContext{}
	spec := ContentWarehouseDocumentSchemaSpec_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	spec.ResourceID = direct.LazyPtr(a.id.DocumentSchema)
	spec.Location = direct.LazyPtr(a.id.Location)
	spec.ProjectRef = &refs.ProjectRef{External: a.id.Project}
	obj.Spec = *spec
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.DocumentSchema)
	u.SetGroupVersionKind(krm.ContentWarehouseDocumentSchemaGVK)

	return u, nil
}

func (a *DocumentSchemaAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting DocumentSchema", "name", a.id)

	name, err := a.fullyQualifiedName(ctx)
	if err != nil {
		return false, err
	}

	_, err = a.gcpClient.Projects.Locations.DocumentSchemas.Delete(name).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent DocumentSchema, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting DocumentSchema %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted DocumentSchema", "name", a.id)
	return true, nil
}

func DocumentSchemaRESTToProto(in *api.GoogleCloudContentwarehouseV1DocumentSchema) (*pb.DocumentSchema, error) {
	if in == nil {
		return nil, nil
	}
	bytes, err := in.MarshalJSON()
	if err != nil {
		return nil, fmt.Errorf("error marshaling REST struct: %w", err)
	}
	out := &pb.DocumentSchema{}
	unmarshalOpts := protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: true,
	}
	if err := unmarshalOpts.Unmarshal(bytes, out); err != nil {
		return nil, fmt.Errorf("error unmarshaling proto struct: %w", err)
	}
	return out, nil
}

func DocumentSchemaProtoToREST(in *pb.DocumentSchema) (*api.GoogleCloudContentwarehouseV1DocumentSchema, error) {
	if in == nil {
		return nil, nil
	}
	marshalOpts := protojson.MarshalOptions{
		UseProtoNames:   false,
		EmitUnpopulated: false,
	}
	bytes, err := marshalOpts.Marshal(in)
	if err != nil {
		return nil, fmt.Errorf("error marshaling proto struct: %w", err)
	}
	out := &api.GoogleCloudContentwarehouseV1DocumentSchema{}
	if err := json.Unmarshal(bytes, out); err != nil {
		return nil, fmt.Errorf("error unmarshaling REST struct: %w", err)
	}
	return out, nil
}
