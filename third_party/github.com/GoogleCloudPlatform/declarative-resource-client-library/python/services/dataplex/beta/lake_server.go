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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/dataplex/beta/dataplex_beta_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/dataplex/beta"
)

// LakeServer implements the gRPC interface for Lake.
type LakeServer struct{}

// ProtoToLakeStateEnum converts a LakeStateEnum enum from its proto representation.
func ProtoToDataplexBetaLakeStateEnum(e betapb.DataplexBetaLakeStateEnum) *beta.LakeStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.DataplexBetaLakeStateEnum_name[int32(e)]; ok {
		e := beta.LakeStateEnum(n[len("DataplexBetaLakeStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToLakeMetastoreStatusStateEnum converts a LakeMetastoreStatusStateEnum enum from its proto representation.
func ProtoToDataplexBetaLakeMetastoreStatusStateEnum(e betapb.DataplexBetaLakeMetastoreStatusStateEnum) *beta.LakeMetastoreStatusStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.DataplexBetaLakeMetastoreStatusStateEnum_name[int32(e)]; ok {
		e := beta.LakeMetastoreStatusStateEnum(n[len("DataplexBetaLakeMetastoreStatusStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToLakeMetastore converts a LakeMetastore object from its proto representation.
func ProtoToDataplexBetaLakeMetastore(p *betapb.DataplexBetaLakeMetastore) *beta.LakeMetastore {
	if p == nil {
		return nil
	}
	obj := &beta.LakeMetastore{
		Service: dcl.StringOrNil(p.GetService()),
	}
	return obj
}

// ProtoToLakeAssetStatus converts a LakeAssetStatus object from its proto representation.
func ProtoToDataplexBetaLakeAssetStatus(p *betapb.DataplexBetaLakeAssetStatus) *beta.LakeAssetStatus {
	if p == nil {
		return nil
	}
	obj := &beta.LakeAssetStatus{
		UpdateTime:                   dcl.StringOrNil(p.GetUpdateTime()),
		ActiveAssets:                 dcl.Int64OrNil(p.GetActiveAssets()),
		SecurityPolicyApplyingAssets: dcl.Int64OrNil(p.GetSecurityPolicyApplyingAssets()),
	}
	return obj
}

// ProtoToLakeMetastoreStatus converts a LakeMetastoreStatus object from its proto representation.
func ProtoToDataplexBetaLakeMetastoreStatus(p *betapb.DataplexBetaLakeMetastoreStatus) *beta.LakeMetastoreStatus {
	if p == nil {
		return nil
	}
	obj := &beta.LakeMetastoreStatus{
		State:      ProtoToDataplexBetaLakeMetastoreStatusStateEnum(p.GetState()),
		Message:    dcl.StringOrNil(p.GetMessage()),
		UpdateTime: dcl.StringOrNil(p.GetUpdateTime()),
		Endpoint:   dcl.StringOrNil(p.GetEndpoint()),
	}
	return obj
}

// ProtoToLake converts a Lake resource from its proto representation.
func ProtoToLake(p *betapb.DataplexBetaLake) *beta.Lake {
	obj := &beta.Lake{
		Name:            dcl.StringOrNil(p.GetName()),
		DisplayName:     dcl.StringOrNil(p.GetDisplayName()),
		Uid:             dcl.StringOrNil(p.GetUid()),
		CreateTime:      dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:      dcl.StringOrNil(p.GetUpdateTime()),
		Description:     dcl.StringOrNil(p.GetDescription()),
		State:           ProtoToDataplexBetaLakeStateEnum(p.GetState()),
		ServiceAccount:  dcl.StringOrNil(p.GetServiceAccount()),
		Metastore:       ProtoToDataplexBetaLakeMetastore(p.GetMetastore()),
		AssetStatus:     ProtoToDataplexBetaLakeAssetStatus(p.GetAssetStatus()),
		MetastoreStatus: ProtoToDataplexBetaLakeMetastoreStatus(p.GetMetastoreStatus()),
		Project:         dcl.StringOrNil(p.GetProject()),
		Location:        dcl.StringOrNil(p.GetLocation()),
	}
	return obj
}

// LakeStateEnumToProto converts a LakeStateEnum enum to its proto representation.
func DataplexBetaLakeStateEnumToProto(e *beta.LakeStateEnum) betapb.DataplexBetaLakeStateEnum {
	if e == nil {
		return betapb.DataplexBetaLakeStateEnum(0)
	}
	if v, ok := betapb.DataplexBetaLakeStateEnum_value["LakeStateEnum"+string(*e)]; ok {
		return betapb.DataplexBetaLakeStateEnum(v)
	}
	return betapb.DataplexBetaLakeStateEnum(0)
}

// LakeMetastoreStatusStateEnumToProto converts a LakeMetastoreStatusStateEnum enum to its proto representation.
func DataplexBetaLakeMetastoreStatusStateEnumToProto(e *beta.LakeMetastoreStatusStateEnum) betapb.DataplexBetaLakeMetastoreStatusStateEnum {
	if e == nil {
		return betapb.DataplexBetaLakeMetastoreStatusStateEnum(0)
	}
	if v, ok := betapb.DataplexBetaLakeMetastoreStatusStateEnum_value["LakeMetastoreStatusStateEnum"+string(*e)]; ok {
		return betapb.DataplexBetaLakeMetastoreStatusStateEnum(v)
	}
	return betapb.DataplexBetaLakeMetastoreStatusStateEnum(0)
}

// LakeMetastoreToProto converts a LakeMetastore object to its proto representation.
func DataplexBetaLakeMetastoreToProto(o *beta.LakeMetastore) *betapb.DataplexBetaLakeMetastore {
	if o == nil {
		return nil
	}
	p := &betapb.DataplexBetaLakeMetastore{}
	p.SetService(dcl.ValueOrEmptyString(o.Service))
	return p
}

// LakeAssetStatusToProto converts a LakeAssetStatus object to its proto representation.
func DataplexBetaLakeAssetStatusToProto(o *beta.LakeAssetStatus) *betapb.DataplexBetaLakeAssetStatus {
	if o == nil {
		return nil
	}
	p := &betapb.DataplexBetaLakeAssetStatus{}
	p.SetUpdateTime(dcl.ValueOrEmptyString(o.UpdateTime))
	p.SetActiveAssets(dcl.ValueOrEmptyInt64(o.ActiveAssets))
	p.SetSecurityPolicyApplyingAssets(dcl.ValueOrEmptyInt64(o.SecurityPolicyApplyingAssets))
	return p
}

// LakeMetastoreStatusToProto converts a LakeMetastoreStatus object to its proto representation.
func DataplexBetaLakeMetastoreStatusToProto(o *beta.LakeMetastoreStatus) *betapb.DataplexBetaLakeMetastoreStatus {
	if o == nil {
		return nil
	}
	p := &betapb.DataplexBetaLakeMetastoreStatus{}
	p.SetState(DataplexBetaLakeMetastoreStatusStateEnumToProto(o.State))
	p.SetMessage(dcl.ValueOrEmptyString(o.Message))
	p.SetUpdateTime(dcl.ValueOrEmptyString(o.UpdateTime))
	p.SetEndpoint(dcl.ValueOrEmptyString(o.Endpoint))
	return p
}

// LakeToProto converts a Lake resource to its proto representation.
func LakeToProto(resource *beta.Lake) *betapb.DataplexBetaLake {
	p := &betapb.DataplexBetaLake{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDisplayName(dcl.ValueOrEmptyString(resource.DisplayName))
	p.SetUid(dcl.ValueOrEmptyString(resource.Uid))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetState(DataplexBetaLakeStateEnumToProto(resource.State))
	p.SetServiceAccount(dcl.ValueOrEmptyString(resource.ServiceAccount))
	p.SetMetastore(DataplexBetaLakeMetastoreToProto(resource.Metastore))
	p.SetAssetStatus(DataplexBetaLakeAssetStatusToProto(resource.AssetStatus))
	p.SetMetastoreStatus(DataplexBetaLakeMetastoreStatusToProto(resource.MetastoreStatus))
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
func (s *LakeServer) applyLake(ctx context.Context, c *beta.Client, request *betapb.ApplyDataplexBetaLakeRequest) (*betapb.DataplexBetaLake, error) {
	p := ProtoToLake(request.GetResource())
	res, err := c.ApplyLake(ctx, p)
	if err != nil {
		return nil, err
	}
	r := LakeToProto(res)
	return r, nil
}

// applyDataplexBetaLake handles the gRPC request by passing it to the underlying Lake Apply() method.
func (s *LakeServer) ApplyDataplexBetaLake(ctx context.Context, request *betapb.ApplyDataplexBetaLakeRequest) (*betapb.DataplexBetaLake, error) {
	cl, err := createConfigLake(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyLake(ctx, cl, request)
}

// DeleteLake handles the gRPC request by passing it to the underlying Lake Delete() method.
func (s *LakeServer) DeleteDataplexBetaLake(ctx context.Context, request *betapb.DeleteDataplexBetaLakeRequest) (*emptypb.Empty, error) {

	cl, err := createConfigLake(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteLake(ctx, ProtoToLake(request.GetResource()))

}

// ListDataplexBetaLake handles the gRPC request by passing it to the underlying LakeList() method.
func (s *LakeServer) ListDataplexBetaLake(ctx context.Context, request *betapb.ListDataplexBetaLakeRequest) (*betapb.ListDataplexBetaLakeResponse, error) {
	cl, err := createConfigLake(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListLake(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.DataplexBetaLake
	for _, r := range resources.Items {
		rp := LakeToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListDataplexBetaLakeResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigLake(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
