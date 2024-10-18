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
// krm.apiVersion: logging.cnrm.cloud.google.com/v1beta1
// krm.kind: LoggingLogExclusion
// proto.service: google.logging.v2.ConfigServiceV2
// proto.resource: LogExclusion

package mocklogging

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/logging/v2"
)

func (s *configService) GetExclusion(ctx context.Context, req *pb.GetExclusionRequest) (*pb.LogExclusion, error) {
	name, err := s.parseLogExclusionName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.LogExclusion{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Exclusion does not exist: %s", name.exclusionName)
		}
		return nil, err
	}

	return obj, nil
}

func (s *configService) CreateExclusion(ctx context.Context, req *pb.CreateExclusionRequest) (*pb.LogExclusion, error) {
	reqName := fmt.Sprintf("%s/exclusions/%s", req.GetParent(), req.GetExclusion().GetName())
	name, err := s.parseLogExclusionName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	obj := proto.Clone(req.GetExclusion()).(*pb.LogExclusion)
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *configService) UpdateExclusion(ctx context.Context, req *pb.UpdateExclusionRequest) (*pb.LogExclusion, error) {
	reqName := req.Name
	name, err := s.parseLogExclusionName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	existing := &pb.LogExclusion{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		return nil, err
	}
	now := time.Now()
	updated := proto.Clone(existing).(*pb.LogExclusion)

	// Required. The update mask applies to the resource.
	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask is required by mock")
	}

	// TODO: Some sort of helper for fieldmask?
	for _, path := range paths {
		switch path {
		case "description":
			updated.Description = req.GetExclusion().GetDescription()
		case "filter":
			updated.Filter = req.GetExclusion().GetFilter()
		case "disabled":
			updated.Disabled = req.GetExclusion().GetDisabled()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
		}
	}

	updated.UpdateTime = timestamppb.New(now)
	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}
	return updated, nil
}

func (s *configService) DeleteExclusion(ctx context.Context, req *pb.DeleteExclusionRequest) (*empty.Empty, error) {
	name, err := s.parseLogExclusionName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deletedObj := &pb.LogExclusion{}
	if err := s.storage.Delete(ctx, fqn, deletedObj); err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}

type logExclusionName struct {
	// Only one of project/folder/organization/billingAccount should be set
	project        *projects.ProjectData
	folder         string
	organization   string
	billingAccount string

	exclusionName string
}

func (n *logExclusionName) String() string {
	if n.organization != "" {
		return fmt.Sprintf("organizations/%s/exclusions/%s", n.organization, n.exclusionName)
	}
	if n.folder != "" {
		return fmt.Sprintf("folders/%s/exclusions/%s", n.folder, n.exclusionName)
	}
	if n.billingAccount != "" {
		return fmt.Sprintf("billingAccounts/%s/exclusions/%s", n.billingAccount, n.exclusionName)
	}
	return fmt.Sprintf("projects/%s/exclusions/%s", n.project.ID, n.exclusionName)
}

// parseLogExclusionName parses a string into a logExclusionName.
// The expected form is `projects/*/exclusions/*`
func (s *MockService) parseLogExclusionName(name string) (*logExclusionName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "exclusions" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}
		name := &logExclusionName{
			project:       project,
			exclusionName: tokens[3],
		}
		return name, nil
	} else if len(tokens) == 4 && tokens[0] == "folders" && tokens[2] == "exclusions" {
		name := &logExclusionName{
			folder:        tokens[1],
			exclusionName: tokens[3],
		}
		return name, nil
	} else if len(tokens) == 4 && tokens[0] == "organizations" && tokens[2] == "exclusions" {
		name := &logExclusionName{
			organization:  tokens[1],
			exclusionName: tokens[3],
		}
		return name, nil
	} else if len(tokens) == 4 && tokens[0] == "billingAccounts" && tokens[2] == "exclusions" {
		name := &logExclusionName{
			billingAccount: tokens[1],
			exclusionName:  tokens[3],
		}
		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
