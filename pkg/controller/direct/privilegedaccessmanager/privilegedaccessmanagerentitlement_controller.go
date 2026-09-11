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

package privilegedaccessmanager

import (
	"context"
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/privilegedaccessmanager/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/tags"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/privilegedaccessmanager/apiv1"
	privilegedaccessmanagerpb "cloud.google.com/go/privilegedaccessmanager/apiv1/privilegedaccessmanagerpb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

const (
	ctrlName      = "privilegedaccessmanager-controller"
	serviceDomain = "//privilegedaccessmanager.googleapis.com"
)

func init() {
	registry.RegisterModel(krm.PrivilegedAccessManagerEntitlementGVK, NewModel)
}

func NewModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &entitlementModel{config: *config}, nil
}

var _ directbase.Model = &entitlementModel{}

type entitlementModel struct {
	config config.ControllerConfig
}

func (m *entitlementModel) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building PrivilegedAccessManager client for Entitlement: %w", err)
	}
	return gcpClient, err
}

func (m *entitlementModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader

	obj := &krm.PrivilegedAccessManagerEntitlement{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	identity, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := identity.(*krm.PrivilegedAccessManagerEntitlementIdentity)

	if obj.Spec.RequesterJustificationConfig.NotMandatory == nil && obj.Spec.RequesterJustificationConfig.Unstructured == nil {
		return nil, fmt.Errorf("one and only one of 'spec.requesterJustificationConfig.notMandatory' " +
			"and 'spec.requesterJustificationConfig.unstructured' should be configured: neither is configured")
	}
	if obj.Spec.RequesterJustificationConfig.NotMandatory != nil && obj.Spec.RequesterJustificationConfig.Unstructured != nil {
		return nil, fmt.Errorf("one and only one of 'spec.requesterJustificationConfig.notMandatory' " +
			"and 'spec.requesterJustificationConfig.unstructured' should be configured: both configured")
	}

	// Get privilegedaccessmanager GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	resourceType, resource, err := getResourceTypeAndResourceFromContainer(id.Container())
	if err != nil {
		return nil, fmt.Errorf("error getting resourceType and resource from container: %w", err)
	}
	hiddenFields := gcpIAMAccessResource{resourceType: resourceType, resource: resource}

	mapCtx := &direct.MapContext{}
	desiredPb := PrivilegedAccessManagerEntitlementSpec_ToProto(mapCtx, &obj.Spec, hiddenFields)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	return &Adapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   desiredPb,
	}, nil
}

func (m *entitlementModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type Adapter struct {
	id        *krm.PrivilegedAccessManagerEntitlementIdentity
	gcpClient *gcp.Client
	desired   *privilegedaccessmanagerpb.Entitlement
	actual    *privilegedaccessmanagerpb.Entitlement
}

var _ directbase.Adapter = &Adapter{}

func (a *Adapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting PrivilegedAccessManagerEntitlement", "name", a.id.FullyQualifiedName())

	req := &privilegedaccessmanagerpb.GetEntitlementRequest{Name: a.id.FullyQualifiedName()}
	entitlementpb, err := a.gcpClient.GetEntitlement(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("error getting PrivilegedAccessManagerEntitlement %q: %w", a.id.FullyQualifiedName(), err)
	}

	a.actual = entitlementpb
	return true, nil
}

func getResourceTypeAndResourceFromContainer(container string) (string, string, error) {
	tokens := strings.Split(container, "/")
	if len(tokens) != 2 {
		return "", "", fmt.Errorf("container should be one of projects/<project>, "+
			"folders/<folder> or organizations/<organization>, but got %s", container)
	}
	resource := fmt.Sprintf("//cloudresourcemanager.googleapis.com/%s", container)
	switch tokens[0] {
	case "projects":
		return "cloudresourcemanager.googleapis.com/Project", resource, nil
	case "folders":
		return "cloudresourcemanager.googleapis.com/Folder", resource, nil
	case "organizations":
		return "cloudresourcemanager.googleapis.com/Organization", resource, nil
	default:
		return "", "", fmt.Errorf("container must start with 'projects', "+
			"'folders', or 'organizations', but it starts with %v", tokens[0])
	}
}

func (a *Adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating PrivilegedAccessManagerEntitlement", "name", a.id.FullyQualifiedName())

	req := &privilegedaccessmanagerpb.CreateEntitlementRequest{
		Parent:        a.id.ParentString(),
		EntitlementId: a.id.Entitlement,
		Entitlement:   a.desired,
	}
	op, err := a.gcpClient.CreateEntitlement(ctx, req)
	if err != nil {
		return fmt.Errorf("error creating PrivilegedAccessManagerEntitlement %s: %w", a.id.FullyQualifiedName(), err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("error waiting for PrivilegedAccessManagerEntitlement %s to be created: %w", a.id.FullyQualifiedName(), err)
	}
	log.V(2).Info("successfully created PrivilegedAccessManagerEntitlement", "name", a.id.FullyQualifiedName())

	return a.updateStatus(ctx, createOp, created)
}

func (a *Adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating PrivilegedAccessManagerEntitlement", "name", a.id.FullyQualifiedName())

	diffs, updateMask, err := comparePrivilegedAccessManagerEntitlement(ctx, a.actual, a.desired, a.id)
	if err != nil {
		return err
	}

	latest := a.actual
	if diffs.HasDiff() {
		diffs.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, diffs)

		desired := proto.Clone(a.desired).(*privilegedaccessmanagerpb.Entitlement)
		desired.Name = a.id.FullyQualifiedName()
		desired.Etag = a.actual.Etag

		req := &privilegedaccessmanagerpb.UpdateEntitlementRequest{
			UpdateMask:  updateMask,
			Entitlement: desired,
		}
		op, err := a.gcpClient.UpdateEntitlement(ctx, req)
		if err != nil {
			return fmt.Errorf("error updating PrivilegedAccessManagerEntitlement %s: %w", a.id.FullyQualifiedName(), err)
		}
		updated, err := op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("error waiting for PrivilegedAccessManagerEntitlement %s to be updated: %w", a.id.FullyQualifiedName(), err)
		}
		log.V(2).Info("successfully updated PrivilegedAccessManagerEntitlement", "name", a.id.FullyQualifiedName())
		latest = updated
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func (a *Adapter) updateStatus(ctx context.Context, op directbase.Operation, latest *privilegedaccessmanagerpb.Entitlement) error {
	mapCtx := &direct.MapContext{}
	status := &krm.PrivilegedAccessManagerEntitlementStatus{}
	observedState := PrivilegedAccessManagerEntitlementObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ObservedState = observedState
	status.ExternalRef = a.id.AsExternalRef()
	return op.UpdateStatus(ctx, status, nil)
}

func comparePrivilegedAccessManagerEntitlement(ctx context.Context, actual, desired *privilegedaccessmanagerpb.Entitlement, id *krm.PrivilegedAccessManagerEntitlementIdentity) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	resourceType, resource, err := getResourceTypeAndResourceFromContainer(id.Container())
	if err != nil {
		return nil, nil, fmt.Errorf("error getting resourceType and resource from container: %w", err)
	}
	hiddenFields := gcpIAMAccessResource{resourceType: resourceType, resource: resource}

	parsedActual := PrivilegedAccessManagerEntitlementSpec_FromProto(&direct.MapContext{}, actual)
	parsedDesired := PrivilegedAccessManagerEntitlementSpec_FromProto(&direct.MapContext{}, desired)

	sortPrincipalsInSpec(parsedActual)
	sortPrincipalsInSpec(parsedDesired)

	// Normalize requireApproverJustification:
	// If it is unset in desired, but false in actual, set it to false in desired
	// so we don't have a false positive diff.
	if isApprovalWorkflowManualApprovalsRequireApproverJustificationUnset(parsedDesired) &&
		isApprovalWorkflowManualApprovalsRequireApproverJustificationSetToFalse(parsedActual) {
		if parsedDesired.ApprovalWorkflow != nil && parsedDesired.ApprovalWorkflow.ManualApprovals != nil {
			f := false
			parsedDesired.ApprovalWorkflow.ManualApprovals.RequireApproverJustification = &f
		}
	}

	mapCtx := &direct.MapContext{}
	maskedActual := PrivilegedAccessManagerEntitlementSpec_ToProto(mapCtx, parsedActual, hiddenFields)
	if mapCtx.Err() != nil {
		return nil, nil, mapCtx.Err()
	}
	clonedDesired := PrivilegedAccessManagerEntitlementSpec_ToProto(mapCtx, parsedDesired, hiddenFields)
	if mapCtx.Err() != nil {
		return nil, nil, mapCtx.Err()
	}

	maskedActual.Name = actual.Name
	clonedDesired.Name = actual.Name

	maskedActual.Etag = actual.Etag
	clonedDesired.Etag = actual.Etag

	diffs, updateMask, err := tags.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}

func isApprovalWorkflowManualApprovalsRequireApproverJustificationUnset(spec *krm.PrivilegedAccessManagerEntitlementSpec) bool {
	return spec.ApprovalWorkflow == nil ||
		spec.ApprovalWorkflow.ManualApprovals == nil ||
		spec.ApprovalWorkflow.ManualApprovals.RequireApproverJustification == nil
}

func isApprovalWorkflowManualApprovalsRequireApproverJustificationSetToFalse(spec *krm.PrivilegedAccessManagerEntitlementSpec) bool {
	return spec.ApprovalWorkflow != nil &&
		spec.ApprovalWorkflow.ManualApprovals != nil &&
		spec.ApprovalWorkflow.ManualApprovals.RequireApproverJustification != nil &&
		*spec.ApprovalWorkflow.ManualApprovals.RequireApproverJustification == false
}

func (a *Adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.PrivilegedAccessManagerEntitlement{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(PrivilegedAccessManagerEntitlementSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	if strings.HasPrefix(a.id.Container(), "projects/") {
		obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Container()}
	} else if strings.HasPrefix(a.id.Container(), "folders/") {
		obj.Spec.FolderRef = &refs.FolderRef{External: a.id.Container()}
	} else if strings.HasPrefix(a.id.Container(), "organizations/") {
		obj.Spec.OrganizationRef = &refs.OrganizationRef{External: a.id.Container()}
	} else {
		return nil, fmt.Errorf("invalid container format for %q", a.id.Container())
	}

	obj.Spec.Location = &a.id.Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}
	u.Object = uObj
	return u, nil
}

// Delete implements the Adapter interface.
func (a *Adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting PrivilegedAccessManagerEntitlement", "name", a.id.FullyQualifiedName())

	req := &privilegedaccessmanagerpb.DeleteEntitlementRequest{Name: a.id.FullyQualifiedName()}
	op, err := a.gcpClient.DeleteEntitlement(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("error deleting PrivilegedAccessManagerEntitlement %s: %w", a.id.FullyQualifiedName(), err)
	}

	_, err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("error waiting for PrivilegedAccessManagerEntitlement %s to be deleted: %w", a.id.FullyQualifiedName(), err)
	}

	log.V(2).Info("successfully deleted PrivilegedAccessManagerEntitlement", "name", a.id.FullyQualifiedName())
	return true, nil
}
