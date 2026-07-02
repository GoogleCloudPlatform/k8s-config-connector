// Copyright 2026 Google LLC
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
// proto.message: google.cloud.video.stitcher.v1.CdnKey
// api.group: videostitcher.cnrm.cloud.google.com

package videostitcher

import (
	pb "cloud.google.com/go/video/stitcher/apiv1/stitcherpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(videostitcherCDNKeyFuzzer())
}

func videostitcherCDNKeyFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.CdnKey{},
		VideoStitcherCDNKeySpec_FromProto, VideoStitcherCDNKeySpec_ToProto,
		VideoStitcherCDNKeyObservedState_FromProto, VideoStitcherCDNKeyObservedState_ToProto,
	)

	// Identity fields that are not in KRM fields
	f.Unimplemented_Identity(".name")

	// Spec fields to fuzz (including both nested objects and their leaf fields)
	f.SpecField(".hostname")

	f.SpecField(".google_cdn_key")
	f.SpecField(".google_cdn_key.private_key")
	f.SpecField(".google_cdn_key.key_name")

	f.SpecField(".akamai_cdn_key")
	f.SpecField(".akamai_cdn_key.token_key")

	f.SpecField(".media_cdn_key")
	f.SpecField(".media_cdn_key.private_key")
	f.SpecField(".media_cdn_key.key_name")
	f.SpecField(".media_cdn_key.token_config")
	f.SpecField(".media_cdn_key.token_config.query_parameter")

	return f
}
