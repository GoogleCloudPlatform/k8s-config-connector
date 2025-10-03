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

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/servicedirectory/v1beta1"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *RegistrationServiceV1) GetService(ctx context.Context, req *pb.GetServiceRequest) (*pb.Service, error) {
	name, err := s.parseServiceName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Service{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *RegistrationServiceV1) CreateService(ctx context.Context, req *pb.CreateServiceRequest) (*pb.Service, error) {
	reqName := req.Parent + "/services/" + req.GetServiceId()
	name, err := s.parseServiceName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.Service).(*pb.Service)
	obj.Name = fqn

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *RegistrationServiceV1) UpdateService(ctx context.Context, req *pb.UpdateServiceRequest) (*pb.Service, error) {
	reqName := req.GetService().GetName()

	name, err := s.parseServiceName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Service{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil

}

func (s *RegistrationServiceV1) DeleteService(ctx context.Context, req *pb.DeleteServiceRequest) (*empty.Empty, error) {
	name, err := s.parseServiceName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	oldObj := &pb.Service{}
	if err := s.storage.Delete(ctx, fqn, oldObj); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

type serviceName struct {
	NamespaceName
	ServiceName string
}

func (n *serviceName) String() string {
	return n.NamespaceName.String() + "/services/" + n.ServiceName
}

// parseServiceName parses a string into a serviceName.
// The expected form is projects/<projectID>/locations/<location>/namespaces/<namespace>/services/<service>
func (s *MockService) parseServiceName(name string) (*serviceName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 8 && tokens[6] == "services" {
		namespacename, err := s.parseNamespaceName(strings.Join(tokens[0:6], "/"))
		if err != nil {
			return nil, err
		}

		name := &serviceName{
			NamespaceName: *namespacename,
			ServiceName:   tokens[7],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
