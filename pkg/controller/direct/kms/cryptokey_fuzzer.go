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
// proto.message: google.cloud.kms.v1.CryptoKey
// api.group: kms.cnrm.cloud.google.com

package kms

import (
	pb "cloud.google.com/go/kms/apiv1/kmspb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(KMSCryptoKeyFuzzer())
}

func KMSCryptoKeyFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.CryptoKey{},
		KMSCryptoKeySpec_FromProto, KMSCryptoKeySpec_ToProto,
		KMSCryptoKeyStatus_FromProto, KMSCryptoKeyStatus_ToProto,
	)

	f.SpecField(".purpose")
	f.SpecField(".rotation_period")
	f.SpecField(".version_template")
	f.SpecField(".version_template.algorithm")
	f.SpecField(".version_template.protection_level")
	f.SpecField(".import_only")
	f.SpecField(".destroy_scheduled_duration")

	f.Unimplemented_Identity(".name")

	f.Unimplemented_NotYetTriaged(".primary")
	f.Unimplemented_NotYetTriaged(".create_time")
	f.Unimplemented_NotYetTriaged(".next_rotation_time")
	f.Unimplemented_NotYetTriaged(".labels")
	f.Unimplemented_NotYetTriaged(".crypto_key_backend")
	f.Unimplemented_NotYetTriaged(".key_access_justifications_policy")

	return f
}
