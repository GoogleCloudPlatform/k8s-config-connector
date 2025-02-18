// Copyright 2024 Google LLC
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

package spanner

import (
	"context"
	"errors"
	"fmt"
	"reflect"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/spanner/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	kccpredicate "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/predicate"

	db "cloud.google.com/go/spanner/admin/database/apiv1"
	gcp "cloud.google.com/go/spanner/admin/instance/apiv1"

	databasepb "cloud.google.com/go/spanner/admin/database/apiv1/databasepb"
	spannerpb "cloud.google.com/go/spanner/admin/instance/apiv1/instancepb"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	ctrlName = "spanner-controller"
)

func init() {
	rg := &InstanceReconcileGate{}
	registry.RegisterModelWithReconcileGate(krm.SpannerInstanceGVK, NewSpannerInstanceModel, rg)
}

type InstanceReconcileGate struct {
	optIn kccpredicate.OptInToDirectReconciliation
}

var _ kccpredicate.ReconcileGate = &InstanceReconcileGate{}

func (r *InstanceReconcileGate) ShouldReconcile(o *unstructured.Unstructured) bool {
	return r.optIn.ShouldReconcile(o)
}

func NewSpannerInstanceModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelSpannerInstance{config: *config}, nil
}

var _ directbase.Model = &modelSpannerInstance{}

type modelSpannerInstance struct {
	config config.ControllerConfig
}

func (m *modelSpannerInstance) client(ctx context.Context) (*gcp.InstanceAdminClient, *db.DatabaseAdminClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, nil, err
	}
	gcpClient, err := gcp.NewInstanceAdminRESTClient(ctx, opts...)
	if err != nil {
		return nil, nil, fmt.Errorf("building Instance client: %w", err)
	}
	backClient, err := db.NewDatabaseAdminRESTClient(ctx, opts...)
	if err != nil {
		return nil, nil, fmt.Errorf("building Backup client: %w", err)
	}
	return gcpClient, backClient, err
}

func (m *modelSpannerInstance) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.SpannerInstance{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewSpannerInstanceIdentity(ctx, reader, obj, u)
	if err != nil {
		return nil, err
	}

	// Get spanner GCP client
	gcpClient, backClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	resourceID := direct.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}
	return &SpannerInstanceAdapter{
		id:           id,
		gcpClient:    gcpClient,
		backupClient: backClient,
		desired:      obj,
	}, nil
}

func (m *modelSpannerInstance) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type SpannerInstanceAdapter struct {
	id           *krm.SpannerInstanceIdentity
	gcpClient    *gcp.InstanceAdminClient
	backupClient *db.DatabaseAdminClient
	desired      *krm.SpannerInstance
	actual       *spannerpb.Instance
}

var _ directbase.Adapter = &SpannerInstanceAdapter{}

func (a *SpannerInstanceAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("getting SpannerInstance", "name", a.id)

	req := &spannerpb.GetInstanceRequest{Name: a.id.String()}
	instancepb, err := a.gcpClient.GetInstance(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting SpannerInstance %q: %w", a.id, err)
	}

	a.actual = instancepb
	return true, nil
}

func (a *SpannerInstanceAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("creating Instance", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	if err := a.SpecValidation(); err != nil {
		return err
	}
	resource := SpannerInstanceSpec_ToProto(mapCtx, &desired.Spec, a.id.SpannerInstanceConfigPrefix())
	// If node count or processing unit and auto-scaling config is not specify,
	// Default NodeCount to 1.
	if resource.NodeCount == 0 && resource.ProcessingUnits == 0 && resource.AutoscalingConfig == nil {
		resource.NodeCount = 1
	}
	resource.Name = a.id.String()
	resource.Labels = desired.Labels
	resource.Labels["managed-by-cnrm"] = "true"
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	req := &spannerpb.CreateInstanceRequest{
		InstanceId: a.id.ID(),
		Instance:   resource,
		Parent:     a.id.Parent().String(),
	}
	op, err := a.gcpClient.CreateInstance(ctx, req)
	if err != nil {
		return fmt.Errorf("creating Instance %s: %w", a.id, err)
	}

	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("Instance %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created Instance", "name", a.id)

	status := SpannerInstanceStatus_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	external := a.id.String()
	status.ExternalRef = &external
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *SpannerInstanceAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("updating Instance", "name", a.id)
	mapCtx := &direct.MapContext{}
	if err := a.SpecValidation(); err != nil {
		return err
	}
	desired := a.desired.DeepCopy()
	resource := SpannerInstanceSpec_ToProto(mapCtx, &desired.Spec, a.id.SpannerInstanceConfigPrefix())
	resource.Name = a.id.String()
	resource.Labels = desired.Labels
	resource.Labels["managed-by-cnrm"] = "true"
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	updateMask := &fieldmaskpb.FieldMask{}
	if !reflect.DeepEqual(a.desired.Spec.DisplayName, a.actual.DisplayName) {
		updateMask.Paths = append(updateMask.Paths, "display_name")
	}
	// If node count is unset, the field become unmanaged.
	// If autoscaling is set, this field become output-only.
	if a.desired.Spec.AutoscalingConfig == nil && a.desired.Spec.NumNodes != nil && !reflect.DeepEqual(resource.NodeCount, a.actual.NodeCount) {
		updateMask.Paths = append(updateMask.Paths, "node_count")
	}
	// If processing unit is unset, the field become unmanaged.
	// If autoscaling is set, this field become output-only.
	if a.desired.Spec.AutoscalingConfig == nil && a.desired.Spec.ProcessingUnits != nil && !reflect.DeepEqual(resource.ProcessingUnits, a.actual.ProcessingUnits) {
		updateMask.Paths = append(updateMask.Paths, "processing_units")
	}
	if !reflect.DeepEqual(resource.Labels, a.actual.Labels) {
		updateMask.Paths = append(updateMask.Paths, "labels")
	}

	if !reflect.DeepEqual(resource.AutoscalingConfig, a.actual.AutoscalingConfig) {
		updateMask.Paths = append(updateMask.Paths, "autoscaling_config")
	}

	if !reflect.DeepEqual(resource.Edition, a.actual.Edition) {
		updateMask.Paths = append(updateMask.Paths, "edition")
	}

	if len(updateMask.Paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		return nil
	}

	req := &spannerpb.UpdateInstanceRequest{
		FieldMask: updateMask,
		Instance:  resource,
	}
	op, err := a.gcpClient.UpdateInstance(ctx, req)
	if err != nil {
		return fmt.Errorf("updating Instance %s: %w", a.id, err)
	}
	updated, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("Instance %s waiting update: %w", a.id, err)
	}
	log.V(2).Info("successfully updated Instance", "name", a.id)

	status := SpannerInstanceStatus_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *SpannerInstanceAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.SpannerInstance{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(SpannerInstanceSpec_FromProto(mapCtx, a.actual, a.id.SpannerInstanceConfigPrefix()))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.actual.Name)
	u.SetGroupVersionKind(krm.SpannerInstanceGVK)

	u.Object = uObj
	return u, nil
}

// Delete implements the Adapter interface.
func (a *SpannerInstanceAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("deleting Instance", "name", a.id)
	v, ok := a.desired.ObjectMeta.Annotations["cnrm.cloud.google.com/force-destroy"]
	if ok && v == "true" {
		log.V(2).Info("Force destroy is enabled, trying to delete all backups...")
		err := a.DeleteBackup(ctx)
		if err != nil {
			return false, err
		}
	}

	req := &spannerpb.DeleteInstanceRequest{Name: a.id.String()}
	err := a.gcpClient.DeleteInstance(ctx, req)
	if err != nil {
		return false, fmt.Errorf("deleting Instance %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted Instance", "name", a.id)
	return true, nil
}

func (a *SpannerInstanceAdapter) SpecValidation() error {
	if a.desired.Spec.NumNodes != nil && a.desired.Spec.ProcessingUnits != nil {
		return fmt.Errorf("Only one field can be set between numNodes and processingUnits.")
	}
	return nil
}

func (a *SpannerInstanceAdapter) DeleteBackup(ctx context.Context) error {
	it := a.backupClient.ListBackups(ctx, &databasepb.ListBackupsRequest{Parent: a.id.String()})
	for {
		backup, err := it.Next()
		if errors.Is(err, iterator.Done) {
			break
		}
		if err != nil {
			klog.Fatalf("Error listing backups: %v", err)
			return err
		}
		err = a.backupClient.DeleteBackup(ctx, &databasepb.DeleteBackupRequest{Name: backup.Name})
		if err != nil {
			klog.Fatalf("Error deleting backup: %v", err)
			return err
		}
	}
	return nil
}
