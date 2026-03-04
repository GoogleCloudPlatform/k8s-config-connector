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

package mockdiscoveryengine

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/discoveryengine/apiv1/discoveryenginepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
)

type siteSearchEngineService struct {
	*MockService
	pb.UnimplementedSiteSearchEngineServiceServer
}

func (s *MockService) RegisterSiteSearchEngine(grpcServer *grpc.Server) {
	pb.RegisterSiteSearchEngineServiceServer(grpcServer, &siteSearchEngineService{MockService: s})
}

func (s *siteSearchEngineService) GetTargetSite(ctx context.Context, req *pb.GetTargetSiteRequest) (*pb.TargetSite, error) {
	name, err := s.parseTargetSiteName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.TargetSite{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "TargetSite %q not found.", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *siteSearchEngineService) CreateTargetSite(ctx context.Context, req *pb.CreateTargetSiteRequest) (*longrunning.Operation, error) {
	parent, err := s.parseSiteSearchEngineName(req.Parent)
	if err != nil {
		return nil, err
	}

	// TargetSite ID is server-generated, but for mock we'll use a dummy one if not provided
	// Actually CreateTargetSite doesn't have TargetSiteId field.
	id := "target-site-123"
	fqn := parent.String() + "/targetSites/" + id

	obj := proto.Clone(req.TargetSite).(*pb.TargetSite)
	obj.Name = fqn
	obj.UpdateTime = timestamppb.Now()

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s/collections/%s/dataStores/%s/siteSearchEngine", parent.Project.ID, parent.Location, parent.Collection, parent.DataStore)
	return s.operations.DoneLRO(ctx, prefix, nil, obj)
}

func (s *siteSearchEngineService) DeleteTargetSite(ctx context.Context, req *pb.DeleteTargetSiteRequest) (*longrunning.Operation, error) {
	name, err := s.parseTargetSiteName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.TargetSite{}
	if err := s.storage.Delete(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "TargetSite %q not found.", fqn)
		}
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s/collections/%s/dataStores/%s/siteSearchEngine", name.Project.ID, name.Location, name.Collection, name.DataStore)
	return s.operations.DoneLRO(ctx, prefix, nil, nil)
}

type targetSiteName struct {
	Project    *projects.ProjectData
	Location   string
	Collection string
	DataStore  string
	TargetSite string
}

func (n *targetSiteName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/collections/%s/dataStores/%s/siteSearchEngine/targetSites/%s", n.Project.ID, n.Location, n.Collection, n.DataStore, n.TargetSite)
}

func (s *MockService) parseTargetSiteName(name string) (*targetSiteName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 11 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "collections" && tokens[6] == "dataStores" && tokens[8] == "siteSearchEngine" && tokens[9] == "targetSites" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		return &targetSiteName{
			Project:    project,
			Location:   tokens[3],
			Collection: tokens[5],
			DataStore:  tokens[7],
			TargetSite: tokens[10],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "invalid target site name %q", name)
}

type siteSearchEngineName struct {
	Project    *projects.ProjectData
	Location   string
	Collection string
	DataStore  string
}

func (n *siteSearchEngineName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/collections/%s/dataStores/%s/siteSearchEngine", n.Project.ID, n.Location, n.Collection, n.DataStore)
}

func (s *MockService) parseSiteSearchEngineName(name string) (*siteSearchEngineName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 9 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "collections" && tokens[6] == "dataStores" && tokens[8] == "siteSearchEngine" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		return &siteSearchEngineName{
			Project:    project,
			Location:   tokens[3],
			Collection: tokens[5],
			DataStore:  tokens[7],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "invalid site search engine name %q", name)
}
