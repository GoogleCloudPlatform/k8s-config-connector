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
	"google.golang.org/api/option"
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
func AddServiceAccountController(mgr manager.Manager, config *controller.Config) error {
	gvk := krm.IAMServiceAccountGVK

	return directbase.Add(mgr, gvk, &model{config: *config})
}

type model struct {
	config controller.Config
}

// model implements the Model interface.
var _ directbase.Model = &model{}

type adapter struct {
	projectID string

	// The account id that is used to generate the service account
	// email address and a stable unique id. It is unique within a project,
	// must be 6-30 characters long, and match the regular expression
	// `[a-z]([-a-z0-9]*[a-z0-9])` to comply with RFC1035.
	accountID string

	desired *krm.IAMServiceAccount
	actual  *krm.IAMServiceAccount

	gcp *api.Service
}

// adapter implements the Adapter interface.
var _ directbase.Adapter = &adapter{}

func (m *model) client(ctx context.Context) (*api.Service, error) {
	var opts []option.ClientOption
	if m.config.UserAgent != "" {
		opts = append(opts, option.WithUserAgent(m.config.UserAgent))
	}
	if m.config.HTTPClient != nil {
		// TODO: Set UserAgent in this scenario (error is: WithHTTPClient is incompatible with gRPC dial options)
		opts = append(opts, option.WithHTTPClient(m.config.HTTPClient))
	}
	if m.config.UserProjectOverride && m.config.BillingProject != "" {
		opts = append(opts, option.WithQuotaProject(m.config.BillingProject))
	}

	// TODO: support endpoints?
	// if m.config.Endpoint != "" {
	// 	opts = append(opts, option.WithEndpoint(m.config.Endpoint))
	// }

	gcpClient, err := api.NewService(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building iam client: %w", err)
	}
	return gcpClient, err
}

// AdapterForObject implements the Model interface.
func (m *model) AdapterForObject(ctx context.Context, u *unstructured.Unstructured) (directbase.Adapter, error) {
	gcp, err := m.client(ctx)
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

	return &adapter{
		projectID: projectID,
		accountID: resourceID,
		desired:   obj,
		gcp:       gcp,
	}, nil
}

// Find implements the Adapter interface.
func (a *adapter) Find(ctx context.Context) (bool, error) {
	if a.accountID == "" {
		return false, nil
	}

	serviceAccount, err := a.gcp.Projects.ServiceAccounts.Get(a.fullyQualifiedName()).Context(ctx).Do()
	if err != nil {
		if IsNotFound(err) {
			klog.Warningf("serviceaccount was not found: %v", err)
			return false, nil
		}
		return false, err
	}

	u := &krm.IAMServiceAccount{}
	if err := serviceAccountToKRM(serviceAccount, u); err != nil {
		return false, err
	}
	a.actual = u

	return true, nil
}

func serviceAccountToKRM(in *api.ServiceAccount, out *krm.IAMServiceAccount) error {
	out.Spec.DisplayName = PtrTo(in.DisplayName)
	out.Spec.Description = PtrTo(in.Description)
	out.Status.Email = PtrTo(in.Email)
	// out.Status.ProjectId = in.GetProjectId()
	out.Status.UniqueId = PtrTo(in.UniqueId)
	// out.Status.Oauth2ClientId = in.GetOauth2ClientId()
	// out.Status.Disabled = in.GetDisabled()
	return nil
}

// Delete implements the Adapter interface.
func (a *adapter) Delete(ctx context.Context) (bool, error) {
	// TODO: Delete via status selfLink?
	_, err := a.gcp.Projects.ServiceAccounts.Delete(a.fullyQualifiedName()).Context(ctx).Do()
	if err != nil {
		if IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting serviceaccount: %w", err)
	}

	return true, nil
}

// Create implements the Adapter interface.
func (a *adapter) Create(ctx context.Context, u *unstructured.Unstructured) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating object", "u", u)

	// The [ServiceAccount][google.iam.admin.v1.ServiceAccount] resource to
	// create. Currently, only the following values are user assignable:
	// `display_name` and `description`.
	serviceAccount := &api.ServiceAccount{
		DisplayName: ValueOf(a.desired.Spec.DisplayName),
		Description: ValueOf(a.desired.Spec.Description),
	}

	req := &api.CreateServiceAccountRequest{
		// Name:           a.fullyQualifiedName(),
		AccountId:      a.accountID,
		ServiceAccount: serviceAccount,
	}

	parent := "projects/" + a.projectID
	created, err := a.gcp.Projects.ServiceAccounts.Create(parent, req).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("creating serviceaccount: %w", err)
	}
	log.V(2).Info("created serviceaccount", "serviceaccount", created)
	// TODO: Return created object
	return nil
}

// Update implements the Adapter interface.
func (a *adapter) Update(ctx context.Context) (*unstructured.Unstructured, error) {
	// TODO: Skip updates at the higher level if no changes?
	updateMask := &fieldmaskpb.FieldMask{}

	// TODO: I think we can do this with a helper
	if !areSame(a.desired.Spec.DisplayName, a.actual.Spec.DisplayName) {
		updateMask.Paths = append(updateMask.Paths, "display_name")
	}
	if !areSame(a.desired.Spec.Description, a.actual.Spec.Description) {
		updateMask.Paths = append(updateMask.Paths, "description")
	}

	if len(updateMask.Paths) == 0 {
		klog.Warningf("unexpected empty update mask, desired: %v, actual: %v", a.desired, a.actual)
		return nil, nil
	}

	// TODO: Where/how do we want to enforce immutability?

	// Currently, only the following fields are updatable:
	// display_name .
	sa := &api.ServiceAccount{
		DisplayName: ValueOf(a.desired.Spec.DisplayName),
		Description: ValueOf(a.desired.Spec.Description),
	}
	req := &api.PatchServiceAccountRequest{
		UpdateMask:     strings.Join(updateMask.GetPaths(), ","),
		ServiceAccount: sa,
	}

	_, err := a.gcp.Projects.ServiceAccounts.Patch(a.fullyQualifiedName(), req).Context(ctx).Do()
	if err != nil {
		return nil, err
	}
	// TODO: Return updated object
	return nil, nil
}

func (a *adapter) fullyQualifiedName() string {
	// The resource name of the service account.
	//
	// Use one of the following formats:
	//
	// * `projects/{PROJECT_ID}/serviceAccounts/{EMAIL_ADDRESS}`
	// * `projects/{PROJECT_ID}/serviceAccounts/{UNIQUE_ID}`
	//

	email := a.accountID + "@" + a.projectID + ".iam.gserviceaccount.com"
	return fmt.Sprintf("projects/%s/serviceAccounts/%s", a.projectID, email)
}
