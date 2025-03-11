in.proto.service: google.bigtable.admin.v2.BigtableTableAdmin
in.proto.service.definition: service BigtableTableAdmin {
  option (google.api.default_host) = "bigtableadmin.googleapis.com";
  option (google.api.oauth_scopes) =
      "https://www.googleapis.com/auth/bigtable.admin,"
      "https://www.googleapis.com/auth/bigtable.admin.table,"
      "https://www.googleapis.com/auth/cloud-bigtable.admin,"
      "https://www.googleapis.com/auth/cloud-bigtable.admin.table,"
      "https://www.googleapis.com/auth/cloud-platform,"
      "https://www.googleapis.com/auth/cloud-platform.read-only";

  // Creates a new table in the specified instance.
  // The table can be created with a full set of initial column families,
  // specified in the request.
  rpc CreateTable(CreateTableRequest) returns (Table) {
    option (google.api.http) = {
      post: "/v2/{parent=projects/*/instances/*}/tables"
      body: "table"
    };
    option (google.api.method_signature) = "parent,table_id,table";
  }

  // Creates a new table from the specified snapshot. The target table must
  // not exist. The snapshot and the table must be in the same instance.
  //
  // Note: This is a private alpha release of Cloud Bigtable snapshots. This
  // feature is not currently available to most Cloud Bigtable customers. This
  // feature might be changed in backward-incompatible ways and is not
  // recommended for production use. It is not subject to any SLA or deprecation
  // policy.
  rpc CreateTableFromSnapshot(CreateTableFromSnapshotRequest)
      returns (google.longrunning.Operation) {
    option (google.api.http) = {
      post: "/v2/{parent=projects/*/instances/*}/tables:createFromSnapshot"
      body: "*"
    };
    option (google.api.method_signature) = "parent,table_id,source_snapshot";
    option (google.longrunning.operation_info) = {
      response_type: "Table"
      metadata_type: "CreateTableFromSnapshotMetadata"
    };
    option (google.api.alpha) = true;
  }

  // Lists all tables served from a specified instance.
  rpc ListTables(ListTablesRequest) returns (ListTablesResponse) {
    option (google.api.http) = {
      get: "/v2/{parent=projects/*/instances/*}/tables"
    };
    option (google.api.method_signature) = "parent";
  }

  // Gets metadata information about the specified table.
  rpc GetTable(GetTableRequest) returns (Table) {
    option (google.api.http) = {
      get: "/v2/{name=projects/*/instances/*/tables/*}"
    };
    option (google.api.method_signature) = "name";
  }

  // Permanently deletes a specified table and all of its data.
  rpc DeleteTable(DeleteTableRequest) returns (google.protobuf.Empty) {
    option (google.api.http) = {
      delete: "/v2/{name=projects/*/instances/*/tables/*}"
    };
    option (google.api.method_signature) = "name";
  }

  // Performs a series of column family modifications on the specified table.
  // Either all or none of the modifications will occur before this method
  // returns, but data requests received prior to that point may see a table
  // where only some modifications have taken effect.
  rpc ModifyColumnFamilies(ModifyColumnFamiliesRequest) returns (Table) {
    option (google.api.http) = {
      post: "/v2/{name=projects/*/instances/*/tables/*}:modifyColumnFamilies"
      body: "*"
    };
    option (google.api.method_signature) = "name,modifications";
  }

  // Permanently drop/delete a row range from a specified table. The request can
  // specify whether to drop all rows in a table, or only those that match a
  // particular prefix.
  rpc DropRowRange(DropRowRangeRequest) returns (google.protobuf.Empty) {
    option (google.api.http) = {
      post: "/v2/{name=projects/*/instances/*/tables/*}:dropRowRange"
      body: "*"
    };
    option (google.api.method_signature) = "name,row_key_prefix";
    option (google.api.method_signature) = "name,delete_all_data_from_table";
  }

  // Generates a series of rpc tokens which can be used to perform large data
  // operations in parallel.  Each of the returned tokens should be used for a
  // single MutateRows operation on the specified table.  These tokens should be
  // used within 30 minutes of creation time, after which they will become
  // invalid.
  rpc GenerateConsistencyToken(GenerateConsistencyTokenRequest)
      returns (GenerateConsistencyTokenResponse) {
    option (google.api.http) = {
      post: "/v2/{name=projects/*/instances/*/tables/*}:generateConsistencyToken"
      body: "*"
    };
    option (google.api.method_signature) = "name";
  }

  // Checks replication consistency based on a consistency token, that is, if
  // replication has caught up based on the conditions specified in the token
  // and the check request.
  rpc CheckConsistency(CheckConsistencyRequest)
      returns (CheckConsistencyResponse) {
    option (google.api.http) = {
      post: "/v2/{name=projects/*/instances/*/tables/*}:checkConsistency"
      body: "*"
    };
    option (google.api.method_signature) = "name,consistency_token";
  }

  // Creates a new snapshot in the specified cluster from the specified
  // source table. The cluster and the table must be in the same instance.
  //
  // Note: This is a private alpha release of Cloud Bigtable snapshots. This
  // feature is not currently available to most Cloud Bigtable customers. This
  // feature might be changed in backward-incompatible ways and is not
  // recommended for production use. It is not subject to any SLA or deprecation
  // policy.
  rpc SnapshotTable(SnapshotTableRequest) returns (google.longrunning.Operation) {
    option (google.api.http) = {
      post: "/v2/{name=projects/*/instances/*/tables/*}:snapshot"
      body: "*"
    };
    option (google.api.method_signature) =
        "name,cluster,snapshot_id,description";
    option (google.longrunning.operation_info) = {
      response_type: "Snapshot"
      metadata_type: "SnapshotTableMetadata"
    };
    option (google.api.alpha) = true;
  }

  // Gets metadata information about the specified snapshot.
  //
  // Note: This is a private alpha release of Cloud Bigtable snapshots. This
  // feature is not currently available to most Cloud Bigtable customers. This
  // feature might be changed in backward-incompatible ways and is not
  // recommended for production use. It is not subject to any SLA or deprecation
  // policy.
  rpc GetSnapshot(GetSnapshotRequest) returns (Snapshot) {
    option (google.api.http) = {
      get: "/v2/{name=projects/*/instances/*/clusters/*/snapshots/*}"
    };
    option (google.api.method_signature) = "name";
    option (google.api.alpha) = true;
  }

  // Lists all snapshots associated with the specified cluster.
  //
  // Note: This is a private alpha release of Cloud Bigtable snapshots. This
  // feature is not currently available to most Cloud Bigtable customers. This
  // feature might be changed in backward-incompatible ways and is not
  // recommended for production use. It is not subject to any SLA or deprecation
  // policy.
  rpc ListSnapshots(ListSnapshotsRequest) returns (ListSnapshotsResponse) {
    option (google.api.http) = {
      get: "/v2/{parent=projects/*/instances/*/clusters/*}/snapshots"
    };
    option (google.api.method_signature) = "parent";
    option (google.api.alpha) = true;
  }

  // Permanently deletes the specified snapshot.
  //
  // Note: This is a private alpha release of Cloud Bigtable snapshots. This
  // feature is not currently available to most Cloud Bigtable customers. This
  // feature might be changed in backward-incompatible ways and is not
  // recommended for production use. It is not subject to any SLA or deprecation
  // policy.
  rpc DeleteSnapshot(DeleteSnapshotRequest) returns (google.protobuf.Empty) {
    option (google.api.http) = {
      delete: "/v2/{name=projects/*/instances/*/clusters/*/snapshots/*}"
    };
    option (google.api.method_signature) = "name";
    option (google.api.alpha) = true;
  }

  // Starts creating a new Cloud Bigtable Backup.  The returned backup
  // [long-running operation][google.longrunning.Operation] can be used to
  // track creation of the backup, by calling
  // [GetOperation][google.longrunning.Operations.GetOperation].
  //
  // Immediately upon completion of this request:
  //
  // * The backup [state][google.bigtable.admin.v2.Backup.state] is set to
  //   `CREATING`. The backup is usable for restores after
  //   its creation completes, i.e. the creation operation finishes successfully.
  //   To check whether a request is complete, use the
  //   [GetOperation][google.longrunning.Operations.GetOperation] method of the
  //   `google.longrunning.Operations` interface:
  //
  //   If the `GetOperation` response's `done` field is `true` and its
  //   [error][google.longrunning.Operation.error] field is empty, the
  //   request has finished.
  //
  // If the request is still in progress:
  //
  //   * You can call `GetOperation` with the `name` of the `Operation`.
  //     This is the value of the
  //     [name][google.longrunning.Operation.name] field returned by the
  //     `CreateBackup` method.
  //
  // If the request is failing, the `Operation` has a non-empty `error` field.
  //
  // * The `metadata.progress.start_time` is set to the initiation
  //   time of the request.
  // * The response is a `Backup` object with all fields populated except
  //   `name`, `expire_time`, `start_time`, `end_time` and `size_bytes`.
  //
  // For the value of `name`, call
  // [GetOperation][google.longrunning.Operations.GetOperation].
  // The field `name` in the `Backup` is of the form
  // `projects/{project}/instances/{instance}/clusters/{cluster}/backups/{backup}`.
  // The cluster is the one in the `CreateBackupRequest`, and the `backup` is
  // the value of the `backup_id` field in the request.
  //
  // The `error` field is an instance of [google.rpc.Status]
  // (https://github.com/googleapis/googleapis/blob/master/google/rpc/status.proto).
  // It has the structure described in
  // https://cloud.google.com/apis/design/errors. Specifically:
  //
  //   * The `status.code` value is a valid [google.rpc.Code]
  //   (https://github.com/googleapis/googleapis/blob/master/google/rpc/code.proto).
  //   * The `status.details` field contains the following:
  //       * The `type_url` value is type.googleapis.com/google.rpc.ErrorInfo.
  //       * The `reason` field value is the value of `status.code`.
  //       * The `domain` field value is bigtable.googleapis.com.
  //       * The `metadata` field contains the request as a JSON encoded
  //       string if available.
  //       * The `metadata` field contains the response (if available).
  //
  //   If the `status.code` is not OK, then the request failed.
  //
  //   * All other fields of the `metadata` are ignored.
  //   * All other fields of the `response` are ignored.
  rpc CreateBackup(CreateBackupRequest) returns (google.longrunning.Operation) {
    option (google.api.http) = {
      post: "/v2/{parent=projects/*/instances/*/clusters/*}/backups"
      body: "backup"
    };
    option (google.api.method_signature) = "parent,backup_id,backup";
    option (google.longrunning.operation_info) = {
      response_type: "Backup"
      metadata_type: "CreateBackupMetadata"
    };
  }

  // Gets metadata on a pending or completed Cloud Bigtable Backup.
  rpc GetBackup(GetBackupRequest) returns (Backup) {
    option (google.api.http) = {
      get: "/v2/{name=projects/*/instances/*/clusters/*/backups/*}"
    };
    option (google.api.method_signature) = "name";
  }

  // Updates a pending or completed Cloud Bigtable Backup.
  rpc UpdateBackup(UpdateBackupRequest) returns (Backup) {
    option (google.api.http) = {
      patch: "/v2/{backup.name=projects/*/instances/*/clusters/*/backups/*}"
      body: "backup"
    };
    option (google.api.method_signature) = "backup,update_mask";
  }

  // Deletes a pending or completed Cloud Bigtable backup.
  rpc DeleteBackup(DeleteBackupRequest) returns (google.protobuf.Empty) {
    option (google.api.http) = {
      delete: "/v2/{name=projects/*/instances/*/clusters/*/backups/*}"
    };
    option (google.api.method_signature) = "name";
  }

  // Lists Cloud Bigtable backups. Returns both completed and pending
  // backups.
  rpc ListBackups(ListBackupsRequest) returns (ListBackupsResponse) {
    option (google.api.http) = {
      get: "/v2/{parent=projects/*/instances/*/clusters/*}/backups"
    };
    option (google.api.method_signature) = "parent";
  }

  // Create a new table by restoring from a completed backup. The
  // new table must be in the same instance as the instance containing
  // the backup.  The returned table
  // [long-running operation][google.longrunning.Operation] can be used to track
  // the progress of the operation, and to cancel it.  The
  // [metadata][google.longrunning.Operation.metadata] field type is
  // [RestoreTableMetadata][google.bigtable.admin.RestoreTableMetadata].
  // The [response][google.longrunning.Operation.response] type is
  // [Table][google.bigtable.admin.v2.Table], if successful.
  //
  // The following values can be updated after restoring the table:
  //
  //     * labels
  //     * replication
  //
  // The `name`, `cluster_states` and `column_families` columns can not be
  // updated.
  //
  // Note: This is a private alpha release of Cloud Bigtable restore table.
  // This feature is not currently available to most Cloud Bigtable customers.
  // This feature might be changed in backward-incompatible ways and is not
  // recommended for production use. It is not subject to any SLA or deprecation
  // policy.
  rpc RestoreTable(RestoreTableRequest) returns (google.longrunning.Operation) {
    option (google.api.http) = {
      post: "/v2/{parent=projects/*/instances/*}/tables:restore"
      body: "*"
    };
    option (google.longrunning.operation_info) = {
      response_type: "Table"
      metadata_type: "RestoreTableMetadata"
    };
    option (google.api.alpha) = true;
  }

  // Get the schema of the restored table by using the backup name.
  //
  // Note: This is a private alpha release of Cloud Bigtable restore table.
  // This feature is not currently available to most Cloud Bigtable customers.
  // This feature might be changed in backward-incompatible ways and is not
  // recommended for production use. It is not subject to any SLA or deprecation
  // policy.
  rpc GetRestoreInfo(GetRestoreInfoRequest) returns (RestoreInfo) {
    option (google.api.http) = {
      get: "/v2/{source_backup=projects/*/instances/*/clusters/*/backups/*}:restoreInfo"
    };
    option (google.api.alpha) = true;
  }

  // Gets the access control policy for a Table or Backup resource.
  // Returns an empty policy if the resource exists but does not have a policy
  // set.
  rpc GetIamPolicy(google.iam.v1.GetIamPolicyRequest) returns (google.iam.v1.Policy) {
    option (google.api.http) = {
      post: "/v2/{resource=projects/*/instances/*/tables/*}:getIamPolicy"
      body: "*"
      additional_bindings {
        post: "/v2/{resource=projects/*/instances/*/clusters/*/backups/*}:getIamPolicy"
        body: "*"
      }
    };
    option (google.api.method_signature) = "resource";
  }

  // Sets the access control policy on a Table or Backup resource.
  // Replaces any existing policy.
  rpc SetIamPolicy(google.iam.v1.SetIamPolicyRequest) returns (google.iam.v1.Policy) {
    option (google.api.http) = {
      post: "/v2/{resource=projects/*/instances/*/tables/*}:setIamPolicy"
      body: "*"
      additional_bindings {
        post: "/v2/{resource=projects/*/instances/*/clusters/*/backups/*}:setIamPolicy"
        body: "*"
      }
    };
    option (google.api.method_signature) = "resource,policy";
  }

  // Returns permissions that the caller has on the specified Table or Backup
  // resource.
  rpc TestIamPermissions(google.iam.v1.TestIamPermissionsRequest) returns (google.iam.v1.TestIamPermissionsResponse) {
    option (google.api.http) = {
      post: "/v2/{resource=projects/*/instances/*/tables/*}:testIamPermissions"
      body: "*"
      additional_bindings {
        post: "/v2/{resource=projects/*/instances/*/clusters/*/backups/*}:testIamPermissions"
        body: "*"
      }
    };
    option (google.api.method_signature) = "resource,permissions";
  }
}
out: // Copyright 2025 Google LLC
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
// proto.service: google.bigtable.admin.v2.BigtableTableAdmin
// proto.message: google.bigtable.admin.v2.Backup

package mockbigtable

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"

	pb "cloud.google.com/go/bigtable/admin/apiv2/adminpb"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

func (s *MockService) CreateBackup(ctx context.Context, req *pb.CreateBackupRequest) (*longrunningpb.Operation, error) {
	reqName := req.Parent + "/backups/" + req.BackupId
	name, err := s.parseBackupName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.Backup).(*pb.Backup)
	obj.Name = fqn

	s.populateDefaultsForBackup(obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/instances/%s/clusters/%s", name.Project.ID, name.Instance, name.Cluster)
	lroMetadata := &pb.CreateBackupMetadata{
		Name:         name.String(),
		SourceTable:  req.Backup.SourceTable,
		StartTime:    timestamppb.New(time.Now()),
		EndTime:      timestamppb.New(time.Now().Add(5 * time.Minute)),
		RequestTime:  timestamppb.Now(),
		Progress:     &pb.OperationProgress{},
		BackupInfo:   &pb.BackupInfo{},
		SourceBackup: req.Backup.SourceBackup,
	}
	return s.operations.StartLRO(ctx, prefix, lroMetadata, func() (proto.Message, error) {
		obj.State = pb.Backup_READY
		obj.SizeBytes = 1562632

		return obj, nil
	})

}

func (s *MockService) populateDefaultsForBackup(obj *pb.Backup) {
	if obj.GetBackupType() == pb.Backup_BACKUP_TYPE_UNSPECIFIED {
		obj.BackupType = pb.Backup_BACKUP_TYPE_UNSPECIFIED
	}
	if obj.GetState() == pb.Backup_STATE_UNSPECIFIED {
		obj.State = pb.Backup_CREATING
	}

	if obj.StartTime == nil {
		obj.StartTime = timestamppb.Now()
	}

	if obj.SourceTable == "" {
		obj.SourceTable = "sample-table"
	}

}

func (s *MockService) GetBackup(ctx context.Context, req *pb.GetBackupRequest) (*pb.Backup, error) {
	name, err := s.parseBackupName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Backup{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "backup %q not found", name)
		}
		return nil, err
	}

	return obj, nil
}

func (s *MockService) UpdateBackup(ctx context.Context, req *pb.UpdateBackupRequest) (*pb.Backup, error) {
	name, err := s.parseBackupName(req.Backup.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Backup{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask must be provided")
	}

	for _, path := range paths {
		switch path {
		case "expire_time":
			obj.ExpireTime = req.Backup.ExpireTime
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
		}
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *MockService) DeleteBackup(ctx context.Context, req *pb.DeleteBackupRequest) (*emptypb.Empty, error) {
	name, err := s.parseBackupName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.Backup{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

type backupName struct {
	Project  *projects.ProjectData
	Instance string
	Cluster  string
	BackupID string
}

func (n *backupName) String() string {
	return fmt.Sprintf("projects/%s/instances/%s/clusters/%s/backups/%s", n.Project.ID, n.Instance, n.Cluster, n.BackupID)
}

// parseBackupName parses a string into a backupName.
// The expected form is `projects/*/instances/*/clusters/*/backups/*`.
func (s *MockService) parseBackupName(name string) (*backupName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "instances" && tokens[4] == "clusters" && tokens[6] == "backups" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &backupName{
			Project:  project,
			Instance: tokens[3],
			Cluster:  tokens[5],
			BackupID: tokens[7],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}


</example>




<example>
in.proto.message: google.cloud.pubsublite.v1.TopicPartitions
in.proto.service.definition: service AdminService {
  option (google.api.default_host) = "pubsublite.googleapis.com";
  option (google.api.oauth_scopes) =
      "https://www.googleapis.com/auth/cloud-platform";

  // Creates a new topic.
  rpc CreateTopic(CreateTopicRequest) returns (Topic) {
    option (google.api.http) = {
      post: "/v1/admin/{parent=projects/*/locations/*}/topics"
      body: "topic"
    };
    option (google.api.method_signature) = "parent,topic,topic_id";
  }

  // Returns the topic configuration.
  rpc GetTopic(GetTopicRequest) returns (Topic) {
    option (google.api.http) = {
      get: "/v1/admin/{name=projects/*/locations/*/topics/*}"
    };
    option (google.api.method_signature) = "name";
  }

  // Returns the partition information for the requested topic.
  rpc GetTopicPartitions(GetTopicPartitionsRequest) returns (TopicPartitions) {
    option (google.api.http) = {
      get: "/v1/admin/{name=projects/*/locations/*/topics/*}/partitions"
    };
    option (google.api.method_signature) = "name";
  }

  // Returns the list of topics for the given project.
  rpc ListTopics(ListTopicsRequest) returns (ListTopicsResponse) {
    option (google.api.http) = {
      get: "/v1/admin/{parent=projects/*/locations/*}/topics"
    };
    option (google.api.method_signature) = "parent";
  }

  // Updates properties of the specified topic.
  rpc UpdateTopic(UpdateTopicRequest) returns (Topic) {
    option (google.api.http) = {
      patch: "/v1/admin/{topic.name=projects/*/locations/*/topics/*}"
      body: "topic"
    };
    option (google.api.method_signature) = "topic,update_mask";
  }

  // Deletes the specified topic.
  rpc DeleteTopic(DeleteTopicRequest) returns (google.protobuf.Empty) {
    option (google.api.http) = {
      delete: "/v1/admin/{name=projects/*/locations/*/topics/*}"
    };
    option (google.api.method_signature) = "name";
  }

  // Lists the subscriptions attached to the specified topic.
  rpc ListTopicSubscriptions(ListTopicSubscriptionsRequest)
      returns (ListTopicSubscriptionsResponse) {
    option (google.api.http) = {
      get: "/v1/admin/{name=projects/*/locations/*/topics/*}/subscriptions"
    };
    option (google.api.method_signature) = "name";
  }

  // Creates a new subscription.
  rpc CreateSubscription(CreateSubscriptionRequest) returns (Subscription) {
    option (google.api.http) = {
      post: "/v1/admin/{parent=projects/*/locations/*}/subscriptions"
      body: "subscription"
    };
    option (google.api.method_signature) =
        "parent,subscription,subscription_id";
  }

  // Returns the subscription configuration.
  rpc GetSubscription(GetSubscriptionRequest) returns (Subscription) {
    option (google.api.http) = {
      get: "/v1/admin/{name=projects/*/locations/*/subscriptions/*}"
    };
    option (google.api.method_signature) = "name";
  }

  // Returns the list of subscriptions for the given project.
  rpc ListSubscriptions(ListSubscriptionsRequest)
      returns (ListSubscriptionsResponse) {
    option (google.api.http) = {
      get: "/v1/admin/{parent=projects/*/locations/*}/subscriptions"
    };
    option (google.api.method_signature) = "parent";
  }

  // Updates properties of the specified subscription.
  rpc UpdateSubscription(UpdateSubscriptionRequest) returns (Subscription) {
    option (google.api.http) = {
      patch: "/v1/admin/{subscription.name=projects/*/locations/*/subscriptions/*}"
      body: "subscription"
    };
    option (google.api.method_signature) = "subscription,update_mask";
  }

  // Deletes the specified subscription.
  rpc DeleteSubscription(DeleteSubscriptionRequest)
      returns (google.protobuf.Empty) {
    option (google.api.http) = {
      delete: "/v1/admin/{name=projects/*/locations/*/subscriptions/*}"
    };
    option (google.api.method_signature) = "name";
  }

  // Performs an out-of-band seek for a subscription to a specified target,
  // which may be timestamps or named positions within the message backlog.
  // Seek translates these targets to cursors for each partition and
  // orchestrates subscribers to start consuming messages from these seek
  // cursors.
  //
  // If an operation is returned, the seek has been registered and subscribers
  // will eventually receive messages from the seek cursors (i.e. eventual
  // consistency), as long as they are using a minimum supported client library
  // version and not a system that tracks cursors independently of Pub/Sub Lite
  // (e.g. Apache Beam, Dataflow, Spark). The seek operation will fail for
  // unsupported clients.
  //
  // If clients would like to know when subscribers react to the seek (or not),
  // they can poll the operation. The seek operation will succeed and complete
  // once subscribers are ready to receive messages from the seek cursors for
  // all partitions of the topic. This means that the seek operation will not
  // complete until all subscribers come online.
  //
  // If the previous seek operation has not yet completed, it will be aborted
  // and the new invocation of seek will supersede it.
  rpc SeekSubscription(SeekSubscriptionRequest)
      returns (google.longrunning.Operation) {
    option (google.api.http) = {
      post: "/v1/admin/{name=projects/*/locations/*/subscriptions/*}:seek"
      body: "*"
    };
    option (google.longrunning.operation_info) = {
      response_type: "SeekSubscriptionResponse"
      metadata_type: "OperationMetadata"
    };
  }

  // Creates a new reservation.
  rpc CreateReservation(CreateReservationRequest) returns (Reservation) {
    option (google.api.http) = {
      post: "/v1/admin/{parent=projects/*/locations/*}/reservations"
      body: "reservation"
    };
    option (google.api.method_signature) = "parent,reservation,reservation_id";
  }

  // Returns the reservation configuration.
  rpc GetReservation(GetReservationRequest) returns (Reservation) {
    option (google.api.http) = {
      get: "/v1/admin/{name=projects/*/locations/*/reservations/*}"
    };
    option (google.api.method_signature) = "name";
  }

  // Returns the list of reservations for the given project.
  rpc ListReservations(ListReservationsRequest)
      returns (ListReservationsResponse) {
    option (google.api.http) = {
      get: "/v1/admin/{parent=projects/*/locations/*}/reservations"
    };
    option (google.api.method_signature) = "parent";
  }

  // Updates properties of the specified reservation.
  rpc UpdateReservation(UpdateReservationRequest) returns (Reservation) {
    option (google.api.http) = {
      patch: "/v1/admin/{reservation.name=projects/*/locations/*/reservations/*}"
      body: "reservation"
    };
    option (google.api.method_signature) = "reservation,update_mask";
  }

  // Deletes the specified reservation.
  rpc DeleteReservation(DeleteReservationRequest)
      returns (google.protobuf.Empty) {
    option (google.api.http) = {
      delete: "/v1/admin/{name=

