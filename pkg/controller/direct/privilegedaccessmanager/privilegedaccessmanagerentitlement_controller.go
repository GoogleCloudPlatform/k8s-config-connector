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

package privilegedaccessmanager

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/privilegedaccessmanager/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/privilegedaccessmanager/apiv1"
	privilegedaccessmanagerpb "cloud.google.com/go/privilegedaccessmanager/apiv1/privilegedaccessmanagerpb"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/api/option"
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
	return &Adapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *entitlementModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type Adapter struct {
	id        *krm.PrivilegedAccessManagerEntitlementIdentity
	gcpClient *gcp.Client
	desired   *krm.PrivilegedAccessManagerEntitlement
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
	u := createOp.GetUnstructured()

	log := klog.FromContext(ctx)
	log.V(2).Info("creating PrivilegedAccessManagerEntitlement", "name", a.id.FullyQualifiedName())
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resourceType, resource, err := getResourceTypeAndResourceFromContainer(a.id.Container())
	if err != nil {
		return fmt.Errorf("error getting resourceType and resource from container: %w", err)
	}
	hiddenFields := gcpIAMAccessResource{resourceType: resourceType, resource: resource}
	entitlement := PrivilegedAccessManagerEntitlementSpec_ToProto(mapCtx, &desired.Spec, hiddenFields)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &privilegedaccessmanagerpb.CreateEntitlementRequest{
		Parent:        a.id.ParentString(),
		EntitlementId: a.id.Entitlement,
		Entitlement:   entitlement,
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

	status := &krm.PrivilegedAccessManagerEntitlementStatus{}
	observedState := PrivilegedAccessManagerEntitlementObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ObservedState = observedState
	status.ExternalRef = a.id.AsExternalRef()
	return setStatus(u, status)
}

func (a *Adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	u := updateOp.GetUnstructured()

	log := klog.FromContext(ctx)
	log.V(2).Info("updating PrivilegedAccessManagerEntitlement", "name", a.id.FullyQualifiedName())
	mapCtx := &direct.MapContext{}

	updateMask := &fieldmaskpb.FieldMask{}

	report := &structuredreporting.Diff{Object: updateOp.GetUnstructured()}

	parsedActual := PrivilegedAccessManagerEntitlementSpec_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return fmt.Errorf("error generating update mask: %w", mapCtx.Err())
	}
	sortPrincipalsInSpec(parsedActual)
	parsedDesired := a.desired.DeepCopy()
	sortPrincipalsInSpec(&parsedDesired.Spec)

	if !reflect.DeepEqual(parsedActual.AdditionalNotificationTargets, parsedDesired.Spec.AdditionalNotificationTargets) {
		report.AddField("additional_notification_targets", parsedActual.AdditionalNotificationTargets, parsedDesired.Spec.AdditionalNotificationTargets)
		log.V(2).Info("'spec.additionalNotificationTargets' field is updated (-old +new)", cmp.Diff(parsedActual.AdditionalNotificationTargets, parsedDesired.Spec.AdditionalNotificationTargets))
		updateMask.Paths = append(updateMask.Paths, "additional_notification_targets")
	}

	// If '.spec.approvalWorkflow.manualApprovals.requireApproverJustification'
	// is unset in the desired state, the corresponding field will be 'false' in
	// the returned actual state when 'approvalWorkflow.manualApprovals' is not
	// an empty struct. So the diffing needs to be handled differently.
	if !(isApprovalWorkflowManualApprovalsRequireApproverJustificationUnset(&parsedDesired.Spec) &&
		isApprovalWorkflowManualApprovalsRequireApproverJustificationSetToFalse(parsedActual)) {
		if !reflect.DeepEqual(parsedActual.ApprovalWorkflow, parsedDesired.Spec.ApprovalWorkflow) {
			report.AddField("approval_workflow", parsedActual.ApprovalWorkflow, parsedDesired.Spec.ApprovalWorkflow)
			log.V(2).Info("'spec.approvalWorkflow' field is updated (-old +new)", cmp.Diff(parsedActual.ApprovalWorkflow, parsedDesired.Spec.ApprovalWorkflow))
			updateMask.Paths = append(updateMask.Paths, "approval_workflow")
		}
	}
	if !reflect.DeepEqual(parsedActual.EligibleUsers, parsedDesired.Spec.EligibleUsers) {
		report.AddField("eligible_users", parsedActual.EligibleUsers, parsedDesired.Spec.EligibleUsers)
		log.V(2).Info("'spec.eligibleUsers' field is updated (-old +new)", cmp.Diff(parsedActual.EligibleUsers, parsedDesired.Spec.EligibleUsers))
		updateMask.Paths = append(updateMask.Paths, "eligible_users")
	}
	if !reflect.DeepEqual(a.actual.MaxRequestDuration.AsDuration(), direct.StringDuration_ToProto(mapCtx, parsedDesired.Spec.MaxRequestDuration).AsDuration()) {
		report.AddField("max_request_duration.as_duration", a.actual.MaxRequestDuration.AsDuration(), direct.StringDuration_ToProto(mapCtx, parsedDesired.Spec.MaxRequestDuration).AsDuration())
		log.V(2).Info("'spec.maxRequestDuration' field is updated (-old +new)", cmp.Diff(a.actual.MaxRequestDuration.AsDuration(), direct.StringDuration_ToProto(mapCtx, parsedDesired.Spec.MaxRequestDuration).AsDuration()))
		updateMask.Paths = append(updateMask.Paths, "max_request_duration")
	}
	if !reflect.DeepEqual(parsedActual.PrivilegedAccess, parsedDesired.Spec.PrivilegedAccess) {
		report.AddField("privileged_access", parsedActual.PrivilegedAccess, parsedDesired.Spec.PrivilegedAccess)
		log.V(2).Info("'spec.privilegedAccess' field is updated (-old +new)", cmp.Diff(parsedActual.PrivilegedAccess, parsedDesired.Spec.PrivilegedAccess))
		updateMask.Paths = append(updateMask.Paths, "privileged_access")
	}
	if !reflect.DeepEqual(parsedActual.RequesterJustificationConfig, parsedDesired.Spec.RequesterJustificationConfig) {
		report.AddField("requester_justification_config", parsedActual.RequesterJustificationConfig, parsedDesired.Spec.RequesterJustificationConfig)
		log.V(2).Info("'spec.requesterJustificationConfig' field is updated (-old +new)", cmp.Diff(parsedActual.RequesterJustificationConfig, parsedDesired.Spec.RequesterJustificationConfig))
		updateMask.Paths = append(updateMask.Paths, "requester_justification_config")
	}
	if len(updateMask.Paths) == 0 {
		log.V(2).Info("underlying PrivilegedAccessManagerEntitlement already up to date", "name", a.id.FullyQualifiedName())

		status := &krm.PrivilegedAccessManagerEntitlementStatus{}
		observedState := PrivilegedAccessManagerEntitlementObservedState_FromProto(mapCtx, a.actual)
		if mapCtx.Err() != nil {
			return mapCtx.Err()
		}
		status.ObservedState = observedState
		status.ExternalRef = a.id.AsExternalRef()
		return setStatus(u, status)
	}

	structuredreporting.ReportDiff(ctx, report)

	desired := a.desired.DeepCopy()
	resourceType, resource, err := getResourceTypeAndResourceFromContainer(a.id.Container())
	if err != nil {
		return fmt.Errorf("error getting resourceType and resource from container: %w", err)
	}
	hiddenFields := gcpIAMAccessResource{resourceType: resourceType, resource: resource}
	entitlement := PrivilegedAccessManagerEntitlementSpec_ToProto(mapCtx, &desired.Spec, hiddenFields)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	entitlement.Name = a.id.FullyQualifiedName()
	entitlement.Etag = a.actual.Etag
	req := &privilegedaccessmanagerpb.UpdateEntitlementRequest{
		UpdateMask:  updateMask,
		Entitlement: entitlement,
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

	status := &krm.PrivilegedAccessManagerEntitlementStatus{}
	observedState := PrivilegedAccessManagerEntitlementObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ObservedState = observedState
	return setStatus(u, status)
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

func setStatus(u *unstructured.Unstructured, typedStatus any) error {
	status, err := runtime.DefaultUnstructuredConverter.ToUnstructured(typedStatus)
	if err != nil {
		return fmt.Errorf("error converting status to unstructured: %w", err)
	}

	old, _, _ := unstructured.NestedMap(u.Object, "status")
	if old != nil {
		status["conditions"] = old["conditions"]
		status["observedGeneration"] = old["observedGeneration"]
		status["externalRef"] = old["externalRef"]
	}

	u.Object["status"] = status

	return nil
}
