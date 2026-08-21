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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/dlp/alpha/dlp_alpha_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/dlp/alpha"
)

// StoredInfoTypeServer implements the gRPC interface for StoredInfoType.
type StoredInfoTypeServer struct{}

// ProtoToStoredInfoTypeLargeCustomDictionary converts a StoredInfoTypeLargeCustomDictionary object from its proto representation.
func ProtoToDlpAlphaStoredInfoTypeLargeCustomDictionary(p *alphapb.DlpAlphaStoredInfoTypeLargeCustomDictionary) *alpha.StoredInfoTypeLargeCustomDictionary {
	if p == nil {
		return nil
	}
	obj := &alpha.StoredInfoTypeLargeCustomDictionary{
		OutputPath:          ProtoToDlpAlphaStoredInfoTypeLargeCustomDictionaryOutputPath(p.GetOutputPath()),
		CloudStorageFileSet: ProtoToDlpAlphaStoredInfoTypeLargeCustomDictionaryCloudStorageFileSet(p.GetCloudStorageFileSet()),
		BigQueryField:       ProtoToDlpAlphaStoredInfoTypeLargeCustomDictionaryBigQueryField(p.GetBigQueryField()),
	}
	return obj
}

// ProtoToStoredInfoTypeLargeCustomDictionaryOutputPath converts a StoredInfoTypeLargeCustomDictionaryOutputPath object from its proto representation.
func ProtoToDlpAlphaStoredInfoTypeLargeCustomDictionaryOutputPath(p *alphapb.DlpAlphaStoredInfoTypeLargeCustomDictionaryOutputPath) *alpha.StoredInfoTypeLargeCustomDictionaryOutputPath {
	if p == nil {
		return nil
	}
	obj := &alpha.StoredInfoTypeLargeCustomDictionaryOutputPath{
		Path: dcl.StringOrNil(p.GetPath()),
	}
	return obj
}

// ProtoToStoredInfoTypeLargeCustomDictionaryCloudStorageFileSet converts a StoredInfoTypeLargeCustomDictionaryCloudStorageFileSet object from its proto representation.
func ProtoToDlpAlphaStoredInfoTypeLargeCustomDictionaryCloudStorageFileSet(p *alphapb.DlpAlphaStoredInfoTypeLargeCustomDictionaryCloudStorageFileSet) *alpha.StoredInfoTypeLargeCustomDictionaryCloudStorageFileSet {
	if p == nil {
		return nil
	}
	obj := &alpha.StoredInfoTypeLargeCustomDictionaryCloudStorageFileSet{
		Url: dcl.StringOrNil(p.GetUrl()),
	}
	return obj
}

// ProtoToStoredInfoTypeLargeCustomDictionaryBigQueryField converts a StoredInfoTypeLargeCustomDictionaryBigQueryField object from its proto representation.
func ProtoToDlpAlphaStoredInfoTypeLargeCustomDictionaryBigQueryField(p *alphapb.DlpAlphaStoredInfoTypeLargeCustomDictionaryBigQueryField) *alpha.StoredInfoTypeLargeCustomDictionaryBigQueryField {
	if p == nil {
		return nil
	}
	obj := &alpha.StoredInfoTypeLargeCustomDictionaryBigQueryField{
		Table: ProtoToDlpAlphaStoredInfoTypeLargeCustomDictionaryBigQueryFieldTable(p.GetTable()),
		Field: ProtoToDlpAlphaStoredInfoTypeLargeCustomDictionaryBigQueryFieldField(p.GetField()),
	}
	return obj
}

// ProtoToStoredInfoTypeLargeCustomDictionaryBigQueryFieldTable converts a StoredInfoTypeLargeCustomDictionaryBigQueryFieldTable object from its proto representation.
func ProtoToDlpAlphaStoredInfoTypeLargeCustomDictionaryBigQueryFieldTable(p *alphapb.DlpAlphaStoredInfoTypeLargeCustomDictionaryBigQueryFieldTable) *alpha.StoredInfoTypeLargeCustomDictionaryBigQueryFieldTable {
	if p == nil {
		return nil
	}
	obj := &alpha.StoredInfoTypeLargeCustomDictionaryBigQueryFieldTable{
		ProjectId: dcl.StringOrNil(p.GetProjectId()),
		DatasetId: dcl.StringOrNil(p.GetDatasetId()),
		TableId:   dcl.StringOrNil(p.GetTableId()),
	}
	return obj
}

// ProtoToStoredInfoTypeLargeCustomDictionaryBigQueryFieldField converts a StoredInfoTypeLargeCustomDictionaryBigQueryFieldField object from its proto representation.
func ProtoToDlpAlphaStoredInfoTypeLargeCustomDictionaryBigQueryFieldField(p *alphapb.DlpAlphaStoredInfoTypeLargeCustomDictionaryBigQueryFieldField) *alpha.StoredInfoTypeLargeCustomDictionaryBigQueryFieldField {
	if p == nil {
		return nil
	}
	obj := &alpha.StoredInfoTypeLargeCustomDictionaryBigQueryFieldField{
		Name: dcl.StringOrNil(p.GetName()),
	}
	return obj
}

// ProtoToStoredInfoTypeDictionary converts a StoredInfoTypeDictionary object from its proto representation.
func ProtoToDlpAlphaStoredInfoTypeDictionary(p *alphapb.DlpAlphaStoredInfoTypeDictionary) *alpha.StoredInfoTypeDictionary {
	if p == nil {
		return nil
	}
	obj := &alpha.StoredInfoTypeDictionary{
		WordList:         ProtoToDlpAlphaStoredInfoTypeDictionaryWordList(p.GetWordList()),
		CloudStoragePath: ProtoToDlpAlphaStoredInfoTypeDictionaryCloudStoragePath(p.GetCloudStoragePath()),
	}
	return obj
}

// ProtoToStoredInfoTypeDictionaryWordList converts a StoredInfoTypeDictionaryWordList object from its proto representation.
func ProtoToDlpAlphaStoredInfoTypeDictionaryWordList(p *alphapb.DlpAlphaStoredInfoTypeDictionaryWordList) *alpha.StoredInfoTypeDictionaryWordList {
	if p == nil {
		return nil
	}
	obj := &alpha.StoredInfoTypeDictionaryWordList{}
	for _, r := range p.GetWords() {
		obj.Words = append(obj.Words, r)
	}
	return obj
}

// ProtoToStoredInfoTypeDictionaryCloudStoragePath converts a StoredInfoTypeDictionaryCloudStoragePath object from its proto representation.
func ProtoToDlpAlphaStoredInfoTypeDictionaryCloudStoragePath(p *alphapb.DlpAlphaStoredInfoTypeDictionaryCloudStoragePath) *alpha.StoredInfoTypeDictionaryCloudStoragePath {
	if p == nil {
		return nil
	}
	obj := &alpha.StoredInfoTypeDictionaryCloudStoragePath{
		Path: dcl.StringOrNil(p.GetPath()),
	}
	return obj
}

// ProtoToStoredInfoTypeRegex converts a StoredInfoTypeRegex object from its proto representation.
func ProtoToDlpAlphaStoredInfoTypeRegex(p *alphapb.DlpAlphaStoredInfoTypeRegex) *alpha.StoredInfoTypeRegex {
	if p == nil {
		return nil
	}
	obj := &alpha.StoredInfoTypeRegex{
		Pattern: dcl.StringOrNil(p.GetPattern()),
	}
	for _, r := range p.GetGroupIndexes() {
		obj.GroupIndexes = append(obj.GroupIndexes, r)
	}
	return obj
}

// ProtoToStoredInfoType converts a StoredInfoType resource from its proto representation.
func ProtoToStoredInfoType(p *alphapb.DlpAlphaStoredInfoType) *alpha.StoredInfoType {
	obj := &alpha.StoredInfoType{
		Name:                  dcl.StringOrNil(p.GetName()),
		DisplayName:           dcl.StringOrNil(p.GetDisplayName()),
		Description:           dcl.StringOrNil(p.GetDescription()),
		LargeCustomDictionary: ProtoToDlpAlphaStoredInfoTypeLargeCustomDictionary(p.GetLargeCustomDictionary()),
		Dictionary:            ProtoToDlpAlphaStoredInfoTypeDictionary(p.GetDictionary()),
		Regex:                 ProtoToDlpAlphaStoredInfoTypeRegex(p.GetRegex()),
		Parent:                dcl.StringOrNil(p.GetParent()),
		Location:              dcl.StringOrNil(p.GetLocation()),
	}
	return obj
}

// StoredInfoTypeLargeCustomDictionaryToProto converts a StoredInfoTypeLargeCustomDictionary object to its proto representation.
func DlpAlphaStoredInfoTypeLargeCustomDictionaryToProto(o *alpha.StoredInfoTypeLargeCustomDictionary) *alphapb.DlpAlphaStoredInfoTypeLargeCustomDictionary {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaStoredInfoTypeLargeCustomDictionary{}
	p.SetOutputPath(DlpAlphaStoredInfoTypeLargeCustomDictionaryOutputPathToProto(o.OutputPath))
	p.SetCloudStorageFileSet(DlpAlphaStoredInfoTypeLargeCustomDictionaryCloudStorageFileSetToProto(o.CloudStorageFileSet))
	p.SetBigQueryField(DlpAlphaStoredInfoTypeLargeCustomDictionaryBigQueryFieldToProto(o.BigQueryField))
	return p
}

// StoredInfoTypeLargeCustomDictionaryOutputPathToProto converts a StoredInfoTypeLargeCustomDictionaryOutputPath object to its proto representation.
func DlpAlphaStoredInfoTypeLargeCustomDictionaryOutputPathToProto(o *alpha.StoredInfoTypeLargeCustomDictionaryOutputPath) *alphapb.DlpAlphaStoredInfoTypeLargeCustomDictionaryOutputPath {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaStoredInfoTypeLargeCustomDictionaryOutputPath{}
	p.SetPath(dcl.ValueOrEmptyString(o.Path))
	return p
}

// StoredInfoTypeLargeCustomDictionaryCloudStorageFileSetToProto converts a StoredInfoTypeLargeCustomDictionaryCloudStorageFileSet object to its proto representation.
func DlpAlphaStoredInfoTypeLargeCustomDictionaryCloudStorageFileSetToProto(o *alpha.StoredInfoTypeLargeCustomDictionaryCloudStorageFileSet) *alphapb.DlpAlphaStoredInfoTypeLargeCustomDictionaryCloudStorageFileSet {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaStoredInfoTypeLargeCustomDictionaryCloudStorageFileSet{}
	p.SetUrl(dcl.ValueOrEmptyString(o.Url))
	return p
}

// StoredInfoTypeLargeCustomDictionaryBigQueryFieldToProto converts a StoredInfoTypeLargeCustomDictionaryBigQueryField object to its proto representation.
func DlpAlphaStoredInfoTypeLargeCustomDictionaryBigQueryFieldToProto(o *alpha.StoredInfoTypeLargeCustomDictionaryBigQueryField) *alphapb.DlpAlphaStoredInfoTypeLargeCustomDictionaryBigQueryField {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaStoredInfoTypeLargeCustomDictionaryBigQueryField{}
	p.SetTable(DlpAlphaStoredInfoTypeLargeCustomDictionaryBigQueryFieldTableToProto(o.Table))
	p.SetField(DlpAlphaStoredInfoTypeLargeCustomDictionaryBigQueryFieldFieldToProto(o.Field))
	return p
}

// StoredInfoTypeLargeCustomDictionaryBigQueryFieldTableToProto converts a StoredInfoTypeLargeCustomDictionaryBigQueryFieldTable object to its proto representation.
func DlpAlphaStoredInfoTypeLargeCustomDictionaryBigQueryFieldTableToProto(o *alpha.StoredInfoTypeLargeCustomDictionaryBigQueryFieldTable) *alphapb.DlpAlphaStoredInfoTypeLargeCustomDictionaryBigQueryFieldTable {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaStoredInfoTypeLargeCustomDictionaryBigQueryFieldTable{}
	p.SetProjectId(dcl.ValueOrEmptyString(o.ProjectId))
	p.SetDatasetId(dcl.ValueOrEmptyString(o.DatasetId))
	p.SetTableId(dcl.ValueOrEmptyString(o.TableId))
	return p
}

// StoredInfoTypeLargeCustomDictionaryBigQueryFieldFieldToProto converts a StoredInfoTypeLargeCustomDictionaryBigQueryFieldField object to its proto representation.
func DlpAlphaStoredInfoTypeLargeCustomDictionaryBigQueryFieldFieldToProto(o *alpha.StoredInfoTypeLargeCustomDictionaryBigQueryFieldField) *alphapb.DlpAlphaStoredInfoTypeLargeCustomDictionaryBigQueryFieldField {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaStoredInfoTypeLargeCustomDictionaryBigQueryFieldField{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	return p
}

// StoredInfoTypeDictionaryToProto converts a StoredInfoTypeDictionary object to its proto representation.
func DlpAlphaStoredInfoTypeDictionaryToProto(o *alpha.StoredInfoTypeDictionary) *alphapb.DlpAlphaStoredInfoTypeDictionary {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaStoredInfoTypeDictionary{}
	p.SetWordList(DlpAlphaStoredInfoTypeDictionaryWordListToProto(o.WordList))
	p.SetCloudStoragePath(DlpAlphaStoredInfoTypeDictionaryCloudStoragePathToProto(o.CloudStoragePath))
	return p
}

// StoredInfoTypeDictionaryWordListToProto converts a StoredInfoTypeDictionaryWordList object to its proto representation.
func DlpAlphaStoredInfoTypeDictionaryWordListToProto(o *alpha.StoredInfoTypeDictionaryWordList) *alphapb.DlpAlphaStoredInfoTypeDictionaryWordList {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaStoredInfoTypeDictionaryWordList{}
	sWords := make([]string, len(o.Words))
	for i, r := range o.Words {
		sWords[i] = r
	}
	p.SetWords(sWords)
	return p
}

// StoredInfoTypeDictionaryCloudStoragePathToProto converts a StoredInfoTypeDictionaryCloudStoragePath object to its proto representation.
func DlpAlphaStoredInfoTypeDictionaryCloudStoragePathToProto(o *alpha.StoredInfoTypeDictionaryCloudStoragePath) *alphapb.DlpAlphaStoredInfoTypeDictionaryCloudStoragePath {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaStoredInfoTypeDictionaryCloudStoragePath{}
	p.SetPath(dcl.ValueOrEmptyString(o.Path))
	return p
}

// StoredInfoTypeRegexToProto converts a StoredInfoTypeRegex object to its proto representation.
func DlpAlphaStoredInfoTypeRegexToProto(o *alpha.StoredInfoTypeRegex) *alphapb.DlpAlphaStoredInfoTypeRegex {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaStoredInfoTypeRegex{}
	p.SetPattern(dcl.ValueOrEmptyString(o.Pattern))
	sGroupIndexes := make([]int64, len(o.GroupIndexes))
	for i, r := range o.GroupIndexes {
		sGroupIndexes[i] = r
	}
	p.SetGroupIndexes(sGroupIndexes)
	return p
}

// StoredInfoTypeToProto converts a StoredInfoType resource to its proto representation.
func StoredInfoTypeToProto(resource *alpha.StoredInfoType) *alphapb.DlpAlphaStoredInfoType {
	p := &alphapb.DlpAlphaStoredInfoType{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDisplayName(dcl.ValueOrEmptyString(resource.DisplayName))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetLargeCustomDictionary(DlpAlphaStoredInfoTypeLargeCustomDictionaryToProto(resource.LargeCustomDictionary))
	p.SetDictionary(DlpAlphaStoredInfoTypeDictionaryToProto(resource.Dictionary))
	p.SetRegex(DlpAlphaStoredInfoTypeRegexToProto(resource.Regex))
	p.SetParent(dcl.ValueOrEmptyString(resource.Parent))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))

	return p
}

// applyStoredInfoType handles the gRPC request by passing it to the underlying StoredInfoType Apply() method.
func (s *StoredInfoTypeServer) applyStoredInfoType(ctx context.Context, c *alpha.Client, request *alphapb.ApplyDlpAlphaStoredInfoTypeRequest) (*alphapb.DlpAlphaStoredInfoType, error) {
	p := ProtoToStoredInfoType(request.GetResource())
	res, err := c.ApplyStoredInfoType(ctx, p)
	if err != nil {
		return nil, err
	}
	r := StoredInfoTypeToProto(res)
	return r, nil
}

// applyDlpAlphaStoredInfoType handles the gRPC request by passing it to the underlying StoredInfoType Apply() method.
func (s *StoredInfoTypeServer) ApplyDlpAlphaStoredInfoType(ctx context.Context, request *alphapb.ApplyDlpAlphaStoredInfoTypeRequest) (*alphapb.DlpAlphaStoredInfoType, error) {
	cl, err := createConfigStoredInfoType(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyStoredInfoType(ctx, cl, request)
}

// DeleteStoredInfoType handles the gRPC request by passing it to the underlying StoredInfoType Delete() method.
func (s *StoredInfoTypeServer) DeleteDlpAlphaStoredInfoType(ctx context.Context, request *alphapb.DeleteDlpAlphaStoredInfoTypeRequest) (*emptypb.Empty, error) {

	cl, err := createConfigStoredInfoType(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteStoredInfoType(ctx, ProtoToStoredInfoType(request.GetResource()))

}

// ListDlpAlphaStoredInfoType handles the gRPC request by passing it to the underlying StoredInfoTypeList() method.
func (s *StoredInfoTypeServer) ListDlpAlphaStoredInfoType(ctx context.Context, request *alphapb.ListDlpAlphaStoredInfoTypeRequest) (*alphapb.ListDlpAlphaStoredInfoTypeResponse, error) {
	cl, err := createConfigStoredInfoType(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListStoredInfoType(ctx, request.GetLocation(), request.GetParent())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.DlpAlphaStoredInfoType
	for _, r := range resources.Items {
		rp := StoredInfoTypeToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListDlpAlphaStoredInfoTypeResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigStoredInfoType(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
