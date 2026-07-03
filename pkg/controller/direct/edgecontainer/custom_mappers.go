// Copyright 2026 Google LLC
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

package edgecontainer

import (
	pb "cloud.google.com/go/edgecontainer/apiv1/edgecontainerpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/edgecontainer/v1beta1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	statuspb "google.golang.org/genproto/googleapis/rpc/status"
)

func Cluster_ControlPlaneEncryption_FromProto(mapCtx *direct.MapContext, in *pb.Cluster_ControlPlaneEncryption) *krm.Cluster_ControlPlaneEncryption {
	if in == nil {
		return nil
	}
	out := &krm.Cluster_ControlPlaneEncryption{}
	if in.GetKmsKey() != "" {
		out.KMSKeyRef = &refsv1beta1.KMSCryptoKeyRef{External: in.GetKmsKey()}
	}
	out.KMSKeyActiveVersion = direct.LazyPtr(in.GetKmsKeyActiveVersion())
	out.KMSKeyState = direct.Enum_FromProto(mapCtx, in.GetKmsKeyState())
	if v := in.GetKmsStatus(); v != nil {
		out.KMSStatus = []krm.KMSStatus{*KMSStatus_FromProto(mapCtx, v)}
	}
	return out
}

func Cluster_ControlPlaneEncryption_ToProto(mapCtx *direct.MapContext, in *krm.Cluster_ControlPlaneEncryption) *pb.Cluster_ControlPlaneEncryption {
	if in == nil {
		return nil
	}
	out := &pb.Cluster_ControlPlaneEncryption{}
	if in.KMSKeyRef != nil {
		out.KmsKey = in.KMSKeyRef.External
	}
	out.KmsKeyActiveVersion = direct.ValueOf(in.KMSKeyActiveVersion)
	out.KmsKeyState = direct.Enum_ToProto[pb.KmsKeyState](mapCtx, in.KMSKeyState)
	if len(in.KMSStatus) > 0 {
		out.KmsStatus = KMSStatus_ToProto(mapCtx, &in.KMSStatus[0])
	}
	return out
}

func KMSStatus_FromProto(mapCtx *direct.MapContext, in *statuspb.Status) *krm.KMSStatus {
	if in == nil {
		return nil
	}
	out := &krm.KMSStatus{}
	out.Code = direct.LazyPtr(in.GetCode())
	out.Message = direct.LazyPtr(in.GetMessage())
	return out
}

func KMSStatus_ToProto(mapCtx *direct.MapContext, in *krm.KMSStatus) *statuspb.Status {
	if in == nil {
		return nil
	}
	out := &statuspb.Status{}
	out.Code = direct.ValueOf(in.Code)
	out.Message = direct.ValueOf(in.Message)
	return out
}
