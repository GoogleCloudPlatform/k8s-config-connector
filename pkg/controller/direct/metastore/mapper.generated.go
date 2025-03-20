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

package metastore

import (
	pb "cloud.google.com/go/metastore/apiv1/metastorepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/metastore/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func BackendMetastore_FromProto(mapCtx *direct.MapContext, in *pb.BackendMetastore) *krm.BackendMetastore {
	if in == nil {
		return nil
	}
	out := &krm.BackendMetastore{}
	// MISSING: Name
	out.MetastoreType = direct.Enum_FromProto(mapCtx, in.GetMetastoreType())
	return out
}
func BackendMetastore_ToProto(mapCtx *direct.MapContext, in *krm.BackendMetastore) *pb.BackendMetastore {
	if in == nil {
		return nil
	}
	out := &pb.BackendMetastore{}
	// MISSING: Name
	out.MetastoreType = direct.Enum_ToProto[pb.BackendMetastore_MetastoreType](mapCtx, in.MetastoreType)
	return out
}
func MetastoreFederationObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Federation) *krm.MetastoreFederationObservedState {
	if in == nil {
		return nil
	}
	out := &krm.MetastoreFederationObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.EndpointURI = direct.LazyPtr(in.GetEndpointUri())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.StateMessage = direct.LazyPtr(in.GetStateMessage())
	out.Uid = direct.LazyPtr(in.GetUid())
	return out
}
func MetastoreFederationObservedState_ToProto(mapCtx *direct.MapContext, in *krm.MetastoreFederationObservedState) *pb.Federation {
	if in == nil {
		return nil
	}
	out := &pb.Federation{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.EndpointUri = direct.ValueOf(in.EndpointURI)
	out.State = direct.Enum_ToProto[pb.Federation_State](mapCtx, in.State)
	out.StateMessage = direct.ValueOf(in.StateMessage)
	out.Uid = direct.ValueOf(in.Uid)
	return out
}
func MetastoreFederationSpec_FromProto(mapCtx *direct.MapContext, in *pb.Federation) *krm.MetastoreFederationSpec {
	if in == nil {
		return nil
	}
	out := &krm.MetastoreFederationSpec{}
	// MISSING: Name
	out.Labels = in.Labels
	out.Version = direct.LazyPtr(in.GetVersion())
	// TODO: map type int32 message for field BackendMetastores
	return out
}
func MetastoreFederationSpec_ToProto(mapCtx *direct.MapContext, in *krm.MetastoreFederationSpec) *pb.Federation {
	if in == nil {
		return nil
	}
	out := &pb.Federation{}
	// MISSING: Name
	out.Labels = in.Labels
	out.Version = direct.ValueOf(in.Version)
	// TODO: map type int32 message for field BackendMetastores
	return out
}
