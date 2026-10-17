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
// proto.service: google.cloud.discoveryengine.v1.SchemaService
// proto.message: google.cloud.discoveryengine.v1.Schema
// crd.type: DiscoveryEngineSchema
// crd.version: v1alpha1

package discoveryengine

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	gcp "cloud.google.com/go/discoveryengine/apiv1"
	pb "cloud.google.com/go/discoveryengine/apiv1/discoveryenginepb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/discoveryengine/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.DiscoveryEngineSchemaGVK, NewSchemaModel)
}

func NewSchemaModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &schemaModel{config: *config}, nil
}

var _ directbase.Model = &schemaModel{}

type schemaModel struct {
	config config.ControllerConfig
}

func (m *schemaModel) client(ctx context.Context, projectID string) (*gcp.SchemaClient, error) {
	var opts []option.ClientOption

	config := m.config

	if !config.UserProjectOverride || config.BillingProject == "" {
		config.UserProjectOverride = true
		config.BillingProject = projectID
	}

	opts, err := config.RESTClientOptions()
	if err != nil {
		return nil, err
	}

	gcpClient, err := gcp.NewSchemaRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building discoveryengine schema client: %w", err)
	}

	return gcpClient, err
}

func (m *schemaModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.DiscoveryEngineSchema{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	identity, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := identity.(*krm.DiscoveryEngineSchemaIdentity)

	mapCtx := &direct.MapContext{}
	desired := DiscoveryEngineSchemaSpec_v1alpha1_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	gcpClient, err := m.client(ctx, id.Project)
	if err != nil {
		return nil, err
	}

	return &schemaAdapter{
		gcpClient: gcpClient,
		id:        id,
		desired:   desired,
	}, nil
}

func (m *schemaModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	log := klog.FromContext(ctx)
	if strings.HasPrefix(url, "//discoveryengine.googleapis.com/") {
		trimmed := strings.TrimPrefix(url, "//discoveryengine.googleapis.com/")
		id := &krm.DiscoveryEngineSchemaIdentity{}
		if err := id.FromExternal(trimmed); err != nil {
			log.V(2).Error(err, "url did not match DiscoveryEngineSchema format", "url", url)
			return nil, nil
		}
		gcpClient, err := m.client(ctx, id.Project)
		if err != nil {
			return nil, err
		}
		return &schemaAdapter{
			gcpClient: gcpClient,
			id:        id,
		}, nil
	}
	return nil, nil
}

type schemaAdapter struct {
	gcpClient *gcp.SchemaClient
	id        *krm.DiscoveryEngineSchemaIdentity
	desired   *pb.Schema
	actual    *pb.Schema
}

var _ directbase.Adapter = &schemaAdapter{}

func (a *schemaAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting discoveryengine schema", "name", a.id.String())

	req := &pb.GetSchemaRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetSchema(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting discoveryengine schema %q from gcp: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *schemaAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating discoveryengine schema", "name", a.id.String())

	desired := proto.Clone(a.desired).(*pb.Schema)
	desired.Name = a.id.String()

	req := &pb.CreateSchemaRequest{
		Parent:   a.id.ParentString(),
		Schema:   desired,
		SchemaId: a.id.Schema,
	}
	op, err := a.gcpClient.CreateSchema(ctx, req)
	if err != nil {
		return fmt.Errorf("creating discoveryengine schema %s: %w", a.id.String(), err)
	}
	_, err = op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for discoveryengine schema creation %s: %w", a.id.String(), err)
	}

	latest, err := a.gcpClient.GetSchema(ctx, &pb.GetSchemaRequest{Name: a.id.String()})
	if err != nil {
		return fmt.Errorf("getting discoveryengine schema %s after creation: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully created discoveryengine schema in gcp", "name", a.id.String())

	return a.updateStatus(ctx, createOp, latest)
}

func (a *schemaAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating discoveryengine schema", "name", a.id.String())

	desired := proto.Clone(a.desired).(*pb.Schema)
	desired.Name = a.id.String()

	diffs, _, err := a.compare(ctx, a.actual, desired)
	if err != nil {
		return err
	}

	if !diffs.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id.String())
		return a.updateStatus(ctx, updateOp, a.actual)
	}

	diffs.Object = updateOp.GetUnstructured()
	structuredreporting.ReportDiff(ctx, diffs)

	req := &pb.UpdateSchemaRequest{
		Schema: desired,
	}
	op, err := a.gcpClient.UpdateSchema(ctx, req)
	if err != nil {
		return fmt.Errorf("updating discoveryengine schema %s: %w", a.id.String(), err)
	}
	_, err = op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for discoveryengine schema update %s: %w", a.id.String(), err)
	}

	latest, err := a.gcpClient.GetSchema(ctx, &pb.GetSchemaRequest{Name: a.id.String()})
	if err != nil {
		return fmt.Errorf("getting discoveryengine schema %s after update: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully updated discoveryengine schema", "name", a.id.String())

	return a.updateStatus(ctx, updateOp, latest)
}

func (a *schemaAdapter) compare(ctx context.Context, actual, desired *pb.Schema) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, DiscoveryEngineSchemaSpec_v1alpha1_FromProto, DiscoveryEngineSchemaSpec_v1alpha1_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = desired.Name // Restore identifier field

	clonedDesired := proto.Clone(desired).(*pb.Schema)

	if actual.GetJsonSchema() != "" && desired.GetJsonSchema() != "" {
		var actualObj, desiredObj any
		if err := json.Unmarshal([]byte(actual.GetJsonSchema()), &actualObj); err == nil {
			if err := json.Unmarshal([]byte(desired.GetJsonSchema()), &desiredObj); err == nil {
				if reflect.DeepEqual(actualObj, desiredObj) {
					maskedActual.Schema = clonedDesired.Schema
				}
			}
		}
	}

	diffs, updateMask, err := common.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}

func (a *schemaAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.Schema) error {
	status := &krm.DiscoveryEngineSchemaStatus{}
	status.ExternalRef = direct.PtrTo(a.id.String())
	return op.UpdateStatus(ctx, status, nil)
}

func (a *schemaAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}

	obj := &krm.DiscoveryEngineSchema{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(DiscoveryEngineSchemaSpec_v1alpha1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.DataStoreRef = &krm.DiscoveryEngineDataStoreRef{External: a.id.ParentString()}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u := &unstructured.Unstructured{Object: uObj}
	u.SetName(a.id.Schema)
	u.SetGroupVersionKind(krm.DiscoveryEngineSchemaGVK)

	return u, nil
}

func (a *schemaAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting discoveryengine schema", "name", a.id.String())

	req := &pb.DeleteSchemaRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteSchema(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting discoveryengine schema %s: %w", a.id.String(), err)
	}
	if err := op.Wait(ctx); err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		if err.Error() != "unsupported result type <nil>: <nil>" {
			return false, fmt.Errorf("waiting for discoveryengine schema deletion %s: %w", a.id.String(), err)
		}
	}
	log.V(2).Info("successfully deleted discoveryengine schema", "name", a.id.String())
	return true, nil
}
