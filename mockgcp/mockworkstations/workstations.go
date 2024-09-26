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

package mockworkstations

import (
	"context"
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"strings"
	"time"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/workstations/v1"
)

type WorkstationsService struct {
	*MockService
	pb.UnimplementedWorkstationsServer
}

func (s *WorkstationsService) GetWorkstationCluster(ctx context.Context, req *pb.GetWorkstationClusterRequest) (*pb.WorkstationCluster, error) {
	fqn := req.GetName()

	obj := &pb.WorkstationCluster{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Requested entity was not found.")
		}
		return nil, err
	}

	return obj, nil
}

func (s *WorkstationsService) CreateWorkstationCluster(ctx context.Context, req *pb.CreateWorkstationClusterRequest) (*longrunningpb.Operation, error) {
	fqn := req.GetParent() + "/workstationClusters/" + req.GetWorkstationClusterId()

	obj := proto.Clone(req.WorkstationCluster).(*pb.WorkstationCluster)
	populateDefaultsForWorkstationCluster(obj, false)
	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	t := timestamppb.New(time.Now())
	metadata := &pb.OperationMetadata{
		CreateTime:            t,
		ApiVersion:            "v1",
		RequestedCancellation: false,
		Target:                fqn,
		Verb:                  "create",
	}
	op, err := s.operations.StartLRO(ctx, req.GetParent(), metadata, func() (proto.Message, error) {
		metadata.EndTime = t
		result := proto.Clone(obj).(*pb.WorkstationCluster)
		populateOutputsForWorkstationCluster(result, fqn)
		s.storage.Update(ctx, fqn, result)
		return result, nil
	})
	return op, err
}

func (s *WorkstationsService) UpdateWorkstationCluster(ctx context.Context, req *pb.UpdateWorkstationClusterRequest) (*longrunningpb.Operation, error) {
	fqn := req.GetWorkstationCluster().GetName()

	existing := &pb.WorkstationCluster{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		return nil, err
	}

	updated := proto.Clone(req.WorkstationCluster).(*pb.WorkstationCluster)
	populateDefaultsForWorkstationCluster(updated, true)
	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}

	op, err := s.operations.StartLRO(ctx, fqn, nil, func() (proto.Message, error) {
		result := proto.Clone(updated).(*pb.WorkstationCluster)
		populateOutputsForWorkstationCluster(result, fqn)
		return result, nil
	})
	if err != nil {
		return op, err
	}
	response, err := anypb.New(updated)
	if err != nil {
		return op, err
	}
	op.Result = &longrunningpb.Operation_Response{
		Response: response,
	}
	return op, err
}

func (s *WorkstationsService) DeleteWorkstationCluster(ctx context.Context, req *pb.DeleteWorkstationClusterRequest) (*longrunningpb.Operation, error) {
	fqn := req.GetName()
	parent, err := getParent(fqn)
	if err != nil {
		return nil, err
	}

	deleted := &pb.WorkstationCluster{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	t := timestamppb.New(time.Now())
	metadata := &pb.OperationMetadata{
		CreateTime:            t,
		ApiVersion:            "v1",
		RequestedCancellation: false,
		Target:                fqn,
		Verb:                  "delete",
	}
	op, err := s.operations.StartLRO(ctx, parent, metadata, func() (proto.Message, error) {
		metadata.EndTime = t
		return &pb.WorkstationCluster{}, nil
	})
	return op, err
}

func getParent(fqn string) (string, error) {
	tokens := strings.Split(fqn, "/")
	if len(tokens) != 6 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "workstationClusters" {
		return "", fmt.Errorf("fqn should be projects/<project>/locations/<location>/workstationClusters/<WorkstationCluster>, got %s", fqn)
	}
	return tokens[0] + "/" + tokens[1] + "/" + tokens[2] + "/" + tokens[3], nil
}

func populateDefaultsForWorkstationCluster(obj *pb.WorkstationCluster, update bool) {
	if obj.Uid == "" {
		obj.Uid = fmt.Sprintf("%x", time.Now().UnixNano())
	}
	t := timestamppb.New(time.Now())
	if obj.CreateTime == nil {
		obj.CreateTime = t
	}
	if obj.UpdateTime == nil || update {
		obj.UpdateTime = t
	}
	if obj.PrivateClusterConfig == nil {
		obj.PrivateClusterConfig = &pb.WorkstationCluster_PrivateClusterConfig{}
	}
	obj.Etag = computeEtag(obj)
}

func populateOutputsForWorkstationCluster(obj *pb.WorkstationCluster, fqn string) {
	if obj.Name == "" {
		obj.Name = fqn
	}
	if obj.ControlPlaneIp == "" {
		obj.ControlPlaneIp = "10.0.0.2"
	}
	if obj.Etag == "" {
		obj.Etag = computeEtag(obj)
	}
	if obj.Uid == "" {
		obj.Uid = fmt.Sprintf("%x", time.Now().UnixNano())
	}
	if obj.PrivateClusterConfig.EnablePrivateEndpoint {
		obj.PrivateClusterConfig.ClusterHostname = "cluster-abcdef.cloudworkstations.dev"
		obj.PrivateClusterConfig.ServiceAttachmentUri = "https://www.googleapis.com/compute/v1/projects/${projectId}/regions/us-west1/serviceAttachments/k8s1-sa-abcdef-cloudshell-gateway-abcdef"
	}
}

func computeEtag(obj proto.Message) string {
	b, err := proto.Marshal(obj)
	if err != nil {
		panic(fmt.Sprintf("converting to proto: %v", err))
	}
	hash := md5.Sum(b)
	return base64.StdEncoding.EncodeToString(hash[:])
}
