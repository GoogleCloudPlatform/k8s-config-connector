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

package mockcloudbuild

import (
	"context"
	"strconv"
	"strings"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/fields"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/devtools/cloudbuild/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var _ pb.CloudBuildServer = &CloudBuildV1{}

type CloudBuildV1 struct {
	*MockService
	pb.UnimplementedCloudBuildServer
}

func (s *CloudBuildV1) GetWorkerPool(ctx context.Context, req *pb.GetWorkerPoolRequest) (*pb.WorkerPool, error) {
	name, err := s.parseWorkerPoolName(req.Name)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.WorkerPool{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Requested entity was not found.")
		}
		return nil, err
	}
	return obj, nil
}

func (s *CloudBuildV1) CreateWorkerPool(ctx context.Context, req *pb.CreateWorkerPoolRequest) (*longrunningpb.Operation, error) {
	workerPoolName := req.GetParent() + "/workerPools/" + req.GetWorkerPoolId()
	name, err := s.parseWorkerPoolName(workerPoolName)
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	now := timestamppb.Now()

	obj := proto.Clone(req.GetWorkerPool()).(*pb.WorkerPool)
	obj.Name = fqn
	obj.CreateTime = now

	populateDefaultsForWorkerPool(obj)
	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}
	metadata := &pb.CreateWorkerPoolOperationMetadata{
		WorkerPool: fqn,
		CreateTime: now,
	}
	return s.operations.StartLRO(ctx, name.String(), metadata, func() (proto.Message, error) {
		return obj, nil
	})
}

func populateDefaultsForWorkerPool(wp *pb.WorkerPool) {
	now := timestamppb.Now()
	network := wp.GetPrivatePoolV1Config().GetNetworkConfig()
	if network != nil {
		tokens := strings.Split(network.PeeredNetwork, "/")
		if len(tokens) == 5 {
			network.PeeredNetwork = tokens[0] + "/" + "${projectNumber}" + "/" + tokens[2] + "/" + tokens[3] + "/" + tokens[4]
		}
	}
	wp.UpdateTime = now
	wp.State = pb.WorkerPool_RUNNING
	wp.Etag = fields.ComputeWeakEtag(wp)
	wp.Uid = "11111111111111111111"
}

func (s *CloudBuildV1) UpdateWorkerPool(ctx context.Context, req *pb.UpdateWorkerPoolRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseWorkerPoolName(req.WorkerPool.Name)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.WorkerPool{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Requested entity was not found.")
		}
		return nil, err
	}
	if err := fields.UpdateByFieldMask(obj, req.WorkerPool, req.UpdateMask.Paths); err != nil {
		return nil, err
	}
	populateDefaultsForWorkerPool(obj)
	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}
	metadata := &pb.UpdateWorkerPoolOperationMetadata{
		WorkerPool: name.String(),
		CreateTime: timestamppb.Now(),
	}
	return s.operations.StartLRO(ctx, name.String(), metadata, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *CloudBuildV1) DeleteWorkerPool(ctx context.Context, req *pb.DeleteWorkerPoolRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseWorkerPoolName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := timestamppb.Now()

	existing := &pb.WorkerPool{}
	err = s.storage.Delete(ctx, fqn, existing)
	if err != nil {
		if status.Code(err) == codes.NotFound && req.AllowMissing {
			return s.operations.NewLRO(ctx)
		}
		return &longrunningpb.Operation{}, err
	}
	metadata := &pb.DeleteWorkerPoolOperationMetadata{
		WorkerPool:   fqn,
		CreateTime:   now,
		CompleteTime: now,
	}
	return s.operations.DoneLRO(ctx, name.String(), metadata, &pb.WorkerPool{})

}

type workerPoolName struct {
	Project *projects.ProjectData
	// TODO: location validation
	Location       string
	WorkerPoolName string
}

func (n *workerPoolName) String() string {
	return n.GetParent() + "/workerPools/" + n.WorkerPoolName
}

func (n *workerPoolName) GetParent() string {
	return "projects/" + strconv.FormatInt(n.Project.Number, 10) + "/locations/" + n.Location
}

func (s *MockService) parseWorkerPoolName(name string) (*workerPoolName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "workerPools" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}
		name := &workerPoolName{
			Project:        project,
			Location:       tokens[3],
			WorkerPoolName: tokens[5],
		}
		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
