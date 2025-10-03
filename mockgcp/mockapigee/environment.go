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

	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/apigee/v1"
)

type environmentName struct {
	Organization string
	Environment  string
}

func (n *environmentName) String() string {
	return fmt.Sprintf("organizations/%v/environments/%v", n.Organization, n.Environment)
}

// parseEnvironmentName parses a string into a environmentName.
// The expected form is organizations/<org>/environments/<env>
func (s *MockService) parseEnvironmentName(name string) (*environmentName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) != 4 || tokens[0] != "organizations" || tokens[2] != "environments" {
		return nil, status.Errorf(codes.InvalidArgument, "invalid environment name %q: must be of the form organizations/org/environments/env", name)
	}

	return &environmentName{
		Organization: tokens[1],
		Environment:  tokens[3],
	}, nil

}

type environmentsServer struct {
	*MockService
	pb.UnimplementedOrganizationsEnvironmentsServerServer
}

func (s *environmentsServer) GetOrganizationsEnvironment(ctx context.Context, req *pb.GetOrganizationsEnvironmentRequest) (*pb.GoogleCloudApigeeV1Environment, error) {
	name, err := s.parseEnvironmentName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.GoogleCloudApigeeV1Environment{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "resource %s not found", fqn)
		}
		return nil, err
	}
	return obj, nil
}

func (s *environmentsServer) CreateOrganizationsEnvironment(ctx context.Context, req *pb.CreateOrganizationsEnvironmentRequest) (*longrunningpb.Operation, error) {
	reqName := req.Parent + "/environments/" + req.OrganizationsEnvironment.Name
	name, err := s.parseEnvironmentName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.OrganizationsEnvironment).(*pb.GoogleCloudApigeeV1Environment)
	obj.Name = req.OrganizationsEnvironment.Name

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	opMetadata := &pb.GoogleCloudApigeeV1OperationMetadata{
		OperationType:      "CREATE",
		State:              "FINISHED",
		TargetResourceName: fqn,
	}
	opPrefix := fmt.Sprintf("organizations/%s", name.Organization)

	return s.operations.DoneLRO(ctx, opPrefix, opMetadata, func() *pb.GoogleCloudApigeeV1Environment {
		obj.Name = name.Environment
		return obj
	}())
}

func (s *environmentsServer) UpdateOrganizationsEnvironment(ctx context.Context, req *pb.UpdateOrganizationsEnvironmentRequest) (*pb.GoogleCloudApigeeV1Environment, error) {
	name, err := s.parseEnvironmentName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.GoogleCloudApigeeV1Environment{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// All fields should be passed by callergst
	updated := ProtoClone(req.GetOrganizationsEnvironment())
	updated.Name = obj.Name

	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}

	return updated, nil
}

func (s *environmentsServer) DeleteOrganizationsEnvironment(ctx context.Context, req *pb.DeleteOrganizationsEnvironmentRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseEnvironmentName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	oldObj := &pb.GoogleCloudApigeeV1Environment{}
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
