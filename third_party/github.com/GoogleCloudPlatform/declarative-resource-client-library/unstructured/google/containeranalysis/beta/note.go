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
package containeranalysis

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/containeranalysis/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
)

type Note struct{}

func NoteToUnstructured(r *dclService.Note) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "containeranalysis",
			Version: "beta",
			Type:    "Note",
		},
		Object: make(map[string]interface{}),
	}
	if r.Attestation != nil && r.Attestation != dclService.EmptyNoteAttestation {
		rAttestation := make(map[string]interface{})
		if r.Attestation.Hint != nil && r.Attestation.Hint != dclService.EmptyNoteAttestationHint {
			rAttestationHint := make(map[string]interface{})
			if r.Attestation.Hint.HumanReadableName != nil {
				rAttestationHint["humanReadableName"] = *r.Attestation.Hint.HumanReadableName
			}
			rAttestation["hint"] = rAttestationHint
		}
		u.Object["attestation"] = rAttestation
	}
	if r.Build != nil && r.Build != dclService.EmptyNoteBuild {
		rBuild := make(map[string]interface{})
		if r.Build.BuilderVersion != nil {
			rBuild["builderVersion"] = *r.Build.BuilderVersion
		}
		if r.Build.Signature != nil && r.Build.Signature != dclService.EmptyNoteBuildSignature {
			rBuildSignature := make(map[string]interface{})
			if r.Build.Signature.KeyId != nil {
				rBuildSignature["keyId"] = *r.Build.Signature.KeyId
			}
			if r.Build.Signature.KeyType != nil {
				rBuildSignature["keyType"] = string(*r.Build.Signature.KeyType)
			}
			if r.Build.Signature.PublicKey != nil {
				rBuildSignature["publicKey"] = *r.Build.Signature.PublicKey
			}
			if r.Build.Signature.Signature != nil {
				rBuildSignature["signature"] = *r.Build.Signature.Signature
			}
			rBuild["signature"] = rBuildSignature
		}
		u.Object["build"] = rBuild
	}
	if r.CreateTime != nil {
		u.Object["createTime"] = *r.CreateTime
	}
	if r.Deployment != nil && r.Deployment != dclService.EmptyNoteDeployment {
		rDeployment := make(map[string]interface{})
		var rDeploymentResourceUri []interface{}
		for _, rDeploymentResourceUriVal := range r.Deployment.ResourceUri {
			rDeploymentResourceUri = append(rDeploymentResourceUri, rDeploymentResourceUriVal)
		}
		rDeployment["resourceUri"] = rDeploymentResourceUri
		u.Object["deployment"] = rDeployment
	}
	if r.Discovery != nil && r.Discovery != dclService.EmptyNoteDiscovery {
		rDiscovery := make(map[string]interface{})
		if r.Discovery.AnalysisKind != nil {
			rDiscovery["analysisKind"] = string(*r.Discovery.AnalysisKind)
		}
		u.Object["discovery"] = rDiscovery
	}
	if r.ExpirationTime != nil {
		u.Object["expirationTime"] = *r.ExpirationTime
	}
	if r.Image != nil && r.Image != dclService.EmptyNoteImage {
		rImage := make(map[string]interface{})
		if r.Image.Fingerprint != nil && r.Image.Fingerprint != dclService.EmptyNoteImageFingerprint {
			rImageFingerprint := make(map[string]interface{})
			if r.Image.Fingerprint.V1Name != nil {
				rImageFingerprint["v1Name"] = *r.Image.Fingerprint.V1Name
			}
			var rImageFingerprintV2Blob []interface{}
			for _, rImageFingerprintV2BlobVal := range r.Image.Fingerprint.V2Blob {
				rImageFingerprintV2Blob = append(rImageFingerprintV2Blob, rImageFingerprintV2BlobVal)
			}
			rImageFingerprint["v2Blob"] = rImageFingerprintV2Blob
			if r.Image.Fingerprint.V2Name != nil {
				rImageFingerprint["v2Name"] = *r.Image.Fingerprint.V2Name
			}
			rImage["fingerprint"] = rImageFingerprint
		}
		if r.Image.ResourceUrl != nil {
			rImage["resourceUrl"] = *r.Image.ResourceUrl
		}
		u.Object["image"] = rImage
	}
	if r.LongDescription != nil {
		u.Object["longDescription"] = *r.LongDescription
	}
	if r.Name != nil {
		u.Object["name"] = *r.Name
	}
	if r.Package != nil && r.Package != dclService.EmptyNotePackage {
		rPackage := make(map[string]interface{})
		var rPackageDistribution []interface{}
		for _, rPackageDistributionVal := range r.Package.Distribution {
			rPackageDistributionObject := make(map[string]interface{})
			if rPackageDistributionVal.Architecture != nil {
				rPackageDistributionObject["architecture"] = string(*rPackageDistributionVal.Architecture)
			}
			if rPackageDistributionVal.CpeUri != nil {
				rPackageDistributionObject["cpeUri"] = *rPackageDistributionVal.CpeUri
			}
			if rPackageDistributionVal.Description != nil {
				rPackageDistributionObject["description"] = *rPackageDistributionVal.Description
			}
			if rPackageDistributionVal.LatestVersion != nil && rPackageDistributionVal.LatestVersion != dclService.EmptyNotePackageDistributionLatestVersion {
				rPackageDistributionValLatestVersion := make(map[string]interface{})
				if rPackageDistributionVal.LatestVersion.Epoch != nil {
					rPackageDistributionValLatestVersion["epoch"] = *rPackageDistributionVal.LatestVersion.Epoch
				}
				if rPackageDistributionVal.LatestVersion.FullName != nil {
					rPackageDistributionValLatestVersion["fullName"] = *rPackageDistributionVal.LatestVersion.FullName
				}
				if rPackageDistributionVal.LatestVersion.Kind != nil {
					rPackageDistributionValLatestVersion["kind"] = string(*rPackageDistributionVal.LatestVersion.Kind)
				}
				if rPackageDistributionVal.LatestVersion.Name != nil {
					rPackageDistributionValLatestVersion["name"] = *rPackageDistributionVal.LatestVersion.Name
				}
				if rPackageDistributionVal.LatestVersion.Revision != nil {
					rPackageDistributionValLatestVersion["revision"] = *rPackageDistributionVal.LatestVersion.Revision
				}
				rPackageDistributionObject["latestVersion"] = rPackageDistributionValLatestVersion
			}
			if rPackageDistributionVal.Maintainer != nil {
				rPackageDistributionObject["maintainer"] = *rPackageDistributionVal.Maintainer
			}
			if rPackageDistributionVal.Url != nil {
				rPackageDistributionObject["url"] = *rPackageDistributionVal.Url
			}
			rPackageDistribution = append(rPackageDistribution, rPackageDistributionObject)
		}
		rPackage["distribution"] = rPackageDistribution
		if r.Package.Name != nil {
			rPackage["name"] = *r.Package.Name
		}
		u.Object["package"] = rPackage
	}
	if r.Project != nil {
		u.Object["project"] = *r.Project
	}
	var rRelatedNoteNames []interface{}
	for _, rRelatedNoteNamesVal := range r.RelatedNoteNames {
		rRelatedNoteNames = append(rRelatedNoteNames, rRelatedNoteNamesVal)
	}
	u.Object["relatedNoteNames"] = rRelatedNoteNames
	var rRelatedUrl []interface{}
	for _, rRelatedUrlVal := range r.RelatedUrl {
		rRelatedUrlObject := make(map[string]interface{})
		if rRelatedUrlVal.Label != nil {
			rRelatedUrlObject["label"] = *rRelatedUrlVal.Label
		}
		if rRelatedUrlVal.Url != nil {
			rRelatedUrlObject["url"] = *rRelatedUrlVal.Url
		}
		rRelatedUrl = append(rRelatedUrl, rRelatedUrlObject)
	}
	u.Object["relatedUrl"] = rRelatedUrl
	if r.ShortDescription != nil {
		u.Object["shortDescription"] = *r.ShortDescription
	}
	if r.UpdateTime != nil {
		u.Object["updateTime"] = *r.UpdateTime
	}
	if r.Vulnerability != nil && r.Vulnerability != dclService.EmptyNoteVulnerability {
		rVulnerability := make(map[string]interface{})
		if r.Vulnerability.CvssScore != nil {
			rVulnerability["cvssScore"] = *r.Vulnerability.CvssScore
		}
		if r.Vulnerability.CvssV3 != nil && r.Vulnerability.CvssV3 != dclService.EmptyNoteVulnerabilityCvssV3 {
			rVulnerabilityCvssV3 := make(map[string]interface{})
			if r.Vulnerability.CvssV3.AttackComplexity != nil {
				rVulnerabilityCvssV3["attackComplexity"] = string(*r.Vulnerability.CvssV3.AttackComplexity)
			}
			if r.Vulnerability.CvssV3.AttackVector != nil {
				rVulnerabilityCvssV3["attackVector"] = string(*r.Vulnerability.CvssV3.AttackVector)
			}
			if r.Vulnerability.CvssV3.AvailabilityImpact != nil {
				rVulnerabilityCvssV3["availabilityImpact"] = string(*r.Vulnerability.CvssV3.AvailabilityImpact)
			}
			if r.Vulnerability.CvssV3.BaseScore != nil {
				rVulnerabilityCvssV3["baseScore"] = *r.Vulnerability.CvssV3.BaseScore
			}
			if r.Vulnerability.CvssV3.ConfidentialityImpact != nil {
				rVulnerabilityCvssV3["confidentialityImpact"] = string(*r.Vulnerability.CvssV3.ConfidentialityImpact)
			}
			if r.Vulnerability.CvssV3.ExploitabilityScore != nil {
				rVulnerabilityCvssV3["exploitabilityScore"] = *r.Vulnerability.CvssV3.ExploitabilityScore
			}
			if r.Vulnerability.CvssV3.ImpactScore != nil {
				rVulnerabilityCvssV3["impactScore"] = *r.Vulnerability.CvssV3.ImpactScore
			}
			if r.Vulnerability.CvssV3.IntegrityImpact != nil {
				rVulnerabilityCvssV3["integrityImpact"] = string(*r.Vulnerability.CvssV3.IntegrityImpact)
			}
			if r.Vulnerability.CvssV3.PrivilegesRequired != nil {
				rVulnerabilityCvssV3["privilegesRequired"] = string(*r.Vulnerability.CvssV3.PrivilegesRequired)
			}
			if r.Vulnerability.CvssV3.Scope != nil {
				rVulnerabilityCvssV3["scope"] = string(*r.Vulnerability.CvssV3.Scope)
			}
			if r.Vulnerability.CvssV3.UserInteraction != nil {
				rVulnerabilityCvssV3["userInteraction"] = string(*r.Vulnerability.CvssV3.UserInteraction)
			}
			rVulnerability["cvssV3"] = rVulnerabilityCvssV3
		}
		var rVulnerabilityDetails []interface{}
		for _, rVulnerabilityDetailsVal := range r.Vulnerability.Details {
			rVulnerabilityDetailsObject := make(map[string]interface{})
			if rVulnerabilityDetailsVal.AffectedCpeUri != nil {
				rVulnerabilityDetailsObject["affectedCpeUri"] = *rVulnerabilityDetailsVal.AffectedCpeUri
			}
			if rVulnerabilityDetailsVal.AffectedPackage != nil {
				rVulnerabilityDetailsObject["affectedPackage"] = *rVulnerabilityDetailsVal.AffectedPackage
			}
			if rVulnerabilityDetailsVal.AffectedVersionEnd != nil && rVulnerabilityDetailsVal.AffectedVersionEnd != dclService.EmptyNoteVulnerabilityDetailsAffectedVersionEnd {
				rVulnerabilityDetailsValAffectedVersionEnd := make(map[string]interface{})
				if rVulnerabilityDetailsVal.AffectedVersionEnd.Epoch != nil {
					rVulnerabilityDetailsValAffectedVersionEnd["epoch"] = *rVulnerabilityDetailsVal.AffectedVersionEnd.Epoch
				}
				if rVulnerabilityDetailsVal.AffectedVersionEnd.FullName != nil {
					rVulnerabilityDetailsValAffectedVersionEnd["fullName"] = *rVulnerabilityDetailsVal.AffectedVersionEnd.FullName
				}
				if rVulnerabilityDetailsVal.AffectedVersionEnd.Kind != nil {
					rVulnerabilityDetailsValAffectedVersionEnd["kind"] = string(*rVulnerabilityDetailsVal.AffectedVersionEnd.Kind)
				}
				if rVulnerabilityDetailsVal.AffectedVersionEnd.Name != nil {
					rVulnerabilityDetailsValAffectedVersionEnd["name"] = *rVulnerabilityDetailsVal.AffectedVersionEnd.Name
				}
				if rVulnerabilityDetailsVal.AffectedVersionEnd.Revision != nil {
					rVulnerabilityDetailsValAffectedVersionEnd["revision"] = *rVulnerabilityDetailsVal.AffectedVersionEnd.Revision
				}
				rVulnerabilityDetailsObject["affectedVersionEnd"] = rVulnerabilityDetailsValAffectedVersionEnd
			}
			if rVulnerabilityDetailsVal.AffectedVersionStart != nil && rVulnerabilityDetailsVal.AffectedVersionStart != dclService.EmptyNoteVulnerabilityDetailsAffectedVersionStart {
				rVulnerabilityDetailsValAffectedVersionStart := make(map[string]interface{})
				if rVulnerabilityDetailsVal.AffectedVersionStart.Epoch != nil {
					rVulnerabilityDetailsValAffectedVersionStart["epoch"] = *rVulnerabilityDetailsVal.AffectedVersionStart.Epoch
				}
				if rVulnerabilityDetailsVal.AffectedVersionStart.FullName != nil {
					rVulnerabilityDetailsValAffectedVersionStart["fullName"] = *rVulnerabilityDetailsVal.AffectedVersionStart.FullName
				}
				if rVulnerabilityDetailsVal.AffectedVersionStart.Kind != nil {
					rVulnerabilityDetailsValAffectedVersionStart["kind"] = string(*rVulnerabilityDetailsVal.AffectedVersionStart.Kind)
				}
				if rVulnerabilityDetailsVal.AffectedVersionStart.Name != nil {
					rVulnerabilityDetailsValAffectedVersionStart["name"] = *rVulnerabilityDetailsVal.AffectedVersionStart.Name
				}
				if rVulnerabilityDetailsVal.AffectedVersionStart.Revision != nil {
					rVulnerabilityDetailsValAffectedVersionStart["revision"] = *rVulnerabilityDetailsVal.AffectedVersionStart.Revision
				}
				rVulnerabilityDetailsObject["affectedVersionStart"] = rVulnerabilityDetailsValAffectedVersionStart
			}
			if rVulnerabilityDetailsVal.Description != nil {
				rVulnerabilityDetailsObject["description"] = *rVulnerabilityDetailsVal.Description
			}
			if rVulnerabilityDetailsVal.FixedCpeUri != nil {
				rVulnerabilityDetailsObject["fixedCpeUri"] = *rVulnerabilityDetailsVal.FixedCpeUri
			}
			if rVulnerabilityDetailsVal.FixedPackage != nil {
				rVulnerabilityDetailsObject["fixedPackage"] = *rVulnerabilityDetailsVal.FixedPackage
			}
			if rVulnerabilityDetailsVal.FixedVersion != nil && rVulnerabilityDetailsVal.FixedVersion != dclService.EmptyNoteVulnerabilityDetailsFixedVersion {
				rVulnerabilityDetailsValFixedVersion := make(map[string]interface{})
				if rVulnerabilityDetailsVal.FixedVersion.Epoch != nil {
					rVulnerabilityDetailsValFixedVersion["epoch"] = *rVulnerabilityDetailsVal.FixedVersion.Epoch
				}
				if rVulnerabilityDetailsVal.FixedVersion.FullName != nil {
					rVulnerabilityDetailsValFixedVersion["fullName"] = *rVulnerabilityDetailsVal.FixedVersion.FullName
				}
				if rVulnerabilityDetailsVal.FixedVersion.Kind != nil {
					rVulnerabilityDetailsValFixedVersion["kind"] = string(*rVulnerabilityDetailsVal.FixedVersion.Kind)
				}
				if rVulnerabilityDetailsVal.FixedVersion.Name != nil {
					rVulnerabilityDetailsValFixedVersion["name"] = *rVulnerabilityDetailsVal.FixedVersion.Name
				}
				if rVulnerabilityDetailsVal.FixedVersion.Revision != nil {
					rVulnerabilityDetailsValFixedVersion["revision"] = *rVulnerabilityDetailsVal.FixedVersion.Revision
				}
				rVulnerabilityDetailsObject["fixedVersion"] = rVulnerabilityDetailsValFixedVersion
			}
			if rVulnerabilityDetailsVal.IsObsolete != nil {
				rVulnerabilityDetailsObject["isObsolete"] = *rVulnerabilityDetailsVal.IsObsolete
			}
			if rVulnerabilityDetailsVal.PackageType != nil {
				rVulnerabilityDetailsObject["packageType"] = *rVulnerabilityDetailsVal.PackageType
			}
			if rVulnerabilityDetailsVal.SeverityName != nil {
				rVulnerabilityDetailsObject["severityName"] = *rVulnerabilityDetailsVal.SeverityName
			}
			if rVulnerabilityDetailsVal.SourceUpdateTime != nil {
				rVulnerabilityDetailsObject["sourceUpdateTime"] = *rVulnerabilityDetailsVal.SourceUpdateTime
			}
			rVulnerabilityDetails = append(rVulnerabilityDetails, rVulnerabilityDetailsObject)
		}
		rVulnerability["details"] = rVulnerabilityDetails
		if r.Vulnerability.Severity != nil {
			rVulnerability["severity"] = string(*r.Vulnerability.Severity)
		}
		if r.Vulnerability.SourceUpdateTime != nil {
			rVulnerability["sourceUpdateTime"] = *r.Vulnerability.SourceUpdateTime
		}
		var rVulnerabilityWindowsDetails []interface{}
		for _, rVulnerabilityWindowsDetailsVal := range r.Vulnerability.WindowsDetails {
			rVulnerabilityWindowsDetailsObject := make(map[string]interface{})
			if rVulnerabilityWindowsDetailsVal.CpeUri != nil {
				rVulnerabilityWindowsDetailsObject["cpeUri"] = *rVulnerabilityWindowsDetailsVal.CpeUri
			}
			if rVulnerabilityWindowsDetailsVal.Description != nil {
				rVulnerabilityWindowsDetailsObject["description"] = *rVulnerabilityWindowsDetailsVal.Description
			}
			var rVulnerabilityWindowsDetailsValFixingKbs []interface{}
			for _, rVulnerabilityWindowsDetailsValFixingKbsVal := range rVulnerabilityWindowsDetailsVal.FixingKbs {
				rVulnerabilityWindowsDetailsValFixingKbsObject := make(map[string]interface{})
				if rVulnerabilityWindowsDetailsValFixingKbsVal.Name != nil {
					rVulnerabilityWindowsDetailsValFixingKbsObject["name"] = *rVulnerabilityWindowsDetailsValFixingKbsVal.Name
				}
				if rVulnerabilityWindowsDetailsValFixingKbsVal.Url != nil {
					rVulnerabilityWindowsDetailsValFixingKbsObject["url"] = *rVulnerabilityWindowsDetailsValFixingKbsVal.Url
				}
				rVulnerabilityWindowsDetailsValFixingKbs = append(rVulnerabilityWindowsDetailsValFixingKbs, rVulnerabilityWindowsDetailsValFixingKbsObject)
			}
			rVulnerabilityWindowsDetailsObject["fixingKbs"] = rVulnerabilityWindowsDetailsValFixingKbs
			if rVulnerabilityWindowsDetailsVal.Name != nil {
				rVulnerabilityWindowsDetailsObject["name"] = *rVulnerabilityWindowsDetailsVal.Name
			}
			rVulnerabilityWindowsDetails = append(rVulnerabilityWindowsDetails, rVulnerabilityWindowsDetailsObject)
		}
		rVulnerability["windowsDetails"] = rVulnerabilityWindowsDetails
		u.Object["vulnerability"] = rVulnerability
	}
	return u
}

func UnstructuredToNote(u *unstructured.Resource) (*dclService.Note, error) {
	r := &dclService.Note{}
	if _, ok := u.Object["attestation"]; ok {
		if rAttestation, ok := u.Object["attestation"].(map[string]interface{}); ok {
			r.Attestation = &dclService.NoteAttestation{}
			if _, ok := rAttestation["hint"]; ok {
				if rAttestationHint, ok := rAttestation["hint"].(map[string]interface{}); ok {
					r.Attestation.Hint = &dclService.NoteAttestationHint{}
					if _, ok := rAttestationHint["humanReadableName"]; ok {
						if s, ok := rAttestationHint["humanReadableName"].(string); ok {
							r.Attestation.Hint.HumanReadableName = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.Attestation.Hint.HumanReadableName: expected string")
						}
					}
				} else {
					return nil, fmt.Errorf("r.Attestation.Hint: expected map[string]interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.Attestation: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["build"]; ok {
		if rBuild, ok := u.Object["build"].(map[string]interface{}); ok {
			r.Build = &dclService.NoteBuild{}
			if _, ok := rBuild["builderVersion"]; ok {
				if s, ok := rBuild["builderVersion"].(string); ok {
					r.Build.BuilderVersion = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.Build.BuilderVersion: expected string")
				}
			}
			if _, ok := rBuild["signature"]; ok {
				if rBuildSignature, ok := rBuild["signature"].(map[string]interface{}); ok {
					r.Build.Signature = &dclService.NoteBuildSignature{}
					if _, ok := rBuildSignature["keyId"]; ok {
						if s, ok := rBuildSignature["keyId"].(string); ok {
							r.Build.Signature.KeyId = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.Build.Signature.KeyId: expected string")
						}
					}
					if _, ok := rBuildSignature["keyType"]; ok {
						if s, ok := rBuildSignature["keyType"].(string); ok {
							r.Build.Signature.KeyType = dclService.NoteBuildSignatureKeyTypeEnumRef(s)
						} else {
							return nil, fmt.Errorf("r.Build.Signature.KeyType: expected string")
						}
					}
					if _, ok := rBuildSignature["publicKey"]; ok {
						if s, ok := rBuildSignature["publicKey"].(string); ok {
							r.Build.Signature.PublicKey = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.Build.Signature.PublicKey: expected string")
						}
					}
					if _, ok := rBuildSignature["signature"]; ok {
						if s, ok := rBuildSignature["signature"].(string); ok {
							r.Build.Signature.Signature = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.Build.Signature.Signature: expected string")
						}
					}
				} else {
					return nil, fmt.Errorf("r.Build.Signature: expected map[string]interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.Build: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["createTime"]; ok {
		if s, ok := u.Object["createTime"].(string); ok {
			r.CreateTime = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.CreateTime: expected string")
		}
	}
	if _, ok := u.Object["deployment"]; ok {
		if rDeployment, ok := u.Object["deployment"].(map[string]interface{}); ok {
			r.Deployment = &dclService.NoteDeployment{}
			if _, ok := rDeployment["resourceUri"]; ok {
				if s, ok := rDeployment["resourceUri"].([]interface{}); ok {
					for _, ss := range s {
						if strval, ok := ss.(string); ok {
							r.Deployment.ResourceUri = append(r.Deployment.ResourceUri, strval)
						}
					}
				} else {
					return nil, fmt.Errorf("r.Deployment.ResourceUri: expected []interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.Deployment: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["discovery"]; ok {
		if rDiscovery, ok := u.Object["discovery"].(map[string]interface{}); ok {
			r.Discovery = &dclService.NoteDiscovery{}
			if _, ok := rDiscovery["analysisKind"]; ok {
				if s, ok := rDiscovery["analysisKind"].(string); ok {
					r.Discovery.AnalysisKind = dclService.NoteDiscoveryAnalysisKindEnumRef(s)
				} else {
					return nil, fmt.Errorf("r.Discovery.AnalysisKind: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.Discovery: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["expirationTime"]; ok {
		if s, ok := u.Object["expirationTime"].(string); ok {
			r.ExpirationTime = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.ExpirationTime: expected string")
		}
	}
	if _, ok := u.Object["image"]; ok {
		if rImage, ok := u.Object["image"].(map[string]interface{}); ok {
			r.Image = &dclService.NoteImage{}
			if _, ok := rImage["fingerprint"]; ok {
				if rImageFingerprint, ok := rImage["fingerprint"].(map[string]interface{}); ok {
					r.Image.Fingerprint = &dclService.NoteImageFingerprint{}
					if _, ok := rImageFingerprint["v1Name"]; ok {
						if s, ok := rImageFingerprint["v1Name"].(string); ok {
							r.Image.Fingerprint.V1Name = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.Image.Fingerprint.V1Name: expected string")
						}
					}
					if _, ok := rImageFingerprint["v2Blob"]; ok {
						if s, ok := rImageFingerprint["v2Blob"].([]interface{}); ok {
							for _, ss := range s {
								if strval, ok := ss.(string); ok {
									r.Image.Fingerprint.V2Blob = append(r.Image.Fingerprint.V2Blob, strval)
								}
							}
						} else {
							return nil, fmt.Errorf("r.Image.Fingerprint.V2Blob: expected []interface{}")
						}
					}
					if _, ok := rImageFingerprint["v2Name"]; ok {
						if s, ok := rImageFingerprint["v2Name"].(string); ok {
							r.Image.Fingerprint.V2Name = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.Image.Fingerprint.V2Name: expected string")
						}
					}
				} else {
					return nil, fmt.Errorf("r.Image.Fingerprint: expected map[string]interface{}")
				}
			}
			if _, ok := rImage["resourceUrl"]; ok {
				if s, ok := rImage["resourceUrl"].(string); ok {
					r.Image.ResourceUrl = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.Image.ResourceUrl: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.Image: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["longDescription"]; ok {
		if s, ok := u.Object["longDescription"].(string); ok {
			r.LongDescription = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.LongDescription: expected string")
		}
	}
	if _, ok := u.Object["name"]; ok {
		if s, ok := u.Object["name"].(string); ok {
			r.Name = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Name: expected string")
		}
	}
	if _, ok := u.Object["package"]; ok {
		if rPackage, ok := u.Object["package"].(map[string]interface{}); ok {
			r.Package = &dclService.NotePackage{}
			if _, ok := rPackage["distribution"]; ok {
				if s, ok := rPackage["distribution"].([]interface{}); ok {
					for _, o := range s {
						if objval, ok := o.(map[string]interface{}); ok {
							var rPackageDistribution dclService.NotePackageDistribution
							if _, ok := objval["architecture"]; ok {
								if s, ok := objval["architecture"].(string); ok {
									rPackageDistribution.Architecture = dclService.NotePackageDistributionArchitectureEnumRef(s)
								} else {
									return nil, fmt.Errorf("rPackageDistribution.Architecture: expected string")
								}
							}
							if _, ok := objval["cpeUri"]; ok {
								if s, ok := objval["cpeUri"].(string); ok {
									rPackageDistribution.CpeUri = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rPackageDistribution.CpeUri: expected string")
								}
							}
							if _, ok := objval["description"]; ok {
								if s, ok := objval["description"].(string); ok {
									rPackageDistribution.Description = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rPackageDistribution.Description: expected string")
								}
							}
							if _, ok := objval["latestVersion"]; ok {
								if rPackageDistributionLatestVersion, ok := objval["latestVersion"].(map[string]interface{}); ok {
									rPackageDistribution.LatestVersion = &dclService.NotePackageDistributionLatestVersion{}
									if _, ok := rPackageDistributionLatestVersion["epoch"]; ok {
										if i, ok := rPackageDistributionLatestVersion["epoch"].(int64); ok {
											rPackageDistribution.LatestVersion.Epoch = dcl.Int64(i)
										} else {
											return nil, fmt.Errorf("rPackageDistribution.LatestVersion.Epoch: expected int64")
										}
									}
									if _, ok := rPackageDistributionLatestVersion["fullName"]; ok {
										if s, ok := rPackageDistributionLatestVersion["fullName"].(string); ok {
											rPackageDistribution.LatestVersion.FullName = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rPackageDistribution.LatestVersion.FullName: expected string")
										}
									}
									if _, ok := rPackageDistributionLatestVersion["kind"]; ok {
										if s, ok := rPackageDistributionLatestVersion["kind"].(string); ok {
											rPackageDistribution.LatestVersion.Kind = dclService.NotePackageDistributionLatestVersionKindEnumRef(s)
										} else {
											return nil, fmt.Errorf("rPackageDistribution.LatestVersion.Kind: expected string")
										}
									}
									if _, ok := rPackageDistributionLatestVersion["name"]; ok {
										if s, ok := rPackageDistributionLatestVersion["name"].(string); ok {
											rPackageDistribution.LatestVersion.Name = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rPackageDistribution.LatestVersion.Name: expected string")
										}
									}
									if _, ok := rPackageDistributionLatestVersion["revision"]; ok {
										if s, ok := rPackageDistributionLatestVersion["revision"].(string); ok {
											rPackageDistribution.LatestVersion.Revision = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rPackageDistribution.LatestVersion.Revision: expected string")
										}
									}
								} else {
									return nil, fmt.Errorf("rPackageDistribution.LatestVersion: expected map[string]interface{}")
								}
							}
							if _, ok := objval["maintainer"]; ok {
								if s, ok := objval["maintainer"].(string); ok {
									rPackageDistribution.Maintainer = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rPackageDistribution.Maintainer: expected string")
								}
							}
							if _, ok := objval["url"]; ok {
								if s, ok := objval["url"].(string); ok {
									rPackageDistribution.Url = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rPackageDistribution.Url: expected string")
								}
							}
							r.Package.Distribution = append(r.Package.Distribution, rPackageDistribution)
						}
					}
				} else {
					return nil, fmt.Errorf("r.Package.Distribution: expected []interface{}")
				}
			}
			if _, ok := rPackage["name"]; ok {
				if s, ok := rPackage["name"].(string); ok {
					r.Package.Name = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.Package.Name: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.Package: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["project"]; ok {
		if s, ok := u.Object["project"].(string); ok {
			r.Project = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Project: expected string")
		}
	}
	if _, ok := u.Object["relatedNoteNames"]; ok {
		if s, ok := u.Object["relatedNoteNames"].([]interface{}); ok {
			for _, ss := range s {
				if strval, ok := ss.(string); ok {
					r.RelatedNoteNames = append(r.RelatedNoteNames, strval)
				}
			}
		} else {
			return nil, fmt.Errorf("r.RelatedNoteNames: expected []interface{}")
		}
	}
	if _, ok := u.Object["relatedUrl"]; ok {
		if s, ok := u.Object["relatedUrl"].([]interface{}); ok {
			for _, o := range s {
				if objval, ok := o.(map[string]interface{}); ok {
					var rRelatedUrl dclService.NoteRelatedUrl
					if _, ok := objval["label"]; ok {
						if s, ok := objval["label"].(string); ok {
							rRelatedUrl.Label = dcl.String(s)
						} else {
							return nil, fmt.Errorf("rRelatedUrl.Label: expected string")
						}
					}
					if _, ok := objval["url"]; ok {
						if s, ok := objval["url"].(string); ok {
							rRelatedUrl.Url = dcl.String(s)
						} else {
							return nil, fmt.Errorf("rRelatedUrl.Url: expected string")
						}
					}
					r.RelatedUrl = append(r.RelatedUrl, rRelatedUrl)
				}
			}
		} else {
			return nil, fmt.Errorf("r.RelatedUrl: expected []interface{}")
		}
	}
	if _, ok := u.Object["shortDescription"]; ok {
		if s, ok := u.Object["shortDescription"].(string); ok {
			r.ShortDescription = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.ShortDescription: expected string")
		}
	}
	if _, ok := u.Object["updateTime"]; ok {
		if s, ok := u.Object["updateTime"].(string); ok {
			r.UpdateTime = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.UpdateTime: expected string")
		}
	}
	if _, ok := u.Object["vulnerability"]; ok {
		if rVulnerability, ok := u.Object["vulnerability"].(map[string]interface{}); ok {
			r.Vulnerability = &dclService.NoteVulnerability{}
			if _, ok := rVulnerability["cvssScore"]; ok {
				if f, ok := rVulnerability["cvssScore"].(float64); ok {
					r.Vulnerability.CvssScore = dcl.Float64(f)
				} else {
					return nil, fmt.Errorf("r.Vulnerability.CvssScore: expected float64")
				}
			}
			if _, ok := rVulnerability["cvssV3"]; ok {
				if rVulnerabilityCvssV3, ok := rVulnerability["cvssV3"].(map[string]interface{}); ok {
					r.Vulnerability.CvssV3 = &dclService.NoteVulnerabilityCvssV3{}
					if _, ok := rVulnerabilityCvssV3["attackComplexity"]; ok {
						if s, ok := rVulnerabilityCvssV3["attackComplexity"].(string); ok {
							r.Vulnerability.CvssV3.AttackComplexity = dclService.NoteVulnerabilityCvssV3AttackComplexityEnumRef(s)
						} else {
							return nil, fmt.Errorf("r.Vulnerability.CvssV3.AttackComplexity: expected string")
						}
					}
					if _, ok := rVulnerabilityCvssV3["attackVector"]; ok {
						if s, ok := rVulnerabilityCvssV3["attackVector"].(string); ok {
							r.Vulnerability.CvssV3.AttackVector = dclService.NoteVulnerabilityCvssV3AttackVectorEnumRef(s)
						} else {
							return nil, fmt.Errorf("r.Vulnerability.CvssV3.AttackVector: expected string")
						}
					}
					if _, ok := rVulnerabilityCvssV3["availabilityImpact"]; ok {
						if s, ok := rVulnerabilityCvssV3["availabilityImpact"].(string); ok {
							r.Vulnerability.CvssV3.AvailabilityImpact = dclService.NoteVulnerabilityCvssV3AvailabilityImpactEnumRef(s)
						} else {
							return nil, fmt.Errorf("r.Vulnerability.CvssV3.AvailabilityImpact: expected string")
						}
					}
					if _, ok := rVulnerabilityCvssV3["baseScore"]; ok {
						if f, ok := rVulnerabilityCvssV3["baseScore"].(float64); ok {
							r.Vulnerability.CvssV3.BaseScore = dcl.Float64(f)
						} else {
							return nil, fmt.Errorf("r.Vulnerability.CvssV3.BaseScore: expected float64")
						}
					}
					if _, ok := rVulnerabilityCvssV3["confidentialityImpact"]; ok {
						if s, ok := rVulnerabilityCvssV3["confidentialityImpact"].(string); ok {
							r.Vulnerability.CvssV3.ConfidentialityImpact = dclService.NoteVulnerabilityCvssV3ConfidentialityImpactEnumRef(s)
						} else {
							return nil, fmt.Errorf("r.Vulnerability.CvssV3.ConfidentialityImpact: expected string")
						}
					}
					if _, ok := rVulnerabilityCvssV3["exploitabilityScore"]; ok {
						if f, ok := rVulnerabilityCvssV3["exploitabilityScore"].(float64); ok {
							r.Vulnerability.CvssV3.ExploitabilityScore = dcl.Float64(f)
						} else {
							return nil, fmt.Errorf("r.Vulnerability.CvssV3.ExploitabilityScore: expected float64")
						}
					}
					if _, ok := rVulnerabilityCvssV3["impactScore"]; ok {
						if f, ok := rVulnerabilityCvssV3["impactScore"].(float64); ok {
							r.Vulnerability.CvssV3.ImpactScore = dcl.Float64(f)
						} else {
							return nil, fmt.Errorf("r.Vulnerability.CvssV3.ImpactScore: expected float64")
						}
					}
					if _, ok := rVulnerabilityCvssV3["integrityImpact"]; ok {
						if s, ok := rVulnerabilityCvssV3["integrityImpact"].(string); ok {
							r.Vulnerability.CvssV3.IntegrityImpact = dclService.NoteVulnerabilityCvssV3IntegrityImpactEnumRef(s)
						} else {
							return nil, fmt.Errorf("r.Vulnerability.CvssV3.IntegrityImpact: expected string")
						}
					}
					if _, ok := rVulnerabilityCvssV3["privilegesRequired"]; ok {
						if s, ok := rVulnerabilityCvssV3["privilegesRequired"].(string); ok {
							r.Vulnerability.CvssV3.PrivilegesRequired = dclService.NoteVulnerabilityCvssV3PrivilegesRequiredEnumRef(s)
						} else {
							return nil, fmt.Errorf("r.Vulnerability.CvssV3.PrivilegesRequired: expected string")
						}
					}
					if _, ok := rVulnerabilityCvssV3["scope"]; ok {
						if s, ok := rVulnerabilityCvssV3["scope"].(string); ok {
							r.Vulnerability.CvssV3.Scope = dclService.NoteVulnerabilityCvssV3ScopeEnumRef(s)
						} else {
							return nil, fmt.Errorf("r.Vulnerability.CvssV3.Scope: expected string")
						}
					}
					if _, ok := rVulnerabilityCvssV3["userInteraction"]; ok {
						if s, ok := rVulnerabilityCvssV3["userInteraction"].(string); ok {
							r.Vulnerability.CvssV3.UserInteraction = dclService.NoteVulnerabilityCvssV3UserInteractionEnumRef(s)
						} else {
							return nil, fmt.Errorf("r.Vulnerability.CvssV3.UserInteraction: expected string")
						}
					}
				} else {
					return nil, fmt.Errorf("r.Vulnerability.CvssV3: expected map[string]interface{}")
				}
			}
			if _, ok := rVulnerability["details"]; ok {
				if s, ok := rVulnerability["details"].([]interface{}); ok {
					for _, o := range s {
						if objval, ok := o.(map[string]interface{}); ok {
							var rVulnerabilityDetails dclService.NoteVulnerabilityDetails
							if _, ok := objval["affectedCpeUri"]; ok {
								if s, ok := objval["affectedCpeUri"].(string); ok {
									rVulnerabilityDetails.AffectedCpeUri = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rVulnerabilityDetails.AffectedCpeUri: expected string")
								}
							}
							if _, ok := objval["affectedPackage"]; ok {
								if s, ok := objval["affectedPackage"].(string); ok {
									rVulnerabilityDetails.AffectedPackage = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rVulnerabilityDetails.AffectedPackage: expected string")
								}
							}
							if _, ok := objval["affectedVersionEnd"]; ok {
								if rVulnerabilityDetailsAffectedVersionEnd, ok := objval["affectedVersionEnd"].(map[string]interface{}); ok {
									rVulnerabilityDetails.AffectedVersionEnd = &dclService.NoteVulnerabilityDetailsAffectedVersionEnd{}
									if _, ok := rVulnerabilityDetailsAffectedVersionEnd["epoch"]; ok {
										if i, ok := rVulnerabilityDetailsAffectedVersionEnd["epoch"].(int64); ok {
											rVulnerabilityDetails.AffectedVersionEnd.Epoch = dcl.Int64(i)
										} else {
											return nil, fmt.Errorf("rVulnerabilityDetails.AffectedVersionEnd.Epoch: expected int64")
										}
									}
									if _, ok := rVulnerabilityDetailsAffectedVersionEnd["fullName"]; ok {
										if s, ok := rVulnerabilityDetailsAffectedVersionEnd["fullName"].(string); ok {
											rVulnerabilityDetails.AffectedVersionEnd.FullName = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rVulnerabilityDetails.AffectedVersionEnd.FullName: expected string")
										}
									}
									if _, ok := rVulnerabilityDetailsAffectedVersionEnd["kind"]; ok {
										if s, ok := rVulnerabilityDetailsAffectedVersionEnd["kind"].(string); ok {
											rVulnerabilityDetails.AffectedVersionEnd.Kind = dclService.NoteVulnerabilityDetailsAffectedVersionEndKindEnumRef(s)
										} else {
											return nil, fmt.Errorf("rVulnerabilityDetails.AffectedVersionEnd.Kind: expected string")
										}
									}
									if _, ok := rVulnerabilityDetailsAffectedVersionEnd["name"]; ok {
										if s, ok := rVulnerabilityDetailsAffectedVersionEnd["name"].(string); ok {
											rVulnerabilityDetails.AffectedVersionEnd.Name = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rVulnerabilityDetails.AffectedVersionEnd.Name: expected string")
										}
									}
									if _, ok := rVulnerabilityDetailsAffectedVersionEnd["revision"]; ok {
										if s, ok := rVulnerabilityDetailsAffectedVersionEnd["revision"].(string); ok {
											rVulnerabilityDetails.AffectedVersionEnd.Revision = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rVulnerabilityDetails.AffectedVersionEnd.Revision: expected string")
										}
									}
								} else {
									return nil, fmt.Errorf("rVulnerabilityDetails.AffectedVersionEnd: expected map[string]interface{}")
								}
							}
							if _, ok := objval["affectedVersionStart"]; ok {
								if rVulnerabilityDetailsAffectedVersionStart, ok := objval["affectedVersionStart"].(map[string]interface{}); ok {
									rVulnerabilityDetails.AffectedVersionStart = &dclService.NoteVulnerabilityDetailsAffectedVersionStart{}
									if _, ok := rVulnerabilityDetailsAffectedVersionStart["epoch"]; ok {
										if i, ok := rVulnerabilityDetailsAffectedVersionStart["epoch"].(int64); ok {
											rVulnerabilityDetails.AffectedVersionStart.Epoch = dcl.Int64(i)
										} else {
											return nil, fmt.Errorf("rVulnerabilityDetails.AffectedVersionStart.Epoch: expected int64")
										}
									}
									if _, ok := rVulnerabilityDetailsAffectedVersionStart["fullName"]; ok {
										if s, ok := rVulnerabilityDetailsAffectedVersionStart["fullName"].(string); ok {
											rVulnerabilityDetails.AffectedVersionStart.FullName = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rVulnerabilityDetails.AffectedVersionStart.FullName: expected string")
										}
									}
									if _, ok := rVulnerabilityDetailsAffectedVersionStart["kind"]; ok {
										if s, ok := rVulnerabilityDetailsAffectedVersionStart["kind"].(string); ok {
											rVulnerabilityDetails.AffectedVersionStart.Kind = dclService.NoteVulnerabilityDetailsAffectedVersionStartKindEnumRef(s)
										} else {
											return nil, fmt.Errorf("rVulnerabilityDetails.AffectedVersionStart.Kind: expected string")
										}
									}
									if _, ok := rVulnerabilityDetailsAffectedVersionStart["name"]; ok {
										if s, ok := rVulnerabilityDetailsAffectedVersionStart["name"].(string); ok {
											rVulnerabilityDetails.AffectedVersionStart.Name = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rVulnerabilityDetails.AffectedVersionStart.Name: expected string")
										}
									}
									if _, ok := rVulnerabilityDetailsAffectedVersionStart["revision"]; ok {
										if s, ok := rVulnerabilityDetailsAffectedVersionStart["revision"].(string); ok {
											rVulnerabilityDetails.AffectedVersionStart.Revision = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rVulnerabilityDetails.AffectedVersionStart.Revision: expected string")
										}
									}
								} else {
									return nil, fmt.Errorf("rVulnerabilityDetails.AffectedVersionStart: expected map[string]interface{}")
								}
							}
							if _, ok := objval["description"]; ok {
								if s, ok := objval["description"].(string); ok {
									rVulnerabilityDetails.Description = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rVulnerabilityDetails.Description: expected string")
								}
							}
							if _, ok := objval["fixedCpeUri"]; ok {
								if s, ok := objval["fixedCpeUri"].(string); ok {
									rVulnerabilityDetails.FixedCpeUri = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rVulnerabilityDetails.FixedCpeUri: expected string")
								}
							}
							if _, ok := objval["fixedPackage"]; ok {
								if s, ok := objval["fixedPackage"].(string); ok {
									rVulnerabilityDetails.FixedPackage = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rVulnerabilityDetails.FixedPackage: expected string")
								}
							}
							if _, ok := objval["fixedVersion"]; ok {
								if rVulnerabilityDetailsFixedVersion, ok := objval["fixedVersion"].(map[string]interface{}); ok {
									rVulnerabilityDetails.FixedVersion = &dclService.NoteVulnerabilityDetailsFixedVersion{}
									if _, ok := rVulnerabilityDetailsFixedVersion["epoch"]; ok {
										if i, ok := rVulnerabilityDetailsFixedVersion["epoch"].(int64); ok {
											rVulnerabilityDetails.FixedVersion.Epoch = dcl.Int64(i)
										} else {
											return nil, fmt.Errorf("rVulnerabilityDetails.FixedVersion.Epoch: expected int64")
										}
									}
									if _, ok := rVulnerabilityDetailsFixedVersion["fullName"]; ok {
										if s, ok := rVulnerabilityDetailsFixedVersion["fullName"].(string); ok {
											rVulnerabilityDetails.FixedVersion.FullName = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rVulnerabilityDetails.FixedVersion.FullName: expected string")
										}
									}
									if _, ok := rVulnerabilityDetailsFixedVersion["kind"]; ok {
										if s, ok := rVulnerabilityDetailsFixedVersion["kind"].(string); ok {
											rVulnerabilityDetails.FixedVersion.Kind = dclService.NoteVulnerabilityDetailsFixedVersionKindEnumRef(s)
										} else {
											return nil, fmt.Errorf("rVulnerabilityDetails.FixedVersion.Kind: expected string")
										}
									}
									if _, ok := rVulnerabilityDetailsFixedVersion["name"]; ok {
										if s, ok := rVulnerabilityDetailsFixedVersion["name"].(string); ok {
											rVulnerabilityDetails.FixedVersion.Name = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rVulnerabilityDetails.FixedVersion.Name: expected string")
										}
									}
									if _, ok := rVulnerabilityDetailsFixedVersion["revision"]; ok {
										if s, ok := rVulnerabilityDetailsFixedVersion["revision"].(string); ok {
											rVulnerabilityDetails.FixedVersion.Revision = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rVulnerabilityDetails.FixedVersion.Revision: expected string")
										}
									}
								} else {
									return nil, fmt.Errorf("rVulnerabilityDetails.FixedVersion: expected map[string]interface{}")
								}
							}
							if _, ok := objval["isObsolete"]; ok {
								if b, ok := objval["isObsolete"].(bool); ok {
									rVulnerabilityDetails.IsObsolete = dcl.Bool(b)
								} else {
									return nil, fmt.Errorf("rVulnerabilityDetails.IsObsolete: expected bool")
								}
							}
							if _, ok := objval["packageType"]; ok {
								if s, ok := objval["packageType"].(string); ok {
									rVulnerabilityDetails.PackageType = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rVulnerabilityDetails.PackageType: expected string")
								}
							}
							if _, ok := objval["severityName"]; ok {
								if s, ok := objval["severityName"].(string); ok {
									rVulnerabilityDetails.SeverityName = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rVulnerabilityDetails.SeverityName: expected string")
								}
							}
							if _, ok := objval["sourceUpdateTime"]; ok {
								if s, ok := objval["sourceUpdateTime"].(string); ok {
									rVulnerabilityDetails.SourceUpdateTime = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rVulnerabilityDetails.SourceUpdateTime: expected string")
								}
							}
							r.Vulnerability.Details = append(r.Vulnerability.Details, rVulnerabilityDetails)
						}
					}
				} else {
					return nil, fmt.Errorf("r.Vulnerability.Details: expected []interface{}")
				}
			}
			if _, ok := rVulnerability["severity"]; ok {
				if s, ok := rVulnerability["severity"].(string); ok {
					r.Vulnerability.Severity = dclService.NoteVulnerabilitySeverityEnumRef(s)
				} else {
					return nil, fmt.Errorf("r.Vulnerability.Severity: expected string")
				}
			}
			if _, ok := rVulnerability["sourceUpdateTime"]; ok {
				if s, ok := rVulnerability["sourceUpdateTime"].(string); ok {
					r.Vulnerability.SourceUpdateTime = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.Vulnerability.SourceUpdateTime: expected string")
				}
			}
			if _, ok := rVulnerability["windowsDetails"]; ok {
				if s, ok := rVulnerability["windowsDetails"].([]interface{}); ok {
					for _, o := range s {
						if objval, ok := o.(map[string]interface{}); ok {
							var rVulnerabilityWindowsDetails dclService.NoteVulnerabilityWindowsDetails
							if _, ok := objval["cpeUri"]; ok {
								if s, ok := objval["cpeUri"].(string); ok {
									rVulnerabilityWindowsDetails.CpeUri = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rVulnerabilityWindowsDetails.CpeUri: expected string")
								}
							}
							if _, ok := objval["description"]; ok {
								if s, ok := objval["description"].(string); ok {
									rVulnerabilityWindowsDetails.Description = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rVulnerabilityWindowsDetails.Description: expected string")
								}
							}
							if _, ok := objval["fixingKbs"]; ok {
								if s, ok := objval["fixingKbs"].([]interface{}); ok {
									for _, o := range s {
										if objval, ok := o.(map[string]interface{}); ok {
											var rVulnerabilityWindowsDetailsFixingKbs dclService.NoteVulnerabilityWindowsDetailsFixingKbs
											if _, ok := objval["name"]; ok {
												if s, ok := objval["name"].(string); ok {
													rVulnerabilityWindowsDetailsFixingKbs.Name = dcl.String(s)
												} else {
													return nil, fmt.Errorf("rVulnerabilityWindowsDetailsFixingKbs.Name: expected string")
												}
											}
											if _, ok := objval["url"]; ok {
												if s, ok := objval["url"].(string); ok {
													rVulnerabilityWindowsDetailsFixingKbs.Url = dcl.String(s)
												} else {
													return nil, fmt.Errorf("rVulnerabilityWindowsDetailsFixingKbs.Url: expected string")
												}
											}
											rVulnerabilityWindowsDetails.FixingKbs = append(rVulnerabilityWindowsDetails.FixingKbs, rVulnerabilityWindowsDetailsFixingKbs)
										}
									}
								} else {
									return nil, fmt.Errorf("rVulnerabilityWindowsDetails.FixingKbs: expected []interface{}")
								}
							}
							if _, ok := objval["name"]; ok {
								if s, ok := objval["name"].(string); ok {
									rVulnerabilityWindowsDetails.Name = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rVulnerabilityWindowsDetails.Name: expected string")
								}
							}
							r.Vulnerability.WindowsDetails = append(r.Vulnerability.WindowsDetails, rVulnerabilityWindowsDetails)
						}
					}
				} else {
					return nil, fmt.Errorf("r.Vulnerability.WindowsDetails: expected []interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.Vulnerability: expected map[string]interface{}")
		}
	}
	return r, nil
}

func GetNote(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToNote(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetNote(ctx, r)
	if err != nil {
		return nil, err
	}
	return NoteToUnstructured(r), nil
}

func ListNote(ctx context.Context, config *dcl.Config, project string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListNote(ctx, project)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, NoteToUnstructured(r))
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

func ApplyNote(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToNote(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToNote(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyNote(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return NoteToUnstructured(r), nil
}

func NoteHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToNote(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToNote(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyNote(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteNote(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToNote(u)
	if err != nil {
		return err
	}
	return c.DeleteNote(ctx, r)
}

func NoteID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToNote(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *Note) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"containeranalysis",
		"Note",
		"beta",
	}
}

func (r *Note) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Note) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Note) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *Note) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Note) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Note) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Note) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetNote(ctx, config, resource)
}

func (r *Note) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyNote(ctx, config, resource, opts...)
}

func (r *Note) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return NoteHasDiff(ctx, config, resource, opts...)
}

func (r *Note) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteNote(ctx, config, resource)
}

func (r *Note) ID(resource *unstructured.Resource) (string, error) {
	return NoteID(resource)
}

func init() {
	unstructured.Register(&Note{})
}
