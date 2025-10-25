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

package compute

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"k8s.io/klog/v2"

	gcp "cloud.google.com/go/compute/apiv1"
	computepb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/label"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.ComputeForwardingRuleGVK, NewForwardingRuleModel)
}

func NewForwardingRuleModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &forwardingRuleModel{config: config}, nil
}

type forwardingRuleModel struct {
	config *config.ControllerConfig
}

// model implements the Model interface.
var _ directbase.Model = &forwardingRuleModel{}

type forwardingRuleAdapter struct {
	id                          *krm.ForwardingRuleIdentity
	forwardingRulesClient       *gcp.ForwardingRulesClient
	globalForwardingRulesClient *gcp.GlobalForwardingRulesClient
	desired                     *krm.ComputeForwardingRule
	actual                      *computepb.ForwardingRule
	reader                      client.Reader
}

var _ directbase.Adapter = &forwardingRuleAdapter{}

func (m *forwardingRuleModel) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.ComputeForwardingRule{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewForwardingRuleIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Handle TF default values
	if obj.Spec.LoadBalancingScheme == nil {
		obj.Spec.LoadBalancingScheme = direct.LazyPtr("EXTERNAL")
	}

	forwardingRuleAdapter := &forwardingRuleAdapter{
		id:      id,
		desired: obj,
		reader:  reader,
	}

	// Get region
	parent := id.Parent()
	region := parent.Location

	// Get GCP client
	gcpClient, err := newGCPClient(m.config)
	if err != nil {
		return nil, fmt.Errorf("building gcp client: %w", err)
	}
	if region == "global" {
		globalForwardingRulesClient, err := gcpClient.newGlobalForwardingRuleClient(ctx)
		if err != nil {
			return nil, err
		}
		forwardingRuleAdapter.globalForwardingRulesClient = globalForwardingRulesClient
	} else {
		forwardingRulesClient, err := gcpClient.forwardingRuleClient(ctx)
		if err != nil {
			return nil, err
		}
		forwardingRuleAdapter.forwardingRulesClient = forwardingRulesClient
	}
	return forwardingRuleAdapter, nil
}

func (m *forwardingRuleModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

func (a *forwardingRuleAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ComputeForwardingRule", "name", a.id)

	var err error
	forwardingRule := &computepb.ForwardingRule{}
	forwardingRule, err = a.get(ctx)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ComputeForwardingRule %q: %w", a.id, err)
	}
	a.actual = forwardingRule
	return true, nil
}

func (a *forwardingRuleAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	u := createOp.GetUnstructured()
	var err error

	err = resolveForwardingRuleRefs(ctx, a.reader, a.desired)
	if err != nil {
		return err
	}

	log := klog.FromContext(ctx)
	log.V(2).Info("creating ComputeForwardingRule", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	sanitizedLabels := label.NewGCPLabelsFromK8sLabels(desired.Labels)

	forwardingRule := ComputeForwardingRuleSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	forwardingRule.Name = direct.LazyPtr(a.id.ID())
	forwardingRule.Labels = sanitizedLabels

	// API restriction: Cannot set labels during creation(by POST). But it can be set later by PATCH SetLabels.
	// API error message: Labels are invalid in Private Service Connect Forwarding Rule.
	// See GH issue for details: https://github.com/hashicorp/terraform-provider-google/issues/16255
	target := direct.ValueOf(forwardingRule.Target)
	// Remove labels for psc forwarding rule
	if target == "all-apis" || target == "vpc-sc" || strings.Contains(target, "/serviceAttachments/") {
		forwardingRule.Labels = nil
	}

	// Create forwarding rule(labels are not set during Insert)
	op := &gcp.Operation{}
	if a.id.Parent().Location == "global" {
		req := &computepb.InsertGlobalForwardingRuleRequest{
			ForwardingRuleResource: forwardingRule,
			Project:                a.id.Parent().ProjectID,
		}
		op, err = a.globalForwardingRulesClient.Insert(ctx, req)
	} else {
		req := &computepb.InsertForwardingRuleRequest{
			ForwardingRuleResource: forwardingRule,
			Region:                 a.id.Parent().Location,
			Project:                a.id.Parent().ProjectID,
		}
		op, err = a.forwardingRulesClient.Insert(ctx, req)
	}
	if err != nil {
		return fmt.Errorf("creating ComputeForwardingRule %s: %w", a.id, err)
	}
	if !op.Done() {
		err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("waiting ComputeForwardingRule %s create failed: %w", a.id, err)
		}
	}
	log.V(2).Info("successfully created ComputeForwardingRule", "name", a.id)

	// Get the created resource
	created := &computepb.ForwardingRule{}
	created, err = a.get(ctx)
	if err != nil {
		return fmt.Errorf("getting ComputeForwardingRule %q: %w", a.id, err)
	}

	// Set labels for the created forwarding rule
	// Add labels back for psc forwarding rule
	if target == "all-apis" || target == "vpc-sc" || strings.Contains(target, "/serviceAttachments/") {
		forwardingRule.Labels = sanitizedLabels
	}
	if forwardingRule.Labels != nil {
		op, err := a.setLabels(ctx, created.LabelFingerprint, forwardingRule.Labels)
		if err != nil {
			return fmt.Errorf("adding ComputeForwardingRule labels %s: %w", a.id, err)
		}
		if !op.Done() {
			err = op.Wait(ctx)
			if err != nil {
				return fmt.Errorf("waiting ComputeForwardingRule %s add labels failed: %w", a.id, err)
			}
		}
		log.V(2).Info("successfully added ComputeForwardingRule labels", "name", a.id)

		// Get the created resource with label added
		created, err = a.get(ctx)
		if err != nil {
			return fmt.Errorf("getting ComputeForwardingRule %q: %w", a.id, err)
		}
	}

	status := &krm.ComputeForwardingRuleStatus{
		LabelFingerprint:  created.LabelFingerprint,
		CreationTimestamp: created.CreationTimestamp,
		SelfLink:          created.SelfLink,
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return setStatus(u, status)
}

func (a *forwardingRuleAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	u := updateOp.GetUnstructured()
	var err error

	if a.id.ID() == "" {
		return fmt.Errorf("resourceID is empty")
	}

	err = resolveForwardingRuleRefs(ctx, a.reader, a.desired)
	if err != nil {
		return err
	}

	log := klog.FromContext(ctx)
	log.V(2).Info("updating ComputeForwardingRule", "name", a.id.ID())
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	sanitizedLabels := label.NewGCPLabelsFromK8sLabels(desired.Labels)
	forwardingRule := ComputeForwardingRuleSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	forwardingRule.Name = direct.LazyPtr(a.id.ID())
	forwardingRule.Labels = sanitizedLabels

	op := &gcp.Operation{}
	updated := &computepb.ForwardingRule{}
	if !reflect.DeepEqual(forwardingRule.AllowGlobalAccess, a.actual.AllowGlobalAccess) {
		// To match the request body in TF-controller log
		// https://github.com/hashicorp/terraform-provider-google/blob/main/google/services/compute/resource_compute_forwarding_rule.go#L1151
		reqBody := &computepb.ForwardingRule{AllowGlobalAccess: forwardingRule.AllowGlobalAccess}
		if a.id.Parent().Location == "global" {
			// TF does not support allowGlobalAccess field for global forwarding rule
			// Underlying API as well, error message: `Field allow-global-access is only supported for regional INTERNAL
			// forwarding rules with backend service/target instance or regional INTERNAL_MANAGED forwarding rules.`
			forwardingRule.AllowGlobalAccess = nil
		} else {
			patchReq := &computepb.PatchForwardingRuleRequest{
				ForwardingRule:         a.id.ID(),
				ForwardingRuleResource: reqBody,
				Project:                a.id.Parent().ProjectID,
				Region:                 a.id.Parent().Location,
			}
			op, err = a.forwardingRulesClient.Patch(ctx, patchReq)
		}
		if err != nil {
			return fmt.Errorf("updating ComputeForwardingRule %s: %w", a.id, err)
		}
		err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("waiting ComputeForwardingRule %s update failed: %w", a.id, err)
		}
		log.V(2).Info("successfully updated ComputeForwardingRule", "name", a.id)
	}

	// Use setTarget and setLabels to update target and labels fields.
	if !reflect.DeepEqual(forwardingRule.Labels, a.actual.Labels) {
		op, err := a.setLabels(ctx, a.actual.LabelFingerprint, forwardingRule.Labels)
		if err != nil {
			return fmt.Errorf("updating ComputeForwardingRule labels %s: %w", a.id, err)
		}
		if !op.Done() {
			err = op.Wait(ctx)
			if err != nil {
				return fmt.Errorf("waiting ComputeForwardingRule %s update labels failed: %w", a.id, err)
			}
		}
		log.V(2).Info("successfully updated ComputeForwardingRule labels", "name", a.id)
	}

	// setTarget request is sent when there are updates to target.
	// IsSelfLinkEqual is a special handling to avoid reconciliation discrepancies caused by resources and
	// their dependencies being managed by different controllers.
	// This can be removed once all Compute resources are migrated to direct controller.
	if !IsSelfLinkEqual(forwardingRule.Target, a.actual.Target) {
		if a.id.Parent().Location == "global" {
			setTargetReq := &computepb.SetTargetGlobalForwardingRuleRequest{
				ForwardingRule:          a.id.ID(),
				TargetReferenceResource: &computepb.TargetReference{Target: forwardingRule.Target},
				Project:                 a.id.Parent().ProjectID,
			}
			op, err = a.globalForwardingRulesClient.SetTarget(ctx, setTargetReq)
		} else {
			setTargetReq := &computepb.SetTargetForwardingRuleRequest{
				ForwardingRule:          a.id.ID(),
				TargetReferenceResource: &computepb.TargetReference{Target: forwardingRule.Target},
				Project:                 a.id.Parent().ProjectID,
				Region:                  a.id.Parent().Location,
			}
			op, err = a.forwardingRulesClient.SetTarget(ctx, setTargetReq)
		}
		if err != nil {
			return fmt.Errorf("updating ComputeForwardingRule target %s: %w", a.id, err)
		}
		if !op.Done() {
			err = op.Wait(ctx)
			if err != nil {
				return fmt.Errorf("waiting ComputeForwardingRule %s update target failed: %w", a.id, err)
			}
		}
		log.V(2).Info("successfully updated ComputeForwardingRule target", "name", a.id)
	}
	// Get the updated resource
	updated, err = a.get(ctx)
	if err != nil {
		return fmt.Errorf("getting ComputeForwardingRule %q: %w", a.id.ID(), err)
	}

	status := &krm.ComputeForwardingRuleStatus{
		LabelFingerprint:  updated.LabelFingerprint,
		CreationTimestamp: updated.CreationTimestamp,
		SelfLink:          updated.SelfLink,
	}
	return setStatus(u, status)
}

func (a *forwardingRuleAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("forwardingrule %q not found", a.id)
	}

	mc := &direct.MapContext{}
	spec := ComputeForwardingRuleSpec_FromProto(mc, a.actual)
	specObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(spec)
	if err != nil {
		return nil, fmt.Errorf("error converting forwardingrule spec to unstructured: %w", err)
	}

	u := &unstructured.Unstructured{
		Object: make(map[string]interface{}),
	}
	u.SetName(a.id.ID())
	u.SetGroupVersionKind(krm.ComputeForwardingRuleGVK)
	u.SetLabels(a.actual.Labels)

	if err := unstructured.SetNestedField(u.Object, specObj, "spec"); err != nil {
		return nil, fmt.Errorf("setting spec: %w", err)
	}

	return u, nil
}

// Delete implements the Adapter interface.
func (a *forwardingRuleAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	if a.id.ID() == "" {
		return false, fmt.Errorf("resourceID is empty")
	}

	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ComputeForwardingRule", "name", a.id.ID())

	var err error
	op := &gcp.Operation{}
	if a.id.Parent().Location == "global" {
		req := &computepb.DeleteGlobalForwardingRuleRequest{
			ForwardingRule: a.id.ID(),
			Project:        a.id.Parent().ProjectID,
		}
		op, err = a.globalForwardingRulesClient.Delete(ctx, req)
	} else {
		req := &computepb.DeleteForwardingRuleRequest{
			ForwardingRule: a.id.ID(),
			Region:         a.id.Parent().Location,
			Project:        a.id.Parent().ProjectID,
		}
		op, err = a.forwardingRulesClient.Delete(ctx, req)
	}
	if err != nil {
		return false, fmt.Errorf("deleting ComputeForwardingRule %s: %w", a.id.ID(), err)
	}
	if !op.Done() {
		err = op.Wait(ctx)
		if err != nil {
			return false, fmt.Errorf("waiting ComputeForwardingRule %s delete failed: %w", a.id.ID(), err)
		}
	}
	log.V(2).Info("successfully deleted ComputeForwardingRule", "name", a.id.ID())
	return true, nil
}

func (a *forwardingRuleAdapter) get(ctx context.Context) (*computepb.ForwardingRule, error) {
	if a.id.Parent().Location == "global" {
		getReq := &computepb.GetGlobalForwardingRuleRequest{
			ForwardingRule: a.id.ID(),
			Project:        a.id.Parent().ProjectID,
		}
		return a.globalForwardingRulesClient.Get(ctx, getReq)
	} else {
		getReq := &computepb.GetForwardingRuleRequest{
			ForwardingRule: a.id.ID(),
			Region:         a.id.Parent().Location,
			Project:        a.id.Parent().ProjectID,
		}
		return a.forwardingRulesClient.Get(ctx, getReq)
	}
}

func (a *forwardingRuleAdapter) setLabels(ctx context.Context, fingerprint *string, labels map[string]string) (*gcp.Operation, error) {
	op := &gcp.Operation{}
	var err error
	if a.id.Parent().Location == "global" {
		setLabelsReq := &computepb.SetLabelsGlobalForwardingRuleRequest{
			Resource:                       a.id.ID(),
			GlobalSetLabelsRequestResource: &computepb.GlobalSetLabelsRequest{LabelFingerprint: fingerprint, Labels: labels},
			Project:                        a.id.Parent().ProjectID,
		}
		op, err = a.globalForwardingRulesClient.SetLabels(ctx, setLabelsReq)
	} else {
		setLabelsReq := &computepb.SetLabelsForwardingRuleRequest{
			Resource:                       a.id.ID(),
			RegionSetLabelsRequestResource: &computepb.RegionSetLabelsRequest{LabelFingerprint: fingerprint, Labels: labels},
			Project:                        a.id.Parent().ProjectID,
			Region:                         a.id.Parent().Location,
		}
		op, err = a.forwardingRulesClient.SetLabels(ctx, setLabelsReq)
	}
	return op, err
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
