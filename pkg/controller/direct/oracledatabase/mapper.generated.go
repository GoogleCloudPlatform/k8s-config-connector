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
func AutonomousDbVersion_FromProto(mapCtx *direct.MapContext, in *pb.AutonomousDbVersion) *krm.AutonomousDbVersion {
	if in == nil {
		return nil
	}
	out := &krm.AutonomousDbVersion{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: Version
	// MISSING: DbWorkload
	// MISSING: WorkloadURI
	return out
}
func AutonomousDbVersion_ToProto(mapCtx *direct.MapContext, in *krm.AutonomousDbVersion) *pb.AutonomousDbVersion {
	if in == nil {
		return nil
	}
	out := &pb.AutonomousDbVersion{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: Version
	// MISSING: DbWorkload
	// MISSING: WorkloadURI
	return out
}
func AutonomousDbVersionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AutonomousDbVersion) *krm.AutonomousDbVersionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AutonomousDbVersionObservedState{}
	// MISSING: Name
	out.Version = direct.LazyPtr(in.GetVersion())
	out.DbWorkload = direct.Enum_FromProto(mapCtx, in.GetDbWorkload())
	out.WorkloadURI = direct.LazyPtr(in.GetWorkloadUri())
	return out
}
func AutonomousDbVersionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AutonomousDbVersionObservedState) *pb.AutonomousDbVersion {
	if in == nil {
		return nil
	}
	out := &pb.AutonomousDbVersion{}
	// MISSING: Name
	out.Version = direct.ValueOf(in.Version)
	out.DbWorkload = direct.Enum_ToProto[pb.DBWorkload](mapCtx, in.DbWorkload)
	out.WorkloadUri = direct.ValueOf(in.WorkloadURI)
	return out
}
func OracledatabaseAutonomousDbVersionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AutonomousDbVersion) *krm.OracledatabaseAutonomousDbVersionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.OracledatabaseAutonomousDbVersionObservedState{}
	// MISSING: Name
	// MISSING: Version
	// MISSING: DbWorkload
	// MISSING: WorkloadURI
	return out
}
func OracledatabaseAutonomousDbVersionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.OracledatabaseAutonomousDbVersionObservedState) *pb.AutonomousDbVersion {
	if in == nil {
		return nil
	}
	out := &pb.AutonomousDbVersion{}
	// MISSING: Name
	// MISSING: Version
	// MISSING: DbWorkload
	// MISSING: WorkloadURI
	return out
}
func OracledatabaseAutonomousDbVersionSpec_FromProto(mapCtx *direct.MapContext, in *pb.AutonomousDbVersion) *krm.OracledatabaseAutonomousDbVersionSpec {
	if in == nil {
		return nil
	}
	out := &krm.OracledatabaseAutonomousDbVersionSpec{}
	// MISSING: Name
	// MISSING: Version
	// MISSING: DbWorkload
	// MISSING: WorkloadURI
	return out
}
func OracledatabaseAutonomousDbVersionSpec_ToProto(mapCtx *direct.MapContext, in *krm.OracledatabaseAutonomousDbVersionSpec) *pb.AutonomousDbVersion {
	if in == nil {
		return nil
	}
	out := &pb.AutonomousDbVersion{}
	// MISSING: Name
	// MISSING: Version
	// MISSING: DbWorkload
	// MISSING: WorkloadURI
	return out
}
