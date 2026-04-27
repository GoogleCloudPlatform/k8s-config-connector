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

package bigtable

import (
	pb "cloud.google.com/go/bigtable/admin/apiv2/adminpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigtable/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func BigtableMaterializedViewObservedState_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.MaterializedView) *krm.BigtableMaterializedViewObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BigtableMaterializedViewObservedState{}
	// MISSING: Name
	// MISSING: Etag
	return out
}

func BigtableMaterializedViewObservedState_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.BigtableMaterializedViewObservedState) *pb.MaterializedView {
	if in == nil {
		return nil
	}
	out := &pb.MaterializedView{}
	// MISSING: Name
	// MISSING: Etag
	return out
}

func BigtableMaterializedViewSpec_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.MaterializedView) *krm.BigtableMaterializedViewSpec {
	if in == nil {
		return nil
	}
	out := &krm.BigtableMaterializedViewSpec{}
	// MISSING: Name
	out.Query = direct.LazyPtr(in.GetQuery())
	// MISSING: Etag
	out.DeletionProtection = direct.LazyPtr(in.GetDeletionProtection())
	return out
}

func BigtableMaterializedViewSpec_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.BigtableMaterializedViewSpec) *pb.MaterializedView {
	if in == nil {
		return nil
	}
	out := &pb.MaterializedView{}
	// MISSING: Name
	out.Query = direct.ValueOf(in.Query)
	// MISSING: Etag
	out.DeletionProtection = direct.ValueOf(in.DeletionProtection)
	return out
}
