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
// proto.message: google.api.cloudquotas.v1beta.QuotaPreference
// api.group: cloudquota.cnrm.cloud.google.com

package cloudquota

import (
	pb "cloud.google.com/go/cloudquotas/apiv1beta/cloudquotaspb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(aPIQuotaPreferenceFuzzer())
}

func aPIQuotaPreferenceFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.QuotaPreference{},
		APIQuotaPreferenceSpec_FromProto, APIQuotaPreferenceSpec_ToProto,
		APIQuotaPreferenceObservedState_FromProto, APIQuotaPreferenceObservedState_ToProto,
	)

	f.SpecFields.Insert(".dimensions")
	f.SpecFields.Insert(".quota_config")
	f.SpecFields.Insert(".service")
	f.SpecFields.Insert(".quota_id")
	f.SpecFields.Insert(".justification")
	f.SpecFields.Insert(".contact_email")

	f.StatusFields.Insert(".quota_config")
	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".update_time")
	f.StatusFields.Insert(".reconciling")
	f.StatusFields.Insert(".etag")

	f.UnimplementedFields.Insert(".name") // special field

	return f
}
