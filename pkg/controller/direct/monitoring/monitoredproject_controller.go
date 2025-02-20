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

package monitoring

import (
	"context"
	"fmt"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/monitoring/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	// TODO(contributor): Update the import with the google cloud client
	gcp "cloud.google.com/go/monitoring/apiv1"

	// TODO(contributor): Update the import with the google cloud client api protobuf
	monitoringpb "cloud.google.com/go/monitoring/v1/monitoringpb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.MonitoringMonitoredProjectGVK, NewMonitoredProjectModel)
}

func NewMonitoredProjectModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelMonitoredProject{config: *config}, nil
}

var _ directbase.Model = &modelMonitoredProject{}

type modelMonitoredProject struct {
	config config.ControllerConfig
}

func (m *modelMonitoredProject) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building MonitoredProject client: %w", err)
	}
	return gcpClient, err
}

func (m *modelMonitoredProject) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.MonitoringMonitoredProject{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewMonitoredProjectIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Get monitoring GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &MonitoredProjectAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelMonitoredProject) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type MonitoredProjectAdapter struct {
	id        *krm.MonitoredProjectIdentity
	gcpClient *gcp.Client
	desired   *krm.MonitoringMonitoredProject
	actual    *monitoringpb.MonitoredProject
}

var _ directbase.Adapter = &MonitoredProjectAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *MonitoredProjectAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting MonitoredProject", "name", a.id)

	req := &monitoringpb.GetMonitoredProjectRequest{Name: a.id.String()}
	monitoredprojectpb, err := a.gcpClient.GetMonitoredProject(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting MonitoredProject %q: %w", a.id, err)
	}

	a.actual = monitoredprojectpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *MonitoredProjectAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating MonitoredProject", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := MonitoringMonitoredProjectSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	// TODO(contributor): Complete the gcp "CREATE" or "INSERT" request.
	req := &monitoringpb.CreateMonitoredProjectRequest{
		Parent:           a.id.Parent().String(),
		MonitoredProject: resource,
	}
	op, err := a.gcpClient.CreateMonitoredProject(ctx, req)
	if err != nil {
		return fmt.Errorf("creating MonitoredProject %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("MonitoredProject %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created MonitoredProject", "name", a.id)

	status := &krm.MonitoringMonitoredProjectStatus{}
	status.ObservedState = MonitoringMonitoredProjectObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *MonitoredProjectAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating MonitoredProject", "name", a.id)
	mapCtx := &direct.MapContext{}

	desiredPb := MonitoringMonitoredProjectSpec_ToProto(mapCtx, &a.desired.DeepCopy().Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	var err error
	paths, err = common.CompareProtoMessage(desiredPb, a.actual, common.BasicDiff)
	if err != nil {
		return err
	}
	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		status := &krm.MonitoringMonitoredProjectStatus{}
		status.ObservedState = MonitoringMonitoredProjectObservedState_FromProto(mapCtx, a.actual)
		if mapCtx.Err() != nil {
			return mapCtx.Err()
		}
		return updateOp.UpdateStatus(ctx, status, nil)
	}
	updateMask := &fieldmaskpb.FieldMask{
		Paths: sets.List(paths)}

	// TODO(contributor): Complete the gcp "UPDATE" or "PATCH" request.
	req := &monitoringpb.UpdateMonitoredProjectRequest{
		Name:             a.id,
		UpdateMask:       updateMask,
		MonitoredProject: desiredPb,
	}
	op, err := a.gcpClient.UpdateMonitoredProject(ctx, req)
	if err != nil {
		return fmt.Errorf("updating MonitoredProject %s: %w", a.id, err)
	}
	updated, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("MonitoredProject %s waiting update: %w", a.id, err)
	}
	log.V(2).Info("successfully updated MonitoredProject", "name", a.id)

	status := &krm.MonitoringMonitoredProjectStatus{}
	status.ObservedState = MonitoringMonitoredProjectObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *MonitoredProjectAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.MonitoringMonitoredProject{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(MonitoringMonitoredProjectSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Parent().ProjectID}
	obj.Spec.Location = a.id.Parent().Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.actual.Id)
	u.SetGroupVersionKind(krm.MonitoringMonitoredProjectGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *MonitoredProjectAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting MonitoredProject", "name", a.id)

	req := &monitoringpb.DeleteMonitoredProjectRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteMonitoredProject(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent MonitoredProject, assuming it was already deleted", "name", a.id.String())
			return true, nil
		}
		return false, fmt.Errorf("deleting MonitoredProject %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted MonitoredProject", "name", a.id)

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete MonitoredProject %s: %w", a.id, err)
	}
	return true, nil
}
