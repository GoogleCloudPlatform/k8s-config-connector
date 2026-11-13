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
// krm.group: compute.cnrm.cloud.google.com
// krm.version: v1alpha1
// krm.kind: ComputeNotificationEndpoint

package compute

import (
	"context"
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1alpha1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/compute/apiv1"
	computepb "cloud.google.com/go/compute/apiv1/computepb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.ComputeNotificationEndpointGVK, NewNotificationEndpointModel)
}

func NewNotificationEndpointModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelNotificationEndpoint{config: *config}, nil
}

var _ directbase.Model = &modelNotificationEndpoint{}

type modelNotificationEndpoint struct {
	config config.ControllerConfig
}

func (m *modelNotificationEndpoint) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ComputeNotificationEndpoint{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if err := common.NormalizeGreenfieldReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.ComputeNotificationEndpointIdentity)

	gcpClient, err := newGCPClient(&m.config)
	if err != nil {
		return nil, fmt.Errorf("building gcp client: %w", err)
	}
	client, err := gcpClient.newRegionNotificationEndpointsClient(ctx)
	if err != nil {
		return nil, err
	}

	return &NotificationEndpointAdapter{
		id:        id,
		gcpClient: client,
		desired:   obj,
	}, nil
}

func (m *modelNotificationEndpoint) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//compute.googleapis.com/") {
		return nil, nil
	}
	id := &krm.ComputeNotificationEndpointIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}

	gcpClient, err := newGCPClient(&m.config)
	if err != nil {
		return nil, fmt.Errorf("building gcp client: %w", err)
	}
	client, err := gcpClient.newRegionNotificationEndpointsClient(ctx)
	if err != nil {
		return nil, err
	}

	return &NotificationEndpointAdapter{
		id:        id,
		gcpClient: client,
	}, nil
}

type NotificationEndpointAdapter struct {
	id        *krm.ComputeNotificationEndpointIdentity
	gcpClient *gcp.RegionNotificationEndpointsClient
	desired   *krm.ComputeNotificationEndpoint
	actual    *computepb.NotificationEndpoint
}

var _ directbase.Adapter = &NotificationEndpointAdapter{}

func (a *NotificationEndpointAdapter) get(ctx context.Context) (*computepb.NotificationEndpoint, error) {
	req := &computepb.GetRegionNotificationEndpointRequest{
		Project:              a.id.Project,
		Region:               a.id.Region,
		NotificationEndpoint: a.id.Notificationendpoint,
	}
	return a.gcpClient.Get(ctx, req)
}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *NotificationEndpointAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting NotificationEndpoint", "name", a.id)

	actual, err := a.get(ctx)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ComputeNotificationEndpoint %q: %w", a.id, err)
	}
	a.actual = actual
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *NotificationEndpointAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating NotificationEndpoint", "name", a.id)

	mapCtx := &direct.MapContext{}
	resource := ComputeNotificationEndpointSpec_v1alpha1_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	resource.Name = direct.LazyPtr(a.id.Notificationendpoint)

	req := &computepb.InsertRegionNotificationEndpointRequest{
		Project:                      a.id.Project,
		Region:                       a.id.Region,
		NotificationEndpointResource: resource,
	}
	op, err := a.gcpClient.Insert(ctx, req)
	if err != nil {
		return fmt.Errorf("creating ComputeNotificationEndpoint %s: %w", a.id, err)
	}
	if err := op.Wait(ctx); err != nil {
		return fmt.Errorf("waiting for ComputeNotificationEndpoint %s creation: %w", a.id, err)
	}

	created, err := a.get(ctx)
	if err != nil {
		return fmt.Errorf("getting ComputeNotificationEndpoint %s after creation: %w", a.id, err)
	}

	return a.setStatus(ctx, createOp, created)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *NotificationEndpointAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating NotificationEndpoint", "name", a.id)

	_, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		ComputeNotificationEndpointSpec_v1alpha1_FromProto,
		ComputeNotificationEndpointSpec_v1alpha1_ToProto,
		nil,
	)
	if err != nil {
		return err
	}
	if report.HasDiff() {
		report.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, report)
		return fmt.Errorf("updating ComputeNotificationEndpoint is not supported, fields modified: %v", updateMask.Paths)
	}

	return a.setStatus(ctx, updateOp, a.actual)
}

func (a *NotificationEndpointAdapter) setStatus(ctx context.Context, op directbase.Operation, actual *computepb.NotificationEndpoint) error {
	status := &krm.ComputeNotificationEndpointStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = ComputeNotificationEndpointObservedState_v1alpha1_FromProto(mapCtx, actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return op.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *NotificationEndpointAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ComputeNotificationEndpoint{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(ComputeNotificationEndpointSpec_v1alpha1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.ProjectRef = &refsv1beta1.ProjectRef{External: a.id.Project}
	obj.Spec.Location = a.id.Region

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.Notificationendpoint)
	u.SetGroupVersionKind(krm.ComputeNotificationEndpointGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *NotificationEndpointAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting NotificationEndpoint", "name", a.id)

	req := &computepb.DeleteRegionNotificationEndpointRequest{
		Project:              a.id.Project,
		Region:               a.id.Region,
		NotificationEndpoint: a.id.Notificationendpoint,
	}
	op, err := a.gcpClient.Delete(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting ComputeNotificationEndpoint %s: %w", a.id, err)
	}
	if err := op.Wait(ctx); err != nil {
		return false, fmt.Errorf("waiting for ComputeNotificationEndpoint %s deletion: %w", a.id, err)
	}

	log.V(2).Info("successfully deleted ComputeNotificationEndpoint", "name", a.id)
	return true, nil
}
