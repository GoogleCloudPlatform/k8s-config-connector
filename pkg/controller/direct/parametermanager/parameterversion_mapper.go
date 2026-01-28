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

// +generated:mapper
// krm.group: parametermanager.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.parametermanager.v1

package parametermanager

import (
	pb "cloud.google.com/go/parametermanager/apiv1/parametermanagerpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/parametermanager/v1alpha1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func ParameterVersionPayload_FromProto(mapCtx *direct.MapContext, in *pb.ParameterVersionPayload) *krm.ParameterVersionPayload {
	if in == nil {
		return nil
	}
	out := &krm.ParameterVersionPayload{}
	out.Data = in.GetData()
	return out
}
func ParameterVersionPayload_ToProto(mapCtx *direct.MapContext, in *krm.ParameterVersionPayload) *pb.ParameterVersionPayload {
	if in == nil {
		return nil
	}
	out := &pb.ParameterVersionPayload{}
	if len(in.Data) > 0 {
		out.Data = in.Data
	}
	return out
}

func ParameterManagerParameterVersionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ParameterVersion) *krm.ParameterManagerParameterVersionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ParameterManagerParameterVersionObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	if in.GetKmsKeyVersion() != "" {
		out.KMSKeyVersion = &refsv1beta1.KMSCryptoKeyRef{External: in.GetKmsKeyVersion()}
	}
	return out
}

func ParameterManagerParameterVersionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ParameterManagerParameterVersionObservedState) *pb.ParameterVersion {
	if in == nil {
		return nil
	}
	out := &pb.ParameterVersion{}
	out.Name = direct.ValueOf(in.Name)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	if in.KMSKeyVersion != nil {
		out.KmsKeyVersion = &in.KMSKeyVersion.External
	}
	return out
}
