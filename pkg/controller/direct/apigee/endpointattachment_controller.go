// Copyright 2024 Google LLC
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

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/apigee/v1alpha1"
	apigeev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/apigee/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	api "google.golang.org/api/apigee/v1"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.ApigeeEndpointAttachmentGVK, NewEndpointAttachmentModel)
}

func NewEndpointAttachmentModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelEndpointAttachment{config: config}, nil
}

var _ directbase.Model = &modelEndpointAttachment{}

type modelEndpointAttachment struct {
	config *config.ControllerConfig
}

func (m *modelEndpointAttachment) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.ApigeeEndpointAttachment{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewApigeeEndpointAttachmentIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	gcpClient, err := newGCPClient(ctx, m.config)
	if err != nil {
		return nil, err
	}
	return &EndpointAttachmentAdapter{
		id:                       id,
		desired:                  obj,
		endpointattachmentclient: gcpClient.endpointAttachmentClient(),
		operationsClient:         gcpClient.operationsClient(),
	}, nil
}

func (m *modelEndpointAttachment) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type EndpointAttachmentAdapter struct {
	id                       *krm.EndpointAttachmentIdentity
	desired                  *krm.ApigeeEndpointAttachment
	actual                   *api.GoogleCloudApigeeV1EndpointAttachment
	endpointattachmentclient *api.OrganizationsEndpointAttachmentsService
	operationsClient         *api.OrganizationsOperationsService
}

var _ directbase.Adapter = &EndpointAttachmentAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *EndpointAttachmentAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting EndpointAttachment", "name", a.id)
	endpointAttachment, err := a.endpointattachmentclient.Get(a.fullyQualifiedName()).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting EndpointAttachment %q: %w", a.id, err)
	}

	a.actual = endpointAttachment
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *EndpointAttachmentAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating EndpointAttachment", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := ApigeeEndpointAttachmentSpec_ToProto(mapCtx, &desired.Spec, a.desired.Name)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	op, err := a.endpointattachmentclient.Create(a.id.Parent().String(), resource).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("creating EndpointAttachment %s: %w", a.id, err)
	}
	if err := WaitForApigeeOp(ctx, a.operationsClient, op); err != nil {
		return fmt.Errorf("waiting for EndpointAttachment %s creation: %w", a.id, err)
	}

	created, err := a.endpointattachmentclient.Get(a.fullyQualifiedName()).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("getting created EndpointAttachment: %w", err)
	}
	log.V(2).Info("successfully created EndpointAttachment", "name", a.id)

	status := &krm.ApigeeEndpointAttachmentStatus{}
	status.ObservedState = ApigeeEndpointAttachmentObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *EndpointAttachmentAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	return fmt.Errorf("Apigee Endpoint Attachment does not support update")
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *EndpointAttachmentAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ApigeeEndpointAttachment{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(ApigeeEndpointAttachmentSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.Parent.OrganizationRef = &apigeev1beta1.OrganizationRef{External: a.id.Parent().String()}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.ID())
	u.SetGroupVersionKind(krm.ApigeeEndpointAttachmentGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *EndpointAttachmentAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting EndpointAttachment", "name", a.id)

	op, err := a.endpointattachmentclient.Delete(a.fullyQualifiedName()).Context(ctx).Do()
	if err != nil {
		return false, fmt.Errorf("deleting EndpointAttachment %s: %w", a.id, err)
	}
	if err := WaitForApigeeOp(ctx, a.operationsClient, op); err != nil {
		return false, fmt.Errorf("waiting delete EndpointAttachment %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted EndpointAttachment", "name", a.id)
	return true, nil
}

func (a *EndpointAttachmentAdapter) fullyQualifiedName() string {
	return a.id.String()
}
