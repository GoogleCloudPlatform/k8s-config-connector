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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/healthcare/alpha/healthcare_alpha_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/healthcare/alpha"
)

// FhirStoreServer implements the gRPC interface for FhirStore.
type FhirStoreServer struct{}

// ProtoToFhirStoreVersionEnum converts a FhirStoreVersionEnum enum from its proto representation.
func ProtoToHealthcareAlphaFhirStoreVersionEnum(e alphapb.HealthcareAlphaFhirStoreVersionEnum) *alpha.FhirStoreVersionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.HealthcareAlphaFhirStoreVersionEnum_name[int32(e)]; ok {
		e := alpha.FhirStoreVersionEnum(n[len("HealthcareAlphaFhirStoreVersionEnum"):])
		return &e
	}
	return nil
}

// ProtoToFhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnum converts a FhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnum enum from its proto representation.
func ProtoToHealthcareAlphaFhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnum(e alphapb.HealthcareAlphaFhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnum) *alpha.FhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.HealthcareAlphaFhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnum_name[int32(e)]; ok {
		e := alpha.FhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnum(n[len("HealthcareAlphaFhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToFhirStoreComplexDataTypeReferenceParsingEnum converts a FhirStoreComplexDataTypeReferenceParsingEnum enum from its proto representation.
func ProtoToHealthcareAlphaFhirStoreComplexDataTypeReferenceParsingEnum(e alphapb.HealthcareAlphaFhirStoreComplexDataTypeReferenceParsingEnum) *alpha.FhirStoreComplexDataTypeReferenceParsingEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.HealthcareAlphaFhirStoreComplexDataTypeReferenceParsingEnum_name[int32(e)]; ok {
		e := alpha.FhirStoreComplexDataTypeReferenceParsingEnum(n[len("HealthcareAlphaFhirStoreComplexDataTypeReferenceParsingEnum"):])
		return &e
	}
	return nil
}

// ProtoToFhirStoreNotificationConfig converts a FhirStoreNotificationConfig object from its proto representation.
func ProtoToHealthcareAlphaFhirStoreNotificationConfig(p *alphapb.HealthcareAlphaFhirStoreNotificationConfig) *alpha.FhirStoreNotificationConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.FhirStoreNotificationConfig{
		PubsubTopic: dcl.StringOrNil(p.GetPubsubTopic()),
	}
	return obj
}

// ProtoToFhirStoreStreamConfigs converts a FhirStoreStreamConfigs object from its proto representation.
func ProtoToHealthcareAlphaFhirStoreStreamConfigs(p *alphapb.HealthcareAlphaFhirStoreStreamConfigs) *alpha.FhirStoreStreamConfigs {
	if p == nil {
		return nil
	}
	obj := &alpha.FhirStoreStreamConfigs{
		BigqueryDestination: ProtoToHealthcareAlphaFhirStoreStreamConfigsBigqueryDestination(p.GetBigqueryDestination()),
	}
	for _, r := range p.GetResourceTypes() {
		obj.ResourceTypes = append(obj.ResourceTypes, r)
	}
	return obj
}

// ProtoToFhirStoreStreamConfigsBigqueryDestination converts a FhirStoreStreamConfigsBigqueryDestination object from its proto representation.
func ProtoToHealthcareAlphaFhirStoreStreamConfigsBigqueryDestination(p *alphapb.HealthcareAlphaFhirStoreStreamConfigsBigqueryDestination) *alpha.FhirStoreStreamConfigsBigqueryDestination {
	if p == nil {
		return nil
	}
	obj := &alpha.FhirStoreStreamConfigsBigqueryDestination{
		DatasetUri:   dcl.StringOrNil(p.GetDatasetUri()),
		SchemaConfig: ProtoToHealthcareAlphaFhirStoreStreamConfigsBigqueryDestinationSchemaConfig(p.GetSchemaConfig()),
	}
	return obj
}

// ProtoToFhirStoreStreamConfigsBigqueryDestinationSchemaConfig converts a FhirStoreStreamConfigsBigqueryDestinationSchemaConfig object from its proto representation.
func ProtoToHealthcareAlphaFhirStoreStreamConfigsBigqueryDestinationSchemaConfig(p *alphapb.HealthcareAlphaFhirStoreStreamConfigsBigqueryDestinationSchemaConfig) *alpha.FhirStoreStreamConfigsBigqueryDestinationSchemaConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.FhirStoreStreamConfigsBigqueryDestinationSchemaConfig{
		SchemaType:              ProtoToHealthcareAlphaFhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnum(p.GetSchemaType()),
		RecursiveStructureDepth: dcl.Int64OrNil(p.GetRecursiveStructureDepth()),
	}
	return obj
}

// ProtoToFhirStoreValidationConfig converts a FhirStoreValidationConfig object from its proto representation.
func ProtoToHealthcareAlphaFhirStoreValidationConfig(p *alphapb.HealthcareAlphaFhirStoreValidationConfig) *alpha.FhirStoreValidationConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.FhirStoreValidationConfig{
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
func ProtoToFhirStore(p *alphapb.HealthcareAlphaFhirStore) *alpha.FhirStore {
	obj := &alpha.FhirStore{
		Name:                            dcl.StringOrNil(p.GetName()),
		EnableUpdateCreate:              dcl.Bool(p.GetEnableUpdateCreate()),
		NotificationConfig:              ProtoToHealthcareAlphaFhirStoreNotificationConfig(p.GetNotificationConfig()),
		DisableReferentialIntegrity:     dcl.Bool(p.GetDisableReferentialIntegrity()),
		ShardNum:                        dcl.Int64OrNil(p.GetShardNum()),
		DisableResourceVersioning:       dcl.Bool(p.GetDisableResourceVersioning()),
		Version:                         ProtoToHealthcareAlphaFhirStoreVersionEnum(p.GetVersion()),
		ValidationConfig:                ProtoToHealthcareAlphaFhirStoreValidationConfig(p.GetValidationConfig()),
		DefaultSearchHandlingStrict:     dcl.Bool(p.GetDefaultSearchHandlingStrict()),
		ComplexDataTypeReferenceParsing: ProtoToHealthcareAlphaFhirStoreComplexDataTypeReferenceParsingEnum(p.GetComplexDataTypeReferenceParsing()),
		Project:                         dcl.StringOrNil(p.GetProject()),
		Location:                        dcl.StringOrNil(p.GetLocation()),
		Dataset:                         dcl.StringOrNil(p.GetDataset()),
	}
	for _, r := range p.GetStreamConfigs() {
		obj.StreamConfigs = append(obj.StreamConfigs, *ProtoToHealthcareAlphaFhirStoreStreamConfigs(r))
	}
	return obj
}

// FhirStoreVersionEnumToProto converts a FhirStoreVersionEnum enum to its proto representation.
func HealthcareAlphaFhirStoreVersionEnumToProto(e *alpha.FhirStoreVersionEnum) alphapb.HealthcareAlphaFhirStoreVersionEnum {
	if e == nil {
		return alphapb.HealthcareAlphaFhirStoreVersionEnum(0)
	}
	if v, ok := alphapb.HealthcareAlphaFhirStoreVersionEnum_value["FhirStoreVersionEnum"+string(*e)]; ok {
		return alphapb.HealthcareAlphaFhirStoreVersionEnum(v)
	}
	return alphapb.HealthcareAlphaFhirStoreVersionEnum(0)
}

// FhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnumToProto converts a FhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnum enum to its proto representation.
func HealthcareAlphaFhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnumToProto(e *alpha.FhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnum) alphapb.HealthcareAlphaFhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnum {
	if e == nil {
		return alphapb.HealthcareAlphaFhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnum(0)
	}
	if v, ok := alphapb.HealthcareAlphaFhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnum_value["FhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnum"+string(*e)]; ok {
		return alphapb.HealthcareAlphaFhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnum(v)
	}
	return alphapb.HealthcareAlphaFhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnum(0)
}

// FhirStoreComplexDataTypeReferenceParsingEnumToProto converts a FhirStoreComplexDataTypeReferenceParsingEnum enum to its proto representation.
func HealthcareAlphaFhirStoreComplexDataTypeReferenceParsingEnumToProto(e *alpha.FhirStoreComplexDataTypeReferenceParsingEnum) alphapb.HealthcareAlphaFhirStoreComplexDataTypeReferenceParsingEnum {
	if e == nil {
		return alphapb.HealthcareAlphaFhirStoreComplexDataTypeReferenceParsingEnum(0)
	}
	if v, ok := alphapb.HealthcareAlphaFhirStoreComplexDataTypeReferenceParsingEnum_value["FhirStoreComplexDataTypeReferenceParsingEnum"+string(*e)]; ok {
		return alphapb.HealthcareAlphaFhirStoreComplexDataTypeReferenceParsingEnum(v)
	}
	return alphapb.HealthcareAlphaFhirStoreComplexDataTypeReferenceParsingEnum(0)
}

// FhirStoreNotificationConfigToProto converts a FhirStoreNotificationConfig object to its proto representation.
func HealthcareAlphaFhirStoreNotificationConfigToProto(o *alpha.FhirStoreNotificationConfig) *alphapb.HealthcareAlphaFhirStoreNotificationConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.HealthcareAlphaFhirStoreNotificationConfig{}
	p.SetPubsubTopic(dcl.ValueOrEmptyString(o.PubsubTopic))
	return p
}

// FhirStoreStreamConfigsToProto converts a FhirStoreStreamConfigs object to its proto representation.
func HealthcareAlphaFhirStoreStreamConfigsToProto(o *alpha.FhirStoreStreamConfigs) *alphapb.HealthcareAlphaFhirStoreStreamConfigs {
	if o == nil {
		return nil
	}
	p := &alphapb.HealthcareAlphaFhirStoreStreamConfigs{}
	p.SetBigqueryDestination(HealthcareAlphaFhirStoreStreamConfigsBigqueryDestinationToProto(o.BigqueryDestination))
	sResourceTypes := make([]string, len(o.ResourceTypes))
	for i, r := range o.ResourceTypes {
		sResourceTypes[i] = r
	}
	p.SetResourceTypes(sResourceTypes)
	return p
}

// FhirStoreStreamConfigsBigqueryDestinationToProto converts a FhirStoreStreamConfigsBigqueryDestination object to its proto representation.
func HealthcareAlphaFhirStoreStreamConfigsBigqueryDestinationToProto(o *alpha.FhirStoreStreamConfigsBigqueryDestination) *alphapb.HealthcareAlphaFhirStoreStreamConfigsBigqueryDestination {
	if o == nil {
		return nil
	}
	p := &alphapb.HealthcareAlphaFhirStoreStreamConfigsBigqueryDestination{}
	p.SetDatasetUri(dcl.ValueOrEmptyString(o.DatasetUri))
	p.SetSchemaConfig(HealthcareAlphaFhirStoreStreamConfigsBigqueryDestinationSchemaConfigToProto(o.SchemaConfig))
	return p
}

// FhirStoreStreamConfigsBigqueryDestinationSchemaConfigToProto converts a FhirStoreStreamConfigsBigqueryDestinationSchemaConfig object to its proto representation.
func HealthcareAlphaFhirStoreStreamConfigsBigqueryDestinationSchemaConfigToProto(o *alpha.FhirStoreStreamConfigsBigqueryDestinationSchemaConfig) *alphapb.HealthcareAlphaFhirStoreStreamConfigsBigqueryDestinationSchemaConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.HealthcareAlphaFhirStoreStreamConfigsBigqueryDestinationSchemaConfig{}
	p.SetSchemaType(HealthcareAlphaFhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnumToProto(o.SchemaType))
	p.SetRecursiveStructureDepth(dcl.ValueOrEmptyInt64(o.RecursiveStructureDepth))
	return p
}

// FhirStoreValidationConfigToProto converts a FhirStoreValidationConfig object to its proto representation.
func HealthcareAlphaFhirStoreValidationConfigToProto(o *alpha.FhirStoreValidationConfig) *alphapb.HealthcareAlphaFhirStoreValidationConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.HealthcareAlphaFhirStoreValidationConfig{}
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
func FhirStoreToProto(resource *alpha.FhirStore) *alphapb.HealthcareAlphaFhirStore {
	p := &alphapb.HealthcareAlphaFhirStore{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetEnableUpdateCreate(dcl.ValueOrEmptyBool(resource.EnableUpdateCreate))
	p.SetNotificationConfig(HealthcareAlphaFhirStoreNotificationConfigToProto(resource.NotificationConfig))
	p.SetDisableReferentialIntegrity(dcl.ValueOrEmptyBool(resource.DisableReferentialIntegrity))
	p.SetShardNum(dcl.ValueOrEmptyInt64(resource.ShardNum))
	p.SetDisableResourceVersioning(dcl.ValueOrEmptyBool(resource.DisableResourceVersioning))
	p.SetVersion(HealthcareAlphaFhirStoreVersionEnumToProto(resource.Version))
	p.SetValidationConfig(HealthcareAlphaFhirStoreValidationConfigToProto(resource.ValidationConfig))
	p.SetDefaultSearchHandlingStrict(dcl.ValueOrEmptyBool(resource.DefaultSearchHandlingStrict))
	p.SetComplexDataTypeReferenceParsing(HealthcareAlphaFhirStoreComplexDataTypeReferenceParsingEnumToProto(resource.ComplexDataTypeReferenceParsing))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	p.SetDataset(dcl.ValueOrEmptyString(resource.Dataset))
	mLabels := make(map[string]string, len(resource.Labels))
	for k, r := range resource.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)
	sStreamConfigs := make([]*alphapb.HealthcareAlphaFhirStoreStreamConfigs, len(resource.StreamConfigs))
	for i, r := range resource.StreamConfigs {
		sStreamConfigs[i] = HealthcareAlphaFhirStoreStreamConfigsToProto(&r)
	}
	p.SetStreamConfigs(sStreamConfigs)

	return p
}

// applyFhirStore handles the gRPC request by passing it to the underlying FhirStore Apply() method.
func (s *FhirStoreServer) applyFhirStore(ctx context.Context, c *alpha.Client, request *alphapb.ApplyHealthcareAlphaFhirStoreRequest) (*alphapb.HealthcareAlphaFhirStore, error) {
	p := ProtoToFhirStore(request.GetResource())
	res, err := c.ApplyFhirStore(ctx, p)
	if err != nil {
		return nil, err
	}
	r := FhirStoreToProto(res)
	return r, nil
}

// applyHealthcareAlphaFhirStore handles the gRPC request by passing it to the underlying FhirStore Apply() method.
func (s *FhirStoreServer) ApplyHealthcareAlphaFhirStore(ctx context.Context, request *alphapb.ApplyHealthcareAlphaFhirStoreRequest) (*alphapb.HealthcareAlphaFhirStore, error) {
	cl, err := createConfigFhirStore(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyFhirStore(ctx, cl, request)
}

// DeleteFhirStore handles the gRPC request by passing it to the underlying FhirStore Delete() method.
func (s *FhirStoreServer) DeleteHealthcareAlphaFhirStore(ctx context.Context, request *alphapb.DeleteHealthcareAlphaFhirStoreRequest) (*emptypb.Empty, error) {

	cl, err := createConfigFhirStore(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteFhirStore(ctx, ProtoToFhirStore(request.GetResource()))

}

// ListHealthcareAlphaFhirStore handles the gRPC request by passing it to the underlying FhirStoreList() method.
func (s *FhirStoreServer) ListHealthcareAlphaFhirStore(ctx context.Context, request *alphapb.ListHealthcareAlphaFhirStoreRequest) (*alphapb.ListHealthcareAlphaFhirStoreResponse, error) {
	cl, err := createConfigFhirStore(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListFhirStore(ctx, request.GetProject(), request.GetLocation(), request.GetDataset())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.HealthcareAlphaFhirStore
	for _, r := range resources.Items {
		rp := FhirStoreToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListHealthcareAlphaFhirStoreResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigFhirStore(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
