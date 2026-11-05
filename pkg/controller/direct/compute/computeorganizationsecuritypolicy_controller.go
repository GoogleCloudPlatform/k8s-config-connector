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

package compute

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	gcp "cloud.google.com/go/compute/apiv1"
	computepb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.ComputeOrganizationSecurityPolicyGVK, NewComputeOrganizationSecurityPolicyModel)
}

func NewComputeOrganizationSecurityPolicyModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &computeOrganizationSecurityPolicyModel{config: config}, nil
}

type computeOrganizationSecurityPolicyModel struct {
	config *config.ControllerConfig
}

var _ directbase.Model = &computeOrganizationSecurityPolicyModel{}

type computeOrganizationSecurityPolicyAdapter struct {
	id        *krm.ComputeOrganizationSecurityPolicyIdentity
	gcpClient *gcp.OrganizationSecurityPoliciesClient
	desired   *krm.ComputeOrganizationSecurityPolicy
	actual    *computepb.SecurityPolicy
	reader    client.Reader
}

var _ directbase.Adapter = &computeOrganizationSecurityPolicyAdapter{}

func (m *computeOrganizationSecurityPolicyModel) client(ctx context.Context) (*gcp.OrganizationSecurityPoliciesClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewOrganizationSecurityPoliciesRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building OrganizationSecurityPolicies REST client: %w", err)
	}
	return gcpClient, nil
}

func (m *computeOrganizationSecurityPolicyModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ComputeOrganizationSecurityPolicy{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	rawID, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id, ok := rawID.(*krm.ComputeOrganizationSecurityPolicyIdentity)
	if !ok {
		return nil, fmt.Errorf("unexpected identity type %T", rawID)
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &computeOrganizationSecurityPolicyAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
		reader:    reader,
	}, nil
}

func (m *computeOrganizationSecurityPolicyModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	id := &krm.ComputeOrganizationSecurityPolicyIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &computeOrganizationSecurityPolicyAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}

func (a *computeOrganizationSecurityPolicyAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ComputeOrganizationSecurityPolicy", "id", a.id)

	if !a.id.HasIdentitySpecified() {
		return false, nil
	}

	policy, err := a.get(ctx)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ComputeOrganizationSecurityPolicy %s: %w", a.id, err)
	}

	a.actual = policy
	return true, nil
}

func (a *computeOrganizationSecurityPolicyAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating ComputeOrganizationSecurityPolicy", "id", a.id)

	mapCtx := &direct.MapContext{}
	desiredProto := ComputeOrganizationSecurityPolicySpec_v1alpha1_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	parent := a.desired.Spec.Parent
	req := &computepb.InsertOrganizationSecurityPolicyRequest{
		ParentId:               &parent,
		SecurityPolicyResource: desiredProto,
	}

	op, err := a.gcpClient.Insert(ctx, req)
	if err != nil {
		return fmt.Errorf("creating ComputeOrganizationSecurityPolicy: %w", err)
	}
	if !op.Done() {
		err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("waiting ComputeOrganizationSecurityPolicy create failed: %w", err)
		}
	}
	log.V(2).Info("successfully created ComputeOrganizationSecurityPolicy")

	var policyID string
	if op.Proto() != nil {
		if op.Proto().TargetId != nil {
			policyID = strconv.FormatUint(*op.Proto().TargetId, 10)
		} else if op.Proto().TargetLink != nil {
			link := *op.Proto().TargetLink
			tokens := strings.Split(link, "/")
			policyID = tokens[len(tokens)-1]
		}
	}

	if policyID == "" {
		return fmt.Errorf("could not determine the security policy ID from operation")
	}

	a.id.SecurityPolicy = policyID
	created, err := a.get(ctx)
	if err != nil {
		return fmt.Errorf("getting created ComputeOrganizationSecurityPolicy %s: %w", a.id, err)
	}

	return a.updateStatus(ctx, createOp, created)
}

func (a *computeOrganizationSecurityPolicyAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating ComputeOrganizationSecurityPolicy", "id", a.id)

	mapCtx := &direct.MapContext{}
	desiredProto := ComputeOrganizationSecurityPolicySpec_v1alpha1_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	diffs, err := compareOrganizationSecurityPolicy(ctx, a.actual, desiredProto)
	if err != nil {
		return err
	}

	if !diffs.HasDiff() {
		log.V(2).Info("no diff detected for ComputeOrganizationSecurityPolicy", "id", a.id)
		return a.updateStatus(ctx, updateOp, a.actual)
	}

	diffs.Object = updateOp.GetUnstructured()
	structuredreporting.ReportDiff(ctx, diffs)

	patchResource := &computepb.SecurityPolicy{
		Description: desiredProto.Description,
		Fingerprint: a.actual.Fingerprint,
	}

	req := &computepb.PatchOrganizationSecurityPolicyRequest{
		SecurityPolicy:         a.id.SecurityPolicy,
		SecurityPolicyResource: patchResource,
	}

	op, err := a.gcpClient.Patch(ctx, req)
	if err != nil {
		return fmt.Errorf("updating ComputeOrganizationSecurityPolicy %s: %w", a.id, err)
	}
	if !op.Done() {
		err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("waiting ComputeOrganizationSecurityPolicy %s update failed: %w", a.id, err)
		}
	}
	log.V(2).Info("successfully updated ComputeOrganizationSecurityPolicy", "id", a.id)

	updated, err := a.get(ctx)
	if err != nil {
		return fmt.Errorf("getting ComputeOrganizationSecurityPolicy %s: %w", a.id, err)
	}

	return a.updateStatus(ctx, updateOp, updated)
}

func (a *computeOrganizationSecurityPolicyAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("ComputeOrganizationSecurityPolicy %s not found", a.id)
	}

	mc := &direct.MapContext{}
	spec := ComputeOrganizationSecurityPolicySpec_v1alpha1_FromProto(mc, a.actual)
	if mc.Err() != nil {
		return nil, mc.Err()
	}

	specObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(spec)
	if err != nil {
		return nil, fmt.Errorf("error converting ComputeOrganizationSecurityPolicy spec to unstructured: %w", err)
	}

	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(krm.ComputeOrganizationSecurityPolicyGVK)

	if err := unstructured.SetNestedField(u.Object, specObj, "spec"); err != nil {
		return nil, fmt.Errorf("setting spec: %w", err)
	}

	return u, nil
}

func (a *computeOrganizationSecurityPolicyAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ComputeOrganizationSecurityPolicy", "id", a.id)

	if !a.id.HasIdentitySpecified() {
		return false, fmt.Errorf("cannot delete ComputeOrganizationSecurityPolicy without an ID")
	}

	req := &computepb.DeleteOrganizationSecurityPolicyRequest{
		SecurityPolicy: a.id.SecurityPolicy,
	}

	op, err := a.gcpClient.Delete(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting ComputeOrganizationSecurityPolicy %s: %w", a.id, err)
	}
	if !op.Done() {
		err = op.Wait(ctx)
		if err != nil {
			return false, fmt.Errorf("waiting ComputeOrganizationSecurityPolicy %s delete failed: %w", a.id, err)
		}
	}
	log.V(2).Info("successfully deleted ComputeOrganizationSecurityPolicy", "id", a.id)
	return true, nil
}

func (a *computeOrganizationSecurityPolicyAdapter) get(ctx context.Context) (*computepb.SecurityPolicy, error) {
	getReq := &computepb.GetOrganizationSecurityPolicyRequest{
		SecurityPolicy: a.id.SecurityPolicy,
	}
	return a.gcpClient.Get(ctx, getReq)
}

func (a *computeOrganizationSecurityPolicyAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *computepb.SecurityPolicy) error {
	mapCtx := &direct.MapContext{}
	status := ComputeOrganizationSecurityPolicyStatus_v1alpha1_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return op.UpdateStatus(ctx, status, nil)
}

func compareOrganizationSecurityPolicy(ctx context.Context, actual, desired *computepb.SecurityPolicy) (*structuredreporting.Diff, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, ComputeOrganizationSecurityPolicySpec_v1alpha1_FromProto, ComputeOrganizationSecurityPolicySpec_v1alpha1_ToProto)
	if err != nil {
		return nil, err
	}
	maskedActual.Name = desired.Name
	maskedActual.Id = desired.Id
	maskedActual.Parent = desired.Parent
	maskedActual.ShortName = desired.ShortName
	maskedActual.Type = desired.Type

	clonedDesired := proto.CloneOf(desired)

	diffs, _, err := common.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, err
	}
	return diffs, nil
}
