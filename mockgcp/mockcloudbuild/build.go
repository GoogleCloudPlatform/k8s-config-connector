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

package mockcloudbuild

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/devtools/cloudbuild/v1"
)

type buildName struct {
	Project  *projects.ProjectData
	Location string
	BuildID  string
}

func (n *buildName) String() string {
	if n.Location != "" {
		return fmt.Sprintf("projects/%d/locations/%s/builds/%s", n.Project.Number, n.Location, n.BuildID)
	}
	return fmt.Sprintf("projects/%d/builds/%s", n.Project.Number, n.BuildID)
}

func (n *buildName) FQN() string {
	if n.Location != "" {
		return fmt.Sprintf("projects/%s/locations/%s/builds/%s", n.Project.ID, n.Location, n.BuildID)
	}
	return fmt.Sprintf("projects/%s/builds/%s", n.Project.ID, n.BuildID)
}

func (s *MockService) parseParent(parent string, projectID string) (*projects.ProjectData, string, error) {
	if parent != "" {
		tokens := strings.Split(parent, "/")
		if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "locations" {
			project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
			if err != nil {
				return nil, "", err
			}
			return project, tokens[3], nil
		}
		if len(tokens) == 2 && tokens[0] == "projects" {
			project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
			if err != nil {
				return nil, "", err
			}
			return project, "", nil
		}
		return nil, "", status.Errorf(codes.InvalidArgument, "invalid parent %q", parent)
	}
	if projectID != "" {
		project, err := s.Projects.GetProjectByIDOrNumber(projectID)
		if err != nil {
			return nil, "", err
		}
		return project, "", nil
	}
	return nil, "", status.Errorf(codes.InvalidArgument, "parent or project_id must be specified")
}

func (s *MockService) parseBuildName(nameStr string, projectID string, id string) (*buildName, error) {
	if nameStr != "" {
		tokens := strings.Split(nameStr, "/")
		if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "builds" {
			project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
			if err != nil {
				return nil, err
			}
			return &buildName{
				Project:  project,
				Location: tokens[3],
				BuildID:  tokens[5],
			}, nil
		}
		if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "builds" {
			project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
			if err != nil {
				return nil, err
			}
			return &buildName{
				Project:  project,
				Location: "",
				BuildID:  tokens[3],
			}, nil
		}
		return nil, status.Errorf(codes.InvalidArgument, "invalid build name %q", nameStr)
	}
	if projectID != "" && id != "" {
		project, err := s.Projects.GetProjectByIDOrNumber(projectID)
		if err != nil {
			return nil, err
		}
		return &buildName{
			Project:  project,
			Location: "",
			BuildID:  id,
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name or project_id and id must be specified")
}

func (s *CloudBuildV1) CreateBuild(ctx context.Context, req *pb.CreateBuildRequest) (*longrunning.Operation, error) {
	project, location, err := s.parseParent(req.GetParent(), req.GetProjectId())
	if err != nil {
		return nil, err
	}

	buildID := uuid.New().String()
	name := &buildName{
		Project:  project,
		Location: location,
		BuildID:  buildID,
	}

	build := proto.Clone(req.GetBuild()).(*pb.Build)
	if build == nil {
		build = &pb.Build{}
	}

	now := timestamppb.New(time.Now())

	build.Name = name.String()
	build.Id = name.BuildID
	build.ProjectId = name.Project.ID
	build.CreateTime = now

	s.populateDefaults(name, build)

	// Prepare initial build (for LRO metadata)
	initialBuild := proto.Clone(build).(*pb.Build)
	initialBuild.Status = pb.Build_QUEUED

	// Prepare final completed build
	finalBuild := proto.Clone(build).(*pb.Build)
	finalBuild.Status = pb.Build_SUCCESS
	finalBuild.StartTime = now
	finalBuild.FinishTime = now
	finalBuild.SourceProvenance = &pb.SourceProvenance{}

	results := &pb.Results{}
	for range finalBuild.Steps {
		results.BuildStepImages = append(results.BuildStepImages, "sha256:589adbc588d8ef3c1ec3c68333faf4baf5dccb65d3cb50f6cf3d78cad2d6e1b0")
		results.BuildStepOutputs = append(results.BuildStepOutputs, []byte(""))
	}
	finalBuild.Results = results

	for _, step := range finalBuild.Steps {
		step.Status = pb.Build_SUCCESS
		step.PullTiming = &pb.TimeSpan{
			StartTime: now,
			EndTime:   now,
		}
		step.Timing = &pb.TimeSpan{
			StartTime: now,
			EndTime:   now,
		}
	}

	finalBuild.Timing = map[string]*pb.TimeSpan{
		"BUILD": {
			StartTime: now,
			EndTime:   now,
		},
	}

	if err := s.storage.Create(ctx, name.FQN(), finalBuild); err != nil {
		return nil, err
	}

	opPrefix := fmt.Sprintf("operations/build/%s/{{operationID}}", name.Project.ID)
	metadata := &pb.BuildOperationMetadata{
		Build: initialBuild,
	}
	return s.operations.StartLRO(ctx, opPrefix, metadata, func() (proto.Message, error) {
		metadata.Build = finalBuild
		return finalBuild, nil
	})
}

func (s *CloudBuildV1) populateDefaults(name *buildName, build *pb.Build) {
	if build.QueueTtl == nil {
		build.QueueTtl = &durationpb.Duration{Seconds: 3600}
	}
	if build.Timeout == nil {
		build.Timeout = &durationpb.Duration{Seconds: 3600}
	}
	if build.ServiceAccount == "" {
		build.ServiceAccount = fmt.Sprintf("projects/%s/serviceAccounts/%d-compute@developer.gserviceaccount.com", name.Project.ID, name.Project.Number)
	}

	if build.Options == nil {
		build.Options = &pb.BuildOptions{}
	}
	if build.Options.Logging == pb.BuildOptions_LOGGING_UNSPECIFIED {
		build.Options.Logging = pb.BuildOptions_LEGACY
	}
	if build.Options.Pool == nil {
		build.Options.Pool = &pb.BuildOptions_PoolOption{}
	}

	if build.Options.Logging != pb.BuildOptions_CLOUD_LOGGING_ONLY && build.LogsBucket == "" {
		if name.Location != "" {
			build.LogsBucket = fmt.Sprintf("gs://%d-%s-cloudbuild-logs", name.Project.Number, name.Location)
		} else {
			build.LogsBucket = fmt.Sprintf("gs://%d-cloudbuild-logs", name.Project.Number)
		}
	}

	if build.LogUrl == "" {
		if name.Location != "" && name.Location != "global" {
			build.LogUrl = fmt.Sprintf("https://console.cloud.google.com/cloud-build/builds;region=%s/%s?project=%d", name.Location, name.BuildID, name.Project.Number)
		} else {
			build.LogUrl = fmt.Sprintf("https://console.cloud.google.com/cloud-build/builds/%s?project=%d", name.BuildID, name.Project.Number)
		}
	}

	// TODO: Document proto3 zero-value gaps. Fields such as dynamicSubstitutions=false
	// or empty results are omitted by proto3 JSON serialization unless explicitly populated.
}

func (s *CloudBuildV1) GetBuild(ctx context.Context, req *pb.GetBuildRequest) (*pb.Build, error) {
	name, err := s.parseBuildName(req.GetName(), req.GetProjectId(), req.GetId())
	if err != nil {
		return nil, err
	}

	fqn := name.FQN()
	obj := &pb.Build{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "build %q not found", name.String())
		}
		return nil, err
	}
	s.populateDefaults(name, obj)
	return obj, nil
}
