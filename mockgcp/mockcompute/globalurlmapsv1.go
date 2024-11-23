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

type GlobalURLMapsV1 struct {
	*MockService
	pb.UnimplementedUrlMapsServer
}

func (s *GlobalURLMapsV1) Get(ctx context.Context, req *pb.GetUrlMapRequest) (*pb.UrlMap, error) {
	reqName := "projects/" + req.GetProject() + "/global" + "/urlMaps/" + req.GetUrlMap()
	name, err := s.parseGlobalUrlMapName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.UrlMap{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, status.Errorf(codes.NotFound, "The resource '%s' was not found", fqn)
	}

	return obj, nil
}

func (s *GlobalURLMapsV1) Insert(ctx context.Context, req *pb.InsertUrlMapRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/global" + "/urlMaps/" + req.GetUrlMapResource().GetName()
	name, err := s.parseGlobalUrlMapName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	id := s.generateID()

	obj := proto.Clone(req.GetUrlMapResource()).(*pb.UrlMap)
	obj.SelfLink = PtrTo(buildComputeSelfLink(ctx, fqn))
	obj.CreationTimestamp = PtrTo(s.nowString())
	obj.Id = &id
	obj.Kind = PtrTo("compute#urlMap")

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return s.newLRO(ctx, name.Project.ID)
}

// Updates a UrlMap resource in the specified project using the data included in the request.
// This method supports PATCH semantics and uses the JSON merge patch format and processing rules.
func (s *GlobalURLMapsV1) Patch(ctx context.Context, req *pb.PatchUrlMapRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/global" + "/urlMaps/" + req.GetUrlMapResource().GetName()
	name, err := s.parseGlobalUrlMapName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.UrlMap{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// TODO: Implement helper to implement the full rules here
	proto.Merge(obj, req.GetUrlMapResource())

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return s.newLRO(ctx, name.Project.ID)
}

// Updates a UrlMap resource in the specified project using the data included in the request.
func (s *GlobalURLMapsV1) Update(ctx context.Context, req *pb.UpdateUrlMapRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/global" + "/urlMaps/" + req.GetUrlMapResource().GetName()
	name, err := s.parseGlobalUrlMapName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.UrlMap{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// TODO: Implement helper to implement the full rules here
	proto.Merge(obj, req.GetUrlMapResource())

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return s.newLRO(ctx, name.Project.ID)
}

func (s *GlobalURLMapsV1) Delete(ctx context.Context, req *pb.DeleteUrlMapRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/global" + "/urlMaps/" + req.GetUrlMap()
	name, err := s.parseGlobalUrlMapName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.UrlMap{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	return s.newLRO(ctx, name.Project.ID)
}

type globalUrlMapName struct {
	Project *projects.ProjectData
	Name    string
}

func (n *globalUrlMapName) String() string {
	return "projects/" + n.Project.ID + "/global" + "/urlMaps/" + n.Name
}

// parseGlobalUrlMapName parses a string into a globalUrlMapName.
// The expected form is `projects/*/regions/*/urlmap/*`.
func (s *MockService) parseGlobalUrlMapName(name string) (*globalUrlMapName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 5 && tokens[0] == "projects" && tokens[2] == "global" && tokens[3] == "urlMaps" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &globalUrlMapName{
			Project: project,
			Name:    tokens[4],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
