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

package mockhealthcare

import (
	"context"
	"strings"

	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/healthcare/v1"
)

type datasetsServer struct {
	*MockService
	pb.UnimplementedProjectsLocationsDatasetsServerServer
}

func (s *datasetsServer) GetProjectsLocationsDataset(ctx context.Context, req *pb.GetProjectsLocationsDatasetRequest) (*pb.Dataset, error) {
	name, err := s.parseDatasetName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Dataset{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *datasetsServer) CreateProjectsLocationsDataset(ctx context.Context, req *pb.CreateProjectsLocationsDatasetRequest) (*longrunning.Operation, error) {
	reqName := req.GetParent() + "/datasets/" + req.GetDatasetId()
	name, err := s.parseDatasetName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.GetProjectsLocationsDataset()).(*pb.Dataset)
	obj.Name = fqn

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return s.operations.DoneLRO(ctx, fqn, nil, obj)
}

func (s *datasetsServer) DeleteProjectsLocationsDataset(ctx context.Context, req *pb.DeleteProjectsLocationsDatasetRequest) (*pb.Empty, error) {
	name, err := s.parseDatasetName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.Dataset{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}

func (s *datasetsServer) PatchProjectsLocationsDataset(ctx context.Context, req *pb.PatchProjectsLocationsDatasetRequest) (*pb.Dataset, error) {
	name, err := s.parseDatasetName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Dataset{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	updateMask := req.GetUpdateMask()
	if updateMask == "" {
		proto.Merge(obj, req.GetProjectsLocationsDataset())
	} else {
		for _, path := range strings.Split(updateMask, ",") {
			switch path {
			case "timeZone":
				obj.TimeZone = req.GetProjectsLocationsDataset().GetTimeZone()
			default:
				return nil, status.Errorf(codes.InvalidArgument, "unsupported update_mask path %q", path)
			}
		}
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}
