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
// proto.service: google.logging.v2.MetricsServiceV2
// proto.message: google.logging.v2.LogMetric

package mocklogging

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/genproto/googleapis/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/logging/apiv2/loggingpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
)

func (s *metricsServiceV2) GetLogMetric(ctx context.Context, req *pb.GetLogMetricRequest) (*pb.LogMetric, error) {
	name, err := s.parseLogMetricName(req.MetricName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.LogMetric{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Metric %s does not exist.", name.MetricName)
		}
		return nil, err
	}

	return redactForReturn(obj), nil
}

func redactForReturn(obj *pb.LogMetric) *pb.LogMetric {
	redacted := proto.Clone(obj).(*pb.LogMetric)
	if redacted.MetricDescriptor != nil {
		redacted.MetricDescriptor.Metadata = nil
		redacted.MetricDescriptor.LaunchStage = api.LaunchStage_LAUNCH_STAGE_UNSPECIFIED
	}

	return redacted
}

func (s *metricsServiceV2) CreateLogMetric(ctx context.Context, req *pb.CreateLogMetricRequest) (*pb.LogMetric, error) {
	reqName := req.Parent + "/metrics/" + req.GetMetric().GetName()
	name, err := s.parseLogMetricName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	obj := proto.Clone(req.GetMetric()).(*pb.LogMetric)
	obj.Name = name.MetricName
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)

	if obj.MetricDescriptor != nil {
		obj.MetricDescriptor.Description = obj.Description
	}

	s.populateDefaultsForLogMetric(name, obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return redactForReturn(obj), nil
}

func (s *metricsServiceV2) populateDefaultsForLogMetric(name *logMetricName, obj *pb.LogMetric) {

	if obj.MetricDescriptor != nil {
		obj.MetricDescriptor.Name = fmt.Sprintf("projects/%s/metricDescriptors/logging.googleapis.com/user/%s", name.Project.ID, name.MetricName)
		obj.MetricDescriptor.Type = fmt.Sprintf("logging.googleapis.com/user/%s", name.MetricName)
	}
}

func (s *metricsServiceV2) UpdateLogMetric(ctx context.Context, req *pb.UpdateLogMetricRequest) (*pb.LogMetric, error) {
	reqName := req.MetricName

	name, err := s.parseLogMetricName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	existing := &pb.LogMetric{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		return nil, err
	}

	now := time.Now()

	updated := proto.Clone(req.GetMetric()).(*pb.LogMetric)
	updated.Name = name.MetricName
	updated.CreateTime = existing.CreateTime
	updated.UpdateTime = timestamppb.New(now)
	if updated.MetricDescriptor == nil {
		updated.MetricDescriptor = existing.MetricDescriptor
	}
	if updated.MetricDescriptor != nil {
		updated.MetricDescriptor.Description = updated.Description
	}
	s.populateDefaultsForLogMetric(name, updated)

	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}
	return redactForReturn(updated), nil
}

func (s *metricsServiceV2) DeleteLogMetric(ctx context.Context, req *pb.DeleteLogMetricRequest) (*empty.Empty, error) {
	name, err := s.parseLogMetricName(req.MetricName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deletedObj := &pb.LogMetric{}
	if err := s.storage.Delete(ctx, fqn, deletedObj); err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}

type logMetricName struct {
	Project    *projects.ProjectData
	MetricName string
}

func (n *logMetricName) String() string {
	return "projects/" + n.Project.ID + "/metrics/" + n.MetricName
}

// parseLogMetricName parses a string into a logmetricName.
// The expected form is `projects/*/locations/global/logmetrices/*`.
func (s *MockService) parseLogMetricName(name string) (*logMetricName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "metrics" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &logMetricName{
			Project:    project,
			MetricName: tokens[3],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
