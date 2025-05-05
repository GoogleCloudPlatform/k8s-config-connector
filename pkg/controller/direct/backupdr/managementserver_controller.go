// Copyright 2025 Google LLC
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
// proto.service: google.cloud.backupdr.v1.BackupDR
// proto.message: google.cloud.backupdr.v1.ManagementServer
// crd.type: BackupDRManagementServer
// crd.version: v1alpha1

package backupdr

import (
	"context"
	"fmt"
	"reflect"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/backupdr/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	gcp "cloud.google.com/go/backupdr/apiv1"
	pb "cloud.google.com/go/backupdr/apiv1/backupdrpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.BackupDRManagementServerGVK, NewManagementServerModel)
}

func NewManagementServerModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelManagementServer{config: *config}, nil
}

var _ directbase.Model = &modelManagementServer{}

type modelManagementServer struct {
	config config.ControllerConfig
}

func (m *modelManagementServer) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.BackupDRManagementServer{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewManagementServerIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Get backupdr GCP client
	gcpClient, err := newGCPClient(ctx, &m.config)
	if err != nil {
		return nil, err
	}
	backupDRClient, err := gcpClient.newBackupDRClient(ctx)
	if err != nil {
		return nil, err
	}
	return &ManagementServerAdapter{
		id:        id,
		gcpClient: backupDRClient,
		desired:   obj,
		reader:    reader,
	}, nil
}

func (m *modelManagementServer) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type ManagementServerAdapter struct {
	id        *krm.ManagementServerIdentity
	gcpClient *gcp.Client
	desired   *krm.BackupDRManagementServer
	actual    *pb.ManagementServer
	reader    client.Reader
}

var _ directbase.Adapter = &ManagementServerAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *ManagementServerAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ManagementServer", "name", a.id)

	req := &pb.GetManagementServerRequest{Name: a.id.String()}
	managementserverpb, err := a.gcpClient.GetManagementServer(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ManagementServer %q: %w", a.id, err)
	}

	a.actual = managementserverpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *ManagementServerAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating ManagementServer", "name", a.id)

	if err := a.normalizeReferenceFields(ctx); err != nil {
		return err
	}

	mapCtx := &direct.MapContext{}
	desired := a.desired.DeepCopy()
	resource := BackupDRManagementServerSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &pb.CreateManagementServerRequest{
		Parent:             a.id.Parent().String(),
		ManagementServerId: a.id.ID(),
		ManagementServer:   resource,
	}
	op, err := a.gcpClient.CreateManagementServer(ctx, req)
	if err != nil {
		return fmt.Errorf("creating ManagementServer %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("ManagementServer %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created ManagementServer", "name", a.id)

	status := &krm.BackupDRManagementServerStatus{}
	status.ObservedState = BackupDRManagementServerObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *ManagementServerAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating ManagementServer", "name", a.id)

	if err := a.normalizeReferenceFields(ctx); err != nil {
		return err
	}

	mapCtx := &direct.MapContext{}
	desired := a.desired.DeepCopy()
	resource := BackupDRManagementServerSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	paths := []string{}
	if desired.Spec.Description != nil && !reflect.DeepEqual(resource.Description, a.actual.Description) {
		paths = append(paths, "description")
	}
	if desired.Spec.Labels != nil && !reflect.DeepEqual(resource.Labels, a.actual.Labels) {
		paths = append(paths, "labels")
	}
	if desired.Spec.Networks != nil && !reflect.DeepEqual(resource.Networks, a.actual.Networks) {
		paths = append(paths, "networks")
	}
	if desired.Spec.Type != nil && !reflect.DeepEqual(resource.Type, a.actual.Type) {
		paths = append(paths, "type")
	}

	if len(paths) != 0 {
		return fmt.Errorf("update ManagementServer is not supported, fields: %v", paths)
	}

	status := &krm.BackupDRManagementServerStatus{}
	status.ObservedState = BackupDRManagementServerObservedState_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *ManagementServerAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.BackupDRManagementServer{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(BackupDRManagementServerSpec_FromProto(mapCtx, a.actual))
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
	u.SetGroupVersionKind(krm.BackupDRManagementServerGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *ManagementServerAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ManagementServer", "name", a.id)

	req := &pb.DeleteManagementServerRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteManagementServer(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent ManagementServer, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting ManagementServer %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted ManagementServer", "name", a.id)

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete ManagementServer %s: %w", a.id, err)
	}
	return true, nil
}

func (a *ManagementServerAdapter) normalizeReferenceFields(ctx context.Context) error {
	obj := a.desired
	for i := range obj.Spec.Networks {
		if obj.Spec.Networks[i].NetworkRef != nil {
			if err := obj.Spec.Networks[i].NetworkRef.Normalize(ctx, a.reader, obj); err != nil {
				return err
			}
		}
	}
	return nil
}
