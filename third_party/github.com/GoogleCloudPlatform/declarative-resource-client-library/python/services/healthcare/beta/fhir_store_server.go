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
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/healthcare/beta/healthcare_beta_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/healthcare/beta"
)

// FhirStoreServer implements the gRPC interface for FhirStore.
type FhirStoreServer struct{}

// ProtoToFhirStoreVersionEnum converts a FhirStoreVersionEnum enum from its proto representation.
func ProtoToHealthcareBetaFhirStoreVersionEnum(e betapb.HealthcareBetaFhirStoreVersionEnum) *beta.FhirStoreVersionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.HealthcareBetaFhirStoreVersionEnum_name[int32(e)]; ok {
		e := beta.FhirStoreVersionEnum(n[len("HealthcareBetaFhirStoreVersionEnum"):])
		return &e
	}
	return nil
}

// ProtoToFhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnum converts a FhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnum enum from its proto representation.
func ProtoToHealthcareBetaFhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnum(e betapb.HealthcareBetaFhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnum) *beta.FhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.HealthcareBetaFhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnum_name[int32(e)]; ok {
		e := beta.FhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnum(n[len("HealthcareBetaFhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToFhirStoreComplexDataTypeReferenceParsingEnum converts a FhirStoreComplexDataTypeReferenceParsingEnum enum from its proto representation.
func ProtoToHealthcareBetaFhirStoreComplexDataTypeReferenceParsingEnum(e betapb.HealthcareBetaFhirStoreComplexDataTypeReferenceParsingEnum) *beta.FhirStoreComplexDataTypeReferenceParsingEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.HealthcareBetaFhirStoreComplexDataTypeReferenceParsingEnum_name[int32(e)]; ok {
		e := beta.FhirStoreComplexDataTypeReferenceParsingEnum(n[len("HealthcareBetaFhirStoreComplexDataTypeReferenceParsingEnum"):])
		return &e
	}
	return nil
}

// ProtoToFhirStoreNotificationConfig converts a FhirStoreNotificationConfig object from its proto representation.
func ProtoToHealthcareBetaFhirStoreNotificationConfig(p *betapb.HealthcareBetaFhirStoreNotificationConfig) *beta.FhirStoreNotificationConfig {
	if p == nil {
		return nil
	}
	obj := &beta.FhirStoreNotificationConfig{
		PubsubTopic: dcl.StringOrNil(p.GetPubsubTopic()),
	}
	return obj
}

// ProtoToFhirStoreStreamConfigs converts a FhirStoreStreamConfigs object from its proto representation.
func ProtoToHealthcareBetaFhirStoreStreamConfigs(p *betapb.HealthcareBetaFhirStoreStreamConfigs) *beta.FhirStoreStreamConfigs {
	if p == nil {
		return nil
	}
	obj := &beta.FhirStoreStreamConfigs{
		BigqueryDestination: ProtoToHealthcareBetaFhirStoreStreamConfigsBigqueryDestination(p.GetBigqueryDestination()),
	}
	for _, r := range p.GetResourceTypes() {
		obj.ResourceTypes = append(obj.ResourceTypes, r)
	}
	return obj
}

// ProtoToFhirStoreStreamConfigsBigqueryDestination converts a FhirStoreStreamConfigsBigqueryDestination object from its proto representation.
func ProtoToHealthcareBetaFhirStoreStreamConfigsBigqueryDestination(p *betapb.HealthcareBetaFhirStoreStreamConfigsBigqueryDestination) *beta.FhirStoreStreamConfigsBigqueryDestination {
	if p == nil {
		return nil
	}
	obj := &beta.FhirStoreStreamConfigsBigqueryDestination{
		DatasetUri:   dcl.StringOrNil(p.GetDatasetUri()),
		SchemaConfig: ProtoToHealthcareBetaFhirStoreStreamConfigsBigqueryDestinationSchemaConfig(p.GetSchemaConfig()),
	}
	return obj
}

// ProtoToFhirStoreStreamConfigsBigqueryDestinationSchemaConfig converts a FhirStoreStreamConfigsBigqueryDestinationSchemaConfig object from its proto representation.
func ProtoToHealthcareBetaFhirStoreStreamConfigsBigqueryDestinationSchemaConfig(p *betapb.HealthcareBetaFhirStoreStreamConfigsBigqueryDestinationSchemaConfig) *beta.FhirStoreStreamConfigsBigqueryDestinationSchemaConfig {
	if p == nil {
		return nil
	}
	obj := &beta.FhirStoreStreamConfigsBigqueryDestinationSchemaConfig{
		SchemaType:              ProtoToHealthcareBetaFhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnum(p.GetSchemaType()),
		RecursiveStructureDepth: dcl.Int64OrNil(p.GetRecursiveStructureDepth()),
	}
	return obj
}

// ProtoToFhirStoreValidationConfig converts a FhirStoreValidationConfig object from its proto representation.
func ProtoToHealthcareBetaFhirStoreValidationConfig(p *betapb.HealthcareBetaFhirStoreValidationConfig) *beta.FhirStoreValidationConfig {
	if p == nil {
		return nil
	}
	obj := &beta.FhirStoreValidationConfig{
		DisableProfileValidation:       dcl.Bool(p.GetDisableProfileValidation()),
		DisableRequiredFieldValidation: dcl.Bool(p.GetDisableRequiredFieldValidation()),
		DisableReferenceTypeValidation: dcl.Bool(p.GetDisableReferenceTypeValidation()),
		DisableFhirpathValidation:      dcl.Bool(p.GetDisableFhirpathValidation()),
	}
	for _, r := range p.GetEnabledImplementationGuides() {
		obj.EnabledImplementationGuides = append(obj.EnabledImplementationGuides, r)
	}
	return obj
}

// ProtoToFhirStore converts a FhirStore resource from its proto representation.
func ProtoToFhirStore(p *betapb.HealthcareBetaFhirStore) *beta.FhirStore {
	obj := &beta.FhirStore{
		Name:                            dcl.StringOrNil(p.GetName()),
		EnableUpdateCreate:              dcl.Bool(p.GetEnableUpdateCreate()),
		NotificationConfig:              ProtoToHealthcareBetaFhirStoreNotificationConfig(p.GetNotificationConfig()),
		DisableReferentialIntegrity:     dcl.Bool(p.GetDisableReferentialIntegrity()),
		ShardNum:                        dcl.Int64OrNil(p.GetShardNum()),
		DisableResourceVersioning:       dcl.Bool(p.GetDisableResourceVersioning()),
		Version:                         ProtoToHealthcareBetaFhirStoreVersionEnum(p.GetVersion()),
		ValidationConfig:                ProtoToHealthcareBetaFhirStoreValidationConfig(p.GetValidationConfig()),
		DefaultSearchHandlingStrict:     dcl.Bool(p.GetDefaultSearchHandlingStrict()),
		ComplexDataTypeReferenceParsing: ProtoToHealthcareBetaFhirStoreComplexDataTypeReferenceParsingEnum(p.GetComplexDataTypeReferenceParsing()),
		Project:                         dcl.StringOrNil(p.GetProject()),
		Location:                        dcl.StringOrNil(p.GetLocation()),
		Dataset:                         dcl.StringOrNil(p.GetDataset()),
	}
	for _, r := range p.GetStreamConfigs() {
		obj.StreamConfigs = append(obj.StreamConfigs, *ProtoToHealthcareBetaFhirStoreStreamConfigs(r))
	}
	return obj
}

// FhirStoreVersionEnumToProto converts a FhirStoreVersionEnum enum to its proto representation.
func HealthcareBetaFhirStoreVersionEnumToProto(e *beta.FhirStoreVersionEnum) betapb.HealthcareBetaFhirStoreVersionEnum {
	if e == nil {
		return betapb.HealthcareBetaFhirStoreVersionEnum(0)
	}
	if v, ok := betapb.HealthcareBetaFhirStoreVersionEnum_value["FhirStoreVersionEnum"+string(*e)]; ok {
		return betapb.HealthcareBetaFhirStoreVersionEnum(v)
	}
	return betapb.HealthcareBetaFhirStoreVersionEnum(0)
}

// FhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnumToProto converts a FhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnum enum to its proto representation.
func HealthcareBetaFhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnumToProto(e *beta.FhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnum) betapb.HealthcareBetaFhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnum {
	if e == nil {
		return betapb.HealthcareBetaFhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnum(0)
	}
	if v, ok := betapb.HealthcareBetaFhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnum_value["FhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnum"+string(*e)]; ok {
		return betapb.HealthcareBetaFhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnum(v)
	}
	return betapb.HealthcareBetaFhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnum(0)
}

// FhirStoreComplexDataTypeReferenceParsingEnumToProto converts a FhirStoreComplexDataTypeReferenceParsingEnum enum to its proto representation.
func HealthcareBetaFhirStoreComplexDataTypeReferenceParsingEnumToProto(e *beta.FhirStoreComplexDataTypeReferenceParsingEnum) betapb.HealthcareBetaFhirStoreComplexDataTypeReferenceParsingEnum {
	if e == nil {
		return betapb.HealthcareBetaFhirStoreComplexDataTypeReferenceParsingEnum(0)
	}
	if v, ok := betapb.HealthcareBetaFhirStoreComplexDataTypeReferenceParsingEnum_value["FhirStoreComplexDataTypeReferenceParsingEnum"+string(*e)]; ok {
		return betapb.HealthcareBetaFhirStoreComplexDataTypeReferenceParsingEnum(v)
	}
	return betapb.HealthcareBetaFhirStoreComplexDataTypeReferenceParsingEnum(0)
}

// FhirStoreNotificationConfigToProto converts a FhirStoreNotificationConfig object to its proto representation.
func HealthcareBetaFhirStoreNotificationConfigToProto(o *beta.FhirStoreNotificationConfig) *betapb.HealthcareBetaFhirStoreNotificationConfig {
	if o == nil {
		return nil
	}
	p := &betapb.HealthcareBetaFhirStoreNotificationConfig{}
	p.SetPubsubTopic(dcl.ValueOrEmptyString(o.PubsubTopic))
	return p
}

// FhirStoreStreamConfigsToProto converts a FhirStoreStreamConfigs object to its proto representation.
func HealthcareBetaFhirStoreStreamConfigsToProto(o *beta.FhirStoreStreamConfigs) *betapb.HealthcareBetaFhirStoreStreamConfigs {
	if o == nil {
		return nil
	}
	p := &betapb.HealthcareBetaFhirStoreStreamConfigs{}
	p.SetBigqueryDestination(HealthcareBetaFhirStoreStreamConfigsBigqueryDestinationToProto(o.BigqueryDestination))
	sResourceTypes := make([]string, len(o.ResourceTypes))
	for i, r := range o.ResourceTypes {
		sResourceTypes[i] = r
	}
	p.SetResourceTypes(sResourceTypes)
	return p
}

// FhirStoreStreamConfigsBigqueryDestinationToProto converts a FhirStoreStreamConfigsBigqueryDestination object to its proto representation.
func HealthcareBetaFhirStoreStreamConfigsBigqueryDestinationToProto(o *beta.FhirStoreStreamConfigsBigqueryDestination) *betapb.HealthcareBetaFhirStoreStreamConfigsBigqueryDestination {
	if o == nil {
		return nil
	}
	p := &betapb.HealthcareBetaFhirStoreStreamConfigsBigqueryDestination{}
	p.SetDatasetUri(dcl.ValueOrEmptyString(o.DatasetUri))
	p.SetSchemaConfig(HealthcareBetaFhirStoreStreamConfigsBigqueryDestinationSchemaConfigToProto(o.SchemaConfig))
	return p
}

// FhirStoreStreamConfigsBigqueryDestinationSchemaConfigToProto converts a FhirStoreStreamConfigsBigqueryDestinationSchemaConfig object to its proto representation.
func HealthcareBetaFhirStoreStreamConfigsBigqueryDestinationSchemaConfigToProto(o *beta.FhirStoreStreamConfigsBigqueryDestinationSchemaConfig) *betapb.HealthcareBetaFhirStoreStreamConfigsBigqueryDestinationSchemaConfig {
	if o == nil {
		return nil
	}
	p := &betapb.HealthcareBetaFhirStoreStreamConfigsBigqueryDestinationSchemaConfig{}
	p.SetSchemaType(HealthcareBetaFhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnumToProto(o.SchemaType))
	p.SetRecursiveStructureDepth(dcl.ValueOrEmptyInt64(o.RecursiveStructureDepth))
	return p
}

// FhirStoreValidationConfigToProto converts a FhirStoreValidationConfig object to its proto representation.
func HealthcareBetaFhirStoreValidationConfigToProto(o *beta.FhirStoreValidationConfig) *betapb.HealthcareBetaFhirStoreValidationConfig {
	if o == nil {
		return nil
	}
	p := &betapb.HealthcareBetaFhirStoreValidationConfig{}
	p.SetDisableProfileValidation(dcl.ValueOrEmptyBool(o.DisableProfileValidation))
	p.SetDisableRequiredFieldValidation(dcl.ValueOrEmptyBool(o.DisableRequiredFieldValidation))
	p.SetDisableReferenceTypeValidation(dcl.ValueOrEmptyBool(o.DisableReferenceTypeValidation))
	p.SetDisableFhirpathValidation(dcl.ValueOrEmptyBool(o.DisableFhirpathValidation))
	sEnabledImplementationGuides := make([]string, len(o.EnabledImplementationGuides))
	for i, r := range o.EnabledImplementationGuides {
		sEnabledImplementationGuides[i] = r
	}
	p.SetEnabledImplementationGuides(sEnabledImplementationGuides)
	return p
}

// FhirStoreToProto converts a FhirStore resource to its proto representation.
func FhirStoreToProto(resource *beta.FhirStore) *betapb.HealthcareBetaFhirStore {
	p := &betapb.HealthcareBetaFhirStore{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetEnableUpdateCreate(dcl.ValueOrEmptyBool(resource.EnableUpdateCreate))
	p.SetNotificationConfig(HealthcareBetaFhirStoreNotificationConfigToProto(resource.NotificationConfig))
	p.SetDisableReferentialIntegrity(dcl.ValueOrEmptyBool(resource.DisableReferentialIntegrity))
	p.SetShardNum(dcl.ValueOrEmptyInt64(resource.ShardNum))
	p.SetDisableResourceVersioning(dcl.ValueOrEmptyBool(resource.DisableResourceVersioning))
	p.SetVersion(HealthcareBetaFhirStoreVersionEnumToProto(resource.Version))
	p.SetValidationConfig(HealthcareBetaFhirStoreValidationConfigToProto(resource.ValidationConfig))
	p.SetDefaultSearchHandlingStrict(dcl.ValueOrEmptyBool(resource.DefaultSearchHandlingStrict))
	p.SetComplexDataTypeReferenceParsing(HealthcareBetaFhirStoreComplexDataTypeReferenceParsingEnumToProto(resource.ComplexDataTypeReferenceParsing))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	p.SetDataset(dcl.ValueOrEmptyString(resource.Dataset))
	mLabels := make(map[string]string, len(resource.Labels))
	for k, r := range resource.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)
	sStreamConfigs := make([]*betapb.HealthcareBetaFhirStoreStreamConfigs, len(resource.StreamConfigs))
	for i, r := range resource.StreamConfigs {
		sStreamConfigs[i] = HealthcareBetaFhirStoreStreamConfigsToProto(&r)
	}
	p.SetStreamConfigs(sStreamConfigs)

	return p
}

// applyFhirStore handles the gRPC request by passing it to the underlying FhirStore Apply() method.
func (s *FhirStoreServer) applyFhirStore(ctx context.Context, c *beta.Client, request *betapb.ApplyHealthcareBetaFhirStoreRequest) (*betapb.HealthcareBetaFhirStore, error) {
	p := ProtoToFhirStore(request.GetResource())
	res, err := c.ApplyFhirStore(ctx, p)
	if err != nil {
		return nil, err
	}
	r := FhirStoreToProto(res)
	return r, nil
}

// applyHealthcareBetaFhirStore handles the gRPC request by passing it to the underlying FhirStore Apply() method.
func (s *FhirStoreServer) ApplyHealthcareBetaFhirStore(ctx context.Context, request *betapb.ApplyHealthcareBetaFhirStoreRequest) (*betapb.HealthcareBetaFhirStore, error) {
	cl, err := createConfigFhirStore(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyFhirStore(ctx, cl, request)
}

// DeleteFhirStore handles the gRPC request by passing it to the underlying FhirStore Delete() method.
func (s *FhirStoreServer) DeleteHealthcareBetaFhirStore(ctx context.Context, request *betapb.DeleteHealthcareBetaFhirStoreRequest) (*emptypb.Empty, error) {

	cl, err := createConfigFhirStore(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteFhirStore(ctx, ProtoToFhirStore(request.GetResource()))

}

// ListHealthcareBetaFhirStore handles the gRPC request by passing it to the underlying FhirStoreList() method.
func (s *FhirStoreServer) ListHealthcareBetaFhirStore(ctx context.Context, request *betapb.ListHealthcareBetaFhirStoreRequest) (*betapb.ListHealthcareBetaFhirStoreResponse, error) {
	cl, err := createConfigFhirStore(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListFhirStore(ctx, request.GetProject(), request.GetLocation(), request.GetDataset())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.HealthcareBetaFhirStore
	for _, r := range resources.Items {
		rp := FhirStoreToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListHealthcareBetaFhirStoreResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigFhirStore(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
