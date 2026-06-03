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
// See the License for the description language governing permissions and
// limitations under the License.

package mocknetworksecurity

import (
	"context"
	"fmt"
	"strings"
	"time"

	pbv1 "cloud.google.com/go/networksecurity/apiv1/networksecuritypb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *MirroringServer) GetMirroringDeployment(ctx context.Context, req *pbv1.GetMirroringDeploymentRequest) (*pbv1.MirroringDeployment, error) {
	name, err := s.parseMirroringDeploymentName(req.Name)
	if err != nil {
		return nil, err
	}

	obj := &pbv1.MirroringDeployment{}
	if err := s.storage.Get(ctx, name.String(), obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *MirroringServer) CreateMirroringDeployment(ctx context.Context, req *pbv1.CreateMirroringDeploymentRequest) (*longrunning.Operation, error) {
	name := req.Parent + "/mirroringDeployments/" + req.MirroringDeploymentId

	fqn := name

	obj := proto.CloneOf(req.MirroringDeployment)
	obj.Name = name

	now := time.Now()
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.State = pbv1.MirroringDeployment_ACTIVE
	obj.Reconciling = false

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroMetadata := &pbv1.OperationMetadata{
		CreateTime:            timestamppb.New(now),
		RequestedCancellation: false,
		Target:                name,
		Verb:                  "create",
		ApiVersion:            "v1",
	}
	return s.operations.StartLRO(ctx, req.Parent, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())
		result := proto.CloneOf(obj)
		return result, nil
	})
}

func (s *MirroringServer) UpdateMirroringDeployment(ctx context.Context, req *pbv1.UpdateMirroringDeploymentRequest) (*longrunning.Operation, error) {
	name, err := s.parseMirroringDeploymentName(req.GetMirroringDeployment().GetName())
	if err != nil {
		return nil, err
	}
	obj := &pbv1.MirroringDeployment{}
	if err := s.storage.Get(ctx, name.String(), obj); err != nil {
		return nil, err
	}

	now := time.Now()
	updated := proto.CloneOf(obj)
	updated.UpdateTime = timestamppb.New(now)
	updated.Description = req.GetMirroringDeployment().GetDescription()
	updated.Labels = req.GetMirroringDeployment().GetLabels()

	if err := s.storage.Update(ctx, name.String(), updated); err != nil {
		return nil, err
	}
	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pbv1.OperationMetadata{
		CreateTime:            timestamppb.New(now),
		RequestedCancellation: false,
		Target:                name.String(),
		Verb:                  "update",
		ApiVersion:            "v1",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())
		result := proto.CloneOf(updated)
		return result, nil
	})
}

func (s *MirroringServer) DeleteMirroringDeployment(ctx context.Context, req *pbv1.DeleteMirroringDeploymentRequest) (*longrunning.Operation, error) {
	name, err := s.parseMirroringDeploymentName(req.Name)
	if err != nil {
		return nil, err
	}

	if err := s.storage.Delete(ctx, name.String(), &pbv1.MirroringDeployment{}); err != nil {
		return nil, err
	}
	now := time.Now()
	lroMetadata := &pbv1.OperationMetadata{
		CreateTime:            timestamppb.New(now),
		RequestedCancellation: false,
		Target:                name.String(),
		Verb:                  "delete",
		ApiVersion:            "v1",
	}
	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())
		return &emptypb.Empty{}, nil
	})
}

type mirroringDeploymentName struct {
	Project  *projects.ProjectData
	Location string
	Name     string
}

func (n *mirroringDeploymentName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/mirroringDeployments/%s", n.Project.ID, n.Location, n.Name)
}

func (s *MockService) parseMirroringDeploymentName(name string) (*mirroringDeploymentName, error) {
	if name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "name must be provided")
	}

	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "mirroringDeployments" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}
		return &mirroringDeploymentName{
			Project:  project,
			Location: tokens[3],
			Name:     tokens[5],
		}, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
