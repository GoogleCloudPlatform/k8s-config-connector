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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/privateca/beta/privateca_beta_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/privateca/beta"
)

// CertificateAuthorityServer implements the gRPC interface for CertificateAuthority.
type CertificateAuthorityServer struct{}

// ProtoToCertificateAuthorityTypeEnum converts a CertificateAuthorityTypeEnum enum from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityTypeEnum(e betapb.PrivatecaBetaCertificateAuthorityTypeEnum) *beta.CertificateAuthorityTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.PrivatecaBetaCertificateAuthorityTypeEnum_name[int32(e)]; ok {
		e := beta.CertificateAuthorityTypeEnum(n[len("PrivatecaBetaCertificateAuthorityTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToCertificateAuthorityConfigPublicKeyFormatEnum converts a CertificateAuthorityConfigPublicKeyFormatEnum enum from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityConfigPublicKeyFormatEnum(e betapb.PrivatecaBetaCertificateAuthorityConfigPublicKeyFormatEnum) *beta.CertificateAuthorityConfigPublicKeyFormatEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.PrivatecaBetaCertificateAuthorityConfigPublicKeyFormatEnum_name[int32(e)]; ok {
		e := beta.CertificateAuthorityConfigPublicKeyFormatEnum(n[len("PrivatecaBetaCertificateAuthorityConfigPublicKeyFormatEnum"):])
		return &e
	}
	return nil
}

// ProtoToCertificateAuthorityKeySpecAlgorithmEnum converts a CertificateAuthorityKeySpecAlgorithmEnum enum from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityKeySpecAlgorithmEnum(e betapb.PrivatecaBetaCertificateAuthorityKeySpecAlgorithmEnum) *beta.CertificateAuthorityKeySpecAlgorithmEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.PrivatecaBetaCertificateAuthorityKeySpecAlgorithmEnum_name[int32(e)]; ok {
		e := beta.CertificateAuthorityKeySpecAlgorithmEnum(n[len("PrivatecaBetaCertificateAuthorityKeySpecAlgorithmEnum"):])
		return &e
	}
	return nil
}

// ProtoToCertificateAuthorityTierEnum converts a CertificateAuthorityTierEnum enum from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityTierEnum(e betapb.PrivatecaBetaCertificateAuthorityTierEnum) *beta.CertificateAuthorityTierEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.PrivatecaBetaCertificateAuthorityTierEnum_name[int32(e)]; ok {
		e := beta.CertificateAuthorityTierEnum(n[len("PrivatecaBetaCertificateAuthorityTierEnum"):])
		return &e
	}
	return nil
}

// ProtoToCertificateAuthorityStateEnum converts a CertificateAuthorityStateEnum enum from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityStateEnum(e betapb.PrivatecaBetaCertificateAuthorityStateEnum) *beta.CertificateAuthorityStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.PrivatecaBetaCertificateAuthorityStateEnum_name[int32(e)]; ok {
		e := beta.CertificateAuthorityStateEnum(n[len("PrivatecaBetaCertificateAuthorityStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToCertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum converts a CertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum enum from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum(e betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum) *beta.CertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum_name[int32(e)]; ok {
		e := beta.CertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum(n[len("PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum"):])
		return &e
	}
	return nil
}

// ProtoToCertificateAuthorityConfig converts a CertificateAuthorityConfig object from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityConfig(p *betapb.PrivatecaBetaCertificateAuthorityConfig) *beta.CertificateAuthorityConfig {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityConfig{
		SubjectConfig: ProtoToPrivatecaBetaCertificateAuthorityConfigSubjectConfig(p.GetSubjectConfig()),
		X509Config:    ProtoToPrivatecaBetaCertificateAuthorityConfigX509Config(p.GetX509Config()),
		PublicKey:     ProtoToPrivatecaBetaCertificateAuthorityConfigPublicKey(p.GetPublicKey()),
	}
	return obj
}

// ProtoToCertificateAuthorityConfigSubjectConfig converts a CertificateAuthorityConfigSubjectConfig object from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityConfigSubjectConfig(p *betapb.PrivatecaBetaCertificateAuthorityConfigSubjectConfig) *beta.CertificateAuthorityConfigSubjectConfig {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityConfigSubjectConfig{
		Subject:        ProtoToPrivatecaBetaCertificateAuthorityConfigSubjectConfigSubject(p.GetSubject()),
		SubjectAltName: ProtoToPrivatecaBetaCertificateAuthorityConfigSubjectConfigSubjectAltName(p.GetSubjectAltName()),
	}
	return obj
}

// ProtoToCertificateAuthorityConfigSubjectConfigSubject converts a CertificateAuthorityConfigSubjectConfigSubject object from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityConfigSubjectConfigSubject(p *betapb.PrivatecaBetaCertificateAuthorityConfigSubjectConfigSubject) *beta.CertificateAuthorityConfigSubjectConfigSubject {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityConfigSubjectConfigSubject{
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
func ProtoToPrivatecaBetaCertificateAuthorityConfigSubjectConfigSubjectAltName(p *betapb.PrivatecaBetaCertificateAuthorityConfigSubjectConfigSubjectAltName) *beta.CertificateAuthorityConfigSubjectConfigSubjectAltName {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityConfigSubjectConfigSubjectAltName{}
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
		obj.CustomSans = append(obj.CustomSans, *ProtoToPrivatecaBetaCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans(r))
	}
	return obj
}

// ProtoToCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans converts a CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans object from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans(p *betapb.PrivatecaBetaCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans) *beta.CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans{
		ObjectId: ProtoToPrivatecaBetaCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId(p.GetObjectId()),
		Critical: dcl.Bool(p.GetCritical()),
		Value:    dcl.StringOrNil(p.GetValue()),
	}
	return obj
}

// ProtoToCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId converts a CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId object from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId(p *betapb.PrivatecaBetaCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId) *beta.CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCertificateAuthorityConfigX509Config converts a CertificateAuthorityConfigX509Config object from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityConfigX509Config(p *betapb.PrivatecaBetaCertificateAuthorityConfigX509Config) *beta.CertificateAuthorityConfigX509Config {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityConfigX509Config{
		KeyUsage:  ProtoToPrivatecaBetaCertificateAuthorityConfigX509ConfigKeyUsage(p.GetKeyUsage()),
		CaOptions: ProtoToPrivatecaBetaCertificateAuthorityConfigX509ConfigCaOptions(p.GetCaOptions()),
	}
	for _, r := range p.GetPolicyIds() {
		obj.PolicyIds = append(obj.PolicyIds, *ProtoToPrivatecaBetaCertificateAuthorityConfigX509ConfigPolicyIds(r))
	}
	for _, r := range p.GetAiaOcspServers() {
		obj.AiaOcspServers = append(obj.AiaOcspServers, r)
	}
	for _, r := range p.GetAdditionalExtensions() {
		obj.AdditionalExtensions = append(obj.AdditionalExtensions, *ProtoToPrivatecaBetaCertificateAuthorityConfigX509ConfigAdditionalExtensions(r))
	}
	return obj
}

// ProtoToCertificateAuthorityConfigX509ConfigKeyUsage converts a CertificateAuthorityConfigX509ConfigKeyUsage object from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityConfigX509ConfigKeyUsage(p *betapb.PrivatecaBetaCertificateAuthorityConfigX509ConfigKeyUsage) *beta.CertificateAuthorityConfigX509ConfigKeyUsage {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityConfigX509ConfigKeyUsage{
		BaseKeyUsage:     ProtoToPrivatecaBetaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsage(p.GetBaseKeyUsage()),
		ExtendedKeyUsage: ProtoToPrivatecaBetaCertificateAuthorityConfigX509ConfigKeyUsageExtendedKeyUsage(p.GetExtendedKeyUsage()),
	}
	for _, r := range p.GetUnknownExtendedKeyUsages() {
		obj.UnknownExtendedKeyUsages = append(obj.UnknownExtendedKeyUsages, *ProtoToPrivatecaBetaCertificateAuthorityConfigX509ConfigKeyUsageUnknownExtendedKeyUsages(r))
	}
	return obj
}

// ProtoToCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsage converts a CertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsage object from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsage(p *betapb.PrivatecaBetaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsage) *beta.CertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsage {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsage{
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
func ProtoToPrivatecaBetaCertificateAuthorityConfigX509ConfigKeyUsageExtendedKeyUsage(p *betapb.PrivatecaBetaCertificateAuthorityConfigX509ConfigKeyUsageExtendedKeyUsage) *beta.CertificateAuthorityConfigX509ConfigKeyUsageExtendedKeyUsage {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityConfigX509ConfigKeyUsageExtendedKeyUsage{
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
func ProtoToPrivatecaBetaCertificateAuthorityConfigX509ConfigKeyUsageUnknownExtendedKeyUsages(p *betapb.PrivatecaBetaCertificateAuthorityConfigX509ConfigKeyUsageUnknownExtendedKeyUsages) *beta.CertificateAuthorityConfigX509ConfigKeyUsageUnknownExtendedKeyUsages {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityConfigX509ConfigKeyUsageUnknownExtendedKeyUsages{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCertificateAuthorityConfigX509ConfigCaOptions converts a CertificateAuthorityConfigX509ConfigCaOptions object from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityConfigX509ConfigCaOptions(p *betapb.PrivatecaBetaCertificateAuthorityConfigX509ConfigCaOptions) *beta.CertificateAuthorityConfigX509ConfigCaOptions {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityConfigX509ConfigCaOptions{
		IsCa:                    dcl.Bool(p.GetIsCa()),
		MaxIssuerPathLength:     dcl.Int64OrNil(p.GetMaxIssuerPathLength()),
		ZeroMaxIssuerPathLength: dcl.Bool(p.GetZeroMaxIssuerPathLength()),
	}
	return obj
}

// ProtoToCertificateAuthorityConfigX509ConfigPolicyIds converts a CertificateAuthorityConfigX509ConfigPolicyIds object from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityConfigX509ConfigPolicyIds(p *betapb.PrivatecaBetaCertificateAuthorityConfigX509ConfigPolicyIds) *beta.CertificateAuthorityConfigX509ConfigPolicyIds {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityConfigX509ConfigPolicyIds{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCertificateAuthorityConfigX509ConfigAdditionalExtensions converts a CertificateAuthorityConfigX509ConfigAdditionalExtensions object from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityConfigX509ConfigAdditionalExtensions(p *betapb.PrivatecaBetaCertificateAuthorityConfigX509ConfigAdditionalExtensions) *beta.CertificateAuthorityConfigX509ConfigAdditionalExtensions {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityConfigX509ConfigAdditionalExtensions{
		ObjectId: ProtoToPrivatecaBetaCertificateAuthorityConfigX509ConfigAdditionalExtensionsObjectId(p.GetObjectId()),
		Critical: dcl.Bool(p.GetCritical()),
		Value:    dcl.StringOrNil(p.GetValue()),
	}
	return obj
}

// ProtoToCertificateAuthorityConfigX509ConfigAdditionalExtensionsObjectId converts a CertificateAuthorityConfigX509ConfigAdditionalExtensionsObjectId object from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityConfigX509ConfigAdditionalExtensionsObjectId(p *betapb.PrivatecaBetaCertificateAuthorityConfigX509ConfigAdditionalExtensionsObjectId) *beta.CertificateAuthorityConfigX509ConfigAdditionalExtensionsObjectId {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityConfigX509ConfigAdditionalExtensionsObjectId{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCertificateAuthorityConfigPublicKey converts a CertificateAuthorityConfigPublicKey object from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityConfigPublicKey(p *betapb.PrivatecaBetaCertificateAuthorityConfigPublicKey) *beta.CertificateAuthorityConfigPublicKey {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityConfigPublicKey{
		Key:    dcl.StringOrNil(p.GetKey()),
		Format: ProtoToPrivatecaBetaCertificateAuthorityConfigPublicKeyFormatEnum(p.GetFormat()),
	}
	return obj
}

// ProtoToCertificateAuthorityKeySpec converts a CertificateAuthorityKeySpec object from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityKeySpec(p *betapb.PrivatecaBetaCertificateAuthorityKeySpec) *beta.CertificateAuthorityKeySpec {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityKeySpec{
		CloudKmsKeyVersion: dcl.StringOrNil(p.GetCloudKmsKeyVersion()),
		Algorithm:          ProtoToPrivatecaBetaCertificateAuthorityKeySpecAlgorithmEnum(p.GetAlgorithm()),
	}
	return obj
}

// ProtoToCertificateAuthoritySubordinateConfig converts a CertificateAuthoritySubordinateConfig object from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthoritySubordinateConfig(p *betapb.PrivatecaBetaCertificateAuthoritySubordinateConfig) *beta.CertificateAuthoritySubordinateConfig {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthoritySubordinateConfig{
		CertificateAuthority: dcl.StringOrNil(p.GetCertificateAuthority()),
		PemIssuerChain:       ProtoToPrivatecaBetaCertificateAuthoritySubordinateConfigPemIssuerChain(p.GetPemIssuerChain()),
	}
	return obj
}

// ProtoToCertificateAuthoritySubordinateConfigPemIssuerChain converts a CertificateAuthoritySubordinateConfigPemIssuerChain object from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthoritySubordinateConfigPemIssuerChain(p *betapb.PrivatecaBetaCertificateAuthoritySubordinateConfigPemIssuerChain) *beta.CertificateAuthoritySubordinateConfigPemIssuerChain {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthoritySubordinateConfigPemIssuerChain{}
	for _, r := range p.GetPemCertificates() {
		obj.PemCertificates = append(obj.PemCertificates, r)
	}
	return obj
}

// ProtoToCertificateAuthorityCaCertificateDescriptions converts a CertificateAuthorityCaCertificateDescriptions object from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptions(p *betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptions) *beta.CertificateAuthorityCaCertificateDescriptions {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityCaCertificateDescriptions{
		SubjectDescription: ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptionsSubjectDescription(p.GetSubjectDescription()),
		X509Description:    ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509Description(p.GetX509Description()),
		PublicKey:          ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptionsPublicKey(p.GetPublicKey()),
		SubjectKeyId:       ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptionsSubjectKeyId(p.GetSubjectKeyId()),
		AuthorityKeyId:     ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptionsAuthorityKeyId(p.GetAuthorityKeyId()),
		CertFingerprint:    ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptionsCertFingerprint(p.GetCertFingerprint()),
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
func ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptionsSubjectDescription(p *betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsSubjectDescription) *beta.CertificateAuthorityCaCertificateDescriptionsSubjectDescription {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityCaCertificateDescriptionsSubjectDescription{
		Subject:         ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject(p.GetSubject()),
		SubjectAltName:  ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName(p.GetSubjectAltName()),
		HexSerialNumber: dcl.StringOrNil(p.GetHexSerialNumber()),
		Lifetime:        dcl.StringOrNil(p.GetLifetime()),
		NotBeforeTime:   dcl.StringOrNil(p.GetNotBeforeTime()),
		NotAfterTime:    dcl.StringOrNil(p.GetNotAfterTime()),
	}
	return obj
}

// ProtoToCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject converts a CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject object from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject(p *betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject) *beta.CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject{
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
func ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName(p *betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName) *beta.CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName{}
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
		obj.CustomSans = append(obj.CustomSans, *ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans(r))
	}
	return obj
}

// ProtoToCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans converts a CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans object from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans(p *betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans) *beta.CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans{
		ObjectId: ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId(p.GetObjectId()),
		Critical: dcl.Bool(p.GetCritical()),
		Value:    dcl.StringOrNil(p.GetValue()),
	}
	return obj
}

// ProtoToCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId converts a CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId object from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId(p *betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId) *beta.CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCertificateAuthorityCaCertificateDescriptionsX509Description converts a CertificateAuthorityCaCertificateDescriptionsX509Description object from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509Description(p *betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509Description) *beta.CertificateAuthorityCaCertificateDescriptionsX509Description {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityCaCertificateDescriptionsX509Description{
		KeyUsage:  ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsage(p.GetKeyUsage()),
		CaOptions: ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionCaOptions(p.GetCaOptions()),
	}
	for _, r := range p.GetPolicyIds() {
		obj.PolicyIds = append(obj.PolicyIds, *ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionPolicyIds(r))
	}
	for _, r := range p.GetAiaOcspServers() {
		obj.AiaOcspServers = append(obj.AiaOcspServers, r)
	}
	for _, r := range p.GetAdditionalExtensions() {
		obj.AdditionalExtensions = append(obj.AdditionalExtensions, *ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensions(r))
	}
	return obj
}

// ProtoToCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsage converts a CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsage object from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsage(p *betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsage) *beta.CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsage {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsage{
		BaseKeyUsage:     ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageBaseKeyUsage(p.GetBaseKeyUsage()),
		ExtendedKeyUsage: ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageExtendedKeyUsage(p.GetExtendedKeyUsage()),
	}
	for _, r := range p.GetUnknownExtendedKeyUsages() {
		obj.UnknownExtendedKeyUsages = append(obj.UnknownExtendedKeyUsages, *ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageUnknownExtendedKeyUsages(r))
	}
	return obj
}

// ProtoToCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageBaseKeyUsage converts a CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageBaseKeyUsage object from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageBaseKeyUsage(p *betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageBaseKeyUsage) *beta.CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageBaseKeyUsage {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageBaseKeyUsage{
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
func ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageExtendedKeyUsage(p *betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageExtendedKeyUsage) *beta.CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageExtendedKeyUsage {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageExtendedKeyUsage{
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
func ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageUnknownExtendedKeyUsages(p *betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageUnknownExtendedKeyUsages) *beta.CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageUnknownExtendedKeyUsages {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageUnknownExtendedKeyUsages{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCertificateAuthorityCaCertificateDescriptionsX509DescriptionCaOptions converts a CertificateAuthorityCaCertificateDescriptionsX509DescriptionCaOptions object from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionCaOptions(p *betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionCaOptions) *beta.CertificateAuthorityCaCertificateDescriptionsX509DescriptionCaOptions {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityCaCertificateDescriptionsX509DescriptionCaOptions{
		IsCa:                dcl.Bool(p.GetIsCa()),
		MaxIssuerPathLength: dcl.Int64OrNil(p.GetMaxIssuerPathLength()),
	}
	return obj
}

// ProtoToCertificateAuthorityCaCertificateDescriptionsX509DescriptionPolicyIds converts a CertificateAuthorityCaCertificateDescriptionsX509DescriptionPolicyIds object from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionPolicyIds(p *betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionPolicyIds) *beta.CertificateAuthorityCaCertificateDescriptionsX509DescriptionPolicyIds {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityCaCertificateDescriptionsX509DescriptionPolicyIds{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensions converts a CertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensions object from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensions(p *betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensions) *beta.CertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensions {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensions{
		ObjectId: ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensionsObjectId(p.GetObjectId()),
		Critical: dcl.Bool(p.GetCritical()),
		Value:    dcl.StringOrNil(p.GetValue()),
	}
	return obj
}

// ProtoToCertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensionsObjectId converts a CertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensionsObjectId object from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensionsObjectId(p *betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensionsObjectId) *beta.CertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensionsObjectId {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensionsObjectId{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCertificateAuthorityCaCertificateDescriptionsPublicKey converts a CertificateAuthorityCaCertificateDescriptionsPublicKey object from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptionsPublicKey(p *betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsPublicKey) *beta.CertificateAuthorityCaCertificateDescriptionsPublicKey {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityCaCertificateDescriptionsPublicKey{
		Key:    dcl.StringOrNil(p.GetKey()),
		Format: ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum(p.GetFormat()),
	}
	return obj
}

// ProtoToCertificateAuthorityCaCertificateDescriptionsSubjectKeyId converts a CertificateAuthorityCaCertificateDescriptionsSubjectKeyId object from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptionsSubjectKeyId(p *betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsSubjectKeyId) *beta.CertificateAuthorityCaCertificateDescriptionsSubjectKeyId {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityCaCertificateDescriptionsSubjectKeyId{
		KeyId: dcl.StringOrNil(p.GetKeyId()),
	}
	return obj
}

// ProtoToCertificateAuthorityCaCertificateDescriptionsAuthorityKeyId converts a CertificateAuthorityCaCertificateDescriptionsAuthorityKeyId object from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptionsAuthorityKeyId(p *betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsAuthorityKeyId) *beta.CertificateAuthorityCaCertificateDescriptionsAuthorityKeyId {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityCaCertificateDescriptionsAuthorityKeyId{
		KeyId: dcl.StringOrNil(p.GetKeyId()),
	}
	return obj
}

// ProtoToCertificateAuthorityCaCertificateDescriptionsCertFingerprint converts a CertificateAuthorityCaCertificateDescriptionsCertFingerprint object from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptionsCertFingerprint(p *betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsCertFingerprint) *beta.CertificateAuthorityCaCertificateDescriptionsCertFingerprint {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityCaCertificateDescriptionsCertFingerprint{
		Sha256Hash: dcl.StringOrNil(p.GetSha256Hash()),
	}
	return obj
}

// ProtoToCertificateAuthorityAccessUrls converts a CertificateAuthorityAccessUrls object from its proto representation.
func ProtoToPrivatecaBetaCertificateAuthorityAccessUrls(p *betapb.PrivatecaBetaCertificateAuthorityAccessUrls) *beta.CertificateAuthorityAccessUrls {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateAuthorityAccessUrls{
		CaCertificateAccessUrl: dcl.StringOrNil(p.GetCaCertificateAccessUrl()),
	}
	for _, r := range p.GetCrlAccessUrls() {
		obj.CrlAccessUrls = append(obj.CrlAccessUrls, r)
	}
	return obj
}

// ProtoToCertificateAuthority converts a CertificateAuthority resource from its proto representation.
func ProtoToCertificateAuthority(p *betapb.PrivatecaBetaCertificateAuthority) *beta.CertificateAuthority {
	obj := &beta.CertificateAuthority{
		Name:              dcl.StringOrNil(p.GetName()),
		Type:              ProtoToPrivatecaBetaCertificateAuthorityTypeEnum(p.GetType()),
		Config:            ProtoToPrivatecaBetaCertificateAuthorityConfig(p.GetConfig()),
		Lifetime:          dcl.StringOrNil(p.GetLifetime()),
		KeySpec:           ProtoToPrivatecaBetaCertificateAuthorityKeySpec(p.GetKeySpec()),
		SubordinateConfig: ProtoToPrivatecaBetaCertificateAuthoritySubordinateConfig(p.GetSubordinateConfig()),
		Tier:              ProtoToPrivatecaBetaCertificateAuthorityTierEnum(p.GetTier()),
		State:             ProtoToPrivatecaBetaCertificateAuthorityStateEnum(p.GetState()),
		GcsBucket:         dcl.StringOrNil(p.GetGcsBucket()),
		AccessUrls:        ProtoToPrivatecaBetaCertificateAuthorityAccessUrls(p.GetAccessUrls()),
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
		obj.CaCertificateDescriptions = append(obj.CaCertificateDescriptions, *ProtoToPrivatecaBetaCertificateAuthorityCaCertificateDescriptions(r))
	}
	return obj
}

// CertificateAuthorityTypeEnumToProto converts a CertificateAuthorityTypeEnum enum to its proto representation.
func PrivatecaBetaCertificateAuthorityTypeEnumToProto(e *beta.CertificateAuthorityTypeEnum) betapb.PrivatecaBetaCertificateAuthorityTypeEnum {
	if e == nil {
		return betapb.PrivatecaBetaCertificateAuthorityTypeEnum(0)
	}
	if v, ok := betapb.PrivatecaBetaCertificateAuthorityTypeEnum_value["CertificateAuthorityTypeEnum"+string(*e)]; ok {
		return betapb.PrivatecaBetaCertificateAuthorityTypeEnum(v)
	}
	return betapb.PrivatecaBetaCertificateAuthorityTypeEnum(0)
}

// CertificateAuthorityConfigPublicKeyFormatEnumToProto converts a CertificateAuthorityConfigPublicKeyFormatEnum enum to its proto representation.
func PrivatecaBetaCertificateAuthorityConfigPublicKeyFormatEnumToProto(e *beta.CertificateAuthorityConfigPublicKeyFormatEnum) betapb.PrivatecaBetaCertificateAuthorityConfigPublicKeyFormatEnum {
	if e == nil {
		return betapb.PrivatecaBetaCertificateAuthorityConfigPublicKeyFormatEnum(0)
	}
	if v, ok := betapb.PrivatecaBetaCertificateAuthorityConfigPublicKeyFormatEnum_value["CertificateAuthorityConfigPublicKeyFormatEnum"+string(*e)]; ok {
		return betapb.PrivatecaBetaCertificateAuthorityConfigPublicKeyFormatEnum(v)
	}
	return betapb.PrivatecaBetaCertificateAuthorityConfigPublicKeyFormatEnum(0)
}

// CertificateAuthorityKeySpecAlgorithmEnumToProto converts a CertificateAuthorityKeySpecAlgorithmEnum enum to its proto representation.
func PrivatecaBetaCertificateAuthorityKeySpecAlgorithmEnumToProto(e *beta.CertificateAuthorityKeySpecAlgorithmEnum) betapb.PrivatecaBetaCertificateAuthorityKeySpecAlgorithmEnum {
	if e == nil {
		return betapb.PrivatecaBetaCertificateAuthorityKeySpecAlgorithmEnum(0)
	}
	if v, ok := betapb.PrivatecaBetaCertificateAuthorityKeySpecAlgorithmEnum_value["CertificateAuthorityKeySpecAlgorithmEnum"+string(*e)]; ok {
		return betapb.PrivatecaBetaCertificateAuthorityKeySpecAlgorithmEnum(v)
	}
	return betapb.PrivatecaBetaCertificateAuthorityKeySpecAlgorithmEnum(0)
}

// CertificateAuthorityTierEnumToProto converts a CertificateAuthorityTierEnum enum to its proto representation.
func PrivatecaBetaCertificateAuthorityTierEnumToProto(e *beta.CertificateAuthorityTierEnum) betapb.PrivatecaBetaCertificateAuthorityTierEnum {
	if e == nil {
		return betapb.PrivatecaBetaCertificateAuthorityTierEnum(0)
	}
	if v, ok := betapb.PrivatecaBetaCertificateAuthorityTierEnum_value["CertificateAuthorityTierEnum"+string(*e)]; ok {
		return betapb.PrivatecaBetaCertificateAuthorityTierEnum(v)
	}
	return betapb.PrivatecaBetaCertificateAuthorityTierEnum(0)
}

// CertificateAuthorityStateEnumToProto converts a CertificateAuthorityStateEnum enum to its proto representation.
func PrivatecaBetaCertificateAuthorityStateEnumToProto(e *beta.CertificateAuthorityStateEnum) betapb.PrivatecaBetaCertificateAuthorityStateEnum {
	if e == nil {
		return betapb.PrivatecaBetaCertificateAuthorityStateEnum(0)
	}
	if v, ok := betapb.PrivatecaBetaCertificateAuthorityStateEnum_value["CertificateAuthorityStateEnum"+string(*e)]; ok {
		return betapb.PrivatecaBetaCertificateAuthorityStateEnum(v)
	}
	return betapb.PrivatecaBetaCertificateAuthorityStateEnum(0)
}

// CertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnumToProto converts a CertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum enum to its proto representation.
func PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnumToProto(e *beta.CertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum) betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum {
	if e == nil {
		return betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum(0)
	}
	if v, ok := betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum_value["CertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum"+string(*e)]; ok {
		return betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum(v)
	}
	return betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum(0)
}

// CertificateAuthorityConfigToProto converts a CertificateAuthorityConfig object to its proto representation.
func PrivatecaBetaCertificateAuthorityConfigToProto(o *beta.CertificateAuthorityConfig) *betapb.PrivatecaBetaCertificateAuthorityConfig {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityConfig{}
	p.SetSubjectConfig(PrivatecaBetaCertificateAuthorityConfigSubjectConfigToProto(o.SubjectConfig))
	p.SetX509Config(PrivatecaBetaCertificateAuthorityConfigX509ConfigToProto(o.X509Config))
	p.SetPublicKey(PrivatecaBetaCertificateAuthorityConfigPublicKeyToProto(o.PublicKey))
	return p
}

// CertificateAuthorityConfigSubjectConfigToProto converts a CertificateAuthorityConfigSubjectConfig object to its proto representation.
func PrivatecaBetaCertificateAuthorityConfigSubjectConfigToProto(o *beta.CertificateAuthorityConfigSubjectConfig) *betapb.PrivatecaBetaCertificateAuthorityConfigSubjectConfig {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityConfigSubjectConfig{}
	p.SetSubject(PrivatecaBetaCertificateAuthorityConfigSubjectConfigSubjectToProto(o.Subject))
	p.SetSubjectAltName(PrivatecaBetaCertificateAuthorityConfigSubjectConfigSubjectAltNameToProto(o.SubjectAltName))
	return p
}

// CertificateAuthorityConfigSubjectConfigSubjectToProto converts a CertificateAuthorityConfigSubjectConfigSubject object to its proto representation.
func PrivatecaBetaCertificateAuthorityConfigSubjectConfigSubjectToProto(o *beta.CertificateAuthorityConfigSubjectConfigSubject) *betapb.PrivatecaBetaCertificateAuthorityConfigSubjectConfigSubject {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityConfigSubjectConfigSubject{}
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
func PrivatecaBetaCertificateAuthorityConfigSubjectConfigSubjectAltNameToProto(o *beta.CertificateAuthorityConfigSubjectConfigSubjectAltName) *betapb.PrivatecaBetaCertificateAuthorityConfigSubjectConfigSubjectAltName {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityConfigSubjectConfigSubjectAltName{}
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
	sCustomSans := make([]*betapb.PrivatecaBetaCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans, len(o.CustomSans))
	for i, r := range o.CustomSans {
		sCustomSans[i] = PrivatecaBetaCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansToProto(&r)
	}
	p.SetCustomSans(sCustomSans)
	return p
}

// CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansToProto converts a CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans object to its proto representation.
func PrivatecaBetaCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansToProto(o *beta.CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans) *betapb.PrivatecaBetaCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans{}
	p.SetObjectId(PrivatecaBetaCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectIdToProto(o.ObjectId))
	p.SetCritical(dcl.ValueOrEmptyBool(o.Critical))
	p.SetValue(dcl.ValueOrEmptyString(o.Value))
	return p
}

// CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectIdToProto converts a CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId object to its proto representation.
func PrivatecaBetaCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectIdToProto(o *beta.CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId) *betapb.PrivatecaBetaCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId{}
	sObjectIdPath := make([]int64, len(o.ObjectIdPath))
	for i, r := range o.ObjectIdPath {
		sObjectIdPath[i] = r
	}
	p.SetObjectIdPath(sObjectIdPath)
	return p
}

// CertificateAuthorityConfigX509ConfigToProto converts a CertificateAuthorityConfigX509Config object to its proto representation.
func PrivatecaBetaCertificateAuthorityConfigX509ConfigToProto(o *beta.CertificateAuthorityConfigX509Config) *betapb.PrivatecaBetaCertificateAuthorityConfigX509Config {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityConfigX509Config{}
	p.SetKeyUsage(PrivatecaBetaCertificateAuthorityConfigX509ConfigKeyUsageToProto(o.KeyUsage))
	p.SetCaOptions(PrivatecaBetaCertificateAuthorityConfigX509ConfigCaOptionsToProto(o.CaOptions))
	sPolicyIds := make([]*betapb.PrivatecaBetaCertificateAuthorityConfigX509ConfigPolicyIds, len(o.PolicyIds))
	for i, r := range o.PolicyIds {
		sPolicyIds[i] = PrivatecaBetaCertificateAuthorityConfigX509ConfigPolicyIdsToProto(&r)
	}
	p.SetPolicyIds(sPolicyIds)
	sAiaOcspServers := make([]string, len(o.AiaOcspServers))
	for i, r := range o.AiaOcspServers {
		sAiaOcspServers[i] = r
	}
	p.SetAiaOcspServers(sAiaOcspServers)
	sAdditionalExtensions := make([]*betapb.PrivatecaBetaCertificateAuthorityConfigX509ConfigAdditionalExtensions, len(o.AdditionalExtensions))
	for i, r := range o.AdditionalExtensions {
		sAdditionalExtensions[i] = PrivatecaBetaCertificateAuthorityConfigX509ConfigAdditionalExtensionsToProto(&r)
	}
	p.SetAdditionalExtensions(sAdditionalExtensions)
	return p
}

// CertificateAuthorityConfigX509ConfigKeyUsageToProto converts a CertificateAuthorityConfigX509ConfigKeyUsage object to its proto representation.
func PrivatecaBetaCertificateAuthorityConfigX509ConfigKeyUsageToProto(o *beta.CertificateAuthorityConfigX509ConfigKeyUsage) *betapb.PrivatecaBetaCertificateAuthorityConfigX509ConfigKeyUsage {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityConfigX509ConfigKeyUsage{}
	p.SetBaseKeyUsage(PrivatecaBetaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageToProto(o.BaseKeyUsage))
	p.SetExtendedKeyUsage(PrivatecaBetaCertificateAuthorityConfigX509ConfigKeyUsageExtendedKeyUsageToProto(o.ExtendedKeyUsage))
	sUnknownExtendedKeyUsages := make([]*betapb.PrivatecaBetaCertificateAuthorityConfigX509ConfigKeyUsageUnknownExtendedKeyUsages, len(o.UnknownExtendedKeyUsages))
	for i, r := range o.UnknownExtendedKeyUsages {
		sUnknownExtendedKeyUsages[i] = PrivatecaBetaCertificateAuthorityConfigX509ConfigKeyUsageUnknownExtendedKeyUsagesToProto(&r)
	}
	p.SetUnknownExtendedKeyUsages(sUnknownExtendedKeyUsages)
	return p
}

// CertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageToProto converts a CertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsage object to its proto representation.
func PrivatecaBetaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageToProto(o *beta.CertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsage) *betapb.PrivatecaBetaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsage {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsage{}
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
func PrivatecaBetaCertificateAuthorityConfigX509ConfigKeyUsageExtendedKeyUsageToProto(o *beta.CertificateAuthorityConfigX509ConfigKeyUsageExtendedKeyUsage) *betapb.PrivatecaBetaCertificateAuthorityConfigX509ConfigKeyUsageExtendedKeyUsage {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityConfigX509ConfigKeyUsageExtendedKeyUsage{}
	p.SetServerAuth(dcl.ValueOrEmptyBool(o.ServerAuth))
	p.SetClientAuth(dcl.ValueOrEmptyBool(o.ClientAuth))
	p.SetCodeSigning(dcl.ValueOrEmptyBool(o.CodeSigning))
	p.SetEmailProtection(dcl.ValueOrEmptyBool(o.EmailProtection))
	p.SetTimeStamping(dcl.ValueOrEmptyBool(o.TimeStamping))
	p.SetOcspSigning(dcl.ValueOrEmptyBool(o.OcspSigning))
	return p
}

// CertificateAuthorityConfigX509ConfigKeyUsageUnknownExtendedKeyUsagesToProto converts a CertificateAuthorityConfigX509ConfigKeyUsageUnknownExtendedKeyUsages object to its proto representation.
func PrivatecaBetaCertificateAuthorityConfigX509ConfigKeyUsageUnknownExtendedKeyUsagesToProto(o *beta.CertificateAuthorityConfigX509ConfigKeyUsageUnknownExtendedKeyUsages) *betapb.PrivatecaBetaCertificateAuthorityConfigX509ConfigKeyUsageUnknownExtendedKeyUsages {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityConfigX509ConfigKeyUsageUnknownExtendedKeyUsages{}
	sObjectIdPath := make([]int64, len(o.ObjectIdPath))
	for i, r := range o.ObjectIdPath {
		sObjectIdPath[i] = r
	}
	p.SetObjectIdPath(sObjectIdPath)
	return p
}

// CertificateAuthorityConfigX509ConfigCaOptionsToProto converts a CertificateAuthorityConfigX509ConfigCaOptions object to its proto representation.
func PrivatecaBetaCertificateAuthorityConfigX509ConfigCaOptionsToProto(o *beta.CertificateAuthorityConfigX509ConfigCaOptions) *betapb.PrivatecaBetaCertificateAuthorityConfigX509ConfigCaOptions {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityConfigX509ConfigCaOptions{}
	p.SetIsCa(dcl.ValueOrEmptyBool(o.IsCa))
	p.SetMaxIssuerPathLength(dcl.ValueOrEmptyInt64(o.MaxIssuerPathLength))
	p.SetZeroMaxIssuerPathLength(dcl.ValueOrEmptyBool(o.ZeroMaxIssuerPathLength))
	return p
}

// CertificateAuthorityConfigX509ConfigPolicyIdsToProto converts a CertificateAuthorityConfigX509ConfigPolicyIds object to its proto representation.
func PrivatecaBetaCertificateAuthorityConfigX509ConfigPolicyIdsToProto(o *beta.CertificateAuthorityConfigX509ConfigPolicyIds) *betapb.PrivatecaBetaCertificateAuthorityConfigX509ConfigPolicyIds {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityConfigX509ConfigPolicyIds{}
	sObjectIdPath := make([]int64, len(o.ObjectIdPath))
	for i, r := range o.ObjectIdPath {
		sObjectIdPath[i] = r
	}
	p.SetObjectIdPath(sObjectIdPath)
	return p
}

// CertificateAuthorityConfigX509ConfigAdditionalExtensionsToProto converts a CertificateAuthorityConfigX509ConfigAdditionalExtensions object to its proto representation.
func PrivatecaBetaCertificateAuthorityConfigX509ConfigAdditionalExtensionsToProto(o *beta.CertificateAuthorityConfigX509ConfigAdditionalExtensions) *betapb.PrivatecaBetaCertificateAuthorityConfigX509ConfigAdditionalExtensions {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityConfigX509ConfigAdditionalExtensions{}
	p.SetObjectId(PrivatecaBetaCertificateAuthorityConfigX509ConfigAdditionalExtensionsObjectIdToProto(o.ObjectId))
	p.SetCritical(dcl.ValueOrEmptyBool(o.Critical))
	p.SetValue(dcl.ValueOrEmptyString(o.Value))
	return p
}

// CertificateAuthorityConfigX509ConfigAdditionalExtensionsObjectIdToProto converts a CertificateAuthorityConfigX509ConfigAdditionalExtensionsObjectId object to its proto representation.
func PrivatecaBetaCertificateAuthorityConfigX509ConfigAdditionalExtensionsObjectIdToProto(o *beta.CertificateAuthorityConfigX509ConfigAdditionalExtensionsObjectId) *betapb.PrivatecaBetaCertificateAuthorityConfigX509ConfigAdditionalExtensionsObjectId {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityConfigX509ConfigAdditionalExtensionsObjectId{}
	sObjectIdPath := make([]int64, len(o.ObjectIdPath))
	for i, r := range o.ObjectIdPath {
		sObjectIdPath[i] = r
	}
	p.SetObjectIdPath(sObjectIdPath)
	return p
}

// CertificateAuthorityConfigPublicKeyToProto converts a CertificateAuthorityConfigPublicKey object to its proto representation.
func PrivatecaBetaCertificateAuthorityConfigPublicKeyToProto(o *beta.CertificateAuthorityConfigPublicKey) *betapb.PrivatecaBetaCertificateAuthorityConfigPublicKey {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityConfigPublicKey{}
	p.SetKey(dcl.ValueOrEmptyString(o.Key))
	p.SetFormat(PrivatecaBetaCertificateAuthorityConfigPublicKeyFormatEnumToProto(o.Format))
	return p
}

// CertificateAuthorityKeySpecToProto converts a CertificateAuthorityKeySpec object to its proto representation.
func PrivatecaBetaCertificateAuthorityKeySpecToProto(o *beta.CertificateAuthorityKeySpec) *betapb.PrivatecaBetaCertificateAuthorityKeySpec {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityKeySpec{}
	p.SetCloudKmsKeyVersion(dcl.ValueOrEmptyString(o.CloudKmsKeyVersion))
	p.SetAlgorithm(PrivatecaBetaCertificateAuthorityKeySpecAlgorithmEnumToProto(o.Algorithm))
	return p
}

// CertificateAuthoritySubordinateConfigToProto converts a CertificateAuthoritySubordinateConfig object to its proto representation.
func PrivatecaBetaCertificateAuthoritySubordinateConfigToProto(o *beta.CertificateAuthoritySubordinateConfig) *betapb.PrivatecaBetaCertificateAuthoritySubordinateConfig {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthoritySubordinateConfig{}
	p.SetCertificateAuthority(dcl.ValueOrEmptyString(o.CertificateAuthority))
	p.SetPemIssuerChain(PrivatecaBetaCertificateAuthoritySubordinateConfigPemIssuerChainToProto(o.PemIssuerChain))
	return p
}

// CertificateAuthoritySubordinateConfigPemIssuerChainToProto converts a CertificateAuthoritySubordinateConfigPemIssuerChain object to its proto representation.
func PrivatecaBetaCertificateAuthoritySubordinateConfigPemIssuerChainToProto(o *beta.CertificateAuthoritySubordinateConfigPemIssuerChain) *betapb.PrivatecaBetaCertificateAuthoritySubordinateConfigPemIssuerChain {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthoritySubordinateConfigPemIssuerChain{}
	sPemCertificates := make([]string, len(o.PemCertificates))
	for i, r := range o.PemCertificates {
		sPemCertificates[i] = r
	}
	p.SetPemCertificates(sPemCertificates)
	return p
}

// CertificateAuthorityCaCertificateDescriptionsToProto converts a CertificateAuthorityCaCertificateDescriptions object to its proto representation.
func PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsToProto(o *beta.CertificateAuthorityCaCertificateDescriptions) *betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptions {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptions{}
	p.SetSubjectDescription(PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionToProto(o.SubjectDescription))
	p.SetX509Description(PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionToProto(o.X509Description))
	p.SetPublicKey(PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsPublicKeyToProto(o.PublicKey))
	p.SetSubjectKeyId(PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsSubjectKeyIdToProto(o.SubjectKeyId))
	p.SetAuthorityKeyId(PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsAuthorityKeyIdToProto(o.AuthorityKeyId))
	p.SetCertFingerprint(PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsCertFingerprintToProto(o.CertFingerprint))
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
func PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionToProto(o *beta.CertificateAuthorityCaCertificateDescriptionsSubjectDescription) *betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsSubjectDescription {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsSubjectDescription{}
	p.SetSubject(PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectToProto(o.Subject))
	p.SetSubjectAltName(PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameToProto(o.SubjectAltName))
	p.SetHexSerialNumber(dcl.ValueOrEmptyString(o.HexSerialNumber))
	p.SetLifetime(dcl.ValueOrEmptyString(o.Lifetime))
	p.SetNotBeforeTime(dcl.ValueOrEmptyString(o.NotBeforeTime))
	p.SetNotAfterTime(dcl.ValueOrEmptyString(o.NotAfterTime))
	return p
}

// CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectToProto converts a CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject object to its proto representation.
func PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectToProto(o *beta.CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject) *betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject{}
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
func PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameToProto(o *beta.CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName) *betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName{}
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
	sCustomSans := make([]*betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans, len(o.CustomSans))
	for i, r := range o.CustomSans {
		sCustomSans[i] = PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansToProto(&r)
	}
	p.SetCustomSans(sCustomSans)
	return p
}

// CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansToProto converts a CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans object to its proto representation.
func PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansToProto(o *beta.CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans) *betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans{}
	p.SetObjectId(PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectIdToProto(o.ObjectId))
	p.SetCritical(dcl.ValueOrEmptyBool(o.Critical))
	p.SetValue(dcl.ValueOrEmptyString(o.Value))
	return p
}

// CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectIdToProto converts a CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId object to its proto representation.
func PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectIdToProto(o *beta.CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId) *betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId{}
	sObjectIdPath := make([]int64, len(o.ObjectIdPath))
	for i, r := range o.ObjectIdPath {
		sObjectIdPath[i] = r
	}
	p.SetObjectIdPath(sObjectIdPath)
	return p
}

// CertificateAuthorityCaCertificateDescriptionsX509DescriptionToProto converts a CertificateAuthorityCaCertificateDescriptionsX509Description object to its proto representation.
func PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionToProto(o *beta.CertificateAuthorityCaCertificateDescriptionsX509Description) *betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509Description {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509Description{}
	p.SetKeyUsage(PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageToProto(o.KeyUsage))
	p.SetCaOptions(PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionCaOptionsToProto(o.CaOptions))
	sPolicyIds := make([]*betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionPolicyIds, len(o.PolicyIds))
	for i, r := range o.PolicyIds {
		sPolicyIds[i] = PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionPolicyIdsToProto(&r)
	}
	p.SetPolicyIds(sPolicyIds)
	sAiaOcspServers := make([]string, len(o.AiaOcspServers))
	for i, r := range o.AiaOcspServers {
		sAiaOcspServers[i] = r
	}
	p.SetAiaOcspServers(sAiaOcspServers)
	sAdditionalExtensions := make([]*betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensions, len(o.AdditionalExtensions))
	for i, r := range o.AdditionalExtensions {
		sAdditionalExtensions[i] = PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensionsToProto(&r)
	}
	p.SetAdditionalExtensions(sAdditionalExtensions)
	return p
}

// CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageToProto converts a CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsage object to its proto representation.
func PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageToProto(o *beta.CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsage) *betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsage {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsage{}
	p.SetBaseKeyUsage(PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageBaseKeyUsageToProto(o.BaseKeyUsage))
	p.SetExtendedKeyUsage(PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageExtendedKeyUsageToProto(o.ExtendedKeyUsage))
	sUnknownExtendedKeyUsages := make([]*betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageUnknownExtendedKeyUsages, len(o.UnknownExtendedKeyUsages))
	for i, r := range o.UnknownExtendedKeyUsages {
		sUnknownExtendedKeyUsages[i] = PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageUnknownExtendedKeyUsagesToProto(&r)
	}
	p.SetUnknownExtendedKeyUsages(sUnknownExtendedKeyUsages)
	return p
}

// CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageBaseKeyUsageToProto converts a CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageBaseKeyUsage object to its proto representation.
func PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageBaseKeyUsageToProto(o *beta.CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageBaseKeyUsage) *betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageBaseKeyUsage {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageBaseKeyUsage{}
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
func PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageExtendedKeyUsageToProto(o *beta.CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageExtendedKeyUsage) *betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageExtendedKeyUsage {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageExtendedKeyUsage{}
	p.SetServerAuth(dcl.ValueOrEmptyBool(o.ServerAuth))
	p.SetClientAuth(dcl.ValueOrEmptyBool(o.ClientAuth))
	p.SetCodeSigning(dcl.ValueOrEmptyBool(o.CodeSigning))
	p.SetEmailProtection(dcl.ValueOrEmptyBool(o.EmailProtection))
	p.SetTimeStamping(dcl.ValueOrEmptyBool(o.TimeStamping))
	p.SetOcspSigning(dcl.ValueOrEmptyBool(o.OcspSigning))
	return p
}

// CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageUnknownExtendedKeyUsagesToProto converts a CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageUnknownExtendedKeyUsages object to its proto representation.
func PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageUnknownExtendedKeyUsagesToProto(o *beta.CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageUnknownExtendedKeyUsages) *betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageUnknownExtendedKeyUsages {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageUnknownExtendedKeyUsages{}
	sObjectIdPath := make([]int64, len(o.ObjectIdPath))
	for i, r := range o.ObjectIdPath {
		sObjectIdPath[i] = r
	}
	p.SetObjectIdPath(sObjectIdPath)
	return p
}

// CertificateAuthorityCaCertificateDescriptionsX509DescriptionCaOptionsToProto converts a CertificateAuthorityCaCertificateDescriptionsX509DescriptionCaOptions object to its proto representation.
func PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionCaOptionsToProto(o *beta.CertificateAuthorityCaCertificateDescriptionsX509DescriptionCaOptions) *betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionCaOptions {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionCaOptions{}
	p.SetIsCa(dcl.ValueOrEmptyBool(o.IsCa))
	p.SetMaxIssuerPathLength(dcl.ValueOrEmptyInt64(o.MaxIssuerPathLength))
	return p
}

// CertificateAuthorityCaCertificateDescriptionsX509DescriptionPolicyIdsToProto converts a CertificateAuthorityCaCertificateDescriptionsX509DescriptionPolicyIds object to its proto representation.
func PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionPolicyIdsToProto(o *beta.CertificateAuthorityCaCertificateDescriptionsX509DescriptionPolicyIds) *betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionPolicyIds {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionPolicyIds{}
	sObjectIdPath := make([]int64, len(o.ObjectIdPath))
	for i, r := range o.ObjectIdPath {
		sObjectIdPath[i] = r
	}
	p.SetObjectIdPath(sObjectIdPath)
	return p
}

// CertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensionsToProto converts a CertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensions object to its proto representation.
func PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensionsToProto(o *beta.CertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensions) *betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensions {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensions{}
	p.SetObjectId(PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensionsObjectIdToProto(o.ObjectId))
	p.SetCritical(dcl.ValueOrEmptyBool(o.Critical))
	p.SetValue(dcl.ValueOrEmptyString(o.Value))
	return p
}

// CertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensionsObjectIdToProto converts a CertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensionsObjectId object to its proto representation.
func PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensionsObjectIdToProto(o *beta.CertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensionsObjectId) *betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensionsObjectId {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensionsObjectId{}
	sObjectIdPath := make([]int64, len(o.ObjectIdPath))
	for i, r := range o.ObjectIdPath {
		sObjectIdPath[i] = r
	}
	p.SetObjectIdPath(sObjectIdPath)
	return p
}

// CertificateAuthorityCaCertificateDescriptionsPublicKeyToProto converts a CertificateAuthorityCaCertificateDescriptionsPublicKey object to its proto representation.
func PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsPublicKeyToProto(o *beta.CertificateAuthorityCaCertificateDescriptionsPublicKey) *betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsPublicKey {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsPublicKey{}
	p.SetKey(dcl.ValueOrEmptyString(o.Key))
	p.SetFormat(PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnumToProto(o.Format))
	return p
}

// CertificateAuthorityCaCertificateDescriptionsSubjectKeyIdToProto converts a CertificateAuthorityCaCertificateDescriptionsSubjectKeyId object to its proto representation.
func PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsSubjectKeyIdToProto(o *beta.CertificateAuthorityCaCertificateDescriptionsSubjectKeyId) *betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsSubjectKeyId {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsSubjectKeyId{}
	p.SetKeyId(dcl.ValueOrEmptyString(o.KeyId))
	return p
}

// CertificateAuthorityCaCertificateDescriptionsAuthorityKeyIdToProto converts a CertificateAuthorityCaCertificateDescriptionsAuthorityKeyId object to its proto representation.
func PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsAuthorityKeyIdToProto(o *beta.CertificateAuthorityCaCertificateDescriptionsAuthorityKeyId) *betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsAuthorityKeyId {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsAuthorityKeyId{}
	p.SetKeyId(dcl.ValueOrEmptyString(o.KeyId))
	return p
}

// CertificateAuthorityCaCertificateDescriptionsCertFingerprintToProto converts a CertificateAuthorityCaCertificateDescriptionsCertFingerprint object to its proto representation.
func PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsCertFingerprintToProto(o *beta.CertificateAuthorityCaCertificateDescriptionsCertFingerprint) *betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsCertFingerprint {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsCertFingerprint{}
	p.SetSha256Hash(dcl.ValueOrEmptyString(o.Sha256Hash))
	return p
}

// CertificateAuthorityAccessUrlsToProto converts a CertificateAuthorityAccessUrls object to its proto representation.
func PrivatecaBetaCertificateAuthorityAccessUrlsToProto(o *beta.CertificateAuthorityAccessUrls) *betapb.PrivatecaBetaCertificateAuthorityAccessUrls {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateAuthorityAccessUrls{}
	p.SetCaCertificateAccessUrl(dcl.ValueOrEmptyString(o.CaCertificateAccessUrl))
	sCrlAccessUrls := make([]string, len(o.CrlAccessUrls))
	for i, r := range o.CrlAccessUrls {
		sCrlAccessUrls[i] = r
	}
	p.SetCrlAccessUrls(sCrlAccessUrls)
	return p
}

// CertificateAuthorityToProto converts a CertificateAuthority resource to its proto representation.
func CertificateAuthorityToProto(resource *beta.CertificateAuthority) *betapb.PrivatecaBetaCertificateAuthority {
	p := &betapb.PrivatecaBetaCertificateAuthority{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetType(PrivatecaBetaCertificateAuthorityTypeEnumToProto(resource.Type))
	p.SetConfig(PrivatecaBetaCertificateAuthorityConfigToProto(resource.Config))
	p.SetLifetime(dcl.ValueOrEmptyString(resource.Lifetime))
	p.SetKeySpec(PrivatecaBetaCertificateAuthorityKeySpecToProto(resource.KeySpec))
	p.SetSubordinateConfig(PrivatecaBetaCertificateAuthoritySubordinateConfigToProto(resource.SubordinateConfig))
	p.SetTier(PrivatecaBetaCertificateAuthorityTierEnumToProto(resource.Tier))
	p.SetState(PrivatecaBetaCertificateAuthorityStateEnumToProto(resource.State))
	p.SetGcsBucket(dcl.ValueOrEmptyString(resource.GcsBucket))
	p.SetAccessUrls(PrivatecaBetaCertificateAuthorityAccessUrlsToProto(resource.AccessUrls))
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
	sCaCertificateDescriptions := make([]*betapb.PrivatecaBetaCertificateAuthorityCaCertificateDescriptions, len(resource.CaCertificateDescriptions))
	for i, r := range resource.CaCertificateDescriptions {
		sCaCertificateDescriptions[i] = PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsToProto(&r)
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
func (s *CertificateAuthorityServer) applyCertificateAuthority(ctx context.Context, c *beta.Client, request *betapb.ApplyPrivatecaBetaCertificateAuthorityRequest) (*betapb.PrivatecaBetaCertificateAuthority, error) {
	p := ProtoToCertificateAuthority(request.GetResource())
	res, err := c.ApplyCertificateAuthority(ctx, p)
	if err != nil {
		return nil, err
	}
	r := CertificateAuthorityToProto(res)
	return r, nil
}

// applyPrivatecaBetaCertificateAuthority handles the gRPC request by passing it to the underlying CertificateAuthority Apply() method.
func (s *CertificateAuthorityServer) ApplyPrivatecaBetaCertificateAuthority(ctx context.Context, request *betapb.ApplyPrivatecaBetaCertificateAuthorityRequest) (*betapb.PrivatecaBetaCertificateAuthority, error) {
	cl, err := createConfigCertificateAuthority(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyCertificateAuthority(ctx, cl, request)
}

// DeleteCertificateAuthority handles the gRPC request by passing it to the underlying CertificateAuthority Delete() method.
func (s *CertificateAuthorityServer) DeletePrivatecaBetaCertificateAuthority(ctx context.Context, request *betapb.DeletePrivatecaBetaCertificateAuthorityRequest) (*emptypb.Empty, error) {

	cl, err := createConfigCertificateAuthority(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteCertificateAuthority(ctx, ProtoToCertificateAuthority(request.GetResource()))

}

// ListPrivatecaBetaCertificateAuthority handles the gRPC request by passing it to the underlying CertificateAuthorityList() method.
func (s *CertificateAuthorityServer) ListPrivatecaBetaCertificateAuthority(ctx context.Context, request *betapb.ListPrivatecaBetaCertificateAuthorityRequest) (*betapb.ListPrivatecaBetaCertificateAuthorityResponse, error) {
	cl, err := createConfigCertificateAuthority(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListCertificateAuthority(ctx, request.GetProject(), request.GetLocation(), request.GetCaPool())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.PrivatecaBetaCertificateAuthority
	for _, r := range resources.Items {
		rp := CertificateAuthorityToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListPrivatecaBetaCertificateAuthorityResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigCertificateAuthority(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
