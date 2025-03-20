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
// proto.service: google.storage.control.v2.StorageControl
// proto.message: google.storage.control.v2.Folder
// crd.type: StorageFolder
// crd.version: v1alpha1

package storagecontrol

import (
	"context"
	"fmt"

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
	registry.RegisterModel(krm.StorageFolderGVK, NewFolderModel)
}

func NewFolderModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelFolder{config: *config}, nil
}

var _ directbase.Model = &modelFolder{}

type modelFolder struct {
	config config.ControllerConfig
}

func (m *modelFolder) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.StorageFolder{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewFolderIdentity(ctx, reader, obj)
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

	return &FolderAdapter{
		id:        id,
		gcpClient: storageControlClient,
		desired:   obj,
		reader:    reader,
	}, nil
}

func (m *modelFolder) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type FolderAdapter struct {
	id        *krm.FolderIdentity
	gcpClient *gcp.StorageControlClient
	desired   *krm.StorageFolder
	actual    *pb.Folder
	reader    client.Reader
}

var _ directbase.Adapter = &FolderAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *FolderAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting Folder", "name", a.id)

	// The format of the request name is verified to use the following format.
	// Reference: https://cloud.google.com/storage/docs/create-folders#storage-create-folders-go
	requestName := fmt.Sprintf("projects/_/buckets/%s/folders/%s", a.id.Parent().BucketName, a.id.ID())
	req := &pb.GetFolderRequest{Name: requestName}
	folderpb, err := a.gcpClient.GetFolder(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting Folder %q: %w", a.id, err)
	}

	a.actual = folderpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *FolderAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating Folder", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := StorageFolderSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	// The format of the request name is verified to use the following format.
	// Reference: https://cloud.google.com/storage/docs/create-folders#storage-create-folders-go
	parent := fmt.Sprintf("projects/_/buckets/%s", a.id.Parent().BucketName)

	// folderID must end with a slash
	req := &pb.CreateFolderRequest{
		Parent:   parent,
		FolderId: a.id.ID() + "/",
		Folder:   resource,
	}
	created, err := a.gcpClient.CreateFolder(ctx, req)
	if err != nil {
		return fmt.Errorf("creating Folder %s: %w", a.id, err)
	}

	log.V(2).Info("successfully created Folder", "name", a.id)

	status := &krm.StorageFolderStatus{}
	status.ObservedState = StorageFolderObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *FolderAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating SQLInstance", "desired", a.desired)

	status := &krm.StorageFolderStatus{}
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *FolderAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.StorageFolder{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(StorageFolderSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.actual.Name)
	u.SetGroupVersionKind(krm.StorageFolderGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *FolderAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting Folder", "name", a.id)

	// The format of the request name is verified to use the following format.
	// Reference: https://cloud.google.com/storage/docs/create-folders#storage-create-folders-go
	requestName := fmt.Sprintf("projects/_/buckets/%s/folders/%s", a.id.Parent().BucketName, a.id.ID())
	req := &pb.DeleteFolderRequest{Name: requestName}
	err := a.gcpClient.DeleteFolder(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent Folder, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting Folder %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted Folder", "name", a.id)
	return true, nil
}
