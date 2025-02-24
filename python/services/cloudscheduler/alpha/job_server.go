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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/cloudscheduler/alpha/cloudscheduler_alpha_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudscheduler/alpha"
)

// JobServer implements the gRPC interface for Job.
type JobServer struct{}

// ProtoToJobAppEngineHttpTargetHttpMethodEnum converts a JobAppEngineHttpTargetHttpMethodEnum enum from its proto representation.
func ProtoToCloudschedulerAlphaJobAppEngineHttpTargetHttpMethodEnum(e alphapb.CloudschedulerAlphaJobAppEngineHttpTargetHttpMethodEnum) *alpha.JobAppEngineHttpTargetHttpMethodEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.CloudschedulerAlphaJobAppEngineHttpTargetHttpMethodEnum_name[int32(e)]; ok {
		e := alpha.JobAppEngineHttpTargetHttpMethodEnum(n[len("CloudschedulerAlphaJobAppEngineHttpTargetHttpMethodEnum"):])
		return &e
	}
	return nil
}

// ProtoToJobHttpTargetHttpMethodEnum converts a JobHttpTargetHttpMethodEnum enum from its proto representation.
func ProtoToCloudschedulerAlphaJobHttpTargetHttpMethodEnum(e alphapb.CloudschedulerAlphaJobHttpTargetHttpMethodEnum) *alpha.JobHttpTargetHttpMethodEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.CloudschedulerAlphaJobHttpTargetHttpMethodEnum_name[int32(e)]; ok {
		e := alpha.JobHttpTargetHttpMethodEnum(n[len("CloudschedulerAlphaJobHttpTargetHttpMethodEnum"):])
		return &e
	}
	return nil
}

// ProtoToJobStateEnum converts a JobStateEnum enum from its proto representation.
func ProtoToCloudschedulerAlphaJobStateEnum(e alphapb.CloudschedulerAlphaJobStateEnum) *alpha.JobStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.CloudschedulerAlphaJobStateEnum_name[int32(e)]; ok {
		e := alpha.JobStateEnum(n[len("CloudschedulerAlphaJobStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToJobPubsubTarget converts a JobPubsubTarget object from its proto representation.
func ProtoToCloudschedulerAlphaJobPubsubTarget(p *alphapb.CloudschedulerAlphaJobPubsubTarget) *alpha.JobPubsubTarget {
	if p == nil {
		return nil
	}
	obj := &alpha.JobPubsubTarget{
		TopicName: dcl.StringOrNil(p.GetTopicName()),
		Data:      dcl.StringOrNil(p.GetData()),
	}
	return obj
}

// ProtoToJobAppEngineHttpTarget converts a JobAppEngineHttpTarget object from its proto representation.
func ProtoToCloudschedulerAlphaJobAppEngineHttpTarget(p *alphapb.CloudschedulerAlphaJobAppEngineHttpTarget) *alpha.JobAppEngineHttpTarget {
	if p == nil {
		return nil
	}
	obj := &alpha.JobAppEngineHttpTarget{
		HttpMethod:       ProtoToCloudschedulerAlphaJobAppEngineHttpTargetHttpMethodEnum(p.GetHttpMethod()),
		AppEngineRouting: ProtoToCloudschedulerAlphaJobAppEngineHttpTargetAppEngineRouting(p.GetAppEngineRouting()),
		RelativeUri:      dcl.StringOrNil(p.GetRelativeUri()),
		Body:             dcl.StringOrNil(p.GetBody()),
	}
	return obj
}

// ProtoToJobAppEngineHttpTargetAppEngineRouting converts a JobAppEngineHttpTargetAppEngineRouting object from its proto representation.
func ProtoToCloudschedulerAlphaJobAppEngineHttpTargetAppEngineRouting(p *alphapb.CloudschedulerAlphaJobAppEngineHttpTargetAppEngineRouting) *alpha.JobAppEngineHttpTargetAppEngineRouting {
	if p == nil {
		return nil
	}
	obj := &alpha.JobAppEngineHttpTargetAppEngineRouting{
		Service:  dcl.StringOrNil(p.GetService()),
		Version:  dcl.StringOrNil(p.GetVersion()),
		Instance: dcl.StringOrNil(p.GetInstance()),
		Host:     dcl.StringOrNil(p.GetHost()),
	}
	return obj
}

// ProtoToJobHttpTarget converts a JobHttpTarget object from its proto representation.
func ProtoToCloudschedulerAlphaJobHttpTarget(p *alphapb.CloudschedulerAlphaJobHttpTarget) *alpha.JobHttpTarget {
	if p == nil {
		return nil
	}
	obj := &alpha.JobHttpTarget{
		Uri:        dcl.StringOrNil(p.GetUri()),
		HttpMethod: ProtoToCloudschedulerAlphaJobHttpTargetHttpMethodEnum(p.GetHttpMethod()),
		Body:       dcl.StringOrNil(p.GetBody()),
		OAuthToken: ProtoToCloudschedulerAlphaJobHttpTargetOAuthToken(p.GetOauthToken()),
		OidcToken:  ProtoToCloudschedulerAlphaJobHttpTargetOidcToken(p.GetOidcToken()),
	}
	return obj
}

// ProtoToJobHttpTargetOAuthToken converts a JobHttpTargetOAuthToken object from its proto representation.
func ProtoToCloudschedulerAlphaJobHttpTargetOAuthToken(p *alphapb.CloudschedulerAlphaJobHttpTargetOAuthToken) *alpha.JobHttpTargetOAuthToken {
	if p == nil {
		return nil
	}
	obj := &alpha.JobHttpTargetOAuthToken{
		ServiceAccountEmail: dcl.StringOrNil(p.GetServiceAccountEmail()),
		Scope:               dcl.StringOrNil(p.GetScope()),
	}
	return obj
}

// ProtoToJobHttpTargetOidcToken converts a JobHttpTargetOidcToken object from its proto representation.
func ProtoToCloudschedulerAlphaJobHttpTargetOidcToken(p *alphapb.CloudschedulerAlphaJobHttpTargetOidcToken) *alpha.JobHttpTargetOidcToken {
	if p == nil {
		return nil
	}
	obj := &alpha.JobHttpTargetOidcToken{
		ServiceAccountEmail: dcl.StringOrNil(p.GetServiceAccountEmail()),
		Audience:            dcl.StringOrNil(p.GetAudience()),
	}
	return obj
}

// ProtoToJobStatus converts a JobStatus object from its proto representation.
func ProtoToCloudschedulerAlphaJobStatus(p *alphapb.CloudschedulerAlphaJobStatus) *alpha.JobStatus {
	if p == nil {
		return nil
	}
	obj := &alpha.JobStatus{
		Code:    dcl.Int64OrNil(p.GetCode()),
		Message: dcl.StringOrNil(p.GetMessage()),
	}
	for _, r := range p.GetDetails() {
		obj.Details = append(obj.Details, *ProtoToCloudschedulerAlphaJobStatusDetails(r))
	}
	return obj
}

// ProtoToJobStatusDetails converts a JobStatusDetails object from its proto representation.
func ProtoToCloudschedulerAlphaJobStatusDetails(p *alphapb.CloudschedulerAlphaJobStatusDetails) *alpha.JobStatusDetails {
	if p == nil {
		return nil
	}
	obj := &alpha.JobStatusDetails{
		TypeUrl: dcl.StringOrNil(p.GetTypeUrl()),
		Value:   dcl.StringOrNil(p.GetValue()),
	}
	return obj
}

// ProtoToJobRetryConfig converts a JobRetryConfig object from its proto representation.
func ProtoToCloudschedulerAlphaJobRetryConfig(p *alphapb.CloudschedulerAlphaJobRetryConfig) *alpha.JobRetryConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.JobRetryConfig{
		RetryCount:         dcl.Int64OrNil(p.GetRetryCount()),
		MaxRetryDuration:   dcl.StringOrNil(p.GetMaxRetryDuration()),
		MinBackoffDuration: dcl.StringOrNil(p.GetMinBackoffDuration()),
		MaxBackoffDuration: dcl.StringOrNil(p.GetMaxBackoffDuration()),
		MaxDoublings:       dcl.Int64OrNil(p.GetMaxDoublings()),
	}
	return obj
}

// ProtoToJob converts a Job resource from its proto representation.
func ProtoToJob(p *alphapb.CloudschedulerAlphaJob) *alpha.Job {
	obj := &alpha.Job{
		Name:                dcl.StringOrNil(p.GetName()),
		Description:         dcl.StringOrNil(p.GetDescription()),
		PubsubTarget:        ProtoToCloudschedulerAlphaJobPubsubTarget(p.GetPubsubTarget()),
		AppEngineHttpTarget: ProtoToCloudschedulerAlphaJobAppEngineHttpTarget(p.GetAppEngineHttpTarget()),
		HttpTarget:          ProtoToCloudschedulerAlphaJobHttpTarget(p.GetHttpTarget()),
		Schedule:            dcl.StringOrNil(p.GetSchedule()),
		TimeZone:            dcl.StringOrNil(p.GetTimeZone()),
		UserUpdateTime:      dcl.StringOrNil(p.GetUserUpdateTime()),
		State:               ProtoToCloudschedulerAlphaJobStateEnum(p.GetState()),
		Status:              ProtoToCloudschedulerAlphaJobStatus(p.GetStatus()),
		ScheduleTime:        dcl.StringOrNil(p.GetScheduleTime()),
		LastAttemptTime:     dcl.StringOrNil(p.GetLastAttemptTime()),
		RetryConfig:         ProtoToCloudschedulerAlphaJobRetryConfig(p.GetRetryConfig()),
		AttemptDeadline:     dcl.StringOrNil(p.GetAttemptDeadline()),
		Project:             dcl.StringOrNil(p.GetProject()),
		Location:            dcl.StringOrNil(p.GetLocation()),
	}
	return obj
}

// JobAppEngineHttpTargetHttpMethodEnumToProto converts a JobAppEngineHttpTargetHttpMethodEnum enum to its proto representation.
func CloudschedulerAlphaJobAppEngineHttpTargetHttpMethodEnumToProto(e *alpha.JobAppEngineHttpTargetHttpMethodEnum) alphapb.CloudschedulerAlphaJobAppEngineHttpTargetHttpMethodEnum {
	if e == nil {
		return alphapb.CloudschedulerAlphaJobAppEngineHttpTargetHttpMethodEnum(0)
	}
	if v, ok := alphapb.CloudschedulerAlphaJobAppEngineHttpTargetHttpMethodEnum_value["JobAppEngineHttpTargetHttpMethodEnum"+string(*e)]; ok {
		return alphapb.CloudschedulerAlphaJobAppEngineHttpTargetHttpMethodEnum(v)
	}
	return alphapb.CloudschedulerAlphaJobAppEngineHttpTargetHttpMethodEnum(0)
}

// JobHttpTargetHttpMethodEnumToProto converts a JobHttpTargetHttpMethodEnum enum to its proto representation.
func CloudschedulerAlphaJobHttpTargetHttpMethodEnumToProto(e *alpha.JobHttpTargetHttpMethodEnum) alphapb.CloudschedulerAlphaJobHttpTargetHttpMethodEnum {
	if e == nil {
		return alphapb.CloudschedulerAlphaJobHttpTargetHttpMethodEnum(0)
	}
	if v, ok := alphapb.CloudschedulerAlphaJobHttpTargetHttpMethodEnum_value["JobHttpTargetHttpMethodEnum"+string(*e)]; ok {
		return alphapb.CloudschedulerAlphaJobHttpTargetHttpMethodEnum(v)
	}
	return alphapb.CloudschedulerAlphaJobHttpTargetHttpMethodEnum(0)
}

// JobStateEnumToProto converts a JobStateEnum enum to its proto representation.
func CloudschedulerAlphaJobStateEnumToProto(e *alpha.JobStateEnum) alphapb.CloudschedulerAlphaJobStateEnum {
	if e == nil {
		return alphapb.CloudschedulerAlphaJobStateEnum(0)
	}
	if v, ok := alphapb.CloudschedulerAlphaJobStateEnum_value["JobStateEnum"+string(*e)]; ok {
		return alphapb.CloudschedulerAlphaJobStateEnum(v)
	}
	return alphapb.CloudschedulerAlphaJobStateEnum(0)
}

// JobPubsubTargetToProto converts a JobPubsubTarget object to its proto representation.
func CloudschedulerAlphaJobPubsubTargetToProto(o *alpha.JobPubsubTarget) *alphapb.CloudschedulerAlphaJobPubsubTarget {
	if o == nil {
		return nil
	}
	p := &alphapb.CloudschedulerAlphaJobPubsubTarget{}
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
func CloudschedulerAlphaJobAppEngineHttpTargetToProto(o *alpha.JobAppEngineHttpTarget) *alphapb.CloudschedulerAlphaJobAppEngineHttpTarget {
	if o == nil {
		return nil
	}
	p := &alphapb.CloudschedulerAlphaJobAppEngineHttpTarget{}
	p.SetHttpMethod(CloudschedulerAlphaJobAppEngineHttpTargetHttpMethodEnumToProto(o.HttpMethod))
	p.SetAppEngineRouting(CloudschedulerAlphaJobAppEngineHttpTargetAppEngineRoutingToProto(o.AppEngineRouting))
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
func CloudschedulerAlphaJobAppEngineHttpTargetAppEngineRoutingToProto(o *alpha.JobAppEngineHttpTargetAppEngineRouting) *alphapb.CloudschedulerAlphaJobAppEngineHttpTargetAppEngineRouting {
	if o == nil {
		return nil
	}
	p := &alphapb.CloudschedulerAlphaJobAppEngineHttpTargetAppEngineRouting{}
	p.SetService(dcl.ValueOrEmptyString(o.Service))
	p.SetVersion(dcl.ValueOrEmptyString(o.Version))
	p.SetInstance(dcl.ValueOrEmptyString(o.Instance))
	p.SetHost(dcl.ValueOrEmptyString(o.Host))
	return p
}

// JobHttpTargetToProto converts a JobHttpTarget object to its proto representation.
func CloudschedulerAlphaJobHttpTargetToProto(o *alpha.JobHttpTarget) *alphapb.CloudschedulerAlphaJobHttpTarget {
	if o == nil {
		return nil
	}
	p := &alphapb.CloudschedulerAlphaJobHttpTarget{}
	p.SetUri(dcl.ValueOrEmptyString(o.Uri))
	p.SetHttpMethod(CloudschedulerAlphaJobHttpTargetHttpMethodEnumToProto(o.HttpMethod))
	p.SetBody(dcl.ValueOrEmptyString(o.Body))
	p.SetOauthToken(CloudschedulerAlphaJobHttpTargetOAuthTokenToProto(o.OAuthToken))
	p.SetOidcToken(CloudschedulerAlphaJobHttpTargetOidcTokenToProto(o.OidcToken))
	mHeaders := make(map[string]string, len(o.Headers))
	for k, r := range o.Headers {
		mHeaders[k] = r
	}
	p.SetHeaders(mHeaders)
	return p
}

// JobHttpTargetOAuthTokenToProto converts a JobHttpTargetOAuthToken object to its proto representation.
func CloudschedulerAlphaJobHttpTargetOAuthTokenToProto(o *alpha.JobHttpTargetOAuthToken) *alphapb.CloudschedulerAlphaJobHttpTargetOAuthToken {
	if o == nil {
		return nil
	}
	p := &alphapb.CloudschedulerAlphaJobHttpTargetOAuthToken{}
	p.SetServiceAccountEmail(dcl.ValueOrEmptyString(o.ServiceAccountEmail))
	p.SetScope(dcl.ValueOrEmptyString(o.Scope))
	return p
}

// JobHttpTargetOidcTokenToProto converts a JobHttpTargetOidcToken object to its proto representation.
func CloudschedulerAlphaJobHttpTargetOidcTokenToProto(o *alpha.JobHttpTargetOidcToken) *alphapb.CloudschedulerAlphaJobHttpTargetOidcToken {
	if o == nil {
		return nil
	}
	p := &alphapb.CloudschedulerAlphaJobHttpTargetOidcToken{}
	p.SetServiceAccountEmail(dcl.ValueOrEmptyString(o.ServiceAccountEmail))
	p.SetAudience(dcl.ValueOrEmptyString(o.Audience))
	return p
}

// JobStatusToProto converts a JobStatus object to its proto representation.
func CloudschedulerAlphaJobStatusToProto(o *alpha.JobStatus) *alphapb.CloudschedulerAlphaJobStatus {
	if o == nil {
		return nil
	}
	p := &alphapb.CloudschedulerAlphaJobStatus{}
	p.SetCode(dcl.ValueOrEmptyInt64(o.Code))
	p.SetMessage(dcl.ValueOrEmptyString(o.Message))
	sDetails := make([]*alphapb.CloudschedulerAlphaJobStatusDetails, len(o.Details))
	for i, r := range o.Details {
		sDetails[i] = CloudschedulerAlphaJobStatusDetailsToProto(&r)
	}
	p.SetDetails(sDetails)
	return p
}

// JobStatusDetailsToProto converts a JobStatusDetails object to its proto representation.
func CloudschedulerAlphaJobStatusDetailsToProto(o *alpha.JobStatusDetails) *alphapb.CloudschedulerAlphaJobStatusDetails {
	if o == nil {
		return nil
	}
	p := &alphapb.CloudschedulerAlphaJobStatusDetails{}
	p.SetTypeUrl(dcl.ValueOrEmptyString(o.TypeUrl))
	p.SetValue(dcl.ValueOrEmptyString(o.Value))
	return p
}

// JobRetryConfigToProto converts a JobRetryConfig object to its proto representation.
func CloudschedulerAlphaJobRetryConfigToProto(o *alpha.JobRetryConfig) *alphapb.CloudschedulerAlphaJobRetryConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.CloudschedulerAlphaJobRetryConfig{}
	p.SetRetryCount(dcl.ValueOrEmptyInt64(o.RetryCount))
	p.SetMaxRetryDuration(dcl.ValueOrEmptyString(o.MaxRetryDuration))
	p.SetMinBackoffDuration(dcl.ValueOrEmptyString(o.MinBackoffDuration))
	p.SetMaxBackoffDuration(dcl.ValueOrEmptyString(o.MaxBackoffDuration))
	p.SetMaxDoublings(dcl.ValueOrEmptyInt64(o.MaxDoublings))
	return p
}

// JobToProto converts a Job resource to its proto representation.
func JobToProto(resource *alpha.Job) *alphapb.CloudschedulerAlphaJob {
	p := &alphapb.CloudschedulerAlphaJob{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetPubsubTarget(CloudschedulerAlphaJobPubsubTargetToProto(resource.PubsubTarget))
	p.SetAppEngineHttpTarget(CloudschedulerAlphaJobAppEngineHttpTargetToProto(resource.AppEngineHttpTarget))
	p.SetHttpTarget(CloudschedulerAlphaJobHttpTargetToProto(resource.HttpTarget))
	p.SetSchedule(dcl.ValueOrEmptyString(resource.Schedule))
	p.SetTimeZone(dcl.ValueOrEmptyString(resource.TimeZone))
	p.SetUserUpdateTime(dcl.ValueOrEmptyString(resource.UserUpdateTime))
	p.SetState(CloudschedulerAlphaJobStateEnumToProto(resource.State))
	p.SetStatus(CloudschedulerAlphaJobStatusToProto(resource.Status))
	p.SetScheduleTime(dcl.ValueOrEmptyString(resource.ScheduleTime))
	p.SetLastAttemptTime(dcl.ValueOrEmptyString(resource.LastAttemptTime))
	p.SetRetryConfig(CloudschedulerAlphaJobRetryConfigToProto(resource.RetryConfig))
	p.SetAttemptDeadline(dcl.ValueOrEmptyString(resource.AttemptDeadline))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))

	return p
}

// applyJob handles the gRPC request by passing it to the underlying Job Apply() method.
func (s *JobServer) applyJob(ctx context.Context, c *alpha.Client, request *alphapb.ApplyCloudschedulerAlphaJobRequest) (*alphapb.CloudschedulerAlphaJob, error) {
	p := ProtoToJob(request.GetResource())
	res, err := c.ApplyJob(ctx, p)
	if err != nil {
		return nil, err
	}
	r := JobToProto(res)
	return r, nil
}

// applyCloudschedulerAlphaJob handles the gRPC request by passing it to the underlying Job Apply() method.
func (s *JobServer) ApplyCloudschedulerAlphaJob(ctx context.Context, request *alphapb.ApplyCloudschedulerAlphaJobRequest) (*alphapb.CloudschedulerAlphaJob, error) {
	cl, err := createConfigJob(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyJob(ctx, cl, request)
}

// DeleteJob handles the gRPC request by passing it to the underlying Job Delete() method.
func (s *JobServer) DeleteCloudschedulerAlphaJob(ctx context.Context, request *alphapb.DeleteCloudschedulerAlphaJobRequest) (*emptypb.Empty, error) {

	cl, err := createConfigJob(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteJob(ctx, ProtoToJob(request.GetResource()))

}

// ListCloudschedulerAlphaJob handles the gRPC request by passing it to the underlying JobList() method.
func (s *JobServer) ListCloudschedulerAlphaJob(ctx context.Context, request *alphapb.ListCloudschedulerAlphaJobRequest) (*alphapb.ListCloudschedulerAlphaJobResponse, error) {
	cl, err := createConfigJob(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListJob(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.CloudschedulerAlphaJob
	for _, r := range resources.Items {
		rp := JobToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListCloudschedulerAlphaJobResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigJob(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
