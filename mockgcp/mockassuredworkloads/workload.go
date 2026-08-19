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

package mockassuredworkloads

import (
	"context"
	"fmt"
	"math/rand"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/assuredworkloads/v1"
)

type AssuredWorkloadsV1 struct {
	*MockService
	pb.UnimplementedAssuredWorkloadsServiceServer
}

func (s *AssuredWorkloadsV1) GetWorkload(ctx context.Context, req *pb.GetWorkloadRequest) (*pb.Workload, error) {
	name, err := s.parseWorkloadName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Workload{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *AssuredWorkloadsV1) CreateWorkload(ctx context.Context, req *pb.CreateWorkloadRequest) (*longrunningpb.Operation, error) {
	reqName := req.Parent + "/workloads/temp" // We'll generate a real name
	name, err := s.parseWorkloadName(reqName)
	if err != nil {
		return nil, err
	}

	// Generate a random ID for the workload
	// Real GCP uses a mix of prefix and random hex, e.g. 00-f987...
	workloadID := fmt.Sprintf("00-%08x-%04x-%04x-%04x-%012x", rand.Uint32(), rand.Uint32()&0xffff, rand.Uint32()&0xffff, rand.Uint32()&0xffff, rand.Uint64())
	name.Workload = workloadID

	fqn := name.String()

	obj := proto.Clone(req.Workload).(*pb.Workload)
	obj.Name = fqn
	obj.CreateTime = timestamppb.Now()
	obj.Etag = "mock-etag"

	// Mock resources
	obj.Resources = []*pb.Workload_ResourceInfo{
		{
			ResourceId:   rand.Int63n(1000000000000),
			ResourceType: pb.Workload_ResourceInfo_CONSUMER_PROJECT,
		},
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// Manual construction of LRO to avoid type mismatch
	lro := &longrunningpb.Operation{
		Name: "operations/mock-lro-" + workloadID,
		Done: true,
	}
	resultAny, err := anypb.New(obj)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error building anypb: %v", err)
	}
	lro.Result = &longrunningpb.Operation_Response{
		Response: resultAny,
	}

	return lro, nil
}

func (s *AssuredWorkloadsV1) UpdateWorkload(ctx context.Context, req *pb.UpdateWorkloadRequest) (*pb.Workload, error) {
	name, err := s.parseWorkloadName(req.Workload.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Workload{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// Currently allows updating of workload display_name and labels.
	if req.UpdateMask != nil {
		for _, path := range req.UpdateMask.Paths {
			switch path {
			case "display_name":
				obj.DisplayName = req.Workload.DisplayName
			case "labels":
				obj.Labels = req.Workload.Labels
			default:
				return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not supported in mock", path)
			}
		}
	} else {
		obj.DisplayName = req.Workload.DisplayName
		obj.Labels = req.Workload.Labels
	}

	obj.Etag = "mock-etag-updated"

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *AssuredWorkloadsV1) DeleteWorkload(ctx context.Context, req *pb.DeleteWorkloadRequest) (*emptypb.Empty, error) {
	name, err := s.parseWorkloadName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	oldObj := &pb.Workload{}
	if err := s.storage.Delete(ctx, fqn, oldObj); err != nil {
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

func (s *MockService) parseWorkloadName(name string) (*workloadName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "organizations" && tokens[2] == "locations" && tokens[4] == "workloads" {
		return &workloadName{
			Organization: tokens[1],
			Location:     tokens[3],
			Workload:     tokens[5],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "invalid workload name %q", name)
}
