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

package apigee

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
	api "google.golang.org/api/apigee/v1"
)

func init() {
	fuzztesting.RegisterKRMFuzzer_NoProto(instanceFuzzer())
}

func instanceFuzzer() fuzztesting.KRMFuzzer_NoProto {
	f := fuzztesting.NewKRMTypedFuzzer_NoProto(&api.GoogleCloudApigeeV1Instance{},
		ApigeeInstanceSpec_FromAPI, ApigeeInstanceSpec_ToAPI,
		ApigeeInstanceObservedState_FromAPI, ApigeeInstanceObservedState_ToAPI,
	)

	f.SpecField(".AccessLoggingConfig")
	f.SpecField(".ConsumerAcceptList")
	f.SpecField(".Description")
	f.SpecField(".DiskEncryptionKeyName")
	f.SpecField(".DisplayName")
	f.SpecField(".IpRange")
	f.SpecField(".Location")
	f.SpecField(".PeeringCidrRange")

	f.StatusField(".CreatedAt")
	f.StatusField(".Host")
	f.StatusField(".LastModifiedAt")
	f.StatusField(".Port")
	f.StatusField(".RuntimeVersion")
	f.StatusField(".ServiceAttachment")
	f.StatusField(".State")

	f.Unimplemented_NotYetTriaged(".MaintenanceUpdatePolicy")
	f.Unimplemented_NotYetTriaged(".ScheduledMaintenance")
	f.Unimplemented_NotYetTriaged(".IsVersionLocked")
	f.Unimplemented_NotYetTriaged(".Name")
	f.Unimplemented_NotYetTriaged(".ForceSendFields")
	f.Unimplemented_NotYetTriaged(".NullFields")
	f.Unimplemented_NotYetTriaged(".ServerResponse")
	f.Unimplemented_NotYetTriaged(".AccessLoggingConfig.ForceSendFields")
	f.Unimplemented_NotYetTriaged(".AccessLoggingConfig.NullFields")

	return f
}
