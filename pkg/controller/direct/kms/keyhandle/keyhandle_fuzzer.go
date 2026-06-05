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
// proto.message: google.cloud.kms.v1.KeyHandle
// api.group: kms.cnrm.cloud.google.com

package keyhandle

import (
	pb "cloud.google.com/go/kms/apiv1/kmspb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(KMSKeyHandleFuzzer())
}

func KMSKeyHandleFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.KeyHandle{},
		KMSKeyHandleSpec_FromProto, KMSKeyHandleSpec_ToProto,
		KMSKeyHandleStatusObservedState_FromProto, KMSKeyHandleStatusObservedState_ToProto,
	)

	f.SpecField(".resource_type_selector")

	f.StatusField(".kms_key")

	f.Unimplemented_Identity(".name")

	return f
}
