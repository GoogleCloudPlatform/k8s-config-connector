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
// proto.service: google.cloud.compute.v1.ResourcePolicies
// proto.message: google.cloud.compute.v1.ResourcePolicy

package mockcompute

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
)

type resourcePolicies struct {
	*MockService
	pb.UnimplementedResourcePoliciesServer
}

func (s *resourcePolicies) Get(ctx context.Context, req *pb.GetResourcePolicyRequest) (*pb.ResourcePolicy, error) {
	reqName := fmt.Sprintf("projects/%s/regions/%s/resourcePolicies/%s", req.GetProject(), req.GetRegion(), req.GetResourcePolicy())
	name, err := s.parseResourcePolicyName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.ResourcePolicy{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "ResourcePolicy %q not found", name)
		}
		return nil, err
	}

	return obj, nil
}

func (s *resourcePolicies) Insert(ctx context.Context, req *pb.InsertResourcePolicyRequest) (*pb.Operation, error) {
	reqName := fmt.Sprintf("projects/%s/regions/%s/resourcePolicies/%s", req.GetProject(), req.GetRegion(), req.GetResourcePolicyResource().GetName())
	name, err := s.parseResourcePolicyName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := proto.Clone(req.GetResourcePolicyResource()).(*pb.ResourcePolicy)
	obj.Id = proto.Uint64(s.generateID())
	obj.SelfLink = PtrTo(buildComputeSelfLink(ctx, fqn))
	obj.Kind = PtrTo("compute#resourcePolicy")
	obj.CreationTimestamp = PtrTo(s.nowString())
	obj.Status = PtrTo("READY")

	obj.Region = PtrTo(makeFullyQualifiedRegion(ctx, name.Project.ID, name.Region))

	s.populateDefaultsForResourcePolicy(obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		OperationType: PtrTo("insert"),
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		User:          PtrTo("user@example.com"),
	}
	return s.computeOperations.startRegionalLRO(ctx, name.Project.ID, name.Region, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *resourcePolicies) populateDefaultsForResourcePolicy(obj *pb.ResourcePolicy) {
	if snapshotSchedulePolicy := obj.GetSnapshotSchedulePolicy(); snapshotSchedulePolicy != nil {
		if retentionPolicy := snapshotSchedulePolicy.GetRetentionPolicy(); retentionPolicy != nil {
			if retentionPolicy.OnSourceDiskDelete == nil {
				retentionPolicy.OnSourceDiskDelete = PtrTo("KEEP_AUTO_SNAPSHOTS")
			}
		}
		if schedule := snapshotSchedulePolicy.GetSchedule(); schedule != nil {
			if dailySchedule := schedule.GetDailySchedule(); dailySchedule != nil {
				if dailySchedule.Duration == nil {
					dailySchedule.Duration = PtrTo("PT14400S")
				}
			}

		}
	}
}

func (s *resourcePolicies) Delete(ctx context.Context, req *pb.DeleteResourcePolicyRequest) (*pb.Operation, error) {
	reqName := fmt.Sprintf("projects/%s/regions/%s/resourcePolicies/%s", req.GetProject(), req.GetRegion(), req.GetResourcePolicy())
	name, err := s.parseResourcePolicyName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.ResourcePolicy{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		OperationType: PtrTo("delete"),
		TargetId:      deleted.Id,
		TargetLink:    deleted.SelfLink,
		User:          PtrTo("user@example.com"),
	}
	return s.computeOperations.startRegionalLRO(ctx, name.Project.ID, name.Region, op, func() (proto.Message, error) {
		return &emptypb.Empty{}, nil
	})
}

func (s *resourcePolicies) Update(ctx context.Context, req *pb.PatchResourcePolicyRequest) (*pb.Operation, error) {
	reqName := fmt.Sprintf("projects/%s/regions/%s/resourcePolicies/%s", req.GetProject(), req.GetRegion(), req.GetResourcePolicy())
	name, err := s.parseResourcePolicyName(reqName)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.ResourcePolicy{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	proto.Merge(obj, req.GetResourcePolicyResource())

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		OperationType: PtrTo("update"),
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		User:          PtrTo("user@example.com"),
	}
	return s.computeOperations.startRegionalLRO(ctx, name.Project.ID, name.Region, op, func() (proto.Message, error) {
		return obj, nil
	})
}

type resourcePolicyName struct {
	Project *projects.ProjectData
	Region  string
	Name    string
}

func (n *resourcePolicyName) String() string {
	return "projects/" + n.Project.ID + "/regions/" + n.Region + "/resourcePolicies/" + n.Name
}

// parseResourcePolicyName parses a string into a resourcePolicyName.
// The expected form is `locations/global/firewallPolicies/*`.
func (s *MockService) parseResourcePolicyName(name string) (*resourcePolicyName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "regions" && tokens[4] == "resourcePolicies" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &resourcePolicyName{
			Project: project,
			Region:  tokens[3],
			Name:    tokens[5],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
