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

// +tool:controller
// proto.service: google.cloud.compute.v1.InstanceGroupManagers
// proto.service: google.cloud.compute.v1.RegionInstanceGroupManagers
// proto.message: google.cloud.compute.v1.InstanceGroupManager
// crd.type: ComputeInstanceGroupManager
// crd.version: v1beta1

package compute

import (
	"context"
	"fmt"

	"google.golang.org/protobuf/proto"

	gcp "cloud.google.com/go/compute/apiv1"
	pb "cloud.google.com/go/compute/apiv1/computepb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.ComputeInstanceGroupManagerGVK, NewComputeInstanceGroupManagerModel)
}

func NewComputeInstanceGroupManagerModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &computeInstanceGroupManagerModel{config: config}, nil
}

var _ directbase.Model = &computeInstanceGroupManagerModel{}

type computeInstanceGroupManagerModel struct {
	config *config.ControllerConfig
}

func (m *computeInstanceGroupManagerModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ComputeInstanceGroupManager{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	managerId := id.(*krm.ComputeInstanceGroupManagerIdentity)
	if err := managerId.Validate(); err != nil {
		return nil, err
	}

	gcpClient, err := newGCPClient(m.config)
	if err != nil {
		return nil, fmt.Errorf("building gcp client: %w", err)
	}

	zonalClient, err := gcpClient.newInstanceGroupManagersClient(ctx)
	if err != nil {
		return nil, err
	}
	regionalClient, err := gcpClient.newRegionInstanceGroupManagersClient(ctx)
	if err != nil {
		return nil, err
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	mapCtx := &direct.MapContext{}
	desired := ComputeInstanceGroupManagerSpec_v1beta1_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	desired.Name = direct.LazyPtr(managerId.InstanceGroupManager)

	return &ComputeInstanceGroupManagerAdapter{
		zonalClient:    zonalClient,
		regionalClient: regionalClient,
		id:             managerId,
		desired:        desired,
		reader:         reader,
	}, nil
}

func (m *computeInstanceGroupManagerModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type ComputeInstanceGroupManagerAdapter struct {
	zonalClient    *gcp.InstanceGroupManagersClient
	regionalClient *gcp.RegionInstanceGroupManagersClient
	id             *krm.ComputeInstanceGroupManagerIdentity
	desired        *pb.InstanceGroupManager
	actual         *pb.InstanceGroupManager
	reader         client.Reader
}

var _ directbase.Adapter = &ComputeInstanceGroupManagerAdapter{}

func (a *ComputeInstanceGroupManagerAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ComputeInstanceGroupManager", "name", a.id)

	if a.id.IsZonal() {
		req := &pb.GetInstanceGroupManagerRequest{
			Project:              a.id.Project,
			Zone:                 a.id.Zone,
			InstanceGroupManager: a.id.InstanceGroupManager,
		}
		actual, err := a.zonalClient.Get(ctx, req)
		if err != nil {
			if direct.IsNotFound(err) {
				return false, nil
			}
			return false, fmt.Errorf("getting zonal ComputeInstanceGroupManager %q: %w", a.id, err)
		}
		a.actual = actual
		return true, nil
	} else if a.id.IsRegional() {
		req := &pb.GetRegionInstanceGroupManagerRequest{
			Project:              a.id.Project,
			Region:               a.id.Region,
			InstanceGroupManager: a.id.InstanceGroupManager,
		}
		actual, err := a.regionalClient.Get(ctx, req)
		if err != nil {
			if direct.IsNotFound(err) {
				return false, nil
			}
			return false, fmt.Errorf("getting regional ComputeInstanceGroupManager %q: %w", a.id, err)
		}
		a.actual = actual
		return true, nil
	} else {
		return false, fmt.Errorf("ComputeInstanceGroupManager %s is neither zonal nor regional", a.id)
	}
}

func (a *ComputeInstanceGroupManagerAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating ComputeInstanceGroupManager", "name", a.id)

	if a.id.IsZonal() {
		req := &pb.InsertInstanceGroupManagerRequest{
			Project:                      a.id.Project,
			Zone:                         a.id.Zone,
			InstanceGroupManagerResource: a.desired,
		}
		op, err := a.zonalClient.Insert(ctx, req)
		if err != nil {
			return fmt.Errorf("creating zonal ComputeInstanceGroupManager %s: %w", a.id, err)
		}
		err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("waiting zonal ComputeInstanceGroupManager %s creation failed: %w", a.id, err)
		}
	} else if a.id.IsRegional() {
		req := &pb.InsertRegionInstanceGroupManagerRequest{
			Project:                      a.id.Project,
			Region:                       a.id.Region,
			InstanceGroupManagerResource: a.desired,
		}
		op, err := a.regionalClient.Insert(ctx, req)
		if err != nil {
			return fmt.Errorf("creating regional ComputeInstanceGroupManager %s: %w", a.id, err)
		}
		err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("waiting regional ComputeInstanceGroupManager %s creation failed: %w", a.id, err)
		}
	} else {
		return fmt.Errorf("ComputeInstanceGroupManager %s is neither zonal nor regional", a.id)
	}

	log.V(2).Info("successfully created ComputeInstanceGroupManager", "name", a.id)

	// Get the created resource to update status
	exists, err := a.Find(ctx)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("created ComputeInstanceGroupManager was not found after creation")
	}

	return a.updateStatus(ctx, createOp, a.actual)
}

func (a *ComputeInstanceGroupManagerAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating ComputeInstanceGroupManager", "name", a.id)

	diffs, err := compareComputeInstanceGroupManager(ctx, a.actual, a.desired, a.id)
	if err != nil {
		return err
	}

	if diffs.HasDiff() {
		diffs.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, diffs)

		updateOp.RecordUpdatingEvent()

		if a.id.IsZonal() {
			req := &pb.PatchInstanceGroupManagerRequest{
				Project:                      a.id.Project,
				Zone:                         a.id.Zone,
				InstanceGroupManager:         a.id.InstanceGroupManager,
				InstanceGroupManagerResource: a.desired,
			}
			op, err := a.zonalClient.Patch(ctx, req)
			if err != nil {
				return fmt.Errorf("patching zonal ComputeInstanceGroupManager %s: %w", a.id, err)
			}
			err = op.Wait(ctx)
			if err != nil {
				return fmt.Errorf("waiting zonal ComputeInstanceGroupManager %s patch failed: %w", a.id, err)
			}
		} else if a.id.IsRegional() {
			req := &pb.PatchRegionInstanceGroupManagerRequest{
				Project:                      a.id.Project,
				Region:                       a.id.Region,
				InstanceGroupManager:         a.id.InstanceGroupManager,
				InstanceGroupManagerResource: a.desired,
			}
			op, err := a.regionalClient.Patch(ctx, req)
			if err != nil {
				return fmt.Errorf("patching regional ComputeInstanceGroupManager %s: %w", a.id, err)
			}
			err = op.Wait(ctx)
			if err != nil {
				return fmt.Errorf("waiting regional ComputeInstanceGroupManager %s patch failed: %w", a.id, err)
			}
		} else {
			return fmt.Errorf("ComputeInstanceGroupManager %s is neither zonal nor regional", a.id)
		}

		log.V(2).Info("successfully updated ComputeInstanceGroupManager", "name", a.id)
	}

	// Fetch latest state to update status
	exists, err := a.Find(ctx)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("updated ComputeInstanceGroupManager was not found after update")
	}

	return a.updateStatus(ctx, updateOp, a.actual)
}

func (a *ComputeInstanceGroupManagerAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ComputeInstanceGroupManager", "name", a.id)

	if a.id.IsZonal() {
		req := &pb.DeleteInstanceGroupManagerRequest{
			Project:              a.id.Project,
			Zone:                 a.id.Zone,
			InstanceGroupManager: a.id.InstanceGroupManager,
		}
		op, err := a.zonalClient.Delete(ctx, req)
		if err != nil {
			if direct.IsNotFound(err) {
				return true, nil
			}
			return false, fmt.Errorf("deleting zonal ComputeInstanceGroupManager %s: %w", a.id, err)
		}
		err = op.Wait(ctx)
		if err != nil {
			return false, fmt.Errorf("waiting zonal ComputeInstanceGroupManager %s delete failed: %w", a.id, err)
		}
	} else if a.id.IsRegional() {
		req := &pb.DeleteRegionInstanceGroupManagerRequest{
			Project:              a.id.Project,
			Region:               a.id.Region,
			InstanceGroupManager: a.id.InstanceGroupManager,
		}
		op, err := a.regionalClient.Delete(ctx, req)
		if err != nil {
			if direct.IsNotFound(err) {
				return true, nil
			}
			return false, fmt.Errorf("deleting regional ComputeInstanceGroupManager %s: %w", a.id, err)
		}
		err = op.Wait(ctx)
		if err != nil {
			return false, fmt.Errorf("waiting regional ComputeInstanceGroupManager %s delete failed: %w", a.id, err)
		}
	} else {
		return false, fmt.Errorf("ComputeInstanceGroupManager %s is neither zonal nor regional", a.id)
	}

	log.V(2).Info("successfully deleted ComputeInstanceGroupManager", "name", a.id)
	return true, nil
}

func (a *ComputeInstanceGroupManagerAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.InstanceGroupManager) error {
	mapCtx := &direct.MapContext{}
	status := ComputeInstanceGroupManagerStatus_v1beta1_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return op.UpdateStatus(ctx, &status, nil)
}

func (a *ComputeInstanceGroupManagerAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ComputeInstanceGroupManager{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(ComputeInstanceGroupManagerSpec_v1beta1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.ProjectRef.External = a.id.Project
	if a.id.IsZonal() {
		obj.Spec.Location = &a.id.Zone
	} else if a.id.IsRegional() {
		obj.Spec.Location = &a.id.Region
	} else {
		return nil, fmt.Errorf("ComputeInstanceGroupManager %s is neither zonal nor regional", a.id)
	}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.actual.GetName())
	u.SetGroupVersionKind(krm.ComputeInstanceGroupManagerGVK)
	return u, nil
}

func compareComputeInstanceGroupManager(ctx context.Context, actual, desired *pb.InstanceGroupManager, id *krm.ComputeInstanceGroupManagerIdentity) (*structuredreporting.Diff, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, ComputeInstanceGroupManagerSpec_v1beta1_FromProto, ComputeInstanceGroupManagerSpec_v1beta1_ToProto)
	if err != nil {
		return nil, err
	}
	maskedActual.Name = actual.Name

	clonedDesired := proto.CloneOf(desired)
	clonedDesired.Name = actual.Name

	populateDefaults := func(obj *pb.InstanceGroupManager) {
		if obj.InstanceTemplate != nil {
			*obj.InstanceTemplate = refs.TrimComputeURIPrefix(*obj.InstanceTemplate)
		}
		for i := range obj.AutoHealingPolicies {
			if obj.AutoHealingPolicies[i].HealthCheck != nil {
				*obj.AutoHealingPolicies[i].HealthCheck = refs.TrimComputeURIPrefix(*obj.AutoHealingPolicies[i].HealthCheck)
			}
		}
		for i := range obj.Versions {
			if obj.Versions[i].InstanceTemplate != nil {
				*obj.Versions[i].InstanceTemplate = refs.TrimComputeURIPrefix(*obj.Versions[i].InstanceTemplate)
			}
		}
		for i := range obj.TargetPools {
			obj.TargetPools[i] = refs.TrimComputeURIPrefix(obj.TargetPools[i])
		}
		if obj.DistributionPolicy != nil {
			for _, z := range obj.DistributionPolicy.Zones {
				if z.Zone != nil {
					*z.Zone = lastComponent(*z.Zone)
				}
			}
		}
	}

	populateDefaults(clonedDesired)
	populateDefaults(maskedActual)

	_, diffs, err := common.CompareProtoMessageStructuredDiff(clonedDesired, maskedActual, common.BasicDiff)
	if err != nil {
		return nil, err
	}
	return diffs, nil
}
