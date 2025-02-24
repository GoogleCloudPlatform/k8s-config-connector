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
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/privateca/alpha/privateca_alpha_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/privateca/alpha"
)

// CertificateAuthorityServer implements the gRPC interface for CertificateAuthority.
type CertificateAuthorityServer struct{}

// ProtoToCertificateAuthorityTypeEnum converts a CertificateAuthorityTypeEnum enum from its proto representation.
func ProtoToPrivatecaAlphaCertificateAuthorityTypeEnum(e alphapb.PrivatecaAlphaCertificateAuthorityTypeEnum) *alpha.CertificateAuthorityTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.PrivatecaAlphaCertificateAuthorityTypeEnum_name[int32(e)]; ok {
		e := alpha.CertificateAuthorityTypeEnum(n[len("PrivatecaAlphaCertificateAuthorityTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToCertificateAuthorityConfigPublicKeyFormatEnum converts a CertificateAuthorityConfigPublicKeyFormatEnum enum from its proto representation.
func ProtoToPrivatecaAlphaCertificateAuthorityConfigPublicKeyFormatEnum(e alphapb.PrivatecaAlphaCertificateAuthorityConfigPublicKeyFormatEnum) *alpha.CertificateAuthorityConfigPublicKeyFormatEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.PrivatecaAlphaCertificateAuthorityConfigPublicKeyFormatEnum_name[int32(e)]; ok {
		e := alpha.CertificateAuthorityConfigPublicKeyFormatEnum(n[len("PrivatecaAlphaCertificateAuthorityConfigPublicKeyFormatEnum"):])
		return &e
	}
	return nil
}

// ProtoToCertificateAuthorityKeySpecAlgorithmEnum converts a CertificateAuthorityKeySpecAlgorithmEnum enum from its proto representation.
func ProtoToPrivatecaAlphaCertificateAuthorityKeySpecAlgorithmEnum(e alphapb.PrivatecaAlphaCertificateAuthorityKeySpecAlgorithmEnum) *alpha.CertificateAuthorityKeySpecAlgorithmEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.PrivatecaAlphaCertificateAuthorityKeySpecAlgorithmEnum_name[int32(e)]; ok {
		e := alpha.CertificateAuthorityKeySpecAlgorithmEnum(n[len("PrivatecaAlphaCertificateAuthorityKeySpecAlgorithmEnum"):])
		return &e
	}
	return nil
}

// ProtoToCertificateAuthorityTierEnum converts a CertificateAuthorityTierEnum enum from its proto representation.
func ProtoToPrivatecaAlphaCertificateAuthorityTierEnum(e alphapb.PrivatecaAlphaCertificateAuthorityTierEnum) *alpha.CertificateAuthorityTierEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.PrivatecaAlphaCertificateAuthorityTierEnum_name[int32(e)]; ok {
		e := alpha.CertificateAuthorityTierEnum(n[len("PrivatecaAlphaCertificateAuthorityTierEnum"):])
		return &e
	}
	return nil
}

// ProtoToCertificateAuthorityStateEnum converts a CertificateAuthorityStateEnum enum from its proto representation.
func ProtoToPrivatecaAlphaCertificateAuthorityStateEnum(e alphapb.PrivatecaAlphaCertificateAuthorityStateEnum) *alpha.CertificateAuthorityStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.PrivatecaAlphaCertificateAuthorityStateEnum_name[int32(e)]; ok {
		e := alpha.CertificateAuthorityStateEnum(n[len("PrivatecaAlphaCertificateAuthorityStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToCertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum converts a CertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum enum from its proto representation.
func ProtoToPrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum(e alphapb.PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum) *alpha.CertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum_name[int32(e)]; ok {
		e := alpha.CertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum(n[len("PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum"):])
		return &e
	}
	return nil
}

// ProtoToCertificateAuthorityConfig converts a CertificateAuthorityConfig object from its proto representation.
func ProtoToPrivatecaAlphaCertificateAuthorityConfig(p *alphapb.PrivatecaAlphaCertificateAuthorityConfig) *alpha.CertificateAuthorityConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.CertificateAuthorityConfig{
		SubjectConfig: ProtoToPrivatecaAlphaCertificateAuthorityConfigSubjectConfig(p.GetSubjectConfig()),
		X509Config:    ProtoToPrivatecaAlphaCertificateAuthorityConfigX509Config(p.GetX509Config()),
		PublicKey:     ProtoToPrivatecaAlphaCertificateAuthorityConfigPublicKey(p.GetPublicKey()),
	}
	return obj
}

// ProtoToCertificateAuthorityConfigSubjectConfig converts a CertificateAuthorityConfigSubjectConfig object from its proto representation.
func ProtoToPrivatecaAlphaCertificateAuthorityConfigSubjectConfig(p *alphapb.PrivatecaAlphaCertificateAuthorityConfigSubjectConfig) *alpha.CertificateAuthorityConfigSubjectConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.CertificateAuthorityConfigSubjectConfig{
		Subject:        ProtoToPrivatecaAlphaCertificateAuthorityConfigSubjectConfigSubject(p.GetSubject()),
		SubjectAltName: ProtoToPrivatecaAlphaCertificateAuthorityConfigSubjectConfigSubjectAltName(p.GetSubjectAltName()),
	}
	return obj
}

// ProtoToCertificateAuthorityConfigSubjectConfigSubject converts a CertificateAuthorityConfigSubjectConfigSubject object from its proto representation.
func ProtoToPrivatecaAlphaCertificateAuthorityConfigSubjectConfigSubject(p *alphapb.PrivatecaAlphaCertificateAuthorityConfigSubjectConfigSubject) *alpha.CertificateAuthorityConfigSubjectConfigSubject {
	if p == nil {
		return nil
	}
	obj := &alpha.CertificateAuthorityConfigSubjectConfigSubject{
		CommonName:         dcl.StringOrNil(p.GetCommonName()),
		CountryCode:        dcl.StringOrNil(p.GetCountryCode()),
		Organization:       dcl.StringOrNil(p.GetOrganization()),
		OrganizationalUnit: dcl.StringOrNil(p.GetOrganizationalUnit()),
		Locality:           dcl.StringOrNil(p.GetLocality()),
		Province:           dcl.StringOrNil(p.GetProvince()),
		StreetAddress:      dcl.StringOrNil(p.GetStreetAddress()),
		PostalCode:         dcl.StringOrNil(p.GetPostalCode()),
	}
	return obj
}

// ProtoToCertificateAuthorityConfigSubjectConfigSubjectAltName converts a CertificateAuthorityConfigSubjectConfigSubjectAltName object from its proto representation.
func ProtoToPrivatecaAlphaCertificateAuthorityConfigSubjectConfigSubjectAltName(p *alphapb.PrivatecaAlphaCertificateAuthorityConfigSubjectConfigSubjectAltName) *alpha.CertificateAuthorityConfigSubjectConfigSubjectAltName {
	if p == nil {
		return nil
	}
	obj := &alpha.CertificateAuthorityConfigSubjectConfigSubjectAltName{}
	for _, r := range p.GetDnsNames() {
		obj.DnsNames = append(obj.DnsNames, r)
	}
	for _, r := range p.GetUris() {
		obj.Uris = append(obj.Uris, r)
	}
	for _, r := range p.GetEmailAddresses() {
		obj.EmailAddresses = append(obj.EmailAddresses, r)
	}
	for _, r := range p.GetIpAddresses() {
		obj.IPAddresses = append(obj.IPAddresses, r)
	}
	for _, r := range p.GetCustomSans() {
		obj.CustomSans = append(obj.CustomSans, *ProtoToPrivatecaAlphaCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans(r))
	}
	return obj
}

// ProtoToCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans converts a CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans object from its proto representation.
func ProtoToPrivatecaAlphaCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans(p *alphapb.PrivatecaAlphaCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans) *alpha.CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans {
	if p == nil {
		return nil
	}
	obj := &alpha.CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans{
		ObjectId: ProtoToPrivatecaAlphaCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId(p.GetObjectId()),
		Critical: dcl.Bool(p.GetCritical()),
		Value:    dcl.StringOrNil(p.GetValue()),
	}
	return obj
}

// ProtoToCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId converts a CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId object from its proto representation.
func ProtoToPrivatecaAlphaCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId(p *alphapb.PrivatecaAlphaCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId) *alpha.CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId {
	if p == nil {
		return nil
	}
	obj := &alpha.CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCertificateAuthorityConfigX509Config converts a CertificateAuthorityConfigX509Config object from its proto representation.
func ProtoToPrivatecaAlphaCertificateAuthorityConfigX509Config(p *alphapb.PrivatecaAlphaCertificateAuthorityConfigX509Config) *alpha.CertificateAuthorityConfigX509Config {
	if p == nil {
		return nil
	}
	obj := &alpha.CertificateAuthorityConfigX509Config{
		KeyUsage:  ProtoToPrivatecaAlphaCertificateAuthorityConfigX509ConfigKeyUsage(p.GetKeyUsage()),
		CaOptions: ProtoToPrivatecaAlphaCertificateAuthorityConfigX509ConfigCaOptions(p.GetCaOptions()),
	}
	for _, r := range p.GetPolicyIds() {
		obj.PolicyIds = append(obj.PolicyIds, *ProtoToPrivatecaAlphaCertificateAuthorityConfigX509ConfigPolicyIds(r))
	}
	for _, r := range p.GetAiaOcspServers() {
		obj.AiaOcspServers = append(obj.AiaOcspServers, r)
	}
	for _, r := range p.GetAdditionalExtensions() {
		obj.AdditionalExtensions = append(obj.AdditionalExtensions, *ProtoToPrivatecaAlphaCertificateAuthorityConfigX509ConfigAdditionalExtensions(r))
	}
	return obj
}

// ProtoToCertificateAuthorityConfigX509ConfigKeyUsage converts a CertificateAuthorityConfigX509ConfigKeyUsage object from its proto representation.
func ProtoToPrivatecaAlphaCertificateAuthorityConfigX509ConfigKeyUsage(p *alphapb.PrivatecaAlphaCertificateAuthorityConfigX509ConfigKeyUsage) *alpha.CertificateAuthorityConfigX509ConfigKeyUsage {
	if p == nil {
		return nil
	}
	obj := &alpha.CertificateAuthorityConfigX509ConfigKeyUsage{
		BaseKeyUsage:     ProtoToPrivatecaAlphaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsage(p.GetBaseKeyUsage()),
		ExtendedKeyUsage: ProtoToPrivatecaAlphaCertificateAuthorityConfigX509ConfigKeyUsageExtendedKeyUsage(p.GetExtendedKeyUsage()),
	}
	for _, r := range p.GetUnknownExtendedKeyUsages() {
		obj.UnknownExtendedKeyUsages = append(obj.UnknownExtendedKeyUsages, *ProtoToPrivatecaAlphaCertificateAuthorityConfigX509ConfigKeyUsageUnknownExtendedKeyUsages(r))
	}
	return obj
}

// ProtoToCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsage converts a CertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsage object from its proto representation.
func ProtoToPrivatecaAlphaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsage(p *alphapb.PrivatecaAlphaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsage) *alpha.CertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsage {
	if p == nil {
		return nil
	}
	obj := &alpha.CertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsage{
		DigitalSignature:  dcl.Bool(p.GetDigitalSignature()),
		ContentCommitment: dcl.Bool(p.GetContentCommitment()),
		KeyEncipherment:   dcl.Bool(p.GetKeyEncipherment()),
		DataEncipherment:  dcl.Bool(p.GetDataEncipherment()),
		KeyAgreement:      dcl.Bool(p.GetKeyAgreement()),
		CertSign:          dcl.Bool(p.GetCertSign()),
		CrlSign:           dcl.Bool(p.GetCrlSign()),
		EncipherOnly:      dcl.Bool(p.GetEncipherOnly()),
		DecipherOnly:      dcl.Bool(p.GetDecipherOnly()),
	}
	return obj
}

// ProtoToCertificateAuthorityConfigX509ConfigKeyUsageExtendedKeyUsage converts a CertificateAuthorityConfigX509ConfigKeyUsageExtendedKeyUsage object from its proto representation.
func ProtoToPrivatecaAlphaCertificateAuthorityConfigX509ConfigKeyUsageExtendedKeyUsage(p *alphapb.PrivatecaAlphaCertificateAuthorityConfigX509ConfigKeyUsageExtendedKeyUsage) *alpha.CertificateAuthorityConfigX509ConfigKeyUsageExtendedKeyUsage {
	if p == nil {
		return nil
	}
	obj := &alpha.CertificateAuthorityConfigX509ConfigKeyUsageExtendedKeyUsage{
		ServerAuth:      dcl.Bool(p.GetServerAuth()),
		ClientAuth:      dcl.Bool(p.GetClientAuth()),
		CodeSigning:     dcl.Bool(p.GetCodeSigning()),
		EmailProtection: dcl.Bool(p.GetEmailProtection()),
		TimeStamping:    dcl.Bool(p.GetTimeStamping()),
		OcspSigning:     dcl.Bool(p.GetOcspSigning()),
	}
	return obj
}

// ProtoToCertificateAuthorityConfigX509ConfigKeyUsageUnknownExtendedKeyUsages converts a CertificateAuthorityConfigX509ConfigKeyUsageUnknownExtendedKeyUsages object from its proto representation.
func ProtoToPrivatecaAlphaCertificateAuthorityConfigX509ConfigKeyUsageUnknownExtendedKeyUsages(p *alphapb.PrivatecaAlphaCertificateAuthorityConfigX509ConfigKeyUsageUnknownExtendedKeyUsages) *alpha.CertificateAuthorityConfigX509ConfigKeyUsageUnknownExtendedKeyUsages {
	if p == nil {
		return nil
	}
	obj := &alpha.CertificateAuthorityConfigX509ConfigKeyUsageUnknownExtendedKeyUsages{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCertificateAuthorityConfigX509ConfigCaOptions converts a CertificateAuthorityConfigX509ConfigCaOptions object from its proto representation.
func ProtoToPrivatecaAlphaCertificateAuthorityConfigX509ConfigCaOptions(p *alphapb.PrivatecaAlphaCertificateAuthorityConfigX509ConfigCaOptions) *alpha.CertificateAuthorityConfigX509ConfigCaOptions {
	if p == nil {
		return nil
	}
	obj := &alpha.CertificateAuthorityConfigX509ConfigCaOptions{
		IsCa:                    dcl.Bool(p.GetIsCa()),
		MaxIssuerPathLength:     dcl.Int64OrNil(p.GetMaxIssuerPathLength()),
		ZeroMaxIssuerPathLength: dcl.Bool(p.GetZeroMaxIssuerPathLength()),
	}
	return obj
}

// ProtoToCertificateAuthorityConfigX509ConfigPolicyIds converts a CertificateAuthorityConfigX509ConfigPolicyIds object from its proto representation.
func ProtoToPrivatecaAlphaCertificateAuthorityConfigX509ConfigPolicyIds(p *alphapb.PrivatecaAlphaCertificateAuthorityConfigX509ConfigPolicyIds) *alpha.CertificateAuthorityConfigX509ConfigPolicyIds {
	if p == nil {
		return nil
	}
	obj := &alpha.CertificateAuthorityConfigX509ConfigPolicyIds{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCertificateAuthorityConfigX509ConfigAdditionalExtensions converts a CertificateAuthorityConfigX509ConfigAdditionalExtensions object from its proto representation.
func ProtoToPrivatecaAlphaCertificateAuthorityConfigX509ConfigAdditionalExtensions(p *alphapb.PrivatecaAlphaCertificateAuthorityConfigX509ConfigAdditionalExtensions) *alpha.CertificateAuthorityConfigX509ConfigAdditionalExtensions {
	if p == nil {
		return nil
	}
	obj := &alpha.CertificateAuthorityConfigX509ConfigAdditionalExtensions{
		ObjectId: ProtoToPrivatecaAlphaCertificateAuthorityConfigX509ConfigAdditionalExtensionsObjectId(p.GetObjectId()),
		Critical: dcl.Bool(p.GetCritical()),
		Value:    dcl.StringOrNil(p.GetValue()),
	}
	return obj
}

// ProtoToCertificateAuthorityConfigX509ConfigAdditionalExtensionsObjectId converts a CertificateAuthorityConfigX509ConfigAdditionalExtensionsObjectId object from its proto representation.
func ProtoToPrivatecaAlphaCertificateAuthorityConfigX509ConfigAdditionalExtensionsObjectId(p *alphapb.PrivatecaAlphaCertificateAuthorityConfigX509ConfigAdditionalExtensionsObjectId) *alpha.CertificateAuthorityConfigX509ConfigAdditionalExtensionsObjectId {
	if p == nil {
		return nil
	}
	obj := &alpha.CertificateAuthorityConfigX509ConfigAdditionalExtensionsObjectId{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCertificateAuthorityConfigPublicKey converts a CertificateAuthorityConfigPublicKey object from its proto representation.
func ProtoToPrivatecaAlphaCertificateAuthorityConfigPublicKey(p *alphapb.PrivatecaAlphaCertificateAuthorityConfigPublicKey) *alpha.CertificateAuthorityConfigPublicKey {
	if p == nil {
		return nil
	}
	obj := &alpha.CertificateAuthorityConfigPublicKey{
		Key:    dcl.StringOrNil(p.GetKey()),
		Format: ProtoToPrivatecaAlphaCertificateAuthorityConfigPublicKeyFormatEnum(p.GetFormat()),
	}
	return obj
}

// ProtoToCertificateAuthorityKeySpec converts a CertificateAuthorityKeySpec object from its proto representation.
func ProtoToPrivatecaAlphaCertificateAuthorityKeySpec(p *alphapb.PrivatecaAlphaCertificateAuthorityKeySpec) *alpha.CertificateAuthorityKeySpec {
	if p == nil {
		return nil
	}
	obj := &alpha.CertificateAuthorityKeySpec{
		CloudKmsKeyVersion: dcl.StringOrNil(p.GetCloudKmsKeyVersion()),
		Algorithm:          ProtoToPrivatecaAlphaCertificateAuthorityKeySpecAlgorithmEnum(p.GetAlgorithm()),
	}
	return obj
}

// ProtoToCertificateAuthoritySubordinateConfig converts a CertificateAuthoritySubordinateConfig object from its proto representation.
func ProtoToPrivatecaAlphaCertificateAuthoritySubordinateConfig(p *alphapb.PrivatecaAlphaCertificateAuthoritySubordinateConfig) *alpha.CertificateAuthoritySubordinateConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.CertificateAuthoritySubordinateConfig{
		CertificateAuthority: dcl.StringOrNil(p.GetCertificateAuthority()),
		PemIssuerChain:       ProtoToPrivatecaAlphaCertificateAuthoritySubordinateConfigPemIssuerChain(p.GetPemIssuerChain()),
	}
	return obj
}

// ProtoToCertificateAuthoritySubordinateConfigPemIssuerChain converts a CertificateAuthoritySubordinateConfigPemIssuerChain object from its proto representation.
func ProtoToPrivatecaAlphaCertificateAuthoritySubordinateConfigPemIssuerChain(p *alphapb.PrivatecaAlphaCertificateAuthoritySubordinateConfigPemIssuerChain) *alpha.CertificateAuthoritySubordinateConfigPemIssuerChain {
	if p == nil {
		return nil
	}
	obj := &alpha.CertificateAuthoritySubordinateConfigPemIssuerChain{}
	for _, r := range p.GetPemCertificates() {
		obj.PemCertificates = append(obj.PemCertificates, r)
	}
	return obj
}

// ProtoToCertificateAuthorityCaCertificateDescriptions converts a CertificateAuthorityCaCertificateDescriptions object from its proto representation.
func ProtoToPrivatecaAlphaCertificateAuthorityCaCertificateDescriptions(p *alphapb.PrivatecaAlphaCertificateAuthorityCaCertificateDescriptions) *alpha.CertificateAuthorityCaCertificateDescriptions {
	if p == nil {
		return nil
	}
	obj := &alpha.CertificateAuthorityCaCertificateDescriptions{
		SubjectDescription: ProtoToPrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsSubjectDescription(p.GetSubjectDescription()),
		X509Description:    ProtoToPrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsX509Description(p.GetX509Description()),
		PublicKey:          ProtoToPrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsPublicKey(p.GetPublicKey()),
		SubjectKeyId:       ProtoToPrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsSubjectKeyId(p.GetSubjectKeyId()),
		AuthorityKeyId:     ProtoToPrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsAuthorityKeyId(p.GetAuthorityKeyId()),
		CertFingerprint:    ProtoToPrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsCertFingerprint(p.GetCertFingerprint()),
	}
	for _, r := range p.GetCrlDistributionPoints() {
		obj.CrlDistributionPoints = append(obj.CrlDistributionPoints, r)
	}
	for _, r := range p.GetAiaIssuingCertificateUrls() {
		obj.AiaIssuingCertificateUrls = append(obj.AiaIssuingCertificateUrls, r)
	}
	return obj
}

// ProtoToCertificateAuthorityCaCertificateDescriptionsSubjectDescription converts a CertificateAuthorityCaCertificateDescriptionsSubjectDescription object from its proto representation.
func ProtoToPrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsSubjectDescription(p *alphapb.PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsSubjectDescription) *alpha.CertificateAuthorityCaCertificateDescriptionsSubjectDescription {
	if p == nil {
		return nil
	}
	obj := &alpha.CertificateAuthorityCaCertificateDescriptionsSubjectDescription{
		Subject:         ProtoToPrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject(p.GetSubject()),
		SubjectAltName:  ProtoToPrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName(p.GetSubjectAltName()),
		HexSerialNumber: dcl.StringOrNil(p.GetHexSerialNumber()),
		Lifetime:        dcl.StringOrNil(p.GetLifetime()),
		NotBeforeTime:   dcl.StringOrNil(p.GetNotBeforeTime()),
		NotAfterTime:    dcl.StringOrNil(p.GetNotAfterTime()),
	}
	return obj
}

// ProtoToCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject converts a CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject object from its proto representation.
func ProtoToPrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject(p *alphapb.PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject) *alpha.CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject {
	if p == nil {
		return nil
	}
	obj := &alpha.CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject{
		CommonName:         dcl.StringOrNil(p.GetCommonName()),
		CountryCode:        dcl.StringOrNil(p.GetCountryCode()),
		Organization:       dcl.StringOrNil(p.GetOrganization()),
		OrganizationalUnit: dcl.StringOrNil(p.GetOrganizationalUnit()),
		Locality:           dcl.StringOrNil(p.GetLocality()),
		Province:           dcl.StringOrNil(p.GetProvince()),
		StreetAddress:      dcl.StringOrNil(p.GetStreetAddress()),
		PostalCode:         dcl.StringOrNil(p.GetPostalCode()),
	}
	return obj
}

// ProtoToCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName converts a CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName object from its proto representation.
func ProtoToPrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName(p *alphapb.PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName) *alpha.CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName {
	if p == nil {
		return nil
	}
	obj := &alpha.CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName{}
	for _, r := range p.GetDnsNames() {
		obj.DnsNames = append(obj.DnsNames, r)
	}
	for _, r := range p.GetUris() {
		obj.Uris = append(obj.Uris, r)
	}
	for _, r := range p.GetEmailAddresses() {
		obj.EmailAddresses = append(obj.EmailAddresses, r)
	}
	for _, r := range p.GetIpAddresses() {
		obj.IPAddresses = append(obj.IPAddresses, r)
	}
	for _, r := range p.GetCustomSans() {
		obj.CustomSans = append(obj.CustomSans, *ProtoToPrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans(r))
	}
	return obj
}

// ProtoToCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans converts a CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans object from its proto representation.
func ProtoToPrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans(p *alphapb.PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans) *alpha.CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans {
	if p == nil {
		return nil
	}
	obj := &alpha.CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans{
		ObjectId: ProtoToPrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId(p.GetObjectId()),
		Critical: dcl.Bool(p.GetCritical()),
		Value:    dcl.StringOrNil(p.GetValue()),
	}
	return obj
}

// ProtoToCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId converts a CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId object from its proto representation.
func ProtoToPrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId(p *alphapb.PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId) *alpha.CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId {
	if p == nil {
		return nil
	}
	obj := &alpha.CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCertificateAuthorityCaCertificateDescriptionsX509Description converts a CertificateAuthorityCaCertificateDescriptionsX509Description object from its proto representation.
func ProtoToPrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsX509Description(p *alphapb.PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsX509Description) *alpha.CertificateAuthorityCaCertificateDescriptionsX509Description {
	if p == nil {
		return nil
	}
	obj := &alpha.CertificateAuthorityCaCertificateDescriptionsX509Description{
		KeyUsage:  ProtoToPrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsage(p.GetKeyUsage()),
		CaOptions: ProtoToPrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsX509DescriptionCaOptions(p.GetCaOptions()),
	}
	for _, r := range p.GetPolicyIds() {
		obj.PolicyIds = append(obj.PolicyIds, *ProtoToPrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsX509DescriptionPolicyIds(r))
	}
	for _, r := range p.GetAiaOcspServers() {
		obj.AiaOcspServers = append(obj.AiaOcspServers, r)
	}
	for _, r := range p.GetAdditionalExtensions() {
		obj.AdditionalExtensions = append(obj.AdditionalExtensions, *ProtoToPrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensions(r))
	}
	return obj
}

// ProtoToCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsage converts a CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsage object from its proto representation.
func ProtoToPrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsage(p *alphapb.PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsage) *alpha.CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsage {
	if p == nil {
		return nil
	}
	obj := &alpha.CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsage{
		BaseKeyUsage:     ProtoToPrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageBaseKeyUsage(p.GetBaseKeyUsage()),
		ExtendedKeyUsage: ProtoToPrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageExtendedKeyUsage(p.GetExtendedKeyUsage()),
	}
	for _, r := range p.GetUnknownExtendedKeyUsages() {
		obj.UnknownExtendedKeyUsages = append(obj.UnknownExtendedKeyUsages, *ProtoToPrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageUnknownExtendedKeyUsages(r))
	}
	return obj
}

// ProtoToCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageBaseKeyUsage converts a CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageBaseKeyUsage object from its proto representation.
func ProtoToPrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageBaseKeyUsage(p *alphapb.PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageBaseKeyUsage) *alpha.CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageBaseKeyUsage {
	if p == nil {
		return nil
	}
	obj := &alpha.CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageBaseKeyUsage{
		DigitalSignature:  dcl.Bool(p.GetDigitalSignature()),
		ContentCommitment: dcl.Bool(p.GetContentCommitment()),
		KeyEncipherment:   dcl.Bool(p.GetKeyEncipherment()),
		DataEncipherment:  dcl.Bool(p.GetDataEncipherment()),
		KeyAgreement:      dcl.Bool(p.GetKeyAgreement()),
		CertSign:          dcl.Bool(p.GetCertSign()),
		CrlSign:           dcl.Bool(p.GetCrlSign()),
		EncipherOnly:      dcl.Bool(p.GetEncipherOnly()),
		DecipherOnly:      dcl.Bool(p.GetDecipherOnly()),
	}
	return obj
}

// ProtoToCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageExtendedKeyUsage converts a CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageExtendedKeyUsage object from its proto representation.
func ProtoToPrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageExtendedKeyUsage(p *alphapb.PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageExtendedKeyUsage) *alpha.CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageExtendedKeyUsage {
	if p == nil {
		return nil
	}
	obj := &alpha.CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageExtendedKeyUsage{
		ServerAuth:      dcl.Bool(p.GetServerAuth()),
		ClientAuth:      dcl.Bool(p.GetClientAuth()),
		CodeSigning:     dcl.Bool(p.GetCodeSigning()),
		EmailProtection: dcl.Bool(p.GetEmailProtection()),
		TimeStamping:    dcl.Bool(p.GetTimeStamping()),
		OcspSigning:     dcl.Bool(p.GetOcspSigning()),
	}
	return obj
}

// ProtoToCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageUnknownExtendedKeyUsages converts a CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageUnknownExtendedKeyUsages object from its proto representation.
func ProtoToPrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageUnknownExtendedKeyUsages(p *alphapb.PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageUnknownExtendedKeyUsages) *alpha.CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageUnknownExtendedKeyUsages {
	if p == nil {
		return nil
	}
	obj := &alpha.CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageUnknownExtendedKeyUsages{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCertificateAuthorityCaCertificateDescriptionsX509DescriptionCaOptions converts a CertificateAuthorityCaCertificateDescriptionsX509DescriptionCaOptions object from its proto representation.
func ProtoToPrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsX509DescriptionCaOptions(p *alphapb.PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsX509DescriptionCaOptions) *alpha.CertificateAuthorityCaCertificateDescriptionsX509DescriptionCaOptions {
	if p == nil {
		return nil
	}
	obj := &alpha.CertificateAuthorityCaCertificateDescriptionsX509DescriptionCaOptions{
		IsCa:                dcl.Bool(p.GetIsCa()),
		MaxIssuerPathLength: dcl.Int64OrNil(p.GetMaxIssuerPathLength()),
	}
	return obj
}

// ProtoToCertificateAuthorityCaCertificateDescriptionsX509DescriptionPolicyIds converts a CertificateAuthorityCaCertificateDescriptionsX509DescriptionPolicyIds object from its proto representation.
func ProtoToPrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsX509DescriptionPolicyIds(p *alphapb.PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsX509DescriptionPolicyIds) *alpha.CertificateAuthorityCaCertificateDescriptionsX509DescriptionPolicyIds {
	if p == nil {
		return nil
	}
	obj := &alpha.CertificateAuthorityCaCertificateDescriptionsX509DescriptionPolicyIds{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensions converts a CertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensions object from its proto representation.
func ProtoToPrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensions(p *alphapb.PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensions) *alpha.CertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensions {
	if p == nil {
		return nil
	}
	obj := &alpha.CertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensions{
		ObjectId: ProtoToPrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensionsObjectId(p.GetObjectId()),
		Critical: dcl.Bool(p.GetCritical()),
		Value:    dcl.StringOrNil(p.GetValue()),
	}
	return obj
}

// ProtoToCertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensionsObjectId converts a CertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensionsObjectId object from its proto representation.
func ProtoToPrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensionsObjectId(p *alphapb.PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensionsObjectId) *alpha.CertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensionsObjectId {
	if p == nil {
		return nil
	}
	obj := &alpha.CertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensionsObjectId{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCertificateAuthorityCaCertificateDescriptionsPublicKey converts a CertificateAuthorityCaCertificateDescriptionsPublicKey object from its proto representation.
func ProtoToPrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsPublicKey(p *alphapb.PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsPublicKey) *alpha.CertificateAuthorityCaCertificateDescriptionsPublicKey {
	if p == nil {
		return nil
	}
	obj := &alpha.CertificateAuthorityCaCertificateDescriptionsPublicKey{
		Key:    dcl.StringOrNil(p.GetKey()),
		Format: ProtoToPrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum(p.GetFormat()),
	}
	return obj
}

// ProtoToCertificateAuthorityCaCertificateDescriptionsSubjectKeyId converts a CertificateAuthorityCaCertificateDescriptionsSubjectKeyId object from its proto representation.
func ProtoToPrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsSubjectKeyId(p *alphapb.PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsSubjectKeyId) *alpha.CertificateAuthorityCaCertificateDescriptionsSubjectKeyId {
	if p == nil {
		return nil
	}
	obj := &alpha.CertificateAuthorityCaCertificateDescriptionsSubjectKeyId{
		KeyId: dcl.StringOrNil(p.GetKeyId()),
	}
	return obj
}

// ProtoToCertificateAuthorityCaCertificateDescriptionsAuthorityKeyId converts a CertificateAuthorityCaCertificateDescriptionsAuthorityKeyId object from its proto representation.
func ProtoToPrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsAuthorityKeyId(p *alphapb.PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsAuthorityKeyId) *alpha.CertificateAuthorityCaCertificateDescriptionsAuthorityKeyId {
	if p == nil {
		return nil
	}
	obj := &alpha.CertificateAuthorityCaCertificateDescriptionsAuthorityKeyId{
		KeyId: dcl.StringOrNil(p.GetKeyId()),
	}
	return obj
}

// ProtoToCertificateAuthorityCaCertificateDescriptionsCertFingerprint converts a CertificateAuthorityCaCertificateDescriptionsCertFingerprint object from its proto representation.
func ProtoToPrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsCertFingerprint(p *alphapb.PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsCertFingerprint) *alpha.CertificateAuthorityCaCertificateDescriptionsCertFingerprint {
	if p == nil {
		return nil
	}
	obj := &alpha.CertificateAuthorityCaCertificateDescriptionsCertFingerprint{
		Sha256Hash: dcl.StringOrNil(p.GetSha256Hash()),
	}
	return obj
}

// ProtoToCertificateAuthorityAccessUrls converts a CertificateAuthorityAccessUrls object from its proto representation.
func ProtoToPrivatecaAlphaCertificateAuthorityAccessUrls(p *alphapb.PrivatecaAlphaCertificateAuthorityAccessUrls) *alpha.CertificateAuthorityAccessUrls {
	if p == nil {
		return nil
	}
	obj := &alpha.CertificateAuthorityAccessUrls{
		CaCertificateAccessUrl: dcl.StringOrNil(p.GetCaCertificateAccessUrl()),
	}
	for _, r := range p.GetCrlAccessUrls() {
		obj.CrlAccessUrls = append(obj.CrlAccessUrls, r)
	}
	return obj
}

// ProtoToCertificateAuthority converts a CertificateAuthority resource from its proto representation.
func ProtoToCertificateAuthority(p *alphapb.PrivatecaAlphaCertificateAuthority) *alpha.CertificateAuthority {
	obj := &alpha.CertificateAuthority{
		Name:              dcl.StringOrNil(p.GetName()),
		Type:              ProtoToPrivatecaAlphaCertificateAuthorityTypeEnum(p.GetType()),
		Config:            ProtoToPrivatecaAlphaCertificateAuthorityConfig(p.GetConfig()),
		Lifetime:          dcl.StringOrNil(p.GetLifetime()),
		KeySpec:           ProtoToPrivatecaAlphaCertificateAuthorityKeySpec(p.GetKeySpec()),
		SubordinateConfig: ProtoToPrivatecaAlphaCertificateAuthoritySubordinateConfig(p.GetSubordinateConfig()),
		Tier:              ProtoToPrivatecaAlphaCertificateAuthorityTierEnum(p.GetTier()),
		State:             ProtoToPrivatecaAlphaCertificateAuthorityStateEnum(p.GetState()),
		GcsBucket:         dcl.StringOrNil(p.GetGcsBucket()),
		AccessUrls:        ProtoToPrivatecaAlphaCertificateAuthorityAccessUrls(p.GetAccessUrls()),
		CreateTime:        dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:        dcl.StringOrNil(p.GetUpdateTime()),
		DeleteTime:        dcl.StringOrNil(p.GetDeleteTime()),
		ExpireTime:        dcl.StringOrNil(p.GetExpireTime()),
		Project:           dcl.StringOrNil(p.GetProject()),
		Location:          dcl.StringOrNil(p.GetLocation()),
		CaPool:            dcl.StringOrNil(p.GetCaPool()),
	}
	for _, r := range p.GetPemCaCertificates() {
		obj.PemCaCertificates = append(obj.PemCaCertificates, r)
	}
	for _, r := range p.GetCaCertificateDescriptions() {
		obj.CaCertificateDescriptions = append(obj.CaCertificateDescriptions, *ProtoToPrivatecaAlphaCertificateAuthorityCaCertificateDescriptions(r))
	}
	return obj
}

// CertificateAuthorityTypeEnumToProto converts a CertificateAuthorityTypeEnum enum to its proto representation.
func PrivatecaAlphaCertificateAuthorityTypeEnumToProto(e *alpha.CertificateAuthorityTypeEnum) alphapb.PrivatecaAlphaCertificateAuthorityTypeEnum {
	if e == nil {
		return alphapb.PrivatecaAlphaCertificateAuthorityTypeEnum(0)
	}
	if v, ok := alphapb.PrivatecaAlphaCertificateAuthorityTypeEnum_value["CertificateAuthorityTypeEnum"+string(*e)]; ok {
		return alphapb.PrivatecaAlphaCertificateAuthorityTypeEnum(v)
	}
	return alphapb.PrivatecaAlphaCertificateAuthorityTypeEnum(0)
}

// CertificateAuthorityConfigPublicKeyFormatEnumToProto converts a CertificateAuthorityConfigPublicKeyFormatEnum enum to its proto representation.
func PrivatecaAlphaCertificateAuthorityConfigPublicKeyFormatEnumToProto(e *alpha.CertificateAuthorityConfigPublicKeyFormatEnum) alphapb.PrivatecaAlphaCertificateAuthorityConfigPublicKeyFormatEnum {
	if e == nil {
		return alphapb.PrivatecaAlphaCertificateAuthorityConfigPublicKeyFormatEnum(0)
	}
	if v, ok := alphapb.PrivatecaAlphaCertificateAuthorityConfigPublicKeyFormatEnum_value["CertificateAuthorityConfigPublicKeyFormatEnum"+string(*e)]; ok {
		return alphapb.PrivatecaAlphaCertificateAuthorityConfigPublicKeyFormatEnum(v)
	}
	return alphapb.PrivatecaAlphaCertificateAuthorityConfigPublicKeyFormatEnum(0)
}

// CertificateAuthorityKeySpecAlgorithmEnumToProto converts a CertificateAuthorityKeySpecAlgorithmEnum enum to its proto representation.
func PrivatecaAlphaCertificateAuthorityKeySpecAlgorithmEnumToProto(e *alpha.CertificateAuthorityKeySpecAlgorithmEnum) alphapb.PrivatecaAlphaCertificateAuthorityKeySpecAlgorithmEnum {
	if e == nil {
		return alphapb.PrivatecaAlphaCertificateAuthorityKeySpecAlgorithmEnum(0)
	}
	if v, ok := alphapb.PrivatecaAlphaCertificateAuthorityKeySpecAlgorithmEnum_value["CertificateAuthorityKeySpecAlgorithmEnum"+string(*e)]; ok {
		return alphapb.PrivatecaAlphaCertificateAuthorityKeySpecAlgorithmEnum(v)
	}
	return alphapb.PrivatecaAlphaCertificateAuthorityKeySpecAlgorithmEnum(0)
}

// CertificateAuthorityTierEnumToProto converts a CertificateAuthorityTierEnum enum to its proto representation.
func PrivatecaAlphaCertificateAuthorityTierEnumToProto(e *alpha.CertificateAuthorityTierEnum) alphapb.PrivatecaAlphaCertificateAuthorityTierEnum {
	if e == nil {
		return alphapb.PrivatecaAlphaCertificateAuthorityTierEnum(0)
	}
	if v, ok := alphapb.PrivatecaAlphaCertificateAuthorityTierEnum_value["CertificateAuthorityTierEnum"+string(*e)]; ok {
		return alphapb.PrivatecaAlphaCertificateAuthorityTierEnum(v)
	}
	return alphapb.PrivatecaAlphaCertificateAuthorityTierEnum(0)
}

// CertificateAuthorityStateEnumToProto converts a CertificateAuthorityStateEnum enum to its proto representation.
func PrivatecaAlphaCertificateAuthorityStateEnumToProto(e *alpha.CertificateAuthorityStateEnum) alphapb.PrivatecaAlphaCertificateAuthorityStateEnum {
	if e == nil {
		return alphapb.PrivatecaAlphaCertificateAuthorityStateEnum(0)
	}
	if v, ok := alphapb.PrivatecaAlphaCertificateAuthorityStateEnum_value["CertificateAuthorityStateEnum"+string(*e)]; ok {
		return alphapb.PrivatecaAlphaCertificateAuthorityStateEnum(v)
	}
	return alphapb.PrivatecaAlphaCertificateAuthorityStateEnum(0)
}

// CertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnumToProto converts a CertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum enum to its proto representation.
func PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnumToProto(e *alpha.CertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum) alphapb.PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum {
	if e == nil {
		return alphapb.PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum(0)
	}
	if v, ok := alphapb.PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum_value["CertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum"+string(*e)]; ok {
		return alphapb.PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum(v)
	}
	return alphapb.PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum(0)
}

// CertificateAuthorityConfigToProto converts a CertificateAuthorityConfig object to its proto representation.
func PrivatecaAlphaCertificateAuthorityConfigToProto(o *alpha.CertificateAuthorityConfig) *alphapb.PrivatecaAlphaCertificateAuthorityConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCertificateAuthorityConfig{}
	p.SetSubjectConfig(PrivatecaAlphaCertificateAuthorityConfigSubjectConfigToProto(o.SubjectConfig))
	p.SetX509Config(PrivatecaAlphaCertificateAuthorityConfigX509ConfigToProto(o.X509Config))
	p.SetPublicKey(PrivatecaAlphaCertificateAuthorityConfigPublicKeyToProto(o.PublicKey))
	return p
}

// CertificateAuthorityConfigSubjectConfigToProto converts a CertificateAuthorityConfigSubjectConfig object to its proto representation.
func PrivatecaAlphaCertificateAuthorityConfigSubjectConfigToProto(o *alpha.CertificateAuthorityConfigSubjectConfig) *alphapb.PrivatecaAlphaCertificateAuthorityConfigSubjectConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCertificateAuthorityConfigSubjectConfig{}
	p.SetSubject(PrivatecaAlphaCertificateAuthorityConfigSubjectConfigSubjectToProto(o.Subject))
	p.SetSubjectAltName(PrivatecaAlphaCertificateAuthorityConfigSubjectConfigSubjectAltNameToProto(o.SubjectAltName))
	return p
}

// CertificateAuthorityConfigSubjectConfigSubjectToProto converts a CertificateAuthorityConfigSubjectConfigSubject object to its proto representation.
func PrivatecaAlphaCertificateAuthorityConfigSubjectConfigSubjectToProto(o *alpha.CertificateAuthorityConfigSubjectConfigSubject) *alphapb.PrivatecaAlphaCertificateAuthorityConfigSubjectConfigSubject {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCertificateAuthorityConfigSubjectConfigSubject{}
	p.SetCommonName(dcl.ValueOrEmptyString(o.CommonName))
	p.SetCountryCode(dcl.ValueOrEmptyString(o.CountryCode))
	p.SetOrganization(dcl.ValueOrEmptyString(o.Organization))
	p.SetOrganizationalUnit(dcl.ValueOrEmptyString(o.OrganizationalUnit))
	p.SetLocality(dcl.ValueOrEmptyString(o.Locality))
	p.SetProvince(dcl.ValueOrEmptyString(o.Province))
	p.SetStreetAddress(dcl.ValueOrEmptyString(o.StreetAddress))
	p.SetPostalCode(dcl.ValueOrEmptyString(o.PostalCode))
	return p
}

// CertificateAuthorityConfigSubjectConfigSubjectAltNameToProto converts a CertificateAuthorityConfigSubjectConfigSubjectAltName object to its proto representation.
func PrivatecaAlphaCertificateAuthorityConfigSubjectConfigSubjectAltNameToProto(o *alpha.CertificateAuthorityConfigSubjectConfigSubjectAltName) *alphapb.PrivatecaAlphaCertificateAuthorityConfigSubjectConfigSubjectAltName {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCertificateAuthorityConfigSubjectConfigSubjectAltName{}
	sDnsNames := make([]string, len(o.DnsNames))
	for i, r := range o.DnsNames {
		sDnsNames[i] = r
	}
	p.SetDnsNames(sDnsNames)
	sUris := make([]string, len(o.Uris))
	for i, r := range o.Uris {
		sUris[i] = r
	}
	p.SetUris(sUris)
	sEmailAddresses := make([]string, len(o.EmailAddresses))
	for i, r := range o.EmailAddresses {
		sEmailAddresses[i] = r
	}
	p.SetEmailAddresses(sEmailAddresses)
	sIPAddresses := make([]string, len(o.IPAddresses))
	for i, r := range o.IPAddresses {
		sIPAddresses[i] = r
	}
	p.SetIpAddresses(sIPAddresses)
	sCustomSans := make([]*alphapb.PrivatecaAlphaCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans, len(o.CustomSans))
	for i, r := range o.CustomSans {
		sCustomSans[i] = PrivatecaAlphaCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansToProto(&r)
	}
	p.SetCustomSans(sCustomSans)
	return p
}

// CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansToProto converts a CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans object to its proto representation.
func PrivatecaAlphaCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansToProto(o *alpha.CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans) *alphapb.PrivatecaAlphaCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans{}
	p.SetObjectId(PrivatecaAlphaCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectIdToProto(o.ObjectId))
	p.SetCritical(dcl.ValueOrEmptyBool(o.Critical))
	p.SetValue(dcl.ValueOrEmptyString(o.Value))
	return p
}

// CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectIdToProto converts a CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId object to its proto representation.
func PrivatecaAlphaCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectIdToProto(o *alpha.CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId) *alphapb.PrivatecaAlphaCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId{}
	sObjectIdPath := make([]int64, len(o.ObjectIdPath))
	for i, r := range o.ObjectIdPath {
		sObjectIdPath[i] = r
	}
	p.SetObjectIdPath(sObjectIdPath)
	return p
}

// CertificateAuthorityConfigX509ConfigToProto converts a CertificateAuthorityConfigX509Config object to its proto representation.
func PrivatecaAlphaCertificateAuthorityConfigX509ConfigToProto(o *alpha.CertificateAuthorityConfigX509Config) *alphapb.PrivatecaAlphaCertificateAuthorityConfigX509Config {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCertificateAuthorityConfigX509Config{}
	p.SetKeyUsage(PrivatecaAlphaCertificateAuthorityConfigX509ConfigKeyUsageToProto(o.KeyUsage))
	p.SetCaOptions(PrivatecaAlphaCertificateAuthorityConfigX509ConfigCaOptionsToProto(o.CaOptions))
	sPolicyIds := make([]*alphapb.PrivatecaAlphaCertificateAuthorityConfigX509ConfigPolicyIds, len(o.PolicyIds))
	for i, r := range o.PolicyIds {
		sPolicyIds[i] = PrivatecaAlphaCertificateAuthorityConfigX509ConfigPolicyIdsToProto(&r)
	}
	p.SetPolicyIds(sPolicyIds)
	sAiaOcspServers := make([]string, len(o.AiaOcspServers))
	for i, r := range o.AiaOcspServers {
		sAiaOcspServers[i] = r
	}
	p.SetAiaOcspServers(sAiaOcspServers)
	sAdditionalExtensions := make([]*alphapb.PrivatecaAlphaCertificateAuthorityConfigX509ConfigAdditionalExtensions, len(o.AdditionalExtensions))
	for i, r := range o.AdditionalExtensions {
		sAdditionalExtensions[i] = PrivatecaAlphaCertificateAuthorityConfigX509ConfigAdditionalExtensionsToProto(&r)
	}
	p.SetAdditionalExtensions(sAdditionalExtensions)
	return p
}

// CertificateAuthorityConfigX509ConfigKeyUsageToProto converts a CertificateAuthorityConfigX509ConfigKeyUsage object to its proto representation.
func PrivatecaAlphaCertificateAuthorityConfigX509ConfigKeyUsageToProto(o *alpha.CertificateAuthorityConfigX509ConfigKeyUsage) *alphapb.PrivatecaAlphaCertificateAuthorityConfigX509ConfigKeyUsage {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCertificateAuthorityConfigX509ConfigKeyUsage{}
	p.SetBaseKeyUsage(PrivatecaAlphaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageToProto(o.BaseKeyUsage))
	p.SetExtendedKeyUsage(PrivatecaAlphaCertificateAuthorityConfigX509ConfigKeyUsageExtendedKeyUsageToProto(o.ExtendedKeyUsage))
	sUnknownExtendedKeyUsages := make([]*alphapb.PrivatecaAlphaCertificateAuthorityConfigX509ConfigKeyUsageUnknownExtendedKeyUsages, len(o.UnknownExtendedKeyUsages))
	for i, r := range o.UnknownExtendedKeyUsages {
		sUnknownExtendedKeyUsages[i] = PrivatecaAlphaCertificateAuthorityConfigX509ConfigKeyUsageUnknownExtendedKeyUsagesToProto(&r)
	}
	p.SetUnknownExtendedKeyUsages(sUnknownExtendedKeyUsages)
	return p
}

// CertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageToProto converts a CertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsage object to its proto representation.
func PrivatecaAlphaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageToProto(o *alpha.CertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsage) *alphapb.PrivatecaAlphaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsage {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsage{}
	p.SetDigitalSignature(dcl.ValueOrEmptyBool(o.DigitalSignature))
	p.SetContentCommitment(dcl.ValueOrEmptyBool(o.ContentCommitment))
	p.SetKeyEncipherment(dcl.ValueOrEmptyBool(o.KeyEncipherment))
	p.SetDataEncipherment(dcl.ValueOrEmptyBool(o.DataEncipherment))
	p.SetKeyAgreement(dcl.ValueOrEmptyBool(o.KeyAgreement))
	p.SetCertSign(dcl.ValueOrEmptyBool(o.CertSign))
	p.SetCrlSign(dcl.ValueOrEmptyBool(o.CrlSign))
	p.SetEncipherOnly(dcl.ValueOrEmptyBool(o.EncipherOnly))
	p.SetDecipherOnly(dcl.ValueOrEmptyBool(o.DecipherOnly))
	return p
}

// CertificateAuthorityConfigX509ConfigKeyUsageExtendedKeyUsageToProto converts a CertificateAuthorityConfigX509ConfigKeyUsageExtendedKeyUsage object to its proto representation.
func PrivatecaAlphaCertificateAuthorityConfigX509ConfigKeyUsageExtendedKeyUsageToProto(o *alpha.CertificateAuthorityConfigX509ConfigKeyUsageExtendedKeyUsage) *alphapb.PrivatecaAlphaCertificateAuthorityConfigX509ConfigKeyUsageExtendedKeyUsage {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCertificateAuthorityConfigX509ConfigKeyUsageExtendedKeyUsage{}
	p.SetServerAuth(dcl.ValueOrEmptyBool(o.ServerAuth))
	p.SetClientAuth(dcl.ValueOrEmptyBool(o.ClientAuth))
	p.SetCodeSigning(dcl.ValueOrEmptyBool(o.CodeSigning))
	p.SetEmailProtection(dcl.ValueOrEmptyBool(o.EmailProtection))
	p.SetTimeStamping(dcl.ValueOrEmptyBool(o.TimeStamping))
	p.SetOcspSigning(dcl.ValueOrEmptyBool(o.OcspSigning))
	return p
}

// CertificateAuthorityConfigX509ConfigKeyUsageUnknownExtendedKeyUsagesToProto converts a CertificateAuthorityConfigX509ConfigKeyUsageUnknownExtendedKeyUsages object to its proto representation.
func PrivatecaAlphaCertificateAuthorityConfigX509ConfigKeyUsageUnknownExtendedKeyUsagesToProto(o *alpha.CertificateAuthorityConfigX509ConfigKeyUsageUnknownExtendedKeyUsages) *alphapb.PrivatecaAlphaCertificateAuthorityConfigX509ConfigKeyUsageUnknownExtendedKeyUsages {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCertificateAuthorityConfigX509ConfigKeyUsageUnknownExtendedKeyUsages{}
	sObjectIdPath := make([]int64, len(o.ObjectIdPath))
	for i, r := range o.ObjectIdPath {
		sObjectIdPath[i] = r
	}
	p.SetObjectIdPath(sObjectIdPath)
	return p
}

// CertificateAuthorityConfigX509ConfigCaOptionsToProto converts a CertificateAuthorityConfigX509ConfigCaOptions object to its proto representation.
func PrivatecaAlphaCertificateAuthorityConfigX509ConfigCaOptionsToProto(o *alpha.CertificateAuthorityConfigX509ConfigCaOptions) *alphapb.PrivatecaAlphaCertificateAuthorityConfigX509ConfigCaOptions {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCertificateAuthorityConfigX509ConfigCaOptions{}
	p.SetIsCa(dcl.ValueOrEmptyBool(o.IsCa))
	p.SetMaxIssuerPathLength(dcl.ValueOrEmptyInt64(o.MaxIssuerPathLength))
	p.SetZeroMaxIssuerPathLength(dcl.ValueOrEmptyBool(o.ZeroMaxIssuerPathLength))
	return p
}

// CertificateAuthorityConfigX509ConfigPolicyIdsToProto converts a CertificateAuthorityConfigX509ConfigPolicyIds object to its proto representation.
func PrivatecaAlphaCertificateAuthorityConfigX509ConfigPolicyIdsToProto(o *alpha.CertificateAuthorityConfigX509ConfigPolicyIds) *alphapb.PrivatecaAlphaCertificateAuthorityConfigX509ConfigPolicyIds {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCertificateAuthorityConfigX509ConfigPolicyIds{}
	sObjectIdPath := make([]int64, len(o.ObjectIdPath))
	for i, r := range o.ObjectIdPath {
		sObjectIdPath[i] = r
	}
	p.SetObjectIdPath(sObjectIdPath)
	return p
}

// CertificateAuthorityConfigX509ConfigAdditionalExtensionsToProto converts a CertificateAuthorityConfigX509ConfigAdditionalExtensions object to its proto representation.
func PrivatecaAlphaCertificateAuthorityConfigX509ConfigAdditionalExtensionsToProto(o *alpha.CertificateAuthorityConfigX509ConfigAdditionalExtensions) *alphapb.PrivatecaAlphaCertificateAuthorityConfigX509ConfigAdditionalExtensions {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCertificateAuthorityConfigX509ConfigAdditionalExtensions{}
	p.SetObjectId(PrivatecaAlphaCertificateAuthorityConfigX509ConfigAdditionalExtensionsObjectIdToProto(o.ObjectId))
	p.SetCritical(dcl.ValueOrEmptyBool(o.Critical))
	p.SetValue(dcl.ValueOrEmptyString(o.Value))
	return p
}

// CertificateAuthorityConfigX509ConfigAdditionalExtensionsObjectIdToProto converts a CertificateAuthorityConfigX509ConfigAdditionalExtensionsObjectId object to its proto representation.
func PrivatecaAlphaCertificateAuthorityConfigX509ConfigAdditionalExtensionsObjectIdToProto(o *alpha.CertificateAuthorityConfigX509ConfigAdditionalExtensionsObjectId) *alphapb.PrivatecaAlphaCertificateAuthorityConfigX509ConfigAdditionalExtensionsObjectId {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCertificateAuthorityConfigX509ConfigAdditionalExtensionsObjectId{}
	sObjectIdPath := make([]int64, len(o.ObjectIdPath))
	for i, r := range o.ObjectIdPath {
		sObjectIdPath[i] = r
	}
	p.SetObjectIdPath(sObjectIdPath)
	return p
}

// CertificateAuthorityConfigPublicKeyToProto converts a CertificateAuthorityConfigPublicKey object to its proto representation.
func PrivatecaAlphaCertificateAuthorityConfigPublicKeyToProto(o *alpha.CertificateAuthorityConfigPublicKey) *alphapb.PrivatecaAlphaCertificateAuthorityConfigPublicKey {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCertificateAuthorityConfigPublicKey{}
	p.SetKey(dcl.ValueOrEmptyString(o.Key))
	p.SetFormat(PrivatecaAlphaCertificateAuthorityConfigPublicKeyFormatEnumToProto(o.Format))
	return p
}

// CertificateAuthorityKeySpecToProto converts a CertificateAuthorityKeySpec object to its proto representation.
func PrivatecaAlphaCertificateAuthorityKeySpecToProto(o *alpha.CertificateAuthorityKeySpec) *alphapb.PrivatecaAlphaCertificateAuthorityKeySpec {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCertificateAuthorityKeySpec{}
	p.SetCloudKmsKeyVersion(dcl.ValueOrEmptyString(o.CloudKmsKeyVersion))
	p.SetAlgorithm(PrivatecaAlphaCertificateAuthorityKeySpecAlgorithmEnumToProto(o.Algorithm))
	return p
}

// CertificateAuthoritySubordinateConfigToProto converts a CertificateAuthoritySubordinateConfig object to its proto representation.
func PrivatecaAlphaCertificateAuthoritySubordinateConfigToProto(o *alpha.CertificateAuthoritySubordinateConfig) *alphapb.PrivatecaAlphaCertificateAuthoritySubordinateConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCertificateAuthoritySubordinateConfig{}
	p.SetCertificateAuthority(dcl.ValueOrEmptyString(o.CertificateAuthority))
	p.SetPemIssuerChain(PrivatecaAlphaCertificateAuthoritySubordinateConfigPemIssuerChainToProto(o.PemIssuerChain))
	return p
}

// CertificateAuthoritySubordinateConfigPemIssuerChainToProto converts a CertificateAuthoritySubordinateConfigPemIssuerChain object to its proto representation.
func PrivatecaAlphaCertificateAuthoritySubordinateConfigPemIssuerChainToProto(o *alpha.CertificateAuthoritySubordinateConfigPemIssuerChain) *alphapb.PrivatecaAlphaCertificateAuthoritySubordinateConfigPemIssuerChain {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCertificateAuthoritySubordinateConfigPemIssuerChain{}
	sPemCertificates := make([]string, len(o.PemCertificates))
	for i, r := range o.PemCertificates {
		sPemCertificates[i] = r
	}
	p.SetPemCertificates(sPemCertificates)
	return p
}

// CertificateAuthorityCaCertificateDescriptionsToProto converts a CertificateAuthorityCaCertificateDescriptions object to its proto representation.
func PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsToProto(o *alpha.CertificateAuthorityCaCertificateDescriptions) *alphapb.PrivatecaAlphaCertificateAuthorityCaCertificateDescriptions {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCertificateAuthorityCaCertificateDescriptions{}
	p.SetSubjectDescription(PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionToProto(o.SubjectDescription))
	p.SetX509Description(PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsX509DescriptionToProto(o.X509Description))
	p.SetPublicKey(PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsPublicKeyToProto(o.PublicKey))
	p.SetSubjectKeyId(PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsSubjectKeyIdToProto(o.SubjectKeyId))
	p.SetAuthorityKeyId(PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsAuthorityKeyIdToProto(o.AuthorityKeyId))
	p.SetCertFingerprint(PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsCertFingerprintToProto(o.CertFingerprint))
	sCrlDistributionPoints := make([]string, len(o.CrlDistributionPoints))
	for i, r := range o.CrlDistributionPoints {
		sCrlDistributionPoints[i] = r
	}
	p.SetCrlDistributionPoints(sCrlDistributionPoints)
	sAiaIssuingCertificateUrls := make([]string, len(o.AiaIssuingCertificateUrls))
	for i, r := range o.AiaIssuingCertificateUrls {
		sAiaIssuingCertificateUrls[i] = r
	}
	p.SetAiaIssuingCertificateUrls(sAiaIssuingCertificateUrls)
	return p
}

// CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionToProto converts a CertificateAuthorityCaCertificateDescriptionsSubjectDescription object to its proto representation.
func PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionToProto(o *alpha.CertificateAuthorityCaCertificateDescriptionsSubjectDescription) *alphapb.PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsSubjectDescription {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsSubjectDescription{}
	p.SetSubject(PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectToProto(o.Subject))
	p.SetSubjectAltName(PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameToProto(o.SubjectAltName))
	p.SetHexSerialNumber(dcl.ValueOrEmptyString(o.HexSerialNumber))
	p.SetLifetime(dcl.ValueOrEmptyString(o.Lifetime))
	p.SetNotBeforeTime(dcl.ValueOrEmptyString(o.NotBeforeTime))
	p.SetNotAfterTime(dcl.ValueOrEmptyString(o.NotAfterTime))
	return p
}

// CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectToProto converts a CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject object to its proto representation.
func PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectToProto(o *alpha.CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject) *alphapb.PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject{}
	p.SetCommonName(dcl.ValueOrEmptyString(o.CommonName))
	p.SetCountryCode(dcl.ValueOrEmptyString(o.CountryCode))
	p.SetOrganization(dcl.ValueOrEmptyString(o.Organization))
	p.SetOrganizationalUnit(dcl.ValueOrEmptyString(o.OrganizationalUnit))
	p.SetLocality(dcl.ValueOrEmptyString(o.Locality))
	p.SetProvince(dcl.ValueOrEmptyString(o.Province))
	p.SetStreetAddress(dcl.ValueOrEmptyString(o.StreetAddress))
	p.SetPostalCode(dcl.ValueOrEmptyString(o.PostalCode))
	return p
}

// CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameToProto converts a CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName object to its proto representation.
func PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameToProto(o *alpha.CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName) *alphapb.PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName{}
	sDnsNames := make([]string, len(o.DnsNames))
	for i, r := range o.DnsNames {
		sDnsNames[i] = r
	}
	p.SetDnsNames(sDnsNames)
	sUris := make([]string, len(o.Uris))
	for i, r := range o.Uris {
		sUris[i] = r
	}
	p.SetUris(sUris)
	sEmailAddresses := make([]string, len(o.EmailAddresses))
	for i, r := range o.EmailAddresses {
		sEmailAddresses[i] = r
	}
	p.SetEmailAddresses(sEmailAddresses)
	sIPAddresses := make([]string, len(o.IPAddresses))
	for i, r := range o.IPAddresses {
		sIPAddresses[i] = r
	}
	p.SetIpAddresses(sIPAddresses)
	sCustomSans := make([]*alphapb.PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans, len(o.CustomSans))
	for i, r := range o.CustomSans {
		sCustomSans[i] = PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansToProto(&r)
	}
	p.SetCustomSans(sCustomSans)
	return p
}

// CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansToProto converts a CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans object to its proto representation.
func PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansToProto(o *alpha.CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans) *alphapb.PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans{}
	p.SetObjectId(PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectIdToProto(o.ObjectId))
	p.SetCritical(dcl.ValueOrEmptyBool(o.Critical))
	p.SetValue(dcl.ValueOrEmptyString(o.Value))
	return p
}

// CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectIdToProto converts a CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId object to its proto representation.
func PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectIdToProto(o *alpha.CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId) *alphapb.PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId{}
	sObjectIdPath := make([]int64, len(o.ObjectIdPath))
	for i, r := range o.ObjectIdPath {
		sObjectIdPath[i] = r
	}
	p.SetObjectIdPath(sObjectIdPath)
	return p
}

// CertificateAuthorityCaCertificateDescriptionsX509DescriptionToProto converts a CertificateAuthorityCaCertificateDescriptionsX509Description object to its proto representation.
func PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsX509DescriptionToProto(o *alpha.CertificateAuthorityCaCertificateDescriptionsX509Description) *alphapb.PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsX509Description {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsX509Description{}
	p.SetKeyUsage(PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageToProto(o.KeyUsage))
	p.SetCaOptions(PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsX509DescriptionCaOptionsToProto(o.CaOptions))
	sPolicyIds := make([]*alphapb.PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsX509DescriptionPolicyIds, len(o.PolicyIds))
	for i, r := range o.PolicyIds {
		sPolicyIds[i] = PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsX509DescriptionPolicyIdsToProto(&r)
	}
	p.SetPolicyIds(sPolicyIds)
	sAiaOcspServers := make([]string, len(o.AiaOcspServers))
	for i, r := range o.AiaOcspServers {
		sAiaOcspServers[i] = r
	}
	p.SetAiaOcspServers(sAiaOcspServers)
	sAdditionalExtensions := make([]*alphapb.PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensions, len(o.AdditionalExtensions))
	for i, r := range o.AdditionalExtensions {
		sAdditionalExtensions[i] = PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensionsToProto(&r)
	}
	p.SetAdditionalExtensions(sAdditionalExtensions)
	return p
}

// CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageToProto converts a CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsage object to its proto representation.
func PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageToProto(o *alpha.CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsage) *alphapb.PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsage {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsage{}
	p.SetBaseKeyUsage(PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageBaseKeyUsageToProto(o.BaseKeyUsage))
	p.SetExtendedKeyUsage(PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageExtendedKeyUsageToProto(o.ExtendedKeyUsage))
	sUnknownExtendedKeyUsages := make([]*alphapb.PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageUnknownExtendedKeyUsages, len(o.UnknownExtendedKeyUsages))
	for i, r := range o.UnknownExtendedKeyUsages {
		sUnknownExtendedKeyUsages[i] = PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageUnknownExtendedKeyUsagesToProto(&r)
	}
	p.SetUnknownExtendedKeyUsages(sUnknownExtendedKeyUsages)
	return p
}

// CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageBaseKeyUsageToProto converts a CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageBaseKeyUsage object to its proto representation.
func PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageBaseKeyUsageToProto(o *alpha.CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageBaseKeyUsage) *alphapb.PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageBaseKeyUsage {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageBaseKeyUsage{}
	p.SetDigitalSignature(dcl.ValueOrEmptyBool(o.DigitalSignature))
	p.SetContentCommitment(dcl.ValueOrEmptyBool(o.ContentCommitment))
	p.SetKeyEncipherment(dcl.ValueOrEmptyBool(o.KeyEncipherment))
	p.SetDataEncipherment(dcl.ValueOrEmptyBool(o.DataEncipherment))
	p.SetKeyAgreement(dcl.ValueOrEmptyBool(o.KeyAgreement))
	p.SetCertSign(dcl.ValueOrEmptyBool(o.CertSign))
	p.SetCrlSign(dcl.ValueOrEmptyBool(o.CrlSign))
	p.SetEncipherOnly(dcl.ValueOrEmptyBool(o.EncipherOnly))
	p.SetDecipherOnly(dcl.ValueOrEmptyBool(o.DecipherOnly))
	return p
}

// CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageExtendedKeyUsageToProto converts a CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageExtendedKeyUsage object to its proto representation.
func PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageExtendedKeyUsageToProto(o *alpha.CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageExtendedKeyUsage) *alphapb.PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageExtendedKeyUsage {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageExtendedKeyUsage{}
	p.SetServerAuth(dcl.ValueOrEmptyBool(o.ServerAuth))
	p.SetClientAuth(dcl.ValueOrEmptyBool(o.ClientAuth))
	p.SetCodeSigning(dcl.ValueOrEmptyBool(o.CodeSigning))
	p.SetEmailProtection(dcl.ValueOrEmptyBool(o.EmailProtection))
	p.SetTimeStamping(dcl.ValueOrEmptyBool(o.TimeStamping))
	p.SetOcspSigning(dcl.ValueOrEmptyBool(o.OcspSigning))
	return p
}

// CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageUnknownExtendedKeyUsagesToProto converts a CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageUnknownExtendedKeyUsages object to its proto representation.
func PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageUnknownExtendedKeyUsagesToProto(o *alpha.CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageUnknownExtendedKeyUsages) *alphapb.PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageUnknownExtendedKeyUsages {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageUnknownExtendedKeyUsages{}
	sObjectIdPath := make([]int64, len(o.ObjectIdPath))
	for i, r := range o.ObjectIdPath {
		sObjectIdPath[i] = r
	}
	p.SetObjectIdPath(sObjectIdPath)
	return p
}

// CertificateAuthorityCaCertificateDescriptionsX509DescriptionCaOptionsToProto converts a CertificateAuthorityCaCertificateDescriptionsX509DescriptionCaOptions object to its proto representation.
func PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsX509DescriptionCaOptionsToProto(o *alpha.CertificateAuthorityCaCertificateDescriptionsX509DescriptionCaOptions) *alphapb.PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsX509DescriptionCaOptions {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsX509DescriptionCaOptions{}
	p.SetIsCa(dcl.ValueOrEmptyBool(o.IsCa))
	p.SetMaxIssuerPathLength(dcl.ValueOrEmptyInt64(o.MaxIssuerPathLength))
	return p
}

// CertificateAuthorityCaCertificateDescriptionsX509DescriptionPolicyIdsToProto converts a CertificateAuthorityCaCertificateDescriptionsX509DescriptionPolicyIds object to its proto representation.
func PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsX509DescriptionPolicyIdsToProto(o *alpha.CertificateAuthorityCaCertificateDescriptionsX509DescriptionPolicyIds) *alphapb.PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsX509DescriptionPolicyIds {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsX509DescriptionPolicyIds{}
	sObjectIdPath := make([]int64, len(o.ObjectIdPath))
	for i, r := range o.ObjectIdPath {
		sObjectIdPath[i] = r
	}
	p.SetObjectIdPath(sObjectIdPath)
	return p
}

// CertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensionsToProto converts a CertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensions object to its proto representation.
func PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensionsToProto(o *alpha.CertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensions) *alphapb.PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensions {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensions{}
	p.SetObjectId(PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensionsObjectIdToProto(o.ObjectId))
	p.SetCritical(dcl.ValueOrEmptyBool(o.Critical))
	p.SetValue(dcl.ValueOrEmptyString(o.Value))
	return p
}

// CertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensionsObjectIdToProto converts a CertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensionsObjectId object to its proto representation.
func PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensionsObjectIdToProto(o *alpha.CertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensionsObjectId) *alphapb.PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensionsObjectId {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensionsObjectId{}
	sObjectIdPath := make([]int64, len(o.ObjectIdPath))
	for i, r := range o.ObjectIdPath {
		sObjectIdPath[i] = r
	}
	p.SetObjectIdPath(sObjectIdPath)
	return p
}

// CertificateAuthorityCaCertificateDescriptionsPublicKeyToProto converts a CertificateAuthorityCaCertificateDescriptionsPublicKey object to its proto representation.
func PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsPublicKeyToProto(o *alpha.CertificateAuthorityCaCertificateDescriptionsPublicKey) *alphapb.PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsPublicKey {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsPublicKey{}
	p.SetKey(dcl.ValueOrEmptyString(o.Key))
	p.SetFormat(PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnumToProto(o.Format))
	return p
}

// CertificateAuthorityCaCertificateDescriptionsSubjectKeyIdToProto converts a CertificateAuthorityCaCertificateDescriptionsSubjectKeyId object to its proto representation.
func PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsSubjectKeyIdToProto(o *alpha.CertificateAuthorityCaCertificateDescriptionsSubjectKeyId) *alphapb.PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsSubjectKeyId {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsSubjectKeyId{}
	p.SetKeyId(dcl.ValueOrEmptyString(o.KeyId))
	return p
}

// CertificateAuthorityCaCertificateDescriptionsAuthorityKeyIdToProto converts a CertificateAuthorityCaCertificateDescriptionsAuthorityKeyId object to its proto representation.
func PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsAuthorityKeyIdToProto(o *alpha.CertificateAuthorityCaCertificateDescriptionsAuthorityKeyId) *alphapb.PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsAuthorityKeyId {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsAuthorityKeyId{}
	p.SetKeyId(dcl.ValueOrEmptyString(o.KeyId))
	return p
}

// CertificateAuthorityCaCertificateDescriptionsCertFingerprintToProto converts a CertificateAuthorityCaCertificateDescriptionsCertFingerprint object to its proto representation.
func PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsCertFingerprintToProto(o *alpha.CertificateAuthorityCaCertificateDescriptionsCertFingerprint) *alphapb.PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsCertFingerprint {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsCertFingerprint{}
	p.SetSha256Hash(dcl.ValueOrEmptyString(o.Sha256Hash))
	return p
}

// CertificateAuthorityAccessUrlsToProto converts a CertificateAuthorityAccessUrls object to its proto representation.
func PrivatecaAlphaCertificateAuthorityAccessUrlsToProto(o *alpha.CertificateAuthorityAccessUrls) *alphapb.PrivatecaAlphaCertificateAuthorityAccessUrls {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCertificateAuthorityAccessUrls{}
	p.SetCaCertificateAccessUrl(dcl.ValueOrEmptyString(o.CaCertificateAccessUrl))
	sCrlAccessUrls := make([]string, len(o.CrlAccessUrls))
	for i, r := range o.CrlAccessUrls {
		sCrlAccessUrls[i] = r
	}
	p.SetCrlAccessUrls(sCrlAccessUrls)
	return p
}

// CertificateAuthorityToProto converts a CertificateAuthority resource to its proto representation.
func CertificateAuthorityToProto(resource *alpha.CertificateAuthority) *alphapb.PrivatecaAlphaCertificateAuthority {
	p := &alphapb.PrivatecaAlphaCertificateAuthority{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetType(PrivatecaAlphaCertificateAuthorityTypeEnumToProto(resource.Type))
	p.SetConfig(PrivatecaAlphaCertificateAuthorityConfigToProto(resource.Config))
	p.SetLifetime(dcl.ValueOrEmptyString(resource.Lifetime))
	p.SetKeySpec(PrivatecaAlphaCertificateAuthorityKeySpecToProto(resource.KeySpec))
	p.SetSubordinateConfig(PrivatecaAlphaCertificateAuthoritySubordinateConfigToProto(resource.SubordinateConfig))
	p.SetTier(PrivatecaAlphaCertificateAuthorityTierEnumToProto(resource.Tier))
	p.SetState(PrivatecaAlphaCertificateAuthorityStateEnumToProto(resource.State))
	p.SetGcsBucket(dcl.ValueOrEmptyString(resource.GcsBucket))
	p.SetAccessUrls(PrivatecaAlphaCertificateAuthorityAccessUrlsToProto(resource.AccessUrls))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetDeleteTime(dcl.ValueOrEmptyString(resource.DeleteTime))
	p.SetExpireTime(dcl.ValueOrEmptyString(resource.ExpireTime))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	p.SetCaPool(dcl.ValueOrEmptyString(resource.CaPool))
	sPemCaCertificates := make([]string, len(resource.PemCaCertificates))
	for i, r := range resource.PemCaCertificates {
		sPemCaCertificates[i] = r
	}
	p.SetPemCaCertificates(sPemCaCertificates)
	sCaCertificateDescriptions := make([]*alphapb.PrivatecaAlphaCertificateAuthorityCaCertificateDescriptions, len(resource.CaCertificateDescriptions))
	for i, r := range resource.CaCertificateDescriptions {
		sCaCertificateDescriptions[i] = PrivatecaAlphaCertificateAuthorityCaCertificateDescriptionsToProto(&r)
	}
	p.SetCaCertificateDescriptions(sCaCertificateDescriptions)
	mLabels := make(map[string]string, len(resource.Labels))
	for k, r := range resource.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)

	return p
}

// applyCertificateAuthority handles the gRPC request by passing it to the underlying CertificateAuthority Apply() method.
func (s *CertificateAuthorityServer) applyCertificateAuthority(ctx context.Context, c *alpha.Client, request *alphapb.ApplyPrivatecaAlphaCertificateAuthorityRequest) (*alphapb.PrivatecaAlphaCertificateAuthority, error) {
	p := ProtoToCertificateAuthority(request.GetResource())
	res, err := c.ApplyCertificateAuthority(ctx, p)
	if err != nil {
		return nil, err
	}
	r := CertificateAuthorityToProto(res)
	return r, nil
}

// applyPrivatecaAlphaCertificateAuthority handles the gRPC request by passing it to the underlying CertificateAuthority Apply() method.
func (s *CertificateAuthorityServer) ApplyPrivatecaAlphaCertificateAuthority(ctx context.Context, request *alphapb.ApplyPrivatecaAlphaCertificateAuthorityRequest) (*alphapb.PrivatecaAlphaCertificateAuthority, error) {
	cl, err := createConfigCertificateAuthority(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyCertificateAuthority(ctx, cl, request)
}

// DeleteCertificateAuthority handles the gRPC request by passing it to the underlying CertificateAuthority Delete() method.
func (s *CertificateAuthorityServer) DeletePrivatecaAlphaCertificateAuthority(ctx context.Context, request *alphapb.DeletePrivatecaAlphaCertificateAuthorityRequest) (*emptypb.Empty, error) {

	cl, err := createConfigCertificateAuthority(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteCertificateAuthority(ctx, ProtoToCertificateAuthority(request.GetResource()))

}

// ListPrivatecaAlphaCertificateAuthority handles the gRPC request by passing it to the underlying CertificateAuthorityList() method.
func (s *CertificateAuthorityServer) ListPrivatecaAlphaCertificateAuthority(ctx context.Context, request *alphapb.ListPrivatecaAlphaCertificateAuthorityRequest) (*alphapb.ListPrivatecaAlphaCertificateAuthorityResponse, error) {
	cl, err := createConfigCertificateAuthority(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListCertificateAuthority(ctx, request.GetProject(), request.GetLocation(), request.GetCaPool())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.PrivatecaAlphaCertificateAuthority
	for _, r := range resources.Items {
		rp := CertificateAuthorityToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListPrivatecaAlphaCertificateAuthorityResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigCertificateAuthority(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
