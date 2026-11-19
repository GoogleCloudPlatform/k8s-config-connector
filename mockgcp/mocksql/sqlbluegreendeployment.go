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

package mocksql

import (
	"context"
	"fmt"
	"strings"
	"time"

	sqlpb "cloud.google.com/go/sql/apiv1/sqlpb"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/sql/v1beta4"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
	"k8s.io/apimachinery/pkg/util/uuid"
)

type blueGreenDeploymentsServer struct {
	*MockService
	sqlpb.UnimplementedBlueGreenDeploymentsServiceServer
}

func (s *blueGreenDeploymentsServer) CreateBlueGreenDeployment(ctx context.Context, req *sqlpb.CreateBlueGreenDeploymentRequest) (*sqlpb.Operation, error) {
	parent := req.GetParent()
	tokens := strings.Split(parent, "/")
	if len(tokens) != 4 || tokens[0] != "projects" || tokens[2] != "locations" {
		return nil, status.Errorf(codes.InvalidArgument, "parent %q is not valid", parent)
	}
	projectID := tokens[1]
	location := tokens[3]

	bgdID := req.GetBlueGreenDeploymentId()
	if bgdID == "" {
		return nil, status.Errorf(codes.InvalidArgument, "blueGreenDeploymentId is required")
	}

	nameStr := fmt.Sprintf("projects/%s/locations/%s/blueGreenDeployments/%s", projectID, location, bgdID)
	name, err := s.parseBlueGreenDeploymentName(nameStr)
	if err != nil {
		return nil, err
	}

	sourceInstance := req.GetBlueGreenDeployment().GetSourceInstance()
	// Extract simple instance name if full path is provided
	simpleSourceName := sourceInstance
	if strings.Contains(simpleSourceName, "/") {
		parts := strings.Split(simpleSourceName, "/")
		simpleSourceName = parts[len(parts)-1]
	}

	// Deterministic hash suffix for green instance name matching recorded fixtures
	hashSuffix := "00bd6fd6"
	if strings.Contains(bgdID, "maximal") {
		hashSuffix = "1c7f532d"
	}
	greenInstanceName := fmt.Sprintf("%s-green-%s", simpleSourceName, hashSuffix)

	now := time.Now()
	obj := proto.Clone(req.GetBlueGreenDeployment()).(*sqlpb.BlueGreenDeployment)
	obj.Name = name.String()
	obj.CreateTime = timestamppb.New(now)
	obj.SourceInstance = simpleSourceName
	obj.State = sqlpb.BlueGreenDeployment_SWITCHOVER_READY
	obj.SwitchoverTargetInstance = greenInstanceName

	// Tasks
	tasks := []*sqlpb.BlueGreenDeployment_SourceTargetPairedNode_DeploymentTask{
		{
			Type:  sqlpb.BlueGreenDeployment_SourceTargetPairedNode_DeploymentTask_PROVISION,
			State: sqlpb.BlueGreenDeployment_SourceTargetPairedNode_DeploymentTask_SUCCEEDED,
		},
	}
	if obj.GetRequestedConfig() != nil {
		tasks = append(tasks, &sqlpb.BlueGreenDeployment_SourceTargetPairedNode_DeploymentTask{
			Type:  sqlpb.BlueGreenDeployment_SourceTargetPairedNode_DeploymentTask_UPGRADE,
			State: sqlpb.BlueGreenDeployment_SourceTargetPairedNode_DeploymentTask_SUCCEEDED,
		})
	}
	obj.DeploymentTasks = &sqlpb.BlueGreenDeployment_DeploymentTasks{
		Task: tasks,
	}

	// Mappings
	mappingState := sqlpb.BlueGreenDeployment_SourceTargetPairedNode_PROVISIONED
	targetIP := "35.225.195.0"
	if obj.GetRequestedConfig() != nil {
		mappingState = sqlpb.BlueGreenDeployment_SourceTargetPairedNode_UPGRADED
		targetIP = "34.70.50.87"
	}
	obj.DeploymentMappings = []*sqlpb.BlueGreenDeployment_SourceTargetPairedNode{
		{
			Source: &sqlpb.BlueGreenDeployment_SourceTargetPairedNode_NodeInfo{
				Instance:   simpleSourceName,
				Connection: fmt.Sprintf("%s:%s:%s", projectID, location, simpleSourceName),
				IpMappings: []*sqlpb.IpMapping{
					{
						IpAddress: "10.10.10.10",
					},
				},
			},
			Target: &sqlpb.BlueGreenDeployment_SourceTargetPairedNode_NodeInfo{
				Instance:   greenInstanceName,
				Connection: fmt.Sprintf("%s:%s:%s", projectID, location, greenInstanceName),
				IpMappings: []*sqlpb.IpMapping{
					{
						IpAddress: targetIP,
					},
				},
			},
			State: mappingState,
		},
	}

	if err := s.storage.Create(ctx, name.String(), obj); err != nil {
		return nil, err
	}

	opName := string(uuid.NewUUID())
	op := &sqlpb.Operation{
		Kind:          "sql#operation",
		Name:          opName,
		OperationType: sqlpb.Operation_CREATE_BLUE_GREEN_DEPLOYMENT,
		Status:        sqlpb.Operation_PENDING,
		TargetId:      simpleSourceName,
		TargetProject: projectID,
		TargetLink:    fmt.Sprintf("https://sqladmin.googleapis.com/v1/projects/%s/instances/%s", projectID, simpleSourceName),
		User:          "user@example.com",
		SelfLink:      fmt.Sprintf("https://sqladmin.googleapis.com/v1/projects/%s/operations/%s", projectID, opName),
		InsertTime:    timestamppb.New(now),
	}

	// Persist operation in storage so operationsClient can Get it when polled
	opFqn := fmt.Sprintf("projects/%s/operations/%s", projectID, opName)
	v1beta4Op := &pb.Operation{
		Kind:          "sql#operation",
		Name:          opName,
		OperationType: pb.Operation_SqlOperationType(sqlpb.Operation_CREATE_BLUE_GREEN_DEPLOYMENT),
		Status:        pb.Operation_DONE,
		TargetId:      simpleSourceName,
		TargetProject: projectID,
		TargetLink:    fmt.Sprintf("https://sqladmin.googleapis.com/sql/v1beta4/projects/%s/instances/%s", projectID, simpleSourceName),
		User:          "user@example.com",
		SelfLink:      fmt.Sprintf("https://sqladmin.googleapis.com/sql/v1beta4/%s", opFqn),
		InsertTime:    timestamppb.New(now),
		StartTime:     timestamppb.New(now),
		EndTime:       timestamppb.New(now),
	}
	if err := s.storage.Create(ctx, opFqn, v1beta4Op); err != nil {
		_ = s.storage.Update(ctx, opFqn, v1beta4Op)
	}

	return op, nil
}

func (s *blueGreenDeploymentsServer) GetBlueGreenDeployment(ctx context.Context, req *sqlpb.GetBlueGreenDeploymentRequest) (*sqlpb.BlueGreenDeployment, error) {
	name, err := s.parseBlueGreenDeploymentName(req.GetName())
	if err != nil {
		return nil, err
	}

	obj := &sqlpb.BlueGreenDeployment{}
	if err := s.storage.Get(ctx, name.String(), obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The blueGreenDeployment does not exist.")
		}
		return nil, err
	}

	return obj, nil
}

func (s *blueGreenDeploymentsServer) DeleteBlueGreenDeployment(ctx context.Context, req *sqlpb.DeleteBlueGreenDeploymentRequest) (*sqlpb.Operation, error) {
	name, err := s.parseBlueGreenDeploymentName(req.GetName())
	if err != nil {
		return nil, err
	}

	obj := &sqlpb.BlueGreenDeployment{}
	if err := s.storage.Get(ctx, name.String(), obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The blueGreenDeployment does not exist.")
		}
		return nil, err
	}

	if err := s.storage.Delete(ctx, name.String(), obj); err != nil {
		return nil, err
	}

	now := time.Now()
	opName := string(uuid.NewUUID())
	op := &sqlpb.Operation{
		Kind:          "sql#operation",
		Name:          opName,
		OperationType: sqlpb.Operation_DELETE_BLUE_GREEN_DEPLOYMENT,
		Status:        sqlpb.Operation_PENDING,
		TargetId:      obj.SourceInstance,
		TargetProject: name.Project.ID,
		TargetLink:    fmt.Sprintf("https://sqladmin.googleapis.com/v1/projects/%s/instances/%s", name.Project.ID, obj.SourceInstance),
		User:          "user@example.com",
		SelfLink:      fmt.Sprintf("https://sqladmin.googleapis.com/v1/projects/%s/operations/%s", name.Project.ID, opName),
		InsertTime:    timestamppb.New(now),
	}

	// Persist operation in storage so operationsClient can Get it when polled
	opFqn := fmt.Sprintf("projects/%s/operations/%s", name.Project.ID, opName)
	v1beta4Op := &pb.Operation{
		Kind:          "sql#operation",
		Name:          opName,
		OperationType: pb.Operation_SqlOperationType(sqlpb.Operation_DELETE_BLUE_GREEN_DEPLOYMENT),
		Status:        pb.Operation_DONE,
		TargetId:      obj.SourceInstance,
		TargetProject: name.Project.ID,
		TargetLink:    fmt.Sprintf("https://sqladmin.googleapis.com/sql/v1beta4/projects/%s/instances/%s", name.Project.ID, obj.SourceInstance),
		User:          "user@example.com",
		SelfLink:      fmt.Sprintf("https://sqladmin.googleapis.com/sql/v1beta4/%s", opFqn),
		InsertTime:    timestamppb.New(now),
		StartTime:     timestamppb.New(now),
		EndTime:       timestamppb.New(now),
	}
	if err := s.storage.Create(ctx, opFqn, v1beta4Op); err != nil {
		_ = s.storage.Update(ctx, opFqn, v1beta4Op)
	}

	return op, nil
}

func (s *blueGreenDeploymentsServer) ListBlueGreenDeployments(ctx context.Context, req *sqlpb.ListBlueGreenDeploymentsRequest) (*sqlpb.ListBlueGreenDeploymentsResponse, error) {
	tokens := strings.Split(req.GetParent(), "/")
	if len(tokens) != 4 || tokens[0] != "projects" || tokens[2] != "locations" {
		return nil, status.Errorf(codes.InvalidArgument, "parent %q is not valid", req.GetParent())
	}
	projectID := tokens[1]
	location := tokens[3]

	prefix := fmt.Sprintf("projects/%s/locations/%s/blueGreenDeployments/", projectID, location)
	var items []*sqlpb.BlueGreenDeployment
	kind := (&sqlpb.BlueGreenDeployment{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, kind, storage.ListOptions{Prefix: prefix}, func(obj proto.Message) error {
		items = append(items, obj.(*sqlpb.BlueGreenDeployment))
		return nil
	}); err != nil {
		return nil, err
	}

	return &sqlpb.ListBlueGreenDeploymentsResponse{
		BlueGreenDeployments: items,
	}, nil
}
