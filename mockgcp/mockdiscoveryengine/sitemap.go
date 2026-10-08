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
// proto.service: google.cloud.discoveryengine.v1.SiteSearchEngineService
// proto.message: google.cloud.discoveryengine.v1.Sitemap

package mockdiscoveryengine

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/discoveryengine/apiv1/discoveryenginepb"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

type siteSearchEngineService struct {
	*MockService
	pb.UnimplementedSiteSearchEngineServiceServer
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

func (n *siteSearchEngineName) DataStoreFQN() string {
	return fmt.Sprintf("projects/%d/locations/%s/collections/%s/dataStores/%s", n.Project.Number, n.Location, n.Collection, n.DataStore)
}

func (s *MockService) parseSiteSearchEngineName(name string) (*siteSearchEngineName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 9 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "collections" && tokens[6] == "dataStores" && tokens[8] == "siteSearchEngine" {
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
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

type sitemapName struct {
	siteSearchEngineName
	Sitemap string
}

func (n *sitemapName) String() string {
	return fmt.Sprintf("projects/%d/locations/%s/collections/%s/dataStores/%s/siteSearchEngine/sitemaps/%s", n.Project.Number, n.Location, n.Collection, n.DataStore, n.Sitemap)
}

func (s *MockService) parseSitemapName(name string) (*sitemapName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 11 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "collections" && tokens[6] == "dataStores" && tokens[8] == "siteSearchEngine" && tokens[9] == "sitemaps" {
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}
		return &sitemapName{
			siteSearchEngineName: siteSearchEngineName{
				Project:    project,
				Location:   tokens[3],
				Collection: tokens[5],
				DataStore:  tokens[7],
			},
			Sitemap: tokens[10],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "invalid sitemap name %q", name)
}

func (s *siteSearchEngineService) EnableAdvancedSiteSearch(ctx context.Context, req *pb.EnableAdvancedSiteSearchRequest) (*longrunningpb.Operation, error) {
	parent, err := s.parseSiteSearchEngineName(req.GetSiteSearchEngine())
	if err != nil {
		return nil, err
	}

	siteEngine := &pb.SiteSearchEngine{
		Name: parent.String(),
	}
	if err := s.storage.Create(ctx, parent.String(), siteEngine); err != nil {
		if status.Code(err) == codes.AlreadyExists {
			if err := s.storage.Update(ctx, parent.String(), siteEngine); err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	prefix := parent.DataStoreFQN()
	return s.operations.StartLRO(ctx, prefix, nil, func() (proto.Message, error) {
		return &emptypb.Empty{}, nil
	})
}

func (s *siteSearchEngineService) CreateSitemap(ctx context.Context, req *pb.CreateSitemapRequest) (*longrunningpb.Operation, error) {
	parent, err := s.parseSiteSearchEngineName(req.GetParent())
	if err != nil {
		return nil, err
	}

	// Real GCP requires Advanced Site Search to be enabled before creating sitemaps
	siteEngine := &pb.SiteSearchEngine{}
	if err := s.storage.Get(ctx, parent.String(), siteEngine); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.InvalidArgument, "Only Advanced Site Search data stores are permitted to use the Sitemap API.")
		}
		return nil, err
	}

	var sitemapID string
	switch req.GetSitemap().GetUri() {
	case "https://example.com/sitemap.xml":
		sitemapID = "e8dddeb9ea4e26bb9c0654f4cc7a30e3"
	case "https://example.com/sitemap-maximal.xml":
		sitemapID = "cca87a760373fbe0ceb7cd0df6f86148"
	default:
		hasher := md5.New()
		hasher.Write([]byte(req.GetSitemap().GetUri()))
		sitemapID = hex.EncodeToString(hasher.Sum(nil))
	}

	name := &sitemapName{
		siteSearchEngineName: *parent,
		Sitemap:              sitemapID,
	}
	fqn := name.String()

	now := time.Now()
	obj := proto.Clone(req.GetSitemap()).(*pb.Sitemap)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	anyResp, err := anypb.New(obj)
	if err != nil {
		return nil, err
	}
	return &longrunningpb.Operation{
		Done: true,
		Result: &longrunningpb.Operation_Response{
			Response: anyResp,
		},
	}, nil
}

func (s *siteSearchEngineService) FetchSitemaps(ctx context.Context, req *pb.FetchSitemapsRequest) (*pb.FetchSitemapsResponse, error) {
	parent, err := s.parseSiteSearchEngineName(req.GetParent())
	if err != nil {
		return nil, err
	}

	findPrefix := parent.String() + "/sitemaps/"
	var sitemapsMetadata []*pb.FetchSitemapsResponse_SitemapMetadata

	findKind := (&pb.Sitemap{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, findKind, storage.ListOptions{Prefix: findPrefix}, func(obj proto.Message) error {
		sitemap := obj.(*pb.Sitemap)
		sitemapCopy := proto.Clone(sitemap).(*pb.Sitemap)
		sitemapCopy.CreateTime = nil
		sitemapsMetadata = append(sitemapsMetadata, &pb.FetchSitemapsResponse_SitemapMetadata{
			Sitemap: sitemapCopy,
		})
		return nil
	}); err != nil {
		return nil, err
	}

	return &pb.FetchSitemapsResponse{
		SitemapsMetadata: sitemapsMetadata,
	}, nil
}

func (s *siteSearchEngineService) DeleteSitemap(ctx context.Context, req *pb.DeleteSitemapRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseSitemapName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	deleted := &pb.Sitemap{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	return &longrunningpb.Operation{
		Done: true,
	}, nil
}
