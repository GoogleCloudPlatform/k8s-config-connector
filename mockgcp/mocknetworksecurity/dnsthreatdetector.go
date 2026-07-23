// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mocknetworksecurity

import (
	"context"
	"strings"
	"time"

	pbv1 "cloud.google.com/go/networksecurity/apiv1/networksecuritypb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type DnsThreatDetectorServer struct {
	*MockService
	pbv1.UnimplementedDnsThreatDetectorServiceServer
}

func (s *DnsThreatDetectorServer) CreateDnsThreatDetector(ctx context.Context, req *pbv1.CreateDnsThreatDetectorRequest) (*pbv1.DnsThreatDetector, error) {
	name := req.Parent + "/dnsThreatDetectors/" + req.DnsThreatDetectorId

	fqn := name

	obj := proto.Clone(req.DnsThreatDetector).(*pbv1.DnsThreatDetector)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(time.Now())
	obj.UpdateTime = timestamppb.New(time.Now())

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *DnsThreatDetectorServer) GetDnsThreatDetector(ctx context.Context, req *pbv1.GetDnsThreatDetectorRequest) (*pbv1.DnsThreatDetector, error) {
	name, err := s.parseDnsThreatDetectorName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pbv1.DnsThreatDetector{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}
	return obj, nil
}

func (s *DnsThreatDetectorServer) UpdateDnsThreatDetector(ctx context.Context, req *pbv1.UpdateDnsThreatDetectorRequest) (*pbv1.DnsThreatDetector, error) {
	name, err := s.parseDnsThreatDetectorName(req.GetDnsThreatDetector().GetName())
	if err != nil {
		return nil, err
	}
	obj := &pbv1.DnsThreatDetector{}
	if err := s.storage.Get(ctx, name.String(), obj); err != nil {
		return nil, err
	}

	updated := proto.Clone(obj).(*pbv1.DnsThreatDetector)
	updated.UpdateTime = timestamppb.New(time.Now())

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		// If update_mask is not provided, all fields should be overwritten.
		updated = proto.Clone(req.GetDnsThreatDetector()).(*pbv1.DnsThreatDetector)
		updated.CreateTime = obj.CreateTime
		updated.UpdateTime = timestamppb.New(time.Now())
		updated.Name = obj.Name
	} else {
		for _, path := range paths {
			switch path {
			case "labels":
				updated.Labels = req.GetDnsThreatDetector().GetLabels()
			case "excluded_networks":
				updated.ExcludedNetworks = req.GetDnsThreatDetector().GetExcludedNetworks()
			case "provider":
				updated.Provider = req.GetDnsThreatDetector().GetProvider()
			default:
				return nil, status.Errorf(codes.InvalidArgument, "field %q is not yet handled in mock", path)
			}
		}
	}

	if err := s.storage.Update(ctx, name.String(), updated); err != nil {
		return nil, err
	}

	return updated, nil
}

func (s *DnsThreatDetectorServer) DeleteDnsThreatDetector(ctx context.Context, req *pbv1.DeleteDnsThreatDetectorRequest) (*emptypb.Empty, error) {
	name, err := s.parseDnsThreatDetectorName(req.Name)
	if err != nil {
		return nil, err
	}

	if err := s.storage.Delete(ctx, name.String(), &pbv1.DnsThreatDetector{}); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

type dnsThreatDetectorName struct {
	Project             *projects.ProjectData
	Location            string
	DnsThreatDetectorID string
}

func (n *dnsThreatDetectorName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/dnsThreatDetectors/" + n.DnsThreatDetectorID
}

func (s *DnsThreatDetectorServer) parseDnsThreatDetectorName(name string) (*dnsThreatDetectorName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "dnsThreatDetectors" {
		project, err := s.Projects.GetProject(&projects.ProjectName{ProjectID: tokens[1]})
		if err != nil {
			return nil, err
		}
		name := &dnsThreatDetectorName{
			Project:             project,
			Location:            tokens[3],
			DnsThreatDetectorID: tokens[5],
		}
		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
