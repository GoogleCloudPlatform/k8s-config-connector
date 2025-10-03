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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/containeranalysis/alpha/containeranalysis_alpha_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/containeranalysis/alpha"
)

// NoteServer implements the gRPC interface for Note.
type NoteServer struct{}

// ProtoToNoteVulnerabilitySeverityEnum converts a NoteVulnerabilitySeverityEnum enum from its proto representation.
func ProtoToContaineranalysisAlphaNoteVulnerabilitySeverityEnum(e alphapb.ContaineranalysisAlphaNoteVulnerabilitySeverityEnum) *alpha.NoteVulnerabilitySeverityEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ContaineranalysisAlphaNoteVulnerabilitySeverityEnum_name[int32(e)]; ok {
		e := alpha.NoteVulnerabilitySeverityEnum(n[len("ContaineranalysisAlphaNoteVulnerabilitySeverityEnum"):])
		return &e
	}
	return nil
}

// ProtoToNoteVulnerabilityDetailsAffectedVersionStartKindEnum converts a NoteVulnerabilityDetailsAffectedVersionStartKindEnum enum from its proto representation.
func ProtoToContaineranalysisAlphaNoteVulnerabilityDetailsAffectedVersionStartKindEnum(e alphapb.ContaineranalysisAlphaNoteVulnerabilityDetailsAffectedVersionStartKindEnum) *alpha.NoteVulnerabilityDetailsAffectedVersionStartKindEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ContaineranalysisAlphaNoteVulnerabilityDetailsAffectedVersionStartKindEnum_name[int32(e)]; ok {
		e := alpha.NoteVulnerabilityDetailsAffectedVersionStartKindEnum(n[len("ContaineranalysisAlphaNoteVulnerabilityDetailsAffectedVersionStartKindEnum"):])
		return &e
	}
	return nil
}

// ProtoToNoteVulnerabilityDetailsAffectedVersionEndKindEnum converts a NoteVulnerabilityDetailsAffectedVersionEndKindEnum enum from its proto representation.
func ProtoToContaineranalysisAlphaNoteVulnerabilityDetailsAffectedVersionEndKindEnum(e alphapb.ContaineranalysisAlphaNoteVulnerabilityDetailsAffectedVersionEndKindEnum) *alpha.NoteVulnerabilityDetailsAffectedVersionEndKindEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ContaineranalysisAlphaNoteVulnerabilityDetailsAffectedVersionEndKindEnum_name[int32(e)]; ok {
		e := alpha.NoteVulnerabilityDetailsAffectedVersionEndKindEnum(n[len("ContaineranalysisAlphaNoteVulnerabilityDetailsAffectedVersionEndKindEnum"):])
		return &e
	}
	return nil
}

// ProtoToNoteVulnerabilityDetailsFixedVersionKindEnum converts a NoteVulnerabilityDetailsFixedVersionKindEnum enum from its proto representation.
func ProtoToContaineranalysisAlphaNoteVulnerabilityDetailsFixedVersionKindEnum(e alphapb.ContaineranalysisAlphaNoteVulnerabilityDetailsFixedVersionKindEnum) *alpha.NoteVulnerabilityDetailsFixedVersionKindEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ContaineranalysisAlphaNoteVulnerabilityDetailsFixedVersionKindEnum_name[int32(e)]; ok {
		e := alpha.NoteVulnerabilityDetailsFixedVersionKindEnum(n[len("ContaineranalysisAlphaNoteVulnerabilityDetailsFixedVersionKindEnum"):])
		return &e
	}
	return nil
}

// ProtoToNoteVulnerabilityCvssV3AttackVectorEnum converts a NoteVulnerabilityCvssV3AttackVectorEnum enum from its proto representation.
func ProtoToContaineranalysisAlphaNoteVulnerabilityCvssV3AttackVectorEnum(e alphapb.ContaineranalysisAlphaNoteVulnerabilityCvssV3AttackVectorEnum) *alpha.NoteVulnerabilityCvssV3AttackVectorEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ContaineranalysisAlphaNoteVulnerabilityCvssV3AttackVectorEnum_name[int32(e)]; ok {
		e := alpha.NoteVulnerabilityCvssV3AttackVectorEnum(n[len("ContaineranalysisAlphaNoteVulnerabilityCvssV3AttackVectorEnum"):])
		return &e
	}
	return nil
}

// ProtoToNoteVulnerabilityCvssV3AttackComplexityEnum converts a NoteVulnerabilityCvssV3AttackComplexityEnum enum from its proto representation.
func ProtoToContaineranalysisAlphaNoteVulnerabilityCvssV3AttackComplexityEnum(e alphapb.ContaineranalysisAlphaNoteVulnerabilityCvssV3AttackComplexityEnum) *alpha.NoteVulnerabilityCvssV3AttackComplexityEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ContaineranalysisAlphaNoteVulnerabilityCvssV3AttackComplexityEnum_name[int32(e)]; ok {
		e := alpha.NoteVulnerabilityCvssV3AttackComplexityEnum(n[len("ContaineranalysisAlphaNoteVulnerabilityCvssV3AttackComplexityEnum"):])
		return &e
	}
	return nil
}

// ProtoToNoteVulnerabilityCvssV3PrivilegesRequiredEnum converts a NoteVulnerabilityCvssV3PrivilegesRequiredEnum enum from its proto representation.
func ProtoToContaineranalysisAlphaNoteVulnerabilityCvssV3PrivilegesRequiredEnum(e alphapb.ContaineranalysisAlphaNoteVulnerabilityCvssV3PrivilegesRequiredEnum) *alpha.NoteVulnerabilityCvssV3PrivilegesRequiredEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ContaineranalysisAlphaNoteVulnerabilityCvssV3PrivilegesRequiredEnum_name[int32(e)]; ok {
		e := alpha.NoteVulnerabilityCvssV3PrivilegesRequiredEnum(n[len("ContaineranalysisAlphaNoteVulnerabilityCvssV3PrivilegesRequiredEnum"):])
		return &e
	}
	return nil
}

// ProtoToNoteVulnerabilityCvssV3UserInteractionEnum converts a NoteVulnerabilityCvssV3UserInteractionEnum enum from its proto representation.
func ProtoToContaineranalysisAlphaNoteVulnerabilityCvssV3UserInteractionEnum(e alphapb.ContaineranalysisAlphaNoteVulnerabilityCvssV3UserInteractionEnum) *alpha.NoteVulnerabilityCvssV3UserInteractionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ContaineranalysisAlphaNoteVulnerabilityCvssV3UserInteractionEnum_name[int32(e)]; ok {
		e := alpha.NoteVulnerabilityCvssV3UserInteractionEnum(n[len("ContaineranalysisAlphaNoteVulnerabilityCvssV3UserInteractionEnum"):])
		return &e
	}
	return nil
}

// ProtoToNoteVulnerabilityCvssV3ScopeEnum converts a NoteVulnerabilityCvssV3ScopeEnum enum from its proto representation.
func ProtoToContaineranalysisAlphaNoteVulnerabilityCvssV3ScopeEnum(e alphapb.ContaineranalysisAlphaNoteVulnerabilityCvssV3ScopeEnum) *alpha.NoteVulnerabilityCvssV3ScopeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ContaineranalysisAlphaNoteVulnerabilityCvssV3ScopeEnum_name[int32(e)]; ok {
		e := alpha.NoteVulnerabilityCvssV3ScopeEnum(n[len("ContaineranalysisAlphaNoteVulnerabilityCvssV3ScopeEnum"):])
		return &e
	}
	return nil
}

// ProtoToNoteVulnerabilityCvssV3ConfidentialityImpactEnum converts a NoteVulnerabilityCvssV3ConfidentialityImpactEnum enum from its proto representation.
func ProtoToContaineranalysisAlphaNoteVulnerabilityCvssV3ConfidentialityImpactEnum(e alphapb.ContaineranalysisAlphaNoteVulnerabilityCvssV3ConfidentialityImpactEnum) *alpha.NoteVulnerabilityCvssV3ConfidentialityImpactEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ContaineranalysisAlphaNoteVulnerabilityCvssV3ConfidentialityImpactEnum_name[int32(e)]; ok {
		e := alpha.NoteVulnerabilityCvssV3ConfidentialityImpactEnum(n[len("ContaineranalysisAlphaNoteVulnerabilityCvssV3ConfidentialityImpactEnum"):])
		return &e
	}
	return nil
}

// ProtoToNoteVulnerabilityCvssV3IntegrityImpactEnum converts a NoteVulnerabilityCvssV3IntegrityImpactEnum enum from its proto representation.
func ProtoToContaineranalysisAlphaNoteVulnerabilityCvssV3IntegrityImpactEnum(e alphapb.ContaineranalysisAlphaNoteVulnerabilityCvssV3IntegrityImpactEnum) *alpha.NoteVulnerabilityCvssV3IntegrityImpactEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ContaineranalysisAlphaNoteVulnerabilityCvssV3IntegrityImpactEnum_name[int32(e)]; ok {
		e := alpha.NoteVulnerabilityCvssV3IntegrityImpactEnum(n[len("ContaineranalysisAlphaNoteVulnerabilityCvssV3IntegrityImpactEnum"):])
		return &e
	}
	return nil
}

// ProtoToNoteVulnerabilityCvssV3AvailabilityImpactEnum converts a NoteVulnerabilityCvssV3AvailabilityImpactEnum enum from its proto representation.
func ProtoToContaineranalysisAlphaNoteVulnerabilityCvssV3AvailabilityImpactEnum(e alphapb.ContaineranalysisAlphaNoteVulnerabilityCvssV3AvailabilityImpactEnum) *alpha.NoteVulnerabilityCvssV3AvailabilityImpactEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ContaineranalysisAlphaNoteVulnerabilityCvssV3AvailabilityImpactEnum_name[int32(e)]; ok {
		e := alpha.NoteVulnerabilityCvssV3AvailabilityImpactEnum(n[len("ContaineranalysisAlphaNoteVulnerabilityCvssV3AvailabilityImpactEnum"):])
		return &e
	}
	return nil
}

// ProtoToNoteBuildSignatureKeyTypeEnum converts a NoteBuildSignatureKeyTypeEnum enum from its proto representation.
func ProtoToContaineranalysisAlphaNoteBuildSignatureKeyTypeEnum(e alphapb.ContaineranalysisAlphaNoteBuildSignatureKeyTypeEnum) *alpha.NoteBuildSignatureKeyTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ContaineranalysisAlphaNoteBuildSignatureKeyTypeEnum_name[int32(e)]; ok {
		e := alpha.NoteBuildSignatureKeyTypeEnum(n[len("ContaineranalysisAlphaNoteBuildSignatureKeyTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToNotePackageDistributionArchitectureEnum converts a NotePackageDistributionArchitectureEnum enum from its proto representation.
func ProtoToContaineranalysisAlphaNotePackageDistributionArchitectureEnum(e alphapb.ContaineranalysisAlphaNotePackageDistributionArchitectureEnum) *alpha.NotePackageDistributionArchitectureEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ContaineranalysisAlphaNotePackageDistributionArchitectureEnum_name[int32(e)]; ok {
		e := alpha.NotePackageDistributionArchitectureEnum(n[len("ContaineranalysisAlphaNotePackageDistributionArchitectureEnum"):])
		return &e
	}
	return nil
}

// ProtoToNotePackageDistributionLatestVersionKindEnum converts a NotePackageDistributionLatestVersionKindEnum enum from its proto representation.
func ProtoToContaineranalysisAlphaNotePackageDistributionLatestVersionKindEnum(e alphapb.ContaineranalysisAlphaNotePackageDistributionLatestVersionKindEnum) *alpha.NotePackageDistributionLatestVersionKindEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ContaineranalysisAlphaNotePackageDistributionLatestVersionKindEnum_name[int32(e)]; ok {
		e := alpha.NotePackageDistributionLatestVersionKindEnum(n[len("ContaineranalysisAlphaNotePackageDistributionLatestVersionKindEnum"):])
		return &e
	}
	return nil
}

// ProtoToNoteDiscoveryAnalysisKindEnum converts a NoteDiscoveryAnalysisKindEnum enum from its proto representation.
func ProtoToContaineranalysisAlphaNoteDiscoveryAnalysisKindEnum(e alphapb.ContaineranalysisAlphaNoteDiscoveryAnalysisKindEnum) *alpha.NoteDiscoveryAnalysisKindEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ContaineranalysisAlphaNoteDiscoveryAnalysisKindEnum_name[int32(e)]; ok {
		e := alpha.NoteDiscoveryAnalysisKindEnum(n[len("ContaineranalysisAlphaNoteDiscoveryAnalysisKindEnum"):])
		return &e
	}
	return nil
}

// ProtoToNoteRelatedUrl converts a NoteRelatedUrl object from its proto representation.
func ProtoToContaineranalysisAlphaNoteRelatedUrl(p *alphapb.ContaineranalysisAlphaNoteRelatedUrl) *alpha.NoteRelatedUrl {
	if p == nil {
		return nil
	}
	obj := &alpha.NoteRelatedUrl{
		Url:   dcl.StringOrNil(p.GetUrl()),
		Label: dcl.StringOrNil(p.GetLabel()),
	}
	return obj
}

// ProtoToNoteVulnerability converts a NoteVulnerability object from its proto representation.
func ProtoToContaineranalysisAlphaNoteVulnerability(p *alphapb.ContaineranalysisAlphaNoteVulnerability) *alpha.NoteVulnerability {
	if p == nil {
		return nil
	}
	obj := &alpha.NoteVulnerability{
		CvssScore:        dcl.Float64OrNil(p.GetCvssScore()),
		Severity:         ProtoToContaineranalysisAlphaNoteVulnerabilitySeverityEnum(p.GetSeverity()),
		CvssV3:           ProtoToContaineranalysisAlphaNoteVulnerabilityCvssV3(p.GetCvssV3()),
		SourceUpdateTime: dcl.StringOrNil(p.GetSourceUpdateTime()),
	}
	for _, r := range p.GetDetails() {
		obj.Details = append(obj.Details, *ProtoToContaineranalysisAlphaNoteVulnerabilityDetails(r))
	}
	for _, r := range p.GetWindowsDetails() {
		obj.WindowsDetails = append(obj.WindowsDetails, *ProtoToContaineranalysisAlphaNoteVulnerabilityWindowsDetails(r))
	}
	return obj
}

// ProtoToNoteVulnerabilityDetails converts a NoteVulnerabilityDetails object from its proto representation.
func ProtoToContaineranalysisAlphaNoteVulnerabilityDetails(p *alphapb.ContaineranalysisAlphaNoteVulnerabilityDetails) *alpha.NoteVulnerabilityDetails {
	if p == nil {
		return nil
	}
	obj := &alpha.NoteVulnerabilityDetails{
		SeverityName:         dcl.StringOrNil(p.GetSeverityName()),
		Description:          dcl.StringOrNil(p.GetDescription()),
		PackageType:          dcl.StringOrNil(p.GetPackageType()),
		AffectedCpeUri:       dcl.StringOrNil(p.GetAffectedCpeUri()),
		AffectedPackage:      dcl.StringOrNil(p.GetAffectedPackage()),
		AffectedVersionStart: ProtoToContaineranalysisAlphaNoteVulnerabilityDetailsAffectedVersionStart(p.GetAffectedVersionStart()),
		AffectedVersionEnd:   ProtoToContaineranalysisAlphaNoteVulnerabilityDetailsAffectedVersionEnd(p.GetAffectedVersionEnd()),
		FixedCpeUri:          dcl.StringOrNil(p.GetFixedCpeUri()),
		FixedPackage:         dcl.StringOrNil(p.GetFixedPackage()),
		FixedVersion:         ProtoToContaineranalysisAlphaNoteVulnerabilityDetailsFixedVersion(p.GetFixedVersion()),
		IsObsolete:           dcl.Bool(p.GetIsObsolete()),
		SourceUpdateTime:     dcl.StringOrNil(p.GetSourceUpdateTime()),
	}
	return obj
}

// ProtoToNoteVulnerabilityDetailsAffectedVersionStart converts a NoteVulnerabilityDetailsAffectedVersionStart object from its proto representation.
func ProtoToContaineranalysisAlphaNoteVulnerabilityDetailsAffectedVersionStart(p *alphapb.ContaineranalysisAlphaNoteVulnerabilityDetailsAffectedVersionStart) *alpha.NoteVulnerabilityDetailsAffectedVersionStart {
	if p == nil {
		return nil
	}
	obj := &alpha.NoteVulnerabilityDetailsAffectedVersionStart{
		Epoch:    dcl.Int64OrNil(p.GetEpoch()),
		Name:     dcl.StringOrNil(p.GetName()),
		Revision: dcl.StringOrNil(p.GetRevision()),
		Kind:     ProtoToContaineranalysisAlphaNoteVulnerabilityDetailsAffectedVersionStartKindEnum(p.GetKind()),
		FullName: dcl.StringOrNil(p.GetFullName()),
	}
	return obj
}

// ProtoToNoteVulnerabilityDetailsAffectedVersionEnd converts a NoteVulnerabilityDetailsAffectedVersionEnd object from its proto representation.
func ProtoToContaineranalysisAlphaNoteVulnerabilityDetailsAffectedVersionEnd(p *alphapb.ContaineranalysisAlphaNoteVulnerabilityDetailsAffectedVersionEnd) *alpha.NoteVulnerabilityDetailsAffectedVersionEnd {
	if p == nil {
		return nil
	}
	obj := &alpha.NoteVulnerabilityDetailsAffectedVersionEnd{
		Epoch:    dcl.Int64OrNil(p.GetEpoch()),
		Name:     dcl.StringOrNil(p.GetName()),
		Revision: dcl.StringOrNil(p.GetRevision()),
		Kind:     ProtoToContaineranalysisAlphaNoteVulnerabilityDetailsAffectedVersionEndKindEnum(p.GetKind()),
		FullName: dcl.StringOrNil(p.GetFullName()),
	}
	return obj
}

// ProtoToNoteVulnerabilityDetailsFixedVersion converts a NoteVulnerabilityDetailsFixedVersion object from its proto representation.
func ProtoToContaineranalysisAlphaNoteVulnerabilityDetailsFixedVersion(p *alphapb.ContaineranalysisAlphaNoteVulnerabilityDetailsFixedVersion) *alpha.NoteVulnerabilityDetailsFixedVersion {
	if p == nil {
		return nil
	}
	obj := &alpha.NoteVulnerabilityDetailsFixedVersion{
		Epoch:    dcl.Int64OrNil(p.GetEpoch()),
		Name:     dcl.StringOrNil(p.GetName()),
		Revision: dcl.StringOrNil(p.GetRevision()),
		Kind:     ProtoToContaineranalysisAlphaNoteVulnerabilityDetailsFixedVersionKindEnum(p.GetKind()),
		FullName: dcl.StringOrNil(p.GetFullName()),
	}
	return obj
}

// ProtoToNoteVulnerabilityCvssV3 converts a NoteVulnerabilityCvssV3 object from its proto representation.
func ProtoToContaineranalysisAlphaNoteVulnerabilityCvssV3(p *alphapb.ContaineranalysisAlphaNoteVulnerabilityCvssV3) *alpha.NoteVulnerabilityCvssV3 {
	if p == nil {
		return nil
	}
	obj := &alpha.NoteVulnerabilityCvssV3{
		BaseScore:             dcl.Float64OrNil(p.GetBaseScore()),
		ExploitabilityScore:   dcl.Float64OrNil(p.GetExploitabilityScore()),
		ImpactScore:           dcl.Float64OrNil(p.GetImpactScore()),
		AttackVector:          ProtoToContaineranalysisAlphaNoteVulnerabilityCvssV3AttackVectorEnum(p.GetAttackVector()),
		AttackComplexity:      ProtoToContaineranalysisAlphaNoteVulnerabilityCvssV3AttackComplexityEnum(p.GetAttackComplexity()),
		PrivilegesRequired:    ProtoToContaineranalysisAlphaNoteVulnerabilityCvssV3PrivilegesRequiredEnum(p.GetPrivilegesRequired()),
		UserInteraction:       ProtoToContaineranalysisAlphaNoteVulnerabilityCvssV3UserInteractionEnum(p.GetUserInteraction()),
		Scope:                 ProtoToContaineranalysisAlphaNoteVulnerabilityCvssV3ScopeEnum(p.GetScope()),
		ConfidentialityImpact: ProtoToContaineranalysisAlphaNoteVulnerabilityCvssV3ConfidentialityImpactEnum(p.GetConfidentialityImpact()),
		IntegrityImpact:       ProtoToContaineranalysisAlphaNoteVulnerabilityCvssV3IntegrityImpactEnum(p.GetIntegrityImpact()),
		AvailabilityImpact:    ProtoToContaineranalysisAlphaNoteVulnerabilityCvssV3AvailabilityImpactEnum(p.GetAvailabilityImpact()),
	}
	return obj
}

// ProtoToNoteVulnerabilityWindowsDetails converts a NoteVulnerabilityWindowsDetails object from its proto representation.
func ProtoToContaineranalysisAlphaNoteVulnerabilityWindowsDetails(p *alphapb.ContaineranalysisAlphaNoteVulnerabilityWindowsDetails) *alpha.NoteVulnerabilityWindowsDetails {
	if p == nil {
		return nil
	}
	obj := &alpha.NoteVulnerabilityWindowsDetails{
		CpeUri:      dcl.StringOrNil(p.GetCpeUri()),
		Name:        dcl.StringOrNil(p.GetName()),
		Description: dcl.StringOrNil(p.GetDescription()),
	}
	for _, r := range p.GetFixingKbs() {
		obj.FixingKbs = append(obj.FixingKbs, *ProtoToContaineranalysisAlphaNoteVulnerabilityWindowsDetailsFixingKbs(r))
	}
	return obj
}

// ProtoToNoteVulnerabilityWindowsDetailsFixingKbs converts a NoteVulnerabilityWindowsDetailsFixingKbs object from its proto representation.
func ProtoToContaineranalysisAlphaNoteVulnerabilityWindowsDetailsFixingKbs(p *alphapb.ContaineranalysisAlphaNoteVulnerabilityWindowsDetailsFixingKbs) *alpha.NoteVulnerabilityWindowsDetailsFixingKbs {
	if p == nil {
		return nil
	}
	obj := &alpha.NoteVulnerabilityWindowsDetailsFixingKbs{
		Name: dcl.StringOrNil(p.GetName()),
		Url:  dcl.StringOrNil(p.GetUrl()),
	}
	return obj
}

// ProtoToNoteBuild converts a NoteBuild object from its proto representation.
func ProtoToContaineranalysisAlphaNoteBuild(p *alphapb.ContaineranalysisAlphaNoteBuild) *alpha.NoteBuild {
	if p == nil {
		return nil
	}
	obj := &alpha.NoteBuild{
		BuilderVersion: dcl.StringOrNil(p.GetBuilderVersion()),
		Signature:      ProtoToContaineranalysisAlphaNoteBuildSignature(p.GetSignature()),
	}
	return obj
}

// ProtoToNoteBuildSignature converts a NoteBuildSignature object from its proto representation.
func ProtoToContaineranalysisAlphaNoteBuildSignature(p *alphapb.ContaineranalysisAlphaNoteBuildSignature) *alpha.NoteBuildSignature {
	if p == nil {
		return nil
	}
	obj := &alpha.NoteBuildSignature{
		PublicKey: dcl.StringOrNil(p.GetPublicKey()),
		Signature: dcl.StringOrNil(p.GetSignature()),
		KeyId:     dcl.StringOrNil(p.GetKeyId()),
		KeyType:   ProtoToContaineranalysisAlphaNoteBuildSignatureKeyTypeEnum(p.GetKeyType()),
	}
	return obj
}

// ProtoToNoteImage converts a NoteImage object from its proto representation.
func ProtoToContaineranalysisAlphaNoteImage(p *alphapb.ContaineranalysisAlphaNoteImage) *alpha.NoteImage {
	if p == nil {
		return nil
	}
	obj := &alpha.NoteImage{
		ResourceUrl: dcl.StringOrNil(p.GetResourceUrl()),
		Fingerprint: ProtoToContaineranalysisAlphaNoteImageFingerprint(p.GetFingerprint()),
	}
	return obj
}

// ProtoToNoteImageFingerprint converts a NoteImageFingerprint object from its proto representation.
func ProtoToContaineranalysisAlphaNoteImageFingerprint(p *alphapb.ContaineranalysisAlphaNoteImageFingerprint) *alpha.NoteImageFingerprint {
	if p == nil {
		return nil
	}
	obj := &alpha.NoteImageFingerprint{
		V1Name: dcl.StringOrNil(p.GetV1Name()),
		V2Name: dcl.StringOrNil(p.GetV2Name()),
	}
	for _, r := range p.GetV2Blob() {
		obj.V2Blob = append(obj.V2Blob, r)
	}
	return obj
}

// ProtoToNotePackage converts a NotePackage object from its proto representation.
func ProtoToContaineranalysisAlphaNotePackage(p *alphapb.ContaineranalysisAlphaNotePackage) *alpha.NotePackage {
	if p == nil {
		return nil
	}
	obj := &alpha.NotePackage{
		Name: dcl.StringOrNil(p.GetName()),
	}
	for _, r := range p.GetDistribution() {
		obj.Distribution = append(obj.Distribution, *ProtoToContaineranalysisAlphaNotePackageDistribution(r))
	}
	return obj
}

// ProtoToNotePackageDistribution converts a NotePackageDistribution object from its proto representation.
func ProtoToContaineranalysisAlphaNotePackageDistribution(p *alphapb.ContaineranalysisAlphaNotePackageDistribution) *alpha.NotePackageDistribution {
	if p == nil {
		return nil
	}
	obj := &alpha.NotePackageDistribution{
		CpeUri:        dcl.StringOrNil(p.GetCpeUri()),
		Architecture:  ProtoToContaineranalysisAlphaNotePackageDistributionArchitectureEnum(p.GetArchitecture()),
		LatestVersion: ProtoToContaineranalysisAlphaNotePackageDistributionLatestVersion(p.GetLatestVersion()),
		Maintainer:    dcl.StringOrNil(p.GetMaintainer()),
		Url:           dcl.StringOrNil(p.GetUrl()),
		Description:   dcl.StringOrNil(p.GetDescription()),
	}
	return obj
}

// ProtoToNotePackageDistributionLatestVersion converts a NotePackageDistributionLatestVersion object from its proto representation.
func ProtoToContaineranalysisAlphaNotePackageDistributionLatestVersion(p *alphapb.ContaineranalysisAlphaNotePackageDistributionLatestVersion) *alpha.NotePackageDistributionLatestVersion {
	if p == nil {
		return nil
	}
	obj := &alpha.NotePackageDistributionLatestVersion{
		Epoch:    dcl.Int64OrNil(p.GetEpoch()),
		Name:     dcl.StringOrNil(p.GetName()),
		Revision: dcl.StringOrNil(p.GetRevision()),
		Kind:     ProtoToContaineranalysisAlphaNotePackageDistributionLatestVersionKindEnum(p.GetKind()),
		FullName: dcl.StringOrNil(p.GetFullName()),
	}
	return obj
}

// ProtoToNoteDiscovery converts a NoteDiscovery object from its proto representation.
func ProtoToContaineranalysisAlphaNoteDiscovery(p *alphapb.ContaineranalysisAlphaNoteDiscovery) *alpha.NoteDiscovery {
	if p == nil {
		return nil
	}
	obj := &alpha.NoteDiscovery{
		AnalysisKind: ProtoToContaineranalysisAlphaNoteDiscoveryAnalysisKindEnum(p.GetAnalysisKind()),
	}
	return obj
}

// ProtoToNoteDeployment converts a NoteDeployment object from its proto representation.
func ProtoToContaineranalysisAlphaNoteDeployment(p *alphapb.ContaineranalysisAlphaNoteDeployment) *alpha.NoteDeployment {
	if p == nil {
		return nil
	}
	obj := &alpha.NoteDeployment{}
	for _, r := range p.GetResourceUri() {
		obj.ResourceUri = append(obj.ResourceUri, r)
	}
	return obj
}

// ProtoToNoteAttestation converts a NoteAttestation object from its proto representation.
func ProtoToContaineranalysisAlphaNoteAttestation(p *alphapb.ContaineranalysisAlphaNoteAttestation) *alpha.NoteAttestation {
	if p == nil {
		return nil
	}
	obj := &alpha.NoteAttestation{
		Hint: ProtoToContaineranalysisAlphaNoteAttestationHint(p.GetHint()),
	}
	return obj
}

// ProtoToNoteAttestationHint converts a NoteAttestationHint object from its proto representation.
func ProtoToContaineranalysisAlphaNoteAttestationHint(p *alphapb.ContaineranalysisAlphaNoteAttestationHint) *alpha.NoteAttestationHint {
	if p == nil {
		return nil
	}
	obj := &alpha.NoteAttestationHint{
		HumanReadableName: dcl.StringOrNil(p.GetHumanReadableName()),
	}
	return obj
}

// ProtoToNote converts a Note resource from its proto representation.
func ProtoToNote(p *alphapb.ContaineranalysisAlphaNote) *alpha.Note {
	obj := &alpha.Note{
		Name:             dcl.StringOrNil(p.GetName()),
		ShortDescription: dcl.StringOrNil(p.GetShortDescription()),
		LongDescription:  dcl.StringOrNil(p.GetLongDescription()),
		ExpirationTime:   dcl.StringOrNil(p.GetExpirationTime()),
		CreateTime:       dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:       dcl.StringOrNil(p.GetUpdateTime()),
		Vulnerability:    ProtoToContaineranalysisAlphaNoteVulnerability(p.GetVulnerability()),
		Build:            ProtoToContaineranalysisAlphaNoteBuild(p.GetBuild_()),
		Image:            ProtoToContaineranalysisAlphaNoteImage(p.GetImage()),
		Package:          ProtoToContaineranalysisAlphaNotePackage(p.GetPackage()),
		Discovery:        ProtoToContaineranalysisAlphaNoteDiscovery(p.GetDiscovery()),
		Deployment:       ProtoToContaineranalysisAlphaNoteDeployment(p.GetDeployment()),
		Attestation:      ProtoToContaineranalysisAlphaNoteAttestation(p.GetAttestation()),
		Project:          dcl.StringOrNil(p.GetProject()),
	}
	for _, r := range p.GetRelatedUrl() {
		obj.RelatedUrl = append(obj.RelatedUrl, *ProtoToContaineranalysisAlphaNoteRelatedUrl(r))
	}
	for _, r := range p.GetRelatedNoteNames() {
		obj.RelatedNoteNames = append(obj.RelatedNoteNames, r)
	}
	return obj
}

// NoteVulnerabilitySeverityEnumToProto converts a NoteVulnerabilitySeverityEnum enum to its proto representation.
func ContaineranalysisAlphaNoteVulnerabilitySeverityEnumToProto(e *alpha.NoteVulnerabilitySeverityEnum) alphapb.ContaineranalysisAlphaNoteVulnerabilitySeverityEnum {
	if e == nil {
		return alphapb.ContaineranalysisAlphaNoteVulnerabilitySeverityEnum(0)
	}
	if v, ok := alphapb.ContaineranalysisAlphaNoteVulnerabilitySeverityEnum_value["NoteVulnerabilitySeverityEnum"+string(*e)]; ok {
		return alphapb.ContaineranalysisAlphaNoteVulnerabilitySeverityEnum(v)
	}
	return alphapb.ContaineranalysisAlphaNoteVulnerabilitySeverityEnum(0)
}

// NoteVulnerabilityDetailsAffectedVersionStartKindEnumToProto converts a NoteVulnerabilityDetailsAffectedVersionStartKindEnum enum to its proto representation.
func ContaineranalysisAlphaNoteVulnerabilityDetailsAffectedVersionStartKindEnumToProto(e *alpha.NoteVulnerabilityDetailsAffectedVersionStartKindEnum) alphapb.ContaineranalysisAlphaNoteVulnerabilityDetailsAffectedVersionStartKindEnum {
	if e == nil {
		return alphapb.ContaineranalysisAlphaNoteVulnerabilityDetailsAffectedVersionStartKindEnum(0)
	}
	if v, ok := alphapb.ContaineranalysisAlphaNoteVulnerabilityDetailsAffectedVersionStartKindEnum_value["NoteVulnerabilityDetailsAffectedVersionStartKindEnum"+string(*e)]; ok {
		return alphapb.ContaineranalysisAlphaNoteVulnerabilityDetailsAffectedVersionStartKindEnum(v)
	}
	return alphapb.ContaineranalysisAlphaNoteVulnerabilityDetailsAffectedVersionStartKindEnum(0)
}

// NoteVulnerabilityDetailsAffectedVersionEndKindEnumToProto converts a NoteVulnerabilityDetailsAffectedVersionEndKindEnum enum to its proto representation.
func ContaineranalysisAlphaNoteVulnerabilityDetailsAffectedVersionEndKindEnumToProto(e *alpha.NoteVulnerabilityDetailsAffectedVersionEndKindEnum) alphapb.ContaineranalysisAlphaNoteVulnerabilityDetailsAffectedVersionEndKindEnum {
	if e == nil {
		return alphapb.ContaineranalysisAlphaNoteVulnerabilityDetailsAffectedVersionEndKindEnum(0)
	}
	if v, ok := alphapb.ContaineranalysisAlphaNoteVulnerabilityDetailsAffectedVersionEndKindEnum_value["NoteVulnerabilityDetailsAffectedVersionEndKindEnum"+string(*e)]; ok {
		return alphapb.ContaineranalysisAlphaNoteVulnerabilityDetailsAffectedVersionEndKindEnum(v)
	}
	return alphapb.ContaineranalysisAlphaNoteVulnerabilityDetailsAffectedVersionEndKindEnum(0)
}

// NoteVulnerabilityDetailsFixedVersionKindEnumToProto converts a NoteVulnerabilityDetailsFixedVersionKindEnum enum to its proto representation.
func ContaineranalysisAlphaNoteVulnerabilityDetailsFixedVersionKindEnumToProto(e *alpha.NoteVulnerabilityDetailsFixedVersionKindEnum) alphapb.ContaineranalysisAlphaNoteVulnerabilityDetailsFixedVersionKindEnum {
	if e == nil {
		return alphapb.ContaineranalysisAlphaNoteVulnerabilityDetailsFixedVersionKindEnum(0)
	}
	if v, ok := alphapb.ContaineranalysisAlphaNoteVulnerabilityDetailsFixedVersionKindEnum_value["NoteVulnerabilityDetailsFixedVersionKindEnum"+string(*e)]; ok {
		return alphapb.ContaineranalysisAlphaNoteVulnerabilityDetailsFixedVersionKindEnum(v)
	}
	return alphapb.ContaineranalysisAlphaNoteVulnerabilityDetailsFixedVersionKindEnum(0)
}

// NoteVulnerabilityCvssV3AttackVectorEnumToProto converts a NoteVulnerabilityCvssV3AttackVectorEnum enum to its proto representation.
func ContaineranalysisAlphaNoteVulnerabilityCvssV3AttackVectorEnumToProto(e *alpha.NoteVulnerabilityCvssV3AttackVectorEnum) alphapb.ContaineranalysisAlphaNoteVulnerabilityCvssV3AttackVectorEnum {
	if e == nil {
		return alphapb.ContaineranalysisAlphaNoteVulnerabilityCvssV3AttackVectorEnum(0)
	}
	if v, ok := alphapb.ContaineranalysisAlphaNoteVulnerabilityCvssV3AttackVectorEnum_value["NoteVulnerabilityCvssV3AttackVectorEnum"+string(*e)]; ok {
		return alphapb.ContaineranalysisAlphaNoteVulnerabilityCvssV3AttackVectorEnum(v)
	}
	return alphapb.ContaineranalysisAlphaNoteVulnerabilityCvssV3AttackVectorEnum(0)
}

// NoteVulnerabilityCvssV3AttackComplexityEnumToProto converts a NoteVulnerabilityCvssV3AttackComplexityEnum enum to its proto representation.
func ContaineranalysisAlphaNoteVulnerabilityCvssV3AttackComplexityEnumToProto(e *alpha.NoteVulnerabilityCvssV3AttackComplexityEnum) alphapb.ContaineranalysisAlphaNoteVulnerabilityCvssV3AttackComplexityEnum {
	if e == nil {
		return alphapb.ContaineranalysisAlphaNoteVulnerabilityCvssV3AttackComplexityEnum(0)
	}
	if v, ok := alphapb.ContaineranalysisAlphaNoteVulnerabilityCvssV3AttackComplexityEnum_value["NoteVulnerabilityCvssV3AttackComplexityEnum"+string(*e)]; ok {
		return alphapb.ContaineranalysisAlphaNoteVulnerabilityCvssV3AttackComplexityEnum(v)
	}
	return alphapb.ContaineranalysisAlphaNoteVulnerabilityCvssV3AttackComplexityEnum(0)
}

// NoteVulnerabilityCvssV3PrivilegesRequiredEnumToProto converts a NoteVulnerabilityCvssV3PrivilegesRequiredEnum enum to its proto representation.
func ContaineranalysisAlphaNoteVulnerabilityCvssV3PrivilegesRequiredEnumToProto(e *alpha.NoteVulnerabilityCvssV3PrivilegesRequiredEnum) alphapb.ContaineranalysisAlphaNoteVulnerabilityCvssV3PrivilegesRequiredEnum {
	if e == nil {
		return alphapb.ContaineranalysisAlphaNoteVulnerabilityCvssV3PrivilegesRequiredEnum(0)
	}
	if v, ok := alphapb.ContaineranalysisAlphaNoteVulnerabilityCvssV3PrivilegesRequiredEnum_value["NoteVulnerabilityCvssV3PrivilegesRequiredEnum"+string(*e)]; ok {
		return alphapb.ContaineranalysisAlphaNoteVulnerabilityCvssV3PrivilegesRequiredEnum(v)
	}
	return alphapb.ContaineranalysisAlphaNoteVulnerabilityCvssV3PrivilegesRequiredEnum(0)
}

// NoteVulnerabilityCvssV3UserInteractionEnumToProto converts a NoteVulnerabilityCvssV3UserInteractionEnum enum to its proto representation.
func ContaineranalysisAlphaNoteVulnerabilityCvssV3UserInteractionEnumToProto(e *alpha.NoteVulnerabilityCvssV3UserInteractionEnum) alphapb.ContaineranalysisAlphaNoteVulnerabilityCvssV3UserInteractionEnum {
	if e == nil {
		return alphapb.ContaineranalysisAlphaNoteVulnerabilityCvssV3UserInteractionEnum(0)
	}
	if v, ok := alphapb.ContaineranalysisAlphaNoteVulnerabilityCvssV3UserInteractionEnum_value["NoteVulnerabilityCvssV3UserInteractionEnum"+string(*e)]; ok {
		return alphapb.ContaineranalysisAlphaNoteVulnerabilityCvssV3UserInteractionEnum(v)
	}
	return alphapb.ContaineranalysisAlphaNoteVulnerabilityCvssV3UserInteractionEnum(0)
}

// NoteVulnerabilityCvssV3ScopeEnumToProto converts a NoteVulnerabilityCvssV3ScopeEnum enum to its proto representation.
func ContaineranalysisAlphaNoteVulnerabilityCvssV3ScopeEnumToProto(e *alpha.NoteVulnerabilityCvssV3ScopeEnum) alphapb.ContaineranalysisAlphaNoteVulnerabilityCvssV3ScopeEnum {
	if e == nil {
		return alphapb.ContaineranalysisAlphaNoteVulnerabilityCvssV3ScopeEnum(0)
	}
	if v, ok := alphapb.ContaineranalysisAlphaNoteVulnerabilityCvssV3ScopeEnum_value["NoteVulnerabilityCvssV3ScopeEnum"+string(*e)]; ok {
		return alphapb.ContaineranalysisAlphaNoteVulnerabilityCvssV3ScopeEnum(v)
	}
	return alphapb.ContaineranalysisAlphaNoteVulnerabilityCvssV3ScopeEnum(0)
}

// NoteVulnerabilityCvssV3ConfidentialityImpactEnumToProto converts a NoteVulnerabilityCvssV3ConfidentialityImpactEnum enum to its proto representation.
func ContaineranalysisAlphaNoteVulnerabilityCvssV3ConfidentialityImpactEnumToProto(e *alpha.NoteVulnerabilityCvssV3ConfidentialityImpactEnum) alphapb.ContaineranalysisAlphaNoteVulnerabilityCvssV3ConfidentialityImpactEnum {
	if e == nil {
		return alphapb.ContaineranalysisAlphaNoteVulnerabilityCvssV3ConfidentialityImpactEnum(0)
	}
	if v, ok := alphapb.ContaineranalysisAlphaNoteVulnerabilityCvssV3ConfidentialityImpactEnum_value["NoteVulnerabilityCvssV3ConfidentialityImpactEnum"+string(*e)]; ok {
		return alphapb.ContaineranalysisAlphaNoteVulnerabilityCvssV3ConfidentialityImpactEnum(v)
	}
	return alphapb.ContaineranalysisAlphaNoteVulnerabilityCvssV3ConfidentialityImpactEnum(0)
}

// NoteVulnerabilityCvssV3IntegrityImpactEnumToProto converts a NoteVulnerabilityCvssV3IntegrityImpactEnum enum to its proto representation.
func ContaineranalysisAlphaNoteVulnerabilityCvssV3IntegrityImpactEnumToProto(e *alpha.NoteVulnerabilityCvssV3IntegrityImpactEnum) alphapb.ContaineranalysisAlphaNoteVulnerabilityCvssV3IntegrityImpactEnum {
	if e == nil {
		return alphapb.ContaineranalysisAlphaNoteVulnerabilityCvssV3IntegrityImpactEnum(0)
	}
	if v, ok := alphapb.ContaineranalysisAlphaNoteVulnerabilityCvssV3IntegrityImpactEnum_value["NoteVulnerabilityCvssV3IntegrityImpactEnum"+string(*e)]; ok {
		return alphapb.ContaineranalysisAlphaNoteVulnerabilityCvssV3IntegrityImpactEnum(v)
	}
	return alphapb.ContaineranalysisAlphaNoteVulnerabilityCvssV3IntegrityImpactEnum(0)
}

// NoteVulnerabilityCvssV3AvailabilityImpactEnumToProto converts a NoteVulnerabilityCvssV3AvailabilityImpactEnum enum to its proto representation.
func ContaineranalysisAlphaNoteVulnerabilityCvssV3AvailabilityImpactEnumToProto(e *alpha.NoteVulnerabilityCvssV3AvailabilityImpactEnum) alphapb.ContaineranalysisAlphaNoteVulnerabilityCvssV3AvailabilityImpactEnum {
	if e == nil {
		return alphapb.ContaineranalysisAlphaNoteVulnerabilityCvssV3AvailabilityImpactEnum(0)
	}
	if v, ok := alphapb.ContaineranalysisAlphaNoteVulnerabilityCvssV3AvailabilityImpactEnum_value["NoteVulnerabilityCvssV3AvailabilityImpactEnum"+string(*e)]; ok {
		return alphapb.ContaineranalysisAlphaNoteVulnerabilityCvssV3AvailabilityImpactEnum(v)
	}
	return alphapb.ContaineranalysisAlphaNoteVulnerabilityCvssV3AvailabilityImpactEnum(0)
}

// NoteBuildSignatureKeyTypeEnumToProto converts a NoteBuildSignatureKeyTypeEnum enum to its proto representation.
func ContaineranalysisAlphaNoteBuildSignatureKeyTypeEnumToProto(e *alpha.NoteBuildSignatureKeyTypeEnum) alphapb.ContaineranalysisAlphaNoteBuildSignatureKeyTypeEnum {
	if e == nil {
		return alphapb.ContaineranalysisAlphaNoteBuildSignatureKeyTypeEnum(0)
	}
	if v, ok := alphapb.ContaineranalysisAlphaNoteBuildSignatureKeyTypeEnum_value["NoteBuildSignatureKeyTypeEnum"+string(*e)]; ok {
		return alphapb.ContaineranalysisAlphaNoteBuildSignatureKeyTypeEnum(v)
	}
	return alphapb.ContaineranalysisAlphaNoteBuildSignatureKeyTypeEnum(0)
}

// NotePackageDistributionArchitectureEnumToProto converts a NotePackageDistributionArchitectureEnum enum to its proto representation.
func ContaineranalysisAlphaNotePackageDistributionArchitectureEnumToProto(e *alpha.NotePackageDistributionArchitectureEnum) alphapb.ContaineranalysisAlphaNotePackageDistributionArchitectureEnum {
	if e == nil {
		return alphapb.ContaineranalysisAlphaNotePackageDistributionArchitectureEnum(0)
	}
	if v, ok := alphapb.ContaineranalysisAlphaNotePackageDistributionArchitectureEnum_value["NotePackageDistributionArchitectureEnum"+string(*e)]; ok {
		return alphapb.ContaineranalysisAlphaNotePackageDistributionArchitectureEnum(v)
	}
	return alphapb.ContaineranalysisAlphaNotePackageDistributionArchitectureEnum(0)
}

// NotePackageDistributionLatestVersionKindEnumToProto converts a NotePackageDistributionLatestVersionKindEnum enum to its proto representation.
func ContaineranalysisAlphaNotePackageDistributionLatestVersionKindEnumToProto(e *alpha.NotePackageDistributionLatestVersionKindEnum) alphapb.ContaineranalysisAlphaNotePackageDistributionLatestVersionKindEnum {
	if e == nil {
		return alphapb.ContaineranalysisAlphaNotePackageDistributionLatestVersionKindEnum(0)
	}
	if v, ok := alphapb.ContaineranalysisAlphaNotePackageDistributionLatestVersionKindEnum_value["NotePackageDistributionLatestVersionKindEnum"+string(*e)]; ok {
		return alphapb.ContaineranalysisAlphaNotePackageDistributionLatestVersionKindEnum(v)
	}
	return alphapb.ContaineranalysisAlphaNotePackageDistributionLatestVersionKindEnum(0)
}

// NoteDiscoveryAnalysisKindEnumToProto converts a NoteDiscoveryAnalysisKindEnum enum to its proto representation.
func ContaineranalysisAlphaNoteDiscoveryAnalysisKindEnumToProto(e *alpha.NoteDiscoveryAnalysisKindEnum) alphapb.ContaineranalysisAlphaNoteDiscoveryAnalysisKindEnum {
	if e == nil {
		return alphapb.ContaineranalysisAlphaNoteDiscoveryAnalysisKindEnum(0)
	}
	if v, ok := alphapb.ContaineranalysisAlphaNoteDiscoveryAnalysisKindEnum_value["NoteDiscoveryAnalysisKindEnum"+string(*e)]; ok {
		return alphapb.ContaineranalysisAlphaNoteDiscoveryAnalysisKindEnum(v)
	}
	return alphapb.ContaineranalysisAlphaNoteDiscoveryAnalysisKindEnum(0)
}

// NoteRelatedUrlToProto converts a NoteRelatedUrl object to its proto representation.
func ContaineranalysisAlphaNoteRelatedUrlToProto(o *alpha.NoteRelatedUrl) *alphapb.ContaineranalysisAlphaNoteRelatedUrl {
	if o == nil {
		return nil
	}
	p := &alphapb.ContaineranalysisAlphaNoteRelatedUrl{}
	p.SetUrl(dcl.ValueOrEmptyString(o.Url))
	p.SetLabel(dcl.ValueOrEmptyString(o.Label))
	return p
}

// NoteVulnerabilityToProto converts a NoteVulnerability object to its proto representation.
func ContaineranalysisAlphaNoteVulnerabilityToProto(o *alpha.NoteVulnerability) *alphapb.ContaineranalysisAlphaNoteVulnerability {
	if o == nil {
		return nil
	}
	p := &alphapb.ContaineranalysisAlphaNoteVulnerability{}
	p.SetCvssScore(dcl.ValueOrEmptyDouble(o.CvssScore))
	p.SetSeverity(ContaineranalysisAlphaNoteVulnerabilitySeverityEnumToProto(o.Severity))
	p.SetCvssV3(ContaineranalysisAlphaNoteVulnerabilityCvssV3ToProto(o.CvssV3))
	p.SetSourceUpdateTime(dcl.ValueOrEmptyString(o.SourceUpdateTime))
	sDetails := make([]*alphapb.ContaineranalysisAlphaNoteVulnerabilityDetails, len(o.Details))
	for i, r := range o.Details {
		sDetails[i] = ContaineranalysisAlphaNoteVulnerabilityDetailsToProto(&r)
	}
	p.SetDetails(sDetails)
	sWindowsDetails := make([]*alphapb.ContaineranalysisAlphaNoteVulnerabilityWindowsDetails, len(o.WindowsDetails))
	for i, r := range o.WindowsDetails {
		sWindowsDetails[i] = ContaineranalysisAlphaNoteVulnerabilityWindowsDetailsToProto(&r)
	}
	p.SetWindowsDetails(sWindowsDetails)
	return p
}

// NoteVulnerabilityDetailsToProto converts a NoteVulnerabilityDetails object to its proto representation.
func ContaineranalysisAlphaNoteVulnerabilityDetailsToProto(o *alpha.NoteVulnerabilityDetails) *alphapb.ContaineranalysisAlphaNoteVulnerabilityDetails {
	if o == nil {
		return nil
	}
	p := &alphapb.ContaineranalysisAlphaNoteVulnerabilityDetails{}
	p.SetSeverityName(dcl.ValueOrEmptyString(o.SeverityName))
	p.SetDescription(dcl.ValueOrEmptyString(o.Description))
	p.SetPackageType(dcl.ValueOrEmptyString(o.PackageType))
	p.SetAffectedCpeUri(dcl.ValueOrEmptyString(o.AffectedCpeUri))
	p.SetAffectedPackage(dcl.ValueOrEmptyString(o.AffectedPackage))
	p.SetAffectedVersionStart(ContaineranalysisAlphaNoteVulnerabilityDetailsAffectedVersionStartToProto(o.AffectedVersionStart))
	p.SetAffectedVersionEnd(ContaineranalysisAlphaNoteVulnerabilityDetailsAffectedVersionEndToProto(o.AffectedVersionEnd))
	p.SetFixedCpeUri(dcl.ValueOrEmptyString(o.FixedCpeUri))
	p.SetFixedPackage(dcl.ValueOrEmptyString(o.FixedPackage))
	p.SetFixedVersion(ContaineranalysisAlphaNoteVulnerabilityDetailsFixedVersionToProto(o.FixedVersion))
	p.SetIsObsolete(dcl.ValueOrEmptyBool(o.IsObsolete))
	p.SetSourceUpdateTime(dcl.ValueOrEmptyString(o.SourceUpdateTime))
	return p
}

// NoteVulnerabilityDetailsAffectedVersionStartToProto converts a NoteVulnerabilityDetailsAffectedVersionStart object to its proto representation.
func ContaineranalysisAlphaNoteVulnerabilityDetailsAffectedVersionStartToProto(o *alpha.NoteVulnerabilityDetailsAffectedVersionStart) *alphapb.ContaineranalysisAlphaNoteVulnerabilityDetailsAffectedVersionStart {
	if o == nil {
		return nil
	}
	p := &alphapb.ContaineranalysisAlphaNoteVulnerabilityDetailsAffectedVersionStart{}
	p.SetEpoch(dcl.ValueOrEmptyInt64(o.Epoch))
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetRevision(dcl.ValueOrEmptyString(o.Revision))
	p.SetKind(ContaineranalysisAlphaNoteVulnerabilityDetailsAffectedVersionStartKindEnumToProto(o.Kind))
	p.SetFullName(dcl.ValueOrEmptyString(o.FullName))
	return p
}

// NoteVulnerabilityDetailsAffectedVersionEndToProto converts a NoteVulnerabilityDetailsAffectedVersionEnd object to its proto representation.
func ContaineranalysisAlphaNoteVulnerabilityDetailsAffectedVersionEndToProto(o *alpha.NoteVulnerabilityDetailsAffectedVersionEnd) *alphapb.ContaineranalysisAlphaNoteVulnerabilityDetailsAffectedVersionEnd {
	if o == nil {
		return nil
	}
	p := &alphapb.ContaineranalysisAlphaNoteVulnerabilityDetailsAffectedVersionEnd{}
	p.SetEpoch(dcl.ValueOrEmptyInt64(o.Epoch))
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetRevision(dcl.ValueOrEmptyString(o.Revision))
	p.SetKind(ContaineranalysisAlphaNoteVulnerabilityDetailsAffectedVersionEndKindEnumToProto(o.Kind))
	p.SetFullName(dcl.ValueOrEmptyString(o.FullName))
	return p
}

// NoteVulnerabilityDetailsFixedVersionToProto converts a NoteVulnerabilityDetailsFixedVersion object to its proto representation.
func ContaineranalysisAlphaNoteVulnerabilityDetailsFixedVersionToProto(o *alpha.NoteVulnerabilityDetailsFixedVersion) *alphapb.ContaineranalysisAlphaNoteVulnerabilityDetailsFixedVersion {
	if o == nil {
		return nil
	}
	p := &alphapb.ContaineranalysisAlphaNoteVulnerabilityDetailsFixedVersion{}
	p.SetEpoch(dcl.ValueOrEmptyInt64(o.Epoch))
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetRevision(dcl.ValueOrEmptyString(o.Revision))
	p.SetKind(ContaineranalysisAlphaNoteVulnerabilityDetailsFixedVersionKindEnumToProto(o.Kind))
	p.SetFullName(dcl.ValueOrEmptyString(o.FullName))
	return p
}

// NoteVulnerabilityCvssV3ToProto converts a NoteVulnerabilityCvssV3 object to its proto representation.
func ContaineranalysisAlphaNoteVulnerabilityCvssV3ToProto(o *alpha.NoteVulnerabilityCvssV3) *alphapb.ContaineranalysisAlphaNoteVulnerabilityCvssV3 {
	if o == nil {
		return nil
	}
	p := &alphapb.ContaineranalysisAlphaNoteVulnerabilityCvssV3{}
	p.SetBaseScore(dcl.ValueOrEmptyDouble(o.BaseScore))
	p.SetExploitabilityScore(dcl.ValueOrEmptyDouble(o.ExploitabilityScore))
	p.SetImpactScore(dcl.ValueOrEmptyDouble(o.ImpactScore))
	p.SetAttackVector(ContaineranalysisAlphaNoteVulnerabilityCvssV3AttackVectorEnumToProto(o.AttackVector))
	p.SetAttackComplexity(ContaineranalysisAlphaNoteVulnerabilityCvssV3AttackComplexityEnumToProto(o.AttackComplexity))
	p.SetPrivilegesRequired(ContaineranalysisAlphaNoteVulnerabilityCvssV3PrivilegesRequiredEnumToProto(o.PrivilegesRequired))
	p.SetUserInteraction(ContaineranalysisAlphaNoteVulnerabilityCvssV3UserInteractionEnumToProto(o.UserInteraction))
	p.SetScope(ContaineranalysisAlphaNoteVulnerabilityCvssV3ScopeEnumToProto(o.Scope))
	p.SetConfidentialityImpact(ContaineranalysisAlphaNoteVulnerabilityCvssV3ConfidentialityImpactEnumToProto(o.ConfidentialityImpact))
	p.SetIntegrityImpact(ContaineranalysisAlphaNoteVulnerabilityCvssV3IntegrityImpactEnumToProto(o.IntegrityImpact))
	p.SetAvailabilityImpact(ContaineranalysisAlphaNoteVulnerabilityCvssV3AvailabilityImpactEnumToProto(o.AvailabilityImpact))
	return p
}

// NoteVulnerabilityWindowsDetailsToProto converts a NoteVulnerabilityWindowsDetails object to its proto representation.
func ContaineranalysisAlphaNoteVulnerabilityWindowsDetailsToProto(o *alpha.NoteVulnerabilityWindowsDetails) *alphapb.ContaineranalysisAlphaNoteVulnerabilityWindowsDetails {
	if o == nil {
		return nil
	}
	p := &alphapb.ContaineranalysisAlphaNoteVulnerabilityWindowsDetails{}
	p.SetCpeUri(dcl.ValueOrEmptyString(o.CpeUri))
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetDescription(dcl.ValueOrEmptyString(o.Description))
	sFixingKbs := make([]*alphapb.ContaineranalysisAlphaNoteVulnerabilityWindowsDetailsFixingKbs, len(o.FixingKbs))
	for i, r := range o.FixingKbs {
		sFixingKbs[i] = ContaineranalysisAlphaNoteVulnerabilityWindowsDetailsFixingKbsToProto(&r)
	}
	p.SetFixingKbs(sFixingKbs)
	return p
}

// NoteVulnerabilityWindowsDetailsFixingKbsToProto converts a NoteVulnerabilityWindowsDetailsFixingKbs object to its proto representation.
func ContaineranalysisAlphaNoteVulnerabilityWindowsDetailsFixingKbsToProto(o *alpha.NoteVulnerabilityWindowsDetailsFixingKbs) *alphapb.ContaineranalysisAlphaNoteVulnerabilityWindowsDetailsFixingKbs {
	if o == nil {
		return nil
	}
	p := &alphapb.ContaineranalysisAlphaNoteVulnerabilityWindowsDetailsFixingKbs{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetUrl(dcl.ValueOrEmptyString(o.Url))
	return p
}

// NoteBuildToProto converts a NoteBuild object to its proto representation.
func ContaineranalysisAlphaNoteBuildToProto(o *alpha.NoteBuild) *alphapb.ContaineranalysisAlphaNoteBuild {
	if o == nil {
		return nil
	}
	p := &alphapb.ContaineranalysisAlphaNoteBuild{}
	p.SetBuilderVersion(dcl.ValueOrEmptyString(o.BuilderVersion))
	p.SetSignature(ContaineranalysisAlphaNoteBuildSignatureToProto(o.Signature))
	return p
}

// NoteBuildSignatureToProto converts a NoteBuildSignature object to its proto representation.
func ContaineranalysisAlphaNoteBuildSignatureToProto(o *alpha.NoteBuildSignature) *alphapb.ContaineranalysisAlphaNoteBuildSignature {
	if o == nil {
		return nil
	}
	p := &alphapb.ContaineranalysisAlphaNoteBuildSignature{}
	p.SetPublicKey(dcl.ValueOrEmptyString(o.PublicKey))
	p.SetSignature(dcl.ValueOrEmptyString(o.Signature))
	p.SetKeyId(dcl.ValueOrEmptyString(o.KeyId))
	p.SetKeyType(ContaineranalysisAlphaNoteBuildSignatureKeyTypeEnumToProto(o.KeyType))
	return p
}

// NoteImageToProto converts a NoteImage object to its proto representation.
func ContaineranalysisAlphaNoteImageToProto(o *alpha.NoteImage) *alphapb.ContaineranalysisAlphaNoteImage {
	if o == nil {
		return nil
	}
	p := &alphapb.ContaineranalysisAlphaNoteImage{}
	p.SetResourceUrl(dcl.ValueOrEmptyString(o.ResourceUrl))
	p.SetFingerprint(ContaineranalysisAlphaNoteImageFingerprintToProto(o.Fingerprint))
	return p
}

// NoteImageFingerprintToProto converts a NoteImageFingerprint object to its proto representation.
func ContaineranalysisAlphaNoteImageFingerprintToProto(o *alpha.NoteImageFingerprint) *alphapb.ContaineranalysisAlphaNoteImageFingerprint {
	if o == nil {
		return nil
	}
	p := &alphapb.ContaineranalysisAlphaNoteImageFingerprint{}
	p.SetV1Name(dcl.ValueOrEmptyString(o.V1Name))
	p.SetV2Name(dcl.ValueOrEmptyString(o.V2Name))
	sV2Blob := make([]string, len(o.V2Blob))
	for i, r := range o.V2Blob {
		sV2Blob[i] = r
	}
	p.SetV2Blob(sV2Blob)
	return p
}

// NotePackageToProto converts a NotePackage object to its proto representation.
func ContaineranalysisAlphaNotePackageToProto(o *alpha.NotePackage) *alphapb.ContaineranalysisAlphaNotePackage {
	if o == nil {
		return nil
	}
	p := &alphapb.ContaineranalysisAlphaNotePackage{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	sDistribution := make([]*alphapb.ContaineranalysisAlphaNotePackageDistribution, len(o.Distribution))
	for i, r := range o.Distribution {
		sDistribution[i] = ContaineranalysisAlphaNotePackageDistributionToProto(&r)
	}
	p.SetDistribution(sDistribution)
	return p
}

// NotePackageDistributionToProto converts a NotePackageDistribution object to its proto representation.
func ContaineranalysisAlphaNotePackageDistributionToProto(o *alpha.NotePackageDistribution) *alphapb.ContaineranalysisAlphaNotePackageDistribution {
	if o == nil {
		return nil
	}
	p := &alphapb.ContaineranalysisAlphaNotePackageDistribution{}
	p.SetCpeUri(dcl.ValueOrEmptyString(o.CpeUri))
	p.SetArchitecture(ContaineranalysisAlphaNotePackageDistributionArchitectureEnumToProto(o.Architecture))
	p.SetLatestVersion(ContaineranalysisAlphaNotePackageDistributionLatestVersionToProto(o.LatestVersion))
	p.SetMaintainer(dcl.ValueOrEmptyString(o.Maintainer))
	p.SetUrl(dcl.ValueOrEmptyString(o.Url))
	p.SetDescription(dcl.ValueOrEmptyString(o.Description))
	return p
}

// NotePackageDistributionLatestVersionToProto converts a NotePackageDistributionLatestVersion object to its proto representation.
func ContaineranalysisAlphaNotePackageDistributionLatestVersionToProto(o *alpha.NotePackageDistributionLatestVersion) *alphapb.ContaineranalysisAlphaNotePackageDistributionLatestVersion {
	if o == nil {
		return nil
	}
	p := &alphapb.ContaineranalysisAlphaNotePackageDistributionLatestVersion{}
	p.SetEpoch(dcl.ValueOrEmptyInt64(o.Epoch))
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetRevision(dcl.ValueOrEmptyString(o.Revision))
	p.SetKind(ContaineranalysisAlphaNotePackageDistributionLatestVersionKindEnumToProto(o.Kind))
	p.SetFullName(dcl.ValueOrEmptyString(o.FullName))
	return p
}

// NoteDiscoveryToProto converts a NoteDiscovery object to its proto representation.
func ContaineranalysisAlphaNoteDiscoveryToProto(o *alpha.NoteDiscovery) *alphapb.ContaineranalysisAlphaNoteDiscovery {
	if o == nil {
		return nil
	}
	p := &alphapb.ContaineranalysisAlphaNoteDiscovery{}
	p.SetAnalysisKind(ContaineranalysisAlphaNoteDiscoveryAnalysisKindEnumToProto(o.AnalysisKind))
	return p
}

// NoteDeploymentToProto converts a NoteDeployment object to its proto representation.
func ContaineranalysisAlphaNoteDeploymentToProto(o *alpha.NoteDeployment) *alphapb.ContaineranalysisAlphaNoteDeployment {
	if o == nil {
		return nil
	}
	p := &alphapb.ContaineranalysisAlphaNoteDeployment{}
	sResourceUri := make([]string, len(o.ResourceUri))
	for i, r := range o.ResourceUri {
		sResourceUri[i] = r
	}
	p.SetResourceUri(sResourceUri)
	return p
}

// NoteAttestationToProto converts a NoteAttestation object to its proto representation.
func ContaineranalysisAlphaNoteAttestationToProto(o *alpha.NoteAttestation) *alphapb.ContaineranalysisAlphaNoteAttestation {
	if o == nil {
		return nil
	}
	p := &alphapb.ContaineranalysisAlphaNoteAttestation{}
	p.SetHint(ContaineranalysisAlphaNoteAttestationHintToProto(o.Hint))
	return p
}

// NoteAttestationHintToProto converts a NoteAttestationHint object to its proto representation.
func ContaineranalysisAlphaNoteAttestationHintToProto(o *alpha.NoteAttestationHint) *alphapb.ContaineranalysisAlphaNoteAttestationHint {
	if o == nil {
		return nil
	}
	p := &alphapb.ContaineranalysisAlphaNoteAttestationHint{}
	p.SetHumanReadableName(dcl.ValueOrEmptyString(o.HumanReadableName))
	return p
}

// NoteToProto converts a Note resource to its proto representation.
func NoteToProto(resource *alpha.Note) *alphapb.ContaineranalysisAlphaNote {
	p := &alphapb.ContaineranalysisAlphaNote{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetShortDescription(dcl.ValueOrEmptyString(resource.ShortDescription))
	p.SetLongDescription(dcl.ValueOrEmptyString(resource.LongDescription))
	p.SetExpirationTime(dcl.ValueOrEmptyString(resource.ExpirationTime))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetVulnerability(ContaineranalysisAlphaNoteVulnerabilityToProto(resource.Vulnerability))
	p.SetBuild_(ContaineranalysisAlphaNoteBuildToProto(resource.Build))
	p.SetImage(ContaineranalysisAlphaNoteImageToProto(resource.Image))
	p.SetPackage(ContaineranalysisAlphaNotePackageToProto(resource.Package))
	p.SetDiscovery(ContaineranalysisAlphaNoteDiscoveryToProto(resource.Discovery))
	p.SetDeployment(ContaineranalysisAlphaNoteDeploymentToProto(resource.Deployment))
	p.SetAttestation(ContaineranalysisAlphaNoteAttestationToProto(resource.Attestation))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	sRelatedUrl := make([]*alphapb.ContaineranalysisAlphaNoteRelatedUrl, len(resource.RelatedUrl))
	for i, r := range resource.RelatedUrl {
		sRelatedUrl[i] = ContaineranalysisAlphaNoteRelatedUrlToProto(&r)
	}
	p.SetRelatedUrl(sRelatedUrl)
	sRelatedNoteNames := make([]string, len(resource.RelatedNoteNames))
	for i, r := range resource.RelatedNoteNames {
		sRelatedNoteNames[i] = r
	}
	p.SetRelatedNoteNames(sRelatedNoteNames)

	return p
}

// applyNote handles the gRPC request by passing it to the underlying Note Apply() method.
func (s *NoteServer) applyNote(ctx context.Context, c *alpha.Client, request *alphapb.ApplyContaineranalysisAlphaNoteRequest) (*alphapb.ContaineranalysisAlphaNote, error) {
	p := ProtoToNote(request.GetResource())
	res, err := c.ApplyNote(ctx, p)
	if err != nil {
		return nil, err
	}
	r := NoteToProto(res)
	return r, nil
}

// applyContaineranalysisAlphaNote handles the gRPC request by passing it to the underlying Note Apply() method.
func (s *NoteServer) ApplyContaineranalysisAlphaNote(ctx context.Context, request *alphapb.ApplyContaineranalysisAlphaNoteRequest) (*alphapb.ContaineranalysisAlphaNote, error) {
	cl, err := createConfigNote(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyNote(ctx, cl, request)
}

// DeleteNote handles the gRPC request by passing it to the underlying Note Delete() method.
func (s *NoteServer) DeleteContaineranalysisAlphaNote(ctx context.Context, request *alphapb.DeleteContaineranalysisAlphaNoteRequest) (*emptypb.Empty, error) {

	cl, err := createConfigNote(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteNote(ctx, ProtoToNote(request.GetResource()))

}

// ListContaineranalysisAlphaNote handles the gRPC request by passing it to the underlying NoteList() method.
func (s *NoteServer) ListContaineranalysisAlphaNote(ctx context.Context, request *alphapb.ListContaineranalysisAlphaNoteRequest) (*alphapb.ListContaineranalysisAlphaNoteResponse, error) {
	cl, err := createConfigNote(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListNote(ctx, request.GetProject())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.ContaineranalysisAlphaNote
	for _, r := range resources.Items {
		rp := NoteToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListContaineranalysisAlphaNoteResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigNote(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
