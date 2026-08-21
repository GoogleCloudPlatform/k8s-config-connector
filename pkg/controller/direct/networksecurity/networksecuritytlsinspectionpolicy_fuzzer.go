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
// proto.message: google.cloud.networksecurity.v1.TlsInspectionPolicy
// api.group: networksecurity.cnrm.cloud.google.com

package networksecurity

import (
	pb "cloud.google.com/go/networksecurity/apiv1/networksecuritypb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(tlsInspectionPolicyFuzzer())
}

func tlsInspectionPolicyFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.TlsInspectionPolicy{},
		NetworkSecurityTLSInspectionPolicySpec_v1alpha1_FromProto, NetworkSecurityTLSInspectionPolicySpec_v1alpha1_ToProto,
		NetworkSecurityTLSInspectionPolicyObservedState_v1alpha1_FromProto, NetworkSecurityTLSInspectionPolicyObservedState_v1alpha1_ToProto,
	)

	f.SpecField(".description")
	f.SpecField(".exclude_public_ca_set")
	f.SpecField(".min_tls_version")
	f.SpecField(".tls_feature_profile")
	f.SpecField(".custom_tls_features")

	// Custom/Reference fields not mapped automatically in mapper.generated.go (handled in adapter instead)
	f.Unimplemented_NotYetTriaged(".ca_pool")
	f.Unimplemented_NotYetTriaged(".trust_config")

	f.StatusField(".create_time")
	f.StatusField(".update_time")

	f.IdentityField(".name")

	return f
}
