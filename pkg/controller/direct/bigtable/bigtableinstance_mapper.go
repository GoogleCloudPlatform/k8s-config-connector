// Copyright 2024 Google LLC
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

package bigtable

import (
	pb "cloud.google.com/go/bigtable/admin/apiv2/adminpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func AppProfile_DataBoostIsolationReadOnly_ComputeBillingOwner_ToProto(mapCtx *direct.MapContext, in *string) *pb.AppProfile_DataBoostIsolationReadOnly_ComputeBillingOwner {
	if in == nil {
		return nil
	}
	v := direct.Enum_ToProto[pb.AppProfile_DataBoostIsolationReadOnly_ComputeBillingOwner](mapCtx, in)
	return &v
}
