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

package mockdiscoveryengine

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"

	pb_v1beta "cloud.google.com/go/discoveryengine/apiv1beta/discoveryenginepb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/fields"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
)

type sampleQuerySetService struct {
	*MockService
	pb_v1beta.UnimplementedSampleQuerySetServiceServer
}

func (s *sampleQuerySetService) GetSampleQuerySet(ctx context.Context, req *pb_v1beta.GetSampleQuerySetRequest) (*pb_v1beta.SampleQuerySet, error) {
	name, err := s.parseSampleQuerySetName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb_v1beta.SampleQuerySet{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "SampleQuerySet %s does not exist.", fqn)
		}
		return nil, err
	}
	return obj, nil
}

func (s *sampleQuerySetService) CreateSampleQuerySet(ctx context.Context, req *pb_v1beta.CreateSampleQuerySetRequest) (*pb_v1beta.SampleQuerySet, error) {
	reqName := fmt.Sprintf("%s/sampleQuerySets/%s", req.GetParent(), req.GetSampleQuerySetId())
	name, err := s.parseSampleQuerySetName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := proto.Clone(req.GetSampleQuerySet()).(*pb_v1beta.SampleQuerySet)
	obj.Name = fqn

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *sampleQuerySetService) UpdateSampleQuerySet(ctx context.Context, req *pb_v1beta.UpdateSampleQuerySetRequest) (*pb_v1beta.SampleQuerySet, error) {
	name, err := s.parseSampleQuerySetName(req.GetSampleQuerySet().GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	obj := &pb_v1beta.SampleQuerySet{}

	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "SampleQuerySet %s does not exist.", fqn)
		}
		return nil, err
	}

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask must be provided")
	}

	if err := fields.UpdateByFieldMask(obj, req.GetSampleQuerySet(), paths); err != nil {
		return nil, err
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *sampleQuerySetService) DeleteSampleQuerySet(ctx context.Context, req *pb_v1beta.DeleteSampleQuerySetRequest) (*emptypb.Empty, error) {
	name, err := s.parseSampleQuerySetName(req.GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	obj := &pb_v1beta.SampleQuerySet{}
	if err := s.storage.Delete(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "SampleQuerySet %s does not exist.", fqn)
		}
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

type sampleQueryService struct {
	*MockService
	pb_v1beta.UnimplementedSampleQueryServiceServer
}

func (s *sampleQueryService) GetSampleQuery(ctx context.Context, req *pb_v1beta.GetSampleQueryRequest) (*pb_v1beta.SampleQuery, error) {
	name, err := s.parseSampleQueryName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb_v1beta.SampleQuery{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "SampleQuery %s does not exist.", fqn)
		}
		return nil, err
	}
	return obj, nil
}

func (s *sampleQueryService) CreateSampleQuery(ctx context.Context, req *pb_v1beta.CreateSampleQueryRequest) (*pb_v1beta.SampleQuery, error) {
	reqName := fmt.Sprintf("%s/sampleQueries/%s", req.GetParent(), req.GetSampleQueryId())
	name, err := s.parseSampleQueryName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := proto.Clone(req.GetSampleQuery()).(*pb_v1beta.SampleQuery)
	obj.Name = fqn

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *sampleQueryService) UpdateSampleQuery(ctx context.Context, req *pb_v1beta.UpdateSampleQueryRequest) (*pb_v1beta.SampleQuery, error) {
	name, err := s.parseSampleQueryName(req.GetSampleQuery().GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	obj := &pb_v1beta.SampleQuery{}

	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "SampleQuery %s does not exist.", fqn)
		}
		return nil, err
	}

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask must be provided")
	}

	if err := fields.UpdateByFieldMask(obj, req.GetSampleQuery(), paths); err != nil {
		return nil, err
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *sampleQueryService) DeleteSampleQuery(ctx context.Context, req *pb_v1beta.DeleteSampleQueryRequest) (*emptypb.Empty, error) {
	name, err := s.parseSampleQueryName(req.GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	obj := &pb_v1beta.SampleQuery{}
	if err := s.storage.Delete(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "SampleQuery %s does not exist.", fqn)
		}
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

type sampleQuerySetName struct {
	Project        *projects.ProjectData
	Location       string
	SampleQuerySet string
}

func (n *sampleQuerySetName) String() string {
	return fmt.Sprintf("projects/%d/locations/%s/sampleQuerySets/%s", n.Project.Number, n.Location, n.SampleQuerySet)
}

func (s *MockService) parseSampleQuerySetName(name string) (*sampleQuerySetName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "sampleQuerySets" {
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}
		return &sampleQuerySetName{
			Project:        project,
			Location:       tokens[3],
			SampleQuerySet: tokens[5],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "invalid sample query set name %q", name)
}

type sampleQueryName struct {
	Project        *projects.ProjectData
	Location       string
	SampleQuerySet string
	SampleQuery    string
}

func (n *sampleQueryName) String() string {
	return fmt.Sprintf("projects/%d/locations/%s/sampleQuerySets/%s/sampleQueries/%s", n.Project.Number, n.Location, n.SampleQuerySet, n.SampleQuery)
}

func (s *MockService) parseSampleQueryName(name string) (*sampleQueryName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "sampleQuerySets" && tokens[6] == "sampleQueries" {
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}
		return &sampleQueryName{
			Project:        project,
			Location:       tokens[3],
			SampleQuerySet: tokens[5],
			SampleQuery:    tokens[7],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "invalid sample query name %q", name)
}
