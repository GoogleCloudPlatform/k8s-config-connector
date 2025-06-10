// Copyright 2022 Google LLC
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

package mockserviceusage

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"k8s.io/klog/v2"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/api/serviceusage/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

type ServiceUsageV1 struct {
	*MockService
	pb.UnimplementedServiceUsageServer
}

func (s *ServiceUsageV1) EnableService(ctx context.Context, req *pb.EnableServiceRequest) (*longrunning.Operation, error) {
	name, err := s.parseServiceName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	// Verify that this is a known service
	if !isKnownService(name.ServiceName) {
		klog.Errorf("enabling service %q not implemented in mock", name.ServiceName)
		return nil, status.Errorf(codes.PermissionDenied, "Not found or permission denied for service(s): %v", name.ServiceName)
	}

	create := false
	service := &pb.Service{}
	if err := s.storage.Get(ctx, fqn, service); err != nil {
		if status.Code(err) == codes.NotFound {
			create = true
		} else {
			return nil, err
		}
	}

	changed := false
	if service.GetState() != pb.State_ENABLED {
		changed = true
		if create {
			service = &pb.Service{
				Name:   fqn,
				Parent: fmt.Sprintf("projects/%d", name.Project.Number),
				State:  pb.State_ENABLED,
			}

			if err := s.storage.Create(ctx, fqn, service); err != nil {
				return nil, err
			}
		} else {
			service.State = pb.State_ENABLED
			if err := s.storage.Update(ctx, fqn, service); err != nil {
				return nil, err
			}
		}
	}

	lroPrefix := ""

	response := &pb.EnableServiceResponse{
		Service: service,
	}
	if !changed {
		return s.operations.DoneLRO(ctx, lroPrefix, nil, response)
	} else {
		return s.operations.StartLRO(ctx, lroPrefix, &emptypb.Empty{}, func() (proto.Message, error) {
			return response, nil
		})
	}
}

func (s *ServiceUsageV1) BatchEnableServices(ctx context.Context, req *pb.BatchEnableServicesRequest) (*longrunning.Operation, error) {
	name, err := s.parseServiceName(req.GetParent() + "/services/placeholder")
	if err != nil {
		return nil, err
	}

	metadata := &pb.OperationMetadata{}
	response := &pb.BatchEnableServicesResponse{}
	for _, serviceID := range req.GetServiceIds() {
		name.ServiceName = serviceID
		fqn := name.String()

		// Verify that this is a known service
		if !isKnownService(name.ServiceName) {
			return nil, status.Errorf(codes.PermissionDenied, "Not found or permission denied for service(s): %v", name.ServiceName)
		}

		exists := true
		service := &pb.Service{}
		if err := s.storage.Get(ctx, fqn, service); err != nil {
			if status.Code(err) == codes.NotFound {
				exists = false
			} else {
				return nil, err
			}
		}

		if service.GetState() != pb.State_ENABLED {
			if !exists {
				service = &pb.Service{
					Name:   fqn,
					Parent: fmt.Sprintf("projects/%d", name.Project.Number),
					State:  pb.State_ENABLED,
				}

				if err := s.storage.Create(ctx, fqn, service); err != nil {
					return nil, err
				}
			} else {
				service.State = pb.State_ENABLED
				if err := s.storage.Update(ctx, fqn, service); err != nil {
					return nil, err
				}
			}
		}

		metadata.ResourceNames = append(metadata.ResourceNames, fmt.Sprintf("services/%s/projectSettings/%d", name.ServiceName, name.Project.Number))
		response.Services = append(response.Services, service)
	}

	prefix := ""
	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		return response, nil
	})
}

func (s *ServiceUsageV1) DisableService(ctx context.Context, req *pb.DisableServiceRequest) (*longrunning.Operation, error) {
	name, err := s.parseServiceName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	// Verify that this is a known service
	if !isKnownService(name.ServiceName) {
		return nil, status.Errorf(codes.PermissionDenied, "Not found or permission denied for service(s): %v", name.ServiceName)
	}

	exists := true
	service := &pb.Service{}
	if err := s.storage.Get(ctx, fqn, service); err != nil {
		if status.Code(err) == codes.NotFound {
			exists = false
		} else {
			return nil, err
		}
	}

	if !exists {
		// Services are disabled by default
	} else {
		service.State = pb.State_DISABLED
		if err := s.storage.Update(ctx, fqn, service); err != nil {
			return nil, err
		}
	}

	prefix := ""
	metadata := &emptypb.Empty{}
	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		response := &pb.DisableServiceResponse{
			Service: service,
		}
		return response, nil
	})
}

func (s *ServiceUsageV1) GetService(ctx context.Context, req *pb.GetServiceRequest) (*pb.Service, error) {
	name, err := s.parseServiceName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Service{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			// Mock: everything is enabled
			obj = &pb.Service{
				Name:  name.String(),
				State: pb.State_DISABLED,
			}
		} else {
			return nil, err
		}
	}

	return obj, nil
}

func (s *ServiceUsageV1) ListServices(ctx context.Context, req *pb.ListServicesRequest) (*pb.ListServicesResponse, error) {
	shouldReturn := make(map[pb.State]bool)
	switch req.Filter {
	case "state:ENABLED":
		shouldReturn[pb.State_ENABLED] = true
	case "state:DISABLED":
		shouldReturn[pb.State_DISABLED] = true
	case "":
		shouldReturn[pb.State_DISABLED] = true
		shouldReturn[pb.State_ENABLED] = true
	default:
		return nil, status.Errorf(codes.InvalidArgument, "unexpected filter %q", req.Filter)
	}

	parent, err := projects.ParseProjectName(req.Parent)
	if err != nil {
		return nil, err
	}

	project, err := s.Projects.GetProject(parent)
	if err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%d/services/", project.Number)

	known := make(map[string]*pb.Service)

	serviceKind := (&pb.Service{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, serviceKind, storage.ListOptions{
		Prefix: prefix,
	}, func(obj proto.Message) error {
		s := obj.(*pb.Service)
		known[lastComponent(s.Name)] = s
		return nil
	}); err != nil {
		return nil, status.Errorf(codes.Internal, "error reading services: %v", err)
	}

	response := &pb.ListServicesResponse{}
	for _, name := range allServices {
		state := pb.State_DISABLED
		if s := known[name]; s != nil {
			state = s.State
		}
		if !shouldReturn[state] {
			continue
		}

		response.Services = append(response.Services, &pb.Service{
			Name:   prefix + name,
			Parent: fmt.Sprintf("projects/%d", project.Number),
			State:  state,
		})
	}

	return response, nil
}

func lastComponent(s string) string {
	i := strings.LastIndex(s, "/")
	return s[i+1:]
}
