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

package mockbigquerydatatransfer

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/bigquery/datatransfer/v1"
)

type dataTransferService struct {
	*MockService
	pb.UnimplementedDataTransferServiceServer
}

func (s *dataTransferService) GetTransferConfig(ctx context.Context, req *pb.GetTransferConfigRequest) (*pb.TransferConfig, error) {
	name, err := s.parseTransferConfigName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.TransferConfig{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Not found: TransferConfig %s", fqn)
		}
		return nil, err
	}

	if err := maskSensitiveFields(obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func maskSensitiveFields(obj *pb.TransferConfig) error {
	if obj.Params == nil || obj.Params.Fields == nil {
		return nil
	}
	if _, found := obj.Params.Fields["connector.authentication.oauth.clientId"]; found {
		token := make([]byte, 32)
		if _, err := rand.Read(token); err != nil {
			return err
		}
		obj.Params.Fields["connector.authentication.oauth.clientId"] = structpb.NewStringValue(base64.RawURLEncoding.EncodeToString(token))
	}
	if _, found := obj.Params.Fields["connector.authentication.oauth.clientSecret"]; found {
		token := make([]byte, 32)
		if _, err := rand.Read(token); err != nil {
			return err
		}
		obj.Params.Fields["connector.authentication.oauth.clientSecret"] = structpb.NewStringValue(base64.RawURLEncoding.EncodeToString(token))
	}
	return nil
}

func (s *dataTransferService) CreateTransferConfig(ctx context.Context, req *pb.CreateTransferConfigRequest) (*pb.TransferConfig, error) {
	if req.Parent == "" {
		return nil, status.Errorf(codes.InvalidArgument, "parent is required")
	}
	// note: req.TransferConfig.Name is ignored when creating a transfer config.
	// instead we should generated ID.
	serviceGeneratedID := fmt.Sprintf("%d", time.Now().UnixNano())
	reqName := req.Parent + "/transferConfigs/" + serviceGeneratedID
	name, err := s.parseTransferConfigName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()
	obj := proto.Clone(req.TransferConfig).(*pb.TransferConfig)
	obj.DatasetRegion = name.Location
	obj.Name = name.String()
	obj.NextRunTime = timestamppb.New(now)
	email := "user@google.com"
	obj.OwnerInfo = &pb.UserInfo{
		Email: &email,
	}
	obj.State = pb.TransferState_PENDING
	obj.UpdateTime = timestamppb.New(now)
	obj.UserId = int64(123)

	objToStore := proto.Clone(obj).(*pb.TransferConfig)
	if objToStore.EmailPreferences == nil { // match the behavior of GCP
		objToStore.EmailPreferences = &pb.EmailPreferences{}
	}

	if obj.EncryptionConfiguration == nil { // match the behavior of GCP
		obj.EncryptionConfiguration = &pb.EncryptionConfiguration{}
	}

	if err := s.storage.Create(ctx, fqn, objToStore); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *dataTransferService) UpdateTransferConfig(ctx context.Context, req *pb.UpdateTransferConfigRequest) (*pb.TransferConfig, error) {
	if req.TransferConfig == nil {
		return nil, status.Errorf(codes.InvalidArgument, "transfer_config is required")
	}
	if req.TransferConfig.Name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "transfer_config.name is required in update")
	}

	name, err := s.parseTransferConfigName(req.TransferConfig.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	obj := &pb.TransferConfig{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	paths := req.GetUpdateMask().GetPaths()
	for _, path := range paths {
		switch path {
		case "dataRefreshWindowDays":
			obj.DataRefreshWindowDays = req.TransferConfig.GetDataRefreshWindowDays()
		case "disabled":
			obj.Disabled = req.TransferConfig.GetDisabled()
		case "displayName":
			obj.DisplayName = req.TransferConfig.GetDisplayName()
		case "emailPreferences":
			obj.EmailPreferences = req.TransferConfig.GetEmailPreferences()
		case "encryptionconfiguration":
			obj.EncryptionConfiguration = req.TransferConfig.GetEncryptionConfiguration()
		case "notificationPubsubTopic":
			obj.NotificationPubsubTopic = req.TransferConfig.GetNotificationPubsubTopic()
		case "params":
			obj.Params = req.TransferConfig.GetParams()
		case "schedule":
			obj.Schedule = req.TransferConfig.GetSchedule()
		case "scheduleOptions":
			obj.ScheduleOptions = req.TransferConfig.GetScheduleOptions()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
		}
	}
	obj.UpdateTime = timestamppb.New(now)
	if obj.DataSourceId != "scheduled_query" { // match the behavior of GCP
		obj.NextRunTime = nil
	}

	objToStore := proto.Clone(obj).(*pb.TransferConfig)
	if objToStore.DataSourceId == "scheduled_query" { // match the behavior of GCP
		if objToStore.ScheduleOptions == nil {
			objToStore.ScheduleOptions = &pb.ScheduleOptions{}
		}
	}
	if err := s.storage.Update(ctx, fqn, objToStore); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *dataTransferService) DeleteTransferConfig(ctx context.Context, req *pb.DeleteTransferConfigRequest) (*empty.Empty, error) {
	if req.Name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "name is required in delete")
	}
	name, err := s.parseTransferConfigName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	oldObj := &pb.TransferConfig{}
	if err := s.storage.Delete(ctx, fqn, oldObj); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

type transferConfigName struct {
	Project    *projects.ProjectData
	Location   string
	ResourceID string
}

func (n *transferConfigName) String() string {
	return "projects/" + strconv.FormatInt(n.Project.Number, 10) + "/locations/" + n.Location + "/transferConfigs/" + n.ResourceID
}

// parseTransferConfigName parses a string into a transferConfigName.
// The expected form is projects/<projectNum>/locations/<location>/transferConfigs/<transferConfigID>
func (s *MockService) parseTransferConfigName(name string) (*transferConfigName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "transferConfigs" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &transferConfigName{
			Project:    project,
			Location:   tokens[3],
			ResourceID: tokens[5],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
