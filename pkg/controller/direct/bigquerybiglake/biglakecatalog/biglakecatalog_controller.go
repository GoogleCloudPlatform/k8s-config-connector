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

package biglakecatalog

import (
	"context"
	"fmt"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigquerybiglake/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/bigquerybiglake"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	gcp "cloud.google.com/go/bigquery/biglake/apiv1"
	bigquerybiglakepb "cloud.google.com/go/bigquery/biglake/apiv1/biglakepb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.BigLakeCatalogGVK, NewCatalogModel)
}

func NewCatalogModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelCatalog{config: *config}, nil
}

var _ directbase.Model = &modelCatalog{}

type modelCatalog struct {
	config config.ControllerConfig
}

func (m *modelCatalog) client(ctx context.Context) (*gcp.MetastoreClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewMetastoreRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building Catalog client: %w", err)
	}
	return gcpClient, err
}

func (m *modelCatalog) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.BigLakeCatalog{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}

	copied := obj.DeepCopy()
	desired := bigquerybiglake.BigLakeCatalogSpec_v1alpha1_ToProto(mapCtx, &copied.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	// Get bigquerybiglake GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &CatalogAdapter{
		id:        id.(*krm.CatalogIdentity),
		gcpClient: gcpClient,
		desired:   desired,
	}, nil
}

func (m *modelCatalog) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type CatalogAdapter struct {
	id        *krm.CatalogIdentity
	gcpClient *gcp.MetastoreClient
	desired   *bigquerybiglakepb.Catalog
	actual    *bigquerybiglakepb.Catalog
}

var _ directbase.Adapter = &CatalogAdapter{}

func (a *CatalogAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting Catalog", "name", a.id)

	req := &bigquerybiglakepb.GetCatalogRequest{Name: a.id.String()}
	catalogpb, err := a.gcpClient.GetCatalog(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting Catalog %q: %w", a.id, err)
	}

	a.actual = catalogpb
	return true, nil
}

func (a *CatalogAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating Catalog", "name", a.id)
	req := &bigquerybiglakepb.CreateCatalogRequest{
		Parent:    a.id.Parent().String(),
		Catalog:   a.desired,
		CatalogId: a.id.ID(),
	}
	created, err := a.gcpClient.CreateCatalog(ctx, req)
	if err != nil {
		return fmt.Errorf("creating Catalog %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created Catalog", "name", a.id)

	mapCtx := &direct.MapContext{}
	status := &krm.BigLakeCatalogStatus{}
	status.ObservedState = bigquerybiglake.BigLakeCatalogObservedState_v1alpha1_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *CatalogAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating Catalog", "name", a.id)

	paths, err := common.CompareProtoMessage(a.desired, a.actual, common.BasicDiff)
	if err != nil {
		return err
	}

	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		return nil
	}

	return fmt.Errorf("Catalog is immutable and cannot be updated")
}

func (a *CatalogAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.BigLakeCatalog{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(bigquerybiglake.BigLakeCatalogSpec_v1alpha1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	externalRef := a.actual.GetName()
	id := &krm.CatalogIdentity{}
	if err := id.FromExternal(externalRef); err != nil {
		return nil, fmt.Errorf("parsing external ref %q: %w", externalRef, err)
	}
	obj.Spec.ProjectAndLocationRef = &parent.ProjectAndLocationRef{
		ProjectRef: &refs.ProjectRef{
			External: id.Parent().ProjectID,
		},
		Location: id.Parent().Location,
	}
	obj.Spec.ResourceID = direct.LazyPtr(a.id.ID())
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.ID())
	u.SetGroupVersionKind(krm.BigLakeCatalogGVK)

	u.Object = uObj
	return u, nil
}

func (a *CatalogAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting Catalog", "name", a.id)

	req := &bigquerybiglakepb.DeleteCatalogRequest{Name: a.id.String()}
	_, err := a.gcpClient.DeleteCatalog(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent Catalog, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting Catalog %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted Catalog", "name", a.id)

	return true, nil
}
