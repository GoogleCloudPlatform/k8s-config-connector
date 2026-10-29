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

package compute

import (
	"context"
	"errors"
	"fmt"
	"strings"

	gcp "cloud.google.com/go/compute/apiv1"
	pb "cloud.google.com/go/compute/apiv1/computepb"
	"google.golang.org/api/iterator"
	"google.golang.org/protobuf/proto"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.ComputeRegionPerInstanceConfigGVK, NewComputeRegionPerInstanceConfigModel)
}

func NewComputeRegionPerInstanceConfigModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &computeRegionPerInstanceConfigModel{config: config}, nil
}

type computeRegionPerInstanceConfigModel struct {
	config *config.ControllerConfig
}

var _ directbase.Model = &computeRegionPerInstanceConfigModel{}

func (m *computeRegionPerInstanceConfigModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ComputeRegionPerInstanceConfig{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	picId, ok := id.(*krm.ComputeRegionPerInstanceConfigIdentity)
	if !ok {
		return nil, fmt.Errorf("unexpected identity type %T", id)
	}

	if picId.Project == "" || picId.Region == "" || picId.RegionInstanceGroupManager == "" || picId.Name == "" {
		return nil, fmt.Errorf("incomplete identity: %+v", picId)
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	gcpClient, err := newGCPClient(m.config)
	if err != nil {
		return nil, err
	}
	regionalClient, err := gcpClient.newRegionInstanceGroupManagersClient(ctx)
	if err != nil {
		return nil, err
	}

	return &ComputeRegionPerInstanceConfigAdapter{
		regionalClient: regionalClient,
		id:             picId,
		desired:        obj,
		reader:         reader,
	}, nil
}

func (m *computeRegionPerInstanceConfigModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type ComputeRegionPerInstanceConfigAdapter struct {
	regionalClient *gcp.RegionInstanceGroupManagersClient
	id             *krm.ComputeRegionPerInstanceConfigIdentity
	desired        *krm.ComputeRegionPerInstanceConfig
	actual         *pb.PerInstanceConfig
	reader         client.Reader
}

var _ directbase.Adapter = &ComputeRegionPerInstanceConfigAdapter{}

func (a *ComputeRegionPerInstanceConfigAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ComputeRegionPerInstanceConfig", "name", a.id.String())

	req := &pb.ListPerInstanceConfigsRegionInstanceGroupManagersRequest{
		Project:              a.id.Project,
		Region:               a.id.Region,
		InstanceGroupManager: a.id.RegionInstanceGroupManager,
	}
	it := a.regionalClient.ListPerInstanceConfigs(ctx, req)
	for {
		item, err := it.Next()
		if errors.Is(err, iterator.Done) {
			break
		}
		if err != nil {
			if direct.IsNotFound(err) {
				return false, nil
			}
			return false, fmt.Errorf("listing per instance configs for %s: %w", a.id.ParentString(), err)
		}
		if item.GetName() == a.id.Name {
			a.actual = item
			return true, nil
		}
	}

	return false, nil
}

func (a *ComputeRegionPerInstanceConfigAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating ComputeRegionPerInstanceConfig", "name", a.id.String())

	mapCtx := &direct.MapContext{}
	desiredPb := ComputeRegionPerInstanceConfigSpec_v1alpha1_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	desiredPb.Name = direct.LazyPtr(a.id.Name)

	req := &pb.CreateInstancesRegionInstanceGroupManagerRequest{
		Project:              a.id.Project,
		Region:               a.id.Region,
		InstanceGroupManager: a.id.RegionInstanceGroupManager,
		RegionInstanceGroupManagersCreateInstancesRequestResource: &pb.RegionInstanceGroupManagersCreateInstancesRequest{
			Instances: []*pb.PerInstanceConfig{desiredPb},
		},
	}
	op, err := a.regionalClient.CreateInstances(ctx, req)
	if err != nil {
		return fmt.Errorf("creating ComputeRegionPerInstanceConfig %s: %w", a.id.String(), err)
	}
	if err := op.Wait(ctx); err != nil {
		return fmt.Errorf("waiting for creation of ComputeRegionPerInstanceConfig %s: %w", a.id.String(), err)
	}

	return a.updateStatus(ctx, createOp)
}

func (a *ComputeRegionPerInstanceConfigAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating ComputeRegionPerInstanceConfig", "name", a.id.String())

	mapCtx := &direct.MapContext{}
	desiredPb := ComputeRegionPerInstanceConfigSpec_v1alpha1_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	desiredPb.Name = direct.LazyPtr(a.id.Name)

	diffs, err := compareComputeRegionPerInstanceConfig(ctx, a.actual, desiredPb)
	if err != nil {
		return err
	}

	if !diffs.HasDiff() {
		log.V(2).Info("no diff detected for ComputeRegionPerInstanceConfig", "name", a.id.String())
		return a.updateStatus(ctx, updateOp)
	}

	structuredreporting.ReportDiff(ctx, diffs)

	req := &pb.UpdatePerInstanceConfigsRegionInstanceGroupManagerRequest{
		Project:              a.id.Project,
		Region:               a.id.Region,
		InstanceGroupManager: a.id.RegionInstanceGroupManager,
		RegionInstanceGroupManagerUpdateInstanceConfigReqResource: &pb.RegionInstanceGroupManagerUpdateInstanceConfigReq{
			PerInstanceConfigs: []*pb.PerInstanceConfig{desiredPb},
		},
	}
	op, err := a.regionalClient.UpdatePerInstanceConfigs(ctx, req)
	if err != nil {
		return fmt.Errorf("updating ComputeRegionPerInstanceConfig %s: %w", a.id.String(), err)
	}
	if err := op.Wait(ctx); err != nil {
		return fmt.Errorf("waiting for update of ComputeRegionPerInstanceConfig %s: %w", a.id.String(), err)
	}

	// Instance name in applyUpdatesToInstances request must include zone
	instanceURL, err := a.findInstanceURL(ctx)
	if err != nil {
		return err
	}

	minAction := "NONE"
	if a.desired.Spec.MinimalAction != nil && *a.desired.Spec.MinimalAction != "" {
		minAction = *a.desired.Spec.MinimalAction
	}
	mostDisruptiveAction := "REPLACE"
	if a.desired.Spec.MostDisruptiveAllowedAction != nil && *a.desired.Spec.MostDisruptiveAllowedAction != "" {
		mostDisruptiveAction = *a.desired.Spec.MostDisruptiveAllowedAction
	}

	applyReq := &pb.ApplyUpdatesToInstancesRegionInstanceGroupManagerRequest{
		Project:              a.id.Project,
		Region:               a.id.Region,
		InstanceGroupManager: a.id.RegionInstanceGroupManager,
		RegionInstanceGroupManagersApplyUpdatesRequestResource: &pb.RegionInstanceGroupManagersApplyUpdatesRequest{
			Instances:                   []string{instanceURL},
			MinimalAction:               direct.LazyPtr(minAction),
			MostDisruptiveAllowedAction: direct.LazyPtr(mostDisruptiveAction),
		},
	}
	applyOp, err := a.regionalClient.ApplyUpdatesToInstances(ctx, applyReq)
	if err != nil {
		return fmt.Errorf("applying updates to ComputeRegionPerInstanceConfig %s: %w", a.id.String(), err)
	}
	if err := applyOp.Wait(ctx); err != nil {
		return fmt.Errorf("waiting for applyUpdatesToInstances of ComputeRegionPerInstanceConfig %s: %w", a.id.String(), err)
	}

	return a.updateStatus(ctx, updateOp)
}

func (a *ComputeRegionPerInstanceConfigAdapter) findInstanceURL(ctx context.Context) (string, error) {
	req := &pb.ListManagedInstancesRegionInstanceGroupManagersRequest{
		Project:              a.id.Project,
		Region:               a.id.Region,
		InstanceGroupManager: a.id.RegionInstanceGroupManager,
		MaxResults:           direct.LazyPtr(uint32(1)),
	}
	instanceNameToFind := "/" + a.id.Name
	it := a.regionalClient.ListManagedInstances(ctx, req)
	for {
		item, err := it.Next()
		if errors.Is(err, iterator.Done) {
			break
		}
		if err != nil {
			return "", fmt.Errorf("listing managed instances for %s: %w", a.id.ParentString(), err)
		}
		if item.GetName() == a.id.Name || strings.HasSuffix(item.GetInstance(), instanceNameToFind) {
			return item.GetInstance(), nil
		}
	}
	return "", fmt.Errorf("could not find instance %q in managed instances of %s", a.id.Name, a.id.ParentString())
}

func (a *ComputeRegionPerInstanceConfigAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ComputeRegionPerInstanceConfig", "name", a.id.String())

	req := &pb.DeletePerInstanceConfigsRegionInstanceGroupManagerRequest{
		Project:              a.id.Project,
		Region:               a.id.Region,
		InstanceGroupManager: a.id.RegionInstanceGroupManager,
		RegionInstanceGroupManagerDeleteInstanceConfigReqResource: &pb.RegionInstanceGroupManagerDeleteInstanceConfigReq{
			Names: []string{a.id.Name},
		},
	}
	op, err := a.regionalClient.DeletePerInstanceConfigs(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting ComputeRegionPerInstanceConfig %s: %w", a.id.String(), err)
	}
	if err := op.Wait(ctx); err != nil {
		return false, fmt.Errorf("waiting for deletion of ComputeRegionPerInstanceConfig %s: %w", a.id.String(), err)
	}

	if a.desired.Spec.RemoveInstanceStateOnDestroy != nil && *a.desired.Spec.RemoveInstanceStateOnDestroy {
		instanceURL, err := a.findInstanceURL(ctx)
		if err == nil {
			applyReq := &pb.ApplyUpdatesToInstancesRegionInstanceGroupManagerRequest{
				Project:              a.id.Project,
				Region:               a.id.Region,
				InstanceGroupManager: a.id.RegionInstanceGroupManager,
				RegionInstanceGroupManagersApplyUpdatesRequestResource: &pb.RegionInstanceGroupManagersApplyUpdatesRequest{
					Instances: []string{instanceURL},
				},
			}
			applyOp, err := a.regionalClient.ApplyUpdatesToInstances(ctx, applyReq)
			if err != nil {
				return false, fmt.Errorf("applying updates on destroy of ComputeRegionPerInstanceConfig %s: %w", a.id.String(), err)
			}
			if err := applyOp.Wait(ctx); err != nil {
				return false, fmt.Errorf("waiting for applyUpdatesToInstances on destroy of ComputeRegionPerInstanceConfig %s: %w", a.id.String(), err)
			}
		}
	}

	log.V(2).Info("successfully deleted ComputeRegionPerInstanceConfig", "name", a.id.String())
	return true, nil
}

func (a *ComputeRegionPerInstanceConfigAdapter) updateStatus(ctx context.Context, op directbase.Operation) error {
	status := &krm.ComputeRegionPerInstanceConfigStatus{}
	return op.UpdateStatus(ctx, status, nil)
}

func (a *ComputeRegionPerInstanceConfigAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ComputeRegionPerInstanceConfig{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(ComputeRegionPerInstanceConfigSpec_v1alpha1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.ProjectRef = refs.ProjectRef{External: a.id.Project}
	obj.Spec.Region = a.id.Region
	obj.Spec.RegionInstanceGroupManagerRef = krm.ComputeRegionInstanceGroupManagerRef{
		External: a.id.RegionInstanceGroupManager,
	}
	obj.Spec.ResourceID = direct.LazyPtr(a.id.Name)

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.actual.GetName())
	u.SetGroupVersionKind(krm.ComputeRegionPerInstanceConfigGVK)
	return u, nil
}

func compareComputeRegionPerInstanceConfig(ctx context.Context, actual, desired *pb.PerInstanceConfig) (*structuredreporting.Diff, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, ComputeRegionPerInstanceConfigSpec_v1alpha1_FromProto, ComputeRegionPerInstanceConfigSpec_v1alpha1_ToProto)
	if err != nil {
		return nil, err
	}
	maskedActual.Name = desired.Name

	clonedDesired := proto.Clone(desired).(*pb.PerInstanceConfig)

	diffs, _, err := common.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, err
	}
	return diffs, nil
}
