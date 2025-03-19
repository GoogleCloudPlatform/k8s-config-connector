in.proto.message.definition: message Folder {
  option (google.api.resource) = {
    type: "storage.googleapis.com/Folder"
    pattern: "projects/{project}/buckets/{bucket}/folders/{folder}"
  };

  // Immutable. Full path name of the folder.
  string name = 1 [(google.api.field_behavior) = IMMUTABLE];

  // User specified attributes.
  map<string, string> custom_attributes = 8;

  // Immutable. Metageneration of the Folder. The field is only populated after
  // the Folder creation.
  int64 metageneration = 2
      [(google.api.field_behavior) = IMMUTABLE, (google.api.field_info) = {
        output_only_behavior: IMMUTABLE
      }];

  // Output only. Resource creation timestamp.
  google.protobuf.Timestamp create_time = 6
      [(google.api.field_behavior) = OUTPUT_ONLY];

  // Output only. Resource update timestamp. It may be either:
  // 1) last time the folder metadata was changed;
  // OR
  // 2) if custom attributes are present, the creation time of the
  //    representation object, i.e. the file whose name is the folder name.
  google.protobuf.Timestamp update_time = 7
      [(google.api.field_behavior) = OUTPUT_ONLY];

  // The managed folder that this folder represents. Can only be set during
  // folder creation, and cannot be updated.
  string managed_folder = 9
      [(google.api.resource_reference) = {
        type: "storagetransfer.googleapis.com/ManagedFolder"
      }];

  // Immutable. User-provided id of the folder, part of name.
  string folder_id = 10 [(google.api.field_behavior) = IMMUTABLE];
}
in.proto.service.definition: service StorageControl {
  option (google.api.default_host) = "storagecontrol.googleapis.com";
  option (google.api.oauth_scopes) = "https://www.googleapis.com/auth/cloud-platform";

  // Creates a new folder.
  rpc CreateFolder(CreateFolderRequest) returns (Folder) {
    option (google.api.http) = {
      post: "/v2/{parent=projects/*/buckets/*}/folders"
      body: "folder"
    };
    option (google.api.method_signature) = "parent,folder";
  }

  // Permanently deletes a folder, and all sub-folders and objects within it.
  // Deletion of folders with underlying objects will only be successful
  // if "force=true" is provided in the request.
  rpc DeleteFolder(DeleteFolderRequest) returns (google.protobuf.Empty) {
    option (google.api.http) = {
      delete: "/v2/{name=projects/*/buckets/*/folders/*}"
    };
    option (google.api.method_signature) = "name";
  }

  // Returns metadata of a folder.
  rpc GetFolder(GetFolderRequest) returns (Folder) {
    option (google.api.http) = {
      get: "/v2/{name=projects/*/buckets/*/folders/*}"
    };
    option (google.api.method_signature) = "name";
  }

  // Returns a (possibly paginated) list of folders for a bucket.
  // The `name` field will be the folder name and there will be no
  // sub-folders, sub-folders can be found via the `prefixes` field.
  rpc ListFolders(ListFoldersRequest) returns (ListFoldersResponse) {
    option (google.api.http) = {
      get: "/v2/{parent=projects/*/buckets/*}/folders"
    };
    option (google.api.method_signature) = "parent";
  }

  // Renames a folder.
  rpc RenameFolder(RenameFolderRequest) returns (Folder) {
    option (google.api.http) = {
      post: "/v2/{name=projects/*/buckets/*/folders/*}:rename"
      body: "*"
    };
    option (google.api.method_signature) = "name,destination_folder_id";
  }
}
out: // Copyright 2024 Google LLC
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

package storage

import (
	"context"
	"fmt"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/storage/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	gcp "cloud.google.com/go/storagecontrol/apiv2"
	pb "cloud.google.com/go/storagecontrol/apiv2/storagecontrolpb"

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
	gcpClient, err := newGCPClient(ctx, &m.config)
	if err != nil {
		return nil, err
	}
	storageControlClient, err := gcpClient.newStorageControlClient(ctx)
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
	gcpClient *gcp.Client
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

	req := &pb.GetFolderRequest{Name: a.id.String()}
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

	req := &pb.CreateFolderRequest{
		Parent:   a.id.Parent().String(),
		FolderId: a.id.ID(),
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
	return fmt.Errorf("update is not implemented")
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

	req := &pb.DeleteFolderRequest{Name: a.id.String()}
	_, err := a.gcpClient.DeleteFolder(ctx, req)
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


