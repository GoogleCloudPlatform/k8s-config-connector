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
// krm.group: iap.cnrm.cloud.google.com
// krm.version: v1alpha1
// krm.kind: IAPTunnelDestGroup

package iap

import (
	"context"
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/iap/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/iap/apiv1"
	iappb "cloud.google.com/go/iap/apiv1/iappb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.IAPTunnelDestGroupGVK, NewTunnelDestGroupModel)
}

func NewTunnelDestGroupModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelTunnelDestGroup{config: *config}, nil
}

var _ directbase.Model = &modelTunnelDestGroup{}

type modelTunnelDestGroup struct {
	config config.ControllerConfig
}

func (m *modelTunnelDestGroup) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.IAPTunnelDestGroup{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.IAPTunnelDestGroupIdentity)

	// Get iap GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &TunnelDestGroupAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelTunnelDestGroup) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//iap.googleapis.com/") {
		return nil, nil
	}
	id := &krm.IAPTunnelDestGroupIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &TunnelDestGroupAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}
func (m *modelTunnelDestGroup) client(ctx context.Context) (*gcp.IdentityAwareProxyAdminClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewIdentityAwareProxyAdminRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building TunnelDestGroup client: %w", err)
	}
	return gcpClient, err
}

type TunnelDestGroupAdapter struct {
	id        *krm.IAPTunnelDestGroupIdentity
	gcpClient *gcp.IdentityAwareProxyAdminClient
	desired   *krm.IAPTunnelDestGroup
	actual    *iappb.TunnelDestGroup
}

var _ directbase.Adapter = &TunnelDestGroupAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *TunnelDestGroupAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting TunnelDestGroup", "name", a.id)

	req := &iappb.GetTunnelDestGroupRequest{
		Name: a.id.String(),
	}
	tunneldestgrouppb, err := a.gcpClient.GetTunnelDestGroup(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting TunnelDestGroup %q: %w", a.id, err)
	}

	a.actual = tunneldestgrouppb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *TunnelDestGroupAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating TunnelDestGroup", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := IAPTunnelDestGroupSpec_v1alpha1_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &iappb.CreateTunnelDestGroupRequest{
		Parent:            a.id.ParentString(),
		TunnelDestGroupId: a.id.DestGroup,
		TunnelDestGroup:   resource,
	}
	_, err := a.gcpClient.CreateTunnelDestGroup(ctx, req)
	if err != nil {
		return fmt.Errorf("creating TunnelDestGroup %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created TunnelDestGroup", "name", a.id)

	status := &krm.IAPTunnelDestGroupStatus{}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *TunnelDestGroupAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating TunnelDestGroup", "name", a.id)

	resource, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		IAPTunnelDestGroupSpec_v1alpha1_FromProto,
		IAPTunnelDestGroupSpec_v1alpha1_ToProto,
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
		req := &iappb.UpdateTunnelDestGroupRequest{
			UpdateMask:      updateMask,
			TunnelDestGroup: resource,
		}
		_, err = a.gcpClient.UpdateTunnelDestGroup(ctx, req)
		if err != nil {
			return fmt.Errorf("updating TunnelDestGroup %s: %w", a.id, err)
		}
		log.V(2).Info("successfully updated TunnelDestGroup", "name", a.id)
	}

	status := &krm.IAPTunnelDestGroupStatus{}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *TunnelDestGroupAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.IAPTunnelDestGroup{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(IAPTunnelDestGroupSpec_v1alpha1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.IAPTunnelDestGroupGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *TunnelDestGroupAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting TunnelDestGroup", "name", a.id)

	req := &iappb.DeleteTunnelDestGroupRequest{
		Name: a.id.String(),
	}
	err := a.gcpClient.DeleteTunnelDestGroup(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent TunnelDestGroup, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting TunnelDestGroup %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted TunnelDestGroup", "name", a.id)
	return true, nil
}
