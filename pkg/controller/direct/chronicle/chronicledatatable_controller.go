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
// krm.group: chronicle.cnrm.cloud.google.com
// krm.version: v1alpha1
// krm.kind: ChronicleDataTable

package chronicle

import (
	"context"
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/chronicle/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/chronicle/apiv1"
	chroniclepb "cloud.google.com/go/chronicle/apiv1/chroniclepb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.ChronicleDataTableGVK, NewDataTableModel)
}

func NewDataTableModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelDataTable{config: *config}, nil
}

var _ directbase.Model = &modelDataTable{}

type modelDataTable struct {
	config config.ControllerConfig
}

func (m *modelDataTable) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ChronicleDataTable{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.ChronicleDataTableIdentity)

	// Get chronicle GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &DataTableAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelDataTable) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//chronicle.googleapis.com/") {
		return nil, nil
	}
	id := &krm.ChronicleDataTableIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &DataTableAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}
func (m *modelDataTable) client(ctx context.Context) (*gcp.DataTableClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewDataTableRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building DataTable client: %w", err)
	}
	return gcpClient, err
}

type DataTableAdapter struct {
	id        *krm.ChronicleDataTableIdentity
	gcpClient *gcp.DataTableClient
	desired   *krm.ChronicleDataTable
	actual    *chroniclepb.DataTable
}

var _ directbase.Adapter = &DataTableAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *DataTableAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting DataTable", "name", a.id)

	req := &chroniclepb.GetDataTableRequest{
		Name: a.id.String(),
	}
	datatablepb, err := a.gcpClient.GetDataTable(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting DataTable %q: %w", a.id, err)
	}

	a.actual = datatablepb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *DataTableAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating DataTable", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := ChronicleDataTableSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &chroniclepb.CreateDataTableRequest{
		Parent:      a.id.ParentString(),
		DataTableId: a.id.DataTable,
		DataTable:   resource,
	}
	created, err := a.gcpClient.CreateDataTable(ctx, req)
	if err != nil {
		return fmt.Errorf("creating DataTable %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created DataTable", "name", a.id)

	status := &krm.ChronicleDataTableStatus{}
	status.ObservedState = ChronicleDataTableObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *DataTableAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating DataTable", "name", a.id)

	resource, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		ChronicleDataTableSpec_FromProto,
		ChronicleDataTableSpec_ToProto,
		nil,
	)
	if err != nil {
		return err
	}
	updated := a.actual
	if !report.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id)
	} else {
		report.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, report)
		resource.Name = a.id.String()
		req := &chroniclepb.UpdateDataTableRequest{
			UpdateMask: updateMask,
			DataTable:  resource,
		}
		updated, err = a.gcpClient.UpdateDataTable(ctx, req)
		if err != nil {
			return fmt.Errorf("updating DataTable %s: %w", a.id, err)
		}
		log.V(2).Info("successfully updated DataTable", "name", a.id)
	}

	status := &krm.ChronicleDataTableStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = ChronicleDataTableObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *DataTableAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ChronicleDataTable{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(ChronicleDataTableSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.ChronicleDataTableGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *DataTableAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting DataTable", "name", a.id)

	req := &chroniclepb.DeleteDataTableRequest{
		Name: a.id.String(),
	}
	err := a.gcpClient.DeleteDataTable(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent DataTable, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting DataTable %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted DataTable", "name", a.id)
	return true, nil
}
