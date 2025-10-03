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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/clouddeploy/beta/clouddeploy_beta_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/clouddeploy/beta"
)

// TargetServer implements the gRPC interface for Target.
type TargetServer struct{}

// ProtoToTargetExecutionConfigsUsagesEnum converts a TargetExecutionConfigsUsagesEnum enum from its proto representation.
func ProtoToClouddeployBetaTargetExecutionConfigsUsagesEnum(e betapb.ClouddeployBetaTargetExecutionConfigsUsagesEnum) *beta.TargetExecutionConfigsUsagesEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ClouddeployBetaTargetExecutionConfigsUsagesEnum_name[int32(e)]; ok {
		e := beta.TargetExecutionConfigsUsagesEnum(n[len("ClouddeployBetaTargetExecutionConfigsUsagesEnum"):])
		return &e
	}
	return nil
}

// ProtoToTargetGke converts a TargetGke object from its proto representation.
func ProtoToClouddeployBetaTargetGke(p *betapb.ClouddeployBetaTargetGke) *beta.TargetGke {
	if p == nil {
		return nil
	}
	obj := &beta.TargetGke{
		Cluster:    dcl.StringOrNil(p.GetCluster()),
		InternalIP: dcl.Bool(p.GetInternalIp()),
	}
	return obj
}

// ProtoToTargetAnthosCluster converts a TargetAnthosCluster object from its proto representation.
func ProtoToClouddeployBetaTargetAnthosCluster(p *betapb.ClouddeployBetaTargetAnthosCluster) *beta.TargetAnthosCluster {
	if p == nil {
		return nil
	}
	obj := &beta.TargetAnthosCluster{
		Membership: dcl.StringOrNil(p.GetMembership()),
	}
	return obj
}

// ProtoToTargetExecutionConfigs converts a TargetExecutionConfigs object from its proto representation.
func ProtoToClouddeployBetaTargetExecutionConfigs(p *betapb.ClouddeployBetaTargetExecutionConfigs) *beta.TargetExecutionConfigs {
	if p == nil {
		return nil
	}
	obj := &beta.TargetExecutionConfigs{
		WorkerPool:       dcl.StringOrNil(p.GetWorkerPool()),
		ServiceAccount:   dcl.StringOrNil(p.GetServiceAccount()),
		ArtifactStorage:  dcl.StringOrNil(p.GetArtifactStorage()),
		ExecutionTimeout: dcl.StringOrNil(p.GetExecutionTimeout()),
	}
	for _, r := range p.GetUsages() {
		obj.Usages = append(obj.Usages, *ProtoToClouddeployBetaTargetExecutionConfigsUsagesEnum(r))
	}
	return obj
}

// ProtoToTargetRun converts a TargetRun object from its proto representation.
func ProtoToClouddeployBetaTargetRun(p *betapb.ClouddeployBetaTargetRun) *beta.TargetRun {
	if p == nil {
		return nil
	}
	obj := &beta.TargetRun{
		Location: dcl.StringOrNil(p.GetLocation()),
	}
	return obj
}

// ProtoToTargetMultiTarget converts a TargetMultiTarget object from its proto representation.
func ProtoToClouddeployBetaTargetMultiTarget(p *betapb.ClouddeployBetaTargetMultiTarget) *beta.TargetMultiTarget {
	if p == nil {
		return nil
	}
	obj := &beta.TargetMultiTarget{}
	for _, r := range p.GetTargetIds() {
		obj.TargetIds = append(obj.TargetIds, r)
	}
	return obj
}

// ProtoToTarget converts a Target resource from its proto representation.
func ProtoToTarget(p *betapb.ClouddeployBetaTarget) *beta.Target {
	obj := &beta.Target{
		Name:            dcl.StringOrNil(p.GetName()),
		TargetId:        dcl.StringOrNil(p.GetTargetId()),
		Uid:             dcl.StringOrNil(p.GetUid()),
		Description:     dcl.StringOrNil(p.GetDescription()),
		RequireApproval: dcl.Bool(p.GetRequireApproval()),
		CreateTime:      dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:      dcl.StringOrNil(p.GetUpdateTime()),
		Gke:             ProtoToClouddeployBetaTargetGke(p.GetGke()),
		AnthosCluster:   ProtoToClouddeployBetaTargetAnthosCluster(p.GetAnthosCluster()),
		Etag:            dcl.StringOrNil(p.GetEtag()),
		Project:         dcl.StringOrNil(p.GetProject()),
		Location:        dcl.StringOrNil(p.GetLocation()),
		Run:             ProtoToClouddeployBetaTargetRun(p.GetRun()),
		MultiTarget:     ProtoToClouddeployBetaTargetMultiTarget(p.GetMultiTarget()),
	}
	for _, r := range p.GetExecutionConfigs() {
		obj.ExecutionConfigs = append(obj.ExecutionConfigs, *ProtoToClouddeployBetaTargetExecutionConfigs(r))
	}
	return obj
}

// TargetExecutionConfigsUsagesEnumToProto converts a TargetExecutionConfigsUsagesEnum enum to its proto representation.
func ClouddeployBetaTargetExecutionConfigsUsagesEnumToProto(e *beta.TargetExecutionConfigsUsagesEnum) betapb.ClouddeployBetaTargetExecutionConfigsUsagesEnum {
	if e == nil {
		return betapb.ClouddeployBetaTargetExecutionConfigsUsagesEnum(0)
	}
	if v, ok := betapb.ClouddeployBetaTargetExecutionConfigsUsagesEnum_value["TargetExecutionConfigsUsagesEnum"+string(*e)]; ok {
		return betapb.ClouddeployBetaTargetExecutionConfigsUsagesEnum(v)
	}
	return betapb.ClouddeployBetaTargetExecutionConfigsUsagesEnum(0)
}

// TargetGkeToProto converts a TargetGke object to its proto representation.
func ClouddeployBetaTargetGkeToProto(o *beta.TargetGke) *betapb.ClouddeployBetaTargetGke {
	if o == nil {
		return nil
	}
	p := &betapb.ClouddeployBetaTargetGke{}
	p.SetCluster(dcl.ValueOrEmptyString(o.Cluster))
	p.SetInternalIp(dcl.ValueOrEmptyBool(o.InternalIP))
	return p
}

// TargetAnthosClusterToProto converts a TargetAnthosCluster object to its proto representation.
func ClouddeployBetaTargetAnthosClusterToProto(o *beta.TargetAnthosCluster) *betapb.ClouddeployBetaTargetAnthosCluster {
	if o == nil {
		return nil
	}
	p := &betapb.ClouddeployBetaTargetAnthosCluster{}
	p.SetMembership(dcl.ValueOrEmptyString(o.Membership))
	return p
}

// TargetExecutionConfigsToProto converts a TargetExecutionConfigs object to its proto representation.
func ClouddeployBetaTargetExecutionConfigsToProto(o *beta.TargetExecutionConfigs) *betapb.ClouddeployBetaTargetExecutionConfigs {
	if o == nil {
		return nil
	}
	p := &betapb.ClouddeployBetaTargetExecutionConfigs{}
	p.SetWorkerPool(dcl.ValueOrEmptyString(o.WorkerPool))
	p.SetServiceAccount(dcl.ValueOrEmptyString(o.ServiceAccount))
	p.SetArtifactStorage(dcl.ValueOrEmptyString(o.ArtifactStorage))
	p.SetExecutionTimeout(dcl.ValueOrEmptyString(o.ExecutionTimeout))
	sUsages := make([]betapb.ClouddeployBetaTargetExecutionConfigsUsagesEnum, len(o.Usages))
	for i, r := range o.Usages {
		sUsages[i] = betapb.ClouddeployBetaTargetExecutionConfigsUsagesEnum(betapb.ClouddeployBetaTargetExecutionConfigsUsagesEnum_value[string(r)])
	}
	p.SetUsages(sUsages)
	return p
}

// TargetRunToProto converts a TargetRun object to its proto representation.
func ClouddeployBetaTargetRunToProto(o *beta.TargetRun) *betapb.ClouddeployBetaTargetRun {
	if o == nil {
		return nil
	}
	p := &betapb.ClouddeployBetaTargetRun{}
	p.SetLocation(dcl.ValueOrEmptyString(o.Location))
	return p
}

// TargetMultiTargetToProto converts a TargetMultiTarget object to its proto representation.
func ClouddeployBetaTargetMultiTargetToProto(o *beta.TargetMultiTarget) *betapb.ClouddeployBetaTargetMultiTarget {
	if o == nil {
		return nil
	}
	p := &betapb.ClouddeployBetaTargetMultiTarget{}
	sTargetIds := make([]string, len(o.TargetIds))
	for i, r := range o.TargetIds {
		sTargetIds[i] = r
	}
	p.SetTargetIds(sTargetIds)
	return p
}

// TargetToProto converts a Target resource to its proto representation.
func TargetToProto(resource *beta.Target) *betapb.ClouddeployBetaTarget {
	p := &betapb.ClouddeployBetaTarget{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetTargetId(dcl.ValueOrEmptyString(resource.TargetId))
	p.SetUid(dcl.ValueOrEmptyString(resource.Uid))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetRequireApproval(dcl.ValueOrEmptyBool(resource.RequireApproval))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetGke(ClouddeployBetaTargetGkeToProto(resource.Gke))
	p.SetAnthosCluster(ClouddeployBetaTargetAnthosClusterToProto(resource.AnthosCluster))
	p.SetEtag(dcl.ValueOrEmptyString(resource.Etag))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	p.SetRun(ClouddeployBetaTargetRunToProto(resource.Run))
	p.SetMultiTarget(ClouddeployBetaTargetMultiTargetToProto(resource.MultiTarget))
	mAnnotations := make(map[string]string, len(resource.Annotations))
	for k, r := range resource.Annotations {
		mAnnotations[k] = r
	}
	p.SetAnnotations(mAnnotations)
	mLabels := make(map[string]string, len(resource.Labels))
	for k, r := range resource.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)
	sExecutionConfigs := make([]*betapb.ClouddeployBetaTargetExecutionConfigs, len(resource.ExecutionConfigs))
	for i, r := range resource.ExecutionConfigs {
		sExecutionConfigs[i] = ClouddeployBetaTargetExecutionConfigsToProto(&r)
	}
	p.SetExecutionConfigs(sExecutionConfigs)
	mDeployParameters := make(map[string]string, len(resource.DeployParameters))
	for k, r := range resource.DeployParameters {
		mDeployParameters[k] = r
	}
	p.SetDeployParameters(mDeployParameters)

	return p
}

// applyTarget handles the gRPC request by passing it to the underlying Target Apply() method.
func (s *TargetServer) applyTarget(ctx context.Context, c *beta.Client, request *betapb.ApplyClouddeployBetaTargetRequest) (*betapb.ClouddeployBetaTarget, error) {
	p := ProtoToTarget(request.GetResource())
	res, err := c.ApplyTarget(ctx, p)
	if err != nil {
		return nil, err
	}
	r := TargetToProto(res)
	return r, nil
}

// applyClouddeployBetaTarget handles the gRPC request by passing it to the underlying Target Apply() method.
func (s *TargetServer) ApplyClouddeployBetaTarget(ctx context.Context, request *betapb.ApplyClouddeployBetaTargetRequest) (*betapb.ClouddeployBetaTarget, error) {
	cl, err := createConfigTarget(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyTarget(ctx, cl, request)
}

// DeleteTarget handles the gRPC request by passing it to the underlying Target Delete() method.
func (s *TargetServer) DeleteClouddeployBetaTarget(ctx context.Context, request *betapb.DeleteClouddeployBetaTargetRequest) (*emptypb.Empty, error) {

	cl, err := createConfigTarget(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteTarget(ctx, ProtoToTarget(request.GetResource()))

}

// ListClouddeployBetaTarget handles the gRPC request by passing it to the underlying TargetList() method.
func (s *TargetServer) ListClouddeployBetaTarget(ctx context.Context, request *betapb.ListClouddeployBetaTargetRequest) (*betapb.ListClouddeployBetaTargetResponse, error) {
	cl, err := createConfigTarget(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListTarget(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.ClouddeployBetaTarget
	for _, r := range resources.Items {
		rp := TargetToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListClouddeployBetaTargetResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigTarget(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
