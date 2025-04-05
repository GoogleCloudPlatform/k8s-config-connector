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
// proto.service: google.cloud.networkconnectivity.v1.HubService
// proto.message: google.cloud.networkconnectivity.v1.RegionalEndpoint
// crd.type: NetworkConnectivityRegionalEndpoint
// crd.version: v1alpha1

package networkconnectivity

import (
	"context"
	"fmt"
	"time"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/networkconnectivity/v1"
	api "google.golang.org/api/networkconnectivity/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networkconnectivity/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
)

func init() {
	registry.RegisterModel(krm.NetworkConnectivityRegionalEndpointGVK, NewRegionalEndpointModel)
}

func NewRegionalEndpointModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &regionalEndpointModel{config: *config}, nil
}

var _ directbase.Model = &regionalEndpointModel{}

type regionalEndpointModel struct {
	config config.ControllerConfig
}

func (m *regionalEndpointModel) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.NetworkConnectivityRegionalEndpoint{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewRegionalEndpointIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// normalize reference fields
	if obj.Spec.NetworkRef != nil {
		if err := obj.Spec.NetworkRef.Normalize(ctx, reader, obj); err != nil {
			return nil, err
		}
	}

	// Get networkconnectivity GCP client
	gcpClient, err := newGCPClient(ctx, &m.config)
	if err != nil {
		return nil, err
	}
	client, err := gcpClient.newNetworkConnectivityClient(ctx)
	if err != nil {
		return nil, err
	}
	return &regionalEndpointAdapter{
		gcpClient: client,
		id:        id,
		desired:   obj,
	}, nil
}

func (m *regionalEndpointModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type regionalEndpointAdapter struct {
	gcpClient *api.Service
	id        *krm.RegionalEndpointIdentity
	desired   *krm.NetworkConnectivityRegionalEndpoint
	actual    *pb.RegionalEndpoint
}

var _ directbase.Adapter = &regionalEndpointAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *regionalEndpointAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting networkconnectivity regionalendpoint", "name", a.id)
	fqn := a.id.String()
	actual, err := a.gcpClient.Projects.Locations.RegionalEndpoints.Get(fqn).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting networkconnectivity regionalendpoint %q from gcp: %w", a.id.String(), err)
	}

	if err := convertAPIToProto(actual, &a.actual); err != nil {
		return false, err
	}

	return true, nil
}

func (a *regionalEndpointAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	u := createOp.GetUnstructured()

	log := klog.FromContext(ctx)
	log.V(2).Info("creating networkconnectivity regionalendpoint", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := NetworkConnectivityRegionalEndpointSpec_ToProto(mapCtx, &a.desired.DeepCopy().Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	req := &api.RegionalEndpoint{}
	err := convertProtoToAPI(desired, req)

	fqn := a.id.String()

	op, err := a.gcpClient.Projects.Locations.RegionalEndpoints.Create(a.id.Parent().String(), req).RegionalEndpointId(a.id.ID()).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("creating networkconnectivity regionalendpoint %s: %w", fqn, err)
	}
	if err := a.waitForOperation(ctx, op); err != nil {
		return fmt.Errorf("waiting for create of regionalendpoint %q: %w", fqn, err)
	}

	log.V(2).Info("successfully created networkconnectivity regionalendpoint in gcp", "name", a.id)

	created, err := a.gcpClient.Projects.Locations.RegionalEndpoints.Get(fqn).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("getting created regionalendpoint %q: %w", fqn, err)
	}

	resourceID := lastComponent(created.Name)
	if err := unstructured.SetNestedField(u.Object, resourceID, "spec", "resourceID"); err != nil {
		return fmt.Errorf("setting spec.resourceID: %w", err)
	}
	var createdPB *pb.RegionalEndpoint
	if err := convertAPIToProto(created, &createdPB); err != nil {
		return err
	}
	status := &krm.NetworkConnectivityRegionalEndpointStatus{}
	status.ObservedState = NetworkConnectivityRegionalEndpointObservedState_FromProto(mapCtx, createdPB)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *regionalEndpointAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("Network Connectivity regional endpoint does not support update", "name", a.id)
	return nil
}

func (a *regionalEndpointAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.NetworkConnectivityRegionalEndpoint{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(NetworkConnectivityRegionalEndpointSpec_FromProto(mapCtx, a.actual))
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
	u.SetGroupVersionKind(krm.NetworkConnectivityRegionalEndpointGVK)
	u.Object = uObj
	return u, nil
}

// Delete implements the Adapter interface.
func (a *regionalEndpointAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting networkconnectivity regionalendpoint", "name", a.id)
	fqn := a.id.String()
	op, err := a.gcpClient.Projects.Locations.ServiceConnectionPolicies.Delete(fqn).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent networkconnectivity regionalendpoint, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting networkconnectivity regionalendpoint %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully deleted networkconnectivity regionalendpoint", "name", a.id)

	if err := a.waitForOperation(ctx, op); err != nil {
		return false, fmt.Errorf("waiting delete networkconnectivity regionalendpoint %s: %w", a.id, err)
	}
	return true, nil
}

func (a *regionalEndpointAdapter) waitForOperation(ctx context.Context, op *api.GoogleLongrunningOperation) error {
	for {
		if err := ctx.Err(); err != nil {
			return err
		}

		latest, err := a.gcpClient.Projects.Locations.Operations.Get(op.Name).Context(ctx).Do()
		if err != nil {
			return fmt.Errorf("getting operation %q: %w", op.Name, err)
		}

		if latest.Done {
			return nil
		}

		time.Sleep(2 * time.Second)
	}
}
