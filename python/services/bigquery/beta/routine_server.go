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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/bigquery/beta/bigquery_beta_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/bigquery/beta"
)

// RoutineServer implements the gRPC interface for Routine.
type RoutineServer struct{}

// ProtoToRoutineRoutineTypeEnum converts a RoutineRoutineTypeEnum enum from its proto representation.
func ProtoToBigqueryBetaRoutineRoutineTypeEnum(e betapb.BigqueryBetaRoutineRoutineTypeEnum) *beta.RoutineRoutineTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.BigqueryBetaRoutineRoutineTypeEnum_name[int32(e)]; ok {
		e := beta.RoutineRoutineTypeEnum(n[len("BigqueryBetaRoutineRoutineTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToRoutineLanguageEnum converts a RoutineLanguageEnum enum from its proto representation.
func ProtoToBigqueryBetaRoutineLanguageEnum(e betapb.BigqueryBetaRoutineLanguageEnum) *beta.RoutineLanguageEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.BigqueryBetaRoutineLanguageEnum_name[int32(e)]; ok {
		e := beta.RoutineLanguageEnum(n[len("BigqueryBetaRoutineLanguageEnum"):])
		return &e
	}
	return nil
}

// ProtoToRoutineArgumentsArgumentKindEnum converts a RoutineArgumentsArgumentKindEnum enum from its proto representation.
func ProtoToBigqueryBetaRoutineArgumentsArgumentKindEnum(e betapb.BigqueryBetaRoutineArgumentsArgumentKindEnum) *beta.RoutineArgumentsArgumentKindEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.BigqueryBetaRoutineArgumentsArgumentKindEnum_name[int32(e)]; ok {
		e := beta.RoutineArgumentsArgumentKindEnum(n[len("BigqueryBetaRoutineArgumentsArgumentKindEnum"):])
		return &e
	}
	return nil
}

// ProtoToRoutineArgumentsModeEnum converts a RoutineArgumentsModeEnum enum from its proto representation.
func ProtoToBigqueryBetaRoutineArgumentsModeEnum(e betapb.BigqueryBetaRoutineArgumentsModeEnum) *beta.RoutineArgumentsModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.BigqueryBetaRoutineArgumentsModeEnum_name[int32(e)]; ok {
		e := beta.RoutineArgumentsModeEnum(n[len("BigqueryBetaRoutineArgumentsModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToRoutineArgumentsDataTypeTypeKindEnum converts a RoutineArgumentsDataTypeTypeKindEnum enum from its proto representation.
func ProtoToBigqueryBetaRoutineArgumentsDataTypeTypeKindEnum(e betapb.BigqueryBetaRoutineArgumentsDataTypeTypeKindEnum) *beta.RoutineArgumentsDataTypeTypeKindEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.BigqueryBetaRoutineArgumentsDataTypeTypeKindEnum_name[int32(e)]; ok {
		e := beta.RoutineArgumentsDataTypeTypeKindEnum(n[len("BigqueryBetaRoutineArgumentsDataTypeTypeKindEnum"):])
		return &e
	}
	return nil
}

// ProtoToRoutineDeterminismLevelEnum converts a RoutineDeterminismLevelEnum enum from its proto representation.
func ProtoToBigqueryBetaRoutineDeterminismLevelEnum(e betapb.BigqueryBetaRoutineDeterminismLevelEnum) *beta.RoutineDeterminismLevelEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.BigqueryBetaRoutineDeterminismLevelEnum_name[int32(e)]; ok {
		e := beta.RoutineDeterminismLevelEnum(n[len("BigqueryBetaRoutineDeterminismLevelEnum"):])
		return &e
	}
	return nil
}

// ProtoToRoutineArguments converts a RoutineArguments object from its proto representation.
func ProtoToBigqueryBetaRoutineArguments(p *betapb.BigqueryBetaRoutineArguments) *beta.RoutineArguments {
	if p == nil {
		return nil
	}
	obj := &beta.RoutineArguments{
		Name:         dcl.StringOrNil(p.GetName()),
		ArgumentKind: ProtoToBigqueryBetaRoutineArgumentsArgumentKindEnum(p.GetArgumentKind()),
		Mode:         ProtoToBigqueryBetaRoutineArgumentsModeEnum(p.GetMode()),
		DataType:     ProtoToBigqueryBetaRoutineArgumentsDataType(p.GetDataType()),
	}
	return obj
}

// ProtoToRoutineArgumentsDataType converts a RoutineArgumentsDataType object from its proto representation.
func ProtoToBigqueryBetaRoutineArgumentsDataType(p *betapb.BigqueryBetaRoutineArgumentsDataType) *beta.RoutineArgumentsDataType {
	if p == nil {
		return nil
	}
	obj := &beta.RoutineArgumentsDataType{
		TypeKind:         ProtoToBigqueryBetaRoutineArgumentsDataTypeTypeKindEnum(p.GetTypeKind()),
		ArrayElementType: ProtoToBigqueryBetaRoutineArgumentsDataType(p.GetArrayElementType()),
		StructType:       ProtoToBigqueryBetaRoutineArgumentsDataTypeStructType(p.GetStructType()),
	}
	return obj
}

// ProtoToRoutineArgumentsDataTypeStructType converts a RoutineArgumentsDataTypeStructType object from its proto representation.
func ProtoToBigqueryBetaRoutineArgumentsDataTypeStructType(p *betapb.BigqueryBetaRoutineArgumentsDataTypeStructType) *beta.RoutineArgumentsDataTypeStructType {
	if p == nil {
		return nil
	}
	obj := &beta.RoutineArgumentsDataTypeStructType{}
	for _, r := range p.GetFields() {
		obj.Fields = append(obj.Fields, *ProtoToBigqueryBetaRoutineArgumentsDataTypeStructTypeFields(r))
	}
	return obj
}

// ProtoToRoutineArgumentsDataTypeStructTypeFields converts a RoutineArgumentsDataTypeStructTypeFields object from its proto representation.
func ProtoToBigqueryBetaRoutineArgumentsDataTypeStructTypeFields(p *betapb.BigqueryBetaRoutineArgumentsDataTypeStructTypeFields) *beta.RoutineArgumentsDataTypeStructTypeFields {
	if p == nil {
		return nil
	}
	obj := &beta.RoutineArgumentsDataTypeStructTypeFields{
		Name: dcl.StringOrNil(p.GetName()),
		Type: ProtoToBigqueryBetaRoutineArgumentsDataType(p.GetType()),
	}
	return obj
}

// ProtoToRoutine converts a Routine resource from its proto representation.
func ProtoToRoutine(p *betapb.BigqueryBetaRoutine) *beta.Routine {
	obj := &beta.Routine{
		Etag:             dcl.StringOrNil(p.GetEtag()),
		Name:             dcl.StringOrNil(p.GetName()),
		Project:          dcl.StringOrNil(p.GetProject()),
		Dataset:          dcl.StringOrNil(p.GetDataset()),
		RoutineType:      ProtoToBigqueryBetaRoutineRoutineTypeEnum(p.GetRoutineType()),
		CreationTime:     dcl.Int64OrNil(p.GetCreationTime()),
		LastModifiedTime: dcl.Int64OrNil(p.GetLastModifiedTime()),
		Language:         ProtoToBigqueryBetaRoutineLanguageEnum(p.GetLanguage()),
		ReturnType:       ProtoToBigqueryBetaRoutineArgumentsDataType(p.GetReturnType()),
		DefinitionBody:   dcl.StringOrNil(p.GetDefinitionBody()),
		Description:      dcl.StringOrNil(p.GetDescription()),
		DeterminismLevel: ProtoToBigqueryBetaRoutineDeterminismLevelEnum(p.GetDeterminismLevel()),
		StrictMode:       dcl.Bool(p.GetStrictMode()),
	}
	for _, r := range p.GetArguments() {
		obj.Arguments = append(obj.Arguments, *ProtoToBigqueryBetaRoutineArguments(r))
	}
	for _, r := range p.GetImportedLibraries() {
		obj.ImportedLibraries = append(obj.ImportedLibraries, r)
	}
	return obj
}

// RoutineRoutineTypeEnumToProto converts a RoutineRoutineTypeEnum enum to its proto representation.
func BigqueryBetaRoutineRoutineTypeEnumToProto(e *beta.RoutineRoutineTypeEnum) betapb.BigqueryBetaRoutineRoutineTypeEnum {
	if e == nil {
		return betapb.BigqueryBetaRoutineRoutineTypeEnum(0)
	}
	if v, ok := betapb.BigqueryBetaRoutineRoutineTypeEnum_value["RoutineRoutineTypeEnum"+string(*e)]; ok {
		return betapb.BigqueryBetaRoutineRoutineTypeEnum(v)
	}
	return betapb.BigqueryBetaRoutineRoutineTypeEnum(0)
}

// RoutineLanguageEnumToProto converts a RoutineLanguageEnum enum to its proto representation.
func BigqueryBetaRoutineLanguageEnumToProto(e *beta.RoutineLanguageEnum) betapb.BigqueryBetaRoutineLanguageEnum {
	if e == nil {
		return betapb.BigqueryBetaRoutineLanguageEnum(0)
	}
	if v, ok := betapb.BigqueryBetaRoutineLanguageEnum_value["RoutineLanguageEnum"+string(*e)]; ok {
		return betapb.BigqueryBetaRoutineLanguageEnum(v)
	}
	return betapb.BigqueryBetaRoutineLanguageEnum(0)
}

// RoutineArgumentsArgumentKindEnumToProto converts a RoutineArgumentsArgumentKindEnum enum to its proto representation.
func BigqueryBetaRoutineArgumentsArgumentKindEnumToProto(e *beta.RoutineArgumentsArgumentKindEnum) betapb.BigqueryBetaRoutineArgumentsArgumentKindEnum {
	if e == nil {
		return betapb.BigqueryBetaRoutineArgumentsArgumentKindEnum(0)
	}
	if v, ok := betapb.BigqueryBetaRoutineArgumentsArgumentKindEnum_value["RoutineArgumentsArgumentKindEnum"+string(*e)]; ok {
		return betapb.BigqueryBetaRoutineArgumentsArgumentKindEnum(v)
	}
	return betapb.BigqueryBetaRoutineArgumentsArgumentKindEnum(0)
}

// RoutineArgumentsModeEnumToProto converts a RoutineArgumentsModeEnum enum to its proto representation.
func BigqueryBetaRoutineArgumentsModeEnumToProto(e *beta.RoutineArgumentsModeEnum) betapb.BigqueryBetaRoutineArgumentsModeEnum {
	if e == nil {
		return betapb.BigqueryBetaRoutineArgumentsModeEnum(0)
	}
	if v, ok := betapb.BigqueryBetaRoutineArgumentsModeEnum_value["RoutineArgumentsModeEnum"+string(*e)]; ok {
		return betapb.BigqueryBetaRoutineArgumentsModeEnum(v)
	}
	return betapb.BigqueryBetaRoutineArgumentsModeEnum(0)
}

// RoutineArgumentsDataTypeTypeKindEnumToProto converts a RoutineArgumentsDataTypeTypeKindEnum enum to its proto representation.
func BigqueryBetaRoutineArgumentsDataTypeTypeKindEnumToProto(e *beta.RoutineArgumentsDataTypeTypeKindEnum) betapb.BigqueryBetaRoutineArgumentsDataTypeTypeKindEnum {
	if e == nil {
		return betapb.BigqueryBetaRoutineArgumentsDataTypeTypeKindEnum(0)
	}
	if v, ok := betapb.BigqueryBetaRoutineArgumentsDataTypeTypeKindEnum_value["RoutineArgumentsDataTypeTypeKindEnum"+string(*e)]; ok {
		return betapb.BigqueryBetaRoutineArgumentsDataTypeTypeKindEnum(v)
	}
	return betapb.BigqueryBetaRoutineArgumentsDataTypeTypeKindEnum(0)
}

// RoutineDeterminismLevelEnumToProto converts a RoutineDeterminismLevelEnum enum to its proto representation.
func BigqueryBetaRoutineDeterminismLevelEnumToProto(e *beta.RoutineDeterminismLevelEnum) betapb.BigqueryBetaRoutineDeterminismLevelEnum {
	if e == nil {
		return betapb.BigqueryBetaRoutineDeterminismLevelEnum(0)
	}
	if v, ok := betapb.BigqueryBetaRoutineDeterminismLevelEnum_value["RoutineDeterminismLevelEnum"+string(*e)]; ok {
		return betapb.BigqueryBetaRoutineDeterminismLevelEnum(v)
	}
	return betapb.BigqueryBetaRoutineDeterminismLevelEnum(0)
}

// RoutineArgumentsToProto converts a RoutineArguments object to its proto representation.
func BigqueryBetaRoutineArgumentsToProto(o *beta.RoutineArguments) *betapb.BigqueryBetaRoutineArguments {
	if o == nil {
		return nil
	}
	p := &betapb.BigqueryBetaRoutineArguments{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetArgumentKind(BigqueryBetaRoutineArgumentsArgumentKindEnumToProto(o.ArgumentKind))
	p.SetMode(BigqueryBetaRoutineArgumentsModeEnumToProto(o.Mode))
	p.SetDataType(BigqueryBetaRoutineArgumentsDataTypeToProto(o.DataType))
	return p
}

// RoutineArgumentsDataTypeToProto converts a RoutineArgumentsDataType object to its proto representation.
func BigqueryBetaRoutineArgumentsDataTypeToProto(o *beta.RoutineArgumentsDataType) *betapb.BigqueryBetaRoutineArgumentsDataType {
	if o == nil {
		return nil
	}
	p := &betapb.BigqueryBetaRoutineArgumentsDataType{}
	p.SetTypeKind(BigqueryBetaRoutineArgumentsDataTypeTypeKindEnumToProto(o.TypeKind))
	p.SetArrayElementType(BigqueryBetaRoutineArgumentsDataTypeToProto(o.ArrayElementType))
	p.SetStructType(BigqueryBetaRoutineArgumentsDataTypeStructTypeToProto(o.StructType))
	return p
}

// RoutineArgumentsDataTypeStructTypeToProto converts a RoutineArgumentsDataTypeStructType object to its proto representation.
func BigqueryBetaRoutineArgumentsDataTypeStructTypeToProto(o *beta.RoutineArgumentsDataTypeStructType) *betapb.BigqueryBetaRoutineArgumentsDataTypeStructType {
	if o == nil {
		return nil
	}
	p := &betapb.BigqueryBetaRoutineArgumentsDataTypeStructType{}
	sFields := make([]*betapb.BigqueryBetaRoutineArgumentsDataTypeStructTypeFields, len(o.Fields))
	for i, r := range o.Fields {
		sFields[i] = BigqueryBetaRoutineArgumentsDataTypeStructTypeFieldsToProto(&r)
	}
	p.SetFields(sFields)
	return p
}

// RoutineArgumentsDataTypeStructTypeFieldsToProto converts a RoutineArgumentsDataTypeStructTypeFields object to its proto representation.
func BigqueryBetaRoutineArgumentsDataTypeStructTypeFieldsToProto(o *beta.RoutineArgumentsDataTypeStructTypeFields) *betapb.BigqueryBetaRoutineArgumentsDataTypeStructTypeFields {
	if o == nil {
		return nil
	}
	p := &betapb.BigqueryBetaRoutineArgumentsDataTypeStructTypeFields{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetType(BigqueryBetaRoutineArgumentsDataTypeToProto(o.Type))
	return p
}

// RoutineToProto converts a Routine resource to its proto representation.
func RoutineToProto(resource *beta.Routine) *betapb.BigqueryBetaRoutine {
	p := &betapb.BigqueryBetaRoutine{}
	p.SetEtag(dcl.ValueOrEmptyString(resource.Etag))
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetDataset(dcl.ValueOrEmptyString(resource.Dataset))
	p.SetRoutineType(BigqueryBetaRoutineRoutineTypeEnumToProto(resource.RoutineType))
	p.SetCreationTime(dcl.ValueOrEmptyInt64(resource.CreationTime))
	p.SetLastModifiedTime(dcl.ValueOrEmptyInt64(resource.LastModifiedTime))
	p.SetLanguage(BigqueryBetaRoutineLanguageEnumToProto(resource.Language))
	p.SetReturnType(BigqueryBetaRoutineArgumentsDataTypeToProto(resource.ReturnType))
	p.SetDefinitionBody(dcl.ValueOrEmptyString(resource.DefinitionBody))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetDeterminismLevel(BigqueryBetaRoutineDeterminismLevelEnumToProto(resource.DeterminismLevel))
	p.SetStrictMode(dcl.ValueOrEmptyBool(resource.StrictMode))
	sArguments := make([]*betapb.BigqueryBetaRoutineArguments, len(resource.Arguments))
	for i, r := range resource.Arguments {
		sArguments[i] = BigqueryBetaRoutineArgumentsToProto(&r)
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
func (s *RoutineServer) applyRoutine(ctx context.Context, c *beta.Client, request *betapb.ApplyBigqueryBetaRoutineRequest) (*betapb.BigqueryBetaRoutine, error) {
	p := ProtoToRoutine(request.GetResource())
	res, err := c.ApplyRoutine(ctx, p)
	if err != nil {
		return nil, err
	}
	r := RoutineToProto(res)
	return r, nil
}

// applyBigqueryBetaRoutine handles the gRPC request by passing it to the underlying Routine Apply() method.
func (s *RoutineServer) ApplyBigqueryBetaRoutine(ctx context.Context, request *betapb.ApplyBigqueryBetaRoutineRequest) (*betapb.BigqueryBetaRoutine, error) {
	cl, err := createConfigRoutine(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyRoutine(ctx, cl, request)
}

// DeleteRoutine handles the gRPC request by passing it to the underlying Routine Delete() method.
func (s *RoutineServer) DeleteBigqueryBetaRoutine(ctx context.Context, request *betapb.DeleteBigqueryBetaRoutineRequest) (*emptypb.Empty, error) {

	cl, err := createConfigRoutine(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteRoutine(ctx, ProtoToRoutine(request.GetResource()))

}

// ListBigqueryBetaRoutine handles the gRPC request by passing it to the underlying RoutineList() method.
func (s *RoutineServer) ListBigqueryBetaRoutine(ctx context.Context, request *betapb.ListBigqueryBetaRoutineRequest) (*betapb.ListBigqueryBetaRoutineResponse, error) {
	cl, err := createConfigRoutine(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListRoutine(ctx, request.GetProject(), request.GetDataset())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.BigqueryBetaRoutine
	for _, r := range resources.Items {
		rp := RoutineToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListBigqueryBetaRoutineResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigRoutine(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
