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
// proto.service: google.cloud.compute.v1.InstanceGroups
// proto.message: google.cloud.compute.v1.InstanceGroup
// crd.type: ComputeInstanceGroup
// crd.version: v1beta1

package compute

import (
	"context"
	"errors"
	"fmt"
	"strings"

	compute "cloud.google.com/go/compute/apiv1"
	pb "cloud.google.com/go/compute/apiv1/computepb"
	"google.golang.org/api/iterator"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	apirefs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/tags"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.ComputeInstanceGroupGVK, NewComputeInstanceGroupModel)
}

func NewComputeInstanceGroupModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &computeInstanceGroupModel{config: config}, nil
}

var _ directbase.Model = &computeInstanceGroupModel{}

type computeInstanceGroupModel struct {
	config *config.ControllerConfig
}

func (m *computeInstanceGroupModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ComputeInstanceGroup{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}

	gcpClient, err := newGCPClient(m.config)
	if err != nil {
		return nil, fmt.Errorf("building gcp client: %w", err)
	}

	instanceGroupsClient, err := gcpClient.newInstanceGroupsClient(ctx)
	if err != nil {
		return nil, err
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	mapCtx := &direct.MapContext{}
	desiredObj := obj.DeepCopy()
	resource := ComputeInstanceGroupSpec_v1beta1_ToProto(mapCtx, &desiredObj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	var desiredInstances []string
	for _, inst := range desiredObj.Spec.Instances {
		if inst.External != "" {
			desiredInstances = append(desiredInstances, inst.External)
		}
	}

	return &ComputeInstanceGroupAdapter{
		gcpClient:        instanceGroupsClient,
		id:               id.(*krm.ComputeInstanceGroupIdentity),
		desired:          resource,
		desiredInstances: desiredInstances,
		reader:           reader,
	}, nil
}

func (m *computeInstanceGroupModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type ComputeInstanceGroupAdapter struct {
	gcpClient        *compute.InstanceGroupsClient
	id               *krm.ComputeInstanceGroupIdentity
	desired          *pb.InstanceGroup
	desiredInstances []string
	actual           *pb.InstanceGroup
	actualInstances  []string
	reader           client.Reader
}

var _ directbase.Adapter = &ComputeInstanceGroupAdapter{}

func (a *ComputeInstanceGroupAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ComputeInstanceGroup", "name", a.id)

	req := &pb.GetInstanceGroupRequest{
		Project:       a.id.Project,
		Zone:          a.id.Zone,
		InstanceGroup: a.id.InstanceGroup,
	}
	actual, err := a.gcpClient.Get(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ComputeInstanceGroup %s: %w", a.id, err)
	}

	a.actual = actual

	// Also retrieve current instances
	listReq := &pb.ListInstancesInstanceGroupsRequest{
		Project:       a.id.Project,
		Zone:          a.id.Zone,
		InstanceGroup: a.id.InstanceGroup,
	}
	it := a.gcpClient.ListInstances(ctx, listReq)
	a.actualInstances = nil
	for {
		item, err := it.Next()
		if err != nil {
			if errors.Is(err, iterator.Done) {
				break
			}
			return false, fmt.Errorf("listing instances of ComputeInstanceGroup %s: %w", a.id, err)
		}
		if item.GetInstance() != "" {
			a.actualInstances = append(a.actualInstances, item.GetInstance())
		}
	}

	return true, nil
}

func (a *ComputeInstanceGroupAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating ComputeInstanceGroup", "name", a.id)

	a.desired.Name = direct.LazyPtr(a.id.InstanceGroup)

	req := &pb.InsertInstanceGroupRequest{
		Project:               a.id.Project,
		Zone:                  a.id.Zone,
		InstanceGroupResource: a.desired,
	}
	op, errInsert := a.gcpClient.Insert(ctx, req)
	if errInsert != nil {
		return fmt.Errorf("creating ComputeInstanceGroup %s: %w", a.id, errInsert)
	}
	err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting ComputeInstanceGroup %s create failed: %w", a.id, err)
	}

	// Now add instances if there are any
	if len(a.desiredInstances) > 0 {
		var instanceRefs []*pb.InstanceReference
		for _, inst := range a.desiredInstances {
			instanceRefs = append(instanceRefs, &pb.InstanceReference{
				Instance: direct.LazyPtr(inst),
			})
		}
		addReq := &pb.AddInstancesInstanceGroupRequest{
			Project:       a.id.Project,
			Zone:          a.id.Zone,
			InstanceGroup: a.id.InstanceGroup,
			InstanceGroupsAddInstancesRequestResource: &pb.InstanceGroupsAddInstancesRequest{
				Instances: instanceRefs,
			},
		}
		addOp, errAdd := a.gcpClient.AddInstances(ctx, addReq)
		if errAdd != nil {
			return fmt.Errorf("adding instances to ComputeInstanceGroup %s: %w", a.id, errAdd)
		}
		if err := addOp.Wait(ctx); err != nil {
			return fmt.Errorf("waiting to add instances to ComputeInstanceGroup %s: %w", a.id, err)
		}
	}

	log.V(2).Info("successfully created ComputeInstanceGroup", "name", a.id)

	// Retrieve actual state to update status
	exists, err := a.Find(ctx)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("created ComputeInstanceGroup was not found after creation")
	}

	return a.updateStatus(ctx, createOp, a.actual)
}

func (a *ComputeInstanceGroupAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating ComputeInstanceGroup", "name", a.id)

	diffs, updateMask, err := compareComputeInstanceGroup(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	if diffs.HasDiff() {
		// Check for immutable field updates
		for _, path := range updateMask.Paths {
			if path != "named_ports" {
				return fmt.Errorf("ComputeInstanceGroup field %q is immutable and cannot be updated", path)
			}
		}

		// Update named ports
		setReq := &pb.SetNamedPortsInstanceGroupRequest{
			Project:       a.id.Project,
			Zone:          a.id.Zone,
			InstanceGroup: a.id.InstanceGroup,
			InstanceGroupsSetNamedPortsRequestResource: &pb.InstanceGroupsSetNamedPortsRequest{
				NamedPorts: a.desired.NamedPorts,
			},
		}
		setOp, errSet := a.gcpClient.SetNamedPorts(ctx, setReq)
		if errSet != nil {
			return fmt.Errorf("updating named ports for ComputeInstanceGroup %s: %w", a.id, errSet)
		}
		if err := setOp.Wait(ctx); err != nil {
			return fmt.Errorf("waiting to update named ports for ComputeInstanceGroup %s: %w", a.id, err)
		}
		structuredreporting.ReportDiff(ctx, diffs)
	}

	// Now handle instances add/remove
	toAdd, toRemove := diffInstances(a.actualInstances, a.desiredInstances)

	if len(toRemove) > 0 {
		var instanceRefs []*pb.InstanceReference
		for _, inst := range toRemove {
			instanceRefs = append(instanceRefs, &pb.InstanceReference{
				Instance: direct.LazyPtr(inst),
			})
		}
		removeReq := &pb.RemoveInstancesInstanceGroupRequest{
			Project:       a.id.Project,
			Zone:          a.id.Zone,
			InstanceGroup: a.id.InstanceGroup,
			InstanceGroupsRemoveInstancesRequestResource: &pb.InstanceGroupsRemoveInstancesRequest{
				Instances: instanceRefs,
			},
		}
		removeOp, errRemove := a.gcpClient.RemoveInstances(ctx, removeReq)
		if errRemove != nil {
			return fmt.Errorf("removing instances from ComputeInstanceGroup %s: %w", a.id, errRemove)
		}
		if err := removeOp.Wait(ctx); err != nil {
			return fmt.Errorf("waiting to remove instances from ComputeInstanceGroup %s: %w", a.id, err)
		}
	}

	if len(toAdd) > 0 {
		var instanceRefs []*pb.InstanceReference
		for _, inst := range toAdd {
			instanceRefs = append(instanceRefs, &pb.InstanceReference{
				Instance: direct.LazyPtr(inst),
			})
		}
		addReq := &pb.AddInstancesInstanceGroupRequest{
			Project:       a.id.Project,
			Zone:          a.id.Zone,
			InstanceGroup: a.id.InstanceGroup,
			InstanceGroupsAddInstancesRequestResource: &pb.InstanceGroupsAddInstancesRequest{
				Instances: instanceRefs,
			},
		}
		addOp, errAdd := a.gcpClient.AddInstances(ctx, addReq)
		if errAdd != nil {
			return fmt.Errorf("adding instances to ComputeInstanceGroup %s: %w", a.id, errAdd)
		}
		if err := addOp.Wait(ctx); err != nil {
			return fmt.Errorf("waiting to add instances to ComputeInstanceGroup %s: %w", a.id, err)
		}
	}

	log.V(2).Info("successfully updated ComputeInstanceGroup", "name", a.id)

	// Fetch latest state to update status
	exists, err := a.Find(ctx)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("updated ComputeInstanceGroup was not found after update")
	}

	return a.updateStatus(ctx, updateOp, a.actual)
}

func (a *ComputeInstanceGroupAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ComputeInstanceGroup", "name", a.id)

	req := &pb.DeleteInstanceGroupRequest{
		Project:       a.id.Project,
		Zone:          a.id.Zone,
		InstanceGroup: a.id.InstanceGroup,
	}
	op, err := a.gcpClient.Delete(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting ComputeInstanceGroup %s: %w", a.id, err)
	}

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting ComputeInstanceGroup %s delete failed: %w", a.id, err)
	}

	log.V(2).Info("successfully deleted ComputeInstanceGroup", "name", a.id)
	return true, nil
}

func (a *ComputeInstanceGroupAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.InstanceGroup) error {
	mapCtx := &direct.MapContext{}
	status := ComputeInstanceGroupStatus_v1beta1_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return op.UpdateStatus(ctx, status, nil)
}

func lastComponent(path string) string {
	idx := strings.LastIndex(path, "/")
	if idx >= 0 {
		return path[idx+1:]
	}
	return path
}

func compareComputeInstanceGroup(ctx context.Context, actual, desired *pb.InstanceGroup) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, ComputeInstanceGroupSpec_v1beta1_FromProto, ComputeInstanceGroupSpec_v1beta1_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = desired.Name

	clonedDesired := proto.Clone(desired).(*pb.InstanceGroup)

	populateDefaults := func(obj *pb.InstanceGroup) {
		if obj.Network != nil {
			*obj.Network = lastComponent(*obj.Network)
		}
		if obj.Zone != nil {
			*obj.Zone = lastComponent(*obj.Zone)
		}
	}
	populateDefaults(maskedActual)
	populateDefaults(clonedDesired)

	diffs, updateMask, err := tags.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}

func diffInstances(actual, desired []string) (toAdd, toRemove []string) {
	actualMap := make(map[string]string)
	for _, inst := range actual {
		normalized := apirefs.TrimComputeURIPrefix(inst)
		actualMap[normalized] = inst
	}

	desiredMap := make(map[string]string)
	for _, inst := range desired {
		normalized := apirefs.TrimComputeURIPrefix(inst)
		desiredMap[normalized] = inst
	}

	for normalized, inst := range desiredMap {
		if _, exists := actualMap[normalized]; !exists {
			toAdd = append(toAdd, inst)
		}
	}

	for normalized, inst := range actualMap {
		if _, exists := desiredMap[normalized]; !exists {
			toRemove = append(toRemove, inst)
		}
	}

	return toAdd, toRemove
}

func (a *ComputeInstanceGroupAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ComputeInstanceGroup{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(ComputeInstanceGroupSpec_v1beta1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	for _, inst := range a.actualInstances {
		obj.Spec.Instances = append(obj.Spec.Instances, krm.InstanceRef{
			External: inst,
		})
	}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.actual.GetName())
	u.SetGroupVersionKind(krm.ComputeInstanceGroupGVK)
	return u, nil
}
