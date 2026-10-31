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
// proto.message: google.cloud.compute.v1.NetworkEndpointGroup
// api.group: compute.cnrm.cloud.google.com

package compute

import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMSpecFuzzer(computeGlobalNetworkEndpointGroupFuzzer())
}

func computeGlobalNetworkEndpointGroupFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedSpecFuzzer[*pb.NetworkEndpointGroup, krm.ComputeGlobalNetworkEndpointGroupSpec](&pb.NetworkEndpointGroup{},
		ComputeGlobalNetworkEndpointGroupSpec_v1alpha1_FromProto, ComputeGlobalNetworkEndpointGroupSpec_v1alpha1_ToProto,
	)

	// Spec fields
	f.SpecField(".default_port")
	f.SpecField(".description")
	f.SpecField(".network_endpoint_type")

	// Unimplemented / identity fields
	f.Unimplemented_Identity(".name")

	// Other fields in NetworkEndpointGroup that are not used by ComputeGlobalNetworkEndpointGroup
	f.Unimplemented_Internal(".annotations")
	f.Unimplemented_Internal(".app_engine")
	f.Unimplemented_Internal(".cloud_function")
	f.Unimplemented_Internal(".cloud_run")
	f.Unimplemented_Internal(".creation_timestamp")
	f.Unimplemented_Internal(".id")
	f.Unimplemented_Internal(".kind")
	f.Unimplemented_Internal(".network")
	f.Unimplemented_Internal(".psc_data")
	f.Unimplemented_Internal(".psc_target_service")
	f.Unimplemented_Internal(".region")
	f.Unimplemented_Internal(".self_link")
	f.Unimplemented_Internal(".size")
	f.Unimplemented_Internal(".subnetwork")
	f.Unimplemented_Internal(".zone")

	return f
}
