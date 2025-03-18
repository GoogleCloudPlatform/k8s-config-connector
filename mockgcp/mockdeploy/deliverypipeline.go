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
// proto.service: google.cloud.deploy.v1.CloudDeploy
// proto.message: google.cloud.deploy.v1.DeliveryPipeline

package mockclouddeploy

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

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/deploy/v1"
	"github.com/google/uuid"
)

type cloudDeploy struct {
	*MockService
	pb.UnimplementedCloudDeployServer
}

func (s *cloudDeploy) GetDeliveryPipeline(ctx context.Context, req *pb.GetDeliveryPipelineRequest) (*pb.DeliveryPipeline, error) {
	name, err := s.parseDeliveryPipelineName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.DeliveryPipeline{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "deliveryPipeline %q not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *cloudDeploy) CreateDeliveryPipeline(ctx context.Context, req *pb.CreateDeliveryPipelineRequest) (*longrunning.Operation, error) {
	reqName := fmt.Sprintf("%s/deliveryPipelines/%s", req.Parent, req.DeliveryPipelineId)
	name, err := s.parseDeliveryPipelineName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := proto.Clone(req.DeliveryPipeline).(*pb.DeliveryPipeline)
	obj.Name = fqn

	obj.Uid = uuid.NewString()
	obj.CreateTime = timestamppb.New(time.Now())
	obj.UpdateTime = timestamppb.New(time.Now())

	if obj.Pipeline == nil {
		obj.Pipeline = &pb.DeliveryPipeline_SerialPipeline{
			SerialPipeline: &pb.SerialPipeline{},
		}
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// By default, immediately finish the LRO with success.
	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(time.Now()),
		Target:     name.String(),
		Verb:       "create",
		ApiVersion: "v1",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}

func (s *cloudDeploy) UpdateDeliveryPipeline(ctx context.Context, req *pb.UpdateDeliveryPipelineRequest) (*longrunning.Operation, error) {
	name, err := s.parseDeliveryPipelineName(req.DeliveryPipeline.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.DeliveryPipeline{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	req.DeliveryPipeline.Uid = obj.GetUid()
	req.DeliveryPipeline.CreateTime = obj.GetCreateTime()
	req.DeliveryPipeline.UpdateTime = timestamppb.New(time.Now())

	// Apply the update mask to the object.
	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask must be provided")
	}
	for _, path := range paths {
		switch path {
		case "description":
			obj.Description = req.DeliveryPipeline.Description
		case "annotations":
			obj.Annotations = req.DeliveryPipeline.Annotations
		case "labels":
			obj.Labels = req.DeliveryPipeline.Labels
		case "serial_pipeline", "serialPipeline":
			obj.Pipeline = &pb.DeliveryPipeline_SerialPipeline{
				SerialPipeline: req.DeliveryPipeline.GetSerialPipeline(),
			}
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
		}
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(time.Now()),
		Target:     name.String(),
		Verb:       "update",
		ApiVersion: "v1",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())
		return obj, nil
	})
}

func (s *cloudDeploy) DeleteDeliveryPipeline(ctx context.Context, req *pb.DeleteDeliveryPipelineRequest) (*longrunning.Operation, error) {
	name, err := s.parseDeliveryPipelineName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.DeliveryPipeline{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	// By default, immediately finish the LRO with success.
	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(time.Now()),
		Target:     name.String(),
		Verb:       "delete",
		ApiVersion: "v1",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())
		return &emptypb.Empty{}, nil
	})
}

type deliveryPipelineName struct {
	Project          *projects.ProjectData
	Location         string
	DeliveryPipeline string
}

func (n *deliveryPipelineName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/deliveryPipelines/%s", n.Project.ID, n.Location, n.DeliveryPipeline)
}

// parseDeliveryPipelineName parses a string into a deliveryPipelineName.
// The expected form is `projects/*/locations/*/deliveryPipelines/*`.
func (s *MockService) parseDeliveryPipelineName(name string) (*deliveryPipelineName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "deliveryPipelines" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &deliveryPipelineName{
			Project:          project,
			Location:         tokens[3],
			DeliveryPipeline: tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

</out>

</example>




<example>
<in.proto.message>google.cloud.resourcemanager.v3.TagValue</in.proto.message>
<in.proto.message.definition>
message TagValue {
  option (google.api.resource) = {
    type: "cloudresourcemanager.googleapis.com/TagValue"
    pattern: "tagValues/{tag_value}"
  };

  // Output only. Resource name for TagValue in the format `tagValues/456`.
  string name = 1 [(google.api.field_behavior) = OUTPUT_ONLY];

  // Immutable. Resource name for TagValue's parent TagKey in the format
  // `tagKeys/123`.
  string parent = 2
      [(google.api.field_behavior) = IMMUTABLE, (google.api.resource_reference) = {
        type: "cloudresourcemanager.googleapis.com/TagKey"
      }];

  // Required. Immutable. User-assigned short name for TagValue. The short name
  // should be unique for TagValues within the same parent TagKey.
  //
  // The short name must be 63 characters or less, beginning and ending with
  // an alphanumeric character ([a-z0-9A-Z]) with dashes (-), underscores (_),
  // dots (.), and alphanumeric characters in between.
  string short_name = 3 [(google.api.field_behavior) = IMMUTABLE, (google.api.field_behavior) = REQUIRED];

  // Output only. The namespaced name of the TagValue. Can be in the form
  // `{organization_id}/{tag_key_short_name}/{tag_value_short_name}` or
  // `{project_id}/{tag_key_short_name}/{tag_value_short_name}` or
  // `{project_number}/{tag_key_short_name}/{tag_value_short_name}`.
  string namespaced_name = 4 [(google.api.field_behavior) = OUTPUT_ONLY];

  // Optional. User-assigned description of the TagValue.
  // Must not exceed 256 characters.
  //
  // Read-write.
  string description = 5 [(google.api.field_behavior) = OPTIONAL];

  // Output only. Creation time.
  google.protobuf.Timestamp create_time = 6
      [(google.api.field_behavior) = OUTPUT_ONLY];

  // Output only. Update time.
  google.protobuf.Timestamp update_time = 7
      [(google.api.field_behavior) = OUTPUT_ONLY];

  // Optional. A string of up to 256 characters, containing any Unicode character.
  // This field is deprecated and will be removed in future versions of Cloud
  // Resource Manager.
  string etag = 8 [(google.api.field_behavior) = OPTIONAL];
}
</in.proto.message.definition>
<in.proto.service>google.cloud.resourcemanager.v3.TagValues</in.proto.service>
<in.proto.service.definition>
service TagValues {
  option (google.api.default_host) = "cloudresourcemanager.googleapis.com";
  option (google.api.oauth_scopes) =
      "https://www.googleapis.com/auth/cloud-platform,"
      "https://www.googleapis.com/auth/cloud-platform.read-only";

  // Lists all TagValues for a specific TagKey.
  rpc ListTagValues(ListTagValuesRequest) returns (ListTagValuesResponse) {
    option (google.api.http) = {
      get: "/v3/tagValues"
    };
    option (google.api.method_signature) = "parent";
  }

  // Retrieves a TagValue. This method will return `PERMISSION_DENIED` if the
  // value does not exist or the user does not have permission to view it.
  rpc GetTagValue(GetTagValueRequest) returns (TagValue) {
    option (google.api.http) = {
      get: "/v3/{name=tagValues/*}"
    };
    option (google.api.method_signature) = "name";
  }

  // Creates a TagValue as a child of the specified TagKey. If a another
  // request with the same parameters is sent while the original request is in
  // process the second request will receive an error. A maximum of 1000
  // TagValues can exist under a TagKey at any given time.
  rpc CreateTagValue(CreateTagValueRequest)
      returns (google.longrunning.Operation) {
    option (google.api.http) = {
      post: "/v3/tagValues"
      body: "tag_value"
    };
    option (google.api.method_signature) = "tag_value";
    option (google.longrunning.operation_info) = {
      response_type: "TagValue"
      metadata_type: "CreateTagValueMetadata"
    };
  }

  // Updates the attributes of the TagValue resource.
  rpc UpdateTagValue(UpdateTagValueRequest)
      returns (google.longrunning.Operation) {
    option (google.api.http) = {
      patch: "/v3/{tag_value.name=tagValues/*}"
      body: "tag_value"
    };
    option (google.api.method_signature) = "tag_value,update_mask";
    option (google.longrunning.operation_info) = {
      response_type: "TagValue"
      metadata_type: "UpdateTagValueMetadata"
    };
  }

  // Deletes a TagValue. The TagValue cannot have any bindings when it is
  // deleted.
  rpc DeleteTagValue(DeleteTagValueRequest)
      returns (google.longrunning.Operation) {
    option (google.api.http) = {
      delete: "/v3/{name=tagValues/*}"
    };
    option (google.api.method_signature) = "name";
    option (google.longrunning.operation_info) = {
      response_type: "TagValue"
      metadata_type: "DeleteTagValueMetadata"
    };
  }

  // Gets the access control policy for a TagValue. The returned policy may be
  // empty if no such policy or resource exists. The `resource` field should
  // be the TagValue's resource name. For example: `tagValues/1234`.
  // The caller must have the
  // `cloudresourcemanager.googleapis.com/tagValues.getIamPolicy` permission on
  // the identified TagValue to get the access control policy.
  rpc GetIamPolicy(google.iam.v1.GetIamPolicyRequest)
      returns (google.iam.v1.Policy) {
    option (google.api.http) = {
      post: "/v3/{resource=tagValues/*}:getIamPolicy"
      body: "*"
    };
    option (google.api.method_signature) = "resource";
  }

  // Sets the access control policy on a TagValue, replacing any existing
  // policy. The `resource` field should be the TagValue's resource name. For
  // example: `tagValues/1234`.
  // The caller must have `resourcemanager.tagValues.setIamPolicy` permission
  // on the identified tagValue.
  rpc SetIamPolicy(google.iam.v1.SetIamPolicyRequest)
      returns (google.iam.v1.Policy) {
    option (google.api.http) = {
      post: "/v3/{resource=tagValues/*}:setIamPolicy"
      body: "*"
    };
    option (google.api.method_signature) = "resource,policy";
  }

  // Returns permissions that a caller has on the specified TagValue.
  // The `resource` field should be the TagValue's resource name. For example:
  // `tagValues/1234`.
  //
  // There are no permissions required for making this API call.
  rpc TestIamPermissions(google.iam.v1.TestIamPermissionsRequest)
      returns (google.iam.v1.TestIamPermissionsResponse) {
    option (google.api.http) = {
      post: "/v3/{resource=tagValues/*}:testIamPermissions"
      body: "*"
    };
    option (google.api.method_signature) = "resource,permissions";
  }

  // Gets the effective policy for a TagValue. This is the result of merging
  // the TagValue's tagValuePolicy with the policies inherited from higher in
  // the hierarchy.
  rpc GetEffectivePolicy(GetEffectivePolicyRequest)
      returns (google.iam.v1.Policy) {
    option (google.api.http) = {
      post: "/v3/{resource=tagValues/*}:getEffectivePolicy"
      body: "*"
    };
    option (google.api.method_signature) = "resource";
  }
}
</in.proto.service.definition>
<out>
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
// proto.service: google.cloud.resourcemanager.v3.TagValues
// proto.message: google.cloud.resourcemanager.v3.TagValue

package mockresourcemanager

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/resourcemanager/v3"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

func (s *MockService) GetTagValue(ctx context.Context, req *pb.GetTagValueRequest) (*pb.TagValue, error) {
	name, err := parseTagValueName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.TagValue{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *MockService) CreateTagValue(ctx context.Context, req *pb.CreateTagValueRequest) (*longrunningpb.Operation, error) {
	reqName := fmt.Sprintf("tagValues/%s", req.GetTagValue().GetShortName())
	name, err := parseTagValueName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.GetTagValue()).(*pb.TagValue)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(time.Now())
	obj.UpdateTime = timestamppb.New(time.Now())

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return s.operations.NewLRO(ctx)
}

type tagValueName struct {
	TagValueID string
}

func (n *tagValueName) String() string {
	return fmt.Sprintf("tagValues/%s", n.TagValueID)
}

// parseTagValueName parses a string into a TagValueName.
// The expected form is `tagValues/*`.
func parseTagValueName(name string) (*tagValueName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 2 && tokens[0] == "tagValues" {
		name := &tagValueName{
			TagValueID: tokens[1],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

</out>

</example>




Can you complete the item? Don't output any additional commentary.
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
// proto.service: google.cloud.dataproc.v1.AutoscalingPolicyService
// proto.message: google.cloud.dataproc.v1.AutoscalingPolicy

package mockdataproc

import (
	"context"
	"fmt"
	"regexp"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "cloud.google.com/go/dataproc/v2/apiv1/dataprocpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

type autoscalingPolicyServiceServer struct {
	*MockService
	pb.UnimplementedAutoscalingPolicyServiceServer
}

func (s *autoscalingPolicyServiceServer) CreateAutoscalingPolicy(ctx context.Context, req *pb.CreateAutoscalingPolicyRequest) (*pb.AutoscalingPolicy, error) {
	reqName := fmt.Sprintf("%s/autoscalingPolicies/%s", req.GetParent(), req.Policy.Id)
	name, err := s.parseAutoscalingPolicyName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.GetPolicy()).(*pb.AutoscalingPolicy)
	obj.Name = fqn
	s.populateDefaultsForAutoscalingPolicy(obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *autoscalingPolicyServiceServer) GetAutoscalingPolicy(ctx context.Context, req *pb.GetAutoscalingPolicyRequest) (*pb.AutoscalingPolicy, error) {
	name, err := s.parseAutoscalingPolicyName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.AutoscalingPolicy{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *autoscalingPolicyServiceServer) DeleteAutoscalingPolicy(ctx context.Context, req *pb.DeleteAutoscalingPolicyRequest) (*emptypb.Empty, error) {
	name, err := s.parseAutoscalingPolicyName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.AutoscalingPolicy{}
	if err := s.storage.Delete(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *autoscalingPolicyServiceServer) UpdateAutoscalingPolicy(ctx context.Context, req *pb.UpdateAutoscalingPolicyRequest) (*pb.AutoscalingPolicy, error) {
	name, err := s.parseAutoscalingPolicyName(req.Policy.Id)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := req.Policy
	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *autoscalingPolicyServiceServer) ListAutoscalingPolicies(ctx context.Context, req *pb.ListAutoscalingPoliciesRequest) (*pb.ListAutoscalingPoliciesResponse, error) {
	response := &pb.ListAutoscalingPoliciesResponse{}

	AutoscalingPolicyKind := (&pb.AutoscalingPolicy{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, AutoscalingPolicyKind, storage.ListOptions{}, func(obj proto.Message) error {
		autoScalingPolicy := obj.(*pb.AutoscalingPolicy)
		if strings.HasPrefix(autoScalingPolicy.GetName(), req.Parent) {
			response.Policies = append(response.Policies, autoScalingPolicy)
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return response, nil
}

func (s *autoscalingPolicyServiceServer) populateDefaultsForAutoscalingPolicy(obj *pb.AutoscalingPolicy) {
	if obj.GetBasicAlgorithm() == nil {
		obj.Algorithm = &pb.AutoscalingPolicy_BasicAlgorithm{}
	}
	if obj.GetBasicAlgorithm().CooldownPeriod == nil {
		obj.GetBasicAlgorithm().CooldownPeriod = durationpb.New(2 * time.Minute)
	}
	if obj.GetBasicAlgorithm().GetYarnConfig() == nil {
		obj.GetBasicAlgorithm().Config = &pb.BasicAutoscalingAlgorithm_YarnConfig{}
	}
	if obj.GetBasicAlgorithm().GetYarnConfig().GracefulDecommissionTimeout == nil {
		obj.GetBasicAlgorithm().GetYarnConfig().GracefulDecommissionTimeout = durationpb.New(24 * time.Hour)
	}
	if obj.GetBasicAlgorithm().GetYarnConfig().ScaleDownFactor == 0 {
		obj.GetBasicAlgorithm().GetYarnConfig().ScaleDownFactor = 1
	}
	if obj.GetBasicAlgorithm().GetYarnConfig().ScaleDownMinWorkerFraction == 0 {
		obj.GetBasicAlgorithm().GetYarnConfig().ScaleDownMinWorkerFraction = 1
	}
	if obj.GetBasicAlgorithm().GetYarnConfig().ScaleUpFactor == 0 {
		obj.GetBasicAlgorithm().GetYarnConfig().ScaleUpFactor = 0.5
	}
	if obj.GetBasicAlgorithm().GetYarnConfig().ScaleUpMinWorkerFraction == 0 {
		obj.GetBasicAlgorithm().GetYarnConfig().ScaleUpMinWorkerFraction = 0.5
	}

	if obj.WorkerConfig == nil {
		obj.WorkerConfig = &pb.InstanceGroupAutoscalingPolicyConfig{}
	}
	if obj.WorkerConfig.MaxInstances == 0 {
		obj.WorkerConfig.MaxInstances = 5
	}

	if obj.SecondaryWorkerConfig == nil {
		obj.SecondaryWorkerConfig = &pb.InstanceGroupAutoscalingPolicyConfig{}
	}
	if obj.SecondaryWorkerConfig.MaxInstances == 0 {
		obj.SecondaryWorkerConfig.MaxInstances = 1
	}
}

type autoscalingPolicyName struct {
	Project            *projects.ProjectData
	Region             string
	AutoscalingPolicy  string
	AutoscalingPolicy2 string
}

func (n *autoscalingPolicyName) String() string {
	return fmt.Sprintf("projects/%s/regions/%s/autoscalingPolicies/%s", n.Project.ID, n.Region, n.AutoscalingPolicy)
}

// parseAutoscalingPolicyName parses a string into an AutoscalingPolicyName.
// The expected form is `projects/*/regions/*/autoscalingPolicies/*`.
func (s *MockService) parseAutoscalingPolicyName(name string) (*autoscalingPolicyName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "regions" && tokens[4] == "autoscalingPolicies" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &autoscalingPolicyName{
			Project:           project,
			Region:            tokens[3],
			AutoscalingPolicy: tokens[5],
		}

		return name, nil
	}

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "autoscalingPolicies" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &autoscalingPolicyName{
			Project:           project,
			Region:            tokens[3],
			AutoscalingPolicy: tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

// buildAutoscalingPolicyName builds a AutoscalingPolicyName from the components.
func (s *MockService) buildAutoscalingPolicyName(projectName, region, cluster string) (*autoscalingPolicyName, error) {
	project, err := s.Projects.GetProjectByID(projectName)
	if err != nil {
		return nil, err
	}

	return &autoscalingPolicyName{
		Project:           project,
		Region:            region,
		AutoscalingPolicy: cluster,
	}, nil
}
```
</out>

</example>

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
// proto.service: google.cloud.resourcemanager.v3.TagValues
// proto.message: google.cloud.resourcemanager.v3.TagValue

package mockresourcemanager

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/resourcemanager/v3"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

func (s *MockService) GetTagValue(ctx context.Context, req *pb.GetTagValueRequest) (*pb.TagValue, error) {
	name, err := parseTagValueName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.TagValue{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *MockService) CreateTagValue(ctx context.Context, req *pb.CreateTagValueRequest) (*longrunningpb.Operation, error) {
	reqName := fmt.Sprintf("tagValues/%s", req.GetTagValue().GetShortName())
	name, err := parseTagValueName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.GetTagValue()).(*pb.TagValue)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(time.Now())
	obj.UpdateTime = timestamppb.New(time.Now())

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return s.operations.NewLRO(ctx)
}

type tagValueName struct {
	TagValueID string
}

func (n *tagValueName) String() string {
	return fmt.Sprintf("tagValues/%s", n.TagValueID)
}

// parseTagValueName parses a string into a TagValueName.
// The expected form is `tagValues/*`.
func parseTagValueName(name string) (*tagValueName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 2 && tokens[0] == "tagValues" {

