// Copyright 2022 Google LLC
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

package mockcompute

import (
	"context"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

type RegionalTargetHTTPProxiesV1 struct {
	*MockService
	pb.UnimplementedRegionTargetHttpProxiesServer
}

func (s *RegionalTargetHTTPProxiesV1) Get(ctx context.Context, req *pb.GetRegionTargetHttpProxyRequest) (*pb.TargetHttpProxy, error) {
	reqName := "projects/" + req.GetProject() + "/regions/" + req.GetRegion() + "/targetHttpProxies/" + req.GetTargetHttpProxy()
	name, err := s.parseRegionalTargetHttpProxyName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.TargetHttpProxy{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *RegionalTargetHTTPProxiesV1) Insert(ctx context.Context, req *pb.InsertRegionTargetHttpProxyRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/regions/" + req.GetRegion() + "/targetHttpProxies/" + req.GetTargetHttpProxyResource().GetName()
	name, err := s.parseRegionalTargetHttpProxyName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	id := s.generateID()

	obj := proto.Clone(req.GetTargetHttpProxyResource()).(*pb.TargetHttpProxy)
	obj.SelfLink = PtrTo(buildComputeSelfLink(ctx, fqn))
	obj.CreationTimestamp = PtrTo(s.nowString())
	obj.Id = &id
	obj.Kind = PtrTo("compute#targetHttpProxy")

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return s.newLRO(ctx, name.Project.ID)
}

func (s *RegionalTargetHTTPProxiesV1) SetUrlMap(ctx context.Context, req *pb.SetUrlMapRegionTargetHttpProxyRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/regions/" + req.GetRegion() + "/targetHttpProxies/" + req.GetTargetHttpProxy()
	name, err := s.parseRegionalTargetHttpProxyName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.TargetHttpProxy{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	obj.UrlMap = req.GetUrlMapReferenceResource().UrlMap

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return s.newLRO(ctx, name.Project.ID)
}

func (s *RegionalTargetHTTPProxiesV1) Delete(ctx context.Context, req *pb.DeleteRegionTargetHttpProxyRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/regions/" + req.GetRegion() + "/targetHttpProxies/" + req.GetTargetHttpProxy()
	name, err := s.parseRegionalTargetHttpProxyName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.TargetHttpProxy{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	return s.newLRO(ctx, name.Project.ID)
}

type regionalTargetHttpProxyName struct {
	Project *projects.ProjectData
	Region  string
	Name    string
}

func (n *regionalTargetHttpProxyName) String() string {
	return "projects/" + n.Project.ID + "/regions/" + n.Region + "/targetHttpProxies/" + n.Name
}

// parseRegionalTargetHttpProxyName parses a string into a targethttpproxyName.
// The expected form is `projects/*/regions/*/targethttpproxy/*`.
func (s *MockService) parseRegionalTargetHttpProxyName(name string) (*regionalTargetHttpProxyName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "regions" && tokens[4] == "targetHttpProxies" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &regionalTargetHttpProxyName{
			Project: project,
			Region:  tokens[3],
			Name:    tokens[5],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
