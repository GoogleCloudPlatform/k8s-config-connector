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
// krm.kind: BigLakeIcebergCatalog

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

	gcp "cloud.google.com/go/biglake/apiv1"
	biglakepb "cloud.google.com/go/biglake/apiv1/biglakepb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.BigLakeIcebergCatalogGVK, NewIcebergCatalogModel)
}

func NewIcebergCatalogModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelIcebergCatalog{config: *config}, nil
}

var _ directbase.Model = &modelIcebergCatalog{}

type modelIcebergCatalog struct {
	config config.ControllerConfig
}

func (m *modelIcebergCatalog) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.BigLakeIcebergCatalog{}
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
	id := i.(*krm.BigLakeIcebergCatalogIdentity)

	// Get bigquerybiglake GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &IcebergCatalogAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelIcebergCatalog) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//biglake.googleapis.com/") {
		return nil, nil
	}
	id := &krm.BigLakeIcebergCatalogIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &IcebergCatalogAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}
func (m *modelIcebergCatalog) client(ctx context.Context) (*gcp.IcebergCatalogClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.GRPCClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewIcebergCatalogClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building IcebergCatalog client: %w", err)
	}
	return gcpClient, err
}

type IcebergCatalogAdapter struct {
	id        *krm.BigLakeIcebergCatalogIdentity
	gcpClient *gcp.IcebergCatalogClient
	desired   *krm.BigLakeIcebergCatalog
	actual    *biglakepb.IcebergCatalog
}

var _ directbase.Adapter = &IcebergCatalogAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *IcebergCatalogAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting IcebergCatalog", "name", a.id)

	req := &biglakepb.GetIcebergCatalogRequest{
		Name: a.id.String(),
	}
	icebergcatalogpb, err := a.gcpClient.GetIcebergCatalog(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting IcebergCatalog %q: %w", a.id, err)
	}

	a.actual = icebergcatalogpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *IcebergCatalogAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating IcebergCatalog", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := BigLakeIcebergCatalogSpec_v1alpha1_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &biglakepb.CreateIcebergCatalogRequest{
		Parent:           a.id.ParentString(),
		IcebergCatalogId: a.id.Catalog,
		IcebergCatalog:   resource,
	}
	created, err := a.gcpClient.CreateIcebergCatalog(ctx, req)
	if err != nil {
		return fmt.Errorf("creating IcebergCatalog %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created IcebergCatalog", "name", a.id)

	status := &krm.BigLakeIcebergCatalogStatus{}
	status.ObservedState = BigLakeIcebergCatalogObservedState_v1alpha1_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *IcebergCatalogAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating IcebergCatalog", "name", a.id)

	resource, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		BigLakeIcebergCatalogSpec_v1alpha1_FromProto,
		BigLakeIcebergCatalogSpec_v1alpha1_ToProto,
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
		req := &biglakepb.UpdateIcebergCatalogRequest{
			UpdateMask:     updateMask,
			IcebergCatalog: resource,
		}
		updated, err = a.gcpClient.UpdateIcebergCatalog(ctx, req)
		if err != nil {
			return fmt.Errorf("updating IcebergCatalog %s: %w", a.id, err)
		}
		log.V(2).Info("successfully updated IcebergCatalog", "name", a.id)
	}

	status := &krm.BigLakeIcebergCatalogStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = BigLakeIcebergCatalogObservedState_v1alpha1_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *IcebergCatalogAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.BigLakeIcebergCatalog{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(BigLakeIcebergCatalogSpec_v1alpha1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.BigLakeIcebergCatalogGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *IcebergCatalogAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting IcebergCatalog", "name", a.id)

	req := &biglakepb.DeleteIcebergCatalogRequest{
		Name: a.id.String(),
	}
	err := a.gcpClient.DeleteIcebergCatalog(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent IcebergCatalog, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting IcebergCatalog %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted IcebergCatalog", "name", a.id)
	return true, nil
}
