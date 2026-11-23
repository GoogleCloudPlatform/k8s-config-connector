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
// krm.kind: BigLakeHiveCatalog

package bigquerybiglake

import (
	"context"
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigquerybiglake/v1alpha1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
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
	registry.RegisterModel(krm.BigLakeHiveCatalogGVK, NewHiveCatalogModel)
}

func NewHiveCatalogModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelHiveCatalog{config: *config}, nil
}

var _ directbase.Model = &modelHiveCatalog{}

type modelHiveCatalog struct {
	config config.ControllerConfig
}

func (m *modelHiveCatalog) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.BigLakeHiveCatalog{}
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
	id := i.(*krm.BigLakeHiveCatalogIdentity)

	// Get bigquerybiglake GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &HiveCatalogAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelHiveCatalog) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//biglake.googleapis.com/") {
		return nil, nil
	}
	id := &krm.BigLakeHiveCatalogIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &HiveCatalogAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}
func (m *modelHiveCatalog) client(ctx context.Context) (*gcp.HiveMetastoreClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.GRPCClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewHiveMetastoreClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building HiveCatalog client: %w", err)
	}
	return gcpClient, err
}

type HiveCatalogAdapter struct {
	id        *krm.BigLakeHiveCatalogIdentity
	gcpClient *gcp.HiveMetastoreClient
	desired   *krm.BigLakeHiveCatalog
	actual    *hivepb.HiveCatalog
}

var _ directbase.Adapter = &HiveCatalogAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *HiveCatalogAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting HiveCatalog", "name", a.id)

	req := &hivepb.GetHiveCatalogRequest{
		Name: a.id.String(),
	}
	hivecatalogpb, err := a.gcpClient.GetHiveCatalog(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting HiveCatalog %q: %w", a.id, err)
	}

	a.actual = hivecatalogpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *HiveCatalogAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating HiveCatalog", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := BigLakeHiveCatalogSpec_v1alpha1_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	primaryLocation := direct.ValueOf(desired.Spec.PrimaryLocation)
	if primaryLocation == "" {
		primaryLocation = "us-central1"
	}

	req := &hivepb.CreateHiveCatalogRequest{
		Parent:          a.id.ParentString(),
		HiveCatalogId:   a.id.Catalog,
		HiveCatalog:     resource,
		PrimaryLocation: primaryLocation,
	}
	created, err := a.gcpClient.CreateHiveCatalog(ctx, req)
	if err != nil {
		return fmt.Errorf("creating HiveCatalog %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created HiveCatalog", "name", a.id)

	status := &krm.BigLakeHiveCatalogStatus{}
	status.ObservedState = BigLakeHiveCatalogObservedState_v1alpha1_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *HiveCatalogAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating HiveCatalog", "name", a.id)

	resource, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		BigLakeHiveCatalogSpec_v1alpha1_FromProto,
		BigLakeHiveCatalogSpec_v1alpha1_ToProto,
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
		req := &hivepb.UpdateHiveCatalogRequest{
			UpdateMask:  updateMask,
			HiveCatalog: resource,
		}
		updated, err = a.gcpClient.UpdateHiveCatalog(ctx, req)
		if err != nil {
			return fmt.Errorf("updating HiveCatalog %s: %w", a.id, err)
		}
		log.V(2).Info("successfully updated HiveCatalog", "name", a.id)
	}

	status := &krm.BigLakeHiveCatalogStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = BigLakeHiveCatalogObservedState_v1alpha1_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *HiveCatalogAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.BigLakeHiveCatalog{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(BigLakeHiveCatalogSpec_v1alpha1_FromProto(mapCtx, a.actual))
	obj.Spec.ProjectRef = &refsv1beta1.ProjectRef{External: a.id.Project}
	for _, replica := range a.actual.GetReplicas() {
		if replica.GetState() == hivepb.HiveCatalog_Replica_STATE_PRIMARY {
			obj.Spec.PrimaryLocation = direct.LazyPtr(replica.GetRegion())
			break
		}
	}
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.Catalog)
	u.SetGroupVersionKind(krm.BigLakeHiveCatalogGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *HiveCatalogAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting HiveCatalog", "name", a.id)

	req := &hivepb.DeleteHiveCatalogRequest{
		Name: a.id.String(),
	}
	err := a.gcpClient.DeleteHiveCatalog(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent HiveCatalog, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting HiveCatalog %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted HiveCatalog", "name", a.id)
	return true, nil
}
