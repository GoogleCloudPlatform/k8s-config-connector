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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/compute/beta/compute_beta_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/compute/beta"
)

// Server implements the gRPC interface for Autoscaler.
type AutoscalerServer struct{}

// ProtoToAutoscalerAutoscalingPolicyCustomMetricUtilizationsUtilizationTargetTypeEnum converts a AutoscalerAutoscalingPolicyCustomMetricUtilizationsUtilizationTargetTypeEnum enum from its proto representation.
func ProtoToComputeBetaAutoscalerAutoscalingPolicyCustomMetricUtilizationsUtilizationTargetTypeEnum(e betapb.ComputeBetaAutoscalerAutoscalingPolicyCustomMetricUtilizationsUtilizationTargetTypeEnum) *beta.AutoscalerAutoscalingPolicyCustomMetricUtilizationsUtilizationTargetTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaAutoscalerAutoscalingPolicyCustomMetricUtilizationsUtilizationTargetTypeEnum_name[int32(e)]; ok {
		e := beta.AutoscalerAutoscalingPolicyCustomMetricUtilizationsUtilizationTargetTypeEnum(n[len("ComputeBetaAutoscalerAutoscalingPolicyCustomMetricUtilizationsUtilizationTargetTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToAutoscalerAutoscalingPolicyModeEnum converts a AutoscalerAutoscalingPolicyModeEnum enum from its proto representation.
func ProtoToComputeBetaAutoscalerAutoscalingPolicyModeEnum(e betapb.ComputeBetaAutoscalerAutoscalingPolicyModeEnum) *beta.AutoscalerAutoscalingPolicyModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaAutoscalerAutoscalingPolicyModeEnum_name[int32(e)]; ok {
		e := beta.AutoscalerAutoscalingPolicyModeEnum(n[len("ComputeBetaAutoscalerAutoscalingPolicyModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToAutoscalerStatusEnum converts a AutoscalerStatusEnum enum from its proto representation.
func ProtoToComputeBetaAutoscalerStatusEnum(e betapb.ComputeBetaAutoscalerStatusEnum) *beta.AutoscalerStatusEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaAutoscalerStatusEnum_name[int32(e)]; ok {
		e := beta.AutoscalerStatusEnum(n[len("ComputeBetaAutoscalerStatusEnum"):])
		return &e
	}
	return nil
}

// ProtoToAutoscalerStatusDetailsTypeEnum converts a AutoscalerStatusDetailsTypeEnum enum from its proto representation.
func ProtoToComputeBetaAutoscalerStatusDetailsTypeEnum(e betapb.ComputeBetaAutoscalerStatusDetailsTypeEnum) *beta.AutoscalerStatusDetailsTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaAutoscalerStatusDetailsTypeEnum_name[int32(e)]; ok {
		e := beta.AutoscalerStatusDetailsTypeEnum(n[len("ComputeBetaAutoscalerStatusDetailsTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToAutoscalerAutoscalingPolicy converts a AutoscalerAutoscalingPolicy resource from its proto representation.
func ProtoToComputeBetaAutoscalerAutoscalingPolicy(p *betapb.ComputeBetaAutoscalerAutoscalingPolicy) *beta.AutoscalerAutoscalingPolicy {
	if p == nil {
		return nil
	}
	obj := &beta.AutoscalerAutoscalingPolicy{
		MinNumReplicas:           dcl.Int64OrNil(p.MinNumReplicas),
		MaxNumReplicas:           dcl.Int64OrNil(p.MaxNumReplicas),
		ScaleInControl:           ProtoToComputeBetaAutoscalerAutoscalingPolicyScaleInControl(p.GetScaleInControl()),
		CoolDownPeriodSec:        dcl.Int64OrNil(p.CoolDownPeriodSec),
		CpuUtilization:           ProtoToComputeBetaAutoscalerAutoscalingPolicyCpuUtilization(p.GetCpuUtilization()),
		LoadBalancingUtilization: ProtoToComputeBetaAutoscalerAutoscalingPolicyLoadBalancingUtilization(p.GetLoadBalancingUtilization()),
		Mode:                     ProtoToComputeBetaAutoscalerAutoscalingPolicyModeEnum(p.GetMode()),
	}
	for _, r := range p.GetCustomMetricUtilizations() {
		obj.CustomMetricUtilizations = append(obj.CustomMetricUtilizations, *ProtoToComputeBetaAutoscalerAutoscalingPolicyCustomMetricUtilizations(r))
	}
	return obj
}

// ProtoToAutoscalerAutoscalingPolicyScaleInControl converts a AutoscalerAutoscalingPolicyScaleInControl resource from its proto representation.
func ProtoToComputeBetaAutoscalerAutoscalingPolicyScaleInControl(p *betapb.ComputeBetaAutoscalerAutoscalingPolicyScaleInControl) *beta.AutoscalerAutoscalingPolicyScaleInControl {
	if p == nil {
		return nil
	}
	obj := &beta.AutoscalerAutoscalingPolicyScaleInControl{
		MaxScaledInReplicas: ProtoToComputeBetaAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas(p.GetMaxScaledInReplicas()),
		TimeWindowSec:       dcl.Int64OrNil(p.TimeWindowSec),
	}
	return obj
}

// ProtoToAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas converts a AutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas resource from its proto representation.
func ProtoToComputeBetaAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas(p *betapb.ComputeBetaAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas) *beta.AutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas {
	if p == nil {
		return nil
	}
	obj := &beta.AutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas{
		Fixed:      dcl.Int64OrNil(p.Fixed),
		Percent:    dcl.Int64OrNil(p.Percent),
		Calculated: dcl.Int64OrNil(p.Calculated),
	}
	return obj
}

// ProtoToAutoscalerAutoscalingPolicyCpuUtilization converts a AutoscalerAutoscalingPolicyCpuUtilization resource from its proto representation.
func ProtoToComputeBetaAutoscalerAutoscalingPolicyCpuUtilization(p *betapb.ComputeBetaAutoscalerAutoscalingPolicyCpuUtilization) *beta.AutoscalerAutoscalingPolicyCpuUtilization {
	if p == nil {
		return nil
	}
	obj := &beta.AutoscalerAutoscalingPolicyCpuUtilization{
		UtilizationTarget: dcl.Float64OrNil(p.UtilizationTarget),
	}
	return obj
}

// ProtoToAutoscalerAutoscalingPolicyCustomMetricUtilizations converts a AutoscalerAutoscalingPolicyCustomMetricUtilizations resource from its proto representation.
func ProtoToComputeBetaAutoscalerAutoscalingPolicyCustomMetricUtilizations(p *betapb.ComputeBetaAutoscalerAutoscalingPolicyCustomMetricUtilizations) *beta.AutoscalerAutoscalingPolicyCustomMetricUtilizations {
	if p == nil {
		return nil
	}
	obj := &beta.AutoscalerAutoscalingPolicyCustomMetricUtilizations{
		Metric:                   dcl.StringOrNil(p.Metric),
		UtilizationTarget:        dcl.Float64OrNil(p.UtilizationTarget),
		UtilizationTargetType:    ProtoToComputeBetaAutoscalerAutoscalingPolicyCustomMetricUtilizationsUtilizationTargetTypeEnum(p.GetUtilizationTargetType()),
		Filter:                   dcl.StringOrNil(p.Filter),
		SingleInstanceAssignment: dcl.Float64OrNil(p.SingleInstanceAssignment),
	}
	return obj
}

// ProtoToAutoscalerAutoscalingPolicyLoadBalancingUtilization converts a AutoscalerAutoscalingPolicyLoadBalancingUtilization resource from its proto representation.
func ProtoToComputeBetaAutoscalerAutoscalingPolicyLoadBalancingUtilization(p *betapb.ComputeBetaAutoscalerAutoscalingPolicyLoadBalancingUtilization) *beta.AutoscalerAutoscalingPolicyLoadBalancingUtilization {
	if p == nil {
		return nil
	}
	obj := &beta.AutoscalerAutoscalingPolicyLoadBalancingUtilization{
		UtilizationTarget: dcl.Float64OrNil(p.UtilizationTarget),
	}
	return obj
}

// ProtoToAutoscalerStatusDetails converts a AutoscalerStatusDetails resource from its proto representation.
func ProtoToComputeBetaAutoscalerStatusDetails(p *betapb.ComputeBetaAutoscalerStatusDetails) *beta.AutoscalerStatusDetails {
	if p == nil {
		return nil
	}
	obj := &beta.AutoscalerStatusDetails{
		Message: dcl.StringOrNil(p.Message),
		Type:    ProtoToComputeBetaAutoscalerStatusDetailsTypeEnum(p.GetType()),
	}
	return obj
}

// ProtoToAutoscaler converts a Autoscaler resource from its proto representation.
func ProtoToAutoscaler(p *betapb.ComputeBetaAutoscaler) *beta.Autoscaler {
	obj := &beta.Autoscaler{
		Id:                dcl.Int64OrNil(p.Id),
		Name:              dcl.StringOrNil(p.Name),
		Description:       dcl.StringOrNil(p.Description),
		Target:            dcl.StringOrNil(p.Target),
		AutoscalingPolicy: ProtoToComputeBetaAutoscalerAutoscalingPolicy(p.GetAutoscalingPolicy()),
		Zone:              dcl.StringOrNil(p.Zone),
		Region:            dcl.StringOrNil(p.Region),
		SelfLink:          dcl.StringOrNil(p.SelfLink),
		Status:            ProtoToComputeBetaAutoscalerStatusEnum(p.GetStatus()),
		RecommendedSize:   dcl.Int64OrNil(p.RecommendedSize),
		SelfLinkWithId:    dcl.StringOrNil(p.SelfLinkWithId),
		Project:           dcl.StringOrNil(p.Project),
		CreationTimestamp: dcl.StringOrNil(p.CreationTimestamp),
		Location:          dcl.StringOrNil(p.Location),
	}
	for _, r := range p.GetStatusDetails() {
		obj.StatusDetails = append(obj.StatusDetails, *ProtoToComputeBetaAutoscalerStatusDetails(r))
	}
	return obj
}

// AutoscalerAutoscalingPolicyCustomMetricUtilizationsUtilizationTargetTypeEnumToProto converts a AutoscalerAutoscalingPolicyCustomMetricUtilizationsUtilizationTargetTypeEnum enum to its proto representation.
func ComputeBetaAutoscalerAutoscalingPolicyCustomMetricUtilizationsUtilizationTargetTypeEnumToProto(e *beta.AutoscalerAutoscalingPolicyCustomMetricUtilizationsUtilizationTargetTypeEnum) betapb.ComputeBetaAutoscalerAutoscalingPolicyCustomMetricUtilizationsUtilizationTargetTypeEnum {
	if e == nil {
		return betapb.ComputeBetaAutoscalerAutoscalingPolicyCustomMetricUtilizationsUtilizationTargetTypeEnum(0)
	}
	if v, ok := betapb.ComputeBetaAutoscalerAutoscalingPolicyCustomMetricUtilizationsUtilizationTargetTypeEnum_value["AutoscalerAutoscalingPolicyCustomMetricUtilizationsUtilizationTargetTypeEnum"+string(*e)]; ok {
		return betapb.ComputeBetaAutoscalerAutoscalingPolicyCustomMetricUtilizationsUtilizationTargetTypeEnum(v)
	}
	return betapb.ComputeBetaAutoscalerAutoscalingPolicyCustomMetricUtilizationsUtilizationTargetTypeEnum(0)
}

// AutoscalerAutoscalingPolicyModeEnumToProto converts a AutoscalerAutoscalingPolicyModeEnum enum to its proto representation.
func ComputeBetaAutoscalerAutoscalingPolicyModeEnumToProto(e *beta.AutoscalerAutoscalingPolicyModeEnum) betapb.ComputeBetaAutoscalerAutoscalingPolicyModeEnum {
	if e == nil {
		return betapb.ComputeBetaAutoscalerAutoscalingPolicyModeEnum(0)
	}
	if v, ok := betapb.ComputeBetaAutoscalerAutoscalingPolicyModeEnum_value["AutoscalerAutoscalingPolicyModeEnum"+string(*e)]; ok {
		return betapb.ComputeBetaAutoscalerAutoscalingPolicyModeEnum(v)
	}
	return betapb.ComputeBetaAutoscalerAutoscalingPolicyModeEnum(0)
}

// AutoscalerStatusEnumToProto converts a AutoscalerStatusEnum enum to its proto representation.
func ComputeBetaAutoscalerStatusEnumToProto(e *beta.AutoscalerStatusEnum) betapb.ComputeBetaAutoscalerStatusEnum {
	if e == nil {
		return betapb.ComputeBetaAutoscalerStatusEnum(0)
	}
	if v, ok := betapb.ComputeBetaAutoscalerStatusEnum_value["AutoscalerStatusEnum"+string(*e)]; ok {
		return betapb.ComputeBetaAutoscalerStatusEnum(v)
	}
	return betapb.ComputeBetaAutoscalerStatusEnum(0)
}

// AutoscalerStatusDetailsTypeEnumToProto converts a AutoscalerStatusDetailsTypeEnum enum to its proto representation.
func ComputeBetaAutoscalerStatusDetailsTypeEnumToProto(e *beta.AutoscalerStatusDetailsTypeEnum) betapb.ComputeBetaAutoscalerStatusDetailsTypeEnum {
	if e == nil {
		return betapb.ComputeBetaAutoscalerStatusDetailsTypeEnum(0)
	}
	if v, ok := betapb.ComputeBetaAutoscalerStatusDetailsTypeEnum_value["AutoscalerStatusDetailsTypeEnum"+string(*e)]; ok {
		return betapb.ComputeBetaAutoscalerStatusDetailsTypeEnum(v)
	}
	return betapb.ComputeBetaAutoscalerStatusDetailsTypeEnum(0)
}

// AutoscalerAutoscalingPolicyToProto converts a AutoscalerAutoscalingPolicy resource to its proto representation.
func ComputeBetaAutoscalerAutoscalingPolicyToProto(o *beta.AutoscalerAutoscalingPolicy) *betapb.ComputeBetaAutoscalerAutoscalingPolicy {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaAutoscalerAutoscalingPolicy{
		MinNumReplicas:           dcl.ValueOrEmptyInt64(o.MinNumReplicas),
		MaxNumReplicas:           dcl.ValueOrEmptyInt64(o.MaxNumReplicas),
		ScaleInControl:           ComputeBetaAutoscalerAutoscalingPolicyScaleInControlToProto(o.ScaleInControl),
		CoolDownPeriodSec:        dcl.ValueOrEmptyInt64(o.CoolDownPeriodSec),
		CpuUtilization:           ComputeBetaAutoscalerAutoscalingPolicyCpuUtilizationToProto(o.CpuUtilization),
		LoadBalancingUtilization: ComputeBetaAutoscalerAutoscalingPolicyLoadBalancingUtilizationToProto(o.LoadBalancingUtilization),
		Mode:                     ComputeBetaAutoscalerAutoscalingPolicyModeEnumToProto(o.Mode),
	}
	for _, r := range o.CustomMetricUtilizations {
		p.CustomMetricUtilizations = append(p.CustomMetricUtilizations, ComputeBetaAutoscalerAutoscalingPolicyCustomMetricUtilizationsToProto(&r))
	}
	return p
}

// AutoscalerAutoscalingPolicyScaleInControlToProto converts a AutoscalerAutoscalingPolicyScaleInControl resource to its proto representation.
func ComputeBetaAutoscalerAutoscalingPolicyScaleInControlToProto(o *beta.AutoscalerAutoscalingPolicyScaleInControl) *betapb.ComputeBetaAutoscalerAutoscalingPolicyScaleInControl {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaAutoscalerAutoscalingPolicyScaleInControl{
		MaxScaledInReplicas: ComputeBetaAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicasToProto(o.MaxScaledInReplicas),
		TimeWindowSec:       dcl.ValueOrEmptyInt64(o.TimeWindowSec),
	}
	return p
}

// AutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicasToProto converts a AutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas resource to its proto representation.
func ComputeBetaAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicasToProto(o *beta.AutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas) *betapb.ComputeBetaAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas{
		Fixed:      dcl.ValueOrEmptyInt64(o.Fixed),
		Percent:    dcl.ValueOrEmptyInt64(o.Percent),
		Calculated: dcl.ValueOrEmptyInt64(o.Calculated),
	}
	return p
}

// AutoscalerAutoscalingPolicyCpuUtilizationToProto converts a AutoscalerAutoscalingPolicyCpuUtilization resource to its proto representation.
func ComputeBetaAutoscalerAutoscalingPolicyCpuUtilizationToProto(o *beta.AutoscalerAutoscalingPolicyCpuUtilization) *betapb.ComputeBetaAutoscalerAutoscalingPolicyCpuUtilization {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaAutoscalerAutoscalingPolicyCpuUtilization{
		UtilizationTarget: dcl.ValueOrEmptyDouble(o.UtilizationTarget),
	}
	return p
}

// AutoscalerAutoscalingPolicyCustomMetricUtilizationsToProto converts a AutoscalerAutoscalingPolicyCustomMetricUtilizations resource to its proto representation.
func ComputeBetaAutoscalerAutoscalingPolicyCustomMetricUtilizationsToProto(o *beta.AutoscalerAutoscalingPolicyCustomMetricUtilizations) *betapb.ComputeBetaAutoscalerAutoscalingPolicyCustomMetricUtilizations {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaAutoscalerAutoscalingPolicyCustomMetricUtilizations{
		Metric:                   dcl.ValueOrEmptyString(o.Metric),
		UtilizationTarget:        dcl.ValueOrEmptyDouble(o.UtilizationTarget),
		UtilizationTargetType:    ComputeBetaAutoscalerAutoscalingPolicyCustomMetricUtilizationsUtilizationTargetTypeEnumToProto(o.UtilizationTargetType),
		Filter:                   dcl.ValueOrEmptyString(o.Filter),
		SingleInstanceAssignment: dcl.ValueOrEmptyDouble(o.SingleInstanceAssignment),
	}
	return p
}

// AutoscalerAutoscalingPolicyLoadBalancingUtilizationToProto converts a AutoscalerAutoscalingPolicyLoadBalancingUtilization resource to its proto representation.
func ComputeBetaAutoscalerAutoscalingPolicyLoadBalancingUtilizationToProto(o *beta.AutoscalerAutoscalingPolicyLoadBalancingUtilization) *betapb.ComputeBetaAutoscalerAutoscalingPolicyLoadBalancingUtilization {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaAutoscalerAutoscalingPolicyLoadBalancingUtilization{
		UtilizationTarget: dcl.ValueOrEmptyDouble(o.UtilizationTarget),
	}
	return p
}

// AutoscalerStatusDetailsToProto converts a AutoscalerStatusDetails resource to its proto representation.
func ComputeBetaAutoscalerStatusDetailsToProto(o *beta.AutoscalerStatusDetails) *betapb.ComputeBetaAutoscalerStatusDetails {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaAutoscalerStatusDetails{
		Message: dcl.ValueOrEmptyString(o.Message),
		Type:    ComputeBetaAutoscalerStatusDetailsTypeEnumToProto(o.Type),
	}
	return p
}

// AutoscalerToProto converts a Autoscaler resource to its proto representation.
func AutoscalerToProto(resource *beta.Autoscaler) *betapb.ComputeBetaAutoscaler {
	p := &betapb.ComputeBetaAutoscaler{
		Id:                dcl.ValueOrEmptyInt64(resource.Id),
		Name:              dcl.ValueOrEmptyString(resource.Name),
		Description:       dcl.ValueOrEmptyString(resource.Description),
		Target:            dcl.ValueOrEmptyString(resource.Target),
		AutoscalingPolicy: ComputeBetaAutoscalerAutoscalingPolicyToProto(resource.AutoscalingPolicy),
		Zone:              dcl.ValueOrEmptyString(resource.Zone),
		Region:            dcl.ValueOrEmptyString(resource.Region),
		SelfLink:          dcl.ValueOrEmptyString(resource.SelfLink),
		Status:            ComputeBetaAutoscalerStatusEnumToProto(resource.Status),
		RecommendedSize:   dcl.ValueOrEmptyInt64(resource.RecommendedSize),
		SelfLinkWithId:    dcl.ValueOrEmptyString(resource.SelfLinkWithId),
		Project:           dcl.ValueOrEmptyString(resource.Project),
		CreationTimestamp: dcl.ValueOrEmptyString(resource.CreationTimestamp),
		Location:          dcl.ValueOrEmptyString(resource.Location),
	}
	for _, r := range resource.StatusDetails {
		p.StatusDetails = append(p.StatusDetails, ComputeBetaAutoscalerStatusDetailsToProto(&r))
	}

	return p
}

// ApplyAutoscaler handles the gRPC request by passing it to the underlying Autoscaler Apply() method.
func (s *AutoscalerServer) applyAutoscaler(ctx context.Context, c *beta.Client, request *betapb.ApplyComputeBetaAutoscalerRequest) (*betapb.ComputeBetaAutoscaler, error) {
	p := ProtoToAutoscaler(request.GetResource())
	res, err := c.ApplyAutoscaler(ctx, p)
	if err != nil {
		return nil, err
	}
	r := AutoscalerToProto(res)
	return r, nil
}

// ApplyAutoscaler handles the gRPC request by passing it to the underlying Autoscaler Apply() method.
func (s *AutoscalerServer) ApplyComputeBetaAutoscaler(ctx context.Context, request *betapb.ApplyComputeBetaAutoscalerRequest) (*betapb.ComputeBetaAutoscaler, error) {
	cl, err := createConfigAutoscaler(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyAutoscaler(ctx, cl, request)
}

// DeleteAutoscaler handles the gRPC request by passing it to the underlying Autoscaler Delete() method.
func (s *AutoscalerServer) DeleteComputeBetaAutoscaler(ctx context.Context, request *betapb.DeleteComputeBetaAutoscalerRequest) (*emptypb.Empty, error) {

	cl, err := createConfigAutoscaler(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteAutoscaler(ctx, ProtoToAutoscaler(request.GetResource()))

}

// ListComputeBetaAutoscaler handles the gRPC request by passing it to the underlying AutoscalerList() method.
func (s *AutoscalerServer) ListComputeBetaAutoscaler(ctx context.Context, request *betapb.ListComputeBetaAutoscalerRequest) (*betapb.ListComputeBetaAutoscalerResponse, error) {
	cl, err := createConfigAutoscaler(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListAutoscaler(ctx, request.Project, request.Location)
	if err != nil {
		return nil, err
	}
	var protos []*betapb.ComputeBetaAutoscaler
	for _, r := range resources.Items {
		rp := AutoscalerToProto(r)
		protos = append(protos, rp)
	}
	return &betapb.ListComputeBetaAutoscalerResponse{Items: protos}, nil
}

func createConfigAutoscaler(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
