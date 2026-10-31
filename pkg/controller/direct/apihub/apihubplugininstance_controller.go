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

// +generated:controller
// krm.group: apihub.cnrm.cloud.google.com
// krm.version: v1alpha1
// krm.kind: APIHubPluginInstance

package apihub

import (
	"context"
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/apihub/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/apihub/apiv1"
	apihubpb "cloud.google.com/go/apihub/apiv1/apihubpb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.APIHubPluginInstanceGVK, NewPluginInstanceModel)
}

func NewPluginInstanceModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelPluginInstance{config: *config}, nil
}

var _ directbase.Model = &modelPluginInstance{}

type modelPluginInstance struct {
	config config.ControllerConfig
}

func (m *modelPluginInstance) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.APIHubPluginInstance{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.APIHubPluginInstanceIdentity)

	// Get apihub GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &PluginInstanceAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelPluginInstance) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//apihub.googleapis.com/") {
		return nil, nil
	}
	id := &krm.APIHubPluginInstanceIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &PluginInstanceAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}
func (m *modelPluginInstance) client(ctx context.Context) (*gcp.ApiHubPluginClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewApiHubPluginRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building PluginInstance client: %w", err)
	}
	return gcpClient, err
}

type PluginInstanceAdapter struct {
	id        *krm.APIHubPluginInstanceIdentity
	gcpClient *gcp.ApiHubPluginClient
	desired   *krm.APIHubPluginInstance
	actual    *apihubpb.PluginInstance
}

var _ directbase.Adapter = &PluginInstanceAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *PluginInstanceAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting PluginInstance", "name", a.id)

	req := &apihubpb.GetPluginInstanceRequest{
		Name: a.id.String(),
	}
	plugininstancepb, err := a.gcpClient.GetPluginInstance(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting PluginInstance %q: %w", a.id, err)
	}

	a.actual = plugininstancepb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *PluginInstanceAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating PluginInstance", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := APIHubPluginInstanceSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &apihubpb.CreatePluginInstanceRequest{
		Parent:           a.id.ParentString(),
		PluginInstanceId: a.id.Instance,
		PluginInstance:   resource,
	}
	op, err := a.gcpClient.CreatePluginInstance(ctx, req)
	if err != nil {
		return fmt.Errorf("creating PluginInstance %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("PluginInstance %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created PluginInstance", "name", a.id)

	status := &krm.APIHubPluginInstanceStatus{}
	status.ObservedState = APIHubPluginInstanceObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *PluginInstanceAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating PluginInstance", "name", a.id)

	resource, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		APIHubPluginInstanceSpec_FromProto,
		APIHubPluginInstanceSpec_ToProto,
		nil,
	)
	if err != nil {
		return err
	}
	updated := a.actual
	if !report.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id)
	} else {
		report.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, report)
		resource.Name = a.id.String()
		req := &apihubpb.UpdatePluginInstanceRequest{
			UpdateMask:     updateMask,
			PluginInstance: resource,
		}
		updated, err = a.gcpClient.UpdatePluginInstance(ctx, req)
		if err != nil {
			return fmt.Errorf("updating PluginInstance %s: %w", a.id, err)
		}
		log.V(2).Info("successfully updated PluginInstance", "name", a.id)
	}

	status := &krm.APIHubPluginInstanceStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = APIHubPluginInstanceObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *PluginInstanceAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.APIHubPluginInstance{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(APIHubPluginInstanceSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.APIHubPluginInstanceGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *PluginInstanceAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting PluginInstance", "name", a.id)

	req := &apihubpb.DeletePluginInstanceRequest{
		Name: a.id.String(),
	}
	op, err := a.gcpClient.DeletePluginInstance(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent PluginInstance, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting PluginInstance %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted PluginInstance", "name", a.id)
	err = op.Wait(ctx)
	if err != nil {
		if direct.IsNotFound(err) || strings.Contains(err.Error(), "missing \"value\" field") {
			return true, nil
		}
		return false, fmt.Errorf("waiting delete PluginInstance %s: %w", a.id, err)
	}
	return true, nil
}
