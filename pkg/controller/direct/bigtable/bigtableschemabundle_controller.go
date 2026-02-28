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

package bigtable

import (
	"context"
	"fmt"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigtable/v1alpha1"
	bigtablev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigtable/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	gcp "cloud.google.com/go/bigtable"
	bigtablepb "cloud.google.com/go/bigtable/admin/apiv2/adminpb"
	"google.golang.org/api/option"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.BigtableSchemaBundleGVK, NewBigtableSchemaBundleModel)
}

func NewBigtableSchemaBundleModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelBigtableSchemaBundle{config: *config}, nil
}

var _ directbase.Model = &modelBigtableSchemaBundle{}

type modelBigtableSchemaBundle struct {
	config config.ControllerConfig
}

func (m *modelBigtableSchemaBundle) client(ctx context.Context, projectID, instanceID string) (*gcp.AdminClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.GRPCClientOptions()
	if err != nil {
		return nil, fmt.Errorf("building BigtableSchemaBundle client options: %w", err)
	}
	gcpClient, err := gcp.NewAdminClient(ctx, projectID, instanceID, opts...)
	if err != nil {
		return nil, fmt.Errorf("building BigtableSchemaBundle client: %w", err)
	}
	return gcpClient, err
}

func (m *modelBigtableSchemaBundle) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.BigtableSchemaBundle{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewSchemaBundleIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	gcpClient, err := m.client(ctx, id.Parent().Parent.Parent.ProjectID, id.Parent().Parent.Id)
	if err != nil {
		return nil, err
	}
	return &BigtableSchemaBundleAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelBigtableSchemaBundle) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type BigtableSchemaBundleAdapter struct {
	id        *krm.SchemaBundleIdentity
	gcpClient *gcp.AdminClient
	desired   *krm.BigtableSchemaBundle
	actual    *gcp.SchemaBundleInfo
}

var _ directbase.Adapter = &BigtableSchemaBundleAdapter{}

func (a *BigtableSchemaBundleAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting BigtableSchemaBundle", "name", a.id)

	bigtableschemabundleinfo, err := a.gcpClient.GetSchemaBundle(ctx, a.id.Parent().Id, a.id.ID())
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting BigtableSchemaBundle %q: %w", a.id, err)
	}

	a.actual = bigtableschemabundleinfo
	return true, nil
}

func (a *BigtableSchemaBundleAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating BigtableSchemaBundle", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := BigtableSchemaBundleSpec_v1alpha1_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	conf := &gcp.SchemaBundleConf{
		TableID:        a.id.Parent().Id,
		SchemaBundleID: a.id.ID(),
		ProtoSchema: &gcp.ProtoSchemaInfo{
			ProtoDescriptors: resource.GetProtoSchema().GetProtoDescriptors(),
		},
	}

	err := a.gcpClient.CreateSchemaBundle(ctx, conf)
	if err != nil {
		return fmt.Errorf("creating BigtableSchemaBundle %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created BigtableSchemaBundle", "name", a.id)

	// Get the created resource to get the etag/status
	created, err := a.gcpClient.GetSchemaBundle(ctx, a.id.Parent().Id, a.id.ID())
	if err != nil {
		return fmt.Errorf("getting created BigtableSchemaBundle %s: %w", a.id, err)
	}

	status := &krm.BigtableSchemaBundleStatus{}
	// Map SchemaBundleInfo back to pb.SchemaBundle for the generated mapper
	pbCreated := &bigtablepb.SchemaBundle{
		Name: a.id.String(),
		Etag: created.Etag,
		Type: &bigtablepb.SchemaBundle_ProtoSchema{
			ProtoSchema: &bigtablepb.ProtoSchema{
				ProtoDescriptors: created.SchemaBundle,
			},
		},
	}
	status.ObservedState = BigtableSchemaBundleObservedState_v1alpha1_FromProto(mapCtx, pbCreated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	if err := createOp.UpdateStatus(ctx, status, nil); err != nil {
		return err
	}

	// Write resourceID into spec.
	if err := unstructured.SetNestedField(createOp.GetUnstructured().Object, a.id.ID(), "spec", "resourceID"); err != nil {
		return fmt.Errorf("error setting spec.resourceID: %w", err)
	}

	return nil
}

func (a *BigtableSchemaBundleAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating BigtableSchemaBundle", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := BigtableSchemaBundleSpec_v1alpha1_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	conf := gcp.UpdateSchemaBundleConf{
		SchemaBundleConf: gcp.SchemaBundleConf{
			TableID:        a.id.Parent().Id,
			SchemaBundleID: a.id.ID(),
			ProtoSchema: &gcp.ProtoSchemaInfo{
				ProtoDescriptors: resource.GetProtoSchema().GetProtoDescriptors(),
			},
			Etag: a.actual.Etag,
		},
	}

	err := a.gcpClient.UpdateSchemaBundle(ctx, conf)
	if err != nil {
		return fmt.Errorf("updating BigtableSchemaBundle %s: %w", a.id, err)
	}
	log.V(2).Info("successfully updated BigtableSchemaBundle", "name", a.id)

	// Get the updated resource
	updated, err := a.gcpClient.GetSchemaBundle(ctx, a.id.Parent().Id, a.id.ID())
	if err != nil {
		return fmt.Errorf("getting updated BigtableSchemaBundle %s: %w", a.id, err)
	}

	status := &krm.BigtableSchemaBundleStatus{}
	pbUpdated := &bigtablepb.SchemaBundle{
		Name: a.id.String(),
		Etag: updated.Etag,
		Type: &bigtablepb.SchemaBundle_ProtoSchema{
			ProtoSchema: &bigtablepb.ProtoSchema{
				ProtoDescriptors: updated.SchemaBundle,
			},
		},
	}
	status.ObservedState = BigtableSchemaBundleObservedState_v1alpha1_FromProto(mapCtx, pbUpdated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *BigtableSchemaBundleAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.BigtableSchemaBundle{}
	mapCtx := &direct.MapContext{}

	// Map SchemaBundleInfo back to pb.SchemaBundle for the generated mapper
	pbActual := &bigtablepb.SchemaBundle{
		Name: a.id.String(),
		Etag: a.actual.Etag,
		Type: &bigtablepb.SchemaBundle_ProtoSchema{
			ProtoSchema: &bigtablepb.ProtoSchema{
				ProtoDescriptors: a.actual.SchemaBundle,
			},
		},
	}

	obj.Spec = direct.ValueOf(BigtableSchemaBundleSpec_v1alpha1_FromProto(mapCtx, pbActual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	// parent tableRef
	tableID, resourceID, err := krm.ParseSchemaBundleExternal(a.id.String())
	if err != nil {
		return nil, err
	}
	obj.Spec.TableRef = bigtablev1beta1.TableRef{External: tableID.String()}
	obj.Spec.ResourceID = direct.LazyPtr(resourceID)

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.ID())
	u.SetGroupVersionKind(krm.BigtableSchemaBundleGVK)

	u.Object = uObj
	return u, nil
}

func (a *BigtableSchemaBundleAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting BigtableSchemaBundle", "name", a.id)

	err := a.gcpClient.DeleteSchemaBundle(ctx, a.id.Parent().Id, a.id.ID())
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent BigtableSchemaBundle, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting BigtableSchemaBundle %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted BigtableSchemaBundle", "name", a.id)

	return true, nil
}
