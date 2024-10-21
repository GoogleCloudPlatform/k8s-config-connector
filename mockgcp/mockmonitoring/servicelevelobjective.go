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
// krm.apiVersion: monitoring.cnrm.cloud.google.com/v1beta1
// krm.kind: MonitoringServiceLevelObjective
// proto.service: google.monitoring.v3.ServiceMonitoringService
// proto.resource: ServiceLevelObjective

package mockmonitoring

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/monitoring/v3"
)

func (s *serviceMonitoringService) GetServiceLevelObjective(ctx context.Context, req *pb.GetServiceLevelObjectiveRequest) (*pb.ServiceLevelObjective, error) {
	name, err := s.parseServiceLevelObjectiveName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.ServiceLevelObjective{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "There is no slo with id %q under service %q in project %q", name.sloID, name.serviceID, name.Project.ID)
		}
		return nil, err
	}
	return obj, nil
}

func (s *serviceMonitoringService) CreateServiceLevelObjective(ctx context.Context, req *pb.CreateServiceLevelObjectiveRequest) (*pb.ServiceLevelObjective, error) {
	reqName := fmt.Sprintf("%s/serviceLevelObjectives/%s", req.GetParent(), req.GetServiceLevelObjectiveId())
	name, err := s.parseServiceLevelObjectiveName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.GetServiceLevelObjective()).(*pb.ServiceLevelObjective)
	obj.Name = fqn
	s.populateDefaultsForServiceLevelObjective(obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *serviceMonitoringService) populateDefaultsForServiceLevelObjective(obj *pb.ServiceLevelObjective) {

}

func (s *serviceMonitoringService) UpdateServiceLevelObjective(ctx context.Context, req *pb.UpdateServiceLevelObjectiveRequest) (*pb.ServiceLevelObjective, error) {
	name, err := s.parseServiceLevelObjectiveName(req.GetServiceLevelObjective().GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.ServiceLevelObjective{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	for _, path := range req.GetUpdateMask().GetPaths() {
		switch path {
		case "display_name", "displayName":
			obj.DisplayName = req.GetServiceLevelObjective().GetDisplayName()
		case "rolling_period", "rollingPeriod":
			obj.Period = req.GetServiceLevelObjective().GetPeriod()
		case "goal":
			obj.Goal = req.ServiceLevelObjective.GetGoal()
		case "userLabels":
			obj.UserLabels = req.ServiceLevelObjective.GetUserLabels()
		case "serviceLevelIndicator.windowsBased.windowPeriod":
			if obj.ServiceLevelIndicator == nil {
				obj.ServiceLevelIndicator = &pb.ServiceLevelIndicator{}
			}
			if obj.ServiceLevelIndicator.GetWindowsBased() == nil {
				obj.ServiceLevelIndicator.Type = &pb.ServiceLevelIndicator_WindowsBased{}
			}
			obj.ServiceLevelIndicator.GetWindowsBased().WindowPeriod = req.GetServiceLevelObjective().GetServiceLevelIndicator().GetWindowsBased().GetWindowPeriod()

		case "serviceLevelIndicator.windowsBased.metricSumInRange.range.max":
			if obj.ServiceLevelIndicator == nil {
				obj.ServiceLevelIndicator = &pb.ServiceLevelIndicator{}
			}
			if obj.ServiceLevelIndicator.GetWindowsBased() == nil {
				obj.ServiceLevelIndicator.Type = &pb.ServiceLevelIndicator_WindowsBased{}
			}
			if obj.ServiceLevelIndicator.GetWindowsBased().GetMetricSumInRange() == nil {
				obj.ServiceLevelIndicator.GetWindowsBased().WindowCriterion = &pb.WindowsBasedSli_MetricSumInRange{}
			}
			if obj.ServiceLevelIndicator.GetWindowsBased().GetMetricSumInRange().GetRange() == nil {
				obj.ServiceLevelIndicator.GetWindowsBased().GetMetricSumInRange().Range = &pb.Range{}
			}
			obj.ServiceLevelIndicator.GetWindowsBased().GetMetricSumInRange().GetRange().Max = req.GetServiceLevelObjective().ServiceLevelIndicator.GetWindowsBased().GetMetricSumInRange().GetRange().GetMax()
		case "serviceLevelIndicator.windowsBased.metricSumInRange.range.min":
			if obj.ServiceLevelIndicator == nil {
				obj.ServiceLevelIndicator = &pb.ServiceLevelIndicator{}
			}
			if obj.ServiceLevelIndicator.GetWindowsBased() == nil {
				obj.ServiceLevelIndicator.Type = &pb.ServiceLevelIndicator_WindowsBased{}
			}
			if obj.ServiceLevelIndicator.GetWindowsBased().GetMetricSumInRange() == nil {
				obj.ServiceLevelIndicator.GetWindowsBased().WindowCriterion = &pb.WindowsBasedSli_MetricSumInRange{}
			}
			if obj.ServiceLevelIndicator.GetWindowsBased().GetMetricSumInRange().GetRange() == nil {
				obj.ServiceLevelIndicator.GetWindowsBased().GetMetricSumInRange().Range = &pb.Range{}
			}
			obj.ServiceLevelIndicator.GetWindowsBased().GetMetricSumInRange().GetRange().Min = req.GetServiceLevelObjective().ServiceLevelIndicator.GetWindowsBased().GetMetricSumInRange().GetRange().GetMin()

		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not supported by mock (full update_mask=%v)", path, req.GetUpdateMask())
		}
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *serviceMonitoringService) DeleteServiceLevelObjective(ctx context.Context, req *pb.DeleteServiceLevelObjectiveRequest) (*emptypb.Empty, error) {
	name, err := s.parseServiceLevelObjectiveName(req.Name)
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	deleted := &pb.ServiceLevelObjective{}

	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

type serviceLevelObjectiveName struct {
	Project   *projects.ProjectData
	serviceID string
	sloID     string
}

func (n *serviceLevelObjectiveName) String() string {
	return fmt.Sprintf("projects/%d/services/%s/serviceLevelObjectives/%s", n.Project.Number, n.serviceID, n.sloID)
}

func (s *MockService) parseServiceLevelObjectiveName(name string) (*serviceLevelObjectiveName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "services" && tokens[4] == "serviceLevelObjectives" {
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}

		n := &serviceLevelObjectiveName{
			Project:   project,
			serviceID: tokens[3],
			sloID:     tokens[5],
		}
		return n, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "invalid service name %q", name)
}
