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
// proto.message: google.cloud.kms.v1.ImportJob
// api.group: kms.cnrm.cloud.google.com

package kms

import (
	pb "cloud.google.com/go/kms/apiv1/kmspb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(KMSImportJobFuzzer())
}

func KMSImportJobFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.ImportJob{},
		KMSImportJobSpec_FromProto, KMSImportJobSpec_ToProto,
		KMSImportJobObservedState_FromProto, KMSImportJobObservedState_ToProto,
	)

	f.SpecField(".import_method")
	f.SpecField(".protection_level")

	f.StatusField(".create_time")
	f.StatusField(".generate_time")
	f.StatusField(".expire_time")
	f.StatusField(".expire_event_time")
	f.StatusField(".state")
	f.StatusField(".public_key")
	f.StatusField(".attestation")

	f.Unimplemented_Identity(".name") // special field
	f.Unimplemented_NotYetTriaged(".crypto_key_backend")

	return f
}
