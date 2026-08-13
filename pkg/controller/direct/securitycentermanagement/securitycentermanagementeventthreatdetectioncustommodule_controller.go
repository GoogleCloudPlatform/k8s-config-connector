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

package securitycentermanagement

import (
	"context"
	"fmt"

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/securitycentermanagement/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	gcp "cloud.google.com/go/securitycentermanagement/apiv1"

	securitycentermanagementpb "cloud.google.com/go/securitycentermanagement/apiv1/securitycentermanagementpb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.SecurityCenterManagementEventThreatDetectionCustomModuleGVK, NewEventThreatDetectionCustomModuleModel)
}

func NewEventThreatDetectionCustomModuleModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelEventThreatDetectionCustomModule{config: *config}, nil
}

var _ directbase.Model = &modelEventThreatDetectionCustomModule{}

type modelEventThreatDetectionCustomModule struct {
	config config.ControllerConfig
}

func (m *modelEventThreatDetectionCustomModule) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building EventThreatDetectionCustomModule client: %w", err)
	}
	return gcpClient, err
}

func (m *modelEventThreatDetectionCustomModule) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.SecurityCenterManagementEventThreatDetectionCustomModule{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}

	// Get securitycentermanagement GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &EventThreatDetectionCustomModuleAdapter{
		id:        id.(*krm.SecurityCenterManagementEventThreatDetectionCustomModuleIdentity),
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelEventThreatDetectionCustomModule) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type EventThreatDetectionCustomModuleAdapter struct {
	id        *krm.SecurityCenterManagementEventThreatDetectionCustomModuleIdentity
	gcpClient *gcp.Client
	desired   *krm.SecurityCenterManagementEventThreatDetectionCustomModule
	actual    *securitycentermanagementpb.EventThreatDetectionCustomModule
}

var _ directbase.Adapter = &EventThreatDetectionCustomModuleAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *EventThreatDetectionCustomModuleAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting EventThreatDetectionCustomModule", "name", a.id)

	req := &securitycentermanagementpb.GetEventThreatDetectionCustomModuleRequest{Name: a.id.String()}
	eventthreatdetectioncustommodulepb, err := a.gcpClient.GetEventThreatDetectionCustomModule(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting EventThreatDetectionCustomModule %q: %w", a.id, err)
	}

	a.actual = eventthreatdetectioncustommodulepb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *EventThreatDetectionCustomModuleAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating EventThreatDetectionCustomModule", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := SecurityCenterManagementEventThreatDetectionCustomModuleSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &securitycentermanagementpb.CreateEventThreatDetectionCustomModuleRequest{
		Parent:                           a.id.ParentString(),
		EventThreatDetectionCustomModule: resource,
	}
	created, err := a.gcpClient.CreateEventThreatDetectionCustomModule(ctx, req)
	if err != nil {
		return fmt.Errorf("creating EventThreatDetectionCustomModule %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created EventThreatDetectionCustomModule", "name", a.id)

	status := &krm.SecurityCenterManagementEventThreatDetectionCustomModuleStatus{}
	status.ObservedState = SecurityCenterManagementEventThreatDetectionCustomModuleObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(created.GetName())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *EventThreatDetectionCustomModuleAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating EventThreatDetectionCustomModule", "name", a.id)
	mapCtx := &direct.MapContext{}

	desiredPb := SecurityCenterManagementEventThreatDetectionCustomModuleSpec_ToProto(mapCtx, &a.desired.DeepCopy().Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	clonedDesired := proto.Clone(desiredPb).(*securitycentermanagementpb.EventThreatDetectionCustomModule)
	actualSpec := SecurityCenterManagementEventThreatDetectionCustomModuleSpec_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	maskedActual := SecurityCenterManagementEventThreatDetectionCustomModuleSpec_ToProto(mapCtx, actualSpec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	diffs, updateMask, err := common.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return err
	}

	updated := a.actual
	if !diffs.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id)
	} else {
		log.V(2).Info("fields need update", "name", a.id, "updateMask", updateMask)

		desiredPb.Name = a.id.String()
		req := &securitycentermanagementpb.UpdateEventThreatDetectionCustomModuleRequest{
			UpdateMask:                       updateMask,
			EventThreatDetectionCustomModule: desiredPb,
		}
		updated, err = a.gcpClient.UpdateEventThreatDetectionCustomModule(ctx, req)
		if err != nil {
			return fmt.Errorf("updating EventThreatDetectionCustomModule %s: %w", a.id, err)
		}
		log.V(2).Info("successfully updated EventThreatDetectionCustomModule", "name", a.id)
	}

	status := &krm.SecurityCenterManagementEventThreatDetectionCustomModuleStatus{}
	status.ObservedState = SecurityCenterManagementEventThreatDetectionCustomModuleObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *EventThreatDetectionCustomModuleAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.SecurityCenterManagementEventThreatDetectionCustomModule{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(SecurityCenterManagementEventThreatDetectionCustomModuleSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.OrganizationRef = &refs.OrganizationRef{External: a.id.Organization}
	obj.Spec.Location = &a.id.Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.EventThreatDetectionCustomModule)
	u.SetGroupVersionKind(krm.SecurityCenterManagementEventThreatDetectionCustomModuleGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *EventThreatDetectionCustomModuleAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting EventThreatDetectionCustomModule", "name", a.id)

	req := &securitycentermanagementpb.DeleteEventThreatDetectionCustomModuleRequest{Name: a.id.String()}
	err := a.gcpClient.DeleteEventThreatDetectionCustomModule(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent EventThreatDetectionCustomModule, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting EventThreatDetectionCustomModule %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted EventThreatDetectionCustomModule", "name", a.id)
	return true, nil
}
