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
// krm.kind: ComputeStoragePool

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
	registry.RegisterModel(krm.ComputeStoragePoolGVK, NewStoragePoolModel)
}

func NewStoragePoolModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelStoragePool{config: *config}, nil
}

var _ directbase.Model = &modelStoragePool{}

type modelStoragePool struct {
	config config.ControllerConfig
}

func (m *modelStoragePool) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ComputeStoragePool{}
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
	id := i.(*krm.ComputeStoragePoolIdentity)

	// Get compute GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &StoragePoolAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelStoragePool) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//compute.googleapis.com/") {
		return nil, nil
	}
	id := &krm.ComputeStoragePoolIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &StoragePoolAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}

func (m *modelStoragePool) client(ctx context.Context) (*gcp.StoragePoolsClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewStoragePoolsRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building StoragePools client: %w", err)
	}
	return gcpClient, err
}

type StoragePoolAdapter struct {
	id        *krm.ComputeStoragePoolIdentity
	gcpClient *gcp.StoragePoolsClient
	desired   *krm.ComputeStoragePool
	actual    *computepb.StoragePool
}

var _ directbase.Adapter = &StoragePoolAdapter{}

func (a *StoragePoolAdapter) get(ctx context.Context) (*computepb.StoragePool, error) {
	req := &computepb.GetStoragePoolRequest{
		Project:     a.id.Project,
		Zone:        a.id.Zone,
		StoragePool: a.id.StoragePool,
	}
	return a.gcpClient.Get(ctx, req)
}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *StoragePoolAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ComputeStoragePool", "name", a.id)

	actual, err := a.get(ctx)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ComputeStoragePool %q: %w", a.id, err)
	}
	a.actual = actual
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *StoragePoolAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating ComputeStoragePool", "name", a.id)

	mapCtx := &direct.MapContext{}
	resource := ComputeStoragePoolSpec_v1alpha1_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	resource.Name = direct.LazyPtr(a.id.StoragePool)
	if resource.StoragePoolType != nil && !strings.Contains(*resource.StoragePoolType, "/") {
		resource.StoragePoolType = direct.LazyPtr(fmt.Sprintf("projects/%s/zones/%s/storagePoolTypes/%s", a.id.Project, a.id.Zone, *resource.StoragePoolType))
	}

	req := &computepb.InsertStoragePoolRequest{
		Project:             a.id.Project,
		Zone:                a.id.Zone,
		StoragePoolResource: resource,
	}
	op, err := a.gcpClient.Insert(ctx, req)
	if err != nil {
		return fmt.Errorf("creating ComputeStoragePool %s: %w", a.id, err)
	}
	if err := op.Wait(ctx); err != nil {
		return fmt.Errorf("waiting for ComputeStoragePool %s creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created ComputeStoragePool", "name", a.id)

	created, err := a.get(ctx)
	if err != nil {
		return fmt.Errorf("getting ComputeStoragePool %s after creation: %w", a.id, err)
	}

	return a.setStatus(ctx, createOp, created)
}

func trimStoragePoolTypePrefix(t string) string {
	normalized := refsv1beta1.TrimComputeURIPrefix(t)
	parts := strings.Split(normalized, "/")
	if len(parts) == 6 && parts[0] == "projects" && parts[4] == "storagePoolTypes" {
		if parts[2] == "zones" {
			return parts[5]
		}
	}
	return t
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *StoragePoolAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating ComputeStoragePool", "name", a.id)

	normalize := func(ctx context.Context, pb *computepb.StoragePool) error {
		if pb.StoragePoolType != nil {
			pb.StoragePoolType = direct.LazyPtr(trimStoragePoolTypePrefix(*pb.StoragePoolType))
		}
		return nil
	}

	_, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		ComputeStoragePoolSpec_v1alpha1_FromProto,
		ComputeStoragePoolSpec_v1alpha1_ToProto,
		normalize,
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
	resource := ComputeStoragePoolSpec_v1alpha1_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	resource.Name = direct.LazyPtr(a.id.StoragePool)
	if resource.StoragePoolType != nil && !strings.Contains(*resource.StoragePoolType, "/") {
		resource.StoragePoolType = direct.LazyPtr(fmt.Sprintf("projects/%s/zones/%s/storagePoolTypes/%s", a.id.Project, a.id.Zone, *resource.StoragePoolType))
	}

	req := &computepb.UpdateStoragePoolRequest{
		Project:             a.id.Project,
		Zone:                a.id.Zone,
		StoragePool:         a.id.StoragePool,
		StoragePoolResource: resource,
	}
	if len(updateMask.Paths) > 0 {
		req.UpdateMask = direct.LazyPtr(strings.Join(updateMask.Paths, ","))
	}
	op, err := a.gcpClient.Update(ctx, req)
	if err != nil {
		return fmt.Errorf("updating ComputeStoragePool %s: %w", a.id, err)
	}
	if err := op.Wait(ctx); err != nil {
		return fmt.Errorf("waiting for ComputeStoragePool %s update: %w", a.id, err)
	}
	log.V(2).Info("successfully updated ComputeStoragePool", "name", a.id)

	updated, err := a.get(ctx)
	if err != nil {
		return fmt.Errorf("getting ComputeStoragePool %s after update: %w", a.id, err)
	}

	return a.setStatus(ctx, updateOp, updated)
}

func (a *StoragePoolAdapter) setStatus(ctx context.Context, op directbase.Operation, actual *computepb.StoragePool) error {
	status := &krm.ComputeStoragePoolStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = ComputeStoragePoolObservedState_v1alpha1_FromProto(mapCtx, actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return op.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *StoragePoolAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ComputeStoragePool{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(ComputeStoragePoolSpec_v1alpha1_FromProto(mapCtx, a.actual))
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
	u.SetName(a.id.StoragePool)
	u.SetGroupVersionKind(krm.ComputeStoragePoolGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *StoragePoolAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ComputeStoragePool", "name", a.id)

	req := &computepb.DeleteStoragePoolRequest{
		Project:     a.id.Project,
		Zone:        a.id.Zone,
		StoragePool: a.id.StoragePool,
	}
	op, err := a.gcpClient.Delete(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting ComputeStoragePool %s: %w", a.id, err)
	}
	if err := op.Wait(ctx); err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("waiting for ComputeStoragePool %s deletion: %w", a.id, err)
	}
	return true, nil
}
