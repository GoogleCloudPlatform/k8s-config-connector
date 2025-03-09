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
			return nil, status.Errorf(codes.NotFound, "backup %q not found", name)
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
	// TODO: EndTime calculation.
	obj.EndTime = timestamppb.New(now.Add(time.Minute * 2))
	obj.State = pb.Backup_ACTIVE

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix := name.Parent().String() + "/operations/" + name.BackupName

	return s.operations.StartLRO(ctx, lroPrefix, nil, func() (proto.Message, error) {
		return obj, nil
	})
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

	prefix := name.Parent().String() + "/operations/" + name.BackupName
	return s.operations.StartLRO(ctx, prefix, nil, func() (proto.Message, error) {
		return &emptypb.Empty{}, nil
	})
}

type serviceName struct {
	Project   *projects.ProjectData
	Location  string
	ServiceID string
}

func (n *serviceName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/services/%s", n.Project.ID, n.Location, n.ServiceID)
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
		Project:   n.Project,
		Location:  n.Location,
		ServiceID: n.Service,
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
