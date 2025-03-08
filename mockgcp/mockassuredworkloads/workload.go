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
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/assuredworkloads/v1"
)

type assuredWorkloadsV1Service struct {
	*MockService
	pb.UnimplementedAssuredWorkloadsServiceServer
}

func (s *assuredWorkloadsV1Service) GetWorkload(ctx context.Context, req *pb.GetWorkloadRequest) (*pb.Workload, error) {
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

func (s *assuredWorkloadsV1Service) ListWorkloads(ctx context.Context, req *pb.ListWorkloadsRequest) (*pb.ListWorkloadsResponse, error) {
	var workloads []*pb.Workload
	workloadKind := (&pb.Workload{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, workloadKind, storage.ListOptions{}, func(obj proto.Message) error {
		workload := obj.(*pb.Workload)
		if strings.HasPrefix(workload.GetName(), req.Parent) {
			workloads = append(workloads, workload)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return &pb.ListWorkloadsResponse{
		Workloads: workloads,
	}, nil
}

func (s *assuredWorkloadsV1Service) CreateWorkload(ctx context.Context, req *pb.CreateWorkloadRequest) (*longrunning.Operation, error) {
	id := fmt.Sprintf("%d", time.Now().UnixNano())
	reqName := req.Parent + "/workloads/" + id
	name, err := s.parseWorkloadName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	obj := proto.Clone(req.Workload).(*pb.Workload)
	// Many fields are not settable because the API lib is 3 years old and needs to be updated.
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.Etag = "abcdef0123A="
	obj.Resources = []*pb.Workload_ResourceInfo{
		{
			ResourceId:   123456,
			ResourceType: pb.Workload_ResourceInfo_CONSUMER_FOLDER,
		},
	}
	obj.BillingAccount = ""

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.CreateWorkloadOperationMetadata{
		ComplianceRegime: obj.ComplianceRegime,
		CreateTime:       timestamppb.New(now),
		DisplayName:      obj.DisplayName,
		Parent:           req.Parent,
	}
	opPrefix := req.Parent
	return s.operations.StartLRO(ctx, opPrefix, op, func() (proto.Message, error) {
		result := proto.Clone(obj).(*pb.Workload)

		return result, nil
	})
}

func (s *assuredWorkloadsV1Service) DeleteWorkload(ctx context.Context, req *pb.DeleteWorkloadRequest) (*emptypb.Empty, error) {
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

func (s *assuredWorkloadsV1Service) parseWorkloadName(name string) (*workloadName, error) {
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
