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

package managedkafka

import (
	pb "cloud.google.com/go/managedkafka/apiv1/managedkafkapb"
	kmsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/managedkafka/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func GcpConfig_FromProto(mapCtx *direct.MapContext, in *pb.GcpConfig) *krm.GcpConfig {
	if in == nil {
		return nil
	}
	out := &krm.GcpConfig{}
	out.AccessConfig = AccessConfig_FromProto(mapCtx, in.GetAccessConfig())
	if in.GetKmsKey() != "" {
		out.KmsKeyRef = &kmsv1beta1.KMSCryptoKeyRef{External: in.GetKmsKey()}
	}
	return out
}
func GcpConfig_ToProto(mapCtx *direct.MapContext, in *krm.GcpConfig) *pb.GcpConfig {
	if in == nil {
		return nil
	}
	out := &pb.GcpConfig{}
	out.AccessConfig = AccessConfig_ToProto(mapCtx, in.AccessConfig)
	if in.KmsKeyRef != nil {
		out.KmsKey = in.KmsKeyRef.External
	}
	return out
}
func NetworkConfig_FromProto(mapCtx *direct.MapContext, in *pb.NetworkConfig) *krm.NetworkConfig {
	if in == nil {
		return nil
	}
	out := &krm.NetworkConfig{}
	if in.GetSubnet() != "" {
		out.SubnetworkRef = &refs.ComputeSubnetworkRef{External: in.GetSubnet()}
	}
	return out
}
func NetworkConfig_ToProto(mapCtx *direct.MapContext, in *krm.NetworkConfig) *pb.NetworkConfig {
	if in == nil {
		return nil
	}
	out := &pb.NetworkConfig{}
	if in.SubnetworkRef != nil {
		out.Subnet = in.SubnetworkRef.External
	}
	return out
}
