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
// proto.service:
// proto.message: google.cloud.netapp.v1.BackupVault

package mocknetapp

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/netapp/v1"
)

type BackupVaultsService struct {
	pb.UnimplementedNetAppServer
	projectStore *projects.ProjectStore
}

func (s *BackupVaultsService) GetBackupVault(ctx context.Context, req *pb.GetBackupVaultRequest) (*pb.BackupVault, error) {
	name, err := s.parseBackupVaultName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.BackupVault{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *BackupVaultsService) CreateBackupVault(ctx context.Context, req *pb.CreateBackupVaultRequest) (*pb.BackupVault, error) {
	reqName := req.Parent + "/backupVaults/" + req.BackupVaultId
	name, err := s.parseBackupVaultName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	obj := req.BackupVault
	obj.Name = fqn

	obj.CreateTime = timestamppb.New(now)
	obj.State = pb.BackupVault_READY

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

type backupVaultName struct {
	Project        *projects.ProjectData
	Location       string
	BackupVaultId string
}

func (n *backupVaultName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/backupVaults/%s", n.Project.ID, n.Location, n.BackupVaultId)
}

// parseBackupVaultName parses a string into a backupVaultName.
// The expected form is `projects/*/locations/*/backupVaults/*`.
func (s *BackupVaultsService) parseBackupVaultName(name string) (*backupVaultName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "backupVaults" {
		project, err := s.projectStore.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &backupVaultName{
			Project:        project,
			Location:       tokens[3],
			BackupVaultId: tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

```
</out>


