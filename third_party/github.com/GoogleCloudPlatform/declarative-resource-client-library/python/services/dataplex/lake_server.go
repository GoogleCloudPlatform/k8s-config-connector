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
	dataplexpb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/dataplex/dataplex_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/dataplex"
)

// LakeServer implements the gRPC interface for Lake.
type LakeServer struct{}

// ProtoToLakeStateEnum converts a LakeStateEnum enum from its proto representation.
func ProtoToDataplexLakeStateEnum(e dataplexpb.DataplexLakeStateEnum) *dataplex.LakeStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := dataplexpb.DataplexLakeStateEnum_name[int32(e)]; ok {
		e := dataplex.LakeStateEnum(n[len("DataplexLakeStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToLakeMetastoreStatusStateEnum converts a LakeMetastoreStatusStateEnum enum from its proto representation.
func ProtoToDataplexLakeMetastoreStatusStateEnum(e dataplexpb.DataplexLakeMetastoreStatusStateEnum) *dataplex.LakeMetastoreStatusStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := dataplexpb.DataplexLakeMetastoreStatusStateEnum_name[int32(e)]; ok {
		e := dataplex.LakeMetastoreStatusStateEnum(n[len("DataplexLakeMetastoreStatusStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToLakeMetastore converts a LakeMetastore object from its proto representation.
func ProtoToDataplexLakeMetastore(p *dataplexpb.DataplexLakeMetastore) *dataplex.LakeMetastore {
	if p == nil {
		return nil
	}
	obj := &dataplex.LakeMetastore{
		Service: dcl.StringOrNil(p.GetService()),
	}
	return obj
}

// ProtoToLakeAssetStatus converts a LakeAssetStatus object from its proto representation.
func ProtoToDataplexLakeAssetStatus(p *dataplexpb.DataplexLakeAssetStatus) *dataplex.LakeAssetStatus {
	if p == nil {
		return nil
	}
	obj := &dataplex.LakeAssetStatus{
		UpdateTime:                   dcl.StringOrNil(p.GetUpdateTime()),
		ActiveAssets:                 dcl.Int64OrNil(p.GetActiveAssets()),
		SecurityPolicyApplyingAssets: dcl.Int64OrNil(p.GetSecurityPolicyApplyingAssets()),
	}
	return obj
}

// ProtoToLakeMetastoreStatus converts a LakeMetastoreStatus object from its proto representation.
func ProtoToDataplexLakeMetastoreStatus(p *dataplexpb.DataplexLakeMetastoreStatus) *dataplex.LakeMetastoreStatus {
	if p == nil {
		return nil
	}
	obj := &dataplex.LakeMetastoreStatus{
		State:      ProtoToDataplexLakeMetastoreStatusStateEnum(p.GetState()),
		Message:    dcl.StringOrNil(p.GetMessage()),
		UpdateTime: dcl.StringOrNil(p.GetUpdateTime()),
		Endpoint:   dcl.StringOrNil(p.GetEndpoint()),
	}
	return obj
}

// ProtoToLake converts a Lake resource from its proto representation.
func ProtoToLake(p *dataplexpb.DataplexLake) *dataplex.Lake {
	obj := &dataplex.Lake{
		Name:            dcl.StringOrNil(p.GetName()),
		DisplayName:     dcl.StringOrNil(p.GetDisplayName()),
		Uid:             dcl.StringOrNil(p.GetUid()),
		CreateTime:      dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:      dcl.StringOrNil(p.GetUpdateTime()),
		Description:     dcl.StringOrNil(p.GetDescription()),
		State:           ProtoToDataplexLakeStateEnum(p.GetState()),
		ServiceAccount:  dcl.StringOrNil(p.GetServiceAccount()),
		Metastore:       ProtoToDataplexLakeMetastore(p.GetMetastore()),
		AssetStatus:     ProtoToDataplexLakeAssetStatus(p.GetAssetStatus()),
		MetastoreStatus: ProtoToDataplexLakeMetastoreStatus(p.GetMetastoreStatus()),
		Project:         dcl.StringOrNil(p.GetProject()),
		Location:        dcl.StringOrNil(p.GetLocation()),
	}
	return obj
}

// LakeStateEnumToProto converts a LakeStateEnum enum to its proto representation.
func DataplexLakeStateEnumToProto(e *dataplex.LakeStateEnum) dataplexpb.DataplexLakeStateEnum {
	if e == nil {
		return dataplexpb.DataplexLakeStateEnum(0)
	}
	if v, ok := dataplexpb.DataplexLakeStateEnum_value["LakeStateEnum"+string(*e)]; ok {
		return dataplexpb.DataplexLakeStateEnum(v)
	}
	return dataplexpb.DataplexLakeStateEnum(0)
}

// LakeMetastoreStatusStateEnumToProto converts a LakeMetastoreStatusStateEnum enum to its proto representation.
func DataplexLakeMetastoreStatusStateEnumToProto(e *dataplex.LakeMetastoreStatusStateEnum) dataplexpb.DataplexLakeMetastoreStatusStateEnum {
	if e == nil {
		return dataplexpb.DataplexLakeMetastoreStatusStateEnum(0)
	}
	if v, ok := dataplexpb.DataplexLakeMetastoreStatusStateEnum_value["LakeMetastoreStatusStateEnum"+string(*e)]; ok {
		return dataplexpb.DataplexLakeMetastoreStatusStateEnum(v)
	}
	return dataplexpb.DataplexLakeMetastoreStatusStateEnum(0)
}

// LakeMetastoreToProto converts a LakeMetastore object to its proto representation.
func DataplexLakeMetastoreToProto(o *dataplex.LakeMetastore) *dataplexpb.DataplexLakeMetastore {
	if o == nil {
		return nil
	}
	p := &dataplexpb.DataplexLakeMetastore{}
	p.SetService(dcl.ValueOrEmptyString(o.Service))
	return p
}

// LakeAssetStatusToProto converts a LakeAssetStatus object to its proto representation.
func DataplexLakeAssetStatusToProto(o *dataplex.LakeAssetStatus) *dataplexpb.DataplexLakeAssetStatus {
	if o == nil {
		return nil
	}
	p := &dataplexpb.DataplexLakeAssetStatus{}
	p.SetUpdateTime(dcl.ValueOrEmptyString(o.UpdateTime))
	p.SetActiveAssets(dcl.ValueOrEmptyInt64(o.ActiveAssets))
	p.SetSecurityPolicyApplyingAssets(dcl.ValueOrEmptyInt64(o.SecurityPolicyApplyingAssets))
	return p
}

// LakeMetastoreStatusToProto converts a LakeMetastoreStatus object to its proto representation.
func DataplexLakeMetastoreStatusToProto(o *dataplex.LakeMetastoreStatus) *dataplexpb.DataplexLakeMetastoreStatus {
	if o == nil {
		return nil
	}
	p := &dataplexpb.DataplexLakeMetastoreStatus{}
	p.SetState(DataplexLakeMetastoreStatusStateEnumToProto(o.State))
	p.SetMessage(dcl.ValueOrEmptyString(o.Message))
	p.SetUpdateTime(dcl.ValueOrEmptyString(o.UpdateTime))
	p.SetEndpoint(dcl.ValueOrEmptyString(o.Endpoint))
	return p
}

// LakeToProto converts a Lake resource to its proto representation.
func LakeToProto(resource *dataplex.Lake) *dataplexpb.DataplexLake {
	p := &dataplexpb.DataplexLake{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDisplayName(dcl.ValueOrEmptyString(resource.DisplayName))
	p.SetUid(dcl.ValueOrEmptyString(resource.Uid))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetState(DataplexLakeStateEnumToProto(resource.State))
	p.SetServiceAccount(dcl.ValueOrEmptyString(resource.ServiceAccount))
	p.SetMetastore(DataplexLakeMetastoreToProto(resource.Metastore))
	p.SetAssetStatus(DataplexLakeAssetStatusToProto(resource.AssetStatus))
	p.SetMetastoreStatus(DataplexLakeMetastoreStatusToProto(resource.MetastoreStatus))
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
func (s *LakeServer) applyLake(ctx context.Context, c *dataplex.Client, request *dataplexpb.ApplyDataplexLakeRequest) (*dataplexpb.DataplexLake, error) {
	p := ProtoToLake(request.GetResource())
	res, err := c.ApplyLake(ctx, p)
	if err != nil {
		return nil, err
	}
	r := LakeToProto(res)
	return r, nil
}

// applyDataplexLake handles the gRPC request by passing it to the underlying Lake Apply() method.
func (s *LakeServer) ApplyDataplexLake(ctx context.Context, request *dataplexpb.ApplyDataplexLakeRequest) (*dataplexpb.DataplexLake, error) {
	cl, err := createConfigLake(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyLake(ctx, cl, request)
}

// DeleteLake handles the gRPC request by passing it to the underlying Lake Delete() method.
func (s *LakeServer) DeleteDataplexLake(ctx context.Context, request *dataplexpb.DeleteDataplexLakeRequest) (*emptypb.Empty, error) {

	cl, err := createConfigLake(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteLake(ctx, ProtoToLake(request.GetResource()))

}

// ListDataplexLake handles the gRPC request by passing it to the underlying LakeList() method.
func (s *LakeServer) ListDataplexLake(ctx context.Context, request *dataplexpb.ListDataplexLakeRequest) (*dataplexpb.ListDataplexLakeResponse, error) {
	cl, err := createConfigLake(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListLake(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*dataplexpb.DataplexLake
	for _, r := range resources.Items {
		rp := LakeToProto(r)
		protos = append(protos, rp)
	}
	p := &dataplexpb.ListDataplexLakeResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigLake(ctx context.Context, service_account_file string) (*dataplex.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return dataplex.NewClient(conf), nil
}
