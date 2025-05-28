// Copyright 2025 Google LLC
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

package cloudidentity

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/cloudidentity/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/google/go-cmp/cmp"
	api "google.golang.org/api/cloudidentity/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.CloudIdentityMembershipGVK, NewMembershipModel)
}

func NewMembershipModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelMembership{config: *config}, nil
}

var _ directbase.Model = &modelMembership{}

type modelMembership struct {
	config config.ControllerConfig
}

func (m *modelMembership) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.CloudIdentityMembership{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewMembershipIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Get cloudidentity GCP client
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := api.NewService(ctx, opts...)
	if err != nil {
		return nil, err
	}
	return &MembershipAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelMembership) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type MembershipAdapter struct {
	id        *krm.MembershipIdentity
	gcpClient *api.Service
	desired   *krm.CloudIdentityMembership
	actual    *api.Membership
}

var _ directbase.Adapter = &MembershipAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *MembershipAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting Membership", "name", a.id)

	// Check whether Config Connector knows the resource identity.
	// If not, Config Connector saves one GCP GET call, and starts the CREATE call directly.
	// This is mostly for GCP services that do not allow user to specify ID, but assign an ID when creating the object.
	if a.id.ID() == "" {
		return false, nil
	}

	resource, err := a.gcpClient.Groups.Memberships.Get(a.id.String()).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting Membership %q: %w", a.id, err)
	}

	a.actual = resource
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *MembershipAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating Membership", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := CloudIdentityMembershipSpec_ToAPI(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	op, err := a.gcpClient.Groups.Memberships.Create(a.id.Parent(), resource).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("creating Membership %s: %w", a.id, err)
	}
	if err := WaitForCloudIdentityOp(ctx, op); err != nil {
		return fmt.Errorf("error waiting Membership %s creation: %w", a.id, err)
	}

	// Get server generated membership name
	var data interface{}
	err = json.Unmarshal(op.Response, &data)
	if err != nil {
		return err
	}
	generatedName := data.(map[string]interface{})["name"].(string)

	created, err := a.gcpClient.Groups.Memberships.Get(generatedName).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("getting created Membership %q: %w", a.id, err)
	}
	log.V(2).Info("successfully created Membership", "name", a.id)

	status := CloudIdentityMembershipStatus_FromAPI(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	externalRef := generatedName
	status.ExternalRef = &externalRef
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *MembershipAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating Membership", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := CloudIdentityMembershipSpec_ToAPI(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	if reflect.DeepEqual(resource.Roles, a.actual.Roles) {
		log.V(2).Info("no field needs update", "name", a.id)
		status := CloudIdentityMembershipStatus_FromAPI(mapCtx, a.actual)
		if mapCtx.Err() != nil {
			return mapCtx.Err()
		}
		return updateOp.UpdateStatus(ctx, status, nil)
	}

	var addRoles []*api.MembershipRole
	var removeRoles []string
	var updateRolesParams []*api.UpdateMembershipRolesParams

	beforeSet := make(map[string]bool)
	if a.actual.Roles != nil {
		for _, item := range a.actual.Roles {
			beforeSet[item.Name] = true
		}
	}

	afterSet := make(map[string]bool)
	if resource.Roles != nil {
		for _, item := range resource.Roles {
			afterSet[item.Name] = true
		}
	}

	for _, item := range resource.Roles {
		if !beforeSet[item.Name] {
			addRoles = append(addRoles, item)
		}
	}

	for _, item := range a.actual.Roles {
		name := item.Name
		if !afterSet[name] {
			removeRoles = append(removeRoles, name)
		}
	}

	for _, afterRole := range resource.Roles {
		for _, beforeRole := range a.actual.Roles {
			if afterRole.Name == beforeRole.Name {
				if afterRole.ExpiryDetail == nil && beforeRole.ExpiryDetail == nil {
					continue
				}
				if afterRole.ExpiryDetail.ExpireTime == "" && beforeRole.ExpiryDetail.ExpireTime == "" {
					continue
				}
				if !cmp.Equal(afterRole, beforeRole) {
					updateRolesParam := &api.UpdateMembershipRolesParams{
						// Only expiry_detail.expire_time is configurable and can be updated
						FieldMask:      "expiry_detail.expire_time",
						MembershipRole: afterRole,
					}
					updateRolesParams = append(updateRolesParams, updateRolesParam)
				}
			}
		}
	}

	req := &api.ModifyMembershipRolesRequest{
		AddRoles:    addRoles,
		RemoveRoles: removeRoles,
		// UpdateRolesParams: The `MembershipRole`s to be updated. Updating roles in
		// the same request as adding or removing roles is not supported. Must not be
		// set if either `add_roles` or `remove_roles` is set.
		// todo: shall we handle this or let the API to handle this?
		UpdateRolesParams: updateRolesParams,
	}

	_, err := a.gcpClient.Groups.Memberships.ModifyMembershipRoles(a.id.String(), req).Context(ctx).Do()

	if err != nil {
		return fmt.Errorf("updating Membership %s: %w", a.id, err)
	}

	updated, err := a.gcpClient.Groups.Memberships.Get(a.id.String()).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("getting updated Membership %q: %w", a.id, err)
	}
	log.V(2).Info("successfully updated Membership", "name", a.id)

	status := CloudIdentityMembershipStatus_FromAPI(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *MembershipAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.CloudIdentityMembership{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(CloudIdentityMembershipSpec_FromAPI(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.actual.Name)
	u.SetGroupVersionKind(krm.CloudIdentityMembershipGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *MembershipAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting Membership", "name", a.id)

	op, err := a.gcpClient.Groups.Memberships.Delete(a.id.String()).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting Membership %q: %w", a.id, err)
	}
	if err := WaitForCloudIdentityOp(ctx, op); err != nil {
		return false, fmt.Errorf("error waiting Membership %s deletion: %w", a.id, err)
	}

	log.V(2).Info("successfully deleted Membership", "name", a.id)
	return true, nil
}
