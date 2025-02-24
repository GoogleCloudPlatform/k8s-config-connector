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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/bigquery/alpha/bigquery_alpha_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/bigquery/alpha"
)

// RoutineServer implements the gRPC interface for Routine.
type RoutineServer struct{}

// ProtoToRoutineRoutineTypeEnum converts a RoutineRoutineTypeEnum enum from its proto representation.
func ProtoToBigqueryAlphaRoutineRoutineTypeEnum(e alphapb.BigqueryAlphaRoutineRoutineTypeEnum) *alpha.RoutineRoutineTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.BigqueryAlphaRoutineRoutineTypeEnum_name[int32(e)]; ok {
		e := alpha.RoutineRoutineTypeEnum(n[len("BigqueryAlphaRoutineRoutineTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToRoutineLanguageEnum converts a RoutineLanguageEnum enum from its proto representation.
func ProtoToBigqueryAlphaRoutineLanguageEnum(e alphapb.BigqueryAlphaRoutineLanguageEnum) *alpha.RoutineLanguageEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.BigqueryAlphaRoutineLanguageEnum_name[int32(e)]; ok {
		e := alpha.RoutineLanguageEnum(n[len("BigqueryAlphaRoutineLanguageEnum"):])
		return &e
	}
	return nil
}

// ProtoToRoutineArgumentsArgumentKindEnum converts a RoutineArgumentsArgumentKindEnum enum from its proto representation.
func ProtoToBigqueryAlphaRoutineArgumentsArgumentKindEnum(e alphapb.BigqueryAlphaRoutineArgumentsArgumentKindEnum) *alpha.RoutineArgumentsArgumentKindEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.BigqueryAlphaRoutineArgumentsArgumentKindEnum_name[int32(e)]; ok {
		e := alpha.RoutineArgumentsArgumentKindEnum(n[len("BigqueryAlphaRoutineArgumentsArgumentKindEnum"):])
		return &e
	}
	return nil
}

// ProtoToRoutineArgumentsModeEnum converts a RoutineArgumentsModeEnum enum from its proto representation.
func ProtoToBigqueryAlphaRoutineArgumentsModeEnum(e alphapb.BigqueryAlphaRoutineArgumentsModeEnum) *alpha.RoutineArgumentsModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.BigqueryAlphaRoutineArgumentsModeEnum_name[int32(e)]; ok {
		e := alpha.RoutineArgumentsModeEnum(n[len("BigqueryAlphaRoutineArgumentsModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToRoutineArgumentsDataTypeTypeKindEnum converts a RoutineArgumentsDataTypeTypeKindEnum enum from its proto representation.
func ProtoToBigqueryAlphaRoutineArgumentsDataTypeTypeKindEnum(e alphapb.BigqueryAlphaRoutineArgumentsDataTypeTypeKindEnum) *alpha.RoutineArgumentsDataTypeTypeKindEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.BigqueryAlphaRoutineArgumentsDataTypeTypeKindEnum_name[int32(e)]; ok {
		e := alpha.RoutineArgumentsDataTypeTypeKindEnum(n[len("BigqueryAlphaRoutineArgumentsDataTypeTypeKindEnum"):])
		return &e
	}
	return nil
}

// ProtoToRoutineDeterminismLevelEnum converts a RoutineDeterminismLevelEnum enum from its proto representation.
func ProtoToBigqueryAlphaRoutineDeterminismLevelEnum(e alphapb.BigqueryAlphaRoutineDeterminismLevelEnum) *alpha.RoutineDeterminismLevelEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.BigqueryAlphaRoutineDeterminismLevelEnum_name[int32(e)]; ok {
		e := alpha.RoutineDeterminismLevelEnum(n[len("BigqueryAlphaRoutineDeterminismLevelEnum"):])
		return &e
	}
	return nil
}

// ProtoToRoutineArguments converts a RoutineArguments object from its proto representation.
func ProtoToBigqueryAlphaRoutineArguments(p *alphapb.BigqueryAlphaRoutineArguments) *alpha.RoutineArguments {
	if p == nil {
		return nil
	}
	obj := &alpha.RoutineArguments{
		Name:         dcl.StringOrNil(p.GetName()),
		ArgumentKind: ProtoToBigqueryAlphaRoutineArgumentsArgumentKindEnum(p.GetArgumentKind()),
		Mode:         ProtoToBigqueryAlphaRoutineArgumentsModeEnum(p.GetMode()),
		DataType:     ProtoToBigqueryAlphaRoutineArgumentsDataType(p.GetDataType()),
	}
	return obj
}

// ProtoToRoutineArgumentsDataType converts a RoutineArgumentsDataType object from its proto representation.
func ProtoToBigqueryAlphaRoutineArgumentsDataType(p *alphapb.BigqueryAlphaRoutineArgumentsDataType) *alpha.RoutineArgumentsDataType {
	if p == nil {
		return nil
	}
	obj := &alpha.RoutineArgumentsDataType{
		TypeKind:         ProtoToBigqueryAlphaRoutineArgumentsDataTypeTypeKindEnum(p.GetTypeKind()),
		ArrayElementType: ProtoToBigqueryAlphaRoutineArgumentsDataType(p.GetArrayElementType()),
		StructType:       ProtoToBigqueryAlphaRoutineArgumentsDataTypeStructType(p.GetStructType()),
	}
	return obj
}

// ProtoToRoutineArgumentsDataTypeStructType converts a RoutineArgumentsDataTypeStructType object from its proto representation.
func ProtoToBigqueryAlphaRoutineArgumentsDataTypeStructType(p *alphapb.BigqueryAlphaRoutineArgumentsDataTypeStructType) *alpha.RoutineArgumentsDataTypeStructType {
	if p == nil {
		return nil
	}
	obj := &alpha.RoutineArgumentsDataTypeStructType{}
	for _, r := range p.GetFields() {
		obj.Fields = append(obj.Fields, *ProtoToBigqueryAlphaRoutineArgumentsDataTypeStructTypeFields(r))
	}
	return obj
}

// ProtoToRoutineArgumentsDataTypeStructTypeFields converts a RoutineArgumentsDataTypeStructTypeFields object from its proto representation.
func ProtoToBigqueryAlphaRoutineArgumentsDataTypeStructTypeFields(p *alphapb.BigqueryAlphaRoutineArgumentsDataTypeStructTypeFields) *alpha.RoutineArgumentsDataTypeStructTypeFields {
	if p == nil {
		return nil
	}
	obj := &alpha.RoutineArgumentsDataTypeStructTypeFields{
		Name: dcl.StringOrNil(p.GetName()),
		Type: ProtoToBigqueryAlphaRoutineArgumentsDataType(p.GetType()),
	}
	return obj
}

// ProtoToRoutine converts a Routine resource from its proto representation.
func ProtoToRoutine(p *alphapb.BigqueryAlphaRoutine) *alpha.Routine {
	obj := &alpha.Routine{
		Etag:             dcl.StringOrNil(p.GetEtag()),
		Name:             dcl.StringOrNil(p.GetName()),
		Project:          dcl.StringOrNil(p.GetProject()),
		Dataset:          dcl.StringOrNil(p.GetDataset()),
		RoutineType:      ProtoToBigqueryAlphaRoutineRoutineTypeEnum(p.GetRoutineType()),
		CreationTime:     dcl.Int64OrNil(p.GetCreationTime()),
		LastModifiedTime: dcl.Int64OrNil(p.GetLastModifiedTime()),
		Language:         ProtoToBigqueryAlphaRoutineLanguageEnum(p.GetLanguage()),
		ReturnType:       ProtoToBigqueryAlphaRoutineArgumentsDataType(p.GetReturnType()),
		DefinitionBody:   dcl.StringOrNil(p.GetDefinitionBody()),
		Description:      dcl.StringOrNil(p.GetDescription()),
		DeterminismLevel: ProtoToBigqueryAlphaRoutineDeterminismLevelEnum(p.GetDeterminismLevel()),
		StrictMode:       dcl.Bool(p.GetStrictMode()),
	}
	for _, r := range p.GetArguments() {
		obj.Arguments = append(obj.Arguments, *ProtoToBigqueryAlphaRoutineArguments(r))
	}
	for _, r := range p.GetImportedLibraries() {
		obj.ImportedLibraries = append(obj.ImportedLibraries, r)
	}
	return obj
}

// RoutineRoutineTypeEnumToProto converts a RoutineRoutineTypeEnum enum to its proto representation.
func BigqueryAlphaRoutineRoutineTypeEnumToProto(e *alpha.RoutineRoutineTypeEnum) alphapb.BigqueryAlphaRoutineRoutineTypeEnum {
	if e == nil {
		return alphapb.BigqueryAlphaRoutineRoutineTypeEnum(0)
	}
	if v, ok := alphapb.BigqueryAlphaRoutineRoutineTypeEnum_value["RoutineRoutineTypeEnum"+string(*e)]; ok {
		return alphapb.BigqueryAlphaRoutineRoutineTypeEnum(v)
	}
	return alphapb.BigqueryAlphaRoutineRoutineTypeEnum(0)
}

// RoutineLanguageEnumToProto converts a RoutineLanguageEnum enum to its proto representation.
func BigqueryAlphaRoutineLanguageEnumToProto(e *alpha.RoutineLanguageEnum) alphapb.BigqueryAlphaRoutineLanguageEnum {
	if e == nil {
		return alphapb.BigqueryAlphaRoutineLanguageEnum(0)
	}
	if v, ok := alphapb.BigqueryAlphaRoutineLanguageEnum_value["RoutineLanguageEnum"+string(*e)]; ok {
		return alphapb.BigqueryAlphaRoutineLanguageEnum(v)
	}
	return alphapb.BigqueryAlphaRoutineLanguageEnum(0)
}

// RoutineArgumentsArgumentKindEnumToProto converts a RoutineArgumentsArgumentKindEnum enum to its proto representation.
func BigqueryAlphaRoutineArgumentsArgumentKindEnumToProto(e *alpha.RoutineArgumentsArgumentKindEnum) alphapb.BigqueryAlphaRoutineArgumentsArgumentKindEnum {
	if e == nil {
		return alphapb.BigqueryAlphaRoutineArgumentsArgumentKindEnum(0)
	}
	if v, ok := alphapb.BigqueryAlphaRoutineArgumentsArgumentKindEnum_value["RoutineArgumentsArgumentKindEnum"+string(*e)]; ok {
		return alphapb.BigqueryAlphaRoutineArgumentsArgumentKindEnum(v)
	}
	return alphapb.BigqueryAlphaRoutineArgumentsArgumentKindEnum(0)
}

// RoutineArgumentsModeEnumToProto converts a RoutineArgumentsModeEnum enum to its proto representation.
func BigqueryAlphaRoutineArgumentsModeEnumToProto(e *alpha.RoutineArgumentsModeEnum) alphapb.BigqueryAlphaRoutineArgumentsModeEnum {
	if e == nil {
		return alphapb.BigqueryAlphaRoutineArgumentsModeEnum(0)
	}
	if v, ok := alphapb.BigqueryAlphaRoutineArgumentsModeEnum_value["RoutineArgumentsModeEnum"+string(*e)]; ok {
		return alphapb.BigqueryAlphaRoutineArgumentsModeEnum(v)
	}
	return alphapb.BigqueryAlphaRoutineArgumentsModeEnum(0)
}

// RoutineArgumentsDataTypeTypeKindEnumToProto converts a RoutineArgumentsDataTypeTypeKindEnum enum to its proto representation.
func BigqueryAlphaRoutineArgumentsDataTypeTypeKindEnumToProto(e *alpha.RoutineArgumentsDataTypeTypeKindEnum) alphapb.BigqueryAlphaRoutineArgumentsDataTypeTypeKindEnum {
	if e == nil {
		return alphapb.BigqueryAlphaRoutineArgumentsDataTypeTypeKindEnum(0)
	}
	if v, ok := alphapb.BigqueryAlphaRoutineArgumentsDataTypeTypeKindEnum_value["RoutineArgumentsDataTypeTypeKindEnum"+string(*e)]; ok {
		return alphapb.BigqueryAlphaRoutineArgumentsDataTypeTypeKindEnum(v)
	}
	return alphapb.BigqueryAlphaRoutineArgumentsDataTypeTypeKindEnum(0)
}

// RoutineDeterminismLevelEnumToProto converts a RoutineDeterminismLevelEnum enum to its proto representation.
func BigqueryAlphaRoutineDeterminismLevelEnumToProto(e *alpha.RoutineDeterminismLevelEnum) alphapb.BigqueryAlphaRoutineDeterminismLevelEnum {
	if e == nil {
		return alphapb.BigqueryAlphaRoutineDeterminismLevelEnum(0)
	}
	if v, ok := alphapb.BigqueryAlphaRoutineDeterminismLevelEnum_value["RoutineDeterminismLevelEnum"+string(*e)]; ok {
		return alphapb.BigqueryAlphaRoutineDeterminismLevelEnum(v)
	}
	return alphapb.BigqueryAlphaRoutineDeterminismLevelEnum(0)
}

// RoutineArgumentsToProto converts a RoutineArguments object to its proto representation.
func BigqueryAlphaRoutineArgumentsToProto(o *alpha.RoutineArguments) *alphapb.BigqueryAlphaRoutineArguments {
	if o == nil {
		return nil
	}
	p := &alphapb.BigqueryAlphaRoutineArguments{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetArgumentKind(BigqueryAlphaRoutineArgumentsArgumentKindEnumToProto(o.ArgumentKind))
	p.SetMode(BigqueryAlphaRoutineArgumentsModeEnumToProto(o.Mode))
	p.SetDataType(BigqueryAlphaRoutineArgumentsDataTypeToProto(o.DataType))
	return p
}

// RoutineArgumentsDataTypeToProto converts a RoutineArgumentsDataType object to its proto representation.
func BigqueryAlphaRoutineArgumentsDataTypeToProto(o *alpha.RoutineArgumentsDataType) *alphapb.BigqueryAlphaRoutineArgumentsDataType {
	if o == nil {
		return nil
	}
	p := &alphapb.BigqueryAlphaRoutineArgumentsDataType{}
	p.SetTypeKind(BigqueryAlphaRoutineArgumentsDataTypeTypeKindEnumToProto(o.TypeKind))
	p.SetArrayElementType(BigqueryAlphaRoutineArgumentsDataTypeToProto(o.ArrayElementType))
	p.SetStructType(BigqueryAlphaRoutineArgumentsDataTypeStructTypeToProto(o.StructType))
	return p
}

// RoutineArgumentsDataTypeStructTypeToProto converts a RoutineArgumentsDataTypeStructType object to its proto representation.
func BigqueryAlphaRoutineArgumentsDataTypeStructTypeToProto(o *alpha.RoutineArgumentsDataTypeStructType) *alphapb.BigqueryAlphaRoutineArgumentsDataTypeStructType {
	if o == nil {
		return nil
	}
	p := &alphapb.BigqueryAlphaRoutineArgumentsDataTypeStructType{}
	sFields := make([]*alphapb.BigqueryAlphaRoutineArgumentsDataTypeStructTypeFields, len(o.Fields))
	for i, r := range o.Fields {
		sFields[i] = BigqueryAlphaRoutineArgumentsDataTypeStructTypeFieldsToProto(&r)
	}
	p.SetFields(sFields)
	return p
}

// RoutineArgumentsDataTypeStructTypeFieldsToProto converts a RoutineArgumentsDataTypeStructTypeFields object to its proto representation.
func BigqueryAlphaRoutineArgumentsDataTypeStructTypeFieldsToProto(o *alpha.RoutineArgumentsDataTypeStructTypeFields) *alphapb.BigqueryAlphaRoutineArgumentsDataTypeStructTypeFields {
	if o == nil {
		return nil
	}
	p := &alphapb.BigqueryAlphaRoutineArgumentsDataTypeStructTypeFields{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetType(BigqueryAlphaRoutineArgumentsDataTypeToProto(o.Type))
	return p
}

// RoutineToProto converts a Routine resource to its proto representation.
func RoutineToProto(resource *alpha.Routine) *alphapb.BigqueryAlphaRoutine {
	p := &alphapb.BigqueryAlphaRoutine{}
	p.SetEtag(dcl.ValueOrEmptyString(resource.Etag))
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetDataset(dcl.ValueOrEmptyString(resource.Dataset))
	p.SetRoutineType(BigqueryAlphaRoutineRoutineTypeEnumToProto(resource.RoutineType))
	p.SetCreationTime(dcl.ValueOrEmptyInt64(resource.CreationTime))
	p.SetLastModifiedTime(dcl.ValueOrEmptyInt64(resource.LastModifiedTime))
	p.SetLanguage(BigqueryAlphaRoutineLanguageEnumToProto(resource.Language))
	p.SetReturnType(BigqueryAlphaRoutineArgumentsDataTypeToProto(resource.ReturnType))
	p.SetDefinitionBody(dcl.ValueOrEmptyString(resource.DefinitionBody))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetDeterminismLevel(BigqueryAlphaRoutineDeterminismLevelEnumToProto(resource.DeterminismLevel))
	p.SetStrictMode(dcl.ValueOrEmptyBool(resource.StrictMode))
	sArguments := make([]*alphapb.BigqueryAlphaRoutineArguments, len(resource.Arguments))
	for i, r := range resource.Arguments {
		sArguments[i] = BigqueryAlphaRoutineArgumentsToProto(&r)
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
func (s *RoutineServer) applyRoutine(ctx context.Context, c *alpha.Client, request *alphapb.ApplyBigqueryAlphaRoutineRequest) (*alphapb.BigqueryAlphaRoutine, error) {
	p := ProtoToRoutine(request.GetResource())
	res, err := c.ApplyRoutine(ctx, p)
	if err != nil {
		return nil, err
	}
	r := RoutineToProto(res)
	return r, nil
}

// applyBigqueryAlphaRoutine handles the gRPC request by passing it to the underlying Routine Apply() method.
func (s *RoutineServer) ApplyBigqueryAlphaRoutine(ctx context.Context, request *alphapb.ApplyBigqueryAlphaRoutineRequest) (*alphapb.BigqueryAlphaRoutine, error) {
	cl, err := createConfigRoutine(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyRoutine(ctx, cl, request)
}

// DeleteRoutine handles the gRPC request by passing it to the underlying Routine Delete() method.
func (s *RoutineServer) DeleteBigqueryAlphaRoutine(ctx context.Context, request *alphapb.DeleteBigqueryAlphaRoutineRequest) (*emptypb.Empty, error) {

	cl, err := createConfigRoutine(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteRoutine(ctx, ProtoToRoutine(request.GetResource()))

}

// ListBigqueryAlphaRoutine handles the gRPC request by passing it to the underlying RoutineList() method.
func (s *RoutineServer) ListBigqueryAlphaRoutine(ctx context.Context, request *alphapb.ListBigqueryAlphaRoutineRequest) (*alphapb.ListBigqueryAlphaRoutineResponse, error) {
	cl, err := createConfigRoutine(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListRoutine(ctx, request.GetProject(), request.GetDataset())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.BigqueryAlphaRoutine
	for _, r := range resources.Items {
		rp := RoutineToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListBigqueryAlphaRoutineResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigRoutine(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
