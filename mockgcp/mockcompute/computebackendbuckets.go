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

// +tool:mockgcp-support
// proto.service: google.cloud.compute.v1.BackendBuckets
// proto.message: google.cloud.compute.v1.BackendBucket

package mockcompute

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
)

type backendBuckets struct {
	*MockService
	pb.UnimplementedBackendBucketsServer
}

func (s *backendBuckets) Get(ctx context.Context, req *pb.GetBackendBucketRequest) (*pb.BackendBucket, error) {
	reqName := fmt.Sprintf("projects/%s/global/backendBuckets/%s", req.GetProject(), req.GetBackendBucket())
	name, err := s.parseBackendBucketName(reqName)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.BackendBucket{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "BackendBucket %q not found", name)
		}
		return nil, err
	}

	return obj, nil
}

func (s *backendBuckets) Insert(ctx context.Context, req *pb.InsertBackendBucketRequest) (*pb.Operation, error) {
	reqName := fmt.Sprintf("projects/%s/global/backendBuckets/%s", req.GetProject(), req.GetBackendBucketResource().GetName())
	name, err := s.parseBackendBucketName(reqName)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := proto.Clone(req.GetBackendBucketResource()).(*pb.BackendBucket)
	obj.Id = proto.Uint64(s.generateID())
	obj.SelfLink = PtrTo(buildComputeSelfLink(ctx, fqn))
	obj.Kind = PtrTo("compute#backendBucket")
	obj.CreationTimestamp = PtrTo(s.nowString())

	if obj.Description == nil {
		obj.Description = PtrTo("")
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		OperationType: PtrTo("insert"),
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		User:          PtrTo("user@example.com"),
	}
	return s.computeOperations.startGlobalLRO(ctx, name.Project.ID, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *backendBuckets) Delete(ctx context.Context, req *pb.DeleteBackendBucketRequest) (*pb.Operation, error) {
	reqName := fmt.Sprintf("projects/%s/global/backendBuckets/%s", req.GetProject(), req.GetBackendBucket())
	name, err := s.parseBackendBucketName(reqName)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	deleted := &pb.BackendBucket{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		OperationType: PtrTo("delete"),
		TargetId:      deleted.Id,
		TargetLink:    deleted.SelfLink,
		User:          PtrTo("user@example.com"),
	}
	return s.computeOperations.startGlobalLRO(ctx, name.Project.ID, op, func() (proto.Message, error) {
		return &emptypb.Empty{}, nil
	})
}

func (s *backendBuckets) Update(ctx context.Context, req *pb.UpdateBackendBucketRequest) (*pb.Operation, error) {
	reqName := fmt.Sprintf("projects/%s/global/backendBuckets/%s", req.GetProject(), req.GetBackendBucket())
	name, err := s.parseBackendBucketName(reqName)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.BackendBucket{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// TODO: Apply field mask.
	proto.Merge(obj, req.GetBackendBucketResource())

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		OperationType: PtrTo("update"),
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		User:          PtrTo("user@example.com"),
	}
	return s.computeOperations.startGlobalLRO(ctx, name.Project.ID, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *backendBuckets) Patch(ctx context.Context, req *pb.PatchBackendBucketRequest) (*pb.Operation, error) {
	reqName := fmt.Sprintf("projects/%s/global/backendBuckets/%s", req.GetProject(), req.GetBackendBucket())
	name, err := s.parseBackendBucketName(reqName)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.BackendBucket{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// TODO: Apply field mask.
	proto.Merge(obj, req.GetBackendBucketResource())

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		OperationType: PtrTo("patch"),
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		User:          PtrTo("user@example.com"),
	}
	return s.computeOperations.startGlobalLRO(ctx, name.Project.ID, op, func() (proto.Message, error) {
		return obj, nil
	})
}

type backendBucketName struct {
	Project *projects.ProjectData
	Name    string
}

func (n *backendBucketName) String() string {
	return "projects/" + n.Project.ID + "/global/backendBuckets/" + n.Name
}

// parseBackendBucketName parses a string into a backendBucketName.
// The expected form is `locations/global/firewallPolicies/*`.
func (s *MockService) parseBackendBucketName(name string) (*backendBucketName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 5 && tokens[0] == "projects" && tokens[2] == "global" && tokens[3] == "backendBuckets" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &backendBucketName{
			Project: project,
			Name:    tokens[4],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
