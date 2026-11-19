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
// proto.message: google.cloud.deploy.v1.Rollout

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

func (s *cloudDeploy) ListRollouts(ctx context.Context, req *pb.ListRolloutsRequest) (*pb.ListRolloutsResponse, error) {
	parent, err := s.parseReleaseName(req.Parent)
	if err != nil {
		return nil, err
	}

	var rollouts []*pb.Rollout
	rolloutKind := (&pb.Rollout{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, rolloutKind, storage.ListOptions{}, func(message proto.Message) error {
		rollout := message.(*pb.Rollout)
		name, err := s.parseRolloutName(rollout.Name)
		if err != nil {
			return nil
		}

		if name.Project.ID == parent.Project.ID && name.Location == parent.Location && name.DeliveryPipeline == parent.DeliveryPipeline && name.Release == parent.Release {
			rolloutCopy := proto.Clone(rollout).(*pb.Rollout)
			populateDefaultsForRollout(rolloutCopy)
			rollouts = append(rollouts, rolloutCopy)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return &pb.ListRolloutsResponse{
		Rollouts: rollouts,
	}, nil
}

func (s *cloudDeploy) GetRollout(ctx context.Context, req *pb.GetRolloutRequest) (*pb.Rollout, error) {
	name, err := s.parseRolloutName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Rollout{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	response := proto.Clone(obj).(*pb.Rollout)
	populateDefaultsForRollout(response)
	return response, nil
}

func (s *cloudDeploy) CreateRollout(ctx context.Context, req *pb.CreateRolloutRequest) (*longrunningpb.Operation, error) {
	parent, err := s.parseReleaseName(req.Parent)
	if err != nil {
		return nil, err
	}

	releaseFQN := parent.String()
	release := &pb.Release{}
	if err := s.storage.Get(ctx, releaseFQN, release); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", releaseFQN)
		}
		return nil, err
	}

	reqName := fmt.Sprintf("%s/rollouts/%s", req.Parent, req.RolloutId)
	name, err := s.parseRolloutName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := proto.Clone(req.Rollout).(*pb.Rollout)
	obj.Name = fqn

	obj.Uid = uuid.NewString()
	obj.CreateTime = timestamppb.New(time.Now())
	obj.Etag = uuid.NewString()
	populateDefaultsForRollout(obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

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

func populateDefaultsForRollout(obj *pb.Rollout) {
	if obj.ApprovalState == pb.Rollout_APPROVAL_STATE_UNSPECIFIED {
		obj.ApprovalState = pb.Rollout_DOES_NOT_NEED_APPROVAL
	}
	if obj.State == pb.Rollout_STATE_UNSPECIFIED {
		obj.State = pb.Rollout_PENDING_RELEASE
	}
	if len(obj.Phases) == 0 {
		obj.Phases = []*pb.Phase{
			{
				Id:    "stable",
				State: pb.Phase_PENDING,
				Jobs: &pb.Phase_DeploymentJobs{
					DeploymentJobs: &pb.DeploymentJobs{
						DeployJob: &pb.Job{
							Id:    "deploy",
							State: pb.Job_PENDING,
							JobType: &pb.Job_DeployJob{
								DeployJob: &pb.DeployJob{},
							},
						},
						VerifyJob: &pb.Job{
							Id:    "verify",
							State: pb.Job_DISABLED,
							JobType: &pb.Job_VerifyJob{
								VerifyJob: &pb.VerifyJob{},
							},
						},
					},
				},
			},
		}
	}
}

type rolloutName struct {
	Project          *projects.ProjectData
	Location         string
	DeliveryPipeline string
	Release          string
	Rollout          string
}

func (n *rolloutName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/deliveryPipelines/%s/releases/%s/rollouts/%s", n.Project.ID, n.Location, n.DeliveryPipeline, n.Release, n.Rollout)
}

func (s *MockService) parseRolloutName(name string) (*rolloutName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 10 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "deliveryPipelines" && tokens[6] == "releases" && tokens[8] == "rollouts" {
		for i := 1; i < len(tokens); i += 2 {
			if tokens[i] == "" {
				return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
			}
		}

		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &rolloutName{
			Project:          project,
			Location:         tokens[3],
			DeliveryPipeline: tokens[5],
			Release:          tokens[7],
			Rollout:          tokens[9],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
