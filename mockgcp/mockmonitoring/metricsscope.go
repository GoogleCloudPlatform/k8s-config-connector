// // Copyright 2024 Google LLC
// //
// // Licensed under the Apache License, Version 2.0 (the "License");
// // you may not use this file except in compliance with the License.
// // You may obtain a copy of the License at
// //
// //      http://www.apache.org/licenses/LICENSE-2.0
// //
// // Unless required by applicable law or agreed to in writing, software
// // distributed under the License is distributed on an "AS IS" BASIS,
// // WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// // See the License for the specific language governing permissions and
// // limitations under the License.

// // +tool:mockgcp-support
// // proto.service: google.monitoring.metricsscope.v1.MetricsScopes
// // proto.message: google.monitoring.metricsscope.v1.MetricsScope

package mockmonitoring

// import (
// 	"context"
// 	"fmt"
// 	"strings"
// 	"time"

// 	"google.golang.org/grpc/codes"
// 	"google.golang.org/grpc/status"
// 	"google.golang.org/protobuf/proto"
// 	"google.golang.org/protobuf/types/known/emptypb"
// 	"google.golang.org/protobuf/types/known/timestamppb"

// 	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
// 	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/monitoring/metricsscope/v1"
// 	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
// 	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
// )

// func (s *metricsScopesV1) GetMetricsScope(ctx context.Context, req *pb.GetMetricsScopeRequest) (*pb.MetricsScope, error) {
// 	name, err := s.parseMetricsScopeName(req.GetName())
// 	if err != nil {
// 		return nil, err
// 	}

// 	fqn := name.String()

// 	obj := &pb.MetricsScope{}
// 	if err := s.storage.Get(ctx, fqn, obj); err != nil {
// 		if status.Code(err) == codes.NotFound {
// 			// MetricsScope is implicitly available for projects
// 			obj = s.createDefaultMetricsScope(name)
// 			if err := s.storage.Create(ctx, fqn, obj); err != nil {
// 				return nil, err
// 			}
// 		} else {
// 			return nil, err
// 		}
// 	}

// 	return obj, nil
// }

// func (s *metricsScopesV1) createDefaultMetricsScope(name *metricsScopeName) *pb.MetricsScope {
// 	now := time.Now()
// 	return &pb.MetricsScope{
// 		Name:       name.String(),
// 		CreateTime: timestamppb.New(now),
// 		UpdateTime: timestamppb.New(now),
// 		MonitoredProjects: []*pb.MonitoredProject{
// 			{
// 				Name:       fmt.Sprintf("locations/global/metricsScopes/%d/projects/%d", name.Project.Number, name.Project.Number),
// 				CreateTime: timestamppb.New(now),
// 			},
// 		},
// 	}
// }

// func (s *metricsScopesV1) CreateMonitoredProject(ctx context.Context, req *pb.CreateMonitoredProjectRequest) (*longrunningpb.Operation, error) {
// 	name, err := s.parseMetricsScopeName(req.GetParent())
// 	if err != nil {
// 		return nil, err
// 	}

// 	fqn := name.String()
// 	metricsScope := &pb.MetricsScope{}
// 	if err := s.storage.Get(ctx, fqn, metricsScope); err != nil {
// 		if status.Code(err) == codes.NotFound {
// 			metricsScope = s.createDefaultMetricsScope(name)
// 			if err := s.storage.Create(ctx, fqn, metricsScope); err != nil {
// 				return nil, err
// 			}
// 		} else {
// 			return nil, err
// 		}
// 	}

// 	monitoredProjectName, err := s.parseMonitoredProjectName(req.GetMonitoredProject().GetName())
// 	if err != nil {
// 		return nil, err
// 	}
// 	monitoredProjectFQN := monitoredProjectName.String()

// 	// Check if already exists
// 	for _, mp := range metricsScope.MonitoredProjects {
// 		if mp.Name == monitoredProjectFQN {
// 			return nil, status.Errorf(codes.AlreadyExists, "Monitored project %s already exists in metrics scope %s", monitoredProjectFQN, name.String())
// 		}
// 	}

// 	now := time.Now()
// 	newMP := &pb.MonitoredProject{
// 		Name:       monitoredProjectFQN,
// 		CreateTime: timestamppb.New(now),
// 	}
// 	metricsScope.MonitoredProjects = append(metricsScope.MonitoredProjects, newMP)
// 	metricsScope.UpdateTime = timestamppb.New(now)

// 	if err := s.storage.Update(ctx, fqn, metricsScope); err != nil {
// 		return nil, err
// 	}

// 	metadata := &pb.OperationMetadata{
// 		State:      pb.OperationMetadata_DONE,
// 		CreateTime: timestamppb.New(now),
// 		UpdateTime: timestamppb.New(now),
// 	}
// 	return s.operations.DoneLRO(ctx, "", metadata, newMP)
// }

// func (s *metricsScopesV1) DeleteMonitoredProject(ctx context.Context, req *pb.DeleteMonitoredProjectRequest) (*longrunningpb.Operation, error) {
// 	mpName, err := s.parseMonitoredProjectName(req.GetName())
// 	if err != nil {
// 		return nil, err
// 	}
// 	metricsScopeName := mpName.MetricsScopeName()
// 	fqn := metricsScopeName.String()

// 	metricsScope := &pb.MetricsScope{}
// 	if err := s.storage.Get(ctx, fqn, metricsScope); err != nil {
// 		return nil, err
// 	}

// 	var filtered []*pb.MonitoredProject
// 	found := false
// 	for _, mp := range metricsScope.MonitoredProjects {
// 		if mp.Name == req.GetName() {
// 			found = true
// 			continue
// 		}
// 		filtered = append(filtered, mp)
// 	}

// 	if !found {
// 		return nil, status.Errorf(codes.NotFound, "Monitored project %s not found in metrics scope", req.GetName())
// 	}

// 	metricsScope.MonitoredProjects = filtered
// 	metricsScope.UpdateTime = timestamppb.Now()

// 	if err := s.storage.Update(ctx, fqn, metricsScope); err != nil {
// 		return nil, err
// 	}

// 	metadata := &pb.OperationMetadata{
// 		State:      pb.OperationMetadata_DONE,
// 		CreateTime: timestamppb.Now(),
// 		UpdateTime: timestamppb.Now(),
// 	}
// 	return s.operations.DoneLRO(ctx, "", metadata, &emptypb.Empty{})
// }

// func (s *metricsScopesV1) ListMetricsScopesByMonitoredProject(ctx context.Context, req *pb.ListMetricsScopesByMonitoredProjectRequest) (*pb.ListMetricsScopesByMonitoredProjectResponse, error) {
// 	response := &pb.ListMetricsScopesByMonitoredProjectResponse{}

// 	scheduleKind := (&pb.MetricsScope{}).ProtoReflect().Descriptor()
// 	if err := s.storage.List(ctx, scheduleKind, storage.ListOptions{}, func(obj proto.Message) error {
// 		metricsScope := obj.(*pb.MetricsScope)

// 		// TODO: Filter

// 		response.MetricsScopes = append(response.MetricsScopes, metricsScope)
// 		return nil
// 	}); err != nil {
// 		return nil, err
// 	}
// 	return response, nil
// }

// type metricsScopeName struct {
// 	Project *projects.ProjectData
// }

// func (n *metricsScopeName) String() string {
// 	return fmt.Sprintf("locations/global/metricsScopes/%d", n.Project.Number)
// }

// func (s *MockService) parseMetricsScopeName(name string) (*metricsScopeName, error) {
// 	tokens := strings.Split(name, "/")
// 	// locations/global/metricsScopes/{SCOPING_PROJECT_ID_OR_NUMBER}
// 	if len(tokens) == 4 && tokens[0] == "locations" && tokens[1] == "global" && tokens[2] == "metricsScopes" {
// 		project, err := s.Projects.GetProjectByIDOrNumber(tokens[3])
// 		if err != nil {
// 			return nil, err
// 		}
// 		return &metricsScopeName{Project: project}, nil
// 	}
// 	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
// }

// type monitoredProjectName struct {
// 	ScopeProject     *projects.ProjectData
// 	MonitoredProject *projects.ProjectData
// }

// func (n *monitoredProjectName) String() string {
// 	return fmt.Sprintf("locations/global/metricsScopes/%d/projects/%d", n.ScopeProject.Number, n.MonitoredProject.Number)
// }

// func (n *monitoredProjectName) MetricsScopeName() *metricsScopeName {
// 	return &metricsScopeName{Project: n.ScopeProject}
// }

// func (s *MockService) parseMonitoredProjectName(name string) (*monitoredProjectName, error) {
// 	tokens := strings.Split(name, "/")
// 	// locations/global/metricsScopes/{SCOPING_PROJECT_ID_OR_NUMBER}/projects/{MONITORED_PROJECT_ID_OR_NUMBER}
// 	if len(tokens) == 6 && tokens[0] == "locations" && tokens[1] == "global" && tokens[2] == "metricsScopes" && tokens[4] == "projects" {
// 		scopeProject, err := s.Projects.GetProjectByIDOrNumber(tokens[3])
// 		if err != nil {
// 			return nil, err
// 		}
// 		monitoredProject, err := s.Projects.GetProjectByIDOrNumber(tokens[5])
// 		if err != nil {
// 			return nil, err
// 		}
// 		return &monitoredProjectName{
// 			ScopeProject:     scopeProject,
// 			MonitoredProject: monitoredProject,
// 		}, nil
// 	}
// 	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
// }
