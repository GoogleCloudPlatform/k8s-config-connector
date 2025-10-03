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
	dataprocpb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/dataproc/dataproc_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/dataproc"
)

// AutoscalingPolicyServer implements the gRPC interface for AutoscalingPolicy.
type AutoscalingPolicyServer struct{}

// ProtoToAutoscalingPolicyBasicAlgorithm converts a AutoscalingPolicyBasicAlgorithm object from its proto representation.
func ProtoToDataprocAutoscalingPolicyBasicAlgorithm(p *dataprocpb.DataprocAutoscalingPolicyBasicAlgorithm) *dataproc.AutoscalingPolicyBasicAlgorithm {
	if p == nil {
		return nil
	}
	obj := &dataproc.AutoscalingPolicyBasicAlgorithm{
		YarnConfig:     ProtoToDataprocAutoscalingPolicyBasicAlgorithmYarnConfig(p.GetYarnConfig()),
		CooldownPeriod: dcl.StringOrNil(p.GetCooldownPeriod()),
	}
	return obj
}

// ProtoToAutoscalingPolicyBasicAlgorithmYarnConfig converts a AutoscalingPolicyBasicAlgorithmYarnConfig object from its proto representation.
func ProtoToDataprocAutoscalingPolicyBasicAlgorithmYarnConfig(p *dataprocpb.DataprocAutoscalingPolicyBasicAlgorithmYarnConfig) *dataproc.AutoscalingPolicyBasicAlgorithmYarnConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.AutoscalingPolicyBasicAlgorithmYarnConfig{
		GracefulDecommissionTimeout: dcl.StringOrNil(p.GetGracefulDecommissionTimeout()),
		ScaleUpFactor:               dcl.Float64OrNil(p.GetScaleUpFactor()),
		ScaleDownFactor:             dcl.Float64OrNil(p.GetScaleDownFactor()),
		ScaleUpMinWorkerFraction:    dcl.Float64OrNil(p.GetScaleUpMinWorkerFraction()),
		ScaleDownMinWorkerFraction:  dcl.Float64OrNil(p.GetScaleDownMinWorkerFraction()),
	}
	return obj
}

// ProtoToAutoscalingPolicyWorkerConfig converts a AutoscalingPolicyWorkerConfig object from its proto representation.
func ProtoToDataprocAutoscalingPolicyWorkerConfig(p *dataprocpb.DataprocAutoscalingPolicyWorkerConfig) *dataproc.AutoscalingPolicyWorkerConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.AutoscalingPolicyWorkerConfig{
		MinInstances: dcl.Int64OrNil(p.GetMinInstances()),
		MaxInstances: dcl.Int64OrNil(p.GetMaxInstances()),
		Weight:       dcl.Int64OrNil(p.GetWeight()),
	}
	return obj
}

// ProtoToAutoscalingPolicySecondaryWorkerConfig converts a AutoscalingPolicySecondaryWorkerConfig object from its proto representation.
func ProtoToDataprocAutoscalingPolicySecondaryWorkerConfig(p *dataprocpb.DataprocAutoscalingPolicySecondaryWorkerConfig) *dataproc.AutoscalingPolicySecondaryWorkerConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.AutoscalingPolicySecondaryWorkerConfig{
		MinInstances: dcl.Int64OrNil(p.GetMinInstances()),
		MaxInstances: dcl.Int64OrNil(p.GetMaxInstances()),
		Weight:       dcl.Int64OrNil(p.GetWeight()),
	}
	return obj
}

// ProtoToAutoscalingPolicy converts a AutoscalingPolicy resource from its proto representation.
func ProtoToAutoscalingPolicy(p *dataprocpb.DataprocAutoscalingPolicy) *dataproc.AutoscalingPolicy {
	obj := &dataproc.AutoscalingPolicy{
		Name:                  dcl.StringOrNil(p.GetName()),
		BasicAlgorithm:        ProtoToDataprocAutoscalingPolicyBasicAlgorithm(p.GetBasicAlgorithm()),
		WorkerConfig:          ProtoToDataprocAutoscalingPolicyWorkerConfig(p.GetWorkerConfig()),
		SecondaryWorkerConfig: ProtoToDataprocAutoscalingPolicySecondaryWorkerConfig(p.GetSecondaryWorkerConfig()),
		Project:               dcl.StringOrNil(p.GetProject()),
		Location:              dcl.StringOrNil(p.GetLocation()),
	}
	return obj
}

// AutoscalingPolicyBasicAlgorithmToProto converts a AutoscalingPolicyBasicAlgorithm object to its proto representation.
func DataprocAutoscalingPolicyBasicAlgorithmToProto(o *dataproc.AutoscalingPolicyBasicAlgorithm) *dataprocpb.DataprocAutoscalingPolicyBasicAlgorithm {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocAutoscalingPolicyBasicAlgorithm{}
	p.SetYarnConfig(DataprocAutoscalingPolicyBasicAlgorithmYarnConfigToProto(o.YarnConfig))
	p.SetCooldownPeriod(dcl.ValueOrEmptyString(o.CooldownPeriod))
	return p
}

// AutoscalingPolicyBasicAlgorithmYarnConfigToProto converts a AutoscalingPolicyBasicAlgorithmYarnConfig object to its proto representation.
func DataprocAutoscalingPolicyBasicAlgorithmYarnConfigToProto(o *dataproc.AutoscalingPolicyBasicAlgorithmYarnConfig) *dataprocpb.DataprocAutoscalingPolicyBasicAlgorithmYarnConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocAutoscalingPolicyBasicAlgorithmYarnConfig{}
	p.SetGracefulDecommissionTimeout(dcl.ValueOrEmptyString(o.GracefulDecommissionTimeout))
	p.SetScaleUpFactor(dcl.ValueOrEmptyDouble(o.ScaleUpFactor))
	p.SetScaleDownFactor(dcl.ValueOrEmptyDouble(o.ScaleDownFactor))
	p.SetScaleUpMinWorkerFraction(dcl.ValueOrEmptyDouble(o.ScaleUpMinWorkerFraction))
	p.SetScaleDownMinWorkerFraction(dcl.ValueOrEmptyDouble(o.ScaleDownMinWorkerFraction))
	return p
}

// AutoscalingPolicyWorkerConfigToProto converts a AutoscalingPolicyWorkerConfig object to its proto representation.
func DataprocAutoscalingPolicyWorkerConfigToProto(o *dataproc.AutoscalingPolicyWorkerConfig) *dataprocpb.DataprocAutoscalingPolicyWorkerConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocAutoscalingPolicyWorkerConfig{}
	p.SetMinInstances(dcl.ValueOrEmptyInt64(o.MinInstances))
	p.SetMaxInstances(dcl.ValueOrEmptyInt64(o.MaxInstances))
	p.SetWeight(dcl.ValueOrEmptyInt64(o.Weight))
	return p
}

// AutoscalingPolicySecondaryWorkerConfigToProto converts a AutoscalingPolicySecondaryWorkerConfig object to its proto representation.
func DataprocAutoscalingPolicySecondaryWorkerConfigToProto(o *dataproc.AutoscalingPolicySecondaryWorkerConfig) *dataprocpb.DataprocAutoscalingPolicySecondaryWorkerConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocAutoscalingPolicySecondaryWorkerConfig{}
	p.SetMinInstances(dcl.ValueOrEmptyInt64(o.MinInstances))
	p.SetMaxInstances(dcl.ValueOrEmptyInt64(o.MaxInstances))
	p.SetWeight(dcl.ValueOrEmptyInt64(o.Weight))
	return p
}

// AutoscalingPolicyToProto converts a AutoscalingPolicy resource to its proto representation.
func AutoscalingPolicyToProto(resource *dataproc.AutoscalingPolicy) *dataprocpb.DataprocAutoscalingPolicy {
	p := &dataprocpb.DataprocAutoscalingPolicy{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetBasicAlgorithm(DataprocAutoscalingPolicyBasicAlgorithmToProto(resource.BasicAlgorithm))
	p.SetWorkerConfig(DataprocAutoscalingPolicyWorkerConfigToProto(resource.WorkerConfig))
	p.SetSecondaryWorkerConfig(DataprocAutoscalingPolicySecondaryWorkerConfigToProto(resource.SecondaryWorkerConfig))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))

	return p
}

// applyAutoscalingPolicy handles the gRPC request by passing it to the underlying AutoscalingPolicy Apply() method.
func (s *AutoscalingPolicyServer) applyAutoscalingPolicy(ctx context.Context, c *dataproc.Client, request *dataprocpb.ApplyDataprocAutoscalingPolicyRequest) (*dataprocpb.DataprocAutoscalingPolicy, error) {
	p := ProtoToAutoscalingPolicy(request.GetResource())
	res, err := c.ApplyAutoscalingPolicy(ctx, p)
	if err != nil {
		return nil, err
	}
	r := AutoscalingPolicyToProto(res)
	return r, nil
}

// applyDataprocAutoscalingPolicy handles the gRPC request by passing it to the underlying AutoscalingPolicy Apply() method.
func (s *AutoscalingPolicyServer) ApplyDataprocAutoscalingPolicy(ctx context.Context, request *dataprocpb.ApplyDataprocAutoscalingPolicyRequest) (*dataprocpb.DataprocAutoscalingPolicy, error) {
	cl, err := createConfigAutoscalingPolicy(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyAutoscalingPolicy(ctx, cl, request)
}

// DeleteAutoscalingPolicy handles the gRPC request by passing it to the underlying AutoscalingPolicy Delete() method.
func (s *AutoscalingPolicyServer) DeleteDataprocAutoscalingPolicy(ctx context.Context, request *dataprocpb.DeleteDataprocAutoscalingPolicyRequest) (*emptypb.Empty, error) {

	cl, err := createConfigAutoscalingPolicy(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteAutoscalingPolicy(ctx, ProtoToAutoscalingPolicy(request.GetResource()))

}

// ListDataprocAutoscalingPolicy handles the gRPC request by passing it to the underlying AutoscalingPolicyList() method.
func (s *AutoscalingPolicyServer) ListDataprocAutoscalingPolicy(ctx context.Context, request *dataprocpb.ListDataprocAutoscalingPolicyRequest) (*dataprocpb.ListDataprocAutoscalingPolicyResponse, error) {
	cl, err := createConfigAutoscalingPolicy(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListAutoscalingPolicy(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*dataprocpb.DataprocAutoscalingPolicy
	for _, r := range resources.Items {
		rp := AutoscalingPolicyToProto(r)
		protos = append(protos, rp)
	}
	p := &dataprocpb.ListDataprocAutoscalingPolicyResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigAutoscalingPolicy(ctx context.Context, service_account_file string) (*dataproc.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return dataproc.NewClient(conf), nil
}
