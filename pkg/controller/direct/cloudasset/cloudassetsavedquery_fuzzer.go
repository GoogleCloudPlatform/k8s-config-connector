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
// proto.message: google.cloud.asset.v1.SavedQuery
// api.group: cloudasset.cnrm.cloud.google.com

package cloudasset

import (
	pb "cloud.google.com/go/asset/apiv1/assetpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(CloudAssetSavedQueryFuzzer())
}

func CloudAssetSavedQueryFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.SavedQuery{},
		CloudAssetSavedQuerySpec_FromProto, CloudAssetSavedQuerySpec_ToProto,
		CloudAssetSavedQueryObservedState_FromProto, CloudAssetSavedQueryObservedState_ToProto,
	)
	f.SpecField(".description")
	f.SpecField(".labels")
	f.SpecField(".content")

	f.StatusField(".create_time")
	f.StatusField(".creator")
	f.StatusField(".last_update_time")
	f.StatusField(".last_updater")

	f.Unimplemented_Identity(".name") // special field
	return f
}
