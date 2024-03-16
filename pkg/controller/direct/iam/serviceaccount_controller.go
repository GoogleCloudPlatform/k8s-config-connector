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

package iam

import (
	"context"
	"fmt"
	"strings"

	api "google.golang.org/api/iam/v1"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/manager"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/iam/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
)

// AddServiceAccountController creates a new controller and adds it to the Manager.
// The Manager will set fields on the Controller and start it when the Manager is started.
func AddServiceAccountController(mgr manager.Manager, config *controller.Config, opts directbase.Deps) error {
	gvk := krm.IAMServiceAccountGVK

	// TODO: Share gcp client (any value in doing so)?
	ctx := context.TODO()
	gcpClient, err := newGCPClient(ctx, config)
	if err != nil {
		return err
	}
	m := &serviceAccountModel{gcpClient: gcpClient}
	return directbase.Add(mgr, gvk, m, opts)
}

type serviceAccountModel struct {
	*gcpClient
}

// model implements the Model interface.
var _ directbase.Model = &serviceAccountModel{}

type serviceAccountAdapter struct {
	projectID string

	// The account id that is used to generate the service account
	// email address and a stable unique id. It is unique within a project,
	// must be 6-30 characters long, and match the regular expression
	// `[a-z]([-a-z0-9]*[a-z0-9])` to comply with RFC1035.
	accountID string

	desiredKRM *krm.IAMServiceAccountSpec
	desired    *api.ServiceAccount
	actual     *api.ServiceAccount

	iamClient *api.Service
}

// adapter implements the Adapter interface.
var _ directbase.Adapter = &serviceAccountAdapter{}

// AdapterForObject implements the Model interface.
func (m *serviceAccountModel) AdapterForObject(ctx context.Context, u *unstructured.Unstructured) (directbase.Adapter, error) {
	iamClient, err := m.newIamClient(ctx)
	if err != nil {
		return nil, err
	}

	// TODO: Just fetch this object?
	obj := &krm.IAMServiceAccount{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	projectID := obj.Annotations["cnrm.cloud.google.com/project-id"]
	if projectID == "" {
		return nil, fmt.Errorf("unable to determine project")
	}

	resourceID := ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("unable to determine resourceID")
	}

	mapContext := &MapContext{}
	desired := ServiceAccountSpec_ToProto(mapContext, &obj.Spec)
	if err := mapContext.Err(); err != nil {
		return nil, mapContext.Err()
	}

	return &serviceAccountAdapter{
		projectID:  projectID,
		accountID:  resourceID,
		desiredKRM: &obj.Spec,
		desired:    desired,
		iamClient:  iamClient,
	}, nil
}

// Find implements the Adapter interface.
func (a *serviceAccountAdapter) Find(ctx context.Context) (bool, error) {
	if a.accountID == "" {
		return false, nil
	}

	serviceAccount, err := a.iamClient.Projects.ServiceAccounts.Get(a.fullyQualifiedName()).Context(ctx).Do()
	if err != nil {
		if IsNotFound(err) {
			klog.Warningf("serviceaccount was not found: %v", err)
			return false, nil
		}
		return false, fmt.Errorf("getting serviceAccount %q: %w", a.fullyQualifiedName(), err)
	}
	a.actual = serviceAccount

	return true, nil
}

func setStatus(u *unstructured.Unstructured, typedStatus any) error {
	status, err := runtime.DefaultUnstructuredConverter.ToUnstructured(typedStatus)
	if err != nil {
		return fmt.Errorf("error converting status to unstructured: %w", err)
	}
	oldStatus := u.Object["status"]

	// Preserve "other" fields
	if oldStatus != nil {
		oldStatusMap := oldStatus.(map[string]any)
		status["conditions"] = oldStatusMap["conditions"]
		status["observedGeneration"] = oldStatusMap["observedGeneration"]
	}

	u.Object["status"] = status

	return nil
}

func setResourceID(u *unstructured.Unstructured, resourceID string) error {
	if err := unstructured.SetNestedField(u.Object, resourceID, "spec", "resourceID"); err != nil {
		return fmt.Errorf("setting spec.resourceID=%q: %w", resourceID, err)
	}

	return nil
}

// Delete implements the Adapter interface.
func (a *serviceAccountAdapter) Delete(ctx context.Context) (bool, error) {
	// TODO: Delete via status selfLink?
	_, err := a.iamClient.Projects.ServiceAccounts.Delete(a.fullyQualifiedName()).Context(ctx).Do()
	if err != nil {
		if IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting serviceaccount: %w", err)
	}

	return true, nil
}

// Create implements the Adapter interface.
func (a *serviceAccountAdapter) Create(ctx context.Context, u *unstructured.Unstructured) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating object", "u", u)

	serviceAccount := a.desired

	req := &api.CreateServiceAccountRequest{
		AccountId:      a.accountID,
		ServiceAccount: serviceAccount,
	}

	parent := "projects/" + a.projectID
	created, err := a.iamClient.Projects.ServiceAccounts.Create(parent, req).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("creating serviceaccount: %w", err)
	}
	log.V(2).Info("created serviceaccount", "serviceaccount", created)

	latest, err := a.applyDisabled(ctx, created)
	if err != nil {
		return err
	}

	if err := setResourceID(u, a.accountID); err != nil {
		return err
	}

	mapContext := &MapContext{}
	status := ServiceAccountStatus_FromProto(mapContext, latest)
	if err := ctx.Err(); err != nil {
		return err
	}

	return setStatus(u, status)
}

// Update implements the Adapter interface.
func (a *serviceAccountAdapter) Update(ctx context.Context, u *unstructured.Unstructured) error {
	// TODO: Skip updates at the higher level if no changes?
	updateMask := &fieldmaskpb.FieldMask{}

	mapContext := &MapContext{}

	latest := a.actual

	latest, err := a.applyEnabled(ctx, latest)
	if err != nil {
		return err
	}

	actualKRM := ServiceAccountSpec_FromProto(mapContext, latest)
	if err := mapContext.Err(); err != nil {
		return err
	}

	// TODO: I think we can do this with a helper
	if a.desiredKRM.DisplayName != nil && ValueOf(a.desiredKRM.DisplayName) != ValueOf(actualKRM.DisplayName) {
		// TODO: Terraform passes display_name; gcloud passes displayName
		updateMask.Paths = append(updateMask.Paths, "display_name")
	}
	if a.desiredKRM.Description != nil && ValueOf(a.desiredKRM.Description) != ValueOf(actualKRM.Description) {
		updateMask.Paths = append(updateMask.Paths, "description")
	}

	if len(updateMask.Paths) != 0 {
		// TODO: Where/how do we want to enforce immutability?

		sa := &api.ServiceAccount{
			DisplayName: a.desired.DisplayName,
			Description: a.desired.Description,
			// ForceSendFields: []string{"DisplayName", "Description"},
		}
		req := &api.PatchServiceAccountRequest{
			UpdateMask:     strings.Join(updateMask.GetPaths(), ","),
			ServiceAccount: sa,
		}

		_, err := a.iamClient.Projects.ServiceAccounts.Patch(a.fullyQualifiedName(), req).Context(ctx).Do()
		if err != nil {
			return err
		}

		// Patch doesn't return the full object
		updated, err := a.iamClient.Projects.ServiceAccounts.Get(a.fullyQualifiedName()).Context(ctx).Do()
		if err != nil {
			return fmt.Errorf("getting serviceAccount %q: %w", a.fullyQualifiedName(), err)
		}
		latest = updated
	} else {
		latest = a.actual
	}

	latest, err = a.applyDisabled(ctx, latest)
	if err != nil {
		return err
	}

	status := ServiceAccountStatus_FromProto(mapContext, latest)
	if err := ctx.Err(); err != nil {
		return err
	}

	return setStatus(u, status)
}

func (a *serviceAccountAdapter) applyEnabled(ctx context.Context, latest *api.ServiceAccount) (*api.ServiceAccount, error) {
	if a.desiredKRM.Disabled == nil || ValueOf(a.desiredKRM.Disabled) == latest.Disabled {
		return latest, nil
	}

	if !ValueOf(a.desiredKRM.Disabled) {
		req := &api.EnableServiceAccountRequest{}
		if _, err := a.iamClient.Projects.ServiceAccounts.Enable(a.fullyQualifiedName(), req).Context(ctx).Do(); err != nil {
			return nil, err
		}
	}

	updated, err := a.iamClient.Projects.ServiceAccounts.Get(a.fullyQualifiedName()).Context(ctx).Do()
	if err != nil {
		return nil, fmt.Errorf("getting serviceAccount %q: %w", a.fullyQualifiedName(), err)
	}
	return updated, nil
}

func (a *serviceAccountAdapter) applyDisabled(ctx context.Context, latest *api.ServiceAccount) (*api.ServiceAccount, error) {
	if a.desiredKRM.Disabled == nil || ValueOf(a.desiredKRM.Disabled) == latest.Disabled {
		return latest, nil
	}

	if ValueOf(a.desiredKRM.Disabled) {
		req := &api.DisableServiceAccountRequest{}
		if _, err := a.iamClient.Projects.ServiceAccounts.Disable(a.fullyQualifiedName(), req).Context(ctx).Do(); err != nil {
			return nil, err
		}
	}

	updated, err := a.iamClient.Projects.ServiceAccounts.Get(a.fullyQualifiedName()).Context(ctx).Do()
	if err != nil {
		return nil, fmt.Errorf("getting serviceAccount %q: %w", a.fullyQualifiedName(), err)
	}
	return updated, nil
}

func (a *serviceAccountAdapter) fullyQualifiedName() string {
	email := a.accountID + "@" + a.projectID + ".iam.gserviceaccount.com"
	return fmt.Sprintf("projects/%s/serviceAccounts/%s", a.projectID, email)
}
