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

package mockaiplatform

import (
	"context"
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"strings"
	"time"

	longrunning "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
)

type tensorboardService struct {
	*MockService
	pb.UnimplementedTensorboardServiceServer
}

func (s *tensorboardService) GetTensorboard(ctx context.Context, req *pb.GetTensorboardRequest) (*pb.Tensorboard, error) {
	name, err := s.parseTensorboardName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Tensorboard{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Tensorboard %s is not found.", req.Name)
		}
		return nil, err
	}

	return obj, nil
}

func (s *tensorboardService) CreateTensorboard(ctx context.Context, req *pb.CreateTensorboardRequest) (*longrunning.Operation, error) {
	id := fmt.Sprintf("%d", time.Now().UnixNano())
	reqName := req.Parent + "/tensorboards/" + id
	name, err := s.parseTensorboardName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	obj := proto.CloneOf(req.Tensorboard)
	obj.Name = fqn

	obj.BlobStoragePathPrefix = "cloud-ai-platform-00000000-1111-2222-3333-444444444444"
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)

	obj.Etag = computeEtag(obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.CreateTensorboardOperationMetadata{}
	op.GenericMetadata = &pb.GenericOperationMetadata{
		CreateTime: timestamppb.New(now),
		UpdateTime: timestamppb.New(now),
	}
	opPrefix := name.String()
	return s.operations.StartLRO(ctx, opPrefix, op, func() (proto.Message, error) {
		// Many fields are not populated in the LRO result
		result := proto.CloneOf(obj)
		result.BlobStoragePathPrefix = ""
		result.CreateTime = nil
		result.UpdateTime = nil
		result.Etag = ""

		return result, nil
	})
}

func (s *tensorboardService) UpdateTensorboard(ctx context.Context, req *pb.UpdateTensorboardRequest) (*longrunning.Operation, error) {
	name, err := s.parseTensorboardName(req.GetTensorboard().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	obj := &pb.Tensorboard{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// See docs for UpdateMask
	updateMask := req.GetUpdateMask()
	for _, path := range updateMask.Paths {
		switch path {
		case "displayName":
			obj.DisplayName = req.GetTensorboard().GetDisplayName()

		case "description":
			obj.Description = req.GetTensorboard().GetDescription()

		default:
			return nil, status.Errorf(codes.InvalidArgument, "field %q is not yet handled in mock", path)
		}
	}

	obj.UpdateTime = timestamppb.New(now)

	obj.Etag = computeEtag(obj)

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.UpdateTensorboardOperationMetadata{}
	op.GenericMetadata = &pb.GenericOperationMetadata{
		CreateTime: timestamppb.New(now),
		UpdateTime: timestamppb.New(now),
	}
	opPrefix := fmt.Sprintf("projects/%d/locations/%s", name.Project.Number, name.Location)
	return s.operations.DoneLRO(ctx, opPrefix, op, obj)
}

func (s *tensorboardService) DeleteTensorboard(ctx context.Context, req *pb.DeleteTensorboardRequest) (*longrunning.Operation, error) {
	name, err := s.parseTensorboardName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	deleted := &pb.Tensorboard{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	op := &pb.DeleteOperationMetadata{}
	op.GenericMetadata = &pb.GenericOperationMetadata{
		CreateTime: timestamppb.New(now),
		UpdateTime: timestamppb.New(now),
	}
	opPrefix := fmt.Sprintf("projects/%d/locations/%s", name.Project.Number, name.Location)
	return s.operations.DoneLRO(ctx, opPrefix, op, &emptypb.Empty{})
}

func (s *tensorboardService) GetTensorboardExperiment(ctx context.Context, req *pb.GetTensorboardExperimentRequest) (*pb.TensorboardExperiment, error) {
	name, err := s.parseTensorboardExperimentName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.TensorboardExperiment{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "TensorboardExperiment %s is not found.", req.Name)
		}
		return nil, err
	}

	return obj, nil
}

func (s *tensorboardService) CreateTensorboardExperiment(ctx context.Context, req *pb.CreateTensorboardExperimentRequest) (*pb.TensorboardExperiment, error) {
	parentName, err := s.parseTensorboardName(req.Parent)
	if err != nil {
		return nil, err
	}

	id := req.TensorboardExperimentId
	if id == "" {
		id = fmt.Sprintf("%d", time.Now().UnixNano())
	}

	fqn := parentName.String() + "/experiments/" + id
	if _, err := s.parseTensorboardExperimentName(fqn); err != nil {
		return nil, err
	}

	now := time.Now()

	obj := proto.Clone(req.TensorboardExperiment).(*pb.TensorboardExperiment)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.Etag = computeEtag(obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *tensorboardService) UpdateTensorboardExperiment(ctx context.Context, req *pb.UpdateTensorboardExperimentRequest) (*pb.TensorboardExperiment, error) {
	name, err := s.parseTensorboardExperimentName(req.GetTensorboardExperiment().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	obj := &pb.TensorboardExperiment{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// Apply the update mask
	paths := req.GetUpdateMask().GetPaths()
	for _, path := range paths {
		switch path {
		case "displayName":
			obj.DisplayName = req.GetTensorboardExperiment().GetDisplayName()
		case "description":
			obj.Description = req.GetTensorboardExperiment().GetDescription()
		case "labels":
			obj.Labels = req.GetTensorboardExperiment().GetLabels()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "field %q is not yet handled in mock", path)
		}
	}

	obj.UpdateTime = timestamppb.New(now)
	obj.Etag = computeEtag(obj)

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *tensorboardService) DeleteTensorboardExperiment(ctx context.Context, req *pb.DeleteTensorboardExperimentRequest) (*longrunning.Operation, error) {
	name, err := s.parseTensorboardExperimentName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	deleted := &pb.TensorboardExperiment{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	op := &pb.DeleteOperationMetadata{}
	op.GenericMetadata = &pb.GenericOperationMetadata{
		CreateTime: timestamppb.New(now),
		UpdateTime: timestamppb.New(now),
	}
	opPrefix := fmt.Sprintf("projects/%d/locations/%s/tensorboards/%s", name.Project.Number, name.Location, name.TensorboardID)
	return s.operations.DoneLRO(ctx, opPrefix, op, &emptypb.Empty{})
}

type TensorboardExperimentName struct {
	Project                 *projects.ProjectData
	Location                string
	TensorboardID           string
	TensorboardExperimentID string
}

func (n *TensorboardExperimentName) String() string {
	return fmt.Sprintf("projects/%d/locations/%s/tensorboards/%s/experiments/%s", n.Project.Number, n.Location, n.TensorboardID, n.TensorboardExperimentID)
}

// parseTensorboardExperimentName parses a string into a TensorboardExperimentName.
func (s *MockService) parseTensorboardExperimentName(name string) (*TensorboardExperimentName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "tensorboards" && tokens[6] == "experiments" {
		projectName, err := projects.ParseProjectName(tokens[0] + "/" + tokens[1])
		if err != nil {
			return nil, err
		}
		project, err := s.Projects.GetProject(projectName)
		if err != nil {
			return nil, err
		}

		name := &TensorboardExperimentName{
			Project:                 project,
			Location:                tokens[3],
			TensorboardID:           tokens[5],
			TensorboardExperimentID: tokens[7],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}

type TensorboardName struct {
	Project       *projects.ProjectData
	Location      string
	TensorboardID string
}

func (n *TensorboardName) String() string {
	return fmt.Sprintf("projects/%d/locations/%s/tensorboards/%s", n.Project.Number, n.Location, n.TensorboardID)
}

// parseTensorboardName parses a string into a TensorboardName.
// The expected form is projects/<projectID>/locations/global/tensorboards/<TensorboardName>
func (s *MockService) parseTensorboardName(name string) (*TensorboardName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "tensorboards" {
		projectName, err := projects.ParseProjectName(tokens[0] + "/" + tokens[1])
		if err != nil {
			return nil, err
		}
		project, err := s.Projects.GetProject(projectName)
		if err != nil {
			return nil, err
		}

		name := &TensorboardName{
			Project:       project,
			Location:      tokens[3],
			TensorboardID: tokens[5],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}

func computeEtag(obj proto.Message) string {
	b, err := proto.Marshal(obj)
	if err != nil {
		panic(fmt.Sprintf("converting to proto: %v", err))
	}
	hash := md5.Sum(b)
	return base64.StdEncoding.EncodeToString(hash[:])
}
