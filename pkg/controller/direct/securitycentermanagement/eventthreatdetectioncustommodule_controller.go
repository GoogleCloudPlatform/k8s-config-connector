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
	"reflect"

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
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
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

	idRaw, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := idRaw.(*krm.EventThreatDetectionCustomModuleIdentity)

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &EventThreatDetectionCustomModuleAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelEventThreatDetectionCustomModule) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type EventThreatDetectionCustomModuleAdapter struct {
	id        *krm.EventThreatDetectionCustomModuleIdentity
	gcpClient *gcp.Client
	desired   *krm.SecurityCenterManagementEventThreatDetectionCustomModule
	actual    *securitycentermanagementpb.EventThreatDetectionCustomModule
}

var _ directbase.Adapter = &EventThreatDetectionCustomModuleAdapter{}

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
		Parent:                           a.id.Parent(),
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
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *EventThreatDetectionCustomModuleAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating EventThreatDetectionCustomModule", "name", a.id)
	mapCtx := &direct.MapContext{}

	desiredPb := SecurityCenterManagementEventThreatDetectionCustomModuleSpec_ToProto(mapCtx, &a.desired.DeepCopy().Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	paths, err := common.CompareProtoMessage(desiredPb, a.actual, common.BasicDiff)
	if err != nil {
		return err
	}

	if !reflect.DeepEqual(a.desired.Spec.DisplayName, a.actual.DisplayName) {
		paths.Insert("display_name")
	}

	updated := a.actual
	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
	} else {
		log.V(2).Info("fields need update", "name", a.id, "paths", paths)
		updateMask := &fieldmaskpb.FieldMask{
			Paths: sets.List(paths),
		}

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

	if a.id.Project != "" {
		obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Project}
	} else if a.id.Folder != "" {
		obj.Spec.FolderRef = &refs.FolderRef{External: a.id.Folder}
	} else if a.id.Organization != "" {
		obj.Spec.OrganizationRef = &refs.OrganizationRef{External: a.id.Organization}
	}
	obj.Spec.Location = direct.LazyPtr(a.id.Location)

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.EventThreatDetectionCustomModule)
	u.SetGroupVersionKind(krm.SecurityCenterManagementEventThreatDetectionCustomModuleGVK)

	u.Object = uObj
	return u, nil
}

func (a *EventThreatDetectionCustomModuleAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting EventThreatDetectionCustomModule", "name", a.id)

	req := &securitycentermanagementpb.DeleteEventThreatDetectionCustomModuleRequest{Name: a.id.String()}
	err := a.gcpClient.DeleteEventThreatDetectionCustomModule(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent EventThreatDetectionCustomModule, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting EventThreatDetectionCustomModule %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted EventThreatDetectionCustomModule", "name", a.id)
	return true, nil
}
