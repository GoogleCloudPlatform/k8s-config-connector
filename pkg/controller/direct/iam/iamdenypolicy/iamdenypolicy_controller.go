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

package iamdenypolicy

import (
	"context"
	"fmt"

	"google.golang.org/api/option"

	gcp_iam "cloud.google.com/go/iam/apiv2"
	pb "cloud.google.com/go/iam/apiv2/iampb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/iam/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	mappers "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/iam"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.IAMDenyPolicyGVK, NewModel)
}

func NewModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &model{config: *config}, nil
}

var _ directbase.Model = &model{}

type model struct {
	config config.ControllerConfig
}

func (m *model) client(ctx context.Context) (*gcp_iam.PoliciesClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp_iam.NewPoliciesRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building policies client: %w", err)
	}
	return gcpClient, err
}

func (m *model) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	obj := &krm.IAMDenyPolicy{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(op.Object.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := obj.GetIdentity(ctx, op.Reader)
	if err != nil {
		return nil, err
	}

	// Get gcp client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &Adapter{
		id:        id.(*krm.IAMDenyPolicyIdentity),
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type Adapter struct {
	id        *krm.IAMDenyPolicyIdentity
	gcpClient *gcp_iam.PoliciesClient
	desired   *krm.IAMDenyPolicy
	actual    *pb.Policy
}

var _ directbase.Adapter = &Adapter{}

func (a *Adapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting IAMDenyPolicy", "name", a.id.String())

	req := &pb.GetPolicyRequest{
		Name: a.id.String(),
	}
	policy, err := a.gcpClient.GetPolicy(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting IAMDenyPolicy %q: %w", a.id.String(), err)
	}

	a.actual = policy

	return true, nil
}

func (a *Adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating IAMDenyPolicy", "name", a.id.String())
	mapCtx := &direct.MapContext{}

	desiredProto := mappers.IAMDenyPolicySpec_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	parent := "policies/" + a.id.AttachmentPoint + "/denypolicies"

	req := &pb.CreatePolicyRequest{
		Parent:   parent,
		Policy:   desiredProto,
		PolicyId: a.id.DenyPolicy,
	}
	op, err := a.gcpClient.CreatePolicy(ctx, req)
	if err != nil {
		return fmt.Errorf("creating IAMDenyPolicy %q: %w", a.id.String(), err)
	}

	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for IAMDenyPolicy %q creation: %w", a.id.String(), err)
	}

	log.V(2).Info("successfully created IAMDenyPolicy", "name", a.id.String())

	observedState := mappers.IAMDenyPolicyObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	a.desired.Status.ObservedState = observedState
	a.desired.Status.ExternalRef = direct.LazyPtr(a.id.String())

	return nil
}

func (a *Adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating IAMDenyPolicy", "name", a.id.String())
	mapCtx := &direct.MapContext{}

	desiredProto := mappers.IAMDenyPolicySpec_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	desiredProto.Name = a.actual.Name
	desiredProto.Uid = a.actual.Uid
	desiredProto.Kind = a.actual.Kind
	desiredProto.Etag = a.actual.Etag
	desiredProto.Annotations = a.actual.Annotations
	desiredProto.CreateTime = a.actual.CreateTime
	desiredProto.UpdateTime = a.actual.UpdateTime
	desiredProto.DeleteTime = a.actual.DeleteTime
	desiredProto.ManagingAuthority = a.actual.ManagingAuthority

	paths, err := common.CompareProtoMessage(desiredProto, a.actual, common.BasicDiff)
	if err != nil {
		return err
	}

	if paths.Len() == 0 {
		log.V(2).Info("no changes detected for IAMDenyPolicy, skipping update", "name", a.id.String())
		return nil
	}

	getReq := &pb.GetPolicyRequest{
		Name: a.id.String(),
	}
	existing, err := a.gcpClient.GetPolicy(ctx, getReq)
	if err != nil {
		return fmt.Errorf("getting IAMDenyPolicy %q before update: %w", a.id.String(), err)
	}
	desiredProto.Etag = existing.Etag

	req := &pb.UpdatePolicyRequest{
		Policy: desiredProto,
	}

	op, err := a.gcpClient.UpdatePolicy(ctx, req)
	if err != nil {
		return fmt.Errorf("updating IAMDenyPolicy %q: %w", a.id.String(), err)
	}

	updated, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for IAMDenyPolicy %q update: %w", a.id.String(), err)
	}

	log.V(2).Info("successfully updated IAMDenyPolicy", "name", a.id.String())

	observedState := mappers.IAMDenyPolicyObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	a.desired.Status.ObservedState = observedState
	a.desired.Status.ExternalRef = direct.LazyPtr(a.id.String())

	return nil
}

func (a *Adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting IAMDenyPolicy", "name", a.id.String())

	req := &pb.DeletePolicyRequest{
		Name: a.id.String(),
	}

	op, err := a.gcpClient.DeletePolicy(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting IAMDenyPolicy %q: %w", a.id.String(), err)
	}

	_, err = op.Wait(ctx)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("waiting for IAMDenyPolicy %q deletion: %w", a.id.String(), err)
	}

	log.V(2).Info("successfully deleted IAMDenyPolicy", "name", a.id.String())
	return true, nil
}

func (a *Adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.IAMDenyPolicy{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(mappers.IAMDenyPolicySpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	// Restore references
	obj.Spec.ProjectRef = a.desired.Spec.ProjectRef
	obj.Spec.FolderRef = a.desired.Spec.FolderRef
	obj.Spec.OrganizationRef = a.desired.Spec.OrganizationRef

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}
	u.Object = uObj
	return u, nil
}
