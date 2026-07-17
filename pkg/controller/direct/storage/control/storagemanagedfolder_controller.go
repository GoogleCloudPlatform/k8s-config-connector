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

// +tool:controller
// proto.service: google.storage.control.v2.StorageControl
// proto.message: google.storage.control.v2.ManagedFolder
// crd.type: StorageManagedFolder
// crd.version: v1alpha1

package storagecontrol

import (
	"context"
	"fmt"

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	storagev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/storage/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/storagecontrol/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	gcp "cloud.google.com/go/storage/control/apiv2"
	pb "cloud.google.com/go/storage/control/apiv2/controlpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.StorageManagedFolderGVK, NewManagedFolderModel)
}

func NewManagedFolderModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelManagedFolder{config: *config}, nil
}

var _ directbase.Model = &modelManagedFolder{}

type modelManagedFolder struct {
	config config.ControllerConfig
}

func (m *modelManagedFolder) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.StorageManagedFolder{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewManagedFolderIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Get storage GCP client
	options, err := m.config.GRPCClientOptions()
	if err != nil {
		return nil, err
	}
	storageControlClient, err := gcp.NewStorageControlClient(ctx, options...)
	if err != nil {
		return nil, err
	}

	return &ManagedFolderAdapter{
		id:        id,
		gcpClient: storageControlClient,
		desired:   obj,
		reader:    reader,
	}, nil
}

func (m *modelManagedFolder) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	id, err := krm.NewManagedFolderIdentityFromExternal(url)
	if err != nil {
		return nil, nil // Not recognized
	}

	options, err := m.config.GRPCClientOptions()
	if err != nil {
		return nil, err
	}
	storageControlClient, err := gcp.NewStorageControlClient(ctx, options...)
	if err != nil {
		return nil, err
	}

	return &ManagedFolderAdapter{
		id:        id,
		gcpClient: storageControlClient,
	}, nil
}

type ManagedFolderAdapter struct {
	id        *krm.ManagedFolderIdentity
	gcpClient *gcp.StorageControlClient
	desired   *krm.StorageManagedFolder
	actual    *pb.ManagedFolder
	reader    client.Reader
}

var _ directbase.Adapter = &ManagedFolderAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *ManagedFolderAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ManagedFolder", "name", a.id)

	requestName := fmt.Sprintf("projects/_/buckets/%s/managedFolders/%s/", a.id.Parent().BucketName, a.id.ID())
	req := &pb.GetManagedFolderRequest{Name: requestName}
	managedFolderpb, err := a.gcpClient.GetManagedFolder(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ManagedFolder %q: %w", a.id, err)
	}

	a.actual = managedFolderpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *ManagedFolderAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating ManagedFolder", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := StorageManagedFolderSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	parent := fmt.Sprintf("projects/_/buckets/%s", a.id.Parent().BucketName)

	req := &pb.CreateManagedFolderRequest{
		Parent:          parent,
		ManagedFolderId: a.id.ID() + "/",
		ManagedFolder:   resource,
	}
	created, err := a.gcpClient.CreateManagedFolder(ctx, req)
	if err != nil {
		return fmt.Errorf("creating ManagedFolder %s: %w", a.id, err)
	}

	log.V(2).Info("successfully created ManagedFolder", "name", a.id)

	status := &krm.StorageManagedFolderStatus{}
	status.ObservedState = StorageManagedFolderObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *ManagedFolderAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating ManagedFolder", "desired", a.desired)

	// ManagedFolder has no mutable fields that we can update.
	// But we still must populate status from actual.
	status := &krm.StorageManagedFolderStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = StorageManagedFolderObservedState_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *ManagedFolderAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.StorageManagedFolder{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(StorageManagedFolderSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.StorageFolderParent = &krm.StorageFolderParent{
		ProjectRef: &refsv1beta1.ProjectRef{
			External: a.id.Parent().ProjectID,
		},
		StorageBucketRef: &storagev1beta1.StorageBucketRef{
			External: a.id.Parent().String(),
		},
	}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.ID())
	u.SetGroupVersionKind(krm.StorageManagedFolderGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *ManagedFolderAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ManagedFolder", "name", a.id)

	requestName := fmt.Sprintf("projects/_/buckets/%s/managedFolders/%s/", a.id.Parent().BucketName, a.id.ID())
	req := &pb.DeleteManagedFolderRequest{Name: requestName}
	err := a.gcpClient.DeleteManagedFolder(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent ManagedFolder, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting ManagedFolder %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted ManagedFolder", "name", a.id)
	return true, nil
}
