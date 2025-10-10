// Copyright 2025 Google LLC
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

	gcp "cloud.google.com/go/bigtable/admin/apiv2"
	bigtablepb "cloud.google.com/go/bigtable/admin/apiv2/bigtablepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigtable/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.BigtableTableGVK, NewTableModel)
}

func NewTableModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelTable{config: *config}, nil
}

var _ directbase.Model = &modelTable{}

type modelTable struct {
	config config.ControllerConfig
}

type TableIdentity struct {
	ProjectID  string
	InstanceID string
	TableID    string
}

func (i *TableIdentity) String() string {
	return fmt.Sprintf("projects/%s/instances/%s/tables/%s", i.ProjectID, i.InstanceID, i.TableID)
}

func (i *TableIdentity) Parent() string {
	return fmt.Sprintf("projects/%s/instances/%s", i.ProjectID, i.InstanceID)
}

func NewTableIdentity(ctx context.Context, reader client.Reader, u *unstructured.Unstructured, obj *krm.BigtableTable) (*TableIdentity, error) {
	projectID, err := common.GetProjectID(u, nil)
	if err != nil {
		return nil, err
	}

	instanceUnstructured, err := direct.ResolveRef(ctx, reader, obj, &obj.Spec.InstanceRef)
	if err != nil {
		return nil, fmt.Errorf("resolving instanceRef: %w", err)
	}

	instanceObj := &krm.BigtableInstance{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(instanceUnstructured.Object, instanceObj); err != nil {
		return nil, fmt.Errorf("error converting instance to %T: %w", instanceObj, err)
	}

	instanceID := instanceObj.Name
	if instanceObj.Spec.ResourceID != nil {
		instanceID = *instanceObj.Spec.ResourceID
	}

	return &TableIdentity{
		ProjectID:  projectID,
		InstanceID: instanceID,
		TableID:    obj.Name,
	}, nil
}

func (m *modelTable) client(ctx context.Context) (*gcp.AdminClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewAdminRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building Table client: %w", err)
	}
	return gcpClient, err
}

func (m *modelTable) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.BigtableTable{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := NewTableIdentity(ctx, reader, u, obj)
	if err != nil {
		return nil, err
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &TableAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelTable) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type TableAdapter struct {
	id        *TableIdentity
	gcpClient *gcp.AdminClient
	desired   *krm.BigtableTable
	actual    *bigtablepb.Table
}

var _ directbase.Adapter = &TableAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *TableAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting Table", "name", a.id)

	req := &bigtablepb.GetTableRequest{Name: a.id.String()}
	tablepb, err := a.gcpClient.GetTable(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting Table %q: %w", a.id, err)
	}

	a.actual = tablepb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *TableAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating Table", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := BigtableTableSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &bigtablepb.CreateTableRequest{
		Parent:  a.id.Parent(),
		Table:   resource,
		TableId: a.id.TableID,
	}
	created, err := a.gcpClient.CreateTable(ctx, req)
	if err != nil {
		return fmt.Errorf("creating Table %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created Table", "name", a.id)

	status := &krm.BigtableTableStatus{}
	status.ObservedState = BigtableTableObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return createOp.UpdateStatus(ctx, status)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *TableAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating Table", "name", a.id)
	mapCtx := &direct.MapContext{}

	desiredPb := BigtableTableSpec_ToProto(mapCtx, &a.desired.DeepCopy().Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	paths := make(sets.Set[string])
	var err error
	paths, err = common.CompareProtoMessage(desiredPb, a.actual, common.BasicDiff)
	if err != nil {
		return err
	}

	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		return nil
	}
	updateMask := &fieldmaskpb.FieldMask{
		Paths: sets.List(paths),
	}
	desiredPb.Name = a.id.String()

	req := &bigtablepb.UpdateTableRequest{
		UpdateMask: updateMask,
		Table:      desiredPb,
	}
	updated, err := a.gcpClient.UpdateTable(ctx, req)
	if err != nil {
		return fmt.Errorf("updating Table %s: %w", a.id, err)
	}

	log.V(2).Info("successfully updated Table", "name", a.id)

	status := &krm.BigtableTableStatus{}
	status.ObservedState = BigtableTableObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *TableAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.BigtableTable{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(BigtableTableSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.TableID)
	u.SetGroupVersionKind(krm.BigtableTableGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *TableAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting Table", "name", a.id)

	req := &bigtablepb.DeleteTableRequest{Name: a.id.String()}
	_, err := a.gcpClient.DeleteTable(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent Table, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting Table %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted Table", "name", a.id)

	return true, nil
}
