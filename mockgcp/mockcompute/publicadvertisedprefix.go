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

type PublicAdvertisedPrefixesV1 struct {
	*MockService
	pb.UnimplementedPublicAdvertisedPrefixesServer
}

func (s *PublicAdvertisedPrefixesV1) Get(ctx context.Context, req *pb.GetPublicAdvertisedPrefixeRequest) (*pb.PublicAdvertisedPrefix, error) {
	reqName := "projects/" + req.GetProject() + "/global/publicAdvertisedPrefixes/" + req.GetPublicAdvertisedPrefix()
	name, err := s.parsePublicAdvertisedPrefixName(reqName)
	if err != nil {
		return nil, err
	}
	return s.getPublicAdvertisedPrefix(ctx, name)
}

func (s *PublicAdvertisedPrefixesV1) Insert(ctx context.Context, req *pb.InsertPublicAdvertisedPrefixeRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/global/publicAdvertisedPrefixes/" + req.GetPublicAdvertisedPrefixResource().GetName()
	name, err := s.parsePublicAdvertisedPrefixName(reqName)
	if err != nil {
		return nil, err
	}

	obj := proto.Clone(req.GetPublicAdvertisedPrefixResource()).(*pb.PublicAdvertisedPrefix)
	obj.Kind = proto.String("compute#publicAdvertisedPrefix")
	obj.SelfLink = proto.String(buildComputeSelfLink(ctx, name.String()))
	obj.CreationTimestamp = proto.String(s.nowString())
	id := s.generateID()
	obj.Id = &id

	if err := s.storage.Create(ctx, name.String(), obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: proto.String("insert"),
		User:          proto.String("user@example.com"),
	}
	return s.startGlobalLRO(ctx, name.Project.ID, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *PublicAdvertisedPrefixesV1) Delete(ctx context.Context, req *pb.DeletePublicAdvertisedPrefixeRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/global/publicAdvertisedPrefixes/" + req.GetPublicAdvertisedPrefix()
	name, err := s.parsePublicAdvertisedPrefixName(reqName)
	if err != nil {
		return nil, err
	}

	deleted := &pb.PublicAdvertisedPrefix{}
	if err := s.storage.Delete(ctx, name.String(), deleted); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      deleted.Id,
		TargetLink:    deleted.SelfLink,
		OperationType: proto.String("delete"),
		User:          proto.String("user@example.com"),
	}

	return s.startGlobalLRO(ctx, name.Project.ID, op, func() (proto.Message, error) {
		return deleted, nil
	})
}

func (s *PublicAdvertisedPrefixesV1) Patch(ctx context.Context, req *pb.PatchPublicAdvertisedPrefixeRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/global/publicAdvertisedPrefixes/" + req.GetPublicAdvertisedPrefix()
	name, err := s.parsePublicAdvertisedPrefixName(reqName)
	if err != nil {
		return nil, err
	}

	obj := &pb.PublicAdvertisedPrefix{}
	if err := s.storage.Get(ctx, name.String(), obj); err != nil {
		return nil, err
	}

	if req.GetPublicAdvertisedPrefixResource().Description != nil {
		obj.Description = req.GetPublicAdvertisedPrefixResource().Description
	}

	if err := s.storage.Update(ctx, name.String(), obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: proto.String("patch"),
		User:          proto.String("user@example.com"),
	}
	return s.startGlobalLRO(ctx, name.Project.ID, op, func() (proto.Message, error) {
		return obj, nil
	})
}

type publicAdvertisedPrefixName struct {
	Project *projects.ProjectData
	Name    string
}

func (n *publicAdvertisedPrefixName) String() string {
	return "projects/" + n.Project.ID + "/global/publicAdvertisedPrefixes/" + n.Name
}

func (s *MockService) parsePublicAdvertisedPrefixName(name string) (*publicAdvertisedPrefixName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 5 && tokens[0] == "projects" && tokens[2] == "global" && tokens[3] == "publicAdvertisedPrefixes" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}
		return &publicAdvertisedPrefixName{
			Project: project,
			Name:    tokens[4],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid (expected projects/*/global/publicAdvertisedPrefixes/*)", name)
}

func (s *MockService) getPublicAdvertisedPrefix(ctx context.Context, name *publicAdvertisedPrefixName) (*pb.PublicAdvertisedPrefix, error) {
	obj := &pb.PublicAdvertisedPrefix{}
	if err := s.storage.Get(ctx, name.String(), obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The resource '%s' was not found", name.String())
		}
		return nil, err
	}
	return obj, nil
}
