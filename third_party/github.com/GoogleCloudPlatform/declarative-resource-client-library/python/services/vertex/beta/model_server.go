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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/vertex/beta/vertex_beta_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/vertex/beta"
)

// ModelServer implements the gRPC interface for Model.
type ModelServer struct{}

// ProtoToModelSupportedExportFormatsExportableContentsEnum converts a ModelSupportedExportFormatsExportableContentsEnum enum from its proto representation.
func ProtoToVertexBetaModelSupportedExportFormatsExportableContentsEnum(e betapb.VertexBetaModelSupportedExportFormatsExportableContentsEnum) *beta.ModelSupportedExportFormatsExportableContentsEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.VertexBetaModelSupportedExportFormatsExportableContentsEnum_name[int32(e)]; ok {
		e := beta.ModelSupportedExportFormatsExportableContentsEnum(n[len("VertexBetaModelSupportedExportFormatsExportableContentsEnum"):])
		return &e
	}
	return nil
}

// ProtoToModelContainerSpecAcceleratorRequirementsTypeEnum converts a ModelContainerSpecAcceleratorRequirementsTypeEnum enum from its proto representation.
func ProtoToVertexBetaModelContainerSpecAcceleratorRequirementsTypeEnum(e betapb.VertexBetaModelContainerSpecAcceleratorRequirementsTypeEnum) *beta.ModelContainerSpecAcceleratorRequirementsTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.VertexBetaModelContainerSpecAcceleratorRequirementsTypeEnum_name[int32(e)]; ok {
		e := beta.ModelContainerSpecAcceleratorRequirementsTypeEnum(n[len("VertexBetaModelContainerSpecAcceleratorRequirementsTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToModelSupportedDeploymentResourcesTypesEnum converts a ModelSupportedDeploymentResourcesTypesEnum enum from its proto representation.
func ProtoToVertexBetaModelSupportedDeploymentResourcesTypesEnum(e betapb.VertexBetaModelSupportedDeploymentResourcesTypesEnum) *beta.ModelSupportedDeploymentResourcesTypesEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.VertexBetaModelSupportedDeploymentResourcesTypesEnum_name[int32(e)]; ok {
		e := beta.ModelSupportedDeploymentResourcesTypesEnum(n[len("VertexBetaModelSupportedDeploymentResourcesTypesEnum"):])
		return &e
	}
	return nil
}

// ProtoToModelSupportedExportFormats converts a ModelSupportedExportFormats object from its proto representation.
func ProtoToVertexBetaModelSupportedExportFormats(p *betapb.VertexBetaModelSupportedExportFormats) *beta.ModelSupportedExportFormats {
	if p == nil {
		return nil
	}
	obj := &beta.ModelSupportedExportFormats{
		Id: dcl.StringOrNil(p.GetId()),
	}
	for _, r := range p.GetExportableContents() {
		obj.ExportableContents = append(obj.ExportableContents, *ProtoToVertexBetaModelSupportedExportFormatsExportableContentsEnum(r))
	}
	return obj
}

// ProtoToModelOriginalModelInfo converts a ModelOriginalModelInfo object from its proto representation.
func ProtoToVertexBetaModelOriginalModelInfo(p *betapb.VertexBetaModelOriginalModelInfo) *beta.ModelOriginalModelInfo {
	if p == nil {
		return nil
	}
	obj := &beta.ModelOriginalModelInfo{
		Model: dcl.StringOrNil(p.GetModel()),
	}
	return obj
}

// ProtoToModelContainerSpec converts a ModelContainerSpec object from its proto representation.
func ProtoToVertexBetaModelContainerSpec(p *betapb.VertexBetaModelContainerSpec) *beta.ModelContainerSpec {
	if p == nil {
		return nil
	}
	obj := &beta.ModelContainerSpec{
		ImageUri:     dcl.StringOrNil(p.GetImageUri()),
		PredictRoute: dcl.StringOrNil(p.GetPredictRoute()),
		HealthRoute:  dcl.StringOrNil(p.GetHealthRoute()),
	}
	for _, r := range p.GetCommand() {
		obj.Command = append(obj.Command, r)
	}
	for _, r := range p.GetArgs() {
		obj.Args = append(obj.Args, r)
	}
	for _, r := range p.GetEnv() {
		obj.Env = append(obj.Env, *ProtoToVertexBetaModelContainerSpecEnv(r))
	}
	for _, r := range p.GetPorts() {
		obj.Ports = append(obj.Ports, *ProtoToVertexBetaModelContainerSpecPorts(r))
	}
	for _, r := range p.GetAcceleratorRequirements() {
		obj.AcceleratorRequirements = append(obj.AcceleratorRequirements, *ProtoToVertexBetaModelContainerSpecAcceleratorRequirements(r))
	}
	return obj
}

// ProtoToModelContainerSpecEnv converts a ModelContainerSpecEnv object from its proto representation.
func ProtoToVertexBetaModelContainerSpecEnv(p *betapb.VertexBetaModelContainerSpecEnv) *beta.ModelContainerSpecEnv {
	if p == nil {
		return nil
	}
	obj := &beta.ModelContainerSpecEnv{
		Name:  dcl.StringOrNil(p.GetName()),
		Value: dcl.StringOrNil(p.GetValue()),
	}
	return obj
}

// ProtoToModelContainerSpecPorts converts a ModelContainerSpecPorts object from its proto representation.
func ProtoToVertexBetaModelContainerSpecPorts(p *betapb.VertexBetaModelContainerSpecPorts) *beta.ModelContainerSpecPorts {
	if p == nil {
		return nil
	}
	obj := &beta.ModelContainerSpecPorts{
		ContainerPort: dcl.Int64OrNil(p.GetContainerPort()),
	}
	return obj
}

// ProtoToModelContainerSpecAcceleratorRequirements converts a ModelContainerSpecAcceleratorRequirements object from its proto representation.
func ProtoToVertexBetaModelContainerSpecAcceleratorRequirements(p *betapb.VertexBetaModelContainerSpecAcceleratorRequirements) *beta.ModelContainerSpecAcceleratorRequirements {
	if p == nil {
		return nil
	}
	obj := &beta.ModelContainerSpecAcceleratorRequirements{
		Type:  ProtoToVertexBetaModelContainerSpecAcceleratorRequirementsTypeEnum(p.GetType()),
		Count: dcl.Int64OrNil(p.GetCount()),
	}
	return obj
}

// ProtoToModelDeployedModels converts a ModelDeployedModels object from its proto representation.
func ProtoToVertexBetaModelDeployedModels(p *betapb.VertexBetaModelDeployedModels) *beta.ModelDeployedModels {
	if p == nil {
		return nil
	}
	obj := &beta.ModelDeployedModels{
		Endpoint:        dcl.StringOrNil(p.GetEndpoint()),
		DeployedModelId: dcl.StringOrNil(p.GetDeployedModelId()),
	}
	return obj
}

// ProtoToModelEncryptionSpec converts a ModelEncryptionSpec object from its proto representation.
func ProtoToVertexBetaModelEncryptionSpec(p *betapb.VertexBetaModelEncryptionSpec) *beta.ModelEncryptionSpec {
	if p == nil {
		return nil
	}
	obj := &beta.ModelEncryptionSpec{
		KmsKeyName: dcl.StringOrNil(p.GetKmsKeyName()),
	}
	return obj
}

// ProtoToModel converts a Model resource from its proto representation.
func ProtoToModel(p *betapb.VertexBetaModel) *beta.Model {
	obj := &beta.Model{
		Name:               dcl.StringOrNil(p.GetName()),
		VersionId:          dcl.StringOrNil(p.GetVersionId()),
		VersionCreateTime:  dcl.StringOrNil(p.GetVersionCreateTime()),
		VersionUpdateTime:  dcl.StringOrNil(p.GetVersionUpdateTime()),
		DisplayName:        dcl.StringOrNil(p.GetDisplayName()),
		Description:        dcl.StringOrNil(p.GetDescription()),
		VersionDescription: dcl.StringOrNil(p.GetVersionDescription()),
		TrainingPipeline:   dcl.StringOrNil(p.GetTrainingPipeline()),
		OriginalModelInfo:  ProtoToVertexBetaModelOriginalModelInfo(p.GetOriginalModelInfo()),
		ContainerSpec:      ProtoToVertexBetaModelContainerSpec(p.GetContainerSpec()),
		ArtifactUri:        dcl.StringOrNil(p.GetArtifactUri()),
		CreateTime:         dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:         dcl.StringOrNil(p.GetUpdateTime()),
		Etag:               dcl.StringOrNil(p.GetEtag()),
		EncryptionSpec:     ProtoToVertexBetaModelEncryptionSpec(p.GetEncryptionSpec()),
		Project:            dcl.StringOrNil(p.GetProject()),
		Location:           dcl.StringOrNil(p.GetLocation()),
	}
	for _, r := range p.GetVersionAliases() {
		obj.VersionAliases = append(obj.VersionAliases, r)
	}
	for _, r := range p.GetSupportedExportFormats() {
		obj.SupportedExportFormats = append(obj.SupportedExportFormats, *ProtoToVertexBetaModelSupportedExportFormats(r))
	}
	for _, r := range p.GetSupportedDeploymentResourcesTypes() {
		obj.SupportedDeploymentResourcesTypes = append(obj.SupportedDeploymentResourcesTypes, *ProtoToVertexBetaModelSupportedDeploymentResourcesTypesEnum(r))
	}
	for _, r := range p.GetSupportedInputStorageFormats() {
		obj.SupportedInputStorageFormats = append(obj.SupportedInputStorageFormats, r)
	}
	for _, r := range p.GetSupportedOutputStorageFormats() {
		obj.SupportedOutputStorageFormats = append(obj.SupportedOutputStorageFormats, r)
	}
	for _, r := range p.GetDeployedModels() {
		obj.DeployedModels = append(obj.DeployedModels, *ProtoToVertexBetaModelDeployedModels(r))
	}
	return obj
}

// ModelSupportedExportFormatsExportableContentsEnumToProto converts a ModelSupportedExportFormatsExportableContentsEnum enum to its proto representation.
func VertexBetaModelSupportedExportFormatsExportableContentsEnumToProto(e *beta.ModelSupportedExportFormatsExportableContentsEnum) betapb.VertexBetaModelSupportedExportFormatsExportableContentsEnum {
	if e == nil {
		return betapb.VertexBetaModelSupportedExportFormatsExportableContentsEnum(0)
	}
	if v, ok := betapb.VertexBetaModelSupportedExportFormatsExportableContentsEnum_value["ModelSupportedExportFormatsExportableContentsEnum"+string(*e)]; ok {
		return betapb.VertexBetaModelSupportedExportFormatsExportableContentsEnum(v)
	}
	return betapb.VertexBetaModelSupportedExportFormatsExportableContentsEnum(0)
}

// ModelContainerSpecAcceleratorRequirementsTypeEnumToProto converts a ModelContainerSpecAcceleratorRequirementsTypeEnum enum to its proto representation.
func VertexBetaModelContainerSpecAcceleratorRequirementsTypeEnumToProto(e *beta.ModelContainerSpecAcceleratorRequirementsTypeEnum) betapb.VertexBetaModelContainerSpecAcceleratorRequirementsTypeEnum {
	if e == nil {
		return betapb.VertexBetaModelContainerSpecAcceleratorRequirementsTypeEnum(0)
	}
	if v, ok := betapb.VertexBetaModelContainerSpecAcceleratorRequirementsTypeEnum_value["ModelContainerSpecAcceleratorRequirementsTypeEnum"+string(*e)]; ok {
		return betapb.VertexBetaModelContainerSpecAcceleratorRequirementsTypeEnum(v)
	}
	return betapb.VertexBetaModelContainerSpecAcceleratorRequirementsTypeEnum(0)
}

// ModelSupportedDeploymentResourcesTypesEnumToProto converts a ModelSupportedDeploymentResourcesTypesEnum enum to its proto representation.
func VertexBetaModelSupportedDeploymentResourcesTypesEnumToProto(e *beta.ModelSupportedDeploymentResourcesTypesEnum) betapb.VertexBetaModelSupportedDeploymentResourcesTypesEnum {
	if e == nil {
		return betapb.VertexBetaModelSupportedDeploymentResourcesTypesEnum(0)
	}
	if v, ok := betapb.VertexBetaModelSupportedDeploymentResourcesTypesEnum_value["ModelSupportedDeploymentResourcesTypesEnum"+string(*e)]; ok {
		return betapb.VertexBetaModelSupportedDeploymentResourcesTypesEnum(v)
	}
	return betapb.VertexBetaModelSupportedDeploymentResourcesTypesEnum(0)
}

// ModelSupportedExportFormatsToProto converts a ModelSupportedExportFormats object to its proto representation.
func VertexBetaModelSupportedExportFormatsToProto(o *beta.ModelSupportedExportFormats) *betapb.VertexBetaModelSupportedExportFormats {
	if o == nil {
		return nil
	}
	p := &betapb.VertexBetaModelSupportedExportFormats{}
	p.SetId(dcl.ValueOrEmptyString(o.Id))
	sExportableContents := make([]betapb.VertexBetaModelSupportedExportFormatsExportableContentsEnum, len(o.ExportableContents))
	for i, r := range o.ExportableContents {
		sExportableContents[i] = betapb.VertexBetaModelSupportedExportFormatsExportableContentsEnum(betapb.VertexBetaModelSupportedExportFormatsExportableContentsEnum_value[string(r)])
	}
	p.SetExportableContents(sExportableContents)
	return p
}

// ModelOriginalModelInfoToProto converts a ModelOriginalModelInfo object to its proto representation.
func VertexBetaModelOriginalModelInfoToProto(o *beta.ModelOriginalModelInfo) *betapb.VertexBetaModelOriginalModelInfo {
	if o == nil {
		return nil
	}
	p := &betapb.VertexBetaModelOriginalModelInfo{}
	p.SetModel(dcl.ValueOrEmptyString(o.Model))
	return p
}

// ModelContainerSpecToProto converts a ModelContainerSpec object to its proto representation.
func VertexBetaModelContainerSpecToProto(o *beta.ModelContainerSpec) *betapb.VertexBetaModelContainerSpec {
	if o == nil {
		return nil
	}
	p := &betapb.VertexBetaModelContainerSpec{}
	p.SetImageUri(dcl.ValueOrEmptyString(o.ImageUri))
	p.SetPredictRoute(dcl.ValueOrEmptyString(o.PredictRoute))
	p.SetHealthRoute(dcl.ValueOrEmptyString(o.HealthRoute))
	sCommand := make([]string, len(o.Command))
	for i, r := range o.Command {
		sCommand[i] = r
	}
	p.SetCommand(sCommand)
	sArgs := make([]string, len(o.Args))
	for i, r := range o.Args {
		sArgs[i] = r
	}
	p.SetArgs(sArgs)
	sEnv := make([]*betapb.VertexBetaModelContainerSpecEnv, len(o.Env))
	for i, r := range o.Env {
		sEnv[i] = VertexBetaModelContainerSpecEnvToProto(&r)
	}
	p.SetEnv(sEnv)
	sPorts := make([]*betapb.VertexBetaModelContainerSpecPorts, len(o.Ports))
	for i, r := range o.Ports {
		sPorts[i] = VertexBetaModelContainerSpecPortsToProto(&r)
	}
	p.SetPorts(sPorts)
	sAcceleratorRequirements := make([]*betapb.VertexBetaModelContainerSpecAcceleratorRequirements, len(o.AcceleratorRequirements))
	for i, r := range o.AcceleratorRequirements {
		sAcceleratorRequirements[i] = VertexBetaModelContainerSpecAcceleratorRequirementsToProto(&r)
	}
	p.SetAcceleratorRequirements(sAcceleratorRequirements)
	return p
}

// ModelContainerSpecEnvToProto converts a ModelContainerSpecEnv object to its proto representation.
func VertexBetaModelContainerSpecEnvToProto(o *beta.ModelContainerSpecEnv) *betapb.VertexBetaModelContainerSpecEnv {
	if o == nil {
		return nil
	}
	p := &betapb.VertexBetaModelContainerSpecEnv{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetValue(dcl.ValueOrEmptyString(o.Value))
	return p
}

// ModelContainerSpecPortsToProto converts a ModelContainerSpecPorts object to its proto representation.
func VertexBetaModelContainerSpecPortsToProto(o *beta.ModelContainerSpecPorts) *betapb.VertexBetaModelContainerSpecPorts {
	if o == nil {
		return nil
	}
	p := &betapb.VertexBetaModelContainerSpecPorts{}
	p.SetContainerPort(dcl.ValueOrEmptyInt64(o.ContainerPort))
	return p
}

// ModelContainerSpecAcceleratorRequirementsToProto converts a ModelContainerSpecAcceleratorRequirements object to its proto representation.
func VertexBetaModelContainerSpecAcceleratorRequirementsToProto(o *beta.ModelContainerSpecAcceleratorRequirements) *betapb.VertexBetaModelContainerSpecAcceleratorRequirements {
	if o == nil {
		return nil
	}
	p := &betapb.VertexBetaModelContainerSpecAcceleratorRequirements{}
	p.SetType(VertexBetaModelContainerSpecAcceleratorRequirementsTypeEnumToProto(o.Type))
	p.SetCount(dcl.ValueOrEmptyInt64(o.Count))
	return p
}

// ModelDeployedModelsToProto converts a ModelDeployedModels object to its proto representation.
func VertexBetaModelDeployedModelsToProto(o *beta.ModelDeployedModels) *betapb.VertexBetaModelDeployedModels {
	if o == nil {
		return nil
	}
	p := &betapb.VertexBetaModelDeployedModels{}
	p.SetEndpoint(dcl.ValueOrEmptyString(o.Endpoint))
	p.SetDeployedModelId(dcl.ValueOrEmptyString(o.DeployedModelId))
	return p
}

// ModelEncryptionSpecToProto converts a ModelEncryptionSpec object to its proto representation.
func VertexBetaModelEncryptionSpecToProto(o *beta.ModelEncryptionSpec) *betapb.VertexBetaModelEncryptionSpec {
	if o == nil {
		return nil
	}
	p := &betapb.VertexBetaModelEncryptionSpec{}
	p.SetKmsKeyName(dcl.ValueOrEmptyString(o.KmsKeyName))
	return p
}

// ModelToProto converts a Model resource to its proto representation.
func ModelToProto(resource *beta.Model) *betapb.VertexBetaModel {
	p := &betapb.VertexBetaModel{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetVersionId(dcl.ValueOrEmptyString(resource.VersionId))
	p.SetVersionCreateTime(dcl.ValueOrEmptyString(resource.VersionCreateTime))
	p.SetVersionUpdateTime(dcl.ValueOrEmptyString(resource.VersionUpdateTime))
	p.SetDisplayName(dcl.ValueOrEmptyString(resource.DisplayName))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetVersionDescription(dcl.ValueOrEmptyString(resource.VersionDescription))
	p.SetTrainingPipeline(dcl.ValueOrEmptyString(resource.TrainingPipeline))
	p.SetOriginalModelInfo(VertexBetaModelOriginalModelInfoToProto(resource.OriginalModelInfo))
	p.SetContainerSpec(VertexBetaModelContainerSpecToProto(resource.ContainerSpec))
	p.SetArtifactUri(dcl.ValueOrEmptyString(resource.ArtifactUri))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetEtag(dcl.ValueOrEmptyString(resource.Etag))
	p.SetEncryptionSpec(VertexBetaModelEncryptionSpecToProto(resource.EncryptionSpec))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	sVersionAliases := make([]string, len(resource.VersionAliases))
	for i, r := range resource.VersionAliases {
		sVersionAliases[i] = r
	}
	p.SetVersionAliases(sVersionAliases)
	sSupportedExportFormats := make([]*betapb.VertexBetaModelSupportedExportFormats, len(resource.SupportedExportFormats))
	for i, r := range resource.SupportedExportFormats {
		sSupportedExportFormats[i] = VertexBetaModelSupportedExportFormatsToProto(&r)
	}
	p.SetSupportedExportFormats(sSupportedExportFormats)
	sSupportedDeploymentResourcesTypes := make([]betapb.VertexBetaModelSupportedDeploymentResourcesTypesEnum, len(resource.SupportedDeploymentResourcesTypes))
	for i, r := range resource.SupportedDeploymentResourcesTypes {
		sSupportedDeploymentResourcesTypes[i] = betapb.VertexBetaModelSupportedDeploymentResourcesTypesEnum(betapb.VertexBetaModelSupportedDeploymentResourcesTypesEnum_value[string(r)])
	}
	p.SetSupportedDeploymentResourcesTypes(sSupportedDeploymentResourcesTypes)
	sSupportedInputStorageFormats := make([]string, len(resource.SupportedInputStorageFormats))
	for i, r := range resource.SupportedInputStorageFormats {
		sSupportedInputStorageFormats[i] = r
	}
	p.SetSupportedInputStorageFormats(sSupportedInputStorageFormats)
	sSupportedOutputStorageFormats := make([]string, len(resource.SupportedOutputStorageFormats))
	for i, r := range resource.SupportedOutputStorageFormats {
		sSupportedOutputStorageFormats[i] = r
	}
	p.SetSupportedOutputStorageFormats(sSupportedOutputStorageFormats)
	sDeployedModels := make([]*betapb.VertexBetaModelDeployedModels, len(resource.DeployedModels))
	for i, r := range resource.DeployedModels {
		sDeployedModels[i] = VertexBetaModelDeployedModelsToProto(&r)
	}
	p.SetDeployedModels(sDeployedModels)
	mLabels := make(map[string]string, len(resource.Labels))
	for k, r := range resource.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)

	return p
}

// applyModel handles the gRPC request by passing it to the underlying Model Apply() method.
func (s *ModelServer) applyModel(ctx context.Context, c *beta.Client, request *betapb.ApplyVertexBetaModelRequest) (*betapb.VertexBetaModel, error) {
	p := ProtoToModel(request.GetResource())
	res, err := c.ApplyModel(ctx, p)
	if err != nil {
		return nil, err
	}
	r := ModelToProto(res)
	return r, nil
}

// applyVertexBetaModel handles the gRPC request by passing it to the underlying Model Apply() method.
func (s *ModelServer) ApplyVertexBetaModel(ctx context.Context, request *betapb.ApplyVertexBetaModelRequest) (*betapb.VertexBetaModel, error) {
	cl, err := createConfigModel(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyModel(ctx, cl, request)
}

// DeleteModel handles the gRPC request by passing it to the underlying Model Delete() method.
func (s *ModelServer) DeleteVertexBetaModel(ctx context.Context, request *betapb.DeleteVertexBetaModelRequest) (*emptypb.Empty, error) {

	cl, err := createConfigModel(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteModel(ctx, ProtoToModel(request.GetResource()))

}

// ListVertexBetaModel handles the gRPC request by passing it to the underlying ModelList() method.
func (s *ModelServer) ListVertexBetaModel(ctx context.Context, request *betapb.ListVertexBetaModelRequest) (*betapb.ListVertexBetaModelResponse, error) {
	cl, err := createConfigModel(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListModel(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.VertexBetaModel
	for _, r := range resources.Items {
		rp := ModelToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListVertexBetaModelResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigModel(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
