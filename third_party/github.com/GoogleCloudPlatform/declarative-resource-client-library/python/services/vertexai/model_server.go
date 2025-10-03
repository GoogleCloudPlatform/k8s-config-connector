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
	vertexaipb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/vertexai/vertexai_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/vertexai"
)

// ModelServer implements the gRPC interface for Model.
type ModelServer struct{}

// ProtoToModelSupportedExportFormatsExportableContentsEnum converts a ModelSupportedExportFormatsExportableContentsEnum enum from its proto representation.
func ProtoToVertexaiModelSupportedExportFormatsExportableContentsEnum(e vertexaipb.VertexaiModelSupportedExportFormatsExportableContentsEnum) *vertexai.ModelSupportedExportFormatsExportableContentsEnum {
	if e == 0 {
		return nil
	}
	if n, ok := vertexaipb.VertexaiModelSupportedExportFormatsExportableContentsEnum_name[int32(e)]; ok {
		e := vertexai.ModelSupportedExportFormatsExportableContentsEnum(n[len("VertexaiModelSupportedExportFormatsExportableContentsEnum"):])
		return &e
	}
	return nil
}

// ProtoToModelSupportedDeploymentResourcesTypesEnum converts a ModelSupportedDeploymentResourcesTypesEnum enum from its proto representation.
func ProtoToVertexaiModelSupportedDeploymentResourcesTypesEnum(e vertexaipb.VertexaiModelSupportedDeploymentResourcesTypesEnum) *vertexai.ModelSupportedDeploymentResourcesTypesEnum {
	if e == 0 {
		return nil
	}
	if n, ok := vertexaipb.VertexaiModelSupportedDeploymentResourcesTypesEnum_name[int32(e)]; ok {
		e := vertexai.ModelSupportedDeploymentResourcesTypesEnum(n[len("VertexaiModelSupportedDeploymentResourcesTypesEnum"):])
		return &e
	}
	return nil
}

// ProtoToModelSupportedExportFormats converts a ModelSupportedExportFormats object from its proto representation.
func ProtoToVertexaiModelSupportedExportFormats(p *vertexaipb.VertexaiModelSupportedExportFormats) *vertexai.ModelSupportedExportFormats {
	if p == nil {
		return nil
	}
	obj := &vertexai.ModelSupportedExportFormats{
		Id: dcl.StringOrNil(p.GetId()),
	}
	for _, r := range p.GetExportableContents() {
		obj.ExportableContents = append(obj.ExportableContents, *ProtoToVertexaiModelSupportedExportFormatsExportableContentsEnum(r))
	}
	return obj
}

// ProtoToModelOriginalModelInfo converts a ModelOriginalModelInfo object from its proto representation.
func ProtoToVertexaiModelOriginalModelInfo(p *vertexaipb.VertexaiModelOriginalModelInfo) *vertexai.ModelOriginalModelInfo {
	if p == nil {
		return nil
	}
	obj := &vertexai.ModelOriginalModelInfo{
		Model: dcl.StringOrNil(p.GetModel()),
	}
	return obj
}

// ProtoToModelContainerSpec converts a ModelContainerSpec object from its proto representation.
func ProtoToVertexaiModelContainerSpec(p *vertexaipb.VertexaiModelContainerSpec) *vertexai.ModelContainerSpec {
	if p == nil {
		return nil
	}
	obj := &vertexai.ModelContainerSpec{
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
		obj.Env = append(obj.Env, *ProtoToVertexaiModelContainerSpecEnv(r))
	}
	for _, r := range p.GetPorts() {
		obj.Ports = append(obj.Ports, *ProtoToVertexaiModelContainerSpecPorts(r))
	}
	return obj
}

// ProtoToModelContainerSpecEnv converts a ModelContainerSpecEnv object from its proto representation.
func ProtoToVertexaiModelContainerSpecEnv(p *vertexaipb.VertexaiModelContainerSpecEnv) *vertexai.ModelContainerSpecEnv {
	if p == nil {
		return nil
	}
	obj := &vertexai.ModelContainerSpecEnv{
		Name:  dcl.StringOrNil(p.GetName()),
		Value: dcl.StringOrNil(p.GetValue()),
	}
	return obj
}

// ProtoToModelContainerSpecPorts converts a ModelContainerSpecPorts object from its proto representation.
func ProtoToVertexaiModelContainerSpecPorts(p *vertexaipb.VertexaiModelContainerSpecPorts) *vertexai.ModelContainerSpecPorts {
	if p == nil {
		return nil
	}
	obj := &vertexai.ModelContainerSpecPorts{
		ContainerPort: dcl.Int64OrNil(p.GetContainerPort()),
	}
	return obj
}

// ProtoToModelDeployedModels converts a ModelDeployedModels object from its proto representation.
func ProtoToVertexaiModelDeployedModels(p *vertexaipb.VertexaiModelDeployedModels) *vertexai.ModelDeployedModels {
	if p == nil {
		return nil
	}
	obj := &vertexai.ModelDeployedModels{
		Endpoint:        dcl.StringOrNil(p.GetEndpoint()),
		DeployedModelId: dcl.StringOrNil(p.GetDeployedModelId()),
	}
	return obj
}

// ProtoToModelEncryptionSpec converts a ModelEncryptionSpec object from its proto representation.
func ProtoToVertexaiModelEncryptionSpec(p *vertexaipb.VertexaiModelEncryptionSpec) *vertexai.ModelEncryptionSpec {
	if p == nil {
		return nil
	}
	obj := &vertexai.ModelEncryptionSpec{
		KmsKeyName: dcl.StringOrNil(p.GetKmsKeyName()),
	}
	return obj
}

// ProtoToModel converts a Model resource from its proto representation.
func ProtoToModel(p *vertexaipb.VertexaiModel) *vertexai.Model {
	obj := &vertexai.Model{
		Name:               dcl.StringOrNil(p.GetName()),
		VersionId:          dcl.StringOrNil(p.GetVersionId()),
		VersionCreateTime:  dcl.StringOrNil(p.GetVersionCreateTime()),
		VersionUpdateTime:  dcl.StringOrNil(p.GetVersionUpdateTime()),
		DisplayName:        dcl.StringOrNil(p.GetDisplayName()),
		Description:        dcl.StringOrNil(p.GetDescription()),
		VersionDescription: dcl.StringOrNil(p.GetVersionDescription()),
		TrainingPipeline:   dcl.StringOrNil(p.GetTrainingPipeline()),
		OriginalModelInfo:  ProtoToVertexaiModelOriginalModelInfo(p.GetOriginalModelInfo()),
		ContainerSpec:      ProtoToVertexaiModelContainerSpec(p.GetContainerSpec()),
		ArtifactUri:        dcl.StringOrNil(p.GetArtifactUri()),
		CreateTime:         dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:         dcl.StringOrNil(p.GetUpdateTime()),
		Etag:               dcl.StringOrNil(p.GetEtag()),
		EncryptionSpec:     ProtoToVertexaiModelEncryptionSpec(p.GetEncryptionSpec()),
		Project:            dcl.StringOrNil(p.GetProject()),
		Location:           dcl.StringOrNil(p.GetLocation()),
	}
	for _, r := range p.GetSupportedExportFormats() {
		obj.SupportedExportFormats = append(obj.SupportedExportFormats, *ProtoToVertexaiModelSupportedExportFormats(r))
	}
	for _, r := range p.GetSupportedDeploymentResourcesTypes() {
		obj.SupportedDeploymentResourcesTypes = append(obj.SupportedDeploymentResourcesTypes, *ProtoToVertexaiModelSupportedDeploymentResourcesTypesEnum(r))
	}
	for _, r := range p.GetSupportedInputStorageFormats() {
		obj.SupportedInputStorageFormats = append(obj.SupportedInputStorageFormats, r)
	}
	for _, r := range p.GetSupportedOutputStorageFormats() {
		obj.SupportedOutputStorageFormats = append(obj.SupportedOutputStorageFormats, r)
	}
	for _, r := range p.GetDeployedModels() {
		obj.DeployedModels = append(obj.DeployedModels, *ProtoToVertexaiModelDeployedModels(r))
	}
	return obj
}

// ModelSupportedExportFormatsExportableContentsEnumToProto converts a ModelSupportedExportFormatsExportableContentsEnum enum to its proto representation.
func VertexaiModelSupportedExportFormatsExportableContentsEnumToProto(e *vertexai.ModelSupportedExportFormatsExportableContentsEnum) vertexaipb.VertexaiModelSupportedExportFormatsExportableContentsEnum {
	if e == nil {
		return vertexaipb.VertexaiModelSupportedExportFormatsExportableContentsEnum(0)
	}
	if v, ok := vertexaipb.VertexaiModelSupportedExportFormatsExportableContentsEnum_value["ModelSupportedExportFormatsExportableContentsEnum"+string(*e)]; ok {
		return vertexaipb.VertexaiModelSupportedExportFormatsExportableContentsEnum(v)
	}
	return vertexaipb.VertexaiModelSupportedExportFormatsExportableContentsEnum(0)
}

// ModelSupportedDeploymentResourcesTypesEnumToProto converts a ModelSupportedDeploymentResourcesTypesEnum enum to its proto representation.
func VertexaiModelSupportedDeploymentResourcesTypesEnumToProto(e *vertexai.ModelSupportedDeploymentResourcesTypesEnum) vertexaipb.VertexaiModelSupportedDeploymentResourcesTypesEnum {
	if e == nil {
		return vertexaipb.VertexaiModelSupportedDeploymentResourcesTypesEnum(0)
	}
	if v, ok := vertexaipb.VertexaiModelSupportedDeploymentResourcesTypesEnum_value["ModelSupportedDeploymentResourcesTypesEnum"+string(*e)]; ok {
		return vertexaipb.VertexaiModelSupportedDeploymentResourcesTypesEnum(v)
	}
	return vertexaipb.VertexaiModelSupportedDeploymentResourcesTypesEnum(0)
}

// ModelSupportedExportFormatsToProto converts a ModelSupportedExportFormats object to its proto representation.
func VertexaiModelSupportedExportFormatsToProto(o *vertexai.ModelSupportedExportFormats) *vertexaipb.VertexaiModelSupportedExportFormats {
	if o == nil {
		return nil
	}
	p := &vertexaipb.VertexaiModelSupportedExportFormats{}
	p.SetId(dcl.ValueOrEmptyString(o.Id))
	sExportableContents := make([]vertexaipb.VertexaiModelSupportedExportFormatsExportableContentsEnum, len(o.ExportableContents))
	for i, r := range o.ExportableContents {
		sExportableContents[i] = vertexaipb.VertexaiModelSupportedExportFormatsExportableContentsEnum(vertexaipb.VertexaiModelSupportedExportFormatsExportableContentsEnum_value[string(r)])
	}
	p.SetExportableContents(sExportableContents)
	return p
}

// ModelOriginalModelInfoToProto converts a ModelOriginalModelInfo object to its proto representation.
func VertexaiModelOriginalModelInfoToProto(o *vertexai.ModelOriginalModelInfo) *vertexaipb.VertexaiModelOriginalModelInfo {
	if o == nil {
		return nil
	}
	p := &vertexaipb.VertexaiModelOriginalModelInfo{}
	p.SetModel(dcl.ValueOrEmptyString(o.Model))
	return p
}

// ModelContainerSpecToProto converts a ModelContainerSpec object to its proto representation.
func VertexaiModelContainerSpecToProto(o *vertexai.ModelContainerSpec) *vertexaipb.VertexaiModelContainerSpec {
	if o == nil {
		return nil
	}
	p := &vertexaipb.VertexaiModelContainerSpec{}
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
	sEnv := make([]*vertexaipb.VertexaiModelContainerSpecEnv, len(o.Env))
	for i, r := range o.Env {
		sEnv[i] = VertexaiModelContainerSpecEnvToProto(&r)
	}
	p.SetEnv(sEnv)
	sPorts := make([]*vertexaipb.VertexaiModelContainerSpecPorts, len(o.Ports))
	for i, r := range o.Ports {
		sPorts[i] = VertexaiModelContainerSpecPortsToProto(&r)
	}
	p.SetPorts(sPorts)
	return p
}

// ModelContainerSpecEnvToProto converts a ModelContainerSpecEnv object to its proto representation.
func VertexaiModelContainerSpecEnvToProto(o *vertexai.ModelContainerSpecEnv) *vertexaipb.VertexaiModelContainerSpecEnv {
	if o == nil {
		return nil
	}
	p := &vertexaipb.VertexaiModelContainerSpecEnv{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetValue(dcl.ValueOrEmptyString(o.Value))
	return p
}

// ModelContainerSpecPortsToProto converts a ModelContainerSpecPorts object to its proto representation.
func VertexaiModelContainerSpecPortsToProto(o *vertexai.ModelContainerSpecPorts) *vertexaipb.VertexaiModelContainerSpecPorts {
	if o == nil {
		return nil
	}
	p := &vertexaipb.VertexaiModelContainerSpecPorts{}
	p.SetContainerPort(dcl.ValueOrEmptyInt64(o.ContainerPort))
	return p
}

// ModelDeployedModelsToProto converts a ModelDeployedModels object to its proto representation.
func VertexaiModelDeployedModelsToProto(o *vertexai.ModelDeployedModels) *vertexaipb.VertexaiModelDeployedModels {
	if o == nil {
		return nil
	}
	p := &vertexaipb.VertexaiModelDeployedModels{}
	p.SetEndpoint(dcl.ValueOrEmptyString(o.Endpoint))
	p.SetDeployedModelId(dcl.ValueOrEmptyString(o.DeployedModelId))
	return p
}

// ModelEncryptionSpecToProto converts a ModelEncryptionSpec object to its proto representation.
func VertexaiModelEncryptionSpecToProto(o *vertexai.ModelEncryptionSpec) *vertexaipb.VertexaiModelEncryptionSpec {
	if o == nil {
		return nil
	}
	p := &vertexaipb.VertexaiModelEncryptionSpec{}
	p.SetKmsKeyName(dcl.ValueOrEmptyString(o.KmsKeyName))
	return p
}

// ModelToProto converts a Model resource to its proto representation.
func ModelToProto(resource *vertexai.Model) *vertexaipb.VertexaiModel {
	p := &vertexaipb.VertexaiModel{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetVersionId(dcl.ValueOrEmptyString(resource.VersionId))
	p.SetVersionCreateTime(dcl.ValueOrEmptyString(resource.VersionCreateTime))
	p.SetVersionUpdateTime(dcl.ValueOrEmptyString(resource.VersionUpdateTime))
	p.SetDisplayName(dcl.ValueOrEmptyString(resource.DisplayName))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetVersionDescription(dcl.ValueOrEmptyString(resource.VersionDescription))
	p.SetTrainingPipeline(dcl.ValueOrEmptyString(resource.TrainingPipeline))
	p.SetOriginalModelInfo(VertexaiModelOriginalModelInfoToProto(resource.OriginalModelInfo))
	p.SetContainerSpec(VertexaiModelContainerSpecToProto(resource.ContainerSpec))
	p.SetArtifactUri(dcl.ValueOrEmptyString(resource.ArtifactUri))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetEtag(dcl.ValueOrEmptyString(resource.Etag))
	p.SetEncryptionSpec(VertexaiModelEncryptionSpecToProto(resource.EncryptionSpec))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	sSupportedExportFormats := make([]*vertexaipb.VertexaiModelSupportedExportFormats, len(resource.SupportedExportFormats))
	for i, r := range resource.SupportedExportFormats {
		sSupportedExportFormats[i] = VertexaiModelSupportedExportFormatsToProto(&r)
	}
	p.SetSupportedExportFormats(sSupportedExportFormats)
	sSupportedDeploymentResourcesTypes := make([]vertexaipb.VertexaiModelSupportedDeploymentResourcesTypesEnum, len(resource.SupportedDeploymentResourcesTypes))
	for i, r := range resource.SupportedDeploymentResourcesTypes {
		sSupportedDeploymentResourcesTypes[i] = vertexaipb.VertexaiModelSupportedDeploymentResourcesTypesEnum(vertexaipb.VertexaiModelSupportedDeploymentResourcesTypesEnum_value[string(r)])
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
	sDeployedModels := make([]*vertexaipb.VertexaiModelDeployedModels, len(resource.DeployedModels))
	for i, r := range resource.DeployedModels {
		sDeployedModels[i] = VertexaiModelDeployedModelsToProto(&r)
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
func (s *ModelServer) applyModel(ctx context.Context, c *vertexai.Client, request *vertexaipb.ApplyVertexaiModelRequest) (*vertexaipb.VertexaiModel, error) {
	p := ProtoToModel(request.GetResource())
	res, err := c.ApplyModel(ctx, p)
	if err != nil {
		return nil, err
	}
	r := ModelToProto(res)
	return r, nil
}

// applyVertexaiModel handles the gRPC request by passing it to the underlying Model Apply() method.
func (s *ModelServer) ApplyVertexaiModel(ctx context.Context, request *vertexaipb.ApplyVertexaiModelRequest) (*vertexaipb.VertexaiModel, error) {
	cl, err := createConfigModel(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyModel(ctx, cl, request)
}

// DeleteModel handles the gRPC request by passing it to the underlying Model Delete() method.
func (s *ModelServer) DeleteVertexaiModel(ctx context.Context, request *vertexaipb.DeleteVertexaiModelRequest) (*emptypb.Empty, error) {

	cl, err := createConfigModel(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteModel(ctx, ProtoToModel(request.GetResource()))

}

// ListVertexaiModel handles the gRPC request by passing it to the underlying ModelList() method.
func (s *ModelServer) ListVertexaiModel(ctx context.Context, request *vertexaipb.ListVertexaiModelRequest) (*vertexaipb.ListVertexaiModelResponse, error) {
	cl, err := createConfigModel(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListModel(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*vertexaipb.VertexaiModel
	for _, r := range resources.Items {
		rp := ModelToProto(r)
		protos = append(protos, rp)
	}
	p := &vertexaipb.ListVertexaiModelResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigModel(ctx context.Context, service_account_file string) (*vertexai.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return vertexai.NewClient(conf), nil
}
