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

package mocknetworkservices

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	pb "cloud.google.com/go/networkservices/apiv1/networkservicespb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	edgecachepb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/networkservices/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type EdgeCacheServicesServer struct {
	*MockService
	edgecachepb.UnimplementedEdgeCacheServicesServerServer
}

type edgeCacheServiceName struct {
	Project  *projects.ProjectData
	Location string
	ID       string
}

func (n *edgeCacheServiceName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/edgeCacheServices/" + n.ID
}

func (s *MockService) parseEdgeCacheServiceName(name string) (*edgeCacheServiceName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "edgeCacheServices" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}
		return &edgeCacheServiceName{
			Project:  project,
			Location: tokens[3],
			ID:       tokens[5],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "invalid edgeCacheService name %q", name)
}

func (s *EdgeCacheServicesServer) GetEdgeCacheService(ctx context.Context, req *edgecachepb.GetEdgeCacheServiceRequest) (*edgecachepb.EdgeCacheService, error) {
	name, err := s.parseEdgeCacheServiceName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &edgecachepb.EdgeCacheService{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}
	return obj, nil
}

func (s *EdgeCacheServicesServer) ListEdgeCacheServices(ctx context.Context, req *edgecachepb.ListEdgeCacheServicesRequest) (*edgecachepb.ListEdgeCacheServicesResponse, error) {
	response := &edgecachepb.ListEdgeCacheServicesResponse{}

	prefixName, err := s.parseEdgeCacheServiceName(req.Parent + "/edgeCacheServices/placeholder")
	if err != nil {
		return nil, err
	}
	prefix := strings.TrimSuffix(prefixName.String(), "placeholder")

	findKind := (&edgecachepb.EdgeCacheService{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, findKind, storage.ListOptions{
		Prefix: prefix,
	}, func(obj proto.Message) error {
		service := obj.(*edgecachepb.EdgeCacheService)
		response.EdgeCacheServices = append(response.EdgeCacheServices, service)
		return nil
	}); err != nil {
		return nil, err
	}
	return response, nil
}

func (s *EdgeCacheServicesServer) populateDefaults(obj *edgecachepb.EdgeCacheService, project *projects.ProjectData) {
	if len(obj.Ipv4Addresses) == 0 {
		obj.Ipv4Addresses = []string{"34.153.6.71"}
	}
	if len(obj.Ipv6Addresses) == 0 {
		obj.Ipv6Addresses = []string{"2600:1900:8010:172e::"}
	}

	if obj.Routing != nil {
		for _, pm := range obj.Routing.PathMatchers {
			for _, rr := range pm.RouteRules {
				if rr.Origin != "" {
					if !strings.HasPrefix(rr.Origin, "projects/") {
						rr.Origin = fmt.Sprintf("projects/%d/locations/global/edgeCacheOrigins/%s", project.Number, rr.Origin)
					} else {
						// Normalize project id to project number in origin URL
						tokens := strings.Split(rr.Origin, "/")
						if len(tokens) >= 6 && tokens[0] == "projects" {
							if p, err := s.Projects.GetProjectByIDOrNumber(tokens[1]); err == nil {
								tokens[1] = strconv.FormatInt(p.Number, 10)
								rr.Origin = strings.Join(tokens, "/")
							}
						}
					}
				}
				if len(rr.MatchRules) == 0 {
					rr.MatchRules = []*edgecachepb.EdgeCacheServiceMatchRule{
						{
							PrefixMatch: "/",
						},
					}
				}
				if rr.HeaderAction == nil {
					rr.HeaderAction = &edgecachepb.EdgeCacheServiceHeaderAction{}
				}
				if rr.RouteAction == nil {
					rr.RouteAction = &edgecachepb.EdgeCacheServiceRouteAction{
						CdnPolicy: &edgecachepb.EdgeCacheServiceCdnPolicy{
							CacheKeyPolicy:    &edgecachepb.EdgeCacheServiceCacheKeyPolicy{},
							CacheMode:         "CACHE_ALL_STATIC",
							ClientTtl:         "3600s",
							DefaultTtl:        "3600s",
							MaxTtl:            "86400s",
							NegativeCaching:   false,
							SignedRequestMode: "DISABLED",
						},
						CompressionMode: "DISABLED",
					}
				}
				if rr.RouteMethods == nil {
					rr.RouteMethods = &edgecachepb.EdgeCacheServiceRouteMethods{
						AllowedMethods: []string{"GET", "HEAD", "OPTIONS"},
					}
				}
			}
		}
	}
}

func (s *EdgeCacheServicesServer) CreateEdgeCacheService(ctx context.Context, req *edgecachepb.CreateEdgeCacheServiceRequest) (*longrunningpb.Operation, error) {
	reqName := req.Parent + "/edgeCacheServices/" + req.EdgeCacheServiceId
	name, err := s.parseEdgeCacheServiceName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	obj := proto.Clone(req.EdgeCacheService).(*edgecachepb.EdgeCacheService)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)

	s.populateDefaults(obj, name.Project)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix := "projects/" + name.Project.ID + "/locations/" + name.Location
	lroMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		Target:     fqn,
		Verb:       "create",
		ApiVersion: "v1",
	}

	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())
		return proto.Clone(obj), nil
	})
}

func (s *EdgeCacheServicesServer) PatchEdgeCacheService(ctx context.Context, req *edgecachepb.PatchEdgeCacheServiceRequest) (*longrunningpb.Operation, error) {
	reqName := req.GetName()
	if reqName == "" && req.GetEdgeCacheService() != nil {
		reqName = req.GetEdgeCacheService().GetName()
	}
	name, err := s.parseEdgeCacheServiceName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	obj := &edgecachepb.EdgeCacheService{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		reqObj := req.GetEdgeCacheService()
		reqObj.CreateTime = obj.CreateTime
		reqObj.UpdateTime = timestamppb.New(now)
		reqObj.Name = obj.Name
		obj = reqObj
	} else {
		for _, path := range paths {
			switch path {
			case "description":
				obj.Description = req.GetEdgeCacheService().GetDescription()
			case "labels":
				obj.Labels = req.GetEdgeCacheService().GetLabels()
			case "disableQuic":
				obj.DisableQuic = req.GetEdgeCacheService().GetDisableQuic()
			case "disableHttp2":
				obj.DisableHttp2 = req.GetEdgeCacheService().GetDisableHttp2()
			case "requireTls":
				obj.RequireTls = req.GetEdgeCacheService().GetRequireTls()
			case "edgeSecurityPolicy":
				obj.EdgeSecurityPolicy = req.GetEdgeCacheService().GetEdgeSecurityPolicy()
			case "sslPolicy":
				obj.SslPolicy = req.GetEdgeCacheService().GetSslPolicy()
			case "edgeSslCertificates":
				obj.EdgeSslCertificates = req.GetEdgeCacheService().GetEdgeSslCertificates()
			case "logConfig":
				obj.LogConfig = req.GetEdgeCacheService().GetLogConfig()
			case "routing":
				obj.Routing = req.GetEdgeCacheService().GetRouting()
			default:
				return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q is not supported", path)
			}
		}
		obj.UpdateTime = timestamppb.New(now)
	}

	s.populateDefaults(obj, name.Project)

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix := "projects/" + name.Project.ID + "/locations/" + name.Location
	lroMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		Target:     fqn,
		Verb:       "update",
		ApiVersion: "v1",
	}

	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())
		return proto.Clone(obj), nil
	})
}

func (s *EdgeCacheServicesServer) DeleteEdgeCacheService(ctx context.Context, req *edgecachepb.DeleteEdgeCacheServiceRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseEdgeCacheServiceName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	deleted := &edgecachepb.EdgeCacheService{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	lroPrefix := "projects/" + name.Project.ID + "/locations/" + name.Location
	lroMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		Target:     fqn,
		Verb:       "delete",
		ApiVersion: "v1",
	}

	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())
		return &emptypb.Empty{}, nil
	})
}
