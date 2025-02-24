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
	privatecapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/privateca/privateca_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/privateca"
)

// CertificateAuthorityServer implements the gRPC interface for CertificateAuthority.
type CertificateAuthorityServer struct{}

// ProtoToCertificateAuthorityTypeEnum converts a CertificateAuthorityTypeEnum enum from its proto representation.
func ProtoToPrivatecaCertificateAuthorityTypeEnum(e privatecapb.PrivatecaCertificateAuthorityTypeEnum) *privateca.CertificateAuthorityTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := privatecapb.PrivatecaCertificateAuthorityTypeEnum_name[int32(e)]; ok {
		e := privateca.CertificateAuthorityTypeEnum(n[len("PrivatecaCertificateAuthorityTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToCertificateAuthorityConfigPublicKeyFormatEnum converts a CertificateAuthorityConfigPublicKeyFormatEnum enum from its proto representation.
func ProtoToPrivatecaCertificateAuthorityConfigPublicKeyFormatEnum(e privatecapb.PrivatecaCertificateAuthorityConfigPublicKeyFormatEnum) *privateca.CertificateAuthorityConfigPublicKeyFormatEnum {
	if e == 0 {
		return nil
	}
	if n, ok := privatecapb.PrivatecaCertificateAuthorityConfigPublicKeyFormatEnum_name[int32(e)]; ok {
		e := privateca.CertificateAuthorityConfigPublicKeyFormatEnum(n[len("PrivatecaCertificateAuthorityConfigPublicKeyFormatEnum"):])
		return &e
	}
	return nil
}

// ProtoToCertificateAuthorityKeySpecAlgorithmEnum converts a CertificateAuthorityKeySpecAlgorithmEnum enum from its proto representation.
func ProtoToPrivatecaCertificateAuthorityKeySpecAlgorithmEnum(e privatecapb.PrivatecaCertificateAuthorityKeySpecAlgorithmEnum) *privateca.CertificateAuthorityKeySpecAlgorithmEnum {
	if e == 0 {
		return nil
	}
	if n, ok := privatecapb.PrivatecaCertificateAuthorityKeySpecAlgorithmEnum_name[int32(e)]; ok {
		e := privateca.CertificateAuthorityKeySpecAlgorithmEnum(n[len("PrivatecaCertificateAuthorityKeySpecAlgorithmEnum"):])
		return &e
	}
	return nil
}

// ProtoToCertificateAuthorityTierEnum converts a CertificateAuthorityTierEnum enum from its proto representation.
func ProtoToPrivatecaCertificateAuthorityTierEnum(e privatecapb.PrivatecaCertificateAuthorityTierEnum) *privateca.CertificateAuthorityTierEnum {
	if e == 0 {
		return nil
	}
	if n, ok := privatecapb.PrivatecaCertificateAuthorityTierEnum_name[int32(e)]; ok {
		e := privateca.CertificateAuthorityTierEnum(n[len("PrivatecaCertificateAuthorityTierEnum"):])
		return &e
	}
	return nil
}

// ProtoToCertificateAuthorityStateEnum converts a CertificateAuthorityStateEnum enum from its proto representation.
func ProtoToPrivatecaCertificateAuthorityStateEnum(e privatecapb.PrivatecaCertificateAuthorityStateEnum) *privateca.CertificateAuthorityStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := privatecapb.PrivatecaCertificateAuthorityStateEnum_name[int32(e)]; ok {
		e := privateca.CertificateAuthorityStateEnum(n[len("PrivatecaCertificateAuthorityStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToCertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum converts a CertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum enum from its proto representation.
func ProtoToPrivatecaCertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum(e privatecapb.PrivatecaCertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum) *privateca.CertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum {
	if e == 0 {
		return nil
	}
	if n, ok := privatecapb.PrivatecaCertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum_name[int32(e)]; ok {
		e := privateca.CertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum(n[len("PrivatecaCertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum"):])
		return &e
	}
	return nil
}

// ProtoToCertificateAuthorityConfig converts a CertificateAuthorityConfig object from its proto representation.
func ProtoToPrivatecaCertificateAuthorityConfig(p *privatecapb.PrivatecaCertificateAuthorityConfig) *privateca.CertificateAuthorityConfig {
	if p == nil {
		return nil
	}
	obj := &privateca.CertificateAuthorityConfig{
		SubjectConfig: ProtoToPrivatecaCertificateAuthorityConfigSubjectConfig(p.GetSubjectConfig()),
		X509Config:    ProtoToPrivatecaCertificateAuthorityConfigX509Config(p.GetX509Config()),
		PublicKey:     ProtoToPrivatecaCertificateAuthorityConfigPublicKey(p.GetPublicKey()),
	}
	return obj
}

// ProtoToCertificateAuthorityConfigSubjectConfig converts a CertificateAuthorityConfigSubjectConfig object from its proto representation.
func ProtoToPrivatecaCertificateAuthorityConfigSubjectConfig(p *privatecapb.PrivatecaCertificateAuthorityConfigSubjectConfig) *privateca.CertificateAuthorityConfigSubjectConfig {
	if p == nil {
		return nil
	}
	obj := &privateca.CertificateAuthorityConfigSubjectConfig{
		Subject:        ProtoToPrivatecaCertificateAuthorityConfigSubjectConfigSubject(p.GetSubject()),
		SubjectAltName: ProtoToPrivatecaCertificateAuthorityConfigSubjectConfigSubjectAltName(p.GetSubjectAltName()),
	}
	return obj
}

// ProtoToCertificateAuthorityConfigSubjectConfigSubject converts a CertificateAuthorityConfigSubjectConfigSubject object from its proto representation.
func ProtoToPrivatecaCertificateAuthorityConfigSubjectConfigSubject(p *privatecapb.PrivatecaCertificateAuthorityConfigSubjectConfigSubject) *privateca.CertificateAuthorityConfigSubjectConfigSubject {
	if p == nil {
		return nil
	}
	obj := &privateca.CertificateAuthorityConfigSubjectConfigSubject{
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
func ProtoToPrivatecaCertificateAuthorityConfigSubjectConfigSubjectAltName(p *privatecapb.PrivatecaCertificateAuthorityConfigSubjectConfigSubjectAltName) *privateca.CertificateAuthorityConfigSubjectConfigSubjectAltName {
	if p == nil {
		return nil
	}
	obj := &privateca.CertificateAuthorityConfigSubjectConfigSubjectAltName{}
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
		obj.CustomSans = append(obj.CustomSans, *ProtoToPrivatecaCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans(r))
	}
	return obj
}

// ProtoToCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans converts a CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans object from its proto representation.
func ProtoToPrivatecaCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans(p *privatecapb.PrivatecaCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans) *privateca.CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans {
	if p == nil {
		return nil
	}
	obj := &privateca.CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans{
		ObjectId: ProtoToPrivatecaCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId(p.GetObjectId()),
		Critical: dcl.Bool(p.GetCritical()),
		Value:    dcl.StringOrNil(p.GetValue()),
	}
	return obj
}

// ProtoToCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId converts a CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId object from its proto representation.
func ProtoToPrivatecaCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId(p *privatecapb.PrivatecaCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId) *privateca.CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId {
	if p == nil {
		return nil
	}
	obj := &privateca.CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCertificateAuthorityConfigX509Config converts a CertificateAuthorityConfigX509Config object from its proto representation.
func ProtoToPrivatecaCertificateAuthorityConfigX509Config(p *privatecapb.PrivatecaCertificateAuthorityConfigX509Config) *privateca.CertificateAuthorityConfigX509Config {
	if p == nil {
		return nil
	}
	obj := &privateca.CertificateAuthorityConfigX509Config{
		KeyUsage:  ProtoToPrivatecaCertificateAuthorityConfigX509ConfigKeyUsage(p.GetKeyUsage()),
		CaOptions: ProtoToPrivatecaCertificateAuthorityConfigX509ConfigCaOptions(p.GetCaOptions()),
	}
	for _, r := range p.GetPolicyIds() {
		obj.PolicyIds = append(obj.PolicyIds, *ProtoToPrivatecaCertificateAuthorityConfigX509ConfigPolicyIds(r))
	}
	for _, r := range p.GetAiaOcspServers() {
		obj.AiaOcspServers = append(obj.AiaOcspServers, r)
	}
	for _, r := range p.GetAdditionalExtensions() {
		obj.AdditionalExtensions = append(obj.AdditionalExtensions, *ProtoToPrivatecaCertificateAuthorityConfigX509ConfigAdditionalExtensions(r))
	}
	return obj
}

// ProtoToCertificateAuthorityConfigX509ConfigKeyUsage converts a CertificateAuthorityConfigX509ConfigKeyUsage object from its proto representation.
func ProtoToPrivatecaCertificateAuthorityConfigX509ConfigKeyUsage(p *privatecapb.PrivatecaCertificateAuthorityConfigX509ConfigKeyUsage) *privateca.CertificateAuthorityConfigX509ConfigKeyUsage {
	if p == nil {
		return nil
	}
	obj := &privateca.CertificateAuthorityConfigX509ConfigKeyUsage{
		BaseKeyUsage:     ProtoToPrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsage(p.GetBaseKeyUsage()),
		ExtendedKeyUsage: ProtoToPrivatecaCertificateAuthorityConfigX509ConfigKeyUsageExtendedKeyUsage(p.GetExtendedKeyUsage()),
	}
	for _, r := range p.GetUnknownExtendedKeyUsages() {
		obj.UnknownExtendedKeyUsages = append(obj.UnknownExtendedKeyUsages, *ProtoToPrivatecaCertificateAuthorityConfigX509ConfigKeyUsageUnknownExtendedKeyUsages(r))
	}
	return obj
}

// ProtoToCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsage converts a CertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsage object from its proto representation.
func ProtoToPrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsage(p *privatecapb.PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsage) *privateca.CertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsage {
	if p == nil {
		return nil
	}
	obj := &privateca.CertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsage{
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
func ProtoToPrivatecaCertificateAuthorityConfigX509ConfigKeyUsageExtendedKeyUsage(p *privatecapb.PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageExtendedKeyUsage) *privateca.CertificateAuthorityConfigX509ConfigKeyUsageExtendedKeyUsage {
	if p == nil {
		return nil
	}
	obj := &privateca.CertificateAuthorityConfigX509ConfigKeyUsageExtendedKeyUsage{
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
func ProtoToPrivatecaCertificateAuthorityConfigX509ConfigKeyUsageUnknownExtendedKeyUsages(p *privatecapb.PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageUnknownExtendedKeyUsages) *privateca.CertificateAuthorityConfigX509ConfigKeyUsageUnknownExtendedKeyUsages {
	if p == nil {
		return nil
	}
	obj := &privateca.CertificateAuthorityConfigX509ConfigKeyUsageUnknownExtendedKeyUsages{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCertificateAuthorityConfigX509ConfigCaOptions converts a CertificateAuthorityConfigX509ConfigCaOptions object from its proto representation.
func ProtoToPrivatecaCertificateAuthorityConfigX509ConfigCaOptions(p *privatecapb.PrivatecaCertificateAuthorityConfigX509ConfigCaOptions) *privateca.CertificateAuthorityConfigX509ConfigCaOptions {
	if p == nil {
		return nil
	}
	obj := &privateca.CertificateAuthorityConfigX509ConfigCaOptions{
		IsCa:                    dcl.Bool(p.GetIsCa()),
		MaxIssuerPathLength:     dcl.Int64OrNil(p.GetMaxIssuerPathLength()),
		ZeroMaxIssuerPathLength: dcl.Bool(p.GetZeroMaxIssuerPathLength()),
	}
	return obj
}

// ProtoToCertificateAuthorityConfigX509ConfigPolicyIds converts a CertificateAuthorityConfigX509ConfigPolicyIds object from its proto representation.
func ProtoToPrivatecaCertificateAuthorityConfigX509ConfigPolicyIds(p *privatecapb.PrivatecaCertificateAuthorityConfigX509ConfigPolicyIds) *privateca.CertificateAuthorityConfigX509ConfigPolicyIds {
	if p == nil {
		return nil
	}
	obj := &privateca.CertificateAuthorityConfigX509ConfigPolicyIds{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCertificateAuthorityConfigX509ConfigAdditionalExtensions converts a CertificateAuthorityConfigX509ConfigAdditionalExtensions object from its proto representation.
func ProtoToPrivatecaCertificateAuthorityConfigX509ConfigAdditionalExtensions(p *privatecapb.PrivatecaCertificateAuthorityConfigX509ConfigAdditionalExtensions) *privateca.CertificateAuthorityConfigX509ConfigAdditionalExtensions {
	if p == nil {
		return nil
	}
	obj := &privateca.CertificateAuthorityConfigX509ConfigAdditionalExtensions{
		ObjectId: ProtoToPrivatecaCertificateAuthorityConfigX509ConfigAdditionalExtensionsObjectId(p.GetObjectId()),
		Critical: dcl.Bool(p.GetCritical()),
		Value:    dcl.StringOrNil(p.GetValue()),
	}
	return obj
}

// ProtoToCertificateAuthorityConfigX509ConfigAdditionalExtensionsObjectId converts a CertificateAuthorityConfigX509ConfigAdditionalExtensionsObjectId object from its proto representation.
func ProtoToPrivatecaCertificateAuthorityConfigX509ConfigAdditionalExtensionsObjectId(p *privatecapb.PrivatecaCertificateAuthorityConfigX509ConfigAdditionalExtensionsObjectId) *privateca.CertificateAuthorityConfigX509ConfigAdditionalExtensionsObjectId {
	if p == nil {
		return nil
	}
	obj := &privateca.CertificateAuthorityConfigX509ConfigAdditionalExtensionsObjectId{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCertificateAuthorityConfigPublicKey converts a CertificateAuthorityConfigPublicKey object from its proto representation.
func ProtoToPrivatecaCertificateAuthorityConfigPublicKey(p *privatecapb.PrivatecaCertificateAuthorityConfigPublicKey) *privateca.CertificateAuthorityConfigPublicKey {
	if p == nil {
		return nil
	}
	obj := &privateca.CertificateAuthorityConfigPublicKey{
		Key:    dcl.StringOrNil(p.GetKey()),
		Format: ProtoToPrivatecaCertificateAuthorityConfigPublicKeyFormatEnum(p.GetFormat()),
	}
	return obj
}

// ProtoToCertificateAuthorityKeySpec converts a CertificateAuthorityKeySpec object from its proto representation.
func ProtoToPrivatecaCertificateAuthorityKeySpec(p *privatecapb.PrivatecaCertificateAuthorityKeySpec) *privateca.CertificateAuthorityKeySpec {
	if p == nil {
		return nil
	}
	obj := &privateca.CertificateAuthorityKeySpec{
		CloudKmsKeyVersion: dcl.StringOrNil(p.GetCloudKmsKeyVersion()),
		Algorithm:          ProtoToPrivatecaCertificateAuthorityKeySpecAlgorithmEnum(p.GetAlgorithm()),
	}
	return obj
}

// ProtoToCertificateAuthoritySubordinateConfig converts a CertificateAuthoritySubordinateConfig object from its proto representation.
func ProtoToPrivatecaCertificateAuthoritySubordinateConfig(p *privatecapb.PrivatecaCertificateAuthoritySubordinateConfig) *privateca.CertificateAuthoritySubordinateConfig {
	if p == nil {
		return nil
	}
	obj := &privateca.CertificateAuthoritySubordinateConfig{
		CertificateAuthority: dcl.StringOrNil(p.GetCertificateAuthority()),
		PemIssuerChain:       ProtoToPrivatecaCertificateAuthoritySubordinateConfigPemIssuerChain(p.GetPemIssuerChain()),
	}
	return obj
}

// ProtoToCertificateAuthoritySubordinateConfigPemIssuerChain converts a CertificateAuthoritySubordinateConfigPemIssuerChain object from its proto representation.
func ProtoToPrivatecaCertificateAuthoritySubordinateConfigPemIssuerChain(p *privatecapb.PrivatecaCertificateAuthoritySubordinateConfigPemIssuerChain) *privateca.CertificateAuthoritySubordinateConfigPemIssuerChain {
	if p == nil {
		return nil
	}
	obj := &privateca.CertificateAuthoritySubordinateConfigPemIssuerChain{}
	for _, r := range p.GetPemCertificates() {
		obj.PemCertificates = append(obj.PemCertificates, r)
	}
	return obj
}

// ProtoToCertificateAuthorityCaCertificateDescriptions converts a CertificateAuthorityCaCertificateDescriptions object from its proto representation.
func ProtoToPrivatecaCertificateAuthorityCaCertificateDescriptions(p *privatecapb.PrivatecaCertificateAuthorityCaCertificateDescriptions) *privateca.CertificateAuthorityCaCertificateDescriptions {
	if p == nil {
		return nil
	}
	obj := &privateca.CertificateAuthorityCaCertificateDescriptions{
		SubjectDescription: ProtoToPrivatecaCertificateAuthorityCaCertificateDescriptionsSubjectDescription(p.GetSubjectDescription()),
		X509Description:    ProtoToPrivatecaCertificateAuthorityCaCertificateDescriptionsX509Description(p.GetX509Description()),
		PublicKey:          ProtoToPrivatecaCertificateAuthorityCaCertificateDescriptionsPublicKey(p.GetPublicKey()),
		SubjectKeyId:       ProtoToPrivatecaCertificateAuthorityCaCertificateDescriptionsSubjectKeyId(p.GetSubjectKeyId()),
		AuthorityKeyId:     ProtoToPrivatecaCertificateAuthorityCaCertificateDescriptionsAuthorityKeyId(p.GetAuthorityKeyId()),
		CertFingerprint:    ProtoToPrivatecaCertificateAuthorityCaCertificateDescriptionsCertFingerprint(p.GetCertFingerprint()),
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
func ProtoToPrivatecaCertificateAuthorityCaCertificateDescriptionsSubjectDescription(p *privatecapb.PrivatecaCertificateAuthorityCaCertificateDescriptionsSubjectDescription) *privateca.CertificateAuthorityCaCertificateDescriptionsSubjectDescription {
	if p == nil {
		return nil
	}
	obj := &privateca.CertificateAuthorityCaCertificateDescriptionsSubjectDescription{
		Subject:         ProtoToPrivatecaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject(p.GetSubject()),
		SubjectAltName:  ProtoToPrivatecaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName(p.GetSubjectAltName()),
		HexSerialNumber: dcl.StringOrNil(p.GetHexSerialNumber()),
		Lifetime:        dcl.StringOrNil(p.GetLifetime()),
		NotBeforeTime:   dcl.StringOrNil(p.GetNotBeforeTime()),
		NotAfterTime:    dcl.StringOrNil(p.GetNotAfterTime()),
	}
	return obj
}

// ProtoToCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject converts a CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject object from its proto representation.
func ProtoToPrivatecaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject(p *privatecapb.PrivatecaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject) *privateca.CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject {
	if p == nil {
		return nil
	}
	obj := &privateca.CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject{
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
func ProtoToPrivatecaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName(p *privatecapb.PrivatecaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName) *privateca.CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName {
	if p == nil {
		return nil
	}
	obj := &privateca.CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName{}
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
		obj.CustomSans = append(obj.CustomSans, *ProtoToPrivatecaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans(r))
	}
	return obj
}

// ProtoToCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans converts a CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans object from its proto representation.
func ProtoToPrivatecaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans(p *privatecapb.PrivatecaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans) *privateca.CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans {
	if p == nil {
		return nil
	}
	obj := &privateca.CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans{
		ObjectId: ProtoToPrivatecaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId(p.GetObjectId()),
		Critical: dcl.Bool(p.GetCritical()),
		Value:    dcl.StringOrNil(p.GetValue()),
	}
	return obj
}

// ProtoToCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId converts a CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId object from its proto representation.
func ProtoToPrivatecaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId(p *privatecapb.PrivatecaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId) *privateca.CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId {
	if p == nil {
		return nil
	}
	obj := &privateca.CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCertificateAuthorityCaCertificateDescriptionsX509Description converts a CertificateAuthorityCaCertificateDescriptionsX509Description object from its proto representation.
func ProtoToPrivatecaCertificateAuthorityCaCertificateDescriptionsX509Description(p *privatecapb.PrivatecaCertificateAuthorityCaCertificateDescriptionsX509Description) *privateca.CertificateAuthorityCaCertificateDescriptionsX509Description {
	if p == nil {
		return nil
	}
	obj := &privateca.CertificateAuthorityCaCertificateDescriptionsX509Description{
		KeyUsage:  ProtoToPrivatecaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsage(p.GetKeyUsage()),
		CaOptions: ProtoToPrivatecaCertificateAuthorityCaCertificateDescriptionsX509DescriptionCaOptions(p.GetCaOptions()),
	}
	for _, r := range p.GetPolicyIds() {
		obj.PolicyIds = append(obj.PolicyIds, *ProtoToPrivatecaCertificateAuthorityCaCertificateDescriptionsX509DescriptionPolicyIds(r))
	}
	for _, r := range p.GetAiaOcspServers() {
		obj.AiaOcspServers = append(obj.AiaOcspServers, r)
	}
	for _, r := range p.GetAdditionalExtensions() {
		obj.AdditionalExtensions = append(obj.AdditionalExtensions, *ProtoToPrivatecaCertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensions(r))
	}
	return obj
}

// ProtoToCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsage converts a CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsage object from its proto representation.
func ProtoToPrivatecaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsage(p *privatecapb.PrivatecaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsage) *privateca.CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsage {
	if p == nil {
		return nil
	}
	obj := &privateca.CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsage{
		BaseKeyUsage:     ProtoToPrivatecaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageBaseKeyUsage(p.GetBaseKeyUsage()),
		ExtendedKeyUsage: ProtoToPrivatecaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageExtendedKeyUsage(p.GetExtendedKeyUsage()),
	}
	for _, r := range p.GetUnknownExtendedKeyUsages() {
		obj.UnknownExtendedKeyUsages = append(obj.UnknownExtendedKeyUsages, *ProtoToPrivatecaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageUnknownExtendedKeyUsages(r))
	}
	return obj
}

// ProtoToCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageBaseKeyUsage converts a CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageBaseKeyUsage object from its proto representation.
func ProtoToPrivatecaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageBaseKeyUsage(p *privatecapb.PrivatecaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageBaseKeyUsage) *privateca.CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageBaseKeyUsage {
	if p == nil {
		return nil
	}
	obj := &privateca.CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageBaseKeyUsage{
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
func ProtoToPrivatecaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageExtendedKeyUsage(p *privatecapb.PrivatecaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageExtendedKeyUsage) *privateca.CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageExtendedKeyUsage {
	if p == nil {
		return nil
	}
	obj := &privateca.CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageExtendedKeyUsage{
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
func ProtoToPrivatecaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageUnknownExtendedKeyUsages(p *privatecapb.PrivatecaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageUnknownExtendedKeyUsages) *privateca.CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageUnknownExtendedKeyUsages {
	if p == nil {
		return nil
	}
	obj := &privateca.CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageUnknownExtendedKeyUsages{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCertificateAuthorityCaCertificateDescriptionsX509DescriptionCaOptions converts a CertificateAuthorityCaCertificateDescriptionsX509DescriptionCaOptions object from its proto representation.
func ProtoToPrivatecaCertificateAuthorityCaCertificateDescriptionsX509DescriptionCaOptions(p *privatecapb.PrivatecaCertificateAuthorityCaCertificateDescriptionsX509DescriptionCaOptions) *privateca.CertificateAuthorityCaCertificateDescriptionsX509DescriptionCaOptions {
	if p == nil {
		return nil
	}
	obj := &privateca.CertificateAuthorityCaCertificateDescriptionsX509DescriptionCaOptions{
		IsCa:                dcl.Bool(p.GetIsCa()),
		MaxIssuerPathLength: dcl.Int64OrNil(p.GetMaxIssuerPathLength()),
	}
	return obj
}

// ProtoToCertificateAuthorityCaCertificateDescriptionsX509DescriptionPolicyIds converts a CertificateAuthorityCaCertificateDescriptionsX509DescriptionPolicyIds object from its proto representation.
func ProtoToPrivatecaCertificateAuthorityCaCertificateDescriptionsX509DescriptionPolicyIds(p *privatecapb.PrivatecaCertificateAuthorityCaCertificateDescriptionsX509DescriptionPolicyIds) *privateca.CertificateAuthorityCaCertificateDescriptionsX509DescriptionPolicyIds {
	if p == nil {
		return nil
	}
	obj := &privateca.CertificateAuthorityCaCertificateDescriptionsX509DescriptionPolicyIds{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensions converts a CertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensions object from its proto representation.
func ProtoToPrivatecaCertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensions(p *privatecapb.PrivatecaCertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensions) *privateca.CertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensions {
	if p == nil {
		return nil
	}
	obj := &privateca.CertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensions{
		ObjectId: ProtoToPrivatecaCertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensionsObjectId(p.GetObjectId()),
		Critical: dcl.Bool(p.GetCritical()),
		Value:    dcl.StringOrNil(p.GetValue()),
	}
	return obj
}

// ProtoToCertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensionsObjectId converts a CertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensionsObjectId object from its proto representation.
func ProtoToPrivatecaCertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensionsObjectId(p *privatecapb.PrivatecaCertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensionsObjectId) *privateca.CertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensionsObjectId {
	if p == nil {
		return nil
	}
	obj := &privateca.CertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensionsObjectId{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCertificateAuthorityCaCertificateDescriptionsPublicKey converts a CertificateAuthorityCaCertificateDescriptionsPublicKey object from its proto representation.
func ProtoToPrivatecaCertificateAuthorityCaCertificateDescriptionsPublicKey(p *privatecapb.PrivatecaCertificateAuthorityCaCertificateDescriptionsPublicKey) *privateca.CertificateAuthorityCaCertificateDescriptionsPublicKey {
	if p == nil {
		return nil
	}
	obj := &privateca.CertificateAuthorityCaCertificateDescriptionsPublicKey{
		Key:    dcl.StringOrNil(p.GetKey()),
		Format: ProtoToPrivatecaCertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum(p.GetFormat()),
	}
	return obj
}

// ProtoToCertificateAuthorityCaCertificateDescriptionsSubjectKeyId converts a CertificateAuthorityCaCertificateDescriptionsSubjectKeyId object from its proto representation.
func ProtoToPrivatecaCertificateAuthorityCaCertificateDescriptionsSubjectKeyId(p *privatecapb.PrivatecaCertificateAuthorityCaCertificateDescriptionsSubjectKeyId) *privateca.CertificateAuthorityCaCertificateDescriptionsSubjectKeyId {
	if p == nil {
		return nil
	}
	obj := &privateca.CertificateAuthorityCaCertificateDescriptionsSubjectKeyId{
		KeyId: dcl.StringOrNil(p.GetKeyId()),
	}
	return obj
}

// ProtoToCertificateAuthorityCaCertificateDescriptionsAuthorityKeyId converts a CertificateAuthorityCaCertificateDescriptionsAuthorityKeyId object from its proto representation.
func ProtoToPrivatecaCertificateAuthorityCaCertificateDescriptionsAuthorityKeyId(p *privatecapb.PrivatecaCertificateAuthorityCaCertificateDescriptionsAuthorityKeyId) *privateca.CertificateAuthorityCaCertificateDescriptionsAuthorityKeyId {
	if p == nil {
		return nil
	}
	obj := &privateca.CertificateAuthorityCaCertificateDescriptionsAuthorityKeyId{
		KeyId: dcl.StringOrNil(p.GetKeyId()),
	}
	return obj
}

// ProtoToCertificateAuthorityCaCertificateDescriptionsCertFingerprint converts a CertificateAuthorityCaCertificateDescriptionsCertFingerprint object from its proto representation.
func ProtoToPrivatecaCertificateAuthorityCaCertificateDescriptionsCertFingerprint(p *privatecapb.PrivatecaCertificateAuthorityCaCertificateDescriptionsCertFingerprint) *privateca.CertificateAuthorityCaCertificateDescriptionsCertFingerprint {
	if p == nil {
		return nil
	}
	obj := &privateca.CertificateAuthorityCaCertificateDescriptionsCertFingerprint{
		Sha256Hash: dcl.StringOrNil(p.GetSha256Hash()),
	}
	return obj
}

// ProtoToCertificateAuthorityAccessUrls converts a CertificateAuthorityAccessUrls object from its proto representation.
func ProtoToPrivatecaCertificateAuthorityAccessUrls(p *privatecapb.PrivatecaCertificateAuthorityAccessUrls) *privateca.CertificateAuthorityAccessUrls {
	if p == nil {
		return nil
	}
	obj := &privateca.CertificateAuthorityAccessUrls{
		CaCertificateAccessUrl: dcl.StringOrNil(p.GetCaCertificateAccessUrl()),
	}
	for _, r := range p.GetCrlAccessUrls() {
		obj.CrlAccessUrls = append(obj.CrlAccessUrls, r)
	}
	return obj
}

// ProtoToCertificateAuthority converts a CertificateAuthority resource from its proto representation.
func ProtoToCertificateAuthority(p *privatecapb.PrivatecaCertificateAuthority) *privateca.CertificateAuthority {
	obj := &privateca.CertificateAuthority{
		Name:              dcl.StringOrNil(p.GetName()),
		Type:              ProtoToPrivatecaCertificateAuthorityTypeEnum(p.GetType()),
		Config:            ProtoToPrivatecaCertificateAuthorityConfig(p.GetConfig()),
		Lifetime:          dcl.StringOrNil(p.GetLifetime()),
		KeySpec:           ProtoToPrivatecaCertificateAuthorityKeySpec(p.GetKeySpec()),
		SubordinateConfig: ProtoToPrivatecaCertificateAuthoritySubordinateConfig(p.GetSubordinateConfig()),
		Tier:              ProtoToPrivatecaCertificateAuthorityTierEnum(p.GetTier()),
		State:             ProtoToPrivatecaCertificateAuthorityStateEnum(p.GetState()),
		GcsBucket:         dcl.StringOrNil(p.GetGcsBucket()),
		AccessUrls:        ProtoToPrivatecaCertificateAuthorityAccessUrls(p.GetAccessUrls()),
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
		obj.CaCertificateDescriptions = append(obj.CaCertificateDescriptions, *ProtoToPrivatecaCertificateAuthorityCaCertificateDescriptions(r))
	}
	return obj
}

// CertificateAuthorityTypeEnumToProto converts a CertificateAuthorityTypeEnum enum to its proto representation.
func PrivatecaCertificateAuthorityTypeEnumToProto(e *privateca.CertificateAuthorityTypeEnum) privatecapb.PrivatecaCertificateAuthorityTypeEnum {
	if e == nil {
		return privatecapb.PrivatecaCertificateAuthorityTypeEnum(0)
	}
	if v, ok := privatecapb.PrivatecaCertificateAuthorityTypeEnum_value["CertificateAuthorityTypeEnum"+string(*e)]; ok {
		return privatecapb.PrivatecaCertificateAuthorityTypeEnum(v)
	}
	return privatecapb.PrivatecaCertificateAuthorityTypeEnum(0)
}

// CertificateAuthorityConfigPublicKeyFormatEnumToProto converts a CertificateAuthorityConfigPublicKeyFormatEnum enum to its proto representation.
func PrivatecaCertificateAuthorityConfigPublicKeyFormatEnumToProto(e *privateca.CertificateAuthorityConfigPublicKeyFormatEnum) privatecapb.PrivatecaCertificateAuthorityConfigPublicKeyFormatEnum {
	if e == nil {
		return privatecapb.PrivatecaCertificateAuthorityConfigPublicKeyFormatEnum(0)
	}
	if v, ok := privatecapb.PrivatecaCertificateAuthorityConfigPublicKeyFormatEnum_value["CertificateAuthorityConfigPublicKeyFormatEnum"+string(*e)]; ok {
		return privatecapb.PrivatecaCertificateAuthorityConfigPublicKeyFormatEnum(v)
	}
	return privatecapb.PrivatecaCertificateAuthorityConfigPublicKeyFormatEnum(0)
}

// CertificateAuthorityKeySpecAlgorithmEnumToProto converts a CertificateAuthorityKeySpecAlgorithmEnum enum to its proto representation.
func PrivatecaCertificateAuthorityKeySpecAlgorithmEnumToProto(e *privateca.CertificateAuthorityKeySpecAlgorithmEnum) privatecapb.PrivatecaCertificateAuthorityKeySpecAlgorithmEnum {
	if e == nil {
		return privatecapb.PrivatecaCertificateAuthorityKeySpecAlgorithmEnum(0)
	}
	if v, ok := privatecapb.PrivatecaCertificateAuthorityKeySpecAlgorithmEnum_value["CertificateAuthorityKeySpecAlgorithmEnum"+string(*e)]; ok {
		return privatecapb.PrivatecaCertificateAuthorityKeySpecAlgorithmEnum(v)
	}
	return privatecapb.PrivatecaCertificateAuthorityKeySpecAlgorithmEnum(0)
}

// CertificateAuthorityTierEnumToProto converts a CertificateAuthorityTierEnum enum to its proto representation.
func PrivatecaCertificateAuthorityTierEnumToProto(e *privateca.CertificateAuthorityTierEnum) privatecapb.PrivatecaCertificateAuthorityTierEnum {
	if e == nil {
		return privatecapb.PrivatecaCertificateAuthorityTierEnum(0)
	}
	if v, ok := privatecapb.PrivatecaCertificateAuthorityTierEnum_value["CertificateAuthorityTierEnum"+string(*e)]; ok {
		return privatecapb.PrivatecaCertificateAuthorityTierEnum(v)
	}
	return privatecapb.PrivatecaCertificateAuthorityTierEnum(0)
}

// CertificateAuthorityStateEnumToProto converts a CertificateAuthorityStateEnum enum to its proto representation.
func PrivatecaCertificateAuthorityStateEnumToProto(e *privateca.CertificateAuthorityStateEnum) privatecapb.PrivatecaCertificateAuthorityStateEnum {
	if e == nil {
		return privatecapb.PrivatecaCertificateAuthorityStateEnum(0)
	}
	if v, ok := privatecapb.PrivatecaCertificateAuthorityStateEnum_value["CertificateAuthorityStateEnum"+string(*e)]; ok {
		return privatecapb.PrivatecaCertificateAuthorityStateEnum(v)
	}
	return privatecapb.PrivatecaCertificateAuthorityStateEnum(0)
}

// CertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnumToProto converts a CertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum enum to its proto representation.
func PrivatecaCertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnumToProto(e *privateca.CertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum) privatecapb.PrivatecaCertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum {
	if e == nil {
		return privatecapb.PrivatecaCertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum(0)
	}
	if v, ok := privatecapb.PrivatecaCertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum_value["CertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum"+string(*e)]; ok {
		return privatecapb.PrivatecaCertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum(v)
	}
	return privatecapb.PrivatecaCertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum(0)
}

// CertificateAuthorityConfigToProto converts a CertificateAuthorityConfig object to its proto representation.
func PrivatecaCertificateAuthorityConfigToProto(o *privateca.CertificateAuthorityConfig) *privatecapb.PrivatecaCertificateAuthorityConfig {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCertificateAuthorityConfig{}
	p.SetSubjectConfig(PrivatecaCertificateAuthorityConfigSubjectConfigToProto(o.SubjectConfig))
	p.SetX509Config(PrivatecaCertificateAuthorityConfigX509ConfigToProto(o.X509Config))
	p.SetPublicKey(PrivatecaCertificateAuthorityConfigPublicKeyToProto(o.PublicKey))
	return p
}

// CertificateAuthorityConfigSubjectConfigToProto converts a CertificateAuthorityConfigSubjectConfig object to its proto representation.
func PrivatecaCertificateAuthorityConfigSubjectConfigToProto(o *privateca.CertificateAuthorityConfigSubjectConfig) *privatecapb.PrivatecaCertificateAuthorityConfigSubjectConfig {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCertificateAuthorityConfigSubjectConfig{}
	p.SetSubject(PrivatecaCertificateAuthorityConfigSubjectConfigSubjectToProto(o.Subject))
	p.SetSubjectAltName(PrivatecaCertificateAuthorityConfigSubjectConfigSubjectAltNameToProto(o.SubjectAltName))
	return p
}

// CertificateAuthorityConfigSubjectConfigSubjectToProto converts a CertificateAuthorityConfigSubjectConfigSubject object to its proto representation.
func PrivatecaCertificateAuthorityConfigSubjectConfigSubjectToProto(o *privateca.CertificateAuthorityConfigSubjectConfigSubject) *privatecapb.PrivatecaCertificateAuthorityConfigSubjectConfigSubject {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCertificateAuthorityConfigSubjectConfigSubject{}
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
func PrivatecaCertificateAuthorityConfigSubjectConfigSubjectAltNameToProto(o *privateca.CertificateAuthorityConfigSubjectConfigSubjectAltName) *privatecapb.PrivatecaCertificateAuthorityConfigSubjectConfigSubjectAltName {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCertificateAuthorityConfigSubjectConfigSubjectAltName{}
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
	sCustomSans := make([]*privatecapb.PrivatecaCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans, len(o.CustomSans))
	for i, r := range o.CustomSans {
		sCustomSans[i] = PrivatecaCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansToProto(&r)
	}
	p.SetCustomSans(sCustomSans)
	return p
}

// CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansToProto converts a CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans object to its proto representation.
func PrivatecaCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansToProto(o *privateca.CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans) *privatecapb.PrivatecaCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans{}
	p.SetObjectId(PrivatecaCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectIdToProto(o.ObjectId))
	p.SetCritical(dcl.ValueOrEmptyBool(o.Critical))
	p.SetValue(dcl.ValueOrEmptyString(o.Value))
	return p
}

// CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectIdToProto converts a CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId object to its proto representation.
func PrivatecaCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectIdToProto(o *privateca.CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId) *privatecapb.PrivatecaCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId{}
	sObjectIdPath := make([]int64, len(o.ObjectIdPath))
	for i, r := range o.ObjectIdPath {
		sObjectIdPath[i] = r
	}
	p.SetObjectIdPath(sObjectIdPath)
	return p
}

// CertificateAuthorityConfigX509ConfigToProto converts a CertificateAuthorityConfigX509Config object to its proto representation.
func PrivatecaCertificateAuthorityConfigX509ConfigToProto(o *privateca.CertificateAuthorityConfigX509Config) *privatecapb.PrivatecaCertificateAuthorityConfigX509Config {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCertificateAuthorityConfigX509Config{}
	p.SetKeyUsage(PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageToProto(o.KeyUsage))
	p.SetCaOptions(PrivatecaCertificateAuthorityConfigX509ConfigCaOptionsToProto(o.CaOptions))
	sPolicyIds := make([]*privatecapb.PrivatecaCertificateAuthorityConfigX509ConfigPolicyIds, len(o.PolicyIds))
	for i, r := range o.PolicyIds {
		sPolicyIds[i] = PrivatecaCertificateAuthorityConfigX509ConfigPolicyIdsToProto(&r)
	}
	p.SetPolicyIds(sPolicyIds)
	sAiaOcspServers := make([]string, len(o.AiaOcspServers))
	for i, r := range o.AiaOcspServers {
		sAiaOcspServers[i] = r
	}
	p.SetAiaOcspServers(sAiaOcspServers)
	sAdditionalExtensions := make([]*privatecapb.PrivatecaCertificateAuthorityConfigX509ConfigAdditionalExtensions, len(o.AdditionalExtensions))
	for i, r := range o.AdditionalExtensions {
		sAdditionalExtensions[i] = PrivatecaCertificateAuthorityConfigX509ConfigAdditionalExtensionsToProto(&r)
	}
	p.SetAdditionalExtensions(sAdditionalExtensions)
	return p
}

// CertificateAuthorityConfigX509ConfigKeyUsageToProto converts a CertificateAuthorityConfigX509ConfigKeyUsage object to its proto representation.
func PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageToProto(o *privateca.CertificateAuthorityConfigX509ConfigKeyUsage) *privatecapb.PrivatecaCertificateAuthorityConfigX509ConfigKeyUsage {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCertificateAuthorityConfigX509ConfigKeyUsage{}
	p.SetBaseKeyUsage(PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageToProto(o.BaseKeyUsage))
	p.SetExtendedKeyUsage(PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageExtendedKeyUsageToProto(o.ExtendedKeyUsage))
	sUnknownExtendedKeyUsages := make([]*privatecapb.PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageUnknownExtendedKeyUsages, len(o.UnknownExtendedKeyUsages))
	for i, r := range o.UnknownExtendedKeyUsages {
		sUnknownExtendedKeyUsages[i] = PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageUnknownExtendedKeyUsagesToProto(&r)
	}
	p.SetUnknownExtendedKeyUsages(sUnknownExtendedKeyUsages)
	return p
}

// CertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageToProto converts a CertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsage object to its proto representation.
func PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageToProto(o *privateca.CertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsage) *privatecapb.PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsage {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsage{}
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
func PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageExtendedKeyUsageToProto(o *privateca.CertificateAuthorityConfigX509ConfigKeyUsageExtendedKeyUsage) *privatecapb.PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageExtendedKeyUsage {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageExtendedKeyUsage{}
	p.SetServerAuth(dcl.ValueOrEmptyBool(o.ServerAuth))
	p.SetClientAuth(dcl.ValueOrEmptyBool(o.ClientAuth))
	p.SetCodeSigning(dcl.ValueOrEmptyBool(o.CodeSigning))
	p.SetEmailProtection(dcl.ValueOrEmptyBool(o.EmailProtection))
	p.SetTimeStamping(dcl.ValueOrEmptyBool(o.TimeStamping))
	p.SetOcspSigning(dcl.ValueOrEmptyBool(o.OcspSigning))
	return p
}

// CertificateAuthorityConfigX509ConfigKeyUsageUnknownExtendedKeyUsagesToProto converts a CertificateAuthorityConfigX509ConfigKeyUsageUnknownExtendedKeyUsages object to its proto representation.
func PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageUnknownExtendedKeyUsagesToProto(o *privateca.CertificateAuthorityConfigX509ConfigKeyUsageUnknownExtendedKeyUsages) *privatecapb.PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageUnknownExtendedKeyUsages {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageUnknownExtendedKeyUsages{}
	sObjectIdPath := make([]int64, len(o.ObjectIdPath))
	for i, r := range o.ObjectIdPath {
		sObjectIdPath[i] = r
	}
	p.SetObjectIdPath(sObjectIdPath)
	return p
}

// CertificateAuthorityConfigX509ConfigCaOptionsToProto converts a CertificateAuthorityConfigX509ConfigCaOptions object to its proto representation.
func PrivatecaCertificateAuthorityConfigX509ConfigCaOptionsToProto(o *privateca.CertificateAuthorityConfigX509ConfigCaOptions) *privatecapb.PrivatecaCertificateAuthorityConfigX509ConfigCaOptions {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCertificateAuthorityConfigX509ConfigCaOptions{}
	p.SetIsCa(dcl.ValueOrEmptyBool(o.IsCa))
	p.SetMaxIssuerPathLength(dcl.ValueOrEmptyInt64(o.MaxIssuerPathLength))
	p.SetZeroMaxIssuerPathLength(dcl.ValueOrEmptyBool(o.ZeroMaxIssuerPathLength))
	return p
}

// CertificateAuthorityConfigX509ConfigPolicyIdsToProto converts a CertificateAuthorityConfigX509ConfigPolicyIds object to its proto representation.
func PrivatecaCertificateAuthorityConfigX509ConfigPolicyIdsToProto(o *privateca.CertificateAuthorityConfigX509ConfigPolicyIds) *privatecapb.PrivatecaCertificateAuthorityConfigX509ConfigPolicyIds {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCertificateAuthorityConfigX509ConfigPolicyIds{}
	sObjectIdPath := make([]int64, len(o.ObjectIdPath))
	for i, r := range o.ObjectIdPath {
		sObjectIdPath[i] = r
	}
	p.SetObjectIdPath(sObjectIdPath)
	return p
}

// CertificateAuthorityConfigX509ConfigAdditionalExtensionsToProto converts a CertificateAuthorityConfigX509ConfigAdditionalExtensions object to its proto representation.
func PrivatecaCertificateAuthorityConfigX509ConfigAdditionalExtensionsToProto(o *privateca.CertificateAuthorityConfigX509ConfigAdditionalExtensions) *privatecapb.PrivatecaCertificateAuthorityConfigX509ConfigAdditionalExtensions {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCertificateAuthorityConfigX509ConfigAdditionalExtensions{}
	p.SetObjectId(PrivatecaCertificateAuthorityConfigX509ConfigAdditionalExtensionsObjectIdToProto(o.ObjectId))
	p.SetCritical(dcl.ValueOrEmptyBool(o.Critical))
	p.SetValue(dcl.ValueOrEmptyString(o.Value))
	return p
}

// CertificateAuthorityConfigX509ConfigAdditionalExtensionsObjectIdToProto converts a CertificateAuthorityConfigX509ConfigAdditionalExtensionsObjectId object to its proto representation.
func PrivatecaCertificateAuthorityConfigX509ConfigAdditionalExtensionsObjectIdToProto(o *privateca.CertificateAuthorityConfigX509ConfigAdditionalExtensionsObjectId) *privatecapb.PrivatecaCertificateAuthorityConfigX509ConfigAdditionalExtensionsObjectId {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCertificateAuthorityConfigX509ConfigAdditionalExtensionsObjectId{}
	sObjectIdPath := make([]int64, len(o.ObjectIdPath))
	for i, r := range o.ObjectIdPath {
		sObjectIdPath[i] = r
	}
	p.SetObjectIdPath(sObjectIdPath)
	return p
}

// CertificateAuthorityConfigPublicKeyToProto converts a CertificateAuthorityConfigPublicKey object to its proto representation.
func PrivatecaCertificateAuthorityConfigPublicKeyToProto(o *privateca.CertificateAuthorityConfigPublicKey) *privatecapb.PrivatecaCertificateAuthorityConfigPublicKey {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCertificateAuthorityConfigPublicKey{}
	p.SetKey(dcl.ValueOrEmptyString(o.Key))
	p.SetFormat(PrivatecaCertificateAuthorityConfigPublicKeyFormatEnumToProto(o.Format))
	return p
}

// CertificateAuthorityKeySpecToProto converts a CertificateAuthorityKeySpec object to its proto representation.
func PrivatecaCertificateAuthorityKeySpecToProto(o *privateca.CertificateAuthorityKeySpec) *privatecapb.PrivatecaCertificateAuthorityKeySpec {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCertificateAuthorityKeySpec{}
	p.SetCloudKmsKeyVersion(dcl.ValueOrEmptyString(o.CloudKmsKeyVersion))
	p.SetAlgorithm(PrivatecaCertificateAuthorityKeySpecAlgorithmEnumToProto(o.Algorithm))
	return p
}

// CertificateAuthoritySubordinateConfigToProto converts a CertificateAuthoritySubordinateConfig object to its proto representation.
func PrivatecaCertificateAuthoritySubordinateConfigToProto(o *privateca.CertificateAuthoritySubordinateConfig) *privatecapb.PrivatecaCertificateAuthoritySubordinateConfig {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCertificateAuthoritySubordinateConfig{}
	p.SetCertificateAuthority(dcl.ValueOrEmptyString(o.CertificateAuthority))
	p.SetPemIssuerChain(PrivatecaCertificateAuthoritySubordinateConfigPemIssuerChainToProto(o.PemIssuerChain))
	return p
}

// CertificateAuthoritySubordinateConfigPemIssuerChainToProto converts a CertificateAuthoritySubordinateConfigPemIssuerChain object to its proto representation.
func PrivatecaCertificateAuthoritySubordinateConfigPemIssuerChainToProto(o *privateca.CertificateAuthoritySubordinateConfigPemIssuerChain) *privatecapb.PrivatecaCertificateAuthoritySubordinateConfigPemIssuerChain {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCertificateAuthoritySubordinateConfigPemIssuerChain{}
	sPemCertificates := make([]string, len(o.PemCertificates))
	for i, r := range o.PemCertificates {
		sPemCertificates[i] = r
	}
	p.SetPemCertificates(sPemCertificates)
	return p
}

// CertificateAuthorityCaCertificateDescriptionsToProto converts a CertificateAuthorityCaCertificateDescriptions object to its proto representation.
func PrivatecaCertificateAuthorityCaCertificateDescriptionsToProto(o *privateca.CertificateAuthorityCaCertificateDescriptions) *privatecapb.PrivatecaCertificateAuthorityCaCertificateDescriptions {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCertificateAuthorityCaCertificateDescriptions{}
	p.SetSubjectDescription(PrivatecaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionToProto(o.SubjectDescription))
	p.SetX509Description(PrivatecaCertificateAuthorityCaCertificateDescriptionsX509DescriptionToProto(o.X509Description))
	p.SetPublicKey(PrivatecaCertificateAuthorityCaCertificateDescriptionsPublicKeyToProto(o.PublicKey))
	p.SetSubjectKeyId(PrivatecaCertificateAuthorityCaCertificateDescriptionsSubjectKeyIdToProto(o.SubjectKeyId))
	p.SetAuthorityKeyId(PrivatecaCertificateAuthorityCaCertificateDescriptionsAuthorityKeyIdToProto(o.AuthorityKeyId))
	p.SetCertFingerprint(PrivatecaCertificateAuthorityCaCertificateDescriptionsCertFingerprintToProto(o.CertFingerprint))
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
func PrivatecaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionToProto(o *privateca.CertificateAuthorityCaCertificateDescriptionsSubjectDescription) *privatecapb.PrivatecaCertificateAuthorityCaCertificateDescriptionsSubjectDescription {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCertificateAuthorityCaCertificateDescriptionsSubjectDescription{}
	p.SetSubject(PrivatecaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectToProto(o.Subject))
	p.SetSubjectAltName(PrivatecaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameToProto(o.SubjectAltName))
	p.SetHexSerialNumber(dcl.ValueOrEmptyString(o.HexSerialNumber))
	p.SetLifetime(dcl.ValueOrEmptyString(o.Lifetime))
	p.SetNotBeforeTime(dcl.ValueOrEmptyString(o.NotBeforeTime))
	p.SetNotAfterTime(dcl.ValueOrEmptyString(o.NotAfterTime))
	return p
}

// CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectToProto converts a CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject object to its proto representation.
func PrivatecaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectToProto(o *privateca.CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject) *privatecapb.PrivatecaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject{}
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
func PrivatecaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameToProto(o *privateca.CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName) *privatecapb.PrivatecaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName{}
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
	sCustomSans := make([]*privatecapb.PrivatecaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans, len(o.CustomSans))
	for i, r := range o.CustomSans {
		sCustomSans[i] = PrivatecaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansToProto(&r)
	}
	p.SetCustomSans(sCustomSans)
	return p
}

// CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansToProto converts a CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans object to its proto representation.
func PrivatecaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansToProto(o *privateca.CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans) *privatecapb.PrivatecaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans{}
	p.SetObjectId(PrivatecaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectIdToProto(o.ObjectId))
	p.SetCritical(dcl.ValueOrEmptyBool(o.Critical))
	p.SetValue(dcl.ValueOrEmptyString(o.Value))
	return p
}

// CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectIdToProto converts a CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId object to its proto representation.
func PrivatecaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectIdToProto(o *privateca.CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId) *privatecapb.PrivatecaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId{}
	sObjectIdPath := make([]int64, len(o.ObjectIdPath))
	for i, r := range o.ObjectIdPath {
		sObjectIdPath[i] = r
	}
	p.SetObjectIdPath(sObjectIdPath)
	return p
}

// CertificateAuthorityCaCertificateDescriptionsX509DescriptionToProto converts a CertificateAuthorityCaCertificateDescriptionsX509Description object to its proto representation.
func PrivatecaCertificateAuthorityCaCertificateDescriptionsX509DescriptionToProto(o *privateca.CertificateAuthorityCaCertificateDescriptionsX509Description) *privatecapb.PrivatecaCertificateAuthorityCaCertificateDescriptionsX509Description {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCertificateAuthorityCaCertificateDescriptionsX509Description{}
	p.SetKeyUsage(PrivatecaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageToProto(o.KeyUsage))
	p.SetCaOptions(PrivatecaCertificateAuthorityCaCertificateDescriptionsX509DescriptionCaOptionsToProto(o.CaOptions))
	sPolicyIds := make([]*privatecapb.PrivatecaCertificateAuthorityCaCertificateDescriptionsX509DescriptionPolicyIds, len(o.PolicyIds))
	for i, r := range o.PolicyIds {
		sPolicyIds[i] = PrivatecaCertificateAuthorityCaCertificateDescriptionsX509DescriptionPolicyIdsToProto(&r)
	}
	p.SetPolicyIds(sPolicyIds)
	sAiaOcspServers := make([]string, len(o.AiaOcspServers))
	for i, r := range o.AiaOcspServers {
		sAiaOcspServers[i] = r
	}
	p.SetAiaOcspServers(sAiaOcspServers)
	sAdditionalExtensions := make([]*privatecapb.PrivatecaCertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensions, len(o.AdditionalExtensions))
	for i, r := range o.AdditionalExtensions {
		sAdditionalExtensions[i] = PrivatecaCertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensionsToProto(&r)
	}
	p.SetAdditionalExtensions(sAdditionalExtensions)
	return p
}

// CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageToProto converts a CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsage object to its proto representation.
func PrivatecaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageToProto(o *privateca.CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsage) *privatecapb.PrivatecaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsage {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsage{}
	p.SetBaseKeyUsage(PrivatecaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageBaseKeyUsageToProto(o.BaseKeyUsage))
	p.SetExtendedKeyUsage(PrivatecaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageExtendedKeyUsageToProto(o.ExtendedKeyUsage))
	sUnknownExtendedKeyUsages := make([]*privatecapb.PrivatecaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageUnknownExtendedKeyUsages, len(o.UnknownExtendedKeyUsages))
	for i, r := range o.UnknownExtendedKeyUsages {
		sUnknownExtendedKeyUsages[i] = PrivatecaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageUnknownExtendedKeyUsagesToProto(&r)
	}
	p.SetUnknownExtendedKeyUsages(sUnknownExtendedKeyUsages)
	return p
}

// CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageBaseKeyUsageToProto converts a CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageBaseKeyUsage object to its proto representation.
func PrivatecaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageBaseKeyUsageToProto(o *privateca.CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageBaseKeyUsage) *privatecapb.PrivatecaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageBaseKeyUsage {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageBaseKeyUsage{}
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
func PrivatecaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageExtendedKeyUsageToProto(o *privateca.CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageExtendedKeyUsage) *privatecapb.PrivatecaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageExtendedKeyUsage {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageExtendedKeyUsage{}
	p.SetServerAuth(dcl.ValueOrEmptyBool(o.ServerAuth))
	p.SetClientAuth(dcl.ValueOrEmptyBool(o.ClientAuth))
	p.SetCodeSigning(dcl.ValueOrEmptyBool(o.CodeSigning))
	p.SetEmailProtection(dcl.ValueOrEmptyBool(o.EmailProtection))
	p.SetTimeStamping(dcl.ValueOrEmptyBool(o.TimeStamping))
	p.SetOcspSigning(dcl.ValueOrEmptyBool(o.OcspSigning))
	return p
}

// CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageUnknownExtendedKeyUsagesToProto converts a CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageUnknownExtendedKeyUsages object to its proto representation.
func PrivatecaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageUnknownExtendedKeyUsagesToProto(o *privateca.CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageUnknownExtendedKeyUsages) *privatecapb.PrivatecaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageUnknownExtendedKeyUsages {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageUnknownExtendedKeyUsages{}
	sObjectIdPath := make([]int64, len(o.ObjectIdPath))
	for i, r := range o.ObjectIdPath {
		sObjectIdPath[i] = r
	}
	p.SetObjectIdPath(sObjectIdPath)
	return p
}

// CertificateAuthorityCaCertificateDescriptionsX509DescriptionCaOptionsToProto converts a CertificateAuthorityCaCertificateDescriptionsX509DescriptionCaOptions object to its proto representation.
func PrivatecaCertificateAuthorityCaCertificateDescriptionsX509DescriptionCaOptionsToProto(o *privateca.CertificateAuthorityCaCertificateDescriptionsX509DescriptionCaOptions) *privatecapb.PrivatecaCertificateAuthorityCaCertificateDescriptionsX509DescriptionCaOptions {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCertificateAuthorityCaCertificateDescriptionsX509DescriptionCaOptions{}
	p.SetIsCa(dcl.ValueOrEmptyBool(o.IsCa))
	p.SetMaxIssuerPathLength(dcl.ValueOrEmptyInt64(o.MaxIssuerPathLength))
	return p
}

// CertificateAuthorityCaCertificateDescriptionsX509DescriptionPolicyIdsToProto converts a CertificateAuthorityCaCertificateDescriptionsX509DescriptionPolicyIds object to its proto representation.
func PrivatecaCertificateAuthorityCaCertificateDescriptionsX509DescriptionPolicyIdsToProto(o *privateca.CertificateAuthorityCaCertificateDescriptionsX509DescriptionPolicyIds) *privatecapb.PrivatecaCertificateAuthorityCaCertificateDescriptionsX509DescriptionPolicyIds {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCertificateAuthorityCaCertificateDescriptionsX509DescriptionPolicyIds{}
	sObjectIdPath := make([]int64, len(o.ObjectIdPath))
	for i, r := range o.ObjectIdPath {
		sObjectIdPath[i] = r
	}
	p.SetObjectIdPath(sObjectIdPath)
	return p
}

// CertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensionsToProto converts a CertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensions object to its proto representation.
func PrivatecaCertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensionsToProto(o *privateca.CertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensions) *privatecapb.PrivatecaCertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensions {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensions{}
	p.SetObjectId(PrivatecaCertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensionsObjectIdToProto(o.ObjectId))
	p.SetCritical(dcl.ValueOrEmptyBool(o.Critical))
	p.SetValue(dcl.ValueOrEmptyString(o.Value))
	return p
}

// CertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensionsObjectIdToProto converts a CertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensionsObjectId object to its proto representation.
func PrivatecaCertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensionsObjectIdToProto(o *privateca.CertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensionsObjectId) *privatecapb.PrivatecaCertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensionsObjectId {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensionsObjectId{}
	sObjectIdPath := make([]int64, len(o.ObjectIdPath))
	for i, r := range o.ObjectIdPath {
		sObjectIdPath[i] = r
	}
	p.SetObjectIdPath(sObjectIdPath)
	return p
}

// CertificateAuthorityCaCertificateDescriptionsPublicKeyToProto converts a CertificateAuthorityCaCertificateDescriptionsPublicKey object to its proto representation.
func PrivatecaCertificateAuthorityCaCertificateDescriptionsPublicKeyToProto(o *privateca.CertificateAuthorityCaCertificateDescriptionsPublicKey) *privatecapb.PrivatecaCertificateAuthorityCaCertificateDescriptionsPublicKey {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCertificateAuthorityCaCertificateDescriptionsPublicKey{}
	p.SetKey(dcl.ValueOrEmptyString(o.Key))
	p.SetFormat(PrivatecaCertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnumToProto(o.Format))
	return p
}

// CertificateAuthorityCaCertificateDescriptionsSubjectKeyIdToProto converts a CertificateAuthorityCaCertificateDescriptionsSubjectKeyId object to its proto representation.
func PrivatecaCertificateAuthorityCaCertificateDescriptionsSubjectKeyIdToProto(o *privateca.CertificateAuthorityCaCertificateDescriptionsSubjectKeyId) *privatecapb.PrivatecaCertificateAuthorityCaCertificateDescriptionsSubjectKeyId {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCertificateAuthorityCaCertificateDescriptionsSubjectKeyId{}
	p.SetKeyId(dcl.ValueOrEmptyString(o.KeyId))
	return p
}

// CertificateAuthorityCaCertificateDescriptionsAuthorityKeyIdToProto converts a CertificateAuthorityCaCertificateDescriptionsAuthorityKeyId object to its proto representation.
func PrivatecaCertificateAuthorityCaCertificateDescriptionsAuthorityKeyIdToProto(o *privateca.CertificateAuthorityCaCertificateDescriptionsAuthorityKeyId) *privatecapb.PrivatecaCertificateAuthorityCaCertificateDescriptionsAuthorityKeyId {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCertificateAuthorityCaCertificateDescriptionsAuthorityKeyId{}
	p.SetKeyId(dcl.ValueOrEmptyString(o.KeyId))
	return p
}

// CertificateAuthorityCaCertificateDescriptionsCertFingerprintToProto converts a CertificateAuthorityCaCertificateDescriptionsCertFingerprint object to its proto representation.
func PrivatecaCertificateAuthorityCaCertificateDescriptionsCertFingerprintToProto(o *privateca.CertificateAuthorityCaCertificateDescriptionsCertFingerprint) *privatecapb.PrivatecaCertificateAuthorityCaCertificateDescriptionsCertFingerprint {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCertificateAuthorityCaCertificateDescriptionsCertFingerprint{}
	p.SetSha256Hash(dcl.ValueOrEmptyString(o.Sha256Hash))
	return p
}

// CertificateAuthorityAccessUrlsToProto converts a CertificateAuthorityAccessUrls object to its proto representation.
func PrivatecaCertificateAuthorityAccessUrlsToProto(o *privateca.CertificateAuthorityAccessUrls) *privatecapb.PrivatecaCertificateAuthorityAccessUrls {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCertificateAuthorityAccessUrls{}
	p.SetCaCertificateAccessUrl(dcl.ValueOrEmptyString(o.CaCertificateAccessUrl))
	sCrlAccessUrls := make([]string, len(o.CrlAccessUrls))
	for i, r := range o.CrlAccessUrls {
		sCrlAccessUrls[i] = r
	}
	p.SetCrlAccessUrls(sCrlAccessUrls)
	return p
}

// CertificateAuthorityToProto converts a CertificateAuthority resource to its proto representation.
func CertificateAuthorityToProto(resource *privateca.CertificateAuthority) *privatecapb.PrivatecaCertificateAuthority {
	p := &privatecapb.PrivatecaCertificateAuthority{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetType(PrivatecaCertificateAuthorityTypeEnumToProto(resource.Type))
	p.SetConfig(PrivatecaCertificateAuthorityConfigToProto(resource.Config))
	p.SetLifetime(dcl.ValueOrEmptyString(resource.Lifetime))
	p.SetKeySpec(PrivatecaCertificateAuthorityKeySpecToProto(resource.KeySpec))
	p.SetSubordinateConfig(PrivatecaCertificateAuthoritySubordinateConfigToProto(resource.SubordinateConfig))
	p.SetTier(PrivatecaCertificateAuthorityTierEnumToProto(resource.Tier))
	p.SetState(PrivatecaCertificateAuthorityStateEnumToProto(resource.State))
	p.SetGcsBucket(dcl.ValueOrEmptyString(resource.GcsBucket))
	p.SetAccessUrls(PrivatecaCertificateAuthorityAccessUrlsToProto(resource.AccessUrls))
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
	sCaCertificateDescriptions := make([]*privatecapb.PrivatecaCertificateAuthorityCaCertificateDescriptions, len(resource.CaCertificateDescriptions))
	for i, r := range resource.CaCertificateDescriptions {
		sCaCertificateDescriptions[i] = PrivatecaCertificateAuthorityCaCertificateDescriptionsToProto(&r)
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
func (s *CertificateAuthorityServer) applyCertificateAuthority(ctx context.Context, c *privateca.Client, request *privatecapb.ApplyPrivatecaCertificateAuthorityRequest) (*privatecapb.PrivatecaCertificateAuthority, error) {
	p := ProtoToCertificateAuthority(request.GetResource())
	res, err := c.ApplyCertificateAuthority(ctx, p)
	if err != nil {
		return nil, err
	}
	r := CertificateAuthorityToProto(res)
	return r, nil
}

// applyPrivatecaCertificateAuthority handles the gRPC request by passing it to the underlying CertificateAuthority Apply() method.
func (s *CertificateAuthorityServer) ApplyPrivatecaCertificateAuthority(ctx context.Context, request *privatecapb.ApplyPrivatecaCertificateAuthorityRequest) (*privatecapb.PrivatecaCertificateAuthority, error) {
	cl, err := createConfigCertificateAuthority(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyCertificateAuthority(ctx, cl, request)
}

// DeleteCertificateAuthority handles the gRPC request by passing it to the underlying CertificateAuthority Delete() method.
func (s *CertificateAuthorityServer) DeletePrivatecaCertificateAuthority(ctx context.Context, request *privatecapb.DeletePrivatecaCertificateAuthorityRequest) (*emptypb.Empty, error) {

	cl, err := createConfigCertificateAuthority(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteCertificateAuthority(ctx, ProtoToCertificateAuthority(request.GetResource()))

}

// ListPrivatecaCertificateAuthority handles the gRPC request by passing it to the underlying CertificateAuthorityList() method.
func (s *CertificateAuthorityServer) ListPrivatecaCertificateAuthority(ctx context.Context, request *privatecapb.ListPrivatecaCertificateAuthorityRequest) (*privatecapb.ListPrivatecaCertificateAuthorityResponse, error) {
	cl, err := createConfigCertificateAuthority(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListCertificateAuthority(ctx, request.GetProject(), request.GetLocation(), request.GetCaPool())
	if err != nil {
		return nil, err
	}
	var protos []*privatecapb.PrivatecaCertificateAuthority
	for _, r := range resources.Items {
		rp := CertificateAuthorityToProto(r)
		protos = append(protos, rp)
	}
	p := &privatecapb.ListPrivatecaCertificateAuthorityResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigCertificateAuthority(ctx context.Context, service_account_file string) (*privateca.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return privateca.NewClient(conf), nil
}
