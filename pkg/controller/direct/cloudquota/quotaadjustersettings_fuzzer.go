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
// proto.message: google.api.cloudquotas.v1beta.QuotaAdjusterSettings
// api.group: cloudquota.cnrm.cloud.google.com

package cloudquota

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"

	pb "cloud.google.com/go/cloudquotas/apiv1beta/cloudquotaspb"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(aPIQuotaAdjusterSettingsFuzzer())
}

func aPIQuotaAdjusterSettingsFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.QuotaAdjusterSettings{},
		APIQuotaAdjusterSettingsSpec_FromProto, APIQuotaAdjusterSettingsSpec_ToProto,
		APIQuotaAdjusterSettingsObservedState_FromProto, APIQuotaAdjusterSettingsObservedState_ToProto,
	)
	f.SpecFields.Insert(".enablement")

	f.StatusFields.Insert(".update_time")
	f.StatusFields.Insert(".etag")

	f.UnimplementedFields.Insert(".name") // special field

	return f
}
