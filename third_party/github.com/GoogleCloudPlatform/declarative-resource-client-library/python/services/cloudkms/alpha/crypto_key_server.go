// Copyright 2024 Google LLC. All Rights Reserved.
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
	"errors"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/cloudkms/alpha/cloudkms_alpha_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudkms/alpha"
)

// CryptoKeyServer implements the gRPC interface for CryptoKey.
type CryptoKeyServer struct{}

// ProtoToCryptoKeyPrimaryStateEnum converts a CryptoKeyPrimaryStateEnum enum from its proto representation.
func ProtoToCloudkmsAlphaCryptoKeyPrimaryStateEnum(e alphapb.CloudkmsAlphaCryptoKeyPrimaryStateEnum) *alpha.CryptoKeyPrimaryStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.CloudkmsAlphaCryptoKeyPrimaryStateEnum_name[int32(e)]; ok {
		e := alpha.CryptoKeyPrimaryStateEnum(n[len("CloudkmsAlphaCryptoKeyPrimaryStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToCryptoKeyPrimaryProtectionLevelEnum converts a CryptoKeyPrimaryProtectionLevelEnum enum from its proto representation.
func ProtoToCloudkmsAlphaCryptoKeyPrimaryProtectionLevelEnum(e alphapb.CloudkmsAlphaCryptoKeyPrimaryProtectionLevelEnum) *alpha.CryptoKeyPrimaryProtectionLevelEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.CloudkmsAlphaCryptoKeyPrimaryProtectionLevelEnum_name[int32(e)]; ok {
		e := alpha.CryptoKeyPrimaryProtectionLevelEnum(n[len("CloudkmsAlphaCryptoKeyPrimaryProtectionLevelEnum"):])
		return &e
	}
	return nil
}

// ProtoToCryptoKeyPrimaryAlgorithmEnum converts a CryptoKeyPrimaryAlgorithmEnum enum from its proto representation.
func ProtoToCloudkmsAlphaCryptoKeyPrimaryAlgorithmEnum(e alphapb.CloudkmsAlphaCryptoKeyPrimaryAlgorithmEnum) *alpha.CryptoKeyPrimaryAlgorithmEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.CloudkmsAlphaCryptoKeyPrimaryAlgorithmEnum_name[int32(e)]; ok {
		e := alpha.CryptoKeyPrimaryAlgorithmEnum(n[len("CloudkmsAlphaCryptoKeyPrimaryAlgorithmEnum"):])
		return &e
	}
	return nil
}

// ProtoToCryptoKeyPrimaryAttestationFormatEnum converts a CryptoKeyPrimaryAttestationFormatEnum enum from its proto representation.
func ProtoToCloudkmsAlphaCryptoKeyPrimaryAttestationFormatEnum(e alphapb.CloudkmsAlphaCryptoKeyPrimaryAttestationFormatEnum) *alpha.CryptoKeyPrimaryAttestationFormatEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.CloudkmsAlphaCryptoKeyPrimaryAttestationFormatEnum_name[int32(e)]; ok {
		e := alpha.CryptoKeyPrimaryAttestationFormatEnum(n[len("CloudkmsAlphaCryptoKeyPrimaryAttestationFormatEnum"):])
		return &e
	}
	return nil
}

// ProtoToCryptoKeyPurposeEnum converts a CryptoKeyPurposeEnum enum from its proto representation.
func ProtoToCloudkmsAlphaCryptoKeyPurposeEnum(e alphapb.CloudkmsAlphaCryptoKeyPurposeEnum) *alpha.CryptoKeyPurposeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.CloudkmsAlphaCryptoKeyPurposeEnum_name[int32(e)]; ok {
		e := alpha.CryptoKeyPurposeEnum(n[len("CloudkmsAlphaCryptoKeyPurposeEnum"):])
		return &e
	}
	return nil
}

// ProtoToCryptoKeyVersionTemplateProtectionLevelEnum converts a CryptoKeyVersionTemplateProtectionLevelEnum enum from its proto representation.
func ProtoToCloudkmsAlphaCryptoKeyVersionTemplateProtectionLevelEnum(e alphapb.CloudkmsAlphaCryptoKeyVersionTemplateProtectionLevelEnum) *alpha.CryptoKeyVersionTemplateProtectionLevelEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.CloudkmsAlphaCryptoKeyVersionTemplateProtectionLevelEnum_name[int32(e)]; ok {
		e := alpha.CryptoKeyVersionTemplateProtectionLevelEnum(n[len("CloudkmsAlphaCryptoKeyVersionTemplateProtectionLevelEnum"):])
		return &e
	}
	return nil
}

// ProtoToCryptoKeyVersionTemplateAlgorithmEnum converts a CryptoKeyVersionTemplateAlgorithmEnum enum from its proto representation.
func ProtoToCloudkmsAlphaCryptoKeyVersionTemplateAlgorithmEnum(e alphapb.CloudkmsAlphaCryptoKeyVersionTemplateAlgorithmEnum) *alpha.CryptoKeyVersionTemplateAlgorithmEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.CloudkmsAlphaCryptoKeyVersionTemplateAlgorithmEnum_name[int32(e)]; ok {
		e := alpha.CryptoKeyVersionTemplateAlgorithmEnum(n[len("CloudkmsAlphaCryptoKeyVersionTemplateAlgorithmEnum"):])
		return &e
	}
	return nil
}

// ProtoToCryptoKeyPrimary converts a CryptoKeyPrimary object from its proto representation.
func ProtoToCloudkmsAlphaCryptoKeyPrimary(p *alphapb.CloudkmsAlphaCryptoKeyPrimary) *alpha.CryptoKeyPrimary {
	if p == nil {
		return nil
	}
	obj := &alpha.CryptoKeyPrimary{
		Name:                           dcl.StringOrNil(p.GetName()),
		State:                          ProtoToCloudkmsAlphaCryptoKeyPrimaryStateEnum(p.GetState()),
		ProtectionLevel:                ProtoToCloudkmsAlphaCryptoKeyPrimaryProtectionLevelEnum(p.GetProtectionLevel()),
		Algorithm:                      ProtoToCloudkmsAlphaCryptoKeyPrimaryAlgorithmEnum(p.GetAlgorithm()),
		Attestation:                    ProtoToCloudkmsAlphaCryptoKeyPrimaryAttestation(p.GetAttestation()),
		CreateTime:                     dcl.StringOrNil(p.GetCreateTime()),
		GenerateTime:                   dcl.StringOrNil(p.GetGenerateTime()),
		DestroyTime:                    dcl.StringOrNil(p.GetDestroyTime()),
		DestroyEventTime:               dcl.StringOrNil(p.GetDestroyEventTime()),
		ImportJob:                      dcl.StringOrNil(p.GetImportJob()),
		ImportTime:                     dcl.StringOrNil(p.GetImportTime()),
		ImportFailureReason:            dcl.StringOrNil(p.GetImportFailureReason()),
		ExternalProtectionLevelOptions: ProtoToCloudkmsAlphaCryptoKeyPrimaryExternalProtectionLevelOptions(p.GetExternalProtectionLevelOptions()),
		ReimportEligible:               dcl.Bool(p.GetReimportEligible()),
	}
	return obj
}

// ProtoToCryptoKeyPrimaryAttestation converts a CryptoKeyPrimaryAttestation object from its proto representation.
func ProtoToCloudkmsAlphaCryptoKeyPrimaryAttestation(p *alphapb.CloudkmsAlphaCryptoKeyPrimaryAttestation) *alpha.CryptoKeyPrimaryAttestation {
	if p == nil {
		return nil
	}
	obj := &alpha.CryptoKeyPrimaryAttestation{
		Format:     ProtoToCloudkmsAlphaCryptoKeyPrimaryAttestationFormatEnum(p.GetFormat()),
		Content:    dcl.StringOrNil(p.GetContent()),
		CertChains: ProtoToCloudkmsAlphaCryptoKeyPrimaryAttestationCertChains(p.GetCertChains()),
	}
	return obj
}

// ProtoToCryptoKeyPrimaryAttestationCertChains converts a CryptoKeyPrimaryAttestationCertChains object from its proto representation.
func ProtoToCloudkmsAlphaCryptoKeyPrimaryAttestationCertChains(p *alphapb.CloudkmsAlphaCryptoKeyPrimaryAttestationCertChains) *alpha.CryptoKeyPrimaryAttestationCertChains {
	if p == nil {
		return nil
	}
	obj := &alpha.CryptoKeyPrimaryAttestationCertChains{}
	for _, r := range p.GetCaviumCerts() {
		obj.CaviumCerts = append(obj.CaviumCerts, r)
	}
	for _, r := range p.GetGoogleCardCerts() {
		obj.GoogleCardCerts = append(obj.GoogleCardCerts, r)
	}
	for _, r := range p.GetGooglePartitionCerts() {
		obj.GooglePartitionCerts = append(obj.GooglePartitionCerts, r)
	}
	return obj
}

// ProtoToCryptoKeyPrimaryExternalProtectionLevelOptions converts a CryptoKeyPrimaryExternalProtectionLevelOptions object from its proto representation.
func ProtoToCloudkmsAlphaCryptoKeyPrimaryExternalProtectionLevelOptions(p *alphapb.CloudkmsAlphaCryptoKeyPrimaryExternalProtectionLevelOptions) *alpha.CryptoKeyPrimaryExternalProtectionLevelOptions {
	if p == nil {
		return nil
	}
	obj := &alpha.CryptoKeyPrimaryExternalProtectionLevelOptions{
		ExternalKeyUri: dcl.StringOrNil(p.GetExternalKeyUri()),
	}
	return obj
}

// ProtoToCryptoKeyVersionTemplate converts a CryptoKeyVersionTemplate object from its proto representation.
func ProtoToCloudkmsAlphaCryptoKeyVersionTemplate(p *alphapb.CloudkmsAlphaCryptoKeyVersionTemplate) *alpha.CryptoKeyVersionTemplate {
	if p == nil {
		return nil
	}
	obj := &alpha.CryptoKeyVersionTemplate{
		ProtectionLevel: ProtoToCloudkmsAlphaCryptoKeyVersionTemplateProtectionLevelEnum(p.GetProtectionLevel()),
		Algorithm:       ProtoToCloudkmsAlphaCryptoKeyVersionTemplateAlgorithmEnum(p.GetAlgorithm()),
	}
	return obj
}

// ProtoToCryptoKey converts a CryptoKey resource from its proto representation.
func ProtoToCryptoKey(p *alphapb.CloudkmsAlphaCryptoKey) *alpha.CryptoKey {
	obj := &alpha.CryptoKey{
		Name:                     dcl.StringOrNil(p.GetName()),
		Primary:                  ProtoToCloudkmsAlphaCryptoKeyPrimary(p.GetPrimary()),
		Purpose:                  ProtoToCloudkmsAlphaCryptoKeyPurposeEnum(p.GetPurpose()),
		CreateTime:               dcl.StringOrNil(p.GetCreateTime()),
		NextRotationTime:         dcl.StringOrNil(p.GetNextRotationTime()),
		RotationPeriod:           dcl.StringOrNil(p.GetRotationPeriod()),
		VersionTemplate:          ProtoToCloudkmsAlphaCryptoKeyVersionTemplate(p.GetVersionTemplate()),
		ImportOnly:               dcl.Bool(p.GetImportOnly()),
		DestroyScheduledDuration: dcl.StringOrNil(p.GetDestroyScheduledDuration()),
		Project:                  dcl.StringOrNil(p.GetProject()),
		Location:                 dcl.StringOrNil(p.GetLocation()),
		KeyRing:                  dcl.StringOrNil(p.GetKeyRing()),
	}
	return obj
}

// CryptoKeyPrimaryStateEnumToProto converts a CryptoKeyPrimaryStateEnum enum to its proto representation.
func CloudkmsAlphaCryptoKeyPrimaryStateEnumToProto(e *alpha.CryptoKeyPrimaryStateEnum) alphapb.CloudkmsAlphaCryptoKeyPrimaryStateEnum {
	if e == nil {
		return alphapb.CloudkmsAlphaCryptoKeyPrimaryStateEnum(0)
	}
	if v, ok := alphapb.CloudkmsAlphaCryptoKeyPrimaryStateEnum_value["CryptoKeyPrimaryStateEnum"+string(*e)]; ok {
		return alphapb.CloudkmsAlphaCryptoKeyPrimaryStateEnum(v)
	}
	return alphapb.CloudkmsAlphaCryptoKeyPrimaryStateEnum(0)
}

// CryptoKeyPrimaryProtectionLevelEnumToProto converts a CryptoKeyPrimaryProtectionLevelEnum enum to its proto representation.
func CloudkmsAlphaCryptoKeyPrimaryProtectionLevelEnumToProto(e *alpha.CryptoKeyPrimaryProtectionLevelEnum) alphapb.CloudkmsAlphaCryptoKeyPrimaryProtectionLevelEnum {
	if e == nil {
		return alphapb.CloudkmsAlphaCryptoKeyPrimaryProtectionLevelEnum(0)
	}
	if v, ok := alphapb.CloudkmsAlphaCryptoKeyPrimaryProtectionLevelEnum_value["CryptoKeyPrimaryProtectionLevelEnum"+string(*e)]; ok {
		return alphapb.CloudkmsAlphaCryptoKeyPrimaryProtectionLevelEnum(v)
	}
	return alphapb.CloudkmsAlphaCryptoKeyPrimaryProtectionLevelEnum(0)
}

// CryptoKeyPrimaryAlgorithmEnumToProto converts a CryptoKeyPrimaryAlgorithmEnum enum to its proto representation.
func CloudkmsAlphaCryptoKeyPrimaryAlgorithmEnumToProto(e *alpha.CryptoKeyPrimaryAlgorithmEnum) alphapb.CloudkmsAlphaCryptoKeyPrimaryAlgorithmEnum {
	if e == nil {
		return alphapb.CloudkmsAlphaCryptoKeyPrimaryAlgorithmEnum(0)
	}
	if v, ok := alphapb.CloudkmsAlphaCryptoKeyPrimaryAlgorithmEnum_value["CryptoKeyPrimaryAlgorithmEnum"+string(*e)]; ok {
		return alphapb.CloudkmsAlphaCryptoKeyPrimaryAlgorithmEnum(v)
	}
	return alphapb.CloudkmsAlphaCryptoKeyPrimaryAlgorithmEnum(0)
}

// CryptoKeyPrimaryAttestationFormatEnumToProto converts a CryptoKeyPrimaryAttestationFormatEnum enum to its proto representation.
func CloudkmsAlphaCryptoKeyPrimaryAttestationFormatEnumToProto(e *alpha.CryptoKeyPrimaryAttestationFormatEnum) alphapb.CloudkmsAlphaCryptoKeyPrimaryAttestationFormatEnum {
	if e == nil {
		return alphapb.CloudkmsAlphaCryptoKeyPrimaryAttestationFormatEnum(0)
	}
	if v, ok := alphapb.CloudkmsAlphaCryptoKeyPrimaryAttestationFormatEnum_value["CryptoKeyPrimaryAttestationFormatEnum"+string(*e)]; ok {
		return alphapb.CloudkmsAlphaCryptoKeyPrimaryAttestationFormatEnum(v)
	}
	return alphapb.CloudkmsAlphaCryptoKeyPrimaryAttestationFormatEnum(0)
}

// CryptoKeyPurposeEnumToProto converts a CryptoKeyPurposeEnum enum to its proto representation.
func CloudkmsAlphaCryptoKeyPurposeEnumToProto(e *alpha.CryptoKeyPurposeEnum) alphapb.CloudkmsAlphaCryptoKeyPurposeEnum {
	if e == nil {
		return alphapb.CloudkmsAlphaCryptoKeyPurposeEnum(0)
	}
	if v, ok := alphapb.CloudkmsAlphaCryptoKeyPurposeEnum_value["CryptoKeyPurposeEnum"+string(*e)]; ok {
		return alphapb.CloudkmsAlphaCryptoKeyPurposeEnum(v)
	}
	return alphapb.CloudkmsAlphaCryptoKeyPurposeEnum(0)
}

// CryptoKeyVersionTemplateProtectionLevelEnumToProto converts a CryptoKeyVersionTemplateProtectionLevelEnum enum to its proto representation.
func CloudkmsAlphaCryptoKeyVersionTemplateProtectionLevelEnumToProto(e *alpha.CryptoKeyVersionTemplateProtectionLevelEnum) alphapb.CloudkmsAlphaCryptoKeyVersionTemplateProtectionLevelEnum {
	if e == nil {
		return alphapb.CloudkmsAlphaCryptoKeyVersionTemplateProtectionLevelEnum(0)
	}
	if v, ok := alphapb.CloudkmsAlphaCryptoKeyVersionTemplateProtectionLevelEnum_value["CryptoKeyVersionTemplateProtectionLevelEnum"+string(*e)]; ok {
		return alphapb.CloudkmsAlphaCryptoKeyVersionTemplateProtectionLevelEnum(v)
	}
	return alphapb.CloudkmsAlphaCryptoKeyVersionTemplateProtectionLevelEnum(0)
}

// CryptoKeyVersionTemplateAlgorithmEnumToProto converts a CryptoKeyVersionTemplateAlgorithmEnum enum to its proto representation.
func CloudkmsAlphaCryptoKeyVersionTemplateAlgorithmEnumToProto(e *alpha.CryptoKeyVersionTemplateAlgorithmEnum) alphapb.CloudkmsAlphaCryptoKeyVersionTemplateAlgorithmEnum {
	if e == nil {
		return alphapb.CloudkmsAlphaCryptoKeyVersionTemplateAlgorithmEnum(0)
	}
	if v, ok := alphapb.CloudkmsAlphaCryptoKeyVersionTemplateAlgorithmEnum_value["CryptoKeyVersionTemplateAlgorithmEnum"+string(*e)]; ok {
		return alphapb.CloudkmsAlphaCryptoKeyVersionTemplateAlgorithmEnum(v)
	}
	return alphapb.CloudkmsAlphaCryptoKeyVersionTemplateAlgorithmEnum(0)
}

// CryptoKeyPrimaryToProto converts a CryptoKeyPrimary object to its proto representation.
func CloudkmsAlphaCryptoKeyPrimaryToProto(o *alpha.CryptoKeyPrimary) *alphapb.CloudkmsAlphaCryptoKeyPrimary {
	if o == nil {
		return nil
	}
	p := &alphapb.CloudkmsAlphaCryptoKeyPrimary{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetState(CloudkmsAlphaCryptoKeyPrimaryStateEnumToProto(o.State))
	p.SetProtectionLevel(CloudkmsAlphaCryptoKeyPrimaryProtectionLevelEnumToProto(o.ProtectionLevel))
	p.SetAlgorithm(CloudkmsAlphaCryptoKeyPrimaryAlgorithmEnumToProto(o.Algorithm))
	p.SetAttestation(CloudkmsAlphaCryptoKeyPrimaryAttestationToProto(o.Attestation))
	p.SetCreateTime(dcl.ValueOrEmptyString(o.CreateTime))
	p.SetGenerateTime(dcl.ValueOrEmptyString(o.GenerateTime))
	p.SetDestroyTime(dcl.ValueOrEmptyString(o.DestroyTime))
	p.SetDestroyEventTime(dcl.ValueOrEmptyString(o.DestroyEventTime))
	p.SetImportJob(dcl.ValueOrEmptyString(o.ImportJob))
	p.SetImportTime(dcl.ValueOrEmptyString(o.ImportTime))
	p.SetImportFailureReason(dcl.ValueOrEmptyString(o.ImportFailureReason))
	p.SetExternalProtectionLevelOptions(CloudkmsAlphaCryptoKeyPrimaryExternalProtectionLevelOptionsToProto(o.ExternalProtectionLevelOptions))
	p.SetReimportEligible(dcl.ValueOrEmptyBool(o.ReimportEligible))
	return p
}

// CryptoKeyPrimaryAttestationToProto converts a CryptoKeyPrimaryAttestation object to its proto representation.
func CloudkmsAlphaCryptoKeyPrimaryAttestationToProto(o *alpha.CryptoKeyPrimaryAttestation) *alphapb.CloudkmsAlphaCryptoKeyPrimaryAttestation {
	if o == nil {
		return nil
	}
	p := &alphapb.CloudkmsAlphaCryptoKeyPrimaryAttestation{}
	p.SetFormat(CloudkmsAlphaCryptoKeyPrimaryAttestationFormatEnumToProto(o.Format))
	p.SetContent(dcl.ValueOrEmptyString(o.Content))
	p.SetCertChains(CloudkmsAlphaCryptoKeyPrimaryAttestationCertChainsToProto(o.CertChains))
	return p
}

// CryptoKeyPrimaryAttestationCertChainsToProto converts a CryptoKeyPrimaryAttestationCertChains object to its proto representation.
func CloudkmsAlphaCryptoKeyPrimaryAttestationCertChainsToProto(o *alpha.CryptoKeyPrimaryAttestationCertChains) *alphapb.CloudkmsAlphaCryptoKeyPrimaryAttestationCertChains {
	if o == nil {
		return nil
	}
	p := &alphapb.CloudkmsAlphaCryptoKeyPrimaryAttestationCertChains{}
	sCaviumCerts := make([]string, len(o.CaviumCerts))
	for i, r := range o.CaviumCerts {
		sCaviumCerts[i] = r
	}
	p.SetCaviumCerts(sCaviumCerts)
	sGoogleCardCerts := make([]string, len(o.GoogleCardCerts))
	for i, r := range o.GoogleCardCerts {
		sGoogleCardCerts[i] = r
	}
	p.SetGoogleCardCerts(sGoogleCardCerts)
	sGooglePartitionCerts := make([]string, len(o.GooglePartitionCerts))
	for i, r := range o.GooglePartitionCerts {
		sGooglePartitionCerts[i] = r
	}
	p.SetGooglePartitionCerts(sGooglePartitionCerts)
	return p
}

// CryptoKeyPrimaryExternalProtectionLevelOptionsToProto converts a CryptoKeyPrimaryExternalProtectionLevelOptions object to its proto representation.
func CloudkmsAlphaCryptoKeyPrimaryExternalProtectionLevelOptionsToProto(o *alpha.CryptoKeyPrimaryExternalProtectionLevelOptions) *alphapb.CloudkmsAlphaCryptoKeyPrimaryExternalProtectionLevelOptions {
	if o == nil {
		return nil
	}
	p := &alphapb.CloudkmsAlphaCryptoKeyPrimaryExternalProtectionLevelOptions{}
	p.SetExternalKeyUri(dcl.ValueOrEmptyString(o.ExternalKeyUri))
	return p
}

// CryptoKeyVersionTemplateToProto converts a CryptoKeyVersionTemplate object to its proto representation.
func CloudkmsAlphaCryptoKeyVersionTemplateToProto(o *alpha.CryptoKeyVersionTemplate) *alphapb.CloudkmsAlphaCryptoKeyVersionTemplate {
	if o == nil {
		return nil
	}
	p := &alphapb.CloudkmsAlphaCryptoKeyVersionTemplate{}
	p.SetProtectionLevel(CloudkmsAlphaCryptoKeyVersionTemplateProtectionLevelEnumToProto(o.ProtectionLevel))
	p.SetAlgorithm(CloudkmsAlphaCryptoKeyVersionTemplateAlgorithmEnumToProto(o.Algorithm))
	return p
}

// CryptoKeyToProto converts a CryptoKey resource to its proto representation.
func CryptoKeyToProto(resource *alpha.CryptoKey) *alphapb.CloudkmsAlphaCryptoKey {
	p := &alphapb.CloudkmsAlphaCryptoKey{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetPrimary(CloudkmsAlphaCryptoKeyPrimaryToProto(resource.Primary))
	p.SetPurpose(CloudkmsAlphaCryptoKeyPurposeEnumToProto(resource.Purpose))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetNextRotationTime(dcl.ValueOrEmptyString(resource.NextRotationTime))
	p.SetRotationPeriod(dcl.ValueOrEmptyString(resource.RotationPeriod))
	p.SetVersionTemplate(CloudkmsAlphaCryptoKeyVersionTemplateToProto(resource.VersionTemplate))
	p.SetImportOnly(dcl.ValueOrEmptyBool(resource.ImportOnly))
	p.SetDestroyScheduledDuration(dcl.ValueOrEmptyString(resource.DestroyScheduledDuration))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	p.SetKeyRing(dcl.ValueOrEmptyString(resource.KeyRing))
	mLabels := make(map[string]string, len(resource.Labels))
	for k, r := range resource.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)

	return p
}

// applyCryptoKey handles the gRPC request by passing it to the underlying CryptoKey Apply() method.
func (s *CryptoKeyServer) applyCryptoKey(ctx context.Context, c *alpha.Client, request *alphapb.ApplyCloudkmsAlphaCryptoKeyRequest) (*alphapb.CloudkmsAlphaCryptoKey, error) {
	p := ProtoToCryptoKey(request.GetResource())
	res, err := c.ApplyCryptoKey(ctx, p)
	if err != nil {
		return nil, err
	}
	r := CryptoKeyToProto(res)
	return r, nil
}

// applyCloudkmsAlphaCryptoKey handles the gRPC request by passing it to the underlying CryptoKey Apply() method.
func (s *CryptoKeyServer) ApplyCloudkmsAlphaCryptoKey(ctx context.Context, request *alphapb.ApplyCloudkmsAlphaCryptoKeyRequest) (*alphapb.CloudkmsAlphaCryptoKey, error) {
	cl, err := createConfigCryptoKey(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyCryptoKey(ctx, cl, request)
}

// DeleteCryptoKey handles the gRPC request by passing it to the underlying CryptoKey Delete() method.
func (s *CryptoKeyServer) DeleteCloudkmsAlphaCryptoKey(ctx context.Context, request *alphapb.DeleteCloudkmsAlphaCryptoKeyRequest) (*emptypb.Empty, error) {

	return nil, errors.New("no delete endpoint for CryptoKey")

}

// ListCloudkmsAlphaCryptoKey handles the gRPC request by passing it to the underlying CryptoKeyList() method.
func (s *CryptoKeyServer) ListCloudkmsAlphaCryptoKey(ctx context.Context, request *alphapb.ListCloudkmsAlphaCryptoKeyRequest) (*alphapb.ListCloudkmsAlphaCryptoKeyResponse, error) {
	cl, err := createConfigCryptoKey(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListCryptoKey(ctx, request.GetProject(), request.GetLocation(), request.GetKeyRing())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.CloudkmsAlphaCryptoKey
	for _, r := range resources.Items {
		rp := CryptoKeyToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListCloudkmsAlphaCryptoKeyResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigCryptoKey(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
