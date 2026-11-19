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
// proto.service: google.cloud.deploy.v1.CloudDeploy
// proto.message: google.cloud.deploy.v1.Release

package mockclouddeploy

import (
	"context"
	"fmt"
	"strings"
	"time"

	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/deploy/apiv1/deploypb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"github.com/google/uuid"
)

func (s *cloudDeploy) ListReleases(ctx context.Context, req *pb.ListReleasesRequest) (*pb.ListReleasesResponse, error) {
	parent, err := s.parseDeliveryPipelineName(req.Parent)
	if err != nil {
		return nil, err
	}

	var releases []*pb.Release
	releaseKind := (&pb.Release{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, releaseKind, storage.ListOptions{}, func(message proto.Message) error {
		release := message.(*pb.Release)
		name, err := s.parseReleaseName(release.Name)
		if err != nil {
			return nil // Should not happen
		}

		if name.Project.ID == parent.Project.ID && name.Location == parent.Location && name.DeliveryPipeline == parent.DeliveryPipeline {
			releases = append(releases, release)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return &pb.ListReleasesResponse{
		Releases: releases,
	}, nil
}

func (s *cloudDeploy) GetRelease(ctx context.Context, req *pb.GetReleaseRequest) (*pb.Release, error) {
	name, err := s.parseReleaseName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Release{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	response := proto.Clone(obj).(*pb.Release)
	// In real GCP, after creation the render process starts asynchronously, populating targetArtifacts and targetRenders.
	if response.TargetArtifacts == nil {
		response.TargetArtifacts = make(map[string]*pb.TargetArtifact)
	}
	if response.TargetRenders == nil {
		response.TargetRenders = make(map[string]*pb.Release_TargetRender)
	}
	for _, targetSnapshot := range response.TargetSnapshots {
		targetID := targetSnapshot.TargetId
		if targetID == "" {
			tokens := strings.Split(targetSnapshot.Name, "/")
			targetID = tokens[len(tokens)-1]
		}
		if _, exists := response.TargetArtifacts[targetID]; !exists {
			response.TargetArtifacts[targetID] = &pb.TargetArtifact{}
		}
		if _, exists := response.TargetRenders[targetID]; !exists {
			response.TargetRenders[targetID] = &pb.Release_TargetRender{
				RenderingBuild: fmt.Sprintf("projects/%d/locations/%s/builds/00000000-0000-0000-0000-000000000001", name.Project.Number, name.Location),
				RenderingState: pb.Release_TargetRender_IN_PROGRESS,
			}
		}
	}

	return response, nil
}

func (s *cloudDeploy) CreateRelease(ctx context.Context, req *pb.CreateReleaseRequest) (*longrunningpb.Operation, error) {
	if req.ReleaseId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "release_id must be provided")
	}

	reqName := fmt.Sprintf("%s/releases/%s", req.Parent, req.ReleaseId)
	name, err := s.parseReleaseName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := proto.CloneOf(req.Release)
	obj.Name = fqn

	obj.Uid = uuid.NewString()
	obj.CreateTime = timestamppb.New(time.Now())
	obj.Etag = uuid.NewString()

	// Default SkaffoldVersion if unset
	if obj.SkaffoldVersion == "" {
		obj.SkaffoldVersion = "cd-skaffold"
	}

	// Default RenderState to IN_PROGRESS
	obj.RenderState = pb.Release_IN_PROGRESS

	// Default Conditions
	obj.Condition = &pb.Release_ReleaseCondition{
		ReleaseReadyCondition: &pb.Release_ReleaseReadyCondition{
			Status: true,
		},
		SkaffoldSupportedCondition: &pb.Release_SkaffoldSupportedCondition{
			Status:               true,
			SkaffoldSupportState: pb.SkaffoldSupportState_SKAFFOLD_SUPPORT_STATE_SUPPORTED,
		},
	}
	// TODO: toolVersionSupportedConditions (docker, helm, kpt, kubectl, kustomize, skaffoldVersion)
	// and toolVersions are present in real GCP responses but not present in deploypb.Release.

	// Take deliveryPipelineSnapshot
	pipelineFQN := fmt.Sprintf("projects/%s/locations/%s/deliveryPipelines/%s", name.Project.ID, name.Location, name.DeliveryPipeline)
	pipeline := &pb.DeliveryPipeline{}
	if err := s.storage.Get(ctx, pipelineFQN, pipeline); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", pipelineFQN)
		}
		return nil, err
	}

	pipelineSnapshot := proto.Clone(pipeline).(*pb.DeliveryPipeline)
	pipelineSnapshot.Name = fmt.Sprintf("projects/%d/locations/%s/deliveryPipelines/%s", name.Project.Number, name.Location, name.DeliveryPipeline)
	obj.DeliveryPipelineSnapshot = pipelineSnapshot

	// Take targetSnapshots from serialPipeline stages
	var targetSnapshots []*pb.Target
	if serialPipeline := pipeline.GetSerialPipeline(); serialPipeline != nil {
		for _, stage := range serialPipeline.GetStages() {
			targetFQN := fmt.Sprintf("projects/%s/locations/%s/targets/%s", name.Project.ID, name.Location, stage.GetTargetId())
			target := &pb.Target{}
			if err := s.storage.Get(ctx, targetFQN, target); err == nil {
				targetSnapshot := proto.Clone(target).(*pb.Target)
				targetSnapshot.Name = fmt.Sprintf("projects/%d/locations/%s/targets/%s", name.Project.Number, name.Location, stage.GetTargetId())
				targetSnapshots = append(targetSnapshots, targetSnapshot)
			}
		}
	}
	obj.TargetSnapshots = targetSnapshots

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// By default, immediately finish the LRO with success.
	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(time.Now()),
		Target:     name.String(),
		Verb:       "create",
		ApiVersion: "v1",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}

func (s *cloudDeploy) AbandonRelease(ctx context.Context, req *pb.AbandonReleaseRequest) (*pb.AbandonReleaseResponse, error) {
	name, err := s.parseReleaseName(req.Name)
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	obj := &pb.Release{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}
	obj.Abandoned = true
	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}
	return &pb.AbandonReleaseResponse{}, nil
}

type releaseName struct {
	Project          *projects.ProjectData
	Location         string
	DeliveryPipeline string
	Release          string
}

func (n *releaseName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/deliveryPipelines/%s/releases/%s", n.Project.ID, n.Location, n.DeliveryPipeline, n.Release)
}

func (s *MockService) parseReleaseName(name string) (*releaseName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "deliveryPipelines" && tokens[6] == "releases" {
		for i := 1; i < len(tokens); i += 2 {
			if tokens[i] == "" {
				return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
			}
		}

		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &releaseName{
			Project:          project,
			Location:         tokens[3],
			DeliveryPipeline: tokens[5],
			Release:          tokens[7],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
