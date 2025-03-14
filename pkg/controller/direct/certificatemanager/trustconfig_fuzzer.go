// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
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
	fuzztesting.RegisterKRMFuzzer(certificateManagerTrustConfigFuzzer())
}

func certificateManagerTrustConfigFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.TrustConfig{},
		CertificateManagerTrustConfigSpec_FromProto, CertificateManagerTrustConfigSpec_ToProto,
		CertificateManagerTrustConfigObservedState_FromProto, CertificateManagerTrustConfigObservedState_ToProto,
	)

	f.SpecFields.Insert(".labels")
	f.SpecFields.Insert(".description")
	f.SpecFields.Insert(".trust_stores")
	f.SpecFields.Insert(".etag")

	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".update_time")

	f.UnimplementedFields.Insert(".name")

	return f
}


