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

// CertificateServer implements the gRPC interface for Certificate.
type CertificateServer struct{}

// ProtoToCertificateConfigPublicKeyFormatEnum converts a CertificateConfigPublicKeyFormatEnum enum from its proto representation.
func ProtoToPrivatecaBetaCertificateConfigPublicKeyFormatEnum(e betapb.PrivatecaBetaCertificateConfigPublicKeyFormatEnum) *beta.CertificateConfigPublicKeyFormatEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.PrivatecaBetaCertificateConfigPublicKeyFormatEnum_name[int32(e)]; ok {
		e := beta.CertificateConfigPublicKeyFormatEnum(n[len("PrivatecaBetaCertificateConfigPublicKeyFormatEnum"):])
		return &e
	}
	return nil
}

// ProtoToCertificateSubjectModeEnum converts a CertificateSubjectModeEnum enum from its proto representation.
func ProtoToPrivatecaBetaCertificateSubjectModeEnum(e betapb.PrivatecaBetaCertificateSubjectModeEnum) *beta.CertificateSubjectModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.PrivatecaBetaCertificateSubjectModeEnum_name[int32(e)]; ok {
		e := beta.CertificateSubjectModeEnum(n[len("PrivatecaBetaCertificateSubjectModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToCertificateRevocationDetailsRevocationStateEnum converts a CertificateRevocationDetailsRevocationStateEnum enum from its proto representation.
func ProtoToPrivatecaBetaCertificateRevocationDetailsRevocationStateEnum(e betapb.PrivatecaBetaCertificateRevocationDetailsRevocationStateEnum) *beta.CertificateRevocationDetailsRevocationStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.PrivatecaBetaCertificateRevocationDetailsRevocationStateEnum_name[int32(e)]; ok {
		e := beta.CertificateRevocationDetailsRevocationStateEnum(n[len("PrivatecaBetaCertificateRevocationDetailsRevocationStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToCertificateCertificateDescriptionPublicKeyFormatEnum converts a CertificateCertificateDescriptionPublicKeyFormatEnum enum from its proto representation.
func ProtoToPrivatecaBetaCertificateCertificateDescriptionPublicKeyFormatEnum(e betapb.PrivatecaBetaCertificateCertificateDescriptionPublicKeyFormatEnum) *beta.CertificateCertificateDescriptionPublicKeyFormatEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.PrivatecaBetaCertificateCertificateDescriptionPublicKeyFormatEnum_name[int32(e)]; ok {
		e := beta.CertificateCertificateDescriptionPublicKeyFormatEnum(n[len("PrivatecaBetaCertificateCertificateDescriptionPublicKeyFormatEnum"):])
		return &e
	}
	return nil
}

// ProtoToCertificateConfig converts a CertificateConfig object from its proto representation.
func ProtoToPrivatecaBetaCertificateConfig(p *betapb.PrivatecaBetaCertificateConfig) *beta.CertificateConfig {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateConfig{
		SubjectConfig: ProtoToPrivatecaBetaCertificateConfigSubjectConfig(p.GetSubjectConfig()),
		X509Config:    ProtoToPrivatecaBetaCertificateConfigX509Config(p.GetX509Config()),
		PublicKey:     ProtoToPrivatecaBetaCertificateConfigPublicKey(p.GetPublicKey()),
	}
	return obj
}

// ProtoToCertificateConfigSubjectConfig converts a CertificateConfigSubjectConfig object from its proto representation.
func ProtoToPrivatecaBetaCertificateConfigSubjectConfig(p *betapb.PrivatecaBetaCertificateConfigSubjectConfig) *beta.CertificateConfigSubjectConfig {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateConfigSubjectConfig{
		Subject:        ProtoToPrivatecaBetaCertificateConfigSubjectConfigSubject(p.GetSubject()),
		SubjectAltName: ProtoToPrivatecaBetaCertificateConfigSubjectConfigSubjectAltName(p.GetSubjectAltName()),
	}
	return obj
}

// ProtoToCertificateConfigSubjectConfigSubject converts a CertificateConfigSubjectConfigSubject object from its proto representation.
func ProtoToPrivatecaBetaCertificateConfigSubjectConfigSubject(p *betapb.PrivatecaBetaCertificateConfigSubjectConfigSubject) *beta.CertificateConfigSubjectConfigSubject {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateConfigSubjectConfigSubject{
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
func ProtoToPrivatecaBetaCertificateConfigSubjectConfigSubjectAltName(p *betapb.PrivatecaBetaCertificateConfigSubjectConfigSubjectAltName) *beta.CertificateConfigSubjectConfigSubjectAltName {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateConfigSubjectConfigSubjectAltName{}
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
func ProtoToPrivatecaBetaCertificateConfigX509Config(p *betapb.PrivatecaBetaCertificateConfigX509Config) *beta.CertificateConfigX509Config {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateConfigX509Config{
		KeyUsage:  ProtoToPrivatecaBetaCertificateConfigX509ConfigKeyUsage(p.GetKeyUsage()),
		CaOptions: ProtoToPrivatecaBetaCertificateConfigX509ConfigCaOptions(p.GetCaOptions()),
	}
	for _, r := range p.GetPolicyIds() {
		obj.PolicyIds = append(obj.PolicyIds, *ProtoToPrivatecaBetaCertificateConfigX509ConfigPolicyIds(r))
	}
	for _, r := range p.GetAiaOcspServers() {
		obj.AiaOcspServers = append(obj.AiaOcspServers, r)
	}
	for _, r := range p.GetAdditionalExtensions() {
		obj.AdditionalExtensions = append(obj.AdditionalExtensions, *ProtoToPrivatecaBetaCertificateConfigX509ConfigAdditionalExtensions(r))
	}
	return obj
}

// ProtoToCertificateConfigX509ConfigKeyUsage converts a CertificateConfigX509ConfigKeyUsage object from its proto representation.
func ProtoToPrivatecaBetaCertificateConfigX509ConfigKeyUsage(p *betapb.PrivatecaBetaCertificateConfigX509ConfigKeyUsage) *beta.CertificateConfigX509ConfigKeyUsage {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateConfigX509ConfigKeyUsage{
		BaseKeyUsage:     ProtoToPrivatecaBetaCertificateConfigX509ConfigKeyUsageBaseKeyUsage(p.GetBaseKeyUsage()),
		ExtendedKeyUsage: ProtoToPrivatecaBetaCertificateConfigX509ConfigKeyUsageExtendedKeyUsage(p.GetExtendedKeyUsage()),
	}
	for _, r := range p.GetUnknownExtendedKeyUsages() {
		obj.UnknownExtendedKeyUsages = append(obj.UnknownExtendedKeyUsages, *ProtoToPrivatecaBetaCertificateConfigX509ConfigKeyUsageUnknownExtendedKeyUsages(r))
	}
	return obj
}

// ProtoToCertificateConfigX509ConfigKeyUsageBaseKeyUsage converts a CertificateConfigX509ConfigKeyUsageBaseKeyUsage object from its proto representation.
func ProtoToPrivatecaBetaCertificateConfigX509ConfigKeyUsageBaseKeyUsage(p *betapb.PrivatecaBetaCertificateConfigX509ConfigKeyUsageBaseKeyUsage) *beta.CertificateConfigX509ConfigKeyUsageBaseKeyUsage {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateConfigX509ConfigKeyUsageBaseKeyUsage{
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
func ProtoToPrivatecaBetaCertificateConfigX509ConfigKeyUsageExtendedKeyUsage(p *betapb.PrivatecaBetaCertificateConfigX509ConfigKeyUsageExtendedKeyUsage) *beta.CertificateConfigX509ConfigKeyUsageExtendedKeyUsage {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateConfigX509ConfigKeyUsageExtendedKeyUsage{
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
func ProtoToPrivatecaBetaCertificateConfigX509ConfigKeyUsageUnknownExtendedKeyUsages(p *betapb.PrivatecaBetaCertificateConfigX509ConfigKeyUsageUnknownExtendedKeyUsages) *beta.CertificateConfigX509ConfigKeyUsageUnknownExtendedKeyUsages {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateConfigX509ConfigKeyUsageUnknownExtendedKeyUsages{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCertificateConfigX509ConfigCaOptions converts a CertificateConfigX509ConfigCaOptions object from its proto representation.
func ProtoToPrivatecaBetaCertificateConfigX509ConfigCaOptions(p *betapb.PrivatecaBetaCertificateConfigX509ConfigCaOptions) *beta.CertificateConfigX509ConfigCaOptions {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateConfigX509ConfigCaOptions{
		IsCa:                    dcl.Bool(p.GetIsCa()),
		NonCa:                   dcl.Bool(p.GetNonCa()),
		MaxIssuerPathLength:     dcl.Int64OrNil(p.GetMaxIssuerPathLength()),
		ZeroMaxIssuerPathLength: dcl.Bool(p.GetZeroMaxIssuerPathLength()),
	}
	return obj
}

// ProtoToCertificateConfigX509ConfigPolicyIds converts a CertificateConfigX509ConfigPolicyIds object from its proto representation.
func ProtoToPrivatecaBetaCertificateConfigX509ConfigPolicyIds(p *betapb.PrivatecaBetaCertificateConfigX509ConfigPolicyIds) *beta.CertificateConfigX509ConfigPolicyIds {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateConfigX509ConfigPolicyIds{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCertificateConfigX509ConfigAdditionalExtensions converts a CertificateConfigX509ConfigAdditionalExtensions object from its proto representation.
func ProtoToPrivatecaBetaCertificateConfigX509ConfigAdditionalExtensions(p *betapb.PrivatecaBetaCertificateConfigX509ConfigAdditionalExtensions) *beta.CertificateConfigX509ConfigAdditionalExtensions {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateConfigX509ConfigAdditionalExtensions{
		ObjectId: ProtoToPrivatecaBetaCertificateConfigX509ConfigAdditionalExtensionsObjectId(p.GetObjectId()),
		Critical: dcl.Bool(p.GetCritical()),
		Value:    dcl.StringOrNil(p.GetValue()),
	}
	return obj
}

// ProtoToCertificateConfigX509ConfigAdditionalExtensionsObjectId converts a CertificateConfigX509ConfigAdditionalExtensionsObjectId object from its proto representation.
func ProtoToPrivatecaBetaCertificateConfigX509ConfigAdditionalExtensionsObjectId(p *betapb.PrivatecaBetaCertificateConfigX509ConfigAdditionalExtensionsObjectId) *beta.CertificateConfigX509ConfigAdditionalExtensionsObjectId {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateConfigX509ConfigAdditionalExtensionsObjectId{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCertificateConfigPublicKey converts a CertificateConfigPublicKey object from its proto representation.
func ProtoToPrivatecaBetaCertificateConfigPublicKey(p *betapb.PrivatecaBetaCertificateConfigPublicKey) *beta.CertificateConfigPublicKey {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateConfigPublicKey{
		Key:    dcl.StringOrNil(p.GetKey()),
		Format: ProtoToPrivatecaBetaCertificateConfigPublicKeyFormatEnum(p.GetFormat()),
	}
	return obj
}

// ProtoToCertificateRevocationDetails converts a CertificateRevocationDetails object from its proto representation.
func ProtoToPrivatecaBetaCertificateRevocationDetails(p *betapb.PrivatecaBetaCertificateRevocationDetails) *beta.CertificateRevocationDetails {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateRevocationDetails{
		RevocationState: ProtoToPrivatecaBetaCertificateRevocationDetailsRevocationStateEnum(p.GetRevocationState()),
		RevocationTime:  dcl.StringOrNil(p.GetRevocationTime()),
	}
	return obj
}

// ProtoToCertificateCertificateDescription converts a CertificateCertificateDescription object from its proto representation.
func ProtoToPrivatecaBetaCertificateCertificateDescription(p *betapb.PrivatecaBetaCertificateCertificateDescription) *beta.CertificateCertificateDescription {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateCertificateDescription{
		SubjectDescription: ProtoToPrivatecaBetaCertificateCertificateDescriptionSubjectDescription(p.GetSubjectDescription()),
		X509Description:    ProtoToPrivatecaBetaCertificateCertificateDescriptionX509Description(p.GetX509Description()),
		PublicKey:          ProtoToPrivatecaBetaCertificateCertificateDescriptionPublicKey(p.GetPublicKey()),
		SubjectKeyId:       ProtoToPrivatecaBetaCertificateCertificateDescriptionSubjectKeyId(p.GetSubjectKeyId()),
		AuthorityKeyId:     ProtoToPrivatecaBetaCertificateCertificateDescriptionAuthorityKeyId(p.GetAuthorityKeyId()),
		CertFingerprint:    ProtoToPrivatecaBetaCertificateCertificateDescriptionCertFingerprint(p.GetCertFingerprint()),
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
func ProtoToPrivatecaBetaCertificateCertificateDescriptionSubjectDescription(p *betapb.PrivatecaBetaCertificateCertificateDescriptionSubjectDescription) *beta.CertificateCertificateDescriptionSubjectDescription {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateCertificateDescriptionSubjectDescription{
		Subject:         ProtoToPrivatecaBetaCertificateCertificateDescriptionSubjectDescriptionSubject(p.GetSubject()),
		SubjectAltName:  ProtoToPrivatecaBetaCertificateCertificateDescriptionSubjectDescriptionSubjectAltName(p.GetSubjectAltName()),
		HexSerialNumber: dcl.StringOrNil(p.GetHexSerialNumber()),
		Lifetime:        dcl.StringOrNil(p.GetLifetime()),
		NotBeforeTime:   dcl.StringOrNil(p.GetNotBeforeTime()),
		NotAfterTime:    dcl.StringOrNil(p.GetNotAfterTime()),
	}
	return obj
}

// ProtoToCertificateCertificateDescriptionSubjectDescriptionSubject converts a CertificateCertificateDescriptionSubjectDescriptionSubject object from its proto representation.
func ProtoToPrivatecaBetaCertificateCertificateDescriptionSubjectDescriptionSubject(p *betapb.PrivatecaBetaCertificateCertificateDescriptionSubjectDescriptionSubject) *beta.CertificateCertificateDescriptionSubjectDescriptionSubject {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateCertificateDescriptionSubjectDescriptionSubject{
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
func ProtoToPrivatecaBetaCertificateCertificateDescriptionSubjectDescriptionSubjectAltName(p *betapb.PrivatecaBetaCertificateCertificateDescriptionSubjectDescriptionSubjectAltName) *beta.CertificateCertificateDescriptionSubjectDescriptionSubjectAltName {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateCertificateDescriptionSubjectDescriptionSubjectAltName{}
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
		obj.CustomSans = append(obj.CustomSans, *ProtoToPrivatecaBetaCertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSans(r))
	}
	return obj
}

// ProtoToCertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSans converts a CertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSans object from its proto representation.
func ProtoToPrivatecaBetaCertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSans(p *betapb.PrivatecaBetaCertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSans) *beta.CertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSans {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSans{
		ObjectId: ProtoToPrivatecaBetaCertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansObjectId(p.GetObjectId()),
		Critical: dcl.Bool(p.GetCritical()),
		Value:    dcl.StringOrNil(p.GetValue()),
	}
	return obj
}

// ProtoToCertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansObjectId converts a CertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansObjectId object from its proto representation.
func ProtoToPrivatecaBetaCertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansObjectId(p *betapb.PrivatecaBetaCertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansObjectId) *beta.CertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansObjectId {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansObjectId{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCertificateCertificateDescriptionX509Description converts a CertificateCertificateDescriptionX509Description object from its proto representation.
func ProtoToPrivatecaBetaCertificateCertificateDescriptionX509Description(p *betapb.PrivatecaBetaCertificateCertificateDescriptionX509Description) *beta.CertificateCertificateDescriptionX509Description {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateCertificateDescriptionX509Description{
		KeyUsage:  ProtoToPrivatecaBetaCertificateCertificateDescriptionX509DescriptionKeyUsage(p.GetKeyUsage()),
		CaOptions: ProtoToPrivatecaBetaCertificateCertificateDescriptionX509DescriptionCaOptions(p.GetCaOptions()),
	}
	for _, r := range p.GetPolicyIds() {
		obj.PolicyIds = append(obj.PolicyIds, *ProtoToPrivatecaBetaCertificateCertificateDescriptionX509DescriptionPolicyIds(r))
	}
	for _, r := range p.GetAiaOcspServers() {
		obj.AiaOcspServers = append(obj.AiaOcspServers, r)
	}
	for _, r := range p.GetAdditionalExtensions() {
		obj.AdditionalExtensions = append(obj.AdditionalExtensions, *ProtoToPrivatecaBetaCertificateCertificateDescriptionX509DescriptionAdditionalExtensions(r))
	}
	return obj
}

// ProtoToCertificateCertificateDescriptionX509DescriptionKeyUsage converts a CertificateCertificateDescriptionX509DescriptionKeyUsage object from its proto representation.
func ProtoToPrivatecaBetaCertificateCertificateDescriptionX509DescriptionKeyUsage(p *betapb.PrivatecaBetaCertificateCertificateDescriptionX509DescriptionKeyUsage) *beta.CertificateCertificateDescriptionX509DescriptionKeyUsage {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateCertificateDescriptionX509DescriptionKeyUsage{
		BaseKeyUsage:     ProtoToPrivatecaBetaCertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsage(p.GetBaseKeyUsage()),
		ExtendedKeyUsage: ProtoToPrivatecaBetaCertificateCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsage(p.GetExtendedKeyUsage()),
	}
	for _, r := range p.GetUnknownExtendedKeyUsages() {
		obj.UnknownExtendedKeyUsages = append(obj.UnknownExtendedKeyUsages, *ProtoToPrivatecaBetaCertificateCertificateDescriptionX509DescriptionKeyUsageUnknownExtendedKeyUsages(r))
	}
	return obj
}

// ProtoToCertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsage converts a CertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsage object from its proto representation.
func ProtoToPrivatecaBetaCertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsage(p *betapb.PrivatecaBetaCertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsage) *beta.CertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsage {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsage{
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
func ProtoToPrivatecaBetaCertificateCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsage(p *betapb.PrivatecaBetaCertificateCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsage) *beta.CertificateCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsage {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsage{
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
func ProtoToPrivatecaBetaCertificateCertificateDescriptionX509DescriptionKeyUsageUnknownExtendedKeyUsages(p *betapb.PrivatecaBetaCertificateCertificateDescriptionX509DescriptionKeyUsageUnknownExtendedKeyUsages) *beta.CertificateCertificateDescriptionX509DescriptionKeyUsageUnknownExtendedKeyUsages {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateCertificateDescriptionX509DescriptionKeyUsageUnknownExtendedKeyUsages{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCertificateCertificateDescriptionX509DescriptionCaOptions converts a CertificateCertificateDescriptionX509DescriptionCaOptions object from its proto representation.
func ProtoToPrivatecaBetaCertificateCertificateDescriptionX509DescriptionCaOptions(p *betapb.PrivatecaBetaCertificateCertificateDescriptionX509DescriptionCaOptions) *beta.CertificateCertificateDescriptionX509DescriptionCaOptions {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateCertificateDescriptionX509DescriptionCaOptions{
		IsCa:                dcl.Bool(p.GetIsCa()),
		MaxIssuerPathLength: dcl.Int64OrNil(p.GetMaxIssuerPathLength()),
	}
	return obj
}

// ProtoToCertificateCertificateDescriptionX509DescriptionPolicyIds converts a CertificateCertificateDescriptionX509DescriptionPolicyIds object from its proto representation.
func ProtoToPrivatecaBetaCertificateCertificateDescriptionX509DescriptionPolicyIds(p *betapb.PrivatecaBetaCertificateCertificateDescriptionX509DescriptionPolicyIds) *beta.CertificateCertificateDescriptionX509DescriptionPolicyIds {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateCertificateDescriptionX509DescriptionPolicyIds{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCertificateCertificateDescriptionX509DescriptionAdditionalExtensions converts a CertificateCertificateDescriptionX509DescriptionAdditionalExtensions object from its proto representation.
func ProtoToPrivatecaBetaCertificateCertificateDescriptionX509DescriptionAdditionalExtensions(p *betapb.PrivatecaBetaCertificateCertificateDescriptionX509DescriptionAdditionalExtensions) *beta.CertificateCertificateDescriptionX509DescriptionAdditionalExtensions {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateCertificateDescriptionX509DescriptionAdditionalExtensions{
		ObjectId: ProtoToPrivatecaBetaCertificateCertificateDescriptionX509DescriptionAdditionalExtensionsObjectId(p.GetObjectId()),
		Critical: dcl.Bool(p.GetCritical()),
		Value:    dcl.StringOrNil(p.GetValue()),
	}
	return obj
}

// ProtoToCertificateCertificateDescriptionX509DescriptionAdditionalExtensionsObjectId converts a CertificateCertificateDescriptionX509DescriptionAdditionalExtensionsObjectId object from its proto representation.
func ProtoToPrivatecaBetaCertificateCertificateDescriptionX509DescriptionAdditionalExtensionsObjectId(p *betapb.PrivatecaBetaCertificateCertificateDescriptionX509DescriptionAdditionalExtensionsObjectId) *beta.CertificateCertificateDescriptionX509DescriptionAdditionalExtensionsObjectId {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateCertificateDescriptionX509DescriptionAdditionalExtensionsObjectId{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCertificateCertificateDescriptionPublicKey converts a CertificateCertificateDescriptionPublicKey object from its proto representation.
func ProtoToPrivatecaBetaCertificateCertificateDescriptionPublicKey(p *betapb.PrivatecaBetaCertificateCertificateDescriptionPublicKey) *beta.CertificateCertificateDescriptionPublicKey {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateCertificateDescriptionPublicKey{
		Key:    dcl.StringOrNil(p.GetKey()),
		Format: ProtoToPrivatecaBetaCertificateCertificateDescriptionPublicKeyFormatEnum(p.GetFormat()),
	}
	return obj
}

// ProtoToCertificateCertificateDescriptionSubjectKeyId converts a CertificateCertificateDescriptionSubjectKeyId object from its proto representation.
func ProtoToPrivatecaBetaCertificateCertificateDescriptionSubjectKeyId(p *betapb.PrivatecaBetaCertificateCertificateDescriptionSubjectKeyId) *beta.CertificateCertificateDescriptionSubjectKeyId {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateCertificateDescriptionSubjectKeyId{
		KeyId: dcl.StringOrNil(p.GetKeyId()),
	}
	return obj
}

// ProtoToCertificateCertificateDescriptionAuthorityKeyId converts a CertificateCertificateDescriptionAuthorityKeyId object from its proto representation.
func ProtoToPrivatecaBetaCertificateCertificateDescriptionAuthorityKeyId(p *betapb.PrivatecaBetaCertificateCertificateDescriptionAuthorityKeyId) *beta.CertificateCertificateDescriptionAuthorityKeyId {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateCertificateDescriptionAuthorityKeyId{
		KeyId: dcl.StringOrNil(p.GetKeyId()),
	}
	return obj
}

// ProtoToCertificateCertificateDescriptionCertFingerprint converts a CertificateCertificateDescriptionCertFingerprint object from its proto representation.
func ProtoToPrivatecaBetaCertificateCertificateDescriptionCertFingerprint(p *betapb.PrivatecaBetaCertificateCertificateDescriptionCertFingerprint) *beta.CertificateCertificateDescriptionCertFingerprint {
	if p == nil {
		return nil
	}
	obj := &beta.CertificateCertificateDescriptionCertFingerprint{
		Sha256Hash: dcl.StringOrNil(p.GetSha256Hash()),
	}
	return obj
}

// ProtoToCertificate converts a Certificate resource from its proto representation.
func ProtoToCertificate(p *betapb.PrivatecaBetaCertificate) *beta.Certificate {
	obj := &beta.Certificate{
		Name:                       dcl.StringOrNil(p.GetName()),
		PemCsr:                     dcl.StringOrNil(p.GetPemCsr()),
		Config:                     ProtoToPrivatecaBetaCertificateConfig(p.GetConfig()),
		IssuerCertificateAuthority: dcl.StringOrNil(p.GetIssuerCertificateAuthority()),
		Lifetime:                   dcl.StringOrNil(p.GetLifetime()),
		CertificateTemplate:        dcl.StringOrNil(p.GetCertificateTemplate()),
		SubjectMode:                ProtoToPrivatecaBetaCertificateSubjectModeEnum(p.GetSubjectMode()),
		RevocationDetails:          ProtoToPrivatecaBetaCertificateRevocationDetails(p.GetRevocationDetails()),
		PemCertificate:             dcl.StringOrNil(p.GetPemCertificate()),
		CertificateDescription:     ProtoToPrivatecaBetaCertificateCertificateDescription(p.GetCertificateDescription()),
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
func PrivatecaBetaCertificateConfigPublicKeyFormatEnumToProto(e *beta.CertificateConfigPublicKeyFormatEnum) betapb.PrivatecaBetaCertificateConfigPublicKeyFormatEnum {
	if e == nil {
		return betapb.PrivatecaBetaCertificateConfigPublicKeyFormatEnum(0)
	}
	if v, ok := betapb.PrivatecaBetaCertificateConfigPublicKeyFormatEnum_value["CertificateConfigPublicKeyFormatEnum"+string(*e)]; ok {
		return betapb.PrivatecaBetaCertificateConfigPublicKeyFormatEnum(v)
	}
	return betapb.PrivatecaBetaCertificateConfigPublicKeyFormatEnum(0)
}

// CertificateSubjectModeEnumToProto converts a CertificateSubjectModeEnum enum to its proto representation.
func PrivatecaBetaCertificateSubjectModeEnumToProto(e *beta.CertificateSubjectModeEnum) betapb.PrivatecaBetaCertificateSubjectModeEnum {
	if e == nil {
		return betapb.PrivatecaBetaCertificateSubjectModeEnum(0)
	}
	if v, ok := betapb.PrivatecaBetaCertificateSubjectModeEnum_value["CertificateSubjectModeEnum"+string(*e)]; ok {
		return betapb.PrivatecaBetaCertificateSubjectModeEnum(v)
	}
	return betapb.PrivatecaBetaCertificateSubjectModeEnum(0)
}

// CertificateRevocationDetailsRevocationStateEnumToProto converts a CertificateRevocationDetailsRevocationStateEnum enum to its proto representation.
func PrivatecaBetaCertificateRevocationDetailsRevocationStateEnumToProto(e *beta.CertificateRevocationDetailsRevocationStateEnum) betapb.PrivatecaBetaCertificateRevocationDetailsRevocationStateEnum {
	if e == nil {
		return betapb.PrivatecaBetaCertificateRevocationDetailsRevocationStateEnum(0)
	}
	if v, ok := betapb.PrivatecaBetaCertificateRevocationDetailsRevocationStateEnum_value["CertificateRevocationDetailsRevocationStateEnum"+string(*e)]; ok {
		return betapb.PrivatecaBetaCertificateRevocationDetailsRevocationStateEnum(v)
	}
	return betapb.PrivatecaBetaCertificateRevocationDetailsRevocationStateEnum(0)
}

// CertificateCertificateDescriptionPublicKeyFormatEnumToProto converts a CertificateCertificateDescriptionPublicKeyFormatEnum enum to its proto representation.
func PrivatecaBetaCertificateCertificateDescriptionPublicKeyFormatEnumToProto(e *beta.CertificateCertificateDescriptionPublicKeyFormatEnum) betapb.PrivatecaBetaCertificateCertificateDescriptionPublicKeyFormatEnum {
	if e == nil {
		return betapb.PrivatecaBetaCertificateCertificateDescriptionPublicKeyFormatEnum(0)
	}
	if v, ok := betapb.PrivatecaBetaCertificateCertificateDescriptionPublicKeyFormatEnum_value["CertificateCertificateDescriptionPublicKeyFormatEnum"+string(*e)]; ok {
		return betapb.PrivatecaBetaCertificateCertificateDescriptionPublicKeyFormatEnum(v)
	}
	return betapb.PrivatecaBetaCertificateCertificateDescriptionPublicKeyFormatEnum(0)
}

// CertificateConfigToProto converts a CertificateConfig object to its proto representation.
func PrivatecaBetaCertificateConfigToProto(o *beta.CertificateConfig) *betapb.PrivatecaBetaCertificateConfig {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateConfig{}
	p.SetSubjectConfig(PrivatecaBetaCertificateConfigSubjectConfigToProto(o.SubjectConfig))
	p.SetX509Config(PrivatecaBetaCertificateConfigX509ConfigToProto(o.X509Config))
	p.SetPublicKey(PrivatecaBetaCertificateConfigPublicKeyToProto(o.PublicKey))
	return p
}

// CertificateConfigSubjectConfigToProto converts a CertificateConfigSubjectConfig object to its proto representation.
func PrivatecaBetaCertificateConfigSubjectConfigToProto(o *beta.CertificateConfigSubjectConfig) *betapb.PrivatecaBetaCertificateConfigSubjectConfig {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateConfigSubjectConfig{}
	p.SetSubject(PrivatecaBetaCertificateConfigSubjectConfigSubjectToProto(o.Subject))
	p.SetSubjectAltName(PrivatecaBetaCertificateConfigSubjectConfigSubjectAltNameToProto(o.SubjectAltName))
	return p
}

// CertificateConfigSubjectConfigSubjectToProto converts a CertificateConfigSubjectConfigSubject object to its proto representation.
func PrivatecaBetaCertificateConfigSubjectConfigSubjectToProto(o *beta.CertificateConfigSubjectConfigSubject) *betapb.PrivatecaBetaCertificateConfigSubjectConfigSubject {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateConfigSubjectConfigSubject{}
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
func PrivatecaBetaCertificateConfigSubjectConfigSubjectAltNameToProto(o *beta.CertificateConfigSubjectConfigSubjectAltName) *betapb.PrivatecaBetaCertificateConfigSubjectConfigSubjectAltName {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateConfigSubjectConfigSubjectAltName{}
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
func PrivatecaBetaCertificateConfigX509ConfigToProto(o *beta.CertificateConfigX509Config) *betapb.PrivatecaBetaCertificateConfigX509Config {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateConfigX509Config{}
	p.SetKeyUsage(PrivatecaBetaCertificateConfigX509ConfigKeyUsageToProto(o.KeyUsage))
	p.SetCaOptions(PrivatecaBetaCertificateConfigX509ConfigCaOptionsToProto(o.CaOptions))
	sPolicyIds := make([]*betapb.PrivatecaBetaCertificateConfigX509ConfigPolicyIds, len(o.PolicyIds))
	for i, r := range o.PolicyIds {
		sPolicyIds[i] = PrivatecaBetaCertificateConfigX509ConfigPolicyIdsToProto(&r)
	}
	p.SetPolicyIds(sPolicyIds)
	sAiaOcspServers := make([]string, len(o.AiaOcspServers))
	for i, r := range o.AiaOcspServers {
		sAiaOcspServers[i] = r
	}
	p.SetAiaOcspServers(sAiaOcspServers)
	sAdditionalExtensions := make([]*betapb.PrivatecaBetaCertificateConfigX509ConfigAdditionalExtensions, len(o.AdditionalExtensions))
	for i, r := range o.AdditionalExtensions {
		sAdditionalExtensions[i] = PrivatecaBetaCertificateConfigX509ConfigAdditionalExtensionsToProto(&r)
	}
	p.SetAdditionalExtensions(sAdditionalExtensions)
	return p
}

// CertificateConfigX509ConfigKeyUsageToProto converts a CertificateConfigX509ConfigKeyUsage object to its proto representation.
func PrivatecaBetaCertificateConfigX509ConfigKeyUsageToProto(o *beta.CertificateConfigX509ConfigKeyUsage) *betapb.PrivatecaBetaCertificateConfigX509ConfigKeyUsage {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateConfigX509ConfigKeyUsage{}
	p.SetBaseKeyUsage(PrivatecaBetaCertificateConfigX509ConfigKeyUsageBaseKeyUsageToProto(o.BaseKeyUsage))
	p.SetExtendedKeyUsage(PrivatecaBetaCertificateConfigX509ConfigKeyUsageExtendedKeyUsageToProto(o.ExtendedKeyUsage))
	sUnknownExtendedKeyUsages := make([]*betapb.PrivatecaBetaCertificateConfigX509ConfigKeyUsageUnknownExtendedKeyUsages, len(o.UnknownExtendedKeyUsages))
	for i, r := range o.UnknownExtendedKeyUsages {
		sUnknownExtendedKeyUsages[i] = PrivatecaBetaCertificateConfigX509ConfigKeyUsageUnknownExtendedKeyUsagesToProto(&r)
	}
	p.SetUnknownExtendedKeyUsages(sUnknownExtendedKeyUsages)
	return p
}

// CertificateConfigX509ConfigKeyUsageBaseKeyUsageToProto converts a CertificateConfigX509ConfigKeyUsageBaseKeyUsage object to its proto representation.
func PrivatecaBetaCertificateConfigX509ConfigKeyUsageBaseKeyUsageToProto(o *beta.CertificateConfigX509ConfigKeyUsageBaseKeyUsage) *betapb.PrivatecaBetaCertificateConfigX509ConfigKeyUsageBaseKeyUsage {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateConfigX509ConfigKeyUsageBaseKeyUsage{}
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
func PrivatecaBetaCertificateConfigX509ConfigKeyUsageExtendedKeyUsageToProto(o *beta.CertificateConfigX509ConfigKeyUsageExtendedKeyUsage) *betapb.PrivatecaBetaCertificateConfigX509ConfigKeyUsageExtendedKeyUsage {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateConfigX509ConfigKeyUsageExtendedKeyUsage{}
	p.SetServerAuth(dcl.ValueOrEmptyBool(o.ServerAuth))
	p.SetClientAuth(dcl.ValueOrEmptyBool(o.ClientAuth))
	p.SetCodeSigning(dcl.ValueOrEmptyBool(o.CodeSigning))
	p.SetEmailProtection(dcl.ValueOrEmptyBool(o.EmailProtection))
	p.SetTimeStamping(dcl.ValueOrEmptyBool(o.TimeStamping))
	p.SetOcspSigning(dcl.ValueOrEmptyBool(o.OcspSigning))
	return p
}

// CertificateConfigX509ConfigKeyUsageUnknownExtendedKeyUsagesToProto converts a CertificateConfigX509ConfigKeyUsageUnknownExtendedKeyUsages object to its proto representation.
func PrivatecaBetaCertificateConfigX509ConfigKeyUsageUnknownExtendedKeyUsagesToProto(o *beta.CertificateConfigX509ConfigKeyUsageUnknownExtendedKeyUsages) *betapb.PrivatecaBetaCertificateConfigX509ConfigKeyUsageUnknownExtendedKeyUsages {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateConfigX509ConfigKeyUsageUnknownExtendedKeyUsages{}
	sObjectIdPath := make([]int64, len(o.ObjectIdPath))
	for i, r := range o.ObjectIdPath {
		sObjectIdPath[i] = r
	}
	p.SetObjectIdPath(sObjectIdPath)
	return p
}

// CertificateConfigX509ConfigCaOptionsToProto converts a CertificateConfigX509ConfigCaOptions object to its proto representation.
func PrivatecaBetaCertificateConfigX509ConfigCaOptionsToProto(o *beta.CertificateConfigX509ConfigCaOptions) *betapb.PrivatecaBetaCertificateConfigX509ConfigCaOptions {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateConfigX509ConfigCaOptions{}
	p.SetIsCa(dcl.ValueOrEmptyBool(o.IsCa))
	p.SetNonCa(dcl.ValueOrEmptyBool(o.NonCa))
	p.SetMaxIssuerPathLength(dcl.ValueOrEmptyInt64(o.MaxIssuerPathLength))
	p.SetZeroMaxIssuerPathLength(dcl.ValueOrEmptyBool(o.ZeroMaxIssuerPathLength))
	return p
}

// CertificateConfigX509ConfigPolicyIdsToProto converts a CertificateConfigX509ConfigPolicyIds object to its proto representation.
func PrivatecaBetaCertificateConfigX509ConfigPolicyIdsToProto(o *beta.CertificateConfigX509ConfigPolicyIds) *betapb.PrivatecaBetaCertificateConfigX509ConfigPolicyIds {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateConfigX509ConfigPolicyIds{}
	sObjectIdPath := make([]int64, len(o.ObjectIdPath))
	for i, r := range o.ObjectIdPath {
		sObjectIdPath[i] = r
	}
	p.SetObjectIdPath(sObjectIdPath)
	return p
}

// CertificateConfigX509ConfigAdditionalExtensionsToProto converts a CertificateConfigX509ConfigAdditionalExtensions object to its proto representation.
func PrivatecaBetaCertificateConfigX509ConfigAdditionalExtensionsToProto(o *beta.CertificateConfigX509ConfigAdditionalExtensions) *betapb.PrivatecaBetaCertificateConfigX509ConfigAdditionalExtensions {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateConfigX509ConfigAdditionalExtensions{}
	p.SetObjectId(PrivatecaBetaCertificateConfigX509ConfigAdditionalExtensionsObjectIdToProto(o.ObjectId))
	p.SetCritical(dcl.ValueOrEmptyBool(o.Critical))
	p.SetValue(dcl.ValueOrEmptyString(o.Value))
	return p
}

// CertificateConfigX509ConfigAdditionalExtensionsObjectIdToProto converts a CertificateConfigX509ConfigAdditionalExtensionsObjectId object to its proto representation.
func PrivatecaBetaCertificateConfigX509ConfigAdditionalExtensionsObjectIdToProto(o *beta.CertificateConfigX509ConfigAdditionalExtensionsObjectId) *betapb.PrivatecaBetaCertificateConfigX509ConfigAdditionalExtensionsObjectId {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateConfigX509ConfigAdditionalExtensionsObjectId{}
	sObjectIdPath := make([]int64, len(o.ObjectIdPath))
	for i, r := range o.ObjectIdPath {
		sObjectIdPath[i] = r
	}
	p.SetObjectIdPath(sObjectIdPath)
	return p
}

// CertificateConfigPublicKeyToProto converts a CertificateConfigPublicKey object to its proto representation.
func PrivatecaBetaCertificateConfigPublicKeyToProto(o *beta.CertificateConfigPublicKey) *betapb.PrivatecaBetaCertificateConfigPublicKey {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateConfigPublicKey{}
	p.SetKey(dcl.ValueOrEmptyString(o.Key))
	p.SetFormat(PrivatecaBetaCertificateConfigPublicKeyFormatEnumToProto(o.Format))
	return p
}

// CertificateRevocationDetailsToProto converts a CertificateRevocationDetails object to its proto representation.
func PrivatecaBetaCertificateRevocationDetailsToProto(o *beta.CertificateRevocationDetails) *betapb.PrivatecaBetaCertificateRevocationDetails {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateRevocationDetails{}
	p.SetRevocationState(PrivatecaBetaCertificateRevocationDetailsRevocationStateEnumToProto(o.RevocationState))
	p.SetRevocationTime(dcl.ValueOrEmptyString(o.RevocationTime))
	return p
}

// CertificateCertificateDescriptionToProto converts a CertificateCertificateDescription object to its proto representation.
func PrivatecaBetaCertificateCertificateDescriptionToProto(o *beta.CertificateCertificateDescription) *betapb.PrivatecaBetaCertificateCertificateDescription {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateCertificateDescription{}
	p.SetSubjectDescription(PrivatecaBetaCertificateCertificateDescriptionSubjectDescriptionToProto(o.SubjectDescription))
	p.SetX509Description(PrivatecaBetaCertificateCertificateDescriptionX509DescriptionToProto(o.X509Description))
	p.SetPublicKey(PrivatecaBetaCertificateCertificateDescriptionPublicKeyToProto(o.PublicKey))
	p.SetSubjectKeyId(PrivatecaBetaCertificateCertificateDescriptionSubjectKeyIdToProto(o.SubjectKeyId))
	p.SetAuthorityKeyId(PrivatecaBetaCertificateCertificateDescriptionAuthorityKeyIdToProto(o.AuthorityKeyId))
	p.SetCertFingerprint(PrivatecaBetaCertificateCertificateDescriptionCertFingerprintToProto(o.CertFingerprint))
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
func PrivatecaBetaCertificateCertificateDescriptionSubjectDescriptionToProto(o *beta.CertificateCertificateDescriptionSubjectDescription) *betapb.PrivatecaBetaCertificateCertificateDescriptionSubjectDescription {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateCertificateDescriptionSubjectDescription{}
	p.SetSubject(PrivatecaBetaCertificateCertificateDescriptionSubjectDescriptionSubjectToProto(o.Subject))
	p.SetSubjectAltName(PrivatecaBetaCertificateCertificateDescriptionSubjectDescriptionSubjectAltNameToProto(o.SubjectAltName))
	p.SetHexSerialNumber(dcl.ValueOrEmptyString(o.HexSerialNumber))
	p.SetLifetime(dcl.ValueOrEmptyString(o.Lifetime))
	p.SetNotBeforeTime(dcl.ValueOrEmptyString(o.NotBeforeTime))
	p.SetNotAfterTime(dcl.ValueOrEmptyString(o.NotAfterTime))
	return p
}

// CertificateCertificateDescriptionSubjectDescriptionSubjectToProto converts a CertificateCertificateDescriptionSubjectDescriptionSubject object to its proto representation.
func PrivatecaBetaCertificateCertificateDescriptionSubjectDescriptionSubjectToProto(o *beta.CertificateCertificateDescriptionSubjectDescriptionSubject) *betapb.PrivatecaBetaCertificateCertificateDescriptionSubjectDescriptionSubject {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateCertificateDescriptionSubjectDescriptionSubject{}
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
func PrivatecaBetaCertificateCertificateDescriptionSubjectDescriptionSubjectAltNameToProto(o *beta.CertificateCertificateDescriptionSubjectDescriptionSubjectAltName) *betapb.PrivatecaBetaCertificateCertificateDescriptionSubjectDescriptionSubjectAltName {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateCertificateDescriptionSubjectDescriptionSubjectAltName{}
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
	sCustomSans := make([]*betapb.PrivatecaBetaCertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSans, len(o.CustomSans))
	for i, r := range o.CustomSans {
		sCustomSans[i] = PrivatecaBetaCertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansToProto(&r)
	}
	p.SetCustomSans(sCustomSans)
	return p
}

// CertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansToProto converts a CertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSans object to its proto representation.
func PrivatecaBetaCertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansToProto(o *beta.CertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSans) *betapb.PrivatecaBetaCertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSans {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSans{}
	p.SetObjectId(PrivatecaBetaCertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansObjectIdToProto(o.ObjectId))
	p.SetCritical(dcl.ValueOrEmptyBool(o.Critical))
	p.SetValue(dcl.ValueOrEmptyString(o.Value))
	return p
}

// CertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansObjectIdToProto converts a CertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansObjectId object to its proto representation.
func PrivatecaBetaCertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansObjectIdToProto(o *beta.CertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansObjectId) *betapb.PrivatecaBetaCertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansObjectId {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansObjectId{}
	sObjectIdPath := make([]int64, len(o.ObjectIdPath))
	for i, r := range o.ObjectIdPath {
		sObjectIdPath[i] = r
	}
	p.SetObjectIdPath(sObjectIdPath)
	return p
}

// CertificateCertificateDescriptionX509DescriptionToProto converts a CertificateCertificateDescriptionX509Description object to its proto representation.
func PrivatecaBetaCertificateCertificateDescriptionX509DescriptionToProto(o *beta.CertificateCertificateDescriptionX509Description) *betapb.PrivatecaBetaCertificateCertificateDescriptionX509Description {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateCertificateDescriptionX509Description{}
	p.SetKeyUsage(PrivatecaBetaCertificateCertificateDescriptionX509DescriptionKeyUsageToProto(o.KeyUsage))
	p.SetCaOptions(PrivatecaBetaCertificateCertificateDescriptionX509DescriptionCaOptionsToProto(o.CaOptions))
	sPolicyIds := make([]*betapb.PrivatecaBetaCertificateCertificateDescriptionX509DescriptionPolicyIds, len(o.PolicyIds))
	for i, r := range o.PolicyIds {
		sPolicyIds[i] = PrivatecaBetaCertificateCertificateDescriptionX509DescriptionPolicyIdsToProto(&r)
	}
	p.SetPolicyIds(sPolicyIds)
	sAiaOcspServers := make([]string, len(o.AiaOcspServers))
	for i, r := range o.AiaOcspServers {
		sAiaOcspServers[i] = r
	}
	p.SetAiaOcspServers(sAiaOcspServers)
	sAdditionalExtensions := make([]*betapb.PrivatecaBetaCertificateCertificateDescriptionX509DescriptionAdditionalExtensions, len(o.AdditionalExtensions))
	for i, r := range o.AdditionalExtensions {
		sAdditionalExtensions[i] = PrivatecaBetaCertificateCertificateDescriptionX509DescriptionAdditionalExtensionsToProto(&r)
	}
	p.SetAdditionalExtensions(sAdditionalExtensions)
	return p
}

// CertificateCertificateDescriptionX509DescriptionKeyUsageToProto converts a CertificateCertificateDescriptionX509DescriptionKeyUsage object to its proto representation.
func PrivatecaBetaCertificateCertificateDescriptionX509DescriptionKeyUsageToProto(o *beta.CertificateCertificateDescriptionX509DescriptionKeyUsage) *betapb.PrivatecaBetaCertificateCertificateDescriptionX509DescriptionKeyUsage {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateCertificateDescriptionX509DescriptionKeyUsage{}
	p.SetBaseKeyUsage(PrivatecaBetaCertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsageToProto(o.BaseKeyUsage))
	p.SetExtendedKeyUsage(PrivatecaBetaCertificateCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsageToProto(o.ExtendedKeyUsage))
	sUnknownExtendedKeyUsages := make([]*betapb.PrivatecaBetaCertificateCertificateDescriptionX509DescriptionKeyUsageUnknownExtendedKeyUsages, len(o.UnknownExtendedKeyUsages))
	for i, r := range o.UnknownExtendedKeyUsages {
		sUnknownExtendedKeyUsages[i] = PrivatecaBetaCertificateCertificateDescriptionX509DescriptionKeyUsageUnknownExtendedKeyUsagesToProto(&r)
	}
	p.SetUnknownExtendedKeyUsages(sUnknownExtendedKeyUsages)
	return p
}

// CertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsageToProto converts a CertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsage object to its proto representation.
func PrivatecaBetaCertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsageToProto(o *beta.CertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsage) *betapb.PrivatecaBetaCertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsage {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsage{}
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
func PrivatecaBetaCertificateCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsageToProto(o *beta.CertificateCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsage) *betapb.PrivatecaBetaCertificateCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsage {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsage{}
	p.SetServerAuth(dcl.ValueOrEmptyBool(o.ServerAuth))
	p.SetClientAuth(dcl.ValueOrEmptyBool(o.ClientAuth))
	p.SetCodeSigning(dcl.ValueOrEmptyBool(o.CodeSigning))
	p.SetEmailProtection(dcl.ValueOrEmptyBool(o.EmailProtection))
	p.SetTimeStamping(dcl.ValueOrEmptyBool(o.TimeStamping))
	p.SetOcspSigning(dcl.ValueOrEmptyBool(o.OcspSigning))
	return p
}

// CertificateCertificateDescriptionX509DescriptionKeyUsageUnknownExtendedKeyUsagesToProto converts a CertificateCertificateDescriptionX509DescriptionKeyUsageUnknownExtendedKeyUsages object to its proto representation.
func PrivatecaBetaCertificateCertificateDescriptionX509DescriptionKeyUsageUnknownExtendedKeyUsagesToProto(o *beta.CertificateCertificateDescriptionX509DescriptionKeyUsageUnknownExtendedKeyUsages) *betapb.PrivatecaBetaCertificateCertificateDescriptionX509DescriptionKeyUsageUnknownExtendedKeyUsages {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateCertificateDescriptionX509DescriptionKeyUsageUnknownExtendedKeyUsages{}
	sObjectIdPath := make([]int64, len(o.ObjectIdPath))
	for i, r := range o.ObjectIdPath {
		sObjectIdPath[i] = r
	}
	p.SetObjectIdPath(sObjectIdPath)
	return p
}

// CertificateCertificateDescriptionX509DescriptionCaOptionsToProto converts a CertificateCertificateDescriptionX509DescriptionCaOptions object to its proto representation.
func PrivatecaBetaCertificateCertificateDescriptionX509DescriptionCaOptionsToProto(o *beta.CertificateCertificateDescriptionX509DescriptionCaOptions) *betapb.PrivatecaBetaCertificateCertificateDescriptionX509DescriptionCaOptions {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateCertificateDescriptionX509DescriptionCaOptions{}
	p.SetIsCa(dcl.ValueOrEmptyBool(o.IsCa))
	p.SetMaxIssuerPathLength(dcl.ValueOrEmptyInt64(o.MaxIssuerPathLength))
	return p
}

// CertificateCertificateDescriptionX509DescriptionPolicyIdsToProto converts a CertificateCertificateDescriptionX509DescriptionPolicyIds object to its proto representation.
func PrivatecaBetaCertificateCertificateDescriptionX509DescriptionPolicyIdsToProto(o *beta.CertificateCertificateDescriptionX509DescriptionPolicyIds) *betapb.PrivatecaBetaCertificateCertificateDescriptionX509DescriptionPolicyIds {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateCertificateDescriptionX509DescriptionPolicyIds{}
	sObjectIdPath := make([]int64, len(o.ObjectIdPath))
	for i, r := range o.ObjectIdPath {
		sObjectIdPath[i] = r
	}
	p.SetObjectIdPath(sObjectIdPath)
	return p
}

// CertificateCertificateDescriptionX509DescriptionAdditionalExtensionsToProto converts a CertificateCertificateDescriptionX509DescriptionAdditionalExtensions object to its proto representation.
func PrivatecaBetaCertificateCertificateDescriptionX509DescriptionAdditionalExtensionsToProto(o *beta.CertificateCertificateDescriptionX509DescriptionAdditionalExtensions) *betapb.PrivatecaBetaCertificateCertificateDescriptionX509DescriptionAdditionalExtensions {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateCertificateDescriptionX509DescriptionAdditionalExtensions{}
	p.SetObjectId(PrivatecaBetaCertificateCertificateDescriptionX509DescriptionAdditionalExtensionsObjectIdToProto(o.ObjectId))
	p.SetCritical(dcl.ValueOrEmptyBool(o.Critical))
	p.SetValue(dcl.ValueOrEmptyString(o.Value))
	return p
}

// CertificateCertificateDescriptionX509DescriptionAdditionalExtensionsObjectIdToProto converts a CertificateCertificateDescriptionX509DescriptionAdditionalExtensionsObjectId object to its proto representation.
func PrivatecaBetaCertificateCertificateDescriptionX509DescriptionAdditionalExtensionsObjectIdToProto(o *beta.CertificateCertificateDescriptionX509DescriptionAdditionalExtensionsObjectId) *betapb.PrivatecaBetaCertificateCertificateDescriptionX509DescriptionAdditionalExtensionsObjectId {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateCertificateDescriptionX509DescriptionAdditionalExtensionsObjectId{}
	sObjectIdPath := make([]int64, len(o.ObjectIdPath))
	for i, r := range o.ObjectIdPath {
		sObjectIdPath[i] = r
	}
	p.SetObjectIdPath(sObjectIdPath)
	return p
}

// CertificateCertificateDescriptionPublicKeyToProto converts a CertificateCertificateDescriptionPublicKey object to its proto representation.
func PrivatecaBetaCertificateCertificateDescriptionPublicKeyToProto(o *beta.CertificateCertificateDescriptionPublicKey) *betapb.PrivatecaBetaCertificateCertificateDescriptionPublicKey {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateCertificateDescriptionPublicKey{}
	p.SetKey(dcl.ValueOrEmptyString(o.Key))
	p.SetFormat(PrivatecaBetaCertificateCertificateDescriptionPublicKeyFormatEnumToProto(o.Format))
	return p
}

// CertificateCertificateDescriptionSubjectKeyIdToProto converts a CertificateCertificateDescriptionSubjectKeyId object to its proto representation.
func PrivatecaBetaCertificateCertificateDescriptionSubjectKeyIdToProto(o *beta.CertificateCertificateDescriptionSubjectKeyId) *betapb.PrivatecaBetaCertificateCertificateDescriptionSubjectKeyId {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateCertificateDescriptionSubjectKeyId{}
	p.SetKeyId(dcl.ValueOrEmptyString(o.KeyId))
	return p
}

// CertificateCertificateDescriptionAuthorityKeyIdToProto converts a CertificateCertificateDescriptionAuthorityKeyId object to its proto representation.
func PrivatecaBetaCertificateCertificateDescriptionAuthorityKeyIdToProto(o *beta.CertificateCertificateDescriptionAuthorityKeyId) *betapb.PrivatecaBetaCertificateCertificateDescriptionAuthorityKeyId {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateCertificateDescriptionAuthorityKeyId{}
	p.SetKeyId(dcl.ValueOrEmptyString(o.KeyId))
	return p
}

// CertificateCertificateDescriptionCertFingerprintToProto converts a CertificateCertificateDescriptionCertFingerprint object to its proto representation.
func PrivatecaBetaCertificateCertificateDescriptionCertFingerprintToProto(o *beta.CertificateCertificateDescriptionCertFingerprint) *betapb.PrivatecaBetaCertificateCertificateDescriptionCertFingerprint {
	if o == nil {
		return nil
	}
	p := &betapb.PrivatecaBetaCertificateCertificateDescriptionCertFingerprint{}
	p.SetSha256Hash(dcl.ValueOrEmptyString(o.Sha256Hash))
	return p
}

// CertificateToProto converts a Certificate resource to its proto representation.
func CertificateToProto(resource *beta.Certificate) *betapb.PrivatecaBetaCertificate {
	p := &betapb.PrivatecaBetaCertificate{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetPemCsr(dcl.ValueOrEmptyString(resource.PemCsr))
	p.SetConfig(PrivatecaBetaCertificateConfigToProto(resource.Config))
	p.SetIssuerCertificateAuthority(dcl.ValueOrEmptyString(resource.IssuerCertificateAuthority))
	p.SetLifetime(dcl.ValueOrEmptyString(resource.Lifetime))
	p.SetCertificateTemplate(dcl.ValueOrEmptyString(resource.CertificateTemplate))
	p.SetSubjectMode(PrivatecaBetaCertificateSubjectModeEnumToProto(resource.SubjectMode))
	p.SetRevocationDetails(PrivatecaBetaCertificateRevocationDetailsToProto(resource.RevocationDetails))
	p.SetPemCertificate(dcl.ValueOrEmptyString(resource.PemCertificate))
	p.SetCertificateDescription(PrivatecaBetaCertificateCertificateDescriptionToProto(resource.CertificateDescription))
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
func (s *CertificateServer) applyCertificate(ctx context.Context, c *beta.Client, request *betapb.ApplyPrivatecaBetaCertificateRequest) (*betapb.PrivatecaBetaCertificate, error) {
	p := ProtoToCertificate(request.GetResource())
	res, err := c.ApplyCertificate(ctx, p)
	if err != nil {
		return nil, err
	}
	r := CertificateToProto(res)
	return r, nil
}

// applyPrivatecaBetaCertificate handles the gRPC request by passing it to the underlying Certificate Apply() method.
func (s *CertificateServer) ApplyPrivatecaBetaCertificate(ctx context.Context, request *betapb.ApplyPrivatecaBetaCertificateRequest) (*betapb.PrivatecaBetaCertificate, error) {
	cl, err := createConfigCertificate(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyCertificate(ctx, cl, request)
}

// DeleteCertificate handles the gRPC request by passing it to the underlying Certificate Delete() method.
func (s *CertificateServer) DeletePrivatecaBetaCertificate(ctx context.Context, request *betapb.DeletePrivatecaBetaCertificateRequest) (*emptypb.Empty, error) {

	cl, err := createConfigCertificate(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteCertificate(ctx, ProtoToCertificate(request.GetResource()))

}

// ListPrivatecaBetaCertificate handles the gRPC request by passing it to the underlying CertificateList() method.
func (s *CertificateServer) ListPrivatecaBetaCertificate(ctx context.Context, request *betapb.ListPrivatecaBetaCertificateRequest) (*betapb.ListPrivatecaBetaCertificateResponse, error) {
	cl, err := createConfigCertificate(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListCertificate(ctx, request.GetProject(), request.GetLocation(), request.GetCaPool())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.PrivatecaBetaCertificate
	for _, r := range resources.Items {
		rp := CertificateToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListPrivatecaBetaCertificateResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigCertificate(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
