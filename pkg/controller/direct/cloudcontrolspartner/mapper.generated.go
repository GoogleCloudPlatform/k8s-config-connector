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

package cloudcontrolspartner

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/cloudcontrolspartner/apiv1/cloudcontrolspartnerpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/cloudcontrolspartner/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func CloudcontrolspartnerCustomerObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Customer) *krm.CloudcontrolspartnerCustomerObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CloudcontrolspartnerCustomerObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: CustomerOnboardingState
	// MISSING: IsOnboarded
	return out
}
func CloudcontrolspartnerCustomerObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CloudcontrolspartnerCustomerObservedState) *pb.Customer {
	if in == nil {
		return nil
	}
	out := &pb.Customer{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: CustomerOnboardingState
	// MISSING: IsOnboarded
	return out
}
func CloudcontrolspartnerCustomerSpec_FromProto(mapCtx *direct.MapContext, in *pb.Customer) *krm.CloudcontrolspartnerCustomerSpec {
	if in == nil {
		return nil
	}
	out := &krm.CloudcontrolspartnerCustomerSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: CustomerOnboardingState
	// MISSING: IsOnboarded
	return out
}
func CloudcontrolspartnerCustomerSpec_ToProto(mapCtx *direct.MapContext, in *krm.CloudcontrolspartnerCustomerSpec) *pb.Customer {
	if in == nil {
		return nil
	}
	out := &pb.Customer{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: CustomerOnboardingState
	// MISSING: IsOnboarded
	return out
}
func CloudcontrolspartnerEkmConnectionsObservedState_FromProto(mapCtx *direct.MapContext, in *pb.EkmConnections) *krm.CloudcontrolspartnerEkmConnectionsObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CloudcontrolspartnerEkmConnectionsObservedState{}
	// MISSING: Name
	// MISSING: EkmConnections
	return out
}
func CloudcontrolspartnerEkmConnectionsObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CloudcontrolspartnerEkmConnectionsObservedState) *pb.EkmConnections {
	if in == nil {
		return nil
	}
	out := &pb.EkmConnections{}
	// MISSING: Name
	// MISSING: EkmConnections
	return out
}
func CloudcontrolspartnerEkmConnectionsSpec_FromProto(mapCtx *direct.MapContext, in *pb.EkmConnections) *krm.CloudcontrolspartnerEkmConnectionsSpec {
	if in == nil {
		return nil
	}
	out := &krm.CloudcontrolspartnerEkmConnectionsSpec{}
	// MISSING: Name
	// MISSING: EkmConnections
	return out
}
func CloudcontrolspartnerEkmConnectionsSpec_ToProto(mapCtx *direct.MapContext, in *krm.CloudcontrolspartnerEkmConnectionsSpec) *pb.EkmConnections {
	if in == nil {
		return nil
	}
	out := &pb.EkmConnections{}
	// MISSING: Name
	// MISSING: EkmConnections
	return out
}
func CloudcontrolspartnerWorkloadObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Workload) *krm.CloudcontrolspartnerWorkloadObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CloudcontrolspartnerWorkloadObservedState{}
	// MISSING: Name
	// MISSING: FolderID
	// MISSING: CreateTime
	// MISSING: Folder
	// MISSING: WorkloadOnboardingState
	// MISSING: IsOnboarded
	// MISSING: KeyManagementProjectID
	// MISSING: Location
	// MISSING: Partner
	return out
}
func CloudcontrolspartnerWorkloadObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CloudcontrolspartnerWorkloadObservedState) *pb.Workload {
	if in == nil {
		return nil
	}
	out := &pb.Workload{}
	// MISSING: Name
	// MISSING: FolderID
	// MISSING: CreateTime
	// MISSING: Folder
	// MISSING: WorkloadOnboardingState
	// MISSING: IsOnboarded
	// MISSING: KeyManagementProjectID
	// MISSING: Location
	// MISSING: Partner
	return out
}
func CloudcontrolspartnerWorkloadSpec_FromProto(mapCtx *direct.MapContext, in *pb.Workload) *krm.CloudcontrolspartnerWorkloadSpec {
	if in == nil {
		return nil
	}
	out := &krm.CloudcontrolspartnerWorkloadSpec{}
	// MISSING: Name
	// MISSING: FolderID
	// MISSING: CreateTime
	// MISSING: Folder
	// MISSING: WorkloadOnboardingState
	// MISSING: IsOnboarded
	// MISSING: KeyManagementProjectID
	// MISSING: Location
	// MISSING: Partner
	return out
}
func CloudcontrolspartnerWorkloadSpec_ToProto(mapCtx *direct.MapContext, in *krm.CloudcontrolspartnerWorkloadSpec) *pb.Workload {
	if in == nil {
		return nil
	}
	out := &pb.Workload{}
	// MISSING: Name
	// MISSING: FolderID
	// MISSING: CreateTime
	// MISSING: Folder
	// MISSING: WorkloadOnboardingState
	// MISSING: IsOnboarded
	// MISSING: KeyManagementProjectID
	// MISSING: Location
	// MISSING: Partner
	return out
}
func EkmConnection_FromProto(mapCtx *direct.MapContext, in *pb.EkmConnection) *krm.EkmConnection {
	if in == nil {
		return nil
	}
	out := &krm.EkmConnection{}
	out.ConnectionName = direct.LazyPtr(in.GetConnectionName())
	// MISSING: ConnectionState
	out.ConnectionError = EkmConnection_ConnectionError_FromProto(mapCtx, in.GetConnectionError())
	return out
}
func EkmConnection_ToProto(mapCtx *direct.MapContext, in *krm.EkmConnection) *pb.EkmConnection {
	if in == nil {
		return nil
	}
	out := &pb.EkmConnection{}
	out.ConnectionName = direct.ValueOf(in.ConnectionName)
	// MISSING: ConnectionState
	out.ConnectionError = EkmConnection_ConnectionError_ToProto(mapCtx, in.ConnectionError)
	return out
}
func EkmConnectionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.EkmConnection) *krm.EkmConnectionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.EkmConnectionObservedState{}
	// MISSING: ConnectionName
	out.ConnectionState = direct.Enum_FromProto(mapCtx, in.GetConnectionState())
	// MISSING: ConnectionError
	return out
}
func EkmConnectionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.EkmConnectionObservedState) *pb.EkmConnection {
	if in == nil {
		return nil
	}
	out := &pb.EkmConnection{}
	// MISSING: ConnectionName
	out.ConnectionState = direct.Enum_ToProto[pb.EkmConnection_ConnectionState](mapCtx, in.ConnectionState)
	// MISSING: ConnectionError
	return out
}
func EkmConnection_ConnectionError_FromProto(mapCtx *direct.MapContext, in *pb.EkmConnection_ConnectionError) *krm.EkmConnection_ConnectionError {
	if in == nil {
		return nil
	}
	out := &krm.EkmConnection_ConnectionError{}
	out.ErrorDomain = direct.LazyPtr(in.GetErrorDomain())
	out.ErrorMessage = direct.LazyPtr(in.GetErrorMessage())
	return out
}
func EkmConnection_ConnectionError_ToProto(mapCtx *direct.MapContext, in *krm.EkmConnection_ConnectionError) *pb.EkmConnection_ConnectionError {
	if in == nil {
		return nil
	}
	out := &pb.EkmConnection_ConnectionError{}
	out.ErrorDomain = direct.ValueOf(in.ErrorDomain)
	out.ErrorMessage = direct.ValueOf(in.ErrorMessage)
	return out
}
func EkmConnections_FromProto(mapCtx *direct.MapContext, in *pb.EkmConnections) *krm.EkmConnections {
	if in == nil {
		return nil
	}
	out := &krm.EkmConnections{}
	out.Name = direct.LazyPtr(in.GetName())
	out.EkmConnections = direct.Slice_FromProto(mapCtx, in.EkmConnections, EkmConnection_FromProto)
	return out
}
func EkmConnections_ToProto(mapCtx *direct.MapContext, in *krm.EkmConnections) *pb.EkmConnections {
	if in == nil {
		return nil
	}
	out := &pb.EkmConnections{}
	out.Name = direct.ValueOf(in.Name)
	out.EkmConnections = direct.Slice_ToProto(mapCtx, in.EkmConnections, EkmConnection_ToProto)
	return out
}
func EkmConnectionsObservedState_FromProto(mapCtx *direct.MapContext, in *pb.EkmConnections) *krm.EkmConnectionsObservedState {
	if in == nil {
		return nil
	}
	out := &krm.EkmConnectionsObservedState{}
	// MISSING: Name
	out.EkmConnections = direct.Slice_FromProto(mapCtx, in.EkmConnections, EkmConnectionObservedState_FromProto)
	return out
}
func EkmConnectionsObservedState_ToProto(mapCtx *direct.MapContext, in *krm.EkmConnectionsObservedState) *pb.EkmConnections {
	if in == nil {
		return nil
	}
	out := &pb.EkmConnections{}
	// MISSING: Name
	out.EkmConnections = direct.Slice_ToProto(mapCtx, in.EkmConnections, EkmConnectionObservedState_ToProto)
	return out
}
