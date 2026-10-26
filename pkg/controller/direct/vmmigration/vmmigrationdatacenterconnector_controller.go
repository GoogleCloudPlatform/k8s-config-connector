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
// krm.group: vmmigration.cnrm.cloud.google.com
// krm.version: v1alpha1
// krm.kind: VMMigrationDatacenterConnector

package vmmigration

import (
	"context"
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vmmigration/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/vmmigration/apiv1"
	vmmigrationpb "cloud.google.com/go/vmmigration/apiv1/vmmigrationpb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.VMMigrationDatacenterConnectorGVK, NewDatacenterConnectorModel)
}

func NewDatacenterConnectorModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelDatacenterConnector{config: *config}, nil
}

var _ directbase.Model = &modelDatacenterConnector{}

type modelDatacenterConnector struct {
	config config.ControllerConfig
}

func (m *modelDatacenterConnector) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.VMMigrationDatacenterConnector{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.VMMigrationDatacenterConnectorIdentity)

	// Get vmmigration GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &DatacenterConnectorAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelDatacenterConnector) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//vmmigration.googleapis.com/") {
		return nil, nil
	}
	id := &krm.VMMigrationDatacenterConnectorIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &DatacenterConnectorAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}
func (m *modelDatacenterConnector) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building DatacenterConnector client: %w", err)
	}
	return gcpClient, err
}

type DatacenterConnectorAdapter struct {
	id        *krm.VMMigrationDatacenterConnectorIdentity
	gcpClient *gcp.Client
	desired   *krm.VMMigrationDatacenterConnector
	actual    *vmmigrationpb.DatacenterConnector
}

var _ directbase.Adapter = &DatacenterConnectorAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *DatacenterConnectorAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting DatacenterConnector", "name", a.id)

	req := &vmmigrationpb.GetDatacenterConnectorRequest{
		Name: a.id.String(),
	}
	datacenterconnectorpb, err := a.gcpClient.GetDatacenterConnector(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting DatacenterConnector %q: %w", a.id, err)
	}

	a.actual = datacenterconnectorpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *DatacenterConnectorAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating DatacenterConnector", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := VMMigrationDatacenterConnectorSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &vmmigrationpb.CreateDatacenterConnectorRequest{
		Parent:                a.id.ParentString(),
		DatacenterConnectorId: a.id.DatacenterConnector,
		DatacenterConnector:   resource,
	}
	op, err := a.gcpClient.CreateDatacenterConnector(ctx, req)
	if err != nil {
		return fmt.Errorf("creating DatacenterConnector %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("DatacenterConnector %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created DatacenterConnector", "name", a.id)

	status := &krm.VMMigrationDatacenterConnectorStatus{}
	status.ObservedState = VMMigrationDatacenterConnectorObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *DatacenterConnectorAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating DatacenterConnector", "name", a.id)

	_, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		VMMigrationDatacenterConnectorSpec_FromProto,
		VMMigrationDatacenterConnectorSpec_ToProto,
		nil,
	)
	if err != nil {
		return err
	}
	if report.HasDiff() {
		report.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, report)
		return fmt.Errorf("updating VMMigrationDatacenterConnector is not supported, fields modified: %v", updateMask.Paths)
	}

	status := &krm.VMMigrationDatacenterConnectorStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = VMMigrationDatacenterConnectorObservedState_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *DatacenterConnectorAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.VMMigrationDatacenterConnector{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(VMMigrationDatacenterConnectorSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.VMMigrationDatacenterConnectorGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *DatacenterConnectorAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting DatacenterConnector", "name", a.id)

	req := &vmmigrationpb.DeleteDatacenterConnectorRequest{
		Name: a.id.String(),
	}
	op, err := a.gcpClient.DeleteDatacenterConnector(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent DatacenterConnector, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting DatacenterConnector %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted DatacenterConnector", "name", a.id)
	err = op.Wait(ctx)
	if err != nil {
		if direct.IsNotFound(err) || strings.Contains(err.Error(), "missing \"value\" field") {
			return true, nil
		}
		return false, fmt.Errorf("waiting delete DatacenterConnector %s: %w", a.id, err)
	}
	return true, nil
}
