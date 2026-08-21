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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/cloudbuild/beta/cloudbuild_beta_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudbuild/beta"
)

// WorkerPoolServer implements the gRPC interface for WorkerPool.
type WorkerPoolServer struct{}

// ProtoToWorkerPoolStateEnum converts a WorkerPoolStateEnum enum from its proto representation.
func ProtoToCloudbuildBetaWorkerPoolStateEnum(e betapb.CloudbuildBetaWorkerPoolStateEnum) *beta.WorkerPoolStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.CloudbuildBetaWorkerPoolStateEnum_name[int32(e)]; ok {
		e := beta.WorkerPoolStateEnum(n[len("CloudbuildBetaWorkerPoolStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToWorkerPoolPrivatePoolV1ConfigNetworkConfigEgressOptionEnum converts a WorkerPoolPrivatePoolV1ConfigNetworkConfigEgressOptionEnum enum from its proto representation.
func ProtoToCloudbuildBetaWorkerPoolPrivatePoolV1ConfigNetworkConfigEgressOptionEnum(e betapb.CloudbuildBetaWorkerPoolPrivatePoolV1ConfigNetworkConfigEgressOptionEnum) *beta.WorkerPoolPrivatePoolV1ConfigNetworkConfigEgressOptionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.CloudbuildBetaWorkerPoolPrivatePoolV1ConfigNetworkConfigEgressOptionEnum_name[int32(e)]; ok {
		e := beta.WorkerPoolPrivatePoolV1ConfigNetworkConfigEgressOptionEnum(n[len("CloudbuildBetaWorkerPoolPrivatePoolV1ConfigNetworkConfigEgressOptionEnum"):])
		return &e
	}
	return nil
}

// ProtoToWorkerPoolPrivatePoolV1Config converts a WorkerPoolPrivatePoolV1Config object from its proto representation.
func ProtoToCloudbuildBetaWorkerPoolPrivatePoolV1Config(p *betapb.CloudbuildBetaWorkerPoolPrivatePoolV1Config) *beta.WorkerPoolPrivatePoolV1Config {
	if p == nil {
		return nil
	}
	obj := &beta.WorkerPoolPrivatePoolV1Config{
		WorkerConfig:  ProtoToCloudbuildBetaWorkerPoolPrivatePoolV1ConfigWorkerConfig(p.GetWorkerConfig()),
		NetworkConfig: ProtoToCloudbuildBetaWorkerPoolPrivatePoolV1ConfigNetworkConfig(p.GetNetworkConfig()),
	}
	return obj
}

// ProtoToWorkerPoolPrivatePoolV1ConfigWorkerConfig converts a WorkerPoolPrivatePoolV1ConfigWorkerConfig object from its proto representation.
func ProtoToCloudbuildBetaWorkerPoolPrivatePoolV1ConfigWorkerConfig(p *betapb.CloudbuildBetaWorkerPoolPrivatePoolV1ConfigWorkerConfig) *beta.WorkerPoolPrivatePoolV1ConfigWorkerConfig {
	if p == nil {
		return nil
	}
	obj := &beta.WorkerPoolPrivatePoolV1ConfigWorkerConfig{
		MachineType: dcl.StringOrNil(p.GetMachineType()),
		DiskSizeGb:  dcl.Int64OrNil(p.GetDiskSizeGb()),
	}
	return obj
}

// ProtoToWorkerPoolPrivatePoolV1ConfigNetworkConfig converts a WorkerPoolPrivatePoolV1ConfigNetworkConfig object from its proto representation.
func ProtoToCloudbuildBetaWorkerPoolPrivatePoolV1ConfigNetworkConfig(p *betapb.CloudbuildBetaWorkerPoolPrivatePoolV1ConfigNetworkConfig) *beta.WorkerPoolPrivatePoolV1ConfigNetworkConfig {
	if p == nil {
		return nil
	}
	obj := &beta.WorkerPoolPrivatePoolV1ConfigNetworkConfig{
		PeeredNetwork:        dcl.StringOrNil(p.GetPeeredNetwork()),
		PeeredNetworkIPRange: dcl.StringOrNil(p.GetPeeredNetworkIpRange()),
		EgressOption:         ProtoToCloudbuildBetaWorkerPoolPrivatePoolV1ConfigNetworkConfigEgressOptionEnum(p.GetEgressOption()),
	}
	return obj
}

// ProtoToWorkerPoolWorkerConfig converts a WorkerPoolWorkerConfig object from its proto representation.
func ProtoToCloudbuildBetaWorkerPoolWorkerConfig(p *betapb.CloudbuildBetaWorkerPoolWorkerConfig) *beta.WorkerPoolWorkerConfig {
	if p == nil {
		return nil
	}
	obj := &beta.WorkerPoolWorkerConfig{
		MachineType:  dcl.StringOrNil(p.GetMachineType()),
		DiskSizeGb:   dcl.Int64OrNil(p.GetDiskSizeGb()),
		NoExternalIP: dcl.Bool(p.GetNoExternalIp()),
	}
	return obj
}

// ProtoToWorkerPoolNetworkConfig converts a WorkerPoolNetworkConfig object from its proto representation.
func ProtoToCloudbuildBetaWorkerPoolNetworkConfig(p *betapb.CloudbuildBetaWorkerPoolNetworkConfig) *beta.WorkerPoolNetworkConfig {
	if p == nil {
		return nil
	}
	obj := &beta.WorkerPoolNetworkConfig{
		PeeredNetwork:        dcl.StringOrNil(p.GetPeeredNetwork()),
		PeeredNetworkIPRange: dcl.StringOrNil(p.GetPeeredNetworkIpRange()),
	}
	return obj
}

// ProtoToWorkerPool converts a WorkerPool resource from its proto representation.
func ProtoToWorkerPool(p *betapb.CloudbuildBetaWorkerPool) *beta.WorkerPool {
	obj := &beta.WorkerPool{
		Name:                dcl.StringOrNil(p.GetName()),
		DisplayName:         dcl.StringOrNil(p.GetDisplayName()),
		Uid:                 dcl.StringOrNil(p.GetUid()),
		CreateTime:          dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:          dcl.StringOrNil(p.GetUpdateTime()),
		DeleteTime:          dcl.StringOrNil(p.GetDeleteTime()),
		State:               ProtoToCloudbuildBetaWorkerPoolStateEnum(p.GetState()),
		PrivatePoolV1Config: ProtoToCloudbuildBetaWorkerPoolPrivatePoolV1Config(p.GetPrivatePoolV1Config()),
		Etag:                dcl.StringOrNil(p.GetEtag()),
		WorkerConfig:        ProtoToCloudbuildBetaWorkerPoolWorkerConfig(p.GetWorkerConfig()),
		NetworkConfig:       ProtoToCloudbuildBetaWorkerPoolNetworkConfig(p.GetNetworkConfig()),
		Project:             dcl.StringOrNil(p.GetProject()),
		Location:            dcl.StringOrNil(p.GetLocation()),
	}
	return obj
}

// WorkerPoolStateEnumToProto converts a WorkerPoolStateEnum enum to its proto representation.
func CloudbuildBetaWorkerPoolStateEnumToProto(e *beta.WorkerPoolStateEnum) betapb.CloudbuildBetaWorkerPoolStateEnum {
	if e == nil {
		return betapb.CloudbuildBetaWorkerPoolStateEnum(0)
	}
	if v, ok := betapb.CloudbuildBetaWorkerPoolStateEnum_value["WorkerPoolStateEnum"+string(*e)]; ok {
		return betapb.CloudbuildBetaWorkerPoolStateEnum(v)
	}
	return betapb.CloudbuildBetaWorkerPoolStateEnum(0)
}

// WorkerPoolPrivatePoolV1ConfigNetworkConfigEgressOptionEnumToProto converts a WorkerPoolPrivatePoolV1ConfigNetworkConfigEgressOptionEnum enum to its proto representation.
func CloudbuildBetaWorkerPoolPrivatePoolV1ConfigNetworkConfigEgressOptionEnumToProto(e *beta.WorkerPoolPrivatePoolV1ConfigNetworkConfigEgressOptionEnum) betapb.CloudbuildBetaWorkerPoolPrivatePoolV1ConfigNetworkConfigEgressOptionEnum {
	if e == nil {
		return betapb.CloudbuildBetaWorkerPoolPrivatePoolV1ConfigNetworkConfigEgressOptionEnum(0)
	}
	if v, ok := betapb.CloudbuildBetaWorkerPoolPrivatePoolV1ConfigNetworkConfigEgressOptionEnum_value["WorkerPoolPrivatePoolV1ConfigNetworkConfigEgressOptionEnum"+string(*e)]; ok {
		return betapb.CloudbuildBetaWorkerPoolPrivatePoolV1ConfigNetworkConfigEgressOptionEnum(v)
	}
	return betapb.CloudbuildBetaWorkerPoolPrivatePoolV1ConfigNetworkConfigEgressOptionEnum(0)
}

// WorkerPoolPrivatePoolV1ConfigToProto converts a WorkerPoolPrivatePoolV1Config object to its proto representation.
func CloudbuildBetaWorkerPoolPrivatePoolV1ConfigToProto(o *beta.WorkerPoolPrivatePoolV1Config) *betapb.CloudbuildBetaWorkerPoolPrivatePoolV1Config {
	if o == nil {
		return nil
	}
	p := &betapb.CloudbuildBetaWorkerPoolPrivatePoolV1Config{}
	p.SetWorkerConfig(CloudbuildBetaWorkerPoolPrivatePoolV1ConfigWorkerConfigToProto(o.WorkerConfig))
	p.SetNetworkConfig(CloudbuildBetaWorkerPoolPrivatePoolV1ConfigNetworkConfigToProto(o.NetworkConfig))
	return p
}

// WorkerPoolPrivatePoolV1ConfigWorkerConfigToProto converts a WorkerPoolPrivatePoolV1ConfigWorkerConfig object to its proto representation.
func CloudbuildBetaWorkerPoolPrivatePoolV1ConfigWorkerConfigToProto(o *beta.WorkerPoolPrivatePoolV1ConfigWorkerConfig) *betapb.CloudbuildBetaWorkerPoolPrivatePoolV1ConfigWorkerConfig {
	if o == nil {
		return nil
	}
	p := &betapb.CloudbuildBetaWorkerPoolPrivatePoolV1ConfigWorkerConfig{}
	p.SetMachineType(dcl.ValueOrEmptyString(o.MachineType))
	p.SetDiskSizeGb(dcl.ValueOrEmptyInt64(o.DiskSizeGb))
	return p
}

// WorkerPoolPrivatePoolV1ConfigNetworkConfigToProto converts a WorkerPoolPrivatePoolV1ConfigNetworkConfig object to its proto representation.
func CloudbuildBetaWorkerPoolPrivatePoolV1ConfigNetworkConfigToProto(o *beta.WorkerPoolPrivatePoolV1ConfigNetworkConfig) *betapb.CloudbuildBetaWorkerPoolPrivatePoolV1ConfigNetworkConfig {
	if o == nil {
		return nil
	}
	p := &betapb.CloudbuildBetaWorkerPoolPrivatePoolV1ConfigNetworkConfig{}
	p.SetPeeredNetwork(dcl.ValueOrEmptyString(o.PeeredNetwork))
	p.SetPeeredNetworkIpRange(dcl.ValueOrEmptyString(o.PeeredNetworkIPRange))
	p.SetEgressOption(CloudbuildBetaWorkerPoolPrivatePoolV1ConfigNetworkConfigEgressOptionEnumToProto(o.EgressOption))
	return p
}

// WorkerPoolWorkerConfigToProto converts a WorkerPoolWorkerConfig object to its proto representation.
func CloudbuildBetaWorkerPoolWorkerConfigToProto(o *beta.WorkerPoolWorkerConfig) *betapb.CloudbuildBetaWorkerPoolWorkerConfig {
	if o == nil {
		return nil
	}
	p := &betapb.CloudbuildBetaWorkerPoolWorkerConfig{}
	p.SetMachineType(dcl.ValueOrEmptyString(o.MachineType))
	p.SetDiskSizeGb(dcl.ValueOrEmptyInt64(o.DiskSizeGb))
	p.SetNoExternalIp(dcl.ValueOrEmptyBool(o.NoExternalIP))
	return p
}

// WorkerPoolNetworkConfigToProto converts a WorkerPoolNetworkConfig object to its proto representation.
func CloudbuildBetaWorkerPoolNetworkConfigToProto(o *beta.WorkerPoolNetworkConfig) *betapb.CloudbuildBetaWorkerPoolNetworkConfig {
	if o == nil {
		return nil
	}
	p := &betapb.CloudbuildBetaWorkerPoolNetworkConfig{}
	p.SetPeeredNetwork(dcl.ValueOrEmptyString(o.PeeredNetwork))
	p.SetPeeredNetworkIpRange(dcl.ValueOrEmptyString(o.PeeredNetworkIPRange))
	return p
}

// WorkerPoolToProto converts a WorkerPool resource to its proto representation.
func WorkerPoolToProto(resource *beta.WorkerPool) *betapb.CloudbuildBetaWorkerPool {
	p := &betapb.CloudbuildBetaWorkerPool{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDisplayName(dcl.ValueOrEmptyString(resource.DisplayName))
	p.SetUid(dcl.ValueOrEmptyString(resource.Uid))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetDeleteTime(dcl.ValueOrEmptyString(resource.DeleteTime))
	p.SetState(CloudbuildBetaWorkerPoolStateEnumToProto(resource.State))
	p.SetPrivatePoolV1Config(CloudbuildBetaWorkerPoolPrivatePoolV1ConfigToProto(resource.PrivatePoolV1Config))
	p.SetEtag(dcl.ValueOrEmptyString(resource.Etag))
	p.SetWorkerConfig(CloudbuildBetaWorkerPoolWorkerConfigToProto(resource.WorkerConfig))
	p.SetNetworkConfig(CloudbuildBetaWorkerPoolNetworkConfigToProto(resource.NetworkConfig))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	mAnnotations := make(map[string]string, len(resource.Annotations))
	for k, r := range resource.Annotations {
		mAnnotations[k] = r
	}
	p.SetAnnotations(mAnnotations)

	return p
}

// applyWorkerPool handles the gRPC request by passing it to the underlying WorkerPool Apply() method.
func (s *WorkerPoolServer) applyWorkerPool(ctx context.Context, c *beta.Client, request *betapb.ApplyCloudbuildBetaWorkerPoolRequest) (*betapb.CloudbuildBetaWorkerPool, error) {
	p := ProtoToWorkerPool(request.GetResource())
	res, err := c.ApplyWorkerPool(ctx, p)
	if err != nil {
		return nil, err
	}
	r := WorkerPoolToProto(res)
	return r, nil
}

// applyCloudbuildBetaWorkerPool handles the gRPC request by passing it to the underlying WorkerPool Apply() method.
func (s *WorkerPoolServer) ApplyCloudbuildBetaWorkerPool(ctx context.Context, request *betapb.ApplyCloudbuildBetaWorkerPoolRequest) (*betapb.CloudbuildBetaWorkerPool, error) {
	cl, err := createConfigWorkerPool(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyWorkerPool(ctx, cl, request)
}

// DeleteWorkerPool handles the gRPC request by passing it to the underlying WorkerPool Delete() method.
func (s *WorkerPoolServer) DeleteCloudbuildBetaWorkerPool(ctx context.Context, request *betapb.DeleteCloudbuildBetaWorkerPoolRequest) (*emptypb.Empty, error) {

	cl, err := createConfigWorkerPool(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteWorkerPool(ctx, ProtoToWorkerPool(request.GetResource()))

}

// ListCloudbuildBetaWorkerPool handles the gRPC request by passing it to the underlying WorkerPoolList() method.
func (s *WorkerPoolServer) ListCloudbuildBetaWorkerPool(ctx context.Context, request *betapb.ListCloudbuildBetaWorkerPoolRequest) (*betapb.ListCloudbuildBetaWorkerPoolResponse, error) {
	cl, err := createConfigWorkerPool(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListWorkerPool(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.CloudbuildBetaWorkerPool
	for _, r := range resources.Items {
		rp := WorkerPoolToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListCloudbuildBetaWorkerPoolResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigWorkerPool(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
