// Copyright 2023 Google LLC
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

package mockcompute

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

type RegionalAddressesV1 struct {
	*MockService
	pb.UnimplementedAddressesServer
}

func (s *RegionalAddressesV1) Get(ctx context.Context, req *pb.GetAddressRequest) (*pb.Address, error) {
	reqName := "projects/" + req.GetProject() + "/regions/" + req.GetRegion() + "/addresses/" + req.GetAddress()
	name, err := s.parseRegionalAddressName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Address{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *RegionalAddressesV1) List(ctx context.Context, req *pb.ListAddressesRequest) (*pb.AddressList, error) {
	reqName := "projects/" + req.GetProject() + "/regions/" + req.GetRegion() + "/addresses/" + "placeholder"
	name, err := s.parseRegionalAddressName(reqName)
	if err != nil {
		return nil, err
	}

	if req.GetFilter() != "" {
		return nil, fmt.Errorf("filter %q not implemented by mockgcp", req.GetFilter())
	}

	findPrefix := strings.TrimSuffix(name.String(), "placeholder")

	response := &pb.AddressList{}
	response.Id = PtrTo("0123456789")
	response.Kind = PtrTo("compute#addressList")
	response.SelfLink = PtrTo(buildComputeSelfLink(ctx, strings.TrimSuffix(findPrefix, "/")))

	findKind := (&pb.Address{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, findKind, storage.ListOptions{Prefix: findPrefix}, func(obj proto.Message) error {
		address := obj.(*pb.Address)
		response.Items = append(response.Items, address)
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}

func (s *RegionalAddressesV1) Insert(ctx context.Context, req *pb.InsertAddressRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/regions/" + req.GetRegion() + "/addresses/" + req.GetAddressResource().GetName()
	name, err := s.parseRegionalAddressName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	id := s.generateID()

	obj := proto.Clone(req.GetAddressResource()).(*pb.Address)
	obj.SelfLink = PtrTo(buildComputeSelfLink(ctx, fqn))
	obj.CreationTimestamp = PtrTo(s.nowString())
	obj.Id = &id
	obj.Kind = PtrTo("compute#address")
	obj.Region = PtrTo(fmt.Sprintf("https://www.googleapis.com/compute/v1/projects/%s/regions/%s", name.Project.ID, name.Region))

	s.populateDefaults(obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("insert"),
		User:          PtrTo("user@example.com"),
	}
	return s.startRegionalLRO(ctx, name.Project.ID, name.Region, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *RegionalAddressesV1) populateDefaults(obj *pb.Address) {
	if obj.Address == nil {
		obj.Address = PtrTo("8.8.8.8")
	}
	if obj.AddressType == nil {
		obj.AddressType = PtrTo("EXTERNAL")
	}
	if obj.Description == nil {
		obj.Description = PtrTo("")
	}
	if obj.NetworkTier == nil {
		obj.NetworkTier = PtrTo("PREMIUM")
	}
	if obj.Status == nil {
		obj.Status = PtrTo("RESERVED")
	}

	obj.LabelFingerprint = PtrTo(labelsFingerprint(obj.Labels))
}

func (s *RegionalAddressesV1) Delete(ctx context.Context, req *pb.DeleteAddressRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/regions/" + req.GetRegion() + "/addresses/" + req.GetAddress()
	name, err := s.parseRegionalAddressName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.Address{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      deleted.Id,
		TargetLink:    deleted.SelfLink,
		OperationType: PtrTo("delete"),
		User:          PtrTo("user@example.com"),
	}
	return s.startRegionalLRO(ctx, name.Project.ID, name.Region, op, func() (proto.Message, error) {
		return deleted, nil
	})
}

func (s *RegionalAddressesV1) SetLabels(ctx context.Context, req *pb.SetLabelsAddressRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/regions/" + req.GetRegion() + "/addresses/" + req.GetResource()
	name, err := s.parseRegionalAddressName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Address{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	obj.Labels = req.GetRegionSetLabelsRequestResource().GetLabels()
	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return s.newLRO(ctx, name.Project.ID)
}

type regionalAddressName struct {
	Project *projects.ProjectData
	Region  string
	Name    string
}

func (n *regionalAddressName) String() string {
	return "projects/" + n.Project.ID + "/regions/" + n.Region + "/addresses/" + n.Name
}

// parseAddressName parses a string into a addressName.
// The expected form is `projects/*/regions/*/address/*`.
func (s *MockService) parseRegionalAddressName(name string) (*regionalAddressName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "regions" && tokens[4] == "addresses" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &regionalAddressName{
			Project: project,
			Region:  tokens[3],
			Name:    tokens[5],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
