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

package binaryauthorizationplatformpolicy

import (
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
	api "google.golang.org/api/binaryauthorization/v1"
	"google.golang.org/api/googleapi"
)

func init() {
	fuzztesting.RegisterKRMFuzzer_NoProto(fuzzBinaryAuthorizationPlatformPolicy())
}

func fuzzBinaryAuthorizationPlatformPolicy() fuzztesting.KRMFuzzer_NoProto {
	f := fuzztesting.NewKRMTypedFuzzer_NoProto(&api.PlatformPolicy{},
		BinaryAuthorizationPlatformPolicySpec_FromProto, BinaryAuthorizationPlatformPolicySpec_ToProto,
		BinaryAuthorizationPlatformPolicyObservedState_FromProto, BinaryAuthorizationPlatformPolicyObservedState_ToProto,
	)

	f.Ignore_JSONBookkeeping(".ForceSendFields")
	f.Ignore_JSONBookkeeping(".NullFields")
	f.Ignore_JSONBookkeeping(".ServerResponse")

	f.Ignore_JSONBookkeeping(".etag")

	f.SpecField(".description")
	f.SpecField(".gke_policy")

	f.StatusField(".name")
	f.StatusField(".update_time")

	// Custom filter to ignore Status/bookkeeping fields within Spec fuzzing
	f.FilterSpec = func(in *api.PlatformPolicy) {
		if in == nil {
			return
		}
		in.Name = ""
		in.UpdateTime = ""
		in.Etag = ""
		in.ForceSendFields = nil
		in.NullFields = nil
		in.ServerResponse = googleapi.ServerResponse{}
		if in.GkePolicy != nil {
		        in.GkePolicy.ForceSendFields = nil
		        in.GkePolicy.NullFields = nil
		        if in.GkePolicy.ImageAllowlist != nil {
		                in.GkePolicy.ImageAllowlist.ForceSendFields = nil
		                in.GkePolicy.ImageAllowlist.NullFields = nil
		        }
		        if len(in.GkePolicy.CheckSets) > 1 {
		                in.GkePolicy.CheckSets = in.GkePolicy.CheckSets[:1]
		        }
		        for _, cs := range in.GkePolicy.CheckSets {
		                if cs != nil {
		                        cs.ForceSendFields = nil
		                        cs.NullFields = nil
		                        if cs.ImageAllowlist != nil {
		                                cs.ImageAllowlist.ForceSendFields = nil
		                                cs.ImageAllowlist.NullFields = nil
		                        }
		                        if cs.Scope != nil {
		                                cs.Scope.ForceSendFields = nil
		                                cs.Scope.NullFields = nil
		                                if cs.Scope.KubernetesServiceAccount != "" {
		                                        parts := strings.Split(cs.Scope.KubernetesServiceAccount, ":")
		                                        if len(parts) == 2 {
		                                                if parts[0] == "" || parts[1] == "" {
		                                                        cs.Scope.KubernetesServiceAccount = "ns:sa"
		                                                }
		                                        } else {
		                                                cs.Scope.KubernetesServiceAccount = "ns:" + cs.Scope.KubernetesServiceAccount
		                                        }
		                                }
		                        }
		                        if len(cs.Checks) > 1 {
		                                cs.Checks = cs.Checks[:1]
		                        }
		                        for _, check := range cs.Checks {						if check != nil {
							check.ForceSendFields = nil
							check.NullFields = nil
							if check.ImageAllowlist != nil {
								check.ImageAllowlist.ForceSendFields = nil
								check.ImageAllowlist.NullFields = nil
							}
							if check.ImageFreshnessCheck != nil {
								check.ImageFreshnessCheck.ForceSendFields = nil
								check.ImageFreshnessCheck.NullFields = nil
								// Prevent overflow truncation diffs
								check.ImageFreshnessCheck.MaxUploadAgeDays = int64(int32(check.ImageFreshnessCheck.MaxUploadAgeDays))
							}
							if check.SigstoreSignatureCheck != nil {
							        check.SigstoreSignatureCheck.ForceSendFields = nil
							        check.SigstoreSignatureCheck.NullFields = nil
							        if len(check.SigstoreSignatureCheck.SigstoreAuthorities) > 1 {
							                check.SigstoreSignatureCheck.SigstoreAuthorities = check.SigstoreSignatureCheck.SigstoreAuthorities[:1]
							        }
							        for _, sa := range check.SigstoreSignatureCheck.SigstoreAuthorities {
							                if sa != nil {
							                       sa.ForceSendFields = nil
							                       sa.NullFields = nil
							                       if sa.PublicKeySet != nil {
							                       sa.PublicKeySet.ForceSendFields = nil
							                       sa.PublicKeySet.NullFields = nil
							                       if len(sa.PublicKeySet.PublicKeys) > 1 {
							                               sa.PublicKeySet.PublicKeys = sa.PublicKeySet.PublicKeys[:1]
							                       }
							                       for _, pk := range sa.PublicKeySet.PublicKeys {
							                       if pk != nil {
							                       pk.ForceSendFields = nil
							                       pk.NullFields = nil
							                       }
							                       }
							                       }
							                }
							        }
							}
							if check.SlsaCheck != nil {
							        check.SlsaCheck.ForceSendFields = nil
							        check.SlsaCheck.NullFields = nil
							        if len(check.SlsaCheck.Rules) > 1 {
							                check.SlsaCheck.Rules = check.SlsaCheck.Rules[:1]
							        }
							        for _, rule := range check.SlsaCheck.Rules {
							                if rule != nil {
							                       rule.ForceSendFields = nil
							                       rule.NullFields = nil
							                       rule.AttestationSource = nil
							                       rule.ConfigBasedBuildRequired = false
							                       rule.CustomConstraints = ""
							                }
							        }
							}
							if check.TrustedDirectoryCheck != nil {
							        check.TrustedDirectoryCheck.ForceSendFields = nil
							        check.TrustedDirectoryCheck.NullFields = nil
							}
							if check.VulnerabilityCheck != nil {
							        check.VulnerabilityCheck.ForceSendFields = nil
							        check.VulnerabilityCheck.NullFields = nil
							}
							if check.SimpleSigningAttestationCheck != nil {
							        check.SimpleSigningAttestationCheck.ForceSendFields = nil
							        check.SimpleSigningAttestationCheck.NullFields = nil
							        if len(check.SimpleSigningAttestationCheck.AttestationAuthenticators) > 1 {
							                check.SimpleSigningAttestationCheck.AttestationAuthenticators = check.SimpleSigningAttestationCheck.AttestationAuthenticators[:1]
							        }
							        for _, auth := range check.SimpleSigningAttestationCheck.AttestationAuthenticators {
							                if auth != nil {
							                       auth.ForceSendFields = nil
							                       auth.NullFields = nil
							                       if auth.PkixPublicKeySet != nil {
							                       auth.PkixPublicKeySet.ForceSendFields = nil
							                       auth.PkixPublicKeySet.NullFields = nil
							                       if len(auth.PkixPublicKeySet.PkixPublicKeys) > 1 {
							                               auth.PkixPublicKeySet.PkixPublicKeys = auth.PkixPublicKeySet.PkixPublicKeys[:1]
							                       }
							                       for _, pk := range auth.PkixPublicKeySet.PkixPublicKeys {
							                       if pk != nil {
							                       pk.ForceSendFields = nil
							                       pk.NullFields = nil
							                       pk.KeyId = ""
							                       }
							                       }
							                       }
							                }
							        }
							}						}
					}
				}
			}
		}
	}

	// Custom filter to ignore Spec/bookkeeping fields within Status fuzzing
	f.FilterStatus = func(in *api.PlatformPolicy) {
		if in == nil {
			return
		}
		in.Description = ""
		in.GkePolicy = nil
		in.Etag = ""
		in.ForceSendFields = nil
		in.NullFields = nil
		in.ServerResponse = googleapi.ServerResponse{}
	}

	return f
}
