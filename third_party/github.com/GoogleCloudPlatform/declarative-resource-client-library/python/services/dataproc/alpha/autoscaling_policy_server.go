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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/dataproc/alpha/dataproc_alpha_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/dataproc/alpha"
)

// AutoscalingPolicyServer implements the gRPC interface for AutoscalingPolicy.
type AutoscalingPolicyServer struct{}

// ProtoToAutoscalingPolicyBasicAlgorithm converts a AutoscalingPolicyBasicAlgorithm object from its proto representation.
func ProtoToDataprocAlphaAutoscalingPolicyBasicAlgorithm(p *alphapb.DataprocAlphaAutoscalingPolicyBasicAlgorithm) *alpha.AutoscalingPolicyBasicAlgorithm {
	if p == nil {
		return nil
	}
	obj := &alpha.AutoscalingPolicyBasicAlgorithm{
		YarnConfig:     ProtoToDataprocAlphaAutoscalingPolicyBasicAlgorithmYarnConfig(p.GetYarnConfig()),
		CooldownPeriod: dcl.StringOrNil(p.GetCooldownPeriod()),
	}
	return obj
}

// ProtoToAutoscalingPolicyBasicAlgorithmYarnConfig converts a AutoscalingPolicyBasicAlgorithmYarnConfig object from its proto representation.
func ProtoToDataprocAlphaAutoscalingPolicyBasicAlgorithmYarnConfig(p *alphapb.DataprocAlphaAutoscalingPolicyBasicAlgorithmYarnConfig) *alpha.AutoscalingPolicyBasicAlgorithmYarnConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.AutoscalingPolicyBasicAlgorithmYarnConfig{
		GracefulDecommissionTimeout: dcl.StringOrNil(p.GetGracefulDecommissionTimeout()),
		ScaleUpFactor:               dcl.Float64OrNil(p.GetScaleUpFactor()),
		ScaleDownFactor:             dcl.Float64OrNil(p.GetScaleDownFactor()),
		ScaleUpMinWorkerFraction:    dcl.Float64OrNil(p.GetScaleUpMinWorkerFraction()),
		ScaleDownMinWorkerFraction:  dcl.Float64OrNil(p.GetScaleDownMinWorkerFraction()),
	}
	return obj
}

// ProtoToAutoscalingPolicyWorkerConfig converts a AutoscalingPolicyWorkerConfig object from its proto representation.
func ProtoToDataprocAlphaAutoscalingPolicyWorkerConfig(p *alphapb.DataprocAlphaAutoscalingPolicyWorkerConfig) *alpha.AutoscalingPolicyWorkerConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.AutoscalingPolicyWorkerConfig{
		MinInstances: dcl.Int64OrNil(p.GetMinInstances()),
		MaxInstances: dcl.Int64OrNil(p.GetMaxInstances()),
		Weight:       dcl.Int64OrNil(p.GetWeight()),
	}
	return obj
}

// ProtoToAutoscalingPolicySecondaryWorkerConfig converts a AutoscalingPolicySecondaryWorkerConfig object from its proto representation.
func ProtoToDataprocAlphaAutoscalingPolicySecondaryWorkerConfig(p *alphapb.DataprocAlphaAutoscalingPolicySecondaryWorkerConfig) *alpha.AutoscalingPolicySecondaryWorkerConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.AutoscalingPolicySecondaryWorkerConfig{
		MinInstances: dcl.Int64OrNil(p.GetMinInstances()),
		MaxInstances: dcl.Int64OrNil(p.GetMaxInstances()),
		Weight:       dcl.Int64OrNil(p.GetWeight()),
	}
	return obj
}

// ProtoToAutoscalingPolicy converts a AutoscalingPolicy resource from its proto representation.
func ProtoToAutoscalingPolicy(p *alphapb.DataprocAlphaAutoscalingPolicy) *alpha.AutoscalingPolicy {
	obj := &alpha.AutoscalingPolicy{
		Name:                  dcl.StringOrNil(p.GetName()),
		BasicAlgorithm:        ProtoToDataprocAlphaAutoscalingPolicyBasicAlgorithm(p.GetBasicAlgorithm()),
		WorkerConfig:          ProtoToDataprocAlphaAutoscalingPolicyWorkerConfig(p.GetWorkerConfig()),
		SecondaryWorkerConfig: ProtoToDataprocAlphaAutoscalingPolicySecondaryWorkerConfig(p.GetSecondaryWorkerConfig()),
		Project:               dcl.StringOrNil(p.GetProject()),
		Location:              dcl.StringOrNil(p.GetLocation()),
	}
	return obj
}

// AutoscalingPolicyBasicAlgorithmToProto converts a AutoscalingPolicyBasicAlgorithm object to its proto representation.
func DataprocAlphaAutoscalingPolicyBasicAlgorithmToProto(o *alpha.AutoscalingPolicyBasicAlgorithm) *alphapb.DataprocAlphaAutoscalingPolicyBasicAlgorithm {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaAutoscalingPolicyBasicAlgorithm{}
	p.SetYarnConfig(DataprocAlphaAutoscalingPolicyBasicAlgorithmYarnConfigToProto(o.YarnConfig))
	p.SetCooldownPeriod(dcl.ValueOrEmptyString(o.CooldownPeriod))
	return p
}

// AutoscalingPolicyBasicAlgorithmYarnConfigToProto converts a AutoscalingPolicyBasicAlgorithmYarnConfig object to its proto representation.
func DataprocAlphaAutoscalingPolicyBasicAlgorithmYarnConfigToProto(o *alpha.AutoscalingPolicyBasicAlgorithmYarnConfig) *alphapb.DataprocAlphaAutoscalingPolicyBasicAlgorithmYarnConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaAutoscalingPolicyBasicAlgorithmYarnConfig{}
	p.SetGracefulDecommissionTimeout(dcl.ValueOrEmptyString(o.GracefulDecommissionTimeout))
	p.SetScaleUpFactor(dcl.ValueOrEmptyDouble(o.ScaleUpFactor))
	p.SetScaleDownFactor(dcl.ValueOrEmptyDouble(o.ScaleDownFactor))
	p.SetScaleUpMinWorkerFraction(dcl.ValueOrEmptyDouble(o.ScaleUpMinWorkerFraction))
	p.SetScaleDownMinWorkerFraction(dcl.ValueOrEmptyDouble(o.ScaleDownMinWorkerFraction))
	return p
}

// AutoscalingPolicyWorkerConfigToProto converts a AutoscalingPolicyWorkerConfig object to its proto representation.
func DataprocAlphaAutoscalingPolicyWorkerConfigToProto(o *alpha.AutoscalingPolicyWorkerConfig) *alphapb.DataprocAlphaAutoscalingPolicyWorkerConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaAutoscalingPolicyWorkerConfig{}
	p.SetMinInstances(dcl.ValueOrEmptyInt64(o.MinInstances))
	p.SetMaxInstances(dcl.ValueOrEmptyInt64(o.MaxInstances))
	p.SetWeight(dcl.ValueOrEmptyInt64(o.Weight))
	return p
}

// AutoscalingPolicySecondaryWorkerConfigToProto converts a AutoscalingPolicySecondaryWorkerConfig object to its proto representation.
func DataprocAlphaAutoscalingPolicySecondaryWorkerConfigToProto(o *alpha.AutoscalingPolicySecondaryWorkerConfig) *alphapb.DataprocAlphaAutoscalingPolicySecondaryWorkerConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaAutoscalingPolicySecondaryWorkerConfig{}
	p.SetMinInstances(dcl.ValueOrEmptyInt64(o.MinInstances))
	p.SetMaxInstances(dcl.ValueOrEmptyInt64(o.MaxInstances))
	p.SetWeight(dcl.ValueOrEmptyInt64(o.Weight))
	return p
}

// AutoscalingPolicyToProto converts a AutoscalingPolicy resource to its proto representation.
func AutoscalingPolicyToProto(resource *alpha.AutoscalingPolicy) *alphapb.DataprocAlphaAutoscalingPolicy {
	p := &alphapb.DataprocAlphaAutoscalingPolicy{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetBasicAlgorithm(DataprocAlphaAutoscalingPolicyBasicAlgorithmToProto(resource.BasicAlgorithm))
	p.SetWorkerConfig(DataprocAlphaAutoscalingPolicyWorkerConfigToProto(resource.WorkerConfig))
	p.SetSecondaryWorkerConfig(DataprocAlphaAutoscalingPolicySecondaryWorkerConfigToProto(resource.SecondaryWorkerConfig))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))

	return p
}

// applyAutoscalingPolicy handles the gRPC request by passing it to the underlying AutoscalingPolicy Apply() method.
func (s *AutoscalingPolicyServer) applyAutoscalingPolicy(ctx context.Context, c *alpha.Client, request *alphapb.ApplyDataprocAlphaAutoscalingPolicyRequest) (*alphapb.DataprocAlphaAutoscalingPolicy, error) {
	p := ProtoToAutoscalingPolicy(request.GetResource())
	res, err := c.ApplyAutoscalingPolicy(ctx, p)
	if err != nil {
		return nil, err
	}
	r := AutoscalingPolicyToProto(res)
	return r, nil
}

// applyDataprocAlphaAutoscalingPolicy handles the gRPC request by passing it to the underlying AutoscalingPolicy Apply() method.
func (s *AutoscalingPolicyServer) ApplyDataprocAlphaAutoscalingPolicy(ctx context.Context, request *alphapb.ApplyDataprocAlphaAutoscalingPolicyRequest) (*alphapb.DataprocAlphaAutoscalingPolicy, error) {
	cl, err := createConfigAutoscalingPolicy(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyAutoscalingPolicy(ctx, cl, request)
}

// DeleteAutoscalingPolicy handles the gRPC request by passing it to the underlying AutoscalingPolicy Delete() method.
func (s *AutoscalingPolicyServer) DeleteDataprocAlphaAutoscalingPolicy(ctx context.Context, request *alphapb.DeleteDataprocAlphaAutoscalingPolicyRequest) (*emptypb.Empty, error) {

	cl, err := createConfigAutoscalingPolicy(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteAutoscalingPolicy(ctx, ProtoToAutoscalingPolicy(request.GetResource()))

}

// ListDataprocAlphaAutoscalingPolicy handles the gRPC request by passing it to the underlying AutoscalingPolicyList() method.
func (s *AutoscalingPolicyServer) ListDataprocAlphaAutoscalingPolicy(ctx context.Context, request *alphapb.ListDataprocAlphaAutoscalingPolicyRequest) (*alphapb.ListDataprocAlphaAutoscalingPolicyResponse, error) {
	cl, err := createConfigAutoscalingPolicy(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListAutoscalingPolicy(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.DataprocAlphaAutoscalingPolicy
	for _, r := range resources.Items {
		rp := AutoscalingPolicyToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListDataprocAlphaAutoscalingPolicyResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigAutoscalingPolicy(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
