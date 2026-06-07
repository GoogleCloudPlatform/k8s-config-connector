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

package networksecurity

import (
	"context"
	"fmt"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networksecurity/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	api "google.golang.org/api/networksecurity/v1"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.NetworkSecurityMirroringDeploymentGroupGVK, NewMirroringDeploymentGroupModel)
}

func NewMirroringDeploymentGroupModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &mirroringDeploymentGroupModel{config: *config}, nil
}

var _ directbase.Model = &mirroringDeploymentGroupModel{}

type mirroringDeploymentGroupModel struct {
	config config.ControllerConfig
}

func (m *mirroringDeploymentGroupModel) client(ctx context.Context) (*api.Service, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := api.NewService(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building MirroringDeploymentGroup client: %w", err)
	}
	return gcpClient, err
}

func (m *mirroringDeploymentGroupModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	obj := &krm.NetworkSecurityMirroringDeploymentGroup{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	idVal, err := obj.GetIdentity(ctx, op.Reader)
	if err != nil {
		return nil, err
	}

	id, ok := idVal.(*krm.NetworkSecurityMirroringDeploymentGroupIdentity)
	if !ok {
		return nil, fmt.Errorf("unexpected identity type: %T", idVal)
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &mirroringDeploymentGroupAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *mirroringDeploymentGroupModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support Import
	return nil, nil
}

type mirroringDeploymentGroupAdapter struct {
	id        *krm.NetworkSecurityMirroringDeploymentGroupIdentity
	gcpClient *api.Service
	desired   *krm.NetworkSecurityMirroringDeploymentGroup
}

var _ directbase.Adapter = &mirroringDeploymentGroupAdapter{}

func (a *mirroringDeploymentGroupAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting MirroringDeploymentGroup", "name", a.id.String())

	mirroringDeploymentGroup, err := a.gcpClient.Projects.Locations.MirroringDeploymentGroups.Get(a.id.String()).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting MirroringDeploymentGroup %q: %w", a.id.String(), err)
	}

	observedState := MirroringDeploymentGroupObservedState_FromAPI(mirroringDeploymentGroup)
	a.desired.Status.ObservedState = observedState
	a.desired.Status.ObservedGeneration = direct.LazyPtr(a.desired.Generation)
	a.desired.Status.ExternalRef = direct.LazyPtr(a.id.String())
	return true, nil
}

func (a *mirroringDeploymentGroupAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating MirroringDeploymentGroup", "name", a.id.String())

	resource := MirroringDeploymentGroupSpec_ToAPI(&a.desired.Spec)

	parent := fmt.Sprintf("projects/%s/locations/%s", a.id.Project, a.id.Location)
	op, err := a.gcpClient.Projects.Locations.MirroringDeploymentGroups.Create(parent, resource).MirroringDeploymentGroupId(a.id.Mirroring_deployment_group).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("creating MirroringDeploymentGroup %s: %w", a.id.String(), err)
	}

	// For lro
	if op.Error != nil {
		return fmt.Errorf("creating MirroringDeploymentGroup %s error: %v", a.id.String(), op.Error)
	}

	log.V(2).Info("successfully started create for MirroringDeploymentGroup", "name", a.id.String())

	status := &krm.NetworkSecurityMirroringDeploymentGroupStatus{}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *mirroringDeploymentGroupAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating MirroringDeploymentGroup", "name", a.id.String())

	resource := MirroringDeploymentGroupSpec_ToAPI(&a.desired.Spec)

	// Build update mask
	updateMask := ""
	if a.desired.Spec.Labels != nil {
		updateMask += "labels,"
	}
	if a.desired.Spec.Description != nil {
		updateMask += "description,"
	}
	if updateMask != "" {
		updateMask = updateMask[:len(updateMask)-1]
	}

	op, err := a.gcpClient.Projects.Locations.MirroringDeploymentGroups.Patch(a.id.String(), resource).UpdateMask(updateMask).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("updating MirroringDeploymentGroup %s: %w", a.id.String(), err)
	}
	if op.Error != nil {
		return fmt.Errorf("updating MirroringDeploymentGroup %s error: %v", a.id.String(), op.Error)
	}

	log.V(2).Info("successfully started update for MirroringDeploymentGroup", "name", a.id.String())

	status := &krm.NetworkSecurityMirroringDeploymentGroupStatus{}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *mirroringDeploymentGroupAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.desired == nil {
		return nil, fmt.Errorf("Export is not supported")
	}
	return nil, fmt.Errorf("Export is not supported")
}

func (a *mirroringDeploymentGroupAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting MirroringDeploymentGroup", "name", a.id.String())

	op, err := a.gcpClient.Projects.Locations.MirroringDeploymentGroups.Delete(a.id.String()).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting MirroringDeploymentGroup %s: %w", a.id.String(), err)
	}
	if op.Error != nil {
		return false, fmt.Errorf("deleting MirroringDeploymentGroup %s error: %v", a.id.String(), op.Error)
	}

	log.V(2).Info("successfully started delete for MirroringDeploymentGroup", "name", a.id.String())

	return true, nil
}

func MirroringDeploymentGroupObservedState_FromAPI(in *api.MirroringDeploymentGroup) *krm.NetworkSecurityMirroringDeploymentGroupObservedState {
	if in == nil {
		return nil
	}
	out := &krm.NetworkSecurityMirroringDeploymentGroupObservedState{}
	out.CreateTime = direct.LazyPtr(in.CreateTime)
	out.UpdateTime = direct.LazyPtr(in.UpdateTime)
	out.State = direct.LazyPtr(in.State)
	out.Reconciling = direct.LazyPtr(in.Reconciling)
	for _, v := range in.ConnectedEndpointGroups {
		out.ConnectedEndpointGroups = append(out.ConnectedEndpointGroups, krm.MirroringDeploymentGroupConnectedEndpointGroupObservedState{
			Name: direct.LazyPtr(v.Name),
		})
	}
	for _, v := range in.NestedDeployments {
		out.NestedDeployments = append(out.NestedDeployments, krm.MirroringDeploymentGroupDeploymentObservedState{
			Name:  direct.LazyPtr(v.Name),
			State: direct.LazyPtr(v.State),
		})
	}
	for _, v := range in.Locations {
		out.Locations = append(out.Locations, krm.MirroringLocationObservedState{
			Location: direct.LazyPtr(v.Location),
			State:    direct.LazyPtr(v.State),
		})
	}
	return out
}

func MirroringDeploymentGroupSpec_ToAPI(in *krm.NetworkSecurityMirroringDeploymentGroupSpec) *api.MirroringDeploymentGroup {
	if in == nil {
		return nil
	}
	out := &api.MirroringDeploymentGroup{}
	out.Description = direct.ValueOf(in.Description)
	out.Labels = in.Labels
	if in.NetworkRef != nil {
		out.Network = in.NetworkRef.External
	}
	return out
}
