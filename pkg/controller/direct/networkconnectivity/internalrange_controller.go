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
// proto.message: google.cloud.networkconnectivity.v1.InternalRange
// crd.type: NetworkConnectivityInternalRange
// crd.version: v1alpha1

package networkconnectivity

import (
	"context"
	"fmt"
	"reflect"
	"strings"
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
	registry.RegisterModel(krm.NetworkConnectivityInternalRangeGVK, NewInternalRangeModel)
}

func NewInternalRangeModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &internalRangeModel{config: *config}, nil
}

var _ directbase.Model = &internalRangeModel{}

type internalRangeModel struct {
	config config.ControllerConfig
}

func (m *internalRangeModel) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.NetworkConnectivityInternalRange{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewInternalRangeIdentity(ctx, reader, obj)
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
	return &internalRangeAdapter{
		gcpClient: client,
		id:        id,
		desired:   obj,
	}, nil
}

func (m *internalRangeModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type internalRangeAdapter struct {
	gcpClient *api.Service
	id        *krm.InternalRangeIdentity
	desired   *krm.NetworkConnectivityInternalRange
	actual    *pb.InternalRange
}

var _ directbase.Adapter = &internalRangeAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *internalRangeAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting networkconnectivity internalrange", "name", a.id)
	fqn := a.id.String()
	actual, err := a.gcpClient.Projects.Locations.InternalRanges.Get(fqn).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting networkconnectivity internalrange %q from gcp: %w", a.id.String(), err)
	}

	if err := convertAPIToProto(actual, &a.actual); err != nil {
		return false, err
	}

	return true, nil
}

func (a *internalRangeAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	u := createOp.GetUnstructured()

	log := klog.FromContext(ctx)
	log.V(2).Info("creating networkconnectivity internalrange", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := NetworkConnectivityInternalRangeSpec_ToProto(mapCtx, &a.desired.DeepCopy().Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	req := &api.InternalRange{}
	err := convertProtoToAPI(desired, req)

	fqn := a.id.String()

	op, err := a.gcpClient.Projects.Locations.InternalRanges.Create(a.id.Parent().String(), req).InternalRangeId(a.id.ID()).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("creating networkconnectivity internalrange %s: %w", fqn, err)
	}
	if err := a.waitForOperation(ctx, op); err != nil {
		return fmt.Errorf("waiting for create of internalrange %q: %w", fqn, err)
	}

	log.V(2).Info("successfully created networkconnectivity internalrange in gcp", "name", a.id)

	created, err := a.gcpClient.Projects.Locations.InternalRanges.Get(fqn).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("getting created internalrange %q: %w", fqn, err)
	}

	resourceID := lastComponent(created.Name)
	if err := unstructured.SetNestedField(u.Object, resourceID, "spec", "resourceID"); err != nil {
		return fmt.Errorf("setting spec.resourceID: %w", err)
	}
	var createdPB *pb.InternalRange
	if err := convertAPIToProto(created, &createdPB); err != nil {
		return err
	}
	status := &krm.NetworkConnectivityInternalRangeStatus{}
	status.ObservedState = NetworkConnectivityInternalRangeObservedState_FromProto(mapCtx, createdPB)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *internalRangeAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating networkconnectivity internalrange", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := NetworkConnectivityInternalRangeSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	paths := []string{}
	if desired.Spec.Description != nil && !reflect.DeepEqual(resource.Description, a.actual.Description) {
		paths = append(paths, "description")
	}
	if desired.Spec.IPCIDRRange != nil && !reflect.DeepEqual(resource.IpCidrRange, a.actual.IpCidrRange) {
		paths = append(paths, "ipCidrRange")
	}
	if desired.Spec.Labels != nil && !reflect.DeepEqual(resource.Labels, a.actual.Labels) {
		paths = append(paths, "labels")
	}
	if desired.Spec.Peering != nil && !reflect.DeepEqual(resource.Peering, a.actual.Peering) {
		paths = append(paths, "peering")
	}
	if desired.Spec.PrefixLength != nil && !reflect.DeepEqual(resource.PrefixLength, a.actual.PrefixLength) {
		paths = append(paths, "prefixLength")
	}
	if desired.Spec.TargetCIDRRange != nil && !reflect.DeepEqual(resource.TargetCidrRange, a.actual.TargetCidrRange) {
		paths = append(paths, "targetCidrRange")
	}
	if desired.Spec.Usage != nil && !reflect.DeepEqual(resource.Usage, a.actual.Usage) {
		paths = append(paths, "usage")
	}

	if len(paths) > 0 {
		resource.Name = a.id.String() // we need to set the name so that GCP API can identify the resource
		req := &api.InternalRange{}
		if err := convertProtoToAPI(resource, req); err != nil {
			return err
		}
		fqn := a.id.String()
		op, err := a.gcpClient.Projects.Locations.InternalRanges.Patch(fqn, req).UpdateMask(strings.Join(paths, ",")).Context(ctx).Do()
		if err != nil {
			return fmt.Errorf("updating networkconnectivity internalrange %s: %w", fqn, err)
		}
		if err := a.waitForOperation(ctx, op); err != nil {
			return fmt.Errorf("waiting for update of internalrange %q: %w", fqn, err)
		}
		log.V(2).Info("successfully updated networkconnectivity internalrange", "name", fqn)
		updatedAPI, err := a.gcpClient.Projects.Locations.InternalRanges.Get(fqn).Context(ctx).Do()
		if err != nil {
			return fmt.Errorf("getting updated internalrange %q: %w", fqn, err)
		}
		if err := convertAPIToProto(updatedAPI, &a.actual); err != nil {
			return err
		}
	}

	status := &krm.NetworkConnectivityInternalRangeStatus{}
	status.ObservedState = NetworkConnectivityInternalRangeObservedState_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *internalRangeAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.NetworkConnectivityInternalRange{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(NetworkConnectivityInternalRangeSpec_FromProto(mapCtx, a.actual))
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
	u.SetGroupVersionKind(krm.NetworkConnectivityInternalRangeGVK)
	u.Object = uObj
	return u, nil
}

// Delete implements the Adapter interface.
func (a *internalRangeAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting networkconnectivity internalrange", "name", a.id)
	fqn := a.id.String()
	op, err := a.gcpClient.Projects.Locations.ServiceConnectionPolicies.Delete(fqn).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent networkconnectivity internalrange, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting networkconnectivity internalrange %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully deleted networkconnectivity internalrange", "name", a.id)

	if err := a.waitForOperation(ctx, op); err != nil {
		return false, fmt.Errorf("waiting delete networkconnectivity internalrange %s: %w", a.id, err)
	}
	return true, nil
}

func (a *internalRangeAdapter) waitForOperation(ctx context.Context, op *api.GoogleLongrunningOperation) error {
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
