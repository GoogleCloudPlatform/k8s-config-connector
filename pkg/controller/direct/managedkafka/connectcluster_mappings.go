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

package managedkafka

import (
	pb "cloud.google.com/go/managedkafka/apiv1/managedkafkapb"
	krmmanagedkafkav1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/managedkafka/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func ConnectGCPConfig_SecretPaths_FromProto(mapCtx *direct.MapContext, in []string) []refs.SecretManagerSecretVersionRef {
	if in == nil {
		return nil
	}
	out := make([]refs.SecretManagerSecretVersionRef, len(in))
	for i, v := range in {
		out[i] = refs.SecretManagerSecretVersionRef{External: v}
	}
	return out
}

func ConnectGCPConfig_SecretPaths_ToProto(mapCtx *direct.MapContext, in []refs.SecretManagerSecretVersionRef) []string {
	if in == nil {
		return nil
	}
	out := make([]string, len(in))
	for i, v := range in {
		out[i] = v.External
	}
	return out
}

func ManagedKafkaConnectClusterSpec_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.ConnectCluster) *krmmanagedkafkav1alpha1.ManagedKafkaConnectClusterSpec {
	if in == nil {
		return nil
	}
	out := &krmmanagedkafkav1alpha1.ManagedKafkaConnectClusterSpec{}
	out.GCPConfig = ConnectGCPConfig_v1alpha1_FromProto(mapCtx, in.GetGcpConfig())
	if in.GetKafkaCluster() != "" {
		out.ClusterRef = &krmmanagedkafkav1alpha1.ClusterRef{External: in.GetKafkaCluster()}
	}
	out.Labels = in.Labels
	out.CapacityConfig = CapacityConfig_v1alpha1_FromProto(mapCtx, in.GetCapacityConfig())
	out.Config = in.Config
	return out
}

func ManagedKafkaConnectClusterSpec_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmmanagedkafkav1alpha1.ManagedKafkaConnectClusterSpec) *pb.ConnectCluster {
	if in == nil {
		return nil
	}
	out := &pb.ConnectCluster{}
	if in.GCPConfig != nil {
		if gcpConfig := ConnectGCPConfig_v1alpha1_ToProto(mapCtx, in.GCPConfig); gcpConfig != nil {
			out.PlatformConfig = &pb.ConnectCluster_GcpConfig{GcpConfig: gcpConfig}
		}
	}
	if in.ClusterRef != nil {
		out.KafkaCluster = in.ClusterRef.External
	}
	out.Labels = in.Labels
	out.CapacityConfig = CapacityConfig_v1alpha1_ToProto(mapCtx, in.CapacityConfig)
	out.Config = in.Config
	return out
}
