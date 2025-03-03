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
	cloudkmspb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/cloudkms/cloudkms_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudkms"
)

// CryptoKeyServer implements the gRPC interface for CryptoKey.
type CryptoKeyServer struct{}

// ProtoToCryptoKeyPrimaryStateEnum converts a CryptoKeyPrimaryStateEnum enum from its proto representation.
func ProtoToCloudkmsCryptoKeyPrimaryStateEnum(e cloudkmspb.CloudkmsCryptoKeyPrimaryStateEnum) *cloudkms.CryptoKeyPrimaryStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := cloudkmspb.CloudkmsCryptoKeyPrimaryStateEnum_name[int32(e)]; ok {
		e := cloudkms.CryptoKeyPrimaryStateEnum(n[len("CloudkmsCryptoKeyPrimaryStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToCryptoKeyPrimaryProtectionLevelEnum converts a CryptoKeyPrimaryProtectionLevelEnum enum from its proto representation.
func ProtoToCloudkmsCryptoKeyPrimaryProtectionLevelEnum(e cloudkmspb.CloudkmsCryptoKeyPrimaryProtectionLevelEnum) *cloudkms.CryptoKeyPrimaryProtectionLevelEnum {
	if e == 0 {
		return nil
	}
	if n, ok := cloudkmspb.CloudkmsCryptoKeyPrimaryProtectionLevelEnum_name[int32(e)]; ok {
		e := cloudkms.CryptoKeyPrimaryProtectionLevelEnum(n[len("CloudkmsCryptoKeyPrimaryProtectionLevelEnum"):])
		return &e
	}
	return nil
}

// ProtoToCryptoKeyPrimaryAlgorithmEnum converts a CryptoKeyPrimaryAlgorithmEnum enum from its proto representation.
func ProtoToCloudkmsCryptoKeyPrimaryAlgorithmEnum(e cloudkmspb.CloudkmsCryptoKeyPrimaryAlgorithmEnum) *cloudkms.CryptoKeyPrimaryAlgorithmEnum {
	if e == 0 {
		return nil
	}
	if n, ok := cloudkmspb.CloudkmsCryptoKeyPrimaryAlgorithmEnum_name[int32(e)]; ok {
		e := cloudkms.CryptoKeyPrimaryAlgorithmEnum(n[len("CloudkmsCryptoKeyPrimaryAlgorithmEnum"):])
		return &e
	}
	return nil
}

// ProtoToCryptoKeyPrimaryAttestationFormatEnum converts a CryptoKeyPrimaryAttestationFormatEnum enum from its proto representation.
func ProtoToCloudkmsCryptoKeyPrimaryAttestationFormatEnum(e cloudkmspb.CloudkmsCryptoKeyPrimaryAttestationFormatEnum) *cloudkms.CryptoKeyPrimaryAttestationFormatEnum {
	if e == 0 {
		return nil
	}
	if n, ok := cloudkmspb.CloudkmsCryptoKeyPrimaryAttestationFormatEnum_name[int32(e)]; ok {
		e := cloudkms.CryptoKeyPrimaryAttestationFormatEnum(n[len("CloudkmsCryptoKeyPrimaryAttestationFormatEnum"):])
		return &e
	}
	return nil
}

// ProtoToCryptoKeyPurposeEnum converts a CryptoKeyPurposeEnum enum from its proto representation.
func ProtoToCloudkmsCryptoKeyPurposeEnum(e cloudkmspb.CloudkmsCryptoKeyPurposeEnum) *cloudkms.CryptoKeyPurposeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := cloudkmspb.CloudkmsCryptoKeyPurposeEnum_name[int32(e)]; ok {
		e := cloudkms.CryptoKeyPurposeEnum(n[len("CloudkmsCryptoKeyPurposeEnum"):])
		return &e
	}
	return nil
}

// ProtoToCryptoKeyVersionTemplateProtectionLevelEnum converts a CryptoKeyVersionTemplateProtectionLevelEnum enum from its proto representation.
func ProtoToCloudkmsCryptoKeyVersionTemplateProtectionLevelEnum(e cloudkmspb.CloudkmsCryptoKeyVersionTemplateProtectionLevelEnum) *cloudkms.CryptoKeyVersionTemplateProtectionLevelEnum {
	if e == 0 {
		return nil
	}
	if n, ok := cloudkmspb.CloudkmsCryptoKeyVersionTemplateProtectionLevelEnum_name[int32(e)]; ok {
		e := cloudkms.CryptoKeyVersionTemplateProtectionLevelEnum(n[len("CloudkmsCryptoKeyVersionTemplateProtectionLevelEnum"):])
		return &e
	}
	return nil
}

// ProtoToCryptoKeyVersionTemplateAlgorithmEnum converts a CryptoKeyVersionTemplateAlgorithmEnum enum from its proto representation.
func ProtoToCloudkmsCryptoKeyVersionTemplateAlgorithmEnum(e cloudkmspb.CloudkmsCryptoKeyVersionTemplateAlgorithmEnum) *cloudkms.CryptoKeyVersionTemplateAlgorithmEnum {
	if e == 0 {
		return nil
	}
	if n, ok := cloudkmspb.CloudkmsCryptoKeyVersionTemplateAlgorithmEnum_name[int32(e)]; ok {
		e := cloudkms.CryptoKeyVersionTemplateAlgorithmEnum(n[len("CloudkmsCryptoKeyVersionTemplateAlgorithmEnum"):])
		return &e
	}
	return nil
}

// ProtoToCryptoKeyPrimary converts a CryptoKeyPrimary object from its proto representation.
func ProtoToCloudkmsCryptoKeyPrimary(p *cloudkmspb.CloudkmsCryptoKeyPrimary) *cloudkms.CryptoKeyPrimary {
	if p == nil {
		return nil
	}
	obj := &cloudkms.CryptoKeyPrimary{
		Name:                           dcl.StringOrNil(p.GetName()),
		State:                          ProtoToCloudkmsCryptoKeyPrimaryStateEnum(p.GetState()),
		ProtectionLevel:                ProtoToCloudkmsCryptoKeyPrimaryProtectionLevelEnum(p.GetProtectionLevel()),
		Algorithm:                      ProtoToCloudkmsCryptoKeyPrimaryAlgorithmEnum(p.GetAlgorithm()),
		Attestation:                    ProtoToCloudkmsCryptoKeyPrimaryAttestation(p.GetAttestation()),
		CreateTime:                     dcl.StringOrNil(p.GetCreateTime()),
		GenerateTime:                   dcl.StringOrNil(p.GetGenerateTime()),
		DestroyTime:                    dcl.StringOrNil(p.GetDestroyTime()),
		DestroyEventTime:               dcl.StringOrNil(p.GetDestroyEventTime()),
		ImportJob:                      dcl.StringOrNil(p.GetImportJob()),
		ImportTime:                     dcl.StringOrNil(p.GetImportTime()),
		ImportFailureReason:            dcl.StringOrNil(p.GetImportFailureReason()),
		ExternalProtectionLevelOptions: ProtoToCloudkmsCryptoKeyPrimaryExternalProtectionLevelOptions(p.GetExternalProtectionLevelOptions()),
		ReimportEligible:               dcl.Bool(p.GetReimportEligible()),
	}
	return obj
}

// ProtoToCryptoKeyPrimaryAttestation converts a CryptoKeyPrimaryAttestation object from its proto representation.
func ProtoToCloudkmsCryptoKeyPrimaryAttestation(p *cloudkmspb.CloudkmsCryptoKeyPrimaryAttestation) *cloudkms.CryptoKeyPrimaryAttestation {
	if p == nil {
		return nil
	}
	obj := &cloudkms.CryptoKeyPrimaryAttestation{
		Format:     ProtoToCloudkmsCryptoKeyPrimaryAttestationFormatEnum(p.GetFormat()),
		Content:    dcl.StringOrNil(p.GetContent()),
		CertChains: ProtoToCloudkmsCryptoKeyPrimaryAttestationCertChains(p.GetCertChains()),
	}
	return obj
}

// ProtoToCryptoKeyPrimaryAttestationCertChains converts a CryptoKeyPrimaryAttestationCertChains object from its proto representation.
func ProtoToCloudkmsCryptoKeyPrimaryAttestationCertChains(p *cloudkmspb.CloudkmsCryptoKeyPrimaryAttestationCertChains) *cloudkms.CryptoKeyPrimaryAttestationCertChains {
	if p == nil {
		return nil
	}
	obj := &cloudkms.CryptoKeyPrimaryAttestationCertChains{}
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
func ProtoToCloudkmsCryptoKeyPrimaryExternalProtectionLevelOptions(p *cloudkmspb.CloudkmsCryptoKeyPrimaryExternalProtectionLevelOptions) *cloudkms.CryptoKeyPrimaryExternalProtectionLevelOptions {
	if p == nil {
		return nil
	}
	obj := &cloudkms.CryptoKeyPrimaryExternalProtectionLevelOptions{
		ExternalKeyUri: dcl.StringOrNil(p.GetExternalKeyUri()),
	}
	return obj
}

// ProtoToCryptoKeyVersionTemplate converts a CryptoKeyVersionTemplate object from its proto representation.
func ProtoToCloudkmsCryptoKeyVersionTemplate(p *cloudkmspb.CloudkmsCryptoKeyVersionTemplate) *cloudkms.CryptoKeyVersionTemplate {
	if p == nil {
		return nil
	}
	obj := &cloudkms.CryptoKeyVersionTemplate{
		ProtectionLevel: ProtoToCloudkmsCryptoKeyVersionTemplateProtectionLevelEnum(p.GetProtectionLevel()),
		Algorithm:       ProtoToCloudkmsCryptoKeyVersionTemplateAlgorithmEnum(p.GetAlgorithm()),
	}
	return obj
}

// ProtoToCryptoKey converts a CryptoKey resource from its proto representation.
func ProtoToCryptoKey(p *cloudkmspb.CloudkmsCryptoKey) *cloudkms.CryptoKey {
	obj := &cloudkms.CryptoKey{
		Name:                     dcl.StringOrNil(p.GetName()),
		Primary:                  ProtoToCloudkmsCryptoKeyPrimary(p.GetPrimary()),
		Purpose:                  ProtoToCloudkmsCryptoKeyPurposeEnum(p.GetPurpose()),
		CreateTime:               dcl.StringOrNil(p.GetCreateTime()),
		NextRotationTime:         dcl.StringOrNil(p.GetNextRotationTime()),
		RotationPeriod:           dcl.StringOrNil(p.GetRotationPeriod()),
		VersionTemplate:          ProtoToCloudkmsCryptoKeyVersionTemplate(p.GetVersionTemplate()),
		ImportOnly:               dcl.Bool(p.GetImportOnly()),
		DestroyScheduledDuration: dcl.StringOrNil(p.GetDestroyScheduledDuration()),
		Project:                  dcl.StringOrNil(p.GetProject()),
		Location:                 dcl.StringOrNil(p.GetLocation()),
		KeyRing:                  dcl.StringOrNil(p.GetKeyRing()),
	}
	return obj
}

// CryptoKeyPrimaryStateEnumToProto converts a CryptoKeyPrimaryStateEnum enum to its proto representation.
func CloudkmsCryptoKeyPrimaryStateEnumToProto(e *cloudkms.CryptoKeyPrimaryStateEnum) cloudkmspb.CloudkmsCryptoKeyPrimaryStateEnum {
	if e == nil {
		return cloudkmspb.CloudkmsCryptoKeyPrimaryStateEnum(0)
	}
	if v, ok := cloudkmspb.CloudkmsCryptoKeyPrimaryStateEnum_value["CryptoKeyPrimaryStateEnum"+string(*e)]; ok {
		return cloudkmspb.CloudkmsCryptoKeyPrimaryStateEnum(v)
	}
	return cloudkmspb.CloudkmsCryptoKeyPrimaryStateEnum(0)
}

// CryptoKeyPrimaryProtectionLevelEnumToProto converts a CryptoKeyPrimaryProtectionLevelEnum enum to its proto representation.
func CloudkmsCryptoKeyPrimaryProtectionLevelEnumToProto(e *cloudkms.CryptoKeyPrimaryProtectionLevelEnum) cloudkmspb.CloudkmsCryptoKeyPrimaryProtectionLevelEnum {
	if e == nil {
		return cloudkmspb.CloudkmsCryptoKeyPrimaryProtectionLevelEnum(0)
	}
	if v, ok := cloudkmspb.CloudkmsCryptoKeyPrimaryProtectionLevelEnum_value["CryptoKeyPrimaryProtectionLevelEnum"+string(*e)]; ok {
		return cloudkmspb.CloudkmsCryptoKeyPrimaryProtectionLevelEnum(v)
	}
	return cloudkmspb.CloudkmsCryptoKeyPrimaryProtectionLevelEnum(0)
}

// CryptoKeyPrimaryAlgorithmEnumToProto converts a CryptoKeyPrimaryAlgorithmEnum enum to its proto representation.
func CloudkmsCryptoKeyPrimaryAlgorithmEnumToProto(e *cloudkms.CryptoKeyPrimaryAlgorithmEnum) cloudkmspb.CloudkmsCryptoKeyPrimaryAlgorithmEnum {
	if e == nil {
		return cloudkmspb.CloudkmsCryptoKeyPrimaryAlgorithmEnum(0)
	}
	if v, ok := cloudkmspb.CloudkmsCryptoKeyPrimaryAlgorithmEnum_value["CryptoKeyPrimaryAlgorithmEnum"+string(*e)]; ok {
		return cloudkmspb.CloudkmsCryptoKeyPrimaryAlgorithmEnum(v)
	}
	return cloudkmspb.CloudkmsCryptoKeyPrimaryAlgorithmEnum(0)
}

// CryptoKeyPrimaryAttestationFormatEnumToProto converts a CryptoKeyPrimaryAttestationFormatEnum enum to its proto representation.
func CloudkmsCryptoKeyPrimaryAttestationFormatEnumToProto(e *cloudkms.CryptoKeyPrimaryAttestationFormatEnum) cloudkmspb.CloudkmsCryptoKeyPrimaryAttestationFormatEnum {
	if e == nil {
		return cloudkmspb.CloudkmsCryptoKeyPrimaryAttestationFormatEnum(0)
	}
	if v, ok := cloudkmspb.CloudkmsCryptoKeyPrimaryAttestationFormatEnum_value["CryptoKeyPrimaryAttestationFormatEnum"+string(*e)]; ok {
		return cloudkmspb.CloudkmsCryptoKeyPrimaryAttestationFormatEnum(v)
	}
	return cloudkmspb.CloudkmsCryptoKeyPrimaryAttestationFormatEnum(0)
}

// CryptoKeyPurposeEnumToProto converts a CryptoKeyPurposeEnum enum to its proto representation.
func CloudkmsCryptoKeyPurposeEnumToProto(e *cloudkms.CryptoKeyPurposeEnum) cloudkmspb.CloudkmsCryptoKeyPurposeEnum {
	if e == nil {
		return cloudkmspb.CloudkmsCryptoKeyPurposeEnum(0)
	}
	if v, ok := cloudkmspb.CloudkmsCryptoKeyPurposeEnum_value["CryptoKeyPurposeEnum"+string(*e)]; ok {
		return cloudkmspb.CloudkmsCryptoKeyPurposeEnum(v)
	}
	return cloudkmspb.CloudkmsCryptoKeyPurposeEnum(0)
}

// CryptoKeyVersionTemplateProtectionLevelEnumToProto converts a CryptoKeyVersionTemplateProtectionLevelEnum enum to its proto representation.
func CloudkmsCryptoKeyVersionTemplateProtectionLevelEnumToProto(e *cloudkms.CryptoKeyVersionTemplateProtectionLevelEnum) cloudkmspb.CloudkmsCryptoKeyVersionTemplateProtectionLevelEnum {
	if e == nil {
		return cloudkmspb.CloudkmsCryptoKeyVersionTemplateProtectionLevelEnum(0)
	}
	if v, ok := cloudkmspb.CloudkmsCryptoKeyVersionTemplateProtectionLevelEnum_value["CryptoKeyVersionTemplateProtectionLevelEnum"+string(*e)]; ok {
		return cloudkmspb.CloudkmsCryptoKeyVersionTemplateProtectionLevelEnum(v)
	}
	return cloudkmspb.CloudkmsCryptoKeyVersionTemplateProtectionLevelEnum(0)
}

// CryptoKeyVersionTemplateAlgorithmEnumToProto converts a CryptoKeyVersionTemplateAlgorithmEnum enum to its proto representation.
func CloudkmsCryptoKeyVersionTemplateAlgorithmEnumToProto(e *cloudkms.CryptoKeyVersionTemplateAlgorithmEnum) cloudkmspb.CloudkmsCryptoKeyVersionTemplateAlgorithmEnum {
	if e == nil {
		return cloudkmspb.CloudkmsCryptoKeyVersionTemplateAlgorithmEnum(0)
	}
	if v, ok := cloudkmspb.CloudkmsCryptoKeyVersionTemplateAlgorithmEnum_value["CryptoKeyVersionTemplateAlgorithmEnum"+string(*e)]; ok {
		return cloudkmspb.CloudkmsCryptoKeyVersionTemplateAlgorithmEnum(v)
	}
	return cloudkmspb.CloudkmsCryptoKeyVersionTemplateAlgorithmEnum(0)
}

// CryptoKeyPrimaryToProto converts a CryptoKeyPrimary object to its proto representation.
func CloudkmsCryptoKeyPrimaryToProto(o *cloudkms.CryptoKeyPrimary) *cloudkmspb.CloudkmsCryptoKeyPrimary {
	if o == nil {
		return nil
	}
	p := &cloudkmspb.CloudkmsCryptoKeyPrimary{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetState(CloudkmsCryptoKeyPrimaryStateEnumToProto(o.State))
	p.SetProtectionLevel(CloudkmsCryptoKeyPrimaryProtectionLevelEnumToProto(o.ProtectionLevel))
	p.SetAlgorithm(CloudkmsCryptoKeyPrimaryAlgorithmEnumToProto(o.Algorithm))
	p.SetAttestation(CloudkmsCryptoKeyPrimaryAttestationToProto(o.Attestation))
	p.SetCreateTime(dcl.ValueOrEmptyString(o.CreateTime))
	p.SetGenerateTime(dcl.ValueOrEmptyString(o.GenerateTime))
	p.SetDestroyTime(dcl.ValueOrEmptyString(o.DestroyTime))
	p.SetDestroyEventTime(dcl.ValueOrEmptyString(o.DestroyEventTime))
	p.SetImportJob(dcl.ValueOrEmptyString(o.ImportJob))
	p.SetImportTime(dcl.ValueOrEmptyString(o.ImportTime))
	p.SetImportFailureReason(dcl.ValueOrEmptyString(o.ImportFailureReason))
	p.SetExternalProtectionLevelOptions(CloudkmsCryptoKeyPrimaryExternalProtectionLevelOptionsToProto(o.ExternalProtectionLevelOptions))
	p.SetReimportEligible(dcl.ValueOrEmptyBool(o.ReimportEligible))
	return p
}

// CryptoKeyPrimaryAttestationToProto converts a CryptoKeyPrimaryAttestation object to its proto representation.
func CloudkmsCryptoKeyPrimaryAttestationToProto(o *cloudkms.CryptoKeyPrimaryAttestation) *cloudkmspb.CloudkmsCryptoKeyPrimaryAttestation {
	if o == nil {
		return nil
	}
	p := &cloudkmspb.CloudkmsCryptoKeyPrimaryAttestation{}
	p.SetFormat(CloudkmsCryptoKeyPrimaryAttestationFormatEnumToProto(o.Format))
	p.SetContent(dcl.ValueOrEmptyString(o.Content))
	p.SetCertChains(CloudkmsCryptoKeyPrimaryAttestationCertChainsToProto(o.CertChains))
	return p
}

// CryptoKeyPrimaryAttestationCertChainsToProto converts a CryptoKeyPrimaryAttestationCertChains object to its proto representation.
func CloudkmsCryptoKeyPrimaryAttestationCertChainsToProto(o *cloudkms.CryptoKeyPrimaryAttestationCertChains) *cloudkmspb.CloudkmsCryptoKeyPrimaryAttestationCertChains {
	if o == nil {
		return nil
	}
	p := &cloudkmspb.CloudkmsCryptoKeyPrimaryAttestationCertChains{}
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
func CloudkmsCryptoKeyPrimaryExternalProtectionLevelOptionsToProto(o *cloudkms.CryptoKeyPrimaryExternalProtectionLevelOptions) *cloudkmspb.CloudkmsCryptoKeyPrimaryExternalProtectionLevelOptions {
	if o == nil {
		return nil
	}
	p := &cloudkmspb.CloudkmsCryptoKeyPrimaryExternalProtectionLevelOptions{}
	p.SetExternalKeyUri(dcl.ValueOrEmptyString(o.ExternalKeyUri))
	return p
}

// CryptoKeyVersionTemplateToProto converts a CryptoKeyVersionTemplate object to its proto representation.
func CloudkmsCryptoKeyVersionTemplateToProto(o *cloudkms.CryptoKeyVersionTemplate) *cloudkmspb.CloudkmsCryptoKeyVersionTemplate {
	if o == nil {
		return nil
	}
	p := &cloudkmspb.CloudkmsCryptoKeyVersionTemplate{}
	p.SetProtectionLevel(CloudkmsCryptoKeyVersionTemplateProtectionLevelEnumToProto(o.ProtectionLevel))
	p.SetAlgorithm(CloudkmsCryptoKeyVersionTemplateAlgorithmEnumToProto(o.Algorithm))
	return p
}

// CryptoKeyToProto converts a CryptoKey resource to its proto representation.
func CryptoKeyToProto(resource *cloudkms.CryptoKey) *cloudkmspb.CloudkmsCryptoKey {
	p := &cloudkmspb.CloudkmsCryptoKey{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetPrimary(CloudkmsCryptoKeyPrimaryToProto(resource.Primary))
	p.SetPurpose(CloudkmsCryptoKeyPurposeEnumToProto(resource.Purpose))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetNextRotationTime(dcl.ValueOrEmptyString(resource.NextRotationTime))
	p.SetRotationPeriod(dcl.ValueOrEmptyString(resource.RotationPeriod))
	p.SetVersionTemplate(CloudkmsCryptoKeyVersionTemplateToProto(resource.VersionTemplate))
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
func (s *CryptoKeyServer) applyCryptoKey(ctx context.Context, c *cloudkms.Client, request *cloudkmspb.ApplyCloudkmsCryptoKeyRequest) (*cloudkmspb.CloudkmsCryptoKey, error) {
	p := ProtoToCryptoKey(request.GetResource())
	res, err := c.ApplyCryptoKey(ctx, p)
	if err != nil {
		return nil, err
	}
	r := CryptoKeyToProto(res)
	return r, nil
}

// applyCloudkmsCryptoKey handles the gRPC request by passing it to the underlying CryptoKey Apply() method.
func (s *CryptoKeyServer) ApplyCloudkmsCryptoKey(ctx context.Context, request *cloudkmspb.ApplyCloudkmsCryptoKeyRequest) (*cloudkmspb.CloudkmsCryptoKey, error) {
	cl, err := createConfigCryptoKey(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyCryptoKey(ctx, cl, request)
}

// DeleteCryptoKey handles the gRPC request by passing it to the underlying CryptoKey Delete() method.
func (s *CryptoKeyServer) DeleteCloudkmsCryptoKey(ctx context.Context, request *cloudkmspb.DeleteCloudkmsCryptoKeyRequest) (*emptypb.Empty, error) {

	return nil, errors.New("no delete endpoint for CryptoKey")

}

// ListCloudkmsCryptoKey handles the gRPC request by passing it to the underlying CryptoKeyList() method.
func (s *CryptoKeyServer) ListCloudkmsCryptoKey(ctx context.Context, request *cloudkmspb.ListCloudkmsCryptoKeyRequest) (*cloudkmspb.ListCloudkmsCryptoKeyResponse, error) {
	cl, err := createConfigCryptoKey(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListCryptoKey(ctx, request.GetProject(), request.GetLocation(), request.GetKeyRing())
	if err != nil {
		return nil, err
	}
	var protos []*cloudkmspb.CloudkmsCryptoKey
	for _, r := range resources.Items {
		rp := CryptoKeyToProto(r)
		protos = append(protos, rp)
	}
	p := &cloudkmspb.ListCloudkmsCryptoKeyResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigCryptoKey(ctx context.Context, service_account_file string) (*cloudkms.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return cloudkms.NewClient(conf), nil
}
