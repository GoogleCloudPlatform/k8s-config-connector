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

package mockfirebasehosting

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/firebasehosting/v1beta1"
)

type ProjectsSitesServer struct {
	*MockService
	pb.UnimplementedProjectsSitesServerServer
}

func (s *ProjectsSitesServer) GetProjectsSite(ctx context.Context, req *pb.GetProjectsSiteRequest) (*pb.Site, error) {
	name, err := s.parseSiteName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Site{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *ProjectsSitesServer) CreateProjectsSite(ctx context.Context, req *pb.CreateProjectsSiteRequest) (*pb.Site, error) {
	siteID := req.GetSiteId()
	if siteID == "" {
		return nil, status.Errorf(codes.InvalidArgument, "site_id is required")
	}

	parentTokens := strings.Split(req.GetParent(), "/")
	if len(parentTokens) != 2 || parentTokens[0] != "projects" {
		return nil, status.Errorf(codes.InvalidArgument, "parent %q is malformed", req.GetParent())
	}

	reqName := fmt.Sprintf("%s/sites/%s", req.GetParent(), siteID)
	name, err := s.parseSiteName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := proto.Clone(req.GetProjectsSite()).(*pb.Site)
	if obj == nil {
		obj = &pb.Site{}
	}

	s.populateDefaultsForSite(obj, name)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *ProjectsSitesServer) DeleteProjectsSite(ctx context.Context, req *pb.DeleteProjectsSiteRequest) (*pb.Empty, error) {
	name, err := s.parseSiteName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	oldObj := &pb.Site{}
	if err := s.storage.Delete(ctx, fqn, oldObj); err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}

func (s *ProjectsSitesServer) PatchProjectsSite(ctx context.Context, req *pb.PatchProjectsSiteRequest) (*pb.Site, error) {
	name, err := s.parseSiteName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	curObj := &pb.Site{}
	if err := s.storage.Get(ctx, fqn, curObj); err != nil {
		return nil, err
	}

	paths := strings.Split(req.GetUpdateMask(), ",")
	for _, path := range paths {
		path = strings.TrimSpace(path)
		switch path {
		case "appId", "app_id":
			curObj.AppId = req.GetProjectsSite().AppId
		default:
			return nil, status.Errorf(codes.InvalidArgument, "field %q is not yet handled in mock", path)
		}
	}

	if err := s.storage.Update(ctx, fqn, curObj); err != nil {
		return nil, err
	}

	return curObj, nil
}

func (s *MockService) populateDefaultsForSite(obj *pb.Site, name *siteName) {
	fullName := name.String()
	obj.Name = &fullName

	if obj.DefaultUrl == nil {
		defaultURL := fmt.Sprintf("https://%s.web.app", name.SiteID)
		obj.DefaultUrl = &defaultURL
	}

	if obj.Type == nil {
		siteType := "USER_SITE"
		obj.Type = &siteType
	}
}
