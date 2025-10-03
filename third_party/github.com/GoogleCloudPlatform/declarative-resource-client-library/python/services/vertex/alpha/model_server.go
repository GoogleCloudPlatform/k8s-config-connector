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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/vertex/alpha/vertex_alpha_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/vertex/alpha"
)

// ModelServer implements the gRPC interface for Model.
type ModelServer struct{}

// ProtoToModelSupportedExportFormatsExportableContentsEnum converts a ModelSupportedExportFormatsExportableContentsEnum enum from its proto representation.
func ProtoToVertexAlphaModelSupportedExportFormatsExportableContentsEnum(e alphapb.VertexAlphaModelSupportedExportFormatsExportableContentsEnum) *alpha.ModelSupportedExportFormatsExportableContentsEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.VertexAlphaModelSupportedExportFormatsExportableContentsEnum_name[int32(e)]; ok {
		e := alpha.ModelSupportedExportFormatsExportableContentsEnum(n[len("VertexAlphaModelSupportedExportFormatsExportableContentsEnum"):])
		return &e
	}
	return nil
}

// ProtoToModelContainerSpecAcceleratorRequirementsTypeEnum converts a ModelContainerSpecAcceleratorRequirementsTypeEnum enum from its proto representation.
func ProtoToVertexAlphaModelContainerSpecAcceleratorRequirementsTypeEnum(e alphapb.VertexAlphaModelContainerSpecAcceleratorRequirementsTypeEnum) *alpha.ModelContainerSpecAcceleratorRequirementsTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.VertexAlphaModelContainerSpecAcceleratorRequirementsTypeEnum_name[int32(e)]; ok {
		e := alpha.ModelContainerSpecAcceleratorRequirementsTypeEnum(n[len("VertexAlphaModelContainerSpecAcceleratorRequirementsTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToModelSupportedDeploymentResourcesTypesEnum converts a ModelSupportedDeploymentResourcesTypesEnum enum from its proto representation.
func ProtoToVertexAlphaModelSupportedDeploymentResourcesTypesEnum(e alphapb.VertexAlphaModelSupportedDeploymentResourcesTypesEnum) *alpha.ModelSupportedDeploymentResourcesTypesEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.VertexAlphaModelSupportedDeploymentResourcesTypesEnum_name[int32(e)]; ok {
		e := alpha.ModelSupportedDeploymentResourcesTypesEnum(n[len("VertexAlphaModelSupportedDeploymentResourcesTypesEnum"):])
		return &e
	}
	return nil
}

// ProtoToModelSupportedExportFormats converts a ModelSupportedExportFormats object from its proto representation.
func ProtoToVertexAlphaModelSupportedExportFormats(p *alphapb.VertexAlphaModelSupportedExportFormats) *alpha.ModelSupportedExportFormats {
	if p == nil {
		return nil
	}
	obj := &alpha.ModelSupportedExportFormats{
		Id: dcl.StringOrNil(p.GetId()),
	}
	for _, r := range p.GetExportableContents() {
		obj.ExportableContents = append(obj.ExportableContents, *ProtoToVertexAlphaModelSupportedExportFormatsExportableContentsEnum(r))
	}
	return obj
}

// ProtoToModelOriginalModelInfo converts a ModelOriginalModelInfo object from its proto representation.
func ProtoToVertexAlphaModelOriginalModelInfo(p *alphapb.VertexAlphaModelOriginalModelInfo) *alpha.ModelOriginalModelInfo {
	if p == nil {
		return nil
	}
	obj := &alpha.ModelOriginalModelInfo{
		Model: dcl.StringOrNil(p.GetModel()),
	}
	return obj
}

// ProtoToModelContainerSpec converts a ModelContainerSpec object from its proto representation.
func ProtoToVertexAlphaModelContainerSpec(p *alphapb.VertexAlphaModelContainerSpec) *alpha.ModelContainerSpec {
	if p == nil {
		return nil
	}
	obj := &alpha.ModelContainerSpec{
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
		obj.Env = append(obj.Env, *ProtoToVertexAlphaModelContainerSpecEnv(r))
	}
	for _, r := range p.GetPorts() {
		obj.Ports = append(obj.Ports, *ProtoToVertexAlphaModelContainerSpecPorts(r))
	}
	for _, r := range p.GetAcceleratorRequirements() {
		obj.AcceleratorRequirements = append(obj.AcceleratorRequirements, *ProtoToVertexAlphaModelContainerSpecAcceleratorRequirements(r))
	}
	return obj
}

// ProtoToModelContainerSpecEnv converts a ModelContainerSpecEnv object from its proto representation.
func ProtoToVertexAlphaModelContainerSpecEnv(p *alphapb.VertexAlphaModelContainerSpecEnv) *alpha.ModelContainerSpecEnv {
	if p == nil {
		return nil
	}
	obj := &alpha.ModelContainerSpecEnv{
		Name:  dcl.StringOrNil(p.GetName()),
		Value: dcl.StringOrNil(p.GetValue()),
	}
	return obj
}

// ProtoToModelContainerSpecPorts converts a ModelContainerSpecPorts object from its proto representation.
func ProtoToVertexAlphaModelContainerSpecPorts(p *alphapb.VertexAlphaModelContainerSpecPorts) *alpha.ModelContainerSpecPorts {
	if p == nil {
		return nil
	}
	obj := &alpha.ModelContainerSpecPorts{
		ContainerPort: dcl.Int64OrNil(p.GetContainerPort()),
	}
	return obj
}

// ProtoToModelContainerSpecAcceleratorRequirements converts a ModelContainerSpecAcceleratorRequirements object from its proto representation.
func ProtoToVertexAlphaModelContainerSpecAcceleratorRequirements(p *alphapb.VertexAlphaModelContainerSpecAcceleratorRequirements) *alpha.ModelContainerSpecAcceleratorRequirements {
	if p == nil {
		return nil
	}
	obj := &alpha.ModelContainerSpecAcceleratorRequirements{
		Type:  ProtoToVertexAlphaModelContainerSpecAcceleratorRequirementsTypeEnum(p.GetType()),
		Count: dcl.Int64OrNil(p.GetCount()),
	}
	return obj
}

// ProtoToModelDeployedModels converts a ModelDeployedModels object from its proto representation.
func ProtoToVertexAlphaModelDeployedModels(p *alphapb.VertexAlphaModelDeployedModels) *alpha.ModelDeployedModels {
	if p == nil {
		return nil
	}
	obj := &alpha.ModelDeployedModels{
		Endpoint:        dcl.StringOrNil(p.GetEndpoint()),
		DeployedModelId: dcl.StringOrNil(p.GetDeployedModelId()),
	}
	return obj
}

// ProtoToModelEncryptionSpec converts a ModelEncryptionSpec object from its proto representation.
func ProtoToVertexAlphaModelEncryptionSpec(p *alphapb.VertexAlphaModelEncryptionSpec) *alpha.ModelEncryptionSpec {
	if p == nil {
		return nil
	}
	obj := &alpha.ModelEncryptionSpec{
		KmsKeyName: dcl.StringOrNil(p.GetKmsKeyName()),
	}
	return obj
}

// ProtoToModel converts a Model resource from its proto representation.
func ProtoToModel(p *alphapb.VertexAlphaModel) *alpha.Model {
	obj := &alpha.Model{
		Name:               dcl.StringOrNil(p.GetName()),
		VersionId:          dcl.StringOrNil(p.GetVersionId()),
		VersionCreateTime:  dcl.StringOrNil(p.GetVersionCreateTime()),
		VersionUpdateTime:  dcl.StringOrNil(p.GetVersionUpdateTime()),
		DisplayName:        dcl.StringOrNil(p.GetDisplayName()),
		Description:        dcl.StringOrNil(p.GetDescription()),
		VersionDescription: dcl.StringOrNil(p.GetVersionDescription()),
		TrainingPipeline:   dcl.StringOrNil(p.GetTrainingPipeline()),
		OriginalModelInfo:  ProtoToVertexAlphaModelOriginalModelInfo(p.GetOriginalModelInfo()),
		ContainerSpec:      ProtoToVertexAlphaModelContainerSpec(p.GetContainerSpec()),
		ArtifactUri:        dcl.StringOrNil(p.GetArtifactUri()),
		CreateTime:         dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:         dcl.StringOrNil(p.GetUpdateTime()),
		Etag:               dcl.StringOrNil(p.GetEtag()),
		EncryptionSpec:     ProtoToVertexAlphaModelEncryptionSpec(p.GetEncryptionSpec()),
		Project:            dcl.StringOrNil(p.GetProject()),
		Location:           dcl.StringOrNil(p.GetLocation()),
	}
	for _, r := range p.GetVersionAliases() {
		obj.VersionAliases = append(obj.VersionAliases, r)
	}
	for _, r := range p.GetSupportedExportFormats() {
		obj.SupportedExportFormats = append(obj.SupportedExportFormats, *ProtoToVertexAlphaModelSupportedExportFormats(r))
	}
	for _, r := range p.GetSupportedDeploymentResourcesTypes() {
		obj.SupportedDeploymentResourcesTypes = append(obj.SupportedDeploymentResourcesTypes, *ProtoToVertexAlphaModelSupportedDeploymentResourcesTypesEnum(r))
	}
	for _, r := range p.GetSupportedInputStorageFormats() {
		obj.SupportedInputStorageFormats = append(obj.SupportedInputStorageFormats, r)
	}
	for _, r := range p.GetSupportedOutputStorageFormats() {
		obj.SupportedOutputStorageFormats = append(obj.SupportedOutputStorageFormats, r)
	}
	for _, r := range p.GetDeployedModels() {
		obj.DeployedModels = append(obj.DeployedModels, *ProtoToVertexAlphaModelDeployedModels(r))
	}
	return obj
}

// ModelSupportedExportFormatsExportableContentsEnumToProto converts a ModelSupportedExportFormatsExportableContentsEnum enum to its proto representation.
func VertexAlphaModelSupportedExportFormatsExportableContentsEnumToProto(e *alpha.ModelSupportedExportFormatsExportableContentsEnum) alphapb.VertexAlphaModelSupportedExportFormatsExportableContentsEnum {
	if e == nil {
		return alphapb.VertexAlphaModelSupportedExportFormatsExportableContentsEnum(0)
	}
	if v, ok := alphapb.VertexAlphaModelSupportedExportFormatsExportableContentsEnum_value["ModelSupportedExportFormatsExportableContentsEnum"+string(*e)]; ok {
		return alphapb.VertexAlphaModelSupportedExportFormatsExportableContentsEnum(v)
	}
	return alphapb.VertexAlphaModelSupportedExportFormatsExportableContentsEnum(0)
}

// ModelContainerSpecAcceleratorRequirementsTypeEnumToProto converts a ModelContainerSpecAcceleratorRequirementsTypeEnum enum to its proto representation.
func VertexAlphaModelContainerSpecAcceleratorRequirementsTypeEnumToProto(e *alpha.ModelContainerSpecAcceleratorRequirementsTypeEnum) alphapb.VertexAlphaModelContainerSpecAcceleratorRequirementsTypeEnum {
	if e == nil {
		return alphapb.VertexAlphaModelContainerSpecAcceleratorRequirementsTypeEnum(0)
	}
	if v, ok := alphapb.VertexAlphaModelContainerSpecAcceleratorRequirementsTypeEnum_value["ModelContainerSpecAcceleratorRequirementsTypeEnum"+string(*e)]; ok {
		return alphapb.VertexAlphaModelContainerSpecAcceleratorRequirementsTypeEnum(v)
	}
	return alphapb.VertexAlphaModelContainerSpecAcceleratorRequirementsTypeEnum(0)
}

// ModelSupportedDeploymentResourcesTypesEnumToProto converts a ModelSupportedDeploymentResourcesTypesEnum enum to its proto representation.
func VertexAlphaModelSupportedDeploymentResourcesTypesEnumToProto(e *alpha.ModelSupportedDeploymentResourcesTypesEnum) alphapb.VertexAlphaModelSupportedDeploymentResourcesTypesEnum {
	if e == nil {
		return alphapb.VertexAlphaModelSupportedDeploymentResourcesTypesEnum(0)
	}
	if v, ok := alphapb.VertexAlphaModelSupportedDeploymentResourcesTypesEnum_value["ModelSupportedDeploymentResourcesTypesEnum"+string(*e)]; ok {
		return alphapb.VertexAlphaModelSupportedDeploymentResourcesTypesEnum(v)
	}
	return alphapb.VertexAlphaModelSupportedDeploymentResourcesTypesEnum(0)
}

// ModelSupportedExportFormatsToProto converts a ModelSupportedExportFormats object to its proto representation.
func VertexAlphaModelSupportedExportFormatsToProto(o *alpha.ModelSupportedExportFormats) *alphapb.VertexAlphaModelSupportedExportFormats {
	if o == nil {
		return nil
	}
	p := &alphapb.VertexAlphaModelSupportedExportFormats{}
	p.SetId(dcl.ValueOrEmptyString(o.Id))
	sExportableContents := make([]alphapb.VertexAlphaModelSupportedExportFormatsExportableContentsEnum, len(o.ExportableContents))
	for i, r := range o.ExportableContents {
		sExportableContents[i] = alphapb.VertexAlphaModelSupportedExportFormatsExportableContentsEnum(alphapb.VertexAlphaModelSupportedExportFormatsExportableContentsEnum_value[string(r)])
	}
	p.SetExportableContents(sExportableContents)
	return p
}

// ModelOriginalModelInfoToProto converts a ModelOriginalModelInfo object to its proto representation.
func VertexAlphaModelOriginalModelInfoToProto(o *alpha.ModelOriginalModelInfo) *alphapb.VertexAlphaModelOriginalModelInfo {
	if o == nil {
		return nil
	}
	p := &alphapb.VertexAlphaModelOriginalModelInfo{}
	p.SetModel(dcl.ValueOrEmptyString(o.Model))
	return p
}

// ModelContainerSpecToProto converts a ModelContainerSpec object to its proto representation.
func VertexAlphaModelContainerSpecToProto(o *alpha.ModelContainerSpec) *alphapb.VertexAlphaModelContainerSpec {
	if o == nil {
		return nil
	}
	p := &alphapb.VertexAlphaModelContainerSpec{}
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
	sEnv := make([]*alphapb.VertexAlphaModelContainerSpecEnv, len(o.Env))
	for i, r := range o.Env {
		sEnv[i] = VertexAlphaModelContainerSpecEnvToProto(&r)
	}
	p.SetEnv(sEnv)
	sPorts := make([]*alphapb.VertexAlphaModelContainerSpecPorts, len(o.Ports))
	for i, r := range o.Ports {
		sPorts[i] = VertexAlphaModelContainerSpecPortsToProto(&r)
	}
	p.SetPorts(sPorts)
	sAcceleratorRequirements := make([]*alphapb.VertexAlphaModelContainerSpecAcceleratorRequirements, len(o.AcceleratorRequirements))
	for i, r := range o.AcceleratorRequirements {
		sAcceleratorRequirements[i] = VertexAlphaModelContainerSpecAcceleratorRequirementsToProto(&r)
	}
	p.SetAcceleratorRequirements(sAcceleratorRequirements)
	return p
}

// ModelContainerSpecEnvToProto converts a ModelContainerSpecEnv object to its proto representation.
func VertexAlphaModelContainerSpecEnvToProto(o *alpha.ModelContainerSpecEnv) *alphapb.VertexAlphaModelContainerSpecEnv {
	if o == nil {
		return nil
	}
	p := &alphapb.VertexAlphaModelContainerSpecEnv{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetValue(dcl.ValueOrEmptyString(o.Value))
	return p
}

// ModelContainerSpecPortsToProto converts a ModelContainerSpecPorts object to its proto representation.
func VertexAlphaModelContainerSpecPortsToProto(o *alpha.ModelContainerSpecPorts) *alphapb.VertexAlphaModelContainerSpecPorts {
	if o == nil {
		return nil
	}
	p := &alphapb.VertexAlphaModelContainerSpecPorts{}
	p.SetContainerPort(dcl.ValueOrEmptyInt64(o.ContainerPort))
	return p
}

// ModelContainerSpecAcceleratorRequirementsToProto converts a ModelContainerSpecAcceleratorRequirements object to its proto representation.
func VertexAlphaModelContainerSpecAcceleratorRequirementsToProto(o *alpha.ModelContainerSpecAcceleratorRequirements) *alphapb.VertexAlphaModelContainerSpecAcceleratorRequirements {
	if o == nil {
		return nil
	}
	p := &alphapb.VertexAlphaModelContainerSpecAcceleratorRequirements{}
	p.SetType(VertexAlphaModelContainerSpecAcceleratorRequirementsTypeEnumToProto(o.Type))
	p.SetCount(dcl.ValueOrEmptyInt64(o.Count))
	return p
}

// ModelDeployedModelsToProto converts a ModelDeployedModels object to its proto representation.
func VertexAlphaModelDeployedModelsToProto(o *alpha.ModelDeployedModels) *alphapb.VertexAlphaModelDeployedModels {
	if o == nil {
		return nil
	}
	p := &alphapb.VertexAlphaModelDeployedModels{}
	p.SetEndpoint(dcl.ValueOrEmptyString(o.Endpoint))
	p.SetDeployedModelId(dcl.ValueOrEmptyString(o.DeployedModelId))
	return p
}

// ModelEncryptionSpecToProto converts a ModelEncryptionSpec object to its proto representation.
func VertexAlphaModelEncryptionSpecToProto(o *alpha.ModelEncryptionSpec) *alphapb.VertexAlphaModelEncryptionSpec {
	if o == nil {
		return nil
	}
	p := &alphapb.VertexAlphaModelEncryptionSpec{}
	p.SetKmsKeyName(dcl.ValueOrEmptyString(o.KmsKeyName))
	return p
}

// ModelToProto converts a Model resource to its proto representation.
func ModelToProto(resource *alpha.Model) *alphapb.VertexAlphaModel {
	p := &alphapb.VertexAlphaModel{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetVersionId(dcl.ValueOrEmptyString(resource.VersionId))
	p.SetVersionCreateTime(dcl.ValueOrEmptyString(resource.VersionCreateTime))
	p.SetVersionUpdateTime(dcl.ValueOrEmptyString(resource.VersionUpdateTime))
	p.SetDisplayName(dcl.ValueOrEmptyString(resource.DisplayName))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetVersionDescription(dcl.ValueOrEmptyString(resource.VersionDescription))
	p.SetTrainingPipeline(dcl.ValueOrEmptyString(resource.TrainingPipeline))
	p.SetOriginalModelInfo(VertexAlphaModelOriginalModelInfoToProto(resource.OriginalModelInfo))
	p.SetContainerSpec(VertexAlphaModelContainerSpecToProto(resource.ContainerSpec))
	p.SetArtifactUri(dcl.ValueOrEmptyString(resource.ArtifactUri))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetEtag(dcl.ValueOrEmptyString(resource.Etag))
	p.SetEncryptionSpec(VertexAlphaModelEncryptionSpecToProto(resource.EncryptionSpec))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	sVersionAliases := make([]string, len(resource.VersionAliases))
	for i, r := range resource.VersionAliases {
		sVersionAliases[i] = r
	}
	p.SetVersionAliases(sVersionAliases)
	sSupportedExportFormats := make([]*alphapb.VertexAlphaModelSupportedExportFormats, len(resource.SupportedExportFormats))
	for i, r := range resource.SupportedExportFormats {
		sSupportedExportFormats[i] = VertexAlphaModelSupportedExportFormatsToProto(&r)
	}
	p.SetSupportedExportFormats(sSupportedExportFormats)
	sSupportedDeploymentResourcesTypes := make([]alphapb.VertexAlphaModelSupportedDeploymentResourcesTypesEnum, len(resource.SupportedDeploymentResourcesTypes))
	for i, r := range resource.SupportedDeploymentResourcesTypes {
		sSupportedDeploymentResourcesTypes[i] = alphapb.VertexAlphaModelSupportedDeploymentResourcesTypesEnum(alphapb.VertexAlphaModelSupportedDeploymentResourcesTypesEnum_value[string(r)])
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
	sDeployedModels := make([]*alphapb.VertexAlphaModelDeployedModels, len(resource.DeployedModels))
	for i, r := range resource.DeployedModels {
		sDeployedModels[i] = VertexAlphaModelDeployedModelsToProto(&r)
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
func (s *ModelServer) applyModel(ctx context.Context, c *alpha.Client, request *alphapb.ApplyVertexAlphaModelRequest) (*alphapb.VertexAlphaModel, error) {
	p := ProtoToModel(request.GetResource())
	res, err := c.ApplyModel(ctx, p)
	if err != nil {
		return nil, err
	}
	r := ModelToProto(res)
	return r, nil
}

// applyVertexAlphaModel handles the gRPC request by passing it to the underlying Model Apply() method.
func (s *ModelServer) ApplyVertexAlphaModel(ctx context.Context, request *alphapb.ApplyVertexAlphaModelRequest) (*alphapb.VertexAlphaModel, error) {
	cl, err := createConfigModel(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyModel(ctx, cl, request)
}

// DeleteModel handles the gRPC request by passing it to the underlying Model Delete() method.
func (s *ModelServer) DeleteVertexAlphaModel(ctx context.Context, request *alphapb.DeleteVertexAlphaModelRequest) (*emptypb.Empty, error) {

	cl, err := createConfigModel(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteModel(ctx, ProtoToModel(request.GetResource()))

}

// ListVertexAlphaModel handles the gRPC request by passing it to the underlying ModelList() method.
func (s *ModelServer) ListVertexAlphaModel(ctx context.Context, request *alphapb.ListVertexAlphaModelRequest) (*alphapb.ListVertexAlphaModelResponse, error) {
	cl, err := createConfigModel(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListModel(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.VertexAlphaModel
	for _, r := range resources.Items {
		rp := ModelToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListVertexAlphaModelResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigModel(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
