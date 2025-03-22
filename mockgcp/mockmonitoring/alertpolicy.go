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
// proto.service: google.monitoring.v3.AlertPolicyService
// proto.message: google.monitoring.v3.AlertPolicy

package mockmonitoring

import (
	"context"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/monitoring/v3"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"github.com/golang/protobuf/ptypes/empty"
)

type AlertPolicyService struct {
	*MockService
	pb.UnimplementedAlertPolicyServiceServer
}

func (s *AlertPolicyService) GetAlertPolicy(ctx context.Context, req *pb.GetAlertPolicyRequest) (*pb.AlertPolicy, error) {
	name, err := s.parseAlertPolicyName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.AlertPolicy{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *AlertPolicyService) ListAlertPolicies(ctx context.Context, req *pb.ListAlertPoliciesRequest) (*pb.ListAlertPoliciesResponse, error) {
	name, err := s.parseAlertPolicyName(req.GetName() + "/alertPolicies/" + "placeholder")
	if err != nil {
		return nil, err
	}

	findPrefix := strings.TrimSuffix(name.String(), "placeholder")

	var alertPolicies []*pb.AlertPolicy

	findKind := (&pb.AlertPolicy{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, findKind, storage.ListOptions{Prefix: findPrefix}, func(obj proto.Message) error {
		alertPolicy := obj.(*pb.AlertPolicy)
		alertPolicies = append(alertPolicies, alertPolicy)
		return nil
	}); err != nil {
		return nil, err
	}

	return &pb.ListAlertPoliciesResponse{
		AlertPolicies: alertPolicies,
		TotalSize:     int32(len(alertPolicies)),
	}, nil
}

func (s *AlertPolicyService) CreateAlertPolicy(ctx context.Context, req *pb.CreateAlertPolicyRequest) (*pb.AlertPolicy, error) {
	now := time.Now()

	alertPolicyID := fmt.Sprintf("%d", now.UnixNano())

	reqName := req.GetName() + "/alertPolicies/" + alertPolicyID
	name, err := s.parseAlertPolicyName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.AlertPolicy).(*pb.AlertPolicy)
	obj.CreationRecord = &pb.MutationRecord{
		MutatedBy:  "user@example.com",
		MutateTime: timestamppb.New(now),
	}
	obj.MutationRecord = &pb.MutationRecord{
		MutatedBy:  "user@example.com",
		MutateTime: timestamppb.New(now),
	}
	obj.Name = fqn
	obj.Enabled = wrapperspb.Bool(true)

	for _, condition := range obj.GetConditions() {
		conditionID := fmt.Sprintf("%d", rand.Int63())
		condition.Name = fmt.Sprintf("projects/%s/alertPolicies/%s/conditions/%s", name.Project.ID, name.AlertPolicyName, conditionID)
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *AlertPolicyService) UpdateAlertPolicy(ctx context.Context, req *pb.UpdateAlertPolicyRequest) (*pb.AlertPolicy, error) {
	name, err := s.parseAlertPolicyName(req.GetAlertPolicy().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	existing := &pb.AlertPolicy{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		return nil, err
	}

	updated := proto.Clone(existing).(*pb.AlertPolicy)
	for _, path := range req.GetUpdateMask().GetPaths() {
		switch path {
		case "displayName", "display_name":
			updated.DisplayName = req.GetAlertPolicy().GetDisplayName()
		case "enabled":
			updated.Enabled = req.GetAlertPolicy().GetEnabled()
		case "conditions":
			updated.Conditions = req.GetAlertPolicy().GetConditions()
		case "documentation":
			updated.Documentation = req.GetAlertPolicy().GetDocumentation()
		case "combiner":
			updated.Combiner = req.GetAlertPolicy().GetCombiner()
		case "notificationChannels":
			updated.NotificationChannels = req.GetAlertPolicy().GetNotificationChannels()
		case "severity":
			updated.Severity = req.GetAlertPolicy().GetSeverity()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not supported by mock (full update_mask=%v)", path, req.GetUpdateMask())
		}
	}

	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}

	return updated, nil
}

func (s *AlertPolicyService) DeleteAlertPolicy(ctx context.Context, req *pb.DeleteAlertPolicyRequest) (*empty.Empty, error) {
	name, err := s.parseAlertPolicyName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.AlertPolicy{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}

type AlertPolicyName struct {
	Project         *projects.ProjectData
	AlertPolicyName string
}

func (n *AlertPolicyName) String() string {
	return fmt.Sprintf("projects/%s/alertPolicies/%s", n.Project.ID, n.AlertPolicyName)
}

// parseAlertPolicyName parses a string into a AlertPolicyName.
// The expected form is projects/[PROJECT_ID_OR_NUMBER]/AlertPolicys/[AlertPolicy_ID]
func (s *MockService) parseAlertPolicyName(name string) (*AlertPolicyName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "alertPolicies" {
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &AlertPolicyName{
			Project:         project,
			AlertPolicyName: tokens[3],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
