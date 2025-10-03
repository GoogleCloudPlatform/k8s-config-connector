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

package mockmonitoring

import (
	"context"
	"fmt"
	"strings"
	"time"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/monitoring/metricsscope/v1"
)

type metricsScopeService struct {
	*MockService
	pb.UnimplementedMetricsScopesServer
}

func (s *metricsScopeService) getMetricsScope(ctx context.Context, name *MetricsScopeName, createIfNotExists bool) (*pb.MetricsScope, error) {
	fqn := name.String()

	metricsScope := &pb.MetricsScope{}
	if err := s.storage.Get(ctx, fqn, metricsScope); err != nil {
		if status.Code(err) == codes.NotFound {
			metricsScope = nil
		} else {
			return nil, err
		}
	}

	if metricsScope == nil && createIfNotExists {
		metricsScope = createDefaultMetricsScopeForProject(ctx, name.Project)

		if err := s.storage.Create(ctx, fqn, metricsScope); err != nil {
			return nil, err
		}
	}

	return metricsScope, nil
}

func (s *metricsScopeService) GetMetricsScope(ctx context.Context, req *pb.GetMetricsScopeRequest) (*pb.MetricsScope, error) {
	name, err := s.parseMetricsScopeName(req.GetName())
	if err != nil {
		return nil, err
	}

	// Because of the timestamps in the metricsScope, we do create if it does not exist (and write it to storage)
	createIfNotExists := true
	obj, err := s.getMetricsScope(ctx, name, createIfNotExists)
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func createDefaultMetricsScopeForProject(ctx context.Context, project *projects.ProjectData) *pb.MetricsScope {
	now := time.Now()

	obj := &pb.MetricsScope{}

	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)

	obj.Name = fmt.Sprintf("locations/global/metricsScopes/%d", project.Number)

	obj.MonitoredProjects = append(obj.MonitoredProjects, &pb.MonitoredProject{
		CreateTime: timestamppb.New(now),
		Name:       fmt.Sprintf("locations/global/metricsScopes/%d/projects/%d", project.Number, project.Number),
	})

	return obj
}

// Adds a `MonitoredProject` with the given project ID
// to the specified `Metrics Scope`.
func (s *metricsScopeService) CreateMonitoredProject(ctx context.Context, req *pb.CreateMonitoredProjectRequest) (*longrunningpb.Operation, error) {
	now := time.Now()

	metricsScopeName, err := s.parseMetricsScopeName(req.GetParent())
	if err != nil {
		return nil, err
	}

	createIfNotExists := true
	metricsScope, err := s.getMetricsScope(ctx, metricsScopeName, createIfNotExists)
	if err != nil {
		return nil, err
	}

	// MonitoredProject is Required. The initial `MonitoredProject` configuration.
	// Specify only the `monitored_project.name` field. All other fields are
	// ignored. The `monitored_project.name` must be in the format:
	// `locations/global/metricsScopes/{SCOPING_PROJECT_ID_OR_NUMBER}/projects/{MONITORED_PROJECT_ID_OR_NUMBER}`

	if req.GetMonitoredProject().GetName() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "monitored_project.name is required")
	}
	monitoredProjectName, err := s.parseMonitoredProjectName(req.GetMonitoredProject().GetName())
	if err != nil {
		return nil, err
	}

	metricsScope.MonitoredProjects = append(metricsScope.MonitoredProjects, &pb.MonitoredProject{
		CreateTime: timestamppb.New(now),
		Name:       monitoredProjectName.String(),
	})

	if err := s.storage.Update(ctx, metricsScopeName.String(), metricsScope); err != nil {
		return nil, err
	}

	metadata := &pb.OperationMetadata{
		// "@type": "type.googleapis.com/google.monitoring.metricsscope.v1.OperationMetadata",
		State:      pb.OperationMetadata_DONE,
		CreateTime: timestamppb.New(now),
		UpdateTime: timestamppb.New(now),
	}

	// Note: returns a MonitoredProject, not a MetricsScope
	response := &pb.MonitoredProject{
		CreateTime: timestamppb.New(now),
		Name:       monitoredProjectName.String(),
	}
	return s.operations.DoneLRO(ctx, "", metadata, response)
}

func (s *metricsScopeService) DeleteMonitoredProject(ctx context.Context, req *pb.DeleteMonitoredProjectRequest) (*longrunningpb.Operation, error) {
	removeMonitoredProjectName, err := s.parseMonitoredProjectName(req.Name)
	if err != nil {
		return nil, err
	}

	metricsScopeName := removeMonitoredProjectName.Parent()

	createIfNotExists := true
	metricsScope, err := s.getMetricsScope(ctx, metricsScopeName, createIfNotExists)
	if err != nil {
		return nil, err
	}

	now := time.Now()

	removed := false
	var keep []*pb.MonitoredProject
	for _, monitoredProject := range metricsScope.MonitoredProjects {
		if monitoredProject.Name == removeMonitoredProjectName.String() {
			removed = true
			continue
		}
		keep = append(keep, monitoredProject)
	}
	metricsScope.MonitoredProjects = keep

	if !removed {
		return nil, status.Errorf(codes.NotFound, "monitored project not found")
	}

	if err := s.storage.Update(ctx, metricsScopeName.String(), metricsScope); err != nil {
		return nil, err
	}

	metadata := &pb.OperationMetadata{
		// "@type": "type.googleapis.com/google.monitoring.metricsscope.v1.OperationMetadata",
		State:      pb.OperationMetadata_DONE,
		CreateTime: timestamppb.New(now),
		UpdateTime: timestamppb.New(now),
	}

	// Note: returns an Empty pb, not a MetricsScope
	response := &emptypb.Empty{}
	return s.operations.DoneLRO(ctx, "", metadata, response)

}

type MetricsScopeName struct {
	Project *projects.ProjectData
}

func (n *MetricsScopeName) String() string {
	return fmt.Sprintf("locations/global/metricsScopes/%d", n.Project.Number)
}

// parseMetricsScopeName parses a string into a MetricsScopeName.
// The expected form is `locations/global/metricsScopes/{SCOPING_PROJECT_ID_OR_NUMBER}`
func (s *MockService) parseMetricsScopeName(name string) (*MetricsScopeName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 4 && tokens[0] == "locations" && tokens[1] == "global" && tokens[2] == "metricsScopes" {
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[3])
		if err != nil {
			return nil, err
		}

		name := &MetricsScopeName{
			Project: project,
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}

type monitoredProjectName struct {
	ParentProject    *projects.ProjectData
	MonitoredProject *projects.ProjectData
}

func (n *monitoredProjectName) String() string {
	return fmt.Sprintf("locations/global/metricsScopes/%d/projects/%d", n.ParentProject.Number, n.MonitoredProject.Number)
}

func (n *monitoredProjectName) Parent() *MetricsScopeName {
	return &MetricsScopeName{
		Project: n.ParentProject,
	}
}

// parseMonitoredProjectName parses a string into a monitoredProjectName.
// The expected form is `locations/global/metricsScopes/{SCOPING_PROJECT_ID_OR_NUMBER}/projects/{MONITORED_PROJECT_ID_OR_NUMBER}`
func (s *MockService) parseMonitoredProjectName(name string) (*monitoredProjectName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "locations" && tokens[1] == "global" && tokens[2] == "metricsScopes" && tokens[4] == "projects" {
		parentProject, err := s.Projects.GetProjectByIDOrNumber(tokens[3])
		if err != nil {
			return nil, err
		}

		monitoredProject, err := s.Projects.GetProjectByIDOrNumber(tokens[5])
		if err != nil {
			return nil, err
		}

		name := &monitoredProjectName{
			ParentProject:    parentProject,
			MonitoredProject: monitoredProject,
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
