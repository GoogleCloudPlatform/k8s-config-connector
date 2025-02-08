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

package connectors

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/connectors/apiv1/connectorspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/connectors/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func ConnectorsSettingsObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Settings) *krm.ConnectorsSettingsObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ConnectorsSettingsObservedState{}
	// MISSING: Name
	// MISSING: Vpcsc
	// MISSING: Payg
	return out
}
func ConnectorsSettingsObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ConnectorsSettingsObservedState) *pb.Settings {
	if in == nil {
		return nil
	}
	out := &pb.Settings{}
	// MISSING: Name
	// MISSING: Vpcsc
	// MISSING: Payg
	return out
}
func ConnectorsSettingsSpec_FromProto(mapCtx *direct.MapContext, in *pb.Settings) *krm.ConnectorsSettingsSpec {
	if in == nil {
		return nil
	}
	out := &krm.ConnectorsSettingsSpec{}
	// MISSING: Name
	// MISSING: Vpcsc
	// MISSING: Payg
	return out
}
func ConnectorsSettingsSpec_ToProto(mapCtx *direct.MapContext, in *krm.ConnectorsSettingsSpec) *pb.Settings {
	if in == nil {
		return nil
	}
	out := &pb.Settings{}
	// MISSING: Name
	// MISSING: Vpcsc
	// MISSING: Payg
	return out
}
func Settings_FromProto(mapCtx *direct.MapContext, in *pb.Settings) *krm.Settings {
	if in == nil {
		return nil
	}
	out := &krm.Settings{}
	// MISSING: Name
	out.Vpcsc = direct.LazyPtr(in.GetVpcsc())
	// MISSING: Payg
	return out
}
func Settings_ToProto(mapCtx *direct.MapContext, in *krm.Settings) *pb.Settings {
	if in == nil {
		return nil
	}
	out := &pb.Settings{}
	// MISSING: Name
	out.Vpcsc = direct.ValueOf(in.Vpcsc)
	// MISSING: Payg
	return out
}
func SettingsObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Settings) *krm.SettingsObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SettingsObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: Vpcsc
	out.Payg = direct.LazyPtr(in.GetPayg())
	return out
}
func SettingsObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SettingsObservedState) *pb.Settings {
	if in == nil {
		return nil
	}
	out := &pb.Settings{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: Vpcsc
	out.Payg = direct.ValueOf(in.Payg)
	return out
}
