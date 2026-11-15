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

package mockbigquery

import (
	"context"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/bigquery/v2"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type rowAccessPoliciesServer struct {
	*MockService
	pb.UnimplementedRowAccessPoliciesServerServer
}

type rowAccessPolicyName struct {
	Project   *projects.ProjectData
	DatasetID string
	TableID   string
	PolicyID  string
}

func (n *rowAccessPolicyName) String() string {
	return "projects/" + n.Project.ID + "/datasets/" + n.DatasetID + "/tables/" + n.TableID + "/rowAccessPolicies/" + n.PolicyID
}

func (s *MockService) buildRowAccessPolicyName(projectName, datasetID, tableID, policyID string) (*rowAccessPolicyName, error) {
	project, err := s.Projects.GetProjectByID(projectName)
	if err != nil {
		return nil, err
	}

	return &rowAccessPolicyName{
		Project:   project,
		DatasetID: datasetID,
		TableID:   tableID,
		PolicyID:  policyID,
	}, nil
}

func (s *rowAccessPoliciesServer) GetRowAccessPolicy(ctx context.Context, req *pb.GetRowAccessPolicyRequest) (*pb.RowAccessPolicy, error) {
	name, err := s.buildRowAccessPolicyName(req.GetProjectId(), req.GetDatasetId(), req.GetTableId(), req.GetPolicyId())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.RowAccessPolicy{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Not found: Row access policy %s on table %s:%s.%s", name.PolicyID, name.Project.ID, name.DatasetID, name.TableID)
		}
		return nil, err
	}

	response := proto.Clone(obj).(*pb.RowAccessPolicy)
	response.Grantees = nil
	response.Etag = nil
	return response, nil
}

func (s *rowAccessPoliciesServer) InsertRowAccessPolicy(ctx context.Context, req *pb.InsertRowAccessPolicyRequest) (*pb.RowAccessPolicy, error) {
	tableName, err := s.buildTableName(req.GetProjectId(), req.GetDatasetId(), req.GetTableId())
	if err != nil {
		return nil, err
	}
	table := &pb.Table{}
	if err := s.storage.Get(ctx, tableName.String(), table); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Not found: Table %s:%s.%s", tableName.Project.ID, tableName.DatasetID, tableName.TableID)
		}
		return nil, err
	}

	policyID := req.GetRowAccessPolicy().GetRowAccessPolicyReference().GetPolicyId()
	name, err := s.buildRowAccessPolicyName(req.GetProjectId(), req.GetDatasetId(), req.GetTableId(), policyID)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := timestamppb.Now()

	obj := proto.Clone(req.GetRowAccessPolicy()).(*pb.RowAccessPolicy)
	if obj.RowAccessPolicyReference == nil {
		obj.RowAccessPolicyReference = &pb.RowAccessPolicyReference{}
	}
	obj.RowAccessPolicyReference.ProjectId = PtrTo(name.Project.ID)
	obj.RowAccessPolicyReference.DatasetId = PtrTo(name.DatasetID)
	obj.RowAccessPolicyReference.TableId = PtrTo(name.TableID)
	obj.RowAccessPolicyReference.PolicyId = PtrTo(name.PolicyID)
	obj.CreationTime = now
	obj.LastModifiedTime = now

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	response := proto.Clone(obj).(*pb.RowAccessPolicy)
	response.Grantees = nil
	response.Etag = nil
	return response, nil
}

func (s *rowAccessPoliciesServer) UpdateRowAccessPolicy(ctx context.Context, req *pb.UpdateRowAccessPolicyRequest) (*pb.RowAccessPolicy, error) {
	name, err := s.buildRowAccessPolicyName(req.GetProjectId(), req.GetDatasetId(), req.GetTableId(), req.GetPolicyId())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	existing := &pb.RowAccessPolicy{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Not found: Row access policy %s on table %s:%s.%s", name.PolicyID, name.Project.ID, name.DatasetID, name.TableID)
		}
		return nil, err
	}

	now := timestamppb.Now()

	updated := proto.Clone(req.GetRowAccessPolicy()).(*pb.RowAccessPolicy)
	updated.CreationTime = existing.CreationTime
	updated.LastModifiedTime = now
	updated.RowAccessPolicyReference = existing.RowAccessPolicyReference

	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}

	response := proto.Clone(updated).(*pb.RowAccessPolicy)
	response.Grantees = nil
	response.Etag = nil
	return response, nil
}

func (s *rowAccessPoliciesServer) DeleteRowAccessPolicy(ctx context.Context, req *pb.DeleteRowAccessPolicyRequest) (*empty.Empty, error) {
	name, err := s.buildRowAccessPolicyName(req.GetProjectId(), req.GetDatasetId(), req.GetTableId(), req.GetPolicyId())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.RowAccessPolicy{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Not found: Row access policy %s on table %s:%s.%s", name.PolicyID, name.Project.ID, name.DatasetID, name.TableID)
		}
		return nil, err
	}

	return &empty.Empty{}, nil
}

func (s *rowAccessPoliciesServer) ListRowAccessPolicies(ctx context.Context, req *pb.ListRowAccessPoliciesRequest) (*pb.ListRowAccessPoliciesResponse, error) {
	name, err := s.buildTableName(req.GetProjectId(), req.GetDatasetId(), req.GetTableId())
	if err != nil {
		return nil, err
	}

	response := &pb.ListRowAccessPoliciesResponse{}
	prefix := name.String() + "/rowAccessPolicies/"
	if err := s.storage.List(ctx, (&pb.RowAccessPolicy{}).ProtoReflect().Descriptor(), storage.ListOptions{
		Prefix: prefix,
	}, func(obj proto.Message) error {
		policy := proto.Clone(obj).(*pb.RowAccessPolicy)
		policy.Grantees = nil
		policy.Etag = nil
		response.RowAccessPolicies = append(response.RowAccessPolicies, policy)
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}
