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
	"reflect"
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
	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}
	metadata := &pb.CreateWorkerPoolOperationMetadata{
		WorkerPool:   fqn,
		CreateTime:   now,
		CompleteTime: now,
	}
	return s.operations.StartLRO(ctx, name.String(), metadata, func() (proto.Message, error) {
		// Many fields are not populated in the LRO result
		result := proto.Clone(obj).(*pb.WorkerPool)
		result.CreateTime = now
		result.UpdateTime = now
		result.State = pb.WorkerPool_RUNNING
		result.Etag = fields.ComputeWeakEtag(result)
		if err := s.storage.Update(ctx, fqn, result); err != nil { // update resource object when LRO is done
			return nil, err
		}
		return result, nil
	})
}

func (s *CloudBuildV1) UpdateWorkerPool(ctx context.Context, req *pb.UpdateWorkerPoolRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseWorkerPoolName(req.WorkerPool.Name)
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	obj := &pb.WorkerPool{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}
	now := timestamppb.Now()
	obj.UpdateTime = now
	source := reflect.ValueOf(req.WorkerPool)
	target := reflect.ValueOf(obj).Elem()
	for _, path := range req.UpdateMask.Paths {
		f := target.FieldByName(path)
		if f.IsValid() && f.CanSet() {
			switch f.Kind() {
			case reflect.Int, reflect.Int64:
				intVal := source.FieldByName(path).Int()
				f.SetInt(intVal)
			case reflect.String:
				stringVal := source.FieldByName(path).String()
				f.SetString(stringVal)
			}

		}
	}
	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}
	metadata := &pb.UpdateWorkerPoolOperationMetadata{
		WorkerPool:   name.String(),
		CreateTime:   now,
		CompleteTime: now,
	}
	return s.operations.StartLRO(ctx, name.String(), metadata, func() (proto.Message, error) {
		// Many fields are not populated in the LRO result
		result := proto.Clone(obj).(*pb.WorkerPool)
		result.UpdateTime = now
		result.State = pb.WorkerPool_RUNNING
		result.Etag = fields.ComputeWeakEtag(result)
		if err := s.storage.Update(ctx, fqn, result); err != nil { // update resource object when LRO is done
			return nil, err
		}
		return result, nil
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
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/workerPools/" + n.WorkerPoolName
}

func (n *workerPoolName) GetParent() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location
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
