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
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/fields"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

type FutureReservationsV1 struct {
	*MockService
	pb.UnimplementedFutureReservationsServer
}

func (s *FutureReservationsV1) Get(ctx context.Context, req *pb.GetFutureReservationRequest) (*pb.FutureReservation, error) {
	reqName := "projects/" + req.GetProject() + "/zones/" + req.GetZone() + "/futureReservations/" + req.GetFutureReservation()
	name, err := s.parseFutureReservationName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.FutureReservation{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *FutureReservationsV1) Insert(ctx context.Context, req *pb.InsertFutureReservationRequest) (*pb.Operation, error) {
	// Simulate realGCP validation
	if req.GetFutureReservationResource().GetTimeWindow() != nil && req.GetFutureReservationResource().GetTimeWindow().GetStartTime() != "" {
		startTimeStr := req.GetFutureReservationResource().GetTimeWindow().GetStartTime()
		startTime, err := time.Parse(time.RFC3339Nano, startTimeStr)
		if err == nil {
			if startTime.Before(time.Now()) {
				return nil, status.Errorf(codes.InvalidArgument, "Invalid value for field 'resource.timeWindow.startTime': '%s'. Future reservation start time is either in the past or too early.", startTimeStr)
			}
		}
	}

	reqName := "projects/" + req.GetProject() + "/zones/" + req.GetZone() + "/futureReservations/" + req.GetFutureReservationResource().GetName()
	name, err := s.parseFutureReservationName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	id := s.generateID()

	obj := proto.CloneOf(req.GetFutureReservationResource())
	obj.CreationTimestamp = PtrTo(s.nowString())
	obj.Id = &id
	obj.SelfLink = PtrTo(BuildComputeSelfLink(ctx, name.String()))
	obj.Kind = PtrTo("compute#futureReservation")
	if obj.Status == nil {
		obj.Status = &pb.FutureReservationStatus{ProcurementStatus: PtrTo("DRAFTING")}
	}
	if obj.PlanningStatus == nil {
		obj.PlanningStatus = PtrTo("DRAFT")
	}
	obj.SelfLinkWithId = PtrTo(BuildComputeSelfLink(ctx, name.String()))
	if obj.SpecificReservationRequired == nil {
		obj.SpecificReservationRequired = PtrTo(false)
	}
	obj.Zone = PtrTo(fmt.Sprintf("https://www.googleapis.com/compute/v1/projects/%s/zones/%s", name.Project.ID, name.Zone))

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
		// Assign API default(nil) values after resource creation

		// GCP API does not have AutoDeleteAutoCreatedReservations field in the response body.
		// If it's false, it will not be populated; If it's true, AutoCreatedReservationsDeleteTime will be populated.
		obj.AutoDeleteAutoCreatedReservations = nil

		if obj.GetEnableEmergentMaintenance() == false {
			obj.EnableEmergentMaintenance = nil
		}
		if obj.ShareSettings != nil && obj.ShareSettings.GetShareType() == "LOCAL" {
			obj.ShareSettings = nil
		}
		if obj.GetSchedulingType() == "FLEXIBLE" {
			obj.SchedulingType = nil
		}

		obj.Status.ExistingMatchingUsageInfo = &pb.FutureReservationStatusExistingMatchingUsageInfo{
			Count:     PtrTo(int64(0)),
			Timestamp: PtrTo(s.nowString()),
		}
		// handle endtime calculated by duration
		if obj.TimeWindow != nil && obj.TimeWindow.Duration != nil {
			end, err := parseDuration(obj.TimeWindow.StartTime, obj.TimeWindow.Duration)
			if err != nil {
				return nil, err
			}
			obj.TimeWindow.EndTime = direct.PtrTo(end.Format(time.RFC3339Nano))
			obj.TimeWindow.Duration = nil
		}

		// handle AutoCreatedReservationsDeleteTime calculated by duration
		if obj.AutoCreatedReservationsDuration != nil {
			end, err := parseDuration(obj.TimeWindow.StartTime, obj.AutoCreatedReservationsDuration)
			if err != nil {
				return nil, err
			}
			obj.AutoCreatedReservationsDeleteTime = direct.PtrTo(end.Format(time.RFC3339Nano))
			obj.AutoCreatedReservationsDuration = nil
		}
		if err := s.storage.Update(ctx, fqn, obj); err != nil {
			return nil, err
		}
		return obj, nil
	})
}

func (s *FutureReservationsV1) Update(ctx context.Context, req *pb.UpdateFutureReservationRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/zones/" + req.GetZone() + "/futureReservations/" + req.GetFutureReservation()
	name, err := s.parseFutureReservationName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.FutureReservation{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	paths := strings.Split(req.GetUpdateMask(), ",")
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask must be provided")
	}

	if err := fields.UpdateByFieldMask(obj, req.GetFutureReservationResource(), paths); err != nil {
		return nil, fmt.Errorf("applying updates: %w", err)
	}

	// handle endtime calculated by duration
	if obj.TimeWindow != nil && obj.TimeWindow.Duration != nil {
		end, err := parseDuration(obj.TimeWindow.StartTime, obj.TimeWindow.Duration)
		if err != nil {
			return nil, err
		}
		obj.TimeWindow.EndTime = direct.PtrTo(end.Format(time.RFC3339Nano))
		obj.TimeWindow.Duration = nil
	}

	// handle AutoCreatedReservationsDeleteTime calculated by duration
	if obj.AutoCreatedReservationsDuration != nil {
		end, err := parseDuration(obj.TimeWindow.StartTime, obj.AutoCreatedReservationsDuration)
		if err != nil {
			return nil, err
		}
		obj.AutoCreatedReservationsDeleteTime = direct.PtrTo(end.Format(time.RFC3339Nano))
		obj.AutoCreatedReservationsDuration = nil
	}

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
		// If AutoDeleteAutoCreatedReservations is enabled and AutoCreatedReservationsDeleteTime/AutoCreatedReservationsDuration is not specified,
		// use EndTime as the default auto delete time.
		if obj.GetAutoDeleteAutoCreatedReservations() == true && obj.AutoCreatedReservationsDeleteTime == nil {
			obj.AutoCreatedReservationsDeleteTime = obj.TimeWindow.EndTime
		}
		// GCP API does not have AutoDeleteAutoCreatedReservations field in the response body.
		// If it's false, it will not be populated; If it's true, AutoCreatedReservationsDeleteTime will be populated.
		obj.AutoDeleteAutoCreatedReservations = nil

		// Simulate realGCP
		if obj.Description == nil {
			obj.Description = direct.PtrTo("")
		}

		if err := s.storage.Update(ctx, fqn, obj); err != nil {
			return nil, err
		}
		return obj, nil
	})
}

func (s *FutureReservationsV1) Cancel(ctx context.Context, req *pb.CancelFutureReservationRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/zones/" + req.GetZone() + "/futureReservations/" + req.GetFutureReservation()
	name, err := s.parseFutureReservationName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.FutureReservation{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// Update status to CANCELLED or similar
	if obj.Status == nil {
		obj.Status = &pb.FutureReservationStatus{}
	}
	obj.Status.ProcurementStatus = PtrTo("CANCELLED")

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("cancel"),
		User:          PtrTo("user@example.com"),
	}
	return s.startZonalLRO(ctx, name.Project.ID, name.Zone, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *FutureReservationsV1) Delete(ctx context.Context, req *pb.DeleteFutureReservationRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/zones/" + req.GetZone() + "/futureReservations/" + req.GetFutureReservation()
	name, err := s.parseFutureReservationName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.FutureReservation{}
	if err := s.storage.Delete(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("delete"),
		User:          PtrTo("user@example.com"),
	}
	return s.startZonalLRO(ctx, name.Project.ID, name.Zone, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *FutureReservationsV1) AggregatedList(ctx context.Context, req *pb.AggregatedListFutureReservationsRequest) (*pb.FutureReservationsAggregatedListResponse, error) {
	response := &pb.FutureReservationsAggregatedListResponse{}
	response.Id = PtrTo("0123456789")
	response.Kind = PtrTo("compute#futureReservationAggregatedList")
	response.SelfLink = PtrTo(BuildComputeSelfLink(ctx, "projects/"+req.GetProject()+"/aggregated/futureReservations"))

	response.Items = make(map[string]*pb.FutureReservationsScopedList)

	findKind := (&pb.FutureReservation{}).ProtoReflect().Descriptor()
	prefix := "projects/" + req.GetProject() + "/zones/"
	if err := s.storage.List(ctx, findKind, storage.ListOptions{Prefix: prefix}, func(obj proto.Message) error {
		fr := obj.(*pb.FutureReservation)

		selfLink := fr.GetSelfLink()
		tokens := strings.Split(selfLink, "/")
		zone := ""
		for i, token := range tokens {
			if token == "zones" && i+1 < len(tokens) {
				zone = "zones/" + tokens[i+1]
				break
			}
		}

		if zone != "" {
			if response.Items[zone] == nil {
				response.Items[zone] = &pb.FutureReservationsScopedList{}
			}
			response.Items[zone].FutureReservations = append(response.Items[zone].FutureReservations, fr)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}

type futureReservationName struct {
	Project *projects.ProjectData
	Zone    string
	Name    string
}

func (n *futureReservationName) String() string {
	return "projects/" + n.Project.ID + "/zones/" + n.Zone + "/futureReservations/" + n.Name
}

func (s *MockService) parseFutureReservationName(name string) (*futureReservationName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "zones" && tokens[4] == "futureReservations" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &futureReservationName{
			Project: project,
			Zone:    tokens[3],
			Name:    tokens[5],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}

func parseDuration(startTime *string, duration *pb.Duration) (time.Time, error) {
	start, err := time.Parse(time.RFC3339Nano, direct.ValueOf(startTime))
	if err != nil {
		return time.Time{}, status.Errorf(codes.InvalidArgument, "invalid start_time: %v", err)
	}

	d := time.Duration(0)
	if duration.Seconds != nil {
		d = d + time.Duration(direct.ValueOf(duration.Seconds))*time.Second
	}
	if duration.Nanos != nil {
		d = d + time.Duration(direct.ValueOf(duration.Nanos))*time.Nanosecond
	}

	return start.Add(d), nil
}
