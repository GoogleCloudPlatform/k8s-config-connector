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
// krm.group: cloudchannel.cnrm.cloud.google.com
// krm.version: v1alpha1
// krm.kind: CloudchannelCustomerRepricingConfig

package cloudchannel

import (
	"context"
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/cloudchannel/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/channel/apiv1"
	channelpb "cloud.google.com/go/channel/apiv1/channelpb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.CloudchannelCustomerRepricingConfigGVK, NewCustomerRepricingConfigModel)
}

func NewCustomerRepricingConfigModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelCustomerRepricingConfig{config: *config}, nil
}

var _ directbase.Model = &modelCustomerRepricingConfig{}

type modelCustomerRepricingConfig struct {
	config config.ControllerConfig
}

func (m *modelCustomerRepricingConfig) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.CloudchannelCustomerRepricingConfig{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.CloudchannelCustomerRepricingConfigIdentity)

	// Get cloudchannel GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &CustomerRepricingConfigAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelCustomerRepricingConfig) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//cloudchannel.googleapis.com/") {
		return nil, nil
	}
	id := &krm.CloudchannelCustomerRepricingConfigIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &CustomerRepricingConfigAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}
func (m *modelCustomerRepricingConfig) client(ctx context.Context) (*gcp.CloudChannelClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewCloudChannelRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building CustomerRepricingConfig client: %w", err)
	}
	return gcpClient, err
}

type CustomerRepricingConfigAdapter struct {
	id        *krm.CloudchannelCustomerRepricingConfigIdentity
	gcpClient *gcp.CloudChannelClient
	desired   *krm.CloudchannelCustomerRepricingConfig
	actual    *channelpb.CustomerRepricingConfig
}

var _ directbase.Adapter = &CustomerRepricingConfigAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *CustomerRepricingConfigAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting CustomerRepricingConfig", "name", a.id)
	if !a.id.HasIdentitySpecified() {
		return false, nil
	}

	req := &channelpb.GetCustomerRepricingConfigRequest{
		Name: a.id.String(),
	}
	customerrepricingconfigpb, err := a.gcpClient.GetCustomerRepricingConfig(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting CustomerRepricingConfig %q: %w", a.id, err)
	}

	a.actual = customerrepricingconfigpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *CustomerRepricingConfigAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating CustomerRepricingConfig", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := CloudchannelCustomerRepricingConfigSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &channelpb.CreateCustomerRepricingConfigRequest{
		Parent:                  a.id.ParentString(),
		CustomerRepricingConfig: resource,
	}
	created, err := a.gcpClient.CreateCustomerRepricingConfig(ctx, req)
	if err != nil {
		return fmt.Errorf("creating CustomerRepricingConfig %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created CustomerRepricingConfig", "name", a.id)

	status := &krm.CloudchannelCustomerRepricingConfigStatus{}
	status.ObservedState = CloudchannelCustomerRepricingConfigObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	parsedID := &krm.CloudchannelCustomerRepricingConfigIdentity{}
	if err := parsedID.FromExternal(created.Name); err != nil {
		return err
	}
	a.id.CustomerRepricingConfig = parsedID.CustomerRepricingConfig
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *CustomerRepricingConfigAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating CustomerRepricingConfig", "name", a.id)

	resource, report, _, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		CloudchannelCustomerRepricingConfigSpec_FromProto,
		CloudchannelCustomerRepricingConfigSpec_ToProto,
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
		req := &channelpb.UpdateCustomerRepricingConfigRequest{
			CustomerRepricingConfig: resource,
		}
		updated, err = a.gcpClient.UpdateCustomerRepricingConfig(ctx, req)
		if err != nil {
			return fmt.Errorf("updating CustomerRepricingConfig %s: %w", a.id, err)
		}
		log.V(2).Info("successfully updated CustomerRepricingConfig", "name", a.id)
	}

	status := &krm.CloudchannelCustomerRepricingConfigStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = CloudchannelCustomerRepricingConfigObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *CustomerRepricingConfigAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.CloudchannelCustomerRepricingConfig{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(CloudchannelCustomerRepricingConfigSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.CloudchannelCustomerRepricingConfigGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *CustomerRepricingConfigAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting CustomerRepricingConfig", "name", a.id)

	req := &channelpb.DeleteCustomerRepricingConfigRequest{
		Name: a.id.String(),
	}
	err := a.gcpClient.DeleteCustomerRepricingConfig(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent CustomerRepricingConfig, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting CustomerRepricingConfig %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted CustomerRepricingConfig", "name", a.id)
	return true, nil
}
