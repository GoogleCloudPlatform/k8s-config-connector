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

package mockbigtable

import (
	"context"
	"fmt"
	"strings"
	"time"

	longrunning "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"k8s.io/apimachinery/pkg/util/sets"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"

	// Note: we use the "real" proto (not mockgcp), because the client uses GRPC.
	pb "cloud.google.com/go/bigtable/admin/apiv2/adminpb"
)

type instanceAdminServer struct {
	*MockService
	pb.UnimplementedBigtableInstanceAdminServer
}

func (s *instanceAdminServer) GetInstance(ctx context.Context, req *pb.GetInstanceRequest) (*pb.Instance, error) {
	name, err := s.parseInstanceName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Instance{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Instance %s not found.", name.InstanceName)
		}
		return nil, err
	}

	return obj, nil
}

func (s *instanceAdminServer) ListInstances(ctx context.Context, req *pb.ListInstancesRequest) (*pb.ListInstancesResponse, error) {
	projectName, err := projects.ParseProjectName(req.GetParent())
	if err != nil {
		return nil, err
	}
	project, err := s.Projects.GetProject(projectName)
	if err != nil {
		return nil, err
	}

	response := &pb.ListInstancesResponse{}
	findKind := (&pb.Instance{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, findKind, storage.ListOptions{
		Prefix: "projects/" + project.ID + "/",
	}, func(obj proto.Message) error {
		instance := obj.(*pb.Instance)
		response.Instances = append(response.Instances, instance)
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}

func (s *instanceAdminServer) CreateInstance(ctx context.Context, req *pb.CreateInstanceRequest) (*longrunning.Operation, error) {
	reqName := req.GetParent() + "/instances/" + req.GetInstanceId()
	name, err := s.parseInstanceName(reqName)
	if err != nil {
		return nil, err
	}

	// "Clean up" the original request; this syncs the metadata OriginalRequest field
	for _, reqCluster := range req.GetClusters() {
		if proto.Equal(reqCluster.EncryptionConfig, &pb.Cluster_EncryptionConfig{}) {
			reqCluster.EncryptionConfig = nil
		}
		if reqCluster.DefaultStorageType == pb.StorageType_STORAGE_TYPE_UNSPECIFIED {
			reqCluster.DefaultStorageType = pb.StorageType_SSD
		}
		reqCluster.Name = ""
	}

	originalRequest := proto.Clone(req).(*pb.CreateInstanceRequest)

	now := time.Now()
	instanceFQN := name.String()

	obj := proto.Clone(req.Instance).(*pb.Instance)
	obj.Name = instanceFQN

	obj.State = pb.Instance_READY
	obj.CreateTime = timestamppb.New(now)

	if err := s.populateDefaultsForInstance(obj); err != nil {
		return nil, err
	}

	if err := s.storage.Create(ctx, instanceFQN, obj); err != nil {
		return nil, err
	}

	// Create cluster objects
	// If this was production, we'd probably want a transaction etc
	for clusterID, cluster := range req.GetClusters() {
		clusterFQN := instanceFQN + "/clusters/" + clusterID
		obj := proto.Clone(cluster).(*pb.Cluster)
		obj.Name = clusterFQN
		if err := s.populateDefaultsForCluster(obj); err != nil {
			return nil, err
		}
		if err := s.storage.Create(ctx, clusterFQN, obj); err != nil {
			return nil, err
		}
	}

	// Create default appProfile
	{
		defaultClusterID := ""
		for clusterID := range req.GetClusters() {
			defaultClusterID = clusterID
		}
		appProfile := &pb.AppProfile{
			Description: "Default application profile for this instance. This profile is used if you do not supply a different profile ID at connection time.",
			Name:        instanceFQN + "/appProfiles/default",
			RoutingPolicy: &pb.AppProfile_SingleClusterRouting_{
				SingleClusterRouting: &pb.AppProfile_SingleClusterRouting{
					AllowTransactionalWrites: true,
					ClusterId:                defaultClusterID,
				},
			},
			Isolation: &pb.AppProfile_StandardIsolation_{
				StandardIsolation: &pb.AppProfile_StandardIsolation{
					Priority: pb.AppProfile_PRIORITY_HIGH,
				},
			},
		}
		appProfileFQN := appProfile.Name
		if err := s.storage.Create(ctx, appProfileFQN, appProfile); err != nil {
			return nil, err
		}
	}

	// Returned request does not include clusterConfig
	for _, cluster := range originalRequest.Clusters {
		cluster.Config = nil
	}

	zone := pickZoneForInstanceOperation(mapValues(req.GetClusters()))

	prefix := fmt.Sprintf("operations/%s/locations/%s", name.String(), zone)
	metadata := &pb.CreateInstanceMetadata{
		RequestTime:     timestamppb.New(now),
		OriginalRequest: originalRequest,
	}
	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.FinishTime = timestamppb.New(time.Now())

		returnObj := proto.Clone(obj).(*pb.Instance)
		returnObj.CreateTime = nil // For some reason, not populated here
		return returnObj, nil
	})
}

func pickZoneForInstanceOperation(clusters []*pb.Cluster) string {
	zones := sets.NewString()
	for _, cluster := range clusters {
		zone := lastComponent(cluster.Location)
		zones.Insert(zone)
	}
	if zones.Len() == 0 {
		return ""
	}
	return zones.List()[0]
}

func (s *instanceAdminServer) PartialUpdateInstance(ctx context.Context, req *pb.PartialUpdateInstanceRequest) (*longrunning.Operation, error) {
	name, err := s.parseInstanceName(req.GetInstance().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Instance{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	clusters, err := s.listClustersForInstance(ctx, name)
	if err != nil {
		return nil, err
	}
	now := time.Now()
	updateMask := req.GetUpdateMask()

	for _, path := range updateMask.GetPaths() {
		switch path {
		case "display_name":
			obj.DisplayName = req.GetInstance().GetDisplayName()
		case "type":
			obj.Type = req.GetInstance().GetType()
		// case "labels":
		// 	obj.Labels = req.GetInstance().GetLabels()
		default:
			return nil, fmt.Errorf("mock does not implement update of %q", path)
		}
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	zone := pickZoneForInstanceOperation(clusters)

	prefix := fmt.Sprintf("operations/%s/locations/%s", name.String(), zone)
	metadata := &pb.UpdateInstanceMetadata{
		RequestTime:     timestamppb.New(now),
		OriginalRequest: req,
	}
	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.FinishTime = timestamppb.Now()

		return obj, nil
	})
}

func (s *instanceAdminServer) UpdateInstance(ctx context.Context, req *pb.Instance) (*pb.Instance, error) {
	name, err := s.parseInstanceName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Instance{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	obj.DisplayName = req.GetDisplayName()
	obj.Type = req.GetType()

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *instanceAdminServer) DeleteInstance(ctx context.Context, req *pb.DeleteInstanceRequest) (*emptypb.Empty, error) {
	name, err := s.parseInstanceName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.Instance{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

type instanceName struct {
	Project      *projects.ProjectData
	InstanceName string
}

func (n *instanceName) String() string {
	return "projects/" + n.Project.ID + "/instances/" + n.InstanceName
}

// parseInstanceName parses a string into a instanceName.
// The expected form is projects/<projectID>/locations/global/instances/<instanceName>
func (s *MockService) parseInstanceName(name string) (*instanceName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "instances" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &instanceName{
			Project:      project,
			InstanceName: tokens[3],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}

func (s *MockService) populateDefaultsForInstance(obj *pb.Instance) error {
	if obj.Type == pb.Instance_TYPE_UNSPECIFIED {
		obj.Type = pb.Instance_PRODUCTION
	}

	return nil
}

func lastComponent(s string) string {
	i := strings.LastIndex(s, "/")
	return s[i+1:]
}

func mapValues[K comparable, V any](m map[K]V) []V {
	var values []V
	for _, v := range m {
		values = append(values, v)
	}
	return values
}
