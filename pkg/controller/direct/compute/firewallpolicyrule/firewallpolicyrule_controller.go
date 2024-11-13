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
	"strconv"
	"strings"

	"google.golang.org/api/option"

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
	id                     *krm.ComputeFirewallPolicyRuleRef
	firewallPoliciesClient *gcp.FirewallPoliciesClient
	desired                *krm.ComputeFirewallPolicyRule
	actual                 *computepb.FirewallPolicyRule
	reader                 client.Reader
}

var _ directbase.Adapter = &firewallPolicyRuleAdapter{}

func (m *firewallPolicyRuleModel) client(ctx context.Context) (*gcp.FirewallPoliciesClient, error) {
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

func (m *firewallPolicyRuleModel) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.ComputeFirewallPolicyRule{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	firewallPolicyRuleRef, err := krm.NewComputeFirewallPolicyRuleRef(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	firewallPolicyRuleAdapter := &firewallPolicyRuleAdapter{
		id:      firewallPolicyRuleRef,
		desired: obj,
		reader:  reader,
	}

	// Get GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, fmt.Errorf("building gcp client: %w", err)
	}
	firewallPolicyRuleAdapter.firewallPoliciesClient = gcpClient

	return firewallPolicyRuleAdapter, nil
}

func (m *firewallPolicyRuleModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

func (a *firewallPolicyRuleAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ComputeFirewallPolicyRule", "name", a.id.External)

	firewallPolicyRule, err := a.get(ctx)
	if err != nil {
		// When a certain rule does not exist, the error has code 400(bad request) instead of 404(not found)
		// example error message:
		// "Invalid value for field 'priority': '9000'. The firewall policy does not contain a rule at priority 9000.",
		if direct.IsBadRequest(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ComputeFirewallPolicyRule %s: %w", a.id.External, err)
	}
	a.actual = firewallPolicyRule
	return true, nil
}

func (a *firewallPolicyRuleAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	err := resolveDependencies(ctx, a.reader, a.desired)
	if err != nil {
		return err
	}

	log := klog.FromContext(ctx)
	log.V(2).Info("creating ComputeFirewallPolicyRule", "name", a.id.External)

	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()

	firewallPolicyRule := ComputeFirewallPolicyRuleSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	parent, err := a.id.Parent()
	if err != nil {
		return fmt.Errorf("get ComputeFirewallPolicyRule parent %s: %w", a.id.External, err)
	}

	req := &computepb.AddRuleFirewallPolicyRequest{
		FirewallPolicyRuleResource: firewallPolicyRule,
		FirewallPolicy:             parent.FirewallPolicy,
	}
	op, err := a.firewallPoliciesClient.AddRule(ctx, req)

	if err != nil {
		return fmt.Errorf("creating ComputeFirewallPolicyRule %s: %w", a.id.External, err)
	}
	if !op.Done() {
		err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("waiting ComputeFirewallPolicyRule %s create failed: %w", a.id.External, err)
		}
	}
	log.V(2).Info("successfully created ComputeFirewallPolicyRule", "name", a.id.External)

	// Get the created resource
	created, err := a.get(ctx)
	if err != nil {
		return fmt.Errorf("getting ComputeFirewallPolicyRule %s: %w", a.id.External, err)
	}

	status := &krm.ComputeFirewallPolicyRuleStatus{}
	status = ComputeFirewallPolicyRuleStatus_FromProto(mapCtx, created)

	parent, err = a.id.Parent()
	if err != nil {
		return err
	}

	priority := strconv.Itoa(int(*created.Priority))
	externalRef := parent.String() + "/rules/" + priority
	status.ExternalRef = &externalRef
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *firewallPolicyRuleAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	var err error

	err = resolveDependencies(ctx, a.reader, a.desired)
	if err != nil {
		return err
	}

	log := klog.FromContext(ctx)
	log.V(2).Info("updating ComputeFirewallPolicyRule", "name", a.id.External)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	firewallPolicyRule := ComputeFirewallPolicyRuleSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	// The field priority should be removed from the patch request body and included as a query parameter.
	// See API doc: https://cloud.google.com/compute/docs/reference/rest/v1/firewallPolicies/patchRule#query-parameters
	firewallPolicyRule.Priority = nil

	updated := &computepb.FirewallPolicyRule{}

	parent, err := a.id.Parent()
	if err != nil {
		return fmt.Errorf("get ComputeFirewallPolicyRule parent %s: %w", a.id.External, err)
	}

	tokens := strings.Split(a.id.External, "/")
	priority, err := strconv.ParseInt(tokens[5], 10, 32)
	if err != nil {
		return fmt.Errorf("get ComputeFirewallPolicyRule priority %s: %w", a.id.External, err)
	}

	updateReq := &computepb.PatchRuleFirewallPolicyRequest{
		FirewallPolicyRuleResource: firewallPolicyRule,
		FirewallPolicy:             parent.FirewallPolicy,
		Priority:                   direct.PtrTo(int32(priority)),
	}
	op, err := a.firewallPoliciesClient.PatchRule(ctx, updateReq)
	if err != nil {
		return fmt.Errorf("updating ComputeFirewallPolicyRule %s: %w", a.id.External, err)
	}
	if !op.Done() {
		err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("waiting ComputeFirewallPolicyRule %s update failed: %w", a.id.External, err)
		}
	}
	log.V(2).Info("successfully updated ComputeFirewallPolicyRule", "name", a.id.External)

	// Get the updated resource
	updated, err = a.get(ctx)
	if err != nil {
		return fmt.Errorf("getting ComputeFirewallPolicyRule %s: %w", a.id.External, err)
	}

	status := &krm.ComputeFirewallPolicyRuleStatus{}
	status = ComputeFirewallPolicyRuleStatus_FromProto(mapCtx, updated)
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *firewallPolicyRuleAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("firewallPolicyRule %s not found", a.id.External)
	}

	mc := &direct.MapContext{}
	spec := ComputeFirewallPolicyRuleSpec_FromProto(mc, a.actual)
	specObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(spec)
	if err != nil {
		return nil, fmt.Errorf("error converting firewallPolicyRule spec to unstructured: %w", err)
	}

	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(krm.ComputeFirewallPolicyRuleGVK)

	if err := unstructured.SetNestedField(u.Object, specObj, "spec"); err != nil {
		return nil, fmt.Errorf("setting spec: %w", err)
	}

	return u, nil
}

// Delete implements the Adapter interface.
func (a *firewallPolicyRuleAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ComputeFirewallPolicyRule", "name", a.id.External)

	parent, err := a.id.Parent()
	if err != nil {
		return false, fmt.Errorf("get ComputeFirewallPolicyRule parent %s: %w", a.id.External, err)
	}

	tokens := strings.Split(a.id.External, "/")
	priority, err := strconv.ParseInt(tokens[5], 10, 32)
	if err != nil {
		return false, fmt.Errorf("get ComputeFirewallPolicyRule parent %s: %w", a.id.External, err)
	}

	delReq := &computepb.RemoveRuleFirewallPolicyRequest{
		FirewallPolicy: parent.FirewallPolicy,
		Priority:       direct.PtrTo(int32(priority)),
	}
	op, err := a.firewallPoliciesClient.RemoveRule(ctx, delReq)
	if err != nil {
		return false, fmt.Errorf("deleting ComputeFirewallPolicyRule %s: %w", a.id.External, err)
	}
	if !op.Done() {
		err = op.Wait(ctx)
		if err != nil {
			return false, fmt.Errorf("waiting ComputeFirewallPolicyRule %s delete failed: %w", a.id.External, err)
		}
	}
	log.V(2).Info("successfully deleted ComputeFirewallPolicyRule", "name", a.id.External)

	// Get the deleted rules
	_, err = a.get(ctx)
	if err != nil {
		return true, nil
	}
	return true, nil
}

func (a *firewallPolicyRuleAdapter) get(ctx context.Context) (*computepb.FirewallPolicyRule, error) {
	parent, err := a.id.Parent()
	if err != nil {
		return nil, fmt.Errorf("get ComputeFirewallPolicyRule parent %s: %w", a.id.External, err)
	}

	tokens := strings.Split(a.id.External, "/")
	priority, err := strconv.ParseInt(tokens[5], 10, 32)
	if err != nil {
		return nil, fmt.Errorf("get ComputeFirewallPolicyRule parent %s: %w", a.id.External, err)
	}

	getReq := &computepb.GetRuleFirewallPolicyRequest{
		FirewallPolicy: parent.FirewallPolicy,
		Priority:       direct.PtrTo(int32(priority)),
	}
	return a.firewallPoliciesClient.GetRule(ctx, getReq)
}
