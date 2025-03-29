// Copyright 2024 Google LLC
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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(gkeMultiCloudAttachedClusterFuzzer())
}

func gkeMultiCloudAttachedClusterFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.AttachedCluster{},
		GkeMultiCloudAttachedClusterSpec_FromProto, GkeMultiCloudAttachedClusterSpec_ToProto,
		GkeMultiCloudAttachedClusterObservedState_FromProto, GkeMultiCloudAttachedClusterObservedState_ToProto,
	)

	f.SpecFields.Insert(".oidc_config")
	f.SpecFields.Insert(".platform_version")
	f.SpecFields.Insert(".distribution")
	f.SpecFields.Insert(".fleet")
	f.SpecFields.Insert(".logging_config")
	f.SpecFields.Insert(".authorization")
	f.SpecFields.Insert(".monitoring_config")
	f.SpecFields.Insert(".binary_authorization")
	f.SpecFields.Insert(".security_posture_config")
	f.SpecFields.Insert(".description")
	f.SpecFields.Insert(".annotations")
	f.SpecFields.Insert(".tags")

	f.StatusFields.Insert(".cluster_region")
	f.StatusFields.Insert(".fleet")
	f.StatusFields.Insert(".state")
	f.StatusFields.Insert(".uid")
	f.StatusFields.Insert(".reconciling")
	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".update_time")
	f.StatusFields.Insert(".kubernetes_version")
	f.StatusFields.Insert(".workload_identity_config")
	f.StatusFields.Insert(".errors")

	f.UnimplementedFields.Insert(".name") // special field
	f.UnimplementedFields.Insert(".proxy_config")
	f.UnimplementedFields.Insert(".etag")

	return f
}
