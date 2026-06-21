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

package mockcompute

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
)

type AutoscalersV1 struct {
	*MockService
	pb.UnimplementedAutoscalersServer
}

func (s *AutoscalersV1) Get(ctx context.Context, req *pb.GetAutoscalerRequest) (*pb.Autoscaler, error) {
	reqName := "projects/" + req.GetProject() + "/zones/" + req.GetZone() + "/autoscalers/" + req.GetAutoscaler()
	name, err := s.parseZonalAutoscalerName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Autoscaler{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *AutoscalersV1) Insert(ctx context.Context, req *pb.InsertAutoscalerRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/zones/" + req.GetZone() + "/autoscalers/" + req.GetAutoscalerResource().GetName()
	name, err := s.parseZonalAutoscalerName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	id := s.generateID()

	obj := proto.Clone(req.GetAutoscalerResource()).(*pb.Autoscaler)
	obj.SelfLink = PtrTo(BuildComputeSelfLink(ctx, fqn))
	obj.CreationTimestamp = PtrTo(s.nowString())
	obj.Id = &id
	obj.Kind = PtrTo("compute#autoscaler")
	obj.Zone = PtrTo(BuildComputeSelfLink(ctx, fmt.Sprintf("projects/%s/zones/%s", name.Project.ID, name.Zone)))
	obj.Status = PtrTo("ACTIVE")

	if obj.AutoscalingPolicy != nil {
		if obj.AutoscalingPolicy.CpuUtilization != nil && obj.AutoscalingPolicy.CpuUtilization.UtilizationTarget == nil {
			obj.AutoscalingPolicy.CpuUtilization.UtilizationTarget = PtrTo(0.6)
		}
		if obj.AutoscalingPolicy.CoolDownPeriodSec == nil {
			obj.AutoscalingPolicy.CoolDownPeriodSec = PtrTo(int32(60))
		}
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("insert"),
		User:          PtrTo("user@example.com"),
	}
	return s.startZonalLRO(ctx, name.Project.ID, name.Zone, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *AutoscalersV1) Update(ctx context.Context, req *pb.UpdateAutoscalerRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/zones/" + req.GetZone() + "/autoscalers/" + req.GetAutoscalerResource().GetName()
	name, err := s.parseZonalAutoscalerName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Autoscaler{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	proto.Merge(obj, req.GetAutoscalerResource())

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("update"),
		User:          PtrTo("user@example.com"),
	}
	return s.startZonalLRO(ctx, name.Project.ID, name.Zone, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *AutoscalersV1) Patch(ctx context.Context, req *pb.PatchAutoscalerRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/zones/" + req.GetZone() + "/autoscalers/" + req.GetAutoscalerResource().GetName()
	name, err := s.parseZonalAutoscalerName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Autoscaler{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	proto.Merge(obj, req.GetAutoscalerResource())

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("patch"),
		User:          PtrTo("user@example.com"),
	}
	return s.startZonalLRO(ctx, name.Project.ID, name.Zone, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *AutoscalersV1) Delete(ctx context.Context, req *pb.DeleteAutoscalerRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/zones/" + req.GetZone() + "/autoscalers/" + req.GetAutoscaler()
	name, err := s.parseZonalAutoscalerName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.Autoscaler{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The resource '%s' was not found", fqn)
		}
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      deleted.Id,
		TargetLink:    deleted.SelfLink,
		OperationType: PtrTo("delete"),
		User:          PtrTo("user@example.com"),
	}
	return s.startZonalLRO(ctx, name.Project.ID, name.Zone, op, func() (proto.Message, error) {
		return deleted, nil
	})
}

type zonalAutoscalerName struct {
	Project *projects.ProjectData
	Zone    string
	Name    string
}

func (n *zonalAutoscalerName) String() string {
	return "projects/" + n.Project.ID + "/zones/" + n.Zone + "/autoscalers/" + n.Name
}

func (s *MockService) parseZonalAutoscalerName(name string) (*zonalAutoscalerName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "zones" && tokens[4] == "autoscalers" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &zonalAutoscalerName{
			Project: project,
			Zone:    tokens[3],
			Name:    tokens[5],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}

// RegionAutoscalersV1

type RegionAutoscalersV1 struct {
	*MockService
	pb.UnimplementedRegionAutoscalersServer
}

func (s *RegionAutoscalersV1) Get(ctx context.Context, req *pb.GetRegionAutoscalerRequest) (*pb.Autoscaler, error) {
	reqName := "projects/" + req.GetProject() + "/regions/" + req.GetRegion() + "/autoscalers/" + req.GetAutoscaler()
	name, err := s.parseRegionalAutoscalerName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Autoscaler{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *RegionAutoscalersV1) Insert(ctx context.Context, req *pb.InsertRegionAutoscalerRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/regions/" + req.GetRegion() + "/autoscalers/" + req.GetAutoscalerResource().GetName()
	name, err := s.parseRegionalAutoscalerName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	id := s.generateID()

	obj := proto.Clone(req.GetAutoscalerResource()).(*pb.Autoscaler)
	obj.SelfLink = PtrTo(BuildComputeSelfLink(ctx, fqn))
	obj.CreationTimestamp = PtrTo(s.nowString())
	obj.Id = &id
	obj.Kind = PtrTo("compute#autoscaler")
	obj.Region = PtrTo(BuildComputeSelfLink(ctx, fmt.Sprintf("projects/%s/regions/%s", name.Project.ID, name.Region)))
	obj.Status = PtrTo("ACTIVE")

	if obj.AutoscalingPolicy != nil {
		if obj.AutoscalingPolicy.CpuUtilization != nil && obj.AutoscalingPolicy.CpuUtilization.UtilizationTarget == nil {
			obj.AutoscalingPolicy.CpuUtilization.UtilizationTarget = PtrTo(0.6)
		}
		if obj.AutoscalingPolicy.CoolDownPeriodSec == nil {
			obj.AutoscalingPolicy.CoolDownPeriodSec = PtrTo(int32(60))
		}
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("insert"),
		User:          PtrTo("user@example.com"),
	}
	return s.startRegionalLRO(ctx, name.Project.ID, name.Region, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *RegionAutoscalersV1) Update(ctx context.Context, req *pb.UpdateRegionAutoscalerRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/regions/" + req.GetRegion() + "/autoscalers/" + req.GetAutoscalerResource().GetName()
	name, err := s.parseRegionalAutoscalerName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Autoscaler{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	proto.Merge(obj, req.GetAutoscalerResource())

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("update"),
		User:          PtrTo("user@example.com"),
	}
	return s.startRegionalLRO(ctx, name.Project.ID, name.Region, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *RegionAutoscalersV1) Patch(ctx context.Context, req *pb.PatchRegionAutoscalerRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/regions/" + req.GetRegion() + "/autoscalers/" + req.GetAutoscalerResource().GetName()
	name, err := s.parseRegionalAutoscalerName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Autoscaler{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	proto.Merge(obj, req.GetAutoscalerResource())

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("patch"),
		User:          PtrTo("user@example.com"),
	}
	return s.startRegionalLRO(ctx, name.Project.ID, name.Region, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *RegionAutoscalersV1) Delete(ctx context.Context, req *pb.DeleteRegionAutoscalerRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/regions/" + req.GetRegion() + "/autoscalers/" + req.GetAutoscaler()
	name, err := s.parseRegionalAutoscalerName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.Autoscaler{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The resource '%s' was not found", fqn)
		}
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      deleted.Id,
		TargetLink:    deleted.SelfLink,
		OperationType: PtrTo("delete"),
		User:          PtrTo("user@example.com"),
	}
	return s.startRegionalLRO(ctx, name.Project.ID, name.Region, op, func() (proto.Message, error) {
		return deleted, nil
	})
}

type regionalAutoscalerName struct {
	Project *projects.ProjectData
	Region  string
	Name    string
}

func (n *regionalAutoscalerName) String() string {
	return "projects/" + n.Project.ID + "/regions/" + n.Region + "/autoscalers/" + n.Name
}

func (s *MockService) parseRegionalAutoscalerName(name string) (*regionalAutoscalerName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "regions" && tokens[4] == "autoscalers" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &regionalAutoscalerName{
			Project: project,
			Region:  tokens[3],
			Name:    tokens[5],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
