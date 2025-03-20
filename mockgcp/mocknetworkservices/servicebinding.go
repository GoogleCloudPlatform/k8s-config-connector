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
// proto.service: google.cloud.networkservices.v1.NetworkServices
// proto.message: google.cloud.networkservices.v1.ServiceBinding

package mocknetworkservices

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/networkservices/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

func (s *NetworkServicesServer) GetServiceBinding(ctx context.Context, req *pb.GetServiceBindingRequest) (*pb.ServiceBinding, error) {
	name, err := s.parseServiceBindingName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.ServiceBinding{}
	obj.Name = fqn
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The ServiceBinding does not exist.")
		}
		return nil, err
	}
	now := time.Now()
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)

	return obj, nil
}

func (s *NetworkServicesServer) ListServiceBindings(ctx context.Context, req *pb.ListServiceBindingsRequest) (*pb.ListServiceBindingsResponse, error) {
	response := &pb.ListServiceBindingsResponse{}

	findKind := (&pb.ServiceBinding{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, findKind, storage.ListOptions{
		Prefix: req.Parent + "/serviceBindings/",
	}, func(obj proto.Message) error {
		serviceBinding := obj.(*pb.ServiceBinding)
		response.ServiceBindings = append(response.ServiceBindings, serviceBinding)
		return nil
	}); err != nil {
		return nil, err
	}
	return response, nil
}

func (s *NetworkServicesServer) CreateServiceBinding(ctx context.Context, req *pb.CreateServiceBindingRequest) (*longrunningpb.Operation, error) {
	reqName := req.Parent + "/serviceBindings/" + req.ServiceBindingId
	name, err := s.parseServiceBindingName(reqName)
	if err != nil {
		return nil, err
	}
	uniqueId := strings.Split(name.ServiceBindingName, "-")[2]

	fqn := name.String()

	now := time.Now()

	obj := proto.Clone(req.ServiceBinding).(*pb.ServiceBinding)
	obj.Name = fqn
	obj.Service = fmt.Sprintf("projects/mock-project/locations/us-central1/namespaces/namespace-%s/services/service-%s", uniqueId, uniqueId)
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		EndTime:    timestamppb.New(now),
		Target:     name.String(),
		Verb:       "create",
		ApiVersion: "v1",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		result := proto.Clone(obj).(*pb.ServiceBinding)
		result.CreateTime = timestamppb.New(now)
		result.UpdateTime = timestamppb.New(now)
		result.Name = reqName
		return result, nil
	})
}

func (s *NetworkServicesServer) DeleteServiceBinding(ctx context.Context, req *pb.DeleteServiceBindingRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseServiceBindingName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.ServiceBinding{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	now := time.Now()
	lroMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		EndTime:    timestamppb.New(now),
		Target:     name.String(),
		Verb:       "delete",
		ApiVersion: "v1",
	}
	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.DoneLRO(ctx, lroPrefix, lroMetadata, &emptypb.Empty{})
}

type serviceBindingName struct {
	Project            *projects.ProjectData
	Location           string
	ServiceBindingName string
}

func (n *serviceBindingName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/serviceBindings/" + n.ServiceBindingName
}

// parseServiceBindingName parses a string into an serviceBindingName.
// The expected form is `projects/*/locations/global/serviceBindings/*`.
func (s *NetworkServicesServer) parseServiceBindingName(name string) (*serviceBindingName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "serviceBindings" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &serviceBindingName{
			Project:            project,
			Location:           tokens[3],
			ServiceBindingName: tokens[5],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
