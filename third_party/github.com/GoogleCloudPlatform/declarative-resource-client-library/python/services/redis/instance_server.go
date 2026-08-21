// Copyright 2021 Google LLC. All Rights Reserved.
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
	redispb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/redis/redis_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/redis"
)

// Server implements the gRPC interface for Instance.
type InstanceServer struct{}

// ProtoToInstanceStateEnum converts a InstanceStateEnum enum from its proto representation.
func ProtoToRedisInstanceStateEnum(e redispb.RedisInstanceStateEnum) *redis.InstanceStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := redispb.RedisInstanceStateEnum_name[int32(e)]; ok {
		e := redis.InstanceStateEnum(n[len("RedisInstanceStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceTierEnum converts a InstanceTierEnum enum from its proto representation.
func ProtoToRedisInstanceTierEnum(e redispb.RedisInstanceTierEnum) *redis.InstanceTierEnum {
	if e == 0 {
		return nil
	}
	if n, ok := redispb.RedisInstanceTierEnum_name[int32(e)]; ok {
		e := redis.InstanceTierEnum(n[len("RedisInstanceTierEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceConnectModeEnum converts a InstanceConnectModeEnum enum from its proto representation.
func ProtoToRedisInstanceConnectModeEnum(e redispb.RedisInstanceConnectModeEnum) *redis.InstanceConnectModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := redispb.RedisInstanceConnectModeEnum_name[int32(e)]; ok {
		e := redis.InstanceConnectModeEnum(n[len("RedisInstanceConnectModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceTransitEncryptionModeEnum converts a InstanceTransitEncryptionModeEnum enum from its proto representation.
func ProtoToRedisInstanceTransitEncryptionModeEnum(e redispb.RedisInstanceTransitEncryptionModeEnum) *redis.InstanceTransitEncryptionModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := redispb.RedisInstanceTransitEncryptionModeEnum_name[int32(e)]; ok {
		e := redis.InstanceTransitEncryptionModeEnum(n[len("RedisInstanceTransitEncryptionModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceMaintenancePolicyWeeklyMaintenanceWindowDayEnum converts a InstanceMaintenancePolicyWeeklyMaintenanceWindowDayEnum enum from its proto representation.
func ProtoToRedisInstanceMaintenancePolicyWeeklyMaintenanceWindowDayEnum(e redispb.RedisInstanceMaintenancePolicyWeeklyMaintenanceWindowDayEnum) *redis.InstanceMaintenancePolicyWeeklyMaintenanceWindowDayEnum {
	if e == 0 {
		return nil
	}
	if n, ok := redispb.RedisInstanceMaintenancePolicyWeeklyMaintenanceWindowDayEnum_name[int32(e)]; ok {
		e := redis.InstanceMaintenancePolicyWeeklyMaintenanceWindowDayEnum(n[len("RedisInstanceMaintenancePolicyWeeklyMaintenanceWindowDayEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceServerCaCerts converts a InstanceServerCaCerts resource from its proto representation.
func ProtoToRedisInstanceServerCaCerts(p *redispb.RedisInstanceServerCaCerts) *redis.InstanceServerCaCerts {
	if p == nil {
		return nil
	}
	obj := &redis.InstanceServerCaCerts{
		SerialNumber:    dcl.StringOrNil(p.SerialNumber),
		Cert:            dcl.StringOrNil(p.Cert),
		CreateTime:      dcl.StringOrNil(p.GetCreateTime()),
		ExpireTime:      dcl.StringOrNil(p.GetExpireTime()),
		Sha1Fingerprint: dcl.StringOrNil(p.Sha1Fingerprint),
	}
	return obj
}

// ProtoToInstanceMaintenancePolicy converts a InstanceMaintenancePolicy resource from its proto representation.
func ProtoToRedisInstanceMaintenancePolicy(p *redispb.RedisInstanceMaintenancePolicy) *redis.InstanceMaintenancePolicy {
	if p == nil {
		return nil
	}
	obj := &redis.InstanceMaintenancePolicy{
		CreateTime:  dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:  dcl.StringOrNil(p.GetUpdateTime()),
		Description: dcl.StringOrNil(p.Description),
	}
	for _, r := range p.GetWeeklyMaintenanceWindow() {
		obj.WeeklyMaintenanceWindow = append(obj.WeeklyMaintenanceWindow, *ProtoToRedisInstanceMaintenancePolicyWeeklyMaintenanceWindow(r))
	}
	return obj
}

// ProtoToInstanceMaintenancePolicyWeeklyMaintenanceWindow converts a InstanceMaintenancePolicyWeeklyMaintenanceWindow resource from its proto representation.
func ProtoToRedisInstanceMaintenancePolicyWeeklyMaintenanceWindow(p *redispb.RedisInstanceMaintenancePolicyWeeklyMaintenanceWindow) *redis.InstanceMaintenancePolicyWeeklyMaintenanceWindow {
	if p == nil {
		return nil
	}
	obj := &redis.InstanceMaintenancePolicyWeeklyMaintenanceWindow{
		Day:       ProtoToRedisInstanceMaintenancePolicyWeeklyMaintenanceWindowDayEnum(p.GetDay()),
		StartTime: ProtoToRedisInstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime(p.GetStartTime()),
		Duration:  dcl.StringOrNil(p.Duration),
	}
	return obj
}

// ProtoToInstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime converts a InstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime resource from its proto representation.
func ProtoToRedisInstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime(p *redispb.RedisInstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime) *redis.InstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime {
	if p == nil {
		return nil
	}
	obj := &redis.InstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime{
		Hours:   dcl.Int64OrNil(p.Hours),
		Minutes: dcl.Int64OrNil(p.Minutes),
		Seconds: dcl.Int64OrNil(p.Seconds),
		Nanos:   dcl.Int64OrNil(p.Nanos),
	}
	return obj
}

// ProtoToInstanceMaintenanceSchedule converts a InstanceMaintenanceSchedule resource from its proto representation.
func ProtoToRedisInstanceMaintenanceSchedule(p *redispb.RedisInstanceMaintenanceSchedule) *redis.InstanceMaintenanceSchedule {
	if p == nil {
		return nil
	}
	obj := &redis.InstanceMaintenanceSchedule{
		StartTime:            dcl.StringOrNil(p.GetStartTime()),
		EndTime:              dcl.StringOrNil(p.GetEndTime()),
		CanReschedule:        dcl.Bool(p.CanReschedule),
		ScheduleDeadlineTime: dcl.StringOrNil(p.GetScheduleDeadlineTime()),
	}
	return obj
}

// ProtoToInstance converts a Instance resource from its proto representation.
func ProtoToInstance(p *redispb.RedisInstance) *redis.Instance {
	obj := &redis.Instance{
		Name:                   dcl.StringOrNil(p.Name),
		DisplayName:            dcl.StringOrNil(p.DisplayName),
		LocationId:             dcl.StringOrNil(p.LocationId),
		AlternativeLocationId:  dcl.StringOrNil(p.AlternativeLocationId),
		RedisVersion:           dcl.StringOrNil(p.RedisVersion),
		ReservedIPRange:        dcl.StringOrNil(p.ReservedIpRange),
		Host:                   dcl.StringOrNil(p.Host),
		Port:                   dcl.Int64OrNil(p.Port),
		CurrentLocationId:      dcl.StringOrNil(p.CurrentLocationId),
		CreateTime:             dcl.StringOrNil(p.GetCreateTime()),
		State:                  ProtoToRedisInstanceStateEnum(p.GetState()),
		StatusMessage:          dcl.StringOrNil(p.StatusMessage),
		Tier:                   ProtoToRedisInstanceTierEnum(p.GetTier()),
		MemorySizeGb:           dcl.Int64OrNil(p.MemorySizeGb),
		AuthorizedNetwork:      dcl.StringOrNil(p.AuthorizedNetwork),
		PersistenceIamIdentity: dcl.StringOrNil(p.PersistenceIamIdentity),
		ConnectMode:            ProtoToRedisInstanceConnectModeEnum(p.GetConnectMode()),
		AuthEnabled:            dcl.Bool(p.AuthEnabled),
		TransitEncryptionMode:  ProtoToRedisInstanceTransitEncryptionModeEnum(p.GetTransitEncryptionMode()),
		MaintenancePolicy:      ProtoToRedisInstanceMaintenancePolicy(p.GetMaintenancePolicy()),
		MaintenanceSchedule:    ProtoToRedisInstanceMaintenanceSchedule(p.GetMaintenanceSchedule()),
		Project:                dcl.StringOrNil(p.Project),
		Location:               dcl.StringOrNil(p.Location),
	}
	for _, r := range p.GetServerCaCerts() {
		obj.ServerCaCerts = append(obj.ServerCaCerts, *ProtoToRedisInstanceServerCaCerts(r))
	}
	return obj
}

// InstanceStateEnumToProto converts a InstanceStateEnum enum to its proto representation.
func RedisInstanceStateEnumToProto(e *redis.InstanceStateEnum) redispb.RedisInstanceStateEnum {
	if e == nil {
		return redispb.RedisInstanceStateEnum(0)
	}
	if v, ok := redispb.RedisInstanceStateEnum_value["InstanceStateEnum"+string(*e)]; ok {
		return redispb.RedisInstanceStateEnum(v)
	}
	return redispb.RedisInstanceStateEnum(0)
}

// InstanceTierEnumToProto converts a InstanceTierEnum enum to its proto representation.
func RedisInstanceTierEnumToProto(e *redis.InstanceTierEnum) redispb.RedisInstanceTierEnum {
	if e == nil {
		return redispb.RedisInstanceTierEnum(0)
	}
	if v, ok := redispb.RedisInstanceTierEnum_value["InstanceTierEnum"+string(*e)]; ok {
		return redispb.RedisInstanceTierEnum(v)
	}
	return redispb.RedisInstanceTierEnum(0)
}

// InstanceConnectModeEnumToProto converts a InstanceConnectModeEnum enum to its proto representation.
func RedisInstanceConnectModeEnumToProto(e *redis.InstanceConnectModeEnum) redispb.RedisInstanceConnectModeEnum {
	if e == nil {
		return redispb.RedisInstanceConnectModeEnum(0)
	}
	if v, ok := redispb.RedisInstanceConnectModeEnum_value["InstanceConnectModeEnum"+string(*e)]; ok {
		return redispb.RedisInstanceConnectModeEnum(v)
	}
	return redispb.RedisInstanceConnectModeEnum(0)
}

// InstanceTransitEncryptionModeEnumToProto converts a InstanceTransitEncryptionModeEnum enum to its proto representation.
func RedisInstanceTransitEncryptionModeEnumToProto(e *redis.InstanceTransitEncryptionModeEnum) redispb.RedisInstanceTransitEncryptionModeEnum {
	if e == nil {
		return redispb.RedisInstanceTransitEncryptionModeEnum(0)
	}
	if v, ok := redispb.RedisInstanceTransitEncryptionModeEnum_value["InstanceTransitEncryptionModeEnum"+string(*e)]; ok {
		return redispb.RedisInstanceTransitEncryptionModeEnum(v)
	}
	return redispb.RedisInstanceTransitEncryptionModeEnum(0)
}

// InstanceMaintenancePolicyWeeklyMaintenanceWindowDayEnumToProto converts a InstanceMaintenancePolicyWeeklyMaintenanceWindowDayEnum enum to its proto representation.
func RedisInstanceMaintenancePolicyWeeklyMaintenanceWindowDayEnumToProto(e *redis.InstanceMaintenancePolicyWeeklyMaintenanceWindowDayEnum) redispb.RedisInstanceMaintenancePolicyWeeklyMaintenanceWindowDayEnum {
	if e == nil {
		return redispb.RedisInstanceMaintenancePolicyWeeklyMaintenanceWindowDayEnum(0)
	}
	if v, ok := redispb.RedisInstanceMaintenancePolicyWeeklyMaintenanceWindowDayEnum_value["InstanceMaintenancePolicyWeeklyMaintenanceWindowDayEnum"+string(*e)]; ok {
		return redispb.RedisInstanceMaintenancePolicyWeeklyMaintenanceWindowDayEnum(v)
	}
	return redispb.RedisInstanceMaintenancePolicyWeeklyMaintenanceWindowDayEnum(0)
}

// InstanceServerCaCertsToProto converts a InstanceServerCaCerts resource to its proto representation.
func RedisInstanceServerCaCertsToProto(o *redis.InstanceServerCaCerts) *redispb.RedisInstanceServerCaCerts {
	if o == nil {
		return nil
	}
	p := &redispb.RedisInstanceServerCaCerts{
		SerialNumber:    dcl.ValueOrEmptyString(o.SerialNumber),
		Cert:            dcl.ValueOrEmptyString(o.Cert),
		CreateTime:      dcl.ValueOrEmptyString(o.CreateTime),
		ExpireTime:      dcl.ValueOrEmptyString(o.ExpireTime),
		Sha1Fingerprint: dcl.ValueOrEmptyString(o.Sha1Fingerprint),
	}
	return p
}

// InstanceMaintenancePolicyToProto converts a InstanceMaintenancePolicy resource to its proto representation.
func RedisInstanceMaintenancePolicyToProto(o *redis.InstanceMaintenancePolicy) *redispb.RedisInstanceMaintenancePolicy {
	if o == nil {
		return nil
	}
	p := &redispb.RedisInstanceMaintenancePolicy{
		CreateTime:  dcl.ValueOrEmptyString(o.CreateTime),
		UpdateTime:  dcl.ValueOrEmptyString(o.UpdateTime),
		Description: dcl.ValueOrEmptyString(o.Description),
	}
	for _, r := range o.WeeklyMaintenanceWindow {
		p.WeeklyMaintenanceWindow = append(p.WeeklyMaintenanceWindow, RedisInstanceMaintenancePolicyWeeklyMaintenanceWindowToProto(&r))
	}
	return p
}

// InstanceMaintenancePolicyWeeklyMaintenanceWindowToProto converts a InstanceMaintenancePolicyWeeklyMaintenanceWindow resource to its proto representation.
func RedisInstanceMaintenancePolicyWeeklyMaintenanceWindowToProto(o *redis.InstanceMaintenancePolicyWeeklyMaintenanceWindow) *redispb.RedisInstanceMaintenancePolicyWeeklyMaintenanceWindow {
	if o == nil {
		return nil
	}
	p := &redispb.RedisInstanceMaintenancePolicyWeeklyMaintenanceWindow{
		Day:       RedisInstanceMaintenancePolicyWeeklyMaintenanceWindowDayEnumToProto(o.Day),
		StartTime: RedisInstanceMaintenancePolicyWeeklyMaintenanceWindowStartTimeToProto(o.StartTime),
		Duration:  dcl.ValueOrEmptyString(o.Duration),
	}
	return p
}

// InstanceMaintenancePolicyWeeklyMaintenanceWindowStartTimeToProto converts a InstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime resource to its proto representation.
func RedisInstanceMaintenancePolicyWeeklyMaintenanceWindowStartTimeToProto(o *redis.InstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime) *redispb.RedisInstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime {
	if o == nil {
		return nil
	}
	p := &redispb.RedisInstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime{
		Hours:   dcl.ValueOrEmptyInt64(o.Hours),
		Minutes: dcl.ValueOrEmptyInt64(o.Minutes),
		Seconds: dcl.ValueOrEmptyInt64(o.Seconds),
		Nanos:   dcl.ValueOrEmptyInt64(o.Nanos),
	}
	return p
}

// InstanceMaintenanceScheduleToProto converts a InstanceMaintenanceSchedule resource to its proto representation.
func RedisInstanceMaintenanceScheduleToProto(o *redis.InstanceMaintenanceSchedule) *redispb.RedisInstanceMaintenanceSchedule {
	if o == nil {
		return nil
	}
	p := &redispb.RedisInstanceMaintenanceSchedule{
		StartTime:            dcl.ValueOrEmptyString(o.StartTime),
		EndTime:              dcl.ValueOrEmptyString(o.EndTime),
		CanReschedule:        dcl.ValueOrEmptyBool(o.CanReschedule),
		ScheduleDeadlineTime: dcl.ValueOrEmptyString(o.ScheduleDeadlineTime),
	}
	return p
}

// InstanceToProto converts a Instance resource to its proto representation.
func InstanceToProto(resource *redis.Instance) *redispb.RedisInstance {
	p := &redispb.RedisInstance{
		Name:                   dcl.ValueOrEmptyString(resource.Name),
		DisplayName:            dcl.ValueOrEmptyString(resource.DisplayName),
		LocationId:             dcl.ValueOrEmptyString(resource.LocationId),
		AlternativeLocationId:  dcl.ValueOrEmptyString(resource.AlternativeLocationId),
		RedisVersion:           dcl.ValueOrEmptyString(resource.RedisVersion),
		ReservedIpRange:        dcl.ValueOrEmptyString(resource.ReservedIPRange),
		Host:                   dcl.ValueOrEmptyString(resource.Host),
		Port:                   dcl.ValueOrEmptyInt64(resource.Port),
		CurrentLocationId:      dcl.ValueOrEmptyString(resource.CurrentLocationId),
		CreateTime:             dcl.ValueOrEmptyString(resource.CreateTime),
		State:                  RedisInstanceStateEnumToProto(resource.State),
		StatusMessage:          dcl.ValueOrEmptyString(resource.StatusMessage),
		Tier:                   RedisInstanceTierEnumToProto(resource.Tier),
		MemorySizeGb:           dcl.ValueOrEmptyInt64(resource.MemorySizeGb),
		AuthorizedNetwork:      dcl.ValueOrEmptyString(resource.AuthorizedNetwork),
		PersistenceIamIdentity: dcl.ValueOrEmptyString(resource.PersistenceIamIdentity),
		ConnectMode:            RedisInstanceConnectModeEnumToProto(resource.ConnectMode),
		AuthEnabled:            dcl.ValueOrEmptyBool(resource.AuthEnabled),
		TransitEncryptionMode:  RedisInstanceTransitEncryptionModeEnumToProto(resource.TransitEncryptionMode),
		MaintenancePolicy:      RedisInstanceMaintenancePolicyToProto(resource.MaintenancePolicy),
		MaintenanceSchedule:    RedisInstanceMaintenanceScheduleToProto(resource.MaintenanceSchedule),
		Project:                dcl.ValueOrEmptyString(resource.Project),
		Location:               dcl.ValueOrEmptyString(resource.Location),
	}
	for _, r := range resource.ServerCaCerts {
		p.ServerCaCerts = append(p.ServerCaCerts, RedisInstanceServerCaCertsToProto(&r))
	}

	return p
}

// ApplyInstance handles the gRPC request by passing it to the underlying Instance Apply() method.
func (s *InstanceServer) applyInstance(ctx context.Context, c *redis.Client, request *redispb.ApplyRedisInstanceRequest) (*redispb.RedisInstance, error) {
	p := ProtoToInstance(request.GetResource())
	res, err := c.ApplyInstance(ctx, p)
	if err != nil {
		return nil, err
	}
	r := InstanceToProto(res)
	return r, nil
}

// ApplyInstance handles the gRPC request by passing it to the underlying Instance Apply() method.
func (s *InstanceServer) ApplyRedisInstance(ctx context.Context, request *redispb.ApplyRedisInstanceRequest) (*redispb.RedisInstance, error) {
	cl, err := createConfigInstance(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyInstance(ctx, cl, request)
}

// DeleteInstance handles the gRPC request by passing it to the underlying Instance Delete() method.
func (s *InstanceServer) DeleteRedisInstance(ctx context.Context, request *redispb.DeleteRedisInstanceRequest) (*emptypb.Empty, error) {

	cl, err := createConfigInstance(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteInstance(ctx, ProtoToInstance(request.GetResource()))

}

// ListRedisInstance handles the gRPC request by passing it to the underlying InstanceList() method.
func (s *InstanceServer) ListRedisInstance(ctx context.Context, request *redispb.ListRedisInstanceRequest) (*redispb.ListRedisInstanceResponse, error) {
	cl, err := createConfigInstance(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListInstance(ctx, request.Project, request.Location)
	if err != nil {
		return nil, err
	}
	var protos []*redispb.RedisInstance
	for _, r := range resources.Items {
		rp := InstanceToProto(r)
		protos = append(protos, rp)
	}
	return &redispb.ListRedisInstanceResponse{Items: protos}, nil
}

func createConfigInstance(ctx context.Context, service_account_file string) (*redis.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return redis.NewClient(conf), nil
}
