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
func GiVersion_FromProto(mapCtx *direct.MapContext, in *pb.GiVersion) *krm.GiVersion {
	if in == nil {
		return nil
	}
	out := &krm.GiVersion{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Version = direct.LazyPtr(in.GetVersion())
	return out
}
func GiVersion_ToProto(mapCtx *direct.MapContext, in *krm.GiVersion) *pb.GiVersion {
	if in == nil {
		return nil
	}
	out := &pb.GiVersion{}
	out.Name = direct.ValueOf(in.Name)
	out.Version = direct.ValueOf(in.Version)
	return out
}
func OracledatabaseGiVersionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.GiVersion) *krm.OracledatabaseGiVersionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.OracledatabaseGiVersionObservedState{}
	// MISSING: Name
	// MISSING: Version
	return out
}
func OracledatabaseGiVersionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.OracledatabaseGiVersionObservedState) *pb.GiVersion {
	if in == nil {
		return nil
	}
	out := &pb.GiVersion{}
	// MISSING: Name
	// MISSING: Version
	return out
}
func OracledatabaseGiVersionSpec_FromProto(mapCtx *direct.MapContext, in *pb.GiVersion) *krm.OracledatabaseGiVersionSpec {
	if in == nil {
		return nil
	}
	out := &krm.OracledatabaseGiVersionSpec{}
	// MISSING: Name
	// MISSING: Version
	return out
}
func OracledatabaseGiVersionSpec_ToProto(mapCtx *direct.MapContext, in *krm.OracledatabaseGiVersionSpec) *pb.GiVersion {
	if in == nil {
		return nil
	}
	out := &pb.GiVersion{}
	// MISSING: Name
	// MISSING: Version
	return out
}
