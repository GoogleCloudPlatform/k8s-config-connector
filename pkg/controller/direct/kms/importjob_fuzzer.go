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

	f.SpecFields.Insert(".import_method")
	f.SpecFields.Insert(".protection_level")

	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".generate_time")
	f.StatusFields.Insert(".expire_time")
	f.StatusFields.Insert(".expire_event_time")
	f.StatusFields.Insert(".state")
	f.StatusFields.Insert(".public_key")
	f.StatusFields.Insert(".attestation")

	f.UnimplementedFields.Insert(".name") // special field

	return f
}
