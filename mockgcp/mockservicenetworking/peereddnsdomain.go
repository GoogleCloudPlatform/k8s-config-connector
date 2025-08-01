// Copyright 2025 Google LLC
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

package mockservicenetworking

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/servicenetworking/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

type peeredDnsDomainsServer struct {
	*MockService
	pb.UnimplementedServicesProjectsGlobalNetworksPeeredDnsDomainsServerServer
}

func (s *peeredDnsDomainsServer) CreateServicesProjectsGlobalNetworksPeeredDnsDomain(ctx context.Context, req *pb.CreateServicesProjectsGlobalNetworksPeeredDnsDomainRequest) (*longrunningpb.Operation, error) {
	reqName := fmt.Sprintf("%s/peeredDnsDomains/%s", req.GetParent(), req.GetServicesProjectsGlobalNetworksPeeredDnsDomain().GetName())
	name, err := s.parsePeeredDnsDomainName(reqName)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := proto.Clone(req.GetServicesProjectsGlobalNetworksPeeredDnsDomain()).(*pb.PeeredDnsDomain)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return s.operations.StartLRO(ctx, "", nil, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *peeredDnsDomainsServer) ListServicesProjectsGlobalNetworksPeeredDnsDomains(ctx context.Context, req *pb.ListServicesProjectsGlobalNetworksPeeredDnsDomainsRequest) (*pb.ListPeeredDnsDomainsResponse, error) {
	reqName := fmt.Sprintf("%s/peeredDnsDomains/placeholder", req.GetParent())
	name, err := s.parsePeeredDnsDomainName(reqName)
	if err != nil {
		return nil, err
	}
	prefix := strings.TrimSuffix(name.String(), "placeholder")

	var domains []*pb.PeeredDnsDomain
	kind := (&pb.PeeredDnsDomain{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, kind, storage.ListOptions{Prefix: prefix}, func(o proto.Message) error {
		domain := o.(*pb.PeeredDnsDomain)
		domains = append(domains, domain)
		return nil
	}); err != nil {
		return nil, err
	}

	return &pb.ListPeeredDnsDomainsResponse{
		PeeredDnsDomains: domains,
	}, nil
}

func (s *peeredDnsDomainsServer) DeleteServicesProjectsGlobalNetworksPeeredDnsDomain(ctx context.Context, req *pb.DeleteServicesProjectsGlobalNetworksPeeredDnsDomainRequest) (*longrunningpb.Operation, error) {
	name, err := s.parsePeeredDnsDomainName(req.GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	deleted := &pb.PeeredDnsDomain{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	return s.operations.StartLRO(ctx, "", nil, func() (proto.Message, error) {
		return &emptypb.Empty{}, nil
	})
}

type peeredDnsDomainName struct {
	Project   *projects.ProjectData
	Service   string
	Network   string
	DnsDomain string
}

func (n *peeredDnsDomainName) String() string {
	return fmt.Sprintf("services/%s/projects/%s/global/networks/%s/peeredDnsDomains/%s", n.Service, n.Project.ID, n.Network, n.DnsDomain)
}

func (s *MockService) parsePeeredDnsDomainName(name string) (*peeredDnsDomainName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 9 && tokens[0] == "services" && tokens[2] == "projects" && tokens[4] == "global" && tokens[5] == "networks" && tokens[7] == "peeredDnsDomains" {
		// Note: servicenetworking requires this to be a number, not a project ID
		project, err := s.Projects.GetProjectByNumber(tokens[3])
		if err != nil {
			return nil, err
		}

		name := &peeredDnsDomainName{
			Project:   project,
			Service:   tokens[1],
			Network:   tokens[6],
			DnsDomain: tokens[8],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
