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

// +tool:controller
// proto.service: google.cloud.datastream.v1.Datastream
// proto.message: google.cloud.datastream.v1.Route
// crd.type: DatastreamRoute
// crd.version: v1alpha1

package datastream

import (
	"context"
	"fmt"
	"reflect"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/datastream/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	gcp "cloud.google.com/go/datastream/apiv1"
	pb "cloud.google.com/go/datastream/apiv1/datastreampb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.DatastreamRouteGVK, NewRouteModel)
}

func NewRouteModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelRoute{config: *config}, nil
}

var _ directbase.Model = &modelRoute{}

type modelRoute struct {
	config config.ControllerConfig
}

func (m *modelRoute) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.DatastreamRoute{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewRouteIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Get datastream GCP client
	gcpClient, err := newGCPClient(ctx, &m.config)
	if err != nil {
		return nil, err
	}
	datastreamClient, err := gcpClient.newDatastreamClient(ctx)
	if err != nil {
		return nil, err
	}
	return &RouteAdapter{
		id:        id,
		gcpClient: datastreamClient,
		desired:   obj,
		reader:    reader,
	}, nil
}

func (m *modelRoute) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type RouteAdapter struct {
	id        *krm.RouteIdentity
	gcpClient *gcp.Client
	desired   *krm.DatastreamRoute
	actual    *pb.Route
	reader    client.Reader
}

var _ directbase.Adapter = &RouteAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *RouteAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting Route", "name", a.id)

	req := &pb.GetRouteRequest{Name: a.id.String()}
	routepb, err := a.gcpClient.GetRoute(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting Route %q: %w", a.id, err)
	}

	a.actual = routepb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *RouteAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating Route", "name", a.id)

	if err := a.normalizeReferenceFields(ctx); err != nil {
		return err
	}

	mapCtx := &direct.MapContext{}
	desired := a.desired.DeepCopy()
	resource := DatastreamRouteSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	parent, err := a.id.Parent(ctx, a.reader)
	if err != nil {
		return err
	}

	req := &pb.CreateRouteRequest{
		Parent:  parent.String(),
		RouteId: a.id.ID(),
		Route:   resource,
	}
	op, err := a.gcpClient.CreateRoute(ctx, req)
	if err != nil {
		return fmt.Errorf("creating Route %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("Route %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created Route", "name", a.id)

	status := &krm.DatastreamRouteStatus{}
	status.ObservedState = DatastreamRouteObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *RouteAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating Route", "name", a.id)

	if err := a.normalizeReferenceFields(ctx); err != nil {
		return err
	}

	mapCtx := &direct.MapContext{}
	desired := a.desired.DeepCopy()
	resource := DatastreamRouteSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	// Check for differences. DatastreamRoute is immutable except for labels and displayName.
	// However, there is no UpdateRoute RPC. Therefore, we cannot update the resource.
	// We only check if an update is needed to decide whether to error out or just update status.
	changed := false
	if desired.Spec.DisplayName != nil && !reflect.DeepEqual(resource.DisplayName, a.actual.DisplayName) {
		changed = true
	}
	if desired.Spec.Labels != nil && !reflect.DeepEqual(resource.Labels, a.actual.Labels) {
		changed = true
	}
	if desired.Spec.DestinationAddress != nil && !reflect.DeepEqual(resource.DestinationAddress, a.actual.DestinationAddress) {
		changed = true
	}
	if desired.Spec.DestinationPort != nil && !reflect.DeepEqual(resource.DestinationPort, a.actual.DestinationPort) {
		changed = true
	}

	if changed {
		return fmt.Errorf("updating Route %q is not supported", a.id)
	}

	log.V(2).Info("no update needed for Route", "name", a.id)

	// Update status even if no GCP changes are made (e.g. acquiring existing resource)
	status := &krm.DatastreamRouteStatus{}
	status.ObservedState = DatastreamRouteObservedState_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *RouteAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.DatastreamRoute{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(DatastreamRouteSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	parentID, err := a.id.Parent(ctx, a.reader)
	if err != nil {
		return nil, fmt.Errorf("getting parent identity: %w", err)
	}

	obj.Spec.ProjectRef = &refs.ProjectRef{External: parentID.Parent().ProjectID}
	obj.Spec.Location = parentID.Parent().Location
	obj.Spec.PrivateConnectionRef = &krm.PrivateConnectionRef{External: parentID.String()}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, fmt.Errorf("converting DatastreamRoute to unstructured: %w", err)
	}

	u.SetName(a.id.ID()) // Use route_id as K8s name
	u.SetGroupVersionKind(krm.DatastreamRouteGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *RouteAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting Route", "name", a.id)

	req := &pb.DeleteRouteRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteRoute(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent Route, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting Route %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted Route", "name", a.id)

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete Route %s: %w", a.id, err)
	}
	return true, nil
}

func (a *RouteAdapter) normalizeReferenceFields(ctx context.Context) error {
	// obj := a.desired
	// Route spec has no fields needing reference normalization,
	// the PrivateConnectionRef is handled by the identity construction.
	return nil
}
