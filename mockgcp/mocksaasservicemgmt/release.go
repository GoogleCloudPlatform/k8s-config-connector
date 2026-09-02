// Copyright 2026 Google LLC
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

package mocksaasservicemgmt

import (
	"context"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/saasplatform/saasservicemgmt/apiv1beta1/saasservicemgmtpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
)

type SaasDeploymentsServer struct {
	*MockService
	pb.UnimplementedSaasDeploymentsServer
}

type releaseName struct {
	Project  *projects.ProjectData
	Location string
	Release  string
}

func (n *releaseName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/releases/" + n.Release
}

func (s *MockService) parseReleaseName(name string) (*releaseName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "releases" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		return &releaseName{
			Project:  project,
			Location: tokens[3],
			Release:  tokens[5],
		}, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}

func (s *SaasDeploymentsServer) GetRelease(ctx context.Context, req *pb.GetReleaseRequest) (*pb.Release, error) {
	name, err := s.parseReleaseName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Release{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *SaasDeploymentsServer) CreateRelease(ctx context.Context, req *pb.CreateReleaseRequest) (*pb.Release, error) {
	reqName := req.Parent + "/releases/" + req.ReleaseId
	name, err := s.parseReleaseName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.Release).(*pb.Release)
	obj.Name = fqn

	now := timestamppb.Now()
	obj.CreateTime = now
	obj.UpdateTime = now
	obj.Etag = "mock-etag"

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	response := proto.Clone(obj).(*pb.Release)
	response.Etag = ""
	return response, nil
}

func (s *SaasDeploymentsServer) UpdateRelease(ctx context.Context, req *pb.UpdateReleaseRequest) (*pb.Release, error) {
	reqName := req.GetRelease().GetName()
	name, err := s.parseReleaseName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Release{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	// Update fields
	desired := req.GetRelease()
	obj.Labels = desired.Labels
	obj.Annotations = desired.Annotations
	obj.InputVariableDefaults = desired.InputVariableDefaults
	obj.ReleaseRequirements = desired.ReleaseRequirements
	obj.Blueprint = desired.Blueprint
	obj.UpdateTime = timestamppb.Now()
	obj.Etag = "mock-etag-updated"

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	response := proto.Clone(obj).(*pb.Release)
	response.Etag = ""
	return response, nil
}

func (s *SaasDeploymentsServer) DeleteRelease(ctx context.Context, req *pb.DeleteReleaseRequest) (*emptypb.Empty, error) {
	name, err := s.parseReleaseName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Release{}
	if err := s.storage.Delete(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
