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
// proto.message: google.cloud.datastream.v1.ConnectionProfile
// crd.type: DatastreamConnectionProfile
// crd.version: v1alpha1

package datastream

import (
	"context"
	"fmt"
	"reflect"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/datastream/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	refsv1beta1secret "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1/secret"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	gcp "cloud.google.com/go/datastream/apiv1"
	pb "cloud.google.com/go/datastream/apiv1/datastreampb"
	"google.golang.org/genproto/protobuf/field_mask"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.DatastreamConnectionProfileGVK, NewConnectionProfileModel)
}

func NewConnectionProfileModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelConnectionProfile{config: *config}, nil
}

var _ directbase.Model = &modelConnectionProfile{}

type modelConnectionProfile struct {
	config config.ControllerConfig
}

func (m *modelConnectionProfile) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.DatastreamConnectionProfile{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewConnectionProfileIdentity(ctx, reader, obj)
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
	return &ConnectionProfileAdapter{
		id:        id,
		gcpClient: datastreamClient,
		desired:   obj,
		reader:    reader,
	}, nil
}

func (m *modelConnectionProfile) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type ConnectionProfileAdapter struct {
	id        *krm.ConnectionProfileIdentity
	gcpClient *gcp.Client
	desired   *krm.DatastreamConnectionProfile
	actual    *pb.ConnectionProfile
	reader    client.Reader
}

var _ directbase.Adapter = &ConnectionProfileAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *ConnectionProfileAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ConnectionProfile", "name", a.id)

	req := &pb.GetConnectionProfileRequest{Name: a.id.String()}
	connectionprofilepb, err := a.gcpClient.GetConnectionProfile(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ConnectionProfile %q: %w", a.id, err)
	}

	a.actual = connectionprofilepb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *ConnectionProfileAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating ConnectionProfile", "name", a.id)

	if err := a.normalizeReferenceFields(ctx); err != nil {
		return err
	}

	mapCtx := &direct.MapContext{}
	desired := a.desired.DeepCopy()
	resource := DatastreamConnectionProfileSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &pb.CreateConnectionProfileRequest{
		Parent:              a.id.Parent().String(),
		ConnectionProfileId: a.id.ID(),
		ConnectionProfile:   resource,
	}
	op, err := a.gcpClient.CreateConnectionProfile(ctx, req)
	if err != nil {
		return fmt.Errorf("creating ConnectionProfile %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("ConnectionProfile %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created ConnectionProfile", "name", a.id)

	status := &krm.DatastreamConnectionProfileStatus{}
	status.ObservedState = DatastreamConnectionProfileObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *ConnectionProfileAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating ConnectionProfile", "name", a.id)

	if err := a.normalizeReferenceFields(ctx); err != nil {
		return err
	}

	mapCtx := &direct.MapContext{}
	desired := a.desired.DeepCopy()
	resource := DatastreamConnectionProfileSpec_ToProto(mapCtx, &desired.Spec)
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
	// TODO: handle other fields

	var updated *pb.ConnectionProfile
	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		updated = a.actual
	} else {
		resource.Name = a.id.String() // we need to set the name so that GCP API can identify the resource
		req := &pb.UpdateConnectionProfileRequest{
			ConnectionProfile: resource,
			UpdateMask:        &field_mask.FieldMask{Paths: paths},
		}
		op, err := a.gcpClient.UpdateConnectionProfile(ctx, req)
		if err != nil {
			return fmt.Errorf("updating ConnectionProfile %s: %w", a.id, err)
		}
		updated, err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("ConnectionProfile %s waiting update: %w", a.id, err)
		}
		log.V(2).Info("successfully updated ConnectionProfile", "name", a.id)
	}

	status := &krm.DatastreamConnectionProfileStatus{}
	status.ObservedState = DatastreamConnectionProfileObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *ConnectionProfileAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.DatastreamConnectionProfile{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(DatastreamConnectionProfileSpec_FromProto(mapCtx, a.actual))
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
	u.SetGroupVersionKind(krm.DatastreamConnectionProfileGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *ConnectionProfileAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ConnectionProfile", "name", a.id)

	req := &pb.DeleteConnectionProfileRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteConnectionProfile(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent ConnectionProfile, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting ConnectionProfile %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted ConnectionProfile", "name", a.id)

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete ConnectionProfile %s: %w", a.id, err)
	}
	return true, nil
}

func (a *ConnectionProfileAdapter) normalizeReferenceFields(ctx context.Context) error {
	obj := a.desired

	if obj.Spec.OracleProfile != nil && obj.Spec.OracleProfile.OracleASMConfig != nil && obj.Spec.OracleProfile.OracleASMConfig.SecretRef != nil {
		if err := refsv1beta1secret.NormalizedSecret(ctx, obj.Spec.OracleProfile.OracleASMConfig.SecretRef, a.reader, obj.GetNamespace()); err != nil {
			return err
		}
	}
	if obj.Spec.ForwardSSHConnectivity != nil && obj.Spec.ForwardSSHConnectivity.SecretRef != nil {
		if err := refsv1beta1secret.NormalizedSecret(ctx, obj.Spec.ForwardSSHConnectivity.SecretRef, a.reader, obj.GetNamespace()); err != nil {
			return err
		}
	}
	if obj.Spec.MySQLProfile != nil && obj.Spec.MySQLProfile.SecretRef != nil {
		if err := refsv1beta1secret.NormalizedSecret(ctx, obj.Spec.MySQLProfile.SecretRef, a.reader, obj.GetNamespace()); err != nil {
			return err
		}
	}
	if obj.Spec.OracleProfile != nil && obj.Spec.OracleProfile.SecretRef != nil {
		if err := refsv1beta1secret.NormalizedSecret(ctx, obj.Spec.OracleProfile.SecretRef, a.reader, obj.GetNamespace()); err != nil {
			return err
		}
	}
	if obj.Spec.OracleProfile != nil && obj.Spec.OracleProfile.SecreteManagerSecretRef != nil {
		if _, err := refs.ResolveSecretManagerSecretRef(ctx, a.reader, obj, obj.Spec.OracleProfile.SecreteManagerSecretRef); err != nil {
			return err
		}
	}
	// TODO: PostgresqlProfile is not implemented yet
	if obj.Spec.SQLServerProfile != nil && obj.Spec.SQLServerProfile.SecretRef != nil {
		if err := refsv1beta1secret.NormalizedSecret(ctx, obj.Spec.SQLServerProfile.SecretRef, a.reader, obj.GetNamespace()); err != nil {
			return err
		}
	}
	if obj.Spec.PrivateConnectivity != nil && obj.Spec.PrivateConnectivity.PrivateConnectionRef != nil {
		if _, err := obj.Spec.PrivateConnectivity.PrivateConnectionRef.NormalizedExternal(ctx, a.reader, obj.GetNamespace()); err != nil {
			return err
		}
	}

	return nil
}
