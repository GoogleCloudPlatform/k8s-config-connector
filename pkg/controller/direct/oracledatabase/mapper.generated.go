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

package oracledatabase

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/oracledatabase/apiv1/oracledatabasepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/oracledatabase/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func AutonomousDatabaseCharacterSet_FromProto(mapCtx *direct.MapContext, in *pb.AutonomousDatabaseCharacterSet) *krm.AutonomousDatabaseCharacterSet {
	if in == nil {
		return nil
	}
	out := &krm.AutonomousDatabaseCharacterSet{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: CharacterSetType
	// MISSING: CharacterSet
	return out
}
func AutonomousDatabaseCharacterSet_ToProto(mapCtx *direct.MapContext, in *krm.AutonomousDatabaseCharacterSet) *pb.AutonomousDatabaseCharacterSet {
	if in == nil {
		return nil
	}
	out := &pb.AutonomousDatabaseCharacterSet{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: CharacterSetType
	// MISSING: CharacterSet
	return out
}
func AutonomousDatabaseCharacterSetObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AutonomousDatabaseCharacterSet) *krm.AutonomousDatabaseCharacterSetObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AutonomousDatabaseCharacterSetObservedState{}
	// MISSING: Name
	out.CharacterSetType = direct.Enum_FromProto(mapCtx, in.GetCharacterSetType())
	out.CharacterSet = direct.LazyPtr(in.GetCharacterSet())
	return out
}
func AutonomousDatabaseCharacterSetObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AutonomousDatabaseCharacterSetObservedState) *pb.AutonomousDatabaseCharacterSet {
	if in == nil {
		return nil
	}
	out := &pb.AutonomousDatabaseCharacterSet{}
	// MISSING: Name
	out.CharacterSetType = direct.Enum_ToProto[pb.AutonomousDatabaseCharacterSet_CharacterSetType](mapCtx, in.CharacterSetType)
	out.CharacterSet = direct.ValueOf(in.CharacterSet)
	return out
}
func OracledatabaseAutonomousDatabaseCharacterSetObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AutonomousDatabaseCharacterSet) *krm.OracledatabaseAutonomousDatabaseCharacterSetObservedState {
	if in == nil {
		return nil
	}
	out := &krm.OracledatabaseAutonomousDatabaseCharacterSetObservedState{}
	// MISSING: Name
	// MISSING: CharacterSetType
	// MISSING: CharacterSet
	return out
}
func OracledatabaseAutonomousDatabaseCharacterSetObservedState_ToProto(mapCtx *direct.MapContext, in *krm.OracledatabaseAutonomousDatabaseCharacterSetObservedState) *pb.AutonomousDatabaseCharacterSet {
	if in == nil {
		return nil
	}
	out := &pb.AutonomousDatabaseCharacterSet{}
	// MISSING: Name
	// MISSING: CharacterSetType
	// MISSING: CharacterSet
	return out
}
func OracledatabaseAutonomousDatabaseCharacterSetSpec_FromProto(mapCtx *direct.MapContext, in *pb.AutonomousDatabaseCharacterSet) *krm.OracledatabaseAutonomousDatabaseCharacterSetSpec {
	if in == nil {
		return nil
	}
	out := &krm.OracledatabaseAutonomousDatabaseCharacterSetSpec{}
	// MISSING: Name
	// MISSING: CharacterSetType
	// MISSING: CharacterSet
	return out
}
func OracledatabaseAutonomousDatabaseCharacterSetSpec_ToProto(mapCtx *direct.MapContext, in *krm.OracledatabaseAutonomousDatabaseCharacterSetSpec) *pb.AutonomousDatabaseCharacterSet {
	if in == nil {
		return nil
	}
	out := &pb.AutonomousDatabaseCharacterSet{}
	// MISSING: Name
	// MISSING: CharacterSetType
	// MISSING: CharacterSet
	return out
}
