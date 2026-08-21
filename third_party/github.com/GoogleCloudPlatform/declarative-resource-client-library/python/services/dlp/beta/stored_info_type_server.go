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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/dlp/beta/dlp_beta_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/dlp/beta"
)

// StoredInfoTypeServer implements the gRPC interface for StoredInfoType.
type StoredInfoTypeServer struct{}

// ProtoToStoredInfoTypeLargeCustomDictionary converts a StoredInfoTypeLargeCustomDictionary object from its proto representation.
func ProtoToDlpBetaStoredInfoTypeLargeCustomDictionary(p *betapb.DlpBetaStoredInfoTypeLargeCustomDictionary) *beta.StoredInfoTypeLargeCustomDictionary {
	if p == nil {
		return nil
	}
	obj := &beta.StoredInfoTypeLargeCustomDictionary{
		OutputPath:          ProtoToDlpBetaStoredInfoTypeLargeCustomDictionaryOutputPath(p.GetOutputPath()),
		CloudStorageFileSet: ProtoToDlpBetaStoredInfoTypeLargeCustomDictionaryCloudStorageFileSet(p.GetCloudStorageFileSet()),
		BigQueryField:       ProtoToDlpBetaStoredInfoTypeLargeCustomDictionaryBigQueryField(p.GetBigQueryField()),
	}
	return obj
}

// ProtoToStoredInfoTypeLargeCustomDictionaryOutputPath converts a StoredInfoTypeLargeCustomDictionaryOutputPath object from its proto representation.
func ProtoToDlpBetaStoredInfoTypeLargeCustomDictionaryOutputPath(p *betapb.DlpBetaStoredInfoTypeLargeCustomDictionaryOutputPath) *beta.StoredInfoTypeLargeCustomDictionaryOutputPath {
	if p == nil {
		return nil
	}
	obj := &beta.StoredInfoTypeLargeCustomDictionaryOutputPath{
		Path: dcl.StringOrNil(p.GetPath()),
	}
	return obj
}

// ProtoToStoredInfoTypeLargeCustomDictionaryCloudStorageFileSet converts a StoredInfoTypeLargeCustomDictionaryCloudStorageFileSet object from its proto representation.
func ProtoToDlpBetaStoredInfoTypeLargeCustomDictionaryCloudStorageFileSet(p *betapb.DlpBetaStoredInfoTypeLargeCustomDictionaryCloudStorageFileSet) *beta.StoredInfoTypeLargeCustomDictionaryCloudStorageFileSet {
	if p == nil {
		return nil
	}
	obj := &beta.StoredInfoTypeLargeCustomDictionaryCloudStorageFileSet{
		Url: dcl.StringOrNil(p.GetUrl()),
	}
	return obj
}

// ProtoToStoredInfoTypeLargeCustomDictionaryBigQueryField converts a StoredInfoTypeLargeCustomDictionaryBigQueryField object from its proto representation.
func ProtoToDlpBetaStoredInfoTypeLargeCustomDictionaryBigQueryField(p *betapb.DlpBetaStoredInfoTypeLargeCustomDictionaryBigQueryField) *beta.StoredInfoTypeLargeCustomDictionaryBigQueryField {
	if p == nil {
		return nil
	}
	obj := &beta.StoredInfoTypeLargeCustomDictionaryBigQueryField{
		Table: ProtoToDlpBetaStoredInfoTypeLargeCustomDictionaryBigQueryFieldTable(p.GetTable()),
		Field: ProtoToDlpBetaStoredInfoTypeLargeCustomDictionaryBigQueryFieldField(p.GetField()),
	}
	return obj
}

// ProtoToStoredInfoTypeLargeCustomDictionaryBigQueryFieldTable converts a StoredInfoTypeLargeCustomDictionaryBigQueryFieldTable object from its proto representation.
func ProtoToDlpBetaStoredInfoTypeLargeCustomDictionaryBigQueryFieldTable(p *betapb.DlpBetaStoredInfoTypeLargeCustomDictionaryBigQueryFieldTable) *beta.StoredInfoTypeLargeCustomDictionaryBigQueryFieldTable {
	if p == nil {
		return nil
	}
	obj := &beta.StoredInfoTypeLargeCustomDictionaryBigQueryFieldTable{
		ProjectId: dcl.StringOrNil(p.GetProjectId()),
		DatasetId: dcl.StringOrNil(p.GetDatasetId()),
		TableId:   dcl.StringOrNil(p.GetTableId()),
	}
	return obj
}

// ProtoToStoredInfoTypeLargeCustomDictionaryBigQueryFieldField converts a StoredInfoTypeLargeCustomDictionaryBigQueryFieldField object from its proto representation.
func ProtoToDlpBetaStoredInfoTypeLargeCustomDictionaryBigQueryFieldField(p *betapb.DlpBetaStoredInfoTypeLargeCustomDictionaryBigQueryFieldField) *beta.StoredInfoTypeLargeCustomDictionaryBigQueryFieldField {
	if p == nil {
		return nil
	}
	obj := &beta.StoredInfoTypeLargeCustomDictionaryBigQueryFieldField{
		Name: dcl.StringOrNil(p.GetName()),
	}
	return obj
}

// ProtoToStoredInfoTypeDictionary converts a StoredInfoTypeDictionary object from its proto representation.
func ProtoToDlpBetaStoredInfoTypeDictionary(p *betapb.DlpBetaStoredInfoTypeDictionary) *beta.StoredInfoTypeDictionary {
	if p == nil {
		return nil
	}
	obj := &beta.StoredInfoTypeDictionary{
		WordList:         ProtoToDlpBetaStoredInfoTypeDictionaryWordList(p.GetWordList()),
		CloudStoragePath: ProtoToDlpBetaStoredInfoTypeDictionaryCloudStoragePath(p.GetCloudStoragePath()),
	}
	return obj
}

// ProtoToStoredInfoTypeDictionaryWordList converts a StoredInfoTypeDictionaryWordList object from its proto representation.
func ProtoToDlpBetaStoredInfoTypeDictionaryWordList(p *betapb.DlpBetaStoredInfoTypeDictionaryWordList) *beta.StoredInfoTypeDictionaryWordList {
	if p == nil {
		return nil
	}
	obj := &beta.StoredInfoTypeDictionaryWordList{}
	for _, r := range p.GetWords() {
		obj.Words = append(obj.Words, r)
	}
	return obj
}

// ProtoToStoredInfoTypeDictionaryCloudStoragePath converts a StoredInfoTypeDictionaryCloudStoragePath object from its proto representation.
func ProtoToDlpBetaStoredInfoTypeDictionaryCloudStoragePath(p *betapb.DlpBetaStoredInfoTypeDictionaryCloudStoragePath) *beta.StoredInfoTypeDictionaryCloudStoragePath {
	if p == nil {
		return nil
	}
	obj := &beta.StoredInfoTypeDictionaryCloudStoragePath{
		Path: dcl.StringOrNil(p.GetPath()),
	}
	return obj
}

// ProtoToStoredInfoTypeRegex converts a StoredInfoTypeRegex object from its proto representation.
func ProtoToDlpBetaStoredInfoTypeRegex(p *betapb.DlpBetaStoredInfoTypeRegex) *beta.StoredInfoTypeRegex {
	if p == nil {
		return nil
	}
	obj := &beta.StoredInfoTypeRegex{
		Pattern: dcl.StringOrNil(p.GetPattern()),
	}
	for _, r := range p.GetGroupIndexes() {
		obj.GroupIndexes = append(obj.GroupIndexes, r)
	}
	return obj
}

// ProtoToStoredInfoType converts a StoredInfoType resource from its proto representation.
func ProtoToStoredInfoType(p *betapb.DlpBetaStoredInfoType) *beta.StoredInfoType {
	obj := &beta.StoredInfoType{
		Name:                  dcl.StringOrNil(p.GetName()),
		DisplayName:           dcl.StringOrNil(p.GetDisplayName()),
		Description:           dcl.StringOrNil(p.GetDescription()),
		LargeCustomDictionary: ProtoToDlpBetaStoredInfoTypeLargeCustomDictionary(p.GetLargeCustomDictionary()),
		Dictionary:            ProtoToDlpBetaStoredInfoTypeDictionary(p.GetDictionary()),
		Regex:                 ProtoToDlpBetaStoredInfoTypeRegex(p.GetRegex()),
		Parent:                dcl.StringOrNil(p.GetParent()),
		Location:              dcl.StringOrNil(p.GetLocation()),
	}
	return obj
}

// StoredInfoTypeLargeCustomDictionaryToProto converts a StoredInfoTypeLargeCustomDictionary object to its proto representation.
func DlpBetaStoredInfoTypeLargeCustomDictionaryToProto(o *beta.StoredInfoTypeLargeCustomDictionary) *betapb.DlpBetaStoredInfoTypeLargeCustomDictionary {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaStoredInfoTypeLargeCustomDictionary{}
	p.SetOutputPath(DlpBetaStoredInfoTypeLargeCustomDictionaryOutputPathToProto(o.OutputPath))
	p.SetCloudStorageFileSet(DlpBetaStoredInfoTypeLargeCustomDictionaryCloudStorageFileSetToProto(o.CloudStorageFileSet))
	p.SetBigQueryField(DlpBetaStoredInfoTypeLargeCustomDictionaryBigQueryFieldToProto(o.BigQueryField))
	return p
}

// StoredInfoTypeLargeCustomDictionaryOutputPathToProto converts a StoredInfoTypeLargeCustomDictionaryOutputPath object to its proto representation.
func DlpBetaStoredInfoTypeLargeCustomDictionaryOutputPathToProto(o *beta.StoredInfoTypeLargeCustomDictionaryOutputPath) *betapb.DlpBetaStoredInfoTypeLargeCustomDictionaryOutputPath {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaStoredInfoTypeLargeCustomDictionaryOutputPath{}
	p.SetPath(dcl.ValueOrEmptyString(o.Path))
	return p
}

// StoredInfoTypeLargeCustomDictionaryCloudStorageFileSetToProto converts a StoredInfoTypeLargeCustomDictionaryCloudStorageFileSet object to its proto representation.
func DlpBetaStoredInfoTypeLargeCustomDictionaryCloudStorageFileSetToProto(o *beta.StoredInfoTypeLargeCustomDictionaryCloudStorageFileSet) *betapb.DlpBetaStoredInfoTypeLargeCustomDictionaryCloudStorageFileSet {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaStoredInfoTypeLargeCustomDictionaryCloudStorageFileSet{}
	p.SetUrl(dcl.ValueOrEmptyString(o.Url))
	return p
}

// StoredInfoTypeLargeCustomDictionaryBigQueryFieldToProto converts a StoredInfoTypeLargeCustomDictionaryBigQueryField object to its proto representation.
func DlpBetaStoredInfoTypeLargeCustomDictionaryBigQueryFieldToProto(o *beta.StoredInfoTypeLargeCustomDictionaryBigQueryField) *betapb.DlpBetaStoredInfoTypeLargeCustomDictionaryBigQueryField {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaStoredInfoTypeLargeCustomDictionaryBigQueryField{}
	p.SetTable(DlpBetaStoredInfoTypeLargeCustomDictionaryBigQueryFieldTableToProto(o.Table))
	p.SetField(DlpBetaStoredInfoTypeLargeCustomDictionaryBigQueryFieldFieldToProto(o.Field))
	return p
}

// StoredInfoTypeLargeCustomDictionaryBigQueryFieldTableToProto converts a StoredInfoTypeLargeCustomDictionaryBigQueryFieldTable object to its proto representation.
func DlpBetaStoredInfoTypeLargeCustomDictionaryBigQueryFieldTableToProto(o *beta.StoredInfoTypeLargeCustomDictionaryBigQueryFieldTable) *betapb.DlpBetaStoredInfoTypeLargeCustomDictionaryBigQueryFieldTable {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaStoredInfoTypeLargeCustomDictionaryBigQueryFieldTable{}
	p.SetProjectId(dcl.ValueOrEmptyString(o.ProjectId))
	p.SetDatasetId(dcl.ValueOrEmptyString(o.DatasetId))
	p.SetTableId(dcl.ValueOrEmptyString(o.TableId))
	return p
}

// StoredInfoTypeLargeCustomDictionaryBigQueryFieldFieldToProto converts a StoredInfoTypeLargeCustomDictionaryBigQueryFieldField object to its proto representation.
func DlpBetaStoredInfoTypeLargeCustomDictionaryBigQueryFieldFieldToProto(o *beta.StoredInfoTypeLargeCustomDictionaryBigQueryFieldField) *betapb.DlpBetaStoredInfoTypeLargeCustomDictionaryBigQueryFieldField {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaStoredInfoTypeLargeCustomDictionaryBigQueryFieldField{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	return p
}

// StoredInfoTypeDictionaryToProto converts a StoredInfoTypeDictionary object to its proto representation.
func DlpBetaStoredInfoTypeDictionaryToProto(o *beta.StoredInfoTypeDictionary) *betapb.DlpBetaStoredInfoTypeDictionary {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaStoredInfoTypeDictionary{}
	p.SetWordList(DlpBetaStoredInfoTypeDictionaryWordListToProto(o.WordList))
	p.SetCloudStoragePath(DlpBetaStoredInfoTypeDictionaryCloudStoragePathToProto(o.CloudStoragePath))
	return p
}

// StoredInfoTypeDictionaryWordListToProto converts a StoredInfoTypeDictionaryWordList object to its proto representation.
func DlpBetaStoredInfoTypeDictionaryWordListToProto(o *beta.StoredInfoTypeDictionaryWordList) *betapb.DlpBetaStoredInfoTypeDictionaryWordList {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaStoredInfoTypeDictionaryWordList{}
	sWords := make([]string, len(o.Words))
	for i, r := range o.Words {
		sWords[i] = r
	}
	p.SetWords(sWords)
	return p
}

// StoredInfoTypeDictionaryCloudStoragePathToProto converts a StoredInfoTypeDictionaryCloudStoragePath object to its proto representation.
func DlpBetaStoredInfoTypeDictionaryCloudStoragePathToProto(o *beta.StoredInfoTypeDictionaryCloudStoragePath) *betapb.DlpBetaStoredInfoTypeDictionaryCloudStoragePath {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaStoredInfoTypeDictionaryCloudStoragePath{}
	p.SetPath(dcl.ValueOrEmptyString(o.Path))
	return p
}

// StoredInfoTypeRegexToProto converts a StoredInfoTypeRegex object to its proto representation.
func DlpBetaStoredInfoTypeRegexToProto(o *beta.StoredInfoTypeRegex) *betapb.DlpBetaStoredInfoTypeRegex {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaStoredInfoTypeRegex{}
	p.SetPattern(dcl.ValueOrEmptyString(o.Pattern))
	sGroupIndexes := make([]int64, len(o.GroupIndexes))
	for i, r := range o.GroupIndexes {
		sGroupIndexes[i] = r
	}
	p.SetGroupIndexes(sGroupIndexes)
	return p
}

// StoredInfoTypeToProto converts a StoredInfoType resource to its proto representation.
func StoredInfoTypeToProto(resource *beta.StoredInfoType) *betapb.DlpBetaStoredInfoType {
	p := &betapb.DlpBetaStoredInfoType{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDisplayName(dcl.ValueOrEmptyString(resource.DisplayName))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetLargeCustomDictionary(DlpBetaStoredInfoTypeLargeCustomDictionaryToProto(resource.LargeCustomDictionary))
	p.SetDictionary(DlpBetaStoredInfoTypeDictionaryToProto(resource.Dictionary))
	p.SetRegex(DlpBetaStoredInfoTypeRegexToProto(resource.Regex))
	p.SetParent(dcl.ValueOrEmptyString(resource.Parent))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))

	return p
}

// applyStoredInfoType handles the gRPC request by passing it to the underlying StoredInfoType Apply() method.
func (s *StoredInfoTypeServer) applyStoredInfoType(ctx context.Context, c *beta.Client, request *betapb.ApplyDlpBetaStoredInfoTypeRequest) (*betapb.DlpBetaStoredInfoType, error) {
	p := ProtoToStoredInfoType(request.GetResource())
	res, err := c.ApplyStoredInfoType(ctx, p)
	if err != nil {
		return nil, err
	}
	r := StoredInfoTypeToProto(res)
	return r, nil
}

// applyDlpBetaStoredInfoType handles the gRPC request by passing it to the underlying StoredInfoType Apply() method.
func (s *StoredInfoTypeServer) ApplyDlpBetaStoredInfoType(ctx context.Context, request *betapb.ApplyDlpBetaStoredInfoTypeRequest) (*betapb.DlpBetaStoredInfoType, error) {
	cl, err := createConfigStoredInfoType(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyStoredInfoType(ctx, cl, request)
}

// DeleteStoredInfoType handles the gRPC request by passing it to the underlying StoredInfoType Delete() method.
func (s *StoredInfoTypeServer) DeleteDlpBetaStoredInfoType(ctx context.Context, request *betapb.DeleteDlpBetaStoredInfoTypeRequest) (*emptypb.Empty, error) {

	cl, err := createConfigStoredInfoType(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteStoredInfoType(ctx, ProtoToStoredInfoType(request.GetResource()))

}

// ListDlpBetaStoredInfoType handles the gRPC request by passing it to the underlying StoredInfoTypeList() method.
func (s *StoredInfoTypeServer) ListDlpBetaStoredInfoType(ctx context.Context, request *betapb.ListDlpBetaStoredInfoTypeRequest) (*betapb.ListDlpBetaStoredInfoTypeResponse, error) {
	cl, err := createConfigStoredInfoType(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListStoredInfoType(ctx, request.GetLocation(), request.GetParent())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.DlpBetaStoredInfoType
	for _, r := range resources.Items {
		rp := StoredInfoTypeToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListDlpBetaStoredInfoTypeResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigStoredInfoType(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
