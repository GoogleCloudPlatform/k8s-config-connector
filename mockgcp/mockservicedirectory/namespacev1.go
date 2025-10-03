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

package mockservicedirectory

import (
	"context"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/servicedirectory/v1beta1"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

type RegistrationServiceV1 struct {
	*MockService
	pb.UnimplementedRegistrationServiceServer
}

func (s *RegistrationServiceV1) GetNamespace(ctx context.Context, req *pb.GetNamespaceRequest) (*pb.Namespace, error) {
	name, err := s.parseNamespaceName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Namespace{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *RegistrationServiceV1) CreateNamespace(ctx context.Context, req *pb.CreateNamespaceRequest) (*pb.Namespace, error) {
	reqName := req.Parent + "/namespaces/" + req.GetNamespaceId()
	name, err := s.parseNamespaceName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.Namespace).(*pb.Namespace)
	obj.Name = fqn

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *RegistrationServiceV1) UpdateNamespace(ctx context.Context, req *pb.UpdateNamespaceRequest) (*pb.Namespace, error) {
	reqName := req.GetNamespace().GetName()

	name, err := s.parseNamespaceName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Namespace{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil

}

func (s *RegistrationServiceV1) DeleteNamespace(ctx context.Context, req *pb.DeleteNamespaceRequest) (*empty.Empty, error) {
	name, err := s.parseNamespaceName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	oldObj := &pb.Namespace{}
	if err := s.storage.Delete(ctx, fqn, oldObj); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

type NamespaceName struct {
	Project       *projects.ProjectData
	Location      string
	NamespaceName string
}

func (n *NamespaceName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/namespaces/" + n.NamespaceName
}

// parseNamespaceName parses a string into a namespaceName.
// The expected form is projects/<projectID>/locations/<location>/namespaces/<namespace>
func (s *MockService) parseNamespaceName(name string) (*NamespaceName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "namespaces" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &NamespaceName{
			Project:       project,
			Location:      tokens[3],
			NamespaceName: tokens[5],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
