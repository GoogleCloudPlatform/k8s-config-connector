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

package contentwarehouse

import (
	"context"
	"fmt"
	"strconv"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/contentwarehouse/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	"google.golang.org/api/option"
	"google.golang.org/api/transport/grpc"
	pb "google.golang.org/genproto/googleapis/cloud/contentwarehouse/v1"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.ContentWarehouseSchemaGVK, NewContentWarehouseSchemaModel)
}

func NewContentWarehouseSchemaModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelContentWarehouseSchema{config: *config}, nil
}

var _ directbase.Model = &modelContentWarehouseSchema{}

type modelContentWarehouseSchema struct {
	config config.ControllerConfig
}

func (m *modelContentWarehouseSchema) client(ctx context.Context) (pb.DocumentSchemaServiceClient, error) {
	var opts []option.ClientOption

	opts, err := m.config.GRPCClientOptions()
	if err != nil {
		return nil, err
	}

	opts = append(opts, option.WithEndpoint("contentwarehouse.googleapis.com:443"))
	opts = append(opts, option.WithScopes("https://www.googleapis.com/auth/cloud-platform"))

	conn, err := grpc.Dial(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("dialing contentwarehouse service: %w", err)
	}

	return pb.NewDocumentSchemaServiceClient(conn), nil
}

func (m *modelContentWarehouseSchema) getProjectNumber(ctx context.Context, projectID string) (string, error) {
	// If the projectID is already numeric, return it.
	if _, err := strconv.ParseInt(projectID, 10, 64); err == nil {
		return projectID, nil
	}

	if m.config.ProjectMapper == nil {
		return "", fmt.Errorf("project mapper not initialized in ControllerConfig")
	}

	projectNumber, err := m.config.ProjectMapper.LookupProjectNumber(ctx, projectID)
	if err != nil {
		return "", fmt.Errorf("resolving project number for %s: %w", projectID, err)
	}

	return fmt.Sprintf("%d", projectNumber), nil
}

func (m *modelContentWarehouseSchema) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ContentWarehouseSchema{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	// Resolve alphanumeric project ID to its numeric project number early to avoid identity mismatch in GetIdentity
	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("resolving project ID: %w", err)
	}
	projectNumber, err := m.getProjectNumber(ctx, projectID)
	if err != nil {
		return nil, fmt.Errorf("resolving project number for %s: %w", projectID, err)
	}
	if obj.Spec.ProjectRef == nil {
		obj.Spec.ProjectRef = &refs.ProjectRef{}
	}
	obj.Spec.ProjectRef.External = projectNumber

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}

	resolvedIdentity := id.(*krm.ContentWarehouseSchemaIdentity)

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desiredPb := ContentWarehouseSchemaSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	desiredPb.Name = resolvedIdentity.String()

	return &ContentWarehouseSchemaAdapter{
		id:        resolvedIdentity,
		gcpClient: gcpClient,
		desired:   desiredPb,
	}, nil
}

func (m *modelContentWarehouseSchema) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type ContentWarehouseSchemaAdapter struct {
	id        *krm.ContentWarehouseSchemaIdentity
	gcpClient pb.DocumentSchemaServiceClient
	desired   *pb.DocumentSchema
	actual    *pb.DocumentSchema
}

var _ directbase.Adapter = &ContentWarehouseSchemaAdapter{}

func (a *ContentWarehouseSchemaAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ContentWarehouseSchema", "name", a.id)

	req := &pb.GetDocumentSchemaRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetDocumentSchema(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ContentWarehouseSchema %q: %w", a.id, err)
	}

	a.actual = actual
	return true, nil
}

func (a *ContentWarehouseSchemaAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	fqn := a.id.String()
	log.V(2).Info("creating ContentWarehouseSchema", "id", fqn)

	parent := a.id.ParentString()

	req := &pb.CreateDocumentSchemaRequest{
		Parent:         parent,
		DocumentSchema: a.desired,
	}
	created, err := a.gcpClient.CreateDocumentSchema(ctx, req)
	if err != nil {
		return fmt.Errorf("creating DocumentSchema %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created DocumentSchema", "name", a.id)

	return a.updateStatus(ctx, createOp, created)
}

func (a *ContentWarehouseSchemaAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating ContentWarehouseSchema", "name", a.id.String())

	diffs, _, err := compareDocumentSchema(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	latest := a.actual
	if diffs.HasDiff() {
		diffs.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, diffs)

		req := &pb.UpdateDocumentSchemaRequest{
			Name:           a.id.String(),
			DocumentSchema: a.desired,
		}

		updated, err := a.gcpClient.UpdateDocumentSchema(ctx, req)
		if err != nil {
			return fmt.Errorf("updating ContentWarehouseSchema %s: %w", a.id.String(), err)
		}
		latest = updated
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func (a *ContentWarehouseSchemaAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.DocumentSchema) error {
	mapCtx := &direct.MapContext{}
	status := &krm.ContentWarehouseSchemaStatus{}
	status.ObservedState = ContentWarehouseSchemaObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return op.UpdateStatus(ctx, status, nil)
}

func (a *ContentWarehouseSchemaAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ContentWarehouseSchema{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(ContentWarehouseSchemaSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.actual.Name)
	u.SetGroupVersionKind(krm.ContentWarehouseSchemaGVK)
	return u, nil
}

func (a *ContentWarehouseSchemaAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting DocumentSchema", "name", a.id)

	req := &pb.DeleteDocumentSchemaRequest{Name: a.id.String()}
	_, err := a.gcpClient.DeleteDocumentSchema(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent DocumentSchema, assuming it was already deleted", "name", a.id.String())
			return true, nil
		}
		return false, fmt.Errorf("deleting DocumentSchema %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted DocumentSchema", "name", a.id)
	return true, nil
}

func compareDocumentSchema(ctx context.Context, actual, desired *pb.DocumentSchema) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, ContentWarehouseSchemaSpec_FromProto, ContentWarehouseSchemaSpec_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = desired.Name // Restore any non-spec identifier fields if needed

	clonedDesired := proto.Clone(desired).(*pb.DocumentSchema)

	populateDefaults := func(obj *pb.DocumentSchema) {
		// Populate any server defaults here if needed
	}
	populateDefaults(maskedActual)
	populateDefaults(clonedDesired)

	diffs, updateMask, err := common.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}
