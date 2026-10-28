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
// proto.message: google.api.apikeys.v2.Key
// api.group: apikeys.cnrm.cloud.google.com

package apikeys

import (
	pb "cloud.google.com/go/apikeys/apiv2/apikeyspb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(apiKeysKeyFuzzer())
}

func apiKeysKeyFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Key{},
		APIKeysKeySpec_FromProto, APIKeysKeySpec_ToProto,
		APIKeysKeyObservedState_FromProto, APIKeysKeyObservedState_ToProto,
	)

	// Documented Field Comparison (KRM Spec -> Proto):
	// - displayName -> pb.Key.display_name (f.SpecField)
	// - restrictions.browserKeyRestrictions -> pb.Key.restrictions.browser_key_restrictions (f.SpecField)
	// - restrictions.serverKeyRestrictions -> pb.Key.restrictions.server_key_restrictions (f.SpecField)
	// - restrictions.androidKeyRestrictions -> pb.Key.restrictions.android_key_restrictions (f.SpecField)
	// - restrictions.iosKeyRestrictions -> pb.Key.restrictions.ios_key_restrictions (f.SpecField)
	// - restrictions.apiTargets -> pb.Key.restrictions.api_targets (f.SpecField)
	// - projectRef -> Unimplemented_Identity (part of pb.Key.name identity)
	// - resourceID -> Unimplemented_Identity (part of pb.Key.name identity)
	f.SpecField(".display_name")
	f.SpecField(".restrictions")
	f.SpecField(".restrictions.browser_key_restrictions.allowed_referrers")
	f.SpecField(".restrictions.server_key_restrictions.allowed_ips")
	f.SpecField(".restrictions.android_key_restrictions.allowed_applications")
	f.SpecField(".restrictions.ios_key_restrictions.allowed_bundle_ids")
	f.SpecField(".restrictions.api_targets")

	// Documented Field Comparison (KRM Status/ObservedState -> Proto):
	// - uid -> pb.Key.uid (f.StatusField)
	f.StatusField(".uid")

	// Identity field
	f.Unimplemented_Identity(".name")

	// Output-only / unmapped proto fields
	f.Unimplemented_NotYetTriaged(".key_string")
	f.Unimplemented_NotYetTriaged(".create_time")
	f.Unimplemented_NotYetTriaged(".update_time")
	f.Unimplemented_NotYetTriaged(".delete_time")
	f.Unimplemented_LabelsAnnotations(".annotations")
	f.Unimplemented_Etag()

	return f
}
