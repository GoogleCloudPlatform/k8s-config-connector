// Copyright 2026 Google LLC
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
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

type natAddressName struct {
	Organization string
	Instance     string
	NatAddress   string
}

func (n *natAddressName) Parent() string {
	return fmt.Sprintf("organizations/%v/instances/%v", n.Organization, n.Instance)
}

func (n *natAddressName) String() string {
	return fmt.Sprintf("organizations/%v/instances/%v/natAddresses/%v", n.Organization, n.Instance, n.NatAddress)
}

func (s *natAddressesServer) parseNatAddressName(name string) (*natAddressName, error) {
	expectedFormat := "organizations/{organization}/instances/{instance}/natAddresses/{natAddress}"
	parts := strings.Split(name, "/")
	if len(parts) != 6 || parts[0] != "organizations" || parts[2] != "instances" || parts[4] != "natAddresses" {
		return nil, fmt.Errorf("name '%s' is not of the form %s", name, expectedFormat)
	}
	return &natAddressName{
		Organization: parts[1],
		Instance:     parts[3],
		NatAddress:   parts[5],
	}, nil
}

type natAddressesServer struct {
	*MockService
	pb.UnimplementedOrganizationsInstancesNatAddressesServerServer
}

func (s *natAddressesServer) GetOrganizationsInstancesNatAddresse(ctx context.Context, req *pb.GetOrganizationsInstancesNatAddresseRequest) (*pb.GoogleCloudApigeeV1NatAddress, error) {
	name, err := s.parseNatAddressName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.GoogleCloudApigeeV1NatAddress{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "generic::not_found: resource %s not found", fqn)
		}
		return nil, err
	}
	return obj, nil
}

func (s *natAddressesServer) CreateOrganizationsInstancesNatAddresse(ctx context.Context, req *pb.CreateOrganizationsInstancesNatAddresseRequest) (*longrunningpb.Operation, error) {
	reqName := req.Parent + "/natAddresses/" + req.OrganizationsInstancesNatAddresse.Name
	name, err := s.parseNatAddressName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	instanceName, err := ParseInstanceName(name.Parent())
	if err != nil {
		return nil, err
	}
	orgID := instanceName.Parent()

	obj := proto.CloneOf(req.OrganizationsInstancesNatAddresse)
	populateDefaultsForOrganizationsInstancesNatAddress(obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	metadata := &pb.GoogleCloudApigeeV1OperationMetadata{
		OperationType:      "INSERT",
		State:              "IN_PROGRESS",
		TargetResourceName: fqn,
	}
	op, err := s.operations.StartLRO(ctx, orgID, metadata, func() (proto.Message, error) {
		metadata.State = "FINISHED"
		result := proto.CloneOf(obj)
		populateOutputsForOrganizationsInstancesNatAddress(result)
		s.storage.Update(ctx, fqn, result)
		return result, nil
	})
	return op, err
}

func (s *natAddressesServer) DeleteOrganizationsInstancesNatAddresse(ctx context.Context, req *pb.DeleteOrganizationsInstancesNatAddresseRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseNatAddressName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	instanceName, err := ParseInstanceName(name.Parent())
	if err != nil {
		return nil, err
	}
	orgID := instanceName.Parent()

	oldObj := &pb.GoogleCloudApigeeV1NatAddress{}
	if err := s.storage.Delete(ctx, fqn, oldObj); err != nil {
		return nil, err
	}

	metadata := &pb.GoogleCloudApigeeV1OperationMetadata{
		OperationType:      "DELETE",
		State:              "IN_PROGRESS",
		TargetResourceName: fqn,
	}
	op, err := s.operations.StartLRO(ctx, orgID, metadata, func() (proto.Message, error) {
		metadata.State = "FINISHED"
		return &emptypb.Empty{}, nil
	})
	return op, err
}

func (s *natAddressesServer) ListOrganizationsInstancesNatAddresses(ctx context.Context, req *pb.ListOrganizationsInstancesNatAddressesRequest) (*pb.GoogleCloudApigeeV1ListNatAddressesResponse, error) {
	instanceName, err := ParseInstanceName(req.Parent)
	if err != nil {
		return nil, err
	}

	response := &pb.GoogleCloudApigeeV1ListNatAddressesResponse{}
	findPrefix := instanceName.String() + "/natAddresses/"

	var findKind = (&pb.GoogleCloudApigeeV1NatAddress{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, findKind, storage.ListOptions{
		Prefix: findPrefix,
	}, func(obj proto.Message) error {
		item := obj.(*pb.GoogleCloudApigeeV1NatAddress)
		response.NatAddresses = append(response.NatAddresses, item)
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}

func (s *natAddressesServer) ActivateOrganizationsInstancesNatAddresse(ctx context.Context, req *pb.ActivateOrganizationsInstancesNatAddresseRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseNatAddressName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	instanceName, err := ParseInstanceName(name.Parent())
	if err != nil {
		return nil, err
	}
	orgID := instanceName.Parent()

	obj := &pb.GoogleCloudApigeeV1NatAddress{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	metadata := &pb.GoogleCloudApigeeV1OperationMetadata{
		OperationType:      "UPDATE",
		State:              "IN_PROGRESS",
		TargetResourceName: fqn,
	}
	op, err := s.operations.StartLRO(ctx, orgID, metadata, func() (proto.Message, error) {
		metadata.State = "FINISHED"
		obj.State = "ACTIVE"
		s.storage.Update(ctx, fqn, obj)
		return obj, nil
	})
	return op, err
}

func populateDefaultsForOrganizationsInstancesNatAddress(obj *pb.GoogleCloudApigeeV1NatAddress) {
	if obj.State == "" {
		obj.State = "RESERVED"
	}
	if obj.IpAddress == "" {
		obj.IpAddress = "136.118.74.194"
	}
}

func populateOutputsForOrganizationsInstancesNatAddress(obj *pb.GoogleCloudApigeeV1NatAddress) {
	if obj.State == "" {
		obj.State = "RESERVED"
	}
	if obj.IpAddress == "" {
		obj.IpAddress = "136.118.74.194"
	}
}
