// Copyright 2022 Google LLC
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
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/apigee/v1"
	field_mask "google.golang.org/protobuf/types/known/fieldmaskpb"
)

type EnvgroupV1 struct {
	*MockService
	pb.UnimplementedOrganizationsEnvgroupsServerServer
}

func (s *EnvgroupV1) GetOrganizationsEnvgroup(ctx context.Context, req *pb.GetOrganizationsEnvgroupRequest) (*pb.GoogleCloudApigeeV1EnvironmentGroup, error) {
	name, err := ParseEnvgroupName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.GoogleCloudApigeeV1EnvironmentGroup{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "generic::not_found: resource %s not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *EnvgroupV1) CreateOrganizationsEnvgroup(ctx context.Context, req *pb.CreateOrganizationsEnvgroupRequest) (*longrunningpb.Operation, error) {
	reqName := req.Parent + "/envgroups/" + req.OrganizationsEnvgroup.Name
	name, err := ParseEnvgroupName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.OrganizationsEnvgroup).(*pb.GoogleCloudApigeeV1EnvironmentGroup)

	now := timestamppb.Now().GetSeconds()
	obj.CreatedAt = now
	obj.LastModifiedAt = now
	obj.State = "ACTIVE"

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		fmt.Printf("An error occurred with %s %v \n", fqn, err)
		return nil, err
	}

	opMetadata := &pb.GoogleCloudApigeeV1OperationMetadata{
		OperationType:      "INSERT",
		State:              "FINISHED",
		TargetResourceName: fqn,
	}
	opPrefix := fmt.Sprintf("organizations/%s", name.Organization)

	return s.operations.DoneLRO(ctx, opPrefix, opMetadata, func() *pb.GoogleCloudApigeeV1EnvironmentGroup {
		retObj := proto.Clone(obj).(*pb.GoogleCloudApigeeV1EnvironmentGroup)
		retObj.State = "ACTIVE"
		retObj.CreatedAt = 0
		retObj.LastModifiedAt = 0
		return retObj
	}())
}

func (s *EnvgroupV1) PatchOrganizationsEnvgroup(ctx context.Context, req *pb.PatchOrganizationsEnvgroupRequest) (*longrunningpb.Operation, error) {
	name, err := ParseEnvgroupName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.GoogleCloudApigeeV1EnvironmentGroup{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// Required. A list of fields to be updated in this request.
	paths := req.GetUpdateMask()
	fieldMask, err := field_mask.New(obj, strings.Split(paths, ",")...)
	if err != nil {
		return nil, err
	}

	for _, path := range fieldMask.GetPaths() {
		switch path {
		case "hostnames":
			obj.Hostnames = req.OrganizationsEnvgroup.Hostnames
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not supported by mockgcp", path)
		}
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	opMetadata := &pb.GoogleCloudApigeeV1OperationMetadata{
		OperationType:      "UPDATE",
		State:              "FINISHED",
		TargetResourceName: fqn,
	}
	opPrefix := fmt.Sprintf("organizations/%s", name.Organization)

	return s.operations.DoneLRO(ctx, opPrefix, opMetadata, func() *pb.GoogleCloudApigeeV1EnvironmentGroup {
		retObj := proto.Clone(obj).(*pb.GoogleCloudApigeeV1EnvironmentGroup)
		return retObj
	}())
}

func (s *EnvgroupV1) DeleteOrganizationsEnvgroup(ctx context.Context, req *pb.DeleteOrganizationsEnvgroupRequest) (*longrunningpb.Operation, error) {
	name, err := ParseEnvgroupName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	oldObj := &pb.GoogleCloudApigeeV1EnvironmentGroup{}
	if err := s.storage.Delete(ctx, fqn, oldObj); err != nil {
		return nil, err
	}

	opMetadata := &pb.GoogleCloudApigeeV1OperationMetadata{
		OperationType:      "DELETE",
		State:              "FINISHED",
		TargetResourceName: fqn,
	}
	opPrefix := fmt.Sprintf("organizations/%s", name.Organization)
	return s.operations.DoneLRO(ctx, opPrefix, opMetadata, &emptypb.Empty{})
}

type EnvGroupName struct {
	Organization string
	EnvGroupName string
}

func (n *EnvGroupName) String() string {
	return "organizations/" + n.Organization + "/envgroups/" + n.EnvGroupName
}

func (n *EnvGroupName) Parent() string {
	return fmt.Sprintf("organizations/%v", n.Organization)
}

// ParseEnvgroupName parses a string into a envgroupName.
// The expected form is organizations/<projectID>/envgroups/<name>
func ParseEnvgroupName(name string) (*EnvGroupName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 4 && tokens[0] == "organizations" && tokens[2] == "envgroups" {
		name := &EnvGroupName{
			Organization: tokens[1],
			EnvGroupName: tokens[3],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
