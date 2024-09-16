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

package firestore

/*
import (
	pb "cloud.google.com/go/firestore/apiv1/admin/adminpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/firestore/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func FirestoreDatabaseObservedState_ToProto(mapCtx *direct.MapContext, in *krm.FirestoreDatabaseObservedState) *pb.Database {
	if in == nil {
		return nil
	}
	out := &pb.Database{}
	// MISSING: Name
	out.Uid = direct.ValueOf(in.Uid)
	out.CreateTime = Database_CreateTime_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = Database_UpdateTime_ToProto(mapCtx, in.UpdateTime)
	// MISSING: LocationID
	// MISSING: Type
	// MISSING: ConcurrencyMode
	out.VersionRetentionPeriod = Database_VersionRetentionPeriod_ToProto(mapCtx, in.VersionRetentionPeriod)
	out.EarliestVersionTime = Database_EarliestVersionTime_ToProto(mapCtx, in.EarliestVersionTime)
	// MISSING: PointInTimeRecoveryEnablement
	// MISSING: AppEngineIntegrationMode
	out.KeyPrefix = direct.ValueOf(in.KeyPrefix)
	// MISSING: DeleteProtectionState
	out.Etag = direct.ValueOf(in.Etag)
	return out
}
*/
