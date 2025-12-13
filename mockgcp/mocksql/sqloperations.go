// Copyright 2023 Google LLC
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
	"sort"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/sql/v1beta4"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

type sqlOperationsService struct {
	*MockService
	pb.UnimplementedSqlOperationsServiceServer
}

func (s *sqlOperationsService) Get(ctx context.Context, req *pb.SqlOperationsGetRequest) (*pb.Operation, error) {
	name, err := s.buildOperationName(req.GetProject(), req.GetOperation())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Operation{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *sqlOperationsService) List(ctx context.Context, req *pb.SqlOperationsListRequest) (*pb.OperationsListResponse, error) {
	project, err := s.projects.GetProjectByID(req.GetProject())
	if err != nil {
		return nil, err
	}

	prefix := "projects/" + project.ID + "/operations/"

	response := &pb.OperationsListResponse{
		Kind: "sql#operationsList",
	}

	operationKind := (&pb.Operation{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, operationKind, storage.ListOptions{
		Prefix: prefix,
	}, func(obj proto.Message) error {
		response.Items = append(response.Items, obj.(*pb.Operation))
		return nil
	}); err != nil {
		return nil, err
	}

	// Sort into reverse order
	sort.Slice(response.Items, func(i, j int) bool {
		return response.Items[i].GetStartTime().AsTime().UnixNano() < response.Items[j].GetStartTime().AsTime().UnixNano()
	})

	return response, nil
}

type OperationName struct {
	Project       *projects.ProjectData
	OperationName string
}

func (n *OperationName) String() string {
	return "projects/" + n.Project.ID + "/operations/" + n.OperationName
}

// parseSQLOperationName parses a string into a OperationName.
// The expected form is projects/<projectID>/Operations/<SQLOperationName>
func (s *MockService) parseOperationName(name string) (*OperationName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "operations" {
		return s.buildOperationName(tokens[1], tokens[3])
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}

func (s *MockService) buildOperationName(projectID, name string) (*OperationName, error) {
	project, err := s.projects.GetProjectByID(projectID)
	if err != nil {
		return nil, err
	}

	return &OperationName{
		Project:       project,
		OperationName: name,
	}, nil
}
