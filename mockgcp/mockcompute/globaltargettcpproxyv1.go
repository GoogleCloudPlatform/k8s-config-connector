// Copyright 2024 Google LLC
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

type GlobalTargetTcpProxyV1 struct {
	*MockService
	pb.UnimplementedTargetTcpProxiesServer
}

func (s *GlobalTargetTcpProxyV1) Get(ctx context.Context, req *pb.GetTargetTcpProxyRequest) (*pb.TargetTcpProxy, error) {
	reqName := "projects/" + req.GetProject() + "/global/targetTcpProxies/" + req.GetTargetTcpProxy()
	name, err := s.parseTargetTcpProxyName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.TargetTcpProxy{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, status.Errorf(codes.NotFound, "The resource '%s' was not found", fqn)
	}

	return obj, nil
}

func (s *GlobalTargetTcpProxyV1) Insert(ctx context.Context, req *pb.InsertTargetTcpProxyRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/global/targetTcpProxies/" + req.GetTargetTcpProxyResource().GetName()
	name, err := s.parseTargetTcpProxyName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	id := s.generateID()

	obj := proto.Clone(req.GetTargetTcpProxyResource()).(*pb.TargetTcpProxy)
	obj.SelfLink = PtrTo("https://www.googleapis.com/compute/v1/" + name.String())
	obj.CreationTimestamp = PtrTo(s.nowString())
	obj.Id = &id
	obj.Kind = PtrTo("compute#targetTcpProxy")

	if obj.ProxyHeader == nil {
		obj.ProxyHeader = PtrTo("NONE")
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("insert"),
		User:          PtrTo("user@example.com"),
	}
	return s.startGlobalLRO(ctx, name.Project.ID, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *GlobalTargetTcpProxyV1) Delete(ctx context.Context, req *pb.DeleteTargetTcpProxyRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/global/targetTcpProxies/" + req.GetTargetTcpProxy()
	name, err := s.parseTargetTcpProxyName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.TargetTcpProxy{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      deleted.Id,
		TargetLink:    deleted.SelfLink,
		OperationType: PtrTo("delete"),
		User:          PtrTo("user@example.com"),
	}
	return s.startGlobalLRO(ctx, name.Project.ID, op, func() (proto.Message, error) {
		return deleted, nil
	})
}

func (s *GlobalTargetTcpProxyV1) SetBackendService(ctx context.Context, req *pb.SetBackendServiceTargetTcpProxyRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/global/targetTcpProxies/" + req.GetTargetTcpProxy()
	name, err := s.parseTargetTcpProxyName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.TargetTcpProxy{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("TargetTcpProxySetBackendService"),
		User:          PtrTo("user@example.com"),
	}
	return s.startGlobalLRO(ctx, name.Project.ID, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *GlobalTargetTcpProxyV1) SetProxyHeader(ctx context.Context, req *pb.SetProxyHeaderTargetTcpProxyRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/global/targetTcpProxies/" + req.GetTargetTcpProxy()
	name, err := s.parseTargetTcpProxyName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.TargetTcpProxy{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	obj.ProxyHeader = req.GetTargetTcpProxiesSetProxyHeaderRequestResource().ProxyHeader
	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("TargetTcpProxySetProxyHeader"),
		User:          PtrTo("user@example.com"),
	}
	return s.startGlobalLRO(ctx, name.Project.ID, op, func() (proto.Message, error) {
		return obj, nil
	})
}

type targetTcpProxyName struct {
	Project *projects.ProjectData
	Name    string
}

func (n *targetTcpProxyName) String() string {
	return "projects/" + n.Project.ID + "/global/targetTcpProxies/" + n.Name
}

// parseTargetTcpProxyName parses a string into a targetTcpProxyName.
// The expected form is `projects/*/global/targetTcpProxies/*`.
func (s *MockService) parseTargetTcpProxyName(name string) (*targetTcpProxyName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 5 && tokens[0] == "projects" && tokens[3] == "targetTcpProxies" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &targetTcpProxyName{
			Project: project,
			Name:    tokens[4],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
