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
func ConnectorsRuntimeConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.RuntimeConfig) *krm.ConnectorsRuntimeConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ConnectorsRuntimeConfigObservedState{}
	// MISSING: LocationID
	// MISSING: ConndTopic
	// MISSING: ConndSubscription
	// MISSING: ControlPlaneTopic
	// MISSING: ControlPlaneSubscription
	// MISSING: RuntimeEndpoint
	// MISSING: State
	// MISSING: SchemaGcsBucket
	// MISSING: ServiceDirectory
	// MISSING: Name
	return out
}
func ConnectorsRuntimeConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ConnectorsRuntimeConfigObservedState) *pb.RuntimeConfig {
	if in == nil {
		return nil
	}
	out := &pb.RuntimeConfig{}
	// MISSING: LocationID
	// MISSING: ConndTopic
	// MISSING: ConndSubscription
	// MISSING: ControlPlaneTopic
	// MISSING: ControlPlaneSubscription
	// MISSING: RuntimeEndpoint
	// MISSING: State
	// MISSING: SchemaGcsBucket
	// MISSING: ServiceDirectory
	// MISSING: Name
	return out
}
func ConnectorsRuntimeConfigSpec_FromProto(mapCtx *direct.MapContext, in *pb.RuntimeConfig) *krm.ConnectorsRuntimeConfigSpec {
	if in == nil {
		return nil
	}
	out := &krm.ConnectorsRuntimeConfigSpec{}
	// MISSING: LocationID
	// MISSING: ConndTopic
	// MISSING: ConndSubscription
	// MISSING: ControlPlaneTopic
	// MISSING: ControlPlaneSubscription
	// MISSING: RuntimeEndpoint
	// MISSING: State
	// MISSING: SchemaGcsBucket
	// MISSING: ServiceDirectory
	// MISSING: Name
	return out
}
func ConnectorsRuntimeConfigSpec_ToProto(mapCtx *direct.MapContext, in *krm.ConnectorsRuntimeConfigSpec) *pb.RuntimeConfig {
	if in == nil {
		return nil
	}
	out := &pb.RuntimeConfig{}
	// MISSING: LocationID
	// MISSING: ConndTopic
	// MISSING: ConndSubscription
	// MISSING: ControlPlaneTopic
	// MISSING: ControlPlaneSubscription
	// MISSING: RuntimeEndpoint
	// MISSING: State
	// MISSING: SchemaGcsBucket
	// MISSING: ServiceDirectory
	// MISSING: Name
	return out
}
func RuntimeConfig_FromProto(mapCtx *direct.MapContext, in *pb.RuntimeConfig) *krm.RuntimeConfig {
	if in == nil {
		return nil
	}
	out := &krm.RuntimeConfig{}
	// MISSING: LocationID
	// MISSING: ConndTopic
	// MISSING: ConndSubscription
	// MISSING: ControlPlaneTopic
	// MISSING: ControlPlaneSubscription
	// MISSING: RuntimeEndpoint
	// MISSING: State
	// MISSING: SchemaGcsBucket
	// MISSING: ServiceDirectory
	// MISSING: Name
	return out
}
func RuntimeConfig_ToProto(mapCtx *direct.MapContext, in *krm.RuntimeConfig) *pb.RuntimeConfig {
	if in == nil {
		return nil
	}
	out := &pb.RuntimeConfig{}
	// MISSING: LocationID
	// MISSING: ConndTopic
	// MISSING: ConndSubscription
	// MISSING: ControlPlaneTopic
	// MISSING: ControlPlaneSubscription
	// MISSING: RuntimeEndpoint
	// MISSING: State
	// MISSING: SchemaGcsBucket
	// MISSING: ServiceDirectory
	// MISSING: Name
	return out
}
func RuntimeConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.RuntimeConfig) *krm.RuntimeConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.RuntimeConfigObservedState{}
	out.LocationID = direct.LazyPtr(in.GetLocationId())
	out.ConndTopic = direct.LazyPtr(in.GetConndTopic())
	out.ConndSubscription = direct.LazyPtr(in.GetConndSubscription())
	out.ControlPlaneTopic = direct.LazyPtr(in.GetControlPlaneTopic())
	out.ControlPlaneSubscription = direct.LazyPtr(in.GetControlPlaneSubscription())
	out.RuntimeEndpoint = direct.LazyPtr(in.GetRuntimeEndpoint())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.SchemaGcsBucket = direct.LazyPtr(in.GetSchemaGcsBucket())
	out.ServiceDirectory = direct.LazyPtr(in.GetServiceDirectory())
	out.Name = direct.LazyPtr(in.GetName())
	return out
}
func RuntimeConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.RuntimeConfigObservedState) *pb.RuntimeConfig {
	if in == nil {
		return nil
	}
	out := &pb.RuntimeConfig{}
	out.LocationId = direct.ValueOf(in.LocationID)
	out.ConndTopic = direct.ValueOf(in.ConndTopic)
	out.ConndSubscription = direct.ValueOf(in.ConndSubscription)
	out.ControlPlaneTopic = direct.ValueOf(in.ControlPlaneTopic)
	out.ControlPlaneSubscription = direct.ValueOf(in.ControlPlaneSubscription)
	out.RuntimeEndpoint = direct.ValueOf(in.RuntimeEndpoint)
	out.State = direct.Enum_ToProto[pb.RuntimeConfig_State](mapCtx, in.State)
	out.SchemaGcsBucket = direct.ValueOf(in.SchemaGcsBucket)
	out.ServiceDirectory = direct.ValueOf(in.ServiceDirectory)
	out.Name = direct.ValueOf(in.Name)
	return out
}
