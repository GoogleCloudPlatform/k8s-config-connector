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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/cloudscheduler/beta/cloudscheduler_beta_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudscheduler/beta"
)

// JobServer implements the gRPC interface for Job.
type JobServer struct{}

// ProtoToJobAppEngineHttpTargetHttpMethodEnum converts a JobAppEngineHttpTargetHttpMethodEnum enum from its proto representation.
func ProtoToCloudschedulerBetaJobAppEngineHttpTargetHttpMethodEnum(e betapb.CloudschedulerBetaJobAppEngineHttpTargetHttpMethodEnum) *beta.JobAppEngineHttpTargetHttpMethodEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.CloudschedulerBetaJobAppEngineHttpTargetHttpMethodEnum_name[int32(e)]; ok {
		e := beta.JobAppEngineHttpTargetHttpMethodEnum(n[len("CloudschedulerBetaJobAppEngineHttpTargetHttpMethodEnum"):])
		return &e
	}
	return nil
}

// ProtoToJobHttpTargetHttpMethodEnum converts a JobHttpTargetHttpMethodEnum enum from its proto representation.
func ProtoToCloudschedulerBetaJobHttpTargetHttpMethodEnum(e betapb.CloudschedulerBetaJobHttpTargetHttpMethodEnum) *beta.JobHttpTargetHttpMethodEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.CloudschedulerBetaJobHttpTargetHttpMethodEnum_name[int32(e)]; ok {
		e := beta.JobHttpTargetHttpMethodEnum(n[len("CloudschedulerBetaJobHttpTargetHttpMethodEnum"):])
		return &e
	}
	return nil
}

// ProtoToJobStateEnum converts a JobStateEnum enum from its proto representation.
func ProtoToCloudschedulerBetaJobStateEnum(e betapb.CloudschedulerBetaJobStateEnum) *beta.JobStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.CloudschedulerBetaJobStateEnum_name[int32(e)]; ok {
		e := beta.JobStateEnum(n[len("CloudschedulerBetaJobStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToJobPubsubTarget converts a JobPubsubTarget object from its proto representation.
func ProtoToCloudschedulerBetaJobPubsubTarget(p *betapb.CloudschedulerBetaJobPubsubTarget) *beta.JobPubsubTarget {
	if p == nil {
		return nil
	}
	obj := &beta.JobPubsubTarget{
		TopicName: dcl.StringOrNil(p.GetTopicName()),
		Data:      dcl.StringOrNil(p.GetData()),
	}
	return obj
}

// ProtoToJobAppEngineHttpTarget converts a JobAppEngineHttpTarget object from its proto representation.
func ProtoToCloudschedulerBetaJobAppEngineHttpTarget(p *betapb.CloudschedulerBetaJobAppEngineHttpTarget) *beta.JobAppEngineHttpTarget {
	if p == nil {
		return nil
	}
	obj := &beta.JobAppEngineHttpTarget{
		HttpMethod:       ProtoToCloudschedulerBetaJobAppEngineHttpTargetHttpMethodEnum(p.GetHttpMethod()),
		AppEngineRouting: ProtoToCloudschedulerBetaJobAppEngineHttpTargetAppEngineRouting(p.GetAppEngineRouting()),
		RelativeUri:      dcl.StringOrNil(p.GetRelativeUri()),
		Body:             dcl.StringOrNil(p.GetBody()),
	}
	return obj
}

// ProtoToJobAppEngineHttpTargetAppEngineRouting converts a JobAppEngineHttpTargetAppEngineRouting object from its proto representation.
func ProtoToCloudschedulerBetaJobAppEngineHttpTargetAppEngineRouting(p *betapb.CloudschedulerBetaJobAppEngineHttpTargetAppEngineRouting) *beta.JobAppEngineHttpTargetAppEngineRouting {
	if p == nil {
		return nil
	}
	obj := &beta.JobAppEngineHttpTargetAppEngineRouting{
		Service:  dcl.StringOrNil(p.GetService()),
		Version:  dcl.StringOrNil(p.GetVersion()),
		Instance: dcl.StringOrNil(p.GetInstance()),
		Host:     dcl.StringOrNil(p.GetHost()),
	}
	return obj
}

// ProtoToJobHttpTarget converts a JobHttpTarget object from its proto representation.
func ProtoToCloudschedulerBetaJobHttpTarget(p *betapb.CloudschedulerBetaJobHttpTarget) *beta.JobHttpTarget {
	if p == nil {
		return nil
	}
	obj := &beta.JobHttpTarget{
		Uri:        dcl.StringOrNil(p.GetUri()),
		HttpMethod: ProtoToCloudschedulerBetaJobHttpTargetHttpMethodEnum(p.GetHttpMethod()),
		Body:       dcl.StringOrNil(p.GetBody()),
		OAuthToken: ProtoToCloudschedulerBetaJobHttpTargetOAuthToken(p.GetOauthToken()),
		OidcToken:  ProtoToCloudschedulerBetaJobHttpTargetOidcToken(p.GetOidcToken()),
	}
	return obj
}

// ProtoToJobHttpTargetOAuthToken converts a JobHttpTargetOAuthToken object from its proto representation.
func ProtoToCloudschedulerBetaJobHttpTargetOAuthToken(p *betapb.CloudschedulerBetaJobHttpTargetOAuthToken) *beta.JobHttpTargetOAuthToken {
	if p == nil {
		return nil
	}
	obj := &beta.JobHttpTargetOAuthToken{
		ServiceAccountEmail: dcl.StringOrNil(p.GetServiceAccountEmail()),
		Scope:               dcl.StringOrNil(p.GetScope()),
	}
	return obj
}

// ProtoToJobHttpTargetOidcToken converts a JobHttpTargetOidcToken object from its proto representation.
func ProtoToCloudschedulerBetaJobHttpTargetOidcToken(p *betapb.CloudschedulerBetaJobHttpTargetOidcToken) *beta.JobHttpTargetOidcToken {
	if p == nil {
		return nil
	}
	obj := &beta.JobHttpTargetOidcToken{
		ServiceAccountEmail: dcl.StringOrNil(p.GetServiceAccountEmail()),
		Audience:            dcl.StringOrNil(p.GetAudience()),
	}
	return obj
}

// ProtoToJobStatus converts a JobStatus object from its proto representation.
func ProtoToCloudschedulerBetaJobStatus(p *betapb.CloudschedulerBetaJobStatus) *beta.JobStatus {
	if p == nil {
		return nil
	}
	obj := &beta.JobStatus{
		Code:    dcl.Int64OrNil(p.GetCode()),
		Message: dcl.StringOrNil(p.GetMessage()),
	}
	for _, r := range p.GetDetails() {
		obj.Details = append(obj.Details, *ProtoToCloudschedulerBetaJobStatusDetails(r))
	}
	return obj
}

// ProtoToJobStatusDetails converts a JobStatusDetails object from its proto representation.
func ProtoToCloudschedulerBetaJobStatusDetails(p *betapb.CloudschedulerBetaJobStatusDetails) *beta.JobStatusDetails {
	if p == nil {
		return nil
	}
	obj := &beta.JobStatusDetails{
		TypeUrl: dcl.StringOrNil(p.GetTypeUrl()),
		Value:   dcl.StringOrNil(p.GetValue()),
	}
	return obj
}

// ProtoToJobRetryConfig converts a JobRetryConfig object from its proto representation.
func ProtoToCloudschedulerBetaJobRetryConfig(p *betapb.CloudschedulerBetaJobRetryConfig) *beta.JobRetryConfig {
	if p == nil {
		return nil
	}
	obj := &beta.JobRetryConfig{
		RetryCount:         dcl.Int64OrNil(p.GetRetryCount()),
		MaxRetryDuration:   dcl.StringOrNil(p.GetMaxRetryDuration()),
		MinBackoffDuration: dcl.StringOrNil(p.GetMinBackoffDuration()),
		MaxBackoffDuration: dcl.StringOrNil(p.GetMaxBackoffDuration()),
		MaxDoublings:       dcl.Int64OrNil(p.GetMaxDoublings()),
	}
	return obj
}

// ProtoToJob converts a Job resource from its proto representation.
func ProtoToJob(p *betapb.CloudschedulerBetaJob) *beta.Job {
	obj := &beta.Job{
		Name:                dcl.StringOrNil(p.GetName()),
		Description:         dcl.StringOrNil(p.GetDescription()),
		PubsubTarget:        ProtoToCloudschedulerBetaJobPubsubTarget(p.GetPubsubTarget()),
		AppEngineHttpTarget: ProtoToCloudschedulerBetaJobAppEngineHttpTarget(p.GetAppEngineHttpTarget()),
		HttpTarget:          ProtoToCloudschedulerBetaJobHttpTarget(p.GetHttpTarget()),
		Schedule:            dcl.StringOrNil(p.GetSchedule()),
		TimeZone:            dcl.StringOrNil(p.GetTimeZone()),
		UserUpdateTime:      dcl.StringOrNil(p.GetUserUpdateTime()),
		State:               ProtoToCloudschedulerBetaJobStateEnum(p.GetState()),
		Status:              ProtoToCloudschedulerBetaJobStatus(p.GetStatus()),
		ScheduleTime:        dcl.StringOrNil(p.GetScheduleTime()),
		LastAttemptTime:     dcl.StringOrNil(p.GetLastAttemptTime()),
		RetryConfig:         ProtoToCloudschedulerBetaJobRetryConfig(p.GetRetryConfig()),
		AttemptDeadline:     dcl.StringOrNil(p.GetAttemptDeadline()),
		Project:             dcl.StringOrNil(p.GetProject()),
		Location:            dcl.StringOrNil(p.GetLocation()),
	}
	return obj
}

// JobAppEngineHttpTargetHttpMethodEnumToProto converts a JobAppEngineHttpTargetHttpMethodEnum enum to its proto representation.
func CloudschedulerBetaJobAppEngineHttpTargetHttpMethodEnumToProto(e *beta.JobAppEngineHttpTargetHttpMethodEnum) betapb.CloudschedulerBetaJobAppEngineHttpTargetHttpMethodEnum {
	if e == nil {
		return betapb.CloudschedulerBetaJobAppEngineHttpTargetHttpMethodEnum(0)
	}
	if v, ok := betapb.CloudschedulerBetaJobAppEngineHttpTargetHttpMethodEnum_value["JobAppEngineHttpTargetHttpMethodEnum"+string(*e)]; ok {
		return betapb.CloudschedulerBetaJobAppEngineHttpTargetHttpMethodEnum(v)
	}
	return betapb.CloudschedulerBetaJobAppEngineHttpTargetHttpMethodEnum(0)
}

// JobHttpTargetHttpMethodEnumToProto converts a JobHttpTargetHttpMethodEnum enum to its proto representation.
func CloudschedulerBetaJobHttpTargetHttpMethodEnumToProto(e *beta.JobHttpTargetHttpMethodEnum) betapb.CloudschedulerBetaJobHttpTargetHttpMethodEnum {
	if e == nil {
		return betapb.CloudschedulerBetaJobHttpTargetHttpMethodEnum(0)
	}
	if v, ok := betapb.CloudschedulerBetaJobHttpTargetHttpMethodEnum_value["JobHttpTargetHttpMethodEnum"+string(*e)]; ok {
		return betapb.CloudschedulerBetaJobHttpTargetHttpMethodEnum(v)
	}
	return betapb.CloudschedulerBetaJobHttpTargetHttpMethodEnum(0)
}

// JobStateEnumToProto converts a JobStateEnum enum to its proto representation.
func CloudschedulerBetaJobStateEnumToProto(e *beta.JobStateEnum) betapb.CloudschedulerBetaJobStateEnum {
	if e == nil {
		return betapb.CloudschedulerBetaJobStateEnum(0)
	}
	if v, ok := betapb.CloudschedulerBetaJobStateEnum_value["JobStateEnum"+string(*e)]; ok {
		return betapb.CloudschedulerBetaJobStateEnum(v)
	}
	return betapb.CloudschedulerBetaJobStateEnum(0)
}

// JobPubsubTargetToProto converts a JobPubsubTarget object to its proto representation.
func CloudschedulerBetaJobPubsubTargetToProto(o *beta.JobPubsubTarget) *betapb.CloudschedulerBetaJobPubsubTarget {
	if o == nil {
		return nil
	}
	p := &betapb.CloudschedulerBetaJobPubsubTarget{}
	p.SetTopicName(dcl.ValueOrEmptyString(o.TopicName))
	p.SetData(dcl.ValueOrEmptyString(o.Data))
	mAttributes := make(map[string]string, len(o.Attributes))
	for k, r := range o.Attributes {
		mAttributes[k] = r
	}
	p.SetAttributes(mAttributes)
	return p
}

// JobAppEngineHttpTargetToProto converts a JobAppEngineHttpTarget object to its proto representation.
func CloudschedulerBetaJobAppEngineHttpTargetToProto(o *beta.JobAppEngineHttpTarget) *betapb.CloudschedulerBetaJobAppEngineHttpTarget {
	if o == nil {
		return nil
	}
	p := &betapb.CloudschedulerBetaJobAppEngineHttpTarget{}
	p.SetHttpMethod(CloudschedulerBetaJobAppEngineHttpTargetHttpMethodEnumToProto(o.HttpMethod))
	p.SetAppEngineRouting(CloudschedulerBetaJobAppEngineHttpTargetAppEngineRoutingToProto(o.AppEngineRouting))
	p.SetRelativeUri(dcl.ValueOrEmptyString(o.RelativeUri))
	p.SetBody(dcl.ValueOrEmptyString(o.Body))
	mHeaders := make(map[string]string, len(o.Headers))
	for k, r := range o.Headers {
		mHeaders[k] = r
	}
	p.SetHeaders(mHeaders)
	return p
}

// JobAppEngineHttpTargetAppEngineRoutingToProto converts a JobAppEngineHttpTargetAppEngineRouting object to its proto representation.
func CloudschedulerBetaJobAppEngineHttpTargetAppEngineRoutingToProto(o *beta.JobAppEngineHttpTargetAppEngineRouting) *betapb.CloudschedulerBetaJobAppEngineHttpTargetAppEngineRouting {
	if o == nil {
		return nil
	}
	p := &betapb.CloudschedulerBetaJobAppEngineHttpTargetAppEngineRouting{}
	p.SetService(dcl.ValueOrEmptyString(o.Service))
	p.SetVersion(dcl.ValueOrEmptyString(o.Version))
	p.SetInstance(dcl.ValueOrEmptyString(o.Instance))
	p.SetHost(dcl.ValueOrEmptyString(o.Host))
	return p
}

// JobHttpTargetToProto converts a JobHttpTarget object to its proto representation.
func CloudschedulerBetaJobHttpTargetToProto(o *beta.JobHttpTarget) *betapb.CloudschedulerBetaJobHttpTarget {
	if o == nil {
		return nil
	}
	p := &betapb.CloudschedulerBetaJobHttpTarget{}
	p.SetUri(dcl.ValueOrEmptyString(o.Uri))
	p.SetHttpMethod(CloudschedulerBetaJobHttpTargetHttpMethodEnumToProto(o.HttpMethod))
	p.SetBody(dcl.ValueOrEmptyString(o.Body))
	p.SetOauthToken(CloudschedulerBetaJobHttpTargetOAuthTokenToProto(o.OAuthToken))
	p.SetOidcToken(CloudschedulerBetaJobHttpTargetOidcTokenToProto(o.OidcToken))
	mHeaders := make(map[string]string, len(o.Headers))
	for k, r := range o.Headers {
		mHeaders[k] = r
	}
	p.SetHeaders(mHeaders)
	return p
}

// JobHttpTargetOAuthTokenToProto converts a JobHttpTargetOAuthToken object to its proto representation.
func CloudschedulerBetaJobHttpTargetOAuthTokenToProto(o *beta.JobHttpTargetOAuthToken) *betapb.CloudschedulerBetaJobHttpTargetOAuthToken {
	if o == nil {
		return nil
	}
	p := &betapb.CloudschedulerBetaJobHttpTargetOAuthToken{}
	p.SetServiceAccountEmail(dcl.ValueOrEmptyString(o.ServiceAccountEmail))
	p.SetScope(dcl.ValueOrEmptyString(o.Scope))
	return p
}

// JobHttpTargetOidcTokenToProto converts a JobHttpTargetOidcToken object to its proto representation.
func CloudschedulerBetaJobHttpTargetOidcTokenToProto(o *beta.JobHttpTargetOidcToken) *betapb.CloudschedulerBetaJobHttpTargetOidcToken {
	if o == nil {
		return nil
	}
	p := &betapb.CloudschedulerBetaJobHttpTargetOidcToken{}
	p.SetServiceAccountEmail(dcl.ValueOrEmptyString(o.ServiceAccountEmail))
	p.SetAudience(dcl.ValueOrEmptyString(o.Audience))
	return p
}

// JobStatusToProto converts a JobStatus object to its proto representation.
func CloudschedulerBetaJobStatusToProto(o *beta.JobStatus) *betapb.CloudschedulerBetaJobStatus {
	if o == nil {
		return nil
	}
	p := &betapb.CloudschedulerBetaJobStatus{}
	p.SetCode(dcl.ValueOrEmptyInt64(o.Code))
	p.SetMessage(dcl.ValueOrEmptyString(o.Message))
	sDetails := make([]*betapb.CloudschedulerBetaJobStatusDetails, len(o.Details))
	for i, r := range o.Details {
		sDetails[i] = CloudschedulerBetaJobStatusDetailsToProto(&r)
	}
	p.SetDetails(sDetails)
	return p
}

// JobStatusDetailsToProto converts a JobStatusDetails object to its proto representation.
func CloudschedulerBetaJobStatusDetailsToProto(o *beta.JobStatusDetails) *betapb.CloudschedulerBetaJobStatusDetails {
	if o == nil {
		return nil
	}
	p := &betapb.CloudschedulerBetaJobStatusDetails{}
	p.SetTypeUrl(dcl.ValueOrEmptyString(o.TypeUrl))
	p.SetValue(dcl.ValueOrEmptyString(o.Value))
	return p
}

// JobRetryConfigToProto converts a JobRetryConfig object to its proto representation.
func CloudschedulerBetaJobRetryConfigToProto(o *beta.JobRetryConfig) *betapb.CloudschedulerBetaJobRetryConfig {
	if o == nil {
		return nil
	}
	p := &betapb.CloudschedulerBetaJobRetryConfig{}
	p.SetRetryCount(dcl.ValueOrEmptyInt64(o.RetryCount))
	p.SetMaxRetryDuration(dcl.ValueOrEmptyString(o.MaxRetryDuration))
	p.SetMinBackoffDuration(dcl.ValueOrEmptyString(o.MinBackoffDuration))
	p.SetMaxBackoffDuration(dcl.ValueOrEmptyString(o.MaxBackoffDuration))
	p.SetMaxDoublings(dcl.ValueOrEmptyInt64(o.MaxDoublings))
	return p
}

// JobToProto converts a Job resource to its proto representation.
func JobToProto(resource *beta.Job) *betapb.CloudschedulerBetaJob {
	p := &betapb.CloudschedulerBetaJob{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetPubsubTarget(CloudschedulerBetaJobPubsubTargetToProto(resource.PubsubTarget))
	p.SetAppEngineHttpTarget(CloudschedulerBetaJobAppEngineHttpTargetToProto(resource.AppEngineHttpTarget))
	p.SetHttpTarget(CloudschedulerBetaJobHttpTargetToProto(resource.HttpTarget))
	p.SetSchedule(dcl.ValueOrEmptyString(resource.Schedule))
	p.SetTimeZone(dcl.ValueOrEmptyString(resource.TimeZone))
	p.SetUserUpdateTime(dcl.ValueOrEmptyString(resource.UserUpdateTime))
	p.SetState(CloudschedulerBetaJobStateEnumToProto(resource.State))
	p.SetStatus(CloudschedulerBetaJobStatusToProto(resource.Status))
	p.SetScheduleTime(dcl.ValueOrEmptyString(resource.ScheduleTime))
	p.SetLastAttemptTime(dcl.ValueOrEmptyString(resource.LastAttemptTime))
	p.SetRetryConfig(CloudschedulerBetaJobRetryConfigToProto(resource.RetryConfig))
	p.SetAttemptDeadline(dcl.ValueOrEmptyString(resource.AttemptDeadline))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))

	return p
}

// applyJob handles the gRPC request by passing it to the underlying Job Apply() method.
func (s *JobServer) applyJob(ctx context.Context, c *beta.Client, request *betapb.ApplyCloudschedulerBetaJobRequest) (*betapb.CloudschedulerBetaJob, error) {
	p := ProtoToJob(request.GetResource())
	res, err := c.ApplyJob(ctx, p)
	if err != nil {
		return nil, err
	}
	r := JobToProto(res)
	return r, nil
}

// applyCloudschedulerBetaJob handles the gRPC request by passing it to the underlying Job Apply() method.
func (s *JobServer) ApplyCloudschedulerBetaJob(ctx context.Context, request *betapb.ApplyCloudschedulerBetaJobRequest) (*betapb.CloudschedulerBetaJob, error) {
	cl, err := createConfigJob(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyJob(ctx, cl, request)
}

// DeleteJob handles the gRPC request by passing it to the underlying Job Delete() method.
func (s *JobServer) DeleteCloudschedulerBetaJob(ctx context.Context, request *betapb.DeleteCloudschedulerBetaJobRequest) (*emptypb.Empty, error) {

	cl, err := createConfigJob(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteJob(ctx, ProtoToJob(request.GetResource()))

}

// ListCloudschedulerBetaJob handles the gRPC request by passing it to the underlying JobList() method.
func (s *JobServer) ListCloudschedulerBetaJob(ctx context.Context, request *betapb.ListCloudschedulerBetaJobRequest) (*betapb.ListCloudschedulerBetaJobResponse, error) {
	cl, err := createConfigJob(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListJob(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.CloudschedulerBetaJob
	for _, r := range resources.Items {
		rp := JobToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListCloudschedulerBetaJobResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigJob(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
