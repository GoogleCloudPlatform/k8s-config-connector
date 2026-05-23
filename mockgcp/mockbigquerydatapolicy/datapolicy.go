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

package mockbigquerydatapolicy

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "cloud.google.com/go/bigquery/datapolicies/apiv1beta1/datapoliciespb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
)

type DataPolicyV1Beta1 struct {
	*MockService
	pb.UnimplementedDataPolicyServiceServer
}

type dataPolicyName struct {
	Project      *projects.ProjectData
	Location     string
	DataPolicyID string
}

func (s *DataPolicyV1Beta1) parseDataPolicyName(name string) (*dataPolicyName, error) {
	if name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "name is required")
	}
	name = strings.ReplaceAll(name, "/dataPolicies/", "/datapolicies/")
	tokens := strings.Split(name, "/")
	if len(tokens) != 6 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "datapolicies" {
		return nil, status.Errorf(codes.InvalidArgument, "invalid DataPolicy name %q", name)
	}
	project, err := s.Projects.GetProjectByID(tokens[1])
	if err != nil {
		return nil, err
	}
	return &dataPolicyName{
		Project:      project,
		Location:     tokens[3],
		DataPolicyID: tokens[5],
	}, nil
}

func (n *dataPolicyName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/dataPolicies/%s", n.Project.ID, n.Location, n.DataPolicyID)
}

func (s *DataPolicyV1Beta1) GetDataPolicy(ctx context.Context, req *pb.GetDataPolicyRequest) (*pb.DataPolicy, error) {
	name, err := s.parseDataPolicyName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.DataPolicy{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Not found: DataPolicy %s", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *DataPolicyV1Beta1) CreateDataPolicy(ctx context.Context, req *pb.CreateDataPolicyRequest) (*pb.DataPolicy, error) {
	if req.Parent == "" {
		return nil, status.Errorf(codes.InvalidArgument, "parent is required")
	}
	if req.DataPolicy == nil {
		return nil, status.Errorf(codes.InvalidArgument, "data_policy is required")
	}

	dataPolicyID := req.DataPolicy.DataPolicyId
	if dataPolicyID == "" {
		return nil, status.Errorf(codes.InvalidArgument, "data_policy_id is required")
	}

	reqName := req.Parent + "/datapolicies/" + dataPolicyID
	name, err := s.parseDataPolicyName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.DataPolicy).(*pb.DataPolicy)
	obj.Name = fqn

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *DataPolicyV1Beta1) UpdateDataPolicy(ctx context.Context, req *pb.UpdateDataPolicyRequest) (*pb.DataPolicy, error) {
	if req.DataPolicy == nil {
		return nil, status.Errorf(codes.InvalidArgument, "data_policy is required")
	}

	name, err := s.parseDataPolicyName(req.DataPolicy.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	existing := &pb.DataPolicy{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		return nil, err
	}

	// For mock purposes, just merge the fields
	updated := proto.Clone(existing).(*pb.DataPolicy)
	if req.DataPolicy.DataPolicyType != pb.DataPolicy_DATA_POLICY_TYPE_UNSPECIFIED {
		updated.DataPolicyType = req.DataPolicy.DataPolicyType
	}
	if req.DataPolicy.MatchingLabel != nil {
		updated.MatchingLabel = req.DataPolicy.MatchingLabel
	}
	if req.DataPolicy.Policy != nil {
		updated.Policy = req.DataPolicy.Policy
	}

	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}

	return updated, nil
}

func (s *DataPolicyV1Beta1) DeleteDataPolicy(ctx context.Context, req *pb.DeleteDataPolicyRequest) (*emptypb.Empty, error) {
	name, err := s.parseDataPolicyName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	deleted := &pb.DataPolicy{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Not found: DataPolicy %s", fqn)
		}
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
