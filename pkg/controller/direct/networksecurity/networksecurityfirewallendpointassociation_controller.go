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
	"strings"
	"time"

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
	registry.RegisterModel(krm.NetworkSecurityFirewallEndpointAssociationGVK, NewFirewallEndpointAssociationModel)
}

func NewFirewallEndpointAssociationModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &firewallEndpointAssociationModel{config: config}, nil
}

type firewallEndpointAssociationModel struct {
	config *config.ControllerConfig
}

func (m *firewallEndpointAssociationModel) client(ctx context.Context) (*api.Service, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := api.NewService(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building NetworkSecurity client: %w", err)
	}
	return gcpClient, nil
}

func (m *firewallEndpointAssociationModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	obj := &krm.NetworkSecurityFirewallEndpointAssociation{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(op.Object.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewFirewallEndpointAssociationIdentity(ctx, op.Reader, obj)
	if err != nil {
		return nil, err
	}

	if err := obj.Spec.NetworkRef.Normalize(ctx, op.Reader, obj.GetNamespace()); err != nil {
		return nil, fmt.Errorf("normalizing NetworkRef: %w", err)
	}
	if err := obj.Spec.FirewallEndpointRef.Normalize(ctx, op.Reader, obj.GetNamespace()); err != nil {
		return nil, fmt.Errorf("normalizing FirewallEndpointRef: %w", err)
	}
	if obj.Spec.TLSInspectionPolicyRef != nil {
		if err := obj.Spec.TLSInspectionPolicyRef.Normalize(ctx, op.Reader, obj.GetNamespace()); err != nil {
			return nil, fmt.Errorf("normalizing TLSInspectionPolicyRef: %w", err)
		}
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &firewallEndpointAssociationAdapter{
		id:        id,
		obj:       obj,
		gcpClient: gcpClient,
	}, nil
}

func (m *firewallEndpointAssociationModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil // Not needed for NetworkSecurityFirewallEndpointAssociation
}

type firewallEndpointAssociationAdapter struct {
	id        *krm.FirewallEndpointAssociationIdentity
	obj       *krm.NetworkSecurityFirewallEndpointAssociation
	gcpClient *api.Service
	actual    *api.FirewallEndpointAssociation
}

func waitOp(ctx context.Context, getOp func() (*api.Operation, error)) error {
	for {
		op, err := getOp()
		if err != nil {
			return fmt.Errorf("getting operation: %w", err)
		}
		if op.Done {
			if op.Error != nil {
				return fmt.Errorf("operation failed: %v", op.Error.Message)
			}
			return nil
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(2 * time.Second):
		}
	}
}

func (a *firewallEndpointAssociationAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting NetworkSecurityFirewallEndpointAssociation", "name", a.id.String())

	association, err := a.gcpClient.Projects.Locations.FirewallEndpointAssociations.Get(a.id.String()).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting NetworkSecurityFirewallEndpointAssociation %q: %w", a.id.String(), err)
	}

	a.actual = association

	mapCtx := &direct.MapContext{}
	a.obj.Status.ObservedState = NetworkSecurityFirewallEndpointAssociationObservedState_FromAPI(mapCtx, association)
	if mapCtx.Err() != nil {
		return false, mapCtx.Err()
	}

	return true, nil
}

func (a *firewallEndpointAssociationAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating NetworkSecurityFirewallEndpointAssociation", "name", a.id.String())

	mapCtx := &direct.MapContext{}
	desired := NetworkSecurityFirewallEndpointAssociationSpec_ToAPI(mapCtx, &a.obj.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	desired.Name = a.id.String()
	desired.Network = a.obj.Spec.NetworkRef.External
	desired.FirewallEndpoint = a.obj.Spec.FirewallEndpointRef.External
	if a.obj.Spec.TLSInspectionPolicyRef != nil {
		desired.TlsInspectionPolicy = a.obj.Spec.TLSInspectionPolicyRef.External
	}

	parent := fmt.Sprintf("projects/%s/locations/%s", a.id.Project, a.id.Location)

	op, err := a.gcpClient.Projects.Locations.FirewallEndpointAssociations.Create(parent, desired).FirewallEndpointAssociationId(a.id.Firewallendpointassociation).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("creating NetworkSecurityFirewallEndpointAssociation %q: %w", a.id.String(), err)
	}

	if err := waitOp(ctx, func() (*api.Operation, error) {
		return a.gcpClient.Projects.Locations.Operations.Get(op.Name).Context(ctx).Do()
	}); err != nil {
		return fmt.Errorf("waiting for create NetworkSecurityFirewallEndpointAssociation %q: %w", a.id.String(), err)
	}

	log.V(2).Info("successfully created NetworkSecurityFirewallEndpointAssociation", "name", a.id.String())

	// Fetch to get the latest state
	association, err := a.gcpClient.Projects.Locations.FirewallEndpointAssociations.Get(a.id.String()).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("getting NetworkSecurityFirewallEndpointAssociation %q: %w", a.id.String(), err)
	}

	a.obj.Status.ObservedState = NetworkSecurityFirewallEndpointAssociationObservedState_FromAPI(mapCtx, association)
	return mapCtx.Err()
}

func (a *firewallEndpointAssociationAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating NetworkSecurityFirewallEndpointAssociation", "name", a.id.String())

	mapCtx := &direct.MapContext{}
	desired := NetworkSecurityFirewallEndpointAssociationSpec_ToAPI(mapCtx, &a.obj.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	desired.Name = a.id.String()
	desired.Network = a.obj.Spec.NetworkRef.External
	desired.FirewallEndpoint = a.obj.Spec.FirewallEndpointRef.External
	if a.obj.Spec.TLSInspectionPolicyRef != nil {
		desired.TlsInspectionPolicy = a.obj.Spec.TLSInspectionPolicyRef.External
	}

	paths := []string{}
	if a.actual != nil {
		if desired.Disabled != a.actual.Disabled {
			paths = append(paths, "disabled")
		}
		if !reflect.DeepEqual(desired.Labels, a.actual.Labels) {
			paths = append(paths, "labels")
		}
		if desired.TlsInspectionPolicy != a.actual.TlsInspectionPolicy {
			paths = append(paths, "tlsInspectionPolicy")
		}
	} else {
		if a.obj.Spec.Disabled != nil {
			paths = append(paths, "disabled")
		}
		if a.obj.Spec.Labels != nil {
			paths = append(paths, "labels")
		}
		if a.obj.Spec.TLSInspectionPolicyRef != nil {
			paths = append(paths, "tlsInspectionPolicy")
		}
	}

	if len(paths) == 0 {
		log.V(2).Info("no updatable fields changed for NetworkSecurityFirewallEndpointAssociation", "name", a.id.String())
		return nil
	}

	updateMask := strings.Join(paths, ",")

	op, err := a.gcpClient.Projects.Locations.FirewallEndpointAssociations.Patch(a.id.String(), desired).UpdateMask(updateMask).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("updating NetworkSecurityFirewallEndpointAssociation %q: %w", a.id.String(), err)
	}

	if err := waitOp(ctx, func() (*api.Operation, error) {
		return a.gcpClient.Projects.Locations.Operations.Get(op.Name).Context(ctx).Do()
	}); err != nil {
		return fmt.Errorf("waiting for update NetworkSecurityFirewallEndpointAssociation %q: %w", a.id.String(), err)
	}

	log.V(2).Info("successfully updated NetworkSecurityFirewallEndpointAssociation", "name", a.id.String())

	association, err := a.gcpClient.Projects.Locations.FirewallEndpointAssociations.Get(a.id.String()).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("getting NetworkSecurityFirewallEndpointAssociation %q: %w", a.id.String(), err)
	}

	a.obj.Status.ObservedState = NetworkSecurityFirewallEndpointAssociationObservedState_FromAPI(mapCtx, association)
	return mapCtx.Err()
}

func (a *firewallEndpointAssociationAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.obj == nil {
		return nil, fmt.Errorf("obj is nil")
	}
	log := klog.FromContext(ctx)
	log.V(2).Info("exporting NetworkSecurityFirewallEndpointAssociation", "name", a.id.String())

	association, err := a.gcpClient.Projects.Locations.FirewallEndpointAssociations.Get(a.id.String()).Context(ctx).Do()
	if err != nil {
		return nil, fmt.Errorf("getting NetworkSecurityFirewallEndpointAssociation %q: %w", a.id.String(), err)
	}

	mapCtx := &direct.MapContext{}
	spec := NetworkSecurityFirewallEndpointAssociationSpec_FromAPI(mapCtx, association)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	spec.NetworkRef = &krm.NetworkRef{External: association.Network}
	spec.FirewallEndpointRef = &krm.FirewallEndpointRef{External: association.FirewallEndpoint}
	if association.TlsInspectionPolicy != "" {
		spec.TLSInspectionPolicyRef = &krm.TLSInspectionPolicyRef{External: association.TlsInspectionPolicy}
	}

	obj := &krm.NetworkSecurityFirewallEndpointAssociation{
		TypeMeta:   a.obj.TypeMeta,
		ObjectMeta: a.obj.ObjectMeta,
		Spec:       *spec,
	}

	exported, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}
	return &unstructured.Unstructured{Object: exported}, nil
}

func (a *firewallEndpointAssociationAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting NetworkSecurityFirewallEndpointAssociation", "name", a.id.String())

	op, err := a.gcpClient.Projects.Locations.FirewallEndpointAssociations.Delete(a.id.String()).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting NetworkSecurityFirewallEndpointAssociation %q: %w", a.id.String(), err)
	}

	if err := waitOp(ctx, func() (*api.Operation, error) {
		return a.gcpClient.Projects.Locations.Operations.Get(op.Name).Context(ctx).Do()
	}); err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("waiting for delete NetworkSecurityFirewallEndpointAssociation %q: %w", a.id.String(), err)
	}

	log.V(2).Info("successfully deleted NetworkSecurityFirewallEndpointAssociation", "name", a.id.String())

	return true, nil
}
