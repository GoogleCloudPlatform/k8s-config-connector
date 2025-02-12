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
	"strings"
	"time"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/cloudidentity/v1beta1"
	cloudidentitygrouppb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/cloudidentity/groups/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	gcp "google.golang.org/api/cloudidentity/v1beta1"
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

	id, err := krm.NewGroupIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Get cloudidentitygroup GCP client
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewService(ctx, opts...)
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
	gcpClient *gcp.Service
	desired   *krm.CloudIdentityGroup
	actual    *cloudidentitygrouppb.Group
}

var _ directbase.Adapter = &GroupAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *GroupAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting Group", "name", a.id)

	desired := a.desired.DeepCopy()
	if desired.Status.ExternalRef == nil {
		return false, nil
	}

	generatedId := direct.ValueOf(desired.Status.ExternalRef)
	resource, err := a.gcpClient.Groups.Get(generatedId).Context(ctx).Do()
	if err != nil {
		return false, fmt.Errorf("getting Group %q: %w", a.id, err)
	}

	if err := convertAPIToProto(resource, &a.actual); err != nil {
		return false, err
	}
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *GroupAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating Group", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := CloudIdentityGroupSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &gcp.Group{}
	if err := convertProtoToAPI(resource, req); err != nil {
		return err
	}

	op, err := a.gcpClient.Groups.Create(req).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("creating Group %s: %w", a.id, err)
	}

	// Get server generated group name
	var data interface{}
	err = json.Unmarshal(op.Response, &data)
	generatedName := data.(map[string]interface{})["name"].(string)

	created, err := a.gcpClient.Groups.Get(generatedName).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("getting created Group %q: %w", a.id, err)
	}

	var createdPB *cloudidentitygrouppb.Group
	if err := convertAPIToProto(created, &createdPB); err != nil {
		return err
	}

	log.V(2).Info("successfully created Group", "name", a.id)

	status := &krm.CloudIdentityGroupStatus{
		Name:       direct.LazyPtr(created.Name),
		CreateTime: direct.LazyPtr(created.CreateTime),
		UpdateTime: direct.LazyPtr(created.UpdateTime),
	}
	status.ObservedState = CloudIdentityGroupObservedState_FromProto(mapCtx, createdPB)
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
	resource := CloudIdentityGroupSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	generatedId := desired.Status.ExternalRef
	resource.Name = generatedId

	paths, err := common.CompareProtoMessage(resource, a.actual, common.BasicDiff)
	if err != nil {
		return err
	}

	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		status := &krm.CloudIdentityGroupStatus{}
		status.ObservedState = CloudIdentityGroupObservedState_FromProto(mapCtx, a.actual)
		if mapCtx.Err() != nil {
			return mapCtx.Err()
		}
		return updateOp.UpdateStatus(ctx, status, nil)
	}

	// updateMask is a comma-separated list of fully qualified names of fields.
	var stringSlice []string
	for path := range paths {
		stringSlice = append(stringSlice, path)
	}
	updateMask := strings.Join(stringSlice, ",")

	req := &gcp.Group{}
	if err := convertProtoToAPI(resource, req); err != nil {
		return err
	}

	_, err = a.gcpClient.Groups.Patch(direct.ValueOf(generatedId), req).UpdateMask(updateMask).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("updating Group %s: %w", a.id, err)
	}

	// A test-google-a.com additionalGroupKey which is auto-added by the service.
	// Wait a bit for it to be added.
	time.Sleep(time.Second)
	updated, err := a.gcpClient.Groups.Get(direct.ValueOf(generatedId)).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("getting updated Group %q: %w", a.id, err)
	}
	var updatedPB *cloudidentitygrouppb.Group
	if err := convertAPIToProto(updated, &updatedPB); err != nil {
		return err
	}
	log.V(2).Info("successfully updated Group", "name", a.id)

	status := &krm.CloudIdentityGroupStatus{
		Name:       direct.LazyPtr(updated.Name),
		CreateTime: direct.LazyPtr(updated.CreateTime),
		UpdateTime: direct.LazyPtr(updated.UpdateTime),
	}
	status.ObservedState = CloudIdentityGroupObservedState_FromProto(mapCtx, updatedPB)
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
	obj.Spec = direct.ValueOf(CloudIdentityGroupSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.actual.GetName())
	u.SetGroupVersionKind(krm.CloudIdentityGroupGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *GroupAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting Group", "name", a.id)

	desired := a.desired.DeepCopy()
	if desired.Status.ExternalRef == nil {
		return false, nil
	}
	generatedId := direct.ValueOf(desired.Status.ExternalRef)

	_, err := a.gcpClient.Groups.Delete(generatedId).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting Group %q: %w", a.id, err)
	}

	log.V(2).Info("successfully deleted Group", "name", a.id)
	return true, nil
}
