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

package bigquerybiglake

import (
	"context"
	"fmt"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigquerybiglake/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	gcp "cloud.google.com/go/bigquery/biglake/apiv1"
	bigquerybiglakepb "cloud.google.com/go/bigquery/biglake/apiv1/biglakepb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.BigLakeTableGVK, NewTableModel)
}

func NewTableModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelTable{config: *config}, nil
}

var _ directbase.Model = &modelTable{}

type modelTable struct {
	config config.ControllerConfig
}

func (m *modelTable) client(ctx context.Context) (*gcp.MetastoreClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewMetastoreRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building Table client: %w", err)
	}
	return gcpClient, err
}

func (m *modelTable) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.BigLakeTable{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewTableIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Get bigquerybiglake GCP client
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
	id        *krm.TableIdentity
	gcpClient *gcp.MetastoreClient
	desired   *krm.BigLakeTable
	actual    *bigquerybiglakepb.Table
}

var _ directbase.Adapter = &TableAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *TableAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting Table", "name", a.id)

	req := &bigquerybiglakepb.GetTableRequest{Name: a.id.String()}
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
	resource := BigLakeTableSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &bigquerybiglakepb.CreateTableRequest{
		Parent:  a.id.Parent().String(),
		Table:   resource,
		TableId: a.id.ID(),
	}
	created, err := a.gcpClient.CreateTable(ctx, req)
	if err != nil {
		return fmt.Errorf("creating Table %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created Table", "name", a.id)

	status := &krm.BigLakeTableStatus{}
	status.ObservedState = BigLakeTableObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *TableAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating Table", "name", a.id)
	mapCtx := &direct.MapContext{}

	desiredPb := BigLakeTableSpec_ToProto(mapCtx, &a.desired.DeepCopy().Spec)
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

	req := &bigquerybiglakepb.UpdateTableRequest{
		UpdateMask: updateMask,
		Table:      desiredPb,
	}
	updated, err := a.gcpClient.UpdateTable(ctx, req)
	if err != nil {
		return fmt.Errorf("updating Table %s: %w", a.id, err)
	}

	log.V(2).Info("successfully updated Table", "name", a.id)

	status := &krm.BigLakeTableStatus{}
	status.ObservedState = BigLakeTableObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *TableAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.BigLakeTable{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(BigLakeTableSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Parent().ProjectID}
	obj.Spec.Location = direct.LazyPtr(a.id.Parent().Location)
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.ID())
	u.SetGroupVersionKind(krm.BigLakeTableGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *TableAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting Table", "name", a.id)

	req := &bigquerybiglakepb.DeleteTableRequest{Name: a.id.String()}
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
