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
// proto.message: google.cloud.certificatemanager.v1.TrustConfig
// api.group: certificatemanager.cnrm.cloud.google.com

package certificatemanager

import (
	pb "cloud.google.com/go/certificatemanager/apiv1/certificatemanagerpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(trustConfigFuzzer())
}

func trustConfigFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.TrustConfig{},
		CertificateManagerTrustConfigSpec_v1alpha1_FromProto, CertificateManagerTrustConfigSpec_v1alpha1_ToProto,
		CertificateManagerTrustConfigObservedState_v1alpha1_FromProto, CertificateManagerTrustConfigObservedState_v1alpha1_ToProto,
	)

	f.SpecField(".description")
	f.SpecField(".trust_stores")

	f.StatusField(".create_time")
	f.StatusField(".update_time")

	f.Unimplemented_Identity(".name")
	f.Unimplemented_LabelsAnnotations(".labels")
	f.Unimplemented_NotYetTriaged(".etag")

	return f
}
