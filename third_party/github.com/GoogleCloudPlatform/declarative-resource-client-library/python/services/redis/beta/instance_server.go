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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/redis/beta/redis_beta_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/redis/beta"
)

// Server implements the gRPC interface for Instance.
type InstanceServer struct{}

// ProtoToInstanceStateEnum converts a InstanceStateEnum enum from its proto representation.
func ProtoToRedisBetaInstanceStateEnum(e betapb.RedisBetaInstanceStateEnum) *beta.InstanceStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.RedisBetaInstanceStateEnum_name[int32(e)]; ok {
		e := beta.InstanceStateEnum(n[len("RedisBetaInstanceStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceTierEnum converts a InstanceTierEnum enum from its proto representation.
func ProtoToRedisBetaInstanceTierEnum(e betapb.RedisBetaInstanceTierEnum) *beta.InstanceTierEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.RedisBetaInstanceTierEnum_name[int32(e)]; ok {
		e := beta.InstanceTierEnum(n[len("RedisBetaInstanceTierEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceConnectModeEnum converts a InstanceConnectModeEnum enum from its proto representation.
func ProtoToRedisBetaInstanceConnectModeEnum(e betapb.RedisBetaInstanceConnectModeEnum) *beta.InstanceConnectModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.RedisBetaInstanceConnectModeEnum_name[int32(e)]; ok {
		e := beta.InstanceConnectModeEnum(n[len("RedisBetaInstanceConnectModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceTransitEncryptionModeEnum converts a InstanceTransitEncryptionModeEnum enum from its proto representation.
func ProtoToRedisBetaInstanceTransitEncryptionModeEnum(e betapb.RedisBetaInstanceTransitEncryptionModeEnum) *beta.InstanceTransitEncryptionModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.RedisBetaInstanceTransitEncryptionModeEnum_name[int32(e)]; ok {
		e := beta.InstanceTransitEncryptionModeEnum(n[len("RedisBetaInstanceTransitEncryptionModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceMaintenancePolicyWeeklyMaintenanceWindowDayEnum converts a InstanceMaintenancePolicyWeeklyMaintenanceWindowDayEnum enum from its proto representation.
func ProtoToRedisBetaInstanceMaintenancePolicyWeeklyMaintenanceWindowDayEnum(e betapb.RedisBetaInstanceMaintenancePolicyWeeklyMaintenanceWindowDayEnum) *beta.InstanceMaintenancePolicyWeeklyMaintenanceWindowDayEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.RedisBetaInstanceMaintenancePolicyWeeklyMaintenanceWindowDayEnum_name[int32(e)]; ok {
		e := beta.InstanceMaintenancePolicyWeeklyMaintenanceWindowDayEnum(n[len("RedisBetaInstanceMaintenancePolicyWeeklyMaintenanceWindowDayEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceServerCaCerts converts a InstanceServerCaCerts resource from its proto representation.
func ProtoToRedisBetaInstanceServerCaCerts(p *betapb.RedisBetaInstanceServerCaCerts) *beta.InstanceServerCaCerts {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceServerCaCerts{
		SerialNumber:    dcl.StringOrNil(p.SerialNumber),
		Cert:            dcl.StringOrNil(p.Cert),
		CreateTime:      dcl.StringOrNil(p.GetCreateTime()),
		ExpireTime:      dcl.StringOrNil(p.GetExpireTime()),
		Sha1Fingerprint: dcl.StringOrNil(p.Sha1Fingerprint),
	}
	return obj
}

// ProtoToInstanceMaintenancePolicy converts a InstanceMaintenancePolicy resource from its proto representation.
func ProtoToRedisBetaInstanceMaintenancePolicy(p *betapb.RedisBetaInstanceMaintenancePolicy) *beta.InstanceMaintenancePolicy {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceMaintenancePolicy{
		CreateTime:  dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:  dcl.StringOrNil(p.GetUpdateTime()),
		Description: dcl.StringOrNil(p.Description),
	}
	for _, r := range p.GetWeeklyMaintenanceWindow() {
		obj.WeeklyMaintenanceWindow = append(obj.WeeklyMaintenanceWindow, *ProtoToRedisBetaInstanceMaintenancePolicyWeeklyMaintenanceWindow(r))
	}
	return obj
}

// ProtoToInstanceMaintenancePolicyWeeklyMaintenanceWindow converts a InstanceMaintenancePolicyWeeklyMaintenanceWindow resource from its proto representation.
func ProtoToRedisBetaInstanceMaintenancePolicyWeeklyMaintenanceWindow(p *betapb.RedisBetaInstanceMaintenancePolicyWeeklyMaintenanceWindow) *beta.InstanceMaintenancePolicyWeeklyMaintenanceWindow {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceMaintenancePolicyWeeklyMaintenanceWindow{
		Day:       ProtoToRedisBetaInstanceMaintenancePolicyWeeklyMaintenanceWindowDayEnum(p.GetDay()),
		StartTime: ProtoToRedisBetaInstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime(p.GetStartTime()),
		Duration:  dcl.StringOrNil(p.Duration),
	}
	return obj
}

// ProtoToInstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime converts a InstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime resource from its proto representation.
func ProtoToRedisBetaInstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime(p *betapb.RedisBetaInstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime) *beta.InstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime{
		Hours:   dcl.Int64OrNil(p.Hours),
		Minutes: dcl.Int64OrNil(p.Minutes),
		Seconds: dcl.Int64OrNil(p.Seconds),
		Nanos:   dcl.Int64OrNil(p.Nanos),
	}
	return obj
}

// ProtoToInstanceMaintenanceSchedule converts a InstanceMaintenanceSchedule resource from its proto representation.
func ProtoToRedisBetaInstanceMaintenanceSchedule(p *betapb.RedisBetaInstanceMaintenanceSchedule) *beta.InstanceMaintenanceSchedule {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceMaintenanceSchedule{
		StartTime:            dcl.StringOrNil(p.GetStartTime()),
		EndTime:              dcl.StringOrNil(p.GetEndTime()),
		CanReschedule:        dcl.Bool(p.CanReschedule),
		ScheduleDeadlineTime: dcl.StringOrNil(p.GetScheduleDeadlineTime()),
	}
	return obj
}

// ProtoToInstance converts a Instance resource from its proto representation.
func ProtoToInstance(p *betapb.RedisBetaInstance) *beta.Instance {
	obj := &beta.Instance{
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
		State:                  ProtoToRedisBetaInstanceStateEnum(p.GetState()),
		StatusMessage:          dcl.StringOrNil(p.StatusMessage),
		Tier:                   ProtoToRedisBetaInstanceTierEnum(p.GetTier()),
		MemorySizeGb:           dcl.Int64OrNil(p.MemorySizeGb),
		AuthorizedNetwork:      dcl.StringOrNil(p.AuthorizedNetwork),
		PersistenceIamIdentity: dcl.StringOrNil(p.PersistenceIamIdentity),
		ConnectMode:            ProtoToRedisBetaInstanceConnectModeEnum(p.GetConnectMode()),
		AuthEnabled:            dcl.Bool(p.AuthEnabled),
		TransitEncryptionMode:  ProtoToRedisBetaInstanceTransitEncryptionModeEnum(p.GetTransitEncryptionMode()),
		MaintenancePolicy:      ProtoToRedisBetaInstanceMaintenancePolicy(p.GetMaintenancePolicy()),
		MaintenanceSchedule:    ProtoToRedisBetaInstanceMaintenanceSchedule(p.GetMaintenanceSchedule()),
		Project:                dcl.StringOrNil(p.Project),
		Location:               dcl.StringOrNil(p.Location),
	}
	for _, r := range p.GetServerCaCerts() {
		obj.ServerCaCerts = append(obj.ServerCaCerts, *ProtoToRedisBetaInstanceServerCaCerts(r))
	}
	return obj
}

// InstanceStateEnumToProto converts a InstanceStateEnum enum to its proto representation.
func RedisBetaInstanceStateEnumToProto(e *beta.InstanceStateEnum) betapb.RedisBetaInstanceStateEnum {
	if e == nil {
		return betapb.RedisBetaInstanceStateEnum(0)
	}
	if v, ok := betapb.RedisBetaInstanceStateEnum_value["InstanceStateEnum"+string(*e)]; ok {
		return betapb.RedisBetaInstanceStateEnum(v)
	}
	return betapb.RedisBetaInstanceStateEnum(0)
}

// InstanceTierEnumToProto converts a InstanceTierEnum enum to its proto representation.
func RedisBetaInstanceTierEnumToProto(e *beta.InstanceTierEnum) betapb.RedisBetaInstanceTierEnum {
	if e == nil {
		return betapb.RedisBetaInstanceTierEnum(0)
	}
	if v, ok := betapb.RedisBetaInstanceTierEnum_value["InstanceTierEnum"+string(*e)]; ok {
		return betapb.RedisBetaInstanceTierEnum(v)
	}
	return betapb.RedisBetaInstanceTierEnum(0)
}

// InstanceConnectModeEnumToProto converts a InstanceConnectModeEnum enum to its proto representation.
func RedisBetaInstanceConnectModeEnumToProto(e *beta.InstanceConnectModeEnum) betapb.RedisBetaInstanceConnectModeEnum {
	if e == nil {
		return betapb.RedisBetaInstanceConnectModeEnum(0)
	}
	if v, ok := betapb.RedisBetaInstanceConnectModeEnum_value["InstanceConnectModeEnum"+string(*e)]; ok {
		return betapb.RedisBetaInstanceConnectModeEnum(v)
	}
	return betapb.RedisBetaInstanceConnectModeEnum(0)
}

// InstanceTransitEncryptionModeEnumToProto converts a InstanceTransitEncryptionModeEnum enum to its proto representation.
func RedisBetaInstanceTransitEncryptionModeEnumToProto(e *beta.InstanceTransitEncryptionModeEnum) betapb.RedisBetaInstanceTransitEncryptionModeEnum {
	if e == nil {
		return betapb.RedisBetaInstanceTransitEncryptionModeEnum(0)
	}
	if v, ok := betapb.RedisBetaInstanceTransitEncryptionModeEnum_value["InstanceTransitEncryptionModeEnum"+string(*e)]; ok {
		return betapb.RedisBetaInstanceTransitEncryptionModeEnum(v)
	}
	return betapb.RedisBetaInstanceTransitEncryptionModeEnum(0)
}

// InstanceMaintenancePolicyWeeklyMaintenanceWindowDayEnumToProto converts a InstanceMaintenancePolicyWeeklyMaintenanceWindowDayEnum enum to its proto representation.
func RedisBetaInstanceMaintenancePolicyWeeklyMaintenanceWindowDayEnumToProto(e *beta.InstanceMaintenancePolicyWeeklyMaintenanceWindowDayEnum) betapb.RedisBetaInstanceMaintenancePolicyWeeklyMaintenanceWindowDayEnum {
	if e == nil {
		return betapb.RedisBetaInstanceMaintenancePolicyWeeklyMaintenanceWindowDayEnum(0)
	}
	if v, ok := betapb.RedisBetaInstanceMaintenancePolicyWeeklyMaintenanceWindowDayEnum_value["InstanceMaintenancePolicyWeeklyMaintenanceWindowDayEnum"+string(*e)]; ok {
		return betapb.RedisBetaInstanceMaintenancePolicyWeeklyMaintenanceWindowDayEnum(v)
	}
	return betapb.RedisBetaInstanceMaintenancePolicyWeeklyMaintenanceWindowDayEnum(0)
}

// InstanceServerCaCertsToProto converts a InstanceServerCaCerts resource to its proto representation.
func RedisBetaInstanceServerCaCertsToProto(o *beta.InstanceServerCaCerts) *betapb.RedisBetaInstanceServerCaCerts {
	if o == nil {
		return nil
	}
	p := &betapb.RedisBetaInstanceServerCaCerts{
		SerialNumber:    dcl.ValueOrEmptyString(o.SerialNumber),
		Cert:            dcl.ValueOrEmptyString(o.Cert),
		CreateTime:      dcl.ValueOrEmptyString(o.CreateTime),
		ExpireTime:      dcl.ValueOrEmptyString(o.ExpireTime),
		Sha1Fingerprint: dcl.ValueOrEmptyString(o.Sha1Fingerprint),
	}
	return p
}

// InstanceMaintenancePolicyToProto converts a InstanceMaintenancePolicy resource to its proto representation.
func RedisBetaInstanceMaintenancePolicyToProto(o *beta.InstanceMaintenancePolicy) *betapb.RedisBetaInstanceMaintenancePolicy {
	if o == nil {
		return nil
	}
	p := &betapb.RedisBetaInstanceMaintenancePolicy{
		CreateTime:  dcl.ValueOrEmptyString(o.CreateTime),
		UpdateTime:  dcl.ValueOrEmptyString(o.UpdateTime),
		Description: dcl.ValueOrEmptyString(o.Description),
	}
	for _, r := range o.WeeklyMaintenanceWindow {
		p.WeeklyMaintenanceWindow = append(p.WeeklyMaintenanceWindow, RedisBetaInstanceMaintenancePolicyWeeklyMaintenanceWindowToProto(&r))
	}
	return p
}

// InstanceMaintenancePolicyWeeklyMaintenanceWindowToProto converts a InstanceMaintenancePolicyWeeklyMaintenanceWindow resource to its proto representation.
func RedisBetaInstanceMaintenancePolicyWeeklyMaintenanceWindowToProto(o *beta.InstanceMaintenancePolicyWeeklyMaintenanceWindow) *betapb.RedisBetaInstanceMaintenancePolicyWeeklyMaintenanceWindow {
	if o == nil {
		return nil
	}
	p := &betapb.RedisBetaInstanceMaintenancePolicyWeeklyMaintenanceWindow{
		Day:       RedisBetaInstanceMaintenancePolicyWeeklyMaintenanceWindowDayEnumToProto(o.Day),
		StartTime: RedisBetaInstanceMaintenancePolicyWeeklyMaintenanceWindowStartTimeToProto(o.StartTime),
		Duration:  dcl.ValueOrEmptyString(o.Duration),
	}
	return p
}

// InstanceMaintenancePolicyWeeklyMaintenanceWindowStartTimeToProto converts a InstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime resource to its proto representation.
func RedisBetaInstanceMaintenancePolicyWeeklyMaintenanceWindowStartTimeToProto(o *beta.InstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime) *betapb.RedisBetaInstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime {
	if o == nil {
		return nil
	}
	p := &betapb.RedisBetaInstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime{
		Hours:   dcl.ValueOrEmptyInt64(o.Hours),
		Minutes: dcl.ValueOrEmptyInt64(o.Minutes),
		Seconds: dcl.ValueOrEmptyInt64(o.Seconds),
		Nanos:   dcl.ValueOrEmptyInt64(o.Nanos),
	}
	return p
}

// InstanceMaintenanceScheduleToProto converts a InstanceMaintenanceSchedule resource to its proto representation.
func RedisBetaInstanceMaintenanceScheduleToProto(o *beta.InstanceMaintenanceSchedule) *betapb.RedisBetaInstanceMaintenanceSchedule {
	if o == nil {
		return nil
	}
	p := &betapb.RedisBetaInstanceMaintenanceSchedule{
		StartTime:            dcl.ValueOrEmptyString(o.StartTime),
		EndTime:              dcl.ValueOrEmptyString(o.EndTime),
		CanReschedule:        dcl.ValueOrEmptyBool(o.CanReschedule),
		ScheduleDeadlineTime: dcl.ValueOrEmptyString(o.ScheduleDeadlineTime),
	}
	return p
}

// InstanceToProto converts a Instance resource to its proto representation.
func InstanceToProto(resource *beta.Instance) *betapb.RedisBetaInstance {
	p := &betapb.RedisBetaInstance{
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
		State:                  RedisBetaInstanceStateEnumToProto(resource.State),
		StatusMessage:          dcl.ValueOrEmptyString(resource.StatusMessage),
		Tier:                   RedisBetaInstanceTierEnumToProto(resource.Tier),
		MemorySizeGb:           dcl.ValueOrEmptyInt64(resource.MemorySizeGb),
		AuthorizedNetwork:      dcl.ValueOrEmptyString(resource.AuthorizedNetwork),
		PersistenceIamIdentity: dcl.ValueOrEmptyString(resource.PersistenceIamIdentity),
		ConnectMode:            RedisBetaInstanceConnectModeEnumToProto(resource.ConnectMode),
		AuthEnabled:            dcl.ValueOrEmptyBool(resource.AuthEnabled),
		TransitEncryptionMode:  RedisBetaInstanceTransitEncryptionModeEnumToProto(resource.TransitEncryptionMode),
		MaintenancePolicy:      RedisBetaInstanceMaintenancePolicyToProto(resource.MaintenancePolicy),
		MaintenanceSchedule:    RedisBetaInstanceMaintenanceScheduleToProto(resource.MaintenanceSchedule),
		Project:                dcl.ValueOrEmptyString(resource.Project),
		Location:               dcl.ValueOrEmptyString(resource.Location),
	}
	for _, r := range resource.ServerCaCerts {
		p.ServerCaCerts = append(p.ServerCaCerts, RedisBetaInstanceServerCaCertsToProto(&r))
	}

	return p
}

// ApplyInstance handles the gRPC request by passing it to the underlying Instance Apply() method.
func (s *InstanceServer) applyInstance(ctx context.Context, c *beta.Client, request *betapb.ApplyRedisBetaInstanceRequest) (*betapb.RedisBetaInstance, error) {
	p := ProtoToInstance(request.GetResource())
	res, err := c.ApplyInstance(ctx, p)
	if err != nil {
		return nil, err
	}
	r := InstanceToProto(res)
	return r, nil
}

// ApplyInstance handles the gRPC request by passing it to the underlying Instance Apply() method.
func (s *InstanceServer) ApplyRedisBetaInstance(ctx context.Context, request *betapb.ApplyRedisBetaInstanceRequest) (*betapb.RedisBetaInstance, error) {
	cl, err := createConfigInstance(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyInstance(ctx, cl, request)
}

// DeleteInstance handles the gRPC request by passing it to the underlying Instance Delete() method.
func (s *InstanceServer) DeleteRedisBetaInstance(ctx context.Context, request *betapb.DeleteRedisBetaInstanceRequest) (*emptypb.Empty, error) {

	cl, err := createConfigInstance(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteInstance(ctx, ProtoToInstance(request.GetResource()))

}

// ListRedisBetaInstance handles the gRPC request by passing it to the underlying InstanceList() method.
func (s *InstanceServer) ListRedisBetaInstance(ctx context.Context, request *betapb.ListRedisBetaInstanceRequest) (*betapb.ListRedisBetaInstanceResponse, error) {
	cl, err := createConfigInstance(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListInstance(ctx, request.Project, request.Location)
	if err != nil {
		return nil, err
	}
	var protos []*betapb.RedisBetaInstance
	for _, r := range resources.Items {
		rp := InstanceToProto(r)
		protos = append(protos, rp)
	}
	return &betapb.ListRedisBetaInstanceResponse{Items: protos}, nil
}

func createConfigInstance(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
