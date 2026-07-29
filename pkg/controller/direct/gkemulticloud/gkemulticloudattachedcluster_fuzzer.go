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

// +tool:fuzz-gen
// proto.message: google.cloud.gkemulticloud.v1.AttachedCluster
// api.group: gkemulticloud.cnrm.cloud.google.com

package gkemulticloud

import (
	pb "cloud.google.com/go/gkemulticloud/apiv1/gkemulticloudpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/gkemulticloud/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(gkeMulticloudAttachedClusterFuzzer())
}

func GKEMulticloudAttachedClusterStatus_FromProto(mapCtx *direct.MapContext, in *pb.AttachedCluster) *krm.GKEMulticloudAttachedClusterStatus {
	out := &krm.GKEMulticloudAttachedClusterStatus{}
	out.ObservedState = GKEMulticloudAttachedClusterObservedState_FromProto(mapCtx, in)
	return out
}

func GKEMulticloudAttachedClusterStatus_ToProto(mapCtx *direct.MapContext, in *krm.GKEMulticloudAttachedClusterStatus) *pb.AttachedCluster {
	if in == nil {
		return nil
	}
	return GKEMulticloudAttachedClusterObservedState_ToProto(mapCtx, in.ObservedState)
}

func gkeMulticloudAttachedClusterFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.AttachedCluster{},
		GKEMulticloudAttachedClusterSpec_FromProto,
		GKEMulticloudAttachedClusterSpec_ToProto,
		GKEMulticloudAttachedClusterStatus_FromProto,
		GKEMulticloudAttachedClusterStatus_ToProto,
	)

	// Spec fields
	f.SpecField(".description")
	f.SpecField(".oidc_config")
	f.SpecField(".platform_version")
	f.SpecField(".distribution")
	f.SpecField(".fleet")
	f.SpecField(".annotations")
	f.SpecField(".logging_config")
	f.SpecField(".authorization")
	f.SpecField(".monitoring_config")
	f.SpecField(".proxy_config")
	f.SpecField(".binary_authorization")
	f.SpecField(".security_posture_config")
	f.SpecField(".tags")

	// Observed fields
	f.StatusField(".cluster_region")
	f.StatusField(".fleet")
	f.StatusField(".state")
	f.StatusField(".uid")
	f.StatusField(".reconciling")
	f.StatusField(".create_time")
	f.StatusField(".update_time")
	f.StatusField(".kubernetes_version")
	f.StatusField(".workload_identity_config")
	f.StatusField(".errors")

	// Unimplemented fields
	f.Unimplemented_Identity(".name")
	f.Unimplemented_NotYetTriaged(".etag")
	f.Unimplemented_NotYetTriaged(".system_components_config")

	return f
}
