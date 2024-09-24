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
	"time"

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

type ContainerType int

const (
	Project ContainerType = iota
	Folder
	Organization
)

func (c ContainerType) String() string {
	return [...]string{"Project", "Folder", "Organization"}[c]
}

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

	projectRef, err := refs.ResolveProject(ctx, reader, obj, obj.Spec.ProjectRef)
	if err != nil {
		return nil, err
	}
	folderRef, err := refs.ResolveFolder(ctx, reader, obj, obj.Spec.FolderRef)
	if err != nil {
		return nil, err
	}
	organizationRef, err := refs.ResolveOrganization(ctx, reader, obj, obj.Spec.OrganizationRef)
	if err != nil {
		return nil, err
	}
	if projectRef == nil && folderRef == nil && organizationRef == nil {
		return nil, fmt.Errorf("one and only one of 'spec.projectRef', " +
			"'spec.folderRef' and 'spec.organizationRef' should be set, but " +
			"none of the three fields was set")
	}
	var containerType ContainerType
	errMsg := fmt.Sprintf("one and only one of 'spec.projectRef', "+
		"'spec.folderRef' and 'spec.organizationRef' should be set, but "+
		"got projectRef: %+v, folderRef: %+v, and organizationRef: %+v",
		obj.Spec.ProjectRef, obj.Spec.FolderRef, obj.Spec.OrganizationRef)
	if projectRef != nil {
		if folderRef != nil || organizationRef != nil {
			return nil, fmt.Errorf(errMsg)
		}
		containerType = Project
	} else if folderRef != nil {
		if organizationRef != nil {
			return nil, fmt.Errorf(errMsg)
		}
		containerType = Folder
	} else {
		containerType = Organization
	}

	projectID := ""
	if projectRef != nil {
		projectID = projectRef.ProjectID
		if projectID == "" {
			return nil, fmt.Errorf("cannot resolve project: project ID is empty")
		}
	}
	folderID := ""
	if folderRef != nil {
		folderID = folderRef.FolderID
		if folderID == "" {
			return nil, fmt.Errorf("cannot resolve folder: folder ID is empty")
		}
	}
	organizationID := ""
	if organizationRef != nil {
		organizationID = organizationRef.OrganizationID
		if organizationID == "" {
			return nil, fmt.Errorf("cannot resolve organization: organization ID is empty")
		}
	}
	// Get location
	location := *obj.Spec.Location

	var id *PrivilegedAccessManagerEntitlementIdentity
	externalRef := direct.ValueOf(obj.Status.ExternalRef)
	if externalRef == "" {
		id = BuildID(projectID, folderID, organizationID, location, resourceID)
	} else {
		id, err = asID(externalRef)
		if err != nil {
			return nil, err
		}

		var existingContainerType ContainerType
		if id.Parent.Project != "" {
			existingContainerType = Project
		} else if id.Parent.Folder != "" {
			existingContainerType = Folder
		} else if id.Parent.Organization != "" {
			existingContainerType = Organization
		} else {
			return nil, fmt.Errorf("status.externalRef doesn't have a parent")
		}
		if containerType != existingContainerType {
			return nil, fmt.Errorf("cannot change container type from %v to %v", existingContainerType.String(), containerType.String())
		}
		if id.Parent.Project != projectID {
			return nil, fmt.Errorf("PrivilegedAccessManagerEntitlement %s/%s has spec.projectRef changed, expect %s, got %s",
				u.GetNamespace(), u.GetName(), id.Parent.Project, projectID)
		}
		if id.Parent.Folder != folderID {
			return nil, fmt.Errorf("PrivilegedAccessManagerEntitlement %s/%s has spec.folderRef changed, expect %s, got %s",
				u.GetNamespace(), u.GetName(), id.Parent.Folder, folderID)
		}
		if id.Parent.Organization != organizationID {
			return nil, fmt.Errorf("PrivilegedAccessManagerEntitlement %s/%s has spec.organizationRef changed, expect %s, got %s",
				u.GetNamespace(), u.GetName(), id.Parent.Organization, organizationID)
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

	gcpIAMAccessProjectRef, err := refs.ResolveProject(ctx, reader, obj, obj.Spec.PrivilegedAccess.GcpIAMAccess.ProjectRef)
	if err != nil {
		return nil, err
	}
	gcpIAMAccessFolderRef, err := refs.ResolveFolder(ctx, reader, obj, obj.Spec.PrivilegedAccess.GcpIAMAccess.FolderRef)
	if err != nil {
		return nil, err
	}
	gcpIAMAccessOrganizationRef, err := refs.ResolveOrganization(ctx, reader, obj, obj.Spec.PrivilegedAccess.GcpIAMAccess.OrganizationRef)
	if err != nil {
		return nil, err
	}

	switch *obj.Spec.PrivilegedAccess.GcpIAMAccess.ResourceType {
	case "cloudresourcemanager.googleapis.com/Project":
		if gcpIAMAccessFolderRef != nil || gcpIAMAccessOrganizationRef != nil {
			return nil, fmt.Errorf("only 'spec.privilegedAccess.gcpIAMAccess.projectRef' " +
				"should be configured because the corresponding resourceType is " +
				"'cloudresourcemanager.googleapis.com/Project'")
		}
		if gcpIAMAccessProjectRef == nil || gcpIAMAccessProjectRef.ProjectID == "" {
			return nil, fmt.Errorf("'spec.privilegedAccess.gcpIAMAccess.projectRef' " +
				"should be non-empty because the corresponding resourceType is " +
				"'cloudresourcemanager.googleapis.com/Project'")
		}
		obj.Spec.PrivilegedAccess.GcpIAMAccess.ProjectRef.External = fmt.Sprintf("projects/%v", gcpIAMAccessProjectRef.ProjectID)
	case "cloudresourcemanager.googleapis.com/Folder":
		if gcpIAMAccessProjectRef != nil || gcpIAMAccessOrganizationRef != nil {
			return nil, fmt.Errorf("only 'spec.privilegedAccess.gcpIAMAccess.folderRef' " +
				"should be configured because the corresponding resourceType is " +
				"'cloudresourcemanager.googleapis.com/Folder'")
		}
		if gcpIAMAccessFolderRef == nil || gcpIAMAccessFolderRef.FolderID == "" {
			return nil, fmt.Errorf("'spec.privilegedAccess.gcpIAMAccess.folderRef' " +
				"should be non-empty because the corresponding resourceType is " +
				"'cloudresourcemanager.googleapis.com/Folder'")
		}
		obj.Spec.PrivilegedAccess.GcpIAMAccess.FolderRef.External = fmt.Sprintf("folders/%v", gcpIAMAccessFolderRef.FolderID)
	case "cloudresourcemanager.googleapis.com/Organization":
		if gcpIAMAccessProjectRef != nil || gcpIAMAccessFolderRef != nil {
			return nil, fmt.Errorf("only 'spec.privilegedAccess.gcpIAMAccess.organizationRef' " +
				"should be configured because the corresponding resourceType is " +
				"'cloudresourcemanager.googleapis.com/Organization'")
		}
		if gcpIAMAccessOrganizationRef == nil || gcpIAMAccessOrganizationRef.OrganizationID == "" {
			return nil, fmt.Errorf("'spec.privilegedAccess.gcpIAMAccess.organizationRef' " +
				"should be non-empty because the corresponding resourceType is " +
				"'cloudresourcemanager.googleapis.com/Organization'")
		}
		obj.Spec.PrivilegedAccess.GcpIAMAccess.OrganizationRef.External = fmt.Sprintf("organizations/%v", gcpIAMAccessOrganizationRef.OrganizationID)
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

	if !reflect.DeepEqual(AdditionalNotificationTargets_FromProto(mapCtx, a.actual.AdditionalNotificationTargets), a.desired.Spec.AdditionalNotificationTargets) {
		log.V(2).Info("'spec.additionalNotificationTargets' field is updated (-old +new)", cmp.Diff(AdditionalNotificationTargets_FromProto(mapCtx, a.actual.AdditionalNotificationTargets), a.desired.Spec.AdditionalNotificationTargets))
		updateMask.Paths = append(updateMask.Paths, "additional_notification_targets")
	}
	if !reflect.DeepEqual(ApprovalWorkflow_FromProto(mapCtx, a.actual.ApprovalWorkflow), a.desired.Spec.ApprovalWorkflow) {
		log.V(2).Info("'spec.approvalWorkflow' field is updated (-old +new)", cmp.Diff(ApprovalWorkflow_FromProto(mapCtx, a.actual.ApprovalWorkflow), a.desired.Spec.ApprovalWorkflow))
		updateMask.Paths = append(updateMask.Paths, "approval_workflow")
	}
	if !reflect.DeepEqual(direct.SliceOfPointers_FromProto(mapCtx, a.actual.EligibleUsers, AccessControlEntry_FromProto), a.desired.Spec.EligibleUsers) {
		log.V(2).Info("'spec.eligibleUsers' field is updated (-old +new)", cmp.Diff(direct.SliceOfPointers_FromProto(mapCtx, a.actual.EligibleUsers, AccessControlEntry_FromProto), a.desired.Spec.EligibleUsers))
		updateMask.Paths = append(updateMask.Paths, "eligible_users")
	}

	actualDuration := a.actual.MaxRequestDuration.AsDuration()
	desiredDuration, err := time.ParseDuration(*a.desired.Spec.MaxRequestDuration)
	if err != nil {
		return fmt.Errorf("error generating update mask: error parsing 'spec.maxRequestDuration': %w", err)
	}
	if !reflect.DeepEqual(actualDuration, desiredDuration) {
		log.V(2).Info("'spec.maxRequestDuration' field is updated (-old +new)", cmp.Diff(actualDuration, desiredDuration))
		updateMask.Paths = append(updateMask.Paths, "max_request_duration")
	}
	if !reflect.DeepEqual(PrivilegedAccess_FromProto(mapCtx, a.actual.PrivilegedAccess), a.desired.Spec.PrivilegedAccess) {
		log.V(2).Info("'spec.privilegedAccess' field is updated (-old +new)", cmp.Diff(PrivilegedAccess_FromProto(mapCtx, a.actual.PrivilegedAccess), a.desired.Spec.PrivilegedAccess))
		updateMask.Paths = append(updateMask.Paths, "privileged_access")
	}
	if !reflect.DeepEqual(RequesterJustificationConfig_FromProto(mapCtx, a.actual.RequesterJustificationConfig), a.desired.Spec.RequesterJustificationConfig) {
		log.V(2).Info("'spec.requesterJustificationConfig' field is updated (-old +new)", cmp.Diff(RequesterJustificationConfig_FromProto(mapCtx, a.actual.RequesterJustificationConfig), a.desired.Spec.RequesterJustificationConfig))
		updateMask.Paths = append(updateMask.Paths, "requester_justification_config")
	}
	if mapCtx.Err() != nil {
		return fmt.Errorf("error generating update mask: %w", mapCtx.Err())
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
	if a.id.Parent.Project != "" {
		obj.Spec.ProjectRef = &refs.ProjectRef{External: fmt.Sprintf("projects/%v", a.id.Parent.Project)}
	} else if a.id.Parent.Folder != "" {
		obj.Spec.FolderRef = &refs.FolderRef{External: fmt.Sprintf("folders/%v", a.id.Parent.Folder)}
	} else {
		obj.Spec.OrganizationRef = &refs.OrganizationRef{External: fmt.Sprintf("organizations/%v", a.id.Parent.Organization)}
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
