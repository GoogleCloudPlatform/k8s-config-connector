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
// proto.service: google.cloud.backupdr.v1.BackupDR
// proto.message: google.cloud.backupdr.v1.ManagementServer

package mockbackupdr

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/fields"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/backupdr/v1"
	"github.com/google/uuid"

	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// GetManagementServer implements the BackupDRServer interface.
func (s *BackupDRV1) GetManagementServer(ctx context.Context, req *pb.GetManagementServerRequest) (*pb.ManagementServer, error) {
	name, err := s.parseManagementServerName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.ManagementServer{}
	if err := s.MockService.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", req.Name)
		}
		return nil, err
	}

	return obj, nil
}

// CreateManagementServer implements the BackupDRServer interface.
func (s *BackupDRV1) CreateManagementServer(ctx context.Context, req *pb.CreateManagementServerRequest) (*longrunningpb.Operation, error) {
	reqName := req.Parent + "/managementServers/" + req.ManagementServerId
	name, err := s.parseManagementServerName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	obj := proto.Clone(req.ManagementServer).(*pb.ManagementServer)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.State = pb.ManagementServer_CREATING
	handleGeneratedFields(name, obj)
	obj.Etag = fields.ComputeWeakEtag(obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	opMetadata := &pb.OperationMetadata{
		ApiVersion:            "v1",
		CreateTime:            timestamppb.New(now),
		RequestedCancellation: false,
		Target:                fmt.Sprintf("projects/%s/locations/%s/managementServers/%s", name.Project.ID, name.Location, req.ManagementServerId),
		Verb:                  "create",
	}
	opPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.MockService.operations.StartLRO(ctx, opPrefix, opMetadata, func() (proto.Message, error) {
		obj.State = pb.ManagementServer_READY
		obj.UpdateTime = timestamppb.New(time.Now())

		opMetadata.EndTime = timestamppb.New(time.Now())

		if err := s.storage.Update(ctx, fqn, obj); err != nil {
			return nil, err
		}

		return obj, nil
	})
}

// DeleteManagementServer implements the BackupDRServer interface.
func (s *BackupDRV1) DeleteManagementServer(ctx context.Context, req *pb.DeleteManagementServerRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseManagementServerName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	obj := &pb.ManagementServer{}
	if err := s.storage.Delete(ctx, fqn, obj); err != nil {
		return nil, err
	}

	obj.State = pb.ManagementServer_DELETING
	obj.UpdateTime = timestamppb.New(now)

	opMetadata := &pb.OperationMetadata{
		ApiVersion: "v1",
		CreateTime: timestamppb.New(now),
		Target:     fqn,
		Verb:       "delete",
	}
	opPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.MockService.operations.StartLRO(ctx, opPrefix, opMetadata, func() (proto.Message, error) {
		opMetadata.EndTime = timestamppb.New(time.Now())
		return &emptypb.Empty{}, nil
	})
}

// ManagementServerName represents a parsed management server resource name.
type ManagementServerName struct {
	Project            *projects.ProjectData
	Location           string
	ManagementServerID string
}

func (n *ManagementServerName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/managementServers/%s",
		n.Project.ID, n.Location, n.ManagementServerID)
}

// parseManagementServerName parses a string into a ManagementServerName.
// The expected form is: projects/<project>/locations/<location>/managementServers/<id>
func (s *BackupDRV1) parseManagementServerName(name string) (*ManagementServerName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "managementServers" {
		projectName, err := projects.ParseProjectName(tokens[0] + "/" + tokens[1])
		if err != nil {
			return nil, err
		}
		project, err := s.Projects.GetProject(projectName)
		if err != nil {
			return nil, err
		}

		name := &ManagementServerName{
			Project:            project,
			Location:           tokens[3],
			ManagementServerID: tokens[5],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}

func handleGeneratedFields(name *ManagementServerName, obj *pb.ManagementServer) {
	obj.ManagementUri = &pb.ManagementURI{
		Api:   fmt.Sprintf("https://bmc-%d-3aclcdbj-dot-%s.backupdr.googleusercontent.com/actifio", name.Project.Number, name.Location),
		WebUi: fmt.Sprintf("https://bmc-%d-3aclcdbj-dot-%s.backupdr.googleusercontent.com", name.Project.Number, name.Location),
	}
	obj.Oauth2ClientId = uuid.New().String()
}
