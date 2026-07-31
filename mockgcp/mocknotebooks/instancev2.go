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

package mocknotebooks

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/notebooks/apiv2/notebookspb"
)

func (s *NotebookServiceV2) GetInstance(ctx context.Context, req *pb.GetInstanceRequest) (*pb.Instance, error) {
	name, err := s.parseInstanceName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Instance{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}
	obj.Name = fqn
	obj.Creator = "someone@somewhere.com"
	obj.Id = "some-random-id"
	obj.HealthState = pb.HealthState_HEALTHY
	obj.HealthInfo = map[string]string{
		"docker_proxy_agent_status": "1",
		"docker_status":             "1",
		"jupyterlab_api_status":     "1",
		"jupyterlab_status":         "1",
	}
	return obj, nil
}

func (s *NotebookServiceV2) CreateInstance(ctx context.Context, req *pb.CreateInstanceRequest) (*longrunning.Operation, error) {
	name, err := s.parseInstanceName(fmt.Sprintf("%s/instances/%s", req.Parent, req.InstanceId))
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.Instance).(*pb.Instance)
	obj.Name = fqn
	obj.Creator = "someone@somewhere.com"
	obj.Id = "some-random-id"
	obj.HealthState = pb.HealthState_HEALTHY
	obj.HealthInfo = map[string]string{
		"docker_proxy_agent_status": "1",
		"docker_status":             "1",
		"jupyterlab_api_status":     "1",
		"jupyterlab_status":         "1",
	}
	obj.CreateTime = timestamppb.New(time.Now())
	obj.UpdateTime = timestamppb.New(time.Now())
	obj.State = pb.State_ACTIVE

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.region)
	metadata := &pb.OperationMetadata{
		CreateTime:            timestamppb.New(time.Now()),
		RequestedCancellation: false,
		Target:                name.String(),
		Verb:                  "create",
		Endpoint:              "CreateInstance",
	}
	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.New(time.Now())
		return obj, nil
	})
}

func (s *NotebookServiceV2) UpdateInstance(ctx context.Context, req *pb.UpdateInstanceRequest) (*longrunning.Operation, error) {
	name, err := s.parseInstanceName(req.Instance.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	existing := &pb.Instance{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		return nil, err
	}

	updated := proto.Clone(req.Instance).(*pb.Instance)
	updated.Name = fqn
	updated.Creator = "someone@somewhere.com"
	updated.Id = "some-random-id"
	updated.HealthState = pb.HealthState_HEALTHY
	updated.HealthInfo = map[string]string{
		"docker_proxy_agent_status": "1",
		"docker_status":             "1",
		"jupyterlab_api_status":     "1",
		"jupyterlab_status":         "1",
	}
	updated.UpdateTime = timestamppb.New(time.Now())
	updated.CreateTime = existing.CreateTime
	updated.State = pb.State_ACTIVE

	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.region)
	metadata := &pb.OperationMetadata{
		CreateTime:            timestamppb.New(time.Now()),
		RequestedCancellation: false,
		Target:                name.String(),
		Verb:                  "update",
		Endpoint:              "UpdateInstance",
	}
	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.New(time.Now())
		return updated, nil
	})
}

func (s *NotebookServiceV2) DeleteInstance(ctx context.Context, req *pb.DeleteInstanceRequest) (*longrunning.Operation, error) {
	name, err := s.parseInstanceName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.Instance{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}
	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.region)
	metadata := &pb.OperationMetadata{
		CreateTime:            timestamppb.New(time.Now()),
		RequestedCancellation: false,
		Target:                name.String(),
		Verb:                  "delete",
		Endpoint:              "DeleteInstance",
	}
	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.New(time.Now())
		return &emptypb.Empty{}, nil
	})
}
