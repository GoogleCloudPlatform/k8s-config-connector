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
	"reflect"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networksecurity/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	api "google.golang.org/api/networksecurity/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.NetworkSecurityFirewallEndpointGVK, NewFirewallEndpointModel)
}

func NewFirewallEndpointModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelFirewallEndpoint{config: *config}, nil
}

var _ directbase.Model = &modelFirewallEndpoint{}

type modelFirewallEndpoint struct {
	config config.ControllerConfig
}

func (m *modelFirewallEndpoint) client(ctx context.Context) (*api.Service, error) {
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := api.NewService(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building FirewallEndpoint client: %w", err)
	}
	return gcpClient, err
}

func (m *modelFirewallEndpoint) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.NetworkSecurityFirewallEndpoint{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewFirewallEndpointIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &FirewallEndpointAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelFirewallEndpoint) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type FirewallEndpointAdapter struct {
	id        *krm.FirewallEndpointIdentity
	gcpClient *api.Service
	desired   *krm.NetworkSecurityFirewallEndpoint
	actual    *api.FirewallEndpoint
}

var _ directbase.Adapter = &FirewallEndpointAdapter{}

// Find retrieves the GCP resource.
func (a *FirewallEndpointAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting FirewallEndpoint", "name", a.id)

	endpoint, err := a.gcpClient.Projects.Locations.FirewallEndpoints.Get(a.id.String()).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting FirewallEndpoint %q: %w", a.id, err)
	}

	a.actual = endpoint
	return true, nil
}

// Create creates the resource in GCP.
func (a *FirewallEndpointAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating FirewallEndpoint", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := FirewallEndpointSpec_ToAPI(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	op, err := a.gcpClient.Projects.Locations.FirewallEndpoints.Create(a.id.Parent().String(), resource).FirewallEndpointId(a.id.ID()).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("creating FirewallEndpoint %s: %w", a.id, err)
	}
	_ = op

	log.V(2).Info("successfully requested creation of FirewallEndpoint", "name", a.id)

	mapCtx = &direct.MapContext{}
	status := &krm.NetworkSecurityFirewallEndpointStatus{}
	status.ObservedState = FirewallEndpointObservedState_FromAPI(mapCtx, a.actual)
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP.
func (a *FirewallEndpointAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating FirewallEndpoint", "name", a.id)
	mapCtx := &direct.MapContext{}

	desiredApi := FirewallEndpointSpec_ToAPI(mapCtx, &a.desired.DeepCopy().Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	var paths []string

	if a.desired.Spec.Description != nil && !reflect.DeepEqual(desiredApi.Description, a.actual.Description) {
		paths = append(paths, "description")
	}

	if a.desired.Spec.Labels != nil && !reflect.DeepEqual(desiredApi.Labels, a.actual.Labels) {
		paths = append(paths, "labels")
	}

	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
	} else {
		log.V(2).Info("fields need update", "name", a.id, "paths", paths)
		req := a.gcpClient.Projects.Locations.FirewallEndpoints.Patch(a.id.String(), desiredApi)
		var updateMask string
		for i, p := range paths {
			if i > 0 {
				updateMask += ","
			}
			updateMask += p
		}
		req.UpdateMask(updateMask)

		op, err := req.Context(ctx).Do()
		if err != nil {
			return fmt.Errorf("updating FirewallEndpoint %s: %w", a.id, err)
		}
		_ = op
		log.V(2).Info("successfully requested update of FirewallEndpoint", "name", a.id)
	}

	status := &krm.NetworkSecurityFirewallEndpointStatus{}
	status.ObservedState = FirewallEndpointObservedState_FromAPI(mapCtx, a.actual)
	status.ExternalRef = direct.LazyPtr(a.id.String())
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object.
func (a *FirewallEndpointAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.NetworkSecurityFirewallEndpoint{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(FirewallEndpointSpec_FromAPI(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Parent().ProjectID}
	obj.Spec.Location = a.id.Parent().Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.ID())
	u.SetGroupVersionKind(krm.NetworkSecurityFirewallEndpointGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP.
func (a *FirewallEndpointAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting FirewallEndpoint", "name", a.id)

	op, err := a.gcpClient.Projects.Locations.FirewallEndpoints.Delete(a.id.String()).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent FirewallEndpoint", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting FirewallEndpoint %s: %w", a.id, err)
	}
	_ = op
	log.V(2).Info("successfully requested deletion of FirewallEndpoint", "name", a.id)
	return true, nil
}
