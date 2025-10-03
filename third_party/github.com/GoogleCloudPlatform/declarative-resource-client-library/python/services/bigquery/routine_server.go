// Copyright 2022 Google LLC. All Rights Reserved.
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
	bigquerypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/bigquery/bigquery_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/bigquery"
)

// RoutineServer implements the gRPC interface for Routine.
type RoutineServer struct{}

// ProtoToRoutineRoutineTypeEnum converts a RoutineRoutineTypeEnum enum from its proto representation.
func ProtoToBigqueryRoutineRoutineTypeEnum(e bigquerypb.BigqueryRoutineRoutineTypeEnum) *bigquery.RoutineRoutineTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := bigquerypb.BigqueryRoutineRoutineTypeEnum_name[int32(e)]; ok {
		e := bigquery.RoutineRoutineTypeEnum(n[len("BigqueryRoutineRoutineTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToRoutineLanguageEnum converts a RoutineLanguageEnum enum from its proto representation.
func ProtoToBigqueryRoutineLanguageEnum(e bigquerypb.BigqueryRoutineLanguageEnum) *bigquery.RoutineLanguageEnum {
	if e == 0 {
		return nil
	}
	if n, ok := bigquerypb.BigqueryRoutineLanguageEnum_name[int32(e)]; ok {
		e := bigquery.RoutineLanguageEnum(n[len("BigqueryRoutineLanguageEnum"):])
		return &e
	}
	return nil
}

// ProtoToRoutineArgumentsArgumentKindEnum converts a RoutineArgumentsArgumentKindEnum enum from its proto representation.
func ProtoToBigqueryRoutineArgumentsArgumentKindEnum(e bigquerypb.BigqueryRoutineArgumentsArgumentKindEnum) *bigquery.RoutineArgumentsArgumentKindEnum {
	if e == 0 {
		return nil
	}
	if n, ok := bigquerypb.BigqueryRoutineArgumentsArgumentKindEnum_name[int32(e)]; ok {
		e := bigquery.RoutineArgumentsArgumentKindEnum(n[len("BigqueryRoutineArgumentsArgumentKindEnum"):])
		return &e
	}
	return nil
}

// ProtoToRoutineArgumentsModeEnum converts a RoutineArgumentsModeEnum enum from its proto representation.
func ProtoToBigqueryRoutineArgumentsModeEnum(e bigquerypb.BigqueryRoutineArgumentsModeEnum) *bigquery.RoutineArgumentsModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := bigquerypb.BigqueryRoutineArgumentsModeEnum_name[int32(e)]; ok {
		e := bigquery.RoutineArgumentsModeEnum(n[len("BigqueryRoutineArgumentsModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToRoutineArgumentsDataTypeTypeKindEnum converts a RoutineArgumentsDataTypeTypeKindEnum enum from its proto representation.
func ProtoToBigqueryRoutineArgumentsDataTypeTypeKindEnum(e bigquerypb.BigqueryRoutineArgumentsDataTypeTypeKindEnum) *bigquery.RoutineArgumentsDataTypeTypeKindEnum {
	if e == 0 {
		return nil
	}
	if n, ok := bigquerypb.BigqueryRoutineArgumentsDataTypeTypeKindEnum_name[int32(e)]; ok {
		e := bigquery.RoutineArgumentsDataTypeTypeKindEnum(n[len("BigqueryRoutineArgumentsDataTypeTypeKindEnum"):])
		return &e
	}
	return nil
}

// ProtoToRoutineDeterminismLevelEnum converts a RoutineDeterminismLevelEnum enum from its proto representation.
func ProtoToBigqueryRoutineDeterminismLevelEnum(e bigquerypb.BigqueryRoutineDeterminismLevelEnum) *bigquery.RoutineDeterminismLevelEnum {
	if e == 0 {
		return nil
	}
	if n, ok := bigquerypb.BigqueryRoutineDeterminismLevelEnum_name[int32(e)]; ok {
		e := bigquery.RoutineDeterminismLevelEnum(n[len("BigqueryRoutineDeterminismLevelEnum"):])
		return &e
	}
	return nil
}

// ProtoToRoutineArguments converts a RoutineArguments object from its proto representation.
func ProtoToBigqueryRoutineArguments(p *bigquerypb.BigqueryRoutineArguments) *bigquery.RoutineArguments {
	if p == nil {
		return nil
	}
	obj := &bigquery.RoutineArguments{
		Name:         dcl.StringOrNil(p.GetName()),
		ArgumentKind: ProtoToBigqueryRoutineArgumentsArgumentKindEnum(p.GetArgumentKind()),
		Mode:         ProtoToBigqueryRoutineArgumentsModeEnum(p.GetMode()),
		DataType:     ProtoToBigqueryRoutineArgumentsDataType(p.GetDataType()),
	}
	return obj
}

// ProtoToRoutineArgumentsDataType converts a RoutineArgumentsDataType object from its proto representation.
func ProtoToBigqueryRoutineArgumentsDataType(p *bigquerypb.BigqueryRoutineArgumentsDataType) *bigquery.RoutineArgumentsDataType {
	if p == nil {
		return nil
	}
	obj := &bigquery.RoutineArgumentsDataType{
		TypeKind:         ProtoToBigqueryRoutineArgumentsDataTypeTypeKindEnum(p.GetTypeKind()),
		ArrayElementType: ProtoToBigqueryRoutineArgumentsDataType(p.GetArrayElementType()),
		StructType:       ProtoToBigqueryRoutineArgumentsDataTypeStructType(p.GetStructType()),
	}
	return obj
}

// ProtoToRoutineArgumentsDataTypeStructType converts a RoutineArgumentsDataTypeStructType object from its proto representation.
func ProtoToBigqueryRoutineArgumentsDataTypeStructType(p *bigquerypb.BigqueryRoutineArgumentsDataTypeStructType) *bigquery.RoutineArgumentsDataTypeStructType {
	if p == nil {
		return nil
	}
	obj := &bigquery.RoutineArgumentsDataTypeStructType{}
	for _, r := range p.GetFields() {
		obj.Fields = append(obj.Fields, *ProtoToBigqueryRoutineArgumentsDataTypeStructTypeFields(r))
	}
	return obj
}

// ProtoToRoutineArgumentsDataTypeStructTypeFields converts a RoutineArgumentsDataTypeStructTypeFields object from its proto representation.
func ProtoToBigqueryRoutineArgumentsDataTypeStructTypeFields(p *bigquerypb.BigqueryRoutineArgumentsDataTypeStructTypeFields) *bigquery.RoutineArgumentsDataTypeStructTypeFields {
	if p == nil {
		return nil
	}
	obj := &bigquery.RoutineArgumentsDataTypeStructTypeFields{
		Name: dcl.StringOrNil(p.GetName()),
		Type: ProtoToBigqueryRoutineArgumentsDataType(p.GetType()),
	}
	return obj
}

// ProtoToRoutine converts a Routine resource from its proto representation.
func ProtoToRoutine(p *bigquerypb.BigqueryRoutine) *bigquery.Routine {
	obj := &bigquery.Routine{
		Etag:             dcl.StringOrNil(p.GetEtag()),
		Name:             dcl.StringOrNil(p.GetName()),
		Project:          dcl.StringOrNil(p.GetProject()),
		Dataset:          dcl.StringOrNil(p.GetDataset()),
		RoutineType:      ProtoToBigqueryRoutineRoutineTypeEnum(p.GetRoutineType()),
		CreationTime:     dcl.Int64OrNil(p.GetCreationTime()),
		LastModifiedTime: dcl.Int64OrNil(p.GetLastModifiedTime()),
		Language:         ProtoToBigqueryRoutineLanguageEnum(p.GetLanguage()),
		ReturnType:       ProtoToBigqueryRoutineArgumentsDataType(p.GetReturnType()),
		DefinitionBody:   dcl.StringOrNil(p.GetDefinitionBody()),
		Description:      dcl.StringOrNil(p.GetDescription()),
		DeterminismLevel: ProtoToBigqueryRoutineDeterminismLevelEnum(p.GetDeterminismLevel()),
		StrictMode:       dcl.Bool(p.GetStrictMode()),
	}
	for _, r := range p.GetArguments() {
		obj.Arguments = append(obj.Arguments, *ProtoToBigqueryRoutineArguments(r))
	}
	for _, r := range p.GetImportedLibraries() {
		obj.ImportedLibraries = append(obj.ImportedLibraries, r)
	}
	return obj
}

// RoutineRoutineTypeEnumToProto converts a RoutineRoutineTypeEnum enum to its proto representation.
func BigqueryRoutineRoutineTypeEnumToProto(e *bigquery.RoutineRoutineTypeEnum) bigquerypb.BigqueryRoutineRoutineTypeEnum {
	if e == nil {
		return bigquerypb.BigqueryRoutineRoutineTypeEnum(0)
	}
	if v, ok := bigquerypb.BigqueryRoutineRoutineTypeEnum_value["RoutineRoutineTypeEnum"+string(*e)]; ok {
		return bigquerypb.BigqueryRoutineRoutineTypeEnum(v)
	}
	return bigquerypb.BigqueryRoutineRoutineTypeEnum(0)
}

// RoutineLanguageEnumToProto converts a RoutineLanguageEnum enum to its proto representation.
func BigqueryRoutineLanguageEnumToProto(e *bigquery.RoutineLanguageEnum) bigquerypb.BigqueryRoutineLanguageEnum {
	if e == nil {
		return bigquerypb.BigqueryRoutineLanguageEnum(0)
	}
	if v, ok := bigquerypb.BigqueryRoutineLanguageEnum_value["RoutineLanguageEnum"+string(*e)]; ok {
		return bigquerypb.BigqueryRoutineLanguageEnum(v)
	}
	return bigquerypb.BigqueryRoutineLanguageEnum(0)
}

// RoutineArgumentsArgumentKindEnumToProto converts a RoutineArgumentsArgumentKindEnum enum to its proto representation.
func BigqueryRoutineArgumentsArgumentKindEnumToProto(e *bigquery.RoutineArgumentsArgumentKindEnum) bigquerypb.BigqueryRoutineArgumentsArgumentKindEnum {
	if e == nil {
		return bigquerypb.BigqueryRoutineArgumentsArgumentKindEnum(0)
	}
	if v, ok := bigquerypb.BigqueryRoutineArgumentsArgumentKindEnum_value["RoutineArgumentsArgumentKindEnum"+string(*e)]; ok {
		return bigquerypb.BigqueryRoutineArgumentsArgumentKindEnum(v)
	}
	return bigquerypb.BigqueryRoutineArgumentsArgumentKindEnum(0)
}

// RoutineArgumentsModeEnumToProto converts a RoutineArgumentsModeEnum enum to its proto representation.
func BigqueryRoutineArgumentsModeEnumToProto(e *bigquery.RoutineArgumentsModeEnum) bigquerypb.BigqueryRoutineArgumentsModeEnum {
	if e == nil {
		return bigquerypb.BigqueryRoutineArgumentsModeEnum(0)
	}
	if v, ok := bigquerypb.BigqueryRoutineArgumentsModeEnum_value["RoutineArgumentsModeEnum"+string(*e)]; ok {
		return bigquerypb.BigqueryRoutineArgumentsModeEnum(v)
	}
	return bigquerypb.BigqueryRoutineArgumentsModeEnum(0)
}

// RoutineArgumentsDataTypeTypeKindEnumToProto converts a RoutineArgumentsDataTypeTypeKindEnum enum to its proto representation.
func BigqueryRoutineArgumentsDataTypeTypeKindEnumToProto(e *bigquery.RoutineArgumentsDataTypeTypeKindEnum) bigquerypb.BigqueryRoutineArgumentsDataTypeTypeKindEnum {
	if e == nil {
		return bigquerypb.BigqueryRoutineArgumentsDataTypeTypeKindEnum(0)
	}
	if v, ok := bigquerypb.BigqueryRoutineArgumentsDataTypeTypeKindEnum_value["RoutineArgumentsDataTypeTypeKindEnum"+string(*e)]; ok {
		return bigquerypb.BigqueryRoutineArgumentsDataTypeTypeKindEnum(v)
	}
	return bigquerypb.BigqueryRoutineArgumentsDataTypeTypeKindEnum(0)
}

// RoutineDeterminismLevelEnumToProto converts a RoutineDeterminismLevelEnum enum to its proto representation.
func BigqueryRoutineDeterminismLevelEnumToProto(e *bigquery.RoutineDeterminismLevelEnum) bigquerypb.BigqueryRoutineDeterminismLevelEnum {
	if e == nil {
		return bigquerypb.BigqueryRoutineDeterminismLevelEnum(0)
	}
	if v, ok := bigquerypb.BigqueryRoutineDeterminismLevelEnum_value["RoutineDeterminismLevelEnum"+string(*e)]; ok {
		return bigquerypb.BigqueryRoutineDeterminismLevelEnum(v)
	}
	return bigquerypb.BigqueryRoutineDeterminismLevelEnum(0)
}

// RoutineArgumentsToProto converts a RoutineArguments object to its proto representation.
func BigqueryRoutineArgumentsToProto(o *bigquery.RoutineArguments) *bigquerypb.BigqueryRoutineArguments {
	if o == nil {
		return nil
	}
	p := &bigquerypb.BigqueryRoutineArguments{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetArgumentKind(BigqueryRoutineArgumentsArgumentKindEnumToProto(o.ArgumentKind))
	p.SetMode(BigqueryRoutineArgumentsModeEnumToProto(o.Mode))
	p.SetDataType(BigqueryRoutineArgumentsDataTypeToProto(o.DataType))
	return p
}

// RoutineArgumentsDataTypeToProto converts a RoutineArgumentsDataType object to its proto representation.
func BigqueryRoutineArgumentsDataTypeToProto(o *bigquery.RoutineArgumentsDataType) *bigquerypb.BigqueryRoutineArgumentsDataType {
	if o == nil {
		return nil
	}
	p := &bigquerypb.BigqueryRoutineArgumentsDataType{}
	p.SetTypeKind(BigqueryRoutineArgumentsDataTypeTypeKindEnumToProto(o.TypeKind))
	p.SetArrayElementType(BigqueryRoutineArgumentsDataTypeToProto(o.ArrayElementType))
	p.SetStructType(BigqueryRoutineArgumentsDataTypeStructTypeToProto(o.StructType))
	return p
}

// RoutineArgumentsDataTypeStructTypeToProto converts a RoutineArgumentsDataTypeStructType object to its proto representation.
func BigqueryRoutineArgumentsDataTypeStructTypeToProto(o *bigquery.RoutineArgumentsDataTypeStructType) *bigquerypb.BigqueryRoutineArgumentsDataTypeStructType {
	if o == nil {
		return nil
	}
	p := &bigquerypb.BigqueryRoutineArgumentsDataTypeStructType{}
	sFields := make([]*bigquerypb.BigqueryRoutineArgumentsDataTypeStructTypeFields, len(o.Fields))
	for i, r := range o.Fields {
		sFields[i] = BigqueryRoutineArgumentsDataTypeStructTypeFieldsToProto(&r)
	}
	p.SetFields(sFields)
	return p
}

// RoutineArgumentsDataTypeStructTypeFieldsToProto converts a RoutineArgumentsDataTypeStructTypeFields object to its proto representation.
func BigqueryRoutineArgumentsDataTypeStructTypeFieldsToProto(o *bigquery.RoutineArgumentsDataTypeStructTypeFields) *bigquerypb.BigqueryRoutineArgumentsDataTypeStructTypeFields {
	if o == nil {
		return nil
	}
	p := &bigquerypb.BigqueryRoutineArgumentsDataTypeStructTypeFields{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetType(BigqueryRoutineArgumentsDataTypeToProto(o.Type))
	return p
}

// RoutineToProto converts a Routine resource to its proto representation.
func RoutineToProto(resource *bigquery.Routine) *bigquerypb.BigqueryRoutine {
	p := &bigquerypb.BigqueryRoutine{}
	p.SetEtag(dcl.ValueOrEmptyString(resource.Etag))
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetDataset(dcl.ValueOrEmptyString(resource.Dataset))
	p.SetRoutineType(BigqueryRoutineRoutineTypeEnumToProto(resource.RoutineType))
	p.SetCreationTime(dcl.ValueOrEmptyInt64(resource.CreationTime))
	p.SetLastModifiedTime(dcl.ValueOrEmptyInt64(resource.LastModifiedTime))
	p.SetLanguage(BigqueryRoutineLanguageEnumToProto(resource.Language))
	p.SetReturnType(BigqueryRoutineArgumentsDataTypeToProto(resource.ReturnType))
	p.SetDefinitionBody(dcl.ValueOrEmptyString(resource.DefinitionBody))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetDeterminismLevel(BigqueryRoutineDeterminismLevelEnumToProto(resource.DeterminismLevel))
	p.SetStrictMode(dcl.ValueOrEmptyBool(resource.StrictMode))
	sArguments := make([]*bigquerypb.BigqueryRoutineArguments, len(resource.Arguments))
	for i, r := range resource.Arguments {
		sArguments[i] = BigqueryRoutineArgumentsToProto(&r)
	}
	p.SetArguments(sArguments)
	sImportedLibraries := make([]string, len(resource.ImportedLibraries))
	for i, r := range resource.ImportedLibraries {
		sImportedLibraries[i] = r
	}
	p.SetImportedLibraries(sImportedLibraries)

	return p
}

// applyRoutine handles the gRPC request by passing it to the underlying Routine Apply() method.
func (s *RoutineServer) applyRoutine(ctx context.Context, c *bigquery.Client, request *bigquerypb.ApplyBigqueryRoutineRequest) (*bigquerypb.BigqueryRoutine, error) {
	p := ProtoToRoutine(request.GetResource())
	res, err := c.ApplyRoutine(ctx, p)
	if err != nil {
		return nil, err
	}
	r := RoutineToProto(res)
	return r, nil
}

// applyBigqueryRoutine handles the gRPC request by passing it to the underlying Routine Apply() method.
func (s *RoutineServer) ApplyBigqueryRoutine(ctx context.Context, request *bigquerypb.ApplyBigqueryRoutineRequest) (*bigquerypb.BigqueryRoutine, error) {
	cl, err := createConfigRoutine(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyRoutine(ctx, cl, request)
}

// DeleteRoutine handles the gRPC request by passing it to the underlying Routine Delete() method.
func (s *RoutineServer) DeleteBigqueryRoutine(ctx context.Context, request *bigquerypb.DeleteBigqueryRoutineRequest) (*emptypb.Empty, error) {

	cl, err := createConfigRoutine(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteRoutine(ctx, ProtoToRoutine(request.GetResource()))

}

// ListBigqueryRoutine handles the gRPC request by passing it to the underlying RoutineList() method.
func (s *RoutineServer) ListBigqueryRoutine(ctx context.Context, request *bigquerypb.ListBigqueryRoutineRequest) (*bigquerypb.ListBigqueryRoutineResponse, error) {
	cl, err := createConfigRoutine(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListRoutine(ctx, request.GetProject(), request.GetDataset())
	if err != nil {
		return nil, err
	}
	var protos []*bigquerypb.BigqueryRoutine
	for _, r := range resources.Items {
		rp := RoutineToProto(r)
		protos = append(protos, rp)
	}
	p := &bigquerypb.ListBigqueryRoutineResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigRoutine(ctx context.Context, service_account_file string) (*bigquery.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return bigquery.NewClient(conf), nil
}
