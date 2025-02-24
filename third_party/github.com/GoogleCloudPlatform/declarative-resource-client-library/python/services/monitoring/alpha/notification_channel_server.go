// Copyright 2024 Google LLC. All Rights Reserved.
// 
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// 
//     http://www.apache.org/licenses/LICENSE-2.0
// 
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package server

import (
	"context"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/monitoring/alpha/monitoring_alpha_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/monitoring/alpha"
)

// NotificationChannelServer implements the gRPC interface for NotificationChannel.
type NotificationChannelServer struct{}

// ProtoToNotificationChannelVerificationStatusEnum converts a NotificationChannelVerificationStatusEnum enum from its proto representation.
func ProtoToMonitoringAlphaNotificationChannelVerificationStatusEnum(e alphapb.MonitoringAlphaNotificationChannelVerificationStatusEnum) *alpha.NotificationChannelVerificationStatusEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.MonitoringAlphaNotificationChannelVerificationStatusEnum_name[int32(e)]; ok {
		e := alpha.NotificationChannelVerificationStatusEnum(n[len("MonitoringAlphaNotificationChannelVerificationStatusEnum"):])
		return &e
	}
	return nil
}

// ProtoToNotificationChannel converts a NotificationChannel resource from its proto representation.
func ProtoToNotificationChannel(p *alphapb.MonitoringAlphaNotificationChannel) *alpha.NotificationChannel {
	obj := &alpha.NotificationChannel{
		Description:        dcl.StringOrNil(p.GetDescription()),
		DisplayName:        dcl.StringOrNil(p.GetDisplayName()),
		Enabled:            dcl.Bool(p.GetEnabled()),
		Name:               dcl.StringOrNil(p.GetName()),
		Type:               dcl.StringOrNil(p.GetType()),
		VerificationStatus: ProtoToMonitoringAlphaNotificationChannelVerificationStatusEnum(p.GetVerificationStatus()),
		Project:            dcl.StringOrNil(p.GetProject()),
	}
	return obj
}

// NotificationChannelVerificationStatusEnumToProto converts a NotificationChannelVerificationStatusEnum enum to its proto representation.
func MonitoringAlphaNotificationChannelVerificationStatusEnumToProto(e *alpha.NotificationChannelVerificationStatusEnum) alphapb.MonitoringAlphaNotificationChannelVerificationStatusEnum {
	if e == nil {
		return alphapb.MonitoringAlphaNotificationChannelVerificationStatusEnum(0)
	}
	if v, ok := alphapb.MonitoringAlphaNotificationChannelVerificationStatusEnum_value["NotificationChannelVerificationStatusEnum"+string(*e)]; ok {
		return alphapb.MonitoringAlphaNotificationChannelVerificationStatusEnum(v)
	}
	return alphapb.MonitoringAlphaNotificationChannelVerificationStatusEnum(0)
}

// NotificationChannelToProto converts a NotificationChannel resource to its proto representation.
func NotificationChannelToProto(resource *alpha.NotificationChannel) *alphapb.MonitoringAlphaNotificationChannel {
	p := &alphapb.MonitoringAlphaNotificationChannel{}
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetDisplayName(dcl.ValueOrEmptyString(resource.DisplayName))
	p.SetEnabled(dcl.ValueOrEmptyBool(resource.Enabled))
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetType(dcl.ValueOrEmptyString(resource.Type))
	p.SetVerificationStatus(MonitoringAlphaNotificationChannelVerificationStatusEnumToProto(resource.VerificationStatus))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	mLabels := make(map[string]string, len(resource.Labels))
	for k, r := range resource.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)
	mUserLabels := make(map[string]string, len(resource.UserLabels))
	for k, r := range resource.UserLabels {
		mUserLabels[k] = r
	}
	p.SetUserLabels(mUserLabels)

	return p
}

// applyNotificationChannel handles the gRPC request by passing it to the underlying NotificationChannel Apply() method.
func (s *NotificationChannelServer) applyNotificationChannel(ctx context.Context, c *alpha.Client, request *alphapb.ApplyMonitoringAlphaNotificationChannelRequest) (*alphapb.MonitoringAlphaNotificationChannel, error) {
	p := ProtoToNotificationChannel(request.GetResource())
	res, err := c.ApplyNotificationChannel(ctx, p)
	if err != nil {
		return nil, err
	}
	r := NotificationChannelToProto(res)
	return r, nil
}

// applyMonitoringAlphaNotificationChannel handles the gRPC request by passing it to the underlying NotificationChannel Apply() method.
func (s *NotificationChannelServer) ApplyMonitoringAlphaNotificationChannel(ctx context.Context, request *alphapb.ApplyMonitoringAlphaNotificationChannelRequest) (*alphapb.MonitoringAlphaNotificationChannel, error) {
	cl, err := createConfigNotificationChannel(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyNotificationChannel(ctx, cl, request)
}

// DeleteNotificationChannel handles the gRPC request by passing it to the underlying NotificationChannel Delete() method.
func (s *NotificationChannelServer) DeleteMonitoringAlphaNotificationChannel(ctx context.Context, request *alphapb.DeleteMonitoringAlphaNotificationChannelRequest) (*emptypb.Empty, error) {

	cl, err := createConfigNotificationChannel(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteNotificationChannel(ctx, ProtoToNotificationChannel(request.GetResource()))

}

// ListMonitoringAlphaNotificationChannel handles the gRPC request by passing it to the underlying NotificationChannelList() method.
func (s *NotificationChannelServer) ListMonitoringAlphaNotificationChannel(ctx context.Context, request *alphapb.ListMonitoringAlphaNotificationChannelRequest) (*alphapb.ListMonitoringAlphaNotificationChannelResponse, error) {
	cl, err := createConfigNotificationChannel(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListNotificationChannel(ctx, request.GetProject())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.MonitoringAlphaNotificationChannel
	for _, r := range resources.Items {
		rp := NotificationChannelToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListMonitoringAlphaNotificationChannelResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigNotificationChannel(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
