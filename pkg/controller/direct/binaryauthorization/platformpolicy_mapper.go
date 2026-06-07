// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This file contains manual mappers for PlatformPolicy.

package binaryauthorization

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/binaryauthorization/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	api "google.golang.org/api/binaryauthorization/v1"
)

func AttestationAuthenticator_FromProto(mapCtx *direct.MapContext, in *api.AttestationAuthenticator) *krm.AttestationAuthenticator {
	if in == nil {
		return nil
	}
	out := &krm.AttestationAuthenticator{}
	out.DisplayName = direct.LazyPtr(in.DisplayName)
	out.PkixPublicKeySet = PkixPublicKeySet_FromProto(mapCtx, in.PkixPublicKeySet)
	return out
}
func AttestationAuthenticator_ToProto(mapCtx *direct.MapContext, in *krm.AttestationAuthenticator) *api.AttestationAuthenticator {
	if in == nil {
		return nil
	}
	out := &api.AttestationAuthenticator{}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.PkixPublicKeySet = PkixPublicKeySet_ToProto(mapCtx, in.PkixPublicKeySet)
	return out
}
func BinaryAuthorizationPlatformPolicyObservedState_FromProto(mapCtx *direct.MapContext, in *api.PlatformPolicy) *krm.BinaryAuthorizationPlatformPolicyObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BinaryAuthorizationPlatformPolicyObservedState{}
	// MISSING: Name
	// MISSING: Etag
	// MISSING: UpdateTime
	// MISSING: GKEPolicy
	return out
}
func BinaryAuthorizationPlatformPolicyObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BinaryAuthorizationPlatformPolicyObservedState) *api.PlatformPolicy {
	if in == nil {
		return nil
	}
	out := &api.PlatformPolicy{}
	// MISSING: Name
	// MISSING: Etag
	// MISSING: UpdateTime
	// MISSING: GKEPolicy
	return out
}
func BinaryAuthorizationPlatformPolicySpec_FromProto(mapCtx *direct.MapContext, in *api.PlatformPolicy) *krm.BinaryAuthorizationPlatformPolicySpec {
	if in == nil {
		return nil
	}
	out := &krm.BinaryAuthorizationPlatformPolicySpec{}
	// MISSING: Name
	out.Description = direct.LazyPtr(in.Description)
	// MISSING: Etag
	// MISSING: UpdateTime
	// MISSING: GKEPolicy
	// (near miss): "GKEPolicy" vs "GkePolicy"
	return out
}
func BinaryAuthorizationPlatformPolicySpec_ToProto(mapCtx *direct.MapContext, in *krm.BinaryAuthorizationPlatformPolicySpec) *api.PlatformPolicy {
	if in == nil {
		return nil
	}
	out := &api.PlatformPolicy{}
	// MISSING: Name
	out.Description = direct.ValueOf(in.Description)
	// MISSING: Etag
	// MISSING: UpdateTime
	// MISSING: GKEPolicy
	// (near miss): "GKEPolicy" vs "GkePolicy"
	return out
}
func Check_FromProto(mapCtx *direct.MapContext, in *api.Check) *krm.Check {
	if in == nil {
		return nil
	}
	out := &krm.Check{}
	out.SimpleSigningAttestationCheck = SimpleSigningAttestationCheck_FromProto(mapCtx, in.SimpleSigningAttestationCheck)
	out.ImageAllowlist = ImageAllowlist_FromProto(mapCtx, in.ImageAllowlist)
	out.AlwaysDeny = direct.LazyPtr(in.AlwaysDeny)
	out.SlsaCheck = SlsaCheck_FromProto(mapCtx, in.SlsaCheck)
	out.TrustedDirectoryCheck = TrustedDirectoryCheck_FromProto(mapCtx, in.TrustedDirectoryCheck)
	out.DisplayName = direct.LazyPtr(in.DisplayName)
	out.ImageFreshnessCheck = ImageFreshnessCheck_FromProto(mapCtx, in.ImageFreshnessCheck)
	out.VulnerabilityCheck = VulnerabilityCheck_FromProto(mapCtx, in.VulnerabilityCheck)
	out.SigstoreSignatureCheck = SigstoreSignatureCheck_FromProto(mapCtx, in.SigstoreSignatureCheck)
	return out
}
func Check_ToProto(mapCtx *direct.MapContext, in *krm.Check) *api.Check {
	if in == nil {
		return nil
	}
	out := &api.Check{}
	out.SimpleSigningAttestationCheck = SimpleSigningAttestationCheck_ToProto(mapCtx, in.SimpleSigningAttestationCheck)
	out.ImageAllowlist = ImageAllowlist_ToProto(mapCtx, in.ImageAllowlist)
	out.AlwaysDeny = direct.ValueOf(in.AlwaysDeny)
	out.SlsaCheck = SlsaCheck_ToProto(mapCtx, in.SlsaCheck)
	out.TrustedDirectoryCheck = TrustedDirectoryCheck_ToProto(mapCtx, in.TrustedDirectoryCheck)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.ImageFreshnessCheck = ImageFreshnessCheck_ToProto(mapCtx, in.ImageFreshnessCheck)
	out.VulnerabilityCheck = VulnerabilityCheck_ToProto(mapCtx, in.VulnerabilityCheck)
	out.SigstoreSignatureCheck = SigstoreSignatureCheck_ToProto(mapCtx, in.SigstoreSignatureCheck)
	return out
}
func CheckSet_FromProto(mapCtx *direct.MapContext, in *api.CheckSet) *krm.CheckSet {
	if in == nil {
		return nil
	}
	out := &krm.CheckSet{}
	out.Checks = direct.Slice_FromProto(mapCtx, in.Checks, Check_FromProto)
	out.DisplayName = direct.LazyPtr(in.DisplayName)
	out.ImageAllowlist = ImageAllowlist_FromProto(mapCtx, in.ImageAllowlist)
	out.Scope = Scope_FromProto(mapCtx, in.Scope)
	return out
}
func CheckSet_ToProto(mapCtx *direct.MapContext, in *krm.CheckSet) *api.CheckSet {
	if in == nil {
		return nil
	}
	out := &api.CheckSet{}
	out.Checks = direct.Slice_ToProto(mapCtx, in.Checks, Check_ToProto)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.ImageAllowlist = ImageAllowlist_ToProto(mapCtx, in.ImageAllowlist)
	out.Scope = Scope_ToProto(mapCtx, in.Scope)
	return out
}
func GKEPolicy_FromProto(mapCtx *direct.MapContext, in *api.GkePolicy) *krm.GKEPolicy {
	if in == nil {
		return nil
	}
	out := &krm.GKEPolicy{}
	out.CheckSets = direct.Slice_FromProto(mapCtx, in.CheckSets, CheckSet_FromProto)
	out.ImageAllowlist = ImageAllowlist_FromProto(mapCtx, in.ImageAllowlist)
	return out
}
func GKEPolicy_ToProto(mapCtx *direct.MapContext, in *krm.GKEPolicy) *api.GkePolicy {
	if in == nil {
		return nil
	}
	out := &api.GkePolicy{}
	out.CheckSets = direct.Slice_ToProto(mapCtx, in.CheckSets, CheckSet_ToProto)
	out.ImageAllowlist = ImageAllowlist_ToProto(mapCtx, in.ImageAllowlist)
	return out
}
func ImageAllowlist_FromProto(mapCtx *direct.MapContext, in *api.ImageAllowlist) *krm.ImageAllowlist {
	if in == nil {
		return nil
	}
	out := &krm.ImageAllowlist{}
	out.AllowPattern = in.AllowPattern
	return out
}
func ImageAllowlist_ToProto(mapCtx *direct.MapContext, in *krm.ImageAllowlist) *api.ImageAllowlist {
	if in == nil {
		return nil
	}
	out := &api.ImageAllowlist{}
	out.AllowPattern = in.AllowPattern
	return out
}
func ImageFreshnessCheck_FromProto(mapCtx *direct.MapContext, in *api.ImageFreshnessCheck) *krm.ImageFreshnessCheck {
	if in == nil {
		return nil
	}
	out := &krm.ImageFreshnessCheck{}
	if in.MaxUploadAgeDays != 0 {
		val := int32(in.MaxUploadAgeDays)
		out.MaxUploadAgeDays = &val
	}
	return out
}
func ImageFreshnessCheck_ToProto(mapCtx *direct.MapContext, in *krm.ImageFreshnessCheck) *api.ImageFreshnessCheck {
	if in == nil {
		return nil
	}
	out := &api.ImageFreshnessCheck{}
	if in.MaxUploadAgeDays != nil {
		out.MaxUploadAgeDays = int64(*in.MaxUploadAgeDays)
	}
	return out
}
func PkixPublicKey_FromProto(mapCtx *direct.MapContext, in *api.PkixPublicKey) *krm.PkixPublicKey {
	if in == nil {
		return nil
	}
	out := &krm.PkixPublicKey{}
	out.PublicKeyPem = direct.LazyPtr(in.PublicKeyPem)
	out.SignatureAlgorithm = direct.LazyPtr(in.SignatureAlgorithm)
	return out
}
func PkixPublicKey_ToProto(mapCtx *direct.MapContext, in *krm.PkixPublicKey) *api.PkixPublicKey {
	if in == nil {
		return nil
	}
	out := &api.PkixPublicKey{}
	out.PublicKeyPem = direct.ValueOf(in.PublicKeyPem)
	out.SignatureAlgorithm = direct.ValueOf(in.SignatureAlgorithm)
	return out
}
func PkixPublicKeySet_FromProto(mapCtx *direct.MapContext, in *api.PkixPublicKeySet) *krm.PkixPublicKeySet {
	if in == nil {
		return nil
	}
	out := &krm.PkixPublicKeySet{}
	out.PkixPublicKeys = direct.Slice_FromProto(mapCtx, in.PkixPublicKeys, PkixPublicKey_FromProto)
	return out
}
func PkixPublicKeySet_ToProto(mapCtx *direct.MapContext, in *krm.PkixPublicKeySet) *api.PkixPublicKeySet {
	if in == nil {
		return nil
	}
	out := &api.PkixPublicKeySet{}
	out.PkixPublicKeys = direct.Slice_ToProto(mapCtx, in.PkixPublicKeys, PkixPublicKey_ToProto)
	return out
}
func PlatformPolicy_FromProto(mapCtx *direct.MapContext, in *api.PlatformPolicy) *krm.PlatformPolicy {
	if in == nil {
		return nil
	}
	out := &krm.PlatformPolicy{}
	// MISSING: Name
	out.Description = direct.LazyPtr(in.Description)
	out.Etag = direct.LazyPtr(in.Etag)
	// MISSING: UpdateTime
	out.GKEPolicy = GKEPolicy_FromProto(mapCtx, in.GkePolicy)
	return out
}
func PlatformPolicy_ToProto(mapCtx *direct.MapContext, in *krm.PlatformPolicy) *api.PlatformPolicy {
	if in == nil {
		return nil
	}
	out := &api.PlatformPolicy{}
	// MISSING: Name
	out.Description = direct.ValueOf(in.Description)
	out.Etag = direct.ValueOf(in.Etag)
	// MISSING: UpdateTime
	out.GkePolicy = GKEPolicy_ToProto(mapCtx, in.GKEPolicy)
	return out
}
func PlatformPolicyObservedState_FromProto(mapCtx *direct.MapContext, in *api.PlatformPolicy) *krm.PlatformPolicyObservedState {
	if in == nil {
		return nil
	}
	out := &krm.PlatformPolicyObservedState{}
	out.Name = direct.LazyPtr(in.Name)
	// MISSING: Description
	// MISSING: Etag
	out.UpdateTime = direct.LazyPtr(in.UpdateTime)
	// MISSING: GKEPolicy
	return out
}
func PlatformPolicyObservedState_ToProto(mapCtx *direct.MapContext, in *krm.PlatformPolicyObservedState) *api.PlatformPolicy {
	if in == nil {
		return nil
	}
	out := &api.PlatformPolicy{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: Description
	// MISSING: Etag
	out.UpdateTime = direct.ValueOf(in.UpdateTime)
	// MISSING: GKEPolicy
	return out
}
func Scope_FromProto(mapCtx *direct.MapContext, in *api.Scope) *krm.Scope {
	if in == nil {
		return nil
	}
	out := &krm.Scope{}
	if in.KubernetesServiceAccount != "" {
		// TODO: parse namespace:name
	}
	out.KubernetesNamespace = direct.LazyPtr(in.KubernetesNamespace)
	return out
}
func Scope_ToProto(mapCtx *direct.MapContext, in *krm.Scope) *api.Scope {
	if in == nil {
		return nil
	}
	out := &api.Scope{}
	if in.KubernetesServiceAccountRef != nil {
		out.KubernetesServiceAccount = in.KubernetesServiceAccountRef.Namespace + ":" + in.KubernetesServiceAccountRef.Name
	}
	out.KubernetesNamespace = direct.ValueOf(in.KubernetesNamespace)
	return out
}
func SigstoreAuthority_FromProto(mapCtx *direct.MapContext, in *api.SigstoreAuthority) *krm.SigstoreAuthority {
	if in == nil {
		return nil
	}
	out := &krm.SigstoreAuthority{}
	out.DisplayName = direct.LazyPtr(in.DisplayName)
	out.PublicKeySet = SigstorePublicKeySet_FromProto(mapCtx, in.PublicKeySet)
	return out
}
func SigstoreAuthority_ToProto(mapCtx *direct.MapContext, in *krm.SigstoreAuthority) *api.SigstoreAuthority {
	if in == nil {
		return nil
	}
	out := &api.SigstoreAuthority{}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.PublicKeySet = SigstorePublicKeySet_ToProto(mapCtx, in.PublicKeySet)
	return out
}
func SigstorePublicKey_FromProto(mapCtx *direct.MapContext, in *api.SigstorePublicKey) *krm.SigstorePublicKey {
	if in == nil {
		return nil
	}
	out := &krm.SigstorePublicKey{}
	out.PublicKeyPem = direct.LazyPtr(in.PublicKeyPem)
	return out
}
func SigstorePublicKey_ToProto(mapCtx *direct.MapContext, in *krm.SigstorePublicKey) *api.SigstorePublicKey {
	if in == nil {
		return nil
	}
	out := &api.SigstorePublicKey{}
	out.PublicKeyPem = direct.ValueOf(in.PublicKeyPem)
	return out
}
func SigstorePublicKeySet_FromProto(mapCtx *direct.MapContext, in *api.SigstorePublicKeySet) *krm.SigstorePublicKeySet {
	if in == nil {
		return nil
	}
	out := &krm.SigstorePublicKeySet{}
	out.PublicKeys = direct.Slice_FromProto(mapCtx, in.PublicKeys, SigstorePublicKey_FromProto)
	return out
}
func SigstorePublicKeySet_ToProto(mapCtx *direct.MapContext, in *krm.SigstorePublicKeySet) *api.SigstorePublicKeySet {
	if in == nil {
		return nil
	}
	out := &api.SigstorePublicKeySet{}
	out.PublicKeys = direct.Slice_ToProto(mapCtx, in.PublicKeys, SigstorePublicKey_ToProto)
	return out
}
func SigstoreSignatureCheck_FromProto(mapCtx *direct.MapContext, in *api.SigstoreSignatureCheck) *krm.SigstoreSignatureCheck {
	if in == nil {
		return nil
	}
	out := &krm.SigstoreSignatureCheck{}
	out.SigstoreAuthorities = direct.Slice_FromProto(mapCtx, in.SigstoreAuthorities, SigstoreAuthority_FromProto)
	return out
}
func SigstoreSignatureCheck_ToProto(mapCtx *direct.MapContext, in *krm.SigstoreSignatureCheck) *api.SigstoreSignatureCheck {
	if in == nil {
		return nil
	}
	out := &api.SigstoreSignatureCheck{}
	out.SigstoreAuthorities = direct.Slice_ToProto(mapCtx, in.SigstoreAuthorities, SigstoreAuthority_ToProto)
	return out
}
func SimpleSigningAttestationCheck_FromProto(mapCtx *direct.MapContext, in *api.SimpleSigningAttestationCheck) *krm.SimpleSigningAttestationCheck {
	if in == nil {
		return nil
	}
	out := &krm.SimpleSigningAttestationCheck{}
	out.ContainerAnalysisAttestationProjects = in.ContainerAnalysisAttestationProjects
	out.AttestationAuthenticators = direct.Slice_FromProto(mapCtx, in.AttestationAuthenticators, AttestationAuthenticator_FromProto)
	return out
}
func SimpleSigningAttestationCheck_ToProto(mapCtx *direct.MapContext, in *krm.SimpleSigningAttestationCheck) *api.SimpleSigningAttestationCheck {
	if in == nil {
		return nil
	}
	out := &api.SimpleSigningAttestationCheck{}
	out.ContainerAnalysisAttestationProjects = in.ContainerAnalysisAttestationProjects
	out.AttestationAuthenticators = direct.Slice_ToProto(mapCtx, in.AttestationAuthenticators, AttestationAuthenticator_ToProto)
	return out
}
func SlsaCheck_FromProto(mapCtx *direct.MapContext, in *api.SlsaCheck) *krm.SlsaCheck {
	if in == nil {
		return nil
	}
	out := &krm.SlsaCheck{}
	out.Rules = direct.Slice_FromProto(mapCtx, in.Rules, VerificationRule_FromProto)
	return out
}
func SlsaCheck_ToProto(mapCtx *direct.MapContext, in *krm.SlsaCheck) *api.SlsaCheck {
	if in == nil {
		return nil
	}
	out := &api.SlsaCheck{}
	out.Rules = direct.Slice_ToProto(mapCtx, in.Rules, VerificationRule_ToProto)
	return out
}
func TrustedDirectoryCheck_FromProto(mapCtx *direct.MapContext, in *api.TrustedDirectoryCheck) *krm.TrustedDirectoryCheck {
	if in == nil {
		return nil
	}
	out := &krm.TrustedDirectoryCheck{}
	out.TrustedDirPatterns = in.TrustedDirPatterns
	return out
}
func TrustedDirectoryCheck_ToProto(mapCtx *direct.MapContext, in *krm.TrustedDirectoryCheck) *api.TrustedDirectoryCheck {
	if in == nil {
		return nil
	}
	out := &api.TrustedDirectoryCheck{}
	out.TrustedDirPatterns = in.TrustedDirPatterns
	return out
}
func VerificationRule_FromProto(mapCtx *direct.MapContext, in *api.VerificationRule) *krm.VerificationRule {
	if in == nil {
		return nil
	}
	out := &krm.VerificationRule{}
	out.TrustedBuilder = direct.LazyPtr(in.TrustedBuilder)
	out.TrustedSourceRepoPatterns = in.TrustedSourceRepoPatterns
	return out
}
func VerificationRule_ToProto(mapCtx *direct.MapContext, in *krm.VerificationRule) *api.VerificationRule {
	if in == nil {
		return nil
	}
	out := &api.VerificationRule{}
	out.TrustedBuilder = direct.ValueOf(in.TrustedBuilder)
	out.TrustedSourceRepoPatterns = in.TrustedSourceRepoPatterns
	return out
}
func VulnerabilityCheck_FromProto(mapCtx *direct.MapContext, in *api.VulnerabilityCheck) *krm.VulnerabilityCheck {
	if in == nil {
		return nil
	}
	out := &krm.VulnerabilityCheck{}
	out.BlockedCves = in.BlockedCves
	out.MaximumUnfixableSeverity = direct.LazyPtr(in.MaximumUnfixableSeverity)
	out.MaximumFixableSeverity = direct.LazyPtr(in.MaximumFixableSeverity)
	out.AllowedCves = in.AllowedCves
	out.ContainerAnalysisVulnerabilityProjects = in.ContainerAnalysisVulnerabilityProjects
	return out
}
func VulnerabilityCheck_ToProto(mapCtx *direct.MapContext, in *krm.VulnerabilityCheck) *api.VulnerabilityCheck {
	if in == nil {
		return nil
	}
	out := &api.VulnerabilityCheck{}
	out.BlockedCves = in.BlockedCves
	out.MaximumUnfixableSeverity = direct.ValueOf(in.MaximumUnfixableSeverity)
	out.MaximumFixableSeverity = direct.ValueOf(in.MaximumFixableSeverity)
	out.AllowedCves = in.AllowedCves
	out.ContainerAnalysisVulnerabilityProjects = in.ContainerAnalysisVulnerabilityProjects
	return out
}
