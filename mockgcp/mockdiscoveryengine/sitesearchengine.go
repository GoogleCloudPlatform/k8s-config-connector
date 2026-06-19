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
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/discoveryengine/apiv1/discoveryenginepb"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
)

type siteSearchEngineService struct {
	*MockService
	pb.UnimplementedSiteSearchEngineServiceServer
}

func (s *siteSearchEngineService) CreateTargetSite(ctx context.Context, req *pb.CreateTargetSiteRequest) (*longrunningpb.Operation, error) {
	parent := req.GetParent() // e.g., projects/*/locations/*/collections/*/dataStores/*/siteSearchEngine
	dsName, err := s.parseSiteSearchEngineParent(parent)
	if err != nil {
		return nil, err
	}

	targetSiteID := fmt.Sprintf("targetsite-%d", time.Now().UnixNano())
	fqn := fmt.Sprintf("%s/targetSites/%s", parent, targetSiteID)

	now := time.Now()
	obj := proto.Clone(req.GetTargetSite()).(*pb.TargetSite)
	obj.Name = fqn
	obj.UpdateTime = timestamppb.New(now)

	// Normalize/Generate Uri Pattern
	obj.GeneratedUriPattern = obj.GetProvidedUriPattern()

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%d/locations/%s/collections/%s", dsName.Project.Number, dsName.Location, dsName.Collection)
	return s.operations.DoneLRO(ctx, prefix, nil, obj)
}

func (s *siteSearchEngineService) GetTargetSite(ctx context.Context, req *pb.GetTargetSiteRequest) (*pb.TargetSite, error) {
	name, err := s.parseTargetSiteName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.TargetSite{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "TargetSite %s not found.", name)
		}
		return nil, err
	}
	return obj, nil
}

func (s *siteSearchEngineService) UpdateTargetSite(ctx context.Context, req *pb.UpdateTargetSiteRequest) (*longrunningpb.Operation, error) {
	reqTargetSite := req.GetTargetSite()
	name, err := s.parseTargetSiteName(reqTargetSite.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.TargetSite{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "TargetSite %s not found.", name)
		}
		return nil, err
	}

	// Update mutable fields
	if reqTargetSite.GetType() != pb.TargetSite_TYPE_UNSPECIFIED {
		obj.Type = reqTargetSite.GetType()
	}
	obj.ExactMatch = reqTargetSite.GetExactMatch()
	if reqTargetSite.GetProvidedUriPattern() != "" {
		obj.ProvidedUriPattern = reqTargetSite.GetProvidedUriPattern()
		obj.GeneratedUriPattern = reqTargetSite.GetProvidedUriPattern()
	}
	obj.UpdateTime = timestamppb.New(time.Now())

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%d/locations/%s/collections/%s", name.Project.Number, name.Location, name.Collection)
	return s.operations.DoneLRO(ctx, prefix, nil, obj)
}

func (s *siteSearchEngineService) DeleteTargetSite(ctx context.Context, req *pb.DeleteTargetSiteRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseTargetSiteName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	deleted := &pb.TargetSite{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%d/locations/%s/collections/%s", name.Project.Number, name.Location, name.Collection)
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
	return fmt.Sprintf("projects/%d/locations/%s/collections/%s/dataStores/%s/siteSearchEngine/targetSites/%s", n.Project.Number, n.Location, n.Collection, n.DataStore, n.TargetSite)
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
	return nil, status.Errorf(codes.InvalidArgument, "invalid name %q", name)
}

func (s *MockService) parseSiteSearchEngineParent(parent string) (*dataStoreName, error) {
	tokens := strings.Split(parent, "/")
	if len(tokens) == 9 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "collections" && tokens[6] == "dataStores" && tokens[8] == "siteSearchEngine" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}
		return &dataStoreName{
			Project:    project,
			Location:   tokens[3],
			Collection: tokens[5],
			DataStore:  tokens[7],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "invalid parent %q", parent)
}
