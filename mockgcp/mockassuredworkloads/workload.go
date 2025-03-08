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

// +tool:mockgcp-support
// proto.service: google.cloud.assuredworkloads.v1
// proto.message: google.cloud.assuredworkloads.v1.Workload

package mockassuredworkloads

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/assuredworkloads/v1"
)

type AssuredworkloadsV1Service struct {
	*MockService
	pb.UnimplementedAssuredWorkloadsServiceServer
}

func (s *AssuredworkloadsV1Service) GetWorkload(ctx context.Context, req *pb.GetWorkloadRequest) (*pb.Workload, error) {
	name, err := s.parseWorkloadName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Workload{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Workload %q not found.", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *AssuredworkloadsV1Service) DeleteWorkload(ctx context.Context, req *pb.DeleteWorkloadRequest) (*emptypb.Empty, error) {
	name, err := s.parseWorkloadName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.Workload{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

type workloadName struct {
	Organization string
	Location     string
	Workload     string
}

func (n *workloadName) String() string {
	return fmt.Sprintf("organizations/%s/locations/%s/workloads/%s", n.Organization, n.Location, n.Workload)
}

func (s *AssuredworkloadsV1Service) parseWorkloadName(name string) (*workloadName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "organizations" && tokens[2] == "locations" && tokens[4] == "workloads" {
		name := &workloadName{
			Organization: tokens[1],
			Location:     tokens[3],
			Workload:     tokens[5],
		}
		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}


