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
	healthcarepb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/healthcare/healthcare_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/healthcare"
)

// FhirStoreServer implements the gRPC interface for FhirStore.
type FhirStoreServer struct{}

// ProtoToFhirStoreVersionEnum converts a FhirStoreVersionEnum enum from its proto representation.
func ProtoToHealthcareFhirStoreVersionEnum(e healthcarepb.HealthcareFhirStoreVersionEnum) *healthcare.FhirStoreVersionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := healthcarepb.HealthcareFhirStoreVersionEnum_name[int32(e)]; ok {
		e := healthcare.FhirStoreVersionEnum(n[len("HealthcareFhirStoreVersionEnum"):])
		return &e
	}
	return nil
}

// ProtoToFhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnum converts a FhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnum enum from its proto representation.
func ProtoToHealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnum(e healthcarepb.HealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnum) *healthcare.FhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := healthcarepb.HealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnum_name[int32(e)]; ok {
		e := healthcare.FhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnum(n[len("HealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToFhirStoreComplexDataTypeReferenceParsingEnum converts a FhirStoreComplexDataTypeReferenceParsingEnum enum from its proto representation.
func ProtoToHealthcareFhirStoreComplexDataTypeReferenceParsingEnum(e healthcarepb.HealthcareFhirStoreComplexDataTypeReferenceParsingEnum) *healthcare.FhirStoreComplexDataTypeReferenceParsingEnum {
	if e == 0 {
		return nil
	}
	if n, ok := healthcarepb.HealthcareFhirStoreComplexDataTypeReferenceParsingEnum_name[int32(e)]; ok {
		e := healthcare.FhirStoreComplexDataTypeReferenceParsingEnum(n[len("HealthcareFhirStoreComplexDataTypeReferenceParsingEnum"):])
		return &e
	}
	return nil
}

// ProtoToFhirStoreNotificationConfig converts a FhirStoreNotificationConfig object from its proto representation.
func ProtoToHealthcareFhirStoreNotificationConfig(p *healthcarepb.HealthcareFhirStoreNotificationConfig) *healthcare.FhirStoreNotificationConfig {
	if p == nil {
		return nil
	}
	obj := &healthcare.FhirStoreNotificationConfig{
		PubsubTopic: dcl.StringOrNil(p.GetPubsubTopic()),
	}
	return obj
}

// ProtoToFhirStoreStreamConfigs converts a FhirStoreStreamConfigs object from its proto representation.
func ProtoToHealthcareFhirStoreStreamConfigs(p *healthcarepb.HealthcareFhirStoreStreamConfigs) *healthcare.FhirStoreStreamConfigs {
	if p == nil {
		return nil
	}
	obj := &healthcare.FhirStoreStreamConfigs{
		BigqueryDestination: ProtoToHealthcareFhirStoreStreamConfigsBigqueryDestination(p.GetBigqueryDestination()),
	}
	for _, r := range p.GetResourceTypes() {
		obj.ResourceTypes = append(obj.ResourceTypes, r)
	}
	return obj
}

// ProtoToFhirStoreStreamConfigsBigqueryDestination converts a FhirStoreStreamConfigsBigqueryDestination object from its proto representation.
func ProtoToHealthcareFhirStoreStreamConfigsBigqueryDestination(p *healthcarepb.HealthcareFhirStoreStreamConfigsBigqueryDestination) *healthcare.FhirStoreStreamConfigsBigqueryDestination {
	if p == nil {
		return nil
	}
	obj := &healthcare.FhirStoreStreamConfigsBigqueryDestination{
		DatasetUri:   dcl.StringOrNil(p.GetDatasetUri()),
		SchemaConfig: ProtoToHealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfig(p.GetSchemaConfig()),
	}
	return obj
}

// ProtoToFhirStoreStreamConfigsBigqueryDestinationSchemaConfig converts a FhirStoreStreamConfigsBigqueryDestinationSchemaConfig object from its proto representation.
func ProtoToHealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfig(p *healthcarepb.HealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfig) *healthcare.FhirStoreStreamConfigsBigqueryDestinationSchemaConfig {
	if p == nil {
		return nil
	}
	obj := &healthcare.FhirStoreStreamConfigsBigqueryDestinationSchemaConfig{
		SchemaType:              ProtoToHealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnum(p.GetSchemaType()),
		RecursiveStructureDepth: dcl.Int64OrNil(p.GetRecursiveStructureDepth()),
	}
	return obj
}

// ProtoToFhirStoreValidationConfig converts a FhirStoreValidationConfig object from its proto representation.
func ProtoToHealthcareFhirStoreValidationConfig(p *healthcarepb.HealthcareFhirStoreValidationConfig) *healthcare.FhirStoreValidationConfig {
	if p == nil {
		return nil
	}
	obj := &healthcare.FhirStoreValidationConfig{
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
func ProtoToFhirStore(p *healthcarepb.HealthcareFhirStore) *healthcare.FhirStore {
	obj := &healthcare.FhirStore{
		Name:                            dcl.StringOrNil(p.GetName()),
		EnableUpdateCreate:              dcl.Bool(p.GetEnableUpdateCreate()),
		NotificationConfig:              ProtoToHealthcareFhirStoreNotificationConfig(p.GetNotificationConfig()),
		DisableReferentialIntegrity:     dcl.Bool(p.GetDisableReferentialIntegrity()),
		ShardNum:                        dcl.Int64OrNil(p.GetShardNum()),
		DisableResourceVersioning:       dcl.Bool(p.GetDisableResourceVersioning()),
		Version:                         ProtoToHealthcareFhirStoreVersionEnum(p.GetVersion()),
		ValidationConfig:                ProtoToHealthcareFhirStoreValidationConfig(p.GetValidationConfig()),
		DefaultSearchHandlingStrict:     dcl.Bool(p.GetDefaultSearchHandlingStrict()),
		ComplexDataTypeReferenceParsing: ProtoToHealthcareFhirStoreComplexDataTypeReferenceParsingEnum(p.GetComplexDataTypeReferenceParsing()),
		Project:                         dcl.StringOrNil(p.GetProject()),
		Location:                        dcl.StringOrNil(p.GetLocation()),
		Dataset:                         dcl.StringOrNil(p.GetDataset()),
	}
	for _, r := range p.GetStreamConfigs() {
		obj.StreamConfigs = append(obj.StreamConfigs, *ProtoToHealthcareFhirStoreStreamConfigs(r))
	}
	return obj
}

// FhirStoreVersionEnumToProto converts a FhirStoreVersionEnum enum to its proto representation.
func HealthcareFhirStoreVersionEnumToProto(e *healthcare.FhirStoreVersionEnum) healthcarepb.HealthcareFhirStoreVersionEnum {
	if e == nil {
		return healthcarepb.HealthcareFhirStoreVersionEnum(0)
	}
	if v, ok := healthcarepb.HealthcareFhirStoreVersionEnum_value["FhirStoreVersionEnum"+string(*e)]; ok {
		return healthcarepb.HealthcareFhirStoreVersionEnum(v)
	}
	return healthcarepb.HealthcareFhirStoreVersionEnum(0)
}

// FhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnumToProto converts a FhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnum enum to its proto representation.
func HealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnumToProto(e *healthcare.FhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnum) healthcarepb.HealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnum {
	if e == nil {
		return healthcarepb.HealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnum(0)
	}
	if v, ok := healthcarepb.HealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnum_value["FhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnum"+string(*e)]; ok {
		return healthcarepb.HealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnum(v)
	}
	return healthcarepb.HealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnum(0)
}

// FhirStoreComplexDataTypeReferenceParsingEnumToProto converts a FhirStoreComplexDataTypeReferenceParsingEnum enum to its proto representation.
func HealthcareFhirStoreComplexDataTypeReferenceParsingEnumToProto(e *healthcare.FhirStoreComplexDataTypeReferenceParsingEnum) healthcarepb.HealthcareFhirStoreComplexDataTypeReferenceParsingEnum {
	if e == nil {
		return healthcarepb.HealthcareFhirStoreComplexDataTypeReferenceParsingEnum(0)
	}
	if v, ok := healthcarepb.HealthcareFhirStoreComplexDataTypeReferenceParsingEnum_value["FhirStoreComplexDataTypeReferenceParsingEnum"+string(*e)]; ok {
		return healthcarepb.HealthcareFhirStoreComplexDataTypeReferenceParsingEnum(v)
	}
	return healthcarepb.HealthcareFhirStoreComplexDataTypeReferenceParsingEnum(0)
}

// FhirStoreNotificationConfigToProto converts a FhirStoreNotificationConfig object to its proto representation.
func HealthcareFhirStoreNotificationConfigToProto(o *healthcare.FhirStoreNotificationConfig) *healthcarepb.HealthcareFhirStoreNotificationConfig {
	if o == nil {
		return nil
	}
	p := &healthcarepb.HealthcareFhirStoreNotificationConfig{}
	p.SetPubsubTopic(dcl.ValueOrEmptyString(o.PubsubTopic))
	return p
}

// FhirStoreStreamConfigsToProto converts a FhirStoreStreamConfigs object to its proto representation.
func HealthcareFhirStoreStreamConfigsToProto(o *healthcare.FhirStoreStreamConfigs) *healthcarepb.HealthcareFhirStoreStreamConfigs {
	if o == nil {
		return nil
	}
	p := &healthcarepb.HealthcareFhirStoreStreamConfigs{}
	p.SetBigqueryDestination(HealthcareFhirStoreStreamConfigsBigqueryDestinationToProto(o.BigqueryDestination))
	sResourceTypes := make([]string, len(o.ResourceTypes))
	for i, r := range o.ResourceTypes {
		sResourceTypes[i] = r
	}
	p.SetResourceTypes(sResourceTypes)
	return p
}

// FhirStoreStreamConfigsBigqueryDestinationToProto converts a FhirStoreStreamConfigsBigqueryDestination object to its proto representation.
func HealthcareFhirStoreStreamConfigsBigqueryDestinationToProto(o *healthcare.FhirStoreStreamConfigsBigqueryDestination) *healthcarepb.HealthcareFhirStoreStreamConfigsBigqueryDestination {
	if o == nil {
		return nil
	}
	p := &healthcarepb.HealthcareFhirStoreStreamConfigsBigqueryDestination{}
	p.SetDatasetUri(dcl.ValueOrEmptyString(o.DatasetUri))
	p.SetSchemaConfig(HealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfigToProto(o.SchemaConfig))
	return p
}

// FhirStoreStreamConfigsBigqueryDestinationSchemaConfigToProto converts a FhirStoreStreamConfigsBigqueryDestinationSchemaConfig object to its proto representation.
func HealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfigToProto(o *healthcare.FhirStoreStreamConfigsBigqueryDestinationSchemaConfig) *healthcarepb.HealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfig {
	if o == nil {
		return nil
	}
	p := &healthcarepb.HealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfig{}
	p.SetSchemaType(HealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaTypeEnumToProto(o.SchemaType))
	p.SetRecursiveStructureDepth(dcl.ValueOrEmptyInt64(o.RecursiveStructureDepth))
	return p
}

// FhirStoreValidationConfigToProto converts a FhirStoreValidationConfig object to its proto representation.
func HealthcareFhirStoreValidationConfigToProto(o *healthcare.FhirStoreValidationConfig) *healthcarepb.HealthcareFhirStoreValidationConfig {
	if o == nil {
		return nil
	}
	p := &healthcarepb.HealthcareFhirStoreValidationConfig{}
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
func FhirStoreToProto(resource *healthcare.FhirStore) *healthcarepb.HealthcareFhirStore {
	p := &healthcarepb.HealthcareFhirStore{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetEnableUpdateCreate(dcl.ValueOrEmptyBool(resource.EnableUpdateCreate))
	p.SetNotificationConfig(HealthcareFhirStoreNotificationConfigToProto(resource.NotificationConfig))
	p.SetDisableReferentialIntegrity(dcl.ValueOrEmptyBool(resource.DisableReferentialIntegrity))
	p.SetShardNum(dcl.ValueOrEmptyInt64(resource.ShardNum))
	p.SetDisableResourceVersioning(dcl.ValueOrEmptyBool(resource.DisableResourceVersioning))
	p.SetVersion(HealthcareFhirStoreVersionEnumToProto(resource.Version))
	p.SetValidationConfig(HealthcareFhirStoreValidationConfigToProto(resource.ValidationConfig))
	p.SetDefaultSearchHandlingStrict(dcl.ValueOrEmptyBool(resource.DefaultSearchHandlingStrict))
	p.SetComplexDataTypeReferenceParsing(HealthcareFhirStoreComplexDataTypeReferenceParsingEnumToProto(resource.ComplexDataTypeReferenceParsing))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	p.SetDataset(dcl.ValueOrEmptyString(resource.Dataset))
	mLabels := make(map[string]string, len(resource.Labels))
	for k, r := range resource.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)
	sStreamConfigs := make([]*healthcarepb.HealthcareFhirStoreStreamConfigs, len(resource.StreamConfigs))
	for i, r := range resource.StreamConfigs {
		sStreamConfigs[i] = HealthcareFhirStoreStreamConfigsToProto(&r)
	}
	p.SetStreamConfigs(sStreamConfigs)

	return p
}

// applyFhirStore handles the gRPC request by passing it to the underlying FhirStore Apply() method.
func (s *FhirStoreServer) applyFhirStore(ctx context.Context, c *healthcare.Client, request *healthcarepb.ApplyHealthcareFhirStoreRequest) (*healthcarepb.HealthcareFhirStore, error) {
	p := ProtoToFhirStore(request.GetResource())
	res, err := c.ApplyFhirStore(ctx, p)
	if err != nil {
		return nil, err
	}
	r := FhirStoreToProto(res)
	return r, nil
}

// applyHealthcareFhirStore handles the gRPC request by passing it to the underlying FhirStore Apply() method.
func (s *FhirStoreServer) ApplyHealthcareFhirStore(ctx context.Context, request *healthcarepb.ApplyHealthcareFhirStoreRequest) (*healthcarepb.HealthcareFhirStore, error) {
	cl, err := createConfigFhirStore(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyFhirStore(ctx, cl, request)
}

// DeleteFhirStore handles the gRPC request by passing it to the underlying FhirStore Delete() method.
func (s *FhirStoreServer) DeleteHealthcareFhirStore(ctx context.Context, request *healthcarepb.DeleteHealthcareFhirStoreRequest) (*emptypb.Empty, error) {

	cl, err := createConfigFhirStore(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteFhirStore(ctx, ProtoToFhirStore(request.GetResource()))

}

// ListHealthcareFhirStore handles the gRPC request by passing it to the underlying FhirStoreList() method.
func (s *FhirStoreServer) ListHealthcareFhirStore(ctx context.Context, request *healthcarepb.ListHealthcareFhirStoreRequest) (*healthcarepb.ListHealthcareFhirStoreResponse, error) {
	cl, err := createConfigFhirStore(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListFhirStore(ctx, request.GetProject(), request.GetLocation(), request.GetDataset())
	if err != nil {
		return nil, err
	}
	var protos []*healthcarepb.HealthcareFhirStore
	for _, r := range resources.Items {
		rp := FhirStoreToProto(r)
		protos = append(protos, rp)
	}
	p := &healthcarepb.ListHealthcareFhirStoreResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigFhirStore(ctx context.Context, service_account_file string) (*healthcare.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return healthcare.NewClient(conf), nil
}
