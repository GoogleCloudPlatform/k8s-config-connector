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

package apigee

import (
	"context"
	"fmt"
	"reflect"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/apigee/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	api "google.golang.org/api/apigee/v1"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.ApigeeSyncAuthorizationGVK, NewApigeeSyncAuthorizationModel)
}

func NewApigeeSyncAuthorizationModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelApigeeSyncAuthorization{config: config}, nil
}

var _ directbase.Model = &modelApigeeSyncAuthorization{}

type modelApigeeSyncAuthorization struct {
	config *config.ControllerConfig
}

func (m *modelApigeeSyncAuthorization) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ApigeeSyncAuthorization{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.ApigeeSyncAuthorizationIdentity)

	gcpClient, err := newGCPClient(ctx, m.config)
	if err != nil {
		return nil, err
	}
	return &ApigeeSyncAuthorizationAdapter{
		id:                  id,
		k8sClient:           reader,
		organizationsClient: gcpClient.organizationsClient(),
		desired:             obj,
	}, nil
}

func (m *modelApigeeSyncAuthorization) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	format := krm.ApigeeSyncAuthorizationIdentityFormat
	id, match, err := format.Parse(url)
	if err != nil {
		return nil, err
	}
	if !match {
		return nil, nil
	}
	gcpClient, err := newGCPClient(ctx, m.config)
	if err != nil {
		return nil, err
	}
	return &ApigeeSyncAuthorizationAdapter{
		id:                  id,
		organizationsClient: gcpClient.organizationsClient(),
	}, nil
}

type ApigeeSyncAuthorizationAdapter struct {
	id                  *krm.ApigeeSyncAuthorizationIdentity
	k8sClient           client.Reader
	organizationsClient *api.OrganizationsService
	desired             *krm.ApigeeSyncAuthorization
	actual              *api.GoogleCloudApigeeV1SyncAuthorization
}

var _ directbase.Adapter = &ApigeeSyncAuthorizationAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *ApigeeSyncAuthorizationAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ApigeeSyncAuthorization", "name", a.id)

	req := &api.GoogleCloudApigeeV1GetSyncAuthorizationRequest{}
	resource, err := a.organizationsClient.GetSyncAuthorization(a.id.ParentString(), req).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ApigeeSyncAuthorization %q: %w", a.id, err)
	}

	a.actual = resource
	return true, nil
}

// Create creates the resource in GCP based on `spec` and updates the Config Connector object `status` based on the GCP response.
func (a *ApigeeSyncAuthorizationAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating ApigeeSyncAuthorization", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := ApigeeSyncAuthorizationSpec_ToAPI(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	if a.actual != nil {
		resource.Etag = a.actual.Etag
	}

	updated, err := a.organizationsClient.SetSyncAuthorization(a.id.ParentString(), resource).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("creating ApigeeSyncAuthorization %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created ApigeeSyncAuthorization", "name", a.id)

	status := &krm.ApigeeSyncAuthorizationStatus{}
	status.Etag = direct.LazyPtr(updated.Etag)
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and updates the Config Connector object `status` based on the GCP response.
func (a *ApigeeSyncAuthorizationAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating ApigeeSyncAuthorization", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := ApigeeSyncAuthorizationSpec_ToAPI(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	report := &structuredreporting.Diff{Object: updateOp.GetUnstructured()}

	// Sort identities to ensure deterministic comparison
	if !equalIdentities(a.actual.Identities, resource.Identities) {
		report.AddField("identities", a.actual.Identities, resource.Identities)
	}

	if len(report.Fields) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		status := &krm.ApigeeSyncAuthorizationStatus{}
		status.Etag = direct.LazyPtr(a.actual.Etag)
		return updateOp.UpdateStatus(ctx, status, nil)
	}

	structuredreporting.ReportDiff(ctx, report)
	resource.Etag = a.actual.Etag

	updated, err := a.organizationsClient.SetSyncAuthorization(a.id.ParentString(), resource).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("updating ApigeeSyncAuthorization %s: %w", a.id, err)
	}
	log.V(2).Info("successfully updated ApigeeSyncAuthorization", "name", a.id)

	status := &krm.ApigeeSyncAuthorizationStatus{}
	status.Etag = direct.LazyPtr(updated.Etag)
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *ApigeeSyncAuthorizationAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ApigeeSyncAuthorization{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(ApigeeSyncAuthorizationSpec_FromAPI(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.Organization)
	u.SetGroupVersionKind(krm.ApigeeSyncAuthorizationGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *ApigeeSyncAuthorizationAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ApigeeSyncAuthorization is a no-op as the resource cannot be deleted from GCP", "name", a.id)
	return true, nil
}

func equalIdentities(a, b []string) bool {
	if len(a) == 0 && len(b) == 0 {
		return true
	}
	return reflect.DeepEqual(asSortedCopy(a), asSortedCopy(b))
}
