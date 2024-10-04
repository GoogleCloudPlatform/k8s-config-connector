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

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/privilegedaccessmanager/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	gcp "cloud.google.com/go/privilegedaccessmanager/apiv1"
	privilegedaccessmanagerpb "cloud.google.com/go/privilegedaccessmanager/apiv1/privilegedaccessmanagerpb"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
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

func (m *entitlementModel) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {

	obj := &krm.PrivilegedAccessManagerEntitlement{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	// Get ResourceID
	resourceID := direct.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	container, err := oneOfContainer(ctx, reader, obj,
		obj.Spec.ProjectRef,
		obj.Spec.FolderRef,
		obj.Spec.OrganizationRef)
	if err != nil {
		return nil, fmt.Errorf("error resolving 'obj.Spec.ProjectRef', "+
			"'obj.Spec.FolderRef' and 'obj.Spec.OrganizationRef': %w", err)
	}

	// Get location
	location := *obj.Spec.Location

	var id *PrivilegedAccessManagerEntitlementIdentity
	externalRef := direct.ValueOf(obj.Status.ExternalRef)
	if externalRef == "" {
		id = BuildID(container, location, resourceID)
	} else {
		id, err = asID(externalRef)
		if err != nil {
			return nil, err
		}

		if id.Parent.Container != container {
			return nil, fmt.Errorf("PrivilegedAccessManagerEntitlement %s/%s has parent container changed, expected %s, got %s",
				u.GetNamespace(), u.GetName(), id.Parent.Container, container)
		}
		if id.Parent.Location != location {
			return nil, fmt.Errorf("PrivilegedAccessManagerEntitlement %s/%s has spec.location changed, expect %s, got %s",
				u.GetNamespace(), u.GetName(), id.Parent.Location, location)
		}
		if id.Entitlement != resourceID {
			return nil, fmt.Errorf("PrivilegedAccessManagerEntitlement %s/%s has metadata.name or spec.resourceID changed, expect %s, got %s",
				u.GetNamespace(), u.GetName(), id.Entitlement, resourceID)
		}
	}

	iamAccessResource, err := oneOfContainer(ctx, reader, obj,
		obj.Spec.PrivilegedAccess.GcpIAMAccess.ProjectRef,
		obj.Spec.PrivilegedAccess.GcpIAMAccess.FolderRef,
		obj.Spec.PrivilegedAccess.GcpIAMAccess.OrganizationRef)
	if err != nil {
		return nil, fmt.Errorf("error resolving 'obj.Spec.PrivilegedAccess.GcpIAMAccess.ProjectRef', "+
			"'obj.Spec.PrivilegedAccess.GcpIAMAccess.FolderRef' and "+
			"'obj.Spec.PrivilegedAccess.GcpIAMAccess.OrganizationRef': %w", err)
	}
	switch *obj.Spec.PrivilegedAccess.GcpIAMAccess.ResourceType {
	case "cloudresourcemanager.googleapis.com/Project":
		if !strings.HasPrefix(iamAccessResource, "projects/") {
			return nil, fmt.Errorf("only 'spec.privilegedAccess.gcpIAMAccess.projectRef' "+
				"should be configured because the corresponding resourceType is "+
				"'cloudresourcemanager.googleapis.com/Project', but got resource %s", iamAccessResource)
		}
		obj.Spec.PrivilegedAccess.GcpIAMAccess.ProjectRef.External = iamAccessResource
	case "cloudresourcemanager.googleapis.com/Folder":
		if !strings.HasPrefix(iamAccessResource, "folders/") {
			return nil, fmt.Errorf("only 'spec.privilegedAccess.gcpIAMAccess.folderRef' "+
				"should be configured because the corresponding resourceType is "+
				"'cloudresourcemanager.googleapis.com/Folder', but got resource %s", iamAccessResource)
		}
		obj.Spec.PrivilegedAccess.GcpIAMAccess.FolderRef.External = iamAccessResource
	case "cloudresourcemanager.googleapis.com/Organization":
		if !strings.HasPrefix(iamAccessResource, "organizations/") {
			return nil, fmt.Errorf("only 'spec.privilegedAccess.gcpIAMAccess.organizationRef' "+
				"should be configured because the corresponding resourceType is "+
				"'cloudresourcemanager.googleapis.com/Organization', but got resource %s", iamAccessResource)
		}
		obj.Spec.PrivilegedAccess.GcpIAMAccess.OrganizationRef.External = iamAccessResource
	default:
		return nil, fmt.Errorf("unrecoganizable resourceType: %v; must be one of "+
			"'cloudresourcemanager.googleapis.com/Project', "+
			"'cloudresourcemanager.googleapis.com/Folder', "+
			"'cloudresourcemanager.googleapis.com/Organization'",
			*obj.Spec.PrivilegedAccess.GcpIAMAccess.ResourceType)
	}

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

func checkExactlyOneOf(values ...interface{}) (bool, interface{}) {
	numOfNonNil := 0
	var nonNilVal interface{}
	for _, value := range values {
		if value != nil && !reflect.ValueOf(value).IsNil() {
			numOfNonNil++
			nonNilVal = value
		}
	}
	if numOfNonNil != 1 {
		return false, nil
	}
	return true, nonNilVal
}

func oneOfContainer(ctx context.Context, reader client.Reader, obj *krm.PrivilegedAccessManagerEntitlement, projectRef *refs.ProjectRef, folderRef *refs.FolderRef, organizationRef *refs.OrganizationRef) (string, error) {
	hasExactlyOneContainer, containerRef := checkExactlyOneOf(projectRef, folderRef, organizationRef)
	if !hasExactlyOneContainer {
		return "", fmt.Errorf("exactly one of 'projectRef', 'folderRef' "+
			"or 'organizationRef' must be set, but got projectRef: %+v, folderRef: %+v, organizationRef: %+v",
			projectRef, folderRef, organizationRef)
	}

	container := ""
	switch containerRef.(type) {
	case *refs.ProjectRef:
		project, err := refs.ResolveProject(ctx, reader, obj, projectRef)
		if err != nil {
			return "", err
		}
		projectID := project.ProjectID
		if projectID == "" {
			return "", fmt.Errorf("cannot resolve project: project ID is empty")
		}
		container = fmt.Sprintf("projects/%s", projectID)
	case *refs.FolderRef:
		folder, err := refs.ResolveFolder(ctx, reader, obj, folderRef)
		if err != nil {
			return "", err
		}
		folderID := folder.FolderID
		if folderID == "" {
			return "", fmt.Errorf("cannot resolve folder: folder ID is empty")
		}
		container = fmt.Sprintf("folders/%s", folderID)
	case *refs.OrganizationRef:
		organization, err := refs.ResolveOrganization(ctx, reader, obj, organizationRef)
		if err != nil {
			return "", err
		}
		organizationID := organization.OrganizationID
		if organizationID == "" {
			return "", fmt.Errorf("cannot resolve organization: organization ID is empty")
		}
		container = fmt.Sprintf("organizations/%s", organizationID)
	default:
		return "", fmt.Errorf("unexpected ref type %T", containerRef)
	}

	return container, nil
}

func (m *entitlementModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type Adapter struct {
	id        *PrivilegedAccessManagerEntitlementIdentity
	gcpClient *gcp.Client
	desired   *krm.PrivilegedAccessManagerEntitlement
	actual    *privilegedaccessmanagerpb.Entitlement
}

var _ directbase.Adapter = &Adapter{}

func (a *Adapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx).WithName(ctrlName)
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

func (a *Adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	u := createOp.GetUnstructured()

	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("creating PrivilegedAccessManagerEntitlement", "name", a.id.FullyQualifiedName())
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := PrivilegedAccessManagerEntitlementSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &privilegedaccessmanagerpb.CreateEntitlementRequest{
		Parent:        a.id.Parent.String(),
		EntitlementId: a.id.Entitlement,
		Entitlement:   resource,
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
	observedState := PrivilegedAccessManagerEntitlementStatusObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ObservedState = observedState
	status.ExternalRef = a.id.AsExternalRef()
	return setStatus(u, status)
}

func (a *Adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	u := updateOp.GetUnstructured()

	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("updating PrivilegedAccessManagerEntitlement", "name", a.id.FullyQualifiedName())
	mapCtx := &direct.MapContext{}

	updateMask := &fieldmaskpb.FieldMask{}

	parsedActual := PrivilegedAccessManagerEntitlementSpec_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return fmt.Errorf("error generating update mask: %w", mapCtx.Err())
	}
	if !reflect.DeepEqual(parsedActual.AdditionalNotificationTargets, a.desired.Spec.AdditionalNotificationTargets) {
		log.V(2).Info("'spec.additionalNotificationTargets' field is updated (-old +new)", cmp.Diff(AdditionalNotificationTargets_FromProto(mapCtx, a.actual.AdditionalNotificationTargets), a.desired.Spec.AdditionalNotificationTargets))
		updateMask.Paths = append(updateMask.Paths, "additional_notification_targets")
	}
	if !reflect.DeepEqual(parsedActual.ApprovalWorkflow, a.desired.Spec.ApprovalWorkflow) {
		log.V(2).Info("'spec.approvalWorkflow' field is updated (-old +new)", cmp.Diff(ApprovalWorkflow_FromProto(mapCtx, a.actual.ApprovalWorkflow), a.desired.Spec.ApprovalWorkflow))
		updateMask.Paths = append(updateMask.Paths, "approval_workflow")
	}
	if !reflect.DeepEqual(parsedActual.EligibleUsers, a.desired.Spec.EligibleUsers) {
		log.V(2).Info("'spec.eligibleUsers' field is updated (-old +new)", cmp.Diff(direct.SliceOfPointers_FromProto(mapCtx, a.actual.EligibleUsers, AccessControlEntry_FromProto), a.desired.Spec.EligibleUsers))
		updateMask.Paths = append(updateMask.Paths, "eligible_users")
	}
	if !reflect.DeepEqual(a.actual.MaxRequestDuration.AsDuration(), direct.StringDuration_ToProto(mapCtx, a.desired.Spec.MaxRequestDuration).AsDuration()) {
		log.V(2).Info("'spec.maxRequestDuration' field is updated (-old +new)", cmp.Diff(a.actual.MaxRequestDuration.AsDuration(), direct.StringDuration_ToProto(mapCtx, a.desired.Spec.MaxRequestDuration).AsDuration()))
		updateMask.Paths = append(updateMask.Paths, "max_request_duration")
	}
	if !reflect.DeepEqual(parsedActual.PrivilegedAccess, a.desired.Spec.PrivilegedAccess) {
		log.V(2).Info("'spec.privilegedAccess' field is updated (-old +new)", cmp.Diff(PrivilegedAccess_FromProto(mapCtx, a.actual.PrivilegedAccess), a.desired.Spec.PrivilegedAccess))
		updateMask.Paths = append(updateMask.Paths, "privileged_access")
	}
	if !reflect.DeepEqual(parsedActual.RequesterJustificationConfig, a.desired.Spec.RequesterJustificationConfig) {
		log.V(2).Info("'spec.requesterJustificationConfig' field is updated (-old +new)", cmp.Diff(RequesterJustificationConfig_FromProto(mapCtx, a.actual.RequesterJustificationConfig), a.desired.Spec.RequesterJustificationConfig))
		updateMask.Paths = append(updateMask.Paths, "requester_justification_config")
	}
	if len(updateMask.Paths) == 0 {
		log.V(2).Info("underlying PrivilegedAccessManagerEntitlement already up to date", "name", a.id.FullyQualifiedName())

		status := &krm.PrivilegedAccessManagerEntitlementStatus{}
		observedState := PrivilegedAccessManagerEntitlementStatusObservedState_FromProto(mapCtx, a.actual)
		if mapCtx.Err() != nil {
			return mapCtx.Err()
		}
		status.ObservedState = observedState
		return setStatus(u, status)
	}
	desired := a.desired.DeepCopy()
	resource := PrivilegedAccessManagerEntitlementSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	resource.Name = a.id.FullyQualifiedName()
	resource.Etag = a.actual.Etag
	req := &privilegedaccessmanagerpb.UpdateEntitlementRequest{
		UpdateMask:  updateMask,
		Entitlement: resource,
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
	observedState := PrivilegedAccessManagerEntitlementStatusObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ObservedState = observedState
	return setStatus(u, status)
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
	if strings.HasPrefix(a.id.Parent.Container, "projects") {
		obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Parent.Container}
	} else if strings.HasPrefix(a.id.Parent.Container, "folders") {
		obj.Spec.FolderRef = &refs.FolderRef{External: a.id.Parent.Container}
	} else {
		obj.Spec.OrganizationRef = &refs.OrganizationRef{External: a.id.Parent.Container}
	}

	obj.Spec.Location = &a.id.Parent.Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}
	u.Object = uObj
	return u, nil
}

// Delete implements the Adapter interface.
func (a *Adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("deleting PrivilegedAccessManagerEntitlement", "name", a.id.FullyQualifiedName())

	req := &privilegedaccessmanagerpb.DeleteEntitlementRequest{Name: a.id.FullyQualifiedName()}
	op, err := a.gcpClient.DeleteEntitlement(ctx, req)
	if err != nil {
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
