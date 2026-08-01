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

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "cloud.google.com/go/discoveryengine/apiv1/discoveryenginepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
)

type siteSearchEngineService struct {
	*MockService
	pb.UnimplementedSiteSearchEngineServiceServer
}

func (s *siteSearchEngineService) GetSiteSearchEngine(ctx context.Context, req *pb.GetSiteSearchEngineRequest) (*pb.SiteSearchEngine, error) {
	name, err := s.parseSiteSearchEngineName(req.GetName())
	if err != nil {
		return nil, err
	}

	// SiteSearchEngine is a singleton nested within a DataStore.
	// Check if the parent DataStore exists.
	datastoreFQN := fmt.Sprintf("projects/%d/locations/%s/collections/%s/dataStores/%s", name.Project.Number, name.Location, name.Collection, name.DataStore)
	datastoreObj := &pb.DataStore{}
	if err := s.storage.Get(ctx, datastoreFQN, datastoreObj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "DataStore %q not found.", datastoreFQN)
		}
		return nil, err
	}

	return &pb.SiteSearchEngine{
		Name: name.String(),
	}, nil
}

type siteSearchEngineName struct {
	Project    *projects.ProjectData
	Location   string
	Collection string
	DataStore  string
}

func (n *siteSearchEngineName) String() string {
	return fmt.Sprintf("projects/%d/locations/%s/collections/%s/dataStores/%s/siteSearchEngine", n.Project.Number, n.Location, n.Collection, n.DataStore)
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
	return nil, status.Errorf(codes.InvalidArgument, "invalid siteSearchEngine name %q", name)
}
