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
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

type TargetGrpcProxyV1 struct {
	*MockService
	pb.UnimplementedTargetGrpcProxiesServer
}

func (s *TargetGrpcProxyV1) Get(ctx context.Context, req *pb.GetTargetGrpcProxyRequest) (*pb.TargetGrpcProxy, error) {
	reqName := "projects/" + req.GetProject() + "/global/targetGrpcProxies/" + req.GetTargetGrpcProxy()
	name, err := s.parseTargetGrpcProxyName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.TargetGrpcProxy{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, status.Errorf(codes.NotFound, "The resource '%s' was not found", fqn)
	}

	return obj, nil
}

func (s *TargetGrpcProxyV1) Insert(ctx context.Context, req *pb.InsertTargetGrpcProxyRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/global/targetGrpcProxies/" + req.GetTargetGrpcProxyResource().GetName()
	name, err := s.parseTargetGrpcProxyName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	id := s.generateID()

	obj := proto.Clone(req.GetTargetGrpcProxyResource()).(*pb.TargetGrpcProxy)
	obj.SelfLink = PtrTo("https://www.googleapis.com/compute/v1/" + name.String())
	obj.SelfLinkWithId = PtrTo(fmt.Sprintf("https://www.googleapis.com/compute/v1/projects/%s/global/targetGrpcProxies/%d", name.Project.ID, id))
	obj.CreationTimestamp = PtrTo(s.nowString())
	obj.Id = &id
	obj.Kind = PtrTo("compute#targetGrpcProxy")
	if obj.Fingerprint == nil {
		obj.Fingerprint = PtrTo(computeFingerprint(obj))
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("compute.targetGrpcProxies.insert"),
		User:          PtrTo("user@example.com"),
	}
	return s.startGlobalLRO(ctx, name.Project.ID, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *TargetGrpcProxyV1) Delete(ctx context.Context, req *pb.DeleteTargetGrpcProxyRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/global/targetGrpcProxies/" + req.GetTargetGrpcProxy()
	name, err := s.parseTargetGrpcProxyName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.TargetGrpcProxy{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      deleted.Id,
		TargetLink:    deleted.SelfLink,
		OperationType: PtrTo("compute.targetGrpcProxies.delete"),
		User:          PtrTo("user@example.com"),
	}
	return s.startGlobalLRO(ctx, name.Project.ID, op, func() (proto.Message, error) {
		return deleted, nil
	})
}

type targetGrpcProxyName struct {
	Project *projects.ProjectData
	Name    string
}

func (n *targetGrpcProxyName) String() string {
	return "projects/" + n.Project.ID + "/global/targetGrpcProxies/" + n.Name
}

// parseTargetGrpcProxyName parses a string into a targetGrpcProxyName.
// The expected form is `projects/*/global/parseTargetGrpcProxies/*`.
func (s *MockService) parseTargetGrpcProxyName(name string) (*targetGrpcProxyName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 5 && tokens[0] == "projects" && tokens[3] == "targetGrpcProxies" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &targetGrpcProxyName{
			Project: project,
			Name:    tokens[4],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
