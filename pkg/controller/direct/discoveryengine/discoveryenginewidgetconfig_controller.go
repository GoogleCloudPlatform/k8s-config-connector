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
// krm.group: discoveryengine.cnrm.cloud.google.com
// krm.version: v1alpha1
// krm.kind: DiscoveryEngineWidgetConfig

package discoveryengine

import (
	"context"
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/discoveryengine/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/discoveryengine/apiv1"
	discoveryenginepb "cloud.google.com/go/discoveryengine/apiv1/discoveryenginepb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.DiscoveryEngineWidgetConfigGVK, NewWidgetConfigModel)
}

func NewWidgetConfigModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelWidgetConfig{config: *config}, nil
}

var _ directbase.Model = &modelWidgetConfig{}

type modelWidgetConfig struct {
	config config.ControllerConfig
}

func (m *modelWidgetConfig) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.DiscoveryEngineWidgetConfig{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.DiscoveryEngineWidgetConfigIdentity)

	if err := common.NormalizeGreenfieldReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	// Get discoveryengine GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &WidgetConfigAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelWidgetConfig) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//discoveryengine.googleapis.com/") {
		return nil, nil
	}
	id := &krm.DiscoveryEngineWidgetConfigIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &WidgetConfigAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}
func (m *modelWidgetConfig) client(ctx context.Context) (*gcp.WidgetConfigClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewWidgetConfigRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building WidgetConfig client: %w", err)
	}
	return gcpClient, err
}

type WidgetConfigAdapter struct {
	id        *krm.DiscoveryEngineWidgetConfigIdentity
	gcpClient *gcp.WidgetConfigClient
	desired   *krm.DiscoveryEngineWidgetConfig
	actual    *discoveryenginepb.WidgetConfig
}

var _ directbase.Adapter = &WidgetConfigAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *WidgetConfigAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting WidgetConfig", "name", a.id)

	req := &discoveryenginepb.GetWidgetConfigRequest{
		Name: a.id.String(),
	}
	widgetconfigpb, err := a.gcpClient.GetWidgetConfig(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting WidgetConfig %q: %w", a.id, err)
	}

	a.actual = widgetconfigpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *WidgetConfigAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating (updating singleton) WidgetConfig", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := DiscoveryEngineWidgetConfigSpec_v1alpha1_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	resource.Name = a.id.String()

	req := &discoveryenginepb.UpdateWidgetConfigRequest{
		WidgetConfig: resource,
	}
	created, err := a.gcpClient.UpdateWidgetConfig(ctx, req)
	if err != nil {
		return fmt.Errorf("creating WidgetConfig %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created WidgetConfig", "name", a.id)

	return a.setStatus(ctx, createOp, created)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *WidgetConfigAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating WidgetConfig", "name", a.id)

	normalize := func(ctx context.Context, p *discoveryenginepb.WidgetConfig) error {
		p.SolutionType = 0
		return nil
	}
	resource, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		DiscoveryEngineWidgetConfigSpec_v1alpha1_FromProto,
		DiscoveryEngineWidgetConfigSpec_v1alpha1_ToProto,
		normalize,
	)
	if err != nil {
		return err
	}
	actual := a.actual
	if !report.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id)
	} else {
		report.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, report)
		resource.Name = a.id.String()
		req := &discoveryenginepb.UpdateWidgetConfigRequest{
			UpdateMask:   updateMask,
			WidgetConfig: resource,
		}
		updated, err := a.gcpClient.UpdateWidgetConfig(ctx, req)
		if err != nil {
			return fmt.Errorf("updating WidgetConfig %s: %w", a.id, err)
		}
		actual = updated
		log.V(2).Info("successfully updated WidgetConfig", "name", a.id)
	}

	return a.setStatus(ctx, updateOp, actual)
}

func (a *WidgetConfigAdapter) setStatus(ctx context.Context, op directbase.Operation, actual *discoveryenginepb.WidgetConfig) error {
	status := &krm.DiscoveryEngineWidgetConfigStatus{}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	mapCtx := &direct.MapContext{}
	status.ObservedState = DiscoveryEngineWidgetConfigObservedState_v1alpha1_FromProto(mapCtx, actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return op.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *WidgetConfigAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.DiscoveryEngineWidgetConfig{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(DiscoveryEngineWidgetConfigSpec_v1alpha1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ParentRef = &krm.DiscoveryEngineDataStoreRef{
		External: a.id.ParentString(),
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.DiscoveryEngineWidgetConfigGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *WidgetConfigAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deletion not supported for WidgetConfig, removing finalizer without deleting GCP resource", "name", a.id)
	return true, nil
}
