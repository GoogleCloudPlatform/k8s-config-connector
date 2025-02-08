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

package servicedirectory

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/servicedirectory/apiv1beta1/servicedirectorypb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/servicedirectory/v1alpha1"
)
func Endpoint_FromProto(mapCtx *direct.MapContext, in *pb.Endpoint) *krm.Endpoint {
	if in == nil {
		return nil
	}
	out := &krm.Endpoint{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Address = direct.LazyPtr(in.GetAddress())
	out.Port = direct.LazyPtr(in.GetPort())
	out.Metadata = in.Metadata
	out.Network = direct.LazyPtr(in.GetNetwork())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Uid
	return out
}
func Endpoint_ToProto(mapCtx *direct.MapContext, in *krm.Endpoint) *pb.Endpoint {
	if in == nil {
		return nil
	}
	out := &pb.Endpoint{}
	out.Name = direct.ValueOf(in.Name)
	out.Address = direct.ValueOf(in.Address)
	out.Port = direct.ValueOf(in.Port)
	out.Metadata = in.Metadata
	out.Network = direct.ValueOf(in.Network)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Uid
	return out
}
func EndpointObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Endpoint) *krm.EndpointObservedState {
	if in == nil {
		return nil
	}
	out := &krm.EndpointObservedState{}
	// MISSING: Name
	// MISSING: Address
	// MISSING: Port
	// MISSING: Metadata
	// MISSING: Network
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.Uid = direct.LazyPtr(in.GetUid())
	return out
}
func EndpointObservedState_ToProto(mapCtx *direct.MapContext, in *krm.EndpointObservedState) *pb.Endpoint {
	if in == nil {
		return nil
	}
	out := &pb.Endpoint{}
	// MISSING: Name
	// MISSING: Address
	// MISSING: Port
	// MISSING: Metadata
	// MISSING: Network
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.Uid = direct.ValueOf(in.Uid)
	return out
}
