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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/dataproc/beta/dataproc_beta_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/dataproc/beta"
)

// AutoscalingPolicyServer implements the gRPC interface for AutoscalingPolicy.
type AutoscalingPolicyServer struct{}

// ProtoToAutoscalingPolicyBasicAlgorithm converts a AutoscalingPolicyBasicAlgorithm object from its proto representation.
func ProtoToDataprocBetaAutoscalingPolicyBasicAlgorithm(p *betapb.DataprocBetaAutoscalingPolicyBasicAlgorithm) *beta.AutoscalingPolicyBasicAlgorithm {
	if p == nil {
		return nil
	}
	obj := &beta.AutoscalingPolicyBasicAlgorithm{
		YarnConfig:     ProtoToDataprocBetaAutoscalingPolicyBasicAlgorithmYarnConfig(p.GetYarnConfig()),
		CooldownPeriod: dcl.StringOrNil(p.GetCooldownPeriod()),
	}
	return obj
}

// ProtoToAutoscalingPolicyBasicAlgorithmYarnConfig converts a AutoscalingPolicyBasicAlgorithmYarnConfig object from its proto representation.
func ProtoToDataprocBetaAutoscalingPolicyBasicAlgorithmYarnConfig(p *betapb.DataprocBetaAutoscalingPolicyBasicAlgorithmYarnConfig) *beta.AutoscalingPolicyBasicAlgorithmYarnConfig {
	if p == nil {
		return nil
	}
	obj := &beta.AutoscalingPolicyBasicAlgorithmYarnConfig{
		GracefulDecommissionTimeout: dcl.StringOrNil(p.GetGracefulDecommissionTimeout()),
		ScaleUpFactor:               dcl.Float64OrNil(p.GetScaleUpFactor()),
		ScaleDownFactor:             dcl.Float64OrNil(p.GetScaleDownFactor()),
		ScaleUpMinWorkerFraction:    dcl.Float64OrNil(p.GetScaleUpMinWorkerFraction()),
		ScaleDownMinWorkerFraction:  dcl.Float64OrNil(p.GetScaleDownMinWorkerFraction()),
	}
	return obj
}

// ProtoToAutoscalingPolicyWorkerConfig converts a AutoscalingPolicyWorkerConfig object from its proto representation.
func ProtoToDataprocBetaAutoscalingPolicyWorkerConfig(p *betapb.DataprocBetaAutoscalingPolicyWorkerConfig) *beta.AutoscalingPolicyWorkerConfig {
	if p == nil {
		return nil
	}
	obj := &beta.AutoscalingPolicyWorkerConfig{
		MinInstances: dcl.Int64OrNil(p.GetMinInstances()),
		MaxInstances: dcl.Int64OrNil(p.GetMaxInstances()),
		Weight:       dcl.Int64OrNil(p.GetWeight()),
	}
	return obj
}

// ProtoToAutoscalingPolicySecondaryWorkerConfig converts a AutoscalingPolicySecondaryWorkerConfig object from its proto representation.
func ProtoToDataprocBetaAutoscalingPolicySecondaryWorkerConfig(p *betapb.DataprocBetaAutoscalingPolicySecondaryWorkerConfig) *beta.AutoscalingPolicySecondaryWorkerConfig {
	if p == nil {
		return nil
	}
	obj := &beta.AutoscalingPolicySecondaryWorkerConfig{
		MinInstances: dcl.Int64OrNil(p.GetMinInstances()),
		MaxInstances: dcl.Int64OrNil(p.GetMaxInstances()),
		Weight:       dcl.Int64OrNil(p.GetWeight()),
	}
	return obj
}

// ProtoToAutoscalingPolicy converts a AutoscalingPolicy resource from its proto representation.
func ProtoToAutoscalingPolicy(p *betapb.DataprocBetaAutoscalingPolicy) *beta.AutoscalingPolicy {
	obj := &beta.AutoscalingPolicy{
		Name:                  dcl.StringOrNil(p.GetName()),
		BasicAlgorithm:        ProtoToDataprocBetaAutoscalingPolicyBasicAlgorithm(p.GetBasicAlgorithm()),
		WorkerConfig:          ProtoToDataprocBetaAutoscalingPolicyWorkerConfig(p.GetWorkerConfig()),
		SecondaryWorkerConfig: ProtoToDataprocBetaAutoscalingPolicySecondaryWorkerConfig(p.GetSecondaryWorkerConfig()),
		Project:               dcl.StringOrNil(p.GetProject()),
		Location:              dcl.StringOrNil(p.GetLocation()),
	}
	return obj
}

// AutoscalingPolicyBasicAlgorithmToProto converts a AutoscalingPolicyBasicAlgorithm object to its proto representation.
func DataprocBetaAutoscalingPolicyBasicAlgorithmToProto(o *beta.AutoscalingPolicyBasicAlgorithm) *betapb.DataprocBetaAutoscalingPolicyBasicAlgorithm {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaAutoscalingPolicyBasicAlgorithm{}
	p.SetYarnConfig(DataprocBetaAutoscalingPolicyBasicAlgorithmYarnConfigToProto(o.YarnConfig))
	p.SetCooldownPeriod(dcl.ValueOrEmptyString(o.CooldownPeriod))
	return p
}

// AutoscalingPolicyBasicAlgorithmYarnConfigToProto converts a AutoscalingPolicyBasicAlgorithmYarnConfig object to its proto representation.
func DataprocBetaAutoscalingPolicyBasicAlgorithmYarnConfigToProto(o *beta.AutoscalingPolicyBasicAlgorithmYarnConfig) *betapb.DataprocBetaAutoscalingPolicyBasicAlgorithmYarnConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaAutoscalingPolicyBasicAlgorithmYarnConfig{}
	p.SetGracefulDecommissionTimeout(dcl.ValueOrEmptyString(o.GracefulDecommissionTimeout))
	p.SetScaleUpFactor(dcl.ValueOrEmptyDouble(o.ScaleUpFactor))
	p.SetScaleDownFactor(dcl.ValueOrEmptyDouble(o.ScaleDownFactor))
	p.SetScaleUpMinWorkerFraction(dcl.ValueOrEmptyDouble(o.ScaleUpMinWorkerFraction))
	p.SetScaleDownMinWorkerFraction(dcl.ValueOrEmptyDouble(o.ScaleDownMinWorkerFraction))
	return p
}

// AutoscalingPolicyWorkerConfigToProto converts a AutoscalingPolicyWorkerConfig object to its proto representation.
func DataprocBetaAutoscalingPolicyWorkerConfigToProto(o *beta.AutoscalingPolicyWorkerConfig) *betapb.DataprocBetaAutoscalingPolicyWorkerConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaAutoscalingPolicyWorkerConfig{}
	p.SetMinInstances(dcl.ValueOrEmptyInt64(o.MinInstances))
	p.SetMaxInstances(dcl.ValueOrEmptyInt64(o.MaxInstances))
	p.SetWeight(dcl.ValueOrEmptyInt64(o.Weight))
	return p
}

// AutoscalingPolicySecondaryWorkerConfigToProto converts a AutoscalingPolicySecondaryWorkerConfig object to its proto representation.
func DataprocBetaAutoscalingPolicySecondaryWorkerConfigToProto(o *beta.AutoscalingPolicySecondaryWorkerConfig) *betapb.DataprocBetaAutoscalingPolicySecondaryWorkerConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaAutoscalingPolicySecondaryWorkerConfig{}
	p.SetMinInstances(dcl.ValueOrEmptyInt64(o.MinInstances))
	p.SetMaxInstances(dcl.ValueOrEmptyInt64(o.MaxInstances))
	p.SetWeight(dcl.ValueOrEmptyInt64(o.Weight))
	return p
}

// AutoscalingPolicyToProto converts a AutoscalingPolicy resource to its proto representation.
func AutoscalingPolicyToProto(resource *beta.AutoscalingPolicy) *betapb.DataprocBetaAutoscalingPolicy {
	p := &betapb.DataprocBetaAutoscalingPolicy{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetBasicAlgorithm(DataprocBetaAutoscalingPolicyBasicAlgorithmToProto(resource.BasicAlgorithm))
	p.SetWorkerConfig(DataprocBetaAutoscalingPolicyWorkerConfigToProto(resource.WorkerConfig))
	p.SetSecondaryWorkerConfig(DataprocBetaAutoscalingPolicySecondaryWorkerConfigToProto(resource.SecondaryWorkerConfig))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))

	return p
}

// applyAutoscalingPolicy handles the gRPC request by passing it to the underlying AutoscalingPolicy Apply() method.
func (s *AutoscalingPolicyServer) applyAutoscalingPolicy(ctx context.Context, c *beta.Client, request *betapb.ApplyDataprocBetaAutoscalingPolicyRequest) (*betapb.DataprocBetaAutoscalingPolicy, error) {
	p := ProtoToAutoscalingPolicy(request.GetResource())
	res, err := c.ApplyAutoscalingPolicy(ctx, p)
	if err != nil {
		return nil, err
	}
	r := AutoscalingPolicyToProto(res)
	return r, nil
}

// applyDataprocBetaAutoscalingPolicy handles the gRPC request by passing it to the underlying AutoscalingPolicy Apply() method.
func (s *AutoscalingPolicyServer) ApplyDataprocBetaAutoscalingPolicy(ctx context.Context, request *betapb.ApplyDataprocBetaAutoscalingPolicyRequest) (*betapb.DataprocBetaAutoscalingPolicy, error) {
	cl, err := createConfigAutoscalingPolicy(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyAutoscalingPolicy(ctx, cl, request)
}

// DeleteAutoscalingPolicy handles the gRPC request by passing it to the underlying AutoscalingPolicy Delete() method.
func (s *AutoscalingPolicyServer) DeleteDataprocBetaAutoscalingPolicy(ctx context.Context, request *betapb.DeleteDataprocBetaAutoscalingPolicyRequest) (*emptypb.Empty, error) {

	cl, err := createConfigAutoscalingPolicy(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteAutoscalingPolicy(ctx, ProtoToAutoscalingPolicy(request.GetResource()))

}

// ListDataprocBetaAutoscalingPolicy handles the gRPC request by passing it to the underlying AutoscalingPolicyList() method.
func (s *AutoscalingPolicyServer) ListDataprocBetaAutoscalingPolicy(ctx context.Context, request *betapb.ListDataprocBetaAutoscalingPolicyRequest) (*betapb.ListDataprocBetaAutoscalingPolicyResponse, error) {
	cl, err := createConfigAutoscalingPolicy(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListAutoscalingPolicy(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.DataprocBetaAutoscalingPolicy
	for _, r := range resources.Items {
		rp := AutoscalingPolicyToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListDataprocBetaAutoscalingPolicyResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigAutoscalingPolicy(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
