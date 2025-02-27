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
// proto.service: google.cloud.notebooks.v1.NotebookService
// proto.message: google.cloud.notebooks.v1.Instance

package mocknotebooks

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/notebooks/v1"
)

type NotebooksV1 struct {
	*MockService
	pb.UnimplementedNotebookServiceServer
}

func (s *NotebooksV1) GetInstance(ctx context.Context, req *pb.GetInstanceRequest) (*pb.Instance, error) {
	name, err := s.parseInstanceName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Instance{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *NotebooksV1) CreateInstance(ctx context.Context, req *pb.CreateInstanceRequest) (*longrunning.Operation, error) {
	reqName := req.Parent + "/instances/" + req.InstanceId
	name, err := s.parseInstanceName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.Instance).(*pb.Instance)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(time.Now())
	obj.State = pb.Instance_ACTIVE
	s.setDefaultServiceAccount(obj, name)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}
	return s.operations.NewLRO(ctx)
}

func (s *NotebooksV1) setDefaultServiceAccount(obj *pb.Instance, name *instanceName) {
	if obj.ServiceAccount == "" {
		obj.ServiceAccount = fmt.Sprintf("%d-compute@developer.gserviceaccount.com", name.Project.Number)
	}
}

func (s *NotebooksV1) DeleteInstance(ctx context.Context, req *pb.DeleteInstanceRequest) (*longrunning.Operation, error) {
	name, err := s.parseInstanceName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.Instance{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	return s.operations.NewLRO(ctx)
}

type instanceName struct {
	Project *projects.ProjectData
	region  string
	name    string
}

func (n *instanceName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/instances/%s", n.Project.ID, n.region, n.name)
}

// parseInstanceName parses a string into an instanceName.
// The expected form is `projects/*/locations/*/instances/*`.
func (s *MockService) parseInstanceName(name string) (*instanceName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "instances" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &instanceName{
			Project: project,
			region:  tokens[3],
			name:    tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}


