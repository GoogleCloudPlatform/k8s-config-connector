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

// +tool:mockgcp-support
// proto.service: google.cloud.metastore.v1.DataprocMetastore
// proto.message: google.cloud.metastore.v1.Backup

package mockmetastore

import (
	"context"
	"fmt"
	"strings"
	"time"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/metastore/v1"
)

func (s *DataprocMetastoreV1) GetBackup(ctx context.Context, req *pb.GetBackupRequest) (*pb.Backup, error) {
	name, err := s.parseBackupName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Backup{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *DataprocMetastoreV1) CreateBackup(ctx context.Context, req *pb.CreateBackupRequest) (*longrunningpb.Operation, error) {
	reqName := req.Parent + "/backups/" + req.BackupId
	name, err := s.parseBackupName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	obj := proto.Clone(req.GetBackup()).(*pb.Backup)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.EndTime = obj.CreateTime // Set EndTime to CreateTime based on diff
	obj.State = pb.Backup_ACTIVE
	// Revert to placeholder ServiceRevision
	service := &pb.Service{}
	if err := s.storage.Get(ctx, req.Parent, service); err != nil {
		return nil, fmt.Errorf("failed to get service: %w", err)
	}
	serviceCopy := proto.Clone(service).(*pb.Service)
	serviceCopy.State = pb.Service_STATE_UNSPECIFIED
	serviceCopy.StateMessage = ""
	serviceCopy.Name = ""
	obj.ServiceRevision = serviceCopy

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// The LRO prefix should be the parent resource path for the operations collection.
	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)

	metadata := &pb.OperationMetadata{
		CreateTime:            timestamppb.New(now),
		Target:                fqn,
		Verb:                  "create",
		RequestedCancellation: false, // Field not expected in initial LRO metadata
		ApiVersion:            "v1",
	}

	op, err := s.operations.StartLRO(ctx, lroPrefix, metadata, func() (proto.Message, error) {
		// We need to fetch the object again to get the updated LRO state
		metadata.EndTime = timestamppb.New(now)
		updatedObj := &pb.Backup{}
		if err := s.storage.Get(ctx, fqn, updatedObj); err != nil {
			// This should not happen in the normal flow
			return nil, fmt.Errorf("failed to get backup during LRO completion: %w", err)
		}
		// Match the expected log: Ensure EndTime equals CreateTime in the final response.
		updatedObj.EndTime = updatedObj.CreateTime
		return updatedObj, nil
	})
	if err != nil {
		return nil, err
	}
	// Initial response should not have `Done` set (it defaults to false)
	// op.Done = false // Reverted based on diff

	return op, nil
}

func (s *DataprocMetastoreV1) DeleteBackup(ctx context.Context, req *pb.DeleteBackupRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseBackupName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.Backup{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	// The LRO prefix should be the parent resource path for the operations collection.
	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	now := time.Now()
	metadata := &pb.OperationMetadata{
		CreateTime:            timestamppb.New(now),
		Target:                fqn,
		Verb:                  "delete",
		RequestedCancellation: false,
		ApiVersion:            "v1",
	}
	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		// Ensure the correct type URL is used for Empty responses
		metadata.EndTime = timestamppb.New(now)
		return &emptypb.Empty{}, nil
	})
}

type backupName struct {
	Project    *projects.ProjectData
	Location   string
	Service    string
	BackupName string
}

func (n *backupName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/services/%s/backups/%s", n.Project.ID, n.Location, n.Service, n.BackupName)
}

func (n *backupName) Parent() *serviceName {
	return &serviceName{
		Project:  n.Project,
		Location: n.Location,
		Name:     n.Service,
	}
}

// parseBackupName parses a string into a backupName.
// The expected form is `projects/*/locations/*/services/*/backups/*`.
func (s *MockService) parseBackupName(name string) (*backupName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "services" && tokens[6] == "backups" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &backupName{
			Project:    project,
			Location:   tokens[3],
			Service:    tokens[5],
			BackupName: tokens[7],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
