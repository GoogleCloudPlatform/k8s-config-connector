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

	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
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
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.NetworkSecurityFirewallEndpointAssociationGVK, NewFirewallEndpointAssociationModel)
}

func NewFirewallEndpointAssociationModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelFirewallEndpointAssociation{config: *config}, nil
}

var _ directbase.Model = &modelFirewallEndpointAssociation{}

type modelFirewallEndpointAssociation struct {
	config config.ControllerConfig
}

func (m *modelFirewallEndpointAssociation) client(ctx context.Context) (*api.Service, error) {
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := api.NewService(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building FirewallEndpointAssociation client: %w", err)
	}
	return gcpClient, err
}

func (m *modelFirewallEndpointAssociation) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.NetworkSecurityFirewallEndpointAssociation{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewFirewallEndpointAssociationIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &FirewallEndpointAssociationAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
		reader:    reader,
	}, nil
}

func (m *modelFirewallEndpointAssociation) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type FirewallEndpointAssociationAdapter struct {
	id        *krm.FirewallEndpointAssociationIdentity
	gcpClient *api.Service
	desired   *krm.NetworkSecurityFirewallEndpointAssociation
	actual    *api.FirewallEndpointAssociation
	reader    client.Reader
}

var _ directbase.Adapter = &FirewallEndpointAssociationAdapter{}

// Find retrieves the GCP resource.
func (a *FirewallEndpointAssociationAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting FirewallEndpointAssociation", "name", a.id)

	association, err := a.gcpClient.Projects.Locations.FirewallEndpointAssociations.Get(a.id.String()).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting FirewallEndpointAssociation %q: %w", a.id, err)
	}

	a.actual = association
	return true, nil
}

// Create creates the resource in GCP.
func (a *FirewallEndpointAssociationAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating FirewallEndpointAssociation", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := FirewallEndpointAssociationSpec_ToAPI(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	if a.desired.Spec.FirewallEndpointRef != nil {
		ep, err := krm.ResolveNetworkSecurityFirewallEndpoint(ctx, a.reader, a.desired, a.desired.Spec.FirewallEndpointRef)
		if err != nil {
			return err
		}
		resource.FirewallEndpoint = ep
	}

	if a.desired.Spec.NetworkRef != nil {
		if err := a.desired.Spec.NetworkRef.Normalize(ctx, a.reader, a.desired.GetNamespace()); err != nil {
			return err
		}
		resource.Network = a.desired.Spec.NetworkRef.External
	}

	if a.desired.Spec.TlsInspectionPolicyRef != nil {
		tls, err := krm.ResolveNetworkSecurityTlsInspectionPolicy(ctx, a.reader, a.desired, a.desired.Spec.TlsInspectionPolicyRef)
		if err != nil {
			return err
		}
		resource.TlsInspectionPolicy = tls
	}

	op, err := a.gcpClient.Projects.Locations.FirewallEndpointAssociations.Create(a.id.Parent().String(), resource).FirewallEndpointAssociationId(a.id.ID()).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("creating FirewallEndpointAssociation %s: %w", a.id, err)
	}
	_ = op

	log.V(2).Info("successfully requested creation of FirewallEndpointAssociation", "name", a.id)

	status := &krm.NetworkSecurityFirewallEndpointAssociationStatus{}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP.
func (a *FirewallEndpointAssociationAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating FirewallEndpointAssociation", "name", a.id)
	mapCtx := &direct.MapContext{}

	desiredApi := FirewallEndpointAssociationSpec_ToAPI(mapCtx, &a.desired.DeepCopy().Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	var paths []string

	if a.desired.Spec.Labels != nil && !reflect.DeepEqual(desiredApi.Labels, a.actual.Labels) {
		paths = append(paths, "labels")
	}

	if a.desired.Spec.Disabled != nil && direct.ValueOf(a.desired.Spec.Disabled) != a.actual.Disabled {
		paths = append(paths, "disabled")
	}

	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
	} else {
		log.V(2).Info("fields need update", "name", a.id, "paths", paths)
		req := a.gcpClient.Projects.Locations.FirewallEndpointAssociations.Patch(a.id.String(), desiredApi)
		req.UpdateMask(strings.Join(paths, ","))

		op, err := req.Context(ctx).Do()
		if err != nil {
			return fmt.Errorf("updating FirewallEndpointAssociation %s: %w", a.id, err)
		}
		_ = op
		log.V(2).Info("successfully requested update of FirewallEndpointAssociation", "name", a.id)
	}

	status := &krm.NetworkSecurityFirewallEndpointAssociationStatus{}
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object.
func (a *FirewallEndpointAssociationAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.NetworkSecurityFirewallEndpointAssociation{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(FirewallEndpointAssociationSpec_FromAPI(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Parent().ProjectID}
	obj.Spec.Location = a.id.Parent().Location

	if a.actual.FirewallEndpoint != "" {
		obj.Spec.FirewallEndpointRef = &krm.NetworkSecurityFirewallEndpointRef{External: a.actual.FirewallEndpoint}
	}
	if a.actual.Network != "" {
		obj.Spec.NetworkRef = &computev1beta1.ComputeNetworkRef{External: a.actual.Network}
	}
	if a.actual.TlsInspectionPolicy != "" {
		obj.Spec.TlsInspectionPolicyRef = &krm.NetworkSecurityTlsInspectionPolicyRef{External: a.actual.TlsInspectionPolicy}
	}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.ID())
	u.SetGroupVersionKind(krm.NetworkSecurityFirewallEndpointAssociationGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP.
func (a *FirewallEndpointAssociationAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting FirewallEndpointAssociation", "name", a.id)

	op, err := a.gcpClient.Projects.Locations.FirewallEndpointAssociations.Delete(a.id.String()).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent FirewallEndpointAssociation", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting FirewallEndpointAssociation %s: %w", a.id, err)
	}
	_ = op
	log.V(2).Info("successfully requested deletion of FirewallEndpointAssociation", "name", a.id)
	return true, nil
}
