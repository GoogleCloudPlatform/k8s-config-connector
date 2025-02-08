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

package gsuiteaddons

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/gsuiteaddons/apiv1/gsuiteaddonspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/gsuiteaddons/v1alpha1"
)
func GsuiteaddonsInstallStatusObservedState_FromProto(mapCtx *direct.MapContext, in *pb.InstallStatus) *krm.GsuiteaddonsInstallStatusObservedState {
	if in == nil {
		return nil
	}
	out := &krm.GsuiteaddonsInstallStatusObservedState{}
	// MISSING: Name
	// MISSING: Installed
	return out
}
func GsuiteaddonsInstallStatusObservedState_ToProto(mapCtx *direct.MapContext, in *krm.GsuiteaddonsInstallStatusObservedState) *pb.InstallStatus {
	if in == nil {
		return nil
	}
	out := &pb.InstallStatus{}
	// MISSING: Name
	// MISSING: Installed
	return out
}
func GsuiteaddonsInstallStatusSpec_FromProto(mapCtx *direct.MapContext, in *pb.InstallStatus) *krm.GsuiteaddonsInstallStatusSpec {
	if in == nil {
		return nil
	}
	out := &krm.GsuiteaddonsInstallStatusSpec{}
	// MISSING: Name
	// MISSING: Installed
	return out
}
func GsuiteaddonsInstallStatusSpec_ToProto(mapCtx *direct.MapContext, in *krm.GsuiteaddonsInstallStatusSpec) *pb.InstallStatus {
	if in == nil {
		return nil
	}
	out := &pb.InstallStatus{}
	// MISSING: Name
	// MISSING: Installed
	return out
}
func InstallStatus_FromProto(mapCtx *direct.MapContext, in *pb.InstallStatus) *krm.InstallStatus {
	if in == nil {
		return nil
	}
	out := &krm.InstallStatus{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Installed = direct.BoolValue_FromProto(mapCtx, in.GetInstalled())
	return out
}
func InstallStatus_ToProto(mapCtx *direct.MapContext, in *krm.InstallStatus) *pb.InstallStatus {
	if in == nil {
		return nil
	}
	out := &pb.InstallStatus{}
	out.Name = direct.ValueOf(in.Name)
	out.Installed = direct.BoolValue_ToProto(mapCtx, in.Installed)
	return out
}
