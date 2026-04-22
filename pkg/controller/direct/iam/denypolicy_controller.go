// Copyright 2025 Google LLC
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

package iam

import (
	"context"
	"fmt"
	"net/url"

	iamv2 "cloud.google.com/go/iam/apiv2"
	pb "cloud.google.com/go/iam/apiv2/iampb"

	krmv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/iam/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krmv1alpha1.IAMDenyPolicyGVK, NewIAMDenyPolicyModel)
}

func NewIAMDenyPolicyModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelIAMDenyPolicy{config: *config}, nil
}

var _ directbase.Model = &modelIAMDenyPolicy{}

type modelIAMDenyPolicy struct {
	config config.ControllerConfig
}

func (m *modelIAMDenyPolicy) client(ctx context.Context) (*iamv2.PoliciesClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := iamv2.NewPoliciesRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building iam policies client: %w", err)
	}
	return gcpClient, err
}

func (m *modelIAMDenyPolicy) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	obj := &krmv1alpha1.IAMDenyPolicy{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(op.GetUnstructured().Object, obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	// Normalize parent reference
	if err := obj.Spec.ParentRef.Normalize(ctx, op.Reader, obj.GetNamespace()); err != nil {
		return nil, fmt.Errorf("normalizing parent reference: %w", err)
	}

	return &IAMDenyPolicyAdapter{
		id:        obj.Spec.ResourceID,
		desired:   obj,
		gcpClient: gcpClient,
	}, nil
}

func (m *modelIAMDenyPolicy) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type IAMDenyPolicyAdapter struct {
	id        *string
	desired   *krmv1alpha1.IAMDenyPolicy
	actual    *pb.Policy
	gcpClient *iamv2.PoliciesClient
}

var _ directbase.Adapter = &IAMDenyPolicyAdapter{}

func (a *IAMDenyPolicyAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx).WithName("IAMDenyPolicyAdapter.Find")
	log.V(2).Info("finding IAMDenyPolicy", "name", a.desired.Name)

	req := &pb.GetPolicyRequest{
		Name: a.fullyQualifiedName(),
	}

	policy, err := a.gcpClient.GetPolicy(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting IAMDenyPolicy %q: %w", req.Name, err)
	}

	a.actual = policy
	return true, nil
}

func (a *IAMDenyPolicyAdapter) Create(ctx context.Context, op *directbase.CreateOperation) error {
	log := klog.FromContext(ctx).WithName("IAMDenyPolicyAdapter.Create")
	log.V(2).Info("creating IAMDenyPolicy", "name", a.desired.Name)

	mapCtx := &direct.MapContext{}
	desiredProto := IAMDenyPolicySpec_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &pb.CreatePolicyRequest{
		Parent:   a.parent(),
		PolicyId: a.resourceID(),
		Policy:   desiredProto,
	}

	opResult, err := a.gcpClient.CreatePolicy(ctx, req)
	if err != nil {
		return fmt.Errorf("creating IAMDenyPolicy %q: %w", req.PolicyId, err)
	}
	created, err := opResult.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for creation of IAMDenyPolicy %q: %w", req.PolicyId, err)
	}
	log.V(2).Info("successfully created IAMDenyPolicy", "name", a.desired.Name)

	status := &krmv1alpha1.IAMDenyPolicyStatus{}
	status.ObservedState = IAMDenyPolicyObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	external := created.Name
	status.ExternalRef = &external
	return op.UpdateStatus(ctx, status, nil)
}

func (a *IAMDenyPolicyAdapter) Update(ctx context.Context, op *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx).WithName("IAMDenyPolicyAdapter.Update")
	log.V(2).Info("updating IAMDenyPolicy", "name", a.desired.Name)

	mapCtx := &direct.MapContext{}
	desiredProto := IAMDenyPolicySpec_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	desiredProto.Name = a.fullyQualifiedName()
	desiredProto.Etag = a.actual.Etag

	req := &pb.UpdatePolicyRequest{
		Policy: desiredProto,
	}

	opResult, err := a.gcpClient.UpdatePolicy(ctx, req)
	if err != nil {
		return fmt.Errorf("updating IAMDenyPolicy %q: %w", req.Policy.Name, err)
	}
	updated, err := opResult.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for update of IAMDenyPolicy %q: %w", req.Policy.Name, err)
	}
	log.V(2).Info("successfully updated IAMDenyPolicy", "name", a.desired.Name)

	status := &krmv1alpha1.IAMDenyPolicyStatus{}
	status.ObservedState = IAMDenyPolicyObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	external := updated.Name
	status.ExternalRef = &external
	return op.UpdateStatus(ctx, status, nil)
}

func (a *IAMDenyPolicyAdapter) Delete(ctx context.Context, op *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx).WithName("IAMDenyPolicyAdapter.Delete")
	log.V(2).Info("deleting IAMDenyPolicy", "name", a.desired.Name)

	req := &pb.DeletePolicyRequest{
		Name: a.fullyQualifiedName(),
	}

	opResult, err := a.gcpClient.DeletePolicy(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting IAMDenyPolicy %q: %w", req.Name, err)
	}
	_, err = opResult.Wait(ctx)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("waiting for deletion of IAMDenyPolicy %q: %w", req.Name, err)
	}
	log.V(2).Info("successfully deleted IAMDenyPolicy", "name", a.desired.Name)
	return true, nil
}

func (a *IAMDenyPolicyAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	return nil, nil
}

func (a *IAMDenyPolicyAdapter) resourceID() string {
	if a.desired.Spec.ResourceID != nil {
		return *a.desired.Spec.ResourceID
	}
	return a.desired.Name
}

func (a *IAMDenyPolicyAdapter) parent() string {
	ref := a.desired.Spec.ParentRef
	var attachmentPoint string

	kind := ref.Kind
	if kind == "" {
		kind = "Project"
	}

	switch kind {
	case "Project":
		attachmentPoint = "cloudresourcemanager.googleapis.com/projects/" + ref.External
	case "Folder":
		attachmentPoint = "cloudresourcemanager.googleapis.com/folders/" + ref.External
	case "Organization":
		attachmentPoint = "cloudresourcemanager.googleapis.com/organizations/" + ref.External
	}

	return "policies/" + url.PathEscape(attachmentPoint)
}

func (a *IAMDenyPolicyAdapter) fullyQualifiedName() string {
	return fmt.Sprintf("%s/denypolicies/%s", a.parent(), a.resourceID())
}
