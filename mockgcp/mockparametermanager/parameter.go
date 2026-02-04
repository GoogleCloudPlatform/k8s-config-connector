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

package mockparametermanager

import (
	"context"
	"strings"

	"cloud.google.com/go/iam/apiv1/iampb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/parametermanager/v1"
)

type ParameterManagerV1 struct {
	*MockService

	pb.UnimplementedParameterManagerServer
}

// Creates a new [Parameter][google.cloud.parametermanager.v1.Parameter].
func (s *ParameterManagerV1) CreateParameter(ctx context.Context, req *pb.CreateParameterRequest) (*pb.Parameter, error) {
	parameterID := req.ParameterId
	if parameterID == "" {
		return nil, status.Errorf(codes.InvalidArgument, "ParameterId is required")
	}

	reqParameterName := req.Parent + "/parameters/" + parameterID

	name, err := s.parseParameterName(reqParameterName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.Parameter).(*pb.Parameter)
	obj.Name = fqn
	obj.CreateTime = timestamppb.Now()
	obj.UpdateTime = obj.CreateTime
	obj.PolicyMember = &iampb.ResourcePolicyMember{}
	obj.PolicyMember.IamPolicyUidPrincipal = "placeholder value"

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

// Gets metadata for a given [Parameter][google.cloud.parametermanager.v1.Parameter].
func (s *ParameterManagerV1) GetParameter(ctx context.Context, req *pb.GetParameterRequest) (*pb.Parameter, error) {
	name, err := s.parseParameterName(req.Name)
	if err != nil {
		return nil, err
	}

	var parameter pb.Parameter
	fqn := name.String()
	if err := s.storage.Get(ctx, fqn, &parameter); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return &parameter, nil
}

// Update metadata for a given [Parameter][google.cloud.parametermanager.v1.Parameter].
func (s *ParameterManagerV1) UpdateParameter(ctx context.Context, req *pb.UpdateParameterRequest) (*pb.Parameter, error) {
	name, err := s.parseParameterName(req.Parameter.Name)
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	existing := &pb.Parameter{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		return nil, err
	}

	updated := proto.Clone(existing).(*pb.Parameter)
	updated.Name = name.String()
	updated.UpdateTime = timestamppb.Now()

	// Required. The update mask applies to the resource.
	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask is required")
	}
	for _, path := range paths {
		switch path {
		case "kmsKey":
			if kmsKey := req.Parameter.GetKmsKey(); kmsKey != "" {
				updated.KmsKey = &kmsKey
			} else {
				updated.KmsKey = nil
			}
		case "labels":
			updated.Labels = req.Parameter.GetLabels()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
		}
	}

	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}
	return updated, nil
}

// Deletes a [Parameter][google.cloud.parametermanager.v1.Parameter].
func (s *ParameterManagerV1) DeleteParameter(ctx context.Context, req *pb.DeleteParameterRequest) (*emptypb.Empty, error) {
	name, err := s.parseParameterName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	oldObj := &pb.Parameter{}
	if err := s.storage.Delete(ctx, fqn, oldObj); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

type parameterName struct {
	Project       *projects.ProjectData
	Location      string
	ParameterName string
}

func (n *parameterName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/parameters/" + n.ParameterName
}

// parseParameterName parses a string into a parameterName.
// The expected form is projects/<projectID>/locations/<location>/parameters/<parameterName>
func (s *MockService) parseParameterName(name string) (*parameterName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "parameters" {
		projectName, err := projects.ParseProjectName("projects/" + tokens[1])
		if err != nil {
			return nil, err
		}

		project, err := s.Projects.GetProject(projectName)
		if err != nil {
			return nil, err
		}

		name := &parameterName{
			Project:       project,
			Location:      tokens[3],
			ParameterName: tokens[5],
		}

		return name, nil
	} else {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"name %q is not valid. expected format is projects/<projectID>/locations/<location>/parameters/<parameterName>",
			name,
		)
	}
}
