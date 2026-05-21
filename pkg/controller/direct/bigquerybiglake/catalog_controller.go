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

package bigquerybiglake

import (
	"context"
	"fmt"

	gcp "cloud.google.com/go/bigquery/biglake/apiv1"
	pb "cloud.google.com/go/bigquery/biglake/apiv1/biglakepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigquerybiglake/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"google.golang.org/api/option"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.BigLakeCatalogGVK, NewCatalogModel)
}

func NewCatalogModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelCatalog{config: *config}, nil
}

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
		return nil, fmt.Errorf("building MetastoreRESTClient: %w", err)
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

	// Get fully qualified identity
	catalogID := id.(*krm.CatalogIdentity)

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &CatalogAdapter{
		id:        catalogID,
		inner:     obj,
		gcpClient: gcpClient,
		reader:    reader,
	}, nil
}

func (m *modelCatalog) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil // not needed
}

type CatalogAdapter struct {
	id        *krm.CatalogIdentity
	inner     *krm.BigLakeCatalog
	gcpClient *gcp.MetastoreClient
	reader    client.Reader
}

var _ directbase.Adapter = &CatalogAdapter{}

func (a *CatalogAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting BigLakeCatalog", "name", a.id.String())

	req := &pb.GetCatalogRequest{
		Name: a.id.String(),
	}
	catalogpb, err := a.gcpClient.GetCatalog(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting BigLakeCatalog %q: %w", a.id.String(), err)
	}

	mapCtx := &direct.MapContext{}
	out := BigLakeCatalogSpec_v1alpha1_FromProto(mapCtx, catalogpb)
	if mapCtx.Err() != nil {
		return false, fmt.Errorf("mapping from proto to krm: %w", mapCtx.Err())
	}
	_ = out
	// update obj status
	status := &krm.BigLakeCatalogStatus{}
	status.ObservedState = BigLakeCatalogObservedState_v1alpha1_FromProto(mapCtx, catalogpb)
	if mapCtx.Err() != nil {
		return false, fmt.Errorf("mapping from proto to krm: %w", mapCtx.Err())
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	a.inner.Status = *status

	return true, nil
}

func (a *CatalogAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating BigLakeCatalog", "name", a.id.String())
	mapCtx := &direct.MapContext{}

	desired := a.inner.DeepCopy()
	resource := BigLakeCatalogSpec_v1alpha1_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &pb.CreateCatalogRequest{
		Parent:    a.id.Parent().String(),
		CatalogId: a.id.ID(),
		Catalog:   resource,
	}

	catalogpb, err := a.gcpClient.CreateCatalog(ctx, req)
	if err != nil {
		return fmt.Errorf("creating BigLakeCatalog %s: %w", a.id.String(), err)
	}

	status := &krm.BigLakeCatalogStatus{}
	status.ObservedState = BigLakeCatalogObservedState_v1alpha1_FromProto(mapCtx, catalogpb)
	if mapCtx.Err() != nil {
		return fmt.Errorf("mapping from proto to krm: %w", mapCtx.Err())
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *CatalogAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating BigLakeCatalog", "name", a.id.String())
	// BigLake Catalog currently has no updatable fields in spec.
	return nil
}

func (a *CatalogAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.inner == nil {
		return nil, fmt.Errorf("export error: inner is nil")
	}
	u := &unstructured.Unstructured{}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(a.inner)
	if err != nil {
		return nil, err
	}
	u.Object = uObj
	return u, nil
}

func (a *CatalogAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting BigLakeCatalog", "name", a.id.String())

	req := &pb.DeleteCatalogRequest{
		Name: a.id.String(),
	}

	_, err := a.gcpClient.DeleteCatalog(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting BigLakeCatalog %s: %w", a.id.String(), err)
	}

	return true, nil
}
