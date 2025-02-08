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

package security

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/security/publicca/apiv1/publiccapb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/security/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func ExternalAccountKey_FromProto(mapCtx *direct.MapContext, in *pb.ExternalAccountKey) *krm.ExternalAccountKey {
	if in == nil {
		return nil
	}
	out := &krm.ExternalAccountKey{}
	// MISSING: Name
	// MISSING: KeyID
	// MISSING: B64MacKey
	return out
}
func ExternalAccountKey_ToProto(mapCtx *direct.MapContext, in *krm.ExternalAccountKey) *pb.ExternalAccountKey {
	if in == nil {
		return nil
	}
	out := &pb.ExternalAccountKey{}
	// MISSING: Name
	// MISSING: KeyID
	// MISSING: B64MacKey
	return out
}
func ExternalAccountKeyObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ExternalAccountKey) *krm.ExternalAccountKeyObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ExternalAccountKeyObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.KeyID = direct.LazyPtr(in.GetKeyId())
	out.B64MacKey = in.GetB64MacKey()
	return out
}
func ExternalAccountKeyObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ExternalAccountKeyObservedState) *pb.ExternalAccountKey {
	if in == nil {
		return nil
	}
	out := &pb.ExternalAccountKey{}
	out.Name = direct.ValueOf(in.Name)
	out.KeyId = direct.ValueOf(in.KeyID)
	out.B64MacKey = in.B64MacKey
	return out
}
func SecurityExternalAccountKeyObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ExternalAccountKey) *krm.SecurityExternalAccountKeyObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SecurityExternalAccountKeyObservedState{}
	// MISSING: Name
	// MISSING: KeyID
	// MISSING: B64MacKey
	return out
}
func SecurityExternalAccountKeyObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SecurityExternalAccountKeyObservedState) *pb.ExternalAccountKey {
	if in == nil {
		return nil
	}
	out := &pb.ExternalAccountKey{}
	// MISSING: Name
	// MISSING: KeyID
	// MISSING: B64MacKey
	return out
}
func SecurityExternalAccountKeySpec_FromProto(mapCtx *direct.MapContext, in *pb.ExternalAccountKey) *krm.SecurityExternalAccountKeySpec {
	if in == nil {
		return nil
	}
	out := &krm.SecurityExternalAccountKeySpec{}
	// MISSING: Name
	// MISSING: KeyID
	// MISSING: B64MacKey
	return out
}
func SecurityExternalAccountKeySpec_ToProto(mapCtx *direct.MapContext, in *krm.SecurityExternalAccountKeySpec) *pb.ExternalAccountKey {
	if in == nil {
		return nil
	}
	out := &pb.ExternalAccountKey{}
	// MISSING: Name
	// MISSING: KeyID
	// MISSING: B64MacKey
	return out
}
