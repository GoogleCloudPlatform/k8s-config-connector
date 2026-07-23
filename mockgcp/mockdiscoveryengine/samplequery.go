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
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pbb "cloud.google.com/go/discoveryengine/apiv1beta/discoveryenginepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
)

type sampleQuerySetService struct {
	*MockService
	pbb.UnimplementedSampleQuerySetServiceServer
}

func (s *sampleQuerySetService) CreateSampleQuerySet(ctx context.Context, req *pbb.CreateSampleQuerySetRequest) (*pbb.SampleQuerySet, error) {
	reqName := fmt.Sprintf("%s/sampleQuerySets/%s", req.GetParent(), req.GetSampleQuerySetId())
	name, err := s.parseSampleQuerySetName(reqName)
	if err != nil {
		return nil, err
	}
	now := time.Now()

	fqn := name.String()
	obj := proto.Clone(req.GetSampleQuerySet()).(*pbb.SampleQuerySet)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *sampleQuerySetService) GetSampleQuerySet(ctx context.Context, req *pbb.GetSampleQuerySetRequest) (*pbb.SampleQuerySet, error) {
	name, err := s.parseSampleQuerySetName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pbb.SampleQuerySet{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "SampleQuerySet %v not found.", name)
		}
		return nil, err
	}
	return obj, nil
}

func (s *sampleQuerySetService) UpdateSampleQuerySet(ctx context.Context, req *pbb.UpdateSampleQuerySetRequest) (*pbb.SampleQuerySet, error) {
	name, err := s.parseSampleQuerySetName(req.GetSampleQuerySet().GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	obj := &pbb.SampleQuerySet{}

	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "SampleQuerySet %v not found.", name)
		}
		return nil, err
	}

	// Simple Merge
	proto.Merge(obj, req.GetSampleQuerySet())

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *sampleQuerySetService) DeleteSampleQuerySet(ctx context.Context, req *pbb.DeleteSampleQuerySetRequest) (*emptypb.Empty, error) {
	name, err := s.parseSampleQuerySetName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	deleted := &pbb.SampleQuerySet{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

type sampleQueryService struct {
	*MockService
	pbb.UnimplementedSampleQueryServiceServer
}

func (s *sampleQueryService) CreateSampleQuery(ctx context.Context, req *pbb.CreateSampleQueryRequest) (*pbb.SampleQuery, error) {
	reqName := fmt.Sprintf("%s/sampleQueries/%s", req.GetParent(), req.GetSampleQueryId())
	name, err := s.parseSampleQueryName(reqName)
	if err != nil {
		return nil, err
	}
	now := time.Now()

	fqn := name.String()
	obj := proto.Clone(req.GetSampleQuery()).(*pbb.SampleQuery)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *sampleQueryService) GetSampleQuery(ctx context.Context, req *pbb.GetSampleQueryRequest) (*pbb.SampleQuery, error) {
	name, err := s.parseSampleQueryName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pbb.SampleQuery{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "SampleQuery %v not found.", name)
		}
		return nil, err
	}
	return obj, nil
}

func (s *sampleQueryService) UpdateSampleQuery(ctx context.Context, req *pbb.UpdateSampleQueryRequest) (*pbb.SampleQuery, error) {
	name, err := s.parseSampleQueryName(req.GetSampleQuery().GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	obj := &pbb.SampleQuery{}

	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "SampleQuery %v not found.", name)
		}
		return nil, err
	}

	obj.Content = req.GetSampleQuery().GetContent()

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *sampleQueryService) DeleteSampleQuery(ctx context.Context, req *pbb.DeleteSampleQueryRequest) (*emptypb.Empty, error) {
	name, err := s.parseSampleQueryName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	deleted := &pbb.SampleQuery{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
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
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}
		return &sampleQuerySetName{
			Project:        project,
			Location:       tokens[3],
			SampleQuerySet: tokens[5],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "invalid name %q", name)
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
		project, err := s.Projects.GetProjectByID(tokens[1])
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
	return nil, status.Errorf(codes.InvalidArgument, "invalid name %q", name)
}
