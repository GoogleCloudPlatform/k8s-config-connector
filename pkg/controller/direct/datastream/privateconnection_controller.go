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
// proto.message: google.cloud.datastream.v1.PrivateConnection
// crd.type: DatastreamPrivateConnection
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
	registry.RegisterModel(krm.DatastreamPrivateConnectionGVK, NewPrivateConnectionModel)
}

func NewPrivateConnectionModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelPrivateConnection{config: *config}, nil
}

var _ directbase.Model = &modelPrivateConnection{}

type modelPrivateConnection struct {
	config config.ControllerConfig
}

func (m *modelPrivateConnection) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.DatastreamPrivateConnection{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewPrivateConnectionIdentity(ctx, reader, obj)
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
	return &PrivateConnectionAdapter{
		id:        id,
		gcpClient: datastreamClient,
		desired:   obj,
		reader:    reader,
	}, nil
}

func (m *modelPrivateConnection) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type PrivateConnectionAdapter struct {
	id        *krm.PrivateConnectionIdentity
	gcpClient *gcp.Client
	desired   *krm.DatastreamPrivateConnection
	actual    *pb.PrivateConnection
	reader    client.Reader
}

var _ directbase.Adapter = &PrivateConnectionAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *PrivateConnectionAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting PrivateConnection", "name", a.id)

	req := &pb.GetPrivateConnectionRequest{Name: a.id.String()}
	privateconnectionpb, err := a.gcpClient.GetPrivateConnection(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting PrivateConnection %q: %w", a.id, err)
	}

	a.actual = privateconnectionpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *PrivateConnectionAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating PrivateConnection", "name", a.id)

	if err := a.normalizeReferenceFields(ctx); err != nil {
		return err
	}

	mapCtx := &direct.MapContext{}
	desired := a.desired.DeepCopy()
	resource := DatastreamPrivateConnectionSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &pb.CreatePrivateConnectionRequest{
		Parent:              a.id.Parent().String(),
		PrivateConnectionId: a.id.ID(),
		PrivateConnection:   resource,
	}
	op, err := a.gcpClient.CreatePrivateConnection(ctx, req)
	if err != nil {
		return fmt.Errorf("creating PrivateConnection %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("PrivateConnection %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created PrivateConnection", "name", a.id)

	status := &krm.DatastreamPrivateConnectionStatus{}
	status.ObservedState = DatastreamPrivateConnectionObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *PrivateConnectionAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating PrivateConnection", "name", a.id)

	if err := a.normalizeReferenceFields(ctx); err != nil {
		return err
	}

	mapCtx := &direct.MapContext{}
	desired := a.desired.DeepCopy()
	resource := DatastreamPrivateConnectionSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	paths := []string{}
	if desired.Spec.DisplayName != nil && !reflect.DeepEqual(resource.DisplayName, a.actual.DisplayName) {
		paths = append(paths, "display_name")
	}
	if desired.Spec.Labels != nil && !reflect.DeepEqual(resource.Labels, a.actual.Labels) {
		paths = append(paths, "labels")
	}
	if desired.Spec.VPCPeeringConfig != nil && !reflect.DeepEqual(resource.VpcPeeringConfig, a.actual.VpcPeeringConfig) {
		paths = append(paths, "vpc_peering_config")
	}

	if len(paths) != 0 {
		return fmt.Errorf("updating PrivateConnection is not supported, fields: %v", paths)
	}

	// still need to update status (in the event of acquiring an existing resource)
	status := &krm.DatastreamPrivateConnectionStatus{}
	status.ObservedState = DatastreamPrivateConnectionObservedState_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *PrivateConnectionAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.DatastreamPrivateConnection{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(DatastreamPrivateConnectionSpec_FromProto(mapCtx, a.actual))
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
	u.SetGroupVersionKind(krm.DatastreamPrivateConnectionGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *PrivateConnectionAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting PrivateConnection", "name", a.id)

	req := &pb.DeletePrivateConnectionRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeletePrivateConnection(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent PrivateConnection, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting PrivateConnection %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted PrivateConnection", "name", a.id)

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete PrivateConnection %s: %w", a.id, err)
	}
	return true, nil
}

func (a *PrivateConnectionAdapter) normalizeReferenceFields(ctx context.Context) error {
	obj := a.desired

	if obj.Spec.VPCPeeringConfig != nil && obj.Spec.VPCPeeringConfig.NetworkRef != nil {
		if err := obj.Spec.VPCPeeringConfig.NetworkRef.Normalize(ctx, a.reader, obj); err != nil {
			return err
		}
	}

	return nil
}
