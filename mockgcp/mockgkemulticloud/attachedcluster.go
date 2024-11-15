// Copyright 2023 Google LLC
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

package mockgkemulticloud

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/gkemulticloud/apiv1/gkemulticloudpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/fields"
)

type GKEMulticloudV1 struct {
	*MockService
	pb.UnimplementedAttachedClustersServer
}

func (s *GKEMulticloudV1) GetAttachedCluster(ctx context.Context, req *pb.GetAttachedClusterRequest) (*pb.AttachedCluster, error) {
	name, err := s.parseAttachedClustersName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.AttachedCluster{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "cluster not found")
		}
		return nil, err
	}

	return obj, nil
}

func (s *GKEMulticloudV1) CreateAttachedCluster(ctx context.Context, req *pb.CreateAttachedClusterRequest) (*longrunning.Operation, error) {
	reqName := req.Parent + "/attachedClusters/" + req.AttachedClusterId
	name, err := s.parseAttachedClustersName(reqName)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	fqn := name.String()

	obj := proto.Clone(req.AttachedCluster).(*pb.AttachedCluster)
	obj.Name = fqn

	if obj.GetBinaryAuthorization() == nil {
		obj.BinaryAuthorization = &pb.BinaryAuthorization{
			EvaluationMode: pb.BinaryAuthorization_DISABLED,
		}
	}
	if obj.GetMonitoringConfig() == nil {
		obj.MonitoringConfig = &pb.MonitoringConfig{
			ManagedPrometheusConfig: &pb.ManagedPrometheusConfig{},
		}
	}
	obj.Fleet.Membership = obj.Fleet.Project + "/locations/global/memberships/" + req.AttachedClusterId
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.Etag = fields.ComputeWeakEtag(obj)
	ver, err := trimPlatformVersion(req.GetAttachedCluster().GetPlatformVersion())
	if err != nil {
		return nil, err
	}
	obj.KubernetesVersion = ver
	obj.State = pb.AttachedCluster_RUNNING
	obj.Uid = "111111111111111111111"

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	opMetadata := &pb.OperationMetadata{
		CreateTime:            timestamppb.New(now),
		Verb:                  "create",
		RequestedCancellation: false,
		Target:                fqn,
	}
	opPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, opPrefix, opMetadata, func() (proto.Message, error) {
		opMetadata.EndTime = timestamppb.Now()
		obj.State = pb.AttachedCluster_RUNNING
		return obj, nil
	})
}

func (s *GKEMulticloudV1) UpdateAttachedCluster(ctx context.Context, req *pb.UpdateAttachedClusterRequest) (*longrunning.Operation, error) {
	name, err := s.parseAttachedClustersName(req.GetAttachedCluster().GetName())
	if err != nil {
		return nil, err
	}

	now := time.Now()
	fqn := name.String()
	obj := &pb.AttachedCluster{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}
	// Mask of fields to update. At least one path must be supplied in
	// this field. The elements of the repeated paths field can only include these
	// fields from
	// [AttachedCluster][mockgcp.cloud.gkemulticloud.v1.AttachedCluster]:
	//
	//   - `annotations`.
	//   - `authorization.admin_users`.
	//   - `binary_authorization.evaluation_mode`.
	//   - `description`.
	//   - `logging_config.component_config.enable_components`.
	//   - `monitoring_config.managed_prometheus_config.enabled`.
	//   - `platform_version`.

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "must specify updateMask")
	}

	// TODO: Some sort of helper for fieldmask?
	for _, path := range paths {
		switch path {
		case "annotations":
			obj.Annotations = req.GetAttachedCluster().GetAnnotations()
		case "authorization.admin_users":
			obj.Authorization.AdminUsers = req.GetAttachedCluster().GetAuthorization().GetAdminUsers()
		case "binary_authorization.evaluation_mode":
			obj.BinaryAuthorization.EvaluationMode = req.GetAttachedCluster().GetBinaryAuthorization().GetEvaluationMode()
		case "description":
			obj.Description = req.GetAttachedCluster().GetDescription()
		case "logging_config.component_config.enable_components":
			obj.LoggingConfig.ComponentConfig.EnableComponents = req.GetAttachedCluster().GetLoggingConfig().GetComponentConfig().GetEnableComponents()
		case "monitoring_config.managed_prometheus_config.enabled":
			obj.MonitoringConfig = req.GetAttachedCluster().GetMonitoringConfig()
		case "platformVersion":
			ver, err := trimPlatformVersion(req.GetAttachedCluster().GetPlatformVersion())
			if err != nil {
				return nil, err
			}
			obj.KubernetesVersion = ver
			obj.PlatformVersion = req.GetAttachedCluster().GetPlatformVersion()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
		}
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}
	opMetadata := &pb.OperationMetadata{
		CreateTime:            timestamppb.New(now),
		Verb:                  "update",
		RequestedCancellation: false,
		Target:                fqn,
	}
	opPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, opPrefix, opMetadata, func() (proto.Message, error) {
		opMetadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}

func (s *GKEMulticloudV1) DeleteAttachedCluster(ctx context.Context, req *pb.DeleteAttachedClusterRequest) (*longrunning.Operation, error) {
	name, err := s.parseAttachedClustersName(req.Name)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	fqn := name.String()

	oldObj := &pb.AttachedCluster{}
	if err := s.storage.Delete(ctx, fqn, oldObj); err != nil {
		return nil, err
	}

	opMetadata := &pb.OperationMetadata{
		CreateTime:            timestamppb.New(now),
		Verb:                  "delete",
		RequestedCancellation: false,
		Target:                fqn,
	}
	opPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, opPrefix, opMetadata, func() (proto.Message, error) {
		opMetadata.EndTime = timestamppb.Now()
		return &emptypb.Empty{}, nil
	})
}

func trimPlatformVersion(platformVersion string) (string, error) {
	tokens := strings.Split(platformVersion, ".")
	if len(tokens) < 2 {
		return "", status.Errorf(codes.InvalidArgument, "platform_version %q is not valid", platformVersion)
	}
	return tokens[0] + "." + tokens[1], nil
}
