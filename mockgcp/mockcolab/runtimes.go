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
// proto.service: google.cloud.aiplatform.v1.NotebookService
// proto.message: google.cloud.aiplatform.v1.NotebookRuntime

package mockaiplatform

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/aiplatform/v1"
)

type notebookServiceV1 struct {
	*MockService
	pb.UnimplementedNotebookServiceServer
}

func (s *notebookServiceV1) GetNotebookRuntime(ctx context.Context, req *pb.GetNotebookRuntimeRequest) (*pb.NotebookRuntime, error) {
	name, err := s.parseNotebookRuntimeName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.NotebookRuntime{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "NotebookRuntime %q not found.", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *notebookServiceV1) DeleteNotebookRuntime(ctx context.Context, req *pb.DeleteNotebookRuntimeRequest) (*longrunning.Operation, error) {
	name, err := s.parseNotebookRuntimeName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.NotebookRuntime{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	op := &pb.DeleteOperationMetadata{}
	opName := name.String()

	return s.operations.DoneLRO(ctx, opName, op, &emptypb.Empty{})
}

type notebookRuntimeName struct {
	Project          *projects.ProjectData
	Location         string
	NotebookRuntime  string
}

func (n *notebookRuntimeName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/notebookRuntimes/%s", n.Project.ID, n.Location, n.NotebookRuntime)
}

// parseNotebookRuntimeName parses a string into a notebookRuntimeName.
// The expected form is `projects/*/locations/*/notebookRuntimes/*`.
func (s *MockService) parseNotebookRuntimeName(name string) (*notebookRuntimeName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "notebookRuntimes" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &notebookRuntimeName{
			Project:          project,
			Location:         tokens[3],
			NotebookRuntime:  tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}


