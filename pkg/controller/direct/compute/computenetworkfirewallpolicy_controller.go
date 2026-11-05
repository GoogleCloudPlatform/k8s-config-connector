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
	"fmt"

	compute "cloud.google.com/go/compute/apiv1"
	computepb "cloud.google.com/go/compute/apiv1/computepb"
	pb "cloud.google.com/go/compute/apiv1/computepb"
	"google.golang.org/protobuf/proto"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
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
	registry.RegisterModel(krm.ComputeNetworkFirewallPolicyGVK, NewComputeNetworkFirewallPolicyModel)
}

func NewComputeNetworkFirewallPolicyModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &computeNetworkFirewallPolicyModel{config: config}, nil
}

type computeNetworkFirewallPolicyModel struct {
	config *config.ControllerConfig
}

var _ directbase.Model = &computeNetworkFirewallPolicyModel{}

func (m *computeNetworkFirewallPolicyModel) client(ctx context.Context) (*compute.NetworkFirewallPoliciesClient, error) {
	gcpClient, err := newGCPClient(m.config)
	if err != nil {
		return nil, fmt.Errorf("building gcp client: %w", err)
	}
	return gcpClient.newNetworkFirewallPoliciesClient(ctx)
}

func (m *computeNetworkFirewallPolicyModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ComputeNetworkFirewallPolicy{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	idVal, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id, ok := idVal.(*krm.ComputeNetworkFirewallPolicyIdentity)
	if !ok {
		return nil, fmt.Errorf("unexpected identity type: %T", idVal)
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &ComputeNetworkFirewallPolicyAdapter{
		gcpClient: gcpClient,
		id:        id,
		desired:   obj,
		reader:    reader,
	}, nil
}

func (m *computeNetworkFirewallPolicyModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	id := &krm.ComputeNetworkFirewallPolicyIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &ComputeNetworkFirewallPolicyAdapter{
		gcpClient: gcpClient,
		id:        id,
	}, nil
}

type ComputeNetworkFirewallPolicyAdapter struct {
	gcpClient *compute.NetworkFirewallPoliciesClient
	id        *krm.ComputeNetworkFirewallPolicyIdentity
	desired   *krm.ComputeNetworkFirewallPolicy
	actual    *pb.FirewallPolicy
	reader    client.Reader
}

var _ directbase.Adapter = &ComputeNetworkFirewallPolicyAdapter{}

func (a *ComputeNetworkFirewallPolicyAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ComputeNetworkFirewallPolicy", "name", a.id.String())

	actual, err := a.get(ctx)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ComputeNetworkFirewallPolicy %q: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *ComputeNetworkFirewallPolicyAdapter) get(ctx context.Context) (*pb.FirewallPolicy, error) {
	req := &computepb.GetNetworkFirewallPolicyRequest{
		Project:        a.id.Project,
		FirewallPolicy: a.id.FirewallPolicy,
	}
	return a.gcpClient.Get(ctx, req)
}

func (a *ComputeNetworkFirewallPolicyAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating ComputeNetworkFirewallPolicy", "name", a.id.String())

	mapCtx := &direct.MapContext{}
	firewallPolicy := ComputeNetworkFirewallPolicySpec_v1beta1_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	firewallPolicy.Name = direct.PtrTo(a.id.FirewallPolicy)

	req := &computepb.InsertNetworkFirewallPolicyRequest{
		Project:                a.id.Project,
		FirewallPolicyResource: firewallPolicy,
	}

	op, err := a.gcpClient.Insert(ctx, req)
	if err != nil {
		return fmt.Errorf("creating ComputeNetworkFirewallPolicy %s: %w", a.id.String(), err)
	}

	if !op.Done() {
		if err := op.Wait(ctx); err != nil {
			return fmt.Errorf("waiting ComputeNetworkFirewallPolicy %s create failed: %w", a.id.String(), err)
		}
	}
	log.V(2).Info("successfully created ComputeNetworkFirewallPolicy", "name", a.id.String())

	created, err := a.get(ctx)
	if err != nil {
		return fmt.Errorf("getting created ComputeNetworkFirewallPolicy %s: %w", a.id.String(), err)
	}

	return a.updateStatus(ctx, createOp, created)
}

func (a *ComputeNetworkFirewallPolicyAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating ComputeNetworkFirewallPolicy", "name", a.id.String())

	mapCtx := &direct.MapContext{}
	desired := ComputeNetworkFirewallPolicySpec_v1beta1_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	desired.Name = direct.PtrTo(a.id.FirewallPolicy)

	diffs, err := compareNetworkFirewallPolicy(ctx, a.actual, desired)
	if err != nil {
		return err
	}

	if !diffs.HasDiff() {
		log.V(2).Info("no diff detected for ComputeNetworkFirewallPolicy", "name", a.id.String())
		return a.updateStatus(ctx, updateOp, a.actual)
	}

	diffs.Object = updateOp.GetUnstructured()
	structuredreporting.ReportDiff(ctx, diffs)

	resourceToPatch := proto.CloneOf(desired)
	resourceToPatch.Name = nil
	resourceToPatch.Fingerprint = a.actual.Fingerprint

	req := &computepb.PatchNetworkFirewallPolicyRequest{
		Project:                a.id.Project,
		FirewallPolicy:         a.id.FirewallPolicy,
		FirewallPolicyResource: resourceToPatch,
	}

	op, err := a.gcpClient.Patch(ctx, req)
	if err != nil {
		return fmt.Errorf("updating ComputeNetworkFirewallPolicy %s: %w", a.id.String(), err)
	}

	if !op.Done() {
		if err := op.Wait(ctx); err != nil {
			return fmt.Errorf("waiting for ComputeNetworkFirewallPolicy %s update: %w", a.id.String(), err)
		}
	}
	log.V(2).Info("successfully updated ComputeNetworkFirewallPolicy", "name", a.id.String())

	updated, err := a.get(ctx)
	if err != nil {
		return fmt.Errorf("getting ComputeNetworkFirewallPolicy %s: %w", a.id.String(), err)
	}

	return a.updateStatus(ctx, updateOp, updated)
}

func (a *ComputeNetworkFirewallPolicyAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("ComputeNetworkFirewallPolicy %s not found", a.id.String())
	}

	mapCtx := &direct.MapContext{}
	spec := ComputeNetworkFirewallPolicySpec_v1beta1_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	spec.ProjectRef = refs.ProjectRef{
		External: a.id.Project,
	}
	spec.ResourceID = direct.PtrTo(a.id.FirewallPolicy)

	obj := &krm.ComputeNetworkFirewallPolicy{}
	obj.Spec = *spec
	obj.Name = a.id.FirewallPolicy
	obj.SetGroupVersionKind(krm.ComputeNetworkFirewallPolicyGVK)

	u, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, fmt.Errorf("converting to unstructured: %w", err)
	}
	return &unstructured.Unstructured{Object: u}, nil
}

func (a *ComputeNetworkFirewallPolicyAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ComputeNetworkFirewallPolicy", "name", a.id.String())

	req := &computepb.DeleteNetworkFirewallPolicyRequest{
		Project:        a.id.Project,
		FirewallPolicy: a.id.FirewallPolicy,
	}

	op, err := a.gcpClient.Delete(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting ComputeNetworkFirewallPolicy %s: %w", a.id.String(), err)
	}

	if !op.Done() {
		if err := op.Wait(ctx); err != nil {
			if direct.IsNotFound(err) {
				return true, nil
			}
			return false, fmt.Errorf("waiting for deletion of ComputeNetworkFirewallPolicy %s: %w", a.id.String(), err)
		}
	}
	log.V(2).Info("successfully deleted ComputeNetworkFirewallPolicy", "name", a.id.String())

	return true, nil
}

func (a *ComputeNetworkFirewallPolicyAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.FirewallPolicy) error {
	mapCtx := &direct.MapContext{}
	status := ComputeNetworkFirewallPolicyStatus_v1beta1_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return op.UpdateStatus(ctx, status, nil)
}

func compareNetworkFirewallPolicy(ctx context.Context, actual, desired *pb.FirewallPolicy) (*structuredreporting.Diff, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, ComputeNetworkFirewallPolicySpec_v1beta1_FromProto, ComputeNetworkFirewallPolicySpec_v1beta1_ToProto)
	if err != nil {
		return nil, err
	}
	maskedActual.Name = desired.Name

	clonedDesired := proto.CloneOf(desired)

	diffs, _, err := common.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, err
	}
	return diffs, nil
}
