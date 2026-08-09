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

package licensemanager

import (
	pb "cloud.google.com/go/licensemanager/apiv1/licensemanagerpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(licenseManagerConfigurationFuzzer())
}

func licenseManagerConfigurationFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Configuration{},
		LicenseManagerConfigurationSpec_FromProto, LicenseManagerConfigurationSpec_ToProto,
		LicenseManagerConfigurationObservedState_FromProto, LicenseManagerConfigurationObservedState_ToProto,
	)

	// Identity Field
	f.Unimplemented_Identity(".name")

	// Spec Fields
	f.SpecField(".display_name")
	f.SpecField(".product")
	f.SpecField(".license_type")
	f.SpecField(".current_billing_info")
	f.SpecField(".next_billing_info")
	f.SpecField(".labels")

	// BillingInfo subfields (UserCountBilling is SpecField; StartTime & EndTime are StatusFields in ObservedState)
	f.SpecField(".current_billing_info.user_count_billing")
	f.SpecField(".current_billing_info.user_count_billing.user_count")
	f.SpecField(".next_billing_info.user_count_billing")
	f.SpecField(".next_billing_info.user_count_billing.user_count")

	f.StatusField(".current_billing_info.start_time")
	f.StatusField(".current_billing_info.end_time")
	f.StatusField(".next_billing_info.start_time")
	f.StatusField(".next_billing_info.end_time")

	// ObservedState Status Fields
	f.StatusField(".create_time")
	f.StatusField(".update_time")
	f.StatusField(".state")

	return f
}
