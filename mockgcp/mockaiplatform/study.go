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

package mockaiplatform

import (
	"context"
	"crypto/sha256"
	"fmt"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

type vizierService struct {
	*MockService
	pb.UnimplementedVizierServiceServer
}

type studyName struct {
	Project  string
	Location string
	Study    string
}

func (n *studyName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/studies/%s", n.Project, n.Location, n.Study)
}

func (s *vizierService) parseStudyName(name string) (*studyName, error) {
	if name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "name must be specified")
	}

	parts := strings.Split(name, "/")
	if len(parts) == 6 && parts[0] == "projects" && parts[2] == "locations" && parts[4] == "studies" {
		return &studyName{
			Project:  parts[1],
			Location: parts[3],
			Study:    parts[5],
		}, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid study name", name)
}

func (s *vizierService) GetStudy(ctx context.Context, req *pb.GetStudyRequest) (*pb.Study, error) {
	name, err := s.parseStudyName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Study{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *vizierService) CreateStudy(ctx context.Context, req *pb.CreateStudyRequest) (*pb.Study, error) {
	// Real Vizier service generates a study ID (e.g., study-123456)
	// We generate it deterministically using a hash of the DisplayName to ensure stable golden files.
	h := sha256.Sum256([]byte(req.Study.GetDisplayName()))
	studyID := fmt.Sprintf("study-%x", h[:8])
	fqn := fmt.Sprintf("%s/studies/%s", req.Parent, studyID)

	now := time.Now()

	obj := proto.Clone(req.Study).(*pb.Study)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.State = pb.Study_ACTIVE

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *vizierService) ListStudies(ctx context.Context, req *pb.ListStudiesRequest) (*pb.ListStudiesResponse, error) {
	prefix := req.Parent + "/studies/"

	var matching []*pb.Study
	err := s.storage.List(ctx, (&pb.Study{}).ProtoReflect().Descriptor(), storage.ListOptions{
		Prefix: prefix,
	}, func(obj proto.Message) error {
		matching = append(matching, obj.(*pb.Study))
		return nil
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "listing studies: %v", err)
	}

	return &pb.ListStudiesResponse{
		Studies: matching,
	}, nil
}

func (s *vizierService) DeleteStudy(ctx context.Context, req *pb.DeleteStudyRequest) (*emptypb.Empty, error) {
	name, err := s.parseStudyName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.Study{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
