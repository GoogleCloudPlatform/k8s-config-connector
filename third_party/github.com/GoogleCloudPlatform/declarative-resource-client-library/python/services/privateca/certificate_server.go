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

// CertificateServer implements the gRPC interface for Certificate.
type CertificateServer struct{}

// ProtoToCertificateConfigPublicKeyFormatEnum converts a CertificateConfigPublicKeyFormatEnum enum from its proto representation.
func ProtoToPrivatecaCertificateConfigPublicKeyFormatEnum(e privatecapb.PrivatecaCertificateConfigPublicKeyFormatEnum) *privateca.CertificateConfigPublicKeyFormatEnum {
	if e == 0 {
		return nil
	}
	if n, ok := privatecapb.PrivatecaCertificateConfigPublicKeyFormatEnum_name[int32(e)]; ok {
		e := privateca.CertificateConfigPublicKeyFormatEnum(n[len("PrivatecaCertificateConfigPublicKeyFormatEnum"):])
		return &e
	}
	return nil
}

// ProtoToCertificateSubjectModeEnum converts a CertificateSubjectModeEnum enum from its proto representation.
func ProtoToPrivatecaCertificateSubjectModeEnum(e privatecapb.PrivatecaCertificateSubjectModeEnum) *privateca.CertificateSubjectModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := privatecapb.PrivatecaCertificateSubjectModeEnum_name[int32(e)]; ok {
		e := privateca.CertificateSubjectModeEnum(n[len("PrivatecaCertificateSubjectModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToCertificateRevocationDetailsRevocationStateEnum converts a CertificateRevocationDetailsRevocationStateEnum enum from its proto representation.
func ProtoToPrivatecaCertificateRevocationDetailsRevocationStateEnum(e privatecapb.PrivatecaCertificateRevocationDetailsRevocationStateEnum) *privateca.CertificateRevocationDetailsRevocationStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := privatecapb.PrivatecaCertificateRevocationDetailsRevocationStateEnum_name[int32(e)]; ok {
		e := privateca.CertificateRevocationDetailsRevocationStateEnum(n[len("PrivatecaCertificateRevocationDetailsRevocationStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToCertificateCertificateDescriptionPublicKeyFormatEnum converts a CertificateCertificateDescriptionPublicKeyFormatEnum enum from its proto representation.
func ProtoToPrivatecaCertificateCertificateDescriptionPublicKeyFormatEnum(e privatecapb.PrivatecaCertificateCertificateDescriptionPublicKeyFormatEnum) *privateca.CertificateCertificateDescriptionPublicKeyFormatEnum {
	if e == 0 {
		return nil
	}
	if n, ok := privatecapb.PrivatecaCertificateCertificateDescriptionPublicKeyFormatEnum_name[int32(e)]; ok {
		e := privateca.CertificateCertificateDescriptionPublicKeyFormatEnum(n[len("PrivatecaCertificateCertificateDescriptionPublicKeyFormatEnum"):])
		return &e
	}
	return nil
}

// ProtoToCertificateConfig converts a CertificateConfig object from its proto representation.
func ProtoToPrivatecaCertificateConfig(p *privatecapb.PrivatecaCertificateConfig) *privateca.CertificateConfig {
	if p == nil {
		return nil
	}
	obj := &privateca.CertificateConfig{
		SubjectConfig: ProtoToPrivatecaCertificateConfigSubjectConfig(p.GetSubjectConfig()),
		X509Config:    ProtoToPrivatecaCertificateConfigX509Config(p.GetX509Config()),
		PublicKey:     ProtoToPrivatecaCertificateConfigPublicKey(p.GetPublicKey()),
	}
	return obj
}

// ProtoToCertificateConfigSubjectConfig converts a CertificateConfigSubjectConfig object from its proto representation.
func ProtoToPrivatecaCertificateConfigSubjectConfig(p *privatecapb.PrivatecaCertificateConfigSubjectConfig) *privateca.CertificateConfigSubjectConfig {
	if p == nil {
		return nil
	}
	obj := &privateca.CertificateConfigSubjectConfig{
		Subject:        ProtoToPrivatecaCertificateConfigSubjectConfigSubject(p.GetSubject()),
		SubjectAltName: ProtoToPrivatecaCertificateConfigSubjectConfigSubjectAltName(p.GetSubjectAltName()),
	}
	return obj
}

// ProtoToCertificateConfigSubjectConfigSubject converts a CertificateConfigSubjectConfigSubject object from its proto representation.
func ProtoToPrivatecaCertificateConfigSubjectConfigSubject(p *privatecapb.PrivatecaCertificateConfigSubjectConfigSubject) *privateca.CertificateConfigSubjectConfigSubject {
	if p == nil {
		return nil
	}
	obj := &privateca.CertificateConfigSubjectConfigSubject{
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
func ProtoToPrivatecaCertificateConfigSubjectConfigSubjectAltName(p *privatecapb.PrivatecaCertificateConfigSubjectConfigSubjectAltName) *privateca.CertificateConfigSubjectConfigSubjectAltName {
	if p == nil {
		return nil
	}
	obj := &privateca.CertificateConfigSubjectConfigSubjectAltName{}
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
func ProtoToPrivatecaCertificateConfigX509Config(p *privatecapb.PrivatecaCertificateConfigX509Config) *privateca.CertificateConfigX509Config {
	if p == nil {
		return nil
	}
	obj := &privateca.CertificateConfigX509Config{
		KeyUsage:  ProtoToPrivatecaCertificateConfigX509ConfigKeyUsage(p.GetKeyUsage()),
		CaOptions: ProtoToPrivatecaCertificateConfigX509ConfigCaOptions(p.GetCaOptions()),
	}
	for _, r := range p.GetPolicyIds() {
		obj.PolicyIds = append(obj.PolicyIds, *ProtoToPrivatecaCertificateConfigX509ConfigPolicyIds(r))
	}
	for _, r := range p.GetAiaOcspServers() {
		obj.AiaOcspServers = append(obj.AiaOcspServers, r)
	}
	for _, r := range p.GetAdditionalExtensions() {
		obj.AdditionalExtensions = append(obj.AdditionalExtensions, *ProtoToPrivatecaCertificateConfigX509ConfigAdditionalExtensions(r))
	}
	return obj
}

// ProtoToCertificateConfigX509ConfigKeyUsage converts a CertificateConfigX509ConfigKeyUsage object from its proto representation.
func ProtoToPrivatecaCertificateConfigX509ConfigKeyUsage(p *privatecapb.PrivatecaCertificateConfigX509ConfigKeyUsage) *privateca.CertificateConfigX509ConfigKeyUsage {
	if p == nil {
		return nil
	}
	obj := &privateca.CertificateConfigX509ConfigKeyUsage{
		BaseKeyUsage:     ProtoToPrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsage(p.GetBaseKeyUsage()),
		ExtendedKeyUsage: ProtoToPrivatecaCertificateConfigX509ConfigKeyUsageExtendedKeyUsage(p.GetExtendedKeyUsage()),
	}
	for _, r := range p.GetUnknownExtendedKeyUsages() {
		obj.UnknownExtendedKeyUsages = append(obj.UnknownExtendedKeyUsages, *ProtoToPrivatecaCertificateConfigX509ConfigKeyUsageUnknownExtendedKeyUsages(r))
	}
	return obj
}

// ProtoToCertificateConfigX509ConfigKeyUsageBaseKeyUsage converts a CertificateConfigX509ConfigKeyUsageBaseKeyUsage object from its proto representation.
func ProtoToPrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsage(p *privatecapb.PrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsage) *privateca.CertificateConfigX509ConfigKeyUsageBaseKeyUsage {
	if p == nil {
		return nil
	}
	obj := &privateca.CertificateConfigX509ConfigKeyUsageBaseKeyUsage{
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
func ProtoToPrivatecaCertificateConfigX509ConfigKeyUsageExtendedKeyUsage(p *privatecapb.PrivatecaCertificateConfigX509ConfigKeyUsageExtendedKeyUsage) *privateca.CertificateConfigX509ConfigKeyUsageExtendedKeyUsage {
	if p == nil {
		return nil
	}
	obj := &privateca.CertificateConfigX509ConfigKeyUsageExtendedKeyUsage{
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
func ProtoToPrivatecaCertificateConfigX509ConfigKeyUsageUnknownExtendedKeyUsages(p *privatecapb.PrivatecaCertificateConfigX509ConfigKeyUsageUnknownExtendedKeyUsages) *privateca.CertificateConfigX509ConfigKeyUsageUnknownExtendedKeyUsages {
	if p == nil {
		return nil
	}
	obj := &privateca.CertificateConfigX509ConfigKeyUsageUnknownExtendedKeyUsages{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCertificateConfigX509ConfigCaOptions converts a CertificateConfigX509ConfigCaOptions object from its proto representation.
func ProtoToPrivatecaCertificateConfigX509ConfigCaOptions(p *privatecapb.PrivatecaCertificateConfigX509ConfigCaOptions) *privateca.CertificateConfigX509ConfigCaOptions {
	if p == nil {
		return nil
	}
	obj := &privateca.CertificateConfigX509ConfigCaOptions{
		IsCa:                    dcl.Bool(p.GetIsCa()),
		NonCa:                   dcl.Bool(p.GetNonCa()),
		MaxIssuerPathLength:     dcl.Int64OrNil(p.GetMaxIssuerPathLength()),
		ZeroMaxIssuerPathLength: dcl.Bool(p.GetZeroMaxIssuerPathLength()),
	}
	return obj
}

// ProtoToCertificateConfigX509ConfigPolicyIds converts a CertificateConfigX509ConfigPolicyIds object from its proto representation.
func ProtoToPrivatecaCertificateConfigX509ConfigPolicyIds(p *privatecapb.PrivatecaCertificateConfigX509ConfigPolicyIds) *privateca.CertificateConfigX509ConfigPolicyIds {
	if p == nil {
		return nil
	}
	obj := &privateca.CertificateConfigX509ConfigPolicyIds{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCertificateConfigX509ConfigAdditionalExtensions converts a CertificateConfigX509ConfigAdditionalExtensions object from its proto representation.
func ProtoToPrivatecaCertificateConfigX509ConfigAdditionalExtensions(p *privatecapb.PrivatecaCertificateConfigX509ConfigAdditionalExtensions) *privateca.CertificateConfigX509ConfigAdditionalExtensions {
	if p == nil {
		return nil
	}
	obj := &privateca.CertificateConfigX509ConfigAdditionalExtensions{
		ObjectId: ProtoToPrivatecaCertificateConfigX509ConfigAdditionalExtensionsObjectId(p.GetObjectId()),
		Critical: dcl.Bool(p.GetCritical()),
		Value:    dcl.StringOrNil(p.GetValue()),
	}
	return obj
}

// ProtoToCertificateConfigX509ConfigAdditionalExtensionsObjectId converts a CertificateConfigX509ConfigAdditionalExtensionsObjectId object from its proto representation.
func ProtoToPrivatecaCertificateConfigX509ConfigAdditionalExtensionsObjectId(p *privatecapb.PrivatecaCertificateConfigX509ConfigAdditionalExtensionsObjectId) *privateca.CertificateConfigX509ConfigAdditionalExtensionsObjectId {
	if p == nil {
		return nil
	}
	obj := &privateca.CertificateConfigX509ConfigAdditionalExtensionsObjectId{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCertificateConfigPublicKey converts a CertificateConfigPublicKey object from its proto representation.
func ProtoToPrivatecaCertificateConfigPublicKey(p *privatecapb.PrivatecaCertificateConfigPublicKey) *privateca.CertificateConfigPublicKey {
	if p == nil {
		return nil
	}
	obj := &privateca.CertificateConfigPublicKey{
		Key:    dcl.StringOrNil(p.GetKey()),
		Format: ProtoToPrivatecaCertificateConfigPublicKeyFormatEnum(p.GetFormat()),
	}
	return obj
}

// ProtoToCertificateRevocationDetails converts a CertificateRevocationDetails object from its proto representation.
func ProtoToPrivatecaCertificateRevocationDetails(p *privatecapb.PrivatecaCertificateRevocationDetails) *privateca.CertificateRevocationDetails {
	if p == nil {
		return nil
	}
	obj := &privateca.CertificateRevocationDetails{
		RevocationState: ProtoToPrivatecaCertificateRevocationDetailsRevocationStateEnum(p.GetRevocationState()),
		RevocationTime:  dcl.StringOrNil(p.GetRevocationTime()),
	}
	return obj
}

// ProtoToCertificateCertificateDescription converts a CertificateCertificateDescription object from its proto representation.
func ProtoToPrivatecaCertificateCertificateDescription(p *privatecapb.PrivatecaCertificateCertificateDescription) *privateca.CertificateCertificateDescription {
	if p == nil {
		return nil
	}
	obj := &privateca.CertificateCertificateDescription{
		SubjectDescription: ProtoToPrivatecaCertificateCertificateDescriptionSubjectDescription(p.GetSubjectDescription()),
		X509Description:    ProtoToPrivatecaCertificateCertificateDescriptionX509Description(p.GetX509Description()),
		PublicKey:          ProtoToPrivatecaCertificateCertificateDescriptionPublicKey(p.GetPublicKey()),
		SubjectKeyId:       ProtoToPrivatecaCertificateCertificateDescriptionSubjectKeyId(p.GetSubjectKeyId()),
		AuthorityKeyId:     ProtoToPrivatecaCertificateCertificateDescriptionAuthorityKeyId(p.GetAuthorityKeyId()),
		CertFingerprint:    ProtoToPrivatecaCertificateCertificateDescriptionCertFingerprint(p.GetCertFingerprint()),
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
func ProtoToPrivatecaCertificateCertificateDescriptionSubjectDescription(p *privatecapb.PrivatecaCertificateCertificateDescriptionSubjectDescription) *privateca.CertificateCertificateDescriptionSubjectDescription {
	if p == nil {
		return nil
	}
	obj := &privateca.CertificateCertificateDescriptionSubjectDescription{
		Subject:         ProtoToPrivatecaCertificateCertificateDescriptionSubjectDescriptionSubject(p.GetSubject()),
		SubjectAltName:  ProtoToPrivatecaCertificateCertificateDescriptionSubjectDescriptionSubjectAltName(p.GetSubjectAltName()),
		HexSerialNumber: dcl.StringOrNil(p.GetHexSerialNumber()),
		Lifetime:        dcl.StringOrNil(p.GetLifetime()),
		NotBeforeTime:   dcl.StringOrNil(p.GetNotBeforeTime()),
		NotAfterTime:    dcl.StringOrNil(p.GetNotAfterTime()),
	}
	return obj
}

// ProtoToCertificateCertificateDescriptionSubjectDescriptionSubject converts a CertificateCertificateDescriptionSubjectDescriptionSubject object from its proto representation.
func ProtoToPrivatecaCertificateCertificateDescriptionSubjectDescriptionSubject(p *privatecapb.PrivatecaCertificateCertificateDescriptionSubjectDescriptionSubject) *privateca.CertificateCertificateDescriptionSubjectDescriptionSubject {
	if p == nil {
		return nil
	}
	obj := &privateca.CertificateCertificateDescriptionSubjectDescriptionSubject{
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
func ProtoToPrivatecaCertificateCertificateDescriptionSubjectDescriptionSubjectAltName(p *privatecapb.PrivatecaCertificateCertificateDescriptionSubjectDescriptionSubjectAltName) *privateca.CertificateCertificateDescriptionSubjectDescriptionSubjectAltName {
	if p == nil {
		return nil
	}
	obj := &privateca.CertificateCertificateDescriptionSubjectDescriptionSubjectAltName{}
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
		obj.CustomSans = append(obj.CustomSans, *ProtoToPrivatecaCertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSans(r))
	}
	return obj
}

// ProtoToCertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSans converts a CertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSans object from its proto representation.
func ProtoToPrivatecaCertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSans(p *privatecapb.PrivatecaCertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSans) *privateca.CertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSans {
	if p == nil {
		return nil
	}
	obj := &privateca.CertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSans{
		ObjectId: ProtoToPrivatecaCertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansObjectId(p.GetObjectId()),
		Critical: dcl.Bool(p.GetCritical()),
		Value:    dcl.StringOrNil(p.GetValue()),
	}
	return obj
}

// ProtoToCertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansObjectId converts a CertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansObjectId object from its proto representation.
func ProtoToPrivatecaCertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansObjectId(p *privatecapb.PrivatecaCertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansObjectId) *privateca.CertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansObjectId {
	if p == nil {
		return nil
	}
	obj := &privateca.CertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansObjectId{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCertificateCertificateDescriptionX509Description converts a CertificateCertificateDescriptionX509Description object from its proto representation.
func ProtoToPrivatecaCertificateCertificateDescriptionX509Description(p *privatecapb.PrivatecaCertificateCertificateDescriptionX509Description) *privateca.CertificateCertificateDescriptionX509Description {
	if p == nil {
		return nil
	}
	obj := &privateca.CertificateCertificateDescriptionX509Description{
		KeyUsage:  ProtoToPrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsage(p.GetKeyUsage()),
		CaOptions: ProtoToPrivatecaCertificateCertificateDescriptionX509DescriptionCaOptions(p.GetCaOptions()),
	}
	for _, r := range p.GetPolicyIds() {
		obj.PolicyIds = append(obj.PolicyIds, *ProtoToPrivatecaCertificateCertificateDescriptionX509DescriptionPolicyIds(r))
	}
	for _, r := range p.GetAiaOcspServers() {
		obj.AiaOcspServers = append(obj.AiaOcspServers, r)
	}
	for _, r := range p.GetAdditionalExtensions() {
		obj.AdditionalExtensions = append(obj.AdditionalExtensions, *ProtoToPrivatecaCertificateCertificateDescriptionX509DescriptionAdditionalExtensions(r))
	}
	return obj
}

// ProtoToCertificateCertificateDescriptionX509DescriptionKeyUsage converts a CertificateCertificateDescriptionX509DescriptionKeyUsage object from its proto representation.
func ProtoToPrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsage(p *privatecapb.PrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsage) *privateca.CertificateCertificateDescriptionX509DescriptionKeyUsage {
	if p == nil {
		return nil
	}
	obj := &privateca.CertificateCertificateDescriptionX509DescriptionKeyUsage{
		BaseKeyUsage:     ProtoToPrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsage(p.GetBaseKeyUsage()),
		ExtendedKeyUsage: ProtoToPrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsage(p.GetExtendedKeyUsage()),
	}
	for _, r := range p.GetUnknownExtendedKeyUsages() {
		obj.UnknownExtendedKeyUsages = append(obj.UnknownExtendedKeyUsages, *ProtoToPrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageUnknownExtendedKeyUsages(r))
	}
	return obj
}

// ProtoToCertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsage converts a CertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsage object from its proto representation.
func ProtoToPrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsage(p *privatecapb.PrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsage) *privateca.CertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsage {
	if p == nil {
		return nil
	}
	obj := &privateca.CertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsage{
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
func ProtoToPrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsage(p *privatecapb.PrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsage) *privateca.CertificateCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsage {
	if p == nil {
		return nil
	}
	obj := &privateca.CertificateCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsage{
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
func ProtoToPrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageUnknownExtendedKeyUsages(p *privatecapb.PrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageUnknownExtendedKeyUsages) *privateca.CertificateCertificateDescriptionX509DescriptionKeyUsageUnknownExtendedKeyUsages {
	if p == nil {
		return nil
	}
	obj := &privateca.CertificateCertificateDescriptionX509DescriptionKeyUsageUnknownExtendedKeyUsages{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCertificateCertificateDescriptionX509DescriptionCaOptions converts a CertificateCertificateDescriptionX509DescriptionCaOptions object from its proto representation.
func ProtoToPrivatecaCertificateCertificateDescriptionX509DescriptionCaOptions(p *privatecapb.PrivatecaCertificateCertificateDescriptionX509DescriptionCaOptions) *privateca.CertificateCertificateDescriptionX509DescriptionCaOptions {
	if p == nil {
		return nil
	}
	obj := &privateca.CertificateCertificateDescriptionX509DescriptionCaOptions{
		IsCa:                dcl.Bool(p.GetIsCa()),
		MaxIssuerPathLength: dcl.Int64OrNil(p.GetMaxIssuerPathLength()),
	}
	return obj
}

// ProtoToCertificateCertificateDescriptionX509DescriptionPolicyIds converts a CertificateCertificateDescriptionX509DescriptionPolicyIds object from its proto representation.
func ProtoToPrivatecaCertificateCertificateDescriptionX509DescriptionPolicyIds(p *privatecapb.PrivatecaCertificateCertificateDescriptionX509DescriptionPolicyIds) *privateca.CertificateCertificateDescriptionX509DescriptionPolicyIds {
	if p == nil {
		return nil
	}
	obj := &privateca.CertificateCertificateDescriptionX509DescriptionPolicyIds{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCertificateCertificateDescriptionX509DescriptionAdditionalExtensions converts a CertificateCertificateDescriptionX509DescriptionAdditionalExtensions object from its proto representation.
func ProtoToPrivatecaCertificateCertificateDescriptionX509DescriptionAdditionalExtensions(p *privatecapb.PrivatecaCertificateCertificateDescriptionX509DescriptionAdditionalExtensions) *privateca.CertificateCertificateDescriptionX509DescriptionAdditionalExtensions {
	if p == nil {
		return nil
	}
	obj := &privateca.CertificateCertificateDescriptionX509DescriptionAdditionalExtensions{
		ObjectId: ProtoToPrivatecaCertificateCertificateDescriptionX509DescriptionAdditionalExtensionsObjectId(p.GetObjectId()),
		Critical: dcl.Bool(p.GetCritical()),
		Value:    dcl.StringOrNil(p.GetValue()),
	}
	return obj
}

// ProtoToCertificateCertificateDescriptionX509DescriptionAdditionalExtensionsObjectId converts a CertificateCertificateDescriptionX509DescriptionAdditionalExtensionsObjectId object from its proto representation.
func ProtoToPrivatecaCertificateCertificateDescriptionX509DescriptionAdditionalExtensionsObjectId(p *privatecapb.PrivatecaCertificateCertificateDescriptionX509DescriptionAdditionalExtensionsObjectId) *privateca.CertificateCertificateDescriptionX509DescriptionAdditionalExtensionsObjectId {
	if p == nil {
		return nil
	}
	obj := &privateca.CertificateCertificateDescriptionX509DescriptionAdditionalExtensionsObjectId{}
	for _, r := range p.GetObjectIdPath() {
		obj.ObjectIdPath = append(obj.ObjectIdPath, r)
	}
	return obj
}

// ProtoToCertificateCertificateDescriptionPublicKey converts a CertificateCertificateDescriptionPublicKey object from its proto representation.
func ProtoToPrivatecaCertificateCertificateDescriptionPublicKey(p *privatecapb.PrivatecaCertificateCertificateDescriptionPublicKey) *privateca.CertificateCertificateDescriptionPublicKey {
	if p == nil {
		return nil
	}
	obj := &privateca.CertificateCertificateDescriptionPublicKey{
		Key:    dcl.StringOrNil(p.GetKey()),
		Format: ProtoToPrivatecaCertificateCertificateDescriptionPublicKeyFormatEnum(p.GetFormat()),
	}
	return obj
}

// ProtoToCertificateCertificateDescriptionSubjectKeyId converts a CertificateCertificateDescriptionSubjectKeyId object from its proto representation.
func ProtoToPrivatecaCertificateCertificateDescriptionSubjectKeyId(p *privatecapb.PrivatecaCertificateCertificateDescriptionSubjectKeyId) *privateca.CertificateCertificateDescriptionSubjectKeyId {
	if p == nil {
		return nil
	}
	obj := &privateca.CertificateCertificateDescriptionSubjectKeyId{
		KeyId: dcl.StringOrNil(p.GetKeyId()),
	}
	return obj
}

// ProtoToCertificateCertificateDescriptionAuthorityKeyId converts a CertificateCertificateDescriptionAuthorityKeyId object from its proto representation.
func ProtoToPrivatecaCertificateCertificateDescriptionAuthorityKeyId(p *privatecapb.PrivatecaCertificateCertificateDescriptionAuthorityKeyId) *privateca.CertificateCertificateDescriptionAuthorityKeyId {
	if p == nil {
		return nil
	}
	obj := &privateca.CertificateCertificateDescriptionAuthorityKeyId{
		KeyId: dcl.StringOrNil(p.GetKeyId()),
	}
	return obj
}

// ProtoToCertificateCertificateDescriptionCertFingerprint converts a CertificateCertificateDescriptionCertFingerprint object from its proto representation.
func ProtoToPrivatecaCertificateCertificateDescriptionCertFingerprint(p *privatecapb.PrivatecaCertificateCertificateDescriptionCertFingerprint) *privateca.CertificateCertificateDescriptionCertFingerprint {
	if p == nil {
		return nil
	}
	obj := &privateca.CertificateCertificateDescriptionCertFingerprint{
		Sha256Hash: dcl.StringOrNil(p.GetSha256Hash()),
	}
	return obj
}

// ProtoToCertificate converts a Certificate resource from its proto representation.
func ProtoToCertificate(p *privatecapb.PrivatecaCertificate) *privateca.Certificate {
	obj := &privateca.Certificate{
		Name:                       dcl.StringOrNil(p.GetName()),
		PemCsr:                     dcl.StringOrNil(p.GetPemCsr()),
		Config:                     ProtoToPrivatecaCertificateConfig(p.GetConfig()),
		IssuerCertificateAuthority: dcl.StringOrNil(p.GetIssuerCertificateAuthority()),
		Lifetime:                   dcl.StringOrNil(p.GetLifetime()),
		CertificateTemplate:        dcl.StringOrNil(p.GetCertificateTemplate()),
		SubjectMode:                ProtoToPrivatecaCertificateSubjectModeEnum(p.GetSubjectMode()),
		RevocationDetails:          ProtoToPrivatecaCertificateRevocationDetails(p.GetRevocationDetails()),
		PemCertificate:             dcl.StringOrNil(p.GetPemCertificate()),
		CertificateDescription:     ProtoToPrivatecaCertificateCertificateDescription(p.GetCertificateDescription()),
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
func PrivatecaCertificateConfigPublicKeyFormatEnumToProto(e *privateca.CertificateConfigPublicKeyFormatEnum) privatecapb.PrivatecaCertificateConfigPublicKeyFormatEnum {
	if e == nil {
		return privatecapb.PrivatecaCertificateConfigPublicKeyFormatEnum(0)
	}
	if v, ok := privatecapb.PrivatecaCertificateConfigPublicKeyFormatEnum_value["CertificateConfigPublicKeyFormatEnum"+string(*e)]; ok {
		return privatecapb.PrivatecaCertificateConfigPublicKeyFormatEnum(v)
	}
	return privatecapb.PrivatecaCertificateConfigPublicKeyFormatEnum(0)
}

// CertificateSubjectModeEnumToProto converts a CertificateSubjectModeEnum enum to its proto representation.
func PrivatecaCertificateSubjectModeEnumToProto(e *privateca.CertificateSubjectModeEnum) privatecapb.PrivatecaCertificateSubjectModeEnum {
	if e == nil {
		return privatecapb.PrivatecaCertificateSubjectModeEnum(0)
	}
	if v, ok := privatecapb.PrivatecaCertificateSubjectModeEnum_value["CertificateSubjectModeEnum"+string(*e)]; ok {
		return privatecapb.PrivatecaCertificateSubjectModeEnum(v)
	}
	return privatecapb.PrivatecaCertificateSubjectModeEnum(0)
}

// CertificateRevocationDetailsRevocationStateEnumToProto converts a CertificateRevocationDetailsRevocationStateEnum enum to its proto representation.
func PrivatecaCertificateRevocationDetailsRevocationStateEnumToProto(e *privateca.CertificateRevocationDetailsRevocationStateEnum) privatecapb.PrivatecaCertificateRevocationDetailsRevocationStateEnum {
	if e == nil {
		return privatecapb.PrivatecaCertificateRevocationDetailsRevocationStateEnum(0)
	}
	if v, ok := privatecapb.PrivatecaCertificateRevocationDetailsRevocationStateEnum_value["CertificateRevocationDetailsRevocationStateEnum"+string(*e)]; ok {
		return privatecapb.PrivatecaCertificateRevocationDetailsRevocationStateEnum(v)
	}
	return privatecapb.PrivatecaCertificateRevocationDetailsRevocationStateEnum(0)
}

// CertificateCertificateDescriptionPublicKeyFormatEnumToProto converts a CertificateCertificateDescriptionPublicKeyFormatEnum enum to its proto representation.
func PrivatecaCertificateCertificateDescriptionPublicKeyFormatEnumToProto(e *privateca.CertificateCertificateDescriptionPublicKeyFormatEnum) privatecapb.PrivatecaCertificateCertificateDescriptionPublicKeyFormatEnum {
	if e == nil {
		return privatecapb.PrivatecaCertificateCertificateDescriptionPublicKeyFormatEnum(0)
	}
	if v, ok := privatecapb.PrivatecaCertificateCertificateDescriptionPublicKeyFormatEnum_value["CertificateCertificateDescriptionPublicKeyFormatEnum"+string(*e)]; ok {
		return privatecapb.PrivatecaCertificateCertificateDescriptionPublicKeyFormatEnum(v)
	}
	return privatecapb.PrivatecaCertificateCertificateDescriptionPublicKeyFormatEnum(0)
}

// CertificateConfigToProto converts a CertificateConfig object to its proto representation.
func PrivatecaCertificateConfigToProto(o *privateca.CertificateConfig) *privatecapb.PrivatecaCertificateConfig {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCertificateConfig{}
	p.SetSubjectConfig(PrivatecaCertificateConfigSubjectConfigToProto(o.SubjectConfig))
	p.SetX509Config(PrivatecaCertificateConfigX509ConfigToProto(o.X509Config))
	p.SetPublicKey(PrivatecaCertificateConfigPublicKeyToProto(o.PublicKey))
	return p
}

// CertificateConfigSubjectConfigToProto converts a CertificateConfigSubjectConfig object to its proto representation.
func PrivatecaCertificateConfigSubjectConfigToProto(o *privateca.CertificateConfigSubjectConfig) *privatecapb.PrivatecaCertificateConfigSubjectConfig {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCertificateConfigSubjectConfig{}
	p.SetSubject(PrivatecaCertificateConfigSubjectConfigSubjectToProto(o.Subject))
	p.SetSubjectAltName(PrivatecaCertificateConfigSubjectConfigSubjectAltNameToProto(o.SubjectAltName))
	return p
}

// CertificateConfigSubjectConfigSubjectToProto converts a CertificateConfigSubjectConfigSubject object to its proto representation.
func PrivatecaCertificateConfigSubjectConfigSubjectToProto(o *privateca.CertificateConfigSubjectConfigSubject) *privatecapb.PrivatecaCertificateConfigSubjectConfigSubject {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCertificateConfigSubjectConfigSubject{}
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
func PrivatecaCertificateConfigSubjectConfigSubjectAltNameToProto(o *privateca.CertificateConfigSubjectConfigSubjectAltName) *privatecapb.PrivatecaCertificateConfigSubjectConfigSubjectAltName {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCertificateConfigSubjectConfigSubjectAltName{}
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
func PrivatecaCertificateConfigX509ConfigToProto(o *privateca.CertificateConfigX509Config) *privatecapb.PrivatecaCertificateConfigX509Config {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCertificateConfigX509Config{}
	p.SetKeyUsage(PrivatecaCertificateConfigX509ConfigKeyUsageToProto(o.KeyUsage))
	p.SetCaOptions(PrivatecaCertificateConfigX509ConfigCaOptionsToProto(o.CaOptions))
	sPolicyIds := make([]*privatecapb.PrivatecaCertificateConfigX509ConfigPolicyIds, len(o.PolicyIds))
	for i, r := range o.PolicyIds {
		sPolicyIds[i] = PrivatecaCertificateConfigX509ConfigPolicyIdsToProto(&r)
	}
	p.SetPolicyIds(sPolicyIds)
	sAiaOcspServers := make([]string, len(o.AiaOcspServers))
	for i, r := range o.AiaOcspServers {
		sAiaOcspServers[i] = r
	}
	p.SetAiaOcspServers(sAiaOcspServers)
	sAdditionalExtensions := make([]*privatecapb.PrivatecaCertificateConfigX509ConfigAdditionalExtensions, len(o.AdditionalExtensions))
	for i, r := range o.AdditionalExtensions {
		sAdditionalExtensions[i] = PrivatecaCertificateConfigX509ConfigAdditionalExtensionsToProto(&r)
	}
	p.SetAdditionalExtensions(sAdditionalExtensions)
	return p
}

// CertificateConfigX509ConfigKeyUsageToProto converts a CertificateConfigX509ConfigKeyUsage object to its proto representation.
func PrivatecaCertificateConfigX509ConfigKeyUsageToProto(o *privateca.CertificateConfigX509ConfigKeyUsage) *privatecapb.PrivatecaCertificateConfigX509ConfigKeyUsage {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCertificateConfigX509ConfigKeyUsage{}
	p.SetBaseKeyUsage(PrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsageToProto(o.BaseKeyUsage))
	p.SetExtendedKeyUsage(PrivatecaCertificateConfigX509ConfigKeyUsageExtendedKeyUsageToProto(o.ExtendedKeyUsage))
	sUnknownExtendedKeyUsages := make([]*privatecapb.PrivatecaCertificateConfigX509ConfigKeyUsageUnknownExtendedKeyUsages, len(o.UnknownExtendedKeyUsages))
	for i, r := range o.UnknownExtendedKeyUsages {
		sUnknownExtendedKeyUsages[i] = PrivatecaCertificateConfigX509ConfigKeyUsageUnknownExtendedKeyUsagesToProto(&r)
	}
	p.SetUnknownExtendedKeyUsages(sUnknownExtendedKeyUsages)
	return p
}

// CertificateConfigX509ConfigKeyUsageBaseKeyUsageToProto converts a CertificateConfigX509ConfigKeyUsageBaseKeyUsage object to its proto representation.
func PrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsageToProto(o *privateca.CertificateConfigX509ConfigKeyUsageBaseKeyUsage) *privatecapb.PrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsage {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsage{}
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
func PrivatecaCertificateConfigX509ConfigKeyUsageExtendedKeyUsageToProto(o *privateca.CertificateConfigX509ConfigKeyUsageExtendedKeyUsage) *privatecapb.PrivatecaCertificateConfigX509ConfigKeyUsageExtendedKeyUsage {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCertificateConfigX509ConfigKeyUsageExtendedKeyUsage{}
	p.SetServerAuth(dcl.ValueOrEmptyBool(o.ServerAuth))
	p.SetClientAuth(dcl.ValueOrEmptyBool(o.ClientAuth))
	p.SetCodeSigning(dcl.ValueOrEmptyBool(o.CodeSigning))
	p.SetEmailProtection(dcl.ValueOrEmptyBool(o.EmailProtection))
	p.SetTimeStamping(dcl.ValueOrEmptyBool(o.TimeStamping))
	p.SetOcspSigning(dcl.ValueOrEmptyBool(o.OcspSigning))
	return p
}

// CertificateConfigX509ConfigKeyUsageUnknownExtendedKeyUsagesToProto converts a CertificateConfigX509ConfigKeyUsageUnknownExtendedKeyUsages object to its proto representation.
func PrivatecaCertificateConfigX509ConfigKeyUsageUnknownExtendedKeyUsagesToProto(o *privateca.CertificateConfigX509ConfigKeyUsageUnknownExtendedKeyUsages) *privatecapb.PrivatecaCertificateConfigX509ConfigKeyUsageUnknownExtendedKeyUsages {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCertificateConfigX509ConfigKeyUsageUnknownExtendedKeyUsages{}
	sObjectIdPath := make([]int64, len(o.ObjectIdPath))
	for i, r := range o.ObjectIdPath {
		sObjectIdPath[i] = r
	}
	p.SetObjectIdPath(sObjectIdPath)
	return p
}

// CertificateConfigX509ConfigCaOptionsToProto converts a CertificateConfigX509ConfigCaOptions object to its proto representation.
func PrivatecaCertificateConfigX509ConfigCaOptionsToProto(o *privateca.CertificateConfigX509ConfigCaOptions) *privatecapb.PrivatecaCertificateConfigX509ConfigCaOptions {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCertificateConfigX509ConfigCaOptions{}
	p.SetIsCa(dcl.ValueOrEmptyBool(o.IsCa))
	p.SetNonCa(dcl.ValueOrEmptyBool(o.NonCa))
	p.SetMaxIssuerPathLength(dcl.ValueOrEmptyInt64(o.MaxIssuerPathLength))
	p.SetZeroMaxIssuerPathLength(dcl.ValueOrEmptyBool(o.ZeroMaxIssuerPathLength))
	return p
}

// CertificateConfigX509ConfigPolicyIdsToProto converts a CertificateConfigX509ConfigPolicyIds object to its proto representation.
func PrivatecaCertificateConfigX509ConfigPolicyIdsToProto(o *privateca.CertificateConfigX509ConfigPolicyIds) *privatecapb.PrivatecaCertificateConfigX509ConfigPolicyIds {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCertificateConfigX509ConfigPolicyIds{}
	sObjectIdPath := make([]int64, len(o.ObjectIdPath))
	for i, r := range o.ObjectIdPath {
		sObjectIdPath[i] = r
	}
	p.SetObjectIdPath(sObjectIdPath)
	return p
}

// CertificateConfigX509ConfigAdditionalExtensionsToProto converts a CertificateConfigX509ConfigAdditionalExtensions object to its proto representation.
func PrivatecaCertificateConfigX509ConfigAdditionalExtensionsToProto(o *privateca.CertificateConfigX509ConfigAdditionalExtensions) *privatecapb.PrivatecaCertificateConfigX509ConfigAdditionalExtensions {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCertificateConfigX509ConfigAdditionalExtensions{}
	p.SetObjectId(PrivatecaCertificateConfigX509ConfigAdditionalExtensionsObjectIdToProto(o.ObjectId))
	p.SetCritical(dcl.ValueOrEmptyBool(o.Critical))
	p.SetValue(dcl.ValueOrEmptyString(o.Value))
	return p
}

// CertificateConfigX509ConfigAdditionalExtensionsObjectIdToProto converts a CertificateConfigX509ConfigAdditionalExtensionsObjectId object to its proto representation.
func PrivatecaCertificateConfigX509ConfigAdditionalExtensionsObjectIdToProto(o *privateca.CertificateConfigX509ConfigAdditionalExtensionsObjectId) *privatecapb.PrivatecaCertificateConfigX509ConfigAdditionalExtensionsObjectId {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCertificateConfigX509ConfigAdditionalExtensionsObjectId{}
	sObjectIdPath := make([]int64, len(o.ObjectIdPath))
	for i, r := range o.ObjectIdPath {
		sObjectIdPath[i] = r
	}
	p.SetObjectIdPath(sObjectIdPath)
	return p
}

// CertificateConfigPublicKeyToProto converts a CertificateConfigPublicKey object to its proto representation.
func PrivatecaCertificateConfigPublicKeyToProto(o *privateca.CertificateConfigPublicKey) *privatecapb.PrivatecaCertificateConfigPublicKey {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCertificateConfigPublicKey{}
	p.SetKey(dcl.ValueOrEmptyString(o.Key))
	p.SetFormat(PrivatecaCertificateConfigPublicKeyFormatEnumToProto(o.Format))
	return p
}

// CertificateRevocationDetailsToProto converts a CertificateRevocationDetails object to its proto representation.
func PrivatecaCertificateRevocationDetailsToProto(o *privateca.CertificateRevocationDetails) *privatecapb.PrivatecaCertificateRevocationDetails {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCertificateRevocationDetails{}
	p.SetRevocationState(PrivatecaCertificateRevocationDetailsRevocationStateEnumToProto(o.RevocationState))
	p.SetRevocationTime(dcl.ValueOrEmptyString(o.RevocationTime))
	return p
}

// CertificateCertificateDescriptionToProto converts a CertificateCertificateDescription object to its proto representation.
func PrivatecaCertificateCertificateDescriptionToProto(o *privateca.CertificateCertificateDescription) *privatecapb.PrivatecaCertificateCertificateDescription {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCertificateCertificateDescription{}
	p.SetSubjectDescription(PrivatecaCertificateCertificateDescriptionSubjectDescriptionToProto(o.SubjectDescription))
	p.SetX509Description(PrivatecaCertificateCertificateDescriptionX509DescriptionToProto(o.X509Description))
	p.SetPublicKey(PrivatecaCertificateCertificateDescriptionPublicKeyToProto(o.PublicKey))
	p.SetSubjectKeyId(PrivatecaCertificateCertificateDescriptionSubjectKeyIdToProto(o.SubjectKeyId))
	p.SetAuthorityKeyId(PrivatecaCertificateCertificateDescriptionAuthorityKeyIdToProto(o.AuthorityKeyId))
	p.SetCertFingerprint(PrivatecaCertificateCertificateDescriptionCertFingerprintToProto(o.CertFingerprint))
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
func PrivatecaCertificateCertificateDescriptionSubjectDescriptionToProto(o *privateca.CertificateCertificateDescriptionSubjectDescription) *privatecapb.PrivatecaCertificateCertificateDescriptionSubjectDescription {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCertificateCertificateDescriptionSubjectDescription{}
	p.SetSubject(PrivatecaCertificateCertificateDescriptionSubjectDescriptionSubjectToProto(o.Subject))
	p.SetSubjectAltName(PrivatecaCertificateCertificateDescriptionSubjectDescriptionSubjectAltNameToProto(o.SubjectAltName))
	p.SetHexSerialNumber(dcl.ValueOrEmptyString(o.HexSerialNumber))
	p.SetLifetime(dcl.ValueOrEmptyString(o.Lifetime))
	p.SetNotBeforeTime(dcl.ValueOrEmptyString(o.NotBeforeTime))
	p.SetNotAfterTime(dcl.ValueOrEmptyString(o.NotAfterTime))
	return p
}

// CertificateCertificateDescriptionSubjectDescriptionSubjectToProto converts a CertificateCertificateDescriptionSubjectDescriptionSubject object to its proto representation.
func PrivatecaCertificateCertificateDescriptionSubjectDescriptionSubjectToProto(o *privateca.CertificateCertificateDescriptionSubjectDescriptionSubject) *privatecapb.PrivatecaCertificateCertificateDescriptionSubjectDescriptionSubject {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCertificateCertificateDescriptionSubjectDescriptionSubject{}
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
func PrivatecaCertificateCertificateDescriptionSubjectDescriptionSubjectAltNameToProto(o *privateca.CertificateCertificateDescriptionSubjectDescriptionSubjectAltName) *privatecapb.PrivatecaCertificateCertificateDescriptionSubjectDescriptionSubjectAltName {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCertificateCertificateDescriptionSubjectDescriptionSubjectAltName{}
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
	sCustomSans := make([]*privatecapb.PrivatecaCertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSans, len(o.CustomSans))
	for i, r := range o.CustomSans {
		sCustomSans[i] = PrivatecaCertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansToProto(&r)
	}
	p.SetCustomSans(sCustomSans)
	return p
}

// CertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansToProto converts a CertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSans object to its proto representation.
func PrivatecaCertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansToProto(o *privateca.CertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSans) *privatecapb.PrivatecaCertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSans {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSans{}
	p.SetObjectId(PrivatecaCertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansObjectIdToProto(o.ObjectId))
	p.SetCritical(dcl.ValueOrEmptyBool(o.Critical))
	p.SetValue(dcl.ValueOrEmptyString(o.Value))
	return p
}

// CertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansObjectIdToProto converts a CertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansObjectId object to its proto representation.
func PrivatecaCertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansObjectIdToProto(o *privateca.CertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansObjectId) *privatecapb.PrivatecaCertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansObjectId {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansObjectId{}
	sObjectIdPath := make([]int64, len(o.ObjectIdPath))
	for i, r := range o.ObjectIdPath {
		sObjectIdPath[i] = r
	}
	p.SetObjectIdPath(sObjectIdPath)
	return p
}

// CertificateCertificateDescriptionX509DescriptionToProto converts a CertificateCertificateDescriptionX509Description object to its proto representation.
func PrivatecaCertificateCertificateDescriptionX509DescriptionToProto(o *privateca.CertificateCertificateDescriptionX509Description) *privatecapb.PrivatecaCertificateCertificateDescriptionX509Description {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCertificateCertificateDescriptionX509Description{}
	p.SetKeyUsage(PrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageToProto(o.KeyUsage))
	p.SetCaOptions(PrivatecaCertificateCertificateDescriptionX509DescriptionCaOptionsToProto(o.CaOptions))
	sPolicyIds := make([]*privatecapb.PrivatecaCertificateCertificateDescriptionX509DescriptionPolicyIds, len(o.PolicyIds))
	for i, r := range o.PolicyIds {
		sPolicyIds[i] = PrivatecaCertificateCertificateDescriptionX509DescriptionPolicyIdsToProto(&r)
	}
	p.SetPolicyIds(sPolicyIds)
	sAiaOcspServers := make([]string, len(o.AiaOcspServers))
	for i, r := range o.AiaOcspServers {
		sAiaOcspServers[i] = r
	}
	p.SetAiaOcspServers(sAiaOcspServers)
	sAdditionalExtensions := make([]*privatecapb.PrivatecaCertificateCertificateDescriptionX509DescriptionAdditionalExtensions, len(o.AdditionalExtensions))
	for i, r := range o.AdditionalExtensions {
		sAdditionalExtensions[i] = PrivatecaCertificateCertificateDescriptionX509DescriptionAdditionalExtensionsToProto(&r)
	}
	p.SetAdditionalExtensions(sAdditionalExtensions)
	return p
}

// CertificateCertificateDescriptionX509DescriptionKeyUsageToProto converts a CertificateCertificateDescriptionX509DescriptionKeyUsage object to its proto representation.
func PrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageToProto(o *privateca.CertificateCertificateDescriptionX509DescriptionKeyUsage) *privatecapb.PrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsage {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsage{}
	p.SetBaseKeyUsage(PrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsageToProto(o.BaseKeyUsage))
	p.SetExtendedKeyUsage(PrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsageToProto(o.ExtendedKeyUsage))
	sUnknownExtendedKeyUsages := make([]*privatecapb.PrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageUnknownExtendedKeyUsages, len(o.UnknownExtendedKeyUsages))
	for i, r := range o.UnknownExtendedKeyUsages {
		sUnknownExtendedKeyUsages[i] = PrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageUnknownExtendedKeyUsagesToProto(&r)
	}
	p.SetUnknownExtendedKeyUsages(sUnknownExtendedKeyUsages)
	return p
}

// CertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsageToProto converts a CertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsage object to its proto representation.
func PrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsageToProto(o *privateca.CertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsage) *privatecapb.PrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsage {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsage{}
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
func PrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsageToProto(o *privateca.CertificateCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsage) *privatecapb.PrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsage {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsage{}
	p.SetServerAuth(dcl.ValueOrEmptyBool(o.ServerAuth))
	p.SetClientAuth(dcl.ValueOrEmptyBool(o.ClientAuth))
	p.SetCodeSigning(dcl.ValueOrEmptyBool(o.CodeSigning))
	p.SetEmailProtection(dcl.ValueOrEmptyBool(o.EmailProtection))
	p.SetTimeStamping(dcl.ValueOrEmptyBool(o.TimeStamping))
	p.SetOcspSigning(dcl.ValueOrEmptyBool(o.OcspSigning))
	return p
}

// CertificateCertificateDescriptionX509DescriptionKeyUsageUnknownExtendedKeyUsagesToProto converts a CertificateCertificateDescriptionX509DescriptionKeyUsageUnknownExtendedKeyUsages object to its proto representation.
func PrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageUnknownExtendedKeyUsagesToProto(o *privateca.CertificateCertificateDescriptionX509DescriptionKeyUsageUnknownExtendedKeyUsages) *privatecapb.PrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageUnknownExtendedKeyUsages {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageUnknownExtendedKeyUsages{}
	sObjectIdPath := make([]int64, len(o.ObjectIdPath))
	for i, r := range o.ObjectIdPath {
		sObjectIdPath[i] = r
	}
	p.SetObjectIdPath(sObjectIdPath)
	return p
}

// CertificateCertificateDescriptionX509DescriptionCaOptionsToProto converts a CertificateCertificateDescriptionX509DescriptionCaOptions object to its proto representation.
func PrivatecaCertificateCertificateDescriptionX509DescriptionCaOptionsToProto(o *privateca.CertificateCertificateDescriptionX509DescriptionCaOptions) *privatecapb.PrivatecaCertificateCertificateDescriptionX509DescriptionCaOptions {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCertificateCertificateDescriptionX509DescriptionCaOptions{}
	p.SetIsCa(dcl.ValueOrEmptyBool(o.IsCa))
	p.SetMaxIssuerPathLength(dcl.ValueOrEmptyInt64(o.MaxIssuerPathLength))
	return p
}

// CertificateCertificateDescriptionX509DescriptionPolicyIdsToProto converts a CertificateCertificateDescriptionX509DescriptionPolicyIds object to its proto representation.
func PrivatecaCertificateCertificateDescriptionX509DescriptionPolicyIdsToProto(o *privateca.CertificateCertificateDescriptionX509DescriptionPolicyIds) *privatecapb.PrivatecaCertificateCertificateDescriptionX509DescriptionPolicyIds {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCertificateCertificateDescriptionX509DescriptionPolicyIds{}
	sObjectIdPath := make([]int64, len(o.ObjectIdPath))
	for i, r := range o.ObjectIdPath {
		sObjectIdPath[i] = r
	}
	p.SetObjectIdPath(sObjectIdPath)
	return p
}

// CertificateCertificateDescriptionX509DescriptionAdditionalExtensionsToProto converts a CertificateCertificateDescriptionX509DescriptionAdditionalExtensions object to its proto representation.
func PrivatecaCertificateCertificateDescriptionX509DescriptionAdditionalExtensionsToProto(o *privateca.CertificateCertificateDescriptionX509DescriptionAdditionalExtensions) *privatecapb.PrivatecaCertificateCertificateDescriptionX509DescriptionAdditionalExtensions {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCertificateCertificateDescriptionX509DescriptionAdditionalExtensions{}
	p.SetObjectId(PrivatecaCertificateCertificateDescriptionX509DescriptionAdditionalExtensionsObjectIdToProto(o.ObjectId))
	p.SetCritical(dcl.ValueOrEmptyBool(o.Critical))
	p.SetValue(dcl.ValueOrEmptyString(o.Value))
	return p
}

// CertificateCertificateDescriptionX509DescriptionAdditionalExtensionsObjectIdToProto converts a CertificateCertificateDescriptionX509DescriptionAdditionalExtensionsObjectId object to its proto representation.
func PrivatecaCertificateCertificateDescriptionX509DescriptionAdditionalExtensionsObjectIdToProto(o *privateca.CertificateCertificateDescriptionX509DescriptionAdditionalExtensionsObjectId) *privatecapb.PrivatecaCertificateCertificateDescriptionX509DescriptionAdditionalExtensionsObjectId {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCertificateCertificateDescriptionX509DescriptionAdditionalExtensionsObjectId{}
	sObjectIdPath := make([]int64, len(o.ObjectIdPath))
	for i, r := range o.ObjectIdPath {
		sObjectIdPath[i] = r
	}
	p.SetObjectIdPath(sObjectIdPath)
	return p
}

// CertificateCertificateDescriptionPublicKeyToProto converts a CertificateCertificateDescriptionPublicKey object to its proto representation.
func PrivatecaCertificateCertificateDescriptionPublicKeyToProto(o *privateca.CertificateCertificateDescriptionPublicKey) *privatecapb.PrivatecaCertificateCertificateDescriptionPublicKey {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCertificateCertificateDescriptionPublicKey{}
	p.SetKey(dcl.ValueOrEmptyString(o.Key))
	p.SetFormat(PrivatecaCertificateCertificateDescriptionPublicKeyFormatEnumToProto(o.Format))
	return p
}

// CertificateCertificateDescriptionSubjectKeyIdToProto converts a CertificateCertificateDescriptionSubjectKeyId object to its proto representation.
func PrivatecaCertificateCertificateDescriptionSubjectKeyIdToProto(o *privateca.CertificateCertificateDescriptionSubjectKeyId) *privatecapb.PrivatecaCertificateCertificateDescriptionSubjectKeyId {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCertificateCertificateDescriptionSubjectKeyId{}
	p.SetKeyId(dcl.ValueOrEmptyString(o.KeyId))
	return p
}

// CertificateCertificateDescriptionAuthorityKeyIdToProto converts a CertificateCertificateDescriptionAuthorityKeyId object to its proto representation.
func PrivatecaCertificateCertificateDescriptionAuthorityKeyIdToProto(o *privateca.CertificateCertificateDescriptionAuthorityKeyId) *privatecapb.PrivatecaCertificateCertificateDescriptionAuthorityKeyId {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCertificateCertificateDescriptionAuthorityKeyId{}
	p.SetKeyId(dcl.ValueOrEmptyString(o.KeyId))
	return p
}

// CertificateCertificateDescriptionCertFingerprintToProto converts a CertificateCertificateDescriptionCertFingerprint object to its proto representation.
func PrivatecaCertificateCertificateDescriptionCertFingerprintToProto(o *privateca.CertificateCertificateDescriptionCertFingerprint) *privatecapb.PrivatecaCertificateCertificateDescriptionCertFingerprint {
	if o == nil {
		return nil
	}
	p := &privatecapb.PrivatecaCertificateCertificateDescriptionCertFingerprint{}
	p.SetSha256Hash(dcl.ValueOrEmptyString(o.Sha256Hash))
	return p
}

// CertificateToProto converts a Certificate resource to its proto representation.
func CertificateToProto(resource *privateca.Certificate) *privatecapb.PrivatecaCertificate {
	p := &privatecapb.PrivatecaCertificate{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetPemCsr(dcl.ValueOrEmptyString(resource.PemCsr))
	p.SetConfig(PrivatecaCertificateConfigToProto(resource.Config))
	p.SetIssuerCertificateAuthority(dcl.ValueOrEmptyString(resource.IssuerCertificateAuthority))
	p.SetLifetime(dcl.ValueOrEmptyString(resource.Lifetime))
	p.SetCertificateTemplate(dcl.ValueOrEmptyString(resource.CertificateTemplate))
	p.SetSubjectMode(PrivatecaCertificateSubjectModeEnumToProto(resource.SubjectMode))
	p.SetRevocationDetails(PrivatecaCertificateRevocationDetailsToProto(resource.RevocationDetails))
	p.SetPemCertificate(dcl.ValueOrEmptyString(resource.PemCertificate))
	p.SetCertificateDescription(PrivatecaCertificateCertificateDescriptionToProto(resource.CertificateDescription))
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
func (s *CertificateServer) applyCertificate(ctx context.Context, c *privateca.Client, request *privatecapb.ApplyPrivatecaCertificateRequest) (*privatecapb.PrivatecaCertificate, error) {
	p := ProtoToCertificate(request.GetResource())
	res, err := c.ApplyCertificate(ctx, p)
	if err != nil {
		return nil, err
	}
	r := CertificateToProto(res)
	return r, nil
}

// applyPrivatecaCertificate handles the gRPC request by passing it to the underlying Certificate Apply() method.
func (s *CertificateServer) ApplyPrivatecaCertificate(ctx context.Context, request *privatecapb.ApplyPrivatecaCertificateRequest) (*privatecapb.PrivatecaCertificate, error) {
	cl, err := createConfigCertificate(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyCertificate(ctx, cl, request)
}

// DeleteCertificate handles the gRPC request by passing it to the underlying Certificate Delete() method.
func (s *CertificateServer) DeletePrivatecaCertificate(ctx context.Context, request *privatecapb.DeletePrivatecaCertificateRequest) (*emptypb.Empty, error) {

	cl, err := createConfigCertificate(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteCertificate(ctx, ProtoToCertificate(request.GetResource()))

}

// ListPrivatecaCertificate handles the gRPC request by passing it to the underlying CertificateList() method.
func (s *CertificateServer) ListPrivatecaCertificate(ctx context.Context, request *privatecapb.ListPrivatecaCertificateRequest) (*privatecapb.ListPrivatecaCertificateResponse, error) {
	cl, err := createConfigCertificate(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListCertificate(ctx, request.GetProject(), request.GetLocation(), request.GetCaPool())
	if err != nil {
		return nil, err
	}
	var protos []*privatecapb.PrivatecaCertificate
	for _, r := range resources.Items {
		rp := CertificateToProto(r)
		protos = append(protos, rp)
	}
	p := &privatecapb.ListPrivatecaCertificateResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigCertificate(ctx context.Context, service_account_file string) (*privateca.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return privateca.NewClient(conf), nil
}
