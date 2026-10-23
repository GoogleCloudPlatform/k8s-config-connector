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

// +generated:controller
// krm.group: retail.cnrm.cloud.google.com
// krm.version: v1alpha1
// krm.kind: RetailServingConfig

package retail

import (
	"context"
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/retail/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/retail/apiv2"
	retailpb "cloud.google.com/go/retail/apiv2/retailpb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.RetailServingConfigGVK, NewServingConfigModel)
}

func NewServingConfigModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelServingConfig{config: *config}, nil
}

var _ directbase.Model = &modelServingConfig{}

type modelServingConfig struct {
	config config.ControllerConfig
}

func (m *modelServingConfig) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.RetailServingConfig{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.RetailServingConfigIdentity)

	// Get retail GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &ServingConfigAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelServingConfig) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//retail.googleapis.com/") {
		return nil, nil
	}
	id := &krm.RetailServingConfigIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &ServingConfigAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}
func (m *modelServingConfig) client(ctx context.Context) (*gcp.ServingConfigClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewServingConfigRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building ServingConfig client: %w", err)
	}
	return gcpClient, err
}

type ServingConfigAdapter struct {
	id        *krm.RetailServingConfigIdentity
	gcpClient *gcp.ServingConfigClient
	desired   *krm.RetailServingConfig
	actual    *retailpb.ServingConfig
}

var _ directbase.Adapter = &ServingConfigAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *ServingConfigAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ServingConfig", "name", a.id)

	req := &retailpb.GetServingConfigRequest{
		Name: a.id.String(),
	}
	servingconfigpb, err := a.gcpClient.GetServingConfig(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ServingConfig %q: %w", a.id, err)
	}

	a.actual = servingconfigpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *ServingConfigAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating ServingConfig", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := RetailServingConfigSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &retailpb.CreateServingConfigRequest{
		Parent:          a.id.ParentString(),
		ServingConfigId: a.id.ServingConfig,
		ServingConfig:   resource,
	}
	_, err := a.gcpClient.CreateServingConfig(ctx, req)
	if err != nil {
		return fmt.Errorf("creating ServingConfig %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created ServingConfig", "name", a.id)

	status := &krm.RetailServingConfigStatus{}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *ServingConfigAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating ServingConfig", "name", a.id)

	resource, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		RetailServingConfigSpec_FromProto,
		RetailServingConfigSpec_ToProto,
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
		req := &retailpb.UpdateServingConfigRequest{
			UpdateMask:    updateMask,
			ServingConfig: resource,
		}
		_, err = a.gcpClient.UpdateServingConfig(ctx, req)
		if err != nil {
			return fmt.Errorf("updating ServingConfig %s: %w", a.id, err)
		}
		log.V(2).Info("successfully updated ServingConfig", "name", a.id)
	}

	status := &krm.RetailServingConfigStatus{}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *ServingConfigAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.RetailServingConfig{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(RetailServingConfigSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.RetailServingConfigGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *ServingConfigAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ServingConfig", "name", a.id)

	req := &retailpb.DeleteServingConfigRequest{
		Name: a.id.String(),
	}
	err := a.gcpClient.DeleteServingConfig(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent ServingConfig, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting ServingConfig %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted ServingConfig", "name", a.id)
	return true, nil
}
