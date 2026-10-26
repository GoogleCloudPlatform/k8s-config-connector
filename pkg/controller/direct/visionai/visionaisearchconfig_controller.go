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
// krm.group: visionai.cnrm.cloud.google.com
// krm.version: v1alpha1
// krm.kind: VisionAISearchConfig

package visionai

import (
	"context"
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/visionai/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/visionai/apiv1"
	visionaipb "cloud.google.com/go/visionai/apiv1/visionaipb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.VisionAISearchConfigGVK, NewSearchConfigModel)
}

func NewSearchConfigModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelSearchConfig{config: *config}, nil
}

var _ directbase.Model = &modelSearchConfig{}

type modelSearchConfig struct {
	config config.ControllerConfig
}

func (m *modelSearchConfig) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.VisionAISearchConfig{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.VisionAISearchConfigIdentity)

	// Get visionai GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &SearchConfigAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelSearchConfig) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//visionai.googleapis.com/") {
		return nil, nil
	}
	id := &krm.VisionAISearchConfigIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &SearchConfigAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}
func (m *modelSearchConfig) client(ctx context.Context) (*gcp.WarehouseClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewWarehouseRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building SearchConfig client: %w", err)
	}
	return gcpClient, err
}

type SearchConfigAdapter struct {
	id        *krm.VisionAISearchConfigIdentity
	gcpClient *gcp.WarehouseClient
	desired   *krm.VisionAISearchConfig
	actual    *visionaipb.SearchConfig
}

var _ directbase.Adapter = &SearchConfigAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *SearchConfigAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting SearchConfig", "name", a.id)

	req := &visionaipb.GetSearchConfigRequest{
		Name: a.id.String(),
	}
	searchconfigpb, err := a.gcpClient.GetSearchConfig(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting SearchConfig %q: %w", a.id, err)
	}

	a.actual = searchconfigpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *SearchConfigAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating SearchConfig", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := VisionAISearchConfigSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &visionaipb.CreateSearchConfigRequest{
		Parent:         a.id.ParentString(),
		SearchConfigId: a.id.SearchConfig,
		SearchConfig:   resource,
	}
	_, err := a.gcpClient.CreateSearchConfig(ctx, req)
	if err != nil {
		return fmt.Errorf("creating SearchConfig %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created SearchConfig", "name", a.id)

	status := &krm.VisionAISearchConfigStatus{}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *SearchConfigAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating SearchConfig", "name", a.id)

	resource, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		VisionAISearchConfigSpec_FromProto,
		VisionAISearchConfigSpec_ToProto,
		nil,
	)
	if err != nil {
		return err
	}
	if !report.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id)
	} else {
		report.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, report)
		resource.Name = a.id.String()
		req := &visionaipb.UpdateSearchConfigRequest{
			UpdateMask:   updateMask,
			SearchConfig: resource,
		}
		_, err = a.gcpClient.UpdateSearchConfig(ctx, req)
		if err != nil {
			return fmt.Errorf("updating SearchConfig %s: %w", a.id, err)
		}
		log.V(2).Info("successfully updated SearchConfig", "name", a.id)
	}

	status := &krm.VisionAISearchConfigStatus{}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *SearchConfigAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.VisionAISearchConfig{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(VisionAISearchConfigSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.VisionAISearchConfigGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *SearchConfigAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting SearchConfig", "name", a.id)

	req := &visionaipb.DeleteSearchConfigRequest{
		Name: a.id.String(),
	}
	err := a.gcpClient.DeleteSearchConfig(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent SearchConfig, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting SearchConfig %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted SearchConfig", "name", a.id)
	return true, nil
}
