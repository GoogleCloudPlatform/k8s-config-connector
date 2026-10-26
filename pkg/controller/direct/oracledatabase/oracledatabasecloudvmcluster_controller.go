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
// krm.group: oracledatabase.cnrm.cloud.google.com
// krm.version: v1alpha1
// krm.kind: OracleDatabaseCloudVMCluster

package oracledatabase

import (
	"context"
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/oracledatabase/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/oracledatabase/apiv1"
	oracledatabasepb "cloud.google.com/go/oracledatabase/apiv1/oracledatabasepb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.OracleDatabaseCloudVMClusterGVK, NewCloudVmClusterModel)
}

func NewCloudVmClusterModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelCloudVmCluster{config: *config}, nil
}

var _ directbase.Model = &modelCloudVmCluster{}

type modelCloudVmCluster struct {
	config config.ControllerConfig
}

func (m *modelCloudVmCluster) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.OracleDatabaseCloudVMCluster{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.OracleDatabaseCloudVMClusterIdentity)

	// Get oracledatabase GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &CloudVmClusterAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelCloudVmCluster) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//oracledatabase.googleapis.com/") {
		return nil, nil
	}
	id := &krm.OracleDatabaseCloudVMClusterIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &CloudVmClusterAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}
func (m *modelCloudVmCluster) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building CloudVmCluster client: %w", err)
	}
	return gcpClient, err
}

type CloudVmClusterAdapter struct {
	id        *krm.OracleDatabaseCloudVMClusterIdentity
	gcpClient *gcp.Client
	desired   *krm.OracleDatabaseCloudVMCluster
	actual    *oracledatabasepb.CloudVmCluster
}

var _ directbase.Adapter = &CloudVmClusterAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *CloudVmClusterAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting CloudVmCluster", "name", a.id)

	req := &oracledatabasepb.GetCloudVmClusterRequest{
		Name: a.id.String(),
	}
	cloudvmclusterpb, err := a.gcpClient.GetCloudVmCluster(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting CloudVmCluster %q: %w", a.id, err)
	}

	a.actual = cloudvmclusterpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *CloudVmClusterAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating CloudVmCluster", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := OracleDatabaseCloudVMClusterSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &oracledatabasepb.CreateCloudVmClusterRequest{
		Parent:           a.id.ParentString(),
		CloudVmClusterId: a.id.CloudVmCluster,
		CloudVmCluster:   resource,
	}
	op, err := a.gcpClient.CreateCloudVmCluster(ctx, req)
	if err != nil {
		return fmt.Errorf("creating CloudVmCluster %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("CloudVmCluster %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created CloudVmCluster", "name", a.id)

	status := &krm.OracleDatabaseCloudVMClusterStatus{}
	status.ObservedState = OracleDatabaseCloudVMClusterObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *CloudVmClusterAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating CloudVmCluster", "name", a.id)

	_, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		OracleDatabaseCloudVMClusterSpec_FromProto,
		OracleDatabaseCloudVMClusterSpec_ToProto,
		nil,
	)
	if err != nil {
		return err
	}
	if report.HasDiff() {
		report.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, report)
		return fmt.Errorf("updating OracleDatabaseCloudVMCluster is not supported, fields modified: %v", updateMask.Paths)
	}

	status := &krm.OracleDatabaseCloudVMClusterStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = OracleDatabaseCloudVMClusterObservedState_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *CloudVmClusterAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.OracleDatabaseCloudVMCluster{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(OracleDatabaseCloudVMClusterSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.OracleDatabaseCloudVMClusterGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *CloudVmClusterAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting CloudVmCluster", "name", a.id)

	req := &oracledatabasepb.DeleteCloudVmClusterRequest{
		Name: a.id.String(),
	}
	op, err := a.gcpClient.DeleteCloudVmCluster(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent CloudVmCluster, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting CloudVmCluster %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted CloudVmCluster", "name", a.id)
	err = op.Wait(ctx)
	if err != nil {
		if direct.IsNotFound(err) || strings.Contains(err.Error(), "missing \"value\" field") {
			return true, nil
		}
		return false, fmt.Errorf("waiting delete CloudVmCluster %s: %w", a.id, err)
	}
	return true, nil
}
