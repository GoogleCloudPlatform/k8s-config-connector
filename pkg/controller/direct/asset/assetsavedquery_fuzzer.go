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
// proto.message: google.cloud.asset.v1.SavedQuery
// api.group: cloudasset.cnrm.cloud.google.com

package asset

import (
	pb "cloud.google.com/go/asset/apiv1/assetpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(AssetSavedQueryFuzzer())
}

func AssetSavedQueryFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.SavedQuery{},
		AssetSavedQuerySpec_FromProto, AssetSavedQuerySpec_ToProto,
		AssetSavedQueryObservedState_FromProto, AssetSavedQueryObservedState_ToProto,
	)
	f.SpecFields.Insert(".description")
	f.SpecFields.Insert(".labels")
	f.SpecFields.Insert(".content")

	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".creator")
	f.StatusFields.Insert(".last_update_time")
	f.StatusFields.Insert(".last_updater")

	f.UnimplementedFields.Insert(".name") // special field
	return f
}
