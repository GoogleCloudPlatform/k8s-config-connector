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

	gcp "cloud.google.com/go/networksecurity/apiv1beta1"
	networksecuritypb "cloud.google.com/go/networksecurity/apiv1beta1/networksecuritypb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

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

func (m *modelFirewallEndpoint) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
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

	// No references to resolve for FirewallEndpoint (at least for now based on discovery doc)

	// Get networksecurity GCP client
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
	gcpClient *gcp.Client
	desired   *krm.NetworkSecurityFirewallEndpoint
	actual    *networksecuritypb.FirewallEndpoint
}

var _ directbase.Adapter = &FirewallEndpointAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *FirewallEndpointAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting FirewallEndpoint", "name", a.id)

	req := &networksecuritypb.GetFirewallEndpointRequest{Name: a.id.String()}
	endpointpb, err := a.gcpClient.GetFirewallEndpoint(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting FirewallEndpoint %q: %w", a.id, err)
	}

	a.actual = endpointpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *FirewallEndpointAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating FirewallEndpoint", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := FirewallEndpointSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &networksecuritypb.CreateFirewallEndpointRequest{
		Parent:             a.id.Parent().String(),
		FirewallEndpoint:   resource,
		FirewallEndpointId: a.id.ID(),
	}
	op, err := a.gcpClient.CreateFirewallEndpoint(ctx, req)
	if err != nil {
		return fmt.Errorf("creating FirewallEndpoint %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("FirewallEndpoint %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created FirewallEndpoint", "name", a.id)

	status := &krm.NetworkSecurityFirewallEndpointStatus{}
	status.ObservedState = FirewallEndpointObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *FirewallEndpointAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating FirewallEndpoint", "name", a.id)
	mapCtx := &direct.MapContext{}

	desiredPb := FirewallEndpointSpec_ToProto(mapCtx, &a.desired.DeepCopy().Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	var paths []string

	// Optional. Description of the firewall endpoint. Max length 2048 characters.
	if a.desired.Spec.Description != nil && !reflect.DeepEqual(desiredPb.GetDescription(), a.actual.GetDescription()) {
		paths = append(paths, "description")
	}

	// Optional. Labels as key value pairs
	if a.desired.Spec.Labels != nil && !reflect.DeepEqual(desiredPb.GetLabels(), a.actual.GetLabels()) {
		paths = append(paths, "labels")
	}

	updated := a.actual
	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
	} else {
		log.V(2).Info("fields need update", "name", a.id, "paths", paths)
		updateMask := &fieldmaskpb.FieldMask{
			Paths: paths,
		}
		req := &networksecuritypb.UpdateFirewallEndpointRequest{
			UpdateMask:       updateMask,
			FirewallEndpoint: desiredPb,
		}
		req.FirewallEndpoint.Name = a.id.String()
		op, err := a.gcpClient.UpdateFirewallEndpoint(ctx, req)
		if err != nil {
			return fmt.Errorf("updating FirewallEndpoint %s: %w", a.id, err)
		}
		updated, err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("FirewallEndpoint %s waiting update: %w", a.id, err)
		}
		log.V(2).Info("successfully updated FirewallEndpoint", "name", a.id)
	}

	status := &krm.NetworkSecurityFirewallEndpointStatus{}
	status.ObservedState = FirewallEndpointObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *FirewallEndpointAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.NetworkSecurityFirewallEndpoint{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(FirewallEndpointSpec_FromProto(mapCtx, a.actual))
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

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *FirewallEndpointAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting FirewallEndpoint", "name", a.id)

	req := &networksecuritypb.DeleteFirewallEndpointRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteFirewallEndpoint(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent FirewallEndpoint, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting FirewallEndpoint %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted FirewallEndpoint", "name", a.id)

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete FirewallEndpoint %s: %w", a.id, err)
	}
	return true, nil
}
