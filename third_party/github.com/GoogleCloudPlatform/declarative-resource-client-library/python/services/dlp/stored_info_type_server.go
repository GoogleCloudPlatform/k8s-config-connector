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
	dlppb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/dlp/dlp_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/dlp"
)

// StoredInfoTypeServer implements the gRPC interface for StoredInfoType.
type StoredInfoTypeServer struct{}

// ProtoToStoredInfoTypeLargeCustomDictionary converts a StoredInfoTypeLargeCustomDictionary object from its proto representation.
func ProtoToDlpStoredInfoTypeLargeCustomDictionary(p *dlppb.DlpStoredInfoTypeLargeCustomDictionary) *dlp.StoredInfoTypeLargeCustomDictionary {
	if p == nil {
		return nil
	}
	obj := &dlp.StoredInfoTypeLargeCustomDictionary{
		OutputPath:          ProtoToDlpStoredInfoTypeLargeCustomDictionaryOutputPath(p.GetOutputPath()),
		CloudStorageFileSet: ProtoToDlpStoredInfoTypeLargeCustomDictionaryCloudStorageFileSet(p.GetCloudStorageFileSet()),
		BigQueryField:       ProtoToDlpStoredInfoTypeLargeCustomDictionaryBigQueryField(p.GetBigQueryField()),
	}
	return obj
}

// ProtoToStoredInfoTypeLargeCustomDictionaryOutputPath converts a StoredInfoTypeLargeCustomDictionaryOutputPath object from its proto representation.
func ProtoToDlpStoredInfoTypeLargeCustomDictionaryOutputPath(p *dlppb.DlpStoredInfoTypeLargeCustomDictionaryOutputPath) *dlp.StoredInfoTypeLargeCustomDictionaryOutputPath {
	if p == nil {
		return nil
	}
	obj := &dlp.StoredInfoTypeLargeCustomDictionaryOutputPath{
		Path: dcl.StringOrNil(p.GetPath()),
	}
	return obj
}

// ProtoToStoredInfoTypeLargeCustomDictionaryCloudStorageFileSet converts a StoredInfoTypeLargeCustomDictionaryCloudStorageFileSet object from its proto representation.
func ProtoToDlpStoredInfoTypeLargeCustomDictionaryCloudStorageFileSet(p *dlppb.DlpStoredInfoTypeLargeCustomDictionaryCloudStorageFileSet) *dlp.StoredInfoTypeLargeCustomDictionaryCloudStorageFileSet {
	if p == nil {
		return nil
	}
	obj := &dlp.StoredInfoTypeLargeCustomDictionaryCloudStorageFileSet{
		Url: dcl.StringOrNil(p.GetUrl()),
	}
	return obj
}

// ProtoToStoredInfoTypeLargeCustomDictionaryBigQueryField converts a StoredInfoTypeLargeCustomDictionaryBigQueryField object from its proto representation.
func ProtoToDlpStoredInfoTypeLargeCustomDictionaryBigQueryField(p *dlppb.DlpStoredInfoTypeLargeCustomDictionaryBigQueryField) *dlp.StoredInfoTypeLargeCustomDictionaryBigQueryField {
	if p == nil {
		return nil
	}
	obj := &dlp.StoredInfoTypeLargeCustomDictionaryBigQueryField{
		Table: ProtoToDlpStoredInfoTypeLargeCustomDictionaryBigQueryFieldTable(p.GetTable()),
		Field: ProtoToDlpStoredInfoTypeLargeCustomDictionaryBigQueryFieldField(p.GetField()),
	}
	return obj
}

// ProtoToStoredInfoTypeLargeCustomDictionaryBigQueryFieldTable converts a StoredInfoTypeLargeCustomDictionaryBigQueryFieldTable object from its proto representation.
func ProtoToDlpStoredInfoTypeLargeCustomDictionaryBigQueryFieldTable(p *dlppb.DlpStoredInfoTypeLargeCustomDictionaryBigQueryFieldTable) *dlp.StoredInfoTypeLargeCustomDictionaryBigQueryFieldTable {
	if p == nil {
		return nil
	}
	obj := &dlp.StoredInfoTypeLargeCustomDictionaryBigQueryFieldTable{
		ProjectId: dcl.StringOrNil(p.GetProjectId()),
		DatasetId: dcl.StringOrNil(p.GetDatasetId()),
		TableId:   dcl.StringOrNil(p.GetTableId()),
	}
	return obj
}

// ProtoToStoredInfoTypeLargeCustomDictionaryBigQueryFieldField converts a StoredInfoTypeLargeCustomDictionaryBigQueryFieldField object from its proto representation.
func ProtoToDlpStoredInfoTypeLargeCustomDictionaryBigQueryFieldField(p *dlppb.DlpStoredInfoTypeLargeCustomDictionaryBigQueryFieldField) *dlp.StoredInfoTypeLargeCustomDictionaryBigQueryFieldField {
	if p == nil {
		return nil
	}
	obj := &dlp.StoredInfoTypeLargeCustomDictionaryBigQueryFieldField{
		Name: dcl.StringOrNil(p.GetName()),
	}
	return obj
}

// ProtoToStoredInfoTypeDictionary converts a StoredInfoTypeDictionary object from its proto representation.
func ProtoToDlpStoredInfoTypeDictionary(p *dlppb.DlpStoredInfoTypeDictionary) *dlp.StoredInfoTypeDictionary {
	if p == nil {
		return nil
	}
	obj := &dlp.StoredInfoTypeDictionary{
		WordList:         ProtoToDlpStoredInfoTypeDictionaryWordList(p.GetWordList()),
		CloudStoragePath: ProtoToDlpStoredInfoTypeDictionaryCloudStoragePath(p.GetCloudStoragePath()),
	}
	return obj
}

// ProtoToStoredInfoTypeDictionaryWordList converts a StoredInfoTypeDictionaryWordList object from its proto representation.
func ProtoToDlpStoredInfoTypeDictionaryWordList(p *dlppb.DlpStoredInfoTypeDictionaryWordList) *dlp.StoredInfoTypeDictionaryWordList {
	if p == nil {
		return nil
	}
	obj := &dlp.StoredInfoTypeDictionaryWordList{}
	for _, r := range p.GetWords() {
		obj.Words = append(obj.Words, r)
	}
	return obj
}

// ProtoToStoredInfoTypeDictionaryCloudStoragePath converts a StoredInfoTypeDictionaryCloudStoragePath object from its proto representation.
func ProtoToDlpStoredInfoTypeDictionaryCloudStoragePath(p *dlppb.DlpStoredInfoTypeDictionaryCloudStoragePath) *dlp.StoredInfoTypeDictionaryCloudStoragePath {
	if p == nil {
		return nil
	}
	obj := &dlp.StoredInfoTypeDictionaryCloudStoragePath{
		Path: dcl.StringOrNil(p.GetPath()),
	}
	return obj
}

// ProtoToStoredInfoTypeRegex converts a StoredInfoTypeRegex object from its proto representation.
func ProtoToDlpStoredInfoTypeRegex(p *dlppb.DlpStoredInfoTypeRegex) *dlp.StoredInfoTypeRegex {
	if p == nil {
		return nil
	}
	obj := &dlp.StoredInfoTypeRegex{
		Pattern: dcl.StringOrNil(p.GetPattern()),
	}
	for _, r := range p.GetGroupIndexes() {
		obj.GroupIndexes = append(obj.GroupIndexes, r)
	}
	return obj
}

// ProtoToStoredInfoType converts a StoredInfoType resource from its proto representation.
func ProtoToStoredInfoType(p *dlppb.DlpStoredInfoType) *dlp.StoredInfoType {
	obj := &dlp.StoredInfoType{
		Name:                  dcl.StringOrNil(p.GetName()),
		DisplayName:           dcl.StringOrNil(p.GetDisplayName()),
		Description:           dcl.StringOrNil(p.GetDescription()),
		LargeCustomDictionary: ProtoToDlpStoredInfoTypeLargeCustomDictionary(p.GetLargeCustomDictionary()),
		Dictionary:            ProtoToDlpStoredInfoTypeDictionary(p.GetDictionary()),
		Regex:                 ProtoToDlpStoredInfoTypeRegex(p.GetRegex()),
		Parent:                dcl.StringOrNil(p.GetParent()),
		Location:              dcl.StringOrNil(p.GetLocation()),
	}
	return obj
}

// StoredInfoTypeLargeCustomDictionaryToProto converts a StoredInfoTypeLargeCustomDictionary object to its proto representation.
func DlpStoredInfoTypeLargeCustomDictionaryToProto(o *dlp.StoredInfoTypeLargeCustomDictionary) *dlppb.DlpStoredInfoTypeLargeCustomDictionary {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpStoredInfoTypeLargeCustomDictionary{}
	p.SetOutputPath(DlpStoredInfoTypeLargeCustomDictionaryOutputPathToProto(o.OutputPath))
	p.SetCloudStorageFileSet(DlpStoredInfoTypeLargeCustomDictionaryCloudStorageFileSetToProto(o.CloudStorageFileSet))
	p.SetBigQueryField(DlpStoredInfoTypeLargeCustomDictionaryBigQueryFieldToProto(o.BigQueryField))
	return p
}

// StoredInfoTypeLargeCustomDictionaryOutputPathToProto converts a StoredInfoTypeLargeCustomDictionaryOutputPath object to its proto representation.
func DlpStoredInfoTypeLargeCustomDictionaryOutputPathToProto(o *dlp.StoredInfoTypeLargeCustomDictionaryOutputPath) *dlppb.DlpStoredInfoTypeLargeCustomDictionaryOutputPath {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpStoredInfoTypeLargeCustomDictionaryOutputPath{}
	p.SetPath(dcl.ValueOrEmptyString(o.Path))
	return p
}

// StoredInfoTypeLargeCustomDictionaryCloudStorageFileSetToProto converts a StoredInfoTypeLargeCustomDictionaryCloudStorageFileSet object to its proto representation.
func DlpStoredInfoTypeLargeCustomDictionaryCloudStorageFileSetToProto(o *dlp.StoredInfoTypeLargeCustomDictionaryCloudStorageFileSet) *dlppb.DlpStoredInfoTypeLargeCustomDictionaryCloudStorageFileSet {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpStoredInfoTypeLargeCustomDictionaryCloudStorageFileSet{}
	p.SetUrl(dcl.ValueOrEmptyString(o.Url))
	return p
}

// StoredInfoTypeLargeCustomDictionaryBigQueryFieldToProto converts a StoredInfoTypeLargeCustomDictionaryBigQueryField object to its proto representation.
func DlpStoredInfoTypeLargeCustomDictionaryBigQueryFieldToProto(o *dlp.StoredInfoTypeLargeCustomDictionaryBigQueryField) *dlppb.DlpStoredInfoTypeLargeCustomDictionaryBigQueryField {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpStoredInfoTypeLargeCustomDictionaryBigQueryField{}
	p.SetTable(DlpStoredInfoTypeLargeCustomDictionaryBigQueryFieldTableToProto(o.Table))
	p.SetField(DlpStoredInfoTypeLargeCustomDictionaryBigQueryFieldFieldToProto(o.Field))
	return p
}

// StoredInfoTypeLargeCustomDictionaryBigQueryFieldTableToProto converts a StoredInfoTypeLargeCustomDictionaryBigQueryFieldTable object to its proto representation.
func DlpStoredInfoTypeLargeCustomDictionaryBigQueryFieldTableToProto(o *dlp.StoredInfoTypeLargeCustomDictionaryBigQueryFieldTable) *dlppb.DlpStoredInfoTypeLargeCustomDictionaryBigQueryFieldTable {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpStoredInfoTypeLargeCustomDictionaryBigQueryFieldTable{}
	p.SetProjectId(dcl.ValueOrEmptyString(o.ProjectId))
	p.SetDatasetId(dcl.ValueOrEmptyString(o.DatasetId))
	p.SetTableId(dcl.ValueOrEmptyString(o.TableId))
	return p
}

// StoredInfoTypeLargeCustomDictionaryBigQueryFieldFieldToProto converts a StoredInfoTypeLargeCustomDictionaryBigQueryFieldField object to its proto representation.
func DlpStoredInfoTypeLargeCustomDictionaryBigQueryFieldFieldToProto(o *dlp.StoredInfoTypeLargeCustomDictionaryBigQueryFieldField) *dlppb.DlpStoredInfoTypeLargeCustomDictionaryBigQueryFieldField {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpStoredInfoTypeLargeCustomDictionaryBigQueryFieldField{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	return p
}

// StoredInfoTypeDictionaryToProto converts a StoredInfoTypeDictionary object to its proto representation.
func DlpStoredInfoTypeDictionaryToProto(o *dlp.StoredInfoTypeDictionary) *dlppb.DlpStoredInfoTypeDictionary {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpStoredInfoTypeDictionary{}
	p.SetWordList(DlpStoredInfoTypeDictionaryWordListToProto(o.WordList))
	p.SetCloudStoragePath(DlpStoredInfoTypeDictionaryCloudStoragePathToProto(o.CloudStoragePath))
	return p
}

// StoredInfoTypeDictionaryWordListToProto converts a StoredInfoTypeDictionaryWordList object to its proto representation.
func DlpStoredInfoTypeDictionaryWordListToProto(o *dlp.StoredInfoTypeDictionaryWordList) *dlppb.DlpStoredInfoTypeDictionaryWordList {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpStoredInfoTypeDictionaryWordList{}
	sWords := make([]string, len(o.Words))
	for i, r := range o.Words {
		sWords[i] = r
	}
	p.SetWords(sWords)
	return p
}

// StoredInfoTypeDictionaryCloudStoragePathToProto converts a StoredInfoTypeDictionaryCloudStoragePath object to its proto representation.
func DlpStoredInfoTypeDictionaryCloudStoragePathToProto(o *dlp.StoredInfoTypeDictionaryCloudStoragePath) *dlppb.DlpStoredInfoTypeDictionaryCloudStoragePath {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpStoredInfoTypeDictionaryCloudStoragePath{}
	p.SetPath(dcl.ValueOrEmptyString(o.Path))
	return p
}

// StoredInfoTypeRegexToProto converts a StoredInfoTypeRegex object to its proto representation.
func DlpStoredInfoTypeRegexToProto(o *dlp.StoredInfoTypeRegex) *dlppb.DlpStoredInfoTypeRegex {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpStoredInfoTypeRegex{}
	p.SetPattern(dcl.ValueOrEmptyString(o.Pattern))
	sGroupIndexes := make([]int64, len(o.GroupIndexes))
	for i, r := range o.GroupIndexes {
		sGroupIndexes[i] = r
	}
	p.SetGroupIndexes(sGroupIndexes)
	return p
}

// StoredInfoTypeToProto converts a StoredInfoType resource to its proto representation.
func StoredInfoTypeToProto(resource *dlp.StoredInfoType) *dlppb.DlpStoredInfoType {
	p := &dlppb.DlpStoredInfoType{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDisplayName(dcl.ValueOrEmptyString(resource.DisplayName))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetLargeCustomDictionary(DlpStoredInfoTypeLargeCustomDictionaryToProto(resource.LargeCustomDictionary))
	p.SetDictionary(DlpStoredInfoTypeDictionaryToProto(resource.Dictionary))
	p.SetRegex(DlpStoredInfoTypeRegexToProto(resource.Regex))
	p.SetParent(dcl.ValueOrEmptyString(resource.Parent))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))

	return p
}

// applyStoredInfoType handles the gRPC request by passing it to the underlying StoredInfoType Apply() method.
func (s *StoredInfoTypeServer) applyStoredInfoType(ctx context.Context, c *dlp.Client, request *dlppb.ApplyDlpStoredInfoTypeRequest) (*dlppb.DlpStoredInfoType, error) {
	p := ProtoToStoredInfoType(request.GetResource())
	res, err := c.ApplyStoredInfoType(ctx, p)
	if err != nil {
		return nil, err
	}
	r := StoredInfoTypeToProto(res)
	return r, nil
}

// applyDlpStoredInfoType handles the gRPC request by passing it to the underlying StoredInfoType Apply() method.
func (s *StoredInfoTypeServer) ApplyDlpStoredInfoType(ctx context.Context, request *dlppb.ApplyDlpStoredInfoTypeRequest) (*dlppb.DlpStoredInfoType, error) {
	cl, err := createConfigStoredInfoType(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyStoredInfoType(ctx, cl, request)
}

// DeleteStoredInfoType handles the gRPC request by passing it to the underlying StoredInfoType Delete() method.
func (s *StoredInfoTypeServer) DeleteDlpStoredInfoType(ctx context.Context, request *dlppb.DeleteDlpStoredInfoTypeRequest) (*emptypb.Empty, error) {

	cl, err := createConfigStoredInfoType(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteStoredInfoType(ctx, ProtoToStoredInfoType(request.GetResource()))

}

// ListDlpStoredInfoType handles the gRPC request by passing it to the underlying StoredInfoTypeList() method.
func (s *StoredInfoTypeServer) ListDlpStoredInfoType(ctx context.Context, request *dlppb.ListDlpStoredInfoTypeRequest) (*dlppb.ListDlpStoredInfoTypeResponse, error) {
	cl, err := createConfigStoredInfoType(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListStoredInfoType(ctx, request.GetLocation(), request.GetParent())
	if err != nil {
		return nil, err
	}
	var protos []*dlppb.DlpStoredInfoType
	for _, r := range resources.Items {
		rp := StoredInfoTypeToProto(r)
		protos = append(protos, rp)
	}
	p := &dlppb.ListDlpStoredInfoTypeResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigStoredInfoType(ctx context.Context, service_account_file string) (*dlp.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return dlp.NewClient(conf), nil
}
