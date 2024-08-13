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

package mockalloydb

import (
	"context"
	"strings"

	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/alloydb/v1beta"
)

func (s *AlloyDBAdminV1) GetInstance(ctx context.Context, req *pb.GetInstanceRequest) (*pb.Instance, error) {
	name, err := s.parseInstanceName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Instance{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *AlloyDBAdminV1) CreateInstance(ctx context.Context, req *pb.CreateInstanceRequest) (*longrunning.Operation, error) {
	reqName := req.Parent + "/instances/" + req.GetInstanceId()
	name, err := s.parseInstanceName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.Instance).(*pb.Instance)
	obj.Name = fqn

	now := timestamppb.Now()
	obj.CreateTime = now

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return s.operations.NewLRO(ctx)
}

func (s *AlloyDBAdminV1) CreateSecondaryInstance(ctx context.Context, req *pb.CreateSecondaryInstanceRequest) (*longrunning.Operation, error) {
	reqName := req.Parent + "/instances/" + req.GetInstanceId()
	name, err := s.parseInstanceName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.Instance).(*pb.Instance)
	obj.Name = fqn

	now := timestamppb.Now()
	obj.CreateTime = now

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return s.operations.NewLRO(ctx)
}

func (s *AlloyDBAdminV1) UpdateInstance(ctx context.Context, req *pb.UpdateInstanceRequest) (*longrunning.Operation, error) {
	reqName := req.GetInstance().GetName()

	name, err := s.parseInstanceName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Instance{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// Required. A list of fields to be updated in this request.
	paths := req.GetUpdateMask().GetPaths()

	// TODO: Some sort of helper for fieldmask?
	for _, path := range paths {
		switch path {
		case "labels":
			obj.Labels = req.Instance.GetLabels()
		case "annotations":
			obj.Annotations = req.Instance.GetAnnotations()
		case "displayName":
			obj.DisplayName = req.Instance.GetDisplayName()
		case "gceZone":
			obj.GceZone = req.Instance.GetGceZone()
		case "databaseFlags":
			obj.DatabaseFlags = req.Instance.GetDatabaseFlags()
		case "availabilityType":
			obj.AvailabilityType = req.Instance.GetAvailabilityType()
		case "readPoolConfig":
			obj.ReadPoolConfig = req.Instance.GetReadPoolConfig()
		case "machineConfig":
			obj.MachineConfig = req.Instance.GetMachineConfig()
		case "pscInstanceConfig":
			obj.PscInstanceConfig = req.Instance.GetPscInstanceConfig()
		case "networkConfig":
			obj.NetworkConfig = req.Instance.GetNetworkConfig()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not supported by mockgcp", path)
		}
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return s.operations.NewLRO(ctx)
}

func (s *AlloyDBAdminV1) DeleteInstance(ctx context.Context, req *pb.DeleteInstanceRequest) (*longrunning.Operation, error) {
	name, err := s.parseInstanceName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	oldObj := &pb.Instance{}
	if err := s.storage.Delete(ctx, fqn, oldObj); err != nil {
		return nil, err
	}

	return s.operations.NewLRO(ctx)
}

type instanceName struct {
	Project      *projects.ProjectData
	Location     string
	ClusterName  string
	InstanceName string
}

func (n *instanceName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/clusters/" + n.ClusterName + "/instances/" + n.InstanceName
}

// parseInstanceName parses a string into an alloyDBInstanceName.
// The expected form is projects/<projectID>/locations/<region>/clusters/<cluster>/instances/<AlloyDBInstanceName>
func (s *MockService) parseInstanceName(name string) (*instanceName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "clusters" && tokens[6] == "instances" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &instanceName{
			Project:      project,
			Location:     tokens[3],
			ClusterName:  tokens[5],
			InstanceName: tokens[7],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
