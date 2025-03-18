google.cloud.dataplex.v1.DataplexService
in.proto.service.definition: service DataplexService {
  option (google.api.default_host) = "dataplex.googleapis.com";
  option (google.api.oauth_scopes) =
      "https://www.googleapis.com/auth/cloud-platform";

  // Creates a lake resource.
  rpc CreateLake(CreateLakeRequest) returns (google.longrunning.Operation) {
    option (google.api.http) = {
      post: "/v1/{parent=projects/*/locations/*}/lakes"
      body: "lake"
    };
    option (google.api.method_signature) = "parent,lake,lake_id";
    option (google.longrunning.operation_info) = {
      response_type: "Lake"
      metadata_type: "OperationMetadata"
    };
  }

  // Updates a lake resource.
  rpc UpdateLake(UpdateLakeRequest) returns (google.longrunning.Operation) {
    option (google.api.http) = {
      patch: "/v1/{lake.name=projects/*/locations/*/lakes/*}"
      body: "lake"
    };
    option (google.api.method_signature) = "lake,update_mask";
    option (google.longrunning.operation_info) = {
      response_type: "Lake"
      metadata_type: "OperationMetadata"
    };
  }

  // Deletes a lake resource.
  rpc DeleteLake(DeleteLakeRequest) returns (google.longrunning.Operation) {
    option (google.api.http) = {
      delete: "/v1/{name=projects/*/locations/*/lakes/*}"
    };
    option (google.api.method_signature) = "name";
    option (google.longrunning.operation_info) = {
      response_type: "google.protobuf.Empty"
      metadata_type: "OperationMetadata"
    };
  }

  // Lists lake resources in a project and location.
  rpc ListLakes(ListLakesRequest) returns (ListLakesResponse) {
    option (google.api.http) = {
      get: "/v1/{parent=projects/*/locations/*}/lakes"
    };
    option (google.api.method_signature) = "parent";
  }

  // Retrieves a lake resource.
  rpc GetLake(GetLakeRequest) returns (Lake) {
    option (google.api.http) = {
      get: "/v1/{name=projects/*/locations/*/lakes/*}"
    };
    option (google.api.method_signature) = "name";
  }

  // Lists action resources in a lake.
  rpc ListLakeActions(ListLakeActionsRequest)
      returns (ListLakeActionsResponse) {
    option (google.api.http) = {
      get: "/v1/{parent=projects/*/locations/*/lakes/*}/actions"
    };
    option (google.api.method_signature) = "parent";
  }

  // Creates a zone resource within a lake.
  rpc CreateZone(CreateZoneRequest) returns (google.longrunning.Operation) {
    option (google.api.http) = {
      post: "/v1/{parent=projects/*/locations/*/lakes/*}/zones"
      body: "zone"
    };
    option (google.api.method_signature) = "parent,zone,zone_id";
    option (google.longrunning.operation_info) = {
      response_type: "Zone"
      metadata_type: "OperationMetadata"
    };
  }

  // Updates a zone resource.
  rpc UpdateZone(UpdateZoneRequest) returns (google.longrunning.Operation) {
    option (google.api.http) = {
      patch: "/v1/{zone.name=projects/*/locations/*/lakes/*/zones/*}"
      body: "zone"
    };
    option (google.api.method_signature) = "zone,update_mask";
    option (google.longrunning.operation_info) = {
      response_type: "Zone"
      metadata_type: "OperationMetadata"
    };
  }

  // Deletes a zone resource. All assets within a zone must be deleted before
  // the zone can be deleted.
  rpc DeleteZone(DeleteZoneRequest) returns (google.longrunning.Operation) {
    option (google.api.http) = {
      delete: "/v1/{name=projects/*/locations/*/lakes/*/zones/*}"
    };
    option (google.api.method_signature) = "name";
    option (google.longrunning.operation_info) = {
      response_type: "google.protobuf.Empty"
      metadata_type: "OperationMetadata"
    };
  }

  // Lists zone resources in a lake.
  rpc ListZones(ListZonesRequest) returns (ListZonesResponse) {
    option (google.api.http) = {
      get: "/v1/{parent=projects/*/locations/*/lakes/*}/zones"
    };
    option (google.api.method_signature) = "parent";
  }

  // Retrieves a zone resource.
  rpc GetZone(GetZoneRequest) returns (Zone) {
    option (google.api.http) = {
      get: "/v1/{name=projects/*/locations/*/lakes/*/zones/*}"
    };
    option (google.api.method_signature) = "name";
  }

  // Lists action resources in a zone.
  rpc ListZoneActions(ListZoneActionsRequest)
      returns (ListZoneActionsResponse) {
    option (google.api.http) = {
      get: "/v1/{parent=projects/*/locations/*/lakes/*/zones/*}/actions"
    };
    option (google.api.method_signature) = "parent";
  }

  // Creates an asset resource.
  rpc CreateAsset(CreateAssetRequest) returns (google.longrunning.Operation) {
    option (google.api.http) = {
      post: "/v1/{parent=projects/*/locations/*/lakes/*/zones/*}/assets"
      body: "asset"
    };
    option (google.api.method_signature) = "parent,asset,asset_id";
    option (google.longrunning.operation_info) = {
      response_type: "Asset"
      metadata_type: "OperationMetadata"
    };
  }

  // Updates an asset resource.
  rpc UpdateAsset(UpdateAssetRequest) returns (google.longrunning.Operation) {
    option (google.api.http) = {
      patch: "/v1/{asset.name=projects/*/locations/*/lakes/*/zones/*/assets/*}"
      body: "asset"
    };
    option (google.api.method_signature) = "asset,update_mask";
    option (google.longrunning.operation_info) = {
      response_type: "Asset"
      metadata_type: "OperationMetadata"
    };
  }

  // Deletes an asset resource. The referenced storage resource is detached
  // (default) or deleted based on the associated data location policy.
  rpc DeleteAsset(DeleteAssetRequest) returns (google.longrunning.Operation) {
    option (google.api.http) = {
      delete: "/v1/{name=projects/*/locations/*/lakes/*/zones/*/assets/*}"
    };
    option (google.api.method_signature) = "name";
    option (google.longrunning.operation_info) = {
      response_type: "google.protobuf.Empty"
      metadata_type: "OperationMetadata"
    };
  }

  // Lists asset resources in a zone.
  rpc ListAssets(ListAssetsRequest) returns (ListAssetsResponse) {
    option (google.api.http) = {
      get: "/v1/{parent=projects/*/locations/*/lakes/*/zones/*}/assets"
    };
    option (google.api.method_signature) = "parent";
  }

  // Retrieves an asset resource.
  rpc GetAsset(GetAssetRequest) returns (Asset) {
    option (google.api.http) = {
      get: "/v1/{name=projects/*/locations/*/lakes/*/zones/*/assets/*}"
    };
    option (google.api.method_signature) = "name";
  }

  // Lists action resources in an asset.
  rpc ListAssetActions(ListAssetActionsRequest)
      returns (ListAssetActionsResponse) {
    option (google.api.http) = {
      get: "/v1/{parent=projects/*/locations/*/lakes/*/zones/*/assets/*}/actions"
    };
    option (google.api.method_signature) = "parent";
  }

  // Lists tasks under the given lake.
  rpc ListTasks(ListTasksRequest) returns (ListTasksResponse) {
    option (google.api.http) = {
      get: "/v1/{parent=projects/*/locations/*/lakes/*}/tasks"
    };
    option (google.api.method_signature) = "parent";
  }

  // Get task resource.
  rpc GetTask(GetTaskRequest) returns (Task) {
    option (google.api.http) = {
      get: "/v1/{name=projects/*/locations/*/lakes/*/tasks/*}"
    };
    option (google.api.method_signature) = "name";
  }

  // Creates a task resource within a lake.
  rpc CreateTask(CreateTaskRequest) returns (google.longrunning.Operation) {
    option (google.api.http) = {
      post: "/v1/{parent=projects/*/locations/*/lakes/*}/tasks"
      body: "task"
    };
    option (google.api.method_signature) = "parent,task,task_id";
    option (google.longrunning.operation_info) = {
      response_type: "Task"
      metadata_type: "OperationMetadata"
    };
  }

  // Update the task resource.
  rpc UpdateTask(UpdateTaskRequest) returns (google.longrunning.Operation) {
    option (google.api.http) = {
      patch: "/v1/{task.name=projects/*/locations/*/lakes/*/tasks/*}"
      body: "task"
    };
    option (google.api.method_signature) = "update_mask,task";
    option (google.longrunning.operation_info) = {
      response_type: "Task"
      metadata_type: "OperationMetadata"
    };
  }

  // Delete the task resource.
  rpc DeleteTask(DeleteTaskRequest) returns (google.longrunning.Operation) {
    option (google.api.http) = {
      delete: "/v1/{name=projects/*/locations/*/lakes/*/tasks/*}"
    };
    option (google.api.method_signature) = "name";
    option (google.longrunning.operation_info) = {
      response_type: "google.protobuf.Empty"
      metadata_type: "OperationMetadata"
    };
  }

  // Lists Jobs under the given task.
  rpc ListJobs(ListJobsRequest) returns (ListJobsResponse) {
    option (google.api.http) = {
      get: "/v1/{parent=projects/*/locations/*/lakes/*/tasks/*}/jobs"
    };
    option (google.api.method_signature) = "parent";
  }

  // Get job resource.
  rpc GetJob(GetJobRequest) returns (Job) {
    option (google.api.http) = {
      get: "/v1/{name=projects/*/locations/*/lakes/*/tasks/*/jobs/*}"
    };
    option (google.api.method_signature) = "name";
  }

  // Cancel task jobs.
  rpc CancelJob(CancelJobRequest) returns (google.protobuf.Empty) {
    option (google.api.http) = {
      post: "/v1/{name=projects/*/locations/*/lakes/*/tasks/*/jobs/*}:cancel"
      body: "*"
    };
    option (google.api.method_signature) = "name";
  }
}
```
out:
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
// proto.service: google.cloud.dataplex.v1.DataplexService
// proto.message: google.cloud.dataplex.v1.Lake

package mockdataplex

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/dataplex/v1"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

type dataplexService struct {
	*MockService
	pb.UnimplementedDataplexServiceServer
}

func (s *dataplexService) GetLake(ctx context.Context, req *pb.GetLakeRequest) (*pb.Lake, error) {
	name, err := s.parseLakeName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Lake{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *dataplexService) CreateLake(ctx context.Context, req *pb.CreateLakeRequest) (*longrunningpb.Operation, error) {
	reqName := fmt.Sprintf("%s/lakes/%s", req.GetParent(), req.GetLakeId())
	name, err := s.parseLakeName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := proto.Clone(req.GetLake()).(*pb.Lake)
	obj.Name = fqn
	now := time.Now()
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.Uid = name.LakeID
	obj.State = pb.State_ACTIVE
	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		Target:     name.String(),
		Verb:       "create",
		ApiVersion: "v1",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}

func (s *dataplexService) DeleteLake(ctx context.Context, req *pb.DeleteLakeRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseLakeName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.Lake{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.Now(),
		Target:     name.String(),
		Verb:       "delete",
		ApiVersion: "v1",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return &emptypb.Empty{}, nil
	})
}

type lakeName struct {
	Project  *projects.ProjectData
	Location string
	LakeID   string
}

func (n *lakeName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/lakes/%s", n.Project.ID, n.Location, n.LakeID)
}

// parseLakeName parses a string into a lakeName.
// The expected form is `projects/*/locations/*/lakes/*`.
func (s *MockService) parseLakeName(name string) (*lakeName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "lakes" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &lakeName{
			Project:  project,
			Location: tokens[3],
			LakeID:   tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}



