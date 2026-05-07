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

// +tool:mockgcp-support
// proto.service: google.cloud.deploy.v1.CloudDeploy
// proto.message: google.cloud.deploy.v1.Target

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
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/fields"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/deploy/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"github.com/google/uuid"
)

func (s *cloudDeploy) GetTarget(ctx context.Context, req *pb.GetTargetRequest) (*pb.Target, error) {
	name, err := s.parseTargetName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Target{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *cloudDeploy) CreateTarget(ctx context.Context, req *pb.CreateTargetRequest) (*longrunningpb.Operation, error) {
	if req.TargetId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "target_id must be provided")
	}

	reqName := fmt.Sprintf("%s/targets/%s", req.Parent, req.TargetId)
	name, err := s.parseTargetName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := proto.Clone(req.Target).(*pb.Target)
	obj.Name = fqn
	obj.TargetId = name.Target

	obj.Uid = uuid.NewString()
	obj.CreateTime = timestamppb.New(time.Now())
	obj.UpdateTime = timestamppb.New(time.Now())
	obj.Etag = uuid.NewString()

	s.defaultTarget(name, obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// By default, immediately finish the LRO with success.
	lroPrefix := name.LocationPrefix()
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

func (s *cloudDeploy) UpdateTarget(ctx context.Context, req *pb.UpdateTargetRequest) (*longrunningpb.Operation, error) {
	if req.Target == nil {
		return nil, status.Errorf(codes.InvalidArgument, "target must be provided")
	}

	name, err := s.parseTargetName(req.Target.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Target{}
	exists := true
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound && req.GetAllowMissing() {
			exists = false
			obj = proto.Clone(req.Target).(*pb.Target)
			obj.Name = fqn
			obj.TargetId = name.Target
			obj.Uid = uuid.NewString()
			obj.CreateTime = timestamppb.New(time.Now())
			obj.UpdateTime = timestamppb.New(time.Now())
			obj.Etag = uuid.NewString()
			s.defaultTarget(name, obj)
		} else {
			if status.Code(err) == codes.NotFound {
				return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
			}
			return nil, err
		}
	}

	if exists {
		// Apply the update mask to the object.
		paths := req.GetUpdateMask().GetPaths()
		if len(paths) == 0 {
			return nil, status.Errorf(codes.InvalidArgument, "update_mask must be provided")
		}

		if err := fields.UpdateByFieldMask(obj, req.Target, req.UpdateMask.Paths); err != nil {
			return nil, fmt.Errorf("update field_mask.paths: %w", err)
		}

		obj.UpdateTime = timestamppb.New(time.Now())
		obj.Etag = uuid.NewString()

		s.defaultTarget(name, obj)

		if err := s.storage.Update(ctx, fqn, obj); err != nil {
			return nil, err
		}
	} else {
		if err := s.storage.Create(ctx, fqn, obj); err != nil {
			return nil, err
		}
	}

	lroPrefix := name.LocationPrefix()
	lroMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(time.Now()),
		Target:     name.String(),
		Verb:       "update",
		ApiVersion: "v1",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}

func (s *cloudDeploy) ListTargets(ctx context.Context, req *pb.ListTargetsRequest) (*pb.ListTargetsResponse, error) {
	parent, err := s.parseLocationName(req.Parent)
	if err != nil {
		return nil, err
	}

	var targets []*pb.Target
	targetKind := (&pb.Target{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, targetKind, storage.ListOptions{}, func(obj proto.Message) error {
		target := obj.(*pb.Target)
		name, err := s.parseTargetName(target.Name)
		if err != nil {
			return nil // Should not happen
		}

		if name.Project.ID == parent.Project.ID && name.Location == parent.Location {
			targets = append(targets, target)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return &pb.ListTargetsResponse{
		Targets: targets,
	}, nil
}

func (s *cloudDeploy) DeleteTarget(ctx context.Context, req *pb.DeleteTargetRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseTargetName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	if err := s.storage.Delete(ctx, fqn, &pb.Target{}); err != nil {
		if status.Code(err) == codes.NotFound && req.GetAllowMissing() {
			// Return success (LRO) if not found and AllowMissing is true
		} else {
			if status.Code(err) == codes.NotFound {
				return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
			}
			return nil, err
		}
	}

	// By default, immediately finish the LRO with success.
	lroPrefix := name.LocationPrefix()
	lroMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(time.Now()),
		Target:     name.String(),
		Verb:       "delete",
		ApiVersion: "v1",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return &emptypb.Empty{}, nil
	})
}

func (s *cloudDeploy) defaultTarget(name *targetName, obj *pb.Target) {
	if len(obj.ExecutionConfigs) == 0 {
		sa := fmt.Sprintf("%d-compute@developer.gserviceaccount.com", name.Project.Number)
		bucket := fmt.Sprintf("gs://%s.deploy-artifacts.%s.appspot.com", name.Location, name.Project.ID)

		obj.ExecutionConfigs = []*pb.ExecutionConfig{
			{
				Usages: []pb.ExecutionConfig_ExecutionEnvironmentUsage{
					pb.ExecutionConfig_RENDER,
					pb.ExecutionConfig_DEPLOY,
					pb.ExecutionConfig_VERIFY,
					pb.ExecutionConfig_PREDEPLOY,
					pb.ExecutionConfig_POSTDEPLOY,
				},
				ExecutionTimeout: durationpb.New(time.Hour),
				ServiceAccount:   sa,
				ArtifactStorage:  bucket,
				ExecutionEnvironment: &pb.ExecutionConfig_DefaultPool{
					DefaultPool: &pb.DefaultPool{
						ServiceAccount:  sa,
						ArtifactStorage: bucket,
					},
				},
			},
		}
	}
}

type targetName struct {
	Project  *projects.ProjectData
	Location string
	Target   string
}

func (n *targetName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/targets/%s", n.Project.ID, n.Location, n.Target)
}

func (n *targetName) LocationPrefix() string {
	return fmt.Sprintf("projects/%s/locations/%s", n.Project.ID, n.Location)
}

// parseTargetName parses a string into a targetName.
// The expected form is `projects/*/locations/*/targets/*`.
func (s *MockService) parseTargetName(name string) (*targetName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "targets" {
		for i := 1; i < len(tokens); i += 2 {
			if tokens[i] == "" {
				return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
			}
		}

		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &targetName{
			Project:  project,
			Location: tokens[3],
			Target:   tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
