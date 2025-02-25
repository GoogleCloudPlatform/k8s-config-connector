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

package mockapigee

import (
	"context"
	"fmt"
	"strings"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/apigee/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

type instanceName struct {
	Organization string
	Instance     string
}

func (n *instanceName) Parent() string {
	return fmt.Sprintf("organizations/%v", n.Organization)
}

func (n *instanceName) String() string {
	return fmt.Sprintf("organizations/%v/instances/%v", n.Organization, n.Instance)
}

// ParseInstanceName parses a string into a instanceName.
// The expected form is organizations/{organization}/instances/{instance}.
func ParseInstanceName(name string) (*instanceName, error) {
	expectedFormat := "organizations/{organization}/instances/{instance}"
	parts := strings.Split(name, "/")
	if len(parts) != 4 || parts[0] != "organizations" || parts[2] != "instances" {
		return nil, fmt.Errorf("name '%s' is not of the form %s", name, expectedFormat)
	}
	return &instanceName{
		Organization: parts[1],
		Instance:     parts[3],
	}, nil
}

type instancesServer struct {
	*MockService
	pb.UnimplementedOrganizationsInstancesServerServer
}

func (s *instancesServer) GetOrganizationsInstance(ctx context.Context, req *pb.GetOrganizationsInstanceRequest) (*pb.GoogleCloudApigeeV1Instance, error) {
	name, err := ParseInstanceName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.GoogleCloudApigeeV1Instance{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "generic::not_found: resource %s not found", fqn)
		}
		return nil, err
	}
	return obj, nil
}

func (s *instancesServer) CreateOrganizationsInstance(ctx context.Context, req *pb.CreateOrganizationsInstanceRequest) (*longrunningpb.Operation, error) {
	reqName := req.Parent + "/instances/" + req.OrganizationsInstance.Name
	name, err := ParseInstanceName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.OrganizationsInstance).(*pb.GoogleCloudApigeeV1Instance)
	obj.Name = req.OrganizationsInstance.Name
	populateDefaultsForOrganizationsInstance(obj)
	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	metadata := &pb.GoogleCloudApigeeV1OperationMetadata{
		OperationType:      "INSERT",
		State:              "IN_PROGRESS",
		TargetResourceName: fqn,
	}
	op, err := s.operations.StartLRO(ctx, req.GetParent(), metadata, func() (proto.Message, error) {
		metadata.Progress = &pb.GoogleCloudApigeeV1OperationMetadataProgress{
			Description: "Succeeded",
			PercentDone: 100,
		}
		metadata.State = "FINISHED"
		result := proto.Clone(obj).(*pb.GoogleCloudApigeeV1Instance)
		populateOutputsForOrganizationsInstance(result)
		s.storage.Update(ctx, fqn, result)
		return result, nil
	})
	return op, err
}

func (s *instancesServer) PatchOrganizationsInstance(ctx context.Context, req *pb.PatchOrganizationsInstanceRequest) (*longrunningpb.Operation, error) {
	name, err := ParseInstanceName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.GoogleCloudApigeeV1Instance{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// Required. A list of fields to be updated in this request.
	paths := req.GetUpdateMask()
	fieldMask, err := fieldmaskpb.New(obj, strings.Split(paths, ",")...)
	if err != nil {
		return nil, err
	}

	for _, path := range fieldMask.GetPaths() {
		switch path {
		case "access_logging_config":
			obj.AccessLoggingConfig = req.OrganizationsInstance.AccessLoggingConfig
		case "consumer_accept_list":
			obj.ConsumerAcceptList = req.OrganizationsInstance.ConsumerAcceptList
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update mask path %q not supported by mockgcp", path)
		}
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	metadata := &pb.GoogleCloudApigeeV1OperationMetadata{
		OperationType:      "PATCH",
		State:              "IN_PROGRESS",
		TargetResourceName: fqn,
	}
	op, err := s.operations.StartLRO(ctx, name.Parent(), metadata, func() (proto.Message, error) {
		metadata.State = "FINISHED"
		result := proto.Clone(obj).(*pb.GoogleCloudApigeeV1Instance)
		return result, nil
	})
	return op, err
}

func (s *instancesServer) DeleteOrganizationsInstance(ctx context.Context, req *pb.DeleteOrganizationsInstanceRequest) (*longrunningpb.Operation, error) {
	name, err := ParseInstanceName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	oldObj := &pb.GoogleCloudApigeeV1Instance{}
	if err := s.storage.Delete(ctx, fqn, oldObj); err != nil {
		return nil, err
	}

	metadata := &pb.GoogleCloudApigeeV1OperationMetadata{
		OperationType:      "DELETE",
		State:              "IN_PROGRESS",
		TargetResourceName: fqn,
	}
	op, err := s.operations.StartLRO(ctx, name.Parent(), metadata, func() (proto.Message, error) {
		metadata.State = "FINISHED"
		return &pb.GoogleCloudApigeeV1Instance{}, nil
	})
	return op, err
}

func populateDefaultsForOrganizationsInstance(obj *pb.GoogleCloudApigeeV1Instance) {
	if len(obj.ConsumerAcceptList) == 0 {
		obj.ConsumerAcceptList = []string{"${projectId}"}
	}
	if obj.IpRange == "" {
		obj.IpRange = "10.39.56.0/22,10.14.0.64/28"
	}
	if obj.PeeringCidrRange == "" {
		obj.PeeringCidrRange = "SLASH_22"
	}
}

func populateOutputsForOrganizationsInstance(obj *pb.GoogleCloudApigeeV1Instance) {
	obj.CreatedAt = 123456789
	obj.Host = "10.39.56.2"
	obj.LastModifiedAt = 123456789
	obj.Port = "443"
	obj.RuntimeVersion = "1-14-0-apigee-4"
	obj.ServiceAttachment = "projects/${projectId}/regions/us-central1/serviceAttachments/apigee-us-central1-abcd"
	obj.State = "ACTIVE"
}
