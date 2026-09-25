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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	admin "cloud.google.com/go/bigtable/admin/apiv2"
	pb "cloud.google.com/go/bigtable/admin/apiv2/adminpb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.BigtableSchemaBundleGVK, NewSchemaBundleModel)
}

func NewSchemaBundleModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelSchemaBundle{config: *config}, nil
}

var _ directbase.Model = &modelSchemaBundle{}

type modelSchemaBundle struct {
	config config.ControllerConfig
}

func (m *modelSchemaBundle) client(ctx context.Context) (*admin.BigtableTableAdminClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.GRPCClientOptions()
	if err != nil {
		return nil, fmt.Errorf("building BigtableSchemaBundle client options: %w", err)
	}
	gcpClient, err := admin.NewBigtableTableAdminClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building BigtableSchemaBundle client: %w", err)
	}
	return gcpClient, nil
}

func (m *modelSchemaBundle) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.BigtableSchemaBundle{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	// NormalizeReferences
	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	id, err := krm.NewBigtableSchemaBundleIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &SchemaBundleAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
		reader:    reader,
	}, nil
}

func (m *modelSchemaBundle) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type SchemaBundleAdapter struct {
	id        *krm.BigtableSchemaBundleIdentity
	gcpClient *admin.BigtableTableAdminClient
	desired   *krm.BigtableSchemaBundle
	reader    client.Reader
	actual    *pb.SchemaBundle
}

var _ directbase.Adapter = &SchemaBundleAdapter{}

func (a *SchemaBundleAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	return nil, nil
}

func (a *SchemaBundleAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting BigtableSchemaBundle", "name", a.id)

	req := &pb.GetSchemaBundleRequest{
		Name: a.id.String(),
	}

	resp, err := a.gcpClient.GetSchemaBundle(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting BigtableSchemaBundle %q: %w", a.id, err)
	}

	a.actual = resp
	return true, nil
}

func (a *SchemaBundleAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating SchemaBundle", "name", a.id)

	mapCtx := &direct.MapContext{}
	schemaBundle := BigtableSchemaBundleSpec_v1alpha1_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	schemaBundle.Name = a.id.String()

	req := &pb.CreateSchemaBundleRequest{
		Parent:         a.id.ParentString(),
		SchemaBundleId: a.id.ID(),
		SchemaBundle:   schemaBundle,
	}

	op, err := a.gcpClient.CreateSchemaBundle(ctx, req)
	if err != nil {
		return fmt.Errorf("creating SchemaBundle %s: %w", a.id, err)
	}

	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for SchemaBundle creation %s: %w", a.id, err)
	}

	log.V(2).Info("successfully created SchemaBundle", "name", a.id)

	latest, err := a.gcpClient.GetSchemaBundle(ctx, &pb.GetSchemaBundleRequest{Name: a.id.String()})
	if err != nil {
		latest = created
	}

	return a.updateStatus(ctx, createOp, latest)
}

func (a *SchemaBundleAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating SchemaBundle", "name", a.id)

	mapCtx := &direct.MapContext{}
	desired := BigtableSchemaBundleSpec_v1alpha1_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	desired.Name = a.id.String()

	diffs, updateMask, err := a.compare(ctx, a.actual, desired)
	if err != nil {
		return err
	}

	if !diffs.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id)
		return a.updateStatus(ctx, updateOp, a.actual)
	}

	log.V(2).Info("fields need update", "name", a.id, "paths", updateMask.Paths)
	diffs.Object = updateOp.GetUnstructured()
	structuredreporting.ReportDiff(ctx, diffs)

	req := &pb.UpdateSchemaBundleRequest{
		SchemaBundle: desired,
		UpdateMask:   updateMask,
	}

	op, err := a.gcpClient.UpdateSchemaBundle(ctx, req)
	if err != nil {
		return fmt.Errorf("updating SchemaBundle %s: %w", a.id, err)
	}

	updated, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for SchemaBundle update %s: %w", a.id, err)
	}

	log.V(2).Info("successfully updated SchemaBundle", "name", a.id)

	latest, err := a.gcpClient.GetSchemaBundle(ctx, &pb.GetSchemaBundleRequest{Name: a.id.String()})
	if err != nil {
		latest = updated
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func (a *SchemaBundleAdapter) compare(ctx context.Context, actual *pb.SchemaBundle, desired *pb.SchemaBundle) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	mapCtx := &direct.MapContext{}
	maskedActual := BigtableSchemaBundleSpec_v1alpha1_ToProto(mapCtx, BigtableSchemaBundleSpec_v1alpha1_FromProto(mapCtx, actual))
	if mapCtx.Err() != nil {
		return nil, nil, mapCtx.Err()
	}
	maskedActual.Name = desired.Name

	clonedDesired := proto.Clone(desired).(*pb.SchemaBundle)

	diffs, updateMask, err := common.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}

func (a *SchemaBundleAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting SchemaBundle", "name", a.id)

	req := &pb.DeleteSchemaBundleRequest{
		Name: a.id.String(),
	}

	err := a.gcpClient.DeleteSchemaBundle(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting SchemaBundle %s: %w", a.id, err)
	}

	log.V(2).Info("successfully deleted SchemaBundle", "name", a.id)
	return true, nil
}

func (a *SchemaBundleAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.SchemaBundle) error {
	status := &krm.BigtableSchemaBundleStatus{}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return op.UpdateStatus(ctx, status, nil)
}
