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
	"strconv"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"google.golang.org/protobuf/proto"

	"google.golang.org/api/option"

	gcp "cloud.google.com/go/compute/apiv1"
	computepb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/tags"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.ComputeFirewallPolicyGVK, NewComputeFirewallPolicyModel)
}

func NewComputeFirewallPolicyModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &computeFirewallPolicyModel{config: config}, nil
}

type computeFirewallPolicyModel struct {
	config *config.ControllerConfig
}

var _ directbase.Model = &computeFirewallPolicyModel{}

type computeFirewallPolicyAdapter struct {
	id                     *krm.ComputeFirewallPolicyIdentity
	firewallPoliciesClient *gcp.FirewallPoliciesClient
	desired                *computepb.FirewallPolicy
	actual                 *computepb.FirewallPolicy
	reader                 client.Reader
	parent                 string
}

var _ directbase.Adapter = &computeFirewallPolicyAdapter{}

func (m *computeFirewallPolicyModel) client(ctx context.Context) (*gcp.FirewallPoliciesClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewFirewallPoliciesRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building FirewallPolicy client: %w", err)
	}
	return gcpClient, err
}

func (m *computeFirewallPolicyModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ComputeFirewallPolicy{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	// We must first resolve/normalize the references in AdapterForObject
	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	mapCtx := &direct.MapContext{}
	desired := ComputeFirewallPolicySpec_v1beta1_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	var parent string
	if obj.Spec.FolderRef != nil {
		parent = obj.Spec.FolderRef.External
	} else if obj.Spec.OrganizationRef != nil {
		parent = obj.Spec.OrganizationRef.External
	}

	if parent == "" {
		return nil, fmt.Errorf("one of FolderRef or OrganizationRef must be specified")
	}

	adapter := &computeFirewallPolicyAdapter{
		desired: desired,
		reader:  reader,
		parent:  parent,
	}

	rawID, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := rawID.(*krm.ComputeFirewallPolicyIdentity)
	adapter.id = id

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, fmt.Errorf("building gcp client: %w", err)
	}
	adapter.firewallPoliciesClient = gcpClient

	return adapter, nil
}

func (m *computeFirewallPolicyModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

func (a *computeFirewallPolicyAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ComputeFirewallPolicy", "name", a.id)

	if a.id.FirewallPolicy == "" {
		return false, nil
	}

	firewallPolicy, err := a.get(ctx)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ComputeFirewallPolicy %s: %w", a.id, err)
	}

	a.actual = firewallPolicy
	return true, nil
}

func (a *computeFirewallPolicyAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating ComputeFirewallPolicy", "name", a.id)

	req := &computepb.InsertFirewallPolicyRequest{
		ParentId:               a.parent,
		FirewallPolicyResource: a.desired,
	}

	op, err := a.firewallPoliciesClient.Insert(ctx, req)
	if err != nil {
		return fmt.Errorf("creating ComputeFirewallPolicy: %w", err)
	}
	if !op.Done() {
		err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("waiting ComputeFirewallPolicy create failed: %w", err)
		}
	}
	log.V(2).Info("successfully created ComputeFirewallPolicy")

	var policyID string
	if op.Proto() != nil {
		if op.Proto().TargetId != nil {
			policyID = strconv.FormatUint(*op.Proto().TargetId, 10)
		} else if op.Proto().TargetLink != nil {
			link := *op.Proto().TargetLink
			tokens := strings.Split(link, "/")
			policyID = tokens[len(tokens)-1]
		}
	}

	if policyID == "" {
		return fmt.Errorf("could not determine the firewall policy ID from operation")
	}

	a.id.FirewallPolicy = policyID
	created, err := a.get(ctx)
	if err != nil {
		return fmt.Errorf("getting created ComputeFirewallPolicy %s: %w", a.id, err)
	}

	return a.updateStatus(ctx, createOp, created)
}

func (a *computeFirewallPolicyAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating ComputeFirewallPolicy", "name", a.id)

	diffs, err := compareFirewallPolicy(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	if !diffs.HasDiff() {
		log.V(2).Info("no diff detected for ComputeFirewallPolicy", "name", a.id)
		return a.updateStatus(ctx, updateOp, a.actual)
	}

	structuredreporting.ReportDiff(ctx, diffs)

	req := &computepb.PatchFirewallPolicyRequest{
		FirewallPolicy:         a.id.FirewallPolicy,
		FirewallPolicyResource: a.desired,
	}

	op, err := a.firewallPoliciesClient.Patch(ctx, req)
	if err != nil {
		return fmt.Errorf("updating ComputeFirewallPolicy %s: %w", a.id, err)
	}
	if !op.Done() {
		err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("waiting ComputeFirewallPolicy %s update failed: %w", a.id, err)
		}
	}
	log.V(2).Info("successfully updated ComputeFirewallPolicy", "name", a.id)

	// Get the updated resource
	updated, err := a.get(ctx)
	if err != nil {
		return fmt.Errorf("getting ComputeFirewallPolicy %s: %w", a.id, err)
	}

	return a.updateStatus(ctx, updateOp, updated)
}

func (a *computeFirewallPolicyAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("ComputeFirewallPolicy %s not found", a.id)
	}

	mc := &direct.MapContext{}
	spec := ComputeFirewallPolicySpec_v1beta1_FromProto(mc, a.actual)

	if a.actual.Parent != nil {
		parent := *a.actual.Parent
		tokens := strings.Split(parent, "/")
		if len(tokens) == 2 {
			if tokens[0] == "folders" {
				spec.FolderRef = &refs.FolderRef{External: parent}
			} else if tokens[0] == "organizations" {
				spec.OrganizationRef = &refs.OrganizationRef{External: parent}
			}
		}
	}

	specObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(spec)
	if err != nil {
		return nil, fmt.Errorf("error converting ComputeFirewallPolicy spec to unstructured: %w", err)
	}

	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(krm.ComputeFirewallPolicyGVK)

	if err := unstructured.SetNestedField(u.Object, specObj, "spec"); err != nil {
		return nil, fmt.Errorf("setting spec: %w", err)
	}

	return u, nil
}

func (a *computeFirewallPolicyAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ComputeFirewallPolicy", "name", a.id)

	if a.id.FirewallPolicy == "" {
		return false, fmt.Errorf("cannot delete ComputeFirewallPolicy without an ID")
	}

	req := &computepb.DeleteFirewallPolicyRequest{
		FirewallPolicy: a.id.FirewallPolicy,
	}

	op, err := a.firewallPoliciesClient.Delete(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting ComputeFirewallPolicy %s: %w", a.id, err)
	}
	if !op.Done() {
		err = op.Wait(ctx)
		if err != nil {
			return false, fmt.Errorf("waiting ComputeFirewallPolicy %s delete failed: %w", a.id, err)
		}
	}
	log.V(2).Info("successfully deleted ComputeFirewallPolicy", "name", a.id)
	return true, nil
}

func (a *computeFirewallPolicyAdapter) get(ctx context.Context) (*computepb.FirewallPolicy, error) {
	getReq := &computepb.GetFirewallPolicyRequest{
		FirewallPolicy: a.id.FirewallPolicy,
	}
	return a.firewallPoliciesClient.Get(ctx, getReq)
}

func (a *computeFirewallPolicyAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *computepb.FirewallPolicy) error {
	mapCtx := &direct.MapContext{}
	status := ComputeFirewallPolicyStatus_v1beta1_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return op.UpdateStatus(ctx, status, nil)
}

func compareFirewallPolicy(ctx context.Context, actual, desired *computepb.FirewallPolicy) (*structuredreporting.Diff, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, ComputeFirewallPolicySpec_v1beta1_FromProto, ComputeFirewallPolicySpec_v1beta1_ToProto)
	if err != nil {
		return nil, err
	}
	maskedActual.Name = desired.Name
	maskedActual.Id = desired.Id
	maskedActual.Parent = desired.Parent

	clonedDesired := proto.CloneOf(desired)

	diffs, _, err := tags.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, err
	}
	return diffs, nil
}
