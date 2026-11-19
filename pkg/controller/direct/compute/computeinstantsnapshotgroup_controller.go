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
// krm.kind: ComputeInstantSnapshotGroup

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
	registry.RegisterModel(krm.ComputeInstantSnapshotGroupGVK, NewInstantSnapshotGroupModel)
}

func NewInstantSnapshotGroupModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelInstantSnapshotGroup{config: *config}, nil
}

var _ directbase.Model = &modelInstantSnapshotGroup{}

type modelInstantSnapshotGroup struct {
	config config.ControllerConfig
}

func (m *modelInstantSnapshotGroup) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ComputeInstantSnapshotGroup{}
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
	id := i.(*krm.ComputeInstantSnapshotGroupIdentity)
	if err := id.Validate(); err != nil {
		return nil, err
	}

	gcpClient, err := newGCPClient(&m.config)
	if err != nil {
		return nil, fmt.Errorf("building gcp client: %w", err)
	}
	instantSnapshotGroupsClient, err := gcpClient.newInstantSnapshotGroupsClient(ctx)
	if err != nil {
		return nil, err
	}
	regionInstantSnapshotGroupsClient, err := gcpClient.newRegionInstantSnapshotGroupsClient(ctx)
	if err != nil {
		return nil, err
	}

	return &InstantSnapshotGroupAdapter{
		id:                                id,
		instantSnapshotGroupsClient:       instantSnapshotGroupsClient,
		regionInstantSnapshotGroupsClient: regionInstantSnapshotGroupsClient,
		desired:                           obj,
	}, nil
}

func (m *modelInstantSnapshotGroup) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//compute.googleapis.com/") {
		return nil, nil
	}
	id := &krm.ComputeInstantSnapshotGroupIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	if err := id.Validate(); err != nil {
		return nil, nil
	}

	gcpClient, err := newGCPClient(&m.config)
	if err != nil {
		return nil, fmt.Errorf("building gcp client: %w", err)
	}
	instantSnapshotGroupsClient, err := gcpClient.newInstantSnapshotGroupsClient(ctx)
	if err != nil {
		return nil, err
	}
	regionInstantSnapshotGroupsClient, err := gcpClient.newRegionInstantSnapshotGroupsClient(ctx)
	if err != nil {
		return nil, err
	}

	return &InstantSnapshotGroupAdapter{
		id:                                id,
		instantSnapshotGroupsClient:       instantSnapshotGroupsClient,
		regionInstantSnapshotGroupsClient: regionInstantSnapshotGroupsClient,
	}, nil
}

type InstantSnapshotGroupAdapter struct {
	id                                *krm.ComputeInstantSnapshotGroupIdentity
	instantSnapshotGroupsClient       *gcp.InstantSnapshotGroupsClient
	regionInstantSnapshotGroupsClient *gcp.RegionInstantSnapshotGroupsClient
	desired                           *krm.ComputeInstantSnapshotGroup
	actual                            *computepb.InstantSnapshotGroup
}

var _ directbase.Adapter = &InstantSnapshotGroupAdapter{}

func (a *InstantSnapshotGroupAdapter) get(ctx context.Context) (*computepb.InstantSnapshotGroup, error) {
	if a.id.IsZonal() {
		req := &computepb.GetInstantSnapshotGroupRequest{
			Project:              a.id.Project,
			Zone:                 a.id.Zone,
			InstantSnapshotGroup: a.id.InstantSnapshotGroup,
		}
		return a.instantSnapshotGroupsClient.Get(ctx, req)
	} else if a.id.IsRegional() {
		req := &computepb.GetRegionInstantSnapshotGroupRequest{
			Project:              a.id.Project,
			Region:               a.id.Region,
			InstantSnapshotGroup: a.id.InstantSnapshotGroup,
		}
		return a.regionInstantSnapshotGroupsClient.Get(ctx, req)
	}
	return nil, fmt.Errorf("ComputeInstantSnapshotGroup %s is neither zonal nor regional", a.id)
}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *InstantSnapshotGroupAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting InstantSnapshotGroup", "name", a.id)

	actual, err := a.get(ctx)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ComputeInstantSnapshotGroup %q: %w", a.id, err)
	}
	a.actual = actual
	return true, nil
}

// Create creates the resource in GCP based on `spec` and updates the Config Connector object `status` based on the GCP response.
func (a *InstantSnapshotGroupAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating InstantSnapshotGroup", "name", a.id)

	mapCtx := &direct.MapContext{}
	resource := ComputeInstantSnapshotGroupSpec_v1alpha1_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	resource.Name = direct.LazyPtr(a.id.InstantSnapshotGroup)
	if resource.SourceConsistencyGroup != nil {
		resource.SourceConsistencyGroup = canonicalizeComputeURL(resource.SourceConsistencyGroup)
	}

	if a.id.IsZonal() {
		req := &computepb.InsertInstantSnapshotGroupRequest{
			Project:                      a.id.Project,
			Zone:                         a.id.Zone,
			InstantSnapshotGroupResource: resource,
		}
		op, err := a.instantSnapshotGroupsClient.Insert(ctx, req)
		if err != nil {
			return fmt.Errorf("creating ComputeInstantSnapshotGroup %s: %w", a.id, err)
		}
		if err := op.Wait(ctx); err != nil {
			return fmt.Errorf("waiting for ComputeInstantSnapshotGroup %s creation: %w", a.id, err)
		}
	} else if a.id.IsRegional() {
		req := &computepb.InsertRegionInstantSnapshotGroupRequest{
			Project:                      a.id.Project,
			Region:                       a.id.Region,
			InstantSnapshotGroupResource: resource,
		}
		op, err := a.regionInstantSnapshotGroupsClient.Insert(ctx, req)
		if err != nil {
			return fmt.Errorf("creating ComputeInstantSnapshotGroup %s: %w", a.id, err)
		}
		if err := op.Wait(ctx); err != nil {
			return fmt.Errorf("waiting for ComputeInstantSnapshotGroup %s creation: %w", a.id, err)
		}
	} else {
		return fmt.Errorf("ComputeInstantSnapshotGroup %s is neither zonal nor regional", a.id)
	}

	created, err := a.get(ctx)
	if err != nil {
		return fmt.Errorf("getting ComputeInstantSnapshotGroup %s after creation: %w", a.id, err)
	}

	return a.setStatus(ctx, createOp, created)
}

// Update updates the resource in GCP based on `spec` and updates the Config Connector object `status` based on the GCP response.
func (a *InstantSnapshotGroupAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating InstantSnapshotGroup", "name", a.id)

	normalize := func(ctx context.Context, pb *computepb.InstantSnapshotGroup) error {
		if pb.SourceConsistencyGroup != nil {
			pb.SourceConsistencyGroup = canonicalizeComputeURL(pb.SourceConsistencyGroup)
		}
		return nil
	}

	_, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		ComputeInstantSnapshotGroupSpec_v1alpha1_FromProto,
		ComputeInstantSnapshotGroupSpec_v1alpha1_ToProto,
		normalize,
	)
	if err != nil {
		return err
	}
	if report.HasDiff() {
		report.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, report)
		return fmt.Errorf("updating ComputeInstantSnapshotGroup is not supported, fields modified: %v", updateMask.Paths)
	}

	return a.setStatus(ctx, updateOp, a.actual)
}

func (a *InstantSnapshotGroupAdapter) setStatus(ctx context.Context, op directbase.Operation, actual *computepb.InstantSnapshotGroup) error {
	status := &krm.ComputeInstantSnapshotGroupStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = ComputeInstantSnapshotGroupObservedState_v1alpha1_FromProto(mapCtx, actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return op.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *InstantSnapshotGroupAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ComputeInstantSnapshotGroup{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(ComputeInstantSnapshotGroupSpec_v1alpha1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.ProjectRef = &refsv1beta1.ProjectRef{External: a.id.Project}
	if a.id.IsZonal() {
		obj.Spec.Location = direct.LazyPtr(a.id.Zone)
	} else if a.id.IsRegional() {
		obj.Spec.Location = direct.LazyPtr(a.id.Region)
	}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.InstantSnapshotGroup)
	u.SetGroupVersionKind(krm.ComputeInstantSnapshotGroupGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *InstantSnapshotGroupAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting InstantSnapshotGroup", "name", a.id)

	if a.id.IsZonal() {
		req := &computepb.DeleteInstantSnapshotGroupRequest{
			Project:              a.id.Project,
			Zone:                 a.id.Zone,
			InstantSnapshotGroup: a.id.InstantSnapshotGroup,
		}
		op, err := a.instantSnapshotGroupsClient.Delete(ctx, req)
		if err != nil {
			if direct.IsNotFound(err) {
				return true, nil
			}
			return false, fmt.Errorf("deleting ComputeInstantSnapshotGroup %s: %w", a.id, err)
		}
		if err := op.Wait(ctx); err != nil {
			return false, fmt.Errorf("waiting for ComputeInstantSnapshotGroup %s deletion: %w", a.id, err)
		}
	} else if a.id.IsRegional() {
		req := &computepb.DeleteRegionInstantSnapshotGroupRequest{
			Project:              a.id.Project,
			Region:               a.id.Region,
			InstantSnapshotGroup: a.id.InstantSnapshotGroup,
		}
		op, err := a.regionInstantSnapshotGroupsClient.Delete(ctx, req)
		if err != nil {
			if direct.IsNotFound(err) {
				return true, nil
			}
			return false, fmt.Errorf("deleting ComputeInstantSnapshotGroup %s: %w", a.id, err)
		}
		if err := op.Wait(ctx); err != nil {
			return false, fmt.Errorf("waiting for ComputeInstantSnapshotGroup %s deletion: %w", a.id, err)
		}
	} else {
		return false, fmt.Errorf("ComputeInstantSnapshotGroup %s is neither zonal nor regional", a.id)
	}

	log.V(2).Info("successfully deleted ComputeInstantSnapshotGroup", "name", a.id)
	return true, nil
}
