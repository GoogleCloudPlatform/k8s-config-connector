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

package mockaiplatform

import (
	"context"
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
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

type modelMonitoringService struct {
	*MockService
	pb.UnimplementedModelMonitoringServiceServer
}

type modelMonitorName struct {
	Project        *projects.ProjectData
	Location       string
	ModelMonitorID string
}

func (n *modelMonitorName) String() string {
	return fmt.Sprintf("projects/%d/locations/%s/modelMonitors/%s", n.Project.Number, n.Location, n.ModelMonitorID)
}

func (s *MockService) parseModelMonitorName(name string) (*modelMonitorName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "modelMonitors" {
		projectName, err := projects.ParseProjectName(tokens[0] + "/" + tokens[1])
		if err != nil {
			return nil, err
		}
		project, err := s.Projects.GetProject(projectName)
		if err != nil {
			return nil, err
		}

		name := &modelMonitorName{
			Project:        project,
			Location:       tokens[3],
			ModelMonitorID: tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

type modelMonitorParent struct {
	Project  *projects.ProjectData
	Location string
}

func (s *MockService) parseModelMonitorParent(parent string) (*modelMonitorParent, error) {
	tokens := strings.Split(parent, "/")
	if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "locations" {
		projectName, err := projects.ParseProjectName(tokens[0] + "/" + tokens[1])
		if err != nil {
			return nil, err
		}
		project, err := s.Projects.GetProject(projectName)
		if err != nil {
			return nil, err
		}
		return &modelMonitorParent{
			Project:  project,
			Location: tokens[3],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "parent %q is not valid", parent)
}

func (s *modelMonitoringService) GetModelMonitor(ctx context.Context, req *pb.GetModelMonitorRequest) (*pb.ModelMonitor, error) {
	name, err := s.parseModelMonitorName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.ModelMonitor{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Model Monitor `%s` is not found or active.", req.GetName())
		}
		return nil, err
	}

	return obj, nil
}

func (s *modelMonitoringService) CreateModelMonitor(ctx context.Context, req *pb.CreateModelMonitorRequest) (*longrunning.Operation, error) {
	reqName := req.GetParent() + "/modelMonitors/" + req.GetModelMonitorId()
	name, err := s.parseModelMonitorName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	obj := proto.Clone(req.GetModelMonitor()).(*pb.ModelMonitor)
	obj.Name = fqn

	if obj.DisplayName == "" {
		obj.DisplayName = req.GetModelMonitorId()
	}

	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.AlreadyExists {
			return nil, status.Errorf(codes.FailedPrecondition, "Already existed a model Monitor `(%d, %s)` for model monitoring target. Please delete the existing one.", name.Project.Number, name.ModelMonitorID)
		}
		return nil, err
	}

	metadata := &pb.CreateModelMonitorOperationMetadata{
		GenericMetadata: &pb.GenericOperationMetadata{
			CreateTime: timestamppb.New(now),
			UpdateTime: timestamppb.New(now),
		},
	}
	opPrefix := fmt.Sprintf("projects/%d/locations/%s/modelMonitors/%s", name.Project.Number, name.Location, name.ModelMonitorID)
	return s.operations.StartLRO(ctx, opPrefix, metadata, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *modelMonitoringService) UpdateModelMonitor(ctx context.Context, req *pb.UpdateModelMonitorRequest) (*longrunning.Operation, error) {
	reqMonitor := req.GetModelMonitor()
	name, err := s.parseModelMonitorName(reqMonitor.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.ModelMonitor{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Model Monitor `%s` is not found or active.", reqMonitor.GetName())
		}
		return nil, err
	}

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		obj.DisplayName = reqMonitor.GetDisplayName()
		obj.ModelMonitoringSchema = reqMonitor.GetModelMonitoringSchema()
		obj.NotificationSpec = reqMonitor.GetNotificationSpec()
		obj.DefaultObjective = reqMonitor.GetDefaultObjective()
	} else {
		for _, path := range paths {
			switch path {
			case "displayName", "display_name":
				obj.DisplayName = reqMonitor.GetDisplayName()
			case "modelMonitoringSchema", "model_monitoring_schema":
				obj.ModelMonitoringSchema = reqMonitor.GetModelMonitoringSchema()
			case "notificationSpec", "notification_spec":
				obj.NotificationSpec = reqMonitor.GetNotificationSpec()
			case "tabularObjective", "tabular_objective", "defaultObjective", "default_objective":
				obj.DefaultObjective = reqMonitor.GetDefaultObjective()
			default:
				return nil, status.Errorf(codes.InvalidArgument, "field %q is not supported for update", path)
			}
		}
	}

	now := time.Now()
	obj.UpdateTime = timestamppb.New(now)

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	metadata := &pb.UpdateModelMonitorOperationMetadata{
		GenericMetadata: &pb.GenericOperationMetadata{
			CreateTime: obj.GetCreateTime(),
			UpdateTime: obj.GetUpdateTime(),
		},
	}
	opPrefix := fmt.Sprintf("projects/%d/locations/%s", name.Project.Number, name.Location)
	respObj := proto.Clone(obj).(*pb.ModelMonitor)
	respObj.CreateTime = nil
	respObj.UpdateTime = nil
	return s.operations.DoneLRO(ctx, opPrefix, metadata, respObj)
}

func (s *modelMonitoringService) DeleteModelMonitor(ctx context.Context, req *pb.DeleteModelMonitorRequest) (*longrunning.Operation, error) {
	name, err := s.parseModelMonitorName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.ModelMonitor{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Model Monitor `%s` is not found or active.", req.GetName())
		}
		return nil, err
	}

	deleted := &pb.ModelMonitor{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	now := time.Now()
	metadata := &pb.DeleteOperationMetadata{
		GenericMetadata: &pb.GenericOperationMetadata{
			CreateTime: timestamppb.New(now),
			UpdateTime: timestamppb.New(now),
		},
	}
	opPrefix := fmt.Sprintf("projects/%d/locations/%s", name.Project.Number, name.Location)
	return s.operations.DoneLRO(ctx, opPrefix, metadata, &emptypb.Empty{})
}

func (s *modelMonitoringService) ListModelMonitors(ctx context.Context, req *pb.ListModelMonitorsRequest) (*pb.ListModelMonitorsResponse, error) {
	parent, err := s.parseModelMonitorParent(req.GetParent())
	if err != nil {
		return nil, err
	}

	findPrefix := fmt.Sprintf("projects/%d/locations/%s/modelMonitors/", parent.Project.Number, parent.Location)
	findPrefixID := fmt.Sprintf("projects/%s/locations/%s/modelMonitors/", parent.Project.ID, parent.Location)

	var modelMonitors []*pb.ModelMonitor
	modelMonitorKind := (&pb.ModelMonitor{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, modelMonitorKind, storage.ListOptions{}, func(obj proto.Message) error {
		mm := obj.(*pb.ModelMonitor)
		if strings.HasPrefix(mm.GetName(), findPrefix) || strings.HasPrefix(mm.GetName(), findPrefixID) {
			modelMonitors = append(modelMonitors, mm)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return &pb.ListModelMonitorsResponse{
		ModelMonitors: modelMonitors,
	}, nil
}
