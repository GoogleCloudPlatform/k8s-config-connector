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

// +tool:mockgcp-support
// proto.service: google.logging.v2.ConfigServiceV2
// proto.message: google.logging.v2.LogExclusion

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

	pb "cloud.google.com/go/logging/apiv2/loggingpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
)

func (s *configServiceV2) GetExclusion(ctx context.Context, req *pb.GetExclusionRequest) (*pb.LogExclusion, error) {
	name, err := s.parseLogExclusionName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.LogExclusion{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Exclusion %s does not exist", req.Name)
		}
		return nil, err
	}
	return obj, nil
}

func (s *configServiceV2) CreateExclusion(ctx context.Context, req *pb.CreateExclusionRequest) (*pb.LogExclusion, error) {
	reqName := fmt.Sprintf("%s/exclusions/%s", req.Parent, req.GetExclusion().GetName())
	name, err := s.parseLogExclusionName(reqName)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := proto.CloneOf(req.GetExclusion())
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(time.Now())
	obj.UpdateTime = timestamppb.New(time.Now())

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *configServiceV2) UpdateExclusion(ctx context.Context, req *pb.UpdateExclusionRequest) (*pb.LogExclusion, error) {
	name, err := s.parseLogExclusionName(req.Name)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	existing := &pb.LogExclusion{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		return nil, err
	}

	updated := proto.CloneOf(existing)

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		// Full update if no update mask
		updated.Description = req.GetExclusion().GetDescription()
		updated.Disabled = req.GetExclusion().GetDisabled()
		updated.Filter = req.GetExclusion().GetFilter()
	} else {
		for _, path := range paths {
			switch path {
			case "description":
				updated.Description = req.GetExclusion().GetDescription()
			case "disabled":
				updated.Disabled = req.GetExclusion().GetDisabled()
			case "filter":
				updated.Filter = req.GetExclusion().GetFilter()
			default:
				return nil, status.Errorf(codes.InvalidArgument, "field %q is read-only or not supported", path)
			}
		}
	}

	updated.UpdateTime = timestamppb.New(time.Now())

	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}
	return updated, nil
}

func (s *configServiceV2) DeleteExclusion(ctx context.Context, req *pb.DeleteExclusionRequest) (*empty.Empty, error) {
	name, err := s.parseLogExclusionName(req.Name)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	deletedObj := &pb.LogExclusion{}
	if err := s.storage.Delete(ctx, fqn, deletedObj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Exclusion %s does not exist", req.Name)
		}
		return nil, err
	}
	return &empty.Empty{}, nil
}

type logExclusionName struct {
	billingAccount string
	folder         string
	organization   string
	project        *projects.ProjectData
	exclusionId    string
}

func (n *logExclusionName) String() string {
	if n.project != nil {
		return fmt.Sprintf("projects/%s/exclusions/%s", n.project.ID, n.exclusionId)
	}
	if n.folder != "" {
		return fmt.Sprintf("folders/%s/exclusions/%s", n.folder, n.exclusionId)
	}
	if n.organization != "" {
		return fmt.Sprintf("organizations/%s/exclusions/%s", n.organization, n.exclusionId)
	}
	if n.billingAccount != "" {
		return fmt.Sprintf("billingAccounts/%s/exclusions/%s", n.billingAccount, n.exclusionId)
	}
	return ""
}

func (s *configServiceV2) parseLogExclusionName(name string) (*logExclusionName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 4 {
		if tokens[2] == "exclusions" {
			if tokens[0] == "projects" {
				project, err := s.Projects.GetProjectByID(tokens[1])
				if err != nil {
					return nil, err
				}
				return &logExclusionName{project: project, exclusionId: tokens[3]}, nil
			}
			if tokens[0] == "folders" {
				return &logExclusionName{folder: tokens[1], exclusionId: tokens[3]}, nil
			}
			if tokens[0] == "organizations" {
				return &logExclusionName{organization: tokens[1], exclusionId: tokens[3]}, nil
			}
			if tokens[0] == "billingAccounts" {
				return &logExclusionName{billingAccount: tokens[1], exclusionId: tokens[3]}, nil
			}
		}
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not a valid LogExclusion name", name)
}
