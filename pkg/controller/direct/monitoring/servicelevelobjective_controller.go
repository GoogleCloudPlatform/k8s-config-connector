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
	"reflect"

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
	monitoringpb "cloud.google.com/go/monitoring/v3/monitoringpb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.MonitoringServiceLevelObjectiveGVK, NewServiceLevelObjectiveModel)
}

func NewServiceLevelObjectiveModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelServiceLevelObjective{config: *config}, nil
}

var _ directbase.Model = &modelServiceLevelObjective{}

type modelServiceLevelObjective struct {
	config config.ControllerConfig
}

func (m *modelServiceLevelObjective) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building ServiceLevelObjective client: %w", err)
	}
	return gcpClient, err
}

func (m *modelServiceLevelObjective) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.MonitoringServiceLevelObjective{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewServiceLevelObjectiveIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Get monitoring GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &ServiceLevelObjectiveAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelServiceLevelObjective) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type ServiceLevelObjectiveAdapter struct {
	id        *krm.ServiceLevelObjectiveIdentity
	gcpClient *gcp.Client
	desired   *krm.MonitoringServiceLevelObjective
	actual    *monitoringpb.ServiceLevelObjective
}

var _ directbase.Adapter = &ServiceLevelObjectiveAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *ServiceLevelObjectiveAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ServiceLevelObjective", "name", a.id)

	req := &monitoringpb.GetServiceLevelObjectiveRequest{Name: a.id}
	servicelevelobjectivepb, err := a.gcpClient.GetServiceLevelObjective(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ServiceLevelObjective %q: %w", a.id, err)
	}

	a.actual = servicelevelobjectivepb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *ServiceLevelObjectiveAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating ServiceLevelObjective", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := MonitoringServiceLevelObjectiveSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	// TODO(contributor): Complete the gcp "CREATE" or "INSERT" request.
	req := &monitoringpb.CreateServiceLevelObjectiveRequest{
		Parent:                a.id.Parent().String(),
		ServiceLevelObjective: resource,
	}
	op, err := a.gcpClient.CreateServiceLevelObjective(ctx, req)
	if err != nil {
		return fmt.Errorf("creating ServiceLevelObjective %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("ServiceLevelObjective %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created ServiceLevelObjective", "name", a.id)

	status := &krm.MonitoringServiceLevelObjectiveStatus{}
	status.ObservedState = MonitoringServiceLevelObjectiveObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = &a.id.External
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *ServiceLevelObjectiveAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating ServiceLevelObjective", "name", a.id)
	mapCtx := &direct.MapContext{}

	desiredPb := MonitoringServiceLevelObjectiveSpec_ToProto(mapCtx, &a.desired.DeepCopy().Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	paths := []string{}
	// Option 1: This option is good for proto that has `field_mask` for output-only, immutable, required/optional.
	// TODO(contributor): If choosing this option, remove the "Option 2" code.
	{
		var err error
		paths, err = common.CompareProtoMessage(desiredPb, a.actual, common.BasicDiff)
		if err != nil {
			return err
		}
	}

	// Option 2: manually add all mutable fields.
	// TODO(contributor): If choosing this option, remove the "Option 1" code.
	{
		if !reflect.DeepEqual(a.desired.Spec.DisplayName, a.actual.DisplayName) {
			paths = append(paths, "display_name")
		}
	}

	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id.External)
		status := &krm.MonitoringServiceLevelObjectiveStatus{}
		status.ObservedState = MonitoringServiceLevelObjectiveObservedState_FromProto(mapCtx, a.actual)
		if mapCtx.Err() != nil {
			return mapCtx.Err()
		}
		return updateOp.UpdateStatus(ctx, status, nil)
	}
	updateMask := &fieldmaskpb.FieldMask{
		Paths: sets.List(paths)}

	// TODO(contributor): Complete the gcp "UPDATE" or "PATCH" request.
	req := &monitoringpb.UpdateServiceLevelObjectiveRequest{
		Name:                  a.id.External,
		UpdateMask:            updateMask,
		ServiceLevelObjective: desiredPb,
	}
	op, err := a.gcpClient.UpdateServiceLevelObjective(ctx, req)
	if err != nil {
		return fmt.Errorf("updating ServiceLevelObjective %s: %w", a.id.External, err)
	}
	updated, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("ServiceLevelObjective %s waiting update: %w", a.id.External, err)
	}
	log.V(2).Info("successfully updated ServiceLevelObjective", "name", a.id.External)

	status := &krm.MonitoringServiceLevelObjectiveStatus{}
	status.ObservedState = MonitoringServiceLevelObjectiveObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *ServiceLevelObjectiveAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.MonitoringServiceLevelObjective{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(MonitoringServiceLevelObjectiveSpec_FromProto(mapCtx, a.actual))
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
	u.SetGroupVersionKind(krm.MonitoringServiceLevelObjectiveGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *ServiceLevelObjectiveAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ServiceLevelObjective", "name", a.id)

	req := &monitoringpb.DeleteServiceLevelObjectiveRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteServiceLevelObjective(ctx, req)
	if err != nil {
		return false, fmt.Errorf("deleting ServiceLevelObjective %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted ServiceLevelObjective", "name", a.id)

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete ServiceLevelObjective %s: %w", a.id, err)
	}
	return true, nil
}
