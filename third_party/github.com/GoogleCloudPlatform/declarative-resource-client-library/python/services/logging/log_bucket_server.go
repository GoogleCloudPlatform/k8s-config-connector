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
	loggingpb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/logging/logging_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/logging"
)

// LogBucketServer implements the gRPC interface for LogBucket.
type LogBucketServer struct{}

// ProtoToLogBucketLifecycleStateEnum converts a LogBucketLifecycleStateEnum enum from its proto representation.
func ProtoToLoggingLogBucketLifecycleStateEnum(e loggingpb.LoggingLogBucketLifecycleStateEnum) *logging.LogBucketLifecycleStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := loggingpb.LoggingLogBucketLifecycleStateEnum_name[int32(e)]; ok {
		e := logging.LogBucketLifecycleStateEnum(n[len("LoggingLogBucketLifecycleStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToLogBucket converts a LogBucket resource from its proto representation.
func ProtoToLogBucket(p *loggingpb.LoggingLogBucket) *logging.LogBucket {
	obj := &logging.LogBucket{
		Name:            dcl.StringOrNil(p.GetName()),
		Description:     dcl.StringOrNil(p.GetDescription()),
		CreateTime:      dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:      dcl.StringOrNil(p.GetUpdateTime()),
		RetentionDays:   dcl.Int64OrNil(p.GetRetentionDays()),
		Locked:          dcl.Bool(p.GetLocked()),
		LifecycleState:  ProtoToLoggingLogBucketLifecycleStateEnum(p.GetLifecycleState()),
		Parent:          dcl.StringOrNil(p.GetParent()),
		Location:        dcl.StringOrNil(p.GetLocation()),
		EnableAnalytics: dcl.Bool(p.GetEnableAnalytics()),
	}
	return obj
}

// LogBucketLifecycleStateEnumToProto converts a LogBucketLifecycleStateEnum enum to its proto representation.
func LoggingLogBucketLifecycleStateEnumToProto(e *logging.LogBucketLifecycleStateEnum) loggingpb.LoggingLogBucketLifecycleStateEnum {
	if e == nil {
		return loggingpb.LoggingLogBucketLifecycleStateEnum(0)
	}
	if v, ok := loggingpb.LoggingLogBucketLifecycleStateEnum_value["LogBucketLifecycleStateEnum"+string(*e)]; ok {
		return loggingpb.LoggingLogBucketLifecycleStateEnum(v)
	}
	return loggingpb.LoggingLogBucketLifecycleStateEnum(0)
}

// LogBucketToProto converts a LogBucket resource to its proto representation.
func LogBucketToProto(resource *logging.LogBucket) *loggingpb.LoggingLogBucket {
	p := &loggingpb.LoggingLogBucket{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetRetentionDays(dcl.ValueOrEmptyInt64(resource.RetentionDays))
	p.SetLocked(dcl.ValueOrEmptyBool(resource.Locked))
	p.SetLifecycleState(LoggingLogBucketLifecycleStateEnumToProto(resource.LifecycleState))
	p.SetParent(dcl.ValueOrEmptyString(resource.Parent))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	p.SetEnableAnalytics(dcl.ValueOrEmptyBool(resource.EnableAnalytics))

	return p
}

// applyLogBucket handles the gRPC request by passing it to the underlying LogBucket Apply() method.
func (s *LogBucketServer) applyLogBucket(ctx context.Context, c *logging.Client, request *loggingpb.ApplyLoggingLogBucketRequest) (*loggingpb.LoggingLogBucket, error) {
	p := ProtoToLogBucket(request.GetResource())
	res, err := c.ApplyLogBucket(ctx, p)
	if err != nil {
		return nil, err
	}
	r := LogBucketToProto(res)
	return r, nil
}

// applyLoggingLogBucket handles the gRPC request by passing it to the underlying LogBucket Apply() method.
func (s *LogBucketServer) ApplyLoggingLogBucket(ctx context.Context, request *loggingpb.ApplyLoggingLogBucketRequest) (*loggingpb.LoggingLogBucket, error) {
	cl, err := createConfigLogBucket(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyLogBucket(ctx, cl, request)
}

// DeleteLogBucket handles the gRPC request by passing it to the underlying LogBucket Delete() method.
func (s *LogBucketServer) DeleteLoggingLogBucket(ctx context.Context, request *loggingpb.DeleteLoggingLogBucketRequest) (*emptypb.Empty, error) {

	cl, err := createConfigLogBucket(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteLogBucket(ctx, ProtoToLogBucket(request.GetResource()))

}

// ListLoggingLogBucket handles the gRPC request by passing it to the underlying LogBucketList() method.
func (s *LogBucketServer) ListLoggingLogBucket(ctx context.Context, request *loggingpb.ListLoggingLogBucketRequest) (*loggingpb.ListLoggingLogBucketResponse, error) {
	cl, err := createConfigLogBucket(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListLogBucket(ctx, request.GetLocation(), request.GetParent())
	if err != nil {
		return nil, err
	}
	var protos []*loggingpb.LoggingLogBucket
	for _, r := range resources.Items {
		rp := LogBucketToProto(r)
		protos = append(protos, rp)
	}
	p := &loggingpb.ListLoggingLogBucketResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigLogBucket(ctx context.Context, service_account_file string) (*logging.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return logging.NewClient(conf), nil
}
