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

package mocknetworkconnectivity

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"k8s.io/klog/v2"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/networkconnectivity/v1"
)

type serviceConnectionMapsServer struct {
	*MockService
	pb.UnimplementedProjectsLocationsServiceConnectionMapsServerServer
}

func (s *serviceConnectionMapsServer) GetProjectsLocationsServiceConnectionMap(ctx context.Context, req *pb.GetProjectsLocationsServiceConnectionMapRequest) (*pb.ServiceConnectionMap, error) {
	name, err := s.parseServiceConnectionMapName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.ServiceConnectionMap{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *serviceConnectionMapsServer) CreateProjectsLocationsServiceConnectionMap(ctx context.Context, req *pb.CreateProjectsLocationsServiceConnectionMapRequest) (*longrunning.Operation, error) {
	reqName := fmt.Sprintf("%s/serviceConnectionMaps/%s", req.GetParent(), req.GetServiceConnectionMapId())
	name, err := s.parseServiceConnectionMapName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	obj := proto.CloneOf(req.GetProjectsLocationsServiceConnectionMap())
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)

	s.populateDefaultsForServiceConnectionMap(name, obj)

	obj.Etag = computeEtag(obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	metadata := &pb.OperationMetadata{
		ApiVersion:            "v1",
		RequestedCancellation: false,
		CreateTime:            timestamppb.New(now),
		Target:                fqn,
		Verb:                  "create",
	}
	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()

		if err := s.storage.Update(ctx, fqn, obj); err != nil {
			return nil, err
		}

		return obj, nil
	})
}

func (s *serviceConnectionMapsServer) PatchProjectsLocationsServiceConnectionMap(ctx context.Context, req *pb.PatchProjectsLocationsServiceConnectionMapRequest) (*longrunning.Operation, error) {
	log := klog.FromContext(ctx)

	reqName := req.GetName()

	name, err := s.parseServiceConnectionMapName(reqName)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	now := time.Now()

	obj := &pb.ServiceConnectionMap{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	obj.UpdateTime = timestamppb.New(now)

	patch := req.GetProjectsLocationsServiceConnectionMap()
	if req.GetUpdateMask() != "" {
		paths := strings.Split(req.GetUpdateMask(), ",")
		for _, path := range paths {
			switch path {
			case "description":
				obj.Description = patch.GetDescription()
			case "labels":
				obj.Labels = patch.GetLabels()
			case "producer_psc_configs", "producerPscConfigs":
				obj.ProducerPscConfigs = patch.GetProducerPscConfigs()
			case "consumer_psc_configs", "consumerPscConfigs":
				obj.ConsumerPscConfigs = patch.GetConsumerPscConfigs()
			case "token":
				obj.Token = patch.GetToken()
			default:
				log.Info("unsupported update_mask", "req", req)
				return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not supported by mock", path)
			}
		}
	} else {
		if patch.Description != "" {
			obj.Description = patch.Description
		}
		if patch.Labels != nil {
			obj.Labels = patch.Labels
		}
		if patch.ProducerPscConfigs != nil {
			obj.ProducerPscConfigs = patch.ProducerPscConfigs
		}
		if patch.ConsumerPscConfigs != nil {
			obj.ConsumerPscConfigs = patch.ConsumerPscConfigs
		}
		if patch.Token != "" {
			obj.Token = patch.Token
		}
	}

	obj.Etag = computeEtag(obj)

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	metadata := &pb.OperationMetadata{
		ApiVersion:            "v1",
		RequestedCancellation: false,
		CreateTime:            timestamppb.New(now),
		Target:                fqn,
		Verb:                  "update",
	}
	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}

func (s *serviceConnectionMapsServer) DeleteProjectsLocationsServiceConnectionMap(ctx context.Context, req *pb.DeleteProjectsLocationsServiceConnectionMapRequest) (*longrunning.Operation, error) {
	name, err := s.parseServiceConnectionMapName(req.GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	now := time.Now()

	oldObj := &pb.ServiceConnectionMap{}
	if err := s.storage.Delete(ctx, fqn, oldObj); err != nil {
		return nil, err
	}

	metadata := &pb.OperationMetadata{
		ApiVersion:            "v1",
		RequestedCancellation: false,
		CreateTime:            timestamppb.New(now),
		Target:                fqn,
		Verb:                  "delete",
	}
	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()
		return &emptypb.Empty{}, nil
	})
}

func (s *serviceConnectionMapsServer) ListProjectsLocationsServiceConnectionMaps(ctx context.Context, req *pb.ListProjectsLocationsServiceConnectionMapsRequest) (*pb.ListServiceConnectionMapsResponse, error) {
	response := &pb.ListServiceConnectionMapsResponse{}
	return response, nil
}

func (s *MockService) populateDefaultsForServiceConnectionMap(name *serviceConnectionMapName, obj *pb.ServiceConnectionMap) {
	if obj.Infrastructure == "" {
		obj.Infrastructure = "PSC"
	}
}

type serviceConnectionMapName struct {
	Project                *projects.ProjectData
	Location               string
	ServiceConnectionMapID string
}

func (n *serviceConnectionMapName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/serviceConnectionMaps/" + n.ServiceConnectionMapID
}

func (s *MockService) parseServiceConnectionMapName(name string) (*serviceConnectionMapName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "serviceConnectionMaps" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &serviceConnectionMapName{
			Project:                project,
			Location:               tokens[3],
			ServiceConnectionMapID: tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
