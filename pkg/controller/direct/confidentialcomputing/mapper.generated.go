// Copyright 2025 Google LLC
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

package confidentialcomputing

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/confidentialcomputing/apiv1/confidentialcomputingpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/confidentialcomputing/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func Challenge_FromProto(mapCtx *direct.MapContext, in *pb.Challenge) *krm.Challenge {
	if in == nil {
		return nil
	}
	out := &krm.Challenge{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: ExpireTime
	// MISSING: Used
	// MISSING: TpmNonce
	return out
}
func Challenge_ToProto(mapCtx *direct.MapContext, in *krm.Challenge) *pb.Challenge {
	if in == nil {
		return nil
	}
	out := &pb.Challenge{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: ExpireTime
	// MISSING: Used
	// MISSING: TpmNonce
	return out
}
func ChallengeObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Challenge) *krm.ChallengeObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ChallengeObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.ExpireTime = direct.StringTimestamp_FromProto(mapCtx, in.GetExpireTime())
	out.Used = direct.LazyPtr(in.GetUsed())
	out.TpmNonce = direct.LazyPtr(in.GetTpmNonce())
	return out
}
func ChallengeObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ChallengeObservedState) *pb.Challenge {
	if in == nil {
		return nil
	}
	out := &pb.Challenge{}
	out.Name = direct.ValueOf(in.Name)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.ExpireTime = direct.StringTimestamp_ToProto(mapCtx, in.ExpireTime)
	out.Used = direct.ValueOf(in.Used)
	out.TpmNonce = direct.ValueOf(in.TpmNonce)
	return out
}
func ConfidentialcomputingChallengeObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Challenge) *krm.ConfidentialcomputingChallengeObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ConfidentialcomputingChallengeObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: ExpireTime
	// MISSING: Used
	// MISSING: TpmNonce
	return out
}
func ConfidentialcomputingChallengeObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ConfidentialcomputingChallengeObservedState) *pb.Challenge {
	if in == nil {
		return nil
	}
	out := &pb.Challenge{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: ExpireTime
	// MISSING: Used
	// MISSING: TpmNonce
	return out
}
func ConfidentialcomputingChallengeSpec_FromProto(mapCtx *direct.MapContext, in *pb.Challenge) *krm.ConfidentialcomputingChallengeSpec {
	if in == nil {
		return nil
	}
	out := &krm.ConfidentialcomputingChallengeSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: ExpireTime
	// MISSING: Used
	// MISSING: TpmNonce
	return out
}
func ConfidentialcomputingChallengeSpec_ToProto(mapCtx *direct.MapContext, in *krm.ConfidentialcomputingChallengeSpec) *pb.Challenge {
	if in == nil {
		return nil
	}
	out := &pb.Challenge{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: ExpireTime
	// MISSING: Used
	// MISSING: TpmNonce
	return out
}
