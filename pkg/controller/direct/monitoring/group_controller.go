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

package monitoring

import (
	"context"
	"fmt"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/monitoring/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	// TODO(contributor): Update the import with the google cloud client
	gcp "cloud.google.com/go/monitoring/apiv3/v2"

	// TODO(contributor): Update the import with the google cloud client api protobuf
	monitoringpb "cloud.google.com/go/monitoring/apiv3/v2/monitoringpb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.MonitoringGroupGVK, NewGroupModel)
}

func NewGroupModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelGroup{config: *config}, nil
}

var _ directbase.Model = &modelGroup{}

type modelGroup struct {
	config config.ControllerConfig
}

func (m *modelGroup) client(ctx context.Context) (*gcp.GroupClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewGroupClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building Group client: %w", err)
	}
	return gcpClient, err
}

func (m *modelGroup) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.MonitoringGroup{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewGroupIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Get monitoring GCP client
	gcpClient, err := m.client(ctx)
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
	gcpClient *gcp.GroupClient
	desired   *krm.MonitoringGroup
	actual    *monitoringpb.Group
}

var _ directbase.Adapter = &GroupAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *GroupAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting Group", "name", a.id)

	req := &monitoringpb.GetGroupRequest{Name: a.id.String()}
	grouppb, err := a.gcpClient.GetGroup(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting Group %q: %w", a.id, err)
	}

	a.actual = grouppb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *GroupAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating Group", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := MonitoringGroupSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	// TODO(contributor): Complete the gcp "CREATE" or "INSERT" request.
	req := &monitoringpb.CreateGroupRequest{
		Name:  a.id.Parent().String(),
		Group: resource,
	}
	created, err := a.gcpClient.CreateGroup(ctx, req)
	if err != nil {
		return fmt.Errorf("creating Group %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created Group", "name", a.id)

	status := &krm.MonitoringGroupStatus{}
	status.ObservedState = MonitoringGroupObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *GroupAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating Group", "name", a.id)
	mapCtx := &direct.MapContext{}

	desiredPb := MonitoringGroupSpec_ToProto(mapCtx, &a.desired.DeepCopy().Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	var err error
	paths, err := common.CompareProtoMessage(desiredPb, a.actual, common.BasicDiff)
	if err != nil {
		return err
	}
	var pathSet sets.Set[string]
	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		status := &krm.MonitoringGroupStatus{}
		status.ObservedState = MonitoringGroupObservedState_FromProto(mapCtx, a.actual)
		if mapCtx.Err() != nil {
			return mapCtx.Err()
		}
		return updateOp.UpdateStatus(ctx, status, nil)
	}
	pathSet = sets.New[string]()
	for p := range paths {
		pathSet.Insert(p)
	}

	// TODO(contributor): Complete the gcp "UPDATE" or "PATCH" request.
	req := &monitoringpb.UpdateGroupRequest{
		Group:        desiredPb,
		ValidateOnly: false,
	}
	updated, err := a.gcpClient.UpdateGroup(ctx, req)
	if err != nil {
		return fmt.Errorf("updating Group %s: %w", a.id, err)
	}
	log.V(2).Info("successfully updated Group", "name", a.id)

	status := &krm.MonitoringGroupStatus{}
	status.ObservedState = MonitoringGroupObservedState_FromProto(mapCtx, updated)
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

	obj := &krm.MonitoringGroup{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(MonitoringGroupSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Parent().ProjectID}
	obj.Spec.Location = a.id.Parent().Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.actual.Name)
	u.SetGroupVersionKind(krm.MonitoringGroupGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *GroupAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting Group", "name", a.id)

	req := &monitoringpb.DeleteGroupRequest{Name: a.id.String()}
	err := a.gcpClient.DeleteGroup(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent Group, assuming it was already deleted", "name", a.id.String())
			return true, nil
		}
		return false, fmt.Errorf("deleting Group %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted Group", "name", a.id)

	if err != nil {
		return false, fmt.Errorf("waiting delete Group %s: %w", a.id, err)
	}
	return true, nil
}
