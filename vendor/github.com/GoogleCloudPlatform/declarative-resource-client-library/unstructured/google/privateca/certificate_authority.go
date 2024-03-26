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
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/privateca"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
)

type CertificateAuthority struct{}

func CertificateAuthorityToUnstructured(r *dclService.CertificateAuthority) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "privateca",
			Version: "ga",
			Type:    "CertificateAuthority",
		},
		Object: make(map[string]interface{}),
	}
	if r.AccessUrls != nil && r.AccessUrls != dclService.EmptyCertificateAuthorityAccessUrls {
		rAccessUrls := make(map[string]interface{})
		if r.AccessUrls.CaCertificateAccessUrl != nil {
			rAccessUrls["caCertificateAccessUrl"] = *r.AccessUrls.CaCertificateAccessUrl
		}
		var rAccessUrlsCrlAccessUrls []interface{}
		for _, rAccessUrlsCrlAccessUrlsVal := range r.AccessUrls.CrlAccessUrls {
			rAccessUrlsCrlAccessUrls = append(rAccessUrlsCrlAccessUrls, rAccessUrlsCrlAccessUrlsVal)
		}
		rAccessUrls["crlAccessUrls"] = rAccessUrlsCrlAccessUrls
		u.Object["accessUrls"] = rAccessUrls
	}
	var rCaCertificateDescriptions []interface{}
	for _, rCaCertificateDescriptionsVal := range r.CaCertificateDescriptions {
		rCaCertificateDescriptionsObject := make(map[string]interface{})
		var rCaCertificateDescriptionsValAiaIssuingCertificateUrls []interface{}
		for _, rCaCertificateDescriptionsValAiaIssuingCertificateUrlsVal := range rCaCertificateDescriptionsVal.AiaIssuingCertificateUrls {
			rCaCertificateDescriptionsValAiaIssuingCertificateUrls = append(rCaCertificateDescriptionsValAiaIssuingCertificateUrls, rCaCertificateDescriptionsValAiaIssuingCertificateUrlsVal)
		}
		rCaCertificateDescriptionsObject["aiaIssuingCertificateUrls"] = rCaCertificateDescriptionsValAiaIssuingCertificateUrls
		if rCaCertificateDescriptionsVal.AuthorityKeyId != nil && rCaCertificateDescriptionsVal.AuthorityKeyId != dclService.EmptyCertificateAuthorityCaCertificateDescriptionsAuthorityKeyId {
			rCaCertificateDescriptionsValAuthorityKeyId := make(map[string]interface{})
			if rCaCertificateDescriptionsVal.AuthorityKeyId.KeyId != nil {
				rCaCertificateDescriptionsValAuthorityKeyId["keyId"] = *rCaCertificateDescriptionsVal.AuthorityKeyId.KeyId
			}
			rCaCertificateDescriptionsObject["authorityKeyId"] = rCaCertificateDescriptionsValAuthorityKeyId
		}
		if rCaCertificateDescriptionsVal.CertFingerprint != nil && rCaCertificateDescriptionsVal.CertFingerprint != dclService.EmptyCertificateAuthorityCaCertificateDescriptionsCertFingerprint {
			rCaCertificateDescriptionsValCertFingerprint := make(map[string]interface{})
			if rCaCertificateDescriptionsVal.CertFingerprint.Sha256Hash != nil {
				rCaCertificateDescriptionsValCertFingerprint["sha256Hash"] = *rCaCertificateDescriptionsVal.CertFingerprint.Sha256Hash
			}
			rCaCertificateDescriptionsObject["certFingerprint"] = rCaCertificateDescriptionsValCertFingerprint
		}
		var rCaCertificateDescriptionsValCrlDistributionPoints []interface{}
		for _, rCaCertificateDescriptionsValCrlDistributionPointsVal := range rCaCertificateDescriptionsVal.CrlDistributionPoints {
			rCaCertificateDescriptionsValCrlDistributionPoints = append(rCaCertificateDescriptionsValCrlDistributionPoints, rCaCertificateDescriptionsValCrlDistributionPointsVal)
		}
		rCaCertificateDescriptionsObject["crlDistributionPoints"] = rCaCertificateDescriptionsValCrlDistributionPoints
		if rCaCertificateDescriptionsVal.PublicKey != nil && rCaCertificateDescriptionsVal.PublicKey != dclService.EmptyCertificateAuthorityCaCertificateDescriptionsPublicKey {
			rCaCertificateDescriptionsValPublicKey := make(map[string]interface{})
			if rCaCertificateDescriptionsVal.PublicKey.Format != nil {
				rCaCertificateDescriptionsValPublicKey["format"] = string(*rCaCertificateDescriptionsVal.PublicKey.Format)
			}
			if rCaCertificateDescriptionsVal.PublicKey.Key != nil {
				rCaCertificateDescriptionsValPublicKey["key"] = *rCaCertificateDescriptionsVal.PublicKey.Key
			}
			rCaCertificateDescriptionsObject["publicKey"] = rCaCertificateDescriptionsValPublicKey
		}
		if rCaCertificateDescriptionsVal.SubjectDescription != nil && rCaCertificateDescriptionsVal.SubjectDescription != dclService.EmptyCertificateAuthorityCaCertificateDescriptionsSubjectDescription {
			rCaCertificateDescriptionsValSubjectDescription := make(map[string]interface{})
			if rCaCertificateDescriptionsVal.SubjectDescription.HexSerialNumber != nil {
				rCaCertificateDescriptionsValSubjectDescription["hexSerialNumber"] = *rCaCertificateDescriptionsVal.SubjectDescription.HexSerialNumber
			}
			if rCaCertificateDescriptionsVal.SubjectDescription.Lifetime != nil {
				rCaCertificateDescriptionsValSubjectDescription["lifetime"] = *rCaCertificateDescriptionsVal.SubjectDescription.Lifetime
			}
			if rCaCertificateDescriptionsVal.SubjectDescription.NotAfterTime != nil {
				rCaCertificateDescriptionsValSubjectDescription["notAfterTime"] = *rCaCertificateDescriptionsVal.SubjectDescription.NotAfterTime
			}
			if rCaCertificateDescriptionsVal.SubjectDescription.NotBeforeTime != nil {
				rCaCertificateDescriptionsValSubjectDescription["notBeforeTime"] = *rCaCertificateDescriptionsVal.SubjectDescription.NotBeforeTime
			}
			if rCaCertificateDescriptionsVal.SubjectDescription.Subject != nil && rCaCertificateDescriptionsVal.SubjectDescription.Subject != dclService.EmptyCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject {
				rCaCertificateDescriptionsValSubjectDescriptionSubject := make(map[string]interface{})
				if rCaCertificateDescriptionsVal.SubjectDescription.Subject.CommonName != nil {
					rCaCertificateDescriptionsValSubjectDescriptionSubject["commonName"] = *rCaCertificateDescriptionsVal.SubjectDescription.Subject.CommonName
				}
				if rCaCertificateDescriptionsVal.SubjectDescription.Subject.CountryCode != nil {
					rCaCertificateDescriptionsValSubjectDescriptionSubject["countryCode"] = *rCaCertificateDescriptionsVal.SubjectDescription.Subject.CountryCode
				}
				if rCaCertificateDescriptionsVal.SubjectDescription.Subject.Locality != nil {
					rCaCertificateDescriptionsValSubjectDescriptionSubject["locality"] = *rCaCertificateDescriptionsVal.SubjectDescription.Subject.Locality
				}
				if rCaCertificateDescriptionsVal.SubjectDescription.Subject.Organization != nil {
					rCaCertificateDescriptionsValSubjectDescriptionSubject["organization"] = *rCaCertificateDescriptionsVal.SubjectDescription.Subject.Organization
				}
				if rCaCertificateDescriptionsVal.SubjectDescription.Subject.OrganizationalUnit != nil {
					rCaCertificateDescriptionsValSubjectDescriptionSubject["organizationalUnit"] = *rCaCertificateDescriptionsVal.SubjectDescription.Subject.OrganizationalUnit
				}
				if rCaCertificateDescriptionsVal.SubjectDescription.Subject.PostalCode != nil {
					rCaCertificateDescriptionsValSubjectDescriptionSubject["postalCode"] = *rCaCertificateDescriptionsVal.SubjectDescription.Subject.PostalCode
				}
				if rCaCertificateDescriptionsVal.SubjectDescription.Subject.Province != nil {
					rCaCertificateDescriptionsValSubjectDescriptionSubject["province"] = *rCaCertificateDescriptionsVal.SubjectDescription.Subject.Province
				}
				if rCaCertificateDescriptionsVal.SubjectDescription.Subject.StreetAddress != nil {
					rCaCertificateDescriptionsValSubjectDescriptionSubject["streetAddress"] = *rCaCertificateDescriptionsVal.SubjectDescription.Subject.StreetAddress
				}
				rCaCertificateDescriptionsValSubjectDescription["subject"] = rCaCertificateDescriptionsValSubjectDescriptionSubject
			}
			if rCaCertificateDescriptionsVal.SubjectDescription.SubjectAltName != nil && rCaCertificateDescriptionsVal.SubjectDescription.SubjectAltName != dclService.EmptyCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName {
				rCaCertificateDescriptionsValSubjectDescriptionSubjectAltName := make(map[string]interface{})
				var rCaCertificateDescriptionsValSubjectDescriptionSubjectAltNameCustomSans []interface{}
				for _, rCaCertificateDescriptionsValSubjectDescriptionSubjectAltNameCustomSansVal := range rCaCertificateDescriptionsVal.SubjectDescription.SubjectAltName.CustomSans {
					rCaCertificateDescriptionsValSubjectDescriptionSubjectAltNameCustomSansObject := make(map[string]interface{})
					if rCaCertificateDescriptionsValSubjectDescriptionSubjectAltNameCustomSansVal.Critical != nil {
						rCaCertificateDescriptionsValSubjectDescriptionSubjectAltNameCustomSansObject["critical"] = *rCaCertificateDescriptionsValSubjectDescriptionSubjectAltNameCustomSansVal.Critical
					}
					if rCaCertificateDescriptionsValSubjectDescriptionSubjectAltNameCustomSansVal.ObjectId != nil && rCaCertificateDescriptionsValSubjectDescriptionSubjectAltNameCustomSansVal.ObjectId != dclService.EmptyCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId {
						rCaCertificateDescriptionsValSubjectDescriptionSubjectAltNameCustomSansValObjectId := make(map[string]interface{})
						var rCaCertificateDescriptionsValSubjectDescriptionSubjectAltNameCustomSansValObjectIdObjectIdPath []interface{}
						for _, rCaCertificateDescriptionsValSubjectDescriptionSubjectAltNameCustomSansValObjectIdObjectIdPathVal := range rCaCertificateDescriptionsValSubjectDescriptionSubjectAltNameCustomSansVal.ObjectId.ObjectIdPath {
							rCaCertificateDescriptionsValSubjectDescriptionSubjectAltNameCustomSansValObjectIdObjectIdPath = append(rCaCertificateDescriptionsValSubjectDescriptionSubjectAltNameCustomSansValObjectIdObjectIdPath, rCaCertificateDescriptionsValSubjectDescriptionSubjectAltNameCustomSansValObjectIdObjectIdPathVal)
						}
						rCaCertificateDescriptionsValSubjectDescriptionSubjectAltNameCustomSansValObjectId["objectIdPath"] = rCaCertificateDescriptionsValSubjectDescriptionSubjectAltNameCustomSansValObjectIdObjectIdPath
						rCaCertificateDescriptionsValSubjectDescriptionSubjectAltNameCustomSansObject["objectId"] = rCaCertificateDescriptionsValSubjectDescriptionSubjectAltNameCustomSansValObjectId
					}
					if rCaCertificateDescriptionsValSubjectDescriptionSubjectAltNameCustomSansVal.Value != nil {
						rCaCertificateDescriptionsValSubjectDescriptionSubjectAltNameCustomSansObject["value"] = *rCaCertificateDescriptionsValSubjectDescriptionSubjectAltNameCustomSansVal.Value
					}
					rCaCertificateDescriptionsValSubjectDescriptionSubjectAltNameCustomSans = append(rCaCertificateDescriptionsValSubjectDescriptionSubjectAltNameCustomSans, rCaCertificateDescriptionsValSubjectDescriptionSubjectAltNameCustomSansObject)
				}
				rCaCertificateDescriptionsValSubjectDescriptionSubjectAltName["customSans"] = rCaCertificateDescriptionsValSubjectDescriptionSubjectAltNameCustomSans
				var rCaCertificateDescriptionsValSubjectDescriptionSubjectAltNameDnsNames []interface{}
				for _, rCaCertificateDescriptionsValSubjectDescriptionSubjectAltNameDnsNamesVal := range rCaCertificateDescriptionsVal.SubjectDescription.SubjectAltName.DnsNames {
					rCaCertificateDescriptionsValSubjectDescriptionSubjectAltNameDnsNames = append(rCaCertificateDescriptionsValSubjectDescriptionSubjectAltNameDnsNames, rCaCertificateDescriptionsValSubjectDescriptionSubjectAltNameDnsNamesVal)
				}
				rCaCertificateDescriptionsValSubjectDescriptionSubjectAltName["dnsNames"] = rCaCertificateDescriptionsValSubjectDescriptionSubjectAltNameDnsNames
				var rCaCertificateDescriptionsValSubjectDescriptionSubjectAltNameEmailAddresses []interface{}
				for _, rCaCertificateDescriptionsValSubjectDescriptionSubjectAltNameEmailAddressesVal := range rCaCertificateDescriptionsVal.SubjectDescription.SubjectAltName.EmailAddresses {
					rCaCertificateDescriptionsValSubjectDescriptionSubjectAltNameEmailAddresses = append(rCaCertificateDescriptionsValSubjectDescriptionSubjectAltNameEmailAddresses, rCaCertificateDescriptionsValSubjectDescriptionSubjectAltNameEmailAddressesVal)
				}
				rCaCertificateDescriptionsValSubjectDescriptionSubjectAltName["emailAddresses"] = rCaCertificateDescriptionsValSubjectDescriptionSubjectAltNameEmailAddresses
				var rCaCertificateDescriptionsValSubjectDescriptionSubjectAltNameIPAddresses []interface{}
				for _, rCaCertificateDescriptionsValSubjectDescriptionSubjectAltNameIPAddressesVal := range rCaCertificateDescriptionsVal.SubjectDescription.SubjectAltName.IPAddresses {
					rCaCertificateDescriptionsValSubjectDescriptionSubjectAltNameIPAddresses = append(rCaCertificateDescriptionsValSubjectDescriptionSubjectAltNameIPAddresses, rCaCertificateDescriptionsValSubjectDescriptionSubjectAltNameIPAddressesVal)
				}
				rCaCertificateDescriptionsValSubjectDescriptionSubjectAltName["ipAddresses"] = rCaCertificateDescriptionsValSubjectDescriptionSubjectAltNameIPAddresses
				var rCaCertificateDescriptionsValSubjectDescriptionSubjectAltNameUris []interface{}
				for _, rCaCertificateDescriptionsValSubjectDescriptionSubjectAltNameUrisVal := range rCaCertificateDescriptionsVal.SubjectDescription.SubjectAltName.Uris {
					rCaCertificateDescriptionsValSubjectDescriptionSubjectAltNameUris = append(rCaCertificateDescriptionsValSubjectDescriptionSubjectAltNameUris, rCaCertificateDescriptionsValSubjectDescriptionSubjectAltNameUrisVal)
				}
				rCaCertificateDescriptionsValSubjectDescriptionSubjectAltName["uris"] = rCaCertificateDescriptionsValSubjectDescriptionSubjectAltNameUris
				rCaCertificateDescriptionsValSubjectDescription["subjectAltName"] = rCaCertificateDescriptionsValSubjectDescriptionSubjectAltName
			}
			rCaCertificateDescriptionsObject["subjectDescription"] = rCaCertificateDescriptionsValSubjectDescription
		}
		if rCaCertificateDescriptionsVal.SubjectKeyId != nil && rCaCertificateDescriptionsVal.SubjectKeyId != dclService.EmptyCertificateAuthorityCaCertificateDescriptionsSubjectKeyId {
			rCaCertificateDescriptionsValSubjectKeyId := make(map[string]interface{})
			if rCaCertificateDescriptionsVal.SubjectKeyId.KeyId != nil {
				rCaCertificateDescriptionsValSubjectKeyId["keyId"] = *rCaCertificateDescriptionsVal.SubjectKeyId.KeyId
			}
			rCaCertificateDescriptionsObject["subjectKeyId"] = rCaCertificateDescriptionsValSubjectKeyId
		}
		if rCaCertificateDescriptionsVal.X509Description != nil && rCaCertificateDescriptionsVal.X509Description != dclService.EmptyCertificateAuthorityCaCertificateDescriptionsX509Description {
			rCaCertificateDescriptionsValX509Description := make(map[string]interface{})
			var rCaCertificateDescriptionsValX509DescriptionAdditionalExtensions []interface{}
			for _, rCaCertificateDescriptionsValX509DescriptionAdditionalExtensionsVal := range rCaCertificateDescriptionsVal.X509Description.AdditionalExtensions {
				rCaCertificateDescriptionsValX509DescriptionAdditionalExtensionsObject := make(map[string]interface{})
				if rCaCertificateDescriptionsValX509DescriptionAdditionalExtensionsVal.Critical != nil {
					rCaCertificateDescriptionsValX509DescriptionAdditionalExtensionsObject["critical"] = *rCaCertificateDescriptionsValX509DescriptionAdditionalExtensionsVal.Critical
				}
				if rCaCertificateDescriptionsValX509DescriptionAdditionalExtensionsVal.ObjectId != nil && rCaCertificateDescriptionsValX509DescriptionAdditionalExtensionsVal.ObjectId != dclService.EmptyCertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensionsObjectId {
					rCaCertificateDescriptionsValX509DescriptionAdditionalExtensionsValObjectId := make(map[string]interface{})
					var rCaCertificateDescriptionsValX509DescriptionAdditionalExtensionsValObjectIdObjectIdPath []interface{}
					for _, rCaCertificateDescriptionsValX509DescriptionAdditionalExtensionsValObjectIdObjectIdPathVal := range rCaCertificateDescriptionsValX509DescriptionAdditionalExtensionsVal.ObjectId.ObjectIdPath {
						rCaCertificateDescriptionsValX509DescriptionAdditionalExtensionsValObjectIdObjectIdPath = append(rCaCertificateDescriptionsValX509DescriptionAdditionalExtensionsValObjectIdObjectIdPath, rCaCertificateDescriptionsValX509DescriptionAdditionalExtensionsValObjectIdObjectIdPathVal)
					}
					rCaCertificateDescriptionsValX509DescriptionAdditionalExtensionsValObjectId["objectIdPath"] = rCaCertificateDescriptionsValX509DescriptionAdditionalExtensionsValObjectIdObjectIdPath
					rCaCertificateDescriptionsValX509DescriptionAdditionalExtensionsObject["objectId"] = rCaCertificateDescriptionsValX509DescriptionAdditionalExtensionsValObjectId
				}
				if rCaCertificateDescriptionsValX509DescriptionAdditionalExtensionsVal.Value != nil {
					rCaCertificateDescriptionsValX509DescriptionAdditionalExtensionsObject["value"] = *rCaCertificateDescriptionsValX509DescriptionAdditionalExtensionsVal.Value
				}
				rCaCertificateDescriptionsValX509DescriptionAdditionalExtensions = append(rCaCertificateDescriptionsValX509DescriptionAdditionalExtensions, rCaCertificateDescriptionsValX509DescriptionAdditionalExtensionsObject)
			}
			rCaCertificateDescriptionsValX509Description["additionalExtensions"] = rCaCertificateDescriptionsValX509DescriptionAdditionalExtensions
			var rCaCertificateDescriptionsValX509DescriptionAiaOcspServers []interface{}
			for _, rCaCertificateDescriptionsValX509DescriptionAiaOcspServersVal := range rCaCertificateDescriptionsVal.X509Description.AiaOcspServers {
				rCaCertificateDescriptionsValX509DescriptionAiaOcspServers = append(rCaCertificateDescriptionsValX509DescriptionAiaOcspServers, rCaCertificateDescriptionsValX509DescriptionAiaOcspServersVal)
			}
			rCaCertificateDescriptionsValX509Description["aiaOcspServers"] = rCaCertificateDescriptionsValX509DescriptionAiaOcspServers
			if rCaCertificateDescriptionsVal.X509Description.CaOptions != nil && rCaCertificateDescriptionsVal.X509Description.CaOptions != dclService.EmptyCertificateAuthorityCaCertificateDescriptionsX509DescriptionCaOptions {
				rCaCertificateDescriptionsValX509DescriptionCaOptions := make(map[string]interface{})
				if rCaCertificateDescriptionsVal.X509Description.CaOptions.IsCa != nil {
					rCaCertificateDescriptionsValX509DescriptionCaOptions["isCa"] = *rCaCertificateDescriptionsVal.X509Description.CaOptions.IsCa
				}
				if rCaCertificateDescriptionsVal.X509Description.CaOptions.MaxIssuerPathLength != nil {
					rCaCertificateDescriptionsValX509DescriptionCaOptions["maxIssuerPathLength"] = *rCaCertificateDescriptionsVal.X509Description.CaOptions.MaxIssuerPathLength
				}
				rCaCertificateDescriptionsValX509Description["caOptions"] = rCaCertificateDescriptionsValX509DescriptionCaOptions
			}
			if rCaCertificateDescriptionsVal.X509Description.KeyUsage != nil && rCaCertificateDescriptionsVal.X509Description.KeyUsage != dclService.EmptyCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsage {
				rCaCertificateDescriptionsValX509DescriptionKeyUsage := make(map[string]interface{})
				if rCaCertificateDescriptionsVal.X509Description.KeyUsage.BaseKeyUsage != nil && rCaCertificateDescriptionsVal.X509Description.KeyUsage.BaseKeyUsage != dclService.EmptyCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageBaseKeyUsage {
					rCaCertificateDescriptionsValX509DescriptionKeyUsageBaseKeyUsage := make(map[string]interface{})
					if rCaCertificateDescriptionsVal.X509Description.KeyUsage.BaseKeyUsage.CertSign != nil {
						rCaCertificateDescriptionsValX509DescriptionKeyUsageBaseKeyUsage["certSign"] = *rCaCertificateDescriptionsVal.X509Description.KeyUsage.BaseKeyUsage.CertSign
					}
					if rCaCertificateDescriptionsVal.X509Description.KeyUsage.BaseKeyUsage.ContentCommitment != nil {
						rCaCertificateDescriptionsValX509DescriptionKeyUsageBaseKeyUsage["contentCommitment"] = *rCaCertificateDescriptionsVal.X509Description.KeyUsage.BaseKeyUsage.ContentCommitment
					}
					if rCaCertificateDescriptionsVal.X509Description.KeyUsage.BaseKeyUsage.CrlSign != nil {
						rCaCertificateDescriptionsValX509DescriptionKeyUsageBaseKeyUsage["crlSign"] = *rCaCertificateDescriptionsVal.X509Description.KeyUsage.BaseKeyUsage.CrlSign
					}
					if rCaCertificateDescriptionsVal.X509Description.KeyUsage.BaseKeyUsage.DataEncipherment != nil {
						rCaCertificateDescriptionsValX509DescriptionKeyUsageBaseKeyUsage["dataEncipherment"] = *rCaCertificateDescriptionsVal.X509Description.KeyUsage.BaseKeyUsage.DataEncipherment
					}
					if rCaCertificateDescriptionsVal.X509Description.KeyUsage.BaseKeyUsage.DecipherOnly != nil {
						rCaCertificateDescriptionsValX509DescriptionKeyUsageBaseKeyUsage["decipherOnly"] = *rCaCertificateDescriptionsVal.X509Description.KeyUsage.BaseKeyUsage.DecipherOnly
					}
					if rCaCertificateDescriptionsVal.X509Description.KeyUsage.BaseKeyUsage.DigitalSignature != nil {
						rCaCertificateDescriptionsValX509DescriptionKeyUsageBaseKeyUsage["digitalSignature"] = *rCaCertificateDescriptionsVal.X509Description.KeyUsage.BaseKeyUsage.DigitalSignature
					}
					if rCaCertificateDescriptionsVal.X509Description.KeyUsage.BaseKeyUsage.EncipherOnly != nil {
						rCaCertificateDescriptionsValX509DescriptionKeyUsageBaseKeyUsage["encipherOnly"] = *rCaCertificateDescriptionsVal.X509Description.KeyUsage.BaseKeyUsage.EncipherOnly
					}
					if rCaCertificateDescriptionsVal.X509Description.KeyUsage.BaseKeyUsage.KeyAgreement != nil {
						rCaCertificateDescriptionsValX509DescriptionKeyUsageBaseKeyUsage["keyAgreement"] = *rCaCertificateDescriptionsVal.X509Description.KeyUsage.BaseKeyUsage.KeyAgreement
					}
					if rCaCertificateDescriptionsVal.X509Description.KeyUsage.BaseKeyUsage.KeyEncipherment != nil {
						rCaCertificateDescriptionsValX509DescriptionKeyUsageBaseKeyUsage["keyEncipherment"] = *rCaCertificateDescriptionsVal.X509Description.KeyUsage.BaseKeyUsage.KeyEncipherment
					}
					rCaCertificateDescriptionsValX509DescriptionKeyUsage["baseKeyUsage"] = rCaCertificateDescriptionsValX509DescriptionKeyUsageBaseKeyUsage
				}
				if rCaCertificateDescriptionsVal.X509Description.KeyUsage.ExtendedKeyUsage != nil && rCaCertificateDescriptionsVal.X509Description.KeyUsage.ExtendedKeyUsage != dclService.EmptyCertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageExtendedKeyUsage {
					rCaCertificateDescriptionsValX509DescriptionKeyUsageExtendedKeyUsage := make(map[string]interface{})
					if rCaCertificateDescriptionsVal.X509Description.KeyUsage.ExtendedKeyUsage.ClientAuth != nil {
						rCaCertificateDescriptionsValX509DescriptionKeyUsageExtendedKeyUsage["clientAuth"] = *rCaCertificateDescriptionsVal.X509Description.KeyUsage.ExtendedKeyUsage.ClientAuth
					}
					if rCaCertificateDescriptionsVal.X509Description.KeyUsage.ExtendedKeyUsage.CodeSigning != nil {
						rCaCertificateDescriptionsValX509DescriptionKeyUsageExtendedKeyUsage["codeSigning"] = *rCaCertificateDescriptionsVal.X509Description.KeyUsage.ExtendedKeyUsage.CodeSigning
					}
					if rCaCertificateDescriptionsVal.X509Description.KeyUsage.ExtendedKeyUsage.EmailProtection != nil {
						rCaCertificateDescriptionsValX509DescriptionKeyUsageExtendedKeyUsage["emailProtection"] = *rCaCertificateDescriptionsVal.X509Description.KeyUsage.ExtendedKeyUsage.EmailProtection
					}
					if rCaCertificateDescriptionsVal.X509Description.KeyUsage.ExtendedKeyUsage.OcspSigning != nil {
						rCaCertificateDescriptionsValX509DescriptionKeyUsageExtendedKeyUsage["ocspSigning"] = *rCaCertificateDescriptionsVal.X509Description.KeyUsage.ExtendedKeyUsage.OcspSigning
					}
					if rCaCertificateDescriptionsVal.X509Description.KeyUsage.ExtendedKeyUsage.ServerAuth != nil {
						rCaCertificateDescriptionsValX509DescriptionKeyUsageExtendedKeyUsage["serverAuth"] = *rCaCertificateDescriptionsVal.X509Description.KeyUsage.ExtendedKeyUsage.ServerAuth
					}
					if rCaCertificateDescriptionsVal.X509Description.KeyUsage.ExtendedKeyUsage.TimeStamping != nil {
						rCaCertificateDescriptionsValX509DescriptionKeyUsageExtendedKeyUsage["timeStamping"] = *rCaCertificateDescriptionsVal.X509Description.KeyUsage.ExtendedKeyUsage.TimeStamping
					}
					rCaCertificateDescriptionsValX509DescriptionKeyUsage["extendedKeyUsage"] = rCaCertificateDescriptionsValX509DescriptionKeyUsageExtendedKeyUsage
				}
				var rCaCertificateDescriptionsValX509DescriptionKeyUsageUnknownExtendedKeyUsages []interface{}
				for _, rCaCertificateDescriptionsValX509DescriptionKeyUsageUnknownExtendedKeyUsagesVal := range rCaCertificateDescriptionsVal.X509Description.KeyUsage.UnknownExtendedKeyUsages {
					rCaCertificateDescriptionsValX509DescriptionKeyUsageUnknownExtendedKeyUsagesObject := make(map[string]interface{})
					var rCaCertificateDescriptionsValX509DescriptionKeyUsageUnknownExtendedKeyUsagesValObjectIdPath []interface{}
					for _, rCaCertificateDescriptionsValX509DescriptionKeyUsageUnknownExtendedKeyUsagesValObjectIdPathVal := range rCaCertificateDescriptionsValX509DescriptionKeyUsageUnknownExtendedKeyUsagesVal.ObjectIdPath {
						rCaCertificateDescriptionsValX509DescriptionKeyUsageUnknownExtendedKeyUsagesValObjectIdPath = append(rCaCertificateDescriptionsValX509DescriptionKeyUsageUnknownExtendedKeyUsagesValObjectIdPath, rCaCertificateDescriptionsValX509DescriptionKeyUsageUnknownExtendedKeyUsagesValObjectIdPathVal)
					}
					rCaCertificateDescriptionsValX509DescriptionKeyUsageUnknownExtendedKeyUsagesObject["objectIdPath"] = rCaCertificateDescriptionsValX509DescriptionKeyUsageUnknownExtendedKeyUsagesValObjectIdPath
					rCaCertificateDescriptionsValX509DescriptionKeyUsageUnknownExtendedKeyUsages = append(rCaCertificateDescriptionsValX509DescriptionKeyUsageUnknownExtendedKeyUsages, rCaCertificateDescriptionsValX509DescriptionKeyUsageUnknownExtendedKeyUsagesObject)
				}
				rCaCertificateDescriptionsValX509DescriptionKeyUsage["unknownExtendedKeyUsages"] = rCaCertificateDescriptionsValX509DescriptionKeyUsageUnknownExtendedKeyUsages
				rCaCertificateDescriptionsValX509Description["keyUsage"] = rCaCertificateDescriptionsValX509DescriptionKeyUsage
			}
			var rCaCertificateDescriptionsValX509DescriptionPolicyIds []interface{}
			for _, rCaCertificateDescriptionsValX509DescriptionPolicyIdsVal := range rCaCertificateDescriptionsVal.X509Description.PolicyIds {
				rCaCertificateDescriptionsValX509DescriptionPolicyIdsObject := make(map[string]interface{})
				var rCaCertificateDescriptionsValX509DescriptionPolicyIdsValObjectIdPath []interface{}
				for _, rCaCertificateDescriptionsValX509DescriptionPolicyIdsValObjectIdPathVal := range rCaCertificateDescriptionsValX509DescriptionPolicyIdsVal.ObjectIdPath {
					rCaCertificateDescriptionsValX509DescriptionPolicyIdsValObjectIdPath = append(rCaCertificateDescriptionsValX509DescriptionPolicyIdsValObjectIdPath, rCaCertificateDescriptionsValX509DescriptionPolicyIdsValObjectIdPathVal)
				}
				rCaCertificateDescriptionsValX509DescriptionPolicyIdsObject["objectIdPath"] = rCaCertificateDescriptionsValX509DescriptionPolicyIdsValObjectIdPath
				rCaCertificateDescriptionsValX509DescriptionPolicyIds = append(rCaCertificateDescriptionsValX509DescriptionPolicyIds, rCaCertificateDescriptionsValX509DescriptionPolicyIdsObject)
			}
			rCaCertificateDescriptionsValX509Description["policyIds"] = rCaCertificateDescriptionsValX509DescriptionPolicyIds
			rCaCertificateDescriptionsObject["x509Description"] = rCaCertificateDescriptionsValX509Description
		}
		rCaCertificateDescriptions = append(rCaCertificateDescriptions, rCaCertificateDescriptionsObject)
	}
	u.Object["caCertificateDescriptions"] = rCaCertificateDescriptions
	if r.CaPool != nil {
		u.Object["caPool"] = *r.CaPool
	}
	if r.Config != nil && r.Config != dclService.EmptyCertificateAuthorityConfig {
		rConfig := make(map[string]interface{})
		if r.Config.PublicKey != nil && r.Config.PublicKey != dclService.EmptyCertificateAuthorityConfigPublicKey {
			rConfigPublicKey := make(map[string]interface{})
			if r.Config.PublicKey.Format != nil {
				rConfigPublicKey["format"] = string(*r.Config.PublicKey.Format)
			}
			if r.Config.PublicKey.Key != nil {
				rConfigPublicKey["key"] = *r.Config.PublicKey.Key
			}
			rConfig["publicKey"] = rConfigPublicKey
		}
		if r.Config.SubjectConfig != nil && r.Config.SubjectConfig != dclService.EmptyCertificateAuthorityConfigSubjectConfig {
			rConfigSubjectConfig := make(map[string]interface{})
			if r.Config.SubjectConfig.Subject != nil && r.Config.SubjectConfig.Subject != dclService.EmptyCertificateAuthorityConfigSubjectConfigSubject {
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
			if r.Config.SubjectConfig.SubjectAltName != nil && r.Config.SubjectConfig.SubjectAltName != dclService.EmptyCertificateAuthorityConfigSubjectConfigSubjectAltName {
				rConfigSubjectConfigSubjectAltName := make(map[string]interface{})
				var rConfigSubjectConfigSubjectAltNameCustomSans []interface{}
				for _, rConfigSubjectConfigSubjectAltNameCustomSansVal := range r.Config.SubjectConfig.SubjectAltName.CustomSans {
					rConfigSubjectConfigSubjectAltNameCustomSansObject := make(map[string]interface{})
					if rConfigSubjectConfigSubjectAltNameCustomSansVal.Critical != nil {
						rConfigSubjectConfigSubjectAltNameCustomSansObject["critical"] = *rConfigSubjectConfigSubjectAltNameCustomSansVal.Critical
					}
					if rConfigSubjectConfigSubjectAltNameCustomSansVal.ObjectId != nil && rConfigSubjectConfigSubjectAltNameCustomSansVal.ObjectId != dclService.EmptyCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId {
						rConfigSubjectConfigSubjectAltNameCustomSansValObjectId := make(map[string]interface{})
						var rConfigSubjectConfigSubjectAltNameCustomSansValObjectIdObjectIdPath []interface{}
						for _, rConfigSubjectConfigSubjectAltNameCustomSansValObjectIdObjectIdPathVal := range rConfigSubjectConfigSubjectAltNameCustomSansVal.ObjectId.ObjectIdPath {
							rConfigSubjectConfigSubjectAltNameCustomSansValObjectIdObjectIdPath = append(rConfigSubjectConfigSubjectAltNameCustomSansValObjectIdObjectIdPath, rConfigSubjectConfigSubjectAltNameCustomSansValObjectIdObjectIdPathVal)
						}
						rConfigSubjectConfigSubjectAltNameCustomSansValObjectId["objectIdPath"] = rConfigSubjectConfigSubjectAltNameCustomSansValObjectIdObjectIdPath
						rConfigSubjectConfigSubjectAltNameCustomSansObject["objectId"] = rConfigSubjectConfigSubjectAltNameCustomSansValObjectId
					}
					if rConfigSubjectConfigSubjectAltNameCustomSansVal.Value != nil {
						rConfigSubjectConfigSubjectAltNameCustomSansObject["value"] = *rConfigSubjectConfigSubjectAltNameCustomSansVal.Value
					}
					rConfigSubjectConfigSubjectAltNameCustomSans = append(rConfigSubjectConfigSubjectAltNameCustomSans, rConfigSubjectConfigSubjectAltNameCustomSansObject)
				}
				rConfigSubjectConfigSubjectAltName["customSans"] = rConfigSubjectConfigSubjectAltNameCustomSans
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
		if r.Config.X509Config != nil && r.Config.X509Config != dclService.EmptyCertificateAuthorityConfigX509Config {
			rConfigX509Config := make(map[string]interface{})
			var rConfigX509ConfigAdditionalExtensions []interface{}
			for _, rConfigX509ConfigAdditionalExtensionsVal := range r.Config.X509Config.AdditionalExtensions {
				rConfigX509ConfigAdditionalExtensionsObject := make(map[string]interface{})
				if rConfigX509ConfigAdditionalExtensionsVal.Critical != nil {
					rConfigX509ConfigAdditionalExtensionsObject["critical"] = *rConfigX509ConfigAdditionalExtensionsVal.Critical
				}
				if rConfigX509ConfigAdditionalExtensionsVal.ObjectId != nil && rConfigX509ConfigAdditionalExtensionsVal.ObjectId != dclService.EmptyCertificateAuthorityConfigX509ConfigAdditionalExtensionsObjectId {
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
			if r.Config.X509Config.CaOptions != nil && r.Config.X509Config.CaOptions != dclService.EmptyCertificateAuthorityConfigX509ConfigCaOptions {
				rConfigX509ConfigCaOptions := make(map[string]interface{})
				if r.Config.X509Config.CaOptions.IsCa != nil {
					rConfigX509ConfigCaOptions["isCa"] = *r.Config.X509Config.CaOptions.IsCa
				}
				if r.Config.X509Config.CaOptions.MaxIssuerPathLength != nil {
					rConfigX509ConfigCaOptions["maxIssuerPathLength"] = *r.Config.X509Config.CaOptions.MaxIssuerPathLength
				}
				if r.Config.X509Config.CaOptions.ZeroMaxIssuerPathLength != nil {
					rConfigX509ConfigCaOptions["zeroMaxIssuerPathLength"] = *r.Config.X509Config.CaOptions.ZeroMaxIssuerPathLength
				}
				rConfigX509Config["caOptions"] = rConfigX509ConfigCaOptions
			}
			if r.Config.X509Config.KeyUsage != nil && r.Config.X509Config.KeyUsage != dclService.EmptyCertificateAuthorityConfigX509ConfigKeyUsage {
				rConfigX509ConfigKeyUsage := make(map[string]interface{})
				if r.Config.X509Config.KeyUsage.BaseKeyUsage != nil && r.Config.X509Config.KeyUsage.BaseKeyUsage != dclService.EmptyCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsage {
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
				if r.Config.X509Config.KeyUsage.ExtendedKeyUsage != nil && r.Config.X509Config.KeyUsage.ExtendedKeyUsage != dclService.EmptyCertificateAuthorityConfigX509ConfigKeyUsageExtendedKeyUsage {
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
	if r.DeleteTime != nil {
		u.Object["deleteTime"] = *r.DeleteTime
	}
	if r.ExpireTime != nil {
		u.Object["expireTime"] = *r.ExpireTime
	}
	if r.GcsBucket != nil {
		u.Object["gcsBucket"] = *r.GcsBucket
	}
	if r.KeySpec != nil && r.KeySpec != dclService.EmptyCertificateAuthorityKeySpec {
		rKeySpec := make(map[string]interface{})
		if r.KeySpec.Algorithm != nil {
			rKeySpec["algorithm"] = string(*r.KeySpec.Algorithm)
		}
		if r.KeySpec.CloudKmsKeyVersion != nil {
			rKeySpec["cloudKmsKeyVersion"] = *r.KeySpec.CloudKmsKeyVersion
		}
		u.Object["keySpec"] = rKeySpec
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
	var rPemCaCertificates []interface{}
	for _, rPemCaCertificatesVal := range r.PemCaCertificates {
		rPemCaCertificates = append(rPemCaCertificates, rPemCaCertificatesVal)
	}
	u.Object["pemCaCertificates"] = rPemCaCertificates
	if r.Project != nil {
		u.Object["project"] = *r.Project
	}
	if r.State != nil {
		u.Object["state"] = string(*r.State)
	}
	if r.SubordinateConfig != nil && r.SubordinateConfig != dclService.EmptyCertificateAuthoritySubordinateConfig {
		rSubordinateConfig := make(map[string]interface{})
		if r.SubordinateConfig.CertificateAuthority != nil {
			rSubordinateConfig["certificateAuthority"] = *r.SubordinateConfig.CertificateAuthority
		}
		if r.SubordinateConfig.PemIssuerChain != nil && r.SubordinateConfig.PemIssuerChain != dclService.EmptyCertificateAuthoritySubordinateConfigPemIssuerChain {
			rSubordinateConfigPemIssuerChain := make(map[string]interface{})
			var rSubordinateConfigPemIssuerChainPemCertificates []interface{}
			for _, rSubordinateConfigPemIssuerChainPemCertificatesVal := range r.SubordinateConfig.PemIssuerChain.PemCertificates {
				rSubordinateConfigPemIssuerChainPemCertificates = append(rSubordinateConfigPemIssuerChainPemCertificates, rSubordinateConfigPemIssuerChainPemCertificatesVal)
			}
			rSubordinateConfigPemIssuerChain["pemCertificates"] = rSubordinateConfigPemIssuerChainPemCertificates
			rSubordinateConfig["pemIssuerChain"] = rSubordinateConfigPemIssuerChain
		}
		u.Object["subordinateConfig"] = rSubordinateConfig
	}
	if r.Tier != nil {
		u.Object["tier"] = string(*r.Tier)
	}
	if r.Type != nil {
		u.Object["type"] = string(*r.Type)
	}
	if r.UpdateTime != nil {
		u.Object["updateTime"] = *r.UpdateTime
	}
	return u
}

func UnstructuredToCertificateAuthority(u *unstructured.Resource) (*dclService.CertificateAuthority, error) {
	r := &dclService.CertificateAuthority{}
	if _, ok := u.Object["accessUrls"]; ok {
		if rAccessUrls, ok := u.Object["accessUrls"].(map[string]interface{}); ok {
			r.AccessUrls = &dclService.CertificateAuthorityAccessUrls{}
			if _, ok := rAccessUrls["caCertificateAccessUrl"]; ok {
				if s, ok := rAccessUrls["caCertificateAccessUrl"].(string); ok {
					r.AccessUrls.CaCertificateAccessUrl = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.AccessUrls.CaCertificateAccessUrl: expected string")
				}
			}
			if _, ok := rAccessUrls["crlAccessUrls"]; ok {
				if s, ok := rAccessUrls["crlAccessUrls"].([]interface{}); ok {
					for _, ss := range s {
						if strval, ok := ss.(string); ok {
							r.AccessUrls.CrlAccessUrls = append(r.AccessUrls.CrlAccessUrls, strval)
						}
					}
				} else {
					return nil, fmt.Errorf("r.AccessUrls.CrlAccessUrls: expected []interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.AccessUrls: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["caCertificateDescriptions"]; ok {
		if s, ok := u.Object["caCertificateDescriptions"].([]interface{}); ok {
			for _, o := range s {
				if objval, ok := o.(map[string]interface{}); ok {
					var rCaCertificateDescriptions dclService.CertificateAuthorityCaCertificateDescriptions
					if _, ok := objval["aiaIssuingCertificateUrls"]; ok {
						if s, ok := objval["aiaIssuingCertificateUrls"].([]interface{}); ok {
							for _, ss := range s {
								if strval, ok := ss.(string); ok {
									rCaCertificateDescriptions.AiaIssuingCertificateUrls = append(rCaCertificateDescriptions.AiaIssuingCertificateUrls, strval)
								}
							}
						} else {
							return nil, fmt.Errorf("rCaCertificateDescriptions.AiaIssuingCertificateUrls: expected []interface{}")
						}
					}
					if _, ok := objval["authorityKeyId"]; ok {
						if rCaCertificateDescriptionsAuthorityKeyId, ok := objval["authorityKeyId"].(map[string]interface{}); ok {
							rCaCertificateDescriptions.AuthorityKeyId = &dclService.CertificateAuthorityCaCertificateDescriptionsAuthorityKeyId{}
							if _, ok := rCaCertificateDescriptionsAuthorityKeyId["keyId"]; ok {
								if s, ok := rCaCertificateDescriptionsAuthorityKeyId["keyId"].(string); ok {
									rCaCertificateDescriptions.AuthorityKeyId.KeyId = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rCaCertificateDescriptions.AuthorityKeyId.KeyId: expected string")
								}
							}
						} else {
							return nil, fmt.Errorf("rCaCertificateDescriptions.AuthorityKeyId: expected map[string]interface{}")
						}
					}
					if _, ok := objval["certFingerprint"]; ok {
						if rCaCertificateDescriptionsCertFingerprint, ok := objval["certFingerprint"].(map[string]interface{}); ok {
							rCaCertificateDescriptions.CertFingerprint = &dclService.CertificateAuthorityCaCertificateDescriptionsCertFingerprint{}
							if _, ok := rCaCertificateDescriptionsCertFingerprint["sha256Hash"]; ok {
								if s, ok := rCaCertificateDescriptionsCertFingerprint["sha256Hash"].(string); ok {
									rCaCertificateDescriptions.CertFingerprint.Sha256Hash = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rCaCertificateDescriptions.CertFingerprint.Sha256Hash: expected string")
								}
							}
						} else {
							return nil, fmt.Errorf("rCaCertificateDescriptions.CertFingerprint: expected map[string]interface{}")
						}
					}
					if _, ok := objval["crlDistributionPoints"]; ok {
						if s, ok := objval["crlDistributionPoints"].([]interface{}); ok {
							for _, ss := range s {
								if strval, ok := ss.(string); ok {
									rCaCertificateDescriptions.CrlDistributionPoints = append(rCaCertificateDescriptions.CrlDistributionPoints, strval)
								}
							}
						} else {
							return nil, fmt.Errorf("rCaCertificateDescriptions.CrlDistributionPoints: expected []interface{}")
						}
					}
					if _, ok := objval["publicKey"]; ok {
						if rCaCertificateDescriptionsPublicKey, ok := objval["publicKey"].(map[string]interface{}); ok {
							rCaCertificateDescriptions.PublicKey = &dclService.CertificateAuthorityCaCertificateDescriptionsPublicKey{}
							if _, ok := rCaCertificateDescriptionsPublicKey["format"]; ok {
								if s, ok := rCaCertificateDescriptionsPublicKey["format"].(string); ok {
									rCaCertificateDescriptions.PublicKey.Format = dclService.CertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnumRef(s)
								} else {
									return nil, fmt.Errorf("rCaCertificateDescriptions.PublicKey.Format: expected string")
								}
							}
							if _, ok := rCaCertificateDescriptionsPublicKey["key"]; ok {
								if s, ok := rCaCertificateDescriptionsPublicKey["key"].(string); ok {
									rCaCertificateDescriptions.PublicKey.Key = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rCaCertificateDescriptions.PublicKey.Key: expected string")
								}
							}
						} else {
							return nil, fmt.Errorf("rCaCertificateDescriptions.PublicKey: expected map[string]interface{}")
						}
					}
					if _, ok := objval["subjectDescription"]; ok {
						if rCaCertificateDescriptionsSubjectDescription, ok := objval["subjectDescription"].(map[string]interface{}); ok {
							rCaCertificateDescriptions.SubjectDescription = &dclService.CertificateAuthorityCaCertificateDescriptionsSubjectDescription{}
							if _, ok := rCaCertificateDescriptionsSubjectDescription["hexSerialNumber"]; ok {
								if s, ok := rCaCertificateDescriptionsSubjectDescription["hexSerialNumber"].(string); ok {
									rCaCertificateDescriptions.SubjectDescription.HexSerialNumber = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rCaCertificateDescriptions.SubjectDescription.HexSerialNumber: expected string")
								}
							}
							if _, ok := rCaCertificateDescriptionsSubjectDescription["lifetime"]; ok {
								if s, ok := rCaCertificateDescriptionsSubjectDescription["lifetime"].(string); ok {
									rCaCertificateDescriptions.SubjectDescription.Lifetime = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rCaCertificateDescriptions.SubjectDescription.Lifetime: expected string")
								}
							}
							if _, ok := rCaCertificateDescriptionsSubjectDescription["notAfterTime"]; ok {
								if s, ok := rCaCertificateDescriptionsSubjectDescription["notAfterTime"].(string); ok {
									rCaCertificateDescriptions.SubjectDescription.NotAfterTime = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rCaCertificateDescriptions.SubjectDescription.NotAfterTime: expected string")
								}
							}
							if _, ok := rCaCertificateDescriptionsSubjectDescription["notBeforeTime"]; ok {
								if s, ok := rCaCertificateDescriptionsSubjectDescription["notBeforeTime"].(string); ok {
									rCaCertificateDescriptions.SubjectDescription.NotBeforeTime = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rCaCertificateDescriptions.SubjectDescription.NotBeforeTime: expected string")
								}
							}
							if _, ok := rCaCertificateDescriptionsSubjectDescription["subject"]; ok {
								if rCaCertificateDescriptionsSubjectDescriptionSubject, ok := rCaCertificateDescriptionsSubjectDescription["subject"].(map[string]interface{}); ok {
									rCaCertificateDescriptions.SubjectDescription.Subject = &dclService.CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject{}
									if _, ok := rCaCertificateDescriptionsSubjectDescriptionSubject["commonName"]; ok {
										if s, ok := rCaCertificateDescriptionsSubjectDescriptionSubject["commonName"].(string); ok {
											rCaCertificateDescriptions.SubjectDescription.Subject.CommonName = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rCaCertificateDescriptions.SubjectDescription.Subject.CommonName: expected string")
										}
									}
									if _, ok := rCaCertificateDescriptionsSubjectDescriptionSubject["countryCode"]; ok {
										if s, ok := rCaCertificateDescriptionsSubjectDescriptionSubject["countryCode"].(string); ok {
											rCaCertificateDescriptions.SubjectDescription.Subject.CountryCode = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rCaCertificateDescriptions.SubjectDescription.Subject.CountryCode: expected string")
										}
									}
									if _, ok := rCaCertificateDescriptionsSubjectDescriptionSubject["locality"]; ok {
										if s, ok := rCaCertificateDescriptionsSubjectDescriptionSubject["locality"].(string); ok {
											rCaCertificateDescriptions.SubjectDescription.Subject.Locality = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rCaCertificateDescriptions.SubjectDescription.Subject.Locality: expected string")
										}
									}
									if _, ok := rCaCertificateDescriptionsSubjectDescriptionSubject["organization"]; ok {
										if s, ok := rCaCertificateDescriptionsSubjectDescriptionSubject["organization"].(string); ok {
											rCaCertificateDescriptions.SubjectDescription.Subject.Organization = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rCaCertificateDescriptions.SubjectDescription.Subject.Organization: expected string")
										}
									}
									if _, ok := rCaCertificateDescriptionsSubjectDescriptionSubject["organizationalUnit"]; ok {
										if s, ok := rCaCertificateDescriptionsSubjectDescriptionSubject["organizationalUnit"].(string); ok {
											rCaCertificateDescriptions.SubjectDescription.Subject.OrganizationalUnit = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rCaCertificateDescriptions.SubjectDescription.Subject.OrganizationalUnit: expected string")
										}
									}
									if _, ok := rCaCertificateDescriptionsSubjectDescriptionSubject["postalCode"]; ok {
										if s, ok := rCaCertificateDescriptionsSubjectDescriptionSubject["postalCode"].(string); ok {
											rCaCertificateDescriptions.SubjectDescription.Subject.PostalCode = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rCaCertificateDescriptions.SubjectDescription.Subject.PostalCode: expected string")
										}
									}
									if _, ok := rCaCertificateDescriptionsSubjectDescriptionSubject["province"]; ok {
										if s, ok := rCaCertificateDescriptionsSubjectDescriptionSubject["province"].(string); ok {
											rCaCertificateDescriptions.SubjectDescription.Subject.Province = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rCaCertificateDescriptions.SubjectDescription.Subject.Province: expected string")
										}
									}
									if _, ok := rCaCertificateDescriptionsSubjectDescriptionSubject["streetAddress"]; ok {
										if s, ok := rCaCertificateDescriptionsSubjectDescriptionSubject["streetAddress"].(string); ok {
											rCaCertificateDescriptions.SubjectDescription.Subject.StreetAddress = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rCaCertificateDescriptions.SubjectDescription.Subject.StreetAddress: expected string")
										}
									}
								} else {
									return nil, fmt.Errorf("rCaCertificateDescriptions.SubjectDescription.Subject: expected map[string]interface{}")
								}
							}
							if _, ok := rCaCertificateDescriptionsSubjectDescription["subjectAltName"]; ok {
								if rCaCertificateDescriptionsSubjectDescriptionSubjectAltName, ok := rCaCertificateDescriptionsSubjectDescription["subjectAltName"].(map[string]interface{}); ok {
									rCaCertificateDescriptions.SubjectDescription.SubjectAltName = &dclService.CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName{}
									if _, ok := rCaCertificateDescriptionsSubjectDescriptionSubjectAltName["customSans"]; ok {
										if s, ok := rCaCertificateDescriptionsSubjectDescriptionSubjectAltName["customSans"].([]interface{}); ok {
											for _, o := range s {
												if objval, ok := o.(map[string]interface{}); ok {
													var rCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans dclService.CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans
													if _, ok := objval["critical"]; ok {
														if b, ok := objval["critical"].(bool); ok {
															rCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans.Critical = dcl.Bool(b)
														} else {
															return nil, fmt.Errorf("rCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans.Critical: expected bool")
														}
													}
													if _, ok := objval["objectId"]; ok {
														if rCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId, ok := objval["objectId"].(map[string]interface{}); ok {
															rCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans.ObjectId = &dclService.CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId{}
															if _, ok := rCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId["objectIdPath"]; ok {
																if s, ok := rCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId["objectIdPath"].([]interface{}); ok {
																	for _, ss := range s {
																		if intval, ok := ss.(int64); ok {
																			rCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans.ObjectId.ObjectIdPath = append(rCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans.ObjectId.ObjectIdPath, intval)
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans.ObjectId.ObjectIdPath: expected []interface{}")
																}
															}
														} else {
															return nil, fmt.Errorf("rCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans.ObjectId: expected map[string]interface{}")
														}
													}
													if _, ok := objval["value"]; ok {
														if s, ok := objval["value"].(string); ok {
															rCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans.Value = dcl.String(s)
														} else {
															return nil, fmt.Errorf("rCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans.Value: expected string")
														}
													}
													rCaCertificateDescriptions.SubjectDescription.SubjectAltName.CustomSans = append(rCaCertificateDescriptions.SubjectDescription.SubjectAltName.CustomSans, rCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans)
												}
											}
										} else {
											return nil, fmt.Errorf("rCaCertificateDescriptions.SubjectDescription.SubjectAltName.CustomSans: expected []interface{}")
										}
									}
									if _, ok := rCaCertificateDescriptionsSubjectDescriptionSubjectAltName["dnsNames"]; ok {
										if s, ok := rCaCertificateDescriptionsSubjectDescriptionSubjectAltName["dnsNames"].([]interface{}); ok {
											for _, ss := range s {
												if strval, ok := ss.(string); ok {
													rCaCertificateDescriptions.SubjectDescription.SubjectAltName.DnsNames = append(rCaCertificateDescriptions.SubjectDescription.SubjectAltName.DnsNames, strval)
												}
											}
										} else {
											return nil, fmt.Errorf("rCaCertificateDescriptions.SubjectDescription.SubjectAltName.DnsNames: expected []interface{}")
										}
									}
									if _, ok := rCaCertificateDescriptionsSubjectDescriptionSubjectAltName["emailAddresses"]; ok {
										if s, ok := rCaCertificateDescriptionsSubjectDescriptionSubjectAltName["emailAddresses"].([]interface{}); ok {
											for _, ss := range s {
												if strval, ok := ss.(string); ok {
													rCaCertificateDescriptions.SubjectDescription.SubjectAltName.EmailAddresses = append(rCaCertificateDescriptions.SubjectDescription.SubjectAltName.EmailAddresses, strval)
												}
											}
										} else {
											return nil, fmt.Errorf("rCaCertificateDescriptions.SubjectDescription.SubjectAltName.EmailAddresses: expected []interface{}")
										}
									}
									if _, ok := rCaCertificateDescriptionsSubjectDescriptionSubjectAltName["ipAddresses"]; ok {
										if s, ok := rCaCertificateDescriptionsSubjectDescriptionSubjectAltName["ipAddresses"].([]interface{}); ok {
											for _, ss := range s {
												if strval, ok := ss.(string); ok {
													rCaCertificateDescriptions.SubjectDescription.SubjectAltName.IPAddresses = append(rCaCertificateDescriptions.SubjectDescription.SubjectAltName.IPAddresses, strval)
												}
											}
										} else {
											return nil, fmt.Errorf("rCaCertificateDescriptions.SubjectDescription.SubjectAltName.IPAddresses: expected []interface{}")
										}
									}
									if _, ok := rCaCertificateDescriptionsSubjectDescriptionSubjectAltName["uris"]; ok {
										if s, ok := rCaCertificateDescriptionsSubjectDescriptionSubjectAltName["uris"].([]interface{}); ok {
											for _, ss := range s {
												if strval, ok := ss.(string); ok {
													rCaCertificateDescriptions.SubjectDescription.SubjectAltName.Uris = append(rCaCertificateDescriptions.SubjectDescription.SubjectAltName.Uris, strval)
												}
											}
										} else {
											return nil, fmt.Errorf("rCaCertificateDescriptions.SubjectDescription.SubjectAltName.Uris: expected []interface{}")
										}
									}
								} else {
									return nil, fmt.Errorf("rCaCertificateDescriptions.SubjectDescription.SubjectAltName: expected map[string]interface{}")
								}
							}
						} else {
							return nil, fmt.Errorf("rCaCertificateDescriptions.SubjectDescription: expected map[string]interface{}")
						}
					}
					if _, ok := objval["subjectKeyId"]; ok {
						if rCaCertificateDescriptionsSubjectKeyId, ok := objval["subjectKeyId"].(map[string]interface{}); ok {
							rCaCertificateDescriptions.SubjectKeyId = &dclService.CertificateAuthorityCaCertificateDescriptionsSubjectKeyId{}
							if _, ok := rCaCertificateDescriptionsSubjectKeyId["keyId"]; ok {
								if s, ok := rCaCertificateDescriptionsSubjectKeyId["keyId"].(string); ok {
									rCaCertificateDescriptions.SubjectKeyId.KeyId = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rCaCertificateDescriptions.SubjectKeyId.KeyId: expected string")
								}
							}
						} else {
							return nil, fmt.Errorf("rCaCertificateDescriptions.SubjectKeyId: expected map[string]interface{}")
						}
					}
					if _, ok := objval["x509Description"]; ok {
						if rCaCertificateDescriptionsX509Description, ok := objval["x509Description"].(map[string]interface{}); ok {
							rCaCertificateDescriptions.X509Description = &dclService.CertificateAuthorityCaCertificateDescriptionsX509Description{}
							if _, ok := rCaCertificateDescriptionsX509Description["additionalExtensions"]; ok {
								if s, ok := rCaCertificateDescriptionsX509Description["additionalExtensions"].([]interface{}); ok {
									for _, o := range s {
										if objval, ok := o.(map[string]interface{}); ok {
											var rCaCertificateDescriptionsX509DescriptionAdditionalExtensions dclService.CertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensions
											if _, ok := objval["critical"]; ok {
												if b, ok := objval["critical"].(bool); ok {
													rCaCertificateDescriptionsX509DescriptionAdditionalExtensions.Critical = dcl.Bool(b)
												} else {
													return nil, fmt.Errorf("rCaCertificateDescriptionsX509DescriptionAdditionalExtensions.Critical: expected bool")
												}
											}
											if _, ok := objval["objectId"]; ok {
												if rCaCertificateDescriptionsX509DescriptionAdditionalExtensionsObjectId, ok := objval["objectId"].(map[string]interface{}); ok {
													rCaCertificateDescriptionsX509DescriptionAdditionalExtensions.ObjectId = &dclService.CertificateAuthorityCaCertificateDescriptionsX509DescriptionAdditionalExtensionsObjectId{}
													if _, ok := rCaCertificateDescriptionsX509DescriptionAdditionalExtensionsObjectId["objectIdPath"]; ok {
														if s, ok := rCaCertificateDescriptionsX509DescriptionAdditionalExtensionsObjectId["objectIdPath"].([]interface{}); ok {
															for _, ss := range s {
																if intval, ok := ss.(int64); ok {
																	rCaCertificateDescriptionsX509DescriptionAdditionalExtensions.ObjectId.ObjectIdPath = append(rCaCertificateDescriptionsX509DescriptionAdditionalExtensions.ObjectId.ObjectIdPath, intval)
																}
															}
														} else {
															return nil, fmt.Errorf("rCaCertificateDescriptionsX509DescriptionAdditionalExtensions.ObjectId.ObjectIdPath: expected []interface{}")
														}
													}
												} else {
													return nil, fmt.Errorf("rCaCertificateDescriptionsX509DescriptionAdditionalExtensions.ObjectId: expected map[string]interface{}")
												}
											}
											if _, ok := objval["value"]; ok {
												if s, ok := objval["value"].(string); ok {
													rCaCertificateDescriptionsX509DescriptionAdditionalExtensions.Value = dcl.String(s)
												} else {
													return nil, fmt.Errorf("rCaCertificateDescriptionsX509DescriptionAdditionalExtensions.Value: expected string")
												}
											}
											rCaCertificateDescriptions.X509Description.AdditionalExtensions = append(rCaCertificateDescriptions.X509Description.AdditionalExtensions, rCaCertificateDescriptionsX509DescriptionAdditionalExtensions)
										}
									}
								} else {
									return nil, fmt.Errorf("rCaCertificateDescriptions.X509Description.AdditionalExtensions: expected []interface{}")
								}
							}
							if _, ok := rCaCertificateDescriptionsX509Description["aiaOcspServers"]; ok {
								if s, ok := rCaCertificateDescriptionsX509Description["aiaOcspServers"].([]interface{}); ok {
									for _, ss := range s {
										if strval, ok := ss.(string); ok {
											rCaCertificateDescriptions.X509Description.AiaOcspServers = append(rCaCertificateDescriptions.X509Description.AiaOcspServers, strval)
										}
									}
								} else {
									return nil, fmt.Errorf("rCaCertificateDescriptions.X509Description.AiaOcspServers: expected []interface{}")
								}
							}
							if _, ok := rCaCertificateDescriptionsX509Description["caOptions"]; ok {
								if rCaCertificateDescriptionsX509DescriptionCaOptions, ok := rCaCertificateDescriptionsX509Description["caOptions"].(map[string]interface{}); ok {
									rCaCertificateDescriptions.X509Description.CaOptions = &dclService.CertificateAuthorityCaCertificateDescriptionsX509DescriptionCaOptions{}
									if _, ok := rCaCertificateDescriptionsX509DescriptionCaOptions["isCa"]; ok {
										if b, ok := rCaCertificateDescriptionsX509DescriptionCaOptions["isCa"].(bool); ok {
											rCaCertificateDescriptions.X509Description.CaOptions.IsCa = dcl.Bool(b)
										} else {
											return nil, fmt.Errorf("rCaCertificateDescriptions.X509Description.CaOptions.IsCa: expected bool")
										}
									}
									if _, ok := rCaCertificateDescriptionsX509DescriptionCaOptions["maxIssuerPathLength"]; ok {
										if i, ok := rCaCertificateDescriptionsX509DescriptionCaOptions["maxIssuerPathLength"].(int64); ok {
											rCaCertificateDescriptions.X509Description.CaOptions.MaxIssuerPathLength = dcl.Int64(i)
										} else {
											return nil, fmt.Errorf("rCaCertificateDescriptions.X509Description.CaOptions.MaxIssuerPathLength: expected int64")
										}
									}
								} else {
									return nil, fmt.Errorf("rCaCertificateDescriptions.X509Description.CaOptions: expected map[string]interface{}")
								}
							}
							if _, ok := rCaCertificateDescriptionsX509Description["keyUsage"]; ok {
								if rCaCertificateDescriptionsX509DescriptionKeyUsage, ok := rCaCertificateDescriptionsX509Description["keyUsage"].(map[string]interface{}); ok {
									rCaCertificateDescriptions.X509Description.KeyUsage = &dclService.CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsage{}
									if _, ok := rCaCertificateDescriptionsX509DescriptionKeyUsage["baseKeyUsage"]; ok {
										if rCaCertificateDescriptionsX509DescriptionKeyUsageBaseKeyUsage, ok := rCaCertificateDescriptionsX509DescriptionKeyUsage["baseKeyUsage"].(map[string]interface{}); ok {
											rCaCertificateDescriptions.X509Description.KeyUsage.BaseKeyUsage = &dclService.CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageBaseKeyUsage{}
											if _, ok := rCaCertificateDescriptionsX509DescriptionKeyUsageBaseKeyUsage["certSign"]; ok {
												if b, ok := rCaCertificateDescriptionsX509DescriptionKeyUsageBaseKeyUsage["certSign"].(bool); ok {
													rCaCertificateDescriptions.X509Description.KeyUsage.BaseKeyUsage.CertSign = dcl.Bool(b)
												} else {
													return nil, fmt.Errorf("rCaCertificateDescriptions.X509Description.KeyUsage.BaseKeyUsage.CertSign: expected bool")
												}
											}
											if _, ok := rCaCertificateDescriptionsX509DescriptionKeyUsageBaseKeyUsage["contentCommitment"]; ok {
												if b, ok := rCaCertificateDescriptionsX509DescriptionKeyUsageBaseKeyUsage["contentCommitment"].(bool); ok {
													rCaCertificateDescriptions.X509Description.KeyUsage.BaseKeyUsage.ContentCommitment = dcl.Bool(b)
												} else {
													return nil, fmt.Errorf("rCaCertificateDescriptions.X509Description.KeyUsage.BaseKeyUsage.ContentCommitment: expected bool")
												}
											}
											if _, ok := rCaCertificateDescriptionsX509DescriptionKeyUsageBaseKeyUsage["crlSign"]; ok {
												if b, ok := rCaCertificateDescriptionsX509DescriptionKeyUsageBaseKeyUsage["crlSign"].(bool); ok {
													rCaCertificateDescriptions.X509Description.KeyUsage.BaseKeyUsage.CrlSign = dcl.Bool(b)
												} else {
													return nil, fmt.Errorf("rCaCertificateDescriptions.X509Description.KeyUsage.BaseKeyUsage.CrlSign: expected bool")
												}
											}
											if _, ok := rCaCertificateDescriptionsX509DescriptionKeyUsageBaseKeyUsage["dataEncipherment"]; ok {
												if b, ok := rCaCertificateDescriptionsX509DescriptionKeyUsageBaseKeyUsage["dataEncipherment"].(bool); ok {
													rCaCertificateDescriptions.X509Description.KeyUsage.BaseKeyUsage.DataEncipherment = dcl.Bool(b)
												} else {
													return nil, fmt.Errorf("rCaCertificateDescriptions.X509Description.KeyUsage.BaseKeyUsage.DataEncipherment: expected bool")
												}
											}
											if _, ok := rCaCertificateDescriptionsX509DescriptionKeyUsageBaseKeyUsage["decipherOnly"]; ok {
												if b, ok := rCaCertificateDescriptionsX509DescriptionKeyUsageBaseKeyUsage["decipherOnly"].(bool); ok {
													rCaCertificateDescriptions.X509Description.KeyUsage.BaseKeyUsage.DecipherOnly = dcl.Bool(b)
												} else {
													return nil, fmt.Errorf("rCaCertificateDescriptions.X509Description.KeyUsage.BaseKeyUsage.DecipherOnly: expected bool")
												}
											}
											if _, ok := rCaCertificateDescriptionsX509DescriptionKeyUsageBaseKeyUsage["digitalSignature"]; ok {
												if b, ok := rCaCertificateDescriptionsX509DescriptionKeyUsageBaseKeyUsage["digitalSignature"].(bool); ok {
													rCaCertificateDescriptions.X509Description.KeyUsage.BaseKeyUsage.DigitalSignature = dcl.Bool(b)
												} else {
													return nil, fmt.Errorf("rCaCertificateDescriptions.X509Description.KeyUsage.BaseKeyUsage.DigitalSignature: expected bool")
												}
											}
											if _, ok := rCaCertificateDescriptionsX509DescriptionKeyUsageBaseKeyUsage["encipherOnly"]; ok {
												if b, ok := rCaCertificateDescriptionsX509DescriptionKeyUsageBaseKeyUsage["encipherOnly"].(bool); ok {
													rCaCertificateDescriptions.X509Description.KeyUsage.BaseKeyUsage.EncipherOnly = dcl.Bool(b)
												} else {
													return nil, fmt.Errorf("rCaCertificateDescriptions.X509Description.KeyUsage.BaseKeyUsage.EncipherOnly: expected bool")
												}
											}
											if _, ok := rCaCertificateDescriptionsX509DescriptionKeyUsageBaseKeyUsage["keyAgreement"]; ok {
												if b, ok := rCaCertificateDescriptionsX509DescriptionKeyUsageBaseKeyUsage["keyAgreement"].(bool); ok {
													rCaCertificateDescriptions.X509Description.KeyUsage.BaseKeyUsage.KeyAgreement = dcl.Bool(b)
												} else {
													return nil, fmt.Errorf("rCaCertificateDescriptions.X509Description.KeyUsage.BaseKeyUsage.KeyAgreement: expected bool")
												}
											}
											if _, ok := rCaCertificateDescriptionsX509DescriptionKeyUsageBaseKeyUsage["keyEncipherment"]; ok {
												if b, ok := rCaCertificateDescriptionsX509DescriptionKeyUsageBaseKeyUsage["keyEncipherment"].(bool); ok {
													rCaCertificateDescriptions.X509Description.KeyUsage.BaseKeyUsage.KeyEncipherment = dcl.Bool(b)
												} else {
													return nil, fmt.Errorf("rCaCertificateDescriptions.X509Description.KeyUsage.BaseKeyUsage.KeyEncipherment: expected bool")
												}
											}
										} else {
											return nil, fmt.Errorf("rCaCertificateDescriptions.X509Description.KeyUsage.BaseKeyUsage: expected map[string]interface{}")
										}
									}
									if _, ok := rCaCertificateDescriptionsX509DescriptionKeyUsage["extendedKeyUsage"]; ok {
										if rCaCertificateDescriptionsX509DescriptionKeyUsageExtendedKeyUsage, ok := rCaCertificateDescriptionsX509DescriptionKeyUsage["extendedKeyUsage"].(map[string]interface{}); ok {
											rCaCertificateDescriptions.X509Description.KeyUsage.ExtendedKeyUsage = &dclService.CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageExtendedKeyUsage{}
											if _, ok := rCaCertificateDescriptionsX509DescriptionKeyUsageExtendedKeyUsage["clientAuth"]; ok {
												if b, ok := rCaCertificateDescriptionsX509DescriptionKeyUsageExtendedKeyUsage["clientAuth"].(bool); ok {
													rCaCertificateDescriptions.X509Description.KeyUsage.ExtendedKeyUsage.ClientAuth = dcl.Bool(b)
												} else {
													return nil, fmt.Errorf("rCaCertificateDescriptions.X509Description.KeyUsage.ExtendedKeyUsage.ClientAuth: expected bool")
												}
											}
											if _, ok := rCaCertificateDescriptionsX509DescriptionKeyUsageExtendedKeyUsage["codeSigning"]; ok {
												if b, ok := rCaCertificateDescriptionsX509DescriptionKeyUsageExtendedKeyUsage["codeSigning"].(bool); ok {
													rCaCertificateDescriptions.X509Description.KeyUsage.ExtendedKeyUsage.CodeSigning = dcl.Bool(b)
												} else {
													return nil, fmt.Errorf("rCaCertificateDescriptions.X509Description.KeyUsage.ExtendedKeyUsage.CodeSigning: expected bool")
												}
											}
											if _, ok := rCaCertificateDescriptionsX509DescriptionKeyUsageExtendedKeyUsage["emailProtection"]; ok {
												if b, ok := rCaCertificateDescriptionsX509DescriptionKeyUsageExtendedKeyUsage["emailProtection"].(bool); ok {
													rCaCertificateDescriptions.X509Description.KeyUsage.ExtendedKeyUsage.EmailProtection = dcl.Bool(b)
												} else {
													return nil, fmt.Errorf("rCaCertificateDescriptions.X509Description.KeyUsage.ExtendedKeyUsage.EmailProtection: expected bool")
												}
											}
											if _, ok := rCaCertificateDescriptionsX509DescriptionKeyUsageExtendedKeyUsage["ocspSigning"]; ok {
												if b, ok := rCaCertificateDescriptionsX509DescriptionKeyUsageExtendedKeyUsage["ocspSigning"].(bool); ok {
													rCaCertificateDescriptions.X509Description.KeyUsage.ExtendedKeyUsage.OcspSigning = dcl.Bool(b)
												} else {
													return nil, fmt.Errorf("rCaCertificateDescriptions.X509Description.KeyUsage.ExtendedKeyUsage.OcspSigning: expected bool")
												}
											}
											if _, ok := rCaCertificateDescriptionsX509DescriptionKeyUsageExtendedKeyUsage["serverAuth"]; ok {
												if b, ok := rCaCertificateDescriptionsX509DescriptionKeyUsageExtendedKeyUsage["serverAuth"].(bool); ok {
													rCaCertificateDescriptions.X509Description.KeyUsage.ExtendedKeyUsage.ServerAuth = dcl.Bool(b)
												} else {
													return nil, fmt.Errorf("rCaCertificateDescriptions.X509Description.KeyUsage.ExtendedKeyUsage.ServerAuth: expected bool")
												}
											}
											if _, ok := rCaCertificateDescriptionsX509DescriptionKeyUsageExtendedKeyUsage["timeStamping"]; ok {
												if b, ok := rCaCertificateDescriptionsX509DescriptionKeyUsageExtendedKeyUsage["timeStamping"].(bool); ok {
													rCaCertificateDescriptions.X509Description.KeyUsage.ExtendedKeyUsage.TimeStamping = dcl.Bool(b)
												} else {
													return nil, fmt.Errorf("rCaCertificateDescriptions.X509Description.KeyUsage.ExtendedKeyUsage.TimeStamping: expected bool")
												}
											}
										} else {
											return nil, fmt.Errorf("rCaCertificateDescriptions.X509Description.KeyUsage.ExtendedKeyUsage: expected map[string]interface{}")
										}
									}
									if _, ok := rCaCertificateDescriptionsX509DescriptionKeyUsage["unknownExtendedKeyUsages"]; ok {
										if s, ok := rCaCertificateDescriptionsX509DescriptionKeyUsage["unknownExtendedKeyUsages"].([]interface{}); ok {
											for _, o := range s {
												if objval, ok := o.(map[string]interface{}); ok {
													var rCaCertificateDescriptionsX509DescriptionKeyUsageUnknownExtendedKeyUsages dclService.CertificateAuthorityCaCertificateDescriptionsX509DescriptionKeyUsageUnknownExtendedKeyUsages
													if _, ok := objval["objectIdPath"]; ok {
														if s, ok := objval["objectIdPath"].([]interface{}); ok {
															for _, ss := range s {
																if intval, ok := ss.(int64); ok {
																	rCaCertificateDescriptionsX509DescriptionKeyUsageUnknownExtendedKeyUsages.ObjectIdPath = append(rCaCertificateDescriptionsX509DescriptionKeyUsageUnknownExtendedKeyUsages.ObjectIdPath, intval)
																}
															}
														} else {
															return nil, fmt.Errorf("rCaCertificateDescriptionsX509DescriptionKeyUsageUnknownExtendedKeyUsages.ObjectIdPath: expected []interface{}")
														}
													}
													rCaCertificateDescriptions.X509Description.KeyUsage.UnknownExtendedKeyUsages = append(rCaCertificateDescriptions.X509Description.KeyUsage.UnknownExtendedKeyUsages, rCaCertificateDescriptionsX509DescriptionKeyUsageUnknownExtendedKeyUsages)
												}
											}
										} else {
											return nil, fmt.Errorf("rCaCertificateDescriptions.X509Description.KeyUsage.UnknownExtendedKeyUsages: expected []interface{}")
										}
									}
								} else {
									return nil, fmt.Errorf("rCaCertificateDescriptions.X509Description.KeyUsage: expected map[string]interface{}")
								}
							}
							if _, ok := rCaCertificateDescriptionsX509Description["policyIds"]; ok {
								if s, ok := rCaCertificateDescriptionsX509Description["policyIds"].([]interface{}); ok {
									for _, o := range s {
										if objval, ok := o.(map[string]interface{}); ok {
											var rCaCertificateDescriptionsX509DescriptionPolicyIds dclService.CertificateAuthorityCaCertificateDescriptionsX509DescriptionPolicyIds
											if _, ok := objval["objectIdPath"]; ok {
												if s, ok := objval["objectIdPath"].([]interface{}); ok {
													for _, ss := range s {
														if intval, ok := ss.(int64); ok {
															rCaCertificateDescriptionsX509DescriptionPolicyIds.ObjectIdPath = append(rCaCertificateDescriptionsX509DescriptionPolicyIds.ObjectIdPath, intval)
														}
													}
												} else {
													return nil, fmt.Errorf("rCaCertificateDescriptionsX509DescriptionPolicyIds.ObjectIdPath: expected []interface{}")
												}
											}
											rCaCertificateDescriptions.X509Description.PolicyIds = append(rCaCertificateDescriptions.X509Description.PolicyIds, rCaCertificateDescriptionsX509DescriptionPolicyIds)
										}
									}
								} else {
									return nil, fmt.Errorf("rCaCertificateDescriptions.X509Description.PolicyIds: expected []interface{}")
								}
							}
						} else {
							return nil, fmt.Errorf("rCaCertificateDescriptions.X509Description: expected map[string]interface{}")
						}
					}
					r.CaCertificateDescriptions = append(r.CaCertificateDescriptions, rCaCertificateDescriptions)
				}
			}
		} else {
			return nil, fmt.Errorf("r.CaCertificateDescriptions: expected []interface{}")
		}
	}
	if _, ok := u.Object["caPool"]; ok {
		if s, ok := u.Object["caPool"].(string); ok {
			r.CaPool = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.CaPool: expected string")
		}
	}
	if _, ok := u.Object["config"]; ok {
		if rConfig, ok := u.Object["config"].(map[string]interface{}); ok {
			r.Config = &dclService.CertificateAuthorityConfig{}
			if _, ok := rConfig["publicKey"]; ok {
				if rConfigPublicKey, ok := rConfig["publicKey"].(map[string]interface{}); ok {
					r.Config.PublicKey = &dclService.CertificateAuthorityConfigPublicKey{}
					if _, ok := rConfigPublicKey["format"]; ok {
						if s, ok := rConfigPublicKey["format"].(string); ok {
							r.Config.PublicKey.Format = dclService.CertificateAuthorityConfigPublicKeyFormatEnumRef(s)
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
					r.Config.SubjectConfig = &dclService.CertificateAuthorityConfigSubjectConfig{}
					if _, ok := rConfigSubjectConfig["subject"]; ok {
						if rConfigSubjectConfigSubject, ok := rConfigSubjectConfig["subject"].(map[string]interface{}); ok {
							r.Config.SubjectConfig.Subject = &dclService.CertificateAuthorityConfigSubjectConfigSubject{}
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
							r.Config.SubjectConfig.SubjectAltName = &dclService.CertificateAuthorityConfigSubjectConfigSubjectAltName{}
							if _, ok := rConfigSubjectConfigSubjectAltName["customSans"]; ok {
								if s, ok := rConfigSubjectConfigSubjectAltName["customSans"].([]interface{}); ok {
									for _, o := range s {
										if objval, ok := o.(map[string]interface{}); ok {
											var rConfigSubjectConfigSubjectAltNameCustomSans dclService.CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans
											if _, ok := objval["critical"]; ok {
												if b, ok := objval["critical"].(bool); ok {
													rConfigSubjectConfigSubjectAltNameCustomSans.Critical = dcl.Bool(b)
												} else {
													return nil, fmt.Errorf("rConfigSubjectConfigSubjectAltNameCustomSans.Critical: expected bool")
												}
											}
											if _, ok := objval["objectId"]; ok {
												if rConfigSubjectConfigSubjectAltNameCustomSansObjectId, ok := objval["objectId"].(map[string]interface{}); ok {
													rConfigSubjectConfigSubjectAltNameCustomSans.ObjectId = &dclService.CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId{}
													if _, ok := rConfigSubjectConfigSubjectAltNameCustomSansObjectId["objectIdPath"]; ok {
														if s, ok := rConfigSubjectConfigSubjectAltNameCustomSansObjectId["objectIdPath"].([]interface{}); ok {
															for _, ss := range s {
																if intval, ok := ss.(int64); ok {
																	rConfigSubjectConfigSubjectAltNameCustomSans.ObjectId.ObjectIdPath = append(rConfigSubjectConfigSubjectAltNameCustomSans.ObjectId.ObjectIdPath, intval)
																}
															}
														} else {
															return nil, fmt.Errorf("rConfigSubjectConfigSubjectAltNameCustomSans.ObjectId.ObjectIdPath: expected []interface{}")
														}
													}
												} else {
													return nil, fmt.Errorf("rConfigSubjectConfigSubjectAltNameCustomSans.ObjectId: expected map[string]interface{}")
												}
											}
											if _, ok := objval["value"]; ok {
												if s, ok := objval["value"].(string); ok {
													rConfigSubjectConfigSubjectAltNameCustomSans.Value = dcl.String(s)
												} else {
													return nil, fmt.Errorf("rConfigSubjectConfigSubjectAltNameCustomSans.Value: expected string")
												}
											}
											r.Config.SubjectConfig.SubjectAltName.CustomSans = append(r.Config.SubjectConfig.SubjectAltName.CustomSans, rConfigSubjectConfigSubjectAltNameCustomSans)
										}
									}
								} else {
									return nil, fmt.Errorf("r.Config.SubjectConfig.SubjectAltName.CustomSans: expected []interface{}")
								}
							}
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
					r.Config.X509Config = &dclService.CertificateAuthorityConfigX509Config{}
					if _, ok := rConfigX509Config["additionalExtensions"]; ok {
						if s, ok := rConfigX509Config["additionalExtensions"].([]interface{}); ok {
							for _, o := range s {
								if objval, ok := o.(map[string]interface{}); ok {
									var rConfigX509ConfigAdditionalExtensions dclService.CertificateAuthorityConfigX509ConfigAdditionalExtensions
									if _, ok := objval["critical"]; ok {
										if b, ok := objval["critical"].(bool); ok {
											rConfigX509ConfigAdditionalExtensions.Critical = dcl.Bool(b)
										} else {
											return nil, fmt.Errorf("rConfigX509ConfigAdditionalExtensions.Critical: expected bool")
										}
									}
									if _, ok := objval["objectId"]; ok {
										if rConfigX509ConfigAdditionalExtensionsObjectId, ok := objval["objectId"].(map[string]interface{}); ok {
											rConfigX509ConfigAdditionalExtensions.ObjectId = &dclService.CertificateAuthorityConfigX509ConfigAdditionalExtensionsObjectId{}
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
							r.Config.X509Config.CaOptions = &dclService.CertificateAuthorityConfigX509ConfigCaOptions{}
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
							r.Config.X509Config.KeyUsage = &dclService.CertificateAuthorityConfigX509ConfigKeyUsage{}
							if _, ok := rConfigX509ConfigKeyUsage["baseKeyUsage"]; ok {
								if rConfigX509ConfigKeyUsageBaseKeyUsage, ok := rConfigX509ConfigKeyUsage["baseKeyUsage"].(map[string]interface{}); ok {
									r.Config.X509Config.KeyUsage.BaseKeyUsage = &dclService.CertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsage{}
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
									r.Config.X509Config.KeyUsage.ExtendedKeyUsage = &dclService.CertificateAuthorityConfigX509ConfigKeyUsageExtendedKeyUsage{}
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
											var rConfigX509ConfigKeyUsageUnknownExtendedKeyUsages dclService.CertificateAuthorityConfigX509ConfigKeyUsageUnknownExtendedKeyUsages
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
									var rConfigX509ConfigPolicyIds dclService.CertificateAuthorityConfigX509ConfigPolicyIds
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
	if _, ok := u.Object["deleteTime"]; ok {
		if s, ok := u.Object["deleteTime"].(string); ok {
			r.DeleteTime = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.DeleteTime: expected string")
		}
	}
	if _, ok := u.Object["expireTime"]; ok {
		if s, ok := u.Object["expireTime"].(string); ok {
			r.ExpireTime = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.ExpireTime: expected string")
		}
	}
	if _, ok := u.Object["gcsBucket"]; ok {
		if s, ok := u.Object["gcsBucket"].(string); ok {
			r.GcsBucket = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.GcsBucket: expected string")
		}
	}
	if _, ok := u.Object["keySpec"]; ok {
		if rKeySpec, ok := u.Object["keySpec"].(map[string]interface{}); ok {
			r.KeySpec = &dclService.CertificateAuthorityKeySpec{}
			if _, ok := rKeySpec["algorithm"]; ok {
				if s, ok := rKeySpec["algorithm"].(string); ok {
					r.KeySpec.Algorithm = dclService.CertificateAuthorityKeySpecAlgorithmEnumRef(s)
				} else {
					return nil, fmt.Errorf("r.KeySpec.Algorithm: expected string")
				}
			}
			if _, ok := rKeySpec["cloudKmsKeyVersion"]; ok {
				if s, ok := rKeySpec["cloudKmsKeyVersion"].(string); ok {
					r.KeySpec.CloudKmsKeyVersion = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.KeySpec.CloudKmsKeyVersion: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.KeySpec: expected map[string]interface{}")
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
	if _, ok := u.Object["pemCaCertificates"]; ok {
		if s, ok := u.Object["pemCaCertificates"].([]interface{}); ok {
			for _, ss := range s {
				if strval, ok := ss.(string); ok {
					r.PemCaCertificates = append(r.PemCaCertificates, strval)
				}
			}
		} else {
			return nil, fmt.Errorf("r.PemCaCertificates: expected []interface{}")
		}
	}
	if _, ok := u.Object["project"]; ok {
		if s, ok := u.Object["project"].(string); ok {
			r.Project = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Project: expected string")
		}
	}
	if _, ok := u.Object["state"]; ok {
		if s, ok := u.Object["state"].(string); ok {
			r.State = dclService.CertificateAuthorityStateEnumRef(s)
		} else {
			return nil, fmt.Errorf("r.State: expected string")
		}
	}
	if _, ok := u.Object["subordinateConfig"]; ok {
		if rSubordinateConfig, ok := u.Object["subordinateConfig"].(map[string]interface{}); ok {
			r.SubordinateConfig = &dclService.CertificateAuthoritySubordinateConfig{}
			if _, ok := rSubordinateConfig["certificateAuthority"]; ok {
				if s, ok := rSubordinateConfig["certificateAuthority"].(string); ok {
					r.SubordinateConfig.CertificateAuthority = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.SubordinateConfig.CertificateAuthority: expected string")
				}
			}
			if _, ok := rSubordinateConfig["pemIssuerChain"]; ok {
				if rSubordinateConfigPemIssuerChain, ok := rSubordinateConfig["pemIssuerChain"].(map[string]interface{}); ok {
					r.SubordinateConfig.PemIssuerChain = &dclService.CertificateAuthoritySubordinateConfigPemIssuerChain{}
					if _, ok := rSubordinateConfigPemIssuerChain["pemCertificates"]; ok {
						if s, ok := rSubordinateConfigPemIssuerChain["pemCertificates"].([]interface{}); ok {
							for _, ss := range s {
								if strval, ok := ss.(string); ok {
									r.SubordinateConfig.PemIssuerChain.PemCertificates = append(r.SubordinateConfig.PemIssuerChain.PemCertificates, strval)
								}
							}
						} else {
							return nil, fmt.Errorf("r.SubordinateConfig.PemIssuerChain.PemCertificates: expected []interface{}")
						}
					}
				} else {
					return nil, fmt.Errorf("r.SubordinateConfig.PemIssuerChain: expected map[string]interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.SubordinateConfig: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["tier"]; ok {
		if s, ok := u.Object["tier"].(string); ok {
			r.Tier = dclService.CertificateAuthorityTierEnumRef(s)
		} else {
			return nil, fmt.Errorf("r.Tier: expected string")
		}
	}
	if _, ok := u.Object["type"]; ok {
		if s, ok := u.Object["type"].(string); ok {
			r.Type = dclService.CertificateAuthorityTypeEnumRef(s)
		} else {
			return nil, fmt.Errorf("r.Type: expected string")
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

func GetCertificateAuthority(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToCertificateAuthority(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetCertificateAuthority(ctx, r)
	if err != nil {
		return nil, err
	}
	return CertificateAuthorityToUnstructured(r), nil
}

func ListCertificateAuthority(ctx context.Context, config *dcl.Config, project string, location string, caPool string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListCertificateAuthority(ctx, project, location, caPool)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, CertificateAuthorityToUnstructured(r))
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

func ApplyCertificateAuthority(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToCertificateAuthority(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToCertificateAuthority(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyCertificateAuthority(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return CertificateAuthorityToUnstructured(r), nil
}

func CertificateAuthorityHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToCertificateAuthority(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToCertificateAuthority(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyCertificateAuthority(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteCertificateAuthority(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToCertificateAuthority(u)
	if err != nil {
		return err
	}
	return c.DeleteCertificateAuthority(ctx, r)
}

func CertificateAuthorityID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToCertificateAuthority(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *CertificateAuthority) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"privateca",
		"CertificateAuthority",
		"ga",
	}
}

func (r *CertificateAuthority) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *CertificateAuthority) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *CertificateAuthority) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *CertificateAuthority) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *CertificateAuthority) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *CertificateAuthority) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *CertificateAuthority) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetCertificateAuthority(ctx, config, resource)
}

func (r *CertificateAuthority) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyCertificateAuthority(ctx, config, resource, opts...)
}

func (r *CertificateAuthority) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return CertificateAuthorityHasDiff(ctx, config, resource, opts...)
}

func (r *CertificateAuthority) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteCertificateAuthority(ctx, config, resource)
}

func (r *CertificateAuthority) ID(resource *unstructured.Resource) (string, error) {
	return CertificateAuthorityID(resource)
}

func init() {
	unstructured.Register(&CertificateAuthority{})
}
