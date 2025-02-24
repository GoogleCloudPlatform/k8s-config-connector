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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/cloudbuild/alpha/cloudbuild_alpha_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudbuild/alpha"
)

// WorkerPoolServer implements the gRPC interface for WorkerPool.
type WorkerPoolServer struct{}

// ProtoToWorkerPoolStateEnum converts a WorkerPoolStateEnum enum from its proto representation.
func ProtoToCloudbuildAlphaWorkerPoolStateEnum(e alphapb.CloudbuildAlphaWorkerPoolStateEnum) *alpha.WorkerPoolStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.CloudbuildAlphaWorkerPoolStateEnum_name[int32(e)]; ok {
		e := alpha.WorkerPoolStateEnum(n[len("CloudbuildAlphaWorkerPoolStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToWorkerPoolPrivatePoolV1ConfigNetworkConfigEgressOptionEnum converts a WorkerPoolPrivatePoolV1ConfigNetworkConfigEgressOptionEnum enum from its proto representation.
func ProtoToCloudbuildAlphaWorkerPoolPrivatePoolV1ConfigNetworkConfigEgressOptionEnum(e alphapb.CloudbuildAlphaWorkerPoolPrivatePoolV1ConfigNetworkConfigEgressOptionEnum) *alpha.WorkerPoolPrivatePoolV1ConfigNetworkConfigEgressOptionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.CloudbuildAlphaWorkerPoolPrivatePoolV1ConfigNetworkConfigEgressOptionEnum_name[int32(e)]; ok {
		e := alpha.WorkerPoolPrivatePoolV1ConfigNetworkConfigEgressOptionEnum(n[len("CloudbuildAlphaWorkerPoolPrivatePoolV1ConfigNetworkConfigEgressOptionEnum"):])
		return &e
	}
	return nil
}

// ProtoToWorkerPoolPrivatePoolV1Config converts a WorkerPoolPrivatePoolV1Config object from its proto representation.
func ProtoToCloudbuildAlphaWorkerPoolPrivatePoolV1Config(p *alphapb.CloudbuildAlphaWorkerPoolPrivatePoolV1Config) *alpha.WorkerPoolPrivatePoolV1Config {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkerPoolPrivatePoolV1Config{
		WorkerConfig:  ProtoToCloudbuildAlphaWorkerPoolPrivatePoolV1ConfigWorkerConfig(p.GetWorkerConfig()),
		NetworkConfig: ProtoToCloudbuildAlphaWorkerPoolPrivatePoolV1ConfigNetworkConfig(p.GetNetworkConfig()),
	}
	return obj
}

// ProtoToWorkerPoolPrivatePoolV1ConfigWorkerConfig converts a WorkerPoolPrivatePoolV1ConfigWorkerConfig object from its proto representation.
func ProtoToCloudbuildAlphaWorkerPoolPrivatePoolV1ConfigWorkerConfig(p *alphapb.CloudbuildAlphaWorkerPoolPrivatePoolV1ConfigWorkerConfig) *alpha.WorkerPoolPrivatePoolV1ConfigWorkerConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkerPoolPrivatePoolV1ConfigWorkerConfig{
		MachineType: dcl.StringOrNil(p.GetMachineType()),
		DiskSizeGb:  dcl.Int64OrNil(p.GetDiskSizeGb()),
	}
	return obj
}

// ProtoToWorkerPoolPrivatePoolV1ConfigNetworkConfig converts a WorkerPoolPrivatePoolV1ConfigNetworkConfig object from its proto representation.
func ProtoToCloudbuildAlphaWorkerPoolPrivatePoolV1ConfigNetworkConfig(p *alphapb.CloudbuildAlphaWorkerPoolPrivatePoolV1ConfigNetworkConfig) *alpha.WorkerPoolPrivatePoolV1ConfigNetworkConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkerPoolPrivatePoolV1ConfigNetworkConfig{
		PeeredNetwork:        dcl.StringOrNil(p.GetPeeredNetwork()),
		PeeredNetworkIPRange: dcl.StringOrNil(p.GetPeeredNetworkIpRange()),
		EgressOption:         ProtoToCloudbuildAlphaWorkerPoolPrivatePoolV1ConfigNetworkConfigEgressOptionEnum(p.GetEgressOption()),
	}
	return obj
}

// ProtoToWorkerPoolWorkerConfig converts a WorkerPoolWorkerConfig object from its proto representation.
func ProtoToCloudbuildAlphaWorkerPoolWorkerConfig(p *alphapb.CloudbuildAlphaWorkerPoolWorkerConfig) *alpha.WorkerPoolWorkerConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkerPoolWorkerConfig{
		MachineType:  dcl.StringOrNil(p.GetMachineType()),
		DiskSizeGb:   dcl.Int64OrNil(p.GetDiskSizeGb()),
		NoExternalIP: dcl.Bool(p.GetNoExternalIp()),
	}
	return obj
}

// ProtoToWorkerPoolNetworkConfig converts a WorkerPoolNetworkConfig object from its proto representation.
func ProtoToCloudbuildAlphaWorkerPoolNetworkConfig(p *alphapb.CloudbuildAlphaWorkerPoolNetworkConfig) *alpha.WorkerPoolNetworkConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkerPoolNetworkConfig{
		PeeredNetwork:        dcl.StringOrNil(p.GetPeeredNetwork()),
		PeeredNetworkIPRange: dcl.StringOrNil(p.GetPeeredNetworkIpRange()),
	}
	return obj
}

// ProtoToWorkerPool converts a WorkerPool resource from its proto representation.
func ProtoToWorkerPool(p *alphapb.CloudbuildAlphaWorkerPool) *alpha.WorkerPool {
	obj := &alpha.WorkerPool{
		Name:                dcl.StringOrNil(p.GetName()),
		DisplayName:         dcl.StringOrNil(p.GetDisplayName()),
		Uid:                 dcl.StringOrNil(p.GetUid()),
		CreateTime:          dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:          dcl.StringOrNil(p.GetUpdateTime()),
		DeleteTime:          dcl.StringOrNil(p.GetDeleteTime()),
		State:               ProtoToCloudbuildAlphaWorkerPoolStateEnum(p.GetState()),
		PrivatePoolV1Config: ProtoToCloudbuildAlphaWorkerPoolPrivatePoolV1Config(p.GetPrivatePoolV1Config()),
		Etag:                dcl.StringOrNil(p.GetEtag()),
		WorkerConfig:        ProtoToCloudbuildAlphaWorkerPoolWorkerConfig(p.GetWorkerConfig()),
		NetworkConfig:       ProtoToCloudbuildAlphaWorkerPoolNetworkConfig(p.GetNetworkConfig()),
		Project:             dcl.StringOrNil(p.GetProject()),
		Location:            dcl.StringOrNil(p.GetLocation()),
	}
	return obj
}

// WorkerPoolStateEnumToProto converts a WorkerPoolStateEnum enum to its proto representation.
func CloudbuildAlphaWorkerPoolStateEnumToProto(e *alpha.WorkerPoolStateEnum) alphapb.CloudbuildAlphaWorkerPoolStateEnum {
	if e == nil {
		return alphapb.CloudbuildAlphaWorkerPoolStateEnum(0)
	}
	if v, ok := alphapb.CloudbuildAlphaWorkerPoolStateEnum_value["WorkerPoolStateEnum"+string(*e)]; ok {
		return alphapb.CloudbuildAlphaWorkerPoolStateEnum(v)
	}
	return alphapb.CloudbuildAlphaWorkerPoolStateEnum(0)
}

// WorkerPoolPrivatePoolV1ConfigNetworkConfigEgressOptionEnumToProto converts a WorkerPoolPrivatePoolV1ConfigNetworkConfigEgressOptionEnum enum to its proto representation.
func CloudbuildAlphaWorkerPoolPrivatePoolV1ConfigNetworkConfigEgressOptionEnumToProto(e *alpha.WorkerPoolPrivatePoolV1ConfigNetworkConfigEgressOptionEnum) alphapb.CloudbuildAlphaWorkerPoolPrivatePoolV1ConfigNetworkConfigEgressOptionEnum {
	if e == nil {
		return alphapb.CloudbuildAlphaWorkerPoolPrivatePoolV1ConfigNetworkConfigEgressOptionEnum(0)
	}
	if v, ok := alphapb.CloudbuildAlphaWorkerPoolPrivatePoolV1ConfigNetworkConfigEgressOptionEnum_value["WorkerPoolPrivatePoolV1ConfigNetworkConfigEgressOptionEnum"+string(*e)]; ok {
		return alphapb.CloudbuildAlphaWorkerPoolPrivatePoolV1ConfigNetworkConfigEgressOptionEnum(v)
	}
	return alphapb.CloudbuildAlphaWorkerPoolPrivatePoolV1ConfigNetworkConfigEgressOptionEnum(0)
}

// WorkerPoolPrivatePoolV1ConfigToProto converts a WorkerPoolPrivatePoolV1Config object to its proto representation.
func CloudbuildAlphaWorkerPoolPrivatePoolV1ConfigToProto(o *alpha.WorkerPoolPrivatePoolV1Config) *alphapb.CloudbuildAlphaWorkerPoolPrivatePoolV1Config {
	if o == nil {
		return nil
	}
	p := &alphapb.CloudbuildAlphaWorkerPoolPrivatePoolV1Config{}
	p.SetWorkerConfig(CloudbuildAlphaWorkerPoolPrivatePoolV1ConfigWorkerConfigToProto(o.WorkerConfig))
	p.SetNetworkConfig(CloudbuildAlphaWorkerPoolPrivatePoolV1ConfigNetworkConfigToProto(o.NetworkConfig))
	return p
}

// WorkerPoolPrivatePoolV1ConfigWorkerConfigToProto converts a WorkerPoolPrivatePoolV1ConfigWorkerConfig object to its proto representation.
func CloudbuildAlphaWorkerPoolPrivatePoolV1ConfigWorkerConfigToProto(o *alpha.WorkerPoolPrivatePoolV1ConfigWorkerConfig) *alphapb.CloudbuildAlphaWorkerPoolPrivatePoolV1ConfigWorkerConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.CloudbuildAlphaWorkerPoolPrivatePoolV1ConfigWorkerConfig{}
	p.SetMachineType(dcl.ValueOrEmptyString(o.MachineType))
	p.SetDiskSizeGb(dcl.ValueOrEmptyInt64(o.DiskSizeGb))
	return p
}

// WorkerPoolPrivatePoolV1ConfigNetworkConfigToProto converts a WorkerPoolPrivatePoolV1ConfigNetworkConfig object to its proto representation.
func CloudbuildAlphaWorkerPoolPrivatePoolV1ConfigNetworkConfigToProto(o *alpha.WorkerPoolPrivatePoolV1ConfigNetworkConfig) *alphapb.CloudbuildAlphaWorkerPoolPrivatePoolV1ConfigNetworkConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.CloudbuildAlphaWorkerPoolPrivatePoolV1ConfigNetworkConfig{}
	p.SetPeeredNetwork(dcl.ValueOrEmptyString(o.PeeredNetwork))
	p.SetPeeredNetworkIpRange(dcl.ValueOrEmptyString(o.PeeredNetworkIPRange))
	p.SetEgressOption(CloudbuildAlphaWorkerPoolPrivatePoolV1ConfigNetworkConfigEgressOptionEnumToProto(o.EgressOption))
	return p
}

// WorkerPoolWorkerConfigToProto converts a WorkerPoolWorkerConfig object to its proto representation.
func CloudbuildAlphaWorkerPoolWorkerConfigToProto(o *alpha.WorkerPoolWorkerConfig) *alphapb.CloudbuildAlphaWorkerPoolWorkerConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.CloudbuildAlphaWorkerPoolWorkerConfig{}
	p.SetMachineType(dcl.ValueOrEmptyString(o.MachineType))
	p.SetDiskSizeGb(dcl.ValueOrEmptyInt64(o.DiskSizeGb))
	p.SetNoExternalIp(dcl.ValueOrEmptyBool(o.NoExternalIP))
	return p
}

// WorkerPoolNetworkConfigToProto converts a WorkerPoolNetworkConfig object to its proto representation.
func CloudbuildAlphaWorkerPoolNetworkConfigToProto(o *alpha.WorkerPoolNetworkConfig) *alphapb.CloudbuildAlphaWorkerPoolNetworkConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.CloudbuildAlphaWorkerPoolNetworkConfig{}
	p.SetPeeredNetwork(dcl.ValueOrEmptyString(o.PeeredNetwork))
	p.SetPeeredNetworkIpRange(dcl.ValueOrEmptyString(o.PeeredNetworkIPRange))
	return p
}

// WorkerPoolToProto converts a WorkerPool resource to its proto representation.
func WorkerPoolToProto(resource *alpha.WorkerPool) *alphapb.CloudbuildAlphaWorkerPool {
	p := &alphapb.CloudbuildAlphaWorkerPool{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDisplayName(dcl.ValueOrEmptyString(resource.DisplayName))
	p.SetUid(dcl.ValueOrEmptyString(resource.Uid))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetDeleteTime(dcl.ValueOrEmptyString(resource.DeleteTime))
	p.SetState(CloudbuildAlphaWorkerPoolStateEnumToProto(resource.State))
	p.SetPrivatePoolV1Config(CloudbuildAlphaWorkerPoolPrivatePoolV1ConfigToProto(resource.PrivatePoolV1Config))
	p.SetEtag(dcl.ValueOrEmptyString(resource.Etag))
	p.SetWorkerConfig(CloudbuildAlphaWorkerPoolWorkerConfigToProto(resource.WorkerConfig))
	p.SetNetworkConfig(CloudbuildAlphaWorkerPoolNetworkConfigToProto(resource.NetworkConfig))
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
func (s *WorkerPoolServer) applyWorkerPool(ctx context.Context, c *alpha.Client, request *alphapb.ApplyCloudbuildAlphaWorkerPoolRequest) (*alphapb.CloudbuildAlphaWorkerPool, error) {
	p := ProtoToWorkerPool(request.GetResource())
	res, err := c.ApplyWorkerPool(ctx, p)
	if err != nil {
		return nil, err
	}
	r := WorkerPoolToProto(res)
	return r, nil
}

// applyCloudbuildAlphaWorkerPool handles the gRPC request by passing it to the underlying WorkerPool Apply() method.
func (s *WorkerPoolServer) ApplyCloudbuildAlphaWorkerPool(ctx context.Context, request *alphapb.ApplyCloudbuildAlphaWorkerPoolRequest) (*alphapb.CloudbuildAlphaWorkerPool, error) {
	cl, err := createConfigWorkerPool(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyWorkerPool(ctx, cl, request)
}

// DeleteWorkerPool handles the gRPC request by passing it to the underlying WorkerPool Delete() method.
func (s *WorkerPoolServer) DeleteCloudbuildAlphaWorkerPool(ctx context.Context, request *alphapb.DeleteCloudbuildAlphaWorkerPoolRequest) (*emptypb.Empty, error) {

	cl, err := createConfigWorkerPool(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteWorkerPool(ctx, ProtoToWorkerPool(request.GetResource()))

}

// ListCloudbuildAlphaWorkerPool handles the gRPC request by passing it to the underlying WorkerPoolList() method.
func (s *WorkerPoolServer) ListCloudbuildAlphaWorkerPool(ctx context.Context, request *alphapb.ListCloudbuildAlphaWorkerPoolRequest) (*alphapb.ListCloudbuildAlphaWorkerPoolResponse, error) {
	cl, err := createConfigWorkerPool(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListWorkerPool(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.CloudbuildAlphaWorkerPool
	for _, r := range resources.Items {
		rp := WorkerPoolToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListCloudbuildAlphaWorkerPoolResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigWorkerPool(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
