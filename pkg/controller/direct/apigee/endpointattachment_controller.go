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
	krmv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/apigee/v1beta1"
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
	registry.RegisterModel(krm.ApigeeEndpointAttachmentGVK, NewApigeeEndpointAttachmentModel)
}

func NewApigeeEndpointAttachmentModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelApigeeEndpointAttachment{config: config}, nil
}

var _ directbase.Model = &modelApigeeEndpointAttachment{}

type modelApigeeEndpointAttachment struct {
	config *config.ControllerConfig
}

func (m *modelApigeeEndpointAttachment) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.ApigeeEndpointAttachment{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.ApigeeEndpointAttachmentIdentity)

	// Get apigee GCP client
	gcpClient, err := newGCPClient(ctx, m.config)
	if err != nil {
		return nil, err
	}
	return &ApigeeEndpointAttachmentAdapter{
		id:                id,
		k8sClient:         reader,
		attachmentsClient: gcpClient.endpointsAttachmentsClient(),
		operationsClient:  gcpClient.operationsClient(),
		desired:           obj,
	}, nil
}

func (m *modelApigeeEndpointAttachment) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type ApigeeEndpointAttachmentAdapter struct {
	id                *krm.ApigeeEndpointAttachmentIdentity
	k8sClient         client.Reader
	attachmentsClient *api.OrganizationsEndpointAttachmentsService
	operationsClient  *api.OrganizationsOperationsService
	desired           *krm.ApigeeEndpointAttachment
	actual            *api.GoogleCloudApigeeV1EndpointAttachment
}

var _ directbase.Adapter = &ApigeeEndpointAttachmentAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *ApigeeEndpointAttachmentAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ApigeeEndpointAttachment", "name", a.id)

	resource, err := a.attachmentsClient.Get(a.id.String()).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ApigeeEndpointAttachment %q: %w", a.id, err)
	}

	a.actual = resource
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *ApigeeEndpointAttachmentAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating ApigeeEndpointAttachment", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()

	// Resolve references
	if err := ResolveApigeeEndpointAttachmentRefs(ctx, a.k8sClient, desired); err != nil {
		return err
	}
	// Convert to proto
	resource := ApigeeEndpointAttachmentSpec_ToAPI(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	op, err := a.attachmentsClient.Create(a.id.ParentID.String(), resource).EndpointAttachmentId(a.id.ResourceID).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("creating ApigeeEndpointAttachment %s: %w", a.id, err)
	}
	if err := WaitForApigeeOp(ctx, a.operationsClient, op); err != nil {
		return fmt.Errorf("ApigeeEndpointAttachment %s waiting creation: %w", a.id, err)
	}

	created, err := a.attachmentsClient.Get(a.id.String()).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("getting created ApigeeEndpointAttachment: %w", err)
	}

	log.V(2).Info("successfully created ApigeeEndpointAttachment", "name", a.id)

	status := &krm.ApigeeEndpointAttachmentStatus{}
	status.ObservedState = ApigeeEndpointAttachmentObservedState_FromAPI(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *ApigeeEndpointAttachmentAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating ApigeeEndpointAttachment", "id", a.id)
	mapCtx := &direct.MapContext{}

	// There are no fields in the GCP ApigeeEndpointAttachment API that can be updated.
	// So, we will only update the KRM status.
	status := &krm.ApigeeEndpointAttachmentStatus{}
	status.ObservedState = ApigeeEndpointAttachmentObservedState_FromAPI(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *ApigeeEndpointAttachmentAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ApigeeEndpointAttachment{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(ApigeeEndpointAttachmentSpec_FromAPI(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.OrganizationRef = &krmv1beta1.ApigeeOrganizationRef{External: a.id.ParentID.String()}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.ResourceID)
	u.SetGroupVersionKind(krm.ApigeeEndpointAttachmentGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *ApigeeEndpointAttachmentAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ApigeeEndpointAttachment", "name", a.id)

	op, err := a.attachmentsClient.Delete(a.id.String()).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent ApigeeEndpointAttachment, assuming it was already deleted", "name", a.id.String())
			return true, nil
		}
		return false, fmt.Errorf("deleting ApigeeEndpointAttachment %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted ApigeeEndpointAttachment", "name", a.id)

	if err := WaitForApigeeOp(ctx, a.operationsClient, op); err != nil {
		return false, fmt.Errorf("waiting delete ApigeeEndpointAttachment %s: %w", a.id, err)
	}
	return true, nil
}
