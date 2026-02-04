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

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/parametermanager/v1"
)

// Creates a new [Parameter][google.cloud.parametermanager.v1.Parameter].
func (s *ParameterManagerV1) CreateParameterVersion(ctx context.Context, req *pb.CreateParameterVersionRequest) (*pb.ParameterVersion, error) {
	ParameterVersionId := req.ParameterVersionId
	if ParameterVersionId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "ParameterVersionId is required")
	}

	parameter, err := s.parseParameterName(req.Parent)
	if err != nil {
		return nil, err
	}

	name := parameterVersionName{
		parent:           parameter,
		ParameterVersion: ParameterVersionId,
	}
	fqn := name.String()

	obj := proto.Clone(req.ParameterVersion).(*pb.ParameterVersion)
	obj.Name = fqn
	obj.CreateTime = timestamppb.Now()
	obj.UpdateTime = timestamppb.Now()

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

// Gets metadata for a given [Parameter][google.cloud.parametermanager.v1.Parameter].
func (s *ParameterManagerV1) GetParameterVersion(ctx context.Context, req *pb.GetParameterVersionRequest) (*pb.ParameterVersion, error) {
	name, err := s.parseParameterVersionName(req.Name)
	if err != nil {
		return nil, err
	}

	var parameter pb.ParameterVersion
	fqn := name.String()
	if err := s.storage.Get(ctx, fqn, &parameter); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found.", fqn)
		}
		return nil, err
	}

	return &parameter, nil
}

// Deletes a [Parameter][google.cloud.parametermanager.v1.Parameter].
func (s *ParameterManagerV1) DeleteParameterVersion(ctx context.Context, req *pb.DeleteParameterVersionRequest) (*emptypb.Empty, error) {
	name, err := s.parseParameterVersionName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	oldObj := &pb.ParameterVersion{}
	if err := s.storage.Delete(ctx, fqn, oldObj); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

// Update metadata for a given [Parameter][google.cloud.parametermanager.v1.Parameter].
func (s *ParameterManagerV1) UpdateParameterVersion(ctx context.Context, req *pb.UpdateParameterVersionRequest) (*pb.ParameterVersion, error) {
	name, err := s.parseParameterVersionName(req.ParameterVersion.Name)
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	existing := &pb.ParameterVersion{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		return nil, err
	}

	updated := proto.Clone(existing).(*pb.ParameterVersion)
	updated.Name = name.String()

	// Required. The update mask applies to the resource.
	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask is required")
	}

	for _, path := range paths {
		switch path {
		case "disabled":
			updated.Disabled = req.ParameterVersion.GetDisabled()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
		}
	}

	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}
	return updated, nil
}

type parameterVersionName struct {
	parent           *parameterName
	ParameterVersion string
}

func (n *parameterVersionName) String() string {
	return n.parent.String() + "/" + "versions" + "/" + n.ParameterVersion
}

// parseParameterName parses a string into a parameterName.
// The expected form is projects/<projectID>/locations/<location>/parameters/<parameterName>
func (s *MockService) parseParameterVersionName(name string) (*parameterVersionName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "parameters" && tokens[6] == "versions" {
		projectName, err := projects.ParseProjectName("projects/" + tokens[1])
		if err != nil {
			return nil, err
		}

		project, err := s.Projects.GetProject(projectName)
		if err != nil {
			return nil, err
		}

		parameterName := &parameterName{
			Project:       project,
			Location:      tokens[3],
			ParameterName: tokens[5],
		}

		parameterVersionName := &parameterVersionName{
			parent:           parameterName,
			ParameterVersion: tokens[7],
		}

		return parameterVersionName, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
