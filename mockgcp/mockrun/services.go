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

package mockrun

import (
	"context"
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/fields"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/run/v2"
	"github.com/golang/protobuf/ptypes/duration"
	"github.com/google/uuid"
	api "google.golang.org/genproto/googleapis/api"
	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ServicesV2 struct {
	*MockService
	pb.UnimplementedServicesServer
}

func (s *ServicesV2) GetService(ctx context.Context, req *pb.GetServiceRequest) (*pb.Service, error) {
	name, err := s.parseServiceName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Service{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *ServicesV2) CreateService(ctx context.Context, req *pb.CreateServiceRequest) (*longrunning.Operation, error) {
	reqName := req.Parent + "/services/" + req.ServiceId
	name, err := s.parseServiceName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.Service).(*pb.Service)
	obj.Name = fqn
	obj.CreateTime = timestamppb.Now()
	obj.UpdateTime = timestamppb.Now()
	obj.Etag = fields.ComputeWeakEtag(obj)
	obj.Creator = "test@google.com"

	obj.LastModifier = "test@google.com"
	obj.Generation = 1

	if obj.Template == nil {
		obj.Template = &pb.RevisionTemplate{}
	}
	if obj.Template.Containers == nil {
		// Just a dummy default
		obj.Template.Containers = []*pb.Container{
			{
				Image: "gcr.io/google-samples/hello-app:1.0",
			},
		}
	}
	for _, container := range obj.Template.Containers {
		if container.Resources == nil {
			container.Resources = &pb.ResourceRequirements{
				Limits: map[string]string{
					"cpu":    "1000m",
					"memory": "512Mi",
				},
			}
		}
	}
	if obj.Template.Timeout == nil {
		obj.Template.Timeout = &duration.Duration{Seconds: 300}
	}
	if obj.Template.ServiceAccount == "" {
		obj.Template.ServiceAccount = fmt.Sprintf("%d-compute@developer.gserviceaccount.com", name.Project.Number)
	}
	if obj.Template.ExecutionEnvironment == 0 {
		obj.Template.ExecutionEnvironment = pb.ExecutionEnvironment_EXECUTION_ENVIRONMENT_GEN2
	}
	if obj.Template.MaxInstanceRequestConcurrency == 0 {
		obj.Template.MaxInstanceRequestConcurrency = 80
	}

	if obj.TerminalCondition == nil {
		obj.TerminalCondition = &pb.Condition{
			LastTransitionTime: timestamppb.Now(),
			State:              pb.Condition_CONDITION_SUCCEEDED,
			Type:               "Ready",
		}
	}
	obj.Uid = uuid.NewString()

	// Server-side defaults
	if obj.LaunchStage == 0 {
		obj.LaunchStage = api.LaunchStage_GA
	}
	if obj.Ingress == 0 {
		obj.Ingress = pb.IngressTraffic_INGRESS_TRAFFIC_ALL
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}
	return s.operations.StartLRO(ctx, req.Parent, obj, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *ServicesV2) UpdateService(ctx context.Context, req *pb.UpdateServiceRequest) (*longrunning.Operation, error) {
	name, err := s.parseServiceName(req.GetService().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Service{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// Update logic from request
	updatedService := req.GetService()
	if updatedService.Template != nil {
		obj.Template = updatedService.Template
	}
	if updatedService.Labels != nil {
		obj.Labels = updatedService.Labels
	}
	if updatedService.Annotations != nil {
		obj.Annotations = updatedService.Annotations
	}
	if updatedService.Description != "" {
		obj.Description = updatedService.Description
	}
	if updatedService.BinaryAuthorization != nil {
		obj.BinaryAuthorization = updatedService.BinaryAuthorization
	}
	if updatedService.Client != "" {
		obj.Client = updatedService.Client
	}
	if updatedService.ClientVersion != "" {
		obj.ClientVersion = updatedService.ClientVersion
	}
	if updatedService.Ingress != 0 {
		obj.Ingress = updatedService.Ingress
	}
	if updatedService.LaunchStage != 0 {
		obj.LaunchStage = updatedService.LaunchStage
	}
	if updatedService.Traffic != nil {
		obj.Traffic = updatedService.Traffic
	}
	if updatedService.Scaling != nil {
		obj.Scaling = updatedService.Scaling
	}

	obj.UpdateTime = timestamppb.Now()

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, lroPrefix, obj, func() (protoreflect.ProtoMessage, error) {
		return obj, nil
	})
}

func (s *ServicesV2) DeleteService(ctx context.Context, req *pb.DeleteServiceRequest) (*longrunning.Operation, error) {
	name, err := s.parseServiceName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Service{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}
	if err := s.storage.Delete(ctx, fqn, &pb.Service{}); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, lroPrefix, obj, func() (protoreflect.ProtoMessage, error) {
		return obj, nil
	})
}

type serviceName struct {
	Project  *projects.ProjectData
	Location string
	Service  string
}

func (n *serviceName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/services/%s", n.Project.ID, n.Location, n.Service)
}

func (s *MockService) parseServiceName(name string) (*serviceName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "services" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		n := &serviceName{
			Project:  project,
			Location: tokens[3],
			Service:  tokens[5],
		}
		return n, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
