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
	cloudschedulerpb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/cloudscheduler/cloudscheduler_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudscheduler"
)

// JobServer implements the gRPC interface for Job.
type JobServer struct{}

// ProtoToJobAppEngineHttpTargetHttpMethodEnum converts a JobAppEngineHttpTargetHttpMethodEnum enum from its proto representation.
func ProtoToCloudschedulerJobAppEngineHttpTargetHttpMethodEnum(e cloudschedulerpb.CloudschedulerJobAppEngineHttpTargetHttpMethodEnum) *cloudscheduler.JobAppEngineHttpTargetHttpMethodEnum {
	if e == 0 {
		return nil
	}
	if n, ok := cloudschedulerpb.CloudschedulerJobAppEngineHttpTargetHttpMethodEnum_name[int32(e)]; ok {
		e := cloudscheduler.JobAppEngineHttpTargetHttpMethodEnum(n[len("CloudschedulerJobAppEngineHttpTargetHttpMethodEnum"):])
		return &e
	}
	return nil
}

// ProtoToJobHttpTargetHttpMethodEnum converts a JobHttpTargetHttpMethodEnum enum from its proto representation.
func ProtoToCloudschedulerJobHttpTargetHttpMethodEnum(e cloudschedulerpb.CloudschedulerJobHttpTargetHttpMethodEnum) *cloudscheduler.JobHttpTargetHttpMethodEnum {
	if e == 0 {
		return nil
	}
	if n, ok := cloudschedulerpb.CloudschedulerJobHttpTargetHttpMethodEnum_name[int32(e)]; ok {
		e := cloudscheduler.JobHttpTargetHttpMethodEnum(n[len("CloudschedulerJobHttpTargetHttpMethodEnum"):])
		return &e
	}
	return nil
}

// ProtoToJobStateEnum converts a JobStateEnum enum from its proto representation.
func ProtoToCloudschedulerJobStateEnum(e cloudschedulerpb.CloudschedulerJobStateEnum) *cloudscheduler.JobStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := cloudschedulerpb.CloudschedulerJobStateEnum_name[int32(e)]; ok {
		e := cloudscheduler.JobStateEnum(n[len("CloudschedulerJobStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToJobPubsubTarget converts a JobPubsubTarget object from its proto representation.
func ProtoToCloudschedulerJobPubsubTarget(p *cloudschedulerpb.CloudschedulerJobPubsubTarget) *cloudscheduler.JobPubsubTarget {
	if p == nil {
		return nil
	}
	obj := &cloudscheduler.JobPubsubTarget{
		TopicName: dcl.StringOrNil(p.GetTopicName()),
		Data:      dcl.StringOrNil(p.GetData()),
	}
	return obj
}

// ProtoToJobAppEngineHttpTarget converts a JobAppEngineHttpTarget object from its proto representation.
func ProtoToCloudschedulerJobAppEngineHttpTarget(p *cloudschedulerpb.CloudschedulerJobAppEngineHttpTarget) *cloudscheduler.JobAppEngineHttpTarget {
	if p == nil {
		return nil
	}
	obj := &cloudscheduler.JobAppEngineHttpTarget{
		HttpMethod:       ProtoToCloudschedulerJobAppEngineHttpTargetHttpMethodEnum(p.GetHttpMethod()),
		AppEngineRouting: ProtoToCloudschedulerJobAppEngineHttpTargetAppEngineRouting(p.GetAppEngineRouting()),
		RelativeUri:      dcl.StringOrNil(p.GetRelativeUri()),
		Body:             dcl.StringOrNil(p.GetBody()),
	}
	return obj
}

// ProtoToJobAppEngineHttpTargetAppEngineRouting converts a JobAppEngineHttpTargetAppEngineRouting object from its proto representation.
func ProtoToCloudschedulerJobAppEngineHttpTargetAppEngineRouting(p *cloudschedulerpb.CloudschedulerJobAppEngineHttpTargetAppEngineRouting) *cloudscheduler.JobAppEngineHttpTargetAppEngineRouting {
	if p == nil {
		return nil
	}
	obj := &cloudscheduler.JobAppEngineHttpTargetAppEngineRouting{
		Service:  dcl.StringOrNil(p.GetService()),
		Version:  dcl.StringOrNil(p.GetVersion()),
		Instance: dcl.StringOrNil(p.GetInstance()),
		Host:     dcl.StringOrNil(p.GetHost()),
	}
	return obj
}

// ProtoToJobHttpTarget converts a JobHttpTarget object from its proto representation.
func ProtoToCloudschedulerJobHttpTarget(p *cloudschedulerpb.CloudschedulerJobHttpTarget) *cloudscheduler.JobHttpTarget {
	if p == nil {
		return nil
	}
	obj := &cloudscheduler.JobHttpTarget{
		Uri:        dcl.StringOrNil(p.GetUri()),
		HttpMethod: ProtoToCloudschedulerJobHttpTargetHttpMethodEnum(p.GetHttpMethod()),
		Body:       dcl.StringOrNil(p.GetBody()),
		OAuthToken: ProtoToCloudschedulerJobHttpTargetOAuthToken(p.GetOauthToken()),
		OidcToken:  ProtoToCloudschedulerJobHttpTargetOidcToken(p.GetOidcToken()),
	}
	return obj
}

// ProtoToJobHttpTargetOAuthToken converts a JobHttpTargetOAuthToken object from its proto representation.
func ProtoToCloudschedulerJobHttpTargetOAuthToken(p *cloudschedulerpb.CloudschedulerJobHttpTargetOAuthToken) *cloudscheduler.JobHttpTargetOAuthToken {
	if p == nil {
		return nil
	}
	obj := &cloudscheduler.JobHttpTargetOAuthToken{
		ServiceAccountEmail: dcl.StringOrNil(p.GetServiceAccountEmail()),
		Scope:               dcl.StringOrNil(p.GetScope()),
	}
	return obj
}

// ProtoToJobHttpTargetOidcToken converts a JobHttpTargetOidcToken object from its proto representation.
func ProtoToCloudschedulerJobHttpTargetOidcToken(p *cloudschedulerpb.CloudschedulerJobHttpTargetOidcToken) *cloudscheduler.JobHttpTargetOidcToken {
	if p == nil {
		return nil
	}
	obj := &cloudscheduler.JobHttpTargetOidcToken{
		ServiceAccountEmail: dcl.StringOrNil(p.GetServiceAccountEmail()),
		Audience:            dcl.StringOrNil(p.GetAudience()),
	}
	return obj
}

// ProtoToJobStatus converts a JobStatus object from its proto representation.
func ProtoToCloudschedulerJobStatus(p *cloudschedulerpb.CloudschedulerJobStatus) *cloudscheduler.JobStatus {
	if p == nil {
		return nil
	}
	obj := &cloudscheduler.JobStatus{
		Code:    dcl.Int64OrNil(p.GetCode()),
		Message: dcl.StringOrNil(p.GetMessage()),
	}
	for _, r := range p.GetDetails() {
		obj.Details = append(obj.Details, *ProtoToCloudschedulerJobStatusDetails(r))
	}
	return obj
}

// ProtoToJobStatusDetails converts a JobStatusDetails object from its proto representation.
func ProtoToCloudschedulerJobStatusDetails(p *cloudschedulerpb.CloudschedulerJobStatusDetails) *cloudscheduler.JobStatusDetails {
	if p == nil {
		return nil
	}
	obj := &cloudscheduler.JobStatusDetails{
		TypeUrl: dcl.StringOrNil(p.GetTypeUrl()),
		Value:   dcl.StringOrNil(p.GetValue()),
	}
	return obj
}

// ProtoToJobRetryConfig converts a JobRetryConfig object from its proto representation.
func ProtoToCloudschedulerJobRetryConfig(p *cloudschedulerpb.CloudschedulerJobRetryConfig) *cloudscheduler.JobRetryConfig {
	if p == nil {
		return nil
	}
	obj := &cloudscheduler.JobRetryConfig{
		RetryCount:         dcl.Int64OrNil(p.GetRetryCount()),
		MaxRetryDuration:   dcl.StringOrNil(p.GetMaxRetryDuration()),
		MinBackoffDuration: dcl.StringOrNil(p.GetMinBackoffDuration()),
		MaxBackoffDuration: dcl.StringOrNil(p.GetMaxBackoffDuration()),
		MaxDoublings:       dcl.Int64OrNil(p.GetMaxDoublings()),
	}
	return obj
}

// ProtoToJob converts a Job resource from its proto representation.
func ProtoToJob(p *cloudschedulerpb.CloudschedulerJob) *cloudscheduler.Job {
	obj := &cloudscheduler.Job{
		Name:                dcl.StringOrNil(p.GetName()),
		Description:         dcl.StringOrNil(p.GetDescription()),
		PubsubTarget:        ProtoToCloudschedulerJobPubsubTarget(p.GetPubsubTarget()),
		AppEngineHttpTarget: ProtoToCloudschedulerJobAppEngineHttpTarget(p.GetAppEngineHttpTarget()),
		HttpTarget:          ProtoToCloudschedulerJobHttpTarget(p.GetHttpTarget()),
		Schedule:            dcl.StringOrNil(p.GetSchedule()),
		TimeZone:            dcl.StringOrNil(p.GetTimeZone()),
		UserUpdateTime:      dcl.StringOrNil(p.GetUserUpdateTime()),
		State:               ProtoToCloudschedulerJobStateEnum(p.GetState()),
		Status:              ProtoToCloudschedulerJobStatus(p.GetStatus()),
		ScheduleTime:        dcl.StringOrNil(p.GetScheduleTime()),
		LastAttemptTime:     dcl.StringOrNil(p.GetLastAttemptTime()),
		RetryConfig:         ProtoToCloudschedulerJobRetryConfig(p.GetRetryConfig()),
		AttemptDeadline:     dcl.StringOrNil(p.GetAttemptDeadline()),
		Project:             dcl.StringOrNil(p.GetProject()),
		Location:            dcl.StringOrNil(p.GetLocation()),
	}
	return obj
}

// JobAppEngineHttpTargetHttpMethodEnumToProto converts a JobAppEngineHttpTargetHttpMethodEnum enum to its proto representation.
func CloudschedulerJobAppEngineHttpTargetHttpMethodEnumToProto(e *cloudscheduler.JobAppEngineHttpTargetHttpMethodEnum) cloudschedulerpb.CloudschedulerJobAppEngineHttpTargetHttpMethodEnum {
	if e == nil {
		return cloudschedulerpb.CloudschedulerJobAppEngineHttpTargetHttpMethodEnum(0)
	}
	if v, ok := cloudschedulerpb.CloudschedulerJobAppEngineHttpTargetHttpMethodEnum_value["JobAppEngineHttpTargetHttpMethodEnum"+string(*e)]; ok {
		return cloudschedulerpb.CloudschedulerJobAppEngineHttpTargetHttpMethodEnum(v)
	}
	return cloudschedulerpb.CloudschedulerJobAppEngineHttpTargetHttpMethodEnum(0)
}

// JobHttpTargetHttpMethodEnumToProto converts a JobHttpTargetHttpMethodEnum enum to its proto representation.
func CloudschedulerJobHttpTargetHttpMethodEnumToProto(e *cloudscheduler.JobHttpTargetHttpMethodEnum) cloudschedulerpb.CloudschedulerJobHttpTargetHttpMethodEnum {
	if e == nil {
		return cloudschedulerpb.CloudschedulerJobHttpTargetHttpMethodEnum(0)
	}
	if v, ok := cloudschedulerpb.CloudschedulerJobHttpTargetHttpMethodEnum_value["JobHttpTargetHttpMethodEnum"+string(*e)]; ok {
		return cloudschedulerpb.CloudschedulerJobHttpTargetHttpMethodEnum(v)
	}
	return cloudschedulerpb.CloudschedulerJobHttpTargetHttpMethodEnum(0)
}

// JobStateEnumToProto converts a JobStateEnum enum to its proto representation.
func CloudschedulerJobStateEnumToProto(e *cloudscheduler.JobStateEnum) cloudschedulerpb.CloudschedulerJobStateEnum {
	if e == nil {
		return cloudschedulerpb.CloudschedulerJobStateEnum(0)
	}
	if v, ok := cloudschedulerpb.CloudschedulerJobStateEnum_value["JobStateEnum"+string(*e)]; ok {
		return cloudschedulerpb.CloudschedulerJobStateEnum(v)
	}
	return cloudschedulerpb.CloudschedulerJobStateEnum(0)
}

// JobPubsubTargetToProto converts a JobPubsubTarget object to its proto representation.
func CloudschedulerJobPubsubTargetToProto(o *cloudscheduler.JobPubsubTarget) *cloudschedulerpb.CloudschedulerJobPubsubTarget {
	if o == nil {
		return nil
	}
	p := &cloudschedulerpb.CloudschedulerJobPubsubTarget{}
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
func CloudschedulerJobAppEngineHttpTargetToProto(o *cloudscheduler.JobAppEngineHttpTarget) *cloudschedulerpb.CloudschedulerJobAppEngineHttpTarget {
	if o == nil {
		return nil
	}
	p := &cloudschedulerpb.CloudschedulerJobAppEngineHttpTarget{}
	p.SetHttpMethod(CloudschedulerJobAppEngineHttpTargetHttpMethodEnumToProto(o.HttpMethod))
	p.SetAppEngineRouting(CloudschedulerJobAppEngineHttpTargetAppEngineRoutingToProto(o.AppEngineRouting))
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
func CloudschedulerJobAppEngineHttpTargetAppEngineRoutingToProto(o *cloudscheduler.JobAppEngineHttpTargetAppEngineRouting) *cloudschedulerpb.CloudschedulerJobAppEngineHttpTargetAppEngineRouting {
	if o == nil {
		return nil
	}
	p := &cloudschedulerpb.CloudschedulerJobAppEngineHttpTargetAppEngineRouting{}
	p.SetService(dcl.ValueOrEmptyString(o.Service))
	p.SetVersion(dcl.ValueOrEmptyString(o.Version))
	p.SetInstance(dcl.ValueOrEmptyString(o.Instance))
	p.SetHost(dcl.ValueOrEmptyString(o.Host))
	return p
}

// JobHttpTargetToProto converts a JobHttpTarget object to its proto representation.
func CloudschedulerJobHttpTargetToProto(o *cloudscheduler.JobHttpTarget) *cloudschedulerpb.CloudschedulerJobHttpTarget {
	if o == nil {
		return nil
	}
	p := &cloudschedulerpb.CloudschedulerJobHttpTarget{}
	p.SetUri(dcl.ValueOrEmptyString(o.Uri))
	p.SetHttpMethod(CloudschedulerJobHttpTargetHttpMethodEnumToProto(o.HttpMethod))
	p.SetBody(dcl.ValueOrEmptyString(o.Body))
	p.SetOauthToken(CloudschedulerJobHttpTargetOAuthTokenToProto(o.OAuthToken))
	p.SetOidcToken(CloudschedulerJobHttpTargetOidcTokenToProto(o.OidcToken))
	mHeaders := make(map[string]string, len(o.Headers))
	for k, r := range o.Headers {
		mHeaders[k] = r
	}
	p.SetHeaders(mHeaders)
	return p
}

// JobHttpTargetOAuthTokenToProto converts a JobHttpTargetOAuthToken object to its proto representation.
func CloudschedulerJobHttpTargetOAuthTokenToProto(o *cloudscheduler.JobHttpTargetOAuthToken) *cloudschedulerpb.CloudschedulerJobHttpTargetOAuthToken {
	if o == nil {
		return nil
	}
	p := &cloudschedulerpb.CloudschedulerJobHttpTargetOAuthToken{}
	p.SetServiceAccountEmail(dcl.ValueOrEmptyString(o.ServiceAccountEmail))
	p.SetScope(dcl.ValueOrEmptyString(o.Scope))
	return p
}

// JobHttpTargetOidcTokenToProto converts a JobHttpTargetOidcToken object to its proto representation.
func CloudschedulerJobHttpTargetOidcTokenToProto(o *cloudscheduler.JobHttpTargetOidcToken) *cloudschedulerpb.CloudschedulerJobHttpTargetOidcToken {
	if o == nil {
		return nil
	}
	p := &cloudschedulerpb.CloudschedulerJobHttpTargetOidcToken{}
	p.SetServiceAccountEmail(dcl.ValueOrEmptyString(o.ServiceAccountEmail))
	p.SetAudience(dcl.ValueOrEmptyString(o.Audience))
	return p
}

// JobStatusToProto converts a JobStatus object to its proto representation.
func CloudschedulerJobStatusToProto(o *cloudscheduler.JobStatus) *cloudschedulerpb.CloudschedulerJobStatus {
	if o == nil {
		return nil
	}
	p := &cloudschedulerpb.CloudschedulerJobStatus{}
	p.SetCode(dcl.ValueOrEmptyInt64(o.Code))
	p.SetMessage(dcl.ValueOrEmptyString(o.Message))
	sDetails := make([]*cloudschedulerpb.CloudschedulerJobStatusDetails, len(o.Details))
	for i, r := range o.Details {
		sDetails[i] = CloudschedulerJobStatusDetailsToProto(&r)
	}
	p.SetDetails(sDetails)
	return p
}

// JobStatusDetailsToProto converts a JobStatusDetails object to its proto representation.
func CloudschedulerJobStatusDetailsToProto(o *cloudscheduler.JobStatusDetails) *cloudschedulerpb.CloudschedulerJobStatusDetails {
	if o == nil {
		return nil
	}
	p := &cloudschedulerpb.CloudschedulerJobStatusDetails{}
	p.SetTypeUrl(dcl.ValueOrEmptyString(o.TypeUrl))
	p.SetValue(dcl.ValueOrEmptyString(o.Value))
	return p
}

// JobRetryConfigToProto converts a JobRetryConfig object to its proto representation.
func CloudschedulerJobRetryConfigToProto(o *cloudscheduler.JobRetryConfig) *cloudschedulerpb.CloudschedulerJobRetryConfig {
	if o == nil {
		return nil
	}
	p := &cloudschedulerpb.CloudschedulerJobRetryConfig{}
	p.SetRetryCount(dcl.ValueOrEmptyInt64(o.RetryCount))
	p.SetMaxRetryDuration(dcl.ValueOrEmptyString(o.MaxRetryDuration))
	p.SetMinBackoffDuration(dcl.ValueOrEmptyString(o.MinBackoffDuration))
	p.SetMaxBackoffDuration(dcl.ValueOrEmptyString(o.MaxBackoffDuration))
	p.SetMaxDoublings(dcl.ValueOrEmptyInt64(o.MaxDoublings))
	return p
}

// JobToProto converts a Job resource to its proto representation.
func JobToProto(resource *cloudscheduler.Job) *cloudschedulerpb.CloudschedulerJob {
	p := &cloudschedulerpb.CloudschedulerJob{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetPubsubTarget(CloudschedulerJobPubsubTargetToProto(resource.PubsubTarget))
	p.SetAppEngineHttpTarget(CloudschedulerJobAppEngineHttpTargetToProto(resource.AppEngineHttpTarget))
	p.SetHttpTarget(CloudschedulerJobHttpTargetToProto(resource.HttpTarget))
	p.SetSchedule(dcl.ValueOrEmptyString(resource.Schedule))
	p.SetTimeZone(dcl.ValueOrEmptyString(resource.TimeZone))
	p.SetUserUpdateTime(dcl.ValueOrEmptyString(resource.UserUpdateTime))
	p.SetState(CloudschedulerJobStateEnumToProto(resource.State))
	p.SetStatus(CloudschedulerJobStatusToProto(resource.Status))
	p.SetScheduleTime(dcl.ValueOrEmptyString(resource.ScheduleTime))
	p.SetLastAttemptTime(dcl.ValueOrEmptyString(resource.LastAttemptTime))
	p.SetRetryConfig(CloudschedulerJobRetryConfigToProto(resource.RetryConfig))
	p.SetAttemptDeadline(dcl.ValueOrEmptyString(resource.AttemptDeadline))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))

	return p
}

// applyJob handles the gRPC request by passing it to the underlying Job Apply() method.
func (s *JobServer) applyJob(ctx context.Context, c *cloudscheduler.Client, request *cloudschedulerpb.ApplyCloudschedulerJobRequest) (*cloudschedulerpb.CloudschedulerJob, error) {
	p := ProtoToJob(request.GetResource())
	res, err := c.ApplyJob(ctx, p)
	if err != nil {
		return nil, err
	}
	r := JobToProto(res)
	return r, nil
}

// applyCloudschedulerJob handles the gRPC request by passing it to the underlying Job Apply() method.
func (s *JobServer) ApplyCloudschedulerJob(ctx context.Context, request *cloudschedulerpb.ApplyCloudschedulerJobRequest) (*cloudschedulerpb.CloudschedulerJob, error) {
	cl, err := createConfigJob(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyJob(ctx, cl, request)
}

// DeleteJob handles the gRPC request by passing it to the underlying Job Delete() method.
func (s *JobServer) DeleteCloudschedulerJob(ctx context.Context, request *cloudschedulerpb.DeleteCloudschedulerJobRequest) (*emptypb.Empty, error) {

	cl, err := createConfigJob(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteJob(ctx, ProtoToJob(request.GetResource()))

}

// ListCloudschedulerJob handles the gRPC request by passing it to the underlying JobList() method.
func (s *JobServer) ListCloudschedulerJob(ctx context.Context, request *cloudschedulerpb.ListCloudschedulerJobRequest) (*cloudschedulerpb.ListCloudschedulerJobResponse, error) {
	cl, err := createConfigJob(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListJob(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*cloudschedulerpb.CloudschedulerJob
	for _, r := range resources.Items {
		rp := JobToProto(r)
		protos = append(protos, rp)
	}
	p := &cloudschedulerpb.ListCloudschedulerJobResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigJob(ctx context.Context, service_account_file string) (*cloudscheduler.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return cloudscheduler.NewClient(conf), nil
}
