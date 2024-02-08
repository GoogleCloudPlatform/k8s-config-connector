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

package mockmonitoring

import (
	"context"
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
	apierrors "k8s.io/apimachinery/pkg/api/errors"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/monitoring/v3"
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
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "alertPolicy %q not found", name)
		} else {
			return nil, status.Errorf(codes.Internal, "error reading alertPolicy: %v", err)
		}
	}

	return obj, nil
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
	obj.Name = fqn

	creationRecord := &pb.MutationRecord{
		MutateTime: timestamppb.New(now),
		MutatedBy:  common.GetUser(ctx),
	}

	obj.CreationRecord = creationRecord
	obj.MutationRecord = creationRecord

	for _, condition := range obj.Conditions {
		n := newID()
		conditionID := fmt.Sprintf("%d", n)
		condition.Name = fqn + "/conditions/" + conditionID
	}
	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, status.Errorf(codes.Internal, "error creating alertPolicy: %v", err)
	}

	return obj, nil
}

func newID() uint64 {
	b := make([]byte, 8)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return binary.LittleEndian.Uint64(b)
}

func (s *AlertPolicyService) UpdateAlertPolicy(ctx context.Context, req *pb.UpdateAlertPolicyRequest) (*pb.AlertPolicy, error) {
	name, err := s.parseAlertPolicyName(req.GetAlertPolicy().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	existing := &pb.AlertPolicy{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "alertPolicy %q not found", name)
		} else {
			return nil, status.Errorf(codes.Internal, "error reading alertPolicy: %v", err)
		}
	}

	now := time.Now()

	updated := proto.Clone(existing).(*pb.AlertPolicy)

	for _, path := range req.GetUpdateMask().GetPaths() {
		switch path {
		case "displayName":
			updated.DisplayName = req.GetAlertPolicy().GetDisplayName()

		case "combiner":
			updated.Combiner = req.GetAlertPolicy().GetCombiner()

		case "conditions":
			updated.Conditions = req.GetAlertPolicy().GetConditions()

		case "documentation":
			updated.Documentation = req.GetAlertPolicy().GetDocumentation()

		case "notificationChannels":
			updated.NotificationChannels = req.GetAlertPolicy().GetNotificationChannels()

		case "enabled":
			updated.Enabled = req.GetAlertPolicy().GetEnabled()

		default:
			return nil, status.Errorf(codes.InvalidArgument, "unknown updateMask %q", path)
		}
	}

	updated.MutationRecord = &pb.MutationRecord{
		MutateTime: timestamppb.New(now),
		MutatedBy:  common.GetUser(ctx),
	}

	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, status.Errorf(codes.Internal, "error updating alertPolicy: %v", err)
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
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "alertPolicy %q not found", name)
		} else {
			return nil, status.Errorf(codes.Internal, "error deleting alertPolicy: %v", err)
		}
	}

	return &empty.Empty{}, nil
}

type alertPolicyName struct {
	Project *projects.ProjectData
	Name    string
}

func (n *alertPolicyName) String() string {
	return "projects/" + n.Project.ID + "/alertPolicies/" + n.Name
}

// parseAlertPolicyName parses a string into a alertPolicyName.
// The expected form is projects/[PROJECT_ID_OR_NUMBER]/alertPolicies/[ALERT_POLICY_ID]
func (s *MockService) parseAlertPolicyName(name string) (*alertPolicyName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "alertPolicies" {
		project, err := s.projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &alertPolicyName{
			Project: project,
			Name:    tokens[3],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
