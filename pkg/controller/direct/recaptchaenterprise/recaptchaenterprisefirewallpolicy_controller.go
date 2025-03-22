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

// +tool:controller
// proto.service: google.cloud.recaptchaenterprise.v1.RecaptchaEnterpriseService
// proto.message: google.cloud.recaptchaenterprise.v1.FirewallPolicy
// crd.type: ReCAPTCHAEnterpriseFirewallPolicy
// crd.version: v1alpha1

package recaptchaenterprise

import (
	"context"
	"fmt"
	"reflect"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/recaptchaenterprise/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util/slice"

	gcp "cloud.google.com/go/recaptchaenterprise/v2/apiv1"
	pb "cloud.google.com/go/recaptchaenterprise/v2/apiv1/recaptchaenterprisepb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.ReCAPTCHAEnterpriseFirewallPolicyGVK, NewFirewallPolicyModel)
}

func NewFirewallPolicyModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelFirewallPolicy{config: *config}, nil
}

var _ directbase.Model = &modelFirewallPolicy{}

type modelFirewallPolicy struct {
	config config.ControllerConfig
}

func (m *modelFirewallPolicy) client(ctx context.Context, projectID string) (*gcp.Client, error) {
	opts := []option.ClientOption{}

	// Workaround for an unusual behaviour (bug?):
	//  the service requires that a quota project be set
	if !m.config.UserProjectOverride || m.config.BillingProject == "" {
		m.config.UserProjectOverride = true
		m.config.BillingProject = projectID
	}

	if optsForClient, err := m.config.RESTClientOptions(); err != nil {
		return nil, err
	} else {
		opts = append(opts, optsForClient...)
	}

	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building recaptchaenterprise client: %w", err)
	}

	return gcpClient, err
}

func (m *modelFirewallPolicy) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.ReCAPTCHAEnterpriseFirewallPolicy{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewFirewallPolicyIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	gcpClient, err := m.client(ctx, id.ProjectID)
	if err != nil {
		return nil, err
	}
	return &FirewallPolicyAdapter{
		gcpClient: gcpClient,
		id:        id,
		desired:   obj,
		reader:    reader,
	}, nil
}

func (m *modelFirewallPolicy) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type FirewallPolicyAdapter struct {
	gcpClient *gcp.Client
	id        *krm.FirewallPolicyIdentity
	desired   *krm.ReCAPTCHAEnterpriseFirewallPolicy
	actual    *pb.FirewallPolicy
	reader    client.Reader
}

var _ directbase.Adapter = &FirewallPolicyAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *FirewallPolicyAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting FirewallPolicy", "name", a.id)

	req := &pb.GetFirewallPolicyRequest{Name: a.id.String()}
	firewallpolicypb, err := a.gcpClient.GetFirewallPolicy(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting FirewallPolicy %q: %w", a.id, err)
	}

	a.actual = firewallpolicypb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *FirewallPolicyAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating FirewallPolicy", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := ReCAPTCHAEnterpriseFirewallPolicySpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	// resource.Name = a.id.String() // NOTE: required for Update, but forbidden for Create

	req := &pb.CreateFirewallPolicyRequest{
		Parent:         a.id.Parent().String(),
		FirewallPolicy: resource,
	}
	created, err := a.gcpClient.CreateFirewallPolicy(ctx, req)
	if err != nil {
		return fmt.Errorf("creating FirewallPolicy %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created FirewallPolicy", "name", a.id)

	status := &krm.ReCAPTCHAEnterpriseFirewallPolicyStatus{}
	status.ObservedState = ReCAPTCHAEnterpriseFirewallPolicyObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *FirewallPolicyAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating FirewallPolicy", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := ReCAPTCHAEnterpriseFirewallPolicySpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	paths := []string{}
	if desired.Spec.Description != nil && !reflect.DeepEqual(resource.Description, a.actual.Description) {
		paths = append(paths, "description")
	}
	if desired.Spec.Path != nil && !reflect.DeepEqual(resource.Path, a.actual.Path) {
		paths = append(paths, "path")
	}
	if desired.Spec.Condition != nil && !reflect.DeepEqual(resource.Condition, a.actual.Condition) {
		paths = append(paths, "condition")
	}
	if desired.Spec.Actions != nil && !slice.ContentEqualsWithOrder(resource.Actions, a.actual.Actions, mapCtx.CompareActions) {
		paths = append(paths, "actions")
	}

	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		return nil
	}

	resource.Name = a.id.String() // we need to set the name so that GCP API can identify the resource
	req := &pb.UpdateFirewallPolicyRequest{
		FirewallPolicy: resource,
		UpdateMask:     &fieldmaskpb.FieldMask{Paths: paths},
	}
	updated, err := a.gcpClient.UpdateFirewallPolicy(ctx, req)
	if err != nil {
		return fmt.Errorf("updating FirewallPolicy %s: %w", a.id, err)
	}
	log.V(2).Info("successfully updated FirewallPolicy", "name", a.id)

	status := &krm.ReCAPTCHAEnterpriseFirewallPolicyStatus{}
	status.ObservedState = ReCAPTCHAEnterpriseFirewallPolicyObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *FirewallPolicyAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ReCAPTCHAEnterpriseFirewallPolicy{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(ReCAPTCHAEnterpriseFirewallPolicySpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = direct.LazyPtr(a.id.Parent().ProjectID)
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(krm.TrimPrefix(a.actual.Name))
	u.SetGroupVersionKind(krm.ReCAPTCHAEnterpriseFirewallPolicyGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *FirewallPolicyAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting FirewallPolicy", "name", a.id)

	req := &pb.DeleteFirewallPolicyRequest{Name: a.id.String()}
	_, err := a.gcpClient.DeleteFirewallPolicy(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent FirewallPolicy, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting FirewallPolicy %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted FirewallPolicy", "name", a.id)

	return true, nil
}
