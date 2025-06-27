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

// +tool:controller
// proto.service: google.cloud.bigquery.biglake.v1.MetastoreService
// proto.message: google.cloud.bigquery.biglake.v1.Catalog
// crd.type: BigqueryCatalog
// crd.version: v1alpha1

package bigquerybiglake

import (
	"context"
	"fmt"

	gcp "cloud.google.com/go/bigquery/biglake/apiv1"
	pb "cloud.google.com/go/bigquery/biglake/apiv1/biglakepb"
	"google.golang.org/api/option"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigquerybiglake/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
)

func init() {
	registry.RegisterModel(krm.BigLakeCatalogGVK, NewCatalogModel)
}

func NewCatalogModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &catalogModel{config: config}, nil
}

var _ directbase.Model = &catalogModel{}

type catalogModel struct {
	config *config.ControllerConfig
}

func (m *catalogModel) client(ctx context.Context, projectID string) (*gcp.MetastoreClient, error) {
	var opts []option.ClientOption

	config := m.config

	// Workaround for an unusual behaviour (bug?):
	//  the service requires that a quota project be set
	if !config.UserProjectOverride || config.BillingProject == "" {
		config.UserProjectOverride = true
		config.BillingProject = projectID
	}

	opts, err := config.RESTClientOptions()
	if err != nil {
		return nil, err
	}

	gcpClient, err := gcp.NewMetastoreRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building biglake catalog client: %w", err)
	}

	return gcpClient, err
}

func (m *catalogModel) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.BigLakeCatalog{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewCatalogIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	gcpClient, err := m.client(ctx, id.Parent().ProjectID)
	if err != nil {
		return nil, err
	}

	return &catalogAdapter{
		gcpClient: gcpClient,
		id:        id,
		desired:   obj,
	}, nil
}

func (m *catalogModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type catalogAdapter struct {
	gcpClient *gcp.MetastoreClient
	id        *krm.CatalogIdentity
	desired   *krm.BigLakeCatalog
	actual    *pb.Catalog
}

var _ directbase.Adapter = &catalogAdapter{}

func (a *catalogAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting biglake catalog", "name", a.id)

	req := &pb.GetCatalogRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetCatalog(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting biglake catalog %q from gcp: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *catalogAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating biglake catalog", "name", a.id)

	mapCtx := &direct.MapContext{}

	// Catalog message is empty, we just need to pass the IDs.
	resource := BigLakeCatalogSpec_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &pb.CreateCatalogRequest{
		Parent:    a.id.Parent().String(),
		CatalogId: a.id.ID(),
		Catalog:   resource,
	}
	created, err := a.gcpClient.CreateCatalog(ctx, req)
	if err != nil {
		return fmt.Errorf("creating biglake catalog %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully created biglake catalog in gcp", "name", a.id)

	status := &krm.BigLakeCatalogStatus{}
	status.ObservedState = BigLakeCatalogObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// BigLakeCatalog does not support update.
func (a *catalogAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating biglake catalog - no-op", "name", a.id)

	// BiglakeCatalog does not have any mutable fields.
	// Just update the status if we are acquiring.
	mapCtx := &direct.MapContext{}
	if a.desired.Status.ExternalRef == nil {
		observedState := BigLakeCatalogObservedState_FromProto(mapCtx, a.actual)
		if mapCtx.Err() != nil {
			return mapCtx.Err()
		}

		a.desired.Status.ExternalRef = direct.PtrTo(a.id.String())
		a.desired.Status.ObservedState = observedState

		return updateOp.UpdateStatus(ctx, a.desired.Status, nil)
	}

	return nil
}

func (a *catalogAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.BigLakeCatalog{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(BigLakeCatalogSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Parent().ProjectID}
	obj.Spec.Location = a.id.Parent().Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.ID())
	u.SetGroupVersionKind(krm.BigLakeCatalogGVK)

	u.Object = uObj
	return u, nil
}

// Delete implements the Adapter interface.
func (a *catalogAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting biglake catalog", "name", a.id)

	req := &pb.DeleteCatalogRequest{Name: a.id.String()}
	_, err := a.gcpClient.DeleteCatalog(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent biglake catalog, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting biglake catalog %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully deleted biglake catalog", "name", a.id)

	return true, nil
}
