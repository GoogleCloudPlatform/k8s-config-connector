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

type GlobalTargetHTTPProxiesV1 struct {
	*MockService
	pb.UnimplementedTargetHttpProxiesServer
}

func (s *GlobalTargetHTTPProxiesV1) Get(ctx context.Context, req *pb.GetTargetHttpProxyRequest) (*pb.TargetHttpProxy, error) {
	reqName := "projects/" + req.GetProject() + "/global" + "/targetHttpProxies/" + req.GetTargetHttpProxy()
	name, err := s.parseGlobalTargetHttpProxyName(reqName)
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

func (s *GlobalTargetHTTPProxiesV1) Insert(ctx context.Context, req *pb.InsertTargetHttpProxyRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/global" + "/targetHttpProxies/" + req.GetTargetHttpProxyResource().GetName()
	name, err := s.parseGlobalTargetHttpProxyName(reqName)
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

// Updates a TargetHttpProxy resource in the specified project using the data included in the request.
// This method supports PATCH semantics and uses the JSON merge patch format and processing rules.
func (s *GlobalTargetHTTPProxiesV1) Patch(ctx context.Context, req *pb.PatchTargetHttpProxyRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/global" + "/targetHttpProxies/" + req.GetTargetHttpProxyResource().GetName()
	name, err := s.parseGlobalTargetHttpProxyName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.TargetHttpProxy{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// TODO: Implement helper to implement the full rules here
	proto.Merge(obj, req.GetTargetHttpProxyResource())

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return s.newLRO(ctx, name.Project.ID)
}

func (s *GlobalTargetHTTPProxiesV1) SetUrlMap(ctx context.Context, req *pb.SetUrlMapTargetHttpProxyRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/global" + "/targetHttpProxies/" + req.GetTargetHttpProxy()
	name, err := s.parseGlobalTargetHttpProxyName(reqName)
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

func (s *GlobalTargetHTTPProxiesV1) Delete(ctx context.Context, req *pb.DeleteTargetHttpProxyRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/global" + "/targetHttpProxies/" + req.GetTargetHttpProxy()
	name, err := s.parseGlobalTargetHttpProxyName(reqName)
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

type globalTargetHttpProxyName struct {
	Project *projects.ProjectData
	Name    string
}

func (n *globalTargetHttpProxyName) String() string {
	return "projects/" + n.Project.ID + "/global" + "/targetHttpProxies/" + n.Name
}

// parseGlobalTargetHttpProxyName parses a string into a globalTargetHttpProxyName.
// The expected form is `projects/*/regions/*/targethttpproxy/*`.
func (s *MockService) parseGlobalTargetHttpProxyName(name string) (*globalTargetHttpProxyName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 5 && tokens[0] == "projects" && tokens[2] == "global" && tokens[3] == "targetHttpProxies" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &globalTargetHttpProxyName{
			Project: project,
			Name:    tokens[4],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
