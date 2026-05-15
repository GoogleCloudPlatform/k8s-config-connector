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

package mockgkebackup

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/gkebackup/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

func (s *BackupForGKEV1) GetBackupPlanBinding(ctx context.Context, req *pb.GetBackupPlanBindingRequest) (*pb.BackupPlanBinding, error) {
	name, err := s.parseBackupPlanBindingName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.BackupPlanBinding{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "BackupPlanBinding %q not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *BackupForGKEV1) ListBackupPlanBindings(ctx context.Context, req *pb.ListBackupPlanBindingsRequest) (*pb.ListBackupPlanBindingsResponse, error) {
	res := &pb.ListBackupPlanBindingsResponse{}
	kind := (&pb.BackupPlanBinding{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, kind, storage.ListOptions{
		Prefix: req.Parent,
	}, func(obj proto.Message) error {
		res.BackupPlanBindings = append(res.BackupPlanBindings, obj.(*pb.BackupPlanBinding))
		return nil
	}); err != nil {
		return nil, err
	}

	return res, nil
}

type backupPlanBindingName struct {
	Project             *projects.ProjectData
	Location            string
	BackupPlanBindingID string
}

func (n *backupPlanBindingName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/backupPlanBindings/%s", n.Project.ID, n.Location, n.BackupPlanBindingID)
}

func (s *MockService) parseBackupPlanBindingName(name string) (*backupPlanBindingName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "backupPlanBindings" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &backupPlanBindingName{
			Project:             project,
			Location:            tokens[3],
			BackupPlanBindingID: tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
