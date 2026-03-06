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
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
)

type PublicDelegatedPrefixesV1 struct {
	*MockService
	pb.UnimplementedPublicDelegatedPrefixesServer
}

func (s *PublicDelegatedPrefixesV1) Get(ctx context.Context, req *pb.GetPublicDelegatedPrefixeRequest) (*pb.PublicDelegatedPrefix, error) {
	reqName := "projects/" + req.GetProject() + "/regions/" + req.GetRegion() + "/publicDelegatedPrefixes/" + req.GetPublicDelegatedPrefix()
	name, err := s.parsePublicDelegatedPrefixName(reqName)
	if err != nil {
		return nil, err
	}
	return s.getPublicDelegatedPrefix(ctx, name)
}

func (s *PublicDelegatedPrefixesV1) Insert(ctx context.Context, req *pb.InsertPublicDelegatedPrefixeRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/regions/" + req.GetRegion() + "/publicDelegatedPrefixes/" + req.GetPublicDelegatedPrefixResource().GetName()
	name, err := s.parsePublicDelegatedPrefixName(reqName)
	if err != nil {
		return nil, err
	}

	obj := proto.Clone(req.GetPublicDelegatedPrefixResource()).(*pb.PublicDelegatedPrefix)
	obj.Kind = proto.String("compute#publicDelegatedPrefix")
	obj.SelfLink = proto.String(buildComputeSelfLink(ctx, name.String()))
	obj.Region = proto.String(fmt.Sprintf("https://www.googleapis.com/compute/v1/projects/%s/regions/%s", name.Project.ID, name.Region))
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
	return s.startRegionalLRO(ctx, name.Project.ID, name.Region, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *PublicDelegatedPrefixesV1) Delete(ctx context.Context, req *pb.DeletePublicDelegatedPrefixeRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/regions/" + req.GetRegion() + "/publicDelegatedPrefixes/" + req.GetPublicDelegatedPrefix()
	name, err := s.parsePublicDelegatedPrefixName(reqName)
	if err != nil {
		return nil, err
	}

	deleted := &pb.PublicDelegatedPrefix{}
	if err := s.storage.Delete(ctx, name.String(), deleted); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      deleted.Id,
		TargetLink:    deleted.SelfLink,
		OperationType: proto.String("delete"),
		User:          proto.String("user@example.com"),
	}

	return s.startRegionalLRO(ctx, name.Project.ID, name.Region, op, func() (proto.Message, error) {
		return deleted, nil
	})
}

func (s *PublicDelegatedPrefixesV1) Patch(ctx context.Context, req *pb.PatchPublicDelegatedPrefixeRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/regions/" + req.GetRegion() + "/publicDelegatedPrefixes/" + req.GetPublicDelegatedPrefix()
	name, err := s.parsePublicDelegatedPrefixName(reqName)
	if err != nil {
		return nil, err
	}

	obj := &pb.PublicDelegatedPrefix{}
	if err := s.storage.Get(ctx, name.String(), obj); err != nil {
		return nil, err
	}

	// Verify fingerprint if provided
	if req.GetPublicDelegatedPrefixResource().Fingerprint != nil {
		if obj.Fingerprint != nil && req.GetPublicDelegatedPrefixResource().GetFingerprint() != obj.GetFingerprint() {
			return nil, status.Errorf(codes.Aborted, "fingerprint mismatch")
		}
	}

	// Update fields
	if req.GetPublicDelegatedPrefixResource().Description != nil {
		obj.Description = req.GetPublicDelegatedPrefixResource().Description
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
	return s.startRegionalLRO(ctx, name.Project.ID, name.Region, op, func() (proto.Message, error) {
		return obj, nil
	})
}

type GlobalPublicDelegatedPrefixesV1 struct {
	*MockService
	pb.UnimplementedGlobalPublicDelegatedPrefixesServer
}

func (s *GlobalPublicDelegatedPrefixesV1) Get(ctx context.Context, req *pb.GetGlobalPublicDelegatedPrefixeRequest) (*pb.PublicDelegatedPrefix, error) {
	reqName := "projects/" + req.GetProject() + "/global/publicDelegatedPrefixes/" + req.GetPublicDelegatedPrefix()
	name, err := s.parseGlobalPublicDelegatedPrefixName(reqName)
	if err != nil {
		return nil, err
	}
	return s.getPublicDelegatedPrefix(ctx, name)
}

func (s *GlobalPublicDelegatedPrefixesV1) Insert(ctx context.Context, req *pb.InsertGlobalPublicDelegatedPrefixeRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/global/publicDelegatedPrefixes/" + req.GetPublicDelegatedPrefixResource().GetName()
	name, err := s.parseGlobalPublicDelegatedPrefixName(reqName)
	if err != nil {
		return nil, err
	}

	obj := proto.Clone(req.GetPublicDelegatedPrefixResource()).(*pb.PublicDelegatedPrefix)
	obj.Kind = proto.String("compute#publicDelegatedPrefix")
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

func (s *GlobalPublicDelegatedPrefixesV1) Delete(ctx context.Context, req *pb.DeleteGlobalPublicDelegatedPrefixeRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/global/publicDelegatedPrefixes/" + req.GetPublicDelegatedPrefix()
	name, err := s.parseGlobalPublicDelegatedPrefixName(reqName)
	if err != nil {
		return nil, err
	}

	deleted := &pb.PublicDelegatedPrefix{}
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

func (s *GlobalPublicDelegatedPrefixesV1) Patch(ctx context.Context, req *pb.PatchGlobalPublicDelegatedPrefixeRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/global/publicDelegatedPrefixes/" + req.GetPublicDelegatedPrefix()
	name, err := s.parseGlobalPublicDelegatedPrefixName(reqName)
	if err != nil {
		return nil, err
	}

	obj := &pb.PublicDelegatedPrefix{}
	if err := s.storage.Get(ctx, name.String(), obj); err != nil {
		return nil, err
	}

	if req.GetPublicDelegatedPrefixResource().Description != nil {
		obj.Description = req.GetPublicDelegatedPrefixResource().Description
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

type publicDelegatedPrefixName struct {
	Project *projects.ProjectData
	Region  string // "global" if global
	Name    string
}

func (n *publicDelegatedPrefixName) String() string {
	if n.Region == "global" {
		return "projects/" + n.Project.ID + "/global/publicDelegatedPrefixes/" + n.Name
	}
	return "projects/" + n.Project.ID + "/regions/" + n.Region + "/publicDelegatedPrefixes/" + n.Name
}

func (s *MockService) parsePublicDelegatedPrefixName(name string) (*publicDelegatedPrefixName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "regions" && tokens[4] == "publicDelegatedPrefixes" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}
		return &publicDelegatedPrefixName{
			Project: project,
			Region:  tokens[3],
			Name:    tokens[5],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid (expected projects/*/regions/*/publicDelegatedPrefixes/*)", name)
}

func (s *MockService) parseGlobalPublicDelegatedPrefixName(name string) (*publicDelegatedPrefixName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 5 && tokens[0] == "projects" && tokens[2] == "global" && tokens[3] == "publicDelegatedPrefixes" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}
		return &publicDelegatedPrefixName{
			Project: project,
			Region:  "global",
			Name:    tokens[4],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid (expected projects/*/global/publicDelegatedPrefixes/*)", name)
}

func (s *MockService) getPublicDelegatedPrefix(ctx context.Context, name *publicDelegatedPrefixName) (*pb.PublicDelegatedPrefix, error) {
	obj := &pb.PublicDelegatedPrefix{}
	if err := s.storage.Get(ctx, name.String(), obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The resource '%s' was not found", name.String())
		}
		return nil, err
	}
	return obj, nil
}
