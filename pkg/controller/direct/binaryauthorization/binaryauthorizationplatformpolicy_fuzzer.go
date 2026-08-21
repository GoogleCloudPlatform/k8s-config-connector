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

// +tool:fuzz-gen
// proto.message: google.cloud.binaryauthorization.v1.PlatformPolicy
// api.group: binaryauthorization.cnrm.cloud.google.com

package binaryauthorization

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
	api "google.golang.org/api/binaryauthorization/v1"
)

func init() {
	fuzztesting.RegisterKRMFuzzer_NoProto(binaryAuthorizationPlatformPolicyFuzzer())
}

func binaryAuthorizationPlatformPolicyFuzzer() fuzztesting.KRMFuzzer_NoProto {
	f := fuzztesting.NewKRMTypedFuzzer_NoProto(&api.PlatformPolicy{},
		BinaryAuthorizationPlatformPolicySpec_FromProto,
		BinaryAuthorizationPlatformPolicySpec_ToProto,
		BinaryAuthorizationPlatformPolicyObservedState_FromProto,
		BinaryAuthorizationPlatformPolicyObservedState_ToProto,
	)

	f.SpecField(".Description")
	f.SpecField(".GkePolicy")

	f.StatusField(".Name")
	f.StatusField(".UpdateTime")

	f.IdentityField(".Name")
	f.IdentityField(".Etag")
	f.IdentityField(".UpdateTime")

	f.Ignore_JSONBookkeeping(".GkePolicy.CheckSets[].Scope.KubernetesServiceAccount")
	f.Ignore_JSONBookkeeping(".GkePolicy.CheckSets[].Checks[].SimpleSigningAttestationCheck.AttestationAuthenticators[].PkixPublicKeySet.PkixPublicKeys[].KeyId")
	f.Ignore_JSONBookkeeping(".GkePolicy.CheckSets[].Checks[].SlsaCheck.Rules[].AttestationSource")
	f.Ignore_JSONBookkeeping(".GkePolicy.CheckSets[].Checks[].SlsaCheck.Rules[].CustomConstraints")
	f.Ignore_JSONBookkeeping(".GkePolicy.CheckSets[].Checks[].SlsaCheck.Rules[].ConfigBasedBuildRequired")

	f.FilterSpec = func(p *api.PlatformPolicy) {
		if p.GkePolicy != nil {
			for _, cs := range p.GkePolicy.CheckSets {
				if cs == nil {
					continue
				}
				for _, c := range cs.Checks {
					if c == nil {
						continue
					}
					if c.ImageFreshnessCheck != nil {
						c.ImageFreshnessCheck.MaxUploadAgeDays = int64(int32(c.ImageFreshnessCheck.MaxUploadAgeDays))
					}
				}
			}
		}
	}

	return f
}
