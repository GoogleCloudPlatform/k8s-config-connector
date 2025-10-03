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

type TargetSslProxyV1 struct {
	*MockService
	pb.UnimplementedTargetSslProxiesServer
}

func (s *TargetSslProxyV1) Get(ctx context.Context, req *pb.GetTargetSslProxyRequest) (*pb.TargetSslProxy, error) {
	reqName := "projects/" + req.GetProject() + "/global/targetSslProxies/" + req.GetTargetSslProxy()
	name, err := s.parseTargetSslProxyName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.TargetSslProxy{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *TargetSslProxyV1) Insert(ctx context.Context, req *pb.InsertTargetSslProxyRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/global/targetSslProxies/" + req.GetTargetSslProxyResource().GetName()
	name, err := s.parseTargetSslProxyName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	id := s.generateID()

	obj := proto.Clone(req.GetTargetSslProxyResource()).(*pb.TargetSslProxy)
	obj.SelfLink = PtrTo(buildComputeSelfLink(ctx, fqn))
	obj.CreationTimestamp = PtrTo(s.nowString())
	obj.Id = &id
	obj.Kind = PtrTo("compute#targetSslProxy")

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

func (s *TargetSslProxyV1) Delete(ctx context.Context, req *pb.DeleteTargetSslProxyRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/global/targetSslProxies/" + req.GetTargetSslProxy()
	name, err := s.parseTargetSslProxyName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.TargetSslProxy{}
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

type targetSslProxyName struct {
	Project *projects.ProjectData
	Name    string
}

func (n *targetSslProxyName) String() string {
	return "projects/" + n.Project.ID + "/global/targetSslProxies/" + n.Name
}

// parseTargetSslProxyName parses a string into a targetSslProxyName.
// The expected form is `projects/*/global/targetSslProxies/*`.
func (s *MockService) parseTargetSslProxyName(name string) (*targetSslProxyName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 5 && tokens[0] == "projects" && tokens[3] == "targetSslProxies" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &targetSslProxyName{
			Project: project,
			Name:    tokens[4],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
