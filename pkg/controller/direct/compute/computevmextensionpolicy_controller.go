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
// krm.kind: ComputeVMExtensionPolicy

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
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.ComputeVMExtensionPolicyGVK, NewVmExtensionPolicyModel)
}

func NewVmExtensionPolicyModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelVmExtensionPolicy{config: *config}, nil
}

var _ directbase.Model = &modelVmExtensionPolicy{}

type modelVmExtensionPolicy struct {
	config config.ControllerConfig
}

func (m *modelVmExtensionPolicy) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ComputeVMExtensionPolicy{}
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
	id := i.(*krm.ComputeVMExtensionPolicyIdentity)

	// Get compute GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &VmExtensionPolicyAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelVmExtensionPolicy) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//compute.googleapis.com/") {
		return nil, nil
	}
	id := &krm.ComputeVMExtensionPolicyIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &VmExtensionPolicyAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}

func (m *modelVmExtensionPolicy) client(ctx context.Context) (*gcp.ZoneVmExtensionPoliciesClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewZoneVmExtensionPoliciesRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building ZoneVmExtensionPolicies client: %w", err)
	}
	return gcpClient, err
}

type VmExtensionPolicyAdapter struct {
	id        *krm.ComputeVMExtensionPolicyIdentity
	gcpClient *gcp.ZoneVmExtensionPoliciesClient
	desired   *krm.ComputeVMExtensionPolicy
	actual    *computepb.VmExtensionPolicy
}

var _ directbase.Adapter = &VmExtensionPolicyAdapter{}

func (a *VmExtensionPolicyAdapter) get(ctx context.Context) (*computepb.VmExtensionPolicy, error) {
	req := &computepb.GetZoneVmExtensionPolicyRequest{
		Project:           a.id.Project,
		Zone:              a.id.Zone,
		VmExtensionPolicy: a.id.Vmextensionpolicy,
	}
	return a.gcpClient.Get(ctx, req)
}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *VmExtensionPolicyAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ComputeVMExtensionPolicy", "name", a.id)

	actual, err := a.get(ctx)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ComputeVMExtensionPolicy %q: %w", a.id, err)
	}
	a.actual = actual
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *VmExtensionPolicyAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating ComputeVMExtensionPolicy", "name", a.id)

	mapCtx := &direct.MapContext{}
	resource := ComputeVMExtensionPolicySpec_v1alpha1_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	resource.Name = direct.LazyPtr(a.id.Vmextensionpolicy)

	req := &computepb.InsertZoneVmExtensionPolicyRequest{
		Project:                   a.id.Project,
		Zone:                      a.id.Zone,
		VmExtensionPolicyResource: resource,
	}
	op, err := a.gcpClient.Insert(ctx, req)
	if err != nil {
		return fmt.Errorf("creating ComputeVMExtensionPolicy %s: %w", a.id, err)
	}
	if err := op.Wait(ctx); err != nil {
		return fmt.Errorf("waiting for ComputeVMExtensionPolicy %s creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created ComputeVMExtensionPolicy", "name", a.id)

	created, err := a.get(ctx)
	if err != nil {
		return fmt.Errorf("getting ComputeVMExtensionPolicy %s after creation: %w", a.id, err)
	}

	return a.setStatus(ctx, createOp, created)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *VmExtensionPolicyAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating ComputeVMExtensionPolicy", "name", a.id)

	_, report, _, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		ComputeVMExtensionPolicySpec_v1alpha1_FromProto,
		ComputeVMExtensionPolicySpec_v1alpha1_ToProto,
		nil,
	)
	if err != nil {
		return err
	}
	if !report.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id.String())
		return a.setStatus(ctx, updateOp, a.actual)
	}

	report.Object = updateOp.GetUnstructured()
	structuredreporting.ReportDiff(ctx, report)

	mapCtx := &direct.MapContext{}
	resource := ComputeVMExtensionPolicySpec_v1alpha1_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	resource.Name = direct.LazyPtr(a.id.Vmextensionpolicy)

	req := &computepb.UpdateZoneVmExtensionPolicyRequest{
		Project:                   a.id.Project,
		Zone:                      a.id.Zone,
		VmExtensionPolicy:         a.id.Vmextensionpolicy,
		VmExtensionPolicyResource: resource,
	}
	op, err := a.gcpClient.Update(ctx, req)
	if err != nil {
		return fmt.Errorf("updating ComputeVMExtensionPolicy %s: %w", a.id, err)
	}
	if err := op.Wait(ctx); err != nil {
		return fmt.Errorf("waiting for ComputeVMExtensionPolicy %s update: %w", a.id, err)
	}
	log.V(2).Info("successfully updated ComputeVMExtensionPolicy", "name", a.id)

	updated, err := a.get(ctx)
	if err != nil {
		return fmt.Errorf("getting ComputeVMExtensionPolicy %s after update: %w", a.id, err)
	}

	return a.setStatus(ctx, updateOp, updated)
}

func (a *VmExtensionPolicyAdapter) setStatus(ctx context.Context, op directbase.Operation, actual *computepb.VmExtensionPolicy) error {
	status := &krm.ComputeVMExtensionPolicyStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = ComputeVMExtensionPolicyObservedState_v1alpha1_FromProto(mapCtx, actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return op.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *VmExtensionPolicyAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ComputeVMExtensionPolicy{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(ComputeVMExtensionPolicySpec_v1alpha1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.ProjectRef = &refsv1beta1.ProjectRef{External: a.id.Project}
	obj.Spec.Location = a.id.Zone

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.Vmextensionpolicy)
	u.SetGroupVersionKind(krm.ComputeVMExtensionPolicyGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *VmExtensionPolicyAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ComputeVMExtensionPolicy", "name", a.id)

	req := &computepb.DeleteZoneVmExtensionPolicyRequest{
		Project:           a.id.Project,
		Zone:              a.id.Zone,
		VmExtensionPolicy: a.id.Vmextensionpolicy,
	}
	op, err := a.gcpClient.Delete(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting ComputeVMExtensionPolicy %s: %w", a.id, err)
	}
	if err := op.Wait(ctx); err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("waiting for ComputeVMExtensionPolicy %s deletion: %w", a.id, err)
	}
	return true, nil
}
