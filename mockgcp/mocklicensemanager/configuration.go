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

package mocklicensemanager

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/licensemanager/apiv1/licensemanagerpb"
)

type LicenseManagerServer struct {
	*MockService
	pb.UnimplementedLicenseManagerServer
}

func (s *LicenseManagerServer) GetConfiguration(ctx context.Context, req *pb.GetConfigurationRequest) (*pb.Configuration, error) {
	name, err := s.parseConfigurationName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Configuration{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *LicenseManagerServer) CreateConfiguration(ctx context.Context, req *pb.CreateConfigurationRequest) (*longrunning.Operation, error) {
	reqName := req.Parent + "/configurations/" + req.ConfigurationId
	name, err := s.parseConfigurationName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	obj := proto.Clone(req.Configuration).(*pb.Configuration)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.State = pb.Configuration_STATE_ACTIVE

	if obj.CurrentBillingInfo != nil {
		obj.CurrentBillingInfo.StartTime = timestamppb.New(now)
	}
	if obj.NextBillingInfo != nil {
		obj.NextBillingInfo.StartTime = timestamppb.New(now)
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, prefix, nil, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *LicenseManagerServer) UpdateConfiguration(ctx context.Context, req *pb.UpdateConfigurationRequest) (*longrunning.Operation, error) {
	name, err := s.parseConfigurationName(req.Configuration.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	existing := &pb.Configuration{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		return nil, err
	}

	updated := proto.Clone(existing).(*pb.Configuration)
	updated.DisplayName = req.Configuration.DisplayName
	updated.Labels = req.Configuration.Labels
	updated.CurrentBillingInfo = req.Configuration.CurrentBillingInfo
	updated.NextBillingInfo = req.Configuration.NextBillingInfo
	updated.UpdateTime = timestamppb.New(now)

	if updated.CurrentBillingInfo != nil && updated.CurrentBillingInfo.StartTime == nil {
		updated.CurrentBillingInfo.StartTime = timestamppb.New(now)
	}
	if updated.NextBillingInfo != nil && updated.NextBillingInfo.StartTime == nil {
		updated.NextBillingInfo.StartTime = timestamppb.New(now)
	}

	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, prefix, nil, func() (proto.Message, error) {
		return updated, nil
	})
}

func (s *LicenseManagerServer) DeleteConfiguration(ctx context.Context, req *pb.DeleteConfigurationRequest) (*longrunning.Operation, error) {
	name, err := s.parseConfigurationName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	oldObj := &pb.Configuration{}
	if err := s.storage.Delete(ctx, fqn, oldObj); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, prefix, nil, func() (proto.Message, error) {
		return &emptypb.Empty{}, nil
	})
}
