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
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/privateca/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
)

type CaPool struct{}

func CaPoolToUnstructured(r *dclService.CaPool) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "privateca",
			Version: "beta",
			Type:    "CaPool",
		},
		Object: make(map[string]interface{}),
	}
	if r.IssuancePolicy != nil && r.IssuancePolicy != dclService.EmptyCaPoolIssuancePolicy {
		rIssuancePolicy := make(map[string]interface{})
		if r.IssuancePolicy.AllowedIssuanceModes != nil && r.IssuancePolicy.AllowedIssuanceModes != dclService.EmptyCaPoolIssuancePolicyAllowedIssuanceModes {
			rIssuancePolicyAllowedIssuanceModes := make(map[string]interface{})
			if r.IssuancePolicy.AllowedIssuanceModes.AllowConfigBasedIssuance != nil {
				rIssuancePolicyAllowedIssuanceModes["allowConfigBasedIssuance"] = *r.IssuancePolicy.AllowedIssuanceModes.AllowConfigBasedIssuance
			}
			if r.IssuancePolicy.AllowedIssuanceModes.AllowCsrBasedIssuance != nil {
				rIssuancePolicyAllowedIssuanceModes["allowCsrBasedIssuance"] = *r.IssuancePolicy.AllowedIssuanceModes.AllowCsrBasedIssuance
			}
			rIssuancePolicy["allowedIssuanceModes"] = rIssuancePolicyAllowedIssuanceModes
		}
		var rIssuancePolicyAllowedKeyTypes []interface{}
		for _, rIssuancePolicyAllowedKeyTypesVal := range r.IssuancePolicy.AllowedKeyTypes {
			rIssuancePolicyAllowedKeyTypesObject := make(map[string]interface{})
			if rIssuancePolicyAllowedKeyTypesVal.EllipticCurve != nil && rIssuancePolicyAllowedKeyTypesVal.EllipticCurve != dclService.EmptyCaPoolIssuancePolicyAllowedKeyTypesEllipticCurve {
				rIssuancePolicyAllowedKeyTypesValEllipticCurve := make(map[string]interface{})
				if rIssuancePolicyAllowedKeyTypesVal.EllipticCurve.SignatureAlgorithm != nil {
					rIssuancePolicyAllowedKeyTypesValEllipticCurve["signatureAlgorithm"] = string(*rIssuancePolicyAllowedKeyTypesVal.EllipticCurve.SignatureAlgorithm)
				}
				rIssuancePolicyAllowedKeyTypesObject["ellipticCurve"] = rIssuancePolicyAllowedKeyTypesValEllipticCurve
			}
			if rIssuancePolicyAllowedKeyTypesVal.Rsa != nil && rIssuancePolicyAllowedKeyTypesVal.Rsa != dclService.EmptyCaPoolIssuancePolicyAllowedKeyTypesRsa {
				rIssuancePolicyAllowedKeyTypesValRsa := make(map[string]interface{})
				if rIssuancePolicyAllowedKeyTypesVal.Rsa.MaxModulusSize != nil {
					rIssuancePolicyAllowedKeyTypesValRsa["maxModulusSize"] = *rIssuancePolicyAllowedKeyTypesVal.Rsa.MaxModulusSize
				}
				if rIssuancePolicyAllowedKeyTypesVal.Rsa.MinModulusSize != nil {
					rIssuancePolicyAllowedKeyTypesValRsa["minModulusSize"] = *rIssuancePolicyAllowedKeyTypesVal.Rsa.MinModulusSize
				}
				rIssuancePolicyAllowedKeyTypesObject["rsa"] = rIssuancePolicyAllowedKeyTypesValRsa
			}
			rIssuancePolicyAllowedKeyTypes = append(rIssuancePolicyAllowedKeyTypes, rIssuancePolicyAllowedKeyTypesObject)
		}
		rIssuancePolicy["allowedKeyTypes"] = rIssuancePolicyAllowedKeyTypes
		if r.IssuancePolicy.BaselineValues != nil && r.IssuancePolicy.BaselineValues != dclService.EmptyCaPoolIssuancePolicyBaselineValues {
			rIssuancePolicyBaselineValues := make(map[string]interface{})
			var rIssuancePolicyBaselineValuesAdditionalExtensions []interface{}
			for _, rIssuancePolicyBaselineValuesAdditionalExtensionsVal := range r.IssuancePolicy.BaselineValues.AdditionalExtensions {
				rIssuancePolicyBaselineValuesAdditionalExtensionsObject := make(map[string]interface{})
				if rIssuancePolicyBaselineValuesAdditionalExtensionsVal.Critical != nil {
					rIssuancePolicyBaselineValuesAdditionalExtensionsObject["critical"] = *rIssuancePolicyBaselineValuesAdditionalExtensionsVal.Critical
				}
				if rIssuancePolicyBaselineValuesAdditionalExtensionsVal.ObjectId != nil && rIssuancePolicyBaselineValuesAdditionalExtensionsVal.ObjectId != dclService.EmptyCaPoolIssuancePolicyBaselineValuesAdditionalExtensionsObjectId {
					rIssuancePolicyBaselineValuesAdditionalExtensionsValObjectId := make(map[string]interface{})
					var rIssuancePolicyBaselineValuesAdditionalExtensionsValObjectIdObjectIdPath []interface{}
					for _, rIssuancePolicyBaselineValuesAdditionalExtensionsValObjectIdObjectIdPathVal := range rIssuancePolicyBaselineValuesAdditionalExtensionsVal.ObjectId.ObjectIdPath {
						rIssuancePolicyBaselineValuesAdditionalExtensionsValObjectIdObjectIdPath = append(rIssuancePolicyBaselineValuesAdditionalExtensionsValObjectIdObjectIdPath, rIssuancePolicyBaselineValuesAdditionalExtensionsValObjectIdObjectIdPathVal)
					}
					rIssuancePolicyBaselineValuesAdditionalExtensionsValObjectId["objectIdPath"] = rIssuancePolicyBaselineValuesAdditionalExtensionsValObjectIdObjectIdPath
					rIssuancePolicyBaselineValuesAdditionalExtensionsObject["objectId"] = rIssuancePolicyBaselineValuesAdditionalExtensionsValObjectId
				}
				if rIssuancePolicyBaselineValuesAdditionalExtensionsVal.Value != nil {
					rIssuancePolicyBaselineValuesAdditionalExtensionsObject["value"] = *rIssuancePolicyBaselineValuesAdditionalExtensionsVal.Value
				}
				rIssuancePolicyBaselineValuesAdditionalExtensions = append(rIssuancePolicyBaselineValuesAdditionalExtensions, rIssuancePolicyBaselineValuesAdditionalExtensionsObject)
			}
			rIssuancePolicyBaselineValues["additionalExtensions"] = rIssuancePolicyBaselineValuesAdditionalExtensions
			var rIssuancePolicyBaselineValuesAiaOcspServers []interface{}
			for _, rIssuancePolicyBaselineValuesAiaOcspServersVal := range r.IssuancePolicy.BaselineValues.AiaOcspServers {
				rIssuancePolicyBaselineValuesAiaOcspServers = append(rIssuancePolicyBaselineValuesAiaOcspServers, rIssuancePolicyBaselineValuesAiaOcspServersVal)
			}
			rIssuancePolicyBaselineValues["aiaOcspServers"] = rIssuancePolicyBaselineValuesAiaOcspServers
			if r.IssuancePolicy.BaselineValues.CaOptions != nil && r.IssuancePolicy.BaselineValues.CaOptions != dclService.EmptyCaPoolIssuancePolicyBaselineValuesCaOptions {
				rIssuancePolicyBaselineValuesCaOptions := make(map[string]interface{})
				if r.IssuancePolicy.BaselineValues.CaOptions.IsCa != nil {
					rIssuancePolicyBaselineValuesCaOptions["isCa"] = *r.IssuancePolicy.BaselineValues.CaOptions.IsCa
				}
				if r.IssuancePolicy.BaselineValues.CaOptions.MaxIssuerPathLength != nil {
					rIssuancePolicyBaselineValuesCaOptions["maxIssuerPathLength"] = *r.IssuancePolicy.BaselineValues.CaOptions.MaxIssuerPathLength
				}
				if r.IssuancePolicy.BaselineValues.CaOptions.ZeroMaxIssuerPathLength != nil {
					rIssuancePolicyBaselineValuesCaOptions["zeroMaxIssuerPathLength"] = *r.IssuancePolicy.BaselineValues.CaOptions.ZeroMaxIssuerPathLength
				}
				rIssuancePolicyBaselineValues["caOptions"] = rIssuancePolicyBaselineValuesCaOptions
			}
			if r.IssuancePolicy.BaselineValues.KeyUsage != nil && r.IssuancePolicy.BaselineValues.KeyUsage != dclService.EmptyCaPoolIssuancePolicyBaselineValuesKeyUsage {
				rIssuancePolicyBaselineValuesKeyUsage := make(map[string]interface{})
				if r.IssuancePolicy.BaselineValues.KeyUsage.BaseKeyUsage != nil && r.IssuancePolicy.BaselineValues.KeyUsage.BaseKeyUsage != dclService.EmptyCaPoolIssuancePolicyBaselineValuesKeyUsageBaseKeyUsage {
					rIssuancePolicyBaselineValuesKeyUsageBaseKeyUsage := make(map[string]interface{})
					if r.IssuancePolicy.BaselineValues.KeyUsage.BaseKeyUsage.CertSign != nil {
						rIssuancePolicyBaselineValuesKeyUsageBaseKeyUsage["certSign"] = *r.IssuancePolicy.BaselineValues.KeyUsage.BaseKeyUsage.CertSign
					}
					if r.IssuancePolicy.BaselineValues.KeyUsage.BaseKeyUsage.ContentCommitment != nil {
						rIssuancePolicyBaselineValuesKeyUsageBaseKeyUsage["contentCommitment"] = *r.IssuancePolicy.BaselineValues.KeyUsage.BaseKeyUsage.ContentCommitment
					}
					if r.IssuancePolicy.BaselineValues.KeyUsage.BaseKeyUsage.CrlSign != nil {
						rIssuancePolicyBaselineValuesKeyUsageBaseKeyUsage["crlSign"] = *r.IssuancePolicy.BaselineValues.KeyUsage.BaseKeyUsage.CrlSign
					}
					if r.IssuancePolicy.BaselineValues.KeyUsage.BaseKeyUsage.DataEncipherment != nil {
						rIssuancePolicyBaselineValuesKeyUsageBaseKeyUsage["dataEncipherment"] = *r.IssuancePolicy.BaselineValues.KeyUsage.BaseKeyUsage.DataEncipherment
					}
					if r.IssuancePolicy.BaselineValues.KeyUsage.BaseKeyUsage.DecipherOnly != nil {
						rIssuancePolicyBaselineValuesKeyUsageBaseKeyUsage["decipherOnly"] = *r.IssuancePolicy.BaselineValues.KeyUsage.BaseKeyUsage.DecipherOnly
					}
					if r.IssuancePolicy.BaselineValues.KeyUsage.BaseKeyUsage.DigitalSignature != nil {
						rIssuancePolicyBaselineValuesKeyUsageBaseKeyUsage["digitalSignature"] = *r.IssuancePolicy.BaselineValues.KeyUsage.BaseKeyUsage.DigitalSignature
					}
					if r.IssuancePolicy.BaselineValues.KeyUsage.BaseKeyUsage.EncipherOnly != nil {
						rIssuancePolicyBaselineValuesKeyUsageBaseKeyUsage["encipherOnly"] = *r.IssuancePolicy.BaselineValues.KeyUsage.BaseKeyUsage.EncipherOnly
					}
					if r.IssuancePolicy.BaselineValues.KeyUsage.BaseKeyUsage.KeyAgreement != nil {
						rIssuancePolicyBaselineValuesKeyUsageBaseKeyUsage["keyAgreement"] = *r.IssuancePolicy.BaselineValues.KeyUsage.BaseKeyUsage.KeyAgreement
					}
					if r.IssuancePolicy.BaselineValues.KeyUsage.BaseKeyUsage.KeyEncipherment != nil {
						rIssuancePolicyBaselineValuesKeyUsageBaseKeyUsage["keyEncipherment"] = *r.IssuancePolicy.BaselineValues.KeyUsage.BaseKeyUsage.KeyEncipherment
					}
					rIssuancePolicyBaselineValuesKeyUsage["baseKeyUsage"] = rIssuancePolicyBaselineValuesKeyUsageBaseKeyUsage
				}
				if r.IssuancePolicy.BaselineValues.KeyUsage.ExtendedKeyUsage != nil && r.IssuancePolicy.BaselineValues.KeyUsage.ExtendedKeyUsage != dclService.EmptyCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsage {
					rIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsage := make(map[string]interface{})
					if r.IssuancePolicy.BaselineValues.KeyUsage.ExtendedKeyUsage.ClientAuth != nil {
						rIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsage["clientAuth"] = *r.IssuancePolicy.BaselineValues.KeyUsage.ExtendedKeyUsage.ClientAuth
					}
					if r.IssuancePolicy.BaselineValues.KeyUsage.ExtendedKeyUsage.CodeSigning != nil {
						rIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsage["codeSigning"] = *r.IssuancePolicy.BaselineValues.KeyUsage.ExtendedKeyUsage.CodeSigning
					}
					if r.IssuancePolicy.BaselineValues.KeyUsage.ExtendedKeyUsage.EmailProtection != nil {
						rIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsage["emailProtection"] = *r.IssuancePolicy.BaselineValues.KeyUsage.ExtendedKeyUsage.EmailProtection
					}
					if r.IssuancePolicy.BaselineValues.KeyUsage.ExtendedKeyUsage.OcspSigning != nil {
						rIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsage["ocspSigning"] = *r.IssuancePolicy.BaselineValues.KeyUsage.ExtendedKeyUsage.OcspSigning
					}
					if r.IssuancePolicy.BaselineValues.KeyUsage.ExtendedKeyUsage.ServerAuth != nil {
						rIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsage["serverAuth"] = *r.IssuancePolicy.BaselineValues.KeyUsage.ExtendedKeyUsage.ServerAuth
					}
					if r.IssuancePolicy.BaselineValues.KeyUsage.ExtendedKeyUsage.TimeStamping != nil {
						rIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsage["timeStamping"] = *r.IssuancePolicy.BaselineValues.KeyUsage.ExtendedKeyUsage.TimeStamping
					}
					rIssuancePolicyBaselineValuesKeyUsage["extendedKeyUsage"] = rIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsage
				}
				var rIssuancePolicyBaselineValuesKeyUsageUnknownExtendedKeyUsages []interface{}
				for _, rIssuancePolicyBaselineValuesKeyUsageUnknownExtendedKeyUsagesVal := range r.IssuancePolicy.BaselineValues.KeyUsage.UnknownExtendedKeyUsages {
					rIssuancePolicyBaselineValuesKeyUsageUnknownExtendedKeyUsagesObject := make(map[string]interface{})
					var rIssuancePolicyBaselineValuesKeyUsageUnknownExtendedKeyUsagesValObjectIdPath []interface{}
					for _, rIssuancePolicyBaselineValuesKeyUsageUnknownExtendedKeyUsagesValObjectIdPathVal := range rIssuancePolicyBaselineValuesKeyUsageUnknownExtendedKeyUsagesVal.ObjectIdPath {
						rIssuancePolicyBaselineValuesKeyUsageUnknownExtendedKeyUsagesValObjectIdPath = append(rIssuancePolicyBaselineValuesKeyUsageUnknownExtendedKeyUsagesValObjectIdPath, rIssuancePolicyBaselineValuesKeyUsageUnknownExtendedKeyUsagesValObjectIdPathVal)
					}
					rIssuancePolicyBaselineValuesKeyUsageUnknownExtendedKeyUsagesObject["objectIdPath"] = rIssuancePolicyBaselineValuesKeyUsageUnknownExtendedKeyUsagesValObjectIdPath
					rIssuancePolicyBaselineValuesKeyUsageUnknownExtendedKeyUsages = append(rIssuancePolicyBaselineValuesKeyUsageUnknownExtendedKeyUsages, rIssuancePolicyBaselineValuesKeyUsageUnknownExtendedKeyUsagesObject)
				}
				rIssuancePolicyBaselineValuesKeyUsage["unknownExtendedKeyUsages"] = rIssuancePolicyBaselineValuesKeyUsageUnknownExtendedKeyUsages
				rIssuancePolicyBaselineValues["keyUsage"] = rIssuancePolicyBaselineValuesKeyUsage
			}
			var rIssuancePolicyBaselineValuesPolicyIds []interface{}
			for _, rIssuancePolicyBaselineValuesPolicyIdsVal := range r.IssuancePolicy.BaselineValues.PolicyIds {
				rIssuancePolicyBaselineValuesPolicyIdsObject := make(map[string]interface{})
				var rIssuancePolicyBaselineValuesPolicyIdsValObjectIdPath []interface{}
				for _, rIssuancePolicyBaselineValuesPolicyIdsValObjectIdPathVal := range rIssuancePolicyBaselineValuesPolicyIdsVal.ObjectIdPath {
					rIssuancePolicyBaselineValuesPolicyIdsValObjectIdPath = append(rIssuancePolicyBaselineValuesPolicyIdsValObjectIdPath, rIssuancePolicyBaselineValuesPolicyIdsValObjectIdPathVal)
				}
				rIssuancePolicyBaselineValuesPolicyIdsObject["objectIdPath"] = rIssuancePolicyBaselineValuesPolicyIdsValObjectIdPath
				rIssuancePolicyBaselineValuesPolicyIds = append(rIssuancePolicyBaselineValuesPolicyIds, rIssuancePolicyBaselineValuesPolicyIdsObject)
			}
			rIssuancePolicyBaselineValues["policyIds"] = rIssuancePolicyBaselineValuesPolicyIds
			rIssuancePolicy["baselineValues"] = rIssuancePolicyBaselineValues
		}
		if r.IssuancePolicy.IdentityConstraints != nil && r.IssuancePolicy.IdentityConstraints != dclService.EmptyCaPoolIssuancePolicyIdentityConstraints {
			rIssuancePolicyIdentityConstraints := make(map[string]interface{})
			if r.IssuancePolicy.IdentityConstraints.AllowSubjectAltNamesPassthrough != nil {
				rIssuancePolicyIdentityConstraints["allowSubjectAltNamesPassthrough"] = *r.IssuancePolicy.IdentityConstraints.AllowSubjectAltNamesPassthrough
			}
			if r.IssuancePolicy.IdentityConstraints.AllowSubjectPassthrough != nil {
				rIssuancePolicyIdentityConstraints["allowSubjectPassthrough"] = *r.IssuancePolicy.IdentityConstraints.AllowSubjectPassthrough
			}
			if r.IssuancePolicy.IdentityConstraints.CelExpression != nil && r.IssuancePolicy.IdentityConstraints.CelExpression != dclService.EmptyCaPoolIssuancePolicyIdentityConstraintsCelExpression {
				rIssuancePolicyIdentityConstraintsCelExpression := make(map[string]interface{})
				if r.IssuancePolicy.IdentityConstraints.CelExpression.Description != nil {
					rIssuancePolicyIdentityConstraintsCelExpression["description"] = *r.IssuancePolicy.IdentityConstraints.CelExpression.Description
				}
				if r.IssuancePolicy.IdentityConstraints.CelExpression.Expression != nil {
					rIssuancePolicyIdentityConstraintsCelExpression["expression"] = *r.IssuancePolicy.IdentityConstraints.CelExpression.Expression
				}
				if r.IssuancePolicy.IdentityConstraints.CelExpression.Location != nil {
					rIssuancePolicyIdentityConstraintsCelExpression["location"] = *r.IssuancePolicy.IdentityConstraints.CelExpression.Location
				}
				if r.IssuancePolicy.IdentityConstraints.CelExpression.Title != nil {
					rIssuancePolicyIdentityConstraintsCelExpression["title"] = *r.IssuancePolicy.IdentityConstraints.CelExpression.Title
				}
				rIssuancePolicyIdentityConstraints["celExpression"] = rIssuancePolicyIdentityConstraintsCelExpression
			}
			rIssuancePolicy["identityConstraints"] = rIssuancePolicyIdentityConstraints
		}
		if r.IssuancePolicy.MaximumLifetime != nil {
			rIssuancePolicy["maximumLifetime"] = *r.IssuancePolicy.MaximumLifetime
		}
		if r.IssuancePolicy.PassthroughExtensions != nil && r.IssuancePolicy.PassthroughExtensions != dclService.EmptyCaPoolIssuancePolicyPassthroughExtensions {
			rIssuancePolicyPassthroughExtensions := make(map[string]interface{})
			var rIssuancePolicyPassthroughExtensionsAdditionalExtensions []interface{}
			for _, rIssuancePolicyPassthroughExtensionsAdditionalExtensionsVal := range r.IssuancePolicy.PassthroughExtensions.AdditionalExtensions {
				rIssuancePolicyPassthroughExtensionsAdditionalExtensionsObject := make(map[string]interface{})
				var rIssuancePolicyPassthroughExtensionsAdditionalExtensionsValObjectIdPath []interface{}
				for _, rIssuancePolicyPassthroughExtensionsAdditionalExtensionsValObjectIdPathVal := range rIssuancePolicyPassthroughExtensionsAdditionalExtensionsVal.ObjectIdPath {
					rIssuancePolicyPassthroughExtensionsAdditionalExtensionsValObjectIdPath = append(rIssuancePolicyPassthroughExtensionsAdditionalExtensionsValObjectIdPath, rIssuancePolicyPassthroughExtensionsAdditionalExtensionsValObjectIdPathVal)
				}
				rIssuancePolicyPassthroughExtensionsAdditionalExtensionsObject["objectIdPath"] = rIssuancePolicyPassthroughExtensionsAdditionalExtensionsValObjectIdPath
				rIssuancePolicyPassthroughExtensionsAdditionalExtensions = append(rIssuancePolicyPassthroughExtensionsAdditionalExtensions, rIssuancePolicyPassthroughExtensionsAdditionalExtensionsObject)
			}
			rIssuancePolicyPassthroughExtensions["additionalExtensions"] = rIssuancePolicyPassthroughExtensionsAdditionalExtensions
			var rIssuancePolicyPassthroughExtensionsKnownExtensions []interface{}
			for _, rIssuancePolicyPassthroughExtensionsKnownExtensionsVal := range r.IssuancePolicy.PassthroughExtensions.KnownExtensions {
				rIssuancePolicyPassthroughExtensionsKnownExtensions = append(rIssuancePolicyPassthroughExtensionsKnownExtensions, string(rIssuancePolicyPassthroughExtensionsKnownExtensionsVal))
			}
			rIssuancePolicyPassthroughExtensions["knownExtensions"] = rIssuancePolicyPassthroughExtensionsKnownExtensions
			rIssuancePolicy["passthroughExtensions"] = rIssuancePolicyPassthroughExtensions
		}
		u.Object["issuancePolicy"] = rIssuancePolicy
	}
	if r.Labels != nil {
		rLabels := make(map[string]interface{})
		for k, v := range r.Labels {
			rLabels[k] = v
		}
		u.Object["labels"] = rLabels
	}
	if r.Location != nil {
		u.Object["location"] = *r.Location
	}
	if r.Name != nil {
		u.Object["name"] = *r.Name
	}
	if r.Project != nil {
		u.Object["project"] = *r.Project
	}
	if r.PublishingOptions != nil && r.PublishingOptions != dclService.EmptyCaPoolPublishingOptions {
		rPublishingOptions := make(map[string]interface{})
		if r.PublishingOptions.PublishCaCert != nil {
			rPublishingOptions["publishCaCert"] = *r.PublishingOptions.PublishCaCert
		}
		if r.PublishingOptions.PublishCrl != nil {
			rPublishingOptions["publishCrl"] = *r.PublishingOptions.PublishCrl
		}
		u.Object["publishingOptions"] = rPublishingOptions
	}
	if r.Tier != nil {
		u.Object["tier"] = string(*r.Tier)
	}
	return u
}

func UnstructuredToCaPool(u *unstructured.Resource) (*dclService.CaPool, error) {
	r := &dclService.CaPool{}
	if _, ok := u.Object["issuancePolicy"]; ok {
		if rIssuancePolicy, ok := u.Object["issuancePolicy"].(map[string]interface{}); ok {
			r.IssuancePolicy = &dclService.CaPoolIssuancePolicy{}
			if _, ok := rIssuancePolicy["allowedIssuanceModes"]; ok {
				if rIssuancePolicyAllowedIssuanceModes, ok := rIssuancePolicy["allowedIssuanceModes"].(map[string]interface{}); ok {
					r.IssuancePolicy.AllowedIssuanceModes = &dclService.CaPoolIssuancePolicyAllowedIssuanceModes{}
					if _, ok := rIssuancePolicyAllowedIssuanceModes["allowConfigBasedIssuance"]; ok {
						if b, ok := rIssuancePolicyAllowedIssuanceModes["allowConfigBasedIssuance"].(bool); ok {
							r.IssuancePolicy.AllowedIssuanceModes.AllowConfigBasedIssuance = dcl.Bool(b)
						} else {
							return nil, fmt.Errorf("r.IssuancePolicy.AllowedIssuanceModes.AllowConfigBasedIssuance: expected bool")
						}
					}
					if _, ok := rIssuancePolicyAllowedIssuanceModes["allowCsrBasedIssuance"]; ok {
						if b, ok := rIssuancePolicyAllowedIssuanceModes["allowCsrBasedIssuance"].(bool); ok {
							r.IssuancePolicy.AllowedIssuanceModes.AllowCsrBasedIssuance = dcl.Bool(b)
						} else {
							return nil, fmt.Errorf("r.IssuancePolicy.AllowedIssuanceModes.AllowCsrBasedIssuance: expected bool")
						}
					}
				} else {
					return nil, fmt.Errorf("r.IssuancePolicy.AllowedIssuanceModes: expected map[string]interface{}")
				}
			}
			if _, ok := rIssuancePolicy["allowedKeyTypes"]; ok {
				if s, ok := rIssuancePolicy["allowedKeyTypes"].([]interface{}); ok {
					for _, o := range s {
						if objval, ok := o.(map[string]interface{}); ok {
							var rIssuancePolicyAllowedKeyTypes dclService.CaPoolIssuancePolicyAllowedKeyTypes
							if _, ok := objval["ellipticCurve"]; ok {
								if rIssuancePolicyAllowedKeyTypesEllipticCurve, ok := objval["ellipticCurve"].(map[string]interface{}); ok {
									rIssuancePolicyAllowedKeyTypes.EllipticCurve = &dclService.CaPoolIssuancePolicyAllowedKeyTypesEllipticCurve{}
									if _, ok := rIssuancePolicyAllowedKeyTypesEllipticCurve["signatureAlgorithm"]; ok {
										if s, ok := rIssuancePolicyAllowedKeyTypesEllipticCurve["signatureAlgorithm"].(string); ok {
											rIssuancePolicyAllowedKeyTypes.EllipticCurve.SignatureAlgorithm = dclService.CaPoolIssuancePolicyAllowedKeyTypesEllipticCurveSignatureAlgorithmEnumRef(s)
										} else {
											return nil, fmt.Errorf("rIssuancePolicyAllowedKeyTypes.EllipticCurve.SignatureAlgorithm: expected string")
										}
									}
								} else {
									return nil, fmt.Errorf("rIssuancePolicyAllowedKeyTypes.EllipticCurve: expected map[string]interface{}")
								}
							}
							if _, ok := objval["rsa"]; ok {
								if rIssuancePolicyAllowedKeyTypesRsa, ok := objval["rsa"].(map[string]interface{}); ok {
									rIssuancePolicyAllowedKeyTypes.Rsa = &dclService.CaPoolIssuancePolicyAllowedKeyTypesRsa{}
									if _, ok := rIssuancePolicyAllowedKeyTypesRsa["maxModulusSize"]; ok {
										if i, ok := rIssuancePolicyAllowedKeyTypesRsa["maxModulusSize"].(int64); ok {
											rIssuancePolicyAllowedKeyTypes.Rsa.MaxModulusSize = dcl.Int64(i)
										} else {
											return nil, fmt.Errorf("rIssuancePolicyAllowedKeyTypes.Rsa.MaxModulusSize: expected int64")
										}
									}
									if _, ok := rIssuancePolicyAllowedKeyTypesRsa["minModulusSize"]; ok {
										if i, ok := rIssuancePolicyAllowedKeyTypesRsa["minModulusSize"].(int64); ok {
											rIssuancePolicyAllowedKeyTypes.Rsa.MinModulusSize = dcl.Int64(i)
										} else {
											return nil, fmt.Errorf("rIssuancePolicyAllowedKeyTypes.Rsa.MinModulusSize: expected int64")
										}
									}
								} else {
									return nil, fmt.Errorf("rIssuancePolicyAllowedKeyTypes.Rsa: expected map[string]interface{}")
								}
							}
							r.IssuancePolicy.AllowedKeyTypes = append(r.IssuancePolicy.AllowedKeyTypes, rIssuancePolicyAllowedKeyTypes)
						}
					}
				} else {
					return nil, fmt.Errorf("r.IssuancePolicy.AllowedKeyTypes: expected []interface{}")
				}
			}
			if _, ok := rIssuancePolicy["baselineValues"]; ok {
				if rIssuancePolicyBaselineValues, ok := rIssuancePolicy["baselineValues"].(map[string]interface{}); ok {
					r.IssuancePolicy.BaselineValues = &dclService.CaPoolIssuancePolicyBaselineValues{}
					if _, ok := rIssuancePolicyBaselineValues["additionalExtensions"]; ok {
						if s, ok := rIssuancePolicyBaselineValues["additionalExtensions"].([]interface{}); ok {
							for _, o := range s {
								if objval, ok := o.(map[string]interface{}); ok {
									var rIssuancePolicyBaselineValuesAdditionalExtensions dclService.CaPoolIssuancePolicyBaselineValuesAdditionalExtensions
									if _, ok := objval["critical"]; ok {
										if b, ok := objval["critical"].(bool); ok {
											rIssuancePolicyBaselineValuesAdditionalExtensions.Critical = dcl.Bool(b)
										} else {
											return nil, fmt.Errorf("rIssuancePolicyBaselineValuesAdditionalExtensions.Critical: expected bool")
										}
									}
									if _, ok := objval["objectId"]; ok {
										if rIssuancePolicyBaselineValuesAdditionalExtensionsObjectId, ok := objval["objectId"].(map[string]interface{}); ok {
											rIssuancePolicyBaselineValuesAdditionalExtensions.ObjectId = &dclService.CaPoolIssuancePolicyBaselineValuesAdditionalExtensionsObjectId{}
											if _, ok := rIssuancePolicyBaselineValuesAdditionalExtensionsObjectId["objectIdPath"]; ok {
												if s, ok := rIssuancePolicyBaselineValuesAdditionalExtensionsObjectId["objectIdPath"].([]interface{}); ok {
													for _, ss := range s {
														if intval, ok := ss.(int64); ok {
															rIssuancePolicyBaselineValuesAdditionalExtensions.ObjectId.ObjectIdPath = append(rIssuancePolicyBaselineValuesAdditionalExtensions.ObjectId.ObjectIdPath, intval)
														}
													}
												} else {
													return nil, fmt.Errorf("rIssuancePolicyBaselineValuesAdditionalExtensions.ObjectId.ObjectIdPath: expected []interface{}")
												}
											}
										} else {
											return nil, fmt.Errorf("rIssuancePolicyBaselineValuesAdditionalExtensions.ObjectId: expected map[string]interface{}")
										}
									}
									if _, ok := objval["value"]; ok {
										if s, ok := objval["value"].(string); ok {
											rIssuancePolicyBaselineValuesAdditionalExtensions.Value = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rIssuancePolicyBaselineValuesAdditionalExtensions.Value: expected string")
										}
									}
									r.IssuancePolicy.BaselineValues.AdditionalExtensions = append(r.IssuancePolicy.BaselineValues.AdditionalExtensions, rIssuancePolicyBaselineValuesAdditionalExtensions)
								}
							}
						} else {
							return nil, fmt.Errorf("r.IssuancePolicy.BaselineValues.AdditionalExtensions: expected []interface{}")
						}
					}
					if _, ok := rIssuancePolicyBaselineValues["aiaOcspServers"]; ok {
						if s, ok := rIssuancePolicyBaselineValues["aiaOcspServers"].([]interface{}); ok {
							for _, ss := range s {
								if strval, ok := ss.(string); ok {
									r.IssuancePolicy.BaselineValues.AiaOcspServers = append(r.IssuancePolicy.BaselineValues.AiaOcspServers, strval)
								}
							}
						} else {
							return nil, fmt.Errorf("r.IssuancePolicy.BaselineValues.AiaOcspServers: expected []interface{}")
						}
					}
					if _, ok := rIssuancePolicyBaselineValues["caOptions"]; ok {
						if rIssuancePolicyBaselineValuesCaOptions, ok := rIssuancePolicyBaselineValues["caOptions"].(map[string]interface{}); ok {
							r.IssuancePolicy.BaselineValues.CaOptions = &dclService.CaPoolIssuancePolicyBaselineValuesCaOptions{}
							if _, ok := rIssuancePolicyBaselineValuesCaOptions["isCa"]; ok {
								if b, ok := rIssuancePolicyBaselineValuesCaOptions["isCa"].(bool); ok {
									r.IssuancePolicy.BaselineValues.CaOptions.IsCa = dcl.Bool(b)
								} else {
									return nil, fmt.Errorf("r.IssuancePolicy.BaselineValues.CaOptions.IsCa: expected bool")
								}
							}
							if _, ok := rIssuancePolicyBaselineValuesCaOptions["maxIssuerPathLength"]; ok {
								if i, ok := rIssuancePolicyBaselineValuesCaOptions["maxIssuerPathLength"].(int64); ok {
									r.IssuancePolicy.BaselineValues.CaOptions.MaxIssuerPathLength = dcl.Int64(i)
								} else {
									return nil, fmt.Errorf("r.IssuancePolicy.BaselineValues.CaOptions.MaxIssuerPathLength: expected int64")
								}
							}
							if _, ok := rIssuancePolicyBaselineValuesCaOptions["zeroMaxIssuerPathLength"]; ok {
								if b, ok := rIssuancePolicyBaselineValuesCaOptions["zeroMaxIssuerPathLength"].(bool); ok {
									r.IssuancePolicy.BaselineValues.CaOptions.ZeroMaxIssuerPathLength = dcl.Bool(b)
								} else {
									return nil, fmt.Errorf("r.IssuancePolicy.BaselineValues.CaOptions.ZeroMaxIssuerPathLength: expected bool")
								}
							}
						} else {
							return nil, fmt.Errorf("r.IssuancePolicy.BaselineValues.CaOptions: expected map[string]interface{}")
						}
					}
					if _, ok := rIssuancePolicyBaselineValues["keyUsage"]; ok {
						if rIssuancePolicyBaselineValuesKeyUsage, ok := rIssuancePolicyBaselineValues["keyUsage"].(map[string]interface{}); ok {
							r.IssuancePolicy.BaselineValues.KeyUsage = &dclService.CaPoolIssuancePolicyBaselineValuesKeyUsage{}
							if _, ok := rIssuancePolicyBaselineValuesKeyUsage["baseKeyUsage"]; ok {
								if rIssuancePolicyBaselineValuesKeyUsageBaseKeyUsage, ok := rIssuancePolicyBaselineValuesKeyUsage["baseKeyUsage"].(map[string]interface{}); ok {
									r.IssuancePolicy.BaselineValues.KeyUsage.BaseKeyUsage = &dclService.CaPoolIssuancePolicyBaselineValuesKeyUsageBaseKeyUsage{}
									if _, ok := rIssuancePolicyBaselineValuesKeyUsageBaseKeyUsage["certSign"]; ok {
										if b, ok := rIssuancePolicyBaselineValuesKeyUsageBaseKeyUsage["certSign"].(bool); ok {
											r.IssuancePolicy.BaselineValues.KeyUsage.BaseKeyUsage.CertSign = dcl.Bool(b)
										} else {
											return nil, fmt.Errorf("r.IssuancePolicy.BaselineValues.KeyUsage.BaseKeyUsage.CertSign: expected bool")
										}
									}
									if _, ok := rIssuancePolicyBaselineValuesKeyUsageBaseKeyUsage["contentCommitment"]; ok {
										if b, ok := rIssuancePolicyBaselineValuesKeyUsageBaseKeyUsage["contentCommitment"].(bool); ok {
											r.IssuancePolicy.BaselineValues.KeyUsage.BaseKeyUsage.ContentCommitment = dcl.Bool(b)
										} else {
											return nil, fmt.Errorf("r.IssuancePolicy.BaselineValues.KeyUsage.BaseKeyUsage.ContentCommitment: expected bool")
										}
									}
									if _, ok := rIssuancePolicyBaselineValuesKeyUsageBaseKeyUsage["crlSign"]; ok {
										if b, ok := rIssuancePolicyBaselineValuesKeyUsageBaseKeyUsage["crlSign"].(bool); ok {
											r.IssuancePolicy.BaselineValues.KeyUsage.BaseKeyUsage.CrlSign = dcl.Bool(b)
										} else {
											return nil, fmt.Errorf("r.IssuancePolicy.BaselineValues.KeyUsage.BaseKeyUsage.CrlSign: expected bool")
										}
									}
									if _, ok := rIssuancePolicyBaselineValuesKeyUsageBaseKeyUsage["dataEncipherment"]; ok {
										if b, ok := rIssuancePolicyBaselineValuesKeyUsageBaseKeyUsage["dataEncipherment"].(bool); ok {
											r.IssuancePolicy.BaselineValues.KeyUsage.BaseKeyUsage.DataEncipherment = dcl.Bool(b)
										} else {
											return nil, fmt.Errorf("r.IssuancePolicy.BaselineValues.KeyUsage.BaseKeyUsage.DataEncipherment: expected bool")
										}
									}
									if _, ok := rIssuancePolicyBaselineValuesKeyUsageBaseKeyUsage["decipherOnly"]; ok {
										if b, ok := rIssuancePolicyBaselineValuesKeyUsageBaseKeyUsage["decipherOnly"].(bool); ok {
											r.IssuancePolicy.BaselineValues.KeyUsage.BaseKeyUsage.DecipherOnly = dcl.Bool(b)
										} else {
											return nil, fmt.Errorf("r.IssuancePolicy.BaselineValues.KeyUsage.BaseKeyUsage.DecipherOnly: expected bool")
										}
									}
									if _, ok := rIssuancePolicyBaselineValuesKeyUsageBaseKeyUsage["digitalSignature"]; ok {
										if b, ok := rIssuancePolicyBaselineValuesKeyUsageBaseKeyUsage["digitalSignature"].(bool); ok {
											r.IssuancePolicy.BaselineValues.KeyUsage.BaseKeyUsage.DigitalSignature = dcl.Bool(b)
										} else {
											return nil, fmt.Errorf("r.IssuancePolicy.BaselineValues.KeyUsage.BaseKeyUsage.DigitalSignature: expected bool")
										}
									}
									if _, ok := rIssuancePolicyBaselineValuesKeyUsageBaseKeyUsage["encipherOnly"]; ok {
										if b, ok := rIssuancePolicyBaselineValuesKeyUsageBaseKeyUsage["encipherOnly"].(bool); ok {
											r.IssuancePolicy.BaselineValues.KeyUsage.BaseKeyUsage.EncipherOnly = dcl.Bool(b)
										} else {
											return nil, fmt.Errorf("r.IssuancePolicy.BaselineValues.KeyUsage.BaseKeyUsage.EncipherOnly: expected bool")
										}
									}
									if _, ok := rIssuancePolicyBaselineValuesKeyUsageBaseKeyUsage["keyAgreement"]; ok {
										if b, ok := rIssuancePolicyBaselineValuesKeyUsageBaseKeyUsage["keyAgreement"].(bool); ok {
											r.IssuancePolicy.BaselineValues.KeyUsage.BaseKeyUsage.KeyAgreement = dcl.Bool(b)
										} else {
											return nil, fmt.Errorf("r.IssuancePolicy.BaselineValues.KeyUsage.BaseKeyUsage.KeyAgreement: expected bool")
										}
									}
									if _, ok := rIssuancePolicyBaselineValuesKeyUsageBaseKeyUsage["keyEncipherment"]; ok {
										if b, ok := rIssuancePolicyBaselineValuesKeyUsageBaseKeyUsage["keyEncipherment"].(bool); ok {
											r.IssuancePolicy.BaselineValues.KeyUsage.BaseKeyUsage.KeyEncipherment = dcl.Bool(b)
										} else {
											return nil, fmt.Errorf("r.IssuancePolicy.BaselineValues.KeyUsage.BaseKeyUsage.KeyEncipherment: expected bool")
										}
									}
								} else {
									return nil, fmt.Errorf("r.IssuancePolicy.BaselineValues.KeyUsage.BaseKeyUsage: expected map[string]interface{}")
								}
							}
							if _, ok := rIssuancePolicyBaselineValuesKeyUsage["extendedKeyUsage"]; ok {
								if rIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsage, ok := rIssuancePolicyBaselineValuesKeyUsage["extendedKeyUsage"].(map[string]interface{}); ok {
									r.IssuancePolicy.BaselineValues.KeyUsage.ExtendedKeyUsage = &dclService.CaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsage{}
									if _, ok := rIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsage["clientAuth"]; ok {
										if b, ok := rIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsage["clientAuth"].(bool); ok {
											r.IssuancePolicy.BaselineValues.KeyUsage.ExtendedKeyUsage.ClientAuth = dcl.Bool(b)
										} else {
											return nil, fmt.Errorf("r.IssuancePolicy.BaselineValues.KeyUsage.ExtendedKeyUsage.ClientAuth: expected bool")
										}
									}
									if _, ok := rIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsage["codeSigning"]; ok {
										if b, ok := rIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsage["codeSigning"].(bool); ok {
											r.IssuancePolicy.BaselineValues.KeyUsage.ExtendedKeyUsage.CodeSigning = dcl.Bool(b)
										} else {
											return nil, fmt.Errorf("r.IssuancePolicy.BaselineValues.KeyUsage.ExtendedKeyUsage.CodeSigning: expected bool")
										}
									}
									if _, ok := rIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsage["emailProtection"]; ok {
										if b, ok := rIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsage["emailProtection"].(bool); ok {
											r.IssuancePolicy.BaselineValues.KeyUsage.ExtendedKeyUsage.EmailProtection = dcl.Bool(b)
										} else {
											return nil, fmt.Errorf("r.IssuancePolicy.BaselineValues.KeyUsage.ExtendedKeyUsage.EmailProtection: expected bool")
										}
									}
									if _, ok := rIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsage["ocspSigning"]; ok {
										if b, ok := rIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsage["ocspSigning"].(bool); ok {
											r.IssuancePolicy.BaselineValues.KeyUsage.ExtendedKeyUsage.OcspSigning = dcl.Bool(b)
										} else {
											return nil, fmt.Errorf("r.IssuancePolicy.BaselineValues.KeyUsage.ExtendedKeyUsage.OcspSigning: expected bool")
										}
									}
									if _, ok := rIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsage["serverAuth"]; ok {
										if b, ok := rIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsage["serverAuth"].(bool); ok {
											r.IssuancePolicy.BaselineValues.KeyUsage.ExtendedKeyUsage.ServerAuth = dcl.Bool(b)
										} else {
											return nil, fmt.Errorf("r.IssuancePolicy.BaselineValues.KeyUsage.ExtendedKeyUsage.ServerAuth: expected bool")
										}
									}
									if _, ok := rIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsage["timeStamping"]; ok {
										if b, ok := rIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsage["timeStamping"].(bool); ok {
											r.IssuancePolicy.BaselineValues.KeyUsage.ExtendedKeyUsage.TimeStamping = dcl.Bool(b)
										} else {
											return nil, fmt.Errorf("r.IssuancePolicy.BaselineValues.KeyUsage.ExtendedKeyUsage.TimeStamping: expected bool")
										}
									}
								} else {
									return nil, fmt.Errorf("r.IssuancePolicy.BaselineValues.KeyUsage.ExtendedKeyUsage: expected map[string]interface{}")
								}
							}
							if _, ok := rIssuancePolicyBaselineValuesKeyUsage["unknownExtendedKeyUsages"]; ok {
								if s, ok := rIssuancePolicyBaselineValuesKeyUsage["unknownExtendedKeyUsages"].([]interface{}); ok {
									for _, o := range s {
										if objval, ok := o.(map[string]interface{}); ok {
											var rIssuancePolicyBaselineValuesKeyUsageUnknownExtendedKeyUsages dclService.CaPoolIssuancePolicyBaselineValuesKeyUsageUnknownExtendedKeyUsages
											if _, ok := objval["objectIdPath"]; ok {
												if s, ok := objval["objectIdPath"].([]interface{}); ok {
													for _, ss := range s {
														if intval, ok := ss.(int64); ok {
															rIssuancePolicyBaselineValuesKeyUsageUnknownExtendedKeyUsages.ObjectIdPath = append(rIssuancePolicyBaselineValuesKeyUsageUnknownExtendedKeyUsages.ObjectIdPath, intval)
														}
													}
												} else {
													return nil, fmt.Errorf("rIssuancePolicyBaselineValuesKeyUsageUnknownExtendedKeyUsages.ObjectIdPath: expected []interface{}")
												}
											}
											r.IssuancePolicy.BaselineValues.KeyUsage.UnknownExtendedKeyUsages = append(r.IssuancePolicy.BaselineValues.KeyUsage.UnknownExtendedKeyUsages, rIssuancePolicyBaselineValuesKeyUsageUnknownExtendedKeyUsages)
										}
									}
								} else {
									return nil, fmt.Errorf("r.IssuancePolicy.BaselineValues.KeyUsage.UnknownExtendedKeyUsages: expected []interface{}")
								}
							}
						} else {
							return nil, fmt.Errorf("r.IssuancePolicy.BaselineValues.KeyUsage: expected map[string]interface{}")
						}
					}
					if _, ok := rIssuancePolicyBaselineValues["policyIds"]; ok {
						if s, ok := rIssuancePolicyBaselineValues["policyIds"].([]interface{}); ok {
							for _, o := range s {
								if objval, ok := o.(map[string]interface{}); ok {
									var rIssuancePolicyBaselineValuesPolicyIds dclService.CaPoolIssuancePolicyBaselineValuesPolicyIds
									if _, ok := objval["objectIdPath"]; ok {
										if s, ok := objval["objectIdPath"].([]interface{}); ok {
											for _, ss := range s {
												if intval, ok := ss.(int64); ok {
													rIssuancePolicyBaselineValuesPolicyIds.ObjectIdPath = append(rIssuancePolicyBaselineValuesPolicyIds.ObjectIdPath, intval)
												}
											}
										} else {
											return nil, fmt.Errorf("rIssuancePolicyBaselineValuesPolicyIds.ObjectIdPath: expected []interface{}")
										}
									}
									r.IssuancePolicy.BaselineValues.PolicyIds = append(r.IssuancePolicy.BaselineValues.PolicyIds, rIssuancePolicyBaselineValuesPolicyIds)
								}
							}
						} else {
							return nil, fmt.Errorf("r.IssuancePolicy.BaselineValues.PolicyIds: expected []interface{}")
						}
					}
				} else {
					return nil, fmt.Errorf("r.IssuancePolicy.BaselineValues: expected map[string]interface{}")
				}
			}
			if _, ok := rIssuancePolicy["identityConstraints"]; ok {
				if rIssuancePolicyIdentityConstraints, ok := rIssuancePolicy["identityConstraints"].(map[string]interface{}); ok {
					r.IssuancePolicy.IdentityConstraints = &dclService.CaPoolIssuancePolicyIdentityConstraints{}
					if _, ok := rIssuancePolicyIdentityConstraints["allowSubjectAltNamesPassthrough"]; ok {
						if b, ok := rIssuancePolicyIdentityConstraints["allowSubjectAltNamesPassthrough"].(bool); ok {
							r.IssuancePolicy.IdentityConstraints.AllowSubjectAltNamesPassthrough = dcl.Bool(b)
						} else {
							return nil, fmt.Errorf("r.IssuancePolicy.IdentityConstraints.AllowSubjectAltNamesPassthrough: expected bool")
						}
					}
					if _, ok := rIssuancePolicyIdentityConstraints["allowSubjectPassthrough"]; ok {
						if b, ok := rIssuancePolicyIdentityConstraints["allowSubjectPassthrough"].(bool); ok {
							r.IssuancePolicy.IdentityConstraints.AllowSubjectPassthrough = dcl.Bool(b)
						} else {
							return nil, fmt.Errorf("r.IssuancePolicy.IdentityConstraints.AllowSubjectPassthrough: expected bool")
						}
					}
					if _, ok := rIssuancePolicyIdentityConstraints["celExpression"]; ok {
						if rIssuancePolicyIdentityConstraintsCelExpression, ok := rIssuancePolicyIdentityConstraints["celExpression"].(map[string]interface{}); ok {
							r.IssuancePolicy.IdentityConstraints.CelExpression = &dclService.CaPoolIssuancePolicyIdentityConstraintsCelExpression{}
							if _, ok := rIssuancePolicyIdentityConstraintsCelExpression["description"]; ok {
								if s, ok := rIssuancePolicyIdentityConstraintsCelExpression["description"].(string); ok {
									r.IssuancePolicy.IdentityConstraints.CelExpression.Description = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.IssuancePolicy.IdentityConstraints.CelExpression.Description: expected string")
								}
							}
							if _, ok := rIssuancePolicyIdentityConstraintsCelExpression["expression"]; ok {
								if s, ok := rIssuancePolicyIdentityConstraintsCelExpression["expression"].(string); ok {
									r.IssuancePolicy.IdentityConstraints.CelExpression.Expression = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.IssuancePolicy.IdentityConstraints.CelExpression.Expression: expected string")
								}
							}
							if _, ok := rIssuancePolicyIdentityConstraintsCelExpression["location"]; ok {
								if s, ok := rIssuancePolicyIdentityConstraintsCelExpression["location"].(string); ok {
									r.IssuancePolicy.IdentityConstraints.CelExpression.Location = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.IssuancePolicy.IdentityConstraints.CelExpression.Location: expected string")
								}
							}
							if _, ok := rIssuancePolicyIdentityConstraintsCelExpression["title"]; ok {
								if s, ok := rIssuancePolicyIdentityConstraintsCelExpression["title"].(string); ok {
									r.IssuancePolicy.IdentityConstraints.CelExpression.Title = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.IssuancePolicy.IdentityConstraints.CelExpression.Title: expected string")
								}
							}
						} else {
							return nil, fmt.Errorf("r.IssuancePolicy.IdentityConstraints.CelExpression: expected map[string]interface{}")
						}
					}
				} else {
					return nil, fmt.Errorf("r.IssuancePolicy.IdentityConstraints: expected map[string]interface{}")
				}
			}
			if _, ok := rIssuancePolicy["maximumLifetime"]; ok {
				if s, ok := rIssuancePolicy["maximumLifetime"].(string); ok {
					r.IssuancePolicy.MaximumLifetime = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.IssuancePolicy.MaximumLifetime: expected string")
				}
			}
			if _, ok := rIssuancePolicy["passthroughExtensions"]; ok {
				if rIssuancePolicyPassthroughExtensions, ok := rIssuancePolicy["passthroughExtensions"].(map[string]interface{}); ok {
					r.IssuancePolicy.PassthroughExtensions = &dclService.CaPoolIssuancePolicyPassthroughExtensions{}
					if _, ok := rIssuancePolicyPassthroughExtensions["additionalExtensions"]; ok {
						if s, ok := rIssuancePolicyPassthroughExtensions["additionalExtensions"].([]interface{}); ok {
							for _, o := range s {
								if objval, ok := o.(map[string]interface{}); ok {
									var rIssuancePolicyPassthroughExtensionsAdditionalExtensions dclService.CaPoolIssuancePolicyPassthroughExtensionsAdditionalExtensions
									if _, ok := objval["objectIdPath"]; ok {
										if s, ok := objval["objectIdPath"].([]interface{}); ok {
											for _, ss := range s {
												if intval, ok := ss.(int64); ok {
													rIssuancePolicyPassthroughExtensionsAdditionalExtensions.ObjectIdPath = append(rIssuancePolicyPassthroughExtensionsAdditionalExtensions.ObjectIdPath, intval)
												}
											}
										} else {
											return nil, fmt.Errorf("rIssuancePolicyPassthroughExtensionsAdditionalExtensions.ObjectIdPath: expected []interface{}")
										}
									}
									r.IssuancePolicy.PassthroughExtensions.AdditionalExtensions = append(r.IssuancePolicy.PassthroughExtensions.AdditionalExtensions, rIssuancePolicyPassthroughExtensionsAdditionalExtensions)
								}
							}
						} else {
							return nil, fmt.Errorf("r.IssuancePolicy.PassthroughExtensions.AdditionalExtensions: expected []interface{}")
						}
					}
					if _, ok := rIssuancePolicyPassthroughExtensions["knownExtensions"]; ok {
						if s, ok := rIssuancePolicyPassthroughExtensions["knownExtensions"].([]interface{}); ok {
							for _, ss := range s {
								if strval, ok := ss.(string); ok {
									r.IssuancePolicy.PassthroughExtensions.KnownExtensions = append(r.IssuancePolicy.PassthroughExtensions.KnownExtensions, dclService.CaPoolIssuancePolicyPassthroughExtensionsKnownExtensionsEnum(strval))
								}
							}
						} else {
							return nil, fmt.Errorf("r.IssuancePolicy.PassthroughExtensions.KnownExtensions: expected []interface{}")
						}
					}
				} else {
					return nil, fmt.Errorf("r.IssuancePolicy.PassthroughExtensions: expected map[string]interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.IssuancePolicy: expected map[string]interface{}")
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
	if _, ok := u.Object["project"]; ok {
		if s, ok := u.Object["project"].(string); ok {
			r.Project = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Project: expected string")
		}
	}
	if _, ok := u.Object["publishingOptions"]; ok {
		if rPublishingOptions, ok := u.Object["publishingOptions"].(map[string]interface{}); ok {
			r.PublishingOptions = &dclService.CaPoolPublishingOptions{}
			if _, ok := rPublishingOptions["publishCaCert"]; ok {
				if b, ok := rPublishingOptions["publishCaCert"].(bool); ok {
					r.PublishingOptions.PublishCaCert = dcl.Bool(b)
				} else {
					return nil, fmt.Errorf("r.PublishingOptions.PublishCaCert: expected bool")
				}
			}
			if _, ok := rPublishingOptions["publishCrl"]; ok {
				if b, ok := rPublishingOptions["publishCrl"].(bool); ok {
					r.PublishingOptions.PublishCrl = dcl.Bool(b)
				} else {
					return nil, fmt.Errorf("r.PublishingOptions.PublishCrl: expected bool")
				}
			}
		} else {
			return nil, fmt.Errorf("r.PublishingOptions: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["tier"]; ok {
		if s, ok := u.Object["tier"].(string); ok {
			r.Tier = dclService.CaPoolTierEnumRef(s)
		} else {
			return nil, fmt.Errorf("r.Tier: expected string")
		}
	}
	return r, nil
}

func GetCaPool(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToCaPool(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetCaPool(ctx, r)
	if err != nil {
		return nil, err
	}
	return CaPoolToUnstructured(r), nil
}

func ListCaPool(ctx context.Context, config *dcl.Config, project string, location string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListCaPool(ctx, project, location)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, CaPoolToUnstructured(r))
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

func ApplyCaPool(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToCaPool(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToCaPool(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyCaPool(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return CaPoolToUnstructured(r), nil
}

func CaPoolHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToCaPool(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToCaPool(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyCaPool(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteCaPool(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToCaPool(u)
	if err != nil {
		return err
	}
	return c.DeleteCaPool(ctx, r)
}

func CaPoolID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToCaPool(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *CaPool) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"privateca",
		"CaPool",
		"beta",
	}
}

func (r *CaPool) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *CaPool) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *CaPool) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *CaPool) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *CaPool) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *CaPool) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *CaPool) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetCaPool(ctx, config, resource)
}

func (r *CaPool) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyCaPool(ctx, config, resource, opts...)
}

func (r *CaPool) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return CaPoolHasDiff(ctx, config, resource, opts...)
}

func (r *CaPool) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteCaPool(ctx, config, resource)
}

func (r *CaPool) ID(resource *unstructured.Resource) (string, error) {
	return CaPoolID(resource)
}

func init() {
	unstructured.Register(&CaPool{})
}
