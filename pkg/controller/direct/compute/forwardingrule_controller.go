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

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"

	gcp "cloud.google.com/go/compute/apiv1"
	computepb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"google.golang.org/api/option"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
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
	resourceID                  string
	projectID                   string
	Location                    string
	forwardingRulesClient       *gcp.ForwardingRulesClient
	globalForwardingRulesClient *gcp.GlobalForwardingRulesClient
	desired                     *krm.ComputeForwardingRule
	actual                      *computepb.ForwardingRule
}

var _ directbase.Adapter = &forwardingRuleAdapter{}

func (m *forwardingRuleModel) globalForwardingRulesClient(ctx context.Context) (*gcp.GlobalForwardingRulesClient, error) {
	var opts []option.ClientOption
	if m.config.UserAgent != "" {
		opts = append(opts, option.WithUserAgent(m.config.UserAgent))
	}
	if m.config.HTTPClient != nil {
		opts = append(opts, option.WithHTTPClient(m.config.HTTPClient))
	}
	if m.config.UserProjectOverride && m.config.BillingProject != "" {
		opts = append(opts, option.WithQuotaProject(m.config.BillingProject))
	}

	globalForwardingRulesClient, err := gcp.NewGlobalForwardingRulesRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building global ComputeForwardingRule client: %w", err)
	}
	return globalForwardingRulesClient, err
}

func (m *forwardingRuleModel) forwardingRulesClient(ctx context.Context) (*gcp.ForwardingRulesClient, error) {
	var opts []option.ClientOption
	if m.config.UserAgent != "" {
		opts = append(opts, option.WithUserAgent(m.config.UserAgent))
	}
	if m.config.HTTPClient != nil {
		opts = append(opts, option.WithHTTPClient(m.config.HTTPClient))
	}
	if m.config.UserProjectOverride && m.config.BillingProject != "" {
		opts = append(opts, option.WithQuotaProject(m.config.BillingProject))
	}

	forwardingRulesClient, err := gcp.NewForwardingRulesRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building ComputeForwardingRule client: %w", err)
	}
	return forwardingRulesClient, err
}

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
		networkRef, err := refs.ResolveComputeNetwork(ctx, reader, obj, obj.Spec.NetworkRef)
		if err != nil {
			return nil, err

		}
		obj.Spec.NetworkRef.External = networkRef.ID()
	}

	// Get compute address
	if obj.Spec.IpAddress.AddressRef != nil {
		computeAddressRef, err := refs.ResolveComputeAddress(ctx, reader, obj, obj.Spec.IpAddress.AddressRef)
		if err != nil {
			return nil, err

		}
		obj.Spec.IpAddress.AddressRef.External = computeAddressRef.GetAddress()
	}

	// Get target targetHTTPProxy
	if obj.Spec.Target.TargetHTTPProxyRef != nil {
		targetHTTPProxyRef, err := refs.ResolveTargetHTTPProxy(ctx, reader, obj, obj.Spec.Target.TargetHTTPProxyRef)
		if err != nil {
			return nil, err

		}
		obj.Spec.Target.TargetHTTPProxyRef.External = targetHTTPProxyRef.Url()
	}

	// Get target TargetVPNGateway
	if obj.Spec.Target.TargetVPNGatewayRef != nil {
		targetVPNGatewayRef, err := refs.ResolveComputeTargetVPNGateway(ctx, reader, obj, obj.Spec.Target.TargetVPNGatewayRef)
		if err != nil {
			return nil, err

		}
		obj.Spec.Target.TargetVPNGatewayRef.External = targetVPNGatewayRef.URL()
	}

	// Get location
	location := obj.Spec.Location

	forwardingRuleAdapter := &forwardingRuleAdapter{
		resourceID: resourceID,
		projectID:  projectID,
		Location:   location,
		desired:    obj,
	}

	// Get compute GCP client
	if location == "global" {
		gcpClient, err := m.globalForwardingRulesClient(ctx)
		if err != nil {
			return nil, err
		}
		forwardingRuleAdapter.globalForwardingRulesClient = gcpClient
	} else {
		gcpClient, err := m.forwardingRulesClient(ctx)
		if err != nil {
			return nil, err
		}
		forwardingRuleAdapter.forwardingRulesClient = gcpClient
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
	} else {
		req := &computepb.GetForwardingRuleRequest{
			ForwardingRule: a.resourceID,
			Region:         a.Location,
			Project:        a.projectID,
		}
		forwardingRule, err = a.forwardingRulesClient.Get(ctx, req)
	}
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ComputeForwardingRule %q failed: %w", a.resourceID, err)
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

	desired := a.desired.DeepCopy()

	forwardingRule := ComputeForwardingRuleSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	forwardingRule.Name = direct.LazyPtr(a.resourceID)
	forwardingRule.Labels = desired.Labels
	// TODO(yuhou): handle default spec values
	if forwardingRule.LoadBalancingScheme == nil {
		forwardingRule.LoadBalancingScheme = direct.LazyPtr("EXTERNAL")
	}

	var err error
	op := &gcp.Operation{}
	created := &computepb.ForwardingRule{}
	if a.Location == "global" {
		req := &computepb.InsertGlobalForwardingRuleRequest{
			ForwardingRuleResource: forwardingRule,
			Project:                a.projectID,
		}
		getReq := &computepb.GetGlobalForwardingRuleRequest{
			ForwardingRule: a.resourceID,
			Project:        a.projectID,
		}
		op, err = a.globalForwardingRulesClient.Insert(ctx, req)
		if err != nil {
			return fmt.Errorf("creating ComputeForwardingRule %s failed: %w", a.resourceID, err)
		}
		log.V(2).Info("successfully created ComputeForwardingRule", "name", a.resourceID)
		// Get the created resource
		created, err = a.globalForwardingRulesClient.Get(ctx, getReq)
		if err != nil {
			if direct.IsNotFound(err) {
				return err
			}
			return fmt.Errorf("getting ComputeForwardingRule %q failed: %w", a.resourceID, err)
		}
	} else {
		req := &computepb.InsertForwardingRuleRequest{
			ForwardingRuleResource: forwardingRule,
			Region:                 a.Location,
			Project:                a.projectID,
		}
		getReq := &computepb.GetForwardingRuleRequest{
			ForwardingRule: a.resourceID,
			Region:         a.Location,
			Project:        a.projectID,
		}
		op, err = a.forwardingRulesClient.Insert(ctx, req)
		if err != nil {
			return fmt.Errorf("creating ComputeForwardingRule %s failed: %w", a.resourceID, err)
		}
		log.V(2).Info("successfully created ComputeForwardingRule", "name", a.resourceID)
		// Get the created resource
		created, err = a.forwardingRulesClient.Get(ctx, getReq)
		if err != nil {
			if direct.IsNotFound(err) {
				return err
			}
			return fmt.Errorf("getting ComputeForwardingRule %q failed: %w", a.resourceID, err)
		}
	}

	status := &krm.ComputeForwardingRuleStatus{
		CreationTimestamp: op.Proto().CreationTimestamp,
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
	// TODO(yuhou): Checked the realGCP logs, setLabels request is being sent even when there are no updates to labels.
	// That might because of the generated labelsFingerPrint?
	if a.Location == "global" {
		setLabelsReq := &computepb.SetLabelsGlobalForwardingRuleRequest{
			Resource:                       a.resourceID,
			GlobalSetLabelsRequestResource: &computepb.GlobalSetLabelsRequest{LabelFingerprint: a.actual.LabelFingerprint, Labels: forwardingRule.Labels},
			Project:                        a.projectID,
		}
		_, err = a.globalForwardingRulesClient.SetLabels(ctx, setLabelsReq)
	} else {
		setLabelsReq := &computepb.SetLabelsForwardingRuleRequest{
			Resource:                       a.resourceID,
			RegionSetLabelsRequestResource: &computepb.RegionSetLabelsRequest{LabelFingerprint: a.actual.LabelFingerprint, Labels: forwardingRule.Labels},
			Project:                        a.projectID,
			Region:                         a.Location,
		}
		_, err = a.forwardingRulesClient.SetLabels(ctx, setLabelsReq)
	}
	if err != nil {
		return fmt.Errorf("updating ComputeForwardingRule labels %s failed: %w", a.resourceID, err)
	}
	log.V(2).Info("successfully updated ComputeForwardingRule labels", "name", a.resourceID)

	// setTarget request is sent when there are updates to target.
	if !reflect.DeepEqual(forwardingRule.Target, a.actual.Target) {
		if a.Location == "global" {
			setTargetReq := &computepb.SetTargetGlobalForwardingRuleRequest{
				ForwardingRule:          a.resourceID,
				TargetReferenceResource: &computepb.TargetReference{Target: forwardingRule.Target},
				Project:                 a.projectID,
			}
			_, err = a.globalForwardingRulesClient.SetTarget(ctx, setTargetReq)
		} else {
			setTargetReq := &computepb.SetTargetForwardingRuleRequest{
				ForwardingRule:          a.resourceID,
				TargetReferenceResource: &computepb.TargetReference{Target: forwardingRule.Target},
				Project:                 a.projectID,
				Region:                  a.Location,
			}
			_, err = a.forwardingRulesClient.SetTarget(ctx, setTargetReq)
		}
		if err != nil {
			return fmt.Errorf("updating ComputeForwardingRule target %s failed: %w", a.resourceID, err)
		}
		log.V(2).Info("successfully updated ComputeForwardingRule target", "name", a.resourceID)
	}

	updated := &computepb.ForwardingRule{}
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
		if direct.IsNotFound(err) {
			return err
		}
		return fmt.Errorf("getting ComputeForwardingRule %q failed: %w", a.resourceID, err)
	}

	status := &krm.ComputeForwardingRuleStatus{}
	// TODO(yuhou): Do we need other status beside conditions and observedGeneration?
	status.SelfLink = updated.SelfLink
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

	if a.Location == "global" {
		req := &computepb.DeleteGlobalForwardingRuleRequest{
			ForwardingRule: a.resourceID,
			Project:        a.projectID,
		}
		_, err = a.globalForwardingRulesClient.Delete(ctx, req)
	} else {
		req := &computepb.DeleteForwardingRuleRequest{
			ForwardingRule: a.resourceID,
			Region:         a.Location,
			Project:        a.projectID,
		}
		_, err = a.forwardingRulesClient.Delete(ctx, req)
	}
	if err != nil {
		return false, fmt.Errorf("deleting ComputeForwardingRule %s failed: %w", a.resourceID, err)
	}
	log.V(2).Info("successfully deleted ComputeForwardingRule", "name", a.resourceID)
	return true, nil
}

func (a *forwardingRuleAdapter) globalFullyQualifiedName() string {
	return fmt.Sprintf("projects/%s/global/forwardingRules/%s", a.projectID, a.resourceID)
}

func (a *forwardingRuleAdapter) fullyQualifiedName() string {
	return fmt.Sprintf("projects/%s/regions/%s/forwardingRules/%s", a.projectID, a.Location, a.resourceID)
}

func (a *forwardingRuleAdapter) getParent() string {
	// TODO(user): Write the GCP URI parent for your resource
	return fmt.Sprintf("projects/%s", a.projectID)
}
