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

package dialogflow

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/dialogflow/apiv2/dialogflowpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dialogflow/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func Version_FromProto(mapCtx *direct.MapContext, in *pb.Version) *krm.Version {
	if in == nil {
		return nil
	}
	out := &krm.Version{}
	// MISSING: Name
	out.Description = direct.LazyPtr(in.GetDescription())
	// MISSING: VersionNumber
	// MISSING: CreateTime
	// MISSING: Status
	return out
}
func Version_ToProto(mapCtx *direct.MapContext, in *krm.Version) *pb.Version {
	if in == nil {
		return nil
	}
	out := &pb.Version{}
	// MISSING: Name
	out.Description = direct.ValueOf(in.Description)
	// MISSING: VersionNumber
	// MISSING: CreateTime
	// MISSING: Status
	return out
}
func VersionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Version) *krm.VersionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VersionObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: Description
	out.VersionNumber = direct.LazyPtr(in.GetVersionNumber())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.Status = direct.Enum_FromProto(mapCtx, in.GetStatus())
	return out
}
func VersionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VersionObservedState) *pb.Version {
	if in == nil {
		return nil
	}
	out := &pb.Version{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: Description
	out.VersionNumber = direct.ValueOf(in.VersionNumber)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.Status = direct.Enum_ToProto[pb.Version_VersionStatus](mapCtx, in.Status)
	return out
}
