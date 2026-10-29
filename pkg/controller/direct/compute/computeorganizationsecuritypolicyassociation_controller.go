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
	"errors"
	"fmt"
	"strings"

	"google.golang.org/api/googleapi"
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
	registry.RegisterModel(krm.ComputeOrganizationSecurityPolicyAssociationGVK, NewComputeOrganizationSecurityPolicyAssociationModel)
}

func NewComputeOrganizationSecurityPolicyAssociationModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &computeOrganizationSecurityPolicyAssociationModel{config: config}, nil
}

type computeOrganizationSecurityPolicyAssociationModel struct {
	config *config.ControllerConfig
}

var _ directbase.Model = &computeOrganizationSecurityPolicyAssociationModel{}

type computeOrganizationSecurityPolicyAssociationAdapter struct {
	id        *krm.ComputeOrganizationSecurityPolicyAssociationIdentity
	gcpClient *gcp.OrganizationSecurityPoliciesClient
	desired   *krm.ComputeOrganizationSecurityPolicyAssociation
	actual    *computepb.SecurityPolicyAssociation
	reader    client.Reader
}

var _ directbase.Adapter = &computeOrganizationSecurityPolicyAssociationAdapter{}

func (m *computeOrganizationSecurityPolicyAssociationModel) client(ctx context.Context) (*gcp.OrganizationSecurityPoliciesClient, error) {
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

func (m *computeOrganizationSecurityPolicyAssociationModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ComputeOrganizationSecurityPolicyAssociation{}
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
	id, ok := rawID.(*krm.ComputeOrganizationSecurityPolicyAssociationIdentity)
	if !ok {
		return nil, fmt.Errorf("unexpected identity type %T", rawID)
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &computeOrganizationSecurityPolicyAssociationAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
		reader:    reader,
	}, nil
}

func (m *computeOrganizationSecurityPolicyAssociationModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	id := &krm.ComputeOrganizationSecurityPolicyAssociationIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &computeOrganizationSecurityPolicyAssociationAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}

func isAssociationNotFound(err error) bool {
	if err == nil {
		return false
	}
	if direct.IsNotFound(err) {
		return true
	}
	var gerr *googleapi.Error
	if errors.As(err, &gerr) {
		if gerr.Code == 400 && strings.Contains(gerr.Message, "An association with that name does not exist") {
			return true
		}
	}
	return false
}

func (a *computeOrganizationSecurityPolicyAssociationAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ComputeOrganizationSecurityPolicyAssociation", "id", a.id)

	if a.id.SecurityPolicy == "" || a.id.Association == "" {
		return false, nil
	}

	association, err := a.get(ctx)
	if err != nil {
		if isAssociationNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ComputeOrganizationSecurityPolicyAssociation %s: %w", a.id, err)
	}

	a.actual = association
	return true, nil
}

func (a *computeOrganizationSecurityPolicyAssociationAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating ComputeOrganizationSecurityPolicyAssociation", "id", a.id)

	mapCtx := &direct.MapContext{}
	desiredProto := ComputeOrganizationSecurityPolicyAssociationSpec_v1alpha1_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	desiredProto.Name = direct.LazyPtr(a.id.Association)

	req := &computepb.AddAssociationOrganizationSecurityPolicyRequest{
		SecurityPolicy:                    a.id.SecurityPolicy,
		SecurityPolicyAssociationResource: desiredProto,
	}

	op, err := a.gcpClient.AddAssociation(ctx, req)
	if err != nil {
		return fmt.Errorf("creating ComputeOrganizationSecurityPolicyAssociation: %w", err)
	}
	if !op.Done() {
		err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("waiting ComputeOrganizationSecurityPolicyAssociation create failed: %w", err)
		}
	}
	log.V(2).Info("successfully created ComputeOrganizationSecurityPolicyAssociation")

	created, err := a.get(ctx)
	if err != nil {
		return fmt.Errorf("getting created ComputeOrganizationSecurityPolicyAssociation %s: %w", a.id, err)
	}

	return a.updateStatus(ctx, createOp, created)
}

func (a *computeOrganizationSecurityPolicyAssociationAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating ComputeOrganizationSecurityPolicyAssociation", "id", a.id)

	mapCtx := &direct.MapContext{}
	desiredProto := ComputeOrganizationSecurityPolicyAssociationSpec_v1alpha1_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	desiredProto.Name = direct.LazyPtr(a.id.Association)

	diffs, err := compareOrganizationSecurityPolicyAssociation(ctx, a.actual, desiredProto)
	if err != nil {
		return err
	}

	if !diffs.HasDiff() {
		log.V(2).Info("no diff detected for ComputeOrganizationSecurityPolicyAssociation", "id", a.id)
		return a.updateStatus(ctx, updateOp, a.actual)
	}

	// ComputeOrganizationSecurityPolicyAssociation in GCP only supports AddAssociation (with replaceExistingAssociation) or RemoveAssociation.
	// However, its spec fields (attachmentId, policyId, securityPolicyRef, resourceID) are immutable.
	return fmt.Errorf("ComputeOrganizationSecurityPolicyAssociation %s is immutable and cannot be updated", a.id)
}

func (a *computeOrganizationSecurityPolicyAssociationAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("ComputeOrganizationSecurityPolicyAssociation %s not found", a.id)
	}

	mc := &direct.MapContext{}
	spec := ComputeOrganizationSecurityPolicyAssociationSpec_v1alpha1_FromProto(mc, a.actual)
	if mc.Err() != nil {
		return nil, mc.Err()
	}

	spec.PolicyID = direct.LazyPtr(a.id.SecurityPolicy)

	specObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(spec)
	if err != nil {
		return nil, fmt.Errorf("error converting ComputeOrganizationSecurityPolicyAssociation spec to unstructured: %w", err)
	}

	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(krm.ComputeOrganizationSecurityPolicyAssociationGVK)

	if err := unstructured.SetNestedField(u.Object, specObj, "spec"); err != nil {
		return nil, fmt.Errorf("setting spec: %w", err)
	}

	return u, nil
}

func (a *computeOrganizationSecurityPolicyAssociationAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ComputeOrganizationSecurityPolicyAssociation", "id", a.id)

	if a.id.SecurityPolicy == "" || a.id.Association == "" {
		return false, fmt.Errorf("cannot delete ComputeOrganizationSecurityPolicyAssociation without an ID")
	}

	req := &computepb.RemoveAssociationOrganizationSecurityPolicyRequest{
		SecurityPolicy: a.id.SecurityPolicy,
		Name:           direct.LazyPtr(a.id.Association),
	}

	op, err := a.gcpClient.RemoveAssociation(ctx, req)
	if err != nil {
		if isAssociationNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting ComputeOrganizationSecurityPolicyAssociation %s: %w", a.id, err)
	}
	if !op.Done() {
		err = op.Wait(ctx)
		if err != nil {
			if isAssociationNotFound(err) {
				return true, nil
			}
			return false, fmt.Errorf("waiting ComputeOrganizationSecurityPolicyAssociation %s delete failed: %w", a.id, err)
		}
	}
	log.V(2).Info("successfully deleted ComputeOrganizationSecurityPolicyAssociation", "id", a.id)
	return true, nil
}

func (a *computeOrganizationSecurityPolicyAssociationAdapter) get(ctx context.Context) (*computepb.SecurityPolicyAssociation, error) {
	getReq := &computepb.GetAssociationOrganizationSecurityPolicyRequest{
		SecurityPolicy: a.id.SecurityPolicy,
		Name:           direct.LazyPtr(a.id.Association),
	}
	return a.gcpClient.GetAssociation(ctx, getReq)
}

func (a *computeOrganizationSecurityPolicyAssociationAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *computepb.SecurityPolicyAssociation) error {
	mapCtx := &direct.MapContext{}
	status := ComputeOrganizationSecurityPolicyAssociationStatus_v1alpha1_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return op.UpdateStatus(ctx, status, nil)
}

func compareOrganizationSecurityPolicyAssociation(ctx context.Context, actual, desired *computepb.SecurityPolicyAssociation) (*structuredreporting.Diff, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, ComputeOrganizationSecurityPolicyAssociationSpec_v1alpha1_FromProto, ComputeOrganizationSecurityPolicyAssociationSpec_v1alpha1_ToProto)
	if err != nil {
		return nil, err
	}
	maskedActual.Name = desired.Name

	clonedDesired := proto.CloneOf(desired)

	diffs, _, err := common.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, err
	}
	return diffs, nil
}
