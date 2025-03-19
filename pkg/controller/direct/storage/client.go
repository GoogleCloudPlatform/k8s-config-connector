in.proto.service.definition: option (google.api.default_host) = "storagetransfer.googleapis.com";
option (google.api.oauth_scopes) =
    "https://www.googleapis.com/auth/cloud-platform";

// Creates a new StorageLayout resource.
rpc CreateStorageLayout(CreateStorageLayoutRequest)
    returns (google.longrunning.Operation) {
  option (google.api.http) = {
    post: "/v2/{parent=projects/*/locations/*}/storageLayouts"
    body: "storage_layout"
  };
  option (google.api.method_signature) = "parent,storage_layout";
  option (google.longrunning.operation_info) = {
    response_type: "StorageLayout"
    metadata_type: "google.protobuf.Empty"
  };
}

// Gets details of a single StorageLayout.
rpc GetStorageLayout(GetStorageLayoutRequest) returns (StorageLayout) {
  option (google.api.http) = {
    get: "/v2/{name=projects/*/locations/*/storageLayouts/*}"
  };
  option (google.api.method_signature) = "name";
}

// Lists the existing StorageLayouts for the project.
rpc ListStorageLayouts(ListStorageLayoutsRequest)
    returns (ListStorageLayoutsResponse) {
  option (google.api.http) = {
    get: "/v2/{parent=projects/*/locations/*}/storageLayouts"
  };
  option (google.api.method_signature) = "parent";
}

// Updates a StorageLayout.
rpc UpdateStorageLayout(UpdateStorageLayoutRequest)
    returns (google.longrunning.Operation) {
  option (google.api.http) = {
    patch: "/v2/{storage_layout.name=projects/*/locations/*/storageLayouts/*}"
    body: "storage_layout"
  };
  option (google.longrunning.operation_info) = {
    response_type: "StorageLayout"
    metadata_type: "google.protobuf.Empty"
  };
}

// Deletes a StorageLayout.
rpc DeleteStorageLayout(DeleteStorageLayoutRequest)
    returns (google.longrunning.Operation) {
  option (google.api.http) = {
    delete: "/v2/{name=projects/*/locations/*/storageLayouts/*}"
  };
  option (google.longrunning.operation_info) = {
    response_type: "google.protobuf.Empty"
    metadata_type: "google.protobuf.Empty"
  };
}

// Creates a new StorageSystem resource.
rpc CreateStorageSystem(CreateStorageSystemRequest)
    returns (google.longrunning.Operation) {
  option (google.api.http) = {
    post: "/v2/{parent=projects/*/locations/*}/storageSystems"
    body: "storage_system"
  };
  option (google.api.method_signature) = "parent,storage_system";
  option (google.longrunning.operation_info) = {
    response_type: "StorageSystem"
    metadata_type: "google.protobuf.Empty"
  };
}

// Gets details of a single StorageSystem.
rpc GetStorageSystem(GetStorageSystemRequest) returns (StorageSystem) {
  option (google.api.http) = {
    get: "/v2/{name=projects/*/locations/*/storageSystems/*}"
  };
  option (google.api.method_signature) = "name";
}

// Lists the existing StorageSystems for the project.
rpc ListStorageSystems(ListStorageSystemsRequest)
    returns (ListStorageSystemsResponse) {
  option (google.api.http) = {
    get: "/v2/{parent=projects/*/locations/*}/storageSystems"
  };
  option (google.api.method_signature) = "parent";
}

// Updates a StorageSystem.
rpc UpdateStorageSystem(UpdateStorageSystemRequest)
    returns (google.longrunning.Operation) {
  option (google.api.http) = {
    patch: "/v2/{storage_system.name=projects/*/locations/*/storageSystems/*}"
    body: "storage_system"
  };
  option (google.longrunning.operation_info) = {
    response_type: "StorageSystem"
    metadata_type: "google.protobuf.Empty"
  };
}

// Deletes a StorageSystem.
rpc DeleteStorageSystem(DeleteStorageSystemRequest)
    returns (google.longrunning.Operation) {
  option (google.api.http) = {
    delete: "/v2/{name=projects/*/locations/*/storageSystems/*}"
  };
  option (google.longrunning.operation_info) = {
    response_type: "google.protobuf.Empty"
    metadata_type: "google.protobuf.Empty"
  };
}

// Creates a new GroupMapping resource.
rpc CreateGroupMapping(CreateGroupMappingRequest)
    returns (google.longrunning.Operation) {
  option (google.api.http) = {
    post: "/v2/{parent=projects/*/locations/*}/groupMappings"
    body: "group_mapping"
  };
  option (google.api.method_signature) = "parent,group_mapping";
  option (google.longrunning.operation_info) = {
    response_type: "GroupMapping"
    metadata_type: "google.protobuf.Empty"
  };
}

// Gets details of a single GroupMapping.
rpc GetGroupMapping(GetGroupMappingRequest) returns (GroupMapping) {
  option (google.api.http) = {
    get: "/v2/{name=projects/*/locations/*/groupMappings/*}"
  };
  option (google.api.method_signature) = "name";
}

// Lists the existing GroupMappings for the project.
rpc ListGroupMappings(ListGroupMappingsRequest)
    returns (ListGroupMappingsResponse) {
  option (google.api.http) = {
    get: "/v2/{parent=projects/*/locations/*}/groupMappings"
  };
  option (google.api.method_signature) = "parent";
}

// Updates a GroupMapping.
rpc UpdateGroupMapping(UpdateGroupMappingRequest)
    returns (google.longrunning.Operation) {
  option (google.api.http) = {
    patch: "/v2/{group_mapping.name=projects/*/locations/*/groupMappings/*}"
    body: "group_mapping"
  };
  option (google.longrunning.operation_info) = {
    response_type: "GroupMapping"
    metadata_type: "google.protobuf.Empty"
  };
}

// Deletes a GroupMapping.
rpc DeleteGroupMapping(DeleteGroupMappingRequest)
    returns (google.longrunning.Operation) {
  option (google.api.http) = {
    delete: "/v2/{name=projects/*/locations/*/groupMappings/*}"
  };
  option (google.longrunning.operation_info) = {
    response_type: "google.protobuf.Empty"
    metadata_type: "google.protobuf.Empty"
  };
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

// +tool:controller-client
// proto.service: google.storage.control.v2.StorageControl

package storagecontrol

import (
	"context"
	"fmt"

	api "cloud.google.com/go/storagecontrol/apiv2"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
)

type gcpClient struct {
	config *config.ControllerConfig
}

func newGCPClient(ctx context.Context, config *config.ControllerConfig) (*gcpClient, error) {
	gcpClient := &gcpClient{
		config: config,
	}
	return gcpClient, nil
}

func (m *gcpClient) newStorageControlClient(ctx context.Context) (*api.Client, error) {
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	client, err := api.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building storagecontrol client: %w", err)
	}
	return client, err
}


