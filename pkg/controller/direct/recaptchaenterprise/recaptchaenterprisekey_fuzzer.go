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
// proto.message: google.cloud.recaptchaenterprise.v1.Key
// api.group: recaptchaenterprise.cnrm.cloud.google.com

package recaptchaenterprise

import (
	pb "cloud.google.com/go/recaptchaenterprise/v2/apiv1/recaptchaenterprisepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(recaptchaEnterpriseKeyFuzzer())
}

func recaptchaEnterpriseKeyFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Key{},
		RecaptchaEnterpriseKeySpec_FromProto, RecaptchaEnterpriseKeySpec_ToProto,
		RecaptchaEnterpriseKeyStatus_FromProto, RecaptchaEnterpriseKeyStatus_ToProto,
	)

	// Explicit comparison of KRM Spec fields with proto fields:
	// - displayName maps to .display_name
	// - webSettings maps to .web_settings
	// - androidSettings maps to .android_settings
	// - iosSettings maps to .ios_settings
	// - testingOptions maps to .testing_options
	// - wafSettings maps to .waf_settings
	// - projectRef and resourceID map to identity .name
	f.SpecField(".display_name")
	f.SpecField(".web_settings")
	f.SpecField(".android_settings")
	f.SpecField(".ios_settings")
	f.SpecField(".testing_options")
	f.SpecField(".waf_settings")

	// Explicit comparison of KRM Status fields with proto fields:
	// - createTime maps to .create_time
	f.StatusField(".create_time")

	// Identity field .name mapped from projectRef and resourceID
	f.Unimplemented_Identity(".name")

	// Unsupported or unimplemented fields
	f.Unimplemented_NotYetTriaged(".express_settings")
	f.Unimplemented_NotYetTriaged(".labels")
	f.Unimplemented_NotYetTriaged(".android_settings.support_non_google_app_store_distribution")
	f.Unimplemented_NotYetTriaged(".ios_settings.apple_developer_id")
	f.Unimplemented_NotYetTriaged(".web_settings.challenge_settings")

	return f
}
