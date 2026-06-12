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

// +tool:controller
// proto.service: google.cloud.compute.v1.SecurityPolicies
// proto.message: google.cloud.compute.v1.SecurityPolicy
// crd.type: ComputeSecurityPolicy
// crd.version: v1beta1

package compute

import (
	"context"
	"fmt"

	compute "cloud.google.com/go/compute/apiv1"
	computepb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.ComputeSecurityPolicyGVK, NewComputeSecurityPolicyModel)
}

func NewComputeSecurityPolicyModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelComputeSecurityPolicy{config: config}, nil
}

var _ directbase.Model = &modelComputeSecurityPolicy{}

type modelComputeSecurityPolicy struct {
	config *config.ControllerConfig
}

func (m *modelComputeSecurityPolicy) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ComputeSecurityPolicy{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}

	gcpClient, err := newGCPClient(m.config)
	if err != nil {
		return nil, err
	}

	adapter := &ComputeSecurityPolicyAdapter{
		id:      id.(*krm.ComputeSecurityPolicyIdentity),
		desired: obj,
		reader:  reader,
	}

	if adapter.id.Region == "" {
		securityPoliciesClient, err := gcpClient.newSecurityPoliciesClient(ctx)
		if err != nil {
			return nil, err
		}
		adapter.securityPoliciesClient = securityPoliciesClient
	} else {
		regionSecurityPoliciesClient, err := gcpClient.newRegionSecurityPoliciesClient(ctx)
		if err != nil {
			return nil, err
		}
		adapter.regionSecurityPoliciesClient = regionSecurityPoliciesClient
	}

	return adapter, nil
}

func (m *modelComputeSecurityPolicy) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type ComputeSecurityPolicyAdapter struct {
	securityPoliciesClient       *compute.SecurityPoliciesClient
	regionSecurityPoliciesClient *compute.RegionSecurityPoliciesClient
	id                           *krm.ComputeSecurityPolicyIdentity
	desired                      *krm.ComputeSecurityPolicy
	actual                       *computepb.SecurityPolicy
	reader                       client.Reader
}

var _ directbase.Adapter = &ComputeSecurityPolicyAdapter{}

func (a *ComputeSecurityPolicyAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ComputeSecurityPolicy", "name", a.id)

	actual, err := a.get(ctx)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ComputeSecurityPolicy %q: %w", a.id, err)
	}

	a.actual = actual
	return true, nil
}

func (a *ComputeSecurityPolicyAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating ComputeSecurityPolicy", "name", a.id)
	mapCtx := &direct.MapContext{}

	if err := common.NormalizeReferences(ctx, a.reader, a.desired, nil); err != nil {
		return fmt.Errorf("normalizing references: %w", err)
	}

	desired := a.desired.DeepCopy()
	resource := ComputeSecurityPolicySpec_v1beta1_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	resource.Name = direct.LazyPtr(a.id.Name)

	var op *compute.Operation
	var err error
	if a.id.Region == "" {
		req := &computepb.InsertSecurityPolicyRequest{
			Project:                a.id.Project,
			SecurityPolicyResource: resource,
		}
		op, err = a.securityPoliciesClient.Insert(ctx, req)
	} else {
		req := &computepb.InsertRegionSecurityPolicyRequest{
			Project:                a.id.Project,
			Region:                 a.id.Region,
			SecurityPolicyResource: resource,
		}
		op, err = a.regionSecurityPoliciesClient.Insert(ctx, req)
	}
	if err != nil {
		return fmt.Errorf("creating ComputeSecurityPolicy %s: %w", a.id, err)
	}

	err = op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("compute ComputeSecurityPolicy %s waiting creation: %w", a.id.String(), err)
	}

	log.V(2).Info("successfully created ComputeSecurityPolicy", "name", a.id)

	created, err := a.get(ctx)
	if err != nil {
		return fmt.Errorf("getting ComputeSecurityPolicy %s: %w", a.id, err)
	}

	status := &krm.ComputeSecurityPolicyStatus{}
	status.Fingerprint = created.Fingerprint
	status.SelfLink = created.SelfLink
	status.ObservedGeneration = direct.LazyPtr(a.desired.Generation)
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *ComputeSecurityPolicyAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating ComputeSecurityPolicy", "name", a.id)
	mapCtx := &direct.MapContext{}

	if err := common.NormalizeReferences(ctx, a.reader, a.desired, nil); err != nil {
		return fmt.Errorf("normalizing references: %w", err)
	}

	desired := a.desired.DeepCopy()
	resource := ComputeSecurityPolicySpec_v1beta1_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	resource.Name = direct.LazyPtr(a.id.Name)
	resource.Fingerprint = a.actual.Fingerprint

	// Exclude output-only fields like Etag, SelfLink, ID etc. from comparison.
	// We're comparing against a.actual.
	paths, err := common.CompareProtoMessage(resource, a.actual, common.BasicDiff)
	if err != nil {
		return err
	}

	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		status := &krm.ComputeSecurityPolicyStatus{}
		status.Fingerprint = a.actual.Fingerprint
		status.SelfLink = a.actual.SelfLink
		status.ObservedGeneration = direct.LazyPtr(a.desired.Generation)
		return updateOp.UpdateStatus(ctx, status, nil)
	}

	report := &structuredreporting.Diff{Object: updateOp.GetUnstructured()}
	for path := range paths {
		report.AddField(path, nil, nil)
	}
	structuredreporting.ReportDiff(ctx, report)

	var op *compute.Operation
	if a.id.Region == "" {
		req := &computepb.PatchSecurityPolicyRequest{
			Project:                a.id.Project,
			SecurityPolicy:         a.id.Name,
			SecurityPolicyResource: resource,
		}
		op, err = a.securityPoliciesClient.Patch(ctx, req)
	} else {
		req := &computepb.PatchRegionSecurityPolicyRequest{
			Project:                a.id.Project,
			Region:                 a.id.Region,
			SecurityPolicy:         a.id.Name,
			SecurityPolicyResource: resource,
		}
		op, err = a.regionSecurityPoliciesClient.Patch(ctx, req)
	}
	if err != nil {
		return fmt.Errorf("updating ComputeSecurityPolicy %s: %w", a.id, err)
	}

	err = op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("compute ComputeSecurityPolicy %s waiting update: %w", a.id.String(), err)
	}

	updated, err := a.get(ctx)
	if err != nil {
		return fmt.Errorf("getting ComputeSecurityPolicy %s: %w", a.id, err)
	}

	status := &krm.ComputeSecurityPolicyStatus{}
	status.Fingerprint = updated.Fingerprint
	status.SelfLink = updated.SelfLink
	status.ObservedGeneration = direct.LazyPtr(a.desired.Generation)
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *ComputeSecurityPolicyAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

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

	u.SetName(a.id.Name)
	u.SetGroupVersionKind(krm.ComputeSecurityPolicyGVK)

	u.Object = uObj
	return u, nil
}

func (a *ComputeSecurityPolicyAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ComputeSecurityPolicy", "name", a.id)

	var op *compute.Operation
	var err error
	if a.id.Region == "" {
		req := &computepb.DeleteSecurityPolicyRequest{
			Project:        a.id.Project,
			SecurityPolicy: a.id.Name,
		}
		op, err = a.securityPoliciesClient.Delete(ctx, req)
	} else {
		req := &computepb.DeleteRegionSecurityPolicyRequest{
			Project:        a.id.Project,
			Region:         a.id.Region,
			SecurityPolicy: a.id.Name,
		}
		op, err = a.regionSecurityPoliciesClient.Delete(ctx, req)
	}
	if err != nil {
		return false, fmt.Errorf("deleting ComputeSecurityPolicy %s: %w", a.id.String(), err)
	}
	log.Info("successfully deleted ComputeSecurityPolicy", "name", a.id)

	if !op.Done() {
		err = op.Wait(ctx)
		if err != nil {
			return false, fmt.Errorf("waiting for deletion of ComputeSecurityPolicy %s: %w", a.id.String(), err)
		}
	}

	return true, nil
}

func (a *ComputeSecurityPolicyAdapter) get(ctx context.Context) (*computepb.SecurityPolicy, error) {
	if a.id.Region == "" {
		getReq := &computepb.GetSecurityPolicyRequest{
			Project:        a.id.Project,
			SecurityPolicy: a.id.Name,
		}
		return a.securityPoliciesClient.Get(ctx, getReq)
	} else {
		getReq := &computepb.GetRegionSecurityPolicyRequest{
			Project:        a.id.Project,
			Region:         a.id.Region,
			SecurityPolicy: a.id.Name,
		}
		return a.regionSecurityPoliciesClient.Get(ctx, getReq)
	}
}
