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
// proto.service: google.cloud.compute.v1.NetworkEndpointGroups
// proto.message: google.cloud.compute.v1.NetworkEndpointGroup
// crd.type: ComputeNetworkEndpointGroup
// crd.version: v1beta1

package compute

import (
	"context"
	"fmt"

	compute "cloud.google.com/go/compute/apiv1"
	pb "cloud.google.com/go/compute/apiv1/computepb"
	"google.golang.org/protobuf/proto"

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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.ComputeNetworkEndpointGroupGVK, NewComputeNetworkEndpointGroupModel)
}

func NewComputeNetworkEndpointGroupModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &computeNetworkEndpointGroupModel{config: config}, nil
}

var _ directbase.Model = &computeNetworkEndpointGroupModel{}

type computeNetworkEndpointGroupModel struct {
	config *config.ControllerConfig
}

func (m *computeNetworkEndpointGroupModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ComputeNetworkEndpointGroup{}
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

	networkEndpointGroupsClient, err := gcpClient.newNetworkEndpointGroupsClient(ctx)
	if err != nil {
		return nil, err
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	mapCtx := &direct.MapContext{}
	desired := obj.DeepCopy()
	resource := ComputeNetworkEndpointGroupSpec_v1beta1_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	negID := id.(*krm.ComputeNetworkEndpointGroupIdentity)
	resource.Name = proto.String(negID.ComputeNetworkEndpointGroup)
	resource.Zone = proto.String(negID.Zone)

	return &ComputeNetworkEndpointGroupAdapter{
		gcpClient: networkEndpointGroupsClient,
		id:        negID,
		desired:   resource,
		reader:    reader,
	}, nil
}

func (m *computeNetworkEndpointGroupModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type ComputeNetworkEndpointGroupAdapter struct {
	gcpClient *compute.NetworkEndpointGroupsClient
	id        *krm.ComputeNetworkEndpointGroupIdentity
	desired   *pb.NetworkEndpointGroup
	actual    *pb.NetworkEndpointGroup
	reader    client.Reader
}

var _ directbase.Adapter = &ComputeNetworkEndpointGroupAdapter{}

func (a *ComputeNetworkEndpointGroupAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ComputeNetworkEndpointGroup", "name", a.id)

	req := &pb.GetNetworkEndpointGroupRequest{
		Project:              a.id.Project,
		Zone:                 a.id.Zone,
		NetworkEndpointGroup: a.id.ComputeNetworkEndpointGroup,
	}
	actual, err := a.gcpClient.Get(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ComputeNetworkEndpointGroup %q: %w", a.id, err)
	}

	a.actual = actual
	return true, nil
}

func (a *ComputeNetworkEndpointGroupAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating ComputeNetworkEndpointGroup", "name", a.id)

	req := &pb.InsertNetworkEndpointGroupRequest{
		Project:                      a.id.Project,
		Zone:                         a.id.Zone,
		NetworkEndpointGroupResource: a.desired,
	}

	op, err := a.gcpClient.Insert(ctx, req)
	if err != nil {
		return fmt.Errorf("creating ComputeNetworkEndpointGroup %s: %w", a.id.String(), err)
	}

	err = op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("compute ComputeNetworkEndpointGroup %s waiting for creation: %w", a.id.String(), err)
	}

	actual, err := a.gcpClient.Get(ctx, &pb.GetNetworkEndpointGroupRequest{
		Project:              a.id.Project,
		Zone:                 a.id.Zone,
		NetworkEndpointGroup: a.id.ComputeNetworkEndpointGroup,
	})
	if err != nil {
		return fmt.Errorf("getting ComputeNetworkEndpointGroup %s after creation: %w", a.id.String(), err)
	}

	return a.updateStatus(ctx, createOp, actual)
}

func (a *ComputeNetworkEndpointGroupAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating ComputeNetworkEndpointGroup", "name", a.id)

	diffs, err := compareComputeNetworkEndpointGroup(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	if diffs.HasDiff() {
		diffs.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, diffs)
		return fmt.Errorf("ComputeNetworkEndpointGroup is immutable and cannot be updated")
	}

	return a.updateStatus(ctx, updateOp, a.actual)
}

func (a *ComputeNetworkEndpointGroupAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ComputeNetworkEndpointGroup", "name", a.id)

	req := &pb.DeleteNetworkEndpointGroupRequest{
		Project:              a.id.Project,
		Zone:                 a.id.Zone,
		NetworkEndpointGroup: a.id.ComputeNetworkEndpointGroup,
	}
	op, err := a.gcpClient.Delete(ctx, req)
	if err != nil {
		return false, fmt.Errorf("deleting ComputeNetworkEndpointGroup %s: %w", a.id.String(), err)
	}

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("compute ComputeNetworkEndpointGroup %s waiting for deletion: %w", a.id.String(), err)
	}

	return true, nil
}

func (a *ComputeNetworkEndpointGroupAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("ComputeNetworkEndpointGroup not found (not populated in Find)")
	}

	mapCtx := &direct.MapContext{}
	spec := ComputeNetworkEndpointGroupSpec_v1beta1_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	spec.Location = a.id.Zone

	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(krm.ComputeNetworkEndpointGroupGVK)
	u.SetName(a.id.ComputeNetworkEndpointGroup)
	u.SetNamespace(a.id.Project)

	specMap, err := runtime.DefaultUnstructuredConverter.ToUnstructured(spec)
	if err != nil {
		return nil, fmt.Errorf("error converting spec to unstructured: %w", err)
	}

	if err := unstructured.SetNestedMap(u.Object, specMap, "spec"); err != nil {
		return nil, fmt.Errorf("error setting spec on unstructured: %w", err)
	}

	return u, nil
}

func (a *ComputeNetworkEndpointGroupAdapter) updateStatus(ctx context.Context, op directbase.Operation, actual *pb.NetworkEndpointGroup) error {
	status := &krm.ComputeNetworkEndpointGroupStatus{}
	mapCtx := &direct.MapContext{}
	status = ComputeNetworkEndpointGroupStatus_v1beta1_FromProto(mapCtx, actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	return op.UpdateStatus(ctx, status, nil)
}

func compareComputeNetworkEndpointGroup(ctx context.Context, actual, desired *pb.NetworkEndpointGroup) (*structuredreporting.Diff, error) {
	diffs := &structuredreporting.Diff{}

	// Immutable fields: description, defaultPort, network, subnetwork, networkEndpointType
	if actual.GetDescription() != desired.GetDescription() {
		diffs.AddField("spec.description", actual.GetDescription(), desired.GetDescription())
	}
	if actual.GetDefaultPort() != desired.GetDefaultPort() {
		diffs.AddField("spec.defaultPort", actual.GetDefaultPort(), desired.GetDefaultPort())
	}
	if !isSameResource(actual.GetNetwork(), desired.GetNetwork()) {
		diffs.AddField("spec.networkRef", actual.GetNetwork(), desired.GetNetwork())
	}
	if !isSameResource(actual.GetSubnetwork(), desired.GetSubnetwork()) {
		diffs.AddField("spec.subnetworkRef", actual.GetSubnetwork(), desired.GetSubnetwork())
	}
	if actual.GetNetworkEndpointType() != desired.GetNetworkEndpointType() {
		diffs.AddField("spec.networkEndpointType", actual.GetNetworkEndpointType(), desired.GetNetworkEndpointType())
	}

	return diffs, nil
}

func isSameResource(a, b string) bool {
	return apirefs.TrimComputeURIPrefix(a) == apirefs.TrimComputeURIPrefix(b)
}
