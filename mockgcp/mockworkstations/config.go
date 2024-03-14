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

package mockworkstations

import (
	"context"
	"fmt"
	"strings"
	"time"

	longrunning "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/workstations/v1beta"
	"github.com/google/uuid"
)

func (s *workstationsServer) GetWorkstationConfig(ctx context.Context, req *pb.GetWorkstationConfigRequest) (*pb.WorkstationConfig, error) {
	name, err := s.parseWorkstationConfigName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.WorkstationConfig{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *workstationsServer) CreateWorkstationConfig(ctx context.Context, req *pb.CreateWorkstationConfigRequest) (*longrunning.Operation, error) {
	reqName := req.GetParent() + "/workstationConfigs/" + req.GetWorkstationConfigId()
	name, err := s.parseWorkstationConfigName(reqName)
	if err != nil {
		return nil, err
	}

	if req.GetValidateOnly() {
		return nil, status.Errorf(codes.Unimplemented, "validate only is not yet implemented in mockgcp")
	}

	fqn := name.String()

	now := time.Now()

	obj := proto.Clone(req.WorkstationConfig).(*pb.WorkstationConfig)
	obj.Name = fqn

	obj.Uid = uuid.New().String()

	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)

	if obj.Container == nil {
		obj.Container = &pb.WorkstationConfig_Container{}
	}
	if obj.Container.Image == "" {
		obj.Container.Image = fmt.Sprintf("%s-docker.pkg.dev/cloud-workstations-images/predefined/code-oss:latest", name.Location)
	}

	if obj.IdleTimeout == nil {
		obj.IdleTimeout = &durationpb.Duration{
			Seconds: 1200,
		}
	}
	if obj.RunningTimeout == nil {
		obj.IdleTimeout = &durationpb.Duration{
			Seconds: 43200,
		}
	}
	if obj.ReplicaZones == nil {
		obj.ReplicaZones = []string{
			"us-central1-a",
			"us-central1-c",
		}
	}

	if obj.Host == nil {
		obj.Host = &pb.WorkstationConfig_Host{}
	}
	if obj.Host.Config == nil {
		obj.Host.Config = &pb.WorkstationConfig_Host_GceInstance_{}
	}
	if gceHostConfig := obj.GetHost().GetGceInstance(); gceHostConfig != nil {
		if gceHostConfig.BootDiskSizeGb == 0 {
			gceHostConfig.BootDiskSizeGb = 50
		}
		if gceHostConfig.ConfidentialInstanceConfig == nil {
			gceHostConfig.ConfidentialInstanceConfig = &pb.WorkstationConfig_Host_GceInstance_GceConfidentialInstanceConfig{}
		}
		if gceHostConfig.MachineType == "" {
			gceHostConfig.MachineType = "e2-standard-4"
		}

		if gceHostConfig.ServiceAccount == "" {
			gceHostConfig.ServiceAccount = fmt.Sprintf("service-%d@gcp-sa-workstationsvm.iam.gserviceaccount.com", name.Project.Number)
		}

		if gceHostConfig.ShieldedInstanceConfig == nil {
			gceHostConfig.ShieldedInstanceConfig = &pb.WorkstationConfig_Host_GceInstance_GceShieldedInstanceConfig{}
		}
	}

	obj.Etag = computeEtag(obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.OperationMetadata{
		CreateTime:            timestamppb.New(now),
		Target:                fqn,
		Verb:                  "create",
		RequestedCancellation: false,
		ApiVersion:            "v1beta",
	}
	opPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, opPrefix, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *workstationsServer) UpdateWorkstationConfig(ctx context.Context, req *pb.UpdateWorkstationConfigRequest) (*longrunning.Operation, error) {
	name, err := s.parseWorkstationConfigName(req.GetWorkstationConfig().GetName())
	if err != nil {
		return nil, err
	}

	if req.GetValidateOnly() {
		return nil, status.Errorf(codes.Unimplemented, "validateOnly is not yet implemented in mockgcp")
	}
	if req.GetAllowMissing() {
		return nil, status.Errorf(codes.Unimplemented, "allowMissing is not yet implemented in mockgcp")
	}

	fqn := name.String()
	now := time.Now()

	obj := &pb.WorkstationConfig{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// See docs for UpdateMask
	updateMask := req.GetUpdateMask()
	for _, path := range updateMask.Paths {
		switch path {
		case "displayName":
			obj.DisplayName = req.GetWorkstationConfig().GetDisplayName()

		default:
			return nil, status.Errorf(codes.InvalidArgument, "field %q is not yet handled in mockgcp", path)
		}
	}

	obj.UpdateTime = timestamppb.New(now)

	obj.Etag = computeEtag(obj)

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.OperationMetadata{
		CreateTime:            timestamppb.New(now),
		Target:                fqn,
		Verb:                  "update",
		RequestedCancellation: false,
		ApiVersion:            "v1beta",
	}
	opPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, opPrefix, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *workstationsServer) DeleteWorkstationConfig(ctx context.Context, req *pb.DeleteWorkstationConfigRequest) (*longrunning.Operation, error) {
	name, err := s.parseWorkstationConfigName(req.GetName())
	if err != nil {
		return nil, err
	}

	if req.GetValidateOnly() {
		return nil, status.Errorf(codes.Unimplemented, "validateOnly is not yet implemented in mockgcp")
	}
	if req.GetForce() {
		return nil, status.Errorf(codes.Unimplemented, "force is not yet implemented in mockgcp")
	}

	fqn := name.String()
	now := time.Now()

	deleted := &pb.WorkstationConfig{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	op := &pb.OperationMetadata{
		CreateTime:            timestamppb.New(now),
		Target:                fqn,
		Verb:                  "delete",
		RequestedCancellation: false,
		ApiVersion:            "v1beta",
	}
	opPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, opPrefix, op, func() (proto.Message, error) {
		return deleted, nil
	})
}

type WorkstationConfigName struct {
	*WorkstationClusterName
	WorkstationConfigID string
}

func (n *WorkstationConfigName) String() string {
	return n.WorkstationClusterName.String() + "/workstationConfigs/" + n.WorkstationConfigID
}

// parseWorkstationConfigsName parses a string into a WorkstationConfigsName.
// The expected form is projects/*/locations/*/workstationClusters/*/workstationConfigs/*
func (s *MockService) parseWorkstationConfigName(name string) (*WorkstationConfigName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 8 && tokens[6] == "workstationConfigs" {
		clusterName, err := s.parseWorkstationClusterName(strings.Join(tokens[0:6], "/"))
		if err != nil {
			return nil, err
		}

		name := &WorkstationConfigName{
			WorkstationClusterName: clusterName,
			WorkstationConfigID:    tokens[7],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
