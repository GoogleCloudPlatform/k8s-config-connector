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

// +tool:mockgcp-support
// proto.service: google.cloud.discoveryengine.v1beta.SampleQuerySetService
// proto.message: google.cloud.discoveryengine.v1beta.SampleQuerySet

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

	pb_v1beta "cloud.google.com/go/discoveryengine/apiv1beta/discoveryenginepb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
)

type sampleQuerySetService struct {
	*MockService
	pb_v1beta.UnimplementedSampleQuerySetServiceServer
}

func (s *sampleQuerySetService) CreateSampleQuerySet(ctx context.Context, req *pb_v1beta.CreateSampleQuerySetRequest) (*pb_v1beta.SampleQuerySet, error) {
	reqName := fmt.Sprintf("%s/sampleQuerySets/%s", req.GetParent(), req.GetSampleQuerySetId())
	name, err := s.parseSampleQuerySetName(reqName)
	if err != nil {
		return nil, err
	}
	now := time.Now()

	fqn := name.String()
	obj := proto.Clone(req.GetSampleQuerySet()).(*pb_v1beta.SampleQuerySet)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *sampleQuerySetService) GetSampleQuerySet(ctx context.Context, req *pb_v1beta.GetSampleQuerySetRequest) (*pb_v1beta.SampleQuerySet, error) {
	name, err := s.parseSampleQuerySetName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb_v1beta.SampleQuerySet{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "SampleQuerySet with name %q does not exist.", fqn)
		}
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
			return nil, status.Errorf(codes.NotFound, "SampleQuerySet %q not found", name)
		}
		return nil, err
	}

	// simple merge for now
	proto.Merge(obj, req.GetSampleQuerySet())
	obj.Name = fqn

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

	deleted := &pb_v1beta.SampleQuerySet{}
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
