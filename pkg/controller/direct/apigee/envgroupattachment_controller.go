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

const aegaCtrlName = "apigee-envgroupattachment-controller"

func init() {
	registry.RegisterModel(krm.EnvgroupAttachmentGVK, NewApigeeEnvgroupAttachmentModel)
}

func NewApigeeEnvgroupAttachmentModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelEnvgroupAttachment{config: config}, nil
}

var _ directbase.Model = &modelEnvgroupAttachment{}

type modelEnvgroupAttachment struct {
	config *config.ControllerConfig
}

func (m *modelEnvgroupAttachment) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	gcpClient, err := newGCPClient(ctx, m.config)
	if err != nil {
		return nil, err
	}

	obj := &krm.ApigeeEnvgroupAttachment{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewEnvgroupAttachmentIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desired := obj
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	desired.Name = id.ID()

	return &EnvgroupAttachmentAdapter{
		id:               id,
		desired:          desired,
		client:           gcpClient.envgroupAttachmentsClient(),
		operationsClient: gcpClient.operationsClient(),
	}, nil
}

func (m *modelEnvgroupAttachment) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type EnvgroupAttachmentAdapter struct {
	id               *krm.EnvgroupAttachmentIdentity
	desired          *krm.ApigeeEnvgroupAttachment
	actual           *api.GoogleCloudApigeeV1EnvironmentGroupAttachment
	client           *api.OrganizationsEnvgroupsAttachmentsService
	operationsClient *api.OrganizationsOperationsService
}

var _ directbase.Adapter = &EnvgroupAttachmentAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers EnvgroupAttachmentAdapter `Update` call.
// Return false means the object is not found. This triggers EnvgroupAttachmentAdapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *EnvgroupAttachmentAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx).WithName(aegaCtrlName)
	log.V(2).Info("getting ApigeeEnvgroupAttachment", "name", a.id)

	ea, err := a.client.Get(a.fullyQualifiedName()).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ApigeeEnvgroupAttachment %q: %w", a.id, err)
	}

	a.actual = ea
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *EnvgroupAttachmentAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx).WithName(aegaCtrlName)
	log.V(2).Info("creating ApigeeEnvgroupAttachment", "name", a.id)
	mapCtx := &direct.MapContext{}

	req := EnvgroupAttachmentSpec_ToAPI(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	req.Name = a.id.ID()

	op, err := a.client.Create(a.id.Parent().String(), req).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("creating ApigeeEnvgroupAttachment %s: %w", a.fullyQualifiedName(), err)
	}

	if err := WaitForApigeeOp(ctx, a.operationsClient, op); err != nil {
		return fmt.Errorf("waiting for ApigeeEnvgroupAttachment %s creation: %w", a.id, err)
	}

	created, err := a.client.Get(a.fullyQualifiedName()).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("getting created ApigeeEnvgroupAttachment: %w", err)
	}

	log.V(2).Info("successfully created ApigeeEnvgroupAttachment", "ApigeeEnvgroupAttachment", created)

	status := &krm.ApigeeEnvgroupAttachmentStatus{}
	status.ObservedState = EnvgroupAttachmentObservedState_FromAPI(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *EnvgroupAttachmentAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx).WithName(aegaCtrlName)

	if *a.desired.Status.ExternalRef == "" {
		// If it is the first reconciliation after switching to direct controller,
		// then update Status to fill out the ExternalRef even if there is
		// no update.
		status := a.desired.Status
		status.ExternalRef = direct.LazyPtr(a.id.String())
		return updateOp.UpdateStatus(ctx, status, nil)
	}

	log.Info("ApigeeEnvgroupAttachment has no update function", "name", a.id)
	return nil
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *EnvgroupAttachmentAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ApigeeEnvgroupAttachment{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(EnvgroupAttachmentSpec_FromAPI(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.OrganizationRef = &krmv1beta1.OrganizationRef{External: a.id.Parent().String()}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.ID())
	u.SetGroupVersionKind(krm.EnvgroupAttachmentGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *EnvgroupAttachmentAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx).WithName(aegaCtrlName)
	log.V(2).Info("deleting ApigeeEnvgroupAttachment", "name", a.id)

	op, err := a.client.Delete(a.fullyQualifiedName()).Context(ctx).Do()

	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent ApigeeEnvgroupAttachment, assuming it was already deleted", "name", a.id.String())
			return true, nil
		}
		return false, fmt.Errorf("deleting ApigeeEnvgroupAttachment %s: %w", a.id, err)
	}

	if err := WaitForApigeeOp(ctx, a.operationsClient, op); err != nil {
		return false, fmt.Errorf("ApigeeEnvgroupAttachment deletion failed: %w", err)
	}
	return true, nil
}

func (a *EnvgroupAttachmentAdapter) fullyQualifiedName() string {
	return a.id.String()
}
