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
	vertexpb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/vertex/vertex_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/vertex"
)

// ModelServer implements the gRPC interface for Model.
type ModelServer struct{}

// ProtoToModelSupportedExportFormatsExportableContentsEnum converts a ModelSupportedExportFormatsExportableContentsEnum enum from its proto representation.
func ProtoToVertexModelSupportedExportFormatsExportableContentsEnum(e vertexpb.VertexModelSupportedExportFormatsExportableContentsEnum) *vertex.ModelSupportedExportFormatsExportableContentsEnum {
	if e == 0 {
		return nil
	}
	if n, ok := vertexpb.VertexModelSupportedExportFormatsExportableContentsEnum_name[int32(e)]; ok {
		e := vertex.ModelSupportedExportFormatsExportableContentsEnum(n[len("VertexModelSupportedExportFormatsExportableContentsEnum"):])
		return &e
	}
	return nil
}

// ProtoToModelSupportedDeploymentResourcesTypesEnum converts a ModelSupportedDeploymentResourcesTypesEnum enum from its proto representation.
func ProtoToVertexModelSupportedDeploymentResourcesTypesEnum(e vertexpb.VertexModelSupportedDeploymentResourcesTypesEnum) *vertex.ModelSupportedDeploymentResourcesTypesEnum {
	if e == 0 {
		return nil
	}
	if n, ok := vertexpb.VertexModelSupportedDeploymentResourcesTypesEnum_name[int32(e)]; ok {
		e := vertex.ModelSupportedDeploymentResourcesTypesEnum(n[len("VertexModelSupportedDeploymentResourcesTypesEnum"):])
		return &e
	}
	return nil
}

// ProtoToModelSupportedExportFormats converts a ModelSupportedExportFormats object from its proto representation.
func ProtoToVertexModelSupportedExportFormats(p *vertexpb.VertexModelSupportedExportFormats) *vertex.ModelSupportedExportFormats {
	if p == nil {
		return nil
	}
	obj := &vertex.ModelSupportedExportFormats{
		Id: dcl.StringOrNil(p.GetId()),
	}
	for _, r := range p.GetExportableContents() {
		obj.ExportableContents = append(obj.ExportableContents, *ProtoToVertexModelSupportedExportFormatsExportableContentsEnum(r))
	}
	return obj
}

// ProtoToModelOriginalModelInfo converts a ModelOriginalModelInfo object from its proto representation.
func ProtoToVertexModelOriginalModelInfo(p *vertexpb.VertexModelOriginalModelInfo) *vertex.ModelOriginalModelInfo {
	if p == nil {
		return nil
	}
	obj := &vertex.ModelOriginalModelInfo{
		Model: dcl.StringOrNil(p.GetModel()),
	}
	return obj
}

// ProtoToModelContainerSpec converts a ModelContainerSpec object from its proto representation.
func ProtoToVertexModelContainerSpec(p *vertexpb.VertexModelContainerSpec) *vertex.ModelContainerSpec {
	if p == nil {
		return nil
	}
	obj := &vertex.ModelContainerSpec{
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
		obj.Env = append(obj.Env, *ProtoToVertexModelContainerSpecEnv(r))
	}
	for _, r := range p.GetPorts() {
		obj.Ports = append(obj.Ports, *ProtoToVertexModelContainerSpecPorts(r))
	}
	return obj
}

// ProtoToModelContainerSpecEnv converts a ModelContainerSpecEnv object from its proto representation.
func ProtoToVertexModelContainerSpecEnv(p *vertexpb.VertexModelContainerSpecEnv) *vertex.ModelContainerSpecEnv {
	if p == nil {
		return nil
	}
	obj := &vertex.ModelContainerSpecEnv{
		Name:  dcl.StringOrNil(p.GetName()),
		Value: dcl.StringOrNil(p.GetValue()),
	}
	return obj
}

// ProtoToModelContainerSpecPorts converts a ModelContainerSpecPorts object from its proto representation.
func ProtoToVertexModelContainerSpecPorts(p *vertexpb.VertexModelContainerSpecPorts) *vertex.ModelContainerSpecPorts {
	if p == nil {
		return nil
	}
	obj := &vertex.ModelContainerSpecPorts{
		ContainerPort: dcl.Int64OrNil(p.GetContainerPort()),
	}
	return obj
}

// ProtoToModelDeployedModels converts a ModelDeployedModels object from its proto representation.
func ProtoToVertexModelDeployedModels(p *vertexpb.VertexModelDeployedModels) *vertex.ModelDeployedModels {
	if p == nil {
		return nil
	}
	obj := &vertex.ModelDeployedModels{
		Endpoint:        dcl.StringOrNil(p.GetEndpoint()),
		DeployedModelId: dcl.StringOrNil(p.GetDeployedModelId()),
	}
	return obj
}

// ProtoToModelEncryptionSpec converts a ModelEncryptionSpec object from its proto representation.
func ProtoToVertexModelEncryptionSpec(p *vertexpb.VertexModelEncryptionSpec) *vertex.ModelEncryptionSpec {
	if p == nil {
		return nil
	}
	obj := &vertex.ModelEncryptionSpec{
		KmsKeyName: dcl.StringOrNil(p.GetKmsKeyName()),
	}
	return obj
}

// ProtoToModel converts a Model resource from its proto representation.
func ProtoToModel(p *vertexpb.VertexModel) *vertex.Model {
	obj := &vertex.Model{
		Name:               dcl.StringOrNil(p.GetName()),
		VersionId:          dcl.StringOrNil(p.GetVersionId()),
		VersionCreateTime:  dcl.StringOrNil(p.GetVersionCreateTime()),
		VersionUpdateTime:  dcl.StringOrNil(p.GetVersionUpdateTime()),
		DisplayName:        dcl.StringOrNil(p.GetDisplayName()),
		Description:        dcl.StringOrNil(p.GetDescription()),
		VersionDescription: dcl.StringOrNil(p.GetVersionDescription()),
		TrainingPipeline:   dcl.StringOrNil(p.GetTrainingPipeline()),
		OriginalModelInfo:  ProtoToVertexModelOriginalModelInfo(p.GetOriginalModelInfo()),
		ContainerSpec:      ProtoToVertexModelContainerSpec(p.GetContainerSpec()),
		ArtifactUri:        dcl.StringOrNil(p.GetArtifactUri()),
		CreateTime:         dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:         dcl.StringOrNil(p.GetUpdateTime()),
		Etag:               dcl.StringOrNil(p.GetEtag()),
		EncryptionSpec:     ProtoToVertexModelEncryptionSpec(p.GetEncryptionSpec()),
		Project:            dcl.StringOrNil(p.GetProject()),
		Location:           dcl.StringOrNil(p.GetLocation()),
	}
	for _, r := range p.GetSupportedExportFormats() {
		obj.SupportedExportFormats = append(obj.SupportedExportFormats, *ProtoToVertexModelSupportedExportFormats(r))
	}
	for _, r := range p.GetSupportedDeploymentResourcesTypes() {
		obj.SupportedDeploymentResourcesTypes = append(obj.SupportedDeploymentResourcesTypes, *ProtoToVertexModelSupportedDeploymentResourcesTypesEnum(r))
	}
	for _, r := range p.GetSupportedInputStorageFormats() {
		obj.SupportedInputStorageFormats = append(obj.SupportedInputStorageFormats, r)
	}
	for _, r := range p.GetSupportedOutputStorageFormats() {
		obj.SupportedOutputStorageFormats = append(obj.SupportedOutputStorageFormats, r)
	}
	for _, r := range p.GetDeployedModels() {
		obj.DeployedModels = append(obj.DeployedModels, *ProtoToVertexModelDeployedModels(r))
	}
	return obj
}

// ModelSupportedExportFormatsExportableContentsEnumToProto converts a ModelSupportedExportFormatsExportableContentsEnum enum to its proto representation.
func VertexModelSupportedExportFormatsExportableContentsEnumToProto(e *vertex.ModelSupportedExportFormatsExportableContentsEnum) vertexpb.VertexModelSupportedExportFormatsExportableContentsEnum {
	if e == nil {
		return vertexpb.VertexModelSupportedExportFormatsExportableContentsEnum(0)
	}
	if v, ok := vertexpb.VertexModelSupportedExportFormatsExportableContentsEnum_value["ModelSupportedExportFormatsExportableContentsEnum"+string(*e)]; ok {
		return vertexpb.VertexModelSupportedExportFormatsExportableContentsEnum(v)
	}
	return vertexpb.VertexModelSupportedExportFormatsExportableContentsEnum(0)
}

// ModelSupportedDeploymentResourcesTypesEnumToProto converts a ModelSupportedDeploymentResourcesTypesEnum enum to its proto representation.
func VertexModelSupportedDeploymentResourcesTypesEnumToProto(e *vertex.ModelSupportedDeploymentResourcesTypesEnum) vertexpb.VertexModelSupportedDeploymentResourcesTypesEnum {
	if e == nil {
		return vertexpb.VertexModelSupportedDeploymentResourcesTypesEnum(0)
	}
	if v, ok := vertexpb.VertexModelSupportedDeploymentResourcesTypesEnum_value["ModelSupportedDeploymentResourcesTypesEnum"+string(*e)]; ok {
		return vertexpb.VertexModelSupportedDeploymentResourcesTypesEnum(v)
	}
	return vertexpb.VertexModelSupportedDeploymentResourcesTypesEnum(0)
}

// ModelSupportedExportFormatsToProto converts a ModelSupportedExportFormats object to its proto representation.
func VertexModelSupportedExportFormatsToProto(o *vertex.ModelSupportedExportFormats) *vertexpb.VertexModelSupportedExportFormats {
	if o == nil {
		return nil
	}
	p := &vertexpb.VertexModelSupportedExportFormats{}
	p.SetId(dcl.ValueOrEmptyString(o.Id))
	sExportableContents := make([]vertexpb.VertexModelSupportedExportFormatsExportableContentsEnum, len(o.ExportableContents))
	for i, r := range o.ExportableContents {
		sExportableContents[i] = vertexpb.VertexModelSupportedExportFormatsExportableContentsEnum(vertexpb.VertexModelSupportedExportFormatsExportableContentsEnum_value[string(r)])
	}
	p.SetExportableContents(sExportableContents)
	return p
}

// ModelOriginalModelInfoToProto converts a ModelOriginalModelInfo object to its proto representation.
func VertexModelOriginalModelInfoToProto(o *vertex.ModelOriginalModelInfo) *vertexpb.VertexModelOriginalModelInfo {
	if o == nil {
		return nil
	}
	p := &vertexpb.VertexModelOriginalModelInfo{}
	p.SetModel(dcl.ValueOrEmptyString(o.Model))
	return p
}

// ModelContainerSpecToProto converts a ModelContainerSpec object to its proto representation.
func VertexModelContainerSpecToProto(o *vertex.ModelContainerSpec) *vertexpb.VertexModelContainerSpec {
	if o == nil {
		return nil
	}
	p := &vertexpb.VertexModelContainerSpec{}
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
	sEnv := make([]*vertexpb.VertexModelContainerSpecEnv, len(o.Env))
	for i, r := range o.Env {
		sEnv[i] = VertexModelContainerSpecEnvToProto(&r)
	}
	p.SetEnv(sEnv)
	sPorts := make([]*vertexpb.VertexModelContainerSpecPorts, len(o.Ports))
	for i, r := range o.Ports {
		sPorts[i] = VertexModelContainerSpecPortsToProto(&r)
	}
	p.SetPorts(sPorts)
	return p
}

// ModelContainerSpecEnvToProto converts a ModelContainerSpecEnv object to its proto representation.
func VertexModelContainerSpecEnvToProto(o *vertex.ModelContainerSpecEnv) *vertexpb.VertexModelContainerSpecEnv {
	if o == nil {
		return nil
	}
	p := &vertexpb.VertexModelContainerSpecEnv{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetValue(dcl.ValueOrEmptyString(o.Value))
	return p
}

// ModelContainerSpecPortsToProto converts a ModelContainerSpecPorts object to its proto representation.
func VertexModelContainerSpecPortsToProto(o *vertex.ModelContainerSpecPorts) *vertexpb.VertexModelContainerSpecPorts {
	if o == nil {
		return nil
	}
	p := &vertexpb.VertexModelContainerSpecPorts{}
	p.SetContainerPort(dcl.ValueOrEmptyInt64(o.ContainerPort))
	return p
}

// ModelDeployedModelsToProto converts a ModelDeployedModels object to its proto representation.
func VertexModelDeployedModelsToProto(o *vertex.ModelDeployedModels) *vertexpb.VertexModelDeployedModels {
	if o == nil {
		return nil
	}
	p := &vertexpb.VertexModelDeployedModels{}
	p.SetEndpoint(dcl.ValueOrEmptyString(o.Endpoint))
	p.SetDeployedModelId(dcl.ValueOrEmptyString(o.DeployedModelId))
	return p
}

// ModelEncryptionSpecToProto converts a ModelEncryptionSpec object to its proto representation.
func VertexModelEncryptionSpecToProto(o *vertex.ModelEncryptionSpec) *vertexpb.VertexModelEncryptionSpec {
	if o == nil {
		return nil
	}
	p := &vertexpb.VertexModelEncryptionSpec{}
	p.SetKmsKeyName(dcl.ValueOrEmptyString(o.KmsKeyName))
	return p
}

// ModelToProto converts a Model resource to its proto representation.
func ModelToProto(resource *vertex.Model) *vertexpb.VertexModel {
	p := &vertexpb.VertexModel{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetVersionId(dcl.ValueOrEmptyString(resource.VersionId))
	p.SetVersionCreateTime(dcl.ValueOrEmptyString(resource.VersionCreateTime))
	p.SetVersionUpdateTime(dcl.ValueOrEmptyString(resource.VersionUpdateTime))
	p.SetDisplayName(dcl.ValueOrEmptyString(resource.DisplayName))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetVersionDescription(dcl.ValueOrEmptyString(resource.VersionDescription))
	p.SetTrainingPipeline(dcl.ValueOrEmptyString(resource.TrainingPipeline))
	p.SetOriginalModelInfo(VertexModelOriginalModelInfoToProto(resource.OriginalModelInfo))
	p.SetContainerSpec(VertexModelContainerSpecToProto(resource.ContainerSpec))
	p.SetArtifactUri(dcl.ValueOrEmptyString(resource.ArtifactUri))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetEtag(dcl.ValueOrEmptyString(resource.Etag))
	p.SetEncryptionSpec(VertexModelEncryptionSpecToProto(resource.EncryptionSpec))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	sSupportedExportFormats := make([]*vertexpb.VertexModelSupportedExportFormats, len(resource.SupportedExportFormats))
	for i, r := range resource.SupportedExportFormats {
		sSupportedExportFormats[i] = VertexModelSupportedExportFormatsToProto(&r)
	}
	p.SetSupportedExportFormats(sSupportedExportFormats)
	sSupportedDeploymentResourcesTypes := make([]vertexpb.VertexModelSupportedDeploymentResourcesTypesEnum, len(resource.SupportedDeploymentResourcesTypes))
	for i, r := range resource.SupportedDeploymentResourcesTypes {
		sSupportedDeploymentResourcesTypes[i] = vertexpb.VertexModelSupportedDeploymentResourcesTypesEnum(vertexpb.VertexModelSupportedDeploymentResourcesTypesEnum_value[string(r)])
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
	sDeployedModels := make([]*vertexpb.VertexModelDeployedModels, len(resource.DeployedModels))
	for i, r := range resource.DeployedModels {
		sDeployedModels[i] = VertexModelDeployedModelsToProto(&r)
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
func (s *ModelServer) applyModel(ctx context.Context, c *vertex.Client, request *vertexpb.ApplyVertexModelRequest) (*vertexpb.VertexModel, error) {
	p := ProtoToModel(request.GetResource())
	res, err := c.ApplyModel(ctx, p)
	if err != nil {
		return nil, err
	}
	r := ModelToProto(res)
	return r, nil
}

// applyVertexModel handles the gRPC request by passing it to the underlying Model Apply() method.
func (s *ModelServer) ApplyVertexModel(ctx context.Context, request *vertexpb.ApplyVertexModelRequest) (*vertexpb.VertexModel, error) {
	cl, err := createConfigModel(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyModel(ctx, cl, request)
}

// DeleteModel handles the gRPC request by passing it to the underlying Model Delete() method.
func (s *ModelServer) DeleteVertexModel(ctx context.Context, request *vertexpb.DeleteVertexModelRequest) (*emptypb.Empty, error) {

	cl, err := createConfigModel(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteModel(ctx, ProtoToModel(request.GetResource()))

}

// ListVertexModel handles the gRPC request by passing it to the underlying ModelList() method.
func (s *ModelServer) ListVertexModel(ctx context.Context, request *vertexpb.ListVertexModelRequest) (*vertexpb.ListVertexModelResponse, error) {
	cl, err := createConfigModel(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListModel(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*vertexpb.VertexModel
	for _, r := range resources.Items {
		rp := ModelToProto(r)
		protos = append(protos, rp)
	}
	p := &vertexpb.ListVertexModelResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigModel(ctx context.Context, service_account_file string) (*vertex.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return vertex.NewClient(conf), nil
}
