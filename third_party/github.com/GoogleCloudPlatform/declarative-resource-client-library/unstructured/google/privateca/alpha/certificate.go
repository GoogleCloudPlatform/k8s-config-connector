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
package privateca

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/privateca/alpha"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
)

type Certificate struct{}

func CertificateToUnstructured(r *dclService.Certificate) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "privateca",
			Version: "alpha",
			Type:    "Certificate",
		},
		Object: make(map[string]interface{}),
	}
	if r.CaPool != nil {
		u.Object["caPool"] = *r.CaPool
	}
	if r.CertificateAuthority != nil {
		u.Object["certificateAuthority"] = *r.CertificateAuthority
	}
	if r.CertificateDescription != nil && r.CertificateDescription != dclService.EmptyCertificateCertificateDescription {
		rCertificateDescription := make(map[string]interface{})
		var rCertificateDescriptionAiaIssuingCertificateUrls []interface{}
		for _, rCertificateDescriptionAiaIssuingCertificateUrlsVal := range r.CertificateDescription.AiaIssuingCertificateUrls {
			rCertificateDescriptionAiaIssuingCertificateUrls = append(rCertificateDescriptionAiaIssuingCertificateUrls, rCertificateDescriptionAiaIssuingCertificateUrlsVal)
		}
		rCertificateDescription["aiaIssuingCertificateUrls"] = rCertificateDescriptionAiaIssuingCertificateUrls
		if r.CertificateDescription.AuthorityKeyId != nil && r.CertificateDescription.AuthorityKeyId != dclService.EmptyCertificateCertificateDescriptionAuthorityKeyId {
			rCertificateDescriptionAuthorityKeyId := make(map[string]interface{})
			if r.CertificateDescription.AuthorityKeyId.KeyId != nil {
				rCertificateDescriptionAuthorityKeyId["keyId"] = *r.CertificateDescription.AuthorityKeyId.KeyId
			}
			rCertificateDescription["authorityKeyId"] = rCertificateDescriptionAuthorityKeyId
		}
		if r.CertificateDescription.CertFingerprint != nil && r.CertificateDescription.CertFingerprint != dclService.EmptyCertificateCertificateDescriptionCertFingerprint {
			rCertificateDescriptionCertFingerprint := make(map[string]interface{})
			if r.CertificateDescription.CertFingerprint.Sha256Hash != nil {
				rCertificateDescriptionCertFingerprint["sha256Hash"] = *r.CertificateDescription.CertFingerprint.Sha256Hash
			}
			rCertificateDescription["certFingerprint"] = rCertificateDescriptionCertFingerprint
		}
		var rCertificateDescriptionCrlDistributionPoints []interface{}
		for _, rCertificateDescriptionCrlDistributionPointsVal := range r.CertificateDescription.CrlDistributionPoints {
			rCertificateDescriptionCrlDistributionPoints = append(rCertificateDescriptionCrlDistributionPoints, rCertificateDescriptionCrlDistributionPointsVal)
		}
		rCertificateDescription["crlDistributionPoints"] = rCertificateDescriptionCrlDistributionPoints
		if r.CertificateDescription.PublicKey != nil && r.CertificateDescription.PublicKey != dclService.EmptyCertificateCertificateDescriptionPublicKey {
			rCertificateDescriptionPublicKey := make(map[string]interface{})
			if r.CertificateDescription.PublicKey.Format != nil {
				rCertificateDescriptionPublicKey["format"] = string(*r.CertificateDescription.PublicKey.Format)
			}
			if r.CertificateDescription.PublicKey.Key != nil {
				rCertificateDescriptionPublicKey["key"] = *r.CertificateDescription.PublicKey.Key
			}
			rCertificateDescription["publicKey"] = rCertificateDescriptionPublicKey
		}
		if r.CertificateDescription.SubjectDescription != nil && r.CertificateDescription.SubjectDescription != dclService.EmptyCertificateCertificateDescriptionSubjectDescription {
			rCertificateDescriptionSubjectDescription := make(map[string]interface{})
			if r.CertificateDescription.SubjectDescription.HexSerialNumber != nil {
				rCertificateDescriptionSubjectDescription["hexSerialNumber"] = *r.CertificateDescription.SubjectDescription.HexSerialNumber
			}
			if r.CertificateDescription.SubjectDescription.Lifetime != nil {
				rCertificateDescriptionSubjectDescription["lifetime"] = *r.CertificateDescription.SubjectDescription.Lifetime
			}
			if r.CertificateDescription.SubjectDescription.NotAfterTime != nil {
				rCertificateDescriptionSubjectDescription["notAfterTime"] = *r.CertificateDescription.SubjectDescription.NotAfterTime
			}
			if r.CertificateDescription.SubjectDescription.NotBeforeTime != nil {
				rCertificateDescriptionSubjectDescription["notBeforeTime"] = *r.CertificateDescription.SubjectDescription.NotBeforeTime
			}
			if r.CertificateDescription.SubjectDescription.Subject != nil && r.CertificateDescription.SubjectDescription.Subject != dclService.EmptyCertificateCertificateDescriptionSubjectDescriptionSubject {
				rCertificateDescriptionSubjectDescriptionSubject := make(map[string]interface{})
				if r.CertificateDescription.SubjectDescription.Subject.CommonName != nil {
					rCertificateDescriptionSubjectDescriptionSubject["commonName"] = *r.CertificateDescription.SubjectDescription.Subject.CommonName
				}
				if r.CertificateDescription.SubjectDescription.Subject.CountryCode != nil {
					rCertificateDescriptionSubjectDescriptionSubject["countryCode"] = *r.CertificateDescription.SubjectDescription.Subject.CountryCode
				}
				if r.CertificateDescription.SubjectDescription.Subject.Locality != nil {
					rCertificateDescriptionSubjectDescriptionSubject["locality"] = *r.CertificateDescription.SubjectDescription.Subject.Locality
				}
				if r.CertificateDescription.SubjectDescription.Subject.Organization != nil {
					rCertificateDescriptionSubjectDescriptionSubject["organization"] = *r.CertificateDescription.SubjectDescription.Subject.Organization
				}
				if r.CertificateDescription.SubjectDescription.Subject.OrganizationalUnit != nil {
					rCertificateDescriptionSubjectDescriptionSubject["organizationalUnit"] = *r.CertificateDescription.SubjectDescription.Subject.OrganizationalUnit
				}
				if r.CertificateDescription.SubjectDescription.Subject.PostalCode != nil {
					rCertificateDescriptionSubjectDescriptionSubject["postalCode"] = *r.CertificateDescription.SubjectDescription.Subject.PostalCode
				}
				if r.CertificateDescription.SubjectDescription.Subject.Province != nil {
					rCertificateDescriptionSubjectDescriptionSubject["province"] = *r.CertificateDescription.SubjectDescription.Subject.Province
				}
				if r.CertificateDescription.SubjectDescription.Subject.StreetAddress != nil {
					rCertificateDescriptionSubjectDescriptionSubject["streetAddress"] = *r.CertificateDescription.SubjectDescription.Subject.StreetAddress
				}
				rCertificateDescriptionSubjectDescription["subject"] = rCertificateDescriptionSubjectDescriptionSubject
			}
			if r.CertificateDescription.SubjectDescription.SubjectAltName != nil && r.CertificateDescription.SubjectDescription.SubjectAltName != dclService.EmptyCertificateCertificateDescriptionSubjectDescriptionSubjectAltName {
				rCertificateDescriptionSubjectDescriptionSubjectAltName := make(map[string]interface{})
				var rCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSans []interface{}
				for _, rCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansVal := range r.CertificateDescription.SubjectDescription.SubjectAltName.CustomSans {
					rCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansObject := make(map[string]interface{})
					if rCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansVal.Critical != nil {
						rCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansObject["critical"] = *rCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansVal.Critical
					}
					if rCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansVal.ObjectId != nil && rCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansVal.ObjectId != dclService.EmptyCertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansObjectId {
						rCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansValObjectId := make(map[string]interface{})
						var rCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansValObjectIdObjectIdPath []interface{}
						for _, rCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansValObjectIdObjectIdPathVal := range rCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansVal.ObjectId.ObjectIdPath {
							rCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansValObjectIdObjectIdPath = append(rCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansValObjectIdObjectIdPath, rCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansValObjectIdObjectIdPathVal)
						}
						rCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansValObjectId["objectIdPath"] = rCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansValObjectIdObjectIdPath
						rCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansObject["objectId"] = rCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansValObjectId
					}
					if rCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansVal.Value != nil {
						rCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansObject["value"] = *rCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansVal.Value
					}
					rCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSans = append(rCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSans, rCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansObject)
				}
				rCertificateDescriptionSubjectDescriptionSubjectAltName["customSans"] = rCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSans
				var rCertificateDescriptionSubjectDescriptionSubjectAltNameDnsNames []interface{}
				for _, rCertificateDescriptionSubjectDescriptionSubjectAltNameDnsNamesVal := range r.CertificateDescription.SubjectDescription.SubjectAltName.DnsNames {
					rCertificateDescriptionSubjectDescriptionSubjectAltNameDnsNames = append(rCertificateDescriptionSubjectDescriptionSubjectAltNameDnsNames, rCertificateDescriptionSubjectDescriptionSubjectAltNameDnsNamesVal)
				}
				rCertificateDescriptionSubjectDescriptionSubjectAltName["dnsNames"] = rCertificateDescriptionSubjectDescriptionSubjectAltNameDnsNames
				var rCertificateDescriptionSubjectDescriptionSubjectAltNameEmailAddresses []interface{}
				for _, rCertificateDescriptionSubjectDescriptionSubjectAltNameEmailAddressesVal := range r.CertificateDescription.SubjectDescription.SubjectAltName.EmailAddresses {
					rCertificateDescriptionSubjectDescriptionSubjectAltNameEmailAddresses = append(rCertificateDescriptionSubjectDescriptionSubjectAltNameEmailAddresses, rCertificateDescriptionSubjectDescriptionSubjectAltNameEmailAddressesVal)
				}
				rCertificateDescriptionSubjectDescriptionSubjectAltName["emailAddresses"] = rCertificateDescriptionSubjectDescriptionSubjectAltNameEmailAddresses
				var rCertificateDescriptionSubjectDescriptionSubjectAltNameIPAddresses []interface{}
				for _, rCertificateDescriptionSubjectDescriptionSubjectAltNameIPAddressesVal := range r.CertificateDescription.SubjectDescription.SubjectAltName.IPAddresses {
					rCertificateDescriptionSubjectDescriptionSubjectAltNameIPAddresses = append(rCertificateDescriptionSubjectDescriptionSubjectAltNameIPAddresses, rCertificateDescriptionSubjectDescriptionSubjectAltNameIPAddressesVal)
				}
				rCertificateDescriptionSubjectDescriptionSubjectAltName["ipAddresses"] = rCertificateDescriptionSubjectDescriptionSubjectAltNameIPAddresses
				var rCertificateDescriptionSubjectDescriptionSubjectAltNameUris []interface{}
				for _, rCertificateDescriptionSubjectDescriptionSubjectAltNameUrisVal := range r.CertificateDescription.SubjectDescription.SubjectAltName.Uris {
					rCertificateDescriptionSubjectDescriptionSubjectAltNameUris = append(rCertificateDescriptionSubjectDescriptionSubjectAltNameUris, rCertificateDescriptionSubjectDescriptionSubjectAltNameUrisVal)
				}
				rCertificateDescriptionSubjectDescriptionSubjectAltName["uris"] = rCertificateDescriptionSubjectDescriptionSubjectAltNameUris
				rCertificateDescriptionSubjectDescription["subjectAltName"] = rCertificateDescriptionSubjectDescriptionSubjectAltName
			}
			rCertificateDescription["subjectDescription"] = rCertificateDescriptionSubjectDescription
		}
		if r.CertificateDescription.SubjectKeyId != nil && r.CertificateDescription.SubjectKeyId != dclService.EmptyCertificateCertificateDescriptionSubjectKeyId {
			rCertificateDescriptionSubjectKeyId := make(map[string]interface{})
			if r.CertificateDescription.SubjectKeyId.KeyId != nil {
				rCertificateDescriptionSubjectKeyId["keyId"] = *r.CertificateDescription.SubjectKeyId.KeyId
			}
			rCertificateDescription["subjectKeyId"] = rCertificateDescriptionSubjectKeyId
		}
		if r.CertificateDescription.X509Description != nil && r.CertificateDescription.X509Description != dclService.EmptyCertificateCertificateDescriptionX509Description {
			rCertificateDescriptionX509Description := make(map[string]interface{})
			var rCertificateDescriptionX509DescriptionAdditionalExtensions []interface{}
			for _, rCertificateDescriptionX509DescriptionAdditionalExtensionsVal := range r.CertificateDescription.X509Description.AdditionalExtensions {
				rCertificateDescriptionX509DescriptionAdditionalExtensionsObject := make(map[string]interface{})
				if rCertificateDescriptionX509DescriptionAdditionalExtensionsVal.Critical != nil {
					rCertificateDescriptionX509DescriptionAdditionalExtensionsObject["critical"] = *rCertificateDescriptionX509DescriptionAdditionalExtensionsVal.Critical
				}
				if rCertificateDescriptionX509DescriptionAdditionalExtensionsVal.ObjectId != nil && rCertificateDescriptionX509DescriptionAdditionalExtensionsVal.ObjectId != dclService.EmptyCertificateCertificateDescriptionX509DescriptionAdditionalExtensionsObjectId {
					rCertificateDescriptionX509DescriptionAdditionalExtensionsValObjectId := make(map[string]interface{})
					var rCertificateDescriptionX509DescriptionAdditionalExtensionsValObjectIdObjectIdPath []interface{}
					for _, rCertificateDescriptionX509DescriptionAdditionalExtensionsValObjectIdObjectIdPathVal := range rCertificateDescriptionX509DescriptionAdditionalExtensionsVal.ObjectId.ObjectIdPath {
						rCertificateDescriptionX509DescriptionAdditionalExtensionsValObjectIdObjectIdPath = append(rCertificateDescriptionX509DescriptionAdditionalExtensionsValObjectIdObjectIdPath, rCertificateDescriptionX509DescriptionAdditionalExtensionsValObjectIdObjectIdPathVal)
					}
					rCertificateDescriptionX509DescriptionAdditionalExtensionsValObjectId["objectIdPath"] = rCertificateDescriptionX509DescriptionAdditionalExtensionsValObjectIdObjectIdPath
					rCertificateDescriptionX509DescriptionAdditionalExtensionsObject["objectId"] = rCertificateDescriptionX509DescriptionAdditionalExtensionsValObjectId
				}
				if rCertificateDescriptionX509DescriptionAdditionalExtensionsVal.Value != nil {
					rCertificateDescriptionX509DescriptionAdditionalExtensionsObject["value"] = *rCertificateDescriptionX509DescriptionAdditionalExtensionsVal.Value
				}
				rCertificateDescriptionX509DescriptionAdditionalExtensions = append(rCertificateDescriptionX509DescriptionAdditionalExtensions, rCertificateDescriptionX509DescriptionAdditionalExtensionsObject)
			}
			rCertificateDescriptionX509Description["additionalExtensions"] = rCertificateDescriptionX509DescriptionAdditionalExtensions
			var rCertificateDescriptionX509DescriptionAiaOcspServers []interface{}
			for _, rCertificateDescriptionX509DescriptionAiaOcspServersVal := range r.CertificateDescription.X509Description.AiaOcspServers {
				rCertificateDescriptionX509DescriptionAiaOcspServers = append(rCertificateDescriptionX509DescriptionAiaOcspServers, rCertificateDescriptionX509DescriptionAiaOcspServersVal)
			}
			rCertificateDescriptionX509Description["aiaOcspServers"] = rCertificateDescriptionX509DescriptionAiaOcspServers
			if r.CertificateDescription.X509Description.CaOptions != nil && r.CertificateDescription.X509Description.CaOptions != dclService.EmptyCertificateCertificateDescriptionX509DescriptionCaOptions {
				rCertificateDescriptionX509DescriptionCaOptions := make(map[string]interface{})
				if r.CertificateDescription.X509Description.CaOptions.IsCa != nil {
					rCertificateDescriptionX509DescriptionCaOptions["isCa"] = *r.CertificateDescription.X509Description.CaOptions.IsCa
				}
				if r.CertificateDescription.X509Description.CaOptions.MaxIssuerPathLength != nil {
					rCertificateDescriptionX509DescriptionCaOptions["maxIssuerPathLength"] = *r.CertificateDescription.X509Description.CaOptions.MaxIssuerPathLength
				}
				rCertificateDescriptionX509Description["caOptions"] = rCertificateDescriptionX509DescriptionCaOptions
			}
			if r.CertificateDescription.X509Description.KeyUsage != nil && r.CertificateDescription.X509Description.KeyUsage != dclService.EmptyCertificateCertificateDescriptionX509DescriptionKeyUsage {
				rCertificateDescriptionX509DescriptionKeyUsage := make(map[string]interface{})
				if r.CertificateDescription.X509Description.KeyUsage.BaseKeyUsage != nil && r.CertificateDescription.X509Description.KeyUsage.BaseKeyUsage != dclService.EmptyCertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsage {
					rCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsage := make(map[string]interface{})
					if r.CertificateDescription.X509Description.KeyUsage.BaseKeyUsage.CertSign != nil {
						rCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsage["certSign"] = *r.CertificateDescription.X509Description.KeyUsage.BaseKeyUsage.CertSign
					}
					if r.CertificateDescription.X509Description.KeyUsage.BaseKeyUsage.ContentCommitment != nil {
						rCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsage["contentCommitment"] = *r.CertificateDescription.X509Description.KeyUsage.BaseKeyUsage.ContentCommitment
					}
					if r.CertificateDescription.X509Description.KeyUsage.BaseKeyUsage.CrlSign != nil {
						rCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsage["crlSign"] = *r.CertificateDescription.X509Description.KeyUsage.BaseKeyUsage.CrlSign
					}
					if r.CertificateDescription.X509Description.KeyUsage.BaseKeyUsage.DataEncipherment != nil {
						rCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsage["dataEncipherment"] = *r.CertificateDescription.X509Description.KeyUsage.BaseKeyUsage.DataEncipherment
					}
					if r.CertificateDescription.X509Description.KeyUsage.BaseKeyUsage.DecipherOnly != nil {
						rCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsage["decipherOnly"] = *r.CertificateDescription.X509Description.KeyUsage.BaseKeyUsage.DecipherOnly
					}
					if r.CertificateDescription.X509Description.KeyUsage.BaseKeyUsage.DigitalSignature != nil {
						rCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsage["digitalSignature"] = *r.CertificateDescription.X509Description.KeyUsage.BaseKeyUsage.DigitalSignature
					}
					if r.CertificateDescription.X509Description.KeyUsage.BaseKeyUsage.EncipherOnly != nil {
						rCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsage["encipherOnly"] = *r.CertificateDescription.X509Description.KeyUsage.BaseKeyUsage.EncipherOnly
					}
					if r.CertificateDescription.X509Description.KeyUsage.BaseKeyUsage.KeyAgreement != nil {
						rCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsage["keyAgreement"] = *r.CertificateDescription.X509Description.KeyUsage.BaseKeyUsage.KeyAgreement
					}
					if r.CertificateDescription.X509Description.KeyUsage.BaseKeyUsage.KeyEncipherment != nil {
						rCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsage["keyEncipherment"] = *r.CertificateDescription.X509Description.KeyUsage.BaseKeyUsage.KeyEncipherment
					}
					rCertificateDescriptionX509DescriptionKeyUsage["baseKeyUsage"] = rCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsage
				}
				if r.CertificateDescription.X509Description.KeyUsage.ExtendedKeyUsage != nil && r.CertificateDescription.X509Description.KeyUsage.ExtendedKeyUsage != dclService.EmptyCertificateCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsage {
					rCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsage := make(map[string]interface{})
					if r.CertificateDescription.X509Description.KeyUsage.ExtendedKeyUsage.ClientAuth != nil {
						rCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsage["clientAuth"] = *r.CertificateDescription.X509Description.KeyUsage.ExtendedKeyUsage.ClientAuth
					}
					if r.CertificateDescription.X509Description.KeyUsage.ExtendedKeyUsage.CodeSigning != nil {
						rCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsage["codeSigning"] = *r.CertificateDescription.X509Description.KeyUsage.ExtendedKeyUsage.CodeSigning
					}
					if r.CertificateDescription.X509Description.KeyUsage.ExtendedKeyUsage.EmailProtection != nil {
						rCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsage["emailProtection"] = *r.CertificateDescription.X509Description.KeyUsage.ExtendedKeyUsage.EmailProtection
					}
					if r.CertificateDescription.X509Description.KeyUsage.ExtendedKeyUsage.OcspSigning != nil {
						rCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsage["ocspSigning"] = *r.CertificateDescription.X509Description.KeyUsage.ExtendedKeyUsage.OcspSigning
					}
					if r.CertificateDescription.X509Description.KeyUsage.ExtendedKeyUsage.ServerAuth != nil {
						rCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsage["serverAuth"] = *r.CertificateDescription.X509Description.KeyUsage.ExtendedKeyUsage.ServerAuth
					}
					if r.CertificateDescription.X509Description.KeyUsage.ExtendedKeyUsage.TimeStamping != nil {
						rCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsage["timeStamping"] = *r.CertificateDescription.X509Description.KeyUsage.ExtendedKeyUsage.TimeStamping
					}
					rCertificateDescriptionX509DescriptionKeyUsage["extendedKeyUsage"] = rCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsage
				}
				var rCertificateDescriptionX509DescriptionKeyUsageUnknownExtendedKeyUsages []interface{}
				for _, rCertificateDescriptionX509DescriptionKeyUsageUnknownExtendedKeyUsagesVal := range r.CertificateDescription.X509Description.KeyUsage.UnknownExtendedKeyUsages {
					rCertificateDescriptionX509DescriptionKeyUsageUnknownExtendedKeyUsagesObject := make(map[string]interface{})
					var rCertificateDescriptionX509DescriptionKeyUsageUnknownExtendedKeyUsagesValObjectIdPath []interface{}
					for _, rCertificateDescriptionX509DescriptionKeyUsageUnknownExtendedKeyUsagesValObjectIdPathVal := range rCertificateDescriptionX509DescriptionKeyUsageUnknownExtendedKeyUsagesVal.ObjectIdPath {
						rCertificateDescriptionX509DescriptionKeyUsageUnknownExtendedKeyUsagesValObjectIdPath = append(rCertificateDescriptionX509DescriptionKeyUsageUnknownExtendedKeyUsagesValObjectIdPath, rCertificateDescriptionX509DescriptionKeyUsageUnknownExtendedKeyUsagesValObjectIdPathVal)
					}
					rCertificateDescriptionX509DescriptionKeyUsageUnknownExtendedKeyUsagesObject["objectIdPath"] = rCertificateDescriptionX509DescriptionKeyUsageUnknownExtendedKeyUsagesValObjectIdPath
					rCertificateDescriptionX509DescriptionKeyUsageUnknownExtendedKeyUsages = append(rCertificateDescriptionX509DescriptionKeyUsageUnknownExtendedKeyUsages, rCertificateDescriptionX509DescriptionKeyUsageUnknownExtendedKeyUsagesObject)
				}
				rCertificateDescriptionX509DescriptionKeyUsage["unknownExtendedKeyUsages"] = rCertificateDescriptionX509DescriptionKeyUsageUnknownExtendedKeyUsages
				rCertificateDescriptionX509Description["keyUsage"] = rCertificateDescriptionX509DescriptionKeyUsage
			}
			var rCertificateDescriptionX509DescriptionPolicyIds []interface{}
			for _, rCertificateDescriptionX509DescriptionPolicyIdsVal := range r.CertificateDescription.X509Description.PolicyIds {
				rCertificateDescriptionX509DescriptionPolicyIdsObject := make(map[string]interface{})
				var rCertificateDescriptionX509DescriptionPolicyIdsValObjectIdPath []interface{}
				for _, rCertificateDescriptionX509DescriptionPolicyIdsValObjectIdPathVal := range rCertificateDescriptionX509DescriptionPolicyIdsVal.ObjectIdPath {
					rCertificateDescriptionX509DescriptionPolicyIdsValObjectIdPath = append(rCertificateDescriptionX509DescriptionPolicyIdsValObjectIdPath, rCertificateDescriptionX509DescriptionPolicyIdsValObjectIdPathVal)
				}
				rCertificateDescriptionX509DescriptionPolicyIdsObject["objectIdPath"] = rCertificateDescriptionX509DescriptionPolicyIdsValObjectIdPath
				rCertificateDescriptionX509DescriptionPolicyIds = append(rCertificateDescriptionX509DescriptionPolicyIds, rCertificateDescriptionX509DescriptionPolicyIdsObject)
			}
			rCertificateDescriptionX509Description["policyIds"] = rCertificateDescriptionX509DescriptionPolicyIds
			rCertificateDescription["x509Description"] = rCertificateDescriptionX509Description
		}
		u.Object["certificateDescription"] = rCertificateDescription
	}
	if r.CertificateTemplate != nil {
		u.Object["certificateTemplate"] = *r.CertificateTemplate
	}
	if r.Config != nil && r.Config != dclService.EmptyCertificateConfig {
		rConfig := make(map[string]interface{})
		if r.Config.PublicKey != nil && r.Config.PublicKey != dclService.EmptyCertificateConfigPublicKey {
			rConfigPublicKey := make(map[string]interface{})
			if r.Config.PublicKey.Format != nil {
				rConfigPublicKey["format"] = string(*r.Config.PublicKey.Format)
			}
			if r.Config.PublicKey.Key != nil {
				rConfigPublicKey["key"] = *r.Config.PublicKey.Key
			}
			rConfig["publicKey"] = rConfigPublicKey
		}
		if r.Config.SubjectConfig != nil && r.Config.SubjectConfig != dclService.EmptyCertificateConfigSubjectConfig {
			rConfigSubjectConfig := make(map[string]interface{})
			if r.Config.SubjectConfig.Subject != nil && r.Config.SubjectConfig.Subject != dclService.EmptyCertificateConfigSubjectConfigSubject {
				rConfigSubjectConfigSubject := make(map[string]interface{})
				if r.Config.SubjectConfig.Subject.CommonName != nil {
					rConfigSubjectConfigSubject["commonName"] = *r.Config.SubjectConfig.Subject.CommonName
				}
				if r.Config.SubjectConfig.Subject.CountryCode != nil {
					rConfigSubjectConfigSubject["countryCode"] = *r.Config.SubjectConfig.Subject.CountryCode
				}
				if r.Config.SubjectConfig.Subject.Locality != nil {
					rConfigSubjectConfigSubject["locality"] = *r.Config.SubjectConfig.Subject.Locality
				}
				if r.Config.SubjectConfig.Subject.Organization != nil {
					rConfigSubjectConfigSubject["organization"] = *r.Config.SubjectConfig.Subject.Organization
				}
				if r.Config.SubjectConfig.Subject.OrganizationalUnit != nil {
					rConfigSubjectConfigSubject["organizationalUnit"] = *r.Config.SubjectConfig.Subject.OrganizationalUnit
				}
				if r.Config.SubjectConfig.Subject.PostalCode != nil {
					rConfigSubjectConfigSubject["postalCode"] = *r.Config.SubjectConfig.Subject.PostalCode
				}
				if r.Config.SubjectConfig.Subject.Province != nil {
					rConfigSubjectConfigSubject["province"] = *r.Config.SubjectConfig.Subject.Province
				}
				if r.Config.SubjectConfig.Subject.StreetAddress != nil {
					rConfigSubjectConfigSubject["streetAddress"] = *r.Config.SubjectConfig.Subject.StreetAddress
				}
				rConfigSubjectConfig["subject"] = rConfigSubjectConfigSubject
			}
			if r.Config.SubjectConfig.SubjectAltName != nil && r.Config.SubjectConfig.SubjectAltName != dclService.EmptyCertificateConfigSubjectConfigSubjectAltName {
				rConfigSubjectConfigSubjectAltName := make(map[string]interface{})
				var rConfigSubjectConfigSubjectAltNameDnsNames []interface{}
				for _, rConfigSubjectConfigSubjectAltNameDnsNamesVal := range r.Config.SubjectConfig.SubjectAltName.DnsNames {
					rConfigSubjectConfigSubjectAltNameDnsNames = append(rConfigSubjectConfigSubjectAltNameDnsNames, rConfigSubjectConfigSubjectAltNameDnsNamesVal)
				}
				rConfigSubjectConfigSubjectAltName["dnsNames"] = rConfigSubjectConfigSubjectAltNameDnsNames
				var rConfigSubjectConfigSubjectAltNameEmailAddresses []interface{}
				for _, rConfigSubjectConfigSubjectAltNameEmailAddressesVal := range r.Config.SubjectConfig.SubjectAltName.EmailAddresses {
					rConfigSubjectConfigSubjectAltNameEmailAddresses = append(rConfigSubjectConfigSubjectAltNameEmailAddresses, rConfigSubjectConfigSubjectAltNameEmailAddressesVal)
				}
				rConfigSubjectConfigSubjectAltName["emailAddresses"] = rConfigSubjectConfigSubjectAltNameEmailAddresses
				var rConfigSubjectConfigSubjectAltNameIPAddresses []interface{}
				for _, rConfigSubjectConfigSubjectAltNameIPAddressesVal := range r.Config.SubjectConfig.SubjectAltName.IPAddresses {
					rConfigSubjectConfigSubjectAltNameIPAddresses = append(rConfigSubjectConfigSubjectAltNameIPAddresses, rConfigSubjectConfigSubjectAltNameIPAddressesVal)
				}
				rConfigSubjectConfigSubjectAltName["ipAddresses"] = rConfigSubjectConfigSubjectAltNameIPAddresses
				var rConfigSubjectConfigSubjectAltNameUris []interface{}
				for _, rConfigSubjectConfigSubjectAltNameUrisVal := range r.Config.SubjectConfig.SubjectAltName.Uris {
					rConfigSubjectConfigSubjectAltNameUris = append(rConfigSubjectConfigSubjectAltNameUris, rConfigSubjectConfigSubjectAltNameUrisVal)
				}
				rConfigSubjectConfigSubjectAltName["uris"] = rConfigSubjectConfigSubjectAltNameUris
				rConfigSubjectConfig["subjectAltName"] = rConfigSubjectConfigSubjectAltName
			}
			rConfig["subjectConfig"] = rConfigSubjectConfig
		}
		if r.Config.X509Config != nil && r.Config.X509Config != dclService.EmptyCertificateConfigX509Config {
			rConfigX509Config := make(map[string]interface{})
			var rConfigX509ConfigAdditionalExtensions []interface{}
			for _, rConfigX509ConfigAdditionalExtensionsVal := range r.Config.X509Config.AdditionalExtensions {
				rConfigX509ConfigAdditionalExtensionsObject := make(map[string]interface{})
				if rConfigX509ConfigAdditionalExtensionsVal.Critical != nil {
					rConfigX509ConfigAdditionalExtensionsObject["critical"] = *rConfigX509ConfigAdditionalExtensionsVal.Critical
				}
				if rConfigX509ConfigAdditionalExtensionsVal.ObjectId != nil && rConfigX509ConfigAdditionalExtensionsVal.ObjectId != dclService.EmptyCertificateConfigX509ConfigAdditionalExtensionsObjectId {
					rConfigX509ConfigAdditionalExtensionsValObjectId := make(map[string]interface{})
					var rConfigX509ConfigAdditionalExtensionsValObjectIdObjectIdPath []interface{}
					for _, rConfigX509ConfigAdditionalExtensionsValObjectIdObjectIdPathVal := range rConfigX509ConfigAdditionalExtensionsVal.ObjectId.ObjectIdPath {
						rConfigX509ConfigAdditionalExtensionsValObjectIdObjectIdPath = append(rConfigX509ConfigAdditionalExtensionsValObjectIdObjectIdPath, rConfigX509ConfigAdditionalExtensionsValObjectIdObjectIdPathVal)
					}
					rConfigX509ConfigAdditionalExtensionsValObjectId["objectIdPath"] = rConfigX509ConfigAdditionalExtensionsValObjectIdObjectIdPath
					rConfigX509ConfigAdditionalExtensionsObject["objectId"] = rConfigX509ConfigAdditionalExtensionsValObjectId
				}
				if rConfigX509ConfigAdditionalExtensionsVal.Value != nil {
					rConfigX509ConfigAdditionalExtensionsObject["value"] = *rConfigX509ConfigAdditionalExtensionsVal.Value
				}
				rConfigX509ConfigAdditionalExtensions = append(rConfigX509ConfigAdditionalExtensions, rConfigX509ConfigAdditionalExtensionsObject)
			}
			rConfigX509Config["additionalExtensions"] = rConfigX509ConfigAdditionalExtensions
			var rConfigX509ConfigAiaOcspServers []interface{}
			for _, rConfigX509ConfigAiaOcspServersVal := range r.Config.X509Config.AiaOcspServers {
				rConfigX509ConfigAiaOcspServers = append(rConfigX509ConfigAiaOcspServers, rConfigX509ConfigAiaOcspServersVal)
			}
			rConfigX509Config["aiaOcspServers"] = rConfigX509ConfigAiaOcspServers
			if r.Config.X509Config.CaOptions != nil && r.Config.X509Config.CaOptions != dclService.EmptyCertificateConfigX509ConfigCaOptions {
				rConfigX509ConfigCaOptions := make(map[string]interface{})
				if r.Config.X509Config.CaOptions.IsCa != nil {
					rConfigX509ConfigCaOptions["isCa"] = *r.Config.X509Config.CaOptions.IsCa
				}
				if r.Config.X509Config.CaOptions.MaxIssuerPathLength != nil {
					rConfigX509ConfigCaOptions["maxIssuerPathLength"] = *r.Config.X509Config.CaOptions.MaxIssuerPathLength
				}
				if r.Config.X509Config.CaOptions.NonCa != nil {
					rConfigX509ConfigCaOptions["nonCa"] = *r.Config.X509Config.CaOptions.NonCa
				}
				if r.Config.X509Config.CaOptions.ZeroMaxIssuerPathLength != nil {
					rConfigX509ConfigCaOptions["zeroMaxIssuerPathLength"] = *r.Config.X509Config.CaOptions.ZeroMaxIssuerPathLength
				}
				rConfigX509Config["caOptions"] = rConfigX509ConfigCaOptions
			}
			if r.Config.X509Config.KeyUsage != nil && r.Config.X509Config.KeyUsage != dclService.EmptyCertificateConfigX509ConfigKeyUsage {
				rConfigX509ConfigKeyUsage := make(map[string]interface{})
				if r.Config.X509Config.KeyUsage.BaseKeyUsage != nil && r.Config.X509Config.KeyUsage.BaseKeyUsage != dclService.EmptyCertificateConfigX509ConfigKeyUsageBaseKeyUsage {
					rConfigX509ConfigKeyUsageBaseKeyUsage := make(map[string]interface{})
					if r.Config.X509Config.KeyUsage.BaseKeyUsage.CertSign != nil {
						rConfigX509ConfigKeyUsageBaseKeyUsage["certSign"] = *r.Config.X509Config.KeyUsage.BaseKeyUsage.CertSign
					}
					if r.Config.X509Config.KeyUsage.BaseKeyUsage.ContentCommitment != nil {
						rConfigX509ConfigKeyUsageBaseKeyUsage["contentCommitment"] = *r.Config.X509Config.KeyUsage.BaseKeyUsage.ContentCommitment
					}
					if r.Config.X509Config.KeyUsage.BaseKeyUsage.CrlSign != nil {
						rConfigX509ConfigKeyUsageBaseKeyUsage["crlSign"] = *r.Config.X509Config.KeyUsage.BaseKeyUsage.CrlSign
					}
					if r.Config.X509Config.KeyUsage.BaseKeyUsage.DataEncipherment != nil {
						rConfigX509ConfigKeyUsageBaseKeyUsage["dataEncipherment"] = *r.Config.X509Config.KeyUsage.BaseKeyUsage.DataEncipherment
					}
					if r.Config.X509Config.KeyUsage.BaseKeyUsage.DecipherOnly != nil {
						rConfigX509ConfigKeyUsageBaseKeyUsage["decipherOnly"] = *r.Config.X509Config.KeyUsage.BaseKeyUsage.DecipherOnly
					}
					if r.Config.X509Config.KeyUsage.BaseKeyUsage.DigitalSignature != nil {
						rConfigX509ConfigKeyUsageBaseKeyUsage["digitalSignature"] = *r.Config.X509Config.KeyUsage.BaseKeyUsage.DigitalSignature
					}
					if r.Config.X509Config.KeyUsage.BaseKeyUsage.EncipherOnly != nil {
						rConfigX509ConfigKeyUsageBaseKeyUsage["encipherOnly"] = *r.Config.X509Config.KeyUsage.BaseKeyUsage.EncipherOnly
					}
					if r.Config.X509Config.KeyUsage.BaseKeyUsage.KeyAgreement != nil {
						rConfigX509ConfigKeyUsageBaseKeyUsage["keyAgreement"] = *r.Config.X509Config.KeyUsage.BaseKeyUsage.KeyAgreement
					}
					if r.Config.X509Config.KeyUsage.BaseKeyUsage.KeyEncipherment != nil {
						rConfigX509ConfigKeyUsageBaseKeyUsage["keyEncipherment"] = *r.Config.X509Config.KeyUsage.BaseKeyUsage.KeyEncipherment
					}
					rConfigX509ConfigKeyUsage["baseKeyUsage"] = rConfigX509ConfigKeyUsageBaseKeyUsage
				}
				if r.Config.X509Config.KeyUsage.ExtendedKeyUsage != nil && r.Config.X509Config.KeyUsage.ExtendedKeyUsage != dclService.EmptyCertificateConfigX509ConfigKeyUsageExtendedKeyUsage {
					rConfigX509ConfigKeyUsageExtendedKeyUsage := make(map[string]interface{})
					if r.Config.X509Config.KeyUsage.ExtendedKeyUsage.ClientAuth != nil {
						rConfigX509ConfigKeyUsageExtendedKeyUsage["clientAuth"] = *r.Config.X509Config.KeyUsage.ExtendedKeyUsage.ClientAuth
					}
					if r.Config.X509Config.KeyUsage.ExtendedKeyUsage.CodeSigning != nil {
						rConfigX509ConfigKeyUsageExtendedKeyUsage["codeSigning"] = *r.Config.X509Config.KeyUsage.ExtendedKeyUsage.CodeSigning
					}
					if r.Config.X509Config.KeyUsage.ExtendedKeyUsage.EmailProtection != nil {
						rConfigX509ConfigKeyUsageExtendedKeyUsage["emailProtection"] = *r.Config.X509Config.KeyUsage.ExtendedKeyUsage.EmailProtection
					}
					if r.Config.X509Config.KeyUsage.ExtendedKeyUsage.OcspSigning != nil {
						rConfigX509ConfigKeyUsageExtendedKeyUsage["ocspSigning"] = *r.Config.X509Config.KeyUsage.ExtendedKeyUsage.OcspSigning
					}
					if r.Config.X509Config.KeyUsage.ExtendedKeyUsage.ServerAuth != nil {
						rConfigX509ConfigKeyUsageExtendedKeyUsage["serverAuth"] = *r.Config.X509Config.KeyUsage.ExtendedKeyUsage.ServerAuth
					}
					if r.Config.X509Config.KeyUsage.ExtendedKeyUsage.TimeStamping != nil {
						rConfigX509ConfigKeyUsageExtendedKeyUsage["timeStamping"] = *r.Config.X509Config.KeyUsage.ExtendedKeyUsage.TimeStamping
					}
					rConfigX509ConfigKeyUsage["extendedKeyUsage"] = rConfigX509ConfigKeyUsageExtendedKeyUsage
				}
				var rConfigX509ConfigKeyUsageUnknownExtendedKeyUsages []interface{}
				for _, rConfigX509ConfigKeyUsageUnknownExtendedKeyUsagesVal := range r.Config.X509Config.KeyUsage.UnknownExtendedKeyUsages {
					rConfigX509ConfigKeyUsageUnknownExtendedKeyUsagesObject := make(map[string]interface{})
					var rConfigX509ConfigKeyUsageUnknownExtendedKeyUsagesValObjectIdPath []interface{}
					for _, rConfigX509ConfigKeyUsageUnknownExtendedKeyUsagesValObjectIdPathVal := range rConfigX509ConfigKeyUsageUnknownExtendedKeyUsagesVal.ObjectIdPath {
						rConfigX509ConfigKeyUsageUnknownExtendedKeyUsagesValObjectIdPath = append(rConfigX509ConfigKeyUsageUnknownExtendedKeyUsagesValObjectIdPath, rConfigX509ConfigKeyUsageUnknownExtendedKeyUsagesValObjectIdPathVal)
					}
					rConfigX509ConfigKeyUsageUnknownExtendedKeyUsagesObject["objectIdPath"] = rConfigX509ConfigKeyUsageUnknownExtendedKeyUsagesValObjectIdPath
					rConfigX509ConfigKeyUsageUnknownExtendedKeyUsages = append(rConfigX509ConfigKeyUsageUnknownExtendedKeyUsages, rConfigX509ConfigKeyUsageUnknownExtendedKeyUsagesObject)
				}
				rConfigX509ConfigKeyUsage["unknownExtendedKeyUsages"] = rConfigX509ConfigKeyUsageUnknownExtendedKeyUsages
				rConfigX509Config["keyUsage"] = rConfigX509ConfigKeyUsage
			}
			var rConfigX509ConfigPolicyIds []interface{}
			for _, rConfigX509ConfigPolicyIdsVal := range r.Config.X509Config.PolicyIds {
				rConfigX509ConfigPolicyIdsObject := make(map[string]interface{})
				var rConfigX509ConfigPolicyIdsValObjectIdPath []interface{}
				for _, rConfigX509ConfigPolicyIdsValObjectIdPathVal := range rConfigX509ConfigPolicyIdsVal.ObjectIdPath {
					rConfigX509ConfigPolicyIdsValObjectIdPath = append(rConfigX509ConfigPolicyIdsValObjectIdPath, rConfigX509ConfigPolicyIdsValObjectIdPathVal)
				}
				rConfigX509ConfigPolicyIdsObject["objectIdPath"] = rConfigX509ConfigPolicyIdsValObjectIdPath
				rConfigX509ConfigPolicyIds = append(rConfigX509ConfigPolicyIds, rConfigX509ConfigPolicyIdsObject)
			}
			rConfigX509Config["policyIds"] = rConfigX509ConfigPolicyIds
			rConfig["x509Config"] = rConfigX509Config
		}
		u.Object["config"] = rConfig
	}
	if r.CreateTime != nil {
		u.Object["createTime"] = *r.CreateTime
	}
	if r.IssuerCertificateAuthority != nil {
		u.Object["issuerCertificateAuthority"] = *r.IssuerCertificateAuthority
	}
	if r.Labels != nil {
		rLabels := make(map[string]interface{})
		for k, v := range r.Labels {
			rLabels[k] = v
		}
		u.Object["labels"] = rLabels
	}
	if r.Lifetime != nil {
		u.Object["lifetime"] = *r.Lifetime
	}
	if r.Location != nil {
		u.Object["location"] = *r.Location
	}
	if r.Name != nil {
		u.Object["name"] = *r.Name
	}
	if r.PemCertificate != nil {
		u.Object["pemCertificate"] = *r.PemCertificate
	}
	var rPemCertificateChain []interface{}
	for _, rPemCertificateChainVal := range r.PemCertificateChain {
		rPemCertificateChain = append(rPemCertificateChain, rPemCertificateChainVal)
	}
	u.Object["pemCertificateChain"] = rPemCertificateChain
	if r.PemCsr != nil {
		u.Object["pemCsr"] = *r.PemCsr
	}
	if r.Project != nil {
		u.Object["project"] = *r.Project
	}
	if r.RevocationDetails != nil && r.RevocationDetails != dclService.EmptyCertificateRevocationDetails {
		rRevocationDetails := make(map[string]interface{})
		if r.RevocationDetails.RevocationState != nil {
			rRevocationDetails["revocationState"] = string(*r.RevocationDetails.RevocationState)
		}
		if r.RevocationDetails.RevocationTime != nil {
			rRevocationDetails["revocationTime"] = *r.RevocationDetails.RevocationTime
		}
		u.Object["revocationDetails"] = rRevocationDetails
	}
	if r.SubjectMode != nil {
		u.Object["subjectMode"] = string(*r.SubjectMode)
	}
	if r.UpdateTime != nil {
		u.Object["updateTime"] = *r.UpdateTime
	}
	return u
}

func UnstructuredToCertificate(u *unstructured.Resource) (*dclService.Certificate, error) {
	r := &dclService.Certificate{}
	if _, ok := u.Object["caPool"]; ok {
		if s, ok := u.Object["caPool"].(string); ok {
			r.CaPool = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.CaPool: expected string")
		}
	}
	if _, ok := u.Object["certificateAuthority"]; ok {
		if s, ok := u.Object["certificateAuthority"].(string); ok {
			r.CertificateAuthority = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.CertificateAuthority: expected string")
		}
	}
	if _, ok := u.Object["certificateDescription"]; ok {
		if rCertificateDescription, ok := u.Object["certificateDescription"].(map[string]interface{}); ok {
			r.CertificateDescription = &dclService.CertificateCertificateDescription{}
			if _, ok := rCertificateDescription["aiaIssuingCertificateUrls"]; ok {
				if s, ok := rCertificateDescription["aiaIssuingCertificateUrls"].([]interface{}); ok {
					for _, ss := range s {
						if strval, ok := ss.(string); ok {
							r.CertificateDescription.AiaIssuingCertificateUrls = append(r.CertificateDescription.AiaIssuingCertificateUrls, strval)
						}
					}
				} else {
					return nil, fmt.Errorf("r.CertificateDescription.AiaIssuingCertificateUrls: expected []interface{}")
				}
			}
			if _, ok := rCertificateDescription["authorityKeyId"]; ok {
				if rCertificateDescriptionAuthorityKeyId, ok := rCertificateDescription["authorityKeyId"].(map[string]interface{}); ok {
					r.CertificateDescription.AuthorityKeyId = &dclService.CertificateCertificateDescriptionAuthorityKeyId{}
					if _, ok := rCertificateDescriptionAuthorityKeyId["keyId"]; ok {
						if s, ok := rCertificateDescriptionAuthorityKeyId["keyId"].(string); ok {
							r.CertificateDescription.AuthorityKeyId.KeyId = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.CertificateDescription.AuthorityKeyId.KeyId: expected string")
						}
					}
				} else {
					return nil, fmt.Errorf("r.CertificateDescription.AuthorityKeyId: expected map[string]interface{}")
				}
			}
			if _, ok := rCertificateDescription["certFingerprint"]; ok {
				if rCertificateDescriptionCertFingerprint, ok := rCertificateDescription["certFingerprint"].(map[string]interface{}); ok {
					r.CertificateDescription.CertFingerprint = &dclService.CertificateCertificateDescriptionCertFingerprint{}
					if _, ok := rCertificateDescriptionCertFingerprint["sha256Hash"]; ok {
						if s, ok := rCertificateDescriptionCertFingerprint["sha256Hash"].(string); ok {
							r.CertificateDescription.CertFingerprint.Sha256Hash = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.CertificateDescription.CertFingerprint.Sha256Hash: expected string")
						}
					}
				} else {
					return nil, fmt.Errorf("r.CertificateDescription.CertFingerprint: expected map[string]interface{}")
				}
			}
			if _, ok := rCertificateDescription["crlDistributionPoints"]; ok {
				if s, ok := rCertificateDescription["crlDistributionPoints"].([]interface{}); ok {
					for _, ss := range s {
						if strval, ok := ss.(string); ok {
							r.CertificateDescription.CrlDistributionPoints = append(r.CertificateDescription.CrlDistributionPoints, strval)
						}
					}
				} else {
					return nil, fmt.Errorf("r.CertificateDescription.CrlDistributionPoints: expected []interface{}")
				}
			}
			if _, ok := rCertificateDescription["publicKey"]; ok {
				if rCertificateDescriptionPublicKey, ok := rCertificateDescription["publicKey"].(map[string]interface{}); ok {
					r.CertificateDescription.PublicKey = &dclService.CertificateCertificateDescriptionPublicKey{}
					if _, ok := rCertificateDescriptionPublicKey["format"]; ok {
						if s, ok := rCertificateDescriptionPublicKey["format"].(string); ok {
							r.CertificateDescription.PublicKey.Format = dclService.CertificateCertificateDescriptionPublicKeyFormatEnumRef(s)
						} else {
							return nil, fmt.Errorf("r.CertificateDescription.PublicKey.Format: expected string")
						}
					}
					if _, ok := rCertificateDescriptionPublicKey["key"]; ok {
						if s, ok := rCertificateDescriptionPublicKey["key"].(string); ok {
							r.CertificateDescription.PublicKey.Key = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.CertificateDescription.PublicKey.Key: expected string")
						}
					}
				} else {
					return nil, fmt.Errorf("r.CertificateDescription.PublicKey: expected map[string]interface{}")
				}
			}
			if _, ok := rCertificateDescription["subjectDescription"]; ok {
				if rCertificateDescriptionSubjectDescription, ok := rCertificateDescription["subjectDescription"].(map[string]interface{}); ok {
					r.CertificateDescription.SubjectDescription = &dclService.CertificateCertificateDescriptionSubjectDescription{}
					if _, ok := rCertificateDescriptionSubjectDescription["hexSerialNumber"]; ok {
						if s, ok := rCertificateDescriptionSubjectDescription["hexSerialNumber"].(string); ok {
							r.CertificateDescription.SubjectDescription.HexSerialNumber = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.CertificateDescription.SubjectDescription.HexSerialNumber: expected string")
						}
					}
					if _, ok := rCertificateDescriptionSubjectDescription["lifetime"]; ok {
						if s, ok := rCertificateDescriptionSubjectDescription["lifetime"].(string); ok {
							r.CertificateDescription.SubjectDescription.Lifetime = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.CertificateDescription.SubjectDescription.Lifetime: expected string")
						}
					}
					if _, ok := rCertificateDescriptionSubjectDescription["notAfterTime"]; ok {
						if s, ok := rCertificateDescriptionSubjectDescription["notAfterTime"].(string); ok {
							r.CertificateDescription.SubjectDescription.NotAfterTime = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.CertificateDescription.SubjectDescription.NotAfterTime: expected string")
						}
					}
					if _, ok := rCertificateDescriptionSubjectDescription["notBeforeTime"]; ok {
						if s, ok := rCertificateDescriptionSubjectDescription["notBeforeTime"].(string); ok {
							r.CertificateDescription.SubjectDescription.NotBeforeTime = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.CertificateDescription.SubjectDescription.NotBeforeTime: expected string")
						}
					}
					if _, ok := rCertificateDescriptionSubjectDescription["subject"]; ok {
						if rCertificateDescriptionSubjectDescriptionSubject, ok := rCertificateDescriptionSubjectDescription["subject"].(map[string]interface{}); ok {
							r.CertificateDescription.SubjectDescription.Subject = &dclService.CertificateCertificateDescriptionSubjectDescriptionSubject{}
							if _, ok := rCertificateDescriptionSubjectDescriptionSubject["commonName"]; ok {
								if s, ok := rCertificateDescriptionSubjectDescriptionSubject["commonName"].(string); ok {
									r.CertificateDescription.SubjectDescription.Subject.CommonName = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.CertificateDescription.SubjectDescription.Subject.CommonName: expected string")
								}
							}
							if _, ok := rCertificateDescriptionSubjectDescriptionSubject["countryCode"]; ok {
								if s, ok := rCertificateDescriptionSubjectDescriptionSubject["countryCode"].(string); ok {
									r.CertificateDescription.SubjectDescription.Subject.CountryCode = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.CertificateDescription.SubjectDescription.Subject.CountryCode: expected string")
								}
							}
							if _, ok := rCertificateDescriptionSubjectDescriptionSubject["locality"]; ok {
								if s, ok := rCertificateDescriptionSubjectDescriptionSubject["locality"].(string); ok {
									r.CertificateDescription.SubjectDescription.Subject.Locality = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.CertificateDescription.SubjectDescription.Subject.Locality: expected string")
								}
							}
							if _, ok := rCertificateDescriptionSubjectDescriptionSubject["organization"]; ok {
								if s, ok := rCertificateDescriptionSubjectDescriptionSubject["organization"].(string); ok {
									r.CertificateDescription.SubjectDescription.Subject.Organization = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.CertificateDescription.SubjectDescription.Subject.Organization: expected string")
								}
							}
							if _, ok := rCertificateDescriptionSubjectDescriptionSubject["organizationalUnit"]; ok {
								if s, ok := rCertificateDescriptionSubjectDescriptionSubject["organizationalUnit"].(string); ok {
									r.CertificateDescription.SubjectDescription.Subject.OrganizationalUnit = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.CertificateDescription.SubjectDescription.Subject.OrganizationalUnit: expected string")
								}
							}
							if _, ok := rCertificateDescriptionSubjectDescriptionSubject["postalCode"]; ok {
								if s, ok := rCertificateDescriptionSubjectDescriptionSubject["postalCode"].(string); ok {
									r.CertificateDescription.SubjectDescription.Subject.PostalCode = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.CertificateDescription.SubjectDescription.Subject.PostalCode: expected string")
								}
							}
							if _, ok := rCertificateDescriptionSubjectDescriptionSubject["province"]; ok {
								if s, ok := rCertificateDescriptionSubjectDescriptionSubject["province"].(string); ok {
									r.CertificateDescription.SubjectDescription.Subject.Province = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.CertificateDescription.SubjectDescription.Subject.Province: expected string")
								}
							}
							if _, ok := rCertificateDescriptionSubjectDescriptionSubject["streetAddress"]; ok {
								if s, ok := rCertificateDescriptionSubjectDescriptionSubject["streetAddress"].(string); ok {
									r.CertificateDescription.SubjectDescription.Subject.StreetAddress = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.CertificateDescription.SubjectDescription.Subject.StreetAddress: expected string")
								}
							}
						} else {
							return nil, fmt.Errorf("r.CertificateDescription.SubjectDescription.Subject: expected map[string]interface{}")
						}
					}
					if _, ok := rCertificateDescriptionSubjectDescription["subjectAltName"]; ok {
						if rCertificateDescriptionSubjectDescriptionSubjectAltName, ok := rCertificateDescriptionSubjectDescription["subjectAltName"].(map[string]interface{}); ok {
							r.CertificateDescription.SubjectDescription.SubjectAltName = &dclService.CertificateCertificateDescriptionSubjectDescriptionSubjectAltName{}
							if _, ok := rCertificateDescriptionSubjectDescriptionSubjectAltName["customSans"]; ok {
								if s, ok := rCertificateDescriptionSubjectDescriptionSubjectAltName["customSans"].([]interface{}); ok {
									for _, o := range s {
										if objval, ok := o.(map[string]interface{}); ok {
											var rCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSans dclService.CertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSans
											if _, ok := objval["critical"]; ok {
												if b, ok := objval["critical"].(bool); ok {
													rCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSans.Critical = dcl.Bool(b)
												} else {
													return nil, fmt.Errorf("rCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSans.Critical: expected bool")
												}
											}
											if _, ok := objval["objectId"]; ok {
												if rCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansObjectId, ok := objval["objectId"].(map[string]interface{}); ok {
													rCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSans.ObjectId = &dclService.CertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansObjectId{}
													if _, ok := rCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansObjectId["objectIdPath"]; ok {
														if s, ok := rCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansObjectId["objectIdPath"].([]interface{}); ok {
															for _, ss := range s {
																if intval, ok := ss.(int64); ok {
																	rCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSans.ObjectId.ObjectIdPath = append(rCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSans.ObjectId.ObjectIdPath, intval)
																}
															}
														} else {
															return nil, fmt.Errorf("rCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSans.ObjectId.ObjectIdPath: expected []interface{}")
														}
													}
												} else {
													return nil, fmt.Errorf("rCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSans.ObjectId: expected map[string]interface{}")
												}
											}
											if _, ok := objval["value"]; ok {
												if s, ok := objval["value"].(string); ok {
													rCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSans.Value = dcl.String(s)
												} else {
													return nil, fmt.Errorf("rCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSans.Value: expected string")
												}
											}
											r.CertificateDescription.SubjectDescription.SubjectAltName.CustomSans = append(r.CertificateDescription.SubjectDescription.SubjectAltName.CustomSans, rCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSans)
										}
									}
								} else {
									return nil, fmt.Errorf("r.CertificateDescription.SubjectDescription.SubjectAltName.CustomSans: expected []interface{}")
								}
							}
							if _, ok := rCertificateDescriptionSubjectDescriptionSubjectAltName["dnsNames"]; ok {
								if s, ok := rCertificateDescriptionSubjectDescriptionSubjectAltName["dnsNames"].([]interface{}); ok {
									for _, ss := range s {
										if strval, ok := ss.(string); ok {
											r.CertificateDescription.SubjectDescription.SubjectAltName.DnsNames = append(r.CertificateDescription.SubjectDescription.SubjectAltName.DnsNames, strval)
										}
									}
								} else {
									return nil, fmt.Errorf("r.CertificateDescription.SubjectDescription.SubjectAltName.DnsNames: expected []interface{}")
								}
							}
							if _, ok := rCertificateDescriptionSubjectDescriptionSubjectAltName["emailAddresses"]; ok {
								if s, ok := rCertificateDescriptionSubjectDescriptionSubjectAltName["emailAddresses"].([]interface{}); ok {
									for _, ss := range s {
										if strval, ok := ss.(string); ok {
											r.CertificateDescription.SubjectDescription.SubjectAltName.EmailAddresses = append(r.CertificateDescription.SubjectDescription.SubjectAltName.EmailAddresses, strval)
										}
									}
								} else {
									return nil, fmt.Errorf("r.CertificateDescription.SubjectDescription.SubjectAltName.EmailAddresses: expected []interface{}")
								}
							}
							if _, ok := rCertificateDescriptionSubjectDescriptionSubjectAltName["ipAddresses"]; ok {
								if s, ok := rCertificateDescriptionSubjectDescriptionSubjectAltName["ipAddresses"].([]interface{}); ok {
									for _, ss := range s {
										if strval, ok := ss.(string); ok {
											r.CertificateDescription.SubjectDescription.SubjectAltName.IPAddresses = append(r.CertificateDescription.SubjectDescription.SubjectAltName.IPAddresses, strval)
										}
									}
								} else {
									return nil, fmt.Errorf("r.CertificateDescription.SubjectDescription.SubjectAltName.IPAddresses: expected []interface{}")
								}
							}
							if _, ok := rCertificateDescriptionSubjectDescriptionSubjectAltName["uris"]; ok {
								if s, ok := rCertificateDescriptionSubjectDescriptionSubjectAltName["uris"].([]interface{}); ok {
									for _, ss := range s {
										if strval, ok := ss.(string); ok {
											r.CertificateDescription.SubjectDescription.SubjectAltName.Uris = append(r.CertificateDescription.SubjectDescription.SubjectAltName.Uris, strval)
										}
									}
								} else {
									return nil, fmt.Errorf("r.CertificateDescription.SubjectDescription.SubjectAltName.Uris: expected []interface{}")
								}
							}
						} else {
							return nil, fmt.Errorf("r.CertificateDescription.SubjectDescription.SubjectAltName: expected map[string]interface{}")
						}
					}
				} else {
					return nil, fmt.Errorf("r.CertificateDescription.SubjectDescription: expected map[string]interface{}")
				}
			}
			if _, ok := rCertificateDescription["subjectKeyId"]; ok {
				if rCertificateDescriptionSubjectKeyId, ok := rCertificateDescription["subjectKeyId"].(map[string]interface{}); ok {
					r.CertificateDescription.SubjectKeyId = &dclService.CertificateCertificateDescriptionSubjectKeyId{}
					if _, ok := rCertificateDescriptionSubjectKeyId["keyId"]; ok {
						if s, ok := rCertificateDescriptionSubjectKeyId["keyId"].(string); ok {
							r.CertificateDescription.SubjectKeyId.KeyId = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.CertificateDescription.SubjectKeyId.KeyId: expected string")
						}
					}
				} else {
					return nil, fmt.Errorf("r.CertificateDescription.SubjectKeyId: expected map[string]interface{}")
				}
			}
			if _, ok := rCertificateDescription["x509Description"]; ok {
				if rCertificateDescriptionX509Description, ok := rCertificateDescription["x509Description"].(map[string]interface{}); ok {
					r.CertificateDescription.X509Description = &dclService.CertificateCertificateDescriptionX509Description{}
					if _, ok := rCertificateDescriptionX509Description["additionalExtensions"]; ok {
						if s, ok := rCertificateDescriptionX509Description["additionalExtensions"].([]interface{}); ok {
							for _, o := range s {
								if objval, ok := o.(map[string]interface{}); ok {
									var rCertificateDescriptionX509DescriptionAdditionalExtensions dclService.CertificateCertificateDescriptionX509DescriptionAdditionalExtensions
									if _, ok := objval["critical"]; ok {
										if b, ok := objval["critical"].(bool); ok {
											rCertificateDescriptionX509DescriptionAdditionalExtensions.Critical = dcl.Bool(b)
										} else {
											return nil, fmt.Errorf("rCertificateDescriptionX509DescriptionAdditionalExtensions.Critical: expected bool")
										}
									}
									if _, ok := objval["objectId"]; ok {
										if rCertificateDescriptionX509DescriptionAdditionalExtensionsObjectId, ok := objval["objectId"].(map[string]interface{}); ok {
											rCertificateDescriptionX509DescriptionAdditionalExtensions.ObjectId = &dclService.CertificateCertificateDescriptionX509DescriptionAdditionalExtensionsObjectId{}
											if _, ok := rCertificateDescriptionX509DescriptionAdditionalExtensionsObjectId["objectIdPath"]; ok {
												if s, ok := rCertificateDescriptionX509DescriptionAdditionalExtensionsObjectId["objectIdPath"].([]interface{}); ok {
													for _, ss := range s {
														if intval, ok := ss.(int64); ok {
															rCertificateDescriptionX509DescriptionAdditionalExtensions.ObjectId.ObjectIdPath = append(rCertificateDescriptionX509DescriptionAdditionalExtensions.ObjectId.ObjectIdPath, intval)
														}
													}
												} else {
													return nil, fmt.Errorf("rCertificateDescriptionX509DescriptionAdditionalExtensions.ObjectId.ObjectIdPath: expected []interface{}")
												}
											}
										} else {
											return nil, fmt.Errorf("rCertificateDescriptionX509DescriptionAdditionalExtensions.ObjectId: expected map[string]interface{}")
										}
									}
									if _, ok := objval["value"]; ok {
										if s, ok := objval["value"].(string); ok {
											rCertificateDescriptionX509DescriptionAdditionalExtensions.Value = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rCertificateDescriptionX509DescriptionAdditionalExtensions.Value: expected string")
										}
									}
									r.CertificateDescription.X509Description.AdditionalExtensions = append(r.CertificateDescription.X509Description.AdditionalExtensions, rCertificateDescriptionX509DescriptionAdditionalExtensions)
								}
							}
						} else {
							return nil, fmt.Errorf("r.CertificateDescription.X509Description.AdditionalExtensions: expected []interface{}")
						}
					}
					if _, ok := rCertificateDescriptionX509Description["aiaOcspServers"]; ok {
						if s, ok := rCertificateDescriptionX509Description["aiaOcspServers"].([]interface{}); ok {
							for _, ss := range s {
								if strval, ok := ss.(string); ok {
									r.CertificateDescription.X509Description.AiaOcspServers = append(r.CertificateDescription.X509Description.AiaOcspServers, strval)
								}
							}
						} else {
							return nil, fmt.Errorf("r.CertificateDescription.X509Description.AiaOcspServers: expected []interface{}")
						}
					}
					if _, ok := rCertificateDescriptionX509Description["caOptions"]; ok {
						if rCertificateDescriptionX509DescriptionCaOptions, ok := rCertificateDescriptionX509Description["caOptions"].(map[string]interface{}); ok {
							r.CertificateDescription.X509Description.CaOptions = &dclService.CertificateCertificateDescriptionX509DescriptionCaOptions{}
							if _, ok := rCertificateDescriptionX509DescriptionCaOptions["isCa"]; ok {
								if b, ok := rCertificateDescriptionX509DescriptionCaOptions["isCa"].(bool); ok {
									r.CertificateDescription.X509Description.CaOptions.IsCa = dcl.Bool(b)
								} else {
									return nil, fmt.Errorf("r.CertificateDescription.X509Description.CaOptions.IsCa: expected bool")
								}
							}
							if _, ok := rCertificateDescriptionX509DescriptionCaOptions["maxIssuerPathLength"]; ok {
								if i, ok := rCertificateDescriptionX509DescriptionCaOptions["maxIssuerPathLength"].(int64); ok {
									r.CertificateDescription.X509Description.CaOptions.MaxIssuerPathLength = dcl.Int64(i)
								} else {
									return nil, fmt.Errorf("r.CertificateDescription.X509Description.CaOptions.MaxIssuerPathLength: expected int64")
								}
							}
						} else {
							return nil, fmt.Errorf("r.CertificateDescription.X509Description.CaOptions: expected map[string]interface{}")
						}
					}
					if _, ok := rCertificateDescriptionX509Description["keyUsage"]; ok {
						if rCertificateDescriptionX509DescriptionKeyUsage, ok := rCertificateDescriptionX509Description["keyUsage"].(map[string]interface{}); ok {
							r.CertificateDescription.X509Description.KeyUsage = &dclService.CertificateCertificateDescriptionX509DescriptionKeyUsage{}
							if _, ok := rCertificateDescriptionX509DescriptionKeyUsage["baseKeyUsage"]; ok {
								if rCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsage, ok := rCertificateDescriptionX509DescriptionKeyUsage["baseKeyUsage"].(map[string]interface{}); ok {
									r.CertificateDescription.X509Description.KeyUsage.BaseKeyUsage = &dclService.CertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsage{}
									if _, ok := rCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsage["certSign"]; ok {
										if b, ok := rCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsage["certSign"].(bool); ok {
											r.CertificateDescription.X509Description.KeyUsage.BaseKeyUsage.CertSign = dcl.Bool(b)
										} else {
											return nil, fmt.Errorf("r.CertificateDescription.X509Description.KeyUsage.BaseKeyUsage.CertSign: expected bool")
										}
									}
									if _, ok := rCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsage["contentCommitment"]; ok {
										if b, ok := rCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsage["contentCommitment"].(bool); ok {
											r.CertificateDescription.X509Description.KeyUsage.BaseKeyUsage.ContentCommitment = dcl.Bool(b)
										} else {
											return nil, fmt.Errorf("r.CertificateDescription.X509Description.KeyUsage.BaseKeyUsage.ContentCommitment: expected bool")
										}
									}
									if _, ok := rCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsage["crlSign"]; ok {
										if b, ok := rCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsage["crlSign"].(bool); ok {
											r.CertificateDescription.X509Description.KeyUsage.BaseKeyUsage.CrlSign = dcl.Bool(b)
										} else {
											return nil, fmt.Errorf("r.CertificateDescription.X509Description.KeyUsage.BaseKeyUsage.CrlSign: expected bool")
										}
									}
									if _, ok := rCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsage["dataEncipherment"]; ok {
										if b, ok := rCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsage["dataEncipherment"].(bool); ok {
											r.CertificateDescription.X509Description.KeyUsage.BaseKeyUsage.DataEncipherment = dcl.Bool(b)
										} else {
											return nil, fmt.Errorf("r.CertificateDescription.X509Description.KeyUsage.BaseKeyUsage.DataEncipherment: expected bool")
										}
									}
									if _, ok := rCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsage["decipherOnly"]; ok {
										if b, ok := rCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsage["decipherOnly"].(bool); ok {
											r.CertificateDescription.X509Description.KeyUsage.BaseKeyUsage.DecipherOnly = dcl.Bool(b)
										} else {
											return nil, fmt.Errorf("r.CertificateDescription.X509Description.KeyUsage.BaseKeyUsage.DecipherOnly: expected bool")
										}
									}
									if _, ok := rCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsage["digitalSignature"]; ok {
										if b, ok := rCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsage["digitalSignature"].(bool); ok {
											r.CertificateDescription.X509Description.KeyUsage.BaseKeyUsage.DigitalSignature = dcl.Bool(b)
										} else {
											return nil, fmt.Errorf("r.CertificateDescription.X509Description.KeyUsage.BaseKeyUsage.DigitalSignature: expected bool")
										}
									}
									if _, ok := rCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsage["encipherOnly"]; ok {
										if b, ok := rCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsage["encipherOnly"].(bool); ok {
											r.CertificateDescription.X509Description.KeyUsage.BaseKeyUsage.EncipherOnly = dcl.Bool(b)
										} else {
											return nil, fmt.Errorf("r.CertificateDescription.X509Description.KeyUsage.BaseKeyUsage.EncipherOnly: expected bool")
										}
									}
									if _, ok := rCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsage["keyAgreement"]; ok {
										if b, ok := rCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsage["keyAgreement"].(bool); ok {
											r.CertificateDescription.X509Description.KeyUsage.BaseKeyUsage.KeyAgreement = dcl.Bool(b)
										} else {
											return nil, fmt.Errorf("r.CertificateDescription.X509Description.KeyUsage.BaseKeyUsage.KeyAgreement: expected bool")
										}
									}
									if _, ok := rCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsage["keyEncipherment"]; ok {
										if b, ok := rCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsage["keyEncipherment"].(bool); ok {
											r.CertificateDescription.X509Description.KeyUsage.BaseKeyUsage.KeyEncipherment = dcl.Bool(b)
										} else {
											return nil, fmt.Errorf("r.CertificateDescription.X509Description.KeyUsage.BaseKeyUsage.KeyEncipherment: expected bool")
										}
									}
								} else {
									return nil, fmt.Errorf("r.CertificateDescription.X509Description.KeyUsage.BaseKeyUsage: expected map[string]interface{}")
								}
							}
							if _, ok := rCertificateDescriptionX509DescriptionKeyUsage["extendedKeyUsage"]; ok {
								if rCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsage, ok := rCertificateDescriptionX509DescriptionKeyUsage["extendedKeyUsage"].(map[string]interface{}); ok {
									r.CertificateDescription.X509Description.KeyUsage.ExtendedKeyUsage = &dclService.CertificateCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsage{}
									if _, ok := rCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsage["clientAuth"]; ok {
										if b, ok := rCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsage["clientAuth"].(bool); ok {
											r.CertificateDescription.X509Description.KeyUsage.ExtendedKeyUsage.ClientAuth = dcl.Bool(b)
										} else {
											return nil, fmt.Errorf("r.CertificateDescription.X509Description.KeyUsage.ExtendedKeyUsage.ClientAuth: expected bool")
										}
									}
									if _, ok := rCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsage["codeSigning"]; ok {
										if b, ok := rCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsage["codeSigning"].(bool); ok {
											r.CertificateDescription.X509Description.KeyUsage.ExtendedKeyUsage.CodeSigning = dcl.Bool(b)
										} else {
											return nil, fmt.Errorf("r.CertificateDescription.X509Description.KeyUsage.ExtendedKeyUsage.CodeSigning: expected bool")
										}
									}
									if _, ok := rCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsage["emailProtection"]; ok {
										if b, ok := rCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsage["emailProtection"].(bool); ok {
											r.CertificateDescription.X509Description.KeyUsage.ExtendedKeyUsage.EmailProtection = dcl.Bool(b)
										} else {
											return nil, fmt.Errorf("r.CertificateDescription.X509Description.KeyUsage.ExtendedKeyUsage.EmailProtection: expected bool")
										}
									}
									if _, ok := rCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsage["ocspSigning"]; ok {
										if b, ok := rCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsage["ocspSigning"].(bool); ok {
											r.CertificateDescription.X509Description.KeyUsage.ExtendedKeyUsage.OcspSigning = dcl.Bool(b)
										} else {
											return nil, fmt.Errorf("r.CertificateDescription.X509Description.KeyUsage.ExtendedKeyUsage.OcspSigning: expected bool")
										}
									}
									if _, ok := rCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsage["serverAuth"]; ok {
										if b, ok := rCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsage["serverAuth"].(bool); ok {
											r.CertificateDescription.X509Description.KeyUsage.ExtendedKeyUsage.ServerAuth = dcl.Bool(b)
										} else {
											return nil, fmt.Errorf("r.CertificateDescription.X509Description.KeyUsage.ExtendedKeyUsage.ServerAuth: expected bool")
										}
									}
									if _, ok := rCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsage["timeStamping"]; ok {
										if b, ok := rCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsage["timeStamping"].(bool); ok {
											r.CertificateDescription.X509Description.KeyUsage.ExtendedKeyUsage.TimeStamping = dcl.Bool(b)
										} else {
											return nil, fmt.Errorf("r.CertificateDescription.X509Description.KeyUsage.ExtendedKeyUsage.TimeStamping: expected bool")
										}
									}
								} else {
									return nil, fmt.Errorf("r.CertificateDescription.X509Description.KeyUsage.ExtendedKeyUsage: expected map[string]interface{}")
								}
							}
							if _, ok := rCertificateDescriptionX509DescriptionKeyUsage["unknownExtendedKeyUsages"]; ok {
								if s, ok := rCertificateDescriptionX509DescriptionKeyUsage["unknownExtendedKeyUsages"].([]interface{}); ok {
									for _, o := range s {
										if objval, ok := o.(map[string]interface{}); ok {
											var rCertificateDescriptionX509DescriptionKeyUsageUnknownExtendedKeyUsages dclService.CertificateCertificateDescriptionX509DescriptionKeyUsageUnknownExtendedKeyUsages
											if _, ok := objval["objectIdPath"]; ok {
												if s, ok := objval["objectIdPath"].([]interface{}); ok {
													for _, ss := range s {
														if intval, ok := ss.(int64); ok {
															rCertificateDescriptionX509DescriptionKeyUsageUnknownExtendedKeyUsages.ObjectIdPath = append(rCertificateDescriptionX509DescriptionKeyUsageUnknownExtendedKeyUsages.ObjectIdPath, intval)
														}
													}
												} else {
													return nil, fmt.Errorf("rCertificateDescriptionX509DescriptionKeyUsageUnknownExtendedKeyUsages.ObjectIdPath: expected []interface{}")
												}
											}
											r.CertificateDescription.X509Description.KeyUsage.UnknownExtendedKeyUsages = append(r.CertificateDescription.X509Description.KeyUsage.UnknownExtendedKeyUsages, rCertificateDescriptionX509DescriptionKeyUsageUnknownExtendedKeyUsages)
										}
									}
								} else {
									return nil, fmt.Errorf("r.CertificateDescription.X509Description.KeyUsage.UnknownExtendedKeyUsages: expected []interface{}")
								}
							}
						} else {
							return nil, fmt.Errorf("r.CertificateDescription.X509Description.KeyUsage: expected map[string]interface{}")
						}
					}
					if _, ok := rCertificateDescriptionX509Description["policyIds"]; ok {
						if s, ok := rCertificateDescriptionX509Description["policyIds"].([]interface{}); ok {
							for _, o := range s {
								if objval, ok := o.(map[string]interface{}); ok {
									var rCertificateDescriptionX509DescriptionPolicyIds dclService.CertificateCertificateDescriptionX509DescriptionPolicyIds
									if _, ok := objval["objectIdPath"]; ok {
										if s, ok := objval["objectIdPath"].([]interface{}); ok {
											for _, ss := range s {
												if intval, ok := ss.(int64); ok {
													rCertificateDescriptionX509DescriptionPolicyIds.ObjectIdPath = append(rCertificateDescriptionX509DescriptionPolicyIds.ObjectIdPath, intval)
												}
											}
										} else {
											return nil, fmt.Errorf("rCertificateDescriptionX509DescriptionPolicyIds.ObjectIdPath: expected []interface{}")
										}
									}
									r.CertificateDescription.X509Description.PolicyIds = append(r.CertificateDescription.X509Description.PolicyIds, rCertificateDescriptionX509DescriptionPolicyIds)
								}
							}
						} else {
							return nil, fmt.Errorf("r.CertificateDescription.X509Description.PolicyIds: expected []interface{}")
						}
					}
				} else {
					return nil, fmt.Errorf("r.CertificateDescription.X509Description: expected map[string]interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.CertificateDescription: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["certificateTemplate"]; ok {
		if s, ok := u.Object["certificateTemplate"].(string); ok {
			r.CertificateTemplate = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.CertificateTemplate: expected string")
		}
	}
	if _, ok := u.Object["config"]; ok {
		if rConfig, ok := u.Object["config"].(map[string]interface{}); ok {
			r.Config = &dclService.CertificateConfig{}
			if _, ok := rConfig["publicKey"]; ok {
				if rConfigPublicKey, ok := rConfig["publicKey"].(map[string]interface{}); ok {
					r.Config.PublicKey = &dclService.CertificateConfigPublicKey{}
					if _, ok := rConfigPublicKey["format"]; ok {
						if s, ok := rConfigPublicKey["format"].(string); ok {
							r.Config.PublicKey.Format = dclService.CertificateConfigPublicKeyFormatEnumRef(s)
						} else {
							return nil, fmt.Errorf("r.Config.PublicKey.Format: expected string")
						}
					}
					if _, ok := rConfigPublicKey["key"]; ok {
						if s, ok := rConfigPublicKey["key"].(string); ok {
							r.Config.PublicKey.Key = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.Config.PublicKey.Key: expected string")
						}
					}
				} else {
					return nil, fmt.Errorf("r.Config.PublicKey: expected map[string]interface{}")
				}
			}
			if _, ok := rConfig["subjectConfig"]; ok {
				if rConfigSubjectConfig, ok := rConfig["subjectConfig"].(map[string]interface{}); ok {
					r.Config.SubjectConfig = &dclService.CertificateConfigSubjectConfig{}
					if _, ok := rConfigSubjectConfig["subject"]; ok {
						if rConfigSubjectConfigSubject, ok := rConfigSubjectConfig["subject"].(map[string]interface{}); ok {
							r.Config.SubjectConfig.Subject = &dclService.CertificateConfigSubjectConfigSubject{}
							if _, ok := rConfigSubjectConfigSubject["commonName"]; ok {
								if s, ok := rConfigSubjectConfigSubject["commonName"].(string); ok {
									r.Config.SubjectConfig.Subject.CommonName = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.Config.SubjectConfig.Subject.CommonName: expected string")
								}
							}
							if _, ok := rConfigSubjectConfigSubject["countryCode"]; ok {
								if s, ok := rConfigSubjectConfigSubject["countryCode"].(string); ok {
									r.Config.SubjectConfig.Subject.CountryCode = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.Config.SubjectConfig.Subject.CountryCode: expected string")
								}
							}
							if _, ok := rConfigSubjectConfigSubject["locality"]; ok {
								if s, ok := rConfigSubjectConfigSubject["locality"].(string); ok {
									r.Config.SubjectConfig.Subject.Locality = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.Config.SubjectConfig.Subject.Locality: expected string")
								}
							}
							if _, ok := rConfigSubjectConfigSubject["organization"]; ok {
								if s, ok := rConfigSubjectConfigSubject["organization"].(string); ok {
									r.Config.SubjectConfig.Subject.Organization = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.Config.SubjectConfig.Subject.Organization: expected string")
								}
							}
							if _, ok := rConfigSubjectConfigSubject["organizationalUnit"]; ok {
								if s, ok := rConfigSubjectConfigSubject["organizationalUnit"].(string); ok {
									r.Config.SubjectConfig.Subject.OrganizationalUnit = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.Config.SubjectConfig.Subject.OrganizationalUnit: expected string")
								}
							}
							if _, ok := rConfigSubjectConfigSubject["postalCode"]; ok {
								if s, ok := rConfigSubjectConfigSubject["postalCode"].(string); ok {
									r.Config.SubjectConfig.Subject.PostalCode = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.Config.SubjectConfig.Subject.PostalCode: expected string")
								}
							}
							if _, ok := rConfigSubjectConfigSubject["province"]; ok {
								if s, ok := rConfigSubjectConfigSubject["province"].(string); ok {
									r.Config.SubjectConfig.Subject.Province = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.Config.SubjectConfig.Subject.Province: expected string")
								}
							}
							if _, ok := rConfigSubjectConfigSubject["streetAddress"]; ok {
								if s, ok := rConfigSubjectConfigSubject["streetAddress"].(string); ok {
									r.Config.SubjectConfig.Subject.StreetAddress = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.Config.SubjectConfig.Subject.StreetAddress: expected string")
								}
							}
						} else {
							return nil, fmt.Errorf("r.Config.SubjectConfig.Subject: expected map[string]interface{}")
						}
					}
					if _, ok := rConfigSubjectConfig["subjectAltName"]; ok {
						if rConfigSubjectConfigSubjectAltName, ok := rConfigSubjectConfig["subjectAltName"].(map[string]interface{}); ok {
							r.Config.SubjectConfig.SubjectAltName = &dclService.CertificateConfigSubjectConfigSubjectAltName{}
							if _, ok := rConfigSubjectConfigSubjectAltName["dnsNames"]; ok {
								if s, ok := rConfigSubjectConfigSubjectAltName["dnsNames"].([]interface{}); ok {
									for _, ss := range s {
										if strval, ok := ss.(string); ok {
											r.Config.SubjectConfig.SubjectAltName.DnsNames = append(r.Config.SubjectConfig.SubjectAltName.DnsNames, strval)
										}
									}
								} else {
									return nil, fmt.Errorf("r.Config.SubjectConfig.SubjectAltName.DnsNames: expected []interface{}")
								}
							}
							if _, ok := rConfigSubjectConfigSubjectAltName["emailAddresses"]; ok {
								if s, ok := rConfigSubjectConfigSubjectAltName["emailAddresses"].([]interface{}); ok {
									for _, ss := range s {
										if strval, ok := ss.(string); ok {
											r.Config.SubjectConfig.SubjectAltName.EmailAddresses = append(r.Config.SubjectConfig.SubjectAltName.EmailAddresses, strval)
										}
									}
								} else {
									return nil, fmt.Errorf("r.Config.SubjectConfig.SubjectAltName.EmailAddresses: expected []interface{}")
								}
							}
							if _, ok := rConfigSubjectConfigSubjectAltName["ipAddresses"]; ok {
								if s, ok := rConfigSubjectConfigSubjectAltName["ipAddresses"].([]interface{}); ok {
									for _, ss := range s {
										if strval, ok := ss.(string); ok {
											r.Config.SubjectConfig.SubjectAltName.IPAddresses = append(r.Config.SubjectConfig.SubjectAltName.IPAddresses, strval)
										}
									}
								} else {
									return nil, fmt.Errorf("r.Config.SubjectConfig.SubjectAltName.IPAddresses: expected []interface{}")
								}
							}
							if _, ok := rConfigSubjectConfigSubjectAltName["uris"]; ok {
								if s, ok := rConfigSubjectConfigSubjectAltName["uris"].([]interface{}); ok {
									for _, ss := range s {
										if strval, ok := ss.(string); ok {
											r.Config.SubjectConfig.SubjectAltName.Uris = append(r.Config.SubjectConfig.SubjectAltName.Uris, strval)
										}
									}
								} else {
									return nil, fmt.Errorf("r.Config.SubjectConfig.SubjectAltName.Uris: expected []interface{}")
								}
							}
						} else {
							return nil, fmt.Errorf("r.Config.SubjectConfig.SubjectAltName: expected map[string]interface{}")
						}
					}
				} else {
					return nil, fmt.Errorf("r.Config.SubjectConfig: expected map[string]interface{}")
				}
			}
			if _, ok := rConfig["x509Config"]; ok {
				if rConfigX509Config, ok := rConfig["x509Config"].(map[string]interface{}); ok {
					r.Config.X509Config = &dclService.CertificateConfigX509Config{}
					if _, ok := rConfigX509Config["additionalExtensions"]; ok {
						if s, ok := rConfigX509Config["additionalExtensions"].([]interface{}); ok {
							for _, o := range s {
								if objval, ok := o.(map[string]interface{}); ok {
									var rConfigX509ConfigAdditionalExtensions dclService.CertificateConfigX509ConfigAdditionalExtensions
									if _, ok := objval["critical"]; ok {
										if b, ok := objval["critical"].(bool); ok {
											rConfigX509ConfigAdditionalExtensions.Critical = dcl.Bool(b)
										} else {
											return nil, fmt.Errorf("rConfigX509ConfigAdditionalExtensions.Critical: expected bool")
										}
									}
									if _, ok := objval["objectId"]; ok {
										if rConfigX509ConfigAdditionalExtensionsObjectId, ok := objval["objectId"].(map[string]interface{}); ok {
											rConfigX509ConfigAdditionalExtensions.ObjectId = &dclService.CertificateConfigX509ConfigAdditionalExtensionsObjectId{}
											if _, ok := rConfigX509ConfigAdditionalExtensionsObjectId["objectIdPath"]; ok {
												if s, ok := rConfigX509ConfigAdditionalExtensionsObjectId["objectIdPath"].([]interface{}); ok {
													for _, ss := range s {
														if intval, ok := ss.(int64); ok {
															rConfigX509ConfigAdditionalExtensions.ObjectId.ObjectIdPath = append(rConfigX509ConfigAdditionalExtensions.ObjectId.ObjectIdPath, intval)
														}
													}
												} else {
													return nil, fmt.Errorf("rConfigX509ConfigAdditionalExtensions.ObjectId.ObjectIdPath: expected []interface{}")
												}
											}
										} else {
											return nil, fmt.Errorf("rConfigX509ConfigAdditionalExtensions.ObjectId: expected map[string]interface{}")
										}
									}
									if _, ok := objval["value"]; ok {
										if s, ok := objval["value"].(string); ok {
											rConfigX509ConfigAdditionalExtensions.Value = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rConfigX509ConfigAdditionalExtensions.Value: expected string")
										}
									}
									r.Config.X509Config.AdditionalExtensions = append(r.Config.X509Config.AdditionalExtensions, rConfigX509ConfigAdditionalExtensions)
								}
							}
						} else {
							return nil, fmt.Errorf("r.Config.X509Config.AdditionalExtensions: expected []interface{}")
						}
					}
					if _, ok := rConfigX509Config["aiaOcspServers"]; ok {
						if s, ok := rConfigX509Config["aiaOcspServers"].([]interface{}); ok {
							for _, ss := range s {
								if strval, ok := ss.(string); ok {
									r.Config.X509Config.AiaOcspServers = append(r.Config.X509Config.AiaOcspServers, strval)
								}
							}
						} else {
							return nil, fmt.Errorf("r.Config.X509Config.AiaOcspServers: expected []interface{}")
						}
					}
					if _, ok := rConfigX509Config["caOptions"]; ok {
						if rConfigX509ConfigCaOptions, ok := rConfigX509Config["caOptions"].(map[string]interface{}); ok {
							r.Config.X509Config.CaOptions = &dclService.CertificateConfigX509ConfigCaOptions{}
							if _, ok := rConfigX509ConfigCaOptions["isCa"]; ok {
								if b, ok := rConfigX509ConfigCaOptions["isCa"].(bool); ok {
									r.Config.X509Config.CaOptions.IsCa = dcl.Bool(b)
								} else {
									return nil, fmt.Errorf("r.Config.X509Config.CaOptions.IsCa: expected bool")
								}
							}
							if _, ok := rConfigX509ConfigCaOptions["maxIssuerPathLength"]; ok {
								if i, ok := rConfigX509ConfigCaOptions["maxIssuerPathLength"].(int64); ok {
									r.Config.X509Config.CaOptions.MaxIssuerPathLength = dcl.Int64(i)
								} else {
									return nil, fmt.Errorf("r.Config.X509Config.CaOptions.MaxIssuerPathLength: expected int64")
								}
							}
							if _, ok := rConfigX509ConfigCaOptions["nonCa"]; ok {
								if b, ok := rConfigX509ConfigCaOptions["nonCa"].(bool); ok {
									r.Config.X509Config.CaOptions.NonCa = dcl.Bool(b)
								} else {
									return nil, fmt.Errorf("r.Config.X509Config.CaOptions.NonCa: expected bool")
								}
							}
							if _, ok := rConfigX509ConfigCaOptions["zeroMaxIssuerPathLength"]; ok {
								if b, ok := rConfigX509ConfigCaOptions["zeroMaxIssuerPathLength"].(bool); ok {
									r.Config.X509Config.CaOptions.ZeroMaxIssuerPathLength = dcl.Bool(b)
								} else {
									return nil, fmt.Errorf("r.Config.X509Config.CaOptions.ZeroMaxIssuerPathLength: expected bool")
								}
							}
						} else {
							return nil, fmt.Errorf("r.Config.X509Config.CaOptions: expected map[string]interface{}")
						}
					}
					if _, ok := rConfigX509Config["keyUsage"]; ok {
						if rConfigX509ConfigKeyUsage, ok := rConfigX509Config["keyUsage"].(map[string]interface{}); ok {
							r.Config.X509Config.KeyUsage = &dclService.CertificateConfigX509ConfigKeyUsage{}
							if _, ok := rConfigX509ConfigKeyUsage["baseKeyUsage"]; ok {
								if rConfigX509ConfigKeyUsageBaseKeyUsage, ok := rConfigX509ConfigKeyUsage["baseKeyUsage"].(map[string]interface{}); ok {
									r.Config.X509Config.KeyUsage.BaseKeyUsage = &dclService.CertificateConfigX509ConfigKeyUsageBaseKeyUsage{}
									if _, ok := rConfigX509ConfigKeyUsageBaseKeyUsage["certSign"]; ok {
										if b, ok := rConfigX509ConfigKeyUsageBaseKeyUsage["certSign"].(bool); ok {
											r.Config.X509Config.KeyUsage.BaseKeyUsage.CertSign = dcl.Bool(b)
										} else {
											return nil, fmt.Errorf("r.Config.X509Config.KeyUsage.BaseKeyUsage.CertSign: expected bool")
										}
									}
									if _, ok := rConfigX509ConfigKeyUsageBaseKeyUsage["contentCommitment"]; ok {
										if b, ok := rConfigX509ConfigKeyUsageBaseKeyUsage["contentCommitment"].(bool); ok {
											r.Config.X509Config.KeyUsage.BaseKeyUsage.ContentCommitment = dcl.Bool(b)
										} else {
											return nil, fmt.Errorf("r.Config.X509Config.KeyUsage.BaseKeyUsage.ContentCommitment: expected bool")
										}
									}
									if _, ok := rConfigX509ConfigKeyUsageBaseKeyUsage["crlSign"]; ok {
										if b, ok := rConfigX509ConfigKeyUsageBaseKeyUsage["crlSign"].(bool); ok {
											r.Config.X509Config.KeyUsage.BaseKeyUsage.CrlSign = dcl.Bool(b)
										} else {
											return nil, fmt.Errorf("r.Config.X509Config.KeyUsage.BaseKeyUsage.CrlSign: expected bool")
										}
									}
									if _, ok := rConfigX509ConfigKeyUsageBaseKeyUsage["dataEncipherment"]; ok {
										if b, ok := rConfigX509ConfigKeyUsageBaseKeyUsage["dataEncipherment"].(bool); ok {
											r.Config.X509Config.KeyUsage.BaseKeyUsage.DataEncipherment = dcl.Bool(b)
										} else {
											return nil, fmt.Errorf("r.Config.X509Config.KeyUsage.BaseKeyUsage.DataEncipherment: expected bool")
										}
									}
									if _, ok := rConfigX509ConfigKeyUsageBaseKeyUsage["decipherOnly"]; ok {
										if b, ok := rConfigX509ConfigKeyUsageBaseKeyUsage["decipherOnly"].(bool); ok {
											r.Config.X509Config.KeyUsage.BaseKeyUsage.DecipherOnly = dcl.Bool(b)
										} else {
											return nil, fmt.Errorf("r.Config.X509Config.KeyUsage.BaseKeyUsage.DecipherOnly: expected bool")
										}
									}
									if _, ok := rConfigX509ConfigKeyUsageBaseKeyUsage["digitalSignature"]; ok {
										if b, ok := rConfigX509ConfigKeyUsageBaseKeyUsage["digitalSignature"].(bool); ok {
											r.Config.X509Config.KeyUsage.BaseKeyUsage.DigitalSignature = dcl.Bool(b)
										} else {
											return nil, fmt.Errorf("r.Config.X509Config.KeyUsage.BaseKeyUsage.DigitalSignature: expected bool")
										}
									}
									if _, ok := rConfigX509ConfigKeyUsageBaseKeyUsage["encipherOnly"]; ok {
										if b, ok := rConfigX509ConfigKeyUsageBaseKeyUsage["encipherOnly"].(bool); ok {
											r.Config.X509Config.KeyUsage.BaseKeyUsage.EncipherOnly = dcl.Bool(b)
										} else {
											return nil, fmt.Errorf("r.Config.X509Config.KeyUsage.BaseKeyUsage.EncipherOnly: expected bool")
										}
									}
									if _, ok := rConfigX509ConfigKeyUsageBaseKeyUsage["keyAgreement"]; ok {
										if b, ok := rConfigX509ConfigKeyUsageBaseKeyUsage["keyAgreement"].(bool); ok {
											r.Config.X509Config.KeyUsage.BaseKeyUsage.KeyAgreement = dcl.Bool(b)
										} else {
											return nil, fmt.Errorf("r.Config.X509Config.KeyUsage.BaseKeyUsage.KeyAgreement: expected bool")
										}
									}
									if _, ok := rConfigX509ConfigKeyUsageBaseKeyUsage["keyEncipherment"]; ok {
										if b, ok := rConfigX509ConfigKeyUsageBaseKeyUsage["keyEncipherment"].(bool); ok {
											r.Config.X509Config.KeyUsage.BaseKeyUsage.KeyEncipherment = dcl.Bool(b)
										} else {
											return nil, fmt.Errorf("r.Config.X509Config.KeyUsage.BaseKeyUsage.KeyEncipherment: expected bool")
										}
									}
								} else {
									return nil, fmt.Errorf("r.Config.X509Config.KeyUsage.BaseKeyUsage: expected map[string]interface{}")
								}
							}
							if _, ok := rConfigX509ConfigKeyUsage["extendedKeyUsage"]; ok {
								if rConfigX509ConfigKeyUsageExtendedKeyUsage, ok := rConfigX509ConfigKeyUsage["extendedKeyUsage"].(map[string]interface{}); ok {
									r.Config.X509Config.KeyUsage.ExtendedKeyUsage = &dclService.CertificateConfigX509ConfigKeyUsageExtendedKeyUsage{}
									if _, ok := rConfigX509ConfigKeyUsageExtendedKeyUsage["clientAuth"]; ok {
										if b, ok := rConfigX509ConfigKeyUsageExtendedKeyUsage["clientAuth"].(bool); ok {
											r.Config.X509Config.KeyUsage.ExtendedKeyUsage.ClientAuth = dcl.Bool(b)
										} else {
											return nil, fmt.Errorf("r.Config.X509Config.KeyUsage.ExtendedKeyUsage.ClientAuth: expected bool")
										}
									}
									if _, ok := rConfigX509ConfigKeyUsageExtendedKeyUsage["codeSigning"]; ok {
										if b, ok := rConfigX509ConfigKeyUsageExtendedKeyUsage["codeSigning"].(bool); ok {
											r.Config.X509Config.KeyUsage.ExtendedKeyUsage.CodeSigning = dcl.Bool(b)
										} else {
											return nil, fmt.Errorf("r.Config.X509Config.KeyUsage.ExtendedKeyUsage.CodeSigning: expected bool")
										}
									}
									if _, ok := rConfigX509ConfigKeyUsageExtendedKeyUsage["emailProtection"]; ok {
										if b, ok := rConfigX509ConfigKeyUsageExtendedKeyUsage["emailProtection"].(bool); ok {
											r.Config.X509Config.KeyUsage.ExtendedKeyUsage.EmailProtection = dcl.Bool(b)
										} else {
											return nil, fmt.Errorf("r.Config.X509Config.KeyUsage.ExtendedKeyUsage.EmailProtection: expected bool")
										}
									}
									if _, ok := rConfigX509ConfigKeyUsageExtendedKeyUsage["ocspSigning"]; ok {
										if b, ok := rConfigX509ConfigKeyUsageExtendedKeyUsage["ocspSigning"].(bool); ok {
											r.Config.X509Config.KeyUsage.ExtendedKeyUsage.OcspSigning = dcl.Bool(b)
										} else {
											return nil, fmt.Errorf("r.Config.X509Config.KeyUsage.ExtendedKeyUsage.OcspSigning: expected bool")
										}
									}
									if _, ok := rConfigX509ConfigKeyUsageExtendedKeyUsage["serverAuth"]; ok {
										if b, ok := rConfigX509ConfigKeyUsageExtendedKeyUsage["serverAuth"].(bool); ok {
											r.Config.X509Config.KeyUsage.ExtendedKeyUsage.ServerAuth = dcl.Bool(b)
										} else {
											return nil, fmt.Errorf("r.Config.X509Config.KeyUsage.ExtendedKeyUsage.ServerAuth: expected bool")
										}
									}
									if _, ok := rConfigX509ConfigKeyUsageExtendedKeyUsage["timeStamping"]; ok {
										if b, ok := rConfigX509ConfigKeyUsageExtendedKeyUsage["timeStamping"].(bool); ok {
											r.Config.X509Config.KeyUsage.ExtendedKeyUsage.TimeStamping = dcl.Bool(b)
										} else {
											return nil, fmt.Errorf("r.Config.X509Config.KeyUsage.ExtendedKeyUsage.TimeStamping: expected bool")
										}
									}
								} else {
									return nil, fmt.Errorf("r.Config.X509Config.KeyUsage.ExtendedKeyUsage: expected map[string]interface{}")
								}
							}
							if _, ok := rConfigX509ConfigKeyUsage["unknownExtendedKeyUsages"]; ok {
								if s, ok := rConfigX509ConfigKeyUsage["unknownExtendedKeyUsages"].([]interface{}); ok {
									for _, o := range s {
										if objval, ok := o.(map[string]interface{}); ok {
											var rConfigX509ConfigKeyUsageUnknownExtendedKeyUsages dclService.CertificateConfigX509ConfigKeyUsageUnknownExtendedKeyUsages
											if _, ok := objval["objectIdPath"]; ok {
												if s, ok := objval["objectIdPath"].([]interface{}); ok {
													for _, ss := range s {
														if intval, ok := ss.(int64); ok {
															rConfigX509ConfigKeyUsageUnknownExtendedKeyUsages.ObjectIdPath = append(rConfigX509ConfigKeyUsageUnknownExtendedKeyUsages.ObjectIdPath, intval)
														}
													}
												} else {
													return nil, fmt.Errorf("rConfigX509ConfigKeyUsageUnknownExtendedKeyUsages.ObjectIdPath: expected []interface{}")
												}
											}
											r.Config.X509Config.KeyUsage.UnknownExtendedKeyUsages = append(r.Config.X509Config.KeyUsage.UnknownExtendedKeyUsages, rConfigX509ConfigKeyUsageUnknownExtendedKeyUsages)
										}
									}
								} else {
									return nil, fmt.Errorf("r.Config.X509Config.KeyUsage.UnknownExtendedKeyUsages: expected []interface{}")
								}
							}
						} else {
							return nil, fmt.Errorf("r.Config.X509Config.KeyUsage: expected map[string]interface{}")
						}
					}
					if _, ok := rConfigX509Config["policyIds"]; ok {
						if s, ok := rConfigX509Config["policyIds"].([]interface{}); ok {
							for _, o := range s {
								if objval, ok := o.(map[string]interface{}); ok {
									var rConfigX509ConfigPolicyIds dclService.CertificateConfigX509ConfigPolicyIds
									if _, ok := objval["objectIdPath"]; ok {
										if s, ok := objval["objectIdPath"].([]interface{}); ok {
											for _, ss := range s {
												if intval, ok := ss.(int64); ok {
													rConfigX509ConfigPolicyIds.ObjectIdPath = append(rConfigX509ConfigPolicyIds.ObjectIdPath, intval)
												}
											}
										} else {
											return nil, fmt.Errorf("rConfigX509ConfigPolicyIds.ObjectIdPath: expected []interface{}")
										}
									}
									r.Config.X509Config.PolicyIds = append(r.Config.X509Config.PolicyIds, rConfigX509ConfigPolicyIds)
								}
							}
						} else {
							return nil, fmt.Errorf("r.Config.X509Config.PolicyIds: expected []interface{}")
						}
					}
				} else {
					return nil, fmt.Errorf("r.Config.X509Config: expected map[string]interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.Config: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["createTime"]; ok {
		if s, ok := u.Object["createTime"].(string); ok {
			r.CreateTime = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.CreateTime: expected string")
		}
	}
	if _, ok := u.Object["issuerCertificateAuthority"]; ok {
		if s, ok := u.Object["issuerCertificateAuthority"].(string); ok {
			r.IssuerCertificateAuthority = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.IssuerCertificateAuthority: expected string")
		}
	}
	if _, ok := u.Object["labels"]; ok {
		if rLabels, ok := u.Object["labels"].(map[string]interface{}); ok {
			m := make(map[string]string)
			for k, v := range rLabels {
				if s, ok := v.(string); ok {
					m[k] = s
				}
			}
			r.Labels = m
		} else {
			return nil, fmt.Errorf("r.Labels: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["lifetime"]; ok {
		if s, ok := u.Object["lifetime"].(string); ok {
			r.Lifetime = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Lifetime: expected string")
		}
	}
	if _, ok := u.Object["location"]; ok {
		if s, ok := u.Object["location"].(string); ok {
			r.Location = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Location: expected string")
		}
	}
	if _, ok := u.Object["name"]; ok {
		if s, ok := u.Object["name"].(string); ok {
			r.Name = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Name: expected string")
		}
	}
	if _, ok := u.Object["pemCertificate"]; ok {
		if s, ok := u.Object["pemCertificate"].(string); ok {
			r.PemCertificate = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.PemCertificate: expected string")
		}
	}
	if _, ok := u.Object["pemCertificateChain"]; ok {
		if s, ok := u.Object["pemCertificateChain"].([]interface{}); ok {
			for _, ss := range s {
				if strval, ok := ss.(string); ok {
					r.PemCertificateChain = append(r.PemCertificateChain, strval)
				}
			}
		} else {
			return nil, fmt.Errorf("r.PemCertificateChain: expected []interface{}")
		}
	}
	if _, ok := u.Object["pemCsr"]; ok {
		if s, ok := u.Object["pemCsr"].(string); ok {
			r.PemCsr = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.PemCsr: expected string")
		}
	}
	if _, ok := u.Object["project"]; ok {
		if s, ok := u.Object["project"].(string); ok {
			r.Project = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Project: expected string")
		}
	}
	if _, ok := u.Object["revocationDetails"]; ok {
		if rRevocationDetails, ok := u.Object["revocationDetails"].(map[string]interface{}); ok {
			r.RevocationDetails = &dclService.CertificateRevocationDetails{}
			if _, ok := rRevocationDetails["revocationState"]; ok {
				if s, ok := rRevocationDetails["revocationState"].(string); ok {
					r.RevocationDetails.RevocationState = dclService.CertificateRevocationDetailsRevocationStateEnumRef(s)
				} else {
					return nil, fmt.Errorf("r.RevocationDetails.RevocationState: expected string")
				}
			}
			if _, ok := rRevocationDetails["revocationTime"]; ok {
				if s, ok := rRevocationDetails["revocationTime"].(string); ok {
					r.RevocationDetails.RevocationTime = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.RevocationDetails.RevocationTime: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.RevocationDetails: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["subjectMode"]; ok {
		if s, ok := u.Object["subjectMode"].(string); ok {
			r.SubjectMode = dclService.CertificateSubjectModeEnumRef(s)
		} else {
			return nil, fmt.Errorf("r.SubjectMode: expected string")
		}
	}
	if _, ok := u.Object["updateTime"]; ok {
		if s, ok := u.Object["updateTime"].(string); ok {
			r.UpdateTime = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.UpdateTime: expected string")
		}
	}
	return r, nil
}

func GetCertificate(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToCertificate(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetCertificate(ctx, r)
	if err != nil {
		return nil, err
	}
	return CertificateToUnstructured(r), nil
}

func ListCertificate(ctx context.Context, config *dcl.Config, project string, location string, caPool string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListCertificate(ctx, project, location, caPool)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, CertificateToUnstructured(r))
		}
		if !l.HasNext() {
			break
		}
		if err := l.Next(ctx, c); err != nil {
			return nil, err
		}
	}
	return resources, nil
}

func ApplyCertificate(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToCertificate(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToCertificate(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyCertificate(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return CertificateToUnstructured(r), nil
}

func CertificateHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToCertificate(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToCertificate(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyCertificate(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteCertificate(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToCertificate(u)
	if err != nil {
		return err
	}
	return c.DeleteCertificate(ctx, r)
}

func CertificateID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToCertificate(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *Certificate) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"privateca",
		"Certificate",
		"alpha",
	}
}

func (r *Certificate) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Certificate) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Certificate) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *Certificate) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Certificate) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Certificate) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Certificate) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetCertificate(ctx, config, resource)
}

func (r *Certificate) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyCertificate(ctx, config, resource, opts...)
}

func (r *Certificate) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return CertificateHasDiff(ctx, config, resource, opts...)
}

func (r *Certificate) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteCertificate(ctx, config, resource)
}

func (r *Certificate) ID(resource *unstructured.Resource) (string, error) {
	return CertificateID(resource)
}

func init() {
	unstructured.Register(&Certificate{})
}
