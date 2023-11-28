// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      https://www.apache.org/licenses/LICENSE-2.0
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

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	apierrors "k8s.io/apimachinery/pkg/api/errors"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
)

type RegionalTargetHTTPSProxiesV1 struct {
	*MockService
	pb.UnimplementedRegionTargetHttpsProxiesServer
}

func (s *RegionalTargetHTTPSProxiesV1) Get(ctx context.Context, req *pb.GetRegionTargetHttpsProxyRequest) (*pb.TargetHttpsProxy, error) {
	reqName := "projects/" + req.GetProject() + "/regions/" + req.GetRegion() + "/targetHttpsProxies/" + req.GetTargetHttpsProxy()
	name, err := s.parseRegionalTargetHttpsProxyName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.TargetHttpsProxy{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "targetHttpsProxy %q not found", name)
		} else {
			return nil, status.Errorf(codes.Internal, "error reading targetHttpsProxy: %v", err)
		}
	}

	return obj, nil
}

func (s *RegionalTargetHTTPSProxiesV1) Insert(ctx context.Context, req *pb.InsertRegionTargetHttpsProxyRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/regions/" + req.GetRegion() + "/targetHttpsProxies/" + req.GetTargetHttpsProxyResource().GetName()
	name, err := s.parseRegionalTargetHttpsProxyName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	id := s.generateID()

	obj := proto.Clone(req.GetTargetHttpsProxyResource()).(*pb.TargetHttpsProxy)
	obj.SelfLink = PtrTo("https://compute.googleapis.com/compute/v1/" + name.String())
	obj.CreationTimestamp = PtrTo(s.nowString())
	obj.Id = &id
	obj.Kind = PtrTo("compute#targetHttpsProxy")

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, status.Errorf(codes.Internal, "error creating targetHttpsProxy: %v", err)
	}

	return s.newLRO(ctx, name.Project.ID)
}

func (s *RegionalTargetHTTPSProxiesV1) SetUrlMap(ctx context.Context, req *pb.SetUrlMapRegionTargetHttpsProxyRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/regions/" + req.GetRegion() + "/targetHttpsProxies/" + req.GetTargetHttpsProxy()
	name, err := s.parseRegionalTargetHttpsProxyName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.TargetHttpsProxy{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "targetHttpsProxy %q not found", fqn)
		}
		return nil, status.Errorf(codes.Internal, "error reading targetHttpsProxy: %v", err)
	}

	obj.UrlMap = req.GetUrlMapReferenceResource().UrlMap

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, status.Errorf(codes.Internal, "error updating targetHttpsProxy: %v", err)
	}

	return s.newLRO(ctx, name.Project.ID)
}

func (s *RegionalTargetHTTPSProxiesV1) Delete(ctx context.Context, req *pb.DeleteRegionTargetHttpsProxyRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/regions/" + req.GetRegion() + "/targetHttpsProxies/" + req.GetTargetHttpsProxy()
	name, err := s.parseRegionalTargetHttpsProxyName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.TargetHttpsProxy{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "targetHttpsProxy %q not found", name)
		} else {
			return nil, status.Errorf(codes.Internal, "error deleting targetHttpsProxy: %v", err)
		}
	}

	return s.newLRO(ctx, name.Project.ID)
}

type regionalTargetHttpsProxyName struct {
	Project *projects.ProjectData
	Region  string
	Name    string
}

func (n *regionalTargetHttpsProxyName) String() string {
	return "projects/" + n.Project.ID + "/regions/" + n.Region + "/targetHttpsProxies/" + n.Name
}

// parseRegionalTargetHttpsProxyName parses a string into a targethttpsproxyName.
// The expected form is `projects/*/regions/*/targethttpsproxy/*`.
func (s *MockService) parseRegionalTargetHttpsProxyName(name string) (*regionalTargetHttpsProxyName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "regions" && tokens[4] == "targetHttpsProxies" {
		project, err := s.projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &regionalTargetHttpsProxyName{
			Project: project,
			Region:  tokens[3],
			Name:    tokens[5],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
