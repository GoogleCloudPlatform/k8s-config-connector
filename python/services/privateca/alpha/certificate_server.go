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

// CertificateServer implements the gRPC interface for Certificate.
type CertificateServer struct{}

// ProtoToCertificateConfigPublicKeyFormatEnum converts a CertificateConfigPublicKeyFormatEnum enum from its proto representation.
func ProtoToPrivatecaAlphaCertificateConfigPublicKeyFormatEnum(e alphapb.PrivatecaAlphaCertificateConfigPublicKeyFormatEnum) *alpha.CertificateConfigPublicKeyFormatEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.PrivatecaAlphaCertificateConfigPublicKeyFormatEnum_name[int32(e)]; ok {
		e := alpha.CertificateConfigPublicKeyFormatEnum(n[len("PrivatecaAlphaCertificateConfigPublicKeyFormatEnum"):])
		return &e
	}
	return nil
}

// ProtoToCertificateSubjectModeEnum converts a CertificateSubjectModeEnum enum from its proto representation.
func ProtoToPrivatecaAlphaCertificateSubjectModeEnum(e alphapb.PrivatecaAlphaCertificateSubjectModeEnum) *alpha.CertificateSubjectModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.PrivatecaAlphaCertificateSubjectModeEnum_name[int32(e)]; ok {
		e := alpha.CertificateSubjectModeEnum(n[len("PrivatecaAlphaCertificateSubjectModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToCertificateRevocationDetailsRevocationStateEnum converts a CertificateRevocationDetailsRevocationStateEnum enum from its proto representation.
func ProtoToPrivatecaAlphaCertificateRevocationDetailsRevocationStateEnum(e alphapb.PrivatecaAlphaCertificateRevocationDetailsRevocationStateEnum) *alpha.CertificateRevocationDetailsRevocationStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.PrivatecaAlphaCertificateRevocationDetailsRevocationStateEnum_name[int32(e)]; ok {
		e := alpha.CertificateRevocationDetailsRevocationStateEnum(n[len("PrivatecaAlphaCertificateRevocationDetailsRevocationStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToCertificateCertificateDescriptionPublicKeyFormatEnum converts a CertificateCertificateDescriptionPublicKeyFormatEnum enum from its proto representation.
func ProtoToPrivatecaAlphaCertificateCertificateDescriptionPublicKeyFormatEnum(e alphapb.PrivatecaAlphaCertificateCertificateDescriptionPublicKeyFormatEnum) *alpha.CertificateCertificateDescriptionPublicKeyFormatEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.PrivatecaAlphaCertificateCertificateDescriptionPublicKeyFormatEnum_name[int32(e)]; ok {
		e := alpha.CertificateCertificateDescriptionPublicKeyFormatEnum(n[len("PrivatecaAlphaCertificateCertificateDescriptionPublicKeyFormatEnum"):])
		return &e
	}
	return nil
}

// ProtoToCertificateConfig converts a CertificateConfig object from its proto representation.
func ProtoToPrivatecaAlphaCertificateConfig(p *alphapb.PrivatecaAlphaCertificateConfig) *alpha.CertificateConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.CertificateConfig{
		SubjectConfig: ProtoToPrivatecaAlphaCertificateConfigSubjectConfig(p.GetSubjectConfig()),
		X509Config:    ProtoToPrivatecaAlphaCertificateConfigX509Config(p.GetX509Config()),
		PublicKey:     ProtoToPrivatecaAlphaCertificateConfigPublicKey(p.GetPublicKey()),
	}
	return obj
}

// ProtoToCertificateConfigSubjectConfig converts a CertificateConfigSubjectConfig object from its proto representation.
func ProtoToPrivatecaAlphaCertificateConfigSubjectConfig(p *alphapb.PrivatecaAlphaCertificateConfigSubjectConfig) *alpha.CertificateConfigSubjectConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.CertificateConfigSubjectConfig{
		Subject:        ProtoToPrivatecaAlphaCertificateConfigSubjectConfigSubject(p.GetSubject()),
		SubjectAltName: ProtoToPrivatecaAlphaCertificateConfigSubjectConfigSubjectAltName(p.GetSubjectAltName()),
	}
	return obj
}

// ProtoToCertificateConfigSubjectConfigSubject converts a CertificateConfigSubjectConfigSubject object from its proto representation.
func ProtoToPrivatecaAlphaCertificateConfigSubjectConfigSubject(p *alphapb.PrivatecaAlphaCertificateConfigSubjectConfigSubject) *alpha.CertificateConfigSubjectConfigSubject {
	if p == nil {
		return nil
	}
	obj := &alpha.CertificateConfigSubjectConfigSubject{
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

// ProtoToCertificateConfigSubjectConfigSubjectAltName converts a CertificateConfigSubjectConfigSubjectAltName object from its proto representation.
func ProtoToPrivatecaAlphaCertificateConfigSubjectConfigSubjectAltName(p *alphapb.PrivatecaAlphaCertificateConfigSubjectConfigSubjectAltName) *alpha.CertificateConfigSubjectConfigSubjectAltName {
	if p == nil {
		return nil
	}
	obj := &alpha.CertificateConfigSubjectConfigSubjectAltName{}
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
	return obj
}

// ProtoToCertificateConfigX509Config converts a CertificateConfigX509Config object from its proto representation.
func ProtoToPrivatecaAlphaCertificateConfigX509Config(p *alphapb.PrivatecaAlphaCertificateConfigX509Config) *alpha.CertificateConfigX509Config {
	if p == nil {
		return nil
	}
	obj := &alpha.CertificateConfigX509Config{
		KeyUsage:  ProtoToPrivatecaAlphaCertificateConfigX509ConfigKeyUsage(p.GetKeyUsage()),
		CaOptions: ProtoToPrivatecaAlphaCertificateConfigX509ConfigCaOptions(p.GetCaOptions()),
	}
	for _, r := range p.GetPolicyIds() {
		obj.PolicyIds = append(obj.PolicyIds, *ProtoToPrivatecaAlphaCertificateConfigX509ConfigPolicyIds(r))
	}
	for _, r := range p.GetAiaOcspServers() {
		obj.AiaOcspServers = append(obj.AiaOcspServers, r)
	}
	for _, r := range p.GetAdditionalExtensions() {
		obj.AdditionalExtensions = append(obj.AdditionalExtensions, *ProtoToPrivatecaAlphaCertificateConfigX509ConfigAdditionalExtensions(r))
	}
	return obj
}

// ProtoToCertificateConfigX509ConfigKeyUsage converts a CertificateConfigX509ConfigKeyUsage object from its proto representation.
func ProtoToPrivatecaAlphaCertificateConfigX509ConfigKeyUsage(p *alphapb.PrivatecaAlphaCertificateConfigX509ConfigKeyUsage) *alpha.CertificateConfigX509ConfigKeyUsage {
	if p == nil {
		return nil
	}
	obj := &alpha.CertificateConfigX509ConfigKeyUsage{
		BaseKeyUsage:     ProtoToPrivatecaAlphaCertificateConfigX509ConfigKeyUsageBaseKeyUsage(p.GetBaseKeyUsage()),
		ExtendedKeyUsage: ProtoToPrivatecaAlphaCertificateConfigX509ConfigKeyUsageExtendedKeyUsage(p.GetExtendedKeyUsage()),
	}
	for _, r := range p.GetUnknownExtendedKeyUsages() {
		obj.UnknownExtendedKeyUsages = append(obj.UnknownExtendedKeyUsages, *ProtoToPrivatecaAlphaCertificateConfigX509ConfigKeyUsageUnknownExtendedKeyUsages(r))
	}
	return obj
}

// ProtoToCertificateConfigX509ConfigKeyUsageBaseKeyUsage converts a CertificateConfigX509ConfigKeyUsageBaseKeyUsage object from its proto representation.
func ProtoToPrivatecaAlphaCertificateConfigX509ConfigKeyUsageBaseKeyUsage(p *alphapb.PrivatecaAlphaCertificateConfigX509ConfigKeyUsageBaseKeyUsage) *alpha.CertificateConfigX509ConfigKeyUsageBaseKeyUsage {
	if p == nil {
		return nil
	}
	obj := &alpha.CertificateConfigX509ConfigKeyUsageBaseKeyUsage{
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

// ProtoToCertificateConfigX509ConfigKeyUsageExtendedKeyUsage converts a CertificateConfigX509ConfigKeyUsageExtendedKeyUsage object from its proto representation.
func ProtoToPrivatecaAlphaCertificateConfigX509ConfigKeyUsageExtendedKeyUsage(p *alphapb.PrivatecaAlphaCertificateConfigX509ConfigKeyUsageExtendedKeyUsage) *alpha.CertificateConfigX509ConfigKeyUsageExtendedKeyUsage {
	if p == nil {
		return nil
	}
	obj := &alpha.CertificateConfigX509ConfigKeyUsageExtendedKeyUsage{
		ServerAuth:      dcl.Bool(p.GetServerAuth()),
		ClientAuth:      dcl.Bool(p.GetClientAuth()),
		CodeSigning:     dcl.Bool(p.GetCodeSigning()),
		EmailProtection: dcl.Bool(p.GetEmailProtection()),
		TimeStamping:    dcl.Bool(p.GetTimeStamping()),
		OcspSigning:     dcl.Bool(p.GetOcspSigning()),
	}
	return obj
}

// ProtoToCertificateConfigX509ConfigKeyUsageUnknownExtendedKeyUsages converts a CertificateConfigX509ConfigKeyUsageUnknownExtendedKeyUsages object from its proto representation.
func ProtoToPrivatecaAlphaCertificateConfigX509ConfigKeyUsageUnknownExtendedKeyUsages(p *alphapb.PrivatecaAlphaCertificateConfigX509ConfigKeyUsageUnknownExtendedKeyUsages) *alpha.CertificateConfigX509ConfigKeyUsageUnknownExtendedKeyUsages {
	if p == nil {
		return nil
	}
	obj := &alpha.CertificateConfigX509ConfigKeyUsageUnknownExtendedKeyUsages{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCertificateConfigX509ConfigCaOptions converts a CertificateConfigX509ConfigCaOptions object from its proto representation.
func ProtoToPrivatecaAlphaCertificateConfigX509ConfigCaOptions(p *alphapb.PrivatecaAlphaCertificateConfigX509ConfigCaOptions) *alpha.CertificateConfigX509ConfigCaOptions {
	if p == nil {
		return nil
	}
	obj := &alpha.CertificateConfigX509ConfigCaOptions{
		IsCa:                    dcl.Bool(p.GetIsCa()),
		NonCa:                   dcl.Bool(p.GetNonCa()),
		MaxIssuerPathLength:     dcl.Int64OrNil(p.GetMaxIssuerPathLength()),
		ZeroMaxIssuerPathLength: dcl.Bool(p.GetZeroMaxIssuerPathLength()),
	}
	return obj
}

// ProtoToCertificateConfigX509ConfigPolicyIds converts a CertificateConfigX509ConfigPolicyIds object from its proto representation.
func ProtoToPrivatecaAlphaCertificateConfigX509ConfigPolicyIds(p *alphapb.PrivatecaAlphaCertificateConfigX509ConfigPolicyIds) *alpha.CertificateConfigX509ConfigPolicyIds {
	if p == nil {
		return nil
	}
	obj := &alpha.CertificateConfigX509ConfigPolicyIds{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCertificateConfigX509ConfigAdditionalExtensions converts a CertificateConfigX509ConfigAdditionalExtensions object from its proto representation.
func ProtoToPrivatecaAlphaCertificateConfigX509ConfigAdditionalExtensions(p *alphapb.PrivatecaAlphaCertificateConfigX509ConfigAdditionalExtensions) *alpha.CertificateConfigX509ConfigAdditionalExtensions {
	if p == nil {
		return nil
	}
	obj := &alpha.CertificateConfigX509ConfigAdditionalExtensions{
		ObjectId: ProtoToPrivatecaAlphaCertificateConfigX509ConfigAdditionalExtensionsObjectId(p.GetObjectId()),
		Critical: dcl.Bool(p.GetCritical()),
		Value:    dcl.StringOrNil(p.GetValue()),
	}
	return obj
}

// ProtoToCertificateConfigX509ConfigAdditionalExtensionsObjectId converts a CertificateConfigX509ConfigAdditionalExtensionsObjectId object from its proto representation.
func ProtoToPrivatecaAlphaCertificateConfigX509ConfigAdditionalExtensionsObjectId(p *alphapb.PrivatecaAlphaCertificateConfigX509ConfigAdditionalExtensionsObjectId) *alpha.CertificateConfigX509ConfigAdditionalExtensionsObjectId {
	if p == nil {
		return nil
	}
	obj := &alpha.CertificateConfigX509ConfigAdditionalExtensionsObjectId{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCertificateConfigPublicKey converts a CertificateConfigPublicKey object from its proto representation.
func ProtoToPrivatecaAlphaCertificateConfigPublicKey(p *alphapb.PrivatecaAlphaCertificateConfigPublicKey) *alpha.CertificateConfigPublicKey {
	if p == nil {
		return nil
	}
	obj := &alpha.CertificateConfigPublicKey{
		Key:    dcl.StringOrNil(p.GetKey()),
		Format: ProtoToPrivatecaAlphaCertificateConfigPublicKeyFormatEnum(p.GetFormat()),
	}
	return obj
}

// ProtoToCertificateRevocationDetails converts a CertificateRevocationDetails object from its proto representation.
func ProtoToPrivatecaAlphaCertificateRevocationDetails(p *alphapb.PrivatecaAlphaCertificateRevocationDetails) *alpha.CertificateRevocationDetails {
	if p == nil {
		return nil
	}
	obj := &alpha.CertificateRevocationDetails{
		RevocationState: ProtoToPrivatecaAlphaCertificateRevocationDetailsRevocationStateEnum(p.GetRevocationState()),
		RevocationTime:  dcl.StringOrNil(p.GetRevocationTime()),
	}
	return obj
}

// ProtoToCertificateCertificateDescription converts a CertificateCertificateDescription object from its proto representation.
func ProtoToPrivatecaAlphaCertificateCertificateDescription(p *alphapb.PrivatecaAlphaCertificateCertificateDescription) *alpha.CertificateCertificateDescription {
	if p == nil {
		return nil
	}
	obj := &alpha.CertificateCertificateDescription{
		SubjectDescription: ProtoToPrivatecaAlphaCertificateCertificateDescriptionSubjectDescription(p.GetSubjectDescription()),
		X509Description:    ProtoToPrivatecaAlphaCertificateCertificateDescriptionX509Description(p.GetX509Description()),
		PublicKey:          ProtoToPrivatecaAlphaCertificateCertificateDescriptionPublicKey(p.GetPublicKey()),
		SubjectKeyId:       ProtoToPrivatecaAlphaCertificateCertificateDescriptionSubjectKeyId(p.GetSubjectKeyId()),
		AuthorityKeyId:     ProtoToPrivatecaAlphaCertificateCertificateDescriptionAuthorityKeyId(p.GetAuthorityKeyId()),
		CertFingerprint:    ProtoToPrivatecaAlphaCertificateCertificateDescriptionCertFingerprint(p.GetCertFingerprint()),
	}
	for _, r := range p.GetCrlDistributionPoints() {
		obj.CrlDistributionPoints = append(obj.CrlDistributionPoints, r)
	}
	for _, r := range p.GetAiaIssuingCertificateUrls() {
		obj.AiaIssuingCertificateUrls = append(obj.AiaIssuingCertificateUrls, r)
	}
	return obj
}

// ProtoToCertificateCertificateDescriptionSubjectDescription converts a CertificateCertificateDescriptionSubjectDescription object from its proto representation.
func ProtoToPrivatecaAlphaCertificateCertificateDescriptionSubjectDescription(p *alphapb.PrivatecaAlphaCertificateCertificateDescriptionSubjectDescription) *alpha.CertificateCertificateDescriptionSubjectDescription {
	if p == nil {
		return nil
	}
	obj := &alpha.CertificateCertificateDescriptionSubjectDescription{
		Subject:         ProtoToPrivatecaAlphaCertificateCertificateDescriptionSubjectDescriptionSubject(p.GetSubject()),
		SubjectAltName:  ProtoToPrivatecaAlphaCertificateCertificateDescriptionSubjectDescriptionSubjectAltName(p.GetSubjectAltName()),
		HexSerialNumber: dcl.StringOrNil(p.GetHexSerialNumber()),
		Lifetime:        dcl.StringOrNil(p.GetLifetime()),
		NotBeforeTime:   dcl.StringOrNil(p.GetNotBeforeTime()),
		NotAfterTime:    dcl.StringOrNil(p.GetNotAfterTime()),
	}
	return obj
}

// ProtoToCertificateCertificateDescriptionSubjectDescriptionSubject converts a CertificateCertificateDescriptionSubjectDescriptionSubject object from its proto representation.
func ProtoToPrivatecaAlphaCertificateCertificateDescriptionSubjectDescriptionSubject(p *alphapb.PrivatecaAlphaCertificateCertificateDescriptionSubjectDescriptionSubject) *alpha.CertificateCertificateDescriptionSubjectDescriptionSubject {
	if p == nil {
		return nil
	}
	obj := &alpha.CertificateCertificateDescriptionSubjectDescriptionSubject{
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

// ProtoToCertificateCertificateDescriptionSubjectDescriptionSubjectAltName converts a CertificateCertificateDescriptionSubjectDescriptionSubjectAltName object from its proto representation.
func ProtoToPrivatecaAlphaCertificateCertificateDescriptionSubjectDescriptionSubjectAltName(p *alphapb.PrivatecaAlphaCertificateCertificateDescriptionSubjectDescriptionSubjectAltName) *alpha.CertificateCertificateDescriptionSubjectDescriptionSubjectAltName {
	if p == nil {
		return nil
	}
	obj := &alpha.CertificateCertificateDescriptionSubjectDescriptionSubjectAltName{}
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
		obj.CustomSans = append(obj.CustomSans, *ProtoToPrivatecaAlphaCertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSans(r))
	}
	return obj
}

// ProtoToCertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSans converts a CertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSans object from its proto representation.
func ProtoToPrivatecaAlphaCertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSans(p *alphapb.PrivatecaAlphaCertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSans) *alpha.CertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSans {
	if p == nil {
		return nil
	}
	obj := &alpha.CertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSans{
		ObjectId: ProtoToPrivatecaAlphaCertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansObjectId(p.GetObjectId()),
		Critical: dcl.Bool(p.GetCritical()),
		Value:    dcl.StringOrNil(p.GetValue()),
	}
	return obj
}

// ProtoToCertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansObjectId converts a CertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansObjectId object from its proto representation.
func ProtoToPrivatecaAlphaCertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansObjectId(p *alphapb.PrivatecaAlphaCertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansObjectId) *alpha.CertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansObjectId {
	if p == nil {
		return nil
	}
	obj := &alpha.CertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansObjectId{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCertificateCertificateDescriptionX509Description converts a CertificateCertificateDescriptionX509Description object from its proto representation.
func ProtoToPrivatecaAlphaCertificateCertificateDescriptionX509Description(p *alphapb.PrivatecaAlphaCertificateCertificateDescriptionX509Description) *alpha.CertificateCertificateDescriptionX509Description {
	if p == nil {
		return nil
	}
	obj := &alpha.CertificateCertificateDescriptionX509Description{
		KeyUsage:  ProtoToPrivatecaAlphaCertificateCertificateDescriptionX509DescriptionKeyUsage(p.GetKeyUsage()),
		CaOptions: ProtoToPrivatecaAlphaCertificateCertificateDescriptionX509DescriptionCaOptions(p.GetCaOptions()),
	}
	for _, r := range p.GetPolicyIds() {
		obj.PolicyIds = append(obj.PolicyIds, *ProtoToPrivatecaAlphaCertificateCertificateDescriptionX509DescriptionPolicyIds(r))
	}
	for _, r := range p.GetAiaOcspServers() {
		obj.AiaOcspServers = append(obj.AiaOcspServers, r)
	}
	for _, r := range p.GetAdditionalExtensions() {
		obj.AdditionalExtensions = append(obj.AdditionalExtensions, *ProtoToPrivatecaAlphaCertificateCertificateDescriptionX509DescriptionAdditionalExtensions(r))
	}
	return obj
}

// ProtoToCertificateCertificateDescriptionX509DescriptionKeyUsage converts a CertificateCertificateDescriptionX509DescriptionKeyUsage object from its proto representation.
func ProtoToPrivatecaAlphaCertificateCertificateDescriptionX509DescriptionKeyUsage(p *alphapb.PrivatecaAlphaCertificateCertificateDescriptionX509DescriptionKeyUsage) *alpha.CertificateCertificateDescriptionX509DescriptionKeyUsage {
	if p == nil {
		return nil
	}
	obj := &alpha.CertificateCertificateDescriptionX509DescriptionKeyUsage{
		BaseKeyUsage:     ProtoToPrivatecaAlphaCertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsage(p.GetBaseKeyUsage()),
		ExtendedKeyUsage: ProtoToPrivatecaAlphaCertificateCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsage(p.GetExtendedKeyUsage()),
	}
	for _, r := range p.GetUnknownExtendedKeyUsages() {
		obj.UnknownExtendedKeyUsages = append(obj.UnknownExtendedKeyUsages, *ProtoToPrivatecaAlphaCertificateCertificateDescriptionX509DescriptionKeyUsageUnknownExtendedKeyUsages(r))
	}
	return obj
}

// ProtoToCertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsage converts a CertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsage object from its proto representation.
func ProtoToPrivatecaAlphaCertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsage(p *alphapb.PrivatecaAlphaCertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsage) *alpha.CertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsage {
	if p == nil {
		return nil
	}
	obj := &alpha.CertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsage{
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

// ProtoToCertificateCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsage converts a CertificateCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsage object from its proto representation.
func ProtoToPrivatecaAlphaCertificateCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsage(p *alphapb.PrivatecaAlphaCertificateCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsage) *alpha.CertificateCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsage {
	if p == nil {
		return nil
	}
	obj := &alpha.CertificateCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsage{
		ServerAuth:      dcl.Bool(p.GetServerAuth()),
		ClientAuth:      dcl.Bool(p.GetClientAuth()),
		CodeSigning:     dcl.Bool(p.GetCodeSigning()),
		EmailProtection: dcl.Bool(p.GetEmailProtection()),
		TimeStamping:    dcl.Bool(p.GetTimeStamping()),
		OcspSigning:     dcl.Bool(p.GetOcspSigning()),
	}
	return obj
}

// ProtoToCertificateCertificateDescriptionX509DescriptionKeyUsageUnknownExtendedKeyUsages converts a CertificateCertificateDescriptionX509DescriptionKeyUsageUnknownExtendedKeyUsages object from its proto representation.
func ProtoToPrivatecaAlphaCertificateCertificateDescriptionX509DescriptionKeyUsageUnknownExtendedKeyUsages(p *alphapb.PrivatecaAlphaCertificateCertificateDescriptionX509DescriptionKeyUsageUnknownExtendedKeyUsages) *alpha.CertificateCertificateDescriptionX509DescriptionKeyUsageUnknownExtendedKeyUsages {
	if p == nil {
		return nil
	}
	obj := &alpha.CertificateCertificateDescriptionX509DescriptionKeyUsageUnknownExtendedKeyUsages{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCertificateCertificateDescriptionX509DescriptionCaOptions converts a CertificateCertificateDescriptionX509DescriptionCaOptions object from its proto representation.
func ProtoToPrivatecaAlphaCertificateCertificateDescriptionX509DescriptionCaOptions(p *alphapb.PrivatecaAlphaCertificateCertificateDescriptionX509DescriptionCaOptions) *alpha.CertificateCertificateDescriptionX509DescriptionCaOptions {
	if p == nil {
		return nil
	}
	obj := &alpha.CertificateCertificateDescriptionX509DescriptionCaOptions{
		IsCa:                dcl.Bool(p.GetIsCa()),
		MaxIssuerPathLength: dcl.Int64OrNil(p.GetMaxIssuerPathLength()),
	}
	return obj
}

// ProtoToCertificateCertificateDescriptionX509DescriptionPolicyIds converts a CertificateCertificateDescriptionX509DescriptionPolicyIds object from its proto representation.
func ProtoToPrivatecaAlphaCertificateCertificateDescriptionX509DescriptionPolicyIds(p *alphapb.PrivatecaAlphaCertificateCertificateDescriptionX509DescriptionPolicyIds) *alpha.CertificateCertificateDescriptionX509DescriptionPolicyIds {
	if p == nil {
		return nil
	}
	obj := &alpha.CertificateCertificateDescriptionX509DescriptionPolicyIds{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCertificateCertificateDescriptionX509DescriptionAdditionalExtensions converts a CertificateCertificateDescriptionX509DescriptionAdditionalExtensions object from its proto representation.
func ProtoToPrivatecaAlphaCertificateCertificateDescriptionX509DescriptionAdditionalExtensions(p *alphapb.PrivatecaAlphaCertificateCertificateDescriptionX509DescriptionAdditionalExtensions) *alpha.CertificateCertificateDescriptionX509DescriptionAdditionalExtensions {
	if p == nil {
		return nil
	}
	obj := &alpha.CertificateCertificateDescriptionX509DescriptionAdditionalExtensions{
		ObjectId: ProtoToPrivatecaAlphaCertificateCertificateDescriptionX509DescriptionAdditionalExtensionsObjectId(p.GetObjectId()),
		Critical: dcl.Bool(p.GetCritical()),
		Value:    dcl.StringOrNil(p.GetValue()),
	}
	return obj
}

// ProtoToCertificateCertificateDescriptionX509DescriptionAdditionalExtensionsObjectId converts a CertificateCertificateDescriptionX509DescriptionAdditionalExtensionsObjectId object from its proto representation.
func ProtoToPrivatecaAlphaCertificateCertificateDescriptionX509DescriptionAdditionalExtensionsObjectId(p *alphapb.PrivatecaAlphaCertificateCertificateDescriptionX509DescriptionAdditionalExtensionsObjectId) *alpha.CertificateCertificateDescriptionX509DescriptionAdditionalExtensionsObjectId {
	if p == nil {
		return nil
	}
	obj := &alpha.CertificateCertificateDescriptionX509DescriptionAdditionalExtensionsObjectId{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCertificateCertificateDescriptionPublicKey converts a CertificateCertificateDescriptionPublicKey object from its proto representation.
func ProtoToPrivatecaAlphaCertificateCertificateDescriptionPublicKey(p *alphapb.PrivatecaAlphaCertificateCertificateDescriptionPublicKey) *alpha.CertificateCertificateDescriptionPublicKey {
	if p == nil {
		return nil
	}
	obj := &alpha.CertificateCertificateDescriptionPublicKey{
		Key:    dcl.StringOrNil(p.GetKey()),
		Format: ProtoToPrivatecaAlphaCertificateCertificateDescriptionPublicKeyFormatEnum(p.GetFormat()),
	}
	return obj
}

// ProtoToCertificateCertificateDescriptionSubjectKeyId converts a CertificateCertificateDescriptionSubjectKeyId object from its proto representation.
func ProtoToPrivatecaAlphaCertificateCertificateDescriptionSubjectKeyId(p *alphapb.PrivatecaAlphaCertificateCertificateDescriptionSubjectKeyId) *alpha.CertificateCertificateDescriptionSubjectKeyId {
	if p == nil {
		return nil
	}
	obj := &alpha.CertificateCertificateDescriptionSubjectKeyId{
		KeyId: dcl.StringOrNil(p.GetKeyId()),
	}
	return obj
}

// ProtoToCertificateCertificateDescriptionAuthorityKeyId converts a CertificateCertificateDescriptionAuthorityKeyId object from its proto representation.
func ProtoToPrivatecaAlphaCertificateCertificateDescriptionAuthorityKeyId(p *alphapb.PrivatecaAlphaCertificateCertificateDescriptionAuthorityKeyId) *alpha.CertificateCertificateDescriptionAuthorityKeyId {
	if p == nil {
		return nil
	}
	obj := &alpha.CertificateCertificateDescriptionAuthorityKeyId{
		KeyId: dcl.StringOrNil(p.GetKeyId()),
	}
	return obj
}

// ProtoToCertificateCertificateDescriptionCertFingerprint converts a CertificateCertificateDescriptionCertFingerprint object from its proto representation.
func ProtoToPrivatecaAlphaCertificateCertificateDescriptionCertFingerprint(p *alphapb.PrivatecaAlphaCertificateCertificateDescriptionCertFingerprint) *alpha.CertificateCertificateDescriptionCertFingerprint {
	if p == nil {
		return nil
	}
	obj := &alpha.CertificateCertificateDescriptionCertFingerprint{
		Sha256Hash: dcl.StringOrNil(p.GetSha256Hash()),
	}
	return obj
}

// ProtoToCertificate converts a Certificate resource from its proto representation.
func ProtoToCertificate(p *alphapb.PrivatecaAlphaCertificate) *alpha.Certificate {
	obj := &alpha.Certificate{
		Name:                       dcl.StringOrNil(p.GetName()),
		PemCsr:                     dcl.StringOrNil(p.GetPemCsr()),
		Config:                     ProtoToPrivatecaAlphaCertificateConfig(p.GetConfig()),
		IssuerCertificateAuthority: dcl.StringOrNil(p.GetIssuerCertificateAuthority()),
		Lifetime:                   dcl.StringOrNil(p.GetLifetime()),
		CertificateTemplate:        dcl.StringOrNil(p.GetCertificateTemplate()),
		SubjectMode:                ProtoToPrivatecaAlphaCertificateSubjectModeEnum(p.GetSubjectMode()),
		RevocationDetails:          ProtoToPrivatecaAlphaCertificateRevocationDetails(p.GetRevocationDetails()),
		PemCertificate:             dcl.StringOrNil(p.GetPemCertificate()),
		CertificateDescription:     ProtoToPrivatecaAlphaCertificateCertificateDescription(p.GetCertificateDescription()),
		CreateTime:                 dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:                 dcl.StringOrNil(p.GetUpdateTime()),
		Project:                    dcl.StringOrNil(p.GetProject()),
		Location:                   dcl.StringOrNil(p.GetLocation()),
		CaPool:                     dcl.StringOrNil(p.GetCaPool()),
		CertificateAuthority:       dcl.StringOrNil(p.GetCertificateAuthority()),
	}
	for _, r := range p.GetPemCertificateChain() {
		obj.PemCertificateChain = append(obj.PemCertificateChain, r)
	}
	return obj
}

// CertificateConfigPublicKeyFormatEnumToProto converts a CertificateConfigPublicKeyFormatEnum enum to its proto representation.
func PrivatecaAlphaCertificateConfigPublicKeyFormatEnumToProto(e *alpha.CertificateConfigPublicKeyFormatEnum) alphapb.PrivatecaAlphaCertificateConfigPublicKeyFormatEnum {
	if e == nil {
		return alphapb.PrivatecaAlphaCertificateConfigPublicKeyFormatEnum(0)
	}
	if v, ok := alphapb.PrivatecaAlphaCertificateConfigPublicKeyFormatEnum_value["CertificateConfigPublicKeyFormatEnum"+string(*e)]; ok {
		return alphapb.PrivatecaAlphaCertificateConfigPublicKeyFormatEnum(v)
	}
	return alphapb.PrivatecaAlphaCertificateConfigPublicKeyFormatEnum(0)
}

// CertificateSubjectModeEnumToProto converts a CertificateSubjectModeEnum enum to its proto representation.
func PrivatecaAlphaCertificateSubjectModeEnumToProto(e *alpha.CertificateSubjectModeEnum) alphapb.PrivatecaAlphaCertificateSubjectModeEnum {
	if e == nil {
		return alphapb.PrivatecaAlphaCertificateSubjectModeEnum(0)
	}
	if v, ok := alphapb.PrivatecaAlphaCertificateSubjectModeEnum_value["CertificateSubjectModeEnum"+string(*e)]; ok {
		return alphapb.PrivatecaAlphaCertificateSubjectModeEnum(v)
	}
	return alphapb.PrivatecaAlphaCertificateSubjectModeEnum(0)
}

// CertificateRevocationDetailsRevocationStateEnumToProto converts a CertificateRevocationDetailsRevocationStateEnum enum to its proto representation.
func PrivatecaAlphaCertificateRevocationDetailsRevocationStateEnumToProto(e *alpha.CertificateRevocationDetailsRevocationStateEnum) alphapb.PrivatecaAlphaCertificateRevocationDetailsRevocationStateEnum {
	if e == nil {
		return alphapb.PrivatecaAlphaCertificateRevocationDetailsRevocationStateEnum(0)
	}
	if v, ok := alphapb.PrivatecaAlphaCertificateRevocationDetailsRevocationStateEnum_value["CertificateRevocationDetailsRevocationStateEnum"+string(*e)]; ok {
		return alphapb.PrivatecaAlphaCertificateRevocationDetailsRevocationStateEnum(v)
	}
	return alphapb.PrivatecaAlphaCertificateRevocationDetailsRevocationStateEnum(0)
}

// CertificateCertificateDescriptionPublicKeyFormatEnumToProto converts a CertificateCertificateDescriptionPublicKeyFormatEnum enum to its proto representation.
func PrivatecaAlphaCertificateCertificateDescriptionPublicKeyFormatEnumToProto(e *alpha.CertificateCertificateDescriptionPublicKeyFormatEnum) alphapb.PrivatecaAlphaCertificateCertificateDescriptionPublicKeyFormatEnum {
	if e == nil {
		return alphapb.PrivatecaAlphaCertificateCertificateDescriptionPublicKeyFormatEnum(0)
	}
	if v, ok := alphapb.PrivatecaAlphaCertificateCertificateDescriptionPublicKeyFormatEnum_value["CertificateCertificateDescriptionPublicKeyFormatEnum"+string(*e)]; ok {
		return alphapb.PrivatecaAlphaCertificateCertificateDescriptionPublicKeyFormatEnum(v)
	}
	return alphapb.PrivatecaAlphaCertificateCertificateDescriptionPublicKeyFormatEnum(0)
}

// CertificateConfigToProto converts a CertificateConfig object to its proto representation.
func PrivatecaAlphaCertificateConfigToProto(o *alpha.CertificateConfig) *alphapb.PrivatecaAlphaCertificateConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCertificateConfig{}
	p.SetSubjectConfig(PrivatecaAlphaCertificateConfigSubjectConfigToProto(o.SubjectConfig))
	p.SetX509Config(PrivatecaAlphaCertificateConfigX509ConfigToProto(o.X509Config))
	p.SetPublicKey(PrivatecaAlphaCertificateConfigPublicKeyToProto(o.PublicKey))
	return p
}

// CertificateConfigSubjectConfigToProto converts a CertificateConfigSubjectConfig object to its proto representation.
func PrivatecaAlphaCertificateConfigSubjectConfigToProto(o *alpha.CertificateConfigSubjectConfig) *alphapb.PrivatecaAlphaCertificateConfigSubjectConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCertificateConfigSubjectConfig{}
	p.SetSubject(PrivatecaAlphaCertificateConfigSubjectConfigSubjectToProto(o.Subject))
	p.SetSubjectAltName(PrivatecaAlphaCertificateConfigSubjectConfigSubjectAltNameToProto(o.SubjectAltName))
	return p
}

// CertificateConfigSubjectConfigSubjectToProto converts a CertificateConfigSubjectConfigSubject object to its proto representation.
func PrivatecaAlphaCertificateConfigSubjectConfigSubjectToProto(o *alpha.CertificateConfigSubjectConfigSubject) *alphapb.PrivatecaAlphaCertificateConfigSubjectConfigSubject {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCertificateConfigSubjectConfigSubject{}
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

// CertificateConfigSubjectConfigSubjectAltNameToProto converts a CertificateConfigSubjectConfigSubjectAltName object to its proto representation.
func PrivatecaAlphaCertificateConfigSubjectConfigSubjectAltNameToProto(o *alpha.CertificateConfigSubjectConfigSubjectAltName) *alphapb.PrivatecaAlphaCertificateConfigSubjectConfigSubjectAltName {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCertificateConfigSubjectConfigSubjectAltName{}
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
	return p
}

// CertificateConfigX509ConfigToProto converts a CertificateConfigX509Config object to its proto representation.
func PrivatecaAlphaCertificateConfigX509ConfigToProto(o *alpha.CertificateConfigX509Config) *alphapb.PrivatecaAlphaCertificateConfigX509Config {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCertificateConfigX509Config{}
	p.SetKeyUsage(PrivatecaAlphaCertificateConfigX509ConfigKeyUsageToProto(o.KeyUsage))
	p.SetCaOptions(PrivatecaAlphaCertificateConfigX509ConfigCaOptionsToProto(o.CaOptions))
	sPolicyIds := make([]*alphapb.PrivatecaAlphaCertificateConfigX509ConfigPolicyIds, len(o.PolicyIds))
	for i, r := range o.PolicyIds {
		sPolicyIds[i] = PrivatecaAlphaCertificateConfigX509ConfigPolicyIdsToProto(&r)
	}
	p.SetPolicyIds(sPolicyIds)
	sAiaOcspServers := make([]string, len(o.AiaOcspServers))
	for i, r := range o.AiaOcspServers {
		sAiaOcspServers[i] = r
	}
	p.SetAiaOcspServers(sAiaOcspServers)
	sAdditionalExtensions := make([]*alphapb.PrivatecaAlphaCertificateConfigX509ConfigAdditionalExtensions, len(o.AdditionalExtensions))
	for i, r := range o.AdditionalExtensions {
		sAdditionalExtensions[i] = PrivatecaAlphaCertificateConfigX509ConfigAdditionalExtensionsToProto(&r)
	}
	p.SetAdditionalExtensions(sAdditionalExtensions)
	return p
}

// CertificateConfigX509ConfigKeyUsageToProto converts a CertificateConfigX509ConfigKeyUsage object to its proto representation.
func PrivatecaAlphaCertificateConfigX509ConfigKeyUsageToProto(o *alpha.CertificateConfigX509ConfigKeyUsage) *alphapb.PrivatecaAlphaCertificateConfigX509ConfigKeyUsage {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCertificateConfigX509ConfigKeyUsage{}
	p.SetBaseKeyUsage(PrivatecaAlphaCertificateConfigX509ConfigKeyUsageBaseKeyUsageToProto(o.BaseKeyUsage))
	p.SetExtendedKeyUsage(PrivatecaAlphaCertificateConfigX509ConfigKeyUsageExtendedKeyUsageToProto(o.ExtendedKeyUsage))
	sUnknownExtendedKeyUsages := make([]*alphapb.PrivatecaAlphaCertificateConfigX509ConfigKeyUsageUnknownExtendedKeyUsages, len(o.UnknownExtendedKeyUsages))
	for i, r := range o.UnknownExtendedKeyUsages {
		sUnknownExtendedKeyUsages[i] = PrivatecaAlphaCertificateConfigX509ConfigKeyUsageUnknownExtendedKeyUsagesToProto(&r)
	}
	p.SetUnknownExtendedKeyUsages(sUnknownExtendedKeyUsages)
	return p
}

// CertificateConfigX509ConfigKeyUsageBaseKeyUsageToProto converts a CertificateConfigX509ConfigKeyUsageBaseKeyUsage object to its proto representation.
func PrivatecaAlphaCertificateConfigX509ConfigKeyUsageBaseKeyUsageToProto(o *alpha.CertificateConfigX509ConfigKeyUsageBaseKeyUsage) *alphapb.PrivatecaAlphaCertificateConfigX509ConfigKeyUsageBaseKeyUsage {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCertificateConfigX509ConfigKeyUsageBaseKeyUsage{}
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

// CertificateConfigX509ConfigKeyUsageExtendedKeyUsageToProto converts a CertificateConfigX509ConfigKeyUsageExtendedKeyUsage object to its proto representation.
func PrivatecaAlphaCertificateConfigX509ConfigKeyUsageExtendedKeyUsageToProto(o *alpha.CertificateConfigX509ConfigKeyUsageExtendedKeyUsage) *alphapb.PrivatecaAlphaCertificateConfigX509ConfigKeyUsageExtendedKeyUsage {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCertificateConfigX509ConfigKeyUsageExtendedKeyUsage{}
	p.SetServerAuth(dcl.ValueOrEmptyBool(o.ServerAuth))
	p.SetClientAuth(dcl.ValueOrEmptyBool(o.ClientAuth))
	p.SetCodeSigning(dcl.ValueOrEmptyBool(o.CodeSigning))
	p.SetEmailProtection(dcl.ValueOrEmptyBool(o.EmailProtection))
	p.SetTimeStamping(dcl.ValueOrEmptyBool(o.TimeStamping))
	p.SetOcspSigning(dcl.ValueOrEmptyBool(o.OcspSigning))
	return p
}

// CertificateConfigX509ConfigKeyUsageUnknownExtendedKeyUsagesToProto converts a CertificateConfigX509ConfigKeyUsageUnknownExtendedKeyUsages object to its proto representation.
func PrivatecaAlphaCertificateConfigX509ConfigKeyUsageUnknownExtendedKeyUsagesToProto(o *alpha.CertificateConfigX509ConfigKeyUsageUnknownExtendedKeyUsages) *alphapb.PrivatecaAlphaCertificateConfigX509ConfigKeyUsageUnknownExtendedKeyUsages {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCertificateConfigX509ConfigKeyUsageUnknownExtendedKeyUsages{}
	sObjectIdPath := make([]int64, len(o.ObjectIdPath))
	for i, r := range o.ObjectIdPath {
		sObjectIdPath[i] = r
	}
	p.SetObjectIdPath(sObjectIdPath)
	return p
}

// CertificateConfigX509ConfigCaOptionsToProto converts a CertificateConfigX509ConfigCaOptions object to its proto representation.
func PrivatecaAlphaCertificateConfigX509ConfigCaOptionsToProto(o *alpha.CertificateConfigX509ConfigCaOptions) *alphapb.PrivatecaAlphaCertificateConfigX509ConfigCaOptions {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCertificateConfigX509ConfigCaOptions{}
	p.SetIsCa(dcl.ValueOrEmptyBool(o.IsCa))
	p.SetNonCa(dcl.ValueOrEmptyBool(o.NonCa))
	p.SetMaxIssuerPathLength(dcl.ValueOrEmptyInt64(o.MaxIssuerPathLength))
	p.SetZeroMaxIssuerPathLength(dcl.ValueOrEmptyBool(o.ZeroMaxIssuerPathLength))
	return p
}

// CertificateConfigX509ConfigPolicyIdsToProto converts a CertificateConfigX509ConfigPolicyIds object to its proto representation.
func PrivatecaAlphaCertificateConfigX509ConfigPolicyIdsToProto(o *alpha.CertificateConfigX509ConfigPolicyIds) *alphapb.PrivatecaAlphaCertificateConfigX509ConfigPolicyIds {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCertificateConfigX509ConfigPolicyIds{}
	sObjectIdPath := make([]int64, len(o.ObjectIdPath))
	for i, r := range o.ObjectIdPath {
		sObjectIdPath[i] = r
	}
	p.SetObjectIdPath(sObjectIdPath)
	return p
}

// CertificateConfigX509ConfigAdditionalExtensionsToProto converts a CertificateConfigX509ConfigAdditionalExtensions object to its proto representation.
func PrivatecaAlphaCertificateConfigX509ConfigAdditionalExtensionsToProto(o *alpha.CertificateConfigX509ConfigAdditionalExtensions) *alphapb.PrivatecaAlphaCertificateConfigX509ConfigAdditionalExtensions {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCertificateConfigX509ConfigAdditionalExtensions{}
	p.SetObjectId(PrivatecaAlphaCertificateConfigX509ConfigAdditionalExtensionsObjectIdToProto(o.ObjectId))
	p.SetCritical(dcl.ValueOrEmptyBool(o.Critical))
	p.SetValue(dcl.ValueOrEmptyString(o.Value))
	return p
}

// CertificateConfigX509ConfigAdditionalExtensionsObjectIdToProto converts a CertificateConfigX509ConfigAdditionalExtensionsObjectId object to its proto representation.
func PrivatecaAlphaCertificateConfigX509ConfigAdditionalExtensionsObjectIdToProto(o *alpha.CertificateConfigX509ConfigAdditionalExtensionsObjectId) *alphapb.PrivatecaAlphaCertificateConfigX509ConfigAdditionalExtensionsObjectId {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCertificateConfigX509ConfigAdditionalExtensionsObjectId{}
	sObjectIdPath := make([]int64, len(o.ObjectIdPath))
	for i, r := range o.ObjectIdPath {
		sObjectIdPath[i] = r
	}
	p.SetObjectIdPath(sObjectIdPath)
	return p
}

// CertificateConfigPublicKeyToProto converts a CertificateConfigPublicKey object to its proto representation.
func PrivatecaAlphaCertificateConfigPublicKeyToProto(o *alpha.CertificateConfigPublicKey) *alphapb.PrivatecaAlphaCertificateConfigPublicKey {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCertificateConfigPublicKey{}
	p.SetKey(dcl.ValueOrEmptyString(o.Key))
	p.SetFormat(PrivatecaAlphaCertificateConfigPublicKeyFormatEnumToProto(o.Format))
	return p
}

// CertificateRevocationDetailsToProto converts a CertificateRevocationDetails object to its proto representation.
func PrivatecaAlphaCertificateRevocationDetailsToProto(o *alpha.CertificateRevocationDetails) *alphapb.PrivatecaAlphaCertificateRevocationDetails {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCertificateRevocationDetails{}
	p.SetRevocationState(PrivatecaAlphaCertificateRevocationDetailsRevocationStateEnumToProto(o.RevocationState))
	p.SetRevocationTime(dcl.ValueOrEmptyString(o.RevocationTime))
	return p
}

// CertificateCertificateDescriptionToProto converts a CertificateCertificateDescription object to its proto representation.
func PrivatecaAlphaCertificateCertificateDescriptionToProto(o *alpha.CertificateCertificateDescription) *alphapb.PrivatecaAlphaCertificateCertificateDescription {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCertificateCertificateDescription{}
	p.SetSubjectDescription(PrivatecaAlphaCertificateCertificateDescriptionSubjectDescriptionToProto(o.SubjectDescription))
	p.SetX509Description(PrivatecaAlphaCertificateCertificateDescriptionX509DescriptionToProto(o.X509Description))
	p.SetPublicKey(PrivatecaAlphaCertificateCertificateDescriptionPublicKeyToProto(o.PublicKey))
	p.SetSubjectKeyId(PrivatecaAlphaCertificateCertificateDescriptionSubjectKeyIdToProto(o.SubjectKeyId))
	p.SetAuthorityKeyId(PrivatecaAlphaCertificateCertificateDescriptionAuthorityKeyIdToProto(o.AuthorityKeyId))
	p.SetCertFingerprint(PrivatecaAlphaCertificateCertificateDescriptionCertFingerprintToProto(o.CertFingerprint))
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

// CertificateCertificateDescriptionSubjectDescriptionToProto converts a CertificateCertificateDescriptionSubjectDescription object to its proto representation.
func PrivatecaAlphaCertificateCertificateDescriptionSubjectDescriptionToProto(o *alpha.CertificateCertificateDescriptionSubjectDescription) *alphapb.PrivatecaAlphaCertificateCertificateDescriptionSubjectDescription {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCertificateCertificateDescriptionSubjectDescription{}
	p.SetSubject(PrivatecaAlphaCertificateCertificateDescriptionSubjectDescriptionSubjectToProto(o.Subject))
	p.SetSubjectAltName(PrivatecaAlphaCertificateCertificateDescriptionSubjectDescriptionSubjectAltNameToProto(o.SubjectAltName))
	p.SetHexSerialNumber(dcl.ValueOrEmptyString(o.HexSerialNumber))
	p.SetLifetime(dcl.ValueOrEmptyString(o.Lifetime))
	p.SetNotBeforeTime(dcl.ValueOrEmptyString(o.NotBeforeTime))
	p.SetNotAfterTime(dcl.ValueOrEmptyString(o.NotAfterTime))
	return p
}

// CertificateCertificateDescriptionSubjectDescriptionSubjectToProto converts a CertificateCertificateDescriptionSubjectDescriptionSubject object to its proto representation.
func PrivatecaAlphaCertificateCertificateDescriptionSubjectDescriptionSubjectToProto(o *alpha.CertificateCertificateDescriptionSubjectDescriptionSubject) *alphapb.PrivatecaAlphaCertificateCertificateDescriptionSubjectDescriptionSubject {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCertificateCertificateDescriptionSubjectDescriptionSubject{}
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

// CertificateCertificateDescriptionSubjectDescriptionSubjectAltNameToProto converts a CertificateCertificateDescriptionSubjectDescriptionSubjectAltName object to its proto representation.
func PrivatecaAlphaCertificateCertificateDescriptionSubjectDescriptionSubjectAltNameToProto(o *alpha.CertificateCertificateDescriptionSubjectDescriptionSubjectAltName) *alphapb.PrivatecaAlphaCertificateCertificateDescriptionSubjectDescriptionSubjectAltName {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCertificateCertificateDescriptionSubjectDescriptionSubjectAltName{}
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
	sCustomSans := make([]*alphapb.PrivatecaAlphaCertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSans, len(o.CustomSans))
	for i, r := range o.CustomSans {
		sCustomSans[i] = PrivatecaAlphaCertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansToProto(&r)
	}
	p.SetCustomSans(sCustomSans)
	return p
}

// CertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansToProto converts a CertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSans object to its proto representation.
func PrivatecaAlphaCertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansToProto(o *alpha.CertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSans) *alphapb.PrivatecaAlphaCertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSans {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSans{}
	p.SetObjectId(PrivatecaAlphaCertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansObjectIdToProto(o.ObjectId))
	p.SetCritical(dcl.ValueOrEmptyBool(o.Critical))
	p.SetValue(dcl.ValueOrEmptyString(o.Value))
	return p
}

// CertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansObjectIdToProto converts a CertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansObjectId object to its proto representation.
func PrivatecaAlphaCertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansObjectIdToProto(o *alpha.CertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansObjectId) *alphapb.PrivatecaAlphaCertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansObjectId {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansObjectId{}
	sObjectIdPath := make([]int64, len(o.ObjectIdPath))
	for i, r := range o.ObjectIdPath {
		sObjectIdPath[i] = r
	}
	p.SetObjectIdPath(sObjectIdPath)
	return p
}

// CertificateCertificateDescriptionX509DescriptionToProto converts a CertificateCertificateDescriptionX509Description object to its proto representation.
func PrivatecaAlphaCertificateCertificateDescriptionX509DescriptionToProto(o *alpha.CertificateCertificateDescriptionX509Description) *alphapb.PrivatecaAlphaCertificateCertificateDescriptionX509Description {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCertificateCertificateDescriptionX509Description{}
	p.SetKeyUsage(PrivatecaAlphaCertificateCertificateDescriptionX509DescriptionKeyUsageToProto(o.KeyUsage))
	p.SetCaOptions(PrivatecaAlphaCertificateCertificateDescriptionX509DescriptionCaOptionsToProto(o.CaOptions))
	sPolicyIds := make([]*alphapb.PrivatecaAlphaCertificateCertificateDescriptionX509DescriptionPolicyIds, len(o.PolicyIds))
	for i, r := range o.PolicyIds {
		sPolicyIds[i] = PrivatecaAlphaCertificateCertificateDescriptionX509DescriptionPolicyIdsToProto(&r)
	}
	p.SetPolicyIds(sPolicyIds)
	sAiaOcspServers := make([]string, len(o.AiaOcspServers))
	for i, r := range o.AiaOcspServers {
		sAiaOcspServers[i] = r
	}
	p.SetAiaOcspServers(sAiaOcspServers)
	sAdditionalExtensions := make([]*alphapb.PrivatecaAlphaCertificateCertificateDescriptionX509DescriptionAdditionalExtensions, len(o.AdditionalExtensions))
	for i, r := range o.AdditionalExtensions {
		sAdditionalExtensions[i] = PrivatecaAlphaCertificateCertificateDescriptionX509DescriptionAdditionalExtensionsToProto(&r)
	}
	p.SetAdditionalExtensions(sAdditionalExtensions)
	return p
}

// CertificateCertificateDescriptionX509DescriptionKeyUsageToProto converts a CertificateCertificateDescriptionX509DescriptionKeyUsage object to its proto representation.
func PrivatecaAlphaCertificateCertificateDescriptionX509DescriptionKeyUsageToProto(o *alpha.CertificateCertificateDescriptionX509DescriptionKeyUsage) *alphapb.PrivatecaAlphaCertificateCertificateDescriptionX509DescriptionKeyUsage {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCertificateCertificateDescriptionX509DescriptionKeyUsage{}
	p.SetBaseKeyUsage(PrivatecaAlphaCertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsageToProto(o.BaseKeyUsage))
	p.SetExtendedKeyUsage(PrivatecaAlphaCertificateCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsageToProto(o.ExtendedKeyUsage))
	sUnknownExtendedKeyUsages := make([]*alphapb.PrivatecaAlphaCertificateCertificateDescriptionX509DescriptionKeyUsageUnknownExtendedKeyUsages, len(o.UnknownExtendedKeyUsages))
	for i, r := range o.UnknownExtendedKeyUsages {
		sUnknownExtendedKeyUsages[i] = PrivatecaAlphaCertificateCertificateDescriptionX509DescriptionKeyUsageUnknownExtendedKeyUsagesToProto(&r)
	}
	p.SetUnknownExtendedKeyUsages(sUnknownExtendedKeyUsages)
	return p
}

// CertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsageToProto converts a CertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsage object to its proto representation.
func PrivatecaAlphaCertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsageToProto(o *alpha.CertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsage) *alphapb.PrivatecaAlphaCertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsage {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsage{}
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

// CertificateCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsageToProto converts a CertificateCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsage object to its proto representation.
func PrivatecaAlphaCertificateCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsageToProto(o *alpha.CertificateCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsage) *alphapb.PrivatecaAlphaCertificateCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsage {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCertificateCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsage{}
	p.SetServerAuth(dcl.ValueOrEmptyBool(o.ServerAuth))
	p.SetClientAuth(dcl.ValueOrEmptyBool(o.ClientAuth))
	p.SetCodeSigning(dcl.ValueOrEmptyBool(o.CodeSigning))
	p.SetEmailProtection(dcl.ValueOrEmptyBool(o.EmailProtection))
	p.SetTimeStamping(dcl.ValueOrEmptyBool(o.TimeStamping))
	p.SetOcspSigning(dcl.ValueOrEmptyBool(o.OcspSigning))
	return p
}

// CertificateCertificateDescriptionX509DescriptionKeyUsageUnknownExtendedKeyUsagesToProto converts a CertificateCertificateDescriptionX509DescriptionKeyUsageUnknownExtendedKeyUsages object to its proto representation.
func PrivatecaAlphaCertificateCertificateDescriptionX509DescriptionKeyUsageUnknownExtendedKeyUsagesToProto(o *alpha.CertificateCertificateDescriptionX509DescriptionKeyUsageUnknownExtendedKeyUsages) *alphapb.PrivatecaAlphaCertificateCertificateDescriptionX509DescriptionKeyUsageUnknownExtendedKeyUsages {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCertificateCertificateDescriptionX509DescriptionKeyUsageUnknownExtendedKeyUsages{}
	sObjectIdPath := make([]int64, len(o.ObjectIdPath))
	for i, r := range o.ObjectIdPath {
		sObjectIdPath[i] = r
	}
	p.SetObjectIdPath(sObjectIdPath)
	return p
}

// CertificateCertificateDescriptionX509DescriptionCaOptionsToProto converts a CertificateCertificateDescriptionX509DescriptionCaOptions object to its proto representation.
func PrivatecaAlphaCertificateCertificateDescriptionX509DescriptionCaOptionsToProto(o *alpha.CertificateCertificateDescriptionX509DescriptionCaOptions) *alphapb.PrivatecaAlphaCertificateCertificateDescriptionX509DescriptionCaOptions {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCertificateCertificateDescriptionX509DescriptionCaOptions{}
	p.SetIsCa(dcl.ValueOrEmptyBool(o.IsCa))
	p.SetMaxIssuerPathLength(dcl.ValueOrEmptyInt64(o.MaxIssuerPathLength))
	return p
}

// CertificateCertificateDescriptionX509DescriptionPolicyIdsToProto converts a CertificateCertificateDescriptionX509DescriptionPolicyIds object to its proto representation.
func PrivatecaAlphaCertificateCertificateDescriptionX509DescriptionPolicyIdsToProto(o *alpha.CertificateCertificateDescriptionX509DescriptionPolicyIds) *alphapb.PrivatecaAlphaCertificateCertificateDescriptionX509DescriptionPolicyIds {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCertificateCertificateDescriptionX509DescriptionPolicyIds{}
	sObjectIdPath := make([]int64, len(o.ObjectIdPath))
	for i, r := range o.ObjectIdPath {
		sObjectIdPath[i] = r
	}
	p.SetObjectIdPath(sObjectIdPath)
	return p
}

// CertificateCertificateDescriptionX509DescriptionAdditionalExtensionsToProto converts a CertificateCertificateDescriptionX509DescriptionAdditionalExtensions object to its proto representation.
func PrivatecaAlphaCertificateCertificateDescriptionX509DescriptionAdditionalExtensionsToProto(o *alpha.CertificateCertificateDescriptionX509DescriptionAdditionalExtensions) *alphapb.PrivatecaAlphaCertificateCertificateDescriptionX509DescriptionAdditionalExtensions {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCertificateCertificateDescriptionX509DescriptionAdditionalExtensions{}
	p.SetObjectId(PrivatecaAlphaCertificateCertificateDescriptionX509DescriptionAdditionalExtensionsObjectIdToProto(o.ObjectId))
	p.SetCritical(dcl.ValueOrEmptyBool(o.Critical))
	p.SetValue(dcl.ValueOrEmptyString(o.Value))
	return p
}

// CertificateCertificateDescriptionX509DescriptionAdditionalExtensionsObjectIdToProto converts a CertificateCertificateDescriptionX509DescriptionAdditionalExtensionsObjectId object to its proto representation.
func PrivatecaAlphaCertificateCertificateDescriptionX509DescriptionAdditionalExtensionsObjectIdToProto(o *alpha.CertificateCertificateDescriptionX509DescriptionAdditionalExtensionsObjectId) *alphapb.PrivatecaAlphaCertificateCertificateDescriptionX509DescriptionAdditionalExtensionsObjectId {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCertificateCertificateDescriptionX509DescriptionAdditionalExtensionsObjectId{}
	sObjectIdPath := make([]int64, len(o.ObjectIdPath))
	for i, r := range o.ObjectIdPath {
		sObjectIdPath[i] = r
	}
	p.SetObjectIdPath(sObjectIdPath)
	return p
}

// CertificateCertificateDescriptionPublicKeyToProto converts a CertificateCertificateDescriptionPublicKey object to its proto representation.
func PrivatecaAlphaCertificateCertificateDescriptionPublicKeyToProto(o *alpha.CertificateCertificateDescriptionPublicKey) *alphapb.PrivatecaAlphaCertificateCertificateDescriptionPublicKey {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCertificateCertificateDescriptionPublicKey{}
	p.SetKey(dcl.ValueOrEmptyString(o.Key))
	p.SetFormat(PrivatecaAlphaCertificateCertificateDescriptionPublicKeyFormatEnumToProto(o.Format))
	return p
}

// CertificateCertificateDescriptionSubjectKeyIdToProto converts a CertificateCertificateDescriptionSubjectKeyId object to its proto representation.
func PrivatecaAlphaCertificateCertificateDescriptionSubjectKeyIdToProto(o *alpha.CertificateCertificateDescriptionSubjectKeyId) *alphapb.PrivatecaAlphaCertificateCertificateDescriptionSubjectKeyId {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCertificateCertificateDescriptionSubjectKeyId{}
	p.SetKeyId(dcl.ValueOrEmptyString(o.KeyId))
	return p
}

// CertificateCertificateDescriptionAuthorityKeyIdToProto converts a CertificateCertificateDescriptionAuthorityKeyId object to its proto representation.
func PrivatecaAlphaCertificateCertificateDescriptionAuthorityKeyIdToProto(o *alpha.CertificateCertificateDescriptionAuthorityKeyId) *alphapb.PrivatecaAlphaCertificateCertificateDescriptionAuthorityKeyId {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCertificateCertificateDescriptionAuthorityKeyId{}
	p.SetKeyId(dcl.ValueOrEmptyString(o.KeyId))
	return p
}

// CertificateCertificateDescriptionCertFingerprintToProto converts a CertificateCertificateDescriptionCertFingerprint object to its proto representation.
func PrivatecaAlphaCertificateCertificateDescriptionCertFingerprintToProto(o *alpha.CertificateCertificateDescriptionCertFingerprint) *alphapb.PrivatecaAlphaCertificateCertificateDescriptionCertFingerprint {
	if o == nil {
		return nil
	}
	p := &alphapb.PrivatecaAlphaCertificateCertificateDescriptionCertFingerprint{}
	p.SetSha256Hash(dcl.ValueOrEmptyString(o.Sha256Hash))
	return p
}

// CertificateToProto converts a Certificate resource to its proto representation.
func CertificateToProto(resource *alpha.Certificate) *alphapb.PrivatecaAlphaCertificate {
	p := &alphapb.PrivatecaAlphaCertificate{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetPemCsr(dcl.ValueOrEmptyString(resource.PemCsr))
	p.SetConfig(PrivatecaAlphaCertificateConfigToProto(resource.Config))
	p.SetIssuerCertificateAuthority(dcl.ValueOrEmptyString(resource.IssuerCertificateAuthority))
	p.SetLifetime(dcl.ValueOrEmptyString(resource.Lifetime))
	p.SetCertificateTemplate(dcl.ValueOrEmptyString(resource.CertificateTemplate))
	p.SetSubjectMode(PrivatecaAlphaCertificateSubjectModeEnumToProto(resource.SubjectMode))
	p.SetRevocationDetails(PrivatecaAlphaCertificateRevocationDetailsToProto(resource.RevocationDetails))
	p.SetPemCertificate(dcl.ValueOrEmptyString(resource.PemCertificate))
	p.SetCertificateDescription(PrivatecaAlphaCertificateCertificateDescriptionToProto(resource.CertificateDescription))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	p.SetCaPool(dcl.ValueOrEmptyString(resource.CaPool))
	p.SetCertificateAuthority(dcl.ValueOrEmptyString(resource.CertificateAuthority))
	sPemCertificateChain := make([]string, len(resource.PemCertificateChain))
	for i, r := range resource.PemCertificateChain {
		sPemCertificateChain[i] = r
	}
	p.SetPemCertificateChain(sPemCertificateChain)
	mLabels := make(map[string]string, len(resource.Labels))
	for k, r := range resource.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)

	return p
}

// applyCertificate handles the gRPC request by passing it to the underlying Certificate Apply() method.
func (s *CertificateServer) applyCertificate(ctx context.Context, c *alpha.Client, request *alphapb.ApplyPrivatecaAlphaCertificateRequest) (*alphapb.PrivatecaAlphaCertificate, error) {
	p := ProtoToCertificate(request.GetResource())
	res, err := c.ApplyCertificate(ctx, p)
	if err != nil {
		return nil, err
	}
	r := CertificateToProto(res)
	return r, nil
}

// applyPrivatecaAlphaCertificate handles the gRPC request by passing it to the underlying Certificate Apply() method.
func (s *CertificateServer) ApplyPrivatecaAlphaCertificate(ctx context.Context, request *alphapb.ApplyPrivatecaAlphaCertificateRequest) (*alphapb.PrivatecaAlphaCertificate, error) {
	cl, err := createConfigCertificate(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyCertificate(ctx, cl, request)
}

// DeleteCertificate handles the gRPC request by passing it to the underlying Certificate Delete() method.
func (s *CertificateServer) DeletePrivatecaAlphaCertificate(ctx context.Context, request *alphapb.DeletePrivatecaAlphaCertificateRequest) (*emptypb.Empty, error) {

	cl, err := createConfigCertificate(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteCertificate(ctx, ProtoToCertificate(request.GetResource()))

}

// ListPrivatecaAlphaCertificate handles the gRPC request by passing it to the underlying CertificateList() method.
func (s *CertificateServer) ListPrivatecaAlphaCertificate(ctx context.Context, request *alphapb.ListPrivatecaAlphaCertificateRequest) (*alphapb.ListPrivatecaAlphaCertificateResponse, error) {
	cl, err := createConfigCertificate(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListCertificate(ctx, request.GetProject(), request.GetLocation(), request.GetCaPool())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.PrivatecaAlphaCertificate
	for _, r := range resources.Items {
		rp := CertificateToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListPrivatecaAlphaCertificateResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigCertificate(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
