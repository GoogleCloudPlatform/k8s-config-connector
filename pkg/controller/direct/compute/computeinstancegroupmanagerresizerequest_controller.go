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
// krm.kind: ComputeInstanceGroupManagerResizeRequest

package compute

import (
	"context"
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	compute "cloud.google.com/go/compute/apiv1"
	computepb "cloud.google.com/go/compute/apiv1/computepb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.ComputeInstanceGroupManagerResizeRequestGVK, NewInstanceGroupManagerResizeRequestModel)
}

func NewInstanceGroupManagerResizeRequestModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelInstanceGroupManagerResizeRequest{config: *config}, nil
}

var _ directbase.Model = &modelInstanceGroupManagerResizeRequest{}

type modelInstanceGroupManagerResizeRequest struct {
	config config.ControllerConfig
}

func (m *modelInstanceGroupManagerResizeRequest) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ComputeInstanceGroupManagerResizeRequest{}
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
	id := i.(*krm.ComputeInstanceGroupManagerResizeRequestIdentity)

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &InstanceGroupManagerResizeRequestAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelInstanceGroupManagerResizeRequest) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//compute.googleapis.com/") {
		return nil, nil
	}
	id := &krm.ComputeInstanceGroupManagerResizeRequestIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &InstanceGroupManagerResizeRequestAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}

func (m *modelInstanceGroupManagerResizeRequest) client(ctx context.Context) (*compute.InstanceGroupManagerResizeRequestsClient, error) {
	gcpClient, err := newGCPClient(&m.config)
	if err != nil {
		return nil, fmt.Errorf("building gcp client: %w", err)
	}
	return gcpClient.newInstanceGroupManagerResizeRequestsClient(ctx)
}

type InstanceGroupManagerResizeRequestAdapter struct {
	id        *krm.ComputeInstanceGroupManagerResizeRequestIdentity
	gcpClient *compute.InstanceGroupManagerResizeRequestsClient
	desired   *krm.ComputeInstanceGroupManagerResizeRequest
	actual    *computepb.InstanceGroupManagerResizeRequest
}

var _ directbase.Adapter = &InstanceGroupManagerResizeRequestAdapter{}

// Find retrieves the GCP resource.
func (a *InstanceGroupManagerResizeRequestAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting InstanceGroupManagerResizeRequest", "name", a.id)

	req := &computepb.GetInstanceGroupManagerResizeRequestRequest{
		Project:              a.id.Project,
		Zone:                 a.id.Zone,
		InstanceGroupManager: a.id.InstanceGroupManager,
		ResizeRequest:        a.id.ResizeRequest,
	}
	actual, err := a.gcpClient.Get(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting InstanceGroupManagerResizeRequest %q: %w", a.id, err)
	}
	a.actual = actual
	return true, nil
}

// Create creates the resource in GCP based on `spec` and updates the Config Connector object `status`.
func (a *InstanceGroupManagerResizeRequestAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating InstanceGroupManagerResizeRequest", "name", a.id)

	mapCtx := &direct.MapContext{}
	resource := ComputeInstanceGroupManagerResizeRequestSpec_v1alpha1_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	resource.Name = direct.LazyPtr(a.id.ResizeRequest)

	req := &computepb.InsertInstanceGroupManagerResizeRequestRequest{
		Project:              a.id.Project,
		Zone:                 a.id.Zone,
		InstanceGroupManager: a.id.InstanceGroupManager,
		InstanceGroupManagerResizeRequestResource: resource,
	}
	op, err := a.gcpClient.Insert(ctx, req)
	if err != nil {
		return fmt.Errorf("creating InstanceGroupManagerResizeRequest %s: %w", a.id, err)
	}
	if err := op.Wait(ctx); err != nil {
		return fmt.Errorf("waiting for InstanceGroupManagerResizeRequest %s creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created InstanceGroupManagerResizeRequest", "name", a.id)

	actual, err := a.gcpClient.Get(ctx, &computepb.GetInstanceGroupManagerResizeRequestRequest{
		Project:              a.id.Project,
		Zone:                 a.id.Zone,
		InstanceGroupManager: a.id.InstanceGroupManager,
		ResizeRequest:        a.id.ResizeRequest,
	})
	if err != nil {
		return fmt.Errorf("getting created InstanceGroupManagerResizeRequest %s: %w", a.id, err)
	}
	a.actual = actual

	status := &krm.ComputeInstanceGroupManagerResizeRequestStatus{}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	status.ObservedState = ComputeInstanceGroupManagerResizeRequestObservedState_v1alpha1_FromProto(mapCtx, actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and updates the Config Connector object `status`.
func (a *InstanceGroupManagerResizeRequestAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating InstanceGroupManagerResizeRequest", "name", a.id)

	_, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		ComputeInstanceGroupManagerResizeRequestSpec_v1alpha1_FromProto,
		ComputeInstanceGroupManagerResizeRequestSpec_v1alpha1_ToProto,
		nil,
	)
	if err != nil {
		return err
	}
	if report.HasDiff() {
		report.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, report)
		return fmt.Errorf("updating ComputeInstanceGroupManagerResizeRequest is not supported, fields modified: %v", updateMask.Paths)
	}

	mapCtx := &direct.MapContext{}
	status := &krm.ComputeInstanceGroupManagerResizeRequestStatus{}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	status.ObservedState = ComputeInstanceGroupManagerResizeRequestObservedState_v1alpha1_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *InstanceGroupManagerResizeRequestAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ComputeInstanceGroupManagerResizeRequest{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(ComputeInstanceGroupManagerResizeRequestSpec_v1alpha1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.ComputeInstanceGroupManagerResizeRequestGVK)

	return u, nil
}

// Delete deletes the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *InstanceGroupManagerResizeRequestAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting InstanceGroupManagerResizeRequest", "name", a.id)

	req := &computepb.DeleteInstanceGroupManagerResizeRequestRequest{
		Project:              a.id.Project,
		Zone:                 a.id.Zone,
		InstanceGroupManager: a.id.InstanceGroupManager,
		ResizeRequest:        a.id.ResizeRequest,
	}
	op, err := a.gcpClient.Delete(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting InstanceGroupManagerResizeRequest %s: %w", a.id, err)
	}
	if err := op.Wait(ctx); err != nil {
		return false, fmt.Errorf("waiting for InstanceGroupManagerResizeRequest %s deletion: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted InstanceGroupManagerResizeRequest", "name", a.id)
	return true, nil
}
