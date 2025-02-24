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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/containeranalysis/beta/containeranalysis_beta_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/containeranalysis/beta"
)

// NoteServer implements the gRPC interface for Note.
type NoteServer struct{}

// ProtoToNoteVulnerabilitySeverityEnum converts a NoteVulnerabilitySeverityEnum enum from its proto representation.
func ProtoToContaineranalysisBetaNoteVulnerabilitySeverityEnum(e betapb.ContaineranalysisBetaNoteVulnerabilitySeverityEnum) *beta.NoteVulnerabilitySeverityEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ContaineranalysisBetaNoteVulnerabilitySeverityEnum_name[int32(e)]; ok {
		e := beta.NoteVulnerabilitySeverityEnum(n[len("ContaineranalysisBetaNoteVulnerabilitySeverityEnum"):])
		return &e
	}
	return nil
}

// ProtoToNoteVulnerabilityDetailsAffectedVersionStartKindEnum converts a NoteVulnerabilityDetailsAffectedVersionStartKindEnum enum from its proto representation.
func ProtoToContaineranalysisBetaNoteVulnerabilityDetailsAffectedVersionStartKindEnum(e betapb.ContaineranalysisBetaNoteVulnerabilityDetailsAffectedVersionStartKindEnum) *beta.NoteVulnerabilityDetailsAffectedVersionStartKindEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ContaineranalysisBetaNoteVulnerabilityDetailsAffectedVersionStartKindEnum_name[int32(e)]; ok {
		e := beta.NoteVulnerabilityDetailsAffectedVersionStartKindEnum(n[len("ContaineranalysisBetaNoteVulnerabilityDetailsAffectedVersionStartKindEnum"):])
		return &e
	}
	return nil
}

// ProtoToNoteVulnerabilityDetailsAffectedVersionEndKindEnum converts a NoteVulnerabilityDetailsAffectedVersionEndKindEnum enum from its proto representation.
func ProtoToContaineranalysisBetaNoteVulnerabilityDetailsAffectedVersionEndKindEnum(e betapb.ContaineranalysisBetaNoteVulnerabilityDetailsAffectedVersionEndKindEnum) *beta.NoteVulnerabilityDetailsAffectedVersionEndKindEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ContaineranalysisBetaNoteVulnerabilityDetailsAffectedVersionEndKindEnum_name[int32(e)]; ok {
		e := beta.NoteVulnerabilityDetailsAffectedVersionEndKindEnum(n[len("ContaineranalysisBetaNoteVulnerabilityDetailsAffectedVersionEndKindEnum"):])
		return &e
	}
	return nil
}

// ProtoToNoteVulnerabilityDetailsFixedVersionKindEnum converts a NoteVulnerabilityDetailsFixedVersionKindEnum enum from its proto representation.
func ProtoToContaineranalysisBetaNoteVulnerabilityDetailsFixedVersionKindEnum(e betapb.ContaineranalysisBetaNoteVulnerabilityDetailsFixedVersionKindEnum) *beta.NoteVulnerabilityDetailsFixedVersionKindEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ContaineranalysisBetaNoteVulnerabilityDetailsFixedVersionKindEnum_name[int32(e)]; ok {
		e := beta.NoteVulnerabilityDetailsFixedVersionKindEnum(n[len("ContaineranalysisBetaNoteVulnerabilityDetailsFixedVersionKindEnum"):])
		return &e
	}
	return nil
}

// ProtoToNoteVulnerabilityCvssV3AttackVectorEnum converts a NoteVulnerabilityCvssV3AttackVectorEnum enum from its proto representation.
func ProtoToContaineranalysisBetaNoteVulnerabilityCvssV3AttackVectorEnum(e betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3AttackVectorEnum) *beta.NoteVulnerabilityCvssV3AttackVectorEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3AttackVectorEnum_name[int32(e)]; ok {
		e := beta.NoteVulnerabilityCvssV3AttackVectorEnum(n[len("ContaineranalysisBetaNoteVulnerabilityCvssV3AttackVectorEnum"):])
		return &e
	}
	return nil
}

// ProtoToNoteVulnerabilityCvssV3AttackComplexityEnum converts a NoteVulnerabilityCvssV3AttackComplexityEnum enum from its proto representation.
func ProtoToContaineranalysisBetaNoteVulnerabilityCvssV3AttackComplexityEnum(e betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3AttackComplexityEnum) *beta.NoteVulnerabilityCvssV3AttackComplexityEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3AttackComplexityEnum_name[int32(e)]; ok {
		e := beta.NoteVulnerabilityCvssV3AttackComplexityEnum(n[len("ContaineranalysisBetaNoteVulnerabilityCvssV3AttackComplexityEnum"):])
		return &e
	}
	return nil
}

// ProtoToNoteVulnerabilityCvssV3PrivilegesRequiredEnum converts a NoteVulnerabilityCvssV3PrivilegesRequiredEnum enum from its proto representation.
func ProtoToContaineranalysisBetaNoteVulnerabilityCvssV3PrivilegesRequiredEnum(e betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3PrivilegesRequiredEnum) *beta.NoteVulnerabilityCvssV3PrivilegesRequiredEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3PrivilegesRequiredEnum_name[int32(e)]; ok {
		e := beta.NoteVulnerabilityCvssV3PrivilegesRequiredEnum(n[len("ContaineranalysisBetaNoteVulnerabilityCvssV3PrivilegesRequiredEnum"):])
		return &e
	}
	return nil
}

// ProtoToNoteVulnerabilityCvssV3UserInteractionEnum converts a NoteVulnerabilityCvssV3UserInteractionEnum enum from its proto representation.
func ProtoToContaineranalysisBetaNoteVulnerabilityCvssV3UserInteractionEnum(e betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3UserInteractionEnum) *beta.NoteVulnerabilityCvssV3UserInteractionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3UserInteractionEnum_name[int32(e)]; ok {
		e := beta.NoteVulnerabilityCvssV3UserInteractionEnum(n[len("ContaineranalysisBetaNoteVulnerabilityCvssV3UserInteractionEnum"):])
		return &e
	}
	return nil
}

// ProtoToNoteVulnerabilityCvssV3ScopeEnum converts a NoteVulnerabilityCvssV3ScopeEnum enum from its proto representation.
func ProtoToContaineranalysisBetaNoteVulnerabilityCvssV3ScopeEnum(e betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3ScopeEnum) *beta.NoteVulnerabilityCvssV3ScopeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3ScopeEnum_name[int32(e)]; ok {
		e := beta.NoteVulnerabilityCvssV3ScopeEnum(n[len("ContaineranalysisBetaNoteVulnerabilityCvssV3ScopeEnum"):])
		return &e
	}
	return nil
}

// ProtoToNoteVulnerabilityCvssV3ConfidentialityImpactEnum converts a NoteVulnerabilityCvssV3ConfidentialityImpactEnum enum from its proto representation.
func ProtoToContaineranalysisBetaNoteVulnerabilityCvssV3ConfidentialityImpactEnum(e betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3ConfidentialityImpactEnum) *beta.NoteVulnerabilityCvssV3ConfidentialityImpactEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3ConfidentialityImpactEnum_name[int32(e)]; ok {
		e := beta.NoteVulnerabilityCvssV3ConfidentialityImpactEnum(n[len("ContaineranalysisBetaNoteVulnerabilityCvssV3ConfidentialityImpactEnum"):])
		return &e
	}
	return nil
}

// ProtoToNoteVulnerabilityCvssV3IntegrityImpactEnum converts a NoteVulnerabilityCvssV3IntegrityImpactEnum enum from its proto representation.
func ProtoToContaineranalysisBetaNoteVulnerabilityCvssV3IntegrityImpactEnum(e betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3IntegrityImpactEnum) *beta.NoteVulnerabilityCvssV3IntegrityImpactEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3IntegrityImpactEnum_name[int32(e)]; ok {
		e := beta.NoteVulnerabilityCvssV3IntegrityImpactEnum(n[len("ContaineranalysisBetaNoteVulnerabilityCvssV3IntegrityImpactEnum"):])
		return &e
	}
	return nil
}

// ProtoToNoteVulnerabilityCvssV3AvailabilityImpactEnum converts a NoteVulnerabilityCvssV3AvailabilityImpactEnum enum from its proto representation.
func ProtoToContaineranalysisBetaNoteVulnerabilityCvssV3AvailabilityImpactEnum(e betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3AvailabilityImpactEnum) *beta.NoteVulnerabilityCvssV3AvailabilityImpactEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3AvailabilityImpactEnum_name[int32(e)]; ok {
		e := beta.NoteVulnerabilityCvssV3AvailabilityImpactEnum(n[len("ContaineranalysisBetaNoteVulnerabilityCvssV3AvailabilityImpactEnum"):])
		return &e
	}
	return nil
}

// ProtoToNoteBuildSignatureKeyTypeEnum converts a NoteBuildSignatureKeyTypeEnum enum from its proto representation.
func ProtoToContaineranalysisBetaNoteBuildSignatureKeyTypeEnum(e betapb.ContaineranalysisBetaNoteBuildSignatureKeyTypeEnum) *beta.NoteBuildSignatureKeyTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ContaineranalysisBetaNoteBuildSignatureKeyTypeEnum_name[int32(e)]; ok {
		e := beta.NoteBuildSignatureKeyTypeEnum(n[len("ContaineranalysisBetaNoteBuildSignatureKeyTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToNotePackageDistributionArchitectureEnum converts a NotePackageDistributionArchitectureEnum enum from its proto representation.
func ProtoToContaineranalysisBetaNotePackageDistributionArchitectureEnum(e betapb.ContaineranalysisBetaNotePackageDistributionArchitectureEnum) *beta.NotePackageDistributionArchitectureEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ContaineranalysisBetaNotePackageDistributionArchitectureEnum_name[int32(e)]; ok {
		e := beta.NotePackageDistributionArchitectureEnum(n[len("ContaineranalysisBetaNotePackageDistributionArchitectureEnum"):])
		return &e
	}
	return nil
}

// ProtoToNotePackageDistributionLatestVersionKindEnum converts a NotePackageDistributionLatestVersionKindEnum enum from its proto representation.
func ProtoToContaineranalysisBetaNotePackageDistributionLatestVersionKindEnum(e betapb.ContaineranalysisBetaNotePackageDistributionLatestVersionKindEnum) *beta.NotePackageDistributionLatestVersionKindEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ContaineranalysisBetaNotePackageDistributionLatestVersionKindEnum_name[int32(e)]; ok {
		e := beta.NotePackageDistributionLatestVersionKindEnum(n[len("ContaineranalysisBetaNotePackageDistributionLatestVersionKindEnum"):])
		return &e
	}
	return nil
}

// ProtoToNoteDiscoveryAnalysisKindEnum converts a NoteDiscoveryAnalysisKindEnum enum from its proto representation.
func ProtoToContaineranalysisBetaNoteDiscoveryAnalysisKindEnum(e betapb.ContaineranalysisBetaNoteDiscoveryAnalysisKindEnum) *beta.NoteDiscoveryAnalysisKindEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ContaineranalysisBetaNoteDiscoveryAnalysisKindEnum_name[int32(e)]; ok {
		e := beta.NoteDiscoveryAnalysisKindEnum(n[len("ContaineranalysisBetaNoteDiscoveryAnalysisKindEnum"):])
		return &e
	}
	return nil
}

// ProtoToNoteRelatedUrl converts a NoteRelatedUrl object from its proto representation.
func ProtoToContaineranalysisBetaNoteRelatedUrl(p *betapb.ContaineranalysisBetaNoteRelatedUrl) *beta.NoteRelatedUrl {
	if p == nil {
		return nil
	}
	obj := &beta.NoteRelatedUrl{
		Url:   dcl.StringOrNil(p.GetUrl()),
		Label: dcl.StringOrNil(p.GetLabel()),
	}
	return obj
}

// ProtoToNoteVulnerability converts a NoteVulnerability object from its proto representation.
func ProtoToContaineranalysisBetaNoteVulnerability(p *betapb.ContaineranalysisBetaNoteVulnerability) *beta.NoteVulnerability {
	if p == nil {
		return nil
	}
	obj := &beta.NoteVulnerability{
		CvssScore:        dcl.Float64OrNil(p.GetCvssScore()),
		Severity:         ProtoToContaineranalysisBetaNoteVulnerabilitySeverityEnum(p.GetSeverity()),
		CvssV3:           ProtoToContaineranalysisBetaNoteVulnerabilityCvssV3(p.GetCvssV3()),
		SourceUpdateTime: dcl.StringOrNil(p.GetSourceUpdateTime()),
	}
	for _, r := range p.GetDetails() {
		obj.Details = append(obj.Details, *ProtoToContaineranalysisBetaNoteVulnerabilityDetails(r))
	}
	for _, r := range p.GetWindowsDetails() {
		obj.WindowsDetails = append(obj.WindowsDetails, *ProtoToContaineranalysisBetaNoteVulnerabilityWindowsDetails(r))
	}
	return obj
}

// ProtoToNoteVulnerabilityDetails converts a NoteVulnerabilityDetails object from its proto representation.
func ProtoToContaineranalysisBetaNoteVulnerabilityDetails(p *betapb.ContaineranalysisBetaNoteVulnerabilityDetails) *beta.NoteVulnerabilityDetails {
	if p == nil {
		return nil
	}
	obj := &beta.NoteVulnerabilityDetails{
		SeverityName:         dcl.StringOrNil(p.GetSeverityName()),
		Description:          dcl.StringOrNil(p.GetDescription()),
		PackageType:          dcl.StringOrNil(p.GetPackageType()),
		AffectedCpeUri:       dcl.StringOrNil(p.GetAffectedCpeUri()),
		AffectedPackage:      dcl.StringOrNil(p.GetAffectedPackage()),
		AffectedVersionStart: ProtoToContaineranalysisBetaNoteVulnerabilityDetailsAffectedVersionStart(p.GetAffectedVersionStart()),
		AffectedVersionEnd:   ProtoToContaineranalysisBetaNoteVulnerabilityDetailsAffectedVersionEnd(p.GetAffectedVersionEnd()),
		FixedCpeUri:          dcl.StringOrNil(p.GetFixedCpeUri()),
		FixedPackage:         dcl.StringOrNil(p.GetFixedPackage()),
		FixedVersion:         ProtoToContaineranalysisBetaNoteVulnerabilityDetailsFixedVersion(p.GetFixedVersion()),
		IsObsolete:           dcl.Bool(p.GetIsObsolete()),
		SourceUpdateTime:     dcl.StringOrNil(p.GetSourceUpdateTime()),
	}
	return obj
}

// ProtoToNoteVulnerabilityDetailsAffectedVersionStart converts a NoteVulnerabilityDetailsAffectedVersionStart object from its proto representation.
func ProtoToContaineranalysisBetaNoteVulnerabilityDetailsAffectedVersionStart(p *betapb.ContaineranalysisBetaNoteVulnerabilityDetailsAffectedVersionStart) *beta.NoteVulnerabilityDetailsAffectedVersionStart {
	if p == nil {
		return nil
	}
	obj := &beta.NoteVulnerabilityDetailsAffectedVersionStart{
		Epoch:    dcl.Int64OrNil(p.GetEpoch()),
		Name:     dcl.StringOrNil(p.GetName()),
		Revision: dcl.StringOrNil(p.GetRevision()),
		Kind:     ProtoToContaineranalysisBetaNoteVulnerabilityDetailsAffectedVersionStartKindEnum(p.GetKind()),
		FullName: dcl.StringOrNil(p.GetFullName()),
	}
	return obj
}

// ProtoToNoteVulnerabilityDetailsAffectedVersionEnd converts a NoteVulnerabilityDetailsAffectedVersionEnd object from its proto representation.
func ProtoToContaineranalysisBetaNoteVulnerabilityDetailsAffectedVersionEnd(p *betapb.ContaineranalysisBetaNoteVulnerabilityDetailsAffectedVersionEnd) *beta.NoteVulnerabilityDetailsAffectedVersionEnd {
	if p == nil {
		return nil
	}
	obj := &beta.NoteVulnerabilityDetailsAffectedVersionEnd{
		Epoch:    dcl.Int64OrNil(p.GetEpoch()),
		Name:     dcl.StringOrNil(p.GetName()),
		Revision: dcl.StringOrNil(p.GetRevision()),
		Kind:     ProtoToContaineranalysisBetaNoteVulnerabilityDetailsAffectedVersionEndKindEnum(p.GetKind()),
		FullName: dcl.StringOrNil(p.GetFullName()),
	}
	return obj
}

// ProtoToNoteVulnerabilityDetailsFixedVersion converts a NoteVulnerabilityDetailsFixedVersion object from its proto representation.
func ProtoToContaineranalysisBetaNoteVulnerabilityDetailsFixedVersion(p *betapb.ContaineranalysisBetaNoteVulnerabilityDetailsFixedVersion) *beta.NoteVulnerabilityDetailsFixedVersion {
	if p == nil {
		return nil
	}
	obj := &beta.NoteVulnerabilityDetailsFixedVersion{
		Epoch:    dcl.Int64OrNil(p.GetEpoch()),
		Name:     dcl.StringOrNil(p.GetName()),
		Revision: dcl.StringOrNil(p.GetRevision()),
		Kind:     ProtoToContaineranalysisBetaNoteVulnerabilityDetailsFixedVersionKindEnum(p.GetKind()),
		FullName: dcl.StringOrNil(p.GetFullName()),
	}
	return obj
}

// ProtoToNoteVulnerabilityCvssV3 converts a NoteVulnerabilityCvssV3 object from its proto representation.
func ProtoToContaineranalysisBetaNoteVulnerabilityCvssV3(p *betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3) *beta.NoteVulnerabilityCvssV3 {
	if p == nil {
		return nil
	}
	obj := &beta.NoteVulnerabilityCvssV3{
		BaseScore:             dcl.Float64OrNil(p.GetBaseScore()),
		ExploitabilityScore:   dcl.Float64OrNil(p.GetExploitabilityScore()),
		ImpactScore:           dcl.Float64OrNil(p.GetImpactScore()),
		AttackVector:          ProtoToContaineranalysisBetaNoteVulnerabilityCvssV3AttackVectorEnum(p.GetAttackVector()),
		AttackComplexity:      ProtoToContaineranalysisBetaNoteVulnerabilityCvssV3AttackComplexityEnum(p.GetAttackComplexity()),
		PrivilegesRequired:    ProtoToContaineranalysisBetaNoteVulnerabilityCvssV3PrivilegesRequiredEnum(p.GetPrivilegesRequired()),
		UserInteraction:       ProtoToContaineranalysisBetaNoteVulnerabilityCvssV3UserInteractionEnum(p.GetUserInteraction()),
		Scope:                 ProtoToContaineranalysisBetaNoteVulnerabilityCvssV3ScopeEnum(p.GetScope()),
		ConfidentialityImpact: ProtoToContaineranalysisBetaNoteVulnerabilityCvssV3ConfidentialityImpactEnum(p.GetConfidentialityImpact()),
		IntegrityImpact:       ProtoToContaineranalysisBetaNoteVulnerabilityCvssV3IntegrityImpactEnum(p.GetIntegrityImpact()),
		AvailabilityImpact:    ProtoToContaineranalysisBetaNoteVulnerabilityCvssV3AvailabilityImpactEnum(p.GetAvailabilityImpact()),
	}
	return obj
}

// ProtoToNoteVulnerabilityWindowsDetails converts a NoteVulnerabilityWindowsDetails object from its proto representation.
func ProtoToContaineranalysisBetaNoteVulnerabilityWindowsDetails(p *betapb.ContaineranalysisBetaNoteVulnerabilityWindowsDetails) *beta.NoteVulnerabilityWindowsDetails {
	if p == nil {
		return nil
	}
	obj := &beta.NoteVulnerabilityWindowsDetails{
		CpeUri:      dcl.StringOrNil(p.GetCpeUri()),
		Name:        dcl.StringOrNil(p.GetName()),
		Description: dcl.StringOrNil(p.GetDescription()),
	}
	for _, r := range p.GetFixingKbs() {
		obj.FixingKbs = append(obj.FixingKbs, *ProtoToContaineranalysisBetaNoteVulnerabilityWindowsDetailsFixingKbs(r))
	}
	return obj
}

// ProtoToNoteVulnerabilityWindowsDetailsFixingKbs converts a NoteVulnerabilityWindowsDetailsFixingKbs object from its proto representation.
func ProtoToContaineranalysisBetaNoteVulnerabilityWindowsDetailsFixingKbs(p *betapb.ContaineranalysisBetaNoteVulnerabilityWindowsDetailsFixingKbs) *beta.NoteVulnerabilityWindowsDetailsFixingKbs {
	if p == nil {
		return nil
	}
	obj := &beta.NoteVulnerabilityWindowsDetailsFixingKbs{
		Name: dcl.StringOrNil(p.GetName()),
		Url:  dcl.StringOrNil(p.GetUrl()),
	}
	return obj
}

// ProtoToNoteBuild converts a NoteBuild object from its proto representation.
func ProtoToContaineranalysisBetaNoteBuild(p *betapb.ContaineranalysisBetaNoteBuild) *beta.NoteBuild {
	if p == nil {
		return nil
	}
	obj := &beta.NoteBuild{
		BuilderVersion: dcl.StringOrNil(p.GetBuilderVersion()),
		Signature:      ProtoToContaineranalysisBetaNoteBuildSignature(p.GetSignature()),
	}
	return obj
}

// ProtoToNoteBuildSignature converts a NoteBuildSignature object from its proto representation.
func ProtoToContaineranalysisBetaNoteBuildSignature(p *betapb.ContaineranalysisBetaNoteBuildSignature) *beta.NoteBuildSignature {
	if p == nil {
		return nil
	}
	obj := &beta.NoteBuildSignature{
		PublicKey: dcl.StringOrNil(p.GetPublicKey()),
		Signature: dcl.StringOrNil(p.GetSignature()),
		KeyId:     dcl.StringOrNil(p.GetKeyId()),
		KeyType:   ProtoToContaineranalysisBetaNoteBuildSignatureKeyTypeEnum(p.GetKeyType()),
	}
	return obj
}

// ProtoToNoteImage converts a NoteImage object from its proto representation.
func ProtoToContaineranalysisBetaNoteImage(p *betapb.ContaineranalysisBetaNoteImage) *beta.NoteImage {
	if p == nil {
		return nil
	}
	obj := &beta.NoteImage{
		ResourceUrl: dcl.StringOrNil(p.GetResourceUrl()),
		Fingerprint: ProtoToContaineranalysisBetaNoteImageFingerprint(p.GetFingerprint()),
	}
	return obj
}

// ProtoToNoteImageFingerprint converts a NoteImageFingerprint object from its proto representation.
func ProtoToContaineranalysisBetaNoteImageFingerprint(p *betapb.ContaineranalysisBetaNoteImageFingerprint) *beta.NoteImageFingerprint {
	if p == nil {
		return nil
	}
	obj := &beta.NoteImageFingerprint{
		V1Name: dcl.StringOrNil(p.GetV1Name()),
		V2Name: dcl.StringOrNil(p.GetV2Name()),
	}
	for _, r := range p.GetV2Blob() {
		obj.V2Blob = append(obj.V2Blob, r)
	}
	return obj
}

// ProtoToNotePackage converts a NotePackage object from its proto representation.
func ProtoToContaineranalysisBetaNotePackage(p *betapb.ContaineranalysisBetaNotePackage) *beta.NotePackage {
	if p == nil {
		return nil
	}
	obj := &beta.NotePackage{
		Name: dcl.StringOrNil(p.GetName()),
	}
	for _, r := range p.GetDistribution() {
		obj.Distribution = append(obj.Distribution, *ProtoToContaineranalysisBetaNotePackageDistribution(r))
	}
	return obj
}

// ProtoToNotePackageDistribution converts a NotePackageDistribution object from its proto representation.
func ProtoToContaineranalysisBetaNotePackageDistribution(p *betapb.ContaineranalysisBetaNotePackageDistribution) *beta.NotePackageDistribution {
	if p == nil {
		return nil
	}
	obj := &beta.NotePackageDistribution{
		CpeUri:        dcl.StringOrNil(p.GetCpeUri()),
		Architecture:  ProtoToContaineranalysisBetaNotePackageDistributionArchitectureEnum(p.GetArchitecture()),
		LatestVersion: ProtoToContaineranalysisBetaNotePackageDistributionLatestVersion(p.GetLatestVersion()),
		Maintainer:    dcl.StringOrNil(p.GetMaintainer()),
		Url:           dcl.StringOrNil(p.GetUrl()),
		Description:   dcl.StringOrNil(p.GetDescription()),
	}
	return obj
}

// ProtoToNotePackageDistributionLatestVersion converts a NotePackageDistributionLatestVersion object from its proto representation.
func ProtoToContaineranalysisBetaNotePackageDistributionLatestVersion(p *betapb.ContaineranalysisBetaNotePackageDistributionLatestVersion) *beta.NotePackageDistributionLatestVersion {
	if p == nil {
		return nil
	}
	obj := &beta.NotePackageDistributionLatestVersion{
		Epoch:    dcl.Int64OrNil(p.GetEpoch()),
		Name:     dcl.StringOrNil(p.GetName()),
		Revision: dcl.StringOrNil(p.GetRevision()),
		Kind:     ProtoToContaineranalysisBetaNotePackageDistributionLatestVersionKindEnum(p.GetKind()),
		FullName: dcl.StringOrNil(p.GetFullName()),
	}
	return obj
}

// ProtoToNoteDiscovery converts a NoteDiscovery object from its proto representation.
func ProtoToContaineranalysisBetaNoteDiscovery(p *betapb.ContaineranalysisBetaNoteDiscovery) *beta.NoteDiscovery {
	if p == nil {
		return nil
	}
	obj := &beta.NoteDiscovery{
		AnalysisKind: ProtoToContaineranalysisBetaNoteDiscoveryAnalysisKindEnum(p.GetAnalysisKind()),
	}
	return obj
}

// ProtoToNoteDeployment converts a NoteDeployment object from its proto representation.
func ProtoToContaineranalysisBetaNoteDeployment(p *betapb.ContaineranalysisBetaNoteDeployment) *beta.NoteDeployment {
	if p == nil {
		return nil
	}
	obj := &beta.NoteDeployment{}
	for _, r := range p.GetResourceUri() {
		obj.ResourceUri = append(obj.ResourceUri, r)
	}
	return obj
}

// ProtoToNoteAttestation converts a NoteAttestation object from its proto representation.
func ProtoToContaineranalysisBetaNoteAttestation(p *betapb.ContaineranalysisBetaNoteAttestation) *beta.NoteAttestation {
	if p == nil {
		return nil
	}
	obj := &beta.NoteAttestation{
		Hint: ProtoToContaineranalysisBetaNoteAttestationHint(p.GetHint()),
	}
	return obj
}

// ProtoToNoteAttestationHint converts a NoteAttestationHint object from its proto representation.
func ProtoToContaineranalysisBetaNoteAttestationHint(p *betapb.ContaineranalysisBetaNoteAttestationHint) *beta.NoteAttestationHint {
	if p == nil {
		return nil
	}
	obj := &beta.NoteAttestationHint{
		HumanReadableName: dcl.StringOrNil(p.GetHumanReadableName()),
	}
	return obj
}

// ProtoToNote converts a Note resource from its proto representation.
func ProtoToNote(p *betapb.ContaineranalysisBetaNote) *beta.Note {
	obj := &beta.Note{
		Name:             dcl.StringOrNil(p.GetName()),
		ShortDescription: dcl.StringOrNil(p.GetShortDescription()),
		LongDescription:  dcl.StringOrNil(p.GetLongDescription()),
		ExpirationTime:   dcl.StringOrNil(p.GetExpirationTime()),
		CreateTime:       dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:       dcl.StringOrNil(p.GetUpdateTime()),
		Vulnerability:    ProtoToContaineranalysisBetaNoteVulnerability(p.GetVulnerability()),
		Build:            ProtoToContaineranalysisBetaNoteBuild(p.GetBuild_()),
		Image:            ProtoToContaineranalysisBetaNoteImage(p.GetImage()),
		Package:          ProtoToContaineranalysisBetaNotePackage(p.GetPackage()),
		Discovery:        ProtoToContaineranalysisBetaNoteDiscovery(p.GetDiscovery()),
		Deployment:       ProtoToContaineranalysisBetaNoteDeployment(p.GetDeployment()),
		Attestation:      ProtoToContaineranalysisBetaNoteAttestation(p.GetAttestation()),
		Project:          dcl.StringOrNil(p.GetProject()),
	}
	for _, r := range p.GetRelatedUrl() {
		obj.RelatedUrl = append(obj.RelatedUrl, *ProtoToContaineranalysisBetaNoteRelatedUrl(r))
	}
	for _, r := range p.GetRelatedNoteNames() {
		obj.RelatedNoteNames = append(obj.RelatedNoteNames, r)
	}
	return obj
}

// NoteVulnerabilitySeverityEnumToProto converts a NoteVulnerabilitySeverityEnum enum to its proto representation.
func ContaineranalysisBetaNoteVulnerabilitySeverityEnumToProto(e *beta.NoteVulnerabilitySeverityEnum) betapb.ContaineranalysisBetaNoteVulnerabilitySeverityEnum {
	if e == nil {
		return betapb.ContaineranalysisBetaNoteVulnerabilitySeverityEnum(0)
	}
	if v, ok := betapb.ContaineranalysisBetaNoteVulnerabilitySeverityEnum_value["NoteVulnerabilitySeverityEnum"+string(*e)]; ok {
		return betapb.ContaineranalysisBetaNoteVulnerabilitySeverityEnum(v)
	}
	return betapb.ContaineranalysisBetaNoteVulnerabilitySeverityEnum(0)
}

// NoteVulnerabilityDetailsAffectedVersionStartKindEnumToProto converts a NoteVulnerabilityDetailsAffectedVersionStartKindEnum enum to its proto representation.
func ContaineranalysisBetaNoteVulnerabilityDetailsAffectedVersionStartKindEnumToProto(e *beta.NoteVulnerabilityDetailsAffectedVersionStartKindEnum) betapb.ContaineranalysisBetaNoteVulnerabilityDetailsAffectedVersionStartKindEnum {
	if e == nil {
		return betapb.ContaineranalysisBetaNoteVulnerabilityDetailsAffectedVersionStartKindEnum(0)
	}
	if v, ok := betapb.ContaineranalysisBetaNoteVulnerabilityDetailsAffectedVersionStartKindEnum_value["NoteVulnerabilityDetailsAffectedVersionStartKindEnum"+string(*e)]; ok {
		return betapb.ContaineranalysisBetaNoteVulnerabilityDetailsAffectedVersionStartKindEnum(v)
	}
	return betapb.ContaineranalysisBetaNoteVulnerabilityDetailsAffectedVersionStartKindEnum(0)
}

// NoteVulnerabilityDetailsAffectedVersionEndKindEnumToProto converts a NoteVulnerabilityDetailsAffectedVersionEndKindEnum enum to its proto representation.
func ContaineranalysisBetaNoteVulnerabilityDetailsAffectedVersionEndKindEnumToProto(e *beta.NoteVulnerabilityDetailsAffectedVersionEndKindEnum) betapb.ContaineranalysisBetaNoteVulnerabilityDetailsAffectedVersionEndKindEnum {
	if e == nil {
		return betapb.ContaineranalysisBetaNoteVulnerabilityDetailsAffectedVersionEndKindEnum(0)
	}
	if v, ok := betapb.ContaineranalysisBetaNoteVulnerabilityDetailsAffectedVersionEndKindEnum_value["NoteVulnerabilityDetailsAffectedVersionEndKindEnum"+string(*e)]; ok {
		return betapb.ContaineranalysisBetaNoteVulnerabilityDetailsAffectedVersionEndKindEnum(v)
	}
	return betapb.ContaineranalysisBetaNoteVulnerabilityDetailsAffectedVersionEndKindEnum(0)
}

// NoteVulnerabilityDetailsFixedVersionKindEnumToProto converts a NoteVulnerabilityDetailsFixedVersionKindEnum enum to its proto representation.
func ContaineranalysisBetaNoteVulnerabilityDetailsFixedVersionKindEnumToProto(e *beta.NoteVulnerabilityDetailsFixedVersionKindEnum) betapb.ContaineranalysisBetaNoteVulnerabilityDetailsFixedVersionKindEnum {
	if e == nil {
		return betapb.ContaineranalysisBetaNoteVulnerabilityDetailsFixedVersionKindEnum(0)
	}
	if v, ok := betapb.ContaineranalysisBetaNoteVulnerabilityDetailsFixedVersionKindEnum_value["NoteVulnerabilityDetailsFixedVersionKindEnum"+string(*e)]; ok {
		return betapb.ContaineranalysisBetaNoteVulnerabilityDetailsFixedVersionKindEnum(v)
	}
	return betapb.ContaineranalysisBetaNoteVulnerabilityDetailsFixedVersionKindEnum(0)
}

// NoteVulnerabilityCvssV3AttackVectorEnumToProto converts a NoteVulnerabilityCvssV3AttackVectorEnum enum to its proto representation.
func ContaineranalysisBetaNoteVulnerabilityCvssV3AttackVectorEnumToProto(e *beta.NoteVulnerabilityCvssV3AttackVectorEnum) betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3AttackVectorEnum {
	if e == nil {
		return betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3AttackVectorEnum(0)
	}
	if v, ok := betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3AttackVectorEnum_value["NoteVulnerabilityCvssV3AttackVectorEnum"+string(*e)]; ok {
		return betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3AttackVectorEnum(v)
	}
	return betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3AttackVectorEnum(0)
}

// NoteVulnerabilityCvssV3AttackComplexityEnumToProto converts a NoteVulnerabilityCvssV3AttackComplexityEnum enum to its proto representation.
func ContaineranalysisBetaNoteVulnerabilityCvssV3AttackComplexityEnumToProto(e *beta.NoteVulnerabilityCvssV3AttackComplexityEnum) betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3AttackComplexityEnum {
	if e == nil {
		return betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3AttackComplexityEnum(0)
	}
	if v, ok := betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3AttackComplexityEnum_value["NoteVulnerabilityCvssV3AttackComplexityEnum"+string(*e)]; ok {
		return betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3AttackComplexityEnum(v)
	}
	return betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3AttackComplexityEnum(0)
}

// NoteVulnerabilityCvssV3PrivilegesRequiredEnumToProto converts a NoteVulnerabilityCvssV3PrivilegesRequiredEnum enum to its proto representation.
func ContaineranalysisBetaNoteVulnerabilityCvssV3PrivilegesRequiredEnumToProto(e *beta.NoteVulnerabilityCvssV3PrivilegesRequiredEnum) betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3PrivilegesRequiredEnum {
	if e == nil {
		return betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3PrivilegesRequiredEnum(0)
	}
	if v, ok := betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3PrivilegesRequiredEnum_value["NoteVulnerabilityCvssV3PrivilegesRequiredEnum"+string(*e)]; ok {
		return betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3PrivilegesRequiredEnum(v)
	}
	return betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3PrivilegesRequiredEnum(0)
}

// NoteVulnerabilityCvssV3UserInteractionEnumToProto converts a NoteVulnerabilityCvssV3UserInteractionEnum enum to its proto representation.
func ContaineranalysisBetaNoteVulnerabilityCvssV3UserInteractionEnumToProto(e *beta.NoteVulnerabilityCvssV3UserInteractionEnum) betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3UserInteractionEnum {
	if e == nil {
		return betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3UserInteractionEnum(0)
	}
	if v, ok := betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3UserInteractionEnum_value["NoteVulnerabilityCvssV3UserInteractionEnum"+string(*e)]; ok {
		return betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3UserInteractionEnum(v)
	}
	return betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3UserInteractionEnum(0)
}

// NoteVulnerabilityCvssV3ScopeEnumToProto converts a NoteVulnerabilityCvssV3ScopeEnum enum to its proto representation.
func ContaineranalysisBetaNoteVulnerabilityCvssV3ScopeEnumToProto(e *beta.NoteVulnerabilityCvssV3ScopeEnum) betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3ScopeEnum {
	if e == nil {
		return betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3ScopeEnum(0)
	}
	if v, ok := betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3ScopeEnum_value["NoteVulnerabilityCvssV3ScopeEnum"+string(*e)]; ok {
		return betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3ScopeEnum(v)
	}
	return betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3ScopeEnum(0)
}

// NoteVulnerabilityCvssV3ConfidentialityImpactEnumToProto converts a NoteVulnerabilityCvssV3ConfidentialityImpactEnum enum to its proto representation.
func ContaineranalysisBetaNoteVulnerabilityCvssV3ConfidentialityImpactEnumToProto(e *beta.NoteVulnerabilityCvssV3ConfidentialityImpactEnum) betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3ConfidentialityImpactEnum {
	if e == nil {
		return betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3ConfidentialityImpactEnum(0)
	}
	if v, ok := betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3ConfidentialityImpactEnum_value["NoteVulnerabilityCvssV3ConfidentialityImpactEnum"+string(*e)]; ok {
		return betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3ConfidentialityImpactEnum(v)
	}
	return betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3ConfidentialityImpactEnum(0)
}

// NoteVulnerabilityCvssV3IntegrityImpactEnumToProto converts a NoteVulnerabilityCvssV3IntegrityImpactEnum enum to its proto representation.
func ContaineranalysisBetaNoteVulnerabilityCvssV3IntegrityImpactEnumToProto(e *beta.NoteVulnerabilityCvssV3IntegrityImpactEnum) betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3IntegrityImpactEnum {
	if e == nil {
		return betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3IntegrityImpactEnum(0)
	}
	if v, ok := betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3IntegrityImpactEnum_value["NoteVulnerabilityCvssV3IntegrityImpactEnum"+string(*e)]; ok {
		return betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3IntegrityImpactEnum(v)
	}
	return betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3IntegrityImpactEnum(0)
}

// NoteVulnerabilityCvssV3AvailabilityImpactEnumToProto converts a NoteVulnerabilityCvssV3AvailabilityImpactEnum enum to its proto representation.
func ContaineranalysisBetaNoteVulnerabilityCvssV3AvailabilityImpactEnumToProto(e *beta.NoteVulnerabilityCvssV3AvailabilityImpactEnum) betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3AvailabilityImpactEnum {
	if e == nil {
		return betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3AvailabilityImpactEnum(0)
	}
	if v, ok := betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3AvailabilityImpactEnum_value["NoteVulnerabilityCvssV3AvailabilityImpactEnum"+string(*e)]; ok {
		return betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3AvailabilityImpactEnum(v)
	}
	return betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3AvailabilityImpactEnum(0)
}

// NoteBuildSignatureKeyTypeEnumToProto converts a NoteBuildSignatureKeyTypeEnum enum to its proto representation.
func ContaineranalysisBetaNoteBuildSignatureKeyTypeEnumToProto(e *beta.NoteBuildSignatureKeyTypeEnum) betapb.ContaineranalysisBetaNoteBuildSignatureKeyTypeEnum {
	if e == nil {
		return betapb.ContaineranalysisBetaNoteBuildSignatureKeyTypeEnum(0)
	}
	if v, ok := betapb.ContaineranalysisBetaNoteBuildSignatureKeyTypeEnum_value["NoteBuildSignatureKeyTypeEnum"+string(*e)]; ok {
		return betapb.ContaineranalysisBetaNoteBuildSignatureKeyTypeEnum(v)
	}
	return betapb.ContaineranalysisBetaNoteBuildSignatureKeyTypeEnum(0)
}

// NotePackageDistributionArchitectureEnumToProto converts a NotePackageDistributionArchitectureEnum enum to its proto representation.
func ContaineranalysisBetaNotePackageDistributionArchitectureEnumToProto(e *beta.NotePackageDistributionArchitectureEnum) betapb.ContaineranalysisBetaNotePackageDistributionArchitectureEnum {
	if e == nil {
		return betapb.ContaineranalysisBetaNotePackageDistributionArchitectureEnum(0)
	}
	if v, ok := betapb.ContaineranalysisBetaNotePackageDistributionArchitectureEnum_value["NotePackageDistributionArchitectureEnum"+string(*e)]; ok {
		return betapb.ContaineranalysisBetaNotePackageDistributionArchitectureEnum(v)
	}
	return betapb.ContaineranalysisBetaNotePackageDistributionArchitectureEnum(0)
}

// NotePackageDistributionLatestVersionKindEnumToProto converts a NotePackageDistributionLatestVersionKindEnum enum to its proto representation.
func ContaineranalysisBetaNotePackageDistributionLatestVersionKindEnumToProto(e *beta.NotePackageDistributionLatestVersionKindEnum) betapb.ContaineranalysisBetaNotePackageDistributionLatestVersionKindEnum {
	if e == nil {
		return betapb.ContaineranalysisBetaNotePackageDistributionLatestVersionKindEnum(0)
	}
	if v, ok := betapb.ContaineranalysisBetaNotePackageDistributionLatestVersionKindEnum_value["NotePackageDistributionLatestVersionKindEnum"+string(*e)]; ok {
		return betapb.ContaineranalysisBetaNotePackageDistributionLatestVersionKindEnum(v)
	}
	return betapb.ContaineranalysisBetaNotePackageDistributionLatestVersionKindEnum(0)
}

// NoteDiscoveryAnalysisKindEnumToProto converts a NoteDiscoveryAnalysisKindEnum enum to its proto representation.
func ContaineranalysisBetaNoteDiscoveryAnalysisKindEnumToProto(e *beta.NoteDiscoveryAnalysisKindEnum) betapb.ContaineranalysisBetaNoteDiscoveryAnalysisKindEnum {
	if e == nil {
		return betapb.ContaineranalysisBetaNoteDiscoveryAnalysisKindEnum(0)
	}
	if v, ok := betapb.ContaineranalysisBetaNoteDiscoveryAnalysisKindEnum_value["NoteDiscoveryAnalysisKindEnum"+string(*e)]; ok {
		return betapb.ContaineranalysisBetaNoteDiscoveryAnalysisKindEnum(v)
	}
	return betapb.ContaineranalysisBetaNoteDiscoveryAnalysisKindEnum(0)
}

// NoteRelatedUrlToProto converts a NoteRelatedUrl object to its proto representation.
func ContaineranalysisBetaNoteRelatedUrlToProto(o *beta.NoteRelatedUrl) *betapb.ContaineranalysisBetaNoteRelatedUrl {
	if o == nil {
		return nil
	}
	p := &betapb.ContaineranalysisBetaNoteRelatedUrl{}
	p.SetUrl(dcl.ValueOrEmptyString(o.Url))
	p.SetLabel(dcl.ValueOrEmptyString(o.Label))
	return p
}

// NoteVulnerabilityToProto converts a NoteVulnerability object to its proto representation.
func ContaineranalysisBetaNoteVulnerabilityToProto(o *beta.NoteVulnerability) *betapb.ContaineranalysisBetaNoteVulnerability {
	if o == nil {
		return nil
	}
	p := &betapb.ContaineranalysisBetaNoteVulnerability{}
	p.SetCvssScore(dcl.ValueOrEmptyDouble(o.CvssScore))
	p.SetSeverity(ContaineranalysisBetaNoteVulnerabilitySeverityEnumToProto(o.Severity))
	p.SetCvssV3(ContaineranalysisBetaNoteVulnerabilityCvssV3ToProto(o.CvssV3))
	p.SetSourceUpdateTime(dcl.ValueOrEmptyString(o.SourceUpdateTime))
	sDetails := make([]*betapb.ContaineranalysisBetaNoteVulnerabilityDetails, len(o.Details))
	for i, r := range o.Details {
		sDetails[i] = ContaineranalysisBetaNoteVulnerabilityDetailsToProto(&r)
	}
	p.SetDetails(sDetails)
	sWindowsDetails := make([]*betapb.ContaineranalysisBetaNoteVulnerabilityWindowsDetails, len(o.WindowsDetails))
	for i, r := range o.WindowsDetails {
		sWindowsDetails[i] = ContaineranalysisBetaNoteVulnerabilityWindowsDetailsToProto(&r)
	}
	p.SetWindowsDetails(sWindowsDetails)
	return p
}

// NoteVulnerabilityDetailsToProto converts a NoteVulnerabilityDetails object to its proto representation.
func ContaineranalysisBetaNoteVulnerabilityDetailsToProto(o *beta.NoteVulnerabilityDetails) *betapb.ContaineranalysisBetaNoteVulnerabilityDetails {
	if o == nil {
		return nil
	}
	p := &betapb.ContaineranalysisBetaNoteVulnerabilityDetails{}
	p.SetSeverityName(dcl.ValueOrEmptyString(o.SeverityName))
	p.SetDescription(dcl.ValueOrEmptyString(o.Description))
	p.SetPackageType(dcl.ValueOrEmptyString(o.PackageType))
	p.SetAffectedCpeUri(dcl.ValueOrEmptyString(o.AffectedCpeUri))
	p.SetAffectedPackage(dcl.ValueOrEmptyString(o.AffectedPackage))
	p.SetAffectedVersionStart(ContaineranalysisBetaNoteVulnerabilityDetailsAffectedVersionStartToProto(o.AffectedVersionStart))
	p.SetAffectedVersionEnd(ContaineranalysisBetaNoteVulnerabilityDetailsAffectedVersionEndToProto(o.AffectedVersionEnd))
	p.SetFixedCpeUri(dcl.ValueOrEmptyString(o.FixedCpeUri))
	p.SetFixedPackage(dcl.ValueOrEmptyString(o.FixedPackage))
	p.SetFixedVersion(ContaineranalysisBetaNoteVulnerabilityDetailsFixedVersionToProto(o.FixedVersion))
	p.SetIsObsolete(dcl.ValueOrEmptyBool(o.IsObsolete))
	p.SetSourceUpdateTime(dcl.ValueOrEmptyString(o.SourceUpdateTime))
	return p
}

// NoteVulnerabilityDetailsAffectedVersionStartToProto converts a NoteVulnerabilityDetailsAffectedVersionStart object to its proto representation.
func ContaineranalysisBetaNoteVulnerabilityDetailsAffectedVersionStartToProto(o *beta.NoteVulnerabilityDetailsAffectedVersionStart) *betapb.ContaineranalysisBetaNoteVulnerabilityDetailsAffectedVersionStart {
	if o == nil {
		return nil
	}
	p := &betapb.ContaineranalysisBetaNoteVulnerabilityDetailsAffectedVersionStart{}
	p.SetEpoch(dcl.ValueOrEmptyInt64(o.Epoch))
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetRevision(dcl.ValueOrEmptyString(o.Revision))
	p.SetKind(ContaineranalysisBetaNoteVulnerabilityDetailsAffectedVersionStartKindEnumToProto(o.Kind))
	p.SetFullName(dcl.ValueOrEmptyString(o.FullName))
	return p
}

// NoteVulnerabilityDetailsAffectedVersionEndToProto converts a NoteVulnerabilityDetailsAffectedVersionEnd object to its proto representation.
func ContaineranalysisBetaNoteVulnerabilityDetailsAffectedVersionEndToProto(o *beta.NoteVulnerabilityDetailsAffectedVersionEnd) *betapb.ContaineranalysisBetaNoteVulnerabilityDetailsAffectedVersionEnd {
	if o == nil {
		return nil
	}
	p := &betapb.ContaineranalysisBetaNoteVulnerabilityDetailsAffectedVersionEnd{}
	p.SetEpoch(dcl.ValueOrEmptyInt64(o.Epoch))
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetRevision(dcl.ValueOrEmptyString(o.Revision))
	p.SetKind(ContaineranalysisBetaNoteVulnerabilityDetailsAffectedVersionEndKindEnumToProto(o.Kind))
	p.SetFullName(dcl.ValueOrEmptyString(o.FullName))
	return p
}

// NoteVulnerabilityDetailsFixedVersionToProto converts a NoteVulnerabilityDetailsFixedVersion object to its proto representation.
func ContaineranalysisBetaNoteVulnerabilityDetailsFixedVersionToProto(o *beta.NoteVulnerabilityDetailsFixedVersion) *betapb.ContaineranalysisBetaNoteVulnerabilityDetailsFixedVersion {
	if o == nil {
		return nil
	}
	p := &betapb.ContaineranalysisBetaNoteVulnerabilityDetailsFixedVersion{}
	p.SetEpoch(dcl.ValueOrEmptyInt64(o.Epoch))
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetRevision(dcl.ValueOrEmptyString(o.Revision))
	p.SetKind(ContaineranalysisBetaNoteVulnerabilityDetailsFixedVersionKindEnumToProto(o.Kind))
	p.SetFullName(dcl.ValueOrEmptyString(o.FullName))
	return p
}

// NoteVulnerabilityCvssV3ToProto converts a NoteVulnerabilityCvssV3 object to its proto representation.
func ContaineranalysisBetaNoteVulnerabilityCvssV3ToProto(o *beta.NoteVulnerabilityCvssV3) *betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3 {
	if o == nil {
		return nil
	}
	p := &betapb.ContaineranalysisBetaNoteVulnerabilityCvssV3{}
	p.SetBaseScore(dcl.ValueOrEmptyDouble(o.BaseScore))
	p.SetExploitabilityScore(dcl.ValueOrEmptyDouble(o.ExploitabilityScore))
	p.SetImpactScore(dcl.ValueOrEmptyDouble(o.ImpactScore))
	p.SetAttackVector(ContaineranalysisBetaNoteVulnerabilityCvssV3AttackVectorEnumToProto(o.AttackVector))
	p.SetAttackComplexity(ContaineranalysisBetaNoteVulnerabilityCvssV3AttackComplexityEnumToProto(o.AttackComplexity))
	p.SetPrivilegesRequired(ContaineranalysisBetaNoteVulnerabilityCvssV3PrivilegesRequiredEnumToProto(o.PrivilegesRequired))
	p.SetUserInteraction(ContaineranalysisBetaNoteVulnerabilityCvssV3UserInteractionEnumToProto(o.UserInteraction))
	p.SetScope(ContaineranalysisBetaNoteVulnerabilityCvssV3ScopeEnumToProto(o.Scope))
	p.SetConfidentialityImpact(ContaineranalysisBetaNoteVulnerabilityCvssV3ConfidentialityImpactEnumToProto(o.ConfidentialityImpact))
	p.SetIntegrityImpact(ContaineranalysisBetaNoteVulnerabilityCvssV3IntegrityImpactEnumToProto(o.IntegrityImpact))
	p.SetAvailabilityImpact(ContaineranalysisBetaNoteVulnerabilityCvssV3AvailabilityImpactEnumToProto(o.AvailabilityImpact))
	return p
}

// NoteVulnerabilityWindowsDetailsToProto converts a NoteVulnerabilityWindowsDetails object to its proto representation.
func ContaineranalysisBetaNoteVulnerabilityWindowsDetailsToProto(o *beta.NoteVulnerabilityWindowsDetails) *betapb.ContaineranalysisBetaNoteVulnerabilityWindowsDetails {
	if o == nil {
		return nil
	}
	p := &betapb.ContaineranalysisBetaNoteVulnerabilityWindowsDetails{}
	p.SetCpeUri(dcl.ValueOrEmptyString(o.CpeUri))
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetDescription(dcl.ValueOrEmptyString(o.Description))
	sFixingKbs := make([]*betapb.ContaineranalysisBetaNoteVulnerabilityWindowsDetailsFixingKbs, len(o.FixingKbs))
	for i, r := range o.FixingKbs {
		sFixingKbs[i] = ContaineranalysisBetaNoteVulnerabilityWindowsDetailsFixingKbsToProto(&r)
	}
	p.SetFixingKbs(sFixingKbs)
	return p
}

// NoteVulnerabilityWindowsDetailsFixingKbsToProto converts a NoteVulnerabilityWindowsDetailsFixingKbs object to its proto representation.
func ContaineranalysisBetaNoteVulnerabilityWindowsDetailsFixingKbsToProto(o *beta.NoteVulnerabilityWindowsDetailsFixingKbs) *betapb.ContaineranalysisBetaNoteVulnerabilityWindowsDetailsFixingKbs {
	if o == nil {
		return nil
	}
	p := &betapb.ContaineranalysisBetaNoteVulnerabilityWindowsDetailsFixingKbs{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetUrl(dcl.ValueOrEmptyString(o.Url))
	return p
}

// NoteBuildToProto converts a NoteBuild object to its proto representation.
func ContaineranalysisBetaNoteBuildToProto(o *beta.NoteBuild) *betapb.ContaineranalysisBetaNoteBuild {
	if o == nil {
		return nil
	}
	p := &betapb.ContaineranalysisBetaNoteBuild{}
	p.SetBuilderVersion(dcl.ValueOrEmptyString(o.BuilderVersion))
	p.SetSignature(ContaineranalysisBetaNoteBuildSignatureToProto(o.Signature))
	return p
}

// NoteBuildSignatureToProto converts a NoteBuildSignature object to its proto representation.
func ContaineranalysisBetaNoteBuildSignatureToProto(o *beta.NoteBuildSignature) *betapb.ContaineranalysisBetaNoteBuildSignature {
	if o == nil {
		return nil
	}
	p := &betapb.ContaineranalysisBetaNoteBuildSignature{}
	p.SetPublicKey(dcl.ValueOrEmptyString(o.PublicKey))
	p.SetSignature(dcl.ValueOrEmptyString(o.Signature))
	p.SetKeyId(dcl.ValueOrEmptyString(o.KeyId))
	p.SetKeyType(ContaineranalysisBetaNoteBuildSignatureKeyTypeEnumToProto(o.KeyType))
	return p
}

// NoteImageToProto converts a NoteImage object to its proto representation.
func ContaineranalysisBetaNoteImageToProto(o *beta.NoteImage) *betapb.ContaineranalysisBetaNoteImage {
	if o == nil {
		return nil
	}
	p := &betapb.ContaineranalysisBetaNoteImage{}
	p.SetResourceUrl(dcl.ValueOrEmptyString(o.ResourceUrl))
	p.SetFingerprint(ContaineranalysisBetaNoteImageFingerprintToProto(o.Fingerprint))
	return p
}

// NoteImageFingerprintToProto converts a NoteImageFingerprint object to its proto representation.
func ContaineranalysisBetaNoteImageFingerprintToProto(o *beta.NoteImageFingerprint) *betapb.ContaineranalysisBetaNoteImageFingerprint {
	if o == nil {
		return nil
	}
	p := &betapb.ContaineranalysisBetaNoteImageFingerprint{}
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
func ContaineranalysisBetaNotePackageToProto(o *beta.NotePackage) *betapb.ContaineranalysisBetaNotePackage {
	if o == nil {
		return nil
	}
	p := &betapb.ContaineranalysisBetaNotePackage{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	sDistribution := make([]*betapb.ContaineranalysisBetaNotePackageDistribution, len(o.Distribution))
	for i, r := range o.Distribution {
		sDistribution[i] = ContaineranalysisBetaNotePackageDistributionToProto(&r)
	}
	p.SetDistribution(sDistribution)
	return p
}

// NotePackageDistributionToProto converts a NotePackageDistribution object to its proto representation.
func ContaineranalysisBetaNotePackageDistributionToProto(o *beta.NotePackageDistribution) *betapb.ContaineranalysisBetaNotePackageDistribution {
	if o == nil {
		return nil
	}
	p := &betapb.ContaineranalysisBetaNotePackageDistribution{}
	p.SetCpeUri(dcl.ValueOrEmptyString(o.CpeUri))
	p.SetArchitecture(ContaineranalysisBetaNotePackageDistributionArchitectureEnumToProto(o.Architecture))
	p.SetLatestVersion(ContaineranalysisBetaNotePackageDistributionLatestVersionToProto(o.LatestVersion))
	p.SetMaintainer(dcl.ValueOrEmptyString(o.Maintainer))
	p.SetUrl(dcl.ValueOrEmptyString(o.Url))
	p.SetDescription(dcl.ValueOrEmptyString(o.Description))
	return p
}

// NotePackageDistributionLatestVersionToProto converts a NotePackageDistributionLatestVersion object to its proto representation.
func ContaineranalysisBetaNotePackageDistributionLatestVersionToProto(o *beta.NotePackageDistributionLatestVersion) *betapb.ContaineranalysisBetaNotePackageDistributionLatestVersion {
	if o == nil {
		return nil
	}
	p := &betapb.ContaineranalysisBetaNotePackageDistributionLatestVersion{}
	p.SetEpoch(dcl.ValueOrEmptyInt64(o.Epoch))
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetRevision(dcl.ValueOrEmptyString(o.Revision))
	p.SetKind(ContaineranalysisBetaNotePackageDistributionLatestVersionKindEnumToProto(o.Kind))
	p.SetFullName(dcl.ValueOrEmptyString(o.FullName))
	return p
}

// NoteDiscoveryToProto converts a NoteDiscovery object to its proto representation.
func ContaineranalysisBetaNoteDiscoveryToProto(o *beta.NoteDiscovery) *betapb.ContaineranalysisBetaNoteDiscovery {
	if o == nil {
		return nil
	}
	p := &betapb.ContaineranalysisBetaNoteDiscovery{}
	p.SetAnalysisKind(ContaineranalysisBetaNoteDiscoveryAnalysisKindEnumToProto(o.AnalysisKind))
	return p
}

// NoteDeploymentToProto converts a NoteDeployment object to its proto representation.
func ContaineranalysisBetaNoteDeploymentToProto(o *beta.NoteDeployment) *betapb.ContaineranalysisBetaNoteDeployment {
	if o == nil {
		return nil
	}
	p := &betapb.ContaineranalysisBetaNoteDeployment{}
	sResourceUri := make([]string, len(o.ResourceUri))
	for i, r := range o.ResourceUri {
		sResourceUri[i] = r
	}
	p.SetResourceUri(sResourceUri)
	return p
}

// NoteAttestationToProto converts a NoteAttestation object to its proto representation.
func ContaineranalysisBetaNoteAttestationToProto(o *beta.NoteAttestation) *betapb.ContaineranalysisBetaNoteAttestation {
	if o == nil {
		return nil
	}
	p := &betapb.ContaineranalysisBetaNoteAttestation{}
	p.SetHint(ContaineranalysisBetaNoteAttestationHintToProto(o.Hint))
	return p
}

// NoteAttestationHintToProto converts a NoteAttestationHint object to its proto representation.
func ContaineranalysisBetaNoteAttestationHintToProto(o *beta.NoteAttestationHint) *betapb.ContaineranalysisBetaNoteAttestationHint {
	if o == nil {
		return nil
	}
	p := &betapb.ContaineranalysisBetaNoteAttestationHint{}
	p.SetHumanReadableName(dcl.ValueOrEmptyString(o.HumanReadableName))
	return p
}

// NoteToProto converts a Note resource to its proto representation.
func NoteToProto(resource *beta.Note) *betapb.ContaineranalysisBetaNote {
	p := &betapb.ContaineranalysisBetaNote{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetShortDescription(dcl.ValueOrEmptyString(resource.ShortDescription))
	p.SetLongDescription(dcl.ValueOrEmptyString(resource.LongDescription))
	p.SetExpirationTime(dcl.ValueOrEmptyString(resource.ExpirationTime))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetVulnerability(ContaineranalysisBetaNoteVulnerabilityToProto(resource.Vulnerability))
	p.SetBuild_(ContaineranalysisBetaNoteBuildToProto(resource.Build))
	p.SetImage(ContaineranalysisBetaNoteImageToProto(resource.Image))
	p.SetPackage(ContaineranalysisBetaNotePackageToProto(resource.Package))
	p.SetDiscovery(ContaineranalysisBetaNoteDiscoveryToProto(resource.Discovery))
	p.SetDeployment(ContaineranalysisBetaNoteDeploymentToProto(resource.Deployment))
	p.SetAttestation(ContaineranalysisBetaNoteAttestationToProto(resource.Attestation))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	sRelatedUrl := make([]*betapb.ContaineranalysisBetaNoteRelatedUrl, len(resource.RelatedUrl))
	for i, r := range resource.RelatedUrl {
		sRelatedUrl[i] = ContaineranalysisBetaNoteRelatedUrlToProto(&r)
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
func (s *NoteServer) applyNote(ctx context.Context, c *beta.Client, request *betapb.ApplyContaineranalysisBetaNoteRequest) (*betapb.ContaineranalysisBetaNote, error) {
	p := ProtoToNote(request.GetResource())
	res, err := c.ApplyNote(ctx, p)
	if err != nil {
		return nil, err
	}
	r := NoteToProto(res)
	return r, nil
}

// applyContaineranalysisBetaNote handles the gRPC request by passing it to the underlying Note Apply() method.
func (s *NoteServer) ApplyContaineranalysisBetaNote(ctx context.Context, request *betapb.ApplyContaineranalysisBetaNoteRequest) (*betapb.ContaineranalysisBetaNote, error) {
	cl, err := createConfigNote(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyNote(ctx, cl, request)
}

// DeleteNote handles the gRPC request by passing it to the underlying Note Delete() method.
func (s *NoteServer) DeleteContaineranalysisBetaNote(ctx context.Context, request *betapb.DeleteContaineranalysisBetaNoteRequest) (*emptypb.Empty, error) {

	cl, err := createConfigNote(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteNote(ctx, ProtoToNote(request.GetResource()))

}

// ListContaineranalysisBetaNote handles the gRPC request by passing it to the underlying NoteList() method.
func (s *NoteServer) ListContaineranalysisBetaNote(ctx context.Context, request *betapb.ListContaineranalysisBetaNoteRequest) (*betapb.ListContaineranalysisBetaNoteResponse, error) {
	cl, err := createConfigNote(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListNote(ctx, request.GetProject())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.ContaineranalysisBetaNote
	for _, r := range resources.Items {
		rp := NoteToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListContaineranalysisBetaNoteResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigNote(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
