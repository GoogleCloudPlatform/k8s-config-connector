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

package mockbinaryauthorization

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/binaryauthorization/v1"
)

type PlatformPolicyV1 struct {
	*MockService

	pb.UnimplementedPlatformPolicyViewerServer
}

func (s *PlatformPolicyV1) GetPlatformPolicy(ctx context.Context, req *pb.GetPlatformPolicyRequest) (*pb.PlatformPolicy, error) {
	name, err := s.parsePlatformPolicyName(req.Name)
	if err != nil {
		return nil, err
	}

	var obj pb.PlatformPolicy
	fqn := name.String()
	if err := s.storage.Get(ctx, fqn, &obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "PlatformPolicy %q not found.", fqn)
		}
		return nil, err
	}

	return &obj, nil
}

func (s *PlatformPolicyV1) CreatePlatformPolicy(ctx context.Context, req *pb.CreatePlatformPolicyRequest) (*pb.PlatformPolicy, error) {
	policyID := req.PolicyId
	if policyID == "" {
		return nil, status.Errorf(codes.InvalidArgument, "PolicyId is required")
	}

	platformName, err := s.parsePlatformName(req.Parent)
	if err != nil {
		return nil, err
	}

	name := platformPolicyName{
		platformName: platformName,
		PolicyId:     policyID,
	}
	fqn := name.String()

	obj := proto.Clone(req.PlatformPolicy).(*pb.PlatformPolicy)
	obj.Name = fqn
	obj.UpdateTime = timestamppb.Now()
	obj.Etag = "etag-12345678" // Arbitrary etag for mocks

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *PlatformPolicyV1) ReplacePlatformPolicy(ctx context.Context, req *pb.ReplacePlatformPolicyRequest) (*pb.PlatformPolicy, error) {
	name, err := s.parsePlatformPolicyName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	var actual pb.PlatformPolicy
	if err := s.storage.Get(ctx, fqn, &actual); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "PlatformPolicy %q not found.", fqn)
		}
		return nil, err
	}

	obj := proto.Clone(req.PlatformPolicy).(*pb.PlatformPolicy)
	obj.Name = fqn
	obj.UpdateTime = timestamppb.Now()
	obj.Etag = "etag-updated"

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *PlatformPolicyV1) DeletePlatformPolicy(ctx context.Context, req *pb.DeletePlatformPolicyRequest) (*emptypb.Empty, error) {
	name, err := s.parsePlatformPolicyName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	var obj pb.PlatformPolicy
	if err := s.storage.Delete(ctx, fqn, &obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "PlatformPolicy %q not found.", fqn)
		}
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

type platformName struct {
	Project  *projects.ProjectData
	Platform string
}

func (p platformName) String() string {
	return fmt.Sprintf("projects/%s/platforms/%s", p.Project.ID, p.Platform)
}

type platformPolicyName struct {
	platformName
	PolicyId string
}

func (p platformPolicyName) String() string {
	return fmt.Sprintf("%s/policies/%s", p.platformName.String(), p.PolicyId)
}

func (s *PlatformPolicyV1) parsePlatformName(name string) (platformName, error) {
	// projects/{project}/platforms/{platform}
	tokens := strings.Split(name, "/")
	if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "platforms" {
		projectName, err := projects.ParseProjectName("projects/" + tokens[1])
		if err != nil {
			return platformName{}, err
		}
		project, err := s.Projects.GetProject(projectName)
		if err != nil {
			return platformName{}, err
		}
		return platformName{
			Project:  project,
			Platform: tokens[3],
		}, nil
	}
	return platformName{}, status.Errorf(codes.InvalidArgument, "invalid platform name %q", name)
}

func (s *PlatformPolicyV1) parsePlatformPolicyName(name string) (platformPolicyName, error) {
	// projects/{project}/platforms/{platform}/policies/{policy}
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "platforms" && tokens[4] == "policies" {
		projectName, err := projects.ParseProjectName("projects/" + tokens[1])
		if err != nil {
			return platformPolicyName{}, err
		}
		project, err := s.Projects.GetProject(projectName)
		if err != nil {
			return platformPolicyName{}, err
		}
		return platformPolicyName{
			platformName: platformName{
				Project:  project,
				Platform: tokens[3],
			},
			PolicyId: tokens[5],
		}, nil
	}
	return platformPolicyName{}, status.Errorf(codes.InvalidArgument, "invalid platform policy name %q", name)
}
