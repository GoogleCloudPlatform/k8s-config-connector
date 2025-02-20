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

package forwardingrule

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"k8s.io/klog/v2"

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"

	gcp "cloud.google.com/go/compute/apiv1"
	computepb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	kccpredicate "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/predicate"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const ctrlName = "forwardingrule-controller"

func init() {
	rg := &ForwardingRuleReconcileGate{}
	registry.RegisterModelWithReconcileGate(krm.ComputeForwardingRuleGVK, NewForwardingRuleModel, rg)
}

func NewForwardingRuleModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &forwardingRuleModel{config: config}, nil
}

type ForwardingRuleReconcileGate struct {
	optIn kccpredicate.OptInToDirectReconciliation
}

var _ kccpredicate.ReconcileGate = &ForwardingRuleReconcileGate{}

func (r *ForwardingRuleReconcileGate) ShouldReconcile(o *unstructured.Unstructured) bool {
	if r.optIn.ShouldReconcile(o) {
		return true
	}
	obj := &krm.ComputeForwardingRule{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(o.Object, &obj); err != nil {
		return false
	}
	// Run the direct reconciler only when spec.Target.GoogleAPIsBundle is specified
	return obj.Spec.Target != nil && obj.Spec.Target.GoogleAPIsBundle != nil
}

type forwardingRuleModel struct {
	config *config.ControllerConfig
}

// model implements the Model interface.
var _ directbase.Model = &forwardingRuleModel{}

type forwardingRuleAdapter struct {
	id                          *ForwardingRuleIdentity
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

	// Get ResourceID
	resourceID, err := refs.GetResourceID(u)
	if err != nil {
		return nil, err
	}

	// Get projectID
	projectID, err := refs.ResolveProjectID(ctx, reader, u)
	if err != nil {
		return nil, err
	}

	// Get location
	location := obj.Spec.Location

	// Set label managed-by-cnrm: true
	obj.ObjectMeta.Labels["managed-by-cnrm"] = "true"

	// Handle TF default values
	if obj.Spec.LoadBalancingScheme == nil {
		obj.Spec.LoadBalancingScheme = direct.LazyPtr("EXTERNAL")
	}

	// Validate ExternalRef
	var id *ForwardingRuleIdentity
	externalRef := direct.ValueOf(obj.Status.ExternalRef)
	if externalRef == "" {
		id = BuildID(projectID, location, resourceID)
	} else {
		id, err = asID(externalRef)
		if err != nil {
			return nil, err
		}

		if id.project != projectID {
			return nil, fmt.Errorf("ComputeForwardingRule %s/%s has spec.projectRef changed, expect %s, got %s",
				u.GetNamespace(), u.GetName(), id.project, projectID)
		}
		if id.location != location {
			return nil, fmt.Errorf("ComputeForwardingRule %s/%s has spec.location changed, expect %s, got %s",
				u.GetNamespace(), u.GetName(), id.location, location)
		}
		// TODO: need to support more cases
		if id.forwardingRule != resourceID {
			return nil, fmt.Errorf("ComputeForwardingRule  %s/%s has metadata.name or spec.resourceID changed, expect %s, got %s",
				u.GetNamespace(), u.GetName(), id.forwardingRule, resourceID)
		}
	}

	forwardingRuleAdapter := &forwardingRuleAdapter{
		id:      id,
		desired: obj,
		reader:  reader,
	}

	// Get GCP client
	gcpClient, err := newGCPClient(ctx, m.config)
	if err != nil {
		return nil, fmt.Errorf("building gcp client: %w", err)
	}
	if location == "global" {
		globalForwardingRulesClient, err := gcpClient.globalForwardingRuleClient(ctx)
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
	if a.id.forwardingRule == "" {
		return false, nil
	}

	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("getting ComputeForwardingRule", "name", a.id.forwardingRule)

	var err error
	forwardingRule := &computepb.ForwardingRule{}
	forwardingRule, err = a.get(ctx)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ComputeForwardingRule %q: %w", a.fullyQualifiedName(), err)
	}
	a.actual = forwardingRule
	return true, nil
}

func (a *forwardingRuleAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	u := createOp.GetUnstructured()
	var err error

	if a.id.project == "" {
		return fmt.Errorf("project is empty")
	}
	if a.id.forwardingRule == "" {
		return fmt.Errorf("resourceID is empty")
	}

	err = resolveDependencies(ctx, a.reader, a.desired)
	if err != nil {
		return err
	}

	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("creating ComputeForwardingRule", "name", a.id.forwardingRule)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()

	forwardingRule := ComputeForwardingRuleSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	forwardingRule.Name = direct.LazyPtr(a.id.forwardingRule)
	forwardingRule.Labels = desired.Labels

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
	if a.id.location == "global" {
		req := &computepb.InsertGlobalForwardingRuleRequest{
			ForwardingRuleResource: forwardingRule,
			Project:                a.id.project,
		}
		op, err = a.globalForwardingRulesClient.Insert(ctx, req)
	} else {
		req := &computepb.InsertForwardingRuleRequest{
			ForwardingRuleResource: forwardingRule,
			Region:                 a.id.location,
			Project:                a.id.project,
		}
		op, err = a.forwardingRulesClient.Insert(ctx, req)
	}
	if err != nil {
		return fmt.Errorf("creating ComputeForwardingRule %s: %w", a.fullyQualifiedName(), err)
	}
	if !op.Done() {
		err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("waiting ComputeForwardingRule %s create failed: %w", a.fullyQualifiedName(), err)
		}
	}
	log.V(2).Info("successfully created ComputeForwardingRule", "name", a.fullyQualifiedName())

	// Get the created resource
	created := &computepb.ForwardingRule{}
	created, err = a.get(ctx)
	if err != nil {
		return fmt.Errorf("getting ComputeForwardingRule %q: %w", a.fullyQualifiedName(), err)
	}

	// Set labels for the created forwarding rule
	// Add labels back for psc forwarding rule
	if target == "all-apis" || target == "vpc-sc" || strings.Contains(target, "/serviceAttachments/") {
		forwardingRule.Labels = desired.Labels
	}
	if forwardingRule.Labels != nil {
		op, err := a.setLabels(ctx, created.LabelFingerprint, forwardingRule.Labels)
		if err != nil {
			return fmt.Errorf("adding ComputeForwardingRule labels %s: %w", a.fullyQualifiedName(), err)
		}
		if !op.Done() {
			err = op.Wait(ctx)
			if err != nil {
				return fmt.Errorf("waiting ComputeForwardingRule %s add labels failed: %w", a.fullyQualifiedName(), err)
			}
		}
		log.V(2).Info("successfully added ComputeForwardingRule labels", "name", a.fullyQualifiedName())

		// Get the created resource with label added
		created, err = a.get(ctx)
		if err != nil {
			return fmt.Errorf("getting ComputeForwardingRule %q: %w", a.fullyQualifiedName(), err)
		}
	}

	status := &krm.ComputeForwardingRuleStatus{
		LabelFingerprint:  created.LabelFingerprint,
		CreationTimestamp: created.CreationTimestamp,
		SelfLink:          created.SelfLink,
	}
	status.ExternalRef = a.id.AsExternalRef()
	return setStatus(u, status)
}

func (a *forwardingRuleAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	u := updateOp.GetUnstructured()
	var err error

	if a.id.forwardingRule == "" {
		return fmt.Errorf("resourceID is empty")
	}

	err = resolveDependencies(ctx, a.reader, a.desired)
	if err != nil {
		return err
	}

	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("updating ComputeForwardingRule", "name", a.id.forwardingRule)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	forwardingRule := ComputeForwardingRuleSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	forwardingRule.Name = direct.LazyPtr(a.id.forwardingRule)
	forwardingRule.Labels = desired.Labels

	op := &gcp.Operation{}
	updated := &computepb.ForwardingRule{}
	if !reflect.DeepEqual(forwardingRule.AllowGlobalAccess, a.actual.AllowGlobalAccess) {
		// To match the request body in TF-controller log
		// https://github.com/hashicorp/terraform-provider-google/blob/main/google/services/compute/resource_compute_forwarding_rule.go#L1151
		reqBody := &computepb.ForwardingRule{AllowGlobalAccess: forwardingRule.AllowGlobalAccess}
		if a.id.location == "global" {
			// TF does not support allowGlobalAccess field for global forwarding rule
			// Underlying API as well, error message: `Field allow-global-access is only supported for regional INTERNAL
			// forwarding rules with backend service/target instance or regional INTERNAL_MANAGED forwarding rules.`
			forwardingRule.AllowGlobalAccess = nil
		} else {
			patchReq := &computepb.PatchForwardingRuleRequest{
				ForwardingRule:         a.id.forwardingRule,
				ForwardingRuleResource: reqBody,
				Project:                a.id.project,
				Region:                 a.id.location,
			}
			op, err = a.forwardingRulesClient.Patch(ctx, patchReq)
		}
		if err != nil {
			return fmt.Errorf("updating ComputeForwardingRule %s: %w", a.fullyQualifiedName(), err)
		}
		err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("waiting ComputeForwardingRule %s update failed: %w", a.fullyQualifiedName(), err)
		}
		log.V(2).Info("successfully updated ComputeForwardingRule", "name", a.fullyQualifiedName())
	}

	// Use setTarget and setLabels to update target and labels fields.
	if !reflect.DeepEqual(forwardingRule.Labels, a.actual.Labels) {
		op, err := a.setLabels(ctx, a.actual.LabelFingerprint, forwardingRule.Labels)
		if err != nil {
			return fmt.Errorf("updating ComputeForwardingRule labels %s: %w", a.fullyQualifiedName(), err)
		}
		if !op.Done() {
			err = op.Wait(ctx)
			if err != nil {
				return fmt.Errorf("waiting ComputeForwardingRule %s update labels failed: %w", a.fullyQualifiedName(), err)
			}
		}
		log.V(2).Info("successfully updated ComputeForwardingRule labels", "name", a.fullyQualifiedName())
	}

	// setTarget request is sent when there are updates to target.
	if !reflect.DeepEqual(forwardingRule.Target, a.actual.Target) {
		if a.id.location == "global" {
			setTargetReq := &computepb.SetTargetGlobalForwardingRuleRequest{
				ForwardingRule:          a.id.forwardingRule,
				TargetReferenceResource: &computepb.TargetReference{Target: forwardingRule.Target},
				Project:                 a.id.project,
			}
			op, err = a.globalForwardingRulesClient.SetTarget(ctx, setTargetReq)
		} else {
			setTargetReq := &computepb.SetTargetForwardingRuleRequest{
				ForwardingRule:          a.id.forwardingRule,
				TargetReferenceResource: &computepb.TargetReference{Target: forwardingRule.Target},
				Project:                 a.id.project,
				Region:                  a.id.location,
			}
			op, err = a.forwardingRulesClient.SetTarget(ctx, setTargetReq)
		}
		if err != nil {
			return fmt.Errorf("updating ComputeForwardingRule target %s: %w", a.fullyQualifiedName(), err)
		}
		if !op.Done() {
			err = op.Wait(ctx)
			if err != nil {
				return fmt.Errorf("waiting ComputeForwardingRule %s update target failed: %w", a.fullyQualifiedName(), err)
			}
		}
		log.V(2).Info("successfully updated ComputeForwardingRule target", "name", a.fullyQualifiedName())
	}
	// Get the updated resource
	updated, err = a.get(ctx)
	if err != nil {
		return fmt.Errorf("getting ComputeForwardingRule %q: %w", a.id.forwardingRule, err)
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
		return nil, fmt.Errorf("forwardingrule %q not found", a.fullyQualifiedName())
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
	u.SetName(a.id.forwardingRule)
	u.SetGroupVersionKind(krm.ComputeForwardingRuleGVK)
	u.SetLabels(a.actual.Labels)

	if err := unstructured.SetNestedField(u.Object, specObj, "spec"); err != nil {
		return nil, fmt.Errorf("setting spec: %w", err)
	}

	return u, nil
}

// Delete implements the Adapter interface.
func (a *forwardingRuleAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	if a.id.forwardingRule == "" {
		return false, fmt.Errorf("resourceID is empty")
	}

	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("deleting ComputeForwardingRule", "name", a.id.forwardingRule)

	var err error
	op := &gcp.Operation{}
	if a.id.location == "global" {
		req := &computepb.DeleteGlobalForwardingRuleRequest{
			ForwardingRule: a.id.forwardingRule,
			Project:        a.id.project,
		}
		op, err = a.globalForwardingRulesClient.Delete(ctx, req)
	} else {
		req := &computepb.DeleteForwardingRuleRequest{
			ForwardingRule: a.id.forwardingRule,
			Region:         a.id.location,
			Project:        a.id.project,
		}
		op, err = a.forwardingRulesClient.Delete(ctx, req)
	}
	if err != nil {
		return false, fmt.Errorf("deleting ComputeForwardingRule %s: %w", a.id.forwardingRule, err)
	}
	if !op.Done() {
		err = op.Wait(ctx)
		if err != nil {
			return false, fmt.Errorf("waiting ComputeForwardingRule %s delete failed: %w", a.id.forwardingRule, err)
		}
	}
	log.V(2).Info("successfully deleted ComputeForwardingRule", "name", a.id.forwardingRule)
	return true, nil
}

func (a *forwardingRuleAdapter) get(ctx context.Context) (*computepb.ForwardingRule, error) {
	if a.id.location == "global" {
		getReq := &computepb.GetGlobalForwardingRuleRequest{
			ForwardingRule: a.id.forwardingRule,
			Project:        a.id.project,
		}
		return a.globalForwardingRulesClient.Get(ctx, getReq)
	} else {
		getReq := &computepb.GetForwardingRuleRequest{
			ForwardingRule: a.id.forwardingRule,
			Region:         a.id.location,
			Project:        a.id.project,
		}
		return a.forwardingRulesClient.Get(ctx, getReq)
	}
}

func (a *forwardingRuleAdapter) setLabels(ctx context.Context, fingerprint *string, labels map[string]string) (*gcp.Operation, error) {
	op := &gcp.Operation{}
	var err error
	if a.id.location == "global" {
		setLabelsReq := &computepb.SetLabelsGlobalForwardingRuleRequest{
			Resource:                       a.id.forwardingRule,
			GlobalSetLabelsRequestResource: &computepb.GlobalSetLabelsRequest{LabelFingerprint: fingerprint, Labels: labels},
			Project:                        a.id.project,
		}
		op, err = a.globalForwardingRulesClient.SetLabels(ctx, setLabelsReq)
	} else {
		setLabelsReq := &computepb.SetLabelsForwardingRuleRequest{
			Resource:                       a.id.forwardingRule,
			RegionSetLabelsRequestResource: &computepb.RegionSetLabelsRequest{LabelFingerprint: fingerprint, Labels: labels},
			Project:                        a.id.project,
			Region:                         a.id.location,
		}
		op, err = a.forwardingRulesClient.SetLabels(ctx, setLabelsReq)
	}
	return op, err
}

func (a *forwardingRuleAdapter) fullyQualifiedName() string {
	if a.id.location == "global" {
		return fmt.Sprintf("projects/%s/global/forwardingRules/%s", a.id.project, a.id.forwardingRule)
	}
	return fmt.Sprintf("projects/%s/regions/%s/forwardingRules/%s", a.id.project, a.id.location, a.id.forwardingRule)

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

func resolveDependencies(ctx context.Context, reader client.Reader, obj *krm.ComputeForwardingRule) error {
	// Get network
	if obj.Spec.NetworkRef != nil {
		networkRef, err := ResolveComputeNetwork(ctx, reader, obj, obj.Spec.NetworkRef)
		if err != nil {
			return err

		}
		obj.Spec.NetworkRef.External = networkRef.External
	}

	// Get subnetwork
	if obj.Spec.SubnetworkRef != nil {
		subnetworkRef, err := ResolveComputeSubnetwork(ctx, reader, obj, obj.Spec.SubnetworkRef)
		if err != nil {
			return err

		}
		obj.Spec.SubnetworkRef.External = subnetworkRef.External
	}

	// Get backend service
	if obj.Spec.BackendServiceRef != nil {
		normalizedExternal, err := obj.Spec.BackendServiceRef.NormalizedExternal(ctx, reader, obj.GetNamespace())
		if err != nil {
			return err

		}
		obj.Spec.BackendServiceRef.External = normalizedExternal
	}

	// Get ip address, ip address is optional
	if obj.Spec.IpAddress != nil && obj.Spec.IpAddress.AddressRef != nil {
		computeAddressRef, err := ResolveComputeAddress(ctx, reader, obj, obj.Spec.IpAddress.AddressRef)
		if err != nil {
			return err

		}
		obj.Spec.IpAddress.AddressRef.External = computeAddressRef.External
	}

	// Get target, target is optional
	if obj.Spec.Target != nil {
		// Get target ServiceAttachment
		if obj.Spec.Target.ServiceAttachmentRef != nil {
			serviceAttachmentRef, err := ResolveComputeServiceAttachment(ctx, reader, obj, obj.Spec.Target.ServiceAttachmentRef)
			if err != nil {
				return err

			}
			obj.Spec.Target.ServiceAttachmentRef.External = serviceAttachmentRef.External
		}

		// Get target ComputeTargetGRPCProxyRef
		if obj.Spec.Target.TargetGRPCProxyRef != nil {
			targetGRPCProxyRef, err := ResolveComputeTargetGrpcProxy(ctx, reader, obj, obj.Spec.Target.TargetGRPCProxyRef)
			if err != nil {
				return err

			}
			obj.Spec.Target.TargetGRPCProxyRef.External = targetGRPCProxyRef.External
		}

		// Get target ComputeTargetHTTPProxy
		if obj.Spec.Target.TargetHTTPProxyRef != nil {
			targetHTTPProxyRef, err := ResolveComputeTargetHTTPProxy(ctx, reader, obj, obj.Spec.Target.TargetHTTPProxyRef)
			if err != nil {
				return err

			}
			obj.Spec.Target.TargetHTTPProxyRef.External = targetHTTPProxyRef.External
		}

		// Get target ComputeTargetHTTPSProxy
		if obj.Spec.Target.TargetHTTPSProxyRef != nil {
			targetHTTPSProxyRef, err := ResolveComputeTargetHTTPSProxy(ctx, reader, obj, obj.Spec.Target.TargetHTTPSProxyRef)
			if err != nil {
				return err

			}
			obj.Spec.Target.TargetHTTPSProxyRef.External = targetHTTPSProxyRef.External
		}

		// Get target TargetVPNGateway
		if obj.Spec.Target.TargetVPNGatewayRef != nil {
			targetVPNGatewayRef, err := ResolveComputeTargetVPNGateway(ctx, reader, obj, obj.Spec.Target.TargetVPNGatewayRef)
			if err != nil {
				return err

			}
			obj.Spec.Target.TargetVPNGatewayRef.External = targetVPNGatewayRef.External
		}

		// Get target SSLProxy
		if obj.Spec.Target.TargetSSLProxyRef != nil {
			targetSSLProxyRef, err := ResolveComputeTargetSSLProxy(ctx, reader, obj, obj.Spec.Target.TargetSSLProxyRef)
			if err != nil {
				return err

			}
			obj.Spec.Target.TargetSSLProxyRef.External = targetSSLProxyRef.External
		}

		// Get target TCPProxy
		if obj.Spec.Target.TargetTCPProxyRef != nil {
			targetTCPProxyRef, err := ResolveComputeTargetTCPProxy(ctx, reader, obj, obj.Spec.Target.TargetTCPProxyRef)
			if err != nil {
				return err

			}
			obj.Spec.Target.TargetTCPProxyRef.External = targetTCPProxyRef.External
		}
	}
	return nil
}
