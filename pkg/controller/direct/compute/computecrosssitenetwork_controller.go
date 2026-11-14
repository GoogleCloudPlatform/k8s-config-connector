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
// krm.kind: ComputeCrossSiteNetwork

package compute

import (
	"context"
	"fmt"
	"strings"

	compute "cloud.google.com/go/compute/apiv1"
	computepb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1alpha1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.ComputeCrossSiteNetworkGVK, NewCrossSiteNetworkModel)
}

func NewCrossSiteNetworkModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelCrossSiteNetwork{config: *config}, nil
}

var _ directbase.Model = &modelCrossSiteNetwork{}

type modelCrossSiteNetwork struct {
	config config.ControllerConfig
}

func (m *modelCrossSiteNetwork) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ComputeCrossSiteNetwork{}
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
	id := i.(*krm.ComputeCrossSiteNetworkIdentity)

	gcpClient, err := newGCPClient(&m.config)
	if err != nil {
		return nil, fmt.Errorf("building gcp client: %w", err)
	}
	crossSiteNetworksClient, err := gcpClient.newCrossSiteNetworksClient(ctx)
	if err != nil {
		return nil, err
	}

	return &CrossSiteNetworkAdapter{
		id:        id,
		gcpClient: crossSiteNetworksClient,
		desired:   obj,
	}, nil
}

func (m *modelCrossSiteNetwork) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//compute.googleapis.com/") {
		return nil, nil
	}
	id := &krm.ComputeCrossSiteNetworkIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}

	gcpClient, err := newGCPClient(&m.config)
	if err != nil {
		return nil, fmt.Errorf("building gcp client: %w", err)
	}
	crossSiteNetworksClient, err := gcpClient.newCrossSiteNetworksClient(ctx)
	if err != nil {
		return nil, err
	}

	return &CrossSiteNetworkAdapter{
		id:        id,
		gcpClient: crossSiteNetworksClient,
	}, nil
}

type CrossSiteNetworkAdapter struct {
	id        *krm.ComputeCrossSiteNetworkIdentity
	gcpClient *compute.CrossSiteNetworksClient
	desired   *krm.ComputeCrossSiteNetwork
	actual    *computepb.CrossSiteNetwork
}

var _ directbase.Adapter = &CrossSiteNetworkAdapter{}

func (a *CrossSiteNetworkAdapter) get(ctx context.Context) (*computepb.CrossSiteNetwork, error) {
	req := &computepb.GetCrossSiteNetworkRequest{
		Project:          a.id.Project,
		CrossSiteNetwork: a.id.CrossSiteNetwork,
	}
	return a.gcpClient.Get(ctx, req)
}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *CrossSiteNetworkAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting CrossSiteNetwork", "name", a.id)

	actual, err := a.get(ctx)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ComputeCrossSiteNetwork %q: %w", a.id, err)
	}
	a.actual = actual
	return true, nil
}

// Create creates the resource in GCP based on `spec` and updates the Config Connector object `status` based on the GCP response.
func (a *CrossSiteNetworkAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating CrossSiteNetwork", "name", a.id)

	mapCtx := &direct.MapContext{}
	resource := ComputeCrossSiteNetworkSpec_v1alpha1_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	resource.Name = direct.LazyPtr(a.id.CrossSiteNetwork)

	req := &computepb.InsertCrossSiteNetworkRequest{
		Project:                  a.id.Project,
		CrossSiteNetworkResource: resource,
	}
	op, err := a.gcpClient.Insert(ctx, req)
	if err != nil {
		return fmt.Errorf("creating ComputeCrossSiteNetwork %s: %w", a.id, err)
	}
	if err := op.Wait(ctx); err != nil {
		return fmt.Errorf("waiting for ComputeCrossSiteNetwork %s creation: %w", a.id, err)
	}

	log.V(2).Info("successfully created CrossSiteNetwork", "name", a.id)

	created, err := a.get(ctx)
	if err != nil {
		return fmt.Errorf("getting ComputeCrossSiteNetwork %s after creation: %w", a.id, err)
	}

	return a.setStatus(ctx, createOp, created)
}

// Update updates the resource in GCP based on `spec` and updates the Config Connector object `status` based on the GCP response.
func (a *CrossSiteNetworkAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating CrossSiteNetwork", "name", a.id)

	resource, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		ComputeCrossSiteNetworkSpec_v1alpha1_FromProto,
		ComputeCrossSiteNetworkSpec_v1alpha1_ToProto,
		nil,
	)
	if err != nil {
		return err
	}
	if !report.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id)
		return a.setStatus(ctx, updateOp, a.actual)
	}

	report.Object = updateOp.GetUnstructured()
	structuredreporting.ReportDiff(ctx, report)

	mask := strings.Join(updateMask.Paths, ",")
	req := &computepb.PatchCrossSiteNetworkRequest{
		Project:                  a.id.Project,
		CrossSiteNetwork:         a.id.CrossSiteNetwork,
		CrossSiteNetworkResource: resource,
		UpdateMask:               direct.LazyPtr(mask),
	}
	op, err := a.gcpClient.Patch(ctx, req)
	if err != nil {
		return fmt.Errorf("updating ComputeCrossSiteNetwork %s: %w", a.id, err)
	}
	if err := op.Wait(ctx); err != nil {
		return fmt.Errorf("waiting for ComputeCrossSiteNetwork %s update: %w", a.id, err)
	}
	log.V(2).Info("successfully updated CrossSiteNetwork", "name", a.id)

	updated, err := a.get(ctx)
	if err != nil {
		return fmt.Errorf("getting ComputeCrossSiteNetwork %s after update: %w", a.id, err)
	}

	return a.setStatus(ctx, updateOp, updated)
}

func (a *CrossSiteNetworkAdapter) setStatus(ctx context.Context, op directbase.Operation, actual *computepb.CrossSiteNetwork) error {
	status := &krm.ComputeCrossSiteNetworkStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = ComputeCrossSiteNetworkObservedState_v1alpha1_FromProto(mapCtx, actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return op.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *CrossSiteNetworkAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ComputeCrossSiteNetwork{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(ComputeCrossSiteNetworkSpec_v1alpha1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.ProjectRef = &refsv1beta1.ProjectRef{External: a.id.Project}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.CrossSiteNetwork)
	u.SetGroupVersionKind(krm.ComputeCrossSiteNetworkGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *CrossSiteNetworkAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting CrossSiteNetwork", "name", a.id)

	req := &computepb.DeleteCrossSiteNetworkRequest{
		Project:          a.id.Project,
		CrossSiteNetwork: a.id.CrossSiteNetwork,
	}
	op, err := a.gcpClient.Delete(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting ComputeCrossSiteNetwork %s: %w", a.id, err)
	}
	if err := op.Wait(ctx); err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("waiting for ComputeCrossSiteNetwork %s deletion: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted CrossSiteNetwork", "name", a.id)
	return true, nil
}
