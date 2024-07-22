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

	"k8s.io/klog/v2"

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"

	gcp "cloud.google.com/go/compute/apiv1"
	computepb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const ctrlName = "forwardingrule-controller"

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
	resourceID                  string
	projectID                   string
	Location                    string
	forwardingRulesClient       *gcp.ForwardingRulesClient
	globalForwardingRulesClient *gcp.GlobalForwardingRulesClient
	desired                     *krm.ComputeForwardingRule
	actual                      *computepb.ForwardingRule
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

	// Get network
	if obj.Spec.NetworkRef != nil {
		networkRef, err := ResolveComputeNetwork(ctx, reader, obj, obj.Spec.NetworkRef)
		if err != nil {
			return nil, err

		}
		obj.Spec.NetworkRef.External = networkRef.External
	}

	// Get compute address
	if obj.Spec.IpAddress.AddressRef != nil {
		computeAddressRef, err := ResolveComputeAddress(ctx, reader, obj, obj.Spec.IpAddress.AddressRef)
		if err != nil {
			return nil, err

		}
		obj.Spec.IpAddress.AddressRef.External = computeAddressRef.External
	}

	// Get target ComputeTargetHTTPProxy
	if obj.Spec.Target.TargetHTTPProxyRef != nil {
		targetHTTPProxyRef, err := ResolveComputeTargetHTTPProxy(ctx, reader, obj, obj.Spec.Target.TargetHTTPProxyRef)
		if err != nil {
			return nil, err

		}
		obj.Spec.Target.TargetHTTPProxyRef.External = targetHTTPProxyRef.External
	}

	// Get target TargetVPNGateway
	if obj.Spec.Target.TargetVPNGatewayRef != nil {
		targetVPNGatewayRef, err := ResolveComputeTargetVPNGateway(ctx, reader, obj, obj.Spec.Target.TargetVPNGatewayRef)
		if err != nil {
			return nil, err

		}
		obj.Spec.Target.TargetVPNGatewayRef.External = targetVPNGatewayRef.External
	}

	// Get location
	location := obj.Spec.Location

	// Set label managed-by-cnrm: true
	obj.ObjectMeta.Labels["managed-by-cnrm"] = "true"

	// Handle TF default values
	if obj.Spec.LoadBalancingScheme == nil {
		obj.Spec.LoadBalancingScheme = direct.LazyPtr("EXTERNAL")
	}

	forwardingRuleAdapter := &forwardingRuleAdapter{
		resourceID: resourceID,
		projectID:  projectID,
		Location:   location,
		desired:    obj,
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
	if a.resourceID == "" {
		return false, nil
	}

	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("getting ComputeForwardingRule", "name", a.resourceID)

	var err error
	forwardingRule := &computepb.ForwardingRule{}
	if a.Location == "global" {
		req := &computepb.GetGlobalForwardingRuleRequest{
			ForwardingRule: a.resourceID,
			Project:        a.projectID,
		}
		forwardingRule, err = a.globalForwardingRulesClient.Get(ctx, req)
		if err != nil {
			if direct.IsNotFound(err) {
				return false, nil
			}
			return false, fmt.Errorf("getting ComputeForwardingRule %q failed: %w", a.fullyQualifiedName(), err)
		}
	} else {
		req := &computepb.GetForwardingRuleRequest{
			ForwardingRule: a.resourceID,
			Region:         a.Location,
			Project:        a.projectID,
		}
		forwardingRule, err = a.forwardingRulesClient.Get(ctx, req)
		if err != nil {
			if direct.IsNotFound(err) {
				return false, nil
			}
			return false, fmt.Errorf("getting ComputeForwardingRule %q failed: %w", a.fullyQualifiedName(), err)
		}
	}
	a.actual = forwardingRule
	return true, nil
}

func (a *forwardingRuleAdapter) Create(ctx context.Context, u *unstructured.Unstructured) error {
	if a.projectID == "" {
		return fmt.Errorf("project is empty")
	}
	if a.resourceID == "" {
		return fmt.Errorf("resourceID is empty")
	}

	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("creating ComputeForwardingRule", "name", a.resourceID)
	mapCtx := &direct.MapContext{}

	a.desired.Labels["managed-by-cnrm"] = "true"
	desired := a.desired.DeepCopy()

	forwardingRule := ComputeForwardingRuleSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	forwardingRule.Name = direct.LazyPtr(a.resourceID)
	forwardingRule.Labels = desired.Labels

	var err error
	op := &gcp.Operation{}
	if a.Location == "global" {
		req := &computepb.InsertGlobalForwardingRuleRequest{
			ForwardingRuleResource: forwardingRule,
			Project:                a.projectID,
		}
		op, err = a.globalForwardingRulesClient.Insert(ctx, req)
	} else {
		req := &computepb.InsertForwardingRuleRequest{
			ForwardingRuleResource: forwardingRule,
			Region:                 a.Location,
			Project:                a.projectID,
		}
		op, err = a.forwardingRulesClient.Insert(ctx, req)
	}
	if err != nil {
		return fmt.Errorf("creating ComputeForwardingRule %s failed: %w", a.fullyQualifiedName(), err)
	}
	log.V(2).Info("successfully created ComputeForwardingRule", "name", a.fullyQualifiedName())
	// Get the created resource
	created := &computepb.ForwardingRule{}
	if a.Location == "global" {
		getReq := &computepb.GetGlobalForwardingRuleRequest{
			ForwardingRule: a.resourceID,
			Project:        a.projectID,
		}
		created, err = a.globalForwardingRulesClient.Get(ctx, getReq)
	} else {
		getReq := &computepb.GetForwardingRuleRequest{
			ForwardingRule: a.resourceID,
			Region:         a.Location,
			Project:        a.projectID,
		}
		created, err = a.forwardingRulesClient.Get(ctx, getReq)
	}
	if err != nil {
		return fmt.Errorf("getting ComputeForwardingRule %q failed: %w", a.fullyQualifiedName(), err)
	}

	status := &krm.ComputeForwardingRuleStatus{
		LabelFingerprint:  created.LabelFingerprint,
		CreationTimestamp: op.Proto().InsertTime,
		SelfLink:          created.SelfLink,
	}
	return setStatus(u, status)
}

func (a *forwardingRuleAdapter) Update(ctx context.Context, u *unstructured.Unstructured) error {
	if a.resourceID == "" {
		return fmt.Errorf("resourceID is empty")
	}

	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("updating ComputeForwardingRule", "name", a.resourceID)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	forwardingRule := ComputeForwardingRuleSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	forwardingRule.Name = direct.LazyPtr(a.resourceID)
	forwardingRule.Labels = desired.Labels

	// Patch only support update on networkTier field, which KCC does not support yet.
	// Use setTarget and setLabels to update target and labels fields.
	var err error
	op := &gcp.Operation{}
	updated := &computepb.ForwardingRule{}
	// TODO(yuhou): Checked the realGCP logs, setLabels request is being sent even when there are no updates to labels.
	// That might because of the generated labelsFingerPrint?
	if a.Location == "global" {
		setLabelsReq := &computepb.SetLabelsGlobalForwardingRuleRequest{
			Resource:                       a.resourceID,
			GlobalSetLabelsRequestResource: &computepb.GlobalSetLabelsRequest{LabelFingerprint: a.actual.LabelFingerprint, Labels: forwardingRule.Labels},
			Project:                        a.projectID,
		}
		op, err = a.globalForwardingRulesClient.SetLabels(ctx, setLabelsReq)
	} else {
		setLabelsReq := &computepb.SetLabelsForwardingRuleRequest{
			Resource:                       a.resourceID,
			RegionSetLabelsRequestResource: &computepb.RegionSetLabelsRequest{LabelFingerprint: a.actual.LabelFingerprint, Labels: forwardingRule.Labels},
			Project:                        a.projectID,
			Region:                         a.Location,
		}
		op, err = a.forwardingRulesClient.SetLabels(ctx, setLabelsReq)
	}
	if err != nil {
		return fmt.Errorf("updating ComputeForwardingRule labels %s failed: %w", a.fullyQualifiedName(), err)
	}
	err = op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting ComputeForwardingRule %s update labels failed: %w", a.fullyQualifiedName(), err)
	}
	log.V(2).Info("successfully updated ComputeForwardingRule labels", "name", a.fullyQualifiedName())
	// Get the updated resource
	if a.Location == "global" {
		getReq := &computepb.GetGlobalForwardingRuleRequest{
			ForwardingRule: a.resourceID,
			Project:        a.projectID,
		}
		updated, err = a.globalForwardingRulesClient.Get(ctx, getReq)
	} else {
		getReq := &computepb.GetForwardingRuleRequest{
			ForwardingRule: a.resourceID,
			Region:         a.Location,
			Project:        a.projectID,
		}
		updated, err = a.forwardingRulesClient.Get(ctx, getReq)
	}
	if err != nil {
		return fmt.Errorf("getting ComputeForwardingRule %q failed: %w", a.resourceID, err)
	}

	// setTarget request is sent when there are updates to target.
	if !reflect.DeepEqual(forwardingRule.Target, a.actual.Target) {
		if a.Location == "global" {
			setTargetReq := &computepb.SetTargetGlobalForwardingRuleRequest{
				ForwardingRule:          a.resourceID,
				TargetReferenceResource: &computepb.TargetReference{Target: forwardingRule.Target},
				Project:                 a.projectID,
			}
			op, err = a.globalForwardingRulesClient.SetTarget(ctx, setTargetReq)
		} else {
			setTargetReq := &computepb.SetTargetForwardingRuleRequest{
				ForwardingRule:          a.resourceID,
				TargetReferenceResource: &computepb.TargetReference{Target: forwardingRule.Target},
				Project:                 a.projectID,
				Region:                  a.Location,
			}
			op, err = a.forwardingRulesClient.SetTarget(ctx, setTargetReq)
		}
		if err != nil {
			return fmt.Errorf("updating ComputeForwardingRule target %s failed: %w", a.fullyQualifiedName(), err)
		}
		err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("waiting ComputeForwardingRule %s update target failed: %w", a.fullyQualifiedName(), err)
		}
		log.V(2).Info("successfully updated ComputeForwardingRule target", "name", a.fullyQualifiedName())
	}

	// Get the updated resource
	if a.Location == "global" {
		getReq := &computepb.GetGlobalForwardingRuleRequest{
			ForwardingRule: a.resourceID,
			Project:        a.projectID,
		}
		updated, err = a.globalForwardingRulesClient.Get(ctx, getReq)
	} else {
		getReq := &computepb.GetForwardingRuleRequest{
			ForwardingRule: a.resourceID,
			Region:         a.Location,
			Project:        a.projectID,
		}
		updated, err = a.forwardingRulesClient.Get(ctx, getReq)
	}
	if err != nil {
		return fmt.Errorf("getting ComputeForwardingRule %q failed: %w", a.resourceID, err)
	}

	status := &krm.ComputeForwardingRuleStatus{
		LabelFingerprint:  updated.LabelFingerprint,
		CreationTimestamp: op.Proto().InsertTime,
		SelfLink:          updated.SelfLink,
	}
	return setStatus(u, status)
}

func (a *forwardingRuleAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	// TODO(kcc)
	return nil, nil
}

// Delete implements the Adapter interface.
func (a *forwardingRuleAdapter) Delete(ctx context.Context) (bool, error) {
	if a.resourceID == "" {
		return false, fmt.Errorf("resourceID is empty")
	}

	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("deleting ComputeForwardingRule", "name", a.resourceID)

	exist, err := a.Find(ctx)
	if err != nil {
		return false, err
	}
	if !exist {
		// return (false, nil) if the object was not found but should be presumed deleted.
		return false, nil
	}

	op := &gcp.Operation{}
	if a.Location == "global" {
		req := &computepb.DeleteGlobalForwardingRuleRequest{
			ForwardingRule: a.resourceID,
			Project:        a.projectID,
		}
		op, err = a.globalForwardingRulesClient.Delete(ctx, req)
	} else {
		req := &computepb.DeleteForwardingRuleRequest{
			ForwardingRule: a.resourceID,
			Region:         a.Location,
			Project:        a.projectID,
		}
		op, err = a.forwardingRulesClient.Delete(ctx, req)
	}
	if err != nil {
		return false, fmt.Errorf("deleting ComputeForwardingRule %s failed: %w", a.resourceID, err)
	}
	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting ComputeForwardingRule %s delete failed: %w", a.resourceID, err)
	}
	log.V(2).Info("successfully deleted ComputeForwardingRule", "name", a.resourceID)
	return true, nil
}

func (a *forwardingRuleAdapter) fullyQualifiedName() string {
	if a.Location == "global" {
		return fmt.Sprintf("projects/%s/global/forwardingRules/%s", a.projectID, a.resourceID)
	}
	return fmt.Sprintf("projects/%s/regions/%s/forwardingRules/%s", a.projectID, a.Location, a.resourceID)
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
	}

	u.Object["status"] = status

	return nil
}
