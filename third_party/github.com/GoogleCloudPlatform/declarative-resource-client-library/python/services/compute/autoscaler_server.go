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
	computepb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/compute/compute_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/compute"
)

// Server implements the gRPC interface for Autoscaler.
type AutoscalerServer struct{}

// ProtoToAutoscalerAutoscalingPolicyCustomMetricUtilizationsUtilizationTargetTypeEnum converts a AutoscalerAutoscalingPolicyCustomMetricUtilizationsUtilizationTargetTypeEnum enum from its proto representation.
func ProtoToComputeAutoscalerAutoscalingPolicyCustomMetricUtilizationsUtilizationTargetTypeEnum(e computepb.ComputeAutoscalerAutoscalingPolicyCustomMetricUtilizationsUtilizationTargetTypeEnum) *compute.AutoscalerAutoscalingPolicyCustomMetricUtilizationsUtilizationTargetTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeAutoscalerAutoscalingPolicyCustomMetricUtilizationsUtilizationTargetTypeEnum_name[int32(e)]; ok {
		e := compute.AutoscalerAutoscalingPolicyCustomMetricUtilizationsUtilizationTargetTypeEnum(n[len("ComputeAutoscalerAutoscalingPolicyCustomMetricUtilizationsUtilizationTargetTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToAutoscalerAutoscalingPolicyModeEnum converts a AutoscalerAutoscalingPolicyModeEnum enum from its proto representation.
func ProtoToComputeAutoscalerAutoscalingPolicyModeEnum(e computepb.ComputeAutoscalerAutoscalingPolicyModeEnum) *compute.AutoscalerAutoscalingPolicyModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeAutoscalerAutoscalingPolicyModeEnum_name[int32(e)]; ok {
		e := compute.AutoscalerAutoscalingPolicyModeEnum(n[len("ComputeAutoscalerAutoscalingPolicyModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToAutoscalerStatusEnum converts a AutoscalerStatusEnum enum from its proto representation.
func ProtoToComputeAutoscalerStatusEnum(e computepb.ComputeAutoscalerStatusEnum) *compute.AutoscalerStatusEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeAutoscalerStatusEnum_name[int32(e)]; ok {
		e := compute.AutoscalerStatusEnum(n[len("ComputeAutoscalerStatusEnum"):])
		return &e
	}
	return nil
}

// ProtoToAutoscalerStatusDetailsTypeEnum converts a AutoscalerStatusDetailsTypeEnum enum from its proto representation.
func ProtoToComputeAutoscalerStatusDetailsTypeEnum(e computepb.ComputeAutoscalerStatusDetailsTypeEnum) *compute.AutoscalerStatusDetailsTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeAutoscalerStatusDetailsTypeEnum_name[int32(e)]; ok {
		e := compute.AutoscalerStatusDetailsTypeEnum(n[len("ComputeAutoscalerStatusDetailsTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToAutoscalerAutoscalingPolicy converts a AutoscalerAutoscalingPolicy resource from its proto representation.
func ProtoToComputeAutoscalerAutoscalingPolicy(p *computepb.ComputeAutoscalerAutoscalingPolicy) *compute.AutoscalerAutoscalingPolicy {
	if p == nil {
		return nil
	}
	obj := &compute.AutoscalerAutoscalingPolicy{
		MinNumReplicas:           dcl.Int64OrNil(p.MinNumReplicas),
		MaxNumReplicas:           dcl.Int64OrNil(p.MaxNumReplicas),
		ScaleInControl:           ProtoToComputeAutoscalerAutoscalingPolicyScaleInControl(p.GetScaleInControl()),
		CoolDownPeriodSec:        dcl.Int64OrNil(p.CoolDownPeriodSec),
		CpuUtilization:           ProtoToComputeAutoscalerAutoscalingPolicyCpuUtilization(p.GetCpuUtilization()),
		LoadBalancingUtilization: ProtoToComputeAutoscalerAutoscalingPolicyLoadBalancingUtilization(p.GetLoadBalancingUtilization()),
		Mode:                     ProtoToComputeAutoscalerAutoscalingPolicyModeEnum(p.GetMode()),
	}
	for _, r := range p.GetCustomMetricUtilizations() {
		obj.CustomMetricUtilizations = append(obj.CustomMetricUtilizations, *ProtoToComputeAutoscalerAutoscalingPolicyCustomMetricUtilizations(r))
	}
	return obj
}

// ProtoToAutoscalerAutoscalingPolicyScaleInControl converts a AutoscalerAutoscalingPolicyScaleInControl resource from its proto representation.
func ProtoToComputeAutoscalerAutoscalingPolicyScaleInControl(p *computepb.ComputeAutoscalerAutoscalingPolicyScaleInControl) *compute.AutoscalerAutoscalingPolicyScaleInControl {
	if p == nil {
		return nil
	}
	obj := &compute.AutoscalerAutoscalingPolicyScaleInControl{
		MaxScaledInReplicas: ProtoToComputeAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas(p.GetMaxScaledInReplicas()),
		TimeWindowSec:       dcl.Int64OrNil(p.TimeWindowSec),
	}
	return obj
}

// ProtoToAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas converts a AutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas resource from its proto representation.
func ProtoToComputeAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas(p *computepb.ComputeAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas) *compute.AutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas {
	if p == nil {
		return nil
	}
	obj := &compute.AutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas{
		Fixed:      dcl.Int64OrNil(p.Fixed),
		Percent:    dcl.Int64OrNil(p.Percent),
		Calculated: dcl.Int64OrNil(p.Calculated),
	}
	return obj
}

// ProtoToAutoscalerAutoscalingPolicyCpuUtilization converts a AutoscalerAutoscalingPolicyCpuUtilization resource from its proto representation.
func ProtoToComputeAutoscalerAutoscalingPolicyCpuUtilization(p *computepb.ComputeAutoscalerAutoscalingPolicyCpuUtilization) *compute.AutoscalerAutoscalingPolicyCpuUtilization {
	if p == nil {
		return nil
	}
	obj := &compute.AutoscalerAutoscalingPolicyCpuUtilization{
		UtilizationTarget: dcl.Float64OrNil(p.UtilizationTarget),
	}
	return obj
}

// ProtoToAutoscalerAutoscalingPolicyCustomMetricUtilizations converts a AutoscalerAutoscalingPolicyCustomMetricUtilizations resource from its proto representation.
func ProtoToComputeAutoscalerAutoscalingPolicyCustomMetricUtilizations(p *computepb.ComputeAutoscalerAutoscalingPolicyCustomMetricUtilizations) *compute.AutoscalerAutoscalingPolicyCustomMetricUtilizations {
	if p == nil {
		return nil
	}
	obj := &compute.AutoscalerAutoscalingPolicyCustomMetricUtilizations{
		Metric:                dcl.StringOrNil(p.Metric),
		UtilizationTarget:     dcl.Float64OrNil(p.UtilizationTarget),
		UtilizationTargetType: ProtoToComputeAutoscalerAutoscalingPolicyCustomMetricUtilizationsUtilizationTargetTypeEnum(p.GetUtilizationTargetType()),
	}
	return obj
}

// ProtoToAutoscalerAutoscalingPolicyLoadBalancingUtilization converts a AutoscalerAutoscalingPolicyLoadBalancingUtilization resource from its proto representation.
func ProtoToComputeAutoscalerAutoscalingPolicyLoadBalancingUtilization(p *computepb.ComputeAutoscalerAutoscalingPolicyLoadBalancingUtilization) *compute.AutoscalerAutoscalingPolicyLoadBalancingUtilization {
	if p == nil {
		return nil
	}
	obj := &compute.AutoscalerAutoscalingPolicyLoadBalancingUtilization{
		UtilizationTarget: dcl.Float64OrNil(p.UtilizationTarget),
	}
	return obj
}

// ProtoToAutoscalerStatusDetails converts a AutoscalerStatusDetails resource from its proto representation.
func ProtoToComputeAutoscalerStatusDetails(p *computepb.ComputeAutoscalerStatusDetails) *compute.AutoscalerStatusDetails {
	if p == nil {
		return nil
	}
	obj := &compute.AutoscalerStatusDetails{
		Message: dcl.StringOrNil(p.Message),
		Type:    ProtoToComputeAutoscalerStatusDetailsTypeEnum(p.GetType()),
	}
	return obj
}

// ProtoToAutoscaler converts a Autoscaler resource from its proto representation.
func ProtoToAutoscaler(p *computepb.ComputeAutoscaler) *compute.Autoscaler {
	obj := &compute.Autoscaler{
		Id:                dcl.Int64OrNil(p.Id),
		Name:              dcl.StringOrNil(p.Name),
		Description:       dcl.StringOrNil(p.Description),
		Target:            dcl.StringOrNil(p.Target),
		AutoscalingPolicy: ProtoToComputeAutoscalerAutoscalingPolicy(p.GetAutoscalingPolicy()),
		Zone:              dcl.StringOrNil(p.Zone),
		Region:            dcl.StringOrNil(p.Region),
		SelfLink:          dcl.StringOrNil(p.SelfLink),
		Status:            ProtoToComputeAutoscalerStatusEnum(p.GetStatus()),
		RecommendedSize:   dcl.Int64OrNil(p.RecommendedSize),
		SelfLinkWithId:    dcl.StringOrNil(p.SelfLinkWithId),
		Project:           dcl.StringOrNil(p.Project),
		CreationTimestamp: dcl.StringOrNil(p.CreationTimestamp),
		Location:          dcl.StringOrNil(p.Location),
	}
	for _, r := range p.GetStatusDetails() {
		obj.StatusDetails = append(obj.StatusDetails, *ProtoToComputeAutoscalerStatusDetails(r))
	}
	return obj
}

// AutoscalerAutoscalingPolicyCustomMetricUtilizationsUtilizationTargetTypeEnumToProto converts a AutoscalerAutoscalingPolicyCustomMetricUtilizationsUtilizationTargetTypeEnum enum to its proto representation.
func ComputeAutoscalerAutoscalingPolicyCustomMetricUtilizationsUtilizationTargetTypeEnumToProto(e *compute.AutoscalerAutoscalingPolicyCustomMetricUtilizationsUtilizationTargetTypeEnum) computepb.ComputeAutoscalerAutoscalingPolicyCustomMetricUtilizationsUtilizationTargetTypeEnum {
	if e == nil {
		return computepb.ComputeAutoscalerAutoscalingPolicyCustomMetricUtilizationsUtilizationTargetTypeEnum(0)
	}
	if v, ok := computepb.ComputeAutoscalerAutoscalingPolicyCustomMetricUtilizationsUtilizationTargetTypeEnum_value["AutoscalerAutoscalingPolicyCustomMetricUtilizationsUtilizationTargetTypeEnum"+string(*e)]; ok {
		return computepb.ComputeAutoscalerAutoscalingPolicyCustomMetricUtilizationsUtilizationTargetTypeEnum(v)
	}
	return computepb.ComputeAutoscalerAutoscalingPolicyCustomMetricUtilizationsUtilizationTargetTypeEnum(0)
}

// AutoscalerAutoscalingPolicyModeEnumToProto converts a AutoscalerAutoscalingPolicyModeEnum enum to its proto representation.
func ComputeAutoscalerAutoscalingPolicyModeEnumToProto(e *compute.AutoscalerAutoscalingPolicyModeEnum) computepb.ComputeAutoscalerAutoscalingPolicyModeEnum {
	if e == nil {
		return computepb.ComputeAutoscalerAutoscalingPolicyModeEnum(0)
	}
	if v, ok := computepb.ComputeAutoscalerAutoscalingPolicyModeEnum_value["AutoscalerAutoscalingPolicyModeEnum"+string(*e)]; ok {
		return computepb.ComputeAutoscalerAutoscalingPolicyModeEnum(v)
	}
	return computepb.ComputeAutoscalerAutoscalingPolicyModeEnum(0)
}

// AutoscalerStatusEnumToProto converts a AutoscalerStatusEnum enum to its proto representation.
func ComputeAutoscalerStatusEnumToProto(e *compute.AutoscalerStatusEnum) computepb.ComputeAutoscalerStatusEnum {
	if e == nil {
		return computepb.ComputeAutoscalerStatusEnum(0)
	}
	if v, ok := computepb.ComputeAutoscalerStatusEnum_value["AutoscalerStatusEnum"+string(*e)]; ok {
		return computepb.ComputeAutoscalerStatusEnum(v)
	}
	return computepb.ComputeAutoscalerStatusEnum(0)
}

// AutoscalerStatusDetailsTypeEnumToProto converts a AutoscalerStatusDetailsTypeEnum enum to its proto representation.
func ComputeAutoscalerStatusDetailsTypeEnumToProto(e *compute.AutoscalerStatusDetailsTypeEnum) computepb.ComputeAutoscalerStatusDetailsTypeEnum {
	if e == nil {
		return computepb.ComputeAutoscalerStatusDetailsTypeEnum(0)
	}
	if v, ok := computepb.ComputeAutoscalerStatusDetailsTypeEnum_value["AutoscalerStatusDetailsTypeEnum"+string(*e)]; ok {
		return computepb.ComputeAutoscalerStatusDetailsTypeEnum(v)
	}
	return computepb.ComputeAutoscalerStatusDetailsTypeEnum(0)
}

// AutoscalerAutoscalingPolicyToProto converts a AutoscalerAutoscalingPolicy resource to its proto representation.
func ComputeAutoscalerAutoscalingPolicyToProto(o *compute.AutoscalerAutoscalingPolicy) *computepb.ComputeAutoscalerAutoscalingPolicy {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeAutoscalerAutoscalingPolicy{
		MinNumReplicas:           dcl.ValueOrEmptyInt64(o.MinNumReplicas),
		MaxNumReplicas:           dcl.ValueOrEmptyInt64(o.MaxNumReplicas),
		ScaleInControl:           ComputeAutoscalerAutoscalingPolicyScaleInControlToProto(o.ScaleInControl),
		CoolDownPeriodSec:        dcl.ValueOrEmptyInt64(o.CoolDownPeriodSec),
		CpuUtilization:           ComputeAutoscalerAutoscalingPolicyCpuUtilizationToProto(o.CpuUtilization),
		LoadBalancingUtilization: ComputeAutoscalerAutoscalingPolicyLoadBalancingUtilizationToProto(o.LoadBalancingUtilization),
		Mode:                     ComputeAutoscalerAutoscalingPolicyModeEnumToProto(o.Mode),
	}
	for _, r := range o.CustomMetricUtilizations {
		p.CustomMetricUtilizations = append(p.CustomMetricUtilizations, ComputeAutoscalerAutoscalingPolicyCustomMetricUtilizationsToProto(&r))
	}
	return p
}

// AutoscalerAutoscalingPolicyScaleInControlToProto converts a AutoscalerAutoscalingPolicyScaleInControl resource to its proto representation.
func ComputeAutoscalerAutoscalingPolicyScaleInControlToProto(o *compute.AutoscalerAutoscalingPolicyScaleInControl) *computepb.ComputeAutoscalerAutoscalingPolicyScaleInControl {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeAutoscalerAutoscalingPolicyScaleInControl{
		MaxScaledInReplicas: ComputeAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicasToProto(o.MaxScaledInReplicas),
		TimeWindowSec:       dcl.ValueOrEmptyInt64(o.TimeWindowSec),
	}
	return p
}

// AutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicasToProto converts a AutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas resource to its proto representation.
func ComputeAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicasToProto(o *compute.AutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas) *computepb.ComputeAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas{
		Fixed:      dcl.ValueOrEmptyInt64(o.Fixed),
		Percent:    dcl.ValueOrEmptyInt64(o.Percent),
		Calculated: dcl.ValueOrEmptyInt64(o.Calculated),
	}
	return p
}

// AutoscalerAutoscalingPolicyCpuUtilizationToProto converts a AutoscalerAutoscalingPolicyCpuUtilization resource to its proto representation.
func ComputeAutoscalerAutoscalingPolicyCpuUtilizationToProto(o *compute.AutoscalerAutoscalingPolicyCpuUtilization) *computepb.ComputeAutoscalerAutoscalingPolicyCpuUtilization {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeAutoscalerAutoscalingPolicyCpuUtilization{
		UtilizationTarget: dcl.ValueOrEmptyDouble(o.UtilizationTarget),
	}
	return p
}

// AutoscalerAutoscalingPolicyCustomMetricUtilizationsToProto converts a AutoscalerAutoscalingPolicyCustomMetricUtilizations resource to its proto representation.
func ComputeAutoscalerAutoscalingPolicyCustomMetricUtilizationsToProto(o *compute.AutoscalerAutoscalingPolicyCustomMetricUtilizations) *computepb.ComputeAutoscalerAutoscalingPolicyCustomMetricUtilizations {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeAutoscalerAutoscalingPolicyCustomMetricUtilizations{
		Metric:                dcl.ValueOrEmptyString(o.Metric),
		UtilizationTarget:     dcl.ValueOrEmptyDouble(o.UtilizationTarget),
		UtilizationTargetType: ComputeAutoscalerAutoscalingPolicyCustomMetricUtilizationsUtilizationTargetTypeEnumToProto(o.UtilizationTargetType),
	}
	return p
}

// AutoscalerAutoscalingPolicyLoadBalancingUtilizationToProto converts a AutoscalerAutoscalingPolicyLoadBalancingUtilization resource to its proto representation.
func ComputeAutoscalerAutoscalingPolicyLoadBalancingUtilizationToProto(o *compute.AutoscalerAutoscalingPolicyLoadBalancingUtilization) *computepb.ComputeAutoscalerAutoscalingPolicyLoadBalancingUtilization {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeAutoscalerAutoscalingPolicyLoadBalancingUtilization{
		UtilizationTarget: dcl.ValueOrEmptyDouble(o.UtilizationTarget),
	}
	return p
}

// AutoscalerStatusDetailsToProto converts a AutoscalerStatusDetails resource to its proto representation.
func ComputeAutoscalerStatusDetailsToProto(o *compute.AutoscalerStatusDetails) *computepb.ComputeAutoscalerStatusDetails {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeAutoscalerStatusDetails{
		Message: dcl.ValueOrEmptyString(o.Message),
		Type:    ComputeAutoscalerStatusDetailsTypeEnumToProto(o.Type),
	}
	return p
}

// AutoscalerToProto converts a Autoscaler resource to its proto representation.
func AutoscalerToProto(resource *compute.Autoscaler) *computepb.ComputeAutoscaler {
	p := &computepb.ComputeAutoscaler{
		Id:                dcl.ValueOrEmptyInt64(resource.Id),
		Name:              dcl.ValueOrEmptyString(resource.Name),
		Description:       dcl.ValueOrEmptyString(resource.Description),
		Target:            dcl.ValueOrEmptyString(resource.Target),
		AutoscalingPolicy: ComputeAutoscalerAutoscalingPolicyToProto(resource.AutoscalingPolicy),
		Zone:              dcl.ValueOrEmptyString(resource.Zone),
		Region:            dcl.ValueOrEmptyString(resource.Region),
		SelfLink:          dcl.ValueOrEmptyString(resource.SelfLink),
		Status:            ComputeAutoscalerStatusEnumToProto(resource.Status),
		RecommendedSize:   dcl.ValueOrEmptyInt64(resource.RecommendedSize),
		SelfLinkWithId:    dcl.ValueOrEmptyString(resource.SelfLinkWithId),
		Project:           dcl.ValueOrEmptyString(resource.Project),
		CreationTimestamp: dcl.ValueOrEmptyString(resource.CreationTimestamp),
		Location:          dcl.ValueOrEmptyString(resource.Location),
	}
	for _, r := range resource.StatusDetails {
		p.StatusDetails = append(p.StatusDetails, ComputeAutoscalerStatusDetailsToProto(&r))
	}

	return p
}

// ApplyAutoscaler handles the gRPC request by passing it to the underlying Autoscaler Apply() method.
func (s *AutoscalerServer) applyAutoscaler(ctx context.Context, c *compute.Client, request *computepb.ApplyComputeAutoscalerRequest) (*computepb.ComputeAutoscaler, error) {
	p := ProtoToAutoscaler(request.GetResource())
	res, err := c.ApplyAutoscaler(ctx, p)
	if err != nil {
		return nil, err
	}
	r := AutoscalerToProto(res)
	return r, nil
}

// ApplyAutoscaler handles the gRPC request by passing it to the underlying Autoscaler Apply() method.
func (s *AutoscalerServer) ApplyComputeAutoscaler(ctx context.Context, request *computepb.ApplyComputeAutoscalerRequest) (*computepb.ComputeAutoscaler, error) {
	cl, err := createConfigAutoscaler(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyAutoscaler(ctx, cl, request)
}

// DeleteAutoscaler handles the gRPC request by passing it to the underlying Autoscaler Delete() method.
func (s *AutoscalerServer) DeleteComputeAutoscaler(ctx context.Context, request *computepb.DeleteComputeAutoscalerRequest) (*emptypb.Empty, error) {

	cl, err := createConfigAutoscaler(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteAutoscaler(ctx, ProtoToAutoscaler(request.GetResource()))

}

// ListComputeAutoscaler handles the gRPC request by passing it to the underlying AutoscalerList() method.
func (s *AutoscalerServer) ListComputeAutoscaler(ctx context.Context, request *computepb.ListComputeAutoscalerRequest) (*computepb.ListComputeAutoscalerResponse, error) {
	cl, err := createConfigAutoscaler(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListAutoscaler(ctx, request.Project, request.Location)
	if err != nil {
		return nil, err
	}
	var protos []*computepb.ComputeAutoscaler
	for _, r := range resources.Items {
		rp := AutoscalerToProto(r)
		protos = append(protos, rp)
	}
	return &computepb.ListComputeAutoscalerResponse{Items: protos}, nil
}

func createConfigAutoscaler(ctx context.Context, service_account_file string) (*compute.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return compute.NewClient(conf), nil
}
