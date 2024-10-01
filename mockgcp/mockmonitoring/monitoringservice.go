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

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/monitoring/v3"
	"github.com/golang/protobuf/ptypes/empty"
)

type serviceMonitoringService struct {
	*MockService
	pb.UnimplementedServiceMonitoringServiceServer
}

func (s *serviceMonitoringService) GetService(ctx context.Context, req *pb.GetServiceRequest) (*pb.Service, error) {
	name, err := s.parseServiceName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Service{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "There is no service with id '%s' in project '%d'", name.ServiceID, name.Project.Number)
		}
		return nil, err
	}

	return obj, nil
}

func (s *serviceMonitoringService) CreateService(ctx context.Context, req *pb.CreateServiceRequest) (*pb.Service, error) {
	now := time.Now()

	serviceID := req.GetServiceId()
	if serviceID == "" {
		serviceID = fmt.Sprintf("%x", now.UnixNano())
	}

	reqName := req.GetParent() + "/services/" + serviceID
	name, err := s.parseServiceName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := ProtoClone(req.Service)
	obj.Name = fqn

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *serviceMonitoringService) UpdateService(ctx context.Context, req *pb.UpdateServiceRequest) (*pb.Service, error) {
	name, err := s.parseServiceName(req.GetService().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	existing := &pb.Service{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		return nil, err
	}

	updated := ProtoClone(existing)
	for _, path := range req.GetUpdateMask().GetPaths() {
		switch path {
		case "displayName":
			updated.DisplayName = req.GetService().GetDisplayName()
		case "telemetry.resourceName":
			if updated.Telemetry == nil {
				updated.Telemetry = &pb.Service_Telemetry{}
			}
			updated.Telemetry.ResourceName = req.GetService().GetTelemetry().GetResourceName()
		case "userLabels":
			updated.UserLabels = req.GetService().GetUserLabels()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not supported by mock (full update_mask=%v)", path, req.GetUpdateMask())
		}
	}

	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}

	return updated, nil
}

func (s *serviceMonitoringService) DeleteService(ctx context.Context, req *pb.DeleteServiceRequest) (*empty.Empty, error) {
	name, err := s.parseServiceName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.Service{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}

type ServiceName struct {
	Project   *projects.ProjectData
	ServiceID string
}

func (n *ServiceName) String() string {
	return fmt.Sprintf("projects/%d/services/%s", n.Project.Number, n.ServiceID)
}

// parseServiceName parses a string into a ServiceName.
// The format is:
//
//	projects/[PROJECT_ID_OR_NUMBER]/services/[SERVICE_ID]
func (s *MockService) parseServiceName(name string) (*ServiceName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "services" {
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &ServiceName{
			Project:   project,
			ServiceID: tokens[3],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
