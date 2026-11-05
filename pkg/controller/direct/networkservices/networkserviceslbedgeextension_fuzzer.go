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
// proto.message: google.cloud.networkservices.v1.LbEdgeExtension
// api.group: networkservices.cnrm.cloud.google.com

package networkservices

import (
	pb "cloud.google.com/go/networkservices/apiv1/networkservicespb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networkservices/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(lbEdgeExtensionFuzzer())
}

func lbEdgeExtensionFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer[*pb.LbEdgeExtension, krm.NetworkServicesLBEdgeExtensionSpec, krm.NetworkServicesLBEdgeExtensionObservedState](
		&pb.LbEdgeExtension{},
		NetworkServicesLBEdgeExtensionSpec_FromProto,
		NetworkServicesLBEdgeExtensionSpec_ToProto,
		NetworkServicesLBEdgeExtensionObservedState_FromProto,
		NetworkServicesLBEdgeExtensionObservedState_ToProto,
	)

	f.SpecField(".description")
	f.SpecField(".forwarding_rules")
	f.SpecField(".extension_chains")
	f.SpecField(".load_balancing_scheme")

	f.StatusField(".create_time")
	f.StatusField(".update_time")

	f.IdentityField(".name")
	f.Unimplemented_LabelsAnnotations(".labels")

	f.SpecField(".extension_chains[].extensions[].service")

	return f
}
