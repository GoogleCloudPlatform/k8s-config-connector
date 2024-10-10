// Copyright 2024 Google LLC
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

package firewallpolicyrule

import (
	"context"
	"fmt"

	gcp "cloud.google.com/go/compute/apiv1"
	computepb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const ctrlName = "firewallpolicyrule-controller"

func init() {
	registry.RegisterModel(krm.ComputeFirewallPolicyRuleGVK, NewFirewallPolicyRuleModel)
}

func NewFirewallPolicyRuleModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &firewallPolicyRuleModel{config: config}, nil
}

type firewallPolicyRuleModel struct {
	config *config.ControllerConfig
}

// model implements the Model interface.
var _ directbase.Model = &firewallPolicyRuleModel{}

type firewallPolicyRuleAdapter struct {
	firewallPolicy         string
	priority               int64
	firewallPoliciesClient *gcp.FirewallPoliciesClient
	desired                *krm.ComputeFirewallPolicyRule
	actual                 *computepb.FirewallPolicyRule
	reader                 client.Reader
}

var _ directbase.Adapter = &firewallPolicyRuleAdapter{}

func (m *firewallPolicyRuleModel) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.ComputeFirewallPolicyRule{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	// Set label managed-by-cnrm: true
	obj.ObjectMeta.Labels["managed-by-cnrm"] = "true"

	// Get firewall policy
	firewallPolicyRef, err := ResolveComputeFirewallPolicy(ctx, reader, obj, obj.Spec.FirewallPolicyRef)
	if err != nil {
		return nil, err

	}
	obj.Spec.FirewallPolicyRef.External = firewallPolicyRef.External
	firewallPolicy := obj.Spec.FirewallPolicyRef.External

	// Get priority
	priority := obj.Spec.Priority

	firewallPolicyRuleAdapter := &firewallPolicyRuleAdapter{
		firewallPolicy: firewallPolicy,
		priority:       priority,
		desired:        obj,
		reader:         reader,
	}

	// Get GCP client
	gcpClient, err := newGCPClient(ctx, m.config)
	if err != nil {
		return nil, fmt.Errorf("building gcp client: %w", err)
	}

	firewallPoliciesClient, err := gcpClient.firewallPoliciesClient(ctx)
	if err != nil {
		return nil, err
	}
	firewallPolicyRuleAdapter.firewallPoliciesClient = firewallPoliciesClient

	return firewallPolicyRuleAdapter, nil
}

func (m *firewallPolicyRuleModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

func (a *firewallPolicyRuleAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("getting ComputeFirewallPolicyRule", "priority", a.priority)

	firewallPolicyRule, err := a.get(ctx)
	if err != nil {
		// When a certain rule does not exist, the error has code 400(invalid) instead of 404(not found)
		// example error message:
		// "Invalid value for field 'priority': '9000'. The firewall policy does not contain a rule at priority 9000.",
		if direct.IsInvalidValue(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ComputeFirewallPolicyRule %d: %w", a.priority, err)
	}
	a.actual = firewallPolicyRule
	return true, nil
}

func (a *firewallPolicyRuleAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	u := createOp.GetUnstructured()
	var err error

	err = resolveDependencies(ctx, a.reader, a.desired)
	if err != nil {
		return err
	}

	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("creating ComputeFirewallPolicyRule", "priority", a.priority)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()

	firewallPolicyRule := ComputeFirewallPolicyRuleSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &computepb.AddRuleFirewallPolicyRequest{
		FirewallPolicyRuleResource: firewallPolicyRule,
		FirewallPolicy:             a.firewallPolicy,
	}
	op, err := a.firewallPoliciesClient.AddRule(ctx, req)

	if err != nil {
		return fmt.Errorf("creating ComputeFirewallPolicyRule %d: %w", a.priority, err)
	}
	if !op.Done() {
		err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("waiting ComputeFirewallPolicyRule %d create failed: %w", a.priority, err)
		}
	}
	log.V(2).Info("successfully created ComputeFirewallPolicyRule", "priority", a.priority)

	// Get the created resource
	created := &computepb.FirewallPolicyRule{}
	created, err = a.get(ctx)
	if err != nil {
		return fmt.Errorf("getting ComputeFirewallPolicyRule %d: %w", a.priority, err)
	}

	status := &krm.ComputeFirewallPolicyRuleStatus{
		RuleTupleCount: direct.PtrTo(int64(*created.RuleTupleCount)),
		Kind:           direct.PtrTo("compute#firewallPolicyRule"),
	}
	return setStatus(u, status)
}

func (a *firewallPolicyRuleAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	u := updateOp.GetUnstructured()
	var err error

	err = resolveDependencies(ctx, a.reader, a.desired)
	if err != nil {
		return err
	}

	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("updating ComputeFirewallPolicyRule", "priority", a.priority)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	firewallPolicyRule := ComputeFirewallPolicyRuleSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	firewallPolicyRule.Priority = nil

	op := &gcp.Operation{}
	updated := &computepb.FirewallPolicyRule{}

	updateReq := &computepb.PatchRuleFirewallPolicyRequest{
		FirewallPolicyRuleResource: firewallPolicyRule,
		FirewallPolicy:             a.firewallPolicy,
		Priority:                   direct.PtrTo(int32(a.priority)),
	}
	op, err = a.firewallPoliciesClient.PatchRule(ctx, updateReq)
	if err != nil {
		return fmt.Errorf("updating ComputeFirewallPolicyRule %d: %w", a.priority, err)
	}
	if !op.Done() {
		err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("waiting ComputeFirewallPolicyRule %d update failed: %w", a.priority, err)
		}
	}
	log.V(2).Info("successfully updated ComputeFirewallPolicyRule", "priority", a.priority)

	// Get the updated resource
	updated, err = a.get(ctx)
	if err != nil {
		return fmt.Errorf("getting ComputeFirewallPolicyRule %d: %w", a.priority, err)
	}

	status := &krm.ComputeFirewallPolicyRuleStatus{
		RuleTupleCount: direct.PtrTo(int64(*updated.RuleTupleCount)),
		Kind:           direct.PtrTo("compute#firewallPolicyRule"),
	}
	return setStatus(u, status)
}

func (a *firewallPolicyRuleAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("firewallPolicyRule %d not found", a.priority)
	}

	mc := &direct.MapContext{}
	spec := ComputeFirewallPolicyRuleSpec_FromProto(mc, a.actual)
	specObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(spec)
	if err != nil {
		return nil, fmt.Errorf("error converting firewallPolicyRule spec to unstructured: %w", err)
	}

	u := &unstructured.Unstructured{
		Object: make(map[string]interface{}),
	}
	u.SetGroupVersionKind(krm.ComputeFirewallPolicyRuleGVK)

	if err := unstructured.SetNestedField(u.Object, specObj, "spec"); err != nil {
		return nil, fmt.Errorf("setting spec: %w", err)
	}

	return u, nil
}

// Delete implements the Adapter interface.
func (a *firewallPolicyRuleAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("deleting ComputeFirewallPolicyRule", "priority", a.priority)

	delReq := &computepb.RemoveRuleFirewallPolicyRequest{
		FirewallPolicy: a.firewallPolicy,
		Priority:       direct.PtrTo(int32(a.priority)),
	}
	op, err := a.firewallPoliciesClient.RemoveRule(ctx, delReq)

	if err != nil {
		return false, fmt.Errorf("deleting ComputeFirewallPolicyRule %d: %w", a.priority, err)
	}
	if !op.Done() {
		err = op.Wait(ctx)
		if err != nil {
			return false, fmt.Errorf("waiting ComputeFirewallPolicyRule %d delete failed: %w", a.priority, err)
		}
	}
	log.V(2).Info("successfully deleted ComputeFirewallPolicyRule", "priority", a.priority)

	// Get the deleted rules
	_, err = a.get(ctx)
	if err != nil {
		return true, nil
	}
	return true, nil
}

func (a *firewallPolicyRuleAdapter) get(ctx context.Context) (*computepb.FirewallPolicyRule, error) {
	getReq := &computepb.GetRuleFirewallPolicyRequest{
		FirewallPolicy: a.firewallPolicy,
		Priority:       direct.PtrTo(int32(a.priority)),
	}
	return a.firewallPoliciesClient.GetRule(ctx, getReq)
}

func setStatus(u *unstructured.Unstructured, typedStatus any) error {
	status, err := runtime.DefaultUnstructuredConverter.ToUnstructured(typedStatus)
	if err != nil {
		return fmt.Errorf("error converting status to unstructured: %w", err)
	}

	old, _, _ := unstructured.NestedMap(u.Object, "status")
	if old != nil {
		status["conditions"] = old["conditions"]
		status["observedGeneration"] = old["observedGeneration"]
		status["externalRef"] = old["externalRef"]
	}

	u.Object["status"] = status

	return nil
}
