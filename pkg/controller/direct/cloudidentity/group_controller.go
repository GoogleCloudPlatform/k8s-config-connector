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
	"sort"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/cloudidentity/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	api "google.golang.org/api/cloudidentity/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.CloudIdentityGroupGVK, NewGroupModel)
}

func NewGroupModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelGroup{config: *config}, nil
}

var _ directbase.Model = &modelGroup{}

type modelGroup struct {
	config config.ControllerConfig
}

func (m *modelGroup) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.CloudIdentityGroup{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	// Handle TF default values
	if obj.Spec.InitialGroupConfig == nil {
		obj.Spec.InitialGroupConfig = direct.LazyPtr("EMPTY")
	}

	id, err := krm.NewGroupIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Get cloudidentitygroup GCP client
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := api.NewService(ctx, opts...)
	if err != nil {
		return nil, err
	}
	return &GroupAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelGroup) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type GroupAdapter struct {
	id        *krm.GroupIdentity
	gcpClient *api.Service
	desired   *krm.CloudIdentityGroup
	actual    *api.Group
}

var _ directbase.Adapter = &GroupAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *GroupAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting Group", "name", a.id)

	// Check whether Config Connector knows the resource identity.
	// If not, Config Connector saves one GCP GET call, and starts the CREATE call directly.
	// This is mostly for GCP services that do not allow user to specify ID, but assign an ID when creating the object.
	if a.id.ID() == "" {
		return false, nil
	}

	resource, err := a.gcpClient.Groups.Get(a.id.String()).Context(ctx).Do()
	if err != nil {
		return false, fmt.Errorf("getting Group %q: %w", a.id, err)
	}

	a.actual = resource
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *GroupAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating Group", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := CloudIdentityGroupSpec_ToAPI(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	initialGroupConfig := direct.ValueOf(desired.Spec.InitialGroupConfig)
	op, err := a.gcpClient.Groups.Create(resource).InitialGroupConfig(initialGroupConfig).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("creating Group %s: %w", a.id, err)
	}
	if err := WaitForCloudIdentityOp(ctx, op); err != nil {
		return fmt.Errorf("error waiting Group %s deletion: %w", a.id, err)
	}

	// Get server generated group name
	var data interface{}
	err = json.Unmarshal(op.Response, &data)
	if err != nil {
		return err
	}
	generatedName := data.(map[string]interface{})["name"].(string)

	created, err := a.gcpClient.Groups.Get(generatedName).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("getting created Group %q: %w", a.id, err)
	}

	log.V(2).Info("successfully created Group", "name", a.id)

	status := &krm.CloudIdentityGroupStatus{}
	status = CloudIdentityGroupStatus_FromAPI(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	externalRef := generatedName
	status.ExternalRef = &externalRef
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *GroupAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating Group", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := CloudIdentityGroupSpec_ToAPI(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	var paths []string
	if !reflect.DeepEqual(resource.DisplayName, a.actual.DisplayName) {
		paths = append(paths, "display_name")
	}
	if !reflect.DeepEqual(resource.Description, a.actual.Description) {
		paths = append(paths, "description")
	}
	if !reflect.DeepEqual(resource.Labels, a.actual.Labels) {
		paths = append(paths, "labels")
	}

	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		status := &krm.CloudIdentityGroupStatus{}
		status = CloudIdentityGroupStatus_FromAPI(mapCtx, a.actual)
		if mapCtx.Err() != nil {
			return mapCtx.Err()
		}
		return updateOp.UpdateStatus(ctx, status, nil)
	}

	// updateMask is a comma-separated list of fully qualified names of fields.
	sort.Strings(paths)
	updateMask := strings.Join(paths, ",")

	op, err := a.gcpClient.Groups.Patch(a.id.String(), resource).UpdateMask(updateMask).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("updating Group %s: %w", a.id, err)
	}
	if err := WaitForCloudIdentityOp(ctx, op); err != nil {
		return fmt.Errorf("error waiting Group %s update: %w", a.id, err)
	}

	updated, err := a.gcpClient.Groups.Get(a.id.String()).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("getting updated Group %q: %w", a.id, err)
	}
	log.V(2).Info("successfully updated Group", "name", a.id)

	status := &krm.CloudIdentityGroupStatus{}
	status = CloudIdentityGroupStatus_FromAPI(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *GroupAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.CloudIdentityGroup{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(CloudIdentityGroupSpec_FromAPI(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.actual.Name)
	u.SetGroupVersionKind(krm.CloudIdentityGroupGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *GroupAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting Group", "name", a.id)

	op, err := a.gcpClient.Groups.Delete(a.id.String()).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting Group %q: %w", a.id, err)
	}
	if err := WaitForCloudIdentityOp(ctx, op); err != nil {
		return false, fmt.Errorf("error waiting Group %s deletion: %w", a.id, err)
	}

	log.V(2).Info("successfully deleted Group", "name", a.id)
	return true, nil
}
