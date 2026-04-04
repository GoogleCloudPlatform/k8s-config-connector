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

package iam

import (
	"context"
	"fmt"
	"net/url"

	api "cloud.google.com/go/iam/apiv2"
	pb "cloud.google.com/go/iam/apiv2/iampb"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/iam/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
)

func init() {
	registry.RegisterModel(krm.IAMDenyPolicyGVK, newDenyPolicyModel)
}

func newDenyPolicyModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &denyPolicyModel{config: *config}, nil
}

type denyPolicyModel struct {
	config config.ControllerConfig
}

// model implements the Model interface.
var _ directbase.Model = &denyPolicyModel{}

type denyPolicyAdapter struct {
	attachmentPoint string
	policyID        string

	desired *krm.IAMDenyPolicy
	actual  *pb.Policy

	gcp *api.PoliciesClient
}

// adapter implements the Adapter interface.
var _ directbase.Adapter = &denyPolicyAdapter{}

func (m *denyPolicyModel) client(ctx context.Context) (*api.PoliciesClient, error) {
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}

	gcpClient, err := api.NewPoliciesRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building iam policies client: %w", err)
	}
	return gcpClient, err
}

// AdapterForObject implements the Model interface.
func (m *denyPolicyModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	gcp, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	obj := &krm.IAMDenyPolicy{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	var attachmentPoint string
	if obj.Spec.ProjectRef != nil {
		projectRef, err := refs.ResolveProject(ctx, op.Reader, obj.GetNamespace(), obj.Spec.ProjectRef)
		if err != nil {
			return nil, err
		}
		attachmentPoint = url.PathEscape(fmt.Sprintf("cloudresourcemanager.googleapis.com/projects/%s", projectRef.ProjectID))
	} else if obj.Spec.FolderRef != nil {
		folderRef, err := refs.ResolveFolder(ctx, op.Reader, obj, obj.Spec.FolderRef)
		if err != nil {
			return nil, err
		}
		attachmentPoint = url.PathEscape(fmt.Sprintf("cloudresourcemanager.googleapis.com/folders/%s", folderRef.FolderID))
	} else if obj.Spec.OrganizationRef != nil {
		orgRef, err := refs.ResolveOrganization(ctx, op.Reader, obj, obj.Spec.OrganizationRef)
		if err != nil {
			return nil, err
		}
		attachmentPoint = url.PathEscape(fmt.Sprintf("cloudresourcemanager.googleapis.com/organizations/%s", orgRef.OrganizationID))
	} else {
		// Default to project from namespace
		projectID, err := refs.ResolveProjectID(ctx, op.Reader, u)
		if err != nil {
			return nil, fmt.Errorf("could not determine attachment point (specify projectRef, folderRef, organizationRef or namespace): %w", err)
		}
		attachmentPoint = url.PathEscape(fmt.Sprintf("cloudresourcemanager.googleapis.com/projects/%s", projectID))
	}

	policyID := direct.ValueOf(obj.Spec.ResourceID)
	if policyID == "" {
		policyID = obj.GetName()
	}

	return &denyPolicyAdapter{
		attachmentPoint: attachmentPoint,
		policyID:        policyID,
		desired:         obj,
		gcp:             gcp,
	}, nil
}

func (m *denyPolicyModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

// Find implements the Adapter interface.
func (a *denyPolicyAdapter) Find(ctx context.Context) (bool, error) {
	if a.policyID == "" {
		return false, nil
	}

	req := &pb.GetPolicyRequest{
		Name: a.fullyQualifiedName(),
	}
	policy, err := a.gcp.GetPolicy(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, err
	}

	a.actual = policy

	return true, nil
}

// Delete implements the Adapter interface.
func (a *denyPolicyAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	req := &pb.DeletePolicyRequest{
		Name: a.fullyQualifiedName(),
		Etag: a.actual.GetEtag(),
	}
	op, err := a.gcp.DeletePolicy(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting IAMDenyPolicy: %w", err)
	}

	if _, err := op.Wait(ctx); err != nil {
		return false, fmt.Errorf("waiting for IAMDenyPolicy deletion: %w", err)
	}

	return true, nil
}

// Create implements the Adapter interface.
func (a *denyPolicyAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	mapCtx := &direct.MapContext{}

	desired := IAMDenyPolicySpec_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &pb.CreatePolicyRequest{
		Parent:   fmt.Sprintf("policies/%s/denypolicies", a.attachmentPoint),
		PolicyId: a.policyID,
		Policy:   desired,
	}

	op, err := a.gcp.CreatePolicy(ctx, req)
	if err != nil {
		return fmt.Errorf("creating IAMDenyPolicy: %w", err)
	}

	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for IAMDenyPolicy creation: %w", err)
	}

	status := &krm.IAMDenyPolicyStatus{}
	status.ExternalRef = direct.LazyPtr(created.GetName())
	status.ObservedState = IAMDenyPolicyObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	return createOp.UpdateStatus(ctx, status, nil)
}

// Update implements the Adapter interface.
func (a *denyPolicyAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	mapCtx := &direct.MapContext{}

	desired := IAMDenyPolicySpec_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	// UpdatePolicy allows 'display_name', 'rules', and 'annotations'
	desired.Name = a.fullyQualifiedName()
	desired.Etag = a.actual.Etag

	req := &pb.UpdatePolicyRequest{
		Policy: desired,
	}

	// The API doesn't seem to support FieldMask in UpdatePolicyRequest,
	// but the mock implementation seems to expect one if we were following standard patterns.
	// Actually, the proto says UpdatePolicyRequest has a Policy, which includes name and etag.

	op, err := a.gcp.UpdatePolicy(ctx, req)
	if err != nil {
		return fmt.Errorf("updating IAMDenyPolicy: %w", err)
	}

	updated, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for IAMDenyPolicy update: %w", err)
	}

	status := &krm.IAMDenyPolicyStatus{}
	status.ExternalRef = direct.LazyPtr(updated.GetName())
	status.ObservedState = IAMDenyPolicyObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *denyPolicyAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("IAMDenyPolicy %q not found", a.fullyQualifiedName())
	}

	mapCtx := &direct.MapContext{}
	spec := IAMDenyPolicySpec_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	u, err := runtime.DefaultUnstructuredConverter.ToUnstructured(spec)
	if err != nil {
		return nil, fmt.Errorf("error converting IAMDenyPolicy spec to unstructured: %w", err)
	}

	return &unstructured.Unstructured{Object: u}, nil
}

func (a *denyPolicyAdapter) fullyQualifiedName() string {
	return fmt.Sprintf("policies/%s/denypolicies/%s", a.attachmentPoint, a.policyID)
}
