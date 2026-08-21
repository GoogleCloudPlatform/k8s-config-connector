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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/vertexai/beta/vertexai_beta_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/vertexai/beta"
)

// ModelServer implements the gRPC interface for Model.
type ModelServer struct{}

// ProtoToModelSupportedExportFormatsExportableContentsEnum converts a ModelSupportedExportFormatsExportableContentsEnum enum from its proto representation.
func ProtoToVertexaiBetaModelSupportedExportFormatsExportableContentsEnum(e betapb.VertexaiBetaModelSupportedExportFormatsExportableContentsEnum) *beta.ModelSupportedExportFormatsExportableContentsEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.VertexaiBetaModelSupportedExportFormatsExportableContentsEnum_name[int32(e)]; ok {
		e := beta.ModelSupportedExportFormatsExportableContentsEnum(n[len("VertexaiBetaModelSupportedExportFormatsExportableContentsEnum"):])
		return &e
	}
	return nil
}

// ProtoToModelSupportedDeploymentResourcesTypesEnum converts a ModelSupportedDeploymentResourcesTypesEnum enum from its proto representation.
func ProtoToVertexaiBetaModelSupportedDeploymentResourcesTypesEnum(e betapb.VertexaiBetaModelSupportedDeploymentResourcesTypesEnum) *beta.ModelSupportedDeploymentResourcesTypesEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.VertexaiBetaModelSupportedDeploymentResourcesTypesEnum_name[int32(e)]; ok {
		e := beta.ModelSupportedDeploymentResourcesTypesEnum(n[len("VertexaiBetaModelSupportedDeploymentResourcesTypesEnum"):])
		return &e
	}
	return nil
}

// ProtoToModelSupportedExportFormats converts a ModelSupportedExportFormats object from its proto representation.
func ProtoToVertexaiBetaModelSupportedExportFormats(p *betapb.VertexaiBetaModelSupportedExportFormats) *beta.ModelSupportedExportFormats {
	if p == nil {
		return nil
	}
	obj := &beta.ModelSupportedExportFormats{
		Id: dcl.StringOrNil(p.GetId()),
	}
	for _, r := range p.GetExportableContents() {
		obj.ExportableContents = append(obj.ExportableContents, *ProtoToVertexaiBetaModelSupportedExportFormatsExportableContentsEnum(r))
	}
	return obj
}

// ProtoToModelOriginalModelInfo converts a ModelOriginalModelInfo object from its proto representation.
func ProtoToVertexaiBetaModelOriginalModelInfo(p *betapb.VertexaiBetaModelOriginalModelInfo) *beta.ModelOriginalModelInfo {
	if p == nil {
		return nil
	}
	obj := &beta.ModelOriginalModelInfo{
		Model: dcl.StringOrNil(p.GetModel()),
	}
	return obj
}

// ProtoToModelContainerSpec converts a ModelContainerSpec object from its proto representation.
func ProtoToVertexaiBetaModelContainerSpec(p *betapb.VertexaiBetaModelContainerSpec) *beta.ModelContainerSpec {
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
		obj.Env = append(obj.Env, *ProtoToVertexaiBetaModelContainerSpecEnv(r))
	}
	for _, r := range p.GetPorts() {
		obj.Ports = append(obj.Ports, *ProtoToVertexaiBetaModelContainerSpecPorts(r))
	}
	return obj
}

// ProtoToModelContainerSpecEnv converts a ModelContainerSpecEnv object from its proto representation.
func ProtoToVertexaiBetaModelContainerSpecEnv(p *betapb.VertexaiBetaModelContainerSpecEnv) *beta.ModelContainerSpecEnv {
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
func ProtoToVertexaiBetaModelContainerSpecPorts(p *betapb.VertexaiBetaModelContainerSpecPorts) *beta.ModelContainerSpecPorts {
	if p == nil {
		return nil
	}
	obj := &beta.ModelContainerSpecPorts{
		ContainerPort: dcl.Int64OrNil(p.GetContainerPort()),
	}
	return obj
}

// ProtoToModelDeployedModels converts a ModelDeployedModels object from its proto representation.
func ProtoToVertexaiBetaModelDeployedModels(p *betapb.VertexaiBetaModelDeployedModels) *beta.ModelDeployedModels {
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
func ProtoToVertexaiBetaModelEncryptionSpec(p *betapb.VertexaiBetaModelEncryptionSpec) *beta.ModelEncryptionSpec {
	if p == nil {
		return nil
	}
	obj := &beta.ModelEncryptionSpec{
		KmsKeyName: dcl.StringOrNil(p.GetKmsKeyName()),
	}
	return obj
}

// ProtoToModel converts a Model resource from its proto representation.
func ProtoToModel(p *betapb.VertexaiBetaModel) *beta.Model {
	obj := &beta.Model{
		Name:               dcl.StringOrNil(p.GetName()),
		VersionId:          dcl.StringOrNil(p.GetVersionId()),
		VersionCreateTime:  dcl.StringOrNil(p.GetVersionCreateTime()),
		VersionUpdateTime:  dcl.StringOrNil(p.GetVersionUpdateTime()),
		DisplayName:        dcl.StringOrNil(p.GetDisplayName()),
		Description:        dcl.StringOrNil(p.GetDescription()),
		VersionDescription: dcl.StringOrNil(p.GetVersionDescription()),
		TrainingPipeline:   dcl.StringOrNil(p.GetTrainingPipeline()),
		OriginalModelInfo:  ProtoToVertexaiBetaModelOriginalModelInfo(p.GetOriginalModelInfo()),
		ContainerSpec:      ProtoToVertexaiBetaModelContainerSpec(p.GetContainerSpec()),
		ArtifactUri:        dcl.StringOrNil(p.GetArtifactUri()),
		CreateTime:         dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:         dcl.StringOrNil(p.GetUpdateTime()),
		Etag:               dcl.StringOrNil(p.GetEtag()),
		EncryptionSpec:     ProtoToVertexaiBetaModelEncryptionSpec(p.GetEncryptionSpec()),
		Project:            dcl.StringOrNil(p.GetProject()),
		Location:           dcl.StringOrNil(p.GetLocation()),
	}
	for _, r := range p.GetVersionAliases() {
		obj.VersionAliases = append(obj.VersionAliases, r)
	}
	for _, r := range p.GetSupportedExportFormats() {
		obj.SupportedExportFormats = append(obj.SupportedExportFormats, *ProtoToVertexaiBetaModelSupportedExportFormats(r))
	}
	for _, r := range p.GetSupportedDeploymentResourcesTypes() {
		obj.SupportedDeploymentResourcesTypes = append(obj.SupportedDeploymentResourcesTypes, *ProtoToVertexaiBetaModelSupportedDeploymentResourcesTypesEnum(r))
	}
	for _, r := range p.GetSupportedInputStorageFormats() {
		obj.SupportedInputStorageFormats = append(obj.SupportedInputStorageFormats, r)
	}
	for _, r := range p.GetSupportedOutputStorageFormats() {
		obj.SupportedOutputStorageFormats = append(obj.SupportedOutputStorageFormats, r)
	}
	for _, r := range p.GetDeployedModels() {
		obj.DeployedModels = append(obj.DeployedModels, *ProtoToVertexaiBetaModelDeployedModels(r))
	}
	return obj
}

// ModelSupportedExportFormatsExportableContentsEnumToProto converts a ModelSupportedExportFormatsExportableContentsEnum enum to its proto representation.
func VertexaiBetaModelSupportedExportFormatsExportableContentsEnumToProto(e *beta.ModelSupportedExportFormatsExportableContentsEnum) betapb.VertexaiBetaModelSupportedExportFormatsExportableContentsEnum {
	if e == nil {
		return betapb.VertexaiBetaModelSupportedExportFormatsExportableContentsEnum(0)
	}
	if v, ok := betapb.VertexaiBetaModelSupportedExportFormatsExportableContentsEnum_value["ModelSupportedExportFormatsExportableContentsEnum"+string(*e)]; ok {
		return betapb.VertexaiBetaModelSupportedExportFormatsExportableContentsEnum(v)
	}
	return betapb.VertexaiBetaModelSupportedExportFormatsExportableContentsEnum(0)
}

// ModelSupportedDeploymentResourcesTypesEnumToProto converts a ModelSupportedDeploymentResourcesTypesEnum enum to its proto representation.
func VertexaiBetaModelSupportedDeploymentResourcesTypesEnumToProto(e *beta.ModelSupportedDeploymentResourcesTypesEnum) betapb.VertexaiBetaModelSupportedDeploymentResourcesTypesEnum {
	if e == nil {
		return betapb.VertexaiBetaModelSupportedDeploymentResourcesTypesEnum(0)
	}
	if v, ok := betapb.VertexaiBetaModelSupportedDeploymentResourcesTypesEnum_value["ModelSupportedDeploymentResourcesTypesEnum"+string(*e)]; ok {
		return betapb.VertexaiBetaModelSupportedDeploymentResourcesTypesEnum(v)
	}
	return betapb.VertexaiBetaModelSupportedDeploymentResourcesTypesEnum(0)
}

// ModelSupportedExportFormatsToProto converts a ModelSupportedExportFormats object to its proto representation.
func VertexaiBetaModelSupportedExportFormatsToProto(o *beta.ModelSupportedExportFormats) *betapb.VertexaiBetaModelSupportedExportFormats {
	if o == nil {
		return nil
	}
	p := &betapb.VertexaiBetaModelSupportedExportFormats{}
	p.SetId(dcl.ValueOrEmptyString(o.Id))
	sExportableContents := make([]betapb.VertexaiBetaModelSupportedExportFormatsExportableContentsEnum, len(o.ExportableContents))
	for i, r := range o.ExportableContents {
		sExportableContents[i] = betapb.VertexaiBetaModelSupportedExportFormatsExportableContentsEnum(betapb.VertexaiBetaModelSupportedExportFormatsExportableContentsEnum_value[string(r)])
	}
	p.SetExportableContents(sExportableContents)
	return p
}

// ModelOriginalModelInfoToProto converts a ModelOriginalModelInfo object to its proto representation.
func VertexaiBetaModelOriginalModelInfoToProto(o *beta.ModelOriginalModelInfo) *betapb.VertexaiBetaModelOriginalModelInfo {
	if o == nil {
		return nil
	}
	p := &betapb.VertexaiBetaModelOriginalModelInfo{}
	p.SetModel(dcl.ValueOrEmptyString(o.Model))
	return p
}

// ModelContainerSpecToProto converts a ModelContainerSpec object to its proto representation.
func VertexaiBetaModelContainerSpecToProto(o *beta.ModelContainerSpec) *betapb.VertexaiBetaModelContainerSpec {
	if o == nil {
		return nil
	}
	p := &betapb.VertexaiBetaModelContainerSpec{}
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
	sEnv := make([]*betapb.VertexaiBetaModelContainerSpecEnv, len(o.Env))
	for i, r := range o.Env {
		sEnv[i] = VertexaiBetaModelContainerSpecEnvToProto(&r)
	}
	p.SetEnv(sEnv)
	sPorts := make([]*betapb.VertexaiBetaModelContainerSpecPorts, len(o.Ports))
	for i, r := range o.Ports {
		sPorts[i] = VertexaiBetaModelContainerSpecPortsToProto(&r)
	}
	p.SetPorts(sPorts)
	return p
}

// ModelContainerSpecEnvToProto converts a ModelContainerSpecEnv object to its proto representation.
func VertexaiBetaModelContainerSpecEnvToProto(o *beta.ModelContainerSpecEnv) *betapb.VertexaiBetaModelContainerSpecEnv {
	if o == nil {
		return nil
	}
	p := &betapb.VertexaiBetaModelContainerSpecEnv{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetValue(dcl.ValueOrEmptyString(o.Value))
	return p
}

// ModelContainerSpecPortsToProto converts a ModelContainerSpecPorts object to its proto representation.
func VertexaiBetaModelContainerSpecPortsToProto(o *beta.ModelContainerSpecPorts) *betapb.VertexaiBetaModelContainerSpecPorts {
	if o == nil {
		return nil
	}
	p := &betapb.VertexaiBetaModelContainerSpecPorts{}
	p.SetContainerPort(dcl.ValueOrEmptyInt64(o.ContainerPort))
	return p
}

// ModelDeployedModelsToProto converts a ModelDeployedModels object to its proto representation.
func VertexaiBetaModelDeployedModelsToProto(o *beta.ModelDeployedModels) *betapb.VertexaiBetaModelDeployedModels {
	if o == nil {
		return nil
	}
	p := &betapb.VertexaiBetaModelDeployedModels{}
	p.SetEndpoint(dcl.ValueOrEmptyString(o.Endpoint))
	p.SetDeployedModelId(dcl.ValueOrEmptyString(o.DeployedModelId))
	return p
}

// ModelEncryptionSpecToProto converts a ModelEncryptionSpec object to its proto representation.
func VertexaiBetaModelEncryptionSpecToProto(o *beta.ModelEncryptionSpec) *betapb.VertexaiBetaModelEncryptionSpec {
	if o == nil {
		return nil
	}
	p := &betapb.VertexaiBetaModelEncryptionSpec{}
	p.SetKmsKeyName(dcl.ValueOrEmptyString(o.KmsKeyName))
	return p
}

// ModelToProto converts a Model resource to its proto representation.
func ModelToProto(resource *beta.Model) *betapb.VertexaiBetaModel {
	p := &betapb.VertexaiBetaModel{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetVersionId(dcl.ValueOrEmptyString(resource.VersionId))
	p.SetVersionCreateTime(dcl.ValueOrEmptyString(resource.VersionCreateTime))
	p.SetVersionUpdateTime(dcl.ValueOrEmptyString(resource.VersionUpdateTime))
	p.SetDisplayName(dcl.ValueOrEmptyString(resource.DisplayName))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetVersionDescription(dcl.ValueOrEmptyString(resource.VersionDescription))
	p.SetTrainingPipeline(dcl.ValueOrEmptyString(resource.TrainingPipeline))
	p.SetOriginalModelInfo(VertexaiBetaModelOriginalModelInfoToProto(resource.OriginalModelInfo))
	p.SetContainerSpec(VertexaiBetaModelContainerSpecToProto(resource.ContainerSpec))
	p.SetArtifactUri(dcl.ValueOrEmptyString(resource.ArtifactUri))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetEtag(dcl.ValueOrEmptyString(resource.Etag))
	p.SetEncryptionSpec(VertexaiBetaModelEncryptionSpecToProto(resource.EncryptionSpec))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	sVersionAliases := make([]string, len(resource.VersionAliases))
	for i, r := range resource.VersionAliases {
		sVersionAliases[i] = r
	}
	p.SetVersionAliases(sVersionAliases)
	sSupportedExportFormats := make([]*betapb.VertexaiBetaModelSupportedExportFormats, len(resource.SupportedExportFormats))
	for i, r := range resource.SupportedExportFormats {
		sSupportedExportFormats[i] = VertexaiBetaModelSupportedExportFormatsToProto(&r)
	}
	p.SetSupportedExportFormats(sSupportedExportFormats)
	sSupportedDeploymentResourcesTypes := make([]betapb.VertexaiBetaModelSupportedDeploymentResourcesTypesEnum, len(resource.SupportedDeploymentResourcesTypes))
	for i, r := range resource.SupportedDeploymentResourcesTypes {
		sSupportedDeploymentResourcesTypes[i] = betapb.VertexaiBetaModelSupportedDeploymentResourcesTypesEnum(betapb.VertexaiBetaModelSupportedDeploymentResourcesTypesEnum_value[string(r)])
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
	sDeployedModels := make([]*betapb.VertexaiBetaModelDeployedModels, len(resource.DeployedModels))
	for i, r := range resource.DeployedModels {
		sDeployedModels[i] = VertexaiBetaModelDeployedModelsToProto(&r)
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
func (s *ModelServer) applyModel(ctx context.Context, c *beta.Client, request *betapb.ApplyVertexaiBetaModelRequest) (*betapb.VertexaiBetaModel, error) {
	p := ProtoToModel(request.GetResource())
	res, err := c.ApplyModel(ctx, p)
	if err != nil {
		return nil, err
	}
	r := ModelToProto(res)
	return r, nil
}

// applyVertexaiBetaModel handles the gRPC request by passing it to the underlying Model Apply() method.
func (s *ModelServer) ApplyVertexaiBetaModel(ctx context.Context, request *betapb.ApplyVertexaiBetaModelRequest) (*betapb.VertexaiBetaModel, error) {
	cl, err := createConfigModel(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyModel(ctx, cl, request)
}

// DeleteModel handles the gRPC request by passing it to the underlying Model Delete() method.
func (s *ModelServer) DeleteVertexaiBetaModel(ctx context.Context, request *betapb.DeleteVertexaiBetaModelRequest) (*emptypb.Empty, error) {

	cl, err := createConfigModel(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteModel(ctx, ProtoToModel(request.GetResource()))

}

// ListVertexaiBetaModel handles the gRPC request by passing it to the underlying ModelList() method.
func (s *ModelServer) ListVertexaiBetaModel(ctx context.Context, request *betapb.ListVertexaiBetaModelRequest) (*betapb.ListVertexaiBetaModelResponse, error) {
	cl, err := createConfigModel(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListModel(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.VertexaiBetaModel
	for _, r := range resources.Items {
		rp := ModelToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListVertexaiBetaModelResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigModel(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
