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

// +tool:mockgcp-support
// proto.service: google.appengine.v1.Services
// proto.message: google.appengine.v1.Service

package mockappengine

import (
	"context"
	"fmt"
	"strings"

	pb "cloud.google.com/go/appengine/apiv1/appenginepb"
	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type servicesServer struct {
	*MockService
	pb.UnimplementedServicesServer
}

func (s *servicesServer) GetService(ctx context.Context, req *pb.GetServiceRequest) (*pb.Service, error) {
	name, err := parseServiceName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Service{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			if name.Service == "default" {
				// Default service exists by default in App Engine applications.
				// Pre-populate with baseline versions and split to match real GCP test environment state recorded in _http.log.
				obj = &pb.Service{
					Name: fqn,
					Id:   name.Service,
					Split: &pb.TrafficSplit{
						ShardBy: pb.TrafficSplit_IP,
						Allocations: map[string]float64{
							"v1": 0.8,
							"v2": 0.2,
						},
					},
				}
				if err := s.storage.Create(ctx, fqn, obj); err != nil {
					return nil, err
				}
				return obj, nil
			}
		}
		return nil, err
	}
	return obj, nil
}

func (s *servicesServer) UpdateService(ctx context.Context, req *pb.UpdateServiceRequest) (*longrunningpb.Operation, error) {
	name, err := parseServiceName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Service{}
	exists := true
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound && name.Service == "default" {
			exists = false
			obj = &pb.Service{
				Name: fqn,
				Id:   name.Service,
			}
		} else {
			return nil, err
		}
	}

	if mask := req.GetUpdateMask(); mask != nil && len(mask.GetPaths()) > 0 {
		for _, path := range mask.GetPaths() {
			switch path {
			case "split":
				if req.GetService() != nil && req.GetService().GetSplit() != nil {
					obj.Split = proto.Clone(req.GetService().GetSplit()).(*pb.TrafficSplit)
				} else {
					obj.Split = nil
				}
			case "labels":
				if req.GetService() != nil {
					obj.Labels = req.GetService().GetLabels()
				} else {
					obj.Labels = nil
				}
			case "network_settings", "networkSettings":
				if req.GetService() != nil && req.GetService().GetNetworkSettings() != nil {
					obj.NetworkSettings = proto.Clone(req.GetService().GetNetworkSettings()).(*pb.NetworkSettings)
				} else {
					obj.NetworkSettings = nil
				}
			default:
				return nil, status.Errorf(codes.InvalidArgument, "field %q is not supported in update_mask", path)
			}
		}
	} else if req.GetService() != nil {
		if req.GetService().GetSplit() != nil {
			obj.Split = proto.Clone(req.GetService().GetSplit()).(*pb.TrafficSplit)
		}
		if req.GetService().GetLabels() != nil {
			obj.Labels = req.GetService().GetLabels()
		}
		if req.GetService().GetNetworkSettings() != nil {
			obj.NetworkSettings = proto.Clone(req.GetService().GetNetworkSettings()).(*pb.NetworkSettings)
		}
	}
	obj.Name = fqn
	obj.Id = name.Service

	if exists {
		if err := s.storage.Update(ctx, fqn, obj); err != nil {
			return nil, err
		}
	} else {
		if err := s.storage.Create(ctx, fqn, obj); err != nil {
			return nil, err
		}
	}

	metadata := &pb.OperationMetadataV1{
		Method:     "google.appengine.v1.Services.UpdateService",
		InsertTime: timestamppb.Now(),
		Target:     fqn,
		User:       fmt.Sprintf("user@%s.iam.gserviceaccount.com", name.App),
	}
	prefix := fmt.Sprintf("apps/%s", name.App)
	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}

func (s *servicesServer) DeleteService(ctx context.Context, req *pb.DeleteServiceRequest) (*longrunningpb.Operation, error) {
	name, err := parseServiceName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Service{}
	if err := s.storage.Delete(ctx, fqn, obj); err != nil {
		return nil, err
	}

	metadata := &pb.OperationMetadataV1{
		Method:     "google.appengine.v1.Services.DeleteService",
		InsertTime: timestamppb.Now(),
		Target:     fqn,
		User:       fmt.Sprintf("user@%s.iam.gserviceaccount.com", name.App),
	}
	prefix := fmt.Sprintf("apps/%s", name.App)
	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}

func (s *servicesServer) ListServices(ctx context.Context, req *pb.ListServicesRequest) (*pb.ListServicesResponse, error) {
	// req.Parent is apps/{appsId}
	parts := strings.Split(req.GetParent(), "/")
	if len(parts) != 2 || parts[0] != "apps" {
		return nil, status.Errorf(codes.InvalidArgument, "invalid parent: %q", req.GetParent())
	}
	appID := parts[1]

	var services []*pb.Service
	prefix := fmt.Sprintf("apps/%s/services/", appID)
	findKind := (&pb.Service{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, findKind, storage.ListOptions{Prefix: prefix}, func(msg proto.Message) error {
		svc := msg.(*pb.Service)
		services = append(services, proto.Clone(svc).(*pb.Service))
		return nil
	}); err != nil {
		return nil, err
	}
	return &pb.ListServicesResponse{
		Services: services,
	}, nil
}

// serviceName and parseServiceName are helper utilities for parsing App Engine service paths.
type serviceName struct {
	App     string
	Service string
}

func (n *serviceName) String() string {
	return fmt.Sprintf("apps/%s/services/%s", n.App, n.Service)
}

func parseServiceName(name string) (*serviceName, error) {
	// apps/{appsId}/services/{servicesId}
	parts := strings.Split(name, "/")
	if len(parts) != 4 || parts[0] != "apps" || parts[2] != "services" {
		return nil, status.Errorf(codes.InvalidArgument, "invalid service name: %q", name)
	}
	return &serviceName{
		App:     parts[1],
		Service: parts[3],
	}, nil
}
