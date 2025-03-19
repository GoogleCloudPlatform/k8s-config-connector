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
// proto.service: google.cloud.backupdr.v1.BackupDR
// proto.message: google.cloud.backupdr.v1.BackupPlan

package mockbackupdr

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/fields"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/backupdr/v1"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

func (s *BackupDRV1) GetBackupPlan(ctx context.Context, req *pb.GetBackupPlanRequest) (*pb.BackupPlan, error) {
	name, err := s.parseBackupPlanName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.BackupPlan{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "backupPlan %q not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *BackupDRV1) CreateBackupPlan(ctx context.Context, req *pb.CreateBackupPlanRequest) (*longrunningpb.Operation, error) {
	reqName := fmt.Sprintf("%s/backupPlans/%s", req.GetParent(), req.GetBackupPlanId())
	name, err := s.parseBackupPlanName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := proto.Clone(req.GetBackupPlan()).(*pb.BackupPlan)
	obj.Name = fqn
	now := time.Now()
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.State = pb.BackupPlan_CREATING
	obj.BackupVaultServiceAccount = fmt.Sprintf("test-service-account-%d@%s.iam.gserviceaccount.com", name.Project.Number, name.Project.ID)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// LRO
	obj.State = pb.BackupPlan_ACTIVE
	obj.Name = strings.ReplaceAll(obj.Name, name.Project.ID, fmt.Sprintf("%v", name.Project.Number))
	obj.BackupVault = strings.ReplaceAll(obj.BackupVault, name.Project.ID, fmt.Sprintf("%v", name.Project.Number))

	lroMetadata := &pb.OperationMetadata{
		ApiVersion: "v1",
		CreateTime: timestamppb.New(now),
		Target:     name.String(),
		Verb:       "create",
	}
	opPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)

	return s.operations.StartLRO(ctx, opPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())
		return obj, nil
	})

}

func (s *BackupDRV1) DeleteBackupPlan(ctx context.Context, req *pb.DeleteBackupPlanRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseBackupPlanName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	deleted := &pb.BackupPlan{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}
	deleted.State = pb.BackupPlan_DELETING

	lroMetadata := &pb.OperationMetadata{
		ApiVersion: "v1",
		CreateTime: timestamppb.New(now),
		Target:     name.String(),
		Verb:       "delete",
	}
	opPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, opPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())
		return &emptypb.Empty{}, nil
	})
}

func (s *BackupDRV1) UpdateBackupPlan(ctx context.Context, req *pb.UpdateBackupPlanRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseBackupPlanName(req.GetBackupPlan().GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.BackupPlan{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// Required. A list of fields to be updated in this request.
	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask must be provided")
	}
	for _, path := range paths {
		switch path {
		case "description":
			obj.Description = req.GetBackupPlan().GetDescription()
		case "backup_rules":
			obj.BackupRules = req.GetBackupPlan().GetBackupRules()
		case "labels":
			obj.Labels = req.GetBackupPlan().GetLabels()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
		}
	}

	obj.UpdateTime = timestamppb.New(time.Now())
	obj.State = pb.BackupPlan_ACTIVE
	// Generate a new etag each time the object is updated
	obj.Etag = fmt.Sprintf("%X", time.Now().UnixNano())
	obj.BackupVault = strings.ReplaceAll(obj.BackupVault, name.Project.ID, fmt.Sprintf("%v", name.Project.Number))
	obj.Name = strings.ReplaceAll(obj.Name, name.Project.ID, fmt.Sprintf("%v", name.Project.Number))
	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}
	// LRO
	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		ApiVersion: "v1",
		CreateTime: timestamppb.Now(),
		Target:     name.String(),
		Verb:       "update",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}

type backupPlanName struct {
	Project    *projects.ProjectData
	Location   string
	BackupPlan string
}

func (n *backupPlanName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/backupPlans/%s", n.Project.ID, n.Location, n.BackupPlan)
}

// parseBackupPlanName parses a string into an backupPlanName.
// The expected form is `projects/*/locations/*/backupPlans/*`.
func (s *MockService) parseBackupPlanName(name string) (*backupPlanName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "backupPlans" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &backupPlanName{
			Project:    project,
			Location:   tokens[3],
			BackupPlan: tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

</out>

</example>




<example>
<in.proto.message>google.cloud.networkconnectivity.v1.InternalRange</in.proto.message>
<in.proto.message.definition>
message InternalRange {
  option (google.api.resource) = {
    type: "networkconnectivity.googleapis.com/InternalRange"
    pattern: "projects/{project}/locations/{location}/internalRanges/{internal_range}"
    plural: "internalRanges"
    singular: "internalRange"
  };

  // Usage specifies how the range is intended to be used.
  enum Usage {
    // Unspecified.
    USAGE_UNSPECIFIED = 0;

    // Only FOR_EXTERNAL_MANAGED_SERVICES are allowed to use this in Create
    // request. It can be used for "default-route" range for example.
    FOR_EXTERNAL_MANAGED_SERVICES = 1;

    // For Private Service Connect endpoints.
    FOR_PSC_PRODUCER_ENDPOINT = 2;
  }

  // Immutable. Identifier. Name of the InternalRange.
  // Format:
  // projects/{project}/locations/{location}/internalRanges/{internal_range}
  // See: https://google.aip.dev/122#fields-representing-resource-names
  string name = 1 [(google.api.field_behavior) = IDENTIFIER];

  // Time when the InternalRange was created.
  google.protobuf.Timestamp create_time = 2
      [(google.api.field_behavior) = OUTPUT_ONLY];

  // Time when the InternalRange was updated.
  google.protobuf.Timestamp update_time = 3
      [(google.api.field_behavior) = OUTPUT_ONLY];

  // User-provided description for the internal range. It is an optional field
  // provided by user for descriptive purposes. Must be no longer than 400
  // characters.
  string description = 4;

  // Optional. The IP range that this InternalRange defines.
  string ip_cidr_range = 5 [(google.api.field_behavior) = OPTIONAL];

  // Immutable. The type of network where the IP range is used. For private
  // range, it will be used for PSC.
  NetworkType network_type = 6
      [(google.api.field_behavior) = IMMUTABLE, (google.api.field_info) = {
        format: ENUM
      }];

  // User-defined labels.
  map<string, string> labels = 7;

  // The network where the IP range is in, in the form of
  // projects/{project_number}/global/networks/{network_id}. This is a
  // relative resource path in the form of URI and not a free form network name.
  // The location of this network can be different from the location of Internal
  // Range. Routes with the range from this IP range to the same network will be
  // treated as if they are on-premise. They will not be exported to the
  // network.
  optional string network = 8 [
    (google.api.field_behavior) = OPTIONAL,
    (google.api.resource_reference) = { type: "compute.googleapis.com/Network" }
  ];

  // The prefix length of the IP range.
  optional int32 prefix_length = 9 [
    (google.api.field_behavior) = OPTIONAL,
    (google.api.field_info) = { format: INT32 }
  ];

  // Optional. Types of resources that are allowed to use the internal range.
  repeated Usage target_resource_types = 10
      [(google.api.field_behavior) = OPTIONAL];

  // Optional. A list of other other subnetworks that are allowed to use this
  // internal range.
  repeated string peerings = 11 [(google.api.field_behavior) = OPTIONAL];

  // Output only. The type of usage set, but cannot be modified using update.
  Usage usage = 12 [
    (google.api.field_behavior) = OUTPUT_ONLY,
    (google.api.field_info) = { format: ENUM }
  ];

  // Optional. The type of CIDR range.
  InternalRangeType range_type = 13 [(google.api.field_behavior) = OPTIONAL];

  // Output only. The list of resources that refer to this internal range.
  // Resources in this list are not allowed to be deleted.
  repeated string users = 14 [
    (google.api.field_behavior) = OUTPUT_ONLY,
    (google.api.resource_reference) = { type: "*" }
  ];
}
</in.proto.message.definition>
<in.proto.service>google.cloud.networkconnectivity.v1.PolicyBasedRoutingService</in.proto.service>
<in.proto.service.definition>
service PolicyBasedRoutingService {
  option (google.api.default_host) = "networkconnectivity.googleapis.com";
  option (google.api.oauth_scopes) =
      "https://www.googleapis.com/auth/cloud-platform";

  // Lists PolicyBasedRoutes in a given project and location.
  rpc ListPolicyBasedRoutes(ListPolicyBasedRoutesRequest)
      returns (ListPolicyBasedRoutesResponse) {
    option (google.api.http) = {
      get: "/v1/{parent=projects/*/locations/global}/policyBasedRoutes"
    };
    option (google.api.method_signature) = "parent";
  }

  // Gets details of the specified

