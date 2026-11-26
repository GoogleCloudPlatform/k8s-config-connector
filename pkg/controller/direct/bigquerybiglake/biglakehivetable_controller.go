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
// krm.group: bigquerybiglake.cnrm.cloud.google.com
// krm.version: v1alpha1
// krm.kind: BigLakeHiveTable

package bigquerybiglake

import (
	"context"
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigquerybiglake/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/biglake/hive/apiv1"
	hivepb "cloud.google.com/go/biglake/hive/apiv1/hivepb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.BigLakeHiveTableGVK, NewHiveTableModel)
}

func NewHiveTableModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelHiveTable{config: *config}, nil
}

var _ directbase.Model = &modelHiveTable{}

type modelHiveTable struct {
	config config.ControllerConfig
}

func (m *modelHiveTable) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.BigLakeHiveTable{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if err := common.NormalizeGreenfieldReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.BigLakeHiveTableIdentity)

	// Get bigquerybiglake GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &HiveTableAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelHiveTable) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//biglake.googleapis.com/") {
		return nil, nil
	}
	id := &krm.BigLakeHiveTableIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &HiveTableAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}
func (m *modelHiveTable) client(ctx context.Context) (*gcp.HiveMetastoreClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewHiveMetastoreRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building HiveTable client: %w", err)
	}
	return gcpClient, err
}

type HiveTableAdapter struct {
	id        *krm.BigLakeHiveTableIdentity
	gcpClient *gcp.HiveMetastoreClient
	desired   *krm.BigLakeHiveTable
	actual    *hivepb.HiveTable
}

var _ directbase.Adapter = &HiveTableAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *HiveTableAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting HiveTable", "name", a.id)

	req := &hivepb.GetHiveTableRequest{
		Name: a.id.String(),
	}
	hivetablepb, err := a.gcpClient.GetHiveTable(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting HiveTable %q: %w", a.id, err)
	}

	a.actual = hivetablepb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *HiveTableAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating HiveTable", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := BigLakeHiveTableSpec_v1alpha1_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &hivepb.CreateHiveTableRequest{
		Parent:      a.id.ParentString(),
		HiveTableId: a.id.Table,
		HiveTable:   resource,
	}
	created, err := a.gcpClient.CreateHiveTable(ctx, req)
	if err != nil {
		return fmt.Errorf("creating HiveTable %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created HiveTable", "name", a.id)

	status := &krm.BigLakeHiveTableStatus{}
	status.ObservedState = BigLakeHiveTableObservedState_v1alpha1_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *HiveTableAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating HiveTable", "name", a.id)

	normalize := func(ctx context.Context, p *hivepb.HiveTable) error {
		if p != nil && p.Parameters != nil {
			delete(p.Parameters, "table_type")
		}
		if p != nil && p.StorageDescriptor != nil {
			// GCP defaults compressed to false and omits it from the JSON/proto response when false.
			if p.StorageDescriptor.Compressed != nil && !*p.StorageDescriptor.Compressed {
				p.StorageDescriptor.Compressed = nil
			}
		}
		return nil
	}

	resource, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		BigLakeHiveTableSpec_v1alpha1_FromProto,
		BigLakeHiveTableSpec_v1alpha1_ToProto,
		normalize,
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
		req := &hivepb.UpdateHiveTableRequest{
			UpdateMask: updateMask,
			HiveTable:  resource,
		}
		updated, err = a.gcpClient.UpdateHiveTable(ctx, req)
		if err != nil {
			return fmt.Errorf("updating HiveTable %s: %w", a.id, err)
		}
		log.V(2).Info("successfully updated HiveTable", "name", a.id)
	}

	status := &krm.BigLakeHiveTableStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = BigLakeHiveTableObservedState_v1alpha1_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *HiveTableAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.BigLakeHiveTable{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(BigLakeHiveTableSpec_v1alpha1_FromProto(mapCtx, a.actual))
	obj.Spec.ParentRef = &krm.BigLakeHiveDatabaseRef{
		External: fmt.Sprintf("projects/%s/catalogs/%s/databases/%s", a.id.Project, a.id.Catalog, a.id.Database),
	}
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.Table)
	u.SetGroupVersionKind(krm.BigLakeHiveTableGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *HiveTableAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting HiveTable", "name", a.id)

	req := &hivepb.DeleteHiveTableRequest{
		Name: a.id.String(),
	}
	err := a.gcpClient.DeleteHiveTable(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent HiveTable, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting HiveTable %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted HiveTable", "name", a.id)
	return true, nil
}
