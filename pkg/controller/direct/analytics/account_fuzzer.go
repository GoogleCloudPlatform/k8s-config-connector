// Copyright 2025 Google LLC
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

package analytics

import (
	pb "cloud.google.com/go/analytics/admin/apiv1alpha/adminpb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(analyticsAccountFuzzer())
}

func analyticsAccountFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Account{},
		AnalyticsAccountSpec_FromProto, AnalyticsAccountSpec_ToProto,
		AnalyticsAccountObservedState_FromProto, AnalyticsAccountObservedState_ToProto,
	)

	f.UnimplementedFields.Insert(".name")

	f.SpecFields.Insert(".display_name")
	f.SpecFields.Insert(".region_code")

	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".update_time")
	f.StatusFields.Insert(".deleted")
	f.StatusFields.Insert(".gmp_organization")

	return f
}
