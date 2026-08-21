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

type CertificateTemplate struct{}

func CertificateTemplateToUnstructured(r *dclService.CertificateTemplate) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "privateca",
			Version: "alpha",
			Type:    "CertificateTemplate",
		},
		Object: make(map[string]interface{}),
	}
	if r.CreateTime != nil {
		u.Object["createTime"] = *r.CreateTime
	}
	if r.Description != nil {
		u.Object["description"] = *r.Description
	}
	if r.IdentityConstraints != nil && r.IdentityConstraints != dclService.EmptyCertificateTemplateIdentityConstraints {
		rIdentityConstraints := make(map[string]interface{})
		if r.IdentityConstraints.AllowSubjectAltNamesPassthrough != nil {
			rIdentityConstraints["allowSubjectAltNamesPassthrough"] = *r.IdentityConstraints.AllowSubjectAltNamesPassthrough
		}
		if r.IdentityConstraints.AllowSubjectPassthrough != nil {
			rIdentityConstraints["allowSubjectPassthrough"] = *r.IdentityConstraints.AllowSubjectPassthrough
		}
		if r.IdentityConstraints.CelExpression != nil && r.IdentityConstraints.CelExpression != dclService.EmptyCertificateTemplateIdentityConstraintsCelExpression {
			rIdentityConstraintsCelExpression := make(map[string]interface{})
			if r.IdentityConstraints.CelExpression.Description != nil {
				rIdentityConstraintsCelExpression["description"] = *r.IdentityConstraints.CelExpression.Description
			}
			if r.IdentityConstraints.CelExpression.Expression != nil {
				rIdentityConstraintsCelExpression["expression"] = *r.IdentityConstraints.CelExpression.Expression
			}
			if r.IdentityConstraints.CelExpression.Location != nil {
				rIdentityConstraintsCelExpression["location"] = *r.IdentityConstraints.CelExpression.Location
			}
			if r.IdentityConstraints.CelExpression.Title != nil {
				rIdentityConstraintsCelExpression["title"] = *r.IdentityConstraints.CelExpression.Title
			}
			rIdentityConstraints["celExpression"] = rIdentityConstraintsCelExpression
		}
		u.Object["identityConstraints"] = rIdentityConstraints
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
	if r.PassthroughExtensions != nil && r.PassthroughExtensions != dclService.EmptyCertificateTemplatePassthroughExtensions {
		rPassthroughExtensions := make(map[string]interface{})
		var rPassthroughExtensionsAdditionalExtensions []interface{}
		for _, rPassthroughExtensionsAdditionalExtensionsVal := range r.PassthroughExtensions.AdditionalExtensions {
			rPassthroughExtensionsAdditionalExtensionsObject := make(map[string]interface{})
			var rPassthroughExtensionsAdditionalExtensionsValObjectIdPath []interface{}
			for _, rPassthroughExtensionsAdditionalExtensionsValObjectIdPathVal := range rPassthroughExtensionsAdditionalExtensionsVal.ObjectIdPath {
				rPassthroughExtensionsAdditionalExtensionsValObjectIdPath = append(rPassthroughExtensionsAdditionalExtensionsValObjectIdPath, rPassthroughExtensionsAdditionalExtensionsValObjectIdPathVal)
			}
			rPassthroughExtensionsAdditionalExtensionsObject["objectIdPath"] = rPassthroughExtensionsAdditionalExtensionsValObjectIdPath
			rPassthroughExtensionsAdditionalExtensions = append(rPassthroughExtensionsAdditionalExtensions, rPassthroughExtensionsAdditionalExtensionsObject)
		}
		rPassthroughExtensions["additionalExtensions"] = rPassthroughExtensionsAdditionalExtensions
		var rPassthroughExtensionsKnownExtensions []interface{}
		for _, rPassthroughExtensionsKnownExtensionsVal := range r.PassthroughExtensions.KnownExtensions {
			rPassthroughExtensionsKnownExtensions = append(rPassthroughExtensionsKnownExtensions, string(rPassthroughExtensionsKnownExtensionsVal))
		}
		rPassthroughExtensions["knownExtensions"] = rPassthroughExtensionsKnownExtensions
		u.Object["passthroughExtensions"] = rPassthroughExtensions
	}
	if r.PredefinedValues != nil && r.PredefinedValues != dclService.EmptyCertificateTemplatePredefinedValues {
		rPredefinedValues := make(map[string]interface{})
		var rPredefinedValuesAdditionalExtensions []interface{}
		for _, rPredefinedValuesAdditionalExtensionsVal := range r.PredefinedValues.AdditionalExtensions {
			rPredefinedValuesAdditionalExtensionsObject := make(map[string]interface{})
			if rPredefinedValuesAdditionalExtensionsVal.Critical != nil {
				rPredefinedValuesAdditionalExtensionsObject["critical"] = *rPredefinedValuesAdditionalExtensionsVal.Critical
			}
			if rPredefinedValuesAdditionalExtensionsVal.ObjectId != nil && rPredefinedValuesAdditionalExtensionsVal.ObjectId != dclService.EmptyCertificateTemplatePredefinedValuesAdditionalExtensionsObjectId {
				rPredefinedValuesAdditionalExtensionsValObjectId := make(map[string]interface{})
				var rPredefinedValuesAdditionalExtensionsValObjectIdObjectIdPath []interface{}
				for _, rPredefinedValuesAdditionalExtensionsValObjectIdObjectIdPathVal := range rPredefinedValuesAdditionalExtensionsVal.ObjectId.ObjectIdPath {
					rPredefinedValuesAdditionalExtensionsValObjectIdObjectIdPath = append(rPredefinedValuesAdditionalExtensionsValObjectIdObjectIdPath, rPredefinedValuesAdditionalExtensionsValObjectIdObjectIdPathVal)
				}
				rPredefinedValuesAdditionalExtensionsValObjectId["objectIdPath"] = rPredefinedValuesAdditionalExtensionsValObjectIdObjectIdPath
				rPredefinedValuesAdditionalExtensionsObject["objectId"] = rPredefinedValuesAdditionalExtensionsValObjectId
			}
			if rPredefinedValuesAdditionalExtensionsVal.Value != nil {
				rPredefinedValuesAdditionalExtensionsObject["value"] = *rPredefinedValuesAdditionalExtensionsVal.Value
			}
			rPredefinedValuesAdditionalExtensions = append(rPredefinedValuesAdditionalExtensions, rPredefinedValuesAdditionalExtensionsObject)
		}
		rPredefinedValues["additionalExtensions"] = rPredefinedValuesAdditionalExtensions
		var rPredefinedValuesAiaOcspServers []interface{}
		for _, rPredefinedValuesAiaOcspServersVal := range r.PredefinedValues.AiaOcspServers {
			rPredefinedValuesAiaOcspServers = append(rPredefinedValuesAiaOcspServers, rPredefinedValuesAiaOcspServersVal)
		}
		rPredefinedValues["aiaOcspServers"] = rPredefinedValuesAiaOcspServers
		if r.PredefinedValues.CaOptions != nil && r.PredefinedValues.CaOptions != dclService.EmptyCertificateTemplatePredefinedValuesCaOptions {
			rPredefinedValuesCaOptions := make(map[string]interface{})
			if r.PredefinedValues.CaOptions.IsCa != nil {
				rPredefinedValuesCaOptions["isCa"] = *r.PredefinedValues.CaOptions.IsCa
			}
			if r.PredefinedValues.CaOptions.MaxIssuerPathLength != nil {
				rPredefinedValuesCaOptions["maxIssuerPathLength"] = *r.PredefinedValues.CaOptions.MaxIssuerPathLength
			}
			rPredefinedValues["caOptions"] = rPredefinedValuesCaOptions
		}
		if r.PredefinedValues.KeyUsage != nil && r.PredefinedValues.KeyUsage != dclService.EmptyCertificateTemplatePredefinedValuesKeyUsage {
			rPredefinedValuesKeyUsage := make(map[string]interface{})
			if r.PredefinedValues.KeyUsage.BaseKeyUsage != nil && r.PredefinedValues.KeyUsage.BaseKeyUsage != dclService.EmptyCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage {
				rPredefinedValuesKeyUsageBaseKeyUsage := make(map[string]interface{})
				if r.PredefinedValues.KeyUsage.BaseKeyUsage.CertSign != nil {
					rPredefinedValuesKeyUsageBaseKeyUsage["certSign"] = *r.PredefinedValues.KeyUsage.BaseKeyUsage.CertSign
				}
				if r.PredefinedValues.KeyUsage.BaseKeyUsage.ContentCommitment != nil {
					rPredefinedValuesKeyUsageBaseKeyUsage["contentCommitment"] = *r.PredefinedValues.KeyUsage.BaseKeyUsage.ContentCommitment
				}
				if r.PredefinedValues.KeyUsage.BaseKeyUsage.CrlSign != nil {
					rPredefinedValuesKeyUsageBaseKeyUsage["crlSign"] = *r.PredefinedValues.KeyUsage.BaseKeyUsage.CrlSign
				}
				if r.PredefinedValues.KeyUsage.BaseKeyUsage.DataEncipherment != nil {
					rPredefinedValuesKeyUsageBaseKeyUsage["dataEncipherment"] = *r.PredefinedValues.KeyUsage.BaseKeyUsage.DataEncipherment
				}
				if r.PredefinedValues.KeyUsage.BaseKeyUsage.DecipherOnly != nil {
					rPredefinedValuesKeyUsageBaseKeyUsage["decipherOnly"] = *r.PredefinedValues.KeyUsage.BaseKeyUsage.DecipherOnly
				}
				if r.PredefinedValues.KeyUsage.BaseKeyUsage.DigitalSignature != nil {
					rPredefinedValuesKeyUsageBaseKeyUsage["digitalSignature"] = *r.PredefinedValues.KeyUsage.BaseKeyUsage.DigitalSignature
				}
				if r.PredefinedValues.KeyUsage.BaseKeyUsage.EncipherOnly != nil {
					rPredefinedValuesKeyUsageBaseKeyUsage["encipherOnly"] = *r.PredefinedValues.KeyUsage.BaseKeyUsage.EncipherOnly
				}
				if r.PredefinedValues.KeyUsage.BaseKeyUsage.KeyAgreement != nil {
					rPredefinedValuesKeyUsageBaseKeyUsage["keyAgreement"] = *r.PredefinedValues.KeyUsage.BaseKeyUsage.KeyAgreement
				}
				if r.PredefinedValues.KeyUsage.BaseKeyUsage.KeyEncipherment != nil {
					rPredefinedValuesKeyUsageBaseKeyUsage["keyEncipherment"] = *r.PredefinedValues.KeyUsage.BaseKeyUsage.KeyEncipherment
				}
				rPredefinedValuesKeyUsage["baseKeyUsage"] = rPredefinedValuesKeyUsageBaseKeyUsage
			}
			if r.PredefinedValues.KeyUsage.ExtendedKeyUsage != nil && r.PredefinedValues.KeyUsage.ExtendedKeyUsage != dclService.EmptyCertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsage {
				rPredefinedValuesKeyUsageExtendedKeyUsage := make(map[string]interface{})
				if r.PredefinedValues.KeyUsage.ExtendedKeyUsage.ClientAuth != nil {
					rPredefinedValuesKeyUsageExtendedKeyUsage["clientAuth"] = *r.PredefinedValues.KeyUsage.ExtendedKeyUsage.ClientAuth
				}
				if r.PredefinedValues.KeyUsage.ExtendedKeyUsage.CodeSigning != nil {
					rPredefinedValuesKeyUsageExtendedKeyUsage["codeSigning"] = *r.PredefinedValues.KeyUsage.ExtendedKeyUsage.CodeSigning
				}
				if r.PredefinedValues.KeyUsage.ExtendedKeyUsage.EmailProtection != nil {
					rPredefinedValuesKeyUsageExtendedKeyUsage["emailProtection"] = *r.PredefinedValues.KeyUsage.ExtendedKeyUsage.EmailProtection
				}
				if r.PredefinedValues.KeyUsage.ExtendedKeyUsage.OcspSigning != nil {
					rPredefinedValuesKeyUsageExtendedKeyUsage["ocspSigning"] = *r.PredefinedValues.KeyUsage.ExtendedKeyUsage.OcspSigning
				}
				if r.PredefinedValues.KeyUsage.ExtendedKeyUsage.ServerAuth != nil {
					rPredefinedValuesKeyUsageExtendedKeyUsage["serverAuth"] = *r.PredefinedValues.KeyUsage.ExtendedKeyUsage.ServerAuth
				}
				if r.PredefinedValues.KeyUsage.ExtendedKeyUsage.TimeStamping != nil {
					rPredefinedValuesKeyUsageExtendedKeyUsage["timeStamping"] = *r.PredefinedValues.KeyUsage.ExtendedKeyUsage.TimeStamping
				}
				rPredefinedValuesKeyUsage["extendedKeyUsage"] = rPredefinedValuesKeyUsageExtendedKeyUsage
			}
			var rPredefinedValuesKeyUsageUnknownExtendedKeyUsages []interface{}
			for _, rPredefinedValuesKeyUsageUnknownExtendedKeyUsagesVal := range r.PredefinedValues.KeyUsage.UnknownExtendedKeyUsages {
				rPredefinedValuesKeyUsageUnknownExtendedKeyUsagesObject := make(map[string]interface{})
				var rPredefinedValuesKeyUsageUnknownExtendedKeyUsagesValObjectIdPath []interface{}
				for _, rPredefinedValuesKeyUsageUnknownExtendedKeyUsagesValObjectIdPathVal := range rPredefinedValuesKeyUsageUnknownExtendedKeyUsagesVal.ObjectIdPath {
					rPredefinedValuesKeyUsageUnknownExtendedKeyUsagesValObjectIdPath = append(rPredefinedValuesKeyUsageUnknownExtendedKeyUsagesValObjectIdPath, rPredefinedValuesKeyUsageUnknownExtendedKeyUsagesValObjectIdPathVal)
				}
				rPredefinedValuesKeyUsageUnknownExtendedKeyUsagesObject["objectIdPath"] = rPredefinedValuesKeyUsageUnknownExtendedKeyUsagesValObjectIdPath
				rPredefinedValuesKeyUsageUnknownExtendedKeyUsages = append(rPredefinedValuesKeyUsageUnknownExtendedKeyUsages, rPredefinedValuesKeyUsageUnknownExtendedKeyUsagesObject)
			}
			rPredefinedValuesKeyUsage["unknownExtendedKeyUsages"] = rPredefinedValuesKeyUsageUnknownExtendedKeyUsages
			rPredefinedValues["keyUsage"] = rPredefinedValuesKeyUsage
		}
		var rPredefinedValuesPolicyIds []interface{}
		for _, rPredefinedValuesPolicyIdsVal := range r.PredefinedValues.PolicyIds {
			rPredefinedValuesPolicyIdsObject := make(map[string]interface{})
			var rPredefinedValuesPolicyIdsValObjectIdPath []interface{}
			for _, rPredefinedValuesPolicyIdsValObjectIdPathVal := range rPredefinedValuesPolicyIdsVal.ObjectIdPath {
				rPredefinedValuesPolicyIdsValObjectIdPath = append(rPredefinedValuesPolicyIdsValObjectIdPath, rPredefinedValuesPolicyIdsValObjectIdPathVal)
			}
			rPredefinedValuesPolicyIdsObject["objectIdPath"] = rPredefinedValuesPolicyIdsValObjectIdPath
			rPredefinedValuesPolicyIds = append(rPredefinedValuesPolicyIds, rPredefinedValuesPolicyIdsObject)
		}
		rPredefinedValues["policyIds"] = rPredefinedValuesPolicyIds
		u.Object["predefinedValues"] = rPredefinedValues
	}
	if r.Project != nil {
		u.Object["project"] = *r.Project
	}
	if r.UpdateTime != nil {
		u.Object["updateTime"] = *r.UpdateTime
	}
	return u
}

func UnstructuredToCertificateTemplate(u *unstructured.Resource) (*dclService.CertificateTemplate, error) {
	r := &dclService.CertificateTemplate{}
	if _, ok := u.Object["createTime"]; ok {
		if s, ok := u.Object["createTime"].(string); ok {
			r.CreateTime = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.CreateTime: expected string")
		}
	}
	if _, ok := u.Object["description"]; ok {
		if s, ok := u.Object["description"].(string); ok {
			r.Description = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Description: expected string")
		}
	}
	if _, ok := u.Object["identityConstraints"]; ok {
		if rIdentityConstraints, ok := u.Object["identityConstraints"].(map[string]interface{}); ok {
			r.IdentityConstraints = &dclService.CertificateTemplateIdentityConstraints{}
			if _, ok := rIdentityConstraints["allowSubjectAltNamesPassthrough"]; ok {
				if b, ok := rIdentityConstraints["allowSubjectAltNamesPassthrough"].(bool); ok {
					r.IdentityConstraints.AllowSubjectAltNamesPassthrough = dcl.Bool(b)
				} else {
					return nil, fmt.Errorf("r.IdentityConstraints.AllowSubjectAltNamesPassthrough: expected bool")
				}
			}
			if _, ok := rIdentityConstraints["allowSubjectPassthrough"]; ok {
				if b, ok := rIdentityConstraints["allowSubjectPassthrough"].(bool); ok {
					r.IdentityConstraints.AllowSubjectPassthrough = dcl.Bool(b)
				} else {
					return nil, fmt.Errorf("r.IdentityConstraints.AllowSubjectPassthrough: expected bool")
				}
			}
			if _, ok := rIdentityConstraints["celExpression"]; ok {
				if rIdentityConstraintsCelExpression, ok := rIdentityConstraints["celExpression"].(map[string]interface{}); ok {
					r.IdentityConstraints.CelExpression = &dclService.CertificateTemplateIdentityConstraintsCelExpression{}
					if _, ok := rIdentityConstraintsCelExpression["description"]; ok {
						if s, ok := rIdentityConstraintsCelExpression["description"].(string); ok {
							r.IdentityConstraints.CelExpression.Description = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.IdentityConstraints.CelExpression.Description: expected string")
						}
					}
					if _, ok := rIdentityConstraintsCelExpression["expression"]; ok {
						if s, ok := rIdentityConstraintsCelExpression["expression"].(string); ok {
							r.IdentityConstraints.CelExpression.Expression = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.IdentityConstraints.CelExpression.Expression: expected string")
						}
					}
					if _, ok := rIdentityConstraintsCelExpression["location"]; ok {
						if s, ok := rIdentityConstraintsCelExpression["location"].(string); ok {
							r.IdentityConstraints.CelExpression.Location = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.IdentityConstraints.CelExpression.Location: expected string")
						}
					}
					if _, ok := rIdentityConstraintsCelExpression["title"]; ok {
						if s, ok := rIdentityConstraintsCelExpression["title"].(string); ok {
							r.IdentityConstraints.CelExpression.Title = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.IdentityConstraints.CelExpression.Title: expected string")
						}
					}
				} else {
					return nil, fmt.Errorf("r.IdentityConstraints.CelExpression: expected map[string]interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.IdentityConstraints: expected map[string]interface{}")
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
	if _, ok := u.Object["passthroughExtensions"]; ok {
		if rPassthroughExtensions, ok := u.Object["passthroughExtensions"].(map[string]interface{}); ok {
			r.PassthroughExtensions = &dclService.CertificateTemplatePassthroughExtensions{}
			if _, ok := rPassthroughExtensions["additionalExtensions"]; ok {
				if s, ok := rPassthroughExtensions["additionalExtensions"].([]interface{}); ok {
					for _, o := range s {
						if objval, ok := o.(map[string]interface{}); ok {
							var rPassthroughExtensionsAdditionalExtensions dclService.CertificateTemplatePassthroughExtensionsAdditionalExtensions
							if _, ok := objval["objectIdPath"]; ok {
								if s, ok := objval["objectIdPath"].([]interface{}); ok {
									for _, ss := range s {
										if intval, ok := ss.(int64); ok {
											rPassthroughExtensionsAdditionalExtensions.ObjectIdPath = append(rPassthroughExtensionsAdditionalExtensions.ObjectIdPath, intval)
										}
									}
								} else {
									return nil, fmt.Errorf("rPassthroughExtensionsAdditionalExtensions.ObjectIdPath: expected []interface{}")
								}
							}
							r.PassthroughExtensions.AdditionalExtensions = append(r.PassthroughExtensions.AdditionalExtensions, rPassthroughExtensionsAdditionalExtensions)
						}
					}
				} else {
					return nil, fmt.Errorf("r.PassthroughExtensions.AdditionalExtensions: expected []interface{}")
				}
			}
			if _, ok := rPassthroughExtensions["knownExtensions"]; ok {
				if s, ok := rPassthroughExtensions["knownExtensions"].([]interface{}); ok {
					for _, ss := range s {
						if strval, ok := ss.(string); ok {
							r.PassthroughExtensions.KnownExtensions = append(r.PassthroughExtensions.KnownExtensions, dclService.CertificateTemplatePassthroughExtensionsKnownExtensionsEnum(strval))
						}
					}
				} else {
					return nil, fmt.Errorf("r.PassthroughExtensions.KnownExtensions: expected []interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.PassthroughExtensions: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["predefinedValues"]; ok {
		if rPredefinedValues, ok := u.Object["predefinedValues"].(map[string]interface{}); ok {
			r.PredefinedValues = &dclService.CertificateTemplatePredefinedValues{}
			if _, ok := rPredefinedValues["additionalExtensions"]; ok {
				if s, ok := rPredefinedValues["additionalExtensions"].([]interface{}); ok {
					for _, o := range s {
						if objval, ok := o.(map[string]interface{}); ok {
							var rPredefinedValuesAdditionalExtensions dclService.CertificateTemplatePredefinedValuesAdditionalExtensions
							if _, ok := objval["critical"]; ok {
								if b, ok := objval["critical"].(bool); ok {
									rPredefinedValuesAdditionalExtensions.Critical = dcl.Bool(b)
								} else {
									return nil, fmt.Errorf("rPredefinedValuesAdditionalExtensions.Critical: expected bool")
								}
							}
							if _, ok := objval["objectId"]; ok {
								if rPredefinedValuesAdditionalExtensionsObjectId, ok := objval["objectId"].(map[string]interface{}); ok {
									rPredefinedValuesAdditionalExtensions.ObjectId = &dclService.CertificateTemplatePredefinedValuesAdditionalExtensionsObjectId{}
									if _, ok := rPredefinedValuesAdditionalExtensionsObjectId["objectIdPath"]; ok {
										if s, ok := rPredefinedValuesAdditionalExtensionsObjectId["objectIdPath"].([]interface{}); ok {
											for _, ss := range s {
												if intval, ok := ss.(int64); ok {
													rPredefinedValuesAdditionalExtensions.ObjectId.ObjectIdPath = append(rPredefinedValuesAdditionalExtensions.ObjectId.ObjectIdPath, intval)
												}
											}
										} else {
											return nil, fmt.Errorf("rPredefinedValuesAdditionalExtensions.ObjectId.ObjectIdPath: expected []interface{}")
										}
									}
								} else {
									return nil, fmt.Errorf("rPredefinedValuesAdditionalExtensions.ObjectId: expected map[string]interface{}")
								}
							}
							if _, ok := objval["value"]; ok {
								if s, ok := objval["value"].(string); ok {
									rPredefinedValuesAdditionalExtensions.Value = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rPredefinedValuesAdditionalExtensions.Value: expected string")
								}
							}
							r.PredefinedValues.AdditionalExtensions = append(r.PredefinedValues.AdditionalExtensions, rPredefinedValuesAdditionalExtensions)
						}
					}
				} else {
					return nil, fmt.Errorf("r.PredefinedValues.AdditionalExtensions: expected []interface{}")
				}
			}
			if _, ok := rPredefinedValues["aiaOcspServers"]; ok {
				if s, ok := rPredefinedValues["aiaOcspServers"].([]interface{}); ok {
					for _, ss := range s {
						if strval, ok := ss.(string); ok {
							r.PredefinedValues.AiaOcspServers = append(r.PredefinedValues.AiaOcspServers, strval)
						}
					}
				} else {
					return nil, fmt.Errorf("r.PredefinedValues.AiaOcspServers: expected []interface{}")
				}
			}
			if _, ok := rPredefinedValues["caOptions"]; ok {
				if rPredefinedValuesCaOptions, ok := rPredefinedValues["caOptions"].(map[string]interface{}); ok {
					r.PredefinedValues.CaOptions = &dclService.CertificateTemplatePredefinedValuesCaOptions{}
					if _, ok := rPredefinedValuesCaOptions["isCa"]; ok {
						if b, ok := rPredefinedValuesCaOptions["isCa"].(bool); ok {
							r.PredefinedValues.CaOptions.IsCa = dcl.Bool(b)
						} else {
							return nil, fmt.Errorf("r.PredefinedValues.CaOptions.IsCa: expected bool")
						}
					}
					if _, ok := rPredefinedValuesCaOptions["maxIssuerPathLength"]; ok {
						if i, ok := rPredefinedValuesCaOptions["maxIssuerPathLength"].(int64); ok {
							r.PredefinedValues.CaOptions.MaxIssuerPathLength = dcl.Int64(i)
						} else {
							return nil, fmt.Errorf("r.PredefinedValues.CaOptions.MaxIssuerPathLength: expected int64")
						}
					}
				} else {
					return nil, fmt.Errorf("r.PredefinedValues.CaOptions: expected map[string]interface{}")
				}
			}
			if _, ok := rPredefinedValues["keyUsage"]; ok {
				if rPredefinedValuesKeyUsage, ok := rPredefinedValues["keyUsage"].(map[string]interface{}); ok {
					r.PredefinedValues.KeyUsage = &dclService.CertificateTemplatePredefinedValuesKeyUsage{}
					if _, ok := rPredefinedValuesKeyUsage["baseKeyUsage"]; ok {
						if rPredefinedValuesKeyUsageBaseKeyUsage, ok := rPredefinedValuesKeyUsage["baseKeyUsage"].(map[string]interface{}); ok {
							r.PredefinedValues.KeyUsage.BaseKeyUsage = &dclService.CertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage{}
							if _, ok := rPredefinedValuesKeyUsageBaseKeyUsage["certSign"]; ok {
								if b, ok := rPredefinedValuesKeyUsageBaseKeyUsage["certSign"].(bool); ok {
									r.PredefinedValues.KeyUsage.BaseKeyUsage.CertSign = dcl.Bool(b)
								} else {
									return nil, fmt.Errorf("r.PredefinedValues.KeyUsage.BaseKeyUsage.CertSign: expected bool")
								}
							}
							if _, ok := rPredefinedValuesKeyUsageBaseKeyUsage["contentCommitment"]; ok {
								if b, ok := rPredefinedValuesKeyUsageBaseKeyUsage["contentCommitment"].(bool); ok {
									r.PredefinedValues.KeyUsage.BaseKeyUsage.ContentCommitment = dcl.Bool(b)
								} else {
									return nil, fmt.Errorf("r.PredefinedValues.KeyUsage.BaseKeyUsage.ContentCommitment: expected bool")
								}
							}
							if _, ok := rPredefinedValuesKeyUsageBaseKeyUsage["crlSign"]; ok {
								if b, ok := rPredefinedValuesKeyUsageBaseKeyUsage["crlSign"].(bool); ok {
									r.PredefinedValues.KeyUsage.BaseKeyUsage.CrlSign = dcl.Bool(b)
								} else {
									return nil, fmt.Errorf("r.PredefinedValues.KeyUsage.BaseKeyUsage.CrlSign: expected bool")
								}
							}
							if _, ok := rPredefinedValuesKeyUsageBaseKeyUsage["dataEncipherment"]; ok {
								if b, ok := rPredefinedValuesKeyUsageBaseKeyUsage["dataEncipherment"].(bool); ok {
									r.PredefinedValues.KeyUsage.BaseKeyUsage.DataEncipherment = dcl.Bool(b)
								} else {
									return nil, fmt.Errorf("r.PredefinedValues.KeyUsage.BaseKeyUsage.DataEncipherment: expected bool")
								}
							}
							if _, ok := rPredefinedValuesKeyUsageBaseKeyUsage["decipherOnly"]; ok {
								if b, ok := rPredefinedValuesKeyUsageBaseKeyUsage["decipherOnly"].(bool); ok {
									r.PredefinedValues.KeyUsage.BaseKeyUsage.DecipherOnly = dcl.Bool(b)
								} else {
									return nil, fmt.Errorf("r.PredefinedValues.KeyUsage.BaseKeyUsage.DecipherOnly: expected bool")
								}
							}
							if _, ok := rPredefinedValuesKeyUsageBaseKeyUsage["digitalSignature"]; ok {
								if b, ok := rPredefinedValuesKeyUsageBaseKeyUsage["digitalSignature"].(bool); ok {
									r.PredefinedValues.KeyUsage.BaseKeyUsage.DigitalSignature = dcl.Bool(b)
								} else {
									return nil, fmt.Errorf("r.PredefinedValues.KeyUsage.BaseKeyUsage.DigitalSignature: expected bool")
								}
							}
							if _, ok := rPredefinedValuesKeyUsageBaseKeyUsage["encipherOnly"]; ok {
								if b, ok := rPredefinedValuesKeyUsageBaseKeyUsage["encipherOnly"].(bool); ok {
									r.PredefinedValues.KeyUsage.BaseKeyUsage.EncipherOnly = dcl.Bool(b)
								} else {
									return nil, fmt.Errorf("r.PredefinedValues.KeyUsage.BaseKeyUsage.EncipherOnly: expected bool")
								}
							}
							if _, ok := rPredefinedValuesKeyUsageBaseKeyUsage["keyAgreement"]; ok {
								if b, ok := rPredefinedValuesKeyUsageBaseKeyUsage["keyAgreement"].(bool); ok {
									r.PredefinedValues.KeyUsage.BaseKeyUsage.KeyAgreement = dcl.Bool(b)
								} else {
									return nil, fmt.Errorf("r.PredefinedValues.KeyUsage.BaseKeyUsage.KeyAgreement: expected bool")
								}
							}
							if _, ok := rPredefinedValuesKeyUsageBaseKeyUsage["keyEncipherment"]; ok {
								if b, ok := rPredefinedValuesKeyUsageBaseKeyUsage["keyEncipherment"].(bool); ok {
									r.PredefinedValues.KeyUsage.BaseKeyUsage.KeyEncipherment = dcl.Bool(b)
								} else {
									return nil, fmt.Errorf("r.PredefinedValues.KeyUsage.BaseKeyUsage.KeyEncipherment: expected bool")
								}
							}
						} else {
							return nil, fmt.Errorf("r.PredefinedValues.KeyUsage.BaseKeyUsage: expected map[string]interface{}")
						}
					}
					if _, ok := rPredefinedValuesKeyUsage["extendedKeyUsage"]; ok {
						if rPredefinedValuesKeyUsageExtendedKeyUsage, ok := rPredefinedValuesKeyUsage["extendedKeyUsage"].(map[string]interface{}); ok {
							r.PredefinedValues.KeyUsage.ExtendedKeyUsage = &dclService.CertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsage{}
							if _, ok := rPredefinedValuesKeyUsageExtendedKeyUsage["clientAuth"]; ok {
								if b, ok := rPredefinedValuesKeyUsageExtendedKeyUsage["clientAuth"].(bool); ok {
									r.PredefinedValues.KeyUsage.ExtendedKeyUsage.ClientAuth = dcl.Bool(b)
								} else {
									return nil, fmt.Errorf("r.PredefinedValues.KeyUsage.ExtendedKeyUsage.ClientAuth: expected bool")
								}
							}
							if _, ok := rPredefinedValuesKeyUsageExtendedKeyUsage["codeSigning"]; ok {
								if b, ok := rPredefinedValuesKeyUsageExtendedKeyUsage["codeSigning"].(bool); ok {
									r.PredefinedValues.KeyUsage.ExtendedKeyUsage.CodeSigning = dcl.Bool(b)
								} else {
									return nil, fmt.Errorf("r.PredefinedValues.KeyUsage.ExtendedKeyUsage.CodeSigning: expected bool")
								}
							}
							if _, ok := rPredefinedValuesKeyUsageExtendedKeyUsage["emailProtection"]; ok {
								if b, ok := rPredefinedValuesKeyUsageExtendedKeyUsage["emailProtection"].(bool); ok {
									r.PredefinedValues.KeyUsage.ExtendedKeyUsage.EmailProtection = dcl.Bool(b)
								} else {
									return nil, fmt.Errorf("r.PredefinedValues.KeyUsage.ExtendedKeyUsage.EmailProtection: expected bool")
								}
							}
							if _, ok := rPredefinedValuesKeyUsageExtendedKeyUsage["ocspSigning"]; ok {
								if b, ok := rPredefinedValuesKeyUsageExtendedKeyUsage["ocspSigning"].(bool); ok {
									r.PredefinedValues.KeyUsage.ExtendedKeyUsage.OcspSigning = dcl.Bool(b)
								} else {
									return nil, fmt.Errorf("r.PredefinedValues.KeyUsage.ExtendedKeyUsage.OcspSigning: expected bool")
								}
							}
							if _, ok := rPredefinedValuesKeyUsageExtendedKeyUsage["serverAuth"]; ok {
								if b, ok := rPredefinedValuesKeyUsageExtendedKeyUsage["serverAuth"].(bool); ok {
									r.PredefinedValues.KeyUsage.ExtendedKeyUsage.ServerAuth = dcl.Bool(b)
								} else {
									return nil, fmt.Errorf("r.PredefinedValues.KeyUsage.ExtendedKeyUsage.ServerAuth: expected bool")
								}
							}
							if _, ok := rPredefinedValuesKeyUsageExtendedKeyUsage["timeStamping"]; ok {
								if b, ok := rPredefinedValuesKeyUsageExtendedKeyUsage["timeStamping"].(bool); ok {
									r.PredefinedValues.KeyUsage.ExtendedKeyUsage.TimeStamping = dcl.Bool(b)
								} else {
									return nil, fmt.Errorf("r.PredefinedValues.KeyUsage.ExtendedKeyUsage.TimeStamping: expected bool")
								}
							}
						} else {
							return nil, fmt.Errorf("r.PredefinedValues.KeyUsage.ExtendedKeyUsage: expected map[string]interface{}")
						}
					}
					if _, ok := rPredefinedValuesKeyUsage["unknownExtendedKeyUsages"]; ok {
						if s, ok := rPredefinedValuesKeyUsage["unknownExtendedKeyUsages"].([]interface{}); ok {
							for _, o := range s {
								if objval, ok := o.(map[string]interface{}); ok {
									var rPredefinedValuesKeyUsageUnknownExtendedKeyUsages dclService.CertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsages
									if _, ok := objval["objectIdPath"]; ok {
										if s, ok := objval["objectIdPath"].([]interface{}); ok {
											for _, ss := range s {
												if intval, ok := ss.(int64); ok {
													rPredefinedValuesKeyUsageUnknownExtendedKeyUsages.ObjectIdPath = append(rPredefinedValuesKeyUsageUnknownExtendedKeyUsages.ObjectIdPath, intval)
												}
											}
										} else {
											return nil, fmt.Errorf("rPredefinedValuesKeyUsageUnknownExtendedKeyUsages.ObjectIdPath: expected []interface{}")
										}
									}
									r.PredefinedValues.KeyUsage.UnknownExtendedKeyUsages = append(r.PredefinedValues.KeyUsage.UnknownExtendedKeyUsages, rPredefinedValuesKeyUsageUnknownExtendedKeyUsages)
								}
							}
						} else {
							return nil, fmt.Errorf("r.PredefinedValues.KeyUsage.UnknownExtendedKeyUsages: expected []interface{}")
						}
					}
				} else {
					return nil, fmt.Errorf("r.PredefinedValues.KeyUsage: expected map[string]interface{}")
				}
			}
			if _, ok := rPredefinedValues["policyIds"]; ok {
				if s, ok := rPredefinedValues["policyIds"].([]interface{}); ok {
					for _, o := range s {
						if objval, ok := o.(map[string]interface{}); ok {
							var rPredefinedValuesPolicyIds dclService.CertificateTemplatePredefinedValuesPolicyIds
							if _, ok := objval["objectIdPath"]; ok {
								if s, ok := objval["objectIdPath"].([]interface{}); ok {
									for _, ss := range s {
										if intval, ok := ss.(int64); ok {
											rPredefinedValuesPolicyIds.ObjectIdPath = append(rPredefinedValuesPolicyIds.ObjectIdPath, intval)
										}
									}
								} else {
									return nil, fmt.Errorf("rPredefinedValuesPolicyIds.ObjectIdPath: expected []interface{}")
								}
							}
							r.PredefinedValues.PolicyIds = append(r.PredefinedValues.PolicyIds, rPredefinedValuesPolicyIds)
						}
					}
				} else {
					return nil, fmt.Errorf("r.PredefinedValues.PolicyIds: expected []interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.PredefinedValues: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["project"]; ok {
		if s, ok := u.Object["project"].(string); ok {
			r.Project = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Project: expected string")
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

func GetCertificateTemplate(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToCertificateTemplate(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetCertificateTemplate(ctx, r)
	if err != nil {
		return nil, err
	}
	return CertificateTemplateToUnstructured(r), nil
}

func ListCertificateTemplate(ctx context.Context, config *dcl.Config, project string, location string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListCertificateTemplate(ctx, project, location)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, CertificateTemplateToUnstructured(r))
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

func ApplyCertificateTemplate(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToCertificateTemplate(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToCertificateTemplate(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyCertificateTemplate(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return CertificateTemplateToUnstructured(r), nil
}

func CertificateTemplateHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToCertificateTemplate(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToCertificateTemplate(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyCertificateTemplate(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteCertificateTemplate(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToCertificateTemplate(u)
	if err != nil {
		return err
	}
	return c.DeleteCertificateTemplate(ctx, r)
}

func CertificateTemplateID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToCertificateTemplate(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *CertificateTemplate) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"privateca",
		"CertificateTemplate",
		"alpha",
	}
}

func (r *CertificateTemplate) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *CertificateTemplate) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *CertificateTemplate) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *CertificateTemplate) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *CertificateTemplate) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *CertificateTemplate) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *CertificateTemplate) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetCertificateTemplate(ctx, config, resource)
}

func (r *CertificateTemplate) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyCertificateTemplate(ctx, config, resource, opts...)
}

func (r *CertificateTemplate) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return CertificateTemplateHasDiff(ctx, config, resource, opts...)
}

func (r *CertificateTemplate) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteCertificateTemplate(ctx, config, resource)
}

func (r *CertificateTemplate) ID(resource *unstructured.Resource) (string, error) {
	return CertificateTemplateID(resource)
}

func init() {
	unstructured.Register(&CertificateTemplate{})
}
