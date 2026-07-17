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
// proto.service: google.pubsub.v1.SchemaService
// proto.message: google.pubsub.v1.Schema
// crd.type: PubSubSchema
// crd.version: v1beta1

package pubsub

import (
	"context"
	"fmt"
	"time"

	api "cloud.google.com/go/pubsub/v2/apiv1"
	pb "cloud.google.com/go/pubsub/v2/apiv1/pubsubpb"

	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/pubsub/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/tags"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.PubSubSchemaGVK, NewSchemaModel)
}

func NewSchemaModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &SchemaModel{config: *config}, nil
}

var _ directbase.Model = &SchemaModel{}

type SchemaModel struct {
	config config.ControllerConfig
}

func (m *SchemaModel) client(ctx context.Context, projectID string) (*api.SchemaClient, error) {
	var opts []option.ClientOption

	config := m.config

	opts, err := config.RESTClientOptions()
	if err != nil {
		return nil, err
	}

	gcpClient, err := api.NewSchemaRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building pubsub schema client: %w", err)
	}

	return gcpClient, err
}

func (m *SchemaModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.PubSubSchema{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	schemaId := id.(*krm.PubSubSchemaIdentity)

	gcpClient, err := m.client(ctx, schemaId.Project)
	if err != nil {
		return nil, err
	}

	return &pubSubSchemaAdapter{
		gcpClient: gcpClient,
		id:        schemaId,
		desired:   obj,
	}, nil
}

func (m *SchemaModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type pubSubSchemaAdapter struct {
	gcpClient *api.SchemaClient
	id        *krm.PubSubSchemaIdentity
	desired   *krm.PubSubSchema
	actual    *pb.Schema
}

var _ directbase.Adapter = &pubSubSchemaAdapter{}

func (a *pubSubSchemaAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.Info("getting pubsub schema", "name", a.id)

	req := &pb.GetSchemaRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetSchema(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting pubsub schema %q from gcp: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *pubSubSchemaAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.Info("creating pubsub schema", "name", a.id)

	mapCtx := &direct.MapContext{}
	desired := PubSubSchemaSpec_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &pb.CreateSchemaRequest{
		Parent:   a.id.ParentString(),
		Schema:   desired,
		SchemaId: a.id.Schema,
	}

	created, err := a.gcpClient.CreateSchema(ctx, req)
	if err != nil {
		return fmt.Errorf("creating pubsub schema %s: %w", a.id.String(), err)
	}
	log.Info("successfully created pubsub schema in gcp", "name", a.id)

	return a.updateStatus(ctx, createOp, created)
}

func (a *pubSubSchemaAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating pubsub schema", "name", a.id)

	mapCtx := &direct.MapContext{}
	desired := PubSubSchemaSpec_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	diffs, _, err := compareSchema(ctx, a.actual, desired)
	if err != nil {
		return err
	}

	if !diffs.HasDiff() {
		log.V(2).Info("no diff detected for pubsub schema", "name", a.id)
		return a.updateStatus(ctx, updateOp, a.actual)
	}

	structuredreporting.ReportDiff(ctx, diffs)

	req := &pb.CommitSchemaRequest{
		Name:   a.id.String(),
		Schema: desired,
	}

	updatedSchema, err := a.gcpClient.CommitSchema(ctx, req)
	if err != nil {
		return fmt.Errorf("updating pubsub schema %s: %w", a.id, err)
	}
	log.V(2).Info("successfully updated pubsub schema", "name", a.id)

	return a.updateStatus(ctx, updateOp, updatedSchema)
}

func (a *pubSubSchemaAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.Info("deleting pubsub schema", "name", a.id)

	req := &pb.DeleteSchemaRequest{Name: a.id.String()}
	err := a.gcpClient.DeleteSchema(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting pubsub schema %s: %w", a.id.String(), err)
	}

	// Confirm deletion by wait-polling until 404 is returned (legacy controller behavior)
	for {
		getReq := &pb.GetSchemaRequest{Name: a.id.String()}
		_, err := a.gcpClient.GetSchema(ctx, getReq)
		if err != nil {
			if direct.IsNotFound(err) {
				break
			}
			return false, fmt.Errorf("waiting for deletion of pubsub schema %s: %w", a.id.String(), err)
		}
		select {
		case <-ctx.Done():
			return false, ctx.Err()
		case <-time.After(2 * time.Second):
		}
	}

	log.Info("successfully deleted pubsub schema", "name", a.id)

	return true, nil
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *pubSubSchemaAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.PubSubSchema{}
	mapCtx := &direct.MapContext{}
	obj.Spec = *PubSubSchemaSpec_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.Schema)
	u.SetGroupVersionKind(krm.PubSubSchemaGVK)

	return u, nil
}

func (a *pubSubSchemaAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.Schema) error {
	status := &krm.PubSubSchemaStatus{}
	return op.UpdateStatus(ctx, status, nil)
}

func compareSchema(ctx context.Context, actual, desired *pb.Schema) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, PubSubSchemaSpec_FromProto, PubSubSchemaSpec_ToProto)
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
