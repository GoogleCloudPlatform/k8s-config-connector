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

	"google.golang.org/protobuf/proto"
	"k8s.io/klog/v2"

	gcp "cloud.google.com/go/compute/apiv1"
	computepb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/export"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/label"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
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
	id                          *krm.ComputeForwardingRuleIdentity
	forwardingRulesClient       *gcp.ForwardingRulesClient
	globalForwardingRulesClient *gcp.GlobalForwardingRulesClient
	desired                     *computepb.ForwardingRule
	actual                      *computepb.ForwardingRule
	reader                      client.Reader
	desiredStatusTarget         *string
}

var _ directbase.Adapter = &forwardingRuleAdapter{}

func (m *forwardingRuleModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ComputeForwardingRule{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.ComputeForwardingRuleIdentity)

	// Handle TF default values
	if obj.Spec.LoadBalancingScheme == nil {
		obj.Spec.LoadBalancingScheme = direct.LazyPtr("EXTERNAL")
	}

	// resolveForwardingRuleRefs is required because certain spec fields on ComputeForwardingRule (like ipAddress, target, etc.)
	// reference other GCP resources where we need to retrieve custom values (such as status.observedState.address
	// or status.selfLink) that are not automatically mapped by standard reference normalization.
	// This must be run before common.NormalizeReferences because it may clear or transform some references (e.g. converting
	// MemorystoreInstanceRef into ServiceAttachmentRef).
	if err := resolveForwardingRuleRefs(ctx, reader, obj); err != nil {
		return nil, fmt.Errorf("resolving remaining references: %w", err)
	}

	// Normalize standard KCC references.
	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	mapCtx := &direct.MapContext{}
	desired := ComputeForwardingRuleSpec_v1beta1_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	desired.Labels = label.NewGCPLabelsFromK8sLabels(obj.Labels)

	forwardingRuleAdapter := &forwardingRuleAdapter{
		id:                  id,
		desired:             desired,
		reader:              reader,
		desiredStatusTarget: obj.Status.Target,
	}

	// Get GCP client
	gcpClient, err := newGCPClient(m.config)
	if err != nil {
		return nil, fmt.Errorf("building gcp client: %w", err)
	}
	if id.IsGlobal() {
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
	id := &krm.ComputeForwardingRuleIdentity{}
	if err := id.FromExternal(url); err != nil {
		// Not recognized
		return nil, nil
	}

	// Get GCP client
	gcpClient, err := newGCPClient(m.config)
	if err != nil {
		return nil, fmt.Errorf("building gcp client: %w", err)
	}

	forwardingRuleAdapter := &forwardingRuleAdapter{
		id: id,
	}

	if id.IsGlobal() {
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
	var err error

	log := klog.FromContext(ctx)
	log.V(2).Info("creating ComputeForwardingRule", "name", a.id)

	forwardingRule := proto.CloneOf(a.desired)
	forwardingRule.Name = direct.LazyPtr(a.id.ForwardingRule)
	desiredLabels := forwardingRule.Labels

	target := direct.ValueOf(forwardingRule.Target)
	isPrivateServiceConnect := target == "all-apis" || target == "vpc-sc" || strings.Contains(target, "/serviceAttachments/")

	// API restriction: Cannot set labels during creation(by POST). But it can be set later by PATCH SetLabels.
	// API error message: Labels are invalid in Private Service Connect Forwarding Rule.
	// See GH issue for details: https://github.com/hashicorp/terraform-provider-google/issues/16255
	if isPrivateServiceConnect {
		forwardingRule.Labels = nil
	}

	// Create forwarding rule(labels are not set during Insert)
	op := &gcp.Operation{}
	if a.id.IsGlobal() {
		req := &computepb.InsertGlobalForwardingRuleRequest{
			ForwardingRuleResource: forwardingRule,
			Project:                a.id.Project,
		}
		op, err = a.globalForwardingRulesClient.Insert(ctx, req)
	} else {
		req := &computepb.InsertForwardingRuleRequest{
			ForwardingRuleResource: forwardingRule,
			Region:                 a.id.Region,
			Project:                a.id.Project,
		}
		op, err = a.forwardingRulesClient.Insert(ctx, req)
	}
	if err != nil {
		return fmt.Errorf("creating ComputeForwardingRule %s: %w", a.id, err)
	}
	if !op.Done() {
		if err = op.Wait(ctx); err != nil {
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
	if isPrivateServiceConnect {
		forwardingRule.Labels = desiredLabels
	}
	if forwardingRule.Labels != nil {
		op, err := a.setLabels(ctx, created.LabelFingerprint, forwardingRule.Labels)
		if err != nil {
			return fmt.Errorf("adding ComputeForwardingRule labels %s: %w", a.id, err)
		}
		if !op.Done() {
			if err = op.Wait(ctx); err != nil {
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

	return a.updateStatus(ctx, createOp, created)
}

func (a *forwardingRuleAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	var err error

	if a.id.ForwardingRule == "" {
		return fmt.Errorf("resourceID is empty")
	}

	log := klog.FromContext(ctx)
	log.V(2).Info("updating ComputeForwardingRule", "name", a.id.ForwardingRule)

	diffs, err := compareForwardingRule(ctx, a.actual, a.desired, a.desiredStatusTarget, updateOp.GetUnstructured())
	if err != nil {
		return err
	}
	if !diffs.HasDiff() {
		log.V(2).Info("no diff detected for ComputeForwardingRule", "name", a.id)
		return a.updateStatus(ctx, updateOp, a.actual)
	}

	structuredreporting.ReportDiff(ctx, diffs)

	if hasFieldDiff(diffs, "allow_global_access") {
		desiredAllowGlobalAccess := direct.ValueOf(a.desired.AllowGlobalAccess)
		if a.id.Region != "global" {
			// To match the request body in TF-controller log
			// https://github.com/hashicorp/terraform-provider-google/blob/main/google/services/compute/resource_compute_forwarding_rule.go#L1151
			reqBody := &computepb.ForwardingRule{AllowGlobalAccess: &desiredAllowGlobalAccess}
			patchReq := &computepb.PatchForwardingRuleRequest{
				ForwardingRule:         a.id.ForwardingRule,
				ForwardingRuleResource: reqBody,
				Project:                a.id.Project,
				Region:                 a.id.Region,
			}
			op, err := a.forwardingRulesClient.Patch(ctx, patchReq)
			if err != nil {
				return fmt.Errorf("updating ComputeForwardingRule %s: %w", a.id, err)
			}
			if !op.Done() {
				if err = op.Wait(ctx); err != nil {
					return fmt.Errorf("waiting ComputeForwardingRule %s update failed: %w", a.id, err)
				}
			}
			log.V(2).Info("successfully updated ComputeForwardingRule", "name", a.id)
		}
	}

	// Use setTarget and setLabels to update target and labels fields.
	if hasFieldDiff(diffs, "labels") {
		labelsToSend := a.desired.Labels
		if labelsToSend == nil {
			labelsToSend = make(map[string]string)
		}
		op, err := a.setLabels(ctx, a.actual.LabelFingerprint, labelsToSend)
		if err != nil {
			return fmt.Errorf("updating ComputeForwardingRule labels %s: %w", a.id, err)
		}
		if !op.Done() {
			if err = op.Wait(ctx); err != nil {
				return fmt.Errorf("waiting ComputeForwardingRule %s update labels failed: %w", a.id, err)
			}
		}
		log.V(2).Info("successfully updated ComputeForwardingRule labels", "name", a.id)
	}

	// setTarget request is sent when there are updates to target.
	if hasFieldDiff(diffs, "target") {
		var op *gcp.Operation
		if a.id.IsGlobal() {
			setTargetReq := &computepb.SetTargetGlobalForwardingRuleRequest{
				ForwardingRule:          a.id.ForwardingRule,
				TargetReferenceResource: &computepb.TargetReference{Target: a.desired.Target},
				Project:                 a.id.Project,
			}
			op, err = a.globalForwardingRulesClient.SetTarget(ctx, setTargetReq)
		} else {
			setTargetReq := &computepb.SetTargetForwardingRuleRequest{
				ForwardingRule:          a.id.ForwardingRule,
				TargetReferenceResource: &computepb.TargetReference{Target: a.desired.Target},
				Project:                 a.id.Project,
				Region:                  a.id.Region,
			}
			op, err = a.forwardingRulesClient.SetTarget(ctx, setTargetReq)
		}
		if err != nil {
			return fmt.Errorf("updating ComputeForwardingRule target %s: %w", a.id, err)
		}
		if !op.Done() {
			if err = op.Wait(ctx); err != nil {
				return fmt.Errorf("waiting ComputeForwardingRule %s update target failed: %w", a.id, err)
			}
		}
		log.V(2).Info("successfully updated ComputeForwardingRule target", "name", a.id)
	}

	// Get the updated resource
	updated, err := a.get(ctx)
	if err != nil {
		return fmt.Errorf("getting ComputeForwardingRule %q: %w", a.id.ForwardingRule, err)
	}

	return a.updateStatus(ctx, updateOp, updated)
}

func compareForwardingRule(ctx context.Context, actual, desired *computepb.ForwardingRule, desiredStatusTarget *string, u *unstructured.Unstructured) (*structuredreporting.Diff, error) {
	diff := &structuredreporting.Diff{Object: u}

	desiredAllowGlobalAccess := false
	if desired.AllowGlobalAccess != nil {
		desiredAllowGlobalAccess = *desired.AllowGlobalAccess
	}
	actualAllowGlobalAccess := false
	if actual.AllowGlobalAccess != nil {
		actualAllowGlobalAccess = *actual.AllowGlobalAccess
	}
	if desiredAllowGlobalAccess != actualAllowGlobalAccess {
		diff.AddField("allow_global_access", actual.AllowGlobalAccess, desired.AllowGlobalAccess)
	}

	if !mapsEqual(desired.Labels, actual.Labels) {
		diff.AddField("labels", actual.Labels, desired.Labels)
	}

	targetMatchSpec := IsSelfLinkEqual(desired.Target, actual.Target)
	targetMatchStatus := IsSelfLinkEqual(desired.Target, desiredStatusTarget)
	if !targetMatchSpec || (desiredStatusTarget != nil && !targetMatchStatus) {
		diff.AddField("target", actual.Target, desired.Target)
	}

	return diff, nil
}

func hasFieldDiff(diff *structuredreporting.Diff, fieldID string) bool {
	for _, f := range diff.Fields {
		if f.ID == fieldID {
			return true
		}
	}
	return false
}

func (a *forwardingRuleAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *computepb.ForwardingRule) error {
	mapCtx := &direct.MapContext{}
	status := ComputeForwardingRuleStatus_v1beta1_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return op.UpdateStatus(ctx, status, nil)
}

func (a *forwardingRuleAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("forwardingrule %q not found", a.id)
	}

	obj := &krm.ComputeForwardingRule{}
	mc := &direct.MapContext{}
	obj.Spec = direct.ValueOf(ComputeForwardingRuleSpec_v1beta1_FromProto(mc, a.actual))
	if mc.Err() != nil {
		return nil, mc.Err()
	}

	// ComputeForwardingRule requires location and resourceID
	obj.Spec.Location = direct.LazyPtr(a.id.Region)
	obj.Spec.ResourceID = direct.LazyPtr(a.id.ForwardingRule)

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u := &unstructured.Unstructured{}
	u.Object = uObj
	u.SetName(a.id.ForwardingRule)
	u.SetGroupVersionKind(krm.ComputeForwardingRuleGVK)

	export.SetProjectID(u, a.id.Project)
	export.SetLabels(u, a.actual.Labels)

	return u, nil
}

// Delete implements the Adapter interface.
func (a *forwardingRuleAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	if a.id.ForwardingRule == "" {
		return false, fmt.Errorf("resourceID is empty")
	}

	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ComputeForwardingRule", "name", a.id.ForwardingRule)

	var err error
	op := &gcp.Operation{}
	if a.id.IsGlobal() {
		req := &computepb.DeleteGlobalForwardingRuleRequest{
			ForwardingRule: a.id.ForwardingRule,
			Project:        a.id.Project,
		}
		op, err = a.globalForwardingRulesClient.Delete(ctx, req)
	} else {
		req := &computepb.DeleteForwardingRuleRequest{
			ForwardingRule: a.id.ForwardingRule,
			Region:         a.id.Region,
			Project:        a.id.Project,
		}
		op, err = a.forwardingRulesClient.Delete(ctx, req)
	}
	if err != nil {
		return false, fmt.Errorf("deleting ComputeForwardingRule %s: %w", a.id.ForwardingRule, err)
	}
	if !op.Done() {
		if err := op.Wait(ctx); err != nil {
			return false, fmt.Errorf("waiting ComputeForwardingRule %s delete failed: %w", a.id.ForwardingRule, err)
		}
	}
	log.V(2).Info("successfully deleted ComputeForwardingRule", "name", a.id.ForwardingRule)
	return true, nil
}

func (a *forwardingRuleAdapter) get(ctx context.Context) (*computepb.ForwardingRule, error) {
	if a.id.IsGlobal() {
		getReq := &computepb.GetGlobalForwardingRuleRequest{
			ForwardingRule: a.id.ForwardingRule,
			Project:        a.id.Project,
		}
		return a.globalForwardingRulesClient.Get(ctx, getReq)
	} else {
		getReq := &computepb.GetForwardingRuleRequest{
			ForwardingRule: a.id.ForwardingRule,
			Region:         a.id.Region,
			Project:        a.id.Project,
		}
		return a.forwardingRulesClient.Get(ctx, getReq)
	}
}

func (a *forwardingRuleAdapter) setLabels(ctx context.Context, fingerprint *string, labels map[string]string) (*gcp.Operation, error) {
	op := &gcp.Operation{}
	var err error
	if a.id.IsGlobal() {
		setLabelsReq := &computepb.SetLabelsGlobalForwardingRuleRequest{
			Resource:                       a.id.ForwardingRule,
			GlobalSetLabelsRequestResource: &computepb.GlobalSetLabelsRequest{LabelFingerprint: fingerprint, Labels: labels},
			Project:                        a.id.Project,
		}
		op, err = a.globalForwardingRulesClient.SetLabels(ctx, setLabelsReq)
	} else {
		setLabelsReq := &computepb.SetLabelsForwardingRuleRequest{
			Resource:                       a.id.ForwardingRule,
			RegionSetLabelsRequestResource: &computepb.RegionSetLabelsRequest{LabelFingerprint: fingerprint, Labels: labels},
			Project:                        a.id.Project,
			Region:                         a.id.Region,
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

func mapsEqual(a, b map[string]string) bool {
	if len(a) == 0 && len(b) == 0 {
		return true
	}
	return reflect.DeepEqual(a, b)
}
