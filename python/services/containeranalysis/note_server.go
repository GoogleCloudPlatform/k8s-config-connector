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
	containeranalysispb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/containeranalysis/containeranalysis_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/containeranalysis"
)

// NoteServer implements the gRPC interface for Note.
type NoteServer struct{}

// ProtoToNoteVulnerabilitySeverityEnum converts a NoteVulnerabilitySeverityEnum enum from its proto representation.
func ProtoToContaineranalysisNoteVulnerabilitySeverityEnum(e containeranalysispb.ContaineranalysisNoteVulnerabilitySeverityEnum) *containeranalysis.NoteVulnerabilitySeverityEnum {
	if e == 0 {
		return nil
	}
	if n, ok := containeranalysispb.ContaineranalysisNoteVulnerabilitySeverityEnum_name[int32(e)]; ok {
		e := containeranalysis.NoteVulnerabilitySeverityEnum(n[len("ContaineranalysisNoteVulnerabilitySeverityEnum"):])
		return &e
	}
	return nil
}

// ProtoToNoteVulnerabilityDetailsAffectedVersionStartKindEnum converts a NoteVulnerabilityDetailsAffectedVersionStartKindEnum enum from its proto representation.
func ProtoToContaineranalysisNoteVulnerabilityDetailsAffectedVersionStartKindEnum(e containeranalysispb.ContaineranalysisNoteVulnerabilityDetailsAffectedVersionStartKindEnum) *containeranalysis.NoteVulnerabilityDetailsAffectedVersionStartKindEnum {
	if e == 0 {
		return nil
	}
	if n, ok := containeranalysispb.ContaineranalysisNoteVulnerabilityDetailsAffectedVersionStartKindEnum_name[int32(e)]; ok {
		e := containeranalysis.NoteVulnerabilityDetailsAffectedVersionStartKindEnum(n[len("ContaineranalysisNoteVulnerabilityDetailsAffectedVersionStartKindEnum"):])
		return &e
	}
	return nil
}

// ProtoToNoteVulnerabilityDetailsAffectedVersionEndKindEnum converts a NoteVulnerabilityDetailsAffectedVersionEndKindEnum enum from its proto representation.
func ProtoToContaineranalysisNoteVulnerabilityDetailsAffectedVersionEndKindEnum(e containeranalysispb.ContaineranalysisNoteVulnerabilityDetailsAffectedVersionEndKindEnum) *containeranalysis.NoteVulnerabilityDetailsAffectedVersionEndKindEnum {
	if e == 0 {
		return nil
	}
	if n, ok := containeranalysispb.ContaineranalysisNoteVulnerabilityDetailsAffectedVersionEndKindEnum_name[int32(e)]; ok {
		e := containeranalysis.NoteVulnerabilityDetailsAffectedVersionEndKindEnum(n[len("ContaineranalysisNoteVulnerabilityDetailsAffectedVersionEndKindEnum"):])
		return &e
	}
	return nil
}

// ProtoToNoteVulnerabilityDetailsFixedVersionKindEnum converts a NoteVulnerabilityDetailsFixedVersionKindEnum enum from its proto representation.
func ProtoToContaineranalysisNoteVulnerabilityDetailsFixedVersionKindEnum(e containeranalysispb.ContaineranalysisNoteVulnerabilityDetailsFixedVersionKindEnum) *containeranalysis.NoteVulnerabilityDetailsFixedVersionKindEnum {
	if e == 0 {
		return nil
	}
	if n, ok := containeranalysispb.ContaineranalysisNoteVulnerabilityDetailsFixedVersionKindEnum_name[int32(e)]; ok {
		e := containeranalysis.NoteVulnerabilityDetailsFixedVersionKindEnum(n[len("ContaineranalysisNoteVulnerabilityDetailsFixedVersionKindEnum"):])
		return &e
	}
	return nil
}

// ProtoToNoteVulnerabilityCvssV3AttackVectorEnum converts a NoteVulnerabilityCvssV3AttackVectorEnum enum from its proto representation.
func ProtoToContaineranalysisNoteVulnerabilityCvssV3AttackVectorEnum(e containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3AttackVectorEnum) *containeranalysis.NoteVulnerabilityCvssV3AttackVectorEnum {
	if e == 0 {
		return nil
	}
	if n, ok := containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3AttackVectorEnum_name[int32(e)]; ok {
		e := containeranalysis.NoteVulnerabilityCvssV3AttackVectorEnum(n[len("ContaineranalysisNoteVulnerabilityCvssV3AttackVectorEnum"):])
		return &e
	}
	return nil
}

// ProtoToNoteVulnerabilityCvssV3AttackComplexityEnum converts a NoteVulnerabilityCvssV3AttackComplexityEnum enum from its proto representation.
func ProtoToContaineranalysisNoteVulnerabilityCvssV3AttackComplexityEnum(e containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3AttackComplexityEnum) *containeranalysis.NoteVulnerabilityCvssV3AttackComplexityEnum {
	if e == 0 {
		return nil
	}
	if n, ok := containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3AttackComplexityEnum_name[int32(e)]; ok {
		e := containeranalysis.NoteVulnerabilityCvssV3AttackComplexityEnum(n[len("ContaineranalysisNoteVulnerabilityCvssV3AttackComplexityEnum"):])
		return &e
	}
	return nil
}

// ProtoToNoteVulnerabilityCvssV3PrivilegesRequiredEnum converts a NoteVulnerabilityCvssV3PrivilegesRequiredEnum enum from its proto representation.
func ProtoToContaineranalysisNoteVulnerabilityCvssV3PrivilegesRequiredEnum(e containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3PrivilegesRequiredEnum) *containeranalysis.NoteVulnerabilityCvssV3PrivilegesRequiredEnum {
	if e == 0 {
		return nil
	}
	if n, ok := containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3PrivilegesRequiredEnum_name[int32(e)]; ok {
		e := containeranalysis.NoteVulnerabilityCvssV3PrivilegesRequiredEnum(n[len("ContaineranalysisNoteVulnerabilityCvssV3PrivilegesRequiredEnum"):])
		return &e
	}
	return nil
}

// ProtoToNoteVulnerabilityCvssV3UserInteractionEnum converts a NoteVulnerabilityCvssV3UserInteractionEnum enum from its proto representation.
func ProtoToContaineranalysisNoteVulnerabilityCvssV3UserInteractionEnum(e containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3UserInteractionEnum) *containeranalysis.NoteVulnerabilityCvssV3UserInteractionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3UserInteractionEnum_name[int32(e)]; ok {
		e := containeranalysis.NoteVulnerabilityCvssV3UserInteractionEnum(n[len("ContaineranalysisNoteVulnerabilityCvssV3UserInteractionEnum"):])
		return &e
	}
	return nil
}

// ProtoToNoteVulnerabilityCvssV3ScopeEnum converts a NoteVulnerabilityCvssV3ScopeEnum enum from its proto representation.
func ProtoToContaineranalysisNoteVulnerabilityCvssV3ScopeEnum(e containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3ScopeEnum) *containeranalysis.NoteVulnerabilityCvssV3ScopeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3ScopeEnum_name[int32(e)]; ok {
		e := containeranalysis.NoteVulnerabilityCvssV3ScopeEnum(n[len("ContaineranalysisNoteVulnerabilityCvssV3ScopeEnum"):])
		return &e
	}
	return nil
}

// ProtoToNoteVulnerabilityCvssV3ConfidentialityImpactEnum converts a NoteVulnerabilityCvssV3ConfidentialityImpactEnum enum from its proto representation.
func ProtoToContaineranalysisNoteVulnerabilityCvssV3ConfidentialityImpactEnum(e containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3ConfidentialityImpactEnum) *containeranalysis.NoteVulnerabilityCvssV3ConfidentialityImpactEnum {
	if e == 0 {
		return nil
	}
	if n, ok := containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3ConfidentialityImpactEnum_name[int32(e)]; ok {
		e := containeranalysis.NoteVulnerabilityCvssV3ConfidentialityImpactEnum(n[len("ContaineranalysisNoteVulnerabilityCvssV3ConfidentialityImpactEnum"):])
		return &e
	}
	return nil
}

// ProtoToNoteVulnerabilityCvssV3IntegrityImpactEnum converts a NoteVulnerabilityCvssV3IntegrityImpactEnum enum from its proto representation.
func ProtoToContaineranalysisNoteVulnerabilityCvssV3IntegrityImpactEnum(e containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3IntegrityImpactEnum) *containeranalysis.NoteVulnerabilityCvssV3IntegrityImpactEnum {
	if e == 0 {
		return nil
	}
	if n, ok := containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3IntegrityImpactEnum_name[int32(e)]; ok {
		e := containeranalysis.NoteVulnerabilityCvssV3IntegrityImpactEnum(n[len("ContaineranalysisNoteVulnerabilityCvssV3IntegrityImpactEnum"):])
		return &e
	}
	return nil
}

// ProtoToNoteVulnerabilityCvssV3AvailabilityImpactEnum converts a NoteVulnerabilityCvssV3AvailabilityImpactEnum enum from its proto representation.
func ProtoToContaineranalysisNoteVulnerabilityCvssV3AvailabilityImpactEnum(e containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3AvailabilityImpactEnum) *containeranalysis.NoteVulnerabilityCvssV3AvailabilityImpactEnum {
	if e == 0 {
		return nil
	}
	if n, ok := containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3AvailabilityImpactEnum_name[int32(e)]; ok {
		e := containeranalysis.NoteVulnerabilityCvssV3AvailabilityImpactEnum(n[len("ContaineranalysisNoteVulnerabilityCvssV3AvailabilityImpactEnum"):])
		return &e
	}
	return nil
}

// ProtoToNotePackageDistributionArchitectureEnum converts a NotePackageDistributionArchitectureEnum enum from its proto representation.
func ProtoToContaineranalysisNotePackageDistributionArchitectureEnum(e containeranalysispb.ContaineranalysisNotePackageDistributionArchitectureEnum) *containeranalysis.NotePackageDistributionArchitectureEnum {
	if e == 0 {
		return nil
	}
	if n, ok := containeranalysispb.ContaineranalysisNotePackageDistributionArchitectureEnum_name[int32(e)]; ok {
		e := containeranalysis.NotePackageDistributionArchitectureEnum(n[len("ContaineranalysisNotePackageDistributionArchitectureEnum"):])
		return &e
	}
	return nil
}

// ProtoToNotePackageDistributionLatestVersionKindEnum converts a NotePackageDistributionLatestVersionKindEnum enum from its proto representation.
func ProtoToContaineranalysisNotePackageDistributionLatestVersionKindEnum(e containeranalysispb.ContaineranalysisNotePackageDistributionLatestVersionKindEnum) *containeranalysis.NotePackageDistributionLatestVersionKindEnum {
	if e == 0 {
		return nil
	}
	if n, ok := containeranalysispb.ContaineranalysisNotePackageDistributionLatestVersionKindEnum_name[int32(e)]; ok {
		e := containeranalysis.NotePackageDistributionLatestVersionKindEnum(n[len("ContaineranalysisNotePackageDistributionLatestVersionKindEnum"):])
		return &e
	}
	return nil
}

// ProtoToNoteDiscoveryAnalysisKindEnum converts a NoteDiscoveryAnalysisKindEnum enum from its proto representation.
func ProtoToContaineranalysisNoteDiscoveryAnalysisKindEnum(e containeranalysispb.ContaineranalysisNoteDiscoveryAnalysisKindEnum) *containeranalysis.NoteDiscoveryAnalysisKindEnum {
	if e == 0 {
		return nil
	}
	if n, ok := containeranalysispb.ContaineranalysisNoteDiscoveryAnalysisKindEnum_name[int32(e)]; ok {
		e := containeranalysis.NoteDiscoveryAnalysisKindEnum(n[len("ContaineranalysisNoteDiscoveryAnalysisKindEnum"):])
		return &e
	}
	return nil
}

// ProtoToNoteRelatedUrl converts a NoteRelatedUrl object from its proto representation.
func ProtoToContaineranalysisNoteRelatedUrl(p *containeranalysispb.ContaineranalysisNoteRelatedUrl) *containeranalysis.NoteRelatedUrl {
	if p == nil {
		return nil
	}
	obj := &containeranalysis.NoteRelatedUrl{
		Url:   dcl.StringOrNil(p.GetUrl()),
		Label: dcl.StringOrNil(p.GetLabel()),
	}
	return obj
}

// ProtoToNoteVulnerability converts a NoteVulnerability object from its proto representation.
func ProtoToContaineranalysisNoteVulnerability(p *containeranalysispb.ContaineranalysisNoteVulnerability) *containeranalysis.NoteVulnerability {
	if p == nil {
		return nil
	}
	obj := &containeranalysis.NoteVulnerability{
		CvssScore:        dcl.Float64OrNil(p.GetCvssScore()),
		Severity:         ProtoToContaineranalysisNoteVulnerabilitySeverityEnum(p.GetSeverity()),
		CvssV3:           ProtoToContaineranalysisNoteVulnerabilityCvssV3(p.GetCvssV3()),
		SourceUpdateTime: dcl.StringOrNil(p.GetSourceUpdateTime()),
	}
	for _, r := range p.GetDetails() {
		obj.Details = append(obj.Details, *ProtoToContaineranalysisNoteVulnerabilityDetails(r))
	}
	for _, r := range p.GetWindowsDetails() {
		obj.WindowsDetails = append(obj.WindowsDetails, *ProtoToContaineranalysisNoteVulnerabilityWindowsDetails(r))
	}
	return obj
}

// ProtoToNoteVulnerabilityDetails converts a NoteVulnerabilityDetails object from its proto representation.
func ProtoToContaineranalysisNoteVulnerabilityDetails(p *containeranalysispb.ContaineranalysisNoteVulnerabilityDetails) *containeranalysis.NoteVulnerabilityDetails {
	if p == nil {
		return nil
	}
	obj := &containeranalysis.NoteVulnerabilityDetails{
		SeverityName:         dcl.StringOrNil(p.GetSeverityName()),
		Description:          dcl.StringOrNil(p.GetDescription()),
		PackageType:          dcl.StringOrNil(p.GetPackageType()),
		AffectedCpeUri:       dcl.StringOrNil(p.GetAffectedCpeUri()),
		AffectedPackage:      dcl.StringOrNil(p.GetAffectedPackage()),
		AffectedVersionStart: ProtoToContaineranalysisNoteVulnerabilityDetailsAffectedVersionStart(p.GetAffectedVersionStart()),
		AffectedVersionEnd:   ProtoToContaineranalysisNoteVulnerabilityDetailsAffectedVersionEnd(p.GetAffectedVersionEnd()),
		FixedCpeUri:          dcl.StringOrNil(p.GetFixedCpeUri()),
		FixedPackage:         dcl.StringOrNil(p.GetFixedPackage()),
		FixedVersion:         ProtoToContaineranalysisNoteVulnerabilityDetailsFixedVersion(p.GetFixedVersion()),
		IsObsolete:           dcl.Bool(p.GetIsObsolete()),
		SourceUpdateTime:     dcl.StringOrNil(p.GetSourceUpdateTime()),
	}
	return obj
}

// ProtoToNoteVulnerabilityDetailsAffectedVersionStart converts a NoteVulnerabilityDetailsAffectedVersionStart object from its proto representation.
func ProtoToContaineranalysisNoteVulnerabilityDetailsAffectedVersionStart(p *containeranalysispb.ContaineranalysisNoteVulnerabilityDetailsAffectedVersionStart) *containeranalysis.NoteVulnerabilityDetailsAffectedVersionStart {
	if p == nil {
		return nil
	}
	obj := &containeranalysis.NoteVulnerabilityDetailsAffectedVersionStart{
		Epoch:    dcl.Int64OrNil(p.GetEpoch()),
		Name:     dcl.StringOrNil(p.GetName()),
		Revision: dcl.StringOrNil(p.GetRevision()),
		Kind:     ProtoToContaineranalysisNoteVulnerabilityDetailsAffectedVersionStartKindEnum(p.GetKind()),
		FullName: dcl.StringOrNil(p.GetFullName()),
	}
	return obj
}

// ProtoToNoteVulnerabilityDetailsAffectedVersionEnd converts a NoteVulnerabilityDetailsAffectedVersionEnd object from its proto representation.
func ProtoToContaineranalysisNoteVulnerabilityDetailsAffectedVersionEnd(p *containeranalysispb.ContaineranalysisNoteVulnerabilityDetailsAffectedVersionEnd) *containeranalysis.NoteVulnerabilityDetailsAffectedVersionEnd {
	if p == nil {
		return nil
	}
	obj := &containeranalysis.NoteVulnerabilityDetailsAffectedVersionEnd{
		Epoch:    dcl.Int64OrNil(p.GetEpoch()),
		Name:     dcl.StringOrNil(p.GetName()),
		Revision: dcl.StringOrNil(p.GetRevision()),
		Kind:     ProtoToContaineranalysisNoteVulnerabilityDetailsAffectedVersionEndKindEnum(p.GetKind()),
		FullName: dcl.StringOrNil(p.GetFullName()),
	}
	return obj
}

// ProtoToNoteVulnerabilityDetailsFixedVersion converts a NoteVulnerabilityDetailsFixedVersion object from its proto representation.
func ProtoToContaineranalysisNoteVulnerabilityDetailsFixedVersion(p *containeranalysispb.ContaineranalysisNoteVulnerabilityDetailsFixedVersion) *containeranalysis.NoteVulnerabilityDetailsFixedVersion {
	if p == nil {
		return nil
	}
	obj := &containeranalysis.NoteVulnerabilityDetailsFixedVersion{
		Epoch:    dcl.Int64OrNil(p.GetEpoch()),
		Name:     dcl.StringOrNil(p.GetName()),
		Revision: dcl.StringOrNil(p.GetRevision()),
		Kind:     ProtoToContaineranalysisNoteVulnerabilityDetailsFixedVersionKindEnum(p.GetKind()),
		FullName: dcl.StringOrNil(p.GetFullName()),
	}
	return obj
}

// ProtoToNoteVulnerabilityCvssV3 converts a NoteVulnerabilityCvssV3 object from its proto representation.
func ProtoToContaineranalysisNoteVulnerabilityCvssV3(p *containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3) *containeranalysis.NoteVulnerabilityCvssV3 {
	if p == nil {
		return nil
	}
	obj := &containeranalysis.NoteVulnerabilityCvssV3{
		BaseScore:             dcl.Float64OrNil(p.GetBaseScore()),
		ExploitabilityScore:   dcl.Float64OrNil(p.GetExploitabilityScore()),
		ImpactScore:           dcl.Float64OrNil(p.GetImpactScore()),
		AttackVector:          ProtoToContaineranalysisNoteVulnerabilityCvssV3AttackVectorEnum(p.GetAttackVector()),
		AttackComplexity:      ProtoToContaineranalysisNoteVulnerabilityCvssV3AttackComplexityEnum(p.GetAttackComplexity()),
		PrivilegesRequired:    ProtoToContaineranalysisNoteVulnerabilityCvssV3PrivilegesRequiredEnum(p.GetPrivilegesRequired()),
		UserInteraction:       ProtoToContaineranalysisNoteVulnerabilityCvssV3UserInteractionEnum(p.GetUserInteraction()),
		Scope:                 ProtoToContaineranalysisNoteVulnerabilityCvssV3ScopeEnum(p.GetScope()),
		ConfidentialityImpact: ProtoToContaineranalysisNoteVulnerabilityCvssV3ConfidentialityImpactEnum(p.GetConfidentialityImpact()),
		IntegrityImpact:       ProtoToContaineranalysisNoteVulnerabilityCvssV3IntegrityImpactEnum(p.GetIntegrityImpact()),
		AvailabilityImpact:    ProtoToContaineranalysisNoteVulnerabilityCvssV3AvailabilityImpactEnum(p.GetAvailabilityImpact()),
	}
	return obj
}

// ProtoToNoteVulnerabilityWindowsDetails converts a NoteVulnerabilityWindowsDetails object from its proto representation.
func ProtoToContaineranalysisNoteVulnerabilityWindowsDetails(p *containeranalysispb.ContaineranalysisNoteVulnerabilityWindowsDetails) *containeranalysis.NoteVulnerabilityWindowsDetails {
	if p == nil {
		return nil
	}
	obj := &containeranalysis.NoteVulnerabilityWindowsDetails{
		CpeUri:      dcl.StringOrNil(p.GetCpeUri()),
		Name:        dcl.StringOrNil(p.GetName()),
		Description: dcl.StringOrNil(p.GetDescription()),
	}
	for _, r := range p.GetFixingKbs() {
		obj.FixingKbs = append(obj.FixingKbs, *ProtoToContaineranalysisNoteVulnerabilityWindowsDetailsFixingKbs(r))
	}
	return obj
}

// ProtoToNoteVulnerabilityWindowsDetailsFixingKbs converts a NoteVulnerabilityWindowsDetailsFixingKbs object from its proto representation.
func ProtoToContaineranalysisNoteVulnerabilityWindowsDetailsFixingKbs(p *containeranalysispb.ContaineranalysisNoteVulnerabilityWindowsDetailsFixingKbs) *containeranalysis.NoteVulnerabilityWindowsDetailsFixingKbs {
	if p == nil {
		return nil
	}
	obj := &containeranalysis.NoteVulnerabilityWindowsDetailsFixingKbs{
		Name: dcl.StringOrNil(p.GetName()),
		Url:  dcl.StringOrNil(p.GetUrl()),
	}
	return obj
}

// ProtoToNoteBuild converts a NoteBuild object from its proto representation.
func ProtoToContaineranalysisNoteBuild(p *containeranalysispb.ContaineranalysisNoteBuild) *containeranalysis.NoteBuild {
	if p == nil {
		return nil
	}
	obj := &containeranalysis.NoteBuild{
		BuilderVersion: dcl.StringOrNil(p.GetBuilderVersion()),
	}
	return obj
}

// ProtoToNoteImage converts a NoteImage object from its proto representation.
func ProtoToContaineranalysisNoteImage(p *containeranalysispb.ContaineranalysisNoteImage) *containeranalysis.NoteImage {
	if p == nil {
		return nil
	}
	obj := &containeranalysis.NoteImage{
		ResourceUrl: dcl.StringOrNil(p.GetResourceUrl()),
		Fingerprint: ProtoToContaineranalysisNoteImageFingerprint(p.GetFingerprint()),
	}
	return obj
}

// ProtoToNoteImageFingerprint converts a NoteImageFingerprint object from its proto representation.
func ProtoToContaineranalysisNoteImageFingerprint(p *containeranalysispb.ContaineranalysisNoteImageFingerprint) *containeranalysis.NoteImageFingerprint {
	if p == nil {
		return nil
	}
	obj := &containeranalysis.NoteImageFingerprint{
		V1Name: dcl.StringOrNil(p.GetV1Name()),
		V2Name: dcl.StringOrNil(p.GetV2Name()),
	}
	for _, r := range p.GetV2Blob() {
		obj.V2Blob = append(obj.V2Blob, r)
	}
	return obj
}

// ProtoToNotePackage converts a NotePackage object from its proto representation.
func ProtoToContaineranalysisNotePackage(p *containeranalysispb.ContaineranalysisNotePackage) *containeranalysis.NotePackage {
	if p == nil {
		return nil
	}
	obj := &containeranalysis.NotePackage{
		Name: dcl.StringOrNil(p.GetName()),
	}
	for _, r := range p.GetDistribution() {
		obj.Distribution = append(obj.Distribution, *ProtoToContaineranalysisNotePackageDistribution(r))
	}
	return obj
}

// ProtoToNotePackageDistribution converts a NotePackageDistribution object from its proto representation.
func ProtoToContaineranalysisNotePackageDistribution(p *containeranalysispb.ContaineranalysisNotePackageDistribution) *containeranalysis.NotePackageDistribution {
	if p == nil {
		return nil
	}
	obj := &containeranalysis.NotePackageDistribution{
		CpeUri:        dcl.StringOrNil(p.GetCpeUri()),
		Architecture:  ProtoToContaineranalysisNotePackageDistributionArchitectureEnum(p.GetArchitecture()),
		LatestVersion: ProtoToContaineranalysisNotePackageDistributionLatestVersion(p.GetLatestVersion()),
		Maintainer:    dcl.StringOrNil(p.GetMaintainer()),
		Url:           dcl.StringOrNil(p.GetUrl()),
		Description:   dcl.StringOrNil(p.GetDescription()),
	}
	return obj
}

// ProtoToNotePackageDistributionLatestVersion converts a NotePackageDistributionLatestVersion object from its proto representation.
func ProtoToContaineranalysisNotePackageDistributionLatestVersion(p *containeranalysispb.ContaineranalysisNotePackageDistributionLatestVersion) *containeranalysis.NotePackageDistributionLatestVersion {
	if p == nil {
		return nil
	}
	obj := &containeranalysis.NotePackageDistributionLatestVersion{
		Epoch:    dcl.Int64OrNil(p.GetEpoch()),
		Name:     dcl.StringOrNil(p.GetName()),
		Revision: dcl.StringOrNil(p.GetRevision()),
		Kind:     ProtoToContaineranalysisNotePackageDistributionLatestVersionKindEnum(p.GetKind()),
		FullName: dcl.StringOrNil(p.GetFullName()),
	}
	return obj
}

// ProtoToNoteDiscovery converts a NoteDiscovery object from its proto representation.
func ProtoToContaineranalysisNoteDiscovery(p *containeranalysispb.ContaineranalysisNoteDiscovery) *containeranalysis.NoteDiscovery {
	if p == nil {
		return nil
	}
	obj := &containeranalysis.NoteDiscovery{
		AnalysisKind: ProtoToContaineranalysisNoteDiscoveryAnalysisKindEnum(p.GetAnalysisKind()),
	}
	return obj
}

// ProtoToNoteDeployment converts a NoteDeployment object from its proto representation.
func ProtoToContaineranalysisNoteDeployment(p *containeranalysispb.ContaineranalysisNoteDeployment) *containeranalysis.NoteDeployment {
	if p == nil {
		return nil
	}
	obj := &containeranalysis.NoteDeployment{}
	for _, r := range p.GetResourceUri() {
		obj.ResourceUri = append(obj.ResourceUri, r)
	}
	return obj
}

// ProtoToNoteAttestation converts a NoteAttestation object from its proto representation.
func ProtoToContaineranalysisNoteAttestation(p *containeranalysispb.ContaineranalysisNoteAttestation) *containeranalysis.NoteAttestation {
	if p == nil {
		return nil
	}
	obj := &containeranalysis.NoteAttestation{
		Hint: ProtoToContaineranalysisNoteAttestationHint(p.GetHint()),
	}
	return obj
}

// ProtoToNoteAttestationHint converts a NoteAttestationHint object from its proto representation.
func ProtoToContaineranalysisNoteAttestationHint(p *containeranalysispb.ContaineranalysisNoteAttestationHint) *containeranalysis.NoteAttestationHint {
	if p == nil {
		return nil
	}
	obj := &containeranalysis.NoteAttestationHint{
		HumanReadableName: dcl.StringOrNil(p.GetHumanReadableName()),
	}
	return obj
}

// ProtoToNote converts a Note resource from its proto representation.
func ProtoToNote(p *containeranalysispb.ContaineranalysisNote) *containeranalysis.Note {
	obj := &containeranalysis.Note{
		Name:             dcl.StringOrNil(p.GetName()),
		ShortDescription: dcl.StringOrNil(p.GetShortDescription()),
		LongDescription:  dcl.StringOrNil(p.GetLongDescription()),
		ExpirationTime:   dcl.StringOrNil(p.GetExpirationTime()),
		CreateTime:       dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:       dcl.StringOrNil(p.GetUpdateTime()),
		Vulnerability:    ProtoToContaineranalysisNoteVulnerability(p.GetVulnerability()),
		Build:            ProtoToContaineranalysisNoteBuild(p.GetBuild_()),
		Image:            ProtoToContaineranalysisNoteImage(p.GetImage()),
		Package:          ProtoToContaineranalysisNotePackage(p.GetPackage()),
		Discovery:        ProtoToContaineranalysisNoteDiscovery(p.GetDiscovery()),
		Deployment:       ProtoToContaineranalysisNoteDeployment(p.GetDeployment()),
		Attestation:      ProtoToContaineranalysisNoteAttestation(p.GetAttestation()),
		Project:          dcl.StringOrNil(p.GetProject()),
	}
	for _, r := range p.GetRelatedUrl() {
		obj.RelatedUrl = append(obj.RelatedUrl, *ProtoToContaineranalysisNoteRelatedUrl(r))
	}
	for _, r := range p.GetRelatedNoteNames() {
		obj.RelatedNoteNames = append(obj.RelatedNoteNames, r)
	}
	return obj
}

// NoteVulnerabilitySeverityEnumToProto converts a NoteVulnerabilitySeverityEnum enum to its proto representation.
func ContaineranalysisNoteVulnerabilitySeverityEnumToProto(e *containeranalysis.NoteVulnerabilitySeverityEnum) containeranalysispb.ContaineranalysisNoteVulnerabilitySeverityEnum {
	if e == nil {
		return containeranalysispb.ContaineranalysisNoteVulnerabilitySeverityEnum(0)
	}
	if v, ok := containeranalysispb.ContaineranalysisNoteVulnerabilitySeverityEnum_value["NoteVulnerabilitySeverityEnum"+string(*e)]; ok {
		return containeranalysispb.ContaineranalysisNoteVulnerabilitySeverityEnum(v)
	}
	return containeranalysispb.ContaineranalysisNoteVulnerabilitySeverityEnum(0)
}

// NoteVulnerabilityDetailsAffectedVersionStartKindEnumToProto converts a NoteVulnerabilityDetailsAffectedVersionStartKindEnum enum to its proto representation.
func ContaineranalysisNoteVulnerabilityDetailsAffectedVersionStartKindEnumToProto(e *containeranalysis.NoteVulnerabilityDetailsAffectedVersionStartKindEnum) containeranalysispb.ContaineranalysisNoteVulnerabilityDetailsAffectedVersionStartKindEnum {
	if e == nil {
		return containeranalysispb.ContaineranalysisNoteVulnerabilityDetailsAffectedVersionStartKindEnum(0)
	}
	if v, ok := containeranalysispb.ContaineranalysisNoteVulnerabilityDetailsAffectedVersionStartKindEnum_value["NoteVulnerabilityDetailsAffectedVersionStartKindEnum"+string(*e)]; ok {
		return containeranalysispb.ContaineranalysisNoteVulnerabilityDetailsAffectedVersionStartKindEnum(v)
	}
	return containeranalysispb.ContaineranalysisNoteVulnerabilityDetailsAffectedVersionStartKindEnum(0)
}

// NoteVulnerabilityDetailsAffectedVersionEndKindEnumToProto converts a NoteVulnerabilityDetailsAffectedVersionEndKindEnum enum to its proto representation.
func ContaineranalysisNoteVulnerabilityDetailsAffectedVersionEndKindEnumToProto(e *containeranalysis.NoteVulnerabilityDetailsAffectedVersionEndKindEnum) containeranalysispb.ContaineranalysisNoteVulnerabilityDetailsAffectedVersionEndKindEnum {
	if e == nil {
		return containeranalysispb.ContaineranalysisNoteVulnerabilityDetailsAffectedVersionEndKindEnum(0)
	}
	if v, ok := containeranalysispb.ContaineranalysisNoteVulnerabilityDetailsAffectedVersionEndKindEnum_value["NoteVulnerabilityDetailsAffectedVersionEndKindEnum"+string(*e)]; ok {
		return containeranalysispb.ContaineranalysisNoteVulnerabilityDetailsAffectedVersionEndKindEnum(v)
	}
	return containeranalysispb.ContaineranalysisNoteVulnerabilityDetailsAffectedVersionEndKindEnum(0)
}

// NoteVulnerabilityDetailsFixedVersionKindEnumToProto converts a NoteVulnerabilityDetailsFixedVersionKindEnum enum to its proto representation.
func ContaineranalysisNoteVulnerabilityDetailsFixedVersionKindEnumToProto(e *containeranalysis.NoteVulnerabilityDetailsFixedVersionKindEnum) containeranalysispb.ContaineranalysisNoteVulnerabilityDetailsFixedVersionKindEnum {
	if e == nil {
		return containeranalysispb.ContaineranalysisNoteVulnerabilityDetailsFixedVersionKindEnum(0)
	}
	if v, ok := containeranalysispb.ContaineranalysisNoteVulnerabilityDetailsFixedVersionKindEnum_value["NoteVulnerabilityDetailsFixedVersionKindEnum"+string(*e)]; ok {
		return containeranalysispb.ContaineranalysisNoteVulnerabilityDetailsFixedVersionKindEnum(v)
	}
	return containeranalysispb.ContaineranalysisNoteVulnerabilityDetailsFixedVersionKindEnum(0)
}

// NoteVulnerabilityCvssV3AttackVectorEnumToProto converts a NoteVulnerabilityCvssV3AttackVectorEnum enum to its proto representation.
func ContaineranalysisNoteVulnerabilityCvssV3AttackVectorEnumToProto(e *containeranalysis.NoteVulnerabilityCvssV3AttackVectorEnum) containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3AttackVectorEnum {
	if e == nil {
		return containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3AttackVectorEnum(0)
	}
	if v, ok := containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3AttackVectorEnum_value["NoteVulnerabilityCvssV3AttackVectorEnum"+string(*e)]; ok {
		return containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3AttackVectorEnum(v)
	}
	return containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3AttackVectorEnum(0)
}

// NoteVulnerabilityCvssV3AttackComplexityEnumToProto converts a NoteVulnerabilityCvssV3AttackComplexityEnum enum to its proto representation.
func ContaineranalysisNoteVulnerabilityCvssV3AttackComplexityEnumToProto(e *containeranalysis.NoteVulnerabilityCvssV3AttackComplexityEnum) containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3AttackComplexityEnum {
	if e == nil {
		return containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3AttackComplexityEnum(0)
	}
	if v, ok := containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3AttackComplexityEnum_value["NoteVulnerabilityCvssV3AttackComplexityEnum"+string(*e)]; ok {
		return containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3AttackComplexityEnum(v)
	}
	return containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3AttackComplexityEnum(0)
}

// NoteVulnerabilityCvssV3PrivilegesRequiredEnumToProto converts a NoteVulnerabilityCvssV3PrivilegesRequiredEnum enum to its proto representation.
func ContaineranalysisNoteVulnerabilityCvssV3PrivilegesRequiredEnumToProto(e *containeranalysis.NoteVulnerabilityCvssV3PrivilegesRequiredEnum) containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3PrivilegesRequiredEnum {
	if e == nil {
		return containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3PrivilegesRequiredEnum(0)
	}
	if v, ok := containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3PrivilegesRequiredEnum_value["NoteVulnerabilityCvssV3PrivilegesRequiredEnum"+string(*e)]; ok {
		return containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3PrivilegesRequiredEnum(v)
	}
	return containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3PrivilegesRequiredEnum(0)
}

// NoteVulnerabilityCvssV3UserInteractionEnumToProto converts a NoteVulnerabilityCvssV3UserInteractionEnum enum to its proto representation.
func ContaineranalysisNoteVulnerabilityCvssV3UserInteractionEnumToProto(e *containeranalysis.NoteVulnerabilityCvssV3UserInteractionEnum) containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3UserInteractionEnum {
	if e == nil {
		return containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3UserInteractionEnum(0)
	}
	if v, ok := containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3UserInteractionEnum_value["NoteVulnerabilityCvssV3UserInteractionEnum"+string(*e)]; ok {
		return containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3UserInteractionEnum(v)
	}
	return containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3UserInteractionEnum(0)
}

// NoteVulnerabilityCvssV3ScopeEnumToProto converts a NoteVulnerabilityCvssV3ScopeEnum enum to its proto representation.
func ContaineranalysisNoteVulnerabilityCvssV3ScopeEnumToProto(e *containeranalysis.NoteVulnerabilityCvssV3ScopeEnum) containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3ScopeEnum {
	if e == nil {
		return containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3ScopeEnum(0)
	}
	if v, ok := containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3ScopeEnum_value["NoteVulnerabilityCvssV3ScopeEnum"+string(*e)]; ok {
		return containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3ScopeEnum(v)
	}
	return containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3ScopeEnum(0)
}

// NoteVulnerabilityCvssV3ConfidentialityImpactEnumToProto converts a NoteVulnerabilityCvssV3ConfidentialityImpactEnum enum to its proto representation.
func ContaineranalysisNoteVulnerabilityCvssV3ConfidentialityImpactEnumToProto(e *containeranalysis.NoteVulnerabilityCvssV3ConfidentialityImpactEnum) containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3ConfidentialityImpactEnum {
	if e == nil {
		return containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3ConfidentialityImpactEnum(0)
	}
	if v, ok := containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3ConfidentialityImpactEnum_value["NoteVulnerabilityCvssV3ConfidentialityImpactEnum"+string(*e)]; ok {
		return containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3ConfidentialityImpactEnum(v)
	}
	return containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3ConfidentialityImpactEnum(0)
}

// NoteVulnerabilityCvssV3IntegrityImpactEnumToProto converts a NoteVulnerabilityCvssV3IntegrityImpactEnum enum to its proto representation.
func ContaineranalysisNoteVulnerabilityCvssV3IntegrityImpactEnumToProto(e *containeranalysis.NoteVulnerabilityCvssV3IntegrityImpactEnum) containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3IntegrityImpactEnum {
	if e == nil {
		return containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3IntegrityImpactEnum(0)
	}
	if v, ok := containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3IntegrityImpactEnum_value["NoteVulnerabilityCvssV3IntegrityImpactEnum"+string(*e)]; ok {
		return containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3IntegrityImpactEnum(v)
	}
	return containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3IntegrityImpactEnum(0)
}

// NoteVulnerabilityCvssV3AvailabilityImpactEnumToProto converts a NoteVulnerabilityCvssV3AvailabilityImpactEnum enum to its proto representation.
func ContaineranalysisNoteVulnerabilityCvssV3AvailabilityImpactEnumToProto(e *containeranalysis.NoteVulnerabilityCvssV3AvailabilityImpactEnum) containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3AvailabilityImpactEnum {
	if e == nil {
		return containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3AvailabilityImpactEnum(0)
	}
	if v, ok := containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3AvailabilityImpactEnum_value["NoteVulnerabilityCvssV3AvailabilityImpactEnum"+string(*e)]; ok {
		return containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3AvailabilityImpactEnum(v)
	}
	return containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3AvailabilityImpactEnum(0)
}

// NotePackageDistributionArchitectureEnumToProto converts a NotePackageDistributionArchitectureEnum enum to its proto representation.
func ContaineranalysisNotePackageDistributionArchitectureEnumToProto(e *containeranalysis.NotePackageDistributionArchitectureEnum) containeranalysispb.ContaineranalysisNotePackageDistributionArchitectureEnum {
	if e == nil {
		return containeranalysispb.ContaineranalysisNotePackageDistributionArchitectureEnum(0)
	}
	if v, ok := containeranalysispb.ContaineranalysisNotePackageDistributionArchitectureEnum_value["NotePackageDistributionArchitectureEnum"+string(*e)]; ok {
		return containeranalysispb.ContaineranalysisNotePackageDistributionArchitectureEnum(v)
	}
	return containeranalysispb.ContaineranalysisNotePackageDistributionArchitectureEnum(0)
}

// NotePackageDistributionLatestVersionKindEnumToProto converts a NotePackageDistributionLatestVersionKindEnum enum to its proto representation.
func ContaineranalysisNotePackageDistributionLatestVersionKindEnumToProto(e *containeranalysis.NotePackageDistributionLatestVersionKindEnum) containeranalysispb.ContaineranalysisNotePackageDistributionLatestVersionKindEnum {
	if e == nil {
		return containeranalysispb.ContaineranalysisNotePackageDistributionLatestVersionKindEnum(0)
	}
	if v, ok := containeranalysispb.ContaineranalysisNotePackageDistributionLatestVersionKindEnum_value["NotePackageDistributionLatestVersionKindEnum"+string(*e)]; ok {
		return containeranalysispb.ContaineranalysisNotePackageDistributionLatestVersionKindEnum(v)
	}
	return containeranalysispb.ContaineranalysisNotePackageDistributionLatestVersionKindEnum(0)
}

// NoteDiscoveryAnalysisKindEnumToProto converts a NoteDiscoveryAnalysisKindEnum enum to its proto representation.
func ContaineranalysisNoteDiscoveryAnalysisKindEnumToProto(e *containeranalysis.NoteDiscoveryAnalysisKindEnum) containeranalysispb.ContaineranalysisNoteDiscoveryAnalysisKindEnum {
	if e == nil {
		return containeranalysispb.ContaineranalysisNoteDiscoveryAnalysisKindEnum(0)
	}
	if v, ok := containeranalysispb.ContaineranalysisNoteDiscoveryAnalysisKindEnum_value["NoteDiscoveryAnalysisKindEnum"+string(*e)]; ok {
		return containeranalysispb.ContaineranalysisNoteDiscoveryAnalysisKindEnum(v)
	}
	return containeranalysispb.ContaineranalysisNoteDiscoveryAnalysisKindEnum(0)
}

// NoteRelatedUrlToProto converts a NoteRelatedUrl object to its proto representation.
func ContaineranalysisNoteRelatedUrlToProto(o *containeranalysis.NoteRelatedUrl) *containeranalysispb.ContaineranalysisNoteRelatedUrl {
	if o == nil {
		return nil
	}
	p := &containeranalysispb.ContaineranalysisNoteRelatedUrl{}
	p.SetUrl(dcl.ValueOrEmptyString(o.Url))
	p.SetLabel(dcl.ValueOrEmptyString(o.Label))
	return p
}

// NoteVulnerabilityToProto converts a NoteVulnerability object to its proto representation.
func ContaineranalysisNoteVulnerabilityToProto(o *containeranalysis.NoteVulnerability) *containeranalysispb.ContaineranalysisNoteVulnerability {
	if o == nil {
		return nil
	}
	p := &containeranalysispb.ContaineranalysisNoteVulnerability{}
	p.SetCvssScore(dcl.ValueOrEmptyDouble(o.CvssScore))
	p.SetSeverity(ContaineranalysisNoteVulnerabilitySeverityEnumToProto(o.Severity))
	p.SetCvssV3(ContaineranalysisNoteVulnerabilityCvssV3ToProto(o.CvssV3))
	p.SetSourceUpdateTime(dcl.ValueOrEmptyString(o.SourceUpdateTime))
	sDetails := make([]*containeranalysispb.ContaineranalysisNoteVulnerabilityDetails, len(o.Details))
	for i, r := range o.Details {
		sDetails[i] = ContaineranalysisNoteVulnerabilityDetailsToProto(&r)
	}
	p.SetDetails(sDetails)
	sWindowsDetails := make([]*containeranalysispb.ContaineranalysisNoteVulnerabilityWindowsDetails, len(o.WindowsDetails))
	for i, r := range o.WindowsDetails {
		sWindowsDetails[i] = ContaineranalysisNoteVulnerabilityWindowsDetailsToProto(&r)
	}
	p.SetWindowsDetails(sWindowsDetails)
	return p
}

// NoteVulnerabilityDetailsToProto converts a NoteVulnerabilityDetails object to its proto representation.
func ContaineranalysisNoteVulnerabilityDetailsToProto(o *containeranalysis.NoteVulnerabilityDetails) *containeranalysispb.ContaineranalysisNoteVulnerabilityDetails {
	if o == nil {
		return nil
	}
	p := &containeranalysispb.ContaineranalysisNoteVulnerabilityDetails{}
	p.SetSeverityName(dcl.ValueOrEmptyString(o.SeverityName))
	p.SetDescription(dcl.ValueOrEmptyString(o.Description))
	p.SetPackageType(dcl.ValueOrEmptyString(o.PackageType))
	p.SetAffectedCpeUri(dcl.ValueOrEmptyString(o.AffectedCpeUri))
	p.SetAffectedPackage(dcl.ValueOrEmptyString(o.AffectedPackage))
	p.SetAffectedVersionStart(ContaineranalysisNoteVulnerabilityDetailsAffectedVersionStartToProto(o.AffectedVersionStart))
	p.SetAffectedVersionEnd(ContaineranalysisNoteVulnerabilityDetailsAffectedVersionEndToProto(o.AffectedVersionEnd))
	p.SetFixedCpeUri(dcl.ValueOrEmptyString(o.FixedCpeUri))
	p.SetFixedPackage(dcl.ValueOrEmptyString(o.FixedPackage))
	p.SetFixedVersion(ContaineranalysisNoteVulnerabilityDetailsFixedVersionToProto(o.FixedVersion))
	p.SetIsObsolete(dcl.ValueOrEmptyBool(o.IsObsolete))
	p.SetSourceUpdateTime(dcl.ValueOrEmptyString(o.SourceUpdateTime))
	return p
}

// NoteVulnerabilityDetailsAffectedVersionStartToProto converts a NoteVulnerabilityDetailsAffectedVersionStart object to its proto representation.
func ContaineranalysisNoteVulnerabilityDetailsAffectedVersionStartToProto(o *containeranalysis.NoteVulnerabilityDetailsAffectedVersionStart) *containeranalysispb.ContaineranalysisNoteVulnerabilityDetailsAffectedVersionStart {
	if o == nil {
		return nil
	}
	p := &containeranalysispb.ContaineranalysisNoteVulnerabilityDetailsAffectedVersionStart{}
	p.SetEpoch(dcl.ValueOrEmptyInt64(o.Epoch))
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetRevision(dcl.ValueOrEmptyString(o.Revision))
	p.SetKind(ContaineranalysisNoteVulnerabilityDetailsAffectedVersionStartKindEnumToProto(o.Kind))
	p.SetFullName(dcl.ValueOrEmptyString(o.FullName))
	return p
}

// NoteVulnerabilityDetailsAffectedVersionEndToProto converts a NoteVulnerabilityDetailsAffectedVersionEnd object to its proto representation.
func ContaineranalysisNoteVulnerabilityDetailsAffectedVersionEndToProto(o *containeranalysis.NoteVulnerabilityDetailsAffectedVersionEnd) *containeranalysispb.ContaineranalysisNoteVulnerabilityDetailsAffectedVersionEnd {
	if o == nil {
		return nil
	}
	p := &containeranalysispb.ContaineranalysisNoteVulnerabilityDetailsAffectedVersionEnd{}
	p.SetEpoch(dcl.ValueOrEmptyInt64(o.Epoch))
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetRevision(dcl.ValueOrEmptyString(o.Revision))
	p.SetKind(ContaineranalysisNoteVulnerabilityDetailsAffectedVersionEndKindEnumToProto(o.Kind))
	p.SetFullName(dcl.ValueOrEmptyString(o.FullName))
	return p
}

// NoteVulnerabilityDetailsFixedVersionToProto converts a NoteVulnerabilityDetailsFixedVersion object to its proto representation.
func ContaineranalysisNoteVulnerabilityDetailsFixedVersionToProto(o *containeranalysis.NoteVulnerabilityDetailsFixedVersion) *containeranalysispb.ContaineranalysisNoteVulnerabilityDetailsFixedVersion {
	if o == nil {
		return nil
	}
	p := &containeranalysispb.ContaineranalysisNoteVulnerabilityDetailsFixedVersion{}
	p.SetEpoch(dcl.ValueOrEmptyInt64(o.Epoch))
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetRevision(dcl.ValueOrEmptyString(o.Revision))
	p.SetKind(ContaineranalysisNoteVulnerabilityDetailsFixedVersionKindEnumToProto(o.Kind))
	p.SetFullName(dcl.ValueOrEmptyString(o.FullName))
	return p
}

// NoteVulnerabilityCvssV3ToProto converts a NoteVulnerabilityCvssV3 object to its proto representation.
func ContaineranalysisNoteVulnerabilityCvssV3ToProto(o *containeranalysis.NoteVulnerabilityCvssV3) *containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3 {
	if o == nil {
		return nil
	}
	p := &containeranalysispb.ContaineranalysisNoteVulnerabilityCvssV3{}
	p.SetBaseScore(dcl.ValueOrEmptyDouble(o.BaseScore))
	p.SetExploitabilityScore(dcl.ValueOrEmptyDouble(o.ExploitabilityScore))
	p.SetImpactScore(dcl.ValueOrEmptyDouble(o.ImpactScore))
	p.SetAttackVector(ContaineranalysisNoteVulnerabilityCvssV3AttackVectorEnumToProto(o.AttackVector))
	p.SetAttackComplexity(ContaineranalysisNoteVulnerabilityCvssV3AttackComplexityEnumToProto(o.AttackComplexity))
	p.SetPrivilegesRequired(ContaineranalysisNoteVulnerabilityCvssV3PrivilegesRequiredEnumToProto(o.PrivilegesRequired))
	p.SetUserInteraction(ContaineranalysisNoteVulnerabilityCvssV3UserInteractionEnumToProto(o.UserInteraction))
	p.SetScope(ContaineranalysisNoteVulnerabilityCvssV3ScopeEnumToProto(o.Scope))
	p.SetConfidentialityImpact(ContaineranalysisNoteVulnerabilityCvssV3ConfidentialityImpactEnumToProto(o.ConfidentialityImpact))
	p.SetIntegrityImpact(ContaineranalysisNoteVulnerabilityCvssV3IntegrityImpactEnumToProto(o.IntegrityImpact))
	p.SetAvailabilityImpact(ContaineranalysisNoteVulnerabilityCvssV3AvailabilityImpactEnumToProto(o.AvailabilityImpact))
	return p
}

// NoteVulnerabilityWindowsDetailsToProto converts a NoteVulnerabilityWindowsDetails object to its proto representation.
func ContaineranalysisNoteVulnerabilityWindowsDetailsToProto(o *containeranalysis.NoteVulnerabilityWindowsDetails) *containeranalysispb.ContaineranalysisNoteVulnerabilityWindowsDetails {
	if o == nil {
		return nil
	}
	p := &containeranalysispb.ContaineranalysisNoteVulnerabilityWindowsDetails{}
	p.SetCpeUri(dcl.ValueOrEmptyString(o.CpeUri))
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetDescription(dcl.ValueOrEmptyString(o.Description))
	sFixingKbs := make([]*containeranalysispb.ContaineranalysisNoteVulnerabilityWindowsDetailsFixingKbs, len(o.FixingKbs))
	for i, r := range o.FixingKbs {
		sFixingKbs[i] = ContaineranalysisNoteVulnerabilityWindowsDetailsFixingKbsToProto(&r)
	}
	p.SetFixingKbs(sFixingKbs)
	return p
}

// NoteVulnerabilityWindowsDetailsFixingKbsToProto converts a NoteVulnerabilityWindowsDetailsFixingKbs object to its proto representation.
func ContaineranalysisNoteVulnerabilityWindowsDetailsFixingKbsToProto(o *containeranalysis.NoteVulnerabilityWindowsDetailsFixingKbs) *containeranalysispb.ContaineranalysisNoteVulnerabilityWindowsDetailsFixingKbs {
	if o == nil {
		return nil
	}
	p := &containeranalysispb.ContaineranalysisNoteVulnerabilityWindowsDetailsFixingKbs{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetUrl(dcl.ValueOrEmptyString(o.Url))
	return p
}

// NoteBuildToProto converts a NoteBuild object to its proto representation.
func ContaineranalysisNoteBuildToProto(o *containeranalysis.NoteBuild) *containeranalysispb.ContaineranalysisNoteBuild {
	if o == nil {
		return nil
	}
	p := &containeranalysispb.ContaineranalysisNoteBuild{}
	p.SetBuilderVersion(dcl.ValueOrEmptyString(o.BuilderVersion))
	return p
}

// NoteImageToProto converts a NoteImage object to its proto representation.
func ContaineranalysisNoteImageToProto(o *containeranalysis.NoteImage) *containeranalysispb.ContaineranalysisNoteImage {
	if o == nil {
		return nil
	}
	p := &containeranalysispb.ContaineranalysisNoteImage{}
	p.SetResourceUrl(dcl.ValueOrEmptyString(o.ResourceUrl))
	p.SetFingerprint(ContaineranalysisNoteImageFingerprintToProto(o.Fingerprint))
	return p
}

// NoteImageFingerprintToProto converts a NoteImageFingerprint object to its proto representation.
func ContaineranalysisNoteImageFingerprintToProto(o *containeranalysis.NoteImageFingerprint) *containeranalysispb.ContaineranalysisNoteImageFingerprint {
	if o == nil {
		return nil
	}
	p := &containeranalysispb.ContaineranalysisNoteImageFingerprint{}
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
func ContaineranalysisNotePackageToProto(o *containeranalysis.NotePackage) *containeranalysispb.ContaineranalysisNotePackage {
	if o == nil {
		return nil
	}
	p := &containeranalysispb.ContaineranalysisNotePackage{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	sDistribution := make([]*containeranalysispb.ContaineranalysisNotePackageDistribution, len(o.Distribution))
	for i, r := range o.Distribution {
		sDistribution[i] = ContaineranalysisNotePackageDistributionToProto(&r)
	}
	p.SetDistribution(sDistribution)
	return p
}

// NotePackageDistributionToProto converts a NotePackageDistribution object to its proto representation.
func ContaineranalysisNotePackageDistributionToProto(o *containeranalysis.NotePackageDistribution) *containeranalysispb.ContaineranalysisNotePackageDistribution {
	if o == nil {
		return nil
	}
	p := &containeranalysispb.ContaineranalysisNotePackageDistribution{}
	p.SetCpeUri(dcl.ValueOrEmptyString(o.CpeUri))
	p.SetArchitecture(ContaineranalysisNotePackageDistributionArchitectureEnumToProto(o.Architecture))
	p.SetLatestVersion(ContaineranalysisNotePackageDistributionLatestVersionToProto(o.LatestVersion))
	p.SetMaintainer(dcl.ValueOrEmptyString(o.Maintainer))
	p.SetUrl(dcl.ValueOrEmptyString(o.Url))
	p.SetDescription(dcl.ValueOrEmptyString(o.Description))
	return p
}

// NotePackageDistributionLatestVersionToProto converts a NotePackageDistributionLatestVersion object to its proto representation.
func ContaineranalysisNotePackageDistributionLatestVersionToProto(o *containeranalysis.NotePackageDistributionLatestVersion) *containeranalysispb.ContaineranalysisNotePackageDistributionLatestVersion {
	if o == nil {
		return nil
	}
	p := &containeranalysispb.ContaineranalysisNotePackageDistributionLatestVersion{}
	p.SetEpoch(dcl.ValueOrEmptyInt64(o.Epoch))
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetRevision(dcl.ValueOrEmptyString(o.Revision))
	p.SetKind(ContaineranalysisNotePackageDistributionLatestVersionKindEnumToProto(o.Kind))
	p.SetFullName(dcl.ValueOrEmptyString(o.FullName))
	return p
}

// NoteDiscoveryToProto converts a NoteDiscovery object to its proto representation.
func ContaineranalysisNoteDiscoveryToProto(o *containeranalysis.NoteDiscovery) *containeranalysispb.ContaineranalysisNoteDiscovery {
	if o == nil {
		return nil
	}
	p := &containeranalysispb.ContaineranalysisNoteDiscovery{}
	p.SetAnalysisKind(ContaineranalysisNoteDiscoveryAnalysisKindEnumToProto(o.AnalysisKind))
	return p
}

// NoteDeploymentToProto converts a NoteDeployment object to its proto representation.
func ContaineranalysisNoteDeploymentToProto(o *containeranalysis.NoteDeployment) *containeranalysispb.ContaineranalysisNoteDeployment {
	if o == nil {
		return nil
	}
	p := &containeranalysispb.ContaineranalysisNoteDeployment{}
	sResourceUri := make([]string, len(o.ResourceUri))
	for i, r := range o.ResourceUri {
		sResourceUri[i] = r
	}
	p.SetResourceUri(sResourceUri)
	return p
}

// NoteAttestationToProto converts a NoteAttestation object to its proto representation.
func ContaineranalysisNoteAttestationToProto(o *containeranalysis.NoteAttestation) *containeranalysispb.ContaineranalysisNoteAttestation {
	if o == nil {
		return nil
	}
	p := &containeranalysispb.ContaineranalysisNoteAttestation{}
	p.SetHint(ContaineranalysisNoteAttestationHintToProto(o.Hint))
	return p
}

// NoteAttestationHintToProto converts a NoteAttestationHint object to its proto representation.
func ContaineranalysisNoteAttestationHintToProto(o *containeranalysis.NoteAttestationHint) *containeranalysispb.ContaineranalysisNoteAttestationHint {
	if o == nil {
		return nil
	}
	p := &containeranalysispb.ContaineranalysisNoteAttestationHint{}
	p.SetHumanReadableName(dcl.ValueOrEmptyString(o.HumanReadableName))
	return p
}

// NoteToProto converts a Note resource to its proto representation.
func NoteToProto(resource *containeranalysis.Note) *containeranalysispb.ContaineranalysisNote {
	p := &containeranalysispb.ContaineranalysisNote{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetShortDescription(dcl.ValueOrEmptyString(resource.ShortDescription))
	p.SetLongDescription(dcl.ValueOrEmptyString(resource.LongDescription))
	p.SetExpirationTime(dcl.ValueOrEmptyString(resource.ExpirationTime))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetVulnerability(ContaineranalysisNoteVulnerabilityToProto(resource.Vulnerability))
	p.SetBuild_(ContaineranalysisNoteBuildToProto(resource.Build))
	p.SetImage(ContaineranalysisNoteImageToProto(resource.Image))
	p.SetPackage(ContaineranalysisNotePackageToProto(resource.Package))
	p.SetDiscovery(ContaineranalysisNoteDiscoveryToProto(resource.Discovery))
	p.SetDeployment(ContaineranalysisNoteDeploymentToProto(resource.Deployment))
	p.SetAttestation(ContaineranalysisNoteAttestationToProto(resource.Attestation))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	sRelatedUrl := make([]*containeranalysispb.ContaineranalysisNoteRelatedUrl, len(resource.RelatedUrl))
	for i, r := range resource.RelatedUrl {
		sRelatedUrl[i] = ContaineranalysisNoteRelatedUrlToProto(&r)
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
func (s *NoteServer) applyNote(ctx context.Context, c *containeranalysis.Client, request *containeranalysispb.ApplyContaineranalysisNoteRequest) (*containeranalysispb.ContaineranalysisNote, error) {
	p := ProtoToNote(request.GetResource())
	res, err := c.ApplyNote(ctx, p)
	if err != nil {
		return nil, err
	}
	r := NoteToProto(res)
	return r, nil
}

// applyContaineranalysisNote handles the gRPC request by passing it to the underlying Note Apply() method.
func (s *NoteServer) ApplyContaineranalysisNote(ctx context.Context, request *containeranalysispb.ApplyContaineranalysisNoteRequest) (*containeranalysispb.ContaineranalysisNote, error) {
	cl, err := createConfigNote(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyNote(ctx, cl, request)
}

// DeleteNote handles the gRPC request by passing it to the underlying Note Delete() method.
func (s *NoteServer) DeleteContaineranalysisNote(ctx context.Context, request *containeranalysispb.DeleteContaineranalysisNoteRequest) (*emptypb.Empty, error) {

	cl, err := createConfigNote(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteNote(ctx, ProtoToNote(request.GetResource()))

}

// ListContaineranalysisNote handles the gRPC request by passing it to the underlying NoteList() method.
func (s *NoteServer) ListContaineranalysisNote(ctx context.Context, request *containeranalysispb.ListContaineranalysisNoteRequest) (*containeranalysispb.ListContaineranalysisNoteResponse, error) {
	cl, err := createConfigNote(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListNote(ctx, request.GetProject())
	if err != nil {
		return nil, err
	}
	var protos []*containeranalysispb.ContaineranalysisNote
	for _, r := range resources.Items {
		rp := NoteToProto(r)
		protos = append(protos, rp)
	}
	p := &containeranalysispb.ListContaineranalysisNoteResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigNote(ctx context.Context, service_account_file string) (*containeranalysis.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return containeranalysis.NewClient(conf), nil
}
