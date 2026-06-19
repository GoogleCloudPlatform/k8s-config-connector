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
// proto.service: google.cloud.compute.v1.NodeGroups
// proto.message: google.cloud.compute.v1.NodeGroup
// crd.type: ComputeNodeGroup
// crd.version: v1beta1

package compute

import (
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"

	compute "cloud.google.com/go/compute/apiv1"
	computepb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.ComputeNodeGroupGVK, NewNodeGroupModel)
}

func NewNodeGroupModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &nodeGroupModel{config: config}, nil
}

var _ directbase.Model = &nodeGroupModel{}

type nodeGroupModel struct {
	config *config.ControllerConfig
}

func (m *nodeGroupModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ComputeNodeGroup{}
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
	nodeGroupsClient, err := gcpClient.newNodeGroupsClient(ctx)
	if err != nil {
		return nil, err
	}

	return &NodeGroupAdapter{
		gcpClient:     nodeGroupsClient,
		id:            id.(*krm.ComputeNodeGroupIdentity),
		desired:       obj,
		reader:        reader,
		projectMapper: m.config.ProjectMapper,
	}, nil
}

func (m *nodeGroupModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type NodeGroupAdapter struct {
	gcpClient     *compute.NodeGroupsClient
	id            *krm.ComputeNodeGroupIdentity
	desired       *krm.ComputeNodeGroup
	actual        *computepb.NodeGroup
	reader        client.Reader
	projectMapper *projects.ProjectMapper
}

var _ directbase.Adapter = &NodeGroupAdapter{}

// Find retrieves the GCP resource.
func (a *NodeGroupAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting NodeGroup", "name", a.id)

	req := &computepb.GetNodeGroupRequest{
		Project:   a.id.Project,
		Zone:      a.id.Zone,
		NodeGroup: a.id.NodeGroup,
	}
	actual, err := a.gcpClient.Get(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting NodeGroup %q: %w", a.id, err)
	}

	a.actual = actual
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *NodeGroupAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating NodeGroup", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	if err := ResolveComputeNodeGroupRefs(ctx, a.reader, a.projectMapper, desired); err != nil {
		return err
	}
	resource := ComputeNodeGroupSpec_v1beta1_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	resource.Name = direct.LazyPtr(a.id.NodeGroup)
	// match realGCP log: "zone": "projects/${projectId}/global/zones/us-central1-a"
	resource.Zone = direct.LazyPtr(fmt.Sprintf("projects/%s/global/zones/%s", a.id.Project, a.id.Zone))

	if resource.Size == nil && desired.Spec.InitialSize != nil {
		resource.Size = direct.PtrInt64ToPtrInt32(desired.Spec.InitialSize)
	}

	// Wait, we need to pass initialNodeCount (which is resource.Size) and the nodeGroupResource
	// In compute.v1, InsertNodeGroupRequest has:
	// - Project
	// - Zone
	// - InitialNodeCount (int32)
	// - NodeGroupResource (*computepb.NodeGroup)
	req := &computepb.InsertNodeGroupRequest{
		Project:           a.id.Project,
		Zone:              a.id.Zone,
		InitialNodeCount:  resource.GetSize(),
		NodeGroupResource: resource,
	}
	op, err := a.gcpClient.Insert(ctx, req)
	if err != nil {
		return fmt.Errorf("creating NodeGroup %s: %w", a.id, err)
	}

	err = op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("compute NodeGroup %s waiting creation: %w", a.id.String(), err)
	}
	log.Info("successfully created compute NodeGroup in gcp", "name", a.id)

	// Get the created resource
	created, err := a.get(ctx)
	if err != nil {
		return fmt.Errorf("getting NodeGroup %s: %w", a.id, err)
	}

	status := ComputeNodeGroupStatus_v1beta1_FromProto(mapCtx, created)
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *NodeGroupAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating NodeGroup", "name", a.id)

	mapCtx := &direct.MapContext{}
	desired := a.desired.DeepCopy()
	if err := ResolveComputeNodeGroupRefs(ctx, a.reader, a.projectMapper, desired); err != nil {
		return err
	}
	resource := ComputeNodeGroupSpec_v1beta1_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	resource.Name = direct.LazyPtr(a.id.NodeGroup)

	// Handle output-only fields from GCP
	a.assignGCPDefaults(resource, a.actual)

	paths, report, err := common.CompareProtoMessageStructuredDiff(resource, a.actual, common.BasicDiff)
	if err != nil {
		return err
	}

	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		if a.desired.Status.ExternalRef == nil {
			status := ComputeNodeGroupStatus_v1beta1_FromProto(mapCtx, a.actual)
			status.ExternalRef = direct.LazyPtr(a.id.String())
			return updateOp.UpdateStatus(ctx, status, nil)
		}
		return nil
	}

	// We only support updating mutable fields
	// Let's check which fields are mutable. On a NodeGroup, sizes, shareSettings, or autoscaling_policy can be patched or updated.
	// But actually, we can support patching them if the API allows.
	// Let's look at what path we can allow or if we just patch the fields.
	// GCP Compute NodeGroup Patch method updates:
	// - autoscalingPolicy, shareSettings, maintenancePolicy
	// - size can be updated via resize method (similar to reservations)
	// Let's implement Patch for mutable fields.
	report.Object = updateOp.GetUnstructured()
	structuredreporting.ReportDiff(ctx, report)

	updateOp.RecordUpdatingEvent()

	// If there is any size change, we can call resize if the size field is changed.
	// But size is usually immutable or resize is used. Let's see if we should call Patch.
	// Let's call Patch for any changes.
	req := &computepb.PatchNodeGroupRequest{
		Project:           a.id.Project,
		Zone:              a.id.Zone,
		NodeGroup:         a.id.NodeGroup,
		NodeGroupResource: resource,
	}
	op, err := a.gcpClient.Patch(ctx, req)
	if err != nil {
		return fmt.Errorf("patching NodeGroup %s: %w", a.id, err)
	}

	err = op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting patch NodeGroup %s: %w", a.id, err)
	}

	log.V(2).Info("successfully patched compute NodeGroup", "name", a.id)

	// Get the updated resource
	updated, err := a.get(ctx)
	if err != nil {
		return fmt.Errorf("getting NodeGroup %s: %w", a.id, err)
	}

	status := ComputeNodeGroupStatus_v1beta1_FromProto(mapCtx, updated)
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *NodeGroupAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting NodeGroup", "name", a.id)

	req := &computepb.DeleteNodeGroupRequest{
		Project:   a.id.Project,
		Zone:      a.id.Zone,
		NodeGroup: a.id.NodeGroup,
	}
	op, err := a.gcpClient.Delete(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting NodeGroup %s: %w", a.id, err)
	}

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete NodeGroup %s: %w", a.id, err)
	}
	log.Info("successfully deleted compute NodeGroup in gcp", "name", a.id)
	return true, nil
}

func (a *NodeGroupAdapter) get(ctx context.Context) (*computepb.NodeGroup, error) {
	req := &computepb.GetNodeGroupRequest{
		Project:   a.id.Project,
		Zone:      a.id.Zone,
		NodeGroup: a.id.NodeGroup,
	}
	return a.gcpClient.Get(ctx, req)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *NodeGroupAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ComputeNodeGroup{}
	mapCtx := &direct.MapContext{}
	obj.Spec = *ComputeNodeGroupSpec_v1beta1_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	// Note: project is not explicitly in spec but inferred from namespace, we set it if needed
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.NodeGroup)
	u.SetGroupVersionKind(krm.ComputeNodeGroupGVK)

	u.Object = uObj
	return u, nil
}

func (a *NodeGroupAdapter) assignGCPDefaults(desired, actual *computepb.NodeGroup) {
	desired.CreationTimestamp = actual.CreationTimestamp
	desired.Id = actual.Id
	desired.SelfLink = actual.SelfLink
	desired.Kind = actual.Kind
	// The zone format in actual is often full URL. Let's make sure they match.
	desired.Zone = actual.Zone
	desired.Status = actual.Status
}
