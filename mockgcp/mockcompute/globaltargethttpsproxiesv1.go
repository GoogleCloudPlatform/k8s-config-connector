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

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	apierrors "k8s.io/apimachinery/pkg/api/errors"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
)

type GlobalTargetHTTPSProxiesV1 struct {
	*MockService
	pb.UnimplementedTargetHttpsProxiesServer
}

func (s *GlobalTargetHTTPSProxiesV1) Get(ctx context.Context, req *pb.GetTargetHttpsProxyRequest) (*pb.TargetHttpsProxy, error) {
	reqName := "projects/" + req.GetProject() + "/global" + "/targetHttpsProxies/" + req.GetTargetHttpsProxy()
	name, err := s.parseGlobalTargetHttpsProxyName(reqName)
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

func (s *GlobalTargetHTTPSProxiesV1) Insert(ctx context.Context, req *pb.InsertTargetHttpsProxyRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/global" + "/targetHttpsProxies/" + req.GetTargetHttpsProxyResource().GetName()
	name, err := s.parseGlobalTargetHttpsProxyName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	id := s.generateID()

	obj := proto.Clone(req.GetTargetHttpsProxyResource()).(*pb.TargetHttpsProxy)
	obj.SelfLink = PtrTo("Https://compute.googleapis.com/compute/v1/" + name.String())
	obj.CreationTimestamp = PtrTo(s.nowString())
	obj.Id = &id
	obj.Kind = PtrTo("compute#targetHttpsProxy")

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, status.Errorf(codes.Internal, "error creating targetHttpsProxy: %v", err)
	}

	return s.newLRO(ctx, name.Project.ID)
}

// Updates a TargetHttpsProxy resource in the specified project using the data included in the request.
// This method supports PATCH semantics and uses the JSON merge patch format and processing rules.
func (s *GlobalTargetHTTPSProxiesV1) Patch(ctx context.Context, req *pb.PatchTargetHttpsProxyRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/global" + "/targetHttpsProxies/" + req.GetTargetHttpsProxyResource().GetName()
	name, err := s.parseGlobalTargetHttpsProxyName(reqName)
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

	// TODO: Implement helper to implement the full rules here
	proto.Merge(obj, req.GetTargetHttpsProxyResource())

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, status.Errorf(codes.Internal, "error updating targetHttpsProxy: %v", err)
	}

	return s.newLRO(ctx, name.Project.ID)
}

func (s *GlobalTargetHTTPSProxiesV1) SetUrlMap(ctx context.Context, req *pb.SetUrlMapTargetHttpsProxyRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/global" + "/targetHttpsProxies/" + req.GetTargetHttpsProxy()
	name, err := s.parseGlobalTargetHttpsProxyName(reqName)
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

func (s *GlobalTargetHTTPSProxiesV1) SetQuicOverride(ctx context.Context, req *pb.SetQuicOverrideTargetHttpsProxyRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/global" + "/targetHttpsProxies/" + req.GetTargetHttpsProxy()
	name, err := s.parseGlobalTargetHttpsProxyName(reqName)
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

	obj.QuicOverride = req.GetTargetHttpsProxiesSetQuicOverrideRequestResource().QuicOverride

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, status.Errorf(codes.Internal, "error updating targetHttpsProxy: %v", err)
	}

	return s.newLRO(ctx, name.Project.ID)
}

func (s *GlobalTargetHTTPSProxiesV1) Delete(ctx context.Context, req *pb.DeleteTargetHttpsProxyRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/global" + "/targetHttpsProxies/" + req.GetTargetHttpsProxy()
	name, err := s.parseGlobalTargetHttpsProxyName(reqName)
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

type globalTargetHttpsProxyName struct {
	Project *projects.ProjectData
	Name    string
}

func (n *globalTargetHttpsProxyName) String() string {
	return "projects/" + n.Project.ID + "/global" + "/targetHttpsProxies/" + n.Name
}

// parseGlobalTargetHttpsProxyName parses a string into a globalTargetHttpsProxyName.
// The expected form is `projects/*/regions/*/targetHttpsproxy/*`.
func (s *MockService) parseGlobalTargetHttpsProxyName(name string) (*globalTargetHttpsProxyName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 5 && tokens[0] == "projects" && tokens[2] == "global" && tokens[3] == "targetHttpsProxies" {
		project, err := s.projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &globalTargetHttpsProxyName{
			Project: project,
			Name:    tokens[4],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
