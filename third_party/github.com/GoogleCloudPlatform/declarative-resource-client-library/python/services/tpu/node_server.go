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
	tpupb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/tpu/tpu_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/tpu"
)

// Server implements the gRPC interface for Node.
type NodeServer struct{}

// ProtoToNodeStateEnum converts a NodeStateEnum enum from its proto representation.
func ProtoToTPUNodeStateEnum(e tpupb.TPUNodeStateEnum) *tpu.NodeStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := tpupb.TPUNodeStateEnum_name[int32(e)]; ok {
		e := tpu.NodeStateEnum(n[len("TPUNodeStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToNodeHealthEnum converts a NodeHealthEnum enum from its proto representation.
func ProtoToTPUNodeHealthEnum(e tpupb.TPUNodeHealthEnum) *tpu.NodeHealthEnum {
	if e == 0 {
		return nil
	}
	if n, ok := tpupb.TPUNodeHealthEnum_name[int32(e)]; ok {
		e := tpu.NodeHealthEnum(n[len("TPUNodeHealthEnum"):])
		return &e
	}
	return nil
}

// ProtoToNodeSymptomsSymptomTypeEnum converts a NodeSymptomsSymptomTypeEnum enum from its proto representation.
func ProtoToTPUNodeSymptomsSymptomTypeEnum(e tpupb.TPUNodeSymptomsSymptomTypeEnum) *tpu.NodeSymptomsSymptomTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := tpupb.TPUNodeSymptomsSymptomTypeEnum_name[int32(e)]; ok {
		e := tpu.NodeSymptomsSymptomTypeEnum(n[len("TPUNodeSymptomsSymptomTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToNodeCreateTime converts a NodeCreateTime resource from its proto representation.
func ProtoToTPUNodeCreateTime(p *tpupb.TPUNodeCreateTime) *tpu.NodeCreateTime {
	if p == nil {
		return nil
	}
	obj := &tpu.NodeCreateTime{
		Seconds: dcl.Int64OrNil(p.Seconds),
		Nanos:   dcl.Int64OrNil(p.Nanos),
	}
	return obj
}

// ProtoToNodeSchedulingConfig converts a NodeSchedulingConfig resource from its proto representation.
func ProtoToTPUNodeSchedulingConfig(p *tpupb.TPUNodeSchedulingConfig) *tpu.NodeSchedulingConfig {
	if p == nil {
		return nil
	}
	obj := &tpu.NodeSchedulingConfig{
		Preemptible: dcl.Bool(p.Preemptible),
		Reserved:    dcl.Bool(p.Reserved),
	}
	return obj
}

// ProtoToNodeNetworkEndpoints converts a NodeNetworkEndpoints resource from its proto representation.
func ProtoToTPUNodeNetworkEndpoints(p *tpupb.TPUNodeNetworkEndpoints) *tpu.NodeNetworkEndpoints {
	if p == nil {
		return nil
	}
	obj := &tpu.NodeNetworkEndpoints{
		IPAddress: dcl.StringOrNil(p.IpAddress),
		Port:      dcl.Int64OrNil(p.Port),
	}
	return obj
}

// ProtoToNodeSymptoms converts a NodeSymptoms resource from its proto representation.
func ProtoToTPUNodeSymptoms(p *tpupb.TPUNodeSymptoms) *tpu.NodeSymptoms {
	if p == nil {
		return nil
	}
	obj := &tpu.NodeSymptoms{
		CreateTime:  ProtoToTPUNodeSymptomsCreateTime(p.GetCreateTime()),
		SymptomType: ProtoToTPUNodeSymptomsSymptomTypeEnum(p.GetSymptomType()),
		Details:     dcl.StringOrNil(p.Details),
		WorkerId:    dcl.StringOrNil(p.WorkerId),
	}
	return obj
}

// ProtoToNodeSymptomsCreateTime converts a NodeSymptomsCreateTime resource from its proto representation.
func ProtoToTPUNodeSymptomsCreateTime(p *tpupb.TPUNodeSymptomsCreateTime) *tpu.NodeSymptomsCreateTime {
	if p == nil {
		return nil
	}
	obj := &tpu.NodeSymptomsCreateTime{
		Seconds: dcl.Int64OrNil(p.Seconds),
		Nanos:   dcl.Int64OrNil(p.Nanos),
	}
	return obj
}

// ProtoToNode converts a Node resource from its proto representation.
func ProtoToNode(p *tpupb.TPUNode) *tpu.Node {
	obj := &tpu.Node{
		Name:                 dcl.StringOrNil(p.Name),
		Description:          dcl.StringOrNil(p.Description),
		AcceleratorType:      dcl.StringOrNil(p.AcceleratorType),
		IPAddress:            dcl.StringOrNil(p.IpAddress),
		Port:                 dcl.StringOrNil(p.Port),
		State:                ProtoToTPUNodeStateEnum(p.GetState()),
		HealthDescription:    dcl.StringOrNil(p.HealthDescription),
		TensorflowVersion:    dcl.StringOrNil(p.TensorflowVersion),
		Network:              dcl.StringOrNil(p.Network),
		CidrBlock:            dcl.StringOrNil(p.CidrBlock),
		ServiceAccount:       dcl.StringOrNil(p.ServiceAccount),
		CreateTime:           ProtoToTPUNodeCreateTime(p.GetCreateTime()),
		SchedulingConfig:     ProtoToTPUNodeSchedulingConfig(p.GetSchedulingConfig()),
		Health:               ProtoToTPUNodeHealthEnum(p.GetHealth()),
		UseServiceNetworking: dcl.Bool(p.UseServiceNetworking),
		Project:              dcl.StringOrNil(p.Project),
		Location:             dcl.StringOrNil(p.Location),
	}
	for _, r := range p.GetNetworkEndpoints() {
		obj.NetworkEndpoints = append(obj.NetworkEndpoints, *ProtoToTPUNodeNetworkEndpoints(r))
	}
	for _, r := range p.GetSymptoms() {
		obj.Symptoms = append(obj.Symptoms, *ProtoToTPUNodeSymptoms(r))
	}
	return obj
}

// NodeStateEnumToProto converts a NodeStateEnum enum to its proto representation.
func TPUNodeStateEnumToProto(e *tpu.NodeStateEnum) tpupb.TPUNodeStateEnum {
	if e == nil {
		return tpupb.TPUNodeStateEnum(0)
	}
	if v, ok := tpupb.TPUNodeStateEnum_value["NodeStateEnum"+string(*e)]; ok {
		return tpupb.TPUNodeStateEnum(v)
	}
	return tpupb.TPUNodeStateEnum(0)
}

// NodeHealthEnumToProto converts a NodeHealthEnum enum to its proto representation.
func TPUNodeHealthEnumToProto(e *tpu.NodeHealthEnum) tpupb.TPUNodeHealthEnum {
	if e == nil {
		return tpupb.TPUNodeHealthEnum(0)
	}
	if v, ok := tpupb.TPUNodeHealthEnum_value["NodeHealthEnum"+string(*e)]; ok {
		return tpupb.TPUNodeHealthEnum(v)
	}
	return tpupb.TPUNodeHealthEnum(0)
}

// NodeSymptomsSymptomTypeEnumToProto converts a NodeSymptomsSymptomTypeEnum enum to its proto representation.
func TPUNodeSymptomsSymptomTypeEnumToProto(e *tpu.NodeSymptomsSymptomTypeEnum) tpupb.TPUNodeSymptomsSymptomTypeEnum {
	if e == nil {
		return tpupb.TPUNodeSymptomsSymptomTypeEnum(0)
	}
	if v, ok := tpupb.TPUNodeSymptomsSymptomTypeEnum_value["NodeSymptomsSymptomTypeEnum"+string(*e)]; ok {
		return tpupb.TPUNodeSymptomsSymptomTypeEnum(v)
	}
	return tpupb.TPUNodeSymptomsSymptomTypeEnum(0)
}

// NodeCreateTimeToProto converts a NodeCreateTime resource to its proto representation.
func TPUNodeCreateTimeToProto(o *tpu.NodeCreateTime) *tpupb.TPUNodeCreateTime {
	if o == nil {
		return nil
	}
	p := &tpupb.TPUNodeCreateTime{
		Seconds: dcl.ValueOrEmptyInt64(o.Seconds),
		Nanos:   dcl.ValueOrEmptyInt64(o.Nanos),
	}
	return p
}

// NodeSchedulingConfigToProto converts a NodeSchedulingConfig resource to its proto representation.
func TPUNodeSchedulingConfigToProto(o *tpu.NodeSchedulingConfig) *tpupb.TPUNodeSchedulingConfig {
	if o == nil {
		return nil
	}
	p := &tpupb.TPUNodeSchedulingConfig{
		Preemptible: dcl.ValueOrEmptyBool(o.Preemptible),
		Reserved:    dcl.ValueOrEmptyBool(o.Reserved),
	}
	return p
}

// NodeNetworkEndpointsToProto converts a NodeNetworkEndpoints resource to its proto representation.
func TPUNodeNetworkEndpointsToProto(o *tpu.NodeNetworkEndpoints) *tpupb.TPUNodeNetworkEndpoints {
	if o == nil {
		return nil
	}
	p := &tpupb.TPUNodeNetworkEndpoints{
		IpAddress: dcl.ValueOrEmptyString(o.IPAddress),
		Port:      dcl.ValueOrEmptyInt64(o.Port),
	}
	return p
}

// NodeSymptomsToProto converts a NodeSymptoms resource to its proto representation.
func TPUNodeSymptomsToProto(o *tpu.NodeSymptoms) *tpupb.TPUNodeSymptoms {
	if o == nil {
		return nil
	}
	p := &tpupb.TPUNodeSymptoms{
		CreateTime:  TPUNodeSymptomsCreateTimeToProto(o.CreateTime),
		SymptomType: TPUNodeSymptomsSymptomTypeEnumToProto(o.SymptomType),
		Details:     dcl.ValueOrEmptyString(o.Details),
		WorkerId:    dcl.ValueOrEmptyString(o.WorkerId),
	}
	return p
}

// NodeSymptomsCreateTimeToProto converts a NodeSymptomsCreateTime resource to its proto representation.
func TPUNodeSymptomsCreateTimeToProto(o *tpu.NodeSymptomsCreateTime) *tpupb.TPUNodeSymptomsCreateTime {
	if o == nil {
		return nil
	}
	p := &tpupb.TPUNodeSymptomsCreateTime{
		Seconds: dcl.ValueOrEmptyInt64(o.Seconds),
		Nanos:   dcl.ValueOrEmptyInt64(o.Nanos),
	}
	return p
}

// NodeToProto converts a Node resource to its proto representation.
func NodeToProto(resource *tpu.Node) *tpupb.TPUNode {
	p := &tpupb.TPUNode{
		Name:                 dcl.ValueOrEmptyString(resource.Name),
		Description:          dcl.ValueOrEmptyString(resource.Description),
		AcceleratorType:      dcl.ValueOrEmptyString(resource.AcceleratorType),
		IpAddress:            dcl.ValueOrEmptyString(resource.IPAddress),
		Port:                 dcl.ValueOrEmptyString(resource.Port),
		State:                TPUNodeStateEnumToProto(resource.State),
		HealthDescription:    dcl.ValueOrEmptyString(resource.HealthDescription),
		TensorflowVersion:    dcl.ValueOrEmptyString(resource.TensorflowVersion),
		Network:              dcl.ValueOrEmptyString(resource.Network),
		CidrBlock:            dcl.ValueOrEmptyString(resource.CidrBlock),
		ServiceAccount:       dcl.ValueOrEmptyString(resource.ServiceAccount),
		CreateTime:           TPUNodeCreateTimeToProto(resource.CreateTime),
		SchedulingConfig:     TPUNodeSchedulingConfigToProto(resource.SchedulingConfig),
		Health:               TPUNodeHealthEnumToProto(resource.Health),
		UseServiceNetworking: dcl.ValueOrEmptyBool(resource.UseServiceNetworking),
		Project:              dcl.ValueOrEmptyString(resource.Project),
		Location:             dcl.ValueOrEmptyString(resource.Location),
	}
	for _, r := range resource.NetworkEndpoints {
		p.NetworkEndpoints = append(p.NetworkEndpoints, TPUNodeNetworkEndpointsToProto(&r))
	}
	for _, r := range resource.Symptoms {
		p.Symptoms = append(p.Symptoms, TPUNodeSymptomsToProto(&r))
	}

	return p
}

// ApplyNode handles the gRPC request by passing it to the underlying Node Apply() method.
func (s *NodeServer) applyNode(ctx context.Context, c *tpu.Client, request *tpupb.ApplyTPUNodeRequest) (*tpupb.TPUNode, error) {
	p := ProtoToNode(request.GetResource())
	res, err := c.ApplyNode(ctx, p)
	if err != nil {
		return nil, err
	}
	r := NodeToProto(res)
	return r, nil
}

// ApplyNode handles the gRPC request by passing it to the underlying Node Apply() method.
func (s *NodeServer) ApplyTPUNode(ctx context.Context, request *tpupb.ApplyTPUNodeRequest) (*tpupb.TPUNode, error) {
	cl, err := createConfigNode(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyNode(ctx, cl, request)
}

// DeleteNode handles the gRPC request by passing it to the underlying Node Delete() method.
func (s *NodeServer) DeleteTPUNode(ctx context.Context, request *tpupb.DeleteTPUNodeRequest) (*emptypb.Empty, error) {

	cl, err := createConfigNode(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteNode(ctx, ProtoToNode(request.GetResource()))

}

// ListTPUNode handles the gRPC request by passing it to the underlying NodeList() method.
func (s *NodeServer) ListTPUNode(ctx context.Context, request *tpupb.ListTPUNodeRequest) (*tpupb.ListTPUNodeResponse, error) {
	cl, err := createConfigNode(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListNode(ctx, request.Project, request.Location)
	if err != nil {
		return nil, err
	}
	var protos []*tpupb.TPUNode
	for _, r := range resources.Items {
		rp := NodeToProto(r)
		protos = append(protos, rp)
	}
	return &tpupb.ListTPUNodeResponse{Items: protos}, nil
}

func createConfigNode(ctx context.Context, service_account_file string) (*tpu.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return tpu.NewClient(conf), nil
}
