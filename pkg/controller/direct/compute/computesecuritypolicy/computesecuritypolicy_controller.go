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

package computesecuritypolicy

import (
	"context"
	"fmt"

	gcp "cloud.google.com/go/compute/apiv1"
	computepb "cloud.google.com/go/compute/apiv1/computepb"
	"google.golang.org/api/option"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.ComputeSecurityPolicyGVK, NewComputeSecurityPolicyModel)
}

func NewComputeSecurityPolicyModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &computesecuritypolicyModel{config: config}, nil
}

type computesecuritypolicyModel struct {
	config *config.ControllerConfig
}

// model implements the Model interface.
var _ directbase.Model = &computesecuritypolicyModel{}

type computesecuritypolicyAdapter struct {
	id             *krm.ComputeSecurityPolicyIdentity
	globalClient   *gcp.SecurityPoliciesClient
	regionalClient *gcp.RegionSecurityPoliciesClient
	desired        *krm.ComputeSecurityPolicy
	actual         *computepb.SecurityPolicy
}

var _ directbase.Adapter = &computesecuritypolicyAdapter{}

func (m *computesecuritypolicyModel) globalClient(ctx context.Context) (*gcp.SecurityPoliciesClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewSecurityPoliciesRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building global SecurityPolicies client: %w", err)
	}
	return gcpClient, nil
}

func (m *computesecuritypolicyModel) regionalClient(ctx context.Context) (*gcp.RegionSecurityPoliciesClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRegionSecurityPoliciesRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building regional SecurityPolicies client: %w", err)
	}
	return gcpClient, nil
}

func (m *computesecuritypolicyModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ComputeSecurityPolicy{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	idInterface, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id, ok := idInterface.(*krm.ComputeSecurityPolicyIdentity)
	if !ok {
		return nil, fmt.Errorf("unexpected identity type: %T", idInterface)
	}

	computesecuritypolicyAdapter := &computesecuritypolicyAdapter{
		id:      id,
		desired: obj,
	}

	if id.Region != "" {
		gcpClient, err := m.regionalClient(ctx)
		if err != nil {
			return nil, fmt.Errorf("building regional gcp client: %w", err)
		}
		computesecuritypolicyAdapter.regionalClient = gcpClient
	} else {
		gcpClient, err := m.globalClient(ctx)
		if err != nil {
			return nil, fmt.Errorf("building global gcp client: %w", err)
		}
		computesecuritypolicyAdapter.globalClient = gcpClient
	}

	return computesecuritypolicyAdapter, nil
}

func (m *computesecuritypolicyModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

func (a *computesecuritypolicyAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ComputeSecurityPolicy", "id", a.id.String())

	var policy *computepb.SecurityPolicy
	var err error

	if a.id.Region != "" {
		req := &computepb.GetRegionSecurityPolicyRequest{
			Project:        a.id.Project,
			Region:         a.id.Region,
			SecurityPolicy: a.id.Name,
		}
		policy, err = a.regionalClient.Get(ctx, req)
	} else {
		req := &computepb.GetSecurityPolicyRequest{
			Project:        a.id.Project,
			SecurityPolicy: a.id.Name,
		}
		policy, err = a.globalClient.Get(ctx, req)
	}

	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ComputeSecurityPolicy %s: %w", a.id.String(), err)
	}

	a.actual = policy
	return true, nil
}

func (a *computesecuritypolicyAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating ComputeSecurityPolicy", "id", a.id.String())

	mapCtx := &direct.MapContext{}
	desired := a.desired.DeepCopy()

	policy := ComputeSecurityPolicySpec_v1beta1_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	policy.Name = direct.LazyPtr(a.id.Name)

	var op *gcp.Operation
	var err error

	if a.id.Region != "" {
		req := &computepb.InsertRegionSecurityPolicyRequest{
			Project:                a.id.Project,
			Region:                 a.id.Region,
			SecurityPolicyResource: policy,
		}
		op, err = a.regionalClient.Insert(ctx, req)
	} else {
		req := &computepb.InsertSecurityPolicyRequest{
			Project:                a.id.Project,
			SecurityPolicyResource: policy,
		}
		op, err = a.globalClient.Insert(ctx, req)
	}

	if err != nil {
		return fmt.Errorf("creating ComputeSecurityPolicy %s: %w", a.id.String(), err)
	}

	err = op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for ComputeSecurityPolicy %s creation: %w", a.id.String(), err)
	}

	// Fetch actual again to get full state (like fingerprint and selfLink)
	if _, err := a.Find(ctx); err != nil {
		return fmt.Errorf("fetching created ComputeSecurityPolicy %s: %w", a.id.String(), err)
	}

	status := &krm.ComputeSecurityPolicyStatus{}
	status.Fingerprint = a.actual.Fingerprint
	status.SelfLink = a.actual.SelfLink

	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *computesecuritypolicyAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating ComputeSecurityPolicy", "id", a.id.String())

	mapCtx := &direct.MapContext{}
	desired := a.desired.DeepCopy()

	policy := ComputeSecurityPolicySpec_v1beta1_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	policy.Name = direct.LazyPtr(a.id.Name)
	policy.Fingerprint = a.actual.Fingerprint

	var op *gcp.Operation
	var err error

	if a.id.Region != "" {
		req := &computepb.PatchRegionSecurityPolicyRequest{
			Project:                a.id.Project,
			Region:                 a.id.Region,
			SecurityPolicy:         a.id.Name,
			SecurityPolicyResource: policy,
		}
		op, err = a.regionalClient.Patch(ctx, req)
	} else {
		req := &computepb.PatchSecurityPolicyRequest{
			Project:                a.id.Project,
			SecurityPolicy:         a.id.Name,
			SecurityPolicyResource: policy,
		}
		op, err = a.globalClient.Patch(ctx, req)
	}

	if err != nil {
		return fmt.Errorf("updating ComputeSecurityPolicy %s: %w", a.id.String(), err)
	}

	err = op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for ComputeSecurityPolicy %s update: %w", a.id.String(), err)
	}

	// Fetch actual again to get updated state (including new fingerprint)
	if _, err := a.Find(ctx); err != nil {
		return fmt.Errorf("fetching updated ComputeSecurityPolicy %s: %w", a.id.String(), err)
	}

	status := &krm.ComputeSecurityPolicyStatus{}
	status.Fingerprint = a.actual.Fingerprint
	status.SelfLink = a.actual.SelfLink

	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *computesecuritypolicyAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ComputeSecurityPolicy", "id", a.id.String())

	var op *gcp.Operation
	var err error

	if a.id.Region != "" {
		req := &computepb.DeleteRegionSecurityPolicyRequest{
			Project:        a.id.Project,
			Region:         a.id.Region,
			SecurityPolicy: a.id.Name,
		}
		op, err = a.regionalClient.Delete(ctx, req)
	} else {
		req := &computepb.DeleteSecurityPolicyRequest{
			Project:        a.id.Project,
			SecurityPolicy: a.id.Name,
		}
		op, err = a.globalClient.Delete(ctx, req)
	}

	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting ComputeSecurityPolicy %s: %w", a.id.String(), err)
	}

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting for ComputeSecurityPolicy %s deletion: %w", a.id.String(), err)
	}

	return true, nil
}

func (a *computesecuritypolicyAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}

	obj := &krm.ComputeSecurityPolicy{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(ComputeSecurityPolicySpec_v1beta1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u := &unstructured.Unstructured{Object: uObj}
	u.SetName(a.id.Name)
	u.SetGroupVersionKind(krm.ComputeSecurityPolicyGVK)
	return u, nil
}
