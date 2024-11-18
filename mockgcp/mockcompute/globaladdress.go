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
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
)

type GlobalAddressesV1 struct {
	*MockService
	pb.UnimplementedGlobalAddressesServer
}

func (s *GlobalAddressesV1) Get(ctx context.Context, req *pb.GetGlobalAddressRequest) (*pb.Address, error) {
	reqName := "projects/" + req.GetProject() + "/global" + "/addresses/" + req.GetAddress()
	name, err := s.parseGlobalAddressName(reqName)
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

func (s *GlobalAddressesV1) Insert(ctx context.Context, req *pb.InsertGlobalAddressRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/global" + "/addresses/" + req.GetAddressResource().GetName()
	name, err := s.parseGlobalAddressName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	id := s.generateID()

	obj := proto.Clone(req.GetAddressResource()).(*pb.Address)
	obj.SelfLink = PtrTo("https://www.googleapis.com/compute/v1/" + name.String())
	obj.CreationTimestamp = PtrTo(s.nowString())
	obj.Id = &id
	obj.Kind = PtrTo("compute#address")
	if obj.Address == nil {
		obj.Address = PtrTo("8.8.8.8")
	}
	if obj.LabelFingerprint == nil {
		obj.LabelFingerprint = PtrTo(computeFingerprint(obj))
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("insert"),
		User:          PtrTo("user@example.com"),
	}
	return s.startGlobalLRO(ctx, name.Project.ID, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *GlobalAddressesV1) Delete(ctx context.Context, req *pb.DeleteGlobalAddressRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/global" + "/addresses/" + req.GetAddress()

	name, err := s.parseGlobalAddressName(reqName)
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
	return s.startGlobalLRO(ctx, name.Project.ID, op, func() (proto.Message, error) {
		return deleted, nil
	})
}

func (s *GlobalAddressesV1) SetLabels(ctx context.Context, req *pb.SetLabelsGlobalAddressRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/global" + "/addresses/" + req.GetResource()
	name, err := s.parseGlobalAddressName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Address{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	obj.Labels = req.GetGlobalSetLabelsRequestResource().GetLabels()

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return s.newLRO(ctx, name.Project.ID)
}

type globalAddressName struct {
	Project *projects.ProjectData
	Name    string
}

func (n *globalAddressName) String() string {
	return "projects/" + n.Project.ID + "/global" + "/addresses/" + n.Name
}

// parseGlobalAddressName parses a string into a globalAddressName.
// The expected form is `projects/*/global/address/*`.
func (s *MockService) parseGlobalAddressName(name string) (*globalAddressName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 5 && tokens[0] == "projects" && tokens[2] == "global" && tokens[3] == "addresses" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &globalAddressName{
			Project: project,
			Name:    tokens[4],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
