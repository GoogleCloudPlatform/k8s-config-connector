 google.bigtable.admin.v2.BigtableInstanceAdmin
in.proto.service.definition: service BigtableInstanceAdmin {
  option (google.api.default_host) = "bigtableadmin.googleapis.com";
  option (google.api.oauth_scopes) =
      "https://www.googleapis.com/auth/bigtable.admin,"
      "https://www.googleapis.com/auth/bigtable.admin.instance,"
      "https://www.googleapis.com/auth/cloud-platform,"
      "https://www.googleapis.com/auth/cloud-platform.read-only";

  // Create an instance within a project.
  //
  // Note that exactly one of Cluster.serve_nodes and
  // Cluster.cluster_config.cluster_autoscaling_config can be set. If
  // serve_nodes is set to non-zero, then the cluster is manually scaled. If
  // cluster_config.cluster_autoscaling_config is non-empty, then autoscaling is
  // enabled.
  rpc CreateInstance(CreateInstanceRequest)
      returns (google.longrunning.Operation) {
    option (google.api.http) = {
      post: "/v2/{parent=projects/*}/instances"
      body: "instance"
    };
    option (google.api.method_signature) =
        "parent,instance_id,instance,clusters";
    option (google.longrunning.operation_info) = {
      response_type: "Instance"
      metadata_type: "CreateInstanceMetadata"
    };
  }

  // Gets information about an instance.
  rpc GetInstance(GetInstanceRequest) returns (Instance) {
    option (google.api.http) = { get: "/v2/{name=projects/*/instances/*}" };
    option (google.api.method_signature) = "name";
  }

  // Lists information about instances in a project.
  rpc ListInstances(ListInstancesRequest) returns (ListInstancesResponse) {
    option (google.api.http) = {
      get: "/v2/{parent=projects/*}/instances"
    };
    option (google.api.method_signature) = "parent";
  }

  // Updates an instance within a project. This method updates only the display
  // name and type for an Instance. To update other Instance properties, such as
  // labels, use PartialUpdateInstance.
  rpc UpdateInstance(Instance) returns (Instance) {
    option (google.api.http) = {
      put: "/v2/{name=projects/*/instances/*}"
      body: "*"
    };
    option (google.api.method_signature) = "name,display_name,type,labels";
  }

  // Partially updates an instance within a project. This method can modify all
  // fields of an Instance and is the preferred way to update an Instance.
  rpc PartialUpdateInstance(PartialUpdateInstanceRequest)
      returns (google.longrunning.Operation) {
    option (google.api.http) = {
      patch: "/v2/{instance.name=projects/*/instances/*}"
      body: "instance"
    };
    option (google.api.method_signature) = "instance,update_mask";
    option (google.longrunning.operation_info) = {
      response_type: "Instance"
      metadata_type: "UpdateInstanceMetadata"
    };
  }

  // Delete an instance from a project.
  rpc DeleteInstance(DeleteInstanceRequest) returns (google.protobuf.Empty) {
    option (google.api.http) = {
      delete: "/v2/{name=projects/*/instances/*}"
    };
    option (google.api.method_signature) = "name";
  }

  // Creates a cluster within an instance.
  //
  // Note that exactly one of Cluster.serve_nodes and
  // Cluster.cluster_config.cluster_autoscaling_config can be set. If
  // serve_nodes is set to non-zero, then the cluster is manually scaled. If
  // cluster_config.cluster_autoscaling_config is non-empty, then autoscaling is
  // enabled.
  rpc CreateCluster(CreateClusterRequest)
      returns (google.longrunning.Operation) {
    option (google.api.http) = {
      post: "/v2/{parent=projects/*/instances/*}/clusters"
      body: "cluster"
    };
    option (google.api.method_signature) = "parent,cluster_id,cluster";
    option (google.longrunning.operation_info) = {
      response_type: "Cluster"
      metadata_type: "CreateClusterMetadata"
    };
  }

  // Gets information about a cluster.
  rpc GetCluster(GetClusterRequest) returns (Cluster) {
    option (google.api.http) = {
      get: "/v2/{name=projects/*/instances/*/clusters/*}"
    };
    option (google.api.method_signature) = "name";
  }

  // Lists information about clusters in an instance.
  rpc ListClusters(ListClustersRequest) returns (ListClustersResponse) {
    option (google.api.http) = {
      get: "/v2/{parent=projects/*/instances/*}/clusters"
    };
    option (google.api.method_signature) = "parent";
  }

  // Updates a cluster within an instance.
  rpc UpdateCluster(Cluster) returns (google.longrunning.Operation) {
    option (google.api.http) = {
      put: "/v2/{name=projects/*/instances/*/clusters/*}"
      body: "*"
    };
    option (google.api.method_signature) = "name,location,serve_nodes";
    option (google.longrunning.operation_info) = {
      response_type: "Cluster"
      metadata_type: "UpdateClusterMetadata"
    };
  }

  // Partially updates a cluster within a project. This method is the preferred
  // way to update a Cluster.
  rpc PartialUpdateCluster(PartialUpdateClusterRequest)
      returns (google.longrunning.Operation) {
    option (google.api.http) = {
      patch: "/v2/{cluster.name=projects/*/instances/*/clusters/*}"
      body: "cluster"
    };
    option (google.longrunning.operation_info) = {
      response_type: "Cluster"
      metadata_type: "UpdateClusterMetadata"
    };
  }

  // Deletes a cluster from an instance.
  rpc DeleteCluster(DeleteClusterRequest) returns (google.protobuf.Empty) {
    option (google.api.http) = {
      delete: "/v2/{name=projects/*/instances/*/clusters/*}"
    };
    option (google.api.method_signature) = "name";
  }

  // Creates an app profile within an instance.
  rpc CreateAppProfile(CreateAppProfileRequest) returns (AppProfile) {
    option (google.api.http) = {
      post: "/v2/{parent=projects/*/instances/*}/appProfiles"
      body: "app_profile"
    };
    option (google.api.method_signature) =
        "parent,app_profile_id,app_profile";
  }

  // Gets information about an app profile.
  rpc GetAppProfile(GetAppProfileRequest) returns (AppProfile) {
    option (google.api.http) = {
      get: "/v2/{name=projects/*/instances/*/appProfiles/*}"
    };
    option (google.api.method_signature) = "name";
  }

  // Lists information about app profiles in an instance.
  rpc ListAppProfiles(ListAppProfilesRequest) returns (ListAppProfilesResponse) {
    option (google.api.http) = {
      get: "/v2/{parent=projects/*/instances/*}/appProfiles"
    };
    option (google.api.method_signature) = "parent";
  }

  // Updates an app profile within an instance.
  rpc UpdateAppProfile(UpdateAppProfileRequest)
      returns (google.longrunning.Operation) {
    option (google.api.http) = {
      patch: "/v2/{app_profile.name=projects/*/instances/*/appProfiles/*}"
      body: "app_profile"
    };
    option (google.api.method_signature) = "app_profile,update_mask";
    option (google.longrunning.operation_info) = {
      response_type: "AppProfile"
      metadata_type: "UpdateAppProfileMetadata"
    };
  }

  // Deletes an app profile from an instance.
  rpc DeleteAppProfile(DeleteAppProfileRequest) returns (google.protobuf.Empty) {
    option (google.api.http) = {
      delete: "/v2/{name=projects/*/instances/*/appProfiles/*}"
    };
    option (google.api.method_signature) = "name,ignore_warnings";
  }

  // Gets the access control policy for an instance resource. Returns an empty
  // policy if an instance exists but does not have a policy set.
  rpc GetIamPolicy(google.iam.v1.GetIamPolicyRequest)
      returns (google.iam.v1.Policy) {
    option (google.api.http) = {
      post: "/v2/{resource=projects/*/instances/*}:getIamPolicy"
      body: "*"
    };
    option (google.api.method_signature) = "resource";
  }

  // Sets the access control policy on an instance resource. Replaces any
  // existing policy.
  rpc SetIamPolicy(google.iam.v1.SetIamPolicyRequest)
      returns (google.iam.v1.Policy) {
    option (google.api.http) = {
      post: "/v2/{resource=projects/*/instances/*}:setIamPolicy"
      body: "*"
    };
    option (google.api.method_signature) = "resource,policy";
  }

  // Returns permissions that the caller has on the specified instance resource.
  rpc TestIamPermissions(google.iam.v1.TestIamPermissionsRequest)
      returns (google.iam.v1.TestIamPermissionsResponse) {
    option (google.api.http) = {
      post: "/v2/{resource=projects/*/instances/*}:testIamPermissions"
      body: "*"
    };
    option (google.api.method_signature) = "resource,permissions";
  }

  // Lists Cloud Bigtable backups. Returns both completed and pending
  // backups.
  rpc ListBackups(ListBackupsRequest) returns (ListBackupsResponse) {
    option (google.api.http) = {
      get: "/v2/{parent=projects/*/instances/*/clusters/*}/backups"
    };
  }

  // Gets metadata on a pending or completed Cloud Bigtable Backup.
  rpc GetBackup(GetBackupRequest) returns (Backup) {
    option (google.api.http) = {
      get: "/v2/{name=projects/*/instances/*/clusters/*/backups/*}"
    };
  }

  // Updates a pending or completed Cloud Bigtable Backup.
  rpc UpdateBackup(UpdateBackupRequest) returns (Backup) {
    option (google.api.http) = {
      patch: "/v2/{backup.name=projects/*/instances/*/clusters/*/backups/*}"
      body: "backup"
    };
  }

  // Deletes a pending or completed Cloud Bigtable backup.
  rpc DeleteBackup(DeleteBackupRequest) returns (google.protobuf.Empty) {
    option (google.api.http) = {
      delete: "/v2/{name=projects/*/instances/*/clusters/*/backups/*}"
    };
  }

  // Create a new table by restoring a completed backup to it, and begin
  // restoring the new table's data from the backup immediately. The new
  // table is readable as soon as (and shortly after) the operation is
  // completed. If the `parent` backup's source cluster's storage_type is
  // `STORAGE_TYPE_UNSPECIFIED`, then the source cluster and the new table must
  // have the same cluster storage type. If the `parent` backup's source
  // cluster's storage type is `HDD`, then the new table must have `HDD`
  // clusters. If the `parent` backup's source cluster's storage type is `SSD`,
  // then the new table must have `SSD` clusters.
  //
  // Any servers in clusters lacking the required storage type will not be able
  // to serve the new table.
  rpc RestoreTable(RestoreTableRequest) returns (google.longrunning.Operation) {
    option (google.api.http) = {
      post: "/v2/{parent=projects/*/instances/*}/tables:restore"
      body: "*"
    };
    option (google.longrunning.operation_info) = {
      response_type: "Table"
      metadata_type: "RestoreTableMetadata"
    };
  }

  // Get a Cloud Bigtable Backup.
  rpc GetBackupIamPolicy(google.iam.v1.GetIamPolicyRequest)
      returns (google.iam.v1.Policy) {
    option (google.api.http) = {
      post: "/v2/{resource=projects/*/instances/*/clusters/*/backups/*}:getIamPolicy"
      body: "*"
    };
    option (google.api.method_signature) = "resource";
  }

  // Set a Cloud Bigtable Backup IAM Policy.
  rpc SetBackupIamPolicy(google.iam.v1.SetIamPolicyRequest)
      returns (google.iam.v1.Policy) {
    option (google.api.http) = {
      post: "/v2/{resource=projects/*/instances/*/clusters/*/backups/*}:setIamPolicy"
      body: "*"
    };
    option (google.api.method_signature) = "resource";
  }

  // Test a Cloud Bigtable Backup IAM Permissions.
  rpc TestBackupIamPermissions(google.iam.v1.TestIamPermissionsRequest)
      returns (google.iam.v1.TestIamPermissionsResponse) {
    option (google.api.http) = {
      post: "/v2/{resource=projects/*/instances/*/clusters/*/backups/*}:testIamPermissions"
      body: "*"
    };
    option (google.api.method_signature) = "resource,permissions";
  }
}

```
out:
```go
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

// +tool:mockgcp-support
// proto.service: google.bigtable.admin.v2.BigtableInstanceAdmin
// proto.message: google.bigtable.admin.v2.AppProfile

package mockbigtable

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/bigtable/admin/v2"
)

func (s *MockService) GetAppProfile(ctx context.Context, req *pb.GetAppProfileRequest) (*pb.AppProfile, error) {
	name, err := s.parseAppProfileName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.AppProfile{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "appProfile %q not found", name)
		}
		return nil, err
	}

	return obj, nil
}

func (s *MockService) CreateAppProfile(ctx context.Context, req *pb.CreateAppProfileRequest) (*pb.AppProfile, error) {
	reqName := req.Parent + "/appProfiles/" + req.AppProfileId
	name, err := s.parseAppProfileName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.AppProfile).(*pb.AppProfile)
	obj.Name = fqn

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *MockService) UpdateAppProfile(ctx context.Context, req *pb.UpdateAppProfileRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseAppProfileName(req.GetAppProfile().GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	existing := &pb.AppProfile{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		return nil, err
	}

	updated := proto.Clone(existing).(*pb.AppProfile)

	// Required. The set of fields to update.
	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask must be provided")
	}

	// TODO: Some sort of helper for fieldmask?
	for _, path := range paths {
		switch path {
		case "description":
			updated.Description = req.GetAppProfile().GetDescription()
		case "multi_cluster_routing_use_any":
			updated.RoutingPolicy = &pb.AppProfile_MultiClusterRoutingUseAny_{
				MultiClusterRoutingUseAny: req.GetAppProfile().GetMultiClusterRoutingUseAny(),
			}
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
		}
	}

	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}

	return s.operations.notImplementedLRO(), nil
}

type appProfileName struct {
	Project    string
	InstanceID string
	AppProfile string
}

func (n *appProfileName) String() string {
	return fmt.Sprintf("projects/%s/instances/%s/appProfiles/%s", n.Project, n.InstanceID, n.AppProfile)
}

// parseAppProfileName parses a string into a appProfileName.
// The expected form is `projects/*/instances/*/appProfiles/*`.
func (s *MockService) parseAppProfileName(name string) (*appProfileName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "instances" && tokens[4] == "appProfiles" {
		name := &appProfileName{
			Project:    tokens[1],
			InstanceID: tokens[3],
			AppProfile: tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}



