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
	"encoding/json"
	"fmt"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/apigee/v1beta1"
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
	registry.RegisterModel(krm.ApigeeEnvgroupAttachmentGVK, NewApigeeEnvgroupAttachmentModel)
}

func NewApigeeEnvgroupAttachmentModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelApigeeEnvgroupAttachment{config: config}, nil
}

var _ directbase.Model = &modelApigeeEnvgroupAttachment{}

type modelApigeeEnvgroupAttachment struct {
	config *config.ControllerConfig
}

func (m *modelApigeeEnvgroupAttachment) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.ApigeeEnvgroupAttachment{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.ApigeeEnvgroupAttachmentIdentity)

	// Get apigee GCP client
	gcpClient, err := newGCPClient(ctx, m.config)
	if err != nil {
		return nil, err
	}
	return &ApigeeEnvgroupAttachmentAdapter{
		id:                id,
		k8sClient:         reader,
		attachmentsClient: gcpClient.envgroupsAttachmentsClient(),
		operationsClient:  gcpClient.operationsClient(),
		desired:           obj,
	}, nil
}

func (m *modelApigeeEnvgroupAttachment) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type ApigeeEnvgroupAttachmentAdapter struct {
	id                *krm.ApigeeEnvgroupAttachmentIdentity
	k8sClient         client.Reader
	attachmentsClient *api.OrganizationsEnvgroupsAttachmentsService
	operationsClient  *api.OrganizationsOperationsService
	desired           *krm.ApigeeEnvgroupAttachment
	actual            *api.GoogleCloudApigeeV1EnvironmentGroupAttachment
}

var _ directbase.Adapter = &ApigeeEnvgroupAttachmentAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *ApigeeEnvgroupAttachmentAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ApigeeInstanceAttachment", "name", a.id)

	if a.id.String() != "" {
		// If resource ID is specified, we can use it to look up the attachment.
		attachment, err := a.attachmentsClient.Get(a.id.String()).Context(ctx).Do()
		if err != nil {
			if direct.IsNotFound(err) {
				return false, nil
			}
			return false, fmt.Errorf("getting ApigeeInstanceAttachment %q: %w", a.id, err)
		}
		a.actual = attachment
		return true, nil
	}

	return false, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *ApigeeEnvgroupAttachmentAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating ApigeeEnvgroupAttachment", "name", a.desired.Name)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()

	// Resolve references
	if err := ResolveApigeeEnvgroupAttachmentRefs(ctx, a.k8sClient, desired); err != nil {
		return err
	}
	// Convert to proto
	resource := ApigeeEnvgroupAttachmentSpec_ToAPI(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	// HACK: Environment field format required by GCP API is name-only, not fully-qualified ID.
	// So, we fix this by editing the environment to be name-only.
	envName, err := GetNameOfEnvironment(resource.Environment)
	if err != nil {
		return err
	}
	resource.Environment = envName

	op, err := a.attachmentsClient.Create(a.id.ParentID.String(), resource).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("creating ApigeeEnvgroupAttachment %s: %w", a.desired.Name, err)
	}
	if err := WaitForApigeeOp(ctx, a.operationsClient, op); err != nil {
		return fmt.Errorf("ApigeeEnvgroupAttachment %s waiting creation: %w", a.desired.Name, err)
	}

	// Get response from completed operation and unmarshal it so that we can determine
	// the server-generated ID for the newly-created envgroup attachment.
	result, err := a.operationsClient.Get(op.Name).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("getting create ApigeeEnvgroupAttachment %s operation: %w", a.desired.Name, err)
	}
	var resultResponse api.GoogleCloudApigeeV1EnvironmentGroupAttachment
	if err := json.Unmarshal(result.Response, &resultResponse); err != nil {
		return fmt.Errorf("unmarshalling create ApigeeEnvgroupAttachment %s operation response: %w", a.desired.Name, err)
	}

	// Update identity to include server-generated ID.
	a.id.ResourceID = resultResponse.Name

	// Get the newly-created envgroup attachment.
	created, err := a.attachmentsClient.Get(a.id.String()).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("getting created ApigeeEnvgroupAttachment %s: %w", a.desired.Name, err)
	}

	log.V(2).Info("successfully created ApigeeEnvgroupAttachment", "name", a.desired.Name, "id", a.id)

	status := &krm.ApigeeEnvgroupAttachmentStatus{}
	status.ObservedState = ApigeeEnvgroupAttachmentObservedState_FromAPI(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *ApigeeEnvgroupAttachmentAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating ApigeeEnvgroupAttachment", "id", a.id)
	mapCtx := &direct.MapContext{}

	// There are no fields in the GCP ApigeeEnvgroupAttachment API that can be updated.
	// So, we will only update the KRM status.
	status := &krm.ApigeeEnvgroupAttachmentStatus{}
	status.ObservedState = ApigeeEnvgroupAttachmentObservedState_FromAPI(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *ApigeeEnvgroupAttachmentAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ApigeeEnvgroupAttachment{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(ApigeeEnvgroupAttachmentSpec_FromAPI(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ResourceID = &a.actual.Name
	obj.Spec.EnvgroupRef = &krm.ApigeeEnvgroupRef{External: a.id.ParentID.String()}

	// HACK: Environment field format returned from GCP API is name-only, not fully-qualified ID.
	// So, we fix this by building a fully-qualified environment ID.
	if obj.Spec.EnvironmentRef != nil {
		environmentID := &krm.ApigeeEnvironmentIdentity{}
		environmentName := obj.Spec.EnvironmentRef.External
		if err := environmentID.FromExternal(a.id.ParentID.ParentID.String() + "/" + krm.ApigeeEnvironmentIDToken + "/" + environmentName); err != nil {
			return nil, err
		}
		obj.Spec.EnvironmentRef.External = environmentID.String()
	}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.ResourceID)
	u.SetGroupVersionKind(krm.ApigeeEnvgroupAttachmentGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *ApigeeEnvgroupAttachmentAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ApigeeEnvgroupAttachment", "id", a.id)

	op, err := a.attachmentsClient.Delete(a.id.String()).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent ApigeeEnvgroupAttachment, assuming it was already deleted", "id", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting ApigeeEnvgroupAttachment %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted ApigeeEnvgroupAttachment", "id", a.id)

	if err := WaitForApigeeOp(ctx, a.operationsClient, op); err != nil {
		return false, fmt.Errorf("waiting delete ApigeeEnvgroupAttachment %s: %w", a.id, err)
	}
	return true, nil
}
