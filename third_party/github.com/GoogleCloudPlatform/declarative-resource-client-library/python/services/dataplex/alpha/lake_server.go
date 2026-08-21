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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/dataplex/alpha/dataplex_alpha_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/dataplex/alpha"
)

// LakeServer implements the gRPC interface for Lake.
type LakeServer struct{}

// ProtoToLakeStateEnum converts a LakeStateEnum enum from its proto representation.
func ProtoToDataplexAlphaLakeStateEnum(e alphapb.DataplexAlphaLakeStateEnum) *alpha.LakeStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.DataplexAlphaLakeStateEnum_name[int32(e)]; ok {
		e := alpha.LakeStateEnum(n[len("DataplexAlphaLakeStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToLakeMetastoreStatusStateEnum converts a LakeMetastoreStatusStateEnum enum from its proto representation.
func ProtoToDataplexAlphaLakeMetastoreStatusStateEnum(e alphapb.DataplexAlphaLakeMetastoreStatusStateEnum) *alpha.LakeMetastoreStatusStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.DataplexAlphaLakeMetastoreStatusStateEnum_name[int32(e)]; ok {
		e := alpha.LakeMetastoreStatusStateEnum(n[len("DataplexAlphaLakeMetastoreStatusStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToLakeMetastore converts a LakeMetastore object from its proto representation.
func ProtoToDataplexAlphaLakeMetastore(p *alphapb.DataplexAlphaLakeMetastore) *alpha.LakeMetastore {
	if p == nil {
		return nil
	}
	obj := &alpha.LakeMetastore{
		Service: dcl.StringOrNil(p.GetService()),
	}
	return obj
}

// ProtoToLakeAssetStatus converts a LakeAssetStatus object from its proto representation.
func ProtoToDataplexAlphaLakeAssetStatus(p *alphapb.DataplexAlphaLakeAssetStatus) *alpha.LakeAssetStatus {
	if p == nil {
		return nil
	}
	obj := &alpha.LakeAssetStatus{
		UpdateTime:                   dcl.StringOrNil(p.GetUpdateTime()),
		ActiveAssets:                 dcl.Int64OrNil(p.GetActiveAssets()),
		SecurityPolicyApplyingAssets: dcl.Int64OrNil(p.GetSecurityPolicyApplyingAssets()),
	}
	return obj
}

// ProtoToLakeMetastoreStatus converts a LakeMetastoreStatus object from its proto representation.
func ProtoToDataplexAlphaLakeMetastoreStatus(p *alphapb.DataplexAlphaLakeMetastoreStatus) *alpha.LakeMetastoreStatus {
	if p == nil {
		return nil
	}
	obj := &alpha.LakeMetastoreStatus{
		State:      ProtoToDataplexAlphaLakeMetastoreStatusStateEnum(p.GetState()),
		Message:    dcl.StringOrNil(p.GetMessage()),
		UpdateTime: dcl.StringOrNil(p.GetUpdateTime()),
		Endpoint:   dcl.StringOrNil(p.GetEndpoint()),
	}
	return obj
}

// ProtoToLake converts a Lake resource from its proto representation.
func ProtoToLake(p *alphapb.DataplexAlphaLake) *alpha.Lake {
	obj := &alpha.Lake{
		Name:            dcl.StringOrNil(p.GetName()),
		DisplayName:     dcl.StringOrNil(p.GetDisplayName()),
		Uid:             dcl.StringOrNil(p.GetUid()),
		CreateTime:      dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:      dcl.StringOrNil(p.GetUpdateTime()),
		Description:     dcl.StringOrNil(p.GetDescription()),
		State:           ProtoToDataplexAlphaLakeStateEnum(p.GetState()),
		ServiceAccount:  dcl.StringOrNil(p.GetServiceAccount()),
		Metastore:       ProtoToDataplexAlphaLakeMetastore(p.GetMetastore()),
		AssetStatus:     ProtoToDataplexAlphaLakeAssetStatus(p.GetAssetStatus()),
		MetastoreStatus: ProtoToDataplexAlphaLakeMetastoreStatus(p.GetMetastoreStatus()),
		Project:         dcl.StringOrNil(p.GetProject()),
		Location:        dcl.StringOrNil(p.GetLocation()),
	}
	return obj
}

// LakeStateEnumToProto converts a LakeStateEnum enum to its proto representation.
func DataplexAlphaLakeStateEnumToProto(e *alpha.LakeStateEnum) alphapb.DataplexAlphaLakeStateEnum {
	if e == nil {
		return alphapb.DataplexAlphaLakeStateEnum(0)
	}
	if v, ok := alphapb.DataplexAlphaLakeStateEnum_value["LakeStateEnum"+string(*e)]; ok {
		return alphapb.DataplexAlphaLakeStateEnum(v)
	}
	return alphapb.DataplexAlphaLakeStateEnum(0)
}

// LakeMetastoreStatusStateEnumToProto converts a LakeMetastoreStatusStateEnum enum to its proto representation.
func DataplexAlphaLakeMetastoreStatusStateEnumToProto(e *alpha.LakeMetastoreStatusStateEnum) alphapb.DataplexAlphaLakeMetastoreStatusStateEnum {
	if e == nil {
		return alphapb.DataplexAlphaLakeMetastoreStatusStateEnum(0)
	}
	if v, ok := alphapb.DataplexAlphaLakeMetastoreStatusStateEnum_value["LakeMetastoreStatusStateEnum"+string(*e)]; ok {
		return alphapb.DataplexAlphaLakeMetastoreStatusStateEnum(v)
	}
	return alphapb.DataplexAlphaLakeMetastoreStatusStateEnum(0)
}

// LakeMetastoreToProto converts a LakeMetastore object to its proto representation.
func DataplexAlphaLakeMetastoreToProto(o *alpha.LakeMetastore) *alphapb.DataplexAlphaLakeMetastore {
	if o == nil {
		return nil
	}
	p := &alphapb.DataplexAlphaLakeMetastore{}
	p.SetService(dcl.ValueOrEmptyString(o.Service))
	return p
}

// LakeAssetStatusToProto converts a LakeAssetStatus object to its proto representation.
func DataplexAlphaLakeAssetStatusToProto(o *alpha.LakeAssetStatus) *alphapb.DataplexAlphaLakeAssetStatus {
	if o == nil {
		return nil
	}
	p := &alphapb.DataplexAlphaLakeAssetStatus{}
	p.SetUpdateTime(dcl.ValueOrEmptyString(o.UpdateTime))
	p.SetActiveAssets(dcl.ValueOrEmptyInt64(o.ActiveAssets))
	p.SetSecurityPolicyApplyingAssets(dcl.ValueOrEmptyInt64(o.SecurityPolicyApplyingAssets))
	return p
}

// LakeMetastoreStatusToProto converts a LakeMetastoreStatus object to its proto representation.
func DataplexAlphaLakeMetastoreStatusToProto(o *alpha.LakeMetastoreStatus) *alphapb.DataplexAlphaLakeMetastoreStatus {
	if o == nil {
		return nil
	}
	p := &alphapb.DataplexAlphaLakeMetastoreStatus{}
	p.SetState(DataplexAlphaLakeMetastoreStatusStateEnumToProto(o.State))
	p.SetMessage(dcl.ValueOrEmptyString(o.Message))
	p.SetUpdateTime(dcl.ValueOrEmptyString(o.UpdateTime))
	p.SetEndpoint(dcl.ValueOrEmptyString(o.Endpoint))
	return p
}

// LakeToProto converts a Lake resource to its proto representation.
func LakeToProto(resource *alpha.Lake) *alphapb.DataplexAlphaLake {
	p := &alphapb.DataplexAlphaLake{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDisplayName(dcl.ValueOrEmptyString(resource.DisplayName))
	p.SetUid(dcl.ValueOrEmptyString(resource.Uid))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetState(DataplexAlphaLakeStateEnumToProto(resource.State))
	p.SetServiceAccount(dcl.ValueOrEmptyString(resource.ServiceAccount))
	p.SetMetastore(DataplexAlphaLakeMetastoreToProto(resource.Metastore))
	p.SetAssetStatus(DataplexAlphaLakeAssetStatusToProto(resource.AssetStatus))
	p.SetMetastoreStatus(DataplexAlphaLakeMetastoreStatusToProto(resource.MetastoreStatus))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	mLabels := make(map[string]string, len(resource.Labels))
	for k, r := range resource.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)

	return p
}

// applyLake handles the gRPC request by passing it to the underlying Lake Apply() method.
func (s *LakeServer) applyLake(ctx context.Context, c *alpha.Client, request *alphapb.ApplyDataplexAlphaLakeRequest) (*alphapb.DataplexAlphaLake, error) {
	p := ProtoToLake(request.GetResource())
	res, err := c.ApplyLake(ctx, p)
	if err != nil {
		return nil, err
	}
	r := LakeToProto(res)
	return r, nil
}

// applyDataplexAlphaLake handles the gRPC request by passing it to the underlying Lake Apply() method.
func (s *LakeServer) ApplyDataplexAlphaLake(ctx context.Context, request *alphapb.ApplyDataplexAlphaLakeRequest) (*alphapb.DataplexAlphaLake, error) {
	cl, err := createConfigLake(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyLake(ctx, cl, request)
}

// DeleteLake handles the gRPC request by passing it to the underlying Lake Delete() method.
func (s *LakeServer) DeleteDataplexAlphaLake(ctx context.Context, request *alphapb.DeleteDataplexAlphaLakeRequest) (*emptypb.Empty, error) {

	cl, err := createConfigLake(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteLake(ctx, ProtoToLake(request.GetResource()))

}

// ListDataplexAlphaLake handles the gRPC request by passing it to the underlying LakeList() method.
func (s *LakeServer) ListDataplexAlphaLake(ctx context.Context, request *alphapb.ListDataplexAlphaLakeRequest) (*alphapb.ListDataplexAlphaLakeResponse, error) {
	cl, err := createConfigLake(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListLake(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.DataplexAlphaLake
	for _, r := range resources.Items {
		rp := LakeToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListDataplexAlphaLakeResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigLake(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
